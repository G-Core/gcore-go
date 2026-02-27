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
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// Consumption statistics is updated in near real-time as a standard practice.
// However, the frequency of updates can vary, but they are typically available
// within a 24-hour period. Exceptions, such as maintenance periods, may delay data
// beyond 24 hours until servers resume and fill in the missing statistics.
//
// StatisticService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatisticService] method instead.
type StatisticService struct {
	Options []option.RequestOption
}

// NewStatisticService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStatisticService(opts ...option.RequestOption) (r StatisticService) {
	r = StatisticService{}
	r.Options = opts
	return
}

// Get the number of CDN resources that used Logs uploader.
//
// Request URL parameters should be added as a query string after the endpoint.
func (r *StatisticService) GetLogsUsageAggregated(ctx context.Context, query StatisticGetLogsUsageAggregatedParams, opts ...option.RequestOption) (res *LogsAggregatedStats, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/statistics/raw_logs_usage/aggregated"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Logs uploader usage statistics for up to 90 days starting today.
//
// Request URL parameters should be added as a query string after the endpoint.
func (r *StatisticService) GetLogsUsageSeries(ctx context.Context, query StatisticGetLogsUsageSeriesParams, opts ...option.RequestOption) (res *UsageSeriesStats, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/statistics/raw_logs_usage/series"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get aggregated CDN resources statistics.
//
// Request URL parameters should be added as a query string after the endpoint.
//
// Aggregated data does not include data for the last two hours.
func (r *StatisticService) GetResourceUsageAggregated(ctx context.Context, query StatisticGetResourceUsageAggregatedParams, opts ...option.RequestOption) (res *ResourceAggregatedStats, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/statistics/aggregate/stats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get CDN resources statistics for up to 365 days starting today.
func (r *StatisticService) GetResourceUsageSeries(ctx context.Context, query StatisticGetResourceUsageSeriesParams, opts ...option.RequestOption) (res *ResourceUsageStats, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/statistics/series"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The number of CDN resources that use origin shielding.
//
// Request URL parameters should be added as a query string after the endpoint.
func (r *StatisticService) GetShieldUsageAggregated(ctx context.Context, query StatisticGetShieldUsageAggregatedParams, opts ...option.RequestOption) (res *ShieldAggregatedStats, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/statistics/shield_usage/aggregated"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get origin shielding usage statistics for up to 365 days starting from today.
//
// Request URL parameters should be added as a query string after the endpoint.
func (r *StatisticService) GetShieldUsageSeries(ctx context.Context, query StatisticGetShieldUsageSeriesParams, opts ...option.RequestOption) (res *UsageSeriesStats, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/statistics/shield_usage/series"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type LogsAggregatedStats struct {
	// CDN resource ID for which statistics data is shown.
	Number1Example any `json:"1 (example)"`
	// Statistics parameters.
	Metrics any `json:"metrics"`
	// Number of resources that used Logs uploader.
	RawLogsUsage string `json:"raw_logs_usage"`
	// Resources IDs by which statistics data is grouped..
	Resource any `json:"resource"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number1Example respjson.Field
		Metrics        respjson.Field
		RawLogsUsage   respjson.Field
		Resource       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsAggregatedStats) RawJSON() string { return r.JSON.raw }
func (r *LogsAggregatedStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResourceAggregatedStats struct {
	// CDN resource ID for which statistics data is shown.
	Number1Example any `json:"1 (example)"`
	// 95 percentile bandwidth value
	Number95Percentile int64 `json:"95_percentile"`
	// Traffic in bytes from Backblaze origin.
	BackblazeBytes int64 `json:"backblaze_bytes"`
	// Formula: 1 - `upstream_bytes` / `sent_bytes`. We deduct the non-cached traffic
	// from the total traffic amount
	CacheHitTrafficRatio int64 `json:"cache_hit_traffic_ratio"`
	// Region by which statistics data is grouped.
	CisExample any `json:"cis (example)"`
	// Maximum bandwidth
	MaxBandwidth int64 `json:"max_bandwidth"`
	// Statistics parameters.
	Metrics any `json:"metrics"`
	// Minimum bandwidth
	MinBandwidth int64 `json:"min_bandwidth"`
	// Regions by which statistics data is grouped.
	Region any `json:"region"`
	// Number of requests to edge servers.
	Requests int64 `json:"requests"`
	// Resources IDs by which statistics data is grouped.
	Resource any `json:"resource"`
	// Statistics by content type. It returns a number of responses for content with
	// different MIME types.
	ResponseTypes any `json:"response_types"`
	// Number of 2xx response codes.
	Responses2xx int64 `json:"responses_2xx"`
	// Number of 3xx response codes.
	Responses3xx int64 `json:"responses_3xx"`
	// Number of 4xx response codes.
	Responses4xx int64 `json:"responses_4xx"`
	// Number of 5xx response codes.
	Responses5xx int64 `json:"responses_5xx"`
	// Number of responses with the header Cache: HIT.
	ResponsesHit int64 `json:"responses_hit"`
	// Number of responses with the header Cache: MISS.
	ResponsesMiss int64 `json:"responses_miss"`
	// Traffic in bytes from CDN servers to clients.
	SentBytes int64 `json:"sent_bytes"`
	// Upstream bytes and `sent_bytes` combined.
	TotalBytes int64 `json:"total_bytes"`
	// Traffic in bytes from the upstream to CDN servers.
	UpstreamBytes int64 `json:"upstream_bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number1Example       respjson.Field
		Number95Percentile   respjson.Field
		BackblazeBytes       respjson.Field
		CacheHitTrafficRatio respjson.Field
		CisExample           respjson.Field
		MaxBandwidth         respjson.Field
		Metrics              respjson.Field
		MinBandwidth         respjson.Field
		Region               respjson.Field
		Requests             respjson.Field
		Resource             respjson.Field
		ResponseTypes        respjson.Field
		Responses2xx         respjson.Field
		Responses3xx         respjson.Field
		Responses4xx         respjson.Field
		Responses5xx         respjson.Field
		ResponsesHit         respjson.Field
		ResponsesMiss        respjson.Field
		SentBytes            respjson.Field
		TotalBytes           respjson.Field
		UpstreamBytes        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResourceAggregatedStats) RawJSON() string { return r.JSON.raw }
func (r *ResourceAggregatedStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResourceUsageStats struct {
	// ID of CDN resource for which statistics data is shown.
	Number1Example any `json:"1 (example)"`
	// BackBlaze bytes from Backblaze origin.
	//
	// Represented by two values:
	//
	// - 1543622400 — Time in the UNIX timestamp when statistics were received.
	// - 17329220573 — Bytes.
	BackblazeBytes []int64 `json:"backblaze_bytes"`
	// Types of statistics data.
	//
	// Possible values:
	//
	//   - **`upstream_bytes`** – Traffic in bytes from an origin server to CDN servers
	//     or to origin shielding when used.
	//   - **`sent_bytes`** – Traffic in bytes from CDN servers to clients.
	//   - **`shield_bytes`** – Traffic in bytes from origin shielding to CDN servers.
	//   - **`backblaze_bytes`** - Traffic in bytes from Backblaze origin.
	//   - **`total_bytes`** – `shield_bytes`, `upstream_bytes` and `sent_bytes`
	//     combined.
	//   - **`cdn_bytes`** – `sent_bytes` and `shield_bytes` combined.
	//   - **requests** – Number of requests to edge servers.
	//   - **`responses_2xx`** – Number of 2xx response codes.
	//   - **`responses_3xx`** – Number of 3xx response codes.
	//   - **`responses_4xx`** – Number of 4xx response codes.
	//   - **`responses_5xx`** – Number of 5xx response codes.
	//   - **`responses_hit`** – Number of responses with the header Cache: HIT.
	//   - **`responses_miss`** – Number of responses with the header Cache: MISS.
	//   - **`response_types`** – Statistics by content type. It returns a number of
	//     responses for content with different MIME types.
	//   - **`cache_hit_traffic_ratio`** – Formula: 1 - `upstream_bytes` / `sent_bytes`.
	//     We deduct the non-cached traffic from the total traffic amount.
	//   - **`cache_hit_requests_ratio`** – Formula: `responses_hit` / requests. The
	//     share of sending cached content.
	//   - **`shield_traffic_ratio`** – Formula: (`shield_bytes` - `upstream_bytes`) /
	//     `shield_bytes`. The efficiency of the Origin Shielding: how much more traffic
	//     is sent from the Origin Shielding than from the origin.
	//   - **`image_processed`** - Number of images transformed on the Image optimization
	//     service.
	//   - **`request_time`** - Time elapsed between the first bytes of a request were
	//     processed and logging after the last bytes were sent to a user.
	//   - **`upstream_response_time`** - Number of milliseconds it took to receive a
	//     response from an origin. If upstream `response_time_` contains several
	//     indications for one request (in case of more than 1 origin), we summarize
	//     them. In case of aggregating several queries, the average of this amount is
	//     calculated.
	//
	// Metrics **`upstream_response_time`** and **`request_time`** should be requested
	// separately from other metrics
	Metrics any `json:"metrics"`
	// Regions for which data is displayed.
	//
	// Possible values:
	//
	// - **na** – North America
	// - **eu** – Europe
	// - **cis** – Commonwealth of Independent States
	// - **asia** – Asia
	// - **au** – Australia
	// - **latam** – Latin America
	// - **me** – Middle East
	// - **africa** - Africa
	// - **sa** - South America
	Region any `json:"region"`
	// Resources IDs by which statistics data is grouped.
	Resource any `json:"resource"`
	// Bytes from CDN servers to the end-users.
	//
	// Represented by two values:
	//
	// - 1543622400 — Time in the UNIX timestamp when statistics were received.
	// - 17329220573 — Bytes.
	SentBytes []int64 `json:"sent_bytes"`
	// Upstream bytes and `sent_bytes` combined.
	//
	// Represented by two values:
	//
	// - 1543622400 — Time in the UNIX timestamp when statistics were received.
	// - 17329220573 — Bytes.
	TotalBytes []int64 `json:"total_bytes"`
	// Bytes from the upstream to the CDN servers.
	//
	// Represented by two values:
	//
	// - 1543622400 — Time in the UNIX timestamp when statistics were received.
	// - 17329220573 — Bytes.
	UpstreamBytes []int64 `json:"upstream_bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number1Example respjson.Field
		BackblazeBytes respjson.Field
		Metrics        respjson.Field
		Region         respjson.Field
		Resource       respjson.Field
		SentBytes      respjson.Field
		TotalBytes     respjson.Field
		UpstreamBytes  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResourceUsageStats) RawJSON() string { return r.JSON.raw }
func (r *ResourceUsageStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShieldAggregatedStats struct {
	// CDN resource ID for which statistics data is shown.
	Number1Example any `json:"1 (example)"`
	// Statistics parameters.
	Metrics any `json:"metrics"`
	// Resources IDs by which statistics data is grouped.
	Resource any `json:"resource"`
	// Number of CDN resources that used origin shielding.
	ShieldUsage string `json:"shield_usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number1Example respjson.Field
		Metrics        respjson.Field
		Resource       respjson.Field
		ShieldUsage    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShieldAggregatedStats) RawJSON() string { return r.JSON.raw }
func (r *ShieldAggregatedStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageSeriesStats []UsageSeriesStat

type UsageSeriesStat struct {
	// Date and time when paid feature was enabled (ISO 8601/RFC 3339 format, UTC.)
	ActiveFrom string `json:"active_from"`
	// Date and time when paid feature was disabled (ISO 8601/RFC 3339 format, UTC.)
	//
	// It returns **null** if the paid feature is enabled.
	ActiveTo string `json:"active_to" api:"nullable"`
	// Client ID.
	ClientID int64 `json:"client_id"`
	// CDN resource CNAME.
	Cname string `json:"cname"`
	// CDN resource ID.
	ResourceID int64 `json:"resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveFrom  respjson.Field
		ActiveTo    respjson.Field
		ClientID    respjson.Field
		Cname       respjson.Field
		ResourceID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageSeriesStat) RawJSON() string { return r.JSON.raw }
func (r *UsageSeriesStat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatisticGetLogsUsageAggregatedParams struct {
	// Beginning of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	From string `query:"from" api:"required" json:"-"`
	// End of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	To string `query:"to" api:"required" json:"-"`
	// The way the parameters are arranged in the response.
	//
	// Possible values:
	//
	// - **true** – Flat structure is used.
	// - **false** – Embedded structure is used (default.)
	Flat param.Opt[bool] `query:"flat,omitzero" json:"-"`
	// Output data grouping.
	//
	// Possible value:
	//
	// - **resource** - Data is grouped by CDN resources.
	GroupBy param.Opt[string] `query:"group_by,omitzero" json:"-"`
	// CDN resources IDs by that statistics data is grouped.
	//
	// To request multiple values, use:
	//
	// - &resource=1&resource=2
	//
	// If CDN resource ID is not specified, data related to all CDN resources is
	// returned.
	Resource param.Opt[int64] `query:"resource,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetLogsUsageAggregatedParams]'s query parameters
// as `url.Values`.
func (r StatisticGetLogsUsageAggregatedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetLogsUsageSeriesParams struct {
	// Beginning of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	From string `query:"from" api:"required" json:"-"`
	// End of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	To string `query:"to" api:"required" json:"-"`
	// CDN resources IDs by that statistics data is grouped.
	//
	// To request multiple values, use:
	//
	// - &resource=1&resource=2
	//
	// If CDN resource ID is not specified, data related to all CDN resources is
	// returned.
	Resource param.Opt[int64] `query:"resource,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetLogsUsageSeriesParams]'s query parameters as
// `url.Values`.
func (r StatisticGetLogsUsageSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetResourceUsageAggregatedParams struct {
	// Beginning of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	From string `query:"from" api:"required" json:"-"`
	// Types of statistics data.
	//
	// Possible values:
	//
	//   - **`upstream_bytes`** – Traffic in bytes from an origin server to CDN servers
	//     or to origin shielding when used.
	//   - **`sent_bytes`** – Traffic in bytes from CDN servers to clients.
	//   - **`shield_bytes`** – Traffic in bytes from origin shielding to CDN servers.
	//   - **`backblaze_bytes`** - Traffic in bytes from Backblaze origin.
	//   - **`total_bytes`** – `shield_bytes`, `upstream_bytes` and `sent_bytes`
	//     combined.
	//   - **`cdn_bytes`** – `sent_bytes` and `shield_bytes` combined.
	//   - **requests** – Number of requests to edge servers.
	//   - **`responses_2xx`** – Number of 2xx response codes.
	//   - **`responses_3xx`** – Number of 3xx response codes.
	//   - **`responses_4xx`** – Number of 4xx response codes.
	//   - **`responses_5xx`** – Number of 5xx response codes.
	//   - **`responses_hit`** – Number of responses with the header Cache: HIT.
	//   - **`responses_miss`** – Number of responses with the header Cache: MISS.
	//   - **`response_types`** – Statistics by content type. It returns a number of
	//     responses for content with different MIME types.
	//   - **`cache_hit_traffic_ratio`** – Formula: 1 - `upstream_bytes` / `sent_bytes`.
	//     We deduct the non-cached traffic from the total traffic amount.
	//   - **`cache_hit_requests_ratio`** – Formula: `responses_hit` / requests. The
	//     share of sending cached content.
	//   - **`shield_traffic_ratio`** – Formula: (`shield_bytes` - `upstream_bytes`) /
	//     `shield_bytes`. The efficiency of the Origin Shielding: how much more traffic
	//     is sent from the Origin Shielding than from the origin.
	//   - **`image_processed`** - Number of images transformed on the Image optimization
	//     service.
	//   - **`request_time`** - Time elapsed between the first bytes of a request were
	//     processed and logging after the last bytes were sent to a user.
	//   - **`upstream_response_time`** - Number of milliseconds it took to receive a
	//     response from an origin. If upstream `response_time_` contains several
	//     indications for one request (in case of more than 1 origin), we summarize
	//     them. In case of aggregating several queries, the average of this amount is
	//     calculated.
	//   - **`95_percentile`** - Represents the 95th percentile of network bandwidth
	//     usage in bytes per second. This means that 95% of the time, the network
	//     resource usage was below this value.
	//   - **`max_bandwidth`** - The maximum network bandwidth that was used during the
	//     selected time represented in bytes per second.
	//   - **`min_bandwidth`** - The minimum network bandwidth that was used during the
	//     selected time represented in bytes per second.
	//
	// Metrics **`upstream_response_time`** and **`request_time`** should be requested
	// separately from other metrics
	Metrics string `query:"metrics" api:"required" json:"-"`
	// Service name.
	//
	// Possible value:
	//
	// - CDN
	Service string `query:"service" api:"required" json:"-"`
	// End of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	To string `query:"to" api:"required" json:"-"`
	// Names of countries for which data should be displayed. English short name from
	// [ISO 3166 standard][1] without the definite article ("the") should be used.
	//
	// To request multiple values, use:
	//
	// - &countries=france&countries=denmark
	//
	// [1]: https://www.iso.org/obp/ui/#search/code/
	Countries param.Opt[string] `query:"countries,omitzero" json:"-"`
	// The way the parameters are arranged in the response.
	//
	// Possible values:
	//
	// - **true** – Flat structure is used.
	// - **false** – Embedded structure is used (default.)
	Flat param.Opt[bool] `query:"flat,omitzero" json:"-"`
	// Output data grouping.
	//
	// Possible values:
	//
	//   - **resource** – Data is grouped by CDN resources IDs.
	//   - **region** – Data is grouped by regions of CDN edge servers.
	//   - **country** – Data is grouped by countries of CDN edge servers.
	//   - **vhost** – Data is grouped by resources CNAMEs.
	//   - **`client_country`** - Data is grouped by countries, based on end-users'
	//     location.
	//
	// To request multiple values, use:
	//
	// - &`group_by`=region&`group_by`=resource
	GroupBy param.Opt[string] `query:"group_by,omitzero" json:"-"`
	// Regions for which data is displayed.
	//
	// Possible values:
	//
	// - **na** – North America
	// - **eu** – Europe
	// - **cis** – Commonwealth of Independent States
	// - **asia** – Asia
	// - **au** – Australia
	// - **latam** – Latin America
	// - **me** – Middle East
	// - **africa** - Africa
	// - **sa** - South America
	Regions param.Opt[string] `query:"regions,omitzero" json:"-"`
	// CDN resources IDs by that statistics data is grouped.
	//
	// To request multiple values, use:
	//
	// - &resource=1&resource=2
	//
	// If CDN resource ID is not specified, data related to all CDN resources is
	// returned.
	Resource param.Opt[int64] `query:"resource,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetResourceUsageAggregatedParams]'s query
// parameters as `url.Values`.
func (r StatisticGetResourceUsageAggregatedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetResourceUsageSeriesParams struct {
	// Beginning of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	From string `query:"from" api:"required" json:"-"`
	// Duration of the time blocks into which the data will be divided.
	//
	// Possible values:
	//
	// - **1m** - available only for up to 1 month in the past.
	// - **5m**
	// - **15m**
	// - **1h**
	// - **1d**
	Granularity string `query:"granularity" api:"required" json:"-"`
	// Types of statistics data.
	//
	// Possible values:
	//
	//   - **`upstream_bytes`** – Traffic in bytes from an origin server to CDN servers
	//     or to origin shielding when used.
	//   - **`sent_bytes`** – Traffic in bytes from CDN servers to clients.
	//   - **`shield_bytes`** – Traffic in bytes from origin shielding to CDN servers.
	//   - **`backblaze_bytes`** - Traffic in bytes from Backblaze origin.
	//   - **`total_bytes`** – `shield_bytes`, `upstream_bytes` and `sent_bytes`
	//     combined.
	//   - **`cdn_bytes`** – `sent_bytes` and `shield_bytes` combined.
	//   - **requests** – Number of requests to edge servers.
	//   - **`responses_2xx`** – Number of 2xx response codes.
	//   - **`responses_3xx`** – Number of 3xx response codes.
	//   - **`responses_4xx`** – Number of 4xx response codes.
	//   - **`responses_5xx`** – Number of 5xx response codes.
	//   - **`responses_hit`** – Number of responses with the header Cache: HIT.
	//   - **`responses_miss`** – Number of responses with the header Cache: MISS.
	//   - **`response_types`** – Statistics by content type. It returns a number of
	//     responses for content with different MIME types.
	//   - **`cache_hit_traffic_ratio`** – Formula: 1 - `upstream_bytes` / `sent_bytes`.
	//     We deduct the non-cached traffic from the total traffic amount.
	//   - **`cache_hit_requests_ratio`** – Formula: `responses_hit` / requests. The
	//     share of sending cached content.
	//   - **`shield_traffic_ratio`** – Formula: (`shield_bytes` - `upstream_bytes`) /
	//     `shield_bytes`. The efficiency of the Origin Shielding: how much more traffic
	//     is sent from the Origin Shielding than from the origin.
	//   - **`image_processed`** - Number of images transformed on the Image optimization
	//     service.
	//   - **`request_time`** - Time elapsed between the first bytes of a request were
	//     processed and logging after the last bytes were sent to a user.
	//   - **`upstream_response_time`** - Number of milliseconds it took to receive a
	//     response from an origin. If upstream `response_time_` contains several
	//     indications for one request (in case of more than 1 origin), we summarize
	//     them. In case of aggregating several queries, the average of this amount is
	//     calculated.
	//
	// Metrics **`upstream_response_time`** and **`request_time`** should be requested
	// separately from other metrics
	Metrics string `query:"metrics" api:"required" json:"-"`
	// Service name.
	//
	// Possible value:
	//
	// - CDN
	Service string `query:"service" api:"required" json:"-"`
	// End of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	To string `query:"to" api:"required" json:"-"`
	// Names of countries for which data should be displayed. English short name from
	// [ISO 3166 standard][1] without the definite article ("the") should be used.
	//
	// To request multiple values, use:
	//
	// - &countries=france&countries=denmark
	//
	// [1]: https://www.iso.org/obp/ui/#search/code/
	Countries param.Opt[string] `query:"countries,omitzero" json:"-"`
	// Output data grouping.
	//
	// Possible values:
	//
	//   - **resource** – Data is grouped by CDN resources IDs.
	//   - **region** – Data is grouped by regions of CDN edge servers.
	//   - **country** – Data is grouped by countries of CDN edge servers.
	//   - **vhost** – Data is grouped by resources CNAMEs.
	//   - **`client_country`** - Data is grouped by countries, based on end-users'
	//     location.
	//
	// To request multiple values, use:
	//
	// - &`group_by`=region&`group_by`=resource
	GroupBy param.Opt[string] `query:"group_by,omitzero" json:"-"`
	// Regions for which data is displayed.
	//
	// Possible values:
	//
	// - **na** – North America
	// - **eu** – Europe
	// - **cis** – Commonwealth of Independent States
	// - **asia** – Asia
	// - **au** – Australia
	// - **latam** – Latin America
	// - **me** – Middle East
	// - **africa** - Africa
	// - **sa** - South America
	Regions param.Opt[string] `query:"regions,omitzero" json:"-"`
	// CDN resources IDs by that statistics data is grouped.
	//
	// To request multiple values, use:
	//
	// - &resource=1&resource=2
	//
	// If CDN resource ID is not specified, data related to all CDN resources is
	// returned.
	Resource param.Opt[int64] `query:"resource,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetResourceUsageSeriesParams]'s query parameters
// as `url.Values`.
func (r StatisticGetResourceUsageSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetShieldUsageAggregatedParams struct {
	// Beginning of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	From string `query:"from" api:"required" json:"-"`
	// End of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	To string `query:"to" api:"required" json:"-"`
	// The way the parameters are arranged in the response.
	//
	// Possible values:
	//
	// - **true** – Flat structure is used.
	// - **false** – Embedded structure is used (default.)
	Flat param.Opt[bool] `query:"flat,omitzero" json:"-"`
	// Output data grouping.
	//
	// Possible value:
	//
	// - **resource** - Data is grouped by CDN resources.
	GroupBy param.Opt[string] `query:"group_by,omitzero" json:"-"`
	// CDN resources IDs by that statistics data is grouped.
	//
	// To request multiple values, use:
	//
	// - &resource=1&resource=2
	//
	// If CDN resource ID is not specified, data related to all CDN resources is
	// returned.
	Resource param.Opt[int64] `query:"resource,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetShieldUsageAggregatedParams]'s query parameters
// as `url.Values`.
func (r StatisticGetShieldUsageAggregatedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetShieldUsageSeriesParams struct {
	// Beginning of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	From string `query:"from" api:"required" json:"-"`
	// End of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	To string `query:"to" api:"required" json:"-"`
	// CDN resources IDs by that statistics data is grouped.
	//
	// To request multiple values, use:
	//
	// - &resource=1&resource=2
	//
	// If CDN resource ID is not specified, data related to all CDN resources is
	// returned.
	Resource param.Opt[int64] `query:"resource,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetShieldUsageSeriesParams]'s query parameters as
// `url.Values`.
func (r StatisticGetShieldUsageSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
