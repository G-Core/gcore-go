// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
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

// Removes a bucket from an S3 storage. The bucket must be empty before deletion.
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

type BucketNewParams struct {
	StorageID int64 `path:"storage_id,required" json:"-"`
	paramObj
}

type BucketDeleteParams struct {
	StorageID int64 `path:"storage_id,required" json:"-"`
	paramObj
}
