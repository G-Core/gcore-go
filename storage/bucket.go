// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// BucketService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketService] method instead.
type BucketService struct {
	Options   []option.RequestOption
	Cors      BucketCorService
	Lifecycle BucketLifecycleService
	Policy    BucketPolicyService
}

// NewBucketService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBucketService(opts ...option.RequestOption) (r BucketService) {
	r = BucketService{}
	r.Options = opts
	r.Cors = NewBucketCorService(opts...)
	r.Lifecycle = NewBucketLifecycleService(opts...)
	r.Policy = NewBucketPolicyService(opts...)
	return
}

// Creates a new bucket within an S3 storage. Only applicable to S3-compatible
// storages.
func (r *BucketService) New(ctx context.Context, bucketName string, body BucketNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/s3/bucket/%s", body.StorageID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Returns the list of buckets for the storage in a wrapped response. Response
// format: count: total number of buckets (independent of pagination) results:
// current page of buckets according to limit/offset
func (r *BucketService) List(ctx context.Context, storageID int64, query BucketListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Bucket], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("storage/provisioning/v2/storage/%v/s3/buckets", storageID)
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

// Returns the list of buckets for the storage in a wrapped response. Response
// format: count: total number of buckets (independent of pagination) results:
// current page of buckets according to limit/offset
func (r *BucketService) ListAutoPaging(ctx context.Context, storageID int64, query BucketListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Bucket] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, storageID, query, opts...))
}

// Removes a bucket from an S3 storage. All objects in the bucket will be
// automatically deleted before the bucket is removed.
func (r *BucketService) Delete(ctx context.Context, bucketName string, body BucketDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/s3/bucket/%s", body.StorageID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// BucketDtoV2 for response
type Bucket struct {
	// Name of the S3 bucket
	Name string `json:"name,required"`
	// Lifecycle policy expiration days (zero if not set)
	Lifecycle int64 `json:"lifecycle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Lifecycle   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Bucket) RawJSON() string { return r.JSON.raw }
func (r *Bucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketNewParams struct {
	StorageID int64 `path:"storage_id,required" json:"-"`
	paramObj
}

type BucketListParams struct {
	// Max number of records in response
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of records to skip before beginning to write in response.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BucketListParams]'s query parameters as `url.Values`.
func (r BucketListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type BucketDeleteParams struct {
	StorageID int64 `path:"storage_id,required" json:"-"`
	paramObj
}
