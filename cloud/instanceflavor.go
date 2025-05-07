// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

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
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
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
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
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
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
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
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/available_flavors", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Instances flavor schema
type InstanceFlavor struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required"`
	// Disabled flavor flag
	Disabled bool `json:"disabled,required"`
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id,required"`
	// Flavor name
	FlavorName string `json:"flavor_name,required"`
	// Flavor operating system
	OsType string `json:"os_type,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Virtual CPU count. For bare metal flavors, it's a physical CPU count
	Vcpus int64 `json:"vcpus,required"`
	// Number of available instances of given configuration
	Capacity int64 `json:"capacity,nullable"`
	// Currency code. Shown if the include_prices query parameter if set to true
	CurrencyCode string `json:"currency_code,nullable"`
	// Additional hardware description
	HardwareDescription map[string]string `json:"hardware_description"`
	// Price per hour. Shown if the include_prices query parameter if set to true
	PricePerHour float64 `json:"price_per_hour,nullable"`
	// Price per month. Shown if the include_prices query parameter if set to true
	PricePerMonth float64 `json:"price_per_month,nullable"`
	// Price status for the UI
	//
	// Any of "error", "hide", "show".
	PriceStatus InstanceFlavorPriceStatus `json:"price_status,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture        respjson.Field
		Disabled            respjson.Field
		FlavorID            respjson.Field
		FlavorName          respjson.Field
		OsType              respjson.Field
		Ram                 respjson.Field
		Vcpus               respjson.Field
		Capacity            respjson.Field
		CurrencyCode        respjson.Field
		HardwareDescription respjson.Field
		PricePerHour        respjson.Field
		PricePerMonth       respjson.Field
		PriceStatus         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavor) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price status for the UI
type InstanceFlavorPriceStatus string

const (
	InstanceFlavorPriceStatusError InstanceFlavorPriceStatus = "error"
	InstanceFlavorPriceStatusHide  InstanceFlavorPriceStatus = "hide"
	InstanceFlavorPriceStatusShow  InstanceFlavorPriceStatus = "show"
)

type InstanceFlavorList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []InstanceFlavor `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorList) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceFlavorListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Flag for filtering disabled flavors in the region. Defaults to true
	Disabled param.Opt[bool] `query:"disabled,omitzero" json:"-"`
	// Set to true to exclude flavors dedicated to linux images. Default False
	ExcludeLinux param.Opt[bool] `query:"exclude_linux,omitzero" json:"-"`
	// Set to true to exclude flavors dedicated to windows images. Default False
	ExcludeWindows param.Opt[bool] `query:"exclude_windows,omitzero" json:"-"`
	// Set to true if the response should include flavor prices
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceFlavorListParams]'s query parameters as
// `url.Values`.
func (r InstanceFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceFlavorListForResizeParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Set to true if flavor listing should include flavor prices
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	paramObj
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Volumes details. Non-important info such as names may be omitted.
	Volumes []InstanceFlavorListSuitableParamsVolume `json:"volumes,omitzero,required"`
	// Set to true if flavor listing should include flavor prices
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	paramObj
}

func (r InstanceFlavorListSuitableParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceFlavorListSuitableParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceFlavorListSuitableParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [InstanceFlavorListSuitableParams]'s query parameters as
// `url.Values`.
func (r InstanceFlavorListSuitableParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The property Source is required.
type InstanceFlavorListSuitableParamsVolume struct {
	// Volume source
	//
	// Any of "apptemplate", "existing-volume", "image", "new-volume", "snapshot".
	Source string            `json:"source,omitzero,required"`
	Name   param.Opt[string] `json:"name,omitzero"`
	// App template ID. Mandatory if volume is created from marketplace template
	ApptemplateID param.Opt[string] `json:"apptemplate_id,omitzero"`
	// 0 should be set for primary boot device Unique positive values for other
	// bootable devices.Negative - boot prohibited
	BootIndex param.Opt[int64] `json:"boot_index,omitzero"`
	// Image ID. Mandatory if volume is created from image
	ImageID param.Opt[string] `json:"image_id,omitzero"`
	// Volume size. Must be specified when source is 'new-volume' or 'image'. If
	// specified for source 'snapshot' or 'existing-volume', value must be equal to
	// respective snapshot or volume size
	Size param.Opt[int64] `json:"size,omitzero"`
	// Volume snapshot ID. Mandatory if volume is created from a snapshot
	SnapshotID param.Opt[string] `json:"snapshot_id,omitzero"`
	// Volume ID. Mandatory if volume is pre-existing volume
	VolumeID param.Opt[string] `json:"volume_id,omitzero"`
	// One of 'standard', 'ssd_hiiops', 'ssd_local', 'ssd_lowlatency', 'cold', 'ultra'
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	TypeName string `json:"type_name,omitzero"`
	paramObj
}

func (r InstanceFlavorListSuitableParamsVolume) MarshalJSON() (data []byte, err error) {
	type shadow InstanceFlavorListSuitableParamsVolume
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceFlavorListSuitableParamsVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InstanceFlavorListSuitableParamsVolume](
		"source", "apptemplate", "existing-volume", "image", "new-volume", "snapshot",
	)
	apijson.RegisterFieldValidator[InstanceFlavorListSuitableParamsVolume](
		"type_name", "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}
