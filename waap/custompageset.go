// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
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
func (r *CustomPageSetService) New(ctx context.Context, body CustomPageSetNewParams, opts ...option.RequestOption) (res *WaapCustomPageSet, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
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
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("waap/v1/custom-page-sets/%v", setID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Retrieve a list of custom page sets available for use
func (r *CustomPageSetService) List(ctx context.Context, query CustomPageSetListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapCustomPageSet], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
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
func (r *CustomPageSetService) ListAutoPaging(ctx context.Context, query CustomPageSetListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapCustomPageSet] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a custom page set
func (r *CustomPageSetService) Delete(ctx context.Context, setID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("waap/v1/custom-page-sets/%v", setID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve a custom page set based on the provided ID
func (r *CustomPageSetService) Get(ctx context.Context, setID int64, opts ...option.RequestOption) (res *WaapCustomPageSet, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("waap/v1/custom-page-sets/%v", setID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Allows to preview a custom page without creating it based on the provided type
// and data
func (r *CustomPageSetService) Preview(ctx context.Context, params CustomPageSetPreviewParams, opts ...option.RequestOption) (res *WaapCustomPagePreview, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "waap/v1/preview-custom-page"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
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
	Name           string                          `json:"name,required"`
	Block          WaapCustomPageSetBlock          `json:"block,nullable"`
	BlockCsrf      WaapCustomPageSetBlockCsrf      `json:"block_csrf,nullable"`
	Captcha        WaapCustomPageSetCaptcha        `json:"captcha,nullable"`
	CookieDisabled WaapCustomPageSetCookieDisabled `json:"cookie_disabled,nullable"`
	// List of domain IDs that are associated with this page set
	Domains            []int64                             `json:"domains,nullable"`
	Handshake          WaapCustomPageSetHandshake          `json:"handshake,nullable"`
	JavascriptDisabled WaapCustomPageSetJavascriptDisabled `json:"javascript_disabled,nullable"`
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

type WaapCustomPageSetBlock struct {
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
func (r WaapCustomPageSetBlock) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomPageSetBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapCustomPageSetBlockCsrf struct {
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
func (r WaapCustomPageSetBlockCsrf) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomPageSetBlockCsrf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapCustomPageSetCaptcha struct {
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
func (r WaapCustomPageSetCaptcha) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomPageSetCaptcha) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapCustomPageSetCookieDisabled struct {
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
func (r WaapCustomPageSetCookieDisabled) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomPageSetCookieDisabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapCustomPageSetHandshake struct {
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
func (r WaapCustomPageSetHandshake) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomPageSetHandshake) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapCustomPageSetJavascriptDisabled struct {
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
func (r WaapCustomPageSetJavascriptDisabled) RawJSON() string { return r.JSON.raw }
func (r *WaapCustomPageSetJavascriptDisabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomPageSetNewParams struct {
	// Name of the custom page set
	Name           string                               `json:"name,required"`
	Block          CustomPageSetNewParamsBlock          `json:"block,omitzero"`
	BlockCsrf      CustomPageSetNewParamsBlockCsrf      `json:"block_csrf,omitzero"`
	Captcha        CustomPageSetNewParamsCaptcha        `json:"captcha,omitzero"`
	CookieDisabled CustomPageSetNewParamsCookieDisabled `json:"cookie_disabled,omitzero"`
	// List of domain IDs that are associated with this page set
	Domains            []int64                                  `json:"domains,omitzero"`
	Handshake          CustomPageSetNewParamsHandshake          `json:"handshake,omitzero"`
	JavascriptDisabled CustomPageSetNewParamsJavascriptDisabled `json:"javascript_disabled,omitzero"`
	paramObj
}

func (r CustomPageSetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type CustomPageSetNewParamsBlock struct {
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

func (r CustomPageSetNewParamsBlock) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetNewParamsBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetNewParamsBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type CustomPageSetNewParamsBlockCsrf struct {
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

func (r CustomPageSetNewParamsBlockCsrf) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetNewParamsBlockCsrf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetNewParamsBlockCsrf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type CustomPageSetNewParamsCaptcha struct {
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

func (r CustomPageSetNewParamsCaptcha) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetNewParamsCaptcha
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetNewParamsCaptcha) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type CustomPageSetNewParamsCookieDisabled struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header param.Opt[string] `json:"header,omitzero"`
	// The text to display in the body of the custom page
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r CustomPageSetNewParamsCookieDisabled) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetNewParamsCookieDisabled
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetNewParamsCookieDisabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type CustomPageSetNewParamsHandshake struct {
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

func (r CustomPageSetNewParamsHandshake) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetNewParamsHandshake
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetNewParamsHandshake) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type CustomPageSetNewParamsJavascriptDisabled struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header param.Opt[string] `json:"header,omitzero"`
	// The text to display in the body of the custom page
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r CustomPageSetNewParamsJavascriptDisabled) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetNewParamsJavascriptDisabled
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetNewParamsJavascriptDisabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomPageSetUpdateParams struct {
	// Name of the custom page set
	Name           param.Opt[string]                       `json:"name,omitzero"`
	Block          CustomPageSetUpdateParamsBlock          `json:"block,omitzero"`
	BlockCsrf      CustomPageSetUpdateParamsBlockCsrf      `json:"block_csrf,omitzero"`
	Captcha        CustomPageSetUpdateParamsCaptcha        `json:"captcha,omitzero"`
	CookieDisabled CustomPageSetUpdateParamsCookieDisabled `json:"cookie_disabled,omitzero"`
	// List of domain IDs that are associated with this page set
	Domains            []int64                                     `json:"domains,omitzero"`
	Handshake          CustomPageSetUpdateParamsHandshake          `json:"handshake,omitzero"`
	JavascriptDisabled CustomPageSetUpdateParamsJavascriptDisabled `json:"javascript_disabled,omitzero"`
	paramObj
}

func (r CustomPageSetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type CustomPageSetUpdateParamsBlock struct {
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

func (r CustomPageSetUpdateParamsBlock) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetUpdateParamsBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetUpdateParamsBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type CustomPageSetUpdateParamsBlockCsrf struct {
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

func (r CustomPageSetUpdateParamsBlockCsrf) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetUpdateParamsBlockCsrf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetUpdateParamsBlockCsrf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type CustomPageSetUpdateParamsCaptcha struct {
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

func (r CustomPageSetUpdateParamsCaptcha) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetUpdateParamsCaptcha
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetUpdateParamsCaptcha) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type CustomPageSetUpdateParamsCookieDisabled struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header param.Opt[string] `json:"header,omitzero"`
	// The text to display in the body of the custom page
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r CustomPageSetUpdateParamsCookieDisabled) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetUpdateParamsCookieDisabled
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetUpdateParamsCookieDisabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type CustomPageSetUpdateParamsHandshake struct {
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

func (r CustomPageSetUpdateParamsHandshake) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetUpdateParamsHandshake
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetUpdateParamsHandshake) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type CustomPageSetUpdateParamsJavascriptDisabled struct {
	// Indicates whether the custom custom page is active or inactive
	Enabled bool `json:"enabled,required"`
	// The text to display in the header of the custom page
	Header param.Opt[string] `json:"header,omitzero"`
	// The text to display in the body of the custom page
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r CustomPageSetUpdateParamsJavascriptDisabled) MarshalJSON() (data []byte, err error) {
	type shadow CustomPageSetUpdateParamsJavascriptDisabled
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomPageSetUpdateParamsJavascriptDisabled) UnmarshalJSON(data []byte) error {
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
