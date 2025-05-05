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
	if !params.ProjectID.Valid() {
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
	if !params.ProjectID.Valid() {
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
	if !body.ProjectID.Valid() {
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
	if !query.ProjectID.Valid() {
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
	if !params.ProjectID.Valid() {
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

type InferenceRegistryCredential struct {
	// Registry credential name.
	Name string `json:"name,required"`
	// Project ID to which the inference registry credentials belongs.
	ProjectID int64 `json:"project_id,required"`
	// Registry URL.
	RegistryURL string `json:"registry_url,required"`
	// Registry username.
	Username string `json:"username,required"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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

type InferenceRegistryCredentialFull struct {
	// Registry credential name.
	Name string `json:"name,required"`
	// Registry password.
	Password string `json:"password,required"`
	// Project ID to which the inference registry credentials belongs.
	ProjectID int64 `json:"project_id,required"`
	// Registry URL.
	RegistryURL string `json:"registry_url,required"`
	// Registry username.
	Username string `json:"username,required"`
	// JSON contains metadata for fields, check presence with [resp.Field.Valid].
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
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Registry credential name.
	Name string `json:"name,required"`
	// Registry password.
	Password string `json:"password,required"`
	// Registry URL.
	RegistryURL string `json:"registry_url,required"`
	// Registry username.
	Username string `json:"username,required"`
	paramObj
}

func (r InferenceRegistryCredentialNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRegistryCredentialNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type InferenceRegistryCredentialListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
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
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

type InferenceRegistryCredentialGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

type InferenceRegistryCredentialReplaceParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Registry password.
	Password string `json:"password,required"`
	// Registry URL.
	RegistryURL string `json:"registry_url,required"`
	// Registry username.
	Username string `json:"username,required"`
	paramObj
}

func (r InferenceRegistryCredentialReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRegistryCredentialReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
