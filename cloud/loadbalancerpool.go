// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
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
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
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
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List load balancer pools
func (r *LoadBalancerPoolService) List(ctx context.Context, params LoadBalancerPoolListParams, opts ...option.RequestOption) (res *LoadBalancerPoolList, err error) {
	opts = append(r.Options[:], opts...)
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
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get load balancer pool
func (r *LoadBalancerPoolService) Get(ctx context.Context, poolID string, query LoadBalancerPoolGetParams, opts ...option.RequestOption) (res *LoadBalancerPoolGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.Valid() {
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

type LoadBalancerPoolGetResponse struct {
	// Pool ID
	ID string `json:"id,required" format:"uuid4"`
	// Secret ID of CA certificate bundle
	CaSecretID string `json:"ca_secret_id,required" format:"uuid4"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required" format:"uuid4"`
	// Secret ID of CA revocation list file
	CrlSecretID string `json:"crl_secret_id,required" format:"uuid4"`
	// Health monitor parameters
	Healthmonitor HealthMonitor `json:"healthmonitor,required"`
	// Load balancer algorithm
	//
	// Any of "LEAST_CONNECTIONS", "ROUND_ROBIN", "SOURCE_IP".
	LbAlgorithm LbAlgorithm `json:"lb_algorithm,required"`
	// Listeners IDs
	Listeners []LoadBalancerPoolGetResponseListener `json:"listeners,required"`
	// Load balancers IDs
	Loadbalancers []LoadBalancerPoolGetResponseLoadbalancer `json:"loadbalancers,required"`
	// Pool members
	Members []Member `json:"members,required"`
	// Pool name
	Name string `json:"name,required"`
	// Pool operating status
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// Protocol
	//
	// Any of "HTTP", "HTTPS", "PROXY", "PROXYV2", "TCP", "UDP".
	Protocol LbPoolProtocol `json:"protocol,required"`
	// Pool lifecycle status
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Secret ID for TLS client authentication to the member servers
	SecretID string `json:"secret_id,required" format:"uuid4"`
	// Session persistence parameters
	SessionPersistence SessionPersistence `json:"session_persistence,required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,required" format:"uuid4"`
	// Frontend client inactivity timeout in milliseconds
	TimeoutClientData int64 `json:"timeout_client_data,required"`
	// Backend member connection timeout in milliseconds
	TimeoutMemberConnect int64 `json:"timeout_member_connect,required"`
	// Backend member inactivity timeout in milliseconds
	TimeoutMemberData int64 `json:"timeout_member_data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CaSecretID           respjson.Field
		CreatorTaskID        respjson.Field
		CrlSecretID          respjson.Field
		Healthmonitor        respjson.Field
		LbAlgorithm          respjson.Field
		Listeners            respjson.Field
		Loadbalancers        respjson.Field
		Members              respjson.Field
		Name                 respjson.Field
		OperatingStatus      respjson.Field
		Protocol             respjson.Field
		ProvisioningStatus   respjson.Field
		SecretID             respjson.Field
		SessionPersistence   respjson.Field
		TaskID               respjson.Field
		TimeoutClientData    respjson.Field
		TimeoutMemberConnect respjson.Field
		TimeoutMemberData    respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerPoolGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerPoolGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolGetResponseListener struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerPoolGetResponseListener) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerPoolGetResponseListener) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolGetResponseLoadbalancer struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerPoolGetResponseLoadbalancer) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerPoolGetResponseLoadbalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Load balancer algorithm
	//
	// Any of "LEAST_CONNECTIONS", "ROUND_ROBIN", "SOURCE_IP".
	LbAlgorithm LbAlgorithm `json:"lb_algorithm,omitzero,required"`
	// Pool name
	Name string `json:"name,required"`
	// Protocol
	//
	// Any of "HTTP", "HTTPS", "PROXY", "PROXYV2", "TCP", "UDP".
	Protocol LbPoolProtocol `json:"protocol,omitzero,required"`
	// Secret ID of CA certificate bundle
	CaSecretID param.Opt[string] `json:"ca_secret_id,omitzero" format:"uuid4"`
	// Secret ID of CA revocation list file
	CrlSecretID param.Opt[string] `json:"crl_secret_id,omitzero" format:"uuid4"`
	// Listener ID
	ListenerID param.Opt[string] `json:"listener_id,omitzero" format:"uuid4"`
	// Loadbalancer ID
	LoadbalancerID param.Opt[string] `json:"loadbalancer_id,omitzero" format:"uuid4"`
	// Secret ID for TLS client authentication to the member servers
	SecretID param.Opt[string] `json:"secret_id,omitzero" format:"uuid4"`
	// Frontend client inactivity timeout in milliseconds
	TimeoutClientData param.Opt[int64] `json:"timeout_client_data,omitzero"`
	// Backend member connection timeout in milliseconds
	TimeoutMemberConnect param.Opt[int64] `json:"timeout_member_connect,omitzero"`
	// Backend member inactivity timeout in milliseconds
	TimeoutMemberData param.Opt[int64] `json:"timeout_member_data,omitzero"`
	// Health monitor details
	Healthmonitor LoadBalancerPoolNewParamsHealthmonitor `json:"healthmonitor,omitzero"`
	// Pool members
	Members []LoadBalancerPoolNewParamsMember `json:"members,omitzero"`
	// Session persistence details
	SessionPersistence LoadBalancerPoolNewParamsSessionPersistence `json:"session_persistence,omitzero"`
	paramObj
}

func (r LoadBalancerPoolNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerPoolNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Health monitor details
//
// The properties Delay, MaxRetries, Timeout, Type are required.
type LoadBalancerPoolNewParamsHealthmonitor struct {
	// The time, in seconds, between sending probes to members
	Delay int64 `json:"delay,required"`
	// Number of successes before the member is switched to ONLINE state
	MaxRetries int64 `json:"max_retries,required"`
	// The maximum time to connect. Must be less than the delay value
	Timeout int64 `json:"timeout,required"`
	// Health monitor type. Once health monitor is created, cannot be changed.
	//
	// Any of "HTTP", "HTTPS", "K8S", "PING", "TCP", "TLS-HELLO", "UDP-CONNECT".
	Type LbHealthMonitorType `json:"type,omitzero,required"`
	// Can only be used together with `HTTP` or `HTTPS` health monitor type.
	ExpectedCodes param.Opt[string] `json:"expected_codes,omitzero"`
	// Number of failures before the member is switched to ERROR state.
	MaxRetriesDown param.Opt[int64] `json:"max_retries_down,omitzero"`
	// URL Path. Defaults to '/'. Can only be used together with `HTTP` or `HTTPS`
	// health monitor type.
	URLPath param.Opt[string] `json:"url_path,omitzero"`
	// HTTP method. Can only be used together with `HTTP` or `HTTPS` health monitor
	// type.
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod HTTPMethod `json:"http_method,omitzero"`
	paramObj
}

func (r LoadBalancerPoolNewParamsHealthmonitor) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolNewParamsHealthmonitor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerPoolNewParamsHealthmonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Address, ProtocolPort are required.
type LoadBalancerPoolNewParamsMember struct {
	// Member IP address
	Address string `json:"address,required" format:"ipvanyaddress"`
	// Member IP port
	ProtocolPort int64 `json:"protocol_port,required"`
	// true if enabled. Defaults to true
	AdminStateUp param.Opt[bool] `json:"admin_state_up,omitzero"`
	// Either subnet_id or instance_id should be provided
	InstanceID param.Opt[string] `json:"instance_id,omitzero" format:"uuid4"`
	// An alternate IP address used for health monitoring of a backend member. Default
	// is null which monitors the member address.
	MonitorAddress param.Opt[string] `json:"monitor_address,omitzero" format:"ipvanyaddress"`
	// An alternate protocol port used for health monitoring of a backend member.
	// Default is null which monitors the member protocol_port.
	MonitorPort param.Opt[int64] `json:"monitor_port,omitzero"`
	// Either subnet_id or instance_id should be provided
	SubnetID param.Opt[string] `json:"subnet_id,omitzero" format:"uuid4"`
	// Member weight. Valid values: 0 to 256, defaults to 1
	Weight param.Opt[int64] `json:"weight,omitzero"`
	paramObj
}

func (r LoadBalancerPoolNewParamsMember) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolNewParamsMember
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerPoolNewParamsMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Session persistence details
//
// The property Type is required.
type LoadBalancerPoolNewParamsSessionPersistence struct {
	// Session persistence type
	//
	// Any of "APP_COOKIE", "HTTP_COOKIE", "SOURCE_IP".
	Type LbSessionPersistenceType `json:"type,omitzero,required"`
	// Should be set if app cookie or http cookie is used
	CookieName param.Opt[string] `json:"cookie_name,omitzero"`
	// Subnet mask if source_ip is used. For UDP ports only
	PersistenceGranularity param.Opt[string] `json:"persistence_granularity,omitzero"`
	// Session persistence timeout. For UDP ports only
	PersistenceTimeout param.Opt[int64] `json:"persistence_timeout,omitzero"`
	paramObj
}

func (r LoadBalancerPoolNewParamsSessionPersistence) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolNewParamsSessionPersistence
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerPoolNewParamsSessionPersistence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Secret ID of CA certificate bundle
	CaSecretID param.Opt[string] `json:"ca_secret_id,omitzero" format:"uuid4"`
	// Secret ID of CA revocation list file
	CrlSecretID param.Opt[string] `json:"crl_secret_id,omitzero" format:"uuid4"`
	// Secret ID for TLS client authentication to the member servers
	SecretID param.Opt[string] `json:"secret_id,omitzero" format:"uuid4"`
	// Frontend client inactivity timeout in milliseconds
	TimeoutClientData param.Opt[int64] `json:"timeout_client_data,omitzero"`
	// Backend member connection timeout in milliseconds
	TimeoutMemberConnect param.Opt[int64] `json:"timeout_member_connect,omitzero"`
	// Backend member inactivity timeout in milliseconds
	TimeoutMemberData param.Opt[int64] `json:"timeout_member_data,omitzero"`
	// New pool name
	Name param.Opt[string] `json:"name,omitzero"`
	// New pool health monitor settings
	Healthmonitor LoadBalancerPoolUpdateParamsHealthmonitor `json:"healthmonitor,omitzero"`
	// New sequence of load balancer pool members. If members are the same (by
	// address + port), they will be kept as is without recreation and downtime.
	Members []LoadBalancerPoolUpdateParamsMember `json:"members,omitzero"`
	// New session persistence settings
	SessionPersistence LoadBalancerPoolUpdateParamsSessionPersistence `json:"session_persistence,omitzero"`
	// New load balancer pool algorithm of how to distribute requests
	//
	// Any of "LEAST_CONNECTIONS", "ROUND_ROBIN", "SOURCE_IP".
	LbAlgorithm LbAlgorithm `json:"lb_algorithm,omitzero"`
	// New communication protocol
	//
	// Any of "HTTP", "HTTPS", "PROXY", "PROXYV2", "TCP", "UDP".
	Protocol LbPoolProtocol `json:"protocol,omitzero"`
	paramObj
}

func (r LoadBalancerPoolUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerPoolUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// New pool health monitor settings
//
// The properties Delay, MaxRetries, Timeout are required.
type LoadBalancerPoolUpdateParamsHealthmonitor struct {
	// The time, in seconds, between sending probes to members
	Delay int64 `json:"delay,required"`
	// Number of successes before the member is switched to ONLINE state
	MaxRetries int64 `json:"max_retries,required"`
	// The maximum time to connect. Must be less than the delay value
	Timeout int64 `json:"timeout,required"`
	// Can only be used together with `HTTP` or `HTTPS` health monitor type.
	ExpectedCodes param.Opt[string] `json:"expected_codes,omitzero"`
	// Number of failures before the member is switched to ERROR state.
	MaxRetriesDown param.Opt[int64] `json:"max_retries_down,omitzero"`
	// URL Path. Defaults to '/'. Can only be used together with `HTTP` or `HTTPS`
	// health monitor type.
	URLPath param.Opt[string] `json:"url_path,omitzero"`
	// HTTP method. Can only be used together with `HTTP` or `HTTPS` health monitor
	// type.
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod HTTPMethod `json:"http_method,omitzero"`
	// Health monitor type. Once health monitor is created, cannot be changed.
	//
	// Any of "HTTP", "HTTPS", "K8S", "PING", "TCP", "TLS-HELLO", "UDP-CONNECT".
	Type LbHealthMonitorType `json:"type,omitzero"`
	paramObj
}

func (r LoadBalancerPoolUpdateParamsHealthmonitor) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolUpdateParamsHealthmonitor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerPoolUpdateParamsHealthmonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Address, ProtocolPort are required.
type LoadBalancerPoolUpdateParamsMember struct {
	// Member IP address
	Address string `json:"address,required" format:"ipvanyaddress"`
	// Member IP port
	ProtocolPort int64 `json:"protocol_port,required"`
	// true if enabled. Defaults to true
	AdminStateUp param.Opt[bool] `json:"admin_state_up,omitzero"`
	// Either subnet_id or instance_id should be provided
	InstanceID param.Opt[string] `json:"instance_id,omitzero" format:"uuid4"`
	// An alternate IP address used for health monitoring of a backend member. Default
	// is null which monitors the member address.
	MonitorAddress param.Opt[string] `json:"monitor_address,omitzero" format:"ipvanyaddress"`
	// An alternate protocol port used for health monitoring of a backend member.
	// Default is null which monitors the member protocol_port.
	MonitorPort param.Opt[int64] `json:"monitor_port,omitzero"`
	// Either subnet_id or instance_id should be provided
	SubnetID param.Opt[string] `json:"subnet_id,omitzero" format:"uuid4"`
	// Member weight. Valid values: 0 to 256, defaults to 1
	Weight param.Opt[int64] `json:"weight,omitzero"`
	paramObj
}

func (r LoadBalancerPoolUpdateParamsMember) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolUpdateParamsMember
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerPoolUpdateParamsMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// New session persistence settings
//
// The property Type is required.
type LoadBalancerPoolUpdateParamsSessionPersistence struct {
	// Session persistence type
	//
	// Any of "APP_COOKIE", "HTTP_COOKIE", "SOURCE_IP".
	Type LbSessionPersistenceType `json:"type,omitzero,required"`
	// Should be set if app cookie or http cookie is used
	CookieName param.Opt[string] `json:"cookie_name,omitzero"`
	// Subnet mask if source_ip is used. For UDP ports only
	PersistenceGranularity param.Opt[string] `json:"persistence_granularity,omitzero"`
	// Session persistence timeout. For UDP ports only
	PersistenceTimeout param.Opt[int64] `json:"persistence_timeout,omitzero"`
	paramObj
}

func (r LoadBalancerPoolUpdateParamsSessionPersistence) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolUpdateParamsSessionPersistence
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerPoolUpdateParamsSessionPersistence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Show members and Health Monitor details
	Details param.Opt[bool] `query:"details,omitzero" json:"-"`
	// Listener ID
	ListenerID param.Opt[string] `query:"listener_id,omitzero" format:"uuid4" json:"-"`
	// Load Balancer ID
	LoadbalancerID param.Opt[string] `query:"loadbalancer_id,omitzero" format:"uuid4" json:"-"`
	paramObj
}

// URLQuery serializes [LoadBalancerPoolListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerPoolListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LoadBalancerPoolDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type LoadBalancerPoolGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
