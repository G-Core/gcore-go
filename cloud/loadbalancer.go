// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/G-Core/gcore-go/shared/constant"
)

// LoadBalancerService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoadBalancerService] method instead.
type LoadBalancerService struct {
	Options    []option.RequestOption
	L7Policies LoadBalancerL7PolicyService
	Flavors    LoadBalancerFlavorService
	Listeners  LoadBalancerListenerService
	Pools      LoadBalancerPoolService
	Metrics    LoadBalancerMetricService
	Statuses   LoadBalancerStatusService
}

// NewLoadBalancerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLoadBalancerService(opts ...option.RequestOption) (r LoadBalancerService) {
	r = LoadBalancerService{}
	r.Options = opts
	r.L7Policies = NewLoadBalancerL7PolicyService(opts...)
	r.Flavors = NewLoadBalancerFlavorService(opts...)
	r.Listeners = NewLoadBalancerListenerService(opts...)
	r.Pools = NewLoadBalancerPoolService(opts...)
	r.Metrics = NewLoadBalancerMetricService(opts...)
	r.Statuses = NewLoadBalancerStatusService(opts...)
	return
}

// Create load balancer
func (r *LoadBalancerService) New(ctx context.Context, params LoadBalancerNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Rename load balancer, activate/deactivate logging, update preferred connectivity
// type and/or modify load balancer tags. The request will only process the fields
// that are provided in the request body. Any fields that are not included will
// remain unchanged.
func (r *LoadBalancerService) Update(ctx context.Context, loadBalancerID string, params LoadBalancerUpdateParams, opts ...option.RequestOption) (res *LoadBalancer, err error) {
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
	if loadBalancerID == "" {
		err = errors.New("missing required load_balancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List load balancers
func (r *LoadBalancerService) List(ctx context.Context, params LoadBalancerListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LoadBalancer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List load balancers
func (r *LoadBalancerService) ListAutoPaging(ctx context.Context, params LoadBalancerListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LoadBalancer] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete load balancer
func (r *LoadBalancerService) Delete(ctx context.Context, loadBalancerID string, body LoadBalancerDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if loadBalancerID == "" {
		err = errors.New("missing required load_balancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Failover load balancer
func (r *LoadBalancerService) Failover(ctx context.Context, loadBalancerID string, params LoadBalancerFailoverParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if loadBalancerID == "" {
		err = errors.New("missing required load_balancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/failover", params.ProjectID.Value, params.RegionID.Value, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get load balancer
func (r *LoadBalancerService) Get(ctx context.Context, loadBalancerID string, params LoadBalancerGetParams, opts ...option.RequestOption) (res *LoadBalancer, err error) {
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
	if loadBalancerID == "" {
		err = errors.New("missing required load_balancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Resize load balancer
func (r *LoadBalancerService) Resize(ctx context.Context, loadBalancerID string, params LoadBalancerResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if loadBalancerID == "" {
		err = errors.New("missing required load_balancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/resize", params.ProjectID.Value, params.RegionID.Value, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type HealthMonitor struct {
	// Health monitor ID
	ID string `json:"id,required" format:"uuid4"`
	// Administrative state of the resource. When set to true, the resource is enabled
	// and operational. When set to false, the resource is disabled and will not
	// process traffic. When null is passed, the value is skipped and defaults to true.
	AdminStateUp bool `json:"admin_state_up,required"`
	// The time, in seconds, between sending probes to members
	Delay int64 `json:"delay,required"`
	// Number of successes before the member is switched to ONLINE state
	MaxRetries int64 `json:"max_retries,required"`
	// Number of failures before the member is switched to ERROR state
	MaxRetriesDown int64 `json:"max_retries_down,required"`
	// Health Monitor operating status
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// Health monitor lifecycle status
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// The maximum time to connect. Must be less than the delay value
	Timeout int64 `json:"timeout,required"`
	// Health monitor type. Once health monitor is created, cannot be changed.
	//
	// Any of "HTTP", "HTTPS", "K8S", "PING", "TCP", "TLS-HELLO", "UDP-CONNECT".
	Type LbHealthMonitorType `json:"type,required"`
	// Expected HTTP response codes. Can be a single code or a range of codes. Can only
	// be used together with `HTTP` or `HTTPS` health monitor type. For example,
	// 200,202,300-302,401,403,404,500-504. If not specified, the default is 200.
	ExpectedCodes string `json:"expected_codes,nullable"`
	// HTTP method
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod HTTPMethod `json:"http_method,nullable"`
	// URL Path. Defaults to '/'
	URLPath string `json:"url_path,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AdminStateUp       respjson.Field
		Delay              respjson.Field
		MaxRetries         respjson.Field
		MaxRetriesDown     respjson.Field
		OperatingStatus    respjson.Field
		ProvisioningStatus respjson.Field
		Timeout            respjson.Field
		Type               respjson.Field
		ExpectedCodes      respjson.Field
		HTTPMethod         respjson.Field
		URLPath            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthMonitor) RawJSON() string { return r.JSON.raw }
func (r *HealthMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HealthMonitorStatus struct {
	// UUID of the entity
	ID string `json:"id,required" format:"uuid"`
	// Operating status of the entity
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// Provisioning status of the entity
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Type of the Health Monitor
	//
	// Any of "HTTP", "HTTPS", "K8S", "PING", "TCP", "TLS-HELLO", "UDP-CONNECT".
	Type LbHealthMonitorType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		OperatingStatus    respjson.Field
		ProvisioningStatus respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthMonitorStatus) RawJSON() string { return r.JSON.raw }
func (r *HealthMonitorStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LbAlgorithm string

const (
	LbAlgorithmLeastConnections LbAlgorithm = "LEAST_CONNECTIONS"
	LbAlgorithmRoundRobin       LbAlgorithm = "ROUND_ROBIN"
	LbAlgorithmSourceIP         LbAlgorithm = "SOURCE_IP"
)

type LbHealthMonitorType string

const (
	LbHealthMonitorTypeHTTP       LbHealthMonitorType = "HTTP"
	LbHealthMonitorTypeHTTPS      LbHealthMonitorType = "HTTPS"
	LbHealthMonitorTypeK8S        LbHealthMonitorType = "K8S"
	LbHealthMonitorTypePing       LbHealthMonitorType = "PING"
	LbHealthMonitorTypeTcp        LbHealthMonitorType = "TCP"
	LbHealthMonitorTypeTlsHello   LbHealthMonitorType = "TLS-HELLO"
	LbHealthMonitorTypeUdpConnect LbHealthMonitorType = "UDP-CONNECT"
)

type LbListenerProtocol string

const (
	LbListenerProtocolHTTP            LbListenerProtocol = "HTTP"
	LbListenerProtocolHTTPS           LbListenerProtocol = "HTTPS"
	LbListenerProtocolPrometheus      LbListenerProtocol = "PROMETHEUS"
	LbListenerProtocolTcp             LbListenerProtocol = "TCP"
	LbListenerProtocolTerminatedHTTPS LbListenerProtocol = "TERMINATED_HTTPS"
	LbListenerProtocolUdp             LbListenerProtocol = "UDP"
)

type LbPoolProtocol string

const (
	LbPoolProtocolHTTP    LbPoolProtocol = "HTTP"
	LbPoolProtocolHTTPS   LbPoolProtocol = "HTTPS"
	LbPoolProtocolProxy   LbPoolProtocol = "PROXY"
	LbPoolProtocolProxyv2 LbPoolProtocol = "PROXYV2"
	LbPoolProtocolTcp     LbPoolProtocol = "TCP"
	LbPoolProtocolUdp     LbPoolProtocol = "UDP"
)

type LbSessionPersistenceType string

const (
	LbSessionPersistenceTypeAppCookie  LbSessionPersistenceType = "APP_COOKIE"
	LbSessionPersistenceTypeHTTPCookie LbSessionPersistenceType = "HTTP_COOKIE"
	LbSessionPersistenceTypeSourceIP   LbSessionPersistenceType = "SOURCE_IP"
)

type ListenerStatus struct {
	// UUID of the entity
	ID string `json:"id,required" format:"uuid"`
	// Name of the load balancer listener
	Name string `json:"name,required"`
	// Operating status of the entity
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// Pools of the Listeners
	Pools []PoolStatus `json:"pools,required"`
	// Provisioning status of the entity
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Name               respjson.Field
		OperatingStatus    respjson.Field
		Pools              respjson.Field
		ProvisioningStatus respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListenerStatus) RawJSON() string { return r.JSON.raw }
func (r *ListenerStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerFlavorDetail struct {
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id,required"`
	// Flavor name
	FlavorName string `json:"flavor_name,required"`
	// Additional hardware description.
	HardwareDescription FlavorHardwareDescription `json:"hardware_description,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Virtual CPU count. For bare metal flavors, it's a physical CPU count
	Vcpus int64 `json:"vcpus,required"`
	// Currency code. Shown if the `include_prices` query parameter if set to true
	CurrencyCode string `json:"currency_code,nullable"`
	// Price per hour. Shown if the `include_prices` query parameter if set to true
	PricePerHour float64 `json:"price_per_hour,nullable"`
	// Price per month. Shown if the `include_prices` query parameter if set to true
	PricePerMonth float64 `json:"price_per_month,nullable"`
	// Price status for the UI
	//
	// Any of "error", "hide", "show".
	PriceStatus LoadBalancerFlavorDetailPriceStatus `json:"price_status,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FlavorID            respjson.Field
		FlavorName          respjson.Field
		HardwareDescription respjson.Field
		Ram                 respjson.Field
		Vcpus               respjson.Field
		CurrencyCode        respjson.Field
		PricePerHour        respjson.Field
		PricePerMonth       respjson.Field
		PriceStatus         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerFlavorDetail) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerFlavorDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price status for the UI
type LoadBalancerFlavorDetailPriceStatus string

const (
	LoadBalancerFlavorDetailPriceStatusError LoadBalancerFlavorDetailPriceStatus = "error"
	LoadBalancerFlavorDetailPriceStatusHide  LoadBalancerFlavorDetailPriceStatus = "hide"
	LoadBalancerFlavorDetailPriceStatusShow  LoadBalancerFlavorDetailPriceStatus = "show"
)

type LoadBalancerFlavorList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []LoadBalancerFlavorDetail `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerFlavorList) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerFlavorList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerL7Policy struct {
	// ID
	ID string `json:"id,required"`
	// Action
	//
	// Any of "REDIRECT_PREFIX", "REDIRECT_TO_POOL", "REDIRECT_TO_URL", "REJECT".
	Action LoadBalancerL7PolicyAction `json:"action,required"`
	// Listener ID
	ListenerID string `json:"listener_id,required"`
	// Human-readable name of the policy
	Name string `json:"name,required"`
	// L7 policy operating status
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// The position of this policy on the listener. Positions start at 1.
	Position int64 `json:"position,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Requests matching this policy will be redirected to the specified URL or Prefix
	// URL with the HTTP response code. Valid if action is `REDIRECT_TO_URL` or
	// `REDIRECT_PREFIX`. Valid options are 301, 302, 303, 307, or 308. Default is 302.
	RedirectHTTPCode int64 `json:"redirect_http_code,required"`
	// Requests matching this policy will be redirected to the pool with this ID. Only
	// valid if action is `REDIRECT_TO_POOL`.
	RedirectPoolID string `json:"redirect_pool_id,required"`
	// Requests matching this policy will be redirected to this Prefix URL. Only valid
	// if action is `REDIRECT_PREFIX`.
	RedirectPrefix string `json:"redirect_prefix,required"`
	// Requests matching this policy will be redirected to this URL. Only valid if
	// action is `REDIRECT_TO_URL`.
	RedirectURL string `json:"redirect_url,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Rules. All the rules associated with a given policy are logically ANDed
	// together. A request must match all the policy’s rules to match the policy.If you
	// need to express a logical OR operation between rules, then do this by creating
	// multiple policies with the same action.
	Rules []LoadBalancerL7PolicyRule `json:"rules,required"`
	// A list of simple strings assigned to the resource.
	Tags []string `json:"tags,required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Action             respjson.Field
		ListenerID         respjson.Field
		Name               respjson.Field
		OperatingStatus    respjson.Field
		Position           respjson.Field
		ProjectID          respjson.Field
		ProvisioningStatus respjson.Field
		RedirectHTTPCode   respjson.Field
		RedirectPoolID     respjson.Field
		RedirectPrefix     respjson.Field
		RedirectURL        respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Rules              respjson.Field
		Tags               respjson.Field
		TaskID             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerL7Policy) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerL7Policy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action
type LoadBalancerL7PolicyAction string

const (
	LoadBalancerL7PolicyActionRedirectPrefix LoadBalancerL7PolicyAction = "REDIRECT_PREFIX"
	LoadBalancerL7PolicyActionRedirectToPool LoadBalancerL7PolicyAction = "REDIRECT_TO_POOL"
	LoadBalancerL7PolicyActionRedirectToURL  LoadBalancerL7PolicyAction = "REDIRECT_TO_URL"
	LoadBalancerL7PolicyActionReject         LoadBalancerL7PolicyAction = "REJECT"
)

type LoadBalancerL7PolicyRule struct {
	// L7Rule ID
	ID string `json:"id,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ProjectID   respjson.Field
		Region      respjson.Field
		RegionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerL7PolicyRule) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerL7PolicyRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerL7PolicyList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []LoadBalancerL7Policy `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerL7PolicyList) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerL7PolicyList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerL7Rule struct {
	// L7Rule ID
	ID string `json:"id,required"`
	// The comparison type for the L7 rule
	//
	// Any of "CONTAINS", "ENDS_WITH", "EQUAL_TO", "REGEX", "STARTS_WITH".
	CompareType LoadBalancerL7RuleCompareType `json:"compare_type,required"`
	// When true the logic of the rule is inverted. For example, with invert true,
	// 'equal to' would become 'not equal to'. Default is false.
	Invert bool `json:"invert,required"`
	// The key to use for the comparison. For example, the name of the cookie to
	// evaluate.
	Key string `json:"key,required"`
	// L7 policy operating status
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// A list of simple strings assigned to the l7 rule
	Tags []string `json:"tags,required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,required" format:"uuid4"`
	// The L7 rule type
	//
	// Any of "COOKIE", "FILE_TYPE", "HEADER", "HOST_NAME", "PATH",
	// "SSL_CONN_HAS_CERT", "SSL_DN_FIELD", "SSL_VERIFY_RESULT".
	Type LoadBalancerL7RuleType `json:"type,required"`
	// The value to use for the comparison. For example, the file type to compare.
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CompareType        respjson.Field
		Invert             respjson.Field
		Key                respjson.Field
		OperatingStatus    respjson.Field
		ProjectID          respjson.Field
		ProvisioningStatus respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Tags               respjson.Field
		TaskID             respjson.Field
		Type               respjson.Field
		Value              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerL7Rule) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerL7Rule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The comparison type for the L7 rule
type LoadBalancerL7RuleCompareType string

const (
	LoadBalancerL7RuleCompareTypeContains   LoadBalancerL7RuleCompareType = "CONTAINS"
	LoadBalancerL7RuleCompareTypeEndsWith   LoadBalancerL7RuleCompareType = "ENDS_WITH"
	LoadBalancerL7RuleCompareTypeEqualTo    LoadBalancerL7RuleCompareType = "EQUAL_TO"
	LoadBalancerL7RuleCompareTypeRegex      LoadBalancerL7RuleCompareType = "REGEX"
	LoadBalancerL7RuleCompareTypeStartsWith LoadBalancerL7RuleCompareType = "STARTS_WITH"
)

// The L7 rule type
type LoadBalancerL7RuleType string

const (
	LoadBalancerL7RuleTypeCookie          LoadBalancerL7RuleType = "COOKIE"
	LoadBalancerL7RuleTypeFileType        LoadBalancerL7RuleType = "FILE_TYPE"
	LoadBalancerL7RuleTypeHeader          LoadBalancerL7RuleType = "HEADER"
	LoadBalancerL7RuleTypeHostName        LoadBalancerL7RuleType = "HOST_NAME"
	LoadBalancerL7RuleTypePath            LoadBalancerL7RuleType = "PATH"
	LoadBalancerL7RuleTypeSslConnHasCert  LoadBalancerL7RuleType = "SSL_CONN_HAS_CERT"
	LoadBalancerL7RuleTypeSslDnField      LoadBalancerL7RuleType = "SSL_DN_FIELD"
	LoadBalancerL7RuleTypeSslVerifyResult LoadBalancerL7RuleType = "SSL_VERIFY_RESULT"
)

type LoadBalancerL7RuleList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []LoadBalancerL7Rule `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerL7RuleList) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerL7RuleList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerListenerDetail struct {
	// Load balancer listener ID
	ID string `json:"id,required" format:"uuid4"`
	// Network CIDRs from which service will be accessible
	AllowedCidrs []string `json:"allowed_cidrs,required" format:"ipvanynetwork"`
	// Limit of simultaneous connections
	ConnectionLimit int64 `json:"connection_limit,required"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required" format:"uuid4"`
	// Dictionary of additional header insertion into HTTP headers. Only used with HTTP
	// and `TERMINATED_HTTPS` protocols.
	InsertHeaders any `json:"insert_headers,required"`
	// Load balancer ID
	LoadBalancerID string `json:"load_balancer_id,required" format:"uuid4"`
	// Load balancer listener name
	Name string `json:"name,required"`
	// Listener operating status
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// Number of pools (for UI)
	PoolCount int64 `json:"pool_count,required"`
	// Load balancer protocol
	//
	// Any of "HTTP", "HTTPS", "PROMETHEUS", "TCP", "TERMINATED_HTTPS", "UDP".
	Protocol LbListenerProtocol `json:"protocol,required"`
	// Protocol port
	ProtocolPort int64 `json:"protocol_port,required"`
	// Listener lifecycle status
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// ID of the secret where PKCS12 file is stored for `TERMINATED_HTTPS` or
	// PROMETHEUS load balancer
	SecretID string `json:"secret_id,required"`
	// List of secret's ID containing PKCS12 format certificate/key bundles for
	// `TERMINATED_HTTPS` or PROMETHEUS listeners
	SniSecretID []string `json:"sni_secret_id,required"`
	// Statistics of the load balancer. It is available only in get functions by a
	// flag.
	Stats LoadBalancerStatistics `json:"stats,required"`
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
	// Load balancer listener users list
	UserList []LoadBalancerListenerDetailUserList `json:"user_list,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AllowedCidrs         respjson.Field
		ConnectionLimit      respjson.Field
		CreatorTaskID        respjson.Field
		InsertHeaders        respjson.Field
		LoadBalancerID       respjson.Field
		Name                 respjson.Field
		OperatingStatus      respjson.Field
		PoolCount            respjson.Field
		Protocol             respjson.Field
		ProtocolPort         respjson.Field
		ProvisioningStatus   respjson.Field
		SecretID             respjson.Field
		SniSecretID          respjson.Field
		Stats                respjson.Field
		TaskID               respjson.Field
		TimeoutClientData    respjson.Field
		TimeoutMemberConnect respjson.Field
		TimeoutMemberData    respjson.Field
		UserList             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerListenerDetail) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerListenerDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerListenerDetailUserList struct {
	// Encrypted password to auth via Basic Authentication
	EncryptedPassword string `json:"encrypted_password,required"`
	// Username to auth via Basic Authentication
	Username string `json:"username,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EncryptedPassword respjson.Field
		Username          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerListenerDetailUserList) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerListenerDetailUserList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerListenerList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []LoadBalancerListenerDetail `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerListenerList) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerListenerList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMetrics struct {
	// CPU utilization, % (max 100% for multi-core)
	CPUUtil float64 `json:"cpu_util,nullable"`
	// RAM utilization, %
	MemoryUtil float64 `json:"memory_util,nullable"`
	// Network out, bytes per second
	NetworkBpsEgress float64 `json:"network_Bps_egress,nullable"`
	// Network in, bytes per second
	NetworkBpsIngress float64 `json:"network_Bps_ingress,nullable"`
	// Network out, packets per second
	NetworkPpsEgress float64 `json:"network_pps_egress,nullable"`
	// Network in, packets per second
	NetworkPpsIngress float64 `json:"network_pps_ingress,nullable"`
	// Timestamp
	Time string `json:"time,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPUUtil           respjson.Field
		MemoryUtil        respjson.Field
		NetworkBpsEgress  respjson.Field
		NetworkBpsIngress respjson.Field
		NetworkPpsEgress  respjson.Field
		NetworkPpsIngress respjson.Field
		Time              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerMetrics) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMetricsList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []LoadBalancerMetrics `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerMetricsList) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerMetricsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPool struct {
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
	Listeners []LoadBalancerPoolListener `json:"listeners,required"`
	// Load balancers IDs
	Loadbalancers []LoadBalancerPoolLoadbalancer `json:"loadbalancers,required"`
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
func (r LoadBalancerPool) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerPool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolListener struct {
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
func (r LoadBalancerPoolListener) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerPoolListener) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolLoadbalancer struct {
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
func (r LoadBalancerPoolLoadbalancer) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerPoolLoadbalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []LoadBalancerPoolListResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerPoolList) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerPoolList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolListResult struct {
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
	Listeners []LoadBalancerPoolListResultListener `json:"listeners,required"`
	// Load balancers IDs
	Loadbalancers []LoadBalancerPoolListResultLoadbalancer `json:"loadbalancers,required"`
	// Pool members
	Members []LoadBalancerPoolListResultMemberUnion `json:"members,required"`
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
func (r LoadBalancerPoolListResult) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerPoolListResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolListResultListener struct {
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
func (r LoadBalancerPoolListResultListener) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerPoolListResultListener) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolListResultLoadbalancer struct {
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
func (r LoadBalancerPoolListResultLoadbalancer) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerPoolListResultLoadbalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LoadBalancerPoolListResultMemberUnion contains all possible properties and
// values from [LoadBalancerPoolListResultMemberLbPoolMemberSerializer], [Member].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LoadBalancerPoolListResultMemberUnion struct {
	ID string `json:"id"`
	// This field is from variant [Member].
	Address string `json:"address"`
	// This field is from variant [Member].
	AdminStateUp bool `json:"admin_state_up"`
	// This field is from variant [Member].
	Backup bool `json:"backup"`
	// This field is from variant [Member].
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status"`
	// This field is from variant [Member].
	ProtocolPort int64 `json:"protocol_port"`
	// This field is from variant [Member].
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status"`
	// This field is from variant [Member].
	SubnetID string `json:"subnet_id"`
	// This field is from variant [Member].
	Weight int64 `json:"weight"`
	// This field is from variant [Member].
	MonitorAddress string `json:"monitor_address"`
	// This field is from variant [Member].
	MonitorPort int64 `json:"monitor_port"`
	JSON        struct {
		ID                 respjson.Field
		Address            respjson.Field
		AdminStateUp       respjson.Field
		Backup             respjson.Field
		OperatingStatus    respjson.Field
		ProtocolPort       respjson.Field
		ProvisioningStatus respjson.Field
		SubnetID           respjson.Field
		Weight             respjson.Field
		MonitorAddress     respjson.Field
		MonitorPort        respjson.Field
		raw                string
	} `json:"-"`
}

func (u LoadBalancerPoolListResultMemberUnion) AsLbPoolMemberSerializer() (v LoadBalancerPoolListResultMemberLbPoolMemberSerializer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LoadBalancerPoolListResultMemberUnion) AsDetailedLbPoolMemberSerializer() (v Member) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LoadBalancerPoolListResultMemberUnion) RawJSON() string { return u.JSON.raw }

func (r *LoadBalancerPoolListResultMemberUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolListResultMemberLbPoolMemberSerializer struct {
	// Member ID must be provided if an existing member is being updated
	ID string `json:"id,required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerPoolListResultMemberLbPoolMemberSerializer) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerPoolListResultMemberLbPoolMemberSerializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerStatus struct {
	// UUID of the entity
	ID string `json:"id,required" format:"uuid"`
	// Listeners of the Load Balancer
	Listeners []ListenerStatus `json:"listeners,required"`
	// Name of the load balancer
	Name string `json:"name,required"`
	// Operating status of the entity
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// Provisioning status of the entity
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Listeners          respjson.Field
		Name               respjson.Field
		OperatingStatus    respjson.Field
		ProvisioningStatus respjson.Field
		Tags               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerStatus) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerStatusList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []LoadBalancerStatus `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerStatusList) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerStatusList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Member struct {
	// Member ID must be provided if an existing member is being updated
	ID string `json:"id,required" format:"uuid4"`
	// Member IP address
	Address string `json:"address,required" format:"ipvanyaddress"`
	// Administrative state of the resource. When set to true, the resource is enabled
	// and operational. When set to false, the resource is disabled and will not
	// process traffic. When null is passed, the value is skipped and defaults to true.
	AdminStateUp bool `json:"admin_state_up,required"`
	// Set to true if the member is a backup member, to which traffic will be sent
	// exclusively when all non-backup members will be unreachable. It allows to
	// realize ACTIVE-BACKUP load balancing without thinking about VRRP and VIP
	// configuration. Default is false
	Backup bool `json:"backup,required"`
	// Member operating status of the entity
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// Member IP port
	ProtocolPort int64 `json:"protocol_port,required"`
	// Pool member lifecycle status
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// `subnet_id` in which `address` is present.
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
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
	Weight int64 `json:"weight,required"`
	// An alternate IP address used for health monitoring of a backend member. Default
	// is null which monitors the member address.
	MonitorAddress string `json:"monitor_address,nullable" format:"ipvanyaddress"`
	// An alternate protocol port used for health monitoring of a backend member.
	// Default is null which monitors the member `protocol_port`.
	MonitorPort int64 `json:"monitor_port,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Address            respjson.Field
		AdminStateUp       respjson.Field
		Backup             respjson.Field
		OperatingStatus    respjson.Field
		ProtocolPort       respjson.Field
		ProvisioningStatus respjson.Field
		SubnetID           respjson.Field
		Weight             respjson.Field
		MonitorAddress     respjson.Field
		MonitorPort        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Member) RawJSON() string { return r.JSON.raw }
func (r *Member) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemberStatus struct {
	// UUID of the entity
	ID string `json:"id,required" format:"uuid"`
	// Address of the member (server)
	Address string `json:"address,required" format:"ipvanyaddress"`
	// Operating status of the entity
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// Port of the member (server)
	ProtocolPort int64 `json:"protocol_port,required"`
	// Provisioning status of the entity
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Address            respjson.Field
		OperatingStatus    respjson.Field
		ProtocolPort       respjson.Field
		ProvisioningStatus respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemberStatus) RawJSON() string { return r.JSON.raw }
func (r *MemberStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PoolStatus struct {
	// UUID of the entity
	ID string `json:"id,required" format:"uuid"`
	// Members (servers) of the pool
	Members []MemberStatus `json:"members,required"`
	// Name of the load balancer pool
	Name string `json:"name,required"`
	// Operating status of the entity
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// Provisioning status of the entity
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Health Monitor of the Pool
	HealthMonitor HealthMonitorStatus `json:"health_monitor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Members            respjson.Field
		Name               respjson.Field
		OperatingStatus    respjson.Field
		ProvisioningStatus respjson.Field
		HealthMonitor      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PoolStatus) RawJSON() string { return r.JSON.raw }
func (r *PoolStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionPersistence struct {
	// Session persistence type
	//
	// Any of "APP_COOKIE", "HTTP_COOKIE", "SOURCE_IP".
	Type LbSessionPersistenceType `json:"type,required"`
	// Should be set if app cookie or http cookie is used
	CookieName string `json:"cookie_name,nullable"`
	// Subnet mask if `source_ip` is used. For UDP ports only
	PersistenceGranularity string `json:"persistence_granularity,nullable"`
	// Session persistence timeout. For UDP ports only
	PersistenceTimeout int64 `json:"persistence_timeout,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                   respjson.Field
		CookieName             respjson.Field
		PersistenceGranularity respjson.Field
		PersistenceTimeout     respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionPersistence) RawJSON() string { return r.JSON.raw }
func (r *SessionPersistence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Load balancer flavor name
	Flavor param.Opt[string] `json:"flavor,omitzero"`
	// Load balancer name. Either `name` or `name_template` should be specified.
	Name param.Opt[string] `json:"name,omitzero"`
	// Load balancer name which will be changed by template. Either `name` or
	// `name_template` should be specified.
	NameTemplate param.Opt[string] `json:"name_template,omitzero"`
	// Network ID for load balancer. If not specified, default external network will be
	// used. Mutually exclusive with `vip_port_id`
	VipNetworkID param.Opt[string] `json:"vip_network_id,omitzero" format:"uuid4"`
	// Existing Reserved Fixed IP port ID for load balancer. Mutually exclusive with
	// `vip_network_id`
	VipPortID param.Opt[string] `json:"vip_port_id,omitzero" format:"uuid4"`
	// Subnet ID for load balancer. If not specified, any subnet from `vip_network_id`
	// will be selected. Ignored when `vip_network_id` is not specified.
	VipSubnetID param.Opt[string] `json:"vip_subnet_id,omitzero" format:"uuid4"`
	// Floating IP configuration for assignment
	FloatingIP LoadBalancerNewParamsFloatingIPUnion `json:"floating_ip,omitzero"`
	// Load balancer listeners. Maximum 50 per LB (excluding Prometheus endpoint
	// listener).
	Listeners []LoadBalancerNewParamsListener `json:"listeners,omitzero"`
	// Logging configuration
	Logging LoadBalancerNewParamsLogging `json:"logging,omitzero"`
	// Preferred option to establish connectivity between load balancer and its pools
	// members. L2 provides best performance, L3 provides less IPs usage. It is taking
	// effect only if `instance_id` + `ip_address` is provided, not `subnet_id` +
	// `ip_address`, because we're considering this as intentional `subnet_id`
	// specification.
	//
	// Any of "L2", "L3".
	PreferredConnectivity LoadBalancerMemberConnectivity `json:"preferred_connectivity,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Some tags are read-only and cannot be modified by the
	// user. Tags are also integrated with cost reports, allowing cost data to be
	// filtered based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	// IP family for load balancer subnet auto-selection if `vip_network_id` is
	// specified
	//
	// Any of "dual", "ipv4", "ipv6".
	VipIPFamily InterfaceIPFamily `json:"vip_ip_family,omitzero"`
	paramObj
}

func (r LoadBalancerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LoadBalancerNewParamsFloatingIPUnion struct {
	OfNew      *LoadBalancerNewParamsFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *LoadBalancerNewParamsFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

func (u LoadBalancerNewParamsFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNew, u.OfExisting)
}
func (u *LoadBalancerNewParamsFloatingIPUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LoadBalancerNewParamsFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LoadBalancerNewParamsFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LoadBalancerNewParamsFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[LoadBalancerNewParamsFloatingIPUnion](
		"source",
		apijson.Discriminator[LoadBalancerNewParamsFloatingIPNew]("new"),
		apijson.Discriminator[LoadBalancerNewParamsFloatingIPExisting]("existing"),
	)
}

func NewLoadBalancerNewParamsFloatingIPNew() LoadBalancerNewParamsFloatingIPNew {
	return LoadBalancerNewParamsFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewLoadBalancerNewParamsFloatingIPNew].
type LoadBalancerNewParamsFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source,required"`
	paramObj
}

func (r LoadBalancerNewParamsFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerNewParamsFloatingIPNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ExistingFloatingID, Source are required.
type LoadBalancerNewParamsFloatingIPExisting struct {
	// An existing available floating IP id must be specified if the source is set to
	// `existing`
	ExistingFloatingID string `json:"existing_floating_id,required" format:"uuid4"`
	// An existing available floating IP will be attached to the instance. A floating
	// IP is a public IP that makes the instance accessible from the internet, even if
	// it only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Source constant.Existing `json:"source,required"`
	paramObj
}

func (r LoadBalancerNewParamsFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerNewParamsFloatingIPExisting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Protocol, ProtocolPort are required.
type LoadBalancerNewParamsListener struct {
	// Load balancer listener name
	Name string `json:"name,required"`
	// Load balancer listener protocol
	//
	// Any of "HTTP", "HTTPS", "PROMETHEUS", "TCP", "TERMINATED_HTTPS", "UDP".
	Protocol LbListenerProtocol `json:"protocol,omitzero,required"`
	// Protocol port
	ProtocolPort int64 `json:"protocol_port,required"`
	// Frontend client inactivity timeout in milliseconds
	TimeoutClientData param.Opt[int64] `json:"timeout_client_data,omitzero"`
	// Backend member connection timeout in milliseconds
	TimeoutMemberConnect param.Opt[int64] `json:"timeout_member_connect,omitzero"`
	// Backend member inactivity timeout in milliseconds
	TimeoutMemberData param.Opt[int64] `json:"timeout_member_data,omitzero"`
	// Limit of the simultaneous connections. If -1 is provided, it is translated to
	// the default value 100000.
	ConnectionLimit param.Opt[int64] `json:"connection_limit,omitzero"`
	// Add headers X-Forwarded-For, X-Forwarded-Port, X-Forwarded-Proto to requests.
	// Only used with HTTP or `TERMINATED_HTTPS` protocols.
	InsertXForwarded param.Opt[bool] `json:"insert_x_forwarded,omitzero"`
	// ID of the secret where PKCS12 file is stored for `TERMINATED_HTTPS` or
	// PROMETHEUS listener
	SecretID param.Opt[string] `json:"secret_id,omitzero"`
	// Network CIDRs from which service will be accessible
	AllowedCidrs []string `json:"allowed_cidrs,omitzero" format:"ipvanynetwork"`
	// Member pools
	Pools []LoadBalancerNewParamsListenerPool `json:"pools,omitzero"`
	// List of secrets IDs containing PKCS12 format certificate/key bundles for
	// `TERMINATED_HTTPS` or PROMETHEUS listeners
	SniSecretID []string `json:"sni_secret_id,omitzero" format:"uuid4"`
	// Load balancer listener list of username and encrypted password items
	UserList []LoadBalancerNewParamsListenerUserList `json:"user_list,omitzero"`
	paramObj
}

func (r LoadBalancerNewParamsListener) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsListener
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerNewParamsListener) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties LbAlgorithm, Name, Protocol are required.
type LoadBalancerNewParamsListenerPool struct {
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
	Healthmonitor LoadBalancerNewParamsListenerPoolHealthmonitor `json:"healthmonitor,omitzero"`
	// Session persistence details
	SessionPersistence LoadBalancerNewParamsListenerPoolSessionPersistence `json:"session_persistence,omitzero"`
	// Pool members
	Members []LoadBalancerNewParamsListenerPoolMember `json:"members,omitzero"`
	paramObj
}

func (r LoadBalancerNewParamsListenerPool) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsListenerPool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerNewParamsListenerPool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Health monitor details
//
// The properties Delay, MaxRetries, Timeout, Type are required.
type LoadBalancerNewParamsListenerPoolHealthmonitor struct {
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

func (r LoadBalancerNewParamsListenerPoolHealthmonitor) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsListenerPoolHealthmonitor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerNewParamsListenerPoolHealthmonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Address, ProtocolPort are required.
type LoadBalancerNewParamsListenerPoolMember struct {
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

func (r LoadBalancerNewParamsListenerPoolMember) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsListenerPoolMember
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerNewParamsListenerPoolMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Session persistence details
//
// The property Type is required.
type LoadBalancerNewParamsListenerPoolSessionPersistence struct {
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

func (r LoadBalancerNewParamsListenerPoolSessionPersistence) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsListenerPoolSessionPersistence
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerNewParamsListenerPoolSessionPersistence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EncryptedPassword, Username are required.
type LoadBalancerNewParamsListenerUserList struct {
	// Encrypted password to auth via Basic Authentication
	EncryptedPassword string `json:"encrypted_password,required"`
	// Username to auth via Basic Authentication
	Username string `json:"username,required"`
	paramObj
}

func (r LoadBalancerNewParamsListenerUserList) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsListenerUserList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerNewParamsListenerUserList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Logging configuration
type LoadBalancerNewParamsLogging struct {
	// Destination region id to which the logs will be written
	DestinationRegionID param.Opt[int64] `json:"destination_region_id,omitzero"`
	// The topic name to which the logs will be written
	TopicName param.Opt[string] `json:"topic_name,omitzero"`
	// Enable/disable forwarding logs to LaaS
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// The logs retention policy
	RetentionPolicy LaasIndexRetentionPolicyParam `json:"retention_policy,omitzero"`
	paramObj
}

func (r LoadBalancerNewParamsLogging) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsLogging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerNewParamsLogging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerUpdateParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Logging configuration
	Logging LoadBalancerUpdateParamsLogging `json:"logging,omitzero"`
	// Preferred option to establish connectivity between load balancer and its pools
	// members
	//
	// Any of "L2", "L3".
	PreferredConnectivity LoadBalancerMemberConnectivity `json:"preferred_connectivity,omitzero"`
	// Update key-value tags using JSON Merge Patch semantics (RFC 7386). Provide
	// key-value pairs to add or update tags. Set tag values to `null` to remove tags.
	// Unspecified tags remain unchanged. Read-only tags are always preserved and
	// cannot be modified.
	//
	// **Examples:**
	//
	//   - **Add/update tags:**
	//     `{'tags': {'environment': 'production', 'team': 'backend'}}` adds new tags or
	//     updates existing ones.
	//
	// - **Delete tags:** `{'tags': {'old_tag': null}}` removes specific tags.
	//
	//   - **Remove all tags:** `{'tags': null}` removes all user-managed tags (read-only
	//     tags are preserved).
	//
	//   - **Partial update:** `{'tags': {'environment': 'staging'}}` only updates
	//     specified tags.
	//
	//   - **Mixed operations:**
	//     `{'tags': {'environment': 'production', 'cost_center': 'engineering', 'deprecated_tag': null}}`
	//     adds/updates 'environment' and '`cost_center`' while removing
	//     '`deprecated_tag`', preserving other existing tags.
	//
	//   - **Replace all:** first delete existing tags with null values, then add new
	//     ones in the same request.
	Tags TagUpdateMap `json:"tags,omitzero"`
	paramObj
}

func (r LoadBalancerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Logging configuration
type LoadBalancerUpdateParamsLogging struct {
	// Destination region id to which the logs will be written
	DestinationRegionID param.Opt[int64] `json:"destination_region_id,omitzero"`
	// The topic name to which the logs will be written
	TopicName param.Opt[string] `json:"topic_name,omitzero"`
	// Enable/disable forwarding logs to LaaS
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// The logs retention policy
	RetentionPolicy LaasIndexRetentionPolicyParam `json:"retention_policy,omitzero"`
	paramObj
}

func (r LoadBalancerUpdateParamsLogging) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerUpdateParamsLogging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerUpdateParamsLogging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// With or without assigned floating IP
	AssignedFloating param.Opt[bool] `query:"assigned_floating,omitzero" json:"-"`
	// Limit the number of returned limit request entities.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// With or without logging
	LoggingEnabled param.Opt[bool] `query:"logging_enabled,omitzero" json:"-"`
	// Filter by name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Offset value is used to exclude the first set of records from the result.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Ordering Load Balancer list result by name, `created_at`, `updated_at`,
	// `operating_status`, `provisioning_status`, `vip_address`, `vip_ip_family` and
	// flavor fields of the load balancer and directions (name.asc), default is
	// "`created_at`.asc"
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// Show statistics
	ShowStats param.Opt[bool] `query:"show_stats,omitzero" json:"-"`
	// Filter by tag key-value pairs. Must be a valid JSON string.
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// Show Advanced DDoS protection profile, if exists
	WithDDOS param.Opt[bool] `query:"with_ddos,omitzero" json:"-"`
	// Filter by tag keys.
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LoadBalancerListParams]'s query parameters as `url.Values`.
func (r LoadBalancerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LoadBalancerDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type LoadBalancerFailoverParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Validate current load balancer status before failover or not.
	Force param.Opt[bool] `json:"force,omitzero"`
	paramObj
}

func (r LoadBalancerFailoverParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerFailoverParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerFailoverParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerGetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Show statistics
	ShowStats param.Opt[bool] `query:"show_stats,omitzero" json:"-"`
	// Show DDoS profile
	WithDDOS param.Opt[bool] `query:"with_ddos,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LoadBalancerGetParams]'s query parameters as `url.Values`.
func (r LoadBalancerGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LoadBalancerResizeParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Name of the desired flavor to resize to.
	Flavor string `json:"flavor,required"`
	paramObj
}

func (r LoadBalancerResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerResizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
