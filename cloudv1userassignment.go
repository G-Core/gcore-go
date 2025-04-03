// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1UserAssignmentService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1UserAssignmentService] method instead.
type CloudV1UserAssignmentService struct {
	Options []option.RequestOption
}

// NewCloudV1UserAssignmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1UserAssignmentService(opts ...option.RequestOption) (r *CloudV1UserAssignmentService) {
	r = &CloudV1UserAssignmentService{}
	r.Options = opts
	return
}

// Modify role assignment to existing user
func (r *CloudV1UserAssignmentService) Update(ctx context.Context, assignmentID int64, body CloudV1UserAssignmentUpdateParams, opts ...option.RequestOption) (res *HTTPResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/users/assignments/%v", assignmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List assignments
func (r *CloudV1UserAssignmentService) List(ctx context.Context, query CloudV1UserAssignmentListParams, opts ...option.RequestOption) (res *CloudV1UserAssignmentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/users/assignments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete role assignment
func (r *CloudV1UserAssignmentService) Delete(ctx context.Context, assignmentID int64, opts ...option.RequestOption) (res *HTTPResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/users/assignments/%v", assignmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Assign role to existing user
func (r *CloudV1UserAssignmentService) AssignRole(ctx context.Context, body CloudV1UserAssignmentAssignRoleParams, opts ...option.RequestOption) (res *RoleAssignment, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/users/assignments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type HTTPResponse struct {
	// True if everything is OK
	Ok bool `json:"ok,required"`
	// Http error response
	Error HTTPResponseError `json:"error,nullable"`
	JSON  httpResponseJSON  `json:"-"`
}

// httpResponseJSON contains the JSON metadata for the struct [HTTPResponse]
type httpResponseJSON struct {
	Ok          apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpResponseJSON) RawJSON() string {
	return r.raw
}

// Http error response
type HTTPResponseError struct {
	Code    int64                 `json:"code,required"`
	Message string                `json:"message,required"`
	JSON    httpResponseErrorJSON `json:"-"`
}

// httpResponseErrorJSON contains the JSON metadata for the struct
// [HTTPResponseError]
type httpResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpResponseErrorJSON) RawJSON() string {
	return r.raw
}

type RequestAssignmentParam struct {
	// User role
	Role param.Field[string] `json:"role,required"`
	// User ID
	UserID param.Field[int64] `json:"user_id,required"`
	// Client ID. Required if project_id is specified
	ClientID param.Field[int64] `json:"client_id"`
	// Project ID
	ProjectID param.Field[int64] `json:"project_id"`
}

func (r RequestAssignmentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RoleAssignment struct {
	// Assignment ID
	ID int64 `json:"id,required"`
	// User role
	Role string `json:"role,required"`
	// User ID
	UserID int64 `json:"user_id,required"`
	// Client ID. Required if project_id is specified
	ClientID int64 `json:"client_id,nullable"`
	// Project ID
	ProjectID int64              `json:"project_id,nullable"`
	JSON      roleAssignmentJSON `json:"-"`
}

// roleAssignmentJSON contains the JSON metadata for the struct [RoleAssignment]
type roleAssignmentJSON struct {
	ID          apijson.Field
	Role        apijson.Field
	UserID      apijson.Field
	ClientID    apijson.Field
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleAssignment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleAssignmentJSON) RawJSON() string {
	return r.raw
}

type CloudV1UserAssignmentListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []RoleAssignment                      `json:"results,required"`
	JSON    cloudV1UserAssignmentListResponseJSON `json:"-"`
}

// cloudV1UserAssignmentListResponseJSON contains the JSON metadata for the struct
// [CloudV1UserAssignmentListResponse]
type cloudV1UserAssignmentListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1UserAssignmentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1UserAssignmentListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1UserAssignmentUpdateParams struct {
	RequestAssignment RequestAssignmentParam `json:"request_assignment,required"`
}

func (r CloudV1UserAssignmentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RequestAssignment)
}

type CloudV1UserAssignmentListParams struct {
	// Client ID. Must be used for users without client_id in jwt
	ClientID param.Field[int64] `query:"client_id"`
	// Limit the number of returned assignments. Falls back to default if not
	// specified. Limited by max limit value
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
	// Project ID
	ProjectID param.Field[int64] `query:"project_id"`
	// User ID for filtering
	UserID param.Field[int64] `query:"user_id"`
}

// URLQuery serializes [CloudV1UserAssignmentListParams]'s query parameters as
// `url.Values`.
func (r CloudV1UserAssignmentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1UserAssignmentAssignRoleParams struct {
	RequestAssignment RequestAssignmentParam `json:"request_assignment,required"`
}

func (r CloudV1UserAssignmentAssignRoleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RequestAssignment)
}
