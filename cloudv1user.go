// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1UserService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1UserService] method instead.
type CloudV1UserService struct {
	Options     []option.RequestOption
	Assignments *CloudV1UserAssignmentService
}

// NewCloudV1UserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1UserService(opts ...option.RequestOption) (r *CloudV1UserService) {
	r = &CloudV1UserService{}
	r.Options = opts
	r.Assignments = NewCloudV1UserAssignmentService(opts...)
	return
}

// List client's users
func (r *CloudV1UserService) List(ctx context.Context, query CloudV1UserListParams, opts ...option.RequestOption) (res *CloudV1UserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List available roles
func (r *CloudV1UserService) ListRoles(ctx context.Context, query CloudV1UserListRolesParams, opts ...option.RequestOption) (res *CloudV1UserListRolesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/users/roles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CloudV1UserListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1UserListResponseResult `json:"results,required"`
	JSON    cloudV1UserListResponseJSON     `json:"-"`
}

// cloudV1UserListResponseJSON contains the JSON metadata for the struct
// [CloudV1UserListResponse]
type cloudV1UserListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1UserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1UserListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1UserListResponseResult struct {
	// User ID
	ID int64 `json:"id,required"`
	// User email
	Email string `json:"email,required" format:"email"`
	// True is the user is active
	Activated bool `json:"activated,nullable"`
	// True if the user has admin role
	IsAdmin bool `json:"is_admin,nullable"`
	// User name
	Name string                            `json:"name"`
	JSON cloudV1UserListResponseResultJSON `json:"-"`
}

// cloudV1UserListResponseResultJSON contains the JSON metadata for the struct
// [CloudV1UserListResponseResult]
type cloudV1UserListResponseResultJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Activated   apijson.Field
	IsAdmin     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1UserListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1UserListResponseResultJSON) RawJSON() string {
	return r.raw
}

type CloudV1UserListRolesResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1UserListRolesResponseResult `json:"results,required"`
	JSON    cloudV1UserListRolesResponseJSON     `json:"-"`
}

// cloudV1UserListRolesResponseJSON contains the JSON metadata for the struct
// [CloudV1UserListRolesResponse]
type cloudV1UserListRolesResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1UserListRolesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1UserListRolesResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1UserListRolesResponseResult struct {
	// Available role
	Role string `json:"role,required"`
	// Role's scope
	Scope string                                 `json:"scope,required"`
	JSON  cloudV1UserListRolesResponseResultJSON `json:"-"`
}

// cloudV1UserListRolesResponseResultJSON contains the JSON metadata for the struct
// [CloudV1UserListRolesResponseResult]
type cloudV1UserListRolesResponseResultJSON struct {
	Role        apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1UserListRolesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1UserListRolesResponseResultJSON) RawJSON() string {
	return r.raw
}

type CloudV1UserListParams struct {
	// Client ID. Must be used for users without client_id in jwt
	ClientID param.Field[int64] `query:"client_id"`
	// Limit the number of returned users. Falls back to default if not specified.
	// Limited by max limit value
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
	// User ID. Specific user search
	UserID param.Field[int64] `query:"user_id"`
}

// URLQuery serializes [CloudV1UserListParams]'s query parameters as `url.Values`.
func (r CloudV1UserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1UserListRolesParams struct {
	// Client ID. Must be used for users without client_id in jwt
	ClientID param.Field[int64] `query:"client_id"`
	// Project ID
	ProjectID param.Field[int64] `query:"project_id"`
}

// URLQuery serializes [CloudV1UserListRolesParams]'s query parameters as
// `url.Values`.
func (r CloudV1UserListRolesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
