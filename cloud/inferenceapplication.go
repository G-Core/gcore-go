// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"github.com/G-Core/gcore-go/option"
)

// InferenceApplicationService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceApplicationService] method instead.
type InferenceApplicationService struct {
	Options     []option.RequestOption
	Deployments InferenceApplicationDeploymentService
	Templates   InferenceApplicationTemplateService
}

// NewInferenceApplicationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInferenceApplicationService(opts ...option.RequestOption) (r InferenceApplicationService) {
	r = InferenceApplicationService{}
	r.Options = opts
	r.Deployments = NewInferenceApplicationDeploymentService(opts...)
	r.Templates = NewInferenceApplicationTemplateService(opts...)
	return
}
