// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/G-Core/gcore-go/shared/constant"
)

// Bare metal servers are dedicated physical machines with direct hardware access,
// supporting provisioning, rebuilding, and network configuration within a cloud
// region.
//
// BaremetalRegionalServerService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBaremetalRegionalServerService] method instead.
type BaremetalRegionalServerService struct {
	Options []option.RequestOption
}

// NewBaremetalRegionalServerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBaremetalRegionalServerService(opts ...option.RequestOption) (r BaremetalRegionalServerService) {
	r = BaremetalRegionalServerService{}
	r.Options = opts
	return
}

// Use
// [POST /v1/projects/{`project_id`}/baremetal/servers](#operation/GlobalBaremetalListView.post)
// instead. This route continues to work. Later, we will deprecate it and then
// remove it. We will tell you in advance.
//
// Create a new bare metal server with the specified configuration.
//
// How to get access:
//
// For Linux,
//
//   - Use the `user_data` field to provide a
//     [cloud-init script](https://cloudinit.readthedocs.io/en/latest/reference/examples.html)
//     in base64 to apply configurations to the instance.
//   - Specify the `username` and `password` to create a new user.
//   - When only `password` is provided, it is set as the password for the default
//     user of the image.
//   - The `user_data` is ignored when the `password` is specified.
//
// For Windows,
//
//   - Use the `user_data` field to provide a
//     [cloudbase-init script](https://cloudbase-init.readthedocs.io/en/latest/userdata.html#cloud-config)
//     in base64 to create new users on Windows.
//   - Use the `password` field to set the password for the 'Admin' user on Windows.
//   - The password of the Admin user cannot be updated via `user_data`.
//   - The `username` cannot be specified in the request.
func (r *BaremetalRegionalServerService) New(ctx context.Context, params BaremetalRegionalServerNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/bminstances/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Use
// [PATCH /v1/projects/{`project_id`}/baremetal/servers/{`server_id`}](#operation/GlobalBaremetalDetailView.patch)
// instead. This route continues to work. Later, we will deprecate it and then
// remove it. We will tell you in advance.
//
// Rename bare metal server or update tags
func (r *BaremetalRegionalServerService) Update(ctx context.Context, serverID string, params BaremetalRegionalServerUpdateParams, opts ...option.RequestOption) (res *BaremetalRegionalServer, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if serverID == "" {
		err = errors.New("missing required server_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/bminstances/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Use
// [GET /v1/projects/{`project_id`}/baremetal/servers](#operation/GlobalBaremetalListView.get)
// instead. This route continues to work. Later, we will deprecate it and then
// remove it. We will tell you in advance.
//
// List all bare metal servers in the specified project and region. Results can be
// filtered by various parameters like name, status, and IP address.
func (r *BaremetalRegionalServerService) List(ctx context.Context, params BaremetalRegionalServerListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[BaremetalRegionalServer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
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

// Use
// [GET /v1/projects/{`project_id`}/baremetal/servers](#operation/GlobalBaremetalListView.get)
// instead. This route continues to work. Later, we will deprecate it and then
// remove it. We will tell you in advance.
//
// List all bare metal servers in the specified project and region. Results can be
// filtered by various parameters like name, status, and IP address.
func (r *BaremetalRegionalServerService) ListAutoPaging(ctx context.Context, params BaremetalRegionalServerListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[BaremetalRegionalServer] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Use
// [DELETE /v1/projects/{`project_id`}/baremetal/servers/{`server_id`}](#operation/GlobalBaremetalDetailView.delete)
// instead. This route continues to work. Later, we will deprecate it and then
// remove it. We will tell you in advance.
//
// Delete the bare metal server and its ports. A reserved fixed IP stays, unless
// `reserved_fixed_ip_ids` names it or `all_reserved_fixed_ips` is `true`. A port
// that you [create as a port](#operation/GlobalPortListView.post) is not a
// reserved fixed IP, so this delete removes it.
//
// To keep the ports, use
// [`DELETE /v1/projects/{project_id}/baremetal/servers/{server_id}`](#operation/GlobalBaremetalDetailView.delete).
// That delete keeps every port by default.
func (r *BaremetalRegionalServerService) Delete(ctx context.Context, serverID string, params BaremetalRegionalServerDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if serverID == "" {
		err = errors.New("missing required server_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/bminstances/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Use
// [POST /v1/projects/{`project_id`}/baremetal/servers/{`server_id`}/action](#operation/GlobalBaremetalActionView.post)
// instead. This route continues to work. Later, we will deprecate it and then
// remove it. We will tell you in advance.
//
// The action can be one of: start, stop, reboot or `reboot_hard`.
func (r *BaremetalRegionalServerService) Action(ctx context.Context, serverID string, params BaremetalRegionalServerActionParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if serverID == "" {
		err = errors.New("missing required server_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v2/baremetal/%v/%v/%s/action", params.ProjectID.Value, params.RegionID.Value, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Use
// [GET /v1/projects/{`project_id`}/baremetal/servers/{`server_id`}](#operation/GlobalBaremetalDetailView.get)
// instead. This route continues to work. Later, we will deprecate it and then
// remove it. We will tell you in advance.
//
// Retrieve detailed information about a specific baremetal instance. This endpoint
// always returns `ddos_profile` (if present) and therefore always requires
// `DDOS_READ`.
func (r *BaremetalRegionalServerService) Get(ctx context.Context, serverID string, query BaremetalRegionalServerGetParams, opts ...option.RequestOption) (res *BaremetalRegionalServer, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !query.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if serverID == "" {
		err = errors.New("missing required server_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/bminstances/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// IP addresses of the trunk port and its subports.
type BaremetalRegionalFixedAddress struct {
	// Address
	Addr string `json:"addr" api:"required"`
	// Interface name. This field will be `null` if `with_interfaces_name=true` is not
	// set in the request when listing servers. It will also be `null` if the
	// `interface_name` was not specified during server creation or when attaching the
	// interface.
	InterfaceName string `json:"interface_name" api:"required"`
	// The unique identifier of the subnet associated with this address.
	SubnetID string `json:"subnet_id" api:"required"`
	// The name of the subnet associated with this address.
	SubnetName string `json:"subnet_name" api:"required"`
	// Type of the address
	Type constant.Fixed `json:"type" default:"fixed"`
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
func (r BaremetalRegionalFixedAddress) RawJSON() string { return r.JSON.raw }
func (r *BaremetalRegionalFixedAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalRegionalFloatingAddress struct {
	// Address
	Addr string `json:"addr" api:"required"`
	// Type of the address
	Type constant.Floating `json:"type" default:"floating"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Addr        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalRegionalFloatingAddress) RawJSON() string { return r.JSON.raw }
func (r *BaremetalRegionalFloatingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalRegionalServer struct {
	// Bare metal server ID
	ID string `json:"id" api:"required" format:"uuid4"`
	// Map of `network_name` to list of addresses in that network
	Addresses map[string][]BaremetalRegionalServerAddressUnion `json:"addresses" api:"required"`
	// IP addresses of the instances that are blackholed by DDoS mitigation system
	BlackholePorts []BlackholePort `json:"blackhole_ports" api:"required"`
	// Datetime when bare metal server was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Task that created this entity. The value is null if the server has no creator
	// task ID.
	CreatorTaskID string `json:"creator_task_id" api:"required"`
	// Bare metal advanced DDoS protection profile. It is always `null` if query
	// parameter `with_ddos=true` is not set.
	DDOSProfile DDOSProfile `json:"ddos_profile" api:"required"`
	// Fixed IP assigned to instance
	FixedIPAssignments []BaremetalRegionalServerFixedIPAssignment `json:"fixed_ip_assignments" api:"required"`
	// Flavor details
	Flavor BaremetalRegionalServerFlavor `json:"flavor" api:"required"`
	// Instance isolation information. The value is null if the server is not isolated.
	InstanceIsolation InstanceIsolation `json:"instance_isolation" api:"required"`
	// Bare metal server name
	Name string `json:"name" api:"required"`
	// Project ID
	ProjectID int64 `json:"project_id" api:"required"`
	// Region name
	Region string `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// SSH key assigned to bare metal server. The value is null if no key pair is set.
	SSHKeyName string `json:"ssh_key_name" api:"required"`
	// Bare metal server status
	//
	// Any of "ACTIVE", "BUILD", "DELETED", "ERROR", "HARD_REBOOT", "MIGRATING",
	// "PASSWORD", "PAUSED", "REBOOT", "REBUILD", "RESCUE", "RESIZE", "REVERT_RESIZE",
	// "SHELVED", "SHELVED_OFFLOADED", "SHUTOFF", "SOFT_DELETED", "SUSPENDED",
	// "UNKNOWN", "VERIFY_RESIZE".
	Status BaremetalRegionalServerStatus `json:"status" api:"required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags" api:"required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id" api:"required"`
	// Task state. The value is null if no operation is in progress.
	TaskState string `json:"task_state" api:"required"`
	// Bare metal server state
	//
	// Any of "active", "building", "deleted", "error", "paused", "rescued", "resized",
	// "shelved", "shelved_offloaded", "soft-deleted", "stopped", "suspended".
	VmState BaremetalRegionalServerVmState `json:"vm_state" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Addresses          respjson.Field
		BlackholePorts     respjson.Field
		CreatedAt          respjson.Field
		CreatorTaskID      respjson.Field
		DDOSProfile        respjson.Field
		FixedIPAssignments respjson.Field
		Flavor             respjson.Field
		InstanceIsolation  respjson.Field
		Name               respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		SSHKeyName         respjson.Field
		Status             respjson.Field
		Tags               respjson.Field
		TaskID             respjson.Field
		TaskState          respjson.Field
		VmState            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalRegionalServer) RawJSON() string { return r.JSON.raw }
func (r *BaremetalRegionalServer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BaremetalRegionalServerAddressUnion contains all possible properties and values
// from [BaremetalRegionalFloatingAddress], [BaremetalRegionalFixedAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BaremetalRegionalServerAddressUnion struct {
	Addr string `json:"addr"`
	Type string `json:"type"`
	// This field is from variant [BaremetalRegionalFixedAddress].
	InterfaceName string `json:"interface_name"`
	// This field is from variant [BaremetalRegionalFixedAddress].
	SubnetID string `json:"subnet_id"`
	// This field is from variant [BaremetalRegionalFixedAddress].
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

func (u BaremetalRegionalServerAddressUnion) AsFloatingIPAddress() (v BaremetalRegionalFloatingAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BaremetalRegionalServerAddressUnion) AsFixedIPAddress() (v BaremetalRegionalFixedAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BaremetalRegionalServerAddressUnion) RawJSON() string { return u.JSON.raw }

func (r *BaremetalRegionalServerAddressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalRegionalServerFixedIPAssignment struct {
	// Is network external
	External bool `json:"external" api:"required"`
	// Ip address
	IPAddress string `json:"ip_address" api:"required"`
	// Interface subnet id
	SubnetID string `json:"subnet_id" api:"required"`
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
func (r BaremetalRegionalServerFixedIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *BaremetalRegionalServerFixedIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor details
type BaremetalRegionalServerFlavor struct {
	// CPU architecture
	Architecture string `json:"architecture" api:"required"`
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id" api:"required"`
	// Flavor name
	FlavorName string `json:"flavor_name" api:"required"`
	// Additional hardware description
	HardwareDescription BaremetalRegionalServerFlavorHardwareDescription `json:"hardware_description" api:"required"`
	// Operating system
	OsType string `json:"os_type" api:"required"`
	// RAM size in MiB
	Ram int64 `json:"ram" api:"required"`
	// Flavor resource class for mapping to hardware capacity
	ResourceClass string `json:"resource_class" api:"required"`
	// Virtual CPU count. For bare metal flavors, it's a physical CPU count
	Vcpus int64 `json:"vcpus" api:"required"`
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
func (r BaremetalRegionalServerFlavor) RawJSON() string { return r.JSON.raw }
func (r *BaremetalRegionalServerFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional hardware description
type BaremetalRegionalServerFlavorHardwareDescription struct {
	// Human-readable CPU description
	CPU string `json:"cpu" api:"required"`
	// Human-readable disk description
	Disk string `json:"disk" api:"required"`
	// If the flavor is licensed, this field contains the license type
	License string `json:"license" api:"required"`
	// Human-readable NIC description
	Network string `json:"network" api:"required"`
	// Human-readable RAM description
	Ram string `json:"ram" api:"required"`
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
func (r BaremetalRegionalServerFlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *BaremetalRegionalServerFlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare metal server status
type BaremetalRegionalServerStatus string

const (
	BaremetalRegionalServerStatusActive           BaremetalRegionalServerStatus = "ACTIVE"
	BaremetalRegionalServerStatusBuild            BaremetalRegionalServerStatus = "BUILD"
	BaremetalRegionalServerStatusDeleted          BaremetalRegionalServerStatus = "DELETED"
	BaremetalRegionalServerStatusError            BaremetalRegionalServerStatus = "ERROR"
	BaremetalRegionalServerStatusHardReboot       BaremetalRegionalServerStatus = "HARD_REBOOT"
	BaremetalRegionalServerStatusMigrating        BaremetalRegionalServerStatus = "MIGRATING"
	BaremetalRegionalServerStatusPassword         BaremetalRegionalServerStatus = "PASSWORD"
	BaremetalRegionalServerStatusPaused           BaremetalRegionalServerStatus = "PAUSED"
	BaremetalRegionalServerStatusReboot           BaremetalRegionalServerStatus = "REBOOT"
	BaremetalRegionalServerStatusRebuild          BaremetalRegionalServerStatus = "REBUILD"
	BaremetalRegionalServerStatusRescue           BaremetalRegionalServerStatus = "RESCUE"
	BaremetalRegionalServerStatusResize           BaremetalRegionalServerStatus = "RESIZE"
	BaremetalRegionalServerStatusRevertResize     BaremetalRegionalServerStatus = "REVERT_RESIZE"
	BaremetalRegionalServerStatusShelved          BaremetalRegionalServerStatus = "SHELVED"
	BaremetalRegionalServerStatusShelvedOffloaded BaremetalRegionalServerStatus = "SHELVED_OFFLOADED"
	BaremetalRegionalServerStatusShutoff          BaremetalRegionalServerStatus = "SHUTOFF"
	BaremetalRegionalServerStatusSoftDeleted      BaremetalRegionalServerStatus = "SOFT_DELETED"
	BaremetalRegionalServerStatusSuspended        BaremetalRegionalServerStatus = "SUSPENDED"
	BaremetalRegionalServerStatusUnknown          BaremetalRegionalServerStatus = "UNKNOWN"
	BaremetalRegionalServerStatusVerifyResize     BaremetalRegionalServerStatus = "VERIFY_RESIZE"
)

// Bare metal server state
type BaremetalRegionalServerVmState string

const (
	BaremetalRegionalServerVmStateActive           BaremetalRegionalServerVmState = "active"
	BaremetalRegionalServerVmStateBuilding         BaremetalRegionalServerVmState = "building"
	BaremetalRegionalServerVmStateDeleted          BaremetalRegionalServerVmState = "deleted"
	BaremetalRegionalServerVmStateError            BaremetalRegionalServerVmState = "error"
	BaremetalRegionalServerVmStatePaused           BaremetalRegionalServerVmState = "paused"
	BaremetalRegionalServerVmStateRescued          BaremetalRegionalServerVmState = "rescued"
	BaremetalRegionalServerVmStateResized          BaremetalRegionalServerVmState = "resized"
	BaremetalRegionalServerVmStateShelved          BaremetalRegionalServerVmState = "shelved"
	BaremetalRegionalServerVmStateShelvedOffloaded BaremetalRegionalServerVmState = "shelved_offloaded"
	BaremetalRegionalServerVmStateSoftDeleted      BaremetalRegionalServerVmState = "soft-deleted"
	BaremetalRegionalServerVmStateStopped          BaremetalRegionalServerVmState = "stopped"
	BaremetalRegionalServerVmStateSuspended        BaremetalRegionalServerVmState = "suspended"
)

type BaremetalRegionalServerNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// The flavor of the instance.
	Flavor string `json:"flavor" api:"required"`
	// A list of network interfaces for the server. You can create one or more
	// interfaces - private, public, or both.
	Interfaces []BaremetalRegionalServerNewParamsInterfaceUnion `json:"interfaces,omitzero" api:"required"`
	// Specifies the name of the SSH keypair, created via the
	// [/v1/`ssh_keys` endpoint](/docs/api-reference/cloud/ssh-keys/add-or-generate-ssh-key).
	SSHKeyName param.Opt[string] `json:"ssh_key_name,omitzero"`
	// Apptemplate ID. Either `image_id` or `apptemplate_id` is required.
	ApptemplateID param.Opt[string] `json:"apptemplate_id,omitzero"`
	// Image ID. Either `image_id` or `apptemplate_id` is required.
	ImageID param.Opt[string] `json:"image_id,omitzero" format:"uuid4"`
	// Server name.
	Name param.Opt[string] `json:"name,omitzero"`
	// If you want server names to be automatically generated based on IP addresses,
	// you can provide a name template instead of specifying the name manually. The
	// template should include a placeholder that will be replaced during provisioning.
	// Supported placeholders are: `{ip_octets}` (last 3 octets of the IP),
	// `{two_ip_octets}`, and `{one_ip_octet}`.
	NameTemplate param.Opt[string] `json:"name_template,omitzero"`
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
	// `user_data`: https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Opt[string] `json:"user_data,omitzero"`
	// For Linux instances, 'username' and 'password' are used to create a new user.
	// For Windows instances, 'username' cannot be specified. Use 'password' field to
	// set the password for the 'Admin' user on Windows.
	Username param.Opt[string] `json:"username,omitzero"`
	// Parameters for the application template if creating the instance from an
	// `apptemplate`.
	AppConfig map[string]any `json:"app_config,omitzero"`
	// Enable advanced DDoS protection for the server
	DDOSProfile BaremetalRegionalServerNewParamsDDOSProfile `json:"ddos_profile,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Tag keys and values cannot contain characters from
	// U+0080 to U+00FF. Some tags are read-only and cannot be modified by the user.
	// Tags are also integrated with cost reports, allowing cost data to be filtered
	// based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r BaremetalRegionalServerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BaremetalRegionalServerNewParamsInterfaceUnion struct {
	OfExternal        *BaremetalRegionalServerNewParamsInterfaceExternal        `json:",omitzero,inline"`
	OfSubnet          *BaremetalRegionalServerNewParamsInterfaceSubnet          `json:",omitzero,inline"`
	OfAnySubnet       *BaremetalRegionalServerNewParamsInterfaceAnySubnet       `json:",omitzero,inline"`
	OfReservedFixedIP *BaremetalRegionalServerNewParamsInterfaceReservedFixedIP `json:",omitzero,inline"`
	paramUnion
}

func (u BaremetalRegionalServerNewParamsInterfaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExternal, u.OfSubnet, u.OfAnySubnet, u.OfReservedFixedIP)
}
func (u *BaremetalRegionalServerNewParamsInterfaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BaremetalRegionalServerNewParamsInterfaceUnion) asAny() any {
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
func (u BaremetalRegionalServerNewParamsInterfaceUnion) GetSubnetID() *string {
	if vt := u.OfSubnet; vt != nil {
		return &vt.SubnetID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalRegionalServerNewParamsInterfaceUnion) GetIPAddress() *string {
	if vt := u.OfAnySubnet; vt != nil && vt.IPAddress.Valid() {
		return &vt.IPAddress.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalRegionalServerNewParamsInterfaceUnion) GetPortID() *string {
	if vt := u.OfReservedFixedIP; vt != nil {
		return &vt.PortID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalRegionalServerNewParamsInterfaceUnion) GetType() *string {
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
func (u BaremetalRegionalServerNewParamsInterfaceUnion) GetInterfaceName() *string {
	if vt := u.OfExternal; vt != nil && vt.InterfaceName.Valid() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfSubnet; vt != nil && vt.InterfaceName.Valid() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfAnySubnet; vt != nil && vt.InterfaceName.Valid() {
		return &vt.InterfaceName.Value
	} else if vt := u.OfReservedFixedIP; vt != nil && vt.InterfaceName.Valid() {
		return &vt.InterfaceName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalRegionalServerNewParamsInterfaceUnion) GetIPFamily() *string {
	if vt := u.OfExternal; vt != nil {
		return (*string)(&vt.IPFamily)
	} else if vt := u.OfAnySubnet; vt != nil {
		return (*string)(&vt.IPFamily)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalRegionalServerNewParamsInterfaceUnion) GetPortGroup() *int64 {
	if vt := u.OfExternal; vt != nil && vt.PortGroup.Valid() {
		return &vt.PortGroup.Value
	} else if vt := u.OfSubnet; vt != nil && vt.PortGroup.Valid() {
		return &vt.PortGroup.Value
	} else if vt := u.OfAnySubnet; vt != nil && vt.PortGroup.Valid() {
		return &vt.PortGroup.Value
	} else if vt := u.OfReservedFixedIP; vt != nil && vt.PortGroup.Valid() {
		return &vt.PortGroup.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalRegionalServerNewParamsInterfaceUnion) GetNetworkID() *string {
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
func (u BaremetalRegionalServerNewParamsInterfaceUnion) GetFloatingIP() (res baremetalRegionalServerNewParamsInterfaceUnionFloatingIP) {
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
// [*BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPNew],
// [*BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPExisting],
// [*BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPNew],
// [*BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPExisting],
// [*BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPNew],
// [*BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting]
type baremetalRegionalServerNewParamsInterfaceUnionFloatingIP struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cloud.BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPNew:
//	case *cloud.BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPExisting:
//	case *cloud.BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPNew:
//	case *cloud.BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPExisting:
//	case *cloud.BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPNew:
//	case *cloud.BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u baremetalRegionalServerNewParamsInterfaceUnionFloatingIP) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u baremetalRegionalServerNewParamsInterfaceUnionFloatingIP) GetSource() *string {
	switch vt := u.any.(type) {
	case *BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPUnion:
		return vt.GetSource()
	case *BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPUnion:
		return vt.GetSource()
	case *BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion:
		return vt.GetSource()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u baremetalRegionalServerNewParamsInterfaceUnionFloatingIP) GetExistingFloatingID() *string {
	switch vt := u.any.(type) {
	case *BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPUnion:
		return vt.GetExistingFloatingID()
	case *BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPUnion:
		return vt.GetExistingFloatingID()
	case *BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion:
		return vt.GetExistingFloatingID()
	}
	return nil
}

func init() {
	apijson.RegisterUnion[BaremetalRegionalServerNewParamsInterfaceUnion](
		"type",
		apijson.Discriminator[BaremetalRegionalServerNewParamsInterfaceExternal]("external"),
		apijson.Discriminator[BaremetalRegionalServerNewParamsInterfaceSubnet]("subnet"),
		apijson.Discriminator[BaremetalRegionalServerNewParamsInterfaceAnySubnet]("any_subnet"),
		apijson.Discriminator[BaremetalRegionalServerNewParamsInterfaceReservedFixedIP]("reserved_fixed_ip"),
	)
}

// Instance will be attached to default external network
//
// The property Type is required.
type BaremetalRegionalServerNewParamsInterfaceExternal struct {
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Specifies the trunk group to which this interface belongs. Applicable only for
	// bare metal servers. Each unique port group is mapped to a separate trunk port.
	// Use this to control how interfaces are grouped across trunks.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Specify `ipv4`, `ipv6`, or `dual` to enable both.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// A public IP address will be assigned to the instance.
	//
	// This field can be elided, and will marshal its zero value as "external".
	Type constant.External `json:"type" default:"external"`
	paramObj
}

func (r BaremetalRegionalServerNewParamsInterfaceExternal) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerNewParamsInterfaceExternal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerNewParamsInterfaceExternal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The instance will get an IP address from the selected network. If you choose to
// add a floating IP, the instance will be reachable from the internet. Otherwise,
// it will only have a private IP within the network.
//
// The properties NetworkID, SubnetID, Type are required.
type BaremetalRegionalServerNewParamsInterfaceSubnet struct {
	// The network where the instance will be connected.
	NetworkID string `json:"network_id" api:"required" format:"uuid4"`
	// The instance will get an IP address from this subnet.
	SubnetID string `json:"subnet_id" api:"required" format:"uuid4"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Specifies the trunk group to which this interface belongs. Applicable only for
	// bare metal servers. Each unique port group is mapped to a separate trunk port.
	// Use this to control how interfaces are grouped across trunks.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPUnion `json:"floating_ip,omitzero"`
	// The instance will get an IP address from the selected network. If you choose to
	// add a floating IP, the instance will be reachable from the internet. Otherwise,
	// it will only have a private IP within the network.
	//
	// This field can be elided, and will marshal its zero value as "subnet".
	Type constant.Subnet `json:"type" default:"subnet"`
	paramObj
}

func (r BaremetalRegionalServerNewParamsInterfaceSubnet) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerNewParamsInterfaceSubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerNewParamsInterfaceSubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPUnion struct {
	OfNew      *BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

func (u BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNew, u.OfExisting)
}
func (u *BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPUnion](
		"source",
		apijson.Discriminator[BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPNew]("new"),
		apijson.Discriminator[BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPExisting]("existing"),
	)
}

func NewBaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPNew() BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPNew {
	return BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewBaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPNew].
type BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source" default:"new"`
	paramObj
}

func (r BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ExistingFloatingID, Source are required.
type BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPExisting struct {
	// An existing available floating IP id must be specified if the source is set to
	// `existing`
	ExistingFloatingID string `json:"existing_floating_id" api:"required" format:"uuid4"`
	// An existing available floating IP will be attached to the instance. A floating
	// IP is a public IP that makes the instance accessible from the internet, even if
	// it only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Source constant.Existing `json:"source" default:"existing"`
	paramObj
}

func (r BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerNewParamsInterfaceSubnetFloatingIPExisting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties NetworkID, Type are required.
type BaremetalRegionalServerNewParamsInterfaceAnySubnet struct {
	// The network where the instance will be connected.
	NetworkID string `json:"network_id" api:"required" format:"uuid4"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// You can specify a specific IP address from your subnet.
	IPAddress param.Opt[string] `json:"ip_address,omitzero" format:"ipvanyaddress"`
	// Specifies the trunk group to which this interface belongs. Applicable only for
	// bare metal servers. Each unique port group is mapped to a separate trunk port.
	// Use this to control how interfaces are grouped across trunks.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPUnion `json:"floating_ip,omitzero"`
	// Specify `ipv4`, `ipv6`, or `dual` to enable both.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily InterfaceIPFamily `json:"ip_family,omitzero"`
	// Instance will be attached to a subnet with the largest count of free IPs.
	//
	// This field can be elided, and will marshal its zero value as "any_subnet".
	Type constant.AnySubnet `json:"type" default:"any_subnet"`
	paramObj
}

func (r BaremetalRegionalServerNewParamsInterfaceAnySubnet) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerNewParamsInterfaceAnySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerNewParamsInterfaceAnySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPUnion struct {
	OfNew      *BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

func (u BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNew, u.OfExisting)
}
func (u *BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPUnion](
		"source",
		apijson.Discriminator[BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPNew]("new"),
		apijson.Discriminator[BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPExisting]("existing"),
	)
}

func NewBaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPNew() BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPNew {
	return BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewBaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPNew].
type BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source" default:"new"`
	paramObj
}

func (r BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ExistingFloatingID, Source are required.
type BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPExisting struct {
	// An existing available floating IP id must be specified if the source is set to
	// `existing`
	ExistingFloatingID string `json:"existing_floating_id" api:"required" format:"uuid4"`
	// An existing available floating IP will be attached to the instance. A floating
	// IP is a public IP that makes the instance accessible from the internet, even if
	// it only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Source constant.Existing `json:"source" default:"existing"`
	paramObj
}

func (r BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerNewParamsInterfaceAnySubnetFloatingIPExisting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties PortID, Type are required.
type BaremetalRegionalServerNewParamsInterfaceReservedFixedIP struct {
	// Network ID the subnet belongs to. Port will be plugged in this network.
	PortID string `json:"port_id" api:"required"`
	// Interface name. Defaults to `null` and is returned as `null` in the API response
	// if not set.
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Specifies the trunk group to which this interface belongs. Applicable only for
	// bare metal servers. Each unique port group is mapped to a separate trunk port.
	// Use this to control how interfaces are grouped across trunks.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Allows the instance to have a public IP that can be reached from the internet.
	FloatingIP BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion `json:"floating_ip,omitzero"`
	// An existing available reserved fixed IP will be attached to the instance. If the
	// reserved IP is not public and you choose to add a floating IP, the instance will
	// be accessible from the internet.
	//
	// This field can be elided, and will marshal its zero value as
	// "reserved_fixed_ip".
	Type constant.ReservedFixedIP `json:"type" default:"reserved_fixed_ip"`
	paramObj
}

func (r BaremetalRegionalServerNewParamsInterfaceReservedFixedIP) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerNewParamsInterfaceReservedFixedIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerNewParamsInterfaceReservedFixedIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion struct {
	OfNew      *BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPNew      `json:",omitzero,inline"`
	OfExisting *BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting `json:",omitzero,inline"`
	paramUnion
}

func (u BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNew, u.OfExisting)
}
func (u *BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfExisting) {
		return u.OfExisting
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion) GetExistingFloatingID() *string {
	if vt := u.OfExisting; vt != nil {
		return &vt.ExistingFloatingID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfExisting; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPUnion](
		"source",
		apijson.Discriminator[BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPNew]("new"),
		apijson.Discriminator[BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting]("existing"),
	)
}

func NewBaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPNew() BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPNew {
	return BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPNew{
		Source: "new",
	}
}

// This struct has a constant value, construct it with
// [NewBaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPNew].
type BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPNew struct {
	// A new floating IP will be created and attached to the instance. A floating IP is
	// a public IP that makes the instance accessible from the internet, even if it
	// only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	Source constant.New `json:"source" default:"new"`
	paramObj
}

func (r BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPNew) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPNew
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ExistingFloatingID, Source are required.
type BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting struct {
	// An existing available floating IP id must be specified if the source is set to
	// `existing`
	ExistingFloatingID string `json:"existing_floating_id" api:"required" format:"uuid4"`
	// An existing available floating IP will be attached to the instance. A floating
	// IP is a public IP that makes the instance accessible from the internet, even if
	// it only has a private IP. It works like SNAT, allowing outgoing and incoming
	// traffic.
	//
	// This field can be elided, and will marshal its zero value as "existing".
	Source constant.Existing `json:"source" default:"existing"`
	paramObj
}

func (r BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerNewParamsInterfaceReservedFixedIPFloatingIPExisting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enable advanced DDoS protection for the server
//
// The property ProfileTemplate is required.
type BaremetalRegionalServerNewParamsDDOSProfile struct {
	// Unique identifier of the DDoS protection template to use for this profile
	ProfileTemplate int64 `json:"profile_template" api:"required"`
	// List of field configurations that customize the protection parameters for this
	// profile
	Fields []BaremetalRegionalServerNewParamsDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r BaremetalRegionalServerNewParamsDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerNewParamsDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerNewParamsDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BaseField is required.
type BaremetalRegionalServerNewParamsDDOSProfileField struct {
	// Unique identifier of the DDoS protection field being configured
	BaseField  int64 `json:"base_field" api:"required"`
	FieldValue any   `json:"field_value,omitzero"`
	paramObj
}

func (r BaremetalRegionalServerNewParamsDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerNewParamsDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerNewParamsDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalRegionalServerUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Name
	Name param.Opt[string] `json:"name,omitzero"`
	// Update key-value tags using JSON Merge Patch semantics (RFC 7396). Provide
	// key-value pairs to add or update tags. Set tag values to `null` to remove tags.
	// Unspecified tags remain unchanged. Read-only tags are always preserved and
	// cannot be modified.
	//
	// **Examples:**
	//
	//   - **Add/update tags:**
	//     `{'tags': {'environment': 'production', 'team': 'backend'}}` adds new tags or
	//     updates existing ones.
	//   - **Delete tags:** `{'tags': {'old_tag': null}}` removes specific tags.
	//   - **Remove all tags:** `{'tags': null}` removes all user-managed tags (read-only
	//     tags are preserved).
	//   - **Partial update:** `{'tags': {'environment': 'staging'}}` only updates
	//     specified tags.
	//   - **Mixed operations:**
	//     `{'tags': {'environment': 'production', 'cost_center': 'engineering', 'deprecated_tag': null}}`
	//     adds/updates 'environment' and 'cost_center' while removing 'deprecated_tag',
	//     preserving other existing tags.
	//   - **Replace all:** first delete existing tags with null values, then add new
	//     ones in the same request.
	Tags TagUpdateMap `json:"tags,omitzero"`
	paramObj
}

func (r BaremetalRegionalServerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalRegionalServerListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Filters the instances by a date and time stamp when the instances last changed.
	ChangesBefore param.Opt[time.Time] `query:"changes-before,omitzero" format:"date-time" json:"-"`
	// Filters the instances by a date and time stamp when the instances last changed
	// status.
	ChangesSince param.Opt[time.Time] `query:"changes-since,omitzero" format:"date-time" json:"-"`
	// Filter out instances by `flavor_id`. Flavor id must match exactly.
	FlavorID param.Opt[string] `query:"flavor_id,omitzero" json:"-"`
	// Filter out instances by `flavor_prefix`.
	FlavorPrefix param.Opt[string] `query:"flavor_prefix,omitzero" json:"-"`
	// Include managed k8s worker nodes
	IncludeK8S param.Opt[bool] `query:"include_k8s,omitzero" json:"-"`
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
	// Filter result by ddos protection profile name. Effective only with `with_ddos`
	// set to true.
	ProfileName param.Opt[string] `query:"profile_name,omitzero" json:"-"`
	// Optional. Filter by tag key-value pairs.
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
	OrderBy BaremetalRegionalServerListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Filter result by DDoS `protection_status`. Effective only with `with_ddos` set
	// to true. (Active, Queued or Error)
	//
	// Any of "Active", "Queued", "Error".
	ProtectionStatus BaremetalRegionalServerListParamsProtectionStatus `query:"protection_status,omitzero" json:"-"`
	// Filters instances by a server status, as a string.
	//
	// Any of "ACTIVE", "BUILD", "ERROR", "HARD_REBOOT", "REBOOT", "REBUILD", "RESCUE",
	// "SHUTOFF", "SUSPENDED".
	Status BaremetalRegionalServerListParamsStatus `query:"status,omitzero" json:"-"`
	// Optional. Filter by tag values. ?`tag_value`=value1&`tag_value`=value2
	TagValue []string `query:"tag_value,omitzero" json:"-"`
	// Return bare metals either only with advanced or only basic DDoS protection.
	// Effective only with `with_ddos` set to true. (advanced or basic)
	//
	// Any of "basic", "advanced".
	TypeDDOSProfile BaremetalRegionalServerListParamsTypeDDOSProfile `query:"type_ddos_profile,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalRegionalServerListParams]'s query parameters as
// `url.Values`.
func (r BaremetalRegionalServerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Order by field and direction.
type BaremetalRegionalServerListParamsOrderBy string

const (
	BaremetalRegionalServerListParamsOrderByCreatedAsc  BaremetalRegionalServerListParamsOrderBy = "created.asc"
	BaremetalRegionalServerListParamsOrderByCreatedDesc BaremetalRegionalServerListParamsOrderBy = "created.desc"
	BaremetalRegionalServerListParamsOrderByNameAsc     BaremetalRegionalServerListParamsOrderBy = "name.asc"
	BaremetalRegionalServerListParamsOrderByNameDesc    BaremetalRegionalServerListParamsOrderBy = "name.desc"
	BaremetalRegionalServerListParamsOrderByStatusAsc   BaremetalRegionalServerListParamsOrderBy = "status.asc"
	BaremetalRegionalServerListParamsOrderByStatusDesc  BaremetalRegionalServerListParamsOrderBy = "status.desc"
)

// Filter result by DDoS `protection_status`. Effective only with `with_ddos` set
// to true. (Active, Queued or Error)
type BaremetalRegionalServerListParamsProtectionStatus string

const (
	BaremetalRegionalServerListParamsProtectionStatusActive BaremetalRegionalServerListParamsProtectionStatus = "Active"
	BaremetalRegionalServerListParamsProtectionStatusQueued BaremetalRegionalServerListParamsProtectionStatus = "Queued"
	BaremetalRegionalServerListParamsProtectionStatusError  BaremetalRegionalServerListParamsProtectionStatus = "Error"
)

// Filters instances by a server status, as a string.
type BaremetalRegionalServerListParamsStatus string

const (
	BaremetalRegionalServerListParamsStatusActive     BaremetalRegionalServerListParamsStatus = "ACTIVE"
	BaremetalRegionalServerListParamsStatusBuild      BaremetalRegionalServerListParamsStatus = "BUILD"
	BaremetalRegionalServerListParamsStatusError      BaremetalRegionalServerListParamsStatus = "ERROR"
	BaremetalRegionalServerListParamsStatusHardReboot BaremetalRegionalServerListParamsStatus = "HARD_REBOOT"
	BaremetalRegionalServerListParamsStatusReboot     BaremetalRegionalServerListParamsStatus = "REBOOT"
	BaremetalRegionalServerListParamsStatusRebuild    BaremetalRegionalServerListParamsStatus = "REBUILD"
	BaremetalRegionalServerListParamsStatusRescue     BaremetalRegionalServerListParamsStatus = "RESCUE"
	BaremetalRegionalServerListParamsStatusShutoff    BaremetalRegionalServerListParamsStatus = "SHUTOFF"
	BaremetalRegionalServerListParamsStatusSuspended  BaremetalRegionalServerListParamsStatus = "SUSPENDED"
)

// Return bare metals either only with advanced or only basic DDoS protection.
// Effective only with `with_ddos` set to true. (advanced or basic)
type BaremetalRegionalServerListParamsTypeDDOSProfile string

const (
	BaremetalRegionalServerListParamsTypeDDOSProfileBasic    BaremetalRegionalServerListParamsTypeDDOSProfile = "basic"
	BaremetalRegionalServerListParamsTypeDDOSProfileAdvanced BaremetalRegionalServerListParamsTypeDDOSProfile = "advanced"
)

type BaremetalRegionalServerDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// True if it is required to delete floating IPs assigned to the instance. Can't be
	// used with `floating_ip_ids`.
	AllFloatingIPs param.Opt[bool] `query:"all_floating_ips,omitzero" json:"-"`
	// True if it is required to delete reserved fixed IPs assigned to the instance.
	// Can't be used with `reserved_fixed_ip_ids`.
	AllReservedFixedIPs param.Opt[bool] `query:"all_reserved_fixed_ips,omitzero" json:"-"`
	// Comma separated list of floating ids that should be deleted. Can't be used with
	// `all_floating_ips`.
	FloatingIPIDs param.Opt[string] `query:"floating_ip_ids,omitzero" json:"-"`
	// Comma separated list of port IDs to be deleted with the instance
	ReservedFixedIPIDs param.Opt[string] `query:"reserved_fixed_ip_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalRegionalServerDeleteParams]'s query parameters as
// `url.Values`.
func (r BaremetalRegionalServerDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BaremetalRegionalServerActionParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfStartActionInstanceSerializer *BaremetalRegionalServerActionParamsBodyStartActionInstanceSerializer `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfBasicBareMetalActionInstanceSerializer *BaremetalRegionalServerActionParamsBodyBasicBareMetalActionInstanceSerializer `json:",inline"`

	paramObj
}

func (u BaremetalRegionalServerActionParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStartActionInstanceSerializer, u.OfBasicBareMetalActionInstanceSerializer)
}
func (r *BaremetalRegionalServerActionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Action is required.
type BaremetalRegionalServerActionParamsBodyStartActionInstanceSerializer struct {
	// Used on start instance to activate Advanced DDoS profile
	ActivateProfile param.Opt[bool] `json:"activate_profile,omitzero"`
	// Instance action name
	//
	// This field can be elided, and will marshal its zero value as "start".
	Action constant.Start `json:"action" default:"start"`
	paramObj
}

func (r BaremetalRegionalServerActionParamsBodyStartActionInstanceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerActionParamsBodyStartActionInstanceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerActionParamsBodyStartActionInstanceSerializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Action is required.
type BaremetalRegionalServerActionParamsBodyBasicBareMetalActionInstanceSerializer struct {
	// Instance action name
	//
	// Any of "reboot", "reboot_hard", "stop".
	Action string `json:"action,omitzero" api:"required"`
	paramObj
}

func (r BaremetalRegionalServerActionParamsBodyBasicBareMetalActionInstanceSerializer) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalRegionalServerActionParamsBodyBasicBareMetalActionInstanceSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalRegionalServerActionParamsBodyBasicBareMetalActionInstanceSerializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BaremetalRegionalServerActionParamsBodyBasicBareMetalActionInstanceSerializer](
		"action", "reboot", "reboot_hard", "stop",
	)
}

type BaremetalRegionalServerGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}
