// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1ResellerImageService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1ResellerImageService] method instead.
type CloudV1ResellerImageService struct {
	Options []option.RequestOption
}

// NewCloudV1ResellerImageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1ResellerImageService(opts ...option.RequestOption) (r *CloudV1ResellerImageService) {
	r = &CloudV1ResellerImageService{}
	r.Options = opts
	return
}

// If the reseller isn't showed or has image_ids = None, all public images
// available to him. If the reseller has image_ids = [], all public images are
// unavailable to the client.
func (r *CloudV1ResellerImageService) Get(ctx context.Context, resellerID int64, opts ...option.RequestOption) (res *CloudV1ResellerImageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/reseller_image/%v", resellerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Reseller and cloud admin can change the set of images, available to reseller
// clients. Firstly, they may limit the number of public images available.
// Secondly, they can share the image of the reseller client to all clients of the
// reseller. If the reseller isn't showed or has image_ids = None, all public
// images available to him. If the reseller has image_ids = [], all public images
// are unavailable to the client.
func (r *CloudV1ResellerImageService) Update(ctx context.Context, body CloudV1ResellerImageUpdateParams, opts ...option.RequestOption) (res *ResellerImage, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/reseller_image"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deprecated. Delete image limits for reseller clients
func (r *CloudV1ResellerImageService) Delete(ctx context.Context, resellerID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("cloud/v1/reseller_image/%v", resellerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Reseller and cloud admin can change the set of images available to reseller
// clients. Firstly, they may limit the number of public images available.
// Secondly, can share the image of the reseller client to all clients of the
// reseller. If image_ids = None for the reseller, all public images will be
// available to them. If image_ids = [] for the reseller, all public images will be
// unavailable to the client.
type ResellerImage struct {
	// List of image IDs available for the clients of this reseller.
	ImageIDs []string `json:"image_ids,required,nullable"`
	// Region ID.
	RegionID int64 `json:"region_id,required"`
	// Reseller ID.
	ResellerID int64             `json:"reseller_id,required"`
	JSON       resellerImageJSON `json:"-"`
}

// resellerImageJSON contains the JSON metadata for the struct [ResellerImage]
type resellerImageJSON struct {
	ImageIDs    apijson.Field
	RegionID    apijson.Field
	ResellerID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResellerImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resellerImageJSON) RawJSON() string {
	return r.raw
}

// Reseller and cloud admin can change the set of images available to reseller
// clients. Firstly, they may limit the number of public images available.
// Secondly, can share the image of the reseller client to all clients of the
// reseller. If image_ids = None for the reseller, all public images will be
// available to them. If image_ids = [] for the reseller, all public images will be
// unavailable to the client.
type ResellerImageParam struct {
	// List of image IDs available for the clients of this reseller.
	ImageIDs param.Field[[]string] `json:"image_ids,required"`
	// Region ID.
	RegionID param.Field[int64] `json:"region_id,required"`
	// Reseller ID.
	ResellerID param.Field[int64] `json:"reseller_id,required"`
}

func (r ResellerImageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1ResellerImageGetResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []ResellerImage                     `json:"results,required"`
	JSON    cloudV1ResellerImageGetResponseJSON `json:"-"`
}

// cloudV1ResellerImageGetResponseJSON contains the JSON metadata for the struct
// [CloudV1ResellerImageGetResponse]
type cloudV1ResellerImageGetResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1ResellerImageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1ResellerImageGetResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1ResellerImageUpdateParams struct {
	// Reseller and cloud admin can change the set of images available to reseller
	// clients. Firstly, they may limit the number of public images available.
	// Secondly, can share the image of the reseller client to all clients of the
	// reseller. If image_ids = None for the reseller, all public images will be
	// available to them. If image_ids = [] for the reseller, all public images will be
	// unavailable to the client.
	ResellerImage ResellerImageParam `json:"reseller_image,required"`
}

func (r CloudV1ResellerImageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ResellerImage)
}
