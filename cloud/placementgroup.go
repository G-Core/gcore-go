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

// Placement Groups allow you to specific a policy that determines whether Virtual
// Machines will be hosted on the same physical server or on different ones.
//
// PlacementGroupService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlacementGroupService] method instead.
type PlacementGroupService struct {
	Options []option.RequestOption
}

// NewPlacementGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlacementGroupService(opts ...option.RequestOption) (r PlacementGroupService) {
	r = PlacementGroupService{}
	r.Options = opts
	return
}

// Create an affinity or anti-affinity or soft-anti-affinity placement group
func (r *PlacementGroupService) New(ctx context.Context, params PlacementGroupNewParams, opts ...option.RequestOption) (res *PlacementGroup, err error) {
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
	path := fmt.Sprintf("cloud/v1/servergroups/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List placement groups
func (r *PlacementGroupService) List(ctx context.Context, query PlacementGroupListParams, opts ...option.RequestOption) (res *PlacementGroupListResponse, err error) {
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
	path := fmt.Sprintf("cloud/v1/servergroups/%v/%v", query.ProjectID.Value, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete placement group
func (r *PlacementGroupService) Delete(ctx context.Context, groupID string, body PlacementGroupDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/servergroups/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get placement group
func (r *PlacementGroupService) Get(ctx context.Context, groupID string, query PlacementGroupGetParams, opts ...option.RequestOption) (res *PlacementGroup, err error) {
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
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/servergroups/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PlacementGroup struct {
	// The list of instances in this server group.
	Instances []PlacementGroupInstance `json:"instances,required"`
	// The name of the server group.
	Name string `json:"name,required"`
	// The server group policy. Options are: anti-affinity, affinity, or
	// soft-anti-affinity.
	Policy string `json:"policy,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// The ID of the server group.
	ServergroupID string `json:"servergroup_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instances     respjson.Field
		Name          respjson.Field
		Policy        respjson.Field
		ProjectID     respjson.Field
		Region        respjson.Field
		RegionID      respjson.Field
		ServergroupID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementGroup) RawJSON() string { return r.JSON.raw }
func (r *PlacementGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementGroupInstance struct {
	// The ID of the instance, corresponding to the attribute 'id'.
	InstanceID string `json:"instance_id,required"`
	// The name of the instance, corresponding to the attribute 'name'.
	InstanceName string `json:"instance_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InstanceID   respjson.Field
		InstanceName respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementGroupInstance) RawJSON() string { return r.JSON.raw }
func (r *PlacementGroupInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementGroupListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []PlacementGroup `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementGroupListResponse) RawJSON() string { return r.JSON.raw }
func (r *PlacementGroupListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementGroupNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// The name of the server group.
	Name string `json:"name,required"`
	// The server group policy.
	//
	// Any of "affinity", "anti-affinity", "soft-anti-affinity".
	Policy PlacementGroupNewParamsPolicy `json:"policy,omitzero,required"`
	paramObj
}

func (r PlacementGroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PlacementGroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlacementGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The server group policy.
type PlacementGroupNewParamsPolicy string

const (
	PlacementGroupNewParamsPolicyAffinity         PlacementGroupNewParamsPolicy = "affinity"
	PlacementGroupNewParamsPolicyAntiAffinity     PlacementGroupNewParamsPolicy = "anti-affinity"
	PlacementGroupNewParamsPolicySoftAntiAffinity PlacementGroupNewParamsPolicy = "soft-anti-affinity"
)

type PlacementGroupListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type PlacementGroupDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type PlacementGroupGetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
