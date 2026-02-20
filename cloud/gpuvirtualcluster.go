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

// GPUVirtualClusterService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUVirtualClusterService] method instead.
type GPUVirtualClusterService struct {
	Options    []option.RequestOption
	Servers    GPUVirtualClusterServerService
	Volumes    GPUVirtualClusterVolumeService
	Interfaces GPUVirtualClusterInterfaceService
	Flavors    GPUVirtualClusterFlavorService
	Images     GPUVirtualClusterImageService
	tasks      TaskService
}

// NewGPUVirtualClusterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGPUVirtualClusterService(opts ...option.RequestOption) (r GPUVirtualClusterService) {
	r = GPUVirtualClusterService{}
	r.Options = opts
	r.Servers = NewGPUVirtualClusterServerService(opts...)
	r.Volumes = NewGPUVirtualClusterVolumeService(opts...)
	r.Interfaces = NewGPUVirtualClusterInterfaceService(opts...)
	r.Flavors = NewGPUVirtualClusterFlavorService(opts...)
	r.Images = NewGPUVirtualClusterImageService(opts...)
	r.tasks = NewTaskService(opts...)
	return
}

// Create a new virtual GPU cluster with the specified configuration.
func (r *GPUVirtualClusterService) New(ctx context.Context, params GPUVirtualClusterNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/clusters", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update the name of an existing virtual GPU cluster.
func (r *GPUVirtualClusterService) Update(ctx context.Context, clusterID string, params GPUVirtualClusterUpdateParams, opts ...option.RequestOption) (res *GPUVirtualCluster, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/clusters/%s", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List all virtual GPU clusters in the specified project and region.
func (r *GPUVirtualClusterService) List(ctx context.Context, params GPUVirtualClusterListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[GPUVirtualCluster], err error) {
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
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/clusters", params.ProjectID.Value, params.RegionID.Value)
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

// List all virtual GPU clusters in the specified project and region.
func (r *GPUVirtualClusterService) ListAutoPaging(ctx context.Context, params GPUVirtualClusterListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[GPUVirtualCluster] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete a virtual GPU cluster and all its associated resources.
func (r *GPUVirtualClusterService) Delete(ctx context.Context, clusterID string, params GPUVirtualClusterDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/clusters/%s", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Perform a specific action on a virtual GPU cluster. Available actions: start,
// stop, soft reboot, hard reboot, resize, update tags.
func (r *GPUVirtualClusterService) Action(ctx context.Context, clusterID string, params GPUVirtualClusterActionParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/clusters/%s/action", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get detailed information about a specific virtual GPU cluster.
func (r *GPUVirtualClusterService) Get(ctx context.Context, clusterID string, query GPUVirtualClusterGetParams, opts ...option.RequestOption) (res *GPUVirtualCluster, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/gpu/virtual/%v/%v/clusters/%s", query.ProjectID.Value, query.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// NewAndPoll creates a new virtual GPU cluster and polls for completion. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *GPUVirtualClusterService) NewAndPoll(ctx context.Context, params GPUVirtualClusterNewParams, opts ...option.RequestOption) (v *GPUVirtualCluster, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams GPUVirtualClusterGetParams
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

	// TODO: extract cluster ID from task created_resources when it becomes available, currently the API is not providing it
	if !task.JSON.Data.Valid() {
		return nil, errors.New("expected task data to be present")
	}
	dataMap, ok := task.Data.(map[string]any)
	if !ok {
		return nil, errors.New("expected task data to be a map")
	}
	clusterID, ok := dataMap["cluster_id"].(string)
	if !ok || clusterID == "" {
		return nil, errors.New("expected cluster_id to be present in task data")
	}

	return r.Get(ctx, clusterID, getParams, opts...)
}

// DeleteAndPoll deletes a virtual GPU cluster and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *GPUVirtualClusterService) DeleteAndPoll(ctx context.Context, clusterID string, params GPUVirtualClusterDeleteParams, opts ...option.RequestOption) error {
	resource, err := r.Delete(ctx, clusterID, params, opts...)
	if err != nil {
		return err
	}

	opts = slices.Concat(r.Options, opts)
	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	return err
}

// ActionAndPoll performs an action on a virtual GPU cluster and polls for completion of the first task. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *GPUVirtualClusterService) ActionAndPoll(ctx context.Context, clusterID string, params GPUVirtualClusterActionParams, opts ...option.RequestOption) (v *GPUVirtualCluster, err error) {
	resource, err := r.Action(ctx, clusterID, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams GPUVirtualClusterGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) == 0 {
		return nil, errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, clusterID, getParams, opts...)
}

type GPUVirtualCluster struct {
	// Cluster unique identifier
	ID string `json:"id,required" format:"uuid4"`
	// Cluster creation date time
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Cluster flavor name
	Flavor string `json:"flavor,required"`
	// Cluster name
	Name string `json:"name,required"`
	// Cluster servers count
	ServersCount int64 `json:"servers_count,required"`
	// List of cluster nodes
	ServersIDs      []string                         `json:"servers_ids,required" format:"uuid4"`
	ServersSettings GPUVirtualClusterServersSettings `json:"servers_settings,required"`
	// Cluster status
	//
	// Any of "active", "creating", "degraded", "deleting", "error", "new",
	// "rebooting", "rebuilding", "resizing", "shutoff".
	Status GPUVirtualClusterStatus `json:"status,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// Cluster update date time
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Flavor          respjson.Field
		Name            respjson.Field
		ServersCount    respjson.Field
		ServersIDs      respjson.Field
		ServersSettings respjson.Field
		Status          respjson.Field
		Tags            respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualCluster) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterServersSettings struct {
	// List of file shares mounted across the cluster.
	FileShares []GPUVirtualClusterServersSettingsFileShare      `json:"file_shares,required"`
	Interfaces []GPUVirtualClusterServersSettingsInterfaceUnion `json:"interfaces,required"`
	// Security groups
	SecurityGroups []GPUVirtualClusterServersSettingsSecurityGroup `json:"security_groups,required"`
	// SSH key name
	SSHKeyName string `json:"ssh_key_name,required"`
	// Optional custom user data
	UserData string `json:"user_data,required"`
	// List of volumes
	Volumes []GPUVirtualClusterServersSettingsVolume `json:"volumes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileShares     respjson.Field
		Interfaces     respjson.Field
		SecurityGroups respjson.Field
		SSHKeyName     respjson.Field
		UserData       respjson.Field
		Volumes        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualClusterServersSettings) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualClusterServersSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterServersSettingsFileShare struct {
	// Unique identifier of the file share in UUID format.
	ID string `json:"id,required" format:"uuid4"`
	// Absolute mount path inside the system where the file share will be mounted.
	MountPath string `json:"mount_path,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		MountPath   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualClusterServersSettingsFileShare) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualClusterServersSettingsFileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUVirtualClusterServersSettingsInterfaceUnion contains all possible properties
// and values from [GPUVirtualClusterServersSettingsInterfaceExternal],
// [GPUVirtualClusterServersSettingsInterfaceSubnet],
// [GPUVirtualClusterServersSettingsInterfaceAnySubnet].
//
// Use the [GPUVirtualClusterServersSettingsInterfaceUnion.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type GPUVirtualClusterServersSettingsInterfaceUnion struct {
	IPFamily string `json:"ip_family"`
	Name     string `json:"name"`
	// Any of "external", "subnet", "any_subnet".
	Type string `json:"type"`
	// This field is a union of
	// [GPUVirtualClusterServersSettingsInterfaceSubnetFloatingIP],
	// [GPUVirtualClusterServersSettingsInterfaceAnySubnetFloatingIP]
	FloatingIP GPUVirtualClusterServersSettingsInterfaceUnionFloatingIP `json:"floating_ip"`
	NetworkID  string                                                   `json:"network_id"`
	// This field is from variant [GPUVirtualClusterServersSettingsInterfaceSubnet].
	SubnetID string `json:"subnet_id"`
	// This field is from variant [GPUVirtualClusterServersSettingsInterfaceAnySubnet].
	IPAddress string `json:"ip_address"`
	JSON      struct {
		IPFamily   respjson.Field
		Name       respjson.Field
		Type       respjson.Field
		FloatingIP respjson.Field
		NetworkID  respjson.Field
		SubnetID   respjson.Field
		IPAddress  respjson.Field
		raw        string
	} `json:"-"`
}

// anyGPUVirtualClusterServersSettingsInterface is implemented by each variant of
// [GPUVirtualClusterServersSettingsInterfaceUnion] to add type safety for the
// return type of [GPUVirtualClusterServersSettingsInterfaceUnion.AsAny]
type anyGPUVirtualClusterServersSettingsInterface interface {
	implGPUVirtualClusterServersSettingsInterfaceUnion()
}

func (GPUVirtualClusterServersSettingsInterfaceExternal) implGPUVirtualClusterServersSettingsInterfaceUnion() {
}
func (GPUVirtualClusterServersSettingsInterfaceSubnet) implGPUVirtualClusterServersSettingsInterfaceUnion() {
}
func (GPUVirtualClusterServersSettingsInterfaceAnySubnet) implGPUVirtualClusterServersSettingsInterfaceUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := GPUVirtualClusterServersSettingsInterfaceUnion.AsAny().(type) {
//	case cloud.GPUVirtualClusterServersSettingsInterfaceExternal:
//	case cloud.GPUVirtualClusterServersSettingsInterfaceSubnet:
//	case cloud.GPUVirtualClusterServersSettingsInterfaceAnySubnet:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u GPUVirtualClusterServersSettingsInterfaceUnion) AsAny() anyGPUVirtualClusterServersSettingsInterface {
	switch u.Type {
	case "external":
		return u.AsExternal()
	case "subnet":
		return u.AsSubnet()
	case "any_subnet":
		return u.AsAnySubnet()
	}
	return nil
}

func (u GPUVirtualClusterServersSettingsInterfaceUnion) AsExternal() (v GPUVirtualClusterServersSettingsInterfaceExternal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GPUVirtualClusterServersSettingsInterfaceUnion) AsSubnet() (v GPUVirtualClusterServersSettingsInterfaceSubnet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GPUVirtualClusterServersSettingsInterfaceUnion) AsAnySubnet() (v GPUVirtualClusterServersSettingsInterfaceAnySubnet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GPUVirtualClusterServersSettingsInterfaceUnion) RawJSON() string { return u.JSON.raw }

func (r *GPUVirtualClusterServersSettingsInterfaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUVirtualClusterServersSettingsInterfaceUnionFloatingIP is an implicit subunion
// of [GPUVirtualClusterServersSettingsInterfaceUnion].
// GPUVirtualClusterServersSettingsInterfaceUnionFloatingIP provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [GPUVirtualClusterServersSettingsInterfaceUnion].
type GPUVirtualClusterServersSettingsInterfaceUnionFloatingIP struct {
	// This field is from variant
	// [GPUVirtualClusterServersSettingsInterfaceSubnetFloatingIP].
	Source constant.New `json:"source"`
	JSON   struct {
		Source respjson.Field
		raw    string
	} `json:"-"`
}

func (r *GPUVirtualClusterServersSettingsInterfaceUnionFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterServersSettingsInterfaceExternal struct {
	// Which subnets should be selected: IPv4, IPv6, or use dual stack.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,required"`
	// Interface name
	Name string            `json:"name,required"`
	Type constant.External `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPFamily    respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualClusterServersSettingsInterfaceExternal) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualClusterServersSettingsInterfaceExternal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterServersSettingsInterfaceSubnet struct {
	// Floating IP config for this subnet attachment
	FloatingIP GPUVirtualClusterServersSettingsInterfaceSubnetFloatingIP `json:"floating_ip,required"`
	// Interface name
	Name string `json:"name,required"`
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID string `json:"network_id,required"`
	// Port is assigned an IP address from this subnet
	SubnetID string          `json:"subnet_id,required" format:"uuid4"`
	Type     constant.Subnet `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FloatingIP  respjson.Field
		Name        respjson.Field
		NetworkID   respjson.Field
		SubnetID    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualClusterServersSettingsInterfaceSubnet) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualClusterServersSettingsInterfaceSubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Floating IP config for this subnet attachment
type GPUVirtualClusterServersSettingsInterfaceSubnetFloatingIP struct {
	Source constant.New `json:"source,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualClusterServersSettingsInterfaceSubnetFloatingIP) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUVirtualClusterServersSettingsInterfaceSubnetFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterServersSettingsInterfaceAnySubnet struct {
	// Floating IP config for this subnet attachment
	FloatingIP GPUVirtualClusterServersSettingsInterfaceAnySubnetFloatingIP `json:"floating_ip,required"`
	// Fixed IP address
	IPAddress string `json:"ip_address,required"`
	// Which subnets should be selected: IPv4, IPv6, or use dual stack
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,required"`
	// Interface name
	Name string `json:"name,required"`
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID string             `json:"network_id,required"`
	Type      constant.AnySubnet `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FloatingIP  respjson.Field
		IPAddress   respjson.Field
		IPFamily    respjson.Field
		Name        respjson.Field
		NetworkID   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualClusterServersSettingsInterfaceAnySubnet) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualClusterServersSettingsInterfaceAnySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Floating IP config for this subnet attachment
type GPUVirtualClusterServersSettingsInterfaceAnySubnetFloatingIP struct {
	Source constant.New `json:"source,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualClusterServersSettingsInterfaceAnySubnetFloatingIP) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUVirtualClusterServersSettingsInterfaceAnySubnetFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterServersSettingsSecurityGroup struct {
	// Security group ID
	ID string `json:"id,required"`
	// Security group name
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualClusterServersSettingsSecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualClusterServersSettingsSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterServersSettingsVolume struct {
	// Boot index of the volume
	BootIndex int64 `json:"boot_index,required"`
	// Flag indicating whether the volume is deleted on instance termination
	DeleteOnTermination bool `json:"delete_on_termination,required"`
	// Image ID for the volume
	ImageID string `json:"image_id,required"`
	// Volume name
	Name string `json:"name,required"`
	// Volume size in GiB
	Size int64 `json:"size,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// Volume type
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BootIndex           respjson.Field
		DeleteOnTermination respjson.Field
		ImageID             respjson.Field
		Name                respjson.Field
		Size                respjson.Field
		Tags                respjson.Field
		Type                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUVirtualClusterServersSettingsVolume) RawJSON() string { return r.JSON.raw }
func (r *GPUVirtualClusterServersSettingsVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cluster status
type GPUVirtualClusterStatus string

const (
	GPUVirtualClusterStatusActive     GPUVirtualClusterStatus = "active"
	GPUVirtualClusterStatusCreating   GPUVirtualClusterStatus = "creating"
	GPUVirtualClusterStatusDegraded   GPUVirtualClusterStatus = "degraded"
	GPUVirtualClusterStatusDeleting   GPUVirtualClusterStatus = "deleting"
	GPUVirtualClusterStatusError      GPUVirtualClusterStatus = "error"
	GPUVirtualClusterStatusNew        GPUVirtualClusterStatus = "new"
	GPUVirtualClusterStatusRebooting  GPUVirtualClusterStatus = "rebooting"
	GPUVirtualClusterStatusRebuilding GPUVirtualClusterStatus = "rebuilding"
	GPUVirtualClusterStatusResizing   GPUVirtualClusterStatus = "resizing"
	GPUVirtualClusterStatusShutoff    GPUVirtualClusterStatus = "shutoff"
)

type GPUVirtualClusterNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Cluster flavor ID
	Flavor string `json:"flavor,required"`
	// Cluster name
	Name string `json:"name,required"`
	// Number of servers in the cluster
	ServersCount int64 `json:"servers_count,required"`
	// Configuration settings for the servers in the cluster
	ServersSettings GPUVirtualClusterNewParamsServersSettings `json:"servers_settings,omitzero,required"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Some tags are read-only and cannot be modified by the
	// user. Tags are also integrated with cost reports, allowing cost data to be
	// filtered based on tag keys or values.
	Tags any `json:"tags,omitzero"`
	paramObj
}

func (r GPUVirtualClusterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration settings for the servers in the cluster
//
// The properties Interfaces, Volumes are required.
type GPUVirtualClusterNewParamsServersSettings struct {
	// Subnet IPs and floating IPs
	Interfaces []GPUVirtualClusterNewParamsServersSettingsInterfaceUnion `json:"interfaces,omitzero,required"`
	// List of volumes
	Volumes []GPUVirtualClusterNewParamsServersSettingsVolumeUnion `json:"volumes,omitzero,required"`
	// Optional custom user data (Base64-encoded)
	UserData param.Opt[string] `json:"user_data,omitzero"`
	// Optional server access credentials
	Credentials GPUVirtualClusterNewParamsServersSettingsCredentials `json:"credentials,omitzero"`
	// List of file shares to be mounted across the cluster.
	FileShares []GPUVirtualClusterNewParamsServersSettingsFileShare `json:"file_shares,omitzero"`
	// List of security groups UUIDs
	SecurityGroups []GPUVirtualClusterNewParamsServersSettingsSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

func (r GPUVirtualClusterNewParamsServersSettings) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterNewParamsServersSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterNewParamsServersSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GPUVirtualClusterNewParamsServersSettingsInterfaceUnion struct {
	OfExternal  *GPUVirtualClusterNewParamsServersSettingsInterfaceExternal  `json:",omitzero,inline"`
	OfSubnet    *GPUVirtualClusterNewParamsServersSettingsInterfaceSubnet    `json:",omitzero,inline"`
	OfAnySubnet *GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnet `json:",omitzero,inline"`
	paramUnion
}

func (u GPUVirtualClusterNewParamsServersSettingsInterfaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExternal, u.OfSubnet, u.OfAnySubnet)
}
func (u *GPUVirtualClusterNewParamsServersSettingsInterfaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GPUVirtualClusterNewParamsServersSettingsInterfaceUnion) asAny() any {
	if !param.IsOmitted(u.OfExternal) {
		return u.OfExternal
	} else if !param.IsOmitted(u.OfSubnet) {
		return u.OfSubnet
	} else if !param.IsOmitted(u.OfAnySubnet) {
		return u.OfAnySubnet
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUVirtualClusterNewParamsServersSettingsInterfaceUnion) GetSubnetID() *string {
	if vt := u.OfSubnet; vt != nil {
		return &vt.SubnetID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUVirtualClusterNewParamsServersSettingsInterfaceUnion) GetType() *string {
	if vt := u.OfExternal; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSubnet; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAnySubnet; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUVirtualClusterNewParamsServersSettingsInterfaceUnion) GetIPFamily() *string {
	if vt := u.OfExternal; vt != nil {
		return (*string)(&vt.IPFamily)
	} else if vt := u.OfAnySubnet; vt != nil {
		return (*string)(&vt.IPFamily)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUVirtualClusterNewParamsServersSettingsInterfaceUnion) GetName() *string {
	if vt := u.OfExternal; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfSubnet; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfAnySubnet; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUVirtualClusterNewParamsServersSettingsInterfaceUnion) GetNetworkID() *string {
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
func (u GPUVirtualClusterNewParamsServersSettingsInterfaceUnion) GetFloatingIP() (res gpuVirtualClusterNewParamsServersSettingsInterfaceUnionFloatingIP) {
	if vt := u.OfSubnet; vt != nil {
		res.any = &vt.FloatingIP
	} else if vt := u.OfAnySubnet; vt != nil {
		res.any = &vt.FloatingIP
	}
	return
}

// Can have the runtime types
// [*GPUVirtualClusterNewParamsServersSettingsInterfaceSubnetFloatingIP],
// [*GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP]
type gpuVirtualClusterNewParamsServersSettingsInterfaceUnionFloatingIP struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cloud.GPUVirtualClusterNewParamsServersSettingsInterfaceSubnetFloatingIP:
//	case *cloud.GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u gpuVirtualClusterNewParamsServersSettingsInterfaceUnionFloatingIP) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u gpuVirtualClusterNewParamsServersSettingsInterfaceUnionFloatingIP) GetSource() *string {
	switch vt := u.any.(type) {
	case *GPUVirtualClusterNewParamsServersSettingsInterfaceSubnetFloatingIP:
		return (*string)(&vt.Source)
	case *GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP:
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[GPUVirtualClusterNewParamsServersSettingsInterfaceUnion](
		"type",
		apijson.Discriminator[GPUVirtualClusterNewParamsServersSettingsInterfaceExternal]("external"),
		apijson.Discriminator[GPUVirtualClusterNewParamsServersSettingsInterfaceSubnet]("subnet"),
		apijson.Discriminator[GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnet]("any_subnet"),
	)
}

// The property Type is required.
type GPUVirtualClusterNewParamsServersSettingsInterfaceExternal struct {
	// Interface name
	Name param.Opt[string] `json:"name,omitzero"`
	// Which subnets should be selected: IPv4, IPv6, or use dual stack.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// This field can be elided, and will marshal its zero value as "external".
	Type constant.External `json:"type,required"`
	paramObj
}

func (r GPUVirtualClusterNewParamsServersSettingsInterfaceExternal) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterNewParamsServersSettingsInterfaceExternal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterNewParamsServersSettingsInterfaceExternal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GPUVirtualClusterNewParamsServersSettingsInterfaceExternal](
		"ip_family", "dual", "ipv4", "ipv6",
	)
}

// The properties NetworkID, SubnetID, Type are required.
type GPUVirtualClusterNewParamsServersSettingsInterfaceSubnet struct {
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID string `json:"network_id,required"`
	// Port is assigned an IP address from this subnet
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Interface name
	Name param.Opt[string] `json:"name,omitzero"`
	// Floating IP config for this subnet attachment
	FloatingIP GPUVirtualClusterNewParamsServersSettingsInterfaceSubnetFloatingIP `json:"floating_ip,omitzero"`
	// This field can be elided, and will marshal its zero value as "subnet".
	Type constant.Subnet `json:"type,required"`
	paramObj
}

func (r GPUVirtualClusterNewParamsServersSettingsInterfaceSubnet) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterNewParamsServersSettingsInterfaceSubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterNewParamsServersSettingsInterfaceSubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewGPUVirtualClusterNewParamsServersSettingsInterfaceSubnetFloatingIP() GPUVirtualClusterNewParamsServersSettingsInterfaceSubnetFloatingIP {
	return GPUVirtualClusterNewParamsServersSettingsInterfaceSubnetFloatingIP{
		Source: "new",
	}
}

// Floating IP config for this subnet attachment
//
// This struct has a constant value, construct it with
// [NewGPUVirtualClusterNewParamsServersSettingsInterfaceSubnetFloatingIP].
type GPUVirtualClusterNewParamsServersSettingsInterfaceSubnetFloatingIP struct {
	Source constant.New `json:"source,required"`
	paramObj
}

func (r GPUVirtualClusterNewParamsServersSettingsInterfaceSubnetFloatingIP) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterNewParamsServersSettingsInterfaceSubnetFloatingIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterNewParamsServersSettingsInterfaceSubnetFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties NetworkID, Type are required.
type GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnet struct {
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID string `json:"network_id,required"`
	// Interface name
	Name param.Opt[string] `json:"name,omitzero"`
	// Floating IP config for this subnet attachment
	FloatingIP GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP `json:"floating_ip,omitzero"`
	// Which subnets should be selected: IPv4, IPv6, or use dual stack
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// This field can be elided, and will marshal its zero value as "any_subnet".
	Type constant.AnySubnet `json:"type,required"`
	paramObj
}

func (r GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnet) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnet](
		"ip_family", "dual", "ipv4", "ipv6",
	)
}

func NewGPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP() GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP {
	return GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP{
		Source: "new",
	}
}

// Floating IP config for this subnet attachment
//
// This struct has a constant value, construct it with
// [NewGPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP].
type GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP struct {
	Source constant.New `json:"source,required"`
	paramObj
}

func (r GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GPUVirtualClusterNewParamsServersSettingsVolumeUnion struct {
	OfNew   *GPUVirtualClusterNewParamsServersSettingsVolumeNew   `json:",omitzero,inline"`
	OfImage *GPUVirtualClusterNewParamsServersSettingsVolumeImage `json:",omitzero,inline"`
	paramUnion
}

func (u GPUVirtualClusterNewParamsServersSettingsVolumeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNew, u.OfImage)
}
func (u *GPUVirtualClusterNewParamsServersSettingsVolumeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GPUVirtualClusterNewParamsServersSettingsVolumeUnion) asAny() any {
	if !param.IsOmitted(u.OfNew) {
		return u.OfNew
	} else if !param.IsOmitted(u.OfImage) {
		return u.OfImage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUVirtualClusterNewParamsServersSettingsVolumeUnion) GetImageID() *string {
	if vt := u.OfImage; vt != nil {
		return &vt.ImageID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUVirtualClusterNewParamsServersSettingsVolumeUnion) GetBootIndex() *int64 {
	if vt := u.OfNew; vt != nil {
		return (*int64)(&vt.BootIndex)
	} else if vt := u.OfImage; vt != nil {
		return (*int64)(&vt.BootIndex)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUVirtualClusterNewParamsServersSettingsVolumeUnion) GetName() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUVirtualClusterNewParamsServersSettingsVolumeUnion) GetSize() *int64 {
	if vt := u.OfNew; vt != nil {
		return (*int64)(&vt.Size)
	} else if vt := u.OfImage; vt != nil {
		return (*int64)(&vt.Size)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUVirtualClusterNewParamsServersSettingsVolumeUnion) GetSource() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Source)
	} else if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Source)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUVirtualClusterNewParamsServersSettingsVolumeUnion) GetType() *string {
	if vt := u.OfNew; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUVirtualClusterNewParamsServersSettingsVolumeUnion) GetDeleteOnTermination() *bool {
	if vt := u.OfNew; vt != nil && vt.DeleteOnTermination.Valid() {
		return &vt.DeleteOnTermination.Value
	} else if vt := u.OfImage; vt != nil && vt.DeleteOnTermination.Valid() {
		return &vt.DeleteOnTermination.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's Tags property, if present.
func (u GPUVirtualClusterNewParamsServersSettingsVolumeUnion) GetTags() map[string]string {
	if vt := u.OfNew; vt != nil {
		return vt.Tags
	} else if vt := u.OfImage; vt != nil {
		return vt.Tags
	}
	return nil
}

func init() {
	apijson.RegisterUnion[GPUVirtualClusterNewParamsServersSettingsVolumeUnion](
		"source",
		apijson.Discriminator[GPUVirtualClusterNewParamsServersSettingsVolumeNew]("new"),
		apijson.Discriminator[GPUVirtualClusterNewParamsServersSettingsVolumeImage]("image"),
	)
}

// The properties BootIndex, Name, Size, Source, Type are required.
type GPUVirtualClusterNewParamsServersSettingsVolumeNew struct {
	// Boot index of the volume
	BootIndex int64 `json:"boot_index,required"`
	// Volume name
	Name string `json:"name,required"`
	// Volume size in GiB
	Size int64 `json:"size,required"`
	// Volume type
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	Type string `json:"type,omitzero,required"`
	// Flag indicating whether the volume is deleted on instance termination
	DeleteOnTermination param.Opt[bool] `json:"delete_on_termination,omitzero"`
	// Tags associated with the volume
	Tags map[string]string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "new".
	Source constant.New `json:"source,required"`
	paramObj
}

func (r GPUVirtualClusterNewParamsServersSettingsVolumeNew) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterNewParamsServersSettingsVolumeNew
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterNewParamsServersSettingsVolumeNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GPUVirtualClusterNewParamsServersSettingsVolumeNew](
		"type", "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

// The properties BootIndex, ImageID, Name, Size, Source, Type are required.
type GPUVirtualClusterNewParamsServersSettingsVolumeImage struct {
	// Boot index of the volume
	BootIndex int64 `json:"boot_index,required"`
	// Image ID for the volume
	ImageID string `json:"image_id,required" format:"uuid4"`
	// Volume name
	Name string `json:"name,required"`
	// Volume size in GiB
	Size int64 `json:"size,required"`
	// Volume type
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	Type string `json:"type,omitzero,required"`
	// Flag indicating whether the volume is deleted on instance termination
	DeleteOnTermination param.Opt[bool] `json:"delete_on_termination,omitzero"`
	// Tags associated with the volume
	Tags map[string]string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "image".
	Source constant.Image `json:"source,required"`
	paramObj
}

func (r GPUVirtualClusterNewParamsServersSettingsVolumeImage) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterNewParamsServersSettingsVolumeImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterNewParamsServersSettingsVolumeImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GPUVirtualClusterNewParamsServersSettingsVolumeImage](
		"type", "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

// Optional server access credentials
type GPUVirtualClusterNewParamsServersSettingsCredentials struct {
	// Used to set the password for the specified 'username' on Linux instances. If
	// 'username' is not provided, the password is applied to the default user of the
	// image. Mutually exclusive with 'user_data' - only one can be specified.
	Password param.Opt[string] `json:"password,omitzero" format:"password"`
	// Specifies the name of the SSH keypair, created via the
	// [/v1/`ssh_keys` endpoint](/docs/api-reference/cloud/ssh-keys/add-or-generate-ssh-key).
	SSHKeyName param.Opt[string] `json:"ssh_key_name,omitzero"`
	// The 'username' and 'password' fields create a new user on the system
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r GPUVirtualClusterNewParamsServersSettingsCredentials) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterNewParamsServersSettingsCredentials
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterNewParamsServersSettingsCredentials) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, MountPath are required.
type GPUVirtualClusterNewParamsServersSettingsFileShare struct {
	// Unique identifier of the file share in UUID format.
	ID string `json:"id,required" format:"uuid4"`
	// Absolute mount path inside the system where the file share will be mounted.
	MountPath string `json:"mount_path,required"`
	paramObj
}

func (r GPUVirtualClusterNewParamsServersSettingsFileShare) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterNewParamsServersSettingsFileShare
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterNewParamsServersSettingsFileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type GPUVirtualClusterNewParamsServersSettingsSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

func (r GPUVirtualClusterNewParamsServersSettingsSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterNewParamsServersSettingsSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterNewParamsServersSettingsSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Cluster name
	Name string `json:"name,required"`
	paramObj
}

func (r GPUVirtualClusterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Limit of items on a single page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset in results list
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GPUVirtualClusterListParams]'s query parameters as
// `url.Values`.
func (r GPUVirtualClusterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type GPUVirtualClusterDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Flag indicating whether the floating ips associated with server / cluster are
	// deleted
	AllFloatingIPs param.Opt[bool] `query:"all_floating_ips,omitzero" json:"-"`
	// Flag indicating whether the reserved fixed ips associated with server / cluster
	// are deleted
	AllReservedFixedIPs param.Opt[bool] `query:"all_reserved_fixed_ips,omitzero" json:"-"`
	// Flag indicating whether all attached volumes are deleted
	AllVolumes param.Opt[bool] `query:"all_volumes,omitzero" json:"-"`
	// Optional list of floating ips to be deleted
	FloatingIPIDs []string `query:"floating_ip_ids,omitzero" format:"uuid4" json:"-"`
	// Optional list of reserved fixed ips to be deleted
	ReservedFixedIPIDs []string `query:"reserved_fixed_ip_ids,omitzero" format:"uuid4" json:"-"`
	// Optional list of volumes to be deleted
	VolumeIDs []string `query:"volume_ids,omitzero" format:"uuid4" json:"-"`
	paramObj
}

// URLQuery serializes [GPUVirtualClusterDeleteParams]'s query parameters as
// `url.Values`.
func (r GPUVirtualClusterDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type GPUVirtualClusterActionParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfStart *GPUVirtualClusterActionParamsBodyStart `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfStop *GPUVirtualClusterActionParamsBodyStop `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSoftReboot *GPUVirtualClusterActionParamsBodySoftReboot `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfHardReboot *GPUVirtualClusterActionParamsBodyHardReboot `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfUpdateTags *GPUVirtualClusterActionParamsBodyUpdateTags `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfResize *GPUVirtualClusterActionParamsBodyResize `json:",inline"`

	paramObj
}

func (u GPUVirtualClusterActionParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStart,
		u.OfStop,
		u.OfSoftReboot,
		u.OfHardReboot,
		u.OfUpdateTags,
		u.OfResize)
}
func (r *GPUVirtualClusterActionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewGPUVirtualClusterActionParamsBodyStart() GPUVirtualClusterActionParamsBodyStart {
	return GPUVirtualClusterActionParamsBodyStart{
		Action: "start",
	}
}

// This struct has a constant value, construct it with
// [NewGPUVirtualClusterActionParamsBodyStart].
type GPUVirtualClusterActionParamsBodyStart struct {
	// Action name
	Action constant.Start `json:"action,required"`
	paramObj
}

func (r GPUVirtualClusterActionParamsBodyStart) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterActionParamsBodyStart
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterActionParamsBodyStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewGPUVirtualClusterActionParamsBodyStop() GPUVirtualClusterActionParamsBodyStop {
	return GPUVirtualClusterActionParamsBodyStop{
		Action: "stop",
	}
}

// This struct has a constant value, construct it with
// [NewGPUVirtualClusterActionParamsBodyStop].
type GPUVirtualClusterActionParamsBodyStop struct {
	// Action name
	Action constant.Stop `json:"action,required"`
	paramObj
}

func (r GPUVirtualClusterActionParamsBodyStop) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterActionParamsBodyStop
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterActionParamsBodyStop) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewGPUVirtualClusterActionParamsBodySoftReboot() GPUVirtualClusterActionParamsBodySoftReboot {
	return GPUVirtualClusterActionParamsBodySoftReboot{
		Action: "soft_reboot",
	}
}

// This struct has a constant value, construct it with
// [NewGPUVirtualClusterActionParamsBodySoftReboot].
type GPUVirtualClusterActionParamsBodySoftReboot struct {
	// Action name
	Action constant.SoftReboot `json:"action,required"`
	paramObj
}

func (r GPUVirtualClusterActionParamsBodySoftReboot) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterActionParamsBodySoftReboot
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterActionParamsBodySoftReboot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewGPUVirtualClusterActionParamsBodyHardReboot() GPUVirtualClusterActionParamsBodyHardReboot {
	return GPUVirtualClusterActionParamsBodyHardReboot{
		Action: "hard_reboot",
	}
}

// This struct has a constant value, construct it with
// [NewGPUVirtualClusterActionParamsBodyHardReboot].
type GPUVirtualClusterActionParamsBodyHardReboot struct {
	// Action name
	Action constant.HardReboot `json:"action,required"`
	paramObj
}

func (r GPUVirtualClusterActionParamsBodyHardReboot) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterActionParamsBodyHardReboot
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterActionParamsBodyHardReboot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, Tags are required.
type GPUVirtualClusterActionParamsBodyUpdateTags struct {
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
	Tags TagUpdateMap `json:"tags,omitzero,required"`
	// Action name
	//
	// This field can be elided, and will marshal its zero value as "update_tags".
	Action constant.UpdateTags `json:"action,required"`
	paramObj
}

func (r GPUVirtualClusterActionParamsBodyUpdateTags) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterActionParamsBodyUpdateTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterActionParamsBodyUpdateTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, ServersCount are required.
type GPUVirtualClusterActionParamsBodyResize struct {
	// Requested servers count
	ServersCount int64 `json:"servers_count,required"`
	// Action name
	//
	// This field can be elided, and will marshal its zero value as "resize".
	Action constant.Resize `json:"action,required"`
	paramObj
}

func (r GPUVirtualClusterActionParamsBodyResize) MarshalJSON() (data []byte, err error) {
	type shadow GPUVirtualClusterActionParamsBodyResize
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUVirtualClusterActionParamsBodyResize) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUVirtualClusterGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
