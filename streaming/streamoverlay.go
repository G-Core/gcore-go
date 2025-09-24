// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package streaming

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	shimjson "github.com/G-Core/gcore-go/internal/encoding/json"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// StreamOverlayService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStreamOverlayService] method instead.
type StreamOverlayService struct {
	Options []option.RequestOption
}

// NewStreamOverlayService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamOverlayService(opts ...option.RequestOption) (r StreamOverlayService) {
	r = StreamOverlayService{}
	r.Options = opts
	return
}

// "Overlay" is a live HTML widget, which rendered and inserted over the live
// stream.
//
// There are can be more that 1 overlay over a stream, which are small or stretched
// over full frame. Overlays can have transparent areas. Frequency of update is 1
// FPS. Automatic size scaling for Adaptative Bitrate qualities is applied.
//
// ![HTML Overlays](https://demo-files.gvideo.io/apidocs/coffee_run_overlays.gif)
//
// How to activate and use in simple steps:
//
// - Activate feature on your account, ask the Support Team
// - Set “`html_overlay`” attribute to "true" for a stream
// - Set array of overlays
// - Start or restart your stream again
// - Enjoy :-)
//
// For the first time an overlay should be enabled **before** start pushing of a
// live stream. If you are pushing the stream already (stream is alive and you are
// activating overlay for the first time), then overlay will become active after
// restart pushing.
//
// Once you activate the overlay for the stream for the first time, you can add,
// change, move, delete widgets on the fly even during a live stream with no
// affection on a result stream.
//
// Tech limits:
//
//   - Max original stream resolution = FullHD.
//   - It is necessary that all widgets must fit into the original frame of the
//     source stream (width x height). If one of the widgets does not fit into the
//     original frame, for example, goes 1 pixel beyond the frame, then all widgets
//     will be hidden.
//   - Attributes of overlays:
//   - url – should be valid http/https url
//   - 0 < width <= 1920
//   - 0 < height <= 1080
//   - 0 <= x < 1920
//   - 0 <= y < 1080
//   - stretch – stretch to full frame. Cannot be used with positioning attributes.
//   - HTML widget can be access by HTTP 80 or HTTPS 443 ports.
//   - HTML page code at the "url" link is read once when starting the stream only.
//     For dynamically updating widgets, you must use either dynamic code via
//     JavaScript or cause a page refresh via HTML meta tag
//     <meta http-equiv="refresh" content="N">.
//   - Widgets can contain scripts, but they must be lightweight and using small
//     amount memory, CPU, and bandwidth. It is prohibited to run heavy scripts,
//     create a heavy load on the network, or run other heavy modules. Such widgets
//     can be stopped automatically, and the ability to insert widgets itself is
//     banned.
//   - If feature is disabled, you will receive HTTP code: 422. Error text: Feature
//     disabled. Contact support to enable.
//
// Please, pay attention to the content of HTML widges you use. If you don't trust
// them, then you shouldn't use them, as their result will be displayed in live
// stream to all users.
//
// **Will there be a widget in the recording?** Right now overlay widgets are sent
// to the end viewer in the HLS/DASH streams, but are not recorded due to technical
// limitations. We are working to ensure that widgets remain in the recordings as
// well. Follow the news.
func (r *StreamOverlayService) New(ctx context.Context, streamID int64, body StreamOverlayNewParams, opts ...option.RequestOption) (res *[]Overlay, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("streaming/streams/%v/overlays", streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates overlay settings
func (r *StreamOverlayService) Update(ctx context.Context, overlayID int64, params StreamOverlayUpdateParams, opts ...option.RequestOption) (res *Overlay, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("streaming/streams/%v/overlays/%v", params.StreamID, overlayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Returns a list of HTML overlay widgets which are attached to a stream
func (r *StreamOverlayService) List(ctx context.Context, streamID int64, opts ...option.RequestOption) (res *[]Overlay, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("streaming/streams/%v/overlays", streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an overlay
func (r *StreamOverlayService) Delete(ctx context.Context, overlayID int64, body StreamOverlayDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("streaming/streams/%v/overlays/%v", body.StreamID, overlayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get overlay details
func (r *StreamOverlayService) Get(ctx context.Context, overlayID int64, query StreamOverlayGetParams, opts ...option.RequestOption) (res *Overlay, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("streaming/streams/%v/overlays/%v", query.StreamID, overlayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates settings for set of overlays
func (r *StreamOverlayService) UpdateMultiple(ctx context.Context, streamID int64, body StreamOverlayUpdateMultipleParams, opts ...option.RequestOption) (res *[]Overlay, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("streaming/streams/%v/overlays", streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type Overlay struct {
	// ID of the overlay
	ID int64 `json:"id,required"`
	// Datetime of creation in ISO 8601
	CreatedAt string `json:"created_at,required"`
	// ID of a stream to which it is attached
	StreamID int64 `json:"stream_id,required"`
	// Datetime of last update in ISO 8601
	UpdatedAt string `json:"updated_at,required"`
	// Valid http/https URL to an HTML page/widget
	URL string `json:"url,required"`
	// Height of the widget
	Height int64 `json:"height"`
	// Switch of auto scaling the widget. Must not be used as "true" simultaneously
	// with the coordinate installation method (w, h, x, y).
	Stretch bool `json:"stretch"`
	// Width of the widget
	Width int64 `json:"width"`
	// Coordinate of left upper corner
	X int64 `json:"x"`
	// Coordinate of left upper corner
	Y int64 `json:"y"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		StreamID    respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		Height      respjson.Field
		Stretch     respjson.Field
		Width       respjson.Field
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Overlay) RawJSON() string { return r.JSON.raw }
func (r *Overlay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamOverlayNewParams struct {
	Body []StreamOverlayNewParamsBody
	paramObj
}

func (r StreamOverlayNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *StreamOverlayNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The property URL is required.
type StreamOverlayNewParamsBody struct {
	// Valid http/https URL to an HTML page/widget
	URL string `json:"url,required"`
	// Height of the widget
	Height param.Opt[int64] `json:"height,omitzero"`
	// Switch of auto scaling the widget. Must not be used as "true" simultaneously
	// with the coordinate installation method (w, h, x, y).
	Stretch param.Opt[bool] `json:"stretch,omitzero"`
	// Width of the widget
	Width param.Opt[int64] `json:"width,omitzero"`
	// Coordinate of left upper corner
	X param.Opt[int64] `json:"x,omitzero"`
	// Coordinate of left upper corner
	Y param.Opt[int64] `json:"y,omitzero"`
	paramObj
}

func (r StreamOverlayNewParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow StreamOverlayNewParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StreamOverlayNewParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamOverlayUpdateParams struct {
	StreamID int64 `path:"stream_id,required" json:"-"`
	// Height of the widget
	Height param.Opt[int64] `json:"height,omitzero"`
	// Switch of auto scaling the widget. Must not be used as "true" simultaneously
	// with the coordinate installation method (w, h, x, y).
	Stretch param.Opt[bool] `json:"stretch,omitzero"`
	// Valid http/https URL to an HTML page/widget
	URL param.Opt[string] `json:"url,omitzero"`
	// Width of the widget
	Width param.Opt[int64] `json:"width,omitzero"`
	// Coordinate of left upper corner
	X param.Opt[int64] `json:"x,omitzero"`
	// Coordinate of left upper corner
	Y param.Opt[int64] `json:"y,omitzero"`
	paramObj
}

func (r StreamOverlayUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow StreamOverlayUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StreamOverlayUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamOverlayDeleteParams struct {
	StreamID int64 `path:"stream_id,required" json:"-"`
	paramObj
}

type StreamOverlayGetParams struct {
	StreamID int64 `path:"stream_id,required" json:"-"`
	paramObj
}

type StreamOverlayUpdateMultipleParams struct {
	Body []StreamOverlayUpdateMultipleParamsBody
	paramObj
}

func (r StreamOverlayUpdateMultipleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *StreamOverlayUpdateMultipleParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The property ID is required.
type StreamOverlayUpdateMultipleParamsBody struct {
	// ID of the overlay
	ID int64 `json:"id,required"`
	// Height of the widget
	Height param.Opt[int64] `json:"height,omitzero"`
	// Switch of auto scaling the widget. Must not be used as "true" simultaneously
	// with the coordinate installation method (w, h, x, y).
	Stretch param.Opt[bool] `json:"stretch,omitzero"`
	// Valid http/https URL to an HTML page/widget
	URL param.Opt[string] `json:"url,omitzero"`
	// Width of the widget
	Width param.Opt[int64] `json:"width,omitzero"`
	// Coordinate of left upper corner
	X param.Opt[int64] `json:"x,omitzero"`
	// Coordinate of left upper corner
	Y param.Opt[int64] `json:"y,omitzero"`
	paramObj
}

func (r StreamOverlayUpdateMultipleParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow StreamOverlayUpdateMultipleParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StreamOverlayUpdateMultipleParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
