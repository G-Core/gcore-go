// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"github.com/G-Core/gcore-go/option"
)

// PricingBaremetalService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPricingBaremetalService] method instead.
type PricingBaremetalService struct {
	Options []option.RequestOption
	Servers PricingBaremetalServerService
	Flavors PricingBaremetalFlavorService
}

// NewPricingBaremetalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPricingBaremetalService(opts ...option.RequestOption) (r PricingBaremetalService) {
	r = PricingBaremetalService{}
	r.Options = opts
	r.Servers = NewPricingBaremetalServerService(opts...)
	r.Flavors = NewPricingBaremetalFlavorService(opts...)
	return
}
