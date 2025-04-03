// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1Service contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1Service] method instead.
type CloudV1Service struct {
	Options  []option.RequestOption
	Projects CloudV1ProjectService
}

// NewCloudV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCloudV1Service(opts ...option.RequestOption) (r CloudV1Service) {
	r = CloudV1Service{}
	r.Options = opts
	r.Projects = NewCloudV1ProjectService(opts...)
	return
}
