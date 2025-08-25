// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package streaming

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	shimjson "github.com/G-Core/gcore-go/internal/encoding/json"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// VideoSubtitleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVideoSubtitleService] method instead.
type VideoSubtitleService struct {
	Options []option.RequestOption
}

// NewVideoSubtitleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVideoSubtitleService(opts ...option.RequestOption) (r VideoSubtitleService) {
	r = VideoSubtitleService{}
	r.Options = opts
	return
}

// Add new subtitle/captions to a video entity.
//
// **Add already exist subtitles** Subtitles must be in one of the following
// formats:
//
//   - SRT – SubRip Text is described on
//     [wikipedia.org](https://en.wikipedia.org/wiki/SubRip#SubRip_file_format). Must
//     start from integer for sequence number. Use calidators to check the subtitles,
//     like
//     [srt-validator](https://taoning2014.github.io/srt-validator-website/index.html).
//   - WebVTT – Web Video Text Tracks Format is described on
//     [developer.mozilla.org](https://developer.mozilla.org/en-US/docs/Web/API/WebVTT_API).
//     Must start from "WEBVTT" header. Use validators to check the subtitles, like
//     [W3C](https://w3c.github.io/webvtt.js/parser.html). Language is 3-letter
//     language code according to ISO-639-2 (bibliographic code). Specify language
//     you need, or just look at our list in the attribute "`audio_language`" of
//     section
//     ["AI Transcribe"](https://api.gcore.com/docs/streaming/docs/api-reference/streaming/ai/create-ai-asr-task).
//     You can add multiple subtitles in the same language, language uniqueness is
//     not required. Size must be up to 5Mb.
//
// The update time for added or changed subtitles is up to 30 seconds. Just like
// videos, subtitles are cached, so it takes time to update the data.
//
// **AI subtitles and transcribing** It is also possible to automatically create
// subtitles based on AI. Read more:
//
//   - What is
//     ["AI Transcribe"](https://api.gcore.com/docs/streaming/docs/api-reference/streaming/ai/create-ai-asr-task).
//   - If the option is enabled via
//     `auto_transcribe_audio_language: auto|<language_code>`, then immediately after
//     successful transcoding, an AI task will be automatically created for
//     transcription.
//   - If you need to translate subtitles from original language to any other, then
//     AI-task of subtitles translation can be applied. Use
//     `auto_translate_subtitles_language: default|<language_codes,>` parameter for
//     that. Also you can point several languages to translate to, then a separate
//     subtitle will be generated for each specified language. The created AI-task(s)
//     will be automatically executed, and result will also be automatically attached
//     to this video as subtitle(s). If AI is disabled in your account, you will
//     receive code 422 in response.
//
// **Where and how subtitles are displayed?** Subtitles are became available in the
// API response and in playback manifests. All added subtitles are automatically
// inserted into the output manifest .m3u8. This way, subtitles become available to
// any player: our player, OS built-in, or other specialized ones. You don't need
// to do anything else. Read more information in the Knowledge Base. Example:
//
// ```
// #EXT-X-MEDIA:TYPE=SUBTITLES,GROUP-ID="subs0",NAME="English",LANGUAGE="en",AUTOSELECT=YES,URI="subs-0.m3u8"
// ```
//
// ![Auto generated subtitles example](https://demo-files.gvideo.io/apidocs/captions.gif)
func (r *VideoSubtitleService) New(ctx context.Context, videoID int64, body VideoSubtitleNewParams, opts ...option.RequestOption) (res *Subtitle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("streaming/videos/%v/subtitles", videoID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Method to update subtitle of a video. You can update all or only some of fields
// you need. If you want to replace the text of subtitles (i.e. found a typo in the
// text, or the timing in the video changed), then:
//
//   - download it using GET method,
//   - change it in an external editor,
//   - and update it using this PATCH method. Just like videos, subtitles are cached,
//     so it takes time to update the data. See POST method for details.
func (r *VideoSubtitleService) Update(ctx context.Context, id int64, params VideoSubtitleUpdateParams, opts ...option.RequestOption) (res *SubtitleBase, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("streaming/videos/%v/subtitles/%v", params.VideoID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Method returns a list of all subtitles that are already attached to a video.
func (r *VideoSubtitleService) List(ctx context.Context, videoID int64, opts ...option.RequestOption) (res *[]Subtitle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("streaming/videos/%v/subtitles", videoID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete specified video subtitle
func (r *VideoSubtitleService) Delete(ctx context.Context, id int64, body VideoSubtitleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("streaming/videos/%v/subtitles/%v", body.VideoID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns information about a specific subtitle for a video.
func (r *VideoSubtitleService) Get(ctx context.Context, id int64, query VideoSubtitleGetParams, opts ...option.RequestOption) (res *Subtitle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("streaming/videos/%v/subtitles/%v", query.VideoID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type VideoSubtitleNewParams struct {
	Body VideoSubtitleNewParamsBody
	paramObj
}

func (r VideoSubtitleNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *VideoSubtitleNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type VideoSubtitleNewParamsBody struct {
	SubtitleBaseParam
}

func (r VideoSubtitleNewParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow VideoSubtitleNewParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}

type VideoSubtitleUpdateParams struct {
	VideoID      int64 `path:"video_id,required" json:"-"`
	SubtitleBase SubtitleBaseParam
	paramObj
}

func (r VideoSubtitleUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SubtitleBase)
}
func (r *VideoSubtitleUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SubtitleBase)
}

type VideoSubtitleDeleteParams struct {
	VideoID int64 `path:"video_id,required" json:"-"`
	paramObj
}

type VideoSubtitleGetParams struct {
	VideoID int64 `path:"video_id,required" json:"-"`
	paramObj
}
