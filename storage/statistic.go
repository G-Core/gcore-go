// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
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

// Consumption statistics is updated in near real-time as a standard practice.
// However, the frequency of updates can vary, but they are typically available
// within a 60 minutes period. Exceptions, such as maintenance periods, may delay
// data beyond 60 minutes until servers resume and backfill missing statistics.
//
// Shows storage total usage data in filtered by storages, locations and interval.
func (r *StatisticService) GetUsageAggregated(ctx context.Context, body StatisticGetUsageAggregatedParams, opts ...option.RequestOption) (res *StorageUsageTotal, err error) {
	opts = append(r.Options[:], opts...)
	path := "storage/stats/v1/storage/usage/total"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Consumption statistics is updated in near real-time as a standard practice.
// However, the frequency of updates can vary, but they are typically available
// within a 60 minutes period. Exceptions, such as maintenance periods, may delay
// data beyond 60 minutes until servers resume and backfill missing statistics.
//
// Shows storage usage data in series format filtered by clients, storages and
// interval.
func (r *StatisticService) GetUsageSeries(ctx context.Context, body StatisticGetUsageSeriesParams, opts ...option.RequestOption) (res *StatisticGetUsageSeriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "storage/stats/v1/storage/usage/series"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type StorageUsageSeries struct {
	// a Clients grouped data
	Clients map[string]StorageUsageSeriesClient `json:"clients"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clients     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageUsageSeries) RawJSON() string { return r.JSON.raw }
func (r *StorageUsageSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageUsageSeriesClient struct {
	// an ID of client
	ID int64 `json:"id"`
	// a FileQuantitySumMax is max sum of files quantity for grouped period
	FileQuantitySumMax int64 `json:"file_quantity_sum_max"`
	// a Locations grouped data
	Locations map[string]StorageUsageSeriesClientLocation `json:"locations"`
	// a RequestsInSum is sum of incoming requests for grouped period
	RequestsInSum int64 `json:"requests_in_sum"`
	// a RequestsOutEdgesSum is sum of out edges requests for grouped period
	RequestsOutEdgesSum int64 `json:"requests_out_edges_sum"`
	// a RequestsOutWoEdgesSum is sum of out no edges requests for grouped period
	RequestsOutWoEdgesSum int64 `json:"requests_out_wo_edges_sum"`
	// a RequestsSum is sum of all requests for grouped period
	RequestsSum int64 `json:"requests_sum"`
	// a SizeSumBytesHour is sum of bytes hour for grouped period
	SizeSumBytesHour int64 `json:"size_sum_bytes_hour"`
	// a SizeSumMax is max sum of all files sizes for grouped period
	SizeSumMax int64 `json:"size_sum_max"`
	// a SizeSumMean is mean sum of all files sizes for grouped period
	SizeSumMean int64 `json:"size_sum_mean"`
	// a TrafficInSum is sum of incoming traffic for grouped period
	TrafficInSum int64 `json:"traffic_in_sum"`
	// a TrafficOutEdgesSum is sum of out edges traffic for grouped period
	TrafficOutEdgesSum int64 `json:"traffic_out_edges_sum"`
	// a TrafficOutWoEdgesSum is sum of out no edges traffic for grouped period
	TrafficOutWoEdgesSum int64 `json:"traffic_out_wo_edges_sum"`
	// a TrafficSum is sum of all traffic for grouped period
	TrafficSum int64 `json:"traffic_sum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		FileQuantitySumMax    respjson.Field
		Locations             respjson.Field
		RequestsInSum         respjson.Field
		RequestsOutEdgesSum   respjson.Field
		RequestsOutWoEdgesSum respjson.Field
		RequestsSum           respjson.Field
		SizeSumBytesHour      respjson.Field
		SizeSumMax            respjson.Field
		SizeSumMean           respjson.Field
		TrafficInSum          respjson.Field
		TrafficOutEdgesSum    respjson.Field
		TrafficOutWoEdgesSum  respjson.Field
		TrafficSum            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageUsageSeriesClient) RawJSON() string { return r.JSON.raw }
func (r *StorageUsageSeriesClient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageUsageSeriesClientLocation struct {
	// a FileQuantitySumMax is max sum of files quantity for grouped period
	FileQuantitySumMax int64 `json:"file_quantity_sum_max"`
	// a Name of location
	Name string `json:"name"`
	// a RequestsInSum is sum of incoming requests for grouped period
	RequestsInSum int64 `json:"requests_in_sum"`
	// a RequestsOutEdgesSum is sum of out edges requests for grouped period
	RequestsOutEdgesSum int64 `json:"requests_out_edges_sum"`
	// a RequestsOutWoEdgesSum is sum of out no edges requests for grouped period
	RequestsOutWoEdgesSum int64 `json:"requests_out_wo_edges_sum"`
	// a RequestsSum is sum of all requests for grouped period
	RequestsSum int64 `json:"requests_sum"`
	// a SizeSumBytesHour is sum of bytes hour for grouped period
	SizeSumBytesHour int64 `json:"size_sum_bytes_hour"`
	// a SizeSumMax is max sum of all files sizes for grouped period
	SizeSumMax int64 `json:"size_sum_max"`
	// a SizeSumMean is mean sum of all files sizes for grouped period
	SizeSumMean int64 `json:"size_sum_mean"`
	// a Storages grouped data
	Storages map[string]StorageUsageSeriesClientLocationStorage `json:"storages"`
	// a TrafficInSum is sum of incoming traffic for grouped period
	TrafficInSum int64 `json:"traffic_in_sum"`
	// a TrafficOutEdgesSum is sum of out edges traffic for grouped period
	TrafficOutEdgesSum int64 `json:"traffic_out_edges_sum"`
	// a TrafficOutWoEdgesSum is sum of out no edges traffic for grouped period
	TrafficOutWoEdgesSum int64 `json:"traffic_out_wo_edges_sum"`
	// a TrafficSum is sum of all traffic for grouped period
	TrafficSum int64 `json:"traffic_sum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileQuantitySumMax    respjson.Field
		Name                  respjson.Field
		RequestsInSum         respjson.Field
		RequestsOutEdgesSum   respjson.Field
		RequestsOutWoEdgesSum respjson.Field
		RequestsSum           respjson.Field
		SizeSumBytesHour      respjson.Field
		SizeSumMax            respjson.Field
		SizeSumMean           respjson.Field
		Storages              respjson.Field
		TrafficInSum          respjson.Field
		TrafficOutEdgesSum    respjson.Field
		TrafficOutWoEdgesSum  respjson.Field
		TrafficSum            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageUsageSeriesClientLocation) RawJSON() string { return r.JSON.raw }
func (r *StorageUsageSeriesClientLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageUsageSeriesClientLocationStorage struct {
	// a BucketsSeries is max bucket files count for grouped period
	// {name:[[timestamp, count]]}
	BucketsSeries map[string][][]any `json:"buckets_series"`
	// a FileQuantitySumMax is max sum of files quantity for grouped period
	FileQuantitySumMax int64 `json:"file_quantity_sum_max"`
	// a Name of storage
	Name string `json:"name"`
	// a RequestsInSeries is sum of incoming requests for grouped period
	// [[timestamp, count]]
	RequestsInSeries [][]any `json:"requests_in_series"`
	// a RequestsInSum is sum of incoming requests for grouped period
	RequestsInSum int64 `json:"requests_in_sum"`
	// a RequestsOutWoEdgesSeries is sum of out requests (only edges) for grouped
	// period [[timestamp, count]]
	RequestsOutEdgesSeries [][]any `json:"requests_out_edges_series"`
	// a RequestsOutEdgesSum is sum of out edges requests for grouped period
	RequestsOutEdgesSum int64 `json:"requests_out_edges_sum"`
	// a RequestsOutWoEdgesSeries is sum of out requests (without edges) for grouped
	// period [[timestamp, count]]
	RequestsOutWoEdgesSeries [][]any `json:"requests_out_wo_edges_series"`
	// a RequestsOutWoEdgesSum is sum of out no edges requests for grouped period
	RequestsOutWoEdgesSum int64 `json:"requests_out_wo_edges_sum"`
	// a RequestsSeries is sum of out requests for grouped period [[timestamp, count]]
	RequestsSeries [][]any `json:"requests_series"`
	// a RequestsSum is sum of all requests for grouped period
	RequestsSum int64 `json:"requests_sum"`
	// a SizeBytesHourSeries is value that displays how many bytes were stored per hour
	// [[timestamp, count]]
	SizeBytesHourSeries [][]any `json:"size_bytes_hour_series"`
	// a SizeMaxSeries is max of files size for grouped period [[timestamp, count]]
	SizeMaxSeries [][]any `json:"size_max_series"`
	// a SizeMeanSeries is mean of files size for grouped period [[timestamp, count]]
	SizeMeanSeries [][]any `json:"size_mean_series"`
	// a SizeSumBytesHour is sum of bytes hour for grouped period
	SizeSumBytesHour int64 `json:"size_sum_bytes_hour"`
	// a SizeSumMax is max sum of all files sizes for grouped period
	SizeSumMax int64 `json:"size_sum_max"`
	// a SizeSumMean is mean sum of all files sizes for grouped period
	SizeSumMean int64 `json:"size_sum_mean"`
	// a TrafficInSeries is sum of incoming traffic bytes for grouped period
	// [[timestamp, count]]
	TrafficInSeries [][]any `json:"traffic_in_series"`
	// a TrafficInSum is sum of incoming traffic for grouped period
	TrafficInSum int64 `json:"traffic_in_sum"`
	// a TrafficOutWoEdgesSeries is sum of out traffic bytes (only edges) for grouped
	// period [[timestamp, count]]
	TrafficOutEdgesSeries [][]any `json:"traffic_out_edges_series"`
	// a TrafficOutEdgesSum is sum of out edges traffic for grouped period
	TrafficOutEdgesSum int64 `json:"traffic_out_edges_sum"`
	// a TrafficOutWoEdgesSeries is sum of out traffic bytes (without edges) for
	// grouped period [[timestamp, count]]
	TrafficOutWoEdgesSeries [][]any `json:"traffic_out_wo_edges_series"`
	// a TrafficOutWoEdgesSum is sum of out no edges traffic for grouped period
	TrafficOutWoEdgesSum int64 `json:"traffic_out_wo_edges_sum"`
	// a TrafficSeries is sum of traffic bytes for grouped period [[timestamp, count]]
	TrafficSeries [][]any `json:"traffic_series"`
	// a TrafficSum is sum of all traffic for grouped period
	TrafficSum int64 `json:"traffic_sum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketsSeries            respjson.Field
		FileQuantitySumMax       respjson.Field
		Name                     respjson.Field
		RequestsInSeries         respjson.Field
		RequestsInSum            respjson.Field
		RequestsOutEdgesSeries   respjson.Field
		RequestsOutEdgesSum      respjson.Field
		RequestsOutWoEdgesSeries respjson.Field
		RequestsOutWoEdgesSum    respjson.Field
		RequestsSeries           respjson.Field
		RequestsSum              respjson.Field
		SizeBytesHourSeries      respjson.Field
		SizeMaxSeries            respjson.Field
		SizeMeanSeries           respjson.Field
		SizeSumBytesHour         respjson.Field
		SizeSumMax               respjson.Field
		SizeSumMean              respjson.Field
		TrafficInSeries          respjson.Field
		TrafficInSum             respjson.Field
		TrafficOutEdgesSeries    respjson.Field
		TrafficOutEdgesSum       respjson.Field
		TrafficOutWoEdgesSeries  respjson.Field
		TrafficOutWoEdgesSum     respjson.Field
		TrafficSeries            respjson.Field
		TrafficSum               respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageUsageSeriesClientLocationStorage) RawJSON() string { return r.JSON.raw }
func (r *StorageUsageSeriesClientLocationStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageUsageTotal struct {
	// StorageUsageTotalRes for response
	Data []StorageUsageTotalData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageUsageTotal) RawJSON() string { return r.JSON.raw }
func (r *StorageUsageTotal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// StorageStatsTotalElement for response
type StorageUsageTotalData struct {
	Metrics StorageUsageTotalDataMetrics `json:"metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageUsageTotalData) RawJSON() string { return r.JSON.raw }
func (r *StorageUsageTotalData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageUsageTotalDataMetrics struct {
	// a FileQuantitySumMax is max sum of files quantity for grouped period
	FileQuantitySumMax int64 `json:"file_quantity_sum_max"`
	// a RequestsInSum is sum of incoming requests for grouped period
	RequestsInSum int64 `json:"requests_in_sum"`
	// a RequestsOutEdgesSum is sum of out edges requests for grouped period
	RequestsOutEdgesSum int64 `json:"requests_out_edges_sum"`
	// a RequestsOutWoEdgesSum is sum of out no edges requests for grouped period
	RequestsOutWoEdgesSum int64 `json:"requests_out_wo_edges_sum"`
	// a RequestsSum is sum of all requests for grouped period
	RequestsSum int64 `json:"requests_sum"`
	// a SizeSumBytesHour is sum of bytes hour for grouped period
	SizeSumBytesHour int64 `json:"size_sum_bytes_hour"`
	// a SizeSumMax is max sum of all files sizes for grouped period
	SizeSumMax int64 `json:"size_sum_max"`
	// a SizeSumMean is mean sum of all files sizes for grouped period
	SizeSumMean int64 `json:"size_sum_mean"`
	// a TrafficInSum is sum of incoming traffic for grouped period
	TrafficInSum int64 `json:"traffic_in_sum"`
	// a TrafficOutEdgesSum is sum of out edges traffic for grouped period
	TrafficOutEdgesSum int64 `json:"traffic_out_edges_sum"`
	// a TrafficOutWoEdgesSum is sum of out no edges traffic for grouped period
	TrafficOutWoEdgesSum int64 `json:"traffic_out_wo_edges_sum"`
	// a TrafficSum is sum of all traffic for grouped period
	TrafficSum int64 `json:"traffic_sum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileQuantitySumMax    respjson.Field
		RequestsInSum         respjson.Field
		RequestsOutEdgesSum   respjson.Field
		RequestsOutWoEdgesSum respjson.Field
		RequestsSum           respjson.Field
		SizeSumBytesHour      respjson.Field
		SizeSumMax            respjson.Field
		SizeSumMean           respjson.Field
		TrafficInSum          respjson.Field
		TrafficOutEdgesSum    respjson.Field
		TrafficOutWoEdgesSum  respjson.Field
		TrafficSum            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageUsageTotalDataMetrics) RawJSON() string { return r.JSON.raw }
func (r *StorageUsageTotalDataMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatisticGetUsageSeriesResponse struct {
	Data StorageUsageSeries `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatisticGetUsageSeriesResponse) RawJSON() string { return r.JSON.raw }
func (r *StatisticGetUsageSeriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatisticGetUsageAggregatedParams struct {
	// a From date filter
	From param.Opt[string] `json:"from,omitzero"`
	// a To date filter
	To param.Opt[string] `json:"to,omitzero"`
	// a Locations list of filter
	Locations []string `json:"locations,omitzero"`
	// a Storages list of filter
	Storages []string `json:"storages,omitzero"`
	paramObj
}

func (r StatisticGetUsageAggregatedParams) MarshalJSON() (data []byte, err error) {
	type shadow StatisticGetUsageAggregatedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StatisticGetUsageAggregatedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatisticGetUsageSeriesParams struct {
	// a From date filter
	From param.Opt[string] `json:"from,omitzero"`
	// a Granularity is period of time for grouping data Valid values are: 1h, 12h, 24h
	Granularity param.Opt[string] `json:"granularity,omitzero"`
	// a Source is deprecated parameter
	Source param.Opt[int64] `json:"source,omitzero"`
	// a To date filter
	To param.Opt[string] `json:"to,omitzero"`
	// a TsString is configurator of response time format switch response from unix
	// time format to RFC3339 (2006-01-02T15:04:05Z07:00)
	TsString param.Opt[bool] `json:"ts_string,omitzero"`
	// a Locations list of filter
	Locations []string `json:"locations,omitzero"`
	// a Storages list of filter
	Storages []string `json:"storages,omitzero"`
	paramObj
}

func (r StatisticGetUsageSeriesParams) MarshalJSON() (data []byte, err error) {
	type shadow StatisticGetUsageSeriesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StatisticGetUsageSeriesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
