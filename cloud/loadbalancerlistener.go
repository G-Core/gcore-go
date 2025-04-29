// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
)

// LoadBalancerListenerService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoadBalancerListenerService] method instead.
type LoadBalancerListenerService struct {
	Options []option.RequestOption
}

// NewLoadBalancerListenerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerListenerService(opts ...option.RequestOption) (r LoadBalancerListenerService) {
	r = LoadBalancerListenerService{}
	r.Options = opts
	return
}

// Create load balancer listener
func (r *LoadBalancerListenerService) New(ctx context.Context, params LoadBalancerListenerNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/lblisteners/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update listener
func (r *LoadBalancerListenerService) Update(ctx context.Context, listenerID string, params LoadBalancerListenerUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if listenerID == "" {
		err = errors.New("missing required listener_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/lblisteners/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, listenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List load balancer listeners
func (r *LoadBalancerListenerService) List(ctx context.Context, params LoadBalancerListenerListParams, opts ...option.RequestOption) (res *LbListenerList, err error) {
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
	path := fmt.Sprintf("cloud/v1/lblisteners/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete load balancer listener
func (r *LoadBalancerListenerService) Delete(ctx context.Context, listenerID string, body LoadBalancerListenerDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if listenerID == "" {
		err = errors.New("missing required listener_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lblisteners/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, listenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get listener
func (r *LoadBalancerListenerService) Get(ctx context.Context, listenerID string, params LoadBalancerListenerGetParams, opts ...option.RequestOption) (res *LbListener, err error) {
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
	if listenerID == "" {
		err = errors.New("missing required listener_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lblisteners/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, listenerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type LoadBalancerListenerNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Flblisteners%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/lblisteners/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flblisteners%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/lblisteners/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateLbListenerSerializer/properties/loadbalancer_id'
	// "$.components.schemas.CreateLbListenerSerializer.properties.loadbalancer_id"
	LoadbalancerID string `json:"loadbalancer_id,required" format:"uuid4"`
	// '#/components/schemas/CreateLbListenerSerializer/properties/name'
	// "$.components.schemas.CreateLbListenerSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateLbListenerSerializer/properties/protocol'
	// "$.components.schemas.CreateLbListenerSerializer.properties.protocol"
	//
	// Any of "HTTP", "HTTPS", "PROMETHEUS", "TCP", "TERMINATED_HTTPS", "UDP".
	Protocol LbListenerProtocol `json:"protocol,omitzero,required"`
	// '#/components/schemas/CreateLbListenerSerializer/properties/protocol_port'
	// "$.components.schemas.CreateLbListenerSerializer.properties.protocol_port"
	ProtocolPort int64 `json:"protocol_port,required"`
	// '#/components/schemas/CreateLbListenerSerializer/properties/timeout_client_data/anyOf/0'
	// "$.components.schemas.CreateLbListenerSerializer.properties.timeout_client_data.anyOf[0]"
	TimeoutClientData param.Opt[int64] `json:"timeout_client_data,omitzero"`
	// '#/components/schemas/CreateLbListenerSerializer/properties/timeout_member_connect/anyOf/0'
	// "$.components.schemas.CreateLbListenerSerializer.properties.timeout_member_connect.anyOf[0]"
	TimeoutMemberConnect param.Opt[int64] `json:"timeout_member_connect,omitzero"`
	// '#/components/schemas/CreateLbListenerSerializer/properties/timeout_member_data/anyOf/0'
	// "$.components.schemas.CreateLbListenerSerializer.properties.timeout_member_data.anyOf[0]"
	TimeoutMemberData param.Opt[int64] `json:"timeout_member_data,omitzero"`
	// '#/components/schemas/CreateLbListenerSerializer/properties/connection_limit'
	// "$.components.schemas.CreateLbListenerSerializer.properties.connection_limit"
	ConnectionLimit param.Opt[int64] `json:"connection_limit,omitzero"`
	// '#/components/schemas/CreateLbListenerSerializer/properties/insert_x_forwarded'
	// "$.components.schemas.CreateLbListenerSerializer.properties.insert_x_forwarded"
	InsertXForwarded param.Opt[bool] `json:"insert_x_forwarded,omitzero"`
	// '#/components/schemas/CreateLbListenerSerializer/properties/secret_id/anyOf/0'
	// "$.components.schemas.CreateLbListenerSerializer.properties.secret_id.anyOf[0]"
	SecretID param.Opt[string] `json:"secret_id,omitzero"`
	// '#/components/schemas/CreateLbListenerSerializer/properties/allowed_cidrs/anyOf/0'
	// "$.components.schemas.CreateLbListenerSerializer.properties.allowed_cidrs.anyOf[0]"
	AllowedCidrs []string `json:"allowed_cidrs,omitzero" format:"ipvanynetwork"`
	// '#/components/schemas/CreateLbListenerSerializer/properties/sni_secret_id'
	// "$.components.schemas.CreateLbListenerSerializer.properties.sni_secret_id"
	SniSecretID []string `json:"sni_secret_id,omitzero" format:"uuid4"`
	// '#/components/schemas/CreateLbListenerSerializer/properties/user_list'
	// "$.components.schemas.CreateLbListenerSerializer.properties.user_list"
	UserList []LoadBalancerListenerNewParamsUserList `json:"user_list,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerListenerNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LoadBalancerListenerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerListenerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateLbListenerSerializer/properties/user_list/items'
// "$.components.schemas.CreateLbListenerSerializer.properties.user_list.items"
//
// The properties EncryptedPassword, Username are required.
type LoadBalancerListenerNewParamsUserList struct {
	// '#/components/schemas/UserListItem/properties/encrypted_password'
	// "$.components.schemas.UserListItem.properties.encrypted_password"
	EncryptedPassword string `json:"encrypted_password,required"`
	// '#/components/schemas/UserListItem/properties/username'
	// "$.components.schemas.UserListItem.properties.username"
	Username string `json:"username,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerListenerNewParamsUserList) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerListenerNewParamsUserList) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerListenerNewParamsUserList
	return param.MarshalObject(r, (*shadow)(&r))
}

type LoadBalancerListenerUpdateParams struct {
	// '#/paths/%2Fcloud%2Fv2%2Flblisteners%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Blistener_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v2/lblisteners/{project_id}/{region_id}/{listener_id}'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Flblisteners%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Blistener_id%7D/patch/parameters/1/schema'
	// "$.paths['/cloud/v2/lblisteners/{project_id}/{region_id}/{listener_id}'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/PatchLbListenerSerializer/properties/secret_id/anyOf/0'
	// "$.components.schemas.PatchLbListenerSerializer.properties.secret_id.anyOf[0]"
	SecretID param.Opt[string] `json:"secret_id,omitzero" format:"uuid4"`
	// '#/components/schemas/PatchLbListenerSerializer/properties/timeout_client_data/anyOf/0'
	// "$.components.schemas.PatchLbListenerSerializer.properties.timeout_client_data.anyOf[0]"
	TimeoutClientData param.Opt[int64] `json:"timeout_client_data,omitzero"`
	// '#/components/schemas/PatchLbListenerSerializer/properties/timeout_member_connect/anyOf/0'
	// "$.components.schemas.PatchLbListenerSerializer.properties.timeout_member_connect.anyOf[0]"
	TimeoutMemberConnect param.Opt[int64] `json:"timeout_member_connect,omitzero"`
	// '#/components/schemas/PatchLbListenerSerializer/properties/timeout_member_data/anyOf/0'
	// "$.components.schemas.PatchLbListenerSerializer.properties.timeout_member_data.anyOf[0]"
	TimeoutMemberData param.Opt[int64] `json:"timeout_member_data,omitzero"`
	// '#/components/schemas/PatchLbListenerSerializer/properties/connection_limit'
	// "$.components.schemas.PatchLbListenerSerializer.properties.connection_limit"
	ConnectionLimit param.Opt[int64] `json:"connection_limit,omitzero"`
	// '#/components/schemas/PatchLbListenerSerializer/properties/name'
	// "$.components.schemas.PatchLbListenerSerializer.properties.name"
	Name param.Opt[string] `json:"name,omitzero"`
	// '#/components/schemas/PatchLbListenerSerializer/properties/allowed_cidrs/anyOf/0'
	// "$.components.schemas.PatchLbListenerSerializer.properties.allowed_cidrs.anyOf[0]"
	AllowedCidrs []string `json:"allowed_cidrs,omitzero" format:"ipvanynetwork"`
	// '#/components/schemas/PatchLbListenerSerializer/properties/sni_secret_id/anyOf/0'
	// "$.components.schemas.PatchLbListenerSerializer.properties.sni_secret_id.anyOf[0]"
	SniSecretID []string `json:"sni_secret_id,omitzero" format:"uuid4"`
	// '#/components/schemas/PatchLbListenerSerializer/properties/user_list/anyOf/0'
	// "$.components.schemas.PatchLbListenerSerializer.properties.user_list.anyOf[0]"
	UserList []LoadBalancerListenerUpdateParamsUserList `json:"user_list,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerListenerUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r LoadBalancerListenerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerListenerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/PatchLbListenerSerializer/properties/user_list/anyOf/0/items'
// "$.components.schemas.PatchLbListenerSerializer.properties.user_list.anyOf[0].items"
//
// The properties EncryptedPassword, Username are required.
type LoadBalancerListenerUpdateParamsUserList struct {
	// '#/components/schemas/UserListItem/properties/encrypted_password'
	// "$.components.schemas.UserListItem.properties.encrypted_password"
	EncryptedPassword string `json:"encrypted_password,required"`
	// '#/components/schemas/UserListItem/properties/username'
	// "$.components.schemas.UserListItem.properties.username"
	Username string `json:"username,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerListenerUpdateParamsUserList) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r LoadBalancerListenerUpdateParamsUserList) MarshalJSON() (data []byte, err error) {
	type shadow LoadBalancerListenerUpdateParamsUserList
	return param.MarshalObject(r, (*shadow)(&r))
}

type LoadBalancerListenerListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Flblisteners%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/lblisteners/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flblisteners%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/lblisteners/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flblisteners%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/lblisteners/{project_id}/{region_id}'].get.parameters[2]"
	LoadbalancerID param.Opt[string] `query:"loadbalancer_id,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flblisteners%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/lblisteners/{project_id}/{region_id}'].get.parameters[3]"
	ShowStats param.Opt[bool] `query:"show_stats,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerListenerListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LoadBalancerListenerListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerListenerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LoadBalancerListenerDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Flblisteners%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Blistener_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/lblisteners/{project_id}/{region_id}/{listener_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flblisteners%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Blistener_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/lblisteners/{project_id}/{region_id}/{listener_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerListenerDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type LoadBalancerListenerGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Flblisteners%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Blistener_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/lblisteners/{project_id}/{region_id}/{listener_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flblisteners%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Blistener_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/lblisteners/{project_id}/{region_id}/{listener_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Flblisteners%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Blistener_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/lblisteners/{project_id}/{region_id}/{listener_id}'].get.parameters[3]"
	ShowStats param.Opt[bool] `query:"show_stats,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LoadBalancerListenerGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [LoadBalancerListenerGetParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerListenerGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
