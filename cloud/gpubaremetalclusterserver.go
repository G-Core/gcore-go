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
)

// GPUBaremetalClusterServerService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUBaremetalClusterServerService] method instead.
type GPUBaremetalClusterServerService struct {
	Options []option.RequestOption
}

// NewGPUBaremetalClusterServerService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGPUBaremetalClusterServerService(opts ...option.RequestOption) (r GPUBaremetalClusterServerService) {
	r = GPUBaremetalClusterServerService{}
	r.Options = opts
	return
}

// List all servers in a bare metal GPU cluster. Results can be filtered and
// paginated.
func (r *GPUBaremetalClusterServerService) List(ctx context.Context, clusterID string, params GPUBaremetalClusterServerListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[GPUBaremetalClusterServer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/clusters/%s/servers", params.ProjectID.Value, params.RegionID.Value, clusterID)
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

// List all servers in a bare metal GPU cluster. Results can be filtered and
// paginated.
func (r *GPUBaremetalClusterServerService) ListAutoPaging(ctx context.Context, clusterID string, params GPUBaremetalClusterServerListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[GPUBaremetalClusterServer] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, clusterID, params, opts...))
}

// Delete a specific node from a GPU cluster. The node must be in a state that
// allows deletion.
func (r *GPUBaremetalClusterServerService) Delete(ctx context.Context, instanceID string, params GPUBaremetalClusterServerDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if params.ClusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/gpu/%v/%v/%s/node/%s", params.ProjectID.Value, params.RegionID.Value, params.ClusterID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Attach interface to bare metal GPU cluster server
func (r *GPUBaremetalClusterServerService) AttachInterface(ctx context.Context, instanceID string, params GPUBaremetalClusterServerAttachInterfaceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/attach_interface", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Detach interface from bare metal GPU cluster server
func (r *GPUBaremetalClusterServerService) DetachInterface(ctx context.Context, instanceID string, params GPUBaremetalClusterServerDetachInterfaceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/detach_interface", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get bare metal GPU cluster server console URL
func (r *GPUBaremetalClusterServerService) GetConsole(ctx context.Context, instanceID string, query GPUBaremetalClusterServerGetConsoleParams, opts ...option.RequestOption) (res *Console, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/get_console", query.ProjectID.Value, query.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Stops and then starts the server, effectively performing a hard reboot.
func (r *GPUBaremetalClusterServerService) Powercycle(ctx context.Context, instanceID string, body GPUBaremetalClusterServerPowercycleParams, opts ...option.RequestOption) (res *GPUBaremetalClusterServerV1, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/powercycle", body.ProjectID.Value, body.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Reboot one bare metal GPU cluster server
func (r *GPUBaremetalClusterServerService) Reboot(ctx context.Context, instanceID string, body GPUBaremetalClusterServerRebootParams, opts ...option.RequestOption) (res *GPUBaremetalClusterServerV1, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/reboot", body.ProjectID.Value, body.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type GPUBaremetalClusterServer struct {
	// Server unique identifier
	ID string `json:"id,required" format:"uuid4"`
	// Server creation date and time
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Unique flavor identifier
	Flavor string `json:"flavor,required"`
	// Server's image UUID
	ImageID string `json:"image_id,required" format:"uuid4"`
	// List of IP addresses
	IPAddresses []string `json:"ip_addresses,required"`
	// Server's name generated using cluster's name
	Name string `json:"name,required"`
	// Security groups names
	SecurityGroups []GPUBaremetalClusterServerSecurityGroup `json:"security_groups,required"`
	// SSH key pair assigned to the server
	SSHKeyName string `json:"ssh_key_name,required"`
	// Current server status
	//
	// Any of "ACTIVE", "BUILD", "DELETED", "ERROR", "HARD_REBOOT", "MIGRATING",
	// "PASSWORD", "PAUSED", "REBOOT", "REBUILD", "RESCUE", "RESIZE", "REVERT_RESIZE",
	// "SHELVED", "SHELVED_OFFLOADED", "SHUTOFF", "SOFT_DELETED", "SUSPENDED",
	// "UNKNOWN", "VERIFY_RESIZE".
	Status GPUBaremetalClusterServerStatus `json:"status,required"`
	// User defined tags
	Tags []Tag `json:"tags,required"`
	// Identifier of the task currently modifying the GPU cluster
	TaskID string `json:"task_id,required" format:"uuid4"`
	// Server update date and time
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Flavor         respjson.Field
		ImageID        respjson.Field
		IPAddresses    respjson.Field
		Name           respjson.Field
		SecurityGroups respjson.Field
		SSHKeyName     respjson.Field
		Status         respjson.Field
		Tags           respjson.Field
		TaskID         respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServer) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServerSecurityGroup struct {
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
func (r GPUBaremetalClusterServerSecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServerSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current server status
type GPUBaremetalClusterServerStatus string

const (
	GPUBaremetalClusterServerStatusActive           GPUBaremetalClusterServerStatus = "ACTIVE"
	GPUBaremetalClusterServerStatusBuild            GPUBaremetalClusterServerStatus = "BUILD"
	GPUBaremetalClusterServerStatusDeleted          GPUBaremetalClusterServerStatus = "DELETED"
	GPUBaremetalClusterServerStatusError            GPUBaremetalClusterServerStatus = "ERROR"
	GPUBaremetalClusterServerStatusHardReboot       GPUBaremetalClusterServerStatus = "HARD_REBOOT"
	GPUBaremetalClusterServerStatusMigrating        GPUBaremetalClusterServerStatus = "MIGRATING"
	GPUBaremetalClusterServerStatusPassword         GPUBaremetalClusterServerStatus = "PASSWORD"
	GPUBaremetalClusterServerStatusPaused           GPUBaremetalClusterServerStatus = "PAUSED"
	GPUBaremetalClusterServerStatusReboot           GPUBaremetalClusterServerStatus = "REBOOT"
	GPUBaremetalClusterServerStatusRebuild          GPUBaremetalClusterServerStatus = "REBUILD"
	GPUBaremetalClusterServerStatusRescue           GPUBaremetalClusterServerStatus = "RESCUE"
	GPUBaremetalClusterServerStatusResize           GPUBaremetalClusterServerStatus = "RESIZE"
	GPUBaremetalClusterServerStatusRevertResize     GPUBaremetalClusterServerStatus = "REVERT_RESIZE"
	GPUBaremetalClusterServerStatusShelved          GPUBaremetalClusterServerStatus = "SHELVED"
	GPUBaremetalClusterServerStatusShelvedOffloaded GPUBaremetalClusterServerStatus = "SHELVED_OFFLOADED"
	GPUBaremetalClusterServerStatusShutoff          GPUBaremetalClusterServerStatus = "SHUTOFF"
	GPUBaremetalClusterServerStatusSoftDeleted      GPUBaremetalClusterServerStatus = "SOFT_DELETED"
	GPUBaremetalClusterServerStatusSuspended        GPUBaremetalClusterServerStatus = "SUSPENDED"
	GPUBaremetalClusterServerStatusUnknown          GPUBaremetalClusterServerStatus = "UNKNOWN"
	GPUBaremetalClusterServerStatusVerifyResize     GPUBaremetalClusterServerStatus = "VERIFY_RESIZE"
)

type GPUBaremetalClusterServerV1 struct {
	// GPU server ID
	ID string `json:"id,required" format:"uuid4"`
	// Map of `network_name` to list of addresses in that network
	Addresses map[string][]GPUBaremetalClusterServerV1AddressUnion `json:"addresses,required"`
	// IP addresses of the instances that are blackholed by DDoS mitigation system
	BlackholePorts []BlackholePort `json:"blackhole_ports,required"`
	// Datetime when GPU server was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required"`
	// Advanced DDoS protection profile. It is always `null` if query parameter
	// `with_ddos=true` is not set.
	DDOSProfile DDOSProfile `json:"ddos_profile,required"`
	// Fixed IP assigned to instance
	FixedIPAssignments []GPUBaremetalClusterServerV1FixedIPAssignment `json:"fixed_ip_assignments,required"`
	// Flavor
	Flavor GPUBaremetalClusterServerV1Flavor `json:"flavor,required"`
	// Instance description
	InstanceDescription string `json:"instance_description,required"`
	// Instance isolation information
	InstanceIsolation InstanceIsolation `json:"instance_isolation,required"`
	// GPU server name
	Name string `json:"name,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Security groups
	SecurityGroups []GPUBaremetalClusterServerV1SecurityGroup `json:"security_groups,required"`
	// SSH key name assigned to instance
	SSHKeyName string `json:"ssh_key_name,required"`
	// Instance status
	//
	// Any of "ACTIVE", "BUILD", "DELETED", "ERROR", "HARD_REBOOT", "MIGRATING",
	// "PASSWORD", "PAUSED", "REBOOT", "REBUILD", "RESCUE", "RESIZE", "REVERT_RESIZE",
	// "SHELVED", "SHELVED_OFFLOADED", "SHUTOFF", "SOFT_DELETED", "SUSPENDED",
	// "UNKNOWN", "VERIFY_RESIZE".
	Status GPUBaremetalClusterServerV1Status `json:"status,required"`
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
	VmState GPUBaremetalClusterServerV1VmState `json:"vm_state,required"`
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
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServerV1) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServerV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUBaremetalClusterServerV1AddressUnion contains all possible properties and
// values from [FloatingAddress], [FixedAddressShort], [FixedAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type GPUBaremetalClusterServerV1AddressUnion struct {
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

func (u GPUBaremetalClusterServerV1AddressUnion) AsFloatingIPAddress() (v FloatingAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GPUBaremetalClusterServerV1AddressUnion) AsFixedIPAddressShort() (v FixedAddressShort) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GPUBaremetalClusterServerV1AddressUnion) AsFixedIPAddress() (v FixedAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GPUBaremetalClusterServerV1AddressUnion) RawJSON() string { return u.JSON.raw }

func (r *GPUBaremetalClusterServerV1AddressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServerV1FixedIPAssignment struct {
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
func (r GPUBaremetalClusterServerV1FixedIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServerV1FixedIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor
type GPUBaremetalClusterServerV1Flavor struct {
	// CPU architecture
	Architecture string `json:"architecture,required"`
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id,required"`
	// Flavor name
	FlavorName string `json:"flavor_name,required"`
	// Additional hardware description
	HardwareDescription GPUBaremetalClusterServerV1FlavorHardwareDescription `json:"hardware_description,required"`
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
func (r GPUBaremetalClusterServerV1Flavor) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServerV1Flavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional hardware description
type GPUBaremetalClusterServerV1FlavorHardwareDescription struct {
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
func (r GPUBaremetalClusterServerV1FlavorHardwareDescription) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServerV1FlavorHardwareDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServerV1SecurityGroup struct {
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
func (r GPUBaremetalClusterServerV1SecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServerV1SecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance status
type GPUBaremetalClusterServerV1Status string

const (
	GPUBaremetalClusterServerV1StatusActive           GPUBaremetalClusterServerV1Status = "ACTIVE"
	GPUBaremetalClusterServerV1StatusBuild            GPUBaremetalClusterServerV1Status = "BUILD"
	GPUBaremetalClusterServerV1StatusDeleted          GPUBaremetalClusterServerV1Status = "DELETED"
	GPUBaremetalClusterServerV1StatusError            GPUBaremetalClusterServerV1Status = "ERROR"
	GPUBaremetalClusterServerV1StatusHardReboot       GPUBaremetalClusterServerV1Status = "HARD_REBOOT"
	GPUBaremetalClusterServerV1StatusMigrating        GPUBaremetalClusterServerV1Status = "MIGRATING"
	GPUBaremetalClusterServerV1StatusPassword         GPUBaremetalClusterServerV1Status = "PASSWORD"
	GPUBaremetalClusterServerV1StatusPaused           GPUBaremetalClusterServerV1Status = "PAUSED"
	GPUBaremetalClusterServerV1StatusReboot           GPUBaremetalClusterServerV1Status = "REBOOT"
	GPUBaremetalClusterServerV1StatusRebuild          GPUBaremetalClusterServerV1Status = "REBUILD"
	GPUBaremetalClusterServerV1StatusRescue           GPUBaremetalClusterServerV1Status = "RESCUE"
	GPUBaremetalClusterServerV1StatusResize           GPUBaremetalClusterServerV1Status = "RESIZE"
	GPUBaremetalClusterServerV1StatusRevertResize     GPUBaremetalClusterServerV1Status = "REVERT_RESIZE"
	GPUBaremetalClusterServerV1StatusShelved          GPUBaremetalClusterServerV1Status = "SHELVED"
	GPUBaremetalClusterServerV1StatusShelvedOffloaded GPUBaremetalClusterServerV1Status = "SHELVED_OFFLOADED"
	GPUBaremetalClusterServerV1StatusShutoff          GPUBaremetalClusterServerV1Status = "SHUTOFF"
	GPUBaremetalClusterServerV1StatusSoftDeleted      GPUBaremetalClusterServerV1Status = "SOFT_DELETED"
	GPUBaremetalClusterServerV1StatusSuspended        GPUBaremetalClusterServerV1Status = "SUSPENDED"
	GPUBaremetalClusterServerV1StatusUnknown          GPUBaremetalClusterServerV1Status = "UNKNOWN"
	GPUBaremetalClusterServerV1StatusVerifyResize     GPUBaremetalClusterServerV1Status = "VERIFY_RESIZE"
)

// Virtual machine state (active)
type GPUBaremetalClusterServerV1VmState string

const (
	GPUBaremetalClusterServerV1VmStateActive           GPUBaremetalClusterServerV1VmState = "active"
	GPUBaremetalClusterServerV1VmStateBuilding         GPUBaremetalClusterServerV1VmState = "building"
	GPUBaremetalClusterServerV1VmStateDeleted          GPUBaremetalClusterServerV1VmState = "deleted"
	GPUBaremetalClusterServerV1VmStateError            GPUBaremetalClusterServerV1VmState = "error"
	GPUBaremetalClusterServerV1VmStatePaused           GPUBaremetalClusterServerV1VmState = "paused"
	GPUBaremetalClusterServerV1VmStateRescued          GPUBaremetalClusterServerV1VmState = "rescued"
	GPUBaremetalClusterServerV1VmStateResized          GPUBaremetalClusterServerV1VmState = "resized"
	GPUBaremetalClusterServerV1VmStateShelved          GPUBaremetalClusterServerV1VmState = "shelved"
	GPUBaremetalClusterServerV1VmStateShelvedOffloaded GPUBaremetalClusterServerV1VmState = "shelved_offloaded"
	GPUBaremetalClusterServerV1VmStateSoftDeleted      GPUBaremetalClusterServerV1VmState = "soft-deleted"
	GPUBaremetalClusterServerV1VmStateStopped          GPUBaremetalClusterServerV1VmState = "stopped"
	GPUBaremetalClusterServerV1VmStateSuspended        GPUBaremetalClusterServerV1VmState = "suspended"
)

type GPUBaremetalClusterServerV1List struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []GPUBaremetalClusterServerV1 `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServerV1List) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServerV1List) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServerListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Filters the results to include only servers whose last change timestamp is less
	// than the specified datetime. Format: ISO 8601.
	ChangedBefore param.Opt[time.Time] `query:"changed_before,omitzero" format:"date-time" json:"-"`
	// Filters the results to include only servers whose last change timestamp is
	// greater than or equal to the specified datetime. Format: ISO 8601.
	ChangedSince param.Opt[time.Time] `query:"changed_since,omitzero" format:"date-time" json:"-"`
	// Filter servers by ip address.
	IPAddress param.Opt[string] `query:"ip_address,omitzero" json:"-"`
	// Limit of items on a single page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter servers by name. You can provide a full or partial name, servers with
	// matching names will be returned. For example, entering 'test' will return all
	// servers that contain 'test' in their name.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Offset in results list
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Order field
	//
	// Any of "created_at.asc", "created_at.desc", "status.asc", "status.desc".
	OrderBy GPUBaremetalClusterServerListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Filters servers by status.
	//
	// Any of "ACTIVE", "BUILD", "ERROR", "HARD_REBOOT", "MIGRATING", "PAUSED",
	// "REBOOT", "REBUILD", "RESIZE", "REVERT_RESIZE", "SHELVED", "SHELVED_OFFLOADED",
	// "SHUTOFF", "SOFT_DELETED", "SUSPENDED", "VERIFY_RESIZE".
	Status GPUBaremetalClusterServerListParamsStatus `query:"status,omitzero" json:"-"`
	// Filter servers by uuid.
	Uuids []string `query:"uuids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GPUBaremetalClusterServerListParams]'s query parameters as
// `url.Values`.
func (r GPUBaremetalClusterServerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Order field
type GPUBaremetalClusterServerListParamsOrderBy string

const (
	GPUBaremetalClusterServerListParamsOrderByCreatedAtAsc  GPUBaremetalClusterServerListParamsOrderBy = "created_at.asc"
	GPUBaremetalClusterServerListParamsOrderByCreatedAtDesc GPUBaremetalClusterServerListParamsOrderBy = "created_at.desc"
	GPUBaremetalClusterServerListParamsOrderByStatusAsc     GPUBaremetalClusterServerListParamsOrderBy = "status.asc"
	GPUBaremetalClusterServerListParamsOrderByStatusDesc    GPUBaremetalClusterServerListParamsOrderBy = "status.desc"
)

// Filters servers by status.
type GPUBaremetalClusterServerListParamsStatus string

const (
	GPUBaremetalClusterServerListParamsStatusActive           GPUBaremetalClusterServerListParamsStatus = "ACTIVE"
	GPUBaremetalClusterServerListParamsStatusBuild            GPUBaremetalClusterServerListParamsStatus = "BUILD"
	GPUBaremetalClusterServerListParamsStatusError            GPUBaremetalClusterServerListParamsStatus = "ERROR"
	GPUBaremetalClusterServerListParamsStatusHardReboot       GPUBaremetalClusterServerListParamsStatus = "HARD_REBOOT"
	GPUBaremetalClusterServerListParamsStatusMigrating        GPUBaremetalClusterServerListParamsStatus = "MIGRATING"
	GPUBaremetalClusterServerListParamsStatusPaused           GPUBaremetalClusterServerListParamsStatus = "PAUSED"
	GPUBaremetalClusterServerListParamsStatusReboot           GPUBaremetalClusterServerListParamsStatus = "REBOOT"
	GPUBaremetalClusterServerListParamsStatusRebuild          GPUBaremetalClusterServerListParamsStatus = "REBUILD"
	GPUBaremetalClusterServerListParamsStatusResize           GPUBaremetalClusterServerListParamsStatus = "RESIZE"
	GPUBaremetalClusterServerListParamsStatusRevertResize     GPUBaremetalClusterServerListParamsStatus = "REVERT_RESIZE"
	GPUBaremetalClusterServerListParamsStatusShelved          GPUBaremetalClusterServerListParamsStatus = "SHELVED"
	GPUBaremetalClusterServerListParamsStatusShelvedOffloaded GPUBaremetalClusterServerListParamsStatus = "SHELVED_OFFLOADED"
	GPUBaremetalClusterServerListParamsStatusShutoff          GPUBaremetalClusterServerListParamsStatus = "SHUTOFF"
	GPUBaremetalClusterServerListParamsStatusSoftDeleted      GPUBaremetalClusterServerListParamsStatus = "SOFT_DELETED"
	GPUBaremetalClusterServerListParamsStatusSuspended        GPUBaremetalClusterServerListParamsStatus = "SUSPENDED"
	GPUBaremetalClusterServerListParamsStatusVerifyResize     GPUBaremetalClusterServerListParamsStatus = "VERIFY_RESIZE"
)

type GPUBaremetalClusterServerDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	ClusterID string           `path:"cluster_id,required" json:"-"`
	// Set False if you do not want to delete assigned floating IPs. By default, it's
	// True.
	DeleteFloatings param.Opt[bool] `query:"delete_floatings,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GPUBaremetalClusterServerDeleteParams]'s query parameters
// as `url.Values`.
func (r GPUBaremetalClusterServerDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type GPUBaremetalClusterServerAttachInterfaceParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to default external network
	OfNewInterfaceExternalExtendSchemaWithDDOS *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to specified subnet
	OfNewInterfaceSpecificSubnetSchema *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchema `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to the network subnet with the largest count of
	// available ips
	OfNewInterfaceAnySubnetSchema *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Instance will be attached to the given port. Floating IP will be created and
	// attached to that IP
	OfNewInterfaceReservedFixedIPSchema *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchema `json:",inline"`

	paramObj
}

func (u GPUBaremetalClusterServerAttachInterfaceParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNewInterfaceExternalExtendSchemaWithDDOS, u.OfNewInterfaceSpecificSubnetSchema, u.OfNewInterfaceAnySubnetSchema, u.OfNewInterfaceReservedFixedIPSchema)
}
func (r *GPUBaremetalClusterServerAttachInterfaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance will be attached to default external network
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS struct {
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be 'external'. Union tag
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile `json:"ddos_profile,omitzero"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// List of security group IDs
	SecurityGroups []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOS](
		"ip_family", "dual", "ipv4", "ipv6",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField struct {
	// ID of DDoS profile field
	BaseField param.Opt[int64] `json:"base_field,omitzero"`
	// Name of DDoS profile field
	FieldName param.Opt[string] `json:"field_name,omitzero"`
	// Basic type value. Only one of 'value' or '`field_value`' must be specified.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value. Only one of 'value' or '`field_value`' must be specified.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MandatoryIdSchema schema
//
// The property ID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceExternalExtendSchemaWithDDOSSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance will be attached to specified subnet
//
// The property SubnetID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchema struct {
	// Port will get an IP address from this subnet
	SubnetID string `json:"subnet_id,required"`
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be 'subnet'
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs
	SecurityGroups []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchema) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField struct {
	// ID of DDoS profile field
	BaseField param.Opt[int64] `json:"base_field,omitzero"`
	// Name of DDoS profile field
	FieldName param.Opt[string] `json:"field_name,omitzero"`
	// Basic type value. Only one of 'value' or '`field_value`' must be specified.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value. Only one of 'value' or '`field_value`' must be specified.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MandatoryIdSchema schema
//
// The property ID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceSpecificSubnetSchemaSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance will be attached to the network subnet with the largest count of
// available ips
//
// The property NetworkID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema struct {
	// Port will get an IP address in this network subnet
	NetworkID string `json:"network_id,required"`
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be '`any_subnet`'
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// List of security group IDs
	SecurityGroups []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchema](
		"ip_family", "dual", "ipv4", "ipv6",
	)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField struct {
	// ID of DDoS profile field
	BaseField param.Opt[int64] `json:"base_field,omitzero"`
	// Name of DDoS profile field
	FieldName param.Opt[string] `json:"field_name,omitzero"`
	// Basic type value. Only one of 'value' or '`field_value`' must be specified.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value. Only one of 'value' or '`field_value`' must be specified.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MandatoryIdSchema schema
//
// The property ID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceAnySubnetSchemaSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance will be attached to the given port. Floating IP will be created and
// attached to that IP
//
// The property PortID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchema struct {
	// Port ID
	PortID string `json:"port_id,required"`
	// Interface name
	InterfaceName param.Opt[string] `json:"interface_name,omitzero"`
	// Each group will be added to the separate trunk.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	// Must be '`reserved_fixed_ip`'. Union tag
	Type param.Opt[string] `json:"type,omitzero"`
	// Advanced DDoS protection.
	DDOSProfile GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile `json:"ddos_profile,omitzero"`
	// List of security group IDs
	SecurityGroups []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchema) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Advanced DDoS protection.
//
// The property ProfileTemplate is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile struct {
	// DDoS profile template ID.
	ProfileTemplate int64 `json:"profile_template,required"`
	// DDoS profile template name.
	ProfileTemplateName param.Opt[string] `json:"profile_template_name,omitzero"`
	// Protection parameters.
	Fields []GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField `json:"fields,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField struct {
	// ID of DDoS profile field
	BaseField param.Opt[int64] `json:"base_field,omitzero"`
	// Name of DDoS profile field
	FieldName param.Opt[string] `json:"field_name,omitzero"`
	// Basic type value. Only one of 'value' or '`field_value`' must be specified.
	Value param.Opt[string] `json:"value,omitzero"`
	// Complex value. Only one of 'value' or '`field_value`' must be specified.
	FieldValue any `json:"field_value,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaDDOSProfileField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MandatoryIdSchema schema
//
// The property ID is required.
type GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerAttachInterfaceParamsBodyNewInterfaceReservedFixedIPSchemaSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServerDetachInterfaceParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// IP address
	IPAddress string `json:"ip_address,required"`
	// ID of the port
	PortID string `json:"port_id,required"`
	paramObj
}

func (r GPUBaremetalClusterServerDetachInterfaceParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterServerDetachInterfaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterServerDetachInterfaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServerGetConsoleParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type GPUBaremetalClusterServerPowercycleParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type GPUBaremetalClusterServerRebootParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
