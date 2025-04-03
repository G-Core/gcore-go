// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2AIService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2AIService] method instead.
type CloudV2AIService struct {
	Options  []option.RequestOption
	Clusters *CloudV2AIClusterService
}

// NewCloudV2AIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV2AIService(opts ...option.RequestOption) (r *CloudV2AIService) {
	r = &CloudV2AIService{}
	r.Options = opts
	r.Clusters = NewCloudV2AIClusterService(opts...)
	return
}
