// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
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

// Placement groups set an affinity or anti-affinity policy for virtual machines,
// in any region of the project.
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

// Create an affinity, anti-affinity, soft-affinity or soft-anti-affinity placement
// group. Returns the ID of the task performing the creation.
//
// The name and the policy cannot change after the create. To add an instance to
// the group, send `placement_group_id` when you create or update the instance.
//
// With `?dry_run=true` the request runs every check and changes nothing. The
// response holds an empty `tasks` array. A failed check returns the same error as
// the real create.
func (r *PlacementGroupService) New(ctx context.Context, params PlacementGroupNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/placement_groups", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List a project's placement groups across all regions. Groups created for a
// managed service such as Kubernetes are returned with a `managed_by` object and
// are read-only.
func (r *PlacementGroupService) List(ctx context.Context, params PlacementGroupListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[PlacementGroup], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/placement_groups", params.ProjectID.Value)
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

// List a project's placement groups across all regions. Groups created for a
// managed service such as Kubernetes are returned with a `managed_by` object and
// are read-only.
func (r *PlacementGroupService) ListAutoPaging(ctx context.Context, params PlacementGroupListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[PlacementGroup] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Returns the ID of the task performing the deletion.
//
// A group created for a managed service such as Kubernetes (non-null `managed_by`)
// is platform-owned and cannot be deleted. Members are not deleted, only
// ungrouped. The delete removes no other resource, and no resource is attached to
// a group. So `resources_to_delete` holds only the group, and
// `resources_to_preserve` is empty.
//
// With `?dry_run=true` the request runs every check and changes nothing. The
// response holds an empty `tasks` array. A failed check returns the same error as
// the real delete.
func (r *PlacementGroupService) Delete(ctx context.Context, placementGroupID string, params PlacementGroupDeleteParams, opts ...option.RequestOption) (res *PlacementGroupDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if placementGroupID == "" {
		err = errors.New("missing required placement_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/placement_groups/%s", params.ProjectID.Value, placementGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Returns a single placement group by ID; the region is resolved from the ID, so
// it is not part of the path.
func (r *PlacementGroupService) Get(ctx context.Context, placementGroupID string, query PlacementGroupGetParams, opts ...option.RequestOption) (res *PlacementGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if placementGroupID == "" {
		err = errors.New("missing required placement_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/placement_groups/%s", query.ProjectID.Value, placementGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PlacementGroup struct {
	// Placement group ID.
	ID string `json:"id" api:"required" format:"uuid4"`
	// Datetime when the placement group was first seen.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Service that operates this placement group. A managed group is read-only and you
	// cannot delete it. Null when you operate the group.
	ManagedBy PlacementGroupManagedBy `json:"managed_by" api:"required"`
	// UUIDs of the instances that are members of this placement group, sorted
	// ascending.
	MemberIDs []string `json:"member_ids" api:"required" format:"uuid4"`
	// Placement group name.
	Name string `json:"name" api:"required"`
	// Placement policy: affinity, anti-affinity, soft-affinity, or soft-anti-affinity.
	//
	// Any of "affinity", "anti-affinity", "soft-affinity", "soft-anti-affinity".
	Policy PlacementGroupPolicy `json:"policy" api:"required"`
	// Region ID.
	RegionID int64 `json:"region_id" api:"required"`
	// Datetime when the placement group last changed.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ManagedBy   respjson.Field
		MemberIDs   respjson.Field
		Name        respjson.Field
		Policy      respjson.Field
		RegionID    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementGroup) RawJSON() string { return r.JSON.raw }
func (r *PlacementGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Service that operates this placement group. A managed group is read-only and you
// cannot delete it. Null when you operate the group.
type PlacementGroupManagedBy struct {
	// Type of the managing service.
	//
	// Any of "system".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementGroupManagedBy) RawJSON() string { return r.JSON.raw }
func (r *PlacementGroupManagedBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Placement policy: affinity, anti-affinity, soft-affinity, or soft-anti-affinity.
type PlacementGroupPolicy string

const (
	PlacementGroupPolicyAffinity         PlacementGroupPolicy = "affinity"
	PlacementGroupPolicyAntiAffinity     PlacementGroupPolicy = "anti-affinity"
	PlacementGroupPolicySoftAffinity     PlacementGroupPolicy = "soft-affinity"
	PlacementGroupPolicySoftAntiAffinity PlacementGroupPolicy = "soft-anti-affinity"
)

type PlacementGroupDeleteResponse struct {
	// The placement group. Nothing else goes with it.
	ResourcesToDelete []PlacementGroupDeleteResponseResourcesToDelete `json:"resources_to_delete" api:"required"`
	// Always empty. No resource is attached to a placement group. The other deletes
	// use the same shape.
	ResourcesToPreserve []PlacementGroupDeleteResponseResourcesToPreserve `json:"resources_to_preserve" api:"required"`
	// List of task IDs representing asynchronous operations. Use these IDs to monitor
	// operation progress:
	//
	//   - `GET /v1/tasks/{task_id}` - Check individual task status and details Poll task
	//     status until completion (`FINISHED`/`ERROR`) before proceeding with dependent
	//     operations.
	Tasks []string `json:"tasks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResourcesToDelete   respjson.Field
		ResourcesToPreserve respjson.Field
		Tasks               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementGroupDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *PlacementGroupDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementGroupDeleteResponseResourcesToDelete struct {
	// The ID of the resource.
	ID string `json:"id" api:"required"`
	// The name of the resource. For a floating IP, the IP address. `null` for a
	// volume, and for a resource without a name.
	Name string `json:"name" api:"required"`
	// The type of the resource.
	Type constant.PlacementGroups `json:"type" default:"placement_groups"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementGroupDeleteResponseResourcesToDelete) RawJSON() string { return r.JSON.raw }
func (r *PlacementGroupDeleteResponseResourcesToDelete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementGroupDeleteResponseResourcesToPreserve struct {
	// The ID of the resource.
	ID string `json:"id" api:"required"`
	// The name of the resource. For a floating IP, the IP address. `null` for a
	// volume, and for a resource without a name.
	Name string `json:"name" api:"required"`
	// The type of the resource.
	Type constant.PlacementGroups `json:"type" default:"placement_groups"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlacementGroupDeleteResponseResourcesToPreserve) RawJSON() string { return r.JSON.raw }
func (r *PlacementGroupDeleteResponseResourcesToPreserve) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementGroupNewParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Placement group name. A name starting with `k8s-` is rejected with 400
	// `validation-error`. The prefix is reserved for the groups that managed
	// Kubernetes creates.
	Name string `json:"name" api:"required"`
	// Placement policy: affinity, anti-affinity, soft-affinity, or soft-anti-affinity.
	//
	// Any of "affinity", "anti-affinity", "soft-affinity", "soft-anti-affinity".
	Policy PlacementGroupNewParamsPolicy `json:"policy,omitzero" api:"required"`
	// Region to create the placement group in.
	RegionID int64 `json:"region_id" api:"required"`
	// Run every check of the real request and change nothing. The response has the
	// status code and the schema of the real call, and an empty list of tasks. With
	// `false`, the default, the request does the operation.
	DryRun param.Opt[bool] `query:"dry_run,omitzero" json:"-"`
	paramObj
}

func (r PlacementGroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PlacementGroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlacementGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PlacementGroupNewParams]'s query parameters as
// `url.Values`.
func (r PlacementGroupNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Placement policy: affinity, anti-affinity, soft-affinity, or soft-anti-affinity.
type PlacementGroupNewParamsPolicy string

const (
	PlacementGroupNewParamsPolicyAffinity         PlacementGroupNewParamsPolicy = "affinity"
	PlacementGroupNewParamsPolicyAntiAffinity     PlacementGroupNewParamsPolicy = "anti-affinity"
	PlacementGroupNewParamsPolicySoftAffinity     PlacementGroupNewParamsPolicy = "soft-affinity"
	PlacementGroupNewParamsPolicySoftAntiAffinity PlacementGroupNewParamsPolicy = "soft-anti-affinity"
)

type PlacementGroupListParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Maximum number of items to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter placement groups that a service operates (true) or not (false).
	Managed param.Opt[bool] `query:"managed,omitzero" json:"-"`
	// Number of items to skip from the start of the result.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by one or more placement group IDs (exact match).
	IDs []string `query:"ids,omitzero" format:"uuid4" json:"-"`
	// Filter by the type of the managing service.
	//
	// Any of "system".
	ManagedBy PlacementGroupListParamsManagedBy `query:"managed_by,omitzero" json:"-"`
	// Filter to placement groups containing any of these instance UUIDs.
	MemberIDs []string `query:"member_ids,omitzero" format:"uuid4" json:"-"`
	// Operator-based name filter, e.g. `name[contains]=web`, `name[prefix]=prod-`.
	// Case-insensitive.
	Name PlacementGroupListParamsName `query:"name,omitzero" json:"-"`
	// Sort field and direction.
	//
	// Any of "created_at.asc", "created_at.desc", "name.asc", "name.desc",
	// "region_id.asc", "region_id.desc", "updated_at.asc", "updated_at.desc".
	OrderBy PlacementGroupListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Filter by one or more placement policies.
	//
	// Any of "affinity", "anti-affinity", "soft-affinity", "soft-anti-affinity".
	Policies []string `query:"policies,omitzero" json:"-"`
	// Filter by one or more region IDs. Omit to list across all regions.
	RegionIDs []int64 `query:"region_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PlacementGroupListParams]'s query parameters as
// `url.Values`.
func (r PlacementGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by the type of the managing service.
type PlacementGroupListParamsManagedBy string

const (
	PlacementGroupListParamsManagedBySystem PlacementGroupListParamsManagedBy = "system"
)

// Operator-based name filter, e.g. `name[contains]=web`, `name[prefix]=prod-`.
// Case-insensitive.
type PlacementGroupListParamsName struct {
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

// URLQuery serializes [PlacementGroupListParamsName]'s query parameters as
// `url.Values`.
func (r PlacementGroupListParamsName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort field and direction.
type PlacementGroupListParamsOrderBy string

const (
	PlacementGroupListParamsOrderByCreatedAtAsc  PlacementGroupListParamsOrderBy = "created_at.asc"
	PlacementGroupListParamsOrderByCreatedAtDesc PlacementGroupListParamsOrderBy = "created_at.desc"
	PlacementGroupListParamsOrderByNameAsc       PlacementGroupListParamsOrderBy = "name.asc"
	PlacementGroupListParamsOrderByNameDesc      PlacementGroupListParamsOrderBy = "name.desc"
	PlacementGroupListParamsOrderByRegionIDAsc   PlacementGroupListParamsOrderBy = "region_id.asc"
	PlacementGroupListParamsOrderByRegionIDDesc  PlacementGroupListParamsOrderBy = "region_id.desc"
	PlacementGroupListParamsOrderByUpdatedAtAsc  PlacementGroupListParamsOrderBy = "updated_at.asc"
	PlacementGroupListParamsOrderByUpdatedAtDesc PlacementGroupListParamsOrderBy = "updated_at.desc"
)

type PlacementGroupDeleteParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Run every check of the real request and change nothing. The response has the
	// status code and the schema of the real call, and an empty list of tasks. With
	// `false`, the default, the request does the operation.
	DryRun param.Opt[bool] `query:"dry_run,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PlacementGroupDeleteParams]'s query parameters as
// `url.Values`.
func (r PlacementGroupDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PlacementGroupGetParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	paramObj
}
