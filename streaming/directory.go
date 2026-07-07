// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package streaming

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// DirectoryService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDirectoryService] method instead.
type DirectoryService struct {
	Options []option.RequestOption
}

// NewDirectoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDirectoryService(opts ...option.RequestOption) (r DirectoryService) {
	r = DirectoryService{}
	r.Options = opts
	return
}

// Use this method to create a new directory entity.
func (r *DirectoryService) New(ctx context.Context, body DirectoryNewParams, opts ...option.RequestOption) (res *DirectoryBase, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "streaming/directories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Change a directory name or move to another "parent_id".
func (r *DirectoryService) Update(ctx context.Context, directoryID int64, body DirectoryUpdateParams, opts ...option.RequestOption) (res *DirectoryBase, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/directories/%v", directoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a directory **and all entities inside**.
//
// After its execution, all contents of the directory will be deleted recursively:
//
// - Subdirectories
// - Videos
//
// The directory and contents are deleted permanently and irreversibly. Therefore,
// it is impossible to restore files after this.
//
// For details, see the Product Documentation.
func (r *DirectoryService) Delete(ctx context.Context, directoryID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("streaming/directories/%v", directoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Complete directory structure with contents. The structure contains both
// subfolders and videos in a continuous list.
func (r *DirectoryService) Get(ctx context.Context, directoryID int64, opts ...option.RequestOption) (res *DirectoryGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/directories/%v", directoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Tree structure of directories.
//
// This endpoint returns hierarchical data about directories in video hosting.
func (r *DirectoryService) GetTree(ctx context.Context, opts ...option.RequestOption) (res *DirectoriesTree, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "streaming/directories/tree"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type DirectoriesTree struct {
	Tree []DirectoriesTreeTree `json:"tree"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tree        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectoriesTree) RawJSON() string { return r.JSON.raw }
func (r *DirectoriesTree) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoriesTreeTree struct {
	// Array of subdirectories, if any.
	Descendants []DirectoriesTree `json:"descendants"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Descendants respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	DirectoryBase
}

// Returns the unmodified JSON received from the API
func (r DirectoriesTreeTree) RawJSON() string { return r.JSON.raw }
func (r *DirectoriesTreeTree) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryBase struct {
	// ID of the directory
	ID int64 `json:"id"`
	// Time of creation. Datetime in ISO 8601 format.
	CreatedAt string `json:"created_at"`
	// Number of objects in this directory. Counting files and folders. The quantity is
	// calculated only at one level (not recursively in all subfolders).
	ItemsCount int64 `json:"items_count"`
	// Title of the directory
	Name string `json:"name"`
	// ID of a parent directory. "null" if it's in the root.
	ParentID int64 `json:"parent_id"`
	// Time of last update of the directory entity. Datetime in ISO 8601 format.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ItemsCount  respjson.Field
		Name        respjson.Field
		ParentID    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectoryBase) RawJSON() string { return r.JSON.raw }
func (r *DirectoryBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryItem struct {
	// Type of the entity: directory, or video
	//
	// Any of "Directory".
	ItemType string `json:"item_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	DirectoryBase
}

// Returns the unmodified JSON received from the API
func (r DirectoryItem) RawJSON() string { return r.JSON.raw }
func (r *DirectoryItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryVideo struct {
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
	ConvertedVideos []DirectoryVideoConvertedVideo `json:"converted_videos"`
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
	// Type of the entity: directory, or video
	//
	// Any of "Video".
	ItemType DirectoryVideoItemType `json:"item_type"`
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
	RecordType DirectoryVideoRecordType `json:"record_type" api:"nullable"`
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
	Status DirectoryVideoStatus `json:"status"`
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
		ItemType            respjson.Field
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
func (r DirectoryVideo) RawJSON() string { return r.JSON.raw }
func (r *DirectoryVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryVideoConvertedVideo struct {
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
func (r DirectoryVideoConvertedVideo) RawJSON() string { return r.JSON.raw }
func (r *DirectoryVideoConvertedVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the entity: directory, or video
type DirectoryVideoItemType string

const (
	DirectoryVideoItemTypeVideo DirectoryVideoItemType = "Video"
)

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
type DirectoryVideoRecordType string

const (
	DirectoryVideoRecordTypeOrigin     DirectoryVideoRecordType = "origin"
	DirectoryVideoRecordTypeTranscoded DirectoryVideoRecordType = "transcoded"
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
type DirectoryVideoStatus string

const (
	DirectoryVideoStatusEmpty    DirectoryVideoStatus = "empty"
	DirectoryVideoStatusPending  DirectoryVideoStatus = "pending"
	DirectoryVideoStatusViewable DirectoryVideoStatus = "viewable"
	DirectoryVideoStatusReady    DirectoryVideoStatus = "ready"
	DirectoryVideoStatusError    DirectoryVideoStatus = "error"
)

type DirectoryGetResponse struct {
	Directory DirectoryGetResponseDirectory `json:"directory"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Directory   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectoryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DirectoryGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryGetResponseDirectory struct {
	// Array of subdirectories, if any.
	Items []DirectoryGetResponseDirectoryItemUnion `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	DirectoryBase
}

// Returns the unmodified JSON received from the API
func (r DirectoryGetResponseDirectory) RawJSON() string { return r.JSON.raw }
func (r *DirectoryGetResponseDirectory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DirectoryGetResponseDirectoryItemUnion contains all possible properties and
// values from [DirectoryVideo], [DirectoryItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DirectoryGetResponseDirectoryItemUnion struct {
	ID int64 `json:"id"`
	// This field is from variant [DirectoryVideo].
	AdID int64 `json:"ad_id"`
	// This field is from variant [DirectoryVideo].
	CDNViews int64 `json:"cdn_views"`
	// This field is from variant [DirectoryVideo].
	ClientID int64 `json:"client_id"`
	// This field is from variant [DirectoryVideo].
	ClientUserID int64 `json:"client_user_id"`
	// This field is from variant [DirectoryVideo].
	ConvertedVideos []DirectoryVideoConvertedVideo `json:"converted_videos"`
	CreatedAt       string                         `json:"created_at"`
	// This field is from variant [DirectoryVideo].
	CustomIframeURL string `json:"custom_iframe_url"`
	// This field is from variant [DirectoryVideo].
	DashURL string `json:"dash_url"`
	// This field is from variant [DirectoryVideo].
	Description string `json:"description"`
	// This field is from variant [DirectoryVideo].
	DirectoryID int64 `json:"directory_id"`
	// This field is from variant [DirectoryVideo].
	Duration int64 `json:"duration"`
	// This field is from variant [DirectoryVideo].
	Error string `json:"error"`
	// This field is from variant [DirectoryVideo].
	HlsCmafURL string `json:"hls_cmaf_url"`
	// This field is from variant [DirectoryVideo].
	HlsURL   string `json:"hls_url"`
	ItemType string `json:"item_type"`
	Name     string `json:"name"`
	// This field is from variant [DirectoryVideo].
	OriginSize int64 `json:"origin_size"`
	// This field is from variant [DirectoryVideo].
	OriginURL string `json:"origin_url"`
	// This field is from variant [DirectoryVideo].
	OriginVideoDuration int64 `json:"origin_video_duration"`
	// This field is from variant [DirectoryVideo].
	Poster string `json:"poster"`
	// This field is from variant [DirectoryVideo].
	PosterThumb string `json:"poster_thumb"`
	// This field is from variant [DirectoryVideo].
	Priority int64 `json:"priority"`
	// This field is from variant [DirectoryVideo].
	Projection string `json:"projection"`
	// This field is from variant [DirectoryVideo].
	QualitySetID int64 `json:"quality_set_id"`
	// This field is from variant [DirectoryVideo].
	RecordType DirectoryVideoRecordType `json:"record_type"`
	// This field is from variant [DirectoryVideo].
	RecordingStartedAt string `json:"recording_started_at"`
	// This field is from variant [DirectoryVideo].
	Screenshot string `json:"screenshot"`
	// This field is from variant [DirectoryVideo].
	ScreenshotID int64 `json:"screenshot_id"`
	// This field is from variant [DirectoryVideo].
	Screenshots []string `json:"screenshots"`
	// This field is from variant [DirectoryVideo].
	ShareURL string `json:"share_url"`
	// This field is from variant [DirectoryVideo].
	Slug string `json:"slug"`
	// This field is from variant [DirectoryVideo].
	Status DirectoryVideoStatus `json:"status"`
	// This field is from variant [DirectoryVideo].
	StreamID  int64  `json:"stream_id"`
	UpdatedAt string `json:"updated_at"`
	// This field is from variant [DirectoryVideo].
	Views int64 `json:"views"`
	// This field is from variant [DirectoryItem].
	ItemsCount int64 `json:"items_count"`
	// This field is from variant [DirectoryItem].
	ParentID int64 `json:"parent_id"`
	JSON     struct {
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
		ItemType            respjson.Field
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
		ItemsCount          respjson.Field
		ParentID            respjson.Field
		raw                 string
	} `json:"-"`
}

func (u DirectoryGetResponseDirectoryItemUnion) AsDirectoryVideo() (v DirectoryVideo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DirectoryGetResponseDirectoryItemUnion) AsDirectoryItem() (v DirectoryItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DirectoryGetResponseDirectoryItemUnion) RawJSON() string { return u.JSON.raw }

func (r *DirectoryGetResponseDirectoryItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryNewParams struct {
	// Title of the directory.
	Name string `json:"name" api:"required"`
	// ID of a parent directory. "null" if it's in the root.
	ParentID param.Opt[int64] `json:"parent_id,omitzero"`
	paramObj
}

func (r DirectoryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DirectoryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirectoryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryUpdateParams struct {
	// Title of the directory. Omit this if you don't want to change.
	Name param.Opt[string] `json:"name,omitzero"`
	// ID of a parent directory. "null" if it's in the root. Omit this if you don't
	// want to change.
	ParentID param.Opt[int64] `json:"parent_id,omitzero"`
	paramObj
}

func (r DirectoryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DirectoryUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirectoryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
