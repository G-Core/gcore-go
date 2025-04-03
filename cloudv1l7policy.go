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

// CloudV1L7policyService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1L7policyService] method instead.
type CloudV1L7policyService struct {
	Options []option.RequestOption
	Rules   *CloudV1L7policyRuleService
}

// NewCloudV1L7policyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1L7policyService(opts ...option.RequestOption) (r *CloudV1L7policyService) {
	r = &CloudV1L7policyService{}
	r.Options = opts
	r.Rules = NewCloudV1L7policyRuleService(opts...)
	return
}

// Create load balancer L7 policy
func (r *CloudV1L7policyService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1L7policyNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get load balancer L7 policy
func (r *CloudV1L7policyService) Get(ctx context.Context, projectID int64, regionID int64, l7policyID string, opts ...option.RequestOption) (res *L7Policy, err error) {
	opts = append(r.Options[:], opts...)
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s", projectID, regionID, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Replace load balancer L7 policy properties
func (r *CloudV1L7policyService) Update(ctx context.Context, projectID int64, regionID int64, l7policyID string, body CloudV1L7policyUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s", projectID, regionID, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List load balancer L7 policies
func (r *CloudV1L7policyService) List(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *CloudV1L7policyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete load balancer L7 policy
func (r *CloudV1L7policyService) Delete(ctx context.Context, projectID int64, regionID int64, l7policyID string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if l7policyID == "" {
		err = errors.New("missing required l7policy_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/l7policies/%v/%v/%s", projectID, regionID, l7policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// L7Policy schema
type L7Policy struct {
	// ID
	ID string `json:"id"`
	// Action
	Action L7PolicyAction `json:"action"`
	// Listener ID
	ListenerID string `json:"listener_id"`
	// Human-readable name of the policy
	Name string `json:"name"`
	// L7 policy operating status
	OperatingStatus L7PolicyOperatingStatus `json:"operating_status"`
	// The position of this policy on the listener. Positions start at 1.
	Position int64 `json:"position"`
	// Project ID
	ProjectID          int64                      `json:"project_id"`
	ProvisioningStatus L7PolicyProvisioningStatus `json:"provisioning_status"`
	// Requests matching this policy will be redirected to the specified URL or Prefix
	// URL with the HTTP response code. Valid if action is REDIRECT_TO_URL or
	// REDIRECT_PREFIX. Valid options are 301, 302, 303, 307, or 308. Default is 302.
	RedirectHTTPCode int64 `json:"redirect_http_code"`
	// Requests matching this policy will be redirected to the pool with this ID. Only
	// valid if action is REDIRECT_TO_POOL.
	RedirectPoolID string `json:"redirect_pool_id"`
	// Requests matching this policy will be redirected to this Prefix URL. Only valid
	// if action is REDIRECT_PREFIX.
	RedirectPrefix string `json:"redirect_prefix"`
	// Requests matching this policy will be redirected to this URL. Only valid if
	// action is REDIRECT_TO_URL.
	RedirectURL string `json:"redirect_url"`
	// Region name
	Region string `json:"region"`
	// Region ID
	RegionID int64 `json:"region_id"`
	// Rules. All the rules associated with a given policy are logically ANDed
	// together. A request must match all the policy’s rules to match the policy.If you
	// need to express a logical OR operation between rules, then do this by creating
	// multiple policies with the same action.
	Rules []L7Rule `json:"rules"`
	// A list of simple strings assigned to the resource.
	Tags []string `json:"tags"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string       `json:"task_id"`
	JSON   l7PolicyJSON `json:"-"`
}

// l7PolicyJSON contains the JSON metadata for the struct [L7Policy]
type l7PolicyJSON struct {
	ID                 apijson.Field
	Action             apijson.Field
	ListenerID         apijson.Field
	Name               apijson.Field
	OperatingStatus    apijson.Field
	Position           apijson.Field
	ProjectID          apijson.Field
	ProvisioningStatus apijson.Field
	RedirectHTTPCode   apijson.Field
	RedirectPoolID     apijson.Field
	RedirectPrefix     apijson.Field
	RedirectURL        apijson.Field
	Region             apijson.Field
	RegionID           apijson.Field
	Rules              apijson.Field
	Tags               apijson.Field
	TaskID             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *L7Policy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r l7PolicyJSON) RawJSON() string {
	return r.raw
}

// Action
type L7PolicyAction string

const (
	L7PolicyActionRedirectPrefix L7PolicyAction = "REDIRECT_PREFIX"
	L7PolicyActionRedirectToPool L7PolicyAction = "REDIRECT_TO_POOL"
	L7PolicyActionRedirectToURL  L7PolicyAction = "REDIRECT_TO_URL"
	L7PolicyActionReject         L7PolicyAction = "REJECT"
)

func (r L7PolicyAction) IsKnown() bool {
	switch r {
	case L7PolicyActionRedirectPrefix, L7PolicyActionRedirectToPool, L7PolicyActionRedirectToURL, L7PolicyActionReject:
		return true
	}
	return false
}

// L7 policy operating status
type L7PolicyOperatingStatus string

const (
	L7PolicyOperatingStatusDegraded  L7PolicyOperatingStatus = "DEGRADED"
	L7PolicyOperatingStatusDraining  L7PolicyOperatingStatus = "DRAINING"
	L7PolicyOperatingStatusError     L7PolicyOperatingStatus = "ERROR"
	L7PolicyOperatingStatusNoMonitor L7PolicyOperatingStatus = "NO_MONITOR"
	L7PolicyOperatingStatusOffline   L7PolicyOperatingStatus = "OFFLINE"
	L7PolicyOperatingStatusOnline    L7PolicyOperatingStatus = "ONLINE"
)

func (r L7PolicyOperatingStatus) IsKnown() bool {
	switch r {
	case L7PolicyOperatingStatusDegraded, L7PolicyOperatingStatusDraining, L7PolicyOperatingStatusError, L7PolicyOperatingStatusNoMonitor, L7PolicyOperatingStatusOffline, L7PolicyOperatingStatusOnline:
		return true
	}
	return false
}

type L7PolicyProvisioningStatus string

const (
	L7PolicyProvisioningStatusActive        L7PolicyProvisioningStatus = "ACTIVE"
	L7PolicyProvisioningStatusDeleted       L7PolicyProvisioningStatus = "DELETED"
	L7PolicyProvisioningStatusError         L7PolicyProvisioningStatus = "ERROR"
	L7PolicyProvisioningStatusPendingCreate L7PolicyProvisioningStatus = "PENDING_CREATE"
	L7PolicyProvisioningStatusPendingDelete L7PolicyProvisioningStatus = "PENDING_DELETE"
	L7PolicyProvisioningStatusPendingUpdate L7PolicyProvisioningStatus = "PENDING_UPDATE"
)

func (r L7PolicyProvisioningStatus) IsKnown() bool {
	switch r {
	case L7PolicyProvisioningStatusActive, L7PolicyProvisioningStatusDeleted, L7PolicyProvisioningStatusError, L7PolicyProvisioningStatusPendingCreate, L7PolicyProvisioningStatusPendingDelete, L7PolicyProvisioningStatusPendingUpdate:
		return true
	}
	return false
}

type CloudV1L7policyListResponse struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []L7Policy                      `json:"results"`
	JSON    cloudV1L7policyListResponseJSON `json:"-"`
}

// cloudV1L7policyListResponseJSON contains the JSON metadata for the struct
// [CloudV1L7policyListResponse]
type cloudV1L7policyListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1L7policyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1L7policyListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1L7policyNewParams struct {
	// Action
	Action param.Field[CloudV1L7policyNewParamsAction] `json:"action,required"`
	// Listener ID
	ListenerID param.Field[string] `json:"listener_id,required"`
	// Human-readable name of the policy
	Name param.Field[string] `json:"name"`
	// The position of this policy on the listener. Positions start at 1.
	Position param.Field[int64] `json:"position"`
	// Requests matching this policy will be redirected to the specified URL or Prefix
	// URL with the HTTP response code. Valid if action is REDIRECT_TO_URL or
	// REDIRECT_PREFIX. Valid options are 301, 302, 303, 307, or 308. Default is 302.
	RedirectHTTPCode param.Field[int64] `json:"redirect_http_code"`
	// Requests matching this policy will be redirected to the pool withthis ID. Only
	// valid if action is REDIRECT_TO_POOL.
	RedirectPoolID param.Field[string] `json:"redirect_pool_id"`
	// Requests matching this policy will be redirected to this Prefix URL. Only valid
	// if action is REDIRECT_PREFIX.
	RedirectPrefix param.Field[string] `json:"redirect_prefix"`
	// Requests matching this policy will be redirected to this URL. Only valid if
	// action is REDIRECT_TO_URL.
	RedirectURL param.Field[string] `json:"redirect_url"`
	// A list of simple strings assigned to the resource.
	Tags param.Field[[]string] `json:"tags"`
}

func (r CloudV1L7policyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action
type CloudV1L7policyNewParamsAction string

const (
	CloudV1L7policyNewParamsActionRedirectPrefix CloudV1L7policyNewParamsAction = "REDIRECT_PREFIX"
	CloudV1L7policyNewParamsActionRedirectToPool CloudV1L7policyNewParamsAction = "REDIRECT_TO_POOL"
	CloudV1L7policyNewParamsActionRedirectToURL  CloudV1L7policyNewParamsAction = "REDIRECT_TO_URL"
	CloudV1L7policyNewParamsActionReject         CloudV1L7policyNewParamsAction = "REJECT"
)

func (r CloudV1L7policyNewParamsAction) IsKnown() bool {
	switch r {
	case CloudV1L7policyNewParamsActionRedirectPrefix, CloudV1L7policyNewParamsActionRedirectToPool, CloudV1L7policyNewParamsActionRedirectToURL, CloudV1L7policyNewParamsActionReject:
		return true
	}
	return false
}

type CloudV1L7policyUpdateParams struct {
	// Action
	Action param.Field[CloudV1L7policyUpdateParamsAction] `json:"action,required"`
	// Human-readable name of the policy
	Name param.Field[string] `json:"name"`
	// The position of this policy on the listener. Positions start at 1.
	Position param.Field[int64] `json:"position"`
	// Requests matching this policy will be redirected to the specified URL or Prefix
	// URL with the HTTP response code. Valid if action is REDIRECT_TO_URL or
	// REDIRECT_PREFIX. Valid options are 301, 302, 303, 307, or 308. Default is 302.
	RedirectHTTPCode param.Field[int64] `json:"redirect_http_code"`
	// Requests matching this policy will be redirected to the pool with this ID. Only
	// valid if action is REDIRECT_TO_POOL.
	RedirectPoolID param.Field[string] `json:"redirect_pool_id"`
	// Requests matching this policy will be redirected to this Prefix URL. Only valid
	// if action is REDIRECT_PREFIX.
	RedirectPrefix param.Field[string] `json:"redirect_prefix"`
	// Requests matching this policy will be redirected to this URL. Only valid if
	// action is REDIRECT_TO_URL.
	RedirectURL param.Field[string] `json:"redirect_url"`
	// A list of simple strings assigned to the resource.
	Tags param.Field[[]string] `json:"tags"`
}

func (r CloudV1L7policyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action
type CloudV1L7policyUpdateParamsAction string

const (
	CloudV1L7policyUpdateParamsActionRedirectPrefix CloudV1L7policyUpdateParamsAction = "REDIRECT_PREFIX"
	CloudV1L7policyUpdateParamsActionRedirectToPool CloudV1L7policyUpdateParamsAction = "REDIRECT_TO_POOL"
	CloudV1L7policyUpdateParamsActionRedirectToURL  CloudV1L7policyUpdateParamsAction = "REDIRECT_TO_URL"
	CloudV1L7policyUpdateParamsActionReject         CloudV1L7policyUpdateParamsAction = "REJECT"
)

func (r CloudV1L7policyUpdateParamsAction) IsKnown() bool {
	switch r {
	case CloudV1L7policyUpdateParamsActionRedirectPrefix, CloudV1L7policyUpdateParamsActionRedirectToPool, CloudV1L7policyUpdateParamsActionRedirectToURL, CloudV1L7policyUpdateParamsActionReject:
		return true
	}
	return false
}
