// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
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

// ProjectService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options []option.RequestOption
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r ProjectService) {
	r = ProjectService{}
	r.Options = opts
	return
}

// Create project
func (r *ProjectService) New(ctx context.Context, body ProjectNewParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List projects
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Project], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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

// List projects
func (r *ProjectService) ListAutoPaging(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Project] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// All cloud resources in all regions that belong to the project will be deleted
// and will not be recoverable
func (r *ProjectService) Delete(ctx context.Context, body ProjectDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/projects/%v", body.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get Project
func (r *ProjectService) Get(ctx context.Context, query ProjectGetParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/projects/%v", query.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Project
func (r *ProjectService) Replace(ctx context.Context, params ProjectReplaceParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/projects/%v", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// '#/components/schemas/ProjectSerializer'
// "$.components.schemas.ProjectSerializer"
type Project struct {
	// '#/components/schemas/ProjectSerializer/properties/id'
	// "$.components.schemas.ProjectSerializer.properties.id"
	ID int64 `json:"id,required"`
	// '#/components/schemas/ProjectSerializer/properties/client_id'
	// "$.components.schemas.ProjectSerializer.properties.client_id"
	ClientID int64 `json:"client_id,required"`
	// '#/components/schemas/ProjectSerializer/properties/created_at'
	// "$.components.schemas.ProjectSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/ProjectSerializer/properties/is_default'
	// "$.components.schemas.ProjectSerializer.properties.is_default"
	IsDefault bool `json:"is_default,required"`
	// '#/components/schemas/ProjectSerializer/properties/name'
	// "$.components.schemas.ProjectSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/ProjectSerializer/properties/state'
	// "$.components.schemas.ProjectSerializer.properties.state"
	State string `json:"state,required"`
	// '#/components/schemas/ProjectSerializer/properties/deleted_at/anyOf/0'
	// "$.components.schemas.ProjectSerializer.properties.deleted_at.anyOf[0]"
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// '#/components/schemas/ProjectSerializer/properties/description/anyOf/0'
	// "$.components.schemas.ProjectSerializer.properties.description.anyOf[0]"
	Description string `json:"description,nullable"`
	// '#/components/schemas/ProjectSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.ProjectSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		ClientID    resp.Field
		CreatedAt   resp.Field
		IsDefault   resp.Field
		Name        resp.Field
		State       resp.Field
		DeletedAt   resp.Field
		Description resp.Field
		TaskID      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Project) RawJSON() string { return r.JSON.raw }
func (r *Project) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectNewParams struct {
	// '#/components/schemas/CreateProjectSerializer/properties/name'
	// "$.components.schemas.CreateProjectSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateProjectSerializer/properties/client_id/anyOf/0'
	// "$.components.schemas.CreateProjectSerializer.properties.client_id.anyOf[0]"
	ClientID param.Opt[int64] `json:"client_id,omitzero"`
	// '#/components/schemas/CreateProjectSerializer/properties/description/anyOf/0'
	// "$.components.schemas.CreateProjectSerializer.properties.description.anyOf[0]"
	Description param.Opt[string] `json:"description,omitzero"`
	// '#/components/schemas/CreateProjectSerializer/properties/state/anyOf/0'
	// "$.components.schemas.CreateProjectSerializer.properties.state.anyOf[0]"
	State param.Opt[string] `json:"state,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ProjectNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type ProjectListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fprojects/get/parameters/0'
	// "$.paths['/cloud/v1/projects'].get.parameters[0]"
	ClientID param.Opt[int64] `query:"client_id,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fprojects/get/parameters/1'
	// "$.paths['/cloud/v1/projects'].get.parameters[1]"
	IncludeDeleted param.Opt[bool] `query:"include_deleted,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fprojects/get/parameters/2'
	// "$.paths['/cloud/v1/projects'].get.parameters[2]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fprojects/get/parameters/3'
	// "$.paths['/cloud/v1/projects'].get.parameters[3]"
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fprojects/get/parameters/4'
	// "$.paths['/cloud/v1/projects'].get.parameters[4]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fprojects/get/parameters/5'
	// "$.paths['/cloud/v1/projects'].get.parameters[5]"
	//
	// Any of "created_at.asc", "created_at.desc", "name.asc", "name.desc".
	OrderBy ProjectListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ProjectListParams]'s query parameters as `url.Values`.
func (r ProjectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// '#/paths/%2Fcloud%2Fv1%2Fprojects/get/parameters/5'
// "$.paths['/cloud/v1/projects'].get.parameters[5]"
type ProjectListParamsOrderBy string

const (
	ProjectListParamsOrderByCreatedAtAsc  ProjectListParamsOrderBy = "created_at.asc"
	ProjectListParamsOrderByCreatedAtDesc ProjectListParamsOrderBy = "created_at.desc"
	ProjectListParamsOrderByNameAsc       ProjectListParamsOrderBy = "name.asc"
	ProjectListParamsOrderByNameDesc      ProjectListParamsOrderBy = "name.desc"
)

type ProjectDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fprojects%2F%7Bproject_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/projects/{project_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type ProjectGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fprojects%2F%7Bproject_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/projects/{project_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type ProjectReplaceParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fprojects%2F%7Bproject_id%7D/put/parameters/0/schema'
	// "$.paths['/cloud/v1/projects/{project_id}'].put.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/components/schemas/NameDescriptionSerializer/properties/name'
	// "$.components.schemas.NameDescriptionSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/NameDescriptionSerializer/properties/description/anyOf/0'
	// "$.components.schemas.NameDescriptionSerializer.properties.description.anyOf[0]"
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ProjectReplaceParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r ProjectReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
