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

// K8ClusterPoolService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewK8ClusterPoolService] method instead.
type K8ClusterPoolService struct {
	Options []option.RequestOption
	Nodes   K8ClusterPoolNodeService
}

// NewK8ClusterPoolService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewK8ClusterPoolService(opts ...option.RequestOption) (r K8ClusterPoolService) {
	r = K8ClusterPoolService{}
	r.Options = opts
	r.Nodes = NewK8ClusterPoolNodeService(opts...)
	return
}

// Create k8s cluster pool
func (r *K8ClusterPoolService) New(ctx context.Context, clusterName string, params K8ClusterPoolNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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

// Update k8s cluster pool
func (r *K8ClusterPoolService) Update(ctx context.Context, poolName string, params K8ClusterPoolUpdateParams, opts ...option.RequestOption) (res *K8sClusterPool, err error) {
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
func (r *K8ClusterPoolService) List(ctx context.Context, clusterName string, query K8ClusterPoolListParams, opts ...option.RequestOption) (res *K8sClusterPoolList, err error) {
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
func (r *K8ClusterPoolService) Delete(ctx context.Context, poolName string, body K8ClusterPoolDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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

// Get k8s cluster pool
func (r *K8ClusterPoolService) Get(ctx context.Context, poolName string, query K8ClusterPoolGetParams, opts ...option.RequestOption) (res *K8sClusterPool, err error) {
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
func (r *K8ClusterPoolService) Resize(ctx context.Context, poolName string, params K8ClusterPoolResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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

type K8sClusterPool struct {
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
func (r K8sClusterPool) RawJSON() string { return r.JSON.raw }
func (r *K8sClusterPool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8sClusterPoolList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []K8sClusterPool `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8sClusterPoolList) RawJSON() string { return r.JSON.raw }
func (r *K8sClusterPoolList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8ClusterPoolNewParams struct {
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
	BootVolumeType K8ClusterPoolNewParamsBootVolumeType `json:"boot_volume_type,omitzero"`
	// Cri-o configuration for pool nodes
	CrioConfig map[string]string `json:"crio_config,omitzero"`
	// Kubelet configuration for pool nodes
	KubeletConfig map[string]string `json:"kubelet_config,omitzero"`
	// Labels applied to the cluster pool
	Labels map[string]string `json:"labels,omitzero"`
	// Server group policy: anti-affinity, soft-anti-affinity or affinity
	//
	// Any of "affinity", "anti-affinity", "soft-anti-affinity".
	ServergroupPolicy K8ClusterPoolNewParamsServergroupPolicy `json:"servergroup_policy,omitzero"`
	// Taints applied to the cluster pool
	Taints map[string]string `json:"taints,omitzero"`
	paramObj
}

func (r K8ClusterPoolNewParams) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterPoolNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterPoolNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boot volume type
type K8ClusterPoolNewParamsBootVolumeType string

const (
	K8ClusterPoolNewParamsBootVolumeTypeCold          K8ClusterPoolNewParamsBootVolumeType = "cold"
	K8ClusterPoolNewParamsBootVolumeTypeSsdHiiops     K8ClusterPoolNewParamsBootVolumeType = "ssd_hiiops"
	K8ClusterPoolNewParamsBootVolumeTypeSsdLocal      K8ClusterPoolNewParamsBootVolumeType = "ssd_local"
	K8ClusterPoolNewParamsBootVolumeTypeSsdLowlatency K8ClusterPoolNewParamsBootVolumeType = "ssd_lowlatency"
	K8ClusterPoolNewParamsBootVolumeTypeStandard      K8ClusterPoolNewParamsBootVolumeType = "standard"
	K8ClusterPoolNewParamsBootVolumeTypeUltra         K8ClusterPoolNewParamsBootVolumeType = "ultra"
)

// Server group policy: anti-affinity, soft-anti-affinity or affinity
type K8ClusterPoolNewParamsServergroupPolicy string

const (
	K8ClusterPoolNewParamsServergroupPolicyAffinity         K8ClusterPoolNewParamsServergroupPolicy = "affinity"
	K8ClusterPoolNewParamsServergroupPolicyAntiAffinity     K8ClusterPoolNewParamsServergroupPolicy = "anti-affinity"
	K8ClusterPoolNewParamsServergroupPolicySoftAntiAffinity K8ClusterPoolNewParamsServergroupPolicy = "soft-anti-affinity"
)

type K8ClusterPoolUpdateParams struct {
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

func (r K8ClusterPoolUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterPoolUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterPoolUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8ClusterPoolListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type K8ClusterPoolDeleteParams struct {
	ProjectID   param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID    param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	ClusterName string           `path:"cluster_name,required" json:"-"`
	paramObj
}

type K8ClusterPoolGetParams struct {
	ProjectID   param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID    param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	ClusterName string           `path:"cluster_name,required" json:"-"`
	paramObj
}

type K8ClusterPoolResizeParams struct {
	ProjectID   param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID    param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	ClusterName string           `path:"cluster_name,required" json:"-"`
	// Target node count
	NodeCount int64 `json:"node_count,required"`
	paramObj
}

func (r K8ClusterPoolResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow K8ClusterPoolResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *K8ClusterPoolResizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
