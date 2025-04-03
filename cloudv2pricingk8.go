// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2PricingK8Service contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2PricingK8Service] method instead.
type CloudV2PricingK8Service struct {
	Options  []option.RequestOption
	Clusters *CloudV2PricingK8ClusterService
}

// NewCloudV2PricingK8Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV2PricingK8Service(opts ...option.RequestOption) (r *CloudV2PricingK8Service) {
	r = &CloudV2PricingK8Service{}
	r.Options = opts
	r.Clusters = NewCloudV2PricingK8ClusterService(opts...)
	return
}
