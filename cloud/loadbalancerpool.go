// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
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
	tasks          TaskService
}

// NewLoadBalancerPoolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerPoolService(opts ...option.RequestOption) (r LoadBalancerPoolService) {
	r = LoadBalancerPoolService{}
	r.Options = opts
	r.HealthMonitors = NewLoadBalancerPoolHealthMonitorService(opts...)
	r.Members = NewLoadBalancerPoolMemberService(opts...)
	r.tasks = NewTaskService(opts...)
	return
}

// Create load balancer pool
func (r *LoadBalancerPoolService) New(ctx context.Context, params LoadBalancerPoolNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates the specified load balancer pool with the provided changes.
//
// **Behavior:**
//
//   - Simple fields (strings, numbers, booleans) will be updated if provided
//   - Complex objects (nested structures like members, health monitors, etc.) must
//     be specified completely - partial updates are not supported for these objects
//   - Undefined fields will remain unchanged
//   - If no change is detected for a specific field compared to the current pool
//     state, that field will be skipped
//   - If no changes are detected at all across all fields, no task will be created
//     and an empty task list will be returned
//
// **Examples of complex objects that require full specification:**
//
// - Pool members: All member properties must be provided when updating members
// - Health monitors: Complete health monitor configuration must be specified
// - Session persistence: Full session persistence settings must be included
func (r *LoadBalancerPoolService) Update(ctx context.Context, poolID string, params LoadBalancerPoolUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v2/lbpools/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List load balancer pools
func (r *LoadBalancerPoolService) List(ctx context.Context, params LoadBalancerPoolListParams, opts ...option.RequestOption) (res *LoadBalancerPoolList, err error) {
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
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete load balancer pool
func (r *LoadBalancerPoolService) Delete(ctx context.Context, poolID string, body LoadBalancerPoolDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
func (r *LoadBalancerPoolService) Get(ctx context.Context, poolID string, query LoadBalancerPoolGetParams, opts ...option.RequestOption) (res *LoadBalancerPool, err error) {
	opts = slices.Concat(r.Options, opts)
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

// NewAndPoll creates a new pool and polls for completion
func (r *LoadBalancerPoolService) NewAndPoll(ctx context.Context, params LoadBalancerPoolNewParams, opts ...option.RequestOption) (v *LoadBalancerPool, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams LoadBalancerPoolGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]
	task, err := r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Pools) != 1 {
		return nil, errors.New("expected exactly one pool to be created in a task")
	}
	resourceID := task.CreatedResources.Pools[0]

	return r.Get(ctx, resourceID, getParams, opts...)
}

// DeleteAndPoll deletes a pool and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *LoadBalancerPoolService) DeleteAndPoll(ctx context.Context, poolID string, params LoadBalancerPoolDeleteParams, opts ...option.RequestOption) error {
	resource, err := r.Delete(ctx, poolID, params, opts...)
	if err != nil {
		return err
	}

	opts = append(r.Options[:], opts...)
	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	return err
}

// UpdateAndPoll updates a pool and polls for completion of the first task. Use the [TaskService.Poll] method if you
// // need to poll for all tasks.
func (r *LoadBalancerPoolService) UpdateAndPoll(ctx context.Context, poolID string, params LoadBalancerPoolUpdateParams, opts ...option.RequestOption) (v *LoadBalancerPool, err error) {
	resource, err := r.Update(ctx, poolID, params, opts...)
	if err != nil {
		return
	}

	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams LoadBalancerPoolGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) == 0 {
		return nil, errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, poolID, getParams, opts...)
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
	LoadBalancerID param.Opt[string] `json:"load_balancer_id,omitzero" format:"uuid4"`
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
	// Expected HTTP response codes. Can be a single code or a range of codes. Can only
	// be used together with `HTTP` or `HTTPS` health monitor type. For example,
	// 200,202,300-302,401,403,404,500-504. If not specified, the default is 200.
	ExpectedCodes param.Opt[string] `json:"expected_codes,omitzero"`
	// URL Path. Defaults to '/'. Can only be used together with `HTTP` or `HTTPS`
	// health monitor type.
	URLPath param.Opt[string] `json:"url_path,omitzero"`
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
	// Either `subnet_id` or `instance_id` should be provided
	InstanceID param.Opt[string] `json:"instance_id,omitzero" format:"uuid4"`
	// An alternate IP address used for health monitoring of a backend member. Default
	// is null which monitors the member address.
	MonitorAddress param.Opt[string] `json:"monitor_address,omitzero" format:"ipvanyaddress"`
	// An alternate protocol port used for health monitoring of a backend member.
	// Default is null which monitors the member `protocol_port`.
	MonitorPort param.Opt[int64] `json:"monitor_port,omitzero"`
	// `subnet_id` in which `address` is present. Either `subnet_id` or `instance_id`
	// should be provided
	SubnetID param.Opt[string] `json:"subnet_id,omitzero" format:"uuid4"`
	// Administrative state of the resource. When set to true, the resource is enabled
	// and operational. When set to false, the resource is disabled and will not
	// process traffic. When null is passed, the value is skipped and defaults to true.
	AdminStateUp param.Opt[bool] `json:"admin_state_up,omitzero"`
	// Set to true if the member is a backup member, to which traffic will be sent
	// exclusively when all non-backup members will be unreachable. It allows to
	// realize ACTIVE-BACKUP load balancing without thinking about VRRP and VIP
	// configuration. Default is false.
	Backup param.Opt[bool] `json:"backup,omitzero"`
	// Member weight. Valid values are 0 < `weight` <= 256, defaults to 1. Controls
	// traffic distribution based on the pool's load balancing algorithm:
	//
	//   - `ROUND_ROBIN`: Distributes connections to each member in turn according to
	//     weights. Higher weight = more turns in the cycle. Example: weights 3 vs 1 =
	//     ~75% vs ~25% of requests.
	//   - `LEAST_CONNECTIONS`: Sends new connections to the member with fewest active
	//     connections, performing round-robin within groups of the same normalized load.
	//     Higher weight = allowed to hold more simultaneous connections before being
	//     considered 'more loaded'. Example: weights 2 vs 1 means 20 vs 10 active
	//     connections is treated as balanced.
	//   - `SOURCE_IP`: Routes clients consistently to the same member by hashing client
	//     source IP; hash result is modulo total weight of running members. Higher
	//     weight = more hash buckets, so more client IPs map to that member. Example:
	//     weights 2 vs 1 = roughly two-thirds of distinct client IPs map to the
	//     higher-weight member.
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
	// Subnet mask if `source_ip` is used. For UDP ports only
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
	// Expected HTTP response codes. Can be a single code or a range of codes. Can only
	// be used together with `HTTP` or `HTTPS` health monitor type. For example,
	// 200,202,300-302,401,403,404,500-504. If not specified, the default is 200.
	ExpectedCodes param.Opt[string] `json:"expected_codes,omitzero"`
	// URL Path. Defaults to '/'. Can only be used together with `HTTP` or `HTTPS`
	// health monitor type.
	URLPath param.Opt[string] `json:"url_path,omitzero"`
	// Number of failures before the member is switched to ERROR state.
	MaxRetriesDown param.Opt[int64] `json:"max_retries_down,omitzero"`
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
	// Either `subnet_id` or `instance_id` should be provided
	InstanceID param.Opt[string] `json:"instance_id,omitzero" format:"uuid4"`
	// An alternate IP address used for health monitoring of a backend member. Default
	// is null which monitors the member address.
	MonitorAddress param.Opt[string] `json:"monitor_address,omitzero" format:"ipvanyaddress"`
	// An alternate protocol port used for health monitoring of a backend member.
	// Default is null which monitors the member `protocol_port`.
	MonitorPort param.Opt[int64] `json:"monitor_port,omitzero"`
	// `subnet_id` in which `address` is present. Either `subnet_id` or `instance_id`
	// should be provided
	SubnetID param.Opt[string] `json:"subnet_id,omitzero" format:"uuid4"`
	// Administrative state of the resource. When set to true, the resource is enabled
	// and operational. When set to false, the resource is disabled and will not
	// process traffic. When null is passed, the value is skipped and defaults to true.
	AdminStateUp param.Opt[bool] `json:"admin_state_up,omitzero"`
	// Set to true if the member is a backup member, to which traffic will be sent
	// exclusively when all non-backup members will be unreachable. It allows to
	// realize ACTIVE-BACKUP load balancing without thinking about VRRP and VIP
	// configuration. Default is false.
	Backup param.Opt[bool] `json:"backup,omitzero"`
	// Member weight. Valid values are 0 < `weight` <= 256, defaults to 1. Controls
	// traffic distribution based on the pool's load balancing algorithm:
	//
	//   - `ROUND_ROBIN`: Distributes connections to each member in turn according to
	//     weights. Higher weight = more turns in the cycle. Example: weights 3 vs 1 =
	//     ~75% vs ~25% of requests.
	//   - `LEAST_CONNECTIONS`: Sends new connections to the member with fewest active
	//     connections, performing round-robin within groups of the same normalized load.
	//     Higher weight = allowed to hold more simultaneous connections before being
	//     considered 'more loaded'. Example: weights 2 vs 1 means 20 vs 10 active
	//     connections is treated as balanced.
	//   - `SOURCE_IP`: Routes clients consistently to the same member by hashing client
	//     source IP; hash result is modulo total weight of running members. Higher
	//     weight = more hash buckets, so more client IPs map to that member. Example:
	//     weights 2 vs 1 = roughly two-thirds of distinct client IPs map to the
	//     higher-weight member.
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
	// Subnet mask if `source_ip` is used. For UDP ports only
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
	LoadBalancerID param.Opt[string] `query:"load_balancer_id,omitzero" format:"uuid4" json:"-"`
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
