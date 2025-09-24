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

// Add a new rule to an existing security group.
func (r *SecurityGroupRuleService) New(ctx context.Context, groupID string, params SecurityGroupRuleNewParams, opts ...option.RequestOption) (res *SecurityGroupRule, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
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
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s/rules", params.ProjectID.Value, params.RegionID.Value, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete a specific rule from a security group.
func (r *SecurityGroupRuleService) Delete(ctx context.Context, ruleID string, body SecurityGroupRuleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
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
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygrouprules/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Update the configuration of an existing security group rule.
func (r *SecurityGroupRuleService) Replace(ctx context.Context, ruleID string, params SecurityGroupRuleReplaceParams, opts ...option.RequestOption) (res *SecurityGroupRule, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
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
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygrouprules/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type SecurityGroupRuleNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
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
	// Ingress or egress, which is the direction in which the security group is applied
	//
	// Any of "egress", "ingress".
	Direction SecurityGroupRuleNewParamsDirection `json:"direction,omitzero"`
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type SecurityGroupRuleReplaceParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Ingress or egress, which is the direction in which the security group rule is
	// applied
	//
	// Any of "egress", "ingress".
	Direction SecurityGroupRuleReplaceParamsDirection `json:"direction,omitzero,required"`
	// Parent security group of this rule
	SecurityGroupID string `json:"security_group_id,required" format:"uuid4"`
	// The maximum port number in the range that is matched by the security group rule
	PortRangeMax param.Opt[int64] `json:"port_range_max,omitzero"`
	// The minimum port number in the range that is matched by the security group rule
	PortRangeMin param.Opt[int64] `json:"port_range_min,omitzero"`
	// The remote group UUID to associate with this security group rule
	RemoteGroupID param.Opt[string] `json:"remote_group_id,omitzero" format:"uuid4"`
	// The remote IP prefix that is matched by this security group rule
	RemoteIPPrefix param.Opt[string] `json:"remote_ip_prefix,omitzero" format:"ipvanynetwork"`
	// Rule description
	Description param.Opt[string] `json:"description,omitzero"`
	// Must be IPv4 or IPv6, and addresses represented in CIDR must match the ingress
	// or egress rules.
	//
	// Any of "IPv4", "IPv6".
	Ethertype SecurityGroupRuleReplaceParamsEthertype `json:"ethertype,omitzero"`
	// Protocol
	//
	// Any of "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap",
	// "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts",
	// "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp".
	Protocol SecurityGroupRuleReplaceParamsProtocol `json:"protocol,omitzero"`
	paramObj
}

func (r SecurityGroupRuleReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupRuleReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecurityGroupRuleReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ingress or egress, which is the direction in which the security group rule is
// applied
type SecurityGroupRuleReplaceParamsDirection string

const (
	SecurityGroupRuleReplaceParamsDirectionEgress  SecurityGroupRuleReplaceParamsDirection = "egress"
	SecurityGroupRuleReplaceParamsDirectionIngress SecurityGroupRuleReplaceParamsDirection = "ingress"
)

// Must be IPv4 or IPv6, and addresses represented in CIDR must match the ingress
// or egress rules.
type SecurityGroupRuleReplaceParamsEthertype string

const (
	SecurityGroupRuleReplaceParamsEthertypeIPv4 SecurityGroupRuleReplaceParamsEthertype = "IPv4"
	SecurityGroupRuleReplaceParamsEthertypeIPv6 SecurityGroupRuleReplaceParamsEthertype = "IPv6"
)

// Protocol
type SecurityGroupRuleReplaceParamsProtocol string

const (
	SecurityGroupRuleReplaceParamsProtocolAh        SecurityGroupRuleReplaceParamsProtocol = "ah"
	SecurityGroupRuleReplaceParamsProtocolAny       SecurityGroupRuleReplaceParamsProtocol = "any"
	SecurityGroupRuleReplaceParamsProtocolDccp      SecurityGroupRuleReplaceParamsProtocol = "dccp"
	SecurityGroupRuleReplaceParamsProtocolEgp       SecurityGroupRuleReplaceParamsProtocol = "egp"
	SecurityGroupRuleReplaceParamsProtocolEsp       SecurityGroupRuleReplaceParamsProtocol = "esp"
	SecurityGroupRuleReplaceParamsProtocolGre       SecurityGroupRuleReplaceParamsProtocol = "gre"
	SecurityGroupRuleReplaceParamsProtocolIcmp      SecurityGroupRuleReplaceParamsProtocol = "icmp"
	SecurityGroupRuleReplaceParamsProtocolIgmp      SecurityGroupRuleReplaceParamsProtocol = "igmp"
	SecurityGroupRuleReplaceParamsProtocolIpencap   SecurityGroupRuleReplaceParamsProtocol = "ipencap"
	SecurityGroupRuleReplaceParamsProtocolIpip      SecurityGroupRuleReplaceParamsProtocol = "ipip"
	SecurityGroupRuleReplaceParamsProtocolIpv6Encap SecurityGroupRuleReplaceParamsProtocol = "ipv6-encap"
	SecurityGroupRuleReplaceParamsProtocolIpv6Frag  SecurityGroupRuleReplaceParamsProtocol = "ipv6-frag"
	SecurityGroupRuleReplaceParamsProtocolIpv6Icmp  SecurityGroupRuleReplaceParamsProtocol = "ipv6-icmp"
	SecurityGroupRuleReplaceParamsProtocolIpv6Nonxt SecurityGroupRuleReplaceParamsProtocol = "ipv6-nonxt"
	SecurityGroupRuleReplaceParamsProtocolIpv6Opts  SecurityGroupRuleReplaceParamsProtocol = "ipv6-opts"
	SecurityGroupRuleReplaceParamsProtocolIpv6Route SecurityGroupRuleReplaceParamsProtocol = "ipv6-route"
	SecurityGroupRuleReplaceParamsProtocolOspf      SecurityGroupRuleReplaceParamsProtocol = "ospf"
	SecurityGroupRuleReplaceParamsProtocolPgm       SecurityGroupRuleReplaceParamsProtocol = "pgm"
	SecurityGroupRuleReplaceParamsProtocolRsvp      SecurityGroupRuleReplaceParamsProtocol = "rsvp"
	SecurityGroupRuleReplaceParamsProtocolSctp      SecurityGroupRuleReplaceParamsProtocol = "sctp"
	SecurityGroupRuleReplaceParamsProtocolTcp       SecurityGroupRuleReplaceParamsProtocol = "tcp"
	SecurityGroupRuleReplaceParamsProtocolUdp       SecurityGroupRuleReplaceParamsProtocol = "udp"
	SecurityGroupRuleReplaceParamsProtocolUdplite   SecurityGroupRuleReplaceParamsProtocol = "udplite"
	SecurityGroupRuleReplaceParamsProtocolVrrp      SecurityGroupRuleReplaceParamsProtocol = "vrrp"
)
