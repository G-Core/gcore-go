// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package streaming

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

// StreamService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStreamService] method instead.
type StreamService struct {
	Options  []option.RequestOption
	Overlays StreamOverlayService
}

// NewStreamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStreamService(opts ...option.RequestOption) (r StreamService) {
	r = StreamService{}
	r.Options = opts
	r.Overlays = NewStreamOverlayService(opts...)
	return
}

// Use this method to create a new live stream entity for broadcasting.
//
// The input in API may contain streams of different formats, including the most
// common ones RTMP, RTMPS, SRT, HLS. Note that multicast MPEG-TS over UDP and
// others are supported too, ask the Support Team please.
//
// For ingestion, you can use both PUSH and PULL methods.
//
// Also you can use the main and backup servers, which are geographically located
// in different locations. By default, any free ingest points in the world are
// used. Settings have been applied that deliver low-latency streams in the optimal
// way. If for some reason you need to set a fixed ingest point, or if you need to
// set the main and backup ingest points in the same region (for example, do not
// send streams outside the EU or US), then contact our Support Team.
//
// The output is HLS and MPEG-DASH with ABR. We transcode video for you by our
// cloud-based infrastructure. ABR ladder supports all qualities from SD to 8K HDR
// 60fps.
//
// All our streams are Low Latency enabled. We support a delay of ±4 seconds for
// video streams by utilizing Common Media Application Format (CMAF) technology. So
// you obtain latency from the traditional 30-50 seconds to ±4 seconds only by
// default. If you need legacy non-low-latency HLS, then look at HLS MPEG-TS
// delivery below.
//
// You have access to additional functions such as:
//
// - DVR
// - Recording
// - Live clipping
// - Restreaming
// - (soon) AI Automatic Speech Recognition for subtitles/captions generating
//
// For more information see specific API methods, and the Knowledge Base. To
// organize streaming with ultra-low latency, look for WebRTC delivery in different
// section in the Knowledge Base.
//
// ![HTML Overlays](https://demo-files.gvideo.io/apidocs/low-latency-football.gif)
func (r *StreamService) New(ctx context.Context, body StreamNewParams, opts ...option.RequestOption) (res *Stream, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "streaming/streams"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates stream settings
func (r *StreamService) Update(ctx context.Context, streamID int64, body StreamUpdateParams, opts ...option.RequestOption) (res *Stream, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/streams/%v", streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of streams
func (r *StreamService) List(ctx context.Context, query StreamListParams, opts ...option.RequestOption) (res *pagination.PageStreaming[Stream], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "streaming/streams"
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

// Returns a list of streams
func (r *StreamService) ListAutoPaging(ctx context.Context, query StreamListParams, opts ...option.RequestOption) *pagination.PageStreamingAutoPager[Stream] {
	return pagination.NewPageStreamingAutoPager(r.List(ctx, query, opts...))
}

// Delete a live stream.
//
// After deleting the live stream, all associated data is deleted: settings, PUSH
// and PULL links, video playback links, etc.
//
// Live stream information is deleted permanently and irreversibly. Therefore, it
// is impossible to restore data and files after this.
//
// But if the live had recordings, they continue to remain independent Video
// entities. The "`stream_id`" parameter will simply point to a stream that no
// longer exists.
//
// Perhaps, instead of deleting, you may use the stream deactivation:
//
// ```
// PATCH /videos/{stream_id}
// { "active": false }
// ```
//
// For details, see the Product Documentation.
func (r *StreamService) Delete(ctx context.Context, streamID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("streaming/streams/%v", streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Clear live stream DVR
func (r *StreamService) ClearDvr(ctx context.Context, streamID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("streaming/streams/%v/dvr_cleanup", streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

// Create an instant clip from on-going live stream.
//
// Instant clips are applicable in cases where there is no time to wait for the
// broadcast to be completed and recorded. For example, for quickly cutting
// highlights in sport events, or cutting an important moment in the news or live
// performance.
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
// "`dvr_enabled`: true". The DVR serves as a source for creating clips:
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
func (r *StreamService) NewClip(ctx context.Context, streamID int64, body StreamNewClipParams, opts ...option.RequestOption) (res *Clip, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/streams/%v/clip_recording", streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns stream details
func (r *StreamService) Get(ctx context.Context, streamID int64, opts ...option.RequestOption) (res *Stream, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/streams/%v", streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get list of non expired instant clips for a stream.
//
// You can now use both MP4 just-in-time packager and HLS for all clips. Get URLs
// from "`hls_master`" and "`mp4_master`".
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
func (r *StreamService) ListClips(ctx context.Context, streamID int64, opts ...option.RequestOption) (res *[]Clip, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/streams/%v/clip_recording", streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Start recording a stream.
//
// Stream will be recorded and automatically saved in our video hosting as a
// separate video VOD:
//
//   - ID of the stream from which the recording was organized is added to
//     "`stream_id`" field. You can find the video by that value later.
//   - Title of the video is based on pattern "Stream Record: {`stream_title`},
//     {`recording_end_time_utc`}".
//   - Recording start time is stored in "`recording_started_at`" field.
//   - You can record the original stream or the transcoded one. Only the transcoded
//     version will contain overlays. Set the appropriate recording method when
//     creating the stream or before calling this recording method. Details in the
//     "`record_type`" parameter of the stream.
//   - If you have access to the premium feature of saving the original stream (so
//     not just transcoded renditions), then the link to the original file will be in
//     the "`origin_url`" field. Look at the description of the field how to use it.
//
// Stream must be live for the recording to start, please check fields "live"
// and/or "`backup_live`". After the recording starts, field "recording" will
// switch to "true", and the recording duration in seconds will appear in the
// "`recording_duration`" field.
//
// Please, keep in mind that recording doesn't start instantly, it takes ±3-7
// seconds to initialize the process after executing this method.
//
// Stream recording stops when:
//
//   - Explicit execution of the method /`stop_recording`. In this case, the file
//     will be completely saved and closed. When you execute the stream recording
//     method again, the recording will be made to a new video file.
//   - When sending the stream stops on the client side, or stops accidentally. In
//     this case, recording process is waiting for 10 seconds to resume recording:
//
//   - If the stream resumes within that period, recording will continue to the same
//     file.
//   - After that period, the file will be completely saved and closed.
//   - If the stream suddenly resumes after this period, the recording will go to a
//     new file, because old file is closed already. Please, also note that if you
//     have long broadcasts, the recording will be cut into 4-hour videos. This value
//     is fixed, but can be changed upon request to the Support Team.
func (r *StreamService) StartRecording(ctx context.Context, streamID int64, opts ...option.RequestOption) (res *StreamStartRecordingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/streams/%v/start_recording", streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Stop recording a stream.
//
// Stream must be in "recording: true" state for recording to be stopped.
//
// If there was a recording, the created video entity will be returned. Otherwise
// the response will be empty. Please see conditions and restrictions for recording
// a stream in the description of method /`start_recording`.
func (r *StreamService) StopRecording(ctx context.Context, streamID int64, opts ...option.RequestOption) (res *Video, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/streams/%v/stop_recording", streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
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
	// Therefore, to change the maximum, use "`dvr_duration`" parameter of this stream.
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
	// of parameter "`mp4_master`".
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

type Stream struct {
	// Stream name.
	//
	// Often used as a human-readable name for the stream, but can contain any text you
	// wish. The values are not unique and may be repeated.
	//
	// Examples:
	//
	// - Conference in July
	// - Stream #10003
	// - Open-Air Camera #31 Backstage
	// - 480fd499-2de2-4988-bc1a-a4eebe9818ee
	Name string `json:"name,required"`
	// Stream ID
	ID int64 `json:"id"`
	// Stream switch between on and off. This is not an indicator of the status "stream
	// is receiving and it is LIVE", but rather an on/off switch.
	//
	// When stream is switched off, there is no way to process it: PULL is deactivated
	// and PUSH will return an error.
	//
	// - true – stream can be processed
	// - false – stream is off, and cannot be processed
	Active bool `json:"active"`
	// Enables autotomatic recording of the stream when it started. So you don't need
	// to call recording manually.
	//
	// Result of recording is automatically added to video hosting. For details see the
	// /streams/`start_recording` method and in knowledge base
	//
	// Values:
	//
	// - true – auto recording is enabled
	// - false – auto recording is disabled
	AutoRecord bool `json:"auto_record"`
	// State of receiving and transcoding master stream from source by backup server if
	// you pushing stream to "`backup_push_url`" or "`backup_push_url_srt`".
	//
	// Displays the backup server status of PUSH method only. For PULL a "live" field
	// is always used, even when origin servers are switched using round robin
	// scheduling (look "uri" field for details).
	BackupLive bool `json:"backup_live"`
	// URL to PUSH master stream to our backup server using RTMP/S protocols. Servers
	// for the main and backup streams are distributed geographically.
	//
	// Mainly sending one stream to main server is enough. But if you need a backup
	// stream, then this is the field to PUSH it.
	//
	// To use RTMPS just manually change the protocol name from "rtmp://" to
	// "rtmps://".
	//
	// The backup logs are as follows: In PUSH mode, you initiate sending a stream from
	// your machine. If your stream stops or breaks for some reason and it stops coming
	// to the main server, then after 3-10 seconds of waiting the stream will turn off
	// or the backup one will be automatically turned on, if you are pushing it too.
	BackupPushURL string `json:"backup_push_url"`
	// URL to PUSH master stream to our backup server using SRT protocol with the same
	// logic of backup-streams
	BackupPushURLSrt string `json:"backup_push_url_srt"`
	// IDs of broadcasts which will include this stream
	BroadcastIDs []int64 `json:"broadcast_ids"`
	// ID of custom CDN resource from which the content will be delivered (only if you
	// know what you do)
	CdnID int64 `json:"cdn_id"`
	// Custom meta field designed to store your own extra information about a video
	// entity: video source, video id, parameters, etc. We do not use this field in any
	// way when processing the stream. You can store any data in any format (string,
	// json, etc), saved as a text string. Example:
	// `client_entity_data = '{ "seq_id": "1234567890", "name": "John Doe", "iat": 1516239022 }'`
	ClientEntityData string `json:"client_entity_data"`
	// Custom meta field for storing the Identifier in your system. We do not use this
	// field in any way when processing the stream. Example: `client_user_id = 1001`
	ClientUserID int64 `json:"client_user_id"`
	// Datetime of creation in ISO 8601
	CreatedAt string `json:"created_at"`
	// MPEG-DASH output. URL for transcoded result stream in MPEG-DASH format, with
	// .mpd link.
	//
	// Low Latency support: YES.
	//
	// This is CMAF-based MPEG-DASH stream. Encoder and packager dynamically assemble
	// the video stream with fMP4 fragments. Chunks have ±2-4 seconds duration
	// depending on the settings. All chunks for DASH are transferred through CDN using
	// chunk transfer technology, which allows to use all the advantages of low latency
	// delivery of DASH.
	//
	//   - by default low latency is ±4 sec, because it's stable for almost all last-mile
	//     use cases.
	//   - and its possible to enable ±2 sec for DASH, just ask our Support Team.
	//
	// Read more information in the article "How Low Latency streaming works" in the
	// Knowledge Base.
	DashURL string `json:"dash_url"`
	// DVR duration in seconds if DVR feature is enabled for the stream. So this is
	// duration of how far the user can rewind the live stream.
	//
	// `dvr_duration` range is [30...14400].
	//
	// Maximum value is 4 hours = 14400 seconds. If you need more, ask the Support Team
	// please.
	DvrDuration int64 `json:"dvr_duration"`
	// Enables DVR for the stream:
	//
	// - true – DVR is enabled
	// - false – DVR is disabled
	DvrEnabled bool `json:"dvr_enabled"`
	// Time when the stream ended for the last time. Datetime in ISO 8601.
	//
	// After restarting the stream, this value is not reset to "null", and the time of
	// the last/previous end is always displayed here. That is, when the start time is
	// greater than the end time, it means the current session is still ongoing and the
	// stream has not ended yet.
	//
	// If you want to see all information about acitivity of the stream, you can get it
	// from another method /streaming/statistics/ffprobe. This method shows aggregated
	// activity parameters during a time, when stream was alive and transcoded. Also
	// you can create graphs to see the activity. For example
	// /streaming/statistics/ffprobe?interval=6000&`date_from`=2023-10-01&`date_to`=2023-10-11&`stream_id`=12345
	FinishedAtPrimary string `json:"finished_at_primary"`
	// Current FPS of the original stream, if stream is transcoding
	FrameRate float64 `json:"frame_rate"`
	// HLS output. URL for transcoded result of stream in HLS CMAF format, with .m3u8
	// link. Recommended for use for all HLS streams.
	//
	// Low Latency support: YES.
	//
	// This is CMAF-based HLS stream. Encoder and packager dynamically assemble the
	// video stream with fMP4 fragments. Chunks have ±2-4 seconds duration depending on
	// the settings. All chunks for LL-HLS are transferred through CDN via dividing
	// into parts (small segments `#EXT-X-PART` of 0.5-1.0 sec duration), which allows
	// to use all the advantages of low latency delivery of LL-HLS.
	//
	//   - by default low latency is ±5 sec, because it's stable for almost all last-mile
	//     use cases.
	//   - and its possible to enable ±3 sec for LL-HLS, just ask our Support Team.
	//
	// It is also possible to use modifier-attributes, which are described in the
	// "`hls_mpegts_url`" field above. If you need to get MPEG-TS (.ts) chunks, look at
	// the attribute "`hls_mpegts_url`".
	//
	// Read more information in the article "How Low Latency streaming works" in the
	// Knowledge Base.
	HlsCmafURL string `json:"hls_cmaf_url"`
	// Add `#EXT-X-ENDLIST` tag within .m3u8 playlist after the last segment of a live
	// stream when broadcast is ended.
	HlsMpegtsEndlistTag bool `json:"hls_mpegts_endlist_tag"`
	// HLS output for legacy devices. URL for transcoded result of stream in HLS
	// MPEG-TS (.ts) format, with .m3u8 link.
	//
	// Low Latency support: NO.
	//
	// Some legacy devices or software may require MPEG-TS (.ts) segments as a format
	// for streaming, so we provide this options keeping backward compatibility with
	// any of your existing workflows. For other cases it's better to use
	// "`hls_cmaf_url`" instead.
	//
	// You can use this legacy HLSv6 format based on MPEG-TS segmenter in parallel with
	// main HLS CMAF. Both formats are sharing same segments size, manifest length
	// (DVR), etc.
	//
	// It is also possible to use additional modifier-attributes:
	//
	//   - ?`get_duration_sec`=true – Adds the real segment duration in seconds to chunk
	//     requests. A chunk duration will be automatically added to a chunk request
	//     string with the "`duration_sec`" attribute. The value is an integer for a
	//     length multiple of whole seconds, or a fractional number separated by a dot
	//     for chunks that are not multiples of seconds. This attribute allows you to
	//     determine duration in seconds at the level of analyzing the logs of CDN
	//     requests and compare it with file size (so to use it in your analytics).
	//
	// Such modifier attributes are applied manually and added to the link obtained
	// from this field. I.e. `<hls_url>?get_duration_sec=true`
	//
	// Example:
	// `https://demo.gvideo.io/mpegts/2675_19146/master_mpegts.m3u8?get_duration_sec=true`
	//
	// ```
	// #EXTM3U
	// #EXT-X-VERSION:6
	// #EXT-X-TARGETDURATION:2
	// ...
	// #EXTINF:2.000000,
	// #EXT-X-PROGRAM-DATE-TIME:2025-08-14T08:15:00
	// seg1.ts?duration_sec=2
	// ...
	// ```
	HlsMpegtsURL string `json:"hls_mpegts_url"`
	// Switch on mode to insert and display real-time HTML overlay widgets on top of
	// live streams
	HTMLOverlay bool `json:"html_overlay"`
	// Array of HTML overlay widgets
	HTMLOverlays []Overlay `json:"html_overlays"`
	// A URL to a built-in HTML web player with the stream inside. It can be inserted
	// into an iframe on your website and the video will automatically play in all
	// browsers.
	//
	// Please, remember that transcoded streams from "`hls_cmaf_url`" with .m3u8 at the
	// end, and from "`dash_url`" with .mpd at the end are to be played inside video
	// players only. For example: AVplayer on iOS, Exoplayer on Android, HTML web
	// player in browser, etc. General bowsers like Chrome, Firefox, etc cannot play
	// transcoded streams with .m3u8 and .mpd at the end. The only exception is Safari,
	// which can only play Apple's HLS .m3u8 format with limits.
	//
	// That's why you may need to use this HTML web player. Please, look Knowledge Base
	// for details.
	//
	// Example of usage on a web page:
	//
	// <iframe width="560" height="315" src="https://player.gvideo.co/streams/2675_201693" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>
	IframeURL string `json:"iframe_url"`
	// State of receiving and transcoding master stream from source by main server
	Live bool `json:"live"`
	// Visualization mode for 360° streams, how the stream is rendered in our web
	// player ONLY. If you would like to show video 360° in an external video player,
	// then use parameters of that video player.
	//
	// Modes:
	//
	// - regular – regular “flat” stream
	// - vr360 – display stream in 360° mode
	// - vr180 – display stream in 180° mode
	// - vr360tb – display stream in 3D 360° mode Top-Bottom
	//
	// Any of "regular", "vr360", "vr180", "vr360tb".
	Projection StreamProjection `json:"projection"`
	// Indicates if stream is pulled from external server or not. Has two possible
	// values:
	//
	//   - true – stream is received by PULL method. Use this when need to get stream
	//     from external server.
	//   - false – stream is received by PUSH method. Use this when need to send stream
	//     from end-device to our Streaming Platform, i.e. from your encoder, mobile app
	//     or OBS Studio.
	Pull bool `json:"pull"`
	// URL to PUSH master stream to our main server using RTMP and RTMPS protocols.
	//
	// To use RTMPS just manually change the protocol name from "rtmp://" to
	// "rtmps://".
	//
	// Use only 1 protocol of sending a master stream: eitheronly RTMP/S (`push_url`),
	// or only SRT (`push_url_srt`).
	//
	// If you see an error like "invalid SSL certificate" try the following:
	//
	//   - Make sure the push URL is correct, and it contains "rtmps://".
	//   - If the URL looks correct but you still get an SSL error, try specifying the
	//     port 443 in the URL. Here’s an example:
	//     rtmps://vp-push.domain.com:443/in/stream?key.
	//   - If you're still having trouble, then your encoder may not support RTMPS.
	//     Double-check the documentation for your encoder.
	//
	// Please note that 1 connection and 1 protocol can be used at a single moment in
	// time per unique stream key input. Trying to send 2+ connection requests into
	// `push_url` to once, or 2+ protocols at once will not lead to a result.
	//
	// For example, transcoding process will fail if:
	//
	//   - you are pushing primary and backup RTMP to the same single `push_url`
	//     simultaneously
	//   - you are pushing RTMP to `push_url` and SRT to `push_url_srt` simultaneously
	//
	// For advanced customers only: For your complexly distributed broadcast systems,
	// it is also possible to additionally output an array of multi-regional ingestion
	// points for manual selection from them. To activate this mode, contact your
	// manager or the Support Team to activate the "`multi_region_push_urls`" attibute.
	// But if you clearly don’t understand why you need this, then it’s best to use the
	// default single URL in the "`push_url`" attribute.
	PushURL string `json:"push_url"`
	// URL to PUSH master stream to our main server using SRT protocol.
	//
	// Use only 1 protocol of sending a master stream: eitheronly RTMP/S (`push_url`),
	// or only SRT (`push_url_srt`).
	//
	// **Setup SRT latency on your sender side**
	//
	// SRT is designed as a low-latency transport protocol, but real networks are not
	// always stable and in some cases the end-to-end path from the venue to the ingest
	// point can be long. For this reason, it is important to configure the latency
	// parameter carefully to match the actual network conditions.
	//
	// Small latency values may lead to packet loss when jitter or retransmissions
	// occur, while very large values introduce unnecessary end-to-end delay.
	// \*Incorrect or low default value is one of the most common reasons for packet
	// loss, frames loss, and bad picture.\*
	//
	// We therefore recommend setting latency manually rather than relying on the
	// default, to ensure the buffer is correctly sized for your environment. A
	// practical range is 400–2000 ms, with the exact value chosen based on RTT,
	// jitter, and expected packet loss.
	//
	// Be sure to check and test SRT settings on your sender side. The default values
	// do not take into account your specific scenarios and do not work well. If
	// necessary, ask us and we will help you.
	//
	// Please note that 1 connection and 1 protocol can be used at a single moment in
	// time per unique stream key input. Trying to send 2+ connection requests into
	// `push_url_srt` to once, or 2+ protocols at once will not lead to a result.
	//
	// For example, transcoding process will fail if:
	//
	//   - you are pushing primary and backup SRT to the same single `push_url_srt`
	//     simultaneously
	//   - you are pushing RTMP to `push_url` and SRT to `push_url_srt` simultaneously
	//
	// See more information and best practices about SRT protocol in the Product
	// Documentation.
	PushURLSrt string `json:"push_url_srt"`
	// URL to PUSH WebRTC stream to our server using WHIP protocol.
	//
	// **WebRTC WHIP to LL-HLS and DASH**
	//
	// Video Streaming supports WebRTC HTTP Ingest Protocol (WHIP), and WebRTC to
	// HLS/DASH converter. As a result you can stream from web broswers natively.
	//
	// **WebRTC WHIP server**
	//
	// We have dedicated WebRTC WHIP servers in our infrastructure. WebRTC WHIP server
	// organizes both signaling and receives video data. Signaling is a term to
	// describe communication between WebRTC endpoints, needed to initiate and maintain
	// a session. WHIP is an open specification for a simple signaling protocol for
	// starting WebRTC sessions in an outgoing direction, (i.e., streaming from your
	// device).
	//
	// There is the primary link only for WHIP, so no backup link.
	//
	// **WebRTC stream encoding parameters**
	//
	// At least one video and audio track both must be present in the stream:
	//
	// - Video must be encoded with H.264.
	// - Audio must be encoded with OPUS.
	//
	// Note. Specifically for WebRTC mode a method of constant transcoding with an
	// initial given resolution is used. This means that if WebRTC in the end-user's
	// browser decides to reduce the quality or resolution of the master stream (to let
	// say 360p) due to restrictions on the end-user's device (network conditions, CPU
	// consumption, etc.), the transcoder will still continue to transcode the reduced
	// stream to the initial resolution (let say 1080p ABR). When the restrictions on
	// the end-user's device are removed, quiality will improve again.
	//
	// **WebRTC WHIP Client**
	//
	// We provide a convenient WebRTC WHIP library for working in browsers. You can use
	// our library, or any other you prefer. Simple example of usage is here:
	// https://stackblitz.com/edit/stackblitz-starters-j2r9ar?file=index.html
	//
	// Also try to use the feature in UI of the Customer Portal. In the Streaming
	// section inside the settings of a specific live stream, a new section "Quick
	// start in browser" has been added.
	//
	// Please note that 1 connection and 1 protocol can be used at a single moment in
	// time per unique stream key input. Trying to send 2+ connection requests into
	// `push_url_whip` to once, or 2+ protocols at once will not lead to a result.
	//
	// For example, transcoding process will fail if:
	//
	//   - you are pushing primary and backup WHIP to the same single `push_url_whip`
	//     simultaneously
	//   - you are pushing WHIP to `push_url_whip` and RTMP to `push_url` simultaneously
	//
	// More information in the Product Documentation on the website.
	PushURLWhip string `json:"push_url_whip"`
	// Custom quality set ID for transcoding, if transcoding is required according to
	// your conditions. Look at GET /`quality_sets` method
	QualitySetID int64 `json:"quality_set_id"`
	// Method of recording a stream. Specifies the source from which the stream will be
	// recorded: original or transcoded.
	//
	// Types:
	//
	//   - "origin" – To record RMTP/SRT/etc original clean media source.
	//   - "transcoded" – To record the output transcoded version of the stream,
	//     including overlays, texts, logos, etc. additional media layers.
	//
	// Any of "origin", "transcoded".
	RecordType StreamRecordType `json:"record_type"`
	// Duration of current recording in seconds if recording is enabled for the stream
	RecordingDuration float64 `json:"recording_duration"`
	// An instant screenshot taken from a live stream, and available as a static JPEG
	// image. Resolution 1080 pixels wide, or less if the original stream has a lower
	// resolution.
	//
	// Screenshot is taken every 10 seconds while the stream is live. This field
	// contains a link to the last screenshot created by the system. Screenshot history
	// is not stored, so if you need a series of screenshots over time, then download
	// them.
	Screenshot string `json:"screenshot"`
	// Time of the last session when backup server started receiving the stream.
	// Datetime in ISO 8601
	StartedAtBackup string `json:"started_at_backup"`
	// Time of the last session when main server started receiving the stream. Datetime
	// in ISO 8601.
	//
	// This means that if the stream was started 1 time, then here will be the time it
	// was started. If the stream was started several times, or restarted on your side,
	// then only the time of the last session is displayed here.
	StartedAtPrimary string `json:"started_at_primary"`
	// Array of qualities to which live stream is transcoded
	TranscodedQualities []string `json:"transcoded_qualities"`
	// Speed of transcoding the stream.
	//
	// Mainly it must be 1.0 for real-time processing. May be less than 1.0 if your
	// stream has problems in delivery due to your local internet provider's
	// conditions, or the stream does not meet stream inbound requirements. See
	// Knowledge Base for details.
	TranscodingSpeed float64 `json:"transcoding_speed"`
	// When using PULL method, this is the URL to pull a stream from.
	//
	// You can specify multiple addresses separated by a space (" "), so you can
	// organize a backup plan. In this case, the specified addresses will be selected
	// one by one using round robin scheduling. If the first address does not respond,
	// then the next one in the list will be automatically requested, returning to the
	// first and so on in a circle. Also, if the sucessfully working stream stops
	// sending data, then the next one will be selected according to the same scheme.
	//
	// After 2 hours of inactivity of your original stream, the system stops PULL
	// requests and the stream is deactivated (the "active" field switches to "false").
	//
	// Please, note that this field is for PULL only, so is not suitable for PUSH. Look
	// at fields "`push_url`" and "`push_url_srt`" from GET method.
	Uri string `json:"uri"`
	// Current height of frame of the original stream, if stream is transcoding
	VideoHeight float64 `json:"video_height"`
	// Current width of frame of the original stream, if stream is transcoding
	VideoWidth float64 `json:"video_width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name                respjson.Field
		ID                  respjson.Field
		Active              respjson.Field
		AutoRecord          respjson.Field
		BackupLive          respjson.Field
		BackupPushURL       respjson.Field
		BackupPushURLSrt    respjson.Field
		BroadcastIDs        respjson.Field
		CdnID               respjson.Field
		ClientEntityData    respjson.Field
		ClientUserID        respjson.Field
		CreatedAt           respjson.Field
		DashURL             respjson.Field
		DvrDuration         respjson.Field
		DvrEnabled          respjson.Field
		FinishedAtPrimary   respjson.Field
		FrameRate           respjson.Field
		HlsCmafURL          respjson.Field
		HlsMpegtsEndlistTag respjson.Field
		HlsMpegtsURL        respjson.Field
		HTMLOverlay         respjson.Field
		HTMLOverlays        respjson.Field
		IframeURL           respjson.Field
		Live                respjson.Field
		Projection          respjson.Field
		Pull                respjson.Field
		PushURL             respjson.Field
		PushURLSrt          respjson.Field
		PushURLWhip         respjson.Field
		QualitySetID        respjson.Field
		RecordType          respjson.Field
		RecordingDuration   respjson.Field
		Screenshot          respjson.Field
		StartedAtBackup     respjson.Field
		StartedAtPrimary    respjson.Field
		TranscodedQualities respjson.Field
		TranscodingSpeed    respjson.Field
		Uri                 respjson.Field
		VideoHeight         respjson.Field
		VideoWidth          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Stream) RawJSON() string { return r.JSON.raw }
func (r *Stream) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visualization mode for 360° streams, how the stream is rendered in our web
// player ONLY. If you would like to show video 360° in an external video player,
// then use parameters of that video player.
//
// Modes:
//
// - regular – regular “flat” stream
// - vr360 – display stream in 360° mode
// - vr180 – display stream in 180° mode
// - vr360tb – display stream in 3D 360° mode Top-Bottom
type StreamProjection string

const (
	StreamProjectionRegular StreamProjection = "regular"
	StreamProjectionVr360   StreamProjection = "vr360"
	StreamProjectionVr180   StreamProjection = "vr180"
	StreamProjectionVr360tb StreamProjection = "vr360tb"
)

// Method of recording a stream. Specifies the source from which the stream will be
// recorded: original or transcoded.
//
// Types:
//
//   - "origin" – To record RMTP/SRT/etc original clean media source.
//   - "transcoded" – To record the output transcoded version of the stream,
//     including overlays, texts, logos, etc. additional media layers.
type StreamRecordType string

const (
	StreamRecordTypeOrigin     StreamRecordType = "origin"
	StreamRecordTypeTranscoded StreamRecordType = "transcoded"
)

type StreamStartRecordingResponse struct {
	// Stream data
	Data StreamStartRecordingResponseData `json:"data"`
	// List of errors received on attempt to start recording process
	Errors []any `json:"errors"`
	// List of warnings received on starting recording process
	Warnings []StreamStartRecordingResponseWarning `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Errors      respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamStartRecordingResponse) RawJSON() string { return r.JSON.raw }
func (r *StreamStartRecordingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stream data
type StreamStartRecordingResponseData struct {
	Stream StreamStartRecordingResponseDataStream `json:"stream"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Stream      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamStartRecordingResponseData) RawJSON() string { return r.JSON.raw }
func (r *StreamStartRecordingResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamStartRecordingResponseDataStream struct {
	// Stream ID
	ID     int64                                        `json:"id"`
	Client StreamStartRecordingResponseDataStreamClient `json:"client"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Client      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamStartRecordingResponseDataStream) RawJSON() string { return r.JSON.raw }
func (r *StreamStartRecordingResponseDataStream) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamStartRecordingResponseDataStreamClient struct {
	// Client ID
	ID int64 `json:"id"`
	// Current storage limit for client by megabytes
	StorageLimitMB int64 `json:"storage_limit_mb"`
	// Current storage usage for client by megabytes
	StorageUsageMB float64 `json:"storage_usage_mb"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		StorageLimitMB respjson.Field
		StorageUsageMB respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamStartRecordingResponseDataStreamClient) RawJSON() string { return r.JSON.raw }
func (r *StreamStartRecordingResponseDataStreamClient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamStartRecordingResponseWarning struct {
	// current warning key
	//
	// Any of "client_storage_almost_exceeded".
	Key string `json:"key"`
	// storage usage state for client
	Meta StreamStartRecordingResponseWarningMeta `json:"meta"`
	// Warning source object
	SourceObject StreamStartRecordingResponseWarningSourceObject `json:"source_object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key          respjson.Field
		Meta         respjson.Field
		SourceObject respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamStartRecordingResponseWarning) RawJSON() string { return r.JSON.raw }
func (r *StreamStartRecordingResponseWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// storage usage state for client
type StreamStartRecordingResponseWarningMeta struct {
	// Current storage limit for client by megabytes
	StorageLimitMB int64 `json:"storage_limit_mb"`
	// Current storage usage for client by megabytes
	StorageUsageMB float64 `json:"storage_usage_mb"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StorageLimitMB respjson.Field
		StorageUsageMB respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamStartRecordingResponseWarningMeta) RawJSON() string { return r.JSON.raw }
func (r *StreamStartRecordingResponseWarningMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Warning source object
type StreamStartRecordingResponseWarningSourceObject struct {
	// Client ID
	ID int64 `json:"id"`
	// Object type (class)
	//
	// Any of "client".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamStartRecordingResponseWarningSourceObject) RawJSON() string { return r.JSON.raw }
func (r *StreamStartRecordingResponseWarningSourceObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamNewParams struct {
	// Stream name.
	//
	// Often used as a human-readable name for the stream, but can contain any text you
	// wish. The values are not unique and may be repeated.
	//
	// Examples:
	//
	// - Conference in July
	// - Stream #10003
	// - Open-Air Camera #31 Backstage
	// - 480fd499-2de2-4988-bc1a-a4eebe9818ee
	Name string `json:"name,required"`
	// Stream switch between on and off. This is not an indicator of the status "stream
	// is receiving and it is LIVE", but rather an on/off switch.
	//
	// When stream is switched off, there is no way to process it: PULL is deactivated
	// and PUSH will return an error.
	//
	// - true – stream can be processed
	// - false – stream is off, and cannot be processed
	Active param.Opt[bool] `json:"active,omitzero"`
	// Enables autotomatic recording of the stream when it started. So you don't need
	// to call recording manually.
	//
	// Result of recording is automatically added to video hosting. For details see the
	// /streams/`start_recording` method and in knowledge base
	//
	// Values:
	//
	// - true – auto recording is enabled
	// - false – auto recording is disabled
	AutoRecord param.Opt[bool] `json:"auto_record,omitzero"`
	// ID of custom CDN resource from which the content will be delivered (only if you
	// know what you do)
	CdnID param.Opt[int64] `json:"cdn_id,omitzero"`
	// Custom meta field designed to store your own extra information about a video
	// entity: video source, video id, parameters, etc. We do not use this field in any
	// way when processing the stream. You can store any data in any format (string,
	// json, etc), saved as a text string. Example:
	// `client_entity_data = '{ "seq_id": "1234567890", "name": "John Doe", "iat": 1516239022 }'`
	ClientEntityData param.Opt[string] `json:"client_entity_data,omitzero"`
	// Custom meta field for storing the Identifier in your system. We do not use this
	// field in any way when processing the stream. Example: `client_user_id = 1001`
	ClientUserID param.Opt[int64] `json:"client_user_id,omitzero"`
	// DVR duration in seconds if DVR feature is enabled for the stream. So this is
	// duration of how far the user can rewind the live stream.
	//
	// `dvr_duration` range is [30...14400].
	//
	// Maximum value is 4 hours = 14400 seconds. If you need more, ask the Support Team
	// please.
	DvrDuration param.Opt[int64] `json:"dvr_duration,omitzero"`
	// Enables DVR for the stream:
	//
	// - true – DVR is enabled
	// - false – DVR is disabled
	DvrEnabled param.Opt[bool] `json:"dvr_enabled,omitzero"`
	// Add `#EXT-X-ENDLIST` tag within .m3u8 playlist after the last segment of a live
	// stream when broadcast is ended.
	HlsMpegtsEndlistTag param.Opt[bool] `json:"hls_mpegts_endlist_tag,omitzero"`
	// Switch on mode to insert and display real-time HTML overlay widgets on top of
	// live streams
	HTMLOverlay param.Opt[bool] `json:"html_overlay,omitzero"`
	// Indicates if stream is pulled from external server or not. Has two possible
	// values:
	//
	//   - true – stream is received by PULL method. Use this when need to get stream
	//     from external server.
	//   - false – stream is received by PUSH method. Use this when need to send stream
	//     from end-device to our Streaming Platform, i.e. from your encoder, mobile app
	//     or OBS Studio.
	Pull param.Opt[bool] `json:"pull,omitzero"`
	// Custom quality set ID for transcoding, if transcoding is required according to
	// your conditions. Look at GET /`quality_sets` method
	QualitySetID param.Opt[int64] `json:"quality_set_id,omitzero"`
	// When using PULL method, this is the URL to pull a stream from.
	//
	// You can specify multiple addresses separated by a space (" "), so you can
	// organize a backup plan. In this case, the specified addresses will be selected
	// one by one using round robin scheduling. If the first address does not respond,
	// then the next one in the list will be automatically requested, returning to the
	// first and so on in a circle. Also, if the sucessfully working stream stops
	// sending data, then the next one will be selected according to the same scheme.
	//
	// After 2 hours of inactivity of your original stream, the system stops PULL
	// requests and the stream is deactivated (the "active" field switches to "false").
	//
	// Please, note that this field is for PULL only, so is not suitable for PUSH. Look
	// at fields "`push_url`" and "`push_url_srt`" from GET method.
	Uri param.Opt[string] `json:"uri,omitzero"`
	// IDs of broadcasts which will include this stream
	BroadcastIDs []int64 `json:"broadcast_ids,omitzero"`
	// Visualization mode for 360° streams, how the stream is rendered in our web
	// player ONLY. If you would like to show video 360° in an external video player,
	// then use parameters of that video player.
	//
	// Modes:
	//
	// - regular – regular “flat” stream
	// - vr360 – display stream in 360° mode
	// - vr180 – display stream in 180° mode
	// - vr360tb – display stream in 3D 360° mode Top-Bottom
	//
	// Any of "regular", "vr360", "vr180", "vr360tb".
	Projection StreamNewParamsProjection `json:"projection,omitzero"`
	// Method of recording a stream. Specifies the source from which the stream will be
	// recorded: original or transcoded.
	//
	// Types:
	//
	//   - "origin" – To record RMTP/SRT/etc original clean media source.
	//   - "transcoded" – To record the output transcoded version of the stream,
	//     including overlays, texts, logos, etc. additional media layers.
	//
	// Any of "origin", "transcoded".
	RecordType StreamNewParamsRecordType `json:"record_type,omitzero"`
	paramObj
}

func (r StreamNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StreamNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StreamNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visualization mode for 360° streams, how the stream is rendered in our web
// player ONLY. If you would like to show video 360° in an external video player,
// then use parameters of that video player.
//
// Modes:
//
// - regular – regular “flat” stream
// - vr360 – display stream in 360° mode
// - vr180 – display stream in 180° mode
// - vr360tb – display stream in 3D 360° mode Top-Bottom
type StreamNewParamsProjection string

const (
	StreamNewParamsProjectionRegular StreamNewParamsProjection = "regular"
	StreamNewParamsProjectionVr360   StreamNewParamsProjection = "vr360"
	StreamNewParamsProjectionVr180   StreamNewParamsProjection = "vr180"
	StreamNewParamsProjectionVr360tb StreamNewParamsProjection = "vr360tb"
)

// Method of recording a stream. Specifies the source from which the stream will be
// recorded: original or transcoded.
//
// Types:
//
//   - "origin" – To record RMTP/SRT/etc original clean media source.
//   - "transcoded" – To record the output transcoded version of the stream,
//     including overlays, texts, logos, etc. additional media layers.
type StreamNewParamsRecordType string

const (
	StreamNewParamsRecordTypeOrigin     StreamNewParamsRecordType = "origin"
	StreamNewParamsRecordTypeTranscoded StreamNewParamsRecordType = "transcoded"
)

type StreamUpdateParams struct {
	Stream StreamUpdateParamsStream `json:"stream,omitzero"`
	paramObj
}

func (r StreamUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow StreamUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StreamUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type StreamUpdateParamsStream struct {
	// Stream name.
	//
	// Often used as a human-readable name for the stream, but can contain any text you
	// wish. The values are not unique and may be repeated.
	//
	// Examples:
	//
	// - Conference in July
	// - Stream #10003
	// - Open-Air Camera #31 Backstage
	// - 480fd499-2de2-4988-bc1a-a4eebe9818ee
	Name string `json:"name,required"`
	// Stream switch between on and off. This is not an indicator of the status "stream
	// is receiving and it is LIVE", but rather an on/off switch.
	//
	// When stream is switched off, there is no way to process it: PULL is deactivated
	// and PUSH will return an error.
	//
	// - true – stream can be processed
	// - false – stream is off, and cannot be processed
	Active param.Opt[bool] `json:"active,omitzero"`
	// Enables autotomatic recording of the stream when it started. So you don't need
	// to call recording manually.
	//
	// Result of recording is automatically added to video hosting. For details see the
	// /streams/`start_recording` method and in knowledge base
	//
	// Values:
	//
	// - true – auto recording is enabled
	// - false – auto recording is disabled
	AutoRecord param.Opt[bool] `json:"auto_record,omitzero"`
	// ID of custom CDN resource from which the content will be delivered (only if you
	// know what you do)
	CdnID param.Opt[int64] `json:"cdn_id,omitzero"`
	// Custom meta field designed to store your own extra information about a video
	// entity: video source, video id, parameters, etc. We do not use this field in any
	// way when processing the stream. You can store any data in any format (string,
	// json, etc), saved as a text string. Example:
	// `client_entity_data = '{ "seq_id": "1234567890", "name": "John Doe", "iat": 1516239022 }'`
	ClientEntityData param.Opt[string] `json:"client_entity_data,omitzero"`
	// Custom meta field for storing the Identifier in your system. We do not use this
	// field in any way when processing the stream. Example: `client_user_id = 1001`
	ClientUserID param.Opt[int64] `json:"client_user_id,omitzero"`
	// DVR duration in seconds if DVR feature is enabled for the stream. So this is
	// duration of how far the user can rewind the live stream.
	//
	// `dvr_duration` range is [30...14400].
	//
	// Maximum value is 4 hours = 14400 seconds. If you need more, ask the Support Team
	// please.
	DvrDuration param.Opt[int64] `json:"dvr_duration,omitzero"`
	// Enables DVR for the stream:
	//
	// - true – DVR is enabled
	// - false – DVR is disabled
	DvrEnabled param.Opt[bool] `json:"dvr_enabled,omitzero"`
	// Add `#EXT-X-ENDLIST` tag within .m3u8 playlist after the last segment of a live
	// stream when broadcast is ended.
	HlsMpegtsEndlistTag param.Opt[bool] `json:"hls_mpegts_endlist_tag,omitzero"`
	// Switch on mode to insert and display real-time HTML overlay widgets on top of
	// live streams
	HTMLOverlay param.Opt[bool] `json:"html_overlay,omitzero"`
	// Indicates if stream is pulled from external server or not. Has two possible
	// values:
	//
	//   - true – stream is received by PULL method. Use this when need to get stream
	//     from external server.
	//   - false – stream is received by PUSH method. Use this when need to send stream
	//     from end-device to our Streaming Platform, i.e. from your encoder, mobile app
	//     or OBS Studio.
	Pull param.Opt[bool] `json:"pull,omitzero"`
	// Custom quality set ID for transcoding, if transcoding is required according to
	// your conditions. Look at GET /`quality_sets` method
	QualitySetID param.Opt[int64] `json:"quality_set_id,omitzero"`
	// When using PULL method, this is the URL to pull a stream from.
	//
	// You can specify multiple addresses separated by a space (" "), so you can
	// organize a backup plan. In this case, the specified addresses will be selected
	// one by one using round robin scheduling. If the first address does not respond,
	// then the next one in the list will be automatically requested, returning to the
	// first and so on in a circle. Also, if the sucessfully working stream stops
	// sending data, then the next one will be selected according to the same scheme.
	//
	// After 2 hours of inactivity of your original stream, the system stops PULL
	// requests and the stream is deactivated (the "active" field switches to "false").
	//
	// Please, note that this field is for PULL only, so is not suitable for PUSH. Look
	// at fields "`push_url`" and "`push_url_srt`" from GET method.
	Uri param.Opt[string] `json:"uri,omitzero"`
	// IDs of broadcasts which will include this stream
	BroadcastIDs []int64 `json:"broadcast_ids,omitzero"`
	// Visualization mode for 360° streams, how the stream is rendered in our web
	// player ONLY. If you would like to show video 360° in an external video player,
	// then use parameters of that video player.
	//
	// Modes:
	//
	// - regular – regular “flat” stream
	// - vr360 – display stream in 360° mode
	// - vr180 – display stream in 180° mode
	// - vr360tb – display stream in 3D 360° mode Top-Bottom
	//
	// Any of "regular", "vr360", "vr180", "vr360tb".
	Projection string `json:"projection,omitzero"`
	// Method of recording a stream. Specifies the source from which the stream will be
	// recorded: original or transcoded.
	//
	// Types:
	//
	//   - "origin" – To record RMTP/SRT/etc original clean media source.
	//   - "transcoded" – To record the output transcoded version of the stream,
	//     including overlays, texts, logos, etc. additional media layers.
	//
	// Any of "origin", "transcoded".
	RecordType string `json:"record_type,omitzero"`
	paramObj
}

func (r StreamUpdateParamsStream) MarshalJSON() (data []byte, err error) {
	type shadow StreamUpdateParamsStream
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StreamUpdateParamsStream) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[StreamUpdateParamsStream](
		"projection", "regular", "vr360", "vr180", "vr360tb",
	)
	apijson.RegisterFieldValidator[StreamUpdateParamsStream](
		"record_type", "origin", "transcoded",
	)
}

type StreamListParams struct {
	// Query parameter. Use it to list the paginated content
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Query parameter. Set to 1 to get details of the broadcasts associated with the
	// stream
	WithBroadcasts param.Opt[int64] `query:"with_broadcasts,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StreamListParams]'s query parameters as `url.Values`.
func (r StreamListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StreamNewClipParams struct {
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
	// Therefore, to change the maximum, use "`dvr_duration`" parameter of this stream.
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

func (r StreamNewClipParams) MarshalJSON() (data []byte, err error) {
	type shadow StreamNewClipParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StreamNewClipParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
