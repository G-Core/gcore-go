// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// WAAP domains enable Web Application and API Protection for monitoring and
// defending web applications against security threats.
//
// DomainService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainService] method instead.
type DomainService struct {
	Options         []option.RequestOption
	Policies        DomainPolicyService
	Settings        DomainSettingService
	APIPaths        DomainAPIPathService
	APIPathGroups   DomainAPIPathGroupService
	APIDiscovery    DomainAPIDiscoveryService
	Insights        DomainInsightService
	InsightSilences DomainInsightSilenceService
	Statistics      DomainStatisticService
	CustomRules     DomainCustomRuleService
	FirewallRules   DomainFirewallRuleService
	AdvancedRules   DomainAdvancedRuleService
}

// NewDomainService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomainService(opts ...option.RequestOption) (r DomainService) {
	r = DomainService{}
	r.Options = opts
	r.Policies = NewDomainPolicyService(opts...)
	r.Settings = NewDomainSettingService(opts...)
	r.APIPaths = NewDomainAPIPathService(opts...)
	r.APIPathGroups = NewDomainAPIPathGroupService(opts...)
	r.APIDiscovery = NewDomainAPIDiscoveryService(opts...)
	r.Insights = NewDomainInsightService(opts...)
	r.InsightSilences = NewDomainInsightSilenceService(opts...)
	r.Statistics = NewDomainStatisticService(opts...)
	r.CustomRules = NewDomainCustomRuleService(opts...)
	r.FirewallRules = NewDomainFirewallRuleService(opts...)
	r.AdvancedRules = NewDomainAdvancedRuleService(opts...)
	return
}

// Update Domain
func (r *DomainService) Update(ctx context.Context, domainID int64, body DomainUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// Retrieve a list of domains associated with the client
func (r *DomainService) List(ctx context.Context, query DomainListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapSummaryDomain], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "waap/v1/domains"
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

// Retrieve a list of domains associated with the client
func (r *DomainService) ListAutoPaging(ctx context.Context, query DomainListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapSummaryDomain] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Delete an inactive domain by ID. Only domains with status 'bypass' can be
// deleted.
func (r *DomainService) Delete(ctx context.Context, domainID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve detailed information about a specific domain
func (r *DomainService) Get(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *WaapDetailedDomain, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve all rule sets linked to a particular domain
func (r *DomainService) ListRuleSets(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *[]WaapRuleSet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/rule-sets", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Represents a WAAP domain, serving as a singular unit within the WAAP service.
//
// Each domain functions autonomously, possessing its own set of rules and
// configurations to manage web application firewall settings and behaviors.
type WaapDetailedDomain struct {
	// The domain ID
	ID int64 `json:"id" api:"required"`
	// The date and time the domain was created in ISO 8601 format
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The ID of the custom page set
	CustomPageSet int64 `json:"custom_page_set" api:"required"`
	// The domain name
	Name string `json:"name" api:"required"`
	// The different statuses a domain can have
	//
	// Any of "active", "bypass", "monitor", "locked".
	Status WaapDetailedDomainStatus `json:"status" api:"required"`
	// Domain level quotas
	Quotas map[string]WaapDetailedDomainQuota `json:"quotas" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		CustomPageSet respjson.Field
		Name          respjson.Field
		Status        respjson.Field
		Quotas        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDetailedDomain) RawJSON() string { return r.JSON.raw }
func (r *WaapDetailedDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The different statuses a domain can have
type WaapDetailedDomainStatus string

const (
	WaapDetailedDomainStatusActive  WaapDetailedDomainStatus = "active"
	WaapDetailedDomainStatusBypass  WaapDetailedDomainStatus = "bypass"
	WaapDetailedDomainStatusMonitor WaapDetailedDomainStatus = "monitor"
	WaapDetailedDomainStatusLocked  WaapDetailedDomainStatus = "locked"
)

type WaapDetailedDomainQuota struct {
	// The maximum allowed number of this resource
	Allowed int64 `json:"allowed" api:"required"`
	// The current number of this resource
	Current int64 `json:"current" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allowed     respjson.Field
		Current     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDetailedDomainQuota) RawJSON() string { return r.JSON.raw }
func (r *WaapDetailedDomainQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API settings of a domain
type WaapDomainAPISettings struct {
	// The API URLs for a domain. If your domain has a common base URL for all API
	// paths, it can be set here
	APIURLs []string `json:"api_urls"`
	// Indicates if the domain is an API domain. All requests to an API domain are
	// treated as API requests. If this is set to true then the `api_urls` field is
	// ignored.
	IsAPI bool `json:"is_api"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIURLs     respjson.Field
		IsAPI       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDomainAPISettings) RawJSON() string { return r.JSON.raw }
func (r *WaapDomainAPISettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DDoS settings for a domain.
type WaapDomainDDOSSettings struct {
	// The burst threshold detects sudden rises in traffic. If it is met and the number
	// of requests is at least five times the last 2-second interval, DDoS protection
	// will activate. Default is 1000.
	BurstThreshold int64 `json:"burst_threshold"`
	// The global threshold is responsible for identifying DDoS attacks with a slow
	// rise in traffic. If the threshold is met and the current number of requests is
	// at least double that of the previous 10-second window, DDoS protection will
	// activate. Default is 5000.
	GlobalThreshold int64 `json:"global_threshold"`
	// The sub-second threshold protects WAAP servers against attacks from traffic
	// bursts. When this threshold is reached, the DDoS mode will activate on the
	// affected WAAP server, not the whole WAAP cluster. Default is 50.
	SubSecondThreshold int64 `json:"sub_second_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BurstThreshold     respjson.Field
		GlobalThreshold    respjson.Field
		SubSecondThreshold respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDomainDDOSSettings) RawJSON() string { return r.JSON.raw }
func (r *WaapDomainDDOSSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for a domain.
type WaapDomainSettingsModel struct {
	// API settings of a domain
	API WaapDomainAPISettings `json:"api" api:"required"`
	// DDoS settings for a domain.
	DDOS WaapDomainDDOSSettings `json:"ddos" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		API         respjson.Field
		DDOS        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDomainSettingsModel) RawJSON() string { return r.JSON.raw }
func (r *WaapDomainSettingsModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request summary used when displaying a list of requests
type WaapRequestSummary struct {
	// Request's unique id
	ID string `json:"id" api:"required"`
	// Action of the triggered rule
	Action string `json:"action" api:"required"`
	// Client's IP address.
	ClientIP string `json:"client_ip" api:"required"`
	// Country code
	Country string `json:"country" api:"required"`
	// The decision made for processing the request through the WAAP.
	//
	// Any of "passed", "allowed", "monitored", "blocked", "".
	Decision WaapRequestSummaryDecision `json:"decision" api:"required"`
	// Domain name
	Domain string `json:"domain" api:"required"`
	// Domain ID
	DomainID int64 `json:"domain_id" api:"required"`
	// HTTP method
	Method string `json:"method" api:"required"`
	// An optional action that may be applied in addition to the primary decision.
	//
	// Any of "captcha", "challenge", "".
	OptionalAction WaapRequestSummaryOptionalAction `json:"optional_action" api:"required"`
	// Organization
	Organization string `json:"organization" api:"required"`
	// Request path
	Path string `json:"path" api:"required"`
	// The reference ID to a sanction that was given to a user.
	ReferenceID string `json:"reference_id" api:"required"`
	// The UNIX timestamp in ms of the date a set of traffic counters was recorded
	RequestTime int64 `json:"request_time" api:"required"`
	// Any of "passed", "blocked", "suppressed", "".
	Result WaapRequestSummaryResult `json:"result" api:"required"`
	// The ID of the triggered rule.
	RuleID string `json:"rule_id" api:"required"`
	// Name of the triggered rule
	RuleName string `json:"rule_name" api:"required"`
	// Status code for http request
	StatusCode int64 `json:"status_code" api:"required"`
	// Comma separated list of traffic types.
	TrafficTypes string `json:"traffic_types" api:"required"`
	// User agent
	UserAgent string `json:"user_agent" api:"required"`
	// Client from parsed User agent header
	UserAgentClient string `json:"user_agent_client" api:"required"`
	// The session ID associated with the request.
	SessionID string `json:"session_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Action          respjson.Field
		ClientIP        respjson.Field
		Country         respjson.Field
		Decision        respjson.Field
		Domain          respjson.Field
		DomainID        respjson.Field
		Method          respjson.Field
		OptionalAction  respjson.Field
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
		SessionID       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRequestSummary) RawJSON() string { return r.JSON.raw }
func (r *WaapRequestSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The decision made for processing the request through the WAAP.
type WaapRequestSummaryDecision string

const (
	WaapRequestSummaryDecisionPassed    WaapRequestSummaryDecision = "passed"
	WaapRequestSummaryDecisionAllowed   WaapRequestSummaryDecision = "allowed"
	WaapRequestSummaryDecisionMonitored WaapRequestSummaryDecision = "monitored"
	WaapRequestSummaryDecisionBlocked   WaapRequestSummaryDecision = "blocked"
	WaapRequestSummaryDecisionEmpty     WaapRequestSummaryDecision = ""
)

// An optional action that may be applied in addition to the primary decision.
type WaapRequestSummaryOptionalAction string

const (
	WaapRequestSummaryOptionalActionCaptcha   WaapRequestSummaryOptionalAction = "captcha"
	WaapRequestSummaryOptionalActionChallenge WaapRequestSummaryOptionalAction = "challenge"
	WaapRequestSummaryOptionalActionEmpty     WaapRequestSummaryOptionalAction = ""
)

type WaapRequestSummaryResult string

const (
	WaapRequestSummaryResultPassed     WaapRequestSummaryResult = "passed"
	WaapRequestSummaryResultBlocked    WaapRequestSummaryResult = "blocked"
	WaapRequestSummaryResultSuppressed WaapRequestSummaryResult = "suppressed"
	WaapRequestSummaryResultEmpty      WaapRequestSummaryResult = ""
)

// Represents a custom rule set.
type WaapRuleSet struct {
	// Identifier of the rule set.
	ID int64 `json:"id" api:"required"`
	// Detailed description of the rule set.
	Description string `json:"description" api:"required"`
	// Indicates if the rule set is currently active.
	IsActive bool `json:"is_active" api:"required"`
	// Name of the rule set.
	Name string `json:"name" api:"required"`
	// Collection of tags associated with the rule set.
	Tags []WaapRuleSetTag `json:"tags" api:"required"`
	// The resource slug associated with the rule set.
	ResourceSlug string            `json:"resource_slug" api:"nullable"`
	Rules        []WaapRuleSetRule `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Description  respjson.Field
		IsActive     respjson.Field
		Name         respjson.Field
		Tags         respjson.Field
		ResourceSlug respjson.Field
		Rules        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRuleSet) RawJSON() string { return r.JSON.raw }
func (r *WaapRuleSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single tag associated with a rule set.
type WaapRuleSetTag struct {
	// Identifier of the tag.
	ID int64 `json:"id" api:"required"`
	// Detailed description of the tag.
	Description string `json:"description" api:"required"`
	// Name of the tag.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRuleSetTag) RawJSON() string { return r.JSON.raw }
func (r *WaapRuleSetTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a configurable WAAP security rule, also known as a policy.
type WaapRuleSetRule struct {
	// Unique identifier for the security rule
	ID string `json:"id" api:"required"`
	// Specifies the action taken by the WAAP upon rule activation
	//
	// Any of "Allow", "Block", "Captcha", "Gateway", "Handshake", "Monitor",
	// "Composite".
	Action string `json:"action" api:"required"`
	// Detailed description of the security rule
	Description string `json:"description" api:"required"`
	// The rule set group name to which the rule belongs
	Group string `json:"group" api:"required"`
	// Indicates if the security rule is active
	Mode bool `json:"mode" api:"required"`
	// Name of the security rule
	Name string `json:"name" api:"required"`
	// Identifier of the rule set to which the rule belongs
	RuleSetID int64 `json:"rule_set_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Action      respjson.Field
		Description respjson.Field
		Group       respjson.Field
		Mode        respjson.Field
		Name        respjson.Field
		RuleSetID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRuleSetRule) RawJSON() string { return r.JSON.raw }
func (r *WaapRuleSetRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a WAAP domain when getting a list of domains.
type WaapSummaryDomain struct {
	// The domain ID
	ID int64 `json:"id" api:"required"`
	// The date and time the domain was created in ISO 8601 format
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The ID of the custom page set
	CustomPageSet int64 `json:"custom_page_set" api:"required"`
	// The domain name
	Name string `json:"name" api:"required"`
	// The different statuses a domain can have
	//
	// Any of "active", "bypass", "monitor", "locked".
	Status WaapSummaryDomainStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		CustomPageSet respjson.Field
		Name          respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapSummaryDomain) RawJSON() string { return r.JSON.raw }
func (r *WaapSummaryDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The different statuses a domain can have
type WaapSummaryDomainStatus string

const (
	WaapSummaryDomainStatusActive  WaapSummaryDomainStatus = "active"
	WaapSummaryDomainStatusBypass  WaapSummaryDomainStatus = "bypass"
	WaapSummaryDomainStatusMonitor WaapSummaryDomainStatus = "monitor"
	WaapSummaryDomainStatusLocked  WaapSummaryDomainStatus = "locked"
)

// Represents the traffic metrics for a domain at a given time window
type WaapTrafficMetrics struct {
	// UNIX timestamp indicating when the traffic data was recorded
	Timestamp int64 `json:"timestamp" api:"required"`
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

type DomainUpdateParams struct {
	// The current status of the domain
	//
	// Any of "active", "monitor".
	Status DomainUpdateParamsStatus `json:"status,omitzero" api:"required"`
	paramObj
}

func (r DomainUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the domain
type DomainUpdateParamsStatus string

const (
	DomainUpdateParamsStatusActive  DomainUpdateParamsStatus = "active"
	DomainUpdateParamsStatusMonitor DomainUpdateParamsStatus = "monitor"
)

type DomainListParams struct {
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter domains based on the domain name. Supports '\*' as a wildcard character
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter domains based on their IDs
	IDs []int64 `query:"ids,omitzero" json:"-"`
	// Sort the response by given field.
	//
	// Any of "id", "name", "status", "created_at", "-id", "-name", "-status",
	// "-created_at".
	Ordering DomainListParamsOrdering `query:"ordering,omitzero" json:"-"`
	// Filter domains based on the domain status
	//
	// Any of "active", "bypass", "monitor", "locked".
	Status DomainListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainListParams]'s query parameters as `url.Values`.
func (r DomainListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort the response by given field.
type DomainListParamsOrdering string

const (
	DomainListParamsOrderingID             DomainListParamsOrdering = "id"
	DomainListParamsOrderingName           DomainListParamsOrdering = "name"
	DomainListParamsOrderingStatus         DomainListParamsOrdering = "status"
	DomainListParamsOrderingCreatedAt      DomainListParamsOrdering = "created_at"
	DomainListParamsOrderingMinusID        DomainListParamsOrdering = "-id"
	DomainListParamsOrderingMinusName      DomainListParamsOrdering = "-name"
	DomainListParamsOrderingMinusStatus    DomainListParamsOrdering = "-status"
	DomainListParamsOrderingMinusCreatedAt DomainListParamsOrdering = "-created_at"
)

// Filter domains based on the domain status
type DomainListParamsStatus string

const (
	DomainListParamsStatusActive  DomainListParamsStatus = "active"
	DomainListParamsStatusBypass  DomainListParamsStatus = "bypass"
	DomainListParamsStatusMonitor DomainListParamsStatus = "monitor"
	DomainListParamsStatusLocked  DomainListParamsStatus = "locked"
)
