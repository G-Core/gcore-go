// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1DbaaService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1DbaaService] method instead.
type CloudV1DbaaService struct {
	Options  []option.RequestOption
	Postgres *CloudV1DbaaPostgreService
	Status   *CloudV1DbaaStatusService
}

// NewCloudV1DbaaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1DbaaService(opts ...option.RequestOption) (r *CloudV1DbaaService) {
	r = &CloudV1DbaaService{}
	r.Options = opts
	r.Postgres = NewCloudV1DbaaPostgreService(opts...)
	r.Status = NewCloudV1DbaaStatusService(opts...)
	return
}
