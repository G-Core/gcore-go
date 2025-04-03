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

// CloudV2PricingInferenceService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2PricingInferenceService] method instead.
type CloudV2PricingInferenceService struct {
	Options []option.RequestOption
}

// NewCloudV2PricingInferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV2PricingInferenceService(opts ...option.RequestOption) (r *CloudV2PricingInferenceService) {
	r = &CloudV2PricingInferenceService{}
	r.Options = opts
	return
}

// Preview billing price of the inference deployment
func (r *CloudV2PricingInferenceService) PreviewDeploymentPrice(ctx context.Context, body CloudV2PricingInferencePreviewDeploymentPriceParams, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/pricing/inference/deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CloudV2PricingInferencePreviewDeploymentPriceParams struct {
	// List of containers for the inference instance.
	Containers param.Field[[]CloudV2PricingInferencePreviewDeploymentPriceParamsContainer] `json:"containers,required"`
	// Flavor name for the inference instance.
	FlavorName param.Field[string] `json:"flavor_name,required"`
}

func (r CloudV2PricingInferencePreviewDeploymentPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2PricingInferencePreviewDeploymentPriceParamsContainer struct {
	// Region ID for the container
	RegionID param.Field[int64] `json:"region_id,required"`
	// Scale for the container
	Scale param.Field[CloudV2PricingInferencePreviewDeploymentPriceParamsContainersScale] `json:"scale,required"`
}

func (r CloudV2PricingInferencePreviewDeploymentPriceParamsContainer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Scale for the container
type CloudV2PricingInferencePreviewDeploymentPriceParamsContainersScale struct {
	// Maximum number of containers
	Max param.Field[int64] `json:"max,required"`
}

func (r CloudV2PricingInferencePreviewDeploymentPriceParamsContainersScale) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
