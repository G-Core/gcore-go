// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"github.com/G-Core/gcore-go/option"
)

// DomainAPIDiscoveryService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainAPIDiscoveryService] method instead.
type DomainAPIDiscoveryService struct {
	Options     []option.RequestOption
	ScanResults DomainAPIDiscoveryScanResultService
	OpenAPI     DomainAPIDiscoveryOpenAPIService
	Settings    DomainAPIDiscoverySettingService
}

// NewDomainAPIDiscoveryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDomainAPIDiscoveryService(opts ...option.RequestOption) (r DomainAPIDiscoveryService) {
	r = DomainAPIDiscoveryService{}
	r.Options = opts
	r.ScanResults = NewDomainAPIDiscoveryScanResultService(opts...)
	r.OpenAPI = NewDomainAPIDiscoveryOpenAPIService(opts...)
	r.Settings = NewDomainAPIDiscoverySettingService(opts...)
	return
}
