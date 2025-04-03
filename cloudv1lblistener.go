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

// CloudV1LblistenerService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1LblistenerService] method instead.
type CloudV1LblistenerService struct {
	Options []option.RequestOption
}

// NewCloudV1LblistenerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1LblistenerService(opts ...option.RequestOption) (r *CloudV1LblistenerService) {
	r = &CloudV1LblistenerService{}
	r.Options = opts
	return
}

// Create load balancer listener
func (r *CloudV1LblistenerService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1LblistenerNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/lblisteners/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get listener
func (r *CloudV1LblistenerService) Get(ctx context.Context, projectID int64, regionID int64, listenerID string, query CloudV1LblistenerGetParams, opts ...option.RequestOption) (res *LbListener, err error) {
	opts = append(r.Options[:], opts...)
	if listenerID == "" {
		err = errors.New("missing required listener_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lblisteners/%v/%v/%s", projectID, regionID, listenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Edit listener name
func (r *CloudV1LblistenerService) Update(ctx context.Context, projectID int64, regionID int64, listenerID string, body CloudV1LblistenerUpdateParams, opts ...option.RequestOption) (res *CloudV1LblistenerUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if listenerID == "" {
		err = errors.New("missing required listener_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lblisteners/%v/%v/%s", projectID, regionID, listenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List load balancer listeners
func (r *CloudV1LblistenerService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1LblistenerListParams, opts ...option.RequestOption) (res *CloudV1LblistenerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/lblisteners/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete load balancer listener
func (r *CloudV1LblistenerService) Delete(ctx context.Context, projectID int64, regionID int64, listenerID string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if listenerID == "" {
		err = errors.New("missing required listener_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lblisteners/%v/%v/%s", projectID, regionID, listenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type LbListener struct {
	// Load balancer listener ID
	ID string `json:"id,required" format:"uuid4"`
	// Limit of simultaneous connections
	ConnectionLimit int64 `json:"connection_limit,required"`
	// Dictionary of additional header insertion into HTTP headers. Only used with HTTP
	// and TERMINATED_HTTPS protocols.
	InsertHeaders interface{} `json:"insert_headers,required"`
	// Load balancer listener name
	Name string `json:"name,required"`
	// Listener operating status
	OperatingStatus OperatingStatus `json:"operating_status,required"`
	// Load balancer protocol
	Protocol LbListenerProtocol `json:"protocol,required"`
	// Protocol port
	ProtocolPort int64 `json:"protocol_port,required"`
	// Listener lifecycle status
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Network CIDRs from which service will be accessible
	AllowedCidrs []string `json:"allowed_cidrs,nullable" format:"ipvanynetwork"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// Load balancer ID
	LoadbalancerID string `json:"loadbalancer_id,nullable" format:"uuid4"`
	// Number of pools (for UI)
	PoolCount int64 `json:"pool_count,nullable"`
	// ID of the secret where PKCS12 file is stored for TERMINATED_HTTPS or PROMETHEUS
	// load balancer
	SecretID string `json:"secret_id,nullable"`
	// List of secret's ID containing PKCS12 format certificate/key bundles for
	// TERMINATED_HTTPS or PROMETHEUS listeners
	SniSecretID []string `json:"sni_secret_id,nullable"`
	// Statistics of the load balancer. It is available only in get functions by a
	// flag.
	Stats LoadbalancerStats `json:"stats,nullable"`
	// Active task. If none, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// Frontend client inactivity timeout in milliseconds
	TimeoutClientData int64 `json:"timeout_client_data,nullable"`
	// Backend member connection timeout in milliseconds
	TimeoutMemberConnect int64 `json:"timeout_member_connect,nullable"`
	// Backend member inactivity timeout in milliseconds
	TimeoutMemberData int64 `json:"timeout_member_data,nullable"`
	// Load balancer listener users list
	UserList []UserListItem `json:"user_list"`
	JSON     lbListenerJSON `json:"-"`
}

// lbListenerJSON contains the JSON metadata for the struct [LbListener]
type lbListenerJSON struct {
	ID                   apijson.Field
	ConnectionLimit      apijson.Field
	InsertHeaders        apijson.Field
	Name                 apijson.Field
	OperatingStatus      apijson.Field
	Protocol             apijson.Field
	ProtocolPort         apijson.Field
	ProvisioningStatus   apijson.Field
	AllowedCidrs         apijson.Field
	CreatorTaskID        apijson.Field
	LoadbalancerID       apijson.Field
	PoolCount            apijson.Field
	SecretID             apijson.Field
	SniSecretID          apijson.Field
	Stats                apijson.Field
	TaskID               apijson.Field
	TimeoutClientData    apijson.Field
	TimeoutMemberConnect apijson.Field
	TimeoutMemberData    apijson.Field
	UserList             apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LbListener) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lbListenerJSON) RawJSON() string {
	return r.raw
}

type LbListenerProtocol string

const (
	LbListenerProtocolHTTP            LbListenerProtocol = "HTTP"
	LbListenerProtocolHTTPS           LbListenerProtocol = "HTTPS"
	LbListenerProtocolPrometheus      LbListenerProtocol = "PROMETHEUS"
	LbListenerProtocolTcp             LbListenerProtocol = "TCP"
	LbListenerProtocolTerminatedHTTPS LbListenerProtocol = "TERMINATED_HTTPS"
	LbListenerProtocolUdp             LbListenerProtocol = "UDP"
)

func (r LbListenerProtocol) IsKnown() bool {
	switch r {
	case LbListenerProtocolHTTP, LbListenerProtocolHTTPS, LbListenerProtocolPrometheus, LbListenerProtocolTcp, LbListenerProtocolTerminatedHTTPS, LbListenerProtocolUdp:
		return true
	}
	return false
}

type LoadbalancerStats struct {
	// Currently active connections
	ActiveConnections int64 `json:"active_connections,required"`
	// Total bytes received
	BytesIn int64 `json:"bytes_in,required"`
	// Total bytes sent
	BytesOut int64 `json:"bytes_out,required"`
	// Total requests that were unable to be fulfilled
	RequestErrors int64 `json:"request_errors,required"`
	// Total connections handled
	TotalConnections int64                 `json:"total_connections,required"`
	JSON             loadbalancerStatsJSON `json:"-"`
}

// loadbalancerStatsJSON contains the JSON metadata for the struct
// [LoadbalancerStats]
type loadbalancerStatsJSON struct {
	ActiveConnections apijson.Field
	BytesIn           apijson.Field
	BytesOut          apijson.Field
	RequestErrors     apijson.Field
	TotalConnections  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LoadbalancerStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadbalancerStatsJSON) RawJSON() string {
	return r.raw
}

type OperatingStatus string

const (
	OperatingStatusDegraded  OperatingStatus = "DEGRADED"
	OperatingStatusDraining  OperatingStatus = "DRAINING"
	OperatingStatusError     OperatingStatus = "ERROR"
	OperatingStatusNoMonitor OperatingStatus = "NO_MONITOR"
	OperatingStatusOffline   OperatingStatus = "OFFLINE"
	OperatingStatusOnline    OperatingStatus = "ONLINE"
)

func (r OperatingStatus) IsKnown() bool {
	switch r {
	case OperatingStatusDegraded, OperatingStatusDraining, OperatingStatusError, OperatingStatusNoMonitor, OperatingStatusOffline, OperatingStatusOnline:
		return true
	}
	return false
}

type ProvisioningStatus string

const (
	ProvisioningStatusActive        ProvisioningStatus = "ACTIVE"
	ProvisioningStatusDeleted       ProvisioningStatus = "DELETED"
	ProvisioningStatusError         ProvisioningStatus = "ERROR"
	ProvisioningStatusPendingCreate ProvisioningStatus = "PENDING_CREATE"
	ProvisioningStatusPendingDelete ProvisioningStatus = "PENDING_DELETE"
	ProvisioningStatusPendingUpdate ProvisioningStatus = "PENDING_UPDATE"
)

func (r ProvisioningStatus) IsKnown() bool {
	switch r {
	case ProvisioningStatusActive, ProvisioningStatusDeleted, ProvisioningStatusError, ProvisioningStatusPendingCreate, ProvisioningStatusPendingDelete, ProvisioningStatusPendingUpdate:
		return true
	}
	return false
}

type UserListItem struct {
	// Encrypted password to auth via Basic Authentication
	EncryptedPassword string `json:"encrypted_password,required"`
	// Username to auth via Basic Authentication
	Username string           `json:"username,required"`
	JSON     userListItemJSON `json:"-"`
}

// userListItemJSON contains the JSON metadata for the struct [UserListItem]
type userListItemJSON struct {
	EncryptedPassword apijson.Field
	Username          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListItemJSON) RawJSON() string {
	return r.raw
}

type UserListItemParam struct {
	// Encrypted password to auth via Basic Authentication
	EncryptedPassword param.Field[string] `json:"encrypted_password,required"`
	// Username to auth via Basic Authentication
	Username param.Field[string] `json:"username,required"`
}

func (r UserListItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LblistenerUpdateResponse struct {
	// Load balancer listener ID
	ID string `json:"id,required" format:"uuid4"`
	// Limit of simultaneous connections
	ConnectionLimit int64 `json:"connection_limit,required"`
	// Dictionary of additional header insertion into HTTP headers. Only used with HTTP
	// and TERMINATED_HTTPS protocols.
	InsertHeaders interface{} `json:"insert_headers,required"`
	// Load balancer listener name
	Name string `json:"name,required"`
	// Listener operating status
	OperatingStatus OperatingStatus `json:"operating_status,required"`
	// Load balancer protocol
	Protocol LbListenerProtocol `json:"protocol,required"`
	// Protocol port
	ProtocolPort int64 `json:"protocol_port,required"`
	// Listener lifecycle status
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Network CIDRs from which service will be accessible
	AllowedCidrs []string `json:"allowed_cidrs,nullable" format:"ipvanynetwork"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// Load balancer ID
	LoadbalancerID string `json:"loadbalancer_id,nullable" format:"uuid4"`
	// Number of pools (for UI)
	PoolCount int64 `json:"pool_count,nullable"`
	// ID of the secret where PKCS12 file is stored for TERMINATED_HTTPS or PROMETHEUS
	// load balancer
	SecretID string `json:"secret_id,nullable"`
	// List of secret's ID containing PKCS12 format certificate/key bundles for
	// TERMINATED_HTTPS or PROMETHEUS listeners
	SniSecretID []string `json:"sni_secret_id,nullable"`
	// Active task. If none, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// Frontend client inactivity timeout in milliseconds
	TimeoutClientData int64 `json:"timeout_client_data,nullable"`
	// Backend member connection timeout in milliseconds
	TimeoutMemberConnect int64 `json:"timeout_member_connect,nullable"`
	// Backend member inactivity timeout in milliseconds
	TimeoutMemberData int64 `json:"timeout_member_data,nullable"`
	// Load balancer listener users list
	UserList []UserListItem                      `json:"user_list"`
	JSON     cloudV1LblistenerUpdateResponseJSON `json:"-"`
}

// cloudV1LblistenerUpdateResponseJSON contains the JSON metadata for the struct
// [CloudV1LblistenerUpdateResponse]
type cloudV1LblistenerUpdateResponseJSON struct {
	ID                   apijson.Field
	ConnectionLimit      apijson.Field
	InsertHeaders        apijson.Field
	Name                 apijson.Field
	OperatingStatus      apijson.Field
	Protocol             apijson.Field
	ProtocolPort         apijson.Field
	ProvisioningStatus   apijson.Field
	AllowedCidrs         apijson.Field
	CreatorTaskID        apijson.Field
	LoadbalancerID       apijson.Field
	PoolCount            apijson.Field
	SecretID             apijson.Field
	SniSecretID          apijson.Field
	TaskID               apijson.Field
	TimeoutClientData    apijson.Field
	TimeoutMemberConnect apijson.Field
	TimeoutMemberData    apijson.Field
	UserList             apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CloudV1LblistenerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1LblistenerUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1LblistenerListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []LbListener                      `json:"results,required"`
	JSON    cloudV1LblistenerListResponseJSON `json:"-"`
}

// cloudV1LblistenerListResponseJSON contains the JSON metadata for the struct
// [CloudV1LblistenerListResponse]
type cloudV1LblistenerListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1LblistenerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1LblistenerListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1LblistenerNewParams struct {
	// Load balancer ID
	LoadbalancerID param.Field[string] `json:"loadbalancer_id,required" format:"uuid4"`
	// Load balancer listener name
	Name param.Field[string] `json:"name,required"`
	// Load balancer listener protocol
	Protocol param.Field[LbListenerProtocol] `json:"protocol,required"`
	// Protocol port
	ProtocolPort param.Field[int64] `json:"protocol_port,required"`
	// Network CIDRs from which service will be accessible
	AllowedCidrs param.Field[[]string] `json:"allowed_cidrs" format:"ipvanynetwork"`
	// Limit of the simultaneous connections
	ConnectionLimit param.Field[int64] `json:"connection_limit"`
	// Add headers X-Forwarded-For, X-Forwarded-Port, X-Forwarded-Proto to requests.
	// Only used with HTTP or TERMINATED_HTTPS protocols.
	InsertXForwarded param.Field[bool] `json:"insert_x_forwarded"`
	// ID of the secret where PKCS12 file is stored for TERMINATED_HTTPS or PROMETHEUS
	// listener
	SecretID param.Field[string] `json:"secret_id"`
	// List of secrets IDs containing PKCS12 format certificate/key bundles for
	// TERMINATED_HTTPS or PROMETHEUS listeners
	SniSecretID param.Field[[]string] `json:"sni_secret_id" format:"uuid4"`
	// Frontend client inactivity timeout in milliseconds
	TimeoutClientData param.Field[int64] `json:"timeout_client_data"`
	// Backend member connection timeout in milliseconds
	TimeoutMemberConnect param.Field[int64] `json:"timeout_member_connect"`
	// Backend member inactivity timeout in milliseconds
	TimeoutMemberData param.Field[int64] `json:"timeout_member_data"`
	// Load balancer listener list of username and encrypted password items
	UserList param.Field[[]UserListItemParam] `json:"user_list"`
}

func (r CloudV1LblistenerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LblistenerGetParams struct {
	// Show statistics
	ShowStats param.Field[bool] `query:"show_stats"`
}

// URLQuery serializes [CloudV1LblistenerGetParams]'s query parameters as
// `url.Values`.
func (r CloudV1LblistenerGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1LblistenerUpdateParams struct {
	NamePydantic NamePydanticParam `json:"name_pydantic,required"`
}

func (r CloudV1LblistenerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.NamePydantic)
}

type CloudV1LblistenerListParams struct {
	// Load balancer ID
	LoadbalancerID param.Field[string] `query:"loadbalancer_id"`
	// Show statistics
	ShowStats param.Field[bool] `query:"show_stats"`
}

// URLQuery serializes [CloudV1LblistenerListParams]'s query parameters as
// `url.Values`.
func (r CloudV1LblistenerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
