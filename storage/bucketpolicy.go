// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
)

// BucketPolicyService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketPolicyService] method instead.
type BucketPolicyService struct {
	Options []option.RequestOption
}

// NewBucketPolicyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBucketPolicyService(opts ...option.RequestOption) (r BucketPolicyService) {
	r = BucketPolicyService{}
	r.Options = opts
	return
}

// Applies a public read policy to the S3 bucket, allowing anonymous users to
// download/access all objects in the bucket via HTTP GET requests. This makes the
// bucket suitable for static website hosting, public file sharing, or CDN
// integration. Only grants read access - users cannot upload, modify, or delete
// objects without proper authentication.
//
// Deprecated: Use PATCH
// /provisioning/v3/storages/{`storage_id`}/buckets/{`bucket_name`} with {"policy":
// {"public": true}} instead.
//
// Deprecated: deprecated
func (r *BucketPolicyService) New(ctx context.Context, bucketName string, body BucketPolicyNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/s3/bucket/%s/policy", body.StorageID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Removes the public read policy from an S3 bucket, making all objects private and
// accessible only with proper authentication credentials. After this operation,
// anonymous users will no longer be able to access bucket contents via HTTP
// requests.
//
// Deprecated: Use PATCH
// /provisioning/v3/storages/{`storage_id`}/buckets/{`bucket_name`} with {"policy":
// {"public": false}} instead.
//
// Deprecated: deprecated
func (r *BucketPolicyService) Delete(ctx context.Context, bucketName string, body BucketPolicyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/s3/bucket/%s/policy", body.StorageID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns whether the S3 bucket is currently configured for public read access.
// Shows if anonymous users can download objects from the bucket via HTTP requests.
//
// Deprecated: Use GET
// /provisioning/v3/storages/{`storage_id`}/buckets/{`bucket_name`} instead, which
// returns policy status along with CORS and lifecycle configuration.
//
// Deprecated: deprecated
func (r *BucketPolicyService) Get(ctx context.Context, bucketName string, query BucketPolicyGetParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/s3/bucket/%s/policy", query.StorageID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BucketPolicyNewParams struct {
	StorageID int64 `path:"storage_id" api:"required" json:"-"`
	paramObj
}

type BucketPolicyDeleteParams struct {
	StorageID int64 `path:"storage_id" api:"required" json:"-"`
	paramObj
}

type BucketPolicyGetParams struct {
	StorageID int64 `path:"storage_id" api:"required" json:"-"`
	paramObj
}
