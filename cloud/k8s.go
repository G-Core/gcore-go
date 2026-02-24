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

// K8SService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewK8SService] method instead.
type K8SService struct {
	Options  []option.RequestOption
	Flavors  K8SFlavorService
	Clusters K8SClusterService
}

// NewK8SService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewK8SService(opts ...option.RequestOption) (r K8SService) {
	r = K8SService{}
	r.Options = opts
	r.Flavors = NewK8SFlavorService(opts...)
	r.Clusters = NewK8SClusterService(opts...)
	return
}

// List available k8s cluster versions for creation
func (r *K8SService) ListVersions(ctx context.Context, query K8SListVersionsParams, opts ...option.RequestOption) (res *K8SClusterVersionList, err error) {
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
	path := fmt.Sprintf("cloud/v2/k8s/%v/%v/create_versions", query.ProjectID.Value, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type K8SClusterVersion struct {
	// List of supported Kubernetes cluster versions
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterVersion) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8SClusterVersionList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []K8SClusterVersion `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SClusterVersionList) RawJSON() string { return r.JSON.raw }
func (r *K8SClusterVersionList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type K8SListVersionsParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}
