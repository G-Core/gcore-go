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
	// GPU bare metal images are custom boot images for bare metal GPU servers.
	Images GPUBaremetalClusterImageService
	tasks  TaskService
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
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/clusters", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update the name, tags, and/or server settings of an existing bare metal GPU
// cluster.
//
// Update tags using JSON Merge Patch semantics (RFC 7396). To add or update tags,
// provide key-value pairs. To remove a tag, set its value to null.
//
// Updating server settings (`servers_settings`, `image_id`) only modifies the
// cluster template. **It does NOT modify or rebuild any existing servers in the
// cluster.** Tags and name will be applied immediately. To apply the rest changes
// to running servers, use the
// `/cloud/v3/gpu/baremetal/{project_id}/{region_id}/clusters/{cluster_id}/apply_settings`
// endpoint.
func (r *GPUBaremetalClusterService) Update(ctx context.Context, clusterID string, params GPUBaremetalClusterUpdateParams, opts ...option.RequestOption) (res *GPUBaremetalCluster, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/clusters/%s", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List all bare metal GPU clusters in the specified project and region.
func (r *GPUBaremetalClusterService) List(ctx context.Context, params GPUBaremetalClusterListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[GPUBaremetalCluster], err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/clusters/%s", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Get detailed information about a specific bare metal GPU cluster.
func (r *GPUBaremetalClusterService) Get(ctx context.Context, clusterID string, query GPUBaremetalClusterGetParams, opts ...option.RequestOption) (res *GPUBaremetalCluster, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/clusters/%s", query.ProjectID.Value, query.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Please use the
// `/v3/gpu/baremetal/{project_id}/{region_id}/clusters/{cluster_id}/action`
// instead.
//
// Deprecated: deprecated
func (r *GPUBaremetalClusterService) PowercycleAllServers(ctx context.Context, clusterID string, body GPUBaremetalClusterPowercycleAllServersParams, opts ...option.RequestOption) (res *GPUBaremetalClusterServerV1List, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s/powercycle", body.ProjectID.Value, body.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Please use the
// `/v3/gpu/baremetal/{project_id}/{region_id}/clusters/{cluster_id}/action`
// instead.
//
// Deprecated: deprecated
func (r *GPUBaremetalClusterService) RebootAllServers(ctx context.Context, clusterID string, body GPUBaremetalClusterRebootAllServersParams, opts ...option.RequestOption) (res *GPUBaremetalClusterServerV1List, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s/reboot", body.ProjectID.Value, body.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Use
// `POST /v3/gpu/baremetal/{project_id}/{region_id}/clusters/{cluster_id}/apply_settings`
// instead.
//
// Perform a rebuild operation on a bare metal GPU cluster. During the rebuild
// process, the servers in cluster receive a new image, SSH key, and user data.
// Important: Before triggering a rebuild, the cluster must have updated server
// settings to apply. These cluster settings must be patched using the following
// endpoint: PATCH
// '/v3/gpu/baremetal/{`project_id`}/{`region_id`}/clusters/{`cluster_id`}'
//
// Deprecated: deprecated
func (r *GPUBaremetalClusterService) Rebuild(ctx context.Context, clusterID string, body GPUBaremetalClusterRebuildParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/clusters/%s/rebuild", body.ProjectID.Value, body.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Change the number of nodes in a GPU cluster. The cluster can be scaled up or
// down.
func (r *GPUBaremetalClusterService) Resize(ctx context.Context, clusterID string, params GPUBaremetalClusterResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/gpu/%v/%v/%s/resize", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// NewAndPoll creates a new GPU bare metal cluster and polls for completion. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *GPUBaremetalClusterService) NewAndPoll(ctx context.Context, params GPUBaremetalClusterNewParams, opts ...option.RequestOption) (v *GPUBaremetalCluster, err error) {
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
	var getParams GPUBaremetalClusterGetParams
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

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Clusters) != 1 {
		return nil, errors.New("expected exactly one cluster to be created in a task")
	}
	clusterID := task.CreatedResources.Clusters[0]

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, clusterID, getParams, getOpts...)
}

// RebuildAndPoll rebuilds a GPU bare metal cluster and polls for completion of the first task. Use the
// [TaskService.Poll] method if you need to poll for all tasks.
func (r *GPUBaremetalClusterService) RebuildAndPoll(ctx context.Context, clusterID string, params GPUBaremetalClusterRebuildParams, opts ...option.RequestOption) (v *GPUBaremetalCluster, err error) {
	// Exclude WithResponseBodyInto for the action (Rebuild returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Rebuild(ctx, clusterID, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
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
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	_, err = r.tasks.Poll(ctx, taskID, pollOpts...)
	if err != nil {
		return
	}

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, clusterID, getParams, getOpts...)
}

// ResizeAndPoll resizes a GPU bare metal cluster and polls for completion of the first task. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *GPUBaremetalClusterService) ResizeAndPoll(ctx context.Context, clusterID string, params GPUBaremetalClusterResizeParams, opts ...option.RequestOption) (v *GPUBaremetalCluster, err error) {
	// Exclude WithResponseBodyInto for the action (Resize returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Resize(ctx, clusterID, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
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
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	_, err = r.tasks.Poll(ctx, taskID, pollOpts...)
	if err != nil {
		return
	}

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, clusterID, getParams, getOpts...)
}

// This operation only modifies cluster settings such as SSH key, image, and user
// data. **It does NOT modify or rebuild any existing servers in the cluster.**
// Use `PATCH /v3/gpu/baremetal/{project_id}/{region_id}/clusters/{cluster_id}`
// instead, which now accepts `servers_settings` and `image_id` fields alongside
// `name` and `tags`.
//
// To apply these configuration changes to running servers, use the
// `/cloud/v3/gpu/baremetal/{project_id}/{region_id}/clusters/{cluster_id}/apply_settings`
// endpoint.
//
// Deprecated: deprecated
func (r *GPUBaremetalClusterService) UpdateServersSettings(ctx context.Context, clusterID string, params GPUBaremetalClusterUpdateServersSettingsParams, opts ...option.RequestOption) (res *GPUBaremetalCluster, err error) {
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
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/clusters/%s/servers_settings", params.ProjectID.Value, params.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// DeleteAndPoll deletes a bare metal GPU cluster and polls for completion of the first task. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *GPUBaremetalClusterService) DeleteAndPoll(ctx context.Context, clusterID string, params GPUBaremetalClusterDeleteParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (Delete returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Delete(ctx, clusterID, params, actionOpts...)
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

type GPUBaremetalCluster struct {
	// Cluster unique identifier
	ID string `json:"id" api:"required" format:"uuid4"`
	// Cluster creation date time
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Cluster flavor name
	Flavor string `json:"flavor" api:"required"`
	// True if any server in the cluster has pending (not yet applied) settings changes
	HasPendingChanges bool `json:"has_pending_changes" api:"required"`
	// Image ID
	ImageID string `json:"image_id" api:"required"`
	// User type managing the resource
	//
	// Any of "k8s", "user".
	ManagedBy GPUBaremetalClusterManagedBy `json:"managed_by" api:"required"`
	// Cluster name
	Name string `json:"name" api:"required"`
	// Cluster servers count
	ServersCount int64 `json:"servers_count" api:"required"`
	// List of cluster nodes
	ServersIDs      []string                           `json:"servers_ids" api:"required" format:"uuid4"`
	ServersSettings GPUBaremetalClusterServersSettings `json:"servers_settings" api:"required"`
	// Cluster status
	//
	// Any of "active", "creating", "degraded", "deleting", "error", "rebooting",
	// "rebuilding", "resizing", "shutoff".
	Status GPUBaremetalClusterStatus `json:"status" api:"required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags" api:"required"`
	// Cluster update date time
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		Flavor            respjson.Field
		HasPendingChanges respjson.Field
		ImageID           respjson.Field
		ManagedBy         respjson.Field
		Name              respjson.Field
		ServersCount      respjson.Field
		ServersIDs        respjson.Field
		ServersSettings   respjson.Field
		Status            respjson.Field
		Tags              respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
	FileShares []GPUBaremetalClusterServersSettingsFileShare      `json:"file_shares" api:"required"`
	Interfaces []GPUBaremetalClusterServersSettingsInterfaceUnion `json:"interfaces" api:"required"`
	// Deprecated. Deduplicated union of security groups across all interfaces; the
	// actual assignment may differ per interface. Use `interfaces[].security_groups`
	// for the authoritative per-interface list.
	//
	// Deprecated: deprecated
	SecurityGroups []GPUBaremetalClusterServersSettingsSecurityGroup `json:"security_groups" api:"required"`
	// SSH key name
	SSHKeyName string `json:"ssh_key_name" api:"required"`
	// Optional custom user data
	UserData string `json:"user_data" api:"required"`
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
	ID string `json:"id" api:"required" format:"uuid4"`
	// Absolute mount path inside the system where the file share will be mounted.
	MountPath string `json:"mount_path" api:"required"`
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
	// This field is a union of
	// [[]GPUBaremetalClusterServersSettingsInterfaceExternalSecurityGroup],
	// [[]GPUBaremetalClusterServersSettingsInterfaceSubnetSecurityGroup],
	// [[]GPUBaremetalClusterServersSettingsInterfaceAnySubnetSecurityGroup]
	SecurityGroups GPUBaremetalClusterServersSettingsInterfaceUnionSecurityGroups `json:"security_groups"`
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
		IPFamily       respjson.Field
		Name           respjson.Field
		SecurityGroups respjson.Field
		Type           respjson.Field
		FloatingIP     respjson.Field
		NetworkID      respjson.Field
		SubnetID       respjson.Field
		IPAddress      respjson.Field
		raw            string
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

// GPUBaremetalClusterServersSettingsInterfaceUnionSecurityGroups is an implicit
// subunion of [GPUBaremetalClusterServersSettingsInterfaceUnion].
// GPUBaremetalClusterServersSettingsInterfaceUnionSecurityGroups provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [GPUBaremetalClusterServersSettingsInterfaceUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSecurityGroups]
type GPUBaremetalClusterServersSettingsInterfaceUnionSecurityGroups struct {
	// This field will be present if the value is a
	// [[]GPUBaremetalClusterServersSettingsInterfaceExternalSecurityGroup] instead of
	// an object.
	OfSecurityGroups []GPUBaremetalClusterServersSettingsInterfaceExternalSecurityGroup `json:",inline"`
	JSON             struct {
		OfSecurityGroups respjson.Field
		raw              string
	} `json:"-"`
}

func (r *GPUBaremetalClusterServersSettingsInterfaceUnionSecurityGroups) UnmarshalJSON(data []byte) error {
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
	IPFamily string `json:"ip_family" api:"required"`
	// Interface name
	Name string `json:"name" api:"required"`
	// Resolved security groups applied to this interface.
	SecurityGroups []GPUBaremetalClusterServersSettingsInterfaceExternalSecurityGroup `json:"security_groups" api:"required"`
	Type           constant.External                                                  `json:"type" default:"external"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPFamily       respjson.Field
		Name           respjson.Field
		SecurityGroups respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServersSettingsInterfaceExternal) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServersSettingsInterfaceExternal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServersSettingsInterfaceExternalSecurityGroup struct {
	// Security group ID
	ID string `json:"id" api:"required"`
	// Security group name
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServersSettingsInterfaceExternalSecurityGroup) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalClusterServersSettingsInterfaceExternalSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServersSettingsInterfaceSubnet struct {
	// Floating IP config for this subnet attachment
	FloatingIP GPUBaremetalClusterServersSettingsInterfaceSubnetFloatingIP `json:"floating_ip" api:"required"`
	// Interface name
	Name string `json:"name" api:"required"`
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID string `json:"network_id" api:"required"`
	// Resolved security groups applied to this interface.
	SecurityGroups []GPUBaremetalClusterServersSettingsInterfaceSubnetSecurityGroup `json:"security_groups" api:"required"`
	// Port is assigned an IP address from this subnet
	SubnetID string          `json:"subnet_id" api:"required" format:"uuid4"`
	Type     constant.Subnet `json:"type" default:"subnet"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FloatingIP     respjson.Field
		Name           respjson.Field
		NetworkID      respjson.Field
		SecurityGroups respjson.Field
		SubnetID       respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServersSettingsInterfaceSubnet) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServersSettingsInterfaceSubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Floating IP config for this subnet attachment
type GPUBaremetalClusterServersSettingsInterfaceSubnetFloatingIP struct {
	Source constant.New `json:"source" default:"new"`
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

type GPUBaremetalClusterServersSettingsInterfaceSubnetSecurityGroup struct {
	// Security group ID
	ID string `json:"id" api:"required"`
	// Security group name
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServersSettingsInterfaceSubnetSecurityGroup) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalClusterServersSettingsInterfaceSubnetSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServersSettingsInterfaceAnySubnet struct {
	// Floating IP config for this subnet attachment
	FloatingIP GPUBaremetalClusterServersSettingsInterfaceAnySubnetFloatingIP `json:"floating_ip" api:"required"`
	// Fixed IP address
	IPAddress string `json:"ip_address" api:"required"`
	// Which subnets should be selected: IPv4, IPv6, or use dual stack
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family" api:"required"`
	// Interface name
	Name string `json:"name" api:"required"`
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID string `json:"network_id" api:"required"`
	// Resolved security groups applied to this interface.
	SecurityGroups []GPUBaremetalClusterServersSettingsInterfaceAnySubnetSecurityGroup `json:"security_groups" api:"required"`
	Type           constant.AnySubnet                                                  `json:"type" default:"any_subnet"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FloatingIP     respjson.Field
		IPAddress      respjson.Field
		IPFamily       respjson.Field
		Name           respjson.Field
		NetworkID      respjson.Field
		SecurityGroups respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServersSettingsInterfaceAnySubnet) RawJSON() string { return r.JSON.raw }
func (r *GPUBaremetalClusterServersSettingsInterfaceAnySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Floating IP config for this subnet attachment
type GPUBaremetalClusterServersSettingsInterfaceAnySubnetFloatingIP struct {
	Source constant.New `json:"source" default:"new"`
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

type GPUBaremetalClusterServersSettingsInterfaceAnySubnetSecurityGroup struct {
	// Security group ID
	ID string `json:"id" api:"required"`
	// Security group name
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GPUBaremetalClusterServersSettingsInterfaceAnySubnetSecurityGroup) RawJSON() string {
	return r.JSON.raw
}
func (r *GPUBaremetalClusterServersSettingsInterfaceAnySubnetSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterServersSettingsSecurityGroup struct {
	// Security group ID
	ID string `json:"id" api:"required"`
	// Security group name
	Name string `json:"name" api:"required"`
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
	GPUBaremetalClusterStatusRebooting  GPUBaremetalClusterStatus = "rebooting"
	GPUBaremetalClusterStatusRebuilding GPUBaremetalClusterStatus = "rebuilding"
	GPUBaremetalClusterStatusResizing   GPUBaremetalClusterStatus = "resizing"
	GPUBaremetalClusterStatusShutoff    GPUBaremetalClusterStatus = "shutoff"
)

type GPUBaremetalClusterNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Cluster flavor ID
	Flavor string `json:"flavor" api:"required"`
	// System image ID
	ImageID string `json:"image_id" api:"required"`
	// Cluster name
	Name string `json:"name" api:"required"`
	// Number of servers in the cluster
	ServersCount int64 `json:"servers_count" api:"required"`
	// Configuration settings for the servers in the cluster
	ServersSettings GPUBaremetalClusterNewParamsServersSettings `json:"servers_settings,omitzero" api:"required"`
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
	Interfaces []GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion `json:"interfaces,omitzero" api:"required"`
	// Optional custom user data (Base64-encoded)
	UserData param.Opt[string] `json:"user_data,omitzero"`
	// Optional server access credentials
	Credentials GPUBaremetalClusterNewParamsServersSettingsCredentials `json:"credentials,omitzero"`
	// List of file shares to be mounted across the cluster.
	FileShares []GPUBaremetalClusterNewParamsServersSettingsFileShare `json:"file_shares,omitzero"`
	// Deprecated. Use per-interface `security_groups` inside `interfaces[]` instead.
	// Cannot be combined with per-interface `security_groups`. If omitted everywhere,
	// the project's default security group is applied.
	//
	// Deprecated: deprecated
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
func (u GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion) GetPortSecurityEnabled() *bool {
	if vt := u.OfExternal; vt != nil && vt.PortSecurityEnabled.Valid() {
		return &vt.PortSecurityEnabled.Value
	} else if vt := u.OfSubnet; vt != nil && vt.PortSecurityEnabled.Valid() {
		return &vt.PortSecurityEnabled.Value
	} else if vt := u.OfAnySubnet; vt != nil && vt.PortSecurityEnabled.Valid() {
		return &vt.PortSecurityEnabled.Value
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
func (u GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion) GetSecurityGroups() (res gpuBaremetalClusterNewParamsServersSettingsInterfaceUnionSecurityGroups) {
	if vt := u.OfExternal; vt != nil {
		res.any = &vt.SecurityGroups
	} else if vt := u.OfSubnet; vt != nil {
		res.any = &vt.SecurityGroups
	} else if vt := u.OfAnySubnet; vt != nil {
		res.any = &vt.SecurityGroups
	}
	return
}

// Can have the runtime types
// [_[]GPUBaremetalClusterNewParamsServersSettingsInterfaceExternalSecurityGroup],
// [_[]GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetSecurityGroup],
// [\*[]GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetSecurityGroup]
type gpuBaremetalClusterNewParamsServersSettingsInterfaceUnionSecurityGroups struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]cloud.GPUBaremetalClusterNewParamsServersSettingsInterfaceExternalSecurityGroup:
//	case *[]cloud.GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetSecurityGroup:
//	case *[]cloud.GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetSecurityGroup:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u gpuBaremetalClusterNewParamsServersSettingsInterfaceUnionSecurityGroups) AsAny() any {
	return u.any
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

func init() {
	apijson.RegisterFieldValidator[GPUBaremetalClusterNewParamsServersSettingsInterfaceExternal](
		"ip_family", "dual", "ipv4", "ipv6",
	)
}

func init() {
	apijson.RegisterFieldValidator[GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnet](
		"ip_family", "dual", "ipv4", "ipv6",
	)
}

// The property Type is required.
type GPUBaremetalClusterNewParamsServersSettingsInterfaceExternal struct {
	// Interface name
	Name param.Opt[string] `json:"name,omitzero"`
	// Controls port security for this interface. When omitted, the default applies
	// (port security enabled, default security group attached). When false, the port
	// is created with port security off and no security group attached;
	// `security_groups` must not be set in that case. Not allowed for interfaces on a
	// public network, nor for bare metal servers without a DPU (their ports cannot
	// enforce port security).
	PortSecurityEnabled param.Opt[bool] `json:"port_security_enabled,omitzero"`
	// Which subnets should be selected: IPv4, IPv6, or use dual stack.
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// Security group UUIDs applied to this interface. If omitted (or empty), the
	// top-level `security_groups` value applies; if both are omitted, the project's
	// default security group is applied.
	SecurityGroups []GPUBaremetalClusterNewParamsServersSettingsInterfaceExternalSecurityGroup `json:"security_groups,omitzero"`
	// This field can be elided, and will marshal its zero value as "external".
	Type constant.External `json:"type" default:"external"`
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

// The property ID is required.
type GPUBaremetalClusterNewParamsServersSettingsInterfaceExternalSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsServersSettingsInterfaceExternalSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsInterfaceExternalSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsInterfaceExternalSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties NetworkID, SubnetID, Type are required.
type GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnet struct {
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID string `json:"network_id" api:"required"`
	// Port is assigned an IP address from this subnet
	SubnetID string `json:"subnet_id" api:"required" format:"uuid4"`
	// Interface name
	Name param.Opt[string] `json:"name,omitzero"`
	// Controls port security for this interface. When omitted, the default applies
	// (port security enabled, default security group attached). When false, the port
	// is created with port security off and no security group attached;
	// `security_groups` must not be set in that case. Not allowed for interfaces on a
	// public network, nor for bare metal servers without a DPU (their ports cannot
	// enforce port security).
	PortSecurityEnabled param.Opt[bool] `json:"port_security_enabled,omitzero"`
	// Floating IP config for this subnet attachment
	FloatingIP GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP `json:"floating_ip,omitzero"`
	// Security group UUIDs applied to this interface. If omitted (or empty), the
	// top-level `security_groups` value applies; if both are omitted, the project's
	// default security group is applied.
	SecurityGroups []GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetSecurityGroup `json:"security_groups,omitzero"`
	// This field can be elided, and will marshal its zero value as "subnet".
	Type constant.Subnet `json:"type" default:"subnet"`
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
	Source constant.New `json:"source" default:"new"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsInterfaceSubnetSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties NetworkID, Type are required.
type GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnet struct {
	// Network ID the subnet belongs to. Port will be plugged in this network
	NetworkID string `json:"network_id" api:"required"`
	// Interface name
	Name param.Opt[string] `json:"name,omitzero"`
	// Controls port security for this interface. When omitted, the default applies
	// (port security enabled, default security group attached). When false, the port
	// is created with port security off and no security group attached;
	// `security_groups` must not be set in that case. Not allowed for interfaces on a
	// public network, nor for bare metal servers without a DPU (their ports cannot
	// enforce port security).
	PortSecurityEnabled param.Opt[bool] `json:"port_security_enabled,omitzero"`
	// Floating IP config for this subnet attachment
	FloatingIP GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP `json:"floating_ip,omitzero"`
	// Which subnets should be selected: IPv4, IPv6, or use dual stack
	//
	// Any of "dual", "ipv4", "ipv6".
	IPFamily string `json:"ip_family,omitzero"`
	// Security group UUIDs applied to this interface. If omitted (or empty), the
	// top-level `security_groups` value applies; if both are omitted, the project's
	// default security group is applied.
	SecurityGroups []GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetSecurityGroup `json:"security_groups,omitzero"`
	// This field can be elided, and will marshal its zero value as "any_subnet".
	Type constant.AnySubnet `json:"type" default:"any_subnet"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnet) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	Source constant.New `json:"source" default:"new"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetFloatingIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetSecurityGroup struct {
	// Resource ID
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsInterfaceAnySubnetSecurityGroup) UnmarshalJSON(data []byte) error {
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
	ID string `json:"id" api:"required" format:"uuid4"`
	// Absolute mount path inside the system where the file share will be mounted.
	MountPath string `json:"mount_path" api:"required"`
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
	ID string `json:"id" api:"required" format:"uuid4"`
	paramObj
}

func (r GPUBaremetalClusterNewParamsServersSettingsSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterNewParamsServersSettingsSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterNewParamsServersSettingsSecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Image ID of the OS image to apply to the cluster template. Use GET
	// /v1/images/{`project_id`}/{`region_id`} to discover available images. Takes
	// effect on existing servers only after a successful POST /`apply_settings` call.
	ImageID param.Opt[string] `json:"image_id,omitzero" format:"uuid4"`
	// Cluster name
	Name param.Opt[string] `json:"name,omitzero"`
	// Configuration settings for the servers in the cluster
	ServersSettings GPUBaremetalClusterUpdateParamsServersSettings `json:"servers_settings,omitzero"`
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

func (r GPUBaremetalClusterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration settings for the servers in the cluster
type GPUBaremetalClusterUpdateParamsServersSettings struct {
	// Optional custom user data (Base64-encoded)
	UserData param.Opt[string] `json:"user_data,omitzero"`
	// Optional server access credentials
	Credentials GPUBaremetalClusterUpdateParamsServersSettingsCredentials `json:"credentials,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterUpdateParamsServersSettings) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterUpdateParamsServersSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterUpdateParamsServersSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional server access credentials
type GPUBaremetalClusterUpdateParamsServersSettingsCredentials struct {
	// Specifies the name of the SSH keypair, created via the
	// [/v1/`ssh_keys` endpoint](/docs/api-reference/cloud/ssh-keys/add-or-generate-ssh-key).
	SSHKeyName param.Opt[string] `json:"ssh_key_name,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterUpdateParamsServersSettingsCredentials) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterUpdateParamsServersSettingsCredentials
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterUpdateParamsServersSettingsCredentials) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Limit of items on a single page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset in results list
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by creation time (UTC), e.g. `created_at[gte]=2026-01-01T00:00:00Z`.
	CreatedAt GPUBaremetalClusterListParamsCreatedAt `query:"created_at,omitzero" json:"-"`
	// Filter by flavor (case-insensitive), e.g. `flavor[prefix]=bm3-`,
	// `flavor[exact]=bm3-ai-1xlarge-h100-80-8`.
	Flavor GPUBaremetalClusterListParamsFlavor `query:"flavor,omitzero" json:"-"`
	// Return only clusters with these IDs, e.g. `ids=<id1>&ids=<id2>`.
	IDs []string `query:"ids,omitzero" format:"uuid4" json:"-"`
	// Return only clusters built from these source image IDs, e.g.
	// `image_ids=<id1>&image_ids=<id2>`.
	ImageIDs []string `query:"image_ids,omitzero" format:"uuid4" json:"-"`
	// Filter by who manages the cluster: `user` (default) or `k8s` (Managed
	// Kubernetes). Pass both to include all, e.g. `managed_by=user&managed_by=k8s`.
	//
	// Any of "k8s", "user".
	ManagedBy []string `query:"managed_by,omitzero" json:"-"`
	// Filter by name (case-insensitive), e.g. `name[contains]=gpu`,
	// `name[prefix]=prod-`.
	Name GPUBaremetalClusterListParamsName `query:"name,omitzero" json:"-"`
	// Filter by node count, e.g. `servers_count[gte]=2`,
	// `servers_count[gte]=2&servers_count[lt]=8`.
	ServersCount GPUBaremetalClusterListParamsServersCount `query:"servers_count,omitzero" json:"-"`
	// Filter by tag key regardless of value, e.g. `tag_key[contains]=team`.
	TagKey GPUBaremetalClusterListParamsTagKey `query:"tag_key,omitzero" json:"-"`
	// Filter by tag value regardless of key, e.g. `tag_value[prefix]=prod`.
	TagValue GPUBaremetalClusterListParamsTagValue `query:"tag_value,omitzero" json:"-"`
	// Filter by exact tag key-value pairs, e.g. `tags[env]=prod&tags[team]=core`.
	// Pairs are ANDed; values match case-insensitively.
	Tags map[string]string `query:"tags,omitzero" json:"-"`
	// Filter by last-change time (UTC), e.g. `updated_at[gte]=2026-06-01T00:00:00Z`.
	UpdatedAt GPUBaremetalClusterListParamsUpdatedAt `query:"updated_at,omitzero" json:"-"`
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

// Filter by creation time (UTC), e.g. `created_at[gte]=2026-01-01T00:00:00Z`.
type GPUBaremetalClusterListParamsCreatedAt struct {
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

// URLQuery serializes [GPUBaremetalClusterListParamsCreatedAt]'s query parameters
// as `url.Values`.
func (r GPUBaremetalClusterListParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by flavor (case-insensitive), e.g. `flavor[prefix]=bm3-`,
// `flavor[exact]=bm3-ai-1xlarge-h100-80-8`.
type GPUBaremetalClusterListParamsFlavor struct {
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

// URLQuery serializes [GPUBaremetalClusterListParamsFlavor]'s query parameters as
// `url.Values`.
func (r GPUBaremetalClusterListParamsFlavor) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by name (case-insensitive), e.g. `name[contains]=gpu`,
// `name[prefix]=prod-`.
type GPUBaremetalClusterListParamsName struct {
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

// URLQuery serializes [GPUBaremetalClusterListParamsName]'s query parameters as
// `url.Values`.
func (r GPUBaremetalClusterListParamsName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by node count, e.g. `servers_count[gte]=2`,
// `servers_count[gte]=2&servers_count[lt]=8`.
type GPUBaremetalClusterListParamsServersCount struct {
	// Strictly greater than, e.g. `[gt]=1`.
	Gt param.Opt[int64] `query:"gt,omitzero" json:"-"`
	// Greater than or equal, e.g. `[gte]=2`.
	Gte param.Opt[int64] `query:"gte,omitzero" json:"-"`
	// Strictly less than, e.g. `[lt]=8`.
	Lt param.Opt[int64] `query:"lt,omitzero" json:"-"`
	// Less than or equal, e.g. `[lte]=4`.
	Lte param.Opt[int64] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GPUBaremetalClusterListParamsServersCount]'s query
// parameters as `url.Values`.
func (r GPUBaremetalClusterListParamsServersCount) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by tag key regardless of value, e.g. `tag_key[contains]=team`.
type GPUBaremetalClusterListParamsTagKey struct {
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

// URLQuery serializes [GPUBaremetalClusterListParamsTagKey]'s query parameters as
// `url.Values`.
func (r GPUBaremetalClusterListParamsTagKey) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by tag value regardless of key, e.g. `tag_value[prefix]=prod`.
type GPUBaremetalClusterListParamsTagValue struct {
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

// URLQuery serializes [GPUBaremetalClusterListParamsTagValue]'s query parameters
// as `url.Values`.
func (r GPUBaremetalClusterListParamsTagValue) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by last-change time (UTC), e.g. `updated_at[gte]=2026-06-01T00:00:00Z`.
type GPUBaremetalClusterListParamsUpdatedAt struct {
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

// URLQuery serializes [GPUBaremetalClusterListParamsUpdatedAt]'s query parameters
// as `url.Values`.
func (r GPUBaremetalClusterListParamsUpdatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type GPUBaremetalClusterDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
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

type GPUBaremetalClusterGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type GPUBaremetalClusterPowercycleAllServersParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type GPUBaremetalClusterRebootAllServersParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type GPUBaremetalClusterRebuildParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type GPUBaremetalClusterResizeParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Resized (total) number of instances
	InstancesCount int64 `json:"instances_count" api:"required"`
	paramObj
}

func (r GPUBaremetalClusterResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterResizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GPUBaremetalClusterUpdateServersSettingsParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// System image ID
	ImageID param.Opt[string] `json:"image_id,omitzero" format:"uuid4"`
	// Configuration settings for the servers in the cluster
	ServersSettings GPUBaremetalClusterUpdateServersSettingsParamsServersSettings `json:"servers_settings,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterUpdateServersSettingsParams) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterUpdateServersSettingsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterUpdateServersSettingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration settings for the servers in the cluster
type GPUBaremetalClusterUpdateServersSettingsParamsServersSettings struct {
	// Optional custom user data (Base64-encoded)
	UserData param.Opt[string] `json:"user_data,omitzero"`
	// Optional server access credentials
	Credentials GPUBaremetalClusterUpdateServersSettingsParamsServersSettingsCredentials `json:"credentials,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterUpdateServersSettingsParamsServersSettings) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterUpdateServersSettingsParamsServersSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterUpdateServersSettingsParamsServersSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional server access credentials
type GPUBaremetalClusterUpdateServersSettingsParamsServersSettingsCredentials struct {
	// Specifies the name of the SSH keypair, created via the
	// [/v1/`ssh_keys` endpoint](/docs/api-reference/cloud/ssh-keys/add-or-generate-ssh-key).
	SSHKeyName param.Opt[string] `json:"ssh_key_name,omitzero"`
	paramObj
}

func (r GPUBaremetalClusterUpdateServersSettingsParamsServersSettingsCredentials) MarshalJSON() (data []byte, err error) {
	type shadow GPUBaremetalClusterUpdateServersSettingsParamsServersSettingsCredentials
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GPUBaremetalClusterUpdateServersSettingsParamsServersSettingsCredentials) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
