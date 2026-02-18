// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package streaming

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// StreamClipService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStreamClipService] method instead.
type StreamClipService struct {
	Options []option.RequestOption
}

// NewStreamClipService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamClipService(opts ...option.RequestOption) (r StreamClipService) {
	r = StreamClipService{}
	r.Options = opts
	return
}

// Create an instant clip from on-going live stream.
//
// Instant clips are applicable in cases where there is no time to wait for the
// broadcast to be completed and recorded. For example, for quickly cutting
// highlights in sport events, or cutting an important moment in the news or live
// performance.
//
// DVR function must be enabled for clip recording. If the DVR is disabled, the
// response will be error 422.
//
// Instant clip becomes available for viewing in the following formats:
//
// - HLS .m3u8,
// - MP4,
// - VOD in video hosting with a permanent link to watch video.
//
// ![HTML Overlays](https://demo-files.gvideo.io/apidocs/clip_recording_mp4_hls.gif)
//
// **Clip lifetime:**
//
// Instant clips are a copy of the stream, created from a live stream. They are
// stored in memory for a limited time, after which the clip ceases to exist and
// you will receive a 404 on the link.
//
// Limits that you should keep in mind:
//
//   - The clip's lifespan is controlled by `expiration` parameter.
//   - The default expiration value is 1 hour. The value can be set from 1 minute to
//     4 hours.
//   - If you want a video for longer or permanent viewing, then create a regular VOD
//     based on the clip. This way you can use the clip's link for the first time,
//     and immediately after the transcoded version is ready, you can change by
//     yourself it to a permanent link of VOD.
//   - The clip becomes available only after it is completely copied from the live
//     stream. So the clip will be available after `start + duration` exact time. If
//     you try to request it before this time, the response will be error code 425
//     "Too Early".
//
// **Cutting a clip from a source:**
//
// In order to use clips recording feature, DVR must be enabled for a stream:
// "dvr_enabled: true". The DVR serves as a source for creating clips:
//
//   - By default live stream DVR is set to 1 hour (3600 seconds). You can create an
//     instant clip using any segment of this time period by specifying the desired
//     start time and duration.
//   - If you create a clip, but the DVR expires, the clip will still exist for the
//     specified time as a copy of the stream.
//
// **Getting permanent VOD:**
//
// To get permanent VOD version of a live clip use this parameter when making a
// request to create a clip: `vod_required: true`.
//
// Later, when the clip is ready, grab `video_id` value from the response and query
// the video by regular GET /video/{id} method.
func (r *StreamClipService) New(ctx context.Context, streamID int64, body StreamClipNewParams, opts ...option.RequestOption) (res *Clip, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/streams/%v/clip_recording", streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get list of non expired instant clips for a stream.
//
// You can now use both MP4 just-in-time packager and HLS for all clips. Get URLs
// from "hls_master" and "mp4_master".
//
// **How to download renditions of clips:**
//
// URLs contain "master" alias by default, which means maximum available quality
// from ABR set (based on height metadata). There is also possibility to access
// individual bitrates from ABR ladder. That works for both HLS and MP4. You can
// replace manually "master" to a value from renditions list in order to get exact
// bitrate/quality from the set. Example:
//
//   - HLS 720p:
//     `https://CID.domain.com/rec/111_1000/rec_d7bsli54p8n4_qsid42_master.m3u8`
//   - HLS 720p:
//     `https://CID.domain.com/rec/111_1000/rec_d7bsli54p8n4_qsid42_media_1_360.m3u8`
//   - MP4 360p:
//     `https://CID.domain.com/rec/111_1000/rec_d7bsli54p8n4_qsid42_master.mp4`
//   - MP4 360p:
//     `https://CID.domain.com/rec/111_1000/rec_d7bsli54p8n4_qsid42_media_1_360.mp4`
func (r *StreamClipService) List(ctx context.Context, streamID int64, opts ...option.RequestOption) (res *[]Clip, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/streams/%v/clip_recording", streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Clip struct {
	// ID of the clip
	ID string `json:"id,required"`
	// Requested segment duration in seconds to be cut.
	//
	// Please, note that cutting is based on the idea of instantly creating a clip,
	// instead of precise timing. So final segment may be:
	//
	//   - Less than the specified value if there is less data in the DVR than the
	//     requested segment.
	//   - Greater than the specified value, because segment is aligned to the first and
	//     last key frames of already stored fragment in DVR, this way -1 and +1 chunks
	//     can be added to left and right.
	//
	// Duration of cutted segment cannot be greater than DVR duration for this stream.
	// Therefore, to change the maximum, use "dvr_duration" parameter of this stream.
	Duration int64 `json:"duration,required"`
	// Creation date and time. Format is date time in ISO 8601
	CreatedAt string `json:"created_at"`
	// Expire time of the clip via a public link.
	//
	// Unix timestamp in seconds, absolute value.
	//
	// This is the time how long the instant clip will be stored in the server memory
	// and can be accessed via public HLS/MP4 links. Download and/or use the instant
	// clip before this time expires.
	//
	// After the time has expired, the clip is deleted from memory and is no longer
	// available via the link. You need to create a new segment, or use
	// `vod_required: true` attribute.
	//
	// If value is omitted, then expiration is counted as +3600 seconds (1 hour) to the
	// end of the clip (i.e. `unix timestamp = <start> + <duration> + 3600`).
	//
	// Allowed range: 1m <= expiration <= 4h.
	//
	// Example:
	// `24.05.2024 14:00:00 (GMT) + 60 seconds of duration + 3600 seconds of expiration = 24.05.2024 15:01:00 (GMT) is Unix timestamp = 1716562860`
	Expiration int64 `json:"expiration"`
	// Link to HLS .m3u8 with immediate clip. The link retains same adaptive bitrate as
	// in the stream for end viewers. For additional restrictions, see the description
	// of parameter "mp4_master".
	HlsMaster string `json:"hls_master"`
	// Link to MP4 with immediate clip. The link points to max rendition quality.
	// Request of the URL can return:
	//
	//   - 200 OK – if the clip exists.
	//   - 404 Not found – if the clip did not exist or has already ceased to exist.
	//   - 425 Too early – if recording is on-going now. The file is incomplete and will
	//     be accessible after start+duration time will come.
	MP4Master string `json:"mp4_master"`
	// List of available rendition heights
	Renditions []string `json:"renditions"`
	// Starting point of the segment to cut.
	//
	// Unix timestamp in seconds, absolute value. Example:
	// `24.05.2024 14:00:00 (GMT) is Unix timestamp = 1716559200`
	//
	// If a value from the past is specified, it is used as the starting point for the
	// segment to cut. If the value is omitted, then clip will start from now.
	Start int64 `json:"start"`
	// ID of the created video if `vod_required`=true
	VideoID int64 `json:"video_id"`
	// Indicates if video needs to be stored as VOD
	VodRequired bool `json:"vod_required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Duration    respjson.Field
		CreatedAt   respjson.Field
		Expiration  respjson.Field
		HlsMaster   respjson.Field
		MP4Master   respjson.Field
		Renditions  respjson.Field
		Start       respjson.Field
		VideoID     respjson.Field
		VodRequired respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Clip) RawJSON() string { return r.JSON.raw }
func (r *Clip) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamClipNewParams struct {
	// Requested segment duration in seconds to be cut.
	//
	// Please, note that cutting is based on the idea of instantly creating a clip,
	// instead of precise timing. So final segment may be:
	//
	//   - Less than the specified value if there is less data in the DVR than the
	//     requested segment.
	//   - Greater than the specified value, because segment is aligned to the first and
	//     last key frames of already stored fragment in DVR, this way -1 and +1 chunks
	//     can be added to left and right.
	//
	// Duration of cutted segment cannot be greater than DVR duration for this stream.
	// Therefore, to change the maximum, use "dvr_duration" parameter of this stream.
	Duration int64 `json:"duration,required"`
	// Expire time of the clip via a public link.
	//
	// Unix timestamp in seconds, absolute value.
	//
	// This is the time how long the instant clip will be stored in the server memory
	// and can be accessed via public HLS/MP4 links. Download and/or use the instant
	// clip before this time expires.
	//
	// After the time has expired, the clip is deleted from memory and is no longer
	// available via the link. You need to create a new segment, or use
	// `vod_required: true` attribute.
	//
	// If value is omitted, then expiration is counted as +3600 seconds (1 hour) to the
	// end of the clip (i.e. `unix timestamp = <start> + <duration> + 3600`).
	//
	// Allowed range: 1m <= expiration <= 4h.
	//
	// Example:
	// `24.05.2024 14:00:00 (GMT) + 60 seconds of duration + 3600 seconds of expiration = 24.05.2024 15:01:00 (GMT) is Unix timestamp = 1716562860`
	Expiration param.Opt[int64] `json:"expiration,omitzero"`
	// Starting point of the segment to cut.
	//
	// Unix timestamp in seconds, absolute value. Example:
	// `24.05.2024 14:00:00 (GMT) is Unix timestamp = 1716559200`
	//
	// If a value from the past is specified, it is used as the starting point for the
	// segment to cut. If the value is omitted, then clip will start from now.
	Start param.Opt[int64] `json:"start,omitzero"`
	// Indicates if video needs to be stored also as permanent VOD
	VodRequired param.Opt[bool] `json:"vod_required,omitzero"`
	paramObj
}

func (r StreamClipNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StreamClipNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StreamClipNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
