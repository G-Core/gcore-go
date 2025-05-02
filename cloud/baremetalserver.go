// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
	"github.com/stainless-sdks/gcore-go/shared/constant"
	"github.com/tidwall/gjson"
)

// BaremetalServerService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBaremetalServerService] method instead.
type BaremetalServerService struct {
	Options []option.RequestOption
}

// NewBaremetalServerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBaremetalServerService(opts ...option.RequestOption) (r BaremetalServerService) {
	r = BaremetalServerService{}
	r.Options = opts
	return
}

// Create a new bare metal server or multiple servers
func (r *BaremetalServerService) New(ctx context.Context, params BaremetalServerNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/bminstances/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List bare metal servers
func (r *BaremetalServerService) List(ctx context.Context, params BaremetalServerListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[BaremetalServer], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/bminstances/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List bare metal servers
func (r *BaremetalServerService) ListAutoPaging(ctx context.Context, params BaremetalServerListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[BaremetalServer] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Rebuild bare metal server
func (r *BaremetalServerService) Rebuild(ctx context.Context, serverID string, params BaremetalServerRebuildParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if serverID == "" {
		err = errors.New("missing required server_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/bminstances/%v/%v/%s/rebuild", params.ProjectID.Value, params.RegionID.Value, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// IP addresses of the trunk port and its subports.
type BaremetalFixedAddress struct {
	// Address
	Addr string `json:"addr,required"`
	// Interface name. This field will be `null` if `with_interfaces_name=true` is not
	// set in the request when listing servers. It will also be `null` if the
	// `interface_name` was not specified during server creation or when attaching the
	// interface.
	InterfaceName string `json:"interface_name,required"`
	// The unique identifier of the subnet associated with this address.
	SubnetID string `json:"subnet_id,required"`
	// The name of the subnet associated with this address.
	SubnetName string `json:"subnet_name,required"`
	// Type of the address
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
func (r BaremetalFixedAddress) RawJSON() string { return r.JSON.raw }
func (r *BaremetalFixedAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalFloatingAddress struct {
	// Address
	Addr string `json:"addr,required"`
	// Type of the address
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
func (r BaremetalFloatingAddress) RawJSON() string { return r.JSON.raw }
func (r *BaremetalFloatingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalServer struct {
	// Map of network_name to list of addresses in that network
	Addresses map[string][]BaremetalServerAddressUnion `json:"addresses,required"`
	// IP addresses of the instances that are blackholed by DDoS mitigation system
	BlackholePorts []BlackholePort `json:"blackhole_ports,required"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required"`
	// Bare metal advanced DDoS protection profile. It is always `null` if query
	// parameter `with_ddos=true` is not set.
	DDOSProfile DDOSProfile `json:"ddos_profile,required"`
	// Fixed IP assigned to instance
	FixedIPAssignments []BaremetalServerFixedIPAssignment `json:"fixed_ip_assignments,required"`
	// Flavor
	Flavor BaremetalServerFlavor `json:"flavor,required"`
	// Datetime when bare metal server was created
	InstanceCreated time.Time `json:"instance_created,required" format:"date-time"`
	// Bare metal server description
	InstanceDescription string `json:"instance_description,required"`
	// Bare metal server ID
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	// Instance isolation information
	InstanceIsolation InstanceIsolation `json:"instance_isolation,required"`
	// Bare metal server name
	InstanceName string `json:"instance_name,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// SSH key assigned to bare metal server
	SSHKeyName string `json:"ssh_key_name,required"`
	// Bare metal server status
	//
	// Any of "ACTIVE", "BUILD", "DELETED", "ERROR", "HARD_REBOOT", "MIGRATING",
	// "PASSWORD", "PAUSED", "REBOOT", "REBUILD", "RESCUE", "RESIZE", "REVERT_RESIZE",
	// "SHELVED", "SHELVED_OFFLOADED", "SHUTOFF", "SOFT_DELETED", "SUSPENDED",
	// "UNKNOWN", "VERIFY_RESIZE".
	Status BaremetalServerStatus `json:"status,required"`
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
	// Bare metal server state
	//
	// Any of "active", "building", "deleted", "error", "paused", "rescued", "resized",
	// "shelved", "shelved_offloaded", "soft-deleted", "stopped", "suspended".
	VmState BaremetalServerVmState `json:"vm_state,required"`
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
		ProjectID           resp.Field
		Region              resp.Field
		RegionID            resp.Field
		SSHKeyName          resp.Field
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
func (r BaremetalServer) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BaremetalServerAddressUnion contains all possible properties and values from
// [BaremetalFloatingAddress], [BaremetalFixedAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BaremetalServerAddressUnion struct {
	Addr string `json:"addr"`
	Type string `json:"type"`
	// This field is from variant [BaremetalFixedAddress].
	InterfaceName string `json:"interface_name"`
	// This field is from variant [BaremetalFixedAddress].
	SubnetID string `json:"subnet_id"`
	// This field is from variant [BaremetalFixedAddress].
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

func (u BaremetalServerAddressUnion) AsFloatingIPAddress() (v BaremetalFloatingAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BaremetalServerAddressUnion) AsFixedIPAddress() (v BaremetalFixedAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BaremetalServerAddressUnion) RawJSON() string { return u.JSON.raw }

func (r *BaremetalServerAddressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalServerFixedIPAssignment struct {
	// Is network external
	External bool `json:"external,required"`
	// Ip address
	IPAddress string `json:"ip_address,required"`
	// Interface subnet id
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
func (r BaremetalServerFixedIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerFixedIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor
type BaremetalServerFlavor struct {
	// CPU architecture
	Architecture string `json:"architecture,required"`
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id,required"`
	// Flavor name
	FlavorName string `json:"flavor_name,required"`
	// Additional hardware description
	HardwareDescription BaremetalServerFlavorHardwareDescription `json:"hardware_description,required"`
	// Operating system
	OsType string `json:"os_type,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Flavor resource class for mapping to hardware capacity
	ResourceClass string `json:"resource_class,required"`
	// Virtual CPU count. For bare metal flavors, it's a physical CPU count
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
func (r BaremetalServerFlavor) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional hardware description
type BaremetalServerFlavorHardwareDescription struct {
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
func (r BaremetalServerFlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerFlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal server status
type BaremetalServerStatus string

const (
	BaremetalServerStatusActive           BaremetalServerStatus = "ACTIVE"
	BaremetalServerStatusBuild            BaremetalServerStatus = "BUILD"
	BaremetalServerStatusDeleted          BaremetalServerStatus = "DELETED"
	BaremetalServerStatusError            BaremetalServerStatus = "ERROR"
	BaremetalServerStatusHardReboot       BaremetalServerStatus = "HARD_REBOOT"
	BaremetalServerStatusMigrating        BaremetalServerStatus = "MIGRATING"
	BaremetalServerStatusPassword         BaremetalServerStatus = "PASSWORD"
	BaremetalServerStatusPaused           BaremetalServerStatus = "PAUSED"
	BaremetalServerStatusReboot           BaremetalServerStatus = "REBOOT"
	BaremetalServerStatusRebuild          BaremetalServerStatus = "REBUILD"
	BaremetalServerStatusRescue           BaremetalServerStatus = "RESCUE"
	BaremetalServerStatusResize           BaremetalServerStatus = "RESIZE"
	BaremetalServerStatusRevertResize     BaremetalServerStatus = "REVERT_RESIZE"
	BaremetalServerStatusShelved          BaremetalServerStatus = "SHELVED"
	BaremetalServerStatusShelvedOffloaded BaremetalServerStatus = "SHELVED_OFFLOADED"
	BaremetalServerStatusShutoff          BaremetalServerStatus = "SHUTOFF"
	BaremetalServerStatusSoftDeleted      BaremetalServerStatus = "SOFT_DELETED"
	BaremetalServerStatusSuspended        BaremetalServerStatus = "SUSPENDED"
	BaremetalServerStatusUnknown          BaremetalServerStatus = "UNKNOWN"
	BaremetalServerStatusVerifyResize     BaremetalServerStatus = "VERIFY_RESIZE"
)

// Bare metal server state
type BaremetalServerVmState string

const (
	BaremetalServerVmStateActive           BaremetalServerVmState = "active"
	BaremetalServerVmStateBuilding         BaremetalServerVmState = "building"
	BaremetalServerVmStateDeleted          BaremetalServerVmState = "deleted"
	BaremetalServerVmStateError            BaremetalServerVmState = "error"
	BaremetalServerVmStatePaused           BaremetalServerVmState = "paused"
	BaremetalServerVmStateRescued          BaremetalServerVmState = "rescued"
	BaremetalServerVmStateResized          BaremetalServerVmState = "resized"
	BaremetalServerVmStateShelved          BaremetalServerVmState = "shelved"
	BaremetalServerVmStateShelvedOffloaded BaremetalServerVmState = "shelved_offloaded"
	BaremetalServerVmStateSoftDeleted      BaremetalServerVmState = "soft-deleted"
	BaremetalServerVmStateStopped          BaremetalServerVmState = "stopped"
	BaremetalServerVmStateSuspended        BaremetalServerVmState = "suspended"
)

type BaremetalServerNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// The flavor of the instance.
	Flavor string `json:"flavor,required"`
	// A list of network interfaces for the server. You can create one or more
	// interfaces—private, public, or both.
	Interfaces []BaremetalServerNewParamsInterfaceUnion `json:"interfaces,omitzero,required"`
	// Specifies the name of the SSH keypair, created via the `/v1/ssh_keys` endpoint.
	SSHKeyName param.Opt[string] `json:"ssh_key_name,omitzero"`
	// Apptemplate ID. Either `image_id` or `apptemplate_id` is required.
	ApptemplateID param.Opt[string] `json:"apptemplate_id,omitzero"`
	// Image ID. Either `image_id` or `apptemplate_id` is required.
	ImageID param.Opt[string] `json:"image_id,omitzero" format:"uuid4"`
	// For Linux instances, 'username' and 'password' are used to create a new user.
	// When only 'password' is provided, it is set as the password for the default user
	// of the image. For Windows instances, 'username' cannot be specified. Use the
	// 'password' field to set the password for the 'Admin' user on Windows. Use the
	// 'user_data' field to provide a script to create new users on Windows. The
	// password of the Admin user cannot be updated via 'user_data'.
	Password param.Opt[string] `json:"password,omitzero"`
	// String in base64 format. For Linux instances, 'user_data' is ignored when
	// 'password' field is provided. For Windows instances, Admin user password is set
	// by 'password' field and cannot be updated via 'user_data'. Examples of the
	// user_data: https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Opt[string] `json:"user_data,omitzero"`
	// For Linux instances, 'username' and 'password' are used to create a new user.
	// For Windows instances, 'username' cannot be specified. Use 'password' field to
	// set the password for the 'Admin' user on Windows.
	Username param.Opt[string] `json:"username,omitzero"`
	// Parameters for the application template if creating the instance from an
	// `apptemplate`.
	AppConfig any `json:"app_config,omitzero"`
	// Enable advanced DDoS protection for the server
	DDOSProfile BaremetalServerNewParamsDDOSProfile `json:"ddos_profile,omitzero"`
	// If you want server names to be automatically generated using IP octets, you can
	// specify name templates instead of setting names manually.Provide a list of
	// templated names that should be replaced using the selected template. The
	// following template formats are supported: `{ip_octets}`, `{two_ip_octets}`, and
	// `{one_ip_octet}`.
	NameTemplates []string `json:"name_templates,omitzero"`
	// List of server names. Specify one name to create a single server.
	Names []string `json:"names,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags TagUpdateList `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r BaremetalServerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BaremetalServerNewParamsInterfaceUnion struct {
	OfExternal        *BaremetalServerNewParamsInterfaceExternal        `json:",omitzero,inline"`
	OfSubnet          *BaremetalServerNewParamsInterfaceSubnet          `json:",omitzero,inline"`
	OfAnySubnet       *BaremetalServerNewParamsInterfaceAnySubnet       `json:",omitzero,inline"`
	OfReservedFixedIP *BaremetalServerNewParamsInterfaceReservedFixedIP `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u BaremetalServerNewParamsInterfaceUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u BaremetalServerNewParamsInterfaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[BaremetalServerNewParamsInterfaceUnion](u.OfExternal, u.OfSubnet, u.OfAnySubnet, u.OfReservedFixedIP)
}

func (u *BaremetalServerNewParamsInterfaceUnion) asAny() any {
	if !param.IsOmitted(u.OfExternal) {
		return u.OfExternal
	} else if !param.IsOmitted(u.OfSubnet) {
		return u.OfSubnet
	} else if !param.IsOmitted(u.OfAnySubnet) {
		return u.OfAnySubnet
	} else if !param.IsOmitted(u.OfReservedFixedIP) {
		return u.OfReservedFixedIP
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalServerNewParamsInterfaceUnion) GetSubnetID() *string {
	if vt := u.OfSubnet; vt != nil {
		return &vt.SubnetID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalServerNewParamsInterfaceUnion) GetIPAddress() *string {
	if vt := u.OfAnySubnet; vt != nil && vt.IPAddress.IsPresent() {
		return &vt.IPAddress.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalServerNewParamsInterfaceUnion) GetPortID() *string {
	if vt := u.OfReservedFixedIP; vt != nil {
		return &vt.PortID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalServerNewParamsInterfaceUnion) GetType() *string {
	if vt := u.OfExternal; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSubnet; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAnySubnet; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfReservedFixedIP; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalServerNewParamsInterfaceUnion) GetInterfaceName() *string {
	if vt := u.OfExternal; vt != nil && vt.InterfaceName.IsPresent() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfSubnet; vt != nil && vt.InterfaceName.IsPresent() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfAnySubnet; vt != nil && vt.InterfaceName.IsPresent() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfReservedFixedIP; vt != nil && vt.InterfaceName.IsPresent() {
		return &vt.InterfaceName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalServerNewParamsInterfaceUnion) GetIPFamily() *string {
	if vt := u.OfExternal; vt != nil {
		return (*string)(&vt.IPFamily)
	} else if vt := u.OfAnySubnet; vt != nil {
		return (*string)(&vt.IPFamily)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalServerNewParamsInterfaceUnion) GetPortGroup() *int64 {
	if vt := u.OfExternal; vt != nil && vt.PortGroup.IsPresent() {
		return &vt.PortGroup.Value
	} else if vt := u.OfSubnet; vt != nil && vt.PortGroup.IsPresent() {
		return &vt.PortGroup.Value
	} else if vt := u.OfAnySubnet; vt != nil && vt.PortGroup.IsPresent() {
		return &vt.PortGroup.Value
	} else if vt := u.OfReservedFixedIP; vt != nil && vt.PortGroup.IsPresent() {
		return &vt.PortGroup.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalServerNewParamsInterfaceUnion) GetNetworkID() *string {
	if vt := u.OfSubnet; vt != nil {
		return (*string)(&vt.NetworkID)
	} else if vt := u.OfAnySubnet; vt != nil {
		return (*string)(&vt.NetworkID)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u BaremetalServerNewParamsInterfaceUnion) GetFloatingIP() (res baremetalServerNewParamsInterfaceUnionFloatingIP) {
	if vt := u.OfSubnet; vt != nil {
		res.any = vt.FloatingIP.asAny()
	} else if vt := u.OfAnySubnet; vt != nil {
		res.any = vt.FloatingIP.asAny()
	} else if vt := u.OfReservedFixedIP; vt != nil {
		res.any = vt.FloatingIP.asAny()
	}
	return
}

// Can have the runtime types
// [*BaremetalServerNewParamsInterfaceSubnetFloatingIPNew],
// [*BaremetalServerNewParamsInterfaceSubnetFloatingIPExisting],
// [*BaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew],
// [*BaremetalServerNewParamsInterfaceAnySubnetFloatingIPExisting],
// [*BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew],
// [*BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting]
type baremetalServerNewParamsInterfaceUnionFloatingIP struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cloud.BaremetalServerNewParamsInterfaceSubnetFloatingIPNew:
//	case *cloud.BaremetalServerNewParamsInterfaceSubnetFloatingIPExisting:
//	case *cloud.BaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew:
//	case *cloud.BaremetalServerNewParamsInterfaceAnySubnetFloatingIPExisting:
//	case *cloud.BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew:
//	case *cloud.BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u baremetalServerNewParamsInterfaceUnionFloatingIP) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u baremetalServerNewParamsInterfaceUnionFloatingIP) GetSource() *string {
	switch vt := u.any.(type) {
	case *BaremetalServerNewParamsInterfaceSubnetFloatingIPUnion:
		return vt.GetSource()
	case *BaremetalServerNewParamsInterfaceAnySubnetFloatingIPUnion:
		return vt.GetSource()
	case *BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion:
		return vt.GetSource()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u baremetalServerNewParamsInterfaceUnionFloatingIP) GetExistingFloatingID() *string {
	switch vt := u.any.(type) {
	case *BaremetalServerNewParamsInterfaceSubnetFloatingIPUnion:
		return vt.GetExistingFloatingID()
	case *BaremetalServerNewParamsInterfaceAnySubnetFloatingIPUnion:
		return vt.GetExistingFloatingID()
	case *BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion:
		return vt.GetExistingFloatingID()
	}
	return nil
}

func init() {
	apijson.RegisterUnion[BaremetalServerNewParamsInterfaceUnion](
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BaremetalServerNewParamsInterfaceExternal{}),
			DiscriminatorValue: "external",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BaremetalServerNewParamsInterfaceSubnet{}),
			DiscriminatorValue: "subnet",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BaremetalServerNewParamsInterfaceAnySubnet{}),
			DiscriminatorValue: "any_subnet",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BaremetalServerNewParamsInterfaceReservedFixedIP{}),
			DiscriminatorValue: "reserved_fixed_ip",
		},
	)
}

// Instance will be attached to default external network
//
// The property Type is required.
type BaremetalServerNewParamsInterfaceExternal struct {
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Applicable only to bare metal. Each group is added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Specify `ipv4`, `ipv6`, or `dual` to enable both.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// A public IP address will be assigned to the instance.
	//
	// This field can be elided, and will marshal its zero value as "external".
	Type constant.External `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerNewParamsInterfaceExternal) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r BaremetalServerNewParamsInterfaceExternal) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParamsInterfaceExternal
	return param.MarshalObject(r, (*shadow)(&r))
}

// The instance will get an IP address from the selected network. If you choose to
// add a floating IP, the instance will be reachable from the internet. Otherwise,
// it will only have a private IP within the network.
//
// The properties NetworkID, SubnetID, Type are required.
type BaremetalServerNewParamsInterfaceSubnet struct {
	// The network where the instance will be connected.
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// The instance will get an IP address from this subnet.
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Applicable only to bare metal. Each group is added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP BaremetalServerNewParamsInterfaceSubnetFloatingIPUnion `json:"floating_ip,omitzero"`
	// The instance will get an IP address from the selected network. If you choose to
	// add a floating IP, the instance will be reachable from the internet. Otherwise,
	// it will only have a private IP within the network.
	//
	// This field can be elided, and will marshal its zero value as "subnet".
	Type constant.Subnet `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerNewParamsInterfaceSubnet) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r BaremetalServerNewParamsInterfaceSubnet) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParamsInterfaceSubnet
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BaremetalServerNewParamsInterfaceSubnetFloatingIPUnion struct {
	OfNew      *BaremetalServerNewParamsInterfaceSubnetFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *BaremetalServerNewParamsInterfaceSubnetFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u BaremetalServerNewParamsInterfaceSubnetFloatingIPUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u BaremetalServerNewParamsInterfaceSubnetFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[BaremetalServerNewParamsInterfaceSubnetFloatingIPUnion](u.OfNew, u.OfExisting)
}

func (u *BaremetalServerNewParamsInterfaceSubnetFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalServerNewParamsInterfaceSubnetFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalServerNewParamsInterfaceSubnetFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[BaremetalServerNewParamsInterfaceSubnetFloatingIPUnion](
		"source",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BaremetalServerNewParamsInterfaceSubnetFloatingIPNew{}),
			DiscriminatorValue: "new",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BaremetalServerNewParamsInterfaceSubnetFloatingIPExisting{}),
			DiscriminatorValue: "existing",
		},
	)
}

func NewBaremetalServerNewParamsInterfaceSubnetFloatingIPNew() BaremetalServerNewParamsInterfaceSubnetFloatingIPNew {
	return BaremetalServerNewParamsInterfaceSubnetFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewBaremetalServerNewParamsInterfaceSubnetFloatingIPNew].
type BaremetalServerNewParamsInterfaceSubnetFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerNewParamsInterfaceSubnetFloatingIPNew) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r BaremetalServerNewParamsInterfaceSubnetFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParamsInterfaceSubnetFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties ExistingFloatingID, Source are required.
type BaremetalServerNewParamsInterfaceSubnetFloatingIPExisting struct {
	// An existing available floating IP id must be specified if the source is set to
	// `existing`
	ExistingFloatingID string `json:"existing_floating_id,required" format:"uuid4"`
	// An existing available floating IP will be attached to the instance. A floating
	// IP is a public IP that makes the instance accessible from the internet, even if
	// it only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Source constant.Existing `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerNewParamsInterfaceSubnetFloatingIPExisting) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r BaremetalServerNewParamsInterfaceSubnetFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParamsInterfaceSubnetFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties NetworkID, Type are required.
type BaremetalServerNewParamsInterfaceAnySubnet struct {
	// The network where the instance will be connected.
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// You can specify a specific IP address from your subnet.
	IPAddress param.Opt[string] `json:"ip_address,omitzero" format:"ipvanyaddress"`
	// Applicable only to bare metal. Each group is added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP BaremetalServerNewParamsInterfaceAnySubnetFloatingIPUnion `json:"floating_ip,omitzero"`
	// Specify `ipv4`, `ipv6`, or `dual` to enable both.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Instance will be attached to a subnet with the largest count of free IPs.
	//
	// This field can be elided, and will marshal its zero value as "any_subnet".
	Type constant.AnySubnet `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerNewParamsInterfaceAnySubnet) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r BaremetalServerNewParamsInterfaceAnySubnet) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParamsInterfaceAnySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BaremetalServerNewParamsInterfaceAnySubnetFloatingIPUnion struct {
	OfNew      *BaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *BaremetalServerNewParamsInterfaceAnySubnetFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u BaremetalServerNewParamsInterfaceAnySubnetFloatingIPUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u BaremetalServerNewParamsInterfaceAnySubnetFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[BaremetalServerNewParamsInterfaceAnySubnetFloatingIPUnion](u.OfNew, u.OfExisting)
}

func (u *BaremetalServerNewParamsInterfaceAnySubnetFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalServerNewParamsInterfaceAnySubnetFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalServerNewParamsInterfaceAnySubnetFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[BaremetalServerNewParamsInterfaceAnySubnetFloatingIPUnion](
		"source",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew{}),
			DiscriminatorValue: "new",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BaremetalServerNewParamsInterfaceAnySubnetFloatingIPExisting{}),
			DiscriminatorValue: "existing",
		},
	)
}

func NewBaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew() BaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew {
	return BaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewBaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew].
type BaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r BaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties ExistingFloatingID, Source are required.
type BaremetalServerNewParamsInterfaceAnySubnetFloatingIPExisting struct {
	// An existing available floating IP id must be specified if the source is set to
	// `existing`
	ExistingFloatingID string `json:"existing_floating_id,required" format:"uuid4"`
	// An existing available floating IP will be attached to the instance. A floating
	// IP is a public IP that makes the instance accessible from the internet, even if
	// it only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Source constant.Existing `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerNewParamsInterfaceAnySubnetFloatingIPExisting) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r BaremetalServerNewParamsInterfaceAnySubnetFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParamsInterfaceAnySubnetFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties PortID, Type are required.
type BaremetalServerNewParamsInterfaceReservedFixedIP struct {
	// Network ID the subnet belongs to. Port will be plugged in this network.
	PortID string `json:"port_id,required"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Applicable only to bare metal. Each group is added to a separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion `json:"floating_ip,omitzero"`
	// An existing available reserved fixed IP will be attached to the instance. If the
	// reserved IP is not public and you choose to add a floating IP, the instance will
	// be accessible from the internet.
	//
	// This field can be elided, and will marshal its zero value as
	// "reserved_fixed_ip".
	Type constant.ReservedFixedIP `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerNewParamsInterfaceReservedFixedIP) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r BaremetalServerNewParamsInterfaceReservedFixedIP) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParamsInterfaceReservedFixedIP
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion struct {
	OfNew      *BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion](u.OfNew, u.OfExisting)
}

func (u *BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion](
		"source",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew{}),
			DiscriminatorValue: "new",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting{}),
			DiscriminatorValue: "existing",
		},
	)
}

func NewBaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew() BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew {
	return BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewBaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew].
type BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}

// The properties ExistingFloatingID, Source are required.
type BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting struct {
	// An existing available floating IP id must be specified if the source is set to
	// `existing`
	ExistingFloatingID string `json:"existing_floating_id,required" format:"uuid4"`
	// An existing available floating IP will be attached to the instance. A floating
	// IP is a public IP that makes the instance accessible from the internet, even if
	// it only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Source constant.Existing `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}

// Enable advanced DDoS protection for the server
//
// The property ProfileTemplate is required.
type BaremetalServerNewParamsDDOSProfile struct {
	// DDoS profile template ID
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// DDoS profile parameters
	Fields []BaremetalServerNewParamsDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerNewParamsDDOSProfile) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r BaremetalServerNewParamsDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParamsDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}

type BaremetalServerNewParamsDDOSProfileField struct {
	// ID of DDoS profile field
	BaseField param.Opt[int64] `json:"base_field,omitzero"`
	// Name of DDoS profile field
	FieldName param.Opt[string] `json:"field_name,omitzero"`
	// Basic type value. Only one of 'value' or 'field_value' must be specified.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value. Only one of 'value' or 'field_value' must be specified.
	FieldValue BaremetalServerNewParamsDDOSProfileFieldFieldValueUnion `json:"field_value,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerNewParamsDDOSProfileField) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r BaremetalServerNewParamsDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParamsDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BaremetalServerNewParamsDDOSProfileFieldFieldValueUnion struct {
	OfAnyArray []any             `json:",omitzero,inline"`
	OfInt      param.Opt[int64]  `json:",omitzero,inline"`
	OfString   param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u BaremetalServerNewParamsDDOSProfileFieldFieldValueUnion) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u BaremetalServerNewParamsDDOSProfileFieldFieldValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[BaremetalServerNewParamsDDOSProfileFieldFieldValueUnion](u.OfAnyArray, u.OfInt, u.OfString)
}

func (u *BaremetalServerNewParamsDDOSProfileFieldFieldValueUnion) asAny() any {
	if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type BaremetalServerListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Filters the instances by a date and time stamp when the instances last changed.
	ChangesBefore param.Opt[time.Time] `query:"changes-before,omitzero" format:"date-time" json:"-"`
	// Filters the instances by a date and time stamp when the instances last changed
	// status.
	ChangesSince param.Opt[time.Time] `query:"changes-since,omitzero" format:"date-time" json:"-"`
	// Filter out instances by flavor_id. Flavor id must match exactly.
	FlavorID param.Opt[string] `query:"flavor_id,omitzero" json:"-"`
	// Filter out instances by flavor_prefix.
	FlavorPrefix param.Opt[string] `query:"flavor_prefix,omitzero" json:"-"`
	// Include managed k8s worker nodes
	IncludeK8s param.Opt[bool] `query:"include_k8s,omitzero" json:"-"`
	// An IPv4 address to filter results by. Note: partial matches are allowed. For
	// example, searching for 192.168.0.1 will return 192.168.0.1, 192.168.0.10,
	// 192.168.0.110, and so on.
	IP param.Opt[string] `query:"ip,omitzero" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter instances by name. You can provide a full or partial name, instances with
	// matching names will be returned. For example, entering 'test' will return all
	// instances that contain 'test' in their name.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Include only isolated instances
	OnlyIsolated param.Opt[bool] `query:"only_isolated,omitzero" json:"-"`
	// Return bare metals only with external fixed IP addresses.
	OnlyWithFixedExternalIP param.Opt[bool] `query:"only_with_fixed_external_ip,omitzero" json:"-"`
	// Filter result by ddos protection profile name. Effective only with with_ddos set
	// to true.
	ProfileName param.Opt[string] `query:"profile_name,omitzero" json:"-"`
	// Optional. Filter by tag key-value pairs. curl -G --data-urlencode
	// "tag_key_value={"key": "value"}" --url
	// "https://example.com/cloud/v1/resource/1/1"
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// Filter the server list result by the UUID of the server. Allowed UUID part
	Uuid param.Opt[string] `query:"uuid,omitzero" json:"-"`
	// Include DDoS profile information for bare-metal servers in the response when set
	// to `true`. Otherwise, the `ddos_profile` field in the response is `null` by
	// default.
	WithDDOS param.Opt[bool] `query:"with_ddos,omitzero" json:"-"`
	// Include `interface_name` in the addresses
	WithInterfacesName param.Opt[bool] `query:"with_interfaces_name,omitzero" json:"-"`
	// Order by field and direction.
	//
	// Any of "created.asc", "created.desc", "name.asc", "name.desc", "status.asc",
	// "status.desc".
	OrderBy BaremetalServerListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Filter result by DDoS protection_status. Effective only with with_ddos set to
	// true. (Active, Queued or Error)
	//
	// Any of "Active", "Queued", "Error".
	ProtectionStatus BaremetalServerListParamsProtectionStatus `query:"protection_status,omitzero" json:"-"`
	// Filters instances by a server status, as a string.
	//
	// Any of "ACTIVE", "BUILD", "ERROR", "HARD_REBOOT", "REBOOT", "REBUILD", "RESCUE",
	// "SHUTOFF", "SUSPENDED".
	Status BaremetalServerListParamsStatus `query:"status,omitzero" json:"-"`
	// Optional. Filter by tag values. ?tag_value=value1&tag_value=value2
	TagValue []string `query:"tag_value,omitzero" json:"-"`
	// Return bare metals either only with advanced or only basic DDoS protection.
	// Effective only with with_ddos set to true. (advanced or basic)
	//
	// Any of "basic", "advanced".
	TypeDDOSProfile BaremetalServerListParamsTypeDDOSProfile `query:"type_ddos_profile,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [BaremetalServerListParams]'s query parameters as
// `url.Values`.
func (r BaremetalServerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Order by field and direction.
type BaremetalServerListParamsOrderBy string

const (
	BaremetalServerListParamsOrderByCreatedAsc  BaremetalServerListParamsOrderBy = "created.asc"
	BaremetalServerListParamsOrderByCreatedDesc BaremetalServerListParamsOrderBy = "created.desc"
	BaremetalServerListParamsOrderByNameAsc     BaremetalServerListParamsOrderBy = "name.asc"
	BaremetalServerListParamsOrderByNameDesc    BaremetalServerListParamsOrderBy = "name.desc"
	BaremetalServerListParamsOrderByStatusAsc   BaremetalServerListParamsOrderBy = "status.asc"
	BaremetalServerListParamsOrderByStatusDesc  BaremetalServerListParamsOrderBy = "status.desc"
)

// Filter result by DDoS protection_status. Effective only with with_ddos set to
// true. (Active, Queued or Error)
type BaremetalServerListParamsProtectionStatus string

const (
	BaremetalServerListParamsProtectionStatusActive BaremetalServerListParamsProtectionStatus = "Active"
	BaremetalServerListParamsProtectionStatusQueued BaremetalServerListParamsProtectionStatus = "Queued"
	BaremetalServerListParamsProtectionStatusError  BaremetalServerListParamsProtectionStatus = "Error"
)

// Filters instances by a server status, as a string.
type BaremetalServerListParamsStatus string

const (
	BaremetalServerListParamsStatusActive     BaremetalServerListParamsStatus = "ACTIVE"
	BaremetalServerListParamsStatusBuild      BaremetalServerListParamsStatus = "BUILD"
	BaremetalServerListParamsStatusError      BaremetalServerListParamsStatus = "ERROR"
	BaremetalServerListParamsStatusHardReboot BaremetalServerListParamsStatus = "HARD_REBOOT"
	BaremetalServerListParamsStatusReboot     BaremetalServerListParamsStatus = "REBOOT"
	BaremetalServerListParamsStatusRebuild    BaremetalServerListParamsStatus = "REBUILD"
	BaremetalServerListParamsStatusRescue     BaremetalServerListParamsStatus = "RESCUE"
	BaremetalServerListParamsStatusShutoff    BaremetalServerListParamsStatus = "SHUTOFF"
	BaremetalServerListParamsStatusSuspended  BaremetalServerListParamsStatus = "SUSPENDED"
)

// Return bare metals either only with advanced or only basic DDoS protection.
// Effective only with with_ddos set to true. (advanced or basic)
type BaremetalServerListParamsTypeDDOSProfile string

const (
	BaremetalServerListParamsTypeDDOSProfileBasic    BaremetalServerListParamsTypeDDOSProfile = "basic"
	BaremetalServerListParamsTypeDDOSProfileAdvanced BaremetalServerListParamsTypeDDOSProfile = "advanced"
)

type BaremetalServerRebuildParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Image ID
	ImageID param.Opt[string] `json:"image_id,omitzero"`
	// String in base64 format. Must not be passed together with 'username' or
	// 'password'. Examples of the user_data:
	// https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Opt[string] `json:"user_data,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalServerRebuildParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r BaremetalServerRebuildParams) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerRebuildParams
	return param.MarshalObject(r, (*shadow)(&r))
}
