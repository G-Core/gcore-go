// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV3Service contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV3Service] method instead.
type CloudV3Service struct {
	Options   []option.RequestOption
	GPU       *CloudV3GPUService
	Inference *CloudV3InferenceService
}

// NewCloudV3Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCloudV3Service(opts ...option.RequestOption) (r *CloudV3Service) {
	r = &CloudV3Service{}
	r.Options = opts
	r.GPU = NewCloudV3GPUService(opts...)
	r.Inference = NewCloudV3InferenceService(opts...)
	return
}
