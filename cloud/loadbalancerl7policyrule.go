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
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s/rules", params.ProjectID.Value, params.RegionID.Value, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List load balancer L7 policy rules
func (r *LoadBalancerL7PolicyRuleService) List(ctx context.Context, l7policyID string, query LoadBalancerL7PolicyRuleListParams, opts ...option.RequestOption) (res *L7RuleList, err error) {
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
func (r *LoadBalancerL7PolicyRuleService) Get(ctx context.Context, l7ruleID string, query LoadBalancerL7PolicyRuleGetParams, opts ...option.RequestOption) (res *L7Rule, err error) {
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
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D%2Frules/post/parameters/0/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}/rules'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D%2Frules/post/parameters/1/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}/rules'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateL7RuleSchema/properties/compare_type'
	// "$.components.schemas.CreateL7RuleSchema.properties.compare_type"
	//
	// Any of "CONTAINS", "ENDS_WITH", "EQUAL_TO", "REGEX", "STARTS_WITH".
	CompareType LoadBalancerL7PolicyRuleNewParamsCompareType `json:"compare_type,omitzero,required"`
	// '#/components/schemas/CreateL7RuleSchema/properties/type'
	// "$.components.schemas.CreateL7RuleSchema.properties.type"
	//
	// Any of "COOKIE", "FILE_TYPE", "HEADER", "HOST_NAME", "PATH",
	// "SSL_CONN_HAS_CERT", "SSL_DN_FIELD", "SSL_VERIFY_RESULT".
	Type LoadBalancerL7PolicyRuleNewParamsType `json:"type,omitzero,required"`
	// '#/components/schemas/CreateL7RuleSchema/properties/value'
	// "$.components.schemas.CreateL7RuleSchema.properties.value"
	Value string `json:"value,required"`
	// '#/components/schemas/CreateL7RuleSchema/properties/invert'
	// "$.components.schemas.CreateL7RuleSchema.properties.invert"
	Invert param.Opt[bool] `json:"invert,omitzero"`
	// '#/components/schemas/CreateL7RuleSchema/properties/key'
	// "$.components.schemas.CreateL7RuleSchema.properties.key"
	Key param.Opt[string] `json:"key,omitzero"`
	// '#/components/schemas/CreateL7RuleSchema/properties/tags'
	// "$.components.schemas.CreateL7RuleSchema.properties.tags"
	Tags []string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerL7PolicyRuleNewParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r LoadBalancerL7PolicyRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateL7RuleSchema/properties/compare_type'
// "$.components.schemas.CreateL7RuleSchema.properties.compare_type"
type LoadBalancerL7PolicyRuleNewParamsCompareType string

const (
	LoadBalancerL7PolicyRuleNewParamsCompareTypeContains   LoadBalancerL7PolicyRuleNewParamsCompareType = "CONTAINS"
	LoadBalancerL7PolicyRuleNewParamsCompareTypeEndsWith   LoadBalancerL7PolicyRuleNewParamsCompareType = "ENDS_WITH"
	LoadBalancerL7PolicyRuleNewParamsCompareTypeEqualTo    LoadBalancerL7PolicyRuleNewParamsCompareType = "EQUAL_TO"
	LoadBalancerL7PolicyRuleNewParamsCompareTypeRegex      LoadBalancerL7PolicyRuleNewParamsCompareType = "REGEX"
	LoadBalancerL7PolicyRuleNewParamsCompareTypeStartsWith LoadBalancerL7PolicyRuleNewParamsCompareType = "STARTS_WITH"
)

// '#/components/schemas/CreateL7RuleSchema/properties/type'
// "$.components.schemas.CreateL7RuleSchema.properties.type"
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
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D%2Frules/get/parameters/0/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}/rules'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D%2Frules/get/parameters/1/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}/rules'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerL7PolicyRuleListParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type LoadBalancerL7PolicyRuleDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D%2Frules%2F%7Bl7rule_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}/rules/{l7rule_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D%2Frules%2F%7Bl7rule_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}/rules/{l7rule_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D%2Frules%2F%7Bl7rule_id%7D/delete/parameters/2/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}/rules/{l7rule_id}']['delete'].parameters[2].schema"
	L7policyID string `path:"l7policy_id,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerL7PolicyRuleDeleteParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type LoadBalancerL7PolicyRuleGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D%2Frules%2F%7Bl7rule_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}/rules/{l7rule_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D%2Frules%2F%7Bl7rule_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}/rules/{l7rule_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D%2Frules%2F%7Bl7rule_id%7D/get/parameters/2/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}/rules/{l7rule_id}'].get.parameters[2].schema"
	L7policyID string `path:"l7policy_id,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerL7PolicyRuleGetParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type LoadBalancerL7PolicyRuleReplaceParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D%2Frules%2F%7Bl7rule_id%7D/put/parameters/0/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}/rules/{l7rule_id}'].put.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D%2Frules%2F%7Bl7rule_id%7D/put/parameters/1/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}/rules/{l7rule_id}'].put.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D%2Frules%2F%7Bl7rule_id%7D/put/parameters/2/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}/rules/{l7rule_id}'].put.parameters[2].schema"
	L7policyID string `path:"l7policy_id,required" json:"-"`
	// '#/components/schemas/UpdateL7RuleSchema/properties/invert'
	// "$.components.schemas.UpdateL7RuleSchema.properties.invert"
	Invert param.Opt[bool] `json:"invert,omitzero"`
	// '#/components/schemas/UpdateL7RuleSchema/properties/key'
	// "$.components.schemas.UpdateL7RuleSchema.properties.key"
	Key param.Opt[string] `json:"key,omitzero"`
	// '#/components/schemas/UpdateL7RuleSchema/properties/value'
	// "$.components.schemas.UpdateL7RuleSchema.properties.value"
	Value param.Opt[string] `json:"value,omitzero"`
	// '#/components/schemas/UpdateL7RuleSchema/properties/compare_type'
	// "$.components.schemas.UpdateL7RuleSchema.properties.compare_type"
	//
	// Any of "CONTAINS", "ENDS_WITH", "EQUAL_TO", "REGEX", "STARTS_WITH".
	CompareType LoadBalancerL7PolicyRuleReplaceParamsCompareType `json:"compare_type,omitzero"`
	// '#/components/schemas/UpdateL7RuleSchema/properties/tags'
	// "$.components.schemas.UpdateL7RuleSchema.properties.tags"
	Tags []string `json:"tags,omitzero"`
	// '#/components/schemas/UpdateL7RuleSchema/properties/type'
	// "$.components.schemas.UpdateL7RuleSchema.properties.type"
	//
	// Any of "COOKIE", "FILE_TYPE", "HEADER", "HOST_NAME", "PATH",
	// "SSL_CONN_HAS_CERT", "SSL_DN_FIELD", "SSL_VERIFY_RESULT".
	Type LoadBalancerL7PolicyRuleReplaceParamsType `json:"type,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerL7PolicyRuleReplaceParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r LoadBalancerL7PolicyRuleReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyRuleReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/UpdateL7RuleSchema/properties/compare_type'
// "$.components.schemas.UpdateL7RuleSchema.properties.compare_type"
type LoadBalancerL7PolicyRuleReplaceParamsCompareType string

const (
	LoadBalancerL7PolicyRuleReplaceParamsCompareTypeContains   LoadBalancerL7PolicyRuleReplaceParamsCompareType = "CONTAINS"
	LoadBalancerL7PolicyRuleReplaceParamsCompareTypeEndsWith   LoadBalancerL7PolicyRuleReplaceParamsCompareType = "ENDS_WITH"
	LoadBalancerL7PolicyRuleReplaceParamsCompareTypeEqualTo    LoadBalancerL7PolicyRuleReplaceParamsCompareType = "EQUAL_TO"
	LoadBalancerL7PolicyRuleReplaceParamsCompareTypeRegex      LoadBalancerL7PolicyRuleReplaceParamsCompareType = "REGEX"
	LoadBalancerL7PolicyRuleReplaceParamsCompareTypeStartsWith LoadBalancerL7PolicyRuleReplaceParamsCompareType = "STARTS_WITH"
)

// '#/components/schemas/UpdateL7RuleSchema/properties/type'
// "$.components.schemas.UpdateL7RuleSchema.properties.type"
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
