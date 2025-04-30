// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
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
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
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
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
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
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
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

// '#/components/schemas/SecurityGroupSerializer'
// "$.components.schemas.SecurityGroupSerializer"
type SecurityGroup struct {
	// '#/components/schemas/SecurityGroupSerializer/properties/id'
	// "$.components.schemas.SecurityGroupSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/SecurityGroupSerializer/properties/created_at'
	// "$.components.schemas.SecurityGroupSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/SecurityGroupSerializer/properties/name'
	// "$.components.schemas.SecurityGroupSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/SecurityGroupSerializer/properties/project_id'
	// "$.components.schemas.SecurityGroupSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/SecurityGroupSerializer/properties/region'
	// "$.components.schemas.SecurityGroupSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/SecurityGroupSerializer/properties/region_id'
	// "$.components.schemas.SecurityGroupSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/SecurityGroupSerializer/properties/revision_number'
	// "$.components.schemas.SecurityGroupSerializer.properties.revision_number"
	RevisionNumber int64 `json:"revision_number,required"`
	// '#/components/schemas/SecurityGroupSerializer/properties/tags_v2'
	// "$.components.schemas.SecurityGroupSerializer.properties.tags_v2"
	TagsV2 []Tag `json:"tags_v2,required"`
	// '#/components/schemas/SecurityGroupSerializer/properties/updated_at'
	// "$.components.schemas.SecurityGroupSerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// '#/components/schemas/SecurityGroupSerializer/properties/description/anyOf/0'
	// "$.components.schemas.SecurityGroupSerializer.properties.description.anyOf[0]"
	Description string `json:"description,nullable"`
	// '#/components/schemas/SecurityGroupSerializer/properties/security_group_rules'
	// "$.components.schemas.SecurityGroupSerializer.properties.security_group_rules"
	SecurityGroupRules []SecurityGroupRule `json:"security_group_rules"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                 resp.Field
		CreatedAt          resp.Field
		Name               resp.Field
		ProjectID          resp.Field
		Region             resp.Field
		RegionID           resp.Field
		RevisionNumber     resp.Field
		TagsV2             resp.Field
		UpdatedAt          resp.Field
		Description        resp.Field
		SecurityGroupRules resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityGroup) RawJSON() string { return r.JSON.raw }
func (r *SecurityGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/SecurityGroupRuleSerializer'
// "$.components.schemas.SecurityGroupRuleSerializer"
type SecurityGroupRule struct {
	// '#/components/schemas/SecurityGroupRuleSerializer/properties/id'
	// "$.components.schemas.SecurityGroupRuleSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/SecurityGroupRuleSerializer/properties/created_at'
	// "$.components.schemas.SecurityGroupRuleSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/SecurityGroupRuleSerializer/properties/direction'
	// "$.components.schemas.SecurityGroupRuleSerializer.properties.direction"
	//
	// Any of "egress", "ingress".
	Direction SecurityGroupRuleDirection `json:"direction,required"`
	// '#/components/schemas/SecurityGroupRuleSerializer/properties/revision_number'
	// "$.components.schemas.SecurityGroupRuleSerializer.properties.revision_number"
	RevisionNumber int64 `json:"revision_number,required"`
	// '#/components/schemas/SecurityGroupRuleSerializer/properties/security_group_id'
	// "$.components.schemas.SecurityGroupRuleSerializer.properties.security_group_id"
	SecurityGroupID string `json:"security_group_id,required" format:"uuid4"`
	// '#/components/schemas/SecurityGroupRuleSerializer/properties/updated_at'
	// "$.components.schemas.SecurityGroupRuleSerializer.properties.updated_at"
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// '#/components/schemas/SecurityGroupRuleSerializer/properties/description/anyOf/0'
	// "$.components.schemas.SecurityGroupRuleSerializer.properties.description.anyOf[0]"
	Description string `json:"description,nullable"`
	// '#/components/schemas/SecurityGroupRuleSerializer/properties/ethertype/anyOf/0'
	// "$.components.schemas.SecurityGroupRuleSerializer.properties.ethertype.anyOf[0]"
	//
	// Any of "IPv4", "IPv6".
	Ethertype SecurityGroupRuleEthertype `json:"ethertype,nullable"`
	// '#/components/schemas/SecurityGroupRuleSerializer/properties/port_range_max/anyOf/0'
	// "$.components.schemas.SecurityGroupRuleSerializer.properties.port_range_max.anyOf[0]"
	PortRangeMax int64 `json:"port_range_max,nullable"`
	// '#/components/schemas/SecurityGroupRuleSerializer/properties/port_range_min/anyOf/0'
	// "$.components.schemas.SecurityGroupRuleSerializer.properties.port_range_min.anyOf[0]"
	PortRangeMin int64 `json:"port_range_min,nullable"`
	// '#/components/schemas/SecurityGroupRuleSerializer/properties/protocol/anyOf/0'
	// "$.components.schemas.SecurityGroupRuleSerializer.properties.protocol.anyOf[0]"
	//
	// Any of "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap",
	// "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts",
	// "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp".
	Protocol SecurityGroupRuleProtocol `json:"protocol,nullable"`
	// '#/components/schemas/SecurityGroupRuleSerializer/properties/remote_group_id/anyOf/0'
	// "$.components.schemas.SecurityGroupRuleSerializer.properties.remote_group_id.anyOf[0]"
	RemoteGroupID string `json:"remote_group_id,nullable" format:"uuid4"`
	// '#/components/schemas/SecurityGroupRuleSerializer/properties/remote_ip_prefix/anyOf/0'
	// "$.components.schemas.SecurityGroupRuleSerializer.properties.remote_ip_prefix.anyOf[0]"
	RemoteIPPrefix string `json:"remote_ip_prefix,nullable" format:"ipvanynetwork"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID              resp.Field
		CreatedAt       resp.Field
		Direction       resp.Field
		RevisionNumber  resp.Field
		SecurityGroupID resp.Field
		UpdatedAt       resp.Field
		Description     resp.Field
		Ethertype       resp.Field
		PortRangeMax    resp.Field
		PortRangeMin    resp.Field
		Protocol        resp.Field
		RemoteGroupID   resp.Field
		RemoteIPPrefix  resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecurityGroupRule) RawJSON() string { return r.JSON.raw }
func (r *SecurityGroupRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/SecurityGroupRuleSerializer/properties/direction'
// "$.components.schemas.SecurityGroupRuleSerializer.properties.direction"
type SecurityGroupRuleDirection string

const (
	SecurityGroupRuleDirectionEgress  SecurityGroupRuleDirection = "egress"
	SecurityGroupRuleDirectionIngress SecurityGroupRuleDirection = "ingress"
)

// '#/components/schemas/SecurityGroupRuleSerializer/properties/ethertype/anyOf/0'
// "$.components.schemas.SecurityGroupRuleSerializer.properties.ethertype.anyOf[0]"
type SecurityGroupRuleEthertype string

const (
	SecurityGroupRuleEthertypeIPv4 SecurityGroupRuleEthertype = "IPv4"
	SecurityGroupRuleEthertypeIPv6 SecurityGroupRuleEthertype = "IPv6"
)

// '#/components/schemas/SecurityGroupRuleSerializer/properties/protocol/anyOf/0'
// "$.components.schemas.SecurityGroupRuleSerializer.properties.protocol.anyOf[0]"
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
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateSecurityGroupSerializer/properties/security_group'
	// "$.components.schemas.CreateSecurityGroupSerializer.properties.security_group"
	SecurityGroup SecurityGroupNewParamsSecurityGroup `json:"security_group,omitzero,required"`
	// '#/components/schemas/CreateSecurityGroupSerializer/properties/instances'
	// "$.components.schemas.CreateSecurityGroupSerializer.properties.instances"
	Instances []string `json:"instances,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecurityGroupNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SecurityGroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateSecurityGroupSerializer/properties/security_group'
// "$.components.schemas.CreateSecurityGroupSerializer.properties.security_group"
//
// The property Name is required.
type SecurityGroupNewParamsSecurityGroup struct {
	// '#/components/schemas/SingleCreateSecurityGroupSerializer/properties/name'
	// "$.components.schemas.SingleCreateSecurityGroupSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/SingleCreateSecurityGroupSerializer/properties/description/anyOf/0'
	// "$.components.schemas.SingleCreateSecurityGroupSerializer.properties.description.anyOf[0]"
	Description param.Opt[string] `json:"description,omitzero"`
	// '#/components/schemas/SingleCreateSecurityGroupSerializer/properties/security_group_rules'
	// "$.components.schemas.SingleCreateSecurityGroupSerializer.properties.security_group_rules"
	SecurityGroupRules []SecurityGroupNewParamsSecurityGroupSecurityGroupRule `json:"security_group_rules,omitzero"`
	// '#/components/schemas/SingleCreateSecurityGroupSerializer/properties/tags'
	// "$.components.schemas.SingleCreateSecurityGroupSerializer.properties.tags"
	Tags map[string]any `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecurityGroupNewParamsSecurityGroup) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r SecurityGroupNewParamsSecurityGroup) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupNewParamsSecurityGroup
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/SingleCreateSecurityGroupSerializer/properties/security_group_rules/items'
// "$.components.schemas.SingleCreateSecurityGroupSerializer.properties.security_group_rules.items"
type SecurityGroupNewParamsSecurityGroupSecurityGroupRule struct {
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
	Direction string `json:"direction,omitzero"`
	// '#/components/schemas/CreateSecurityGroupRuleSerializer/properties/ethertype'
	// "$.components.schemas.CreateSecurityGroupRuleSerializer.properties.ethertype"
	//
	// Any of "IPv4", "IPv6".
	Ethertype string `json:"ethertype,omitzero"`
	// '#/components/schemas/CreateSecurityGroupRuleSerializer/properties/protocol'
	// "$.components.schemas.CreateSecurityGroupRuleSerializer.properties.protocol"
	//
	// Any of "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap",
	// "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts",
	// "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp".
	Protocol string `json:"protocol,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecurityGroupNewParamsSecurityGroupSecurityGroupRule) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
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
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}/{group_id}'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}/{group_id}'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/UpdateSecurityGroupSerializer/properties/name'
	// "$.components.schemas.UpdateSecurityGroupSerializer.properties.name"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/UpdateSecurityGroupSerializer/properties/changed_rules'
	// "$.components.schemas.UpdateSecurityGroupSerializer.properties.changed_rules"
	ChangedRules []SecurityGroupUpdateParamsChangedRule `json:"changed_rules,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecurityGroupUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SecurityGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/UpdateSecurityGroupSerializer/properties/changed_rules/items'
// "$.components.schemas.UpdateSecurityGroupSerializer.properties.changed_rules.items"
//
// The property Action is required.
type SecurityGroupUpdateParamsChangedRule struct {
	// '#/components/schemas/UpdateSecurityGroupRuleSerializer/properties/action'
	// "$.components.schemas.UpdateSecurityGroupRuleSerializer.properties.action"
	//
	// Any of "create", "delete".
	Action string `json:"action,omitzero,required"`
	// '#/components/schemas/UpdateSecurityGroupRuleSerializer/properties/remote_group_id/anyOf/0'
	// "$.components.schemas.UpdateSecurityGroupRuleSerializer.properties.remote_group_id.anyOf[0]"
	RemoteGroupID param.Opt[string] `json:"remote_group_id,omitzero" format:"uuid4"`
	// '#/components/schemas/UpdateSecurityGroupRuleSerializer/properties/remote_ip_prefix/anyOf/0'
	// "$.components.schemas.UpdateSecurityGroupRuleSerializer.properties.remote_ip_prefix.anyOf[0]"
	RemoteIPPrefix param.Opt[string] `json:"remote_ip_prefix,omitzero" format:"ipvanynetwork"`
	// '#/components/schemas/UpdateSecurityGroupRuleSerializer/properties/security_group_rule_id/anyOf/0'
	// "$.components.schemas.UpdateSecurityGroupRuleSerializer.properties.security_group_rule_id.anyOf[0]"
	SecurityGroupRuleID param.Opt[string] `json:"security_group_rule_id,omitzero" format:"uuid4"`
	// '#/components/schemas/UpdateSecurityGroupRuleSerializer/properties/description'
	// "$.components.schemas.UpdateSecurityGroupRuleSerializer.properties.description"
	Description param.Opt[string] `json:"description,omitzero"`
	// '#/components/schemas/UpdateSecurityGroupRuleSerializer/properties/port_range_max'
	// "$.components.schemas.UpdateSecurityGroupRuleSerializer.properties.port_range_max"
	PortRangeMax param.Opt[int64] `json:"port_range_max,omitzero"`
	// '#/components/schemas/UpdateSecurityGroupRuleSerializer/properties/port_range_min'
	// "$.components.schemas.UpdateSecurityGroupRuleSerializer.properties.port_range_min"
	PortRangeMin param.Opt[int64] `json:"port_range_min,omitzero"`
	// '#/components/schemas/UpdateSecurityGroupRuleSerializer/properties/ethertype/anyOf/0'
	// "$.components.schemas.UpdateSecurityGroupRuleSerializer.properties.ethertype.anyOf[0]"
	//
	// Any of "IPv4", "IPv6".
	Ethertype string `json:"ethertype,omitzero"`
	// '#/components/schemas/UpdateSecurityGroupRuleSerializer/properties/direction'
	// "$.components.schemas.UpdateSecurityGroupRuleSerializer.properties.direction"
	//
	// Any of "egress", "ingress".
	Direction string `json:"direction,omitzero"`
	// '#/components/schemas/UpdateSecurityGroupRuleSerializer/properties/protocol'
	// "$.components.schemas.UpdateSecurityGroupRuleSerializer.properties.protocol"
	//
	// Any of "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap",
	// "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts",
	// "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp".
	Protocol string `json:"protocol,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecurityGroupUpdateParamsChangedRule) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
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
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}'].get.parameters[2]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}'].get.parameters[3]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}'].get.parameters[5]"
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}'].get.parameters[4]"
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecurityGroupListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SecurityGroupListParams]'s query parameters as
// `url.Values`.
func (r SecurityGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SecurityGroupDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}/{group_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}/{group_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecurityGroupDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type SecurityGroupCopyParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D%2Fcopy/post/parameters/0/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}/{group_id}/copy'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D%2Fcopy/post/parameters/1/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}/{group_id}/copy'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/NameSerializerPydantic/properties/name'
	// "$.components.schemas.NameSerializerPydantic.properties.name"
	Name string `json:"name,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecurityGroupCopyParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SecurityGroupCopyParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupCopyParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type SecurityGroupGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}/{group_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}/{group_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecurityGroupGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type SecurityGroupRevertToDefaultParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D%2Frevert/post/parameters/0/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}/{group_id}/revert'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fsecuritygroups%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bgroup_id%7D%2Frevert/post/parameters/1/schema'
	// "$.paths['/cloud/v1/securitygroups/{project_id}/{region_id}/{group_id}/revert'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SecurityGroupRevertToDefaultParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
