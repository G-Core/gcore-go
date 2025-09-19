// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"github.com/G-Core/gcore-go/packages/respjson"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// LoadBalancerListenerService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoadBalancerListenerService] method instead.
type LoadBalancerListenerService struct {
	Options []option.RequestOption
	tasks   TaskService
}

// NewLoadBalancerListenerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerListenerService(opts ...option.RequestOption) (r LoadBalancerListenerService) {
	r = LoadBalancerListenerService{}
	r.Options = opts
	r.tasks = NewTaskService(opts...)
	return
}

// Create load balancer listener
func (r *LoadBalancerListenerService) New(ctx context.Context, params LoadBalancerListenerNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/lblisteners/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update load balancer listener
func (r *LoadBalancerListenerService) Update(ctx context.Context, listenerID string, params LoadBalancerListenerUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if listenerID == "" {
		err = errors.New("missing required listener_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/lblisteners/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, listenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List load balancer listeners
func (r *LoadBalancerListenerService) List(ctx context.Context, params LoadBalancerListenerListParams, opts ...option.RequestOption) (res *LoadBalancerListenerList, err error) {
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
	path := fmt.Sprintf("cloud/v1/lblisteners/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete load balancer listener
func (r *LoadBalancerListenerService) Delete(ctx context.Context, listenerID string, body LoadBalancerListenerDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if listenerID == "" {
		err = errors.New("missing required listener_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lblisteners/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, listenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get load balancer listener
func (r *LoadBalancerListenerService) Get(ctx context.Context, listenerID string, params LoadBalancerListenerGetParams, opts ...option.RequestOption) (res *LoadBalancerListenerDetail, err error) {
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
	if listenerID == "" {
		err = errors.New("missing required listener_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lblisteners/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, listenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// NewAndPoll creates a new listener and polls for completion
func (r *LoadBalancerListenerService) NewAndPoll(ctx context.Context, params LoadBalancerListenerNewParams, opts ...option.RequestOption) (v *LoadBalancerListenerDetail, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams LoadBalancerListenerGetParams
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

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Listeners) != 1 {
		return nil, errors.New("expected exactly one listener to be created in a task")
	}
	resourceID := task.CreatedResources.Listeners[0]

	return r.Get(ctx, resourceID, getParams, opts...)
}

// DeleteAndPoll deletes a listener and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *LoadBalancerListenerService) DeleteAndPoll(ctx context.Context, listenerID string, params LoadBalancerListenerDeleteParams, opts ...option.RequestOption) error {
	resource, err := r.Delete(ctx, listenerID, params, opts...)
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

// UpdateAndPoll updates a listener and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *LoadBalancerListenerService) UpdateAndPoll(ctx context.Context, listenerID string, params LoadBalancerListenerUpdateParams, opts ...option.RequestOption) (v *LoadBalancerListenerDetail, err error) {
	resource, err := r.Update(ctx, listenerID, params, opts...)
	if err != nil {
		return
	}

	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams LoadBalancerListenerGetParams
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

	return r.Get(ctx, listenerID, getParams, opts...)
}

type LoadBalancerListenerListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []LoadBalancerListenerDetail `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerListenerListResponse) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerListenerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerListenerNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Load balancer ID
	LoadbalancerID string `json:"loadbalancer_id,required" format:"uuid4"`
	// Load balancer listener name
	Name string `json:"name,required"`
	// Load balancer listener protocol
	//
	// Any of "HTTP", "HTTPS", "PROMETHEUS", "TCP", "TERMINATED_HTTPS", "UDP".
	Protocol LbListenerProtocol `json:"protocol,omitzero,required"`
	// Protocol port
	ProtocolPort int64 `json:"protocol_port,required"`
	// Frontend client inactivity timeout in milliseconds
	TimeoutClientData param.Opt[int64] `json:"timeout_client_data,omitzero"`
	// Backend member connection timeout in milliseconds
	TimeoutMemberConnect param.Opt[int64] `json:"timeout_member_connect,omitzero"`
	// Backend member inactivity timeout in milliseconds
	TimeoutMemberData param.Opt[int64] `json:"timeout_member_data,omitzero"`
	// Limit of the simultaneous connections
	ConnectionLimit param.Opt[int64] `json:"connection_limit,omitzero"`
	// Add headers X-Forwarded-For, X-Forwarded-Port, X-Forwarded-Proto to requests.
	// Only used with HTTP or `TERMINATED_HTTPS` protocols.
	InsertXForwarded param.Opt[bool] `json:"insert_x_forwarded,omitzero"`
	// ID of the secret where PKCS12 file is stored for `TERMINATED_HTTPS` or
	// PROMETHEUS listener
	SecretID param.Opt[string] `json:"secret_id,omitzero"`
	// Network CIDRs from which service will be accessible
	AllowedCidrs []string `json:"allowed_cidrs,omitzero" format:"ipvanynetwork"`
	// List of secrets IDs containing PKCS12 format certificate/key bundles for
	// `TERMINATED_HTTPS` or PROMETHEUS listeners
	SniSecretID []string `json:"sni_secret_id,omitzero" format:"uuid4"`
	// Load balancer listener list of username and encrypted password items
	UserList []LoadBalancerListenerNewParamsUserList `json:"user_list,omitzero"`
	paramObj
}

func (r LoadBalancerListenerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerListenerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerListenerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EncryptedPassword, Username are required.
type LoadBalancerListenerNewParamsUserList struct {
	// Encrypted password to auth via Basic Authentication
	EncryptedPassword string `json:"encrypted_password,required"`
	// Username to auth via Basic Authentication
	Username string `json:"username,required"`
	paramObj
}

func (r LoadBalancerListenerNewParamsUserList) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerListenerNewParamsUserList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerListenerNewParamsUserList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerListenerUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// ID of the secret where PKCS12 file is stored for `TERMINATED_HTTPS` or
	// PROMETHEUS load balancer
	SecretID param.Opt[string] `json:"secret_id,omitzero" format:"uuid4"`
	// Frontend client inactivity timeout in milliseconds
	TimeoutClientData param.Opt[int64] `json:"timeout_client_data,omitzero"`
	// Backend member connection timeout in milliseconds
	TimeoutMemberConnect param.Opt[int64] `json:"timeout_member_connect,omitzero"`
	// Backend member inactivity timeout in milliseconds
	TimeoutMemberData param.Opt[int64] `json:"timeout_member_data,omitzero"`
	// Limit of simultaneous connections
	ConnectionLimit param.Opt[int64] `json:"connection_limit,omitzero"`
	// Load balancer listener name
	Name param.Opt[string] `json:"name,omitzero"`
	// Network CIDRs from which service will be accessible
	AllowedCidrs []string `json:"allowed_cidrs,omitzero" format:"ipvanynetwork"`
	// List of secret's ID containing PKCS12 format certificate/key bundfles for
	// `TERMINATED_HTTPS` or PROMETHEUS listeners
	SniSecretID []string `json:"sni_secret_id,omitzero" format:"uuid4"`
	// Load balancer listener users list
	UserList []LoadBalancerListenerUpdateParamsUserList `json:"user_list,omitzero"`
	paramObj
}

func (r LoadBalancerListenerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerListenerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerListenerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EncryptedPassword, Username are required.
type LoadBalancerListenerUpdateParamsUserList struct {
	// Encrypted password to auth via Basic Authentication
	EncryptedPassword string `json:"encrypted_password,required"`
	// Username to auth via Basic Authentication
	Username string `json:"username,required"`
	paramObj
}

func (r LoadBalancerListenerUpdateParamsUserList) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerListenerUpdateParamsUserList
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoadBalancerListenerUpdateParamsUserList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerListenerListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Load Balancer ID
	LoadbalancerID param.Opt[string] `query:"loadbalancer_id,omitzero" format:"uuid4" json:"-"`
	// Show stats
	ShowStats param.Opt[bool] `query:"show_stats,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LoadBalancerListenerListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerListenerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LoadBalancerListenerDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type LoadBalancerListenerGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Show stats
	ShowStats param.Opt[bool] `query:"show_stats,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LoadBalancerListenerGetParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerListenerGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
