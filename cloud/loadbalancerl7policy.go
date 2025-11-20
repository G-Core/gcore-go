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
	"github.com/G-Core/gcore-go/packages/respjson"
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

// List load balancer L7 policies
func (r *LoadBalancerL7PolicyService) List(ctx context.Context, query LoadBalancerL7PolicyListParams, opts ...option.RequestOption) (res *LoadBalancerL7PolicyListResponse, err error) {
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
func (r *LoadBalancerL7PolicyService) Get(ctx context.Context, l7policyID string, query LoadBalancerL7PolicyGetParams, opts ...option.RequestOption) (res *LoadBalancerL7PolicyGetResponse, err error) {
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

// Replaces the entire L7 policy configuration with the provided data. Any fields
// omitted from the request body will be unset (set to null) or reset to their
// default values (such as "`redirect_http_code`") depending on the "action". This
// is a destructive operation that overwrites the complete policy configuration.
func (r *LoadBalancerL7PolicyService) Replace(ctx context.Context, l7policyID string, params LoadBalancerL7PolicyReplaceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type LoadBalancerL7PolicyListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []LoadBalancerL7PolicyListResponseResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerL7PolicyListResponse) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerL7PolicyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerL7PolicyListResponseResult struct {
	// ID
	ID string `json:"id,required"`
	// Action
	//
	// Any of "REDIRECT_PREFIX", "REDIRECT_TO_POOL", "REDIRECT_TO_URL", "REJECT".
	Action string `json:"action,required"`
	// Listener ID
	ListenerID string `json:"listener_id,required"`
	// Human-readable name of the policy
	Name string `json:"name,required"`
	// L7 policy operating status
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// The position of this policy on the listener. Positions start at 1.
	Position int64 `json:"position,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Requests matching this policy will be redirected to the specified URL or Prefix
	// URL with the HTTP response code. Valid if action is `REDIRECT_TO_URL` or
	// `REDIRECT_PREFIX`. Valid options are 301, 302, 303, 307, or 308. Default is 302.
	RedirectHTTPCode int64 `json:"redirect_http_code,required"`
	// Requests matching this policy will be redirected to the pool with this ID. Only
	// valid if action is `REDIRECT_TO_POOL`.
	RedirectPoolID string `json:"redirect_pool_id,required"`
	// Requests matching this policy will be redirected to this Prefix URL. Only valid
	// if action is `REDIRECT_PREFIX`.
	RedirectPrefix string `json:"redirect_prefix,required"`
	// Requests matching this policy will be redirected to this URL. Only valid if
	// action is `REDIRECT_TO_URL`.
	RedirectURL string `json:"redirect_url,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Rules. All the rules associated with a given policy are logically ANDed
	// together. A request must match all the policy’s rules to match the policy.If you
	// need to express a logical OR operation between rules, then do this by creating
	// multiple policies with the same action.
	Rules []LoadBalancerL7PolicyListResponseResultRule `json:"rules,required"`
	// A list of simple strings assigned to the resource.
	Tags []string `json:"tags,required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Action             respjson.Field
		ListenerID         respjson.Field
		Name               respjson.Field
		OperatingStatus    respjson.Field
		Position           respjson.Field
		ProjectID          respjson.Field
		ProvisioningStatus respjson.Field
		RedirectHTTPCode   respjson.Field
		RedirectPoolID     respjson.Field
		RedirectPrefix     respjson.Field
		RedirectURL        respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Rules              respjson.Field
		Tags               respjson.Field
		TaskID             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerL7PolicyListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerL7PolicyListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerL7PolicyListResponseResultRule struct {
	// L7Rule ID
	ID string `json:"id,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ProjectID   respjson.Field
		Region      respjson.Field
		RegionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerL7PolicyListResponseResultRule) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerL7PolicyListResponseResultRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerL7PolicyGetResponse struct {
	// ID
	ID string `json:"id,required"`
	// Action
	//
	// Any of "REDIRECT_PREFIX", "REDIRECT_TO_POOL", "REDIRECT_TO_URL", "REJECT".
	Action LoadBalancerL7PolicyGetResponseAction `json:"action,required"`
	// Listener ID
	ListenerID string `json:"listener_id,required"`
	// Human-readable name of the policy
	Name string `json:"name,required"`
	// L7 policy operating status
	//
	// Any of "DEGRADED", "DRAINING", "ERROR", "NO_MONITOR", "OFFLINE", "ONLINE".
	OperatingStatus LoadBalancerOperatingStatus `json:"operating_status,required"`
	// The position of this policy on the listener. Positions start at 1.
	Position int64 `json:"position,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Any of "ACTIVE", "DELETED", "ERROR", "PENDING_CREATE", "PENDING_DELETE",
	// "PENDING_UPDATE".
	ProvisioningStatus ProvisioningStatus `json:"provisioning_status,required"`
	// Requests matching this policy will be redirected to the specified URL or Prefix
	// URL with the HTTP response code. Valid if action is `REDIRECT_TO_URL` or
	// `REDIRECT_PREFIX`. Valid options are 301, 302, 303, 307, or 308. Default is 302.
	RedirectHTTPCode int64 `json:"redirect_http_code,required"`
	// Requests matching this policy will be redirected to the pool with this ID. Only
	// valid if action is `REDIRECT_TO_POOL`.
	RedirectPoolID string `json:"redirect_pool_id,required"`
	// Requests matching this policy will be redirected to this Prefix URL. Only valid
	// if action is `REDIRECT_PREFIX`.
	RedirectPrefix string `json:"redirect_prefix,required"`
	// Requests matching this policy will be redirected to this URL. Only valid if
	// action is `REDIRECT_TO_URL`.
	RedirectURL string `json:"redirect_url,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Rules. All the rules associated with a given policy are logically ANDed
	// together. A request must match all the policy’s rules to match the policy.If you
	// need to express a logical OR operation between rules, then do this by creating
	// multiple policies with the same action.
	Rules []LoadBalancerL7PolicyGetResponseRule `json:"rules,required"`
	// A list of simple strings assigned to the resource.
	Tags []string `json:"tags,required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,required" format:"uuid4"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Action             respjson.Field
		ListenerID         respjson.Field
		Name               respjson.Field
		OperatingStatus    respjson.Field
		Position           respjson.Field
		ProjectID          respjson.Field
		ProvisioningStatus respjson.Field
		RedirectHTTPCode   respjson.Field
		RedirectPoolID     respjson.Field
		RedirectPrefix     respjson.Field
		RedirectURL        respjson.Field
		Region             respjson.Field
		RegionID           respjson.Field
		Rules              respjson.Field
		Tags               respjson.Field
		TaskID             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerL7PolicyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerL7PolicyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action
type LoadBalancerL7PolicyGetResponseAction string

const (
	LoadBalancerL7PolicyGetResponseActionRedirectPrefix LoadBalancerL7PolicyGetResponseAction = "REDIRECT_PREFIX"
	LoadBalancerL7PolicyGetResponseActionRedirectToPool LoadBalancerL7PolicyGetResponseAction = "REDIRECT_TO_POOL"
	LoadBalancerL7PolicyGetResponseActionRedirectToURL  LoadBalancerL7PolicyGetResponseAction = "REDIRECT_TO_URL"
	LoadBalancerL7PolicyGetResponseActionReject         LoadBalancerL7PolicyGetResponseAction = "REJECT"
)

type LoadBalancerL7PolicyGetResponseRule struct {
	// L7Rule ID
	ID string `json:"id,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ProjectID   respjson.Field
		Region      respjson.Field
		RegionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerL7PolicyGetResponseRule) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerL7PolicyGetResponseRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type LoadBalancerL7PolicyReplaceParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfRedirectToURL *LoadBalancerL7PolicyReplaceParamsBodyRedirectToURL `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfRedirectPrefix *LoadBalancerL7PolicyReplaceParamsBodyRedirectPrefix `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfRedirectToPool *LoadBalancerL7PolicyReplaceParamsBodyRedirectToPool `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfReject *LoadBalancerL7PolicyReplaceParamsBodyReject `json:",inline"`

	paramObj
}

func (u LoadBalancerL7PolicyReplaceParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRedirectToURL, u.OfRedirectPrefix, u.OfRedirectToPool, u.OfReject)
}
func (r *LoadBalancerL7PolicyReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, RedirectURL are required.
type LoadBalancerL7PolicyReplaceParamsBodyRedirectToURL struct {
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

func (r LoadBalancerL7PolicyReplaceParamsBodyRedirectToURL) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyReplaceParamsBodyRedirectToURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyReplaceParamsBodyRedirectToURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, RedirectPrefix are required.
type LoadBalancerL7PolicyReplaceParamsBodyRedirectPrefix struct {
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

func (r LoadBalancerL7PolicyReplaceParamsBodyRedirectPrefix) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyReplaceParamsBodyRedirectPrefix
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyReplaceParamsBodyRedirectPrefix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, RedirectPoolID are required.
type LoadBalancerL7PolicyReplaceParamsBodyRedirectToPool struct {
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

func (r LoadBalancerL7PolicyReplaceParamsBodyRedirectToPool) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyReplaceParamsBodyRedirectToPool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyReplaceParamsBodyRedirectToPool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Action is required.
type LoadBalancerL7PolicyReplaceParamsBodyReject struct {
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

func (r LoadBalancerL7PolicyReplaceParamsBodyReject) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyReplaceParamsBodyReject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyReplaceParamsBodyReject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
