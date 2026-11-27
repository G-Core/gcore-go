// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// GPU pricing lists the public list prices of bare metal and virtual GPUs in EUR
// and USD.
//
// PricingGPUService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPricingGPUService] method instead.
type PricingGPUService struct {
	Options []option.RequestOption
}

// NewPricingGPUService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPricingGPUService(opts ...option.RequestOption) (r PricingGPUService) {
	r = PricingGPUService{}
	r.Options = opts
	return
}

// Returns list prices for GPUs, bare metal and virtual, in EUR and USD.
//
// The response has no pagination. The catalog is small. The list sorts by price
// per hour, from low to high.
func (r *PricingGPUService) List(ctx context.Context, opts ...option.RequestOption) (res *GPUPriceCatalog, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cloud/public/v1/pricing/gpu"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type GPUCurrency string

const (
	GPUCurrencyEur GPUCurrency = "EUR"
	GPUCurrencyUsd GPUCurrency = "USD"
)

type GPUDeploymentType string

const (
	GPUDeploymentTypeBaremetal GPUDeploymentType = "baremetal"
	GPUDeploymentTypeVirtual   GPUDeploymentType = "virtual"
)

type GPUPrice struct {
	// ISO 4217 currency code.
	//
	// Any of "EUR", "USD".
	Currency GPUCurrency `json:"currency" api:"required"`
	// Baremetal or virtual.
	//
	// Any of "baremetal", "virtual".
	DeploymentType GPUDeploymentType `json:"deployment_type" api:"required"`
	// Number of GPU cards this price covers.
	GPUCount int64 `json:"gpu_count" api:"required"`
	// GPU model.
	GPUModel string `json:"gpu_model" api:"required"`
	// On-demand price per GPU per hour.
	PricePerGPUHour string `json:"price_per_gpu_hour" api:"required"`
	// On-demand price per node/instance per hour.
	PricePerHour string `json:"price_per_hour" api:"required"`
	// Pricing-specific detail.
	PricingDetail GPUPriceDetail `json:"pricing_detail" api:"required"`
	// Region this price applies to.
	RegionName string `json:"region_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency        respjson.Field
		DeploymentType  respjson.Field
		GPUCount        respjson.Field
		GPUModel        respjson.Field
		PricePerGPUHour respjson.Field
		PricePerHour    respjson.Field
		PricingDetail   respjson.Field
		RegionName      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUPrice) RawJSON() string { return r.JSON.raw }
func (r *GPUPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUPriceCatalog struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// When this catalog snapshot was built.
	GeneratedAt time.Time `json:"generated_at" api:"required" format:"date-time"`
	// Objects
	Results []GPUPrice `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		GeneratedAt respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUPriceCatalog) RawJSON() string { return r.JSON.raw }
func (r *GPUPriceCatalog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUPriceDetail struct {
	// Flavor names that share this SKU and price.
	FlavorNames []string `json:"flavor_names" api:"required"`
	// SKU id for this price row.
	SKU string `json:"sku" api:"required"`
	// When this row was last refreshed.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FlavorNames respjson.Field
		SKU         respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUPriceDetail) RawJSON() string { return r.JSON.raw }
func (r *GPUPriceDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
