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

// GPUBaremetalClusterService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUBaremetalClusterService] method instead.
type GPUBaremetalClusterService struct {
	Options    []option.RequestOption
	Interfaces GPUBaremetalClusterInterfaceService
	Servers    GPUBaremetalClusterServerService
	Flavors    GPUBaremetalClusterFlavorService
	Images     GPUBaremetalClusterImageService
	tasks      TaskService
}

// NewGPUBaremetalClusterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGPUBaremetalClusterService(opts ...option.RequestOption) (r GPUBaremetalClusterService) {
	r = GPUBaremetalClusterService{}
	r.Options = opts
	r.Interfaces = NewGPUBaremetalClusterInterfaceService(opts...)
	r.Servers = NewGPUBaremetalClusterServerService(opts...)
	r.Flavors = NewGPUBaremetalClusterFlavorService(opts...)
	r.Images = NewGPUBaremetalClusterImageService(opts...)
	r.tasks = NewTaskService(opts...)
	return
}

// Create a new bare metal GPU cluster with the specified configuration.
func (r *GPUBaremetalClusterService) New(ctx context.Context, params GPUBaremetalClusterNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/clusters", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List all bare metal GPU clusters in the specified project and region.
func (r *GPUBaremetalClusterService) List(ctx context.Context, params GPUBaremetalClusterListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[GPUBaremetalCluster], err error) {
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
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/clusters", params.ProjectID.Value, params.RegionID.Value)
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

// List all bare metal GPU clusters in the specified project and region.
func (r *GPUBaremetalClusterService) ListAutoPaging(ctx context.Context, params GPUBaremetalClusterListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[GPUBaremetalCluster] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete a bare metal GPU cluster and all its associated resources.
func (r *GPUBaremetalClusterService) Delete(ctx context.Context, clusterID string, params GPUBaremetalClusterDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/clusters/%s", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Perform a specific action on a baremetal GPU cluster. Available actions: update
// tags.
func (r *GPUBaremetalClusterService) Action(ctx context.Context, clusterID string, params GPUBaremetalClusterActionParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/clusters/%s/action", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get detailed information about a specific bare metal GPU cluster.
func (r *GPUBaremetalClusterService) Get(ctx context.Context, clusterID string, query GPUBaremetalClusterGetParams, opts ...option.RequestOption) (res *GPUBaremetalCluster, err error) {
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
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/clusters/%s", query.ProjectID.Value, query.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Stops and then starts all cluster servers, effectively performing a hard reboot.
func (r *GPUBaremetalClusterService) PowercycleAllServers(ctx context.Context, clusterID string, body GPUBaremetalClusterPowercycleAllServersParams, opts ...option.RequestOption) (res *GPUBaremetalClusterServerV1List, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s/powercycle", body.ProjectID.Value, body.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Reboot all bare metal GPU cluster servers
func (r *GPUBaremetalClusterService) RebootAllServers(ctx context.Context, clusterID string, body GPUBaremetalClusterRebootAllServersParams, opts ...option.RequestOption) (res *GPUBaremetalClusterServerV1List, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s/reboot", body.ProjectID.Value, body.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Rebuild one or more nodes in a GPU cluster. All cluster nodes must be specified
// to update the cluster image.
func (r *GPUBaremetalClusterService) Rebuild(ctx context.Context, clusterID string, params GPUBaremetalClusterRebuildParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/ai/clusters/gpu/%v/%v/%s/rebuild", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Change the number of nodes in a GPU cluster. The cluster can be scaled up or
// down.
func (r *GPUBaremetalClusterService) Resize(ctx context.Context, clusterID string, params GPUBaremetalClusterResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/ai/clusters/gpu/%v/%v/%s/resize", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// NewAndPoll creates a new GPU bare metal cluster and polls for completion. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *GPUBaremetalClusterService) NewAndPoll(ctx context.Context, params GPUBaremetalClusterNewParams, opts ...option.RequestOption) (v *GPUBaremetalCluster, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams GPUBaremetalClusterGetParams
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

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Clusters) != 1 {
		return nil, errors.New("expected exactly one cluster to be created in a task")
	}
	clusterID := task.CreatedResources.Clusters[0]

	return r.Get(ctx, clusterID, getParams, opts...)
}

// RebuildAndPoll rebuilds a GPU bare metal cluster and polls for completion of the first task. Use the
// [TaskService.Poll] method if you need to poll for all tasks.
func (r *GPUBaremetalClusterService) RebuildAndPoll(ctx context.Context, clusterID string, params GPUBaremetalClusterRebuildParams, opts ...option.RequestOption) (v *GPUBaremetalCluster, err error) {
	resource, err := r.Rebuild(ctx, clusterID, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams GPUBaremetalClusterGetParams
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
		return
	}

	return r.Get(ctx, clusterID, getParams, opts...)
}

// ResizeAndPoll resizes a GPU bare metal cluster and polls for completion of the first task. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *GPUBaremetalClusterService) ResizeAndPoll(ctx context.Context, clusterID string, params GPUBaremetalClusterResizeParams, opts ...option.RequestOption) (v *GPUBaremetalCluster, err error) {
	resource, err := r.Resize(ctx, clusterID, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams GPUBaremetalClusterGetParams
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
		return
	}

	return r.Get(ctx, clusterID, getParams, opts...)
}

type GPUBaremetalCluster struct {
	// Cluster unique identifier
	ID string `json:"id,required" format:"uuid4"`
	// Cluster creation date time
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Cluster flavor name
	Flavor string `json:"flavor,required"`
	// Image ID
	ImageID string `json:"image_id,required"`
	// User type managing the resource
	//
	// Any of "k8s", "user".
	ManagedBy GPUBaremetalClusterManagedBy `json:"managed_by,required"`
	// Cluster name
	Name string `json:"name,required"`
	// Cluster servers count
	ServersCount int64 `json:"servers_count,required"`
	// List of cluster nodes
	ServersIDs      []string                           `json:"servers_ids,required" format:"uuid4"`
	ServersSettings GPUBaremetalClusterServersSettings `json:"servers_settings,required"`
	// Cluster status
	//
	// Any of "active", "creating", "degraded", "deleting", "error", "new",
	// "rebooting", "rebuilding", "resizing", "shutoff".
	Status GPUBaremetalClusterStatus `json:"status,required"`
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
		ImageID         respjson.Field
		ManagedBy       respjson.Field
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
func (r GPUBaremetalCluster) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User type managing the resource
type GPUBaremetalClusterManagedBy string

const (
	GPUBaremetalClusterManagedByK8S  GPUBaremetalClusterManagedBy = "k8s"
	GPUBaremetalClusterManagedByUser GPUBaremetalClusterManagedBy = "user"
)

type GPUBaremetalClusterServersSettings struct {
	// List of file shares mounted across the cluster.
	FileShares []GPUBaremetalClusterServersSettingsFileShare      `json:"file_shares,required"`
	Interfaces []GPUBaremetalClusterServersSettingsInterfaceUnion `json:"interfaces,required"`
	// Security groups
	SecurityGroups []GPUBaremetalClusterServersSettingsSecurityGroup `json:"security_groups,required"`
	// SSH key name
	SSHKeyName string `json:"ssh_key_name,required"`
	// Optional custom user data
	UserData string `json:"user_data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileShares     respjson.Field
		Interfaces     respjson.Field
		SecurityGroups respjson.Field
		SSHKeyName     respjson.Field
		UserData       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServersSettings) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServersSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServersSettingsFileShare struct {
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
func (r GPUBaremetalClusterServersSettingsFileShare) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServersSettingsFileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUBaremetalClusterServersSettingsInterfaceUnion contains all possible
// properties and values from
// [GPUBaremetalClusterServersSettingsInterfaceExternal],
// [GPUBaremetalClusterServersSettingsInterfaceSubnet],
// [GPUBaremetalClusterServersSettingsInterfaceAnySubnet].
//
// Use the [GPUBaremetalClusterServersSettingsInterfaceUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type GPUBaremetalClusterServersSettingsInterfaceUnion struct {
	IPFamily string `json:"ip_family"`
	Name     string `json:"name"`
	// Any of "external", "subnet", "any_subnet".
	Type string `json:"type"`
	// This field is a union of
	// [GPUBaremetalClusterServersSettingsInterfaceSubnetFloatingIP],
	// [GPUBaremetalClusterServersSettingsInterfaceAnySubnetFloatingIP]
	FloatingIP GPUBaremetalClusterServersSettingsInterfaceUnionFloatingIP `json:"floating_ip"`
	NetworkID  string                                                     `json:"network_id"`
	// This field is from variant [GPUBaremetalClusterServersSettingsInterfaceSubnet].
	SubnetID string `json:"subnet_id"`
	// This field is from variant
	// [GPUBaremetalClusterServersSettingsInterfaceAnySubnet].
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

// anyGPUBaremetalClusterServersSettingsInterface is implemented by each variant of
// [GPUBaremetalClusterServersSettingsInterfaceUnion] to add type safety for the
// return type of [GPUBaremetalClusterServersSettingsInterfaceUnion.AsAny]
type anyGPUBaremetalClusterServersSettingsInterface interface {
	implGPUBaremetalClusterServersSettingsInterfaceUnion()
}

func (GPUBaremetalClusterServersSettingsInterfaceExternal) implGPUBaremetalClusterServersSettingsInterfaceUnion() {
}
func (GPUBaremetalClusterServersSettingsInterfaceSubnet) implGPUBaremetalClusterServersSettingsInterfaceUnion() {
}
func (GPUBaremetalClusterServersSettingsInterfaceAnySubnet) implGPUBaremetalClusterServersSettingsInterfaceUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := GPUBaremetalClusterServersSettingsInterfaceUnion.AsAny().(type) {
//	case cloud.GPUBaremetalClusterServersSettingsInterfaceExternal:
//	case cloud.GPUBaremetalClusterServersSettingsInterfaceSubnet:
//	case cloud.GPUBaremetalClusterServersSettingsInterfaceAnySubnet:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u GPUBaremetalClusterServersSettingsInterfaceUnion) AsAny() anyGPUBaremetalClusterServersSettingsInterface {
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

func (u GPUBaremetalClusterServersSettingsInterfaceUnion) AsExternal() (v GPUBaremetalClusterServersSettingsInterfaceExternal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GPUBaremetalClusterServersSettingsInterfaceUnion) AsSubnet() (v GPUBaremetalClusterServersSettingsInterfaceSubnet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GPUBaremetalClusterServersSettingsInterfaceUnion) AsAnySubnet() (v GPUBaremetalClusterServersSettingsInterfaceAnySubnet) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GPUBaremetalClusterServersSettingsInterfaceUnion) RawJSON() string { return u.JSON.raw }

func (r *GPUBaremetalClusterServersSettingsInterfaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPUBaremetalClusterServersSettingsInterfaceUnionFloatingIP is an implicit
// subunion of [GPUBaremetalClusterServersSettingsInterfaceUnion].
// GPUBaremetalClusterServersSettingsInterfaceUnionFloatingIP provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [GPUBaremetalClusterServersSettingsInterfaceUnion].
type GPUBaremetalClusterServersSettingsInterfaceUnionFloatingIP struct {
	// This field is from variant
	// [GPUBaremetalClusterServersSettingsInterfaceSubnetFloatingIP].
	Source constant.New `json:"source"`
	JSON   struct {
		Source respjson.Field
		raw    string
	} `json:"-"`
}

func (r *GPUBaremetalClusterServersSettingsInterfaceUnionFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServersSettingsInterfaceExternal struct {
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
func (r GPUBaremetalClusterServersSettingsInterfaceExternal) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServersSettingsInterfaceExternal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServersSettingsInterfaceSubnet struct {
	// Floating IP config for this subnet attachment
	FloatingIP GPUBaremetalClusterServersSettingsInterfaceSubnetFloatingIP `json:"floating_ip,required"`
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
func (r GPUBaremetalClusterServersSettingsInterfaceSubnet) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServersSettingsInterfaceSubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Floating IP config for this subnet attachment
type GPUBaremetalClusterServersSettingsInterfaceSubnetFloatingIP struct {
	Source constant.New `json:"source,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServersSettingsInterfaceSubnetFloatingIP) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalClusterServersSettingsInterfaceSubnetFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServersSettingsInterfaceAnySubnet struct {
	// Floating IP config for this subnet attachment
	FloatingIP GPUBaremetalClusterServersSettingsInterfaceAnySubnetFloatingIP `json:"floating_ip,required"`
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
func (r GPUBaremetalClusterServersSettingsInterfaceAnySubnet) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServersSettingsInterfaceAnySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Floating IP config for this subnet attachment
type GPUBaremetalClusterServersSettingsInterfaceAnySubnetFloatingIP struct {
	Source constant.New `json:"source,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServersSettingsInterfaceAnySubnetFloatingIP) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalClusterServersSettingsInterfaceAnySubnetFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServersSettingsSecurityGroup struct {
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
func (r GPUBaremetalClusterServersSettingsSecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServersSettingsSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cluster status
type GPUBaremetalClusterStatus string

const (
	GPUBaremetalClusterStatusActive     GPUBaremetalClusterStatus = "active"
	GPUBaremetalClusterStatusCreating   GPUBaremetalClusterStatus = "creating"
	GPUBaremetalClusterStatusDegraded   GPUBaremetalClusterStatus = "degraded"
	GPUBaremetalClusterStatusDeleting   GPUBaremetalClusterStatus = "deleting"
	GPUBaremetalClusterStatusError      GPUBaremetalClusterStatus = "error"
	GPUBaremetalClusterStatusNew        GPUBaremetalClusterStatus = "new"
	GPUBaremetalClusterStatusRebooting  GPUBaremetalClusterStatus = "rebooting"
	GPUBaremetalClusterStatusRebuilding GPUBaremetalClusterStatus = "rebuilding"
	GPUBaremetalClusterStatusResizing   GPUBaremetalClusterStatus = "resizing"
	GPUBaremetalClusterStatusShutoff    GPUBaremetalClusterStatus = "shutoff"
)

type GPUBaremetalClusterNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Cluster flavor ID
	Flavor string `json:"flavor,required"`
	// System image ID
	ImageID string `json:"image_id,required"`
	// Cluster name
	Name string `json:"name,required"`
	// Number of servers in the cluster
	ServersCount int64 `json:"servers_count,required"`
	// Configuration settings for the servers in the cluster
	ServersSettings GPUBaremetalClusterNewParamsServersSettings `json:"servers_settings,omitzero,required"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Some tags are read-only and cannot be modified by the
	// user. Tags are also integrated with cost reports, allowing cost data to be
	// filtered based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration settings for the servers in the cluster
//
// The property Interfaces is required.
type GPUBaremetalClusterNewParamsServersSettings struct {
	// Subnet IPs and floating IPs
	Interfaces []GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion `json:"interfaces,omitzero,required"`
	// Optional custom user data (Base64-encoded)
	UserData param.Opt[string] `json:"user_data,omitzero"`
	// Optional server access credentials
	Credentials GPUBaremetalClusterNewParamsServersSettingsCredentials `json:"credentials,omitzero"`
	// List of file shares to be mounted across the cluster.
	FileShares []GPUBaremetalClusterNewParamsServersSettingsFileShare `json:"file_shares,omitzero"`
	// List of security groups UUIDs
	SecurityGroups []GPUBaremetalClusterNewParamsServersSettingsSecurityGroup `json:"security_groups,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsServersSettings) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion struct {
	OfExternal  *GPUBaremetalClusterNewParamsServersSettingsInterfaceExternal  `json:",omitzero,inline"`
	OfSubnet    *GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnet    `json:",omitzero,inline"`
	OfAnySubnet *GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnet `json:",omitzero,inline"`
	paramUnion
}

func (u GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExternal, u.OfSubnet, u.OfAnySubnet)
}
func (u *GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion) asAny() any {
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
func (u GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion) GetSubnetID() *string {
	if vt := u.OfSubnet; vt != nil {
		return &vt.SubnetID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion) GetType() *string {
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
func (u GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion) GetIPFamily() *string {
	if vt := u.OfExternal; vt != nil {
		return (*string)(&vt.IPFamily)
	} else if vt := u.OfAnySubnet; vt != nil {
		return (*string)(&vt.IPFamily)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion) GetName() *string {
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
func (u GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion) GetNetworkID() *string {
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
func (u GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion) GetFloatingIP() (res gpuBaremetalClusterNewParamsServersSettingsInterfaceUnionFloatingIP) {
	if vt := u.OfSubnet; vt != nil {
		res.any = &vt.FloatingIP
	} else if vt := u.OfAnySubnet; vt != nil {
		res.any = &vt.FloatingIP
	}
	return
}

// Can have the runtime types
// [*GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP],
// [*GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP]
type gpuBaremetalClusterNewParamsServersSettingsInterfaceUnionFloatingIP struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cloud.GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP:
//	case *cloud.GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u gpuBaremetalClusterNewParamsServersSettingsInterfaceUnionFloatingIP) AsAny() any {
	return u.any
}

// Returns a pointer to the underlying variant's property, if present.
func (u gpuBaremetalClusterNewParamsServersSettingsInterfaceUnionFloatingIP) GetSource() *string {
	switch vt := u.any.(type) {
	case *GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP:
		return (*string)(&vt.Source)
	case *GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP:
		return (*string)(&vt.Source)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion](
		"type",
		apijson.Discriminator[GPUBaremetalClusterNewParamsServersSettingsInterfaceExternal]("external"),
		apijson.Discriminator[GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnet]("subnet"),
		apijson.Discriminator[GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnet]("any_subnet"),
	)
}

// The property Type is required.
type GPUBaremetalClusterNewParamsServersSettingsInterfaceExternal struct {
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

func (r GPUBaremetalClusterNewParamsServersSettingsInterfaceExternal) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsInterfaceExternal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsInterfaceExternal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GPUBaremetalClusterNewParamsServersSettingsInterfaceExternal](
		"ip_family", "dual", "ipv4", "ipv6",
	)
}

// The properties NetworkID, SubnetID, Type are required.
type GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnet struct {
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID string `json:"network_id,required"`
	// Port is assigned an IP address from this subnet
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Interface name
	Name param.Opt[string] `json:"name,omitzero"`
	// Floating IP config for this subnet attachment
	FloatingIP GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP `json:"floating_ip,omitzero"`
	// This field can be elided, and will marshal its zero value as "subnet".
	Type constant.Subnet `json:"type,required"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnet) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewGPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP() GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP {
	return GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP{
		Source: "new",
	}
}

// Floating IP config for this subnet attachment
//
// This struct has a constant value, construct it with
// [NewGPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP].
type GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP struct {
	Source constant.New `json:"source,required"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties NetworkID, Type are required.
type GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnet struct {
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID string `json:"network_id,required"`
	// Interface name
	Name param.Opt[string] `json:"name,omitzero"`
	// Floating IP config for this subnet attachment
	FloatingIP GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP `json:"floating_ip,omitzero"`
	// Which subnets should be selected: IPv4, IPv6, or use dual stack
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// This field can be elided, and will marshal its zero value as "any_subnet".
	Type constant.AnySubnet `json:"type,required"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnet) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnet](
		"ip_family", "dual", "ipv4", "ipv6",
	)
}

func NewGPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP() GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP {
	return GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP{
		Source: "new",
	}
}

// Floating IP config for this subnet attachment
//
// This struct has a constant value, construct it with
// [NewGPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP].
type GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP struct {
	Source constant.New `json:"source,required"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional server access credentials
type GPUBaremetalClusterNewParamsServersSettingsCredentials struct {
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

func (r GPUBaremetalClusterNewParamsServersSettingsCredentials) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsCredentials
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsCredentials) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, MountPath are required.
type GPUBaremetalClusterNewParamsServersSettingsFileShare struct {
	// Unique identifier of the file share in UUID format.
	ID string `json:"id,required" format:"uuid4"`
	// Absolute mount path inside the system where the file share will be mounted.
	MountPath string `json:"mount_path,required"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsServersSettingsFileShare) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsFileShare
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsFileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type GPUBaremetalClusterNewParamsServersSettingsSecurityGroup struct {
	// Resource ID
	ID string `json:"id,required" format:"uuid4"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsServersSettingsSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Limit of items on a single page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset in results list
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Specifies the entity responsible for managing the resource.
	//
	//   - `user`: The resource (cluster) is created and maintained directly by the user.
	//   - `k8s`: The resource is created and maintained automatically by Managed
	//     Kubernetes service
	//
	// Any of "k8s", "user".
	ManagedBy []string `query:"managed_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GPUBaremetalClusterListParams]'s query parameters as
// `url.Values`.
func (r GPUBaremetalClusterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type GPUBaremetalClusterDeleteParams struct {
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
	// Optional list of floating ips to be deleted
	FloatingIPIDs []string `query:"floating_ip_ids,omitzero" format:"uuid4" json:"-"`
	// Optional list of reserved fixed ips to be deleted
	ReservedFixedIPIDs []string `query:"reserved_fixed_ip_ids,omitzero" format:"uuid4" json:"-"`
	paramObj
}

// URLQuery serializes [GPUBaremetalClusterDeleteParams]'s query parameters as
// `url.Values`.
func (r GPUBaremetalClusterDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type GPUBaremetalClusterActionParams struct {
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

func (r GPUBaremetalClusterActionParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterActionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterActionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type GPUBaremetalClusterPowercycleAllServersParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type GPUBaremetalClusterRebootAllServersParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type GPUBaremetalClusterRebuildParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// List of nodes uuids to be rebuild
	Nodes []string `json:"nodes,omitzero,required"`
	// AI GPU image ID
	ImageID param.Opt[string] `json:"image_id,omitzero"`
	// String in base64 format.Examples of the `user_data`:
	// https://cloudinit.readthedocs.io/en/latest/topics/examples.html
	UserData param.Opt[string] `json:"user_data,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterRebuildParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterRebuildParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterRebuildParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterResizeParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Resized (total) number of instances
	InstancesCount int64 `json:"instances_count,required"`
	paramObj
}

func (r GPUBaremetalClusterResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterResizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
