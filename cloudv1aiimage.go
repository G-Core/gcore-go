// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1AIImageService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1AIImageService] method instead.
type CloudV1AIImageService struct {
	Options []option.RequestOption
	GPU     *CloudV1AIImageGPUService
}

// NewCloudV1AIImageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1AIImageService(opts ...option.RequestOption) (r *CloudV1AIImageService) {
	r = &CloudV1AIImageService{}
	r.Options = opts
	r.GPU = NewCloudV1AIImageGPUService(opts...)
	return
}
