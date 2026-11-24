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
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// Instance flavors define available CPU, memory, and disk configurations for
// creating cloud instances.
//
// InstancesRegionalFlavorService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstancesRegionalFlavorService] method instead.
type InstancesRegionalFlavorService struct {
	Options []option.RequestOption
}

// NewInstancesRegionalFlavorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInstancesRegionalFlavorService(opts ...option.RequestOption) (r InstancesRegionalFlavorService) {
	r = InstancesRegionalFlavorService{}
	r.Options = opts
	return
}

// Use
// [GET /v1/projects/{`project_id`}/instances/flavors](#operation/GlobalInstanceFlavorListView.get)
// instead. This route also returns the basic VM and the bare metal flavors. For
// them, use [List basic VM flavors](#operation/GlobalBasicVmFlavorListView.get)
// and [List bare metal flavors](#operation/GlobalBaremetalFlavorListView.get). The
// new routes do not return prices. For prices, use
// [List instance flavor prices](#operation/InstanceFlavorPriceListView.get),
// [List basic VM flavor prices](#operation/BasicVmFlavorPriceListView.get) and
// [List bare metal flavor prices](#operation/BaremetalFlavorPriceListView.get).
// This route continues to work. Later, we will deprecate it and then remove it. We
// will tell you in advance.
//
// Retrieve a list of available instance flavors in the project and region. When
// `include_prices` is specified, the list includes pricing information. Trial mode
// clients see all prices as 0. Contact support for pricing errors.
func (r *InstancesRegionalFlavorService) List(ctx context.Context, params InstancesRegionalFlavorListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[InstanceRegionalFlavorDetailedUnion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/flavors/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Use
// [GET /v1/projects/{`project_id`}/instances/flavors](#operation/GlobalInstanceFlavorListView.get)
// instead. This route also returns the basic VM and the bare metal flavors. For
// them, use [List basic VM flavors](#operation/GlobalBasicVmFlavorListView.get)
// and [List bare metal flavors](#operation/GlobalBaremetalFlavorListView.get). The
// new routes do not return prices. For prices, use
// [List instance flavor prices](#operation/InstanceFlavorPriceListView.get),
// [List basic VM flavor prices](#operation/BasicVmFlavorPriceListView.get) and
// [List bare metal flavor prices](#operation/BaremetalFlavorPriceListView.get).
// This route continues to work. Later, we will deprecate it and then remove it. We
// will tell you in advance.
//
// Retrieve a list of available instance flavors in the project and region. When
// `include_prices` is specified, the list includes pricing information. Trial mode
// clients see all prices as 0. Contact support for pricing errors.
func (r *InstancesRegionalFlavorService) ListAutoPaging(ctx context.Context, params InstancesRegionalFlavorListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[InstanceRegionalFlavorDetailedUnion] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// InstanceRegionalFlavorDetailedUnion contains all possible properties and values
// from [InstanceRegionalFlavorDetailedInstancesFlavorSchemaWithoutPrice],
// [InstanceRegionalFlavorDetailedInstancesFlavorSchemaWithPrice].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InstanceRegionalFlavorDetailedUnion struct {
	Architecture        string  `json:"architecture"`
	CurrencyCode        string  `json:"currency_code"`
	Disabled            bool    `json:"disabled"`
	FlavorID            string  `json:"flavor_id"`
	FlavorName          string  `json:"flavor_name"`
	HardwareDescription string  `json:"hardware_description"`
	OsType              string  `json:"os_type"`
	PricePerHour        float64 `json:"price_per_hour"`
	PricePerMonth       float64 `json:"price_per_month"`
	PriceStatus         string  `json:"price_status"`
	Ram                 int64   `json:"ram"`
	Vcpus               int64   `json:"vcpus"`
	JSON                struct {
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
		raw                 string
	} `json:"-"`
}

func (u InstanceRegionalFlavorDetailedUnion) AsInstancesFlavorSchemaWithoutPrice() (v InstanceRegionalFlavorDetailedInstancesFlavorSchemaWithoutPrice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InstanceRegionalFlavorDetailedUnion) AsInstancesFlavorSchemaWithPrice() (v InstanceRegionalFlavorDetailedInstancesFlavorSchemaWithPrice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InstanceRegionalFlavorDetailedUnion) RawJSON() string { return u.JSON.raw }

func (r *InstanceRegionalFlavorDetailedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instances flavor schema without price information
type InstanceRegionalFlavorDetailedInstancesFlavorSchemaWithoutPrice struct {
	// Flavor architecture type
	Architecture string `json:"architecture" api:"required"`
	// Currency code. The value is null if the price is not available.
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
	// Price per hour. The value is null if the price is not available.
	PricePerHour float64 `json:"price_per_hour" api:"required"`
	// Price per month. The value is null if the price is not available.
	PricePerMonth float64 `json:"price_per_month" api:"required"`
	// Price status for the UI. The value is null if the price is not available.
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
func (r InstanceRegionalFlavorDetailedInstancesFlavorSchemaWithoutPrice) RawJSON() string {
	return r.JSON.raw
}
func (r *InstanceRegionalFlavorDetailedInstancesFlavorSchemaWithoutPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instances flavor schema with price information
type InstanceRegionalFlavorDetailedInstancesFlavorSchemaWithPrice struct {
	// Flavor architecture type
	Architecture string `json:"architecture" api:"required"`
	// Currency code. The value is null if the price is not available.
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
	// Price per hour. The value is null if the price is not available.
	PricePerHour float64 `json:"price_per_hour" api:"required"`
	// Price per month. The value is null if the price is not available.
	PricePerMonth float64 `json:"price_per_month" api:"required"`
	// Price status for the UI. The value is null if the price is not available.
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
func (r InstanceRegionalFlavorDetailedInstancesFlavorSchemaWithPrice) RawJSON() string {
	return r.JSON.raw
}
func (r *InstanceRegionalFlavorDetailedInstancesFlavorSchemaWithPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceRegionalFlavorList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []InstanceRegionalFlavorDetailedUnion `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceRegionalFlavorList) RawJSON() string { return r.JSON.raw }
func (r *InstanceRegionalFlavorList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstancesRegionalFlavorListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Flag for filtering disabled flavors in the region. Defaults to true
	Disabled param.Opt[bool] `query:"disabled,omitzero" json:"-"`
	// Set to true to exclude flavors dedicated to linux images
	ExcludeLinux param.Opt[bool] `query:"exclude_linux,omitzero" json:"-"`
	// Set to true to exclude flavors dedicated to windows images
	ExcludeWindows param.Opt[bool] `query:"exclude_windows,omitzero" json:"-"`
	// Set to true if the response should include flavor prices
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstancesRegionalFlavorListParams]'s query parameters as
// `url.Values`.
func (r InstancesRegionalFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
