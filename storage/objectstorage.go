// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/polling"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/tidwall/sjson"
)

// ObjectStorageService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectStorageService] method instead.
type ObjectStorageService struct {
	Options []option.RequestOption
	// Object storage access keys provide secure credentials for API access to object
	// storage resources.
	AccessKeys ObjectStorageAccessKeyService
	Buckets    ObjectStorageBucketService
}

// NewObjectStorageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectStorageService(opts ...option.RequestOption) (r ObjectStorageService) {
	r = ObjectStorageService{}
	r.Options = opts
	r.AccessKeys = NewObjectStorageAccessKeyService(opts...)
	r.Buckets = NewObjectStorageBucketService(opts...)
	return
}

// Creates a new S3-compatible storage instance in the specified location and
// returns the storage details including credentials.
func (r *ObjectStorageService) New(ctx context.Context, body ObjectStorageNewParams, opts ...option.RequestOption) (res *S3StorageCreated, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/v4/object_storages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

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
		// the raw response so polling can proceed; only viable when the body
		// shape preserves it (**http.Response). The one-time AccessKeys are
		// only present here — they are never replayed on subsequent Gets.
		created = &S3StorageCreated{}
		rawBytes, err = polling.RecoverActionBody(raw, created)
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

// Returns a paginated list of S3-compatible storage instances for the
// authenticated client.
func (r *ObjectStorageService) List(ctx context.Context, query ObjectStorageListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[S3Storage], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "storage/v4/object_storages"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a paginated list of S3-compatible storage instances for the
// authenticated client.
func (r *ObjectStorageService) ListAutoPaging(ctx context.Context, query ObjectStorageListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[S3Storage] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes an S3-compatible storage instance. This operation cannot be undone.
func (r *ObjectStorageService) Delete(ctx context.Context, storageID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("storage/v4/object_storages/%v", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
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

// Returns details of a specific S3-compatible storage instance.
func (r *ObjectStorageService) Get(ctx context.Context, storageID int64, opts ...option.RequestOption) (res *S3Storage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("storage/v4/object_storages/%v", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Restores a previously deleted S3-compatible storage instance if it was deleted
// within the last 2 weeks.
func (r *ObjectStorageService) Restore(ctx context.Context, storageID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("storage/v4/object_storages/%v/restore", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// RestoreAndPoll restores a previously deleted S3-compatible storage (within the
// 2-week window) and polls until provisioning_status reaches "active". Polling
// reuses the polling_interval_seconds and polling_timeout_seconds client options.
//
// Returns the final S3Storage on success, or an error if Restore itself fails,
// the storage transitions to deleting/deleted during restore, or polling times
// out / is cancelled.
func (r *ObjectStorageService) RestoreAndPoll(ctx context.Context, storageID int64, opts ...option.RequestOption) (res *S3Storage, err error) {
	if err := r.Restore(ctx, storageID, opts...); err != nil {
		return nil, err
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
		s3, err := r.Get(pollingCtx, storageID, pollOpts...)
		if err != nil {
			return nil, fmt.Errorf("failed to get object storage status: %w", err)
		}

		if s3.ProvisioningStatus == S3StorageProvisioningStatusActive {
			// Restore returns no action body, so the caller's WithResponseBodyInto
			// (terraform provider pattern) has nothing to capture from the POST.
			// Replay the final Get with the caller's full opts so their captured
			// response is populated with the active S3Storage payload.
			return r.Get(pollingCtx, storageID, opts...)
		}

		if s3.ProvisioningStatus == S3StorageProvisioningStatusDeleting ||
			s3.ProvisioningStatus == S3StorageProvisioningStatusDeleted {
			return nil, fmt.Errorf("object storage %d entered terminal state %q during restore", storageID, s3.ProvisioningStatus)
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

type S3Storage struct {
	// Unique identifier for the storage instance
	ID int64 `json:"id" api:"required"`
	// Full hostname/address for accessing the storage endpoint
	Address string `json:"address" api:"required"`
	// ISO 8601 timestamp when the storage was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Read-only internal full name of the storage, composed as "{`client_id`}-{name}".
	// Used internally by the backend. Clients should continue to identify the storage
	// by `name`.
	FullName string `json:"full_name" api:"required"`
	// Geographic location code where the storage is provisioned
	LocationName string `json:"location_name" api:"required"`
	// User-defined name for the storage instance, as supplied at creation time.
	Name string `json:"name" api:"required"`
	// Lifecycle status of the storage. Use this to check readiness before operations.
	//
	// Any of "creating", "active", "updating", "deleting", "deleted".
	ProvisioningStatus S3StorageProvisioningStatus `json:"provisioning_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Address            respjson.Field
		CreatedAt          respjson.Field
		FullName           respjson.Field
		LocationName       respjson.Field
		Name               respjson.Field
		ProvisioningStatus respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r S3Storage) RawJSON() string { return r.JSON.raw }
func (r *S3Storage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle status of the storage. Use this to check readiness before operations.
type S3StorageProvisioningStatus string

const (
	S3StorageProvisioningStatusCreating S3StorageProvisioningStatus = "creating"
	S3StorageProvisioningStatusActive   S3StorageProvisioningStatus = "active"
	S3StorageProvisioningStatusUpdating S3StorageProvisioningStatus = "updating"
	S3StorageProvisioningStatusDeleting S3StorageProvisioningStatus = "deleting"
	S3StorageProvisioningStatusDeleted  S3StorageProvisioningStatus = "deleted"
)

type S3StorageCreated struct {
	// Unique identifier for the storage instance
	ID int64 `json:"id" api:"required"`
	// S3 access keys
	AccessKeys []S3StorageCreatedAccessKey `json:"access_keys" api:"required"`
	// Full hostname/address for accessing the storage endpoint
	Address string `json:"address" api:"required"`
	// ISO 8601 timestamp when the storage was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Read-only internal full name of the storage, composed as "{`client_id`}-{name}".
	// Used internally by the backend. Clients should continue to identify the storage
	// by `name`.
	FullName string `json:"full_name" api:"required"`
	// Geographic location code where the storage is provisioned
	LocationName string `json:"location_name" api:"required"`
	// User-defined name for the storage instance, as supplied at creation time.
	Name string `json:"name" api:"required"`
	// Lifecycle status of the storage. Use this to check readiness before operations.
	//
	// Any of "creating", "active", "updating", "deleting", "deleted".
	ProvisioningStatus S3StorageCreatedProvisioningStatus `json:"provisioning_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AccessKeys         respjson.Field
		Address            respjson.Field
		CreatedAt          respjson.Field
		FullName           respjson.Field
		LocationName       respjson.Field
		Name               respjson.Field
		ProvisioningStatus respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r S3StorageCreated) RawJSON() string { return r.JSON.raw }
func (r *S3StorageCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type S3StorageCreatedAccessKey struct {
	// Access key ID used as the username in S3 authentication. Pass this in the
	// `AWS_ACCESS_KEY_ID` field of your S3 client.
	AccessKey string `json:"access_key" api:"required"`
	// Secret key used as the password in S3 authentication. Save this now — it cannot
	// be retrieved again.
	SecretKey string `json:"secret_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessKey   respjson.Field
		SecretKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r S3StorageCreatedAccessKey) RawJSON() string { return r.JSON.raw }
func (r *S3StorageCreatedAccessKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle status of the storage. Use this to check readiness before operations.
type S3StorageCreatedProvisioningStatus string

const (
	S3StorageCreatedProvisioningStatusCreating S3StorageCreatedProvisioningStatus = "creating"
	S3StorageCreatedProvisioningStatusActive   S3StorageCreatedProvisioningStatus = "active"
	S3StorageCreatedProvisioningStatusUpdating S3StorageCreatedProvisioningStatus = "updating"
	S3StorageCreatedProvisioningStatusDeleting S3StorageCreatedProvisioningStatus = "deleting"
	S3StorageCreatedProvisioningStatusDeleted  S3StorageCreatedProvisioningStatus = "deleted"
)

type ObjectStorageNewParams struct {
	// Location code where the storage should be created
	LocationName string `json:"location_name" api:"required"`
	// User-defined name for the storage instance
	Name string `json:"name" api:"required"`
	paramObj
}

func (r ObjectStorageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectStorageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectStorageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectStorageListParams struct {
	// Filter by storage ID
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Max number of records in response
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by storage location/region
	LocationName param.Opt[string] `query:"location_name,omitzero" json:"-"`
	// Filter by storage name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Number of records to skip
	Offset  param.Opt[int64]  `query:"offset,omitzero" json:"-"`
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// Include deleted storages
	ShowDeleted param.Opt[bool] `query:"show_deleted,omitzero" json:"-"`
	// Filter by provisioning status
	//
	// Any of "active", "creating", "updating", "deleting", "deleted".
	ProvisioningStatus ObjectStorageListParamsProvisioningStatus `query:"provisioning_status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectStorageListParams]'s query parameters as
// `url.Values`.
func (r ObjectStorageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by provisioning status
type ObjectStorageListParamsProvisioningStatus string

const (
	ObjectStorageListParamsProvisioningStatusActive   ObjectStorageListParamsProvisioningStatus = "active"
	ObjectStorageListParamsProvisioningStatusCreating ObjectStorageListParamsProvisioningStatus = "creating"
	ObjectStorageListParamsProvisioningStatusUpdating ObjectStorageListParamsProvisioningStatus = "updating"
	ObjectStorageListParamsProvisioningStatusDeleting ObjectStorageListParamsProvisioningStatus = "deleting"
	ObjectStorageListParamsProvisioningStatusDeleted  ObjectStorageListParamsProvisioningStatus = "deleted"
)
