// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// InferenceAPIKeyService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceAPIKeyService] method instead.
type InferenceAPIKeyService struct {
	Options []option.RequestOption
}

// NewInferenceAPIKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInferenceAPIKeyService(opts ...option.RequestOption) (r InferenceAPIKeyService) {
	r = InferenceAPIKeyService{}
	r.Options = opts
	return
}

// This endpoint creates a new API key for everywhere inference. It returs api
// key's actual secret only once after creation.
func (r *InferenceAPIKeyService) New(ctx context.Context, params InferenceAPIKeyNewParams, opts ...option.RequestOption) (res *InferenceAPIKeyCreate, err error) {
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
	path := fmt.Sprintf("cloud/v3/inference/%v/api_keys", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// This endpoint updates a specific API key for everywhere inference.
func (r *InferenceAPIKeyService) Update(ctx context.Context, apiKeyName string, params InferenceAPIKeyUpdateParams, opts ...option.RequestOption) (res *InferenceAPIKey, err error) {
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
	if apiKeyName == "" {
		err = errors.New("missing required api_key_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/api_keys/%s", params.ProjectID.Value, apiKeyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// This endpoint retrieves a list of API keys for everywhere inference.
func (r *InferenceAPIKeyService) List(ctx context.Context, params InferenceAPIKeyListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[InferenceAPIKey], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("cloud/v3/inference/%v/api_keys", params.ProjectID.Value)
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

// This endpoint retrieves a list of API keys for everywhere inference.
func (r *InferenceAPIKeyService) ListAutoPaging(ctx context.Context, params InferenceAPIKeyListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[InferenceAPIKey] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// This endpoint deletes a specific API key for everywhere inference. If the API
// key is attached to any inference deployments, it will not be removed.
// ConflictError will be raised
func (r *InferenceAPIKeyService) Delete(ctx context.Context, apiKeyName string, body InferenceAPIKeyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if apiKeyName == "" {
		err = errors.New("missing required api_key_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/api_keys/%s", body.ProjectID.Value, apiKeyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// This endpoint retrieves a specific API key for everywhere inference.
func (r *InferenceAPIKeyService) Get(ctx context.Context, apiKeyName string, query InferenceAPIKeyGetParams, opts ...option.RequestOption) (res *InferenceAPIKey, err error) {
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
	if apiKeyName == "" {
		err = errors.New("missing required api_key_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/api_keys/%s", query.ProjectID.Value, apiKeyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type InferenceAPIKey struct {
	// Timestamp when the API Key was created.
	CreatedAt string `json:"created_at,required"`
	// List of inference deployment names to which this API Key has been attached.
	DeploymentNames []string `json:"deployment_names,required"`
	// Description of the API Key.
	Description string `json:"description,required"`
	// Timestamp when the API Key will expire.
	ExpiresAt string `json:"expires_at,required"`
	// API Key name.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt       respjson.Field
		DeploymentNames respjson.Field
		Description     respjson.Field
		ExpiresAt       respjson.Field
		Name            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceAPIKey) RawJSON() string { return r.JSON.raw }
func (r *InferenceAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceAPIKeyCreate struct {
	// Timestamp when the API Key was created.
	CreatedAt string `json:"created_at,required"`
	// List of inference deployment names to which this API Key has been attached.
	DeploymentNames []string `json:"deployment_names,required"`
	// Description of the API Key.
	Description string `json:"description,required"`
	// Timestamp when the API Key will expire.
	ExpiresAt string `json:"expires_at,required"`
	// API Key name.
	Name string `json:"name,required"`
	// The actual API Key secret.
	Secret string `json:"secret,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt       respjson.Field
		DeploymentNames respjson.Field
		Description     respjson.Field
		ExpiresAt       respjson.Field
		Name            respjson.Field
		Secret          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceAPIKeyCreate) RawJSON() string { return r.JSON.raw }
func (r *InferenceAPIKeyCreate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceAPIKeyNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Name of the API Key.
	Name string `json:"name,required"`
	// Description of the API Key.
	Description param.Opt[string] `json:"description,omitzero"`
	// Expiration date of the API Key in ISO 8601 format.
	ExpiresAt param.Opt[string] `json:"expires_at,omitzero"`
	paramObj
}

func (r InferenceAPIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceAPIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceAPIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceAPIKeyUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Description of the API Key.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r InferenceAPIKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceAPIKeyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceAPIKeyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceAPIKeyListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InferenceAPIKeyListParams]'s query parameters as
// `url.Values`.
func (r InferenceAPIKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InferenceAPIKeyDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

type InferenceAPIKeyGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}
