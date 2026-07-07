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
	shimjson "github.com/G-Core/gcore-go/internal/encoding/json"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// VideoService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVideoService] method instead.
type VideoService struct {
	Options   []option.RequestOption
	Subtitles VideoSubtitleService
}

// NewVideoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVideoService(opts ...option.RequestOption) (r VideoService) {
	r = VideoService{}
	r.Options = opts
	r.Subtitles = NewVideoSubtitleService(opts...)
	return
}

// Use this method to create a new video entity.
//
// **Methods of creating**
//
// To upload the original video file to the server, there are several possible
// scenarios:
//
//   - **Copy from another server** – If your video is accessible via "http://",
//     "https://", or "sftp://" public link, then you can use this method to copy a
//     file from an external server. Set `origin_url` parameter with the link to the
//     original video file (i.e. "https://domain.com/video.mp4"). After method
//     execution file will be uploaded and will be sent to transcoding automatically,
//     you don't have to do anything else. Use extra field `origin_http_headers` if
//     authorization is required on the external server.
//   - **Direct upload from a local device** – If you need to upload video directly
//     from your local device or from a mobile app, then use this method. Keep
//     `origin_url` empty and use TUS protocol ([tus.io](https://tus.io)) to upload
//     file. More details are here
//     ["Get TUS' upload"](/api-reference/streaming/videos/get-tus-parameters-for-direct-upload)
//
// After getting the video, it is processed through the queue. There are 2 priority
// criteria: global and local. Global is determined automatically by the system as
// converters are ready to get next video, so your videos rarely queue longer than
// usual (when you don't have a dedicated region). Local priority works at the
// level of your account and you have full control over it, look at "priority"
// attribute.
//
// **AI processing**
//
// When uploading a video, it is possible to automatically create subtitles based
// on AI.
//
// Read more:
//
//   - What is
//     ["AI Speech Recognition"](/api-reference/streaming/ai/create-ai-asr-task).
//   - If the option is enabled via
//     `auto_transcribe_audio_language: auto|<language_code>`, then immediately after
//     successful transcoding, an AI task will be automatically created for
//     transcription.
//   - If you need to translate subtitles from original language to any other, then
//     AI-task of subtitles translation can be applied. Use
//     `auto_translate_subtitles_language: default|<language_codes,>` parameter for
//     that. Also you can point several languages to translate to, then a separate
//     subtitle will be generated for each specified language.
//   - How to
//     ["add AI-generated subtitles to an exist video"](/api-reference/streaming/subtitles/add-subtitle).
//
// The created AI-task(s) will be automatically executed, and result will also be
// automatically attached to this video as subtitle(s).
//
// Please note that transcription is done automatically for all videos uploaded to
// our video hosting. If necessary, you can disable automatic creation of
// subtitles. If AI is disabled in your account, no AI functionality is called.
//
// **Advanced Features** For details on the requirements for incoming original
// files, and output video parameters after transcoding, refer to the Knowledge
// Base documentation. By default, video will be transcoded into H.264 (AVC)
// according to the original resolution (up to 4K), and a suitable quality ladder
// will be applied. More advanced codecs such as HEVC, AV1, and VP9 are available
// as part of our premium encoding features. There is no automatic upscaling; the
// maximum quality is taken from the original video. If you want to upload specific
// files not explicitly listed in requirements or wish to modify the standard
// quality ladder (i.e. decrease quality or add new non-standard qualities), then
// such customization is possible. Please reach out to us for assistance.
//
// Additionally, check the Knowledge Base for any supplementary information you may
// need.
func (r *VideoService) New(ctx context.Context, body VideoNewParams, opts ...option.RequestOption) (res *[]Video, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "streaming/videos"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Changes parameters of the video to new values.
//
// It's allowed to update only those public parameters that are described in POST
// method to create a new “video” entity. So it's not possible to change calculated
// parameters like "id", "duration", "hls_url", etc.
//
// Examples of changing:
//
// - Name: `{ "name": "new name of the video" }`
// - Move the video to a new directory: ` { "directory_id": 200 }`
//
// Please note that some parameters are used on initial step (before transcoding)
// only, so after transcoding there is no use in changing their values. For
// example, "origin_url" parameter is used for downloading an original file from a
// source and never used after transcoding; or "priority" parameter is used to set
// priority of processing and never used after transcoding.
func (r *VideoService) Update(ctx context.Context, videoID int64, body VideoUpdateParams, opts ...option.RequestOption) (res *Video, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/videos/%v", videoID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a set of videos by the given criteria.
func (r *VideoService) List(ctx context.Context, query VideoListParams, opts ...option.RequestOption) (res *pagination.PageStreaming[VideoListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "streaming/videos"
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

// Returns a set of videos by the given criteria.
func (r *VideoService) ListAutoPaging(ctx context.Context, query VideoListParams, opts ...option.RequestOption) *pagination.PageStreamingAutoPager[VideoListResponse] {
	return pagination.NewPageStreamingAutoPager(r.List(ctx, query, opts...))
}

// Operation to delete video entity.
//
// When you delete a video, all transcoded qualities and all associated files such
// as subtitles and screenshots, as well as other data, are deleted from cloud
// storage.
//
// The video is deleted permanently and irreversibly. Therefore, it is impossible
// to restore files after this.
//
// For detailed information and information on calculating your maximum monthly
// storage usage, please refer to the Product Documentation.
func (r *VideoService) Delete(ctx context.Context, videoID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("streaming/videos/%v", videoID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Mass upload of your videos. Method is used to set the task of creating videos in
// the form of 1 aggregated request instead of a large number of single requests.
//
// An additional advantage is the ability to specify subtitles in the same request.
// Whereas for a normal single upload, subtitles are uploaded in separate requests.
//
// All videos in the request will be processed in queue in order of priority. Use
// "priority" attribute and look at general description in POST /videos method.
//
// Limits:
//
// - Batch max size = 500 videos.
// - Max body size (payload) = 64MB.
// - API connection timeout = 30 sec.
func (r *VideoService) NewMultiple(ctx context.Context, params VideoNewMultipleParams, opts ...option.RequestOption) (res *[]Video, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "streaming/videos/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Information about a video entity.
//
// Contains all the data about the video: meta-data, data for streaming and
// renditions, static media data, data about original video.
//
// You can use different methods to play video:
//
//   - `iframe_url` – a URL to a built-in HTML video player with automatically
//     configured video playback.
//   - `hls_url` – a URLs to HLS TS .m3u8 manifest, which can be played in video
//     players.
//   - `hls_cmaf_url` – a URL to HLS CMAF .m3u8 manifest with chunks in fMP4 format,
//     which can be played in most modern video players.
//   - `dash_url` – a URL to MPEG-DASH .mpd manifest, which can be played in most
//     modern video players. Preferable for Android and Windows devices.
//   - `converted_videos`/`mp4_url` – a URL to MP4 file of specific rendition.
//
// ![Video player](https://demo-files.gvideo.io/apidocs/coffee-run-player.jpg)
func (r *VideoService) Get(ctx context.Context, videoID int64, opts ...option.RequestOption) (res *Video, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/videos/%v", videoID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Use this method to get TUS' session parameters: hostname of the server to
// upload, secure token.
//
// The general sequence of actions for a direct upload of a video is as follows:
//
//   - Create video entity via POST method
//     ["Create video"](/api-reference/streaming/videos/create-video)
//   - Get TUS' session parameters (you are here now)
//   - Upload file via TUS client, choose your implementation on
//     [tus.io](https://tus.io/implementations)
//
// Final endpoint for uploading is constructed using the following template:
// "https://{hostname}/upload/". Also you have to provide token, `client_id`,
// `video_id` as metadata too.
//
// A short javascript example is shown below, based on tus-js-client. Variable
// "data" below is the result of this API request. Please, note that we support 2.x
// version only of tus-js-client.
//
// ```
//
//	uploads[data.video.id] = new tus.Upload(file, {
//	  endpoint: `https://${data.servers[0].hostname}/upload/`,
//	  metadata: {
//	    filename: data.video.name,
//	    token: data.token,
//	    video_id: data.video.id,
//	    client_id: data.video.client_id
//	  },
//	  onSuccess: function() {
//	    ...
//	  }
//	}
//	uploads[data.video.id].start();
//
// ```
func (r *VideoService) GetParametersForDirectUpload(ctx context.Context, videoID int64, opts ...option.RequestOption) (res *DirectUploadParametersResp, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/videos/%v/upload", videoID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns names for specified video IDs
func (r *VideoService) ListNames(ctx context.Context, query VideoListNamesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "streaming/videos/names"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

type DirectUploadParametersResp struct {
	// Token
	Token string `json:"token"`
	// An array which contains information about servers you can upload a video to.
	//
	//	**Server;** type — object.
	//
	// ---
	//
	// Server has the following fields:
	//
	//   - **id;** type — integer
	//     Server ID
	//   - **hostname;** type — string
	//     Server hostname
	Servers []any `json:"servers"`
	// Contains information about the created video. See the full description in the
	// Get video request
	Video any `json:"video"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Servers     respjson.Field
		Video       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectUploadParametersResp) RawJSON() string { return r.JSON.raw }
func (r *DirectUploadParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Subtitle struct {
	// ID of subtitle file
	ID int64 `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	SubtitleBase
}

// Returns the unmodified JSON received from the API
func (r Subtitle) RawJSON() string { return r.JSON.raw }
func (r *Subtitle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubtitleBase struct {
	// 3-letter language code according to ISO-639-2 (bibliographic code)
	Language string `json:"language"`
	// Name of subtitle file
	Name string `json:"name"`
	// Full text of subtitles/captions, with escaped "\n" ("\r") symbol of new line
	Vtt string `json:"vtt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Name        respjson.Field
		Vtt         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubtitleBase) RawJSON() string { return r.JSON.raw }
func (r *SubtitleBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SubtitleBase to a SubtitleBaseParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SubtitleBaseParam.Overrides()
func (r SubtitleBase) ToParam() SubtitleBaseParam {
	return param.Override[SubtitleBaseParam](json.RawMessage(r.RawJSON()))
}

type SubtitleBaseParam struct {
	// 3-letter language code according to ISO-639-2 (bibliographic code)
	Language param.Opt[string] `json:"language,omitzero"`
	// Name of subtitle file
	Name param.Opt[string] `json:"name,omitzero"`
	// Full text of subtitles/captions, with escaped "\n" ("\r") symbol of new line
	Vtt param.Opt[string] `json:"vtt,omitzero"`
	paramObj
}

func (r SubtitleBaseParam) MarshalJSON() (data []byte, err error) {
	type shadow SubtitleBaseParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubtitleBaseParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoListResponse struct {
	// Video ID
	ID int64 `json:"id"`
	// ID of ad that should be shown. If empty the default ad is show. If there is no
	// default ad, no ad is shown.
	AdID int64 `json:"ad_id"`
	// Total number of video views. It is calculated based on the analysis of all
	// views, no matter in which player.
	CDNViews int64 `json:"cdn_views"`
	// Client ID
	ClientID int64 `json:"client_id"`
	// Custom meta field for storing the Identifier in your system. We do not use this
	// field in any way when processing the stream. Example: `client_user_id = 1001`
	ClientUserID int64 `json:"client_user_id"`
	// Array of data about each transcoded quality
	ConvertedVideos []VideoListResponseConvertedVideo `json:"converted_videos"`
	// Time of creation. Datetime in ISO 8601 format.
	CreatedAt string `json:"created_at"`
	// Custom URL of Iframe for video player to be used in share panel in player. Auto
	// generated Iframe URL provided by default.
	CustomIframeURL string `json:"custom_iframe_url"`
	// A URL to a master playlist MPEG-DASH (master.mpd) with CMAF or WebM based
	// chunks.
	//
	// Chunk type will be selected automatically for each quality:
	//
	// - CMAF for H264 and H265 codecs.
	// - WebM for AV1 codec.
	//
	// This URL is a link to the main manifest. But you can also manually specify
	// suffix-options that will allow you to change the manifest to your request:
	//
	// ```
	// /videos/{client_id}_{slug}/master[-min-N][-max-N][-(h264|hevc|av1)].mpd
	// ```
	//
	// List of suffix-options:
	//
	//   - [-min-N] – ABR soft limitation of qualities from below.
	//   - [-max-N] – ABR soft limitation of qualities from above.
	//   - [-(h264|hevc|av1) – Video codec soft limitation. Applicable if the video was
	//     transcoded into multiple codecs H264, H265 and AV1 at once, but you want to
	//     return just 1 video codec in a manifest. Read the Product Documentation for
	//     details.
	//
	// Read more what is ABR soft-limiting in the "hls_url" field above.
	//
	// Caution. Solely master.mpd is officially documented and intended for your use.
	// Any additional internal manifests, sub-manifests, parameters, chunk names, file
	// extensions, and related components are internal infrastructure entities. These
	// may undergo modifications without prior notice, in any manner or form. It is
	// strongly advised not to store them in your database or cache them on your end.
	DashURL string `json:"dash_url"`
	// Additional text field for video description
	Description string `json:"description"`
	// ID of the directory where the video is stored.
	DirectoryID int64 `json:"directory_id"`
	// Video duration in milliseconds. May differ from "origin_video_duration" value if
	// the video was uploaded with clipping through the parameters "clip_start_seconds"
	// and "clip_duration_seconds"
	Duration int64 `json:"duration"`
	// Video processing error text will be saved here if "status: error"
	Error string `json:"error"`
	// A URL to a master playlist HLS (master-cmaf.m3u8) with CMAF-based chunks. Chunks
	// are in fMP4 container. It's a code-agnostic container, which allows to use any
	// like H264, H265, AV1, etc.
	//
	// It is possible to use the same suffix-options as described in the "hls_url"
	// attribute.
	//
	// Caution. Solely master.m3u8 (and master[-options].m3u8) is officially documented
	// and intended for your use. Any additional internal manifests, sub-manifests,
	// parameters, chunk names, file extensions, and related components are internal
	// infrastructure entities. These may undergo modifications without prior notice,
	// in any manner or form. It is strongly advised not to store them in your database
	// or cache them on your end.
	HlsCmafURL string `json:"hls_cmaf_url"`
	// A URL to a master playlist HLS (master.m3u8). Chunk type will be selected
	// automatically:
	//
	//   - TS if your video was encoded to H264 only.
	//   - CMAF if your video was encoded additionally to H265 and/or AV1 codecs (as
	//     Apple does not support these codecs over MPEG TS, and they are not
	//     standardized in TS-container).
	//
	// You can also manually specify suffix-options that will allow you to change the
	// manifest to your request:
	//
	// ```
	// /videos/{client_id}_{video_slug}/master[-cmaf][-min-N][-max-N][-img][-(h264|hevc|av1)].m3u8
	// ```
	//
	// List of suffix-options:
	//
	//   - [-cmaf] – getting HLS CMAF version of the manifest. Look at the `hls_cmaf_url`
	//     field.
	//   - [-min-N] – ABR soft limitation of qualities from below.
	//   - [-max-N] – ABR soft limitation of qualities from above.
	//   - [-img] – Roku trick play: to add tiles directly into .m3u8 manifest. Read the
	//     Product Documentation for details.
	//   - [-(h264|hevc|av1) – Video codec soft limitation. Applicable if the video was
	//     transcoded into multiple codecs H264, H265 and AV1 at once, but you want to
	//     return just 1 video codec in a manifest. Read the Product Documentation for
	//     details.
	//
	// ABR soft-limiting: Soft limitation of the list of qualities allows you to return
	// not the entire list of transcoded qualities for a video, but only those you
	// need. For example, the video is available in 7 qualities from 360p to 4K, but
	// you want to return not more than 480p only due to the conditions of distribution
	// of content to a specific end-user (i.e. free account): ABR soft-limiting
	// examples:
	//
	//   - To a generic `.../master.m3u8` manifest
	//   - Add a suffix-option to limit quality `.../master-max-480.m3u8`
	//   - Add a suffix-option to limit quality and codec
	//     `.../master-min-320-max-320-h264.m3u8` For more details look at the Product
	//     Documentation.
	//
	// Caution. Solely master.m3u8 (and master[-options].m3u8) is officially documented
	// and intended for your use. Any additional internal manifests, sub-manifests,
	// parameters, chunk names, file extensions, and related components are internal
	// infrastructure entities. These may undergo modifications without prior notice,
	// in any manner or form. It is strongly advised not to store them in your database
	// or cache them on your end.
	HlsURL string `json:"hls_url"`
	// Title of the video.
	//
	// Often used as a human-readable name of the video, but can contain any text you
	// wish. The values are not unique and may be repeated. Examples:
	//
	// - Educational training 2024-03-29
	// - Series X S3E14, The empire strikes back
	// - 480fd499-2de2-4988-bc1a-a4eebe9818ee
	Name string `json:"name"`
	// Size of original file
	OriginSize int64 `json:"origin_size"`
	// URL to an original file from which the information for transcoding was taken.
	//
	// May contain a link for scenarios:
	//
	// - If the video was downloaded from another origin
	// - If the video is a recording of a live stream
	// - Otherwise it is "null"
	//
	// **Copy from another server** URL to an original file that was downloaded. Look
	// at method "Copy from another server" in POST /videos. **Recording of an original
	// live stream**
	//
	// URL to the original non-transcoded stream recording with original quality, saved
	// in MP4 format. File is created immediately after the completion of the stream
	// recording. The stream from which the recording was made is reflected in
	// "stream_id" field.
	//
	// Can be used for internal operations when a recording needs to be received faster
	// than the transcoded versions are ready. But this version is not intended for
	// public distribution. Views and downloads occur in the usual way, like viewing an
	// MP4 rendition.
	//
	// The MP4 file becomes available for downloading when the video entity "status"
	// changes from "new" to "pending".
	//
	// Format of URL is `/videos/<cid>_<slug>/origin_<bitrate>_<height>.mp4` Where:
	//
	// - `<bitrate>` – Encoding bitrate in Kbps.
	// - `<height>` – Video height.
	//
	// The original file is stored for up to 7 days, after which it is deleted
	// automatically. By default, the retention policy for original recorded files
	// cannot be changed. For enterprise customers, it can be adjusted individually
	// upon request.
	//
	// This is a premium feature, available only upon request through your manager or
	// support team.
	OriginURL string `json:"origin_url"`
	// Original video duration in milliseconds
	OriginVideoDuration int64 `json:"origin_video_duration"`
	// Poster is your own static image which can be displayed before the video begins
	// playing. This is often a frame of the video or a custom title screen.
	//
	// Field contains a link to your own uploaded image.
	//
	// Also look at "screenshot" attribute.
	Poster string `json:"poster"`
	// Field contains a link to minimized poster image. Original "poster" image is
	// proportionally scaled to a size of 200 pixels in height.
	PosterThumb string `json:"poster_thumb"`
	// This value indicates the priority with which the transcoding task was assigned.
	// After transcoding is complete, this value no longer has any effect and simply
	// displays historical data.
	Priority int64 `json:"priority"`
	// Regulates the video format:
	//
	// - **regular** — plays the video as usual
	// - **vr360** — plays the video in 360 degree mode
	// - **vr180** — plays the video in 180 degree mode
	// - **vr360tb** — plays the video in 3D 360 degree mode Top-Bottom.
	//
	// Default is regular
	Projection string `json:"projection"`
	// Custom quality set ID used for transcoding.
	QualitySetID int64 `json:"quality_set_id"`
	// Method of recording a stream. Specifies the source from which the stream will be
	// recorded: original or transcoded.
	//
	// Types:
	//
	//   - null – indicates that the video was uploaded in the standard way (and not
	//     recorded from a stream).
	//   - "origin" – To record RMTP/SRT/etc original clean media source.
	//   - "transcoded" – To record the output transcoded version of the stream,
	//     including overlays, texts, logos, etc. additional media layers.
	//
	// Any of "origin", "transcoded".
	RecordType VideoListResponseRecordType `json:"record_type" api:"nullable"`
	// If the video was saved from a stream, then start time of the stream recording is
	// saved here. Format is date time in ISO 8601
	RecordingStartedAt string `json:"recording_started_at"`
	// A URL to the default screenshot is here. The image is selected from an array of
	// all screenshots based on the “`screenshot_id`” attribute. If you use your own
	// "poster", the link to it will be here too.
	//
	// Our video player uses this field to display the static image before the video
	// starts playing. As soon as the user hits "play" the image will go away. If you
	// use your own external video player, then you can use the value of this field to
	// set the poster/thumbnail in your player.
	//
	// Example:
	//
	// - `video_js`.poster: `api.screenshot`
	// - clappr.poster: `api.screenshot`
	Screenshot string `json:"screenshot"`
	// ID of auto generated screenshots to be used for default screenshot.
	//
	// Counting from 0. A value of -1 sets the "screenshot" attribute to the URL of
	// your own image from the "poster" attribute.
	ScreenshotID int64 `json:"screenshot_id"`
	// Array of auto generated screenshots from the video. By default 5 static
	// screenshots are taken from different places in the video. If the video is short,
	// there may be fewer screenshots.
	//
	// Screenshots are created automatically, so they may contain not very good frames
	// from the video. To use your own image look at "poster" attribute.
	Screenshots []string `json:"screenshots"`
	// Custom URL or iframe displayed in the link field when a user clicks on a sharing
	// button in player. If empty, the link field and social network sharing is
	// disabled
	ShareURL string `json:"share_url"`
	// A unique alphanumeric identifier used in public URLs to retrieve and view the
	// video. It is unique for each video, generated randomly and set automatically by
	// the system.
	//
	// Format of usage in URL is _.../videos/{`client_id`}\_{slug}/..._
	//
	// Example:
	//
	// - Player: /videos/`12345_neAq1bYZ2`
	// - Manifest: /videos/`12345_neAq1bYZ2`/master.m3u8
	// - Rendition: /videos/`12345_neAq1bYZ2`/`qid90v1_720`.mp4
	Slug string `json:"slug"`
	// Video processing status:
	//
	//   - empty – initial status, when video-entity is created, but video-file has not
	//     yet been fully uploaded (TUS uploading, or downloading from an origin is not
	//     finished yet)
	//   - pending – video is in queue to be processed
	//   - viewable – video has at least 1 quality and can already be viewed via a link,
	//     but not all qualities are ready yet
	//   - ready – video is completely ready, available for viewing with all qualities
	//   - error – error while processing a video, look at "error" field
	//
	// Any of "empty", "pending", "viewable", "ready", "error".
	Status VideoListResponseStatus `json:"status"`
	// If the video was saved from a stream, then ID of that stream is saved here
	StreamID int64 `json:"stream_id"`
	// Time of last update of the video entity. Datetime in ISO 8601 format.
	UpdatedAt string `json:"updated_at"`
	// Number of video views through the built-in HTML video player of the Streaming
	// Platform only. This attribute does not count views from other external players
	// and native OS players, so here may be less number of views than in "cdn_views".
	Views int64 `json:"views"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AdID                respjson.Field
		CDNViews            respjson.Field
		ClientID            respjson.Field
		ClientUserID        respjson.Field
		ConvertedVideos     respjson.Field
		CreatedAt           respjson.Field
		CustomIframeURL     respjson.Field
		DashURL             respjson.Field
		Description         respjson.Field
		DirectoryID         respjson.Field
		Duration            respjson.Field
		Error               respjson.Field
		HlsCmafURL          respjson.Field
		HlsURL              respjson.Field
		Name                respjson.Field
		OriginSize          respjson.Field
		OriginURL           respjson.Field
		OriginVideoDuration respjson.Field
		Poster              respjson.Field
		PosterThumb         respjson.Field
		Priority            respjson.Field
		Projection          respjson.Field
		QualitySetID        respjson.Field
		RecordType          respjson.Field
		RecordingStartedAt  respjson.Field
		Screenshot          respjson.Field
		ScreenshotID        respjson.Field
		Screenshots         respjson.Field
		ShareURL            respjson.Field
		Slug                respjson.Field
		Status              respjson.Field
		StreamID            respjson.Field
		UpdatedAt           respjson.Field
		Views               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoListResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoListResponseConvertedVideo struct {
	// ID of the converted file of the specific quality
	ID int64 `json:"id"`
	// Video processing error text in this quality
	Error string `json:"error"`
	// Height in pixels of the converted video file of the specific quality. Can be
	// `null` for audio-only files.
	Height int64 `json:"height"`
	// A URL to a rendition file of the specified quality in MP4 format for
	// downloading.
	//
	// **Download methods**
	//
	// For each converted video, additional download endpoints are available under
	// `converted_videos`/`mp4_urls`. An MP4 download endpoints:
	//
	// 1. `/videos/{client_id}_{slug}/{filename}.mp4`
	// 2. `/videos/{client_id}_{slug}/{filename}.mp4/download`
	// 3. `/videos/{client_id}_{slug}/{filename}.mp4/download={custom_filename}`
	//
	// The first option returns the file as is. Response will be:
	//
	// ```
	// GET .mp4
	// ...
	// content-type: video/mp4
	// ```
	//
	// The second option with `/download` will respond with HTTP response header that
	// directly tells browsers to download the file instead of playing it in the
	// browser:
	//
	// ```
	// GET .mp4/download
	// ...
	// content-type: video/mp4
	// content-disposition: attachment
	// access-control-expose-headers: Content-Disposition
	// ```
	//
	// The third option allows you to set a custom name for the file being downloaded.
	// You can optionally specify a custom filename (just name excluding the .mp4
	// extension) using the download= query.
	//
	// Filename constraints:
	//
	// - Length: 1-255 characters
	// - Must NOT include the .mp4 extension (it is added automatically)
	// - Allowed characters: a-z, A-Z, 0-9, \_(underscore), -(dash), .(dot)
	// - First character cannot be .(dot)
	// - Example valid filenames: `holiday2025`, `_backup.final`, `clip-v1.2`
	//
	// ```
	// GET .mp4/download={custom_filename}
	// ...
	// content-type: video/mp4
	// content-disposition: attachment; filename="{custom_filename}.mp4"
	// access-control-expose-headers: Content-Disposition
	// ```
	//
	// Examples:
	//
	//   - MP4:
	//     `https://demo-public.gvideo.io/videos/2675_1OFgHZ1FWZNNvx1A/qid3567v1_h264_4050_1080.mp4/download`
	//   - MP4 with custom download filename:
	//     `https://demo-public.gvideo.io/videos/2675_1OFgHZ1FWZNNvx1A/qid3567v1_h264_4050_1080.mp4/download=highlights_v1.1_2025-05-30`
	//
	// **Default MP4 file name structure**
	//
	// Link to the file {filename} contains information about the encoding method using
	// format:
	//
	// `<quality_version>_<codec>_<bitrate>_<height>.mp4`
	//
	//   - `<quality_version>` – Internal quality identifier and file version. Please do
	//     not use it, can be changed at any time without any notice.
	//   - `<codec>` – Codec name that was used to encode the video, or audio codec if it
	//     is an audio-only file.
	//   - `<bitrate>` – Encoding bitrate in Kbps.
	//   - `<height>` – Video height, or word "audio" if it is an audio-only file.
	//
	// Note that this link format has been applied since 14.08.2024. If the video
	// entity was uploaded earlier, links may have old simplified format.
	//
	// Example: `/videos/{client_id}_{slug}/qid3567v1_h264_4050_1080.mp4`
	//
	// **Dynamic speed limiting** This mode sets different limits for different users
	// or for different types of content. The speed is adjusted based on requests with
	// the “speed” and “buffer” arguments.
	//
	// Example: `?speed=50k&buffer=500k`
	//
	// Read more in Product Documentation in CDN section "Network limits".
	//
	// **Secure token authentication for MP4 (updated)**
	//
	// Access to MP4 download links only can be protected using advanced secure tokens
	// passed as query parameters.
	//
	// Token generation uses the entire MP4 path, which ensures the token only grants
	// access to a specific quality/version of the video. This prevents unintended
	// access to other bitrate versions of an ABR stream.
	//
	// Token Query Parameters:
	//
	// - token: The generated hash
	// - expires: Expiration timestamp
	// - speed: (optional) Speed limit in bytes/sec, or empty string
	// - buffer: (optional) Buffer size in bytes, or empty string
	//
	// Optional (for IP-bound tokens):
	//
	//   - ip: The user’s IP address Example:
	//     `?md5=QX39c77lbQKvYgMMAvpyMQ&expires=1743167062`
	//
	// Read more in Product Documentation in Streaming section "Protected temporarily
	// link".
	MP4URL string `json:"mp4_url"`
	// Specific quality name
	Name string `json:"name"`
	// Status of transcoding into the specific quality, from 0 to 100
	Progress int64 `json:"progress"`
	// Size in bytes of the converted file of the specific quality. Can be `null` until
	// transcoding is fully completed.
	Size int64 `json:"size"`
	// Status of transcoding:
	//
	// - processing – video is being transcoded to this quality,
	// - complete – quality is fully processed,
	// - error – quality processing error, see parameter "error".
	//
	// Any of "processing", "complete", "error".
	Status string `json:"status"`
	// Width in pixels of the converted video file of the specified quality. Can be
	// `null` for audio files.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Error       respjson.Field
		Height      respjson.Field
		MP4URL      respjson.Field
		Name        respjson.Field
		Progress    respjson.Field
		Size        respjson.Field
		Status      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoListResponseConvertedVideo) RawJSON() string { return r.JSON.raw }
func (r *VideoListResponseConvertedVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Method of recording a stream. Specifies the source from which the stream will be
// recorded: original or transcoded.
//
// Types:
//
//   - null – indicates that the video was uploaded in the standard way (and not
//     recorded from a stream).
//   - "origin" – To record RMTP/SRT/etc original clean media source.
//   - "transcoded" – To record the output transcoded version of the stream,
//     including overlays, texts, logos, etc. additional media layers.
type VideoListResponseRecordType string

const (
	VideoListResponseRecordTypeOrigin     VideoListResponseRecordType = "origin"
	VideoListResponseRecordTypeTranscoded VideoListResponseRecordType = "transcoded"
)

// Video processing status:
//
//   - empty – initial status, when video-entity is created, but video-file has not
//     yet been fully uploaded (TUS uploading, or downloading from an origin is not
//     finished yet)
//   - pending – video is in queue to be processed
//   - viewable – video has at least 1 quality and can already be viewed via a link,
//     but not all qualities are ready yet
//   - ready – video is completely ready, available for viewing with all qualities
//   - error – error while processing a video, look at "error" field
type VideoListResponseStatus string

const (
	VideoListResponseStatusEmpty    VideoListResponseStatus = "empty"
	VideoListResponseStatusPending  VideoListResponseStatus = "pending"
	VideoListResponseStatusViewable VideoListResponseStatus = "viewable"
	VideoListResponseStatusReady    VideoListResponseStatus = "ready"
	VideoListResponseStatusError    VideoListResponseStatus = "error"
)

type VideoNewParams struct {
	Video CreateVideoParam `json:"video,omitzero"`
	paramObj
}

func (r VideoNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VideoNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoUpdateParams struct {
	CreateVideo CreateVideoParam
	paramObj
}

func (r VideoUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateVideo)
}
func (r *VideoUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoListParams struct {
	// IDs of the videos to find. You can specify one or more identifiers separated by
	// commas. Example, ?id=1,101,1001
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Find videos where "client_user_id" meta field is equal to the search value
	ClientUserID param.Opt[int64] `query:"client_user_id,omitzero" json:"-"`
	// Restriction to return only the specified attributes, instead of the entire
	// dataset. Specify, if you need to get short response. The following fields are
	// available for specifying: id, name, duration, status, `created_at`,
	// `updated_at`, `hls_url`, screenshots, `converted_videos`, priority, `stream_id`.
	// Example, ?fields=id,name,`hls_url`
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// Page number. Use it to list the paginated content
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Items per page number. Use it to list the paginated content
	PerPage param.Opt[int64] `query:"per_page,omitzero" json:"-"`
	// Aggregated search condition. If set, the video list is filtered by one combined
	// SQL criterion:
	//
	// - id={s} OR slug={s} OR name like {s}
	//
	// i.e. "/videos?search=1000" returns list of videos where id=1000 or slug=1000 or
	// name contains "1000".
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Use it to get videos filtered by their status. Possible values:
	//
	// - empty
	// - pending
	// - viewable
	// - ready
	// - error
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Find videos recorded from a specific stream, so for which "stream_id" field is
	// equal to the search value
	StreamID param.Opt[int64] `query:"stream_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VideoListParams]'s query parameters as `url.Values`.
func (r VideoListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type VideoNewMultipleParams struct {
	// Restriction to return only the specified attributes, instead of the entire
	// dataset. Specify, if you need to get short response. The following fields are
	// available for specifying: id, name, duration, status, `created_at`,
	// `updated_at`, `hls_url`, screenshots, `converted_videos`, priority. Example,
	// ?fields=id,name,`hls_url`
	Fields param.Opt[string]             `query:"fields,omitzero" json:"-"`
	Videos []VideoNewMultipleParamsVideo `json:"videos,omitzero"`
	paramObj
}

func (r VideoNewMultipleParams) MarshalJSON() (data []byte, err error) {
	type shadow VideoNewMultipleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoNewMultipleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [VideoNewMultipleParams]'s query parameters as `url.Values`.
func (r VideoNewMultipleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type VideoNewMultipleParamsVideo struct {
	Subtitles []SubtitleBaseParam `json:"subtitles,omitzero"`
	CreateVideoParam
}

func (r VideoNewMultipleParamsVideo) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*VideoNewMultipleParamsVideo
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type VideoListNamesParams struct {
	// Comma-separated set of video IDs. Example, ?ids=7,17
	IDs []int64 `query:"ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VideoListNamesParams]'s query parameters as `url.Values`.
func (r VideoListNamesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
