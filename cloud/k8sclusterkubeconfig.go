// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// Kubeconfig provides the necessary configuration and credentials to access a
// Kubernetes cluster using kubectl or other Kubernetes clients.
//
// K8SClusterKubeconfigService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewK8SClusterKubeconfigService] method instead.
type K8SClusterKubeconfigService struct {
	Options []option.RequestOption
}

// NewK8SClusterKubeconfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewK8SClusterKubeconfigService(opts ...option.RequestOption) (r K8SClusterKubeconfigService) {
	r = K8SClusterKubeconfigService{}
	r.Options = opts
	return
}

// Get k8s cluster kubeconfig
func (r *K8SClusterKubeconfigService) Get(ctx context.Context, clusterName string, query K8SClusterKubeconfigGetParams, opts ...option.RequestOption) (res *K8SClusterKubeconfig, err error) {
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
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v2/k8s/clusters/%v/%v/%s/config", query.ProjectID.Value, query.RegionID.Value, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type K8SClusterKubeconfigGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}
