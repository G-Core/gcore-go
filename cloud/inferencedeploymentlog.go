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
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
)

// InferenceDeploymentLogService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceDeploymentLogService] method instead.
type InferenceDeploymentLogService struct {
	Options []option.RequestOption
}

// NewInferenceDeploymentLogService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInferenceDeploymentLogService(opts ...option.RequestOption) (r InferenceDeploymentLogService) {
	r = InferenceDeploymentLogService{}
	r.Options = opts
	return
}

// Get inference deployment logs
func (r *InferenceDeploymentLogService) List(ctx context.Context, deploymentName string, params InferenceDeploymentLogListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[InferenceLog], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if deploymentName == "" {
		err = errors.New("missing required deployment_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments/%s/logs", params.ProjectID.Value, deploymentName)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Get inference deployment logs
func (r *InferenceDeploymentLogService) ListAutoPaging(ctx context.Context, deploymentName string, params InferenceDeploymentLogListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[InferenceLog] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, deploymentName, params, opts...))
}

type InferenceDeploymentLogListParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments%2F%7Bdeployment_name%7D%2Flogs/get/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments/{deployment_name}/logs'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments%2F%7Bdeployment_name%7D%2Flogs/get/parameters/5/schema/anyOf/0'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments/{deployment_name}/logs'].get.parameters[5].schema.anyOf[0]"
	RegionID param.Opt[int64] `query:"region_id,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments%2F%7Bdeployment_name%7D%2Flogs/get/parameters/2'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments/{deployment_name}/logs'].get.parameters[2]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments%2F%7Bdeployment_name%7D%2Flogs/get/parameters/3'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments/{deployment_name}/logs'].get.parameters[3]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments%2F%7Bdeployment_name%7D%2Flogs/get/parameters/4'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments/{deployment_name}/logs'].get.parameters[4]"
	//
	// Any of "time.asc", "time.desc".
	OrderBy InferenceDeploymentLogListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentLogListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [InferenceDeploymentLogListParams]'s query parameters as
// `url.Values`.
func (r InferenceDeploymentLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments%2F%7Bdeployment_name%7D%2Flogs/get/parameters/4'
// "$.paths['/cloud/v3/inference/{project_id}/deployments/{deployment_name}/logs'].get.parameters[4]"
type InferenceDeploymentLogListParamsOrderBy string

const (
	InferenceDeploymentLogListParamsOrderByTimeAsc  InferenceDeploymentLogListParamsOrderBy = "time.asc"
	InferenceDeploymentLogListParamsOrderByTimeDesc InferenceDeploymentLogListParamsOrderBy = "time.desc"
)
