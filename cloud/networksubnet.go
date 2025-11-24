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

// NetworkSubnetService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkSubnetService] method instead.
type NetworkSubnetService struct {
	Options []option.RequestOption
	tasks   TaskService
}

// NewNetworkSubnetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkSubnetService(opts ...option.RequestOption) (r NetworkSubnetService) {
	r = NetworkSubnetService{}
	r.Options = opts
	r.tasks = NewTaskService(opts...)
	return
}

// Create subnet
func (r *NetworkSubnetService) New(ctx context.Context, params NetworkSubnetNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// NewAndPoll creates a subnet and then polls the task until it's completed.
func (r *NetworkSubnetService) NewAndPoll(ctx context.Context, params NetworkSubnetNewParams, opts ...option.RequestOption) (res *Subnet, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams NetworkSubnetGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]
	task, err := r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Subnets) != 1 {
		return nil, errors.New("expected exactly one subnet to be created in a task")
	}
	resourceID := task.CreatedResources.Subnets[0]

	return r.Get(ctx, resourceID, getParams, opts...)
}

// Update subnet
func (r *NetworkSubnetService) Update(ctx context.Context, subnetID string, params NetworkSubnetUpdateParams, opts ...option.RequestOption) (res *Subnet, err error) {
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
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List subnets
func (r *NetworkSubnetService) List(ctx context.Context, params NetworkSubnetListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Subnet], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v", params.ProjectID.Value, params.RegionID.Value)
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

// List subnets
func (r *NetworkSubnetService) ListAutoPaging(ctx context.Context, params NetworkSubnetListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Subnet] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete subnet
func (r *NetworkSubnetService) Delete(ctx context.Context, subnetID string, body NetworkSubnetDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// DeleteAndPoll deletes a network and polls for completion of the first task. Use the
// [TaskService.Poll] method if you need to poll for all tasks.
func (r *NetworkSubnetService) DeleteAndPoll(ctx context.Context, subnetID string, body NetworkSubnetDeleteParams, opts ...option.RequestOption) (err error) {
	resource, err := r.Delete(ctx, subnetID, body, opts...)
	if err != nil {
		return err
	}

	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	return err
}

// Get subnet
func (r *NetworkSubnetService) Get(ctx context.Context, subnetID string, query NetworkSubnetGetParams, opts ...option.RequestOption) (res *Subnet, err error) {
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
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type NetworkSubnetNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// CIDR
	Cidr string `json:"cidr,required" format:"ipvanynetwork"`
	// Subnet name
	Name string `json:"name,required"`
	// Network ID
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// Default GW IPv4 address to advertise in DHCP routes in this subnet. Omit this
	// field to let the cloud backend allocate it automatically. Set to null if no
	// gateway must be advertised by this subnet's DHCP (useful when attaching
	// instances to multiple subnets in order to prevent default route conflicts).
	GatewayIP param.Opt[string] `json:"gateway_ip,omitzero" format:"ipvanyaddress"`
	// ID of the router to connect to. Requires `connect_to_network_router` set to
	// true. If not specified, attempts to find a router created during network
	// creation.
	RouterIDToConnect param.Opt[string] `json:"router_id_to_connect,omitzero" format:"uuid4"`
	// True if the network's router should get a gateway in this subnet. Must be
	// explicitly 'false' when `gateway_ip` is null.
	ConnectToNetworkRouter param.Opt[bool] `json:"connect_to_network_router,omitzero"`
	// True if DHCP should be enabled
	EnableDhcp param.Opt[bool] `json:"enable_dhcp,omitzero"`
	// List IP addresses of DNS servers to advertise via DHCP.
	DNSNameservers []string `json:"dns_nameservers,omitzero" format:"ipvanyaddress"`
	// List of custom static routes to advertise via DHCP.
	HostRoutes []NetworkSubnetNewParamsHostRoute `json:"host_routes,omitzero"`
	// IP version
	//
	// Any of 4, 6.
	IPVersion int64 `json:"ip_version,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Some tags are read-only and cannot be modified by the
	// user. Tags are also integrated with cost reports, allowing cost data to be
	// filtered based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r NetworkSubnetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkSubnetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkSubnetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Destination, Nexthop are required.
type NetworkSubnetNewParamsHostRoute struct {
	// CIDR of destination IPv4 subnet.
	Destination string `json:"destination,required" format:"ipvanynetwork"`
	// IPv4 address to forward traffic to if it's destination IP matches 'destination'
	// CIDR.
	Nexthop string `json:"nexthop,required" format:"ipvanyaddress"`
	paramObj
}

func (r NetworkSubnetNewParamsHostRoute) MarshalJSON() (data []byte, err error) {
	type shadow NetworkSubnetNewParamsHostRoute
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkSubnetNewParamsHostRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkSubnetUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// True if DHCP should be enabled
	EnableDhcp param.Opt[bool] `json:"enable_dhcp,omitzero"`
	// Default GW IPv4 address to advertise in DHCP routes in this subnet. Omit this
	// field to let the cloud backend allocate it automatically. Set to null if no
	// gateway must be advertised by this subnet's DHCP (useful when attaching
	// instances to multiple subnets in order to prevent default route conflicts).
	GatewayIP param.Opt[string] `json:"gateway_ip,omitzero" format:"ipvanyaddress"`
	// Name
	Name param.Opt[string] `json:"name,omitzero"`
	// List IP addresses of DNS servers to advertise via DHCP.
	DNSNameservers []string `json:"dns_nameservers,omitzero" format:"ipvanyaddress"`
	// List of custom static routes to advertise via DHCP.
	HostRoutes []NetworkSubnetUpdateParamsHostRoute `json:"host_routes,omitzero"`
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
	//
	// - **Delete tags:** `{'tags': {'old_tag': null}}` removes specific tags.
	//
	//   - **Remove all tags:** `{'tags': null}` removes all user-managed tags (read-only
	//     tags are preserved).
	//
	//   - **Partial update:** `{'tags': {'environment': 'staging'}}` only updates
	//     specified tags.
	//
	//   - **Mixed operations:**
	//     `{'tags': {'environment': 'production', 'cost_center': 'engineering', 'deprecated_tag': null}}`
	//     adds/updates 'environment' and '`cost_center`' while removing
	//     '`deprecated_tag`', preserving other existing tags.
	//
	//   - **Replace all:** first delete existing tags with null values, then add new
	//     ones in the same request.
	Tags TagUpdateMap `json:"tags,omitzero"`
	paramObj
}

func (r NetworkSubnetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkSubnetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkSubnetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Destination, Nexthop are required.
type NetworkSubnetUpdateParamsHostRoute struct {
	// CIDR of destination IPv4 subnet.
	Destination string `json:"destination,required" format:"ipvanynetwork"`
	// IPv4 address to forward traffic to if it's destination IP matches 'destination'
	// CIDR.
	Nexthop string `json:"nexthop,required" format:"ipvanyaddress"`
	paramObj
}

func (r NetworkSubnetUpdateParamsHostRoute) MarshalJSON() (data []byte, err error) {
	type shadow NetworkSubnetUpdateParamsHostRoute
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkSubnetUpdateParamsHostRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkSubnetListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only list subnets of this network
	NetworkID param.Opt[string] `query:"network_id,omitzero" format:"uuid4" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Optional. Filter by tag key-value pairs.
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// Ordering subnets list result by `name`, `created_at`, `updated_at`,
	// `available_ips`, `total_ips`, and `cidr` (default) fields of the subnet and
	// directions (`name.asc`).
	//
	// Any of "available_ips.asc", "available_ips.desc", "cidr.asc", "cidr.desc",
	// "created_at.asc", "created_at.desc", "name.asc", "name.desc", "total_ips.asc",
	// "total_ips.desc", "updated_at.asc", "updated_at.desc".
	OrderBy NetworkSubnetListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Optional. Filter by tag keys. ?`tag_key`=key1&`tag_key`=key2
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkSubnetListParams]'s query parameters as
// `url.Values`.
func (r NetworkSubnetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Ordering subnets list result by `name`, `created_at`, `updated_at`,
// `available_ips`, `total_ips`, and `cidr` (default) fields of the subnet and
// directions (`name.asc`).
type NetworkSubnetListParamsOrderBy string

const (
	NetworkSubnetListParamsOrderByAvailableIPsAsc  NetworkSubnetListParamsOrderBy = "available_ips.asc"
	NetworkSubnetListParamsOrderByAvailableIPsDesc NetworkSubnetListParamsOrderBy = "available_ips.desc"
	NetworkSubnetListParamsOrderByCidrAsc          NetworkSubnetListParamsOrderBy = "cidr.asc"
	NetworkSubnetListParamsOrderByCidrDesc         NetworkSubnetListParamsOrderBy = "cidr.desc"
	NetworkSubnetListParamsOrderByCreatedAtAsc     NetworkSubnetListParamsOrderBy = "created_at.asc"
	NetworkSubnetListParamsOrderByCreatedAtDesc    NetworkSubnetListParamsOrderBy = "created_at.desc"
	NetworkSubnetListParamsOrderByNameAsc          NetworkSubnetListParamsOrderBy = "name.asc"
	NetworkSubnetListParamsOrderByNameDesc         NetworkSubnetListParamsOrderBy = "name.desc"
	NetworkSubnetListParamsOrderByTotalIPsAsc      NetworkSubnetListParamsOrderBy = "total_ips.asc"
	NetworkSubnetListParamsOrderByTotalIPsDesc     NetworkSubnetListParamsOrderBy = "total_ips.desc"
	NetworkSubnetListParamsOrderByUpdatedAtAsc     NetworkSubnetListParamsOrderBy = "updated_at.asc"
	NetworkSubnetListParamsOrderByUpdatedAtDesc    NetworkSubnetListParamsOrderBy = "updated_at.desc"
)

type NetworkSubnetDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type NetworkSubnetGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
