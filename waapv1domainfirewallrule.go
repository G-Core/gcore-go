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

// WaapV1DomainFirewallRuleService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1DomainFirewallRuleService] method instead.
type WaapV1DomainFirewallRuleService struct {
	Options []option.RequestOption
}

// NewWaapV1DomainFirewallRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWaapV1DomainFirewallRuleService(opts ...option.RequestOption) (r *WaapV1DomainFirewallRuleService) {
	r = &WaapV1DomainFirewallRuleService{}
	r.Options = opts
	return
}

// Create a firewall rule
func (r *WaapV1DomainFirewallRuleService) New(ctx context.Context, domainID int64, body WaapV1DomainFirewallRuleNewParams, opts ...option.RequestOption) (res *FirewallRule, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/firewall-rules", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Extracts a specific firewall rule assigned to a domain
func (r *WaapV1DomainFirewallRuleService) Get(ctx context.Context, domainID int64, ruleID int64, opts ...option.RequestOption) (res *FirewallRule, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/firewall-rules/%v", domainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Only properties present in the request will be updated
func (r *WaapV1DomainFirewallRuleService) Update(ctx context.Context, domainID int64, ruleID int64, body WaapV1DomainFirewallRuleUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/firewall-rules/%v", domainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Extracts a list of firewall rules assigned to a domain, offering filter,
// ordering, and pagination capabilities
func (r *WaapV1DomainFirewallRuleService) List(ctx context.Context, domainID int64, query WaapV1DomainFirewallRuleListParams, opts ...option.RequestOption) (res *WaapV1DomainFirewallRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/firewall-rules", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a firewall rule
func (r *WaapV1DomainFirewallRuleService) Delete(ctx context.Context, domainID int64, ruleID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/firewall-rules/%v", domainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Delete multiple WAAP rules
func (r *WaapV1DomainFirewallRuleService) BulkDelete(ctx context.Context, domainID int64, body WaapV1DomainFirewallRuleBulkDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/firewall-rules/bulk_delete", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Toggle a firewall rule
func (r *WaapV1DomainFirewallRuleService) Toggle(ctx context.Context, domainID int64, ruleID int64, action CustomerRuleState, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/firewall-rules/%v/%v", domainID, ruleID, action)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, nil, opts...)
	return
}

type FirewallRule struct {
	// The unique identifier of the rule
	ID int64 `json:"id,required"`
	// The action that a firewall rule takes when triggered
	Action FirewallRuleAction `json:"action,required"`
	// The condition required for the WAAP engine to trigger the rule.
	Conditions []FirewallRuleCondition `json:"conditions,required"`
	// Whether or not the rule is enabled
	Enabled bool `json:"enabled,required"`
	// The name assigned to the rule
	Name string `json:"name,required"`
	// The description assigned to the rule
	Description string           `json:"description,nullable"`
	JSON        firewallRuleJSON `json:"-"`
}

// firewallRuleJSON contains the JSON metadata for the struct [FirewallRule]
type firewallRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Conditions  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleJSON) RawJSON() string {
	return r.raw
}

// The action that a firewall rule takes when triggered
type FirewallRuleAction struct {
	// The WAAP allowed the request
	Allow RuleAllowAction `json:"allow,nullable"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block RuleBlockAction        `json:"block,nullable"`
	JSON  firewallRuleActionJSON `json:"-"`
}

// firewallRuleActionJSON contains the JSON metadata for the struct
// [FirewallRuleAction]
type firewallRuleActionJSON struct {
	Allow       apijson.Field
	Block       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleActionJSON) RawJSON() string {
	return r.raw
}

// The criteria of an incoming web request and the models of the various values
// those criteria can take
type FirewallRuleCondition struct {
	// Match the incoming request against a single IP address
	IP IPCondition `json:"ip,nullable"`
	// Match the incoming request against an IP range
	IPRange IPRangeCondition          `json:"ip_range,nullable"`
	JSON    firewallRuleConditionJSON `json:"-"`
}

// firewallRuleConditionJSON contains the JSON metadata for the struct
// [FirewallRuleCondition]
type firewallRuleConditionJSON struct {
	IP          apijson.Field
	IPRange     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleConditionJSON) RawJSON() string {
	return r.raw
}

// The criteria of an incoming web request and the models of the various values
// those criteria can take
type FirewallRuleConditionParam struct {
	// Match the incoming request against a single IP address
	IP param.Field[IPConditionParam] `json:"ip"`
	// Match the incoming request against an IP range
	IPRange param.Field[IPRangeConditionParam] `json:"ip_range"`
}

func (r FirewallRuleConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaapV1DomainFirewallRuleListResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []FirewallRule                           `json:"results,required"`
	JSON    waapV1DomainFirewallRuleListResponseJSON `json:"-"`
}

// waapV1DomainFirewallRuleListResponseJSON contains the JSON metadata for the
// struct [WaapV1DomainFirewallRuleListResponse]
type waapV1DomainFirewallRuleListResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainFirewallRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainFirewallRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainFirewallRuleNewParams struct {
	// The action that a firewall rule takes when triggered
	Action param.Field[WaapV1DomainFirewallRuleNewParamsAction] `json:"action,required"`
	// The condition required for the WAAP engine to trigger the rule.
	Conditions param.Field[[]FirewallRuleConditionParam] `json:"conditions,required"`
	// Whether or not the rule is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// The name assigned to the rule
	Name param.Field[string] `json:"name,required"`
	// The description assigned to the rule
	Description param.Field[string] `json:"description"`
}

func (r WaapV1DomainFirewallRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action that a firewall rule takes when triggered
type WaapV1DomainFirewallRuleNewParamsAction struct {
	// The WAAP allowed the request
	Allow param.Field[RuleAllowActionParam] `json:"allow"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block param.Field[RuleBlockActionParam] `json:"block"`
}

func (r WaapV1DomainFirewallRuleNewParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaapV1DomainFirewallRuleUpdateParams struct {
	// The action that a WAAP rule takes when triggered
	Action param.Field[CustomerRuleActionInputParam] `json:"action"`
	// The condition required for the WAAP engine to trigger the rule.
	Conditions param.Field[[]FirewallRuleConditionParam] `json:"conditions"`
	// The description assigned to the rule
	Description param.Field[string] `json:"description"`
	// Whether or not the rule is enabled
	Enabled param.Field[bool] `json:"enabled"`
	// The name assigned to the rule
	Name param.Field[string] `json:"name"`
}

func (r WaapV1DomainFirewallRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaapV1DomainFirewallRuleListParams struct {
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
	Ordering param.Field[WaapV1DomainFirewallRuleListParamsOrdering] `query:"ordering"`
}

// URLQuery serializes [WaapV1DomainFirewallRuleListParams]'s query parameters as
// `url.Values`.
func (r WaapV1DomainFirewallRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Determine the field to order results by
type WaapV1DomainFirewallRuleListParamsOrdering string

const (
	WaapV1DomainFirewallRuleListParamsOrderingID               WaapV1DomainFirewallRuleListParamsOrdering = "id"
	WaapV1DomainFirewallRuleListParamsOrderingName             WaapV1DomainFirewallRuleListParamsOrdering = "name"
	WaapV1DomainFirewallRuleListParamsOrderingDescription      WaapV1DomainFirewallRuleListParamsOrdering = "description"
	WaapV1DomainFirewallRuleListParamsOrderingEnabled          WaapV1DomainFirewallRuleListParamsOrdering = "enabled"
	WaapV1DomainFirewallRuleListParamsOrderingAction           WaapV1DomainFirewallRuleListParamsOrdering = "action"
	WaapV1DomainFirewallRuleListParamsOrderingMinusID          WaapV1DomainFirewallRuleListParamsOrdering = "-id"
	WaapV1DomainFirewallRuleListParamsOrderingMinusName        WaapV1DomainFirewallRuleListParamsOrdering = "-name"
	WaapV1DomainFirewallRuleListParamsOrderingMinusDescription WaapV1DomainFirewallRuleListParamsOrdering = "-description"
	WaapV1DomainFirewallRuleListParamsOrderingMinusEnabled     WaapV1DomainFirewallRuleListParamsOrdering = "-enabled"
	WaapV1DomainFirewallRuleListParamsOrderingMinusAction      WaapV1DomainFirewallRuleListParamsOrdering = "-action"
)

func (r WaapV1DomainFirewallRuleListParamsOrdering) IsKnown() bool {
	switch r {
	case WaapV1DomainFirewallRuleListParamsOrderingID, WaapV1DomainFirewallRuleListParamsOrderingName, WaapV1DomainFirewallRuleListParamsOrderingDescription, WaapV1DomainFirewallRuleListParamsOrderingEnabled, WaapV1DomainFirewallRuleListParamsOrderingAction, WaapV1DomainFirewallRuleListParamsOrderingMinusID, WaapV1DomainFirewallRuleListParamsOrderingMinusName, WaapV1DomainFirewallRuleListParamsOrderingMinusDescription, WaapV1DomainFirewallRuleListParamsOrderingMinusEnabled, WaapV1DomainFirewallRuleListParamsOrderingMinusAction:
		return true
	}
	return false
}

type WaapV1DomainFirewallRuleBulkDeleteParams struct {
	// A request to delete a list of rules
	RulesBulkDelete RulesBulkDeleteParam `json:"rules_bulk_delete,required"`
}

func (r WaapV1DomainFirewallRuleBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RulesBulkDelete)
}
