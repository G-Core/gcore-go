// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
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

// InstanceFlavorDetailedUnion contains all possible properties and values from
// [InstanceFlavorDetailedInstancesFlavorSchemaWithoutPrice],
// [InstanceFlavorDetailedInstancesFlavorSchemaWithPrice].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InstanceFlavorDetailedUnion struct {
	Architecture        string `json:"architecture"`
	Disabled            bool   `json:"disabled"`
	FlavorID            string `json:"flavor_id"`
	FlavorName          string `json:"flavor_name"`
	HardwareDescription string `json:"hardware_description"`
	OsType              string `json:"os_type"`
	Ram                 int64  `json:"ram"`
	Vcpus               int64  `json:"vcpus"`
	// This field is from variant
	// [InstanceFlavorDetailedInstancesFlavorSchemaWithPrice].
	CurrencyCode string `json:"currency_code"`
	// This field is from variant
	// [InstanceFlavorDetailedInstancesFlavorSchemaWithPrice].
	PricePerHour float64 `json:"price_per_hour"`
	// This field is from variant
	// [InstanceFlavorDetailedInstancesFlavorSchemaWithPrice].
	PricePerMonth float64 `json:"price_per_month"`
	// This field is from variant
	// [InstanceFlavorDetailedInstancesFlavorSchemaWithPrice].
	PriceStatus string `json:"price_status"`
	JSON        struct {
		Architecture        respjson.Field
		Disabled            respjson.Field
		FlavorID            respjson.Field
		FlavorName          respjson.Field
		HardwareDescription respjson.Field
		OsType              respjson.Field
		Ram                 respjson.Field
		Vcpus               respjson.Field
		CurrencyCode        respjson.Field
		PricePerHour        respjson.Field
		PricePerMonth       respjson.Field
		PriceStatus         respjson.Field
		raw                 string
	} `json:"-"`
}

func (u InstanceFlavorDetailedUnion) AsInstancesFlavorSchemaWithoutPrice() (v InstanceFlavorDetailedInstancesFlavorSchemaWithoutPrice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InstanceFlavorDetailedUnion) AsInstancesFlavorSchemaWithPrice() (v InstanceFlavorDetailedInstancesFlavorSchemaWithPrice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InstanceFlavorDetailedUnion) RawJSON() string { return u.JSON.raw }

func (r *InstanceFlavorDetailedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instances flavor schema without price information
type InstanceFlavorDetailedInstancesFlavorSchemaWithoutPrice struct {
	// Flavor architecture type
	Architecture string `json:"architecture" api:"required"`
	// Disabled flavor flag
	Disabled bool `json:"disabled" api:"required"`
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id" api:"required"`
	// Flavor name
	FlavorName string `json:"flavor_name" api:"required"`
	// Additional hardware description
	HardwareDescription map[string]string `json:"hardware_description" api:"required"`
	// Flavor operating system
	OsType string `json:"os_type" api:"required"`
	// RAM size in MiB
	Ram int64 `json:"ram" api:"required"`
	// Virtual CPU count
	Vcpus int64 `json:"vcpus" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture        respjson.Field
		Disabled            respjson.Field
		FlavorID            respjson.Field
		FlavorName          respjson.Field
		HardwareDescription respjson.Field
		OsType              respjson.Field
		Ram                 respjson.Field
		Vcpus               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorDetailedInstancesFlavorSchemaWithoutPrice) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorDetailedInstancesFlavorSchemaWithoutPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instances flavor schema with price information
type InstanceFlavorDetailedInstancesFlavorSchemaWithPrice struct {
	// Flavor architecture type
	Architecture string `json:"architecture" api:"required"`
	// Currency code
	CurrencyCode string `json:"currency_code" api:"required"`
	// Disabled flavor flag
	Disabled bool `json:"disabled" api:"required"`
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id" api:"required"`
	// Flavor name
	FlavorName string `json:"flavor_name" api:"required"`
	// Additional hardware description
	HardwareDescription map[string]string `json:"hardware_description" api:"required"`
	// Flavor operating system
	OsType string `json:"os_type" api:"required"`
	// Price per hour
	PricePerHour float64 `json:"price_per_hour" api:"required"`
	// Price per month
	PricePerMonth float64 `json:"price_per_month" api:"required"`
	// Price status for the UI
	//
	// Any of "error", "hide", "show".
	PriceStatus string `json:"price_status" api:"required"`
	// RAM size in MiB
	Ram int64 `json:"ram" api:"required"`
	// Virtual CPU count
	Vcpus int64 `json:"vcpus" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture        respjson.Field
		CurrencyCode        respjson.Field
		Disabled            respjson.Field
		FlavorID            respjson.Field
		FlavorName          respjson.Field
		HardwareDescription respjson.Field
		OsType              respjson.Field
		PricePerHour        respjson.Field
		PricePerMonth       respjson.Field
		PriceStatus         respjson.Field
		Ram                 respjson.Field
		Vcpus               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorDetailedInstancesFlavorSchemaWithPrice) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorDetailedInstancesFlavorSchemaWithPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceFlavorList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []InstanceFlavorDetailedUnion `json:"results" api:"required"`
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
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
