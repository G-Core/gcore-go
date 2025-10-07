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

// LogsUploaderPolicyService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogsUploaderPolicyService] method instead.
type LogsUploaderPolicyService struct {
	Options []option.RequestOption
}

// NewLogsUploaderPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogsUploaderPolicyService(opts ...option.RequestOption) (r LogsUploaderPolicyService) {
	r = LogsUploaderPolicyService{}
	r.Options = opts
	return
}

// Create logs uploader policy.
func (r *LogsUploaderPolicyService) New(ctx context.Context, body LogsUploaderPolicyNewParams, opts ...option.RequestOption) (res *LogsUploaderPolicy, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/logs_uploader/policies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change logs uploader policy partially.
func (r *LogsUploaderPolicyService) Update(ctx context.Context, id int64, body LogsUploaderPolicyUpdateParams, opts ...option.RequestOption) (res *LogsUploaderPolicy, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/policies/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get list of logs uploader policies.
func (r *LogsUploaderPolicyService) List(ctx context.Context, query LogsUploaderPolicyListParams, opts ...option.RequestOption) (res *LogsUploaderPolicyList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/logs_uploader/policies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete the logs uploader policy from the system permanently.
//
// Notes:
//
//   - **Irreversibility**: This action is irreversible. Once deleted, the logs
//     uploader policy cannot be recovered.
func (r *LogsUploaderPolicyService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("cdn/logs_uploader/policies/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get information about logs uploader policy.
func (r *LogsUploaderPolicyService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *LogsUploaderPolicy, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/policies/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get list of available fields for logs uploader policy.
func (r *LogsUploaderPolicyService) ListFields(ctx context.Context, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/logs_uploader/policies/fields"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change logs uploader policy.
func (r *LogsUploaderPolicyService) Replace(ctx context.Context, id int64, body LogsUploaderPolicyReplaceParams, opts ...option.RequestOption) (res *LogsUploaderPolicy, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/policies/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type LogsUploaderPolicy struct {
	ID int64 `json:"id"`
	// Client that owns the policy.
	ClientID int64 `json:"client_id"`
	// Time when logs uploader policy was created.
	Created time.Time `json:"created" format:"date-time"`
	// Date format for logs.
	DateFormat string `json:"date_format"`
	// Description of the policy.
	Description string `json:"description"`
	// Field delimiter for logs.
	FieldDelimiter string `json:"field_delimiter"`
	// Field separator for logs.
	FieldSeparator string `json:"field_separator"`
	// List of fields to include in logs.
	Fields []string `json:"fields"`
	// Template for log file name.
	FileNameTemplate string `json:"file_name_template"`
	// Format type for logs.
	FormatType string `json:"format_type"`
	// Include empty logs in the upload.
	IncludeEmptyLogs bool `json:"include_empty_logs"`
	// Include logs from origin shielding in the upload.
	IncludeShieldLogs bool `json:"include_shield_logs"`
	// Name of the policy.
	Name string `json:"name"`
	// List of logs uploader configs that use this policy.
	RelatedUploaderConfigs []int64 `json:"related_uploader_configs"`
	// Interval in minutes to retry failed uploads.
	RetryIntervalMinutes int64 `json:"retry_interval_minutes"`
	// Interval in minutes to rotate logs.
	RotateIntervalMinutes int64 `json:"rotate_interval_minutes"`
	// Threshold in lines to rotate logs.
	RotateThresholdLines int64 `json:"rotate_threshold_lines"`
	// Threshold in MB to rotate logs.
	RotateThresholdMB int64 `json:"rotate_threshold_mb,nullable"`
	// Tags allow for dynamic decoration of logs by adding predefined fields to the log
	// format. These tags serve as customizable key-value pairs that can be included in
	// log entries to enhance context and readability.
	Tags map[string]string `json:"tags"`
	// Time when logs uploader policy was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		ClientID               respjson.Field
		Created                respjson.Field
		DateFormat             respjson.Field
		Description            respjson.Field
		FieldDelimiter         respjson.Field
		FieldSeparator         respjson.Field
		Fields                 respjson.Field
		FileNameTemplate       respjson.Field
		FormatType             respjson.Field
		IncludeEmptyLogs       respjson.Field
		IncludeShieldLogs      respjson.Field
		Name                   respjson.Field
		RelatedUploaderConfigs respjson.Field
		RetryIntervalMinutes   respjson.Field
		RotateIntervalMinutes  respjson.Field
		RotateThresholdLines   respjson.Field
		RotateThresholdMB      respjson.Field
		Tags                   respjson.Field
		Updated                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderPolicy) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderPolicyList []LogsUploaderPolicy

type LogsUploaderPolicyNewParams struct {
	// Threshold in MB to rotate logs.
	RotateThresholdMB param.Opt[int64] `json:"rotate_threshold_mb,omitzero"`
	// Date format for logs.
	DateFormat param.Opt[string] `json:"date_format,omitzero"`
	// Description of the policy.
	Description param.Opt[string] `json:"description,omitzero"`
	// Field delimiter for logs.
	FieldDelimiter param.Opt[string] `json:"field_delimiter,omitzero"`
	// Field separator for logs.
	FieldSeparator param.Opt[string] `json:"field_separator,omitzero"`
	// Template for log file name.
	FileNameTemplate param.Opt[string] `json:"file_name_template,omitzero"`
	// Format type for logs.
	FormatType param.Opt[string] `json:"format_type,omitzero"`
	// Include empty logs in the upload.
	IncludeEmptyLogs param.Opt[bool] `json:"include_empty_logs,omitzero"`
	// Include logs from origin shielding in the upload.
	IncludeShieldLogs param.Opt[bool] `json:"include_shield_logs,omitzero"`
	// Name of the policy.
	Name param.Opt[string] `json:"name,omitzero"`
	// Interval in minutes to retry failed uploads.
	RetryIntervalMinutes param.Opt[int64] `json:"retry_interval_minutes,omitzero"`
	// Interval in minutes to rotate logs.
	RotateIntervalMinutes param.Opt[int64] `json:"rotate_interval_minutes,omitzero"`
	// Threshold in lines to rotate logs.
	RotateThresholdLines param.Opt[int64] `json:"rotate_threshold_lines,omitzero"`
	// List of fields to include in logs.
	Fields []string `json:"fields,omitzero"`
	// Tags allow for dynamic decoration of logs by adding predefined fields to the log
	// format. These tags serve as customizable key-value pairs that can be included in
	// log entries to enhance context and readability.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r LogsUploaderPolicyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderPolicyUpdateParams struct {
	// Threshold in MB to rotate logs.
	RotateThresholdMB param.Opt[int64] `json:"rotate_threshold_mb,omitzero"`
	// Date format for logs.
	DateFormat param.Opt[string] `json:"date_format,omitzero"`
	// Description of the policy.
	Description param.Opt[string] `json:"description,omitzero"`
	// Field delimiter for logs.
	FieldDelimiter param.Opt[string] `json:"field_delimiter,omitzero"`
	// Field separator for logs.
	FieldSeparator param.Opt[string] `json:"field_separator,omitzero"`
	// Template for log file name.
	FileNameTemplate param.Opt[string] `json:"file_name_template,omitzero"`
	// Format type for logs.
	FormatType param.Opt[string] `json:"format_type,omitzero"`
	// Include empty logs in the upload.
	IncludeEmptyLogs param.Opt[bool] `json:"include_empty_logs,omitzero"`
	// Include logs from origin shielding in the upload.
	IncludeShieldLogs param.Opt[bool] `json:"include_shield_logs,omitzero"`
	// Name of the policy.
	Name param.Opt[string] `json:"name,omitzero"`
	// Interval in minutes to retry failed uploads.
	RetryIntervalMinutes param.Opt[int64] `json:"retry_interval_minutes,omitzero"`
	// Interval in minutes to rotate logs.
	RotateIntervalMinutes param.Opt[int64] `json:"rotate_interval_minutes,omitzero"`
	// Threshold in lines to rotate logs.
	RotateThresholdLines param.Opt[int64] `json:"rotate_threshold_lines,omitzero"`
	// List of fields to include in logs.
	Fields []string `json:"fields,omitzero"`
	// Tags allow for dynamic decoration of logs by adding predefined fields to the log
	// format. These tags serve as customizable key-value pairs that can be included in
	// log entries to enhance context and readability.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r LogsUploaderPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderPolicyListParams struct {
	// Search by policy name or id.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Filter by ids of related logs uploader configs that use given policy.
	ConfigIDs []int64 `query:"config_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LogsUploaderPolicyListParams]'s query parameters as
// `url.Values`.
func (r LogsUploaderPolicyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogsUploaderPolicyReplaceParams struct {
	// Threshold in MB to rotate logs.
	RotateThresholdMB param.Opt[int64] `json:"rotate_threshold_mb,omitzero"`
	// Date format for logs.
	DateFormat param.Opt[string] `json:"date_format,omitzero"`
	// Description of the policy.
	Description param.Opt[string] `json:"description,omitzero"`
	// Field delimiter for logs.
	FieldDelimiter param.Opt[string] `json:"field_delimiter,omitzero"`
	// Field separator for logs.
	FieldSeparator param.Opt[string] `json:"field_separator,omitzero"`
	// Template for log file name.
	FileNameTemplate param.Opt[string] `json:"file_name_template,omitzero"`
	// Format type for logs.
	FormatType param.Opt[string] `json:"format_type,omitzero"`
	// Include empty logs in the upload.
	IncludeEmptyLogs param.Opt[bool] `json:"include_empty_logs,omitzero"`
	// Include logs from origin shielding in the upload.
	IncludeShieldLogs param.Opt[bool] `json:"include_shield_logs,omitzero"`
	// Name of the policy.
	Name param.Opt[string] `json:"name,omitzero"`
	// Interval in minutes to retry failed uploads.
	RetryIntervalMinutes param.Opt[int64] `json:"retry_interval_minutes,omitzero"`
	// Interval in minutes to rotate logs.
	RotateIntervalMinutes param.Opt[int64] `json:"rotate_interval_minutes,omitzero"`
	// Threshold in lines to rotate logs.
	RotateThresholdLines param.Opt[int64] `json:"rotate_threshold_lines,omitzero"`
	// List of fields to include in logs.
	Fields []string `json:"fields,omitzero"`
	// Tags allow for dynamic decoration of logs by adding predefined fields to the log
	// format. These tags serve as customizable key-value pairs that can be included in
	// log entries to enhance context and readability.
	Tags map[string]string `json:"tags,omitzero"`
	paramObj
}

func (r LogsUploaderPolicyReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
