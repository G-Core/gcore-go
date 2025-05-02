// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// UserRoleAssignmentService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserRoleAssignmentService] method instead.
type UserRoleAssignmentService struct {
	Options []option.RequestOption
}

// NewUserRoleAssignmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserRoleAssignmentService(opts ...option.RequestOption) (r UserRoleAssignmentService) {
	r = UserRoleAssignmentService{}
	r.Options = opts
	return
}

// Assign role to existing user
func (r *UserRoleAssignmentService) New(ctx context.Context, body UserRoleAssignmentNewParams, opts ...option.RequestOption) (res *RoleAssignment, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/users/assignments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Modify role assignment to existing user
func (r *UserRoleAssignmentService) Update(ctx context.Context, assignmentID int64, body UserRoleAssignmentUpdateParams, opts ...option.RequestOption) (res *RoleAssignmentUpdateDelete, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/users/assignments/%v", assignmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List assignments
func (r *UserRoleAssignmentService) List(ctx context.Context, query UserRoleAssignmentListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[RoleAssignment], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cloud/v1/users/assignments"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List assignments
func (r *UserRoleAssignmentService) ListAutoPaging(ctx context.Context, query UserRoleAssignmentListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[RoleAssignment] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Delete role assignment
func (r *UserRoleAssignmentService) Delete(ctx context.Context, assignmentID int64, opts ...option.RequestOption) (res *RoleAssignmentUpdateDelete, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/users/assignments/%v", assignmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type RoleAssignment struct {
	// Assignment ID
	ID         int64 `json:"id,required"`
	AssignedBy int64 `json:"assigned_by,required"`
	// Client ID
	ClientID int64 `json:"client_id,required"`
	// Created timestamp
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// User role
	Role string `json:"role,required"`
	// Updated timestamp
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// User ID
	UserID int64 `json:"user_id,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		AssignedBy  resp.Field
		ClientID    resp.Field
		CreatedAt   resp.Field
		ProjectID   resp.Field
		Role        resp.Field
		UpdatedAt   resp.Field
		UserID      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoleAssignment) RawJSON() string { return r.JSON.raw }
func (r *RoleAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RoleAssignmentUpdateDelete struct {
	// Assignment ID
	AssignmentID int64 `json:"assignment_id,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AssignmentID resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoleAssignmentUpdateDelete) RawJSON() string { return r.JSON.raw }
func (r *RoleAssignmentUpdateDelete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRoleAssignmentNewParams struct {
	// User role
	Role string `json:"role,required"`
	// User ID
	UserID int64 `json:"user_id,required"`
	// Client ID. Required if project_id is specified
	ClientID param.Opt[int64] `json:"client_id,omitzero"`
	// Project ID
	ProjectID param.Opt[int64] `json:"project_id,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f UserRoleAssignmentNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r UserRoleAssignmentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UserRoleAssignmentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type UserRoleAssignmentUpdateParams struct {
	// User role
	Role string `json:"role,required"`
	// User ID
	UserID int64 `json:"user_id,required"`
	// Client ID. Required if project_id is specified
	ClientID param.Opt[int64] `json:"client_id,omitzero"`
	// Project ID
	ProjectID param.Opt[int64] `json:"project_id,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f UserRoleAssignmentUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r UserRoleAssignmentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UserRoleAssignmentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type UserRoleAssignmentListParams struct {
	// Limit the number of returned items. Falls back to default of 1000 if not
	// specified. Limited by max limit value of 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Project ID
	ProjectID param.Opt[int64] `query:"project_id,omitzero" json:"-"`
	// User ID for filtering
	UserID param.Opt[int64] `query:"user_id,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f UserRoleAssignmentListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [UserRoleAssignmentListParams]'s query parameters as
// `url.Values`.
func (r UserRoleAssignmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
