// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// GPUVirtualClusterInterfaceService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUVirtualClusterInterfaceService] method instead.
type GPUVirtualClusterInterfaceService struct {
	Options []option.RequestOption
}

// NewGPUVirtualClusterInterfaceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGPUVirtualClusterInterfaceService(opts ...option.RequestOption) (r GPUVirtualClusterInterfaceService) {
	r = GPUVirtualClusterInterfaceService{}
	r.Options = opts
	return
}

// List all network interfaces for servers in a virtual GPU cluster.
func (r *GPUVirtualClusterInterfaceService) List(ctx context.Context, clusterID string, query GPUVirtualClusterInterfaceListParams, opts ...option.RequestOption) (res *GPUVirtualInterfaceList, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/clusters/%s/interfaces", query.ProjectID.Value, query.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type GPUVirtualInterface struct {
	// Bodies of floatingips that are NAT-ing ips of this port
	FloatingIPs []GPUVirtualInterfaceFloatingIP `json:"floating_ips,required"`
	// IP addresses assigned to this port
	IPAssignments []GPUVirtualInterfaceIPAssignment `json:"ip_assignments,required"`
	// MAC address of the virtual port
	MacAddress string `json:"mac_address,required"`
	// Body of the network this port is attached to
	Network GPUVirtualInterfaceNetwork `json:"network,required"`
	// ID of the network the port is attached to
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// ID of virtual ethernet port object
	PortID string `json:"port_id,required" format:"uuid4"`
	// Port security status
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FloatingIPs         respjson.Field
		IPAssignments       respjson.Field
		MacAddress          respjson.Field
		Network             respjson.Field
		NetworkID           respjson.Field
		PortID              respjson.Field
		PortSecurityEnabled respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualInterface) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualInterfaceFloatingIP struct {
	// Floating IP ID
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the floating IP was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// IP address of the port the floating IP is attached to
	FixedIPAddress string `json:"fixed_ip_address,required" format:"ipvanyaddress"`
	// IP Address of the floating IP
	FloatingIPAddress string `json:"floating_ip_address,required" format:"ipvanyaddress"`
	// Port ID the floating IP is attached to. The `fixed_ip_address` is the IP address
	// of the port.
	PortID string `json:"port_id,required" format:"uuid4"`
	// Router ID
	RouterID string `json:"router_id,required" format:"uuid4"`
	// Floating IP status
	//
	// Any of "ACTIVE", "DOWN", "ERROR".
	Status FloatingIPStatus `json:"status,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// Datetime when the floating IP was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		FixedIPAddress    respjson.Field
		FloatingIPAddress respjson.Field
		PortID            respjson.Field
		RouterID          respjson.Field
		Status            respjson.Field
		Tags              respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualInterfaceFloatingIP) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualInterfaceFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualInterfaceIPAssignment struct {
	// The IP address assigned to the port from the specified subnet
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// ID of the subnet that allocated the IP
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress   respjson.Field
		SubnetID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualInterfaceIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualInterfaceIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Body of the network this port is attached to
type GPUVirtualInterfaceNetwork struct {
	// Network ID
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the network was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// True if the network `router:external` attribute
	External bool `json:"external,required"`
	// MTU (maximum transmission unit)
	Mtu int64 `json:"mtu,required"`
	// Network name
	Name string `json:"name,required"`
	// Indicates `port_security_enabled` status of all newly created in the network
	// ports.
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// Id of network segment
	SegmentationID int64 `json:"segmentation_id,required"`
	// True when the network is shared with your project by external owner
	Shared bool `json:"shared,required"`
	// List of subnetworks
	Subnets []GPUVirtualInterfaceNetworkSubnet `json:"subnets,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// Network type (vlan, vxlan)
	Type string `json:"type,required"`
	// Datetime when the network was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		External            respjson.Field
		Mtu                 respjson.Field
		Name                respjson.Field
		PortSecurityEnabled respjson.Field
		SegmentationID      respjson.Field
		Shared              respjson.Field
		Subnets             respjson.Field
		Tags                respjson.Field
		Type                respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualInterfaceNetwork) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualInterfaceNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualInterfaceNetworkSubnet struct {
	// Subnet id.
	ID string `json:"id,required" format:"uuid4"`
	// Number of available ips in subnet
	AvailableIPs int64 `json:"available_ips,required"`
	// CIDR
	Cidr string `json:"cidr,required" format:"ipvanynetwork"`
	// Datetime when the subnet was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// List IP addresses of a DNS resolver reachable from the network
	DNSNameservers []string `json:"dns_nameservers,required" format:"ipvanyaddress"`
	// Indicates whether DHCP is enabled for this subnet. If true, IP addresses will be
	// assigned automatically
	EnableDhcp bool `json:"enable_dhcp,required"`
	// Default GW IPv4 address, advertised in DHCP routes of this subnet. If null, no
	// gateway is advertised by this subnet.
	GatewayIP string `json:"gateway_ip,required" format:"ipvanyaddress"`
	// Deprecated. Always returns `false`.
	//
	// Deprecated: deprecated
	HasRouter bool `json:"has_router,required"`
	// List of custom static routes to advertise via DHCP.
	HostRoutes []Route `json:"host_routes,required"`
	// IP version used by the subnet (IPv4 or IPv6)
	//
	// Any of 4, 6.
	IPVersion int64 `json:"ip_version,required"`
	// Subnet name
	Name string `json:"name,required"`
	// Network ID
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// Total number of ips in subnet
	TotalIPs int64 `json:"total_ips,required"`
	// Datetime when the subnet was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AvailableIPs   respjson.Field
		Cidr           respjson.Field
		CreatedAt      respjson.Field
		DNSNameservers respjson.Field
		EnableDhcp     respjson.Field
		GatewayIP      respjson.Field
		HasRouter      respjson.Field
		HostRoutes     respjson.Field
		IPVersion      respjson.Field
		Name           respjson.Field
		NetworkID      respjson.Field
		Tags           respjson.Field
		TotalIPs       respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualInterfaceNetworkSubnet) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualInterfaceNetworkSubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualInterfaceList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []GPUVirtualInterface `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualInterfaceList) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualInterfaceList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterInterfaceListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
