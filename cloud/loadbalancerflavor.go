// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
)

// LoadBalancerFlavorService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoadBalancerFlavorService] method instead.
type LoadBalancerFlavorService struct {
	Options []option.RequestOption
}

// NewLoadBalancerFlavorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerFlavorService(opts ...option.RequestOption) (r LoadBalancerFlavorService) {
	r = LoadBalancerFlavorService{}
	r.Options = opts
	return
}

// Retrieve a list of load balancer flavors. When the include_prices query
// parameter is specified, the list shows prices. A client in trial mode gets all
// price values as 0. If you get Pricing Error contact the support
func (r *LoadBalancerFlavorService) List(ctx context.Context, params LoadBalancerFlavorListParams, opts ...option.RequestOption) (res *LbFlavorList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbflavors/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type LoadBalancerFlavorListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Flbflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/lbflavors/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flbflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/lbflavors/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flbflavors%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/lbflavors/{project_id}/{region_id}'].get.parameters[2]"
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerFlavorListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LoadBalancerFlavorListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
