// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package streaming

import (
	"context"
	"encoding/json"
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

// PlayerService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlayerService] method instead.
type PlayerService struct {
	Options []option.RequestOption
}

// NewPlayerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPlayerService(opts ...option.RequestOption) (r PlayerService) {
	r = PlayerService{}
	r.Options = opts
	return
}

// Create player
func (r *PlayerService) New(ctx context.Context, body PlayerNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "streaming/players"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Updates player settings
func (r *PlayerService) Update(ctx context.Context, playerID int64, body PlayerUpdateParams, opts ...option.RequestOption) (res *Player, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/players/%v", playerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of created players
func (r *PlayerService) List(ctx context.Context, query PlayerListParams, opts ...option.RequestOption) (res *pagination.PageStreaming[Player], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "streaming/players"
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

// Returns a list of created players
func (r *PlayerService) ListAutoPaging(ctx context.Context, query PlayerListParams, opts ...option.RequestOption) *pagination.PageStreamingAutoPager[Player] {
	return pagination.NewPageStreamingAutoPager(r.List(ctx, query, opts...))
}

// Delete player
func (r *PlayerService) Delete(ctx context.Context, playerID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("streaming/players/%v", playerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns player settings
func (r *PlayerService) Get(ctx context.Context, playerID int64, opts ...option.RequestOption) (res *Player, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/players/%v", playerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns player configuration in HTML
func (r *PlayerService) Preview(ctx context.Context, playerID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("streaming/players/%v/preview", playerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Set of properties for displaying videos. All parameters may be blank to inherit
// their values from default Streaming player.
type Player struct {
	// Player name
	Name string `json:"name,required"`
	// Player ID
	ID int64 `json:"id"`
	// Enables video playback right after player load:
	//
	// - **true** — video starts playing right after player loads
	// - **false** — video isn’t played automatically. A user must click play to start
	//
	// Default is false
	Autoplay bool `json:"autoplay"`
	// Color of skin background in format #AAAAAA
	BgColor string `json:"bg_color"`
	// Client ID
	ClientID int64 `json:"client_id"`
	// Custom CSS to be added to player iframe
	CustomCss string `json:"custom_css"`
	// String to be rendered as JS parameters to player
	Design string `json:"design"`
	// Enables/Disables player skin:
	//
	// - **true** — player skin is disabled
	// - **false** — player skin is enabled
	//
	// Default is false
	DisableSkin bool `json:"disable_skin"`
	// Color of skin foreground (elements) in format #AAAAAA
	FgColor string `json:"fg_color"`
	// Player framework type
	Framework string `json:"framework"`
	// Color of foreground elements when mouse is over in format #AAAAAA
	HoverColor string `json:"hover_color"`
	// Player main JS file URL. Leave empty to use JS URL from the default player
	JsURL string `json:"js_url"`
	// URL to logo image
	Logo string `json:"logo"`
	// Logotype position.
	//
	//	Has four possible values:
	//
	// - **tl** — top left
	// - **tr** — top right
	// - **bl** — bottom left
	// - **br** — bottom right
	//
	// Default is null
	LogoPosition string `json:"logo_position"`
	// Regulates the sound volume:
	//
	// - **true** — video starts with volume off
	// - **false** — video starts with volume on
	//
	// Default is false
	Mute bool `json:"mute"`
	// Enables/Disables saving volume and other options in cookies:
	//
	// - **true** — user settings will be saved
	// - **false** — user settings will not be saved
	//
	// Default is true
	SaveOptionsToCookies bool `json:"save_options_to_cookies"`
	// Enables/Disables sharing button display:
	//
	// - **true** — sharing button is displayed
	// - **false** — no sharing button is displayed
	//
	// Default is true
	ShowSharing bool `json:"show_sharing"`
	// URL to custom skin JS file
	SkinIsURL string `json:"skin_is_url"`
	// Enables/Disables speed control button display:
	//
	// - **true** — sharing button is displayed
	// - **false** — no sharing button is displayed
	//
	// Default is false
	SpeedControl bool `json:"speed_control"`
	// Color of skin text elements in format #AAAAAA
	TextColor string `json:"text_color"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name                 respjson.Field
		ID                   respjson.Field
		Autoplay             respjson.Field
		BgColor              respjson.Field
		ClientID             respjson.Field
		CustomCss            respjson.Field
		Design               respjson.Field
		DisableSkin          respjson.Field
		FgColor              respjson.Field
		Framework            respjson.Field
		HoverColor           respjson.Field
		JsURL                respjson.Field
		Logo                 respjson.Field
		LogoPosition         respjson.Field
		Mute                 respjson.Field
		SaveOptionsToCookies respjson.Field
		ShowSharing          respjson.Field
		SkinIsURL            respjson.Field
		SpeedControl         respjson.Field
		TextColor            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Player) RawJSON() string { return r.JSON.raw }
func (r *Player) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Player to a PlayerParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PlayerParam.Overrides()
func (r Player) ToParam() PlayerParam {
	return param.Override[PlayerParam](json.RawMessage(r.RawJSON()))
}

// Set of properties for displaying videos. All parameters may be blank to inherit
// their values from default Streaming player.
//
// The property Name is required.
type PlayerParam struct {
	// Player name
	Name string `json:"name,required"`
	// Player ID
	ID param.Opt[int64] `json:"id,omitzero"`
	// Enables video playback right after player load:
	//
	// - **true** — video starts playing right after player loads
	// - **false** — video isn’t played automatically. A user must click play to start
	//
	// Default is false
	Autoplay param.Opt[bool] `json:"autoplay,omitzero"`
	// Color of skin background in format #AAAAAA
	BgColor param.Opt[string] `json:"bg_color,omitzero"`
	// Client ID
	ClientID param.Opt[int64] `json:"client_id,omitzero"`
	// Custom CSS to be added to player iframe
	CustomCss param.Opt[string] `json:"custom_css,omitzero"`
	// String to be rendered as JS parameters to player
	Design param.Opt[string] `json:"design,omitzero"`
	// Enables/Disables player skin:
	//
	// - **true** — player skin is disabled
	// - **false** — player skin is enabled
	//
	// Default is false
	DisableSkin param.Opt[bool] `json:"disable_skin,omitzero"`
	// Color of skin foreground (elements) in format #AAAAAA
	FgColor param.Opt[string] `json:"fg_color,omitzero"`
	// Player framework type
	Framework param.Opt[string] `json:"framework,omitzero"`
	// Color of foreground elements when mouse is over in format #AAAAAA
	HoverColor param.Opt[string] `json:"hover_color,omitzero"`
	// Player main JS file URL. Leave empty to use JS URL from the default player
	JsURL param.Opt[string] `json:"js_url,omitzero"`
	// URL to logo image
	Logo param.Opt[string] `json:"logo,omitzero"`
	// Logotype position.
	//
	//	Has four possible values:
	//
	// - **tl** — top left
	// - **tr** — top right
	// - **bl** — bottom left
	// - **br** — bottom right
	//
	// Default is null
	LogoPosition param.Opt[string] `json:"logo_position,omitzero"`
	// Regulates the sound volume:
	//
	// - **true** — video starts with volume off
	// - **false** — video starts with volume on
	//
	// Default is false
	Mute param.Opt[bool] `json:"mute,omitzero"`
	// Enables/Disables saving volume and other options in cookies:
	//
	// - **true** — user settings will be saved
	// - **false** — user settings will not be saved
	//
	// Default is true
	SaveOptionsToCookies param.Opt[bool] `json:"save_options_to_cookies,omitzero"`
	// Enables/Disables sharing button display:
	//
	// - **true** — sharing button is displayed
	// - **false** — no sharing button is displayed
	//
	// Default is true
	ShowSharing param.Opt[bool] `json:"show_sharing,omitzero"`
	// URL to custom skin JS file
	SkinIsURL param.Opt[string] `json:"skin_is_url,omitzero"`
	// Enables/Disables speed control button display:
	//
	// - **true** — sharing button is displayed
	// - **false** — no sharing button is displayed
	//
	// Default is false
	SpeedControl param.Opt[bool] `json:"speed_control,omitzero"`
	// Color of skin text elements in format #AAAAAA
	TextColor param.Opt[string] `json:"text_color,omitzero"`
	paramObj
}

func (r PlayerParam) MarshalJSON() (data []byte, err error) {
	type shadow PlayerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlayerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlayerNewParams struct {
	// Set of properties for displaying videos. All parameters may be blank to inherit
	// their values from default Streaming player.
	Player PlayerParam `json:"player,omitzero"`
	paramObj
}

func (r PlayerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PlayerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlayerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlayerUpdateParams struct {
	// Set of properties for displaying videos. All parameters may be blank to inherit
	// their values from default Streaming player.
	Player PlayerParam `json:"player,omitzero"`
	paramObj
}

func (r PlayerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PlayerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlayerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlayerListParams struct {
	// Query parameter. Use it to list the paginated content
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PlayerListParams]'s query parameters as `url.Values`.
func (r PlayerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
