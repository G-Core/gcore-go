// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1PortService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1PortService] method instead.
type CloudV1PortService struct {
	Options []option.RequestOption
}

// NewCloudV1PortService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1PortService(opts ...option.RequestOption) (r *CloudV1PortService) {
	r = &CloudV1PortService{}
	r.Options = opts
	return
}

// Deprecated. Assign allowed address pairs for instance port
func (r *CloudV1PortService) AllowAddressPairs(ctx context.Context, projectID int64, regionID int64, portID string, body CloudV1PortAllowAddressPairsParams, opts ...option.RequestOption) (res *CloudV1PortAllowAddressPairsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ports/%v/%v/%s/allow_address_pairs", projectID, regionID, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Disable port security for instance interface
func (r *CloudV1PortService) DisablePortSecurity(ctx context.Context, projectID int64, regionID int64, portID string, opts ...option.RequestOption) (res *InstancePort, err error) {
	opts = append(r.Options[:], opts...)
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ports/%v/%v/%s/disable_port_security", projectID, regionID, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Enable port security for instance interface
func (r *CloudV1PortService) EnablePortSecurity(ctx context.Context, projectID int64, regionID int64, portID string, opts ...option.RequestOption) (res *InstancePort, err error) {
	opts = append(r.Options[:], opts...)
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ports/%v/%v/%s/enable_port_security", projectID, regionID, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AllowedAddressPairs struct {
	// Subnet mask or IP address of the port specified in allowed_address_pairs
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// MAC address of the port specified in allowed_address_pairs
	MacAddress string                  `json:"mac_address,nullable"`
	JSON       allowedAddressPairsJSON `json:"-"`
}

// allowedAddressPairsJSON contains the JSON metadata for the struct
// [AllowedAddressPairs]
type allowedAddressPairsJSON struct {
	IPAddress   apijson.Field
	MacAddress  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AllowedAddressPairs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r allowedAddressPairsJSON) RawJSON() string {
	return r.raw
}

type AllowedAddressPairsParam struct {
	// Subnet mask or IP address of the port specified in allowed_address_pairs
	IPAddress param.Field[string] `json:"ip_address,required" format:"ipvanyaddress"`
	// MAC address of the port specified in allowed_address_pairs
	MacAddress param.Field[string] `json:"mac_address"`
}

func (r AllowedAddressPairsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Floating IP object
type DeprecatedFloatingIP struct {
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
	// Metadata
	Metadata interface{} `json:"metadata"`
	// If provided, floating IP will be immediately attached to the specified port
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
	Status DeprecatedFloatingIPStatus `json:"status"`
	// Subnet ID
	SubnetID string `json:"subnet_id"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id"`
	// Datetime when the floating IP was last updated
	UpdatedAt time.Time                `json:"updated_at" format:"date-time"`
	JSON      deprecatedFloatingIPJSON `json:"-"`
}

// deprecatedFloatingIPJSON contains the JSON metadata for the struct
// [DeprecatedFloatingIP]
type deprecatedFloatingIPJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	CreatorTaskID     apijson.Field
	DNSDomain         apijson.Field
	DNSName           apijson.Field
	FixedIPAddress    apijson.Field
	FloatingIPAddress apijson.Field
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

func (r *DeprecatedFloatingIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedFloatingIPJSON) RawJSON() string {
	return r.raw
}

// Floating IP status
type DeprecatedFloatingIPStatus string

const (
	DeprecatedFloatingIPStatusActive DeprecatedFloatingIPStatus = "ACTIVE"
	DeprecatedFloatingIPStatusDown   DeprecatedFloatingIPStatus = "DOWN"
	DeprecatedFloatingIPStatusError  DeprecatedFloatingIPStatus = "ERROR"
)

func (r DeprecatedFloatingIPStatus) IsKnown() bool {
	switch r {
	case DeprecatedFloatingIPStatusActive, DeprecatedFloatingIPStatusDown, DeprecatedFloatingIPStatusError:
		return true
	}
	return false
}

// InstancePortSchema schema
type InstancePort struct {
	// bodies of floatingips that are NAT-ing ips of this port.
	FloatingipDetails []DeprecatedFloatingIP `json:"floatingip_details,required"`
	// IP addresses assigned to this port
	IPAssignments []PortIP `json:"ip_assignments,required"`
	// Network schema with subnetwork details
	NetworkDetails InstancePortNetworkDetails `json:"network_details,required"`
	// ID of the network the port is attached to
	NetworkID string `json:"network_id,required"`
	// ID of virtual ethernet port object
	PortID string `json:"port_id,required"`
	// port security status
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// Group of subnet masks and/or IP addresses that share the current IP as VIP
	AllowedAddressPairs []InstancePortAllowedAddressPair `json:"allowed_address_pairs"`
	// Interface name
	InterfaceName string `json:"interface_name"`
	// MAC address of the virtual port
	MacAddress string           `json:"mac_address"`
	JSON       instancePortJSON `json:"-"`
}

// instancePortJSON contains the JSON metadata for the struct [InstancePort]
type instancePortJSON struct {
	FloatingipDetails   apijson.Field
	IPAssignments       apijson.Field
	NetworkDetails      apijson.Field
	NetworkID           apijson.Field
	PortID              apijson.Field
	PortSecurityEnabled apijson.Field
	AllowedAddressPairs apijson.Field
	InterfaceName       apijson.Field
	MacAddress          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InstancePort) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instancePortJSON) RawJSON() string {
	return r.raw
}

// Network schema with subnetwork details
type InstancePortNetworkDetails struct {
	// Network name
	Name string `json:"name,required"`
	// Network ID
	ID string `json:"id"`
	// Datetime when the network was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Task that created this entity
	CreatorTaskID interface{} `json:"creator_task_id"`
	// True if the network has is_default attribute
	Default bool `json:"default"`
	// True if the network has `router:external` attribute
	External bool `json:"external"`
	// Network metadata
	Metadata interface{} `json:"metadata"`
	// MTU (maximum transmission unit). Default value is 1450
	Mtu int64 `json:"mtu"`
	// Project ID
	ProjectID int64 `json:"project_id"`
	// Region name
	Region string `json:"region"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// Id of network segment
	SegmentationID int64 `json:"segmentation_id"`
	// True when the network is shared with your project by external owner
	Shared  bool                               `json:"shared"`
	Subnets []InstancePortNetworkDetailsSubnet `json:"subnets"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id"`
	// Network type (vlan, vxlan)
	Type string `json:"type"`
	// Datetime when the network was last updated
	UpdatedAt time.Time                      `json:"updated_at" format:"date-time"`
	JSON      instancePortNetworkDetailsJSON `json:"-"`
}

// instancePortNetworkDetailsJSON contains the JSON metadata for the struct
// [InstancePortNetworkDetails]
type instancePortNetworkDetailsJSON struct {
	Name           apijson.Field
	ID             apijson.Field
	CreatedAt      apijson.Field
	CreatorTaskID  apijson.Field
	Default        apijson.Field
	External       apijson.Field
	Metadata       apijson.Field
	Mtu            apijson.Field
	ProjectID      apijson.Field
	Region         apijson.Field
	RegionID       apijson.Field
	SegmentationID apijson.Field
	Shared         apijson.Field
	Subnets        apijson.Field
	TaskID         apijson.Field
	Type           apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstancePortNetworkDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instancePortNetworkDetailsJSON) RawJSON() string {
	return r.raw
}

// Subnet schema
type InstancePortNetworkDetailsSubnet struct {
	// CIDR
	Cidr string `json:"cidr,required"`
	// Subnet name
	Name string `json:"name,required"`
	// Network ID
	NetworkID string `json:"network_id,required"`
	// Subnet id.
	ID string `json:"id"`
	// Number of available ips in subnet
	AvailableIPs int64 `json:"available_ips"`
	// Datetime when the subnet was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Task that created this entity
	CreatorTaskID interface{} `json:"creator_task_id"`
	// List IP addresses of a DNS resolver reachable from the network
	DNSNameservers []string `json:"dns_nameservers"`
	// True if DHCP should be enabled
	EnableDhcp bool `json:"enable_dhcp"`
	// Default GW IPv4 address, advertised in DHCP routes of this subnet. If null, no
	// gateway is advertised by this subnet.
	GatewayIP string `json:"gateway_ip,nullable"`
	// Subnet has router attached to it
	HasRouter bool `json:"has_router"`
	// List of custom static routes to advertise via DHCP.
	HostRoutes []InstancePortNetworkDetailsSubnetsHostRoute `json:"host_routes"`
	// IP version
	IPVersion InstancePortNetworkDetailsSubnetsIPVersion `json:"ip_version"`
	// Subnetwork metadata
	Metadata interface{} `json:"metadata"`
	// Project ID
	ProjectID int64 `json:"project_id"`
	// Region name
	Region string `json:"region"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id"`
	// Total number of ips in subnet
	TotalIPs int64 `json:"total_ips"`
	// Datetime when the subnet was updated
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      instancePortNetworkDetailsSubnetJSON `json:"-"`
}

// instancePortNetworkDetailsSubnetJSON contains the JSON metadata for the struct
// [InstancePortNetworkDetailsSubnet]
type instancePortNetworkDetailsSubnetJSON struct {
	Cidr           apijson.Field
	Name           apijson.Field
	NetworkID      apijson.Field
	ID             apijson.Field
	AvailableIPs   apijson.Field
	CreatedAt      apijson.Field
	CreatorTaskID  apijson.Field
	DNSNameservers apijson.Field
	EnableDhcp     apijson.Field
	GatewayIP      apijson.Field
	HasRouter      apijson.Field
	HostRoutes     apijson.Field
	IPVersion      apijson.Field
	Metadata       apijson.Field
	ProjectID      apijson.Field
	Region         apijson.Field
	RegionID       apijson.Field
	TaskID         apijson.Field
	TotalIPs       apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstancePortNetworkDetailsSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instancePortNetworkDetailsSubnetJSON) RawJSON() string {
	return r.raw
}

// Custom route definition.
type InstancePortNetworkDetailsSubnetsHostRoute struct {
	// CIDR of destination IPv4 subnet.
	Destination string `json:"destination,required"`
	// IPv4 address to forward traffic to if it's destination IP matches 'destination'
	// CIDR.
	Nexthop string                                         `json:"nexthop,required"`
	JSON    instancePortNetworkDetailsSubnetsHostRouteJSON `json:"-"`
}

// instancePortNetworkDetailsSubnetsHostRouteJSON contains the JSON metadata for
// the struct [InstancePortNetworkDetailsSubnetsHostRoute]
type instancePortNetworkDetailsSubnetsHostRouteJSON struct {
	Destination apijson.Field
	Nexthop     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstancePortNetworkDetailsSubnetsHostRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instancePortNetworkDetailsSubnetsHostRouteJSON) RawJSON() string {
	return r.raw
}

// IP version
type InstancePortNetworkDetailsSubnetsIPVersion int64

const (
	InstancePortNetworkDetailsSubnetsIPVersion4 InstancePortNetworkDetailsSubnetsIPVersion = 4
	InstancePortNetworkDetailsSubnetsIPVersion6 InstancePortNetworkDetailsSubnetsIPVersion = 6
)

func (r InstancePortNetworkDetailsSubnetsIPVersion) IsKnown() bool {
	switch r {
	case InstancePortNetworkDetailsSubnetsIPVersion4, InstancePortNetworkDetailsSubnetsIPVersion6:
		return true
	}
	return false
}

type InstancePortAllowedAddressPair struct {
	// Subnet mask or IP address of the port specified in allowed_address_pairs
	IPAddress string `json:"ip_address"`
	// MAC address of the port specified in allowed_address_pairs
	MacAddress string                             `json:"mac_address"`
	JSON       instancePortAllowedAddressPairJSON `json:"-"`
}

// instancePortAllowedAddressPairJSON contains the JSON metadata for the struct
// [InstancePortAllowedAddressPair]
type instancePortAllowedAddressPairJSON struct {
	IPAddress   apijson.Field
	MacAddress  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstancePortAllowedAddressPair) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instancePortAllowedAddressPairJSON) RawJSON() string {
	return r.raw
}

// PortIpSchema schema
type PortIP struct {
	// IP address
	IPAddress string `json:"ip_address,required"`
	// ID of the subnet that allocated the IP
	SubnetID string     `json:"subnet_id,required"`
	JSON     portIPJSON `json:"-"`
}

// portIPJSON contains the JSON metadata for the struct [PortIP]
type portIPJSON struct {
	IPAddress   apijson.Field
	SubnetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PortIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r portIPJSON) RawJSON() string {
	return r.raw
}

type CloudV1PortAllowAddressPairsResponse struct {
	// A set of zero or more allowed port address pair and/or subnet masks
	AllowedAddressPairs []AllowedAddressPairs `json:"allowed_address_pairs"`
	// The ID of the instance that uses this port
	InstanceID string `json:"instance_id,nullable" format:"uuid4"`
	// The ID of the attached network
	NetworkID string `json:"network_id,nullable" format:"uuid4"`
	// The ID of the port
	PortID string                                   `json:"port_id,nullable" format:"uuid4"`
	JSON   cloudV1PortAllowAddressPairsResponseJSON `json:"-"`
}

// cloudV1PortAllowAddressPairsResponseJSON contains the JSON metadata for the
// struct [CloudV1PortAllowAddressPairsResponse]
type cloudV1PortAllowAddressPairsResponseJSON struct {
	AllowedAddressPairs apijson.Field
	InstanceID          apijson.Field
	NetworkID           apijson.Field
	PortID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudV1PortAllowAddressPairsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1PortAllowAddressPairsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1PortAllowAddressPairsParams struct {
	// A set of zero or more allowed port address pair and/or subnet masks
	AllowedAddressPairs param.Field[[]AllowedAddressPairsParam] `json:"allowed_address_pairs"`
	// The ID of the instance that uses this port
	InstanceID param.Field[string] `json:"instance_id" format:"uuid4"`
	// The ID of the attached network
	NetworkID param.Field[string] `json:"network_id" format:"uuid4"`
	// The ID of the port
	PortID param.Field[string] `json:"port_id" format:"uuid4"`
}

func (r CloudV1PortAllowAddressPairsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
