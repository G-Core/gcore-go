// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// ProjectService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options []option.RequestOption
	tasks   TaskService
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r ProjectService) {
	r = ProjectService{}
	r.Options = opts
	r.tasks = NewTaskService(opts...)
	return
}

// Create a new project for a client. Project management must be enabled to perform
// this operation.
func (r *ProjectService) New(ctx context.Context, body ProjectNewParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cloud/v1/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update project name and description.
func (r *ProjectService) Update(ctx context.Context, params ProjectUpdateParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/projects/%v", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of projects for a client. Results can be filtered by name and
// ordered by various fields.
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Project], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cloud/v1/projects"
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

// Retrieve a list of projects for a client. Results can be filtered by name and
// ordered by various fields.
func (r *ProjectService) ListAutoPaging(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Project] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a project and all its associated cloud resources across all regions. This
// operation is irreversible and cannot be undone. Default projects cannot be
// deleted.
func (r *ProjectService) Delete(ctx context.Context, body ProjectDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/projects/%v", body.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// DeleteAndPoll deletes a project and polls for completion of the first task. Use the
// [TaskService.Poll] method if you need to poll for all tasks.
func (r *ProjectService) DeleteAndPoll(ctx context.Context, body ProjectDeleteParams, opts ...option.RequestOption) (err error) {
	resource, err := r.Delete(ctx, body, opts...)
	if err != nil {
		return err
	}

	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	return err
}

// Retrieve detailed information about a specific project.
func (r *ProjectService) Get(ctx context.Context, query ProjectGetParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/projects/%v", query.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Project struct {
	// Project ID, which is automatically generated upon creation.
	ID int64 `json:"id,required"`
	// ID associated with the client.
	ClientID int64 `json:"client_id,required"`
	// Datetime of creation, which is automatically generated.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Datetime of deletion, which is automatically generated if the project is
	// deleted.
	DeletedAt time.Time `json:"deleted_at,required" format:"date-time"`
	// Description of the project.
	Description string `json:"description,required"`
	// Indicates if the project is the default one. Each client always has one default
	// project.
	IsDefault bool `json:"is_default,required"`
	// Unique project name for a client.
	Name string `json:"name,required"`
	// The state of the project.
	State string `json:"state,required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ClientID    respjson.Field
		CreatedAt   respjson.Field
		DeletedAt   respjson.Field
		Description respjson.Field
		IsDefault   respjson.Field
		Name        respjson.Field
		State       respjson.Field
		TaskID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Project) RawJSON() string { return r.JSON.raw }
func (r *Project) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectNewParams struct {
	// Unique project name for a client. Each client always has one "default" project.
	Name string `json:"name,required"`
	// Description of the project.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ProjectNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectUpdateParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Name of the entity, following a specific format.
	Name string `json:"name,required"`
	// Description of the project.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListParams struct {
	// Client ID filter for administrators.
	ClientID param.Opt[int64] `query:"client_id,omitzero" json:"-"`
	// Whether to include deleted projects in the response.
	IncludeDeleted param.Opt[bool] `query:"include_deleted,omitzero" json:"-"`
	// Limit value is used to limit the number of records in the result
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Name to filter the results by.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Order by field and direction.
	//
	// Any of "created_at.asc", "created_at.desc", "name.asc", "name.desc".
	OrderBy ProjectListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectListParams]'s query parameters as `url.Values`.
func (r ProjectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Order by field and direction.
type ProjectListParamsOrderBy string

const (
	ProjectListParamsOrderByCreatedAtAsc  ProjectListParamsOrderBy = "created_at.asc"
	ProjectListParamsOrderByCreatedAtDesc ProjectListParamsOrderBy = "created_at.desc"
	ProjectListParamsOrderByNameAsc       ProjectListParamsOrderBy = "name.asc"
	ProjectListParamsOrderByNameDesc      ProjectListParamsOrderBy = "name.desc"
)

type ProjectDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

type ProjectGetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}
