// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
)

// GPUBaremetalClusterInterfaceService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUBaremetalClusterInterfaceService] method instead.
type GPUBaremetalClusterInterfaceService struct {
	Options []option.RequestOption
}

// NewGPUBaremetalClusterInterfaceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGPUBaremetalClusterInterfaceService(opts ...option.RequestOption) (r GPUBaremetalClusterInterfaceService) {
	r = GPUBaremetalClusterInterfaceService{}
	r.Options = opts
	return
}

// List network interfaces attached to GPU cluster servers
func (r *GPUBaremetalClusterInterfaceService) List(ctx context.Context, clusterID string, query GPUBaremetalClusterInterfaceListParams, opts ...option.RequestOption) (res *NetworkInterfaceList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/interfaces", query.ProjectID.Value, query.RegionID.Value, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type GPUBaremetalClusterInterfaceListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D%2Finterfaces/get/parameters/0/schema'
	// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{cluster_id}/interfaces'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fai%2Fclusters%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bcluster_id%7D%2Finterfaces/get/parameters/1/schema'
	// "$.paths['/cloud/v1/ai/clusters/{project_id}/{region_id}/{cluster_id}/interfaces'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f GPUBaremetalClusterInterfaceListParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
