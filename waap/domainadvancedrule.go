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

// DomainAdvancedRuleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainAdvancedRuleService] method instead.
type DomainAdvancedRuleService struct {
	Options []option.RequestOption
}

// NewDomainAdvancedRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDomainAdvancedRuleService(opts ...option.RequestOption) (r DomainAdvancedRuleService) {
	r = DomainAdvancedRuleService{}
	r.Options = opts
	return
}

// Create an advanced rule
func (r *DomainAdvancedRuleService) New(ctx context.Context, domainID int64, body DomainAdvancedRuleNewParams, opts ...option.RequestOption) (res *WaapAdvancedRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/advanced-rules", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Only properties present in the request will be updated
func (r *DomainAdvancedRuleService) Update(ctx context.Context, ruleID int64, params DomainAdvancedRuleUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/advanced-rules/%v", params.DomainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return
}

// Retrieve a list of advanced rules assigned to a domain, offering filter,
// ordering, and pagination capabilities
func (r *DomainAdvancedRuleService) List(ctx context.Context, domainID int64, query DomainAdvancedRuleListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapAdvancedRule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/advanced-rules", domainID)
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

// Retrieve a list of advanced rules assigned to a domain, offering filter,
// ordering, and pagination capabilities
func (r *DomainAdvancedRuleService) ListAutoPaging(ctx context.Context, domainID int64, query DomainAdvancedRuleListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapAdvancedRule] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, domainID, query, opts...))
}

// Delete an advanced rule
func (r *DomainAdvancedRuleService) Delete(ctx context.Context, ruleID int64, body DomainAdvancedRuleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/advanced-rules/%v", body.DomainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve a specific advanced rule assigned to a domain
func (r *DomainAdvancedRuleService) Get(ctx context.Context, ruleID int64, query DomainAdvancedRuleGetParams, opts ...option.RequestOption) (res *WaapAdvancedRule, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/advanced-rules/%v", query.DomainID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Toggle an advanced rule
func (r *DomainAdvancedRuleService) Toggle(ctx context.Context, action DomainAdvancedRuleToggleParamsAction, body DomainAdvancedRuleToggleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/advanced-rules/%v/%v", body.DomainID, body.RuleID, action)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, nil, opts...)
	return
}

// An advanced WAAP rule applied to a domain
type WaapAdvancedRule struct {
	// The unique identifier for the rule
	ID int64 `json:"id,required"`
	// The action that the rule takes when triggered. Only one action can be set per
	// rule.
	Action WaapAdvancedRuleAction `json:"action,required"`
	// Whether or not the rule is enabled
	Enabled bool `json:"enabled,required"`
	// The name assigned to the rule
	Name string `json:"name,required"`
	// A CEL syntax expression that contains the rule's conditions. Allowed objects
	// are: request, whois, session, response, tags, `user_defined_tags`, `user_agent`,
	// `client_data`.
	//
	// More info can be found here:
	// https://gcore.com/docs/waap/waap-rules/advanced-rules
	Source string `json:"source,required"`
	// The description assigned to the rule
	Description string `json:"description"`
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
	//
	// Any of "access", "header_filter", "body_filter".
	Phase WaapAdvancedRulePhase `json:"phase,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Action      respjson.Field
		Enabled     respjson.Field
		Name        respjson.Field
		Source      respjson.Field
		Description respjson.Field
		Phase       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAdvancedRule) RawJSON() string { return r.JSON.raw }
func (r *WaapAdvancedRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the rule takes when triggered. Only one action can be set per
// rule.
type WaapAdvancedRuleAction struct {
	// The WAAP allowed the request
	Allow any `json:"allow"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block WaapAdvancedRuleActionBlock `json:"block"`
	// The WAAP presented the user with a captcha
	Captcha any `json:"captcha"`
	// The WAAP performed automatic browser validation
	Handshake any `json:"handshake"`
	// The WAAP monitored the request but took no action
	Monitor any `json:"monitor"`
	// WAAP tag action gets a list of tags to tag the request scope with
	Tag WaapAdvancedRuleActionTag `json:"tag"`
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
func (r WaapAdvancedRuleAction) RawJSON() string { return r.JSON.raw }
func (r *WaapAdvancedRuleAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WAAP block action behavior could be configured with response status code and
// action duration.
type WaapAdvancedRuleActionBlock struct {
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
func (r WaapAdvancedRuleActionBlock) RawJSON() string { return r.JSON.raw }
func (r *WaapAdvancedRuleActionBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WAAP tag action gets a list of tags to tag the request scope with
type WaapAdvancedRuleActionTag struct {
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
func (r WaapAdvancedRuleActionTag) RawJSON() string { return r.JSON.raw }
func (r *WaapAdvancedRuleActionTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
type WaapAdvancedRulePhase string

const (
	WaapAdvancedRulePhaseAccess       WaapAdvancedRulePhase = "access"
	WaapAdvancedRulePhaseHeaderFilter WaapAdvancedRulePhase = "header_filter"
	WaapAdvancedRulePhaseBodyFilter   WaapAdvancedRulePhase = "body_filter"
)

type DomainAdvancedRuleNewParams struct {
	// The action that the rule takes when triggered. Only one action can be set per
	// rule.
	Action DomainAdvancedRuleNewParamsAction `json:"action,omitzero,required"`
	// Whether or not the rule is enabled
	Enabled bool `json:"enabled,required"`
	// The name assigned to the rule
	Name string `json:"name,required"`
	// A CEL syntax expression that contains the rule's conditions. Allowed objects
	// are: request, whois, session, response, tags, `user_defined_tags`, `user_agent`,
	// `client_data`.
	//
	// More info can be found here:
	// https://gcore.com/docs/waap/waap-rules/advanced-rules
	Source string `json:"source,required"`
	// The description assigned to the rule
	Description param.Opt[string] `json:"description,omitzero"`
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
	//
	// Any of "access", "header_filter", "body_filter".
	Phase DomainAdvancedRuleNewParamsPhase `json:"phase,omitzero"`
	paramObj
}

func (r DomainAdvancedRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainAdvancedRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainAdvancedRuleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the rule takes when triggered. Only one action can be set per
// rule.
type DomainAdvancedRuleNewParamsAction struct {
	// The WAAP allowed the request
	Allow any `json:"allow,omitzero"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block DomainAdvancedRuleNewParamsActionBlock `json:"block,omitzero"`
	// The WAAP presented the user with a captcha
	Captcha any `json:"captcha,omitzero"`
	// The WAAP performed automatic browser validation
	Handshake any `json:"handshake,omitzero"`
	// The WAAP monitored the request but took no action
	Monitor any `json:"monitor,omitzero"`
	// WAAP tag action gets a list of tags to tag the request scope with
	Tag DomainAdvancedRuleNewParamsActionTag `json:"tag,omitzero"`
	paramObj
}

func (r DomainAdvancedRuleNewParamsAction) MarshalJSON() (data []byte, err error) {
	type shadow DomainAdvancedRuleNewParamsAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainAdvancedRuleNewParamsAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WAAP block action behavior could be configured with response status code and
// action duration.
type DomainAdvancedRuleNewParamsActionBlock struct {
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

func (r DomainAdvancedRuleNewParamsActionBlock) MarshalJSON() (data []byte, err error) {
	type shadow DomainAdvancedRuleNewParamsActionBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainAdvancedRuleNewParamsActionBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainAdvancedRuleNewParamsActionBlock](
		"status_code", 403, 405, 418, 429,
	)
}

// WAAP tag action gets a list of tags to tag the request scope with
//
// The property Tags is required.
type DomainAdvancedRuleNewParamsActionTag struct {
	// The list of user defined tags to tag the request with
	Tags []string `json:"tags,omitzero,required"`
	paramObj
}

func (r DomainAdvancedRuleNewParamsActionTag) MarshalJSON() (data []byte, err error) {
	type shadow DomainAdvancedRuleNewParamsActionTag
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainAdvancedRuleNewParamsActionTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
type DomainAdvancedRuleNewParamsPhase string

const (
	DomainAdvancedRuleNewParamsPhaseAccess       DomainAdvancedRuleNewParamsPhase = "access"
	DomainAdvancedRuleNewParamsPhaseHeaderFilter DomainAdvancedRuleNewParamsPhase = "header_filter"
	DomainAdvancedRuleNewParamsPhaseBodyFilter   DomainAdvancedRuleNewParamsPhase = "body_filter"
)

type DomainAdvancedRuleUpdateParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	// The description assigned to the rule
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether or not the rule is enabled
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// The name assigned to the rule
	Name param.Opt[string] `json:"name,omitzero"`
	// A CEL syntax expression that contains the rule's conditions. Allowed objects
	// are: request, whois, session, response, tags, `user_defined_tags`, `user_agent`,
	// `client_data`.
	//
	// More info can be found here:
	// https://gcore.com/docs/waap/waap-rules/advanced-rules
	Source param.Opt[string] `json:"source,omitzero"`
	// The action that a WAAP rule takes when triggered.
	Action DomainAdvancedRuleUpdateParamsAction `json:"action,omitzero"`
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
	//
	// Any of "access", "header_filter", "body_filter".
	Phase DomainAdvancedRuleUpdateParamsPhase `json:"phase,omitzero"`
	paramObj
}

func (r DomainAdvancedRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainAdvancedRuleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainAdvancedRuleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action that a WAAP rule takes when triggered.
type DomainAdvancedRuleUpdateParamsAction struct {
	// The WAAP allowed the request
	Allow any `json:"allow,omitzero"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block DomainAdvancedRuleUpdateParamsActionBlock `json:"block,omitzero"`
	// The WAAP presented the user with a captcha
	Captcha any `json:"captcha,omitzero"`
	// The WAAP performed automatic browser validation
	Handshake any `json:"handshake,omitzero"`
	// The WAAP monitored the request but took no action
	Monitor any `json:"monitor,omitzero"`
	// WAAP tag action gets a list of tags to tag the request scope with
	Tag DomainAdvancedRuleUpdateParamsActionTag `json:"tag,omitzero"`
	paramObj
}

func (r DomainAdvancedRuleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	type shadow DomainAdvancedRuleUpdateParamsAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainAdvancedRuleUpdateParamsAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WAAP block action behavior could be configured with response status code and
// action duration.
type DomainAdvancedRuleUpdateParamsActionBlock struct {
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

func (r DomainAdvancedRuleUpdateParamsActionBlock) MarshalJSON() (data []byte, err error) {
	type shadow DomainAdvancedRuleUpdateParamsActionBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainAdvancedRuleUpdateParamsActionBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DomainAdvancedRuleUpdateParamsActionBlock](
		"status_code", 403, 405, 418, 429,
	)
}

// WAAP tag action gets a list of tags to tag the request scope with
//
// The property Tags is required.
type DomainAdvancedRuleUpdateParamsActionTag struct {
	// The list of user defined tags to tag the request with
	Tags []string `json:"tags,omitzero,required"`
	paramObj
}

func (r DomainAdvancedRuleUpdateParamsActionTag) MarshalJSON() (data []byte, err error) {
	type shadow DomainAdvancedRuleUpdateParamsActionTag
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainAdvancedRuleUpdateParamsActionTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
type DomainAdvancedRuleUpdateParamsPhase string

const (
	DomainAdvancedRuleUpdateParamsPhaseAccess       DomainAdvancedRuleUpdateParamsPhase = "access"
	DomainAdvancedRuleUpdateParamsPhaseHeaderFilter DomainAdvancedRuleUpdateParamsPhase = "header_filter"
	DomainAdvancedRuleUpdateParamsPhaseBodyFilter   DomainAdvancedRuleUpdateParamsPhase = "body_filter"
)

type DomainAdvancedRuleListParams struct {
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
	// Any of "id", "name", "description", "enabled", "action", "phase", "-id",
	// "-name", "-description", "-enabled", "-action", "-phase".
	Ordering DomainAdvancedRuleListParamsOrdering `query:"ordering,omitzero" json:"-"`
	// Filter to refine results by specific actions
	//
	// Any of "allow", "block", "captcha", "handshake", "monitor", "tag".
	Action DomainAdvancedRuleListParamsAction `query:"action,omitzero" json:"-"`
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
	//
	// Any of "access", "header_filter", "body_filter".
	Phase DomainAdvancedRuleListParamsPhase `query:"phase,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainAdvancedRuleListParams]'s query parameters as
// `url.Values`.
func (r DomainAdvancedRuleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter to refine results by specific actions
type DomainAdvancedRuleListParamsAction string

const (
	DomainAdvancedRuleListParamsActionAllow     DomainAdvancedRuleListParamsAction = "allow"
	DomainAdvancedRuleListParamsActionBlock     DomainAdvancedRuleListParamsAction = "block"
	DomainAdvancedRuleListParamsActionCaptcha   DomainAdvancedRuleListParamsAction = "captcha"
	DomainAdvancedRuleListParamsActionHandshake DomainAdvancedRuleListParamsAction = "handshake"
	DomainAdvancedRuleListParamsActionMonitor   DomainAdvancedRuleListParamsAction = "monitor"
	DomainAdvancedRuleListParamsActionTag       DomainAdvancedRuleListParamsAction = "tag"
)

// Determine the field to order results by
type DomainAdvancedRuleListParamsOrdering string

const (
	DomainAdvancedRuleListParamsOrderingID               DomainAdvancedRuleListParamsOrdering = "id"
	DomainAdvancedRuleListParamsOrderingName             DomainAdvancedRuleListParamsOrdering = "name"
	DomainAdvancedRuleListParamsOrderingDescription      DomainAdvancedRuleListParamsOrdering = "description"
	DomainAdvancedRuleListParamsOrderingEnabled          DomainAdvancedRuleListParamsOrdering = "enabled"
	DomainAdvancedRuleListParamsOrderingAction           DomainAdvancedRuleListParamsOrdering = "action"
	DomainAdvancedRuleListParamsOrderingPhase            DomainAdvancedRuleListParamsOrdering = "phase"
	DomainAdvancedRuleListParamsOrderingMinusID          DomainAdvancedRuleListParamsOrdering = "-id"
	DomainAdvancedRuleListParamsOrderingMinusName        DomainAdvancedRuleListParamsOrdering = "-name"
	DomainAdvancedRuleListParamsOrderingMinusDescription DomainAdvancedRuleListParamsOrdering = "-description"
	DomainAdvancedRuleListParamsOrderingMinusEnabled     DomainAdvancedRuleListParamsOrdering = "-enabled"
	DomainAdvancedRuleListParamsOrderingMinusAction      DomainAdvancedRuleListParamsOrdering = "-action"
	DomainAdvancedRuleListParamsOrderingMinusPhase       DomainAdvancedRuleListParamsOrdering = "-phase"
)

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
type DomainAdvancedRuleListParamsPhase string

const (
	DomainAdvancedRuleListParamsPhaseAccess       DomainAdvancedRuleListParamsPhase = "access"
	DomainAdvancedRuleListParamsPhaseHeaderFilter DomainAdvancedRuleListParamsPhase = "header_filter"
	DomainAdvancedRuleListParamsPhaseBodyFilter   DomainAdvancedRuleListParamsPhase = "body_filter"
)

type DomainAdvancedRuleDeleteParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}

type DomainAdvancedRuleGetParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}

type DomainAdvancedRuleToggleParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	// The advanced rule ID
	RuleID int64 `path:"rule_id,required" json:"-"`
	paramObj
}

// Enable or disable an advanced rule
type DomainAdvancedRuleToggleParamsAction string

const (
	DomainAdvancedRuleToggleParamsActionEnable  DomainAdvancedRuleToggleParamsAction = "enable"
	DomainAdvancedRuleToggleParamsActionDisable DomainAdvancedRuleToggleParamsAction = "disable"
)
