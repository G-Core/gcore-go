// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
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

// Retrieve a time-series of attacks originating from a specified IP address.
func (r *IPInfoService) GetAttackTimeSeries(ctx context.Context, query IPInfoGetAttackTimeSeriesParams, opts ...option.RequestOption) (res *[]WaapTimeSeriesAttack, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/attack-time-series"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve metrics, which enumerate blocked requests originating from a specific
// IP to a domain, grouped by rule name and taken action. Each metric provides
// insights into the request count blocked under a specific rule and the
// corresponding action that was executed.
func (r *IPInfoService) GetBlockedRequests(ctx context.Context, query IPInfoGetBlockedRequestsParams, opts ...option.RequestOption) (res *[]WaapRuleBlockedRequests, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/blocked-requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch and analyze DDoS (Distributed Denial of Service) attack metrics for a
// specified IP address. The endpoint provides time-series data, enabling users to
// evaluate the frequency and intensity of attacks across various time intervals,
// and it returns metrics in Prometheus format to offer a systematic view of DDoS
// attack patterns.
func (r *IPInfoService) GetDDOSAttackSeries(ctx context.Context, query IPInfoGetDDOSAttackSeriesParams, opts ...option.RequestOption) (res *WaapIPDDOSInfoModel, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/ddos"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch details about a particular IP address, including WHOIS data, risk score,
// and additional tags.
func (r *IPInfoService) GetIPInfo(ctx context.Context, query IPInfoGetIPInfoParams, opts ...option.RequestOption) (res *IPInfoGetIPInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/ip-info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of the top 10 URLs accessed by a specified IP address within a
// specific domain. This data is vital to understand user navigation patterns,
// pinpoint high-traffic pages, and facilitate more targeted enhancements or
// security monitoring based on URL popularity.
func (r *IPInfoService) GetTopURLs(ctx context.Context, query IPInfoGetTopURLsParams, opts ...option.RequestOption) (res *[]IPInfoGetTopURLsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/top-urls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve the top 10 user agents interacting with a specified domain, filtered by
// IP.
func (r *IPInfoService) GetTopUserAgents(ctx context.Context, query IPInfoGetTopUserAgentsParams, opts ...option.RequestOption) (res *[]WaapTopUserAgent, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/top-user-agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Obtain the top 10 user sessions interfacing with a particular domain, identified
// by IP.
func (r *IPInfoService) GetTopUserSessions(ctx context.Context, query IPInfoGetTopUserSessionsParams, opts ...option.RequestOption) (res *[]WaapTopSession, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/top-sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a list of countries attacked by the specified IP address
func (r *IPInfoService) ListAttackedCountries(ctx context.Context, query IPInfoListAttackedCountriesParams, opts ...option.RequestOption) (res *[]WaapIPCountryAttack, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/attack-map"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WaapIPCountryAttack struct {
	// The number of attacks from the specified IP address to the country
	Count int64 `json:"count,required"`
	// An ISO 3166-1 alpha-2 formatted string representing the country that was
	// attacked
	Country string `json:"country,required"`
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
	BotnetClient bool `json:"botnet_client,required"`
	// The time series data for the DDoS attacks from the IP address
	TimeSeries []WaapIPDDOSInfoModelTimeSery `json:"time_series,required"`
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

type WaapIPDDOSInfoModelTimeSery struct {
	// The number of attacks
	Count int64 `json:"count,required"`
	// The timestamp of the time series item as a POSIX timestamp
	Timestamp int64 `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapIPDDOSInfoModelTimeSery) RawJSON() string { return r.JSON.raw }
func (r *WaapIPDDOSInfoModelTimeSery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapRuleBlockedRequests struct {
	// The action taken by the rule
	Action string `json:"action,required"`
	// The number of requests blocked by the rule
	Count int64 `json:"count,required"`
	// The name of the rule that blocked the request
	RuleName string `json:"rule_name,required"`
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
	AttackType string `json:"attack_type,required"`
	// The time series data
	Values []WaapTimeSeriesAttackValue `json:"values,required"`
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
	Count int64 `json:"count,required"`
	// The timestamp of the time series item as a POSIX timestamp
	Timestamp int64 `json:"timestamp,required"`
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

type WaapTopSession struct {
	// The number of blocked requests in the session
	Blocked int64 `json:"blocked,required"`
	// The duration of the session in seconds
	Duration float64 `json:"duration,required"`
	// The number of requests in the session
	Requests int64 `json:"requests,required"`
	// The session ID
	SessionID string `json:"session_id,required" format:"uuid"`
	// The start time of the session as a POSIX timestamp
	StartTime time.Time `json:"start_time,required" format:"date-time"`
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

type WaapTopUserAgent struct {
	// The number of requests made with the user agent
	Count int64 `json:"count,required"`
	// The user agent that was used
	UserAgent string `json:"user_agent,required"`
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

type IPInfoGetIPInfoResponse struct {
	// The risk score of the IP address
	//
	// Any of "NO_RISK", "LOW", "MEDIUM", "HIGH", "EXTREME", "NOT_ENOUGH_DATA".
	RiskScore IPInfoGetIPInfoResponseRiskScore `json:"risk_score,required"`
	// The tags associated with the IP address that affect the risk score
	Tags []string `json:"tags,required"`
	// The WHOIS information for the IP address
	Whois IPInfoGetIPInfoResponseWhois `json:"whois,required"`
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
func (r IPInfoGetIPInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *IPInfoGetIPInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The risk score of the IP address
type IPInfoGetIPInfoResponseRiskScore string

const (
	IPInfoGetIPInfoResponseRiskScoreNoRisk        IPInfoGetIPInfoResponseRiskScore = "NO_RISK"
	IPInfoGetIPInfoResponseRiskScoreLow           IPInfoGetIPInfoResponseRiskScore = "LOW"
	IPInfoGetIPInfoResponseRiskScoreMedium        IPInfoGetIPInfoResponseRiskScore = "MEDIUM"
	IPInfoGetIPInfoResponseRiskScoreHigh          IPInfoGetIPInfoResponseRiskScore = "HIGH"
	IPInfoGetIPInfoResponseRiskScoreExtreme       IPInfoGetIPInfoResponseRiskScore = "EXTREME"
	IPInfoGetIPInfoResponseRiskScoreNotEnoughData IPInfoGetIPInfoResponseRiskScore = "NOT_ENOUGH_DATA"
)

// The WHOIS information for the IP address
type IPInfoGetIPInfoResponseWhois struct {
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
	State string `json:"state,nullable"`
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
func (r IPInfoGetIPInfoResponseWhois) RawJSON() string { return r.JSON.raw }
func (r *IPInfoGetIPInfoResponseWhois) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPInfoGetTopURLsResponse struct {
	// The number of attacks to the URL
	Count int64 `json:"count,required"`
	// The URL that was attacked
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPInfoGetTopURLsResponse) RawJSON() string { return r.JSON.raw }
func (r *IPInfoGetTopURLsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPInfoGetAttackTimeSeriesParams struct {
	// The IP address to check
	IP string `query:"ip,required" format:"ipv4" json:"-"`
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
	DomainID int64 `query:"domain_id,required" json:"-"`
	// The IP address to check
	IP string `query:"ip,required" format:"ipv4" json:"-"`
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
	IP string `query:"ip,required" format:"ipv4" json:"-"`
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
	IP string `query:"ip,required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetIPInfoParams]'s query parameters as `url.Values`.
func (r IPInfoGetIPInfoParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPInfoGetTopURLsParams struct {
	// The identifier for a domain. When specified, the response will exclusively
	// contain data pertinent to the indicated domain, filtering out information from
	// other domains.
	DomainID int64 `query:"domain_id,required" json:"-"`
	// The IP address to check
	IP string `query:"ip,required" format:"ipv4" json:"-"`
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
	DomainID int64 `query:"domain_id,required" json:"-"`
	// The IP address to check
	IP string `query:"ip,required" format:"ipv4" json:"-"`
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
	DomainID int64 `query:"domain_id,required" json:"-"`
	// The IP address to check
	IP string `query:"ip,required" format:"ipv4" json:"-"`
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
	IP string `query:"ip,required" format:"ipv4" json:"-"`
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
