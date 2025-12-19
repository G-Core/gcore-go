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
//   - **Copy from another server** – If your video is accessable via "http://",
//     "https://", or "sftp://" public link, then you can use this method to copy a
//     file from an external server. Set `origin_url` parameter with the link to the
//     original video file (i.e. "https://domain.com/video.mp4"). After method
//     execution file will be uploaded and will be sent to transcoding automatically,
//     you don't have to do anything else. Use extra field `origin_http_headers` if
//     authorization is required on the external server.
//
//   - **Direct upload from a local device** – If you need to upload video directly
//     from your local device or from a mobile app, then use this method. Keep
//     `origin_url` empty and use TUS protocol ([tus.io](https://tus.io)) to upload
//     file. More details are here
//     ["Get TUS' upload"](/docs/api-reference/streaming/videos/get-tus-parameters-for-direct-upload)
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
//     ["AI Speech Recognition"](/docs/api-reference/streaming/ai/create-ai-asr-task).
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
//     ["add AI-generated subtitles to an exist video"](/docs/api-reference/streaming/subtitles/add-subtitle).
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
// Base documentation. By default video will be transcoded according to the
// original resolution, and a quality ladder suitable for your original video will
// be applied. There is no automatic upscaling; the maximum quality is taken from
// the original video. If you want to upload specific files not explicitly listed
// in requirements or wish to modify the standard quality ladder (i.e. decrease
// quality or add new non-standard qualities), then such customization is possible.
// Please reach out to us for assistance.
//
// Additionally, check the Knowledge Base for any supplementary information you may
// need.
func (r *VideoService) New(ctx context.Context, body VideoNewParams, opts ...option.RequestOption) (res *[]Video, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "streaming/videos"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes parameters of the video to new values.
//
// It's allowed to update only those public parameters that are described in POST
// method to create a new “video” entity. So it's not possible to change calculated
// parameters like "id", "duration", "`hls_url`", etc.
//
// Examples of changing:
//
// - Name: `{ "name": "new name of the video" }`
// - Move the video to a new directory: ` { "directory_id": 200 }`
//
// Please note that some parameters are used on initial step (before transcoding)
// only, so after transcoding there is no use in changing their values. For
// example, "`origin_url`" parameter is used for downloading an original file from
// a source and never used after transcoding; or "priority" parameter is used to
// set priority of processing and never used after transcoding.
func (r *VideoService) Update(ctx context.Context, videoID int64, body VideoUpdateParams, opts ...option.RequestOption) (res *Video, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/videos/%v", videoID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a set of videos by the given criteria.
func (r *VideoService) List(ctx context.Context, query VideoListParams, opts ...option.RequestOption) (res *pagination.PageStreaming[Video], err error) {
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
func (r *VideoService) ListAutoPaging(ctx context.Context, query VideoListParams, opts ...option.RequestOption) *pagination.PageStreamingAutoPager[Video] {
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
	return
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
	return
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
	return
}

// Use this method to get TUS' session parameters: hostname of the server to
// upload, secure token.
//
// The general sequence of actions for a direct upload of a video is as follows:
//
//   - Create video entity via POST method
//     ["Create video"](/docs/api-reference/streaming/videos/create-video)
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
	return
}

// Returns names for specified video IDs
func (r *VideoService) ListNames(ctx context.Context, query VideoListNamesParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "streaming/videos/names"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
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
	return json.Unmarshal(data, &r.CreateVideo)
}

type VideoListParams struct {
	// IDs of the videos to find. You can specify one or more identifiers separated by
	// commas. Example, ?id=1,101,1001
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Find videos where "`client_user_id`" meta field is equal to the search value
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
	// Find videos recorded from a specific stream, so for which "`stream_id`" field is
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
