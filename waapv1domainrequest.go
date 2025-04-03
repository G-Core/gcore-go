// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// WaapV1DomainRequestService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1DomainRequestService] method instead.
type WaapV1DomainRequestService struct {
	Options []option.RequestOption
}

// NewWaapV1DomainRequestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWaapV1DomainRequestService(opts ...option.RequestOption) (r *WaapV1DomainRequestService) {
	r = &WaapV1DomainRequestService{}
	r.Options = opts
	return
}

// Retrieve a domain's requests data.
func (r *WaapV1DomainRequestService) List(ctx context.Context, domainID int64, query WaapV1DomainRequestListParams, opts ...option.RequestOption) (res *WaapV1DomainRequestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/requests", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves all the available information for a request that matches a given
// request id
func (r *WaapV1DomainRequestService) GetDetails(ctx context.Context, domainID int64, requestID string, opts ...option.RequestOption) (res *WaapV1DomainRequestGetDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/requests/%s/details", domainID, requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type WaapV1DomainRequestListResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []WaapV1DomainRequestListResponseResult `json:"results,required"`
	JSON    waapV1DomainRequestListResponseJSON     `json:"-"`
}

// waapV1DomainRequestListResponseJSON contains the JSON metadata for the struct
// [WaapV1DomainRequestListResponse]
type waapV1DomainRequestListResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainRequestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainRequestListResponseJSON) RawJSON() string {
	return r.raw
}

// Request summary used when displaying a list of requests
type WaapV1DomainRequestListResponseResult struct {
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
	RequestTime int64                                        `json:"request_time,required"`
	Result      WaapV1DomainRequestListResponseResultsResult `json:"result,required"`
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
	UserAgentClient string                                    `json:"user_agent_client,required"`
	JSON            waapV1DomainRequestListResponseResultJSON `json:"-"`
}

// waapV1DomainRequestListResponseResultJSON contains the JSON metadata for the
// struct [WaapV1DomainRequestListResponseResult]
type waapV1DomainRequestListResponseResultJSON struct {
	ID              apijson.Field
	Action          apijson.Field
	ClientIP        apijson.Field
	Country         apijson.Field
	Domain          apijson.Field
	Method          apijson.Field
	Organization    apijson.Field
	Path            apijson.Field
	ReferenceID     apijson.Field
	RequestTime     apijson.Field
	Result          apijson.Field
	RuleID          apijson.Field
	RuleName        apijson.Field
	StatusCode      apijson.Field
	TrafficTypes    apijson.Field
	UserAgent       apijson.Field
	UserAgentClient apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WaapV1DomainRequestListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainRequestListResponseResultJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainRequestListResponseResultsResult string

const (
	WaapV1DomainRequestListResponseResultsResultPassed     WaapV1DomainRequestListResponseResultsResult = "passed"
	WaapV1DomainRequestListResponseResultsResultBlocked    WaapV1DomainRequestListResponseResultsResult = "blocked"
	WaapV1DomainRequestListResponseResultsResultSuppressed WaapV1DomainRequestListResponseResultsResult = "suppressed"
	WaapV1DomainRequestListResponseResultsResultEmpty      WaapV1DomainRequestListResponseResultsResult = ""
)

func (r WaapV1DomainRequestListResponseResultsResult) IsKnown() bool {
	switch r {
	case WaapV1DomainRequestListResponseResultsResultPassed, WaapV1DomainRequestListResponseResultsResultBlocked, WaapV1DomainRequestListResponseResultsResultSuppressed, WaapV1DomainRequestListResponseResultsResultEmpty:
		return true
	}
	return false
}

// Request's details used when displaying a single request.
type WaapV1DomainRequestGetDetailsResponse struct {
	// Request ID
	ID string `json:"id,required"`
	// Request action
	Action string `json:"action,required"`
	// List of common tags
	CommonTags []WaapV1DomainRequestGetDetailsResponseCommonTag `json:"common_tags,required"`
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
	Network WaapV1DomainRequestGetDetailsResponseNetwork `json:"network,required"`
	// Request path
	Path string `json:"path,required"`
	// List of shield tags
	PatternMatchedTags []WaapV1DomainRequestGetDetailsResponsePatternMatchedTag `json:"pattern_matched_tags,required"`
	// The query string of the request
	QueryString string `json:"query_string,required"`
	// Reference ID to identify user sanction
	ReferenceID string `json:"reference_id,required"`
	// HTTP request headers
	RequestHeaders interface{} `json:"request_headers,required"`
	// The time of the request
	RequestTime string `json:"request_time,required"`
	// The type of the request that generated an event
	RequestType string `json:"request_type,required"`
	// The real domain name
	RequestedDomain string `json:"requested_domain,required"`
	// Time took to process all request
	ResponseTime string `json:"response_time,required"`
	// The result of a request
	Result WaapV1DomainRequestGetDetailsResponseResult `json:"result,required"`
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
	// User agent details
	UserAgent WaapV1DomainRequestGetDetailsResponseUserAgent `json:"user_agent,required"`
	JSON      waapV1DomainRequestGetDetailsResponseJSON      `json:"-"`
}

// waapV1DomainRequestGetDetailsResponseJSON contains the JSON metadata for the
// struct [WaapV1DomainRequestGetDetailsResponse]
type waapV1DomainRequestGetDetailsResponseJSON struct {
	ID                  apijson.Field
	Action              apijson.Field
	CommonTags          apijson.Field
	ContentType         apijson.Field
	Domain              apijson.Field
	HTTPStatusCode      apijson.Field
	HTTPVersion         apijson.Field
	IncidentID          apijson.Field
	Method              apijson.Field
	Network             apijson.Field
	Path                apijson.Field
	PatternMatchedTags  apijson.Field
	QueryString         apijson.Field
	ReferenceID         apijson.Field
	RequestHeaders      apijson.Field
	RequestTime         apijson.Field
	RequestType         apijson.Field
	RequestedDomain     apijson.Field
	ResponseTime        apijson.Field
	Result              apijson.Field
	RuleID              apijson.Field
	RuleName            apijson.Field
	Scheme              apijson.Field
	SessionRequestCount apijson.Field
	TrafficTypes        apijson.Field
	UserAgent           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *WaapV1DomainRequestGetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainRequestGetDetailsResponseJSON) RawJSON() string {
	return r.raw
}

// Common tag details
type WaapV1DomainRequestGetDetailsResponseCommonTag struct {
	// Tag description information
	Description string `json:"description,required"`
	// The tag's display name
	DisplayName string `json:"display_name,required"`
	// Tag name
	Tag  string                                             `json:"tag,required"`
	JSON waapV1DomainRequestGetDetailsResponseCommonTagJSON `json:"-"`
}

// waapV1DomainRequestGetDetailsResponseCommonTagJSON contains the JSON metadata
// for the struct [WaapV1DomainRequestGetDetailsResponseCommonTag]
type waapV1DomainRequestGetDetailsResponseCommonTagJSON struct {
	Description apijson.Field
	DisplayName apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainRequestGetDetailsResponseCommonTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainRequestGetDetailsResponseCommonTagJSON) RawJSON() string {
	return r.raw
}

// Network details
type WaapV1DomainRequestGetDetailsResponseNetwork struct {
	// Client IP
	ClientIP string `json:"client_ip,required"`
	// Country code
	Country string `json:"country,required"`
	// Organization details
	Organization WaapV1DomainRequestGetDetailsResponseNetworkOrganization `json:"organization,required"`
	JSON         waapV1DomainRequestGetDetailsResponseNetworkJSON         `json:"-"`
}

// waapV1DomainRequestGetDetailsResponseNetworkJSON contains the JSON metadata for
// the struct [WaapV1DomainRequestGetDetailsResponseNetwork]
type waapV1DomainRequestGetDetailsResponseNetworkJSON struct {
	ClientIP     apijson.Field
	Country      apijson.Field
	Organization apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WaapV1DomainRequestGetDetailsResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainRequestGetDetailsResponseNetworkJSON) RawJSON() string {
	return r.raw
}

// Organization details
type WaapV1DomainRequestGetDetailsResponseNetworkOrganization struct {
	// Organization name
	Name string `json:"name,required"`
	// Network range
	Subnet string                                                       `json:"subnet,required"`
	JSON   waapV1DomainRequestGetDetailsResponseNetworkOrganizationJSON `json:"-"`
}

// waapV1DomainRequestGetDetailsResponseNetworkOrganizationJSON contains the JSON
// metadata for the struct
// [WaapV1DomainRequestGetDetailsResponseNetworkOrganization]
type waapV1DomainRequestGetDetailsResponseNetworkOrganizationJSON struct {
	Name        apijson.Field
	Subnet      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainRequestGetDetailsResponseNetworkOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainRequestGetDetailsResponseNetworkOrganizationJSON) RawJSON() string {
	return r.raw
}

// Pattern matched tag details
type WaapV1DomainRequestGetDetailsResponsePatternMatchedTag struct {
	// Tag description information
	Description string `json:"description,required"`
	// The tag's display name
	DisplayName string `json:"display_name,required"`
	// The phase in which the tag was triggered: access -> Request, header_filter ->
	// response_header, body_filter -> response_body
	ExecutionPhase string `json:"execution_phase,required"`
	// The entity to which the variable that triggered the tag belong to. For example:
	// request_headers, uri, cookies etc.
	Field string `json:"field,required"`
	// The name of the variable which holds the value that triggered the tag
	FieldName string `json:"field_name,required"`
	// The name of the detected regexp pattern
	PatternName string `json:"pattern_name,required"`
	// The pattern which triggered the tag
	PatternValue string `json:"pattern_value,required"`
	// Tag name
	Tag  string                                                     `json:"tag,required"`
	JSON waapV1DomainRequestGetDetailsResponsePatternMatchedTagJSON `json:"-"`
}

// waapV1DomainRequestGetDetailsResponsePatternMatchedTagJSON contains the JSON
// metadata for the struct [WaapV1DomainRequestGetDetailsResponsePatternMatchedTag]
type waapV1DomainRequestGetDetailsResponsePatternMatchedTagJSON struct {
	Description    apijson.Field
	DisplayName    apijson.Field
	ExecutionPhase apijson.Field
	Field          apijson.Field
	FieldName      apijson.Field
	PatternName    apijson.Field
	PatternValue   apijson.Field
	Tag            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WaapV1DomainRequestGetDetailsResponsePatternMatchedTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainRequestGetDetailsResponsePatternMatchedTagJSON) RawJSON() string {
	return r.raw
}

// The result of a request
type WaapV1DomainRequestGetDetailsResponseResult string

const (
	WaapV1DomainRequestGetDetailsResponseResultPassed     WaapV1DomainRequestGetDetailsResponseResult = "passed"
	WaapV1DomainRequestGetDetailsResponseResultBlocked    WaapV1DomainRequestGetDetailsResponseResult = "blocked"
	WaapV1DomainRequestGetDetailsResponseResultSuppressed WaapV1DomainRequestGetDetailsResponseResult = "suppressed"
	WaapV1DomainRequestGetDetailsResponseResultEmpty      WaapV1DomainRequestGetDetailsResponseResult = ""
)

func (r WaapV1DomainRequestGetDetailsResponseResult) IsKnown() bool {
	switch r {
	case WaapV1DomainRequestGetDetailsResponseResultPassed, WaapV1DomainRequestGetDetailsResponseResultBlocked, WaapV1DomainRequestGetDetailsResponseResultSuppressed, WaapV1DomainRequestGetDetailsResponseResultEmpty:
		return true
	}
	return false
}

// User agent details
type WaapV1DomainRequestGetDetailsResponseUserAgent struct {
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
	RenderingEngine string                                             `json:"rendering_engine,required"`
	JSON            waapV1DomainRequestGetDetailsResponseUserAgentJSON `json:"-"`
}

// waapV1DomainRequestGetDetailsResponseUserAgentJSON contains the JSON metadata
// for the struct [WaapV1DomainRequestGetDetailsResponseUserAgent]
type waapV1DomainRequestGetDetailsResponseUserAgentJSON struct {
	BaseBrowser        apijson.Field
	BaseBrowserVersion apijson.Field
	Client             apijson.Field
	ClientType         apijson.Field
	ClientVersion      apijson.Field
	CPU                apijson.Field
	Device             apijson.Field
	DeviceType         apijson.Field
	FullString         apijson.Field
	Os                 apijson.Field
	RenderingEngine    apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WaapV1DomainRequestGetDetailsResponseUserAgent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainRequestGetDetailsResponseUserAgentJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainRequestListParams struct {
	// Filter traffic starting from a specified date in ISO 8601 format
	Start param.Field[time.Time] `query:"start,required" format:"date-time"`
	// Filter the response by actions.
	Actions param.Field[[]WaapV1DomainRequestListParamsAction] `query:"actions"`
	// Filter the response by country codes in ISO 3166-1 alpha-2 format.
	Countries param.Field[[]string] `query:"countries"`
	// Filter traffic up to a specified end date in ISO 8601 format. If not provided,
	// defaults to the current date and time.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// Filter the response by IP.
	IP param.Field[string] `query:"ip"`
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip
	Offset param.Field[int64] `query:"offset"`
	// Sort the response by given field.
	Ordering param.Field[string] `query:"ordering"`
	// Filter the response by reference ID.
	ReferenceID param.Field[string] `query:"reference_id"`
	// Filter the response by security rule name.
	SecurityRuleName param.Field[string] `query:"security_rule_name"`
	// Filter the response by response code.
	StatusCode param.Field[int64] `query:"status_code"`
	// Filter the response by traffic types.
	TrafficTypes param.Field[[]WaapV1DomainRequestListParamsTrafficType] `query:"traffic_types"`
}

// URLQuery serializes [WaapV1DomainRequestListParams]'s query parameters as
// `url.Values`.
func (r WaapV1DomainRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WaapV1DomainRequestListParamsAction string

const (
	WaapV1DomainRequestListParamsActionAllow     WaapV1DomainRequestListParamsAction = "allow"
	WaapV1DomainRequestListParamsActionBlock     WaapV1DomainRequestListParamsAction = "block"
	WaapV1DomainRequestListParamsActionCaptcha   WaapV1DomainRequestListParamsAction = "captcha"
	WaapV1DomainRequestListParamsActionHandshake WaapV1DomainRequestListParamsAction = "handshake"
)

func (r WaapV1DomainRequestListParamsAction) IsKnown() bool {
	switch r {
	case WaapV1DomainRequestListParamsActionAllow, WaapV1DomainRequestListParamsActionBlock, WaapV1DomainRequestListParamsActionCaptcha, WaapV1DomainRequestListParamsActionHandshake:
		return true
	}
	return false
}

type WaapV1DomainRequestListParamsTrafficType string

const (
	WaapV1DomainRequestListParamsTrafficTypePolicyAllowed     WaapV1DomainRequestListParamsTrafficType = "policy_allowed"
	WaapV1DomainRequestListParamsTrafficTypePolicyBlocked     WaapV1DomainRequestListParamsTrafficType = "policy_blocked"
	WaapV1DomainRequestListParamsTrafficTypeCustomRuleAllowed WaapV1DomainRequestListParamsTrafficType = "custom_rule_allowed"
	WaapV1DomainRequestListParamsTrafficTypeCustomBlocked     WaapV1DomainRequestListParamsTrafficType = "custom_blocked"
	WaapV1DomainRequestListParamsTrafficTypeLegitRequests     WaapV1DomainRequestListParamsTrafficType = "legit_requests"
	WaapV1DomainRequestListParamsTrafficTypeSanctioned        WaapV1DomainRequestListParamsTrafficType = "sanctioned"
	WaapV1DomainRequestListParamsTrafficTypeDynamic           WaapV1DomainRequestListParamsTrafficType = "dynamic"
	WaapV1DomainRequestListParamsTrafficTypeAPI               WaapV1DomainRequestListParamsTrafficType = "api"
	WaapV1DomainRequestListParamsTrafficTypeStatic            WaapV1DomainRequestListParamsTrafficType = "static"
	WaapV1DomainRequestListParamsTrafficTypeAjax              WaapV1DomainRequestListParamsTrafficType = "ajax"
	WaapV1DomainRequestListParamsTrafficTypeRedirects         WaapV1DomainRequestListParamsTrafficType = "redirects"
	WaapV1DomainRequestListParamsTrafficTypeMonitor           WaapV1DomainRequestListParamsTrafficType = "monitor"
	WaapV1DomainRequestListParamsTrafficTypeErr40x            WaapV1DomainRequestListParamsTrafficType = "err_40x"
	WaapV1DomainRequestListParamsTrafficTypeErr50x            WaapV1DomainRequestListParamsTrafficType = "err_50x"
	WaapV1DomainRequestListParamsTrafficTypePassedToOrigin    WaapV1DomainRequestListParamsTrafficType = "passed_to_origin"
	WaapV1DomainRequestListParamsTrafficTypeTimeout           WaapV1DomainRequestListParamsTrafficType = "timeout"
	WaapV1DomainRequestListParamsTrafficTypeOther             WaapV1DomainRequestListParamsTrafficType = "other"
	WaapV1DomainRequestListParamsTrafficTypeDdos              WaapV1DomainRequestListParamsTrafficType = "ddos"
	WaapV1DomainRequestListParamsTrafficTypeLegit             WaapV1DomainRequestListParamsTrafficType = "legit"
	WaapV1DomainRequestListParamsTrafficTypeMonitored         WaapV1DomainRequestListParamsTrafficType = "monitored"
)

func (r WaapV1DomainRequestListParamsTrafficType) IsKnown() bool {
	switch r {
	case WaapV1DomainRequestListParamsTrafficTypePolicyAllowed, WaapV1DomainRequestListParamsTrafficTypePolicyBlocked, WaapV1DomainRequestListParamsTrafficTypeCustomRuleAllowed, WaapV1DomainRequestListParamsTrafficTypeCustomBlocked, WaapV1DomainRequestListParamsTrafficTypeLegitRequests, WaapV1DomainRequestListParamsTrafficTypeSanctioned, WaapV1DomainRequestListParamsTrafficTypeDynamic, WaapV1DomainRequestListParamsTrafficTypeAPI, WaapV1DomainRequestListParamsTrafficTypeStatic, WaapV1DomainRequestListParamsTrafficTypeAjax, WaapV1DomainRequestListParamsTrafficTypeRedirects, WaapV1DomainRequestListParamsTrafficTypeMonitor, WaapV1DomainRequestListParamsTrafficTypeErr40x, WaapV1DomainRequestListParamsTrafficTypeErr50x, WaapV1DomainRequestListParamsTrafficTypePassedToOrigin, WaapV1DomainRequestListParamsTrafficTypeTimeout, WaapV1DomainRequestListParamsTrafficTypeOther, WaapV1DomainRequestListParamsTrafficTypeDdos, WaapV1DomainRequestListParamsTrafficTypeLegit, WaapV1DomainRequestListParamsTrafficTypeMonitored:
		return true
	}
	return false
}
