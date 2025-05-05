// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/respjson"
)

// InferenceSecretService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceSecretService] method instead.
type InferenceSecretService struct {
	Options []option.RequestOption
}

// NewInferenceSecretService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInferenceSecretService(opts ...option.RequestOption) (r InferenceSecretService) {
	r = InferenceSecretService{}
	r.Options = opts
	return
}

// Create Inference Secret
func (r *InferenceSecretService) New(ctx context.Context, params InferenceSecretNewParams, opts ...option.RequestOption) (res *InferenceSecret, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/secrets", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List Secrets for Inference
func (r *InferenceSecretService) List(ctx context.Context, params InferenceSecretListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[InferenceSecret], err error) {
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
	path := fmt.Sprintf("cloud/v3/inference/%v/secrets", params.ProjectID.Value)
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

// List Secrets for Inference
func (r *InferenceSecretService) ListAutoPaging(ctx context.Context, params InferenceSecretListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[InferenceSecret] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete Inference Secret
func (r *InferenceSecretService) Delete(ctx context.Context, secretName string, body InferenceSecretDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/secrets/%s", body.ProjectID.Value, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get Inference Secret
func (r *InferenceSecretService) Get(ctx context.Context, secretName string, query InferenceSecretGetParams, opts ...option.RequestOption) (res *InferenceSecret, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/secrets/%s", query.ProjectID.Value, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Inference Secret
func (r *InferenceSecretService) Replace(ctx context.Context, secretName string, params InferenceSecretReplaceParams, opts ...option.RequestOption) (res *InferenceSecret, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if secretName == "" {
		err = errors.New("missing required secret_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/secrets/%s", params.ProjectID.Value, secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type InferenceSecret struct {
	// Secret data.
	Data AwsIamData `json:"data,required"`
	// Secret name.
	Name string `json:"name,required"`
	// Secret type.
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceSecret) RawJSON() string { return r.JSON.raw }
func (r *InferenceSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceSecretNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Secret data.
	Data AwsIamDataParam `json:"data,omitzero,required"`
	// Secret name.
	Name string `json:"name,required"`
	// Secret type. Currently only `aws-iam` is supported.
	Type string `json:"type,required"`
	paramObj
}

func (r InferenceSecretNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceSecretNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type InferenceSecretListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InferenceSecretListParams]'s query parameters as
// `url.Values`.
func (r InferenceSecretListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InferenceSecretDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

type InferenceSecretGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

type InferenceSecretReplaceParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Secret data.
	Data AwsIamDataParam `json:"data,omitzero,required"`
	// Secret type.
	Type string `json:"type,required"`
	paramObj
}

func (r InferenceSecretReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceSecretReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
