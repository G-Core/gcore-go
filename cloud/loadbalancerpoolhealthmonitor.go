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

// LoadBalancerPoolHealthMonitorService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoadBalancerPoolHealthMonitorService] method instead.
type LoadBalancerPoolHealthMonitorService struct {
	Options []option.RequestOption
}

// NewLoadBalancerPoolHealthMonitorService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLoadBalancerPoolHealthMonitorService(opts ...option.RequestOption) (r LoadBalancerPoolHealthMonitorService) {
	r = LoadBalancerPoolHealthMonitorService{}
	r.Options = opts
	return
}

// Creates a health monitor for a load balancer pool to automatically check the
// health status of pool members. The health monitor performs periodic checks on
// pool members and removes unhealthy members from rotation, ensuring only healthy
// servers receive traffic.
func (r *LoadBalancerPoolHealthMonitorService) New(ctx context.Context, poolID string, params LoadBalancerPoolHealthMonitorNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s/healthmonitor", params.ProjectID.Value, params.RegionID.Value, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Removes the health monitor from a load balancer pool. After deletion, the pool
// will no longer perform automatic health checks on its members, and all members
// will remain in rotation regardless of their actual health status.
func (r *LoadBalancerPoolHealthMonitorService) Delete(ctx context.Context, poolID string, body LoadBalancerPoolHealthMonitorDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s/healthmonitor", body.ProjectID.Value, body.RegionID.Value, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type LoadBalancerPoolHealthMonitorNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// The time, in seconds, between sending probes to members
	Delay int64 `json:"delay" api:"required"`
	// Number of successes before the member is switched to ONLINE state
	MaxRetries int64 `json:"max_retries" api:"required"`
	// The maximum time to connect. Must be less than the delay value
	Timeout int64 `json:"timeout" api:"required"`
	// Health monitor type. Once health monitor is created, cannot be changed.
	//
	// Any of "HTTP", "HTTPS", "K8S", "PING", "TCP", "TLS-HELLO", "UDP-CONNECT".
	Type LbHealthMonitorType `json:"type,omitzero" api:"required"`
	// Expected HTTP response codes. Can be a single code or a range of codes. Can only
	// be used together with `HTTP` or `HTTPS` health monitor type. For example,
	// 200,202,300-302,401,403,404,500-504. If not specified, the default is 200.
	ExpectedCodes param.Opt[string] `json:"expected_codes,omitzero"`
	// URL Path. Defaults to '/'. Can only be used together with `HTTP` or `HTTPS`
	// health monitor type.
	URLPath param.Opt[string] `json:"url_path,omitzero"`
	// Administrative state of the resource. When set to true, the resource is enabled
	// and operational. When set to false, the resource is disabled and will not
	// process traffic. Defaults to true.
	AdminStateUp param.Opt[bool] `json:"admin_state_up,omitzero"`
	// Number of failures before the member is switched to ERROR state.
	MaxRetriesDown param.Opt[int64] `json:"max_retries_down,omitzero"`
	// HTTP method. Can only be used together with `HTTP` or `HTTPS` health monitor
	// type.
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod HTTPMethod `json:"http_method,omitzero"`
	paramObj
}

func (r LoadBalancerPoolHealthMonitorNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolHealthMonitorNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerPoolHealthMonitorNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolHealthMonitorDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}
