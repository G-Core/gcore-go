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

// PlaylistService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlaylistService] method instead.
type PlaylistService struct {
	Options []option.RequestOption
}

// NewPlaylistService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlaylistService(opts ...option.RequestOption) (r PlaylistService) {
	r = PlaylistService{}
	r.Options = opts
	return
}

// Playlist is a curated collection of video content organized in a sequential
// manner.
//
// This method offers several advantages and features that are typical of live
// streaming but with more control over the content. Here's how it works:
//
//   - Playlist always consists only of static VOD videos you previously uploaded to
//     the system.
//   - Playlist is always played as a "Live stream" for end-users, so without the
//     ability to fast forward the stream to the “future”. Manifest will contain
//     chunks as for live stream too.
//   - Playlist can be looped endlessly. In this case, all the videos in the list
//     will be constantly repeated through the list.
//   - Playlist can be programmed to be played at a specific time in the future. In
//     that case, before the start time there will be empty manifest.
//
// You can add new videos to the list, remove unnecessary videos, or change the
// order of videos in the list. But please pay attention to when the video list
// changes, it is updated instantly on the server. This means that after saving the
// changed list, the playlist will be reloaded for all users and it will start
// plays from the very first element.
//
// Maximum video limit = 128 videos in a row.
//
// Examples of usage:
//
// - Looped video playback
// - Scheduled playback
//
// **Looped video playback**
//
// It can be used to simulate TV channel pre-programmed behaviour.
//
//   - Selection: Choose a series of videos, such as TV show episodes, movies,
//     tutorials, or any other relevant content.
//   - Order: Arrange the selected videos in the desired sequence, much like setting
//     a broadcast schedule.
//   - Looping: Optionally, the playlist can be set to loop, replaying the sequence
//     once it finishes to maintain a continuous stream.
//
// Example:
//
// ```
//
//	active: true
//	loop: true
//	name: "Playlist: TV channel 'The world around us' (Programmed broadcast for 24 hours)"
//
// ```
//
// **Scheduled playback**
//
// It can be used to simulate live events such as virtual concerts, webinars, or
// any special broadcasts without the logistical challenges of an actual live
// stream.
//
//   - Timing: Set specific start time, creating the illusion of a live broadcast
//     schedule.
//   - Selection: Choose a video or series of videos to be played at the specified
//     time.
//   - No Pauses: Unlike on-demand streaming where users can pause and skip, this
//     emulated live stream runs continuously, mirroring the constraints of
//     traditional live broadcasts.
//
// ```
//
//	active: true
//	loop: false
//	name: "Playlist: Webinar 'Onboarding for new employees on working with the corporate portal'"
//	start_time: "2024-07-01T11:00:00Z"
//
// ```
func (r *PlaylistService) New(ctx context.Context, body PlaylistNewParams, opts ...option.RequestOption) (res *PlaylistCreate, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/playlists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change playlist
func (r *PlaylistService) Update(ctx context.Context, playlistID int64, body PlaylistUpdateParams, opts ...option.RequestOption) (res *Playlist, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("streaming/playlists/%v", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of created playlists
func (r *PlaylistService) List(ctx context.Context, query PlaylistListParams, opts ...option.RequestOption) (res *pagination.PageStreaming[Playlist], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/playlists"
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

// Returns a list of created playlists
func (r *PlaylistService) ListAutoPaging(ctx context.Context, query PlaylistListParams, opts ...option.RequestOption) *pagination.PageStreamingAutoPager[Playlist] {
	return pagination.NewPageStreamingAutoPager(r.List(ctx, query, opts...))
}

// Delete playlist
func (r *PlaylistService) Delete(ctx context.Context, playlistID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("streaming/playlists/%v", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns a playlist details
func (r *PlaylistService) Get(ctx context.Context, playlistID int64, opts ...option.RequestOption) (res *Playlist, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("streaming/playlists/%v", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Shows ordered array of playlist videos
func (r *PlaylistService) ListVideos(ctx context.Context, playlistID int64, opts ...option.RequestOption) (res *[]PlaylistVideo, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("streaming/playlists/%v/videos", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Playlist struct {
	// Enables/Disables playlist. Has two possible values:
	//
	// - true – Playlist can be played.
	// - false – Playlist is disabled. No broadcast while it's desabled.
	Active bool `json:"active"`
	// The advertisement ID that will be inserted into the video
	AdID int64 `json:"ad_id"`
	// Current playlist client ID
	ClientID int64 `json:"client_id"`
	// Custom field where you can specify user ID in your system
	ClientUserID int64 `json:"client_user_id"`
	// Enables countdown before playlist start with `playlist_type: live`
	Countdown bool `json:"countdown"`
	// A URL to a master playlist HLS (master-cmaf.m3u8) with CMAF-based chunks. Chunks
	// are in fMP4 container.
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
	// A URL to a master playlist HLS (master.m3u8) with MPEG TS container.
	//
	// This URL is a link to the main manifest. But you can also manually specify
	// suffix-options that will allow you to change the manifest to your request:
	// `/playlists/{client_id}_{playlist_id}/master[-cmaf][-min-N][-max-N][-img][-(h264|hevc|av1)].m3u8`
	// Please see the details in `hls_url` attribute of /videos/{id} method.
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
	// browsers.
	//
	// The player can be opened or shared via this direct link. Also the video player
	// can be integrated into your web pages using the Iframe tag.
	//
	// Please see the details in `iframe_url` attribute of /videos/{id} method.
	IframeURL string `json:"iframe_url"`
	// Enables/Disables playlist loop
	Loop bool `json:"loop"`
	// Playlist name
	Name string `json:"name"`
	// The player ID with which the video will be played
	PlayerID int64 `json:"player_id"`
	// Determines whether the playlist:
	//
	// - `live` - playlist for live-streaming
	// - `vod` - playlist is for video on demand access
	//
	// Any of "live", "vod".
	PlaylistType PlaylistPlaylistType `json:"playlist_type"`
	// Playlist start time. Playlist won't be available before the specified time.
	// Datetime in ISO 8601 format.
	StartTime string `json:"start_time"`
	// A list of VOD IDs included in the playlist. Order of videos in a playlist
	// reflects the order of IDs in the array.
	//
	// Maximum video limit = 128.
	VideoIDs []int64 `json:"video_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active       respjson.Field
		AdID         respjson.Field
		ClientID     respjson.Field
		ClientUserID respjson.Field
		Countdown    respjson.Field
		HlsCmafURL   respjson.Field
		HlsURL       respjson.Field
		IframeURL    respjson.Field
		Loop         respjson.Field
		Name         respjson.Field
		PlayerID     respjson.Field
		PlaylistType respjson.Field
		StartTime    respjson.Field
		VideoIDs     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Playlist) RawJSON() string { return r.JSON.raw }
func (r *Playlist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Playlist to a PlaylistParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PlaylistParam.Overrides()
func (r Playlist) ToParam() PlaylistParam {
	return param.Override[PlaylistParam](json.RawMessage(r.RawJSON()))
}

// Determines whether the playlist:
//
// - `live` - playlist for live-streaming
// - `vod` - playlist is for video on demand access
type PlaylistPlaylistType string

const (
	PlaylistPlaylistTypeLive PlaylistPlaylistType = "live"
	PlaylistPlaylistTypeVod  PlaylistPlaylistType = "vod"
)

type PlaylistParam struct {
	// Enables/Disables playlist. Has two possible values:
	//
	// - true – Playlist can be played.
	// - false – Playlist is disabled. No broadcast while it's desabled.
	Active param.Opt[bool] `json:"active,omitzero"`
	// The advertisement ID that will be inserted into the video
	AdID param.Opt[int64] `json:"ad_id,omitzero"`
	// Current playlist client ID
	ClientID param.Opt[int64] `json:"client_id,omitzero"`
	// Custom field where you can specify user ID in your system
	ClientUserID param.Opt[int64] `json:"client_user_id,omitzero"`
	// Enables countdown before playlist start with `playlist_type: live`
	Countdown param.Opt[bool] `json:"countdown,omitzero"`
	// A URL to a master playlist HLS (master-cmaf.m3u8) with CMAF-based chunks. Chunks
	// are in fMP4 container.
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
	HlsCmafURL param.Opt[string] `json:"hls_cmaf_url,omitzero"`
	// A URL to a master playlist HLS (master.m3u8) with MPEG TS container.
	//
	// This URL is a link to the main manifest. But you can also manually specify
	// suffix-options that will allow you to change the manifest to your request:
	// `/playlists/{client_id}_{playlist_id}/master[-cmaf][-min-N][-max-N][-img][-(h264|hevc|av1)].m3u8`
	// Please see the details in `hls_url` attribute of /videos/{id} method.
	//
	// Caution. Solely master.m3u8 (and master[-options].m3u8) is officially documented
	// and intended for your use. Any additional internal manifests, sub-manifests,
	// parameters, chunk names, file extensions, and related components are internal
	// infrastructure entities. These may undergo modifications without prior notice,
	// in any manner or form. It is strongly advised not to store them in your database
	// or cache them on your end.
	HlsURL param.Opt[string] `json:"hls_url,omitzero"`
	// A URL to a built-in HTML video player with the video inside. It can be inserted
	// into an iframe on your website and the video will automatically play in all
	// browsers.
	//
	// The player can be opened or shared via this direct link. Also the video player
	// can be integrated into your web pages using the Iframe tag.
	//
	// Please see the details in `iframe_url` attribute of /videos/{id} method.
	IframeURL param.Opt[string] `json:"iframe_url,omitzero"`
	// Enables/Disables playlist loop
	Loop param.Opt[bool] `json:"loop,omitzero"`
	// Playlist name
	Name param.Opt[string] `json:"name,omitzero"`
	// The player ID with which the video will be played
	PlayerID param.Opt[int64] `json:"player_id,omitzero"`
	// Playlist start time. Playlist won't be available before the specified time.
	// Datetime in ISO 8601 format.
	StartTime param.Opt[string] `json:"start_time,omitzero"`
	// Determines whether the playlist:
	//
	// - `live` - playlist for live-streaming
	// - `vod` - playlist is for video on demand access
	//
	// Any of "live", "vod".
	PlaylistType PlaylistPlaylistType `json:"playlist_type,omitzero"`
	// A list of VOD IDs included in the playlist. Order of videos in a playlist
	// reflects the order of IDs in the array.
	//
	// Maximum video limit = 128.
	VideoIDs []int64 `json:"video_ids,omitzero"`
	paramObj
}

func (r PlaylistParam) MarshalJSON() (data []byte, err error) {
	type shadow PlaylistParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlaylistParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaylistCreate struct {
	// Playlist ID
	ID int64 `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Playlist
}

// Returns the unmodified JSON received from the API
func (r PlaylistCreate) RawJSON() string { return r.JSON.raw }
func (r *PlaylistCreate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaylistVideo struct {
	// Video name
	Name string `json:"name,required"`
	// Automatic creation of subtitles by transcribing the audio track.
	//
	// Values:
	//
	//   - disable – Do not transcribe.
	//   - auto – Automatically detects the activation of the option based on the
	//     settings in your account. If generation is activated, then automatic language
	//     detection while transcribing.
	//   - \ – Transcribe from specific language. Can be used to specify the exact
	//     language spoken in the audio track, or when auto language detection fails.
	//     Language is set by 3-letter language code according to ISO-639-2
	//     (bibliographic code). List of languages is available in `audio_language`
	//     attribute of API POST /streaming/ai/transcribe .
	//
	// Example:
	//
	// ```
	// auto_transcribe_audio_language: "auto"
	// auto_transcribe_audio_language: "ger"
	// ```
	//
	// More details:
	//
	//   - List of AI tasks – API
	//     [GET /streaming/ai/tasks](/docs/api-reference/streaming/ai/get-list-of-ai-tasks)
	//   - Add subtitles to an exist video – API
	//     [POST /streaming/videos/{`video_id`}/subtitles](/docs/api-reference/streaming/subtitles/add-subtitle).
	//
	// Any of "disable", "auto", "<language_code>".
	AutoTranscribeAudioLanguage PlaylistVideoAutoTranscribeAudioLanguage `json:"auto_transcribe_audio_language"`
	// Automatic translation of auto-transcribed subtitles to the specified
	// language(s). Can be used both together with `auto_transcribe_audio_language`
	// option only.
	//
	// Use it when you want to make automatic subtitles in languages other than the
	// original language in audio.
	//
	// Values:
	//
	//   - disable – Do not translate.
	//   - default – There are 3 default languages: eng,fre,ger
	//   - \ – Explicit language to translate to, or list of languages separated by a
	//     comma. Look at list of available languages in description of AI ASR task
	//     creation.
	//
	// If several languages are specified for translation, a separate subtitle will be
	// generated for each language.
	//
	// Example:
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
	AutoTranslateSubtitlesLanguage PlaylistVideoAutoTranslateSubtitlesLanguage `json:"auto_translate_subtitles_language"`
	// Custom field where you can specify user ID in your system
	ClientUserID int64 `json:"client_user_id"`
	// The length of the trimmed segment to transcode, instead of the entire length of
	// the video. Is only used in conjunction with specifying the start of a segment.
	// Transcoding duration is a number in seconds.
	ClipDurationSeconds int64 `json:"clip_duration_seconds"`
	// If you want to transcode only a trimmed segment of a video instead of entire
	// length if the video, then you can provide timecodes of starting point and
	// duration of a segment to process. Start encoding from is a number in seconds.
	ClipStartSeconds int64 `json:"clip_start_seconds"`
	// Deprecated.
	//
	// Custom URL of IFrame for video player to be used in share panel in player. Auto
	// generated IFrame URL provided by default
	CustomIframeURL string `json:"custom_iframe_url"`
	// Video details; not visible to the end-users
	Description string `json:"description"`
	// ID of the directory where the video should be uploaded. (beta)
	DirectoryID int64 `json:"directory_id"`
	// Authorization HTTP request header. Will be used as credentials to authenticate a
	// request to download a file (specified in "`origin_url`" parameter) on an
	// external server.
	//
	// Syntax: `Authorization: <auth-scheme> <authorization-parameters>`
	//
	// Examples:
	//
	//   - "`origin_http_headers`": "Authorization: Basic ..."
	//   - "`origin_http_headers`": "Authorization: Bearer ..."
	//   - "`origin_http_headers`": "Authorization: APIKey ..." Example of usage when
	//     downloading a file from Google Drive:
	//
	// ```
	// POST https://api.gcore.com/streaming/videos
	//
	//	"video": {
	//	  "name": "IBC 2024 intro.mp4",
	//	  "origin_url": "https://www.googleapis.com/drive/v3/files/...?alt=media",
	//	  "origin_http_headers": "Authorization: Bearer ABC"
	//	}
	//
	// ```
	OriginHTTPHeaders string `json:"origin_http_headers"`
	// URL to an original file which you want to copy from external storage. If
	// specified, system will download the file and will use it as video source for
	// transcoding.
	OriginURL string `json:"origin_url"`
	// Poster is your own static image which can be displayed before the video starts.
	//
	// After uploading the video, the system will automatically create several
	// screenshots (they will be stored in "screenshots" attribute) from which you can
	// select an default screenshot. This "poster" field is for uploading your own
	// image. Also use attribute "`screenshot_id`" to select poster as a default
	// screnshot.
	//
	// Attribute accepts single image as base64-encoded string
	// [(RFC 2397 – The "data" URL scheme)](https://www.rfc-editor.org/rfc/rfc2397). In
	// format: `data:[<mediatype>];base64,<data>`
	//
	// MIME-types are image/jpeg, image/webp, and image/png and file sizes up to 1Mb.
	//
	// Examples:
	//
	// - `data:image/jpeg;base64,/9j/4AA...qf/2Q==`
	// - `data:image/png;base64,iVBORw0KGg...ggg==`
	// - `data:image/webp;base64,UklGRt.../DgAAAAA`
	Poster string `json:"poster"`
	// Priority allows you to adjust the urgency of processing some videos before
	// others in your account, if your algorithm requires it. For example, when there
	// are very urgent video and some regular ones that can wait in the queue.
	//
	// Value range, integer [-10..10]. -10 is the lowest down-priority, 10 is the
	// highest up-priority. Default priority is 0.
	Priority int64 `json:"priority"`
	// Deprecated.
	//
	// Regulates the video format:
	//
	// - **regular** — plays the video as usual
	// - **vr360** — plays the video in 360 degree mode
	// - **vr180** — plays the video in 180 degree mode
	// - **vr360tb** — plays the video in 3D 360 degree mode Top-Bottom.
	//
	// Default is regular
	Projection string `json:"projection"`
	// Custom quality set ID for transcoding, if transcoding is required according to
	// your conditions. Look at GET /`quality_sets` method
	QualitySetID int64 `json:"quality_set_id"`
	// Poster URL to download from external resource, instead of uploading via "poster"
	// attribute.
	//
	// It has the same restrictions as "poster" attribute.
	RemotePosterURL string `json:"remote_poster_url"`
	// Set it to true to remove poster
	RemovePoster bool `json:"remove_poster"`
	// Default screenshot index.
	//
	// Specify an ID from the "screenshots" array, so that the URL of the required
	// screenshot appears in the "screenshot" attribute as the default screenshot. By
	// default 5 static screenshots will be taken from different places in the video
	// after transcoding. If the video is short, there may be fewer screenshots.
	//
	// Counting from 0. A value of -1 sets the default screenshot to the URL of your
	// own image from the "poster" attribute.
	//
	// Look at "screenshot" attribute in GET /videos/{`video_id`} for details.
	ScreenshotID int64 `json:"screenshot_id"`
	// Deprecated.
	//
	// Custom URL or iframe displayed in the link field when a user clicks on a sharing
	// button in player. If empty, the link field and social network sharing is
	// disabled
	ShareURL string `json:"share_url"`
	// The option allows you to set the video transcoding rule so that the output
	// bitrate in ABR ladder is not exceeding the bitrate of the original video.
	//
	// This option is for advanced users only.
	//
	// By default `source_bitrate_limit: true` this option allows you to have the
	// output bitrate not more than in the original video, thus to transcode video
	// faster and to deliver it to end-viewers faster as well. At the same time, the
	// quality will be similar to the original.
	//
	// If for some reason you need more byte-space in the output quality when encoding,
	// you can set this option to `source_bitrate_limit: false`. Then, when
	// transcoding, the quality ceiling will be raised from the bitrate of the original
	// video to the maximum possible limit specified in our the Product Documentation.
	// For example, this may be needed when:
	//
	// - to improve the visual quality parameters using PSNR, SSIM, VMAF metrics,
	// - to improve the picture quality on dynamic scenes,
	// - etc.
	//
	// The option is applied only at the video creation stage and cannot be changed
	// later. If you want to re-transcode the video using new value, then you need to
	// create and upload a new video only.
	SourceBitrateLimit bool `json:"source_bitrate_limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name                           respjson.Field
		AutoTranscribeAudioLanguage    respjson.Field
		AutoTranslateSubtitlesLanguage respjson.Field
		ClientUserID                   respjson.Field
		ClipDurationSeconds            respjson.Field
		ClipStartSeconds               respjson.Field
		CustomIframeURL                respjson.Field
		Description                    respjson.Field
		DirectoryID                    respjson.Field
		OriginHTTPHeaders              respjson.Field
		OriginURL                      respjson.Field
		Poster                         respjson.Field
		Priority                       respjson.Field
		Projection                     respjson.Field
		QualitySetID                   respjson.Field
		RemotePosterURL                respjson.Field
		RemovePoster                   respjson.Field
		ScreenshotID                   respjson.Field
		ShareURL                       respjson.Field
		SourceBitrateLimit             respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaylistVideo) RawJSON() string { return r.JSON.raw }
func (r *PlaylistVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Automatic creation of subtitles by transcribing the audio track.
//
// Values:
//
//   - disable – Do not transcribe.
//   - auto – Automatically detects the activation of the option based on the
//     settings in your account. If generation is activated, then automatic language
//     detection while transcribing.
//   - \ – Transcribe from specific language. Can be used to specify the exact
//     language spoken in the audio track, or when auto language detection fails.
//     Language is set by 3-letter language code according to ISO-639-2
//     (bibliographic code). List of languages is available in `audio_language`
//     attribute of API POST /streaming/ai/transcribe .
//
// Example:
//
// ```
// auto_transcribe_audio_language: "auto"
// auto_transcribe_audio_language: "ger"
// ```
//
// More details:
//
//   - List of AI tasks – API
//     [GET /streaming/ai/tasks](/docs/api-reference/streaming/ai/get-list-of-ai-tasks)
//   - Add subtitles to an exist video – API
//     [POST /streaming/videos/{`video_id`}/subtitles](/docs/api-reference/streaming/subtitles/add-subtitle).
type PlaylistVideoAutoTranscribeAudioLanguage string

const (
	PlaylistVideoAutoTranscribeAudioLanguageDisable      PlaylistVideoAutoTranscribeAudioLanguage = "disable"
	PlaylistVideoAutoTranscribeAudioLanguageAuto         PlaylistVideoAutoTranscribeAudioLanguage = "auto"
	PlaylistVideoAutoTranscribeAudioLanguageLanguageCode PlaylistVideoAutoTranscribeAudioLanguage = "<language_code>"
)

// Automatic translation of auto-transcribed subtitles to the specified
// language(s). Can be used both together with `auto_transcribe_audio_language`
// option only.
//
// Use it when you want to make automatic subtitles in languages other than the
// original language in audio.
//
// Values:
//
//   - disable – Do not translate.
//   - default – There are 3 default languages: eng,fre,ger
//   - \ – Explicit language to translate to, or list of languages separated by a
//     comma. Look at list of available languages in description of AI ASR task
//     creation.
//
// If several languages are specified for translation, a separate subtitle will be
// generated for each language.
//
// Example:
//
// ```
// auto_translate_subtitles_language: default
// auto_translate_subtitles_language: eng,fre,ger
// ```
//
// Please note that subtitle translation is done separately and after
// transcription. Thus separate AI-tasks are created for translation.
type PlaylistVideoAutoTranslateSubtitlesLanguage string

const (
	PlaylistVideoAutoTranslateSubtitlesLanguageDisable       PlaylistVideoAutoTranslateSubtitlesLanguage = "disable"
	PlaylistVideoAutoTranslateSubtitlesLanguageDefault       PlaylistVideoAutoTranslateSubtitlesLanguage = "default"
	PlaylistVideoAutoTranslateSubtitlesLanguageLanguageCodes PlaylistVideoAutoTranslateSubtitlesLanguage = "<language_codes,>"
)

type PlaylistNewParams struct {
	Playlist PlaylistParam
	paramObj
}

func (r PlaylistNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Playlist)
}
func (r *PlaylistNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Playlist)
}

type PlaylistUpdateParams struct {
	Playlist PlaylistParam
	paramObj
}

func (r PlaylistUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Playlist)
}
func (r *PlaylistUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Playlist)
}

type PlaylistListParams struct {
	// Query parameter. Use it to list the paginated content
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PlaylistListParams]'s query parameters as `url.Values`.
func (r PlaylistListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
