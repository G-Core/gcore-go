// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// BucketCorService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketCorService] method instead.
type BucketCorService struct {
	Options []option.RequestOption
}

// NewBucketCorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBucketCorService(opts ...option.RequestOption) (r BucketCorService) {
	r = BucketCorService{}
	r.Options = opts
	return
}

// Configures Cross-Origin Resource Sharing (CORS) rules for an S3 bucket, allowing
// web applications from specified domains to access bucket resources directly from
// browsers.
func (r *BucketCorService) New(ctx context.Context, bucketName string, params BucketCorNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/s3/bucket/%s/cors", params.StorageID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Retrieves the current Cross-Origin Resource Sharing (CORS) configuration for an
// S3 bucket, showing which domains are allowed to access the bucket from web
// browsers.
func (r *BucketCorService) Get(ctx context.Context, bucketName string, query BucketCorGetParams, opts ...option.RequestOption) (res *BucketCors, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/s3/bucket/%s/cors", query.StorageID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// StorageGetBucketCorsEndpointRes output
type BucketCors struct {
	// List of allowed origins for Cross-Origin Resource Sharing (CORS) requests.
	// Contains domains/URLs that are permitted to make cross-origin requests to this
	// bucket.
	AllowedOrigins []string `json:"allowedOrigins"`
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

type BucketCorNewParams struct {
	StorageID int64 `path:"storage_id,required" json:"-"`
	// List of allowed origins for CORS requests
	AllowedOrigins []string `json:"allowedOrigins,omitzero"`
	paramObj
}

func (r BucketCorNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BucketCorNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketCorNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketCorGetParams struct {
	StorageID int64 `path:"storage_id,required" json:"-"`
	paramObj
}
