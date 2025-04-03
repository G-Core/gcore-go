// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1LoadbalancerService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1LoadbalancerService] method instead.
type CloudV1LoadbalancerService struct {
	Options      []option.RequestOption
	Status       *CloudV1LoadbalancerStatusService
	Metadata     *CloudV1LoadbalancerMetadataService
	MetadataItem *CloudV1LoadbalancerMetadataItemService
}

// NewCloudV1LoadbalancerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1LoadbalancerService(opts ...option.RequestOption) (r *CloudV1LoadbalancerService) {
	r = &CloudV1LoadbalancerService{}
	r.Options = opts
	r.Status = NewCloudV1LoadbalancerStatusService(opts...)
	r.Metadata = NewCloudV1LoadbalancerMetadataService(opts...)
	r.MetadataItem = NewCloudV1LoadbalancerMetadataItemService(opts...)
	return
}

// Create load balancer
func (r *CloudV1LoadbalancerService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1LoadbalancerNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get load balancer
func (r *CloudV1LoadbalancerService) Get(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, query CloudV1LoadbalancerGetParams, opts ...option.RequestOption) (res *Loadbalancer, err error) {
	opts = append(r.Options[:], opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Rename load balancer, activate/deactivate logs or update preferred connectivity
// for load balancer
func (r *CloudV1LoadbalancerService) Update(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, body CloudV1LoadbalancerUpdateParams, opts ...option.RequestOption) (res *Loadbalancer, err error) {
	opts = append(r.Options[:], opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List load balancers
func (r *CloudV1LoadbalancerService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1LoadbalancerListParams, opts ...option.RequestOption) (res *CloudV1LoadbalancerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete load balancer
func (r *CloudV1LoadbalancerService) Delete(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Check if regional quota is exceeded, if yes the number of additional quotas
// needed to create the specified load balancer will be calculated
func (r *CloudV1LoadbalancerService) CheckQuota(ctx context.Context, projectID int64, regionID int64, body CloudV1LoadbalancerCheckQuotaParams, opts ...option.RequestOption) (res *RegionalDiffQuotas, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/check_limits", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Failover loadbalancer
func (r *CloudV1LoadbalancerService) Failover(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, body CloudV1LoadbalancerFailoverParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/failover", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get loadbalancer metrics, including cpu, memory and network
func (r *CloudV1LoadbalancerService) GetMetrics(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, body CloudV1LoadbalancerGetMetricsParams, opts ...option.RequestOption) (res *CloudV1LoadbalancerGetMetricsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/metrics", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Resize loadbalancer
func (r *CloudV1LoadbalancerService) Resize(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, body CloudV1LoadbalancerResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/resize", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GetClientProfile struct {
	// DDoS protection profile ID
	ID int64 `json:"id,required"`
	// Template data
	ProfileTemplate ClientProfileTemplate   `json:"profile_template,required"`
	Fields          []ClientProfileField    `json:"fields"`
	Options         GetClientProfileOptions `json:"options,nullable"`
	// DDoS profile template description
	ProfileTemplateDescription string `json:"profile_template_description,nullable"`
	// List of protocols
	Protocols []interface{}          `json:"protocols,nullable"`
	Site      string                 `json:"site,nullable"`
	Status    GetClientProfileStatus `json:"status,nullable"`
	JSON      getClientProfileJSON   `json:"-"`
}

// getClientProfileJSON contains the JSON metadata for the struct
// [GetClientProfile]
type getClientProfileJSON struct {
	ID                         apijson.Field
	ProfileTemplate            apijson.Field
	Fields                     apijson.Field
	Options                    apijson.Field
	ProfileTemplateDescription apijson.Field
	Protocols                  apijson.Field
	Site                       apijson.Field
	Status                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GetClientProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getClientProfileJSON) RawJSON() string {
	return r.raw
}

type GetClientProfileOptions struct {
	// Activate profile.
	Active bool `json:"active,nullable"`
	// Activate BGP protocol.
	Bgp  bool                        `json:"bgp,nullable"`
	JSON getClientProfileOptionsJSON `json:"-"`
}

// getClientProfileOptionsJSON contains the JSON metadata for the struct
// [GetClientProfileOptions]
type getClientProfileOptionsJSON struct {
	Active      apijson.Field
	Bgp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetClientProfileOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getClientProfileOptionsJSON) RawJSON() string {
	return r.raw
}

type GetClientProfileStatus struct {
	// Description of the error, if it exists
	ErrorDescription string `json:"error_description,required"`
	// Profile status
	Status string                     `json:"status,required"`
	JSON   getClientProfileStatusJSON `json:"-"`
}

// getClientProfileStatusJSON contains the JSON metadata for the struct
// [GetClientProfileStatus]
type getClientProfileStatusJSON struct {
	ErrorDescription apijson.Field
	Status           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *GetClientProfileStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getClientProfileStatusJSON) RawJSON() string {
	return r.raw
}

type InterfaceIPFamily string

const (
	InterfaceIPFamilyDual InterfaceIPFamily = "dual"
	InterfaceIPFamilyIpv4 InterfaceIPFamily = "ipv4"
	InterfaceIPFamilyIpv6 InterfaceIPFamily = "ipv6"
)

func (r InterfaceIPFamily) IsKnown() bool {
	switch r {
	case InterfaceIPFamilyDual, InterfaceIPFamilyIpv4, InterfaceIPFamilyIpv6:
		return true
	}
	return false
}

type LaasIndexRetentionPolicy struct {
	// Duration of days for which logs must be kept.
	Period int64                        `json:"period,nullable"`
	JSON   laasIndexRetentionPolicyJSON `json:"-"`
}

// laasIndexRetentionPolicyJSON contains the JSON metadata for the struct
// [LaasIndexRetentionPolicy]
type laasIndexRetentionPolicyJSON struct {
	Period      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LaasIndexRetentionPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r laasIndexRetentionPolicyJSON) RawJSON() string {
	return r.raw
}

type LaasIndexRetentionPolicyParam struct {
	// Duration of days for which logs must be kept.
	Period param.Field[int64] `json:"period"`
}

func (r LaasIndexRetentionPolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Loadbalancer struct {
	// Load balancer ID
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the load balancer was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Load balancer name
	Name string `json:"name,required"`
	// Load balancer operating status
	OperatingStatus OperatingStatus `json:"operating_status,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Load balancer lifecycle status
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// List of additional IP addresses
	AdditionalVips []LoadbalancerAdditionalVip `json:"additional_vips"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// Loadbalancer advanced DDoS protection profile.
	DdosProfile GetClientProfile `json:"ddos_profile,nullable"`
	// Load balancer flavor (if not default)
	Flavor LoadbalancerFlavor `json:"flavor,nullable"`
	// List of assigned floating IPs
	FloatingIPs []FloatingIP `json:"floating_ips"`
	// Load balancer listeners
	Listeners []LoadbalancerListener `json:"listeners"`
	// Logging configuration
	Logging LoadbalancerLogging `json:"logging,nullable"`
	// Metadata items for a loadbalancer
	Metadata []DetailedMetadata `json:"metadata"`
	// Preferred option to establish connectivity between load balancer and its pools
	// members
	PreferredConnectivity MemberConnectivity `json:"preferred_connectivity"`
	// Statistics of load balancer.
	Stats LoadbalancerStats `json:"stats,nullable"`
	// Active task. If null, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// Datetime when the load balancer was last updated
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// Load balancer IP address
	VipAddress string `json:"vip_address,nullable" format:"ipvanyaddress"`
	// Load balancer IP family
	VipIPFamily InterfaceIPFamily `json:"vip_ip_family,nullable"`
	// The ID of the Virtual IP (VIP) port.
	VipPortID string `json:"vip_port_id,nullable" format:"uuid4"`
	// List of VRRP IP addresses
	VrrpIPs []LoadbalancerVrrpIP `json:"vrrp_ips"`
	JSON    loadbalancerJSON     `json:"-"`
}

// loadbalancerJSON contains the JSON metadata for the struct [Loadbalancer]
type loadbalancerJSON struct {
	ID                    apijson.Field
	CreatedAt             apijson.Field
	Name                  apijson.Field
	OperatingStatus       apijson.Field
	ProjectID             apijson.Field
	ProvisioningStatus    apijson.Field
	Region                apijson.Field
	RegionID              apijson.Field
	AdditionalVips        apijson.Field
	CreatorTaskID         apijson.Field
	DdosProfile           apijson.Field
	Flavor                apijson.Field
	FloatingIPs           apijson.Field
	Listeners             apijson.Field
	Logging               apijson.Field
	Metadata              apijson.Field
	PreferredConnectivity apijson.Field
	Stats                 apijson.Field
	TaskID                apijson.Field
	UpdatedAt             apijson.Field
	VipAddress            apijson.Field
	VipIPFamily           apijson.Field
	VipPortID             apijson.Field
	VrrpIPs               apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Loadbalancer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadbalancerJSON) RawJSON() string {
	return r.raw
}

type LoadbalancerAdditionalVip struct {
	// IP address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// Subnet UUID
	SubnetID string                        `json:"subnet_id,required" format:"uuid4"`
	JSON     loadbalancerAdditionalVipJSON `json:"-"`
}

// loadbalancerAdditionalVipJSON contains the JSON metadata for the struct
// [LoadbalancerAdditionalVip]
type loadbalancerAdditionalVipJSON struct {
	IPAddress   apijson.Field
	SubnetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadbalancerAdditionalVip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadbalancerAdditionalVipJSON) RawJSON() string {
	return r.raw
}

// Load balancer flavor (if not default)
type LoadbalancerFlavor struct {
	// Short ID
	FlavorID string `json:"flavor_id,required"`
	// Human readable name
	FlavorName string `json:"flavor_name,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Virtual CPU count
	Vcpus int64                  `json:"vcpus,required"`
	JSON  loadbalancerFlavorJSON `json:"-"`
}

// loadbalancerFlavorJSON contains the JSON metadata for the struct
// [LoadbalancerFlavor]
type loadbalancerFlavorJSON struct {
	FlavorID    apijson.Field
	FlavorName  apijson.Field
	Ram         apijson.Field
	Vcpus       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadbalancerFlavor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadbalancerFlavorJSON) RawJSON() string {
	return r.raw
}

type LoadbalancerListener struct {
	// Listener ID
	ID   string                   `json:"id,required" format:"uuid4"`
	JSON loadbalancerListenerJSON `json:"-"`
}

// loadbalancerListenerJSON contains the JSON metadata for the struct
// [LoadbalancerListener]
type loadbalancerListenerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadbalancerListener) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadbalancerListenerJSON) RawJSON() string {
	return r.raw
}

type LoadbalancerVrrpIP struct {
	// IP address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// LoadBalancer instance role to which VRRP IP belong
	Role LoadbalancerVrrpIPsRole `json:"role,required"`
	// Subnet UUID
	SubnetID string                 `json:"subnet_id,required" format:"uuid4"`
	JSON     loadbalancerVrrpIPJSON `json:"-"`
}

// loadbalancerVrrpIPJSON contains the JSON metadata for the struct
// [LoadbalancerVrrpIP]
type loadbalancerVrrpIPJSON struct {
	IPAddress   apijson.Field
	Role        apijson.Field
	SubnetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadbalancerVrrpIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadbalancerVrrpIPJSON) RawJSON() string {
	return r.raw
}

// LoadBalancer instance role to which VRRP IP belong
type LoadbalancerVrrpIPsRole string

const (
	LoadbalancerVrrpIPsRoleBackup     LoadbalancerVrrpIPsRole = "BACKUP"
	LoadbalancerVrrpIPsRoleMaster     LoadbalancerVrrpIPsRole = "MASTER"
	LoadbalancerVrrpIPsRoleStandalone LoadbalancerVrrpIPsRole = "STANDALONE"
)

func (r LoadbalancerVrrpIPsRole) IsKnown() bool {
	switch r {
	case LoadbalancerVrrpIPsRoleBackup, LoadbalancerVrrpIPsRoleMaster, LoadbalancerVrrpIPsRoleStandalone:
		return true
	}
	return false
}

type LoadbalancerLogging struct {
	// Destination region id to which the logs will be written
	DestinationRegionID int64 `json:"destination_region_id,nullable"`
	// Enable/disable forwarding logs to LaaS
	Enabled bool `json:"enabled"`
	// The logs retention policy
	RetentionPolicy LaasIndexRetentionPolicyPydantic `json:"retention_policy,nullable"`
	// The topic name to which the logs will be written
	TopicName string                  `json:"topic_name,nullable"`
	JSON      loadbalancerLoggingJSON `json:"-"`
}

// loadbalancerLoggingJSON contains the JSON metadata for the struct
// [LoadbalancerLogging]
type loadbalancerLoggingJSON struct {
	DestinationRegionID apijson.Field
	Enabled             apijson.Field
	RetentionPolicy     apijson.Field
	TopicName           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadbalancerLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadbalancerLoggingJSON) RawJSON() string {
	return r.raw
}

type LoadbalancerLoggingParam struct {
	// Destination region id to which the logs will be written
	DestinationRegionID param.Field[int64] `json:"destination_region_id"`
	// Enable/disable forwarding logs to LaaS
	Enabled param.Field[bool] `json:"enabled"`
	// The logs retention policy
	RetentionPolicy param.Field[LaasIndexRetentionPolicyPydanticParam] `json:"retention_policy"`
	// The topic name to which the logs will be written
	TopicName param.Field[string] `json:"topic_name"`
}

func (r LoadbalancerLoggingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberConnectivity string

const (
	MemberConnectivityL2 MemberConnectivity = "L2"
	MemberConnectivityL3 MemberConnectivity = "L3"
)

func (r MemberConnectivity) IsKnown() bool {
	switch r {
	case MemberConnectivityL2, MemberConnectivityL3:
		return true
	}
	return false
}

// NewInterfaceFloatingIpSerializer schema
type NewInterfaceFloatingIPParam struct {
	// Will the floating for this subnet's IP be created or reused?
	Source param.Field[NewInterfaceFloatingIPSource] `json:"source,required"`
	// If source is 'existing', existing floating IP ID must be specified
	ExistingFloatingID param.Field[string] `json:"existing_floating_id"`
}

func (r NewInterfaceFloatingIPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Will the floating for this subnet's IP be created or reused?
type NewInterfaceFloatingIPSource string

const (
	NewInterfaceFloatingIPSourceExisting NewInterfaceFloatingIPSource = "existing"
	NewInterfaceFloatingIPSourceNew      NewInterfaceFloatingIPSource = "new"
)

func (r NewInterfaceFloatingIPSource) IsKnown() bool {
	switch r {
	case NewInterfaceFloatingIPSourceExisting, NewInterfaceFloatingIPSourceNew:
		return true
	}
	return false
}

type NewInterfaceFloatingIPPydanticParam struct {
	// 'New' if the floating for this subnet's IP will be created, 'Existing' if it
	// will be reused
	Source param.Field[NewInterfaceFloatingIPPydanticSource] `json:"source,required"`
	// An existing floating IP id must be specified if the source is set to 'existing'
	ExistingFloatingID param.Field[string] `json:"existing_floating_id"`
}

func (r NewInterfaceFloatingIPPydanticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// 'New' if the floating for this subnet's IP will be created, 'Existing' if it
// will be reused
type NewInterfaceFloatingIPPydanticSource string

const (
	NewInterfaceFloatingIPPydanticSourceExisting NewInterfaceFloatingIPPydanticSource = "existing"
	NewInterfaceFloatingIPPydanticSourceNew      NewInterfaceFloatingIPPydanticSource = "new"
)

func (r NewInterfaceFloatingIPPydanticSource) IsKnown() bool {
	switch r {
	case NewInterfaceFloatingIPPydanticSourceExisting, NewInterfaceFloatingIPPydanticSourceNew:
		return true
	}
	return false
}

type CloudV1LoadbalancerListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Loadbalancer                      `json:"results,required"`
	JSON    cloudV1LoadbalancerListResponseJSON `json:"-"`
}

// cloudV1LoadbalancerListResponseJSON contains the JSON metadata for the struct
// [CloudV1LoadbalancerListResponse]
type cloudV1LoadbalancerListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1LoadbalancerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1LoadbalancerListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1LoadbalancerGetMetricsResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1LoadbalancerGetMetricsResponseResult `json:"results,required"`
	JSON    cloudV1LoadbalancerGetMetricsResponseJSON     `json:"-"`
}

// cloudV1LoadbalancerGetMetricsResponseJSON contains the JSON metadata for the
// struct [CloudV1LoadbalancerGetMetricsResponse]
type cloudV1LoadbalancerGetMetricsResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1LoadbalancerGetMetricsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1LoadbalancerGetMetricsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1LoadbalancerGetMetricsResponseResult struct {
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
	Time string                                          `json:"time,nullable"`
	JSON cloudV1LoadbalancerGetMetricsResponseResultJSON `json:"-"`
}

// cloudV1LoadbalancerGetMetricsResponseResultJSON contains the JSON metadata for
// the struct [CloudV1LoadbalancerGetMetricsResponseResult]
type cloudV1LoadbalancerGetMetricsResponseResultJSON struct {
	CPUUtil           apijson.Field
	MemoryUtil        apijson.Field
	NetworkBpsEgress  apijson.Field
	NetworkBpsIngress apijson.Field
	NetworkPpsEgress  apijson.Field
	NetworkPpsIngress apijson.Field
	Time              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1LoadbalancerGetMetricsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1LoadbalancerGetMetricsResponseResultJSON) RawJSON() string {
	return r.raw
}

type CloudV1LoadbalancerNewParams struct {
	// Load balancer flavor name
	Flavor param.Field[string] `json:"flavor"`
	// Floating IP configuration for assignment
	FloatingIP param.Field[NewInterfaceFloatingIPPydanticParam] `json:"floating_ip"`
	// Load balancer listeners. Maximum 50 per LB (excluding Prometheus endpoint
	// listener).
	Listeners param.Field[[]CloudV1LoadbalancerNewParamsListener] `json:"listeners"`
	// Logging configuration
	Logging param.Field[LoadbalancerLoggingParam] `json:"logging"`
	// A list of simple strings assigned to the load balancer
	Metadata param.Field[interface{}] `json:"metadata"`
	// Load balancer name
	Name param.Field[string] `json:"name"`
	// Load balancer name which will be changed by template.
	NameTemplate param.Field[string] `json:"name_template"`
	// Preferred option to establish connectivity between load balancer and its pools
	// members. L2 provides best performance, L3 provides less IPs usage. It is taking
	// effect only if instance_id + ip_address is provided, not subnet_id + ip_address,
	// because we're considering this as intentional subnet_id specification.
	PreferredConnectivity param.Field[MemberConnectivity] `json:"preferred_connectivity"`
	// A list of inner tags (k8s, k8s_ingress) assigned to the load balancer
	Tag param.Field[[]string] `json:"tag"`
	// IP family for load balancer subnet auto-selection if vip_network_id is specified
	VipIPFamily param.Field[InterfaceIPFamily] `json:"vip_ip_family"`
	// Network ID for load balancer. If not specified, default external network will be
	// used. Mutually exclusive with vip_port_id
	VipNetworkID param.Field[string] `json:"vip_network_id" format:"uuid4"`
	// Existing Reserved Fixed IP port ID for load balancer. Mutually exclusive with
	// vip_network_id
	VipPortID param.Field[string] `json:"vip_port_id" format:"uuid4"`
	// Subnet ID for load balancer. If not specified, any subnet from vip_network_id
	// will be selected. Ignored when vip_network_id is not specified.
	VipSubnetID param.Field[string] `json:"vip_subnet_id" format:"uuid4"`
}

func (r CloudV1LoadbalancerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LoadbalancerNewParamsListener struct {
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
	// Member pools
	Pools param.Field[[]CreateLbPoolParam] `json:"pools"`
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

func (r CloudV1LoadbalancerNewParamsListener) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LoadbalancerGetParams struct {
	// Show statistics
	ShowStats param.Field[bool] `query:"show_stats"`
	// Show DDoS profile
	WithDdos param.Field[bool] `query:"with_ddos"`
}

// URLQuery serializes [CloudV1LoadbalancerGetParams]'s query parameters as
// `url.Values`.
func (r CloudV1LoadbalancerGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1LoadbalancerUpdateParams struct {
	// Logging configuration
	Logging param.Field[LoadbalancerLoggingParam] `json:"logging"`
	// Name.
	Name param.Field[string] `json:"name"`
	// Preferred option to establish connectivity between load balancer and its pools
	// members
	PreferredConnectivity param.Field[MemberConnectivity] `json:"preferred_connectivity"`
}

func (r CloudV1LoadbalancerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LoadbalancerListParams struct {
	// With or without assigned floating IP
	AssignedFloating param.Field[bool] `query:"assigned_floating"`
	// Limit the number of returned limit request entities.
	Limit param.Field[int64] `query:"limit"`
	// With or without logging
	LoggingEnabled param.Field[bool] `query:"logging_enabled"`
	// Filter by metadata keys. Must be a valid JSON string. curl -G --data-urlencode
	// "metadata_k=["value", "sense"]" --url
	// "http://localhost:1111/v1/loadbalancers/1/1"
	MetadataK param.Field[string] `query:"metadata_k"`
	// Filter by metadata key-value pairs. Must be a valid JSON string. curl -G
	// --data-urlencode "metadata_kv={"key": "value"}" --url
	// "http://localhost:1111/v1/loadbalancers/1/1"
	MetadataKv param.Field[string] `query:"metadata_kv"`
	// Filter by name
	Name param.Field[string] `query:"name"`
	// Offset value is used to exclude the first set of records from the result.
	Offset param.Field[int64] `query:"offset"`
	// Ordering Load Balancer list result by name, created_at, updated_at,
	// operating_status, provisioning_status, vip_address, vip_ip_family and flavor
	// fields of the load balancer and directions (name.asc), default is
	// "created_at.asc"
	OrderBy param.Field[string] `query:"order_by"`
	// Show statistics
	ShowStats param.Field[bool] `query:"show_stats"`
	// Show Advanced DDoS protection profile, if exists
	WithDdos param.Field[bool] `query:"with_ddos"`
}

// URLQuery serializes [CloudV1LoadbalancerListParams]'s query parameters as
// `url.Values`.
func (r CloudV1LoadbalancerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1LoadbalancerCheckQuotaParams struct {
	// NewInterfaceFloatingIpSerializer schema
	FloatingIP param.Field[NewInterfaceFloatingIPParam] `json:"floating_ip"`
	// Logging configuration
	Logging param.Field[CloudV1LoadbalancerCheckQuotaParamsLogging] `json:"logging"`
}

func (r CloudV1LoadbalancerCheckQuotaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logging configuration
type CloudV1LoadbalancerCheckQuotaParamsLogging struct {
	// Destination region id to which the logs will be written
	DestinationRegionID param.Field[int64] `json:"destination_region_id"`
	// 'True' to activate, 'False' to deactivate logging
	Enabled         param.Field[bool]                          `json:"enabled"`
	RetentionPolicy param.Field[LaasIndexRetentionPolicyParam] `json:"retention_policy"`
	// Existing logging topic name or None. Use only with 'enabled': True
	TopicName param.Field[string] `json:"topic_name"`
}

func (r CloudV1LoadbalancerCheckQuotaParamsLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LoadbalancerFailoverParams struct {
	// Validate current load balancer status before failover or not.
	Force param.Field[bool] `json:"force"`
}

func (r CloudV1LoadbalancerFailoverParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LoadbalancerGetMetricsParams struct {
	// Time interval
	TimeInterval param.Field[int64] `json:"time_interval,required"`
	// Time interval unit
	TimeUnit param.Field[InstanceMetricsTimeUnit] `json:"time_unit,required"`
}

func (r CloudV1LoadbalancerGetMetricsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LoadbalancerResizeParams struct {
	// Name of the desired flavor to resize to.
	Flavor param.Field[string] `json:"flavor,required"`
}

func (r CloudV1LoadbalancerResizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
