// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fastedge

import (
	"github.com/G-Core/gcore-go/option"
)

// AppLogService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppLogService] method instead.
type AppLogService struct {
	Options []option.RequestOption
}

// NewAppLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppLogService(opts ...option.RequestOption) (r AppLogService) {
	r = AppLogService{}
	r.Options = opts
	return
}
