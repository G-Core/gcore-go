// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fastedge

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	shimjson "github.com/G-Core/gcore-go/internal/encoding/json"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// AppService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
	Logs    AppLogService
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r AppService) {
	r = AppService{}
	r.Options = opts
	r.Logs = NewAppLogService(opts...)
	return
}

// Add a new app
func (r *AppService) New(ctx context.Context, body AppNewParams, opts ...option.RequestOption) (res *AppShort, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fastedge/v1/apps"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update app
func (r *AppService) Update(ctx context.Context, id int64, body AppUpdateParams, opts ...option.RequestOption) (res *AppShort, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("fastedge/v1/apps/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List client's apps
func (r *AppService) List(ctx context.Context, query AppListParams, opts ...option.RequestOption) (res *pagination.OffsetPageFastedgeApps[AppShort], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "fastedge/v1/apps"
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

// List client's apps
func (r *AppService) ListAutoPaging(ctx context.Context, query AppListParams, opts ...option.RequestOption) *pagination.OffsetPageFastedgeAppsAutoPager[AppShort] {
	return pagination.NewOffsetPageFastedgeAppsAutoPager(r.List(ctx, query, opts...))
}

// Delete app
func (r *AppService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("fastedge/v1/apps/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get app details
func (r *AppService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *App, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("fastedge/v1/apps/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an app
func (r *AppService) Replace(ctx context.Context, id int64, body AppReplaceParams, opts ...option.RequestOption) (res *AppShort, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("fastedge/v1/apps/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type App struct {
	// Wasm API type
	APIType string `json:"api_type"`
	// Binary ID
	Binary int64 `json:"binary"`
	// App description
	Comment string `json:"comment"`
	// Switch on logging for 30 minutes (switched off by default)
	Debug bool `json:"debug"`
	// When debugging finishes
	DebugUntil time.Time `json:"debug_until" format:"date-time"`
	// Environment variables
	Env map[string]string `json:"env"`
	// Logging channel (by default - kafka, which allows exploring logs with API)
	//
	// Any of "kafka", "none".
	Log AppLog `json:"log,nullable"`
	// App name
	Name string `json:"name"`
	// Networks
	Networks []string `json:"networks"`
	// Plan name
	Plan string `json:"plan"`
	// Plan ID
	PlanID int64 `json:"plan_id"`
	// Extra headers to add to the response
	RspHeaders map[string]string `json:"rsp_headers"`
	// Application secrets
	Secrets map[string]AppSecret `json:"secrets"`
	// Status code:
	// 0 - draft (inactive)
	// 1 - enabled
	// 2 - disabled
	// 3 - hourly call limit exceeded
	// 4 - daily call limit exceeded
	// 5 - suspended
	Status int64 `json:"status"`
	// Application edge stores
	Stores map[string]AppStore `json:"stores"`
	// Template ID
	Template int64 `json:"template"`
	// Template name
	TemplateName string `json:"template_name"`
	// App URL
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIType      respjson.Field
		Binary       respjson.Field
		Comment      respjson.Field
		Debug        respjson.Field
		DebugUntil   respjson.Field
		Env          respjson.Field
		Log          respjson.Field
		Name         respjson.Field
		Networks     respjson.Field
		Plan         respjson.Field
		PlanID       respjson.Field
		RspHeaders   respjson.Field
		Secrets      respjson.Field
		Status       respjson.Field
		Stores       respjson.Field
		Template     respjson.Field
		TemplateName respjson.Field
		URL          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r App) RawJSON() string { return r.JSON.raw }
func (r *App) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this App to a AppParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AppParam.Overrides()
func (r App) ToParam() AppParam {
	return param.Override[AppParam](json.RawMessage(r.RawJSON()))
}

// Logging channel (by default - kafka, which allows exploring logs with API)
type AppLog string

const (
	AppLogKafka AppLog = "kafka"
	AppLogNone  AppLog = "none"
)

// Application secret short description
type AppSecret struct {
	// The unique identifier of the secret.
	ID int64 `json:"id,required"`
	// A description or comment about the secret.
	Comment string `json:"comment"`
	// The unique name of the secret.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Comment     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppSecret) RawJSON() string { return r.JSON.raw }
func (r *AppSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Application stores
type AppStore struct {
	// The identifier of the store
	ID int64 `json:"id,required"`
	// The name of the store
	Name string `json:"name,required"`
	// A description of the store
	Comment string `json:"comment"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Comment     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppStore) RawJSON() string { return r.JSON.raw }
func (r *AppStore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppParam struct {
	// Binary ID
	Binary param.Opt[int64] `json:"binary,omitzero"`
	// App description
	Comment param.Opt[string] `json:"comment,omitzero"`
	// Switch on logging for 30 minutes (switched off by default)
	Debug param.Opt[bool] `json:"debug,omitzero"`
	// App name
	Name param.Opt[string] `json:"name,omitzero"`
	// Status code:
	// 0 - draft (inactive)
	// 1 - enabled
	// 2 - disabled
	// 3 - hourly call limit exceeded
	// 4 - daily call limit exceeded
	// 5 - suspended
	Status param.Opt[int64] `json:"status,omitzero"`
	// Template ID
	Template param.Opt[int64] `json:"template,omitzero"`
	// Logging channel (by default - kafka, which allows exploring logs with API)
	//
	// Any of "kafka", "none".
	Log AppLog `json:"log,omitzero"`
	// Environment variables
	Env map[string]string `json:"env,omitzero"`
	// Extra headers to add to the response
	RspHeaders map[string]string `json:"rsp_headers,omitzero"`
	// Application secrets
	Secrets map[string]AppSecretParam `json:"secrets,omitzero"`
	// Application edge stores
	Stores map[string]AppStoreParam `json:"stores,omitzero"`
	paramObj
}

func (r AppParam) MarshalJSON() (data []byte, err error) {
	type shadow AppParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Application secret short description
//
// The property ID is required.
type AppSecretParam struct {
	// The unique identifier of the secret.
	ID int64 `json:"id,required"`
	paramObj
}

func (r AppSecretParam) MarshalJSON() (data []byte, err error) {
	type shadow AppSecretParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppSecretParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Application stores
//
// The properties ID, Name are required.
type AppStoreParam struct {
	// The identifier of the store
	ID int64 `json:"id,required"`
	paramObj
}

func (r AppStoreParam) MarshalJSON() (data []byte, err error) {
	type shadow AppStoreParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppStoreParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppShort struct {
	// App ID
	ID int64 `json:"id,required"`
	// Wasm API type
	APIType string `json:"api_type,required"`
	// Binary ID
	Binary int64 `json:"binary,required"`
	// App name
	Name string `json:"name,required"`
	// Application plan ID
	PlanID int64 `json:"plan_id,required"`
	// Status code:
	// 0 - draft (inactive)
	// 1 - enabled
	// 2 - disabled
	// 3 - hourly call limit exceeded
	// 4 - daily call limit exceeded
	// 5 - suspended
	Status int64 `json:"status,required"`
	// Description of the binary
	Comment string `json:"comment"`
	// Switch on logging for 30 minutes (switched off by default)
	Debug bool `json:"debug"`
	// When debugging finishes
	DebugUntil time.Time `json:"debug_until" format:"date-time"`
	// Networks
	Networks []string `json:"networks"`
	// Application plan name
	Plan string `json:"plan"`
	// Template ID
	Template int64 `json:"template"`
	// Template name
	TemplateName string `json:"template_name"`
	// ID of the binary the app can be upgraded to
	UpgradeableTo int64 `json:"upgradeable_to"`
	// App URL
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		APIType       respjson.Field
		Binary        respjson.Field
		Name          respjson.Field
		PlanID        respjson.Field
		Status        respjson.Field
		Comment       respjson.Field
		Debug         respjson.Field
		DebugUntil    respjson.Field
		Networks      respjson.Field
		Plan          respjson.Field
		Template      respjson.Field
		TemplateName  respjson.Field
		UpgradeableTo respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppShort) RawJSON() string { return r.JSON.raw }
func (r *AppShort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppNewParams struct {
	App AppParam
	paramObj
}

func (r AppNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.App)
}
func (r *AppNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.App)
}

type AppUpdateParams struct {
	App AppParam
	paramObj
}

func (r AppUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.App)
}
func (r *AppUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.App)
}

type AppListParams struct {
	// Binary ID
	Binary param.Opt[int64] `query:"binary,omitzero" json:"-"`
	// Limit for pagination
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Name of the app
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Offset for pagination
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Plan ID
	Plan param.Opt[int64] `query:"plan,omitzero" json:"-"`
	// Status code:
	// 0 - draft (inactive)
	// 1 - enabled
	// 2 - disabled
	// 3 - hourly call limit exceeded
	// 4 - daily call limit exceeded
	// 5 - suspended
	Status param.Opt[int64] `query:"status,omitzero" json:"-"`
	// Template ID
	Template param.Opt[int64] `query:"template,omitzero" json:"-"`
	// API type:
	// wasi-http - WASI with HTTP entry point
	// proxy-wasm - Proxy-Wasm app, callable from CDN
	//
	// Any of "wasi-http", "proxy-wasm".
	APIType AppListParamsAPIType `query:"api_type,omitzero" json:"-"`
	// Ordering
	//
	// Any of "name", "-name", "status", "-status", "id", "-id", "template",
	// "-template", "binary", "-binary", "plan", "-plan".
	Ordering AppListParamsOrdering `query:"ordering,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AppListParams]'s query parameters as `url.Values`.
func (r AppListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// API type:
// wasi-http - WASI with HTTP entry point
// proxy-wasm - Proxy-Wasm app, callable from CDN
type AppListParamsAPIType string

const (
	AppListParamsAPITypeWasiHTTP  AppListParamsAPIType = "wasi-http"
	AppListParamsAPITypeProxyWasm AppListParamsAPIType = "proxy-wasm"
)

// Ordering
type AppListParamsOrdering string

const (
	AppListParamsOrderingName          AppListParamsOrdering = "name"
	AppListParamsOrderingMinusName     AppListParamsOrdering = "-name"
	AppListParamsOrderingStatus        AppListParamsOrdering = "status"
	AppListParamsOrderingMinusStatus   AppListParamsOrdering = "-status"
	AppListParamsOrderingID            AppListParamsOrdering = "id"
	AppListParamsOrderingMinusID       AppListParamsOrdering = "-id"
	AppListParamsOrderingTemplate      AppListParamsOrdering = "template"
	AppListParamsOrderingMinusTemplate AppListParamsOrdering = "-template"
	AppListParamsOrderingBinary        AppListParamsOrdering = "binary"
	AppListParamsOrderingMinusBinary   AppListParamsOrdering = "-binary"
	AppListParamsOrderingPlan          AppListParamsOrdering = "plan"
	AppListParamsOrderingMinusPlan     AppListParamsOrdering = "-plan"
)

type AppReplaceParams struct {
	Body AppReplaceParamsBody
	paramObj
}

func (r AppReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AppReplaceParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type AppReplaceParamsBody struct {
	AppParam
}

func (r AppReplaceParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*AppReplaceParamsBody
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}
