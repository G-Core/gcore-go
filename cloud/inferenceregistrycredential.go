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

// InferenceRegistryCredentialService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceRegistryCredentialService] method instead.
type InferenceRegistryCredentialService struct {
	Options []option.RequestOption
}

// NewInferenceRegistryCredentialService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewInferenceRegistryCredentialService(opts ...option.RequestOption) (r InferenceRegistryCredentialService) {
	r = InferenceRegistryCredentialService{}
	r.Options = opts
	return
}

// Create inference registry credential
func (r *InferenceRegistryCredentialService) New(ctx context.Context, params InferenceRegistryCredentialNewParams, opts ...option.RequestOption) (res *InferenceRegistryCredentialFull, err error) {
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
	path := fmt.Sprintf("cloud/v3/inference/%v/registry_credentials", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List inference registry credentials
func (r *InferenceRegistryCredentialService) List(ctx context.Context, params InferenceRegistryCredentialListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[InferenceRegistryCredential], err error) {
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
	path := fmt.Sprintf("cloud/v3/inference/%v/registry_credentials", params.ProjectID.Value)
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

// List inference registry credentials
func (r *InferenceRegistryCredentialService) ListAutoPaging(ctx context.Context, params InferenceRegistryCredentialListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[InferenceRegistryCredential] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete inference registry credential
func (r *InferenceRegistryCredentialService) Delete(ctx context.Context, credentialName string, body InferenceRegistryCredentialDeleteParams, opts ...option.RequestOption) (err error) {
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
	if credentialName == "" {
		err = errors.New("missing required credential_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/registry_credentials/%s", body.ProjectID.Value, credentialName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get inference registry credential
func (r *InferenceRegistryCredentialService) Get(ctx context.Context, credentialName string, query InferenceRegistryCredentialGetParams, opts ...option.RequestOption) (res *InferenceRegistryCredential, err error) {
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
	if credentialName == "" {
		err = errors.New("missing required credential_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/registry_credentials/%s", query.ProjectID.Value, credentialName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update inference registry credential
func (r *InferenceRegistryCredentialService) Replace(ctx context.Context, credentialName string, params InferenceRegistryCredentialReplaceParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if credentialName == "" {
		err = errors.New("missing required credential_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/registry_credentials/%s", params.ProjectID.Value, credentialName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

// '#/components/schemas/InferenceRegistryCredentialOutSerializer'
// "$.components.schemas.InferenceRegistryCredentialOutSerializer"
type InferenceRegistryCredential struct {
	// '#/components/schemas/InferenceRegistryCredentialOutSerializer/properties/name'
	// "$.components.schemas.InferenceRegistryCredentialOutSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/InferenceRegistryCredentialOutSerializer/properties/project_id'
	// "$.components.schemas.InferenceRegistryCredentialOutSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/InferenceRegistryCredentialOutSerializer/properties/registry_url'
	// "$.components.schemas.InferenceRegistryCredentialOutSerializer.properties.registry_url"
	RegistryURL string `json:"registry_url,required"`
	// '#/components/schemas/InferenceRegistryCredentialOutSerializer/properties/username'
	// "$.components.schemas.InferenceRegistryCredentialOutSerializer.properties.username"
	Username string `json:"username,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		ProjectID   resp.Field
		RegistryURL resp.Field
		Username    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceRegistryCredential) RawJSON() string { return r.JSON.raw }
func (r *InferenceRegistryCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InferenceRegistryCredentialOutFullSerializer'
// "$.components.schemas.InferenceRegistryCredentialOutFullSerializer"
type InferenceRegistryCredentialFull struct {
	// '#/components/schemas/InferenceRegistryCredentialOutFullSerializer/properties/name'
	// "$.components.schemas.InferenceRegistryCredentialOutFullSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/InferenceRegistryCredentialOutFullSerializer/properties/password'
	// "$.components.schemas.InferenceRegistryCredentialOutFullSerializer.properties.password"
	Password string `json:"password,required"`
	// '#/components/schemas/InferenceRegistryCredentialOutFullSerializer/properties/project_id'
	// "$.components.schemas.InferenceRegistryCredentialOutFullSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/InferenceRegistryCredentialOutFullSerializer/properties/registry_url'
	// "$.components.schemas.InferenceRegistryCredentialOutFullSerializer.properties.registry_url"
	RegistryURL string `json:"registry_url,required"`
	// '#/components/schemas/InferenceRegistryCredentialOutFullSerializer/properties/username'
	// "$.components.schemas.InferenceRegistryCredentialOutFullSerializer.properties.username"
	Username string `json:"username,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		Password    resp.Field
		ProjectID   resp.Field
		RegistryURL resp.Field
		Username    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceRegistryCredentialFull) RawJSON() string { return r.JSON.raw }
func (r *InferenceRegistryCredentialFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceRegistryCredentialNewParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fregistry_credentials/post/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/registry_credentials'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/components/schemas/InferenceRegistryCredentialInSerializer/properties/name'
	// "$.components.schemas.InferenceRegistryCredentialInSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/InferenceRegistryCredentialInSerializer/properties/password'
	// "$.components.schemas.InferenceRegistryCredentialInSerializer.properties.password"
	Password string `json:"password,required"`
	// '#/components/schemas/InferenceRegistryCredentialInSerializer/properties/registry_url'
	// "$.components.schemas.InferenceRegistryCredentialInSerializer.properties.registry_url"
	RegistryURL string `json:"registry_url,required"`
	// '#/components/schemas/InferenceRegistryCredentialInSerializer/properties/username'
	// "$.components.schemas.InferenceRegistryCredentialInSerializer.properties.username"
	Username string `json:"username,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceRegistryCredentialNewParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r InferenceRegistryCredentialNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRegistryCredentialNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type InferenceRegistryCredentialListParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fregistry_credentials/get/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/registry_credentials'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fregistry_credentials/get/parameters/1'
	// "$.paths['/cloud/v3/inference/{project_id}/registry_credentials'].get.parameters[1]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fregistry_credentials/get/parameters/2'
	// "$.paths['/cloud/v3/inference/{project_id}/registry_credentials'].get.parameters[2]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceRegistryCredentialListParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [InferenceRegistryCredentialListParams]'s query parameters
// as `url.Values`.
func (r InferenceRegistryCredentialListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InferenceRegistryCredentialDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fregistry_credentials%2F%7Bcredential_name%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/registry_credentials/{credential_name}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceRegistryCredentialDeleteParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type InferenceRegistryCredentialGetParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fregistry_credentials%2F%7Bcredential_name%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/registry_credentials/{credential_name}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceRegistryCredentialGetParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type InferenceRegistryCredentialReplaceParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fregistry_credentials%2F%7Bcredential_name%7D/put/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/registry_credentials/{credential_name}'].put.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/components/schemas/InferenceRegistryCredentialInUpdateSerializer/properties/password'
	// "$.components.schemas.InferenceRegistryCredentialInUpdateSerializer.properties.password"
	Password string `json:"password,required"`
	// '#/components/schemas/InferenceRegistryCredentialInUpdateSerializer/properties/registry_url'
	// "$.components.schemas.InferenceRegistryCredentialInUpdateSerializer.properties.registry_url"
	RegistryURL string `json:"registry_url,required"`
	// '#/components/schemas/InferenceRegistryCredentialInUpdateSerializer/properties/username'
	// "$.components.schemas.InferenceRegistryCredentialInUpdateSerializer.properties.username"
	Username string `json:"username,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceRegistryCredentialReplaceParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r InferenceRegistryCredentialReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRegistryCredentialReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
