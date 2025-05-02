// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"encoding/json"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/resp"
	"github.com/stainless-sdks/gcore-go/shared/constant"
)

// CloudService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudService] method instead.
type CloudService struct {
	Options              []option.RequestOption
	Projects             ProjectService
	Tasks                TaskService
	Regions              RegionService
	Quotas               QuotaService
	Secrets              SecretService
	SSHKeys              SSHKeyService
	IPRanges             IPRangeService
	LoadBalancers        LoadBalancerService
	ReservedFixedIPs     ReservedFixedIPService
	Networks             NetworkService
	Volumes              VolumeService
	FloatingIPs          FloatingIPService
	SecurityGroups       SecurityGroupService
	Users                UserService
	Inference            InferenceService
	PlacementGroups      PlacementGroupService
	Baremetal            BaremetalService
	Registries           RegistryService
	FileShares           FileShareService
	BillingReservations  BillingReservationService
	GPUBaremetalClusters GPUBaremetalClusterService
	Instances            InstanceService
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
	r.LoadBalancers = NewLoadBalancerService(opts...)
	r.ReservedFixedIPs = NewReservedFixedIPService(opts...)
	r.Networks = NewNetworkService(opts...)
	r.Volumes = NewVolumeService(opts...)
	r.FloatingIPs = NewFloatingIPService(opts...)
	r.SecurityGroups = NewSecurityGroupService(opts...)
	r.Users = NewUserService(opts...)
	r.Inference = NewInferenceService(opts...)
	r.PlacementGroups = NewPlacementGroupService(opts...)
	r.Baremetal = NewBaremetalService(opts...)
	r.Registries = NewRegistryService(opts...)
	r.FileShares = NewFileShareService(opts...)
	r.BillingReservations = NewBillingReservationService(opts...)
	r.GPUBaremetalClusters = NewGPUBaremetalClusterService(opts...)
	r.Instances = NewInstanceService(opts...)
	return
}

// '#/components/schemas/RemoteConsoleSerializer'
// "$.components.schemas.RemoteConsoleSerializer"
type Console struct {
	// '#/components/schemas/RemoteConsoleSerializer/properties/remote_console'
	// "$.components.schemas.RemoteConsoleSerializer.properties.remote_console"
	RemoteConsole ConsoleRemoteConsole `json:"remote_console,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		RemoteConsole resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Console) RawJSON() string { return r.JSON.raw }
func (r *Console) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/RemoteConsoleSerializer/properties/remote_console'
// "$.components.schemas.RemoteConsoleSerializer.properties.remote_console"
type ConsoleRemoteConsole struct {
	// '#/components/schemas/RemoteConsoleData/properties/protocol'
	// "$.components.schemas.RemoteConsoleData.properties.protocol"
	Protocol string `json:"protocol,required"`
	// '#/components/schemas/RemoteConsoleData/properties/type'
	// "$.components.schemas.RemoteConsoleData.properties.type"
	Type string `json:"type,required"`
	// '#/components/schemas/RemoteConsoleData/properties/url'
	// "$.components.schemas.RemoteConsoleData.properties.url"
	URL string `json:"url,required" format:"uri"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Protocol    resp.Field
		Type        resp.Field
		URL         resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsoleRemoteConsole) RawJSON() string { return r.JSON.raw }
func (r *ConsoleRemoteConsole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/GetClientProfileSerializer'
// "$.components.schemas.GetClientProfileSerializer"
type DDOSProfile struct {
	// '#/components/schemas/GetClientProfileSerializer/properties/id'
	// "$.components.schemas.GetClientProfileSerializer.properties.id"
	ID int64 `json:"id,required"`
	// '#/components/schemas/GetClientProfileSerializer/properties/profile_template'
	// "$.components.schemas.GetClientProfileSerializer.properties.profile_template"
	ProfileTemplate DDOSProfileTemplate `json:"profile_template,required"`
	// '#/components/schemas/GetClientProfileSerializer/properties/fields'
	// "$.components.schemas.GetClientProfileSerializer.properties.fields"
	Fields []DDOSProfileField `json:"fields"`
	// '#/components/schemas/GetClientProfileSerializer/properties/options/anyOf/0'
	// "$.components.schemas.GetClientProfileSerializer.properties.options.anyOf[0]"
	Options DDOSProfileOptionList `json:"options,nullable"`
	// '#/components/schemas/GetClientProfileSerializer/properties/profile_template_description/anyOf/0'
	// "$.components.schemas.GetClientProfileSerializer.properties.profile_template_description.anyOf[0]"
	ProfileTemplateDescription string `json:"profile_template_description,nullable"`
	// '#/components/schemas/GetClientProfileSerializer/properties/protocols/anyOf/0'
	// "$.components.schemas.GetClientProfileSerializer.properties.protocols.anyOf[0]"
	Protocols []any `json:"protocols,nullable"`
	// '#/components/schemas/GetClientProfileSerializer/properties/site/anyOf/0'
	// "$.components.schemas.GetClientProfileSerializer.properties.site.anyOf[0]"
	Site string `json:"site,nullable"`
	// '#/components/schemas/GetClientProfileSerializer/properties/status/anyOf/0'
	// "$.components.schemas.GetClientProfileSerializer.properties.status.anyOf[0]"
	Status DDOSProfileStatus `json:"status,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                         resp.Field
		ProfileTemplate            resp.Field
		Fields                     resp.Field
		Options                    resp.Field
		ProfileTemplateDescription resp.Field
		Protocols                  resp.Field
		Site                       resp.Field
		Status                     resp.Field
		ExtraFields                map[string]resp.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DDOSProfile) RawJSON() string { return r.JSON.raw }
func (r *DDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ClientProfileFieldSerializer'
// "$.components.schemas.ClientProfileFieldSerializer"
type DDOSProfileField struct {
	// '#/components/schemas/ClientProfileFieldSerializer/properties/id'
	// "$.components.schemas.ClientProfileFieldSerializer.properties.id"
	ID int64 `json:"id,required"`
	// '#/components/schemas/ClientProfileFieldSerializer/properties/default'
	// "$.components.schemas.ClientProfileFieldSerializer.properties['default']"
	Default any `json:"default,required"`
	// '#/components/schemas/ClientProfileFieldSerializer/properties/description'
	// "$.components.schemas.ClientProfileFieldSerializer.properties.description"
	Description string `json:"description,required"`
	// '#/components/schemas/ClientProfileFieldSerializer/properties/field_value'
	// "$.components.schemas.ClientProfileFieldSerializer.properties.field_value"
	FieldValue any `json:"field_value,required"`
	// '#/components/schemas/ClientProfileFieldSerializer/properties/name'
	// "$.components.schemas.ClientProfileFieldSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/ClientProfileFieldSerializer/properties/base_field/anyOf/0'
	// "$.components.schemas.ClientProfileFieldSerializer.properties.base_field.anyOf[0]"
	BaseField int64 `json:"base_field,nullable"`
	// '#/components/schemas/ClientProfileFieldSerializer/properties/field_name/anyOf/0'
	// "$.components.schemas.ClientProfileFieldSerializer.properties.field_name.anyOf[0]"
	FieldName string `json:"field_name,nullable"`
	// '#/components/schemas/ClientProfileFieldSerializer/properties/field_type/anyOf/0'
	// "$.components.schemas.ClientProfileFieldSerializer.properties.field_type.anyOf[0]"
	FieldType string `json:"field_type,nullable"`
	// '#/components/schemas/ClientProfileFieldSerializer/properties/required/anyOf/0'
	// "$.components.schemas.ClientProfileFieldSerializer.properties.required.anyOf[0]"
	Required bool `json:"required,nullable"`
	// '#/components/schemas/ClientProfileFieldSerializer/properties/validation_schema'
	// "$.components.schemas.ClientProfileFieldSerializer.properties.validation_schema"
	ValidationSchema any `json:"validation_schema"`
	// '#/components/schemas/ClientProfileFieldSerializer/properties/value/anyOf/0'
	// "$.components.schemas.ClientProfileFieldSerializer.properties.value.anyOf[0]"
	Value string `json:"value,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID               resp.Field
		Default          resp.Field
		Description      resp.Field
		FieldValue       resp.Field
		Name             resp.Field
		BaseField        resp.Field
		FieldName        resp.Field
		FieldType        resp.Field
		Required         resp.Field
		ValidationSchema resp.Field
		Value            resp.Field
		ExtraFields      map[string]resp.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DDOSProfileField) RawJSON() string { return r.JSON.raw }
func (r *DDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ProfileOptionsSerializer'
// "$.components.schemas.ProfileOptionsSerializer"
type DDOSProfileOptionList struct {
	// '#/components/schemas/ProfileOptionsSerializer/properties/active/anyOf/0'
	// "$.components.schemas.ProfileOptionsSerializer.properties.active.anyOf[0]"
	Active bool `json:"active,nullable"`
	// '#/components/schemas/ProfileOptionsSerializer/properties/bgp/anyOf/0'
	// "$.components.schemas.ProfileOptionsSerializer.properties.bgp.anyOf[0]"
	Bgp bool `json:"bgp,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Active      resp.Field
		Bgp         resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DDOSProfileOptionList) RawJSON() string { return r.JSON.raw }
func (r *DDOSProfileOptionList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/DdosProfileStatusSerializer'
// "$.components.schemas.DdosProfileStatusSerializer"
type DDOSProfileStatus struct {
	// '#/components/schemas/DdosProfileStatusSerializer/properties/error_description'
	// "$.components.schemas.DdosProfileStatusSerializer.properties.error_description"
	ErrorDescription string `json:"error_description,required"`
	// '#/components/schemas/DdosProfileStatusSerializer/properties/status'
	// "$.components.schemas.DdosProfileStatusSerializer.properties.status"
	Status string `json:"status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ErrorDescription resp.Field
		Status           resp.Field
		ExtraFields      map[string]resp.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DDOSProfileStatus) RawJSON() string { return r.JSON.raw }
func (r *DDOSProfileStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ClientProfileTemplateSerializer'
// "$.components.schemas.ClientProfileTemplateSerializer"
type DDOSProfileTemplate struct {
	// '#/components/schemas/ClientProfileTemplateSerializer/properties/id'
	// "$.components.schemas.ClientProfileTemplateSerializer.properties.id"
	ID int64 `json:"id,required"`
	// '#/components/schemas/ClientProfileTemplateSerializer/properties/name'
	// "$.components.schemas.ClientProfileTemplateSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/ClientProfileTemplateSerializer/properties/description/anyOf/0'
	// "$.components.schemas.ClientProfileTemplateSerializer.properties.description.anyOf[0]"
	Description string `json:"description,nullable"`
	// '#/components/schemas/ClientProfileTemplateSerializer/properties/fields'
	// "$.components.schemas.ClientProfileTemplateSerializer.properties.fields"
	Fields []DDOSProfileTemplateField `json:"fields"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Name        resp.Field
		Description resp.Field
		Fields      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DDOSProfileTemplate) RawJSON() string { return r.JSON.raw }
func (r *DDOSProfileTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ClientProfileTemplateFieldSerializer'
// "$.components.schemas.ClientProfileTemplateFieldSerializer"
type DDOSProfileTemplateField struct {
	// '#/components/schemas/ClientProfileTemplateFieldSerializer/properties/id'
	// "$.components.schemas.ClientProfileTemplateFieldSerializer.properties.id"
	ID int64 `json:"id,required"`
	// '#/components/schemas/ClientProfileTemplateFieldSerializer/properties/name'
	// "$.components.schemas.ClientProfileTemplateFieldSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/ClientProfileTemplateFieldSerializer/properties/default/anyOf/0'
	// "$.components.schemas.ClientProfileTemplateFieldSerializer.properties['default'].anyOf[0]"
	Default string `json:"default,nullable"`
	// '#/components/schemas/ClientProfileTemplateFieldSerializer/properties/description/anyOf/0'
	// "$.components.schemas.ClientProfileTemplateFieldSerializer.properties.description.anyOf[0]"
	Description string `json:"description,nullable"`
	// '#/components/schemas/ClientProfileTemplateFieldSerializer/properties/field_type/anyOf/0'
	// "$.components.schemas.ClientProfileTemplateFieldSerializer.properties.field_type.anyOf[0]"
	FieldType string `json:"field_type,nullable"`
	// '#/components/schemas/ClientProfileTemplateFieldSerializer/properties/required/anyOf/0'
	// "$.components.schemas.ClientProfileTemplateFieldSerializer.properties.required.anyOf[0]"
	Required bool `json:"required,nullable"`
	// '#/components/schemas/ClientProfileTemplateFieldSerializer/properties/validation_schema'
	// "$.components.schemas.ClientProfileTemplateFieldSerializer.properties.validation_schema"
	ValidationSchema any `json:"validation_schema"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID               resp.Field
		Name             resp.Field
		Default          resp.Field
		Description      resp.Field
		FieldType        resp.Field
		Required         resp.Field
		ValidationSchema resp.Field
		ExtraFields      map[string]resp.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DDOSProfileTemplateField) RawJSON() string { return r.JSON.raw }
func (r *DDOSProfileTemplateField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/FlavorHardwareDescriptionSerializer'
// "$.components.schemas.FlavorHardwareDescriptionSerializer"
type FlavorHardwareDescription struct {
	// '#/components/schemas/FlavorHardwareDescriptionSerializer/properties/cpu/anyOf/0'
	// "$.components.schemas.FlavorHardwareDescriptionSerializer.properties.cpu.anyOf[0]"
	CPU string `json:"cpu,nullable"`
	// '#/components/schemas/FlavorHardwareDescriptionSerializer/properties/disk/anyOf/0'
	// "$.components.schemas.FlavorHardwareDescriptionSerializer.properties.disk.anyOf[0]"
	Disk string `json:"disk,nullable"`
	// '#/components/schemas/FlavorHardwareDescriptionSerializer/properties/ephemeral/anyOf/0'
	// "$.components.schemas.FlavorHardwareDescriptionSerializer.properties.ephemeral.anyOf[0]"
	Ephemeral string `json:"ephemeral,nullable"`
	// '#/components/schemas/FlavorHardwareDescriptionSerializer/properties/gpu/anyOf/0'
	// "$.components.schemas.FlavorHardwareDescriptionSerializer.properties.gpu.anyOf[0]"
	GPU string `json:"gpu,nullable"`
	// '#/components/schemas/FlavorHardwareDescriptionSerializer/properties/ipu/anyOf/0'
	// "$.components.schemas.FlavorHardwareDescriptionSerializer.properties.ipu.anyOf[0]"
	Ipu string `json:"ipu,nullable"`
	// '#/components/schemas/FlavorHardwareDescriptionSerializer/properties/network/anyOf/0'
	// "$.components.schemas.FlavorHardwareDescriptionSerializer.properties.network.anyOf[0]"
	Network string `json:"network,nullable"`
	// '#/components/schemas/FlavorHardwareDescriptionSerializer/properties/poplar_count/anyOf/0'
	// "$.components.schemas.FlavorHardwareDescriptionSerializer.properties.poplar_count.anyOf[0]"
	PoplarCount int64 `json:"poplar_count,nullable"`
	// '#/components/schemas/FlavorHardwareDescriptionSerializer/properties/ram/anyOf/0'
	// "$.components.schemas.FlavorHardwareDescriptionSerializer.properties.ram.anyOf[0]"
	Ram string `json:"ram,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CPU         resp.Field
		Disk        resp.Field
		Ephemeral   resp.Field
		GPU         resp.Field
		Ipu         resp.Field
		Network     resp.Field
		PoplarCount resp.Field
		Ram         resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *FlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/FloatingIPSerializer'
// "$.components.schemas.FloatingIPSerializer"
type FloatingIP struct {
	// '#/components/schemas/FloatingIPSerializer/properties/id/anyOf/0'
	// "$.components.schemas.FloatingIPSerializer.properties.id.anyOf[0]"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/FloatingIPSerializer/properties/created_at'
	// "$.components.schemas.FloatingIPSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/FloatingIPSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.FloatingIPSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,required" format:"uuid4"`
	// '#/components/schemas/FloatingIPSerializer/properties/dns_domain/anyOf/0'
	// "$.components.schemas.FloatingIPSerializer.properties.dns_domain.anyOf[0]"
	//
	// Deprecated: deprecated
	DNSDomain string `json:"dns_domain,required"`
	// '#/components/schemas/FloatingIPSerializer/properties/dns_name/anyOf/0'
	// "$.components.schemas.FloatingIPSerializer.properties.dns_name.anyOf[0]"
	//
	// Deprecated: deprecated
	DNSName string `json:"dns_name,required"`
	// '#/components/schemas/FloatingIPSerializer/properties/fixed_ip_address/anyOf/0'
	// "$.components.schemas.FloatingIPSerializer.properties.fixed_ip_address.anyOf[0]"
	FixedIPAddress string `json:"fixed_ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/FloatingIPSerializer/properties/floating_ip_address/anyOf/0'
	// "$.components.schemas.FloatingIPSerializer.properties.floating_ip_address.anyOf[0]"
	FloatingIPAddress string `json:"floating_ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/FloatingIPSerializer/properties/port_id/anyOf/0'
	// "$.components.schemas.FloatingIPSerializer.properties.port_id.anyOf[0]"
	PortID string `json:"port_id,required" format:"uuid4"`
	// '#/components/schemas/FloatingIPSerializer/properties/project_id'
	// "$.components.schemas.FloatingIPSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/FloatingIPSerializer/properties/region'
	// "$.components.schemas.FloatingIPSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/FloatingIPSerializer/properties/region_id'
	// "$.components.schemas.FloatingIPSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/FloatingIPSerializer/properties/router_id/anyOf/0'
	// "$.components.schemas.FloatingIPSerializer.properties.router_id.anyOf[0]"
	RouterID string `json:"router_id,required" format:"uuid4"`
	// '#/components/schemas/FloatingIPSerializer/properties/status/anyOf/0'
	// "$.components.schemas.FloatingIPSerializer.properties.status.anyOf[0]"
	//
	// Any of "ACTIVE", "DOWN", "ERROR".
	Status FloatingIPStatus `json:"status,required"`
	// '#/components/schemas/FloatingIPSerializer/properties/subnet_id/anyOf/0'
	// "$.components.schemas.FloatingIPSerializer.properties.subnet_id.anyOf[0]"
	//
	// Deprecated: deprecated
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// '#/components/schemas/FloatingIPSerializer/properties/tags'
	// "$.components.schemas.FloatingIPSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/FloatingIPSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.FloatingIPSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,required" format:"uuid4"`
	// '#/components/schemas/FloatingIPSerializer/properties/updated_at'
	// "$.components.schemas.FloatingIPSerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                resp.Field
		CreatedAt         resp.Field
		CreatorTaskID     resp.Field
		DNSDomain         resp.Field
		DNSName           resp.Field
		FixedIPAddress    resp.Field
		FloatingIPAddress resp.Field
		PortID            resp.Field
		ProjectID         resp.Field
		Region            resp.Field
		RegionID          resp.Field
		RouterID          resp.Field
		Status            resp.Field
		SubnetID          resp.Field
		Tags              resp.Field
		TaskID            resp.Field
		UpdatedAt         resp.Field
		ExtraFields       map[string]resp.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatingIP) RawJSON() string { return r.JSON.raw }
func (r *FloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/FloatingIPStatus' "$.components.schemas.FloatingIPStatus"
type FloatingIPStatus string

const (
	FloatingIPStatusActive FloatingIPStatus = "ACTIVE"
	FloatingIPStatusDown   FloatingIPStatus = "DOWN"
	FloatingIPStatusError  FloatingIPStatus = "ERROR"
)

// '#/components/schemas/GPUClusterServerSerializer'
// "$.components.schemas.GPUClusterServerSerializer"
type GPUClusterServer struct {
	// '#/components/schemas/GPUClusterServerSerializer/properties/addresses'
	// "$.components.schemas.GPUClusterServerSerializer.properties.addresses"
	Addresses map[string][]GPUClusterServerAddressUnion `json:"addresses,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/blackhole_ports'
	// "$.components.schemas.GPUClusterServerSerializer.properties.blackhole_ports"
	BlackholePorts []GPUClusterServerBlackholePort `json:"blackhole_ports,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.GPUClusterServerSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/ddos_profile/anyOf/0'
	// "$.components.schemas.GPUClusterServerSerializer.properties.ddos_profile.anyOf[0]"
	DDOSProfile DDOSProfile `json:"ddos_profile,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/fixed_ip_assignments/anyOf/0'
	// "$.components.schemas.GPUClusterServerSerializer.properties.fixed_ip_assignments.anyOf[0]"
	FixedIPAssignments []GPUClusterServerFixedIPAssignment `json:"fixed_ip_assignments,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/flavor'
	// "$.components.schemas.GPUClusterServerSerializer.properties.flavor"
	Flavor GPUClusterServerFlavor `json:"flavor,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/instance_created'
	// "$.components.schemas.GPUClusterServerSerializer.properties.instance_created"
	InstanceCreated time.Time `json:"instance_created,required" format:"date-time"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/instance_description/anyOf/0'
	// "$.components.schemas.GPUClusterServerSerializer.properties.instance_description.anyOf[0]"
	InstanceDescription string `json:"instance_description,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/instance_id'
	// "$.components.schemas.GPUClusterServerSerializer.properties.instance_id"
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/instance_isolation/anyOf/0'
	// "$.components.schemas.GPUClusterServerSerializer.properties.instance_isolation.anyOf[0]"
	InstanceIsolation GPUClusterServerInstanceIsolation `json:"instance_isolation,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/instance_name'
	// "$.components.schemas.GPUClusterServerSerializer.properties.instance_name"
	InstanceName string `json:"instance_name,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/keypair_name/anyOf/0'
	// "$.components.schemas.GPUClusterServerSerializer.properties.keypair_name.anyOf[0]"
	KeypairName string `json:"keypair_name,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/project_id'
	// "$.components.schemas.GPUClusterServerSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/region'
	// "$.components.schemas.GPUClusterServerSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/region_id'
	// "$.components.schemas.GPUClusterServerSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/security_groups'
	// "$.components.schemas.GPUClusterServerSerializer.properties.security_groups"
	SecurityGroups []GPUClusterServerSecurityGroup `json:"security_groups,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/status'
	// "$.components.schemas.GPUClusterServerSerializer.properties.status"
	//
	// Any of "ACTIVE", "BUILD", "DELETED", "ERROR", "HARD_REBOOT", "MIGRATING",
	// "PASSWORD", "PAUSED", "REBOOT", "REBUILD", "RESCUE", "RESIZE", "REVERT_RESIZE",
	// "SHELVED", "SHELVED_OFFLOADED", "SHUTOFF", "SOFT_DELETED", "SUSPENDED",
	// "UNKNOWN", "VERIFY_RESIZE".
	Status GPUClusterServerStatus `json:"status,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/tags'
	// "$.components.schemas.GPUClusterServerSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.GPUClusterServerSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/task_state/anyOf/0'
	// "$.components.schemas.GPUClusterServerSerializer.properties.task_state.anyOf[0]"
	TaskState string `json:"task_state,required"`
	// '#/components/schemas/GPUClusterServerSerializer/properties/vm_state'
	// "$.components.schemas.GPUClusterServerSerializer.properties.vm_state"
	//
	// Any of "active", "building", "deleted", "error", "paused", "rescued", "resized",
	// "shelved", "shelved_offloaded", "soft-deleted", "stopped", "suspended".
	VmState GPUClusterServerVmState `json:"vm_state,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Addresses           resp.Field
		BlackholePorts      resp.Field
		CreatorTaskID       resp.Field
		DDOSProfile         resp.Field
		FixedIPAssignments  resp.Field
		Flavor              resp.Field
		InstanceCreated     resp.Field
		InstanceDescription resp.Field
		InstanceID          resp.Field
		InstanceIsolation   resp.Field
		InstanceName        resp.Field
		KeypairName         resp.Field
		ProjectID           resp.Field
		Region              resp.Field
		RegionID            resp.Field
		SecurityGroups      resp.Field
		Status              resp.Field
		Tags                resp.Field
		TaskID              resp.Field
		TaskState           resp.Field
		VmState             resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUClusterServer) RawJSON() string { return r.JSON.raw }
func (r *GPUClusterServer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUClusterServerAddressUnion contains all possible properties and values from
// [GPUClusterServerAddressFloatingIPAddress],
// [GPUClusterServerAddressFixedIPAddressShort],
// [GPUClusterServerAddressFixedIPAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type GPUClusterServerAddressUnion struct {
	Addr          string `json:"addr"`
	Type          string `json:"type"`
	InterfaceName string `json:"interface_name"`
	// This field is from variant [GPUClusterServerAddressFixedIPAddress].
	SubnetID string `json:"subnet_id"`
	// This field is from variant [GPUClusterServerAddressFixedIPAddress].
	SubnetName string `json:"subnet_name"`
	JSON       struct {
		Addr          resp.Field
		Type          resp.Field
		InterfaceName resp.Field
		SubnetID      resp.Field
		SubnetName    resp.Field
		raw           string
	} `json:"-"`
}

func (u GPUClusterServerAddressUnion) AsFloatingIPAddress() (v GPUClusterServerAddressFloatingIPAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GPUClusterServerAddressUnion) AsFixedIPAddressShort() (v GPUClusterServerAddressFixedIPAddressShort) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GPUClusterServerAddressUnion) AsFixedIPAddress() (v GPUClusterServerAddressFixedIPAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GPUClusterServerAddressUnion) RawJSON() string { return u.JSON.raw }

func (r *GPUClusterServerAddressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/GPUClusterServerSerializer/properties/addresses/additionalProperties/items/anyOf/0'
// "$.components.schemas.GPUClusterServerSerializer.properties.addresses.additionalProperties.items.anyOf[0]"
type GPUClusterServerAddressFloatingIPAddress struct {
	// '#/components/schemas/InstanceFloatingAddressSerializer/properties/addr'
	// "$.components.schemas.InstanceFloatingAddressSerializer.properties.addr"
	Addr string `json:"addr,required"`
	// '#/components/schemas/InstanceFloatingAddressSerializer/properties/type'
	// "$.components.schemas.InstanceFloatingAddressSerializer.properties.type"
	Type constant.Floating `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Addr        resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUClusterServerAddressFloatingIPAddress) RawJSON() string { return r.JSON.raw }
func (r *GPUClusterServerAddressFloatingIPAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/GPUClusterServerSerializer/properties/addresses/additionalProperties/items/anyOf/1'
// "$.components.schemas.GPUClusterServerSerializer.properties.addresses.additionalProperties.items.anyOf[1]"
type GPUClusterServerAddressFixedIPAddressShort struct {
	// '#/components/schemas/InstanceFixedAddressShortSerializer/properties/addr'
	// "$.components.schemas.InstanceFixedAddressShortSerializer.properties.addr"
	Addr string `json:"addr,required"`
	// '#/components/schemas/InstanceFixedAddressShortSerializer/properties/interface_name/anyOf/0'
	// "$.components.schemas.InstanceFixedAddressShortSerializer.properties.interface_name.anyOf[0]"
	InterfaceName string `json:"interface_name,required"`
	// '#/components/schemas/InstanceFixedAddressShortSerializer/properties/type'
	// "$.components.schemas.InstanceFixedAddressShortSerializer.properties.type"
	Type constant.Fixed `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Addr          resp.Field
		InterfaceName resp.Field
		Type          resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUClusterServerAddressFixedIPAddressShort) RawJSON() string { return r.JSON.raw }
func (r *GPUClusterServerAddressFixedIPAddressShort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/GPUClusterServerSerializer/properties/addresses/additionalProperties/items/anyOf/2'
// "$.components.schemas.GPUClusterServerSerializer.properties.addresses.additionalProperties.items.anyOf[2]"
type GPUClusterServerAddressFixedIPAddress struct {
	// '#/components/schemas/InstanceFixedAddressSerializer/properties/addr'
	// "$.components.schemas.InstanceFixedAddressSerializer.properties.addr"
	Addr string `json:"addr,required"`
	// '#/components/schemas/InstanceFixedAddressSerializer/properties/interface_name/anyOf/0'
	// "$.components.schemas.InstanceFixedAddressSerializer.properties.interface_name.anyOf[0]"
	InterfaceName string `json:"interface_name,required"`
	// '#/components/schemas/InstanceFixedAddressSerializer/properties/subnet_id'
	// "$.components.schemas.InstanceFixedAddressSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required"`
	// '#/components/schemas/InstanceFixedAddressSerializer/properties/subnet_name'
	// "$.components.schemas.InstanceFixedAddressSerializer.properties.subnet_name"
	SubnetName string `json:"subnet_name,required"`
	// '#/components/schemas/InstanceFixedAddressSerializer/properties/type'
	// "$.components.schemas.InstanceFixedAddressSerializer.properties.type"
	Type constant.Fixed `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Addr          resp.Field
		InterfaceName resp.Field
		SubnetID      resp.Field
		SubnetName    resp.Field
		Type          resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUClusterServerAddressFixedIPAddress) RawJSON() string { return r.JSON.raw }
func (r *GPUClusterServerAddressFixedIPAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/GPUClusterServerSerializer/properties/blackhole_ports/items'
// "$.components.schemas.GPUClusterServerSerializer.properties.blackhole_ports.items"
type GPUClusterServerBlackholePort struct {
	// '#/components/schemas/BlackholePortSerializer/properties/AlarmEnd'
	// "$.components.schemas.BlackholePortSerializer.properties.AlarmEnd"
	AlarmEnd time.Time `json:"AlarmEnd,required" format:"date-time"`
	// '#/components/schemas/BlackholePortSerializer/properties/AlarmStart'
	// "$.components.schemas.BlackholePortSerializer.properties.AlarmStart"
	AlarmStart time.Time `json:"AlarmStart,required" format:"date-time"`
	// '#/components/schemas/BlackholePortSerializer/properties/AlarmState'
	// "$.components.schemas.BlackholePortSerializer.properties.AlarmState"
	//
	// Any of "ACK_REQ", "ALARM", "ARCHIVED", "CLEAR", "CLEARING", "CLEARING_FAIL",
	// "END_GRACE", "END_WAIT", "MANUAL_CLEAR", "MANUAL_CLEARING",
	// "MANUAL_CLEARING_FAIL", "MANUAL_MITIGATING", "MANUAL_STARTING",
	// "MANUAL_STARTING_FAIL", "MITIGATING", "STARTING", "STARTING_FAIL", "START_WAIT",
	// "ack_req", "alarm", "archived", "clear", "clearing", "clearing_fail",
	// "end_grace", "end_wait", "manual_clear", "manual_clearing",
	// "manual_clearing_fail", "manual_mitigating", "manual_starting",
	// "manual_starting_fail", "mitigating", "start_wait", "starting", "starting_fail".
	AlarmState string `json:"AlarmState,required"`
	// '#/components/schemas/BlackholePortSerializer/properties/AlertDuration'
	// "$.components.schemas.BlackholePortSerializer.properties.AlertDuration"
	AlertDuration string `json:"AlertDuration,required"`
	// '#/components/schemas/BlackholePortSerializer/properties/DestinationIP'
	// "$.components.schemas.BlackholePortSerializer.properties.DestinationIP"
	DestinationIP string `json:"DestinationIP,required"`
	// '#/components/schemas/BlackholePortSerializer/properties/ID'
	// "$.components.schemas.BlackholePortSerializer.properties.ID"
	ID int64 `json:"ID,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AlarmEnd      resp.Field
		AlarmStart    resp.Field
		AlarmState    resp.Field
		AlertDuration resp.Field
		DestinationIP resp.Field
		ID            resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUClusterServerBlackholePort) RawJSON() string { return r.JSON.raw }
func (r *GPUClusterServerBlackholePort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/GPUClusterServerSerializer/properties/fixed_ip_assignments/anyOf/0/items'
// "$.components.schemas.GPUClusterServerSerializer.properties.fixed_ip_assignments.anyOf[0].items"
type GPUClusterServerFixedIPAssignment struct {
	// '#/components/schemas/IpAssignmentsSerializer/properties/external'
	// "$.components.schemas.IpAssignmentsSerializer.properties.external"
	External bool `json:"external,required"`
	// '#/components/schemas/IpAssignmentsSerializer/properties/ip_address'
	// "$.components.schemas.IpAssignmentsSerializer.properties.ip_address"
	IPAddress string `json:"ip_address,required"`
	// '#/components/schemas/IpAssignmentsSerializer/properties/subnet_id'
	// "$.components.schemas.IpAssignmentsSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		External    resp.Field
		IPAddress   resp.Field
		SubnetID    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUClusterServerFixedIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *GPUClusterServerFixedIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/GPUClusterServerSerializer/properties/flavor'
// "$.components.schemas.GPUClusterServerSerializer.properties.flavor"
type GPUClusterServerFlavor struct {
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/architecture'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.architecture"
	Architecture string `json:"architecture,required"`
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/flavor_id'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.flavor_id"
	FlavorID string `json:"flavor_id,required"`
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/flavor_name'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.flavor_name"
	FlavorName string `json:"flavor_name,required"`
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/hardware_description'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.hardware_description"
	HardwareDescription GPUClusterServerFlavorHardwareDescription `json:"hardware_description,required"`
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/os_type'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.os_type"
	OsType string `json:"os_type,required"`
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/ram'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.ram"
	Ram int64 `json:"ram,required"`
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/resource_class'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.resource_class"
	ResourceClass string `json:"resource_class,required"`
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/vcpus'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.vcpus"
	Vcpus int64 `json:"vcpus,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Architecture        resp.Field
		FlavorID            resp.Field
		FlavorName          resp.Field
		HardwareDescription resp.Field
		OsType              resp.Field
		Ram                 resp.Field
		ResourceClass       resp.Field
		Vcpus               resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUClusterServerFlavor) RawJSON() string { return r.JSON.raw }
func (r *GPUClusterServerFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/hardware_description'
// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.hardware_description"
type GPUClusterServerFlavorHardwareDescription struct {
	// '#/components/schemas/DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer/properties/cpu'
	// "$.components.schemas.DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer.properties.cpu"
	CPU string `json:"cpu,required"`
	// '#/components/schemas/DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer/properties/disk'
	// "$.components.schemas.DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer.properties.disk"
	Disk string `json:"disk,required"`
	// '#/components/schemas/DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer/properties/gpu'
	// "$.components.schemas.DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer.properties.gpu"
	GPU string `json:"gpu,required"`
	// '#/components/schemas/DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer/properties/license'
	// "$.components.schemas.DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer.properties.license"
	License string `json:"license,required"`
	// '#/components/schemas/DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer/properties/network'
	// "$.components.schemas.DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer.properties.network"
	Network string `json:"network,required"`
	// '#/components/schemas/DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer/properties/ram'
	// "$.components.schemas.DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer.properties.ram"
	Ram string `json:"ram,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CPU         resp.Field
		Disk        resp.Field
		GPU         resp.Field
		License     resp.Field
		Network     resp.Field
		Ram         resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUClusterServerFlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *GPUClusterServerFlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/GPUClusterServerSerializer/properties/instance_isolation/anyOf/0'
// "$.components.schemas.GPUClusterServerSerializer.properties.instance_isolation.anyOf[0]"
type GPUClusterServerInstanceIsolation struct {
	// '#/components/schemas/IsolationSerializer/properties/reason/anyOf/0'
	// "$.components.schemas.IsolationSerializer.properties.reason.anyOf[0]"
	Reason string `json:"reason,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Reason      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUClusterServerInstanceIsolation) RawJSON() string { return r.JSON.raw }
func (r *GPUClusterServerInstanceIsolation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/GPUClusterServerSerializer/properties/security_groups/items'
// "$.components.schemas.GPUClusterServerSerializer.properties.security_groups.items"
type GPUClusterServerSecurityGroup struct {
	// '#/components/schemas/NameSerializerPydantic/properties/name'
	// "$.components.schemas.NameSerializerPydantic.properties.name"
	Name string `json:"name,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUClusterServerSecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *GPUClusterServerSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/GPUClusterServerSerializer/properties/status'
// "$.components.schemas.GPUClusterServerSerializer.properties.status"
type GPUClusterServerStatus string

const (
	GPUClusterServerStatusActive           GPUClusterServerStatus = "ACTIVE"
	GPUClusterServerStatusBuild            GPUClusterServerStatus = "BUILD"
	GPUClusterServerStatusDeleted          GPUClusterServerStatus = "DELETED"
	GPUClusterServerStatusError            GPUClusterServerStatus = "ERROR"
	GPUClusterServerStatusHardReboot       GPUClusterServerStatus = "HARD_REBOOT"
	GPUClusterServerStatusMigrating        GPUClusterServerStatus = "MIGRATING"
	GPUClusterServerStatusPassword         GPUClusterServerStatus = "PASSWORD"
	GPUClusterServerStatusPaused           GPUClusterServerStatus = "PAUSED"
	GPUClusterServerStatusReboot           GPUClusterServerStatus = "REBOOT"
	GPUClusterServerStatusRebuild          GPUClusterServerStatus = "REBUILD"
	GPUClusterServerStatusRescue           GPUClusterServerStatus = "RESCUE"
	GPUClusterServerStatusResize           GPUClusterServerStatus = "RESIZE"
	GPUClusterServerStatusRevertResize     GPUClusterServerStatus = "REVERT_RESIZE"
	GPUClusterServerStatusShelved          GPUClusterServerStatus = "SHELVED"
	GPUClusterServerStatusShelvedOffloaded GPUClusterServerStatus = "SHELVED_OFFLOADED"
	GPUClusterServerStatusShutoff          GPUClusterServerStatus = "SHUTOFF"
	GPUClusterServerStatusSoftDeleted      GPUClusterServerStatus = "SOFT_DELETED"
	GPUClusterServerStatusSuspended        GPUClusterServerStatus = "SUSPENDED"
	GPUClusterServerStatusUnknown          GPUClusterServerStatus = "UNKNOWN"
	GPUClusterServerStatusVerifyResize     GPUClusterServerStatus = "VERIFY_RESIZE"
)

// '#/components/schemas/GPUClusterServerSerializer/properties/vm_state'
// "$.components.schemas.GPUClusterServerSerializer.properties.vm_state"
type GPUClusterServerVmState string

const (
	GPUClusterServerVmStateActive           GPUClusterServerVmState = "active"
	GPUClusterServerVmStateBuilding         GPUClusterServerVmState = "building"
	GPUClusterServerVmStateDeleted          GPUClusterServerVmState = "deleted"
	GPUClusterServerVmStateError            GPUClusterServerVmState = "error"
	GPUClusterServerVmStatePaused           GPUClusterServerVmState = "paused"
	GPUClusterServerVmStateRescued          GPUClusterServerVmState = "rescued"
	GPUClusterServerVmStateResized          GPUClusterServerVmState = "resized"
	GPUClusterServerVmStateShelved          GPUClusterServerVmState = "shelved"
	GPUClusterServerVmStateShelvedOffloaded GPUClusterServerVmState = "shelved_offloaded"
	GPUClusterServerVmStateSoftDeleted      GPUClusterServerVmState = "soft-deleted"
	GPUClusterServerVmStateStopped          GPUClusterServerVmState = "stopped"
	GPUClusterServerVmStateSuspended        GPUClusterServerVmState = "suspended"
)

// '#/components/schemas/GPUClusterServerCollectionSerializer'
// "$.components.schemas.GPUClusterServerCollectionSerializer"
type GPUClusterServerList struct {
	// '#/components/schemas/GPUClusterServerCollectionSerializer/properties/count'
	// "$.components.schemas.GPUClusterServerCollectionSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/GPUClusterServerCollectionSerializer/properties/results'
	// "$.components.schemas.GPUClusterServerCollectionSerializer.properties.results"
	Results []GPUClusterServer `json:"results,required"`
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
func (r GPUClusterServerList) RawJSON() string { return r.JSON.raw }
func (r *GPUClusterServerList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/GpuImageSerializer'
// "$.components.schemas.GpuImageSerializer"
type GPUImage struct {
	// '#/components/schemas/GpuImageSerializer/properties/id'
	// "$.components.schemas.GpuImageSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/GpuImageSerializer/properties/created_at'
	// "$.components.schemas.GpuImageSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/GpuImageSerializer/properties/min_disk'
	// "$.components.schemas.GpuImageSerializer.properties.min_disk"
	MinDisk int64 `json:"min_disk,required"`
	// '#/components/schemas/GpuImageSerializer/properties/min_ram'
	// "$.components.schemas.GpuImageSerializer.properties.min_ram"
	MinRam int64 `json:"min_ram,required"`
	// '#/components/schemas/GpuImageSerializer/properties/name'
	// "$.components.schemas.GpuImageSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/GpuImageSerializer/properties/status'
	// "$.components.schemas.GpuImageSerializer.properties.status"
	Status string `json:"status,required"`
	// '#/components/schemas/GpuImageSerializer/properties/tags'
	// "$.components.schemas.GpuImageSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/GpuImageSerializer/properties/updated_at'
	// "$.components.schemas.GpuImageSerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// '#/components/schemas/GpuImageSerializer/properties/visibility'
	// "$.components.schemas.GpuImageSerializer.properties.visibility"
	Visibility string `json:"visibility,required"`
	// '#/components/schemas/GpuImageSerializer/properties/architecture/anyOf/0'
	// "$.components.schemas.GpuImageSerializer.properties.architecture.anyOf[0]"
	Architecture string `json:"architecture,nullable"`
	// '#/components/schemas/GpuImageSerializer/properties/os_distro/anyOf/0'
	// "$.components.schemas.GpuImageSerializer.properties.os_distro.anyOf[0]"
	OsDistro string `json:"os_distro,nullable"`
	// '#/components/schemas/GpuImageSerializer/properties/os_type/anyOf/0'
	// "$.components.schemas.GpuImageSerializer.properties.os_type.anyOf[0]"
	OsType string `json:"os_type,nullable"`
	// '#/components/schemas/GpuImageSerializer/properties/os_version/anyOf/0'
	// "$.components.schemas.GpuImageSerializer.properties.os_version.anyOf[0]"
	OsVersion string `json:"os_version,nullable"`
	// '#/components/schemas/GpuImageSerializer/properties/size'
	// "$.components.schemas.GpuImageSerializer.properties.size"
	Size int64 `json:"size"`
	// '#/components/schemas/GpuImageSerializer/properties/ssh_key/anyOf/0'
	// "$.components.schemas.GpuImageSerializer.properties.ssh_key.anyOf[0]"
	SSHKey string `json:"ssh_key,nullable"`
	// '#/components/schemas/GpuImageSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.GpuImageSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID           resp.Field
		CreatedAt    resp.Field
		MinDisk      resp.Field
		MinRam       resp.Field
		Name         resp.Field
		Status       resp.Field
		Tags         resp.Field
		UpdatedAt    resp.Field
		Visibility   resp.Field
		Architecture resp.Field
		OsDistro     resp.Field
		OsType       resp.Field
		OsVersion    resp.Field
		Size         resp.Field
		SSHKey       resp.Field
		TaskID       resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUImage) RawJSON() string { return r.JSON.raw }
func (r *GPUImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ListGpuImageSerializer'
// "$.components.schemas.ListGpuImageSerializer"
type GPUImageList struct {
	// '#/components/schemas/ListGpuImageSerializer/properties/count'
	// "$.components.schemas.ListGpuImageSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/ListGpuImageSerializer/properties/results'
	// "$.components.schemas.ListGpuImageSerializer.properties.results"
	Results []GPUImage `json:"results,required"`
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
func (r GPUImageList) RawJSON() string { return r.JSON.raw }
func (r *GPUImageList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ImageSerializer' "$.components.schemas.ImageSerializer"
type Image struct {
	// '#/components/schemas/ImageSerializer/properties/id'
	// "$.components.schemas.ImageSerializer.properties.id"
	ID string `json:"id,required"`
	// '#/components/schemas/ImageSerializer/properties/created_at'
	// "$.components.schemas.ImageSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/ImageSerializer/properties/disk_format'
	// "$.components.schemas.ImageSerializer.properties.disk_format"
	DiskFormat string `json:"disk_format,required"`
	// '#/components/schemas/ImageSerializer/properties/min_disk'
	// "$.components.schemas.ImageSerializer.properties.min_disk"
	MinDisk int64 `json:"min_disk,required"`
	// '#/components/schemas/ImageSerializer/properties/min_ram'
	// "$.components.schemas.ImageSerializer.properties.min_ram"
	MinRam int64 `json:"min_ram,required"`
	// '#/components/schemas/ImageSerializer/properties/name'
	// "$.components.schemas.ImageSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/ImageSerializer/properties/os_distro'
	// "$.components.schemas.ImageSerializer.properties.os_distro"
	OsDistro string `json:"os_distro,required"`
	// '#/components/schemas/ImageSerializer/properties/os_type'
	// "$.components.schemas.ImageSerializer.properties.os_type"
	//
	// Any of "linux", "windows".
	OsType ImageOsType `json:"os_type,required"`
	// '#/components/schemas/ImageSerializer/properties/os_version'
	// "$.components.schemas.ImageSerializer.properties.os_version"
	OsVersion string `json:"os_version,required"`
	// '#/components/schemas/ImageSerializer/properties/project_id'
	// "$.components.schemas.ImageSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/ImageSerializer/properties/region'
	// "$.components.schemas.ImageSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/ImageSerializer/properties/region_id'
	// "$.components.schemas.ImageSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/ImageSerializer/properties/size'
	// "$.components.schemas.ImageSerializer.properties.size"
	Size int64 `json:"size,required"`
	// '#/components/schemas/ImageSerializer/properties/status'
	// "$.components.schemas.ImageSerializer.properties.status"
	Status string `json:"status,required"`
	// '#/components/schemas/ImageSerializer/properties/tags'
	// "$.components.schemas.ImageSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/ImageSerializer/properties/updated_at'
	// "$.components.schemas.ImageSerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// '#/components/schemas/ImageSerializer/properties/visibility'
	// "$.components.schemas.ImageSerializer.properties.visibility"
	Visibility string `json:"visibility,required"`
	// '#/components/schemas/ImageSerializer/properties/architecture'
	// "$.components.schemas.ImageSerializer.properties.architecture"
	//
	// Any of "aarch64", "x86_64".
	Architecture ImageArchitecture `json:"architecture"`
	// '#/components/schemas/ImageSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,nullable"`
	// '#/components/schemas/ImageSerializer/properties/description/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.description.anyOf[0]"
	Description string `json:"description,nullable"`
	// '#/components/schemas/ImageSerializer/properties/display_order/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.display_order.anyOf[0]"
	DisplayOrder int64 `json:"display_order,nullable"`
	// '#/components/schemas/ImageSerializer/properties/hw_firmware_type/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.hw_firmware_type.anyOf[0]"
	//
	// Any of "bios", "uefi".
	HwFirmwareType ImageHwFirmwareType `json:"hw_firmware_type,nullable"`
	// '#/components/schemas/ImageSerializer/properties/hw_machine_type/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.hw_machine_type.anyOf[0]"
	//
	// Any of "pc", "q35".
	HwMachineType ImageHwMachineType `json:"hw_machine_type,nullable"`
	// '#/components/schemas/ImageSerializer/properties/is_baremetal/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.is_baremetal.anyOf[0]"
	IsBaremetal bool `json:"is_baremetal,nullable"`
	// '#/components/schemas/ImageSerializer/properties/ssh_key/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.ssh_key.anyOf[0]"
	//
	// Any of "allow", "deny", "required".
	SSHKey ImageSSHKey `json:"ssh_key,nullable"`
	// '#/components/schemas/ImageSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.ImageSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID             resp.Field
		CreatedAt      resp.Field
		DiskFormat     resp.Field
		MinDisk        resp.Field
		MinRam         resp.Field
		Name           resp.Field
		OsDistro       resp.Field
		OsType         resp.Field
		OsVersion      resp.Field
		ProjectID      resp.Field
		Region         resp.Field
		RegionID       resp.Field
		Size           resp.Field
		Status         resp.Field
		Tags           resp.Field
		UpdatedAt      resp.Field
		Visibility     resp.Field
		Architecture   resp.Field
		CreatorTaskID  resp.Field
		Description    resp.Field
		DisplayOrder   resp.Field
		HwFirmwareType resp.Field
		HwMachineType  resp.Field
		IsBaremetal    resp.Field
		SSHKey         resp.Field
		TaskID         resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Image) RawJSON() string { return r.JSON.raw }
func (r *Image) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ImageSerializer/properties/os_type'
// "$.components.schemas.ImageSerializer.properties.os_type"
type ImageOsType string

const (
	ImageOsTypeLinux   ImageOsType = "linux"
	ImageOsTypeWindows ImageOsType = "windows"
)

// '#/components/schemas/ImageSerializer/properties/architecture'
// "$.components.schemas.ImageSerializer.properties.architecture"
type ImageArchitecture string

const (
	ImageArchitectureAarch64 ImageArchitecture = "aarch64"
	ImageArchitectureX86_64  ImageArchitecture = "x86_64"
)

// '#/components/schemas/ImageSerializer/properties/hw_firmware_type/anyOf/0'
// "$.components.schemas.ImageSerializer.properties.hw_firmware_type.anyOf[0]"
type ImageHwFirmwareType string

const (
	ImageHwFirmwareTypeBios ImageHwFirmwareType = "bios"
	ImageHwFirmwareTypeUefi ImageHwFirmwareType = "uefi"
)

// '#/components/schemas/ImageSerializer/properties/hw_machine_type/anyOf/0'
// "$.components.schemas.ImageSerializer.properties.hw_machine_type.anyOf[0]"
type ImageHwMachineType string

const (
	ImageHwMachineTypePc  ImageHwMachineType = "pc"
	ImageHwMachineTypeQ35 ImageHwMachineType = "q35"
)

// '#/components/schemas/ImageSerializer/properties/ssh_key/anyOf/0'
// "$.components.schemas.ImageSerializer.properties.ssh_key.anyOf[0]"
type ImageSSHKey string

const (
	ImageSSHKeyAllow    ImageSSHKey = "allow"
	ImageSSHKeyDeny     ImageSSHKey = "deny"
	ImageSSHKeyRequired ImageSSHKey = "required"
)

// '#/components/schemas/ImageCollectionSerializer'
// "$.components.schemas.ImageCollectionSerializer"
type ImageList struct {
	// '#/components/schemas/ImageCollectionSerializer/properties/count'
	// "$.components.schemas.ImageCollectionSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/ImageCollectionSerializer/properties/results'
	// "$.components.schemas.ImageCollectionSerializer.properties.results"
	Results []Image `json:"results,required"`
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
func (r ImageList) RawJSON() string { return r.JSON.raw }
func (r *ImageList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceSerializer'
// "$.components.schemas.InstanceSerializer"
type Instance struct {
	// '#/components/schemas/InstanceSerializer/properties/addresses'
	// "$.components.schemas.InstanceSerializer.properties.addresses"
	Addresses map[string][]InstanceAddressUnion `json:"addresses,required"`
	// '#/components/schemas/InstanceSerializer/properties/blackhole_ports'
	// "$.components.schemas.InstanceSerializer.properties.blackhole_ports"
	BlackholePorts []InstanceBlackholePort `json:"blackhole_ports,required"`
	// '#/components/schemas/InstanceSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.InstanceSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,required"`
	// '#/components/schemas/InstanceSerializer/properties/ddos_profile/anyOf/0'
	// "$.components.schemas.InstanceSerializer.properties.ddos_profile.anyOf[0]"
	DDOSProfile DDOSProfile `json:"ddos_profile,required"`
	// '#/components/schemas/InstanceSerializer/properties/fixed_ip_assignments/anyOf/0'
	// "$.components.schemas.InstanceSerializer.properties.fixed_ip_assignments.anyOf[0]"
	FixedIPAssignments []InstanceFixedIPAssignment `json:"fixed_ip_assignments,required"`
	// '#/components/schemas/InstanceSerializer/properties/flavor'
	// "$.components.schemas.InstanceSerializer.properties.flavor"
	Flavor InstanceFlavorUnion `json:"flavor,required"`
	// '#/components/schemas/InstanceSerializer/properties/instance_created'
	// "$.components.schemas.InstanceSerializer.properties.instance_created"
	InstanceCreated time.Time `json:"instance_created,required" format:"date-time"`
	// '#/components/schemas/InstanceSerializer/properties/instance_description/anyOf/0'
	// "$.components.schemas.InstanceSerializer.properties.instance_description.anyOf[0]"
	InstanceDescription string `json:"instance_description,required"`
	// '#/components/schemas/InstanceSerializer/properties/instance_id'
	// "$.components.schemas.InstanceSerializer.properties.instance_id"
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	// '#/components/schemas/InstanceSerializer/properties/instance_isolation/anyOf/0'
	// "$.components.schemas.InstanceSerializer.properties.instance_isolation.anyOf[0]"
	InstanceIsolation InstanceInstanceIsolation `json:"instance_isolation,required"`
	// '#/components/schemas/InstanceSerializer/properties/instance_name'
	// "$.components.schemas.InstanceSerializer.properties.instance_name"
	InstanceName string `json:"instance_name,required"`
	// '#/components/schemas/InstanceSerializer/properties/keypair_name/anyOf/0'
	// "$.components.schemas.InstanceSerializer.properties.keypair_name.anyOf[0]"
	KeypairName string `json:"keypair_name,required"`
	// '#/components/schemas/InstanceSerializer/properties/project_id'
	// "$.components.schemas.InstanceSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/InstanceSerializer/properties/region'
	// "$.components.schemas.InstanceSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/InstanceSerializer/properties/region_id'
	// "$.components.schemas.InstanceSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/InstanceSerializer/properties/security_groups'
	// "$.components.schemas.InstanceSerializer.properties.security_groups"
	SecurityGroups []InstanceSecurityGroup `json:"security_groups,required"`
	// '#/components/schemas/InstanceSerializer/properties/status'
	// "$.components.schemas.InstanceSerializer.properties.status"
	//
	// Any of "ACTIVE", "BUILD", "DELETED", "ERROR", "HARD_REBOOT", "MIGRATING",
	// "PASSWORD", "PAUSED", "REBOOT", "REBUILD", "RESCUE", "RESIZE", "REVERT_RESIZE",
	// "SHELVED", "SHELVED_OFFLOADED", "SHUTOFF", "SOFT_DELETED", "SUSPENDED",
	// "UNKNOWN", "VERIFY_RESIZE".
	Status InstanceStatus `json:"status,required"`
	// '#/components/schemas/InstanceSerializer/properties/tags'
	// "$.components.schemas.InstanceSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/InstanceSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.InstanceSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,required"`
	// '#/components/schemas/InstanceSerializer/properties/task_state/anyOf/0'
	// "$.components.schemas.InstanceSerializer.properties.task_state.anyOf[0]"
	TaskState string `json:"task_state,required"`
	// '#/components/schemas/InstanceSerializer/properties/vm_state'
	// "$.components.schemas.InstanceSerializer.properties.vm_state"
	//
	// Any of "active", "building", "deleted", "error", "paused", "rescued", "resized",
	// "shelved", "shelved_offloaded", "soft-deleted", "stopped", "suspended".
	VmState InstanceVmState `json:"vm_state,required"`
	// '#/components/schemas/InstanceSerializer/properties/volumes'
	// "$.components.schemas.InstanceSerializer.properties.volumes"
	Volumes []InstanceVolume `json:"volumes,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Addresses           resp.Field
		BlackholePorts      resp.Field
		CreatorTaskID       resp.Field
		DDOSProfile         resp.Field
		FixedIPAssignments  resp.Field
		Flavor              resp.Field
		InstanceCreated     resp.Field
		InstanceDescription resp.Field
		InstanceID          resp.Field
		InstanceIsolation   resp.Field
		InstanceName        resp.Field
		KeypairName         resp.Field
		ProjectID           resp.Field
		Region              resp.Field
		RegionID            resp.Field
		SecurityGroups      resp.Field
		Status              resp.Field
		Tags                resp.Field
		TaskID              resp.Field
		TaskState           resp.Field
		VmState             resp.Field
		Volumes             resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Instance) RawJSON() string { return r.JSON.raw }
func (r *Instance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InstanceAddressUnion contains all possible properties and values from
// [InstanceAddressFloatingIPAddress], [InstanceAddressFixedIPAddressShort],
// [InstanceAddressFixedIPAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InstanceAddressUnion struct {
	Addr          string `json:"addr"`
	Type          string `json:"type"`
	InterfaceName string `json:"interface_name"`
	// This field is from variant [InstanceAddressFixedIPAddress].
	SubnetID string `json:"subnet_id"`
	// This field is from variant [InstanceAddressFixedIPAddress].
	SubnetName string `json:"subnet_name"`
	JSON       struct {
		Addr          resp.Field
		Type          resp.Field
		InterfaceName resp.Field
		SubnetID      resp.Field
		SubnetName    resp.Field
		raw           string
	} `json:"-"`
}

func (u InstanceAddressUnion) AsFloatingIPAddress() (v InstanceAddressFloatingIPAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InstanceAddressUnion) AsFixedIPAddressShort() (v InstanceAddressFixedIPAddressShort) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InstanceAddressUnion) AsFixedIPAddress() (v InstanceAddressFixedIPAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InstanceAddressUnion) RawJSON() string { return u.JSON.raw }

func (r *InstanceAddressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceSerializer/properties/addresses/additionalProperties/items/anyOf/0'
// "$.components.schemas.InstanceSerializer.properties.addresses.additionalProperties.items.anyOf[0]"
type InstanceAddressFloatingIPAddress struct {
	// '#/components/schemas/InstanceFloatingAddressSerializer/properties/addr'
	// "$.components.schemas.InstanceFloatingAddressSerializer.properties.addr"
	Addr string `json:"addr,required"`
	// '#/components/schemas/InstanceFloatingAddressSerializer/properties/type'
	// "$.components.schemas.InstanceFloatingAddressSerializer.properties.type"
	Type constant.Floating `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Addr        resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceAddressFloatingIPAddress) RawJSON() string { return r.JSON.raw }
func (r *InstanceAddressFloatingIPAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceSerializer/properties/addresses/additionalProperties/items/anyOf/1'
// "$.components.schemas.InstanceSerializer.properties.addresses.additionalProperties.items.anyOf[1]"
type InstanceAddressFixedIPAddressShort struct {
	// '#/components/schemas/InstanceFixedAddressShortSerializer/properties/addr'
	// "$.components.schemas.InstanceFixedAddressShortSerializer.properties.addr"
	Addr string `json:"addr,required"`
	// '#/components/schemas/InstanceFixedAddressShortSerializer/properties/interface_name/anyOf/0'
	// "$.components.schemas.InstanceFixedAddressShortSerializer.properties.interface_name.anyOf[0]"
	InterfaceName string `json:"interface_name,required"`
	// '#/components/schemas/InstanceFixedAddressShortSerializer/properties/type'
	// "$.components.schemas.InstanceFixedAddressShortSerializer.properties.type"
	Type constant.Fixed `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Addr          resp.Field
		InterfaceName resp.Field
		Type          resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceAddressFixedIPAddressShort) RawJSON() string { return r.JSON.raw }
func (r *InstanceAddressFixedIPAddressShort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceSerializer/properties/addresses/additionalProperties/items/anyOf/2'
// "$.components.schemas.InstanceSerializer.properties.addresses.additionalProperties.items.anyOf[2]"
type InstanceAddressFixedIPAddress struct {
	// '#/components/schemas/InstanceFixedAddressSerializer/properties/addr'
	// "$.components.schemas.InstanceFixedAddressSerializer.properties.addr"
	Addr string `json:"addr,required"`
	// '#/components/schemas/InstanceFixedAddressSerializer/properties/interface_name/anyOf/0'
	// "$.components.schemas.InstanceFixedAddressSerializer.properties.interface_name.anyOf[0]"
	InterfaceName string `json:"interface_name,required"`
	// '#/components/schemas/InstanceFixedAddressSerializer/properties/subnet_id'
	// "$.components.schemas.InstanceFixedAddressSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required"`
	// '#/components/schemas/InstanceFixedAddressSerializer/properties/subnet_name'
	// "$.components.schemas.InstanceFixedAddressSerializer.properties.subnet_name"
	SubnetName string `json:"subnet_name,required"`
	// '#/components/schemas/InstanceFixedAddressSerializer/properties/type'
	// "$.components.schemas.InstanceFixedAddressSerializer.properties.type"
	Type constant.Fixed `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Addr          resp.Field
		InterfaceName resp.Field
		SubnetID      resp.Field
		SubnetName    resp.Field
		Type          resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceAddressFixedIPAddress) RawJSON() string { return r.JSON.raw }
func (r *InstanceAddressFixedIPAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceSerializer/properties/blackhole_ports/items'
// "$.components.schemas.InstanceSerializer.properties.blackhole_ports.items"
type InstanceBlackholePort struct {
	// '#/components/schemas/BlackholePortSerializer/properties/AlarmEnd'
	// "$.components.schemas.BlackholePortSerializer.properties.AlarmEnd"
	AlarmEnd time.Time `json:"AlarmEnd,required" format:"date-time"`
	// '#/components/schemas/BlackholePortSerializer/properties/AlarmStart'
	// "$.components.schemas.BlackholePortSerializer.properties.AlarmStart"
	AlarmStart time.Time `json:"AlarmStart,required" format:"date-time"`
	// '#/components/schemas/BlackholePortSerializer/properties/AlarmState'
	// "$.components.schemas.BlackholePortSerializer.properties.AlarmState"
	//
	// Any of "ACK_REQ", "ALARM", "ARCHIVED", "CLEAR", "CLEARING", "CLEARING_FAIL",
	// "END_GRACE", "END_WAIT", "MANUAL_CLEAR", "MANUAL_CLEARING",
	// "MANUAL_CLEARING_FAIL", "MANUAL_MITIGATING", "MANUAL_STARTING",
	// "MANUAL_STARTING_FAIL", "MITIGATING", "STARTING", "STARTING_FAIL", "START_WAIT",
	// "ack_req", "alarm", "archived", "clear", "clearing", "clearing_fail",
	// "end_grace", "end_wait", "manual_clear", "manual_clearing",
	// "manual_clearing_fail", "manual_mitigating", "manual_starting",
	// "manual_starting_fail", "mitigating", "start_wait", "starting", "starting_fail".
	AlarmState string `json:"AlarmState,required"`
	// '#/components/schemas/BlackholePortSerializer/properties/AlertDuration'
	// "$.components.schemas.BlackholePortSerializer.properties.AlertDuration"
	AlertDuration string `json:"AlertDuration,required"`
	// '#/components/schemas/BlackholePortSerializer/properties/DestinationIP'
	// "$.components.schemas.BlackholePortSerializer.properties.DestinationIP"
	DestinationIP string `json:"DestinationIP,required"`
	// '#/components/schemas/BlackholePortSerializer/properties/ID'
	// "$.components.schemas.BlackholePortSerializer.properties.ID"
	ID int64 `json:"ID,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AlarmEnd      resp.Field
		AlarmStart    resp.Field
		AlarmState    resp.Field
		AlertDuration resp.Field
		DestinationIP resp.Field
		ID            resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceBlackholePort) RawJSON() string { return r.JSON.raw }
func (r *InstanceBlackholePort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceSerializer/properties/fixed_ip_assignments/anyOf/0/items'
// "$.components.schemas.InstanceSerializer.properties.fixed_ip_assignments.anyOf[0].items"
type InstanceFixedIPAssignment struct {
	// '#/components/schemas/IpAssignmentsSerializer/properties/external'
	// "$.components.schemas.IpAssignmentsSerializer.properties.external"
	External bool `json:"external,required"`
	// '#/components/schemas/IpAssignmentsSerializer/properties/ip_address'
	// "$.components.schemas.IpAssignmentsSerializer.properties.ip_address"
	IPAddress string `json:"ip_address,required"`
	// '#/components/schemas/IpAssignmentsSerializer/properties/subnet_id'
	// "$.components.schemas.IpAssignmentsSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		External    resp.Field
		IPAddress   resp.Field
		SubnetID    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFixedIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *InstanceFixedIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InstanceFlavorUnion contains all possible properties and values from
// [InstanceFlavorInstanceFlavor], [InstanceFlavorBareMetalFlavor],
// [InstanceFlavorGPUClusterFlavor].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InstanceFlavorUnion struct {
	Architecture string `json:"architecture"`
	FlavorID     string `json:"flavor_id"`
	FlavorName   string `json:"flavor_name"`
	// This field is a union of [InstanceFlavorInstanceFlavorHardwareDescription],
	// [InstanceFlavorBareMetalFlavorHardwareDescription],
	// [InstanceFlavorGPUClusterFlavorHardwareDescription]
	HardwareDescription InstanceFlavorUnionHardwareDescription `json:"hardware_description"`
	OsType              string                                 `json:"os_type"`
	Ram                 int64                                  `json:"ram"`
	Vcpus               int64                                  `json:"vcpus"`
	ResourceClass       string                                 `json:"resource_class"`
	JSON                struct {
		Architecture        resp.Field
		FlavorID            resp.Field
		FlavorName          resp.Field
		HardwareDescription resp.Field
		OsType              resp.Field
		Ram                 resp.Field
		Vcpus               resp.Field
		ResourceClass       resp.Field
		raw                 string
	} `json:"-"`
}

func (u InstanceFlavorUnion) AsInstanceFlavor() (v InstanceFlavorInstanceFlavor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InstanceFlavorUnion) AsBareMetalFlavor() (v InstanceFlavorBareMetalFlavor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InstanceFlavorUnion) AsGPUClusterFlavor() (v InstanceFlavorGPUClusterFlavor) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InstanceFlavorUnion) RawJSON() string { return u.JSON.raw }

func (r *InstanceFlavorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InstanceFlavorUnionHardwareDescription is an implicit subunion of
// [InstanceFlavorUnion]. InstanceFlavorUnionHardwareDescription provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [InstanceFlavorUnion].
type InstanceFlavorUnionHardwareDescription struct {
	Ram string `json:"ram"`
	// This field is from variant [InstanceFlavorInstanceFlavorHardwareDescription].
	Vcpus   string `json:"vcpus"`
	CPU     string `json:"cpu"`
	Disk    string `json:"disk"`
	License string `json:"license"`
	Network string `json:"network"`
	// This field is from variant [InstanceFlavorGPUClusterFlavorHardwareDescription].
	GPU  string `json:"gpu"`
	JSON struct {
		Ram     resp.Field
		Vcpus   resp.Field
		CPU     resp.Field
		Disk    resp.Field
		License resp.Field
		Network resp.Field
		GPU     resp.Field
		raw     string
	} `json:"-"`
}

func (r *InstanceFlavorUnionHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceSerializer/properties/flavor/anyOf/0'
// "$.components.schemas.InstanceSerializer.properties.flavor.anyOf[0]"
type InstanceFlavorInstanceFlavor struct {
	// '#/components/schemas/InstanceFlavorSerializer/properties/architecture'
	// "$.components.schemas.InstanceFlavorSerializer.properties.architecture"
	Architecture string `json:"architecture,required"`
	// '#/components/schemas/InstanceFlavorSerializer/properties/flavor_id'
	// "$.components.schemas.InstanceFlavorSerializer.properties.flavor_id"
	FlavorID string `json:"flavor_id,required"`
	// '#/components/schemas/InstanceFlavorSerializer/properties/flavor_name'
	// "$.components.schemas.InstanceFlavorSerializer.properties.flavor_name"
	FlavorName string `json:"flavor_name,required"`
	// '#/components/schemas/InstanceFlavorSerializer/properties/hardware_description'
	// "$.components.schemas.InstanceFlavorSerializer.properties.hardware_description"
	HardwareDescription InstanceFlavorInstanceFlavorHardwareDescription `json:"hardware_description,required"`
	// '#/components/schemas/InstanceFlavorSerializer/properties/os_type'
	// "$.components.schemas.InstanceFlavorSerializer.properties.os_type"
	OsType string `json:"os_type,required"`
	// '#/components/schemas/InstanceFlavorSerializer/properties/ram'
	// "$.components.schemas.InstanceFlavorSerializer.properties.ram"
	Ram int64 `json:"ram,required"`
	// '#/components/schemas/InstanceFlavorSerializer/properties/vcpus'
	// "$.components.schemas.InstanceFlavorSerializer.properties.vcpus"
	Vcpus int64 `json:"vcpus,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Architecture        resp.Field
		FlavorID            resp.Field
		FlavorName          resp.Field
		HardwareDescription resp.Field
		OsType              resp.Field
		Ram                 resp.Field
		Vcpus               resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorInstanceFlavor) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorInstanceFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceFlavorSerializer/properties/hardware_description'
// "$.components.schemas.InstanceFlavorSerializer.properties.hardware_description"
type InstanceFlavorInstanceFlavorHardwareDescription struct {
	// '#/components/schemas/InstanceFlavorHardwareDescriptionSerializer/properties/ram'
	// "$.components.schemas.InstanceFlavorHardwareDescriptionSerializer.properties.ram"
	Ram string `json:"ram,required"`
	// '#/components/schemas/InstanceFlavorHardwareDescriptionSerializer/properties/vcpus'
	// "$.components.schemas.InstanceFlavorHardwareDescriptionSerializer.properties.vcpus"
	Vcpus string `json:"vcpus,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Ram         resp.Field
		Vcpus       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorInstanceFlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorInstanceFlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceSerializer/properties/flavor/anyOf/1'
// "$.components.schemas.InstanceSerializer.properties.flavor.anyOf[1]"
type InstanceFlavorBareMetalFlavor struct {
	// '#/components/schemas/BareMetalFlavorSerializer/properties/architecture'
	// "$.components.schemas.BareMetalFlavorSerializer.properties.architecture"
	Architecture string `json:"architecture,required"`
	// '#/components/schemas/BareMetalFlavorSerializer/properties/flavor_id'
	// "$.components.schemas.BareMetalFlavorSerializer.properties.flavor_id"
	FlavorID string `json:"flavor_id,required"`
	// '#/components/schemas/BareMetalFlavorSerializer/properties/flavor_name'
	// "$.components.schemas.BareMetalFlavorSerializer.properties.flavor_name"
	FlavorName string `json:"flavor_name,required"`
	// '#/components/schemas/BareMetalFlavorSerializer/properties/hardware_description'
	// "$.components.schemas.BareMetalFlavorSerializer.properties.hardware_description"
	HardwareDescription InstanceFlavorBareMetalFlavorHardwareDescription `json:"hardware_description,required"`
	// '#/components/schemas/BareMetalFlavorSerializer/properties/os_type'
	// "$.components.schemas.BareMetalFlavorSerializer.properties.os_type"
	OsType string `json:"os_type,required"`
	// '#/components/schemas/BareMetalFlavorSerializer/properties/ram'
	// "$.components.schemas.BareMetalFlavorSerializer.properties.ram"
	Ram int64 `json:"ram,required"`
	// '#/components/schemas/BareMetalFlavorSerializer/properties/resource_class'
	// "$.components.schemas.BareMetalFlavorSerializer.properties.resource_class"
	ResourceClass string `json:"resource_class,required"`
	// '#/components/schemas/BareMetalFlavorSerializer/properties/vcpus'
	// "$.components.schemas.BareMetalFlavorSerializer.properties.vcpus"
	Vcpus int64 `json:"vcpus,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Architecture        resp.Field
		FlavorID            resp.Field
		FlavorName          resp.Field
		HardwareDescription resp.Field
		OsType              resp.Field
		Ram                 resp.Field
		ResourceClass       resp.Field
		Vcpus               resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorBareMetalFlavor) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorBareMetalFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/BareMetalFlavorSerializer/properties/hardware_description'
// "$.components.schemas.BareMetalFlavorSerializer.properties.hardware_description"
type InstanceFlavorBareMetalFlavorHardwareDescription struct {
	// '#/components/schemas/BareMetalFlavorHardwareDescriptionSerializer/properties/cpu'
	// "$.components.schemas.BareMetalFlavorHardwareDescriptionSerializer.properties.cpu"
	CPU string `json:"cpu,required"`
	// '#/components/schemas/BareMetalFlavorHardwareDescriptionSerializer/properties/disk'
	// "$.components.schemas.BareMetalFlavorHardwareDescriptionSerializer.properties.disk"
	Disk string `json:"disk,required"`
	// '#/components/schemas/BareMetalFlavorHardwareDescriptionSerializer/properties/license'
	// "$.components.schemas.BareMetalFlavorHardwareDescriptionSerializer.properties.license"
	License string `json:"license,required"`
	// '#/components/schemas/BareMetalFlavorHardwareDescriptionSerializer/properties/network'
	// "$.components.schemas.BareMetalFlavorHardwareDescriptionSerializer.properties.network"
	Network string `json:"network,required"`
	// '#/components/schemas/BareMetalFlavorHardwareDescriptionSerializer/properties/ram'
	// "$.components.schemas.BareMetalFlavorHardwareDescriptionSerializer.properties.ram"
	Ram string `json:"ram,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CPU         resp.Field
		Disk        resp.Field
		License     resp.Field
		Network     resp.Field
		Ram         resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorBareMetalFlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorBareMetalFlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceSerializer/properties/flavor/anyOf/2'
// "$.components.schemas.InstanceSerializer.properties.flavor.anyOf[2]"
type InstanceFlavorGPUClusterFlavor struct {
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/architecture'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.architecture"
	Architecture string `json:"architecture,required"`
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/flavor_id'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.flavor_id"
	FlavorID string `json:"flavor_id,required"`
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/flavor_name'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.flavor_name"
	FlavorName string `json:"flavor_name,required"`
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/hardware_description'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.hardware_description"
	HardwareDescription InstanceFlavorGPUClusterFlavorHardwareDescription `json:"hardware_description,required"`
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/os_type'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.os_type"
	OsType string `json:"os_type,required"`
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/ram'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.ram"
	Ram int64 `json:"ram,required"`
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/resource_class'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.resource_class"
	ResourceClass string `json:"resource_class,required"`
	// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/vcpus'
	// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.vcpus"
	Vcpus int64 `json:"vcpus,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Architecture        resp.Field
		FlavorID            resp.Field
		FlavorName          resp.Field
		HardwareDescription resp.Field
		OsType              resp.Field
		Ram                 resp.Field
		ResourceClass       resp.Field
		Vcpus               resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorGPUClusterFlavor) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorGPUClusterFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/DeprecatedGpuClusterFlavorSerializer/properties/hardware_description'
// "$.components.schemas.DeprecatedGpuClusterFlavorSerializer.properties.hardware_description"
type InstanceFlavorGPUClusterFlavorHardwareDescription struct {
	// '#/components/schemas/DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer/properties/cpu'
	// "$.components.schemas.DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer.properties.cpu"
	CPU string `json:"cpu,required"`
	// '#/components/schemas/DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer/properties/disk'
	// "$.components.schemas.DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer.properties.disk"
	Disk string `json:"disk,required"`
	// '#/components/schemas/DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer/properties/gpu'
	// "$.components.schemas.DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer.properties.gpu"
	GPU string `json:"gpu,required"`
	// '#/components/schemas/DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer/properties/license'
	// "$.components.schemas.DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer.properties.license"
	License string `json:"license,required"`
	// '#/components/schemas/DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer/properties/network'
	// "$.components.schemas.DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer.properties.network"
	Network string `json:"network,required"`
	// '#/components/schemas/DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer/properties/ram'
	// "$.components.schemas.DeprecatedAIClusterServerFlavorHardwareDescriptionSerializer.properties.ram"
	Ram string `json:"ram,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CPU         resp.Field
		Disk        resp.Field
		GPU         resp.Field
		License     resp.Field
		Network     resp.Field
		Ram         resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorGPUClusterFlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorGPUClusterFlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceSerializer/properties/instance_isolation/anyOf/0'
// "$.components.schemas.InstanceSerializer.properties.instance_isolation.anyOf[0]"
type InstanceInstanceIsolation struct {
	// '#/components/schemas/IsolationSerializer/properties/reason/anyOf/0'
	// "$.components.schemas.IsolationSerializer.properties.reason.anyOf[0]"
	Reason string `json:"reason,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Reason      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceInstanceIsolation) RawJSON() string { return r.JSON.raw }
func (r *InstanceInstanceIsolation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceSerializer/properties/security_groups/items'
// "$.components.schemas.InstanceSerializer.properties.security_groups.items"
type InstanceSecurityGroup struct {
	// '#/components/schemas/NameSerializerPydantic/properties/name'
	// "$.components.schemas.NameSerializerPydantic.properties.name"
	Name string `json:"name,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceSecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *InstanceSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceSerializer/properties/status'
// "$.components.schemas.InstanceSerializer.properties.status"
type InstanceStatus string

const (
	InstanceStatusActive           InstanceStatus = "ACTIVE"
	InstanceStatusBuild            InstanceStatus = "BUILD"
	InstanceStatusDeleted          InstanceStatus = "DELETED"
	InstanceStatusError            InstanceStatus = "ERROR"
	InstanceStatusHardReboot       InstanceStatus = "HARD_REBOOT"
	InstanceStatusMigrating        InstanceStatus = "MIGRATING"
	InstanceStatusPassword         InstanceStatus = "PASSWORD"
	InstanceStatusPaused           InstanceStatus = "PAUSED"
	InstanceStatusReboot           InstanceStatus = "REBOOT"
	InstanceStatusRebuild          InstanceStatus = "REBUILD"
	InstanceStatusRescue           InstanceStatus = "RESCUE"
	InstanceStatusResize           InstanceStatus = "RESIZE"
	InstanceStatusRevertResize     InstanceStatus = "REVERT_RESIZE"
	InstanceStatusShelved          InstanceStatus = "SHELVED"
	InstanceStatusShelvedOffloaded InstanceStatus = "SHELVED_OFFLOADED"
	InstanceStatusShutoff          InstanceStatus = "SHUTOFF"
	InstanceStatusSoftDeleted      InstanceStatus = "SOFT_DELETED"
	InstanceStatusSuspended        InstanceStatus = "SUSPENDED"
	InstanceStatusUnknown          InstanceStatus = "UNKNOWN"
	InstanceStatusVerifyResize     InstanceStatus = "VERIFY_RESIZE"
)

// '#/components/schemas/InstanceSerializer/properties/vm_state'
// "$.components.schemas.InstanceSerializer.properties.vm_state"
type InstanceVmState string

const (
	InstanceVmStateActive           InstanceVmState = "active"
	InstanceVmStateBuilding         InstanceVmState = "building"
	InstanceVmStateDeleted          InstanceVmState = "deleted"
	InstanceVmStateError            InstanceVmState = "error"
	InstanceVmStatePaused           InstanceVmState = "paused"
	InstanceVmStateRescued          InstanceVmState = "rescued"
	InstanceVmStateResized          InstanceVmState = "resized"
	InstanceVmStateShelved          InstanceVmState = "shelved"
	InstanceVmStateShelvedOffloaded InstanceVmState = "shelved_offloaded"
	InstanceVmStateSoftDeleted      InstanceVmState = "soft-deleted"
	InstanceVmStateStopped          InstanceVmState = "stopped"
	InstanceVmStateSuspended        InstanceVmState = "suspended"
)

// '#/components/schemas/InstanceSerializer/properties/volumes/items'
// "$.components.schemas.InstanceSerializer.properties.volumes.items"
type InstanceVolume struct {
	// '#/components/schemas/InstanceVolumeSerializer/properties/id'
	// "$.components.schemas.InstanceVolumeSerializer.properties.id"
	ID string `json:"id,required"`
	// '#/components/schemas/InstanceVolumeSerializer/properties/delete_on_termination'
	// "$.components.schemas.InstanceVolumeSerializer.properties.delete_on_termination"
	DeleteOnTermination bool `json:"delete_on_termination,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                  resp.Field
		DeleteOnTermination resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceVolume) RawJSON() string { return r.JSON.raw }
func (r *InstanceVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceCollectionSerializer'
// "$.components.schemas.InstanceCollectionSerializer"
type InstanceList struct {
	// '#/components/schemas/InstanceCollectionSerializer/properties/count'
	// "$.components.schemas.InstanceCollectionSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/InstanceCollectionSerializer/properties/results'
	// "$.components.schemas.InstanceCollectionSerializer.properties.results"
	Results []Instance `json:"results,required"`
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
func (r InstanceList) RawJSON() string { return r.JSON.raw }
func (r *InstanceList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceMetricsTimeUnitEnum'
// "$.components.schemas.InstanceMetricsTimeUnitEnum"
type InstanceMetricsTimeUnit string

const (
	InstanceMetricsTimeUnitDay  InstanceMetricsTimeUnit = "day"
	InstanceMetricsTimeUnitHour InstanceMetricsTimeUnit = "hour"
)

// '#/components/schemas/InterfaceIPFamily'
// "$.components.schemas.InterfaceIPFamily"
type InterfaceIPFamily string

const (
	InterfaceIPFamilyDual InterfaceIPFamily = "dual"
	InterfaceIPFamilyIpv4 InterfaceIPFamily = "ipv4"
	InterfaceIPFamilyIpv6 InterfaceIPFamily = "ipv6"
)

// '#/components/schemas/IPVersion' "$.components.schemas.IPVersion"
type IPVersion int64

const (
	IPVersion4 IPVersion = 4
	IPVersion6 IPVersion = 6
)

// '#/components/schemas/LoadbalancerSerializer'
// "$.components.schemas.LoadbalancerSerializer"
type LoadBalancer struct {
	// '#/components/schemas/LoadbalancerSerializer/properties/id'
	// "$.components.schemas.LoadbalancerSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/LoadbalancerSerializer/properties/created_at'
	// "$.components.schemas.LoadbalancerSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/LoadbalancerSerializer/properties/name'
	// "$.components.schemas.LoadbalancerSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/LoadbalancerSerializer/properties/operating_status'
	// "$.components.schemas.LoadbalancerSerializer.properties.operating_status"
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// '#/components/schemas/LoadbalancerSerializer/properties/project_id'
	// "$.components.schemas.LoadbalancerSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/LoadbalancerSerializer/properties/provisioning_status'
	// "$.components.schemas.LoadbalancerSerializer.properties.provisioning_status"
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// '#/components/schemas/LoadbalancerSerializer/properties/region'
	// "$.components.schemas.LoadbalancerSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/LoadbalancerSerializer/properties/region_id'
	// "$.components.schemas.LoadbalancerSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/LoadbalancerSerializer/properties/tags_v2'
	// "$.components.schemas.LoadbalancerSerializer.properties.tags_v2"
	TagsV2 []Tag `json:"tags_v2,required"`
	// '#/components/schemas/LoadbalancerSerializer/properties/additional_vips'
	// "$.components.schemas.LoadbalancerSerializer.properties.additional_vips"
	AdditionalVips []LoadBalancerAdditionalVip `json:"additional_vips"`
	// '#/components/schemas/LoadbalancerSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.LoadbalancerSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// '#/components/schemas/LoadbalancerSerializer/properties/ddos_profile/anyOf/0'
	// "$.components.schemas.LoadbalancerSerializer.properties.ddos_profile.anyOf[0]"
	DDOSProfile DDOSProfile `json:"ddos_profile,nullable"`
	// '#/components/schemas/LoadbalancerSerializer/properties/flavor/anyOf/0'
	// "$.components.schemas.LoadbalancerSerializer.properties.flavor.anyOf[0]"
	Flavor LoadBalancerFlavor `json:"flavor,nullable"`
	// '#/components/schemas/LoadbalancerSerializer/properties/floating_ips'
	// "$.components.schemas.LoadbalancerSerializer.properties.floating_ips"
	FloatingIPs []FloatingIP `json:"floating_ips"`
	// '#/components/schemas/LoadbalancerSerializer/properties/listeners'
	// "$.components.schemas.LoadbalancerSerializer.properties.listeners"
	Listeners []LoadBalancerListener `json:"listeners"`
	// '#/components/schemas/LoadbalancerSerializer/properties/logging/anyOf/0'
	// "$.components.schemas.LoadbalancerSerializer.properties.logging.anyOf[0]"
	Logging Logging `json:"logging,nullable"`
	// '#/components/schemas/LoadbalancerSerializer/properties/preferred_connectivity'
	// "$.components.schemas.LoadbalancerSerializer.properties.preferred_connectivity"
	//
	// Any of "L2", "L3".
	PreferredConnectivity LoadBalancerMemberConnectivity `json:"preferred_connectivity"`
	// '#/components/schemas/LoadbalancerSerializer/properties/stats/anyOf/0'
	// "$.components.schemas.LoadbalancerSerializer.properties.stats.anyOf[0]"
	Stats LoadBalancerStatistics `json:"stats,nullable"`
	// '#/components/schemas/LoadbalancerSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.LoadbalancerSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// '#/components/schemas/LoadbalancerSerializer/properties/updated_at/anyOf/0'
	// "$.components.schemas.LoadbalancerSerializer.properties.updated_at.anyOf[0]"
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// '#/components/schemas/LoadbalancerSerializer/properties/vip_address/anyOf/0'
	// "$.components.schemas.LoadbalancerSerializer.properties.vip_address.anyOf[0]"
	VipAddress string `json:"vip_address,nullable" format:"ipvanyaddress"`
	// '#/components/schemas/LoadbalancerSerializer/properties/vip_ip_family/anyOf/0'
	// "$.components.schemas.LoadbalancerSerializer.properties.vip_ip_family.anyOf[0]"
	//
	// Any of "dual", "ipv4", "ipv6".
	VipIPFamily InterfaceIPFamily `json:"vip_ip_family,nullable"`
	// '#/components/schemas/LoadbalancerSerializer/properties/vip_port_id/anyOf/0'
	// "$.components.schemas.LoadbalancerSerializer.properties.vip_port_id.anyOf[0]"
	VipPortID string `json:"vip_port_id,nullable" format:"uuid4"`
	// '#/components/schemas/LoadbalancerSerializer/properties/vrrp_ips'
	// "$.components.schemas.LoadbalancerSerializer.properties.vrrp_ips"
	VrrpIPs []LoadBalancerVrrpIP `json:"vrrp_ips"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                    resp.Field
		CreatedAt             resp.Field
		Name                  resp.Field
		OperatingStatus       resp.Field
		ProjectID             resp.Field
		ProvisioningStatus    resp.Field
		Region                resp.Field
		RegionID              resp.Field
		TagsV2                resp.Field
		AdditionalVips        resp.Field
		CreatorTaskID         resp.Field
		DDOSProfile           resp.Field
		Flavor                resp.Field
		FloatingIPs           resp.Field
		Listeners             resp.Field
		Logging               resp.Field
		PreferredConnectivity resp.Field
		Stats                 resp.Field
		TaskID                resp.Field
		UpdatedAt             resp.Field
		VipAddress            resp.Field
		VipIPFamily           resp.Field
		VipPortID             resp.Field
		VrrpIPs               resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancer) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LoadbalancerSerializer/properties/additional_vips/items'
// "$.components.schemas.LoadbalancerSerializer.properties.additional_vips.items"
type LoadBalancerAdditionalVip struct {
	// '#/components/schemas/NetworkPortFixedIp/properties/ip_address'
	// "$.components.schemas.NetworkPortFixedIp.properties.ip_address"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/NetworkPortFixedIp/properties/subnet_id'
	// "$.components.schemas.NetworkPortFixedIp.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IPAddress   resp.Field
		SubnetID    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerAdditionalVip) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerAdditionalVip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LoadbalancerSerializer/properties/flavor/anyOf/0'
// "$.components.schemas.LoadbalancerSerializer.properties.flavor.anyOf[0]"
type LoadBalancerFlavor struct {
	// '#/components/schemas/LbFlavorSerializer/properties/flavor_id'
	// "$.components.schemas.LbFlavorSerializer.properties.flavor_id"
	FlavorID string `json:"flavor_id,required"`
	// '#/components/schemas/LbFlavorSerializer/properties/flavor_name'
	// "$.components.schemas.LbFlavorSerializer.properties.flavor_name"
	FlavorName string `json:"flavor_name,required"`
	// '#/components/schemas/LbFlavorSerializer/properties/ram'
	// "$.components.schemas.LbFlavorSerializer.properties.ram"
	Ram int64 `json:"ram,required"`
	// '#/components/schemas/LbFlavorSerializer/properties/vcpus'
	// "$.components.schemas.LbFlavorSerializer.properties.vcpus"
	Vcpus int64 `json:"vcpus,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		FlavorID    resp.Field
		FlavorName  resp.Field
		Ram         resp.Field
		Vcpus       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerFlavor) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LoadbalancerSerializer/properties/listeners/items'
// "$.components.schemas.LoadbalancerSerializer.properties.listeners.items"
type LoadBalancerListener struct {
	// '#/components/schemas/ListenerSerializer/properties/id'
	// "$.components.schemas.ListenerSerializer.properties.id"
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
func (r LoadBalancerListener) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerListener) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LoadbalancerSerializer/properties/vrrp_ips/items'
// "$.components.schemas.LoadbalancerSerializer.properties.vrrp_ips.items"
type LoadBalancerVrrpIP struct {
	// '#/components/schemas/VRRPIP/properties/ip_address'
	// "$.components.schemas.VRRPIP.properties.ip_address"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/VRRPIP/properties/role'
	// "$.components.schemas.VRRPIP.properties.role"
	//
	// Any of "BACKUP", "MASTER", "STANDALONE".
	Role LoadBalancerInstanceRole `json:"role,required"`
	// '#/components/schemas/VRRPIP/properties/subnet_id'
	// "$.components.schemas.VRRPIP.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IPAddress   resp.Field
		Role        resp.Field
		SubnetID    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerVrrpIP) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerVrrpIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/LoadBalancerInstanceRoleEnum'
// "$.components.schemas.LoadBalancerInstanceRoleEnum"
type LoadBalancerInstanceRole string

const (
	LoadBalancerInstanceRoleBackup     LoadBalancerInstanceRole = "BACKUP"
	LoadBalancerInstanceRoleMaster     LoadBalancerInstanceRole = "MASTER"
	LoadBalancerInstanceRoleStandalone LoadBalancerInstanceRole = "STANDALONE"
)

// '#/components/schemas/MemberConnectivity'
// "$.components.schemas.MemberConnectivity"
type LoadBalancerMemberConnectivity string

const (
	LoadBalancerMemberConnectivityL2 LoadBalancerMemberConnectivity = "L2"
	LoadBalancerMemberConnectivityL3 LoadBalancerMemberConnectivity = "L3"
)

// '#/components/schemas/OperatingStatusEnum'
// "$.components.schemas.OperatingStatusEnum"
type LoadBalancerOperatingStatus string

const (
	LoadBalancerOperatingStatusDegraded  LoadBalancerOperatingStatus = "DEGRADED"
	LoadBalancerOperatingStatusDraining  LoadBalancerOperatingStatus = "DRAINING"
	LoadBalancerOperatingStatusError     LoadBalancerOperatingStatus = "ERROR"
	LoadBalancerOperatingStatusNoMonitor LoadBalancerOperatingStatus = "NO_MONITOR"
	LoadBalancerOperatingStatusOffline   LoadBalancerOperatingStatus = "OFFLINE"
	LoadBalancerOperatingStatusOnline    LoadBalancerOperatingStatus = "ONLINE"
)

// '#/components/schemas/LoadbalancerStatsSerializer'
// "$.components.schemas.LoadbalancerStatsSerializer"
type LoadBalancerStatistics struct {
	// '#/components/schemas/LoadbalancerStatsSerializer/properties/active_connections'
	// "$.components.schemas.LoadbalancerStatsSerializer.properties.active_connections"
	ActiveConnections int64 `json:"active_connections,required"`
	// '#/components/schemas/LoadbalancerStatsSerializer/properties/bytes_in'
	// "$.components.schemas.LoadbalancerStatsSerializer.properties.bytes_in"
	BytesIn int64 `json:"bytes_in,required"`
	// '#/components/schemas/LoadbalancerStatsSerializer/properties/bytes_out'
	// "$.components.schemas.LoadbalancerStatsSerializer.properties.bytes_out"
	BytesOut int64 `json:"bytes_out,required"`
	// '#/components/schemas/LoadbalancerStatsSerializer/properties/request_errors'
	// "$.components.schemas.LoadbalancerStatsSerializer.properties.request_errors"
	RequestErrors int64 `json:"request_errors,required"`
	// '#/components/schemas/LoadbalancerStatsSerializer/properties/total_connections'
	// "$.components.schemas.LoadbalancerStatsSerializer.properties.total_connections"
	TotalConnections int64 `json:"total_connections,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ActiveConnections resp.Field
		BytesIn           resp.Field
		BytesOut          resp.Field
		RequestErrors     resp.Field
		TotalConnections  resp.Field
		ExtraFields       map[string]resp.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerStatistics) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

// '#/components/schemas/InstanceInterfaceTrunkSerializer'
// "$.components.schemas.InstanceInterfaceTrunkSerializer"
type NetworkInterface struct {
	// '#/components/schemas/InstanceInterfaceTrunkSerializer/properties/allowed_address_pairs'
	// "$.components.schemas.InstanceInterfaceTrunkSerializer.properties.allowed_address_pairs"
	AllowedAddressPairs []NetworkInterfaceAllowedAddressPair `json:"allowed_address_pairs,required"`
	// '#/components/schemas/InstanceInterfaceTrunkSerializer/properties/floatingip_details'
	// "$.components.schemas.InstanceInterfaceTrunkSerializer.properties.floatingip_details"
	FloatingipDetails []FloatingIP `json:"floatingip_details,required"`
	// '#/components/schemas/InstanceInterfaceTrunkSerializer/properties/ip_assignments'
	// "$.components.schemas.InstanceInterfaceTrunkSerializer.properties.ip_assignments"
	IPAssignments []NetworkInterfaceIPAssignment `json:"ip_assignments,required"`
	// '#/components/schemas/InstanceInterfaceTrunkSerializer/properties/network_details'
	// "$.components.schemas.InstanceInterfaceTrunkSerializer.properties.network_details"
	NetworkDetails NetworkInterfaceNetworkDetails `json:"network_details,required"`
	// '#/components/schemas/InstanceInterfaceTrunkSerializer/properties/network_id'
	// "$.components.schemas.InstanceInterfaceTrunkSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/InstanceInterfaceTrunkSerializer/properties/port_id'
	// "$.components.schemas.InstanceInterfaceTrunkSerializer.properties.port_id"
	PortID string `json:"port_id,required" format:"uuid4"`
	// '#/components/schemas/InstanceInterfaceTrunkSerializer/properties/port_security_enabled'
	// "$.components.schemas.InstanceInterfaceTrunkSerializer.properties.port_security_enabled"
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// '#/components/schemas/InstanceInterfaceTrunkSerializer/properties/sub_ports'
	// "$.components.schemas.InstanceInterfaceTrunkSerializer.properties.sub_ports"
	SubPorts []NetworkInterfaceSubPort `json:"sub_ports,required"`
	// '#/components/schemas/InstanceInterfaceTrunkSerializer/properties/interface_name/anyOf/0'
	// "$.components.schemas.InstanceInterfaceTrunkSerializer.properties.interface_name.anyOf[0]"
	InterfaceName string `json:"interface_name,nullable"`
	// '#/components/schemas/InstanceInterfaceTrunkSerializer/properties/mac_address/anyOf/0'
	// "$.components.schemas.InstanceInterfaceTrunkSerializer.properties.mac_address.anyOf[0]"
	MacAddress string `json:"mac_address,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AllowedAddressPairs resp.Field
		FloatingipDetails   resp.Field
		IPAssignments       resp.Field
		NetworkDetails      resp.Field
		NetworkID           resp.Field
		PortID              resp.Field
		PortSecurityEnabled resp.Field
		SubPorts            resp.Field
		InterfaceName       resp.Field
		MacAddress          resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkInterface) RawJSON() string { return r.JSON.raw }
func (r *NetworkInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInterfaceTrunkSerializer/properties/allowed_address_pairs/items'
// "$.components.schemas.InstanceInterfaceTrunkSerializer.properties.allowed_address_pairs.items"
type NetworkInterfaceAllowedAddressPair struct {
	// '#/components/schemas/AllowedAddressPairsSerializer/properties/ip_address/anyOf/0'
	// "$.components.schemas.AllowedAddressPairsSerializer.properties.ip_address.anyOf[0]"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/AllowedAddressPairsSerializer/properties/mac_address/anyOf/0'
	// "$.components.schemas.AllowedAddressPairsSerializer.properties.mac_address.anyOf[0]"
	MacAddress string `json:"mac_address,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IPAddress   resp.Field
		MacAddress  resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkInterfaceAllowedAddressPair) RawJSON() string { return r.JSON.raw }
func (r *NetworkInterfaceAllowedAddressPair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInterfaceTrunkSerializer/properties/ip_assignments/items'
// "$.components.schemas.InstanceInterfaceTrunkSerializer.properties.ip_assignments.items"
type NetworkInterfaceIPAssignment struct {
	// '#/components/schemas/PortIpSubnetIdSerializer/properties/ip_address'
	// "$.components.schemas.PortIpSubnetIdSerializer.properties.ip_address"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/PortIpSubnetIdSerializer/properties/subnet_id'
	// "$.components.schemas.PortIpSubnetIdSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IPAddress   resp.Field
		SubnetID    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkInterfaceIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *NetworkInterfaceIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInterfaceTrunkSerializer/properties/network_details'
// "$.components.schemas.InstanceInterfaceTrunkSerializer.properties.network_details"
type NetworkInterfaceNetworkDetails struct {
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/id'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/created_at'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/external'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.external"
	External bool `json:"external,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/name'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/port_security_enabled'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.port_security_enabled"
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/region'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/region_id'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/shared'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.shared"
	Shared bool `json:"shared,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/tags'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/type'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.type"
	Type string `json:"type,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/updated_at'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/default/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties['default'].anyOf[0]"
	Default bool `json:"default,nullable"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/mtu'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.mtu"
	Mtu int64 `json:"mtu"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/project_id/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.project_id.anyOf[0]"
	ProjectID int64 `json:"project_id,nullable"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/segmentation_id/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.segmentation_id.anyOf[0]"
	SegmentationID int64 `json:"segmentation_id,nullable"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/subnets/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.subnets.anyOf[0]"
	Subnets []Subnet `json:"subnets,nullable"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                  resp.Field
		CreatedAt           resp.Field
		External            resp.Field
		Name                resp.Field
		PortSecurityEnabled resp.Field
		Region              resp.Field
		RegionID            resp.Field
		Shared              resp.Field
		Tags                resp.Field
		Type                resp.Field
		UpdatedAt           resp.Field
		CreatorTaskID       resp.Field
		Default             resp.Field
		Mtu                 resp.Field
		ProjectID           resp.Field
		SegmentationID      resp.Field
		Subnets             resp.Field
		TaskID              resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkInterfaceNetworkDetails) RawJSON() string { return r.JSON.raw }
func (r *NetworkInterfaceNetworkDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInterfaceTrunkSerializer/properties/sub_ports/items'
// "$.components.schemas.InstanceInterfaceTrunkSerializer.properties.sub_ports.items"
type NetworkInterfaceSubPort struct {
	// '#/components/schemas/InstanceInterfaceSubportSerializer/properties/allowed_address_pairs'
	// "$.components.schemas.InstanceInterfaceSubportSerializer.properties.allowed_address_pairs"
	AllowedAddressPairs []NetworkInterfaceSubPortAllowedAddressPair `json:"allowed_address_pairs,required"`
	// '#/components/schemas/InstanceInterfaceSubportSerializer/properties/floatingip_details'
	// "$.components.schemas.InstanceInterfaceSubportSerializer.properties.floatingip_details"
	FloatingipDetails []FloatingIP `json:"floatingip_details,required"`
	// '#/components/schemas/InstanceInterfaceSubportSerializer/properties/ip_assignments'
	// "$.components.schemas.InstanceInterfaceSubportSerializer.properties.ip_assignments"
	IPAssignments []NetworkInterfaceSubPortIPAssignment `json:"ip_assignments,required"`
	// '#/components/schemas/InstanceInterfaceSubportSerializer/properties/network_details'
	// "$.components.schemas.InstanceInterfaceSubportSerializer.properties.network_details"
	NetworkDetails NetworkInterfaceSubPortNetworkDetails `json:"network_details,required"`
	// '#/components/schemas/InstanceInterfaceSubportSerializer/properties/network_id'
	// "$.components.schemas.InstanceInterfaceSubportSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/InstanceInterfaceSubportSerializer/properties/port_id'
	// "$.components.schemas.InstanceInterfaceSubportSerializer.properties.port_id"
	PortID string `json:"port_id,required" format:"uuid4"`
	// '#/components/schemas/InstanceInterfaceSubportSerializer/properties/port_security_enabled'
	// "$.components.schemas.InstanceInterfaceSubportSerializer.properties.port_security_enabled"
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// '#/components/schemas/InstanceInterfaceSubportSerializer/properties/segmentation_id'
	// "$.components.schemas.InstanceInterfaceSubportSerializer.properties.segmentation_id"
	SegmentationID int64 `json:"segmentation_id,required"`
	// '#/components/schemas/InstanceInterfaceSubportSerializer/properties/segmentation_type'
	// "$.components.schemas.InstanceInterfaceSubportSerializer.properties.segmentation_type"
	SegmentationType string `json:"segmentation_type,required"`
	// '#/components/schemas/InstanceInterfaceSubportSerializer/properties/interface_name/anyOf/0'
	// "$.components.schemas.InstanceInterfaceSubportSerializer.properties.interface_name.anyOf[0]"
	InterfaceName string `json:"interface_name,nullable"`
	// '#/components/schemas/InstanceInterfaceSubportSerializer/properties/mac_address/anyOf/0'
	// "$.components.schemas.InstanceInterfaceSubportSerializer.properties.mac_address.anyOf[0]"
	MacAddress string `json:"mac_address,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AllowedAddressPairs resp.Field
		FloatingipDetails   resp.Field
		IPAssignments       resp.Field
		NetworkDetails      resp.Field
		NetworkID           resp.Field
		PortID              resp.Field
		PortSecurityEnabled resp.Field
		SegmentationID      resp.Field
		SegmentationType    resp.Field
		InterfaceName       resp.Field
		MacAddress          resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkInterfaceSubPort) RawJSON() string { return r.JSON.raw }
func (r *NetworkInterfaceSubPort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInterfaceSubportSerializer/properties/allowed_address_pairs/items'
// "$.components.schemas.InstanceInterfaceSubportSerializer.properties.allowed_address_pairs.items"
type NetworkInterfaceSubPortAllowedAddressPair struct {
	// '#/components/schemas/AllowedAddressPairsSerializer/properties/ip_address/anyOf/0'
	// "$.components.schemas.AllowedAddressPairsSerializer.properties.ip_address.anyOf[0]"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/AllowedAddressPairsSerializer/properties/mac_address/anyOf/0'
	// "$.components.schemas.AllowedAddressPairsSerializer.properties.mac_address.anyOf[0]"
	MacAddress string `json:"mac_address,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IPAddress   resp.Field
		MacAddress  resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkInterfaceSubPortAllowedAddressPair) RawJSON() string { return r.JSON.raw }
func (r *NetworkInterfaceSubPortAllowedAddressPair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInterfaceSubportSerializer/properties/ip_assignments/items'
// "$.components.schemas.InstanceInterfaceSubportSerializer.properties.ip_assignments.items"
type NetworkInterfaceSubPortIPAssignment struct {
	// '#/components/schemas/PortIpSubnetIdSerializer/properties/ip_address'
	// "$.components.schemas.PortIpSubnetIdSerializer.properties.ip_address"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// '#/components/schemas/PortIpSubnetIdSerializer/properties/subnet_id'
	// "$.components.schemas.PortIpSubnetIdSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IPAddress   resp.Field
		SubnetID    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkInterfaceSubPortIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *NetworkInterfaceSubPortIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInterfaceSubportSerializer/properties/network_details'
// "$.components.schemas.InstanceInterfaceSubportSerializer.properties.network_details"
type NetworkInterfaceSubPortNetworkDetails struct {
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/id'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/created_at'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/external'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.external"
	External bool `json:"external,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/name'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/port_security_enabled'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.port_security_enabled"
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/region'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/region_id'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/shared'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.shared"
	Shared bool `json:"shared,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/tags'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/type'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.type"
	Type string `json:"type,required"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/updated_at'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/default/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties['default'].anyOf[0]"
	Default bool `json:"default,nullable"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/mtu'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.mtu"
	Mtu int64 `json:"mtu"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/project_id/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.project_id.anyOf[0]"
	ProjectID int64 `json:"project_id,nullable"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/segmentation_id/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.segmentation_id.anyOf[0]"
	SegmentationID int64 `json:"segmentation_id,nullable"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/subnets/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.subnets.anyOf[0]"
	Subnets []Subnet `json:"subnets,nullable"`
	// '#/components/schemas/NetworkSubnetworkSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.NetworkSubnetworkSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                  resp.Field
		CreatedAt           resp.Field
		External            resp.Field
		Name                resp.Field
		PortSecurityEnabled resp.Field
		Region              resp.Field
		RegionID            resp.Field
		Shared              resp.Field
		Tags                resp.Field
		Type                resp.Field
		UpdatedAt           resp.Field
		CreatorTaskID       resp.Field
		Default             resp.Field
		Mtu                 resp.Field
		ProjectID           resp.Field
		SegmentationID      resp.Field
		Subnets             resp.Field
		TaskID              resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkInterfaceSubPortNetworkDetails) RawJSON() string { return r.JSON.raw }
func (r *NetworkInterfaceSubPortNetworkDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InstanceInterfaceTrunkCollectionSerializer'
// "$.components.schemas.InstanceInterfaceTrunkCollectionSerializer"
type NetworkInterfaceList struct {
	// '#/components/schemas/InstanceInterfaceTrunkCollectionSerializer/properties/count'
	// "$.components.schemas.InstanceInterfaceTrunkCollectionSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/InstanceInterfaceTrunkCollectionSerializer/properties/results'
	// "$.components.schemas.InstanceInterfaceTrunkCollectionSerializer.properties.results"
	Results []NetworkInterface `json:"results,required"`
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
func (r NetworkInterfaceList) RawJSON() string { return r.JSON.raw }
func (r *NetworkInterfaceList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ProvisioningStatusEnum'
// "$.components.schemas.ProvisioningStatusEnum"
type ProvisioningStatus string

const (
	ProvisioningStatusActive        ProvisioningStatus = "ACTIVE"
	ProvisioningStatusDeleted       ProvisioningStatus = "DELETED"
	ProvisioningStatusError         ProvisioningStatus = "ERROR"
	ProvisioningStatusPendingCreate ProvisioningStatus = "PENDING_CREATE"
	ProvisioningStatusPendingDelete ProvisioningStatus = "PENDING_DELETE"
	ProvisioningStatusPendingUpdate ProvisioningStatus = "PENDING_UPDATE"
)

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
	// '#/components/schemas/RouteOutSerializer/properties/destination'
	// "$.components.schemas.RouteOutSerializer.properties.destination"
	Destination string `json:"destination,required" format:"ipvanynetwork"`
	// '#/components/schemas/RouteOutSerializer/properties/nexthop'
	// "$.components.schemas.RouteOutSerializer.properties.nexthop"
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

type TagUpdateList map[string]string

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
