// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// InstanceFlavorService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceFlavorService] method instead.
type InstanceFlavorService struct {
	Options []option.RequestOption
}

// NewInstanceFlavorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceFlavorService(opts ...option.RequestOption) (r InstanceFlavorService) {
	r = InstanceFlavorService{}
	r.Options = opts
	return
}

// Retrieve a list of flavors. When the include_prices query parameter is
// specified, the list shows prices. A client in trial mode gets all price values
// as 0. If you get Pricing Error contact the support
func (r *InstanceFlavorService) List(ctx context.Context, params InstanceFlavorListParams, opts ...option.RequestOption) (res *InstanceFlavorList, err error) {
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
	path := fmt.Sprintf("cloud/v1/flavors/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// List suitable flavors for instance resize
func (r *InstanceFlavorService) ListForResize(ctx context.Context, instanceID string, params InstanceFlavorListForResizeParams, opts ...option.RequestOption) (res *InstanceFlavorList, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/available_flavors", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// List suitable flavors for instance creation
func (r *InstanceFlavorService) ListSuitable(ctx context.Context, params InstanceFlavorListSuitableParams, opts ...option.RequestOption) (res *InstanceFlavorList, err error) {
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
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/available_flavors", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// '#/components/schemas/InstanceFlavorExtendedSerializer'
// "$.components.schemas.InstanceFlavorExtendedSerializer"
type InstanceFlavor struct {
	// '#/components/schemas/InstanceFlavorExtendedSerializer/properties/architecture'
	// "$.components.schemas.InstanceFlavorExtendedSerializer.properties.architecture"
	Architecture string `json:"architecture,required"`
	// '#/components/schemas/InstanceFlavorExtendedSerializer/properties/disabled'
	// "$.components.schemas.InstanceFlavorExtendedSerializer.properties.disabled"
	Disabled bool `json:"disabled,required"`
	// '#/components/schemas/InstanceFlavorExtendedSerializer/properties/flavor_id'
	// "$.components.schemas.InstanceFlavorExtendedSerializer.properties.flavor_id"
	FlavorID string `json:"flavor_id,required"`
	// '#/components/schemas/InstanceFlavorExtendedSerializer/properties/flavor_name'
	// "$.components.schemas.InstanceFlavorExtendedSerializer.properties.flavor_name"
	FlavorName string `json:"flavor_name,required"`
	// '#/components/schemas/InstanceFlavorExtendedSerializer/properties/os_type'
	// "$.components.schemas.InstanceFlavorExtendedSerializer.properties.os_type"
	OsType string `json:"os_type,required"`
	// '#/components/schemas/InstanceFlavorExtendedSerializer/properties/ram'
	// "$.components.schemas.InstanceFlavorExtendedSerializer.properties.ram"
	Ram int64 `json:"ram,required"`
	// '#/components/schemas/InstanceFlavorExtendedSerializer/properties/vcpus'
	// "$.components.schemas.InstanceFlavorExtendedSerializer.properties.vcpus"
	Vcpus int64 `json:"vcpus,required"`
	// '#/components/schemas/InstanceFlavorExtendedSerializer/properties/capacity/anyOf/0'
	// "$.components.schemas.InstanceFlavorExtendedSerializer.properties.capacity.anyOf[0]"
	Capacity int64 `json:"capacity,nullable"`
	// '#/components/schemas/InstanceFlavorExtendedSerializer/properties/currency_code/anyOf/0'
	// "$.components.schemas.InstanceFlavorExtendedSerializer.properties.currency_code.anyOf[0]"
	CurrencyCode string `json:"currency_code,nullable"`
	// '#/components/schemas/InstanceFlavorExtendedSerializer/properties/hardware_description'
	// "$.components.schemas.InstanceFlavorExtendedSerializer.properties.hardware_description"
	HardwareDescription map[string]string `json:"hardware_description"`
	// '#/components/schemas/InstanceFlavorExtendedSerializer/properties/price_per_hour/anyOf/0'
	// "$.components.schemas.InstanceFlavorExtendedSerializer.properties.price_per_hour.anyOf[0]"
	PricePerHour float64 `json:"price_per_hour,nullable"`
	// '#/components/schemas/InstanceFlavorExtendedSerializer/properties/price_per_month/anyOf/0'
	// "$.components.schemas.InstanceFlavorExtendedSerializer.properties.price_per_month.anyOf[0]"
	PricePerMonth float64 `json:"price_per_month,nullable"`
	// '#/components/schemas/InstanceFlavorExtendedSerializer/properties/price_status/anyOf/0'
	// "$.components.schemas.InstanceFlavorExtendedSerializer.properties.price_status.anyOf[0]"
	//
	// Any of "error", "hide", "show".
	PriceStatus InstanceFlavorPriceStatus `json:"price_status,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Architecture        resp.Field
		Disabled            resp.Field
		FlavorID            resp.Field
		FlavorName          resp.Field
		OsType              resp.Field
		Ram                 resp.Field
		Vcpus               resp.Field
		Capacity            resp.Field
		CurrencyCode        resp.Field
		HardwareDescription resp.Field
		PricePerHour        resp.Field
		PricePerMonth       resp.Field
		PriceStatus         resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavor) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceFlavorExtendedSerializer/properties/price_status/anyOf/0'
// "$.components.schemas.InstanceFlavorExtendedSerializer.properties.price_status.anyOf[0]"
type InstanceFlavorPriceStatus string

const (
	InstanceFlavorPriceStatusError InstanceFlavorPriceStatus = "error"
	InstanceFlavorPriceStatusHide  InstanceFlavorPriceStatus = "hide"
	InstanceFlavorPriceStatusShow  InstanceFlavorPriceStatus = "show"
)

// '#/components/schemas/InstanceFlavorExtendedCollectionSerializer'
// "$.components.schemas.InstanceFlavorExtendedCollectionSerializer"
type InstanceFlavorList struct {
	// '#/components/schemas/InstanceFlavorExtendedCollectionSerializer/properties/count'
	// "$.components.schemas.InstanceFlavorExtendedCollectionSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/InstanceFlavorExtendedCollectionSerializer/properties/results'
	// "$.components.schemas.InstanceFlavorExtendedCollectionSerializer.properties.results"
	Results []InstanceFlavor `json:"results,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorList) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceFlavorListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/flavors/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/flavors/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/flavors/{project_id}/{region_id}'].get.parameters[2]"
	Disabled param.Opt[bool] `query:"disabled,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/flavors/{project_id}/{region_id}'].get.parameters[3]"
	ExcludeLinux param.Opt[bool] `query:"exclude_linux,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/flavors/{project_id}/{region_id}'].get.parameters[4]"
	ExcludeWindows param.Opt[bool] `query:"exclude_windows,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/flavors/{project_id}/{region_id}'].get.parameters[5]"
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceFlavorListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [InstanceFlavorListParams]'s query parameters as
// `url.Values`.
func (r InstanceFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceFlavorListForResizeParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Favailable_flavors/get/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/available_flavors'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Favailable_flavors/get/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/available_flavors'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Binstance_id%7D%2Favailable_flavors/get/parameters/3'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/{instance_id}/available_flavors'].get.parameters[3]"
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceFlavorListForResizeParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [InstanceFlavorListForResizeParams]'s query parameters as
// `url.Values`.
func (r InstanceFlavorListForResizeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceFlavorListSuitableParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Favailable_flavors/post/parameters/0/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/available_flavors'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Favailable_flavors/post/parameters/1/schema'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/available_flavors'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateInstanceVolumeListSchema/properties/volumes'
	// "$.components.schemas.CreateInstanceVolumeListSchema.properties.volumes"
	Volumes []InstanceFlavorListSuitableParamsVolume `json:"volumes,omitzero,required"`
	// '#/paths/%2Fcloud%2Fv1%2Finstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Favailable_flavors/post/parameters/2'
	// "$.paths['/cloud/v1/instances/{project_id}/{region_id}/available_flavors'].post.parameters[2]"
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceFlavorListSuitableParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InstanceFlavorListSuitableParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceFlavorListSuitableParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// URLQuery serializes [InstanceFlavorListSuitableParams]'s query parameters as
// `url.Values`.
func (r InstanceFlavorListSuitableParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// '#/components/schemas/CreateInstanceVolumeListSchema/properties/volumes/items'
// "$.components.schemas.CreateInstanceVolumeListSchema.properties.volumes.items"
//
// The property Source is required.
type InstanceFlavorListSuitableParamsVolume struct {
	// '#/components/schemas/CheckFlavorVolumeSchema/properties/source'
	// "$.components.schemas.CheckFlavorVolumeSchema.properties.source"
	//
	// Any of "apptemplate", "existing-volume", "image", "new-volume", "snapshot".
	Source string `json:"source,omitzero,required"`
	// '#/components/schemas/CheckFlavorVolumeSchema/properties/name'
	// "$.components.schemas.CheckFlavorVolumeSchema.properties.name"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/CheckFlavorVolumeSchema/properties/apptemplate_id'
	// "$.components.schemas.CheckFlavorVolumeSchema.properties.apptemplate_id"
	ApptemplateID param.Opt[string] `json:"apptemplate_id,omitzero"`
	// '#/components/schemas/CheckFlavorVolumeSchema/properties/boot_index'
	// "$.components.schemas.CheckFlavorVolumeSchema.properties.boot_index"
	BootIndex param.Opt[int64] `json:"boot_index,omitzero"`
	// '#/components/schemas/CheckFlavorVolumeSchema/properties/image_id'
	// "$.components.schemas.CheckFlavorVolumeSchema.properties.image_id"
	ImageID param.Opt[string] `json:"image_id,omitzero"`
	// '#/components/schemas/CheckFlavorVolumeSchema/properties/size'
	// "$.components.schemas.CheckFlavorVolumeSchema.properties.size"
	Size param.Opt[int64] `json:"size,omitzero"`
	// '#/components/schemas/CheckFlavorVolumeSchema/properties/snapshot_id'
	// "$.components.schemas.CheckFlavorVolumeSchema.properties.snapshot_id"
	SnapshotID param.Opt[string] `json:"snapshot_id,omitzero"`
	// '#/components/schemas/CheckFlavorVolumeSchema/properties/volume_id'
	// "$.components.schemas.CheckFlavorVolumeSchema.properties.volume_id"
	VolumeID param.Opt[string] `json:"volume_id,omitzero"`
	// '#/components/schemas/CheckFlavorVolumeSchema/properties/type_name'
	// "$.components.schemas.CheckFlavorVolumeSchema.properties.type_name"
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	TypeName string `json:"type_name,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InstanceFlavorListSuitableParamsVolume) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InstanceFlavorListSuitableParamsVolume) MarshalJSON() (data []byte, err error) {
	type shadow InstanceFlavorListSuitableParamsVolume
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[InstanceFlavorListSuitableParamsVolume](
		"Source", false, "apptemplate", "existing-volume", "image", "new-volume", "snapshot",
	)
	apijson.RegisterFieldValidator[InstanceFlavorListSuitableParamsVolume](
		"TypeName", false, "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}
