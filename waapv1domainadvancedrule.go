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

// WaapV1DomainAdvancedRuleService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1DomainAdvancedRuleService] method instead.
type WaapV1DomainAdvancedRuleService struct {
	Options []option.RequestOption
}

// NewWaapV1DomainAdvancedRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWaapV1DomainAdvancedRuleService(opts ...option.RequestOption) (r *WaapV1DomainAdvancedRuleService) {
	r = &WaapV1DomainAdvancedRuleService{}
	r.Options = opts
	return
}

// Create an advanced rule
func (r *WaapV1DomainAdvancedRuleService) New(ctx context.Context, domainID int64, body WaapV1DomainAdvancedRuleNewParams, opts ...option.RequestOption) (res *AdvancedRule, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/advanced-rules", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific advanced rule assigned to a domain
func (r *WaapV1DomainAdvancedRuleService) Get(ctx context.Context, domainID int64, ruleID int64, opts ...option.RequestOption) (res *AdvancedRule, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/advanced-rules/%v", domainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Only properties present in the request will be updated
func (r *WaapV1DomainAdvancedRuleService) Update(ctx context.Context, domainID int64, ruleID int64, body WaapV1DomainAdvancedRuleUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/advanced-rules/%v", domainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Retrieve a list of advanced rules assigned to a domain, offering filter,
// ordering, and pagination capabilities
func (r *WaapV1DomainAdvancedRuleService) List(ctx context.Context, domainID int64, query WaapV1DomainAdvancedRuleListParams, opts ...option.RequestOption) (res *WaapV1DomainAdvancedRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/advanced-rules", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an advanced rule
func (r *WaapV1DomainAdvancedRuleService) Delete(ctx context.Context, domainID int64, ruleID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/advanced-rules/%v", domainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Toggle an advanced rule
func (r *WaapV1DomainAdvancedRuleService) Toggle(ctx context.Context, domainID int64, ruleID int64, action CustomerRuleState, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/advanced-rules/%v/%v", domainID, ruleID, action)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, nil, opts...)
	return
}

// An advanced WAAP rule applied to a domain
type AdvancedRule struct {
	// The unique identifier for the rule
	ID int64 `json:"id,required"`
	// The action that a WAAP rule takes when triggered
	Action CustomerRuleActionOutput `json:"action,required"`
	// Whether or not the rule is enabled
	Enabled bool `json:"enabled,required"`
	// The name assigned to the rule
	Name string `json:"name,required"`
	// A CEL syntax expression that contains the rule's conditions. Allowed objects
	// are: request, whois, session, response, tags, user_defined_tags, user_agent,
	// client_data.
	//
	// More info can be found here:
	// https://gcore.com/docs/waap/waap-rules/advanced-rules
	Source string `json:"source,required"`
	// The description assigned to the rule
	Description string `json:"description,nullable"`
	// The WAAP request/response phase for applying the rule. Default is "access".
	//
	// The "access" phase is responsible for modifying the request before it is sent to
	// the origin server.
	//
	// The "header_filter" phase is responsible for modifying the HTTP headers of a
	// response before they are sent back to the client.
	//
	// The "body_filter" phase is responsible for modifying the body of a response
	// before it is sent back to the client.
	Phase AdvancedRulePhase `json:"phase,nullable"`
	JSON  advancedRuleJSON  `json:"-"`
}

// advancedRuleJSON contains the JSON metadata for the struct [AdvancedRule]
type advancedRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Description apijson.Field
	Phase       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedRuleJSON) RawJSON() string {
	return r.raw
}

// The WAAP request/response phase for applying the rule. Default is "access".
//
// The "access" phase is responsible for modifying the request before it is sent to
// the origin server.
//
// The "header_filter" phase is responsible for modifying the HTTP headers of a
// response before they are sent back to the client.
//
// The "body_filter" phase is responsible for modifying the body of a response
// before it is sent back to the client.
type AdvancedRulePhase string

const (
	AdvancedRulePhaseAccess       AdvancedRulePhase = "access"
	AdvancedRulePhaseHeaderFilter AdvancedRulePhase = "header_filter"
	AdvancedRulePhaseBodyFilter   AdvancedRulePhase = "body_filter"
)

func (r AdvancedRulePhase) IsKnown() bool {
	switch r {
	case AdvancedRulePhaseAccess, AdvancedRulePhaseHeaderFilter, AdvancedRulePhaseBodyFilter:
		return true
	}
	return false
}

type WaapV1DomainAdvancedRuleListResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []AdvancedRule                           `json:"results,required"`
	JSON    waapV1DomainAdvancedRuleListResponseJSON `json:"-"`
}

// waapV1DomainAdvancedRuleListResponseJSON contains the JSON metadata for the
// struct [WaapV1DomainAdvancedRuleListResponse]
type waapV1DomainAdvancedRuleListResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainAdvancedRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainAdvancedRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainAdvancedRuleNewParams struct {
	// The action that a WAAP rule takes when triggered
	Action param.Field[CustomerRuleActionInputParam] `json:"action,required"`
	// Whether or not the rule is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// The name assigned to the rule
	Name param.Field[string] `json:"name,required"`
	// A CEL syntax expression that contains the rule's conditions. Allowed objects
	// are: request, whois, session, response, tags, user_defined_tags, user_agent,
	// client_data.
	//
	// More info can be found here:
	// https://gcore.com/docs/waap/waap-rules/advanced-rules
	Source param.Field[string] `json:"source,required"`
	// The description assigned to the rule
	Description param.Field[string] `json:"description"`
	// The WAAP request/response phase for applying the rule. Default is "access".
	//
	// The "access" phase is responsible for modifying the request before it is sent to
	// the origin server.
	//
	// The "header_filter" phase is responsible for modifying the HTTP headers of a
	// response before they are sent back to the client.
	//
	// The "body_filter" phase is responsible for modifying the body of a response
	// before it is sent back to the client.
	Phase param.Field[WaapV1DomainAdvancedRuleNewParamsPhase] `json:"phase"`
}

func (r WaapV1DomainAdvancedRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The WAAP request/response phase for applying the rule. Default is "access".
//
// The "access" phase is responsible for modifying the request before it is sent to
// the origin server.
//
// The "header_filter" phase is responsible for modifying the HTTP headers of a
// response before they are sent back to the client.
//
// The "body_filter" phase is responsible for modifying the body of a response
// before it is sent back to the client.
type WaapV1DomainAdvancedRuleNewParamsPhase string

const (
	WaapV1DomainAdvancedRuleNewParamsPhaseAccess       WaapV1DomainAdvancedRuleNewParamsPhase = "access"
	WaapV1DomainAdvancedRuleNewParamsPhaseHeaderFilter WaapV1DomainAdvancedRuleNewParamsPhase = "header_filter"
	WaapV1DomainAdvancedRuleNewParamsPhaseBodyFilter   WaapV1DomainAdvancedRuleNewParamsPhase = "body_filter"
)

func (r WaapV1DomainAdvancedRuleNewParamsPhase) IsKnown() bool {
	switch r {
	case WaapV1DomainAdvancedRuleNewParamsPhaseAccess, WaapV1DomainAdvancedRuleNewParamsPhaseHeaderFilter, WaapV1DomainAdvancedRuleNewParamsPhaseBodyFilter:
		return true
	}
	return false
}

type WaapV1DomainAdvancedRuleUpdateParams struct {
	// The action that a WAAP rule takes when triggered
	Action param.Field[CustomerRuleActionInputParam] `json:"action"`
	// The description assigned to the rule
	Description param.Field[string] `json:"description"`
	// Whether or not the rule is enabled
	Enabled param.Field[bool] `json:"enabled"`
	// The name assigned to the rule
	Name param.Field[string] `json:"name"`
	// The WAAP request/response phase for applying the rule.
	//
	// The "access" phase is responsible for modifying the request before it is sent to
	// the origin server.
	//
	// The "header_filter" phase is responsible for modifying the HTTP headers of a
	// response before they are sent back to the client.
	//
	// The "body_filter" phase is responsible for modifying the body of a response
	// before it is sent back to the client.
	Phase param.Field[WaapV1DomainAdvancedRuleUpdateParamsPhase] `json:"phase"`
	// A CEL syntax expression that contains the rule's conditions. Allowed objects
	// are: request, whois, session, response, tags, user_defined_tags, user_agent,
	// client_data.
	//
	// More info can be found here:
	// https://gcore.com/docs/waap/waap-rules/advanced-rules
	Source param.Field[string] `json:"source"`
}

func (r WaapV1DomainAdvancedRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The WAAP request/response phase for applying the rule.
//
// The "access" phase is responsible for modifying the request before it is sent to
// the origin server.
//
// The "header_filter" phase is responsible for modifying the HTTP headers of a
// response before they are sent back to the client.
//
// The "body_filter" phase is responsible for modifying the body of a response
// before it is sent back to the client.
type WaapV1DomainAdvancedRuleUpdateParamsPhase string

const (
	WaapV1DomainAdvancedRuleUpdateParamsPhaseAccess       WaapV1DomainAdvancedRuleUpdateParamsPhase = "access"
	WaapV1DomainAdvancedRuleUpdateParamsPhaseHeaderFilter WaapV1DomainAdvancedRuleUpdateParamsPhase = "header_filter"
	WaapV1DomainAdvancedRuleUpdateParamsPhaseBodyFilter   WaapV1DomainAdvancedRuleUpdateParamsPhase = "body_filter"
)

func (r WaapV1DomainAdvancedRuleUpdateParamsPhase) IsKnown() bool {
	switch r {
	case WaapV1DomainAdvancedRuleUpdateParamsPhaseAccess, WaapV1DomainAdvancedRuleUpdateParamsPhaseHeaderFilter, WaapV1DomainAdvancedRuleUpdateParamsPhaseBodyFilter:
		return true
	}
	return false
}

type WaapV1DomainAdvancedRuleListParams struct {
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
	Ordering param.Field[WaapV1DomainAdvancedRuleListParamsOrdering] `query:"ordering"`
	// Filter rules based on the WAAP request/response phase for applying the rule.
	//
	// The "access" phase is responsible for modifying the request before it is sent to
	// the origin server.
	//
	// The "header_filter" phase is responsible for modifying the HTTP headers of a
	// response before they are sent back to the client.
	//
	// The "body_filter" phase is responsible for modifying the body of a response
	// before it is sent back to the client.
	Phase param.Field[WaapV1DomainAdvancedRuleListParamsPhase] `query:"phase"`
}

// URLQuery serializes [WaapV1DomainAdvancedRuleListParams]'s query parameters as
// `url.Values`.
func (r WaapV1DomainAdvancedRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Determine the field to order results by
type WaapV1DomainAdvancedRuleListParamsOrdering string

const (
	WaapV1DomainAdvancedRuleListParamsOrderingID               WaapV1DomainAdvancedRuleListParamsOrdering = "id"
	WaapV1DomainAdvancedRuleListParamsOrderingName             WaapV1DomainAdvancedRuleListParamsOrdering = "name"
	WaapV1DomainAdvancedRuleListParamsOrderingDescription      WaapV1DomainAdvancedRuleListParamsOrdering = "description"
	WaapV1DomainAdvancedRuleListParamsOrderingEnabled          WaapV1DomainAdvancedRuleListParamsOrdering = "enabled"
	WaapV1DomainAdvancedRuleListParamsOrderingAction           WaapV1DomainAdvancedRuleListParamsOrdering = "action"
	WaapV1DomainAdvancedRuleListParamsOrderingPhase            WaapV1DomainAdvancedRuleListParamsOrdering = "phase"
	WaapV1DomainAdvancedRuleListParamsOrderingMinusID          WaapV1DomainAdvancedRuleListParamsOrdering = "-id"
	WaapV1DomainAdvancedRuleListParamsOrderingMinusName        WaapV1DomainAdvancedRuleListParamsOrdering = "-name"
	WaapV1DomainAdvancedRuleListParamsOrderingMinusDescription WaapV1DomainAdvancedRuleListParamsOrdering = "-description"
	WaapV1DomainAdvancedRuleListParamsOrderingMinusEnabled     WaapV1DomainAdvancedRuleListParamsOrdering = "-enabled"
	WaapV1DomainAdvancedRuleListParamsOrderingMinusAction      WaapV1DomainAdvancedRuleListParamsOrdering = "-action"
	WaapV1DomainAdvancedRuleListParamsOrderingMinusPhase       WaapV1DomainAdvancedRuleListParamsOrdering = "-phase"
)

func (r WaapV1DomainAdvancedRuleListParamsOrdering) IsKnown() bool {
	switch r {
	case WaapV1DomainAdvancedRuleListParamsOrderingID, WaapV1DomainAdvancedRuleListParamsOrderingName, WaapV1DomainAdvancedRuleListParamsOrderingDescription, WaapV1DomainAdvancedRuleListParamsOrderingEnabled, WaapV1DomainAdvancedRuleListParamsOrderingAction, WaapV1DomainAdvancedRuleListParamsOrderingPhase, WaapV1DomainAdvancedRuleListParamsOrderingMinusID, WaapV1DomainAdvancedRuleListParamsOrderingMinusName, WaapV1DomainAdvancedRuleListParamsOrderingMinusDescription, WaapV1DomainAdvancedRuleListParamsOrderingMinusEnabled, WaapV1DomainAdvancedRuleListParamsOrderingMinusAction, WaapV1DomainAdvancedRuleListParamsOrderingMinusPhase:
		return true
	}
	return false
}

// Filter rules based on the WAAP request/response phase for applying the rule.
//
// The "access" phase is responsible for modifying the request before it is sent to
// the origin server.
//
// The "header_filter" phase is responsible for modifying the HTTP headers of a
// response before they are sent back to the client.
//
// The "body_filter" phase is responsible for modifying the body of a response
// before it is sent back to the client.
type WaapV1DomainAdvancedRuleListParamsPhase string

const (
	WaapV1DomainAdvancedRuleListParamsPhaseAccess       WaapV1DomainAdvancedRuleListParamsPhase = "access"
	WaapV1DomainAdvancedRuleListParamsPhaseHeaderFilter WaapV1DomainAdvancedRuleListParamsPhase = "header_filter"
	WaapV1DomainAdvancedRuleListParamsPhaseBodyFilter   WaapV1DomainAdvancedRuleListParamsPhase = "body_filter"
)

func (r WaapV1DomainAdvancedRuleListParamsPhase) IsKnown() bool {
	switch r {
	case WaapV1DomainAdvancedRuleListParamsPhaseAccess, WaapV1DomainAdvancedRuleListParamsPhaseHeaderFilter, WaapV1DomainAdvancedRuleListParamsPhaseBodyFilter:
		return true
	}
	return false
}
