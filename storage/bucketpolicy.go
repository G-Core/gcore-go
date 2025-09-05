// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
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

// Creates or updates access policy for an S3 bucket, controlling permissions for
// bucket operations.
func (r *BucketPolicyService) New(ctx context.Context, bucketName string, body BucketPolicyNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/s3/bucket/%s/policy", body.StorageID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Removes the access policy from an S3 bucket, reverting to default permissions.
func (r *BucketPolicyService) Delete(ctx context.Context, bucketName string, body BucketPolicyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/s3/bucket/%s/policy", body.StorageID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieves the current access policy configuration for an S3 bucket.
func (r *BucketPolicyService) Get(ctx context.Context, bucketName string, query BucketPolicyGetParams, opts ...option.RequestOption) (res *StorageBucketPolicy, err error) {
	opts = append(r.Options[:], opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/s3/bucket/%s/policy", query.StorageID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// StorageGetBucketPolicyEndpointRes output
type StorageBucketPolicy struct {
	EnabledHTTPAccess bool `json:"enabledHttpAccess"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnabledHTTPAccess respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageBucketPolicy) RawJSON() string { return r.JSON.raw }
func (r *StorageBucketPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketPolicyNewParams struct {
	StorageID int64 `path:"storage_id,required" json:"-"`
	paramObj
}

type BucketPolicyDeleteParams struct {
	StorageID int64 `path:"storage_id,required" json:"-"`
	paramObj
}

type BucketPolicyGetParams struct {
	StorageID int64 `path:"storage_id,required" json:"-"`
	paramObj
}
