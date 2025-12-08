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
	tasks   TaskService
}

// NewLoadBalancerL7PolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerL7PolicyService(opts ...option.RequestOption) (r LoadBalancerL7PolicyService) {
	r = LoadBalancerL7PolicyService{}
	r.Options = opts
	r.Rules = NewLoadBalancerL7PolicyRuleService(opts...)
	r.tasks = NewTaskService(opts...)
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

// NewAndPoll creates a new L7 policy and polls for completion
func (r *LoadBalancerL7PolicyService) NewAndPoll(ctx context.Context, params LoadBalancerL7PolicyNewParams, opts ...option.RequestOption) (v *LoadBalancerL7Policy, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams LoadBalancerL7PolicyGetParams
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

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.L7polices) != 1 {
		return nil, errors.New("expected exactly one L7 policy to be created in a task")
	}
	resourceID := task.CreatedResources.L7polices[0]

	return r.Get(ctx, resourceID, getParams, opts...)
}

// DeleteAndPoll deletes an L7 policy and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *LoadBalancerL7PolicyService) DeleteAndPoll(ctx context.Context, l7policyID string, body LoadBalancerL7PolicyDeleteParams, opts ...option.RequestOption) error {
	resource, err := r.Delete(ctx, l7policyID, body, opts...)
	if err != nil {
		return err
	}

	opts = slices.Concat(r.Options, opts)
	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	return err
}

// ReplaceAndPoll replaces an L7 policy and polls for completion of the first task. Use the [TaskService.Poll] method if
// you need to poll for all tasks.
func (r *LoadBalancerL7PolicyService) ReplaceAndPoll(ctx context.Context, l7policyID string, params LoadBalancerL7PolicyReplaceParams, opts ...option.RequestOption) (v *LoadBalancerL7Policy, err error) {
	resource, err := r.Replace(ctx, l7policyID, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams LoadBalancerL7PolicyGetParams
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
		return
	}

	return r.Get(ctx, l7policyID, getParams, opts...)
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
