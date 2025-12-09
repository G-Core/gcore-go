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
	"github.com/G-Core/gcore-go/shared/constant"
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
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates only provided fields; omitted ones stay unchanged.
func (r *LoadBalancerL7PolicyService) Update(ctx context.Context, l7policyID string, params LoadBalancerL7PolicyUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List load balancer L7 policies
func (r *LoadBalancerL7PolicyService) List(ctx context.Context, query LoadBalancerL7PolicyListParams, opts ...option.RequestOption) (res *LoadBalancerL7PolicyList, err error) {
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
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v", query.ProjectID.Value, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete load balancer L7 policy
func (r *LoadBalancerL7PolicyService) Delete(ctx context.Context, l7policyID string, body LoadBalancerL7PolicyDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get load balancer L7 policy
func (r *LoadBalancerL7PolicyService) Get(ctx context.Context, l7policyID string, query LoadBalancerL7PolicyGetParams, opts ...option.RequestOption) (res *LoadBalancerL7Policy, err error) {
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
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LoadBalancerL7PolicyNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfRedirectToURL *LoadBalancerL7PolicyNewParamsBodyRedirectToURL `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfRedirectPrefix *LoadBalancerL7PolicyNewParamsBodyRedirectPrefix `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfRedirectToPool *LoadBalancerL7PolicyNewParamsBodyRedirectToPool `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfReject *LoadBalancerL7PolicyNewParamsBodyReject `json:",inline"`

	paramObj
}

func (u LoadBalancerL7PolicyNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRedirectToURL, u.OfRedirectPrefix, u.OfRedirectToPool, u.OfReject)
}
func (r *LoadBalancerL7PolicyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, ListenerID, RedirectURL are required.
type LoadBalancerL7PolicyNewParamsBodyRedirectToURL struct {
	// Listener ID
	ListenerID string `json:"listener_id,required"`
	// Requests matching this policy will be redirected to this URL.
	RedirectURL string `json:"redirect_url,required"`
	// Human-readable name of the policy
	Name param.Opt[string] `json:"name,omitzero"`
	// The position of this policy on the listener
	Position param.Opt[int64] `json:"position,omitzero"`
	// Requests matching this policy will be redirected to the specified URL or Prefix
	// URL with the HTTP response code. Valid options are 301, 302, 303, 307, or 308.
	// Default is 302.
	RedirectHTTPCode param.Opt[int64] `json:"redirect_http_code,omitzero"`
	// A list of simple strings assigned to the resource.
	Tags []string `json:"tags,omitzero"`
	// Action
	//
	// This field can be elided, and will marshal its zero value as "REDIRECT_TO_URL".
	Action constant.RedirectToURL `json:"action,required"`
	paramObj
}

func (r LoadBalancerL7PolicyNewParamsBodyRedirectToURL) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyNewParamsBodyRedirectToURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyNewParamsBodyRedirectToURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, ListenerID, RedirectPrefix are required.
type LoadBalancerL7PolicyNewParamsBodyRedirectPrefix struct {
	// Listener ID
	ListenerID string `json:"listener_id,required"`
	// Requests matching this policy will be redirected to this Prefix URL.
	RedirectPrefix string `json:"redirect_prefix,required"`
	// Human-readable name of the policy
	Name param.Opt[string] `json:"name,omitzero"`
	// The position of this policy on the listener
	Position param.Opt[int64] `json:"position,omitzero"`
	// Requests matching this policy will be redirected to the specified URL or Prefix
	// URL with the HTTP response code. Valid options are 301, 302, 303, 307, or 308.
	// Default is 302.
	RedirectHTTPCode param.Opt[int64] `json:"redirect_http_code,omitzero"`
	// A list of simple strings assigned to the resource.
	Tags []string `json:"tags,omitzero"`
	// Action
	//
	// This field can be elided, and will marshal its zero value as "REDIRECT_PREFIX".
	Action constant.RedirectPrefix `json:"action,required"`
	paramObj
}

func (r LoadBalancerL7PolicyNewParamsBodyRedirectPrefix) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyNewParamsBodyRedirectPrefix
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyNewParamsBodyRedirectPrefix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, ListenerID, RedirectPoolID are required.
type LoadBalancerL7PolicyNewParamsBodyRedirectToPool struct {
	// Listener ID
	ListenerID string `json:"listener_id,required"`
	// Requests matching this policy will be redirected to the pool with this ID.
	RedirectPoolID string `json:"redirect_pool_id,required"`
	// Human-readable name of the policy
	Name param.Opt[string] `json:"name,omitzero"`
	// The position of this policy on the listener
	Position param.Opt[int64] `json:"position,omitzero"`
	// A list of simple strings assigned to the resource.
	Tags []string `json:"tags,omitzero"`
	// Action
	//
	// This field can be elided, and will marshal its zero value as "REDIRECT_TO_POOL".
	Action constant.RedirectToPool `json:"action,required"`
	paramObj
}

func (r LoadBalancerL7PolicyNewParamsBodyRedirectToPool) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyNewParamsBodyRedirectToPool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyNewParamsBodyRedirectToPool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, ListenerID are required.
type LoadBalancerL7PolicyNewParamsBodyReject struct {
	// Listener ID
	ListenerID string `json:"listener_id,required"`
	// Human-readable name of the policy
	Name param.Opt[string] `json:"name,omitzero"`
	// The position of this policy on the listener
	Position param.Opt[int64] `json:"position,omitzero"`
	// A list of simple strings assigned to the resource.
	Tags []string `json:"tags,omitzero"`
	// Action
	//
	// This field can be elided, and will marshal its zero value as "REJECT".
	Action constant.Reject `json:"action,required"`
	paramObj
}

func (r LoadBalancerL7PolicyNewParamsBodyReject) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyNewParamsBodyReject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyNewParamsBodyReject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerL7PolicyUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfRedirectToURL *LoadBalancerL7PolicyUpdateParamsBodyRedirectToURL `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfRedirectPrefix *LoadBalancerL7PolicyUpdateParamsBodyRedirectPrefix `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfRedirectToPool *LoadBalancerL7PolicyUpdateParamsBodyRedirectToPool `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfReject *LoadBalancerL7PolicyUpdateParamsBodyReject `json:",inline"`

	paramObj
}

func (u LoadBalancerL7PolicyUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRedirectToURL, u.OfRedirectPrefix, u.OfRedirectToPool, u.OfReject)
}
func (r *LoadBalancerL7PolicyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, RedirectURL are required.
type LoadBalancerL7PolicyUpdateParamsBodyRedirectToURL struct {
	// Requests matching this policy will be redirected to this URL. Only valid if
	// action is `REDIRECT_TO_URL`.
	RedirectURL string `json:"redirect_url,required"`
	// Human-readable name of the policy
	Name param.Opt[string] `json:"name,omitzero"`
	// The position of this policy on the listener
	Position param.Opt[int64] `json:"position,omitzero"`
	// Requests matching this policy will be redirected to the specified URL or Prefix
	// URL with the HTTP response code. Valid if action is `REDIRECT_TO_URL` or
	// `REDIRECT_PREFIX`. Valid options are 301, 302, 303, 307, or 308. Default is 302.
	RedirectHTTPCode param.Opt[int64] `json:"redirect_http_code,omitzero"`
	// A list of simple strings assigned to the resource.
	Tags []string `json:"tags,omitzero"`
	// Action
	//
	// This field can be elided, and will marshal its zero value as "REDIRECT_TO_URL".
	Action constant.RedirectToURL `json:"action,required"`
	paramObj
}

func (r LoadBalancerL7PolicyUpdateParamsBodyRedirectToURL) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyUpdateParamsBodyRedirectToURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyUpdateParamsBodyRedirectToURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, RedirectPrefix are required.
type LoadBalancerL7PolicyUpdateParamsBodyRedirectPrefix struct {
	// Requests matching this policy will be redirected to this Prefix URL.
	RedirectPrefix string `json:"redirect_prefix,required"`
	// Human-readable name of the policy
	Name param.Opt[string] `json:"name,omitzero"`
	// The position of this policy on the listener
	Position param.Opt[int64] `json:"position,omitzero"`
	// Requests matching this policy will be redirected to the specified URL or Prefix
	// URL with the HTTP response code. Valid options are 301, 302, 303, 307, or 308.
	// Default is 302.
	RedirectHTTPCode param.Opt[int64] `json:"redirect_http_code,omitzero"`
	// A list of simple strings assigned to the resource.
	Tags []string `json:"tags,omitzero"`
	// Action
	//
	// This field can be elided, and will marshal its zero value as "REDIRECT_PREFIX".
	Action constant.RedirectPrefix `json:"action,required"`
	paramObj
}

func (r LoadBalancerL7PolicyUpdateParamsBodyRedirectPrefix) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyUpdateParamsBodyRedirectPrefix
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyUpdateParamsBodyRedirectPrefix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, RedirectPoolID are required.
type LoadBalancerL7PolicyUpdateParamsBodyRedirectToPool struct {
	// Requests matching this policy will be redirected to the pool with this ID.
	RedirectPoolID string `json:"redirect_pool_id,required"`
	// Human-readable name of the policy
	Name param.Opt[string] `json:"name,omitzero"`
	// The position of this policy on the listener
	Position param.Opt[int64] `json:"position,omitzero"`
	// A list of simple strings assigned to the resource.
	Tags []string `json:"tags,omitzero"`
	// Action
	//
	// This field can be elided, and will marshal its zero value as "REDIRECT_TO_POOL".
	Action constant.RedirectToPool `json:"action,required"`
	paramObj
}

func (r LoadBalancerL7PolicyUpdateParamsBodyRedirectToPool) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyUpdateParamsBodyRedirectToPool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyUpdateParamsBodyRedirectToPool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Action is required.
type LoadBalancerL7PolicyUpdateParamsBodyReject struct {
	// Human-readable name of the policy
	Name param.Opt[string] `json:"name,omitzero"`
	// The position of this policy on the listener
	Position param.Opt[int64] `json:"position,omitzero"`
	// A list of simple strings assigned to the resource.
	Tags []string `json:"tags,omitzero"`
	// Action
	//
	// This field can be elided, and will marshal its zero value as "REJECT".
	Action constant.Reject `json:"action,required"`
	paramObj
}

func (r LoadBalancerL7PolicyUpdateParamsBodyReject) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyUpdateParamsBodyReject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyUpdateParamsBodyReject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerL7PolicyListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type LoadBalancerL7PolicyDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type LoadBalancerL7PolicyGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
