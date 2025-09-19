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
)

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
func (r *LoadBalancerPoolMemberService) Add(ctx context.Context, poolID string, params LoadBalancerPoolMemberAddParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s/member", params.ProjectID.Value, params.RegionID.Value, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete load balancer pool member
func (r *LoadBalancerPoolMemberService) Remove(ctx context.Context, memberID string, body LoadBalancerPoolMemberRemoveParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if body.PoolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s/member/%s", body.ProjectID.Value, body.RegionID.Value, body.PoolID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type LoadBalancerPoolMemberAddParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Member IP address
	Address string `json:"address,required" format:"ipvanyaddress"`
	// Member IP port
	ProtocolPort int64 `json:"protocol_port,required"`
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
	// process traffic. When null is passed, the value is skipped and defaults to true.
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

func (r LoadBalancerPoolMemberAddParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolMemberAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerPoolMemberAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolMemberRemoveParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Pool ID
	PoolID string `path:"pool_id,required" format:"uuid4" json:"-"`
	paramObj
}
