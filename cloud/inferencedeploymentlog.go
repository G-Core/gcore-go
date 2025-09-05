// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
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
func (r *InferenceDeploymentLogService) List(ctx context.Context, deploymentName string, params InferenceDeploymentLogListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[InferenceDeploymentLog], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
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
func (r *InferenceDeploymentLogService) ListAutoPaging(ctx context.Context, deploymentName string, params InferenceDeploymentLogListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[InferenceDeploymentLog] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, deploymentName, params, opts...))
}

type InferenceDeploymentLog struct {
	// Log message.
	Message string `json:"message,required"`
	// Pod name.
	Pod string `json:"pod,required"`
	// Region ID where the container is deployed.
	RegionID int64 `json:"region_id,required"`
	// Log message timestamp in ISO 8601 format.
	Time time.Time `json:"time,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Pod         respjson.Field
		RegionID    respjson.Field
		Time        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceDeploymentLog) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceDeploymentLogListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `query:"region_id,omitzero" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Order by field
	//
	// Any of "time.asc", "time.desc".
	OrderBy InferenceDeploymentLogListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InferenceDeploymentLogListParams]'s query parameters as
// `url.Values`.
func (r InferenceDeploymentLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Order by field
type InferenceDeploymentLogListParamsOrderBy string

const (
	InferenceDeploymentLogListParamsOrderByTimeAsc  InferenceDeploymentLogListParamsOrderBy = "time.asc"
	InferenceDeploymentLogListParamsOrderByTimeDesc InferenceDeploymentLogListParamsOrderBy = "time.desc"
)
