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

// LoadBalancerL7PolicyService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoadBalancerL7PolicyService] method instead.
type LoadBalancerL7PolicyService struct {
	Options []option.RequestOption
	Rules   LoadBalancerL7PolicyRuleService
}

// NewLoadBalancerL7PolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerL7PolicyService(opts ...option.RequestOption) (r LoadBalancerL7PolicyService) {
	r = LoadBalancerL7PolicyService{}
	r.Options = opts
	r.Rules = NewLoadBalancerL7PolicyRuleService(opts...)
	return
}

// Create load balancer L7 policy
func (r *LoadBalancerL7PolicyService) New(ctx context.Context, params LoadBalancerL7PolicyNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List load balancer L7 policies
func (r *LoadBalancerL7PolicyService) List(ctx context.Context, query LoadBalancerL7PolicyListParams, opts ...option.RequestOption) (res *L7PolicyList, err error) {
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
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v", query.ProjectID.Value, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete load balancer L7 policy
func (r *LoadBalancerL7PolicyService) Delete(ctx context.Context, l7policyID string, body LoadBalancerL7PolicyDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get load balancer L7 policy
func (r *LoadBalancerL7PolicyService) Get(ctx context.Context, l7policyID string, query LoadBalancerL7PolicyGetParams, opts ...option.RequestOption) (res *L7Policy, err error) {
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
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Replace load balancer L7 policy properties
func (r *LoadBalancerL7PolicyService) Replace(ctx context.Context, l7policyID string, params LoadBalancerL7PolicyReplaceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type LoadBalancerL7PolicyNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateL7PolicySchema/properties/action'
	// "$.components.schemas.CreateL7PolicySchema.properties.action"
	//
	// Any of "REDIRECT_PREFIX", "REDIRECT_TO_POOL", "REDIRECT_TO_URL", "REJECT".
	Action LoadBalancerL7PolicyNewParamsAction `json:"action,omitzero,required"`
	// '#/components/schemas/CreateL7PolicySchema/properties/listener_id'
	// "$.components.schemas.CreateL7PolicySchema.properties.listener_id"
	ListenerID string `json:"listener_id,required"`
	// '#/components/schemas/CreateL7PolicySchema/properties/name'
	// "$.components.schemas.CreateL7PolicySchema.properties.name"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/CreateL7PolicySchema/properties/position'
	// "$.components.schemas.CreateL7PolicySchema.properties.position"
	Position param.Opt[int64] `json:"position,omitzero"`
	// '#/components/schemas/CreateL7PolicySchema/properties/redirect_http_code'
	// "$.components.schemas.CreateL7PolicySchema.properties.redirect_http_code"
	RedirectHTTPCode param.Opt[int64] `json:"redirect_http_code,omitzero"`
	// '#/components/schemas/CreateL7PolicySchema/properties/redirect_pool_id'
	// "$.components.schemas.CreateL7PolicySchema.properties.redirect_pool_id"
	RedirectPoolID param.Opt[string] `json:"redirect_pool_id,omitzero"`
	// '#/components/schemas/CreateL7PolicySchema/properties/redirect_prefix'
	// "$.components.schemas.CreateL7PolicySchema.properties.redirect_prefix"
	RedirectPrefix param.Opt[string] `json:"redirect_prefix,omitzero"`
	// '#/components/schemas/CreateL7PolicySchema/properties/redirect_url'
	// "$.components.schemas.CreateL7PolicySchema.properties.redirect_url"
	RedirectURL param.Opt[string] `json:"redirect_url,omitzero"`
	// '#/components/schemas/CreateL7PolicySchema/properties/tags'
	// "$.components.schemas.CreateL7PolicySchema.properties.tags"
	Tags []string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerL7PolicyNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LoadBalancerL7PolicyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateL7PolicySchema/properties/action'
// "$.components.schemas.CreateL7PolicySchema.properties.action"
type LoadBalancerL7PolicyNewParamsAction string

const (
	LoadBalancerL7PolicyNewParamsActionRedirectPrefix LoadBalancerL7PolicyNewParamsAction = "REDIRECT_PREFIX"
	LoadBalancerL7PolicyNewParamsActionRedirectToPool LoadBalancerL7PolicyNewParamsAction = "REDIRECT_TO_POOL"
	LoadBalancerL7PolicyNewParamsActionRedirectToURL  LoadBalancerL7PolicyNewParamsAction = "REDIRECT_TO_URL"
	LoadBalancerL7PolicyNewParamsActionReject         LoadBalancerL7PolicyNewParamsAction = "REJECT"
)

type LoadBalancerL7PolicyListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerL7PolicyListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type LoadBalancerL7PolicyDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerL7PolicyDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type LoadBalancerL7PolicyGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerL7PolicyGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type LoadBalancerL7PolicyReplaceParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D/put/parameters/0/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}'].put.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fl7policies%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bl7policy_id%7D/put/parameters/1/schema'
	// "$.paths['/cloud/v1/l7policies/{project_id}/{region_id}/{l7policy_id}'].put.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/UpdateL7PolicySchema/properties/action'
	// "$.components.schemas.UpdateL7PolicySchema.properties.action"
	//
	// Any of "REDIRECT_PREFIX", "REDIRECT_TO_POOL", "REDIRECT_TO_URL", "REJECT".
	Action LoadBalancerL7PolicyReplaceParamsAction `json:"action,omitzero,required"`
	// '#/components/schemas/UpdateL7PolicySchema/properties/name'
	// "$.components.schemas.UpdateL7PolicySchema.properties.name"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/UpdateL7PolicySchema/properties/position'
	// "$.components.schemas.UpdateL7PolicySchema.properties.position"
	Position param.Opt[int64] `json:"position,omitzero"`
	// '#/components/schemas/UpdateL7PolicySchema/properties/redirect_http_code'
	// "$.components.schemas.UpdateL7PolicySchema.properties.redirect_http_code"
	RedirectHTTPCode param.Opt[int64] `json:"redirect_http_code,omitzero"`
	// '#/components/schemas/UpdateL7PolicySchema/properties/redirect_pool_id'
	// "$.components.schemas.UpdateL7PolicySchema.properties.redirect_pool_id"
	RedirectPoolID param.Opt[string] `json:"redirect_pool_id,omitzero"`
	// '#/components/schemas/UpdateL7PolicySchema/properties/redirect_prefix'
	// "$.components.schemas.UpdateL7PolicySchema.properties.redirect_prefix"
	RedirectPrefix param.Opt[string] `json:"redirect_prefix,omitzero"`
	// '#/components/schemas/UpdateL7PolicySchema/properties/redirect_url'
	// "$.components.schemas.UpdateL7PolicySchema.properties.redirect_url"
	RedirectURL param.Opt[string] `json:"redirect_url,omitzero"`
	// '#/components/schemas/UpdateL7PolicySchema/properties/tags'
	// "$.components.schemas.UpdateL7PolicySchema.properties.tags"
	Tags []string `json:"tags,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerL7PolicyReplaceParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r LoadBalancerL7PolicyReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/UpdateL7PolicySchema/properties/action'
// "$.components.schemas.UpdateL7PolicySchema.properties.action"
type LoadBalancerL7PolicyReplaceParamsAction string

const (
	LoadBalancerL7PolicyReplaceParamsActionRedirectPrefix LoadBalancerL7PolicyReplaceParamsAction = "REDIRECT_PREFIX"
	LoadBalancerL7PolicyReplaceParamsActionRedirectToPool LoadBalancerL7PolicyReplaceParamsAction = "REDIRECT_TO_POOL"
	LoadBalancerL7PolicyReplaceParamsActionRedirectToURL  LoadBalancerL7PolicyReplaceParamsAction = "REDIRECT_TO_URL"
	LoadBalancerL7PolicyReplaceParamsActionReject         LoadBalancerL7PolicyReplaceParamsAction = "REJECT"
)
