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

// Create Load Balancer Pool Health Monitor
func (r *LoadBalancerPoolHealthMonitorService) New(ctx context.Context, poolID string, params LoadBalancerPoolHealthMonitorNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s/healthmonitor", params.ProjectID.Value, params.RegionID.Value, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete load balancer pool health monitor
func (r *LoadBalancerPoolHealthMonitorService) Delete(ctx context.Context, poolID string, body LoadBalancerPoolHealthMonitorDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
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
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bpool_id%7D%2Fhealthmonitor/post/parameters/0/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}/{pool_id}/healthmonitor'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bpool_id%7D%2Fhealthmonitor/post/parameters/1/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}/{pool_id}/healthmonitor'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateLbHealthMonitorSerializer/properties/delay'
	// "$.components.schemas.CreateLbHealthMonitorSerializer.properties.delay"
	Delay int64 `json:"delay,required"`
	// '#/components/schemas/CreateLbHealthMonitorSerializer/properties/max_retries'
	// "$.components.schemas.CreateLbHealthMonitorSerializer.properties.max_retries"
	MaxRetries int64 `json:"max_retries,required"`
	// '#/components/schemas/CreateLbHealthMonitorSerializer/properties/timeout'
	// "$.components.schemas.CreateLbHealthMonitorSerializer.properties.timeout"
	Timeout int64 `json:"timeout,required"`
	// '#/components/schemas/CreateLbHealthMonitorSerializer/properties/type'
	// "$.components.schemas.CreateLbHealthMonitorSerializer.properties.type"
	//
	// Any of "HTTP", "HTTPS", "K8S", "PING", "TCP", "TLS-HELLO", "UDP-CONNECT".
	Type HealthMonitorType `json:"type,omitzero,required"`
	// '#/components/schemas/CreateLbHealthMonitorSerializer/properties/expected_codes/anyOf/0'
	// "$.components.schemas.CreateLbHealthMonitorSerializer.properties.expected_codes.anyOf[0]"
	ExpectedCodes param.Opt[string] `json:"expected_codes,omitzero"`
	// '#/components/schemas/CreateLbHealthMonitorSerializer/properties/max_retries_down/anyOf/0'
	// "$.components.schemas.CreateLbHealthMonitorSerializer.properties.max_retries_down.anyOf[0]"
	MaxRetriesDown param.Opt[int64] `json:"max_retries_down,omitzero"`
	// '#/components/schemas/CreateLbHealthMonitorSerializer/properties/url_path/anyOf/0'
	// "$.components.schemas.CreateLbHealthMonitorSerializer.properties.url_path.anyOf[0]"
	URLPath param.Opt[string] `json:"url_path,omitzero"`
	// '#/components/schemas/CreateLbHealthMonitorSerializer/properties/http_method/anyOf/0'
	// "$.components.schemas.CreateLbHealthMonitorSerializer.properties.http_method.anyOf[0]"
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod HTTPMethod `json:"http_method,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerPoolHealthMonitorNewParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r LoadBalancerPoolHealthMonitorNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolHealthMonitorNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type LoadBalancerPoolHealthMonitorDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bpool_id%7D%2Fhealthmonitor/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}/{pool_id}/healthmonitor']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bpool_id%7D%2Fhealthmonitor/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}/{pool_id}/healthmonitor']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerPoolHealthMonitorDeleteParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
