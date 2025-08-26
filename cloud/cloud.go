// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"encoding/json"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/G-Core/gcore-go/shared/constant"
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
	AuditLogs            AuditLogService
	CostReports          CostReportService
	UsageReports         UsageReportService
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
	r.AuditLogs = NewAuditLogService(opts...)
	r.CostReports = NewCostReportService(opts...)
	r.UsageReports = NewUsageReportService(opts...)
	return
}

type AllowedAddressPairs struct {
	// Subnet mask or IP address of the port specified in `allowed_address_pairs`
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// MAC address of the port specified in `allowed_address_pairs`
	MacAddress string `json:"mac_address,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress   respjson.Field
		MacAddress  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllowedAddressPairs) RawJSON() string { return r.JSON.raw }
func (r *AllowedAddressPairs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal flavor schema
type BaremetalFlavor struct {
	// Flavor architecture type
	Architecture string `json:"architecture,required"`
	// Disabled flavor flag
	Disabled bool `json:"disabled,required"`
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id,required"`
	// Flavor name
	FlavorName string `json:"flavor_name,required"`
	// Flavor operating system
	OsType string `json:"os_type,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Flavor resource class for mapping to hardware capacity
	ResourceClass string `json:"resource_class,required"`
	// Virtual CPU count. For bare metal flavors, it's a physical CPU count
	Vcpus int64 `json:"vcpus,required"`
	// Number of available instances of given configuration
	Capacity int64 `json:"capacity,nullable"`
	// Currency code. Shown if the `include_prices` query parameter if set to true
	CurrencyCode string `json:"currency_code,nullable"`
	// Additional hardware description
	HardwareDescription map[string]string `json:"hardware_description"`
	// Price per hour. Shown if the `include_prices` query parameter if set to true
	PricePerHour float64 `json:"price_per_hour,nullable"`
	// Price per month. Shown if the `include_prices` query parameter if set to true
	PricePerMonth float64 `json:"price_per_month,nullable"`
	// Price status for the UI
	//
	// Any of "error", "hide", "show".
	PriceStatus BaremetalFlavorPriceStatus `json:"price_status,nullable"`
	// Count of reserved but not used nodes. If a client don't have reservations for
	// the flavor, it's None.
	ReservedInStock int64 `json:"reserved_in_stock,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture        respjson.Field
		Disabled            respjson.Field
		FlavorID            respjson.Field
		FlavorName          respjson.Field
		OsType              respjson.Field
		Ram                 respjson.Field
		ResourceClass       respjson.Field
		Vcpus               respjson.Field
		Capacity            respjson.Field
		CurrencyCode        respjson.Field
		HardwareDescription respjson.Field
		PricePerHour        respjson.Field
		PricePerMonth       respjson.Field
		PriceStatus         respjson.Field
		ReservedInStock     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalFlavor) RawJSON() string { return r.JSON.raw }
func (r *BaremetalFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price status for the UI
type BaremetalFlavorPriceStatus string

const (
	BaremetalFlavorPriceStatusError BaremetalFlavorPriceStatus = "error"
	BaremetalFlavorPriceStatusHide  BaremetalFlavorPriceStatus = "hide"
	BaremetalFlavorPriceStatusShow  BaremetalFlavorPriceStatus = "show"
)

type BaremetalFlavorList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []BaremetalFlavor `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalFlavorList) RawJSON() string { return r.JSON.raw }
func (r *BaremetalFlavorList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlackholePort struct {
	// A date-time string giving the time that the alarm ended
	AlarmEnd time.Time `json:"AlarmEnd,required" format:"date-time"`
	// A date-time string giving the time that the alarm started
	AlarmStart time.Time `json:"AlarmStart,required" format:"date-time"`
	// Current state of alarm
	//
	// Any of "ACK_REQ", "ALARM", "ARCHIVED", "CLEAR", "CLEARING", "CLEARING_FAIL",
	// "END_GRACE", "END_WAIT", "MANUAL_CLEAR", "MANUAL_CLEARING",
	// "MANUAL_CLEARING_FAIL", "MANUAL_MITIGATING", "MANUAL_STARTING",
	// "MANUAL_STARTING_FAIL", "MITIGATING", "STARTING", "STARTING_FAIL", "START_WAIT",
	// "ack_req", "alarm", "archived", "clear", "clearing", "clearing_fail",
	// "end_grace", "end_wait", "manual_clear", "manual_clearing",
	// "manual_clearing_fail", "manual_mitigating", "manual_starting",
	// "manual_starting_fail", "mitigating", "start_wait", "starting", "starting_fail".
	AlarmState BlackholePortAlarmState `json:"AlarmState,required"`
	// Total alert duration
	AlertDuration string `json:"AlertDuration,required"`
	// Notification destination IP address
	DestinationIP string `json:"DestinationIP,required"`
	ID            int64  `json:"ID,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AlarmEnd      respjson.Field
		AlarmStart    respjson.Field
		AlarmState    respjson.Field
		AlertDuration respjson.Field
		DestinationIP respjson.Field
		ID            respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BlackholePort) RawJSON() string { return r.JSON.raw }
func (r *BlackholePort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current state of alarm
type BlackholePortAlarmState string

const (
	BlackholePortAlarmStateAckReqUppercase             BlackholePortAlarmState = "ACK_REQ"
	BlackholePortAlarmStateAlarmUppercase              BlackholePortAlarmState = "ALARM"
	BlackholePortAlarmStateArchivedUppercase           BlackholePortAlarmState = "ARCHIVED"
	BlackholePortAlarmStateClearUppercase              BlackholePortAlarmState = "CLEAR"
	BlackholePortAlarmStateClearingUppercase           BlackholePortAlarmState = "CLEARING"
	BlackholePortAlarmStateClearingFailUppercase       BlackholePortAlarmState = "CLEARING_FAIL"
	BlackholePortAlarmStateEndGraceUppercase           BlackholePortAlarmState = "END_GRACE"
	BlackholePortAlarmStateEndWaitUppercase            BlackholePortAlarmState = "END_WAIT"
	BlackholePortAlarmStateManualClearUppercase        BlackholePortAlarmState = "MANUAL_CLEAR"
	BlackholePortAlarmStateManualClearingUppercase     BlackholePortAlarmState = "MANUAL_CLEARING"
	BlackholePortAlarmStateManualClearingFailUppercase BlackholePortAlarmState = "MANUAL_CLEARING_FAIL"
	BlackholePortAlarmStateManualMitigatingUppercase   BlackholePortAlarmState = "MANUAL_MITIGATING"
	BlackholePortAlarmStateManualStartingUppercase     BlackholePortAlarmState = "MANUAL_STARTING"
	BlackholePortAlarmStateManualStartingFailUppercase BlackholePortAlarmState = "MANUAL_STARTING_FAIL"
	BlackholePortAlarmStateMitigatingUppercase         BlackholePortAlarmState = "MITIGATING"
	BlackholePortAlarmStateStartingUppercase           BlackholePortAlarmState = "STARTING"
	BlackholePortAlarmStateStartingFailUppercase       BlackholePortAlarmState = "STARTING_FAIL"
	BlackholePortAlarmStateStartWaitUppercase          BlackholePortAlarmState = "START_WAIT"
	BlackholePortAlarmStateAckReq                      BlackholePortAlarmState = "ack_req"
	BlackholePortAlarmStateAlarm                       BlackholePortAlarmState = "alarm"
	BlackholePortAlarmStateArchived                    BlackholePortAlarmState = "archived"
	BlackholePortAlarmStateClear                       BlackholePortAlarmState = "clear"
	BlackholePortAlarmStateClearing                    BlackholePortAlarmState = "clearing"
	BlackholePortAlarmStateClearingFail                BlackholePortAlarmState = "clearing_fail"
	BlackholePortAlarmStateEndGrace                    BlackholePortAlarmState = "end_grace"
	BlackholePortAlarmStateEndWait                     BlackholePortAlarmState = "end_wait"
	BlackholePortAlarmStateManualClear                 BlackholePortAlarmState = "manual_clear"
	BlackholePortAlarmStateManualClearing              BlackholePortAlarmState = "manual_clearing"
	BlackholePortAlarmStateManualClearingFail          BlackholePortAlarmState = "manual_clearing_fail"
	BlackholePortAlarmStateManualMitigating            BlackholePortAlarmState = "manual_mitigating"
	BlackholePortAlarmStateManualStarting              BlackholePortAlarmState = "manual_starting"
	BlackholePortAlarmStateManualStartingFail          BlackholePortAlarmState = "manual_starting_fail"
	BlackholePortAlarmStateMitigating                  BlackholePortAlarmState = "mitigating"
	BlackholePortAlarmStateStartWait                   BlackholePortAlarmState = "start_wait"
	BlackholePortAlarmStateStarting                    BlackholePortAlarmState = "starting"
	BlackholePortAlarmStateStartingFail                BlackholePortAlarmState = "starting_fail"
)

type Console struct {
	// Remote console information
	RemoteConsole ConsoleRemoteConsole `json:"remote_console,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RemoteConsole respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Console) RawJSON() string { return r.JSON.raw }
func (r *Console) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Remote console information
type ConsoleRemoteConsole struct {
	Protocol string `json:"protocol,required"`
	Type     string `json:"type,required"`
	URL      string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Protocol    respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsoleRemoteConsole) RawJSON() string { return r.JSON.raw }
func (r *ConsoleRemoteConsole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DDOSProfile struct {
	// Unique identifier for the DDoS protection profile
	ID int64 `json:"id,required"`
	// List of configured field values for the protection profile
	Fields []DDOSProfileField `json:"fields,required"`
	// Configuration options controlling profile activation and BGP routing
	Options DDOSProfileOptionList `json:"options,required"`
	// Complete template configuration data used for this profile
	ProfileTemplate DDOSProfileTemplate `json:"profile_template,required"`
	// Detailed description of the protection template used for this profile
	ProfileTemplateDescription string `json:"profile_template_description,required"`
	// List of network protocols and ports configured for protection
	Protocols []DDOSProfileProtocol `json:"protocols,required"`
	// Geographic site identifier where the protection is deployed
	Site string `json:"site,required"`
	// Current operational status and any error information for the profile
	Status DDOSProfileStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		Fields                     respjson.Field
		Options                    respjson.Field
		ProfileTemplate            respjson.Field
		ProfileTemplateDescription respjson.Field
		Protocols                  respjson.Field
		Site                       respjson.Field
		Status                     respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DDOSProfile) RawJSON() string { return r.JSON.raw }
func (r *DDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DDOSProfileProtocol struct {
	// Network port number for which protocols are configured
	Port string `json:"port,required"`
	// List of network protocols enabled on the specified port
	Protocols []string `json:"protocols,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Port        respjson.Field
		Protocols   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DDOSProfileProtocol) RawJSON() string { return r.JSON.raw }
func (r *DDOSProfileProtocol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DDOSProfileField struct {
	// Unique identifier for the DDoS protection field
	ID int64 `json:"id,required"`
	// ID of DDoS profile field
	BaseField int64 `json:"base_field,required"`
	// Predefined default value for the field if not specified
	Default string `json:"default,required"`
	// Detailed description explaining the field's purpose and usage guidelines
	Description string `json:"description,required"`
	// Name of DDoS profile field
	FieldName string `json:"field_name,required"`
	// Data type classification of the field (e.g., string, integer, array)
	FieldType string `json:"field_type,required"`
	// Complex value. Only one of 'value' or '`field_value`' must be specified.
	FieldValue any `json:"field_value,required"`
	// Human-readable name of the protection field
	Name string `json:"name,required"`
	// Indicates whether this field must be provided when creating a protection profile
	Required bool `json:"required,required"`
	// JSON schema defining validation rules and constraints for the field value
	ValidationSchema any `json:"validation_schema,required"`
	// Basic type value. Only one of 'value' or '`field_value`' must be specified.
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		BaseField        respjson.Field
		Default          respjson.Field
		Description      respjson.Field
		FieldName        respjson.Field
		FieldType        respjson.Field
		FieldValue       respjson.Field
		Name             respjson.Field
		Required         respjson.Field
		ValidationSchema respjson.Field
		Value            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DDOSProfileField) RawJSON() string { return r.JSON.raw }
func (r *DDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DDOSProfileOptionList struct {
	// Controls whether the DDoS protection profile is enabled and actively protecting
	// the resource
	Active bool `json:"active,required"`
	// Enables Border Gateway Protocol (BGP) routing for DDoS protection traffic
	Bgp bool `json:"bgp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Bgp         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DDOSProfileOptionList) RawJSON() string { return r.JSON.raw }
func (r *DDOSProfileOptionList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DDOSProfileStatus struct {
	// Detailed error message describing any issues with the profile operation
	ErrorDescription string `json:"error_description,required"`
	// Current operational status of the DDoS protection profile
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorDescription respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DDOSProfileStatus) RawJSON() string { return r.JSON.raw }
func (r *DDOSProfileStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DDOSProfileTemplate struct {
	// Unique identifier for the DDoS protection template
	ID int64 `json:"id,required"`
	// Detailed description explaining the template's purpose and use cases
	Description string `json:"description,required"`
	// List of configurable fields that define the template's protection parameters
	Fields []DDOSProfileTemplateField `json:"fields,required"`
	// Human-readable name of the protection template
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Fields      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DDOSProfileTemplate) RawJSON() string { return r.JSON.raw }
func (r *DDOSProfileTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DDOSProfileTemplateField struct {
	// Unique identifier for the DDoS protection field
	ID int64 `json:"id,required"`
	// Predefined default value for the field if not specified
	Default string `json:"default,required"`
	// Detailed description explaining the field's purpose and usage guidelines
	Description string `json:"description,required"`
	// Data type classification of the field (e.g., string, integer, array)
	FieldType string `json:"field_type,required"`
	// Human-readable name of the protection field
	Name string `json:"name,required"`
	// Indicates whether this field must be provided when creating a protection profile
	Required bool `json:"required,required"`
	// JSON schema defining validation rules and constraints for the field value
	ValidationSchema any `json:"validation_schema,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Default          respjson.Field
		Description      respjson.Field
		FieldType        respjson.Field
		Name             respjson.Field
		Required         respjson.Field
		ValidationSchema respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DDOSProfileTemplateField) RawJSON() string { return r.JSON.raw }
func (r *DDOSProfileTemplateField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for `fixed` addresses. This schema is used when fetching a single
// instance.
type FixedAddress struct {
	// IP address
	Addr string `json:"addr,required"`
	// Interface name. This field will be `null` if `with_interfaces_name=true` is not
	// set in the request when listing instances. It will also be `null` if the
	// `interface_name` was not specified during instance creation or when attaching
	// the interface.
	InterfaceName string `json:"interface_name,required"`
	// The unique identifier of the subnet associated with this address. Included only
	// in the response for a single-resource lookup (GET by ID). For the trunk
	// subports, this field is always set.
	SubnetID string `json:"subnet_id,required"`
	// The name of the subnet associated with this address. Included only in the
	// response for a single-resource lookup (GET by ID). For the trunk subports, this
	// field is always set.
	SubnetName string `json:"subnet_name,required"`
	// Type of the address
	Type constant.Fixed `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Addr          respjson.Field
		InterfaceName respjson.Field
		SubnetID      respjson.Field
		SubnetName    respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FixedAddress) RawJSON() string { return r.JSON.raw }
func (r *FixedAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for `fixed` addresses. This schema is used when listing instances. It
// omits the `subnet_name` and `subnet_id` fields.
type FixedAddressShort struct {
	// IP address
	Addr string `json:"addr,required"`
	// Interface name. This field will be `null` if `with_interfaces_name=true` is not
	// set in the request when listing instances. It will also be `null` if the
	// `interface_name` was not specified during instance creation or when attaching
	// the interface.
	InterfaceName string `json:"interface_name,required"`
	// Type of the address
	Type constant.Fixed `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Addr          respjson.Field
		InterfaceName respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FixedAddressShort) RawJSON() string { return r.JSON.raw }
func (r *FixedAddressShort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FlavorHardwareDescription struct {
	// Human-readable CPU description
	CPU string `json:"cpu,nullable"`
	// Human-readable disk description
	Disk string `json:"disk,nullable"`
	// Human-readable ephemeral disk description
	Ephemeral string `json:"ephemeral,nullable"`
	// Human-readable GPU description
	GPU string `json:"gpu,nullable"`
	// Human-readable NIC description
	Network string `json:"network,nullable"`
	// Human-readable RAM description
	Ram string `json:"ram,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU         respjson.Field
		Disk        respjson.Field
		Ephemeral   respjson.Field
		GPU         respjson.Field
		Network     respjson.Field
		Ram         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *FlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for `floating` addresses.
type FloatingAddress struct {
	// Address
	Addr string `json:"addr,required"`
	// Type of the address
	Type constant.Floating `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Addr        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatingAddress) RawJSON() string { return r.JSON.raw }
func (r *FloatingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIP struct {
	// Floating IP ID
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the floating IP was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required" format:"uuid4"`
	// IP address of the port the floating IP is attached to
	FixedIPAddress string `json:"fixed_ip_address,required" format:"ipvanyaddress"`
	// IP Address of the floating IP
	FloatingIPAddress string `json:"floating_ip_address,required" format:"ipvanyaddress"`
	// Port ID the floating IP is attached to. The `fixed_ip_address` is the IP address
	// of the port.
	PortID string `json:"port_id,required" format:"uuid4"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
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
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,required" format:"uuid4"`
	// Datetime when the floating IP was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		CreatorTaskID     respjson.Field
		FixedIPAddress    respjson.Field
		FloatingIPAddress respjson.Field
		PortID            respjson.Field
		ProjectID         respjson.Field
		Region            respjson.Field
		RegionID          respjson.Field
		RouterID          respjson.Field
		Status            respjson.Field
		Tags              respjson.Field
		TaskID            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatingIP) RawJSON() string { return r.JSON.raw }
func (r *FloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIPStatus string

const (
	FloatingIPStatusActive FloatingIPStatus = "ACTIVE"
	FloatingIPStatusDown   FloatingIPStatus = "DOWN"
	FloatingIPStatusError  FloatingIPStatus = "ERROR"
)

type GPUImage struct {
	// Image ID
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the image was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Minimal boot volume required
	MinDisk int64 `json:"min_disk,required"`
	// Minimal VM RAM required
	MinRam int64 `json:"min_ram,required"`
	// Image name
	Name string `json:"name,required"`
	// Image status
	Status string `json:"status,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// Datetime when the image was updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Image visibility. Globally visible images are public
	Visibility string `json:"visibility,required"`
	// Image architecture type
	Architecture string `json:"architecture,nullable"`
	// OS Distribution
	OsDistro string `json:"os_distro,nullable"`
	// The operating system installed on the image
	OsType string `json:"os_type,nullable"`
	// OS version, i.e. 19.04 (for Ubuntu) or 9.4 for Debian
	OsVersion string `json:"os_version,nullable"`
	// Image size in bytes.
	Size int64 `json:"size"`
	// Whether the image supports SSH key or not
	SSHKey string `json:"ssh_key,nullable"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		MinDisk      respjson.Field
		MinRam       respjson.Field
		Name         respjson.Field
		Status       respjson.Field
		Tags         respjson.Field
		UpdatedAt    respjson.Field
		Visibility   respjson.Field
		Architecture respjson.Field
		OsDistro     respjson.Field
		OsType       respjson.Field
		OsVersion    respjson.Field
		Size         respjson.Field
		SSHKey       respjson.Field
		TaskID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUImage) RawJSON() string { return r.JSON.raw }
func (r *GPUImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUImageList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []GPUImage `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUImageList) RawJSON() string { return r.JSON.raw }
func (r *GPUImageList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type Image struct {
	// Image ID
	ID string `json:"id,required"`
	// Datetime when the image was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Disk format
	DiskFormat string `json:"disk_format,required"`
	// Minimal boot volume required
	MinDisk int64 `json:"min_disk,required"`
	// Minimal VM RAM required
	MinRam int64 `json:"min_ram,required"`
	// Image display name
	Name string `json:"name,required"`
	// OS Distribution, i.e. Debian, CentOS, Ubuntu, CoreOS etc.
	OsDistro string `json:"os_distro,required"`
	// The operating system installed on the image.
	//
	// Any of "linux", "windows".
	OsType ImageOsType `json:"os_type,required"`
	// OS version, i.e. 19.04 (for Ubuntu) or 9.4 for Debian
	OsVersion string `json:"os_version,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Image size in bytes
	Size int64 `json:"size,required"`
	// Image status, i.e. active
	Status string `json:"status,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// Datetime when the image was updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Image visibility. Globally visible images are public
	Visibility string `json:"visibility,required"`
	// An image architecture type: aarch64, `x86_64`
	//
	// Any of "aarch64", "x86_64".
	Architecture ImageArchitecture `json:"architecture"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable"`
	// Image description
	Description  string `json:"description,nullable"`
	DisplayOrder int64  `json:"display_order,nullable"`
	// Specifies the type of firmware with which to boot the guest.
	//
	// Any of "bios", "uefi".
	HwFirmwareType ImageHwFirmwareType `json:"hw_firmware_type,nullable"`
	// A virtual chipset type.
	//
	// Any of "pc", "q35".
	HwMachineType ImageHwMachineType `json:"hw_machine_type,nullable"`
	// Set to true if the image will be used by bare metal servers. Defaults to false.
	IsBaremetal bool `json:"is_baremetal,nullable"`
	// Whether the image supports SSH key or not
	//
	// Any of "allow", "deny", "required".
	SSHKey ImageSSHKey `json:"ssh_key,nullable"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		DiskFormat     respjson.Field
		MinDisk        respjson.Field
		MinRam         respjson.Field
		Name           respjson.Field
		OsDistro       respjson.Field
		OsType         respjson.Field
		OsVersion      respjson.Field
		ProjectID      respjson.Field
		Region         respjson.Field
		RegionID       respjson.Field
		Size           respjson.Field
		Status         respjson.Field
		Tags           respjson.Field
		UpdatedAt      respjson.Field
		Visibility     respjson.Field
		Architecture   respjson.Field
		CreatorTaskID  respjson.Field
		Description    respjson.Field
		DisplayOrder   respjson.Field
		HwFirmwareType respjson.Field
		HwMachineType  respjson.Field
		IsBaremetal    respjson.Field
		SSHKey         respjson.Field
		TaskID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Image) RawJSON() string { return r.JSON.raw }
func (r *Image) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The operating system installed on the image.
type ImageOsType string

const (
	ImageOsTypeLinux   ImageOsType = "linux"
	ImageOsTypeWindows ImageOsType = "windows"
)

// An image architecture type: aarch64, `x86_64`
type ImageArchitecture string

const (
	ImageArchitectureAarch64 ImageArchitecture = "aarch64"
	ImageArchitectureX86_64  ImageArchitecture = "x86_64"
)

// Specifies the type of firmware with which to boot the guest.
type ImageHwFirmwareType string

const (
	ImageHwFirmwareTypeBios ImageHwFirmwareType = "bios"
	ImageHwFirmwareTypeUefi ImageHwFirmwareType = "uefi"
)

// A virtual chipset type.
type ImageHwMachineType string

const (
	ImageHwMachineTypePc  ImageHwMachineType = "pc"
	ImageHwMachineTypeQ35 ImageHwMachineType = "q35"
)

// Whether the image supports SSH key or not
type ImageSSHKey string

const (
	ImageSSHKeyAllow    ImageSSHKey = "allow"
	ImageSSHKeyDeny     ImageSSHKey = "deny"
	ImageSSHKeyRequired ImageSSHKey = "required"
)

type ImageList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Image `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageList) RawJSON() string { return r.JSON.raw }
func (r *ImageList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Instance struct {
	// Instance ID
	ID string `json:"id,required" format:"uuid4"`
	// Map of `network_name` to list of addresses in that network
	Addresses map[string][]InstanceAddressUnion `json:"addresses,required"`
	// IP addresses of the instances that are blackholed by DDoS mitigation system
	BlackholePorts []BlackholePort `json:"blackhole_ports,required"`
	// Datetime when instance was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required"`
	// Advanced DDoS protection profile. It is always `null` if query parameter
	// `with_ddos=true` is not set.
	DDOSProfile DDOSProfile `json:"ddos_profile,required"`
	// Fixed IP assigned to instance
	FixedIPAssignments []InstanceFixedIPAssignment `json:"fixed_ip_assignments,required"`
	// Flavor
	Flavor InstanceFlavorUnion `json:"flavor,required"`
	// Instance description
	InstanceDescription string `json:"instance_description,required"`
	// Instance isolation information
	InstanceIsolation InstanceIsolation `json:"instance_isolation,required"`
	// Instance name
	Name string `json:"name,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Security groups
	SecurityGroups []InstanceSecurityGroup `json:"security_groups,required"`
	// SSH key assigned to instance
	SSHKeyName string `json:"ssh_key_name,required"`
	// Instance status
	//
	// Any of "ACTIVE", "BUILD", "DELETED", "ERROR", "HARD_REBOOT", "MIGRATING",
	// "PASSWORD", "PAUSED", "REBOOT", "REBUILD", "RESCUE", "RESIZE", "REVERT_RESIZE",
	// "SHELVED", "SHELVED_OFFLOADED", "SHUTOFF", "SOFT_DELETED", "SUSPENDED",
	// "UNKNOWN", "VERIFY_RESIZE".
	Status InstanceStatus `json:"status,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,required"`
	// Task state
	TaskState string `json:"task_state,required"`
	// Virtual machine state (active)
	//
	// Any of "active", "building", "deleted", "error", "paused", "rescued", "resized",
	// "shelved", "shelved_offloaded", "soft-deleted", "stopped", "suspended".
	VmState InstanceVmState `json:"vm_state,required"`
	// List of volumes
	Volumes []InstanceVolume `json:"volumes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Addresses           respjson.Field
		BlackholePorts      respjson.Field
		CreatedAt           respjson.Field
		CreatorTaskID       respjson.Field
		DDOSProfile         respjson.Field
		FixedIPAssignments  respjson.Field
		Flavor              respjson.Field
		InstanceDescription respjson.Field
		InstanceIsolation   respjson.Field
		Name                respjson.Field
		ProjectID           respjson.Field
		Region              respjson.Field
		RegionID            respjson.Field
		SecurityGroups      respjson.Field
		SSHKeyName          respjson.Field
		Status              respjson.Field
		Tags                respjson.Field
		TaskID              respjson.Field
		TaskState           respjson.Field
		VmState             respjson.Field
		Volumes             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Instance) RawJSON() string { return r.JSON.raw }
func (r *Instance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InstanceAddressUnion contains all possible properties and values from
// [FloatingAddress], [FixedAddressShort], [FixedAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InstanceAddressUnion struct {
	Addr          string `json:"addr"`
	Type          string `json:"type"`
	InterfaceName string `json:"interface_name"`
	// This field is from variant [FixedAddress].
	SubnetID string `json:"subnet_id"`
	// This field is from variant [FixedAddress].
	SubnetName string `json:"subnet_name"`
	JSON       struct {
		Addr          respjson.Field
		Type          respjson.Field
		InterfaceName respjson.Field
		SubnetID      respjson.Field
		SubnetName    respjson.Field
		raw           string
	} `json:"-"`
}

func (u InstanceAddressUnion) AsFloatingIPAddress() (v FloatingAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InstanceAddressUnion) AsFixedIPAddressShort() (v FixedAddressShort) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InstanceAddressUnion) AsFixedIPAddress() (v FixedAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InstanceAddressUnion) RawJSON() string { return u.JSON.raw }

func (r *InstanceAddressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceFixedIPAssignment struct {
	// Is network external
	External bool `json:"external,required"`
	// Ip address
	IPAddress string `json:"ip_address,required"`
	// Interface subnet id
	SubnetID string `json:"subnet_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		External    respjson.Field
		IPAddress   respjson.Field
		SubnetID    respjson.Field
		ExtraFields map[string]respjson.Field
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
		Architecture        respjson.Field
		FlavorID            respjson.Field
		FlavorName          respjson.Field
		HardwareDescription respjson.Field
		OsType              respjson.Field
		Ram                 respjson.Field
		Vcpus               respjson.Field
		ResourceClass       respjson.Field
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
		Ram     respjson.Field
		Vcpus   respjson.Field
		CPU     respjson.Field
		Disk    respjson.Field
		License respjson.Field
		Network respjson.Field
		GPU     respjson.Field
		raw     string
	} `json:"-"`
}

func (r *InstanceFlavorUnionHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instances flavor schema embedded into instance schema
type InstanceFlavorInstanceFlavor struct {
	// CPU architecture
	Architecture string `json:"architecture,required"`
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id,required"`
	// Flavor name
	FlavorName string `json:"flavor_name,required"`
	// Additional hardware description
	HardwareDescription InstanceFlavorInstanceFlavorHardwareDescription `json:"hardware_description,required"`
	// Flavor operating system
	OsType string `json:"os_type,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Virtual CPU count. For bare metal flavors, it's a physical CPU count
	Vcpus int64 `json:"vcpus,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture        respjson.Field
		FlavorID            respjson.Field
		FlavorName          respjson.Field
		HardwareDescription respjson.Field
		OsType              respjson.Field
		Ram                 respjson.Field
		Vcpus               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorInstanceFlavor) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorInstanceFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional hardware description
type InstanceFlavorInstanceFlavorHardwareDescription struct {
	// RAM description
	Ram string `json:"ram,required"`
	// CPU description
	Vcpus string `json:"vcpus,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ram         respjson.Field
		Vcpus       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorInstanceFlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorInstanceFlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal flavor schema embedded into instance schema
type InstanceFlavorBareMetalFlavor struct {
	// CPU architecture
	Architecture string `json:"architecture,required"`
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id,required"`
	// Flavor name
	FlavorName string `json:"flavor_name,required"`
	// Additional hardware description
	HardwareDescription InstanceFlavorBareMetalFlavorHardwareDescription `json:"hardware_description,required"`
	// Operating system
	OsType string `json:"os_type,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Flavor resource class for mapping to hardware capacity
	ResourceClass string `json:"resource_class,required"`
	// Virtual CPU count. For bare metal flavors, it's a physical CPU count
	Vcpus int64 `json:"vcpus,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture        respjson.Field
		FlavorID            respjson.Field
		FlavorName          respjson.Field
		HardwareDescription respjson.Field
		OsType              respjson.Field
		Ram                 respjson.Field
		ResourceClass       respjson.Field
		Vcpus               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorBareMetalFlavor) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorBareMetalFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional hardware description
type InstanceFlavorBareMetalFlavorHardwareDescription struct {
	// Human-readable CPU description
	CPU string `json:"cpu,required"`
	// Human-readable disk description
	Disk string `json:"disk,required"`
	// If the flavor is licensed, this field contains the license type
	License string `json:"license,required"`
	// Human-readable NIC description
	Network string `json:"network,required"`
	// Human-readable RAM description
	Ram string `json:"ram,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU         respjson.Field
		Disk        respjson.Field
		License     respjson.Field
		Network     respjson.Field
		Ram         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorBareMetalFlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorBareMetalFlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPU cluster flavor schema embedded into instance schema
type InstanceFlavorGPUClusterFlavor struct {
	// CPU architecture
	Architecture string `json:"architecture,required"`
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id,required"`
	// Flavor name
	FlavorName string `json:"flavor_name,required"`
	// Additional hardware description
	HardwareDescription InstanceFlavorGPUClusterFlavorHardwareDescription `json:"hardware_description,required"`
	// Operating system
	OsType string `json:"os_type,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Flavor resource class for mapping to hardware capacity
	ResourceClass string `json:"resource_class,required"`
	// Virtual CPU count. For bare metal flavors, it's a physical CPU count
	Vcpus int64 `json:"vcpus,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture        respjson.Field
		FlavorID            respjson.Field
		FlavorName          respjson.Field
		HardwareDescription respjson.Field
		OsType              respjson.Field
		Ram                 respjson.Field
		ResourceClass       respjson.Field
		Vcpus               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorGPUClusterFlavor) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorGPUClusterFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional hardware description
type InstanceFlavorGPUClusterFlavorHardwareDescription struct {
	// Human-readable CPU description
	CPU string `json:"cpu,required"`
	// Human-readable disk description
	Disk string `json:"disk,required"`
	// Human-readable GPU description
	GPU string `json:"gpu,required"`
	// If the flavor is licensed, this field contains the license type
	License string `json:"license,required"`
	// Human-readable NIC description
	Network string `json:"network,required"`
	// Human-readable RAM description
	Ram string `json:"ram,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU         respjson.Field
		Disk        respjson.Field
		GPU         respjson.Field
		License     respjson.Field
		Network     respjson.Field
		Ram         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceFlavorGPUClusterFlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *InstanceFlavorGPUClusterFlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceSecurityGroup struct {
	// Name.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceSecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *InstanceSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance status
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

// Virtual machine state (active)
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

type InstanceVolume struct {
	// Volume ID
	ID string `json:"id,required"`
	// Whether the volume is deleted together with the VM
	DeleteOnTermination bool `json:"delete_on_termination,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		DeleteOnTermination respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceVolume) RawJSON() string { return r.JSON.raw }
func (r *InstanceVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceIsolation struct {
	// The reason of instance isolation if it is isolated from external internet.
	Reason string `json:"reason,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceIsolation) RawJSON() string { return r.JSON.raw }
func (r *InstanceIsolation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Instance `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceList) RawJSON() string { return r.JSON.raw }
func (r *InstanceList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceMetricsTimeUnit string

const (
	InstanceMetricsTimeUnitDay  InstanceMetricsTimeUnit = "day"
	InstanceMetricsTimeUnitHour InstanceMetricsTimeUnit = "hour"
)

type InterfaceIPFamily string

const (
	InterfaceIPFamilyDual InterfaceIPFamily = "dual"
	InterfaceIPFamilyIpv4 InterfaceIPFamily = "ipv4"
	InterfaceIPFamilyIpv6 InterfaceIPFamily = "ipv6"
)

type IPAssignment struct {
	// IP address
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
func (r IPAssignment) RawJSON() string { return r.JSON.raw }
func (r *IPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPVersion int64

const (
	IPVersion4 IPVersion = 4
	IPVersion6 IPVersion = 6
)

type LaasIndexRetentionPolicy struct {
	// Duration of days for which logs must be kept.
	Period int64 `json:"period,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Period      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaasIndexRetentionPolicy) RawJSON() string { return r.JSON.raw }
func (r *LaasIndexRetentionPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LaasIndexRetentionPolicy to a
// LaasIndexRetentionPolicyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LaasIndexRetentionPolicyParam.Overrides()
func (r LaasIndexRetentionPolicy) ToParam() LaasIndexRetentionPolicyParam {
	return param.Override[LaasIndexRetentionPolicyParam](json.RawMessage(r.RawJSON()))
}

// The property Period is required.
type LaasIndexRetentionPolicyParam struct {
	// Duration of days for which logs must be kept.
	Period param.Opt[int64] `json:"period,omitzero,required"`
	paramObj
}

func (r LaasIndexRetentionPolicyParam) MarshalJSON() (data []byte, err error) {
	type shadow LaasIndexRetentionPolicyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaasIndexRetentionPolicyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancer struct {
	// Load balancer ID
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the load balancer was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Load balancer name
	Name string `json:"name,required"`
	// Load balancer operating status
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Load balancer lifecycle status
	//
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	TagsV2 []Tag `json:"tags_v2,required"`
	// List of additional IP addresses
	AdditionalVips []LoadBalancerAdditionalVip `json:"additional_vips"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// Loadbalancer advanced DDoS protection profile.
	DDOSProfile DDOSProfile `json:"ddos_profile,nullable"`
	// Load balancer flavor (if not default)
	Flavor LoadBalancerFlavor `json:"flavor,nullable"`
	// List of assigned floating IPs
	FloatingIPs []FloatingIP `json:"floating_ips"`
	// Load balancer listeners
	Listeners []LoadBalancerListener `json:"listeners"`
	// Logging configuration
	Logging Logging `json:"logging,nullable"`
	// Preferred option to establish connectivity between load balancer and its pools
	// members
	//
	// Any of "L2", "L3".
	PreferredConnectivity LoadBalancerMemberConnectivity `json:"preferred_connectivity"`
	// Statistics of load balancer.
	Stats LoadBalancerStatistics `json:"stats,nullable"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// Datetime when the load balancer was last updated
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// Load balancer IP address
	VipAddress string `json:"vip_address,nullable" format:"ipvanyaddress"`
	// Load balancer IP family
	//
	// Any of "dual", "ipv4", "ipv6".
	VipIPFamily InterfaceIPFamily `json:"vip_ip_family,nullable"`
	// The ID of the Virtual IP (VIP) port.
	VipPortID string `json:"vip_port_id,nullable" format:"uuid4"`
	// List of VRRP IP addresses
	VrrpIPs []LoadBalancerVrrpIP `json:"vrrp_ips"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		Name                  respjson.Field
		OperatingStatus       respjson.Field
		ProjectID             respjson.Field
		ProvisioningStatus    respjson.Field
		Region                respjson.Field
		RegionID              respjson.Field
		TagsV2                respjson.Field
		AdditionalVips        respjson.Field
		CreatorTaskID         respjson.Field
		DDOSProfile           respjson.Field
		Flavor                respjson.Field
		FloatingIPs           respjson.Field
		Listeners             respjson.Field
		Logging               respjson.Field
		PreferredConnectivity respjson.Field
		Stats                 respjson.Field
		TaskID                respjson.Field
		UpdatedAt             respjson.Field
		VipAddress            respjson.Field
		VipIPFamily           respjson.Field
		VipPortID             respjson.Field
		VrrpIPs               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancer) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerAdditionalVip struct {
	// IP address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// Subnet UUID
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
func (r LoadBalancerAdditionalVip) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerAdditionalVip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Load balancer flavor (if not default)
type LoadBalancerFlavor struct {
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id,required"`
	// Flavor name
	FlavorName string `json:"flavor_name,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Virtual CPU count. For bare metal flavors, it's a physical CPU count
	Vcpus int64 `json:"vcpus,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FlavorID    respjson.Field
		FlavorName  respjson.Field
		Ram         respjson.Field
		Vcpus       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerFlavor) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerListener struct {
	// Listener ID
	ID string `json:"id,required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerListener) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerListener) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerVrrpIP struct {
	// IP address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// LoadBalancer instance role to which VRRP IP belong
	//
	// Any of "BACKUP", "MASTER", "STANDALONE".
	Role LoadBalancerInstanceRole `json:"role,required"`
	// Subnet UUID
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress   respjson.Field
		Role        respjson.Field
		SubnetID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerVrrpIP) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerVrrpIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerInstanceRole string

const (
	LoadBalancerInstanceRoleBackup     LoadBalancerInstanceRole = "BACKUP"
	LoadBalancerInstanceRoleMaster     LoadBalancerInstanceRole = "MASTER"
	LoadBalancerInstanceRoleStandalone LoadBalancerInstanceRole = "STANDALONE"
)

type LoadBalancerMemberConnectivity string

const (
	LoadBalancerMemberConnectivityL2 LoadBalancerMemberConnectivity = "L2"
	LoadBalancerMemberConnectivityL3 LoadBalancerMemberConnectivity = "L3"
)

type LoadBalancerOperatingStatus string

const (
	LoadBalancerOperatingStatusDegraded  LoadBalancerOperatingStatus = "DEGRADED"
	LoadBalancerOperatingStatusDraining  LoadBalancerOperatingStatus = "DRAINING"
	LoadBalancerOperatingStatusError     LoadBalancerOperatingStatus = "ERROR"
	LoadBalancerOperatingStatusNoMonitor LoadBalancerOperatingStatus = "NO_MONITOR"
	LoadBalancerOperatingStatusOffline   LoadBalancerOperatingStatus = "OFFLINE"
	LoadBalancerOperatingStatusOnline    LoadBalancerOperatingStatus = "ONLINE"
)

type LoadBalancerStatistics struct {
	// Currently active connections
	ActiveConnections int64 `json:"active_connections,required"`
	// Total bytes received
	BytesIn int64 `json:"bytes_in,required"`
	// Total bytes sent
	BytesOut int64 `json:"bytes_out,required"`
	// Total requests that were unable to be fulfilled
	RequestErrors int64 `json:"request_errors,required"`
	// Total connections handled
	TotalConnections int64 `json:"total_connections,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveConnections respjson.Field
		BytesIn           respjson.Field
		BytesOut          respjson.Field
		RequestErrors     respjson.Field
		TotalConnections  respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerStatistics) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Logging struct {
	// ID of the region in which the logs will be stored
	DestinationRegionID int64 `json:"destination_region_id,required"`
	// Indicates if log streaming is enabled or disabled
	Enabled bool `json:"enabled,required"`
	// The topic name to stream logs to
	TopicName string `json:"topic_name,required"`
	// Logs retention policy
	RetentionPolicy LaasIndexRetentionPolicy `json:"retention_policy,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationRegionID respjson.Field
		Enabled             respjson.Field
		TopicName           respjson.Field
		RetentionPolicy     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Logging) RawJSON() string { return r.JSON.raw }
func (r *Logging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Network struct {
	// Network ID
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the network was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required" format:"uuid4"`
	// True if network has `is_default` attribute
	Default bool `json:"default,required"`
	// True if the network `router:external` attribute
	External bool `json:"external,required"`
	// MTU (maximum transmission unit). Default value is 1450
	Mtu int64 `json:"mtu,required"`
	// Network name
	Name string `json:"name,required"`
	// Indicates `port_security_enabled` status of all newly created in the network
	// ports.
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Id of network segment
	SegmentationID int64 `json:"segmentation_id,required"`
	// True when the network is shared with your project by external owner
	Shared bool `json:"shared,required"`
	// List of subnetworks
	Subnets []string `json:"subnets,required" format:"uuid4"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,required" format:"uuid4"`
	// Network type (vlan, vxlan)
	Type string `json:"type,required"`
	// Datetime when the network was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		CreatorTaskID       respjson.Field
		Default             respjson.Field
		External            respjson.Field
		Mtu                 respjson.Field
		Name                respjson.Field
		PortSecurityEnabled respjson.Field
		ProjectID           respjson.Field
		Region              respjson.Field
		RegionID            respjson.Field
		SegmentationID      respjson.Field
		Shared              respjson.Field
		Subnets             respjson.Field
		Tags                respjson.Field
		TaskID              respjson.Field
		Type                respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Network) RawJSON() string { return r.JSON.raw }
func (r *Network) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkDetails struct {
	// Network ID
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the network was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required" format:"uuid4"`
	// True if network has `is_default` attribute
	Default bool `json:"default,required"`
	// True if the network `router:external` attribute
	External bool `json:"external,required"`
	// MTU (maximum transmission unit). Default value is 1450
	Mtu int64 `json:"mtu,required"`
	// Network name
	Name string `json:"name,required"`
	// Indicates `port_security_enabled` status of all newly created in the network
	// ports.
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Id of network segment
	SegmentationID int64 `json:"segmentation_id,required"`
	// True when the network is shared with your project by external owner
	Shared bool `json:"shared,required"`
	// List of subnets associated with the network
	Subnets []Subnet `json:"subnets,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,required" format:"uuid4"`
	// Network type (vlan, vxlan)
	Type string `json:"type,required"`
	// Datetime when the network was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		CreatorTaskID       respjson.Field
		Default             respjson.Field
		External            respjson.Field
		Mtu                 respjson.Field
		Name                respjson.Field
		PortSecurityEnabled respjson.Field
		ProjectID           respjson.Field
		Region              respjson.Field
		RegionID            respjson.Field
		SegmentationID      respjson.Field
		Shared              respjson.Field
		Subnets             respjson.Field
		Tags                respjson.Field
		TaskID              respjson.Field
		Type                respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkDetails) RawJSON() string { return r.JSON.raw }
func (r *NetworkDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkInterface struct {
	// Group of subnet masks and/or IP addresses that share the current IP as VIP
	AllowedAddressPairs []AllowedAddressPairs `json:"allowed_address_pairs,required"`
	// Bodies of floating IPs that are NAT-ing IPs of this port
	FloatingipDetails []FloatingIP `json:"floatingip_details,required"`
	// IP addresses assigned to this port
	IPAssignments []IPAssignment `json:"ip_assignments,required"`
	// Body of the network this port is attached to
	NetworkDetails NetworkDetails `json:"network_details,required"`
	// ID of the network the port is attached to
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// ID of virtual ethernet port object
	PortID string `json:"port_id,required" format:"uuid4"`
	// Port security status
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// body of ports that are included into trunk port
	SubPorts []NetworkInterfaceSubPort `json:"sub_ports,required"`
	// Interface name
	InterfaceName string `json:"interface_name,nullable"`
	// MAC address of the virtual port
	MacAddress string `json:"mac_address,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedAddressPairs respjson.Field
		FloatingipDetails   respjson.Field
		IPAssignments       respjson.Field
		NetworkDetails      respjson.Field
		NetworkID           respjson.Field
		PortID              respjson.Field
		PortSecurityEnabled respjson.Field
		SubPorts            respjson.Field
		InterfaceName       respjson.Field
		MacAddress          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkInterface) RawJSON() string { return r.JSON.raw }
func (r *NetworkInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkInterfaceSubPort struct {
	// Group of subnet masks and/or IP addresses that share the current IP as VIP
	AllowedAddressPairs []AllowedAddressPairs `json:"allowed_address_pairs,required"`
	// Bodies of floating IPs that are NAT-ing IPs of this port
	FloatingipDetails []FloatingIP `json:"floatingip_details,required"`
	// IP addresses assigned to this port
	IPAssignments []IPAssignment `json:"ip_assignments,required"`
	// Body of the network this port is attached to
	NetworkDetails NetworkDetails `json:"network_details,required"`
	// ID of the network the port is attached to
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// ID of virtual ethernet port object
	PortID string `json:"port_id,required" format:"uuid4"`
	// Port security status
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// id of network segment
	SegmentationID int64 `json:"segmentation_id,required"`
	// type of network segment
	SegmentationType string `json:"segmentation_type,required"`
	// Interface name
	InterfaceName string `json:"interface_name,nullable"`
	// MAC address of the virtual port
	MacAddress string `json:"mac_address,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedAddressPairs respjson.Field
		FloatingipDetails   respjson.Field
		IPAssignments       respjson.Field
		NetworkDetails      respjson.Field
		NetworkID           respjson.Field
		PortID              respjson.Field
		PortSecurityEnabled respjson.Field
		SegmentationID      respjson.Field
		SegmentationType    respjson.Field
		InterfaceName       respjson.Field
		MacAddress          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkInterfaceSubPort) RawJSON() string { return r.JSON.raw }
func (r *NetworkInterfaceSubPort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkInterfaceList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []NetworkInterface `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkInterfaceList) RawJSON() string { return r.JSON.raw }
func (r *NetworkInterfaceList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type Route struct {
	// CIDR of destination IPv4 subnet.
	Destination string `json:"destination,required" format:"ipvanynetwork"`
	// IPv4 address to forward traffic to if it's destination IP matches 'destination'
	// CIDR.
	Nexthop string `json:"nexthop,required" format:"ipvanyaddress"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Destination respjson.Field
		Nexthop     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Route) RawJSON() string { return r.JSON.raw }
func (r *Route) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Subnet struct {
	// CIDR
	Cidr string `json:"cidr,required" format:"ipvanynetwork"`
	// Datetime when the subnet was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// True if DHCP should be enabled
	EnableDhcp bool `json:"enable_dhcp,required"`
	// IP version
	//
	// Any of 4, 6.
	IPVersion int64 `json:"ip_version,required"`
	// Subnet name
	Name string `json:"name,required"`
	// Network ID
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// Datetime when the subnet was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Subnet id.
	ID string `json:"id,nullable" format:"uuid4"`
	// Number of available ips in subnet
	AvailableIPs int64 `json:"available_ips,nullable"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// List IP addresses of a DNS resolver reachable from the network
	DNSNameservers []string `json:"dns_nameservers,nullable" format:"ipvanyaddress"`
	// Default GW IPv4 address, advertised in DHCP routes of this subnet. If null, no
	// gateway is advertised by this subnet.
	GatewayIP string `json:"gateway_ip,nullable" format:"ipvanyaddress"`
	// Deprecated. Always returns `false`.
	//
	// Deprecated: deprecated
	HasRouter bool `json:"has_router"`
	// List of custom static routes to advertise via DHCP.
	HostRoutes []Route `json:"host_routes,nullable"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,nullable" format:"uuid4"`
	// Total number of ips in subnet
	TotalIPs int64 `json:"total_ips,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cidr           respjson.Field
		CreatedAt      respjson.Field
		EnableDhcp     respjson.Field
		IPVersion      respjson.Field
		Name           respjson.Field
		NetworkID      respjson.Field
		ProjectID      respjson.Field
		Region         respjson.Field
		RegionID       respjson.Field
		Tags           respjson.Field
		UpdatedAt      respjson.Field
		ID             respjson.Field
		AvailableIPs   respjson.Field
		CreatorTaskID  respjson.Field
		DNSNameservers respjson.Field
		GatewayIP      respjson.Field
		HasRouter      respjson.Field
		HostRoutes     respjson.Field
		TaskID         respjson.Field
		TotalIPs       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Subnet) RawJSON() string { return r.JSON.raw }
func (r *Subnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A tag is a key-value pair that can be associated with a resource, enabling
// efficient filtering and grouping for better organization and management. Some
// tags are read-only and cannot be modified by the user. Tags are also integrated
// with cost reports, allowing cost data to be filtered based on tag keys or
// values.
type Tag struct {
	// Tag key. The maximum size for a key is 255 bytes.
	Key string `json:"key,required"`
	// If true, the tag is read-only and cannot be modified by the user
	ReadOnly bool `json:"read_only,required"`
	// Tag value. The maximum size for a value is 1024 bytes.
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		ReadOnly    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tag) RawJSON() string { return r.JSON.raw }
func (r *Tag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagUpdateMap map[string]string

type TaskIDList struct {
	// List of task IDs representing asynchronous operations. Use these IDs to monitor
	// operation progress:
	//
	//   - `GET /v1/tasks/{`task_id`}` - Check individual task status and details Poll
	//     task status until completion (`FINISHED`/`ERROR`) before proceeding with
	//     dependent operations.
	Tasks []string `json:"tasks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tasks       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskIDList) RawJSON() string { return r.JSON.raw }
func (r *TaskIDList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
