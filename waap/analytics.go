// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"encoding/json"
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

// AnalyticsService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsService] method instead.
type AnalyticsService struct {
	Options []option.RequestOption
}

// NewAnalyticsService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnalyticsService(opts ...option.RequestOption) (r AnalyticsService) {
	r = AnalyticsService{}
	r.Options = opts
	return
}

// Retrieves event statistics per given dimension of a request characteristics. A
// WAAP _Event_ represents a request observed by the system. The report contains
// the total, blocked, suppressed, and allowed event counts for top ten points in
// the selected dimension.
func (r *AnalyticsService) GetEventStatistics(ctx context.Context, dimension AnalyticsGetEventStatisticsParamsDimension, query AnalyticsGetEventStatisticsParams, opts ...option.RequestOption) (res *WaapSimpleEventStatistics, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/analytics/stats/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve request log data over account's domains. The log records every request
// passing through WAAP towards the origin server.
func (r *AnalyticsService) GetRequests(ctx context.Context, query AnalyticsGetRequestsParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapRequestSummary], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "waap/v1/analytics/requests"
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

// Retrieve request log data over account's domains. The log records every request
// passing through WAAP towards the origin server.
func (r *AnalyticsService) GetRequestsAutoPaging(ctx context.Context, query AnalyticsGetRequestsParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapRequestSummary] {
	return pagination.NewOffsetPageAutoPager(r.GetRequests(ctx, query, opts...))
}

// Retrieves a comprehensive report on traffic statistics for a set of domains. The
// report includes details such as API requests, blocked events, error counts, and
// many more traffic-related metrics.
func (r *AnalyticsService) GetTraffic(ctx context.Context, query AnalyticsGetTrafficParams, opts ...option.RequestOption) (res *[]WaapTrafficMetrics, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/analytics/traffic"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieves a traffic time series data over a set of domains. The response is
// suitable for plotting a time series chart. This method allows filtering the
// traffic data by various criteria.
func (r *AnalyticsService) GetTrafficFiltered(ctx context.Context, query AnalyticsGetTrafficFilteredParams, opts ...option.RequestOption) (res *[]WaapCompactTrafficMetrics, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/analytics/traffic-filtered"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Represents the compact traffic metrics for a domain at a given time window
type WaapCompactTrafficMetrics struct {
	// UNIX timestamp indicating when the traffic data was recorded
	Timestamp int64 `json:"timestamp" api:"required"`
	// Traffic passed due to a permissive security rule
	Allowed int64 `json:"allowed"`
	// Traffic blocked by a security policy, custom rule, or DDoS protection
	Blocked int64 `json:"blocked"`
	// Traffic that was identified as malicious by security policies (for monitored
	// domains)
	Monitored int64 `json:"monitored"`
	// Traffic (number of requests) that was passed to the origin and didn't trigger
	// any rules
	Passed int64 `json:"passed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timestamp   respjson.Field
		Allowed     respjson.Field
		Blocked     respjson.Field
		Monitored   respjson.Field
		Passed      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCompactTrafficMetrics) RawJSON() string { return r.JSON.raw }
func (r *WaapCompactTrafficMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of event counts per dimension values over a time span.
type WaapSimpleEventStatistics struct {
	// Number of requests passed due to a permissive security rule.
	Allowed [][]any `json:"allowed" api:"required"`
	// Number of blocked request events, with the same keys and formatted the same as
	// in the total field.
	Blocked [][]any `json:"blocked" api:"required"`
	// Additional information, depending on the selected dimension. Keys refer to the
	// items in the result (first value of a tuple).
	Reference map[string]WaapSimpleEventStatisticsReferenceUnion `json:"reference" api:"required"`
	// Number of requests observed in monitoring mode that would have been blocked
	// otherwise.
	Suppressed [][]any `json:"suppressed" api:"required"`
	// Total number of observed requests. First element of the tuple is a key, the
	// second one is its counter value. The key refers to a point in the requested
	// dimension (e.g., an IP address). Results are ordered by the counter value in
	// descending order.
	Total [][]any `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allowed     respjson.Field
		Blocked     respjson.Field
		Reference   respjson.Field
		Suppressed  respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapSimpleEventStatistics) RawJSON() string { return r.JSON.raw }
func (r *WaapSimpleEventStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapSimpleEventStatisticsReferenceUnion contains all possible properties and
// values from [WaapSimpleEventStatisticsReferenceIPReference],
// [WaapSimpleEventStatisticsReferenceDomainReference].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WaapSimpleEventStatisticsReferenceUnion struct {
	// This field is from variant [WaapSimpleEventStatisticsReferenceIPReference].
	Country string `json:"country"`
	// This field is from variant [WaapSimpleEventStatisticsReferenceIPReference].
	Org string `json:"org"`
	// This field is from variant [WaapSimpleEventStatisticsReferenceDomainReference].
	DomainID int64 `json:"domain_id"`
	JSON     struct {
		Country  respjson.Field
		Org      respjson.Field
		DomainID respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapSimpleEventStatisticsReferenceUnion) AsIPReference() (v WaapSimpleEventStatisticsReferenceIPReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapSimpleEventStatisticsReferenceUnion) AsDomainReference() (v WaapSimpleEventStatisticsReferenceDomainReference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapSimpleEventStatisticsReferenceUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapSimpleEventStatisticsReferenceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapSimpleEventStatisticsReferenceIPReference struct {
	Country string `json:"country" api:"required"`
	Org     string `json:"org" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		Org         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapSimpleEventStatisticsReferenceIPReference) RawJSON() string { return r.JSON.raw }
func (r *WaapSimpleEventStatisticsReferenceIPReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapSimpleEventStatisticsReferenceDomainReference struct {
	DomainID int64 `json:"domain_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DomainID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapSimpleEventStatisticsReferenceDomainReference) RawJSON() string { return r.JSON.raw }
func (r *WaapSimpleEventStatisticsReferenceDomainReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsGetEventStatisticsParams struct {
	// Filter data items starting from a specified date in ISO 8601 format
	Start string `query:"start" api:"required" json:"-"`
	// Filter data items up to a specified end date in ISO 8601 format. If not
	// provided, defaults to the current date and time.
	End param.Opt[string] `query:"end,omitzero" json:"-"`
	// List of domain IDs. Empty list means all domains belonging to the current
	// account.
	Domains []int64 `query:"domains,omitzero" json:"-"`
	// Filter statistics by client IP addresses (max 10).
	IPs []string `query:"ips,omitzero" format:"ipvanyaddress" json:"-"`
	// Filter data by name of a security rule matched the request.
	SecurityRuleNames []string `query:"security_rule_names,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AnalyticsGetEventStatisticsParams]'s query parameters as
// `url.Values`.
func (r AnalyticsGetEventStatisticsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// A request characteristics dimension
type AnalyticsGetEventStatisticsParamsDimension string

const (
	AnalyticsGetEventStatisticsParamsDimensionCountry   AnalyticsGetEventStatisticsParamsDimension = "country"
	AnalyticsGetEventStatisticsParamsDimensionRule      AnalyticsGetEventStatisticsParamsDimension = "rule"
	AnalyticsGetEventStatisticsParamsDimensionOrg       AnalyticsGetEventStatisticsParamsDimension = "org"
	AnalyticsGetEventStatisticsParamsDimensionIP        AnalyticsGetEventStatisticsParamsDimension = "ip"
	AnalyticsGetEventStatisticsParamsDimensionUseragent AnalyticsGetEventStatisticsParamsDimension = "useragent"
	AnalyticsGetEventStatisticsParamsDimensionTarget    AnalyticsGetEventStatisticsParamsDimension = "target"
)

type AnalyticsGetRequestsParams struct {
	// Filter data items starting from a specified date in ISO 8601 format
	Start string `query:"start" api:"required" json:"-"`
	// Filter data items up to a specified end date in ISO 8601 format. If not
	// provided, defaults to the current date and time.
	End param.Opt[string] `query:"end,omitzero" json:"-"`
	// Filter by URL path with a glob-like pattern.
	Path param.Opt[string] `query:"path,omitzero" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Sort data by given field.
	Ordering param.Opt[string] `query:"ordering,omitzero" json:"-"`
	// Filter data by a country code of the originating IP address in ISO 3166-1
	// alpha-2 format.
	Countries []string `query:"countries,omitzero" json:"-"`
	// Filter data by decision.
	//
	// Any of "blocked", "monitored", "allowed", "passed".
	Decision []string `query:"decision,omitzero" json:"-"`
	// List of domain IDs. Empty list means all domains belonging to the current
	// account.
	Domains []int64 `query:"domains,omitzero" json:"-"`
	// Exclude data by a country code of the originating IP address in ISO 3166-1
	// alpha-2 format.
	ExcludeCountries []string `query:"exclude_countries,omitzero" json:"-"`
	// Exclude data by domain ID.
	ExcludeDomains []int64 `query:"exclude_domains,omitzero" json:"-"`
	// Exclude traffic data by client IP.
	ExcludeIPs []string `query:"exclude_ips,omitzero" format:"ipvanyaddress" json:"-"`
	// Exclude data by reference IDs.
	ExcludeReferenceIDs []string `query:"exclude_reference_ids,omitzero" json:"-"`
	// Exclude data by name of a security rule matched the request.
	ExcludeSecurityRuleNames []string `query:"exclude_security_rule_names,omitzero" json:"-"`
	// Exclude data by session IDs.
	ExcludeSessionIDs []string `query:"exclude_session_ids,omitzero" json:"-"`
	// Filter by HTTP methods
	//
	// Any of "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT", "TRACE".
	HTTPMethods []string `query:"http_methods,omitzero" json:"-"`
	// Filter traffic data by client IP.
	IPs []string `query:"ips,omitzero" format:"ipvanyaddress" json:"-"`
	// Filter data by optional action.
	//
	// Any of "captcha", "challenge".
	OptionalAction []string `query:"optional_action,omitzero" json:"-"`
	// Filter data by reference IDs.
	ReferenceIDs []string `query:"reference_ids,omitzero" json:"-"`
	// Filter data by request IDs.
	RequestIDs []string `query:"request_ids,omitzero" json:"-"`
	// Filter data by name of a security rule matched the request.
	SecurityRuleNames []string `query:"security_rule_names,omitzero" json:"-"`
	// Filter data by session IDs.
	SessionIDs []string `query:"session_ids,omitzero" json:"-"`
	// Filter data by HTTP response status code.
	StatusCodes []int64 `query:"status_codes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AnalyticsGetRequestsParams]'s query parameters as
// `url.Values`.
func (r AnalyticsGetRequestsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AnalyticsGetTrafficParams struct {
	// Specifies the granularity of the result data.
	//
	// Any of "daily", "hourly", "minutely".
	Resolution AnalyticsGetTrafficParamsResolution `query:"resolution,omitzero" api:"required" json:"-"`
	// Filter data items starting from a specified date in ISO 8601 format
	Start string `query:"start" api:"required" json:"-"`
	// Filter data items up to a specified end date in ISO 8601 format. If not
	// provided, defaults to the current date and time.
	End param.Opt[string] `query:"end,omitzero" json:"-"`
	// Optional explicit aggregation bucket width in seconds. When supplied,
	// `bucket_size` supersedes `resolution` for aggregation granularity.
	//
	// Any of 60, 300, 600, 900, 1800, 3600, 7200, 10800, 21600, 43200, 86400.
	BucketSize int64 `query:"bucket_size,omitzero" json:"-"`
	// List of domain IDs. Empty list means all domains belonging to the current
	// account.
	Domains []int64 `query:"domains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AnalyticsGetTrafficParams]'s query parameters as
// `url.Values`.
func (r AnalyticsGetTrafficParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the granularity of the result data.
type AnalyticsGetTrafficParamsResolution string

const (
	AnalyticsGetTrafficParamsResolutionDaily    AnalyticsGetTrafficParamsResolution = "daily"
	AnalyticsGetTrafficParamsResolutionHourly   AnalyticsGetTrafficParamsResolution = "hourly"
	AnalyticsGetTrafficParamsResolutionMinutely AnalyticsGetTrafficParamsResolution = "minutely"
)

type AnalyticsGetTrafficFilteredParams struct {
	// Specifies the granularity of the result data.
	//
	// Any of "daily", "hourly", "minutely".
	Resolution AnalyticsGetTrafficFilteredParamsResolution `query:"resolution,omitzero" api:"required" json:"-"`
	// Filter data items starting from a specified date in ISO 8601 format
	Start string `query:"start" api:"required" json:"-"`
	// Filter data items up to a specified end date in ISO 8601 format. If not
	// provided, defaults to the current date and time.
	End param.Opt[string] `query:"end,omitzero" json:"-"`
	// Filter by URL path with a glob-like pattern.
	Path param.Opt[string] `query:"path,omitzero" json:"-"`
	// Optional explicit aggregation bucket width in seconds. When supplied,
	// `bucket_size` supersedes `resolution` for aggregation granularity.
	//
	// Any of 60, 300, 600, 900, 1800, 3600, 7200, 10800, 21600, 43200, 86400.
	BucketSize int64 `query:"bucket_size,omitzero" json:"-"`
	// Filter data by a country code of the originating IP address in ISO 3166-1
	// alpha-2 format.
	Countries []string `query:"countries,omitzero" json:"-"`
	// Filter data by decision.
	//
	// Any of "blocked", "monitored", "allowed", "passed".
	Decision []string `query:"decision,omitzero" json:"-"`
	// List of domain IDs. Empty list means all domains belonging to the current
	// account.
	Domains []int64 `query:"domains,omitzero" json:"-"`
	// Exclude data by a country code of the originating IP address in ISO 3166-1
	// alpha-2 format.
	ExcludeCountries []string `query:"exclude_countries,omitzero" json:"-"`
	// Exclude data by domain ID.
	ExcludeDomains []int64 `query:"exclude_domains,omitzero" json:"-"`
	// Exclude traffic data by client IP.
	ExcludeIPs []string `query:"exclude_ips,omitzero" format:"ipvanyaddress" json:"-"`
	// Exclude data by reference IDs.
	ExcludeReferenceIDs []string `query:"exclude_reference_ids,omitzero" json:"-"`
	// Exclude data by name of a security rule matched the request.
	ExcludeSecurityRuleNames []string `query:"exclude_security_rule_names,omitzero" json:"-"`
	// Exclude data by session IDs.
	ExcludeSessionIDs []string `query:"exclude_session_ids,omitzero" json:"-"`
	// Filter by HTTP methods
	//
	// Any of "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT", "TRACE".
	HTTPMethods []string `query:"http_methods,omitzero" json:"-"`
	// Filter traffic data by client IP.
	IPs []string `query:"ips,omitzero" format:"ipvanyaddress" json:"-"`
	// Filter data by optional action.
	//
	// Any of "captcha", "challenge".
	OptionalAction []string `query:"optional_action,omitzero" json:"-"`
	// Filter data by reference IDs.
	ReferenceIDs []string `query:"reference_ids,omitzero" json:"-"`
	// Filter data by request IDs.
	RequestIDs []string `query:"request_ids,omitzero" json:"-"`
	// Filter data by name of a security rule matched the request.
	SecurityRuleNames []string `query:"security_rule_names,omitzero" json:"-"`
	// Filter data by session IDs.
	SessionIDs []string `query:"session_ids,omitzero" json:"-"`
	// Filter data by HTTP response status code.
	StatusCodes []int64 `query:"status_codes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AnalyticsGetTrafficFilteredParams]'s query parameters as
// `url.Values`.
func (r AnalyticsGetTrafficFilteredParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the granularity of the result data.
type AnalyticsGetTrafficFilteredParamsResolution string

const (
	AnalyticsGetTrafficFilteredParamsResolutionDaily    AnalyticsGetTrafficFilteredParamsResolution = "daily"
	AnalyticsGetTrafficFilteredParamsResolutionHourly   AnalyticsGetTrafficFilteredParamsResolution = "hourly"
	AnalyticsGetTrafficFilteredParamsResolutionMinutely AnalyticsGetTrafficFilteredParamsResolution = "minutely"
)
