// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
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

// Log viewer provides you with general information about CDN operation. This
// information does not contain all possible sets of fields and restricted by time.
// To receive full data, use Logs Uploader.
//
// LogService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogService] method instead.
type LogService struct {
	Options []option.RequestOption
}

// NewLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogService(opts ...option.RequestOption) (r LogService) {
	r = LogService{}
	r.Options = opts
	return
}

// Get CDN logs for up to 3 days starting today.
//
// You can filter logs using query parameters by client IP, CDN resource, date,
// path and etc.
//
// To filter the CDN logs by 2xx status codes, use:
//
// - &`status__gte`=200&`status__lt`=300
func (r *LogService) List(ctx context.Context, query LogListParams, opts ...option.RequestOption) (res *pagination.OffsetPageCDNLogs[CDNLogEntryData], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cdn/advanced/v1/logs"
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

// Get CDN logs for up to 3 days starting today.
//
// You can filter logs using query parameters by client IP, CDN resource, date,
// path and etc.
//
// To filter the CDN logs by 2xx status codes, use:
//
// - &`status__gte`=200&`status__lt`=300
func (r *LogService) ListAutoPaging(ctx context.Context, query LogListParams, opts ...option.RequestOption) *pagination.OffsetPageCDNLogsAutoPager[CDNLogEntryData] {
	return pagination.NewOffsetPageCDNLogsAutoPager(r.List(ctx, query, opts...))
}

// Download CDN logs for up to 3 days starting today.
//
// You can filter logs using query params by client IP, CDN resource, date, path
// and etc.
func (r *LogService) Download(ctx context.Context, query LogDownloadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/zip")}, opts...)
	path := "cdn/advanced/v1/logs/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CDNLogEntry struct {
	// Contains requested logs.
	Data []CDNLogEntryData `json:"data"`
	// Contains meta-information.
	Meta CDNLogEntryMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNLogEntry) RawJSON() string { return r.JSON.raw }
func (r *CDNLogEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CDNLogEntryData struct {
	// Cache status: HIT, MISS, etc.
	CacheStatus string `json:"cache_status"`
	// IP address from that the request was received.
	ClientIP string `json:"client_ip"`
	// CDN resource custom domain.
	Cname string `json:"cname"`
	// Data center where the request was processed.
	Datacenter string `json:"datacenter"`
	// HTTP method used in the request.
	Method string `json:"method"`
	// Path requested.
	Path string `json:"path"`
	// Value of 'Referer' header.
	Referer string `json:"referer"`
	// CDN resource ID.
	ResourceID int64 `json:"resource_id"`
	// Value of the Content-Type HTTP header, indicating the MIME type of the resource
	// being transmitted.
	SentHTTPContentType string `json:"sent_http_content_type"`
	// Response size in bytes.
	Size int64 `json:"size"`
	// HTTP status code.
	Status int64 `json:"status"`
	// Time required to transmit a complete TCP segment: from the first bit to the
	// last.
	TcpinfoRtt int64 `json:"tcpinfo_rtt"`
	// Log timestamp.
	Timestamp int64 `json:"timestamp"`
	// Value of 'User-Agent' header.
	UserAgent string `json:"user_agent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheStatus         respjson.Field
		ClientIP            respjson.Field
		Cname               respjson.Field
		Datacenter          respjson.Field
		Method              respjson.Field
		Path                respjson.Field
		Referer             respjson.Field
		ResourceID          respjson.Field
		SentHTTPContentType respjson.Field
		Size                respjson.Field
		Status              respjson.Field
		TcpinfoRtt          respjson.Field
		Timestamp           respjson.Field
		UserAgent           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNLogEntryData) RawJSON() string { return r.JSON.raw }
func (r *CDNLogEntryData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contains meta-information.
type CDNLogEntryMeta struct {
	// Total number of records which match given parameters.
	Count int64 `json:"count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CDNLogEntryMeta) RawJSON() string { return r.JSON.raw }
func (r *CDNLogEntryMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogListParams struct {
	// Start date and time of the requested time period (ISO 8601/RFC 3339 format,
	// UTC.)
	//
	// Difference between "from" and "to" cannot exceed 6 hours.
	//
	// Examples:
	//
	// - &from=2021-06-14T00:00:00Z
	// - &from=2021-06-14T00:00:00.000Z
	From string `query:"from" api:"required" json:"-"`
	// End date and time of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	//
	// Difference between "from" and "to" cannot exceed 6 hours.
	//
	// Examples:
	//
	// - &to=2021-06-15T00:00:00Z
	// - &to=2021-06-15T00:00:00.000Z
	To string `query:"to" api:"required" json:"-"`
	// Caching status. Possible values: 'MISS', 'BYPASS', 'EXPIRED', 'STALE',
	// 'PENDING', 'UPDATING', 'REVALIDATED', 'HIT', '-'.
	CacheStatusEq param.Opt[string] `query:"cache_status__eq,omitzero" json:"-"`
	// List of caching statuses. Possible values: 'MISS', 'BYPASS', 'EXPIRED', 'STALE',
	// 'PENDING', 'UPDATING', 'REVALIDATED', 'HIT', '-'. Values should be separated by
	// a comma.
	CacheStatusIn param.Opt[string] `query:"cache_status__in,omitzero" json:"-"`
	// Caching status not equal to the specified value. Possible values: 'MISS',
	// 'BYPASS', 'EXPIRED', 'STALE', 'PENDING', 'UPDATING', 'REVALIDATED', 'HIT', '-'.
	CacheStatusNe param.Opt[string] `query:"cache_status__ne,omitzero" json:"-"`
	// List of caching statuses not equal to the specified values. Possible values:
	// 'MISS', 'BYPASS', 'EXPIRED', 'STALE', 'PENDING', 'UPDATING', 'REVALIDATED',
	// 'HIT', '-'. Values should be separated by a comma.
	CacheStatusNotIn param.Opt[string] `query:"cache_status__not_in,omitzero" json:"-"`
	// IP address of the client who sent the request.
	ClientIPEq param.Opt[string] `query:"client_ip__eq,omitzero" json:"-"`
	// List of IP addresses of the clients who sent the request.
	ClientIPIn param.Opt[string] `query:"client_ip__in,omitzero" json:"-"`
	// IP address of the client who did not send the request.
	ClientIPNe param.Opt[string] `query:"client_ip__ne,omitzero" json:"-"`
	// List of IP addresses of the clients who did not send the request.
	ClientIPNotIn param.Opt[string] `query:"client_ip__not_in,omitzero" json:"-"`
	// Part of the custom domain of the requested CDN resource. Minimum length is 3
	// characters.
	CnameContains param.Opt[string] `query:"cname__contains,omitzero" json:"-"`
	// Custom domain of the requested CDN resource.
	CnameEq param.Opt[string] `query:"cname__eq,omitzero" json:"-"`
	// List of custom domains of the requested CDN resource. Values should be separated
	// by a comma.
	CnameIn param.Opt[string] `query:"cname__in,omitzero" json:"-"`
	// Custom domain of the requested CDN resource not equal to the specified value.
	CnameNe param.Opt[string] `query:"cname__ne,omitzero" json:"-"`
	// List of custom domains of the requested CDN resource not equal to the specified
	// values. Values should be separated by a comma.
	CnameNotIn param.Opt[string] `query:"cname__not_in,omitzero" json:"-"`
	// Data center where request was processed.
	DatacenterEq param.Opt[string] `query:"datacenter__eq,omitzero" json:"-"`
	// List of data centers where request was processed. Values should be separated by
	// a comma.
	DatacenterIn param.Opt[string] `query:"datacenter__in,omitzero" json:"-"`
	// Data center where request was not processed.
	DatacenterNe param.Opt[string] `query:"datacenter__ne,omitzero" json:"-"`
	// List of data centers where request was not processed. Values should be separated
	// by a comma.
	DatacenterNotIn param.Opt[string] `query:"datacenter__not_in,omitzero" json:"-"`
	// A comma-separated list of returned fields.
	//
	// Supported fields are presented in the responses section.
	//
	// Example:
	//
	// - &fields=timestamp,path,status
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Maximum number of log records in the response.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Request HTTP method. Possible values: 'CONNECT', 'DELETE', 'GET', 'HEAD',
	// 'OPTIONS', 'PATCH', 'POST', 'PUT', 'TRACE'.
	MethodEq param.Opt[string] `query:"method__eq,omitzero" json:"-"`
	// Request HTTP method. Possible values: 'CONNECT', 'DELETE', 'GET', 'HEAD',
	// 'OPTIONS', 'PATCH', 'POST', 'PUT', 'TRACE'. Values should be separated by a
	// comma.
	MethodIn param.Opt[string] `query:"method__in,omitzero" json:"-"`
	// Request HTTP method. Possible values: 'CONNECT', 'DELETE', 'GET', 'HEAD',
	// 'OPTIONS', 'PATCH', 'POST', 'PUT', 'TRACE'.
	MethodNe param.Opt[string] `query:"method__ne,omitzero" json:"-"`
	// Request HTTP method. Possible values: 'CONNECT', 'DELETE', 'GET', 'HEAD',
	// 'OPTIONS', 'PATCH', 'POST', 'PUT', 'TRACE'. Values should be separated by a
	// comma.
	MethodNotIn param.Opt[string] `query:"method__not_in,omitzero" json:"-"`
	// Number of log records to skip starting from the beginning of the requested
	// period.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Sorting rules.
	//
	// Possible values:
	//
	// - **method** - Request HTTP method.
	// - **`client_ip`** - IP address of the client who sent the request.
	// - **status** - Status code in the response.
	// - **size** - Response size in bytes.
	// - **cname** - Custom domain of the requested resource.
	// - **`resource_id`** - ID of the requested CDN resource.
	// - **`cache_status`** - Caching status.
	// - **datacenter** - Data center where request was processed.
	// - **timestamp** - Date and time when the request was made.
	//
	// Parameter may have multiple values separated by a comma.
	//
	// By default, ascending sorting is applied. To sort in descending order, add '-'
	// prefix.
	//
	// Example:
	//
	// - &ordering=-timestamp,status
	Ordering param.Opt[string] `query:"ordering,omitzero" json:"-"`
	// ID of the requested CDN resource equal to the specified value.
	ResourceIDEq param.Opt[int64] `query:"resource_id__eq,omitzero" json:"-"`
	// ID of the requested CDN resource greater than the specified value.
	ResourceIDGt param.Opt[int64] `query:"resource_id__gt,omitzero" json:"-"`
	// ID of the requested CDN resource greater than or equal to the specified value.
	ResourceIDGte param.Opt[int64] `query:"resource_id__gte,omitzero" json:"-"`
	// List of IDs of the requested CDN resource. Values should be separated by a
	// comma.
	ResourceIDIn param.Opt[string] `query:"resource_id__in,omitzero" json:"-"`
	// ID of the requested CDN resource less than the specified value.
	ResourceIDLt param.Opt[int64] `query:"resource_id__lt,omitzero" json:"-"`
	// ID of the requested CDN resource less than or equal to the specified value.
	ResourceIDLte param.Opt[int64] `query:"resource_id__lte,omitzero" json:"-"`
	// ID of the requested CDN resource not equal to the specified value.
	ResourceIDNe param.Opt[int64] `query:"resource_id__ne,omitzero" json:"-"`
	// List of IDs of the requested CDN resource not equal to the specified values.
	// Values should be separated by a comma.
	ResourceIDNotIn param.Opt[string] `query:"resource_id__not_in,omitzero" json:"-"`
	// Response size in bytes equal to the specified value.
	SizeEq param.Opt[int64] `query:"size__eq,omitzero" json:"-"`
	// Response size in bytes greater than the specified value.
	SizeGt param.Opt[int64] `query:"size__gt,omitzero" json:"-"`
	// Response size in bytes greater than or equal to the specified value.
	SizeGte param.Opt[int64] `query:"size__gte,omitzero" json:"-"`
	// List of response sizes in bytes. Values should be separated by a comma.
	SizeIn param.Opt[string] `query:"size__in,omitzero" json:"-"`
	// Response size in bytes less than the specified value.
	SizeLt param.Opt[int64] `query:"size__lt,omitzero" json:"-"`
	// Response size in bytes less than or equal to the specified value.
	SizeLte param.Opt[int64] `query:"size__lte,omitzero" json:"-"`
	// Response size in bytes not equal to the specified value.
	SizeNe param.Opt[int64] `query:"size__ne,omitzero" json:"-"`
	// List of response sizes in bytes not equal to the specified values. Values should
	// be separated by
	SizeNotIn param.Opt[string] `query:"size__not_in,omitzero" json:"-"`
	// Status code in the response equal to the specified value.
	StatusEq param.Opt[int64] `query:"status__eq,omitzero" json:"-"`
	// Status code in the response greater than the specified value.
	StatusGt param.Opt[int64] `query:"status__gt,omitzero" json:"-"`
	// Status code in the response greater than or equal to the specified value.
	StatusGte param.Opt[int64] `query:"status__gte,omitzero" json:"-"`
	// List of status codes in the response. Values should be separated by a comma.
	StatusIn param.Opt[string] `query:"status__in,omitzero" json:"-"`
	// Status code in the response less than the specified value.
	StatusLt param.Opt[int64] `query:"status__lt,omitzero" json:"-"`
	// Status code in the response less than or equal to the specified value.
	StatusLte param.Opt[int64] `query:"status__lte,omitzero" json:"-"`
	// Status code in the response not equal to the specified value.
	StatusNe param.Opt[int64] `query:"status__ne,omitzero" json:"-"`
	// List of status codes not in the response. Values should be separated by a comma.
	StatusNotIn param.Opt[string] `query:"status__not_in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LogListParams]'s query parameters as `url.Values`.
func (r LogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogDownloadParams struct {
	// Output format.
	//
	// Possible values:
	//
	// - csv
	// - tsv
	Format string `query:"format" api:"required" json:"-"`
	// Start date and time of the requested time period (ISO 8601/RFC 3339 format,
	// UTC.)
	//
	// Difference between "from" and "to" cannot exceed 6 hours.
	//
	// Examples:
	//
	// - &from=2021-06-14T00:00:00Z
	// - &from=2021-06-14T00:00:00.000Z
	From string `query:"from" api:"required" json:"-"`
	// End date and time of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	//
	// Difference between "from" and "to" cannot exceed 6 hours.
	//
	// Examples:
	//
	// - &to=2021-06-15T00:00:00Z
	// - &to=2021-06-15T00:00:00.000Z
	To string `query:"to" api:"required" json:"-"`
	// Caching status. Possible values: 'MISS', 'BYPASS', 'EXPIRED', 'STALE',
	// 'PENDING', 'UPDATING', 'REVALIDATED', 'HIT', '-'.
	CacheStatusEq param.Opt[string] `query:"cache_status__eq,omitzero" json:"-"`
	// List of caching statuses. Possible values: 'MISS', 'BYPASS', 'EXPIRED', 'STALE',
	// 'PENDING', 'UPDATING', 'REVALIDATED', 'HIT', '-'. Values should be separated by
	// a comma.
	CacheStatusIn param.Opt[string] `query:"cache_status__in,omitzero" json:"-"`
	// Caching status not equal to the specified value. Possible values: 'MISS',
	// 'BYPASS', 'EXPIRED', 'STALE', 'PENDING', 'UPDATING', 'REVALIDATED', 'HIT', '-'.
	CacheStatusNe param.Opt[string] `query:"cache_status__ne,omitzero" json:"-"`
	// List of caching statuses not equal to the specified values. Possible values:
	// 'MISS', 'BYPASS', 'EXPIRED', 'STALE', 'PENDING', 'UPDATING', 'REVALIDATED',
	// 'HIT', '-'. Values should be separated by a comma.
	CacheStatusNotIn param.Opt[string] `query:"cache_status__not_in,omitzero" json:"-"`
	// IP address of the client who sent the request.
	ClientIPEq param.Opt[string] `query:"client_ip__eq,omitzero" json:"-"`
	// List of IP addresses of the clients who sent the request.
	ClientIPIn param.Opt[string] `query:"client_ip__in,omitzero" json:"-"`
	// IP address of the client who did not send the request.
	ClientIPNe param.Opt[string] `query:"client_ip__ne,omitzero" json:"-"`
	// List of IP addresses of the clients who did not send the request.
	ClientIPNotIn param.Opt[string] `query:"client_ip__not_in,omitzero" json:"-"`
	// Part of the custom domain of the requested CDN resource. Minimum length is 3
	// characters.
	CnameContains param.Opt[string] `query:"cname__contains,omitzero" json:"-"`
	// Custom domain of the requested CDN resource.
	CnameEq param.Opt[string] `query:"cname__eq,omitzero" json:"-"`
	// List of custom domains of the requested CDN resource. Values should be separated
	// by a comma.
	CnameIn param.Opt[string] `query:"cname__in,omitzero" json:"-"`
	// Custom domain of the requested CDN resource not equal to the specified value.
	CnameNe param.Opt[string] `query:"cname__ne,omitzero" json:"-"`
	// List of custom domains of the requested CDN resource not equal to the specified
	// values. Values should be separated by a comma.
	CnameNotIn param.Opt[string] `query:"cname__not_in,omitzero" json:"-"`
	// Data center where request was processed.
	DatacenterEq param.Opt[string] `query:"datacenter__eq,omitzero" json:"-"`
	// List of data centers where request was processed. Values should be separated by
	// a comma.
	DatacenterIn param.Opt[string] `query:"datacenter__in,omitzero" json:"-"`
	// Data center where request was not processed.
	DatacenterNe param.Opt[string] `query:"datacenter__ne,omitzero" json:"-"`
	// List of data centers where request was not processed. Values should be separated
	// by a comma.
	DatacenterNotIn param.Opt[string] `query:"datacenter__not_in,omitzero" json:"-"`
	// A comma-separated list of returned fields.
	//
	// Supported fields are presented in the responses section.
	//
	// Example:
	//
	// - &fields=timestamp,path,status
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Maximum number of log records in the response.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Request HTTP method. Possible values: 'CONNECT', 'DELETE', 'GET', 'HEAD',
	// 'OPTIONS', 'PATCH', 'POST', 'PUT', 'TRACE'.
	MethodEq param.Opt[string] `query:"method__eq,omitzero" json:"-"`
	// Request HTTP method. Possible values: 'CONNECT', 'DELETE', 'GET', 'HEAD',
	// 'OPTIONS', 'PATCH', 'POST', 'PUT', 'TRACE'. Values should be separated by a
	// comma.
	MethodIn param.Opt[string] `query:"method__in,omitzero" json:"-"`
	// Request HTTP method. Possible values: 'CONNECT', 'DELETE', 'GET', 'HEAD',
	// 'OPTIONS', 'PATCH', 'POST', 'PUT', 'TRACE'.
	MethodNe param.Opt[string] `query:"method__ne,omitzero" json:"-"`
	// Request HTTP method. Possible values: 'CONNECT', 'DELETE', 'GET', 'HEAD',
	// 'OPTIONS', 'PATCH', 'POST', 'PUT', 'TRACE'. Values should be separated by a
	// comma.
	MethodNotIn param.Opt[string] `query:"method__not_in,omitzero" json:"-"`
	// Number of log records to skip starting from the beginning of the requested
	// period.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// ID of the requested CDN resource equal to the specified value.
	ResourceIDEq param.Opt[int64] `query:"resource_id__eq,omitzero" json:"-"`
	// ID of the requested CDN resource greater than the specified value.
	ResourceIDGt param.Opt[int64] `query:"resource_id__gt,omitzero" json:"-"`
	// ID of the requested CDN resource greater than or equal to the specified value.
	ResourceIDGte param.Opt[int64] `query:"resource_id__gte,omitzero" json:"-"`
	// List of IDs of the requested CDN resource. Values should be separated by a
	// comma.
	ResourceIDIn param.Opt[string] `query:"resource_id__in,omitzero" json:"-"`
	// ID of the requested CDN resource less than the specified value.
	ResourceIDLt param.Opt[int64] `query:"resource_id__lt,omitzero" json:"-"`
	// ID of the requested CDN resource less than or equal to the specified value.
	ResourceIDLte param.Opt[int64] `query:"resource_id__lte,omitzero" json:"-"`
	// ID of the requested CDN resource not equal to the specified value.
	ResourceIDNe param.Opt[int64] `query:"resource_id__ne,omitzero" json:"-"`
	// List of IDs of the requested CDN resource not equal to the specified values.
	// Values should be separated by a comma.
	ResourceIDNotIn param.Opt[string] `query:"resource_id__not_in,omitzero" json:"-"`
	// Response size in bytes equal to the specified value.
	SizeEq param.Opt[int64] `query:"size__eq,omitzero" json:"-"`
	// Response size in bytes greater than the specified value.
	SizeGt param.Opt[int64] `query:"size__gt,omitzero" json:"-"`
	// Response size in bytes greater than or equal to the specified value.
	SizeGte param.Opt[int64] `query:"size__gte,omitzero" json:"-"`
	// List of response sizes in bytes. Values should be separated by a comma.
	SizeIn param.Opt[string] `query:"size__in,omitzero" json:"-"`
	// Response size in bytes less than the specified value.
	SizeLt param.Opt[int64] `query:"size__lt,omitzero" json:"-"`
	// Response size in bytes less than or equal to the specified value.
	SizeLte param.Opt[int64] `query:"size__lte,omitzero" json:"-"`
	// Response size in bytes not equal to the specified value.
	SizeNe param.Opt[int64] `query:"size__ne,omitzero" json:"-"`
	// List of response sizes in bytes not equal to the specified values. Values should
	// be separated by
	SizeNotIn param.Opt[string] `query:"size__not_in,omitzero" json:"-"`
	// Sorting rules.
	//
	// Possible values:
	//
	// - **method** - Request HTTP method.
	// - **`client_ip`** - IP address of the client who sent the request.
	// - **status** - Status code in the response.
	// - **size** - Response size in bytes.
	// - **cname** - Custom domain of the requested resource.
	// - **`resource_id`** - ID of the requested CDN resource.
	// - **`cache_status`** - Caching status.
	// - **datacenter** - Data center where request was processed.
	// - **timestamp** - Date and time when the request was made.
	//
	// May include multiple values separated by a comma.
	//
	// Example:
	//
	// - &sort=-timestamp,status
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	// Status code in the response equal to the specified value.
	StatusEq param.Opt[int64] `query:"status__eq,omitzero" json:"-"`
	// Status code in the response greater than the specified value.
	StatusGt param.Opt[int64] `query:"status__gt,omitzero" json:"-"`
	// Status code in the response greater than or equal to the specified value.
	StatusGte param.Opt[int64] `query:"status__gte,omitzero" json:"-"`
	// List of status codes in the response. Values should be separated by a comma.
	StatusIn param.Opt[string] `query:"status__in,omitzero" json:"-"`
	// Status code in the response less than the specified value.
	StatusLt param.Opt[int64] `query:"status__lt,omitzero" json:"-"`
	// Status code in the response less than or equal to the specified value.
	StatusLte param.Opt[int64] `query:"status__lte,omitzero" json:"-"`
	// Status code in the response not equal to the specified value.
	StatusNe param.Opt[int64] `query:"status__ne,omitzero" json:"-"`
	// List of status codes not in the response. Values should be separated by a comma.
	StatusNotIn param.Opt[string] `query:"status__not_in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LogDownloadParams]'s query parameters as `url.Values`.
func (r LogDownloadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
