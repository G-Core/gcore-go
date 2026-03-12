// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fastedge

import (
	"context"
	"encoding/json"
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

// FastEdge applications combine a WebAssembly binary with configuration,
// environment variables, and secrets for deployment at the CDN edge.
//
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

// Create a new edge application from a WebAssembly binary. The app is configured
// with execution limits, environment variables, and network assignments. Apps
// start in draft status (0) and must be explicitly enabled to receive traffic.
func (r *AppService) New(ctx context.Context, body AppNewParams, opts ...option.RequestOption) (res *AppShort, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fastedge/v1/apps"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of applications owned by the authenticated client.
// Results can be filtered by name, API type, status, template, binary, or plan,
// and sorted by various fields.
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

// Retrieve a paginated list of applications owned by the authenticated client.
// Results can be filtered by name, API type, status, template, binary, or plan,
// and sorted by various fields.
func (r *AppService) ListAutoPaging(ctx context.Context, query AppListParams, opts ...option.RequestOption) *pagination.OffsetPageFastedgeAppsAutoPager[AppShort] {
	return pagination.NewOffsetPageFastedgeAppsAutoPager(r.List(ctx, query, opts...))
}

type AppParam struct {
	// ID of the WebAssembly binary to deploy
	Binary param.Opt[int64] `json:"binary,omitzero"`
	// Optional human-readable description of the application's purpose
	Comment param.Opt[string] `json:"comment,omitzero"`
	// Enable verbose debug logging for 30 minutes. Automatically expires to prevent
	// performance impact.
	Debug param.Opt[bool] `json:"debug,omitzero"`
	// Unique application name (alphanumeric, hyphens allowed)
	Name param.Opt[string] `json:"name,omitzero"`
	// Status code:
	// 0 - draft (inactive)
	// 1 - enabled
	// 2 - disabled
	// 5 - suspended
	Status param.Opt[int64] `json:"status,omitzero"`
	// Template ID
	Template param.Opt[int64] `json:"template,omitzero"`
	// Logging channel. Use 'kafka' to enable log collection (queryable via API), or
	// 'none' to disable logging.
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

// Logging channel. Use 'kafka' to enable log collection (queryable via API), or
// 'none' to disable logging.
type AppLog string

const (
	AppLogKafka AppLog = "kafka"
	AppLogNone  AppLog = "none"
)

// Application secret short description
//
// The property ID is required.
type AppSecretParam struct {
	// The unique identifier of the secret.
	ID int64 `json:"id" api:"required"`
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
// The property ID is required.
type AppStoreParam struct {
	// The identifier of the store
	ID int64 `json:"id" api:"required"`
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
	ID int64 `json:"id" api:"required"`
	// Wasm API type
	APIType string `json:"api_type" api:"required"`
	// Binary ID
	Binary int64 `json:"binary" api:"required"`
	// App name
	Name string `json:"name" api:"required"`
	// Application plan ID
	PlanID int64 `json:"plan_id" api:"required"`
	// Status code:
	// 0 - draft (inactive)
	// 1 - enabled
	// 2 - disabled
	// 3 - hourly call limit exceeded
	// 4 - daily call limit exceeded
	// 5 - suspended
	Status int64 `json:"status" api:"required"`
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

type AppListParams struct {
	// Filter by binary ID (shows apps using this binary)
	Binary param.Opt[int64] `query:"binary,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by application name (case-insensitive partial match)
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Number of results to skip for pagination
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by plan ID
	Plan param.Opt[int64] `query:"plan,omitzero" json:"-"`
	// Status code:
	// 0 - draft (inactive)
	// 1 - enabled
	// 2 - disabled
	// 3 - hourly call limit exceeded
	// 4 - daily call limit exceeded
	// 5 - suspended
	Status param.Opt[int64] `query:"status,omitzero" json:"-"`
	// Filter by template ID (shows apps created from this template)
	Template param.Opt[int64] `query:"template,omitzero" json:"-"`
	// API type:
	// wasi-http - WASI with HTTP entry point
	// proxy-wasm - Proxy-Wasm app, callable from CDN
	//
	// Any of "wasi-http", "proxy-wasm".
	APIType AppListParamsAPIType `query:"api_type,omitzero" json:"-"`
	// Sort order. Use - prefix for descending (e.g., -name sorts by name descending)
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

// Sort order. Use - prefix for descending (e.g., -name sorts by name descending)
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
