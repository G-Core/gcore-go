// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"errors"
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

// ObjectStorageAccessKeyService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectStorageAccessKeyService] method instead.
type ObjectStorageAccessKeyService struct {
	Options []option.RequestOption
}

// NewObjectStorageAccessKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectStorageAccessKeyService(opts ...option.RequestOption) (r ObjectStorageAccessKeyService) {
	r = ObjectStorageAccessKeyService{}
	r.Options = opts
	return
}

// Creates a new access key for an S3-compatible storage. Returns the new access
// key and secret key. Maximum 2 access keys per storage.
func (r *ObjectStorageAccessKeyService) New(ctx context.Context, storageID int64, opts ...option.RequestOption) (res *AccessKeyCreated, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("storage/v4/object_storages/%v/access_keys", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns a list of access keys for an S3-compatible storage.
func (r *ObjectStorageAccessKeyService) List(ctx context.Context, storageID int64, query ObjectStorageAccessKeyListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[AccessKey], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("storage/v4/object_storages/%v/access_keys", storageID)
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

// Returns a list of access keys for an S3-compatible storage.
func (r *ObjectStorageAccessKeyService) ListAutoPaging(ctx context.Context, storageID int64, query ObjectStorageAccessKeyListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[AccessKey] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, storageID, query, opts...))
}

// Deletes an access key from an S3-compatible storage.
func (r *ObjectStorageAccessKeyService) Delete(ctx context.Context, accessKey string, body ObjectStorageAccessKeyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if accessKey == "" {
		err = errors.New("missing required access_key parameter")
		return err
	}
	path := fmt.Sprintf("storage/v4/object_storages/%v/access_keys/%s", body.StorageID, accessKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Returns details of a specific access key for an S3-compatible storage.
func (r *ObjectStorageAccessKeyService) Get(ctx context.Context, accessKey string, query ObjectStorageAccessKeyGetParams, opts ...option.RequestOption) (res *AccessKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if accessKey == "" {
		err = errors.New("missing required access_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("storage/v4/object_storages/%v/access_keys/%s", query.StorageID, accessKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AccessKey struct {
	// Access key ID used as the username in S3 authentication. Pass this in the
	// `AWS_ACCESS_KEY_ID` field of your S3 client configuration.
	AccessKey string `json:"access_key" api:"required"`
	// ISO 8601 timestamp when the access key was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessKey   respjson.Field
		CreatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessKey) RawJSON() string { return r.JSON.raw }
func (r *AccessKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessKeyCreated struct {
	// Access key ID used as the username in S3 authentication. Pass this in the
	// `AWS_ACCESS_KEY_ID` field of your S3 client.
	AccessKey string `json:"access_key" api:"required"`
	// ISO 8601 timestamp when the access key was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Secret key used as the password in S3 authentication. Save this now — it cannot
	// be retrieved again.
	SecretKey string `json:"secret_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessKey   respjson.Field
		CreatedAt   respjson.Field
		SecretKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessKeyCreated) RawJSON() string { return r.JSON.raw }
func (r *AccessKeyCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectStorageAccessKeyListParams struct {
	// Max number of records in response
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip before beginning to return results
	Offset  param.Opt[int64]  `query:"offset,omitzero" json:"-"`
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectStorageAccessKeyListParams]'s query parameters as
// `url.Values`.
func (r ObjectStorageAccessKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ObjectStorageAccessKeyDeleteParams struct {
	StorageID int64 `path:"storage_id" api:"required" json:"-"`
	paramObj
}

type ObjectStorageAccessKeyGetParams struct {
	StorageID int64 `path:"storage_id" api:"required" json:"-"`
	paramObj
}
