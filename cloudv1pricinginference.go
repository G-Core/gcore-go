// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1PricingInferenceService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1PricingInferenceService] method instead.
type CloudV1PricingInferenceService struct {
	Options []option.RequestOption
}

// NewCloudV1PricingInferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1PricingInferenceService(opts ...option.RequestOption) (r *CloudV1PricingInferenceService) {
	r = &CloudV1PricingInferenceService{}
	r.Options = opts
	return
}

// Deprecated
//
// Preview billing price of the inference deployment based on the provided
// specifications.
func (r *CloudV1PricingInferenceService) PreviewDeploymentPrice(ctx context.Context, body CloudV1PricingInferencePreviewDeploymentPriceParams, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/pricing/inference/deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CloudV1PricingInferencePreviewDeploymentPriceParams struct {
	// List of containers for the inference instance.
	Containers param.Field[[]CloudV1PricingInferencePreviewDeploymentPriceParamsContainer] `json:"containers,required"`
	// Flavor ID for the inference instance.
	FlavorID param.Field[string] `json:"flavor_id,required" format:"uuid"`
}

func (r CloudV1PricingInferencePreviewDeploymentPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1PricingInferencePreviewDeploymentPriceParamsContainer struct {
	// Region ID for the container
	RegionID param.Field[int64] `json:"region_id,required"`
	// Scale for the container
	Scale param.Field[CloudV1PricingInferencePreviewDeploymentPriceParamsContainersScale] `json:"scale,required"`
}

func (r CloudV1PricingInferencePreviewDeploymentPriceParamsContainer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Scale for the container
type CloudV1PricingInferencePreviewDeploymentPriceParamsContainersScale struct {
	// Maximum number of containers
	Max param.Field[int64] `json:"max,required"`
}

func (r CloudV1PricingInferencePreviewDeploymentPriceParamsContainersScale) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
