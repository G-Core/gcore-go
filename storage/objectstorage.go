// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// ObjectStorageService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectStorageService] method instead.
type ObjectStorageService struct {
	Options    []option.RequestOption
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

type S3Storage struct {
	// Unique identifier for the storage instance
	ID int64 `json:"id" api:"required"`
	// Full hostname/address for accessing the storage endpoint
	Address string `json:"address" api:"required"`
	// ISO 8601 timestamp when the storage was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Geographic location code where the storage is provisioned
	LocationName string `json:"location_name" api:"required"`
	// User-defined name for the storage instance
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
	// Geographic location code where the storage is provisioned
	LocationName string `json:"location_name" api:"required"`
	// User-defined name for the storage instance
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
