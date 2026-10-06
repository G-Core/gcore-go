// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
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

// FirewallRuleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFirewallRuleService] method instead.
type FirewallRuleService struct {
	Options []option.RequestOption
}

// NewFirewallRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFirewallRuleService(opts ...option.RequestOption) (r FirewallRuleService) {
	r = FirewallRuleService{}
	r.Options = opts
	return
}

// Returns all firewall rules across all domains belonging to the authenticated
// account. Supports filtering, ordering, and pagination.
func (r *FirewallRuleService) List(ctx context.Context, query FirewallRuleListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapAccountFirewallRule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "waap/v2/firewall-rules"
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

// Returns all firewall rules across all domains belonging to the authenticated
// account. Supports filtering, ordering, and pagination.
func (r *FirewallRuleService) ListAutoPaging(ctx context.Context, query FirewallRuleListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapAccountFirewallRule] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Firewall rule response for account-level queries with domain information
type WaapAccountFirewallRule struct {
	// The unique identifier of the rule
	ID int64 `json:"id" api:"required"`
	// The action that the rule takes when triggered
	Action WaapAccountFirewallRuleAction `json:"action" api:"required"`
	// The condition required for the WAAP engine to trigger the rule.
	Conditions []WaapAccountFirewallRuleCondition `json:"conditions" api:"required"`
	// The domain this firewall rule belongs to
	Domain WaapAccountFirewallRuleDomain `json:"domain" api:"required"`
	// Whether or not the rule is enabled
	Enabled bool `json:"enabled" api:"required"`
	// The name assigned to the rule
	Name string `json:"name" api:"required"`
	// The description assigned to the rule
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Action      respjson.Field
		Conditions  respjson.Field
		Domain      respjson.Field
		Enabled     respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAccountFirewallRule) RawJSON() string { return r.JSON.raw }
func (r *WaapAccountFirewallRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the rule takes when triggered
type WaapAccountFirewallRuleAction struct {
	// The WAAP allows the request
	Allow map[string]any `json:"allow" api:"nullable"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block WaapAccountFirewallRuleActionBlock `json:"block" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allow       respjson.Field
		Block       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAccountFirewallRuleAction) RawJSON() string { return r.JSON.raw }
func (r *WaapAccountFirewallRuleAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WAAP block action behavior could be configured with response status code and
// action duration.
type WaapAccountFirewallRuleActionBlock struct {
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
func (r WaapAccountFirewallRuleActionBlock) RawJSON() string { return r.JSON.raw }
func (r *WaapAccountFirewallRuleActionBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The criteria of an incoming web request and the models of the various values
// those criteria can take
type WaapAccountFirewallRuleCondition struct {
	// Match the incoming request against a single IP address
	IP WaapAccountFirewallRuleConditionIP `json:"ip"`
	// Match the incoming request against an IP range
	IPRange WaapAccountFirewallRuleConditionIPRange `json:"ip_range"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IP          respjson.Field
		IPRange     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAccountFirewallRuleCondition) RawJSON() string { return r.JSON.raw }
func (r *WaapAccountFirewallRuleCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against a single IP address
type WaapAccountFirewallRuleConditionIP struct {
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
func (r WaapAccountFirewallRuleConditionIP) RawJSON() string { return r.JSON.raw }
func (r *WaapAccountFirewallRuleConditionIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against an IP range
type WaapAccountFirewallRuleConditionIPRange struct {
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
func (r WaapAccountFirewallRuleConditionIPRange) RawJSON() string { return r.JSON.raw }
func (r *WaapAccountFirewallRuleConditionIPRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The domain this firewall rule belongs to
type WaapAccountFirewallRuleDomain struct {
	// The domain ID
	ID int64 `json:"id" api:"required"`
	// The domain name
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
func (r WaapAccountFirewallRuleDomain) RawJSON() string { return r.JSON.raw }
func (r *WaapAccountFirewallRuleDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleListParams struct {
	// Filter rules by description (supports wildcards)
	Description param.Opt[string] `query:"description,omitzero" json:"-"`
	// Filter by enabled status
	Enabled param.Opt[bool] `query:"enabled,omitzero" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter rules by name (supports wildcards)
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Order results by field (prefix with - for descending)
	//
	// Any of "id", "name", "description", "enabled", "action", "-id", "-name",
	// "-description", "-enabled", "-action".
	Ordering FirewallRuleListParamsOrdering `query:"ordering,omitzero" json:"-"`
	// Filter by action type
	//
	// Any of "allow", "block".
	Action FirewallRuleListParamsAction `query:"action,omitzero" json:"-"`
	// Filter by domain IDs (can specify multiple)
	DomainIDs []int64 `query:"domain_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FirewallRuleListParams]'s query parameters as `url.Values`.
func (r FirewallRuleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by action type
type FirewallRuleListParamsAction string

const (
	FirewallRuleListParamsActionAllow FirewallRuleListParamsAction = "allow"
	FirewallRuleListParamsActionBlock FirewallRuleListParamsAction = "block"
)

// Order results by field (prefix with - for descending)
type FirewallRuleListParamsOrdering string

const (
	FirewallRuleListParamsOrderingID               FirewallRuleListParamsOrdering = "id"
	FirewallRuleListParamsOrderingName             FirewallRuleListParamsOrdering = "name"
	FirewallRuleListParamsOrderingDescription      FirewallRuleListParamsOrdering = "description"
	FirewallRuleListParamsOrderingEnabled          FirewallRuleListParamsOrdering = "enabled"
	FirewallRuleListParamsOrderingAction           FirewallRuleListParamsOrdering = "action"
	FirewallRuleListParamsOrderingMinusID          FirewallRuleListParamsOrdering = "-id"
	FirewallRuleListParamsOrderingMinusName        FirewallRuleListParamsOrdering = "-name"
	FirewallRuleListParamsOrderingMinusDescription FirewallRuleListParamsOrdering = "-description"
	FirewallRuleListParamsOrderingMinusEnabled     FirewallRuleListParamsOrdering = "-enabled"
	FirewallRuleListParamsOrderingMinusAction      FirewallRuleListParamsOrdering = "-action"
)
