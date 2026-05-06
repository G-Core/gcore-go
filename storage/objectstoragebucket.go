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

// ObjectStorageBucketService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectStorageBucketService] method instead.
type ObjectStorageBucketService struct {
	Options []option.RequestOption
}

// NewObjectStorageBucketService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectStorageBucketService(opts ...option.RequestOption) (r ObjectStorageBucketService) {
	r = ObjectStorageBucketService{}
	r.Options = opts
	return
}

// Creates a new bucket within an S3-compatible storage.
func (r *ObjectStorageBucketService) New(ctx context.Context, storageID int64, body ObjectStorageBucketNewParams, opts ...option.RequestOption) (res *Bucket, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("storage/v4/object_storages/%v/buckets", storageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of buckets for an S3-compatible storage, including each
// bucket's CORS, Lifecycle, and Policy configuration. Results are sorted
// alphabetically by bucket name (ascending).
func (r *ObjectStorageBucketService) List(ctx context.Context, storageID int64, query ObjectStorageBucketListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Bucket], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("storage/v4/object_storages/%v/buckets", storageID)
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

// Returns a paginated list of buckets for an S3-compatible storage, including each
// bucket's CORS, Lifecycle, and Policy configuration. Results are sorted
// alphabetically by bucket name (ascending).
func (r *ObjectStorageBucketService) ListAutoPaging(ctx context.Context, storageID int64, query ObjectStorageBucketListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Bucket] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, storageID, query, opts...))
}

type Bucket struct {
	Cors      BucketCors      `json:"cors" api:"required"`
	Lifecycle BucketLifecycle `json:"lifecycle" api:"required"`
	// Globally unique bucket name within the storage. Used as the path prefix when
	// accessing objects via S3 API.
	Name   string       `json:"name" api:"required"`
	Policy BucketPolicy `json:"policy" api:"required"`
	// Parent storage this bucket belongs to. Use this ID in the URL path for bucket
	// operations.
	StorageID int64 `json:"storage_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cors        respjson.Field
		Lifecycle   respjson.Field
		Name        respjson.Field
		Policy      respjson.Field
		StorageID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Bucket) RawJSON() string { return r.JSON.raw }
func (r *Bucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketCors struct {
	// Web domains allowed to make direct browser requests to this bucket (e.g.,
	// "https://myapp.com"). Use "\*" to allow any origin.
	AllowedOrigins []string `json:"allowed_origins" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedOrigins respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BucketCors) RawJSON() string { return r.JSON.raw }
func (r *BucketCors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketLifecycle struct {
	// Days after upload before objects are automatically deleted. For example, 30
	// means files are removed 30 days after creation.
	ExpirationDays int64 `json:"expiration_days" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpirationDays respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BucketLifecycle) RawJSON() string { return r.JSON.raw }
func (r *BucketLifecycle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketPolicy struct {
	// When true, anyone can download objects without credentials. When false, all
	// requests require valid S3 authentication.
	IsPublic bool `json:"is_public" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsPublic    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BucketPolicy) RawJSON() string { return r.JSON.raw }
func (r *BucketPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectStorageBucketNewParams struct {
	// Name of the bucket to create
	Name string `json:"name" api:"required"`
	paramObj
}

func (r ObjectStorageBucketNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectStorageBucketNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectStorageBucketNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectStorageBucketListParams struct {
	// Max number of records in response
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip before beginning to return results
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectStorageBucketListParams]'s query parameters as
// `url.Values`.
func (r ObjectStorageBucketListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
