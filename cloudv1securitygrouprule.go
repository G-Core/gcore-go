// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1SecuritygroupruleService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1SecuritygroupruleService] method instead.
type CloudV1SecuritygroupruleService struct {
	Options []option.RequestOption
}

// NewCloudV1SecuritygroupruleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1SecuritygroupruleService(opts ...option.RequestOption) (r *CloudV1SecuritygroupruleService) {
	r = &CloudV1SecuritygroupruleService{}
	r.Options = opts
	return
}

// Edit the security group rule: delete old and create new rule
func (r *CloudV1SecuritygroupruleService) Update(ctx context.Context, projectID int64, regionID int64, pk string, body CloudV1SecuritygroupruleUpdateParams, opts ...option.RequestOption) (res *SecurityGroupRule, err error) {
	opts = append(r.Options[:], opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygrouprules/%v/%v/%s", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete security group rule
func (r *CloudV1SecuritygroupruleService) Delete(ctx context.Context, projectID int64, regionID int64, pk string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygrouprules/%v/%v/%s", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type SecurityGroupProtocol string

const (
	SecurityGroupProtocolAh        SecurityGroupProtocol = "ah"
	SecurityGroupProtocolAny       SecurityGroupProtocol = "any"
	SecurityGroupProtocolDccp      SecurityGroupProtocol = "dccp"
	SecurityGroupProtocolEgp       SecurityGroupProtocol = "egp"
	SecurityGroupProtocolEsp       SecurityGroupProtocol = "esp"
	SecurityGroupProtocolGre       SecurityGroupProtocol = "gre"
	SecurityGroupProtocolIcmp      SecurityGroupProtocol = "icmp"
	SecurityGroupProtocolIgmp      SecurityGroupProtocol = "igmp"
	SecurityGroupProtocolIpencap   SecurityGroupProtocol = "ipencap"
	SecurityGroupProtocolIpip      SecurityGroupProtocol = "ipip"
	SecurityGroupProtocolIpv6Encap SecurityGroupProtocol = "ipv6-encap"
	SecurityGroupProtocolIpv6Frag  SecurityGroupProtocol = "ipv6-frag"
	SecurityGroupProtocolIpv6Icmp  SecurityGroupProtocol = "ipv6-icmp"
	SecurityGroupProtocolIpv6Nonxt SecurityGroupProtocol = "ipv6-nonxt"
	SecurityGroupProtocolIpv6Opts  SecurityGroupProtocol = "ipv6-opts"
	SecurityGroupProtocolIpv6Route SecurityGroupProtocol = "ipv6-route"
	SecurityGroupProtocolOspf      SecurityGroupProtocol = "ospf"
	SecurityGroupProtocolPgm       SecurityGroupProtocol = "pgm"
	SecurityGroupProtocolRsvp      SecurityGroupProtocol = "rsvp"
	SecurityGroupProtocolSctp      SecurityGroupProtocol = "sctp"
	SecurityGroupProtocolTcp       SecurityGroupProtocol = "tcp"
	SecurityGroupProtocolUdp       SecurityGroupProtocol = "udp"
	SecurityGroupProtocolUdplite   SecurityGroupProtocol = "udplite"
	SecurityGroupProtocolVrrp      SecurityGroupProtocol = "vrrp"
)

func (r SecurityGroupProtocol) IsKnown() bool {
	switch r {
	case SecurityGroupProtocolAh, SecurityGroupProtocolAny, SecurityGroupProtocolDccp, SecurityGroupProtocolEgp, SecurityGroupProtocolEsp, SecurityGroupProtocolGre, SecurityGroupProtocolIcmp, SecurityGroupProtocolIgmp, SecurityGroupProtocolIpencap, SecurityGroupProtocolIpip, SecurityGroupProtocolIpv6Encap, SecurityGroupProtocolIpv6Frag, SecurityGroupProtocolIpv6Icmp, SecurityGroupProtocolIpv6Nonxt, SecurityGroupProtocolIpv6Opts, SecurityGroupProtocolIpv6Route, SecurityGroupProtocolOspf, SecurityGroupProtocolPgm, SecurityGroupProtocolRsvp, SecurityGroupProtocolSctp, SecurityGroupProtocolTcp, SecurityGroupProtocolUdp, SecurityGroupProtocolUdplite, SecurityGroupProtocolVrrp:
		return true
	}
	return false
}

type SecurityGroupRule struct {
	// The ID of the security group rule
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the rule was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Ingress or egress, which is the direction in which the security group rule is
	// applied
	Direction SecurityGroupRuleDirection `json:"direction,required"`
	// The revision number of the resource
	RevisionNumber int64 `json:"revision_number,required"`
	// The security group ID to associate with this security group rule
	SecurityGroupID string `json:"security_group_id,required" format:"uuid4"`
	// Datetime when the rule was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Rule description
	Description string `json:"description,nullable"`
	// Must be IPv4 or IPv6, and addresses represented in CIDR must match the ingress
	// or egress rules.
	Ethertype SecurityGroupRuleEthertype `json:"ethertype,nullable"`
	// The maximum port number in the range that is matched by the security group rule
	PortRangeMax int64 `json:"port_range_max,nullable"`
	// The minimum port number in the range that is matched by the security group rule
	PortRangeMin int64 `json:"port_range_min,nullable"`
	// Protocol
	Protocol SecurityGroupProtocol `json:"protocol,nullable"`
	// The remote group UUID to associate with this security group rule
	RemoteGroupID string `json:"remote_group_id,nullable" format:"uuid4"`
	// The remote IP prefix that is matched by this security group rule
	RemoteIPPrefix string                `json:"remote_ip_prefix,nullable" format:"ipvanynetwork"`
	JSON           securityGroupRuleJSON `json:"-"`
}

// securityGroupRuleJSON contains the JSON metadata for the struct
// [SecurityGroupRule]
type securityGroupRuleJSON struct {
	ID              apijson.Field
	CreatedAt       apijson.Field
	Direction       apijson.Field
	RevisionNumber  apijson.Field
	SecurityGroupID apijson.Field
	UpdatedAt       apijson.Field
	Description     apijson.Field
	Ethertype       apijson.Field
	PortRangeMax    apijson.Field
	PortRangeMin    apijson.Field
	Protocol        apijson.Field
	RemoteGroupID   apijson.Field
	RemoteIPPrefix  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SecurityGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityGroupRuleJSON) RawJSON() string {
	return r.raw
}

// Ingress or egress, which is the direction in which the security group rule is
// applied
type SecurityGroupRuleDirection string

const (
	SecurityGroupRuleDirectionEgress  SecurityGroupRuleDirection = "egress"
	SecurityGroupRuleDirectionIngress SecurityGroupRuleDirection = "ingress"
)

func (r SecurityGroupRuleDirection) IsKnown() bool {
	switch r {
	case SecurityGroupRuleDirectionEgress, SecurityGroupRuleDirectionIngress:
		return true
	}
	return false
}

// Must be IPv4 or IPv6, and addresses represented in CIDR must match the ingress
// or egress rules.
type SecurityGroupRuleEthertype string

const (
	SecurityGroupRuleEthertypeIPv4 SecurityGroupRuleEthertype = "IPv4"
	SecurityGroupRuleEthertypeIPv6 SecurityGroupRuleEthertype = "IPv6"
)

func (r SecurityGroupRuleEthertype) IsKnown() bool {
	switch r {
	case SecurityGroupRuleEthertypeIPv4, SecurityGroupRuleEthertypeIPv6:
		return true
	}
	return false
}

type CloudV1SecuritygroupruleUpdateParams struct {
	// Ingress or egress, which is the direction in which the security group rule is
	// applied
	Direction param.Field[CloudV1SecuritygroupruleUpdateParamsDirection] `json:"direction,required"`
	// Parent security group of this rule
	SecurityGroupID param.Field[string] `json:"security_group_id,required" format:"uuid4"`
	// Rule description
	Description param.Field[string] `json:"description"`
	// Must be IPv4 or IPv6, and addresses represented in CIDR must match the ingress
	// or egress rules.
	Ethertype param.Field[CloudV1SecuritygroupruleUpdateParamsEthertype] `json:"ethertype"`
	// The maximum port number in the range that is matched by the security group rule
	PortRangeMax param.Field[int64] `json:"port_range_max"`
	// The minimum port number in the range that is matched by the security group rule
	PortRangeMin param.Field[int64] `json:"port_range_min"`
	// Protocol
	Protocol param.Field[SecurityGroupProtocol] `json:"protocol"`
	// The remote group UUID to associate with this security group rule
	RemoteGroupID param.Field[string] `json:"remote_group_id" format:"uuid4"`
	// The remote IP prefix that is matched by this security group rule
	RemoteIPPrefix param.Field[string] `json:"remote_ip_prefix" format:"ipvanynetwork"`
}

func (r CloudV1SecuritygroupruleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Ingress or egress, which is the direction in which the security group rule is
// applied
type CloudV1SecuritygroupruleUpdateParamsDirection string

const (
	CloudV1SecuritygroupruleUpdateParamsDirectionEgress  CloudV1SecuritygroupruleUpdateParamsDirection = "egress"
	CloudV1SecuritygroupruleUpdateParamsDirectionIngress CloudV1SecuritygroupruleUpdateParamsDirection = "ingress"
)

func (r CloudV1SecuritygroupruleUpdateParamsDirection) IsKnown() bool {
	switch r {
	case CloudV1SecuritygroupruleUpdateParamsDirectionEgress, CloudV1SecuritygroupruleUpdateParamsDirectionIngress:
		return true
	}
	return false
}

// Must be IPv4 or IPv6, and addresses represented in CIDR must match the ingress
// or egress rules.
type CloudV1SecuritygroupruleUpdateParamsEthertype string

const (
	CloudV1SecuritygroupruleUpdateParamsEthertypeIPv4 CloudV1SecuritygroupruleUpdateParamsEthertype = "IPv4"
	CloudV1SecuritygroupruleUpdateParamsEthertypeIPv6 CloudV1SecuritygroupruleUpdateParamsEthertype = "IPv6"
)

func (r CloudV1SecuritygroupruleUpdateParamsEthertype) IsKnown() bool {
	switch r {
	case CloudV1SecuritygroupruleUpdateParamsEthertypeIPv4, CloudV1SecuritygroupruleUpdateParamsEthertypeIPv6:
		return true
	}
	return false
}
