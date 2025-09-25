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
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// NewAndPoll creates a floating IP and waits for the operation to complete.
func (r *FloatingIPService) NewAndPoll(ctx context.Context, params FloatingIPNewParams, opts ...option.RequestOption) (res *FloatingIP, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
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
	task, err := r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Floatingips) != 1 {
		return nil, errors.New("expected exactly one network to be created in a task")
	}
	resourceID := task.CreatedResources.Floatingips[0]

	return r.Get(ctx, resourceID, getParams, opts...)
}

// Update floating IP
func (r *FloatingIPService) Update(ctx context.Context, floatingIPID string, params FloatingIPUpdateParams, opts ...option.RequestOption) (res *FloatingIP, err error) {
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
	if floatingIPID == "" {
		err = errors.New("missing required floating_ip_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, floatingIPID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List floating IPs
func (r *FloatingIPService) List(ctx context.Context, params FloatingIPListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[FloatingIPDetailed], err error) {
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
	resource, err := r.Delete(ctx, floatingIPID, body, opts...)
	if err != nil {
		return err
	}

	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	return err
}

// Assign floating IP to instance or loadbalancer
func (r *FloatingIPService) Assign(ctx context.Context, floatingIPID string, params FloatingIPAssignParams, opts ...option.RequestOption) (res *FloatingIP, err error) {
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
	if floatingIPID == "" {
		err = errors.New("missing required floating_ip_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, floatingIPID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Unassign floating IP
func (r *FloatingIPService) Unassign(ctx context.Context, floatingIPID string, body FloatingIPUnassignParams, opts ...option.RequestOption) (res *FloatingIP, err error) {
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
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the floating IP was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required" format:"uuid4"`
	// IP address of the port the floating IP is attached to
	FixedIPAddress string `json:"fixed_ip_address,required" format:"ipvanyaddress"`
	// IP Address of the floating IP
	FloatingIPAddress string `json:"floating_ip_address,required" format:"ipvanyaddress"`
	// Instance the floating IP is attached to
	Instance FloatingIPDetailedInstance `json:"instance,required"`
	// Load balancer the floating IP is attached to
	Loadbalancer LoadBalancer `json:"loadbalancer,required"`
	// Port ID
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
	ID string `json:"id,required" format:"uuid4"`
	// Map of `network_name` to list of addresses in that network
	Addresses map[string][]FloatingIPDetailedInstanceAddressUnion `json:"addresses,required"`
	// Datetime when instance was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required"`
	// Flavor
	Flavor FloatingIPDetailedInstanceFlavor `json:"flavor,required"`
	// Instance description
	InstanceDescription string `json:"instance_description,required"`
	// Instance name
	Name string `json:"name,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Security groups
	SecurityGroups []FloatingIPDetailedInstanceSecurityGroup `json:"security_groups,required"`
	// SSH key name assigned to instance
	SSHKeyName string `json:"ssh_key_name,required"`
	// Instance status
	//
	// Any of "ACTIVE", "BUILD", "DELETED", "ERROR", "HARD_REBOOT", "MIGRATING",
	// "PASSWORD", "PAUSED", "REBOOT", "REBUILD", "RESCUE", "RESIZE", "REVERT_RESIZE",
	// "SHELVED", "SHELVED_OFFLOADED", "SHUTOFF", "SOFT_DELETED", "SUSPENDED",
	// "UNKNOWN", "VERIFY_RESIZE".
	Status string `json:"status,required"`
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
	VmState string `json:"vm_state,required"`
	// List of volumes
	Volumes []FloatingIPDetailedInstanceVolume `json:"volumes,required"`
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
func (r FloatingIPDetailedInstanceFlavor) RawJSON() string { return r.JSON.raw }
func (r *FloatingIPDetailedInstanceFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIPDetailedInstanceSecurityGroup struct {
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
func (r FloatingIPDetailedInstanceSecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *FloatingIPDetailedInstanceSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIPDetailedInstanceVolume struct {
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
func (r FloatingIPDetailedInstanceVolume) RawJSON() string { return r.JSON.raw }
func (r *FloatingIPDetailedInstanceVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIPNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// If the port has multiple IP addresses, a specific one can be selected using this
	// field. If not specified, the first IP in the port's list will be used by
	// default.
	FixedIPAddress param.Opt[string] `json:"fixed_ip_address,omitzero" format:"ipv4"`
	// If provided, the floating IP will be immediately attached to the specified port.
	PortID param.Opt[string] `json:"port_id,omitzero" format:"uuid4"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
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
	//   - **Delete tags:** `{'tags': {'`old_tag`': null}}` removes specific tags.
	//   - **Remove all tags:** `{'tags': null}` removes all user-managed tags (read-only
	//     tags are preserved).
	//   - **Partial update:** `{'tags': {'environment': 'staging'}}` only updates
	//     specified tags.
	//   - **Mixed operations:**
	//     `{'tags': {'environment': 'production', '`cost_center`': 'engineering', '`deprecated_tag`': null}}`
	//     adds/updates 'environment' and '`cost_center`' while removing
	//     '`deprecated_tag`', preserving other existing tags.
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Optional. Filter by tag key-value pairs.
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type FloatingIPAssignParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Port ID
	PortID string `json:"port_id,required" format:"uuid4"`
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type FloatingIPUnassignParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
