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

// Add new rule to security group
func (r *SecurityGroupRuleService) New(ctx context.Context, groupID string, params SecurityGroupRuleNewParams, opts ...option.RequestOption) (res *SecurityGroupRule, err error) {
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
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s/rules", params.ProjectID.Value, params.RegionID.Value, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete security group rule
func (r *SecurityGroupRuleService) Delete(ctx context.Context, ruleID string, body SecurityGroupRuleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygrouprules/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Edit the security group rule: delete old and create new rule
func (r *SecurityGroupRuleService) Replace(ctx context.Context, ruleID string, params SecurityGroupRuleReplaceParams, opts ...option.RequestOption) (res *SecurityGroupRule, err error) {
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
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygrouprules/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type SecurityGroupRuleNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D%2Frules/post/parameters/0/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}/{group_id}/rules'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D%2Frules/post/parameters/1/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}/{group_id}/rules'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateSecurityGroupRuleSerializer/properties/port_range_max/anyOf/0'
	// "$.components.schemas.CreateSecurityGroupRuleSerializer.properties.port_range_max.anyOf[0]"
	PortRangeMax param.Opt[int64] `json:"port_range_max,omitzero"`
	// '#/components/schemas/CreateSecurityGroupRuleSerializer/properties/port_range_min/anyOf/0'
	// "$.components.schemas.CreateSecurityGroupRuleSerializer.properties.port_range_min.anyOf[0]"
	PortRangeMin param.Opt[int64] `json:"port_range_min,omitzero"`
	// '#/components/schemas/CreateSecurityGroupRuleSerializer/properties/remote_ip_prefix/anyOf/0'
	// "$.components.schemas.CreateSecurityGroupRuleSerializer.properties.remote_ip_prefix.anyOf[0]"
	RemoteIPPrefix param.Opt[string] `json:"remote_ip_prefix,omitzero" format:"ipvanynetwork"`
	// '#/components/schemas/CreateSecurityGroupRuleSerializer/properties/description'
	// "$.components.schemas.CreateSecurityGroupRuleSerializer.properties.description"
	Description param.Opt[string] `json:"description,omitzero"`
	// '#/components/schemas/CreateSecurityGroupRuleSerializer/properties/remote_group_id'
	// "$.components.schemas.CreateSecurityGroupRuleSerializer.properties.remote_group_id"
	RemoteGroupID param.Opt[string] `json:"remote_group_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateSecurityGroupRuleSerializer/properties/direction'
	// "$.components.schemas.CreateSecurityGroupRuleSerializer.properties.direction"
	//
	// Any of "egress", "ingress".
	Direction SecurityGroupRuleNewParamsDirection `json:"direction,omitzero"`
	// '#/components/schemas/CreateSecurityGroupRuleSerializer/properties/ethertype'
	// "$.components.schemas.CreateSecurityGroupRuleSerializer.properties.ethertype"
	//
	// Any of "IPv4", "IPv6".
	Ethertype SecurityGroupRuleNewParamsEthertype `json:"ethertype,omitzero"`
	// '#/components/schemas/CreateSecurityGroupRuleSerializer/properties/protocol'
	// "$.components.schemas.CreateSecurityGroupRuleSerializer.properties.protocol"
	//
	// Any of "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap",
	// "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts",
	// "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp".
	Protocol SecurityGroupRuleNewParamsProtocol `json:"protocol,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecurityGroupRuleNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SecurityGroupRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateSecurityGroupRuleSerializer/properties/direction'
// "$.components.schemas.CreateSecurityGroupRuleSerializer.properties.direction"
type SecurityGroupRuleNewParamsDirection string

const (
	SecurityGroupRuleNewParamsDirectionEgress  SecurityGroupRuleNewParamsDirection = "egress"
	SecurityGroupRuleNewParamsDirectionIngress SecurityGroupRuleNewParamsDirection = "ingress"
)

// '#/components/schemas/CreateSecurityGroupRuleSerializer/properties/ethertype'
// "$.components.schemas.CreateSecurityGroupRuleSerializer.properties.ethertype"
type SecurityGroupRuleNewParamsEthertype string

const (
	SecurityGroupRuleNewParamsEthertypeIPv4 SecurityGroupRuleNewParamsEthertype = "IPv4"
	SecurityGroupRuleNewParamsEthertypeIPv6 SecurityGroupRuleNewParamsEthertype = "IPv6"
)

// '#/components/schemas/CreateSecurityGroupRuleSerializer/properties/protocol'
// "$.components.schemas.CreateSecurityGroupRuleSerializer.properties.protocol"
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
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygrouprules%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brule_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/securitygrouprules/{project_id}/{region_id}/{rule_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygrouprules%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brule_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/securitygrouprules/{project_id}/{region_id}/{rule_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecurityGroupRuleDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type SecurityGroupRuleReplaceParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygrouprules%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brule_id%7D/put/parameters/0/schema'
	// "$.paths['/cloud/v1/securitygrouprules/{project_id}/{region_id}/{rule_id}'].put.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygrouprules%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Brule_id%7D/put/parameters/1/schema'
	// "$.paths['/cloud/v1/securitygrouprules/{project_id}/{region_id}/{rule_id}'].put.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/ChangeSecurityGroupRuleSerializer/properties/direction'
	// "$.components.schemas.ChangeSecurityGroupRuleSerializer.properties.direction"
	//
	// Any of "egress", "ingress".
	Direction SecurityGroupRuleReplaceParamsDirection `json:"direction,omitzero,required"`
	// '#/components/schemas/ChangeSecurityGroupRuleSerializer/properties/security_group_id'
	// "$.components.schemas.ChangeSecurityGroupRuleSerializer.properties.security_group_id"
	SecurityGroupID string `json:"security_group_id,required" format:"uuid4"`
	// '#/components/schemas/ChangeSecurityGroupRuleSerializer/properties/port_range_max/anyOf/0'
	// "$.components.schemas.ChangeSecurityGroupRuleSerializer.properties.port_range_max.anyOf[0]"
	PortRangeMax param.Opt[int64] `json:"port_range_max,omitzero"`
	// '#/components/schemas/ChangeSecurityGroupRuleSerializer/properties/port_range_min/anyOf/0'
	// "$.components.schemas.ChangeSecurityGroupRuleSerializer.properties.port_range_min.anyOf[0]"
	PortRangeMin param.Opt[int64] `json:"port_range_min,omitzero"`
	// '#/components/schemas/ChangeSecurityGroupRuleSerializer/properties/remote_group_id/anyOf/0'
	// "$.components.schemas.ChangeSecurityGroupRuleSerializer.properties.remote_group_id.anyOf[0]"
	RemoteGroupID param.Opt[string] `json:"remote_group_id,omitzero" format:"uuid4"`
	// '#/components/schemas/ChangeSecurityGroupRuleSerializer/properties/remote_ip_prefix/anyOf/0'
	// "$.components.schemas.ChangeSecurityGroupRuleSerializer.properties.remote_ip_prefix.anyOf[0]"
	RemoteIPPrefix param.Opt[string] `json:"remote_ip_prefix,omitzero" format:"ipvanynetwork"`
	// '#/components/schemas/ChangeSecurityGroupRuleSerializer/properties/description'
	// "$.components.schemas.ChangeSecurityGroupRuleSerializer.properties.description"
	Description param.Opt[string] `json:"description,omitzero"`
	// '#/components/schemas/ChangeSecurityGroupRuleSerializer/properties/ethertype/anyOf/0'
	// "$.components.schemas.ChangeSecurityGroupRuleSerializer.properties.ethertype.anyOf[0]"
	//
	// Any of "IPv4", "IPv6".
	Ethertype SecurityGroupRuleReplaceParamsEthertype `json:"ethertype,omitzero"`
	// '#/components/schemas/ChangeSecurityGroupRuleSerializer/properties/protocol'
	// "$.components.schemas.ChangeSecurityGroupRuleSerializer.properties.protocol"
	//
	// Any of "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap",
	// "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts",
	// "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp".
	Protocol SecurityGroupRuleReplaceParamsProtocol `json:"protocol,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecurityGroupRuleReplaceParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SecurityGroupRuleReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupRuleReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ChangeSecurityGroupRuleSerializer/properties/direction'
// "$.components.schemas.ChangeSecurityGroupRuleSerializer.properties.direction"
type SecurityGroupRuleReplaceParamsDirection string

const (
	SecurityGroupRuleReplaceParamsDirectionEgress  SecurityGroupRuleReplaceParamsDirection = "egress"
	SecurityGroupRuleReplaceParamsDirectionIngress SecurityGroupRuleReplaceParamsDirection = "ingress"
)

// '#/components/schemas/ChangeSecurityGroupRuleSerializer/properties/ethertype/anyOf/0'
// "$.components.schemas.ChangeSecurityGroupRuleSerializer.properties.ethertype.anyOf[0]"
type SecurityGroupRuleReplaceParamsEthertype string

const (
	SecurityGroupRuleReplaceParamsEthertypeIPv4 SecurityGroupRuleReplaceParamsEthertype = "IPv4"
	SecurityGroupRuleReplaceParamsEthertypeIPv6 SecurityGroupRuleReplaceParamsEthertype = "IPv6"
)

// '#/components/schemas/ChangeSecurityGroupRuleSerializer/properties/protocol'
// "$.components.schemas.ChangeSecurityGroupRuleSerializer.properties.protocol"
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
