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

// DomainFirewallRuleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainFirewallRuleService] method instead.
type DomainFirewallRuleService struct {
	Options []option.RequestOption
}

// NewDomainFirewallRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDomainFirewallRuleService(opts ...option.RequestOption) (r DomainFirewallRuleService) {
	r = DomainFirewallRuleService{}
	r.Options = opts
	return
}

// Create a firewall rule
func (r *DomainFirewallRuleService) New(ctx context.Context, domainID int64, body DomainFirewallRuleNewParams, opts ...option.RequestOption) (res *WaapFirewallRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/firewall-rules", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Only properties present in the request will be updated
func (r *DomainFirewallRuleService) Update(ctx context.Context, ruleID int64, params DomainFirewallRuleUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/firewall-rules/%v", params.DomainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return
}

// Extracts a list of firewall rules assigned to a domain, offering filter,
// ordering, and pagination capabilities
func (r *DomainFirewallRuleService) List(ctx context.Context, domainID int64, query DomainFirewallRuleListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapFirewallRule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/firewall-rules", domainID)
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

// Extracts a list of firewall rules assigned to a domain, offering filter,
// ordering, and pagination capabilities
func (r *DomainFirewallRuleService) ListAutoPaging(ctx context.Context, domainID int64, query DomainFirewallRuleListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapFirewallRule] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, domainID, query, opts...))
}

// Delete a firewall rule
func (r *DomainFirewallRuleService) Delete(ctx context.Context, ruleID int64, body DomainFirewallRuleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/firewall-rules/%v", body.DomainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Delete multiple WAAP rules
func (r *DomainFirewallRuleService) DeleteMultiple(ctx context.Context, domainID int64, body DomainFirewallRuleDeleteMultipleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/firewall-rules/bulk_delete", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Extracts a specific firewall rule assigned to a domain
func (r *DomainFirewallRuleService) Get(ctx context.Context, ruleID int64, query DomainFirewallRuleGetParams, opts ...option.RequestOption) (res *WaapFirewallRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/firewall-rules/%v", query.DomainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Toggle a firewall rule
func (r *DomainFirewallRuleService) Toggle(ctx context.Context, action DomainFirewallRuleToggleParamsAction, body DomainFirewallRuleToggleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/firewall-rules/%v/%v", body.DomainID, body.RuleID, action)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, nil, opts...)
	return
}

type WaapFirewallRule struct {
	// The unique identifier of the rule
	ID int64 `json:"id" api:"required"`
	// The action that the rule takes when triggered
	Action WaapFirewallRuleAction `json:"action" api:"required"`
	// The condition required for the WAAP engine to trigger the rule.
	Conditions []WaapFirewallRuleCondition `json:"conditions" api:"required"`
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
		Enabled     respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapFirewallRule) RawJSON() string { return r.JSON.raw }
func (r *WaapFirewallRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the rule takes when triggered
type WaapFirewallRuleAction struct {
	// The WAAP allows the request
	Allow map[string]any `json:"allow" api:"nullable"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block WaapFirewallRuleActionBlock `json:"block" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allow       respjson.Field
		Block       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapFirewallRuleAction) RawJSON() string { return r.JSON.raw }
func (r *WaapFirewallRuleAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WAAP block action behavior could be configured with response status code and
// action duration.
type WaapFirewallRuleActionBlock struct {
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
func (r WaapFirewallRuleActionBlock) RawJSON() string { return r.JSON.raw }
func (r *WaapFirewallRuleActionBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The criteria of an incoming web request and the models of the various values
// those criteria can take
type WaapFirewallRuleCondition struct {
	// Match the incoming request against a single IP address
	IP WaapFirewallRuleConditionIP `json:"ip"`
	// Match the incoming request against an IP range
	IPRange WaapFirewallRuleConditionIPRange `json:"ip_range"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IP          respjson.Field
		IPRange     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapFirewallRuleCondition) RawJSON() string { return r.JSON.raw }
func (r *WaapFirewallRuleCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against a single IP address
type WaapFirewallRuleConditionIP struct {
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
func (r WaapFirewallRuleConditionIP) RawJSON() string { return r.JSON.raw }
func (r *WaapFirewallRuleConditionIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against an IP range
type WaapFirewallRuleConditionIPRange struct {
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
func (r WaapFirewallRuleConditionIPRange) RawJSON() string { return r.JSON.raw }
func (r *WaapFirewallRuleConditionIPRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainFirewallRuleNewParams struct {
	// The action that the rule takes when triggered
	Action DomainFirewallRuleNewParamsAction `json:"action,omitzero" api:"required"`
	// The condition required for the WAAP engine to trigger the rule.
	Conditions []DomainFirewallRuleNewParamsCondition `json:"conditions,omitzero" api:"required"`
	// Whether or not the rule is enabled
	Enabled bool `json:"enabled" api:"required"`
	// The name assigned to the rule
	Name string `json:"name" api:"required"`
	// The description assigned to the rule
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r DomainFirewallRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainFirewallRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainFirewallRuleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the rule takes when triggered
type DomainFirewallRuleNewParamsAction struct {
	// The WAAP allows the request
	Allow map[string]any `json:"allow,omitzero"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block DomainFirewallRuleNewParamsActionBlock `json:"block,omitzero"`
	paramObj
}

func (r DomainFirewallRuleNewParamsAction) MarshalJSON() (data []byte, err error) {
	type shadow DomainFirewallRuleNewParamsAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainFirewallRuleNewParamsAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WAAP block action behavior could be configured with response status code and
// action duration.
type DomainFirewallRuleNewParamsActionBlock struct {
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

func (r DomainFirewallRuleNewParamsActionBlock) MarshalJSON() (data []byte, err error) {
	type shadow DomainFirewallRuleNewParamsActionBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainFirewallRuleNewParamsActionBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainFirewallRuleNewParamsActionBlock](
		"status_code", 403, 405, 418, 429,
	)
}

// The criteria of an incoming web request and the models of the various values
// those criteria can take
type DomainFirewallRuleNewParamsCondition struct {
	// Match the incoming request against a single IP address
	IP DomainFirewallRuleNewParamsConditionIP `json:"ip,omitzero"`
	// Match the incoming request against an IP range
	IPRange DomainFirewallRuleNewParamsConditionIPRange `json:"ip_range,omitzero"`
	paramObj
}

func (r DomainFirewallRuleNewParamsCondition) MarshalJSON() (data []byte, err error) {
	type shadow DomainFirewallRuleNewParamsCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainFirewallRuleNewParamsCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against a single IP address
//
// The property IPAddress is required.
type DomainFirewallRuleNewParamsConditionIP struct {
	// A single IPv4 or IPv6 address
	IPAddress string `json:"ip_address" api:"required" format:"ipvanyaddress"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainFirewallRuleNewParamsConditionIP) MarshalJSON() (data []byte, err error) {
	type shadow DomainFirewallRuleNewParamsConditionIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainFirewallRuleNewParamsConditionIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against an IP range
//
// The properties LowerBound, UpperBound are required.
type DomainFirewallRuleNewParamsConditionIPRange struct {
	// The lower bound IPv4 or IPv6 address to match against
	LowerBound string `json:"lower_bound" api:"required" format:"ipvanyaddress"`
	// The upper bound IPv4 or IPv6 address to match against
	UpperBound string `json:"upper_bound" api:"required" format:"ipvanyaddress"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainFirewallRuleNewParamsConditionIPRange) MarshalJSON() (data []byte, err error) {
	type shadow DomainFirewallRuleNewParamsConditionIPRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainFirewallRuleNewParamsConditionIPRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainFirewallRuleUpdateParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id" api:"required" json:"-"`
	// The description assigned to the rule
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether or not the rule is enabled
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// The name assigned to the rule
	Name param.Opt[string] `json:"name,omitzero"`
	// The action that a firewall rule takes when triggered
	Action DomainFirewallRuleUpdateParamsAction `json:"action,omitzero"`
	// The condition required for the WAAP engine to trigger the rule.
	Conditions []DomainFirewallRuleUpdateParamsCondition `json:"conditions,omitzero"`
	paramObj
}

func (r DomainFirewallRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainFirewallRuleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainFirewallRuleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action that a firewall rule takes when triggered
type DomainFirewallRuleUpdateParamsAction struct {
	// The WAAP allows the request
	Allow map[string]any `json:"allow,omitzero"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block DomainFirewallRuleUpdateParamsActionBlock `json:"block,omitzero"`
	paramObj
}

func (r DomainFirewallRuleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	type shadow DomainFirewallRuleUpdateParamsAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainFirewallRuleUpdateParamsAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WAAP block action behavior could be configured with response status code and
// action duration.
type DomainFirewallRuleUpdateParamsActionBlock struct {
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

func (r DomainFirewallRuleUpdateParamsActionBlock) MarshalJSON() (data []byte, err error) {
	type shadow DomainFirewallRuleUpdateParamsActionBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainFirewallRuleUpdateParamsActionBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainFirewallRuleUpdateParamsActionBlock](
		"status_code", 403, 405, 418, 429,
	)
}

// The criteria of an incoming web request and the models of the various values
// those criteria can take
type DomainFirewallRuleUpdateParamsCondition struct {
	// Match the incoming request against a single IP address
	IP DomainFirewallRuleUpdateParamsConditionIP `json:"ip,omitzero"`
	// Match the incoming request against an IP range
	IPRange DomainFirewallRuleUpdateParamsConditionIPRange `json:"ip_range,omitzero"`
	paramObj
}

func (r DomainFirewallRuleUpdateParamsCondition) MarshalJSON() (data []byte, err error) {
	type shadow DomainFirewallRuleUpdateParamsCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainFirewallRuleUpdateParamsCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against a single IP address
//
// The property IPAddress is required.
type DomainFirewallRuleUpdateParamsConditionIP struct {
	// A single IPv4 or IPv6 address
	IPAddress string `json:"ip_address" api:"required" format:"ipvanyaddress"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainFirewallRuleUpdateParamsConditionIP) MarshalJSON() (data []byte, err error) {
	type shadow DomainFirewallRuleUpdateParamsConditionIP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainFirewallRuleUpdateParamsConditionIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Match the incoming request against an IP range
//
// The properties LowerBound, UpperBound are required.
type DomainFirewallRuleUpdateParamsConditionIPRange struct {
	// The lower bound IPv4 or IPv6 address to match against
	LowerBound string `json:"lower_bound" api:"required" format:"ipvanyaddress"`
	// The upper bound IPv4 or IPv6 address to match against
	UpperBound string `json:"upper_bound" api:"required" format:"ipvanyaddress"`
	// Whether or not to apply a boolean NOT operation to the rule's condition
	Negation param.Opt[bool] `json:"negation,omitzero"`
	paramObj
}

func (r DomainFirewallRuleUpdateParamsConditionIPRange) MarshalJSON() (data []byte, err error) {
	type shadow DomainFirewallRuleUpdateParamsConditionIPRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainFirewallRuleUpdateParamsConditionIPRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainFirewallRuleListParams struct {
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
	Ordering DomainFirewallRuleListParamsOrdering `query:"ordering,omitzero" json:"-"`
	// Filter to refine results by specific firewall actions
	//
	// Any of "allow", "block".
	Action DomainFirewallRuleListParamsAction `query:"action,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainFirewallRuleListParams]'s query parameters as
// `url.Values`.
func (r DomainFirewallRuleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter to refine results by specific firewall actions
type DomainFirewallRuleListParamsAction string

const (
	DomainFirewallRuleListParamsActionAllow DomainFirewallRuleListParamsAction = "allow"
	DomainFirewallRuleListParamsActionBlock DomainFirewallRuleListParamsAction = "block"
)

// Determine the field to order results by
type DomainFirewallRuleListParamsOrdering string

const (
	DomainFirewallRuleListParamsOrderingID               DomainFirewallRuleListParamsOrdering = "id"
	DomainFirewallRuleListParamsOrderingName             DomainFirewallRuleListParamsOrdering = "name"
	DomainFirewallRuleListParamsOrderingDescription      DomainFirewallRuleListParamsOrdering = "description"
	DomainFirewallRuleListParamsOrderingEnabled          DomainFirewallRuleListParamsOrdering = "enabled"
	DomainFirewallRuleListParamsOrderingAction           DomainFirewallRuleListParamsOrdering = "action"
	DomainFirewallRuleListParamsOrderingMinusID          DomainFirewallRuleListParamsOrdering = "-id"
	DomainFirewallRuleListParamsOrderingMinusName        DomainFirewallRuleListParamsOrdering = "-name"
	DomainFirewallRuleListParamsOrderingMinusDescription DomainFirewallRuleListParamsOrdering = "-description"
	DomainFirewallRuleListParamsOrderingMinusEnabled     DomainFirewallRuleListParamsOrdering = "-enabled"
	DomainFirewallRuleListParamsOrderingMinusAction      DomainFirewallRuleListParamsOrdering = "-action"
)

type DomainFirewallRuleDeleteParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id" api:"required" json:"-"`
	paramObj
}

type DomainFirewallRuleDeleteMultipleParams struct {
	// The IDs of the rules to delete
	RuleIDs []int64 `json:"rule_ids,omitzero" api:"required"`
	paramObj
}

func (r DomainFirewallRuleDeleteMultipleParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainFirewallRuleDeleteMultipleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainFirewallRuleDeleteMultipleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainFirewallRuleGetParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id" api:"required" json:"-"`
	paramObj
}

type DomainFirewallRuleToggleParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id" api:"required" json:"-"`
	// The firewall rule ID
	RuleID int64 `path:"rule_id" api:"required" json:"-"`
	paramObj
}

// Enable or disable a firewall rule
type DomainFirewallRuleToggleParamsAction string

const (
	DomainFirewallRuleToggleParamsActionEnable  DomainFirewallRuleToggleParamsAction = "enable"
	DomainFirewallRuleToggleParamsActionDisable DomainFirewallRuleToggleParamsAction = "disable"
)
