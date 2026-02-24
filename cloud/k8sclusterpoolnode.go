// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// K8SClusterPoolNodeService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewK8SClusterPoolNodeService] method instead.
type K8SClusterPoolNodeService struct {
	Options []option.RequestOption
}

// NewK8SClusterPoolNodeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewK8SClusterPoolNodeService(opts ...option.RequestOption) (r K8SClusterPoolNodeService) {
	r = K8SClusterPoolNodeService{}
	r.Options = opts
	return
}

// List k8s cluster pool nodes
func (r *K8SClusterPoolNodeService) List(ctx context.Context, poolName string, params K8SClusterPoolNodeListParams, opts ...option.RequestOption) (res *InstanceList, err error) {
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
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools/%s/instances", params.ProjectID.Value, params.RegionID.Value, params.ClusterName, poolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// After deletion, the node will be automatically recreated to maintain the desired
// pool size.
func (r *K8SClusterPoolNodeService) Delete(ctx context.Context, instanceID string, body K8SClusterPoolNodeDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
	if body.PoolName == "" {
		err = errors.New("missing required pool_name parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/pools/%s/instances/%s", body.ProjectID.Value, body.RegionID.Value, body.ClusterName, body.PoolName, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type K8SClusterPoolNodeListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Cluster name
	ClusterName string `path:"cluster_name" api:"required" json:"-"`
	// Include DDoS profile information if set to true. Default is false.
	WithDDOS param.Opt[bool] `query:"with_ddos,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [K8SClusterPoolNodeListParams]'s query parameters as
// `url.Values`.
func (r K8SClusterPoolNodeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type K8SClusterPoolNodeDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Cluster name
	ClusterName string `path:"cluster_name" api:"required" json:"-"`
	// Pool name
	PoolName string `path:"pool_name" api:"required" json:"-"`
	paramObj
}
