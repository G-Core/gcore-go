// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"github.com/G-Core/gcore-go/option"
)

// GPUVirtualService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUVirtualService] method instead.
type GPUVirtualService struct {
	Options  []option.RequestOption
	Clusters GPUVirtualClusterService
}

// NewGPUVirtualService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGPUVirtualService(opts ...option.RequestOption) (r GPUVirtualService) {
	r = GPUVirtualService{}
	r.Options = opts
	r.Clusters = NewGPUVirtualClusterService(opts...)
	return
}
