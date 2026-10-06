// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
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

// IPInfoService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPInfoService] method instead.
type IPInfoService struct {
	Options []option.RequestOption
	Metrics IPInfoMetricService
}

// NewIPInfoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPInfoService(opts ...option.RequestOption) (r IPInfoService) {
	r = IPInfoService{}
	r.Options = opts
	r.Metrics = NewIPInfoMetricService(opts...)
	return
}

// Retrieve a time-series of attacks originating from a specified IP address,
// scoped to the authenticated account's domains
func (r *IPInfoService) GetAccountAttackTimeSeries(ctx context.Context, query IPInfoGetAccountAttackTimeSeriesParams, opts ...option.RequestOption) (res *[]WaapTimeSeriesAttack, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v2/ip-info/attack-time-series"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve comprehensive metrics for an IP address including distinct counts of
// attack types, domains, paths, and user agents across all account domains
func (r *IPInfoService) GetAccountCounts(ctx context.Context, query IPInfoGetAccountCountsParams, opts ...option.RequestOption) (res *WaapIPInfoCountsV2, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v2/ip-info/counts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve up to 10 top user agents associated with an IP across all account
// domains.
func (r *IPInfoService) GetAccountTopUserAgents(ctx context.Context, query IPInfoGetAccountTopUserAgentsParams, opts ...option.RequestOption) (res *[]WaapTopUserAgentDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v2/ip-info/top-user-agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve a time-series of attacks originating from a specified IP address.
func (r *IPInfoService) GetAttackTimeSeries(ctx context.Context, query IPInfoGetAttackTimeSeriesParams, opts ...option.RequestOption) (res *[]WaapTimeSeriesAttack, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/ip-info/attack-time-series"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve metrics, which enumerate blocked requests originating from a specific
// IP to a domain, grouped by rule name and taken action. Each metric provides
// insights into the request count blocked under a specific rule and the
// corresponding action that was executed.
func (r *IPInfoService) GetBlockedRequests(ctx context.Context, query IPInfoGetBlockedRequestsParams, opts ...option.RequestOption) (res *[]WaapRuleBlockedRequests, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/ip-info/blocked-requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Fetch and analyze DDoS (Distributed Denial of Service) attack metrics for a
// specified IP address. The endpoint provides time-series data, enabling users to
// evaluate the frequency and intensity of attacks across various time intervals,
// and it returns metrics in Prometheus format to offer a systematic view of DDoS
// attack patterns.
func (r *IPInfoService) GetDDOSAttackSeries(ctx context.Context, query IPInfoGetDDOSAttackSeriesParams, opts ...option.RequestOption) (res *WaapIPDDOSInfoModel, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/ip-info/ddos"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Fetch details about a particular IP address, including WHOIS data, risk score,
// and additional tags.
func (r *IPInfoService) GetIPInfo(ctx context.Context, query IPInfoGetIPInfoParams, opts ...option.RequestOption) (res *WaapIPInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/ip-info/ip-info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns top attack types with associated domains, paths, and counts across all
// account domains
func (r *IPInfoService) GetTopAttacks(ctx context.Context, query IPInfoGetTopAttacksParams, opts ...option.RequestOption) (res *[]WaapTopAttack, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/ip-info/top-attacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve top attacked domains for an IP address, scoped to the authenticated
// account's domains
func (r *IPInfoService) GetTopDomains(ctx context.Context, query IPInfoGetTopDomainsParams, opts ...option.RequestOption) (res *[]WaapTopDomainDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/ip-info/top-domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve top attacked paths for an IP address, scoped to the authenticated
// account's domains
func (r *IPInfoService) GetTopPaths(ctx context.Context, query IPInfoGetTopPathsParams, opts ...option.RequestOption) (res *[]WaapTopPathDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v2/ip-info/top-paths"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a list of the top 10 URLs accessed by a specified IP address within a
// specific domain. This data is vital to understand user navigation patterns,
// pinpoint high-traffic pages, and facilitate more targeted enhancements or
// security monitoring based on URL popularity.
func (r *IPInfoService) GetTopURLs(ctx context.Context, query IPInfoGetTopURLsParams, opts ...option.RequestOption) (res *[]WaapTopURL, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/ip-info/top-urls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the top 10 user agents interacting with a specified domain, filtered by
// IP.
func (r *IPInfoService) GetTopUserAgents(ctx context.Context, query IPInfoGetTopUserAgentsParams, opts ...option.RequestOption) (res *[]WaapTopUserAgent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/ip-info/top-user-agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Obtain the top 10 user sessions interfacing with a particular domain, identified
// by IP.
func (r *IPInfoService) GetTopUserSessions(ctx context.Context, query IPInfoGetTopUserSessionsParams, opts ...option.RequestOption) (res *[]WaapTopSession, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/ip-info/top-sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve a list of countries attacked by the specified IP address
func (r *IPInfoService) ListAttackedCountries(ctx context.Context, query IPInfoListAttackedCountriesParams, opts ...option.RequestOption) (res *[]WaapIPCountryAttack, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/ip-info/attack-map"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type WaapIPCountryAttack struct {
	// The number of attacks from the specified IP address to the country
	Count int64 `json:"count" api:"required"`
	// An ISO 3166-1 alpha-2 formatted string representing the country that was
	// attacked
	Country string `json:"country" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Country     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapIPCountryAttack) RawJSON() string { return r.JSON.raw }
func (r *WaapIPCountryAttack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapIPDDOSInfoModel struct {
	// Indicates if the IP is tagged as a botnet client
	BotnetClient bool `json:"botnet_client" api:"required"`
	// The time series data for the DDoS attacks from the IP address
	TimeSeries []WaapIPDDOSInfoModelTimeSeries `json:"time_series" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BotnetClient respjson.Field
		TimeSeries   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapIPDDOSInfoModel) RawJSON() string { return r.JSON.raw }
func (r *WaapIPDDOSInfoModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapIPDDOSInfoModelTimeSeries struct {
	// The number of attacks
	Count int64 `json:"count" api:"required"`
	// The timestamp of the time series item as a POSIX timestamp
	Timestamp int64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapIPDDOSInfoModelTimeSeries) RawJSON() string { return r.JSON.raw }
func (r *WaapIPDDOSInfoModelTimeSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapIPInfo struct {
	// The risk score of the IP address
	//
	// Any of "NO_RISK", "LOW", "MEDIUM", "HIGH", "EXTREME", "NOT_ENOUGH_DATA".
	RiskScore WaapIPInfoRiskScore `json:"risk_score" api:"required"`
	// The tags associated with the IP address that affect the risk score
	Tags []string `json:"tags" api:"required"`
	// The WHOIS information for the IP address
	Whois WaapIPInfoWhois `json:"whois" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RiskScore   respjson.Field
		Tags        respjson.Field
		Whois       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapIPInfo) RawJSON() string { return r.JSON.raw }
func (r *WaapIPInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The risk score of the IP address
type WaapIPInfoRiskScore string

const (
	WaapIPInfoRiskScoreNoRisk        WaapIPInfoRiskScore = "NO_RISK"
	WaapIPInfoRiskScoreLow           WaapIPInfoRiskScore = "LOW"
	WaapIPInfoRiskScoreMedium        WaapIPInfoRiskScore = "MEDIUM"
	WaapIPInfoRiskScoreHigh          WaapIPInfoRiskScore = "HIGH"
	WaapIPInfoRiskScoreExtreme       WaapIPInfoRiskScore = "EXTREME"
	WaapIPInfoRiskScoreNotEnoughData WaapIPInfoRiskScore = "NOT_ENOUGH_DATA"
)

// The WHOIS information for the IP address
type WaapIPInfoWhois struct {
	// The abuse mail
	AbuseMail string `json:"abuse_mail" api:"nullable"`
	// The CIDR
	Cidr int64 `json:"cidr" api:"nullable"`
	// The country
	Country string `json:"country" api:"nullable"`
	// The network description
	NetDescription string `json:"net_description" api:"nullable"`
	// The network name
	NetName string `json:"net_name" api:"nullable"`
	// The network range
	NetRange string `json:"net_range" api:"nullable"`
	// The network type
	NetType string `json:"net_type" api:"nullable"`
	// The organization ID
	OrgID string `json:"org_id" api:"nullable"`
	// The organization name
	OrgName string `json:"org_name" api:"nullable"`
	// The owner type
	OwnerType string `json:"owner_type" api:"nullable"`
	// The RIR
	Rir string `json:"rir" api:"nullable"`
	// The state
	State string `json:"state" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AbuseMail      respjson.Field
		Cidr           respjson.Field
		Country        respjson.Field
		NetDescription respjson.Field
		NetName        respjson.Field
		NetRange       respjson.Field
		NetType        respjson.Field
		OrgID          respjson.Field
		OrgName        respjson.Field
		OwnerType      respjson.Field
		Rir            respjson.Field
		State          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapIPInfoWhois) RawJSON() string { return r.JSON.raw }
func (r *WaapIPInfoWhois) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapIPInfoCountsV2 struct {
	// The number of distinct attack types (rules) that blocked requests from this IP
	AttackTypes int64 `json:"attack_types" api:"required"`
	// The number of requests from the IP address that were blocked
	BlockedRequests int64 `json:"blocked_requests" api:"required"`
	// The number of distinct domains accessed by this IP
	Domains int64 `json:"domains" api:"required"`
	// The number of distinct paths (URLs) accessed by this IP
	Paths int64 `json:"paths" api:"required"`
	// The total number of requests made by the IP address
	TotalRequests int64 `json:"total_requests" api:"required"`
	// The number of unique sessions from the IP address
	UniqueSessions int64 `json:"unique_sessions" api:"required"`
	// The number of distinct user agents used by this IP
	UserAgents int64 `json:"user_agents" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttackTypes     respjson.Field
		BlockedRequests respjson.Field
		Domains         respjson.Field
		Paths           respjson.Field
		TotalRequests   respjson.Field
		UniqueSessions  respjson.Field
		UserAgents      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapIPInfoCountsV2) RawJSON() string { return r.JSON.raw }
func (r *WaapIPInfoCountsV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapRuleBlockedRequests struct {
	// The action taken by the rule
	Action string `json:"action" api:"required"`
	// The number of requests blocked by the rule
	Count int64 `json:"count" api:"required"`
	// The name of the rule that blocked the request
	RuleName string `json:"rule_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Count       respjson.Field
		RuleName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRuleBlockedRequests) RawJSON() string { return r.JSON.raw }
func (r *WaapRuleBlockedRequests) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapTimeSeriesAttack struct {
	// The type of attack
	AttackType string `json:"attack_type" api:"required"`
	// The time series data
	Values []WaapTimeSeriesAttackValue `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttackType  respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTimeSeriesAttack) RawJSON() string { return r.JSON.raw }
func (r *WaapTimeSeriesAttack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapTimeSeriesAttackValue struct {
	// The number of attacks
	Count int64 `json:"count" api:"required"`
	// The timestamp of the time series item as a POSIX timestamp
	Timestamp int64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTimeSeriesAttackValue) RawJSON() string { return r.JSON.raw }
func (r *WaapTimeSeriesAttackValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapTopAttack struct {
	// The attack type (rule name)
	AttackType string `json:"attack_type" api:"required"`
	// Total number of attack requests (blocked + monitored)
	Attacks int64 `json:"attacks" api:"required"`
	// Number of blocked attack requests
	Blocked int64 `json:"blocked" api:"required"`
	// Full list of domains where this attack type was seen
	Domains []string `json:"domains"`
	// Full list of paths where this attack type was seen
	Paths []string `json:"paths"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttackType  respjson.Field
		Attacks     respjson.Field
		Blocked     respjson.Field
		Domains     respjson.Field
		Paths       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTopAttack) RawJSON() string { return r.JSON.raw }
func (r *WaapTopAttack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapTopDomainDetail struct {
	// Distinct rule names seen for this IP+domain, sorted alphabetically.
	AttackTypes []string `json:"attack_types" api:"required"`
	// Total detected attacks.
	Attacks int64 `json:"attacks" api:"required"`
	// Blocked attacks.
	Blocked int64 `json:"blocked" api:"required"`
	// First-seen timestamp for the IP+domain within retention period.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Domain name.
	Domain string `json:"domain" api:"required"`
	// `last_seen` - `first_seen` in seconds within retention.
	Duration int64 `json:"duration" api:"required"`
	// Distinct URL paths excluding query strings.
	Paths int64 `json:"paths" api:"required"`
	// Total requests.
	Requests int64 `json:"requests" api:"required"`
	// Distinct sessions.
	Sessions int64 `json:"sessions" api:"required"`
	// Distinct user agents.
	UserAgents int64 `json:"user_agents" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttackTypes respjson.Field
		Attacks     respjson.Field
		Blocked     respjson.Field
		Date        respjson.Field
		Domain      respjson.Field
		Duration    respjson.Field
		Paths       respjson.Field
		Requests    respjson.Field
		Sessions    respjson.Field
		UserAgents  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTopDomainDetail) RawJSON() string { return r.JSON.raw }
func (r *WaapTopDomainDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapTopPathDetail struct {
	// First seen timestamp for the IP+path within retention period
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Path value as stored in logs
	Path string `json:"path" api:"required"`
	// Distinct rule names seen for this IP+path
	AttackTypes []string `json:"attack_types"`
	// Total detected attacks (blocked + monitored/suppressed outcomes)
	Attacks int64 `json:"attacks"`
	// Blocked attacks
	Blocked int64 `json:"blocked"`
	// `last_seen` - `first_seen` in seconds within retention
	Duration int64 `json:"duration"`
	// Total requests
	Requests int64 `json:"requests"`
	// Distinct sessions
	Sessions int64 `json:"sessions"`
	// Distinct user agents
	UserAgents int64 `json:"user_agents"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Path        respjson.Field
		AttackTypes respjson.Field
		Attacks     respjson.Field
		Blocked     respjson.Field
		Duration    respjson.Field
		Requests    respjson.Field
		Sessions    respjson.Field
		UserAgents  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTopPathDetail) RawJSON() string { return r.JSON.raw }
func (r *WaapTopPathDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapTopSession struct {
	// The number of blocked requests in the session
	Blocked int64 `json:"blocked" api:"required"`
	// The duration of the session in seconds
	Duration float64 `json:"duration" api:"required"`
	// The number of requests in the session
	Requests int64 `json:"requests" api:"required"`
	// The session ID
	SessionID string `json:"session_id" api:"required" format:"uuid"`
	// The start time of the session as a POSIX timestamp
	StartTime time.Time `json:"start_time" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blocked     respjson.Field
		Duration    respjson.Field
		Requests    respjson.Field
		SessionID   respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTopSession) RawJSON() string { return r.JSON.raw }
func (r *WaapTopSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapTopURL struct {
	// The number of attacks to the URL
	Count int64 `json:"count" api:"required"`
	// The URL that was attacked
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTopURL) RawJSON() string { return r.JSON.raw }
func (r *WaapTopURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapTopUserAgent struct {
	// The number of requests made with the user agent
	Count int64 `json:"count" api:"required"`
	// The user agent that was used
	UserAgent string `json:"user_agent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		UserAgent   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTopUserAgent) RawJSON() string { return r.JSON.raw }
func (r *WaapTopUserAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapTopUserAgentDetail struct {
	// Total number of detected attacks.
	Attacks int64 `json:"attacks"`
	// Number of blocked attacks.
	Blocked int64 `json:"blocked"`
	// Number of distinct domains.
	Domains int64 `json:"domains"`
	// Number of distinct URL paths excluding query strings.
	Paths int64 `json:"paths"`
	// Total number of requests.
	Requests int64 `json:"requests"`
	// Number of distinct sessions.
	Sessions int64 `json:"sessions"`
	// The user agent string; missing values are represented as "".
	UserAgent string `json:"user_agent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attacks     respjson.Field
		Blocked     respjson.Field
		Domains     respjson.Field
		Paths       respjson.Field
		Requests    respjson.Field
		Sessions    respjson.Field
		UserAgent   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTopUserAgentDetail) RawJSON() string { return r.JSON.raw }
func (r *WaapTopUserAgentDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPInfoGetAccountAttackTimeSeriesParams struct {
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	// Number of days to look back for attack data
	Days param.Opt[int64] `query:"days,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetAccountAttackTimeSeriesParams]'s query parameters
// as `url.Values`.
func (r IPInfoGetAccountAttackTimeSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPInfoGetAccountCountsParams struct {
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetAccountCountsParams]'s query parameters as
// `url.Values`.
func (r IPInfoGetAccountCountsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPInfoGetAccountTopUserAgentsParams struct {
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetAccountTopUserAgentsParams]'s query parameters as
// `url.Values`.
func (r IPInfoGetAccountTopUserAgentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPInfoGetAttackTimeSeriesParams struct {
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetAttackTimeSeriesParams]'s query parameters as
// `url.Values`.
func (r IPInfoGetAttackTimeSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPInfoGetBlockedRequestsParams struct {
	// The identifier for a domain. When specified, the response will exclusively
	// contain data pertinent to the indicated domain, filtering out information from
	// other domains.
	DomainID int64 `query:"domain_id" api:"required" json:"-"`
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetBlockedRequestsParams]'s query parameters as
// `url.Values`.
func (r IPInfoGetBlockedRequestsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPInfoGetDDOSAttackSeriesParams struct {
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetDDOSAttackSeriesParams]'s query parameters as
// `url.Values`.
func (r IPInfoGetDDOSAttackSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPInfoGetIPInfoParams struct {
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetIPInfoParams]'s query parameters as `url.Values`.
func (r IPInfoGetIPInfoParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPInfoGetTopAttacksParams struct {
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetTopAttacksParams]'s query parameters as
// `url.Values`.
func (r IPInfoGetTopAttacksParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPInfoGetTopDomainsParams struct {
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetTopDomainsParams]'s query parameters as
// `url.Values`.
func (r IPInfoGetTopDomainsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPInfoGetTopPathsParams struct {
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetTopPathsParams]'s query parameters as
// `url.Values`.
func (r IPInfoGetTopPathsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPInfoGetTopURLsParams struct {
	// The identifier for a domain. When specified, the response will exclusively
	// contain data pertinent to the indicated domain, filtering out information from
	// other domains.
	DomainID int64 `query:"domain_id" api:"required" json:"-"`
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetTopURLsParams]'s query parameters as `url.Values`.
func (r IPInfoGetTopURLsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPInfoGetTopUserAgentsParams struct {
	// The identifier for a domain. When specified, the response will exclusively
	// contain data pertinent to the indicated domain, filtering out information from
	// other domains.
	DomainID int64 `query:"domain_id" api:"required" json:"-"`
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetTopUserAgentsParams]'s query parameters as
// `url.Values`.
func (r IPInfoGetTopUserAgentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPInfoGetTopUserSessionsParams struct {
	// The identifier for a domain. When specified, the response will exclusively
	// contain data pertinent to the indicated domain, filtering out information from
	// other domains.
	DomainID int64 `query:"domain_id" api:"required" json:"-"`
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetTopUserSessionsParams]'s query parameters as
// `url.Values`.
func (r IPInfoGetTopUserSessionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPInfoListAttackedCountriesParams struct {
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoListAttackedCountriesParams]'s query parameters as
// `url.Values`.
func (r IPInfoListAttackedCountriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
