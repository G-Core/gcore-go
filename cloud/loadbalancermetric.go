// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// LoadBalancerMetricService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoadBalancerMetricService] method instead.
type LoadBalancerMetricService struct {
	Options []option.RequestOption
}

// NewLoadBalancerMetricService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerMetricService(opts ...option.RequestOption) (r LoadBalancerMetricService) {
	r = LoadBalancerMetricService{}
	r.Options = opts
	return
}

// Get load balancer metrics, including cpu, memory and network
func (r *LoadBalancerMetricService) List(ctx context.Context, loadbalancerID string, params LoadBalancerMetricListParams, opts ...option.RequestOption) (res *LoadBalancerMetricsList, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/metrics", params.ProjectID.Value, params.RegionID.Value, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type LoadBalancerMetricListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Time interval
	TimeInterval int64 `json:"time_interval,required"`
	// Time interval unit
	//
	// Any of "day", "hour".
	TimeUnit InstanceMetricsTimeUnit `json:"time_unit,omitzero,required"`
	paramObj
}

func (r LoadBalancerMetricListParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerMetricListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerMetricListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
