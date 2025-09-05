// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
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

// Updates storage configuration such as expiration date and server alias.
func (r *StorageService) Update(ctx context.Context, storageID int64, body StorageUpdateParams, opts ...option.RequestOption) (res *Storage, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Permanently deletes a storage and all its data. This action cannot be undone.
func (r *StorageService) Delete(ctx context.Context, storageID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieves detailed information about a specific storage including its
// configuration, credentials, and current status.
func (r *StorageService) Get(ctx context.Context, storageID int64, opts ...option.RequestOption) (res *Storage, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Associates an SSH public key with an SFTP storage, enabling passwordless
// authentication. Only works with SFTP storage types - not applicable to
// S3-compatible storage.
func (r *StorageService) LinkSSHKey(ctx context.Context, keyID int64, body StorageLinkSSHKeyParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/key/%v/link", body.StorageID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Restores a previously deleted S3 storage if it was deleted within the last 2
// weeks. SFTP storages cannot be restored.
func (r *StorageService) Restore(ctx context.Context, storageID int64, body StorageRestoreParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/restore", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Removes SSH key association from an SFTP storage, disabling passwordless
// authentication for that key. The key itself remains available for other
// storages.
func (r *StorageService) UnlinkSSHKey(ctx context.Context, keyID int64, body StorageUnlinkSSHKeyParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/key/%v/unlink", body.StorageID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type Storage struct {
	ID               int64              `json:"id"`
	Address          string             `json:"address"`
	CanRestore       bool               `json:"can_restore"`
	ClientID         int64              `json:"client_id"`
	CreatedAt        string             `json:"created_at"`
	Credentials      StorageCredentials `json:"credentials"`
	CustomConfigFile bool               `json:"custom_config_file"`
	DeletedAt        string             `json:"deleted_at"`
	DisableHTTP      bool               `json:"disable_http"`
	Expires          string             `json:"expires"`
	// Any of "s-ed1", "s-drc2", "s-sgc1", "s-nhn2", "s-darz", "s-ws1", "ams", "sin",
	// "fra", "mia".
	Location           StorageLocation   `json:"location"`
	Name               string            `json:"name"`
	ProvisioningStatus string            `json:"provisioning_status"`
	ResellerID         int64             `json:"reseller_id"`
	RewriteRules       map[string]string `json:"rewrite_rules"`
	ServerAlias        string            `json:"server_alias"`
	// Any of "sftp", "s3".
	Type StorageType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Address            respjson.Field
		CanRestore         respjson.Field
		ClientID           respjson.Field
		CreatedAt          respjson.Field
		Credentials        respjson.Field
		CustomConfigFile   respjson.Field
		DeletedAt          respjson.Field
		DisableHTTP        respjson.Field
		Expires            respjson.Field
		Location           respjson.Field
		Name               respjson.Field
		ProvisioningStatus respjson.Field
		ResellerID         respjson.Field
		RewriteRules       respjson.Field
		ServerAlias        respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Storage) RawJSON() string { return r.JSON.raw }
func (r *Storage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageCredentials struct {
	Keys         []StorageCredentialsKey `json:"keys"`
	S3           StorageCredentialsS3    `json:"s3"`
	SftpPassword string                  `json:"sftp_password"`
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
	ID        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	Name      string `json:"name"`
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
	AccessKey string `json:"access_key"`
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

type StorageLocation string

const (
	StorageLocationSEd1  StorageLocation = "s-ed1"
	StorageLocationSDrc2 StorageLocation = "s-drc2"
	StorageLocationSSgc1 StorageLocation = "s-sgc1"
	StorageLocationSNhn2 StorageLocation = "s-nhn2"
	StorageLocationSDarz StorageLocation = "s-darz"
	StorageLocationSWs1  StorageLocation = "s-ws1"
	StorageLocationAms   StorageLocation = "ams"
	StorageLocationSin   StorageLocation = "sin"
	StorageLocationFra   StorageLocation = "fra"
	StorageLocationMia   StorageLocation = "mia"
)

type StorageType string

const (
	StorageTypeSftp StorageType = "sftp"
	StorageTypeS3   StorageType = "s3"
)

type StorageUpdateParams struct {
	Expires     param.Opt[string] `json:"expires,omitzero"`
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
