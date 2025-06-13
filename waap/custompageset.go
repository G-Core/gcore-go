// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// CustomPageSetService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomPageSetService] method instead.
type CustomPageSetService struct {
	Options []option.RequestOption
}

// NewCustomPageSetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomPageSetService(opts ...option.RequestOption) (r CustomPageSetService) {
	r = CustomPageSetService{}
	r.Options = opts
	return
}

// Create a custom page set based on the provided data. For any custom page type
// (block, `block_csrf`, etc) that is not provided the default page will be used.
func (r *CustomPageSetService) New(ctx context.Context, body CustomPageSetNewParams, opts ...option.RequestOption) (res *CustomPageSet, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/custom-page-sets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a custom page set based on the provided parameters. To update a field,
// provide the field with the new value. To remove a field, provide it as null. To
// keep a field unaltered, do not include it in the request. Note: `name` cannot be
// removed. When updating a custom page, include all the fields that you want it to
// have. Any field not included will be removed.
func (r *CustomPageSetService) Update(ctx context.Context, setID int64, body CustomPageSetUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/custom-page-sets/%v", setID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Retrieve a list of custom page sets available for use
func (r *CustomPageSetService) List(ctx context.Context, query CustomPageSetListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[CustomPageSet], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "waap/v1/custom-page-sets"
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

// Retrieve a list of custom page sets available for use
func (r *CustomPageSetService) ListAutoPaging(ctx context.Context, query CustomPageSetListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[CustomPageSet] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a custom page set
func (r *CustomPageSetService) Delete(ctx context.Context, setID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/custom-page-sets/%v", setID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve a custom page set based on the provided ID
func (r *CustomPageSetService) Get(ctx context.Context, setID int64, opts ...option.RequestOption) (res *CustomPageSet, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/custom-page-sets/%v", setID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Allows to preview a custom page without creating it based on the provided type
// and data
func (r *CustomPageSetService) Preview(ctx context.Context, params CustomPageSetPreviewParams, opts ...option.RequestOption) (res *PreviewCustomPage, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/preview-custom-page"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type BlockCsrfPageData struct {
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
func (r BlockCsrfPageData) RawJSON() string { return r.JSON.raw }
func (r *BlockCsrfPageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BlockCsrfPageData to a BlockCsrfPageDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BlockCsrfPageDataParam.Overrides()
func (r BlockCsrfPageData) ToParam() BlockCsrfPageDataParam {
	return param.Override[BlockCsrfPageDataParam](json.RawMessage(r.RawJSON()))
}

// The property Enabled is required.
type BlockCsrfPageDataParam struct {
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

func (r BlockCsrfPageDataParam) MarshalJSON() (data []byte, err error) {
	type shadow BlockCsrfPageDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlockCsrfPageDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlockPageData struct {
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
func (r BlockPageData) RawJSON() string { return r.JSON.raw }
func (r *BlockPageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BlockPageData to a BlockPageDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BlockPageDataParam.Overrides()
func (r BlockPageData) ToParam() BlockPageDataParam {
	return param.Override[BlockPageDataParam](json.RawMessage(r.RawJSON()))
}

// The property Enabled is required.
type BlockPageDataParam struct {
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

func (r BlockPageDataParam) MarshalJSON() (data []byte, err error) {
	type shadow BlockPageDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlockPageDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CaptchaPageData struct {
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
func (r CaptchaPageData) RawJSON() string { return r.JSON.raw }
func (r *CaptchaPageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CaptchaPageData to a CaptchaPageDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CaptchaPageDataParam.Overrides()
func (r CaptchaPageData) ToParam() CaptchaPageDataParam {
	return param.Override[CaptchaPageDataParam](json.RawMessage(r.RawJSON()))
}

// The property Enabled is required.
type CaptchaPageDataParam struct {
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

func (r CaptchaPageDataParam) MarshalJSON() (data []byte, err error) {
	type shadow CaptchaPageDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CaptchaPageDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CookieDisabledPageData struct {
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
func (r CookieDisabledPageData) RawJSON() string { return r.JSON.raw }
func (r *CookieDisabledPageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CookieDisabledPageData to a CookieDisabledPageDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CookieDisabledPageDataParam.Overrides()
func (r CookieDisabledPageData) ToParam() CookieDisabledPageDataParam {
	return param.Override[CookieDisabledPageDataParam](json.RawMessage(r.RawJSON()))
}

// The property Enabled is required.
type CookieDisabledPageDataParam struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header param.Opt[string] `json:"header,omitzero"`
	// The text to display in the body of the custom page
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r CookieDisabledPageDataParam) MarshalJSON() (data []byte, err error) {
	type shadow CookieDisabledPageDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CookieDisabledPageDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomPageSet struct {
	// The ID of the custom page set
	ID int64 `json:"id,required"`
	// Name of the custom page set
	Name           string                 `json:"name,required"`
	Block          BlockPageData          `json:"block,nullable"`
	BlockCsrf      BlockCsrfPageData      `json:"block_csrf,nullable"`
	Captcha        CaptchaPageData        `json:"captcha,nullable"`
	CookieDisabled CookieDisabledPageData `json:"cookie_disabled,nullable"`
	// List of domain IDs that are associated with this page set
	Domains            []int64                    `json:"domains,nullable"`
	Handshake          HandshakePageData          `json:"handshake,nullable"`
	JavascriptDisabled JavascriptDisabledPageData `json:"javascript_disabled,nullable"`
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
func (r CustomPageSet) RawJSON() string { return r.JSON.raw }
func (r *CustomPageSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HandshakePageData struct {
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
func (r HandshakePageData) RawJSON() string { return r.JSON.raw }
func (r *HandshakePageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this HandshakePageData to a HandshakePageDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// HandshakePageDataParam.Overrides()
func (r HandshakePageData) ToParam() HandshakePageDataParam {
	return param.Override[HandshakePageDataParam](json.RawMessage(r.RawJSON()))
}

// The property Enabled is required.
type HandshakePageDataParam struct {
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

func (r HandshakePageDataParam) MarshalJSON() (data []byte, err error) {
	type shadow HandshakePageDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HandshakePageDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JavascriptDisabledPageData struct {
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
func (r JavascriptDisabledPageData) RawJSON() string { return r.JSON.raw }
func (r *JavascriptDisabledPageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JavascriptDisabledPageData to a
// JavascriptDisabledPageDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JavascriptDisabledPageDataParam.Overrides()
func (r JavascriptDisabledPageData) ToParam() JavascriptDisabledPageDataParam {
	return param.Override[JavascriptDisabledPageDataParam](json.RawMessage(r.RawJSON()))
}

// The property Enabled is required.
type JavascriptDisabledPageDataParam struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header param.Opt[string] `json:"header,omitzero"`
	// The text to display in the body of the custom page
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r JavascriptDisabledPageDataParam) MarshalJSON() (data []byte, err error) {
	type shadow JavascriptDisabledPageDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JavascriptDisabledPageDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreviewCustomPage struct {
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
func (r PreviewCustomPage) RawJSON() string { return r.JSON.raw }
func (r *PreviewCustomPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomPageSetNewParams struct {
	// Name of the custom page set
	Name string `json:"name,required"`
	// List of domain IDs that are associated with this page set
	Domains            []int64                         `json:"domains,omitzero"`
	Block              BlockPageDataParam              `json:"block,omitzero"`
	BlockCsrf          BlockCsrfPageDataParam          `json:"block_csrf,omitzero"`
	Captcha            CaptchaPageDataParam            `json:"captcha,omitzero"`
	CookieDisabled     CookieDisabledPageDataParam     `json:"cookie_disabled,omitzero"`
	Handshake          HandshakePageDataParam          `json:"handshake,omitzero"`
	JavascriptDisabled JavascriptDisabledPageDataParam `json:"javascript_disabled,omitzero"`
	paramObj
}

func (r CustomPageSetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomPageSetUpdateParams struct {
	// Name of the custom page set
	Name param.Opt[string] `json:"name,omitzero"`
	// List of domain IDs that are associated with this page set
	Domains            []int64                         `json:"domains,omitzero"`
	Block              BlockPageDataParam              `json:"block,omitzero"`
	BlockCsrf          BlockCsrfPageDataParam          `json:"block_csrf,omitzero"`
	Captcha            CaptchaPageDataParam            `json:"captcha,omitzero"`
	CookieDisabled     CookieDisabledPageDataParam     `json:"cookie_disabled,omitzero"`
	Handshake          HandshakePageDataParam          `json:"handshake,omitzero"`
	JavascriptDisabled JavascriptDisabledPageDataParam `json:"javascript_disabled,omitzero"`
	paramObj
}

func (r CustomPageSetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomPageSetListParams struct {
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter page sets based on their name. Supports '\*' as a wildcard character
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter page sets based on their IDs
	IDs []int64 `query:"ids,omitzero" json:"-"`
	// Sort the response by given field.
	//
	// Any of "name", "-name", "id", "-id".
	Ordering CustomPageSetListParamsOrdering `query:"ordering,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomPageSetListParams]'s query parameters as
// `url.Values`.
func (r CustomPageSetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort the response by given field.
type CustomPageSetListParamsOrdering string

const (
	CustomPageSetListParamsOrderingName      CustomPageSetListParamsOrdering = "name"
	CustomPageSetListParamsOrderingMinusName CustomPageSetListParamsOrdering = "-name"
	CustomPageSetListParamsOrderingID        CustomPageSetListParamsOrdering = "id"
	CustomPageSetListParamsOrderingMinusID   CustomPageSetListParamsOrdering = "-id"
)

type CustomPageSetPreviewParams struct {
	// The type of the custom page
	//
	// Any of "block.html", "block_csrf.html", "captcha.html", "cookieDisabled.html",
	// "handshake.html", "javascriptDisabled.html".
	PageType CustomPageSetPreviewParamsPageType `query:"page_type,omitzero,required" json:"-"`
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

func (r CustomPageSetPreviewParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetPreviewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetPreviewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CustomPageSetPreviewParams]'s query parameters as
// `url.Values`.
func (r CustomPageSetPreviewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The type of the custom page
type CustomPageSetPreviewParamsPageType string

const (
	CustomPageSetPreviewParamsPageTypeBlockHTML              CustomPageSetPreviewParamsPageType = "block.html"
	CustomPageSetPreviewParamsPageTypeBlockCsrfHTML          CustomPageSetPreviewParamsPageType = "block_csrf.html"
	CustomPageSetPreviewParamsPageTypeCaptchaHTML            CustomPageSetPreviewParamsPageType = "captcha.html"
	CustomPageSetPreviewParamsPageTypeCookieDisabledHTML     CustomPageSetPreviewParamsPageType = "cookieDisabled.html"
	CustomPageSetPreviewParamsPageTypeHandshakeHTML          CustomPageSetPreviewParamsPageType = "handshake.html"
	CustomPageSetPreviewParamsPageTypeJavascriptDisabledHTML CustomPageSetPreviewParamsPageType = "javascriptDisabled.html"
)
