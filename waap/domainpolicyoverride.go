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

// DomainPolicyOverrideService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainPolicyOverrideService] method instead.
type DomainPolicyOverrideService struct {
	Options []option.RequestOption
}

// NewDomainPolicyOverrideService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDomainPolicyOverrideService(opts ...option.RequestOption) (r DomainPolicyOverrideService) {
	r = DomainPolicyOverrideService{}
	r.Options = opts
	return
}

// Create a policy override
func (r *DomainPolicyOverrideService) New(ctx context.Context, domainID int64, body DomainPolicyOverrideNewParams, opts ...option.RequestOption) (res *WaapPolicyOverride, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/default-rule-overrides", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a policy override
func (r *DomainPolicyOverrideService) Update(ctx context.Context, policyOverrideID int64, params DomainPolicyOverrideUpdateParams, opts ...option.RequestOption) (res *WaapPolicyOverride, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/default-rule-overrides/%v", params.DomainID, policyOverrideID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieve policy overrides
func (r *DomainPolicyOverrideService) List(ctx context.Context, domainID int64, query DomainPolicyOverrideListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapPolicyOverride], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/default-rule-overrides", domainID)
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

// Retrieve policy overrides
func (r *DomainPolicyOverrideService) ListAutoPaging(ctx context.Context, domainID int64, query DomainPolicyOverrideListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapPolicyOverride] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, domainID, query, opts...))
}

// Delete a policy override
func (r *DomainPolicyOverrideService) Delete(ctx context.Context, policyOverrideID int64, body DomainPolicyOverrideDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/default-rule-overrides/%v", body.DomainID, policyOverrideID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve a policy override
func (r *DomainPolicyOverrideService) Get(ctx context.Context, policyOverrideID int64, query DomainPolicyOverrideGetParams, opts ...option.RequestOption) (res *WaapPolicyOverride, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/default-rule-overrides/%v", query.DomainID, policyOverrideID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List OWASP Top Threats targets selectable as a policy override target
func (r *DomainPolicyOverrideService) ListOwaspTopThreatsTargets(ctx context.Context, domainID int64, query DomainPolicyOverrideListOwaspTopThreatsTargetsParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapOwaspTopThreatsTarget], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/policy-override-targets/owasp-top-threats", domainID)
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

// List OWASP Top Threats targets selectable as a policy override target
func (r *DomainPolicyOverrideService) ListOwaspTopThreatsTargetsAutoPaging(ctx context.Context, domainID int64, query DomainPolicyOverrideListOwaspTopThreatsTargetsParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapOwaspTopThreatsTarget] {
	return pagination.NewOffsetPageAutoPager(r.ListOwaspTopThreatsTargets(ctx, domainID, query, opts...))
}

// A default OWASP Top Threats target that can be selected as the target of a
// policy override, so this specific managed detection stops blocking matching
// traffic while the rest of OWASP Top Threats protection stays unchanged.
type WaapOwaspTopThreatsTarget struct {
	// The rule's ID; use this as a `target_id` when creating a policy override
	ID int64 `json:"id" api:"required"`
	// The action this rule takes when it triggers
	Action string `json:"action" api:"required"`
	// The rule's display name
	Name string `json:"name" api:"required"`
	// The parent policy this rule is grouped under, if any
	Policy WaapOwaspTopThreatsTargetPolicy `json:"policy" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Action      respjson.Field
		Name        respjson.Field
		Policy      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapOwaspTopThreatsTarget) RawJSON() string { return r.JSON.raw }
func (r *WaapOwaspTopThreatsTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parent policy this rule is grouped under, if any
type WaapOwaspTopThreatsTargetPolicy struct {
	// The parent policy's ID
	ID int64 `json:"id" api:"required"`
	// The parent policy's name
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapOwaspTopThreatsTargetPolicy) RawJSON() string { return r.JSON.raw }
func (r *WaapOwaspTopThreatsTargetPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapPolicyOverride struct {
	// The policy override's ID
	ID int64 `json:"id" api:"required"`
	// The conditions required for the override to apply. Empty means the override
	// applies to every request on the owning domain that hits a linked target.
	Conditions []WaapPolicyOverrideCondition `json:"conditions" api:"required"`
	// When the override was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The domain this override belongs to
	DomainID int64 `json:"domain_id" api:"required"`
	// Whether the override is active
	Enabled bool `json:"enabled" api:"required"`
	// The policy override's display name
	Name string `json:"name" api:"required"`
	// Whether a linked target has been deleted, disabled, or otherwise become invalid.
	// A stale override is not published and does not suppress anything until it's
	// edited.
	Stale bool `json:"stale" api:"required"`
	// Whether targets are detectors or rules
	//
	// Any of "detector", "rule".
	TargetType WaapPolicyOverrideTargetType `json:"target_type" api:"required"`
	// The detectors or rules this override suppresses
	Targets []WaapPolicyOverrideTarget `json:"targets" api:"required"`
	// When the override was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional note explaining the override
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Conditions  respjson.Field
		CreatedAt   respjson.Field
		DomainID    respjson.Field
		Enabled     respjson.Field
		Name        respjson.Field
		Stale       respjson.Field
		TargetType  respjson.Field
		Targets     respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapPolicyOverride) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single policy override condition. Only one of the fields below may be set.
//
// Limited to the subset of `CustomRuleConditionScope` scopes that policy overrides
// support (Lime wire-format translation is reused from Custom Rules' own
// `app.models.customer_rules.lime_translation.ConditionMapper`).
type WaapPolicyOverrideCondition struct {
	// Match the country that the request originated from
	Country WaapPolicyOverrideConditionCountry `json:"country"`
	// Match an incoming request header
	Header WaapPolicyOverrideConditionHeader `json:"header"`
	// Match when an incoming request header is present
	HeaderExists WaapPolicyOverrideConditionHeaderExists `json:"header_exists"`
	// Match the incoming HTTP method
	HTTPMethod WaapPolicyOverrideConditionHTTPMethod `json:"http_method"`
	// Match the incoming request against a single IP address
	IP WaapPolicyOverrideConditionIP `json:"ip"`
	// Match the incoming request against an IP range
	IPRange WaapPolicyOverrideConditionIPRange `json:"ip_range"`
	// Match the JA3 TLS client fingerprint of the request
	Ja3 WaapPolicyOverrideConditionJa3 `json:"ja3"`
	// Match the JA4 TLS client fingerprint of the request
	Ja4 WaapPolicyOverrideConditionJa4 `json:"ja4"`
	// Match the organization the request originated from, as determined by a WHOIS
	// lookup of the requesting IP
	Organization WaapPolicyOverrideConditionOrganization `json:"organization"`
	// Match the incoming request URL
	URL WaapPolicyOverrideConditionURL `json:"url"`
	// Match the user agent making the request
	UserAgent WaapPolicyOverrideConditionUserAgent `json:"user_agent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country      respjson.Field
		Header       respjson.Field
		HeaderExists respjson.Field
		HTTPMethod   respjson.Field
		IP           respjson.Field
		IPRange      respjson.Field
		Ja3          respjson.Field
		Ja4          respjson.Field
		Organization respjson.Field
		URL          respjson.Field
		UserAgent    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapPolicyOverrideCondition) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverrideCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the country that the request originated from
type WaapPolicyOverrideConditionCountry struct {
	// A list of ISO 3166-1 alpha-2 formatted strings representing the countries to
	// match against
	CountryCode []string `json:"country_code" api:"required"`
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
func (r WaapPolicyOverrideConditionCountry) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverrideConditionCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match an incoming request header
type WaapPolicyOverrideConditionHeader struct {
	// The request header name
	Header string `json:"header" api:"required"`
	// The request header value
	Value string `json:"value" api:"required"`
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
func (r WaapPolicyOverrideConditionHeader) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverrideConditionHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match when an incoming request header is present
type WaapPolicyOverrideConditionHeaderExists struct {
	// The request header name
	Header string `json:"header" api:"required"`
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
func (r WaapPolicyOverrideConditionHeaderExists) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverrideConditionHeaderExists) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming HTTP method
type WaapPolicyOverrideConditionHTTPMethod struct {
	// HTTP methods of a request
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod string `json:"http_method" api:"required"`
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
func (r WaapPolicyOverrideConditionHTTPMethod) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverrideConditionHTTPMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against a single IP address
type WaapPolicyOverrideConditionIP struct {
	// A single IPv4 or IPv6 address
	IPAddress string `json:"ip_address" api:"required" format:"ipvanyaddress"`
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
func (r WaapPolicyOverrideConditionIP) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverrideConditionIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against an IP range
type WaapPolicyOverrideConditionIPRange struct {
	// The lower bound IPv4 or IPv6 address to match against
	LowerBound string `json:"lower_bound" api:"required" format:"ipvanyaddress"`
	// The upper bound IPv4 or IPv6 address to match against
	UpperBound string `json:"upper_bound" api:"required" format:"ipvanyaddress"`
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
func (r WaapPolicyOverrideConditionIPRange) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverrideConditionIPRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the JA3 TLS client fingerprint of the request
type WaapPolicyOverrideConditionJa3 struct {
	// A list of JA3 TLS client fingerprints to match against the request
	Ja3Fingerprints []string `json:"ja3_fingerprints" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ja3Fingerprints respjson.Field
		Negation        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapPolicyOverrideConditionJa3) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverrideConditionJa3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the JA4 TLS client fingerprint of the request
type WaapPolicyOverrideConditionJa4 struct {
	// A list of JA4 TLS client fingerprints to match against the request
	Ja4Fingerprints []string `json:"ja4_fingerprints" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation bool `json:"negation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ja4Fingerprints respjson.Field
		Negation        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapPolicyOverrideConditionJa4) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverrideConditionJa4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the organization the request originated from, as determined by a WHOIS
// lookup of the requesting IP
type WaapPolicyOverrideConditionOrganization struct {
	// The organization to match against
	Organization string `json:"organization" api:"required"`
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
func (r WaapPolicyOverrideConditionOrganization) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverrideConditionOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request URL
type WaapPolicyOverrideConditionURL struct {
	// The pattern to match against the request URL. Constraints depend on
	// `match_type`:
	//
	//   - **Exact/Contains**: plain text matching (e.g., `/admin`, must comply with
	//     `^[\w!\$~:#\[\]@\(\)*\+,=\/\-\.\%]+$`).
	//   - **Regex**: a valid regular expression (e.g., `^/upload(/\d+)?/\w+`).
	//     Lookahead/lookbehind constructs are forbidden.
	URL string `json:"url" api:"required"`
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
func (r WaapPolicyOverrideConditionURL) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverrideConditionURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the user agent making the request
type WaapPolicyOverrideConditionUserAgent struct {
	// The user agent value to match
	UserAgent string `json:"user_agent" api:"required"`
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
func (r WaapPolicyOverrideConditionUserAgent) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverrideConditionUserAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether targets are detectors or rules
type WaapPolicyOverrideTargetType string

const (
	WaapPolicyOverrideTargetTypeDetector WaapPolicyOverrideTargetType = "detector"
	WaapPolicyOverrideTargetTypeRule     WaapPolicyOverrideTargetType = "rule"
)

type WaapPolicyOverrideTarget struct {
	// The policy override target's ID
	ID int64 `json:"id" api:"required"`
	// The stable reference edge uses to match this target
	Ref string `json:"ref" api:"required"`
	// A point-in-time record of the target's details, retained for display after the
	// linked object is deleted
	Snapshot WaapPolicyOverrideTargetSnapshot `json:"snapshot" api:"required"`
	// The linked detector's ID, if `target_type` is detector
	DetectorID int64 `json:"detector_id" api:"nullable"`
	// The linked rule's ID, if `target_type` is rule
	RuleID int64 `json:"rule_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Ref         respjson.Field
		Snapshot    respjson.Field
		DetectorID  respjson.Field
		RuleID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapPolicyOverrideTarget) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverrideTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A point-in-time record of the target's details, retained for display after the
// linked object is deleted
type WaapPolicyOverrideTargetSnapshot struct {
	// The target's display name when it was linked
	Name string `json:"name" api:"required"`
	// The action the target took when it triggered, if `target_type` is rule
	Action string `json:"action" api:"nullable"`
	// The attack category the target detected, if `target_type` is rule
	AttackType string `json:"attack_type" api:"nullable"`
	// The target's rule category (policy or template), if `target_type` is rule
	Category string `json:"category" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Action      respjson.Field
		AttackType  respjson.Field
		Category    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapPolicyOverrideTargetSnapshot) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyOverrideTargetSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainPolicyOverrideNewParams struct {
	// Whether the override is active
	Enabled bool `json:"enabled" api:"required"`
	// The policy override's display name
	Name string `json:"name" api:"required"`
	// The detectors or rules this override suppresses
	TargetIDs []int64 `json:"target_ids,omitzero" api:"required"`
	// Whether `target_ids` are detector IDs or rule IDs
	//
	// Any of "detector", "rule".
	TargetType DomainPolicyOverrideNewParamsTargetType `json:"target_type,omitzero" api:"required"`
	// An optional note explaining the override
	Description param.Opt[string] `json:"description,omitzero"`
	// The ID of the request that prompted this override, if created from Event
	// Details. Preserves a snapshot of that request for audit purposes; has no effect
	// on override behavior and is not returned by this API.
	RequestID param.Opt[string] `json:"request_id,omitzero"`
	// The conditions required for the override to apply. An override may have up to 5
	// conditions; all must pass for the override to apply. Without conditions the
	// override applies to every request on the owning domain that hits a linked
	// target.
	Conditions []DomainPolicyOverrideNewParamsCondition `json:"conditions,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether `target_ids` are detector IDs or rule IDs
type DomainPolicyOverrideNewParamsTargetType string

const (
	DomainPolicyOverrideNewParamsTargetTypeDetector DomainPolicyOverrideNewParamsTargetType = "detector"
	DomainPolicyOverrideNewParamsTargetTypeRule     DomainPolicyOverrideNewParamsTargetType = "rule"
)

// A single policy override condition. Only one of the fields below may be set.
//
// Limited to the subset of `CustomRuleConditionScope` scopes that policy overrides
// support (Lime wire-format translation is reused from Custom Rules' own
// `app.models.customer_rules.lime_translation.ConditionMapper`).
type DomainPolicyOverrideNewParamsCondition struct {
	// Match the country that the request originated from
	Country DomainPolicyOverrideNewParamsConditionCountry `json:"country,omitzero"`
	// Match an incoming request header
	Header DomainPolicyOverrideNewParamsConditionHeader `json:"header,omitzero"`
	// Match when an incoming request header is present
	HeaderExists DomainPolicyOverrideNewParamsConditionHeaderExists `json:"header_exists,omitzero"`
	// Match the incoming HTTP method
	HTTPMethod DomainPolicyOverrideNewParamsConditionHTTPMethod `json:"http_method,omitzero"`
	// Match the incoming request against a single IP address
	IP DomainPolicyOverrideNewParamsConditionIP `json:"ip,omitzero"`
	// Match the incoming request against an IP range
	IPRange DomainPolicyOverrideNewParamsConditionIPRange `json:"ip_range,omitzero"`
	// Match the JA3 TLS client fingerprint of the request
	Ja3 DomainPolicyOverrideNewParamsConditionJa3 `json:"ja3,omitzero"`
	// Match the JA4 TLS client fingerprint of the request
	Ja4 DomainPolicyOverrideNewParamsConditionJa4 `json:"ja4,omitzero"`
	// Match the organization the request originated from, as determined by a WHOIS
	// lookup of the requesting IP
	Organization DomainPolicyOverrideNewParamsConditionOrganization `json:"organization,omitzero"`
	// Match the incoming request URL
	URL DomainPolicyOverrideNewParamsConditionURL `json:"url,omitzero"`
	// Match the user agent making the request
	UserAgent DomainPolicyOverrideNewParamsConditionUserAgent `json:"user_agent,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideNewParamsCondition) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideNewParamsCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideNewParamsCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the country that the request originated from
//
// The property CountryCode is required.
type DomainPolicyOverrideNewParamsConditionCountry struct {
	// A list of ISO 3166-1 alpha-2 formatted strings representing the countries to
	// match against
	CountryCode []string `json:"country_code,omitzero" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideNewParamsConditionCountry) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideNewParamsConditionCountry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideNewParamsConditionCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match an incoming request header
//
// The properties Header, Value are required.
type DomainPolicyOverrideNewParamsConditionHeader struct {
	// The request header name
	Header string `json:"header" api:"required"`
	// The request header value
	Value string `json:"value" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// The type of matching condition for header and value.
	//
	// Any of "Exact", "Contains".
	MatchType string `json:"match_type,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideNewParamsConditionHeader) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideNewParamsConditionHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideNewParamsConditionHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainPolicyOverrideNewParamsConditionHeader](
		"match_type", "Exact", "Contains",
	)
}

// Match when an incoming request header is present
//
// The property Header is required.
type DomainPolicyOverrideNewParamsConditionHeaderExists struct {
	// The request header name
	Header string `json:"header" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideNewParamsConditionHeaderExists) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideNewParamsConditionHeaderExists
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideNewParamsConditionHeaderExists) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming HTTP method
//
// The property HTTPMethod is required.
type DomainPolicyOverrideNewParamsConditionHTTPMethod struct {
	// HTTP methods of a request
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod string `json:"http_method,omitzero" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideNewParamsConditionHTTPMethod) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideNewParamsConditionHTTPMethod
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideNewParamsConditionHTTPMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainPolicyOverrideNewParamsConditionHTTPMethod](
		"http_method", "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT", "TRACE",
	)
}

// Match the incoming request against a single IP address
//
// The property IPAddress is required.
type DomainPolicyOverrideNewParamsConditionIP struct {
	// A single IPv4 or IPv6 address
	IPAddress string `json:"ip_address" api:"required" format:"ipvanyaddress"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideNewParamsConditionIP) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideNewParamsConditionIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideNewParamsConditionIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against an IP range
//
// The properties LowerBound, UpperBound are required.
type DomainPolicyOverrideNewParamsConditionIPRange struct {
	// The lower bound IPv4 or IPv6 address to match against
	LowerBound string `json:"lower_bound" api:"required" format:"ipvanyaddress"`
	// The upper bound IPv4 or IPv6 address to match against
	UpperBound string `json:"upper_bound" api:"required" format:"ipvanyaddress"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideNewParamsConditionIPRange) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideNewParamsConditionIPRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideNewParamsConditionIPRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the JA3 TLS client fingerprint of the request
//
// The property Ja3Fingerprints is required.
type DomainPolicyOverrideNewParamsConditionJa3 struct {
	// A list of JA3 TLS client fingerprints to match against the request
	Ja3Fingerprints []string `json:"ja3_fingerprints,omitzero" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideNewParamsConditionJa3) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideNewParamsConditionJa3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideNewParamsConditionJa3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the JA4 TLS client fingerprint of the request
//
// The property Ja4Fingerprints is required.
type DomainPolicyOverrideNewParamsConditionJa4 struct {
	// A list of JA4 TLS client fingerprints to match against the request
	Ja4Fingerprints []string `json:"ja4_fingerprints,omitzero" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideNewParamsConditionJa4) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideNewParamsConditionJa4
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideNewParamsConditionJa4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the organization the request originated from, as determined by a WHOIS
// lookup of the requesting IP
//
// The property Organization is required.
type DomainPolicyOverrideNewParamsConditionOrganization struct {
	// The organization to match against
	Organization string `json:"organization" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideNewParamsConditionOrganization) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideNewParamsConditionOrganization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideNewParamsConditionOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request URL
//
// The property URL is required.
type DomainPolicyOverrideNewParamsConditionURL struct {
	// The pattern to match against the request URL. Constraints depend on
	// `match_type`:
	//
	//   - **Exact/Contains**: plain text matching (e.g., `/admin`, must comply with
	//     `^[\w!\$~:#\[\]@\(\)*\+,=\/\-\.\%]+$`).
	//   - **Regex**: a valid regular expression (e.g., `^/upload(/\d+)?/\w+`).
	//     Lookahead/lookbehind constructs are forbidden.
	URL string `json:"url" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// The type of matching condition.
	//
	// Any of "Exact", "Contains", "Regex".
	MatchType string `json:"match_type,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideNewParamsConditionURL) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideNewParamsConditionURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideNewParamsConditionURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainPolicyOverrideNewParamsConditionURL](
		"match_type", "Exact", "Contains", "Regex",
	)
}

// Match the user agent making the request
//
// The property UserAgent is required.
type DomainPolicyOverrideNewParamsConditionUserAgent struct {
	// The user agent value to match
	UserAgent string `json:"user_agent" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// The type of matching condition.
	//
	// Any of "Exact", "Contains".
	MatchType string `json:"match_type,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideNewParamsConditionUserAgent) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideNewParamsConditionUserAgent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideNewParamsConditionUserAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainPolicyOverrideNewParamsConditionUserAgent](
		"match_type", "Exact", "Contains",
	)
}

type DomainPolicyOverrideUpdateParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id" api:"required" json:"-"`
	// An optional note explaining the override
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the override is active
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// The policy override's display name
	Name param.Opt[string] `json:"name,omitzero"`
	// The conditions required for the override to apply. An override may have up to 5
	// conditions; all must pass for the override to apply. Without conditions the
	// override applies to every request on the owning domain that hits a linked
	// target.
	Conditions []DomainPolicyOverrideUpdateParamsCondition `json:"conditions,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single policy override condition. Only one of the fields below may be set.
//
// Limited to the subset of `CustomRuleConditionScope` scopes that policy overrides
// support (Lime wire-format translation is reused from Custom Rules' own
// `app.models.customer_rules.lime_translation.ConditionMapper`).
type DomainPolicyOverrideUpdateParamsCondition struct {
	// Match the country that the request originated from
	Country DomainPolicyOverrideUpdateParamsConditionCountry `json:"country,omitzero"`
	// Match an incoming request header
	Header DomainPolicyOverrideUpdateParamsConditionHeader `json:"header,omitzero"`
	// Match when an incoming request header is present
	HeaderExists DomainPolicyOverrideUpdateParamsConditionHeaderExists `json:"header_exists,omitzero"`
	// Match the incoming HTTP method
	HTTPMethod DomainPolicyOverrideUpdateParamsConditionHTTPMethod `json:"http_method,omitzero"`
	// Match the incoming request against a single IP address
	IP DomainPolicyOverrideUpdateParamsConditionIP `json:"ip,omitzero"`
	// Match the incoming request against an IP range
	IPRange DomainPolicyOverrideUpdateParamsConditionIPRange `json:"ip_range,omitzero"`
	// Match the JA3 TLS client fingerprint of the request
	Ja3 DomainPolicyOverrideUpdateParamsConditionJa3 `json:"ja3,omitzero"`
	// Match the JA4 TLS client fingerprint of the request
	Ja4 DomainPolicyOverrideUpdateParamsConditionJa4 `json:"ja4,omitzero"`
	// Match the organization the request originated from, as determined by a WHOIS
	// lookup of the requesting IP
	Organization DomainPolicyOverrideUpdateParamsConditionOrganization `json:"organization,omitzero"`
	// Match the incoming request URL
	URL DomainPolicyOverrideUpdateParamsConditionURL `json:"url,omitzero"`
	// Match the user agent making the request
	UserAgent DomainPolicyOverrideUpdateParamsConditionUserAgent `json:"user_agent,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideUpdateParamsCondition) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideUpdateParamsCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideUpdateParamsCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the country that the request originated from
//
// The property CountryCode is required.
type DomainPolicyOverrideUpdateParamsConditionCountry struct {
	// A list of ISO 3166-1 alpha-2 formatted strings representing the countries to
	// match against
	CountryCode []string `json:"country_code,omitzero" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideUpdateParamsConditionCountry) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideUpdateParamsConditionCountry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideUpdateParamsConditionCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match an incoming request header
//
// The properties Header, Value are required.
type DomainPolicyOverrideUpdateParamsConditionHeader struct {
	// The request header name
	Header string `json:"header" api:"required"`
	// The request header value
	Value string `json:"value" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// The type of matching condition for header and value.
	//
	// Any of "Exact", "Contains".
	MatchType string `json:"match_type,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideUpdateParamsConditionHeader) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideUpdateParamsConditionHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideUpdateParamsConditionHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainPolicyOverrideUpdateParamsConditionHeader](
		"match_type", "Exact", "Contains",
	)
}

// Match when an incoming request header is present
//
// The property Header is required.
type DomainPolicyOverrideUpdateParamsConditionHeaderExists struct {
	// The request header name
	Header string `json:"header" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideUpdateParamsConditionHeaderExists) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideUpdateParamsConditionHeaderExists
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideUpdateParamsConditionHeaderExists) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming HTTP method
//
// The property HTTPMethod is required.
type DomainPolicyOverrideUpdateParamsConditionHTTPMethod struct {
	// HTTP methods of a request
	//
	// Any of "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT",
	// "TRACE".
	HTTPMethod string `json:"http_method,omitzero" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideUpdateParamsConditionHTTPMethod) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideUpdateParamsConditionHTTPMethod
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideUpdateParamsConditionHTTPMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainPolicyOverrideUpdateParamsConditionHTTPMethod](
		"http_method", "CONNECT", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT", "TRACE",
	)
}

// Match the incoming request against a single IP address
//
// The property IPAddress is required.
type DomainPolicyOverrideUpdateParamsConditionIP struct {
	// A single IPv4 or IPv6 address
	IPAddress string `json:"ip_address" api:"required" format:"ipvanyaddress"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideUpdateParamsConditionIP) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideUpdateParamsConditionIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideUpdateParamsConditionIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against an IP range
//
// The properties LowerBound, UpperBound are required.
type DomainPolicyOverrideUpdateParamsConditionIPRange struct {
	// The lower bound IPv4 or IPv6 address to match against
	LowerBound string `json:"lower_bound" api:"required" format:"ipvanyaddress"`
	// The upper bound IPv4 or IPv6 address to match against
	UpperBound string `json:"upper_bound" api:"required" format:"ipvanyaddress"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideUpdateParamsConditionIPRange) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideUpdateParamsConditionIPRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideUpdateParamsConditionIPRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the JA3 TLS client fingerprint of the request
//
// The property Ja3Fingerprints is required.
type DomainPolicyOverrideUpdateParamsConditionJa3 struct {
	// A list of JA3 TLS client fingerprints to match against the request
	Ja3Fingerprints []string `json:"ja3_fingerprints,omitzero" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideUpdateParamsConditionJa3) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideUpdateParamsConditionJa3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideUpdateParamsConditionJa3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the JA4 TLS client fingerprint of the request
//
// The property Ja4Fingerprints is required.
type DomainPolicyOverrideUpdateParamsConditionJa4 struct {
	// A list of JA4 TLS client fingerprints to match against the request
	Ja4Fingerprints []string `json:"ja4_fingerprints,omitzero" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideUpdateParamsConditionJa4) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideUpdateParamsConditionJa4
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideUpdateParamsConditionJa4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the organization the request originated from, as determined by a WHOIS
// lookup of the requesting IP
//
// The property Organization is required.
type DomainPolicyOverrideUpdateParamsConditionOrganization struct {
	// The organization to match against
	Organization string `json:"organization" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideUpdateParamsConditionOrganization) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideUpdateParamsConditionOrganization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideUpdateParamsConditionOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request URL
//
// The property URL is required.
type DomainPolicyOverrideUpdateParamsConditionURL struct {
	// The pattern to match against the request URL. Constraints depend on
	// `match_type`:
	//
	//   - **Exact/Contains**: plain text matching (e.g., `/admin`, must comply with
	//     `^[\w!\$~:#\[\]@\(\)*\+,=\/\-\.\%]+$`).
	//   - **Regex**: a valid regular expression (e.g., `^/upload(/\d+)?/\w+`).
	//     Lookahead/lookbehind constructs are forbidden.
	URL string `json:"url" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// The type of matching condition.
	//
	// Any of "Exact", "Contains", "Regex".
	MatchType string `json:"match_type,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideUpdateParamsConditionURL) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideUpdateParamsConditionURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideUpdateParamsConditionURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainPolicyOverrideUpdateParamsConditionURL](
		"match_type", "Exact", "Contains", "Regex",
	)
}

// Match the user agent making the request
//
// The property UserAgent is required.
type DomainPolicyOverrideUpdateParamsConditionUserAgent struct {
	// The user agent value to match
	UserAgent string `json:"user_agent" api:"required"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	// The type of matching condition.
	//
	// Any of "Exact", "Contains".
	MatchType string `json:"match_type,omitzero"`
	paramObj
}

func (r DomainPolicyOverrideUpdateParamsConditionUserAgent) MarshalJSON() (data []byte, err error) {
	type shadow DomainPolicyOverrideUpdateParamsConditionUserAgent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainPolicyOverrideUpdateParamsConditionUserAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainPolicyOverrideUpdateParamsConditionUserAgent](
		"match_type", "Exact", "Contains",
	)
}

type DomainPolicyOverrideListParams struct {
	// Filter by enabled state.
	Enabled param.Opt[bool] `query:"enabled,omitzero" json:"-"`
	// Search policy override names and descriptions.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Determine the field used to order results.
	//
	// Any of "id", "name", "target_type", "enabled", "stale", "created_at",
	// "updated_at", "-id", "-name", "-target_type", "-enabled", "-stale",
	// "-created_at", "-updated_at".
	Ordering DomainPolicyOverrideListParamsOrdering `query:"ordering,omitzero" json:"-"`
	// Filter by linked target type.
	//
	// Any of "detector", "rule".
	TargetType DomainPolicyOverrideListParamsTargetType `query:"target_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainPolicyOverrideListParams]'s query parameters as
// `url.Values`.
func (r DomainPolicyOverrideListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Determine the field used to order results.
type DomainPolicyOverrideListParamsOrdering string

const (
	DomainPolicyOverrideListParamsOrderingID              DomainPolicyOverrideListParamsOrdering = "id"
	DomainPolicyOverrideListParamsOrderingName            DomainPolicyOverrideListParamsOrdering = "name"
	DomainPolicyOverrideListParamsOrderingTargetType      DomainPolicyOverrideListParamsOrdering = "target_type"
	DomainPolicyOverrideListParamsOrderingEnabled         DomainPolicyOverrideListParamsOrdering = "enabled"
	DomainPolicyOverrideListParamsOrderingStale           DomainPolicyOverrideListParamsOrdering = "stale"
	DomainPolicyOverrideListParamsOrderingCreatedAt       DomainPolicyOverrideListParamsOrdering = "created_at"
	DomainPolicyOverrideListParamsOrderingUpdatedAt       DomainPolicyOverrideListParamsOrdering = "updated_at"
	DomainPolicyOverrideListParamsOrderingMinusID         DomainPolicyOverrideListParamsOrdering = "-id"
	DomainPolicyOverrideListParamsOrderingMinusName       DomainPolicyOverrideListParamsOrdering = "-name"
	DomainPolicyOverrideListParamsOrderingMinusTargetType DomainPolicyOverrideListParamsOrdering = "-target_type"
	DomainPolicyOverrideListParamsOrderingMinusEnabled    DomainPolicyOverrideListParamsOrdering = "-enabled"
	DomainPolicyOverrideListParamsOrderingMinusStale      DomainPolicyOverrideListParamsOrdering = "-stale"
	DomainPolicyOverrideListParamsOrderingMinusCreatedAt  DomainPolicyOverrideListParamsOrdering = "-created_at"
	DomainPolicyOverrideListParamsOrderingMinusUpdatedAt  DomainPolicyOverrideListParamsOrdering = "-updated_at"
)

// Filter by linked target type.
type DomainPolicyOverrideListParamsTargetType string

const (
	DomainPolicyOverrideListParamsTargetTypeDetector DomainPolicyOverrideListParamsTargetType = "detector"
	DomainPolicyOverrideListParamsTargetTypeRule     DomainPolicyOverrideListParamsTargetType = "rule"
)

type DomainPolicyOverrideDeleteParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id" api:"required" json:"-"`
	paramObj
}

type DomainPolicyOverrideGetParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id" api:"required" json:"-"`
	paramObj
}

type DomainPolicyOverrideListOwaspTopThreatsTargetsParams struct {
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainPolicyOverrideListOwaspTopThreatsTargetsParams]'s
// query parameters as `url.Values`.
func (r DomainPolicyOverrideListOwaspTopThreatsTargetsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
