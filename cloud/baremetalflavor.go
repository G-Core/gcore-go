// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
)

// BaremetalFlavorService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBaremetalFlavorService] method instead.
type BaremetalFlavorService struct {
	Options []option.RequestOption
}

// NewBaremetalFlavorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBaremetalFlavorService(opts ...option.RequestOption) (r BaremetalFlavorService) {
	r = BaremetalFlavorService{}
	r.Options = opts
	return
}

// Retrieve a list of flavors. When the include_prices query parameter is
// specified, the list shows prices. A client in trial mode gets all price values
// as 0. If you get Pricing Error contact the support
func (r *BaremetalFlavorService) List(ctx context.Context, params BaremetalFlavorListParams, opts ...option.RequestOption) (res *BaremetalFlavorList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/bmflavors/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// List suitalbe flavors for bare metal server creation
func (r *BaremetalFlavorService) ListSuitable(ctx context.Context, params BaremetalFlavorListSuitableParams, opts ...option.RequestOption) (res *BaremetalFlavorList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/bminstances/%v/%v/available_flavors", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type BaremetalFlavorListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fbmflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/bmflavors/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbmflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/bmflavors/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbmflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/bmflavors/{project_id}/{region_id}'].get.parameters[2]"
	Disabled param.Opt[bool] `query:"disabled,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbmflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/bmflavors/{project_id}/{region_id}'].get.parameters[3]"
	ExcludeLinux param.Opt[bool] `query:"exclude_linux,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbmflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/bmflavors/{project_id}/{region_id}'].get.parameters[4]"
	ExcludeWindows param.Opt[bool] `query:"exclude_windows,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbmflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/bmflavors/{project_id}/{region_id}'].get.parameters[5]"
	IncludeCapacity param.Opt[bool] `query:"include_capacity,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbmflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/6'
	// "$.paths['/cloud/v1/bmflavors/{project_id}/{region_id}'].get.parameters[6]"
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbmflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/7'
	// "$.paths['/cloud/v1/bmflavors/{project_id}/{region_id}'].get.parameters[7]"
	IncludeReservationStock param.Opt[bool] `query:"include_reservation_stock,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalFlavorListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [BaremetalFlavorListParams]'s query parameters as
// `url.Values`.
func (r BaremetalFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type BaremetalFlavorListSuitableParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Favailable_flavors/post/parameters/0/schema'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}/available_flavors'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Favailable_flavors/post/parameters/1/schema'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}/available_flavors'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Favailable_flavors/post/parameters/2'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}/available_flavors'].post.parameters[2]"
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	// '#/components/schemas/AvailableBaremetalFlavorsRequestSchema/properties/apptemplate_id'
	// "$.components.schemas.AvailableBaremetalFlavorsRequestSchema.properties.apptemplate_id"
	ApptemplateID param.Opt[string] `json:"apptemplate_id,omitzero"`
	// '#/components/schemas/AvailableBaremetalFlavorsRequestSchema/properties/image_id'
	// "$.components.schemas.AvailableBaremetalFlavorsRequestSchema.properties.image_id"
	ImageID param.Opt[string] `json:"image_id,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalFlavorListSuitableParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r BaremetalFlavorListSuitableParams) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalFlavorListSuitableParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// URLQuery serializes [BaremetalFlavorListSuitableParams]'s query parameters as
// `url.Values`.
func (r BaremetalFlavorListSuitableParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
