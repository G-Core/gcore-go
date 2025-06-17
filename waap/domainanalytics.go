// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
)

// DomainAnalyticsService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainAnalyticsService] method instead.
type DomainAnalyticsService struct {
	Options  []option.RequestOption
	Requests DomainAnalyticsRequestService
}

// NewDomainAnalyticsService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDomainAnalyticsService(opts ...option.RequestOption) (r DomainAnalyticsService) {
	r = DomainAnalyticsService{}
	r.Options = opts
	r.Requests = NewDomainAnalyticsRequestService(opts...)
	return
}

// Retrieve an domain's event statistics
func (r *DomainAnalyticsService) GetEventStatistics(ctx context.Context, domainID int64, query DomainAnalyticsGetEventStatisticsParams, opts ...option.RequestOption) (res *WaapEventStatistics, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/stats", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a domain's DDoS attacks
func (r *DomainAnalyticsService) ListDDOSAttacks(ctx context.Context, domainID int64, query DomainAnalyticsListDDOSAttacksParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapDDOSAttack], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/ddos-attacks", domainID)
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

// Retrieve a domain's DDoS attacks
func (r *DomainAnalyticsService) ListDDOSAttacksAutoPaging(ctx context.Context, domainID int64, query DomainAnalyticsListDDOSAttacksParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapDDOSAttack] {
	return pagination.NewOffsetPageAutoPager(r.ListDDOSAttacks(ctx, domainID, query, opts...))
}

// Returns the top DDoS counts grouped by URL, User-Agent or IP
func (r *DomainAnalyticsService) ListDDOSInfo(ctx context.Context, domainID int64, query DomainAnalyticsListDDOSInfoParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapDDOSInfo], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/ddos-info", domainID)
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

// Returns the top DDoS counts grouped by URL, User-Agent or IP
func (r *DomainAnalyticsService) ListDDOSInfoAutoPaging(ctx context.Context, domainID int64, query DomainAnalyticsListDDOSInfoParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapDDOSInfo] {
	return pagination.NewOffsetPageAutoPager(r.ListDDOSInfo(ctx, domainID, query, opts...))
}

// Retrieves a comprehensive report on a domain's traffic statistics based on
// Clickhouse. The report includes details such as API requests, blocked events,
// error counts, and many more traffic-related metrics.
func (r *DomainAnalyticsService) ListEventTraffic(ctx context.Context, domainID int64, query DomainAnalyticsListEventTrafficParams, opts ...option.RequestOption) (res *[]WaapTrafficMetrics, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/traffic", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DomainAnalyticsGetEventStatisticsParams struct {
	// Filter traffic starting from a specified date in ISO 8601 format
	Start time.Time `query:"start,required" format:"date-time" json:"-"`
	// Filter traffic up to a specified end date in ISO 8601 format. If not provided,
	// defaults to the current date and time.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date-time" json:"-"`
	// A list of action names to filter on.
	//
	// Any of "block", "captcha", "handshake", "monitor".
	Action []string `query:"action,omitzero" json:"-"`
	// A list of IPs to filter event statistics.
	IP []string `query:"ip,omitzero" format:"ipvanyaddress" json:"-"`
	// A list of reference IDs to filter event statistics.
	ReferenceID []string `query:"reference_id,omitzero" json:"-"`
	// A list of results to filter event statistics.
	//
	// Any of "passed", "blocked", "monitored", "allowed".
	Result []string `query:"result,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainAnalyticsGetEventStatisticsParams]'s query parameters
// as `url.Values`.
func (r DomainAnalyticsGetEventStatisticsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DomainAnalyticsListDDOSAttacksParams struct {
	// Filter attacks up to a specified end date in ISO 8601 format
	EndTime param.Opt[time.Time] `query:"end_time,omitzero" format:"date-time" json:"-"`
	// Filter attacks starting from a specified date in ISO 8601 format
	StartTime param.Opt[time.Time] `query:"start_time,omitzero" format:"date-time" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Sort the response by given field.
	//
	// Any of "start_time", "-start_time", "end_time", "-end_time".
	Ordering DomainAnalyticsListDDOSAttacksParamsOrdering `query:"ordering,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainAnalyticsListDDOSAttacksParams]'s query parameters as
// `url.Values`.
func (r DomainAnalyticsListDDOSAttacksParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort the response by given field.
type DomainAnalyticsListDDOSAttacksParamsOrdering string

const (
	DomainAnalyticsListDDOSAttacksParamsOrderingStartTime      DomainAnalyticsListDDOSAttacksParamsOrdering = "start_time"
	DomainAnalyticsListDDOSAttacksParamsOrderingMinusStartTime DomainAnalyticsListDDOSAttacksParamsOrdering = "-start_time"
	DomainAnalyticsListDDOSAttacksParamsOrderingEndTime        DomainAnalyticsListDDOSAttacksParamsOrdering = "end_time"
	DomainAnalyticsListDDOSAttacksParamsOrderingMinusEndTime   DomainAnalyticsListDDOSAttacksParamsOrdering = "-end_time"
)

type DomainAnalyticsListDDOSInfoParams struct {
	// The identity of the requests to group by
	//
	// Any of "URL", "User-Agent", "IP".
	GroupBy DomainAnalyticsListDDOSInfoParamsGroupBy `query:"group_by,omitzero,required" json:"-"`
	// Filter traffic starting from a specified date in ISO 8601 format
	Start time.Time `query:"start,required" format:"date-time" json:"-"`
	// Filter traffic up to a specified end date in ISO 8601 format. If not provided,
	// defaults to the current date and time.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date-time" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainAnalyticsListDDOSInfoParams]'s query parameters as
// `url.Values`.
func (r DomainAnalyticsListDDOSInfoParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The identity of the requests to group by
type DomainAnalyticsListDDOSInfoParamsGroupBy string

const (
	DomainAnalyticsListDDOSInfoParamsGroupByURL       DomainAnalyticsListDDOSInfoParamsGroupBy = "URL"
	DomainAnalyticsListDDOSInfoParamsGroupByUserAgent DomainAnalyticsListDDOSInfoParamsGroupBy = "User-Agent"
	DomainAnalyticsListDDOSInfoParamsGroupByIP        DomainAnalyticsListDDOSInfoParamsGroupBy = "IP"
)

type DomainAnalyticsListEventTrafficParams struct {
	// Specifies the granularity of the result data.
	//
	// Any of "daily", "hourly", "minutely".
	Resolution WaapResolution `query:"resolution,omitzero,required" json:"-"`
	// Filter traffic starting from a specified date in ISO 8601 format
	Start time.Time `query:"start,required" format:"date-time" json:"-"`
	// Filter traffic up to a specified end date in ISO 8601 format. If not provided,
	// defaults to the current date and time.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [DomainAnalyticsListEventTrafficParams]'s query parameters
// as `url.Values`.
func (r DomainAnalyticsListEventTrafficParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
