// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1ProjectService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1ProjectService] method instead.
type CloudV1ProjectService struct {
	Options []option.RequestOption
}

// NewCloudV1ProjectService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1ProjectService(opts ...option.RequestOption) (r *CloudV1ProjectService) {
	r = &CloudV1ProjectService{}
	r.Options = opts
	return
}

// Create project
func (r *CloudV1ProjectService) New(ctx context.Context, body CloudV1ProjectNewParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Project
func (r *CloudV1ProjectService) Get(ctx context.Context, projectID int64, opts ...option.RequestOption) (res *Project, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/projects/%v", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Project
func (r *CloudV1ProjectService) Update(ctx context.Context, projectID int64, body CloudV1ProjectUpdateParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/projects/%v", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List projects
func (r *CloudV1ProjectService) List(ctx context.Context, query CloudV1ProjectListParams, opts ...option.RequestOption) (res *CloudV1ProjectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// All cloud resources in all regions that belong to the project will be deleted
// and will not be recoverable
func (r *CloudV1ProjectService) Delete(ctx context.Context, projectID int64, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/projects/%v", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Project struct {
	// Project ID, which is automatically generated upon creation.
	ID int64 `json:"id,required"`
	// ID associated with the client.
	ClientID int64 `json:"client_id,required"`
	// Datetime of creation, which is automatically generated.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Indicates if the project is the default one. Each client always has one default
	// project.
	IsDefault bool `json:"is_default,required"`
	// Unique project name for a client.
	Name string `json:"name,required"`
	// The state of the project.
	State string `json:"state,required"`
	// Datetime of deletion, which is automatically generated if the project is
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Description of the project.
	Description string `json:"description,nullable"`
	// ID of the Task entity responsible for handling the project's state transition.
	TaskID string      `json:"task_id,nullable"`
	JSON   projectJSON `json:"-"`
}

// projectJSON contains the JSON metadata for the struct [Project]
type projectJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	IsDefault   apijson.Field
	Name        apijson.Field
	State       apijson.Field
	DeletedAt   apijson.Field
	Description apijson.Field
	TaskID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Project) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectJSON) RawJSON() string {
	return r.raw
}

type CloudV1ProjectListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Project                      `json:"results,required"`
	JSON    cloudV1ProjectListResponseJSON `json:"-"`
}

// cloudV1ProjectListResponseJSON contains the JSON metadata for the struct
// [CloudV1ProjectListResponse]
type cloudV1ProjectListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1ProjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1ProjectListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1ProjectNewParams struct {
	// Unique project name for a client. Each client always has one "default" project.
	Name param.Field[string] `json:"name,required"`
	// ID associated with the client.
	ClientID param.Field[int64] `json:"client_id"`
	// Description of the project.
	Description param.Field[string] `json:"description"`
	// State of the project.
	State param.Field[string] `json:"state"`
}

func (r CloudV1ProjectNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1ProjectUpdateParams struct {
	// Name of the entity, following a specific format.
	Name param.Field[string] `json:"name,required"`
	// Description of the project.
	Description param.Field[string] `json:"description"`
}

func (r CloudV1ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1ProjectListParams struct {
	// Client ID filter for administrators.
	ClientID param.Field[int64] `query:"client_id"`
	// Whether to include deleted entries in the response.
	IncludeDeleted param.Field[bool] `query:"include_deleted"`
	// Name to filter the results by.
	Name param.Field[string] `query:"name"`
	// Order by field and direction. Supports multiple values.
	OrderBy param.Field[[]CloudV1ProjectListParamsOrderBy] `query:"order_by"`
}

// URLQuery serializes [CloudV1ProjectListParams]'s query parameters as
// `url.Values`.
func (r CloudV1ProjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1ProjectListParamsOrderBy string

const (
	CloudV1ProjectListParamsOrderByCreatedAtAsc  CloudV1ProjectListParamsOrderBy = "created_at.asc"
	CloudV1ProjectListParamsOrderByCreatedAtDesc CloudV1ProjectListParamsOrderBy = "created_at.desc"
	CloudV1ProjectListParamsOrderByNameAsc       CloudV1ProjectListParamsOrderBy = "name.asc"
	CloudV1ProjectListParamsOrderByNameDesc      CloudV1ProjectListParamsOrderBy = "name.desc"
)

func (r CloudV1ProjectListParamsOrderBy) IsKnown() bool {
	switch r {
	case CloudV1ProjectListParamsOrderByCreatedAtAsc, CloudV1ProjectListParamsOrderByCreatedAtDesc, CloudV1ProjectListParamsOrderByNameAsc, CloudV1ProjectListParamsOrderByNameDesc:
		return true
	}
	return false
}
