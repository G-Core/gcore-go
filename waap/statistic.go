// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
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

// Retrieve statistics data as a time series. The `from` and `to` parameters are
// rounded down and up according to the `granularity`. This means that if the
// `granularity` is set to `1h`, the `from` and `to` parameters will be rounded
// down and up to the nearest hour, respectively. If the `granularity` is set to
// `1d`, the `from` and `to` parameters will be rounded down and up to the nearest
// day, respectively. The response will include explicit 0 values for any missing
// data points.
func (r *StatisticService) GetUsageSeries(ctx context.Context, query StatisticGetUsageSeriesParams, opts ...option.RequestOption) (res *StatisticsSeries, err error) {
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
	Value int64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateTime    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatisticItem) RawJSON() string { return r.JSON.raw }
func (r *StatisticItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response model for the statistics series
type StatisticsSeries struct {
	// Will be returned if `total_bytes` is requested in the metrics parameter
	TotalBytes []StatisticItem `json:"total_bytes,nullable"`
	// Will be included if `total_requests` is requested in the metrics parameter
	TotalRequests []StatisticItem `json:"total_requests,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TotalBytes    respjson.Field
		TotalRequests respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatisticsSeries) RawJSON() string { return r.JSON.raw }
func (r *StatisticsSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatisticGetUsageSeriesParams struct {
	// Beginning of the requested time period (ISO 8601 format, UTC)
	From time.Time `query:"from,required" format:"date-time" json:"-"`
	// Duration of the time blocks into which the data will be divided.
	//
	// Any of "1h", "1d".
	Granularity StatisticGetUsageSeriesParamsGranularity `query:"granularity,omitzero,required" json:"-"`
	// List of metric types to retrieve statistics for.
	//
	// Any of "total_bytes", "total_requests".
	Metrics []string `query:"metrics,omitzero,required" json:"-"`
	// End of the requested time period (ISO 8601 format, UTC)
	To time.Time `query:"to,required" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetUsageSeriesParams]'s query parameters as
// `url.Values`.
func (r StatisticGetUsageSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Duration of the time blocks into which the data will be divided.
type StatisticGetUsageSeriesParamsGranularity string

const (
	StatisticGetUsageSeriesParamsGranularity1h StatisticGetUsageSeriesParamsGranularity = "1h"
	StatisticGetUsageSeriesParamsGranularity1d StatisticGetUsageSeriesParamsGranularity = "1d"
)
