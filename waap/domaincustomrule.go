// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
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

// DomainCustomRuleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainCustomRuleService] method instead.
type DomainCustomRuleService struct {
	Options []option.RequestOption
}

// NewDomainCustomRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDomainCustomRuleService(opts ...option.RequestOption) (r DomainCustomRuleService) {
	r = DomainCustomRuleService{}
	r.Options = opts
	return
}

// Create a custom rule
func (r *DomainCustomRuleService) New(ctx context.Context, domainID int64, body DomainCustomRuleNewParams, opts ...option.RequestOption) (res *WaapCustomRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/custom-rules", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Only properties present in the request will be updated
func (r *DomainCustomRuleService) Update(ctx context.Context, ruleID int64, params DomainCustomRuleUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/custom-rules/%v", params.DomainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return
}

// Extracts a list of custom rules assigned to a domain, offering filter, ordering,
// and pagination capabilities
func (r *DomainCustomRuleService) List(ctx context.Context, domainID int64, query DomainCustomRuleListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapCustomRule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/custom-rules", domainID)
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

// Extracts a list of custom rules assigned to a domain, offering filter, ordering,
// and pagination capabilities
func (r *DomainCustomRuleService) ListAutoPaging(ctx context.Context, domainID int64, query DomainCustomRuleListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapCustomRule] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, domainID, query, opts...))
}

// Delete a custom rule
func (r *DomainCustomRuleService) Delete(ctx context.Context, ruleID int64, body DomainCustomRuleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/custom-rules/%v", body.DomainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Delete multiple WAAP rules
func (r *DomainCustomRuleService) DeleteMultiple(ctx context.Context, domainID int64, body DomainCustomRuleDeleteMultipleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/custom-rules/bulk_delete", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Extracts a specific custom rule assigned to a domain
func (r *DomainCustomRuleService) Get(ctx context.Context, ruleID int64, query DomainCustomRuleGetParams, opts ...option.RequestOption) (res *WaapCustomRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/custom-rules/%v", query.DomainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Toggle a custom rule
func (r *DomainCustomRuleService) Toggle(ctx context.Context, action DomainCustomRuleToggleParamsAction, body DomainCustomRuleToggleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/custom-rules/%v/%v", body.DomainID, body.RuleID, action)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, nil, opts...)
	return
}

// An WAAP rule applied to a domain
type WaapCustomRule struct {
	// The unique identifier for the rule
	ID int64 `json:"id,required"`
	// The action that the rule takes when triggered. Only one action can be set per
	// rule.
	Action WaapCustomRuleAction `json:"action,required"`
	// The conditions required for the WAAP engine to trigger the rule. Rules may have
	// between 1 and 5 conditions. All conditions must pass for the rule to trigger
	Conditions []WaapCustomRuleCondition `json:"conditions,required"`
	// Whether or not the rule is enabled
	Enabled bool `json:"enabled,required"`
	// The name assigned to the rule
	Name string `json:"name,required"`
	// The description assigned to the rule
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Action      respjson.Field
		Conditions  respjson.Field
		Enabled     respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRule) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the rule takes when triggered. Only one action can be set per
// rule.
type WaapCustomRuleAction struct {
	// The WAAP allowed the request
	Allow any `json:"allow"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block WaapCustomRuleActionBlock `json:"block"`
	// The WAAP presented the user with a captcha
	Captcha any `json:"captcha"`
	// The WAAP performed automatic browser validation
	Handshake any `json:"handshake"`
	// The WAAP monitored the request but took no action
	Monitor any `json:"monitor"`
	// WAAP tag action gets a list of tags to tag the request scope with
	Tag WaapCustomRuleActionTag `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allow       respjson.Field
		Block       respjson.Field
		Captcha     respjson.Field
		Handshake   respjson.Field
		Monitor     respjson.Field
		Tag         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleAction) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WAAP block action behavior could be configured with response status code and
// action duration.
type WaapCustomRuleActionBlock struct {
	// How long a rule's block action will apply to subsequent requests. Can be
	// specified in seconds or by using a numeral followed by 's', 'm', 'h', or 'd' to
	// represent time format (seconds, minutes, hours, or days). Empty time intervals
	// are not allowed.
	ActionDuration string `json:"action_duration"`
	// A custom HTTP status code that the WAAP returns if a rule blocks a request
	//
	// Any of 403, 405, 418, 429.
	StatusCode int64 `json:"status_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionDuration respjson.Field
		StatusCode     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleActionBlock) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleActionBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WAAP tag action gets a list of tags to tag the request scope with
type WaapCustomRuleActionTag struct {
	// The list of user defined tags to tag the request with
	Tags []string `json:"tags,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleActionTag) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleActionTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The criteria of an incoming web request and the models of the various values
// those criteria can take
type WaapCustomRuleCondition struct {
	// Match the requested Content-Type
	ContentType WaapCustomRuleConditionContentType `json:"content_type"`
	// Match the country that the request originated from
	Country WaapCustomRuleConditionCountry `json:"country"`
	// Match the incoming file extension
	FileExtension WaapCustomRuleConditionFileExtension `json:"file_extension"`
	// Match an incoming request header
	Header WaapCustomRuleConditionHeader `json:"header"`
	// Match when an incoming request header is present
	HeaderExists WaapCustomRuleConditionHeaderExists `json:"header_exists"`
	// Match the incoming HTTP method
	HTTPMethod WaapCustomRuleConditionHTTPMethod `json:"http_method"`
	// Match the incoming request against a single IP address
	IP WaapCustomRuleConditionIP `json:"ip"`
	// Match the incoming request against an IP range
	IPRange WaapCustomRuleConditionIPRange `json:"ip_range"`
	// Match the organization the request originated from, as determined by a WHOIS
	// lookup of the requesting IP
	Organization WaapCustomRuleConditionOrganization `json:"organization"`
	// Match the type of organization that owns the IP address making an incoming
	// request
	OwnerTypes WaapCustomRuleConditionOwnerTypes `json:"owner_types"`
	// Match the rate at which requests come in that match certain conditions
	RequestRate WaapCustomRuleConditionRequestRate `json:"request_rate"`
	// Match a response header
	ResponseHeader WaapCustomRuleConditionResponseHeader `json:"response_header"`
	// Match when a response header is present
	ResponseHeaderExists WaapCustomRuleConditionResponseHeaderExists `json:"response_header_exists"`
	// Match the number of dynamic page requests made in a WAAP session
	SessionRequestCount WaapCustomRuleConditionSessionRequestCount `json:"session_request_count"`
	// Matches requests based on specified tags
	Tags WaapCustomRuleConditionTags `json:"tags"`
	// Match the incoming request URL
	URL WaapCustomRuleConditionURL `json:"url"`
	// Match the user agent making the request
	UserAgent WaapCustomRuleConditionUserAgent `json:"user_agent"`
	// Matches requests based on user-defined tags
	UserDefinedTags WaapCustomRuleConditionUserDefinedTags `json:"user_defined_tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType          respjson.Field
		Country              respjson.Field
		FileExtension        respjson.Field
		Header               respjson.Field
		HeaderExists         respjson.Field
		HTTPMethod           respjson.Field
		IP                   respjson.Field
		IPRange              respjson.Field
		Organization         respjson.Field
		OwnerTypes           respjson.Field
		RequestRate          respjson.Field
		ResponseHeader       respjson.Field
		ResponseHeaderExists respjson.Field
		SessionRequestCount  respjson.Field
		Tags                 respjson.Field
		URL                  respjson.Field
		UserAgent            respjson.Field
		UserDefinedTags      respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleCondition) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the requested Content-Type
type WaapCustomRuleConditionContentType struct {
	// The list of content types to match against
	ContentType []string `json:"content_type,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		Negation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionContentType) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionContentType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the country that the request originated from
type WaapCustomRuleConditionCountry struct {
	// A list of ISO 3166-1 alpha-2 formatted strings representing the countries to
	// match against
	CountryCode []string `json:"country_code,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode respjson.Field
		Negation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionCountry) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming file extension
type WaapCustomRuleConditionFileExtension struct {
	// The list of file extensions to match against
	FileExtension []string `json:"file_extension,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileExtension respjson.Field
		Negation      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionFileExtension) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionFileExtension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match an incoming request header
type WaapCustomRuleConditionHeader struct {
	// The request header name
	Header string `json:"header,required"`
	// The request header value
	Value string `json:"value,required"`
	// The type of matching condition for header and value.
	//
	// Any of "Exact", "Contains".
	MatchType string `json:"match_type"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Header      respjson.Field
		Value       respjson.Field
		MatchType   respjson.Field
		Negation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionHeader) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match when an incoming request header is present
type WaapCustomRuleConditionHeaderExists struct {
	// The request header name
	Header string `json:"header,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Header      respjson.Field
		Negation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionHeaderExists) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionHeaderExists) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming HTTP method
type WaapCustomRuleConditionHTTPMethod struct {
	// HTTP methods of a request
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod string `json:"http_method,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HTTPMethod  respjson.Field
		Negation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionHTTPMethod) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionHTTPMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against a single IP address
type WaapCustomRuleConditionIP struct {
	// A single IPv4 or IPv6 address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress   respjson.Field
		Negation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionIP) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against an IP range
type WaapCustomRuleConditionIPRange struct {
	// The lower bound IPv4 or IPv6 address to match against
	LowerBound string `json:"lower_bound,required" format:"ipvanyaddress"`
	// The upper bound IPv4 or IPv6 address to match against
	UpperBound string `json:"upper_bound,required" format:"ipvanyaddress"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LowerBound  respjson.Field
		UpperBound  respjson.Field
		Negation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionIPRange) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionIPRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the organization the request originated from, as determined by a WHOIS
// lookup of the requesting IP
type WaapCustomRuleConditionOrganization struct {
	// The organization to match against
	Organization string `json:"organization,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organization respjson.Field
		Negation     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionOrganization) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the type of organization that owns the IP address making an incoming
// request
type WaapCustomRuleConditionOwnerTypes struct {
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// Match the type of organization that owns the IP address making an incoming
	// request
	//
	// Any of "COMMERCIAL", "EDUCATIONAL", "GOVERNMENT", "HOSTING_SERVICES", "ISP",
	// "MOBILE_NETWORK", "NETWORK", "RESERVED".
	OwnerTypes []string `json:"owner_types"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Negation    respjson.Field
		OwnerTypes  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionOwnerTypes) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionOwnerTypes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the rate at which requests come in that match certain conditions
type WaapCustomRuleConditionRequestRate struct {
	// A regular expression matching the URL path of the incoming request
	PathPattern string `json:"path_pattern,required"`
	// The number of incoming requests over the given time that can trigger a request
	// rate condition
	Requests int64 `json:"requests,required"`
	// The number of seconds that the WAAP measures incoming requests over before
	// triggering a request rate condition
	Time int64 `json:"time,required"`
	// Possible HTTP request methods that can trigger a request rate condition
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethods []string `json:"http_methods"`
	// A list of source IPs that can trigger a request rate condition
	IPs []string `json:"ips" format:"ipvanyaddress"`
	// A user-defined tag that can be included in incoming requests and used to trigger
	// a request rate condition
	UserDefinedTag string `json:"user_defined_tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PathPattern    respjson.Field
		Requests       respjson.Field
		Time           respjson.Field
		HTTPMethods    respjson.Field
		IPs            respjson.Field
		UserDefinedTag respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionRequestRate) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionRequestRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match a response header
type WaapCustomRuleConditionResponseHeader struct {
	// The response header name
	Header string `json:"header,required"`
	// The response header value
	Value string `json:"value,required"`
	// The type of matching condition for header and value.
	//
	// Any of "Exact", "Contains".
	MatchType string `json:"match_type"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Header      respjson.Field
		Value       respjson.Field
		MatchType   respjson.Field
		Negation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionResponseHeader) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionResponseHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match when a response header is present
type WaapCustomRuleConditionResponseHeaderExists struct {
	// The response header name
	Header string `json:"header,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Header      respjson.Field
		Negation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionResponseHeaderExists) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionResponseHeaderExists) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the number of dynamic page requests made in a WAAP session
type WaapCustomRuleConditionSessionRequestCount struct {
	// The number of dynamic requests in the session
	RequestCount int64 `json:"request_count,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestCount respjson.Field
		Negation     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionSessionRequestCount) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionSessionRequestCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Matches requests based on specified tags
type WaapCustomRuleConditionTags struct {
	// A list of tags to match against the request tags
	Tags []string `json:"tags,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tags        respjson.Field
		Negation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionTags) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request URL
type WaapCustomRuleConditionURL struct {
	// The pattern to match against the request URL. Constraints depend on
	// `match_type`:
	//
	//   - **Exact/Contains**: plain text matching (e.g., `/admin`, must comply with
	//     `^[\w!\$~:#\[\]@\(\)\\*\+,=\/\-\.\%]+$`).
	//   - **Regex**: a valid regular expression (e.g., `^/upload(/\d+)?/\w+`).
	//     Lookahead/lookbehind constructs are forbidden.
	URL string `json:"url,required"`
	// The type of matching condition.
	//
	// Any of "Exact", "Contains", "Regex".
	MatchType string `json:"match_type"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		MatchType   respjson.Field
		Negation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionURL) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the user agent making the request
type WaapCustomRuleConditionUserAgent struct {
	// The user agent value to match
	UserAgent string `json:"user_agent,required"`
	// The type of matching condition.
	//
	// Any of "Exact", "Contains".
	MatchType string `json:"match_type"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserAgent   respjson.Field
		MatchType   respjson.Field
		Negation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionUserAgent) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionUserAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Matches requests based on user-defined tags
type WaapCustomRuleConditionUserDefinedTags struct {
	// A list of user-defined tags to match against the request tags
	Tags []string `json:"tags,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tags        respjson.Field
		Negation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomRuleConditionUserDefinedTags) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomRuleConditionUserDefinedTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainCustomRuleNewParams struct {
	// The action that the rule takes when triggered. Only one action can be set per
	// rule.
	Action DomainCustomRuleNewParamsAction `json:"action,omitzero,required"`
	// The conditions required for the WAAP engine to trigger the rule. Rules may have
	// between 1 and 5 conditions. All conditions must pass for the rule to trigger
	Conditions []DomainCustomRuleNewParamsCondition `json:"conditions,omitzero,required"`
	// Whether or not the rule is enabled
	Enabled bool `json:"enabled,required"`
	// The name assigned to the rule
	Name string `json:"name,required"`
	// The description assigned to the rule
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the rule takes when triggered. Only one action can be set per
// rule.
type DomainCustomRuleNewParamsAction struct {
	// The WAAP allowed the request
	Allow any `json:"allow,omitzero"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block DomainCustomRuleNewParamsActionBlock `json:"block,omitzero"`
	// The WAAP presented the user with a captcha
	Captcha any `json:"captcha,omitzero"`
	// The WAAP performed automatic browser validation
	Handshake any `json:"handshake,omitzero"`
	// The WAAP monitored the request but took no action
	Monitor any `json:"monitor,omitzero"`
	// WAAP tag action gets a list of tags to tag the request scope with
	Tag DomainCustomRuleNewParamsActionTag `json:"tag,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsAction) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WAAP block action behavior could be configured with response status code and
// action duration.
type DomainCustomRuleNewParamsActionBlock struct {
	// How long a rule's block action will apply to subsequent requests. Can be
	// specified in seconds or by using a numeral followed by 's', 'm', 'h', or 'd' to
	// represent time format (seconds, minutes, hours, or days). Empty time intervals
	// are not allowed.
	ActionDuration param.Opt[string] `json:"action_duration,omitzero"`
	// A custom HTTP status code that the WAAP returns if a rule blocks a request
	//
	// Any of 403, 405, 418, 429.
	StatusCode int64 `json:"status_code,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsActionBlock) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsActionBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsActionBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainCustomRuleNewParamsActionBlock](
		"status_code", 403, 405, 418, 429,
	)
}

// WAAP tag action gets a list of tags to tag the request scope with
//
// The property Tags is required.
type DomainCustomRuleNewParamsActionTag struct {
	// The list of user defined tags to tag the request with
	Tags []string `json:"tags,omitzero,required"`
	paramObj
}

func (r DomainCustomRuleNewParamsActionTag) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsActionTag
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsActionTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The criteria of an incoming web request and the models of the various values
// those criteria can take
type DomainCustomRuleNewParamsCondition struct {
	// Match the requested Content-Type
	ContentType DomainCustomRuleNewParamsConditionContentType `json:"content_type,omitzero"`
	// Match the country that the request originated from
	Country DomainCustomRuleNewParamsConditionCountry `json:"country,omitzero"`
	// Match the incoming file extension
	FileExtension DomainCustomRuleNewParamsConditionFileExtension `json:"file_extension,omitzero"`
	// Match an incoming request header
	Header DomainCustomRuleNewParamsConditionHeader `json:"header,omitzero"`
	// Match when an incoming request header is present
	HeaderExists DomainCustomRuleNewParamsConditionHeaderExists `json:"header_exists,omitzero"`
	// Match the incoming HTTP method
	HTTPMethod DomainCustomRuleNewParamsConditionHTTPMethod `json:"http_method,omitzero"`
	// Match the incoming request against a single IP address
	IP DomainCustomRuleNewParamsConditionIP `json:"ip,omitzero"`
	// Match the incoming request against an IP range
	IPRange DomainCustomRuleNewParamsConditionIPRange `json:"ip_range,omitzero"`
	// Match the organization the request originated from, as determined by a WHOIS
	// lookup of the requesting IP
	Organization DomainCustomRuleNewParamsConditionOrganization `json:"organization,omitzero"`
	// Match the type of organization that owns the IP address making an incoming
	// request
	OwnerTypes DomainCustomRuleNewParamsConditionOwnerTypes `json:"owner_types,omitzero"`
	// Match the rate at which requests come in that match certain conditions
	RequestRate DomainCustomRuleNewParamsConditionRequestRate `json:"request_rate,omitzero"`
	// Match a response header
	ResponseHeader DomainCustomRuleNewParamsConditionResponseHeader `json:"response_header,omitzero"`
	// Match when a response header is present
	ResponseHeaderExists DomainCustomRuleNewParamsConditionResponseHeaderExists `json:"response_header_exists,omitzero"`
	// Match the number of dynamic page requests made in a WAAP session
	SessionRequestCount DomainCustomRuleNewParamsConditionSessionRequestCount `json:"session_request_count,omitzero"`
	// Matches requests based on specified tags
	Tags DomainCustomRuleNewParamsConditionTags `json:"tags,omitzero"`
	// Match the incoming request URL
	URL DomainCustomRuleNewParamsConditionURL `json:"url,omitzero"`
	// Match the user agent making the request
	UserAgent DomainCustomRuleNewParamsConditionUserAgent `json:"user_agent,omitzero"`
	// Matches requests based on user-defined tags
	UserDefinedTags DomainCustomRuleNewParamsConditionUserDefinedTags `json:"user_defined_tags,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsCondition) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the requested Content-Type
//
// The property ContentType is required.
type DomainCustomRuleNewParamsConditionContentType struct {
	// The list of content types to match against
	ContentType []string `json:"content_type,omitzero,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionContentType) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionContentType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionContentType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the country that the request originated from
//
// The property CountryCode is required.
type DomainCustomRuleNewParamsConditionCountry struct {
	// A list of ISO 3166-1 alpha-2 formatted strings representing the countries to
	// match against
	CountryCode []string `json:"country_code,omitzero,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionCountry) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionCountry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming file extension
//
// The property FileExtension is required.
type DomainCustomRuleNewParamsConditionFileExtension struct {
	// The list of file extensions to match against
	FileExtension []string `json:"file_extension,omitzero,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionFileExtension) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionFileExtension
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionFileExtension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match an incoming request header
//
// The properties Header, Value are required.
type DomainCustomRuleNewParamsConditionHeader struct {
	// The request header name
	Header string `json:"header,required"`
	// The request header value
	Value string `json:"value,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// The type of matching condition for header and value.
	//
	// Any of "Exact", "Contains".
	MatchType string `json:"match_type,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionHeader) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainCustomRuleNewParamsConditionHeader](
		"match_type", "Exact", "Contains",
	)
}

// Match when an incoming request header is present
//
// The property Header is required.
type DomainCustomRuleNewParamsConditionHeaderExists struct {
	// The request header name
	Header string `json:"header,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionHeaderExists) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionHeaderExists
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionHeaderExists) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming HTTP method
//
// The property HTTPMethod is required.
type DomainCustomRuleNewParamsConditionHTTPMethod struct {
	// HTTP methods of a request
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod string `json:"http_method,omitzero,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionHTTPMethod) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionHTTPMethod
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionHTTPMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainCustomRuleNewParamsConditionHTTPMethod](
		"http_method", "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT", "TRACE",
	)
}

// Match the incoming request against a single IP address
//
// The property IPAddress is required.
type DomainCustomRuleNewParamsConditionIP struct {
	// A single IPv4 or IPv6 address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionIP) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against an IP range
//
// The properties LowerBound, UpperBound are required.
type DomainCustomRuleNewParamsConditionIPRange struct {
	// The lower bound IPv4 or IPv6 address to match against
	LowerBound string `json:"lower_bound,required" format:"ipvanyaddress"`
	// The upper bound IPv4 or IPv6 address to match against
	UpperBound string `json:"upper_bound,required" format:"ipvanyaddress"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionIPRange) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionIPRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionIPRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the organization the request originated from, as determined by a WHOIS
// lookup of the requesting IP
//
// The property Organization is required.
type DomainCustomRuleNewParamsConditionOrganization struct {
	// The organization to match against
	Organization string `json:"organization,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionOrganization) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionOrganization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the type of organization that owns the IP address making an incoming
// request
type DomainCustomRuleNewParamsConditionOwnerTypes struct {
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// Match the type of organization that owns the IP address making an incoming
	// request
	//
	// Any of "COMMERCIAL", "EDUCATIONAL", "GOVERNMENT", "HOSTING_SERVICES", "ISP",
	// "MOBILE_NETWORK", "NETWORK", "RESERVED".
	OwnerTypes []string `json:"owner_types,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionOwnerTypes) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionOwnerTypes
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionOwnerTypes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the rate at which requests come in that match certain conditions
//
// The properties PathPattern, Requests, Time are required.
type DomainCustomRuleNewParamsConditionRequestRate struct {
	// A regular expression matching the URL path of the incoming request
	PathPattern string `json:"path_pattern,required"`
	// The number of incoming requests over the given time that can trigger a request
	// rate condition
	Requests int64 `json:"requests,required"`
	// The number of seconds that the WAAP measures incoming requests over before
	// triggering a request rate condition
	Time int64 `json:"time,required"`
	// A user-defined tag that can be included in incoming requests and used to trigger
	// a request rate condition
	UserDefinedTag param.Opt[string] `json:"user_defined_tag,omitzero"`
	// Possible HTTP request methods that can trigger a request rate condition
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethods []string `json:"http_methods,omitzero"`
	// A list of source IPs that can trigger a request rate condition
	IPs []string `json:"ips,omitzero" format:"ipvanyaddress"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionRequestRate) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionRequestRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionRequestRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match a response header
//
// The properties Header, Value are required.
type DomainCustomRuleNewParamsConditionResponseHeader struct {
	// The response header name
	Header string `json:"header,required"`
	// The response header value
	Value string `json:"value,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// The type of matching condition for header and value.
	//
	// Any of "Exact", "Contains".
	MatchType string `json:"match_type,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionResponseHeader) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionResponseHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionResponseHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainCustomRuleNewParamsConditionResponseHeader](
		"match_type", "Exact", "Contains",
	)
}

// Match when a response header is present
//
// The property Header is required.
type DomainCustomRuleNewParamsConditionResponseHeaderExists struct {
	// The response header name
	Header string `json:"header,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionResponseHeaderExists) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionResponseHeaderExists
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionResponseHeaderExists) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the number of dynamic page requests made in a WAAP session
//
// The property RequestCount is required.
type DomainCustomRuleNewParamsConditionSessionRequestCount struct {
	// The number of dynamic requests in the session
	RequestCount int64 `json:"request_count,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionSessionRequestCount) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionSessionRequestCount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionSessionRequestCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Matches requests based on specified tags
//
// The property Tags is required.
type DomainCustomRuleNewParamsConditionTags struct {
	// A list of tags to match against the request tags
	Tags []string `json:"tags,omitzero,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionTags) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request URL
//
// The property URL is required.
type DomainCustomRuleNewParamsConditionURL struct {
	// The pattern to match against the request URL. Constraints depend on
	// `match_type`:
	//
	//   - **Exact/Contains**: plain text matching (e.g., `/admin`, must comply with
	//     `^[\w!\$~:#\[\]@\(\)\\*\+,=\/\-\.\%]+$`).
	//   - **Regex**: a valid regular expression (e.g., `^/upload(/\d+)?/\w+`).
	//     Lookahead/lookbehind constructs are forbidden.
	URL string `json:"url,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// The type of matching condition.
	//
	// Any of "Exact", "Contains", "Regex".
	MatchType string `json:"match_type,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionURL) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainCustomRuleNewParamsConditionURL](
		"match_type", "Exact", "Contains", "Regex",
	)
}

// Match the user agent making the request
//
// The property UserAgent is required.
type DomainCustomRuleNewParamsConditionUserAgent struct {
	// The user agent value to match
	UserAgent string `json:"user_agent,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// The type of matching condition.
	//
	// Any of "Exact", "Contains".
	MatchType string `json:"match_type,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionUserAgent) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionUserAgent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionUserAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainCustomRuleNewParamsConditionUserAgent](
		"match_type", "Exact", "Contains",
	)
}

// Matches requests based on user-defined tags
//
// The property Tags is required.
type DomainCustomRuleNewParamsConditionUserDefinedTags struct {
	// A list of user-defined tags to match against the request tags
	Tags []string `json:"tags,omitzero,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleNewParamsConditionUserDefinedTags) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleNewParamsConditionUserDefinedTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleNewParamsConditionUserDefinedTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainCustomRuleUpdateParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	// The description assigned to the rule
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether or not the rule is enabled
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// The name assigned to the rule
	Name param.Opt[string] `json:"name,omitzero"`
	// The action that a WAAP rule takes when triggered.
	Action DomainCustomRuleUpdateParamsAction `json:"action,omitzero"`
	// The conditions required for the WAAP engine to trigger the rule. Rules may have
	// between 1 and 5 conditions. All conditions must pass for the rule to trigger
	Conditions []DomainCustomRuleUpdateParamsCondition `json:"conditions,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action that a WAAP rule takes when triggered.
type DomainCustomRuleUpdateParamsAction struct {
	// The WAAP allowed the request
	Allow any `json:"allow,omitzero"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block DomainCustomRuleUpdateParamsActionBlock `json:"block,omitzero"`
	// The WAAP presented the user with a captcha
	Captcha any `json:"captcha,omitzero"`
	// The WAAP performed automatic browser validation
	Handshake any `json:"handshake,omitzero"`
	// The WAAP monitored the request but took no action
	Monitor any `json:"monitor,omitzero"`
	// WAAP tag action gets a list of tags to tag the request scope with
	Tag DomainCustomRuleUpdateParamsActionTag `json:"tag,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WAAP block action behavior could be configured with response status code and
// action duration.
type DomainCustomRuleUpdateParamsActionBlock struct {
	// How long a rule's block action will apply to subsequent requests. Can be
	// specified in seconds or by using a numeral followed by 's', 'm', 'h', or 'd' to
	// represent time format (seconds, minutes, hours, or days). Empty time intervals
	// are not allowed.
	ActionDuration param.Opt[string] `json:"action_duration,omitzero"`
	// A custom HTTP status code that the WAAP returns if a rule blocks a request
	//
	// Any of 403, 405, 418, 429.
	StatusCode int64 `json:"status_code,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsActionBlock) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsActionBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsActionBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainCustomRuleUpdateParamsActionBlock](
		"status_code", 403, 405, 418, 429,
	)
}

// WAAP tag action gets a list of tags to tag the request scope with
//
// The property Tags is required.
type DomainCustomRuleUpdateParamsActionTag struct {
	// The list of user defined tags to tag the request with
	Tags []string `json:"tags,omitzero,required"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsActionTag) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsActionTag
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsActionTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The criteria of an incoming web request and the models of the various values
// those criteria can take
type DomainCustomRuleUpdateParamsCondition struct {
	// Match the requested Content-Type
	ContentType DomainCustomRuleUpdateParamsConditionContentType `json:"content_type,omitzero"`
	// Match the country that the request originated from
	Country DomainCustomRuleUpdateParamsConditionCountry `json:"country,omitzero"`
	// Match the incoming file extension
	FileExtension DomainCustomRuleUpdateParamsConditionFileExtension `json:"file_extension,omitzero"`
	// Match an incoming request header
	Header DomainCustomRuleUpdateParamsConditionHeader `json:"header,omitzero"`
	// Match when an incoming request header is present
	HeaderExists DomainCustomRuleUpdateParamsConditionHeaderExists `json:"header_exists,omitzero"`
	// Match the incoming HTTP method
	HTTPMethod DomainCustomRuleUpdateParamsConditionHTTPMethod `json:"http_method,omitzero"`
	// Match the incoming request against a single IP address
	IP DomainCustomRuleUpdateParamsConditionIP `json:"ip,omitzero"`
	// Match the incoming request against an IP range
	IPRange DomainCustomRuleUpdateParamsConditionIPRange `json:"ip_range,omitzero"`
	// Match the organization the request originated from, as determined by a WHOIS
	// lookup of the requesting IP
	Organization DomainCustomRuleUpdateParamsConditionOrganization `json:"organization,omitzero"`
	// Match the type of organization that owns the IP address making an incoming
	// request
	OwnerTypes DomainCustomRuleUpdateParamsConditionOwnerTypes `json:"owner_types,omitzero"`
	// Match the rate at which requests come in that match certain conditions
	RequestRate DomainCustomRuleUpdateParamsConditionRequestRate `json:"request_rate,omitzero"`
	// Match a response header
	ResponseHeader DomainCustomRuleUpdateParamsConditionResponseHeader `json:"response_header,omitzero"`
	// Match when a response header is present
	ResponseHeaderExists DomainCustomRuleUpdateParamsConditionResponseHeaderExists `json:"response_header_exists,omitzero"`
	// Match the number of dynamic page requests made in a WAAP session
	SessionRequestCount DomainCustomRuleUpdateParamsConditionSessionRequestCount `json:"session_request_count,omitzero"`
	// Matches requests based on specified tags
	Tags DomainCustomRuleUpdateParamsConditionTags `json:"tags,omitzero"`
	// Match the incoming request URL
	URL DomainCustomRuleUpdateParamsConditionURL `json:"url,omitzero"`
	// Match the user agent making the request
	UserAgent DomainCustomRuleUpdateParamsConditionUserAgent `json:"user_agent,omitzero"`
	// Matches requests based on user-defined tags
	UserDefinedTags DomainCustomRuleUpdateParamsConditionUserDefinedTags `json:"user_defined_tags,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsCondition) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the requested Content-Type
//
// The property ContentType is required.
type DomainCustomRuleUpdateParamsConditionContentType struct {
	// The list of content types to match against
	ContentType []string `json:"content_type,omitzero,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionContentType) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionContentType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionContentType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the country that the request originated from
//
// The property CountryCode is required.
type DomainCustomRuleUpdateParamsConditionCountry struct {
	// A list of ISO 3166-1 alpha-2 formatted strings representing the countries to
	// match against
	CountryCode []string `json:"country_code,omitzero,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionCountry) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionCountry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming file extension
//
// The property FileExtension is required.
type DomainCustomRuleUpdateParamsConditionFileExtension struct {
	// The list of file extensions to match against
	FileExtension []string `json:"file_extension,omitzero,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionFileExtension) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionFileExtension
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionFileExtension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match an incoming request header
//
// The properties Header, Value are required.
type DomainCustomRuleUpdateParamsConditionHeader struct {
	// The request header name
	Header string `json:"header,required"`
	// The request header value
	Value string `json:"value,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// The type of matching condition for header and value.
	//
	// Any of "Exact", "Contains".
	MatchType string `json:"match_type,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionHeader) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainCustomRuleUpdateParamsConditionHeader](
		"match_type", "Exact", "Contains",
	)
}

// Match when an incoming request header is present
//
// The property Header is required.
type DomainCustomRuleUpdateParamsConditionHeaderExists struct {
	// The request header name
	Header string `json:"header,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionHeaderExists) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionHeaderExists
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionHeaderExists) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming HTTP method
//
// The property HTTPMethod is required.
type DomainCustomRuleUpdateParamsConditionHTTPMethod struct {
	// HTTP methods of a request
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod string `json:"http_method,omitzero,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionHTTPMethod) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionHTTPMethod
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionHTTPMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainCustomRuleUpdateParamsConditionHTTPMethod](
		"http_method", "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT", "TRACE",
	)
}

// Match the incoming request against a single IP address
//
// The property IPAddress is required.
type DomainCustomRuleUpdateParamsConditionIP struct {
	// A single IPv4 or IPv6 address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionIP) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against an IP range
//
// The properties LowerBound, UpperBound are required.
type DomainCustomRuleUpdateParamsConditionIPRange struct {
	// The lower bound IPv4 or IPv6 address to match against
	LowerBound string `json:"lower_bound,required" format:"ipvanyaddress"`
	// The upper bound IPv4 or IPv6 address to match against
	UpperBound string `json:"upper_bound,required" format:"ipvanyaddress"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionIPRange) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionIPRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionIPRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the organization the request originated from, as determined by a WHOIS
// lookup of the requesting IP
//
// The property Organization is required.
type DomainCustomRuleUpdateParamsConditionOrganization struct {
	// The organization to match against
	Organization string `json:"organization,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionOrganization) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionOrganization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the type of organization that owns the IP address making an incoming
// request
type DomainCustomRuleUpdateParamsConditionOwnerTypes struct {
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// Match the type of organization that owns the IP address making an incoming
	// request
	//
	// Any of "COMMERCIAL", "EDUCATIONAL", "GOVERNMENT", "HOSTING_SERVICES", "ISP",
	// "MOBILE_NETWORK", "NETWORK", "RESERVED".
	OwnerTypes []string `json:"owner_types,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionOwnerTypes) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionOwnerTypes
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionOwnerTypes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the rate at which requests come in that match certain conditions
//
// The properties PathPattern, Requests, Time are required.
type DomainCustomRuleUpdateParamsConditionRequestRate struct {
	// A regular expression matching the URL path of the incoming request
	PathPattern string `json:"path_pattern,required"`
	// The number of incoming requests over the given time that can trigger a request
	// rate condition
	Requests int64 `json:"requests,required"`
	// The number of seconds that the WAAP measures incoming requests over before
	// triggering a request rate condition
	Time int64 `json:"time,required"`
	// A user-defined tag that can be included in incoming requests and used to trigger
	// a request rate condition
	UserDefinedTag param.Opt[string] `json:"user_defined_tag,omitzero"`
	// Possible HTTP request methods that can trigger a request rate condition
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethods []string `json:"http_methods,omitzero"`
	// A list of source IPs that can trigger a request rate condition
	IPs []string `json:"ips,omitzero" format:"ipvanyaddress"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionRequestRate) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionRequestRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionRequestRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match a response header
//
// The properties Header, Value are required.
type DomainCustomRuleUpdateParamsConditionResponseHeader struct {
	// The response header name
	Header string `json:"header,required"`
	// The response header value
	Value string `json:"value,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// The type of matching condition for header and value.
	//
	// Any of "Exact", "Contains".
	MatchType string `json:"match_type,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionResponseHeader) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionResponseHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionResponseHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainCustomRuleUpdateParamsConditionResponseHeader](
		"match_type", "Exact", "Contains",
	)
}

// Match when a response header is present
//
// The property Header is required.
type DomainCustomRuleUpdateParamsConditionResponseHeaderExists struct {
	// The response header name
	Header string `json:"header,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionResponseHeaderExists) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionResponseHeaderExists
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionResponseHeaderExists) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the number of dynamic page requests made in a WAAP session
//
// The property RequestCount is required.
type DomainCustomRuleUpdateParamsConditionSessionRequestCount struct {
	// The number of dynamic requests in the session
	RequestCount int64 `json:"request_count,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionSessionRequestCount) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionSessionRequestCount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionSessionRequestCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Matches requests based on specified tags
//
// The property Tags is required.
type DomainCustomRuleUpdateParamsConditionTags struct {
	// A list of tags to match against the request tags
	Tags []string `json:"tags,omitzero,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionTags) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request URL
//
// The property URL is required.
type DomainCustomRuleUpdateParamsConditionURL struct {
	// The pattern to match against the request URL. Constraints depend on
	// `match_type`:
	//
	//   - **Exact/Contains**: plain text matching (e.g., `/admin`, must comply with
	//     `^[\w!\$~:#\[\]@\(\)\\*\+,=\/\-\.\%]+$`).
	//   - **Regex**: a valid regular expression (e.g., `^/upload(/\d+)?/\w+`).
	//     Lookahead/lookbehind constructs are forbidden.
	URL string `json:"url,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// The type of matching condition.
	//
	// Any of "Exact", "Contains", "Regex".
	MatchType string `json:"match_type,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionURL) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainCustomRuleUpdateParamsConditionURL](
		"match_type", "Exact", "Contains", "Regex",
	)
}

// Match the user agent making the request
//
// The property UserAgent is required.
type DomainCustomRuleUpdateParamsConditionUserAgent struct {
	// The user agent value to match
	UserAgent string `json:"user_agent,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// The type of matching condition.
	//
	// Any of "Exact", "Contains".
	MatchType string `json:"match_type,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionUserAgent) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionUserAgent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionUserAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainCustomRuleUpdateParamsConditionUserAgent](
		"match_type", "Exact", "Contains",
	)
}

// Matches requests based on user-defined tags
//
// The property Tags is required.
type DomainCustomRuleUpdateParamsConditionUserDefinedTags struct {
	// A list of user-defined tags to match against the request tags
	Tags []string `json:"tags,omitzero,required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainCustomRuleUpdateParamsConditionUserDefinedTags) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleUpdateParamsConditionUserDefinedTags
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleUpdateParamsConditionUserDefinedTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainCustomRuleListParams struct {
	// Filter rules based on their description. Supports '\*' as a wildcard character.
	Description param.Opt[string] `query:"description,omitzero" json:"-"`
	// Filter rules based on their active status
	Enabled param.Opt[bool] `query:"enabled,omitzero" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter rules based on their name. Supports '\*' as a wildcard character.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Determine the field to order results by
	//
	// Any of "id", "name", "description", "enabled", "action", "-id", "-name",
	// "-description", "-enabled", "-action".
	Ordering DomainCustomRuleListParamsOrdering `query:"ordering,omitzero" json:"-"`
	// Filter to refine results by specific actions
	//
	// Any of "allow", "block", "captcha", "handshake", "monitor", "tag".
	Action DomainCustomRuleListParamsAction `query:"action,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainCustomRuleListParams]'s query parameters as
// `url.Values`.
func (r DomainCustomRuleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter to refine results by specific actions
type DomainCustomRuleListParamsAction string

const (
	DomainCustomRuleListParamsActionAllow     DomainCustomRuleListParamsAction = "allow"
	DomainCustomRuleListParamsActionBlock     DomainCustomRuleListParamsAction = "block"
	DomainCustomRuleListParamsActionCaptcha   DomainCustomRuleListParamsAction = "captcha"
	DomainCustomRuleListParamsActionHandshake DomainCustomRuleListParamsAction = "handshake"
	DomainCustomRuleListParamsActionMonitor   DomainCustomRuleListParamsAction = "monitor"
	DomainCustomRuleListParamsActionTag       DomainCustomRuleListParamsAction = "tag"
)

// Determine the field to order results by
type DomainCustomRuleListParamsOrdering string

const (
	DomainCustomRuleListParamsOrderingID               DomainCustomRuleListParamsOrdering = "id"
	DomainCustomRuleListParamsOrderingName             DomainCustomRuleListParamsOrdering = "name"
	DomainCustomRuleListParamsOrderingDescription      DomainCustomRuleListParamsOrdering = "description"
	DomainCustomRuleListParamsOrderingEnabled          DomainCustomRuleListParamsOrdering = "enabled"
	DomainCustomRuleListParamsOrderingAction           DomainCustomRuleListParamsOrdering = "action"
	DomainCustomRuleListParamsOrderingMinusID          DomainCustomRuleListParamsOrdering = "-id"
	DomainCustomRuleListParamsOrderingMinusName        DomainCustomRuleListParamsOrdering = "-name"
	DomainCustomRuleListParamsOrderingMinusDescription DomainCustomRuleListParamsOrdering = "-description"
	DomainCustomRuleListParamsOrderingMinusEnabled     DomainCustomRuleListParamsOrdering = "-enabled"
	DomainCustomRuleListParamsOrderingMinusAction      DomainCustomRuleListParamsOrdering = "-action"
)

type DomainCustomRuleDeleteParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}

type DomainCustomRuleDeleteMultipleParams struct {
	// The IDs of the rules to delete
	RuleIDs []int64 `json:"rule_ids,omitzero,required"`
	paramObj
}

func (r DomainCustomRuleDeleteMultipleParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainCustomRuleDeleteMultipleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainCustomRuleDeleteMultipleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainCustomRuleGetParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}

type DomainCustomRuleToggleParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	// The custom rule ID
	RuleID int64 `path:"rule_id,required" json:"-"`
	paramObj
}

// Enable or disable a custom rule
type DomainCustomRuleToggleParamsAction string

const (
	DomainCustomRuleToggleParamsActionEnable  DomainCustomRuleToggleParamsAction = "enable"
	DomainCustomRuleToggleParamsActionDisable DomainCustomRuleToggleParamsAction = "disable"
)
