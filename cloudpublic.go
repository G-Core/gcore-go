// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudPublicService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudPublicService] method instead.
type CloudPublicService struct {
	Options []option.RequestOption
	V1      *CloudPublicV1Service
}

// NewCloudPublicService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudPublicService(opts ...option.RequestOption) (r *CloudPublicService) {
	r = &CloudPublicService{}
	r.Options = opts
	r.V1 = NewCloudPublicV1Service(opts...)
	return
}
