// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

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

// Retrieve a list of available instance flavors in the project and region. When
// `include_prices` is specified, the list includes pricing information. Trial mode
// clients see all prices as 0. Contact support for pricing errors.
func (r *InstanceFlavorService) List(ctx context.Context, params InstanceFlavorListParams, opts ...option.RequestOption) (res *InstanceFlavorList, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
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
	// Currency code. Shown if the `include_prices` query parameter if set to true
	CurrencyCode string `json:"currency_code,nullable"`
	// Additional hardware description
	HardwareDescription map[string]string `json:"hardware_description"`
	// Price per hour. Shown if the `include_prices` query parameter if set to true
	PricePerHour float64 `json:"price_per_hour,nullable"`
	// Price per month. Shown if the `include_prices` query parameter if set to true
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
