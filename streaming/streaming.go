// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package streaming

import (
	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// StreamingService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStreamingService] method instead.
type StreamingService struct {
	Options     []option.RequestOption
	AITasks     AITaskService
	Broadcasts  BroadcastService
	Directories DirectoryService
	Players     PlayerService
	QualitySets QualitySetService
	Playlists   PlaylistService
	Videos      VideoService
	Streams     StreamService
	Restreams   RestreamService
	Statistics  StatisticService
}

// NewStreamingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamingService(opts ...option.RequestOption) (r StreamingService) {
	r = StreamingService{}
	r.Options = opts
	r.AITasks = NewAITaskService(opts...)
	r.Broadcasts = NewBroadcastService(opts...)
	r.Directories = NewDirectoryService(opts...)
	r.Players = NewPlayerService(opts...)
	r.QualitySets = NewQualitySetService(opts...)
	r.Playlists = NewPlaylistService(opts...)
	r.Videos = NewVideoService(opts...)
	r.Streams = NewStreamService(opts...)
	r.Restreams = NewRestreamService(opts...)
	r.Statistics = NewStatisticService(opts...)
	return
}

// The property Name is required.
type CreateVideoParam struct {
	// Video name
	Name string `json:"name,required"`
	// Custom field where you can specify user ID in your system
	ClientUserID param.Opt[int64] `json:"client_user_id,omitzero"`
	// The length of the trimmed segment to transcode, instead of the entire length of
	// the video. Is only used in conjunction with specifying the start of a segment.
	// Transcoding duration is a number in seconds.
	ClipDurationSeconds param.Opt[int64] `json:"clip_duration_seconds,omitzero"`
	// If you want to transcode only a trimmed segment of a video instead of entire
	// length if the video, then you can provide timecodes of starting point and
	// duration of a segment to process. Start encoding from is a number in seconds.
	ClipStartSeconds param.Opt[int64] `json:"clip_start_seconds,omitzero"`
	// Deprecated. Custom URL of IFrame for video player to be used in share panel in
	// player. Auto generated IFrame URL provided by default
	CustomIframeURL param.Opt[string] `json:"custom_iframe_url,omitzero"`
	// Video details; not visible to the end-users
	Description param.Opt[string] `json:"description,omitzero"`
	// ID of the directory where the video should be uploaded. (beta)
	DirectoryID param.Opt[int64] `json:"directory_id,omitzero"`
	// Authorization HTTP request header. Will be used as credentials to authenticate a
	// request to download a file (specified in "`origin_url`" parameter) on an
	// external server. Syntax: `Authorization: ` Examples:
	//
	//   - "`origin_http_headers`": "Authorization: Basic ..."
	//   - "`origin_http_headers`": "Authorization: Bearer ..."
	//   - "`origin_http_headers`": "Authorization: APIKey ..." Example of usage when
	//     downloading a file from Google Drive:
	//
	// ```
	// POST https://api.gcore.com/streaming/videos
	// "video": {
	// "name": "IBC 2024 intro.mp4",
	// "origin_url": "https://www.googleapis.com/drive/v3/files/...?alt=media",
	// "origin_http_headers": "Authorization: Bearer ABC"
	// }
	// ```
	OriginHTTPHeaders param.Opt[string] `json:"origin_http_headers,omitzero"`
	// URL to an original file which you want to copy from external storage. If
	// specified, system will download the file and will use it as video source for
	// transcoding.
	OriginURL param.Opt[string] `json:"origin_url,omitzero"`
	// Poster is your own static image which can be displayed before the video starts.
	// After uploading the video, the system will automatically create several
	// screenshots (they will be stored in "screenshots" attribute) from which you can
	// select an default screenshot. This "poster" field is for uploading your own
	// image. Also use attribute "`screenshot_id`" to select poster as a default
	// screnshot. Attribute accepts single image as base64-encoded string
	// [(RFC 2397 – The "data" URL scheme)](https://www.rfc-editor.org/rfc/rfc2397). In
	// format: `data:[];base64,` MIME-types are image/jpeg, image/webp, and image/png
	// and file sizes up to 1Mb. Examples:
	//
	// - `data:image/jpeg;base64,/9j/4AA...qf/2Q==`
	// - `data:image/png;base64,iVBORw0KGg...ggg==`
	// - `data:image/webp;base64,UklGRt.../DgAAAAA`
	Poster param.Opt[string] `json:"poster,omitzero"`
	// Priority allows you to adjust the urgency of processing some videos before
	// others in your account, if your algorithm requires it. For example, when there
	// are very urgent video and some regular ones that can wait in the queue. Value
	// range, integer [-10..10]. -10 is the lowest down-priority, 10 is the highest
	// up-priority. Default priority is 0.
	Priority param.Opt[int64] `json:"priority,omitzero"`
	// Deprecated. Regulates the video format:
	//
	// - **regular** — plays the video as usual
	// - **vr360** — plays the video in 360 degree mode
	// - **vr180** — plays the video in 180 degree mode
	// - **vr360tb** — plays the video in 3D 360 degree mode Top-Bottom.
	//
	// Default is regular
	Projection param.Opt[string] `json:"projection,omitzero"`
	// Custom quality set ID for transcoding, if transcoding is required according to
	// your conditions. Look at GET /`quality_sets` method
	QualitySetID param.Opt[int64] `json:"quality_set_id,omitzero"`
	// Poster URL to download from external resource, instead of uploading via "poster"
	// attribute. It has the same restrictions as "poster" attribute.
	RemotePosterURL param.Opt[string] `json:"remote_poster_url,omitzero"`
	// Set it to true to remove poster
	RemovePoster param.Opt[bool] `json:"remove_poster,omitzero"`
	// Default screenshot index. Specify an ID from the "screenshots" array, so that
	// the URL of the required screenshot appears in the "screenshot" attribute as the
	// default screenshot. By default 5 static screenshots will be taken from different
	// places in the video after transcoding. If the video is short, there may be fewer
	// screenshots. Counting from 0. A value of -1 sets the default screenshot to the
	// URL of your own image from the "poster" attribute. Look at "screenshot"
	// attribute in GET /videos/{`video_id`} for details.
	ScreenshotID param.Opt[int64] `json:"screenshot_id,omitzero"`
	// Deprecated. Custom URL or iframe displayed in the link field when a user clicks
	// on a sharing button in player. If empty, the link field and social network
	// sharing is disabled
	ShareURL param.Opt[string] `json:"share_url,omitzero"`
	// The option allows you to set the video transcoding rule so that the output
	// bitrate in ABR ladder is not exceeding the bitrate of the original video.
	//
	// This option is for advanced users only.
	//
	// By default `source_bitrate_limit: true` this option allows you to have the
	// output bitrate not more than in the original video, thus to transcode video
	// faster and to deliver it to end-viewers faster as well. At the same time, the
	// quality will be similar to the original. If for some reason you need more
	// byte-space in the output quality when encoding, you can set this option to
	// `source_bitrate_limit: false`. Then, when transcoding, the quality ceiling will
	// be raised from the bitrate of the original video to the maximum possible limit
	// specified in our the Product Documentation. For example, this may be needed
	// when:
	//
	//   - to improve the visual quality parameters using PSNR, SSIM, VMAF metrics,
	//   - to improve the picture quality on dynamic scenes,
	//   - etc. The option is applied only at the video creation stage and cannot be
	//     changed later. If you want to re-transcode the video using new value, then you
	//     need to create and upload a new video only.
	SourceBitrateLimit param.Opt[bool] `json:"source_bitrate_limit,omitzero"`
	// Automatic creation of subtitles by transcribing the audio track. Values:
	//
	//   - disable – Do not transcribe.
	//   - auto – Automatically detects the activation of the option based on the
	//     settings in your account. If generation is activated, then automatic language
	//     detection while transcribing.
	//   - \ – Transcribe from specific language. Can be used to specify the exact
	//     language spoken in the audio track, or when auto language detection fails.
	//     Language is set by 3-letter language code according to ISO-639-2
	//     (bibliographic code). List of languages is available in `audio_language`
	//     attribute of API POST /streaming/ai/transcribe . Example:
	//
	// ```
	// auto_transcribe_audio_language: "auto"
	// auto_transcribe_audio_language: "ger"
	// ```
	//
	// More details:
	//
	//   - List of AI tasks – API
	//     [GET /streaming/ai/tasks](https://api.gcore.com/docs/streaming#tag/AI/operation/get_ai_results)
	//   - Add subtitles to an exist video – API
	//     [POST /streaming/videos/{`video_id`}/subtitles](https://api.gcore.com/docs/streaming#tag/Subtitles/operation/post_api_videos_video_id_subtitles).
	//
	// Any of "disable", "auto", "<language_code>".
	AutoTranscribeAudioLanguage CreateVideoAutoTranscribeAudioLanguage `json:"auto_transcribe_audio_language,omitzero"`
	// Automatic translation of auto-transcribed subtitles to the specified
	// language(s). Can be used both together with `auto_transcribe_audio_language`
	// option only. Use it when you want to make automatic subtitles in languages other
	// than the original language in audio. Values:
	//
	//   - disable – Do not translate.
	//   - default – There are 3 default languages: eng,fre,ger
	//   - \ – Explicit language to translate to, or list of languages separated by a
	//     comma. Look at list of available languages in description of AI ASR task
	//     creation. If several languages are specified for translation, a separate
	//     subtitle will be generated for each language. Example:
	//
	// ```
	// auto_translate_subtitles_language: default
	// auto_translate_subtitles_language: eng,fre,ger
	// ```
	//
	// Please note that subtitle translation is done separately and after
	// transcription. Thus separate AI-tasks are created for translation.
	//
	// Any of "disable", "default", "<language_codes,>".
	AutoTranslateSubtitlesLanguage CreateVideoAutoTranslateSubtitlesLanguage `json:"auto_translate_subtitles_language,omitzero"`
	paramObj
}

func (r CreateVideoParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateVideoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateVideoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Automatic creation of subtitles by transcribing the audio track. Values:
//
//   - disable – Do not transcribe.
//   - auto – Automatically detects the activation of the option based on the
//     settings in your account. If generation is activated, then automatic language
//     detection while transcribing.
//   - \ – Transcribe from specific language. Can be used to specify the exact
//     language spoken in the audio track, or when auto language detection fails.
//     Language is set by 3-letter language code according to ISO-639-2
//     (bibliographic code). List of languages is available in `audio_language`
//     attribute of API POST /streaming/ai/transcribe . Example:
//
// ```
// auto_transcribe_audio_language: "auto"
// auto_transcribe_audio_language: "ger"
// ```
//
// More details:
//
//   - List of AI tasks – API
//     [GET /streaming/ai/tasks](https://api.gcore.com/docs/streaming#tag/AI/operation/get_ai_results)
//   - Add subtitles to an exist video – API
//     [POST /streaming/videos/{`video_id`}/subtitles](https://api.gcore.com/docs/streaming#tag/Subtitles/operation/post_api_videos_video_id_subtitles).
type CreateVideoAutoTranscribeAudioLanguage string

const (
	CreateVideoAutoTranscribeAudioLanguageDisable      CreateVideoAutoTranscribeAudioLanguage = "disable"
	CreateVideoAutoTranscribeAudioLanguageAuto         CreateVideoAutoTranscribeAudioLanguage = "auto"
	CreateVideoAutoTranscribeAudioLanguageLanguageCode CreateVideoAutoTranscribeAudioLanguage = "<language_code>"
)

// Automatic translation of auto-transcribed subtitles to the specified
// language(s). Can be used both together with `auto_transcribe_audio_language`
// option only. Use it when you want to make automatic subtitles in languages other
// than the original language in audio. Values:
//
//   - disable – Do not translate.
//   - default – There are 3 default languages: eng,fre,ger
//   - \ – Explicit language to translate to, or list of languages separated by a
//     comma. Look at list of available languages in description of AI ASR task
//     creation. If several languages are specified for translation, a separate
//     subtitle will be generated for each language. Example:
//
// ```
// auto_translate_subtitles_language: default
// auto_translate_subtitles_language: eng,fre,ger
// ```
//
// Please note that subtitle translation is done separately and after
// transcription. Thus separate AI-tasks are created for translation.
type CreateVideoAutoTranslateSubtitlesLanguage string

const (
	CreateVideoAutoTranslateSubtitlesLanguageDisable       CreateVideoAutoTranslateSubtitlesLanguage = "disable"
	CreateVideoAutoTranslateSubtitlesLanguageDefault       CreateVideoAutoTranslateSubtitlesLanguage = "default"
	CreateVideoAutoTranslateSubtitlesLanguageLanguageCodes CreateVideoAutoTranslateSubtitlesLanguage = "<language_codes,>"
)

type Video struct {
	// Video ID
	ID int64 `json:"id"`
	// ID of ad that should be shown. If empty the default ad is show. If there is no
	// default ad, no ad is shownю
	AdID int64 `json:"ad_id"`
	// Total number of video views. It is calculated based on the analysis of all
	// views, no matter in which player.
	CdnViews int64 `json:"cdn_views"`
	// Client ID
	ClientID int64 `json:"client_id"`
	// Custom meta field for storing the Identifier in your system. We do not use this
	// field in any way when processing the stream. Example: `client_user_id = 1001`
	ClientUserID int64 `json:"client_user_id"`
	// Array of data about each transcoded quality
	ConvertedVideos []VideoConvertedVideo `json:"converted_videos"`
	// Custom URL of Iframe for video player to be used in share panel in player. Auto
	// generated Iframe URL provided by default.
	CustomIframeURL string `json:"custom_iframe_url"`
	// A URL to a master playlist MPEG-DASH (master.mpd) with CMAF or WebM based
	// chunks. Chunk type will be selected automatically for each quality:
	//
	// - CMAF for H264 and H265 codecs.
	// - WebM for AV1 codec.
	//
	// This URL is a link to the main manifest. But you can also manually specify
	// suffix-options that will allow you to change the manifest to your request:
	// `/videos/{client_id}_{slug}/master[-min-N][-max-N][-(h264|hevc|av1)].mpd` List
	// of suffix-options:
	//
	//   - [-min-N] – ABR soft limitation of qualities from below.
	//   - [-max-N] – ABR soft limitation of qualities from above.
	//   - [-(h264|hevc|av1) – Video codec soft limitation. Applicable if the video was
	//     transcoded into multiple codecs H264, H265 and AV1 at once, but you want to
	//     return just 1 video codec in a manifest. Read the Product Documentation for
	//     details. Read more what is ABR soft-limiting in the "`hls_url`" field above.
	//
	// Caution. Solely master.mpd is officially documented and intended for your use.
	// Any additional internal manifests, sub-manifests, parameters, chunk names, file
	// extensions, and related components are internal infrastructure entities. These
	// may undergo modifications without prior notice, in any manner or form. It is
	// strongly advised not to store them in your database or cache them on your end.
	DashURL string `json:"dash_url"`
	// Additional text field for video description
	Description string `json:"description"`
	// Video duration in milliseconds. May differ from "`origin_video_duration`" value
	// if the video was uploaded with clipping through the parameters
	// "`clip_start_seconds`" and "`clip_duration_seconds`"
	Duration int64 `json:"duration"`
	// Video processing error text will be saved here if "status: error"
	Error string `json:"error"`
	// A URL to a master playlist HLS (master-cmaf.m3u8) with CMAF-based chunks. Chunks
	// are in fMP4 container. It's a code-agnostic container, which allows to use any
	// like H264, H265, AV1, etc.
	//
	// It is possible to use the same suffix-options as described in the "`hls_url`"
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
	// `/videos/{client_id}_{video_slug}/master[-cmaf][-min-N][-max-N][-img][-(h264|hevc|av1)].m3u8`
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
	//     details. ABR soft-limiting: Soft limitation of the list of qualities allows
	//     you to return not the entire list of transcoded qualities for a video, but
	//     only those you need. For more details look at the Product Documentation. For
	//     example, the video is available in 7 qualities from 360p to 4K, but you want
	//     to return not more than 480p only due to the conditions of distribution of
	//     content to a specific end-user (i.e. free account):
	//   - To a generic `.../master.m3u8` manifest
	//   - Add a suffix-option to limit quality `.../master-max-480.m3u8`
	//   - Add a suffix-option to limit quality and codec
	//     `.../master-min-320-max-320-h264.m3u8`
	//
	// Caution. Solely master.m3u8 (and master[-options].m3u8) is officially documented
	// and intended for your use. Any additional internal manifests, sub-manifests,
	// parameters, chunk names, file extensions, and related components are internal
	// infrastructure entities. These may undergo modifications without prior notice,
	// in any manner or form. It is strongly advised not to store them in your database
	// or cache them on your end.
	HlsURL string `json:"hls_url"`
	// A URL to a built-in HTML video player with the video inside. It can be inserted
	// into an iframe on your website and the video will automatically play in all
	// browsers. The player can be opened or shared via this direct link. Also the
	// video player can be integrated into your web pages using the Iframe tag. Example
	// of usage on a web page:
	//
	// <iframe width="100%" height="100%" src="https://player.gvideo.co/videos/2675_FnlHXwA16ZMxmUr" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>
	//
	// There are some link modificators you can specify and add manually:
	// - ?`no_low_latency` – player is forced to use non-low-latency streams HLS MPEG TS, instead of MPEG-DASH CMAF or HLS/LL-HLS CMAF.
	// - ?t=(integer) – time to start playback from specified point in the video. Applicable for VOD only.
	// - ?`sub_lang`=(language) – force subtitles to specific language (2 letters ISO 639 code of a language).
	// - Read more in the Product Documentation.
	IframeURL string `json:"iframe_url"`
	// Title of the video. Often used as a human-readable name of the video, but can
	// contain any text you wish. The values are not unique and may be repeated.
	// Examples:
	//
	// - Educational training 2024-03-29
	// - Series X S3E14, The empire strikes back
	// - 480fd499-2de2-4988-bc1a-a4eebe9818ee
	Name string `json:"name"`
	// Size of original file
	OriginSize int64 `json:"origin_size"`
	// URL to an original file from which the information for transcoding was taken.
	// May contain a link for scenarios:
	//
	//   - If the video was downloaded from another origin
	//
	//   - If the video is a recording of a live stream
	//
	//   - Otherwise it is "null" **Copy from another server** URL to an original file
	//     that was downloaded. Look at method "Copy from another server" in POST
	//     /videos. **Recording of an original live stream** URL to the original
	//     non-transcoded stream recording with original quality, saved in MP4 format.
	//     File is created immediately after the completion of the stream recording. The
	//     stream from which the recording was made is reflected in "`stream_id`" field.
	//     Can be used for internal operations when a recording needs to be received
	//     faster than the transcoded versions are ready. But this version is not
	//     intended for public distribution. Views and downloads occur in the usual way,
	//     like viewing an MP4 rendition. The MP4 file becomes available for downloading
	//     when the video entity "status" changes from "new" to "pending". The file is
	//     stored for 7 days, after which it will be automatically deleted. Format of URL
	//     is `/videos/_/origin__.mp4` Where:
	//
	//   - ```– Encoding bitrate in Kbps.
	//
	//     ```
	//   - ```– Video height.
	//     This is a premium feature, available only upon request through your manager or support team.
	//     ```
	OriginURL string `json:"origin_url"`
	// Original video duration in milliseconds
	OriginVideoDuration int64 `json:"origin_video_duration"`
	// Poster is your own static image which can be displayed before the video begins
	// playing. This is often a frame of the video or a custom title screen. Field
	// contains a link to your own uploaded image. Also look at "screenshot" attribute.
	Poster string `json:"poster"`
	// Field contains a link to minimized poster image. Original "poster" image is
	// proportionally scaled to a size of 200 pixels in height.
	PosterThumb string `json:"poster_thumb"`
	// Regulates the video format:
	//
	// - **regular** — plays the video as usual
	// - **vr360** — plays the video in 360 degree mode
	// - **vr180** — plays the video in 180 degree mode
	// - **vr360tb** — plays the video in 3D 360 degree mode Top-Bottom.
	//
	// Default is regular
	Projection string `json:"projection"`
	// If the video was saved from a stream, then start time of the stream recording is
	// saved here. Format is date time in ISO 8601
	RecordingStartedAt string `json:"recording_started_at"`
	// A URL to the default screenshot is here. The image is selected from an array of
	// all screenshots based on the “`screenshot_id`” attribute. If you use your own
	// "poster", the link to it will be here too. Our video player uses this field to
	// display the static image before the video starts playing. As soon as the user
	// hits "play" the image will go away. If you use your own external video player,
	// then you can use the value of this field to set the poster/thumbnail in your
	// player. Example:
	//
	// - `video_js`.poster: `api.screenshot`
	// - clappr.poster: `api.screenshot`
	Screenshot string `json:"screenshot"`
	// ID of auto generated screenshots to be used for default screenshot. Counting
	// from 0. A value of -1 sets the "screenshot" attribute to the URL of your own
	// image from the "poster" attribute.
	ScreenshotID int64 `json:"screenshot_id"`
	// Array of auto generated screenshots from the video. By default 5 static
	// screenshots are taken from different places in the video. If the video is short,
	// there may be fewer screenshots. Screenshots are created automatically, so they
	// may contain not very good frames from the video. To use your own image look at
	// "poster" attribute.
	Screenshots []string `json:"screenshots"`
	// Custom URL or iframe displayed in the link field when a user clicks on a sharing
	// button in player. If empty, the link field and social network sharing is
	// disabled
	ShareURL string `json:"share_url"`
	// A unique alphanumeric identifier used in public URLs to retrieve and view the
	// video. It is unique for each video, generated randomly and set automatically by
	// the system. Format of usage in URL is \*.../videos/{`client_id`}\_{slug}/...\*
	// Example:
	//
	// - Player: /videos/`12345_neAq1bYZ2`
	// - Manifest: /videos/`12345_neAq1bYZ2`/master.m3u8
	// - Rendition: /videos/`12345_neAq1bYZ2`/`qid90v1_720`.mp4
	Slug string `json:"slug"`
	// Link to picture with video storyboard. Image in JPG format. The picture is a set
	// of rectangles with frames from the video. Typically storyboard is used to show
	// preview images when hovering the video's timeline.
	Sprite string `json:"sprite"`
	// Storyboard in VTT format. This format implies an explicit indication of the
	// timing and frame area from a large sprite image.
	SpriteVtt string `json:"sprite_vtt"`
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
	Status VideoStatus `json:"status"`
	// If the video was saved from a stream, then ID of that stream is saved here
	StreamID int64 `json:"stream_id"`
	// Number of video views through the built-in HTML video player of the Streaming
	// Platform only. This attribute does not count views from other external players
	// and native OS players, so here may be less number of views than in
	// "`cdn_views`".
	Views int64 `json:"views"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AdID                respjson.Field
		CdnViews            respjson.Field
		ClientID            respjson.Field
		ClientUserID        respjson.Field
		ConvertedVideos     respjson.Field
		CustomIframeURL     respjson.Field
		DashURL             respjson.Field
		Description         respjson.Field
		Duration            respjson.Field
		Error               respjson.Field
		HlsCmafURL          respjson.Field
		HlsURL              respjson.Field
		IframeURL           respjson.Field
		Name                respjson.Field
		OriginSize          respjson.Field
		OriginURL           respjson.Field
		OriginVideoDuration respjson.Field
		Poster              respjson.Field
		PosterThumb         respjson.Field
		Projection          respjson.Field
		RecordingStartedAt  respjson.Field
		Screenshot          respjson.Field
		ScreenshotID        respjson.Field
		Screenshots         respjson.Field
		ShareURL            respjson.Field
		Slug                respjson.Field
		Sprite              respjson.Field
		SpriteVtt           respjson.Field
		Status              respjson.Field
		StreamID            respjson.Field
		Views               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Video) RawJSON() string { return r.JSON.raw }
func (r *Video) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoConvertedVideo struct {
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
	// **Download methods** For each converted video, additional download endpoints are
	// available under `converted_videos`/`mp4_urls`. An MP4 download enpoints:
	//
	//   - /videos/{`client_id`}\_{slug}/{filename}.mp4
	//   - /videos/{`client_id`}\_{slug}/{filename}.mp4/download
	//   - /videos/{`client_id`}\_{slug}/{filename}.mp4/download={`custom_filename`} The
	//     first option returns the file as is. The following options respond with the
	//     header that directly tells browsers to download the file instead of playing it
	//     in the browser.
	//
	// ```
	// Content-Disposition: attachment
	// ```
	//
	// The third option allows you to set a custom name for the file being downloaded.
	// You can optionally specify a custom filename (just name excluding the .mp4
	// extension) using the download= query. Filename Constraints
	//
	//   - Length: 1-255 characters
	//   - Must NOT include the .mp4 extension (it is added automatically)
	//   - Allowed characters: a-z, A-Z, 0-9, \_(underscore), -(dash), .(dot)
	//   - First character cannot be .(dot) Example valid filenames: `holiday2025`,
	//     `_backup.final`, `clip-v1.2`
	//
	// **Default MP4 file name structure** Link to the file {filename} contains
	// information about the encoding method using format: `___.mp4`
	//
	// - ```– Internal quality identifier and file version. Please do not use it, can be changed at any time without any notice.
	//
	//	```
	//
	// - ```– Codec name that was used to encode the video, or audio codec if it is an audio-only file.
	//
	//	```
	//
	// - ```– Encoding bitrate in Kbps.
	//
	//	```
	//
	//   - ````– Video height, or word "audio" if it is an audio-only file.
	//     Note that this link format has been applied since 14.08.2024. If the video entity was uploaded earlier, links may have old simplified format.
	//     Example: ``` /videos/{client_id}_{slug}/qid3567v1_h264_4050_1080.mp4 ```
	//
	//     ````
	//
	// **Dynamic speed limiting** This mode sets different limits for different users
	// or for different types of content. The speed is adjusted based on requests with
	// the “speed” and “buffer” arguments. Example: `?speed=50k&buffer=500k` Read more
	// in Product Documentation in CDN section "Network limits".
	//
	// **Secure token authentication (updated)** Access to MP4 download links can be
	// protected using secure tokens passed as query parameters. The token generation
	// logic has been updated to allow fine-grained protection per file and bitrate.
	// Token generation uses the entire MP4 path, which ensures the token only grants
	// access to a specific quality/version of the video. This prevents unintended
	// access to other bitrate versions of an ABR stream. Token Query Parameters:
	//
	//   - token: The generated hash
	//   - expires: Expiration timestamp
	//   - speed: (optional) Speed limit in bytes/sec, or empty string
	//   - buffer: (optional) Buffer size in bytes, or empty string Optional (for
	//     IP-bound tokens):
	//   - ip: The user’s IP address Example:
	//     `?md5=QX39c77lbQKvYgMMAvpyMQ&expires=1743167062` Read more in Product
	//     Documentation in Streaming section "Protected temporarily link".
	//
	// **Examples**
	//
	//   - Audio-only:
	//     `https://demo-public.gvideo.io/videos/2675_JNnccG5l97XPxsov/qid3585v1_aac_128_audio.mp4`
	//   - Video:
	//     `https://demo-public.gvideo.io/videos/2675_3MlggU4xDb1Ssa5Y/qid3567v1_h264_4050_1080.mp4/download`
	//   - Video with custom download filename:
	//     `https://demo-public.gvideo.io/videos/2675_XtMKxzJM6Xt7SBUV/1080.mp4/download=highlights_v1.1_2025-05-30`
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
func (r VideoConvertedVideo) RawJSON() string { return r.JSON.raw }
func (r *VideoConvertedVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
type VideoStatus string

const (
	VideoStatusEmpty    VideoStatus = "empty"
	VideoStatusPending  VideoStatus = "pending"
	VideoStatusViewable VideoStatus = "viewable"
	VideoStatusReady    VideoStatus = "ready"
	VideoStatusError    VideoStatus = "error"
)
