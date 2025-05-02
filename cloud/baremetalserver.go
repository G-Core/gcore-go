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

// '#/components/schemas/BareMetalServerSerializer'
// "$.components.schemas.BareMetalServerSerializer"
type BaremetalServer struct {
	// '#/components/schemas/BareMetalServerSerializer/properties/addresses'
	// "$.components.schemas.BareMetalServerSerializer.properties.addresses"
	Addresses map[string][]BaremetalServerAddressUnion `json:"addresses,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/blackhole_ports'
	// "$.components.schemas.BareMetalServerSerializer.properties.blackhole_ports"
	BlackholePorts []BaremetalServerBlackholePort `json:"blackhole_ports,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.BareMetalServerSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/ddos_profile/anyOf/0'
	// "$.components.schemas.BareMetalServerSerializer.properties.ddos_profile.anyOf[0]"
	DDOSProfile DDOSProfile `json:"ddos_profile,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/fixed_ip_assignments'
	// "$.components.schemas.BareMetalServerSerializer.properties.fixed_ip_assignments"
	FixedIPAssignments []BaremetalServerFixedIPAssignment `json:"fixed_ip_assignments,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/flavor'
	// "$.components.schemas.BareMetalServerSerializer.properties.flavor"
	Flavor BaremetalServerFlavor `json:"flavor,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/instance_created'
	// "$.components.schemas.BareMetalServerSerializer.properties.instance_created"
	InstanceCreated time.Time `json:"instance_created,required" format:"date-time"`
	// '#/components/schemas/BareMetalServerSerializer/properties/instance_description/anyOf/0'
	// "$.components.schemas.BareMetalServerSerializer.properties.instance_description.anyOf[0]"
	InstanceDescription string `json:"instance_description,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/instance_id'
	// "$.components.schemas.BareMetalServerSerializer.properties.instance_id"
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	// '#/components/schemas/BareMetalServerSerializer/properties/instance_isolation/anyOf/0'
	// "$.components.schemas.BareMetalServerSerializer.properties.instance_isolation.anyOf[0]"
	InstanceIsolation BaremetalServerInstanceIsolation `json:"instance_isolation,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/instance_name'
	// "$.components.schemas.BareMetalServerSerializer.properties.instance_name"
	InstanceName string `json:"instance_name,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/project_id'
	// "$.components.schemas.BareMetalServerSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/region'
	// "$.components.schemas.BareMetalServerSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/region_id'
	// "$.components.schemas.BareMetalServerSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/ssh_key_name/anyOf/0'
	// "$.components.schemas.BareMetalServerSerializer.properties.ssh_key_name.anyOf[0]"
	SSHKeyName string `json:"ssh_key_name,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/status'
	// "$.components.schemas.BareMetalServerSerializer.properties.status"
	//
	// Any of "ACTIVE", "BUILD", "DELETED", "ERROR", "HARD_REBOOT", "MIGRATING",
	// "PASSWORD", "PAUSED", "REBOOT", "REBUILD", "RESCUE", "RESIZE", "REVERT_RESIZE",
	// "SHELVED", "SHELVED_OFFLOADED", "SHUTOFF", "SOFT_DELETED", "SUSPENDED",
	// "UNKNOWN", "VERIFY_RESIZE".
	Status BaremetalServerStatus `json:"status,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/tags'
	// "$.components.schemas.BareMetalServerSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.BareMetalServerSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/task_state/anyOf/0'
	// "$.components.schemas.BareMetalServerSerializer.properties.task_state.anyOf[0]"
	TaskState string `json:"task_state,required"`
	// '#/components/schemas/BareMetalServerSerializer/properties/vm_state'
	// "$.components.schemas.BareMetalServerSerializer.properties.vm_state"
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
// [BaremetalServerAddressFloatingIPAddress],
// [BaremetalServerAddressFixedIPAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BaremetalServerAddressUnion struct {
	Addr string `json:"addr"`
	Type string `json:"type"`
	// This field is from variant [BaremetalServerAddressFixedIPAddress].
	InterfaceName string `json:"interface_name"`
	// This field is from variant [BaremetalServerAddressFixedIPAddress].
	SubnetID string `json:"subnet_id"`
	// This field is from variant [BaremetalServerAddressFixedIPAddress].
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

func (u BaremetalServerAddressUnion) AsFloatingIPAddress() (v BaremetalServerAddressFloatingIPAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BaremetalServerAddressUnion) AsFixedIPAddress() (v BaremetalServerAddressFixedIPAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BaremetalServerAddressUnion) RawJSON() string { return u.JSON.raw }

func (r *BaremetalServerAddressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/BareMetalServerSerializer/properties/addresses/additionalProperties/items/anyOf/0'
// "$.components.schemas.BareMetalServerSerializer.properties.addresses.additionalProperties.items.anyOf[0]"
type BaremetalServerAddressFloatingIPAddress struct {
	// '#/components/schemas/BareMetalFloatingAddressSerializer/properties/addr'
	// "$.components.schemas.BareMetalFloatingAddressSerializer.properties.addr"
	Addr string `json:"addr,required"`
	// '#/components/schemas/BareMetalFloatingAddressSerializer/properties/type'
	// "$.components.schemas.BareMetalFloatingAddressSerializer.properties.type"
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
func (r BaremetalServerAddressFloatingIPAddress) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerAddressFloatingIPAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/BareMetalServerSerializer/properties/addresses/additionalProperties/items/anyOf/1'
// "$.components.schemas.BareMetalServerSerializer.properties.addresses.additionalProperties.items.anyOf[1]"
type BaremetalServerAddressFixedIPAddress struct {
	// '#/components/schemas/BareMetalFixedAddressSerializer/properties/addr'
	// "$.components.schemas.BareMetalFixedAddressSerializer.properties.addr"
	Addr string `json:"addr,required"`
	// '#/components/schemas/BareMetalFixedAddressSerializer/properties/interface_name/anyOf/0'
	// "$.components.schemas.BareMetalFixedAddressSerializer.properties.interface_name.anyOf[0]"
	InterfaceName string `json:"interface_name,required"`
	// '#/components/schemas/BareMetalFixedAddressSerializer/properties/subnet_id'
	// "$.components.schemas.BareMetalFixedAddressSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required"`
	// '#/components/schemas/BareMetalFixedAddressSerializer/properties/subnet_name'
	// "$.components.schemas.BareMetalFixedAddressSerializer.properties.subnet_name"
	SubnetName string `json:"subnet_name,required"`
	// '#/components/schemas/BareMetalFixedAddressSerializer/properties/type'
	// "$.components.schemas.BareMetalFixedAddressSerializer.properties.type"
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
func (r BaremetalServerAddressFixedIPAddress) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerAddressFixedIPAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/BareMetalServerSerializer/properties/blackhole_ports/items'
// "$.components.schemas.BareMetalServerSerializer.properties.blackhole_ports.items"
type BaremetalServerBlackholePort struct {
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
func (r BaremetalServerBlackholePort) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerBlackholePort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/BareMetalServerSerializer/properties/fixed_ip_assignments/items'
// "$.components.schemas.BareMetalServerSerializer.properties.fixed_ip_assignments.items"
type BaremetalServerFixedIPAssignment struct {
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
func (r BaremetalServerFixedIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerFixedIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/BareMetalServerSerializer/properties/flavor'
// "$.components.schemas.BareMetalServerSerializer.properties.flavor"
type BaremetalServerFlavor struct {
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
	HardwareDescription BaremetalServerFlavorHardwareDescription `json:"hardware_description,required"`
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
func (r BaremetalServerFlavor) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/BareMetalFlavorSerializer/properties/hardware_description'
// "$.components.schemas.BareMetalFlavorSerializer.properties.hardware_description"
type BaremetalServerFlavorHardwareDescription struct {
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
func (r BaremetalServerFlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerFlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/BareMetalServerSerializer/properties/instance_isolation/anyOf/0'
// "$.components.schemas.BareMetalServerSerializer.properties.instance_isolation.anyOf[0]"
type BaremetalServerInstanceIsolation struct {
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
func (r BaremetalServerInstanceIsolation) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerInstanceIsolation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/BareMetalServerSerializer/properties/status'
// "$.components.schemas.BareMetalServerSerializer.properties.status"
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

// '#/components/schemas/BareMetalServerSerializer/properties/vm_state'
// "$.components.schemas.BareMetalServerSerializer.properties.vm_state"
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
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateBareMetalServerSerializer/properties/flavor'
	// "$.components.schemas.CreateBareMetalServerSerializer.properties.flavor"
	Flavor string `json:"flavor,required"`
	// '#/components/schemas/CreateBareMetalServerSerializer/properties/interfaces'
	// "$.components.schemas.CreateBareMetalServerSerializer.properties.interfaces"
	Interfaces []BaremetalServerNewParamsInterfaceUnion `json:"interfaces,omitzero,required"`
	// '#/components/schemas/CreateBareMetalServerSerializer/properties/ssh_key_name/anyOf/0'
	// "$.components.schemas.CreateBareMetalServerSerializer.properties.ssh_key_name.anyOf[0]"
	SSHKeyName param.Opt[string] `json:"ssh_key_name,omitzero"`
	// '#/components/schemas/CreateBareMetalServerSerializer/properties/apptemplate_id'
	// "$.components.schemas.CreateBareMetalServerSerializer.properties.apptemplate_id"
	ApptemplateID param.Opt[string] `json:"apptemplate_id,omitzero"`
	// '#/components/schemas/CreateBareMetalServerSerializer/properties/image_id'
	// "$.components.schemas.CreateBareMetalServerSerializer.properties.image_id"
	ImageID param.Opt[string] `json:"image_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateBareMetalServerSerializer/properties/password'
	// "$.components.schemas.CreateBareMetalServerSerializer.properties.password"
	Password param.Opt[string] `json:"password,omitzero"`
	// '#/components/schemas/CreateBareMetalServerSerializer/properties/user_data'
	// "$.components.schemas.CreateBareMetalServerSerializer.properties.user_data"
	UserData param.Opt[string] `json:"user_data,omitzero"`
	// '#/components/schemas/CreateBareMetalServerSerializer/properties/username'
	// "$.components.schemas.CreateBareMetalServerSerializer.properties.username"
	Username param.Opt[string] `json:"username,omitzero"`
	// '#/components/schemas/CreateBareMetalServerSerializer/properties/app_config/anyOf/0'
	// "$.components.schemas.CreateBareMetalServerSerializer.properties.app_config.anyOf[0]"
	AppConfig any `json:"app_config,omitzero"`
	// '#/components/schemas/CreateBareMetalServerSerializer/properties/ddos_profile'
	// "$.components.schemas.CreateBareMetalServerSerializer.properties.ddos_profile"
	DDOSProfile BaremetalServerNewParamsDDOSProfile `json:"ddos_profile,omitzero"`
	// '#/components/schemas/CreateBareMetalServerSerializer/properties/name_templates'
	// "$.components.schemas.CreateBareMetalServerSerializer.properties.name_templates"
	NameTemplates []string `json:"name_templates,omitzero"`
	// '#/components/schemas/CreateBareMetalServerSerializer/properties/names'
	// "$.components.schemas.CreateBareMetalServerSerializer.properties.names"
	Names []string `json:"names,omitzero"`
	// '#/components/schemas/CreateBareMetalServerSerializer/properties/tags'
	// "$.components.schemas.CreateBareMetalServerSerializer.properties.tags"
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
	// '#/components/schemas/CreateBareMetalInterfaceSerializers/anyOf/0'
	// "$.components.schemas.CreateBareMetalInterfaceSerializers.anyOf[0]"
	OfExternal *BaremetalServerNewParamsInterfaceExternal `json:",omitzero,inline"`
	// '#/components/schemas/CreateBareMetalInterfaceSerializers/anyOf/1'
	// "$.components.schemas.CreateBareMetalInterfaceSerializers.anyOf[1]"
	OfSubnet *BaremetalServerNewParamsInterfaceSubnet `json:",omitzero,inline"`
	// '#/components/schemas/CreateBareMetalInterfaceSerializers/anyOf/2'
	// "$.components.schemas.CreateBareMetalInterfaceSerializers.anyOf[2]"
	OfAnySubnet *BaremetalServerNewParamsInterfaceAnySubnet `json:",omitzero,inline"`
	// '#/components/schemas/CreateBareMetalInterfaceSerializers/anyOf/3'
	// "$.components.schemas.CreateBareMetalInterfaceSerializers.anyOf[3]"
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

// '#/components/schemas/CreateBareMetalInterfaceSerializers/anyOf/0'
// "$.components.schemas.CreateBareMetalInterfaceSerializers.anyOf[0]"
//
// The property Type is required.
type BaremetalServerNewParamsInterfaceExternal struct {
	// '#/components/schemas/CreateBareMetalExternalInterfaceSerializer/properties/interface_name'
	// "$.components.schemas.CreateBareMetalExternalInterfaceSerializer.properties.interface_name"
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// '#/components/schemas/CreateBareMetalExternalInterfaceSerializer/properties/port_group'
	// "$.components.schemas.CreateBareMetalExternalInterfaceSerializer.properties.port_group"
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// '#/components/schemas/CreateBareMetalExternalInterfaceSerializer/properties/ip_family/anyOf/0'
	// "$.components.schemas.CreateBareMetalExternalInterfaceSerializer.properties.ip_family.anyOf[0]"
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// '#/components/schemas/CreateBareMetalExternalInterfaceSerializer/properties/type'
	// "$.components.schemas.CreateBareMetalExternalInterfaceSerializer.properties.type"
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

// '#/components/schemas/CreateBareMetalInterfaceSerializers/anyOf/1'
// "$.components.schemas.CreateBareMetalInterfaceSerializers.anyOf[1]"
//
// The properties NetworkID, SubnetID, Type are required.
type BaremetalServerNewParamsInterfaceSubnet struct {
	// '#/components/schemas/CreateBareMetalSubnetInterfaceSerializer/properties/network_id'
	// "$.components.schemas.CreateBareMetalSubnetInterfaceSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/CreateBareMetalSubnetInterfaceSerializer/properties/subnet_id'
	// "$.components.schemas.CreateBareMetalSubnetInterfaceSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// '#/components/schemas/CreateBareMetalSubnetInterfaceSerializer/properties/interface_name'
	// "$.components.schemas.CreateBareMetalSubnetInterfaceSerializer.properties.interface_name"
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// '#/components/schemas/CreateBareMetalSubnetInterfaceSerializer/properties/port_group'
	// "$.components.schemas.CreateBareMetalSubnetInterfaceSerializer.properties.port_group"
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// '#/components/schemas/CreateBareMetalSubnetInterfaceSerializer/properties/floating_ip'
	// "$.components.schemas.CreateBareMetalSubnetInterfaceSerializer.properties.floating_ip"
	FloatingIP BaremetalServerNewParamsInterfaceSubnetFloatingIPUnion `json:"floating_ip,omitzero"`
	// '#/components/schemas/CreateBareMetalSubnetInterfaceSerializer/properties/type'
	// "$.components.schemas.CreateBareMetalSubnetInterfaceSerializer.properties.type"
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
	// '#/components/schemas/CreateBareMetalSubnetInterfaceSerializer/properties/floating_ip/anyOf/0'
	// "$.components.schemas.CreateBareMetalSubnetInterfaceSerializer.properties.floating_ip.anyOf[0]"
	OfNew *BaremetalServerNewParamsInterfaceSubnetFloatingIPNew `json:",omitzero,inline"`
	// '#/components/schemas/CreateBareMetalSubnetInterfaceSerializer/properties/floating_ip/anyOf/1'
	// "$.components.schemas.CreateBareMetalSubnetInterfaceSerializer.properties.floating_ip.anyOf[1]"
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

// '#/components/schemas/CreateBareMetalSubnetInterfaceSerializer/properties/floating_ip/anyOf/0'
// "$.components.schemas.CreateBareMetalSubnetInterfaceSerializer.properties.floating_ip.anyOf[0]"
//
// This struct has a constant value, construct it with
// [NewBaremetalServerNewParamsInterfaceSubnetFloatingIPNew].
type BaremetalServerNewParamsInterfaceSubnetFloatingIPNew struct {
	// '#/components/schemas/NewInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.NewInstanceFloatingIpInterfaceSerializer.properties.source"
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

// '#/components/schemas/CreateBareMetalSubnetInterfaceSerializer/properties/floating_ip/anyOf/1'
// "$.components.schemas.CreateBareMetalSubnetInterfaceSerializer.properties.floating_ip.anyOf[1]"
//
// The properties ExistingFloatingID, Source are required.
type BaremetalServerNewParamsInterfaceSubnetFloatingIPExisting struct {
	// '#/components/schemas/ExistingInstanceFloatingIpInterfaceSerializer/properties/existing_floating_id'
	// "$.components.schemas.ExistingInstanceFloatingIpInterfaceSerializer.properties.existing_floating_id"
	ExistingFloatingID string `json:"existing_floating_id,required" format:"uuid4"`
	// '#/components/schemas/ExistingInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.ExistingInstanceFloatingIpInterfaceSerializer.properties.source"
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

// '#/components/schemas/CreateBareMetalInterfaceSerializers/anyOf/2'
// "$.components.schemas.CreateBareMetalInterfaceSerializers.anyOf[2]"
//
// The properties NetworkID, Type are required.
type BaremetalServerNewParamsInterfaceAnySubnet struct {
	// '#/components/schemas/CreateBareMetalAnySubnetInterfaceSerializer/properties/network_id'
	// "$.components.schemas.CreateBareMetalAnySubnetInterfaceSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/CreateBareMetalAnySubnetInterfaceSerializer/properties/interface_name'
	// "$.components.schemas.CreateBareMetalAnySubnetInterfaceSerializer.properties.interface_name"
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// '#/components/schemas/CreateBareMetalAnySubnetInterfaceSerializer/properties/ip_address'
	// "$.components.schemas.CreateBareMetalAnySubnetInterfaceSerializer.properties.ip_address"
	IPAddress param.Opt[string] `json:"ip_address,omitzero" format:"ipvanyaddress"`
	// '#/components/schemas/CreateBareMetalAnySubnetInterfaceSerializer/properties/port_group'
	// "$.components.schemas.CreateBareMetalAnySubnetInterfaceSerializer.properties.port_group"
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// '#/components/schemas/CreateBareMetalAnySubnetInterfaceSerializer/properties/floating_ip'
	// "$.components.schemas.CreateBareMetalAnySubnetInterfaceSerializer.properties.floating_ip"
	FloatingIP BaremetalServerNewParamsInterfaceAnySubnetFloatingIPUnion `json:"floating_ip,omitzero"`
	// '#/components/schemas/CreateBareMetalAnySubnetInterfaceSerializer/properties/ip_family/anyOf/0'
	// "$.components.schemas.CreateBareMetalAnySubnetInterfaceSerializer.properties.ip_family.anyOf[0]"
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// '#/components/schemas/CreateBareMetalAnySubnetInterfaceSerializer/properties/type'
	// "$.components.schemas.CreateBareMetalAnySubnetInterfaceSerializer.properties.type"
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
	// '#/components/schemas/CreateBareMetalAnySubnetInterfaceSerializer/properties/floating_ip/anyOf/0'
	// "$.components.schemas.CreateBareMetalAnySubnetInterfaceSerializer.properties.floating_ip.anyOf[0]"
	OfNew *BaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew `json:",omitzero,inline"`
	// '#/components/schemas/CreateBareMetalAnySubnetInterfaceSerializer/properties/floating_ip/anyOf/1'
	// "$.components.schemas.CreateBareMetalAnySubnetInterfaceSerializer.properties.floating_ip.anyOf[1]"
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

// '#/components/schemas/CreateBareMetalAnySubnetInterfaceSerializer/properties/floating_ip/anyOf/0'
// "$.components.schemas.CreateBareMetalAnySubnetInterfaceSerializer.properties.floating_ip.anyOf[0]"
//
// This struct has a constant value, construct it with
// [NewBaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew].
type BaremetalServerNewParamsInterfaceAnySubnetFloatingIPNew struct {
	// '#/components/schemas/NewInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.NewInstanceFloatingIpInterfaceSerializer.properties.source"
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

// '#/components/schemas/CreateBareMetalAnySubnetInterfaceSerializer/properties/floating_ip/anyOf/1'
// "$.components.schemas.CreateBareMetalAnySubnetInterfaceSerializer.properties.floating_ip.anyOf[1]"
//
// The properties ExistingFloatingID, Source are required.
type BaremetalServerNewParamsInterfaceAnySubnetFloatingIPExisting struct {
	// '#/components/schemas/ExistingInstanceFloatingIpInterfaceSerializer/properties/existing_floating_id'
	// "$.components.schemas.ExistingInstanceFloatingIpInterfaceSerializer.properties.existing_floating_id"
	ExistingFloatingID string `json:"existing_floating_id,required" format:"uuid4"`
	// '#/components/schemas/ExistingInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.ExistingInstanceFloatingIpInterfaceSerializer.properties.source"
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

// '#/components/schemas/CreateBareMetalInterfaceSerializers/anyOf/3'
// "$.components.schemas.CreateBareMetalInterfaceSerializers.anyOf[3]"
//
// The properties PortID, Type are required.
type BaremetalServerNewParamsInterfaceReservedFixedIP struct {
	// '#/components/schemas/CreateBareMetalReservedFixedIpInterfaceSerializer/properties/port_id'
	// "$.components.schemas.CreateBareMetalReservedFixedIpInterfaceSerializer.properties.port_id"
	PortID string `json:"port_id,required"`
	// '#/components/schemas/CreateBareMetalReservedFixedIpInterfaceSerializer/properties/interface_name'
	// "$.components.schemas.CreateBareMetalReservedFixedIpInterfaceSerializer.properties.interface_name"
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// '#/components/schemas/CreateBareMetalReservedFixedIpInterfaceSerializer/properties/port_group'
	// "$.components.schemas.CreateBareMetalReservedFixedIpInterfaceSerializer.properties.port_group"
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// '#/components/schemas/CreateBareMetalReservedFixedIpInterfaceSerializer/properties/floating_ip'
	// "$.components.schemas.CreateBareMetalReservedFixedIpInterfaceSerializer.properties.floating_ip"
	FloatingIP BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion `json:"floating_ip,omitzero"`
	// '#/components/schemas/CreateBareMetalReservedFixedIpInterfaceSerializer/properties/type'
	// "$.components.schemas.CreateBareMetalReservedFixedIpInterfaceSerializer.properties.type"
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
	// '#/components/schemas/CreateBareMetalReservedFixedIpInterfaceSerializer/properties/floating_ip/anyOf/0'
	// "$.components.schemas.CreateBareMetalReservedFixedIpInterfaceSerializer.properties.floating_ip.anyOf[0]"
	OfNew *BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew `json:",omitzero,inline"`
	// '#/components/schemas/CreateBareMetalReservedFixedIpInterfaceSerializer/properties/floating_ip/anyOf/1'
	// "$.components.schemas.CreateBareMetalReservedFixedIpInterfaceSerializer.properties.floating_ip.anyOf[1]"
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

// '#/components/schemas/CreateBareMetalReservedFixedIpInterfaceSerializer/properties/floating_ip/anyOf/0'
// "$.components.schemas.CreateBareMetalReservedFixedIpInterfaceSerializer.properties.floating_ip.anyOf[0]"
//
// This struct has a constant value, construct it with
// [NewBaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew].
type BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPNew struct {
	// '#/components/schemas/NewInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.NewInstanceFloatingIpInterfaceSerializer.properties.source"
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

// '#/components/schemas/CreateBareMetalReservedFixedIpInterfaceSerializer/properties/floating_ip/anyOf/1'
// "$.components.schemas.CreateBareMetalReservedFixedIpInterfaceSerializer.properties.floating_ip.anyOf[1]"
//
// The properties ExistingFloatingID, Source are required.
type BaremetalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting struct {
	// '#/components/schemas/ExistingInstanceFloatingIpInterfaceSerializer/properties/existing_floating_id'
	// "$.components.schemas.ExistingInstanceFloatingIpInterfaceSerializer.properties.existing_floating_id"
	ExistingFloatingID string `json:"existing_floating_id,required" format:"uuid4"`
	// '#/components/schemas/ExistingInstanceFloatingIpInterfaceSerializer/properties/source'
	// "$.components.schemas.ExistingInstanceFloatingIpInterfaceSerializer.properties.source"
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

// '#/components/schemas/CreateBareMetalServerSerializer/properties/ddos_profile'
// "$.components.schemas.CreateBareMetalServerSerializer.properties.ddos_profile"
//
// The property ProfileTemplate is required.
type BaremetalServerNewParamsDDOSProfile struct {
	// '#/components/schemas/CreateDDoSProfileSerializer/properties/profile_template'
	// "$.components.schemas.CreateDDoSProfileSerializer.properties.profile_template"
	ProfileTemplate int64 `json:"profile_template,required"`
	// '#/components/schemas/CreateDDoSProfileSerializer/properties/profile_template_name/anyOf/0'
	// "$.components.schemas.CreateDDoSProfileSerializer.properties.profile_template_name.anyOf[0]"
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// '#/components/schemas/CreateDDoSProfileSerializer/properties/fields/anyOf/0'
	// "$.components.schemas.CreateDDoSProfileSerializer.properties.fields.anyOf[0]"
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

// '#/components/schemas/CreateDDoSProfileSerializer/properties/fields/anyOf/0/items'
// "$.components.schemas.CreateDDoSProfileSerializer.properties.fields.anyOf[0].items"
type BaremetalServerNewParamsDDOSProfileField struct {
	// '#/components/schemas/CreateBareMetalDDoSProfileFieldSerializer/properties/base_field/anyOf/0'
	// "$.components.schemas.CreateBareMetalDDoSProfileFieldSerializer.properties.base_field.anyOf[0]"
	BaseField param.Opt[int64] `json:"base_field,omitzero"`
	// '#/components/schemas/CreateBareMetalDDoSProfileFieldSerializer/properties/field_name/anyOf/0'
	// "$.components.schemas.CreateBareMetalDDoSProfileFieldSerializer.properties.field_name.anyOf[0]"
	FieldName param.Opt[string] `json:"field_name,omitzero"`
	// '#/components/schemas/CreateBareMetalDDoSProfileFieldSerializer/properties/value/anyOf/0'
	// "$.components.schemas.CreateBareMetalDDoSProfileFieldSerializer.properties.value.anyOf[0]"
	Value param.Opt[string] `json:"value,omitzero"`
	// '#/components/schemas/CreateBareMetalDDoSProfileFieldSerializer/properties/field_value'
	// "$.components.schemas.CreateBareMetalDDoSProfileFieldSerializer.properties.field_value"
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
	// '#/components/schemas/CreateBareMetalDDoSProfileFieldSerializer/properties/field_value/anyOf/0'
	// "$.components.schemas.CreateBareMetalDDoSProfileFieldSerializer.properties.field_value.anyOf[0]"
	OfAnyArray []any `json:",omitzero,inline"`
	// '#/components/schemas/CreateBareMetalDDoSProfileFieldSerializer/properties/field_value/anyOf/1'
	// "$.components.schemas.CreateBareMetalDDoSProfileFieldSerializer.properties.field_value.anyOf[1]"
	OfInt param.Opt[int64] `json:",omitzero,inline"`
	// '#/components/schemas/CreateBareMetalDDoSProfileFieldSerializer/properties/field_value/anyOf/2'
	// "$.components.schemas.CreateBareMetalDDoSProfileFieldSerializer.properties.field_value.anyOf[2]"
	OfString param.Opt[string] `json:",omitzero,inline"`
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
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[2]"
	ChangesBefore param.Opt[time.Time] `query:"changes-before,omitzero" format:"date-time" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[3]"
	ChangesSince param.Opt[time.Time] `query:"changes-since,omitzero" format:"date-time" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[4]"
	FlavorID param.Opt[string] `query:"flavor_id,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[5]"
	FlavorPrefix param.Opt[string] `query:"flavor_prefix,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/6'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[6]"
	IncludeK8s param.Opt[bool] `query:"include_k8s,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/7'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[7]"
	IP param.Opt[string] `query:"ip,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/8'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[8]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/9'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[9]"
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/10'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[10]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/11'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[11]"
	OnlyIsolated param.Opt[bool] `query:"only_isolated,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/12'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[12]"
	OnlyWithFixedExternalIP param.Opt[bool] `query:"only_with_fixed_external_ip,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/14'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[14]"
	ProfileName param.Opt[string] `query:"profile_name,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/17'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[17]"
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/20'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[20]"
	Uuid param.Opt[string] `query:"uuid,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/21'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[21]"
	WithDDOS param.Opt[bool] `query:"with_ddos,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/22'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[22]"
	WithInterfacesName param.Opt[bool] `query:"with_interfaces_name,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/13'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[13]"
	//
	// Any of "created.asc", "created.desc", "name.asc", "name.desc", "status.asc",
	// "status.desc".
	OrderBy BaremetalServerListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/15'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[15]"
	//
	// Any of "Active", "Queued", "Error".
	ProtectionStatus BaremetalServerListParamsProtectionStatus `query:"protection_status,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/16'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[16]"
	//
	// Any of "ACTIVE", "BUILD", "ERROR", "HARD_REBOOT", "REBOOT", "REBUILD", "RESCUE",
	// "SHUTOFF", "SUSPENDED".
	Status BaremetalServerListParamsStatus `query:"status,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/18'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[18]"
	TagValue []string `query:"tag_value,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/19'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[19]"
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

// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/13'
// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[13]"
type BaremetalServerListParamsOrderBy string

const (
	BaremetalServerListParamsOrderByCreatedAsc  BaremetalServerListParamsOrderBy = "created.asc"
	BaremetalServerListParamsOrderByCreatedDesc BaremetalServerListParamsOrderBy = "created.desc"
	BaremetalServerListParamsOrderByNameAsc     BaremetalServerListParamsOrderBy = "name.asc"
	BaremetalServerListParamsOrderByNameDesc    BaremetalServerListParamsOrderBy = "name.desc"
	BaremetalServerListParamsOrderByStatusAsc   BaremetalServerListParamsOrderBy = "status.asc"
	BaremetalServerListParamsOrderByStatusDesc  BaremetalServerListParamsOrderBy = "status.desc"
)

// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/15'
// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[15]"
type BaremetalServerListParamsProtectionStatus string

const (
	BaremetalServerListParamsProtectionStatusActive BaremetalServerListParamsProtectionStatus = "Active"
	BaremetalServerListParamsProtectionStatusQueued BaremetalServerListParamsProtectionStatus = "Queued"
	BaremetalServerListParamsProtectionStatusError  BaremetalServerListParamsProtectionStatus = "Error"
)

// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/16'
// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[16]"
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

// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/19'
// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}'].get.parameters[19]"
type BaremetalServerListParamsTypeDDOSProfile string

const (
	BaremetalServerListParamsTypeDDOSProfileBasic    BaremetalServerListParamsTypeDDOSProfile = "basic"
	BaremetalServerListParamsTypeDDOSProfileAdvanced BaremetalServerListParamsTypeDDOSProfile = "advanced"
)

type BaremetalServerRebuildParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bserver_id%7D%2Frebuild/post/parameters/0/schema'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}/{server_id}/rebuild'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fbminstances%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bserver_id%7D%2Frebuild/post/parameters/1/schema'
	// "$.paths['/cloud/v1/bminstances/{project_id}/{region_id}/{server_id}/rebuild'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/RebuildBaremetalSchema/properties/image_id'
	// "$.components.schemas.RebuildBaremetalSchema.properties.image_id"
	ImageID param.Opt[string] `json:"image_id,omitzero"`
	// '#/components/schemas/RebuildBaremetalSchema/properties/user_data'
	// "$.components.schemas.RebuildBaremetalSchema.properties.user_data"
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
