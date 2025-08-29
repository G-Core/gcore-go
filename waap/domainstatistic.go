// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// DomainStatisticService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainStatisticService] method instead.
type DomainStatisticService struct {
	Options []option.RequestOption
}

// NewDomainStatisticService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDomainStatisticService(opts ...option.RequestOption) (r DomainStatisticService) {
	r = DomainStatisticService{}
	r.Options = opts
	return
}

// Retrieve a domain's DDoS attacks
func (r *DomainStatisticService) GetDDOSAttacks(ctx context.Context, domainID int64, query DomainStatisticGetDDOSAttacksParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapDDOSAttack], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/ddos-attacks", domainID)
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

// Retrieve a domain's DDoS attacks
func (r *DomainStatisticService) GetDDOSAttacksAutoPaging(ctx context.Context, domainID int64, query DomainStatisticGetDDOSAttacksParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapDDOSAttack] {
	return pagination.NewOffsetPageAutoPager(r.GetDDOSAttacks(ctx, domainID, query, opts...))
}

// Returns the top DDoS counts grouped by URL, User-Agent or IP
func (r *DomainStatisticService) GetDDOSInfo(ctx context.Context, domainID int64, query DomainStatisticGetDDOSInfoParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapDDOSInfo], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/ddos-info", domainID)
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

// Returns the top DDoS counts grouped by URL, User-Agent or IP
func (r *DomainStatisticService) GetDDOSInfoAutoPaging(ctx context.Context, domainID int64, query DomainStatisticGetDDOSInfoParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapDDOSInfo] {
	return pagination.NewOffsetPageAutoPager(r.GetDDOSInfo(ctx, domainID, query, opts...))
}

// Retrieve an domain's event statistics
func (r *DomainStatisticService) GetEventsAggregated(ctx context.Context, domainID int64, query DomainStatisticGetEventsAggregatedParams, opts ...option.RequestOption) (res *WaapEventStatistics, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/stats", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves all the available information for a request that matches a given
// request id
func (r *DomainStatisticService) GetRequestDetails(ctx context.Context, requestID string, query DomainStatisticGetRequestDetailsParams, opts ...option.RequestOption) (res *WaapRequestDetails, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/requests/%s/details", query.DomainID, requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a domain's requests data.
func (r *DomainStatisticService) GetRequestsSeries(ctx context.Context, domainID int64, query DomainStatisticGetRequestsSeriesParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapRequestSummary], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/requests", domainID)
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

// Retrieve a domain's requests data.
func (r *DomainStatisticService) GetRequestsSeriesAutoPaging(ctx context.Context, domainID int64, query DomainStatisticGetRequestsSeriesParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapRequestSummary] {
	return pagination.NewOffsetPageAutoPager(r.GetRequestsSeries(ctx, domainID, query, opts...))
}

// Retrieves a comprehensive report on a domain's traffic statistics based on
// Clickhouse. The report includes details such as API requests, blocked events,
// error counts, and many more traffic-related metrics.
func (r *DomainStatisticService) GetTrafficSeries(ctx context.Context, domainID int64, query DomainStatisticGetTrafficSeriesParams, opts ...option.RequestOption) (res *[]WaapTrafficMetrics, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/traffic", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A collection of total numbers of events with blocked results per criteria
type WaapBlockedStatistics struct {
	// A collection of event counts per action. The first item is the action's
	// abbreviation/full action name, and the second item is the number of events
	Action [][]WaapBlockedStatisticsActionUnion `json:"action,required"`
	// A collection of event counts per country of origin. The first item is the
	// country's ISO 3166-1 alpha-2, and the second item is the number of events
	Country [][]WaapBlockedStatisticsCountryUnion `json:"country,required"`
	// A collection of event counts per organization that owns the event's client IP.
	// The first item is the organization's name, and the second item is the number of
	// events
	Org [][]WaapBlockedStatisticsOrgUnion `json:"org,required"`
	// A collection of event counts per rule that triggered the event. The first item
	// is the rule's name, and the second item is the number of events
	RuleName [][]WaapBlockedStatisticsRuleNameUnion `json:"rule_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Country     respjson.Field
		Org         respjson.Field
		RuleName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapBlockedStatistics) RawJSON() string { return r.JSON.raw }
func (r *WaapBlockedStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapBlockedStatisticsActionUnion contains all possible properties and values
// from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapBlockedStatisticsActionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapBlockedStatisticsActionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapBlockedStatisticsActionUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapBlockedStatisticsActionUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapBlockedStatisticsActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapBlockedStatisticsCountryUnion contains all possible properties and values
// from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapBlockedStatisticsCountryUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapBlockedStatisticsCountryUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapBlockedStatisticsCountryUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapBlockedStatisticsCountryUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapBlockedStatisticsCountryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapBlockedStatisticsOrgUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapBlockedStatisticsOrgUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapBlockedStatisticsOrgUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapBlockedStatisticsOrgUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapBlockedStatisticsOrgUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapBlockedStatisticsOrgUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapBlockedStatisticsRuleNameUnion contains all possible properties and values
// from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapBlockedStatisticsRuleNameUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapBlockedStatisticsRuleNameUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapBlockedStatisticsRuleNameUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapBlockedStatisticsRuleNameUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapBlockedStatisticsRuleNameUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of total numbers of events per criteria
type WaapCountStatistics struct {
	// A collection of event counts per action. The first item is the action's
	// abbreviation/full action name, and the second item is the number of events
	Action [][]WaapCountStatisticsActionUnion `json:"action,required"`
	// A collection of event counts per country of origin. The first item is the
	// country's ISO 3166-1 alpha-2, and the second item is the number of events
	Country [][]WaapCountStatisticsCountryUnion `json:"country,required"`
	// A collection of event counts per organization that owns the event's client IP.
	// The first item is the organization's name, and the second item is the number of
	// events
	Org [][]WaapCountStatisticsOrgUnion `json:"org,required"`
	// A collection of event counts per rule that triggered the event. The first item
	// is the rule's name, and the second item is the number of events
	RuleName [][]WaapCountStatisticsRuleNameUnion `json:"rule_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Country     respjson.Field
		Org         respjson.Field
		RuleName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCountStatistics) RawJSON() string { return r.JSON.raw }
func (r *WaapCountStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapCountStatisticsActionUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapCountStatisticsActionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapCountStatisticsActionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapCountStatisticsActionUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapCountStatisticsActionUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapCountStatisticsActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapCountStatisticsCountryUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapCountStatisticsCountryUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapCountStatisticsCountryUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapCountStatisticsCountryUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapCountStatisticsCountryUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapCountStatisticsCountryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapCountStatisticsOrgUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapCountStatisticsOrgUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapCountStatisticsOrgUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapCountStatisticsOrgUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapCountStatisticsOrgUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapCountStatisticsOrgUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapCountStatisticsRuleNameUnion contains all possible properties and values
// from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapCountStatisticsRuleNameUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapCountStatisticsRuleNameUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapCountStatisticsRuleNameUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapCountStatisticsRuleNameUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapCountStatisticsRuleNameUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapDDOSAttack struct {
	// End time of DDoS attack
	EndTime time.Time `json:"end_time,nullable" format:"date-time"`
	// Start time of DDoS attack
	StartTime time.Time `json:"start_time" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTime     respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDDOSAttack) RawJSON() string { return r.JSON.raw }
func (r *WaapDDOSAttack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapDDOSInfo struct {
	// The number of requests made
	Count int64 `json:"count,required"`
	// The value for the grouped by type
	Identity string `json:"identity,required"`
	// Any of "URL", "IP", "User-Agent".
	Type WaapDDOSInfoType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Identity    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDDOSInfo) RawJSON() string { return r.JSON.raw }
func (r *WaapDDOSInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapDDOSInfoType string

const (
	WaapDDOSInfoTypeURL       WaapDDOSInfoType = "URL"
	WaapDDOSInfoTypeIP        WaapDDOSInfoType = "IP"
	WaapDDOSInfoTypeUserAgent WaapDDOSInfoType = "User-Agent"
)

// A collection of event metrics over a time span
type WaapEventStatistics struct {
	// A collection of total numbers of events with blocked results per criteria
	Blocked WaapBlockedStatistics `json:"blocked,required"`
	// A collection of total numbers of events per criteria
	Count WaapCountStatistics `json:"count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blocked     respjson.Field
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapEventStatistics) RawJSON() string { return r.JSON.raw }
func (r *WaapEventStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request's details used when displaying a single request.
type WaapRequestDetails struct {
	// Request ID
	ID string `json:"id,required"`
	// Request action
	Action string `json:"action,required"`
	// List of common tags
	CommonTags []WaapRequestDetailsCommonTag `json:"common_tags,required"`
	// Content type of request
	ContentType string `json:"content_type,required"`
	// Domain name
	Domain string `json:"domain,required"`
	// Status code for http request
	HTTPStatusCode int64 `json:"http_status_code,required"`
	// HTTP version of request
	HTTPVersion string `json:"http_version,required"`
	// ID of challenge that was generated
	IncidentID string `json:"incident_id,required"`
	// Request method
	Method string `json:"method,required"`
	// Network details
	Network WaapRequestDetailsNetwork `json:"network,required"`
	// Request path
	Path string `json:"path,required"`
	// List of shield tags
	PatternMatchedTags []WaapRequestDetailsPatternMatchedTag `json:"pattern_matched_tags,required"`
	// The query string of the request
	QueryString string `json:"query_string,required"`
	// Reference ID to identify user sanction
	ReferenceID string `json:"reference_id,required"`
	// HTTP request headers
	RequestHeaders map[string]any `json:"request_headers,required"`
	// The time of the request
	RequestTime string `json:"request_time,required"`
	// The type of the request that generated an event
	RequestType string `json:"request_type,required"`
	// The real domain name
	RequestedDomain string `json:"requested_domain,required"`
	// Time took to process all request
	ResponseTime string `json:"response_time,required"`
	// The result of a request
	//
	// Any of "passed", "blocked", "suppressed", "".
	Result WaapRequestDetailsResult `json:"result,required"`
	// ID of the triggered rule
	RuleID string `json:"rule_id,required"`
	// Name of the triggered rule
	RuleName string `json:"rule_name,required"`
	// The HTTP scheme of the request that generated an event
	Scheme string `json:"scheme,required"`
	// The number requests in session
	SessionRequestCount string `json:"session_request_count,required"`
	// List of traffic types
	TrafficTypes []string `json:"traffic_types,required"`
	// User agent
	UserAgent WaapRequestDetailsUserAgent `json:"user_agent,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Action              respjson.Field
		CommonTags          respjson.Field
		ContentType         respjson.Field
		Domain              respjson.Field
		HTTPStatusCode      respjson.Field
		HTTPVersion         respjson.Field
		IncidentID          respjson.Field
		Method              respjson.Field
		Network             respjson.Field
		Path                respjson.Field
		PatternMatchedTags  respjson.Field
		QueryString         respjson.Field
		ReferenceID         respjson.Field
		RequestHeaders      respjson.Field
		RequestTime         respjson.Field
		RequestType         respjson.Field
		RequestedDomain     respjson.Field
		ResponseTime        respjson.Field
		Result              respjson.Field
		RuleID              respjson.Field
		RuleName            respjson.Field
		Scheme              respjson.Field
		SessionRequestCount respjson.Field
		TrafficTypes        respjson.Field
		UserAgent           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRequestDetails) RawJSON() string { return r.JSON.raw }
func (r *WaapRequestDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Common tag details
type WaapRequestDetailsCommonTag struct {
	// Tag description information
	Description string `json:"description,required"`
	// The tag's display name
	DisplayName string `json:"display_name,required"`
	// Tag name
	Tag string `json:"tag,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		DisplayName respjson.Field
		Tag         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRequestDetailsCommonTag) RawJSON() string { return r.JSON.raw }
func (r *WaapRequestDetailsCommonTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Network details
type WaapRequestDetailsNetwork struct {
	// Client IP
	ClientIP string `json:"client_ip,required"`
	// Country code
	Country string `json:"country,required"`
	// Organization details
	Organization WaapRequestDetailsNetworkOrganization `json:"organization,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientIP     respjson.Field
		Country      respjson.Field
		Organization respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRequestDetailsNetwork) RawJSON() string { return r.JSON.raw }
func (r *WaapRequestDetailsNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Organization details
type WaapRequestDetailsNetworkOrganization struct {
	// Organization name
	Name string `json:"name,required"`
	// Network range
	Subnet string `json:"subnet,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Subnet      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRequestDetailsNetworkOrganization) RawJSON() string { return r.JSON.raw }
func (r *WaapRequestDetailsNetworkOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pattern matched tag details
type WaapRequestDetailsPatternMatchedTag struct {
	// Tag description information
	Description string `json:"description,required"`
	// The tag's display name
	DisplayName string `json:"display_name,required"`
	// The phase in which the tag was triggered: access -> Request, `header_filter` ->
	// `response_header`, `body_filter` -> `response_body`
	ExecutionPhase string `json:"execution_phase,required"`
	// The entity to which the variable that triggered the tag belong to. For example:
	// `request_headers`, uri, cookies etc.
	Field string `json:"field,required"`
	// The name of the variable which holds the value that triggered the tag
	FieldName string `json:"field_name,required"`
	// The name of the detected regexp pattern
	PatternName string `json:"pattern_name,required"`
	// The pattern which triggered the tag
	PatternValue string `json:"pattern_value,required"`
	// Tag name
	Tag string `json:"tag,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		DisplayName    respjson.Field
		ExecutionPhase respjson.Field
		Field          respjson.Field
		FieldName      respjson.Field
		PatternName    respjson.Field
		PatternValue   respjson.Field
		Tag            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRequestDetailsPatternMatchedTag) RawJSON() string { return r.JSON.raw }
func (r *WaapRequestDetailsPatternMatchedTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The result of a request
type WaapRequestDetailsResult string

const (
	WaapRequestDetailsResultPassed     WaapRequestDetailsResult = "passed"
	WaapRequestDetailsResultBlocked    WaapRequestDetailsResult = "blocked"
	WaapRequestDetailsResultSuppressed WaapRequestDetailsResult = "suppressed"
	WaapRequestDetailsResultEmpty      WaapRequestDetailsResult = ""
)

// User agent
type WaapRequestDetailsUserAgent struct {
	// User agent browser
	BaseBrowser string `json:"base_browser,required"`
	// User agent browser version
	BaseBrowserVersion string `json:"base_browser_version,required"`
	// Client from User agent header
	Client string `json:"client,required"`
	// User agent client type
	ClientType string `json:"client_type,required"`
	// User agent client version
	ClientVersion string `json:"client_version,required"`
	// User agent cpu
	CPU string `json:"cpu,required"`
	// User agent device
	Device string `json:"device,required"`
	// User agent device type
	DeviceType string `json:"device_type,required"`
	// User agent
	FullString string `json:"full_string,required"`
	// User agent os
	Os string `json:"os,required"`
	// User agent engine
	RenderingEngine string `json:"rendering_engine,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseBrowser        respjson.Field
		BaseBrowserVersion respjson.Field
		Client             respjson.Field
		ClientType         respjson.Field
		ClientVersion      respjson.Field
		CPU                respjson.Field
		Device             respjson.Field
		DeviceType         respjson.Field
		FullString         respjson.Field
		Os                 respjson.Field
		RenderingEngine    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRequestDetailsUserAgent) RawJSON() string { return r.JSON.raw }
func (r *WaapRequestDetailsUserAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request summary used when displaying a list of requests
type WaapRequestSummary struct {
	// Request's unique id
	ID string `json:"id,required"`
	// Action of the triggered rule
	Action string `json:"action,required"`
	// Client's IP address.
	ClientIP string `json:"client_ip,required"`
	// Country code
	Country string `json:"country,required"`
	// Domain name
	Domain string `json:"domain,required"`
	// HTTP method
	Method string `json:"method,required"`
	// Organization
	Organization string `json:"organization,required"`
	// Request path
	Path string `json:"path,required"`
	// The reference ID to a sanction that was given to a user.
	ReferenceID string `json:"reference_id,required"`
	// The UNIX timestamp in ms of the date a set of traffic counters was recorded
	RequestTime int64 `json:"request_time,required"`
	// Any of "passed", "blocked", "suppressed", "".
	Result WaapRequestSummaryResult `json:"result,required"`
	// The ID of the triggered rule.
	RuleID string `json:"rule_id,required"`
	// Name of the triggered rule
	RuleName string `json:"rule_name,required"`
	// Status code for http request
	StatusCode int64 `json:"status_code,required"`
	// Comma separated list of traffic types.
	TrafficTypes string `json:"traffic_types,required"`
	// User agent
	UserAgent string `json:"user_agent,required"`
	// Client from parsed User agent header
	UserAgentClient string `json:"user_agent_client,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Action          respjson.Field
		ClientIP        respjson.Field
		Country         respjson.Field
		Domain          respjson.Field
		Method          respjson.Field
		Organization    respjson.Field
		Path            respjson.Field
		ReferenceID     respjson.Field
		RequestTime     respjson.Field
		Result          respjson.Field
		RuleID          respjson.Field
		RuleName        respjson.Field
		StatusCode      respjson.Field
		TrafficTypes    respjson.Field
		UserAgent       respjson.Field
		UserAgentClient respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRequestSummary) RawJSON() string { return r.JSON.raw }
func (r *WaapRequestSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapRequestSummaryResult string

const (
	WaapRequestSummaryResultPassed     WaapRequestSummaryResult = "passed"
	WaapRequestSummaryResultBlocked    WaapRequestSummaryResult = "blocked"
	WaapRequestSummaryResultSuppressed WaapRequestSummaryResult = "suppressed"
	WaapRequestSummaryResultEmpty      WaapRequestSummaryResult = ""
)

// Represents the traffic metrics for a domain at a given time window
type WaapTrafficMetrics struct {
	// UNIX timestamp indicating when the traffic data was recorded
	Timestamp int64 `json:"timestamp,required"`
	// Number of AJAX requests made
	Ajax int64 `json:"ajax"`
	// Number of API requests made
	API int64 `json:"api"`
	// Number of requests allowed through custom rules
	CustomAllowed int64 `json:"customAllowed"`
	// Number of requests blocked due to custom rules
	CustomBlocked int64 `json:"customBlocked"`
	// Number of DDoS attack attempts successfully blocked
	DDOSBlocked int64 `json:"ddosBlocked"`
	// Number of requests triggering monitoring actions
	Monitored int64 `json:"monitored"`
	// Number of successful HTTP 2xx responses from the origin server
	Origin2xx int64 `json:"origin2xx"`
	// Number of HTTP 3xx redirects issued by the origin server
	Origin3xx int64 `json:"origin3xx"`
	// Number of HTTP 4xx errors from the origin server
	OriginError4xx int64 `json:"originError4xx"`
	// Number of HTTP 5xx errors from the origin server
	OriginError5xx int64 `json:"originError5xx"`
	// Number of timeouts experienced at the origin server
	OriginTimeout int64 `json:"originTimeout"`
	// Number of requests served directly by the origin server
	PassedToOrigin int64 `json:"passedToOrigin"`
	// Number of requests allowed by security policies
	PolicyAllowed int64 `json:"policyAllowed"`
	// Number of requests blocked by security policies
	PolicyBlocked int64 `json:"policyBlocked"`
	// Average origin server response time in milliseconds
	ResponseTime int64 `json:"responseTime"`
	// Number of static asset requests
	Static int64 `json:"static"`
	// Total number of requests
	Total int64 `json:"total"`
	// Requests resulting in neither blocks nor sanctions
	Uncategorized int64 `json:"uncategorized"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timestamp      respjson.Field
		Ajax           respjson.Field
		API            respjson.Field
		CustomAllowed  respjson.Field
		CustomBlocked  respjson.Field
		DDOSBlocked    respjson.Field
		Monitored      respjson.Field
		Origin2xx      respjson.Field
		Origin3xx      respjson.Field
		OriginError4xx respjson.Field
		OriginError5xx respjson.Field
		OriginTimeout  respjson.Field
		PassedToOrigin respjson.Field
		PolicyAllowed  respjson.Field
		PolicyBlocked  respjson.Field
		ResponseTime   respjson.Field
		Static         respjson.Field
		Total          respjson.Field
		Uncategorized  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTrafficMetrics) RawJSON() string { return r.JSON.raw }
func (r *WaapTrafficMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainStatisticGetDDOSAttacksParams struct {
	// Filter attacks up to a specified end date in ISO 8601 format
	EndTime param.Opt[time.Time] `query:"end_time,omitzero" format:"date-time" json:"-"`
	// Filter attacks starting from a specified date in ISO 8601 format
	StartTime param.Opt[time.Time] `query:"start_time,omitzero" format:"date-time" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Sort the response by given field.
	//
	// Any of "start_time", "-start_time", "end_time", "-end_time".
	Ordering DomainStatisticGetDDOSAttacksParamsOrdering `query:"ordering,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainStatisticGetDDOSAttacksParams]'s query parameters as
// `url.Values`.
func (r DomainStatisticGetDDOSAttacksParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort the response by given field.
type DomainStatisticGetDDOSAttacksParamsOrdering string

const (
	DomainStatisticGetDDOSAttacksParamsOrderingStartTime      DomainStatisticGetDDOSAttacksParamsOrdering = "start_time"
	DomainStatisticGetDDOSAttacksParamsOrderingMinusStartTime DomainStatisticGetDDOSAttacksParamsOrdering = "-start_time"
	DomainStatisticGetDDOSAttacksParamsOrderingEndTime        DomainStatisticGetDDOSAttacksParamsOrdering = "end_time"
	DomainStatisticGetDDOSAttacksParamsOrderingMinusEndTime   DomainStatisticGetDDOSAttacksParamsOrdering = "-end_time"
)

type DomainStatisticGetDDOSInfoParams struct {
	// The identity of the requests to group by
	//
	// Any of "URL", "User-Agent", "IP".
	GroupBy DomainStatisticGetDDOSInfoParamsGroupBy `query:"group_by,omitzero,required" json:"-"`
	// Filter data items starting from a specified date in ISO 8601 format
	Start string `query:"start,required" json:"-"`
	// Filter data items up to a specified end date in ISO 8601 format. If not
	// provided, defaults to the current date and time.
	End param.Opt[string] `query:"end,omitzero" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainStatisticGetDDOSInfoParams]'s query parameters as
// `url.Values`.
func (r DomainStatisticGetDDOSInfoParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The identity of the requests to group by
type DomainStatisticGetDDOSInfoParamsGroupBy string

const (
	DomainStatisticGetDDOSInfoParamsGroupByURL       DomainStatisticGetDDOSInfoParamsGroupBy = "URL"
	DomainStatisticGetDDOSInfoParamsGroupByUserAgent DomainStatisticGetDDOSInfoParamsGroupBy = "User-Agent"
	DomainStatisticGetDDOSInfoParamsGroupByIP        DomainStatisticGetDDOSInfoParamsGroupBy = "IP"
)

type DomainStatisticGetEventsAggregatedParams struct {
	// Filter data items starting from a specified date in ISO 8601 format
	Start string `query:"start,required" json:"-"`
	// Filter data items up to a specified end date in ISO 8601 format. If not
	// provided, defaults to the current date and time.
	End param.Opt[string] `query:"end,omitzero" json:"-"`
	// A list of action names to filter on.
	//
	// Any of "block", "captcha", "handshake", "monitor".
	Action []string `query:"action,omitzero" json:"-"`
	// A list of IPs to filter event statistics.
	IP []string `query:"ip,omitzero" format:"ipvanyaddress" json:"-"`
	// A list of reference IDs to filter event statistics.
	ReferenceID []string `query:"reference_id,omitzero" json:"-"`
	// A list of results to filter event statistics.
	//
	// Any of "passed", "blocked", "monitored", "allowed".
	Result []string `query:"result,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainStatisticGetEventsAggregatedParams]'s query
// parameters as `url.Values`.
func (r DomainStatisticGetEventsAggregatedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DomainStatisticGetRequestDetailsParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}

type DomainStatisticGetRequestsSeriesParams struct {
	// Filter data items starting from a specified date in ISO 8601 format
	Start string `query:"start,required" json:"-"`
	// Filter data items up to a specified end date in ISO 8601 format. If not
	// provided, defaults to the current date and time.
	End param.Opt[string] `query:"end,omitzero" json:"-"`
	// Filter the response by IP.
	IP param.Opt[string] `query:"ip,omitzero" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Sort the response by given field.
	Ordering param.Opt[string] `query:"ordering,omitzero" json:"-"`
	// Filter the response by reference ID.
	ReferenceID param.Opt[string] `query:"reference_id,omitzero" json:"-"`
	// Filter the response by security rule name.
	SecurityRuleName param.Opt[string] `query:"security_rule_name,omitzero" json:"-"`
	// Filter the response by response code.
	StatusCode param.Opt[int64] `query:"status_code,omitzero" json:"-"`
	// Filter the response by actions.
	//
	// Any of "allow", "block", "captcha", "handshake".
	Actions []string `query:"actions,omitzero" json:"-"`
	// Filter the response by country codes in ISO 3166-1 alpha-2 format.
	Countries []string `query:"countries,omitzero" json:"-"`
	// Filter the response by traffic types.
	//
	// Any of "policy_allowed", "policy_blocked", "custom_rule_allowed",
	// "custom_blocked", "legit_requests", "sanctioned", "dynamic", "api", "static",
	// "ajax", "redirects", "monitor", "err_40x", "err_50x", "passed_to_origin",
	// "timeout", "other", "ddos", "legit", "monitored".
	TrafficTypes []string `query:"traffic_types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainStatisticGetRequestsSeriesParams]'s query parameters
// as `url.Values`.
func (r DomainStatisticGetRequestsSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DomainStatisticGetTrafficSeriesParams struct {
	// Specifies the granularity of the result data.
	//
	// Any of "daily", "hourly", "minutely".
	Resolution DomainStatisticGetTrafficSeriesParamsResolution `query:"resolution,omitzero,required" json:"-"`
	// Filter data items starting from a specified date in ISO 8601 format
	Start string `query:"start,required" json:"-"`
	// Filter data items up to a specified end date in ISO 8601 format. If not
	// provided, defaults to the current date and time.
	End param.Opt[string] `query:"end,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainStatisticGetTrafficSeriesParams]'s query parameters
// as `url.Values`.
func (r DomainStatisticGetTrafficSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the granularity of the result data.
type DomainStatisticGetTrafficSeriesParamsResolution string

const (
	DomainStatisticGetTrafficSeriesParamsResolutionDaily    DomainStatisticGetTrafficSeriesParamsResolution = "daily"
	DomainStatisticGetTrafficSeriesParamsResolutionHourly   DomainStatisticGetTrafficSeriesParamsResolution = "hourly"
	DomainStatisticGetTrafficSeriesParamsResolutionMinutely DomainStatisticGetTrafficSeriesParamsResolution = "minutely"
)
