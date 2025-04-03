// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudPublicV1Service contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudPublicV1Service] method instead.
type CloudPublicV1Service struct {
	Options  []option.RequestOption
	Ipranges *CloudPublicV1IprangeService
}

// NewCloudPublicV1Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudPublicV1Service(opts ...option.RequestOption) (r *CloudPublicV1Service) {
	r = &CloudPublicV1Service{}
	r.Options = opts
	r.Ipranges = NewCloudPublicV1IprangeService(opts...)
	return
}
