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
	"github.com/stainless-sdks/gcore-go/packages/resp"
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
	if !params.ProjectID.IsPresent() {
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
	if !params.ProjectID.IsPresent() {
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
	if !body.ProjectID.IsPresent() {
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
	if !query.ProjectID.IsPresent() {
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
	if !params.ProjectID.IsPresent() {
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

// '#/components/schemas/InferenceBoxSecretsSerializer'
// "$.components.schemas.InferenceBoxSecretsSerializer"
type InferenceSecret struct {
	// '#/components/schemas/InferenceBoxSecretsSerializer/properties/data'
	// "$.components.schemas.InferenceBoxSecretsSerializer.properties.data"
	Data AwsIamData `json:"data,required"`
	// '#/components/schemas/InferenceBoxSecretsSerializer/properties/name'
	// "$.components.schemas.InferenceBoxSecretsSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/InferenceBoxSecretsSerializer/properties/type'
	// "$.components.schemas.InferenceBoxSecretsSerializer.properties.type"
	Type string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data        resp.Field
		Name        resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceSecret) RawJSON() string { return r.JSON.raw }
func (r *InferenceSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceSecretNewParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fsecrets/post/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/secrets'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/components/schemas/InferenceBoxSecretsInSerializer/properties/data'
	// "$.components.schemas.InferenceBoxSecretsInSerializer.properties.data"
	Data AwsIamDataParam `json:"data,omitzero,required"`
	// '#/components/schemas/InferenceBoxSecretsInSerializer/properties/name'
	// "$.components.schemas.InferenceBoxSecretsInSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/InferenceBoxSecretsInSerializer/properties/type'
	// "$.components.schemas.InferenceBoxSecretsInSerializer.properties.type"
	Type string `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceSecretNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InferenceSecretNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceSecretNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type InferenceSecretListParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fsecrets/get/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/secrets'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fsecrets/get/parameters/1'
	// "$.paths['/cloud/v3/inference/{project_id}/secrets'].get.parameters[1]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fsecrets/get/parameters/2'
	// "$.paths['/cloud/v3/inference/{project_id}/secrets'].get.parameters[2]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceSecretListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [InferenceSecretListParams]'s query parameters as
// `url.Values`.
func (r InferenceSecretListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InferenceSecretDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fsecrets%2F%7Bsecret_name%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/secrets/{secret_name}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceSecretDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type InferenceSecretGetParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fsecrets%2F%7Bsecret_name%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/secrets/{secret_name}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceSecretGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type InferenceSecretReplaceParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fsecrets%2F%7Bsecret_name%7D/put/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/secrets/{secret_name}'].put.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/components/schemas/InferenceSecretInUpdateSerializer/properties/data'
	// "$.components.schemas.InferenceSecretInUpdateSerializer.properties.data"
	Data AwsIamDataParam `json:"data,omitzero,required"`
	// '#/components/schemas/InferenceSecretInUpdateSerializer/properties/type'
	// "$.components.schemas.InferenceSecretInUpdateSerializer.properties.type"
	Type string `json:"type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceSecretReplaceParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InferenceSecretReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceSecretReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
