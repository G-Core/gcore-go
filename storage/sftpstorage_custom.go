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

// NewAndPoll creates a new SFTP storage and polls until provisioning is complete,
// returning the original create response (with the one-time Password preserved
// when password_mode is "auto" or "set") and the provisioning status promoted to
// "active". Polling reuses the polling_interval_seconds and
// polling_timeout_seconds client options.
//
// JSON.raw on the returned value is the original POST body, so RawJSON() will
// still report provisioning_status="creating". Use the typed fields for
// post-provisioning state.
func (r *SftpStorageService) NewAndPoll(ctx context.Context, body SftpStorageNewParams, opts ...option.RequestOption) (res *SftpStorageCreated, err error) {
	var raw *http.Response
	actionOpts := slices.Concat(opts, []option.RequestOption{option.WithResponseInto(&raw)})

	created, err := r.New(ctx, body, actionOpts...)
	if err != nil {
		return nil, err
	}

	var rawBytes []byte
	if created == nil {
		// Caller's WithResponseBodyInto overrode the default deserialization
		// target. Recover the typed struct from the raw response so polling can
		// proceed; supports **http.Response and *[]byte body shapes.
		created = &SftpStorageCreated{}
		rawBytes, err = polling.RecoverActionBody(raw, created, opts...)
		if err != nil {
			return nil, fmt.Errorf("sftp NewAndPoll: %w", err)
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

	// Exclude WithResponseBodyInto and clear request body for intermediate Gets (SftpStorage must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)

	for {
		s, err := r.Get(pollingCtx, created.ID, pollOpts...)
		if err != nil {
			return nil, fmt.Errorf("failed to get sftp storage status: %w", err)
		}

		if s.ProvisioningStatus == SftpStorageProvisioningStatusActive {
			// Promote the polled status onto the create response (which still
			// carries the one-time Password). All other fields are populated
			// synchronously on the POST response and don't change during
			// provisioning.
			created.ProvisioningStatus = SftpStorageCreatedProvisioningStatus(s.ProvisioningStatus)
			if rawBytes != nil && raw != nil {
				enriched, sjErr := sjson.SetBytes(rawBytes, "provisioning_status", string(s.ProvisioningStatus))
				if sjErr != nil {
					return nil, fmt.Errorf("failed to enrich sftp create response with polled provisioning_status: %w", sjErr)
				}
				raw.Body = io.NopCloser(bytes.NewReader(enriched))
				polling.WriteResponseBodyInto(opts, enriched)
			}
			return created, nil
		}

		if s.ProvisioningStatus == SftpStorageProvisioningStatusDeleting ||
			s.ProvisioningStatus == SftpStorageProvisioningStatusDeleted {
			return nil, fmt.Errorf("sftp storage %d entered terminal state %q during creation", created.ID, s.ProvisioningStatus)
		}

		select {
		case <-pollingCtx.Done():
			return nil, pollingCtx.Err()
		case <-time.After(pollingInterval):
		}
	}
}

// UpdateAndPoll updates SFTP storage configuration and polls until provisioning
// is complete, returning the original update response (with a regenerated
// Password preserved when password_mode is "auto") and the provisioning status
// promoted to "active". Polling reuses the polling_interval_seconds and
// polling_timeout_seconds client options.
func (r *SftpStorageService) UpdateAndPoll(ctx context.Context, storageID int64, body SftpStorageUpdateParams, opts ...option.RequestOption) (res *SftpStorageCreated, err error) {
	var raw *http.Response
	actionOpts := slices.Concat(opts, []option.RequestOption{option.WithResponseInto(&raw)})

	updated, err := r.Update(ctx, storageID, body, actionOpts...)
	if err != nil {
		return nil, err
	}

	var rawBytes []byte
	if updated == nil {
		updated = &SftpStorageCreated{}
		rawBytes, err = polling.RecoverActionBody(raw, updated, opts...)
		if err != nil {
			return nil, fmt.Errorf("sftp UpdateAndPoll: %w", err)
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

	// Exclude WithResponseBodyInto and clear request body for intermediate Gets (SftpStorage must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)

	for {
		s, err := r.Get(pollingCtx, storageID, pollOpts...)
		if err != nil {
			return nil, fmt.Errorf("failed to get sftp storage status: %w", err)
		}

		if s.ProvisioningStatus == SftpStorageProvisioningStatusActive {
			// Promote the polled status onto the PATCH response (which still
			// carries any regenerated Password). All other fields are already
			// authoritative on the PATCH response.
			updated.ProvisioningStatus = SftpStorageCreatedProvisioningStatus(s.ProvisioningStatus)
			if rawBytes != nil && raw != nil {
				enriched, sjErr := sjson.SetBytes(rawBytes, "provisioning_status", string(s.ProvisioningStatus))
				if sjErr != nil {
					return nil, fmt.Errorf("failed to enrich sftp update response with polled provisioning_status: %w", sjErr)
				}
				raw.Body = io.NopCloser(bytes.NewReader(enriched))
				polling.WriteResponseBodyInto(opts, enriched)
			}
			return updated, nil
		}

		if s.ProvisioningStatus == SftpStorageProvisioningStatusDeleting ||
			s.ProvisioningStatus == SftpStorageProvisioningStatusDeleted {
			return nil, fmt.Errorf("sftp storage %d entered terminal state %q during update", storageID, s.ProvisioningStatus)
		}

		select {
		case <-pollingCtx.Done():
			return nil, pollingCtx.Err()
		case <-time.After(pollingInterval):
		}
	}
}

// DeleteAndPoll deletes an SFTP storage and polls until the resource is fully
// removed. Polling completes when Get returns 404 or the polled
// provisioning_status reaches "deleted". Polling reuses the
// polling_interval_seconds and polling_timeout_seconds client options.
func (r *SftpStorageService) DeleteAndPoll(ctx context.Context, storageID int64, opts ...option.RequestOption) error {
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

	// Exclude WithResponseBodyInto and clear request body for intermediate Gets (SftpStorage must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)

	for {
		s, err := r.Get(pollingCtx, storageID, pollOpts...)
		if err != nil {
			var apiErr *Error
			if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
				return nil
			}
			return fmt.Errorf("failed to get sftp storage status: %w", err)
		}

		if s.ProvisioningStatus == SftpStorageProvisioningStatusDeleted {
			return nil
		}

		select {
		case <-pollingCtx.Done():
			return pollingCtx.Err()
		case <-time.After(pollingInterval):
		}
	}
}
