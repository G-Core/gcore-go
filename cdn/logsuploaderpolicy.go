// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

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
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/G-Core/gcore-go/shared/constant"
)

// Logs uploader policies define how CDN logs are formatted and delivered,
// including field selection, field ordering, delimiters, delivery frequency, and
// file size limits.
//
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
	return res, err
}

// Change logs uploader policy partially.
func (r *LogsUploaderPolicyService) Update(ctx context.Context, id int64, body LogsUploaderPolicyUpdateParams, opts ...option.RequestOption) (res *LogsUploaderPolicy, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/policies/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get list of logs uploader policies.
func (r *LogsUploaderPolicyService) List(ctx context.Context, query LogsUploaderPolicyListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[LogsUploaderPolicy], err error) {
	// CUSTOM CODE: CDN API only envelopes when limit>=1; default so List always paginates.
	if !query.Limit.Valid() {
		query.Limit = param.NewOpt[int64](1000)
	}
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cdn/logs_uploader/policies"
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

// Get list of logs uploader policies.
func (r *LogsUploaderPolicyService) ListAutoPaging(ctx context.Context, query LogsUploaderPolicyListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[LogsUploaderPolicy] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Delete the logs uploader policy from the system permanently.
//
// Notes:
//
//   - **Irreversibility**: This action is irreversible. Once deleted, the logs
//     uploader policy cannot be recovered.
func (r *LogsUploaderPolicyService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("cdn/logs_uploader/policies/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get information about logs uploader policy.
func (r *LogsUploaderPolicyService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *LogsUploaderPolicy, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/policies/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get list of available fields for logs uploader policy.
//
// `/cdn/v2/logs_uploader/policies/fields` returns the same fields together with
// the conversion types each one permits.
func (r *LogsUploaderPolicyService) ListFields(ctx context.Context, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/logs_uploader/policies/fields"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the available fields for a logs uploader policy, each with the conversion
// types it permits in `field_conversions`.
//
// Supersedes `/cdn/logs_uploader/policies/fields`, which returns field names only.
//
// `-` is the placeholder for a skipped column: it may appear in `fields` and
// permits no conversion.
func (r *LogsUploaderPolicyService) ListFieldsAllowedConversions(ctx context.Context, opts ...option.RequestOption) (res *[]LogsUploaderPolicyField, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/v2/logs_uploader/policies/fields"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Change logs uploader policy.
func (r *LogsUploaderPolicyService) Replace(ctx context.Context, id int64, body LogsUploaderPolicyReplaceParams, opts ...option.RequestOption) (res *LogsUploaderPolicy, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/logs_uploader/policies/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
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
	// When set to true, the service sanitizes string values by escaping characters
	// that may be unsafe for transport, logging, or downstream processing.
	//
	// The following categories of characters are escaped:
	//
	// - Control and non-printable characters
	// - Quotation marks and escape characters
	// - Characters outside the standard ASCII range
	//
	// The resulting output contains only printable ASCII characters.
	EscapeSpecialCharacters bool `json:"escape_special_characters"`
	// Per-field value conversions for exported logs. Maps a canonical Gcore field name
	// to the pipeline applied to its values. Field names are limited to 255 characters
	// and must not be empty. Each key must be present in `fields`, and each conversion
	// type must be listed in that field's `allowed_conversions` from
	// `/cdn/v2/logs_uploader/policies/fields`. Conversions in a pipeline are applied
	// in array order. Values are converted independently of `field_remap`, which
	// renames the exported field: both are keyed on the canonical field name.
	FieldConversions map[string]LogsUploaderPolicyFieldConversion `json:"field_conversions"`
	// Field delimiter for logs.
	FieldDelimiter string `json:"field_delimiter"`
	// Per-field output-name remap for exported logs. Maps a canonical Gcore field name
	// (from `/cdn/logs_uploader/policies/fields`, and must be present in `fields`) to
	// the field name it should have in the exported logs. Unmapped fields keep their
	// canonical name. Output names (after remapping) must be unique.
	FieldRemap map[string]string `json:"field_remap"`
	// Field separator for logs.
	FieldSeparator string `json:"field_separator"`
	// List of fields to include in logs. Duplicate names are allowed for plain text
	// output, but rejected when `format_type` is `json` or a `field_remap` is set
	// (each field becomes a distinct output key).
	Fields []string `json:"fields"`
	// Template for log file name.
	FileNameTemplate string `json:"file_name_template"`
	// Format type for logs.
	//
	// Possible values:
	//
	//   - **""** - empty, it means it will apply the format configurations from the
	//     policy.
	//   - **"json"** - output the logs as json lines.
	//
	// Any of "json", "".
	FormatType LogsUploaderPolicyFormatType `json:"format_type"`
	// Include empty logs in the upload.
	IncludeEmptyLogs bool `json:"include_empty_logs"`
	// Include logs from origin shielding in the upload.
	IncludeShieldLogs bool `json:"include_shield_logs"`
	// Sampling rate for logs. A value between 0 and 1 that determines the fraction of
	// log entries to collect.
	//
	//   - **1** - collect all logs (default).
	//   - **0.5** - collect approximately 50% of logs.
	//   - **0** - collect no logs (effectively disables logging without removing the
	//     policy).
	LogSampleRate float64 `json:"log_sample_rate"`
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
	RotateThresholdMB int64 `json:"rotate_threshold_mb" api:"nullable"`
	// Tags allow for dynamic decoration of logs by adding predefined fields to the log
	// format. These tags serve as customizable key-value pairs that can be included in
	// log entries to enhance context and readability.
	Tags map[string]string `json:"tags"`
	// Time when logs uploader policy was updated.
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		ClientID                respjson.Field
		Created                 respjson.Field
		DateFormat              respjson.Field
		Description             respjson.Field
		EscapeSpecialCharacters respjson.Field
		FieldConversions        respjson.Field
		FieldDelimiter          respjson.Field
		FieldRemap              respjson.Field
		FieldSeparator          respjson.Field
		Fields                  respjson.Field
		FileNameTemplate        respjson.Field
		FormatType              respjson.Field
		IncludeEmptyLogs        respjson.Field
		IncludeShieldLogs       respjson.Field
		LogSampleRate           respjson.Field
		Name                    respjson.Field
		RelatedUploaderConfigs  respjson.Field
		RetryIntervalMinutes    respjson.Field
		RotateIntervalMinutes   respjson.Field
		RotateThresholdLines    respjson.Field
		RotateThresholdMB       respjson.Field
		Tags                    respjson.Field
		Updated                 respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderPolicy) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderPolicyFieldConversion struct {
	Conversions []LogsUploaderPolicyFieldConversionConversionUnion `json:"conversions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conversions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderPolicyFieldConversion) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderPolicyFieldConversion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LogsUploaderPolicyFieldConversionConversionUnion contains all possible
// properties and values from [LogsUploaderPolicyFieldConversionConversionScale],
// [LogsUploaderPolicyFieldConversionConversionReplace].
//
// Use the [LogsUploaderPolicyFieldConversionConversionUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LogsUploaderPolicyFieldConversionConversionUnion struct {
	// This field is a union of
	// [LogsUploaderPolicyFieldConversionConversionScaleConfig],
	// [LogsUploaderPolicyFieldConversionConversionReplaceConfig]
	Config LogsUploaderPolicyFieldConversionConversionUnionConfig `json:"config"`
	// Any of "scale", "replace".
	Type string `json:"type"`
	JSON struct {
		Config respjson.Field
		Type   respjson.Field
		raw    string
	} `json:"-"`
}

// anyLogsUploaderPolicyFieldConversionConversion is implemented by each variant of
// [LogsUploaderPolicyFieldConversionConversionUnion] to add type safety for the
// return type of [LogsUploaderPolicyFieldConversionConversionUnion.AsAny]
type anyLogsUploaderPolicyFieldConversionConversion interface {
	implLogsUploaderPolicyFieldConversionConversionUnion()
}

func (LogsUploaderPolicyFieldConversionConversionScale) implLogsUploaderPolicyFieldConversionConversionUnion() {
}
func (LogsUploaderPolicyFieldConversionConversionReplace) implLogsUploaderPolicyFieldConversionConversionUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := LogsUploaderPolicyFieldConversionConversionUnion.AsAny().(type) {
//	case cdn.LogsUploaderPolicyFieldConversionConversionScale:
//	case cdn.LogsUploaderPolicyFieldConversionConversionReplace:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u LogsUploaderPolicyFieldConversionConversionUnion) AsAny() anyLogsUploaderPolicyFieldConversionConversion {
	switch u.Type {
	case "scale":
		return u.AsScale()
	case "replace":
		return u.AsReplace()
	}
	return nil
}

func (u LogsUploaderPolicyFieldConversionConversionUnion) AsScale() (v LogsUploaderPolicyFieldConversionConversionScale) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogsUploaderPolicyFieldConversionConversionUnion) AsReplace() (v LogsUploaderPolicyFieldConversionConversionReplace) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LogsUploaderPolicyFieldConversionConversionUnion) RawJSON() string { return u.JSON.raw }

func (r *LogsUploaderPolicyFieldConversionConversionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LogsUploaderPolicyFieldConversionConversionUnionConfig is an implicit subunion
// of [LogsUploaderPolicyFieldConversionConversionUnion].
// LogsUploaderPolicyFieldConversionConversionUnionConfig provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [LogsUploaderPolicyFieldConversionConversionUnion].
type LogsUploaderPolicyFieldConversionConversionUnionConfig struct {
	// This field is from variant
	// [LogsUploaderPolicyFieldConversionConversionScaleConfig].
	Factor float64 `json:"factor"`
	// This field is from variant
	// [LogsUploaderPolicyFieldConversionConversionScaleConfig].
	Precision int64 `json:"precision"`
	// This field is from variant
	// [LogsUploaderPolicyFieldConversionConversionScaleConfig].
	Rounding string `json:"rounding"`
	// This field is from variant
	// [LogsUploaderPolicyFieldConversionConversionReplaceConfig].
	Values map[string]string `json:"values"`
	// This field is from variant
	// [LogsUploaderPolicyFieldConversionConversionReplaceConfig].
	Default string `json:"default"`
	JSON    struct {
		Factor    respjson.Field
		Precision respjson.Field
		Rounding  respjson.Field
		Values    respjson.Field
		Default   respjson.Field
		raw       string
	} `json:"-"`
}

func (r *LogsUploaderPolicyFieldConversionConversionUnionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderPolicyFieldConversionConversionScale struct {
	Config LogsUploaderPolicyFieldConversionConversionScaleConfig `json:"config" api:"required"`
	Type   constant.Scale                                         `json:"type" default:"scale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderPolicyFieldConversionConversionScale) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderPolicyFieldConversionConversionScale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderPolicyFieldConversionConversionScaleConfig struct {
	// Multiplier applied to the field value.
	Factor float64 `json:"factor" api:"required"`
	// Optional number of decimal places in the converted value. Must be specified
	// together with `rounding`; returned as `null` when unspecified.
	Precision int64 `json:"precision" api:"nullable"`
	// Optional rounding mode. Must be specified together with `precision`; returned as
	// `null` when unspecified.
	//
	// Any of "nearest", "down", "up".
	Rounding string `json:"rounding" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Factor      respjson.Field
		Precision   respjson.Field
		Rounding    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderPolicyFieldConversionConversionScaleConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderPolicyFieldConversionConversionScaleConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderPolicyFieldConversionConversionReplace struct {
	Config LogsUploaderPolicyFieldConversionConversionReplaceConfig `json:"config" api:"required"`
	Type   constant.Replace                                         `json:"type" default:"replace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderPolicyFieldConversionConversionReplace) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderPolicyFieldConversionConversionReplace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderPolicyFieldConversionConversionReplaceConfig struct {
	// Exact, case-sensitive replacements, matched and written without trimming. Keys
	// must not be empty and are limited to 255 characters; values are limited to 100
	// characters and may be empty.
	Values map[string]string `json:"values" api:"required"`
	// Value used when the input does not match any key in `values`.
	Default string `json:"default"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Values      respjson.Field
		Default     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderPolicyFieldConversionConversionReplaceConfig) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderPolicyFieldConversionConversionReplaceConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Format type for logs.
//
// Possible values:
//
//   - **""** - empty, it means it will apply the format configurations from the
//     policy.
//   - **"json"** - output the logs as json lines.
type LogsUploaderPolicyFormatType string

const (
	LogsUploaderPolicyFormatTypeJson  LogsUploaderPolicyFormatType = "json"
	LogsUploaderPolicyFormatTypeEmpty LogsUploaderPolicyFormatType = ""
)

type LogsUploaderPolicyField struct {
	// Conversion types this field permits in `field_conversions`. Empty when the field
	// permits none.
	//
	// Any of "scale", "replace".
	AllowedConversions []string `json:"allowed_conversions" api:"required"`
	// Canonical Gcore field name, or `-` as the skipped-column placeholder, selectable
	// in a policy's `fields`.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedConversions respjson.Field
		Name               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderPolicyField) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderPolicyField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderPolicyList struct {
	// Total number of items.
	Count int64 `json:"count" api:"required"`
	// URL to the next page of results. Null if current page is the last one.
	Next string `json:"next" api:"required"`
	// URL to the previous page of results. Null if current page is the first one.
	Previous string               `json:"previous" api:"required"`
	Results  []LogsUploaderPolicy `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Next        respjson.Field
		Previous    respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsUploaderPolicyList) RawJSON() string { return r.JSON.raw }
func (r *LogsUploaderPolicyList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsUploaderPolicyNewParams struct {
	// Threshold in MB to rotate logs.
	RotateThresholdMB param.Opt[int64] `json:"rotate_threshold_mb,omitzero"`
	// Date format for logs.
	DateFormat param.Opt[string] `json:"date_format,omitzero"`
	// Description of the policy.
	Description param.Opt[string] `json:"description,omitzero"`
	// When set to true, the service sanitizes string values by escaping characters
	// that may be unsafe for transport, logging, or downstream processing.
	//
	// The following categories of characters are escaped:
	//
	// - Control and non-printable characters
	// - Quotation marks and escape characters
	// - Characters outside the standard ASCII range
	//
	// The resulting output contains only printable ASCII characters.
	EscapeSpecialCharacters param.Opt[bool] `json:"escape_special_characters,omitzero"`
	// Field delimiter for logs.
	FieldDelimiter param.Opt[string] `json:"field_delimiter,omitzero"`
	// Field separator for logs.
	FieldSeparator param.Opt[string] `json:"field_separator,omitzero"`
	// Template for log file name.
	FileNameTemplate param.Opt[string] `json:"file_name_template,omitzero"`
	// Include empty logs in the upload.
	IncludeEmptyLogs param.Opt[bool] `json:"include_empty_logs,omitzero"`
	// Include logs from origin shielding in the upload.
	IncludeShieldLogs param.Opt[bool] `json:"include_shield_logs,omitzero"`
	// Sampling rate for logs. A value between 0 and 1 that determines the fraction of
	// log entries to collect.
	//
	//   - **1** - collect all logs (default).
	//   - **0.5** - collect approximately 50% of logs.
	//   - **0** - collect no logs (effectively disables logging without removing the
	//     policy).
	LogSampleRate param.Opt[float64] `json:"log_sample_rate,omitzero"`
	// Name of the policy.
	Name param.Opt[string] `json:"name,omitzero"`
	// Interval in minutes to retry failed uploads.
	RetryIntervalMinutes param.Opt[int64] `json:"retry_interval_minutes,omitzero"`
	// Interval in minutes to rotate logs.
	RotateIntervalMinutes param.Opt[int64] `json:"rotate_interval_minutes,omitzero"`
	// Threshold in lines to rotate logs.
	RotateThresholdLines param.Opt[int64] `json:"rotate_threshold_lines,omitzero"`
	// Per-field value conversions for exported logs. Maps a canonical Gcore field name
	// to the pipeline applied to its values. Field names are limited to 255 characters
	// and must not be empty. Each key must be present in `fields`, and each conversion
	// type must be listed in that field's `allowed_conversions` from
	// `/cdn/v2/logs_uploader/policies/fields`. Conversions in a pipeline are applied
	// in array order. Values are converted independently of `field_remap`, which
	// renames the exported field: both are keyed on the canonical field name.
	FieldConversions map[string]LogsUploaderPolicyNewParamsFieldConversion `json:"field_conversions,omitzero"`
	// Per-field output-name remap for exported logs. Maps a canonical Gcore field name
	// (from `/cdn/logs_uploader/policies/fields`, and must be present in `fields`) to
	// the field name it should have in the exported logs. Unmapped fields keep their
	// canonical name. Output names (after remapping) must be unique.
	FieldRemap map[string]string `json:"field_remap,omitzero"`
	// List of fields to include in logs. Duplicate names are allowed for plain text
	// output, but rejected when `format_type` is `json` or a `field_remap` is set
	// (each field becomes a distinct output key).
	Fields []string `json:"fields,omitzero"`
	// Format type for logs.
	//
	// Possible values:
	//
	//   - **""** - empty, it means it will apply the format configurations from the
	//     policy.
	//   - **"json"** - output the logs as json lines.
	//
	// Any of "json", "".
	FormatType LogsUploaderPolicyNewParamsFormatType `json:"format_type,omitzero"`
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

// The property Conversions is required.
type LogsUploaderPolicyNewParamsFieldConversion struct {
	Conversions []LogsUploaderPolicyNewParamsFieldConversionConversionUnion `json:"conversions,omitzero" api:"required"`
	paramObj
}

func (r LogsUploaderPolicyNewParamsFieldConversion) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyNewParamsFieldConversion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyNewParamsFieldConversion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LogsUploaderPolicyNewParamsFieldConversionConversionUnion struct {
	OfScale   *LogsUploaderPolicyNewParamsFieldConversionConversionScale   `json:",omitzero,inline"`
	OfReplace *LogsUploaderPolicyNewParamsFieldConversionConversionReplace `json:",omitzero,inline"`
	paramUnion
}

func (u LogsUploaderPolicyNewParamsFieldConversionConversionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfScale, u.OfReplace)
}
func (u *LogsUploaderPolicyNewParamsFieldConversionConversionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LogsUploaderPolicyNewParamsFieldConversionConversionUnion) asAny() any {
	if !param.IsOmitted(u.OfScale) {
		return u.OfScale
	} else if !param.IsOmitted(u.OfReplace) {
		return u.OfReplace
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderPolicyNewParamsFieldConversionConversionUnion) GetType() *string {
	if vt := u.OfScale; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfReplace; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u LogsUploaderPolicyNewParamsFieldConversionConversionUnion) GetConfig() (res logsUploaderPolicyNewParamsFieldConversionConversionUnionConfig) {
	if vt := u.OfScale; vt != nil {
		res.any = &vt.Config
	} else if vt := u.OfReplace; vt != nil {
		res.any = &vt.Config
	}
	return
}

// Can have the runtime types
// [*LogsUploaderPolicyNewParamsFieldConversionConversionScaleConfig],
// [*LogsUploaderPolicyNewParamsFieldConversionConversionReplaceConfig]
type logsUploaderPolicyNewParamsFieldConversionConversionUnionConfig struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cdn.LogsUploaderPolicyNewParamsFieldConversionConversionScaleConfig:
//	case *cdn.LogsUploaderPolicyNewParamsFieldConversionConversionReplaceConfig:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u logsUploaderPolicyNewParamsFieldConversionConversionUnionConfig) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[LogsUploaderPolicyNewParamsFieldConversionConversionUnion](
		"type",
		apijson.Discriminator[LogsUploaderPolicyNewParamsFieldConversionConversionScale]("scale"),
		apijson.Discriminator[LogsUploaderPolicyNewParamsFieldConversionConversionReplace]("replace"),
	)
}

// The properties Config, Type are required.
type LogsUploaderPolicyNewParamsFieldConversionConversionScale struct {
	Config LogsUploaderPolicyNewParamsFieldConversionConversionScaleConfig `json:"config,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "scale".
	Type constant.Scale `json:"type" default:"scale"`
	paramObj
}

func (r LogsUploaderPolicyNewParamsFieldConversionConversionScale) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyNewParamsFieldConversionConversionScale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyNewParamsFieldConversionConversionScale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Factor is required.
type LogsUploaderPolicyNewParamsFieldConversionConversionScaleConfig struct {
	// Multiplier applied to the field value.
	Factor float64 `json:"factor" api:"required"`
	// Optional number of decimal places in the converted value. Must be specified
	// together with `rounding`; returned as `null` when unspecified.
	Precision param.Opt[int64] `json:"precision,omitzero"`
	// Optional rounding mode. Must be specified together with `precision`; returned as
	// `null` when unspecified.
	//
	// Any of "nearest", "down", "up".
	Rounding string `json:"rounding,omitzero"`
	paramObj
}

func (r LogsUploaderPolicyNewParamsFieldConversionConversionScaleConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyNewParamsFieldConversionConversionScaleConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyNewParamsFieldConversionConversionScaleConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderPolicyNewParamsFieldConversionConversionScaleConfig](
		"rounding", "nearest", "down", "up",
	)
}

// The properties Config, Type are required.
type LogsUploaderPolicyNewParamsFieldConversionConversionReplace struct {
	Config LogsUploaderPolicyNewParamsFieldConversionConversionReplaceConfig `json:"config,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "replace".
	Type constant.Replace `json:"type" default:"replace"`
	paramObj
}

func (r LogsUploaderPolicyNewParamsFieldConversionConversionReplace) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyNewParamsFieldConversionConversionReplace
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyNewParamsFieldConversionConversionReplace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Values is required.
type LogsUploaderPolicyNewParamsFieldConversionConversionReplaceConfig struct {
	// Exact, case-sensitive replacements, matched and written without trimming. Keys
	// must not be empty and are limited to 255 characters; values are limited to 100
	// characters and may be empty.
	Values map[string]string `json:"values,omitzero" api:"required"`
	// Value used when the input does not match any key in `values`.
	Default param.Opt[string] `json:"default,omitzero"`
	paramObj
}

func (r LogsUploaderPolicyNewParamsFieldConversionConversionReplaceConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyNewParamsFieldConversionConversionReplaceConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyNewParamsFieldConversionConversionReplaceConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Format type for logs.
//
// Possible values:
//
//   - **""** - empty, it means it will apply the format configurations from the
//     policy.
//   - **"json"** - output the logs as json lines.
type LogsUploaderPolicyNewParamsFormatType string

const (
	LogsUploaderPolicyNewParamsFormatTypeJson  LogsUploaderPolicyNewParamsFormatType = "json"
	LogsUploaderPolicyNewParamsFormatTypeEmpty LogsUploaderPolicyNewParamsFormatType = ""
)

type LogsUploaderPolicyUpdateParams struct {
	// Threshold in MB to rotate logs.
	RotateThresholdMB param.Opt[int64] `json:"rotate_threshold_mb,omitzero"`
	// Date format for logs.
	DateFormat param.Opt[string] `json:"date_format,omitzero"`
	// Description of the policy.
	Description param.Opt[string] `json:"description,omitzero"`
	// When set to true, the service sanitizes string values by escaping characters
	// that may be unsafe for transport, logging, or downstream processing.
	//
	// The following categories of characters are escaped:
	//
	// - Control and non-printable characters
	// - Quotation marks and escape characters
	// - Characters outside the standard ASCII range
	//
	// The resulting output contains only printable ASCII characters.
	EscapeSpecialCharacters param.Opt[bool] `json:"escape_special_characters,omitzero"`
	// Field delimiter for logs.
	FieldDelimiter param.Opt[string] `json:"field_delimiter,omitzero"`
	// Field separator for logs.
	FieldSeparator param.Opt[string] `json:"field_separator,omitzero"`
	// Template for log file name.
	FileNameTemplate param.Opt[string] `json:"file_name_template,omitzero"`
	// Include empty logs in the upload.
	IncludeEmptyLogs param.Opt[bool] `json:"include_empty_logs,omitzero"`
	// Include logs from origin shielding in the upload.
	IncludeShieldLogs param.Opt[bool] `json:"include_shield_logs,omitzero"`
	// Sampling rate for logs. A value between 0 and 1 that determines the fraction of
	// log entries to collect.
	//
	//   - **1** - collect all logs (default).
	//   - **0.5** - collect approximately 50% of logs.
	//   - **0** - collect no logs (effectively disables logging without removing the
	//     policy).
	LogSampleRate param.Opt[float64] `json:"log_sample_rate,omitzero"`
	// Name of the policy.
	Name param.Opt[string] `json:"name,omitzero"`
	// Interval in minutes to retry failed uploads.
	RetryIntervalMinutes param.Opt[int64] `json:"retry_interval_minutes,omitzero"`
	// Interval in minutes to rotate logs.
	RotateIntervalMinutes param.Opt[int64] `json:"rotate_interval_minutes,omitzero"`
	// Threshold in lines to rotate logs.
	RotateThresholdLines param.Opt[int64] `json:"rotate_threshold_lines,omitzero"`
	// Per-field value conversions for exported logs. Maps a canonical Gcore field name
	// to the pipeline applied to its values. Field names are limited to 255 characters
	// and must not be empty. Each key must be present in `fields`, and each conversion
	// type must be listed in that field's `allowed_conversions` from
	// `/cdn/v2/logs_uploader/policies/fields`. Conversions in a pipeline are applied
	// in array order. Values are converted independently of `field_remap`, which
	// renames the exported field: both are keyed on the canonical field name.
	FieldConversions map[string]LogsUploaderPolicyUpdateParamsFieldConversion `json:"field_conversions,omitzero"`
	// Per-field output-name remap for exported logs. Maps a canonical Gcore field name
	// (from `/cdn/logs_uploader/policies/fields`, and must be present in `fields`) to
	// the field name it should have in the exported logs. Unmapped fields keep their
	// canonical name. Output names (after remapping) must be unique.
	FieldRemap map[string]string `json:"field_remap,omitzero"`
	// List of fields to include in logs. Duplicate names are allowed for plain text
	// output, but rejected when `format_type` is `json` or a `field_remap` is set
	// (each field becomes a distinct output key).
	Fields []string `json:"fields,omitzero"`
	// Format type for logs.
	//
	// Possible values:
	//
	//   - **""** - empty, it means it will apply the format configurations from the
	//     policy.
	//   - **"json"** - output the logs as json lines.
	//
	// Any of "json", "".
	FormatType LogsUploaderPolicyUpdateParamsFormatType `json:"format_type,omitzero"`
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

// The property Conversions is required.
type LogsUploaderPolicyUpdateParamsFieldConversion struct {
	Conversions []LogsUploaderPolicyUpdateParamsFieldConversionConversionUnion `json:"conversions,omitzero" api:"required"`
	paramObj
}

func (r LogsUploaderPolicyUpdateParamsFieldConversion) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyUpdateParamsFieldConversion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyUpdateParamsFieldConversion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LogsUploaderPolicyUpdateParamsFieldConversionConversionUnion struct {
	OfScale   *LogsUploaderPolicyUpdateParamsFieldConversionConversionScale   `json:",omitzero,inline"`
	OfReplace *LogsUploaderPolicyUpdateParamsFieldConversionConversionReplace `json:",omitzero,inline"`
	paramUnion
}

func (u LogsUploaderPolicyUpdateParamsFieldConversionConversionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfScale, u.OfReplace)
}
func (u *LogsUploaderPolicyUpdateParamsFieldConversionConversionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LogsUploaderPolicyUpdateParamsFieldConversionConversionUnion) asAny() any {
	if !param.IsOmitted(u.OfScale) {
		return u.OfScale
	} else if !param.IsOmitted(u.OfReplace) {
		return u.OfReplace
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderPolicyUpdateParamsFieldConversionConversionUnion) GetType() *string {
	if vt := u.OfScale; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfReplace; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u LogsUploaderPolicyUpdateParamsFieldConversionConversionUnion) GetConfig() (res logsUploaderPolicyUpdateParamsFieldConversionConversionUnionConfig) {
	if vt := u.OfScale; vt != nil {
		res.any = &vt.Config
	} else if vt := u.OfReplace; vt != nil {
		res.any = &vt.Config
	}
	return
}

// Can have the runtime types
// [*LogsUploaderPolicyUpdateParamsFieldConversionConversionScaleConfig],
// [*LogsUploaderPolicyUpdateParamsFieldConversionConversionReplaceConfig]
type logsUploaderPolicyUpdateParamsFieldConversionConversionUnionConfig struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cdn.LogsUploaderPolicyUpdateParamsFieldConversionConversionScaleConfig:
//	case *cdn.LogsUploaderPolicyUpdateParamsFieldConversionConversionReplaceConfig:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u logsUploaderPolicyUpdateParamsFieldConversionConversionUnionConfig) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[LogsUploaderPolicyUpdateParamsFieldConversionConversionUnion](
		"type",
		apijson.Discriminator[LogsUploaderPolicyUpdateParamsFieldConversionConversionScale]("scale"),
		apijson.Discriminator[LogsUploaderPolicyUpdateParamsFieldConversionConversionReplace]("replace"),
	)
}

// The properties Config, Type are required.
type LogsUploaderPolicyUpdateParamsFieldConversionConversionScale struct {
	Config LogsUploaderPolicyUpdateParamsFieldConversionConversionScaleConfig `json:"config,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "scale".
	Type constant.Scale `json:"type" default:"scale"`
	paramObj
}

func (r LogsUploaderPolicyUpdateParamsFieldConversionConversionScale) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyUpdateParamsFieldConversionConversionScale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyUpdateParamsFieldConversionConversionScale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Factor is required.
type LogsUploaderPolicyUpdateParamsFieldConversionConversionScaleConfig struct {
	// Multiplier applied to the field value.
	Factor float64 `json:"factor" api:"required"`
	// Optional number of decimal places in the converted value. Must be specified
	// together with `rounding`; returned as `null` when unspecified.
	Precision param.Opt[int64] `json:"precision,omitzero"`
	// Optional rounding mode. Must be specified together with `precision`; returned as
	// `null` when unspecified.
	//
	// Any of "nearest", "down", "up".
	Rounding string `json:"rounding,omitzero"`
	paramObj
}

func (r LogsUploaderPolicyUpdateParamsFieldConversionConversionScaleConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyUpdateParamsFieldConversionConversionScaleConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyUpdateParamsFieldConversionConversionScaleConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderPolicyUpdateParamsFieldConversionConversionScaleConfig](
		"rounding", "nearest", "down", "up",
	)
}

// The properties Config, Type are required.
type LogsUploaderPolicyUpdateParamsFieldConversionConversionReplace struct {
	Config LogsUploaderPolicyUpdateParamsFieldConversionConversionReplaceConfig `json:"config,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "replace".
	Type constant.Replace `json:"type" default:"replace"`
	paramObj
}

func (r LogsUploaderPolicyUpdateParamsFieldConversionConversionReplace) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyUpdateParamsFieldConversionConversionReplace
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyUpdateParamsFieldConversionConversionReplace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Values is required.
type LogsUploaderPolicyUpdateParamsFieldConversionConversionReplaceConfig struct {
	// Exact, case-sensitive replacements, matched and written without trimming. Keys
	// must not be empty and are limited to 255 characters; values are limited to 100
	// characters and may be empty.
	Values map[string]string `json:"values,omitzero" api:"required"`
	// Value used when the input does not match any key in `values`.
	Default param.Opt[string] `json:"default,omitzero"`
	paramObj
}

func (r LogsUploaderPolicyUpdateParamsFieldConversionConversionReplaceConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyUpdateParamsFieldConversionConversionReplaceConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyUpdateParamsFieldConversionConversionReplaceConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Format type for logs.
//
// Possible values:
//
//   - **""** - empty, it means it will apply the format configurations from the
//     policy.
//   - **"json"** - output the logs as json lines.
type LogsUploaderPolicyUpdateParamsFormatType string

const (
	LogsUploaderPolicyUpdateParamsFormatTypeJson  LogsUploaderPolicyUpdateParamsFormatType = "json"
	LogsUploaderPolicyUpdateParamsFormatTypeEmpty LogsUploaderPolicyUpdateParamsFormatType = ""
)

type LogsUploaderPolicyListParams struct {
	// Maximum number of items to return in the response. Cannot exceed 1000.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip from the beginning of the list.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
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
	// When set to true, the service sanitizes string values by escaping characters
	// that may be unsafe for transport, logging, or downstream processing.
	//
	// The following categories of characters are escaped:
	//
	// - Control and non-printable characters
	// - Quotation marks and escape characters
	// - Characters outside the standard ASCII range
	//
	// The resulting output contains only printable ASCII characters.
	EscapeSpecialCharacters param.Opt[bool] `json:"escape_special_characters,omitzero"`
	// Field delimiter for logs.
	FieldDelimiter param.Opt[string] `json:"field_delimiter,omitzero"`
	// Field separator for logs.
	FieldSeparator param.Opt[string] `json:"field_separator,omitzero"`
	// Template for log file name.
	FileNameTemplate param.Opt[string] `json:"file_name_template,omitzero"`
	// Include empty logs in the upload.
	IncludeEmptyLogs param.Opt[bool] `json:"include_empty_logs,omitzero"`
	// Include logs from origin shielding in the upload.
	IncludeShieldLogs param.Opt[bool] `json:"include_shield_logs,omitzero"`
	// Sampling rate for logs. A value between 0 and 1 that determines the fraction of
	// log entries to collect.
	//
	//   - **1** - collect all logs (default).
	//   - **0.5** - collect approximately 50% of logs.
	//   - **0** - collect no logs (effectively disables logging without removing the
	//     policy).
	LogSampleRate param.Opt[float64] `json:"log_sample_rate,omitzero"`
	// Name of the policy.
	Name param.Opt[string] `json:"name,omitzero"`
	// Interval in minutes to retry failed uploads.
	RetryIntervalMinutes param.Opt[int64] `json:"retry_interval_minutes,omitzero"`
	// Interval in minutes to rotate logs.
	RotateIntervalMinutes param.Opt[int64] `json:"rotate_interval_minutes,omitzero"`
	// Threshold in lines to rotate logs.
	RotateThresholdLines param.Opt[int64] `json:"rotate_threshold_lines,omitzero"`
	// Per-field value conversions for exported logs. Maps a canonical Gcore field name
	// to the pipeline applied to its values. Field names are limited to 255 characters
	// and must not be empty. Each key must be present in `fields`, and each conversion
	// type must be listed in that field's `allowed_conversions` from
	// `/cdn/v2/logs_uploader/policies/fields`. Conversions in a pipeline are applied
	// in array order. Values are converted independently of `field_remap`, which
	// renames the exported field: both are keyed on the canonical field name.
	FieldConversions map[string]LogsUploaderPolicyReplaceParamsFieldConversion `json:"field_conversions,omitzero"`
	// Per-field output-name remap for exported logs. Maps a canonical Gcore field name
	// (from `/cdn/logs_uploader/policies/fields`, and must be present in `fields`) to
	// the field name it should have in the exported logs. Unmapped fields keep their
	// canonical name. Output names (after remapping) must be unique.
	FieldRemap map[string]string `json:"field_remap,omitzero"`
	// List of fields to include in logs. Duplicate names are allowed for plain text
	// output, but rejected when `format_type` is `json` or a `field_remap` is set
	// (each field becomes a distinct output key).
	Fields []string `json:"fields,omitzero"`
	// Format type for logs.
	//
	// Possible values:
	//
	//   - **""** - empty, it means it will apply the format configurations from the
	//     policy.
	//   - **"json"** - output the logs as json lines.
	//
	// Any of "json", "".
	FormatType LogsUploaderPolicyReplaceParamsFormatType `json:"format_type,omitzero"`
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

// The property Conversions is required.
type LogsUploaderPolicyReplaceParamsFieldConversion struct {
	Conversions []LogsUploaderPolicyReplaceParamsFieldConversionConversionUnion `json:"conversions,omitzero" api:"required"`
	paramObj
}

func (r LogsUploaderPolicyReplaceParamsFieldConversion) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyReplaceParamsFieldConversion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyReplaceParamsFieldConversion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type LogsUploaderPolicyReplaceParamsFieldConversionConversionUnion struct {
	OfScale   *LogsUploaderPolicyReplaceParamsFieldConversionConversionScale   `json:",omitzero,inline"`
	OfReplace *LogsUploaderPolicyReplaceParamsFieldConversionConversionReplace `json:",omitzero,inline"`
	paramUnion
}

func (u LogsUploaderPolicyReplaceParamsFieldConversionConversionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfScale, u.OfReplace)
}
func (u *LogsUploaderPolicyReplaceParamsFieldConversionConversionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *LogsUploaderPolicyReplaceParamsFieldConversionConversionUnion) asAny() any {
	if !param.IsOmitted(u.OfScale) {
		return u.OfScale
	} else if !param.IsOmitted(u.OfReplace) {
		return u.OfReplace
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u LogsUploaderPolicyReplaceParamsFieldConversionConversionUnion) GetType() *string {
	if vt := u.OfScale; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfReplace; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u LogsUploaderPolicyReplaceParamsFieldConversionConversionUnion) GetConfig() (res logsUploaderPolicyReplaceParamsFieldConversionConversionUnionConfig) {
	if vt := u.OfScale; vt != nil {
		res.any = &vt.Config
	} else if vt := u.OfReplace; vt != nil {
		res.any = &vt.Config
	}
	return
}

// Can have the runtime types
// [*LogsUploaderPolicyReplaceParamsFieldConversionConversionScaleConfig],
// [*LogsUploaderPolicyReplaceParamsFieldConversionConversionReplaceConfig]
type logsUploaderPolicyReplaceParamsFieldConversionConversionUnionConfig struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *cdn.LogsUploaderPolicyReplaceParamsFieldConversionConversionScaleConfig:
//	case *cdn.LogsUploaderPolicyReplaceParamsFieldConversionConversionReplaceConfig:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u logsUploaderPolicyReplaceParamsFieldConversionConversionUnionConfig) AsAny() any {
	return u.any
}

func init() {
	apijson.RegisterUnion[LogsUploaderPolicyReplaceParamsFieldConversionConversionUnion](
		"type",
		apijson.Discriminator[LogsUploaderPolicyReplaceParamsFieldConversionConversionScale]("scale"),
		apijson.Discriminator[LogsUploaderPolicyReplaceParamsFieldConversionConversionReplace]("replace"),
	)
}

// The properties Config, Type are required.
type LogsUploaderPolicyReplaceParamsFieldConversionConversionScale struct {
	Config LogsUploaderPolicyReplaceParamsFieldConversionConversionScaleConfig `json:"config,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "scale".
	Type constant.Scale `json:"type" default:"scale"`
	paramObj
}

func (r LogsUploaderPolicyReplaceParamsFieldConversionConversionScale) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyReplaceParamsFieldConversionConversionScale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyReplaceParamsFieldConversionConversionScale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Factor is required.
type LogsUploaderPolicyReplaceParamsFieldConversionConversionScaleConfig struct {
	// Multiplier applied to the field value.
	Factor float64 `json:"factor" api:"required"`
	// Optional number of decimal places in the converted value. Must be specified
	// together with `rounding`; returned as `null` when unspecified.
	Precision param.Opt[int64] `json:"precision,omitzero"`
	// Optional rounding mode. Must be specified together with `precision`; returned as
	// `null` when unspecified.
	//
	// Any of "nearest", "down", "up".
	Rounding string `json:"rounding,omitzero"`
	paramObj
}

func (r LogsUploaderPolicyReplaceParamsFieldConversionConversionScaleConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyReplaceParamsFieldConversionConversionScaleConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyReplaceParamsFieldConversionConversionScaleConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LogsUploaderPolicyReplaceParamsFieldConversionConversionScaleConfig](
		"rounding", "nearest", "down", "up",
	)
}

// The properties Config, Type are required.
type LogsUploaderPolicyReplaceParamsFieldConversionConversionReplace struct {
	Config LogsUploaderPolicyReplaceParamsFieldConversionConversionReplaceConfig `json:"config,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "replace".
	Type constant.Replace `json:"type" default:"replace"`
	paramObj
}

func (r LogsUploaderPolicyReplaceParamsFieldConversionConversionReplace) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyReplaceParamsFieldConversionConversionReplace
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyReplaceParamsFieldConversionConversionReplace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Values is required.
type LogsUploaderPolicyReplaceParamsFieldConversionConversionReplaceConfig struct {
	// Exact, case-sensitive replacements, matched and written without trimming. Keys
	// must not be empty and are limited to 255 characters; values are limited to 100
	// characters and may be empty.
	Values map[string]string `json:"values,omitzero" api:"required"`
	// Value used when the input does not match any key in `values`.
	Default param.Opt[string] `json:"default,omitzero"`
	paramObj
}

func (r LogsUploaderPolicyReplaceParamsFieldConversionConversionReplaceConfig) MarshalJSON() (data []byte, err error) {
	type shadow LogsUploaderPolicyReplaceParamsFieldConversionConversionReplaceConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogsUploaderPolicyReplaceParamsFieldConversionConversionReplaceConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Format type for logs.
//
// Possible values:
//
//   - **""** - empty, it means it will apply the format configurations from the
//     policy.
//   - **"json"** - output the logs as json lines.
type LogsUploaderPolicyReplaceParamsFormatType string

const (
	LogsUploaderPolicyReplaceParamsFormatTypeJson  LogsUploaderPolicyReplaceParamsFormatType = "json"
	LogsUploaderPolicyReplaceParamsFormatTypeEmpty LogsUploaderPolicyReplaceParamsFormatType = ""
)
