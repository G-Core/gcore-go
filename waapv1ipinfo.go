// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// WaapV1IPInfoService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1IPInfoService] method instead.
type WaapV1IPInfoService struct {
	Options []option.RequestOption
}

// NewWaapV1IPInfoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWaapV1IPInfoService(opts ...option.RequestOption) (r *WaapV1IPInfoService) {
	r = &WaapV1IPInfoService{}
	r.Options = opts
	return
}

// Retrieve a list of countries attacked by the specified IP address
func (r *WaapV1IPInfoService) GetAttackMap(ctx context.Context, query WaapV1IPInfoGetAttackMapParams, opts ...option.RequestOption) (res *[]WaapV1IPInfoGetAttackMapResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/attack-map"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a time-series of attacks originating from a specified IP address.
func (r *WaapV1IPInfoService) GetAttackTimeSeries(ctx context.Context, query WaapV1IPInfoGetAttackTimeSeriesParams, opts ...option.RequestOption) (res *[]WaapV1IPInfoGetAttackTimeSeriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/attack-time-series"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve metrics, which enumerate blocked requests originating from a specific
// IP to a domain, grouped by rule name and taken action. Each metric provides
// insights into the request count blocked under a specific rule and the
// corresponding action that was executed.
func (r *WaapV1IPInfoService) GetBlockedRequests(ctx context.Context, query WaapV1IPInfoGetBlockedRequestsParams, opts ...option.RequestOption) (res *[]WaapV1IPInfoGetBlockedRequestsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/blocked-requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve metrics encompassing the counts of total requests, blocked requests and
// unique sessions associated with a specified IP address. Metrics provide a
// statistical overview, aiding in analyzing the interaction and access patterns of
// the IP address in context.
func (r *WaapV1IPInfoService) GetCounts(ctx context.Context, query WaapV1IPInfoGetCountsParams, opts ...option.RequestOption) (res *WaapV1IPInfoGetCountsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/counts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch and analyze DDoS (Distributed Denial of Service) attack metrics for a
// specified IP address. The endpoint provides time-series data, enabling users to
// evaluate the frequency and intensity of attacks across various time intervals,
// and it returns metrics in Prometheus format to offer a systematic view of DDoS
// attack patterns.
func (r *WaapV1IPInfoService) GetDdosMetrics(ctx context.Context, query WaapV1IPInfoGetDdosMetricsParams, opts ...option.RequestOption) (res *WaapV1IPInfoGetDdosMetricsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/ddos"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch details about a particular IP address, including WHOIS data, risk score,
// and additional tags.
func (r *WaapV1IPInfoService) GetIPInfo(ctx context.Context, query WaapV1IPInfoGetIPInfoParams, opts ...option.RequestOption) (res *WaapV1IPInfoGetIPInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/ip-info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Obtain the top 10 user sessions interfacing with a particular domain, identified
// by IP.
func (r *WaapV1IPInfoService) GetTopSessions(ctx context.Context, query WaapV1IPInfoGetTopSessionsParams, opts ...option.RequestOption) (res *[]WaapV1IPInfoGetTopSessionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/top-sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of the top 10 URLs accessed by a specified IP address within a
// specific domain. This data is vital to understand user navigation patterns,
// pinpoint high-traffic pages, and facilitate more targeted enhancements or
// security monitoring based on URL popularity.
func (r *WaapV1IPInfoService) GetTopURLs(ctx context.Context, query WaapV1IPInfoGetTopURLsParams, opts ...option.RequestOption) (res *[]WaapV1IPInfoGetTopURLsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/top-urls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve the top 10 user agents interacting with a specified domain, filtered by
// IP.
func (r *WaapV1IPInfoService) GetTopUserAgents(ctx context.Context, query WaapV1IPInfoGetTopUserAgentsParams, opts ...option.RequestOption) (res *[]WaapV1IPInfoGetTopUserAgentsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/top-user-agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TimeSeriesItem struct {
	// The number of attacks
	Count int64 `json:"count,required"`
	// The timestamp of the time series item as a POSIX timestamp
	Timestamp int64              `json:"timestamp,required"`
	JSON      timeSeriesItemJSON `json:"-"`
}

// timeSeriesItemJSON contains the JSON metadata for the struct [TimeSeriesItem]
type timeSeriesItemJSON struct {
	Count       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimeSeriesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timeSeriesItemJSON) RawJSON() string {
	return r.raw
}

type WaapV1IPInfoGetAttackMapResponse struct {
	// The number of attacks from the specified IP address to the country
	Count int64 `json:"count,required"`
	// An ISO 3166-1 alpha-2 formatted string representing the country that was
	// attacked
	Country string                               `json:"country,required"`
	JSON    waapV1IPInfoGetAttackMapResponseJSON `json:"-"`
}

// waapV1IPInfoGetAttackMapResponseJSON contains the JSON metadata for the struct
// [WaapV1IPInfoGetAttackMapResponse]
type waapV1IPInfoGetAttackMapResponseJSON struct {
	Count       apijson.Field
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1IPInfoGetAttackMapResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1IPInfoGetAttackMapResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1IPInfoGetAttackTimeSeriesResponse struct {
	// The type of attack
	AttackType string `json:"attack_type,required"`
	// The time series data
	Values []TimeSeriesItem                            `json:"values,required"`
	JSON   waapV1IPInfoGetAttackTimeSeriesResponseJSON `json:"-"`
}

// waapV1IPInfoGetAttackTimeSeriesResponseJSON contains the JSON metadata for the
// struct [WaapV1IPInfoGetAttackTimeSeriesResponse]
type waapV1IPInfoGetAttackTimeSeriesResponseJSON struct {
	AttackType  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1IPInfoGetAttackTimeSeriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1IPInfoGetAttackTimeSeriesResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1IPInfoGetBlockedRequestsResponse struct {
	// The action taken by the rule
	Action string `json:"action,required"`
	// The number of requests blocked by the rule
	Count int64 `json:"count,required"`
	// The name of the rule that blocked the request
	RuleName string                                     `json:"rule_name,required"`
	JSON     waapV1IPInfoGetBlockedRequestsResponseJSON `json:"-"`
}

// waapV1IPInfoGetBlockedRequestsResponseJSON contains the JSON metadata for the
// struct [WaapV1IPInfoGetBlockedRequestsResponse]
type waapV1IPInfoGetBlockedRequestsResponseJSON struct {
	Action      apijson.Field
	Count       apijson.Field
	RuleName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1IPInfoGetBlockedRequestsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1IPInfoGetBlockedRequestsResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1IPInfoGetCountsResponse struct {
	// The number of requests from the IP address that were blocked
	BlockedRequests int64 `json:"blocked_requests,required"`
	// The total number of requests made by the IP address
	TotalRequests int64 `json:"total_requests,required"`
	// The number of unique sessions from the IP address
	UniqueSessions int64                             `json:"unique_sessions,required"`
	JSON           waapV1IPInfoGetCountsResponseJSON `json:"-"`
}

// waapV1IPInfoGetCountsResponseJSON contains the JSON metadata for the struct
// [WaapV1IPInfoGetCountsResponse]
type waapV1IPInfoGetCountsResponseJSON struct {
	BlockedRequests apijson.Field
	TotalRequests   apijson.Field
	UniqueSessions  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WaapV1IPInfoGetCountsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1IPInfoGetCountsResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1IPInfoGetDdosMetricsResponse struct {
	// Indicates if the IP is tagged as a botnet client
	BotnetClient bool `json:"botnet_client,required"`
	// The time series data for the DDoS attacks from the IP address
	TimeSeries []TimeSeriesItem                       `json:"time_series,required"`
	JSON       waapV1IPInfoGetDdosMetricsResponseJSON `json:"-"`
}

// waapV1IPInfoGetDdosMetricsResponseJSON contains the JSON metadata for the struct
// [WaapV1IPInfoGetDdosMetricsResponse]
type waapV1IPInfoGetDdosMetricsResponseJSON struct {
	BotnetClient apijson.Field
	TimeSeries   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WaapV1IPInfoGetDdosMetricsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1IPInfoGetDdosMetricsResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1IPInfoGetIPInfoResponse struct {
	// The risk score of the IP address
	RiskScore WaapV1IPInfoGetIPInfoResponseRiskScore `json:"risk_score,required"`
	// The tags associated with the IP address that affect the risk score
	Tags []string `json:"tags,required"`
	// The WHOIS information for the IP address
	Whois WaapV1IPInfoGetIPInfoResponseWhois `json:"whois,required"`
	JSON  waapV1IPInfoGetIPInfoResponseJSON  `json:"-"`
}

// waapV1IPInfoGetIPInfoResponseJSON contains the JSON metadata for the struct
// [WaapV1IPInfoGetIPInfoResponse]
type waapV1IPInfoGetIPInfoResponseJSON struct {
	RiskScore   apijson.Field
	Tags        apijson.Field
	Whois       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1IPInfoGetIPInfoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1IPInfoGetIPInfoResponseJSON) RawJSON() string {
	return r.raw
}

// The risk score of the IP address
type WaapV1IPInfoGetIPInfoResponseRiskScore string

const (
	WaapV1IPInfoGetIPInfoResponseRiskScoreNoRisk        WaapV1IPInfoGetIPInfoResponseRiskScore = "NO_RISK"
	WaapV1IPInfoGetIPInfoResponseRiskScoreLow           WaapV1IPInfoGetIPInfoResponseRiskScore = "LOW"
	WaapV1IPInfoGetIPInfoResponseRiskScoreMedium        WaapV1IPInfoGetIPInfoResponseRiskScore = "MEDIUM"
	WaapV1IPInfoGetIPInfoResponseRiskScoreHigh          WaapV1IPInfoGetIPInfoResponseRiskScore = "HIGH"
	WaapV1IPInfoGetIPInfoResponseRiskScoreExtreme       WaapV1IPInfoGetIPInfoResponseRiskScore = "EXTREME"
	WaapV1IPInfoGetIPInfoResponseRiskScoreNotEnoughData WaapV1IPInfoGetIPInfoResponseRiskScore = "NOT_ENOUGH_DATA"
)

func (r WaapV1IPInfoGetIPInfoResponseRiskScore) IsKnown() bool {
	switch r {
	case WaapV1IPInfoGetIPInfoResponseRiskScoreNoRisk, WaapV1IPInfoGetIPInfoResponseRiskScoreLow, WaapV1IPInfoGetIPInfoResponseRiskScoreMedium, WaapV1IPInfoGetIPInfoResponseRiskScoreHigh, WaapV1IPInfoGetIPInfoResponseRiskScoreExtreme, WaapV1IPInfoGetIPInfoResponseRiskScoreNotEnoughData:
		return true
	}
	return false
}

// The WHOIS information for the IP address
type WaapV1IPInfoGetIPInfoResponseWhois struct {
	// The abuse mail
	AbuseMail string `json:"abuse_mail,nullable"`
	// The CIDR
	Cidr int64 `json:"cidr,nullable"`
	// The country
	Country string `json:"country,nullable"`
	// The network description
	NetDescription string `json:"net_description,nullable"`
	// The network name
	NetName string `json:"net_name,nullable"`
	// The network range
	NetRange string `json:"net_range,nullable"`
	// The network type
	NetType string `json:"net_type,nullable"`
	// The organization ID
	OrgID string `json:"org_id,nullable"`
	// The organization name
	OrgName string `json:"org_name,nullable"`
	// The owner type
	OwnerType string `json:"owner_type,nullable"`
	// The RIR
	Rir string `json:"rir,nullable"`
	// The state
	State string                                 `json:"state,nullable"`
	JSON  waapV1IPInfoGetIPInfoResponseWhoisJSON `json:"-"`
}

// waapV1IPInfoGetIPInfoResponseWhoisJSON contains the JSON metadata for the struct
// [WaapV1IPInfoGetIPInfoResponseWhois]
type waapV1IPInfoGetIPInfoResponseWhoisJSON struct {
	AbuseMail      apijson.Field
	Cidr           apijson.Field
	Country        apijson.Field
	NetDescription apijson.Field
	NetName        apijson.Field
	NetRange       apijson.Field
	NetType        apijson.Field
	OrgID          apijson.Field
	OrgName        apijson.Field
	OwnerType      apijson.Field
	Rir            apijson.Field
	State          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WaapV1IPInfoGetIPInfoResponseWhois) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1IPInfoGetIPInfoResponseWhoisJSON) RawJSON() string {
	return r.raw
}

type WaapV1IPInfoGetTopSessionsResponse struct {
	// The number of blocked requests in the session
	Blocked int64 `json:"blocked,required"`
	// The duration of the session in seconds
	Duration float64 `json:"duration,required"`
	// The number of requests in the session
	Requests int64 `json:"requests,required"`
	// The session ID
	SessionID string `json:"session_id,required" format:"uuid"`
	// The start time of the session as a POSIX timestamp
	StartTime time.Time                              `json:"start_time,required" format:"date-time"`
	JSON      waapV1IPInfoGetTopSessionsResponseJSON `json:"-"`
}

// waapV1IPInfoGetTopSessionsResponseJSON contains the JSON metadata for the struct
// [WaapV1IPInfoGetTopSessionsResponse]
type waapV1IPInfoGetTopSessionsResponseJSON struct {
	Blocked     apijson.Field
	Duration    apijson.Field
	Requests    apijson.Field
	SessionID   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1IPInfoGetTopSessionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1IPInfoGetTopSessionsResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1IPInfoGetTopURLsResponse struct {
	// The number of attacks to the URL
	Count int64 `json:"count,required"`
	// The URL that was attacked
	URL  string                             `json:"url,required"`
	JSON waapV1IPInfoGetTopURLsResponseJSON `json:"-"`
}

// waapV1IPInfoGetTopURLsResponseJSON contains the JSON metadata for the struct
// [WaapV1IPInfoGetTopURLsResponse]
type waapV1IPInfoGetTopURLsResponseJSON struct {
	Count       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1IPInfoGetTopURLsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1IPInfoGetTopURLsResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1IPInfoGetTopUserAgentsResponse struct {
	// The number of requests made with the user agent
	Count int64 `json:"count,required"`
	// The user agent that was used
	UserAgent string                                   `json:"user_agent,required"`
	JSON      waapV1IPInfoGetTopUserAgentsResponseJSON `json:"-"`
}

// waapV1IPInfoGetTopUserAgentsResponseJSON contains the JSON metadata for the
// struct [WaapV1IPInfoGetTopUserAgentsResponse]
type waapV1IPInfoGetTopUserAgentsResponseJSON struct {
	Count       apijson.Field
	UserAgent   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1IPInfoGetTopUserAgentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1IPInfoGetTopUserAgentsResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1IPInfoGetAttackMapParams struct {
	// The IP address to check
	IP param.Field[string] `query:"ip,required" format:"ipv4"`
}

// URLQuery serializes [WaapV1IPInfoGetAttackMapParams]'s query parameters as
// `url.Values`.
func (r WaapV1IPInfoGetAttackMapParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WaapV1IPInfoGetAttackTimeSeriesParams struct {
	// The IP address to check
	IP param.Field[string] `query:"ip,required" format:"ipv4"`
}

// URLQuery serializes [WaapV1IPInfoGetAttackTimeSeriesParams]'s query parameters
// as `url.Values`.
func (r WaapV1IPInfoGetAttackTimeSeriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WaapV1IPInfoGetBlockedRequestsParams struct {
	// The domain ID
	DomainID param.Field[int64] `query:"domain_id,required"`
	// The IP address to check
	IP param.Field[string] `query:"ip,required" format:"ipv4"`
}

// URLQuery serializes [WaapV1IPInfoGetBlockedRequestsParams]'s query parameters as
// `url.Values`.
func (r WaapV1IPInfoGetBlockedRequestsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WaapV1IPInfoGetCountsParams struct {
	// The IP address to check
	IP param.Field[string] `query:"ip,required" format:"ipv4"`
	// The identifier for a domain. When specified, the response will exclusively
	// contain data pertinent to the indicated domain, filtering out information from
	// other domains.
	DomainID param.Field[int64] `query:"domain_id"`
}

// URLQuery serializes [WaapV1IPInfoGetCountsParams]'s query parameters as
// `url.Values`.
func (r WaapV1IPInfoGetCountsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WaapV1IPInfoGetDdosMetricsParams struct {
	// The IP address to check
	IP param.Field[string] `query:"ip,required" format:"ipv4"`
}

// URLQuery serializes [WaapV1IPInfoGetDdosMetricsParams]'s query parameters as
// `url.Values`.
func (r WaapV1IPInfoGetDdosMetricsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WaapV1IPInfoGetIPInfoParams struct {
	// The IP address to check
	IP param.Field[string] `query:"ip,required" format:"ipv4"`
}

// URLQuery serializes [WaapV1IPInfoGetIPInfoParams]'s query parameters as
// `url.Values`.
func (r WaapV1IPInfoGetIPInfoParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WaapV1IPInfoGetTopSessionsParams struct {
	// The domain ID
	DomainID param.Field[int64] `query:"domain_id,required"`
	// The IP address to check
	IP param.Field[string] `query:"ip,required" format:"ipv4"`
}

// URLQuery serializes [WaapV1IPInfoGetTopSessionsParams]'s query parameters as
// `url.Values`.
func (r WaapV1IPInfoGetTopSessionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WaapV1IPInfoGetTopURLsParams struct {
	// The domain ID
	DomainID param.Field[int64] `query:"domain_id,required"`
	// The IP address to check
	IP param.Field[string] `query:"ip,required" format:"ipv4"`
}

// URLQuery serializes [WaapV1IPInfoGetTopURLsParams]'s query parameters as
// `url.Values`.
func (r WaapV1IPInfoGetTopURLsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WaapV1IPInfoGetTopUserAgentsParams struct {
	// The domain ID
	DomainID param.Field[int64] `query:"domain_id,required"`
	// The IP address to check
	IP param.Field[string] `query:"ip,required" format:"ipv4"`
}

// URLQuery serializes [WaapV1IPInfoGetTopUserAgentsParams]'s query parameters as
// `url.Values`.
func (r WaapV1IPInfoGetTopUserAgentsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
