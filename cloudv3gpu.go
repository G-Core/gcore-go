// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV3GPUService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV3GPUService] method instead.
type CloudV3GPUService struct {
	Options   []option.RequestOption
	Baremetal *CloudV3GPUBaremetalService
	Virtual   *CloudV3GPUVirtualService
}

// NewCloudV3GPUService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV3GPUService(opts ...option.RequestOption) (r *CloudV3GPUService) {
	r = &CloudV3GPUService{}
	r.Options = opts
	r.Baremetal = NewCloudV3GPUBaremetalService(opts...)
	r.Virtual = NewCloudV3GPUVirtualService(opts...)
	return
}
