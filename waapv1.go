// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// WaapV1Service contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1Service] method instead.
type WaapV1Service struct {
	Options          []option.RequestOption
	Clients          *WaapV1ClientService
	Statistics       *WaapV1StatisticService
	Domains          *WaapV1DomainService
	CustomPageSets   *WaapV1CustomPageSetService
	AdvancedRules    *WaapV1AdvancedRuleService
	IPInfo           *WaapV1IPInfoService
	SecurityInsights *WaapV1SecurityInsightService
}

// NewWaapV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWaapV1Service(opts ...option.RequestOption) (r *WaapV1Service) {
	r = &WaapV1Service{}
	r.Options = opts
	r.Clients = NewWaapV1ClientService(opts...)
	r.Statistics = NewWaapV1StatisticService(opts...)
	r.Domains = NewWaapV1DomainService(opts...)
	r.CustomPageSets = NewWaapV1CustomPageSetService(opts...)
	r.AdvancedRules = NewWaapV1AdvancedRuleService(opts...)
	r.IPInfo = NewWaapV1IPInfoService(opts...)
	r.SecurityInsights = NewWaapV1SecurityInsightService(opts...)
	return
}

// This endpoint retrieves a list of network organizations that own IP ranges as
// identified by the Whois service.It supports pagination, filtering based on
// various parameters, and ordering of results.
func (r *WaapV1Service) ListOrganizations(ctx context.Context, query WaapV1ListOrganizationsParams, opts ...option.RequestOption) (res *WaapV1ListOrganizationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Allows to preview a custom page without creating it based on the provided type
// and data
func (r *WaapV1Service) PreviewCustomPage(ctx context.Context, params WaapV1PreviewCustomPageParams, opts ...option.RequestOption) (res *WaapV1PreviewCustomPageResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/preview-custom-page"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Tags are shortcuts for the rules used in WAAP policies for the creation of more
// complex WAAP rules
func (r *WaapV1Service) GetTags(ctx context.Context, query WaapV1GetTagsParams, opts ...option.RequestOption) (res *WaapV1GetTagsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WaapV1ListOrganizationsResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []WaapV1ListOrganizationsResponseResult `json:"results,required"`
	JSON    waapV1ListOrganizationsResponseJSON     `json:"-"`
}

// waapV1ListOrganizationsResponseJSON contains the JSON metadata for the struct
// [WaapV1ListOrganizationsResponse]
type waapV1ListOrganizationsResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1ListOrganizationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1ListOrganizationsResponseJSON) RawJSON() string {
	return r.raw
}

// Represents an IP range owner organization
type WaapV1ListOrganizationsResponseResult struct {
	// The ID of an organization
	ID int64 `json:"id,required"`
	// The name of an organization
	Name string                                    `json:"name,required"`
	JSON waapV1ListOrganizationsResponseResultJSON `json:"-"`
}

// waapV1ListOrganizationsResponseResultJSON contains the JSON metadata for the
// struct [WaapV1ListOrganizationsResponseResult]
type waapV1ListOrganizationsResponseResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1ListOrganizationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1ListOrganizationsResponseResultJSON) RawJSON() string {
	return r.raw
}

type WaapV1PreviewCustomPageResponse struct {
	// HTML content of the custom page
	HTML string                              `json:"html,required"`
	JSON waapV1PreviewCustomPageResponseJSON `json:"-"`
}

// waapV1PreviewCustomPageResponseJSON contains the JSON metadata for the struct
// [WaapV1PreviewCustomPageResponse]
type waapV1PreviewCustomPageResponseJSON struct {
	HTML        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1PreviewCustomPageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1PreviewCustomPageResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1GetTagsResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []WaapV1GetTagsResponseResult `json:"results,required"`
	JSON    waapV1GetTagsResponseJSON     `json:"-"`
}

// waapV1GetTagsResponseJSON contains the JSON metadata for the struct
// [WaapV1GetTagsResponse]
type waapV1GetTagsResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1GetTagsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1GetTagsResponseJSON) RawJSON() string {
	return r.raw
}

// Tags provide shortcuts for the rules used in WAAP policies for the creation of
// more complex WAAP rules.
type WaapV1GetTagsResponseResult struct {
	// A tag's human readable description
	Description string `json:"description,required"`
	// The name of a tag that should be used in a WAAP rule condition
	Name string `json:"name,required"`
	// The display name of the tag
	ReadableName string                          `json:"readable_name,required"`
	JSON         waapV1GetTagsResponseResultJSON `json:"-"`
}

// waapV1GetTagsResponseResultJSON contains the JSON metadata for the struct
// [WaapV1GetTagsResponseResult]
type waapV1GetTagsResponseResultJSON struct {
	Description  apijson.Field
	Name         apijson.Field
	ReadableName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WaapV1GetTagsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1GetTagsResponseResultJSON) RawJSON() string {
	return r.raw
}

type WaapV1ListOrganizationsParams struct {
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Filter organizations by their name. Supports '\*' as a wildcard character.
	Name param.Field[string] `query:"name"`
	// Number of items to skip
	Offset param.Field[int64] `query:"offset"`
	// Determine the field to order results by
	Ordering param.Field[WaapV1ListOrganizationsParamsOrdering] `query:"ordering"`
}

// URLQuery serializes [WaapV1ListOrganizationsParams]'s query parameters as
// `url.Values`.
func (r WaapV1ListOrganizationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Determine the field to order results by
type WaapV1ListOrganizationsParamsOrdering string

const (
	WaapV1ListOrganizationsParamsOrderingName      WaapV1ListOrganizationsParamsOrdering = "name"
	WaapV1ListOrganizationsParamsOrderingID        WaapV1ListOrganizationsParamsOrdering = "id"
	WaapV1ListOrganizationsParamsOrderingMinusName WaapV1ListOrganizationsParamsOrdering = "-name"
	WaapV1ListOrganizationsParamsOrderingMinusID   WaapV1ListOrganizationsParamsOrdering = "-id"
)

func (r WaapV1ListOrganizationsParamsOrdering) IsKnown() bool {
	switch r {
	case WaapV1ListOrganizationsParamsOrderingName, WaapV1ListOrganizationsParamsOrderingID, WaapV1ListOrganizationsParamsOrderingMinusName, WaapV1ListOrganizationsParamsOrderingMinusID:
		return true
	}
	return false
}

type WaapV1PreviewCustomPageParams struct {
	// The type of the custom page
	PageType param.Field[WaapV1PreviewCustomPageParamsPageType] `query:"page_type,required"`
	// Error message
	Error param.Field[string] `json:"error"`
	// The text to display in the header of the custom page
	Header param.Field[string] `json:"header"`
	// Supported image types are JPEG, PNG and JPG, size is limited to width 450px,
	// height 130px. This should be a base 64 encoding of the full HTML img tag
	// compatible image, with the header included.
	Logo param.Field[string] `json:"logo"`
	// The text to display in the body of the custom page
	Text param.Field[string] `json:"text"`
	// The text to display in the title of the custom page
	Title param.Field[string] `json:"title"`
}

func (r WaapV1PreviewCustomPageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [WaapV1PreviewCustomPageParams]'s query parameters as
// `url.Values`.
func (r WaapV1PreviewCustomPageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of the custom page
type WaapV1PreviewCustomPageParamsPageType string

const (
	WaapV1PreviewCustomPageParamsPageTypeBlockHTML              WaapV1PreviewCustomPageParamsPageType = "block.html"
	WaapV1PreviewCustomPageParamsPageTypeBlockCsrfHTML          WaapV1PreviewCustomPageParamsPageType = "block_csrf.html"
	WaapV1PreviewCustomPageParamsPageTypeCaptchaHTML            WaapV1PreviewCustomPageParamsPageType = "captcha.html"
	WaapV1PreviewCustomPageParamsPageTypeCookieDisabledHTML     WaapV1PreviewCustomPageParamsPageType = "cookieDisabled.html"
	WaapV1PreviewCustomPageParamsPageTypeHandshakeHTML          WaapV1PreviewCustomPageParamsPageType = "handshake.html"
	WaapV1PreviewCustomPageParamsPageTypeJavascriptDisabledHTML WaapV1PreviewCustomPageParamsPageType = "javascriptDisabled.html"
)

func (r WaapV1PreviewCustomPageParamsPageType) IsKnown() bool {
	switch r {
	case WaapV1PreviewCustomPageParamsPageTypeBlockHTML, WaapV1PreviewCustomPageParamsPageTypeBlockCsrfHTML, WaapV1PreviewCustomPageParamsPageTypeCaptchaHTML, WaapV1PreviewCustomPageParamsPageTypeCookieDisabledHTML, WaapV1PreviewCustomPageParamsPageTypeHandshakeHTML, WaapV1PreviewCustomPageParamsPageTypeJavascriptDisabledHTML:
		return true
	}
	return false
}

type WaapV1GetTagsParams struct {
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Filter tags by their name. Supports '\*' as a wildcard character.
	Name param.Field[string] `query:"name"`
	// Number of items to skip
	Offset param.Field[int64] `query:"offset"`
	// Determine the field to order results by
	Ordering param.Field[WaapV1GetTagsParamsOrdering] `query:"ordering"`
	// Filter to include only reserved tags.
	Reserved param.Field[bool] `query:"reserved"`
}

// URLQuery serializes [WaapV1GetTagsParams]'s query parameters as `url.Values`.
func (r WaapV1GetTagsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Determine the field to order results by
type WaapV1GetTagsParamsOrdering string

const (
	WaapV1GetTagsParamsOrderingName              WaapV1GetTagsParamsOrdering = "name"
	WaapV1GetTagsParamsOrderingReadableName      WaapV1GetTagsParamsOrdering = "readable_name"
	WaapV1GetTagsParamsOrderingReserved          WaapV1GetTagsParamsOrdering = "reserved"
	WaapV1GetTagsParamsOrderingMinusName         WaapV1GetTagsParamsOrdering = "-name"
	WaapV1GetTagsParamsOrderingMinusReadableName WaapV1GetTagsParamsOrdering = "-readable_name"
	WaapV1GetTagsParamsOrderingMinusReserved     WaapV1GetTagsParamsOrdering = "-reserved"
)

func (r WaapV1GetTagsParamsOrdering) IsKnown() bool {
	switch r {
	case WaapV1GetTagsParamsOrderingName, WaapV1GetTagsParamsOrderingReadableName, WaapV1GetTagsParamsOrderingReserved, WaapV1GetTagsParamsOrderingMinusName, WaapV1GetTagsParamsOrderingMinusReadableName, WaapV1GetTagsParamsOrderingMinusReserved:
		return true
	}
	return false
}
