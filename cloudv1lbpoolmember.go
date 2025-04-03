// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1LbpoolMemberService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1LbpoolMemberService] method instead.
type CloudV1LbpoolMemberService struct {
	Options []option.RequestOption
}

// NewCloudV1LbpoolMemberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1LbpoolMemberService(opts ...option.RequestOption) (r *CloudV1LbpoolMemberService) {
	r = &CloudV1LbpoolMemberService{}
	r.Options = opts
	return
}

// Create load balancer pool member
func (r *CloudV1LbpoolMemberService) New(ctx context.Context, projectID int64, regionID int64, poolID string, body CloudV1LbpoolMemberNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s/member", projectID, regionID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete load balancer pool member
func (r *CloudV1LbpoolMemberService) Delete(ctx context.Context, projectID int64, regionID int64, poolID string, memberID string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbpools/%v/%v/%s/member/%s", projectID, regionID, poolID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CreateLbPoolMemberParam struct {
	// Member IP address
	Address param.Field[string] `json:"address,required" format:"ipvanyaddress"`
	// Member IP port
	ProtocolPort param.Field[int64] `json:"protocol_port,required"`
	// true if enabled. Defaults to true
	AdminStateUp param.Field[bool] `json:"admin_state_up"`
	// Either subnet_id or instance_id should be provided
	InstanceID param.Field[string] `json:"instance_id" format:"uuid4"`
	// An alternate IP address used for health monitoring of a backend member. Default
	// is null which monitors the member address.
	MonitorAddress param.Field[string] `json:"monitor_address" format:"ipvanyaddress"`
	// An alternate protocol port used for health monitoring of a backend member.
	// Default is null which monitors the member protocol_port.
	MonitorPort param.Field[int64] `json:"monitor_port"`
	// Either subnet_id or instance_id should be provided
	SubnetID param.Field[string] `json:"subnet_id" format:"uuid4"`
	// Member weight. Valid values: 0 to 256, defaults to 1
	Weight param.Field[int64] `json:"weight"`
}

func (r CreateLbPoolMemberParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1LbpoolMemberNewParams struct {
	CreateLbPoolMember CreateLbPoolMemberParam `json:"create_lb_pool_member,required"`
}

func (r CloudV1LbpoolMemberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateLbPoolMember)
}
