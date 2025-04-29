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

// LoadBalancerPoolService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoadBalancerPoolService] method instead.
type LoadBalancerPoolService struct {
	Options        []option.RequestOption
	HealthMonitors LoadBalancerPoolHealthMonitorService
	Members        LoadBalancerPoolMemberService
}

// NewLoadBalancerPoolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerPoolService(opts ...option.RequestOption) (r LoadBalancerPoolService) {
	r = LoadBalancerPoolService{}
	r.Options = opts
	r.HealthMonitors = NewLoadBalancerPoolHealthMonitorService(opts...)
	r.Members = NewLoadBalancerPoolMemberService(opts...)
	return
}

// Create load balancer pool
func (r *LoadBalancerPoolService) New(ctx context.Context, params LoadBalancerPoolNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Changes provided here will overwrite existing load balancer pool settings.
// Undefined fields will be kept as is. Complex objects need to be specified fully,
// they will be overwritten.
func (r *LoadBalancerPoolService) Update(ctx context.Context, poolID string, params LoadBalancerPoolUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List load balancer pools
func (r *LoadBalancerPoolService) List(ctx context.Context, params LoadBalancerPoolListParams, opts ...option.RequestOption) (res *DetailedLbPoolList, err error) {
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
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete load balancer pool
func (r *LoadBalancerPoolService) Delete(ctx context.Context, poolID string, body LoadBalancerPoolDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
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
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get load balancer pool
func (r *LoadBalancerPoolService) Get(ctx context.Context, poolID string, query LoadBalancerPoolGetParams, opts ...option.RequestOption) (res *DetailedLbPool, err error) {
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
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LoadBalancerPoolNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/lb_algorithm'
	// "$.components.schemas.CreateLbPoolSerializer.properties.lb_algorithm"
	//
	// Any of "LEAST_CONNECTIONS", "ROUND_ROBIN", "SOURCE_IP".
	LbAlgorithm LbAlgorithm `json:"lb_algorithm,omitzero,required"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/name'
	// "$.components.schemas.CreateLbPoolSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/protocol'
	// "$.components.schemas.CreateLbPoolSerializer.properties.protocol"
	//
	// Any of "HTTP", "HTTPS", "PROXY", "PROXYV2", "TCP", "UDP".
	Protocol LbPoolProtocol `json:"protocol,omitzero,required"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/ca_secret_id/anyOf/0'
	// "$.components.schemas.CreateLbPoolSerializer.properties.ca_secret_id.anyOf[0]"
	CaSecretID param.Opt[string] `json:"ca_secret_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/crl_secret_id/anyOf/0'
	// "$.components.schemas.CreateLbPoolSerializer.properties.crl_secret_id.anyOf[0]"
	CrlSecretID param.Opt[string] `json:"crl_secret_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/listener_id/anyOf/0'
	// "$.components.schemas.CreateLbPoolSerializer.properties.listener_id.anyOf[0]"
	ListenerID param.Opt[string] `json:"listener_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/loadbalancer_id/anyOf/0'
	// "$.components.schemas.CreateLbPoolSerializer.properties.loadbalancer_id.anyOf[0]"
	LoadbalancerID param.Opt[string] `json:"loadbalancer_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/secret_id/anyOf/0'
	// "$.components.schemas.CreateLbPoolSerializer.properties.secret_id.anyOf[0]"
	SecretID param.Opt[string] `json:"secret_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/timeout_client_data/anyOf/0'
	// "$.components.schemas.CreateLbPoolSerializer.properties.timeout_client_data.anyOf[0]"
	TimeoutClientData param.Opt[int64] `json:"timeout_client_data,omitzero"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/timeout_member_connect/anyOf/0'
	// "$.components.schemas.CreateLbPoolSerializer.properties.timeout_member_connect.anyOf[0]"
	TimeoutMemberConnect param.Opt[int64] `json:"timeout_member_connect,omitzero"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/timeout_member_data/anyOf/0'
	// "$.components.schemas.CreateLbPoolSerializer.properties.timeout_member_data.anyOf[0]"
	TimeoutMemberData param.Opt[int64] `json:"timeout_member_data,omitzero"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/healthmonitor/anyOf/0'
	// "$.components.schemas.CreateLbPoolSerializer.properties.healthmonitor.anyOf[0]"
	Healthmonitor LoadBalancerPoolNewParamsHealthmonitor `json:"healthmonitor,omitzero"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/members/anyOf/0'
	// "$.components.schemas.CreateLbPoolSerializer.properties.members.anyOf[0]"
	Members []LoadBalancerPoolNewParamsMember `json:"members,omitzero"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/session_persistence/anyOf/0'
	// "$.components.schemas.CreateLbPoolSerializer.properties.session_persistence.anyOf[0]"
	SessionPersistence LoadBalancerPoolNewParamsSessionPersistence `json:"session_persistence,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerPoolNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LoadBalancerPoolNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateLbPoolSerializer/properties/healthmonitor/anyOf/0'
// "$.components.schemas.CreateLbPoolSerializer.properties.healthmonitor.anyOf[0]"
//
// The properties Delay, MaxRetries, Timeout, Type are required.
type LoadBalancerPoolNewParamsHealthmonitor struct {
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
func (f LoadBalancerPoolNewParamsHealthmonitor) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerPoolNewParamsHealthmonitor) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolNewParamsHealthmonitor
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateLbPoolSerializer/properties/members/anyOf/0/items'
// "$.components.schemas.CreateLbPoolSerializer.properties.members.anyOf[0].items"
//
// The properties Address, ProtocolPort are required.
type LoadBalancerPoolNewParamsMember struct {
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/address'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.address"
	Address string `json:"address,required" format:"ipvanyaddress"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/protocol_port'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.protocol_port"
	ProtocolPort int64 `json:"protocol_port,required"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/admin_state_up/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.admin_state_up.anyOf[0]"
	AdminStateUp param.Opt[bool] `json:"admin_state_up,omitzero"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/instance_id/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.instance_id.anyOf[0]"
	InstanceID param.Opt[string] `json:"instance_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/monitor_address/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.monitor_address.anyOf[0]"
	MonitorAddress param.Opt[string] `json:"monitor_address,omitzero" format:"ipvanyaddress"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/monitor_port/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.monitor_port.anyOf[0]"
	MonitorPort param.Opt[int64] `json:"monitor_port,omitzero"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/subnet_id/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.subnet_id.anyOf[0]"
	SubnetID param.Opt[string] `json:"subnet_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/weight/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.weight.anyOf[0]"
	Weight param.Opt[int64] `json:"weight,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerPoolNewParamsMember) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r LoadBalancerPoolNewParamsMember) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolNewParamsMember
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateLbPoolSerializer/properties/session_persistence/anyOf/0'
// "$.components.schemas.CreateLbPoolSerializer.properties.session_persistence.anyOf[0]"
//
// The property Type is required.
type LoadBalancerPoolNewParamsSessionPersistence struct {
	// '#/components/schemas/MutateLbSessionPersistence/properties/type'
	// "$.components.schemas.MutateLbSessionPersistence.properties.type"
	//
	// Any of "APP_COOKIE", "HTTP_COOKIE", "SOURCE_IP".
	Type SessionPersistenceType `json:"type,omitzero,required"`
	// '#/components/schemas/MutateLbSessionPersistence/properties/cookie_name/anyOf/0'
	// "$.components.schemas.MutateLbSessionPersistence.properties.cookie_name.anyOf[0]"
	CookieName param.Opt[string] `json:"cookie_name,omitzero"`
	// '#/components/schemas/MutateLbSessionPersistence/properties/persistence_granularity/anyOf/0'
	// "$.components.schemas.MutateLbSessionPersistence.properties.persistence_granularity.anyOf[0]"
	PersistenceGranularity param.Opt[string] `json:"persistence_granularity,omitzero"`
	// '#/components/schemas/MutateLbSessionPersistence/properties/persistence_timeout/anyOf/0'
	// "$.components.schemas.MutateLbSessionPersistence.properties.persistence_timeout.anyOf[0]"
	PersistenceTimeout param.Opt[int64] `json:"persistence_timeout,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerPoolNewParamsSessionPersistence) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerPoolNewParamsSessionPersistence) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolNewParamsSessionPersistence
	return param.MarshalObject(r, (*shadow)(&r))
}

type LoadBalancerPoolUpdateParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bpool_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}/{pool_id}'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bpool_id%7D/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}/{pool_id}'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/PatchLbPoolSerializer/properties/ca_secret_id/anyOf/0'
	// "$.components.schemas.PatchLbPoolSerializer.properties.ca_secret_id.anyOf[0]"
	CaSecretID param.Opt[string] `json:"ca_secret_id,omitzero" format:"uuid4"`
	// '#/components/schemas/PatchLbPoolSerializer/properties/crl_secret_id/anyOf/0'
	// "$.components.schemas.PatchLbPoolSerializer.properties.crl_secret_id.anyOf[0]"
	CrlSecretID param.Opt[string] `json:"crl_secret_id,omitzero" format:"uuid4"`
	// '#/components/schemas/PatchLbPoolSerializer/properties/secret_id/anyOf/0'
	// "$.components.schemas.PatchLbPoolSerializer.properties.secret_id.anyOf[0]"
	SecretID param.Opt[string] `json:"secret_id,omitzero" format:"uuid4"`
	// '#/components/schemas/PatchLbPoolSerializer/properties/timeout_client_data/anyOf/0'
	// "$.components.schemas.PatchLbPoolSerializer.properties.timeout_client_data.anyOf[0]"
	TimeoutClientData param.Opt[int64] `json:"timeout_client_data,omitzero"`
	// '#/components/schemas/PatchLbPoolSerializer/properties/timeout_member_connect/anyOf/0'
	// "$.components.schemas.PatchLbPoolSerializer.properties.timeout_member_connect.anyOf[0]"
	TimeoutMemberConnect param.Opt[int64] `json:"timeout_member_connect,omitzero"`
	// '#/components/schemas/PatchLbPoolSerializer/properties/timeout_member_data/anyOf/0'
	// "$.components.schemas.PatchLbPoolSerializer.properties.timeout_member_data.anyOf[0]"
	TimeoutMemberData param.Opt[int64] `json:"timeout_member_data,omitzero"`
	// '#/components/schemas/PatchLbPoolSerializer/properties/name'
	// "$.components.schemas.PatchLbPoolSerializer.properties.name"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/PatchLbPoolSerializer/properties/healthmonitor/anyOf/0'
	// "$.components.schemas.PatchLbPoolSerializer.properties.healthmonitor.anyOf[0]"
	Healthmonitor LoadBalancerPoolUpdateParamsHealthmonitor `json:"healthmonitor,omitzero"`
	// '#/components/schemas/PatchLbPoolSerializer/properties/members/anyOf/0'
	// "$.components.schemas.PatchLbPoolSerializer.properties.members.anyOf[0]"
	Members []LoadBalancerPoolUpdateParamsMember `json:"members,omitzero"`
	// '#/components/schemas/PatchLbPoolSerializer/properties/session_persistence/anyOf/0'
	// "$.components.schemas.PatchLbPoolSerializer.properties.session_persistence.anyOf[0]"
	SessionPersistence LoadBalancerPoolUpdateParamsSessionPersistence `json:"session_persistence,omitzero"`
	// '#/components/schemas/PatchLbPoolSerializer/properties/lb_algorithm'
	// "$.components.schemas.PatchLbPoolSerializer.properties.lb_algorithm"
	//
	// Any of "LEAST_CONNECTIONS", "ROUND_ROBIN", "SOURCE_IP".
	LbAlgorithm LbAlgorithm `json:"lb_algorithm,omitzero"`
	// '#/components/schemas/PatchLbPoolSerializer/properties/protocol'
	// "$.components.schemas.PatchLbPoolSerializer.properties.protocol"
	//
	// Any of "HTTP", "HTTPS", "PROXY", "PROXYV2", "TCP", "UDP".
	Protocol LbPoolProtocol `json:"protocol,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerPoolUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LoadBalancerPoolUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/PatchLbPoolSerializer/properties/healthmonitor/anyOf/0'
// "$.components.schemas.PatchLbPoolSerializer.properties.healthmonitor.anyOf[0]"
//
// The properties Delay, MaxRetries, Timeout are required.
type LoadBalancerPoolUpdateParamsHealthmonitor struct {
	// '#/components/schemas/PatchLbHealthMonitorSerializer/properties/delay'
	// "$.components.schemas.PatchLbHealthMonitorSerializer.properties.delay"
	Delay int64 `json:"delay,required"`
	// '#/components/schemas/PatchLbHealthMonitorSerializer/properties/max_retries'
	// "$.components.schemas.PatchLbHealthMonitorSerializer.properties.max_retries"
	MaxRetries int64 `json:"max_retries,required"`
	// '#/components/schemas/PatchLbHealthMonitorSerializer/properties/timeout'
	// "$.components.schemas.PatchLbHealthMonitorSerializer.properties.timeout"
	Timeout int64 `json:"timeout,required"`
	// '#/components/schemas/PatchLbHealthMonitorSerializer/properties/expected_codes/anyOf/0'
	// "$.components.schemas.PatchLbHealthMonitorSerializer.properties.expected_codes.anyOf[0]"
	ExpectedCodes param.Opt[string] `json:"expected_codes,omitzero"`
	// '#/components/schemas/PatchLbHealthMonitorSerializer/properties/max_retries_down/anyOf/0'
	// "$.components.schemas.PatchLbHealthMonitorSerializer.properties.max_retries_down.anyOf[0]"
	MaxRetriesDown param.Opt[int64] `json:"max_retries_down,omitzero"`
	// '#/components/schemas/PatchLbHealthMonitorSerializer/properties/url_path/anyOf/0'
	// "$.components.schemas.PatchLbHealthMonitorSerializer.properties.url_path.anyOf[0]"
	URLPath param.Opt[string] `json:"url_path,omitzero"`
	// '#/components/schemas/PatchLbHealthMonitorSerializer/properties/http_method/anyOf/0'
	// "$.components.schemas.PatchLbHealthMonitorSerializer.properties.http_method.anyOf[0]"
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod HTTPMethod `json:"http_method,omitzero"`
	// '#/components/schemas/PatchLbHealthMonitorSerializer/properties/type/anyOf/0'
	// "$.components.schemas.PatchLbHealthMonitorSerializer.properties.type.anyOf[0]"
	//
	// Any of "HTTP", "HTTPS", "K8S", "PING", "TCP", "TLS-HELLO", "UDP-CONNECT".
	Type HealthMonitorType `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerPoolUpdateParamsHealthmonitor) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerPoolUpdateParamsHealthmonitor) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolUpdateParamsHealthmonitor
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/PatchLbPoolSerializer/properties/members/anyOf/0/items'
// "$.components.schemas.PatchLbPoolSerializer.properties.members.anyOf[0].items"
//
// The properties Address, ProtocolPort are required.
type LoadBalancerPoolUpdateParamsMember struct {
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/address'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.address"
	Address string `json:"address,required" format:"ipvanyaddress"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/protocol_port'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.protocol_port"
	ProtocolPort int64 `json:"protocol_port,required"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/admin_state_up/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.admin_state_up.anyOf[0]"
	AdminStateUp param.Opt[bool] `json:"admin_state_up,omitzero"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/instance_id/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.instance_id.anyOf[0]"
	InstanceID param.Opt[string] `json:"instance_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/monitor_address/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.monitor_address.anyOf[0]"
	MonitorAddress param.Opt[string] `json:"monitor_address,omitzero" format:"ipvanyaddress"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/monitor_port/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.monitor_port.anyOf[0]"
	MonitorPort param.Opt[int64] `json:"monitor_port,omitzero"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/subnet_id/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.subnet_id.anyOf[0]"
	SubnetID param.Opt[string] `json:"subnet_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/weight/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.weight.anyOf[0]"
	Weight param.Opt[int64] `json:"weight,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerPoolUpdateParamsMember) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerPoolUpdateParamsMember) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolUpdateParamsMember
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/PatchLbPoolSerializer/properties/session_persistence/anyOf/0'
// "$.components.schemas.PatchLbPoolSerializer.properties.session_persistence.anyOf[0]"
//
// The property Type is required.
type LoadBalancerPoolUpdateParamsSessionPersistence struct {
	// '#/components/schemas/MutateLbSessionPersistence/properties/type'
	// "$.components.schemas.MutateLbSessionPersistence.properties.type"
	//
	// Any of "APP_COOKIE", "HTTP_COOKIE", "SOURCE_IP".
	Type SessionPersistenceType `json:"type,omitzero,required"`
	// '#/components/schemas/MutateLbSessionPersistence/properties/cookie_name/anyOf/0'
	// "$.components.schemas.MutateLbSessionPersistence.properties.cookie_name.anyOf[0]"
	CookieName param.Opt[string] `json:"cookie_name,omitzero"`
	// '#/components/schemas/MutateLbSessionPersistence/properties/persistence_granularity/anyOf/0'
	// "$.components.schemas.MutateLbSessionPersistence.properties.persistence_granularity.anyOf[0]"
	PersistenceGranularity param.Opt[string] `json:"persistence_granularity,omitzero"`
	// '#/components/schemas/MutateLbSessionPersistence/properties/persistence_timeout/anyOf/0'
	// "$.components.schemas.MutateLbSessionPersistence.properties.persistence_timeout.anyOf[0]"
	PersistenceTimeout param.Opt[int64] `json:"persistence_timeout,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerPoolUpdateParamsSessionPersistence) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerPoolUpdateParamsSessionPersistence) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolUpdateParamsSessionPersistence
	return param.MarshalObject(r, (*shadow)(&r))
}

type LoadBalancerPoolListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}'].get.parameters[2]"
	Details param.Opt[bool] `query:"details,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}'].get.parameters[3]"
	ListenerID param.Opt[string] `query:"listener_id,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}'].get.parameters[4]"
	LoadbalancerID param.Opt[string] `query:"loadbalancer_id,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerPoolListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LoadBalancerPoolListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerPoolListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LoadBalancerPoolDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bpool_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}/{pool_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bpool_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}/{pool_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerPoolDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type LoadBalancerPoolGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bpool_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}/{pool_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bpool_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}/{pool_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerPoolGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
