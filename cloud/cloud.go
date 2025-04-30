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
	Options             []option.RequestOption
	Projects            ProjectService
	Tasks               TaskService
	Regions             RegionService
	Quotas              QuotaService
	Secrets             SecretService
	SSHKeys             SSHKeyService
	IPRanges            IPRangeService
	LoadBalancers       LoadBalancerService
	ReservedFixedIPs    ReservedFixedIPService
	Networks            NetworkService
	Volumes             VolumeService
	FloatingIPs         FloatingIPService
	SecurityGroups      SecurityGroupService
	Users               UserService
	Inference           InferenceService
	PlacementGroups     PlacementGroupService
	Baremetal           BaremetalService
	Instances           InstanceService
	Registries          RegistryService
	FileShares          FileShareService
	BillingReservations BillingReservationService
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
	r.Instances = NewInstanceService(opts...)
	r.Registries = NewRegistryService(opts...)
	r.FileShares = NewFileShareService(opts...)
	r.BillingReservations = NewBillingReservationService(opts...)
	return
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
