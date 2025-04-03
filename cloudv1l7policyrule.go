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

// CloudV1L7policyRuleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1L7policyRuleService] method instead.
type CloudV1L7policyRuleService struct {
	Options []option.RequestOption
}

// NewCloudV1L7policyRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1L7policyRuleService(opts ...option.RequestOption) (r *CloudV1L7policyRuleService) {
	r = &CloudV1L7policyRuleService{}
	r.Options = opts
	return
}

// Create load balancer L7 rule
func (r *CloudV1L7policyRuleService) New(ctx context.Context, projectID int64, regionID int64, l7policyID string, body CloudV1L7policyRuleNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s/rules", projectID, regionID, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get load balancer L7 rule
func (r *CloudV1L7policyRuleService) Get(ctx context.Context, projectID int64, regionID int64, l7policyID string, l7ruleID string, opts ...option.RequestOption) (res *L7Rule, err error) {
	opts = append(r.Options[:], opts...)
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	if l7ruleID == "" {
		err = errors.New("missing required l7rule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s/rules/%s", projectID, regionID, l7policyID, l7ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Replace load balancer L7 rule properties
func (r *CloudV1L7policyRuleService) Update(ctx context.Context, projectID int64, regionID int64, l7policyID string, l7ruleID string, body CloudV1L7policyRuleUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	if l7ruleID == "" {
		err = errors.New("missing required l7rule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s/rules/%s", projectID, regionID, l7policyID, l7ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List load balancer L7 policy rules
func (r *CloudV1L7policyRuleService) List(ctx context.Context, projectID int64, regionID int64, l7policyID string, opts ...option.RequestOption) (res *CloudV1L7policyRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s/rules", projectID, regionID, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete load balancer L7 rule
func (r *CloudV1L7policyRuleService) Delete(ctx context.Context, projectID int64, regionID int64, l7policyID string, l7ruleID string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	if l7ruleID == "" {
		err = errors.New("missing required l7rule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s/rules/%s", projectID, regionID, l7policyID, l7ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// L7Rule schema All the rules associated with a given policy are logically ANDed
// together. A request must match all the policy’s rules to match the policy. If
// you need to express a logical OR operation between rules, then do this by
// creating multiple policies with the same action.
type L7Rule struct {
	// L7Rule ID
	ID string `json:"id"`
	// The comparison type for the L7 rule
	CompareType L7RuleCompareType `json:"compare_type"`
	// When true the logic of the rule is inverted. For example, with invert true,
	// 'equal to' would become 'not equal to'. Default is false.
	Invert bool `json:"invert"`
	// The key to use for the comparison. For example, the name of the cookie to
	// evaluate.
	Key string `json:"key"`
	// L7 policy operating status
	OperatingStatus L7RuleOperatingStatus `json:"operating_status"`
	// Project ID
	ProjectID          int64                    `json:"project_id"`
	ProvisioningStatus L7RuleProvisioningStatus `json:"provisioning_status"`
	// Region name
	Region string `json:"region"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// A list of simple strings assigned to the l7 rule
	Tags []string `json:"tags"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string `json:"task_id"`
	// The L7 rule type
	Type L7RuleType `json:"type"`
	// The value to use for the comparison. For example, the file type to compare.
	Value string     `json:"value"`
	JSON  l7RuleJSON `json:"-"`
}

// l7RuleJSON contains the JSON metadata for the struct [L7Rule]
type l7RuleJSON struct {
	ID                 apijson.Field
	CompareType        apijson.Field
	Invert             apijson.Field
	Key                apijson.Field
	OperatingStatus    apijson.Field
	ProjectID          apijson.Field
	ProvisioningStatus apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Tags               apijson.Field
	TaskID             apijson.Field
	Type               apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *L7Rule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r l7RuleJSON) RawJSON() string {
	return r.raw
}

// The comparison type for the L7 rule
type L7RuleCompareType string

const (
	L7RuleCompareTypeContains   L7RuleCompareType = "CONTAINS"
	L7RuleCompareTypeEndsWith   L7RuleCompareType = "ENDS_WITH"
	L7RuleCompareTypeEqualTo    L7RuleCompareType = "EQUAL_TO"
	L7RuleCompareTypeRegex      L7RuleCompareType = "REGEX"
	L7RuleCompareTypeStartsWith L7RuleCompareType = "STARTS_WITH"
)

func (r L7RuleCompareType) IsKnown() bool {
	switch r {
	case L7RuleCompareTypeContains, L7RuleCompareTypeEndsWith, L7RuleCompareTypeEqualTo, L7RuleCompareTypeRegex, L7RuleCompareTypeStartsWith:
		return true
	}
	return false
}

// L7 policy operating status
type L7RuleOperatingStatus string

const (
	L7RuleOperatingStatusDegraded  L7RuleOperatingStatus = "DEGRADED"
	L7RuleOperatingStatusDraining  L7RuleOperatingStatus = "DRAINING"
	L7RuleOperatingStatusError     L7RuleOperatingStatus = "ERROR"
	L7RuleOperatingStatusNoMonitor L7RuleOperatingStatus = "NO_MONITOR"
	L7RuleOperatingStatusOffline   L7RuleOperatingStatus = "OFFLINE"
	L7RuleOperatingStatusOnline    L7RuleOperatingStatus = "ONLINE"
)

func (r L7RuleOperatingStatus) IsKnown() bool {
	switch r {
	case L7RuleOperatingStatusDegraded, L7RuleOperatingStatusDraining, L7RuleOperatingStatusError, L7RuleOperatingStatusNoMonitor, L7RuleOperatingStatusOffline, L7RuleOperatingStatusOnline:
		return true
	}
	return false
}

type L7RuleProvisioningStatus string

const (
	L7RuleProvisioningStatusActive        L7RuleProvisioningStatus = "ACTIVE"
	L7RuleProvisioningStatusDeleted       L7RuleProvisioningStatus = "DELETED"
	L7RuleProvisioningStatusError         L7RuleProvisioningStatus = "ERROR"
	L7RuleProvisioningStatusPendingCreate L7RuleProvisioningStatus = "PENDING_CREATE"
	L7RuleProvisioningStatusPendingDelete L7RuleProvisioningStatus = "PENDING_DELETE"
	L7RuleProvisioningStatusPendingUpdate L7RuleProvisioningStatus = "PENDING_UPDATE"
)

func (r L7RuleProvisioningStatus) IsKnown() bool {
	switch r {
	case L7RuleProvisioningStatusActive, L7RuleProvisioningStatusDeleted, L7RuleProvisioningStatusError, L7RuleProvisioningStatusPendingCreate, L7RuleProvisioningStatusPendingDelete, L7RuleProvisioningStatusPendingUpdate:
		return true
	}
	return false
}

// The L7 rule type
type L7RuleType string

const (
	L7RuleTypeCookie          L7RuleType = "COOKIE"
	L7RuleTypeFileType        L7RuleType = "FILE_TYPE"
	L7RuleTypeHeader          L7RuleType = "HEADER"
	L7RuleTypeHostName        L7RuleType = "HOST_NAME"
	L7RuleTypePath            L7RuleType = "PATH"
	L7RuleTypeSslConnHasCert  L7RuleType = "SSL_CONN_HAS_CERT"
	L7RuleTypeSslDnField      L7RuleType = "SSL_DN_FIELD"
	L7RuleTypeSslVerifyResult L7RuleType = "SSL_VERIFY_RESULT"
)

func (r L7RuleType) IsKnown() bool {
	switch r {
	case L7RuleTypeCookie, L7RuleTypeFileType, L7RuleTypeHeader, L7RuleTypeHostName, L7RuleTypePath, L7RuleTypeSslConnHasCert, L7RuleTypeSslDnField, L7RuleTypeSslVerifyResult:
		return true
	}
	return false
}

type CloudV1L7policyRuleListResponse struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []L7Rule                            `json:"results"`
	JSON    cloudV1L7policyRuleListResponseJSON `json:"-"`
}

// cloudV1L7policyRuleListResponseJSON contains the JSON metadata for the struct
// [CloudV1L7policyRuleListResponse]
type cloudV1L7policyRuleListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1L7policyRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1L7policyRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1L7policyRuleNewParams struct {
	// The comparison type for the L7 rule
	CompareType param.Field[CloudV1L7policyRuleNewParamsCompareType] `json:"compare_type,required"`
	// The L7 rule type
	Type param.Field[CloudV1L7policyRuleNewParamsType] `json:"type,required"`
	// The value to use for the comparison. For example, the file type to compare
	Value param.Field[string] `json:"value,required"`
	// When true the logic of the rule is inverted. For example, with invert true,
	// 'equal to' would become 'not equal to'. Default is false.
	Invert param.Field[bool] `json:"invert"`
	// The key to use for the comparison. For example, the name of the cookie to
	// evaluate.
	Key param.Field[string] `json:"key"`
	// A list of simple strings assigned to the l7 rule
	Tags param.Field[[]string] `json:"tags"`
}

func (r CloudV1L7policyRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The comparison type for the L7 rule
type CloudV1L7policyRuleNewParamsCompareType string

const (
	CloudV1L7policyRuleNewParamsCompareTypeContains   CloudV1L7policyRuleNewParamsCompareType = "CONTAINS"
	CloudV1L7policyRuleNewParamsCompareTypeEndsWith   CloudV1L7policyRuleNewParamsCompareType = "ENDS_WITH"
	CloudV1L7policyRuleNewParamsCompareTypeEqualTo    CloudV1L7policyRuleNewParamsCompareType = "EQUAL_TO"
	CloudV1L7policyRuleNewParamsCompareTypeRegex      CloudV1L7policyRuleNewParamsCompareType = "REGEX"
	CloudV1L7policyRuleNewParamsCompareTypeStartsWith CloudV1L7policyRuleNewParamsCompareType = "STARTS_WITH"
)

func (r CloudV1L7policyRuleNewParamsCompareType) IsKnown() bool {
	switch r {
	case CloudV1L7policyRuleNewParamsCompareTypeContains, CloudV1L7policyRuleNewParamsCompareTypeEndsWith, CloudV1L7policyRuleNewParamsCompareTypeEqualTo, CloudV1L7policyRuleNewParamsCompareTypeRegex, CloudV1L7policyRuleNewParamsCompareTypeStartsWith:
		return true
	}
	return false
}

// The L7 rule type
type CloudV1L7policyRuleNewParamsType string

const (
	CloudV1L7policyRuleNewParamsTypeCookie          CloudV1L7policyRuleNewParamsType = "COOKIE"
	CloudV1L7policyRuleNewParamsTypeFileType        CloudV1L7policyRuleNewParamsType = "FILE_TYPE"
	CloudV1L7policyRuleNewParamsTypeHeader          CloudV1L7policyRuleNewParamsType = "HEADER"
	CloudV1L7policyRuleNewParamsTypeHostName        CloudV1L7policyRuleNewParamsType = "HOST_NAME"
	CloudV1L7policyRuleNewParamsTypePath            CloudV1L7policyRuleNewParamsType = "PATH"
	CloudV1L7policyRuleNewParamsTypeSslConnHasCert  CloudV1L7policyRuleNewParamsType = "SSL_CONN_HAS_CERT"
	CloudV1L7policyRuleNewParamsTypeSslDnField      CloudV1L7policyRuleNewParamsType = "SSL_DN_FIELD"
	CloudV1L7policyRuleNewParamsTypeSslVerifyResult CloudV1L7policyRuleNewParamsType = "SSL_VERIFY_RESULT"
)

func (r CloudV1L7policyRuleNewParamsType) IsKnown() bool {
	switch r {
	case CloudV1L7policyRuleNewParamsTypeCookie, CloudV1L7policyRuleNewParamsTypeFileType, CloudV1L7policyRuleNewParamsTypeHeader, CloudV1L7policyRuleNewParamsTypeHostName, CloudV1L7policyRuleNewParamsTypePath, CloudV1L7policyRuleNewParamsTypeSslConnHasCert, CloudV1L7policyRuleNewParamsTypeSslDnField, CloudV1L7policyRuleNewParamsTypeSslVerifyResult:
		return true
	}
	return false
}

type CloudV1L7policyRuleUpdateParams struct {
	// The comparison type for the L7 rule
	CompareType param.Field[CloudV1L7policyRuleUpdateParamsCompareType] `json:"compare_type"`
	// When true the logic of the rule is inverted. For example, with invert true,
	// 'equal to' would become 'not equal to'. Default is false.
	Invert param.Field[bool] `json:"invert"`
	// The key to use for the comparison. For example, the name of the cookie to
	// evaluate.
	Key param.Field[string] `json:"key"`
	// A list of simple strings assigned to the l7 rule
	Tags param.Field[[]string] `json:"tags"`
	// The L7 rule type
	Type param.Field[CloudV1L7policyRuleUpdateParamsType] `json:"type"`
	// The value to use for the comparison. For example, the file type to compare
	Value param.Field[string] `json:"value"`
}

func (r CloudV1L7policyRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The comparison type for the L7 rule
type CloudV1L7policyRuleUpdateParamsCompareType string

const (
	CloudV1L7policyRuleUpdateParamsCompareTypeContains   CloudV1L7policyRuleUpdateParamsCompareType = "CONTAINS"
	CloudV1L7policyRuleUpdateParamsCompareTypeEndsWith   CloudV1L7policyRuleUpdateParamsCompareType = "ENDS_WITH"
	CloudV1L7policyRuleUpdateParamsCompareTypeEqualTo    CloudV1L7policyRuleUpdateParamsCompareType = "EQUAL_TO"
	CloudV1L7policyRuleUpdateParamsCompareTypeRegex      CloudV1L7policyRuleUpdateParamsCompareType = "REGEX"
	CloudV1L7policyRuleUpdateParamsCompareTypeStartsWith CloudV1L7policyRuleUpdateParamsCompareType = "STARTS_WITH"
)

func (r CloudV1L7policyRuleUpdateParamsCompareType) IsKnown() bool {
	switch r {
	case CloudV1L7policyRuleUpdateParamsCompareTypeContains, CloudV1L7policyRuleUpdateParamsCompareTypeEndsWith, CloudV1L7policyRuleUpdateParamsCompareTypeEqualTo, CloudV1L7policyRuleUpdateParamsCompareTypeRegex, CloudV1L7policyRuleUpdateParamsCompareTypeStartsWith:
		return true
	}
	return false
}

// The L7 rule type
type CloudV1L7policyRuleUpdateParamsType string

const (
	CloudV1L7policyRuleUpdateParamsTypeCookie          CloudV1L7policyRuleUpdateParamsType = "COOKIE"
	CloudV1L7policyRuleUpdateParamsTypeFileType        CloudV1L7policyRuleUpdateParamsType = "FILE_TYPE"
	CloudV1L7policyRuleUpdateParamsTypeHeader          CloudV1L7policyRuleUpdateParamsType = "HEADER"
	CloudV1L7policyRuleUpdateParamsTypeHostName        CloudV1L7policyRuleUpdateParamsType = "HOST_NAME"
	CloudV1L7policyRuleUpdateParamsTypePath            CloudV1L7policyRuleUpdateParamsType = "PATH"
	CloudV1L7policyRuleUpdateParamsTypeSslConnHasCert  CloudV1L7policyRuleUpdateParamsType = "SSL_CONN_HAS_CERT"
	CloudV1L7policyRuleUpdateParamsTypeSslDnField      CloudV1L7policyRuleUpdateParamsType = "SSL_DN_FIELD"
	CloudV1L7policyRuleUpdateParamsTypeSslVerifyResult CloudV1L7policyRuleUpdateParamsType = "SSL_VERIFY_RESULT"
)

func (r CloudV1L7policyRuleUpdateParamsType) IsKnown() bool {
	switch r {
	case CloudV1L7policyRuleUpdateParamsTypeCookie, CloudV1L7policyRuleUpdateParamsTypeFileType, CloudV1L7policyRuleUpdateParamsTypeHeader, CloudV1L7policyRuleUpdateParamsTypeHostName, CloudV1L7policyRuleUpdateParamsTypePath, CloudV1L7policyRuleUpdateParamsTypeSslConnHasCert, CloudV1L7policyRuleUpdateParamsTypeSslDnField, CloudV1L7policyRuleUpdateParamsTypeSslVerifyResult:
		return true
	}
	return false
}
