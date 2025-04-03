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

// CloudV3GPUVirtualService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV3GPUVirtualService] method instead.
type CloudV3GPUVirtualService struct {
	Options []option.RequestOption
	Images  *CloudV3GPUVirtualImageService
}

// NewCloudV3GPUVirtualService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV3GPUVirtualService(opts ...option.RequestOption) (r *CloudV3GPUVirtualService) {
	r = &CloudV3GPUVirtualService{}
	r.Options = opts
	r.Images = NewCloudV3GPUVirtualImageService(opts...)
	return
}

// List virtual GPU flavors
func (r *CloudV3GPUVirtualService) ListFlavors(ctx context.Context, projectID int64, regionID int64, query CloudV3GPUVirtualListFlavorsParams, opts ...option.RequestOption) (res *CloudV3GPUVirtualListFlavorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/flavors", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type FlavorHardwareDescriptionVirtual struct {
	// Human-readable GPU description
	GPU string `json:"gpu,required,nullable"`
	// RAM size in MiB
	Ram int64 `json:"ram,required,nullable"`
	// Virtual CPU count
	Vcpus int64                                `json:"vcpus,required,nullable"`
	JSON  flavorHardwareDescriptionVirtualJSON `json:"-"`
}

// flavorHardwareDescriptionVirtualJSON contains the JSON metadata for the struct
// [FlavorHardwareDescriptionVirtual]
type flavorHardwareDescriptionVirtualJSON struct {
	GPU         apijson.Field
	Ram         apijson.Field
	Vcpus       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FlavorHardwareDescriptionVirtual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r flavorHardwareDescriptionVirtualJSON) RawJSON() string {
	return r.raw
}

type CloudV3GPUVirtualListFlavorsResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV3GPUVirtualListFlavorsResponseResult `json:"results,required"`
	JSON    cloudV3GPUVirtualListFlavorsResponseJSON     `json:"-"`
}

// cloudV3GPUVirtualListFlavorsResponseJSON contains the JSON metadata for the
// struct [CloudV3GPUVirtualListFlavorsResponse]
type cloudV3GPUVirtualListFlavorsResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV3GPUVirtualListFlavorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3GPUVirtualListFlavorsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV3GPUVirtualListFlavorsResponseResult struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required,nullable"`
	// Number of available instances of given flavor
	Capacity int64 `json:"capacity,required"`
	// If the flavor is disabled, new resources cannot be created using this flavor.
	Disabled bool `json:"disabled,required"`
	// Additional virtual hardware description
	HardwareDescription FlavorHardwareDescriptionVirtual `json:"hardware_description,required"`
	// Additional virtual hardware properties
	HardwareProperties FlavorHardwareProperties `json:"hardware_properties,required"`
	// Flavor name
	Name string `json:"name,required"`
	// Flavor price.
	Price FlavorPrice                                    `json:"price"`
	JSON  cloudV3GPUVirtualListFlavorsResponseResultJSON `json:"-"`
	union CloudV3GPUVirtualListFlavorsResponseResultsUnion
}

// cloudV3GPUVirtualListFlavorsResponseResultJSON contains the JSON metadata for
// the struct [CloudV3GPUVirtualListFlavorsResponseResult]
type cloudV3GPUVirtualListFlavorsResponseResultJSON struct {
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

func (r cloudV3GPUVirtualListFlavorsResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *CloudV3GPUVirtualListFlavorsResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = CloudV3GPUVirtualListFlavorsResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CloudV3GPUVirtualListFlavorsResponseResultsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithoutPrice],
// [CloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithPrices].
func (r CloudV3GPUVirtualListFlavorsResponseResult) AsUnion() CloudV3GPUVirtualListFlavorsResponseResultsUnion {
	return r.union
}

// Union satisfied by
// [CloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithoutPrice]
// or
// [CloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithPrices].
type CloudV3GPUVirtualListFlavorsResponseResultsUnion interface {
	implementsCloudV3GPUVirtualListFlavorsResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV3GPUVirtualListFlavorsResponseResultsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithoutPrice{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithPrices{}),
		},
	)
}

type CloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithoutPrice struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required,nullable"`
	// Number of available instances of given flavor
	Capacity int64 `json:"capacity,required"`
	// If the flavor is disabled, new resources cannot be created using this flavor.
	Disabled bool `json:"disabled,required"`
	// Additional virtual hardware description
	HardwareDescription FlavorHardwareDescriptionVirtual `json:"hardware_description,required"`
	// Additional virtual hardware properties
	HardwareProperties FlavorHardwareProperties `json:"hardware_properties,required"`
	// Flavor name
	Name string                                                                                `json:"name,required"`
	JSON cloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithoutPriceJSON `json:"-"`
}

// cloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithoutPriceJSON
// contains the JSON metadata for the struct
// [CloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithoutPrice]
type cloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithoutPriceJSON struct {
	Architecture        apijson.Field
	Capacity            apijson.Field
	Disabled            apijson.Field
	HardwareDescription apijson.Field
	HardwareProperties  apijson.Field
	Name                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithoutPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithoutPriceJSON) RawJSON() string {
	return r.raw
}

func (r CloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithoutPrice) implementsCloudV3GPUVirtualListFlavorsResponseResult() {
}

type CloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithPrices struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required,nullable"`
	// Number of available instances of given flavor
	Capacity int64 `json:"capacity,required"`
	// If the flavor is disabled, new resources cannot be created using this flavor.
	Disabled bool `json:"disabled,required"`
	// Additional virtual hardware description
	HardwareDescription FlavorHardwareDescriptionVirtual `json:"hardware_description,required"`
	// Additional virtual hardware properties
	HardwareProperties FlavorHardwareProperties `json:"hardware_properties,required"`
	// Flavor name
	Name string `json:"name,required"`
	// Flavor price.
	Price FlavorPrice                                                                         `json:"price,required"`
	JSON  cloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithPricesJSON `json:"-"`
}

// cloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithPricesJSON
// contains the JSON metadata for the struct
// [CloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithPrices]
type cloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithPricesJSON struct {
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

func (r *CloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithPrices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithPricesJSON) RawJSON() string {
	return r.raw
}

func (r CloudV3GPUVirtualListFlavorsResponseResultsGPUVirtualFlavorSerializerWithPrices) implementsCloudV3GPUVirtualListFlavorsResponseResult() {
}

type CloudV3GPUVirtualListFlavorsParams struct {
	// Flag for filtering disabled flavors in the region.
	HideDisabled param.Field[bool] `query:"hide_disabled"`
	// Set to true if the response should include flavor prices.
	IncludePrices param.Field[bool] `query:"include_prices"`
}

// URLQuery serializes [CloudV3GPUVirtualListFlavorsParams]'s query parameters as
// `url.Values`.
func (r CloudV3GPUVirtualListFlavorsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
