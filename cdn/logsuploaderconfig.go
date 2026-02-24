// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// LogsUploaderConfigService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogsUploaderConfigService] method instead.
type LogsUploaderConfigService struct {
	Options []option.RequestOption
}

// NewLogsUploaderConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogsUploaderConfigService(opts ...option.RequestOption) (r LogsUploaderConfigService) {
	r = LogsUploaderConfigService{}
	r.Options = opts
	return
}

// Create logs uploader config.
func (r *LogsUploaderConfigService) New(ctx context.Context, body LogsUploaderConfigNewParams, opts ...option.RequestOption) (res *LogsUploaderConfig, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/logs_uploader/configs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change logs uploader config partially.
func (r *LogsUploaderConfigService) Update(ctx context.Context, id int64, body LogsUploaderConfigUpdateParams, opts ...option.RequestOption) (res *LogsUploaderConfig, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/configs/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get list of logs uploader configs.
func (r *LogsUploaderConfigService) List(ctx context.Context, query LogsUploaderConfigListParams, opts ...option.RequestOption) (res *LogsUploaderConfigList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/logs_uploader/configs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete the logs uploader config from the system permanently.
//
// Notes:
//
//   - **Irreversibility**: This action is irreversible. Once deleted, the logs
//     uploader config cannot be recovered.
func (r *LogsUploaderConfigService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/logs_uploader/configs/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get information about logs uploader config.
func (r *LogsUploaderConfigService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *LogsUploaderConfig, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/configs/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change logs uploader config.
func (r *LogsUploaderConfigService) Replace(ctx context.Context, id int64, body LogsUploaderConfigReplaceParams, opts ...option.RequestOption) (res *LogsUploaderConfig, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/configs/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Validate logs uploader config.
func (r *LogsUploaderConfigService) Validate(ctx context.Context, id int64, opts ...option.RequestOption) (res *LogsUploaderValidation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/configs/%v/validate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type LogsUploaderConfig struct {
	ID int64 `json:"id"`
	// Client that owns the config.
	ClientID int64 `json:"client_id"`
	// Time when the config was created.
	Created time.Time `json:"created" format:"date-time"`
	// Enables or disables the config.
	Enabled bool `json:"enabled"`
	// If set to true, the config will be applied to all CDN resources. If set to
	// false, the config will be applied to the resources specified in the `resources`
	// field.
	ForAllResources bool `json:"for_all_resources"`
	// Name of the config.
	Name string `json:"name"`
	// ID of the policy that should be assigned to given config.
	Policy int64 `json:"policy"`
	// List of resource IDs to which the config should be applied.
	Resources []int64 `json:"resources"`
	// Validation status of the logs uploader config.
	Status LogsUploaderConfigStatus `json:"status"`
	// ID of the target to which logs should be uploaded.
	Target int64 `json:"target"`
	// Time when the config was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ClientID        respjson.Field
		Created         respjson.Field
		Enabled         respjson.Field
		ForAllResources respjson.Field
		Name            respjson.Field
		Policy          respjson.Field
		Resources       respjson.Field
		Status          respjson.Field
		Target          respjson.Field
		Updated         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Validation status of the logs uploader config.
type LogsUploaderConfigStatus struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	LogsUploaderValidation
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderConfigStatus) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderConfigStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderConfigList []LogsUploaderConfig

type LogsUploaderConfigNewParams struct {
	// Name of the config.
	Name string `json:"name" api:"required"`
	// ID of the policy that should be assigned to given config.
	Policy int64 `json:"policy" api:"required"`
	// ID of the target to which logs should be uploaded.
	Target int64 `json:"target" api:"required"`
	// Enables or disables the config.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// If set to true, the config will be applied to all CDN resources. If set to
	// false, the config will be applied to the resources specified in the `resources`
	// field.
	ForAllResources param.Opt[bool] `json:"for_all_resources,omitzero"`
	// List of resource IDs to which the config should be applied.
	Resources []int64 `json:"resources,omitzero"`
	paramObj
}

func (r LogsUploaderConfigNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderConfigNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderConfigNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderConfigUpdateParams struct {
	// Enables or disables the config.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// If set to true, the config will be applied to all CDN resources. If set to
	// false, the config will be applied to the resources specified in the `resources`
	// field.
	ForAllResources param.Opt[bool] `json:"for_all_resources,omitzero"`
	// Name of the config.
	Name param.Opt[string] `json:"name,omitzero"`
	// ID of the policy that should be assigned to given config.
	Policy param.Opt[int64] `json:"policy,omitzero"`
	// ID of the target to which logs should be uploaded.
	Target param.Opt[int64] `json:"target,omitzero"`
	// List of resource IDs to which the config should be applied.
	Resources []int64 `json:"resources,omitzero"`
	paramObj
}

func (r LogsUploaderConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderConfigUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderConfigUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderConfigListParams struct {
	// Search by config name or id.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Filter by ids of CDN resources that are assigned to given config.
	ResourceIDs []int64 `query:"resource_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LogsUploaderConfigListParams]'s query parameters as
// `url.Values`.
func (r LogsUploaderConfigListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogsUploaderConfigReplaceParams struct {
	// Name of the config.
	Name string `json:"name" api:"required"`
	// ID of the policy that should be assigned to given config.
	Policy int64 `json:"policy" api:"required"`
	// ID of the target to which logs should be uploaded.
	Target int64 `json:"target" api:"required"`
	// Enables or disables the config.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// If set to true, the config will be applied to all CDN resources. If set to
	// false, the config will be applied to the resources specified in the `resources`
	// field.
	ForAllResources param.Opt[bool] `json:"for_all_resources,omitzero"`
	// List of resource IDs to which the config should be applied.
	Resources []int64 `json:"resources,omitzero"`
	paramObj
}

func (r LogsUploaderConfigReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderConfigReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderConfigReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
