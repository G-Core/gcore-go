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
	shimjson "github.com/G-Core/gcore-go/internal/encoding/json"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
)

// Pool members represent backend instances that receive load-balanced traffic from
// a pool.
//
// LoadBalancerPoolMemberService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoadBalancerPoolMemberService] method instead.
type LoadBalancerPoolMemberService struct {
	Options []option.RequestOption
}

// NewLoadBalancerPoolMemberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerPoolMemberService(opts ...option.RequestOption) (r LoadBalancerPoolMemberService) {
	r = LoadBalancerPoolMemberService{}
	r.Options = opts
	return
}

// Create load balancer pool member
func (r *LoadBalancerPoolMemberService) New(ctx context.Context, poolID string, params LoadBalancerPoolMemberNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/pools/%s/members", params.ProjectID.Value, params.RegionID.Value, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Updates the specified member's mutable settings. `address`, `protocol_port` and
// `subnet_id` cannot be changed after creation. If no changes are detected, no
// task is created and an empty task list is returned.
func (r *LoadBalancerPoolMemberService) Update(ctx context.Context, memberID string, params LoadBalancerPoolMemberUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if params.PoolID == "" {
		err = errors.New("missing required pool_id parameter")
		return nil, err
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/pools/%s/members/%s", params.ProjectID.Value, params.RegionID.Value, params.PoolID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List load balancer pool members
func (r *LoadBalancerPoolMemberService) List(ctx context.Context, poolID string, params LoadBalancerPoolMemberListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Member], err error) {
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
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/pools/%s/members", params.ProjectID.Value, params.RegionID.Value, poolID)
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

// List load balancer pool members
func (r *LoadBalancerPoolMemberService) ListAutoPaging(ctx context.Context, poolID string, params LoadBalancerPoolMemberListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Member] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, poolID, params, opts...))
}

// Delete load balancer pool member
func (r *LoadBalancerPoolMemberService) Delete(ctx context.Context, memberID string, body LoadBalancerPoolMemberDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if body.PoolID == "" {
		err = errors.New("missing required pool_id parameter")
		return nil, err
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/pools/%s/members/%s", body.ProjectID.Value, body.RegionID.Value, body.PoolID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get load balancer pool member
func (r *LoadBalancerPoolMemberService) Get(ctx context.Context, memberID string, query LoadBalancerPoolMemberGetParams, opts ...option.RequestOption) (res *Member, err error) {
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
	if query.PoolID == "" {
		err = errors.New("missing required pool_id parameter")
		return nil, err
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/pools/%s/members/%s", query.ProjectID.Value, query.RegionID.Value, query.PoolID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Replaces the full set of pool members with the provided list (declarative,
// desired-state semantics): members present in the payload are created or updated,
// members currently on the pool but absent from the payload are deleted. If a
// member is unchanged (same address + port), it is kept as is without recreation
// and downtime. If no changes are detected, no task is created and an empty task
// list is returned.
//
// For updating a single existing member without affecting the rest of the pool,
// use
// `PATCH /v1/loadbalancers/{project_id}/{region_id}/pools/{pool_id}/members/{member_id}`
// instead.
func (r *LoadBalancerPoolMemberService) Replace(ctx context.Context, poolID string, params LoadBalancerPoolMemberReplaceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/pools/%s/members", params.ProjectID.Value, params.RegionID.Value, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

type LoadBalancerPoolMemberNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Member IP address
	Address string `json:"address" api:"required" format:"ipvanyaddress"`
	// Member IP port
	ProtocolPort int64 `json:"protocol_port" api:"required"`
	// Either `subnet_id` or `instance_id` should be provided
	InstanceID param.Opt[string] `json:"instance_id,omitzero" format:"uuid4"`
	// An alternate IP address used for health monitoring of a backend member. Default
	// is null which monitors the member address.
	MonitorAddress param.Opt[string] `json:"monitor_address,omitzero" format:"ipvanyaddress"`
	// An alternate protocol port used for health monitoring of a backend member.
	// Default is null which monitors the member `protocol_port`.
	MonitorPort param.Opt[int64] `json:"monitor_port,omitzero"`
	// `subnet_id` in which `address` is present. Either `subnet_id` or `instance_id`
	// should be provided
	SubnetID param.Opt[string] `json:"subnet_id,omitzero" format:"uuid4"`
	// Administrative state of the resource. When set to true, the resource is enabled
	// and operational. When set to false, the resource is disabled and will not
	// process traffic. Defaults to true.
	AdminStateUp param.Opt[bool] `json:"admin_state_up,omitzero"`
	// Set to true if the member is a backup member, to which traffic will be sent
	// exclusively when all non-backup members will be unreachable. It allows to
	// realize ACTIVE-BACKUP load balancing without thinking about VRRP and VIP
	// configuration. Default is false.
	Backup param.Opt[bool] `json:"backup,omitzero"`
	// Member weight. Valid values are 0 < `weight` <= 256, defaults to 1. Controls
	// traffic distribution based on the pool's load balancing algorithm:
	//
	//   - `ROUND_ROBIN`: Distributes connections to each member in turn according to
	//     weights. Higher weight = more turns in the cycle. Example: weights 3 vs 1 =
	//     ~75% vs ~25% of requests.
	//   - `LEAST_CONNECTIONS`: Sends new connections to the member with fewest active
	//     connections, performing round-robin within groups of the same normalized load.
	//     Higher weight = allowed to hold more simultaneous connections before being
	//     considered 'more loaded'. Example: weights 2 vs 1 means 20 vs 10 active
	//     connections is treated as balanced.
	//   - `SOURCE_IP`: Routes clients consistently to the same member by hashing client
	//     source IP; hash result is modulo total weight of running members. Higher
	//     weight = more hash buckets, so more client IPs map to that member. Example:
	//     weights 2 vs 1 = roughly two-thirds of distinct client IPs map to the
	//     higher-weight member.
	Weight param.Opt[int64] `json:"weight,omitzero"`
	paramObj
}

func (r LoadBalancerPoolMemberNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolMemberNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerPoolMemberNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolMemberUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Pool ID
	PoolID string `path:"pool_id" api:"required" format:"uuid4" json:"-"`
	// Alternate IP address used for health monitoring of a backend member. Set to
	// `null` to clear it and fall back to the member address; omit to leave unchanged.
	MonitorAddress param.Opt[string] `json:"monitor_address,omitzero" format:"ipvanyaddress"`
	// Alternate protocol port used for health monitoring of a backend member. Set to
	// `null` to clear it and fall back to the member `protocol_port`; omit to leave
	// unchanged.
	MonitorPort param.Opt[int64] `json:"monitor_port,omitzero"`
	// Administrative state of the member. Omit to leave unchanged; `false` disables
	// the member so it receives no traffic.
	AdminStateUp param.Opt[bool] `json:"admin_state_up,omitzero"`
	// Set to true if the member is a backup member, to which traffic will be sent
	// exclusively when all non-backup members will be unreachable. Omit to leave
	// unchanged.
	Backup param.Opt[bool] `json:"backup,omitzero"`
	// Member weight. Valid values are 0 < `weight` <= 256. Omit to leave unchanged.
	// Controls traffic distribution based on the pool's load balancing algorithm:
	//
	//   - `ROUND_ROBIN`: Distributes connections to each member in turn according to
	//     weights. Higher weight = more turns in the cycle. Example: weights 3 vs 1 =
	//     ~75% vs ~25% of requests.
	//   - `LEAST_CONNECTIONS`: Sends new connections to the member with fewest active
	//     connections, performing round-robin within groups of the same normalized load.
	//     Higher weight = allowed to hold more simultaneous connections before being
	//     considered 'more loaded'. Example: weights 2 vs 1 means 20 vs 10 active
	//     connections is treated as balanced.
	//   - `SOURCE_IP`: Routes clients consistently to the same member by hashing client
	//     source IP; hash result is modulo total weight of running members. Higher
	//     weight = more hash buckets, so more client IPs map to that member. Example:
	//     weights 2 vs 1 = roughly two-thirds of distinct client IPs map to the
	//     higher-weight member.
	Weight param.Opt[int64] `json:"weight,omitzero"`
	paramObj
}

func (r LoadBalancerPoolMemberUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolMemberUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerPoolMemberUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolMemberListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Ordering pool members list result by `address` or `created_at` fields and
	// directions (e.g. `address.desc`). Default is `address.asc`.
	//
	// Any of "address.asc", "address.desc", "created_at.asc", "created_at.desc".
	OrderBy LoadBalancerPoolMemberListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LoadBalancerPoolMemberListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerPoolMemberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Ordering pool members list result by `address` or `created_at` fields and
// directions (e.g. `address.desc`). Default is `address.asc`.
type LoadBalancerPoolMemberListParamsOrderBy string

const (
	LoadBalancerPoolMemberListParamsOrderByAddressAsc    LoadBalancerPoolMemberListParamsOrderBy = "address.asc"
	LoadBalancerPoolMemberListParamsOrderByAddressDesc   LoadBalancerPoolMemberListParamsOrderBy = "address.desc"
	LoadBalancerPoolMemberListParamsOrderByCreatedAtAsc  LoadBalancerPoolMemberListParamsOrderBy = "created_at.asc"
	LoadBalancerPoolMemberListParamsOrderByCreatedAtDesc LoadBalancerPoolMemberListParamsOrderBy = "created_at.desc"
)

type LoadBalancerPoolMemberDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Pool ID
	PoolID string `path:"pool_id" api:"required" format:"uuid4" json:"-"`
	paramObj
}

type LoadBalancerPoolMemberGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Pool ID
	PoolID string `path:"pool_id" api:"required" format:"uuid4" json:"-"`
	paramObj
}

type LoadBalancerPoolMemberReplaceParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// New sequence of load balancer pool members. If members are the same (by
	// address + port), they will be kept as is without recreation and downtime.
	Body []LoadBalancerPoolMemberReplaceParamsBody
	paramObj
}

func (r LoadBalancerPoolMemberReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *LoadBalancerPoolMemberReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Address, ProtocolPort are required.
type LoadBalancerPoolMemberReplaceParamsBody struct {
	// Member IP address
	Address string `json:"address" api:"required" format:"ipvanyaddress"`
	// Member IP port
	ProtocolPort int64 `json:"protocol_port" api:"required"`
	// Either `subnet_id` or `instance_id` should be provided
	InstanceID param.Opt[string] `json:"instance_id,omitzero" format:"uuid4"`
	// An alternate IP address used for health monitoring of a backend member. Default
	// is null which monitors the member address.
	MonitorAddress param.Opt[string] `json:"monitor_address,omitzero" format:"ipvanyaddress"`
	// An alternate protocol port used for health monitoring of a backend member.
	// Default is null which monitors the member `protocol_port`.
	MonitorPort param.Opt[int64] `json:"monitor_port,omitzero"`
	// `subnet_id` in which `address` is present. Either `subnet_id` or `instance_id`
	// should be provided
	SubnetID param.Opt[string] `json:"subnet_id,omitzero" format:"uuid4"`
	// Administrative state of the resource. When set to true, the resource is enabled
	// and operational. When set to false, the resource is disabled and will not
	// process traffic. Defaults to true.
	AdminStateUp param.Opt[bool] `json:"admin_state_up,omitzero"`
	// Set to true if the member is a backup member, to which traffic will be sent
	// exclusively when all non-backup members will be unreachable. It allows to
	// realize ACTIVE-BACKUP load balancing without thinking about VRRP and VIP
	// configuration. Default is false.
	Backup param.Opt[bool] `json:"backup,omitzero"`
	// Member weight. Valid values are 0 < `weight` <= 256, defaults to 1. Controls
	// traffic distribution based on the pool's load balancing algorithm:
	//
	//   - `ROUND_ROBIN`: Distributes connections to each member in turn according to
	//     weights. Higher weight = more turns in the cycle. Example: weights 3 vs 1 =
	//     ~75% vs ~25% of requests.
	//   - `LEAST_CONNECTIONS`: Sends new connections to the member with fewest active
	//     connections, performing round-robin within groups of the same normalized load.
	//     Higher weight = allowed to hold more simultaneous connections before being
	//     considered 'more loaded'. Example: weights 2 vs 1 means 20 vs 10 active
	//     connections is treated as balanced.
	//   - `SOURCE_IP`: Routes clients consistently to the same member by hashing client
	//     source IP; hash result is modulo total weight of running members. Higher
	//     weight = more hash buckets, so more client IPs map to that member. Example:
	//     weights 2 vs 1 = roughly two-thirds of distinct client IPs map to the
	//     higher-weight member.
	Weight param.Opt[int64] `json:"weight,omitzero"`
	paramObj
}

func (r LoadBalancerPoolMemberReplaceParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolMemberReplaceParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerPoolMemberReplaceParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
