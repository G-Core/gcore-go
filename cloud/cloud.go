// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// CloudService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudService] method instead.
type CloudService struct {
	Options          []option.RequestOption
	Projects         ProjectService
	Tasks            TaskService
	Regions          RegionService
	Quotas           QuotaService
	Secrets          SecretService
	SSHKeys          SSHKeyService
	IPRanges         IPRangeService
	ReservedFixedIPs ReservedFixedIPService
}

// NewCloudService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCloudService(opts ...option.RequestOption) (r CloudService) {
	r = CloudService{}
	r.Options = opts
	r.Projects = NewProjectService(opts...)
	r.Tasks = NewTaskService(opts...)
	r.Regions = NewRegionService(opts...)
	r.Quotas = NewQuotaService(opts...)
	r.Secrets = NewSecretService(opts...)
	r.SSHKeys = NewSSHKeyService(opts...)
	r.IPRanges = NewIPRangeService(opts...)
	r.ReservedFixedIPs = NewReservedFixedIPService(opts...)
	return
}

// '#/components/schemas/InterfaceIPFamily'
// "$.components.schemas.InterfaceIPFamily"
type InterfaceIPFamily string

const (
	InterfaceIPFamilyDual InterfaceIPFamily = "dual"
	InterfaceIPFamilyIpv4 InterfaceIPFamily = "ipv4"
	InterfaceIPFamilyIpv6 InterfaceIPFamily = "ipv6"
)

// '#/components/schemas/NetworkSerializer'
// "$.components.schemas.NetworkSerializer"
type Network struct {
	// '#/components/schemas/NetworkSerializer/properties/id'
	// "$.components.schemas.NetworkSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/NetworkSerializer/properties/created_at'
	// "$.components.schemas.NetworkSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/NetworkSerializer/properties/external'
	// "$.components.schemas.NetworkSerializer.properties.external"
	External bool `json:"external,required"`
	// '#/components/schemas/NetworkSerializer/properties/metadata'
	// "$.components.schemas.NetworkSerializer.properties.metadata"
	//
	// Deprecated: deprecated
	Metadata []Tag `json:"metadata,required"`
	// '#/components/schemas/NetworkSerializer/properties/name'
	// "$.components.schemas.NetworkSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/NetworkSerializer/properties/port_security_enabled'
	// "$.components.schemas.NetworkSerializer.properties.port_security_enabled"
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// '#/components/schemas/NetworkSerializer/properties/region'
	// "$.components.schemas.NetworkSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/NetworkSerializer/properties/region_id'
	// "$.components.schemas.NetworkSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/NetworkSerializer/properties/shared'
	// "$.components.schemas.NetworkSerializer.properties.shared"
	Shared bool `json:"shared,required"`
	// '#/components/schemas/NetworkSerializer/properties/subnets'
	// "$.components.schemas.NetworkSerializer.properties.subnets"
	Subnets []string `json:"subnets,required" format:"uuid4"`
	// '#/components/schemas/NetworkSerializer/properties/tags'
	// "$.components.schemas.NetworkSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/NetworkSerializer/properties/type'
	// "$.components.schemas.NetworkSerializer.properties.type"
	Type string `json:"type,required"`
	// '#/components/schemas/NetworkSerializer/properties/updated_at'
	// "$.components.schemas.NetworkSerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// '#/components/schemas/NetworkSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.NetworkSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// '#/components/schemas/NetworkSerializer/properties/default/anyOf/0'
	// "$.components.schemas.NetworkSerializer.properties['default'].anyOf[0]"
	Default bool `json:"default,nullable"`
	// '#/components/schemas/NetworkSerializer/properties/mtu'
	// "$.components.schemas.NetworkSerializer.properties.mtu"
	Mtu int64 `json:"mtu"`
	// '#/components/schemas/NetworkSerializer/properties/project_id/anyOf/0'
	// "$.components.schemas.NetworkSerializer.properties.project_id.anyOf[0]"
	ProjectID int64 `json:"project_id,nullable"`
	// '#/components/schemas/NetworkSerializer/properties/segmentation_id/anyOf/0'
	// "$.components.schemas.NetworkSerializer.properties.segmentation_id.anyOf[0]"
	SegmentationID int64 `json:"segmentation_id,nullable"`
	// '#/components/schemas/NetworkSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.NetworkSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                  resp.Field
		CreatedAt           resp.Field
		External            resp.Field
		Metadata            resp.Field
		Name                resp.Field
		PortSecurityEnabled resp.Field
		Region              resp.Field
		RegionID            resp.Field
		Shared              resp.Field
		Subnets             resp.Field
		Tags                resp.Field
		Type                resp.Field
		UpdatedAt           resp.Field
		CreatorTaskID       resp.Field
		Default             resp.Field
		Mtu                 resp.Field
		ProjectID           resp.Field
		SegmentationID      resp.Field
		TaskID              resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Network) RawJSON() string { return r.JSON.raw }
func (r *Network) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/SubnetSerializer' "$.components.schemas.SubnetSerializer"
type Subnet struct {
	// '#/components/schemas/SubnetSerializer/properties/cidr'
	// "$.components.schemas.SubnetSerializer.properties.cidr"
	Cidr string `json:"cidr,required" format:"ipvanynetwork"`
	// '#/components/schemas/SubnetSerializer/properties/created_at'
	// "$.components.schemas.SubnetSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/SubnetSerializer/properties/enable_dhcp'
	// "$.components.schemas.SubnetSerializer.properties.enable_dhcp"
	EnableDhcp bool `json:"enable_dhcp,required"`
	// '#/components/schemas/SubnetSerializer/properties/ip_version'
	// "$.components.schemas.SubnetSerializer.properties.ip_version"
	//
	// Any of 4, 6.
	IPVersion int64 `json:"ip_version,required"`
	// '#/components/schemas/SubnetSerializer/properties/name'
	// "$.components.schemas.SubnetSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/SubnetSerializer/properties/network_id'
	// "$.components.schemas.SubnetSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/SubnetSerializer/properties/project_id'
	// "$.components.schemas.SubnetSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/SubnetSerializer/properties/region'
	// "$.components.schemas.SubnetSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/SubnetSerializer/properties/region_id'
	// "$.components.schemas.SubnetSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/SubnetSerializer/properties/tags'
	// "$.components.schemas.SubnetSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/SubnetSerializer/properties/updated_at'
	// "$.components.schemas.SubnetSerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// '#/components/schemas/SubnetSerializer/properties/id/anyOf/0'
	// "$.components.schemas.SubnetSerializer.properties.id.anyOf[0]"
	ID string `json:"id,nullable" format:"uuid4"`
	// '#/components/schemas/SubnetSerializer/properties/available_ips/anyOf/0'
	// "$.components.schemas.SubnetSerializer.properties.available_ips.anyOf[0]"
	AvailableIPs int64 `json:"available_ips,nullable"`
	// '#/components/schemas/SubnetSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.SubnetSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// '#/components/schemas/SubnetSerializer/properties/dns_nameservers/anyOf/0'
	// "$.components.schemas.SubnetSerializer.properties.dns_nameservers.anyOf[0]"
	DNSNameservers []string `json:"dns_nameservers,nullable" format:"ipvanyaddress"`
	// '#/components/schemas/SubnetSerializer/properties/gateway_ip/anyOf/0'
	// "$.components.schemas.SubnetSerializer.properties.gateway_ip.anyOf[0]"
	GatewayIP string `json:"gateway_ip,nullable" format:"ipvanyaddress"`
	// '#/components/schemas/SubnetSerializer/properties/has_router'
	// "$.components.schemas.SubnetSerializer.properties.has_router"
	HasRouter bool `json:"has_router"`
	// '#/components/schemas/SubnetSerializer/properties/host_routes/anyOf/0'
	// "$.components.schemas.SubnetSerializer.properties.host_routes.anyOf[0]"
	HostRoutes []SubnetHostRoute `json:"host_routes,nullable"`
	// '#/components/schemas/SubnetSerializer/properties/metadata'
	// "$.components.schemas.SubnetSerializer.properties.metadata"
	//
	// Deprecated: deprecated
	Metadata []Tag `json:"metadata"`
	// '#/components/schemas/SubnetSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.SubnetSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// '#/components/schemas/SubnetSerializer/properties/total_ips/anyOf/0'
	// "$.components.schemas.SubnetSerializer.properties.total_ips.anyOf[0]"
	TotalIPs int64 `json:"total_ips,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Cidr           resp.Field
		CreatedAt      resp.Field
		EnableDhcp     resp.Field
		IPVersion      resp.Field
		Name           resp.Field
		NetworkID      resp.Field
		ProjectID      resp.Field
		Region         resp.Field
		RegionID       resp.Field
		Tags           resp.Field
		UpdatedAt      resp.Field
		ID             resp.Field
		AvailableIPs   resp.Field
		CreatorTaskID  resp.Field
		DNSNameservers resp.Field
		GatewayIP      resp.Field
		HasRouter      resp.Field
		HostRoutes     resp.Field
		Metadata       resp.Field
		TaskID         resp.Field
		TotalIPs       resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Subnet) RawJSON() string { return r.JSON.raw }
func (r *Subnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/SubnetSerializer/properties/host_routes/anyOf/0/items'
// "$.components.schemas.SubnetSerializer.properties.host_routes.anyOf[0].items"
type SubnetHostRoute struct {
	// '#/components/schemas/NeutronRouteSerializer/properties/destination'
	// "$.components.schemas.NeutronRouteSerializer.properties.destination"
	Destination string `json:"destination,required" format:"ipvanynetwork"`
	// '#/components/schemas/NeutronRouteSerializer/properties/nexthop'
	// "$.components.schemas.NeutronRouteSerializer.properties.nexthop"
	Nexthop string `json:"nexthop,required" format:"ipvanyaddress"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Destination resp.Field
		Nexthop     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubnetHostRoute) RawJSON() string { return r.JSON.raw }
func (r *SubnetHostRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/TagSerializer' "$.components.schemas.TagSerializer"
type Tag struct {
	// '#/components/schemas/TagSerializer/properties/key'
	// "$.components.schemas.TagSerializer.properties.key"
	Key string `json:"key,required"`
	// '#/components/schemas/TagSerializer/properties/read_only'
	// "$.components.schemas.TagSerializer.properties.read_only"
	ReadOnly bool `json:"read_only,required"`
	// '#/components/schemas/TagSerializer/properties/value'
	// "$.components.schemas.TagSerializer.properties.value"
	Value string `json:"value,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Key         resp.Field
		ReadOnly    resp.Field
		Value       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tag) RawJSON() string { return r.JSON.raw }
func (r *Tag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/TaskIDsSerializer'
// "$.components.schemas.TaskIDsSerializer"
type TaskIDList struct {
	// '#/components/schemas/TaskIDsSerializer/properties/tasks'
	// "$.components.schemas.TaskIDsSerializer.properties.tasks"
	Tasks []string `json:"tasks,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Tasks       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskIDList) RawJSON() string { return r.JSON.raw }
func (r *TaskIDList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
