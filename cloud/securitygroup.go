// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// SecurityGroupService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecurityGroupService] method instead.
type SecurityGroupService struct {
	Options []option.RequestOption
	Rules   SecurityGroupRuleService
}

// NewSecurityGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecurityGroupService(opts ...option.RequestOption) (r SecurityGroupService) {
	r = SecurityGroupService{}
	r.Options = opts
	r.Rules = NewSecurityGroupRuleService(opts...)
	return
}

// Create security group
func (r *SecurityGroupService) New(ctx context.Context, params SecurityGroupNewParams, opts ...option.RequestOption) (res *SecurityGroup, err error) {
	opts = append(r.Options[:], opts...)
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
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Change security group
func (r *SecurityGroupService) Update(ctx context.Context, groupID string, params SecurityGroupUpdateParams, opts ...option.RequestOption) (res *SecurityGroup, err error) {
	opts = append(r.Options[:], opts...)
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
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get security groups
func (r *SecurityGroupService) List(ctx context.Context, params SecurityGroupListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SecurityGroup], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v", params.ProjectID.Value, params.RegionID.Value)
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

// Get security groups
func (r *SecurityGroupService) ListAutoPaging(ctx context.Context, params SecurityGroupListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SecurityGroup] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete security group
func (r *SecurityGroupService) Delete(ctx context.Context, groupID string, body SecurityGroupDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Create a deep copy of security group
func (r *SecurityGroupService) Copy(ctx context.Context, groupID string, params SecurityGroupCopyParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s/copy", params.ProjectID.Value, params.RegionID.Value, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Get security group
func (r *SecurityGroupService) Get(ctx context.Context, groupID string, query SecurityGroupGetParams, opts ...option.RequestOption) (res *SecurityGroup, err error) {
	opts = append(r.Options[:], opts...)
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
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Revert security group
func (r *SecurityGroupService) RevertToDefault(ctx context.Context, groupID string, body SecurityGroupRevertToDefaultParams, opts ...option.RequestOption) (res *SecurityGroup, err error) {
	opts = append(r.Options[:], opts...)
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
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s/revert", body.ProjectID.Value, body.RegionID.Value, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type SecurityGroup struct {
	// Security group ID
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the security group was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Security group name
	Name string `json:"name,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// The number of revisions
	RevisionNumber int64 `json:"revision_number,required"`
	// Tags for a security group
	TagsV2 []Tag `json:"tags_v2,required"`
	// Datetime when the security group was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Security group description
	Description string `json:"description,nullable"`
	// Security group rules
	SecurityGroupRules []SecurityGroupRule `json:"security_group_rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Name               respjson.Field
		ProjectID          respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		RevisionNumber     respjson.Field
		TagsV2             respjson.Field
		UpdatedAt          respjson.Field
		Description        respjson.Field
		SecurityGroupRules respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *SecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityGroupRule struct {
	// The ID of the security group rule
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the rule was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Ingress or egress, which is the direction in which the security group rule is
	// applied
	//
	// Any of "egress", "ingress".
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
	//
	// Any of "IPv4", "IPv6".
	Ethertype SecurityGroupRuleEthertype `json:"ethertype,nullable"`
	// The maximum port number in the range that is matched by the security group rule
	PortRangeMax int64 `json:"port_range_max,nullable"`
	// The minimum port number in the range that is matched by the security group rule
	PortRangeMin int64 `json:"port_range_min,nullable"`
	// Protocol
	//
	// Any of "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap",
	// "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts",
	// "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp".
	Protocol SecurityGroupRuleProtocol `json:"protocol,nullable"`
	// The remote group UUID to associate with this security group rule
	RemoteGroupID string `json:"remote_group_id,nullable" format:"uuid4"`
	// The remote IP prefix that is matched by this security group rule
	RemoteIPPrefix string `json:"remote_ip_prefix,nullable" format:"ipvanynetwork"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Direction       respjson.Field
		RevisionNumber  respjson.Field
		SecurityGroupID respjson.Field
		UpdatedAt       respjson.Field
		Description     respjson.Field
		Ethertype       respjson.Field
		PortRangeMax    respjson.Field
		PortRangeMin    respjson.Field
		Protocol        respjson.Field
		RemoteGroupID   respjson.Field
		RemoteIPPrefix  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityGroupRule) RawJSON() string { return r.JSON.raw }
func (r *SecurityGroupRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ingress or egress, which is the direction in which the security group rule is
// applied
type SecurityGroupRuleDirection string

const (
	SecurityGroupRuleDirectionEgress  SecurityGroupRuleDirection = "egress"
	SecurityGroupRuleDirectionIngress SecurityGroupRuleDirection = "ingress"
)

// Must be IPv4 or IPv6, and addresses represented in CIDR must match the ingress
// or egress rules.
type SecurityGroupRuleEthertype string

const (
	SecurityGroupRuleEthertypeIPv4 SecurityGroupRuleEthertype = "IPv4"
	SecurityGroupRuleEthertypeIPv6 SecurityGroupRuleEthertype = "IPv6"
)

// Protocol
type SecurityGroupRuleProtocol string

const (
	SecurityGroupRuleProtocolAh        SecurityGroupRuleProtocol = "ah"
	SecurityGroupRuleProtocolAny       SecurityGroupRuleProtocol = "any"
	SecurityGroupRuleProtocolDccp      SecurityGroupRuleProtocol = "dccp"
	SecurityGroupRuleProtocolEgp       SecurityGroupRuleProtocol = "egp"
	SecurityGroupRuleProtocolEsp       SecurityGroupRuleProtocol = "esp"
	SecurityGroupRuleProtocolGre       SecurityGroupRuleProtocol = "gre"
	SecurityGroupRuleProtocolIcmp      SecurityGroupRuleProtocol = "icmp"
	SecurityGroupRuleProtocolIgmp      SecurityGroupRuleProtocol = "igmp"
	SecurityGroupRuleProtocolIpencap   SecurityGroupRuleProtocol = "ipencap"
	SecurityGroupRuleProtocolIpip      SecurityGroupRuleProtocol = "ipip"
	SecurityGroupRuleProtocolIpv6Encap SecurityGroupRuleProtocol = "ipv6-encap"
	SecurityGroupRuleProtocolIpv6Frag  SecurityGroupRuleProtocol = "ipv6-frag"
	SecurityGroupRuleProtocolIpv6Icmp  SecurityGroupRuleProtocol = "ipv6-icmp"
	SecurityGroupRuleProtocolIpv6Nonxt SecurityGroupRuleProtocol = "ipv6-nonxt"
	SecurityGroupRuleProtocolIpv6Opts  SecurityGroupRuleProtocol = "ipv6-opts"
	SecurityGroupRuleProtocolIpv6Route SecurityGroupRuleProtocol = "ipv6-route"
	SecurityGroupRuleProtocolOspf      SecurityGroupRuleProtocol = "ospf"
	SecurityGroupRuleProtocolPgm       SecurityGroupRuleProtocol = "pgm"
	SecurityGroupRuleProtocolRsvp      SecurityGroupRuleProtocol = "rsvp"
	SecurityGroupRuleProtocolSctp      SecurityGroupRuleProtocol = "sctp"
	SecurityGroupRuleProtocolTcp       SecurityGroupRuleProtocol = "tcp"
	SecurityGroupRuleProtocolUdp       SecurityGroupRuleProtocol = "udp"
	SecurityGroupRuleProtocolUdplite   SecurityGroupRuleProtocol = "udplite"
	SecurityGroupRuleProtocolVrrp      SecurityGroupRuleProtocol = "vrrp"
)

type SecurityGroupNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Security group
	SecurityGroup SecurityGroupNewParamsSecurityGroup `json:"security_group,omitzero,required"`
	// List of instances
	Instances []string `json:"instances,omitzero"`
	paramObj
}

func (r SecurityGroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Security group
//
// The property Name is required.
type SecurityGroupNewParamsSecurityGroup struct {
	// Security group name
	Name string `json:"name,required"`
	// Security group description
	Description param.Opt[string] `json:"description,omitzero"`
	// Security group rules
	SecurityGroupRules []SecurityGroupNewParamsSecurityGroupSecurityGroupRule `json:"security_group_rules,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags map[string]any `json:"tags,omitzero"`
	paramObj
}

func (r SecurityGroupNewParamsSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupNewParamsSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

type SecurityGroupNewParamsSecurityGroupSecurityGroupRule struct {
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
	Direction string `json:"direction,omitzero"`
	// Ether type
	//
	// Any of "IPv4", "IPv6".
	Ethertype string `json:"ethertype,omitzero"`
	// Protocol
	//
	// Any of "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap",
	// "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts",
	// "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp".
	Protocol string `json:"protocol,omitzero"`
	paramObj
}

func (r SecurityGroupNewParamsSecurityGroupSecurityGroupRule) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupNewParamsSecurityGroupSecurityGroupRule
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[SecurityGroupNewParamsSecurityGroupSecurityGroupRule](
		"Direction", false, "egress", "ingress",
	)
	apijson.RegisterFieldValidator[SecurityGroupNewParamsSecurityGroupSecurityGroupRule](
		"Ethertype", false, "IPv4", "IPv6",
	)
	apijson.RegisterFieldValidator[SecurityGroupNewParamsSecurityGroupSecurityGroupRule](
		"Protocol", false, "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap", "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts", "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp",
	)
}

type SecurityGroupUpdateParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Name
	Name param.Opt[string] `json:"name,omitzero"`
	// List of rules to create or delete
	ChangedRules []SecurityGroupUpdateParamsChangedRule `json:"changed_rules,omitzero"`
	paramObj
}

func (r SecurityGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// The property Action is required.
type SecurityGroupUpdateParamsChangedRule struct {
	// Action for a rule
	//
	// Any of "create", "delete".
	Action string `json:"action,omitzero,required"`
	// The remote group UUID to associate with this security group rule
	RemoteGroupID param.Opt[string] `json:"remote_group_id,omitzero" format:"uuid4"`
	// The remote IP prefix that is matched by this security group rule
	RemoteIPPrefix param.Opt[string] `json:"remote_ip_prefix,omitzero" format:"ipvanynetwork"`
	// UUID of rule to be deleted. Required for action 'delete' only
	SecurityGroupRuleID param.Opt[string] `json:"security_group_rule_id,omitzero" format:"uuid4"`
	// Security grpup rule description
	Description param.Opt[string] `json:"description,omitzero"`
	// The maximum port number in the range that is matched by the security group rule
	PortRangeMax param.Opt[int64] `json:"port_range_max,omitzero"`
	// The minimum port number in the range that is matched by the security group rule
	PortRangeMin param.Opt[int64] `json:"port_range_min,omitzero"`
	// Must be IPv4 or IPv6, and addresses represented in CIDR must match the ingress
	// or egress rules.
	//
	// Any of "IPv4", "IPv6".
	Ethertype string `json:"ethertype,omitzero"`
	// Ingress or egress, which is the direction in which the security group rule is
	// applied
	//
	// Any of "egress", "ingress".
	Direction string `json:"direction,omitzero"`
	// Protocol
	//
	// Any of "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap",
	// "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts",
	// "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp".
	Protocol string `json:"protocol,omitzero"`
	paramObj
}

func (r SecurityGroupUpdateParamsChangedRule) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupUpdateParamsChangedRule
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[SecurityGroupUpdateParamsChangedRule](
		"Action", false, "create", "delete",
	)
	apijson.RegisterFieldValidator[SecurityGroupUpdateParamsChangedRule](
		"Direction", false, "egress", "ingress",
	)
	apijson.RegisterFieldValidator[SecurityGroupUpdateParamsChangedRule](
		"Ethertype", true, "IPv4", "IPv6",
	)
	apijson.RegisterFieldValidator[SecurityGroupUpdateParamsChangedRule](
		"Protocol", false, "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap", "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts", "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp",
	)
}

type SecurityGroupListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Limit the number of returned limit request entities.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset value is used to exclude the first set of records from the result.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by tag key-value pairs. Must be a valid JSON string. curl -G
	// --data-urlencode "tag_key_value={"key": "value"}" --url
	// "http://localhost:1111/v1/securitygroups/1/1"
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// Filter by tag keys.
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SecurityGroupListParams]'s query parameters as
// `url.Values`.
func (r SecurityGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SecurityGroupDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type SecurityGroupCopyParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Name.
	Name string `json:"name,required"`
	paramObj
}

func (r SecurityGroupCopyParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupCopyParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type SecurityGroupGetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type SecurityGroupRevertToDefaultParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
