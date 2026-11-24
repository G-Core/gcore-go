// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"github.com/G-Core/gcore-go/option"
)

// Region-pathed bare metal servers and flavors. This client is temporary, and a
// future release will remove it. Use `cloud.baremetal` in new code.
//
// BaremetalRegionalService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBaremetalRegionalService] method instead.
type BaremetalRegionalService struct {
	Options []option.RequestOption
	Flavors BaremetalRegionalFlavorService
	// Bare metal servers are dedicated physical machines with direct hardware access,
	// supporting provisioning, rebuilding, and network configuration within a cloud
	// region.
	Servers BaremetalRegionalServerService
}

// NewBaremetalRegionalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBaremetalRegionalService(opts ...option.RequestOption) (r BaremetalRegionalService) {
	r = BaremetalRegionalService{}
	r.Options = opts
	r.Flavors = NewBaremetalRegionalFlavorService(opts...)
	r.Servers = NewBaremetalRegionalServerService(opts...)
	return
}
