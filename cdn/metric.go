// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// MetricService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMetricService] method instead.
type MetricService struct {
	Options []option.RequestOption
}

// NewMetricService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMetricService(opts ...option.RequestOption) (r MetricService) {
	r = MetricService{}
	r.Options = opts
	return
}

// Get CDN metrics
func (r *MetricService) List(ctx context.Context, body MetricListParams, opts ...option.RequestOption) (res *CdnMetrics, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/advanced/v1/metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CdnMetrics struct {
	// If no grouping was requested then "data" holds an array of metric values. If at
	// least one field is specified in "group_by" then "data" is an object whose
	// properties are groups, which may include other groups; the last group will hold
	// array of metrics values.
	Data CdnMetricsDataUnion `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnMetrics) RawJSON() string { return r.JSON.raw }
func (r *CdnMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CdnMetricsDataUnion contains all possible properties and values from
// [CdnMetricsValues], [CdnMetricsGroups].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCdnMetricsValues]
type CdnMetricsDataUnion struct {
	// This field will be present if the value is a [CdnMetricsValues] instead of an
	// object.
	OfCdnMetricsValues CdnMetricsValues `json:",inline"`
	// This field is from variant [CdnMetricsGroups].
	Group CdnMetricsValues `json:"group"`
	JSON  struct {
		OfCdnMetricsValues respjson.Field
		Group              respjson.Field
		raw                string
	} `json:"-"`
}

func (u CdnMetricsDataUnion) AsCdnMetricsValues() (v CdnMetricsValues) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CdnMetricsDataUnion) AsCdnMetricsGroups() (v CdnMetricsGroups) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CdnMetricsDataUnion) RawJSON() string { return u.JSON.raw }

func (r *CdnMetricsDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdnMetricsGroups struct {
	// List of requested metrics sorted by timestamp in ascending order.
	Group CdnMetricsValues `json:"group"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Group       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnMetricsGroups) RawJSON() string { return r.JSON.raw }
func (r *CdnMetricsGroups) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdnMetricsValues []CdnMetricsValue

type CdnMetricsValue struct {
	// Metrics value.
	Metric float64 `json:"metric"`
	// Start timestamp of interval.
	Timestamp int64 `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metric      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnMetricsValue) RawJSON() string { return r.JSON.raw }
func (r *CdnMetricsValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MetricListParams struct {
	// Beginning period to fetch metrics (ISO 8601/RFC 3339 format, UTC.)
	//
	// Examples:
	//
	// - 2021-06-14T00:00:00Z
	// - 2021-06-14T00:00:00.000Z
	//
	// The total number of points, which is determined as the difference between "from"
	// and "to" divided by "granularity", cannot exceed 1440. Exception: "speed"
	// metrics are limited to 72 points.
	From string `json:"from,required"`
	// Possible values:
	//
	//   - **`edge_bandwidth`** - Bandwidth from client to CDN (bit/s.)
	//   - **`edge_requests`** - Number of requests per interval (requests/s.)
	//   - **`edge_requests_total`** - Total number of requests per interval.
	//   - **`edge_status_1xx`** - Number of 1xx status codes from edge.
	//   - **`edge_status_200`** - Number of 200 status codes from edge.
	//   - **`edge_status_204`** - Number of 204 status codes from edge.
	//   - **`edge_status_206`** - Number of 206 status codes from edge.
	//   - **`edge_status_2xx`** - Number of 2xx status codes from edge.
	//   - **`edge_status_301`** - Number of 301 status codes from edge.
	//   - **`edge_status_302`** - Number of 302 status codes from edge.
	//   - **`edge_status_304`** - Number of 304 status codes from edge.
	//   - **`edge_status_3xx`** - Number of 3xx status codes from edge.
	//   - **`edge_status_400`** - Number of 400 status codes from edge.
	//   - **`edge_status_401`** - Number of 401 status codes from edge.
	//   - **`edge_status_403`** - Number of 403 status codes from edge.
	//   - **`edge_status_404`** - Number of 404 status codes from edge.
	//   - **`edge_status_416`** - Number of 416 status codes from edge.
	//   - **`edge_status_429`** - Number of 429 status codes from edge.
	//   - **`edge_status_4xx`** - Number of 4xx status codes from edge.
	//   - **`edge_status_500`** - Number of 500 status codes from edge.
	//   - **`edge_status_501`** - Number of 501 status codes from edge.
	//   - **`edge_status_502`** - Number of 502 status codes from edge.
	//   - **`edge_status_503`** - Number of 503 status codes from edge.
	//   - **`edge_status_504`** - Number of 504 status codes from edge.
	//   - **`edge_status_505`** - Number of 505 status codes from edge.
	//   - **`edge_status_5xx`** - Number of 5xx status codes from edge.
	//   - **`edge_hit_ratio`** - Percent of cache hits (0.0 - 1.0).
	//   - **`edge_hit_bytes`** - Number of bytes sent back when cache hits.
	//   - **`origin_bandwidth`** - Bandwidth from CDN to Origin (bit/s.)
	//   - **`origin_requests`** - Number of requests per interval (requests/s.)
	//   - **`origin_status_1xx`** - Number of 1xx status from origin.
	//   - **`origin_status_200`** - Number of 200 status from origin.
	//   - **`origin_status_204`** - Number of 204 status from origin.
	//   - **`origin_status_206`** - Number of 206 status from origin.
	//   - **`origin_status_2xx`** - Number of 2xx status from origin.
	//   - **`origin_status_301`** - Number of 301 status from origin.
	//   - **`origin_status_302`** - Number of 302 status from origin.
	//   - **`origin_status_304`** - Number of 304 status from origin.
	//   - **`origin_status_3xx`** - Number of 3xx status from origin.
	//   - **`origin_status_400`** - Number of 400 status from origin.
	//   - **`origin_status_401`** - Number of 401 status from origin.
	//   - **`origin_status_403`** - Number of 403 status from origin.
	//   - **`origin_status_404`** - Number of 404 status from origin.
	//   - **`origin_status_416`** - Number of 416 status from origin.
	//   - **`origin_status_429`** - Number of 426 status from origin.
	//   - **`origin_status_4xx`** - Number of 4xx status from origin.
	//   - **`origin_status_500`** - Number of 500 status from origin.
	//   - **`origin_status_501`** - Number of 501 status from origin.
	//   - **`origin_status_502`** - Number of 502 status from origin.
	//   - **`origin_status_503`** - Number of 503 status from origin.
	//   - **`origin_status_504`** - Number of 504 status from origin.
	//   - **`origin_status_505`** - Number of 505 status from origin.
	//   - **`origin_status_5xx`** - Number of 5xx status from origin.
	//   - **`edge_download_speed`** - Download speed from edge in KB/s (includes only
	//     requests that status was in the range [200, 300].)
	//   - **`origin_download_speed`** - Download speed from origin in KB/s (includes
	//     only requests that status was in the range [200, 300].)
	Metrics []string `json:"metrics,omitzero,required"`
	// Specifies ending period to fetch metrics (ISO 8601/RFC 3339 format, UTC)
	//
	// Examples:
	//
	// - 2021-06-15T00:00:00Z
	// - 2021-06-15T00:00:00.000Z
	//
	// The total number of points, which is determined as the difference between "from"
	// and "to" divided by "granularity", cannot exceed 1440. Exception: "speed"
	// metrics are limited to 72 points.
	To string `json:"to,required"`
	// Duration of the time blocks into which the data is divided. The value must
	// correspond to the ISO 8601 period format.
	//
	// Examples:
	//
	// - P1D
	// - PT5M
	//
	// Notes:
	//
	//   - The total number of points, which is determined as the difference between
	//     "from" and "to" divided by "granularity", cannot exceed 1440. Exception:
	//     "speed" metrics are limited to 72 points.
	//   - For "speed" metrics the value must be a multiple of 5.
	Granularity param.Opt[string] `json:"granularity,omitzero" format:"P(n)Y(n)M(n)DT(n)H(n)M)"`
	// Each item represents one filter statement.
	FilterBy []MetricListParamsFilterBy `json:"filter_by,omitzero"`
	// Output data grouping.
	//
	// Possible values:
	//
	//   - **resource** - Data is grouped by CDN resource.
	//   - **cname** - Data is grouped by common names.
	//   - **region** – Data is grouped by regions (continents.) Available for "speed"
	//     metrics only.
	//   - **isp** - Data is grouped by ISP names. Available for "speed" metrics only.
	GroupBy []string `json:"group_by,omitzero"`
	paramObj
}

func (r MetricListParams) MarshalJSON() (data []byte, err error) {
	type shadow MetricListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MetricListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Field, Op, Values are required.
type MetricListParamsFilterBy struct {
	// Defines the parameters by that data can be filtered.
	//
	// Possible values:
	//
	//   - **resource** - Data is filtered by CDN resource ID.
	//   - **cname** - Data is filtered by common name.
	//   - **region** - Data is filtered by region (continent.) Available for "speed"
	//     metrics only.
	//   - **isp** - Data is filtered by ISP name. Available for "speed" metrics only.
	Field string `json:"field,required"`
	// Comparison operator to be applied.
	//
	// Possible values:
	//
	// - **in** - 'IN' operator.
	// - **`not_in`** - 'NOT IN' operator.
	// - **gt** - '>' operator.
	// - **gte** - '>=' operator.
	// - **lt** - '<' operator.
	// - **lte** - '<=' operator.
	// - **eq** - '==' operator.
	// - **ne** - '!=' operator.
	// - **like** - 'LIKE' operator.
	// - **`not_like`** - 'NOT LIKE' operator.
	Op string `json:"op,required"`
	// Contains one or more values to be compared against.
	Values []MetricListParamsFilterByValueUnion `json:"values,omitzero,required"`
	paramObj
}

func (r MetricListParamsFilterBy) MarshalJSON() (data []byte, err error) {
	type shadow MetricListParamsFilterBy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MetricListParamsFilterBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MetricListParamsFilterByValueUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u MetricListParamsFilterByValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *MetricListParamsFilterByValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MetricListParamsFilterByValueUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}
