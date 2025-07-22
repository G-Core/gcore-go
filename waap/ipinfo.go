// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// IPInfoService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPInfoService] method instead.
type IPInfoService struct {
	Options []option.RequestOption
}

// NewIPInfoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPInfoService(opts ...option.RequestOption) (r IPInfoService) {
	r = IPInfoService{}
	r.Options = opts
	return
}

// Fetch details about a particular IP address, including WHOIS data, risk score,
// and additional tags.
func (r *IPInfoService) Get(ctx context.Context, query IPInfoGetParams, opts ...option.RequestOption) (res *WaapIPInfo, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/ip-info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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

// Retrieve metrics encompassing the counts of total requests, blocked requests and
// unique sessions associated with a specified IP address. Metrics provide a
// statistical overview, aiding in analyzing the interaction and access patterns of
// the IP address in context.
func (r *IPInfoService) GetCounts(ctx context.Context, query IPInfoGetCountsParams, opts ...option.RequestOption) (res *WaapIPInfoCounts, err error) {
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
func (r *IPInfoService) GetDDOSAttackSeries(ctx context.Context, query IPInfoGetDDOSAttackSeriesParams, opts ...option.RequestOption) (res *WaapIPDDOSInfoModel, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/ddos"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Obtain the top 10 user sessions interfacing with a particular domain, identified
// by IP.
func (r *IPInfoService) GetTopSessions(ctx context.Context, query IPInfoGetTopSessionsParams, opts ...option.RequestOption) (res *[]WaapTopSession, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/top-sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of the top 10 URLs accessed by a specified IP address within a
// specific domain. This data is vital to understand user navigation patterns,
// pinpoint high-traffic pages, and facilitate more targeted enhancements or
// security monitoring based on URL popularity.
func (r *IPInfoService) GetTopURLs(ctx context.Context, query IPInfoGetTopURLsParams, opts ...option.RequestOption) (res *[]WaapTopURL, err error) {
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

// Retrieve a list of countries attacked by the specified IP address
func (r *IPInfoService) ListAttackedCountries(ctx context.Context, query IPInfoListAttackedCountriesParams, opts ...option.RequestOption) (res *[]WaapIPCountryAttack, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/ip-info/attack-map"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type IPInfoGetParams struct {
	// The IP address to check
	IP string `query:"ip,required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetParams]'s query parameters as `url.Values`.
func (r IPInfoGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
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

type IPInfoGetCountsParams struct {
	// The IP address to check
	IP string `query:"ip,required" format:"ipv4" json:"-"`
	// The identifier for a domain. When specified, the response will exclusively
	// contain data pertinent to the indicated domain, filtering out information from
	// other domains.
	DomainID param.Opt[int64] `query:"domain_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetCountsParams]'s query parameters as `url.Values`.
func (r IPInfoGetCountsParams) URLQuery() (v url.Values, err error) {
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

type IPInfoGetTopSessionsParams struct {
	// The identifier for a domain. When specified, the response will exclusively
	// contain data pertinent to the indicated domain, filtering out information from
	// other domains.
	DomainID int64 `query:"domain_id,required" json:"-"`
	// The IP address to check
	IP string `query:"ip,required" format:"ipv4" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoGetTopSessionsParams]'s query parameters as
// `url.Values`.
func (r IPInfoGetTopSessionsParams) URLQuery() (v url.Values, err error) {
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
