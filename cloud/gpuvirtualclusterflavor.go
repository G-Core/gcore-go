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

// GPUVirtualClusterFlavorService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUVirtualClusterFlavorService] method instead.
type GPUVirtualClusterFlavorService struct {
	Options []option.RequestOption
}

// NewGPUVirtualClusterFlavorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGPUVirtualClusterFlavorService(opts ...option.RequestOption) (r GPUVirtualClusterFlavorService) {
	r = GPUVirtualClusterFlavorService{}
	r.Options = opts
	return
}

// List virtual GPU flavors
func (r *GPUVirtualClusterFlavorService) List(ctx context.Context, params GPUVirtualClusterFlavorListParams, opts ...option.RequestOption) (res *GPUVirtualFlavorList, err error) {
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
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/flavors", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// GPUVirtualFlavorUnion contains all possible properties and values from
// [GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPrice],
// [GPUVirtualFlavorVirtualGPUFlavorsChemaWithPrice].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type GPUVirtualFlavorUnion struct {
	Architecture string `json:"architecture"`
	Capacity     int64  `json:"capacity"`
	Disabled     bool   `json:"disabled"`
	// This field is a union of
	// [GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceHardwareDescription],
	// [GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceHardwareDescription]
	HardwareDescription GPUVirtualFlavorUnionHardwareDescription `json:"hardware_description"`
	// This field is a union of
	// [GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceHardwareProperties],
	// [GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceHardwareProperties]
	HardwareProperties GPUVirtualFlavorUnionHardwareProperties `json:"hardware_properties"`
	Name               string                                  `json:"name"`
	// This field is a union of
	// [GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceSupportedFeatures],
	// [GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceSupportedFeatures]
	SupportedFeatures GPUVirtualFlavorUnionSupportedFeatures `json:"supported_features"`
	// This field is from variant [GPUVirtualFlavorVirtualGPUFlavorsChemaWithPrice].
	Price GPUVirtualFlavorVirtualGPUFlavorsChemaWithPricePrice `json:"price"`
	JSON  struct {
		Architecture        respjson.Field
		Capacity            respjson.Field
		Disabled            respjson.Field
		HardwareDescription respjson.Field
		HardwareProperties  respjson.Field
		Name                respjson.Field
		SupportedFeatures   respjson.Field
		Price               respjson.Field
		raw                 string
	} `json:"-"`
}

func (u GPUVirtualFlavorUnion) AsVirtualGPUFlavorsChemaWithoutPrice() (v GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPrice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GPUVirtualFlavorUnion) AsVirtualGPUFlavorsChemaWithPrice() (v GPUVirtualFlavorVirtualGPUFlavorsChemaWithPrice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GPUVirtualFlavorUnion) RawJSON() string { return u.JSON.raw }

func (r *GPUVirtualFlavorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUVirtualFlavorUnionHardwareDescription is an implicit subunion of
// [GPUVirtualFlavorUnion]. GPUVirtualFlavorUnionHardwareDescription provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [GPUVirtualFlavorUnion].
type GPUVirtualFlavorUnionHardwareDescription struct {
	GPU          string `json:"gpu"`
	LocalStorage int64  `json:"local_storage"`
	Ram          int64  `json:"ram"`
	Vcpus        int64  `json:"vcpus"`
	JSON         struct {
		GPU          respjson.Field
		LocalStorage respjson.Field
		Ram          respjson.Field
		Vcpus        respjson.Field
		raw          string
	} `json:"-"`
}

func (r *GPUVirtualFlavorUnionHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUVirtualFlavorUnionHardwareProperties is an implicit subunion of
// [GPUVirtualFlavorUnion]. GPUVirtualFlavorUnionHardwareProperties provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [GPUVirtualFlavorUnion].
type GPUVirtualFlavorUnionHardwareProperties struct {
	GPUCount        int64  `json:"gpu_count"`
	GPUManufacturer string `json:"gpu_manufacturer"`
	GPUModel        string `json:"gpu_model"`
	NicEth          string `json:"nic_eth"`
	NicIb           string `json:"nic_ib"`
	JSON            struct {
		GPUCount        respjson.Field
		GPUManufacturer respjson.Field
		GPUModel        respjson.Field
		NicEth          respjson.Field
		NicIb           respjson.Field
		raw             string
	} `json:"-"`
}

func (r *GPUVirtualFlavorUnionHardwareProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUVirtualFlavorUnionSupportedFeatures is an implicit subunion of
// [GPUVirtualFlavorUnion]. GPUVirtualFlavorUnionSupportedFeatures provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [GPUVirtualFlavorUnion].
type GPUVirtualFlavorUnionSupportedFeatures struct {
	SecurityGroups bool `json:"security_groups"`
	JSON           struct {
		SecurityGroups respjson.Field
		raw            string
	} `json:"-"`
}

func (r *GPUVirtualFlavorUnionSupportedFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPrice struct {
	// Flavor architecture type
	Architecture string `json:"architecture" api:"required"`
	// Number of available instances of given flavor
	Capacity int64 `json:"capacity" api:"required"`
	// If the flavor is disabled, new resources cannot be created using this flavor.
	Disabled bool `json:"disabled" api:"required"`
	// Additional virtual hardware description
	HardwareDescription GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceHardwareDescription `json:"hardware_description" api:"required"`
	// Additional virtual hardware properties
	HardwareProperties GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceHardwareProperties `json:"hardware_properties" api:"required"`
	// Flavor name
	Name string `json:"name" api:"required"`
	// Set of enabled features based on the flavor's type and configuration
	SupportedFeatures GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceSupportedFeatures `json:"supported_features" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture        respjson.Field
		Capacity            respjson.Field
		Disabled            respjson.Field
		HardwareDescription respjson.Field
		HardwareProperties  respjson.Field
		Name                respjson.Field
		SupportedFeatures   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPrice) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional virtual hardware description
type GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceHardwareDescription struct {
	// Human-readable GPU description
	GPU string `json:"gpu" api:"required"`
	// Local storage capacity in GiB
	LocalStorage int64 `json:"local_storage" api:"required"`
	// RAM size in MiB
	Ram int64 `json:"ram" api:"required"`
	// Virtual CPU count
	Vcpus int64 `json:"vcpus" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GPU          respjson.Field
		LocalStorage respjson.Field
		Ram          respjson.Field
		Vcpus        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceHardwareDescription) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional virtual hardware properties
type GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceHardwareProperties struct {
	// The total count of available GPUs.
	GPUCount int64 `json:"gpu_count" api:"required"`
	// The manufacturer of the graphics processing GPU
	GPUManufacturer string `json:"gpu_manufacturer" api:"required"`
	// GPU model
	GPUModel string `json:"gpu_model" api:"required"`
	// The configuration of the Ethernet ports
	NicEth string `json:"nic_eth" api:"required"`
	// The configuration of the InfiniBand ports
	NicIb string `json:"nic_ib" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GPUCount        respjson.Field
		GPUManufacturer respjson.Field
		GPUModel        respjson.Field
		NicEth          respjson.Field
		NicIb           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceHardwareProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceHardwareProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Set of enabled features based on the flavor's type and configuration
type GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceSupportedFeatures struct {
	SecurityGroups bool `json:"security_groups" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SecurityGroups respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceSupportedFeatures) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUVirtualFlavorVirtualGPUFlavorsChemaWithoutPriceSupportedFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualFlavorVirtualGPUFlavorsChemaWithPrice struct {
	// Flavor architecture type
	Architecture string `json:"architecture" api:"required"`
	// Number of available instances of given flavor
	Capacity int64 `json:"capacity" api:"required"`
	// If the flavor is disabled, new resources cannot be created using this flavor.
	Disabled bool `json:"disabled" api:"required"`
	// Additional virtual hardware description
	HardwareDescription GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceHardwareDescription `json:"hardware_description" api:"required"`
	// Additional virtual hardware properties
	HardwareProperties GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceHardwareProperties `json:"hardware_properties" api:"required"`
	// Flavor name
	Name string `json:"name" api:"required"`
	// Flavor price.
	Price GPUVirtualFlavorVirtualGPUFlavorsChemaWithPricePrice `json:"price" api:"required"`
	// Set of enabled features based on the flavor's type and configuration
	SupportedFeatures GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceSupportedFeatures `json:"supported_features" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture        respjson.Field
		Capacity            respjson.Field
		Disabled            respjson.Field
		HardwareDescription respjson.Field
		HardwareProperties  respjson.Field
		Name                respjson.Field
		Price               respjson.Field
		SupportedFeatures   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualFlavorVirtualGPUFlavorsChemaWithPrice) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualFlavorVirtualGPUFlavorsChemaWithPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional virtual hardware description
type GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceHardwareDescription struct {
	// Human-readable GPU description
	GPU string `json:"gpu" api:"required"`
	// Local storage capacity in GiB
	LocalStorage int64 `json:"local_storage" api:"required"`
	// RAM size in MiB
	Ram int64 `json:"ram" api:"required"`
	// Virtual CPU count
	Vcpus int64 `json:"vcpus" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GPU          respjson.Field
		LocalStorage respjson.Field
		Ram          respjson.Field
		Vcpus        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceHardwareDescription) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional virtual hardware properties
type GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceHardwareProperties struct {
	// The total count of available GPUs.
	GPUCount int64 `json:"gpu_count" api:"required"`
	// The manufacturer of the graphics processing GPU
	GPUManufacturer string `json:"gpu_manufacturer" api:"required"`
	// GPU model
	GPUModel string `json:"gpu_model" api:"required"`
	// The configuration of the Ethernet ports
	NicEth string `json:"nic_eth" api:"required"`
	// The configuration of the InfiniBand ports
	NicIb string `json:"nic_ib" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GPUCount        respjson.Field
		GPUManufacturer respjson.Field
		GPUModel        respjson.Field
		NicEth          respjson.Field
		NicIb           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceHardwareProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceHardwareProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor price.
type GPUVirtualFlavorVirtualGPUFlavorsChemaWithPricePrice struct {
	// Currency code. Shown if the `include_prices` query parameter if set to true
	CurrencyCode string `json:"currency_code" api:"required"`
	// Price per hour. Shown if the `include_prices` query parameter if set to true
	PricePerHour float64 `json:"price_per_hour" api:"required"`
	// Price per month. Shown if the `include_prices` query parameter if set to true
	PricePerMonth float64 `json:"price_per_month" api:"required"`
	// Price status for the UI
	//
	// Any of "error", "hide", "show".
	PriceStatus string `json:"price_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrencyCode  respjson.Field
		PricePerHour  respjson.Field
		PricePerMonth respjson.Field
		PriceStatus   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualFlavorVirtualGPUFlavorsChemaWithPricePrice) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualFlavorVirtualGPUFlavorsChemaWithPricePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Set of enabled features based on the flavor's type and configuration
type GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceSupportedFeatures struct {
	SecurityGroups bool `json:"security_groups" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SecurityGroups respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceSupportedFeatures) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUVirtualFlavorVirtualGPUFlavorsChemaWithPriceSupportedFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualFlavorList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []GPUVirtualFlavorUnion `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualFlavorList) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualFlavorList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterFlavorListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Set to `true` to remove the disabled flavors from the response.
	HideDisabled param.Opt[bool] `query:"hide_disabled,omitzero" json:"-"`
	// Set to `true` if the response should include flavor prices.
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GPUVirtualClusterFlavorListParams]'s query parameters as
// `url.Values`.
func (r GPUVirtualClusterFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
