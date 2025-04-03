// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2K8ClusterPoolService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2K8ClusterPoolService] method instead.
type CloudV2K8ClusterPoolService struct {
	Options []option.RequestOption
}

// NewCloudV2K8ClusterPoolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV2K8ClusterPoolService(opts ...option.RequestOption) (r *CloudV2K8ClusterPoolService) {
	r = &CloudV2K8ClusterPoolService{}
	r.Options = opts
	return
}

// Create k8s cluster pool
func (r *CloudV2K8ClusterPoolService) New(ctx context.Context, projectID int64, regionID int64, clusterName string, body CloudV2K8ClusterPoolNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools", projectID, regionID, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get k8s cluster pool details
func (r *CloudV2K8ClusterPoolService) Get(ctx context.Context, projectID int64, regionID int64, clusterName string, poolName string, opts ...option.RequestOption) (res *Pool, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	if poolName == "" {
		err = errors.New("missing required pool_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools/%s", projectID, regionID, clusterName, poolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update k8s cluster pool
func (r *CloudV2K8ClusterPoolService) Update(ctx context.Context, projectID int64, regionID int64, clusterName string, poolName string, body CloudV2K8ClusterPoolUpdateParams, opts ...option.RequestOption) (res *Pool, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	if poolName == "" {
		err = errors.New("missing required pool_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools/%s", projectID, regionID, clusterName, poolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List k8s cluster pools
func (r *CloudV2K8ClusterPoolService) List(ctx context.Context, projectID int64, regionID int64, clusterName string, opts ...option.RequestOption) (res *CloudV2K8ClusterPoolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools", projectID, regionID, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete k8s cluster pool
func (r *CloudV2K8ClusterPoolService) Delete(ctx context.Context, projectID int64, regionID int64, clusterName string, poolName string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	if poolName == "" {
		err = errors.New("missing required pool_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools/%s", projectID, regionID, clusterName, poolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List instances in k8s cluster pool
func (r *CloudV2K8ClusterPoolService) ListInstances(ctx context.Context, projectID int64, regionID int64, clusterName string, poolName string, query CloudV2K8ClusterPoolListInstancesParams, opts ...option.RequestOption) (res *DeprecatedInstanceList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	if poolName == "" {
		err = errors.New("missing required pool_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools/%s/instances", projectID, regionID, clusterName, poolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Resize k8s cluster pool
func (r *CloudV2K8ClusterPoolService) Resize(ctx context.Context, projectID int64, regionID int64, clusterName string, poolName string, body CloudV2K8ClusterPoolResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	if poolName == "" {
		err = errors.New("missing required pool_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools/%s/resize", projectID, regionID, clusterName, poolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Pool struct {
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
	ServergroupPolicy string   `json:"servergroup_policy"`
	JSON              poolJSON `json:"-"`
}

// poolJSON contains the JSON metadata for the struct [Pool]
type poolJSON struct {
	ID                 apijson.Field
	AutoHealingEnabled apijson.Field
	BootVolumeSize     apijson.Field
	BootVolumeType     apijson.Field
	CreatedAt          apijson.Field
	CrioConfig         apijson.Field
	FlavorID           apijson.Field
	IsPublicIpv4       apijson.Field
	KubeletConfig      apijson.Field
	Labels             apijson.Field
	MaxNodeCount       apijson.Field
	MinNodeCount       apijson.Field
	Name               apijson.Field
	NodeCount          apijson.Field
	Status             apijson.Field
	Taints             apijson.Field
	ServergroupID      apijson.Field
	ServergroupName    apijson.Field
	ServergroupPolicy  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Pool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolJSON) RawJSON() string {
	return r.raw
}

type PoolCreateParam struct {
	// Flavor ID
	FlavorID param.Field[string] `json:"flavor_id,required"`
	// Minimum node count
	MinNodeCount param.Field[int64] `json:"min_node_count,required"`
	// Pool's name
	Name param.Field[string] `json:"name,required"`
	// Enable auto healing
	AutoHealingEnabled param.Field[bool] `json:"auto_healing_enabled"`
	// Boot volume size
	BootVolumeSize param.Field[int64] `json:"boot_volume_size"`
	// Boot volume type
	BootVolumeType param.Field[VolumeTypeName] `json:"boot_volume_type"`
	// Cri-o configuration for pool nodes
	CrioConfig param.Field[map[string]string] `json:"crio_config"`
	// Enable public v4 address
	IsPublicIpv4 param.Field[bool] `json:"is_public_ipv4"`
	// Kubelet configuration for pool nodes
	KubeletConfig param.Field[map[string]string] `json:"kubelet_config"`
	// Labels applied to the cluster pool
	Labels param.Field[map[string]string] `json:"labels"`
	// Maximum node count
	MaxNodeCount param.Field[int64] `json:"max_node_count"`
	// Server group policy: anti-affinity, soft-anti-affinity or affinity
	ServergroupPolicy param.Field[ServerGroupPolicy] `json:"servergroup_policy"`
	// Taints applied to the cluster pool
	Taints param.Field[map[string]string] `json:"taints"`
}

func (r PoolCreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2K8ClusterPoolListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Pool                               `json:"results,required"`
	JSON    cloudV2K8ClusterPoolListResponseJSON `json:"-"`
}

// cloudV2K8ClusterPoolListResponseJSON contains the JSON metadata for the struct
// [CloudV2K8ClusterPoolListResponse]
type cloudV2K8ClusterPoolListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV2K8ClusterPoolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV2K8ClusterPoolListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV2K8ClusterPoolNewParams struct {
	PoolCreate PoolCreateParam `json:"pool_create,required"`
}

func (r CloudV2K8ClusterPoolNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PoolCreate)
}

type CloudV2K8ClusterPoolUpdateParams struct {
	// Enable/disable auto healing
	AutoHealingEnabled param.Field[bool] `json:"auto_healing_enabled"`
	// Labels applied to the cluster pool
	Labels param.Field[map[string]string] `json:"labels"`
	// Maximum node count
	MaxNodeCount param.Field[int64] `json:"max_node_count"`
	// Minimum node count
	MinNodeCount param.Field[int64] `json:"min_node_count"`
	// Current node count
	NodeCount param.Field[int64] `json:"node_count"`
	// Taints applied to the cluster pool
	Taints param.Field[map[string]string] `json:"taints"`
}

func (r CloudV2K8ClusterPoolUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2K8ClusterPoolListInstancesParams struct {
	// Include DDoS profile information if set to true. The default value is false.
	WithDdos param.Field[bool] `query:"with_ddos"`
}

// URLQuery serializes [CloudV2K8ClusterPoolListInstancesParams]'s query parameters
// as `url.Values`.
func (r CloudV2K8ClusterPoolListInstancesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV2K8ClusterPoolResizeParams struct {
	// Target node count
	NodeCount param.Field[int64] `json:"node_count,required"`
}

func (r CloudV2K8ClusterPoolResizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
