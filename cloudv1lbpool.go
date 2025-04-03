// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1LbpoolService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1LbpoolService] method instead.
type CloudV1LbpoolService struct {
	Options       []option.RequestOption
	Healthmonitor *CloudV1LbpoolHealthmonitorService
	Member        *CloudV1LbpoolMemberService
}

// NewCloudV1LbpoolService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1LbpoolService(opts ...option.RequestOption) (r *CloudV1LbpoolService) {
	r = &CloudV1LbpoolService{}
	r.Options = opts
	r.Healthmonitor = NewCloudV1LbpoolHealthmonitorService(opts...)
	r.Member = NewCloudV1LbpoolMemberService(opts...)
	return
}

// Create load balancer pool
func (r *CloudV1LbpoolService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1LbpoolNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get load balancer pool
func (r *CloudV1LbpoolService) Get(ctx context.Context, projectID int64, regionID int64, poolID string, opts ...option.RequestOption) (res *DetailedLbPool, err error) {
	opts = append(r.Options[:], opts...)
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s", projectID, regionID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Changes provided here will overwrite existing load balancer pool settings.
// Undefined fields will be kept as is. Complex objects need to be specified fully,
// they will be overwritten.
func (r *CloudV1LbpoolService) Update(ctx context.Context, projectID int64, regionID int64, poolID string, body CloudV1LbpoolUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s", projectID, regionID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List load balancer pools
func (r *CloudV1LbpoolService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1LbpoolListParams, opts ...option.RequestOption) (res *CloudV1LbpoolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete load balancer pool
func (r *CloudV1LbpoolService) Delete(ctx context.Context, projectID int64, regionID int64, poolID string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s", projectID, regionID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CreateLbPoolParam struct {
	// Load balancer algorithm
	LbAlgorithm param.Field[LbAlgorithm] `json:"lb_algorithm,required"`
	// Pool name
	Name param.Field[string] `json:"name,required"`
	// Protocol
	Protocol param.Field[LbPoolProtocol] `json:"protocol,required"`
	// Secret ID of CA certificate bundle
	CaSecretID param.Field[string] `json:"ca_secret_id" format:"uuid4"`
	// Secret ID of CA revocation list file
	CrlSecretID param.Field[string] `json:"crl_secret_id" format:"uuid4"`
	// Health monitor details
	Healthmonitor param.Field[CreateLbHealthMonitorParam] `json:"healthmonitor"`
	// Listener ID
	ListenerID param.Field[string] `json:"listener_id" format:"uuid4"`
	// Loadbalancer ID
	LoadbalancerID param.Field[string] `json:"loadbalancer_id" format:"uuid4"`
	// Pool members
	Members param.Field[[]CreateLbPoolMemberParam] `json:"members"`
	// Secret ID for TLS client authentication to the member servers
	SecretID param.Field[string] `json:"secret_id" format:"uuid4"`
	// Session persistence details
	SessionPersistence param.Field[MutateLbSessionPersistenceParam] `json:"session_persistence"`
	// Frontend client inactivity timeout in milliseconds
	TimeoutClientData param.Field[int64] `json:"timeout_client_data"`
	// Backend member connection timeout in milliseconds
	TimeoutMemberConnect param.Field[int64] `json:"timeout_member_connect"`
	// Backend member inactivity timeout in milliseconds
	TimeoutMemberData param.Field[int64] `json:"timeout_member_data"`
}

func (r CreateLbPoolParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DetailedLbPool struct {
	// Pool ID
	ID string `json:"id,required" format:"uuid4"`
	// Secret ID of CA certificate bundle
	CaSecretID string `json:"ca_secret_id,required,nullable" format:"uuid4"`
	// Secret ID of CA revocation list file
	CrlSecretID string `json:"crl_secret_id,required,nullable" format:"uuid4"`
	// Load balancer algorithm
	LbAlgorithm LbAlgorithm `json:"lb_algorithm,required"`
	// Listeners IDs
	Listeners []MandatoryIDPydantic `json:"listeners,required"`
	// Load balancers IDs
	Loadbalancers []MandatoryIDPydantic `json:"loadbalancers,required"`
	// Pool members
	Members []DetailedLbPoolMember `json:"members,required"`
	// Pool name
	Name string `json:"name,required"`
	// Pool operating status
	OperatingStatus OperatingStatus `json:"operating_status,required"`
	// Protocol
	Protocol LbPoolProtocol `json:"protocol,required"`
	// Pool lifecycle status
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Secret ID for TLS client authentication to the member servers
	SecretID string `json:"secret_id,required,nullable" format:"uuid4"`
	// Session persistence parameters
	SessionPersistence DetailedLbPoolSessionPersistence `json:"session_persistence,required,nullable"`
	// Frontend client inactivity timeout in milliseconds
	TimeoutClientData int64 `json:"timeout_client_data,required,nullable"`
	// Backend member connection timeout in milliseconds
	TimeoutMemberConnect int64 `json:"timeout_member_connect,required,nullable"`
	// Backend member inactivity timeout in milliseconds
	TimeoutMemberData int64 `json:"timeout_member_data,required,nullable"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// Health monitor parameters
	Healthmonitor DetailedLbPoolHealthmonitor `json:"healthmonitor,nullable"`
	// Active task. If null, action has been performed immediately in the request
	// itself.
	TaskID string             `json:"task_id,nullable" format:"uuid4"`
	JSON   detailedLbPoolJSON `json:"-"`
}

// detailedLbPoolJSON contains the JSON metadata for the struct [DetailedLbPool]
type detailedLbPoolJSON struct {
	ID                   apijson.Field
	CaSecretID           apijson.Field
	CrlSecretID          apijson.Field
	LbAlgorithm          apijson.Field
	Listeners            apijson.Field
	Loadbalancers        apijson.Field
	Members              apijson.Field
	Name                 apijson.Field
	OperatingStatus      apijson.Field
	Protocol             apijson.Field
	ProvisioningStatus   apijson.Field
	SecretID             apijson.Field
	SessionPersistence   apijson.Field
	TimeoutClientData    apijson.Field
	TimeoutMemberConnect apijson.Field
	TimeoutMemberData    apijson.Field
	CreatorTaskID        apijson.Field
	Healthmonitor        apijson.Field
	TaskID               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DetailedLbPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailedLbPoolJSON) RawJSON() string {
	return r.raw
}

type DetailedLbPoolMember struct {
	// Member ID must be provided if an existing member is being updated
	ID string `json:"id,required" format:"uuid4"`
	// Member IP address
	Address string `json:"address,required" format:"ipvanyaddress"`
	// true if enabled. Defaults to true
	AdminStateUp bool `json:"admin_state_up,required"`
	// Member operating status of the entity
	OperatingStatus OperatingStatus `json:"operating_status,required"`
	// Member IP port
	ProtocolPort int64 `json:"protocol_port,required"`
	// Pool member lifecycle status
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Member weight. Valid values: 0 to 256, defaults to 1
	Weight int64 `json:"weight,required"`
	// An alternate IP address used for health monitoring of a backend member. Default
	// is null which monitors the member address.
	MonitorAddress string `json:"monitor_address,nullable" format:"ipvanyaddress"`
	// An alternate protocol port used for health monitoring of a backend member.
	// Default is null which monitors the member protocol_port.
	MonitorPort int64 `json:"monitor_port,nullable"`
	// Either subnet_id or instance_id should be provided
	SubnetID string                   `json:"subnet_id,nullable" format:"uuid4"`
	JSON     detailedLbPoolMemberJSON `json:"-"`
}

// detailedLbPoolMemberJSON contains the JSON metadata for the struct
// [DetailedLbPoolMember]
type detailedLbPoolMemberJSON struct {
	ID                 apijson.Field
	Address            apijson.Field
	AdminStateUp       apijson.Field
	OperatingStatus    apijson.Field
	ProtocolPort       apijson.Field
	ProvisioningStatus apijson.Field
	Weight             apijson.Field
	MonitorAddress     apijson.Field
	MonitorPort        apijson.Field
	SubnetID           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DetailedLbPoolMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailedLbPoolMemberJSON) RawJSON() string {
	return r.raw
}

// Session persistence parameters
type DetailedLbPoolSessionPersistence struct {
	// Session persistence type
	Type SessionPersistenceType `json:"type,required"`
	// Should be set if app cookie or http cookie is used
	CookieName string `json:"cookie_name,nullable"`
	// Subnet mask if source_ip is used. For UDP ports only
	PersistenceGranularity string `json:"persistence_granularity,nullable"`
	// Session persistence timeout. For UDP ports only
	PersistenceTimeout int64                                `json:"persistence_timeout,nullable"`
	JSON               detailedLbPoolSessionPersistenceJSON `json:"-"`
}

// detailedLbPoolSessionPersistenceJSON contains the JSON metadata for the struct
// [DetailedLbPoolSessionPersistence]
type detailedLbPoolSessionPersistenceJSON struct {
	Type                   apijson.Field
	CookieName             apijson.Field
	PersistenceGranularity apijson.Field
	PersistenceTimeout     apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *DetailedLbPoolSessionPersistence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailedLbPoolSessionPersistenceJSON) RawJSON() string {
	return r.raw
}

// Health monitor parameters
type DetailedLbPoolHealthmonitor struct {
	// Health monitor ID
	ID string `json:"id,required" format:"uuid4"`
	// true if enabled. Defaults to true
	AdminStateUp bool `json:"admin_state_up,required"`
	// The time, in seconds, between sending probes to members
	Delay int64 `json:"delay,required"`
	// Number of successes before the member is switched to ONLINE state
	MaxRetries int64 `json:"max_retries,required"`
	// Number of failures before the member is switched to ERROR state
	MaxRetriesDown int64 `json:"max_retries_down,required"`
	// Health Monitor operating status
	OperatingStatus OperatingStatus `json:"operating_status,required"`
	// Health monitor lifecycle status
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// The maximum time to connect. Must be less than the delay value
	Timeout int64 `json:"timeout,required"`
	// Health monitor type. Once health monitor is created, cannot be changed.
	Type          HealthMonitorType `json:"type,required"`
	ExpectedCodes string            `json:"expected_codes,nullable"`
	// HTTP method
	HTTPMethod HTTPMethodHealthmonitor `json:"http_method,nullable"`
	// URL Path. Defaults to '/'
	URLPath string                          `json:"url_path,nullable"`
	JSON    detailedLbPoolHealthmonitorJSON `json:"-"`
}

// detailedLbPoolHealthmonitorJSON contains the JSON metadata for the struct
// [DetailedLbPoolHealthmonitor]
type detailedLbPoolHealthmonitorJSON struct {
	ID                 apijson.Field
	AdminStateUp       apijson.Field
	Delay              apijson.Field
	MaxRetries         apijson.Field
	MaxRetriesDown     apijson.Field
	OperatingStatus    apijson.Field
	ProvisioningStatus apijson.Field
	Timeout            apijson.Field
	Type               apijson.Field
	ExpectedCodes      apijson.Field
	HTTPMethod         apijson.Field
	URLPath            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DetailedLbPoolHealthmonitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailedLbPoolHealthmonitorJSON) RawJSON() string {
	return r.raw
}

type LbAlgorithm string

const (
	LbAlgorithmLeastConnections LbAlgorithm = "LEAST_CONNECTIONS"
	LbAlgorithmRoundRobin       LbAlgorithm = "ROUND_ROBIN"
	LbAlgorithmSourceIP         LbAlgorithm = "SOURCE_IP"
)

func (r LbAlgorithm) IsKnown() bool {
	switch r {
	case LbAlgorithmLeastConnections, LbAlgorithmRoundRobin, LbAlgorithmSourceIP:
		return true
	}
	return false
}

type LbPoolProtocol string

const (
	LbPoolProtocolHTTP    LbPoolProtocol = "HTTP"
	LbPoolProtocolHTTPS   LbPoolProtocol = "HTTPS"
	LbPoolProtocolProxy   LbPoolProtocol = "PROXY"
	LbPoolProtocolProxyv2 LbPoolProtocol = "PROXYV2"
	LbPoolProtocolTcp     LbPoolProtocol = "TCP"
	LbPoolProtocolUdp     LbPoolProtocol = "UDP"
)

func (r LbPoolProtocol) IsKnown() bool {
	switch r {
	case LbPoolProtocolHTTP, LbPoolProtocolHTTPS, LbPoolProtocolProxy, LbPoolProtocolProxyv2, LbPoolProtocolTcp, LbPoolProtocolUdp:
		return true
	}
	return false
}

type MandatoryIDPydantic struct {
	// Resource ID
	ID   string                  `json:"id,required"`
	JSON mandatoryIDPydanticJSON `json:"-"`
}

// mandatoryIDPydanticJSON contains the JSON metadata for the struct
// [MandatoryIDPydantic]
type mandatoryIDPydanticJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MandatoryIDPydantic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mandatoryIDPydanticJSON) RawJSON() string {
	return r.raw
}

type MutateLbSessionPersistenceParam struct {
	// Session persistence type
	Type param.Field[SessionPersistenceType] `json:"type,required"`
	// Should be set if app cookie or http cookie is used
	CookieName param.Field[string] `json:"cookie_name"`
	// Subnet mask if source_ip is used. For UDP ports only
	PersistenceGranularity param.Field[string] `json:"persistence_granularity"`
	// Session persistence timeout. For UDP ports only
	PersistenceTimeout param.Field[int64] `json:"persistence_timeout"`
}

func (r MutateLbSessionPersistenceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SessionPersistenceType string

const (
	SessionPersistenceTypeAppCookie  SessionPersistenceType = "APP_COOKIE"
	SessionPersistenceTypeHTTPCookie SessionPersistenceType = "HTTP_COOKIE"
	SessionPersistenceTypeSourceIP   SessionPersistenceType = "SOURCE_IP"
)

func (r SessionPersistenceType) IsKnown() bool {
	switch r {
	case SessionPersistenceTypeAppCookie, SessionPersistenceTypeHTTPCookie, SessionPersistenceTypeSourceIP:
		return true
	}
	return false
}

type CloudV1LbpoolListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []DetailedLbPool              `json:"results,required"`
	JSON    cloudV1LbpoolListResponseJSON `json:"-"`
}

// cloudV1LbpoolListResponseJSON contains the JSON metadata for the struct
// [CloudV1LbpoolListResponse]
type cloudV1LbpoolListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1LbpoolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1LbpoolListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1LbpoolNewParams struct {
	CreateLbPool CreateLbPoolParam `json:"create_lb_pool,required"`
}

func (r CloudV1LbpoolNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateLbPool)
}

type CloudV1LbpoolUpdateParams struct {
	// Secret ID of CA certificate bundle
	CaSecretID param.Field[string] `json:"ca_secret_id" format:"uuid4"`
	// Secret ID of CA revocation list file
	CrlSecretID param.Field[string] `json:"crl_secret_id" format:"uuid4"`
	// New pool health monitor settings
	Healthmonitor param.Field[CloudV1LbpoolUpdateParamsHealthmonitor] `json:"healthmonitor"`
	// New load balancer pool algorithm of how to distribute requests
	LbAlgorithm param.Field[LbAlgorithm] `json:"lb_algorithm"`
	// New sequence of load balancer pool members. If members are the same (by
	// address + port), they will be kept as is without recreation and downtime.
	Members param.Field[[]CreateLbPoolMemberParam] `json:"members"`
	// New pool name
	Name param.Field[string] `json:"name"`
	// New communication protocol
	Protocol param.Field[LbPoolProtocol] `json:"protocol"`
	// Secret ID for TLS client authentication to the member servers
	SecretID param.Field[string] `json:"secret_id" format:"uuid4"`
	// New session persistence settings
	SessionPersistence param.Field[MutateLbSessionPersistenceParam] `json:"session_persistence"`
	// Frontend client inactivity timeout in milliseconds
	TimeoutClientData param.Field[int64] `json:"timeout_client_data"`
	// Backend member connection timeout in milliseconds
	TimeoutMemberConnect param.Field[int64] `json:"timeout_member_connect"`
	// Backend member inactivity timeout in milliseconds
	TimeoutMemberData param.Field[int64] `json:"timeout_member_data"`
}

func (r CloudV1LbpoolUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// New pool health monitor settings
type CloudV1LbpoolUpdateParamsHealthmonitor struct {
	// The time, in seconds, between sending probes to members
	Delay param.Field[int64] `json:"delay,required"`
	// Number of successes before the member is switched to ONLINE state
	MaxRetries param.Field[int64] `json:"max_retries,required"`
	// The maximum time to connect. Must be less than the delay value
	Timeout       param.Field[int64]  `json:"timeout,required"`
	ExpectedCodes param.Field[string] `json:"expected_codes"`
	// HTTP method
	HTTPMethod param.Field[HTTPMethodHealthmonitor] `json:"http_method"`
	// Number of failures before the member is switched to ERROR state
	MaxRetriesDown param.Field[int64] `json:"max_retries_down"`
	// Health monitor type. Once health monitor is created, cannot be changed.
	Type param.Field[HealthMonitorType] `json:"type"`
	// URL Path. Defaults to '/'
	URLPath param.Field[string] `json:"url_path"`
}

func (r CloudV1LbpoolUpdateParamsHealthmonitor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LbpoolListParams struct {
	// If true, show member and healthmonitor details of each pool (increases request
	// time)
	Details param.Field[bool] `query:"details"`
	// Load balancer listener ID
	ListenerID param.Field[string] `query:"listener_id"`
	// Load balancer ID
	LoadbalancerID param.Field[string] `query:"loadbalancer_id"`
}

// URLQuery serializes [CloudV1LbpoolListParams]'s query parameters as
// `url.Values`.
func (r CloudV1LbpoolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
