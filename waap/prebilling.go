// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// PrebillingService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrebillingService] method instead.
type PrebillingService struct {
	Options []option.RequestOption
}

// NewPrebillingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrebillingService(opts ...option.RequestOption) (r PrebillingService) {
	r = PrebillingService{}
	r.Options = opts
	return
}

// Retrieve a report on the calling client's own billable consumption by given
// metric with per-account breakdown. The report is paginated by accounts: `limit`,
// `offset` and `count` are counted in accounts, and each account is returned with
// all its items on exactly one page
func (r *PrebillingService) GetBreakdown(ctx context.Context, metric PrebillingGetBreakdownParamsMetric, query PrebillingGetBreakdownParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapClientBillingReport], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("waap/v2/prebilling/breakdown/%v", metric)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve a report on the calling client's own billable consumption by given
// metric with per-account breakdown. The report is paginated by accounts: `limit`,
// `offset` and `count` are counted in accounts, and each account is returned with
// all its items on exactly one page
func (r *PrebillingService) GetBreakdownAutoPaging(ctx context.Context, metric PrebillingGetBreakdownParamsMetric, query PrebillingGetBreakdownParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapClientBillingReport] {
	return pagination.NewOffsetPageAutoPager(r.GetBreakdown(ctx, metric, query, opts...))
}

// Retrieve an aggregated report on the calling client's own billable consumption
// by given metric
func (r *PrebillingService) GetTotal(ctx context.Context, metric PrebillingGetTotalParamsMetric, query PrebillingGetTotalParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapBillingReportItem], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("waap/v1/prebilling/total/%v", metric)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve an aggregated report on the calling client's own billable consumption
// by given metric
func (r *PrebillingService) GetTotalAutoPaging(ctx context.Context, metric PrebillingGetTotalParamsMetric, query PrebillingGetTotalParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapBillingReportItem] {
	return pagination.NewOffsetPageAutoPager(r.GetTotal(ctx, metric, query, opts...))
}

type WaapBillingReportItem struct {
	// The amount of the billable resource accrued on the date
	Amount int64 `json:"amount" api:"required"`
	// The date of the accrued billable consumption item
	Date string `json:"date" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Date        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapBillingReportItem) RawJSON() string { return r.JSON.raw }
func (r *WaapBillingReportItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapClientBillingReport struct {
	// The client ID
	ClientID int64 `json:"client_id" api:"required"`
	// The billable consumption items of the client, ordered by date
	Items []WaapBillingReportItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID    respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapClientBillingReport) RawJSON() string { return r.JSON.raw }
func (r *WaapClientBillingReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PrebillingGetBreakdownParams struct {
	// The report start date (inclusive) in ISO 8601 format
	From time.Time `query:"from" api:"required" format:"date" json:"-"`
	// The report end date (exclusive) in ISO 8601 format
	To time.Time `query:"to" api:"required" format:"date" json:"-"`
	// Number of clients to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of clients to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PrebillingGetBreakdownParams]'s query parameters as
// `url.Values`.
func (r PrebillingGetBreakdownParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The metric to get billable consumption for
type PrebillingGetBreakdownParamsMetric string

const (
	PrebillingGetBreakdownParamsMetricDomains       PrebillingGetBreakdownParamsMetric = "domains"
	PrebillingGetBreakdownParamsMetricVolume        PrebillingGetBreakdownParamsMetric = "volume"
	PrebillingGetBreakdownParamsMetricRequests      PrebillingGetBreakdownParamsMetric = "requests"
	PrebillingGetBreakdownParamsMetricRequestsClean PrebillingGetBreakdownParamsMetric = "requests_clean"
)

type PrebillingGetTotalParams struct {
	// The report start date (inclusive) in ISO 8601 format
	From time.Time `query:"from" api:"required" format:"date" json:"-"`
	// The report end date (exclusive) in ISO 8601 format
	To time.Time `query:"to" api:"required" format:"date" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PrebillingGetTotalParams]'s query parameters as
// `url.Values`.
func (r PrebillingGetTotalParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The metric to get billable consumption for
type PrebillingGetTotalParamsMetric string

const (
	PrebillingGetTotalParamsMetricDomains       PrebillingGetTotalParamsMetric = "domains"
	PrebillingGetTotalParamsMetricVolume        PrebillingGetTotalParamsMetric = "volume"
	PrebillingGetTotalParamsMetricRequests      PrebillingGetTotalParamsMetric = "requests"
	PrebillingGetTotalParamsMetricRequestsClean PrebillingGetTotalParamsMetric = "requests_clean"
)
