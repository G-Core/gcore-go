// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
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
func (r *StatisticService) GetUsageSeries(ctx context.Context, query StatisticGetUsageSeriesParams, opts ...option.RequestOption) (res *WaapStatisticsSeries, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/statistics/series"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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
