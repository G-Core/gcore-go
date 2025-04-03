// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// WaapV1DomainCustomRuleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1DomainCustomRuleService] method instead.
type WaapV1DomainCustomRuleService struct {
	Options []option.RequestOption
}

// NewWaapV1DomainCustomRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWaapV1DomainCustomRuleService(opts ...option.RequestOption) (r *WaapV1DomainCustomRuleService) {
	r = &WaapV1DomainCustomRuleService{}
	r.Options = opts
	return
}

// Create a custom rule
func (r *WaapV1DomainCustomRuleService) New(ctx context.Context, domainID int64, body WaapV1DomainCustomRuleNewParams, opts ...option.RequestOption) (res *CustomRule, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/custom-rules", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Extracts a specific custom rule assigned to a domain
func (r *WaapV1DomainCustomRuleService) Get(ctx context.Context, domainID int64, ruleID int64, opts ...option.RequestOption) (res *CustomRule, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/custom-rules/%v", domainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Only properties present in the request will be updated
func (r *WaapV1DomainCustomRuleService) Update(ctx context.Context, domainID int64, ruleID int64, body WaapV1DomainCustomRuleUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/custom-rules/%v", domainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Extracts a list of custom rules assigned to a domain, offering filter, ordering,
// and pagination capabilities
func (r *WaapV1DomainCustomRuleService) List(ctx context.Context, domainID int64, query WaapV1DomainCustomRuleListParams, opts ...option.RequestOption) (res *WaapV1DomainCustomRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/custom-rules", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a custom rule
func (r *WaapV1DomainCustomRuleService) Delete(ctx context.Context, domainID int64, ruleID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/custom-rules/%v", domainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Delete multiple WAAP rules
func (r *WaapV1DomainCustomRuleService) BulkDelete(ctx context.Context, domainID int64, body WaapV1DomainCustomRuleBulkDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/custom-rules/bulk_delete", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Toggle a custom rule
func (r *WaapV1DomainCustomRuleService) Toggle(ctx context.Context, domainID int64, ruleID int64, action CustomerRuleState, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/custom-rules/%v/%v", domainID, ruleID, action)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, nil, opts...)
	return
}

// Match the requested Content-Type
type ContentTypeCondition struct {
	// The list of content types to match against
	ContentType []string `json:"content_type,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool                     `json:"negation"`
	JSON     contentTypeConditionJSON `json:"-"`
}

// contentTypeConditionJSON contains the JSON metadata for the struct
// [ContentTypeCondition]
type contentTypeConditionJSON struct {
	ContentType apijson.Field
	Negation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentTypeCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentTypeConditionJSON) RawJSON() string {
	return r.raw
}

// Match the requested Content-Type
type ContentTypeConditionParam struct {
	// The list of content types to match against
	ContentType param.Field[[]string] `json:"content_type,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r ContentTypeConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match the country that the request originated from
type CountryCondition struct {
	// A list of ISO 3166-1 alpha-2 formatted strings representing the countries to
	// match against
	CountryCode []string `json:"country_code,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool                 `json:"negation"`
	JSON     countryConditionJSON `json:"-"`
}

// countryConditionJSON contains the JSON metadata for the struct
// [CountryCondition]
type countryConditionJSON struct {
	CountryCode apijson.Field
	Negation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CountryCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r countryConditionJSON) RawJSON() string {
	return r.raw
}

// Match the country that the request originated from
type CountryConditionParam struct {
	// A list of ISO 3166-1 alpha-2 formatted strings representing the countries to
	// match against
	CountryCode param.Field[[]string] `json:"country_code,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r CountryConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An WAAP rule applied to a domain
type CustomRule struct {
	// The unique identifier for the rule
	ID int64 `json:"id,required"`
	// The action that a WAAP rule takes when triggered
	Action CustomerRuleActionOutput `json:"action,required"`
	// The conditions required for the WAAP engine to trigger the rule. Rules may have
	// between 1 and 5 conditions. All conditions must pass for the rule to trigger
	Conditions []CustomRuleCondition `json:"conditions,required"`
	// Whether or not the rule is enabled
	Enabled bool `json:"enabled,required"`
	// The name assigned to the rule
	Name string `json:"name,required"`
	// The description assigned to the rule
	Description string         `json:"description,nullable"`
	JSON        customRuleJSON `json:"-"`
}

// customRuleJSON contains the JSON metadata for the struct [CustomRule]
type customRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Conditions  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customRuleJSON) RawJSON() string {
	return r.raw
}

// The criteria of an incoming web request and the models of the various values
// those criteria can take
type CustomRuleCondition struct {
	// Match the requested Content-Type
	ContentType ContentTypeCondition `json:"content_type,nullable"`
	// Match the country that the request originated from
	Country CountryCondition `json:"country,nullable"`
	// Match the incoming file extension
	FileExtension FileExtensionCondition `json:"file_extension,nullable"`
	// Match an incoming request header
	Header HeaderCondition `json:"header,nullable"`
	// Match when an incoming request header is present
	HeaderExists HeaderExistsCondition `json:"header_exists,nullable"`
	// Match the incoming HTTP method
	HTTPMethod HTTPMethodCondition `json:"http_method,nullable"`
	// Match the incoming request against a single IP address
	IP IPCondition `json:"ip,nullable"`
	// Match the incoming request against an IP range
	IPRange IPRangeCondition `json:"ip_range,nullable"`
	// Match the organization the request originated from, as determined by a WHOIS
	// lookup of the requesting IP
	Organization OrganizationCondition `json:"organization,nullable"`
	// Match the type of organization that owns the IP address making an incoming
	// request
	OwnerTypes OwnerTypesCondition `json:"owner_types,nullable"`
	// Match the rate at which requests come in that match certain conditions
	RequestRate RequestRateCondition `json:"request_rate,nullable"`
	// Match a response header
	ResponseHeader ResponseHeaderCondition `json:"response_header,nullable"`
	// Match when a response header is present
	ResponseHeaderExists ResponseHeaderExistsCondition `json:"response_header_exists,nullable"`
	// Match the number of dynamic page requests made in a WAAP session
	SessionRequestCount SessionRequestCountCondition `json:"session_request_count,nullable"`
	// Matches requests based on specified tags
	Tags TagsCondition `json:"tags,nullable"`
	// Match the incoming request URL
	URL URLCondition `json:"url,nullable"`
	// Match the user agent making the request
	UserAgent UserAgentCondition `json:"user_agent,nullable"`
	// Matches requests based on user-defined tags
	UserDefinedTags UserDefinedTagsCondition `json:"user_defined_tags,nullable"`
	JSON            customRuleConditionJSON  `json:"-"`
}

// customRuleConditionJSON contains the JSON metadata for the struct
// [CustomRuleCondition]
type customRuleConditionJSON struct {
	ContentType          apijson.Field
	Country              apijson.Field
	FileExtension        apijson.Field
	Header               apijson.Field
	HeaderExists         apijson.Field
	HTTPMethod           apijson.Field
	IP                   apijson.Field
	IPRange              apijson.Field
	Organization         apijson.Field
	OwnerTypes           apijson.Field
	RequestRate          apijson.Field
	ResponseHeader       apijson.Field
	ResponseHeaderExists apijson.Field
	SessionRequestCount  apijson.Field
	Tags                 apijson.Field
	URL                  apijson.Field
	UserAgent            apijson.Field
	UserDefinedTags      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CustomRuleCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customRuleConditionJSON) RawJSON() string {
	return r.raw
}

// The criteria of an incoming web request and the models of the various values
// those criteria can take
type CustomRuleConditionInputParam struct {
	// Match the requested Content-Type
	ContentType param.Field[ContentTypeConditionParam] `json:"content_type"`
	// Match the country that the request originated from
	Country param.Field[CountryConditionParam] `json:"country"`
	// Match the incoming file extension
	FileExtension param.Field[FileExtensionConditionParam] `json:"file_extension"`
	// Match an incoming request header
	Header param.Field[HeaderConditionParam] `json:"header"`
	// Match when an incoming request header is present
	HeaderExists param.Field[HeaderExistsConditionParam] `json:"header_exists"`
	// Match the incoming HTTP method
	HTTPMethod param.Field[HTTPMethodConditionParam] `json:"http_method"`
	// Match the incoming request against a single IP address
	IP param.Field[IPConditionParam] `json:"ip"`
	// Match the incoming request against an IP range
	IPRange param.Field[IPRangeConditionParam] `json:"ip_range"`
	// Match the organization the request originated from, as determined by a WHOIS
	// lookup of the requesting IP
	Organization param.Field[OrganizationConditionParam] `json:"organization"`
	// Match the type of organization that owns the IP address making an incoming
	// request
	OwnerTypes param.Field[OwnerTypesConditionParam] `json:"owner_types"`
	// Match the rate at which requests come in that match certain conditions
	RequestRate param.Field[RequestRateConditionParam] `json:"request_rate"`
	// Match a response header
	ResponseHeader param.Field[ResponseHeaderConditionParam] `json:"response_header"`
	// Match when a response header is present
	ResponseHeaderExists param.Field[ResponseHeaderExistsConditionParam] `json:"response_header_exists"`
	// Match the number of dynamic page requests made in a WAAP session
	SessionRequestCount param.Field[SessionRequestCountConditionParam] `json:"session_request_count"`
	// Matches requests based on specified tags
	Tags param.Field[TagsConditionParam] `json:"tags"`
	// Match the incoming request URL
	URL param.Field[URLConditionParam] `json:"url"`
	// Match the user agent making the request
	UserAgent param.Field[UserAgentConditionParam] `json:"user_agent"`
	// Matches requests based on user-defined tags
	UserDefinedTags param.Field[UserDefinedTagsConditionParam] `json:"user_defined_tags"`
}

func (r CustomRuleConditionInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action that a WAAP rule takes when triggered
type CustomerRuleActionInputParam struct {
	// The WAAP allowed the request
	Allow param.Field[RuleAllowActionParam] `json:"allow"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block param.Field[RuleBlockActionParam] `json:"block"`
	// The WAAP presented the user with a captcha
	Captcha param.Field[RuleCaptchaActionParam] `json:"captcha"`
	// The WAAP performed automatic browser validation
	Handshake param.Field[RuleHandshakeActionParam] `json:"handshake"`
	// The WAAP monitored the request but took no action
	Monitor param.Field[RuleMonitorActionParam] `json:"monitor"`
	// WAAP tag action gets a list of tags to tag the request scope with
	Tag param.Field[RuleTagActionParam] `json:"tag"`
}

func (r CustomerRuleActionInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action that a WAAP rule takes when triggered
type CustomerRuleActionOutput struct {
	// The WAAP allowed the request
	Allow RuleAllowAction `json:"allow,nullable"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block RuleBlockAction `json:"block,nullable"`
	// The WAAP presented the user with a captcha
	Captcha RuleCaptchaAction `json:"captcha,nullable"`
	// The WAAP performed automatic browser validation
	Handshake RuleHandshakeAction `json:"handshake,nullable"`
	// The WAAP monitored the request but took no action
	Monitor RuleMonitorAction `json:"monitor,nullable"`
	// WAAP tag action gets a list of tags to tag the request scope with
	Tag  RuleTagAction                `json:"tag,nullable"`
	JSON customerRuleActionOutputJSON `json:"-"`
}

// customerRuleActionOutputJSON contains the JSON metadata for the struct
// [CustomerRuleActionOutput]
type customerRuleActionOutputJSON struct {
	Allow       apijson.Field
	Block       apijson.Field
	Captcha     apijson.Field
	Handshake   apijson.Field
	Monitor     apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerRuleActionOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerRuleActionOutputJSON) RawJSON() string {
	return r.raw
}

type CustomerRuleState string

const (
	CustomerRuleStateEnable  CustomerRuleState = "enable"
	CustomerRuleStateDisable CustomerRuleState = "disable"
)

func (r CustomerRuleState) IsKnown() bool {
	switch r {
	case CustomerRuleStateEnable, CustomerRuleStateDisable:
		return true
	}
	return false
}

// Match the incoming file extension
type FileExtensionCondition struct {
	// The list of file extensions to match against
	FileExtension []string `json:"file_extension,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool                       `json:"negation"`
	JSON     fileExtensionConditionJSON `json:"-"`
}

// fileExtensionConditionJSON contains the JSON metadata for the struct
// [FileExtensionCondition]
type fileExtensionConditionJSON struct {
	FileExtension apijson.Field
	Negation      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FileExtensionCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileExtensionConditionJSON) RawJSON() string {
	return r.raw
}

// Match the incoming file extension
type FileExtensionConditionParam struct {
	// The list of file extensions to match against
	FileExtension param.Field[[]string] `json:"file_extension,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r FileExtensionConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an incoming request header
type HeaderCondition struct {
	// The request header name
	Header string `json:"header,required"`
	// The request header value
	Value string `json:"value,required"`
	// The type of matching condition for header and value.
	MatchType HeaderConditionMatchType `json:"match_type"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool                `json:"negation"`
	JSON     headerConditionJSON `json:"-"`
}

// headerConditionJSON contains the JSON metadata for the struct [HeaderCondition]
type headerConditionJSON struct {
	Header      apijson.Field
	Value       apijson.Field
	MatchType   apijson.Field
	Negation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HeaderCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r headerConditionJSON) RawJSON() string {
	return r.raw
}

// The type of matching condition for header and value.
type HeaderConditionMatchType string

const (
	HeaderConditionMatchTypeExact    HeaderConditionMatchType = "Exact"
	HeaderConditionMatchTypeContains HeaderConditionMatchType = "Contains"
)

func (r HeaderConditionMatchType) IsKnown() bool {
	switch r {
	case HeaderConditionMatchTypeExact, HeaderConditionMatchTypeContains:
		return true
	}
	return false
}

// Match an incoming request header
type HeaderConditionParam struct {
	// The request header name
	Header param.Field[string] `json:"header,required"`
	// The request header value
	Value param.Field[string] `json:"value,required"`
	// The type of matching condition for header and value.
	MatchType param.Field[HeaderConditionMatchType] `json:"match_type"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r HeaderConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match when an incoming request header is present
type HeaderExistsCondition struct {
	// The request header name
	Header string `json:"header,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool                      `json:"negation"`
	JSON     headerExistsConditionJSON `json:"-"`
}

// headerExistsConditionJSON contains the JSON metadata for the struct
// [HeaderExistsCondition]
type headerExistsConditionJSON struct {
	Header      apijson.Field
	Negation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HeaderExistsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r headerExistsConditionJSON) RawJSON() string {
	return r.raw
}

// Match when an incoming request header is present
type HeaderExistsConditionParam struct {
	// The request header name
	Header param.Field[string] `json:"header,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r HeaderExistsConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match the incoming HTTP method
type HTTPMethodCondition struct {
	// HTTP methods and descriptions
	//
	// Methods from the following RFCs are all observed:
	//
	//   - RFC 7231: Hypertext Transfer Protocol (HTTP/1.1), obsoletes 2616
	//   - RFC 5789: PATCH Method for HTTP
	HTTPMethod HTTPMethodCustomRule `json:"http_method,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool                    `json:"negation"`
	JSON     httpMethodConditionJSON `json:"-"`
}

// httpMethodConditionJSON contains the JSON metadata for the struct
// [HTTPMethodCondition]
type httpMethodConditionJSON struct {
	HTTPMethod  apijson.Field
	Negation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPMethodCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpMethodConditionJSON) RawJSON() string {
	return r.raw
}

// Match the incoming HTTP method
type HTTPMethodConditionParam struct {
	// HTTP methods and descriptions
	//
	// Methods from the following RFCs are all observed:
	//
	//   - RFC 7231: Hypertext Transfer Protocol (HTTP/1.1), obsoletes 2616
	//   - RFC 5789: PATCH Method for HTTP
	HTTPMethod param.Field[HTTPMethodCustomRule] `json:"http_method,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r HTTPMethodConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// HTTP methods and descriptions
//
// Methods from the following RFCs are all observed:
//
//   - RFC 7231: Hypertext Transfer Protocol (HTTP/1.1), obsoletes 2616
//   - RFC 5789: PATCH Method for HTTP
type HTTPMethodCustomRule string

const (
	HTTPMethodCustomRuleConnect HTTPMethodCustomRule = "CONNECT"
	HTTPMethodCustomRuleDelete  HTTPMethodCustomRule = "DELETE"
	HTTPMethodCustomRuleGet     HTTPMethodCustomRule = "GET"
	HTTPMethodCustomRuleHead    HTTPMethodCustomRule = "HEAD"
	HTTPMethodCustomRuleOptions HTTPMethodCustomRule = "OPTIONS"
	HTTPMethodCustomRulePatch   HTTPMethodCustomRule = "PATCH"
	HTTPMethodCustomRulePost    HTTPMethodCustomRule = "POST"
	HTTPMethodCustomRulePut     HTTPMethodCustomRule = "PUT"
	HTTPMethodCustomRuleTrace   HTTPMethodCustomRule = "TRACE"
)

func (r HTTPMethodCustomRule) IsKnown() bool {
	switch r {
	case HTTPMethodCustomRuleConnect, HTTPMethodCustomRuleDelete, HTTPMethodCustomRuleGet, HTTPMethodCustomRuleHead, HTTPMethodCustomRuleOptions, HTTPMethodCustomRulePatch, HTTPMethodCustomRulePost, HTTPMethodCustomRulePut, HTTPMethodCustomRuleTrace:
		return true
	}
	return false
}

// Match the incoming request against a single IP address
type IPCondition struct {
	// A single IPv4 or IPv6 address
	IPAddress string `json:"ip_address,required" format:"ipv4"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool            `json:"negation"`
	JSON     ipConditionJSON `json:"-"`
}

// ipConditionJSON contains the JSON metadata for the struct [IPCondition]
type ipConditionJSON struct {
	IPAddress   apijson.Field
	Negation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipConditionJSON) RawJSON() string {
	return r.raw
}

// Match the incoming request against a single IP address
type IPConditionParam struct {
	// A single IPv4 or IPv6 address
	IPAddress param.Field[string] `json:"ip_address,required" format:"ipv4"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r IPConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match the incoming request against an IP range
type IPRangeCondition struct {
	// The lower bound IPv4 or IPv6 address to match against
	LowerBound string `json:"lower_bound,required" format:"ipv4"`
	// The upper bound IPv4 or IPv6 address to match against
	UpperBound string `json:"upper_bound,required" format:"ipv4"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool                 `json:"negation"`
	JSON     ipRangeConditionJSON `json:"-"`
}

// ipRangeConditionJSON contains the JSON metadata for the struct
// [IPRangeCondition]
type ipRangeConditionJSON struct {
	LowerBound  apijson.Field
	UpperBound  apijson.Field
	Negation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPRangeCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipRangeConditionJSON) RawJSON() string {
	return r.raw
}

// Match the incoming request against an IP range
type IPRangeConditionParam struct {
	// The lower bound IPv4 or IPv6 address to match against
	LowerBound param.Field[string] `json:"lower_bound,required" format:"ipv4"`
	// The upper bound IPv4 or IPv6 address to match against
	UpperBound param.Field[string] `json:"upper_bound,required" format:"ipv4"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r IPRangeConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match the organization the request originated from, as determined by a WHOIS
// lookup of the requesting IP
type OrganizationCondition struct {
	// The organization to match against
	Organization string `json:"organization,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool                      `json:"negation"`
	JSON     organizationConditionJSON `json:"-"`
}

// organizationConditionJSON contains the JSON metadata for the struct
// [OrganizationCondition]
type organizationConditionJSON struct {
	Organization apijson.Field
	Negation     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationConditionJSON) RawJSON() string {
	return r.raw
}

// Match the organization the request originated from, as determined by a WHOIS
// lookup of the requesting IP
type OrganizationConditionParam struct {
	// The organization to match against
	Organization param.Field[string] `json:"organization,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r OrganizationConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match the type of organization that owns the IP address making an incoming
// request
type OwnerTypesCondition struct {
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// Match the type of organization that owns the IP address making an incoming
	// request
	OwnerTypes []OwnerTypesConditionOwnerType `json:"owner_types"`
	JSON       ownerTypesConditionJSON        `json:"-"`
}

// ownerTypesConditionJSON contains the JSON metadata for the struct
// [OwnerTypesCondition]
type ownerTypesConditionJSON struct {
	Negation    apijson.Field
	OwnerTypes  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnerTypesCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownerTypesConditionJSON) RawJSON() string {
	return r.raw
}

type OwnerTypesConditionOwnerType string

const (
	OwnerTypesConditionOwnerTypeCommercial      OwnerTypesConditionOwnerType = "COMMERCIAL"
	OwnerTypesConditionOwnerTypeEducational     OwnerTypesConditionOwnerType = "EDUCATIONAL"
	OwnerTypesConditionOwnerTypeGovernment      OwnerTypesConditionOwnerType = "GOVERNMENT"
	OwnerTypesConditionOwnerTypeHostingServices OwnerTypesConditionOwnerType = "HOSTING_SERVICES"
	OwnerTypesConditionOwnerTypeIsp             OwnerTypesConditionOwnerType = "ISP"
	OwnerTypesConditionOwnerTypeMobileNetwork   OwnerTypesConditionOwnerType = "MOBILE_NETWORK"
	OwnerTypesConditionOwnerTypeNetwork         OwnerTypesConditionOwnerType = "NETWORK"
	OwnerTypesConditionOwnerTypeReserved        OwnerTypesConditionOwnerType = "RESERVED"
)

func (r OwnerTypesConditionOwnerType) IsKnown() bool {
	switch r {
	case OwnerTypesConditionOwnerTypeCommercial, OwnerTypesConditionOwnerTypeEducational, OwnerTypesConditionOwnerTypeGovernment, OwnerTypesConditionOwnerTypeHostingServices, OwnerTypesConditionOwnerTypeIsp, OwnerTypesConditionOwnerTypeMobileNetwork, OwnerTypesConditionOwnerTypeNetwork, OwnerTypesConditionOwnerTypeReserved:
		return true
	}
	return false
}

// Match the type of organization that owns the IP address making an incoming
// request
type OwnerTypesConditionParam struct {
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
	// Match the type of organization that owns the IP address making an incoming
	// request
	OwnerTypes param.Field[[]OwnerTypesConditionOwnerType] `json:"owner_types"`
}

func (r OwnerTypesConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match the rate at which requests come in that match certain conditions
type RequestRateCondition struct {
	// A regular expression matching the URL path of the incoming request
	PathPattern string `json:"path_pattern,required"`
	// The number of incoming requests over the given time that can trigger a request
	// rate condition
	Requests int64 `json:"requests,required"`
	// The number of seconds that the WAAP measures incoming requests over before
	// triggering a request rate condition
	Time int64 `json:"time,required"`
	// Possible HTTP request methods that can trigger a request rate condition
	HTTPMethods []HTTPMethodCustomRule `json:"http_methods,nullable"`
	// A list of source IPs that can trigger a request rate condition
	IPs []string `json:"ips,nullable" format:"ipv4"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// A user-defined tag that can be included in incoming requests and used to trigger
	// a request rate condition
	UserDefinedTag string                   `json:"user_defined_tag,nullable"`
	JSON           requestRateConditionJSON `json:"-"`
}

// requestRateConditionJSON contains the JSON metadata for the struct
// [RequestRateCondition]
type requestRateConditionJSON struct {
	PathPattern    apijson.Field
	Requests       apijson.Field
	Time           apijson.Field
	HTTPMethods    apijson.Field
	IPs            apijson.Field
	Negation       apijson.Field
	UserDefinedTag apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RequestRateCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestRateConditionJSON) RawJSON() string {
	return r.raw
}

// Match the rate at which requests come in that match certain conditions
type RequestRateConditionParam struct {
	// A regular expression matching the URL path of the incoming request
	PathPattern param.Field[string] `json:"path_pattern,required"`
	// The number of incoming requests over the given time that can trigger a request
	// rate condition
	Requests param.Field[int64] `json:"requests,required"`
	// The number of seconds that the WAAP measures incoming requests over before
	// triggering a request rate condition
	Time param.Field[int64] `json:"time,required"`
	// Possible HTTP request methods that can trigger a request rate condition
	HTTPMethods param.Field[[]HTTPMethodCustomRule] `json:"http_methods"`
	// A list of source IPs that can trigger a request rate condition
	IPs param.Field[[]string] `json:"ips" format:"ipv4"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
	// A user-defined tag that can be included in incoming requests and used to trigger
	// a request rate condition
	UserDefinedTag param.Field[string] `json:"user_defined_tag"`
}

func (r RequestRateConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match a response header
type ResponseHeaderCondition struct {
	// The response header name
	Header string `json:"header,required"`
	// The response header value
	Value string `json:"value,required"`
	// The type of matching condition for header and value.
	MatchType ResponseHeaderConditionMatchType `json:"match_type"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool                        `json:"negation"`
	JSON     responseHeaderConditionJSON `json:"-"`
}

// responseHeaderConditionJSON contains the JSON metadata for the struct
// [ResponseHeaderCondition]
type responseHeaderConditionJSON struct {
	Header      apijson.Field
	Value       apijson.Field
	MatchType   apijson.Field
	Negation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseHeaderCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseHeaderConditionJSON) RawJSON() string {
	return r.raw
}

// The type of matching condition for header and value.
type ResponseHeaderConditionMatchType string

const (
	ResponseHeaderConditionMatchTypeExact    ResponseHeaderConditionMatchType = "Exact"
	ResponseHeaderConditionMatchTypeContains ResponseHeaderConditionMatchType = "Contains"
)

func (r ResponseHeaderConditionMatchType) IsKnown() bool {
	switch r {
	case ResponseHeaderConditionMatchTypeExact, ResponseHeaderConditionMatchTypeContains:
		return true
	}
	return false
}

// Match a response header
type ResponseHeaderConditionParam struct {
	// The response header name
	Header param.Field[string] `json:"header,required"`
	// The response header value
	Value param.Field[string] `json:"value,required"`
	// The type of matching condition for header and value.
	MatchType param.Field[ResponseHeaderConditionMatchType] `json:"match_type"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r ResponseHeaderConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match when a response header is present
type ResponseHeaderExistsCondition struct {
	// The response header name
	Header string `json:"header,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool                              `json:"negation"`
	JSON     responseHeaderExistsConditionJSON `json:"-"`
}

// responseHeaderExistsConditionJSON contains the JSON metadata for the struct
// [ResponseHeaderExistsCondition]
type responseHeaderExistsConditionJSON struct {
	Header      apijson.Field
	Negation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseHeaderExistsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseHeaderExistsConditionJSON) RawJSON() string {
	return r.raw
}

// Match when a response header is present
type ResponseHeaderExistsConditionParam struct {
	// The response header name
	Header param.Field[string] `json:"header,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r ResponseHeaderExistsConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleActionType string

const (
	RuleActionTypeAllow     RuleActionType = "allow"
	RuleActionTypeBlock     RuleActionType = "block"
	RuleActionTypeCaptcha   RuleActionType = "captcha"
	RuleActionTypeHandshake RuleActionType = "handshake"
	RuleActionTypeMonitor   RuleActionType = "monitor"
	RuleActionTypeTag       RuleActionType = "tag"
)

func (r RuleActionType) IsKnown() bool {
	switch r {
	case RuleActionTypeAllow, RuleActionTypeBlock, RuleActionTypeCaptcha, RuleActionTypeHandshake, RuleActionTypeMonitor, RuleActionTypeTag:
		return true
	}
	return false
}

type RuleAllowAction = interface{}

// WAAP block action behavior could be configured with response status code and
// action duration.
type RuleBlockAction struct {
	// How long a rule's block action will apply to subsequent requests. Can be
	// specified in seconds or by using a numeral followed by 's', 'm', 'h', or 'd' to
	// represent time format (seconds, minutes, hours, or days)
	ActionDuration string `json:"action_duration,nullable"`
	// Designates the HTTP status code to deliver when a request is blocked.
	StatusCode RuleBlockActionStatusCode `json:"status_code,nullable"`
	JSON       ruleBlockActionJSON       `json:"-"`
}

// ruleBlockActionJSON contains the JSON metadata for the struct [RuleBlockAction]
type ruleBlockActionJSON struct {
	ActionDuration apijson.Field
	StatusCode     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuleBlockAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleBlockActionJSON) RawJSON() string {
	return r.raw
}

// Designates the HTTP status code to deliver when a request is blocked.
type RuleBlockActionStatusCode int64

const (
	RuleBlockActionStatusCode403 RuleBlockActionStatusCode = 403
	RuleBlockActionStatusCode405 RuleBlockActionStatusCode = 405
	RuleBlockActionStatusCode418 RuleBlockActionStatusCode = 418
	RuleBlockActionStatusCode429 RuleBlockActionStatusCode = 429
)

func (r RuleBlockActionStatusCode) IsKnown() bool {
	switch r {
	case RuleBlockActionStatusCode403, RuleBlockActionStatusCode405, RuleBlockActionStatusCode418, RuleBlockActionStatusCode429:
		return true
	}
	return false
}

// WAAP block action behavior could be configured with response status code and
// action duration.
type RuleBlockActionParam struct {
	// How long a rule's block action will apply to subsequent requests. Can be
	// specified in seconds or by using a numeral followed by 's', 'm', 'h', or 'd' to
	// represent time format (seconds, minutes, hours, or days)
	ActionDuration param.Field[string] `json:"action_duration"`
	// Designates the HTTP status code to deliver when a request is blocked.
	StatusCode param.Field[RuleBlockActionStatusCode] `json:"status_code"`
}

func (r RuleBlockActionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleCaptchaAction = interface{}

type RuleHandshakeAction = interface{}

type RuleMonitorAction = interface{}

// WAAP tag action gets a list of tags to tag the request scope with
type RuleTagAction struct {
	// The list of user defined tags to tag the request with
	Tags []string          `json:"tags,required"`
	JSON ruleTagActionJSON `json:"-"`
}

// ruleTagActionJSON contains the JSON metadata for the struct [RuleTagAction]
type ruleTagActionJSON struct {
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleTagAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleTagActionJSON) RawJSON() string {
	return r.raw
}

// WAAP tag action gets a list of tags to tag the request scope with
type RuleTagActionParam struct {
	// The list of user defined tags to tag the request with
	Tags param.Field[[]string] `json:"tags,required"`
}

func (r RuleTagActionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A request to delete a list of rules
type RulesBulkDeleteParam struct {
	// The IDs of the rules to delete
	RuleIDs param.Field[[]int64] `json:"rule_ids,required"`
}

func (r RulesBulkDeleteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match the number of dynamic page requests made in a WAAP session
type SessionRequestCountCondition struct {
	// The number of dynamic requests in the session
	RequestCount int64 `json:"request_count,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool                             `json:"negation"`
	JSON     sessionRequestCountConditionJSON `json:"-"`
}

// sessionRequestCountConditionJSON contains the JSON metadata for the struct
// [SessionRequestCountCondition]
type sessionRequestCountConditionJSON struct {
	RequestCount apijson.Field
	Negation     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SessionRequestCountCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionRequestCountConditionJSON) RawJSON() string {
	return r.raw
}

// Match the number of dynamic page requests made in a WAAP session
type SessionRequestCountConditionParam struct {
	// The number of dynamic requests in the session
	RequestCount param.Field[int64] `json:"request_count,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r SessionRequestCountConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches requests based on specified tags
type TagsCondition struct {
	// A list of tags to match against the request tags
	Tags []string `json:"tags,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool              `json:"negation"`
	JSON     tagsConditionJSON `json:"-"`
}

// tagsConditionJSON contains the JSON metadata for the struct [TagsCondition]
type tagsConditionJSON struct {
	Tags        apijson.Field
	Negation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TagsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tagsConditionJSON) RawJSON() string {
	return r.raw
}

// Matches requests based on specified tags
type TagsConditionParam struct {
	// A list of tags to match against the request tags
	Tags param.Field[[]string] `json:"tags,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r TagsConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match the incoming request URL
type URLCondition struct {
	// The URL to match
	URL string `json:"url,required"`
	// The type of matching condition.
	MatchType URLConditionMatchType `json:"match_type"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool             `json:"negation"`
	JSON     urlConditionJSON `json:"-"`
}

// urlConditionJSON contains the JSON metadata for the struct [URLCondition]
type urlConditionJSON struct {
	URL         apijson.Field
	MatchType   apijson.Field
	Negation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlConditionJSON) RawJSON() string {
	return r.raw
}

// The type of matching condition.
type URLConditionMatchType string

const (
	URLConditionMatchTypeExact    URLConditionMatchType = "Exact"
	URLConditionMatchTypeContains URLConditionMatchType = "Contains"
	URLConditionMatchTypeRegex    URLConditionMatchType = "Regex"
)

func (r URLConditionMatchType) IsKnown() bool {
	switch r {
	case URLConditionMatchTypeExact, URLConditionMatchTypeContains, URLConditionMatchTypeRegex:
		return true
	}
	return false
}

// Match the incoming request URL
type URLConditionParam struct {
	// The URL to match
	URL param.Field[string] `json:"url,required"`
	// The type of matching condition.
	MatchType param.Field[URLConditionMatchType] `json:"match_type"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r URLConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match the user agent making the request
type UserAgentCondition struct {
	// The user agent value to match
	UserAgent string `json:"user_agent,required"`
	// The type of matching condition.
	MatchType UserAgentConditionMatchType `json:"match_type"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool                   `json:"negation"`
	JSON     userAgentConditionJSON `json:"-"`
}

// userAgentConditionJSON contains the JSON metadata for the struct
// [UserAgentCondition]
type userAgentConditionJSON struct {
	UserAgent   apijson.Field
	MatchType   apijson.Field
	Negation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAgentCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAgentConditionJSON) RawJSON() string {
	return r.raw
}

// The type of matching condition.
type UserAgentConditionMatchType string

const (
	UserAgentConditionMatchTypeExact    UserAgentConditionMatchType = "Exact"
	UserAgentConditionMatchTypeContains UserAgentConditionMatchType = "Contains"
)

func (r UserAgentConditionMatchType) IsKnown() bool {
	switch r {
	case UserAgentConditionMatchTypeExact, UserAgentConditionMatchTypeContains:
		return true
	}
	return false
}

// Match the user agent making the request
type UserAgentConditionParam struct {
	// The user agent value to match
	UserAgent param.Field[string] `json:"user_agent,required"`
	// The type of matching condition.
	MatchType param.Field[UserAgentConditionMatchType] `json:"match_type"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r UserAgentConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches requests based on user-defined tags
type UserDefinedTagsCondition struct {
	// A list of user-defined tags to match against the request tags
	Tags []string `json:"tags,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool                         `json:"negation"`
	JSON     userDefinedTagsConditionJSON `json:"-"`
}

// userDefinedTagsConditionJSON contains the JSON metadata for the struct
// [UserDefinedTagsCondition]
type userDefinedTagsConditionJSON struct {
	Tags        apijson.Field
	Negation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserDefinedTagsCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userDefinedTagsConditionJSON) RawJSON() string {
	return r.raw
}

// Matches requests based on user-defined tags
type UserDefinedTagsConditionParam struct {
	// A list of user-defined tags to match against the request tags
	Tags param.Field[[]string] `json:"tags,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Field[bool] `json:"negation"`
}

func (r UserDefinedTagsConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaapV1DomainCustomRuleListResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []CustomRule                           `json:"results,required"`
	JSON    waapV1DomainCustomRuleListResponseJSON `json:"-"`
}

// waapV1DomainCustomRuleListResponseJSON contains the JSON metadata for the struct
// [WaapV1DomainCustomRuleListResponse]
type waapV1DomainCustomRuleListResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainCustomRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainCustomRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainCustomRuleNewParams struct {
	// The action that a WAAP rule takes when triggered
	Action param.Field[CustomerRuleActionInputParam] `json:"action,required"`
	// The conditions required for the WAAP engine to trigger the rule. Rules may have
	// between 1 and 5 conditions. All conditions must pass for the rule to trigger
	Conditions param.Field[[]CustomRuleConditionInputParam] `json:"conditions,required"`
	// Whether or not the rule is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// The name assigned to the rule
	Name param.Field[string] `json:"name,required"`
	// The description assigned to the rule
	Description param.Field[string] `json:"description"`
}

func (r WaapV1DomainCustomRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaapV1DomainCustomRuleUpdateParams struct {
	// The action that a WAAP rule takes when triggered
	Action param.Field[CustomerRuleActionInputParam] `json:"action"`
	// The conditions required for the WAAP engine to trigger the rule. Rules may have
	// between 1 and 5 conditions. All conditions must pass for the rule to trigger
	Conditions param.Field[[]CustomRuleConditionInputParam] `json:"conditions"`
	// The description assigned to the rule
	Description param.Field[string] `json:"description"`
	// Whether or not the rule is enabled
	Enabled param.Field[bool] `json:"enabled"`
	// The name assigned to the rule
	Name param.Field[string] `json:"name"`
}

func (r WaapV1DomainCustomRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaapV1DomainCustomRuleListParams struct {
	// Filter to refine results by specific actions
	Action param.Field[RuleActionType] `query:"action"`
	// Filter rules based on their description. Supports '\*' as a wildcard character.
	Description param.Field[string] `query:"description"`
	// Filter rules based on their active status
	Enabled param.Field[bool] `query:"enabled"`
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Filter rules based on their name. Supports '\*' as a wildcard character.
	Name param.Field[string] `query:"name"`
	// Number of items to skip
	Offset param.Field[int64] `query:"offset"`
	// Determine the field to order results by
	Ordering param.Field[WaapV1DomainCustomRuleListParamsOrdering] `query:"ordering"`
}

// URLQuery serializes [WaapV1DomainCustomRuleListParams]'s query parameters as
// `url.Values`.
func (r WaapV1DomainCustomRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Determine the field to order results by
type WaapV1DomainCustomRuleListParamsOrdering string

const (
	WaapV1DomainCustomRuleListParamsOrderingID               WaapV1DomainCustomRuleListParamsOrdering = "id"
	WaapV1DomainCustomRuleListParamsOrderingName             WaapV1DomainCustomRuleListParamsOrdering = "name"
	WaapV1DomainCustomRuleListParamsOrderingDescription      WaapV1DomainCustomRuleListParamsOrdering = "description"
	WaapV1DomainCustomRuleListParamsOrderingEnabled          WaapV1DomainCustomRuleListParamsOrdering = "enabled"
	WaapV1DomainCustomRuleListParamsOrderingAction           WaapV1DomainCustomRuleListParamsOrdering = "action"
	WaapV1DomainCustomRuleListParamsOrderingMinusID          WaapV1DomainCustomRuleListParamsOrdering = "-id"
	WaapV1DomainCustomRuleListParamsOrderingMinusName        WaapV1DomainCustomRuleListParamsOrdering = "-name"
	WaapV1DomainCustomRuleListParamsOrderingMinusDescription WaapV1DomainCustomRuleListParamsOrdering = "-description"
	WaapV1DomainCustomRuleListParamsOrderingMinusEnabled     WaapV1DomainCustomRuleListParamsOrdering = "-enabled"
	WaapV1DomainCustomRuleListParamsOrderingMinusAction      WaapV1DomainCustomRuleListParamsOrdering = "-action"
)

func (r WaapV1DomainCustomRuleListParamsOrdering) IsKnown() bool {
	switch r {
	case WaapV1DomainCustomRuleListParamsOrderingID, WaapV1DomainCustomRuleListParamsOrderingName, WaapV1DomainCustomRuleListParamsOrderingDescription, WaapV1DomainCustomRuleListParamsOrderingEnabled, WaapV1DomainCustomRuleListParamsOrderingAction, WaapV1DomainCustomRuleListParamsOrderingMinusID, WaapV1DomainCustomRuleListParamsOrderingMinusName, WaapV1DomainCustomRuleListParamsOrderingMinusDescription, WaapV1DomainCustomRuleListParamsOrderingMinusEnabled, WaapV1DomainCustomRuleListParamsOrderingMinusAction:
		return true
	}
	return false
}

type WaapV1DomainCustomRuleBulkDeleteParams struct {
	// A request to delete a list of rules
	RulesBulkDelete RulesBulkDeleteParam `json:"rules_bulk_delete,required"`
}

func (r WaapV1DomainCustomRuleBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RulesBulkDelete)
}
