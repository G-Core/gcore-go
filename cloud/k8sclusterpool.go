// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// K8SClusterPoolService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewK8SClusterPoolService] method instead.
type K8SClusterPoolService struct {
	Options []option.RequestOption
	Nodes   K8SClusterPoolNodeService
	tasks   TaskService
}

// NewK8SClusterPoolService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewK8SClusterPoolService(opts ...option.RequestOption) (r K8SClusterPoolService) {
	r = K8SClusterPoolService{}
	r.Options = opts
	r.Nodes = NewK8SClusterPoolNodeService(opts...)
	r.tasks = NewTaskService(opts...)
	return
}

// Create k8s cluster pool
func (r *K8SClusterPoolService) New(ctx context.Context, clusterName string, params K8SClusterPoolNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools", params.ProjectID.Value, params.RegionID.Value, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// NewAndPoll creates a new k8s cluster pool and polls for completion
func (r *K8SClusterPoolService) NewAndPoll(ctx context.Context, clusterName string, params K8SClusterPoolNewParams, opts ...option.RequestOption) (v *K8SClusterPool, err error) {
	resource, err := r.New(ctx, clusterName, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams K8SClusterPoolGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID
	getParams.ClusterName = clusterName

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	// for k8s cluster pool creation the task.CreatedResources.K8SPools only contains the cluster pool ID and not the
	// cluster pool name, which is the path parameter required to retrieve the created cluster pool. Therefore, we use
	// the params.Name instead.
	return r.Get(ctx, params.Name, getParams, opts...)
}

// Update k8s cluster pool
func (r *K8SClusterPoolService) Update(ctx context.Context, poolName string, params K8SClusterPoolUpdateParams, opts ...option.RequestOption) (res *K8SClusterPool, err error) {
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
	if params.ClusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	if poolName == "" {
		err = errors.New("missing required pool_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools/%s", params.ProjectID.Value, params.RegionID.Value, params.ClusterName, poolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List k8s cluster pools
func (r *K8SClusterPoolService) List(ctx context.Context, clusterName string, query K8SClusterPoolListParams, opts ...option.RequestOption) (res *K8SClusterPoolListResponse, err error) {
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
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools", query.ProjectID.Value, query.RegionID.Value, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete k8s cluster pool
func (r *K8SClusterPoolService) Delete(ctx context.Context, poolName string, body K8SClusterPoolDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if body.ClusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	if poolName == "" {
		err = errors.New("missing required pool_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools/%s", body.ProjectID.Value, body.RegionID.Value, body.ClusterName, poolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// DeleteAndPoll deletes a k8s cluster pool and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *K8SClusterPoolService) DeleteAndPoll(ctx context.Context, poolName string, body K8SClusterPoolDeleteParams, opts ...option.RequestOption) error {
	resource, err := r.Delete(ctx, poolName, body, opts...)
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

// Calculate quota requirements for a new cluster pool before creation. Returns
// exceeded quotas if regional limits would be violated. Use before pool creation
// to validate resource availability. Checks: CPU, RAM, volumes, VMs, GPUs, and
// baremetal quotas depending on flavor type.
func (r *K8SClusterPoolService) CheckQuota(ctx context.Context, params K8SClusterPoolCheckQuotaParams, opts ...option.RequestOption) (res *K8SClusterPoolQuota, err error) {
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
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/pools/check_limits", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get k8s cluster pool
func (r *K8SClusterPoolService) Get(ctx context.Context, poolName string, query K8SClusterPoolGetParams, opts ...option.RequestOption) (res *K8SClusterPool, err error) {
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
	if query.ClusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	if poolName == "" {
		err = errors.New("missing required pool_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools/%s", query.ProjectID.Value, query.RegionID.Value, query.ClusterName, poolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Resize k8s cluster pool
func (r *K8SClusterPoolService) Resize(ctx context.Context, poolName string, params K8SClusterPoolResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if params.ClusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	if poolName == "" {
		err = errors.New("missing required pool_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools/%s/resize", params.ProjectID.Value, params.RegionID.Value, params.ClusterName, poolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ResizeAndPoll resizes a k8s cluster pool and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *K8SClusterPoolService) ResizeAndPoll(ctx context.Context, poolName string, params K8SClusterPoolResizeParams, opts ...option.RequestOption) (v *K8SClusterPool, err error) {
	resource, err := r.Resize(ctx, poolName, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams K8SClusterPoolGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID
	getParams.ClusterName = params.ClusterName

	if len(resource.Tasks) == 0 {
		return nil, errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, poolName, getParams, opts...)
}

type K8SClusterPool struct {
	// UUID of the cluster pool
	ID string `json:"id,required"`
	// Indicates the status of auto healing
	AutoHealingEnabled bool `json:"auto_healing_enabled,required"`
	// Size of the boot volume
	BootVolumeSize int64 `json:"boot_volume_size,required"`
	// Type of the boot volume
	BootVolumeType string `json:"boot_volume_type,required"`
	// Date of function creation
	CreatedAt string `json:"created_at,required"`
	// Crio configuration for pool nodes
	CrioConfig map[string]string `json:"crio_config,required"`
	// ID of the cluster pool flavor
	FlavorID string `json:"flavor_id,required"`
	// Indicates if the pool is public
	IsPublicIpv4 bool `json:"is_public_ipv4,required"`
	// Kubelet configuration for pool nodes
	KubeletConfig map[string]string `json:"kubelet_config,required"`
	// Labels applied to the cluster pool
	Labels map[string]string `json:"labels,required"`
	// Maximum node count in the cluster pool
	MaxNodeCount int64 `json:"max_node_count,required"`
	// Minimum node count in the cluster pool
	MinNodeCount int64 `json:"min_node_count,required"`
	// Name of the cluster pool
	Name string `json:"name,required"`
	// Node count in the cluster pool
	NodeCount int64 `json:"node_count,required"`
	// Status of the cluster pool
	Status string `json:"status,required"`
	// Taints applied to the cluster pool
	Taints map[string]string `json:"taints,required"`
	// Server group ID
	ServergroupID string `json:"servergroup_id"`
	// Server group name
	ServergroupName string `json:"servergroup_name"`
	// Anti-affinity, affinity or soft-anti-affinity server group policy
	ServergroupPolicy string `json:"servergroup_policy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AutoHealingEnabled respjson.Field
		BootVolumeSize     respjson.Field
		BootVolumeType     respjson.Field
		CreatedAt          respjson.Field
		CrioConfig         respjson.Field
		FlavorID           respjson.Field
		IsPublicIpv4       respjson.Field
		KubeletConfig      respjson.Field
		Labels             respjson.Field
		MaxNodeCount       respjson.Field
		MinNodeCount       respjson.Field
		Name               respjson.Field
		NodeCount          respjson.Field
		Status             respjson.Field
		Taints             respjson.Field
		ServergroupID      respjson.Field
		ServergroupName    respjson.Field
		ServergroupPolicy  respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterPool) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterPool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response schema for K8s cluster quota check.
//
// Returns quota fields that are exceeded. Fields are only included when regional
// limits would be violated. Empty response means no quotas exceeded.
type K8SClusterPoolQuota struct {
	// Bare metal A100 GPU server count limit
	BaremetalGPUA100CountLimit int64 `json:"baremetal_gpu_a100_count_limit"`
	// Bare metal A100 GPU server count requested
	BaremetalGPUA100CountRequested int64 `json:"baremetal_gpu_a100_count_requested"`
	// Bare metal A100 GPU server count usage
	BaremetalGPUA100CountUsage int64 `json:"baremetal_gpu_a100_count_usage"`
	// Bare metal H100 GPU server count limit
	BaremetalGPUH100CountLimit int64 `json:"baremetal_gpu_h100_count_limit"`
	// Bare metal H100 GPU server count requested
	BaremetalGPUH100CountRequested int64 `json:"baremetal_gpu_h100_count_requested"`
	// Bare metal H100 GPU server count usage
	BaremetalGPUH100CountUsage int64 `json:"baremetal_gpu_h100_count_usage"`
	// Bare metal H200 GPU server count limit
	BaremetalGPUH200CountLimit int64 `json:"baremetal_gpu_h200_count_limit"`
	// Bare metal H200 GPU server count requested
	BaremetalGPUH200CountRequested int64 `json:"baremetal_gpu_h200_count_requested"`
	// Bare metal H200 GPU server count usage
	BaremetalGPUH200CountUsage int64 `json:"baremetal_gpu_h200_count_usage"`
	// Bare metal L40S GPU server count limit
	BaremetalGPUL40sCountLimit int64 `json:"baremetal_gpu_l40s_count_limit"`
	// Bare metal L40S GPU server count requested
	BaremetalGPUL40sCountRequested int64 `json:"baremetal_gpu_l40s_count_requested"`
	// Bare metal L40S GPU server count usage
	BaremetalGPUL40sCountUsage int64 `json:"baremetal_gpu_l40s_count_usage"`
	// High-frequency bare metal servers count limit
	BaremetalHfCountLimit int64 `json:"baremetal_hf_count_limit"`
	// High-frequency bare metal servers count requested
	BaremetalHfCountRequested int64 `json:"baremetal_hf_count_requested"`
	// High-frequency bare metal servers count usage
	BaremetalHfCountUsage int64 `json:"baremetal_hf_count_usage"`
	// K8s clusters count limit
	ClusterCountLimit int64 `json:"cluster_count_limit"`
	// K8s clusters count requested
	ClusterCountRequested int64 `json:"cluster_count_requested"`
	// K8s clusters count usage
	ClusterCountUsage int64 `json:"cluster_count_usage"`
	// vCPU Count limit
	CPUCountLimit int64 `json:"cpu_count_limit"`
	// vCPU Count requested
	CPUCountRequested int64 `json:"cpu_count_requested"`
	// vCPU Count usage
	CPUCountUsage int64 `json:"cpu_count_usage"`
	// Firewalls Count limit
	FirewallCountLimit int64 `json:"firewall_count_limit"`
	// Firewalls Count requested
	FirewallCountRequested int64 `json:"firewall_count_requested"`
	// Firewalls Count usage
	FirewallCountUsage int64 `json:"firewall_count_usage"`
	// Floating IP Count limit
	FloatingCountLimit int64 `json:"floating_count_limit"`
	// Floating IP Count requested
	FloatingCountRequested int64 `json:"floating_count_requested"`
	// Floating IP Count usage
	FloatingCountUsage int64 `json:"floating_count_usage"`
	// GPU Count limit
	GPUCountLimit int64 `json:"gpu_count_limit"`
	// GPU Count requested
	GPUCountRequested int64 `json:"gpu_count_requested"`
	// GPU Count usage
	GPUCountUsage int64 `json:"gpu_count_usage"`
	// Virtual A100 GPU card count limit
	GPUVirtualA100CountLimit int64 `json:"gpu_virtual_a100_count_limit"`
	// Virtual A100 GPU card count requested
	GPUVirtualA100CountRequested int64 `json:"gpu_virtual_a100_count_requested"`
	// Virtual A100 GPU card count usage
	GPUVirtualA100CountUsage int64 `json:"gpu_virtual_a100_count_usage"`
	// Virtual H100 GPU card count limit
	GPUVirtualH100CountLimit int64 `json:"gpu_virtual_h100_count_limit"`
	// Virtual H100 GPU card count requested
	GPUVirtualH100CountRequested int64 `json:"gpu_virtual_h100_count_requested"`
	// Virtual H100 GPU card count usage
	GPUVirtualH100CountUsage int64 `json:"gpu_virtual_h100_count_usage"`
	// Virtual H200 GPU card count limit
	GPUVirtualH200CountLimit int64 `json:"gpu_virtual_h200_count_limit"`
	// Virtual H200 GPU card count requested
	GPUVirtualH200CountRequested int64 `json:"gpu_virtual_h200_count_requested"`
	// Virtual H200 GPU card count usage
	GPUVirtualH200CountUsage int64 `json:"gpu_virtual_h200_count_usage"`
	// Virtual L40S GPU card count limit
	GPUVirtualL40sCountLimit int64 `json:"gpu_virtual_l40s_count_limit"`
	// Virtual L40S GPU card count requested
	GPUVirtualL40sCountRequested int64 `json:"gpu_virtual_l40s_count_requested"`
	// Virtual L40S GPU card count usage
	GPUVirtualL40sCountUsage int64 `json:"gpu_virtual_l40s_count_usage"`
	// LaaS Topics Count limit
	LaasTopicCountLimit int64 `json:"laas_topic_count_limit"`
	// LaaS Topics Count requested
	LaasTopicCountRequested int64 `json:"laas_topic_count_requested"`
	// LaaS Topics Count usage
	LaasTopicCountUsage int64 `json:"laas_topic_count_usage"`
	// Load Balancers Count limit
	LoadbalancerCountLimit int64 `json:"loadbalancer_count_limit"`
	// Load Balancers Count requested
	LoadbalancerCountRequested int64 `json:"loadbalancer_count_requested"`
	// Load Balancers Count usage
	LoadbalancerCountUsage int64 `json:"loadbalancer_count_usage"`
	// RAM Size, MiB limit
	RamLimit int64 `json:"ram_limit"`
	// RAM Size, MiB requested
	RamRequested int64 `json:"ram_requested"`
	// RAM Size, MiB usage
	RamUsage int64 `json:"ram_usage"`
	// Placement Group Count limit
	ServergroupCountLimit int64 `json:"servergroup_count_limit"`
	// Placement Group Count requested
	ServergroupCountRequested int64 `json:"servergroup_count_requested"`
	// Placement Group Count usage
	ServergroupCountUsage int64 `json:"servergroup_count_usage"`
	// VMs Count limit
	VmCountLimit int64 `json:"vm_count_limit"`
	// VMs Count requested
	VmCountRequested int64 `json:"vm_count_requested"`
	// VMs Count usage
	VmCountUsage int64 `json:"vm_count_usage"`
	// Volumes Count limit
	VolumeCountLimit int64 `json:"volume_count_limit"`
	// Volumes Count requested
	VolumeCountRequested int64 `json:"volume_count_requested"`
	// Volumes Count usage
	VolumeCountUsage int64 `json:"volume_count_usage"`
	// Volumes Size, GiB limit
	VolumeSizeLimit int64 `json:"volume_size_limit"`
	// Volumes Size, GiB requested
	VolumeSizeRequested int64 `json:"volume_size_requested"`
	// Volumes Size, GiB usage
	VolumeSizeUsage int64 `json:"volume_size_usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaremetalGPUA100CountLimit     respjson.Field
		BaremetalGPUA100CountRequested respjson.Field
		BaremetalGPUA100CountUsage     respjson.Field
		BaremetalGPUH100CountLimit     respjson.Field
		BaremetalGPUH100CountRequested respjson.Field
		BaremetalGPUH100CountUsage     respjson.Field
		BaremetalGPUH200CountLimit     respjson.Field
		BaremetalGPUH200CountRequested respjson.Field
		BaremetalGPUH200CountUsage     respjson.Field
		BaremetalGPUL40sCountLimit     respjson.Field
		BaremetalGPUL40sCountRequested respjson.Field
		BaremetalGPUL40sCountUsage     respjson.Field
		BaremetalHfCountLimit          respjson.Field
		BaremetalHfCountRequested      respjson.Field
		BaremetalHfCountUsage          respjson.Field
		ClusterCountLimit              respjson.Field
		ClusterCountRequested          respjson.Field
		ClusterCountUsage              respjson.Field
		CPUCountLimit                  respjson.Field
		CPUCountRequested              respjson.Field
		CPUCountUsage                  respjson.Field
		FirewallCountLimit             respjson.Field
		FirewallCountRequested         respjson.Field
		FirewallCountUsage             respjson.Field
		FloatingCountLimit             respjson.Field
		FloatingCountRequested         respjson.Field
		FloatingCountUsage             respjson.Field
		GPUCountLimit                  respjson.Field
		GPUCountRequested              respjson.Field
		GPUCountUsage                  respjson.Field
		GPUVirtualA100CountLimit       respjson.Field
		GPUVirtualA100CountRequested   respjson.Field
		GPUVirtualA100CountUsage       respjson.Field
		GPUVirtualH100CountLimit       respjson.Field
		GPUVirtualH100CountRequested   respjson.Field
		GPUVirtualH100CountUsage       respjson.Field
		GPUVirtualH200CountLimit       respjson.Field
		GPUVirtualH200CountRequested   respjson.Field
		GPUVirtualH200CountUsage       respjson.Field
		GPUVirtualL40sCountLimit       respjson.Field
		GPUVirtualL40sCountRequested   respjson.Field
		GPUVirtualL40sCountUsage       respjson.Field
		LaasTopicCountLimit            respjson.Field
		LaasTopicCountRequested        respjson.Field
		LaasTopicCountUsage            respjson.Field
		LoadbalancerCountLimit         respjson.Field
		LoadbalancerCountRequested     respjson.Field
		LoadbalancerCountUsage         respjson.Field
		RamLimit                       respjson.Field
		RamRequested                   respjson.Field
		RamUsage                       respjson.Field
		ServergroupCountLimit          respjson.Field
		ServergroupCountRequested      respjson.Field
		ServergroupCountUsage          respjson.Field
		VmCountLimit                   respjson.Field
		VmCountRequested               respjson.Field
		VmCountUsage                   respjson.Field
		VolumeCountLimit               respjson.Field
		VolumeCountRequested           respjson.Field
		VolumeCountUsage               respjson.Field
		VolumeSizeLimit                respjson.Field
		VolumeSizeRequested            respjson.Field
		VolumeSizeUsage                respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterPoolQuota) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterPoolQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8SClusterPoolListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []K8SClusterPool `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterPoolListResponse) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterPoolListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8SClusterPoolNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Flavor ID
	FlavorID string `json:"flavor_id,required"`
	// Minimum node count
	MinNodeCount int64 `json:"min_node_count,required"`
	// Pool's name
	Name string `json:"name,required"`
	// Enable auto healing
	AutoHealingEnabled param.Opt[bool] `json:"auto_healing_enabled,omitzero"`
	// Boot volume size
	BootVolumeSize param.Opt[int64] `json:"boot_volume_size,omitzero"`
	// Enable public v4 address
	IsPublicIpv4 param.Opt[bool] `json:"is_public_ipv4,omitzero"`
	// Maximum node count
	MaxNodeCount param.Opt[int64] `json:"max_node_count,omitzero"`
	// Boot volume type
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	BootVolumeType K8SClusterPoolNewParamsBootVolumeType `json:"boot_volume_type,omitzero"`
	// Cri-o configuration for pool nodes
	CrioConfig map[string]string `json:"crio_config,omitzero"`
	// Kubelet configuration for pool nodes
	KubeletConfig map[string]string `json:"kubelet_config,omitzero"`
	// Labels applied to the cluster pool
	Labels map[string]string `json:"labels,omitzero"`
	// Server group policy: anti-affinity, soft-anti-affinity or affinity
	//
	// Any of "affinity", "anti-affinity", "soft-anti-affinity".
	ServergroupPolicy K8SClusterPoolNewParamsServergroupPolicy `json:"servergroup_policy,omitzero"`
	// Taints applied to the cluster pool
	Taints map[string]string `json:"taints,omitzero"`
	paramObj
}

func (r K8SClusterPoolNewParams) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterPoolNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterPoolNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boot volume type
type K8SClusterPoolNewParamsBootVolumeType string

const (
	K8SClusterPoolNewParamsBootVolumeTypeCold          K8SClusterPoolNewParamsBootVolumeType = "cold"
	K8SClusterPoolNewParamsBootVolumeTypeSsdHiiops     K8SClusterPoolNewParamsBootVolumeType = "ssd_hiiops"
	K8SClusterPoolNewParamsBootVolumeTypeSsdLocal      K8SClusterPoolNewParamsBootVolumeType = "ssd_local"
	K8SClusterPoolNewParamsBootVolumeTypeSsdLowlatency K8SClusterPoolNewParamsBootVolumeType = "ssd_lowlatency"
	K8SClusterPoolNewParamsBootVolumeTypeStandard      K8SClusterPoolNewParamsBootVolumeType = "standard"
	K8SClusterPoolNewParamsBootVolumeTypeUltra         K8SClusterPoolNewParamsBootVolumeType = "ultra"
)

// Server group policy: anti-affinity, soft-anti-affinity or affinity
type K8SClusterPoolNewParamsServergroupPolicy string

const (
	K8SClusterPoolNewParamsServergroupPolicyAffinity         K8SClusterPoolNewParamsServergroupPolicy = "affinity"
	K8SClusterPoolNewParamsServergroupPolicyAntiAffinity     K8SClusterPoolNewParamsServergroupPolicy = "anti-affinity"
	K8SClusterPoolNewParamsServergroupPolicySoftAntiAffinity K8SClusterPoolNewParamsServergroupPolicy = "soft-anti-affinity"
)

type K8SClusterPoolUpdateParams struct {
	ProjectID   param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID    param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	ClusterName string           `path:"cluster_name,required" json:"-"`
	// Enable/disable auto healing
	AutoHealingEnabled param.Opt[bool] `json:"auto_healing_enabled,omitzero"`
	// Maximum node count
	MaxNodeCount param.Opt[int64] `json:"max_node_count,omitzero"`
	// Minimum node count
	MinNodeCount param.Opt[int64] `json:"min_node_count,omitzero"`
	// This field is deprecated. Please use the cluster pool resize handler instead.
	NodeCount param.Opt[int64] `json:"node_count,omitzero"`
	// Labels applied to the cluster pool
	Labels map[string]string `json:"labels,omitzero"`
	// Taints applied to the cluster pool
	Taints map[string]string `json:"taints,omitzero"`
	paramObj
}

func (r K8SClusterPoolUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterPoolUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterPoolUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8SClusterPoolListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type K8SClusterPoolDeleteParams struct {
	ProjectID   param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID    param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	ClusterName string           `path:"cluster_name,required" json:"-"`
	paramObj
}

type K8SClusterPoolCheckQuotaParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Flavor ID
	FlavorID string `json:"flavor_id,required"`
	// Boot volume size
	BootVolumeSize param.Opt[int64] `json:"boot_volume_size,omitzero"`
	// Maximum node count
	MaxNodeCount param.Opt[int64] `json:"max_node_count,omitzero"`
	// Minimum node count
	MinNodeCount param.Opt[int64] `json:"min_node_count,omitzero"`
	// Name of the cluster pool
	Name param.Opt[string] `json:"name,omitzero"`
	// Maximum node count
	NodeCount param.Opt[int64] `json:"node_count,omitzero"`
	// Server group policy: anti-affinity, soft-anti-affinity or affinity
	//
	// Any of "affinity", "anti-affinity", "soft-anti-affinity".
	ServergroupPolicy K8SClusterPoolCheckQuotaParamsServergroupPolicy `json:"servergroup_policy,omitzero"`
	paramObj
}

func (r K8SClusterPoolCheckQuotaParams) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterPoolCheckQuotaParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterPoolCheckQuotaParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Server group policy: anti-affinity, soft-anti-affinity or affinity
type K8SClusterPoolCheckQuotaParamsServergroupPolicy string

const (
	K8SClusterPoolCheckQuotaParamsServergroupPolicyAffinity         K8SClusterPoolCheckQuotaParamsServergroupPolicy = "affinity"
	K8SClusterPoolCheckQuotaParamsServergroupPolicyAntiAffinity     K8SClusterPoolCheckQuotaParamsServergroupPolicy = "anti-affinity"
	K8SClusterPoolCheckQuotaParamsServergroupPolicySoftAntiAffinity K8SClusterPoolCheckQuotaParamsServergroupPolicy = "soft-anti-affinity"
)

type K8SClusterPoolGetParams struct {
	ProjectID   param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID    param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	ClusterName string           `path:"cluster_name,required" json:"-"`
	paramObj
}

type K8SClusterPoolResizeParams struct {
	ProjectID   param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID    param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	ClusterName string           `path:"cluster_name,required" json:"-"`
	// Target node count
	NodeCount int64 `json:"node_count,required"`
	paramObj
}

func (r K8SClusterPoolResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow K8SClusterPoolResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8SClusterPoolResizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
