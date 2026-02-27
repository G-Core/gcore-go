// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fastedge

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

// Apps are descriptions of edge apps, that reference the binary and may contain
// app-specific settings, such as environment variables.
//
// AppLogService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppLogService] method instead.
type AppLogService struct {
	Options []option.RequestOption
}

// NewAppLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppLogService(opts ...option.RequestOption) (r AppLogService) {
	r = AppLogService{}
	r.Options = opts
	return
}

// List logs for the app
func (r *AppLogService) List(ctx context.Context, id int64, query AppLogListParams, opts ...option.RequestOption) (res *pagination.OffsetPageFastedgeAppLogs[Log], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("fastedge/v1/apps/%v/logs", id)
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

// List logs for the app
func (r *AppLogService) ListAutoPaging(ctx context.Context, id int64, query AppLogListParams, opts ...option.RequestOption) *pagination.OffsetPageFastedgeAppLogsAutoPager[Log] {
	return pagination.NewOffsetPageFastedgeAppLogsAutoPager(r.List(ctx, id, query, opts...))
}

type Log struct {
	// Id of the log
	ID string `json:"id"`
	// Name of the application
	AppName string `json:"app_name"`
	// Client IP
	ClientIP string `json:"client_ip"`
	// Edge name
	Edge string `json:"edge"`
	// Log message
	Log string `json:"log"`
	// Timestamp of a log in RFC3339 format
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AppName     respjson.Field
		ClientIP    respjson.Field
		Edge        respjson.Field
		Log         respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Log) RawJSON() string { return r.JSON.raw }
func (r *Log) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppLogListParams struct {
	// Search by client IP
	ClientIP param.Opt[string] `query:"client_ip,omitzero" format:"ipv4" json:"-"`
	// Edge name
	Edge param.Opt[string] `query:"edge,omitzero" format:"string" json:"-"`
	// Reporting period start time, RFC3339 format. Default 1 hour ago.
	From param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	// Limit for pagination
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset for pagination
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Search string
	Search param.Opt[string] `query:"search,omitzero" format:"string" json:"-"`
	// Reporting period end time, RFC3339 format. Default current time in UTC.
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Sort order (default desc)
	//
	// Any of "desc", "asc".
	Sort AppLogListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AppLogListParams]'s query parameters as `url.Values`.
func (r AppLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort order (default desc)
type AppLogListParamsSort string

const (
	AppLogListParamsSortDesc AppLogListParamsSort = "desc"
	AppLogListParamsSortAsc  AppLogListParamsSort = "asc"
)
