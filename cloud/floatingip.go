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

// A floating IP is a static IP address that points to one of your Instances. It
// allows you to redirect network traffic to any of your Instances in the same
// datacenter.
//
// FloatingIPService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFloatingIPService] method instead.
type FloatingIPService struct {
	Options []option.RequestOption
	tasks   TaskService
}

// NewFloatingIPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFloatingIPService(opts ...option.RequestOption) (r FloatingIPService) {
	r = FloatingIPService{}
	r.Options = opts
	r.tasks = NewTaskService(opts...)
	return
}

// Create floating IP
func (r *FloatingIPService) New(ctx context.Context, params FloatingIPNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// NewAndPoll creates a floating IP and waits for the operation to complete.
func (r *FloatingIPService) NewAndPoll(ctx context.Context, params FloatingIPNewParams, opts ...option.RequestOption) (res *FloatingIP, err error) {
	// Exclude WithResponseBodyInto for the action (New returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.New(ctx, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams FloatingIPGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	task, err := r.tasks.Poll(ctx, taskID, pollOpts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Floatingips) != 1 {
		return nil, errors.New("expected exactly one floating IP to be created in a task")
	}
	resourceID := task.CreatedResources.Floatingips[0]

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, resourceID, getParams, getOpts...)
}

// This endpoint updates the association and tags of an existing Floating IP. The
// behavior depends on the current association state and the provided fields:
//
// Parameters:
//
// `port_id`: The unique identifier of the network interface (port) to which the
// Floating IP should be assigned. This ID can be retrieved from the "Get instance"
// or "List network interfaces" endpoints.
//
// `fixed_ip_address`: The private IP address assigned to the network interface.
// This must be one of the IP addresses currently assigned to the specified port.
// You can retrieve available fixed IP addresses from the "Get instance" or "List
// network interfaces" endpoints.
//
// When the Floating IP has no port associated (`port_id` is null):
//
//   - Patch with both `port_id` and `fixed_ip_address`: Assign the Floating IP to
//     the specified port and the provided `fixed_ip_address`, if that
//     `fixed_ip_address` exists on the port and is not yet used by another Floating
//     IP.
//   - Patch with `port_id` only (`fixed_ip_address` omitted): Assign the Floating IP
//     to the specified port using the first available IPv4 fixed IP of that port.
//
// When the Floating IP is already associated with a port:
//
//   - Patch with both `port_id` and `fixed_ip_address`: Re-assign the Floating IP to
//     the specified port and address if all validations pass.
//   - Patch with `port_id` only (`fixed_ip_address` omitted): Re-assign the Floating
//     IP to the specified port using the first available IPv4 fixed IP of that port.
//   - Patch with `port_id` = null: Unassign the Floating IP from its current port.
//
// Tags:
//
//   - You can update tags alongside association changes. Tags are provided as a list
//     of key-value pairs.
//
// Idempotency and task creation:
//
//   - No worker task is created if the requested state is already actual, i.e., the
//     requested `port_id` equals the current `port_id` and/or the requested
//     `fixed_ip_address` equals the current `fixed_ip_address`, and the tags already
//     match the current tags. In such cases, the endpoint returns an empty tasks
//     list.
func (r *FloatingIPService) Update(ctx context.Context, floatingIPID string, params FloatingIPUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if floatingIPID == "" {
		err = errors.New("missing required floating_ip_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/floatingips/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, floatingIPID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// UpdateAndPoll updates a floating IP and polls for completion of the first task. Use the [TaskService.Poll] method if
// you need to poll for all tasks.
func (r *FloatingIPService) UpdateAndPoll(ctx context.Context, floatingIPID string, params FloatingIPUpdateParams, opts ...option.RequestOption) (res *FloatingIP, err error) {
	// Exclude WithResponseBodyInto for the action (Update returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Update(ctx, floatingIPID, params, actionOpts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams FloatingIPGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) == 0 {
		return nil, errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	_, err = r.tasks.Poll(ctx, taskID, pollOpts...)
	if err != nil {
		return nil, err
	}

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, floatingIPID, getParams, getOpts...)
}

// List floating IPs
func (r *FloatingIPService) List(ctx context.Context, params FloatingIPListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[FloatingIPDetailed], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v", params.ProjectID.Value, params.RegionID.Value)
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

// List floating IPs
func (r *FloatingIPService) ListAutoPaging(ctx context.Context, params FloatingIPListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[FloatingIPDetailed] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete floating IP
func (r *FloatingIPService) Delete(ctx context.Context, floatingIPID string, body FloatingIPDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if floatingIPID == "" {
		err = errors.New("missing required floating_ip_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, floatingIPID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// DeleteAndPoll deletes a floating IP and polls for completion of the first task. Use the
// [TaskService.Poll] method if you need to poll for all tasks.
func (r *FloatingIPService) DeleteAndPoll(ctx context.Context, floatingIPID string, body FloatingIPDeleteParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (Delete returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Delete(ctx, floatingIPID, body, actionOpts...)
	if err != nil {
		return err
	}

	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	_, err = r.tasks.Poll(ctx, taskID, pollOpts...)
	return err
}

// Assign floating IP to instance or loadbalancer.
//
// **Deprecated**: Use
// `PATCH /v2/floatingips/{project_id}/{region_id}/{floating_ip_id}` instead.
//
// Deprecated: deprecated
func (r *FloatingIPService) Assign(ctx context.Context, floatingIPID string, params FloatingIPAssignParams, opts ...option.RequestOption) (res *FloatingIP, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if floatingIPID == "" {
		err = errors.New("missing required floating_ip_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s/assign", params.ProjectID.Value, params.RegionID.Value, floatingIPID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get floating IP
func (r *FloatingIPService) Get(ctx context.Context, floatingIPID string, query FloatingIPGetParams, opts ...option.RequestOption) (res *FloatingIP, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if floatingIPID == "" {
		err = errors.New("missing required floating_ip_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, floatingIPID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// **Deprecated**: Use
// `PATCH /v2/floatingips/{project_id}/{region_id}/{floating_ip_id}` instead.
//
// Deprecated: deprecated
func (r *FloatingIPService) Unassign(ctx context.Context, floatingIPID string, body FloatingIPUnassignParams, opts ...option.RequestOption) (res *FloatingIP, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if floatingIPID == "" {
		err = errors.New("missing required floating_ip_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s/unassign", body.ProjectID.Value, body.RegionID.Value, floatingIPID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type FloatingIPDetailed struct {
	// Floating IP ID
	ID string `json:"id" api:"required" format:"uuid4"`
	// Datetime when the floating IP was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id" api:"required" format:"uuid4"`
	// IP address of the port the floating IP is attached to
	FixedIPAddress string `json:"fixed_ip_address" api:"required" format:"ipvanyaddress"`
	// IP Address of the floating IP
	FloatingIPAddress string `json:"floating_ip_address" api:"required" format:"ipvanyaddress"`
	// Instance the floating IP is attached to
	Instance FloatingIPDetailedInstance `json:"instance" api:"required"`
	// Load balancer the floating IP is attached to
	Loadbalancer LoadBalancer `json:"loadbalancer" api:"required"`
	// Port ID
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	// Project ID
	ProjectID int64 `json:"project_id" api:"required"`
	// Region name
	Region string `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// Router ID
	RouterID string `json:"router_id" api:"required" format:"uuid4"`
	// Floating IP status. DOWN - unassigned (available). ACTIVE - attached to a port
	// (in use). ERROR - error state.
	//
	// Any of "ACTIVE", "DOWN", "ERROR".
	Status FloatingIPStatus `json:"status" api:"required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags" api:"required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id" api:"required" format:"uuid4"`
	// Datetime when the floating IP was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		CreatorTaskID     respjson.Field
		FixedIPAddress    respjson.Field
		FloatingIPAddress respjson.Field
		Instance          respjson.Field
		Loadbalancer      respjson.Field
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
func (r FloatingIPDetailed) RawJSON() string { return r.JSON.raw }
func (r *FloatingIPDetailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance the floating IP is attached to
type FloatingIPDetailedInstance struct {
	// Instance ID
	ID string `json:"id" api:"required" format:"uuid4"`
	// Map of `network_name` to list of addresses in that network
	Addresses map[string][]FloatingIPDetailedInstanceAddressUnion `json:"addresses" api:"required"`
	// Datetime when instance was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id" api:"required"`
	// Flavor
	Flavor FloatingIPDetailedInstanceFlavor `json:"flavor" api:"required"`
	// Instance description
	InstanceDescription string `json:"instance_description" api:"required"`
	// Instance name
	Name string `json:"name" api:"required"`
	// Project ID
	ProjectID int64 `json:"project_id" api:"required"`
	// Region name
	Region string `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// Security groups
	SecurityGroups []FloatingIPDetailedInstanceSecurityGroup `json:"security_groups" api:"required"`
	// SSH key name assigned to instance
	SSHKeyName string `json:"ssh_key_name" api:"required"`
	// Instance status
	//
	// Any of "ACTIVE", "BUILD", "DELETED", "ERROR", "HARD_REBOOT", "MIGRATING",
	// "PASSWORD", "PAUSED", "REBOOT", "REBUILD", "RESCUE", "RESIZE", "REVERT_RESIZE",
	// "SHELVED", "SHELVED_OFFLOADED", "SHUTOFF", "SOFT_DELETED", "SUSPENDED",
	// "UNKNOWN", "VERIFY_RESIZE".
	Status string `json:"status" api:"required"`
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
	// Task state
	TaskState string `json:"task_state" api:"required"`
	// Virtual machine state (active)
	//
	// Any of "active", "building", "deleted", "error", "paused", "rescued", "resized",
	// "shelved", "shelved_offloaded", "soft-deleted", "stopped", "suspended".
	VmState string `json:"vm_state" api:"required"`
	// List of volumes
	Volumes []FloatingIPDetailedInstanceVolume `json:"volumes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Addresses           respjson.Field
		CreatedAt           respjson.Field
		CreatorTaskID       respjson.Field
		Flavor              respjson.Field
		InstanceDescription respjson.Field
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
func (r FloatingIPDetailedInstance) RawJSON() string { return r.JSON.raw }
func (r *FloatingIPDetailedInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FloatingIPDetailedInstanceAddressUnion contains all possible properties and
// values from [FloatingAddress], [FixedAddressShort], [FixedAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FloatingIPDetailedInstanceAddressUnion struct {
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

func (u FloatingIPDetailedInstanceAddressUnion) AsFloatingIPAddress() (v FloatingAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FloatingIPDetailedInstanceAddressUnion) AsFixedIPAddressShort() (v FixedAddressShort) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FloatingIPDetailedInstanceAddressUnion) AsFixedIPAddress() (v FixedAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FloatingIPDetailedInstanceAddressUnion) RawJSON() string { return u.JSON.raw }

func (r *FloatingIPDetailedInstanceAddressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor
type FloatingIPDetailedInstanceFlavor struct {
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id" api:"required"`
	// Flavor name
	FlavorName string `json:"flavor_name" api:"required"`
	// RAM size in MiB
	Ram int64 `json:"ram" api:"required"`
	// Virtual CPU count. For bare metal flavors, it's a physical CPU count
	Vcpus int64 `json:"vcpus" api:"required"`
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
func (r FloatingIPDetailedInstanceFlavor) RawJSON() string { return r.JSON.raw }
func (r *FloatingIPDetailedInstanceFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIPDetailedInstanceSecurityGroup struct {
	// Name.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatingIPDetailedInstanceSecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *FloatingIPDetailedInstanceSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIPDetailedInstanceVolume struct {
	// Volume ID
	ID string `json:"id" api:"required"`
	// Whether the volume is deleted together with the VM
	DeleteOnTermination bool `json:"delete_on_termination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		DeleteOnTermination respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FloatingIPDetailedInstanceVolume) RawJSON() string { return r.JSON.raw }
func (r *FloatingIPDetailedInstanceVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIPNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// If the port has multiple IP addresses, a specific one can be selected using this
	// field. If not specified, the first IP in the port's list will be used by
	// default.
	FixedIPAddress param.Opt[string] `json:"fixed_ip_address,omitzero" format:"ipv4"`
	// If provided, the floating IP will be immediately attached to the specified port.
	PortID param.Opt[string] `json:"port_id,omitzero" format:"uuid4"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Some tags are read-only and cannot be modified by the
	// user. Tags are also integrated with cost reports, allowing cost data to be
	// filtered based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r FloatingIPNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FloatingIPNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FloatingIPNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIPUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Fixed IP address
	FixedIPAddress param.Opt[string] `json:"fixed_ip_address,omitzero" format:"ipv4"`
	// Port ID
	PortID param.Opt[string] `json:"port_id,omitzero" format:"uuid4"`
	// Update key-value tags using JSON Merge Patch semantics (RFC 7386). Provide
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

func (r FloatingIPUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FloatingIPUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FloatingIPUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIPListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Optional. Filter by tag key-value pairs.
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// Filter by floating IP status. DOWN - unassigned (available). ACTIVE - attached
	// to a port (in use). ERROR - error state.
	//
	// Any of "ACTIVE", "DOWN", "ERROR".
	Status FloatingIPStatus `query:"status,omitzero" json:"-"`
	// Optional. Filter by tag keys. ?`tag_key`=key1&`tag_key`=key2
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FloatingIPListParams]'s query parameters as `url.Values`.
func (r FloatingIPListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FloatingIPDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type FloatingIPAssignParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Port ID
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	// Fixed IP address
	FixedIPAddress param.Opt[string] `json:"fixed_ip_address,omitzero" format:"ipvanyaddress"`
	paramObj
}

func (r FloatingIPAssignParams) MarshalJSON() (data []byte, err error) {
	type shadow FloatingIPAssignParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FloatingIPAssignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIPGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type FloatingIPUnassignParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}
