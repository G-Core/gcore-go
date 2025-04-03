// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2PricingK8ClusterService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2PricingK8ClusterService] method instead.
type CloudV2PricingK8ClusterService struct {
	Options []option.RequestOption
}

// NewCloudV2PricingK8ClusterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV2PricingK8ClusterService(opts ...option.RequestOption) (r *CloudV2PricingK8ClusterService) {
	r = &CloudV2PricingK8ClusterService{}
	r.Options = opts
	return
}

// Pricing scheme for various resources
type ResourcesPricing struct {
	Instances []ItemPriceResponse `json:"instances,required"`
	Publicips []ItemPriceResponse `json:"publicips,required"`
	// Response with prices per hour and per month
	TotalPrice ItemPriceResponse    `json:"total_price,required"`
	Volumes    []ItemPriceResponse  `json:"volumes,required"`
	JSON       resourcesPricingJSON `json:"-"`
}

// resourcesPricingJSON contains the JSON metadata for the struct
// [ResourcesPricing]
type resourcesPricingJSON struct {
	Instances   apijson.Field
	Publicips   apijson.Field
	TotalPrice  apijson.Field
	Volumes     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourcesPricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourcesPricingJSON) RawJSON() string {
	return r.raw
}
