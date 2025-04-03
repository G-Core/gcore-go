// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1PricingFaaService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1PricingFaaService] method instead.
type CloudV1PricingFaaService struct {
	Options []option.RequestOption
}

// NewCloudV1PricingFaaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1PricingFaaService(opts ...option.RequestOption) (r *CloudV1PricingFaaService) {
	r = &CloudV1PricingFaaService{}
	r.Options = opts
	return
}

// Preview billing price of the function
func (r *CloudV1PricingFaaService) PreviewFunctionPrice(ctx context.Context, projectID int64, regionID int64, body CloudV1PricingFaaPreviewFunctionPriceParams, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/pricing/%v/%v/faas/functions", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CloudV1PricingFaaPreviewFunctionPriceParams struct {
	// Autoscaling configuration for the function. Keys must be 'min_instances' or
	// 'max_instances', and values must be integers between 0 and 50.
	Autoscaling param.Field[FaasAutoscalingParam] `json:"autoscaling,required"`
	// The name of the flavor associated with the function.
	Flavor param.Field[string] `json:"flavor,required"`
}

func (r CloudV1PricingFaaPreviewFunctionPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
