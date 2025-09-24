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

// LoadBalancerL7PolicyRuleService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoadBalancerL7PolicyRuleService] method instead.
type LoadBalancerL7PolicyRuleService struct {
	Options []option.RequestOption
}

// NewLoadBalancerL7PolicyRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLoadBalancerL7PolicyRuleService(opts ...option.RequestOption) (r LoadBalancerL7PolicyRuleService) {
	r = LoadBalancerL7PolicyRuleService{}
	r.Options = opts
	return
}

// Create load balancer L7 rule
func (r *LoadBalancerL7PolicyRuleService) New(ctx context.Context, l7policyID string, params LoadBalancerL7PolicyRuleNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s/rules", params.ProjectID.Value, params.RegionID.Value, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List load balancer L7 policy rules
func (r *LoadBalancerL7PolicyRuleService) List(ctx context.Context, l7policyID string, query LoadBalancerL7PolicyRuleListParams, opts ...option.RequestOption) (res *LoadBalancerL7RuleList, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
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
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s/rules", query.ProjectID.Value, query.RegionID.Value, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete load balancer L7 rule
func (r *LoadBalancerL7PolicyRuleService) Delete(ctx context.Context, l7ruleID string, body LoadBalancerL7PolicyRuleDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	if body.L7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	if l7ruleID == "" {
		err = errors.New("missing required l7rule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s/rules/%s", body.ProjectID.Value, body.RegionID.Value, body.L7policyID, l7ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get load balancer L7 rule
func (r *LoadBalancerL7PolicyRuleService) Get(ctx context.Context, l7ruleID string, query LoadBalancerL7PolicyRuleGetParams, opts ...option.RequestOption) (res *LoadBalancerL7Rule, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
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
	if query.L7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	if l7ruleID == "" {
		err = errors.New("missing required l7rule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s/rules/%s", query.ProjectID.Value, query.RegionID.Value, query.L7policyID, l7ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Replace load balancer L7 rule properties
func (r *LoadBalancerL7PolicyRuleService) Replace(ctx context.Context, l7ruleID string, params LoadBalancerL7PolicyRuleReplaceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if params.L7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	if l7ruleID == "" {
		err = errors.New("missing required l7rule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s/rules/%s", params.ProjectID.Value, params.RegionID.Value, params.L7policyID, l7ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type LoadBalancerL7PolicyRuleNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// The comparison type for the L7 rule
	//
	// Any of "CONTAINS", "ENDS_WITH", "EQUAL_TO", "REGEX", "STARTS_WITH".
	CompareType LoadBalancerL7PolicyRuleNewParamsCompareType `json:"compare_type,omitzero,required"`
	// The L7 rule type
	//
	// Any of "COOKIE", "FILE_TYPE", "HEADER", "HOST_NAME", "PATH",
	// "SSL_CONN_HAS_CERT", "SSL_DN_FIELD", "SSL_VERIFY_RESULT".
	Type LoadBalancerL7PolicyRuleNewParamsType `json:"type,omitzero,required"`
	// The value to use for the comparison. For example, the file type to compare
	Value string `json:"value,required"`
	// When true the logic of the rule is inverted. For example, with invert true,
	// 'equal to' would become 'not equal to'. Default is false.
	Invert param.Opt[bool] `json:"invert,omitzero"`
	// The key to use for the comparison. For example, the name of the cookie to
	// evaluate.
	Key param.Opt[string] `json:"key,omitzero"`
	// A list of simple strings assigned to the l7 rule
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r LoadBalancerL7PolicyRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyRuleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The comparison type for the L7 rule
type LoadBalancerL7PolicyRuleNewParamsCompareType string

const (
	LoadBalancerL7PolicyRuleNewParamsCompareTypeContains   LoadBalancerL7PolicyRuleNewParamsCompareType = "CONTAINS"
	LoadBalancerL7PolicyRuleNewParamsCompareTypeEndsWith   LoadBalancerL7PolicyRuleNewParamsCompareType = "ENDS_WITH"
	LoadBalancerL7PolicyRuleNewParamsCompareTypeEqualTo    LoadBalancerL7PolicyRuleNewParamsCompareType = "EQUAL_TO"
	LoadBalancerL7PolicyRuleNewParamsCompareTypeRegex      LoadBalancerL7PolicyRuleNewParamsCompareType = "REGEX"
	LoadBalancerL7PolicyRuleNewParamsCompareTypeStartsWith LoadBalancerL7PolicyRuleNewParamsCompareType = "STARTS_WITH"
)

// The L7 rule type
type LoadBalancerL7PolicyRuleNewParamsType string

const (
	LoadBalancerL7PolicyRuleNewParamsTypeCookie          LoadBalancerL7PolicyRuleNewParamsType = "COOKIE"
	LoadBalancerL7PolicyRuleNewParamsTypeFileType        LoadBalancerL7PolicyRuleNewParamsType = "FILE_TYPE"
	LoadBalancerL7PolicyRuleNewParamsTypeHeader          LoadBalancerL7PolicyRuleNewParamsType = "HEADER"
	LoadBalancerL7PolicyRuleNewParamsTypeHostName        LoadBalancerL7PolicyRuleNewParamsType = "HOST_NAME"
	LoadBalancerL7PolicyRuleNewParamsTypePath            LoadBalancerL7PolicyRuleNewParamsType = "PATH"
	LoadBalancerL7PolicyRuleNewParamsTypeSslConnHasCert  LoadBalancerL7PolicyRuleNewParamsType = "SSL_CONN_HAS_CERT"
	LoadBalancerL7PolicyRuleNewParamsTypeSslDnField      LoadBalancerL7PolicyRuleNewParamsType = "SSL_DN_FIELD"
	LoadBalancerL7PolicyRuleNewParamsTypeSslVerifyResult LoadBalancerL7PolicyRuleNewParamsType = "SSL_VERIFY_RESULT"
)

type LoadBalancerL7PolicyRuleListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type LoadBalancerL7PolicyRuleDeleteParams struct {
	ProjectID  param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID   param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	L7policyID string           `path:"l7policy_id,required" json:"-"`
	paramObj
}

type LoadBalancerL7PolicyRuleGetParams struct {
	ProjectID  param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID   param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	L7policyID string           `path:"l7policy_id,required" json:"-"`
	paramObj
}

type LoadBalancerL7PolicyRuleReplaceParams struct {
	ProjectID  param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID   param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	L7policyID string           `path:"l7policy_id,required" json:"-"`
	// When true the logic of the rule is inverted. For example, with invert true,
	// 'equal to' would become 'not equal to'. Default is false.
	Invert param.Opt[bool] `json:"invert,omitzero"`
	// The key to use for the comparison. For example, the name of the cookie to
	// evaluate.
	Key param.Opt[string] `json:"key,omitzero"`
	// The value to use for the comparison. For example, the file type to compare
	Value param.Opt[string] `json:"value,omitzero"`
	// The comparison type for the L7 rule
	//
	// Any of "CONTAINS", "ENDS_WITH", "EQUAL_TO", "REGEX", "STARTS_WITH".
	CompareType LoadBalancerL7PolicyRuleReplaceParamsCompareType `json:"compare_type,omitzero"`
	// A list of simple strings assigned to the l7 rule
	Tags []string `json:"tags,omitzero"`
	// The L7 rule type
	//
	// Any of "COOKIE", "FILE_TYPE", "HEADER", "HOST_NAME", "PATH",
	// "SSL_CONN_HAS_CERT", "SSL_DN_FIELD", "SSL_VERIFY_RESULT".
	Type LoadBalancerL7PolicyRuleReplaceParamsType `json:"type,omitzero"`
	paramObj
}

func (r LoadBalancerL7PolicyRuleReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyRuleReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyRuleReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The comparison type for the L7 rule
type LoadBalancerL7PolicyRuleReplaceParamsCompareType string

const (
	LoadBalancerL7PolicyRuleReplaceParamsCompareTypeContains   LoadBalancerL7PolicyRuleReplaceParamsCompareType = "CONTAINS"
	LoadBalancerL7PolicyRuleReplaceParamsCompareTypeEndsWith   LoadBalancerL7PolicyRuleReplaceParamsCompareType = "ENDS_WITH"
	LoadBalancerL7PolicyRuleReplaceParamsCompareTypeEqualTo    LoadBalancerL7PolicyRuleReplaceParamsCompareType = "EQUAL_TO"
	LoadBalancerL7PolicyRuleReplaceParamsCompareTypeRegex      LoadBalancerL7PolicyRuleReplaceParamsCompareType = "REGEX"
	LoadBalancerL7PolicyRuleReplaceParamsCompareTypeStartsWith LoadBalancerL7PolicyRuleReplaceParamsCompareType = "STARTS_WITH"
)

// The L7 rule type
type LoadBalancerL7PolicyRuleReplaceParamsType string

const (
	LoadBalancerL7PolicyRuleReplaceParamsTypeCookie          LoadBalancerL7PolicyRuleReplaceParamsType = "COOKIE"
	LoadBalancerL7PolicyRuleReplaceParamsTypeFileType        LoadBalancerL7PolicyRuleReplaceParamsType = "FILE_TYPE"
	LoadBalancerL7PolicyRuleReplaceParamsTypeHeader          LoadBalancerL7PolicyRuleReplaceParamsType = "HEADER"
	LoadBalancerL7PolicyRuleReplaceParamsTypeHostName        LoadBalancerL7PolicyRuleReplaceParamsType = "HOST_NAME"
	LoadBalancerL7PolicyRuleReplaceParamsTypePath            LoadBalancerL7PolicyRuleReplaceParamsType = "PATH"
	LoadBalancerL7PolicyRuleReplaceParamsTypeSslConnHasCert  LoadBalancerL7PolicyRuleReplaceParamsType = "SSL_CONN_HAS_CERT"
	LoadBalancerL7PolicyRuleReplaceParamsTypeSslDnField      LoadBalancerL7PolicyRuleReplaceParamsType = "SSL_DN_FIELD"
	LoadBalancerL7PolicyRuleReplaceParamsTypeSslVerifyResult LoadBalancerL7PolicyRuleReplaceParamsType = "SSL_VERIFY_RESULT"
)
