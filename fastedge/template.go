// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fastedge

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	shimjson "github.com/G-Core/gcore-go/internal/encoding/json"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// FastEdge templates encapsulate reusable configurations for FastEdge
// applications, including a WebAssembly binary reference and configurable
// parameters.
//
// TemplateService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTemplateService] method instead.
type TemplateService struct {
	Options []option.RequestOption
}

// NewTemplateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTemplateService(opts ...option.RequestOption) (r TemplateService) {
	r = TemplateService{}
	r.Options = opts
	return
}

// Create a new application template from an existing binary and configuration.
// Templates can be shared with groups to enable collaborative application
// development.
func (r *TemplateService) New(ctx context.Context, body TemplateNewParams, opts ...option.RequestOption) (res *TemplateShort, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fastedge/v1/template"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve available application templates including system templates and
// user-created templates. Templates provide pre-configured application settings
// that can be reused for quick deployment.
func (r *TemplateService) List(ctx context.Context, query TemplateListParams, opts ...option.RequestOption) (res *pagination.OffsetPageFastedgeTemplates[TemplateShort], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "fastedge/v1/template"
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

// Retrieve available application templates including system templates and
// user-created templates. Templates provide pre-configured application settings
// that can be reused for quick deployment.
func (r *TemplateService) ListAutoPaging(ctx context.Context, query TemplateListParams, opts ...option.RequestOption) *pagination.OffsetPageFastedgeTemplatesAutoPager[TemplateShort] {
	return pagination.NewOffsetPageFastedgeTemplatesAutoPager(r.List(ctx, query, opts...))
}

// Remove a template from the platform. Templates shared with groups require
// force=true parameter to delete.
func (r *TemplateService) Delete(ctx context.Context, templateID int64, body TemplateDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("fastedge/v1/template/%v", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Retrieve complete configuration and metadata for a specific template. Use this
// to inspect template parameters before creating applications from it.
func (r *TemplateService) Get(ctx context.Context, templateID int64, opts ...option.RequestOption) (res *Template, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("fastedge/v1/template/%v", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Modify an existing template's configuration. Updates affect future applications
// created from this template; existing apps are not changed.
func (r *TemplateService) Replace(ctx context.Context, templateID int64, body TemplateReplaceParams, opts ...option.RequestOption) (res *TemplateShort, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("fastedge/v1/template/%v", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type Template struct {
	// Wasm API type
	APIType string `json:"api_type" api:"required"`
	// ID of the WebAssembly binary to use for this template
	BinaryID int64 `json:"binary_id" api:"required"`
	// Unique name for the template (used for identification and searching)
	Name string `json:"name" api:"required"`
	// Is the template owned by user?
	Owned bool `json:"owned" api:"required"`
	// Parameters
	Params []TemplateParameterResp `json:"params" api:"required"`
	// Detailed markdown description explaining template features and usage
	LongDescr string `json:"long_descr"`
	// Brief one-line description displayed in template listings
	ShortDescr string `json:"short_descr"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIType     respjson.Field
		BinaryID    respjson.Field
		Name        respjson.Field
		Owned       respjson.Field
		Params      respjson.Field
		LongDescr   respjson.Field
		ShortDescr  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Template) RawJSON() string { return r.JSON.raw }
func (r *Template) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Template to a TemplateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TemplateParam.Overrides()
func (r Template) ToParam() TemplateParam {
	return param.Override[TemplateParam](json.RawMessage(r.RawJSON()))
}

// The properties APIType, BinaryID, Name, Owned, Params are required.
type TemplateParam struct {
	// ID of the WebAssembly binary to use for this template
	BinaryID int64 `json:"binary_id" api:"required"`
	// Unique name for the template (used for identification and searching)
	Name string `json:"name" api:"required"`
	// Is the template owned by user?
	Owned bool `json:"owned" api:"required"`
	// Parameters
	Params []TemplateParameter `json:"params,omitzero" api:"required"`
	// Detailed markdown description explaining template features and usage
	LongDescr param.Opt[string] `json:"long_descr,omitzero"`
	// Brief one-line description displayed in template listings
	ShortDescr param.Opt[string] `json:"short_descr,omitzero"`
	paramObj
}

func (r TemplateParam) MarshalJSON() (data []byte, err error) {
	type shadow TemplateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateParameterResp struct {
	// Parameter type determines validation and UI rendering:
	// string - text input
	// number - numeric input
	// date - date picker
	// time - time picker
	// secret - references a secret
	// store - references an edge store
	// bool - boolean toggle
	// json - JSON editor or multiline text area with JSON validation
	// enum - dropdown/select with allowed values defined via parameter metadata
	//
	// Any of "string", "number", "date", "time", "secret", "store", "bool", "json",
	// "enum".
	DataType TemplateParameterDataType `json:"data_type" api:"required"`
	// If true, this parameter must be provided when instantiating the template
	Mandatory bool `json:"mandatory" api:"required"`
	// Parameter name used as a placeholder in template (e.g., `API_KEY`,
	// `DATABASE_URL`)
	Name string `json:"name" api:"required"`
	// Human-readable explanation of what this parameter controls
	Descr string `json:"descr"`
	// Optional JSON-encoded string for additional parameter metadata, such as allowed
	// values for enum types
	Metadata string `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataType    respjson.Field
		Mandatory   respjson.Field
		Name        respjson.Field
		Descr       respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateParameterResp) RawJSON() string { return r.JSON.raw }
func (r *TemplateParameterResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TemplateParameterResp to a TemplateParameter.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TemplateParameter.Overrides()
func (r TemplateParameterResp) ToParam() TemplateParameter {
	return param.Override[TemplateParameter](json.RawMessage(r.RawJSON()))
}

// Parameter type determines validation and UI rendering:
// string - text input
// number - numeric input
// date - date picker
// time - time picker
// secret - references a secret
// store - references an edge store
// bool - boolean toggle
// json - JSON editor or multiline text area with JSON validation
// enum - dropdown/select with allowed values defined via parameter metadata
type TemplateParameterDataType string

const (
	TemplateParameterDataTypeString TemplateParameterDataType = "string"
	TemplateParameterDataTypeNumber TemplateParameterDataType = "number"
	TemplateParameterDataTypeDate   TemplateParameterDataType = "date"
	TemplateParameterDataTypeTime   TemplateParameterDataType = "time"
	TemplateParameterDataTypeSecret TemplateParameterDataType = "secret"
	TemplateParameterDataTypeStore  TemplateParameterDataType = "store"
	TemplateParameterDataTypeBool   TemplateParameterDataType = "bool"
	TemplateParameterDataTypeJson   TemplateParameterDataType = "json"
	TemplateParameterDataTypeEnum   TemplateParameterDataType = "enum"
)

// The properties DataType, Mandatory, Name are required.
type TemplateParameter struct {
	// Parameter type determines validation and UI rendering:
	// string - text input
	// number - numeric input
	// date - date picker
	// time - time picker
	// secret - references a secret
	// store - references an edge store
	// bool - boolean toggle
	// json - JSON editor or multiline text area with JSON validation
	// enum - dropdown/select with allowed values defined via parameter metadata
	//
	// Any of "string", "number", "date", "time", "secret", "store", "bool", "json",
	// "enum".
	DataType TemplateParameterDataType `json:"data_type,omitzero" api:"required"`
	// If true, this parameter must be provided when instantiating the template
	Mandatory bool `json:"mandatory" api:"required"`
	// Parameter name used as a placeholder in template (e.g., `API_KEY`,
	// `DATABASE_URL`)
	Name string `json:"name" api:"required"`
	// Human-readable explanation of what this parameter controls
	Descr param.Opt[string] `json:"descr,omitzero"`
	// Optional JSON-encoded string for additional parameter metadata, such as allowed
	// values for enum types
	Metadata param.Opt[string] `json:"metadata,omitzero"`
	paramObj
}

func (r TemplateParameter) MarshalJSON() (data []byte, err error) {
	type shadow TemplateParameter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateShort struct {
	// Template ID
	ID int64 `json:"id" api:"required"`
	// Wasm API type
	APIType string `json:"api_type" api:"required"`
	// Name of the template
	Name string `json:"name" api:"required"`
	// Is the template owned by user?
	Owned bool `json:"owned" api:"required"`
	// Long description of the template
	LongDescr string `json:"long_descr"`
	// Short description of the template
	ShortDescr string `json:"short_descr"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		APIType     respjson.Field
		Name        respjson.Field
		Owned       respjson.Field
		LongDescr   respjson.Field
		ShortDescr  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateShort) RawJSON() string { return r.JSON.raw }
func (r *TemplateShort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateNewParams struct {
	Template TemplateParam
	paramObj
}

func (r TemplateNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Template)
}
func (r *TemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateListParams struct {
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of results to skip for pagination
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// When true, returns only templates created by the client. When false, includes
	// shared templates.
	OnlyMine param.Opt[bool] `query:"only_mine,omitzero" json:"-"`
	// API type:
	// wasi-http - WASI with HTTP entry point
	// proxy-wasm - Proxy-Wasm app, callable from CDN
	//
	// Any of "wasi-http", "proxy-wasm".
	APIType TemplateListParamsAPIType `query:"api_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TemplateListParams]'s query parameters as `url.Values`.
func (r TemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// API type:
// wasi-http - WASI with HTTP entry point
// proxy-wasm - Proxy-Wasm app, callable from CDN
type TemplateListParamsAPIType string

const (
	TemplateListParamsAPITypeWasiHTTP  TemplateListParamsAPIType = "wasi-http"
	TemplateListParamsAPITypeProxyWasm TemplateListParamsAPIType = "proxy-wasm"
)

type TemplateDeleteParams struct {
	// When true, deletes template even if shared with groups. Defaults to false.
	Force param.Opt[bool] `query:"force,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TemplateDeleteParams]'s query parameters as `url.Values`.
func (r TemplateDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type TemplateReplaceParams struct {
	Template TemplateParam
	paramObj
}

func (r TemplateReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Template)
}
func (r *TemplateReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
