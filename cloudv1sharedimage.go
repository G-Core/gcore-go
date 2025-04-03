// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1SharedImageService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1SharedImageService] method instead.
type CloudV1SharedImageService struct {
	Options []option.RequestOption
}

// NewCloudV1SharedImageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1SharedImageService(opts ...option.RequestOption) (r *CloudV1SharedImageService) {
	r = &CloudV1SharedImageService{}
	r.Options = opts
	return
}

// Create or update shared image
func (r *CloudV1SharedImageService) Update(ctx context.Context, body CloudV1SharedImageUpdateParams, opts ...option.RequestOption) (res *SharedImageList, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/shared_image"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete entity shared image
func (r *CloudV1SharedImageService) Delete(ctx context.Context, body CloudV1SharedImageDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cloud/v1/shared_image"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// If the entity isn't showed or has image_ids = None, all public images available
// to him. If the entity has image_ids = [], all public images are unavailable to
// the client.
func (r *CloudV1SharedImageService) GetAvailableIDs(ctx context.Context, query CloudV1SharedImageGetAvailableIDsParams, opts ...option.RequestOption) (res *SharedImageList, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/shared_image"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Cloud admin can change the set of images. Firstly, they may limit the number of
// public images available. Secondly, he can share the image to all clients of the
// reseller, with a specific client and with the project. If the entity isn't
// showed or has image_ids = None, all public imagesavailable to him. If the entity
// has image_ids = [], all public images are unavailable to the client.
type SharedImage struct {
	// Entity ID
	EntityID int64 `json:"entity_id,required"`
	// Entity type
	EntityType SharedImageEntityType `json:"entity_type,required"`
	// List of shared image IDs available for the projects, clients or resellers
	ImageIDs []string `json:"image_ids,required,nullable"`
	// Region ID
	RegionID int64           `json:"region_id,required"`
	ID       int64           `json:"id"`
	JSON     sharedImageJSON `json:"-"`
}

// sharedImageJSON contains the JSON metadata for the struct [SharedImage]
type sharedImageJSON struct {
	EntityID    apijson.Field
	EntityType  apijson.Field
	ImageIDs    apijson.Field
	RegionID    apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SharedImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sharedImageJSON) RawJSON() string {
	return r.raw
}

// Entity type
type SharedImageEntityType string

const (
	SharedImageEntityTypeClient   SharedImageEntityType = "client"
	SharedImageEntityTypeProject  SharedImageEntityType = "project"
	SharedImageEntityTypeReseller SharedImageEntityType = "reseller"
)

func (r SharedImageEntityType) IsKnown() bool {
	switch r {
	case SharedImageEntityTypeClient, SharedImageEntityTypeProject, SharedImageEntityTypeReseller:
		return true
	}
	return false
}

// Cloud admin can change the set of images. Firstly, they may limit the number of
// public images available. Secondly, he can share the image to all clients of the
// reseller, with a specific client and with the project. If the entity isn't
// showed or has image_ids = None, all public imagesavailable to him. If the entity
// has image_ids = [], all public images are unavailable to the client.
type SharedImageParam struct {
	// Entity ID
	EntityID param.Field[int64] `json:"entity_id,required"`
	// Entity type
	EntityType param.Field[SharedImageEntityType] `json:"entity_type,required"`
	// List of shared image IDs available for the projects, clients or resellers
	ImageIDs param.Field[[]string] `json:"image_ids,required"`
	// Region ID
	RegionID param.Field[int64] `json:"region_id,required"`
}

func (r SharedImageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SharedImageList struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []SharedImage       `json:"results"`
	JSON    sharedImageListJSON `json:"-"`
}

// sharedImageListJSON contains the JSON metadata for the struct [SharedImageList]
type sharedImageListJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SharedImageList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sharedImageListJSON) RawJSON() string {
	return r.raw
}

type CloudV1SharedImageUpdateParams struct {
	// Cloud admin can change the set of images. Firstly, they may limit the number of
	// public images available. Secondly, he can share the image to all clients of the
	// reseller, with a specific client and with the project. If the entity isn't
	// showed or has image_ids = None, all public imagesavailable to him. If the entity
	// has image_ids = [], all public images are unavailable to the client.
	SharedImage SharedImageParam `json:"shared_image,required"`
}

func (r CloudV1SharedImageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SharedImage)
}

type CloudV1SharedImageDeleteParams struct {
	// Entity type
	EntityType param.Field[CloudV1SharedImageDeleteParamsEntityType] `query:"entity_type"`
}

// URLQuery serializes [CloudV1SharedImageDeleteParams]'s query parameters as
// `url.Values`.
func (r CloudV1SharedImageDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Entity type
type CloudV1SharedImageDeleteParamsEntityType string

const (
	CloudV1SharedImageDeleteParamsEntityTypeReseller CloudV1SharedImageDeleteParamsEntityType = "reseller"
	CloudV1SharedImageDeleteParamsEntityTypeClient   CloudV1SharedImageDeleteParamsEntityType = "client"
	CloudV1SharedImageDeleteParamsEntityTypeProject  CloudV1SharedImageDeleteParamsEntityType = "project"
)

func (r CloudV1SharedImageDeleteParamsEntityType) IsKnown() bool {
	switch r {
	case CloudV1SharedImageDeleteParamsEntityTypeReseller, CloudV1SharedImageDeleteParamsEntityTypeClient, CloudV1SharedImageDeleteParamsEntityTypeProject:
		return true
	}
	return false
}

type CloudV1SharedImageGetAvailableIDsParams struct {
	// Entity type
	EntityType param.Field[CloudV1SharedImageGetAvailableIDsParamsEntityType] `query:"entity_type"`
}

// URLQuery serializes [CloudV1SharedImageGetAvailableIDsParams]'s query parameters
// as `url.Values`.
func (r CloudV1SharedImageGetAvailableIDsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Entity type
type CloudV1SharedImageGetAvailableIDsParamsEntityType string

const (
	CloudV1SharedImageGetAvailableIDsParamsEntityTypeReseller CloudV1SharedImageGetAvailableIDsParamsEntityType = "reseller"
	CloudV1SharedImageGetAvailableIDsParamsEntityTypeClient   CloudV1SharedImageGetAvailableIDsParamsEntityType = "client"
	CloudV1SharedImageGetAvailableIDsParamsEntityTypeProject  CloudV1SharedImageGetAvailableIDsParamsEntityType = "project"
)

func (r CloudV1SharedImageGetAvailableIDsParamsEntityType) IsKnown() bool {
	switch r {
	case CloudV1SharedImageGetAvailableIDsParamsEntityTypeReseller, CloudV1SharedImageGetAvailableIDsParamsEntityTypeClient, CloudV1SharedImageGetAvailableIDsParamsEntityTypeProject:
		return true
	}
	return false
}
