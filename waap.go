// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"github.com/stainless-sdks/gcore-go/option"
)

// WaapService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapService] method instead.
type WaapService struct {
	Options []option.RequestOption
	V1      *WaapV1Service
}

// NewWaapService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWaapService(opts ...option.RequestOption) (r *WaapService) {
	r = &WaapService{}
	r.Options = opts
	r.V1 = NewWaapV1Service(opts...)
	return
}
