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
	"github.com/G-Core/gcore-go/packages/param"
)

// BucketLifecycleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketLifecycleService] method instead.
type BucketLifecycleService struct {
	Options []option.RequestOption
}

// NewBucketLifecycleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBucketLifecycleService(opts ...option.RequestOption) (r BucketLifecycleService) {
	r = BucketLifecycleService{}
	r.Options = opts
	return
}

// Configures automatic object expiration rules for an S3 bucket. Objects older
// than the specified number of days will be automatically deleted to manage
// storage costs and compliance.
func (r *BucketLifecycleService) New(ctx context.Context, bucketName string, params BucketLifecycleNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/s3/bucket/%s/lifecycle", params.StorageID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Removes all lifecycle rules from an S3 bucket, disabling automatic object
// expiration. Objects will no longer be automatically deleted based on age.
func (r *BucketLifecycleService) Delete(ctx context.Context, bucketName string, body BucketLifecycleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("storage/provisioning/v1/storage/%v/s3/bucket/%s/lifecycle", body.StorageID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type BucketLifecycleNewParams struct {
	StorageID      int64            `path:"storage_id,required" json:"-"`
	ExpirationDays param.Opt[int64] `json:"expiration_days,omitzero"`
	paramObj
}

func (r BucketLifecycleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BucketLifecycleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketLifecycleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketLifecycleDeleteParams struct {
	StorageID int64 `path:"storage_id,required" json:"-"`
	paramObj
}
