// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/tidwall/gjson"
)

// CloudV3GPUBaremetalService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV3GPUBaremetalService] method instead.
type CloudV3GPUBaremetalService struct {
	Options []option.RequestOption
	Images  *CloudV3GPUBaremetalImageService
}

// NewCloudV3GPUBaremetalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV3GPUBaremetalService(opts ...option.RequestOption) (r *CloudV3GPUBaremetalService) {
	r = &CloudV3GPUBaremetalService{}
	r.Options = opts
	r.Images = NewCloudV3GPUBaremetalImageService(opts...)
	return
}

// List bare metal GPU flavors
func (r *CloudV3GPUBaremetalService) ListFlavors(ctx context.Context, projectID int64, regionID int64, query CloudV3GPUBaremetalListFlavorsParams, opts ...option.RequestOption) (res *CloudV3GPUBaremetalListFlavorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/flavors", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type FlavorHardwareDescriptionBaremetal struct {
	// Human-readable CPU description
	CPU string `json:"cpu,required,nullable"`
	// Human-readable disk description
	Disk string `json:"disk,required,nullable"`
	// Human-readable GPU description
	GPU string `json:"gpu,required,nullable"`
	// Human-readable NIC description
	Network string `json:"network,required,nullable"`
	// Human-readable RAM description
	Ram  string                                 `json:"ram,required,nullable"`
	JSON flavorHardwareDescriptionBaremetalJSON `json:"-"`
}

// flavorHardwareDescriptionBaremetalJSON contains the JSON metadata for the struct
// [FlavorHardwareDescriptionBaremetal]
type flavorHardwareDescriptionBaremetalJSON struct {
	CPU         apijson.Field
	Disk        apijson.Field
	GPU         apijson.Field
	Network     apijson.Field
	Ram         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FlavorHardwareDescriptionBaremetal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r flavorHardwareDescriptionBaremetalJSON) RawJSON() string {
	return r.raw
}

type FlavorHardwareProperties struct {
	// The total count of available GPUs.
	GPUCount int64 `json:"gpu_count,required,nullable"`
	// The manufacturer of the graphics processing GPU
	GPUManufacturer string `json:"gpu_manufacturer,required,nullable"`
	// GPU model
	GPUModel string                       `json:"gpu_model,required,nullable"`
	JSON     flavorHardwarePropertiesJSON `json:"-"`
}

// flavorHardwarePropertiesJSON contains the JSON metadata for the struct
// [FlavorHardwareProperties]
type flavorHardwarePropertiesJSON struct {
	GPUCount        apijson.Field
	GPUManufacturer apijson.Field
	GPUModel        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FlavorHardwareProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r flavorHardwarePropertiesJSON) RawJSON() string {
	return r.raw
}

type FlavorPrice struct {
	// Currency code. Shown if the include_prices query parameter if set to true
	CurrencyCode string `json:"currency_code,required,nullable"`
	// Price per hour. Shown if the include_prices query parameter if set to true
	PricePerHour float64 `json:"price_per_hour,required,nullable"`
	// Price per month. Shown if the include_prices query parameter if set to true
	PricePerMonth float64 `json:"price_per_month,required,nullable"`
	// Price status for the UI
	PriceStatus PriceDisplayStatus `json:"price_status,required,nullable"`
	JSON        flavorPriceJSON    `json:"-"`
}

// flavorPriceJSON contains the JSON metadata for the struct [FlavorPrice]
type flavorPriceJSON struct {
	CurrencyCode  apijson.Field
	PricePerHour  apijson.Field
	PricePerMonth apijson.Field
	PriceStatus   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FlavorPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r flavorPriceJSON) RawJSON() string {
	return r.raw
}

type CloudV3GPUBaremetalListFlavorsResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV3GPUBaremetalListFlavorsResponseResult `json:"results,required"`
	JSON    cloudV3GPUBaremetalListFlavorsResponseJSON     `json:"-"`
}

// cloudV3GPUBaremetalListFlavorsResponseJSON contains the JSON metadata for the
// struct [CloudV3GPUBaremetalListFlavorsResponse]
type cloudV3GPUBaremetalListFlavorsResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV3GPUBaremetalListFlavorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3GPUBaremetalListFlavorsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV3GPUBaremetalListFlavorsResponseResult struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required,nullable"`
	// Number of available instances of given flavor
	Capacity int64 `json:"capacity,required"`
	// If the flavor is disabled, new resources cannot be created using this flavor.
	Disabled bool `json:"disabled,required"`
	// Additional bare metal hardware description
	HardwareDescription FlavorHardwareDescriptionBaremetal `json:"hardware_description,required"`
	// Additional bare metal hardware properties
	HardwareProperties FlavorHardwareProperties `json:"hardware_properties,required"`
	// Flavor name
	Name string `json:"name,required"`
	// Flavor price
	Price FlavorPrice                                      `json:"price"`
	JSON  cloudV3GPUBaremetalListFlavorsResponseResultJSON `json:"-"`
	union CloudV3GPUBaremetalListFlavorsResponseResultsUnion
}

// cloudV3GPUBaremetalListFlavorsResponseResultJSON contains the JSON metadata for
// the struct [CloudV3GPUBaremetalListFlavorsResponseResult]
type cloudV3GPUBaremetalListFlavorsResponseResultJSON struct {
	Architecture        apijson.Field
	Capacity            apijson.Field
	Disabled            apijson.Field
	HardwareDescription apijson.Field
	HardwareProperties  apijson.Field
	Name                apijson.Field
	Price               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r cloudV3GPUBaremetalListFlavorsResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *CloudV3GPUBaremetalListFlavorsResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = CloudV3GPUBaremetalListFlavorsResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CloudV3GPUBaremetalListFlavorsResponseResultsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithoutPrice],
// [CloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithPrices].
func (r CloudV3GPUBaremetalListFlavorsResponseResult) AsUnion() CloudV3GPUBaremetalListFlavorsResponseResultsUnion {
	return r.union
}

// Union satisfied by
// [CloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithoutPrice]
// or
// [CloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithPrices].
type CloudV3GPUBaremetalListFlavorsResponseResultsUnion interface {
	implementsCloudV3GPUBaremetalListFlavorsResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV3GPUBaremetalListFlavorsResponseResultsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithoutPrice{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithPrices{}),
		},
	)
}

type CloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithoutPrice struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required,nullable"`
	// Number of available instances of given flavor
	Capacity int64 `json:"capacity,required"`
	// If the flavor is disabled, new resources cannot be created using this flavor.
	Disabled bool `json:"disabled,required"`
	// Additional bare metal hardware description
	HardwareDescription FlavorHardwareDescriptionBaremetal `json:"hardware_description,required"`
	// Additional bare metal hardware properties
	HardwareProperties FlavorHardwareProperties `json:"hardware_properties,required"`
	// Flavor name
	Name string                                                                                    `json:"name,required"`
	JSON cloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithoutPriceJSON `json:"-"`
}

// cloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithoutPriceJSON
// contains the JSON metadata for the struct
// [CloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithoutPrice]
type cloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithoutPriceJSON struct {
	Architecture        apijson.Field
	Capacity            apijson.Field
	Disabled            apijson.Field
	HardwareDescription apijson.Field
	HardwareProperties  apijson.Field
	Name                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithoutPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithoutPriceJSON) RawJSON() string {
	return r.raw
}

func (r CloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithoutPrice) implementsCloudV3GPUBaremetalListFlavorsResponseResult() {
}

type CloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithPrices struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required,nullable"`
	// Number of available instances of given flavor
	Capacity int64 `json:"capacity,required"`
	// If the flavor is disabled, new resources cannot be created using this flavor.
	Disabled bool `json:"disabled,required"`
	// Additional virtual hardware description
	HardwareDescription FlavorHardwareDescriptionBaremetal `json:"hardware_description,required"`
	// Additional bare metal hardware properties
	HardwareProperties FlavorHardwareProperties `json:"hardware_properties,required"`
	// Flavor name
	Name string `json:"name,required"`
	// Flavor price
	Price FlavorPrice                                                                             `json:"price,required"`
	JSON  cloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithPricesJSON `json:"-"`
}

// cloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithPricesJSON
// contains the JSON metadata for the struct
// [CloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithPrices]
type cloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithPricesJSON struct {
	Architecture        apijson.Field
	Capacity            apijson.Field
	Disabled            apijson.Field
	HardwareDescription apijson.Field
	HardwareProperties  apijson.Field
	Name                apijson.Field
	Price               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithPrices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithPricesJSON) RawJSON() string {
	return r.raw
}

func (r CloudV3GPUBaremetalListFlavorsResponseResultsGPUBaremetalFlavorSerializerWithPrices) implementsCloudV3GPUBaremetalListFlavorsResponseResult() {
}

type CloudV3GPUBaremetalListFlavorsParams struct {
	// Flag for filtering disabled flavors in the region.
	HideDisabled param.Field[bool] `query:"hide_disabled"`
	// Set to true if the response should include flavor prices.
	IncludePrices param.Field[bool] `query:"include_prices"`
}

// URLQuery serializes [CloudV3GPUBaremetalListFlavorsParams]'s query parameters as
// `url.Values`.
func (r CloudV3GPUBaremetalListFlavorsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
