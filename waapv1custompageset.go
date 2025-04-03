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

// WaapV1CustomPageSetService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1CustomPageSetService] method instead.
type WaapV1CustomPageSetService struct {
	Options []option.RequestOption
}

// NewWaapV1CustomPageSetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWaapV1CustomPageSetService(opts ...option.RequestOption) (r *WaapV1CustomPageSetService) {
	r = &WaapV1CustomPageSetService{}
	r.Options = opts
	return
}

// Create a custom page set based on the provided data. For any custom page type
// (block, block_csrf, etc) that is not provided the default page will be used.
func (r *WaapV1CustomPageSetService) New(ctx context.Context, body WaapV1CustomPageSetNewParams, opts ...option.RequestOption) (res *CustomPageSet, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/custom-page-sets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a custom page set based on the provided ID
func (r *WaapV1CustomPageSetService) Get(ctx context.Context, setID int64, opts ...option.RequestOption) (res *CustomPageSet, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/custom-page-sets/%v", setID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a custom page set based on the provided parameters. To update a field,
// provide the field with the new value. To remove a field, provide it as null. To
// keep a field unaltered, do not include it in the request. Note: `name` cannot be
// removed. When updating a custom page, include all the fields that you want it to
// have. Any field not included will be removed.
func (r *WaapV1CustomPageSetService) Update(ctx context.Context, setID int64, body WaapV1CustomPageSetUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/custom-page-sets/%v", setID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Retrieve a list of custom page sets available for use
func (r *WaapV1CustomPageSetService) List(ctx context.Context, query WaapV1CustomPageSetListParams, opts ...option.RequestOption) (res *WaapV1CustomPageSetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/custom-page-sets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a custom page set
func (r *WaapV1CustomPageSetService) Delete(ctx context.Context, setID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/custom-page-sets/%v", setID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type BlockCsrfPage struct {
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
	Title string            `json:"title"`
	JSON  blockCsrfPageJSON `json:"-"`
}

// blockCsrfPageJSON contains the JSON metadata for the struct [BlockCsrfPage]
type blockCsrfPageJSON struct {
	Enabled     apijson.Field
	Header      apijson.Field
	Logo        apijson.Field
	Text        apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockCsrfPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockCsrfPageJSON) RawJSON() string {
	return r.raw
}

type BlockCsrfPageParam struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled param.Field[bool] `json:"enabled,required"`
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

func (r BlockCsrfPageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BlockPage struct {
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
	Title string        `json:"title"`
	JSON  blockPageJSON `json:"-"`
}

// blockPageJSON contains the JSON metadata for the struct [BlockPage]
type blockPageJSON struct {
	Enabled     apijson.Field
	Header      apijson.Field
	Logo        apijson.Field
	Text        apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockPageJSON) RawJSON() string {
	return r.raw
}

type BlockPageParam struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled param.Field[bool] `json:"enabled,required"`
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

func (r BlockPageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CaptchaPage struct {
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
	Title string          `json:"title"`
	JSON  captchaPageJSON `json:"-"`
}

// captchaPageJSON contains the JSON metadata for the struct [CaptchaPage]
type captchaPageJSON struct {
	Enabled     apijson.Field
	Error       apijson.Field
	Header      apijson.Field
	Logo        apijson.Field
	Text        apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptchaPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captchaPageJSON) RawJSON() string {
	return r.raw
}

type CaptchaPageParam struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled param.Field[bool] `json:"enabled,required"`
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

func (r CaptchaPageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CookieDisabledPage struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header string `json:"header"`
	// The text to display in the body of the custom page
	Text string                 `json:"text"`
	JSON cookieDisabledPageJSON `json:"-"`
}

// cookieDisabledPageJSON contains the JSON metadata for the struct
// [CookieDisabledPage]
type cookieDisabledPageJSON struct {
	Enabled     apijson.Field
	Header      apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CookieDisabledPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cookieDisabledPageJSON) RawJSON() string {
	return r.raw
}

type CookieDisabledPageParam struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled param.Field[bool] `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header param.Field[string] `json:"header"`
	// The text to display in the body of the custom page
	Text param.Field[string] `json:"text"`
}

func (r CookieDisabledPageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomPageSet struct {
	// The ID of the custom page set
	ID int64 `json:"id,required"`
	// Name of the custom page set
	Name           string             `json:"name,required"`
	Block          BlockPage          `json:"block,nullable"`
	BlockCsrf      BlockCsrfPage      `json:"block_csrf,nullable"`
	Captcha        CaptchaPage        `json:"captcha,nullable"`
	CookieDisabled CookieDisabledPage `json:"cookie_disabled,nullable"`
	// List of domain IDs that are associated with this page set
	Domains            []int64                `json:"domains,nullable"`
	Handshake          HandshakePage          `json:"handshake,nullable"`
	JavascriptDisabled JavascriptDisabledPage `json:"javascript_disabled,nullable"`
	JSON               customPageSetJSON      `json:"-"`
}

// customPageSetJSON contains the JSON metadata for the struct [CustomPageSet]
type customPageSetJSON struct {
	ID                 apijson.Field
	Name               apijson.Field
	Block              apijson.Field
	BlockCsrf          apijson.Field
	Captcha            apijson.Field
	CookieDisabled     apijson.Field
	Domains            apijson.Field
	Handshake          apijson.Field
	JavascriptDisabled apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomPageSet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageSetJSON) RawJSON() string {
	return r.raw
}

type HandshakePage struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header string `json:"header"`
	// Supported image types are JPEG, PNG and JPG, size is limited to width 450px,
	// height 130px. This should be a base 64 encoding of the full HTML img tag
	// compatible image, with the header included.
	Logo string `json:"logo"`
	// The text to display in the title of the custom page
	Title string            `json:"title"`
	JSON  handshakePageJSON `json:"-"`
}

// handshakePageJSON contains the JSON metadata for the struct [HandshakePage]
type handshakePageJSON struct {
	Enabled     apijson.Field
	Header      apijson.Field
	Logo        apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HandshakePage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r handshakePageJSON) RawJSON() string {
	return r.raw
}

type HandshakePageParam struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled param.Field[bool] `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header param.Field[string] `json:"header"`
	// Supported image types are JPEG, PNG and JPG, size is limited to width 450px,
	// height 130px. This should be a base 64 encoding of the full HTML img tag
	// compatible image, with the header included.
	Logo param.Field[string] `json:"logo"`
	// The text to display in the title of the custom page
	Title param.Field[string] `json:"title"`
}

func (r HandshakePageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type JavascriptDisabledPage struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header string `json:"header"`
	// The text to display in the body of the custom page
	Text string                     `json:"text"`
	JSON javascriptDisabledPageJSON `json:"-"`
}

// javascriptDisabledPageJSON contains the JSON metadata for the struct
// [JavascriptDisabledPage]
type javascriptDisabledPageJSON struct {
	Enabled     apijson.Field
	Header      apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JavascriptDisabledPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r javascriptDisabledPageJSON) RawJSON() string {
	return r.raw
}

type JavascriptDisabledPageParam struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled param.Field[bool] `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header param.Field[string] `json:"header"`
	// The text to display in the body of the custom page
	Text param.Field[string] `json:"text"`
}

func (r JavascriptDisabledPageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaapV1CustomPageSetListResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []CustomPageSet                     `json:"results,required"`
	JSON    waapV1CustomPageSetListResponseJSON `json:"-"`
}

// waapV1CustomPageSetListResponseJSON contains the JSON metadata for the struct
// [WaapV1CustomPageSetListResponse]
type waapV1CustomPageSetListResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1CustomPageSetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1CustomPageSetListResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1CustomPageSetNewParams struct {
	// Name of the custom page set
	Name           param.Field[string]                  `json:"name,required"`
	Block          param.Field[BlockPageParam]          `json:"block"`
	BlockCsrf      param.Field[BlockCsrfPageParam]      `json:"block_csrf"`
	Captcha        param.Field[CaptchaPageParam]        `json:"captcha"`
	CookieDisabled param.Field[CookieDisabledPageParam] `json:"cookie_disabled"`
	// List of domain IDs that are associated with this page set
	Domains            param.Field[[]int64]                     `json:"domains"`
	Handshake          param.Field[HandshakePageParam]          `json:"handshake"`
	JavascriptDisabled param.Field[JavascriptDisabledPageParam] `json:"javascript_disabled"`
}

func (r WaapV1CustomPageSetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaapV1CustomPageSetUpdateParams struct {
	Block          param.Field[BlockPageParam]          `json:"block"`
	BlockCsrf      param.Field[BlockCsrfPageParam]      `json:"block_csrf"`
	Captcha        param.Field[CaptchaPageParam]        `json:"captcha"`
	CookieDisabled param.Field[CookieDisabledPageParam] `json:"cookie_disabled"`
	// List of domain IDs that are associated with this page set
	Domains            param.Field[[]int64]                     `json:"domains"`
	Handshake          param.Field[HandshakePageParam]          `json:"handshake"`
	JavascriptDisabled param.Field[JavascriptDisabledPageParam] `json:"javascript_disabled"`
	// Name of the custom page set
	Name param.Field[string] `json:"name"`
}

func (r WaapV1CustomPageSetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WaapV1CustomPageSetListParams struct {
	// Filter page sets based on their IDs
	IDs param.Field[[]int64] `query:"ids"`
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Filter page sets based on their name. Supports '\*' as a wildcard character
	Name param.Field[string] `query:"name"`
	// Number of items to skip
	Offset param.Field[int64] `query:"offset"`
	// Sort the response by given field.
	Ordering param.Field[WaapV1CustomPageSetListParamsOrdering] `query:"ordering"`
}

// URLQuery serializes [WaapV1CustomPageSetListParams]'s query parameters as
// `url.Values`.
func (r WaapV1CustomPageSetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort the response by given field.
type WaapV1CustomPageSetListParamsOrdering string

const (
	WaapV1CustomPageSetListParamsOrderingName      WaapV1CustomPageSetListParamsOrdering = "name"
	WaapV1CustomPageSetListParamsOrderingMinusName WaapV1CustomPageSetListParamsOrdering = "-name"
	WaapV1CustomPageSetListParamsOrderingID        WaapV1CustomPageSetListParamsOrdering = "id"
	WaapV1CustomPageSetListParamsOrderingMinusID   WaapV1CustomPageSetListParamsOrdering = "-id"
)

func (r WaapV1CustomPageSetListParamsOrdering) IsKnown() bool {
	switch r {
	case WaapV1CustomPageSetListParamsOrderingName, WaapV1CustomPageSetListParamsOrderingMinusName, WaapV1CustomPageSetListParamsOrderingID, WaapV1CustomPageSetListParamsOrderingMinusID:
		return true
	}
	return false
}
