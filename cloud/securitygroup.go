// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
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
	tasks   TaskService
}

// NewSecurityGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecurityGroupService(opts ...option.RequestOption) (r SecurityGroupService) {
	r = SecurityGroupService{}
	r.Options = opts
	r.Rules = NewSecurityGroupRuleService(opts...)
	r.tasks = NewTaskService(opts...)
	return
}

// Creates a new security group with the specified configuration. If no egress
// rules are provided, default set of egress rules will be applied If rules are
// explicitly set to empty, no rules will be created.
func (r *SecurityGroupService) New(ctx context.Context, params SecurityGroupNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v2/security_groups/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// NewAndPoll creates a security group and polls for completion of the first task. Use the [TaskService.Poll] method if
// you need to poll for all tasks.
func (r *SecurityGroupService) NewAndPoll(ctx context.Context, params SecurityGroupNewParams, opts ...option.RequestOption) (res *SecurityGroup, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams SecurityGroupGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]
	task, err := r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.SecurityGroups) != 1 {
		return nil, errors.New("expected exactly one security group to be created in a task")
	}
	resourceID := task.CreatedResources.SecurityGroups[0]

	return r.Get(ctx, resourceID, getParams, opts...)
}

// Updates the specified security group with the provided changes.
//
// **Behavior:**
//
//   - Simple fields (name, description) will be updated if provided
//   - Undefined fields will remain unchanged
//   - If no change is detected for a specific field compared to the current security
//     group state, that field will be skipped
//   - If no changes are detected at all across all fields, no task will be created
//     and an empty task list will be returned
//
// **Important - Security Group Rules:**
//
//   - Rules must be specified completely as the desired final state
//   - The system compares the provided rules against the current state
//   - Rules that exist in the request but not in the current state will be added
//   - Rules that exist in the current state but not in the request will be removed
//   - To keep existing rules, they must be included in the request alongside any new
//     rules
func (r *SecurityGroupService) Update(ctx context.Context, groupID string, params SecurityGroupUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v2/security_groups/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// UpdateAndPoll updates a security group and polls for completion of the first task. Use the [TaskService.Poll] method if
// you need to poll for all tasks.
func (r *SecurityGroupService) UpdateAndPoll(ctx context.Context, groupID string, params SecurityGroupUpdateParams, opts ...option.RequestOption) (res *SecurityGroup, err error) {
	resource, err := r.Update(ctx, groupID, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams SecurityGroupGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) == 0 {
		return nil, errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, groupID, getParams, opts...)
}

// List all security groups in the specified project and region.
func (r *SecurityGroupService) List(ctx context.Context, params SecurityGroupListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SecurityGroup], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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

// List all security groups in the specified project and region.
func (r *SecurityGroupService) ListAutoPaging(ctx context.Context, params SecurityGroupListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SecurityGroup] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete a specific security group and all its associated rules.
func (r *SecurityGroupService) Delete(ctx context.Context, groupID string, body SecurityGroupDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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

// Create a deep copy of an existing security group.
func (r *SecurityGroupService) Copy(ctx context.Context, groupID string, params SecurityGroupCopyParams, opts ...option.RequestOption) (res *SecurityGroup, err error) {
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
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s/copy", params.ProjectID.Value, params.RegionID.Value, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get detailed information about a specific security group.
func (r *SecurityGroupService) Get(ctx context.Context, groupID string, query SecurityGroupGetParams, opts ...option.RequestOption) (res *SecurityGroup, err error) {
	opts = slices.Concat(r.Options, opts)
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

// Revert a security group to its previous state.
func (r *SecurityGroupService) RevertToDefault(ctx context.Context, groupID string, body SecurityGroupRevertToDefaultParams, opts ...option.RequestOption) (res *SecurityGroup, err error) {
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
	ID string `json:"id" api:"required" format:"uuid4"`
	// Datetime when the security group was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Security group name
	Name string `json:"name" api:"required"`
	// Project ID
	ProjectID int64 `json:"project_id" api:"required"`
	// Region name
	Region string `json:"region" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// The number of revisions
	RevisionNumber int64 `json:"revision_number" api:"required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	TagsV2 []Tag `json:"tags_v2" api:"required"`
	// Datetime when the security group was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Security group description
	Description string `json:"description" api:"nullable"`
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
	ID string `json:"id" api:"required" format:"uuid4"`
	// Datetime when the rule was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Rule description
	Description string `json:"description" api:"required"`
	// Ingress or egress, which is the direction in which the security group rule is
	// applied
	//
	// Any of "egress", "ingress".
	Direction SecurityGroupRuleDirection `json:"direction" api:"required"`
	// Must be IPv4 or IPv6, and addresses represented in CIDR must match the ingress
	// or egress rules.
	//
	// Any of "IPv4", "IPv6".
	Ethertype SecurityGroupRuleEthertype `json:"ethertype" api:"required"`
	// The maximum port number in the range that is matched by the security group rule
	PortRangeMax int64 `json:"port_range_max" api:"required"`
	// The minimum port number in the range that is matched by the security group rule
	PortRangeMin int64 `json:"port_range_min" api:"required"`
	// Protocol
	//
	// Any of "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap",
	// "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts",
	// "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp".
	Protocol SecurityGroupRuleProtocol `json:"protocol" api:"required"`
	// The remote group UUID to associate with this security group rule
	RemoteGroupID string `json:"remote_group_id" api:"required" format:"uuid4"`
	// The remote IP prefix that is matched by this security group rule
	RemoteIPPrefix string `json:"remote_ip_prefix" api:"required" format:"ipvanynetwork"`
	// The revision number of the resource
	RevisionNumber int64 `json:"revision_number" api:"required"`
	// The security group ID to associate with this security group rule
	SecurityGroupID string `json:"security_group_id" api:"required" format:"uuid4"`
	// Datetime when the rule was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Description     respjson.Field
		Direction       respjson.Field
		Ethertype       respjson.Field
		PortRangeMax    respjson.Field
		PortRangeMin    respjson.Field
		Protocol        respjson.Field
		RemoteGroupID   respjson.Field
		RemoteIPPrefix  respjson.Field
		RevisionNumber  respjson.Field
		SecurityGroupID respjson.Field
		UpdatedAt       respjson.Field
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
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Security group name
	Name string `json:"name" api:"required"`
	// Security group description
	Description param.Opt[string] `json:"description,omitzero"`
	// Security group rules
	Rules []SecurityGroupNewParamsRule `json:"rules,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Both tag keys and values have a maximum
	// length of 255 characters. Some tags are read-only and cannot be modified by the
	// user. Tags are also integrated with cost reports, allowing cost data to be
	// filtered based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r SecurityGroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecurityGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Direction is required.
type SecurityGroupNewParamsRule struct {
	// Ingress or egress, which is the direction in which the security group is applied
	//
	// Any of "egress", "ingress".
	Direction string `json:"direction,omitzero" api:"required"`
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
	Ethertype string `json:"ethertype,omitzero"`
	// Protocol
	//
	// Any of "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap",
	// "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts",
	// "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp".
	Protocol string `json:"protocol,omitzero"`
	paramObj
}

func (r SecurityGroupNewParamsRule) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupNewParamsRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecurityGroupNewParamsRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SecurityGroupNewParamsRule](
		"direction", "egress", "ingress",
	)
	apijson.RegisterFieldValidator[SecurityGroupNewParamsRule](
		"ethertype", "IPv4", "IPv6",
	)
	apijson.RegisterFieldValidator[SecurityGroupNewParamsRule](
		"protocol", "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap", "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts", "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp",
	)
}

type SecurityGroupUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Security group description
	Description param.Opt[string] `json:"description,omitzero"`
	// Name
	Name param.Opt[string] `json:"name,omitzero"`
	// Security group rules
	Rules []SecurityGroupUpdateParamsRule `json:"rules,omitzero"`
	// Update key-value tags using JSON Merge Patch semantics (RFC 7386). Provide
	// key-value pairs to add or update tags. Set tag values to `null` to remove tags.
	// Unspecified tags remain unchanged. Read-only tags are always preserved and
	// cannot be modified.
	//
	// **Examples:**
	//
	//   - **Add/update tags:**
	//     `{'tags': {'environment': 'production', 'team': 'backend'}}` adds new tags or
	//     updates existing ones.
	//   - **Delete tags:** `{'tags': {'old_tag': null}}` removes specific tags.
	//   - **Remove all tags:** `{'tags': null}` removes all user-managed tags (read-only
	//     tags are preserved).
	//   - **Partial update:** `{'tags': {'environment': 'staging'}}` only updates
	//     specified tags.
	//   - **Mixed operations:**
	//     `{'tags': {'environment': 'production', 'cost_center': 'engineering', 'deprecated_tag': null}}`
	//     adds/updates 'environment' and 'cost_center' while removing 'deprecated_tag',
	//     preserving other existing tags.
	//   - **Replace all:** first delete existing tags with null values, then add new
	//     ones in the same request.
	Tags TagUpdateMap `json:"tags,omitzero"`
	paramObj
}

func (r SecurityGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecurityGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityGroupUpdateParamsRule struct {
	// Security grpup rule description
	Description param.Opt[string] `json:"description,omitzero"`
	// The maximum port number in the range that is matched by the security group rule
	PortRangeMax param.Opt[int64] `json:"port_range_max,omitzero"`
	// The minimum port number in the range that is matched by the security group rule
	PortRangeMin param.Opt[int64] `json:"port_range_min,omitzero"`
	// The remote group UUID to associate with this security group rule
	RemoteGroupID param.Opt[string] `json:"remote_group_id,omitzero" format:"uuid4"`
	// The remote IP prefix that is matched by this security group rule
	RemoteIPPrefix param.Opt[string] `json:"remote_ip_prefix,omitzero" format:"ipvanynetwork"`
	// Ingress or egress, which is the direction in which the security group rule is
	// applied
	//
	// Any of "egress", "ingress".
	Direction string `json:"direction,omitzero"`
	// Must be IPv4 or IPv6, and addresses represented in CIDR must match the ingress
	// or egress rules.
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

func (r SecurityGroupUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupUpdateParamsRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecurityGroupUpdateParamsRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SecurityGroupUpdateParamsRule](
		"direction", "egress", "ingress",
	)
	apijson.RegisterFieldValidator[SecurityGroupUpdateParamsRule](
		"ethertype", "IPv4", "IPv6",
	)
	apijson.RegisterFieldValidator[SecurityGroupUpdateParamsRule](
		"protocol", "ah", "any", "dccp", "egp", "esp", "gre", "icmp", "igmp", "ipencap", "ipip", "ipv6-encap", "ipv6-frag", "ipv6-icmp", "ipv6-nonxt", "ipv6-opts", "ipv6-route", "ospf", "pgm", "rsvp", "sctp", "tcp", "udp", "udplite", "vrrp",
	)
}

type SecurityGroupListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Limit of items on a single page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Filter by name. Must be specified a full name of the security group.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Offset in results list
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Optional. Filter by tag key-value pairs.
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// Optional. Filter by tag keys. ?`tag_key`=key1&`tag_key`=key2
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
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type SecurityGroupCopyParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Name.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r SecurityGroupCopyParams) MarshalJSON() (data []byte, err error) {
	type shadow SecurityGroupCopyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecurityGroupCopyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecurityGroupGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}

type SecurityGroupRevertToDefaultParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	paramObj
}
