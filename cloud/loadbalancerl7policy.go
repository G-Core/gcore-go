// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
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
	opts = append(r.Options[:], opts...)
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
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List load balancer L7 policies
func (r *LoadBalancerL7PolicyService) List(ctx context.Context, query LoadBalancerL7PolicyListParams, opts ...option.RequestOption) (res *LoadBalancerL7PolicyList, err error) {
	opts = append(r.Options[:], opts...)
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
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v", query.ProjectID.Value, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete load balancer L7 policy
func (r *LoadBalancerL7PolicyService) Delete(ctx context.Context, l7policyID string, body LoadBalancerL7PolicyDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Replace load balancer L7 policy
func (r *LoadBalancerL7PolicyService) Replace(ctx context.Context, l7policyID string, params LoadBalancerL7PolicyReplaceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
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
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// NewAndPoll creates a new L7 policy and polls for completion
func (r *LoadBalancerL7PolicyService) NewAndPoll(ctx context.Context, params LoadBalancerL7PolicyNewParams, opts ...option.RequestOption) (v *LoadBalancerL7Policy, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = append(r.Options[:], opts...)
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

	opts = append(r.Options[:], opts...)
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

	opts = append(r.Options[:], opts...)
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
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Action
	//
	// Any of "REDIRECT_PREFIX", "REDIRECT_TO_POOL", "REDIRECT_TO_URL", "REJECT".
	Action LoadBalancerL7PolicyNewParamsAction `json:"action,omitzero,required"`
	// Listener ID
	ListenerID string `json:"listener_id,required"`
	// Human-readable name of the policy
	Name param.Opt[string] `json:"name,omitzero"`
	// The position of this policy on the listener. Positions start at 1.
	Position param.Opt[int64] `json:"position,omitzero"`
	// Requests matching this policy will be redirected to the specified URL or Prefix
	// URL with the HTTP response code. Valid if action is `REDIRECT_TO_URL` or
	// `REDIRECT_PREFIX`. Valid options are 301, 302, 303, 307, or 308. Default is 302.
	RedirectHTTPCode param.Opt[int64] `json:"redirect_http_code,omitzero"`
	// Requests matching this policy will be redirected to the pool withthis ID. Only
	// valid if action is `REDIRECT_TO_POOL`.
	RedirectPoolID param.Opt[string] `json:"redirect_pool_id,omitzero"`
	// Requests matching this policy will be redirected to this Prefix URL. Only valid
	// if action is `REDIRECT_PREFIX`.
	RedirectPrefix param.Opt[string] `json:"redirect_prefix,omitzero"`
	// Requests matching this policy will be redirected to this URL. Only valid if
	// action is `REDIRECT_TO_URL`.
	RedirectURL param.Opt[string] `json:"redirect_url,omitzero"`
	// A list of simple strings assigned to the resource.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r LoadBalancerL7PolicyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action
type LoadBalancerL7PolicyNewParamsAction string

const (
	LoadBalancerL7PolicyNewParamsActionRedirectPrefix LoadBalancerL7PolicyNewParamsAction = "REDIRECT_PREFIX"
	LoadBalancerL7PolicyNewParamsActionRedirectToPool LoadBalancerL7PolicyNewParamsAction = "REDIRECT_TO_POOL"
	LoadBalancerL7PolicyNewParamsActionRedirectToURL  LoadBalancerL7PolicyNewParamsAction = "REDIRECT_TO_URL"
	LoadBalancerL7PolicyNewParamsActionReject         LoadBalancerL7PolicyNewParamsAction = "REJECT"
)

type LoadBalancerL7PolicyListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type LoadBalancerL7PolicyDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type LoadBalancerL7PolicyGetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type LoadBalancerL7PolicyReplaceParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Action
	//
	// Any of "REDIRECT_PREFIX", "REDIRECT_TO_POOL", "REDIRECT_TO_URL", "REJECT".
	Action LoadBalancerL7PolicyReplaceParamsAction `json:"action,omitzero,required"`
	// Human-readable name of the policy
	Name param.Opt[string] `json:"name,omitzero"`
	// The position of this policy on the listener. Positions start at 1.
	Position param.Opt[int64] `json:"position,omitzero"`
	// Requests matching this policy will be redirected to the specified URL or Prefix
	// URL with the HTTP response code. Valid if action is `REDIRECT_TO_URL` or
	// `REDIRECT_PREFIX`. Valid options are 301, 302, 303, 307, or 308. Default is 302.
	RedirectHTTPCode param.Opt[int64] `json:"redirect_http_code,omitzero"`
	// Requests matching this policy will be redirected to the pool with this ID. Only
	// valid if action is `REDIRECT_TO_POOL`.
	RedirectPoolID param.Opt[string] `json:"redirect_pool_id,omitzero"`
	// Requests matching this policy will be redirected to this Prefix URL. Only valid
	// if action is `REDIRECT_PREFIX`.
	RedirectPrefix param.Opt[string] `json:"redirect_prefix,omitzero"`
	// Requests matching this policy will be redirected to this URL. Only valid if
	// action is `REDIRECT_TO_URL`.
	RedirectURL param.Opt[string] `json:"redirect_url,omitzero"`
	// A list of simple strings assigned to the resource.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r LoadBalancerL7PolicyReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerL7PolicyReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerL7PolicyReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action
type LoadBalancerL7PolicyReplaceParamsAction string

const (
	LoadBalancerL7PolicyReplaceParamsActionRedirectPrefix LoadBalancerL7PolicyReplaceParamsAction = "REDIRECT_PREFIX"
	LoadBalancerL7PolicyReplaceParamsActionRedirectToPool LoadBalancerL7PolicyReplaceParamsAction = "REDIRECT_TO_POOL"
	LoadBalancerL7PolicyReplaceParamsActionRedirectToURL  LoadBalancerL7PolicyReplaceParamsAction = "REDIRECT_TO_URL"
	LoadBalancerL7PolicyReplaceParamsActionReject         LoadBalancerL7PolicyReplaceParamsAction = "REJECT"
)
