// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
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

// Instances are cloud virtual machines that boot from a volume and connect to
// networks through ports, in any region of the project.
//
// InstanceService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceService] method instead.
type InstanceService struct {
	Options       []option.RequestOption
	Console       InstanceConsoleService
	ResizeFlavors InstanceResizeFlavorService
	// Instance flavors define the CPU, memory, and OS type of an instance, in all
	// regions of the project.
	Flavors InstanceFlavorService
	// Instance images are operating system images (public, private, or shared) used to
	// boot cloud instances.
	Images  InstanceImageService
	Metrics InstanceMetricService
}

// NewInstanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceService(opts ...option.RequestOption) (r InstanceService) {
	r = InstanceService{}
	r.Options = opts
	r.Console = NewInstanceConsoleService(opts...)
	r.ResizeFlavors = NewInstanceResizeFlavorService(opts...)
	r.Flavors = NewInstanceFlavorService(opts...)
	r.Images = NewInstanceImageService(opts...)
	r.Metrics = NewInstanceMetricService(opts...)
	return
}

// Create a virtual instance that boots from a volume. The instance uses a volume
// and ports that you create first, in the region of the instance.
//
// To create one:
//
//  1. Pick an image with [List images](#operation/ImagesViewSet.get).
//  2. [Create a volume](#operation/VolumeViewSet.post) from the image. Wait until
//     its status is `available`.
//  3. Pick a flavor with
//     [List instance flavors](#operation/GlobalInstanceFlavorListView.get). Send
//     `compatible_with_boot_volume_id` to get only the flavors that can boot the
//     volume.
//  4. [Create a port](#operation/GlobalPortListView.post). A body with only
//     `region_id` gives one public IPv4 address.
//
// You can also attach more volumes, put the instance in a
// [placement group](#operation/GlobalPlacementGroupListView.post), and add an
// [SSH key](/docs/api-reference/cloud/ssh-keys/add-or-generate-ssh-key).
//
// The response holds the IDs of the tasks that create the instance.
//
// With `?dry_run=true` the request runs every check and changes nothing. The
// response holds an empty `tasks` array. A failed check returns the same error as
// the real create.
func (r *InstanceService) New(ctx context.Context, params InstanceNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/projects/%v/instances", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Send only the fields to change. One task makes all the changes. A field with the
// current value changes nothing.
//
//   - Rename: send `name`.
//   - Resize: send `flavor_name`. Pick it with
//     [List instance resize flavors](#operation/GlobalInstanceResizeFlavorListView.get).
//     A resize restarts the instance.
//   - Add a port: [create a port](#operation/GlobalPortListView.post). Then send
//     `interfaces` with the current ports in their current order and the new port at
//     the end.
//   - Remove a port: send `interfaces` without that port. The port stays and you can
//     use it again. For a port change, the instance `status` must be `active` or
//     `stopped`. Resume a suspended instance first.
//   - Attach or detach a data volume: send the full list in `data_volume_ids`. A new
//     volume must be `available`.
//   - Move to a [placement group](#operation/GlobalPlacementGroupListView.post):
//     send `placement_group_id`. Send `null` to remove the instance from its group.
//     The instance `status` must be `active`. The instance can move to another host.
//   - Change tags: send `tags` as a JSON Merge Patch.
//
// To change the security groups,
// [update the ports](#operation/GlobalPortDetailView.patch). You cannot change
// `boot_volume_id` or `ssh_key_id`. To change them, create a new instance.
//
// With `?dry_run=true` the request runs every check and changes nothing. The
// response holds an empty `tasks` array. A failed check returns the same error as
// the real update. If a resize exceeds a quota, the error names every quota that
// it exceeds.
func (r *InstanceService) Update(ctx context.Context, instanceID string, params InstanceUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/instances/%s", params.ProjectID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List a project's instances across all regions.
func (r *InstanceService) List(ctx context.Context, params InstanceListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Instance], err error) {
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
	path := fmt.Sprintf("cloud/v1/projects/%v/instances", params.ProjectID.Value)
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

// List a project's instances across all regions.
func (r *InstanceService) ListAutoPaging(ctx context.Context, params InstanceListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Instance] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Cascade-delete volumes, ports, floating IPs based on preservation params. By
// default, all are preserved (detached/disassociated).
//
// `resources_to_delete` holds:
//
//   - the instance itself;
//   - the volumes that `volume_ids` names, or every volume if `all_volumes` is
//     `true`, and the snapshots of these volumes;
//   - the ports that `port_ids` names, or every port if `all_ports` is `true`, and
//     the sub-ports of their trunks;
//   - the floating IPs that `floating_ip_ids` names, or every floating IP if
//     `all_floating_ips` is `true`.
//
// `resources_to_preserve` holds every other volume, port, sub-port and floating IP
// of the instance. A floating IP on a deleted port stays, with no port.
//
// [`DELETE /v1/instances/{project_id}/{region_id}/{instance_id}`](#operation/InstanceInstanceViewSet.delete)
// is different. It deletes every port, except the reserved fixed IPs that
// `reserved_fixed_ips` does not name.
//
// With `?dry_run=true` the request runs every check and changes nothing. The
// response holds an empty `tasks` array and the same two lists. A failed check
// returns the same error as the real delete.
func (r *InstanceService) Delete(ctx context.Context, instanceID string, params InstanceDeleteParams, opts ...option.RequestOption) (res *InstanceDeleteResponse, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/instances/%s", params.ProjectID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Change the power state of an instance. The action is one of: `start`, `stop`,
// `reboot`, `reboot_hard`, `suspend` or `resume`. The action does not change the
// Advanced DDoS profile of the instance. Use the DDoS profile resource to activate
// or to deactivate protection. Returns the ID of the task that runs the action.
//
// With `?dry_run=true` the request runs every check and changes nothing. The
// response holds an empty `tasks` array. A failed check returns the same error as
// the real action.
func (r *InstanceService) Action(ctx context.Context, instanceID string, params InstanceActionParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/instances/%s/action", params.ProjectID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a single instance by ID, resolved across all regions.
func (r *InstanceService) Get(ctx context.Context, instanceID string, query InstanceGetParams, opts ...option.RequestOption) (res *Instance, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/instances/%s", query.ProjectID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Rebuild instance with a new image while preserving its configuration.
func (r *InstanceService) Rebuild(ctx context.Context, instanceID string, params InstanceRebuildParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v3/instances/%v/%v/%s/rebuild", params.ProjectID.Value, params.RegionID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type Instance struct {
	// Instance ID.
	ID string `json:"id" api:"required" format:"uuid4"`
	// Boot volume ID, or null for an image-backed instance.
	BootVolumeID string `json:"boot_volume_id" api:"required" format:"uuid4"`
	// Datetime when the instance was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Non-boot attached volume IDs.
	DataVolumeIDs []string `json:"data_volume_ids" api:"required" format:"uuid4"`
	// Flavor name.
	FlavorName string `json:"flavor_name" api:"required"`
	// Attached network interfaces.
	Interfaces []InstanceInterface `json:"interfaces" api:"required"`
	// Service managing this instance, or null if customer-operated.
	ManagedBy InstanceManagedBy `json:"managed_by" api:"required"`
	// Instance name.
	Name string `json:"name" api:"required"`
	// Placement group this instance belongs to, or null.
	PlacementGroupID string `json:"placement_group_id" api:"required" format:"uuid4"`
	// Region ID.
	RegionID int64 `json:"region_id" api:"required"`
	// ID of the SSH key of the default user, or null.
	SSHKeyID string `json:"ssh_key_id" api:"required" format:"uuid4"`
	// Canonical lifecycle status.
	//
	// Any of "active", "deleted", "deleting", "error", "provisioning", "rebooting",
	// "rebuilding", "resizing", "stopped", "unknown", "updating".
	Status InstanceStatus `json:"status" api:"required"`
	// Read-only key-value tags set by the platform.
	SystemTags map[string]string `json:"system_tags" api:"required"`
	// Customer-owned key-value tags.
	Tags map[string]string `json:"tags" api:"required"`
	// Datetime when the instance was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		BootVolumeID     respjson.Field
		CreatedAt        respjson.Field
		DataVolumeIDs    respjson.Field
		FlavorName       respjson.Field
		Interfaces       respjson.Field
		ManagedBy        respjson.Field
		Name             respjson.Field
		PlacementGroupID respjson.Field
		RegionID         respjson.Field
		SSHKeyID         respjson.Field
		Status           respjson.Field
		SystemTags       respjson.Field
		Tags             respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Instance) RawJSON() string { return r.JSON.raw }
func (r *Instance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceInterface struct {
	// Fixed IPs assigned to the port, each with its floating IP.
	IPAssignments []InstanceInterfaceIPAssignment `json:"ip_assignments" api:"required"`
	// Port MAC address.
	MacAddress string `json:"mac_address" api:"required"`
	// Network the port is attached to.
	NetworkID string `json:"network_id" api:"required" format:"uuid4"`
	// ID of the port that connects the server to the network.
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAssignments respjson.Field
		MacAddress    respjson.Field
		NetworkID     respjson.Field
		PortID        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstanceInterface) RawJSON() string { return r.JSON.raw }
func (r *InstanceInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceInterfaceIPAssignment struct {
	// Floating IP that translates this address. `null` when the address has no
	// floating IP. An IPv6 address always has `null`.
	FloatingIP InstanceInterfaceIPAssignmentFloatingIP `json:"floating_ip" api:"required"`
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
func (r InstanceInterfaceIPAssignment) RawJSON() string { return r.JSON.raw }
func (r *InstanceInterfaceIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Floating IP that translates this address. `null` when the address has no
// floating IP. An IPv6 address always has `null`.
type InstanceInterfaceIPAssignmentFloatingIP struct {
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
func (r InstanceInterfaceIPAssignmentFloatingIP) RawJSON() string { return r.JSON.raw }
func (r *InstanceInterfaceIPAssignmentFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Service managing this instance, or null if customer-operated.
type InstanceManagedBy struct {
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
func (r InstanceManagedBy) RawJSON() string { return r.JSON.raw }
func (r *InstanceManagedBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Canonical lifecycle status.
type InstanceStatus string

const (
	InstanceStatusActive       InstanceStatus = "active"
	InstanceStatusDeleted      InstanceStatus = "deleted"
	InstanceStatusDeleting     InstanceStatus = "deleting"
	InstanceStatusError        InstanceStatus = "error"
	InstanceStatusProvisioning InstanceStatus = "provisioning"
	InstanceStatusRebooting    InstanceStatus = "rebooting"
	InstanceStatusRebuilding   InstanceStatus = "rebuilding"
	InstanceStatusResizing     InstanceStatus = "resizing"
	InstanceStatusStopped      InstanceStatus = "stopped"
	InstanceStatusUnknown      InstanceStatus = "unknown"
	InstanceStatusUpdating     InstanceStatus = "updating"
)

type InstanceDeleteResponse struct {
	// Every resource that this request deletes: the instance, and each volume,
	// snapshot, port and floating IP that goes with it. Sorted by `type`, then by
	// `id`.
	ResourcesToDelete []InstanceDeleteResponseResourcesToDelete `json:"resources_to_delete" api:"required"`
	// The volumes, ports and floating IPs of the instance that stay after the delete.
	// Sorted by `type`, then by `id`.
	ResourcesToPreserve []InstanceDeleteResponseResourcesToPreserve `json:"resources_to_preserve" api:"required"`
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
func (r InstanceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *InstanceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceDeleteResponseResourcesToDelete struct {
	// The ID of the resource.
	ID string `json:"id" api:"required"`
	// The name of the resource. For a floating IP, the IP address. `null` for a
	// volume, and for a resource without a name.
	Name string `json:"name" api:"required"`
	// The type of the resource.
	//
	// Any of "floating_ips", "instances", "ports", "volume_snapshots", "volumes".
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
func (r InstanceDeleteResponseResourcesToDelete) RawJSON() string { return r.JSON.raw }
func (r *InstanceDeleteResponseResourcesToDelete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceDeleteResponseResourcesToPreserve struct {
	// The ID of the resource.
	ID string `json:"id" api:"required"`
	// The name of the resource. For a floating IP, the IP address. `null` for a
	// volume, and for a resource without a name.
	Name string `json:"name" api:"required"`
	// The type of the resource.
	//
	// Any of "floating_ips", "instances", "ports", "volume_snapshots", "volumes".
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
func (r InstanceDeleteResponseResourcesToPreserve) RawJSON() string { return r.JSON.raw }
func (r *InstanceDeleteResponseResourcesToPreserve) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceNewParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Pre-created boot volume UUID.
	BootVolumeID string `json:"boot_volume_id" api:"required" format:"uuid4"`
	// Name of an instance flavor with no GPU. Use
	// `GET /v1/projects/{project_id}/instances/flavors` to list them.
	FlavorName string `json:"flavor_name" api:"required"`
	// Pre-created ports to attach, in NIC/boot order.
	Interfaces []InstanceNewParamsInterface `json:"interfaces,omitzero" api:"required"`
	// Region to provision the instance in.
	RegionID int64 `json:"region_id" api:"required"`
	// Run every check of the real request and change nothing. The response has the
	// status code and the schema of the real call, and an empty list of tasks. With
	// `false`, the default, the request does the operation.
	DryRun param.Opt[bool] `query:"dry_run,omitzero" json:"-"`
	// Instance name. If you omit it, the name template of the account gives the name.
	// Without a template, the name is generated. The account can forbid custom names.
	// Then the name must be equal to the name that the template gives.
	Name param.Opt[string] `json:"name,omitzero"`
	// Password of the default user. No response returns it. A Windows image needs a
	// password. For a Windows image, the create checks these rules: the password has a
	// minimum of 6 characters, it contains no word 'admin', and it contains 3 of these
	// character types: digits, lowercase letters, uppercase letters, other characters.
	// For a Linux image, the password enables SSH login with a password. Send
	// `password` or `user_data`, not both.
	Password param.Opt[string] `json:"password,omitzero" format:"password"`
	// Pre-created placement group UUID for affinity/anti-affinity.
	PlacementGroupID param.Opt[string] `json:"placement_group_id,omitzero" format:"uuid4"`
	// ID of the SSH key to add to the default user.
	SSHKeyID param.Opt[string] `json:"ssh_key_id,omitzero" format:"uuid4"`
	// cloud-init user data, raw or base64. Up to 65535 bytes once decoded.
	UserData param.Opt[string] `json:"user_data,omitzero"`
	// Pre-created data volume UUIDs to attach.
	DataVolumeIDs []string `json:"data_volume_ids,omitzero" format:"uuid4"`
	// Customer-owned key-value tags.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r InstanceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [InstanceNewParams]'s query parameters as `url.Values`.
func (r InstanceNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The property PortID is required.
type InstanceNewParamsInterface struct {
	// Existing port to attach. Array order is the NIC/boot order.
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstanceNewParamsInterface) MarshalJSON() (data []byte, err error) {
	type shadow InstanceNewParamsInterface
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceNewParamsInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceUpdateParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// ID of the placement group. Send null to remove the instance from its placement
	// group.
	PlacementGroupID param.Opt[string] `json:"placement_group_id,omitzero" format:"uuid4"`
	// Run every check of the real request and change nothing. The response has the
	// status code and the schema of the real call, and an empty list of tasks. With
	// `false`, the default, the request does the operation.
	DryRun param.Opt[bool] `query:"dry_run,omitzero" json:"-"`
	// Name of the new flavor. The instance is resized to this flavor.
	FlavorName param.Opt[string] `json:"flavor_name,omitzero"`
	// Instance name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Full list of the data volumes of the instance. Volumes that are not in the list
	// are detached. New volumes are attached.
	DataVolumeIDs []string `json:"data_volume_ids,omitzero" format:"uuid4"`
	// Full list of the ports of the instance. Ports that are not in the list are
	// detached. New ports are attached. Keep the current ports in their current order.
	Interfaces []InstanceUpdateParamsInterface `json:"interfaces,omitzero"`
	// Tags as JSON Merge Patch. A key with the value null removes that tag. The value
	// null removes all tags.
	Tags TagUpdateMap `json:"tags,omitzero"`
	paramObj
}

func (r InstanceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [InstanceUpdateParams]'s query parameters as `url.Values`.
func (r InstanceUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The property PortID is required.
type InstanceUpdateParamsInterface struct {
	// Existing port to attach. Array order is the NIC/boot order.
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	paramObj
}

func (r InstanceUpdateParamsInterface) MarshalJSON() (data []byte, err error) {
	type shadow InstanceUpdateParamsInterface
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceUpdateParamsInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstanceListParams struct {
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
	CreatedAt InstanceListParamsCreatedAt `query:"created_at,omitzero" json:"-"`
	// Operator flavor name filter.
	FlavorName InstanceListParamsFlavorName `query:"flavor_name,omitzero" json:"-"`
	// Filter by instance UUID(s).
	IDs []string `query:"ids,omitzero" format:"uuid4" json:"-"`
	// Filter by an IP address of a port of the instance. The filter compares the fixed
	// IPs and the floating IPs of the port.
	IPAddress InstanceListParamsIPAddress `query:"ip_address,omitzero" json:"-"`
	// Filter by manager type, e.g. `k8s_cluster`.
	//
	// Any of "k8s_cluster".
	ManagedBy InstanceListParamsManagedBy `query:"managed_by,omitzero" json:"-"`
	// Filter by managing resource ID(s).
	ManagedByResourceIDs []string `query:"managed_by_resource_ids,omitzero" json:"-"`
	// Filter by managing resource name(s).
	ManagedByResourceNames []string `query:"managed_by_resource_names,omitzero" json:"-"`
	// Operator name filter, e.g. name[contains]=web.
	Name InstanceListParamsName `query:"name,omitzero" json:"-"`
	// Filter by network ID(s) the instance has a port on.
	NetworkIDs []string `query:"network_ids,omitzero" format:"uuid4" json:"-"`
	// Sort field and direction.
	//
	// Any of "created_at.asc", "created_at.desc", "name.asc", "name.desc",
	// "region_id.asc", "region_id.desc", "status.asc", "status.desc".
	OrderBy InstanceListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Filter by region ID(s). Omit = all regions.
	RegionIDs []int64 `query:"region_ids,omitzero" json:"-"`
	// Filter by security group ID(s).
	SecurityGroupIDs []string `query:"security_group_ids,omitzero" format:"uuid4" json:"-"`
	// Filter by lifecycle status(es).
	//
	// Any of "active", "deleted", "deleting", "error", "provisioning", "rebooting",
	// "rebuilding", "resizing", "stopped", "unknown", "updating".
	Statuses []string `query:"statuses,omitzero" json:"-"`
	// Filter by subnet ID(s).
	SubnetIDs []string `query:"subnet_ids,omitzero" format:"uuid4" json:"-"`
	// Filter by tag key regardless of value, e.g. `tag_key[contains]=team`.
	TagKey InstanceListParamsTagKey `query:"tag_key,omitzero" json:"-"`
	// Filter by tag value regardless of key, e.g. `tag_value[prefix]=prod`.
	TagValue InstanceListParamsTagValue `query:"tag_value,omitzero" json:"-"`
	// Filter by exact tag key-value pairs, e.g. `tags[env]=prod&tags[team]=core`.
	// Repeat a key to match any of several values for it, e.g.
	// `tags[env]=prod&tags[env]=dev`. Values for one key are ORed, different keys are
	// ANDed; values match case-insensitively.
	Tags map[string][]string `query:"tags,omitzero" json:"-"`
	// Range filter on last-update time.
	UpdatedAt InstanceListParamsUpdatedAt `query:"updated_at,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceListParams]'s query parameters as `url.Values`.
func (r InstanceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Range filter on creation time.
type InstanceListParamsCreatedAt struct {
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

// URLQuery serializes [InstanceListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r InstanceListParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Operator flavor name filter.
type InstanceListParamsFlavorName struct {
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

// URLQuery serializes [InstanceListParamsFlavorName]'s query parameters as
// `url.Values`.
func (r InstanceListParamsFlavorName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by an IP address of a port of the instance. The filter compares the fixed
// IPs and the floating IPs of the port.
type InstanceListParamsIPAddress struct {
	// Substring match on IP address, e.g. `[contains]=192.168`. Repeat the key to
	// match any substring.
	Contains []string `query:"contains,omitzero" json:"-"`
	// Exact IP match, e.g. `[exact]=192.168.1.10`. Repeat the key to match any of
	// several.
	Exact []string `query:"exact,omitzero" format:"ipvanyaddress" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceListParamsIPAddress]'s query parameters as
// `url.Values`.
func (r InstanceListParamsIPAddress) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by manager type, e.g. `k8s_cluster`.
type InstanceListParamsManagedBy string

const (
	InstanceListParamsManagedByK8SCluster InstanceListParamsManagedBy = "k8s_cluster"
)

// Operator name filter, e.g. name[contains]=web.
type InstanceListParamsName struct {
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

// URLQuery serializes [InstanceListParamsName]'s query parameters as `url.Values`.
func (r InstanceListParamsName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort field and direction.
type InstanceListParamsOrderBy string

const (
	InstanceListParamsOrderByCreatedAtAsc  InstanceListParamsOrderBy = "created_at.asc"
	InstanceListParamsOrderByCreatedAtDesc InstanceListParamsOrderBy = "created_at.desc"
	InstanceListParamsOrderByNameAsc       InstanceListParamsOrderBy = "name.asc"
	InstanceListParamsOrderByNameDesc      InstanceListParamsOrderBy = "name.desc"
	InstanceListParamsOrderByRegionIDAsc   InstanceListParamsOrderBy = "region_id.asc"
	InstanceListParamsOrderByRegionIDDesc  InstanceListParamsOrderBy = "region_id.desc"
	InstanceListParamsOrderByStatusAsc     InstanceListParamsOrderBy = "status.asc"
	InstanceListParamsOrderByStatusDesc    InstanceListParamsOrderBy = "status.desc"
)

// Filter by tag key regardless of value, e.g. `tag_key[contains]=team`.
type InstanceListParamsTagKey struct {
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

// URLQuery serializes [InstanceListParamsTagKey]'s query parameters as
// `url.Values`.
func (r InstanceListParamsTagKey) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by tag value regardless of key, e.g. `tag_value[prefix]=prod`.
type InstanceListParamsTagValue struct {
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

// URLQuery serializes [InstanceListParamsTagValue]'s query parameters as
// `url.Values`.
func (r InstanceListParamsTagValue) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Range filter on last-update time.
type InstanceListParamsUpdatedAt struct {
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

// URLQuery serializes [InstanceListParamsUpdatedAt]'s query parameters as
// `url.Values`.
func (r InstanceListParamsUpdatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InstanceDeleteParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Delete all associated floating IPs (default: preserve).
	AllFloatingIPs param.Opt[bool] `query:"all_floating_ips,omitzero" json:"-"`
	// Delete all attached ports (default: preserve).
	AllPorts param.Opt[bool] `query:"all_ports,omitzero" json:"-"`
	// Delete all attached volumes (default: preserve).
	AllVolumes param.Opt[bool] `query:"all_volumes,omitzero" json:"-"`
	// Run every check of the real request and change nothing. The response has the
	// status code and the schema of the real call, and an empty list of tasks. With
	// `false`, the default, the request does the operation.
	DryRun param.Opt[bool] `query:"dry_run,omitzero" json:"-"`
	// Delete only these floating IP IDs (mutually exclusive with `all_floating_ips`).
	FloatingIPIDs []string `query:"floating_ip_ids,omitzero" format:"uuid4" json:"-"`
	// Delete only these port IDs (mutually exclusive with `all_ports`).
	PortIDs []string `query:"port_ids,omitzero" format:"uuid4" json:"-"`
	// Delete only these volume IDs (mutually exclusive with `all_volumes`).
	VolumeIDs []string `query:"volume_ids,omitzero" format:"uuid4" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceDeleteParams]'s query parameters as `url.Values`.
func (r InstanceDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InstanceActionParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Power action to run on the instance. `reboot_hard` replaces the legacy
	// `powercycle` action.
	//
	// Any of "reboot", "reboot_hard", "resume", "start", "stop", "suspend".
	Action InstanceActionParamsAction `json:"action,omitzero" api:"required"`
	// Run every check of the real request and change nothing. The response has the
	// status code and the schema of the real call, and an empty list of tasks. With
	// `false`, the default, the request does the operation.
	DryRun param.Opt[bool] `query:"dry_run,omitzero" json:"-"`
	paramObj
}

func (r InstanceActionParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceActionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceActionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [InstanceActionParams]'s query parameters as `url.Values`.
func (r InstanceActionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Power action to run on the instance. `reboot_hard` replaces the legacy
// `powercycle` action.
type InstanceActionParamsAction string

const (
	InstanceActionParamsActionReboot     InstanceActionParamsAction = "reboot"
	InstanceActionParamsActionRebootHard InstanceActionParamsAction = "reboot_hard"
	InstanceActionParamsActionResume     InstanceActionParamsAction = "resume"
	InstanceActionParamsActionStart      InstanceActionParamsAction = "start"
	InstanceActionParamsActionStop       InstanceActionParamsAction = "stop"
	InstanceActionParamsActionSuspend    InstanceActionParamsAction = "suspend"
)

type InstanceGetParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	paramObj
}

type InstanceRebuildParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Specifies the name of the SSH keypair, created via the
	// [/v1/`ssh_keys` endpoint](/docs/api-reference/cloud/ssh-keys/add-or-generate-ssh-key).
	// Use null to remove the SSH key from the instance. If you omit it, the instance
	// keeps the current SSH key.
	SSHKeyName param.Opt[string] `json:"ssh_key_name,omitzero"`
	// Image ID. If you omit it, the instance keeps the current image.
	ImageID param.Opt[string] `json:"image_id,omitzero" format:"uuid4"`
	// String in base64 format. Examples of the `user_data`:
	// https://cloudinit.readthedocs.io/en/latest/topics/examples.html. If you omit it,
	// the instance keeps the current user data.
	UserData param.Opt[string] `json:"user_data,omitzero"`
	paramObj
}

func (r InstanceRebuildParams) MarshalJSON() (data []byte, err error) {
	type shadow InstanceRebuildParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstanceRebuildParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
