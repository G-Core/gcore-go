// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1PricingAIService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1PricingAIService] method instead.
type CloudV1PricingAIService struct {
	Options  []option.RequestOption
	Clusters *CloudV1PricingAIClusterService
}

// NewCloudV1PricingAIService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1PricingAIService(opts ...option.RequestOption) (r *CloudV1PricingAIService) {
	r = &CloudV1PricingAIService{}
	r.Options = opts
	r.Clusters = NewCloudV1PricingAIClusterService(opts...)
	return
}
