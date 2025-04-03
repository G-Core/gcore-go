// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1SecuritygroupService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1SecuritygroupService] method instead.
type CloudV1SecuritygroupService struct {
	Options      []option.RequestOption
	MetadataItem *CloudV1SecuritygroupMetadataItemService
	Metadata     *CloudV1SecuritygroupMetadataService
}

// NewCloudV1SecuritygroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1SecuritygroupService(opts ...option.RequestOption) (r *CloudV1SecuritygroupService) {
	r = &CloudV1SecuritygroupService{}
	r.Options = opts
	r.MetadataItem = NewCloudV1SecuritygroupMetadataItemService(opts...)
	r.Metadata = NewCloudV1SecuritygroupMetadataService(opts...)
	return
}

// Create security group
func (r *CloudV1SecuritygroupService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1SecuritygroupNewParams, opts ...option.RequestOption) (res *SecurityGroup, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get security group
func (r *CloudV1SecuritygroupService) Get(ctx context.Context, projectID int64, regionID int64, pk string, opts ...option.RequestOption) (res *SecurityGroup, err error) {
	opts = append(r.Options[:], opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change security group
func (r *CloudV1SecuritygroupService) Update(ctx context.Context, projectID int64, regionID int64, pk string, body CloudV1SecuritygroupUpdateParams, opts ...option.RequestOption) (res *SecurityGroup, err error) {
	opts = append(r.Options[:], opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get security groups
func (r *CloudV1SecuritygroupService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1SecuritygroupListParams, opts ...option.RequestOption) (res *CloudV1SecuritygroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete security group
func (r *CloudV1SecuritygroupService) Delete(ctx context.Context, projectID int64, regionID int64, pk string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Add new rule to security group
func (r *CloudV1SecuritygroupService) AddRule(ctx context.Context, projectID int64, regionID int64, pk string, body CloudV1SecuritygroupAddRuleParams, opts ...option.RequestOption) (res *SecurityGroupRule, err error) {
	opts = append(r.Options[:], opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s/rules", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a deep copy of security group
func (r *CloudV1SecuritygroupService) Copy(ctx context.Context, projectID int64, regionID int64, pk string, body CloudV1SecuritygroupCopyParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s/copy", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve the instances list (filter by security group)
func (r *CloudV1SecuritygroupService) FilterInstances(ctx context.Context, projectID int64, regionID int64, secgroupID string, opts ...option.RequestOption) (res *DeprecatedInstanceList, err error) {
	opts = append(r.Options[:], opts...)
	if secgroupID == "" {
		err = errors.New("missing required secgroup_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s/instances", projectID, regionID, secgroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Revert security group
func (r *CloudV1SecuritygroupService) Revert(ctx context.Context, projectID int64, regionID int64, pk string, opts ...option.RequestOption) (res *SecurityGroup, err error) {
	opts = append(r.Options[:], opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s/revert", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type CreateSecurityGroupRuleParam struct {
	// Rule description
	Description param.Field[string] `json:"description"`
	// Ingress or egress, which is the direction in which the security group is applied
	Direction param.Field[CreateSecurityGroupRuleDirection] `json:"direction"`
	// Ether type
	Ethertype param.Field[CreateSecurityGroupRuleEthertype] `json:"ethertype"`
	// The maximum port number in the range that is matched by the security group rule
	PortRangeMax param.Field[int64] `json:"port_range_max"`
	// The minimum port number in the range that is matched by the security group rule
	PortRangeMin param.Field[int64] `json:"port_range_min"`
	// Protocol
	Protocol param.Field[SecurityGroupProtocol] `json:"protocol"`
	// The remote group UUID to associate with this security group
	RemoteGroupID param.Field[string] `json:"remote_group_id" format:"uuid4"`
	// The remote IP prefix that is matched by this security group rule
	RemoteIPPrefix param.Field[string] `json:"remote_ip_prefix" format:"ipvanynetwork"`
}

func (r CreateSecurityGroupRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Ingress or egress, which is the direction in which the security group is applied
type CreateSecurityGroupRuleDirection string

const (
	CreateSecurityGroupRuleDirectionEgress  CreateSecurityGroupRuleDirection = "egress"
	CreateSecurityGroupRuleDirectionIngress CreateSecurityGroupRuleDirection = "ingress"
)

func (r CreateSecurityGroupRuleDirection) IsKnown() bool {
	switch r {
	case CreateSecurityGroupRuleDirectionEgress, CreateSecurityGroupRuleDirectionIngress:
		return true
	}
	return false
}

// Ether type
type CreateSecurityGroupRuleEthertype string

const (
	CreateSecurityGroupRuleEthertypeIPv4 CreateSecurityGroupRuleEthertype = "IPv4"
	CreateSecurityGroupRuleEthertypeIPv6 CreateSecurityGroupRuleEthertype = "IPv6"
)

func (r CreateSecurityGroupRuleEthertype) IsKnown() bool {
	switch r {
	case CreateSecurityGroupRuleEthertypeIPv4, CreateSecurityGroupRuleEthertypeIPv6:
		return true
	}
	return false
}

type NamePydantic struct {
	// Name.
	Name string           `json:"name,required"`
	JSON namePydanticJSON `json:"-"`
}

// namePydanticJSON contains the JSON metadata for the struct [NamePydantic]
type namePydanticJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamePydantic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namePydanticJSON) RawJSON() string {
	return r.raw
}

type NamePydanticParam struct {
	// Name.
	Name param.Field[string] `json:"name,required"`
}

func (r NamePydanticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// Datetime when the security group was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Security group description
	Description string `json:"description,nullable"`
	// Metadata items for a security group
	Metadata []DetailedMetadata `json:"metadata"`
	// Security group rules
	SecurityGroupRules []SecurityGroupRule `json:"security_group_rules"`
	Tags               []string            `json:"tags,nullable"`
	JSON               securityGroupJSON   `json:"-"`
}

// securityGroupJSON contains the JSON metadata for the struct [SecurityGroup]
type securityGroupJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	Name               apijson.Field
	ProjectID          apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	RevisionNumber     apijson.Field
	UpdatedAt          apijson.Field
	Description        apijson.Field
	Metadata           apijson.Field
	SecurityGroupRules apijson.Field
	Tags               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SecurityGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityGroupJSON) RawJSON() string {
	return r.raw
}

type CloudV1SecuritygroupListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []SecurityGroup                      `json:"results,required"`
	JSON    cloudV1SecuritygroupListResponseJSON `json:"-"`
}

// cloudV1SecuritygroupListResponseJSON contains the JSON metadata for the struct
// [CloudV1SecuritygroupListResponse]
type cloudV1SecuritygroupListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1SecuritygroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1SecuritygroupListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1SecuritygroupNewParams struct {
	// Security group
	SecurityGroup param.Field[CloudV1SecuritygroupNewParamsSecurityGroup] `json:"security_group,required"`
	// List of instances
	Instances param.Field[[]string] `json:"instances"`
}

func (r CloudV1SecuritygroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Security group
type CloudV1SecuritygroupNewParamsSecurityGroup struct {
	// Security group name
	Name param.Field[string] `json:"name,required"`
	// Security group description
	Description param.Field[string] `json:"description"`
	// Create one or more metadata items for a security group
	Metadata param.Field[RawMetadataParam] `json:"metadata"`
	// Security group rules
	SecurityGroupRules param.Field[[]CreateSecurityGroupRuleParam] `json:"security_group_rules"`
	Tags               param.Field[[]string]                       `json:"tags"`
}

func (r CloudV1SecuritygroupNewParamsSecurityGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1SecuritygroupUpdateParams struct {
	// List of rules to create or delete
	ChangedRules param.Field[[]CloudV1SecuritygroupUpdateParamsChangedRule] `json:"changed_rules"`
	// Name
	Name param.Field[string] `json:"name"`
}

func (r CloudV1SecuritygroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1SecuritygroupUpdateParamsChangedRule struct {
	// Action for a rule
	Action param.Field[CloudV1SecuritygroupUpdateParamsChangedRulesAction] `json:"action,required"`
	// Security grpup rule description
	Description param.Field[string] `json:"description"`
	// Ingress or egress, which is the direction in which the security group rule is
	// applied
	Direction param.Field[CloudV1SecuritygroupUpdateParamsChangedRulesDirection] `json:"direction"`
	// Must be IPv4 or IPv6, and addresses represented in CIDR must match the ingress
	// or egress rules.
	Ethertype param.Field[CloudV1SecuritygroupUpdateParamsChangedRulesEthertype] `json:"ethertype"`
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
	// UUID of rule to be deleted. Required for action 'delete' only
	SecurityGroupRuleID param.Field[string] `json:"security_group_rule_id" format:"uuid4"`
}

func (r CloudV1SecuritygroupUpdateParamsChangedRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action for a rule
type CloudV1SecuritygroupUpdateParamsChangedRulesAction string

const (
	CloudV1SecuritygroupUpdateParamsChangedRulesActionCreate CloudV1SecuritygroupUpdateParamsChangedRulesAction = "create"
	CloudV1SecuritygroupUpdateParamsChangedRulesActionDelete CloudV1SecuritygroupUpdateParamsChangedRulesAction = "delete"
)

func (r CloudV1SecuritygroupUpdateParamsChangedRulesAction) IsKnown() bool {
	switch r {
	case CloudV1SecuritygroupUpdateParamsChangedRulesActionCreate, CloudV1SecuritygroupUpdateParamsChangedRulesActionDelete:
		return true
	}
	return false
}

// Ingress or egress, which is the direction in which the security group rule is
// applied
type CloudV1SecuritygroupUpdateParamsChangedRulesDirection string

const (
	CloudV1SecuritygroupUpdateParamsChangedRulesDirectionEgress  CloudV1SecuritygroupUpdateParamsChangedRulesDirection = "egress"
	CloudV1SecuritygroupUpdateParamsChangedRulesDirectionIngress CloudV1SecuritygroupUpdateParamsChangedRulesDirection = "ingress"
)

func (r CloudV1SecuritygroupUpdateParamsChangedRulesDirection) IsKnown() bool {
	switch r {
	case CloudV1SecuritygroupUpdateParamsChangedRulesDirectionEgress, CloudV1SecuritygroupUpdateParamsChangedRulesDirectionIngress:
		return true
	}
	return false
}

// Must be IPv4 or IPv6, and addresses represented in CIDR must match the ingress
// or egress rules.
type CloudV1SecuritygroupUpdateParamsChangedRulesEthertype string

const (
	CloudV1SecuritygroupUpdateParamsChangedRulesEthertypeIPv4 CloudV1SecuritygroupUpdateParamsChangedRulesEthertype = "IPv4"
	CloudV1SecuritygroupUpdateParamsChangedRulesEthertypeIPv6 CloudV1SecuritygroupUpdateParamsChangedRulesEthertype = "IPv6"
)

func (r CloudV1SecuritygroupUpdateParamsChangedRulesEthertype) IsKnown() bool {
	switch r {
	case CloudV1SecuritygroupUpdateParamsChangedRulesEthertypeIPv4, CloudV1SecuritygroupUpdateParamsChangedRulesEthertypeIPv6:
		return true
	}
	return false
}

type CloudV1SecuritygroupListParams struct {
	// Limit the number of returned limit request entities.
	Limit param.Field[int64] `query:"limit"`
	// Filter by metadata keys. Must be a valid JSON string. curl -G --data-urlencode
	// "metadata_k=["value", "sense"]" --url
	// "http://localhost:1111/v1/securitygroups/1/1"
	MetadataK param.Field[string] `query:"metadata_k"`
	// Filter by metadata key-value pairs. Must be a valid JSON string. curl -G
	// --data-urlencode "metadata_kv={"key": "value"}" --url
	// "http://localhost:1111/v1/securitygroups/1/1"
	MetadataKv param.Field[string] `query:"metadata_kv"`
	// Offset value is used to exclude the first set of records from the result.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV1SecuritygroupListParams]'s query parameters as
// `url.Values`.
func (r CloudV1SecuritygroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1SecuritygroupAddRuleParams struct {
	CreateSecurityGroupRule CreateSecurityGroupRuleParam `json:"create_security_group_rule"`
}

func (r CloudV1SecuritygroupAddRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateSecurityGroupRule)
}

type CloudV1SecuritygroupCopyParams struct {
	NamePydantic NamePydanticParam `json:"name_pydantic,required"`
}

func (r CloudV1SecuritygroupCopyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.NamePydantic)
}
