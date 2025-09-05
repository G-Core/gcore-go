// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package streaming

import (
	"context"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// StatisticService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatisticService] method instead.
type StatisticService struct {
	Options []option.RequestOption
}

// NewStatisticService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStatisticService(opts ...option.RequestOption) (r StatisticService) {
	r = StatisticService{}
	r.Options = opts
	return
}

// Aggregates data for the specified video stream in the specified time interval.
// "interval" and "units" params working together to point aggregation interval.
// You can use this method to watch when stream was alive in time, and when it was
// off.
func (r *StatisticService) GetFfprobes(ctx context.Context, query StatisticGetFfprobesParams, opts ...option.RequestOption) (res *Ffprobes, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/ffprobe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Calculates time series of unique viewers of Live streams via CDN. The statistics
// are taken from the data of CDN and work regardless of which player the views
// were made with. Works similar to the method `/statistics/cdn/uniqs`. But this
// allows you to break down data with the specified granularity: minutes, hours,
// days. Based on this method, a graph of unique views in the Customer Portal is
// built.
// ![Unique viewers via CDN in Customer Portal](https://demo-files.gvideo.io/apidocs/cdn_unique_viewers.png)
func (r *StatisticService) GetLiveUniqueViewers(ctx context.Context, query StatisticGetLiveUniqueViewersParams, opts ...option.RequestOption) (res *[]StatisticGetLiveUniqueViewersResponse, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/stream/viewers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Calculates a time series of live streams watching duration in minutes. Views of
// only those streams that meet the specified filters are summed up. The statistics
// are taken from the data of CDN and work regardless of which player the views
// were made with. Please note that the result for each time interval is in
// minutes, it is rounded to the nearest upper integer. You cannot use the sum of
// all intervals as the total watch time value; instead, use the /total method.
func (r *StatisticService) GetLiveWatchTimeCdn(ctx context.Context, query StatisticGetLiveWatchTimeCdnParams, opts ...option.RequestOption) (res *StreamSeries, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/stream/watching_duration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Calculates the total duration of live streams watching in minutes. Views of only
// those streams that meet the specified filters are summed up. The statistics are
// taken from the data of CDN and work regardless of which player the views were
// made with.
func (r *StatisticService) GetLiveWatchTimeTotalCdn(ctx context.Context, query StatisticGetLiveWatchTimeTotalCdnParams, opts ...option.RequestOption) (res *VodTotalStreamDurationSeries, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/stream/watching_duration/total"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Calculates time series of the amount of simultaneous streams. The data is
// updated near realtime.
func (r *StatisticService) GetMaxStreamsSeries(ctx context.Context, query StatisticGetMaxStreamsSeriesParams, opts ...option.RequestOption) (res *MaxStreamSeries, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/max_stream"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Aggregates the number of views for all client videos, grouping them by id and
// sort from most popular to less in the built-in player. Note. This method
// operates only on data collected by the built-in HTML player. It will not show
// statistics if you are using another player or viewing in native OS players
// through direct .m3u8/.mpd/.mp4 links. For such cases, use calculations through
// CDN (look at method /statistics/cdn/uniqs) or statistics of the players you have
// chosen.
func (r *StatisticService) GetPopularVideos(ctx context.Context, query StatisticGetPopularVideosParams, opts ...option.RequestOption) (res *PopularVideos, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/popular"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Calculates time series of the size of disk space in bytes for all processed and
// undeleted client videos. The data is updated every 8 hours, it does not make
// sense to set the granulation less than 1 day.
func (r *StatisticService) GetStorageSeries(ctx context.Context, query StatisticGetStorageSeriesParams, opts ...option.RequestOption) (res *StorageSeries, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/storage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Calculates time series of the transcoding minutes of all streams. The data is
// updated near realtime.
func (r *StatisticService) GetStreamSeries(ctx context.Context, query StatisticGetStreamSeriesParams, opts ...option.RequestOption) (res *StreamSeries, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/stream"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the number of unique viewers in the built-in player. Counts the number of
// unique IPs. Allows flexible grouping and filtering. The fields in the response
// depend on the selected grouping. Note. This method operates only on data
// collected by the built-in HTML player. It will not show statistics if you are
// using another player or viewing in native OS players through direct
// .m3u8/.mpd/.mp4 links. For such cases, use calculations through CDN (look at
// method /statistics/cdn/uniqs) or statistics of the players you have chosen.
func (r *StatisticService) GetUniqueViewers(ctx context.Context, query StatisticGetUniqueViewersParams, opts ...option.RequestOption) (res *UniqueViewers, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/uniqs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Сounts the number of unique viewers of a video entity over CDN. It doesn't
// matter what player you used. All unique viewers for the specified period of time
// are counted. **How does it work?** Calculating the number of unique viewers for
// a Live stream or VOD over CDN involves aggregating and analyzing various metrics
// to ensure each individual viewer is counted only once, regardless of how many
// times they connect or disconnect during the stream. This method provides
// statistics for any video viewing by unique users, regardless of viewing method
// and a player you used. Thus, this is the most important difference from viewing
// through the built-in player:
//
//   - In method /statistics/uniqs viewers of the built-in player are tracked only.
//   - But this method tracks all viewers from everywhere. This method is a
//     combination of two other Live and VOD detailed methods. If you need detailed
//     information, then see the methods: `/statistics/stream/viewers` and
//     `/statistics/vod/viewers`. **Data Processing and Deduplication** We us IP
//     Address & User-Agent combination. Each unique combination of IP address and
//     User-Agent string might be considered a unique viewer. This approach allows to
//     accurately estimate the number of unique viewers. However, this is not
//     foolproof due to NAT (Network Address Translation) and shared networks. Thus
//     if your users fall under such restrictions, then the number of unique viewers
//     may be higher than calculated. **Why is there no "Unique Views" method?**
//     Based on CDN data, we can calculate the number of unique viewers only. Thus
//     only your player will be able to count the number of unique views (clicks on
//     the Play button) within the player session (i.e. how many times 1 unique
//     viewer clicked the Play button within a unique player's session).
func (r *StatisticService) GetUniqueViewersCdn(ctx context.Context, query StatisticGetUniqueViewersCdnParams, opts ...option.RequestOption) (res *UniqueViewersCdn, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/cdn/uniqs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the number of views in the built-in player. Allows flexible grouping and
// filtering. The fields in the response depend on the selected grouping. Note.
// This method operates only on data collected by the built-in HTML player. It will
// not show statistics if you are using another player or viewing in native OS
// players through direct .m3u8/.mpd/.mp4 links. For such cases, use calculations
// through CDN (look at method /statistics/cdn/uniqs) or statistics of the players
// you have chosen.
func (r *StatisticService) GetViews(ctx context.Context, query StatisticGetViewsParams, opts ...option.RequestOption) (res *Views, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/views"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Aggregates the number of views for all client videos, grouping them by browsers
// in the built-in player. Note. This method operates only on data collected by the
// built-in HTML player. It will not show statistics if you are using another
// player or viewing in native OS players through direct .m3u8/.mpd/.mp4 links. For
// such cases, use calculations through CDN (look at method /statistics/cdn/uniqs)
// or statistics of the players you have chosen.
func (r *StatisticService) GetViewsByBrowsers(ctx context.Context, query StatisticGetViewsByBrowsersParams, opts ...option.RequestOption) (res *ViewsByBrowser, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/browsers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Aggregates the number of views grouping them by country in the built-in player.
// Note. This method operates only on data collected by the built-in HTML player.
// It will not show statistics if you are using another player or viewing in native
// OS players through direct .m3u8/.mpd/.mp4 links. For such cases, use
// calculations through CDN (look at method /statistics/cdn/uniqs) or statistics of
// the players you have chosen.
func (r *StatisticService) GetViewsByCountry(ctx context.Context, query StatisticGetViewsByCountryParams, opts ...option.RequestOption) (res *ViewsByCountry, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/countries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Aggregates the number of views, grouping them by "host" domain name the built-in
// player was embeded to. Note. This method operates only on data collected by the
// built-in HTML player. It will not show statistics if you are using another
// player or viewing in native OS players through direct .m3u8/.mpd/.mp4 links. For
// such cases, use calculations through CDN (look at method /statistics/cdn/uniqs)
// or statistics of the players you have chosen.
func (r *StatisticService) GetViewsByHostname(ctx context.Context, query StatisticGetViewsByHostnameParams, opts ...option.RequestOption) (res *ViewsByHostname, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/hosts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Aggregates the number of views for all client videos, grouping them by device
// OSs in the built-in player. Note. This method operates only on data collected by
// the built-in HTML player. It will not show statistics if you are using another
// player or viewing in native OS players through direct .m3u8/.mpd/.mp4 links. For
// such cases, use calculations through CDN (look at method /statistics/cdn/uniqs)
// or statistics of the players you have chosen.
func (r *StatisticService) GetViewsByOperatingSystem(ctx context.Context, query StatisticGetViewsByOperatingSystemParams, opts ...option.RequestOption) (res *ViewsByOperatingSystem, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/systems"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Aggregates the number of views, grouping them by "referer" URL of pages the
// built-in player was embeded to. Note. This method operates only on data
// collected by the built-in HTML player. It will not show statistics if you are
// using another player or viewing in native OS players through direct
// .m3u8/.mpd/.mp4 links. For such cases, use calculations through CDN (look at
// method /statistics/cdn/uniqs) or statistics of the players you have chosen.
func (r *StatisticService) GetViewsByReferer(ctx context.Context, query StatisticGetViewsByRefererParams, opts ...option.RequestOption) (res *ViewsByReferer, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/embeds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Aggregates the number of views grouping them by regions of countries in the
// built-in player. Note. This method operates only on data collected by the
// built-in HTML player. It will not show statistics if you are using another
// player or viewing in native OS players through direct .m3u8/.mpd/.mp4 links. For
// such cases, use calculations through CDN (look at method /statistics/cdn/uniqs)
// or statistics of the players you have chosen.
func (r *StatisticService) GetViewsByRegion(ctx context.Context, query StatisticGetViewsByRegionParams, opts ...option.RequestOption) (res *ViewsByRegion, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/regions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Shows information about what part of the video your viewers watched in the
// built-in player. This way you can find out how many viewers started watching the
// video, and where they stopped watching instead of watching the entire video to
// the end. Has different format of response depends on query param "type". Note.
// This method operates only on data collected by the built-in HTML player. It will
// not show statistics if you are using another player or viewing in native OS
// players through direct .m3u8/.mpd/.mp4 links. For such cases, use calculations
// through CDN (look at method /statistics/cdn/uniqs) or statistics of the players
// you have chosen.
func (r *StatisticService) GetViewsHeatmap(ctx context.Context, query StatisticGetViewsHeatmapParams, opts ...option.RequestOption) (res *ViewsHeatmap, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/heatmap"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Calculates time series of the duration in minutes for all processed and
// undeleted client videos. The data is updated every 8 hours, it does not make
// sense to set the granulation less than 1 day.
func (r *StatisticService) GetVodStorageVolume(ctx context.Context, query StatisticGetVodStorageVolumeParams, opts ...option.RequestOption) (res *VodStatisticsSeries, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/vod/storage_duration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Calculates time series of the transcoding time in minutes for all processed
// client videos. The data is updated every 8 hours, it does not make sense to set
// the granulation less than 1 day.
func (r *StatisticService) GetVodTranscodingDuration(ctx context.Context, query StatisticGetVodTranscodingDurationParams, opts ...option.RequestOption) (res *VodStatisticsSeries, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/vod/transcoding_duration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Calculates time series of unique viewers of VOD via CDN. The statistics are
// taken from the data of CDN and work regardless of which player the views were
// made with. Works similar to the method `/statistics/cdn/uniqs`. But this allows
// you to break down data with the specified granularity: minutes, hours, days.
// Based on this method, a graph of unique views in the Customer Portal is built.
// ![Unique viewers via CDN in Customer Portal](https://demo-files.gvideo.io/apidocs/cdn_unique_viewers.png)
func (r *StatisticService) GetVodUniqueViewersCdn(ctx context.Context, query StatisticGetVodUniqueViewersCdnParams, opts ...option.RequestOption) (res *VodStatisticsSeries, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/vod/viewers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Calculates a time series of video watching duration in minutes. Views of only
// those videos that meet the specified filters are summed up. The statistics are
// taken from the data of CDN and work regardless of which player the views were
// made with. Please note that the result for each time interval is in minutes, it
// is rounded to the nearest upper integer. You cannot use the sum of all intervals
// as the total watch time value; instead, use the /total method.
func (r *StatisticService) GetVodWatchTimeCdn(ctx context.Context, query StatisticGetVodWatchTimeCdnParams, opts ...option.RequestOption) (res *VodStatisticsSeries, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/vod/watching_duration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Calculates the total duration of video watching in minutes. Views of only those
// videos that meet the specified filters are summed up. The statistics are taken
// from the data of CDN and work regardless of which player the views were made
// with.
func (r *StatisticService) GetVodWatchTimeTotalCdn(ctx context.Context, query StatisticGetVodWatchTimeTotalCdnParams, opts ...option.RequestOption) (res *[]StatisticGetVodWatchTimeTotalCdnResponse, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/statistics/vod/watching_duration/total"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Ffprobes struct {
	Data []FfprobesData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Ffprobes) RawJSON() string { return r.JSON.raw }
func (r *Ffprobes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FfprobesData struct {
	AvgBitrate          float64 `json:"avg_bitrate,required"`
	MaxFps              float64 `json:"max_fps,required"`
	MaxHeight           int64   `json:"max_height,required"`
	MaxKeyframeInterval int64   `json:"max_keyframe_interval,required"`
	SumFrames           int64   `json:"sum_frames,required"`
	Time                string  `json:"time,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgBitrate          respjson.Field
		MaxFps              respjson.Field
		MaxHeight           respjson.Field
		MaxKeyframeInterval respjson.Field
		SumFrames           respjson.Field
		Time                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FfprobesData) RawJSON() string { return r.JSON.raw }
func (r *FfprobesData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MaxStreamSeries []MaxStreamSeriesItem

type MaxStreamSeriesItem struct {
	Client  int64                      `json:"client,required"`
	Metrics MaxStreamSeriesItemMetrics `json:"metrics,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Client      respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MaxStreamSeriesItem) RawJSON() string { return r.JSON.raw }
func (r *MaxStreamSeriesItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MaxStreamSeriesItemMetrics struct {
	Streams []int64 `json:"streams,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Streams     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MaxStreamSeriesItemMetrics) RawJSON() string { return r.JSON.raw }
func (r *MaxStreamSeriesItemMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PopularVideos struct {
	Data []PopularVideosData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PopularVideos) RawJSON() string { return r.JSON.raw }
func (r *PopularVideos) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PopularVideosData struct {
	ID    string `json:"id,required"`
	Views int64  `json:"views,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PopularVideosData) RawJSON() string { return r.JSON.raw }
func (r *PopularVideosData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageSeries []StorageSeriesItem

type StorageSeriesItem struct {
	Client  int64                    `json:"client,required"`
	Metrics StorageSeriesItemMetrics `json:"metrics,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Client      respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageSeriesItem) RawJSON() string { return r.JSON.raw }
func (r *StorageSeriesItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageSeriesItemMetrics struct {
	MaxVolumeUsage []int64   `json:"max_volume_usage,required"`
	Storage        [][]int64 `json:"storage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxVolumeUsage respjson.Field
		Storage        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageSeriesItemMetrics) RawJSON() string { return r.JSON.raw }
func (r *StorageSeriesItemMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamSeries []StreamSeriesItem

type StreamSeriesItem struct {
	Client  int64                   `json:"client,required"`
	Metrics StreamSeriesItemMetrics `json:"metrics,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Client      respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamSeriesItem) RawJSON() string { return r.JSON.raw }
func (r *StreamSeriesItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamSeriesItemMetrics struct {
	Streams []int64 `json:"streams,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Streams     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamSeriesItemMetrics) RawJSON() string { return r.JSON.raw }
func (r *StreamSeriesItemMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UniqueViewers struct {
	Data []UniqueViewersData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UniqueViewers) RawJSON() string { return r.JSON.raw }
func (r *UniqueViewers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UniqueViewersData struct {
	Date      string `json:"date,required"`
	UniqueIPs int64  `json:"unique_ips,required"`
	ID        int64  `json:"id"`
	Browser   string `json:"browser"`
	Country   string `json:"country"`
	Event     string `json:"event"`
	Host      string `json:"host"`
	IP        string `json:"ip"`
	Os        string `json:"os"`
	Platform  string `json:"platform"`
	Type      string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		UniqueIPs   respjson.Field
		ID          respjson.Field
		Browser     respjson.Field
		Country     respjson.Field
		Event       respjson.Field
		Host        respjson.Field
		IP          respjson.Field
		Os          respjson.Field
		Platform    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UniqueViewersData) RawJSON() string { return r.JSON.raw }
func (r *UniqueViewersData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UniqueViewersCdn struct {
	Data []UniqueViewersCdnData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UniqueViewersCdn) RawJSON() string { return r.JSON.raw }
func (r *UniqueViewersCdn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UniqueViewersCdnData struct {
	Type  string `json:"type,required"`
	Uniqs int64  `json:"uniqs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Uniqs       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UniqueViewersCdnData) RawJSON() string { return r.JSON.raw }
func (r *UniqueViewersCdnData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Views struct {
	Data []ViewsData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Views) RawJSON() string { return r.JSON.raw }
func (r *Views) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsData struct {
	Date     string `json:"date,required"`
	Type     string `json:"type,required"`
	Views    int64  `json:"views,required"`
	ID       int64  `json:"id"`
	Browser  string `json:"browser"`
	Country  string `json:"country"`
	Event    string `json:"event"`
	Host     string `json:"host"`
	IP       string `json:"ip"`
	Os       string `json:"os"`
	Platform string `json:"platform"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Type        respjson.Field
		Views       respjson.Field
		ID          respjson.Field
		Browser     respjson.Field
		Country     respjson.Field
		Event       respjson.Field
		Host        respjson.Field
		IP          respjson.Field
		Os          respjson.Field
		Platform    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsData) RawJSON() string { return r.JSON.raw }
func (r *ViewsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsByBrowser struct {
	Data []ViewsByBrowserData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsByBrowser) RawJSON() string { return r.JSON.raw }
func (r *ViewsByBrowser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsByBrowserData struct {
	Browser string `json:"browser,required"`
	Views   int64  `json:"views,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Browser     respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsByBrowserData) RawJSON() string { return r.JSON.raw }
func (r *ViewsByBrowserData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsByCountry struct {
	Data []ViewsByCountryData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsByCountry) RawJSON() string { return r.JSON.raw }
func (r *ViewsByCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsByCountryData struct {
	Country     string `json:"country,required"`
	CountryName string `json:"country_name,required"`
	Views       int64  `json:"views,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		CountryName respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsByCountryData) RawJSON() string { return r.JSON.raw }
func (r *ViewsByCountryData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsByHostname struct {
	Data []ViewsByHostnameData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsByHostname) RawJSON() string { return r.JSON.raw }
func (r *ViewsByHostname) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsByHostnameData struct {
	Host  string `json:"host,required"`
	Views int64  `json:"views,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Host        respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsByHostnameData) RawJSON() string { return r.JSON.raw }
func (r *ViewsByHostnameData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsByOperatingSystem struct {
	Data []ViewsByOperatingSystemData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsByOperatingSystem) RawJSON() string { return r.JSON.raw }
func (r *ViewsByOperatingSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsByOperatingSystemData struct {
	Os    string `json:"os,required"`
	Views int64  `json:"views,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Os          respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsByOperatingSystemData) RawJSON() string { return r.JSON.raw }
func (r *ViewsByOperatingSystemData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsByReferer struct {
	Data []ViewsByRefererData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsByReferer) RawJSON() string { return r.JSON.raw }
func (r *ViewsByReferer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsByRefererData struct {
	EmbedURL string `json:"embed_url,required"`
	Views    int64  `json:"views,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmbedURL    respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsByRefererData) RawJSON() string { return r.JSON.raw }
func (r *ViewsByRefererData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsByRegion struct {
	Data []ViewsByRegionData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsByRegion) RawJSON() string { return r.JSON.raw }
func (r *ViewsByRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsByRegionData struct {
	Region     string `json:"region,required"`
	RegionName string `json:"region_name,required"`
	Views      int64  `json:"views,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Region      respjson.Field
		RegionName  respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsByRegionData) RawJSON() string { return r.JSON.raw }
func (r *ViewsByRegionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsHeatmap struct {
	Data []ViewsHeatmapData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsHeatmap) RawJSON() string { return r.JSON.raw }
func (r *ViewsHeatmap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewsHeatmapData struct {
	Viewers int64  `json:"viewers,required"`
	Seconds int64  `json:"seconds"`
	Time    string `json:"time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Viewers     respjson.Field
		Seconds     respjson.Field
		Time        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewsHeatmapData) RawJSON() string { return r.JSON.raw }
func (r *ViewsHeatmapData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VodStatisticsSeries []VodStatisticsSeriesItem

type VodStatisticsSeriesItem struct {
	Client  int64                          `json:"client,required"`
	Metrics VodStatisticsSeriesItemMetrics `json:"metrics,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Client      respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VodStatisticsSeriesItem) RawJSON() string { return r.JSON.raw }
func (r *VodStatisticsSeriesItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VodStatisticsSeriesItemMetrics struct {
	Vod []int64 `json:"vod,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Vod         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VodStatisticsSeriesItemMetrics) RawJSON() string { return r.JSON.raw }
func (r *VodStatisticsSeriesItemMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VodTotalStreamDurationSeries []VodTotalStreamDurationSeriesItem

type VodTotalStreamDurationSeriesItem struct {
	Client int64 `json:"client,required"`
	// count of minutes
	Duration     int64  `json:"duration,required"`
	ClientUserID int64  `json:"client_user_id"`
	StreamID     string `json:"stream_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Client       respjson.Field
		Duration     respjson.Field
		ClientUserID respjson.Field
		StreamID     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VodTotalStreamDurationSeriesItem) RawJSON() string { return r.JSON.raw }
func (r *VodTotalStreamDurationSeriesItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatisticGetLiveUniqueViewersResponse struct {
	Client  int64                                        `json:"client,required"`
	Metrics StatisticGetLiveUniqueViewersResponseMetrics `json:"metrics,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Client      respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatisticGetLiveUniqueViewersResponse) RawJSON() string { return r.JSON.raw }
func (r *StatisticGetLiveUniqueViewersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatisticGetLiveUniqueViewersResponseMetrics struct {
	Streams []int64 `json:"streams,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Streams     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatisticGetLiveUniqueViewersResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *StatisticGetLiveUniqueViewersResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatisticGetVodWatchTimeTotalCdnResponse struct {
	Client int64 `json:"client,required"`
	// count of minutes
	Duration     int64  `json:"duration,required"`
	ClientUserID int64  `json:"client_user_id"`
	Slug         string `json:"slug"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Client       respjson.Field
		Duration     respjson.Field
		ClientUserID respjson.Field
		Slug         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatisticGetVodWatchTimeTotalCdnResponse) RawJSON() string { return r.JSON.raw }
func (r *StatisticGetVodWatchTimeTotalCdnResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatisticGetFfprobesParams struct {
	// Start of time frame. Format is ISO 8601.
	DateFrom string `query:"date_from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	DateTo string `query:"date_to,required" json:"-"`
	// Stream ID
	StreamID string           `query:"stream_id,required" json:"-"`
	Interval param.Opt[int64] `query:"interval,omitzero" json:"-"`
	// Any of "second", "minute", "hour", "day", "week", "month".
	Units StatisticGetFfprobesParamsUnits `query:"units,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetFfprobesParams]'s query parameters as
// `url.Values`.
func (r StatisticGetFfprobesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetFfprobesParamsUnits string

const (
	StatisticGetFfprobesParamsUnitsSecond StatisticGetFfprobesParamsUnits = "second"
	StatisticGetFfprobesParamsUnitsMinute StatisticGetFfprobesParamsUnits = "minute"
	StatisticGetFfprobesParamsUnitsHour   StatisticGetFfprobesParamsUnits = "hour"
	StatisticGetFfprobesParamsUnitsDay    StatisticGetFfprobesParamsUnits = "day"
	StatisticGetFfprobesParamsUnitsWeek   StatisticGetFfprobesParamsUnits = "week"
	StatisticGetFfprobesParamsUnitsMonth  StatisticGetFfprobesParamsUnits = "month"
)

type StatisticGetLiveUniqueViewersParams struct {
	// Start of time frame. Format is date time in ISO 8601
	From string `query:"from,required" json:"-"`
	// End of time frame. Format is date time in ISO 8601
	To string `query:"to,required" json:"-"`
	// Filter by "`client_user_id`"
	ClientUserID param.Opt[int64] `query:"client_user_id,omitzero" json:"-"`
	// Filter by "`stream_id`"
	StreamID param.Opt[int64] `query:"stream_id,omitzero" json:"-"`
	// Specifies the time interval for grouping data
	//
	// Any of "1m", "5m", "15m", "1h", "1d".
	Granularity StatisticGetLiveUniqueViewersParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetLiveUniqueViewersParams]'s query parameters as
// `url.Values`.
func (r StatisticGetLiveUniqueViewersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the time interval for grouping data
type StatisticGetLiveUniqueViewersParamsGranularity string

const (
	StatisticGetLiveUniqueViewersParamsGranularity1m  StatisticGetLiveUniqueViewersParamsGranularity = "1m"
	StatisticGetLiveUniqueViewersParamsGranularity5m  StatisticGetLiveUniqueViewersParamsGranularity = "5m"
	StatisticGetLiveUniqueViewersParamsGranularity15m StatisticGetLiveUniqueViewersParamsGranularity = "15m"
	StatisticGetLiveUniqueViewersParamsGranularity1h  StatisticGetLiveUniqueViewersParamsGranularity = "1h"
	StatisticGetLiveUniqueViewersParamsGranularity1d  StatisticGetLiveUniqueViewersParamsGranularity = "1d"
)

type StatisticGetLiveWatchTimeCdnParams struct {
	// Start of the time period for counting minutes of watching. Format is date time
	// in ISO 8601.
	From string `query:"from,required" json:"-"`
	// Filter by field "`client_user_id`"
	ClientUserID param.Opt[int64] `query:"client_user_id,omitzero" json:"-"`
	// Filter by `stream_id`
	StreamID param.Opt[int64] `query:"stream_id,omitzero" json:"-"`
	// End of time frame. Datetime in ISO 8601 format. If omitted, then the current
	// time is taken
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	// Data is grouped by the specified time interval
	//
	// Any of "1m", "5m", "15m", "1h", "1d", "1mo".
	Granularity StatisticGetLiveWatchTimeCdnParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetLiveWatchTimeCdnParams]'s query parameters as
// `url.Values`.
func (r StatisticGetLiveWatchTimeCdnParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Data is grouped by the specified time interval
type StatisticGetLiveWatchTimeCdnParamsGranularity string

const (
	StatisticGetLiveWatchTimeCdnParamsGranularity1m  StatisticGetLiveWatchTimeCdnParamsGranularity = "1m"
	StatisticGetLiveWatchTimeCdnParamsGranularity5m  StatisticGetLiveWatchTimeCdnParamsGranularity = "5m"
	StatisticGetLiveWatchTimeCdnParamsGranularity15m StatisticGetLiveWatchTimeCdnParamsGranularity = "15m"
	StatisticGetLiveWatchTimeCdnParamsGranularity1h  StatisticGetLiveWatchTimeCdnParamsGranularity = "1h"
	StatisticGetLiveWatchTimeCdnParamsGranularity1d  StatisticGetLiveWatchTimeCdnParamsGranularity = "1d"
	StatisticGetLiveWatchTimeCdnParamsGranularity1mo StatisticGetLiveWatchTimeCdnParamsGranularity = "1mo"
)

type StatisticGetLiveWatchTimeTotalCdnParams struct {
	// Filter by field "`client_user_id`"
	ClientUserID param.Opt[int64] `query:"client_user_id,omitzero" json:"-"`
	// Start of the time period for counting minutes of watching. Format is date time
	// in ISO 8601. If omitted, the earliest start time for viewing is taken
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// Filter by `stream_id`
	StreamID param.Opt[int64] `query:"stream_id,omitzero" json:"-"`
	// End of time frame. Datetime in ISO 8601 format. If missed, then the current time
	// is taken
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetLiveWatchTimeTotalCdnParams]'s query parameters
// as `url.Values`.
func (r StatisticGetLiveWatchTimeTotalCdnParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetMaxStreamsSeriesParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	From string `query:"from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	To string `query:"to,required" json:"-"`
	// specifies the time interval for grouping data
	//
	// Any of "1m", "5m", "15m", "1h", "1d".
	Granularity StatisticGetMaxStreamsSeriesParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetMaxStreamsSeriesParams]'s query parameters as
// `url.Values`.
func (r StatisticGetMaxStreamsSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// specifies the time interval for grouping data
type StatisticGetMaxStreamsSeriesParamsGranularity string

const (
	StatisticGetMaxStreamsSeriesParamsGranularity1m  StatisticGetMaxStreamsSeriesParamsGranularity = "1m"
	StatisticGetMaxStreamsSeriesParamsGranularity5m  StatisticGetMaxStreamsSeriesParamsGranularity = "5m"
	StatisticGetMaxStreamsSeriesParamsGranularity15m StatisticGetMaxStreamsSeriesParamsGranularity = "15m"
	StatisticGetMaxStreamsSeriesParamsGranularity1h  StatisticGetMaxStreamsSeriesParamsGranularity = "1h"
	StatisticGetMaxStreamsSeriesParamsGranularity1d  StatisticGetMaxStreamsSeriesParamsGranularity = "1d"
)

type StatisticGetPopularVideosParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	DateFrom string `query:"date_from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	DateTo string `query:"date_to,required" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetPopularVideosParams]'s query parameters as
// `url.Values`.
func (r StatisticGetPopularVideosParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetStorageSeriesParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	From string `query:"from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	To string `query:"to,required" json:"-"`
	// specifies the time interval for grouping data
	//
	// Any of "1m", "5m", "15m", "1h", "1d".
	Granularity StatisticGetStorageSeriesParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetStorageSeriesParams]'s query parameters as
// `url.Values`.
func (r StatisticGetStorageSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// specifies the time interval for grouping data
type StatisticGetStorageSeriesParamsGranularity string

const (
	StatisticGetStorageSeriesParamsGranularity1m  StatisticGetStorageSeriesParamsGranularity = "1m"
	StatisticGetStorageSeriesParamsGranularity5m  StatisticGetStorageSeriesParamsGranularity = "5m"
	StatisticGetStorageSeriesParamsGranularity15m StatisticGetStorageSeriesParamsGranularity = "15m"
	StatisticGetStorageSeriesParamsGranularity1h  StatisticGetStorageSeriesParamsGranularity = "1h"
	StatisticGetStorageSeriesParamsGranularity1d  StatisticGetStorageSeriesParamsGranularity = "1d"
)

type StatisticGetStreamSeriesParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	From string `query:"from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	To string `query:"to,required" json:"-"`
	// specifies the time interval for grouping data
	//
	// Any of "1m", "5m", "15m", "1h", "1d".
	Granularity StatisticGetStreamSeriesParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetStreamSeriesParams]'s query parameters as
// `url.Values`.
func (r StatisticGetStreamSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// specifies the time interval for grouping data
type StatisticGetStreamSeriesParamsGranularity string

const (
	StatisticGetStreamSeriesParamsGranularity1m  StatisticGetStreamSeriesParamsGranularity = "1m"
	StatisticGetStreamSeriesParamsGranularity5m  StatisticGetStreamSeriesParamsGranularity = "5m"
	StatisticGetStreamSeriesParamsGranularity15m StatisticGetStreamSeriesParamsGranularity = "15m"
	StatisticGetStreamSeriesParamsGranularity1h  StatisticGetStreamSeriesParamsGranularity = "1h"
	StatisticGetStreamSeriesParamsGranularity1d  StatisticGetStreamSeriesParamsGranularity = "1d"
)

type StatisticGetUniqueViewersParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	DateFrom string `query:"date_from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	DateTo string `query:"date_to,required" json:"-"`
	// filter by entity's id
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// filter by country
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// filter by host
	Host param.Opt[string] `query:"host,omitzero" json:"-"`
	// filter by event's name
	//
	// Any of "init", "start", "watch".
	Event StatisticGetUniqueViewersParamsEvent `query:"event,omitzero" json:"-"`
	// group=1,2,4 OR group=1&group=2&group=3
	//
	// Any of "date", "host", "os", "browser", "platform", "ip", "country", "event",
	// "id".
	Group []string `query:"group,omitzero" json:"-"`
	// filter by entity's type
	//
	// Any of "live", "vod", "playlist".
	Type StatisticGetUniqueViewersParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetUniqueViewersParams]'s query parameters as
// `url.Values`.
func (r StatisticGetUniqueViewersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// filter by event's name
type StatisticGetUniqueViewersParamsEvent string

const (
	StatisticGetUniqueViewersParamsEventInit  StatisticGetUniqueViewersParamsEvent = "init"
	StatisticGetUniqueViewersParamsEventStart StatisticGetUniqueViewersParamsEvent = "start"
	StatisticGetUniqueViewersParamsEventWatch StatisticGetUniqueViewersParamsEvent = "watch"
)

// filter by entity's type
type StatisticGetUniqueViewersParamsType string

const (
	StatisticGetUniqueViewersParamsTypeLive     StatisticGetUniqueViewersParamsType = "live"
	StatisticGetUniqueViewersParamsTypeVod      StatisticGetUniqueViewersParamsType = "vod"
	StatisticGetUniqueViewersParamsTypePlaylist StatisticGetUniqueViewersParamsType = "playlist"
)

type StatisticGetUniqueViewersCdnParams struct {
	// Start of time frame. Format is date time in ISO 8601.
	DateFrom string `query:"date_from,required" json:"-"`
	// End of time frame. Format is date time in ISO 8601.
	DateTo string `query:"date_to,required" json:"-"`
	// Filter by entity's id. Put ID of a Live stream, VOD or a playlist to be
	// calculated. If the value is omitted, then the calculation is done for all
	// videos/streams of the specified type. When using this "id" parameter, be sure to
	// specify the "type" parameter too. If you do not specify a type, the "id" will be
	// ignored.
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Filter by entity's type
	//
	// Any of "live", "vod", "playlist".
	Type StatisticGetUniqueViewersCdnParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetUniqueViewersCdnParams]'s query parameters as
// `url.Values`.
func (r StatisticGetUniqueViewersCdnParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by entity's type
type StatisticGetUniqueViewersCdnParamsType string

const (
	StatisticGetUniqueViewersCdnParamsTypeLive     StatisticGetUniqueViewersCdnParamsType = "live"
	StatisticGetUniqueViewersCdnParamsTypeVod      StatisticGetUniqueViewersCdnParamsType = "vod"
	StatisticGetUniqueViewersCdnParamsTypePlaylist StatisticGetUniqueViewersCdnParamsType = "playlist"
)

type StatisticGetViewsParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	DateFrom string `query:"date_from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	DateTo string `query:"date_to,required" json:"-"`
	// filter by entity's id
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// filter by country
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// filter by host
	Host param.Opt[string] `query:"host,omitzero" json:"-"`
	// filter by event's name
	//
	// Any of "init", "start", "watch".
	Event StatisticGetViewsParamsEvent `query:"event,omitzero" json:"-"`
	// group=1,2,4 OR group=1&group=2&group=3
	//
	// Any of "host", "os", "browser", "platform", "ip", "country", "event", "id".
	Group []string `query:"group,omitzero" json:"-"`
	// filter by entity's type
	//
	// Any of "live", "vod", "playlist".
	Type StatisticGetViewsParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetViewsParams]'s query parameters as
// `url.Values`.
func (r StatisticGetViewsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// filter by event's name
type StatisticGetViewsParamsEvent string

const (
	StatisticGetViewsParamsEventInit  StatisticGetViewsParamsEvent = "init"
	StatisticGetViewsParamsEventStart StatisticGetViewsParamsEvent = "start"
	StatisticGetViewsParamsEventWatch StatisticGetViewsParamsEvent = "watch"
)

// filter by entity's type
type StatisticGetViewsParamsType string

const (
	StatisticGetViewsParamsTypeLive     StatisticGetViewsParamsType = "live"
	StatisticGetViewsParamsTypeVod      StatisticGetViewsParamsType = "vod"
	StatisticGetViewsParamsTypePlaylist StatisticGetViewsParamsType = "playlist"
)

type StatisticGetViewsByBrowsersParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	DateFrom string `query:"date_from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	DateTo string `query:"date_to,required" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetViewsByBrowsersParams]'s query parameters as
// `url.Values`.
func (r StatisticGetViewsByBrowsersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetViewsByCountryParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	DateFrom string `query:"date_from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	DateTo string `query:"date_to,required" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetViewsByCountryParams]'s query parameters as
// `url.Values`.
func (r StatisticGetViewsByCountryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetViewsByHostnameParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	DateFrom string `query:"date_from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	DateTo string `query:"date_to,required" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetViewsByHostnameParams]'s query parameters as
// `url.Values`.
func (r StatisticGetViewsByHostnameParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetViewsByOperatingSystemParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	DateFrom string `query:"date_from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	DateTo string `query:"date_to,required" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetViewsByOperatingSystemParams]'s query
// parameters as `url.Values`.
func (r StatisticGetViewsByOperatingSystemParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetViewsByRefererParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	DateFrom string `query:"date_from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	DateTo string `query:"date_to,required" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetViewsByRefererParams]'s query parameters as
// `url.Values`.
func (r StatisticGetViewsByRefererParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetViewsByRegionParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	DateFrom string `query:"date_from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	DateTo string `query:"date_to,required" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetViewsByRegionParams]'s query parameters as
// `url.Values`.
func (r StatisticGetViewsByRegionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetViewsHeatmapParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	DateFrom string `query:"date_from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	DateTo string `query:"date_to,required" json:"-"`
	// video streaming ID
	StreamID string `query:"stream_id,required" json:"-"`
	// entity's type
	//
	// Any of "live", "vod", "playlist".
	Type StatisticGetViewsHeatmapParamsType `query:"type,omitzero,required" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetViewsHeatmapParams]'s query parameters as
// `url.Values`.
func (r StatisticGetViewsHeatmapParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// entity's type
type StatisticGetViewsHeatmapParamsType string

const (
	StatisticGetViewsHeatmapParamsTypeLive     StatisticGetViewsHeatmapParamsType = "live"
	StatisticGetViewsHeatmapParamsTypeVod      StatisticGetViewsHeatmapParamsType = "vod"
	StatisticGetViewsHeatmapParamsTypePlaylist StatisticGetViewsHeatmapParamsType = "playlist"
)

type StatisticGetVodStorageVolumeParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	From string `query:"from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	To string `query:"to,required" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetVodStorageVolumeParams]'s query parameters as
// `url.Values`.
func (r StatisticGetVodStorageVolumeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetVodTranscodingDurationParams struct {
	// Start of time frame. Datetime in ISO 8601 format.
	From string `query:"from,required" json:"-"`
	// End of time frame. Datetime in ISO 8601 format.
	To string `query:"to,required" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetVodTranscodingDurationParams]'s query
// parameters as `url.Values`.
func (r StatisticGetVodTranscodingDurationParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetVodUniqueViewersCdnParams struct {
	// Start of time frame. Format is date time in ISO 8601
	From string `query:"from,required" json:"-"`
	// End of time frame. Format is date time in ISO 8601
	To string `query:"to,required" json:"-"`
	// Filter by user "id"
	ClientUserID param.Opt[int64] `query:"client_user_id,omitzero" json:"-"`
	// Filter by video "slug"
	Slug param.Opt[string] `query:"slug,omitzero" json:"-"`
	// Specifies the time interval for grouping data
	//
	// Any of "1m", "5m", "15m", "1h", "1d".
	Granularity StatisticGetVodUniqueViewersCdnParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetVodUniqueViewersCdnParams]'s query parameters
// as `url.Values`.
func (r StatisticGetVodUniqueViewersCdnParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the time interval for grouping data
type StatisticGetVodUniqueViewersCdnParamsGranularity string

const (
	StatisticGetVodUniqueViewersCdnParamsGranularity1m  StatisticGetVodUniqueViewersCdnParamsGranularity = "1m"
	StatisticGetVodUniqueViewersCdnParamsGranularity5m  StatisticGetVodUniqueViewersCdnParamsGranularity = "5m"
	StatisticGetVodUniqueViewersCdnParamsGranularity15m StatisticGetVodUniqueViewersCdnParamsGranularity = "15m"
	StatisticGetVodUniqueViewersCdnParamsGranularity1h  StatisticGetVodUniqueViewersCdnParamsGranularity = "1h"
	StatisticGetVodUniqueViewersCdnParamsGranularity1d  StatisticGetVodUniqueViewersCdnParamsGranularity = "1d"
)

type StatisticGetVodWatchTimeCdnParams struct {
	// Start of the time period for counting minutes of watching. Format is date time
	// in ISO 8601.
	From string `query:"from,required" json:"-"`
	// Filter by field "`client_user_id`"
	ClientUserID param.Opt[int64] `query:"client_user_id,omitzero" json:"-"`
	// Filter by video's slug
	Slug param.Opt[string] `query:"slug,omitzero" json:"-"`
	// End of time frame. Datetime in ISO 8601 format. If omitted, then the current
	// time is taken.
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	// Data is grouped by the specified time interval
	//
	// Any of "1m", "5m", "15m", "1h", "1d", "1mo".
	Granularity StatisticGetVodWatchTimeCdnParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetVodWatchTimeCdnParams]'s query parameters as
// `url.Values`.
func (r StatisticGetVodWatchTimeCdnParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Data is grouped by the specified time interval
type StatisticGetVodWatchTimeCdnParamsGranularity string

const (
	StatisticGetVodWatchTimeCdnParamsGranularity1m  StatisticGetVodWatchTimeCdnParamsGranularity = "1m"
	StatisticGetVodWatchTimeCdnParamsGranularity5m  StatisticGetVodWatchTimeCdnParamsGranularity = "5m"
	StatisticGetVodWatchTimeCdnParamsGranularity15m StatisticGetVodWatchTimeCdnParamsGranularity = "15m"
	StatisticGetVodWatchTimeCdnParamsGranularity1h  StatisticGetVodWatchTimeCdnParamsGranularity = "1h"
	StatisticGetVodWatchTimeCdnParamsGranularity1d  StatisticGetVodWatchTimeCdnParamsGranularity = "1d"
	StatisticGetVodWatchTimeCdnParamsGranularity1mo StatisticGetVodWatchTimeCdnParamsGranularity = "1mo"
)

type StatisticGetVodWatchTimeTotalCdnParams struct {
	// Filter by field "`client_user_id`"
	ClientUserID param.Opt[int64] `query:"client_user_id,omitzero" json:"-"`
	// Start of the time period for counting minutes of watching. Format is date time
	// in ISO 8601. If omitted, the earliest start time for viewing is taken
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// Filter by video's slug
	Slug param.Opt[string] `query:"slug,omitzero" json:"-"`
	// End of time frame. Datetime in ISO 8601 format. If omitted, then the current
	// time is taken
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetVodWatchTimeTotalCdnParams]'s query parameters
// as `url.Values`.
func (r StatisticGetVodWatchTimeTotalCdnParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
