// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
)

// Networks provide software-defined networking infrastructure for connecting
// instances and other cloud resources within a region.
//
// NetworkService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkService] method instead.
type NetworkService struct {
	Options []option.RequestOption
	// Subnets define IP address ranges within a network for instance connectivity,
	// with support for DHCP and DNS configuration.
	Subnets NetworkSubnetService
	// Routers interconnect subnets and manage network routing, including external
	// gateway connectivity and static routes.
	Routers NetworkRouterService
	tasks   TaskService
}

// NewNetworkService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNetworkService(opts ...option.RequestOption) (r NetworkService) {
	r = NetworkService{}
	r.Options = opts
	r.Subnets = NewNetworkSubnetService(opts...)
	r.Routers = NewNetworkRouterService(opts...)
	r.tasks = NewTaskService(opts...)
	return
}

// Create network
func (r *NetworkService) New(ctx context.Context, params NetworkNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

func (r *NetworkService) NewAndPoll(ctx context.Context, params NetworkNewParams, opts ...option.RequestOption) (res *Network, err error) {
	// Exclude WithResponseBodyInto for the action (New returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.New(ctx, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams NetworkGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	task, err := r.tasks.Poll(ctx, taskID, pollOpts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Networks) != 1 {
		return nil, errors.New("expected exactly one network to be created in a task")
	}
	resourceID := task.CreatedResources.Networks[0]

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, resourceID, getParams, getOpts...)
}

// Rename network and/or update network tags. The request will only process the
// fields that are provided in the request body. Any fields that are not included
// will remain unchanged.
func (r *NetworkService) Update(ctx context.Context, networkID string, params NetworkUpdateParams, opts ...option.RequestOption) (res *Network, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a list of networks. Use the `owned_by` query parameter to control which
// networks are returned: `project` (default) returns only networks owned by the
// project, `any` returns all networks the project can use, including shared
// networks.
func (r *NetworkService) List(ctx context.Context, params NetworkListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Network], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v", params.ProjectID.Value, params.RegionID.Value)
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

// Returns a list of networks. Use the `owned_by` query parameter to control which
// networks are returned: `project` (default) returns only networks owned by the
// project, `any` returns all networks the project can use, including shared
// networks.
func (r *NetworkService) ListAutoPaging(ctx context.Context, params NetworkListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Network] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete network
func (r *NetworkService) Delete(ctx context.Context, networkID string, body NetworkDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// DeleteAndPoll deletes a network and polls for completion of the first task. Use the
// [TaskService.Poll] method if you need to poll for all tasks.
func (r *NetworkService) DeleteAndPoll(ctx context.Context, networkID string, body NetworkDeleteParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (Delete returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Delete(ctx, networkID, body, actionOpts...)
	if err != nil {
		return err
	}

	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	_, err = r.tasks.Poll(ctx, taskID, pollOpts...)
	return err
}

// Get network
func (r *NetworkService) Get(ctx context.Context, networkID string, query NetworkGetParams, opts ...option.RequestOption) (res *Network, err error) {
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
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type NetworkNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Network name
	Name string `json:"name" api:"required"`
	// Defaults to True
	CreateRouter param.Opt[bool] `json:"create_router,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Some tags are read-only and cannot be modified by the
	// user. Tags are also integrated with cost reports, allowing cost data to be
	// filtered based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	// vlan or vxlan network type is allowed. Default value is vxlan
	//
	// Any of "vlan", "vxlan".
	Type NetworkNewParamsType `json:"type,omitzero"`
	paramObj
}

func (r NetworkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// vlan or vxlan network type is allowed. Default value is vxlan
type NetworkNewParamsType string

const (
	NetworkNewParamsTypeVlan  NetworkNewParamsType = "vlan"
	NetworkNewParamsTypeVxlan NetworkNewParamsType = "vxlan"
)

type NetworkUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Update key-value tags using JSON Merge Patch semantics (RFC 7386). Provide
	// key-value pairs to add or update tags. Set tag values to `null` to remove tags.
	// Unspecified tags remain unchanged. Read-only tags are always preserved and
	// cannot be modified.
	//
	// **Examples:**
	//
	//   - **Add/update tags:**
	//     `{'tags': {'environment': 'production', 'team': 'backend'}}` adds new tags or
	//     updates existing ones.
	//   - **Delete tags:** `{'tags': {'old_tag': null}}` removes specific tags.
	//   - **Remove all tags:** `{'tags': null}` removes all user-managed tags (read-only
	//     tags are preserved).
	//   - **Partial update:** `{'tags': {'environment': 'staging'}}` only updates
	//     specified tags.
	//   - **Mixed operations:**
	//     `{'tags': {'environment': 'production', 'cost_center': 'engineering', 'deprecated_tag': null}}`
	//     adds/updates 'environment' and 'cost_center' while removing 'deprecated_tag',
	//     preserving other existing tags.
	//   - **Replace all:** first delete existing tags with null values, then add new
	//     ones in the same request.
	Tags TagUpdateMap `json:"tags,omitzero"`
	paramObj
}

func (r NetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Filter by external network status
	External param.Opt[bool] `query:"external,omitzero" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter networks by name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Optional. Filter by tag key-value pairs.
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// Filter by network type (vlan or vxlan)
	//
	// Any of "vlan", "vxlan".
	NetworkType NetworkListParamsNetworkType `query:"network_type,omitzero" json:"-"`
	// Ordering networks list result by `name`, `created_at` or `priority` fields and
	// directions (e.g. `created_at.desc`). Default is `created_at.desc`. Use
	// `priority.desc` to sort by shared network priority (relevant when
	// `owned_by=any`).
	//
	// Any of "created_at.asc", "created_at.desc", "name.asc", "name.desc",
	// "priority.desc".
	OrderBy NetworkListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Controls which networks are returned. 'project' (default) returns only networks
	// owned by the project. 'any' returns all networks that the project can use,
	// including shared networks from other projects.
	//
	// Any of "any", "project".
	OwnedBy NetworkListParamsOwnedBy `query:"owned_by,omitzero" json:"-"`
	// Optional. Filter by tag keys. ?`tag_key`=key1&`tag_key`=key2
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkListParams]'s query parameters as `url.Values`.
func (r NetworkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by network type (vlan or vxlan)
type NetworkListParamsNetworkType string

const (
	NetworkListParamsNetworkTypeVlan  NetworkListParamsNetworkType = "vlan"
	NetworkListParamsNetworkTypeVxlan NetworkListParamsNetworkType = "vxlan"
)

// Ordering networks list result by `name`, `created_at` or `priority` fields and
// directions (e.g. `created_at.desc`). Default is `created_at.desc`. Use
// `priority.desc` to sort by shared network priority (relevant when
// `owned_by=any`).
type NetworkListParamsOrderBy string

const (
	NetworkListParamsOrderByCreatedAtAsc  NetworkListParamsOrderBy = "created_at.asc"
	NetworkListParamsOrderByCreatedAtDesc NetworkListParamsOrderBy = "created_at.desc"
	NetworkListParamsOrderByNameAsc       NetworkListParamsOrderBy = "name.asc"
	NetworkListParamsOrderByNameDesc      NetworkListParamsOrderBy = "name.desc"
	NetworkListParamsOrderByPriorityDesc  NetworkListParamsOrderBy = "priority.desc"
)

// Controls which networks are returned. 'project' (default) returns only networks
// owned by the project. 'any' returns all networks that the project can use,
// including shared networks from other projects.
type NetworkListParamsOwnedBy string

const (
	NetworkListParamsOwnedByAny     NetworkListParamsOwnedBy = "any"
	NetworkListParamsOwnedByProject NetworkListParamsOwnedBy = "project"
)

type NetworkDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type NetworkGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}
