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
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s/member", params.ProjectID.Value, params.RegionID.Value, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete load balancer pool member
func (r *LoadBalancerPoolMemberService) Remove(ctx context.Context, poolID string, memberID string, body LoadBalancerPoolMemberRemoveParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s/member/%s", body.ProjectID.Value, body.RegionID.Value, poolID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type LoadBalancerPoolMemberAddParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bpool_id%7D%2Fmember/post/parameters/0/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}/{pool_id}/member'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bpool_id%7D%2Fmember/post/parameters/1/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}/{pool_id}/member'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/address'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.address"
	Address string `json:"address,required" format:"ipvanyaddress"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/protocol_port'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.protocol_port"
	ProtocolPort int64 `json:"protocol_port,required"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/admin_state_up/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.admin_state_up.anyOf[0]"
	AdminStateUp param.Opt[bool] `json:"admin_state_up,omitzero"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/instance_id/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.instance_id.anyOf[0]"
	InstanceID param.Opt[string] `json:"instance_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/monitor_address/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.monitor_address.anyOf[0]"
	MonitorAddress param.Opt[string] `json:"monitor_address,omitzero" format:"ipvanyaddress"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/monitor_port/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.monitor_port.anyOf[0]"
	MonitorPort param.Opt[int64] `json:"monitor_port,omitzero"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/subnet_id/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.subnet_id.anyOf[0]"
	SubnetID param.Opt[string] `json:"subnet_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLbPoolMemberSerializer/properties/weight/anyOf/0'
	// "$.components.schemas.CreateLbPoolMemberSerializer.properties.weight.anyOf[0]"
	Weight param.Opt[int64] `json:"weight,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerPoolMemberAddParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LoadBalancerPoolMemberAddParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerPoolMemberAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type LoadBalancerPoolMemberRemoveParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bpool_id%7D%2Fmember%2F%7Bmember_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}/{pool_id}/member/{member_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flbpools%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bpool_id%7D%2Fmember%2F%7Bmember_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/lbpools/{project_id}/{region_id}/{pool_id}/member/{member_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerPoolMemberRemoveParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
