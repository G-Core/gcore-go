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

// GPUBaremetalClusterFlavorService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUBaremetalClusterFlavorService] method instead.
type GPUBaremetalClusterFlavorService struct {
	Options []option.RequestOption
}

// NewGPUBaremetalClusterFlavorService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGPUBaremetalClusterFlavorService(opts ...option.RequestOption) (r GPUBaremetalClusterFlavorService) {
	r = GPUBaremetalClusterFlavorService{}
	r.Options = opts
	return
}

// List bare metal GPU flavors
func (r *GPUBaremetalClusterFlavorService) List(ctx context.Context, params GPUBaremetalClusterFlavorListParams, opts ...option.RequestOption) (res *GPUBaremetalFlavorList, err error) {
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
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/flavors", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// GPUBaremetalFlavorUnion contains all possible properties and values from
// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPrice],
// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPrice].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type GPUBaremetalFlavorUnion struct {
	Architecture string `json:"architecture"`
	Capacity     int64  `json:"capacity"`
	Disabled     bool   `json:"disabled"`
	// This field is a union of
	// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareDescription],
	// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareDescription]
	HardwareDescription GPUBaremetalFlavorUnionHardwareDescription `json:"hardware_description"`
	// This field is a union of
	// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareProperties],
	// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareProperties]
	HardwareProperties GPUBaremetalFlavorUnionHardwareProperties `json:"hardware_properties"`
	Name               string                                    `json:"name"`
	// This field is a union of
	// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceSupportedFeatures],
	// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceSupportedFeatures]
	SupportedFeatures GPUBaremetalFlavorUnionSupportedFeatures `json:"supported_features"`
	// This field is from variant
	// [GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPrice].
	Price GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPricePrice `json:"price"`
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

func (u GPUBaremetalFlavorUnion) AsBareMetalGPUFlavorsChemaWithoutPrice() (v GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPrice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GPUBaremetalFlavorUnion) AsBareMetalGPUFlavorsChemaWithPrice() (v GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPrice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GPUBaremetalFlavorUnion) RawJSON() string { return u.JSON.raw }

func (r *GPUBaremetalFlavorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUBaremetalFlavorUnionHardwareDescription is an implicit subunion of
// [GPUBaremetalFlavorUnion]. GPUBaremetalFlavorUnionHardwareDescription provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [GPUBaremetalFlavorUnion].
type GPUBaremetalFlavorUnionHardwareDescription struct {
	CPU     string `json:"cpu"`
	Disk    string `json:"disk"`
	GPU     string `json:"gpu"`
	Network string `json:"network"`
	Ram     string `json:"ram"`
	JSON    struct {
		CPU     respjson.Field
		Disk    respjson.Field
		GPU     respjson.Field
		Network respjson.Field
		Ram     respjson.Field
		raw     string
	} `json:"-"`
}

func (r *GPUBaremetalFlavorUnionHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUBaremetalFlavorUnionHardwareProperties is an implicit subunion of
// [GPUBaremetalFlavorUnion]. GPUBaremetalFlavorUnionHardwareProperties provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [GPUBaremetalFlavorUnion].
type GPUBaremetalFlavorUnionHardwareProperties struct {
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

func (r *GPUBaremetalFlavorUnionHardwareProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUBaremetalFlavorUnionSupportedFeatures is an implicit subunion of
// [GPUBaremetalFlavorUnion]. GPUBaremetalFlavorUnionSupportedFeatures provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [GPUBaremetalFlavorUnion].
type GPUBaremetalFlavorUnionSupportedFeatures struct {
	SecurityGroups bool `json:"security_groups"`
	JSON           struct {
		SecurityGroups respjson.Field
		raw            string
	} `json:"-"`
}

func (r *GPUBaremetalFlavorUnionSupportedFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPrice struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required"`
	// Number of available instances of given flavor
	Capacity int64 `json:"capacity,required"`
	// If the flavor is disabled, new resources cannot be created using this flavor.
	Disabled bool `json:"disabled,required"`
	// Additional bare metal hardware description
	HardwareDescription GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareDescription `json:"hardware_description,required"`
	// Additional bare metal hardware properties
	HardwareProperties GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareProperties `json:"hardware_properties,required"`
	// Flavor name
	Name string `json:"name,required"`
	// Set of enabled features based on the flavor's type and configuration
	SupportedFeatures GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceSupportedFeatures `json:"supported_features,required"`
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
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPrice) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional bare metal hardware description
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareDescription struct {
	// Human-readable CPU description
	CPU string `json:"cpu,required"`
	// Human-readable disk description
	Disk string `json:"disk,required"`
	// Human-readable GPU description
	GPU string `json:"gpu,required"`
	// Human-readable NIC description
	Network string `json:"network,required"`
	// Human-readable RAM description
	Ram string `json:"ram,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU         respjson.Field
		Disk        respjson.Field
		GPU         respjson.Field
		Network     respjson.Field
		Ram         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareDescription) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional bare metal hardware properties
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareProperties struct {
	// The total count of available GPUs.
	GPUCount int64 `json:"gpu_count,required"`
	// The manufacturer of the graphics processing GPU
	GPUManufacturer string `json:"gpu_manufacturer,required"`
	// GPU model
	GPUModel string `json:"gpu_model,required"`
	// The configuration of the Ethernet ports
	NicEth string `json:"nic_eth,required"`
	// The configuration of the InfiniBand ports
	NicIb string `json:"nic_ib,required"`
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
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceHardwareProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Set of enabled features based on the flavor's type and configuration
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceSupportedFeatures struct {
	SecurityGroups bool `json:"security_groups,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SecurityGroups respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceSupportedFeatures) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPriceSupportedFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPrice struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required"`
	// Number of available instances of given flavor
	Capacity int64 `json:"capacity,required"`
	// If the flavor is disabled, new resources cannot be created using this flavor.
	Disabled bool `json:"disabled,required"`
	// Additional virtual hardware description
	HardwareDescription GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareDescription `json:"hardware_description,required"`
	// Additional bare metal hardware properties
	HardwareProperties GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareProperties `json:"hardware_properties,required"`
	// Flavor name
	Name string `json:"name,required"`
	// Flavor price
	Price GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPricePrice `json:"price,required"`
	// Set of enabled features based on the flavor's type and configuration
	SupportedFeatures GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceSupportedFeatures `json:"supported_features,required"`
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
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPrice) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional virtual hardware description
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareDescription struct {
	// Human-readable CPU description
	CPU string `json:"cpu,required"`
	// Human-readable disk description
	Disk string `json:"disk,required"`
	// Human-readable GPU description
	GPU string `json:"gpu,required"`
	// Human-readable NIC description
	Network string `json:"network,required"`
	// Human-readable RAM description
	Ram string `json:"ram,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU         respjson.Field
		Disk        respjson.Field
		GPU         respjson.Field
		Network     respjson.Field
		Ram         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareDescription) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional bare metal hardware properties
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareProperties struct {
	// The total count of available GPUs.
	GPUCount int64 `json:"gpu_count,required"`
	// The manufacturer of the graphics processing GPU
	GPUManufacturer string `json:"gpu_manufacturer,required"`
	// GPU model
	GPUModel string `json:"gpu_model,required"`
	// The configuration of the Ethernet ports
	NicEth string `json:"nic_eth,required"`
	// The configuration of the InfiniBand ports
	NicIb string `json:"nic_ib,required"`
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
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareProperties) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceHardwareProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor price
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPricePrice struct {
	// Currency code. Shown if the `include_prices` query parameter if set to true
	CurrencyCode string `json:"currency_code,required"`
	// Price per hour. Shown if the `include_prices` query parameter if set to true
	PricePerHour float64 `json:"price_per_hour,required"`
	// Price per month. Shown if the `include_prices` query parameter if set to true
	PricePerMonth float64 `json:"price_per_month,required"`
	// Price status for the UI
	//
	// Any of "error", "hide", "show".
	PriceStatus string `json:"price_status,required"`
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
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPricePrice) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPricePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Set of enabled features based on the flavor's type and configuration
type GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceSupportedFeatures struct {
	SecurityGroups bool `json:"security_groups,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SecurityGroups respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceSupportedFeatures) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPriceSupportedFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalFlavorList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []GPUBaremetalFlavorUnion `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalFlavorList) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalFlavorList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterFlavorListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Set to `true` to remove the disabled flavors from the response.
	HideDisabled param.Opt[bool] `query:"hide_disabled,omitzero" json:"-"`
	// Set to `true` if the response should include flavor prices.
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GPUBaremetalClusterFlavorListParams]'s query parameters as
// `url.Values`.
func (r GPUBaremetalClusterFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
