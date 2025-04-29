// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
)

// LoadBalancerStatusService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoadBalancerStatusService] method instead.
type LoadBalancerStatusService struct {
	Options []option.RequestOption
}

// NewLoadBalancerStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerStatusService(opts ...option.RequestOption) (r LoadBalancerStatusService) {
	r = LoadBalancerStatusService{}
	r.Options = opts
	return
}

// List load balancers statuses
func (r *LoadBalancerStatusService) List(ctx context.Context, query LoadBalancerStatusListParams, opts ...option.RequestOption) (res *LoadBalancerStatusList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/status", query.ProjectID.Value, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get load balancer status
func (r *LoadBalancerStatusService) Get(ctx context.Context, loadbalancerID string, query LoadBalancerStatusGetParams, opts ...option.RequestOption) (res *LoadBalancerStatus, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/status", query.ProjectID.Value, query.RegionID.Value, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LoadBalancerStatusListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Fstatus/get/parameters/0/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/status'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2Fstatus/get/parameters/1/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/status'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerStatusListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type LoadBalancerStatusGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bloadbalancer_id%7D%2Fstatus/get/parameters/0/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/{loadbalancer_id}/status'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bloadbalancer_id%7D%2Fstatus/get/parameters/1/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/{loadbalancer_id}/status'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerStatusGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
