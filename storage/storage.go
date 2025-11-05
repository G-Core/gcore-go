// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// StorageService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageService] method instead.
type StorageService struct {
	Options     []option.RequestOption
	Locations   LocationService
	Statistics  StatisticService
	Credentials CredentialService
	Buckets     BucketService
}

// NewStorageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStorageService(opts ...option.RequestOption) (r StorageService) {
	r = StorageService{}
	r.Options = opts
	r.Locations = NewLocationService(opts...)
	r.Statistics = NewStatisticService(opts...)
	r.Credentials = NewCredentialService(opts...)
	r.Buckets = NewBucketService(opts...)
	return
}

// Creates a new storage instance (S3 or SFTP) in the specified location and
// returns the storage details including credentials.
func (r *StorageService) New(ctx context.Context, body StorageNewParams, opts ...option.RequestOption) (res *Storage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/provisioning/v2/storage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates storage configuration such as expiration date and server alias. Used for
// SFTP storages.
func (r *StorageService) Update(ctx context.Context, storageID int64, body StorageUpdateParams, opts ...option.RequestOption) (res *Storage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("storage/provisioning/v2/storage/%v", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns storages with the same filtering and pagination as v2, but in a
// simplified response shape for easier client consumption.
//
// Response format: count: total number of storages matching the filter
// (independent of pagination) results: the current page of storages according to
// limit/offset
func (r *StorageService) List(ctx context.Context, query StorageListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Storage], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "storage/provisioning/v3/storage"
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

// Returns storages with the same filtering and pagination as v2, but in a
// simplified response shape for easier client consumption.
//
// Response format: count: total number of storages matching the filter
// (independent of pagination) results: the current page of storages according to
// limit/offset
func (r *StorageService) ListAutoPaging(ctx context.Context, query StorageListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Storage] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a storage and all its data. This action cannot be undone.
func (r *StorageService) Delete(ctx context.Context, storageID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieves detailed information about a specific storage including its
// configuration, credentials, and current status.
func (r *StorageService) Get(ctx context.Context, storageID int64, opts ...option.RequestOption) (res *Storage, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Associates an SSH public key with an SFTP storage, enabling passwordless
// authentication. Only works with SFTP storage types - not applicable to
// S3-compatible storage.
func (r *StorageService) LinkSSHKey(ctx context.Context, keyID int64, body StorageLinkSSHKeyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/key/%v/link", body.StorageID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Restores a previously deleted S3 storage if it was deleted within the last 2
// weeks. SFTP storages cannot be restored.
func (r *StorageService) Restore(ctx context.Context, storageID int64, body StorageRestoreParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/restore", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Removes SSH key association from an SFTP storage, disabling passwordless
// authentication for that key. The key itself remains available for other
// storages.
func (r *StorageService) UnlinkSSHKey(ctx context.Context, keyID int64, body StorageUnlinkSSHKeyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/key/%v/unlink", body.StorageID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type Storage struct {
	// Unique identifier for the storage instance
	ID int64 `json:"id,required"`
	// Full hostname/address for accessing the storage endpoint
	Address string `json:"address,required"`
	// Client identifier who owns this storage
	ClientID int64 `json:"client_id,required"`
	// ISO 8601 timestamp when the storage was created
	CreatedAt string `json:"created_at,required"`
	// Geographic location code where the storage is provisioned
	Location string `json:"location,required"`
	// User-defined name for the storage instance
	Name string `json:"name,required"`
	// Current provisioning status of the storage instance
	//
	// Any of "creating", "ok", "updating", "deleting", "deleted".
	ProvisioningStatus StorageProvisioningStatus `json:"provisioning_status,required"`
	// Reseller technical client ID associated with the client
	ResellerID int64 `json:"reseller_id,required"`
	// Storage protocol type - either S3-compatible object storage or SFTP file
	// transfer
	//
	// Any of "sftp", "s3".
	Type StorageType `json:"type,required"`
	// Whether this storage can be restored if deleted (S3 storages only, within 2
	// weeks)
	CanRestore  bool               `json:"can_restore"`
	Credentials StorageCredentials `json:"credentials"`
	// Whether custom configuration file is used for this storage
	CustomConfigFile bool `json:"custom_config_file"`
	// ISO 8601 timestamp when the storage was deleted (only present for deleted
	// storages)
	DeletedAt string `json:"deleted_at"`
	// Whether HTTP access is disabled for this storage (HTTPS only)
	DisableHTTP bool `json:"disable_http"`
	// ISO 8601 timestamp when the storage will expire (if set)
	Expires string `json:"expires"`
	// Custom URL rewrite rules for the storage (admin-configurable)
	RewriteRules map[string]string `json:"rewrite_rules"`
	// Custom domain alias for accessing the storage
	ServerAlias string `json:"server_alias"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Address            respjson.Field
		ClientID           respjson.Field
		CreatedAt          respjson.Field
		Location           respjson.Field
		Name               respjson.Field
		ProvisioningStatus respjson.Field
		ResellerID         respjson.Field
		Type               respjson.Field
		CanRestore         respjson.Field
		Credentials        respjson.Field
		CustomConfigFile   respjson.Field
		DeletedAt          respjson.Field
		DisableHTTP        respjson.Field
		Expires            respjson.Field
		RewriteRules       respjson.Field
		ServerAlias        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Storage) RawJSON() string { return r.JSON.raw }
func (r *Storage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current provisioning status of the storage instance
type StorageProvisioningStatus string

const (
	StorageProvisioningStatusCreating StorageProvisioningStatus = "creating"
	StorageProvisioningStatusOk       StorageProvisioningStatus = "ok"
	StorageProvisioningStatusUpdating StorageProvisioningStatus = "updating"
	StorageProvisioningStatusDeleting StorageProvisioningStatus = "deleting"
	StorageProvisioningStatusDeleted  StorageProvisioningStatus = "deleted"
)

// Storage protocol type - either S3-compatible object storage or SFTP file
// transfer
type StorageType string

const (
	StorageTypeSftp StorageType = "sftp"
	StorageTypeS3   StorageType = "s3"
)

type StorageCredentials struct {
	// SSH public keys associated with SFTP storage for passwordless authentication
	Keys []StorageCredentialsKey `json:"keys"`
	S3   StorageCredentialsS3    `json:"s3"`
	// Generated or user-provided password for SFTP access (only present for SFTP
	// storage type)
	SftpPassword string `json:"sftp_password"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Keys         respjson.Field
		S3           respjson.Field
		SftpPassword respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageCredentials) RawJSON() string { return r.JSON.raw }
func (r *StorageCredentials) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageCredentialsKey struct {
	// Unique identifier for the SSH key
	ID int64 `json:"id"`
	// ISO 8601 timestamp when the SSH key was created
	CreatedAt string `json:"created_at"`
	// User-defined name for the SSH key
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageCredentialsKey) RawJSON() string { return r.JSON.raw }
func (r *StorageCredentialsKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageCredentialsS3 struct {
	// S3-compatible access key identifier for authentication
	AccessKey string `json:"access_key"`
	// S3-compatible secret key for authentication (keep secure)
	SecretKey string `json:"secret_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessKey   respjson.Field
		SecretKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageCredentialsS3) RawJSON() string { return r.JSON.raw }
func (r *StorageCredentialsS3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageNewParams struct {
	// Geographic location where the storage will be provisioned. Each location
	// represents a specific data center region.
	Location string `json:"location,required"`
	// Unique storage name identifier. Must contain only letters, numbers, dashes, and
	// underscores. Cannot be empty and must be less than 256 characters.
	Name string `json:"name,required"`
	// Storage protocol type. Choose 's3' for S3-compatible object storage with API
	// access, or `sftp` for SFTP file transfer protocol.
	//
	// Any of "sftp", "s3".
	Type StorageNewParamsType `json:"type,omitzero,required"`
	// Automatically generate a secure password for SFTP storage access. Only
	// applicable when type is `sftp`. When `true`, a random password will be generated
	// and returned in the response.
	GenerateSftpPassword param.Opt[bool] `json:"generate_sftp_password,omitzero"`
	// Custom password for SFTP storage access. Only applicable when type is `sftp`. If
	// not provided and `generate_sftp_password` is `false`, no password authentication
	// will be available.
	SftpPassword param.Opt[string] `json:"sftp_password,omitzero"`
	paramObj
}

func (r StorageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StorageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Storage protocol type. Choose 's3' for S3-compatible object storage with API
// access, or `sftp` for SFTP file transfer protocol.
type StorageNewParamsType string

const (
	StorageNewParamsTypeSftp StorageNewParamsType = "sftp"
	StorageNewParamsTypeS3   StorageNewParamsType = "s3"
)

type StorageUpdateParams struct {
	// Duration when the storage should expire in format like "1 years 6 months 2 weeks
	// 3 days 5 hours 10 minutes 15 seconds". Set empty to remove expiration.
	Expires param.Opt[string] `json:"expires,omitzero"`
	// Custom domain alias for accessing the storage. Set empty to remove alias.
	ServerAlias param.Opt[string] `json:"server_alias,omitzero"`
	paramObj
}

func (r StorageUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow StorageUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageListParams struct {
	// Filter by storage ID
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Max number of records in response
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by storage location/region
	Location param.Opt[string] `query:"location,omitzero" json:"-"`
	// Filter by storage name (exact match)
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Number of records to skip before beginning to write in response.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Field name to sort by
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// Include deleted storages in the response
	ShowDeleted param.Opt[bool] `query:"show_deleted,omitzero" json:"-"`
	// Ascending or descending order
	//
	// Any of "asc", "desc".
	OrderDirection StorageListParamsOrderDirection `query:"order_direction,omitzero" json:"-"`
	// Filter by storage status
	//
	// Any of "active", "suspended", "deleted", "pending".
	Status StorageListParamsStatus `query:"status,omitzero" json:"-"`
	// Filter by storage type
	//
	// Any of "s3", "sftp".
	Type StorageListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StorageListParams]'s query parameters as `url.Values`.
func (r StorageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Ascending or descending order
type StorageListParamsOrderDirection string

const (
	StorageListParamsOrderDirectionAsc  StorageListParamsOrderDirection = "asc"
	StorageListParamsOrderDirectionDesc StorageListParamsOrderDirection = "desc"
)

// Filter by storage status
type StorageListParamsStatus string

const (
	StorageListParamsStatusActive    StorageListParamsStatus = "active"
	StorageListParamsStatusSuspended StorageListParamsStatus = "suspended"
	StorageListParamsStatusDeleted   StorageListParamsStatus = "deleted"
	StorageListParamsStatusPending   StorageListParamsStatus = "pending"
)

// Filter by storage type
type StorageListParamsType string

const (
	StorageListParamsTypeS3   StorageListParamsType = "s3"
	StorageListParamsTypeSftp StorageListParamsType = "sftp"
)

type StorageLinkSSHKeyParams struct {
	StorageID int64 `path:"storage_id,required" json:"-"`
	paramObj
}

type StorageRestoreParams struct {
	ClientID param.Opt[int64] `query:"client_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StorageRestoreParams]'s query parameters as `url.Values`.
func (r StorageRestoreParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StorageUnlinkSSHKeyParams struct {
	StorageID int64 `path:"storage_id,required" json:"-"`
	paramObj
}
