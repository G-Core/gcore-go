// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

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
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/servergroups/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List placement groups
func (r *PlacementGroupService) List(ctx context.Context, query PlacementGroupListParams, opts ...option.RequestOption) (res *PlacementGroupList, err error) {
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
	path := fmt.Sprintf("cloud/v1/servergroups/%v/%v", query.ProjectID.Value, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete placement group
func (r *PlacementGroupService) Delete(ctx context.Context, groupID string, body PlacementGroupDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
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
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/servergroups/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// '#/components/schemas/ServerGroupSerializer'
// "$.components.schemas.ServerGroupSerializer"
type PlacementGroup struct {
	// '#/components/schemas/ServerGroupSerializer/properties/instances'
	// "$.components.schemas.ServerGroupSerializer.properties.instances"
	Instances []PlacementGroupInstance `json:"instances,required"`
	// '#/components/schemas/ServerGroupSerializer/properties/name'
	// "$.components.schemas.ServerGroupSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/ServerGroupSerializer/properties/policy'
	// "$.components.schemas.ServerGroupSerializer.properties.policy"
	Policy string `json:"policy,required"`
	// '#/components/schemas/ServerGroupSerializer/properties/project_id'
	// "$.components.schemas.ServerGroupSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/ServerGroupSerializer/properties/region'
	// "$.components.schemas.ServerGroupSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/ServerGroupSerializer/properties/region_id'
	// "$.components.schemas.ServerGroupSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/ServerGroupSerializer/properties/servergroup_id'
	// "$.components.schemas.ServerGroupSerializer.properties.servergroup_id"
	ServergroupID string `json:"servergroup_id,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Instances     resp.Field
		Name          resp.Field
		Policy        resp.Field
		ProjectID     resp.Field
		Region        resp.Field
		RegionID      resp.Field
		ServergroupID resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementGroup) RawJSON() string { return r.JSON.raw }
func (r *PlacementGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ServerGroupSerializer/properties/instances/items'
// "$.components.schemas.ServerGroupSerializer.properties.instances.items"
type PlacementGroupInstance struct {
	// '#/components/schemas/InstanceBriefSerializer/properties/instance_id'
	// "$.components.schemas.InstanceBriefSerializer.properties.instance_id"
	InstanceID string `json:"instance_id,required"`
	// '#/components/schemas/InstanceBriefSerializer/properties/instance_name'
	// "$.components.schemas.InstanceBriefSerializer.properties.instance_name"
	InstanceName string `json:"instance_name,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		InstanceID   resp.Field
		InstanceName resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementGroupInstance) RawJSON() string { return r.JSON.raw }
func (r *PlacementGroupInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ServerGroupSerializerList'
// "$.components.schemas.ServerGroupSerializerList"
type PlacementGroupList struct {
	// '#/components/schemas/ServerGroupSerializerList/properties/count'
	// "$.components.schemas.ServerGroupSerializerList.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/ServerGroupSerializerList/properties/results'
	// "$.components.schemas.ServerGroupSerializerList.properties.results"
	Results []PlacementGroup `json:"results,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementGroupList) RawJSON() string { return r.JSON.raw }
func (r *PlacementGroupList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementGroupNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fservergroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/servergroups/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fservergroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/servergroups/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateServerGroupSerializer/properties/name'
	// "$.components.schemas.CreateServerGroupSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateServerGroupSerializer/properties/policy'
	// "$.components.schemas.CreateServerGroupSerializer.properties.policy"
	//
	// Any of "affinity", "anti-affinity", "soft-anti-affinity".
	Policy PlacementGroupNewParamsPolicy `json:"policy,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PlacementGroupNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r PlacementGroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PlacementGroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateServerGroupSerializer/properties/policy'
// "$.components.schemas.CreateServerGroupSerializer.properties.policy"
type PlacementGroupNewParamsPolicy string

const (
	PlacementGroupNewParamsPolicyAffinity         PlacementGroupNewParamsPolicy = "affinity"
	PlacementGroupNewParamsPolicyAntiAffinity     PlacementGroupNewParamsPolicy = "anti-affinity"
	PlacementGroupNewParamsPolicySoftAntiAffinity PlacementGroupNewParamsPolicy = "soft-anti-affinity"
)

type PlacementGroupListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fservergroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/servergroups/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fservergroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/servergroups/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PlacementGroupListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type PlacementGroupDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fservergroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/servergroups/{project_id}/{region_id}/{group_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fservergroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/servergroups/{project_id}/{region_id}/{group_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PlacementGroupDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type PlacementGroupGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fservergroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/servergroups/{project_id}/{region_id}/{group_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fservergroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/servergroups/{project_id}/{region_id}/{group_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PlacementGroupGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
