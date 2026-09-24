// Custom code. This file is not generated and is preserved across codegen runs.
// It isolates hand-written *AndPoll convenience methods from generated code to
// eliminate merge conflicts.

package storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/polling"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/tidwall/sjson"
)

// NewAndPoll creates a new S3-compatible storage and polls until provisioning is
// complete, returning the create response (with one-time access keys preserved)
// and the provisioning status promoted to "active". Polling reuses the
// polling_interval_seconds and polling_timeout_seconds client options.
//
// Callers that capture the raw HTTP response via option.WithResponseBodyInto
// (the gcore-terraform provider pattern) get a body whose provisioning_status
// reflects the polled state, with every other byte — including access_keys —
// preserved verbatim from the original POST. Callers that pass no extra
// options see no behavioral change vs. plain New plus a manual wait.
func (r *ObjectStorageService) NewAndPoll(ctx context.Context, body ObjectStorageNewParams, opts ...option.RequestOption) (res *S3StorageCreated, err error) {
	var raw *http.Response
	actionOpts := slices.Concat(opts, []option.RequestOption{option.WithResponseInto(&raw)})

	created, err := r.New(ctx, body, actionOpts...)
	if err != nil {
		return nil, err
	}

	var rawBytes []byte
	if created == nil {
		// Caller's WithResponseBodyInto overrode the default deserialization
		// target (terraform provider pattern). Recover the typed struct from
		// the raw response so polling can proceed; supports **http.Response
		// and *[]byte body shapes. The one-time AccessKeys are only present
		// here — they are never replayed on subsequent Gets.
		created = &S3StorageCreated{}
		rawBytes, err = polling.RecoverActionBody(raw, created, opts...)
		if err != nil {
			return nil, fmt.Errorf("object storage NewAndPoll: %w", err)
		}
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return nil, err
	}
	pollingInterval := time.Duration(precfg.PollingIntervalSeconds) * time.Second
	if pollingInterval < time.Second {
		pollingInterval = time.Second
	}

	pollingCtx := ctx
	var cancel context.CancelFunc
	if precfg.PollingTimeoutSeconds > 0 {
		pollingTimeout := time.Duration(precfg.PollingTimeoutSeconds) * time.Second
		pollingCtx, cancel = context.WithTimeout(ctx, pollingTimeout)
		defer cancel()
	}

	// Exclude WithResponseBodyInto and clear request body for intermediate Gets (S3Storage must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)

	for {
		s3, err := r.Get(pollingCtx, created.ID, pollOpts...)
		if err != nil {
			return nil, fmt.Errorf("failed to get object storage status: %w", err)
		}

		if s3.ProvisioningStatus == S3StorageProvisioningStatusActive {
			// Only provisioning_status changes between the POST response and
			// the active state. address / full_name are deterministic
			// (location-derived hostname and {client_id}-{name}) and set
			// synchronously on the create response; the rest (id, name,
			// location_name, created_at, access_keys) is immutable per the
			// OAS.
			created.ProvisioningStatus = S3StorageCreatedProvisioningStatus(s3.ProvisioningStatus)
			if rawBytes != nil && raw != nil {
				// Caller captured the response via WithResponseBodyInto; rewrite the
				// body so it reflects the polled provisioning_status while preserving
				// every other byte, including the one-time access_keys.
				enriched, sjErr := sjson.SetBytes(rawBytes, "provisioning_status", string(s3.ProvisioningStatus))
				if sjErr != nil {
					return nil, fmt.Errorf("failed to enrich object storage create response with polled provisioning_status: %w", sjErr)
				}
				raw.Body = io.NopCloser(bytes.NewReader(enriched))
				polling.WriteResponseBodyInto(opts, enriched)
			}
			return created, nil
		}

		if s3.ProvisioningStatus == S3StorageProvisioningStatusDeleting ||
			s3.ProvisioningStatus == S3StorageProvisioningStatusDeleted {
			return nil, fmt.Errorf("object storage %d entered terminal state %q during creation", created.ID, s3.ProvisioningStatus)
		}

		// check if the context is done before sleeping
		select {
		// handles both timeout and cancellation
		case <-pollingCtx.Done():
			return nil, pollingCtx.Err()
		case <-time.After(pollingInterval):
		}
	}
}

// DeleteAndPoll deletes an S3-compatible storage and polls until the storage is
// gone — either Get returns 404 or provisioning_status reaches "deleted". Polling
// reuses the polling_interval_seconds and polling_timeout_seconds client options.
//
// Returns nil on confirmed deletion. Returns an error if Delete itself fails or
// polling times out / is cancelled.
func (r *ObjectStorageService) DeleteAndPoll(ctx context.Context, storageID int64, opts ...option.RequestOption) error {
	if err := r.Delete(ctx, storageID, opts...); err != nil {
		return err
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return err
	}
	pollingInterval := time.Duration(precfg.PollingIntervalSeconds) * time.Second
	if pollingInterval < time.Second {
		pollingInterval = time.Second
	}

	pollingCtx := ctx
	var cancel context.CancelFunc
	if precfg.PollingTimeoutSeconds > 0 {
		pollingTimeout := time.Duration(precfg.PollingTimeoutSeconds) * time.Second
		pollingCtx, cancel = context.WithTimeout(ctx, pollingTimeout)
		defer cancel()
	}

	// Exclude WithResponseBodyInto and clear request body for intermediate Gets (S3Storage must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)

	for {
		s3, err := r.Get(pollingCtx, storageID, pollOpts...)
		if err != nil {
			var apiErr *Error
			if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
				return nil
			}
			return fmt.Errorf("failed to get object storage status: %w", err)
		}

		if s3.ProvisioningStatus == S3StorageProvisioningStatusDeleted {
			return nil
		}

		// check if the context is done before sleeping
		select {
		// handles both timeout and cancellation
		case <-pollingCtx.Done():
			return pollingCtx.Err()
		case <-time.After(pollingInterval):
		}
	}
}

// RestoreAndPoll restores a previously deleted S3-compatible storage (within the
// 2-week window) and polls until provisioning_status reaches "active". Polling
// reuses the polling_interval_seconds and polling_timeout_seconds client options.
//
// Returns nil on confirmed restore. Returns an error if Restore itself fails,
// the storage transitions to deleting/deleted during restore, or polling times
// out / is cancelled. The signature mirrors [ObjectStorageService.Restore] —
// callers that need the active S3Storage afterwards can call
// [ObjectStorageService.Get], or pass option.WithResponseBodyInto to capture it
// from the final replay Get below.
func (r *ObjectStorageService) RestoreAndPoll(ctx context.Context, storageID int64, opts ...option.RequestOption) error {
	if err := r.Restore(ctx, storageID, opts...); err != nil {
		return err
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return err
	}
	pollingInterval := time.Duration(precfg.PollingIntervalSeconds) * time.Second
	if pollingInterval < time.Second {
		pollingInterval = time.Second
	}

	pollingCtx := ctx
	var cancel context.CancelFunc
	if precfg.PollingTimeoutSeconds > 0 {
		pollingTimeout := time.Duration(precfg.PollingTimeoutSeconds) * time.Second
		pollingCtx, cancel = context.WithTimeout(ctx, pollingTimeout)
		defer cancel()
	}

	// Exclude WithResponseBodyInto and clear request body for intermediate Gets (S3Storage must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)

	for {
		s3, err := r.Get(pollingCtx, storageID, pollOpts...)
		if err != nil {
			return fmt.Errorf("failed to get object storage status: %w", err)
		}

		if s3.ProvisioningStatus == S3StorageProvisioningStatusActive {
			// Restore returns no action body, so the caller's WithResponseBodyInto
			// (terraform provider pattern) has nothing to capture from the POST.
			// Replay the final Get with the caller's full opts so their captured
			// response is populated with the active S3Storage payload.
			_, err := r.Get(pollingCtx, storageID, opts...)
			return err
		}

		if s3.ProvisioningStatus == S3StorageProvisioningStatusDeleting ||
			s3.ProvisioningStatus == S3StorageProvisioningStatusDeleted {
			return fmt.Errorf("object storage %d entered terminal state %q during restore", storageID, s3.ProvisioningStatus)
		}

		// check if the context is done before sleeping
		select {
		// handles both timeout and cancellation
		case <-pollingCtx.Done():
			return pollingCtx.Err()
		case <-time.After(pollingInterval):
		}
	}
}
