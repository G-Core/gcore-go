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

// Bare metal servers are dedicated physical machines that boot an OS image from a
// local disk and connect to networks through ports, in any region of the project.
//
// BaremetalServerService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBaremetalServerService] method instead.
type BaremetalServerService struct {
	Options []option.RequestOption
	Console BaremetalServerConsoleService
}

// NewBaremetalServerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBaremetalServerService(opts ...option.RequestOption) (r BaremetalServerService) {
	r = BaremetalServerService{}
	r.Options = opts
	r.Console = NewBaremetalServerConsoleService(opts...)
	return
}

// Create a bare metal server. The server writes an OS image to its local disk and
// boots from it. The server uses ports that you create first, in the region of the
// server.
//
// To create one:
//
//  1. Pick an image with [List bare metal images](#operation/BmImagesViewSet.get).
//  2. Pick a flavor with
//     [List bare metal flavors](#operation/GlobalBaremetalFlavorListView.get). Send
//     `compatible_with_image_id` to get only the flavors that can boot the image.
//     The `networking_capabilities` of the flavor tell which port settings it
//     accepts.
//  3. [Create a port](#operation/GlobalPortListView.post) for each interface. A
//     body with only `region_id` gives one public IPv4 address.
//
// The order of `interfaces` is the NIC order and the boot order. Omit `port_group`
// to put all interfaces in the first bond. You can also add an
// [SSH key](/docs/api-reference/cloud/ssh-keys/add-or-generate-ssh-key).
//
// The response holds the IDs of the tasks that create the server.
//
// With `?dry_run=true` the request runs every check and changes nothing. The
// response holds an empty `tasks` array. A failed check returns the same error as
// the real create.
func (r *BaremetalServerService) New(ctx context.Context, params BaremetalServerNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/baremetal/servers", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Send only the fields to change. The response holds the IDs of the tasks that
// make the changes.
//
//   - Rename: send `name`.
//   - Add an interface: [create a port](#operation/GlobalPortListView.post). Then
//     send `interfaces` with the current interfaces in their current order and the
//     new port at the end.
//   - Remove an interface: send `interfaces` without it. The port stays and you can
//     use it again. The server keeps at least one interface.
//   - Change tags: send `tags` as a JSON Merge Patch.
//
// To change the security groups,
// [update the ports](#operation/GlobalPortDetailView.patch). The update does not
// change the flavor or the image.
//
// With `?dry_run=true` the request runs every check and changes nothing. The
// response holds an empty `tasks` array. A failed check returns the same error as
// the real update.
func (r *BaremetalServerService) Update(ctx context.Context, serverID string, params BaremetalServerUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if serverID == "" {
		err = errors.New("missing required server_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/baremetal/servers/%s", params.ProjectID.Value, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List a project's bare metal servers across all regions.
func (r *BaremetalServerService) List(ctx context.Context, params BaremetalServerListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[BaremetalServer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/baremetal/servers", params.ProjectID.Value)
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

// List a project's bare metal servers across all regions.
func (r *BaremetalServerService) ListAutoPaging(ctx context.Context, params BaremetalServerListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[BaremetalServer] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete a bare metal server. Ports and floating IPs are preserved by default:
// ports are detached (the trunk is dissolved, each port survives standalone) and
// floating IPs are disassociated. Asynchronous - returns the ID of the task
// performing the deletion.
//
// `resources_to_delete` holds:
//
//   - the server itself;
//   - the ports and sub-ports that `port_ids` names, or every port if `all_ports` is
//     `true`;
//   - the floating IPs that `floating_ip_ids` names, or every floating IP if
//     `all_floating_ips` is `true`.
//
// `resources_to_preserve` holds every other port, sub-port and floating IP of the
// server. A floating IP on a deleted port stays, with no port.
//
// [`DELETE /v1/bminstances/{project_id}/{region_id}/{server_id}`](#operation/BMInstancesViewSet.delete)
// is different. It deletes every port, except the reserved fixed IPs that
// `reserved_fixed_ip_ids` and `all_reserved_fixed_ips` do not select.
//
// With `?dry_run=true` the request runs every check and changes nothing. The
// response holds an empty `tasks` array and the same two lists. A failed check
// returns the same error as the real delete.
func (r *BaremetalServerService) Delete(ctx context.Context, serverID string, params BaremetalServerDeleteParams, opts ...option.RequestOption) (res *BaremetalServerDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if serverID == "" {
		err = errors.New("missing required server_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/baremetal/servers/%s", params.ProjectID.Value, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Change the power state of a bare metal server. The action is one of: `start`,
// `stop`, `reboot` or `reboot_hard`. The action does not change the Advanced DDoS
// profile of the server. Use the DDoS profile resource to activate or to
// deactivate protection. Returns the ID of the task that runs the action.
//
// With `?dry_run=true` the request runs every check and changes nothing. The
// response holds an empty `tasks` array. A failed check returns the same error as
// the real action.
func (r *BaremetalServerService) Action(ctx context.Context, serverID string, params BaremetalServerActionParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if serverID == "" {
		err = errors.New("missing required server_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/baremetal/servers/%s/action", params.ProjectID.Value, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a single bare metal server by ID, resolved across all regions.
func (r *BaremetalServerService) Get(ctx context.Context, serverID string, query BaremetalServerGetParams, opts ...option.RequestOption) (res *BaremetalServer, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if serverID == "" {
		err = errors.New("missing required server_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/baremetal/servers/%s", query.ProjectID.Value, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Rebuild a bare metal server with a new image while preserving its configuration.
func (r *BaremetalServerService) Rebuild(ctx context.Context, serverID string, params BaremetalServerRebuildParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/bminstances/%v/%v/%s/rebuild", params.ProjectID.Value, params.RegionID.Value, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type BaremetalServer struct {
	// Baremetal server ID.
	ID string `json:"id" api:"required" format:"uuid4"`
	// Datetime when the server was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Flavor name.
	FlavorName string `json:"flavor_name" api:"required"`
	// Image ID the server was provisioned from.
	ImageID string `json:"image_id" api:"required" format:"uuid4"`
	// Attached network interfaces with trunk/bond metadata.
	Interfaces []BaremetalServerInterface `json:"interfaces" api:"required"`
	// Service managing this server, or null if customer-operated.
	ManagedBy BaremetalServerManagedBy `json:"managed_by" api:"required"`
	// Baremetal server name.
	Name string `json:"name" api:"required"`
	// Region ID.
	RegionID int64 `json:"region_id" api:"required"`
	// ID of the SSH key of the default user, or null.
	SSHKeyID string `json:"ssh_key_id" api:"required" format:"uuid4"`
	// Canonical lifecycle status.
	//
	// Any of "active", "deleted", "deleting", "error", "provisioning", "rebooting",
	// "rebuilding", "resizing", "stopped", "unknown", "updating".
	Status BaremetalServerStatus `json:"status" api:"required"`
	// Read-only key-value tags set by the platform.
	SystemTags map[string]string `json:"system_tags" api:"required"`
	// Customer-owned key-value tags.
	Tags map[string]string `json:"tags" api:"required"`
	// Datetime when the server was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		FlavorName  respjson.Field
		ImageID     respjson.Field
		Interfaces  respjson.Field
		ManagedBy   respjson.Field
		Name        respjson.Field
		RegionID    respjson.Field
		SSHKeyID    respjson.Field
		Status      respjson.Field
		SystemTags  respjson.Field
		Tags        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalServer) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare-metal interface: the VM interface fields plus trunk/bond metadata.
//
// `port_group` is the 0-based bond index (`null` for DPU/SmartNIC, which has no
// trunk). `trunk` conveys the role within the bond.
type BaremetalServerInterface struct {
	// Fixed IPs assigned to the port, each with its floating IP.
	IPAssignments []BaremetalServerInterfaceIPAssignment `json:"ip_assignments" api:"required"`
	// Port MAC address.
	MacAddress string `json:"mac_address" api:"required"`
	// Network the port is attached to.
	NetworkID string `json:"network_id" api:"required" format:"uuid4"`
	// 0-based bond index. `null` on DPU/SmartNIC bare metal (no trunk).
	PortGroup int64 `json:"port_group" api:"required"`
	// ID of the port that connects the server to the network.
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	// Trunk role within the bond. `null` when the port is not part of a trunk.
	Trunk BaremetalServerInterfaceTrunkUnion `json:"trunk" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAssignments respjson.Field
		MacAddress    respjson.Field
		NetworkID     respjson.Field
		PortGroup     respjson.Field
		PortID        respjson.Field
		Trunk         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalServerInterface) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalServerInterfaceIPAssignment struct {
	// Floating IP that translates this address. `null` when the address has no
	// floating IP. An IPv6 address always has `null`.
	FloatingIP BaremetalServerInterfaceIPAssignmentFloatingIP `json:"floating_ip" api:"required"`
	// The IP address assigned to the port from the specified subnet
	IPAddress string `json:"ip_address" api:"required" format:"ipvanyaddress"`
	// ID of the subnet that allocated the IP
	SubnetID string `json:"subnet_id" api:"required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FloatingIP  respjson.Field
		IPAddress   respjson.Field
		SubnetID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalServerInterfaceIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerInterfaceIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Floating IP that translates this address. `null` when the address has no
// floating IP. An IPv6 address always has `null`.
type BaremetalServerInterfaceIPAssignmentFloatingIP struct {
	// Floating IP ID.
	ID string `json:"id" api:"required" format:"uuid4"`
	// Public IP address of the floating IP.
	IPAddress string `json:"ip_address" api:"required" format:"ipvanyaddress"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IPAddress   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalServerInterfaceIPAssignmentFloatingIP) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerInterfaceIPAssignmentFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BaremetalServerInterfaceTrunkUnion contains all possible properties and values
// from [BaremetalServerInterfaceTrunkParent],
// [BaremetalServerInterfaceTrunkSubPort].
//
// Use the [BaremetalServerInterfaceTrunkUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BaremetalServerInterfaceTrunkUnion struct {
	// Any of "parent", "sub_port".
	Role string `json:"role"`
	// This field is from variant [BaremetalServerInterfaceTrunkSubPort].
	ParentPortID string `json:"parent_port_id"`
	JSON         struct {
		Role         respjson.Field
		ParentPortID respjson.Field
		raw          string
	} `json:"-"`
}

// anyBaremetalServerInterfaceTrunk is implemented by each variant of
// [BaremetalServerInterfaceTrunkUnion] to add type safety for the return type of
// [BaremetalServerInterfaceTrunkUnion.AsAny]
type anyBaremetalServerInterfaceTrunk interface {
	implBaremetalServerInterfaceTrunkUnion()
}

func (BaremetalServerInterfaceTrunkParent) implBaremetalServerInterfaceTrunkUnion()  {}
func (BaremetalServerInterfaceTrunkSubPort) implBaremetalServerInterfaceTrunkUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := BaremetalServerInterfaceTrunkUnion.AsAny().(type) {
//	case cloud.BaremetalServerInterfaceTrunkParent:
//	case cloud.BaremetalServerInterfaceTrunkSubPort:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u BaremetalServerInterfaceTrunkUnion) AsAny() anyBaremetalServerInterfaceTrunk {
	switch u.Role {
	case "parent":
		return u.AsParent()
	case "sub_port":
		return u.AsSubPort()
	}
	return nil
}

func (u BaremetalServerInterfaceTrunkUnion) AsParent() (v BaremetalServerInterfaceTrunkParent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BaremetalServerInterfaceTrunkUnion) AsSubPort() (v BaremetalServerInterfaceTrunkSubPort) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BaremetalServerInterfaceTrunkUnion) RawJSON() string { return u.JSON.raw }

func (r *BaremetalServerInterfaceTrunkUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalServerInterfaceTrunkParent struct {
	// Role of this interface within its bond/trunk.
	Role constant.Parent `json:"role" default:"parent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalServerInterfaceTrunkParent) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerInterfaceTrunkParent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalServerInterfaceTrunkSubPort struct {
	// Parent (bond) port ID.
	ParentPortID string `json:"parent_port_id" api:"required" format:"uuid4"`
	// Role of this interface within its bond/trunk.
	Role constant.SubPort `json:"role" default:"sub_port"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ParentPortID respjson.Field
		Role         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalServerInterfaceTrunkSubPort) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerInterfaceTrunkSubPort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Service managing this server, or null if customer-operated.
type BaremetalServerManagedBy struct {
	// ID of the managing resource. Null when it is not known.
	ResourceID string `json:"resource_id" api:"required"`
	// Name of the managing resource.
	ResourceName string `json:"resource_name" api:"required"`
	// Type of the managing service.
	//
	// Any of "k8s_cluster".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResourceID   respjson.Field
		ResourceName respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalServerManagedBy) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerManagedBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Canonical lifecycle status.
type BaremetalServerStatus string

const (
	BaremetalServerStatusActive       BaremetalServerStatus = "active"
	BaremetalServerStatusDeleted      BaremetalServerStatus = "deleted"
	BaremetalServerStatusDeleting     BaremetalServerStatus = "deleting"
	BaremetalServerStatusError        BaremetalServerStatus = "error"
	BaremetalServerStatusProvisioning BaremetalServerStatus = "provisioning"
	BaremetalServerStatusRebooting    BaremetalServerStatus = "rebooting"
	BaremetalServerStatusRebuilding   BaremetalServerStatus = "rebuilding"
	BaremetalServerStatusResizing     BaremetalServerStatus = "resizing"
	BaremetalServerStatusStopped      BaremetalServerStatus = "stopped"
	BaremetalServerStatusUnknown      BaremetalServerStatus = "unknown"
	BaremetalServerStatusUpdating     BaremetalServerStatus = "updating"
)

type BaremetalServerDeleteResponse struct {
	// Every resource that this request deletes: the server, and each port and floating
	// IP that goes with it. Sorted by `type`, then by `id`.
	ResourcesToDelete []BaremetalServerDeleteResponseResourcesToDelete `json:"resources_to_delete" api:"required"`
	// The ports and floating IPs of the server that stay after the delete. Sorted by
	// `type`, then by `id`.
	ResourcesToPreserve []BaremetalServerDeleteResponseResourcesToPreserve `json:"resources_to_preserve" api:"required"`
	// List of task IDs representing asynchronous operations. Use these IDs to monitor
	// operation progress:
	//
	//   - `GET /v1/tasks/{task_id}` - Check individual task status and details Poll task
	//     status until completion (`FINISHED`/`ERROR`) before proceeding with dependent
	//     operations.
	Tasks []string `json:"tasks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResourcesToDelete   respjson.Field
		ResourcesToPreserve respjson.Field
		Tasks               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalServerDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalServerDeleteResponseResourcesToDelete struct {
	// The ID of the resource.
	ID string `json:"id" api:"required"`
	// The name of the resource. For a floating IP, the IP address. `null` for a
	// volume, and for a resource without a name.
	Name string `json:"name" api:"required"`
	// The type of the resource.
	//
	// Any of "baremetal_servers", "floating_ips", "ports".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalServerDeleteResponseResourcesToDelete) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerDeleteResponseResourcesToDelete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalServerDeleteResponseResourcesToPreserve struct {
	// The ID of the resource.
	ID string `json:"id" api:"required"`
	// The name of the resource. For a floating IP, the IP address. `null` for a
	// volume, and for a resource without a name.
	Name string `json:"name" api:"required"`
	// The type of the resource.
	//
	// Any of "baremetal_servers", "floating_ips", "ports".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaremetalServerDeleteResponseResourcesToPreserve) RawJSON() string { return r.JSON.raw }
func (r *BaremetalServerDeleteResponseResourcesToPreserve) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalServerNewParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Name of a bare metal flavor. The flavor sets the CPU, the memory, the disks and
	// the network cards of the server.
	// `GET /v1/projects/{project_id}/baremetal/flavors` lists the flavors.
	FlavorName string `json:"flavor_name" api:"required"`
	// OS image to write to the server's local disk.
	ImageID string `json:"image_id" api:"required" format:"uuid4"`
	// Pre-created ports to attach, in NIC/boot order.
	Interfaces []BaremetalServerNewParamsInterface `json:"interfaces,omitzero" api:"required"`
	// Region to provision the server in.
	RegionID int64 `json:"region_id" api:"required"`
	// Run every check of the real request and change nothing. The response has the
	// status code and the schema of the real call, and an empty list of tasks. With
	// `false`, the default, the request does the operation.
	DryRun param.Opt[bool] `query:"dry_run,omitzero" json:"-"`
	// Server name. Auto-generated if omitted.
	Name param.Opt[string] `json:"name,omitzero"`
	// Password of the default user. No response returns it. A Windows image needs a
	// password. For a Windows image, the create checks these rules: the password has a
	// minimum of 6 characters, it contains no word 'admin', and it contains 3 of these
	// character types: digits, lowercase letters, uppercase letters, other characters.
	// For a Linux image, the password enables SSH login with a password. Send
	// `password` or `user_data`, not both.
	Password param.Opt[string] `json:"password,omitzero" format:"password"`
	// ID of the SSH key to add to the default user.
	SSHKeyID param.Opt[string] `json:"ssh_key_id,omitzero" format:"uuid4"`
	// cloud-init user data (raw or base64). Windows: may be combined with password
	// (password sets the Administrator account). Linux: mutually exclusive with
	// password.
	UserData param.Opt[string] `json:"user_data,omitzero"`
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

func (r BaremetalServerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalServerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BaremetalServerNewParams]'s query parameters as
// `url.Values`.
func (r BaremetalServerNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The property PortID is required.
type BaremetalServerNewParamsInterface struct {
	// Existing port to attach. Array order is the NIC/boot order.
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	// 0-based bond/trunk group, dense across all interfaces. Omitted defaults to
	// group 0. Must be omitted on DPU/SmartNIC flavors.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	paramObj
}

func (r BaremetalServerNewParamsInterface) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerNewParamsInterface
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalServerNewParamsInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalServerUpdateParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Run every check of the real request and change nothing. The response has the
	// status code and the schema of the real call, and an empty list of tasks. With
	// `false`, the default, the request does the operation.
	DryRun param.Opt[bool] `query:"dry_run,omitzero" json:"-"`
	// Server name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Full list of the interfaces of the server. Interfaces that are not in the list
	// are detached. Keep the current interfaces in their current order. Add new
	// interfaces at the end of the list. Send at least one interface.
	Interfaces []BaremetalServerUpdateParamsInterface `json:"interfaces,omitzero"`
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

func (r BaremetalServerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalServerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BaremetalServerUpdateParams]'s query parameters as
// `url.Values`.
func (r BaremetalServerUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The property PortID is required.
type BaremetalServerUpdateParamsInterface struct {
	// Existing port to attach. Array order is the NIC/boot order.
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	// 0-based bond/trunk group, dense across all interfaces. Omitted defaults to
	// group 0. Must be omitted on DPU/SmartNIC flavors.
	PortGroup param.Opt[int64] `json:"port_group,omitzero"`
	paramObj
}

func (r BaremetalServerUpdateParamsInterface) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerUpdateParamsInterface
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalServerUpdateParamsInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaremetalServerListParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Max items to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by managed status. true = managed by a service, false =
	// customer-operated.
	Managed param.Opt[bool] `query:"managed,omitzero" json:"-"`
	// Items to skip.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Range filter on creation time.
	CreatedAt BaremetalServerListParamsCreatedAt `query:"created_at,omitzero" json:"-"`
	// Operator flavor name filter.
	FlavorName BaremetalServerListParamsFlavorName `query:"flavor_name,omitzero" json:"-"`
	// Filter by baremetal server UUID(s).
	IDs []string `query:"ids,omitzero" format:"uuid4" json:"-"`
	// Filter by an IP address of a port of the server. The filter compares the fixed
	// IPs and the floating IPs of the port.
	IPAddress BaremetalServerListParamsIPAddress `query:"ip_address,omitzero" json:"-"`
	// Filter by manager type, e.g. `k8s_cluster`.
	//
	// Any of "k8s_cluster".
	ManagedBy BaremetalServerListParamsManagedBy `query:"managed_by,omitzero" json:"-"`
	// Filter by managing resource ID(s).
	ManagedByResourceIDs []string `query:"managed_by_resource_ids,omitzero" json:"-"`
	// Filter by managing resource name(s).
	ManagedByResourceNames []string `query:"managed_by_resource_names,omitzero" json:"-"`
	// Operator name filter, e.g. name[contains]=web.
	Name BaremetalServerListParamsName `query:"name,omitzero" json:"-"`
	// Filter by network ID(s) the server has a port on.
	NetworkIDs []string `query:"network_ids,omitzero" format:"uuid4" json:"-"`
	// Sort field and direction.
	//
	// Any of "created_at.asc", "created_at.desc", "name.asc", "name.desc",
	// "region_id.asc", "region_id.desc", "status.asc", "status.desc".
	OrderBy BaremetalServerListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Filter by region ID(s). Omit = all regions.
	RegionIDs []int64 `query:"region_ids,omitzero" json:"-"`
	// Filter by lifecycle status(es).
	//
	// Any of "active", "deleted", "deleting", "error", "provisioning", "rebooting",
	// "rebuilding", "resizing", "stopped", "unknown", "updating".
	Statuses []string `query:"statuses,omitzero" json:"-"`
	// Filter by subnet ID(s).
	SubnetIDs []string `query:"subnet_ids,omitzero" format:"uuid4" json:"-"`
	// Filter by tag key regardless of value, e.g. `tag_key[contains]=team`.
	TagKey BaremetalServerListParamsTagKey `query:"tag_key,omitzero" json:"-"`
	// Filter by tag value regardless of key, e.g. `tag_value[prefix]=prod`.
	TagValue BaremetalServerListParamsTagValue `query:"tag_value,omitzero" json:"-"`
	// Filter by exact tag key-value pairs, e.g. `tags[env]=prod&tags[team]=core`.
	// Repeat a key to match any of several values for it, e.g.
	// `tags[env]=prod&tags[env]=dev`. Values for one key are ORed, different keys are
	// ANDed; values match case-insensitively.
	Tags map[string][]string `query:"tags,omitzero" json:"-"`
	// Range filter on last-update time.
	UpdatedAt BaremetalServerListParamsUpdatedAt `query:"updated_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalServerListParams]'s query parameters as
// `url.Values`.
func (r BaremetalServerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Range filter on creation time.
type BaremetalServerListParamsCreatedAt struct {
	// Strictly after this timestamp, e.g. `[gt]=2026-01-01T00:00:00Z`.
	Gt param.Opt[time.Time] `query:"gt,omitzero" format:"date-time" json:"-"`
	// At or after this timestamp (inclusive), e.g. `[gte]=2026-01-01T00:00:00Z`.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// Strictly before this timestamp, e.g. `[lt]=2026-02-01T00:00:00Z`.
	Lt param.Opt[time.Time] `query:"lt,omitzero" format:"date-time" json:"-"`
	// At or before this timestamp (inclusive), e.g. `[lte]=2026-02-01T00:00:00Z`.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalServerListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r BaremetalServerListParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Operator flavor name filter.
type BaremetalServerListParamsFlavorName struct {
	// Case-insensitive substring, e.g. `[contains]=web`. Repeat the key to match any
	// substring.
	Contains []string `query:"contains,omitzero" json:"-"`
	// Case-insensitive exact match, e.g. `[exact]=web-1`. Repeat the key to match any
	// of several.
	Exact []string `query:"exact,omitzero" json:"-"`
	// Case-insensitive starts-with, e.g. `[prefix]=prod-`. Repeat the key to match any
	// prefix.
	Prefix []string `query:"prefix,omitzero" json:"-"`
	// Case-insensitive ends-with, e.g. `[suffix]=-db`. Repeat the key to match any
	// suffix.
	Suffix []string `query:"suffix,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalServerListParamsFlavorName]'s query parameters as
// `url.Values`.
func (r BaremetalServerListParamsFlavorName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by an IP address of a port of the server. The filter compares the fixed
// IPs and the floating IPs of the port.
type BaremetalServerListParamsIPAddress struct {
	// Substring match on IP address, e.g. `[contains]=192.168`. Repeat the key to
	// match any substring.
	Contains []string `query:"contains,omitzero" json:"-"`
	// Exact IP match, e.g. `[exact]=192.168.1.10`. Repeat the key to match any of
	// several.
	Exact []string `query:"exact,omitzero" format:"ipvanyaddress" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalServerListParamsIPAddress]'s query parameters as
// `url.Values`.
func (r BaremetalServerListParamsIPAddress) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by manager type, e.g. `k8s_cluster`.
type BaremetalServerListParamsManagedBy string

const (
	BaremetalServerListParamsManagedByK8SCluster BaremetalServerListParamsManagedBy = "k8s_cluster"
)

// Operator name filter, e.g. name[contains]=web.
type BaremetalServerListParamsName struct {
	// Case-insensitive substring, e.g. `[contains]=web`. Repeat the key to match any
	// substring.
	Contains []string `query:"contains,omitzero" json:"-"`
	// Case-insensitive exact match, e.g. `[exact]=web-1`. Repeat the key to match any
	// of several.
	Exact []string `query:"exact,omitzero" json:"-"`
	// Case-insensitive starts-with, e.g. `[prefix]=prod-`. Repeat the key to match any
	// prefix.
	Prefix []string `query:"prefix,omitzero" json:"-"`
	// Case-insensitive ends-with, e.g. `[suffix]=-db`. Repeat the key to match any
	// suffix.
	Suffix []string `query:"suffix,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalServerListParamsName]'s query parameters as
// `url.Values`.
func (r BaremetalServerListParamsName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort field and direction.
type BaremetalServerListParamsOrderBy string

const (
	BaremetalServerListParamsOrderByCreatedAtAsc  BaremetalServerListParamsOrderBy = "created_at.asc"
	BaremetalServerListParamsOrderByCreatedAtDesc BaremetalServerListParamsOrderBy = "created_at.desc"
	BaremetalServerListParamsOrderByNameAsc       BaremetalServerListParamsOrderBy = "name.asc"
	BaremetalServerListParamsOrderByNameDesc      BaremetalServerListParamsOrderBy = "name.desc"
	BaremetalServerListParamsOrderByRegionIDAsc   BaremetalServerListParamsOrderBy = "region_id.asc"
	BaremetalServerListParamsOrderByRegionIDDesc  BaremetalServerListParamsOrderBy = "region_id.desc"
	BaremetalServerListParamsOrderByStatusAsc     BaremetalServerListParamsOrderBy = "status.asc"
	BaremetalServerListParamsOrderByStatusDesc    BaremetalServerListParamsOrderBy = "status.desc"
)

// Filter by tag key regardless of value, e.g. `tag_key[contains]=team`.
type BaremetalServerListParamsTagKey struct {
	// Case-insensitive substring, e.g. `[contains]=web`. Repeat the key to match any
	// substring.
	Contains []string `query:"contains,omitzero" json:"-"`
	// Case-insensitive exact match, e.g. `[exact]=web-1`. Repeat the key to match any
	// of several.
	Exact []string `query:"exact,omitzero" json:"-"`
	// Case-insensitive starts-with, e.g. `[prefix]=prod-`. Repeat the key to match any
	// prefix.
	Prefix []string `query:"prefix,omitzero" json:"-"`
	// Case-insensitive ends-with, e.g. `[suffix]=-db`. Repeat the key to match any
	// suffix.
	Suffix []string `query:"suffix,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalServerListParamsTagKey]'s query parameters as
// `url.Values`.
func (r BaremetalServerListParamsTagKey) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by tag value regardless of key, e.g. `tag_value[prefix]=prod`.
type BaremetalServerListParamsTagValue struct {
	// Case-insensitive substring, e.g. `[contains]=web`. Repeat the key to match any
	// substring.
	Contains []string `query:"contains,omitzero" json:"-"`
	// Case-insensitive exact match, e.g. `[exact]=web-1`. Repeat the key to match any
	// of several.
	Exact []string `query:"exact,omitzero" json:"-"`
	// Case-insensitive starts-with, e.g. `[prefix]=prod-`. Repeat the key to match any
	// prefix.
	Prefix []string `query:"prefix,omitzero" json:"-"`
	// Case-insensitive ends-with, e.g. `[suffix]=-db`. Repeat the key to match any
	// suffix.
	Suffix []string `query:"suffix,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalServerListParamsTagValue]'s query parameters as
// `url.Values`.
func (r BaremetalServerListParamsTagValue) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Range filter on last-update time.
type BaremetalServerListParamsUpdatedAt struct {
	// Strictly after this timestamp, e.g. `[gt]=2026-01-01T00:00:00Z`.
	Gt param.Opt[time.Time] `query:"gt,omitzero" format:"date-time" json:"-"`
	// At or after this timestamp (inclusive), e.g. `[gte]=2026-01-01T00:00:00Z`.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// Strictly before this timestamp, e.g. `[lt]=2026-02-01T00:00:00Z`.
	Lt param.Opt[time.Time] `query:"lt,omitzero" format:"date-time" json:"-"`
	// At or before this timestamp (inclusive), e.g. `[lte]=2026-02-01T00:00:00Z`.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalServerListParamsUpdatedAt]'s query parameters as
// `url.Values`.
func (r BaremetalServerListParamsUpdatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BaremetalServerDeleteParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Delete all floating IPs associated with the server's ports. Default preserves
	// them (disassociate only). Mutually exclusive with `floating_ip_ids`.
	AllFloatingIPs param.Opt[bool] `query:"all_floating_ips,omitzero" json:"-"`
	// Delete all ports attached to the bare metal server. Default preserves them
	// (detach only; the trunk is dissolved but every port survives, reusable
	// elsewhere). Mutually exclusive with `port_ids`.
	AllPorts param.Opt[bool] `query:"all_ports,omitzero" json:"-"`
	// Run every check of the real request and change nothing. The response has the
	// status code and the schema of the real call, and an empty list of tasks. With
	// `false`, the default, the request does the operation.
	DryRun param.Opt[bool] `query:"dry_run,omitzero" json:"-"`
	// Floating IP IDs to delete; the rest are preserved. Mutually exclusive with
	// `all_floating_ips`.
	FloatingIPIDs []string `query:"floating_ip_ids,omitzero" format:"uuid4" json:"-"`
	// Port IDs to delete; the rest are preserved. Mutually exclusive with `all_ports`.
	PortIDs []string `query:"port_ids,omitzero" format:"uuid4" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalServerDeleteParams]'s query parameters as
// `url.Values`.
func (r BaremetalServerDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BaremetalServerActionParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Power action to run on the bare metal server.
	//
	// Any of "reboot", "reboot_hard", "start", "stop".
	Action BaremetalServerActionParamsAction `json:"action,omitzero" api:"required"`
	// Run every check of the real request and change nothing. The response has the
	// status code and the schema of the real call, and an empty list of tasks. With
	// `false`, the default, the request does the operation.
	DryRun param.Opt[bool] `query:"dry_run,omitzero" json:"-"`
	paramObj
}

func (r BaremetalServerActionParams) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerActionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalServerActionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BaremetalServerActionParams]'s query parameters as
// `url.Values`.
func (r BaremetalServerActionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Power action to run on the bare metal server.
type BaremetalServerActionParamsAction string

const (
	BaremetalServerActionParamsActionReboot     BaremetalServerActionParamsAction = "reboot"
	BaremetalServerActionParamsActionRebootHard BaremetalServerActionParamsAction = "reboot_hard"
	BaremetalServerActionParamsActionStart      BaremetalServerActionParamsAction = "start"
	BaremetalServerActionParamsActionStop       BaremetalServerActionParamsAction = "stop"
)

type BaremetalServerGetParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	paramObj
}

type BaremetalServerRebuildParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Image ID
	ImageID param.Opt[string] `json:"image_id,omitzero"`
	// String in base64 format. Must not be passed together with 'username' or
	// 'password'. Examples of the `user_data`:
	// https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Opt[string] `json:"user_data,omitzero"`
	paramObj
}

func (r BaremetalServerRebuildParams) MarshalJSON() (data []byte, err error) {
	type shadow BaremetalServerRebuildParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaremetalServerRebuildParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
