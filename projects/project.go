// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package projects

import (
	"github.com/stainless-sdks/gcore-go/option"
)

// ProjectService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options []option.RequestOption
	V1      V1Service
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r ProjectService) {
	r = ProjectService{}
	r.Options = opts
	r.V1 = NewV1Service(opts...)
	return
}
