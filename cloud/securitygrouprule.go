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

// SecurityGroupRuleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecurityGroupRuleService] method instead.
type SecurityGroupRuleService struct {
	Options []option.RequestOption
}

// NewSecurityGroupRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSecurityGroupRuleService(opts ...option.RequestOption) (r SecurityGroupRuleService) {
	r = SecurityGroupRuleService{}
	r.Options = opts
	return
}

// Add a new rule to an existing security group. Returns a task ID for tracking the
// asynchronous operation.
func (r *SecurityGroupRuleService) New(ctx context.Context, groupID string, params SecurityGroupRuleNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/security_groups/%v/%v/%s/rules", params.ProjectID.Value, params.RegionID.Value, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete a specific rule from a security group. Returns a task ID for tracking the
// asynchronous operation.
func (r *SecurityGroupRuleService) Delete(ctx context.Context, ruleID string, body SecurityGroupRuleDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if body.GroupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/security_groups/%v/%v/%s/rules/%s", body.ProjectID.Value, body.RegionID.Value, body.GroupID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type SecurityGroupRuleNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Ingress or egress, which is the direction in which the security group is applied
	//
	// Any of "egress", "ingress".
	Direction SecurityGroupRuleNewParamsDirection `json:"direction,omitzero,required"`
	// The maximum port number in the range that is matched by the security group rule
	PortRangeMax param.Opt[int64] `json:"port_range_max,omitzero"`
	// The minimum port number in the range that is matched by the security group rule
	PortRangeMin param.Opt[int64] `json:"port_range_min,omitzero"`
	// The remote IP prefix that is matched by this security group rule
	RemoteIPPrefix param.Opt[string] `json:"remote_ip_prefix,omitzero" format:"ipvanynetwork"`
	// Rule description
	Description param.Opt[string] `json:"description,omitzero"`
	// The remote group UUID to associate with this security group
	RemoteGroupID param.Opt[string] `json:"remote_group_id,omitzero" format:"uuid4"`
	// Ether type
	//
	// Any of "IPv4", "IPv6".
	Ethertype SecurityGroupRuleNewParamsEthertype `json:"ethertype,omitzero"`
	// Protocol
	//
	// Any of "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap",
	// "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts",
	// "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp".
	Protocol SecurityGroupRuleNewParamsProtocol `json:"protocol,omitzero"`
	paramObj
}

func (r SecurityGroupRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecurityGroupRuleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ingress or egress, which is the direction in which the security group is applied
type SecurityGroupRuleNewParamsDirection string

const (
	SecurityGroupRuleNewParamsDirectionEgress  SecurityGroupRuleNewParamsDirection = "egress"
	SecurityGroupRuleNewParamsDirectionIngress SecurityGroupRuleNewParamsDirection = "ingress"
)

// Ether type
type SecurityGroupRuleNewParamsEthertype string

const (
	SecurityGroupRuleNewParamsEthertypeIPv4 SecurityGroupRuleNewParamsEthertype = "IPv4"
	SecurityGroupRuleNewParamsEthertypeIPv6 SecurityGroupRuleNewParamsEthertype = "IPv6"
)

// Protocol
type SecurityGroupRuleNewParamsProtocol string

const (
	SecurityGroupRuleNewParamsProtocolAh        SecurityGroupRuleNewParamsProtocol = "ah"
	SecurityGroupRuleNewParamsProtocolAny       SecurityGroupRuleNewParamsProtocol = "any"
	SecurityGroupRuleNewParamsProtocolDccp      SecurityGroupRuleNewParamsProtocol = "dccp"
	SecurityGroupRuleNewParamsProtocolEgp       SecurityGroupRuleNewParamsProtocol = "egp"
	SecurityGroupRuleNewParamsProtocolEsp       SecurityGroupRuleNewParamsProtocol = "esp"
	SecurityGroupRuleNewParamsProtocolGre       SecurityGroupRuleNewParamsProtocol = "gre"
	SecurityGroupRuleNewParamsProtocolIcmp      SecurityGroupRuleNewParamsProtocol = "icmp"
	SecurityGroupRuleNewParamsProtocolIgmp      SecurityGroupRuleNewParamsProtocol = "igmp"
	SecurityGroupRuleNewParamsProtocolIpencap   SecurityGroupRuleNewParamsProtocol = "ipencap"
	SecurityGroupRuleNewParamsProtocolIpip      SecurityGroupRuleNewParamsProtocol = "ipip"
	SecurityGroupRuleNewParamsProtocolIpv6Encap SecurityGroupRuleNewParamsProtocol = "ipv6-encap"
	SecurityGroupRuleNewParamsProtocolIpv6Frag  SecurityGroupRuleNewParamsProtocol = "ipv6-frag"
	SecurityGroupRuleNewParamsProtocolIpv6Icmp  SecurityGroupRuleNewParamsProtocol = "ipv6-icmp"
	SecurityGroupRuleNewParamsProtocolIpv6Nonxt SecurityGroupRuleNewParamsProtocol = "ipv6-nonxt"
	SecurityGroupRuleNewParamsProtocolIpv6Opts  SecurityGroupRuleNewParamsProtocol = "ipv6-opts"
	SecurityGroupRuleNewParamsProtocolIpv6Route SecurityGroupRuleNewParamsProtocol = "ipv6-route"
	SecurityGroupRuleNewParamsProtocolOspf      SecurityGroupRuleNewParamsProtocol = "ospf"
	SecurityGroupRuleNewParamsProtocolPgm       SecurityGroupRuleNewParamsProtocol = "pgm"
	SecurityGroupRuleNewParamsProtocolRsvp      SecurityGroupRuleNewParamsProtocol = "rsvp"
	SecurityGroupRuleNewParamsProtocolSctp      SecurityGroupRuleNewParamsProtocol = "sctp"
	SecurityGroupRuleNewParamsProtocolTcp       SecurityGroupRuleNewParamsProtocol = "tcp"
	SecurityGroupRuleNewParamsProtocolUdp       SecurityGroupRuleNewParamsProtocol = "udp"
	SecurityGroupRuleNewParamsProtocolUdplite   SecurityGroupRuleNewParamsProtocol = "udplite"
	SecurityGroupRuleNewParamsProtocolVrrp      SecurityGroupRuleNewParamsProtocol = "vrrp"
)

type SecurityGroupRuleDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Security group ID
	GroupID string `path:"group_id,required" format:"uuid4" json:"-"`
	paramObj
}
