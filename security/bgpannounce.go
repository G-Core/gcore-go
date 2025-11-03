// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security

import (
	"github.com/G-Core/gcore-go/option"
)

// BgpAnnounceService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBgpAnnounceService] method instead.
type BgpAnnounceService struct {
	Options []option.RequestOption
}

// NewBgpAnnounceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBgpAnnounceService(opts ...option.RequestOption) (r BgpAnnounceService) {
	r = BgpAnnounceService{}
	r.Options = opts
	return
}
