// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
	"github.com/stainless-sdks/gcore-go/shared/constant"
	"github.com/tidwall/gjson"
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
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Rename load balancer, activate/deactivate logs or update preferred connectivity
// for load balancer
func (r *LoadBalancerService) Update(ctx context.Context, loadbalancerID string, params LoadBalancerUpdateParams, opts ...option.RequestOption) (res *LoadBalancer, err error) {
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
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List load balancers
func (r *LoadBalancerService) List(ctx context.Context, params LoadBalancerListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LoadBalancer], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
func (r *LoadBalancerService) Delete(ctx context.Context, loadbalancerID string, body LoadBalancerDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Failover loadbalancer
func (r *LoadBalancerService) Failover(ctx context.Context, loadbalancerID string, params LoadBalancerFailoverParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/failover", params.ProjectID.Value, params.RegionID.Value, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get load balancer
func (r *LoadBalancerService) Get(ctx context.Context, loadbalancerID string, params LoadBalancerGetParams, opts ...option.RequestOption) (res *LoadBalancer, err error) {
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
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Resize loadbalancer
func (r *LoadBalancerService) Resize(ctx context.Context, loadbalancerID string, params LoadBalancerResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/resize", params.ProjectID.Value, params.RegionID.Value, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// '#/components/schemas/DetailedLbPoolSerializer'
// "$.components.schemas.DetailedLbPoolSerializer"
type DetailedLbPool struct {
	// '#/components/schemas/DetailedLbPoolSerializer/properties/id'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/ca_secret_id/anyOf/0'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.ca_secret_id.anyOf[0]"
	CaSecretID string `json:"ca_secret_id,required" format:"uuid4"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/crl_secret_id/anyOf/0'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.crl_secret_id.anyOf[0]"
	CrlSecretID string `json:"crl_secret_id,required" format:"uuid4"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/lb_algorithm'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.lb_algorithm"
	//
	// Any of "LEAST_CONNECTIONS", "ROUND_ROBIN", "SOURCE_IP".
	LbAlgorithm LbAlgorithm `json:"lb_algorithm,required"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/listeners'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.listeners"
	Listeners []DetailedLbPoolListener `json:"listeners,required"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/loadbalancers'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.loadbalancers"
	Loadbalancers []DetailedLbPoolLoadbalancer `json:"loadbalancers,required"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/members'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.members"
	Members []DetailedLbPoolMember `json:"members,required"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/name'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/operating_status'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.operating_status"
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/protocol'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.protocol"
	//
	// Any of "HTTP", "HTTPS", "PROXY", "PROXYV2", "TCP", "UDP".
	Protocol LbPoolProtocol `json:"protocol,required"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/provisioning_status'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.provisioning_status"
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/secret_id/anyOf/0'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.secret_id.anyOf[0]"
	SecretID string `json:"secret_id,required" format:"uuid4"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/session_persistence/anyOf/0'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.session_persistence.anyOf[0]"
	SessionPersistence LbSessionPersistence `json:"session_persistence,required"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/timeout_client_data/anyOf/0'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.timeout_client_data.anyOf[0]"
	TimeoutClientData int64 `json:"timeout_client_data,required"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/timeout_member_connect/anyOf/0'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.timeout_member_connect.anyOf[0]"
	TimeoutMemberConnect int64 `json:"timeout_member_connect,required"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/timeout_member_data/anyOf/0'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.timeout_member_data.anyOf[0]"
	TimeoutMemberData int64 `json:"timeout_member_data,required"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/healthmonitor/anyOf/0'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.healthmonitor.anyOf[0]"
	Healthmonitor LbHealthMonitor `json:"healthmonitor,nullable"`
	// '#/components/schemas/DetailedLbPoolSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.DetailedLbPoolSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                   resp.Field
		CaSecretID           resp.Field
		CrlSecretID          resp.Field
		LbAlgorithm          resp.Field
		Listeners            resp.Field
		Loadbalancers        resp.Field
		Members              resp.Field
		Name                 resp.Field
		OperatingStatus      resp.Field
		Protocol             resp.Field
		ProvisioningStatus   resp.Field
		SecretID             resp.Field
		SessionPersistence   resp.Field
		TimeoutClientData    resp.Field
		TimeoutMemberConnect resp.Field
		TimeoutMemberData    resp.Field
		CreatorTaskID        resp.Field
		Healthmonitor        resp.Field
		TaskID               resp.Field
		ExtraFields          map[string]resp.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DetailedLbPool) RawJSON() string { return r.JSON.raw }
func (r *DetailedLbPool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/DetailedLbPoolSerializer/properties/listeners/items'
// "$.components.schemas.DetailedLbPoolSerializer.properties.listeners.items"
type DetailedLbPoolListener struct {
	// '#/components/schemas/MandatoryIdSerializerPydantic/properties/id'
	// "$.components.schemas.MandatoryIdSerializerPydantic.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DetailedLbPoolListener) RawJSON() string { return r.JSON.raw }
func (r *DetailedLbPoolListener) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/DetailedLbPoolSerializer/properties/loadbalancers/items'
// "$.components.schemas.DetailedLbPoolSerializer.properties.loadbalancers.items"
type DetailedLbPoolLoadbalancer struct {
	// '#/components/schemas/MandatoryIdSerializerPydantic/properties/id'
	// "$.components.schemas.MandatoryIdSerializerPydantic.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DetailedLbPoolLoadbalancer) RawJSON() string { return r.JSON.raw }
func (r *DetailedLbPoolLoadbalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/DetailedLbPoolSerializerList'
// "$.components.schemas.DetailedLbPoolSerializerList"
type DetailedLbPoolList struct {
	// '#/components/schemas/DetailedLbPoolSerializerList/properties/count'
	// "$.components.schemas.DetailedLbPoolSerializerList.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/DetailedLbPoolSerializerList/properties/results'
	// "$.components.schemas.DetailedLbPoolSerializerList.properties.results"
	Results []DetailedLbPool `json:"results,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DetailedLbPoolList) RawJSON() string { return r.JSON.raw }
func (r *DetailedLbPoolList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/DetailedLbPoolMemberSerializer'
// "$.components.schemas.DetailedLbPoolMemberSerializer"
type DetailedLbPoolMember struct {
	// '#/components/schemas/DetailedLbPoolMemberSerializer/properties/id'
	// "$.components.schemas.DetailedLbPoolMemberSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/DetailedLbPoolMemberSerializer/properties/address'
	// "$.components.schemas.DetailedLbPoolMemberSerializer.properties.address"
	Address string `json:"address,required" format:"ipvanyaddress"`
	// '#/components/schemas/DetailedLbPoolMemberSerializer/properties/admin_state_up'
	// "$.components.schemas.DetailedLbPoolMemberSerializer.properties.admin_state_up"
	AdminStateUp bool `json:"admin_state_up,required"`
	// '#/components/schemas/DetailedLbPoolMemberSerializer/properties/operating_status'
	// "$.components.schemas.DetailedLbPoolMemberSerializer.properties.operating_status"
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// '#/components/schemas/DetailedLbPoolMemberSerializer/properties/protocol_port'
	// "$.components.schemas.DetailedLbPoolMemberSerializer.properties.protocol_port"
	ProtocolPort int64 `json:"protocol_port,required"`
	// '#/components/schemas/DetailedLbPoolMemberSerializer/properties/provisioning_status'
	// "$.components.schemas.DetailedLbPoolMemberSerializer.properties.provisioning_status"
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// '#/components/schemas/DetailedLbPoolMemberSerializer/properties/weight'
	// "$.components.schemas.DetailedLbPoolMemberSerializer.properties.weight"
	Weight int64 `json:"weight,required"`
	// '#/components/schemas/DetailedLbPoolMemberSerializer/properties/monitor_address/anyOf/0'
	// "$.components.schemas.DetailedLbPoolMemberSerializer.properties.monitor_address.anyOf[0]"
	MonitorAddress string `json:"monitor_address,nullable" format:"ipvanyaddress"`
	// '#/components/schemas/DetailedLbPoolMemberSerializer/properties/monitor_port/anyOf/0'
	// "$.components.schemas.DetailedLbPoolMemberSerializer.properties.monitor_port.anyOf[0]"
	MonitorPort int64 `json:"monitor_port,nullable"`
	// '#/components/schemas/DetailedLbPoolMemberSerializer/properties/subnet_id/anyOf/0'
	// "$.components.schemas.DetailedLbPoolMemberSerializer.properties.subnet_id.anyOf[0]"
	SubnetID string `json:"subnet_id,nullable" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                 resp.Field
		Address            resp.Field
		AdminStateUp       resp.Field
		OperatingStatus    resp.Field
		ProtocolPort       resp.Field
		ProvisioningStatus resp.Field
		Weight             resp.Field
		MonitorAddress     resp.Field
		MonitorPort        resp.Field
		SubnetID           resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DetailedLbPoolMember) RawJSON() string { return r.JSON.raw }
func (r *DetailedLbPoolMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/HealthMonitorStatusSerializer'
// "$.components.schemas.HealthMonitorStatusSerializer"
type HealthMonitorStatus struct {
	// '#/components/schemas/HealthMonitorStatusSerializer/properties/id'
	// "$.components.schemas.HealthMonitorStatusSerializer.properties.id"
	ID string `json:"id,required" format:"uuid"`
	// '#/components/schemas/HealthMonitorStatusSerializer/properties/operating_status'
	// "$.components.schemas.HealthMonitorStatusSerializer.properties.operating_status"
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// '#/components/schemas/HealthMonitorStatusSerializer/properties/provisioning_status'
	// "$.components.schemas.HealthMonitorStatusSerializer.properties.provisioning_status"
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// '#/components/schemas/HealthMonitorStatusSerializer/properties/type'
	// "$.components.schemas.HealthMonitorStatusSerializer.properties.type"
	//
	// Any of "HTTP", "HTTPS", "K8S", "PING", "TCP", "TLS-HELLO", "UDP-CONNECT".
	Type HealthMonitorType `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                 resp.Field
		OperatingStatus    resp.Field
		ProvisioningStatus resp.Field
		Type               resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthMonitorStatus) RawJSON() string { return r.JSON.raw }
func (r *HealthMonitorStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/HealthMonitorTypeEnum'
// "$.components.schemas.HealthMonitorTypeEnum"
type HealthMonitorType string

const (
	HealthMonitorTypeHTTP       HealthMonitorType = "HTTP"
	HealthMonitorTypeHTTPS      HealthMonitorType = "HTTPS"
	HealthMonitorTypeK8S        HealthMonitorType = "K8S"
	HealthMonitorTypePing       HealthMonitorType = "PING"
	HealthMonitorTypeTcp        HealthMonitorType = "TCP"
	HealthMonitorTypeTlsHello   HealthMonitorType = "TLS-HELLO"
	HealthMonitorTypeUdpConnect HealthMonitorType = "UDP-CONNECT"
)

// '#/components/schemas/HttpMethodEnum' "$.components.schemas.HttpMethodEnum"
type HTTPMethod string

const (
	HTTPMethodConnect HTTPMethod = "CONNECT"
	HTTPMethodDelete  HTTPMethod = "DELETE"
	HTTPMethodGet     HTTPMethod = "GET"
	HTTPMethodHead    HTTPMethod = "HEAD"
	HTTPMethodOptions HTTPMethod = "OPTIONS"
	HTTPMethodPatch   HTTPMethod = "PATCH"
	HTTPMethodPost    HTTPMethod = "POST"
	HTTPMethodPut     HTTPMethod = "PUT"
	HTTPMethodTrace   HTTPMethod = "TRACE"
)

// '#/components/schemas/L7PolicySchema' "$.components.schemas.L7PolicySchema"
type L7Policy struct {
	// '#/components/schemas/L7PolicySchema/properties/id'
	// "$.components.schemas.L7PolicySchema.properties.id"
	ID string `json:"id"`
	// '#/components/schemas/L7PolicySchema/properties/action'
	// "$.components.schemas.L7PolicySchema.properties.action"
	//
	// Any of "REDIRECT_PREFIX", "REDIRECT_TO_POOL", "REDIRECT_TO_URL", "REJECT".
	Action L7PolicyAction `json:"action"`
	// '#/components/schemas/L7PolicySchema/properties/listener_id'
	// "$.components.schemas.L7PolicySchema.properties.listener_id"
	ListenerID string `json:"listener_id"`
	// '#/components/schemas/L7PolicySchema/properties/name'
	// "$.components.schemas.L7PolicySchema.properties.name"
	Name string `json:"name"`
	// '#/components/schemas/L7PolicySchema/properties/operating_status'
	// "$.components.schemas.L7PolicySchema.properties.operating_status"
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus L7PolicyOperatingStatus `json:"operating_status"`
	// '#/components/schemas/L7PolicySchema/properties/position'
	// "$.components.schemas.L7PolicySchema.properties.position"
	Position int64 `json:"position"`
	// '#/components/schemas/L7PolicySchema/properties/project_id'
	// "$.components.schemas.L7PolicySchema.properties.project_id"
	ProjectID int64 `json:"project_id"`
	// '#/components/schemas/L7PolicySchema/properties/provisioning_status'
	// "$.components.schemas.L7PolicySchema.properties.provisioning_status"
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus L7PolicyProvisioningStatus `json:"provisioning_status"`
	// '#/components/schemas/L7PolicySchema/properties/redirect_http_code'
	// "$.components.schemas.L7PolicySchema.properties.redirect_http_code"
	RedirectHTTPCode int64 `json:"redirect_http_code"`
	// '#/components/schemas/L7PolicySchema/properties/redirect_pool_id'
	// "$.components.schemas.L7PolicySchema.properties.redirect_pool_id"
	RedirectPoolID string `json:"redirect_pool_id"`
	// '#/components/schemas/L7PolicySchema/properties/redirect_prefix'
	// "$.components.schemas.L7PolicySchema.properties.redirect_prefix"
	RedirectPrefix string `json:"redirect_prefix"`
	// '#/components/schemas/L7PolicySchema/properties/redirect_url'
	// "$.components.schemas.L7PolicySchema.properties.redirect_url"
	RedirectURL string `json:"redirect_url"`
	// '#/components/schemas/L7PolicySchema/properties/region'
	// "$.components.schemas.L7PolicySchema.properties.region"
	Region string `json:"region"`
	// '#/components/schemas/L7PolicySchema/properties/region_id'
	// "$.components.schemas.L7PolicySchema.properties.region_id"
	RegionID int64 `json:"region_id"`
	// '#/components/schemas/L7PolicySchema/properties/rules'
	// "$.components.schemas.L7PolicySchema.properties.rules"
	Rules []L7Rule `json:"rules"`
	// '#/components/schemas/L7PolicySchema/properties/tags'
	// "$.components.schemas.L7PolicySchema.properties.tags"
	Tags []string `json:"tags"`
	// '#/components/schemas/L7PolicySchema/properties/task_id'
	// "$.components.schemas.L7PolicySchema.properties.task_id"
	TaskID string `json:"task_id"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                 resp.Field
		Action             resp.Field
		ListenerID         resp.Field
		Name               resp.Field
		OperatingStatus    resp.Field
		Position           resp.Field
		ProjectID          resp.Field
		ProvisioningStatus resp.Field
		RedirectHTTPCode   resp.Field
		RedirectPoolID     resp.Field
		RedirectPrefix     resp.Field
		RedirectURL        resp.Field
		Region             resp.Field
		RegionID           resp.Field
		Rules              resp.Field
		Tags               resp.Field
		TaskID             resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r L7Policy) RawJSON() string { return r.JSON.raw }
func (r *L7Policy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/L7PolicySchema/properties/action'
// "$.components.schemas.L7PolicySchema.properties.action"
type L7PolicyAction string

const (
	L7PolicyActionRedirectPrefix L7PolicyAction = "REDIRECT_PREFIX"
	L7PolicyActionRedirectToPool L7PolicyAction = "REDIRECT_TO_POOL"
	L7PolicyActionRedirectToURL  L7PolicyAction = "REDIRECT_TO_URL"
	L7PolicyActionReject         L7PolicyAction = "REJECT"
)

// '#/components/schemas/L7PolicySchema/properties/operating_status'
// "$.components.schemas.L7PolicySchema.properties.operating_status"
type L7PolicyOperatingStatus string

const (
	L7PolicyOperatingStatusDegraded  L7PolicyOperatingStatus = "DEGRADED"
	L7PolicyOperatingStatusDraining  L7PolicyOperatingStatus = "DRAINING"
	L7PolicyOperatingStatusError     L7PolicyOperatingStatus = "ERROR"
	L7PolicyOperatingStatusNoMonitor L7PolicyOperatingStatus = "NO_MONITOR"
	L7PolicyOperatingStatusOffline   L7PolicyOperatingStatus = "OFFLINE"
	L7PolicyOperatingStatusOnline    L7PolicyOperatingStatus = "ONLINE"
)

// '#/components/schemas/L7PolicySchema/properties/provisioning_status'
// "$.components.schemas.L7PolicySchema.properties.provisioning_status"
type L7PolicyProvisioningStatus string

const (
	L7PolicyProvisioningStatusActive        L7PolicyProvisioningStatus = "ACTIVE"
	L7PolicyProvisioningStatusDeleted       L7PolicyProvisioningStatus = "DELETED"
	L7PolicyProvisioningStatusError         L7PolicyProvisioningStatus = "ERROR"
	L7PolicyProvisioningStatusPendingCreate L7PolicyProvisioningStatus = "PENDING_CREATE"
	L7PolicyProvisioningStatusPendingDelete L7PolicyProvisioningStatus = "PENDING_DELETE"
	L7PolicyProvisioningStatusPendingUpdate L7PolicyProvisioningStatus = "PENDING_UPDATE"
)

// '#/components/schemas/L7PolicyListSchema'
// "$.components.schemas.L7PolicyListSchema"
type L7PolicyList struct {
	// '#/components/schemas/L7PolicyListSchema/properties/count'
	// "$.components.schemas.L7PolicyListSchema.properties.count"
	Count int64 `json:"count"`
	// '#/components/schemas/L7PolicyListSchema/properties/results'
	// "$.components.schemas.L7PolicyListSchema.properties.results"
	Results []L7Policy `json:"results"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r L7PolicyList) RawJSON() string { return r.JSON.raw }
func (r *L7PolicyList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/L7RuleSchema' "$.components.schemas.L7RuleSchema"
type L7Rule struct {
	// '#/components/schemas/L7RuleSchema/properties/id'
	// "$.components.schemas.L7RuleSchema.properties.id"
	ID string `json:"id"`
	// '#/components/schemas/L7RuleSchema/properties/compare_type'
	// "$.components.schemas.L7RuleSchema.properties.compare_type"
	//
	// Any of "CONTAINS", "ENDS_WITH", "EQUAL_TO", "REGEX", "STARTS_WITH".
	CompareType L7RuleCompareType `json:"compare_type"`
	// '#/components/schemas/L7RuleSchema/properties/invert'
	// "$.components.schemas.L7RuleSchema.properties.invert"
	Invert bool `json:"invert"`
	// '#/components/schemas/L7RuleSchema/properties/key'
	// "$.components.schemas.L7RuleSchema.properties.key"
	Key string `json:"key"`
	// '#/components/schemas/L7RuleSchema/properties/operating_status'
	// "$.components.schemas.L7RuleSchema.properties.operating_status"
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus L7RuleOperatingStatus `json:"operating_status"`
	// '#/components/schemas/L7RuleSchema/properties/project_id'
	// "$.components.schemas.L7RuleSchema.properties.project_id"
	ProjectID int64 `json:"project_id"`
	// '#/components/schemas/L7RuleSchema/properties/provisioning_status'
	// "$.components.schemas.L7RuleSchema.properties.provisioning_status"
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus L7RuleProvisioningStatus `json:"provisioning_status"`
	// '#/components/schemas/L7RuleSchema/properties/region'
	// "$.components.schemas.L7RuleSchema.properties.region"
	Region string `json:"region"`
	// '#/components/schemas/L7RuleSchema/properties/region_id'
	// "$.components.schemas.L7RuleSchema.properties.region_id"
	RegionID int64 `json:"region_id"`
	// '#/components/schemas/L7RuleSchema/properties/tags'
	// "$.components.schemas.L7RuleSchema.properties.tags"
	Tags []string `json:"tags"`
	// '#/components/schemas/L7RuleSchema/properties/task_id'
	// "$.components.schemas.L7RuleSchema.properties.task_id"
	TaskID string `json:"task_id"`
	// '#/components/schemas/L7RuleSchema/properties/type'
	// "$.components.schemas.L7RuleSchema.properties.type"
	//
	// Any of "COOKIE", "FILE_TYPE", "HEADER", "HOST_NAME", "PATH",
	// "SSL_CONN_HAS_CERT", "SSL_DN_FIELD", "SSL_VERIFY_RESULT".
	Type L7RuleType `json:"type"`
	// '#/components/schemas/L7RuleSchema/properties/value'
	// "$.components.schemas.L7RuleSchema.properties.value"
	Value string `json:"value"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                 resp.Field
		CompareType        resp.Field
		Invert             resp.Field
		Key                resp.Field
		OperatingStatus    resp.Field
		ProjectID          resp.Field
		ProvisioningStatus resp.Field
		Region             resp.Field
		RegionID           resp.Field
		Tags               resp.Field
		TaskID             resp.Field
		Type               resp.Field
		Value              resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r L7Rule) RawJSON() string { return r.JSON.raw }
func (r *L7Rule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/L7RuleSchema/properties/compare_type'
// "$.components.schemas.L7RuleSchema.properties.compare_type"
type L7RuleCompareType string

const (
	L7RuleCompareTypeContains   L7RuleCompareType = "CONTAINS"
	L7RuleCompareTypeEndsWith   L7RuleCompareType = "ENDS_WITH"
	L7RuleCompareTypeEqualTo    L7RuleCompareType = "EQUAL_TO"
	L7RuleCompareTypeRegex      L7RuleCompareType = "REGEX"
	L7RuleCompareTypeStartsWith L7RuleCompareType = "STARTS_WITH"
)

// '#/components/schemas/L7RuleSchema/properties/operating_status'
// "$.components.schemas.L7RuleSchema.properties.operating_status"
type L7RuleOperatingStatus string

const (
	L7RuleOperatingStatusDegraded  L7RuleOperatingStatus = "DEGRADED"
	L7RuleOperatingStatusDraining  L7RuleOperatingStatus = "DRAINING"
	L7RuleOperatingStatusError     L7RuleOperatingStatus = "ERROR"
	L7RuleOperatingStatusNoMonitor L7RuleOperatingStatus = "NO_MONITOR"
	L7RuleOperatingStatusOffline   L7RuleOperatingStatus = "OFFLINE"
	L7RuleOperatingStatusOnline    L7RuleOperatingStatus = "ONLINE"
)

// '#/components/schemas/L7RuleSchema/properties/provisioning_status'
// "$.components.schemas.L7RuleSchema.properties.provisioning_status"
type L7RuleProvisioningStatus string

const (
	L7RuleProvisioningStatusActive        L7RuleProvisioningStatus = "ACTIVE"
	L7RuleProvisioningStatusDeleted       L7RuleProvisioningStatus = "DELETED"
	L7RuleProvisioningStatusError         L7RuleProvisioningStatus = "ERROR"
	L7RuleProvisioningStatusPendingCreate L7RuleProvisioningStatus = "PENDING_CREATE"
	L7RuleProvisioningStatusPendingDelete L7RuleProvisioningStatus = "PENDING_DELETE"
	L7RuleProvisioningStatusPendingUpdate L7RuleProvisioningStatus = "PENDING_UPDATE"
)

// '#/components/schemas/L7RuleSchema/properties/type'
// "$.components.schemas.L7RuleSchema.properties.type"
type L7RuleType string

const (
	L7RuleTypeCookie          L7RuleType = "COOKIE"
	L7RuleTypeFileType        L7RuleType = "FILE_TYPE"
	L7RuleTypeHeader          L7RuleType = "HEADER"
	L7RuleTypeHostName        L7RuleType = "HOST_NAME"
	L7RuleTypePath            L7RuleType = "PATH"
	L7RuleTypeSslConnHasCert  L7RuleType = "SSL_CONN_HAS_CERT"
	L7RuleTypeSslDnField      L7RuleType = "SSL_DN_FIELD"
	L7RuleTypeSslVerifyResult L7RuleType = "SSL_VERIFY_RESULT"
)

// '#/components/schemas/L7RuleListSchema' "$.components.schemas.L7RuleListSchema"
type L7RuleList struct {
	// '#/components/schemas/L7RuleListSchema/properties/count'
	// "$.components.schemas.L7RuleListSchema.properties.count"
	Count int64 `json:"count"`
	// '#/components/schemas/L7RuleListSchema/properties/results'
	// "$.components.schemas.L7RuleListSchema.properties.results"
	Results []L7Rule `json:"results"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r L7RuleList) RawJSON() string { return r.JSON.raw }
func (r *L7RuleList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LbAlgorithmEnum' "$.components.schemas.LbAlgorithmEnum"
type LbAlgorithm string

const (
	LbAlgorithmLeastConnections LbAlgorithm = "LEAST_CONNECTIONS"
	LbAlgorithmRoundRobin       LbAlgorithm = "ROUND_ROBIN"
	LbAlgorithmSourceIP         LbAlgorithm = "SOURCE_IP"
)

// '#/components/schemas/LbFlavorPricingCollectionSerializer'
// "$.components.schemas.LbFlavorPricingCollectionSerializer"
type LbFlavorList struct {
	// '#/components/schemas/LbFlavorPricingCollectionSerializer/properties/count'
	// "$.components.schemas.LbFlavorPricingCollectionSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/LbFlavorPricingCollectionSerializer/properties/results'
	// "$.components.schemas.LbFlavorPricingCollectionSerializer.properties.results"
	Results []LbFlavorListResult `json:"results,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LbFlavorList) RawJSON() string { return r.JSON.raw }
func (r *LbFlavorList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LbFlavorPricingCollectionSerializer/properties/results/items'
// "$.components.schemas.LbFlavorPricingCollectionSerializer.properties.results.items"
type LbFlavorListResult struct {
	// '#/components/schemas/LbFlavorPricingSerializer/properties/flavor_id'
	// "$.components.schemas.LbFlavorPricingSerializer.properties.flavor_id"
	FlavorID string `json:"flavor_id,required"`
	// '#/components/schemas/LbFlavorPricingSerializer/properties/flavor_name'
	// "$.components.schemas.LbFlavorPricingSerializer.properties.flavor_name"
	FlavorName string `json:"flavor_name,required"`
	// '#/components/schemas/LbFlavorPricingSerializer/properties/hardware_description'
	// "$.components.schemas.LbFlavorPricingSerializer.properties.hardware_description"
	HardwareDescription FlavorHardwareDescription `json:"hardware_description,required"`
	// '#/components/schemas/LbFlavorPricingSerializer/properties/ram'
	// "$.components.schemas.LbFlavorPricingSerializer.properties.ram"
	Ram int64 `json:"ram,required"`
	// '#/components/schemas/LbFlavorPricingSerializer/properties/vcpus'
	// "$.components.schemas.LbFlavorPricingSerializer.properties.vcpus"
	Vcpus int64 `json:"vcpus,required"`
	// '#/components/schemas/LbFlavorPricingSerializer/properties/currency_code/anyOf/0'
	// "$.components.schemas.LbFlavorPricingSerializer.properties.currency_code.anyOf[0]"
	CurrencyCode string `json:"currency_code,nullable"`
	// '#/components/schemas/LbFlavorPricingSerializer/properties/price_per_hour/anyOf/0'
	// "$.components.schemas.LbFlavorPricingSerializer.properties.price_per_hour.anyOf[0]"
	PricePerHour float64 `json:"price_per_hour,nullable"`
	// '#/components/schemas/LbFlavorPricingSerializer/properties/price_per_month/anyOf/0'
	// "$.components.schemas.LbFlavorPricingSerializer.properties.price_per_month.anyOf[0]"
	PricePerMonth float64 `json:"price_per_month,nullable"`
	// '#/components/schemas/LbFlavorPricingSerializer/properties/price_status/anyOf/0'
	// "$.components.schemas.LbFlavorPricingSerializer.properties.price_status.anyOf[0]"
	//
	// Any of "error", "hide", "show".
	PriceStatus string `json:"price_status,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		FlavorID            resp.Field
		FlavorName          resp.Field
		HardwareDescription resp.Field
		Ram                 resp.Field
		Vcpus               resp.Field
		CurrencyCode        resp.Field
		PricePerHour        resp.Field
		PricePerMonth       resp.Field
		PriceStatus         resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LbFlavorListResult) RawJSON() string { return r.JSON.raw }
func (r *LbFlavorListResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LbHealthMonitorSerializer'
// "$.components.schemas.LbHealthMonitorSerializer"
type LbHealthMonitor struct {
	// '#/components/schemas/LbHealthMonitorSerializer/properties/id'
	// "$.components.schemas.LbHealthMonitorSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/LbHealthMonitorSerializer/properties/admin_state_up'
	// "$.components.schemas.LbHealthMonitorSerializer.properties.admin_state_up"
	AdminStateUp bool `json:"admin_state_up,required"`
	// '#/components/schemas/LbHealthMonitorSerializer/properties/delay'
	// "$.components.schemas.LbHealthMonitorSerializer.properties.delay"
	Delay int64 `json:"delay,required"`
	// '#/components/schemas/LbHealthMonitorSerializer/properties/max_retries'
	// "$.components.schemas.LbHealthMonitorSerializer.properties.max_retries"
	MaxRetries int64 `json:"max_retries,required"`
	// '#/components/schemas/LbHealthMonitorSerializer/properties/max_retries_down'
	// "$.components.schemas.LbHealthMonitorSerializer.properties.max_retries_down"
	MaxRetriesDown int64 `json:"max_retries_down,required"`
	// '#/components/schemas/LbHealthMonitorSerializer/properties/operating_status'
	// "$.components.schemas.LbHealthMonitorSerializer.properties.operating_status"
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// '#/components/schemas/LbHealthMonitorSerializer/properties/provisioning_status'
	// "$.components.schemas.LbHealthMonitorSerializer.properties.provisioning_status"
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// '#/components/schemas/LbHealthMonitorSerializer/properties/timeout'
	// "$.components.schemas.LbHealthMonitorSerializer.properties.timeout"
	Timeout int64 `json:"timeout,required"`
	// '#/components/schemas/LbHealthMonitorSerializer/properties/type'
	// "$.components.schemas.LbHealthMonitorSerializer.properties.type"
	//
	// Any of "HTTP", "HTTPS", "K8S", "PING", "TCP", "TLS-HELLO", "UDP-CONNECT".
	Type HealthMonitorType `json:"type,required"`
	// '#/components/schemas/LbHealthMonitorSerializer/properties/expected_codes/anyOf/0'
	// "$.components.schemas.LbHealthMonitorSerializer.properties.expected_codes.anyOf[0]"
	ExpectedCodes string `json:"expected_codes,nullable"`
	// '#/components/schemas/LbHealthMonitorSerializer/properties/http_method/anyOf/0'
	// "$.components.schemas.LbHealthMonitorSerializer.properties.http_method.anyOf[0]"
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod HTTPMethod `json:"http_method,nullable"`
	// '#/components/schemas/LbHealthMonitorSerializer/properties/url_path/anyOf/0'
	// "$.components.schemas.LbHealthMonitorSerializer.properties.url_path.anyOf[0]"
	URLPath string `json:"url_path,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                 resp.Field
		AdminStateUp       resp.Field
		Delay              resp.Field
		MaxRetries         resp.Field
		MaxRetriesDown     resp.Field
		OperatingStatus    resp.Field
		ProvisioningStatus resp.Field
		Timeout            resp.Field
		Type               resp.Field
		ExpectedCodes      resp.Field
		HTTPMethod         resp.Field
		URLPath            resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LbHealthMonitor) RawJSON() string { return r.JSON.raw }
func (r *LbHealthMonitor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LbListenerSerializer'
// "$.components.schemas.LbListenerSerializer"
type LbListener struct {
	// '#/components/schemas/LbListenerSerializer/properties/id'
	// "$.components.schemas.LbListenerSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/LbListenerSerializer/properties/connection_limit'
	// "$.components.schemas.LbListenerSerializer.properties.connection_limit"
	ConnectionLimit int64 `json:"connection_limit,required"`
	// '#/components/schemas/LbListenerSerializer/properties/insert_headers'
	// "$.components.schemas.LbListenerSerializer.properties.insert_headers"
	InsertHeaders any `json:"insert_headers,required"`
	// '#/components/schemas/LbListenerSerializer/properties/name'
	// "$.components.schemas.LbListenerSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/LbListenerSerializer/properties/operating_status'
	// "$.components.schemas.LbListenerSerializer.properties.operating_status"
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// '#/components/schemas/LbListenerSerializer/properties/protocol'
	// "$.components.schemas.LbListenerSerializer.properties.protocol"
	//
	// Any of "HTTP", "HTTPS", "PROMETHEUS", "TCP", "TERMINATED_HTTPS", "UDP".
	Protocol LbListenerProtocol `json:"protocol,required"`
	// '#/components/schemas/LbListenerSerializer/properties/protocol_port'
	// "$.components.schemas.LbListenerSerializer.properties.protocol_port"
	ProtocolPort int64 `json:"protocol_port,required"`
	// '#/components/schemas/LbListenerSerializer/properties/provisioning_status'
	// "$.components.schemas.LbListenerSerializer.properties.provisioning_status"
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// '#/components/schemas/LbListenerSerializer/properties/allowed_cidrs/anyOf/0'
	// "$.components.schemas.LbListenerSerializer.properties.allowed_cidrs.anyOf[0]"
	AllowedCidrs []string `json:"allowed_cidrs,nullable" format:"ipvanynetwork"`
	// '#/components/schemas/LbListenerSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.LbListenerSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// '#/components/schemas/LbListenerSerializer/properties/loadbalancer_id/anyOf/0'
	// "$.components.schemas.LbListenerSerializer.properties.loadbalancer_id.anyOf[0]"
	LoadbalancerID string `json:"loadbalancer_id,nullable" format:"uuid4"`
	// '#/components/schemas/LbListenerSerializer/properties/pool_count/anyOf/0'
	// "$.components.schemas.LbListenerSerializer.properties.pool_count.anyOf[0]"
	PoolCount int64 `json:"pool_count,nullable"`
	// '#/components/schemas/LbListenerSerializer/properties/secret_id/anyOf/0'
	// "$.components.schemas.LbListenerSerializer.properties.secret_id.anyOf[0]"
	SecretID string `json:"secret_id,nullable"`
	// '#/components/schemas/LbListenerSerializer/properties/sni_secret_id/anyOf/0'
	// "$.components.schemas.LbListenerSerializer.properties.sni_secret_id.anyOf[0]"
	SniSecretID []string `json:"sni_secret_id,nullable"`
	// '#/components/schemas/LbListenerSerializer/properties/stats/anyOf/0'
	// "$.components.schemas.LbListenerSerializer.properties.stats.anyOf[0]"
	Stats LoadBalancerStatistics `json:"stats,nullable"`
	// '#/components/schemas/LbListenerSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.LbListenerSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// '#/components/schemas/LbListenerSerializer/properties/timeout_client_data/anyOf/0'
	// "$.components.schemas.LbListenerSerializer.properties.timeout_client_data.anyOf[0]"
	TimeoutClientData int64 `json:"timeout_client_data,nullable"`
	// '#/components/schemas/LbListenerSerializer/properties/timeout_member_connect/anyOf/0'
	// "$.components.schemas.LbListenerSerializer.properties.timeout_member_connect.anyOf[0]"
	TimeoutMemberConnect int64 `json:"timeout_member_connect,nullable"`
	// '#/components/schemas/LbListenerSerializer/properties/timeout_member_data/anyOf/0'
	// "$.components.schemas.LbListenerSerializer.properties.timeout_member_data.anyOf[0]"
	TimeoutMemberData int64 `json:"timeout_member_data,nullable"`
	// '#/components/schemas/LbListenerSerializer/properties/user_list'
	// "$.components.schemas.LbListenerSerializer.properties.user_list"
	UserList []LbListenerUserList `json:"user_list"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                   resp.Field
		ConnectionLimit      resp.Field
		InsertHeaders        resp.Field
		Name                 resp.Field
		OperatingStatus      resp.Field
		Protocol             resp.Field
		ProtocolPort         resp.Field
		ProvisioningStatus   resp.Field
		AllowedCidrs         resp.Field
		CreatorTaskID        resp.Field
		LoadbalancerID       resp.Field
		PoolCount            resp.Field
		SecretID             resp.Field
		SniSecretID          resp.Field
		Stats                resp.Field
		TaskID               resp.Field
		TimeoutClientData    resp.Field
		TimeoutMemberConnect resp.Field
		TimeoutMemberData    resp.Field
		UserList             resp.Field
		ExtraFields          map[string]resp.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LbListener) RawJSON() string { return r.JSON.raw }
func (r *LbListener) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LbListenerSerializer/properties/user_list/items'
// "$.components.schemas.LbListenerSerializer.properties.user_list.items"
type LbListenerUserList struct {
	// '#/components/schemas/UserListItem/properties/encrypted_password'
	// "$.components.schemas.UserListItem.properties.encrypted_password"
	EncryptedPassword string `json:"encrypted_password,required"`
	// '#/components/schemas/UserListItem/properties/username'
	// "$.components.schemas.UserListItem.properties.username"
	Username string `json:"username,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		EncryptedPassword resp.Field
		Username          resp.Field
		ExtraFields       map[string]resp.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LbListenerUserList) RawJSON() string { return r.JSON.raw }
func (r *LbListenerUserList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LbListenerSerializerList'
// "$.components.schemas.LbListenerSerializerList"
type LbListenerList struct {
	// '#/components/schemas/LbListenerSerializerList/properties/count'
	// "$.components.schemas.LbListenerSerializerList.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/LbListenerSerializerList/properties/results'
	// "$.components.schemas.LbListenerSerializerList.properties.results"
	Results []LbListener `json:"results,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LbListenerList) RawJSON() string { return r.JSON.raw }
func (r *LbListenerList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LbListenerProtocolEnum'
// "$.components.schemas.LbListenerProtocolEnum"
type LbListenerProtocol string

const (
	LbListenerProtocolHTTP            LbListenerProtocol = "HTTP"
	LbListenerProtocolHTTPS           LbListenerProtocol = "HTTPS"
	LbListenerProtocolPrometheus      LbListenerProtocol = "PROMETHEUS"
	LbListenerProtocolTcp             LbListenerProtocol = "TCP"
	LbListenerProtocolTerminatedHTTPS LbListenerProtocol = "TERMINATED_HTTPS"
	LbListenerProtocolUdp             LbListenerProtocol = "UDP"
)

// '#/components/schemas/LbPoolProtocolEnum'
// "$.components.schemas.LbPoolProtocolEnum"
type LbPoolProtocol string

const (
	LbPoolProtocolHTTP    LbPoolProtocol = "HTTP"
	LbPoolProtocolHTTPS   LbPoolProtocol = "HTTPS"
	LbPoolProtocolProxy   LbPoolProtocol = "PROXY"
	LbPoolProtocolProxyv2 LbPoolProtocol = "PROXYV2"
	LbPoolProtocolTcp     LbPoolProtocol = "TCP"
	LbPoolProtocolUdp     LbPoolProtocol = "UDP"
)

// '#/components/schemas/LbSessionPersistence'
// "$.components.schemas.LbSessionPersistence"
type LbSessionPersistence struct {
	// '#/components/schemas/LbSessionPersistence/properties/type'
	// "$.components.schemas.LbSessionPersistence.properties.type"
	//
	// Any of "APP_COOKIE", "HTTP_COOKIE", "SOURCE_IP".
	Type SessionPersistenceType `json:"type,required"`
	// '#/components/schemas/LbSessionPersistence/properties/cookie_name/anyOf/0'
	// "$.components.schemas.LbSessionPersistence.properties.cookie_name.anyOf[0]"
	CookieName string `json:"cookie_name,nullable"`
	// '#/components/schemas/LbSessionPersistence/properties/persistence_granularity/anyOf/0'
	// "$.components.schemas.LbSessionPersistence.properties.persistence_granularity.anyOf[0]"
	PersistenceGranularity string `json:"persistence_granularity,nullable"`
	// '#/components/schemas/LbSessionPersistence/properties/persistence_timeout/anyOf/0'
	// "$.components.schemas.LbSessionPersistence.properties.persistence_timeout.anyOf[0]"
	PersistenceTimeout int64 `json:"persistence_timeout,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Type                   resp.Field
		CookieName             resp.Field
		PersistenceGranularity resp.Field
		PersistenceTimeout     resp.Field
		ExtraFields            map[string]resp.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LbSessionPersistence) RawJSON() string { return r.JSON.raw }
func (r *LbSessionPersistence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ListenerStatusSerializer'
// "$.components.schemas.ListenerStatusSerializer"
type ListenerStatus struct {
	// '#/components/schemas/ListenerStatusSerializer/properties/id'
	// "$.components.schemas.ListenerStatusSerializer.properties.id"
	ID string `json:"id,required" format:"uuid"`
	// '#/components/schemas/ListenerStatusSerializer/properties/name'
	// "$.components.schemas.ListenerStatusSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/ListenerStatusSerializer/properties/operating_status'
	// "$.components.schemas.ListenerStatusSerializer.properties.operating_status"
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// '#/components/schemas/ListenerStatusSerializer/properties/pools'
	// "$.components.schemas.ListenerStatusSerializer.properties.pools"
	Pools []PoolStatus `json:"pools,required"`
	// '#/components/schemas/ListenerStatusSerializer/properties/provisioning_status'
	// "$.components.schemas.ListenerStatusSerializer.properties.provisioning_status"
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                 resp.Field
		Name               resp.Field
		OperatingStatus    resp.Field
		Pools              resp.Field
		ProvisioningStatus resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListenerStatus) RawJSON() string { return r.JSON.raw }
func (r *ListenerStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LoadBalancerStatusSerializer'
// "$.components.schemas.LoadBalancerStatusSerializer"
type LoadBalancerStatus struct {
	// '#/components/schemas/LoadBalancerStatusSerializer/properties/id'
	// "$.components.schemas.LoadBalancerStatusSerializer.properties.id"
	ID string `json:"id,required" format:"uuid"`
	// '#/components/schemas/LoadBalancerStatusSerializer/properties/listeners'
	// "$.components.schemas.LoadBalancerStatusSerializer.properties.listeners"
	Listeners []ListenerStatus `json:"listeners,required"`
	// '#/components/schemas/LoadBalancerStatusSerializer/properties/name'
	// "$.components.schemas.LoadBalancerStatusSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/LoadBalancerStatusSerializer/properties/operating_status'
	// "$.components.schemas.LoadBalancerStatusSerializer.properties.operating_status"
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// '#/components/schemas/LoadBalancerStatusSerializer/properties/provisioning_status'
	// "$.components.schemas.LoadBalancerStatusSerializer.properties.provisioning_status"
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// '#/components/schemas/LoadBalancerStatusSerializer/properties/tags'
	// "$.components.schemas.LoadBalancerStatusSerializer.properties.tags"
	Tags []Tag `json:"tags"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                 resp.Field
		Listeners          resp.Field
		Name               resp.Field
		OperatingStatus    resp.Field
		ProvisioningStatus resp.Field
		Tags               resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerStatus) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LoadBalancerStatusSerializerList'
// "$.components.schemas.LoadBalancerStatusSerializerList"
type LoadBalancerStatusList struct {
	// '#/components/schemas/LoadBalancerStatusSerializerList/properties/count'
	// "$.components.schemas.LoadBalancerStatusSerializerList.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/LoadBalancerStatusSerializerList/properties/results'
	// "$.components.schemas.LoadBalancerStatusSerializerList.properties.results"
	Results []LoadBalancerStatus `json:"results,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerStatusList) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerStatusList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LoadbalancerMetricsSerializer'
// "$.components.schemas.LoadbalancerMetricsSerializer"
type LoadbalancerMetrics struct {
	// '#/components/schemas/LoadbalancerMetricsSerializer/properties/cpu_util/anyOf/0'
	// "$.components.schemas.LoadbalancerMetricsSerializer.properties.cpu_util.anyOf[0]"
	CPUUtil float64 `json:"cpu_util,nullable"`
	// '#/components/schemas/LoadbalancerMetricsSerializer/properties/memory_util/anyOf/0'
	// "$.components.schemas.LoadbalancerMetricsSerializer.properties.memory_util.anyOf[0]"
	MemoryUtil float64 `json:"memory_util,nullable"`
	// '#/components/schemas/LoadbalancerMetricsSerializer/properties/network_Bps_egress/anyOf/0'
	// "$.components.schemas.LoadbalancerMetricsSerializer.properties.network_Bps_egress.anyOf[0]"
	NetworkBpsEgress float64 `json:"network_Bps_egress,nullable"`
	// '#/components/schemas/LoadbalancerMetricsSerializer/properties/network_Bps_ingress/anyOf/0'
	// "$.components.schemas.LoadbalancerMetricsSerializer.properties.network_Bps_ingress.anyOf[0]"
	NetworkBpsIngress float64 `json:"network_Bps_ingress,nullable"`
	// '#/components/schemas/LoadbalancerMetricsSerializer/properties/network_pps_egress/anyOf/0'
	// "$.components.schemas.LoadbalancerMetricsSerializer.properties.network_pps_egress.anyOf[0]"
	NetworkPpsEgress float64 `json:"network_pps_egress,nullable"`
	// '#/components/schemas/LoadbalancerMetricsSerializer/properties/network_pps_ingress/anyOf/0'
	// "$.components.schemas.LoadbalancerMetricsSerializer.properties.network_pps_ingress.anyOf[0]"
	NetworkPpsIngress float64 `json:"network_pps_ingress,nullable"`
	// '#/components/schemas/LoadbalancerMetricsSerializer/properties/time/anyOf/0'
	// "$.components.schemas.LoadbalancerMetricsSerializer.properties.time.anyOf[0]"
	Time string `json:"time,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CPUUtil           resp.Field
		MemoryUtil        resp.Field
		NetworkBpsEgress  resp.Field
		NetworkBpsIngress resp.Field
		NetworkPpsEgress  resp.Field
		NetworkPpsIngress resp.Field
		Time              resp.Field
		ExtraFields       map[string]resp.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadbalancerMetrics) RawJSON() string { return r.JSON.raw }
func (r *LoadbalancerMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LoadbalancerMetricsSerializerList'
// "$.components.schemas.LoadbalancerMetricsSerializerList"
type LoadbalancerMetricsList struct {
	// '#/components/schemas/LoadbalancerMetricsSerializerList/properties/count'
	// "$.components.schemas.LoadbalancerMetricsSerializerList.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/LoadbalancerMetricsSerializerList/properties/results'
	// "$.components.schemas.LoadbalancerMetricsSerializerList.properties.results"
	Results []LoadbalancerMetrics `json:"results,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadbalancerMetricsList) RawJSON() string { return r.JSON.raw }
func (r *LoadbalancerMetricsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/MemberStatusSerializer'
// "$.components.schemas.MemberStatusSerializer"
type MemberStatus struct {
	// '#/components/schemas/MemberStatusSerializer/properties/id'
	// "$.components.schemas.MemberStatusSerializer.properties.id"
	ID string `json:"id,required" format:"uuid"`
	// '#/components/schemas/MemberStatusSerializer/properties/address'
	// "$.components.schemas.MemberStatusSerializer.properties.address"
	Address string `json:"address,required" format:"ipvanyaddress"`
	// '#/components/schemas/MemberStatusSerializer/properties/operating_status'
	// "$.components.schemas.MemberStatusSerializer.properties.operating_status"
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// '#/components/schemas/MemberStatusSerializer/properties/protocol_port'
	// "$.components.schemas.MemberStatusSerializer.properties.protocol_port"
	ProtocolPort int64 `json:"protocol_port,required"`
	// '#/components/schemas/MemberStatusSerializer/properties/provisioning_status'
	// "$.components.schemas.MemberStatusSerializer.properties.provisioning_status"
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                 resp.Field
		Address            resp.Field
		OperatingStatus    resp.Field
		ProtocolPort       resp.Field
		ProvisioningStatus resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemberStatus) RawJSON() string { return r.JSON.raw }
func (r *MemberStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/PoolStatusSerializer'
// "$.components.schemas.PoolStatusSerializer"
type PoolStatus struct {
	// '#/components/schemas/PoolStatusSerializer/properties/id'
	// "$.components.schemas.PoolStatusSerializer.properties.id"
	ID string `json:"id,required" format:"uuid"`
	// '#/components/schemas/PoolStatusSerializer/properties/members'
	// "$.components.schemas.PoolStatusSerializer.properties.members"
	Members []MemberStatus `json:"members,required"`
	// '#/components/schemas/PoolStatusSerializer/properties/name'
	// "$.components.schemas.PoolStatusSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/PoolStatusSerializer/properties/operating_status'
	// "$.components.schemas.PoolStatusSerializer.properties.operating_status"
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// '#/components/schemas/PoolStatusSerializer/properties/provisioning_status'
	// "$.components.schemas.PoolStatusSerializer.properties.provisioning_status"
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// '#/components/schemas/PoolStatusSerializer/properties/health_monitor'
	// "$.components.schemas.PoolStatusSerializer.properties.health_monitor"
	HealthMonitor HealthMonitorStatus `json:"health_monitor"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                 resp.Field
		Members            resp.Field
		Name               resp.Field
		OperatingStatus    resp.Field
		ProvisioningStatus resp.Field
		HealthMonitor      resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PoolStatus) RawJSON() string { return r.JSON.raw }
func (r *PoolStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/SessionPersistenceTypeEnum'
// "$.components.schemas.SessionPersistenceTypeEnum"
type SessionPersistenceType string

const (
	SessionPersistenceTypeAppCookie  SessionPersistenceType = "APP_COOKIE"
	SessionPersistenceTypeHTTPCookie SessionPersistenceType = "HTTP_COOKIE"
	SessionPersistenceTypeSourceIP   SessionPersistenceType = "SOURCE_IP"
)

type LoadBalancerNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateLoadbalancerSerializer/properties/flavor'
	// "$.components.schemas.CreateLoadbalancerSerializer.properties.flavor"
	Flavor param.Opt[string] `json:"flavor,omitzero"`
	// '#/components/schemas/CreateLoadbalancerSerializer/properties/name'
	// "$.components.schemas.CreateLoadbalancerSerializer.properties.name"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/CreateLoadbalancerSerializer/properties/name_template'
	// "$.components.schemas.CreateLoadbalancerSerializer.properties.name_template"
	NameTemplate param.Opt[string] `json:"name_template,omitzero"`
	// '#/components/schemas/CreateLoadbalancerSerializer/properties/vip_network_id'
	// "$.components.schemas.CreateLoadbalancerSerializer.properties.vip_network_id"
	VipNetworkID param.Opt[string] `json:"vip_network_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLoadbalancerSerializer/properties/vip_port_id'
	// "$.components.schemas.CreateLoadbalancerSerializer.properties.vip_port_id"
	VipPortID param.Opt[string] `json:"vip_port_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLoadbalancerSerializer/properties/vip_subnet_id'
	// "$.components.schemas.CreateLoadbalancerSerializer.properties.vip_subnet_id"
	VipSubnetID param.Opt[string] `json:"vip_subnet_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLoadbalancerSerializer/properties/floating_ip'
	// "$.components.schemas.CreateLoadbalancerSerializer.properties.floating_ip"
	FloatingIP LoadBalancerNewParamsFloatingIPUnion `json:"floating_ip,omitzero"`
	// '#/components/schemas/CreateLoadbalancerSerializer/properties/listeners'
	// "$.components.schemas.CreateLoadbalancerSerializer.properties.listeners"
	Listeners []LoadBalancerNewParamsListener `json:"listeners,omitzero"`
	// '#/components/schemas/CreateLoadbalancerSerializer/properties/logging'
	// "$.components.schemas.CreateLoadbalancerSerializer.properties.logging"
	Logging LoadBalancerNewParamsLogging `json:"logging,omitzero"`
	// '#/components/schemas/CreateLoadbalancerSerializer/properties/preferred_connectivity'
	// "$.components.schemas.CreateLoadbalancerSerializer.properties.preferred_connectivity"
	//
	// Any of "L2", "L3".
	PreferredConnectivity LoadBalancerMemberConnectivity `json:"preferred_connectivity,omitzero"`
	// '#/components/schemas/CreateLoadbalancerSerializer/properties/tags'
	// "$.components.schemas.CreateLoadbalancerSerializer.properties.tags"
	Tags TagUpdateList `json:"tags,omitzero"`
	// '#/components/schemas/CreateLoadbalancerSerializer/properties/vip_ip_family'
	// "$.components.schemas.CreateLoadbalancerSerializer.properties.vip_ip_family"
	//
	// Any of "dual", "ipv4", "ipv6".
	VipIPFamily InterfaceIPFamily `json:"vip_ip_family,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LoadBalancerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LoadBalancerNewParamsFloatingIPUnion struct {
	OfNewFloatingIP      *LoadBalancerNewParamsFloatingIPNewInstanceFloatingIPInterfaceSerializer      `json:",omitzero,inline"`
	OfExistingFloatingIP *LoadBalancerNewParamsFloatingIPExistingInstanceFloatingIPInterfaceSerializer `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u LoadBalancerNewParamsFloatingIPUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u LoadBalancerNewParamsFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[LoadBalancerNewParamsFloatingIPUnion](u.OfNewFloatingIP, u.OfExistingFloatingIP)
}

func (u *LoadBalancerNewParamsFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNewFloatingIP) {
		return u.OfNewFloatingIP
	} else if !param.IsOmitted(u.OfExistingFloatingIP) {
		return u.OfExistingFloatingIP
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LoadBalancerNewParamsFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExistingFloatingIP; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LoadBalancerNewParamsFloatingIPUnion) GetSource() *string {
	if vt := u.OfNewFloatingIP; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExistingFloatingIP; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[LoadBalancerNewParamsFloatingIPUnion](
		"source",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LoadBalancerNewParamsFloatingIPNewInstanceFloatingIPInterfaceSerializer{}),
			DiscriminatorValue: "new",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LoadBalancerNewParamsFloatingIPExistingInstanceFloatingIPInterfaceSerializer{}),
			DiscriminatorValue: "existing",
		},
	)
}

func NewLoadBalancerNewParamsFloatingIPNewInstanceFloatingIPInterfaceSerializer() LoadBalancerNewParamsFloatingIPNewInstanceFloatingIPInterfaceSerializer {
	return LoadBalancerNewParamsFloatingIPNewInstanceFloatingIPInterfaceSerializer{
		Source: "new",
	}
}

// '#/components/schemas/CreateLoadbalancerSerializer/properties/floating_ip/anyOf/0'
// "$.components.schemas.CreateLoadbalancerSerializer.properties.floating_ip.anyOf[0]"
//
// This struct has a constant value, construct it with
// [NewLoadBalancerNewParamsFloatingIPNewInstanceFloatingIPInterfaceSerializer].
type LoadBalancerNewParamsFloatingIPNewInstanceFloatingIPInterfaceSerializer struct {
	// '#/components/schemas/NewInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.NewInstanceFloatingIpInterfaceSerializer.properties.source"
	Source constant.New `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerNewParamsFloatingIPNewInstanceFloatingIPInterfaceSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerNewParamsFloatingIPNewInstanceFloatingIPInterfaceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsFloatingIPNewInstanceFloatingIPInterfaceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateLoadbalancerSerializer/properties/floating_ip/anyOf/1'
// "$.components.schemas.CreateLoadbalancerSerializer.properties.floating_ip.anyOf[1]"
//
// The properties ExistingFloatingID, Source are required.
type LoadBalancerNewParamsFloatingIPExistingInstanceFloatingIPInterfaceSerializer struct {
	// '#/components/schemas/ExistingInstanceFloatingIpInterfaceSerializer/properties/existing_floating_id'
	// "$.components.schemas.ExistingInstanceFloatingIpInterfaceSerializer.properties.existing_floating_id"
	ExistingFloatingID string `json:"existing_floating_id,required" format:"uuid4"`
	// '#/components/schemas/ExistingInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.ExistingInstanceFloatingIpInterfaceSerializer.properties.source"
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Source constant.Existing `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerNewParamsFloatingIPExistingInstanceFloatingIPInterfaceSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerNewParamsFloatingIPExistingInstanceFloatingIPInterfaceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsFloatingIPExistingInstanceFloatingIPInterfaceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateLoadbalancerSerializer/properties/listeners/items'
// "$.components.schemas.CreateLoadbalancerSerializer.properties.listeners.items"
//
// The properties Name, Protocol, ProtocolPort are required.
type LoadBalancerNewParamsListener struct {
	// '#/components/schemas/CreateListenerSerializer/properties/name'
	// "$.components.schemas.CreateListenerSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateListenerSerializer/properties/protocol'
	// "$.components.schemas.CreateListenerSerializer.properties.protocol"
	//
	// Any of "HTTP", "HTTPS", "PROMETHEUS", "TCP", "TERMINATED_HTTPS", "UDP".
	Protocol LbListenerProtocol `json:"protocol,omitzero,required"`
	// '#/components/schemas/CreateListenerSerializer/properties/protocol_port'
	// "$.components.schemas.CreateListenerSerializer.properties.protocol_port"
	ProtocolPort int64 `json:"protocol_port,required"`
	// '#/components/schemas/CreateListenerSerializer/properties/timeout_client_data/anyOf/0'
	// "$.components.schemas.CreateListenerSerializer.properties.timeout_client_data.anyOf[0]"
	TimeoutClientData param.Opt[int64] `json:"timeout_client_data,omitzero"`
	// '#/components/schemas/CreateListenerSerializer/properties/timeout_member_connect/anyOf/0'
	// "$.components.schemas.CreateListenerSerializer.properties.timeout_member_connect.anyOf[0]"
	TimeoutMemberConnect param.Opt[int64] `json:"timeout_member_connect,omitzero"`
	// '#/components/schemas/CreateListenerSerializer/properties/timeout_member_data/anyOf/0'
	// "$.components.schemas.CreateListenerSerializer.properties.timeout_member_data.anyOf[0]"
	TimeoutMemberData param.Opt[int64] `json:"timeout_member_data,omitzero"`
	// '#/components/schemas/CreateListenerSerializer/properties/connection_limit'
	// "$.components.schemas.CreateListenerSerializer.properties.connection_limit"
	ConnectionLimit param.Opt[int64] `json:"connection_limit,omitzero"`
	// '#/components/schemas/CreateListenerSerializer/properties/insert_x_forwarded'
	// "$.components.schemas.CreateListenerSerializer.properties.insert_x_forwarded"
	InsertXForwarded param.Opt[bool] `json:"insert_x_forwarded,omitzero"`
	// '#/components/schemas/CreateListenerSerializer/properties/secret_id/anyOf/0'
	// "$.components.schemas.CreateListenerSerializer.properties.secret_id.anyOf[0]"
	SecretID param.Opt[string] `json:"secret_id,omitzero"`
	// '#/components/schemas/CreateListenerSerializer/properties/allowed_cidrs/anyOf/0'
	// "$.components.schemas.CreateListenerSerializer.properties.allowed_cidrs.anyOf[0]"
	AllowedCidrs []string `json:"allowed_cidrs,omitzero" format:"ipvanynetwork"`
	// '#/components/schemas/CreateListenerSerializer/properties/pools'
	// "$.components.schemas.CreateListenerSerializer.properties.pools"
	Pools []LoadBalancerNewParamsListenerPool `json:"pools,omitzero"`
	// '#/components/schemas/CreateListenerSerializer/properties/sni_secret_id'
	// "$.components.schemas.CreateListenerSerializer.properties.sni_secret_id"
	SniSecretID []string `json:"sni_secret_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateListenerSerializer/properties/user_list'
	// "$.components.schemas.CreateListenerSerializer.properties.user_list"
	UserList []LoadBalancerNewParamsListenerUserList `json:"user_list,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerNewParamsListener) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r LoadBalancerNewParamsListener) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsListener
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateListenerSerializer/properties/pools/items'
// "$.components.schemas.CreateListenerSerializer.properties.pools.items"
//
// The properties LbAlgorithm, Name, Protocol are required.
type LoadBalancerNewParamsListenerPool struct {
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
	Healthmonitor LoadBalancerNewParamsListenerPoolHealthmonitor `json:"healthmonitor,omitzero"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/members/anyOf/0'
	// "$.components.schemas.CreateLbPoolSerializer.properties.members.anyOf[0]"
	Members []LoadBalancerNewParamsListenerPoolMember `json:"members,omitzero"`
	// '#/components/schemas/CreateLbPoolSerializer/properties/session_persistence/anyOf/0'
	// "$.components.schemas.CreateLbPoolSerializer.properties.session_persistence.anyOf[0]"
	SessionPersistence LoadBalancerNewParamsListenerPoolSessionPersistence `json:"session_persistence,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerNewParamsListenerPool) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerNewParamsListenerPool) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsListenerPool
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateLbPoolSerializer/properties/healthmonitor/anyOf/0'
// "$.components.schemas.CreateLbPoolSerializer.properties.healthmonitor.anyOf[0]"
//
// The properties Delay, MaxRetries, Timeout, Type are required.
type LoadBalancerNewParamsListenerPoolHealthmonitor struct {
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
func (f LoadBalancerNewParamsListenerPoolHealthmonitor) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerNewParamsListenerPoolHealthmonitor) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsListenerPoolHealthmonitor
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateLbPoolSerializer/properties/members/anyOf/0/items'
// "$.components.schemas.CreateLbPoolSerializer.properties.members.anyOf[0].items"
//
// The properties Address, ProtocolPort are required.
type LoadBalancerNewParamsListenerPoolMember struct {
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
func (f LoadBalancerNewParamsListenerPoolMember) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerNewParamsListenerPoolMember) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsListenerPoolMember
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateLbPoolSerializer/properties/session_persistence/anyOf/0'
// "$.components.schemas.CreateLbPoolSerializer.properties.session_persistence.anyOf[0]"
//
// The property Type is required.
type LoadBalancerNewParamsListenerPoolSessionPersistence struct {
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
func (f LoadBalancerNewParamsListenerPoolSessionPersistence) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerNewParamsListenerPoolSessionPersistence) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsListenerPoolSessionPersistence
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateListenerSerializer/properties/user_list/items'
// "$.components.schemas.CreateListenerSerializer.properties.user_list.items"
//
// The properties EncryptedPassword, Username are required.
type LoadBalancerNewParamsListenerUserList struct {
	// '#/components/schemas/UserListItem/properties/encrypted_password'
	// "$.components.schemas.UserListItem.properties.encrypted_password"
	EncryptedPassword string `json:"encrypted_password,required"`
	// '#/components/schemas/UserListItem/properties/username'
	// "$.components.schemas.UserListItem.properties.username"
	Username string `json:"username,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerNewParamsListenerUserList) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerNewParamsListenerUserList) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsListenerUserList
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateLoadbalancerSerializer/properties/logging'
// "$.components.schemas.CreateLoadbalancerSerializer.properties.logging"
type LoadBalancerNewParamsLogging struct {
	// '#/components/schemas/LoadbalancerLoggingSerializer/properties/destination_region_id/anyOf/0'
	// "$.components.schemas.LoadbalancerLoggingSerializer.properties.destination_region_id.anyOf[0]"
	DestinationRegionID param.Opt[int64] `json:"destination_region_id,omitzero"`
	// '#/components/schemas/LoadbalancerLoggingSerializer/properties/topic_name/anyOf/0'
	// "$.components.schemas.LoadbalancerLoggingSerializer.properties.topic_name.anyOf[0]"
	TopicName param.Opt[string] `json:"topic_name,omitzero"`
	// '#/components/schemas/LoadbalancerLoggingSerializer/properties/enabled'
	// "$.components.schemas.LoadbalancerLoggingSerializer.properties.enabled"
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// '#/components/schemas/LoadbalancerLoggingSerializer/properties/retention_policy/anyOf/0'
	// "$.components.schemas.LoadbalancerLoggingSerializer.properties.retention_policy.anyOf[0]"
	RetentionPolicy LoadBalancerNewParamsLoggingRetentionPolicy `json:"retention_policy,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerNewParamsLogging) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r LoadBalancerNewParamsLogging) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsLogging
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/LoadbalancerLoggingSerializer/properties/retention_policy/anyOf/0'
// "$.components.schemas.LoadbalancerLoggingSerializer.properties.retention_policy.anyOf[0]"
//
// The property Period is required.
type LoadBalancerNewParamsLoggingRetentionPolicy struct {
	// '#/components/schemas/LaasIndexRetentionPolicyPydanticSerializer/properties/period/anyOf/0'
	// "$.components.schemas.LaasIndexRetentionPolicyPydanticSerializer.properties.period.anyOf[0]"
	Period param.Opt[int64] `json:"period,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerNewParamsLoggingRetentionPolicy) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerNewParamsLoggingRetentionPolicy) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerNewParamsLoggingRetentionPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}

type LoadBalancerUpdateParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bloadbalancer_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/{loadbalancer_id}'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bloadbalancer_id%7D/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/{loadbalancer_id}'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/LoadBalancerPatchSerializer/properties/name'
	// "$.components.schemas.LoadBalancerPatchSerializer.properties.name"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/LoadBalancerPatchSerializer/properties/logging'
	// "$.components.schemas.LoadBalancerPatchSerializer.properties.logging"
	Logging LoadBalancerUpdateParamsLogging `json:"logging,omitzero"`
	// '#/components/schemas/LoadBalancerPatchSerializer/properties/preferred_connectivity'
	// "$.components.schemas.LoadBalancerPatchSerializer.properties.preferred_connectivity"
	//
	// Any of "L2", "L3".
	PreferredConnectivity LoadBalancerMemberConnectivity `json:"preferred_connectivity,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LoadBalancerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/LoadBalancerPatchSerializer/properties/logging'
// "$.components.schemas.LoadBalancerPatchSerializer.properties.logging"
type LoadBalancerUpdateParamsLogging struct {
	// '#/components/schemas/LoadbalancerLoggingSerializer/properties/destination_region_id/anyOf/0'
	// "$.components.schemas.LoadbalancerLoggingSerializer.properties.destination_region_id.anyOf[0]"
	DestinationRegionID param.Opt[int64] `json:"destination_region_id,omitzero"`
	// '#/components/schemas/LoadbalancerLoggingSerializer/properties/topic_name/anyOf/0'
	// "$.components.schemas.LoadbalancerLoggingSerializer.properties.topic_name.anyOf[0]"
	TopicName param.Opt[string] `json:"topic_name,omitzero"`
	// '#/components/schemas/LoadbalancerLoggingSerializer/properties/enabled'
	// "$.components.schemas.LoadbalancerLoggingSerializer.properties.enabled"
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// '#/components/schemas/LoadbalancerLoggingSerializer/properties/retention_policy/anyOf/0'
	// "$.components.schemas.LoadbalancerLoggingSerializer.properties.retention_policy.anyOf[0]"
	RetentionPolicy LoadBalancerUpdateParamsLoggingRetentionPolicy `json:"retention_policy,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerUpdateParamsLogging) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r LoadBalancerUpdateParamsLogging) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerUpdateParamsLogging
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/LoadbalancerLoggingSerializer/properties/retention_policy/anyOf/0'
// "$.components.schemas.LoadbalancerLoggingSerializer.properties.retention_policy.anyOf[0]"
//
// The property Period is required.
type LoadBalancerUpdateParamsLoggingRetentionPolicy struct {
	// '#/components/schemas/LaasIndexRetentionPolicyPydanticSerializer/properties/period/anyOf/0'
	// "$.components.schemas.LaasIndexRetentionPolicyPydanticSerializer.properties.period.anyOf[0]"
	Period param.Opt[int64] `json:"period,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerUpdateParamsLoggingRetentionPolicy) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerUpdateParamsLoggingRetentionPolicy) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerUpdateParamsLoggingRetentionPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}

type LoadBalancerListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}'].get.parameters[2]"
	AssignedFloating param.Opt[bool] `query:"assigned_floating,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}'].get.parameters[3]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}'].get.parameters[4]"
	LoggingEnabled param.Opt[bool] `query:"logging_enabled,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}'].get.parameters[5]"
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/6'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}'].get.parameters[6]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/7'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}'].get.parameters[7]"
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/8'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}'].get.parameters[8]"
	ShowStats param.Opt[bool] `query:"show_stats,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/10'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}'].get.parameters[10]"
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/11'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}'].get.parameters[11]"
	WithDDOS param.Opt[bool] `query:"with_ddos,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/9'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}'].get.parameters[9]"
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LoadBalancerListParams]'s query parameters as `url.Values`.
func (r LoadBalancerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LoadBalancerDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bloadbalancer_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/{loadbalancer_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bloadbalancer_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/{loadbalancer_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type LoadBalancerFailoverParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bloadbalancer_id%7D%2Ffailover/post/parameters/0/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/{loadbalancer_id}/failover'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bloadbalancer_id%7D%2Ffailover/post/parameters/1/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/{loadbalancer_id}/failover'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/FailoverLoadBalancer/properties/force'
	// "$.components.schemas.FailoverLoadBalancer.properties.force"
	Force param.Opt[bool] `json:"force,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerFailoverParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LoadBalancerFailoverParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerFailoverParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type LoadBalancerGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bloadbalancer_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/{loadbalancer_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bloadbalancer_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/{loadbalancer_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bloadbalancer_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/{loadbalancer_id}'].get.parameters[3]"
	ShowStats param.Opt[bool] `query:"show_stats,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bloadbalancer_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/{loadbalancer_id}'].get.parameters[4]"
	WithDDOS param.Opt[bool] `query:"with_ddos,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LoadBalancerGetParams]'s query parameters as `url.Values`.
func (r LoadBalancerGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LoadBalancerResizeParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bloadbalancer_id%7D%2Fresize/post/parameters/0/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/{loadbalancer_id}/resize'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Floadbalancers%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bloadbalancer_id%7D%2Fresize/post/parameters/1/schema'
	// "$.paths['/cloud/v1/loadbalancers/{project_id}/{region_id}/{loadbalancer_id}/resize'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/ResizeLoadBalancer/properties/flavor'
	// "$.components.schemas.ResizeLoadBalancer.properties.flavor"
	Flavor string `json:"flavor,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerResizeParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LoadBalancerResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
