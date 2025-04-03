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

// CloudV1FloatingipService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1FloatingipService] method instead.
type CloudV1FloatingipService struct {
	Options      []option.RequestOption
	Metadata     *CloudV1FloatingipMetadataService
	MetadataItem *CloudV1FloatingipMetadataItemService
}

// NewCloudV1FloatingipService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1FloatingipService(opts ...option.RequestOption) (r *CloudV1FloatingipService) {
	r = &CloudV1FloatingipService{}
	r.Options = opts
	r.Metadata = NewCloudV1FloatingipMetadataService(opts...)
	r.MetadataItem = NewCloudV1FloatingipMetadataItemService(opts...)
	return
}

// Create floating IP
func (r *CloudV1FloatingipService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1FloatingipNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get floating IP
func (r *CloudV1FloatingipService) Get(ctx context.Context, projectID int64, regionID int64, pk string, opts ...option.RequestOption) (res *FloatingIP, err error) {
	opts = append(r.Options[:], opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List floating IPs
func (r *CloudV1FloatingipService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1FloatingipListParams, opts ...option.RequestOption) (res *CloudV1FloatingipListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete floating IP
func (r *CloudV1FloatingipService) Delete(ctx context.Context, projectID int64, regionID int64, pk string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Assign floating IP to instance or loadbalancer
func (r *CloudV1FloatingipService) Assign(ctx context.Context, projectID int64, regionID int64, pk string, body CloudV1FloatingipAssignParams, opts ...option.RequestOption) (res *FloatingIP, err error) {
	opts = append(r.Options[:], opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s/assign", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unassign floating IP from the instance
func (r *CloudV1FloatingipService) Unassign(ctx context.Context, projectID int64, regionID int64, pk string, opts ...option.RequestOption) (res *FloatingIP, err error) {
	opts = append(r.Options[:], opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s/unassign", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type DeprecatedNetworkPortFixedIP struct {
	// IP Address
	IPAddress string `json:"ip_address"`
	// Subnet UUID
	SubnetID string                           `json:"subnet_id" format:"uuid"`
	JSON     deprecatedNetworkPortFixedIPJSON `json:"-"`
}

// deprecatedNetworkPortFixedIPJSON contains the JSON metadata for the struct
// [DeprecatedNetworkPortFixedIP]
type deprecatedNetworkPortFixedIPJSON struct {
	IPAddress   apijson.Field
	SubnetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeprecatedNetworkPortFixedIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedNetworkPortFixedIPJSON) RawJSON() string {
	return r.raw
}

type FloatingIP struct {
	// Datetime when the floating IP was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Datetime when the floating IP was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Floating IP ID
	ID string `json:"id,nullable" format:"uuid4"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// Associated DNS domain
	DNSDomain string `json:"dns_domain,nullable"`
	// DNS name
	DNSName string `json:"dns_name,nullable"`
	// Fixed IP address
	FixedIPAddress string `json:"fixed_ip_address,nullable" format:"ipvanyaddress"`
	// IP Address
	FloatingIPAddress string `json:"floating_ip_address,nullable" format:"ipvanyaddress"`
	// Metadata
	Metadata []DetailedMetadata `json:"metadata"`
	// Port ID
	PortID string `json:"port_id,nullable" format:"uuid4"`
	// Router ID
	RouterID string `json:"router_id,nullable" format:"uuid4"`
	// Floating IP status
	Status FloatingIPStatus `json:"status,nullable"`
	// Subnet ID
	SubnetID string `json:"subnet_id,nullable" format:"uuid4"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string         `json:"task_id,nullable" format:"uuid4"`
	JSON   floatingIPJSON `json:"-"`
}

// floatingIPJSON contains the JSON metadata for the struct [FloatingIP]
type floatingIPJSON struct {
	CreatedAt         apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	UpdatedAt         apijson.Field
	ID                apijson.Field
	CreatorTaskID     apijson.Field
	DNSDomain         apijson.Field
	DNSName           apijson.Field
	FixedIPAddress    apijson.Field
	FloatingIPAddress apijson.Field
	Metadata          apijson.Field
	PortID            apijson.Field
	RouterID          apijson.Field
	Status            apijson.Field
	SubnetID          apijson.Field
	TaskID            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *FloatingIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r floatingIPJSON) RawJSON() string {
	return r.raw
}

// Floating IP status
type FloatingIPStatus string

const (
	FloatingIPStatusActive FloatingIPStatus = "ACTIVE"
	FloatingIPStatusDown   FloatingIPStatus = "DOWN"
	FloatingIPStatusError  FloatingIPStatus = "ERROR"
)

func (r FloatingIPStatus) IsKnown() bool {
	switch r {
	case FloatingIPStatusActive, FloatingIPStatusDown, FloatingIPStatusError:
		return true
	}
	return false
}

type CloudV1FloatingipListResponse struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []CloudV1FloatingipListResponseResult `json:"results"`
	JSON    cloudV1FloatingipListResponseJSON     `json:"-"`
}

// cloudV1FloatingipListResponseJSON contains the JSON metadata for the struct
// [CloudV1FloatingipListResponse]
type cloudV1FloatingipListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1FloatingipListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FloatingipListResponseJSON) RawJSON() string {
	return r.raw
}

// FloatingIPDetailed schema
type CloudV1FloatingipListResponseResult struct {
	// Floating IP ID
	ID string `json:"id"`
	// Datetime when the floating IP was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id"`
	// Associated DNS domain
	DNSDomain string `json:"dns_domain"`
	// DNS name
	DNSName string `json:"dns_name"`
	// Fixed IP address
	FixedIPAddress string `json:"fixed_ip_address"`
	// IP Address
	FloatingIPAddress string `json:"floating_ip_address"`
	// Instance schema
	Instance DeprecatedInstance `json:"instance"`
	// Loadbalancer schema
	Loadbalancer CloudV1FloatingipListResponseResultsLoadbalancer `json:"loadbalancer"`
	// Metadata
	Metadata interface{} `json:"metadata"`
	// Port ID
	PortID string `json:"port_id,nullable"`
	// Project ID
	ProjectID int64 `json:"project_id"`
	// Region name
	Region string `json:"region"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// Router ID
	RouterID string `json:"router_id"`
	// Floating IP status
	Status CloudV1FloatingipListResponseResultsStatus `json:"status"`
	// Subnet ID
	SubnetID string `json:"subnet_id"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id"`
	// Datetime when the floating IP was last updated
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	JSON      cloudV1FloatingipListResponseResultJSON `json:"-"`
}

// cloudV1FloatingipListResponseResultJSON contains the JSON metadata for the
// struct [CloudV1FloatingipListResponseResult]
type cloudV1FloatingipListResponseResultJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	CreatorTaskID     apijson.Field
	DNSDomain         apijson.Field
	DNSName           apijson.Field
	FixedIPAddress    apijson.Field
	FloatingIPAddress apijson.Field
	Instance          apijson.Field
	Loadbalancer      apijson.Field
	Metadata          apijson.Field
	PortID            apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	RouterID          apijson.Field
	Status            apijson.Field
	SubnetID          apijson.Field
	TaskID            apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1FloatingipListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FloatingipListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Loadbalancer schema
type CloudV1FloatingipListResponseResultsLoadbalancer struct {
	// Load balancer name
	Name string `json:"name,required"`
	// Load balancer ID
	ID string `json:"id"`
	// List of additional IP addresses
	AdditionalVips []DeprecatedNetworkPortFixedIP `json:"additional_vips"`
	// Datetime. Example value: "2019-06-13T13:58:12+0000"
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id"`
	// Get DDoS Protection service client profile
	DdosProfile DeprecatedGetClientProfile `json:"ddos_profile"`
	// Flavor object
	Flavor CloudV1FloatingipListResponseResultsLoadbalancerFlavor `json:"flavor"`
	// List of assigned floating IPs
	FloatingIPs []DeprecatedFloatingIP `json:"floating_ips"`
	// Load balancer listeners
	Listeners []CloudV1FloatingipListResponseResultsLoadbalancerListener `json:"listeners"`
	// Logging configuration
	Logging CloudV1FloatingipListResponseResultsLoadbalancerLogging `json:"logging"`
	// Create one or more metadata items for a loadbalancers
	Metadata interface{} `json:"metadata"`
	// Load balancer operating status
	OperatingStatus CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatus `json:"operating_status"`
	// Project ID
	ProjectID int64 `json:"project_id"`
	// Load balancer lifecycle status
	ProvisioningStatus CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatus `json:"provisioning_status"`
	// Region name
	Region string `json:"region"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// Statistic schema of a load balancer or its listener
	Stats CloudV1FloatingipListResponseResultsLoadbalancerStats `json:"stats"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id"`
	// Datetime. Example value: "2019-06-13T13:58:12+0000"
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Load balancer IP address
	VipAddress string `json:"vip_address"`
	// Load balancer IP family
	VipIPFamily CloudV1FloatingipListResponseResultsLoadbalancerVipIPFamily `json:"vip_ip_family"`
	VipPortID   string                                                      `json:"vip_port_id"`
	// List of VRRP IP addresses
	VrrpIPs []DeprecatedNetworkPortFixedIP                       `json:"vrrp_ips"`
	JSON    cloudV1FloatingipListResponseResultsLoadbalancerJSON `json:"-"`
}

// cloudV1FloatingipListResponseResultsLoadbalancerJSON contains the JSON metadata
// for the struct [CloudV1FloatingipListResponseResultsLoadbalancer]
type cloudV1FloatingipListResponseResultsLoadbalancerJSON struct {
	Name               apijson.Field
	ID                 apijson.Field
	AdditionalVips     apijson.Field
	CreatedAt          apijson.Field
	CreatorTaskID      apijson.Field
	DdosProfile        apijson.Field
	Flavor             apijson.Field
	FloatingIPs        apijson.Field
	Listeners          apijson.Field
	Logging            apijson.Field
	Metadata           apijson.Field
	OperatingStatus    apijson.Field
	ProjectID          apijson.Field
	ProvisioningStatus apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Stats              apijson.Field
	TaskID             apijson.Field
	UpdatedAt          apijson.Field
	VipAddress         apijson.Field
	VipIPFamily        apijson.Field
	VipPortID          apijson.Field
	VrrpIPs            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1FloatingipListResponseResultsLoadbalancer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FloatingipListResponseResultsLoadbalancerJSON) RawJSON() string {
	return r.raw
}

// Flavor object
type CloudV1FloatingipListResponseResultsLoadbalancerFlavor struct {
	// Flavor architecture type
	Architecture string `json:"architecture"`
	// Disabled flavor flag
	Disabled bool `json:"disabled"`
	// Short ID
	FlavorID string `json:"flavor_id"`
	// Human readable name
	FlavorName string `json:"flavor_name"`
	// Various hardware hints for flavor.
	HardwareDescription DeprecatedFlavorHardwareDescription `json:"hardware_description"`
	// Flavor operating system
	OsType string `json:"os_type"`
	// RAM size in MiB
	Ram int64 `json:"ram"`
	// Flavor resource class for mapping to hardware capacity
	ResourceClass string `json:"resource_class"`
	// Virtual CPU count
	Vcpus int64                                                      `json:"vcpus"`
	JSON  cloudV1FloatingipListResponseResultsLoadbalancerFlavorJSON `json:"-"`
}

// cloudV1FloatingipListResponseResultsLoadbalancerFlavorJSON contains the JSON
// metadata for the struct [CloudV1FloatingipListResponseResultsLoadbalancerFlavor]
type cloudV1FloatingipListResponseResultsLoadbalancerFlavorJSON struct {
	Architecture        apijson.Field
	Disabled            apijson.Field
	FlavorID            apijson.Field
	FlavorName          apijson.Field
	HardwareDescription apijson.Field
	OsType              apijson.Field
	Ram                 apijson.Field
	ResourceClass       apijson.Field
	Vcpus               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudV1FloatingipListResponseResultsLoadbalancerFlavor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FloatingipListResponseResultsLoadbalancerFlavorJSON) RawJSON() string {
	return r.raw
}

type CloudV1FloatingipListResponseResultsLoadbalancerListener struct {
	Name string `json:"name,required"`
	ID   string `json:"id"`
	// Network CIDRs from which service will be accessible
	AllowedCidrs []string `json:"allowed_cidrs,nullable"`
	// Limit of the simultaneous connections
	ConnectionLimit int64  `json:"connection_limit"`
	Description     string `json:"description"`
	// Listener operating status
	OperatingStatus    CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatus    `json:"operating_status"`
	Protocol           string                                                                      `json:"protocol"`
	ProtocolPort       int64                                                                       `json:"protocol_port"`
	ProvisioningStatus CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatus `json:"provisioning_status"`
	// Load balancer listener users list
	UserList []CloudV1FloatingipListResponseResultsLoadbalancerListenersUserList `json:"user_list"`
	JSON     cloudV1FloatingipListResponseResultsLoadbalancerListenerJSON        `json:"-"`
}

// cloudV1FloatingipListResponseResultsLoadbalancerListenerJSON contains the JSON
// metadata for the struct
// [CloudV1FloatingipListResponseResultsLoadbalancerListener]
type cloudV1FloatingipListResponseResultsLoadbalancerListenerJSON struct {
	Name               apijson.Field
	ID                 apijson.Field
	AllowedCidrs       apijson.Field
	ConnectionLimit    apijson.Field
	Description        apijson.Field
	OperatingStatus    apijson.Field
	Protocol           apijson.Field
	ProtocolPort       apijson.Field
	ProvisioningStatus apijson.Field
	UserList           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudV1FloatingipListResponseResultsLoadbalancerListener) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FloatingipListResponseResultsLoadbalancerListenerJSON) RawJSON() string {
	return r.raw
}

// Listener operating status
type CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatus string

const (
	CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatusDegraded  CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatus = "DEGRADED"
	CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatusDraining  CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatus = "DRAINING"
	CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatusError     CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatus = "ERROR"
	CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatusNoMonitor CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatus = "NO_MONITOR"
	CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatusOffline   CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatus = "OFFLINE"
	CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatusOnline    CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatus = "ONLINE"
)

func (r CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatus) IsKnown() bool {
	switch r {
	case CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatusDegraded, CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatusDraining, CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatusError, CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatusNoMonitor, CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatusOffline, CloudV1FloatingipListResponseResultsLoadbalancerListenersOperatingStatusOnline:
		return true
	}
	return false
}

type CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatus string

const (
	CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatusActive        CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatus = "ACTIVE"
	CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatusDeleted       CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatus = "DELETED"
	CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatusError         CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatus = "ERROR"
	CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatusPendingCreate CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatus = "PENDING_CREATE"
	CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatusPendingDelete CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatus = "PENDING_DELETE"
	CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatusPendingUpdate CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatus = "PENDING_UPDATE"
)

func (r CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatus) IsKnown() bool {
	switch r {
	case CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatusActive, CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatusDeleted, CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatusError, CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatusPendingCreate, CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatusPendingDelete, CloudV1FloatingipListResponseResultsLoadbalancerListenersProvisioningStatusPendingUpdate:
		return true
	}
	return false
}

type CloudV1FloatingipListResponseResultsLoadbalancerListenersUserList struct {
	// Encrypted password to auth via Basic Authentication
	EncryptedPassword string `json:"encrypted_password,required"`
	// Username to auth via Basic Authentication
	Username string                                                                `json:"username,required"`
	JSON     cloudV1FloatingipListResponseResultsLoadbalancerListenersUserListJSON `json:"-"`
}

// cloudV1FloatingipListResponseResultsLoadbalancerListenersUserListJSON contains
// the JSON metadata for the struct
// [CloudV1FloatingipListResponseResultsLoadbalancerListenersUserList]
type cloudV1FloatingipListResponseResultsLoadbalancerListenersUserListJSON struct {
	EncryptedPassword apijson.Field
	Username          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1FloatingipListResponseResultsLoadbalancerListenersUserList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FloatingipListResponseResultsLoadbalancerListenersUserListJSON) RawJSON() string {
	return r.raw
}

// Logging configuration
type CloudV1FloatingipListResponseResultsLoadbalancerLogging struct {
	// Destination region id to which the logs will be written
	DestinationRegionID int64 `json:"destination_region_id,nullable"`
	// 'True' to activate, 'False' to deactivate logging
	Enabled         bool                     `json:"enabled"`
	RetentionPolicy LaasIndexRetentionPolicy `json:"retention_policy"`
	// Existing logging topic name or None. Use only with 'enabled': True
	TopicName string                                                      `json:"topic_name,nullable"`
	JSON      cloudV1FloatingipListResponseResultsLoadbalancerLoggingJSON `json:"-"`
}

// cloudV1FloatingipListResponseResultsLoadbalancerLoggingJSON contains the JSON
// metadata for the struct
// [CloudV1FloatingipListResponseResultsLoadbalancerLogging]
type cloudV1FloatingipListResponseResultsLoadbalancerLoggingJSON struct {
	DestinationRegionID apijson.Field
	Enabled             apijson.Field
	RetentionPolicy     apijson.Field
	TopicName           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudV1FloatingipListResponseResultsLoadbalancerLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FloatingipListResponseResultsLoadbalancerLoggingJSON) RawJSON() string {
	return r.raw
}

// Load balancer operating status
type CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatus string

const (
	CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatusDegraded  CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatus = "DEGRADED"
	CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatusDraining  CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatus = "DRAINING"
	CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatusError     CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatus = "ERROR"
	CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatusNoMonitor CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatus = "NO_MONITOR"
	CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatusOffline   CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatus = "OFFLINE"
	CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatusOnline    CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatus = "ONLINE"
)

func (r CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatus) IsKnown() bool {
	switch r {
	case CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatusDegraded, CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatusDraining, CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatusError, CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatusNoMonitor, CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatusOffline, CloudV1FloatingipListResponseResultsLoadbalancerOperatingStatusOnline:
		return true
	}
	return false
}

// Load balancer lifecycle status
type CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatus string

const (
	CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatusActive        CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatus = "ACTIVE"
	CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatusDeleted       CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatus = "DELETED"
	CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatusError         CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatus = "ERROR"
	CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatusPendingCreate CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatus = "PENDING_CREATE"
	CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatusPendingDelete CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatus = "PENDING_DELETE"
	CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatusPendingUpdate CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatus = "PENDING_UPDATE"
)

func (r CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatus) IsKnown() bool {
	switch r {
	case CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatusActive, CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatusDeleted, CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatusError, CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatusPendingCreate, CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatusPendingDelete, CloudV1FloatingipListResponseResultsLoadbalancerProvisioningStatusPendingUpdate:
		return true
	}
	return false
}

// Statistic schema of a load balancer or its listener
type CloudV1FloatingipListResponseResultsLoadbalancerStats struct {
	// The currently active connections
	ActiveConnections int64 `json:"active_connections,required"`
	// The total bytes received
	BytesIn int64 `json:"bytes_in,required"`
	// The total bytes sent
	BytesOut int64 `json:"bytes_out,required"`
	// The total requests that were unable to be fulfilled
	RequestErrors int64 `json:"request_errors,required"`
	// The total connections handled
	TotalConnections int64                                                     `json:"total_connections,required"`
	JSON             cloudV1FloatingipListResponseResultsLoadbalancerStatsJSON `json:"-"`
}

// cloudV1FloatingipListResponseResultsLoadbalancerStatsJSON contains the JSON
// metadata for the struct [CloudV1FloatingipListResponseResultsLoadbalancerStats]
type cloudV1FloatingipListResponseResultsLoadbalancerStatsJSON struct {
	ActiveConnections apijson.Field
	BytesIn           apijson.Field
	BytesOut          apijson.Field
	RequestErrors     apijson.Field
	TotalConnections  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1FloatingipListResponseResultsLoadbalancerStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FloatingipListResponseResultsLoadbalancerStatsJSON) RawJSON() string {
	return r.raw
}

// Load balancer IP family
type CloudV1FloatingipListResponseResultsLoadbalancerVipIPFamily string

const (
	CloudV1FloatingipListResponseResultsLoadbalancerVipIPFamilyDual CloudV1FloatingipListResponseResultsLoadbalancerVipIPFamily = "dual"
	CloudV1FloatingipListResponseResultsLoadbalancerVipIPFamilyIpv4 CloudV1FloatingipListResponseResultsLoadbalancerVipIPFamily = "ipv4"
	CloudV1FloatingipListResponseResultsLoadbalancerVipIPFamilyIpv6 CloudV1FloatingipListResponseResultsLoadbalancerVipIPFamily = "ipv6"
)

func (r CloudV1FloatingipListResponseResultsLoadbalancerVipIPFamily) IsKnown() bool {
	switch r {
	case CloudV1FloatingipListResponseResultsLoadbalancerVipIPFamilyDual, CloudV1FloatingipListResponseResultsLoadbalancerVipIPFamilyIpv4, CloudV1FloatingipListResponseResultsLoadbalancerVipIPFamilyIpv6:
		return true
	}
	return false
}

// Floating IP status
type CloudV1FloatingipListResponseResultsStatus string

const (
	CloudV1FloatingipListResponseResultsStatusActive CloudV1FloatingipListResponseResultsStatus = "ACTIVE"
	CloudV1FloatingipListResponseResultsStatusDown   CloudV1FloatingipListResponseResultsStatus = "DOWN"
	CloudV1FloatingipListResponseResultsStatusError  CloudV1FloatingipListResponseResultsStatus = "ERROR"
)

func (r CloudV1FloatingipListResponseResultsStatus) IsKnown() bool {
	switch r {
	case CloudV1FloatingipListResponseResultsStatusActive, CloudV1FloatingipListResponseResultsStatusDown, CloudV1FloatingipListResponseResultsStatusError:
		return true
	}
	return false
}

type CloudV1FloatingipNewParams struct {
	// In case the port has multiple IPs, specific address can be selected using this
	// field. If unspecified, first IP in port's list is used
	FixedIPAddress param.Field[string] `json:"fixed_ip_address"`
	// Create one or more metadata items for a floating IP
	Metadata param.Field[interface{}] `json:"metadata"`
	// If provided, floating IP will be immediately attached to the specified port
	PortID param.Field[string] `json:"port_id"`
}

func (r CloudV1FloatingipNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1FloatingipListParams struct {
	// Limit the number of returned limit request entities.
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV1FloatingipListParams]'s query parameters as
// `url.Values`.
func (r CloudV1FloatingipListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1FloatingipAssignParams struct {
	// Port ID
	PortID param.Field[string] `json:"port_id,required" format:"uuid4"`
	// Fixed IP address
	FixedIPAddress param.Field[string] `json:"fixed_ip_address" format:"ipvanyaddress"`
}

func (r CloudV1FloatingipAssignParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
