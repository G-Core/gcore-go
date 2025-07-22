// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// WaapService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapService] method instead.
type WaapService struct {
	Options        []option.RequestOption
	Statistics     StatisticService
	Domains        DomainService
	CustomPageSets CustomPageSetService
	AdvancedRules  AdvancedRuleService
	Tags           TagService
	Organizations  OrganizationService
	IPInfo         IPInfoService
}

// NewWaapService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWaapService(opts ...option.RequestOption) (r WaapService) {
	r = WaapService{}
	r.Options = opts
	r.Statistics = NewStatisticService(opts...)
	r.Domains = NewDomainService(opts...)
	r.CustomPageSets = NewCustomPageSetService(opts...)
	r.AdvancedRules = NewAdvancedRuleService(opts...)
	r.Tags = NewTagService(opts...)
	r.Organizations = NewOrganizationService(opts...)
	r.IPInfo = NewIPInfoService(opts...)
	return
}

// Get information about WAAP service for the client
func (r *WaapService) GetAccountOverview(ctx context.Context, opts ...option.RequestOption) (res *WaapGetAccountOverviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/clients/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// An advanced WAAP rule applied to a domain
type WaapAdvancedRule struct {
	// The unique identifier for the rule
	ID int64 `json:"id,required"`
	// The action that a WAAP rule takes when triggered
	Action WaapAdvancedRuleAction `json:"action,required"`
	// Whether or not the rule is enabled
	Enabled bool `json:"enabled,required"`
	// The name assigned to the rule
	Name string `json:"name,required"`
	// A CEL syntax expression that contains the rule's conditions. Allowed objects
	// are: request, whois, session, response, tags, `user_defined_tags`, `user_agent`,
	// `client_data`. More info can be found here:
	// https://gcore.com/docs/waap/waap-rules/advanced-rules
	Source string `json:"source,required"`
	// The description assigned to the rule
	Description string `json:"description,nullable"`
	// The WAAP request/response phase for applying the rule. Default is "access". The
	// "access" phase is responsible for modifying the request before it is sent to the
	// origin server. The "`header_filter`" phase is responsible for modifying the HTTP
	// headers of a response before they are sent back to the client. The
	// "`body_filter`" phase is responsible for modifying the body of a response before
	// it is sent back to the client.
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

// The action that a WAAP rule takes when triggered
type WaapAdvancedRuleAction struct {
	// The WAAP allowed the request
	Allow any `json:"allow,nullable"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block WaapAdvancedRuleActionBlock `json:"block,nullable"`
	// The WAAP presented the user with a captcha
	Captcha any `json:"captcha,nullable"`
	// The WAAP performed automatic browser validation
	Handshake any `json:"handshake,nullable"`
	// The WAAP monitored the request but took no action
	Monitor any `json:"monitor,nullable"`
	// WAAP tag action gets a list of tags to tag the request scope with
	Tag WaapAdvancedRuleActionTag `json:"tag,nullable"`
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
	// represent time format (seconds, minutes, hours, or days)
	ActionDuration string `json:"action_duration,nullable"`
	// Designates the HTTP status code to deliver when a request is blocked.
	//
	// Any of 403, 405, 418, 429.
	StatusCode int64 `json:"status_code,nullable"`
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

// The WAAP request/response phase for applying the rule. Default is "access". The
// "access" phase is responsible for modifying the request before it is sent to the
// origin server. The "`header_filter`" phase is responsible for modifying the HTTP
// headers of a response before they are sent back to the client. The
// "`body_filter`" phase is responsible for modifying the body of a response before
// it is sent back to the client.
type WaapAdvancedRulePhase string

const (
	WaapAdvancedRulePhaseAccess       WaapAdvancedRulePhase = "access"
	WaapAdvancedRulePhaseHeaderFilter WaapAdvancedRulePhase = "header_filter"
	WaapAdvancedRulePhaseBodyFilter   WaapAdvancedRulePhase = "body_filter"
)

// Advanced rules descriptor object
type WaapAdvancedRuleDescriptor struct {
	// The object's name
	Name string `json:"name,required"`
	// The object's type
	Type string `json:"type,required"`
	// The object's attributes list
	Attrs []WaapAdvancedRuleDescriptorAttr `json:"attrs,nullable"`
	// The object's description
	Description string `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Attrs       respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAdvancedRuleDescriptor) RawJSON() string { return r.JSON.raw }
func (r *WaapAdvancedRuleDescriptor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An attribute of a descriptor's object
type WaapAdvancedRuleDescriptorAttr struct {
	// The attribute's name
	Name string `json:"name,required"`
	// The attribute's type
	Type string `json:"type,required"`
	// A list of arguments for the attribute
	Args []WaapAdvancedRuleDescriptorAttrArg `json:"args,nullable"`
	// The attribute's description
	Description string `json:"description,nullable"`
	// The attribute's hint
	Hint string `json:"hint,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Args        respjson.Field
		Description respjson.Field
		Hint        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAdvancedRuleDescriptorAttr) RawJSON() string { return r.JSON.raw }
func (r *WaapAdvancedRuleDescriptorAttr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An argument of a descriptor's object
type WaapAdvancedRuleDescriptorAttrArg struct {
	// The argument's name
	Name string `json:"name,required"`
	// The argument's type
	Type string `json:"type,required"`
	// The argument's description
	Description string `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAdvancedRuleDescriptorAttrArg) RawJSON() string { return r.JSON.raw }
func (r *WaapAdvancedRuleDescriptorAttrArg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response from a request to retrieve an advanced rules descriptor
type WaapAdvancedRuleDescriptorList struct {
	// The descriptor's version
	Version string                       `json:"version,required"`
	Objects []WaapAdvancedRuleDescriptor `json:"objects,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		Objects     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapAdvancedRuleDescriptorList) RawJSON() string { return r.JSON.raw }
func (r *WaapAdvancedRuleDescriptorList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapBlockCsrfPageData struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header string `json:"header"`
	// Supported image types are JPEG, PNG and JPG, size is limited to width 450px,
	// height 130px. This should be a base 64 encoding of the full HTML img tag
	// compatible image, with the header included.
	Logo string `json:"logo"`
	// The text to display in the body of the custom page
	Text string `json:"text"`
	// The text to display in the title of the custom page
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Header      respjson.Field
		Logo        respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapBlockCsrfPageData) RawJSON() string { return r.JSON.raw }
func (r *WaapBlockCsrfPageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WaapBlockCsrfPageData to a WaapBlockCsrfPageDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WaapBlockCsrfPageDataParam.Overrides()
func (r WaapBlockCsrfPageData) ToParam() WaapBlockCsrfPageDataParam {
	return param.Override[WaapBlockCsrfPageDataParam](json.RawMessage(r.RawJSON()))
}

// The property Enabled is required.
type WaapBlockCsrfPageDataParam struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header param.Opt[string] `json:"header,omitzero"`
	// Supported image types are JPEG, PNG and JPG, size is limited to width 450px,
	// height 130px. This should be a base 64 encoding of the full HTML img tag
	// compatible image, with the header included.
	Logo param.Opt[string] `json:"logo,omitzero"`
	// The text to display in the body of the custom page
	Text param.Opt[string] `json:"text,omitzero"`
	// The text to display in the title of the custom page
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r WaapBlockCsrfPageDataParam) MarshalJSON() (data []byte, err error) {
	type shadow WaapBlockCsrfPageDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaapBlockCsrfPageDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapBlockPageData struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header string `json:"header"`
	// Supported image types are JPEG, PNG and JPG, size is limited to width 450px,
	// height 130px. This should be a base 64 encoding of the full HTML img tag
	// compatible image, with the header included.
	Logo string `json:"logo"`
	// The text to display in the body of the custom page
	Text string `json:"text"`
	// The text to display in the title of the custom page
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Header      respjson.Field
		Logo        respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapBlockPageData) RawJSON() string { return r.JSON.raw }
func (r *WaapBlockPageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WaapBlockPageData to a WaapBlockPageDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WaapBlockPageDataParam.Overrides()
func (r WaapBlockPageData) ToParam() WaapBlockPageDataParam {
	return param.Override[WaapBlockPageDataParam](json.RawMessage(r.RawJSON()))
}

// The property Enabled is required.
type WaapBlockPageDataParam struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header param.Opt[string] `json:"header,omitzero"`
	// Supported image types are JPEG, PNG and JPG, size is limited to width 450px,
	// height 130px. This should be a base 64 encoding of the full HTML img tag
	// compatible image, with the header included.
	Logo param.Opt[string] `json:"logo,omitzero"`
	// The text to display in the body of the custom page
	Text param.Opt[string] `json:"text,omitzero"`
	// The text to display in the title of the custom page
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r WaapBlockPageDataParam) MarshalJSON() (data []byte, err error) {
	type shadow WaapBlockPageDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaapBlockPageDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of total numbers of events with blocked results per criteria
type WaapBlockedStatistics struct {
	// A collection of event counts per action. The first item is the action's
	// abbreviation/full action name, and the second item is the number of events
	Action [][]WaapBlockedStatisticsActionUnion `json:"action,required"`
	// A collection of event counts per country of origin. The first item is the
	// country's ISO 3166-1 alpha-2, and the second item is the number of events
	Country [][]WaapBlockedStatisticsCountryUnion `json:"country,required"`
	// A collection of event counts per organization that owns the event's client IP.
	// The first item is the organization's name, and the second item is the number of
	// events
	Org [][]WaapBlockedStatisticsOrgUnion `json:"org,required"`
	// A collection of event counts per rule that triggered the event. The first item
	// is the rule's name, and the second item is the number of events
	RuleName [][]WaapBlockedStatisticsRuleNameUnion `json:"rule_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Country     respjson.Field
		Org         respjson.Field
		RuleName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapBlockedStatistics) RawJSON() string { return r.JSON.raw }
func (r *WaapBlockedStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapBlockedStatisticsActionUnion contains all possible properties and values
// from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapBlockedStatisticsActionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapBlockedStatisticsActionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapBlockedStatisticsActionUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapBlockedStatisticsActionUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapBlockedStatisticsActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapBlockedStatisticsCountryUnion contains all possible properties and values
// from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapBlockedStatisticsCountryUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapBlockedStatisticsCountryUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapBlockedStatisticsCountryUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapBlockedStatisticsCountryUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapBlockedStatisticsCountryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapBlockedStatisticsOrgUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapBlockedStatisticsOrgUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapBlockedStatisticsOrgUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapBlockedStatisticsOrgUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapBlockedStatisticsOrgUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapBlockedStatisticsOrgUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapBlockedStatisticsRuleNameUnion contains all possible properties and values
// from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapBlockedStatisticsRuleNameUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapBlockedStatisticsRuleNameUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapBlockedStatisticsRuleNameUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapBlockedStatisticsRuleNameUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapBlockedStatisticsRuleNameUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapCaptchaPageData struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// Error message
	Error string `json:"error"`
	// The text to display in the header of the custom page
	Header string `json:"header"`
	// Supported image types are JPEG, PNG and JPG, size is limited to width 450px,
	// height 130px. This should be a base 64 encoding of the full HTML img tag
	// compatible image, with the header included.
	Logo string `json:"logo"`
	// The text to display in the body of the custom page
	Text string `json:"text"`
	// The text to display in the title of the custom page
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Error       respjson.Field
		Header      respjson.Field
		Logo        respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCaptchaPageData) RawJSON() string { return r.JSON.raw }
func (r *WaapCaptchaPageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WaapCaptchaPageData to a WaapCaptchaPageDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WaapCaptchaPageDataParam.Overrides()
func (r WaapCaptchaPageData) ToParam() WaapCaptchaPageDataParam {
	return param.Override[WaapCaptchaPageDataParam](json.RawMessage(r.RawJSON()))
}

// The property Enabled is required.
type WaapCaptchaPageDataParam struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// Error message
	Error param.Opt[string] `json:"error,omitzero"`
	// The text to display in the header of the custom page
	Header param.Opt[string] `json:"header,omitzero"`
	// Supported image types are JPEG, PNG and JPG, size is limited to width 450px,
	// height 130px. This should be a base 64 encoding of the full HTML img tag
	// compatible image, with the header included.
	Logo param.Opt[string] `json:"logo,omitzero"`
	// The text to display in the body of the custom page
	Text param.Opt[string] `json:"text,omitzero"`
	// The text to display in the title of the custom page
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r WaapCaptchaPageDataParam) MarshalJSON() (data []byte, err error) {
	type shadow WaapCaptchaPageDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaapCaptchaPageDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Common tag details
type WaapCommonTag struct {
	// Tag description information
	Description string `json:"description,required"`
	// The tag's display name
	DisplayName string `json:"display_name,required"`
	// Tag name
	Tag string `json:"tag,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		DisplayName respjson.Field
		Tag         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCommonTag) RawJSON() string { return r.JSON.raw }
func (r *WaapCommonTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapCookieDisabledPageData struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header string `json:"header"`
	// The text to display in the body of the custom page
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Header      respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCookieDisabledPageData) RawJSON() string { return r.JSON.raw }
func (r *WaapCookieDisabledPageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WaapCookieDisabledPageData to a
// WaapCookieDisabledPageDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WaapCookieDisabledPageDataParam.Overrides()
func (r WaapCookieDisabledPageData) ToParam() WaapCookieDisabledPageDataParam {
	return param.Override[WaapCookieDisabledPageDataParam](json.RawMessage(r.RawJSON()))
}

// The property Enabled is required.
type WaapCookieDisabledPageDataParam struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header param.Opt[string] `json:"header,omitzero"`
	// The text to display in the body of the custom page
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r WaapCookieDisabledPageDataParam) MarshalJSON() (data []byte, err error) {
	type shadow WaapCookieDisabledPageDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaapCookieDisabledPageDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of total numbers of events per criteria
type WaapCountStatistics struct {
	// A collection of event counts per action. The first item is the action's
	// abbreviation/full action name, and the second item is the number of events
	Action [][]WaapCountStatisticsActionUnion `json:"action,required"`
	// A collection of event counts per country of origin. The first item is the
	// country's ISO 3166-1 alpha-2, and the second item is the number of events
	Country [][]WaapCountStatisticsCountryUnion `json:"country,required"`
	// A collection of event counts per organization that owns the event's client IP.
	// The first item is the organization's name, and the second item is the number of
	// events
	Org [][]WaapCountStatisticsOrgUnion `json:"org,required"`
	// A collection of event counts per rule that triggered the event. The first item
	// is the rule's name, and the second item is the number of events
	RuleName [][]WaapCountStatisticsRuleNameUnion `json:"rule_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Country     respjson.Field
		Org         respjson.Field
		RuleName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCountStatistics) RawJSON() string { return r.JSON.raw }
func (r *WaapCountStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapCountStatisticsActionUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapCountStatisticsActionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapCountStatisticsActionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapCountStatisticsActionUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapCountStatisticsActionUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapCountStatisticsActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapCountStatisticsCountryUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapCountStatisticsCountryUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapCountStatisticsCountryUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapCountStatisticsCountryUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapCountStatisticsCountryUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapCountStatisticsCountryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapCountStatisticsOrgUnion contains all possible properties and values from
// [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapCountStatisticsOrgUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapCountStatisticsOrgUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapCountStatisticsOrgUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapCountStatisticsOrgUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapCountStatisticsOrgUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WaapCountStatisticsRuleNameUnion contains all possible properties and values
// from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type WaapCountStatisticsRuleNameUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u WaapCountStatisticsRuleNameUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WaapCountStatisticsRuleNameUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WaapCountStatisticsRuleNameUnion) RawJSON() string { return u.JSON.raw }

func (r *WaapCountStatisticsRuleNameUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapCustomPagePreview struct {
	// HTML content of the custom page
	HTML string `json:"html,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HTML        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomPagePreview) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomPagePreview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapCustomPageSet struct {
	// The ID of the custom page set
	ID int64 `json:"id,required"`
	// Name of the custom page set
	Name           string                     `json:"name,required"`
	Block          WaapBlockPageData          `json:"block,nullable"`
	BlockCsrf      WaapBlockCsrfPageData      `json:"block_csrf,nullable"`
	Captcha        WaapCaptchaPageData        `json:"captcha,nullable"`
	CookieDisabled WaapCookieDisabledPageData `json:"cookie_disabled,nullable"`
	// List of domain IDs that are associated with this page set
	Domains            []int64                        `json:"domains,nullable"`
	Handshake          WaapHandshakePageData          `json:"handshake,nullable"`
	JavascriptDisabled WaapJavascriptDisabledPageData `json:"javascript_disabled,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Name               respjson.Field
		Block              respjson.Field
		BlockCsrf          respjson.Field
		Captcha            respjson.Field
		CookieDisabled     respjson.Field
		Domains            respjson.Field
		Handshake          respjson.Field
		JavascriptDisabled respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapCustomPageSet) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomPageSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An WAAP rule applied to a domain
type WaapCustomRule struct {
	// The unique identifier for the rule
	ID int64 `json:"id,required"`
	// The action that a WAAP rule takes when triggered
	Action WaapCustomRuleAction `json:"action,required"`
	// The conditions required for the WAAP engine to trigger the rule. Rules may have
	// between 1 and 5 conditions. All conditions must pass for the rule to trigger
	Conditions []WaapCustomRuleCondition `json:"conditions,required"`
	// Whether or not the rule is enabled
	Enabled bool `json:"enabled,required"`
	// The name assigned to the rule
	Name string `json:"name,required"`
	// The description assigned to the rule
	Description string `json:"description,nullable"`
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

// The action that a WAAP rule takes when triggered
type WaapCustomRuleAction struct {
	// The WAAP allowed the request
	Allow any `json:"allow,nullable"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block WaapCustomRuleActionBlock `json:"block,nullable"`
	// The WAAP presented the user with a captcha
	Captcha any `json:"captcha,nullable"`
	// The WAAP performed automatic browser validation
	Handshake any `json:"handshake,nullable"`
	// The WAAP monitored the request but took no action
	Monitor any `json:"monitor,nullable"`
	// WAAP tag action gets a list of tags to tag the request scope with
	Tag WaapCustomRuleActionTag `json:"tag,nullable"`
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
	// represent time format (seconds, minutes, hours, or days)
	ActionDuration string `json:"action_duration,nullable"`
	// Designates the HTTP status code to deliver when a request is blocked.
	//
	// Any of 403, 405, 418, 429.
	StatusCode int64 `json:"status_code,nullable"`
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
	ContentType WaapCustomRuleConditionContentType `json:"content_type,nullable"`
	// Match the country that the request originated from
	Country WaapCustomRuleConditionCountry `json:"country,nullable"`
	// Match the incoming file extension
	FileExtension WaapCustomRuleConditionFileExtension `json:"file_extension,nullable"`
	// Match an incoming request header
	Header WaapCustomRuleConditionHeader `json:"header,nullable"`
	// Match when an incoming request header is present
	HeaderExists WaapCustomRuleConditionHeaderExists `json:"header_exists,nullable"`
	// Match the incoming HTTP method
	HTTPMethod WaapCustomRuleConditionHTTPMethod `json:"http_method,nullable"`
	// Match the incoming request against a single IP address
	IP WaapCustomRuleConditionIP `json:"ip,nullable"`
	// Match the incoming request against an IP range
	IPRange WaapCustomRuleConditionIPRange `json:"ip_range,nullable"`
	// Match the organization the request originated from, as determined by a WHOIS
	// lookup of the requesting IP
	Organization WaapCustomRuleConditionOrganization `json:"organization,nullable"`
	// Match the type of organization that owns the IP address making an incoming
	// request
	OwnerTypes WaapCustomRuleConditionOwnerTypes `json:"owner_types,nullable"`
	// Match the rate at which requests come in that match certain conditions
	RequestRate WaapCustomRuleConditionRequestRate `json:"request_rate,nullable"`
	// Match a response header
	ResponseHeader WaapCustomRuleConditionResponseHeader `json:"response_header,nullable"`
	// Match when a response header is present
	ResponseHeaderExists WaapCustomRuleConditionResponseHeaderExists `json:"response_header_exists,nullable"`
	// Match the number of dynamic page requests made in a WAAP session
	SessionRequestCount WaapCustomRuleConditionSessionRequestCount `json:"session_request_count,nullable"`
	// Matches requests based on specified tags
	Tags WaapCustomRuleConditionTags `json:"tags,nullable"`
	// Match the incoming request URL
	URL WaapCustomRuleConditionURL `json:"url,nullable"`
	// Match the user agent making the request
	UserAgent WaapCustomRuleConditionUserAgent `json:"user_agent,nullable"`
	// Matches requests based on user-defined tags
	UserDefinedTags WaapCustomRuleConditionUserDefinedTags `json:"user_defined_tags,nullable"`
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
	// HTTP methods and descriptions Methods from the following RFCs are all observed:
	//
	// - RFC 7231: Hypertext Transfer Protocol (HTTP/1.1), obsoletes 2616
	// - RFC 5789: PATCH Method for HTTP
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
	IPAddress string `json:"ip_address,required" format:"ipv4"`
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
	LowerBound string `json:"lower_bound,required" format:"ipv4"`
	// The upper bound IPv4 or IPv6 address to match against
	UpperBound string `json:"upper_bound,required" format:"ipv4"`
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
	HTTPMethods []string `json:"http_methods,nullable"`
	// A list of source IPs that can trigger a request rate condition
	IPs []string `json:"ips,nullable" format:"ipv4"`
	// A user-defined tag that can be included in incoming requests and used to trigger
	// a request rate condition
	UserDefinedTag string `json:"user_defined_tag,nullable"`
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

type WaapCustomerRuleState string

const (
	WaapCustomerRuleStateEnable  WaapCustomerRuleState = "enable"
	WaapCustomerRuleStateDisable WaapCustomerRuleState = "disable"
)

type WaapDDOSAttack struct {
	// End time of DDoS attack
	EndTime time.Time `json:"end_time,nullable" format:"date-time"`
	// Start time of DDoS attack
	StartTime time.Time `json:"start_time" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTime     respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDDOSAttack) RawJSON() string { return r.JSON.raw }
func (r *WaapDDOSAttack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapDDOSInfo struct {
	// The number of requests made
	Count int64 `json:"count,required"`
	// The value for the grouped by type
	Identity string `json:"identity,required"`
	// Any of "URL", "IP", "User-Agent".
	Type WaapDDOSInfoType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Identity    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDDOSInfo) RawJSON() string { return r.JSON.raw }
func (r *WaapDDOSInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapDDOSInfoType string

const (
	WaapDDOSInfoTypeURL       WaapDDOSInfoType = "URL"
	WaapDDOSInfoTypeIP        WaapDDOSInfoType = "IP"
	WaapDDOSInfoTypeUserAgent WaapDDOSInfoType = "User-Agent"
)

// Represents a WAAP domain, serving as a singular unit within the WAAP service.
// Each domain functions autonomously, possessing its own set of rules and
// configurations to manage web application firewall settings and behaviors.
type WaapDetailedDomain struct {
	// The domain ID
	ID int64 `json:"id,required"`
	// The date and time the domain was created in ISO 8601 format
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The ID of the custom page set
	CustomPageSet int64 `json:"custom_page_set,required"`
	// The domain name
	Name string `json:"name,required"`
	// The different statuses a domain can have
	//
	// Any of "active", "bypass", "monitor", "locked".
	Status WaapDomainStatus `json:"status,required"`
	// Domain level quotas
	Quotas map[string]WaapDetailedDomainQuota `json:"quotas,nullable"`
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

type WaapDetailedDomainQuota struct {
	// The maximum allowed number of this resource
	Allowed int64 `json:"allowed,required"`
	// The current number of this resource
	Current int64 `json:"current,required"`
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

// Represents a configurable WAAP security rule, also known as a policy.
type WaapDomainPolicy struct {
	// Unique identifier for the security rule
	ID string `json:"id,required"`
	// The action taken by the WAAP upon rule activation.
	//
	// Any of "Allow", "Block", "Captcha", "Gateway", "Handshake", "Monitor",
	// "Composite".
	Action WaapPolicyAction `json:"action,required"`
	// Detailed description of the security rule
	Description string `json:"description,required"`
	// The rule set group name to which the rule belongs
	Group string `json:"group,required"`
	// Indicates if the security rule is active
	Mode bool `json:"mode,required"`
	// Name of the security rule
	Name string `json:"name,required"`
	// Identifier of the rule set to which the rule belongs
	RuleSetID int64 `json:"rule_set_id,required"`
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
func (r WaapDomainPolicy) RawJSON() string { return r.JSON.raw }
func (r *WaapDomainPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for a domain.
type WaapDomainSettingsModel struct {
	// API settings of a domain
	API WaapDomainAPISettings `json:"api,required"`
	// DDoS settings for a domain.
	DDOS WaapDomainDDOSSettings `json:"ddos,required"`
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

// The different statuses a domain can have
type WaapDomainStatus string

const (
	WaapDomainStatusActive  WaapDomainStatus = "active"
	WaapDomainStatusBypass  WaapDomainStatus = "bypass"
	WaapDomainStatusMonitor WaapDomainStatus = "monitor"
	WaapDomainStatusLocked  WaapDomainStatus = "locked"
)

// A collection of event metrics over a time span
type WaapEventStatistics struct {
	// A collection of total numbers of events with blocked results per criteria
	Blocked WaapBlockedStatistics `json:"blocked,required"`
	// A collection of total numbers of events per criteria
	Count WaapCountStatistics `json:"count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blocked     respjson.Field
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapEventStatistics) RawJSON() string { return r.JSON.raw }
func (r *WaapEventStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapFirewallRule struct {
	// The unique identifier of the rule
	ID int64 `json:"id,required"`
	// The action that a firewall rule takes when triggered
	Action WaapFirewallRuleAction `json:"action,required"`
	// The condition required for the WAAP engine to trigger the rule.
	Conditions []WaapFirewallRuleCondition `json:"conditions,required"`
	// Whether or not the rule is enabled
	Enabled bool `json:"enabled,required"`
	// The name assigned to the rule
	Name string `json:"name,required"`
	// The description assigned to the rule
	Description string `json:"description,nullable"`
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

// The action that a firewall rule takes when triggered
type WaapFirewallRuleAction struct {
	// The WAAP allowed the request
	Allow any `json:"allow,nullable"`
	// WAAP block action behavior could be configured with response status code and
	// action duration.
	Block WaapFirewallRuleActionBlock `json:"block,nullable"`
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
	// represent time format (seconds, minutes, hours, or days)
	ActionDuration string `json:"action_duration,nullable"`
	// Designates the HTTP status code to deliver when a request is blocked.
	//
	// Any of 403, 405, 418, 429.
	StatusCode int64 `json:"status_code,nullable"`
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
	IP WaapFirewallRuleConditionIP `json:"ip,nullable"`
	// Match the incoming request against an IP range
	IPRange WaapFirewallRuleConditionIPRange `json:"ip_range,nullable"`
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
	IPAddress string `json:"ip_address,required" format:"ipv4"`
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
	LowerBound string `json:"lower_bound,required" format:"ipv4"`
	// The upper bound IPv4 or IPv6 address to match against
	UpperBound string `json:"upper_bound,required" format:"ipv4"`
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

type WaapHandshakePageData struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header string `json:"header"`
	// Supported image types are JPEG, PNG and JPG, size is limited to width 450px,
	// height 130px. This should be a base 64 encoding of the full HTML img tag
	// compatible image, with the header included.
	Logo string `json:"logo"`
	// The text to display in the title of the custom page
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Header      respjson.Field
		Logo        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapHandshakePageData) RawJSON() string { return r.JSON.raw }
func (r *WaapHandshakePageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WaapHandshakePageData to a WaapHandshakePageDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WaapHandshakePageDataParam.Overrides()
func (r WaapHandshakePageData) ToParam() WaapHandshakePageDataParam {
	return param.Override[WaapHandshakePageDataParam](json.RawMessage(r.RawJSON()))
}

// The property Enabled is required.
type WaapHandshakePageDataParam struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header param.Opt[string] `json:"header,omitzero"`
	// Supported image types are JPEG, PNG and JPG, size is limited to width 450px,
	// height 130px. This should be a base 64 encoding of the full HTML img tag
	// compatible image, with the header included.
	Logo param.Opt[string] `json:"logo,omitzero"`
	// The text to display in the title of the custom page
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r WaapHandshakePageDataParam) MarshalJSON() (data []byte, err error) {
	type shadow WaapHandshakePageDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaapHandshakePageDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapInsight struct {
	// A generated unique identifier for the insight
	ID string `json:"id,required" format:"uuid"`
	// The description of the insight
	Description string `json:"description,required"`
	// The date and time the insight was first seen in ISO 8601 format
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// The type of the insight represented as a slug
	InsightType string `json:"insight_type,required"`
	// A hash table of label names and values that apply to the insight
	Labels map[string]string `json:"labels,required"`
	// The date and time the insight was last seen in ISO 8601 format
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// The date and time the insight was last seen in ISO 8601 format
	LastStatusChange time.Time `json:"last_status_change,required" format:"date-time"`
	// The recommended action to perform to resolve the insight
	Recommendation string `json:"recommendation,required"`
	// The different statuses an insight can have
	//
	// Any of "OPEN", "ACKED", "CLOSED".
	Status WaapInsightStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Description      respjson.Field
		FirstSeen        respjson.Field
		InsightType      respjson.Field
		Labels           respjson.Field
		LastSeen         respjson.Field
		LastStatusChange respjson.Field
		Recommendation   respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapInsight) RawJSON() string { return r.JSON.raw }
func (r *WaapInsight) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapInsightSilence struct {
	// A generated unique identifier for the silence
	ID string `json:"id,required" format:"uuid"`
	// The author of the silence
	Author string `json:"author,required"`
	// A comment explaining the reason for the silence
	Comment string `json:"comment,required"`
	// The date and time the silence expires in ISO 8601 format
	ExpireAt time.Time `json:"expire_at,required" format:"date-time"`
	// The slug of the insight type
	InsightType string `json:"insight_type,required"`
	// A hash table of label names and values that apply to the insight silence
	Labels map[string]string `json:"labels,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Author      respjson.Field
		Comment     respjson.Field
		ExpireAt    respjson.Field
		InsightType respjson.Field
		Labels      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapInsightSilence) RawJSON() string { return r.JSON.raw }
func (r *WaapInsightSilence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapInsightSilenceSortBy string

const (
	WaapInsightSilenceSortByID               WaapInsightSilenceSortBy = "id"
	WaapInsightSilenceSortByMinusID          WaapInsightSilenceSortBy = "-id"
	WaapInsightSilenceSortByInsightType      WaapInsightSilenceSortBy = "insight_type"
	WaapInsightSilenceSortByMinusInsightType WaapInsightSilenceSortBy = "-insight_type"
	WaapInsightSilenceSortByComment          WaapInsightSilenceSortBy = "comment"
	WaapInsightSilenceSortByMinusComment     WaapInsightSilenceSortBy = "-comment"
	WaapInsightSilenceSortByAuthor           WaapInsightSilenceSortBy = "author"
	WaapInsightSilenceSortByMinusAuthor      WaapInsightSilenceSortBy = "-author"
	WaapInsightSilenceSortByExpireAt         WaapInsightSilenceSortBy = "expire_at"
	WaapInsightSilenceSortByMinusExpireAt    WaapInsightSilenceSortBy = "-expire_at"
)

type WaapInsightSortBy string

const (
	WaapInsightSortByID                    WaapInsightSortBy = "id"
	WaapInsightSortByMinusID               WaapInsightSortBy = "-id"
	WaapInsightSortByInsightType           WaapInsightSortBy = "insight_type"
	WaapInsightSortByMinusInsightType      WaapInsightSortBy = "-insight_type"
	WaapInsightSortByFirstSeen             WaapInsightSortBy = "first_seen"
	WaapInsightSortByMinusFirstSeen        WaapInsightSortBy = "-first_seen"
	WaapInsightSortByLastSeen              WaapInsightSortBy = "last_seen"
	WaapInsightSortByMinusLastSeen         WaapInsightSortBy = "-last_seen"
	WaapInsightSortByLastStatusChange      WaapInsightSortBy = "last_status_change"
	WaapInsightSortByMinusLastStatusChange WaapInsightSortBy = "-last_status_change"
	WaapInsightSortByStatus                WaapInsightSortBy = "status"
	WaapInsightSortByMinusStatus           WaapInsightSortBy = "-status"
)

// The different statuses an insight can have
type WaapInsightStatus string

const (
	WaapInsightStatusOpen   WaapInsightStatus = "OPEN"
	WaapInsightStatusAcked  WaapInsightStatus = "ACKED"
	WaapInsightStatusClosed WaapInsightStatus = "CLOSED"
)

type WaapIPCountryAttack struct {
	// The number of attacks from the specified IP address to the country
	Count int64 `json:"count,required"`
	// An ISO 3166-1 alpha-2 formatted string representing the country that was
	// attacked
	Country string `json:"country,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Country     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapIPCountryAttack) RawJSON() string { return r.JSON.raw }
func (r *WaapIPCountryAttack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapIPDDOSInfoModel struct {
	// Indicates if the IP is tagged as a botnet client
	BotnetClient bool `json:"botnet_client,required"`
	// The time series data for the DDoS attacks from the IP address
	TimeSeries []WaapIPDDOSInfoModelTimeSery `json:"time_series,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BotnetClient respjson.Field
		TimeSeries   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapIPDDOSInfoModel) RawJSON() string { return r.JSON.raw }
func (r *WaapIPDDOSInfoModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapIPDDOSInfoModelTimeSery struct {
	// The number of attacks
	Count int64 `json:"count,required"`
	// The timestamp of the time series item as a POSIX timestamp
	Timestamp int64 `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapIPDDOSInfoModelTimeSery) RawJSON() string { return r.JSON.raw }
func (r *WaapIPDDOSInfoModelTimeSery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapIPInfo struct {
	// The risk score of the IP address
	//
	// Any of "NO_RISK", "LOW", "MEDIUM", "HIGH", "EXTREME", "NOT_ENOUGH_DATA".
	RiskScore WaapIPInfoRiskScore `json:"risk_score,required"`
	// The tags associated with the IP address that affect the risk score
	Tags []string `json:"tags,required"`
	// The WHOIS information for the IP address
	Whois WaapIPInfoWhois `json:"whois,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RiskScore   respjson.Field
		Tags        respjson.Field
		Whois       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapIPInfo) RawJSON() string { return r.JSON.raw }
func (r *WaapIPInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The risk score of the IP address
type WaapIPInfoRiskScore string

const (
	WaapIPInfoRiskScoreNoRisk        WaapIPInfoRiskScore = "NO_RISK"
	WaapIPInfoRiskScoreLow           WaapIPInfoRiskScore = "LOW"
	WaapIPInfoRiskScoreMedium        WaapIPInfoRiskScore = "MEDIUM"
	WaapIPInfoRiskScoreHigh          WaapIPInfoRiskScore = "HIGH"
	WaapIPInfoRiskScoreExtreme       WaapIPInfoRiskScore = "EXTREME"
	WaapIPInfoRiskScoreNotEnoughData WaapIPInfoRiskScore = "NOT_ENOUGH_DATA"
)

// The WHOIS information for the IP address
type WaapIPInfoWhois struct {
	// The abuse mail
	AbuseMail string `json:"abuse_mail,nullable"`
	// The CIDR
	Cidr int64 `json:"cidr,nullable"`
	// The country
	Country string `json:"country,nullable"`
	// The network description
	NetDescription string `json:"net_description,nullable"`
	// The network name
	NetName string `json:"net_name,nullable"`
	// The network range
	NetRange string `json:"net_range,nullable"`
	// The network type
	NetType string `json:"net_type,nullable"`
	// The organization ID
	OrgID string `json:"org_id,nullable"`
	// The organization name
	OrgName string `json:"org_name,nullable"`
	// The owner type
	OwnerType string `json:"owner_type,nullable"`
	// The RIR
	Rir string `json:"rir,nullable"`
	// The state
	State string `json:"state,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AbuseMail      respjson.Field
		Cidr           respjson.Field
		Country        respjson.Field
		NetDescription respjson.Field
		NetName        respjson.Field
		NetRange       respjson.Field
		NetType        respjson.Field
		OrgID          respjson.Field
		OrgName        respjson.Field
		OwnerType      respjson.Field
		Rir            respjson.Field
		State          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapIPInfoWhois) RawJSON() string { return r.JSON.raw }
func (r *WaapIPInfoWhois) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapIPInfoCounts struct {
	// The number of requests from the IP address that were blocked
	BlockedRequests int64 `json:"blocked_requests,required"`
	// The total number of requests made by the IP address
	TotalRequests int64 `json:"total_requests,required"`
	// The number of unique sessions from the IP address
	UniqueSessions int64 `json:"unique_sessions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockedRequests respjson.Field
		TotalRequests   respjson.Field
		UniqueSessions  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapIPInfoCounts) RawJSON() string { return r.JSON.raw }
func (r *WaapIPInfoCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapJavascriptDisabledPageData struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header string `json:"header"`
	// The text to display in the body of the custom page
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Header      respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapJavascriptDisabledPageData) RawJSON() string { return r.JSON.raw }
func (r *WaapJavascriptDisabledPageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WaapJavascriptDisabledPageData to a
// WaapJavascriptDisabledPageDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WaapJavascriptDisabledPageDataParam.Overrides()
func (r WaapJavascriptDisabledPageData) ToParam() WaapJavascriptDisabledPageDataParam {
	return param.Override[WaapJavascriptDisabledPageDataParam](json.RawMessage(r.RawJSON()))
}

// The property Enabled is required.
type WaapJavascriptDisabledPageDataParam struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header param.Opt[string] `json:"header,omitzero"`
	// The text to display in the body of the custom page
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r WaapJavascriptDisabledPageDataParam) MarshalJSON() (data []byte, err error) {
	type shadow WaapJavascriptDisabledPageDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaapJavascriptDisabledPageDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Network details
type WaapNetworkDetails struct {
	// Client IP
	ClientIP string `json:"client_ip,required"`
	// Country code
	Country string `json:"country,required"`
	// Organization details
	Organization WaapRequestOrganization `json:"organization,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientIP     respjson.Field
		Country      respjson.Field
		Organization respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapNetworkDetails) RawJSON() string { return r.JSON.raw }
func (r *WaapNetworkDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an IP range owner organization
type WaapOrganization struct {
	// The ID of an organization
	ID int64 `json:"id,required"`
	// The name of an organization
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapOrganization) RawJSON() string { return r.JSON.raw }
func (r *WaapOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the custom page
type WaapPageType string

const (
	WaapPageTypeBlockHTML              WaapPageType = "block.html"
	WaapPageTypeBlockCsrfHTML          WaapPageType = "block_csrf.html"
	WaapPageTypeCaptchaHTML            WaapPageType = "captcha.html"
	WaapPageTypeCookieDisabledHTML     WaapPageType = "cookieDisabled.html"
	WaapPageTypeHandshakeHTML          WaapPageType = "handshake.html"
	WaapPageTypeJavascriptDisabledHTML WaapPageType = "javascriptDisabled.html"
)

type WaapPaginatedCustomPageSet struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []WaapCustomPageSet `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Limit       respjson.Field
		Offset      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapPaginatedCustomPageSet) RawJSON() string { return r.JSON.raw }
func (r *WaapPaginatedCustomPageSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapPaginatedDDOSAttack struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []WaapDDOSAttack `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Limit       respjson.Field
		Offset      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapPaginatedDDOSAttack) RawJSON() string { return r.JSON.raw }
func (r *WaapPaginatedDDOSAttack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapPaginatedDDOSInfo struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []WaapDDOSInfo `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Limit       respjson.Field
		Offset      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapPaginatedDDOSInfo) RawJSON() string { return r.JSON.raw }
func (r *WaapPaginatedDDOSInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapPaginatedRequestSummary struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []WaapRequestSummary `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Limit       respjson.Field
		Offset      respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapPaginatedRequestSummary) RawJSON() string { return r.JSON.raw }
func (r *WaapPaginatedRequestSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pattern matched tag details
type WaapPatternMatchedTag struct {
	// Tag description information
	Description string `json:"description,required"`
	// The tag's display name
	DisplayName string `json:"display_name,required"`
	// The phase in which the tag was triggered: access -> Request, `header_filter` ->
	// `response_header`, `body_filter` -> `response_body`
	ExecutionPhase string `json:"execution_phase,required"`
	// The entity to which the variable that triggered the tag belong to. For example:
	// `request_headers`, uri, cookies etc.
	Field string `json:"field,required"`
	// The name of the variable which holds the value that triggered the tag
	FieldName string `json:"field_name,required"`
	// The name of the detected regexp pattern
	PatternName string `json:"pattern_name,required"`
	// The pattern which triggered the tag
	PatternValue string `json:"pattern_value,required"`
	// Tag name
	Tag string `json:"tag,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		DisplayName    respjson.Field
		ExecutionPhase respjson.Field
		Field          respjson.Field
		FieldName      respjson.Field
		PatternName    respjson.Field
		PatternValue   respjson.Field
		Tag            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapPatternMatchedTag) RawJSON() string { return r.JSON.raw }
func (r *WaapPatternMatchedTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action taken by the WAAP upon rule activation.
type WaapPolicyAction string

const (
	WaapPolicyActionAllow     WaapPolicyAction = "Allow"
	WaapPolicyActionBlock     WaapPolicyAction = "Block"
	WaapPolicyActionCaptcha   WaapPolicyAction = "Captcha"
	WaapPolicyActionGateway   WaapPolicyAction = "Gateway"
	WaapPolicyActionHandshake WaapPolicyAction = "Handshake"
	WaapPolicyActionMonitor   WaapPolicyAction = "Monitor"
	WaapPolicyActionComposite WaapPolicyAction = "Composite"
)

// Represents the mode of a security rule.
type WaapPolicyMode struct {
	// Indicates if the security rule is active
	Mode bool `json:"mode,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapPolicyMode) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyMode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request's details used when displaying a single request.
type WaapRequestDetails struct {
	// Request ID
	ID string `json:"id,required"`
	// Request action
	Action string `json:"action,required"`
	// List of common tags
	CommonTags []WaapCommonTag `json:"common_tags,required"`
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
	Network WaapNetworkDetails `json:"network,required"`
	// Request path
	Path string `json:"path,required"`
	// List of shield tags
	PatternMatchedTags []WaapPatternMatchedTag `json:"pattern_matched_tags,required"`
	// The query string of the request
	QueryString string `json:"query_string,required"`
	// Reference ID to identify user sanction
	ReferenceID string `json:"reference_id,required"`
	// HTTP request headers
	RequestHeaders any `json:"request_headers,required"`
	// The time of the request
	RequestTime string `json:"request_time,required"`
	// The type of the request that generated an event
	RequestType string `json:"request_type,required"`
	// The real domain name
	RequestedDomain string `json:"requested_domain,required"`
	// Time took to process all request
	ResponseTime string `json:"response_time,required"`
	// The result of a request
	//
	// Any of "passed", "blocked", "suppressed", "".
	Result WaapRequestDetailsResult `json:"result,required"`
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
	UserAgent WaapUserAgentDetails `json:"user_agent,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Action              respjson.Field
		CommonTags          respjson.Field
		ContentType         respjson.Field
		Domain              respjson.Field
		HTTPStatusCode      respjson.Field
		HTTPVersion         respjson.Field
		IncidentID          respjson.Field
		Method              respjson.Field
		Network             respjson.Field
		Path                respjson.Field
		PatternMatchedTags  respjson.Field
		QueryString         respjson.Field
		ReferenceID         respjson.Field
		RequestHeaders      respjson.Field
		RequestTime         respjson.Field
		RequestType         respjson.Field
		RequestedDomain     respjson.Field
		ResponseTime        respjson.Field
		Result              respjson.Field
		RuleID              respjson.Field
		RuleName            respjson.Field
		Scheme              respjson.Field
		SessionRequestCount respjson.Field
		TrafficTypes        respjson.Field
		UserAgent           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRequestDetails) RawJSON() string { return r.JSON.raw }
func (r *WaapRequestDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The result of a request
type WaapRequestDetailsResult string

const (
	WaapRequestDetailsResultPassed     WaapRequestDetailsResult = "passed"
	WaapRequestDetailsResultBlocked    WaapRequestDetailsResult = "blocked"
	WaapRequestDetailsResultSuppressed WaapRequestDetailsResult = "suppressed"
	WaapRequestDetailsResultEmpty      WaapRequestDetailsResult = ""
)

// Organization details
type WaapRequestOrganization struct {
	// Organization name
	Name string `json:"name,required"`
	// Network range
	Subnet string `json:"subnet,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Subnet      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRequestOrganization) RawJSON() string { return r.JSON.raw }
func (r *WaapRequestOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request summary used when displaying a list of requests
type WaapRequestSummary struct {
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
	RequestTime int64 `json:"request_time,required"`
	// Any of "passed", "blocked", "suppressed", "".
	Result WaapRequestSummaryResult `json:"result,required"`
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
	UserAgentClient string `json:"user_agent_client,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Action          respjson.Field
		ClientIP        respjson.Field
		Country         respjson.Field
		Domain          respjson.Field
		Method          respjson.Field
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
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRequestSummary) RawJSON() string { return r.JSON.raw }
func (r *WaapRequestSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapRequestSummaryResult string

const (
	WaapRequestSummaryResultPassed     WaapRequestSummaryResult = "passed"
	WaapRequestSummaryResultBlocked    WaapRequestSummaryResult = "blocked"
	WaapRequestSummaryResultSuppressed WaapRequestSummaryResult = "suppressed"
	WaapRequestSummaryResultEmpty      WaapRequestSummaryResult = ""
)

// Specifies the granularity of the result data.
type WaapResolution string

const (
	WaapResolutionDaily    WaapResolution = "daily"
	WaapResolutionHourly   WaapResolution = "hourly"
	WaapResolutionMinutely WaapResolution = "minutely"
)

type WaapRuleActionType string

const (
	WaapRuleActionTypeAllow     WaapRuleActionType = "allow"
	WaapRuleActionTypeBlock     WaapRuleActionType = "block"
	WaapRuleActionTypeCaptcha   WaapRuleActionType = "captcha"
	WaapRuleActionTypeHandshake WaapRuleActionType = "handshake"
	WaapRuleActionTypeMonitor   WaapRuleActionType = "monitor"
	WaapRuleActionTypeTag       WaapRuleActionType = "tag"
)

type WaapRuleBlockedRequests struct {
	// The action taken by the rule
	Action string `json:"action,required"`
	// The number of requests blocked by the rule
	Count int64 `json:"count,required"`
	// The name of the rule that blocked the request
	RuleName string `json:"rule_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Count       respjson.Field
		RuleName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRuleBlockedRequests) RawJSON() string { return r.JSON.raw }
func (r *WaapRuleBlockedRequests) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a custom rule set.
type WaapRuleSet struct {
	// Identifier of the rule set.
	ID int64 `json:"id,required"`
	// Detailed description of the rule set.
	Description string `json:"description,required"`
	// Indicates if the rule set is currently active.
	IsActive bool `json:"is_active,required"`
	// Name of the rule set.
	Name string `json:"name,required"`
	// Collection of tags associated with the rule set.
	Tags []WaapRuleSetTag `json:"tags,required"`
	// The resource slug associated with the rule set.
	ResourceSlug string             `json:"resource_slug,nullable"`
	Rules        []WaapDomainPolicy `json:"rules"`
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
	ID int64 `json:"id,required"`
	// Detailed description of the tag.
	Description string `json:"description,required"`
	// Name of the tag.
	Name string `json:"name,required"`
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

// Response model for the statistics item
type WaapStatisticItem struct {
	// The date and time for the statistic in ISO 8601 format
	DateTime time.Time `json:"date_time,required" format:"date-time"`
	// The value for the statistic. If there is no data for the given time, the value
	// will be 0.
	Value int64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateTime    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapStatisticItem) RawJSON() string { return r.JSON.raw }
func (r *WaapStatisticItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response model for the statistics series
type WaapStatisticsSeries struct {
	// Will be returned if `total_bytes` is requested in the metrics parameter
	TotalBytes []WaapStatisticItem `json:"total_bytes,nullable"`
	// Will be included if `total_requests` is requested in the metrics parameter
	TotalRequests []WaapStatisticItem `json:"total_requests,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TotalBytes    respjson.Field
		TotalRequests respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapStatisticsSeries) RawJSON() string { return r.JSON.raw }
func (r *WaapStatisticsSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a WAAP domain when getting a list of domains.
type WaapSummaryDomain struct {
	// The domain ID
	ID int64 `json:"id,required"`
	// The date and time the domain was created in ISO 8601 format
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The ID of the custom page set
	CustomPageSet int64 `json:"custom_page_set,required"`
	// The domain name
	Name string `json:"name,required"`
	// The different statuses a domain can have
	//
	// Any of "active", "bypass", "monitor", "locked".
	Status WaapDomainStatus `json:"status,required"`
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

// Tags provide shortcuts for the rules used in WAAP policies for the creation of
// more complex WAAP rules.
type WaapTag struct {
	// A tag's human readable description
	Description string `json:"description,required"`
	// The name of a tag that should be used in a WAAP rule condition
	Name string `json:"name,required"`
	// The display name of the tag
	ReadableName string `json:"readable_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description  respjson.Field
		Name         respjson.Field
		ReadableName respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTag) RawJSON() string { return r.JSON.raw }
func (r *WaapTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapTimeSeriesAttack struct {
	// The type of attack
	AttackType string `json:"attack_type,required"`
	// The time series data
	Values []WaapTimeSeriesAttackValue `json:"values,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttackType  respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTimeSeriesAttack) RawJSON() string { return r.JSON.raw }
func (r *WaapTimeSeriesAttack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapTimeSeriesAttackValue struct {
	// The number of attacks
	Count int64 `json:"count,required"`
	// The timestamp of the time series item as a POSIX timestamp
	Timestamp int64 `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTimeSeriesAttackValue) RawJSON() string { return r.JSON.raw }
func (r *WaapTimeSeriesAttackValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapTopSession struct {
	// The number of blocked requests in the session
	Blocked int64 `json:"blocked,required"`
	// The duration of the session in seconds
	Duration float64 `json:"duration,required"`
	// The number of requests in the session
	Requests int64 `json:"requests,required"`
	// The session ID
	SessionID string `json:"session_id,required" format:"uuid"`
	// The start time of the session as a POSIX timestamp
	StartTime time.Time `json:"start_time,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blocked     respjson.Field
		Duration    respjson.Field
		Requests    respjson.Field
		SessionID   respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTopSession) RawJSON() string { return r.JSON.raw }
func (r *WaapTopSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapTopURL struct {
	// The number of attacks to the URL
	Count int64 `json:"count,required"`
	// The URL that was attacked
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTopURL) RawJSON() string { return r.JSON.raw }
func (r *WaapTopURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapTopUserAgent struct {
	// The number of requests made with the user agent
	Count int64 `json:"count,required"`
	// The user agent that was used
	UserAgent string `json:"user_agent,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		UserAgent   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapTopUserAgent) RawJSON() string { return r.JSON.raw }
func (r *WaapTopUserAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents the traffic metrics for a domain at a given time window
type WaapTrafficMetrics struct {
	// UNIX timestamp indicating when the traffic data was recorded
	Timestamp int64 `json:"timestamp,required"`
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

type WaapTrafficType string

const (
	WaapTrafficTypePolicyAllowed     WaapTrafficType = "policy_allowed"
	WaapTrafficTypePolicyBlocked     WaapTrafficType = "policy_blocked"
	WaapTrafficTypeCustomRuleAllowed WaapTrafficType = "custom_rule_allowed"
	WaapTrafficTypeCustomBlocked     WaapTrafficType = "custom_blocked"
	WaapTrafficTypeLegitRequests     WaapTrafficType = "legit_requests"
	WaapTrafficTypeSanctioned        WaapTrafficType = "sanctioned"
	WaapTrafficTypeDynamic           WaapTrafficType = "dynamic"
	WaapTrafficTypeAPI               WaapTrafficType = "api"
	WaapTrafficTypeStatic            WaapTrafficType = "static"
	WaapTrafficTypeAjax              WaapTrafficType = "ajax"
	WaapTrafficTypeRedirects         WaapTrafficType = "redirects"
	WaapTrafficTypeMonitor           WaapTrafficType = "monitor"
	WaapTrafficTypeErr40x            WaapTrafficType = "err_40x"
	WaapTrafficTypeErr50x            WaapTrafficType = "err_50x"
	WaapTrafficTypePassedToOrigin    WaapTrafficType = "passed_to_origin"
	WaapTrafficTypeTimeout           WaapTrafficType = "timeout"
	WaapTrafficTypeOther             WaapTrafficType = "other"
	WaapTrafficTypeDDOS              WaapTrafficType = "ddos"
	WaapTrafficTypeLegit             WaapTrafficType = "legit"
	WaapTrafficTypeMonitored         WaapTrafficType = "monitored"
)

// User agent details
type WaapUserAgentDetails struct {
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
	RenderingEngine string `json:"rendering_engine,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseBrowser        respjson.Field
		BaseBrowserVersion respjson.Field
		Client             respjson.Field
		ClientType         respjson.Field
		ClientVersion      respjson.Field
		CPU                respjson.Field
		Device             respjson.Field
		DeviceType         respjson.Field
		FullString         respjson.Field
		Os                 respjson.Field
		RenderingEngine    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapUserAgentDetails) RawJSON() string { return r.JSON.raw }
func (r *WaapUserAgentDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents the WAAP service information for a client
type WaapGetAccountOverviewResponse struct {
	// The client ID
	ID int64 `json:"id,required"`
	// List of enabled features
	Features []string `json:"features,required"`
	// Quotas for the client
	Quotas map[string]WaapGetAccountOverviewResponseQuota `json:"quotas,required"`
	// Information about the WAAP service status
	Service WaapGetAccountOverviewResponseService `json:"service,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Features    respjson.Field
		Quotas      respjson.Field
		Service     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapGetAccountOverviewResponse) RawJSON() string { return r.JSON.raw }
func (r *WaapGetAccountOverviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapGetAccountOverviewResponseQuota struct {
	// The maximum allowed number of this resource
	Allowed int64 `json:"allowed,required"`
	// The current number of this resource
	Current int64 `json:"current,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allowed     respjson.Field
		Current     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapGetAccountOverviewResponseQuota) RawJSON() string { return r.JSON.raw }
func (r *WaapGetAccountOverviewResponseQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the WAAP service status
type WaapGetAccountOverviewResponseService struct {
	// Whether the service is enabled
	Enabled bool `json:"enabled,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapGetAccountOverviewResponseService) RawJSON() string { return r.JSON.raw }
func (r *WaapGetAccountOverviewResponseService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
