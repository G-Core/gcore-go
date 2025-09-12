// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package streaming

import (
	"context"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// QualitySetService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQualitySetService] method instead.
type QualitySetService struct {
	Options []option.RequestOption
}

// NewQualitySetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQualitySetService(opts ...option.RequestOption) (r QualitySetService) {
	r = QualitySetService{}
	r.Options = opts
	return
}

// Method returns a list of all custom quality sets.
//
// Transcoding is designed to minimize video file size while maintaining maximum
// visual quality. This is done so that the video can be delivered and viewed on
// any device, on any Internet connection, anywhere in the world. It's always a
// compromise between video/audio quality and delivery+viewing quality (QoE).
//
// Our experts have selected the optimal parameters for transcoding, to ensure
// maximum video/audio quality with the best compression. Default quality sets are
// described in the
// [documentation](https://gcore.com/docs/streaming-platform/live-streams-and-videos-protocols-and-codecs/output-parameters-after-transcoding-bitrate-frame-rate-and-codecs).
// These values are the default for everyone. There is no need to configure
// anything additional. Read more about qiality in our blog
// [How we lowered the bitrate for live and VOD streaming by 32.5% without sacrificing quality](https://gcore.com/blog/how-we-lowered-the-bitrate-for-live-and-vod-streaming-by-32-5-without-sacrificing-quality).
// ![Quality ladder](https://demo-files.gvideo.io/apidocs/encoding_ladder.png) Only
// for those cases when, in addition to the main parameters, it is necessary to use
// your own, then it is necessary to use custom quality sets. How to use:
//
//  1. By default custom quality set is empty – `{ "live":[],"vod":[] }`
//  2. Request the use of custom quality sets from your manager or the Support Team.
//  3. Please forward your requirements to us, since the parameters are set not by
//     you, but by our engineers. (We are working to ensure that later you can
//     create qualities by yourself.)
//  4. Use the created quality sets through the these specified API methods. Here
//     are some common parameters of quality settings:
//
//   - Resolution: Determines the size of the video frame. I.e. 720p, 1080p, 4K, etc.
//   - Bitrate: Refers to the amount of data processed per unit of time.
//   - Codec: Codec used for transcoding can significantly affect quality. Popular
//     codecs include H.264 (AVC), H.265 (HEVC), and AV1.
//   - Frame Rate: Determines how many frames per second are displayed. Common frame
//     rates include 24fps, 30fps, and 60fps.
//   - Color Depth and Chroma Subsampling: These settings determine the accuracy of
//     color representation in the video.
//   - Audio Bitrate and Codec: Don't forget about the audio :) Bitrate and codec
//     used for audio can also affect the overall quality. Note: Custom quality set
//     is a paid feature.
func (r *QualitySetService) List(ctx context.Context, opts ...option.RequestOption) (res *QualitySets, err error) {
	opts = append(r.Options[:], opts...)
	path := "streaming/quality_sets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Method to set default quality set for VOD and Live transcoding. For changing
// default quality set, specify the ID of the custom quality set from the method
// GET /`quality_sets`. Default value can be reverted to the system defaults
// (cleared) by setting `"id": null`.
//
// Live transcoding management:
//
//   - You can specify quality set explicitly in POST /streams method, look at
//     attribute "`quality_set_id`".
//   - Otherwise these default values will be used by the system by default. VOD
//     transcoding management:
//   - You can specify quality set explicitly in POST /videos method, look at
//     attribute "`quality_set_id`".
//   - Otherwise these default values will be used by the system by default.
func (r *QualitySetService) SetDefault(ctx context.Context, body QualitySetSetDefaultParams, opts ...option.RequestOption) (res *QualitySets, err error) {
	opts = append(r.Options[:], opts...)
	path := "streaming/quality_sets/default"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type QualitySets struct {
	Live []QualitySetsLive `json:"live"`
	Vod  []QualitySetsVod  `json:"vod"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Live        respjson.Field
		Vod         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QualitySets) RawJSON() string { return r.JSON.raw }
func (r *QualitySets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QualitySetsLive struct {
	// ID of the custom quality set
	ID int64 `json:"id"`
	// States if this preset is default for a client profile
	Default bool `json:"default"`
	// Human readable name of the quality set
	Name string `json:"name"`
	// Array of associated qualities
	Qualities []QualitySetsLiveQuality `json:"qualities"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Default     respjson.Field
		Name        respjson.Field
		Qualities   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QualitySetsLive) RawJSON() string { return r.JSON.raw }
func (r *QualitySetsLive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QualitySetsLiveQuality struct {
	// ID of the quality
	ID int64 `json:"id"`
	// Name of the quality
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QualitySetsLiveQuality) RawJSON() string { return r.JSON.raw }
func (r *QualitySetsLiveQuality) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QualitySetsVod struct {
	// ID of the custom quality set
	ID int64 `json:"id"`
	// States if this preset is default for a client profile
	Default bool `json:"default"`
	// Human readable name of the quality set
	Name string `json:"name"`
	// Array of associated qualities
	Qualities []QualitySetsVodQuality `json:"qualities"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Default     respjson.Field
		Name        respjson.Field
		Qualities   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QualitySetsVod) RawJSON() string { return r.JSON.raw }
func (r *QualitySetsVod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QualitySetsVodQuality struct {
	// ID of the quality
	ID int64 `json:"id"`
	// Name of the quality
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QualitySetsVodQuality) RawJSON() string { return r.JSON.raw }
func (r *QualitySetsVodQuality) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QualitySetSetDefaultParams struct {
	Live QualitySetSetDefaultParamsLive `json:"live,omitzero"`
	Vod  QualitySetSetDefaultParamsVod  `json:"vod,omitzero"`
	paramObj
}

func (r QualitySetSetDefaultParams) MarshalJSON() (data []byte, err error) {
	type shadow QualitySetSetDefaultParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QualitySetSetDefaultParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QualitySetSetDefaultParamsLive struct {
	// ID of the custom quality set, or "null" for the system default
	ID param.Opt[int64] `json:"id,omitzero"`
	paramObj
}

func (r QualitySetSetDefaultParamsLive) MarshalJSON() (data []byte, err error) {
	type shadow QualitySetSetDefaultParamsLive
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QualitySetSetDefaultParamsLive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QualitySetSetDefaultParamsVod struct {
	// ID of the custom quality set, or "null" for the system default
	ID param.Opt[int64] `json:"id,omitzero"`
	paramObj
}

func (r QualitySetSetDefaultParamsVod) MarshalJSON() (data []byte, err error) {
	type shadow QualitySetSetDefaultParamsVod
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QualitySetSetDefaultParamsVod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
