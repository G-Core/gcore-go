// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// WaapV1StatisticService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1StatisticService] method instead.
type WaapV1StatisticService struct {
	Options []option.RequestOption
}

// NewWaapV1StatisticService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWaapV1StatisticService(opts ...option.RequestOption) (r *WaapV1StatisticService) {
	r = &WaapV1StatisticService{}
	r.Options = opts
	return
}

// Retrieve statistics data as a time series. The `from` and `to` parameters are
// rounded down and up according to the `granularity`. This means that if the
// `granularity` is set to `1h`, the `from` and `to` parameters will be rounded
// down and up to the nearest hour, respectively. If the `granularity` is set to
// `1d`, the `from` and `to` parameters will be rounded down and up to the nearest
// day, respectively. The response will include explicit 0 values for any missing
// data points.
func (r *WaapV1StatisticService) GetSeries(ctx context.Context, query WaapV1StatisticGetSeriesParams, opts ...option.RequestOption) (res *WaapV1StatisticGetSeriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/statistics/series"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response model for the statistics item
type StatisticItem struct {
	// The date and time for the statistic in ISO 8601 format
	DateTime time.Time `json:"date_time,required" format:"date-time"`
	// The value for the statistic. If there is no data for the given time, the value
	// will be 0.
	Value int64             `json:"value,required"`
	JSON  statisticItemJSON `json:"-"`
}

// statisticItemJSON contains the JSON metadata for the struct [StatisticItem]
type statisticItemJSON struct {
	DateTime    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatisticItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statisticItemJSON) RawJSON() string {
	return r.raw
}

// Response model for the statistics series
type WaapV1StatisticGetSeriesResponse struct {
	// Will be returned if `total_bytes` is requested in the metrics parameter
	TotalBytes []StatisticItem `json:"total_bytes,nullable"`
	// Will be included if `total_requests` is requested in the metrics parameter
	TotalRequests []StatisticItem                      `json:"total_requests,nullable"`
	JSON          waapV1StatisticGetSeriesResponseJSON `json:"-"`
}

// waapV1StatisticGetSeriesResponseJSON contains the JSON metadata for the struct
// [WaapV1StatisticGetSeriesResponse]
type waapV1StatisticGetSeriesResponseJSON struct {
	TotalBytes    apijson.Field
	TotalRequests apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WaapV1StatisticGetSeriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1StatisticGetSeriesResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1StatisticGetSeriesParams struct {
	// Beginning of the requested time period (ISO 8601 format, UTC)
	From param.Field[time.Time] `query:"from,required" format:"date-time"`
	// Duration of the time blocks into which the data will be divided.
	Granularity param.Field[WaapV1StatisticGetSeriesParamsGranularity] `query:"granularity,required"`
	// List of metric types to retrieve statistics for.
	Metrics param.Field[[]WaapV1StatisticGetSeriesParamsMetric] `query:"metrics,required"`
	// End of the requested time period (ISO 8601 format, UTC)
	To param.Field[time.Time] `query:"to,required" format:"date-time"`
}

// URLQuery serializes [WaapV1StatisticGetSeriesParams]'s query parameters as
// `url.Values`.
func (r WaapV1StatisticGetSeriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Duration of the time blocks into which the data will be divided.
type WaapV1StatisticGetSeriesParamsGranularity string

const (
	WaapV1StatisticGetSeriesParamsGranularity1h WaapV1StatisticGetSeriesParamsGranularity = "1h"
	WaapV1StatisticGetSeriesParamsGranularity1d WaapV1StatisticGetSeriesParamsGranularity = "1d"
)

func (r WaapV1StatisticGetSeriesParamsGranularity) IsKnown() bool {
	switch r {
	case WaapV1StatisticGetSeriesParamsGranularity1h, WaapV1StatisticGetSeriesParamsGranularity1d:
		return true
	}
	return false
}

type WaapV1StatisticGetSeriesParamsMetric string

const (
	WaapV1StatisticGetSeriesParamsMetricTotalBytes    WaapV1StatisticGetSeriesParamsMetric = "total_bytes"
	WaapV1StatisticGetSeriesParamsMetricTotalRequests WaapV1StatisticGetSeriesParamsMetric = "total_requests"
)

func (r WaapV1StatisticGetSeriesParamsMetric) IsKnown() bool {
	switch r {
	case WaapV1StatisticGetSeriesParamsMetricTotalBytes, WaapV1StatisticGetSeriesParamsMetricTotalRequests:
		return true
	}
	return false
}
