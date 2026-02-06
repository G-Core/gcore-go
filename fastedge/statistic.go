// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fastedge

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

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

// Call statistics
func (r *StatisticService) GetCallSeries(ctx context.Context, query StatisticGetCallSeriesParams, opts ...option.RequestOption) (res *StatisticGetCallSeriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fastedge/v1/stats/calls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Execution duration statistics
func (r *StatisticService) GetDurationSeries(ctx context.Context, query StatisticGetDurationSeriesParams, opts ...option.RequestOption) (res *StatisticGetDurationSeriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fastedge/v1/stats/app_duration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Edge app call statistics
type CallStatus struct {
	// Count by status
	CountByStatus []CallStatusCountByStatus `json:"count_by_status,required"`
	// Beginning ot reporting slot
	Time time.Time `json:"time,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountByStatus respjson.Field
		Time          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallStatus) RawJSON() string { return r.JSON.raw }
func (r *CallStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallStatusCountByStatus struct {
	// Number of app calls
	Count int64 `json:"count,required"`
	// HTTP status
	Status int64 `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallStatusCountByStatus) RawJSON() string { return r.JSON.raw }
func (r *CallStatusCountByStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Edge app execution duration statistics
type DurationStats struct {
	// Average duration in usec
	Avg int64 `json:"avg,required"`
	// Max duration in usec
	Max int64 `json:"max,required"`
	// Median (50% percentile) duration in usec
	Median int64 `json:"median,required"`
	// Min duration in usec
	Min int64 `json:"min,required"`
	// 75% percentile duration in usec
	Perc75 int64 `json:"perc75,required"`
	// 90% percentile duration in usec
	Perc90 int64 `json:"perc90,required"`
	// Beginning of reporting slot
	Time time.Time `json:"time,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Avg         respjson.Field
		Max         respjson.Field
		Median      respjson.Field
		Min         respjson.Field
		Perc75      respjson.Field
		Perc90      respjson.Field
		Time        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DurationStats) RawJSON() string { return r.JSON.raw }
func (r *DurationStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatisticGetCallSeriesResponse struct {
	Stats []CallStatus `json:"stats,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Stats       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatisticGetCallSeriesResponse) RawJSON() string { return r.JSON.raw }
func (r *StatisticGetCallSeriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatisticGetDurationSeriesResponse struct {
	Stats []DurationStats `json:"stats,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Stats       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatisticGetDurationSeriesResponse) RawJSON() string { return r.JSON.raw }
func (r *StatisticGetDurationSeriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatisticGetCallSeriesParams struct {
	// Reporting period start time, RFC3339 format
	From time.Time `query:"from,required" format:"date-time" json:"-"`
	// Reporting granularity, in seconds
	Step int64 `query:"step,required" json:"-"`
	// Reporting period end time (not included into reporting period), RFC3339 format
	To time.Time `query:"to,required" format:"date-time" json:"-"`
	// App ID
	ID param.Opt[int64] `query:"id,omitzero" json:"-"`
	// Network name
	Network param.Opt[string] `query:"network,omitzero" format:"string" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetCallSeriesParams]'s query parameters as
// `url.Values`.
func (r StatisticGetCallSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StatisticGetDurationSeriesParams struct {
	// Reporting period start time, RFC3339 format
	From time.Time `query:"from,required" format:"date-time" json:"-"`
	// Reporting granularity, in seconds
	Step int64 `query:"step,required" json:"-"`
	// Reporting period end time (not included into reporting period), RFC3339 format
	To time.Time `query:"to,required" format:"date-time" json:"-"`
	// App ID
	ID param.Opt[int64] `query:"id,omitzero" json:"-"`
	// Network name
	Network param.Opt[string] `query:"network,omitzero" format:"string" json:"-"`
	paramObj
}

// URLQuery serializes [StatisticGetDurationSeriesParams]'s query parameters as
// `url.Values`.
func (r StatisticGetDurationSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
