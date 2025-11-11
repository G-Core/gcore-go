// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// AuditLogService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuditLogService] method instead.
type AuditLogService struct {
	Options []option.RequestOption
}

// NewAuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuditLogService(opts ...option.RequestOption) (r AuditLogService) {
	r = AuditLogService{}
	r.Options = opts
	return
}

// Get information about all CDN activity logs records.
func (r *AuditLogService) List(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[CdnAuditLogEntry], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cdn/activity_log/requests"
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

// Get information about all CDN activity logs records.
func (r *AuditLogService) ListAutoPaging(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[CdnAuditLogEntry] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Get information about CDN activity logs record.
func (r *AuditLogService) Get(ctx context.Context, logID int64, opts ...option.RequestOption) (res *CdnAuditLogEntry, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/activity_log/requests/%v", logID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CdnAuditLogEntry struct {
	// Activity logs record ID.
	ID int64 `json:"id"`
	// State of a requested object before and after the request.
	Actions []CdnAuditLogEntryAction `json:"actions"`
	// ID of the client who made the request.
	ClientID int64 `json:"client_id"`
	// Request body.
	Data any `json:"data"`
	// Host from which the request was made.
	Host string `json:"host"`
	// Request HTTP method.
	Method string `json:"method"`
	// Request URL.
	Path string `json:"path"`
	// Request parameters.
	QueryParams string `json:"query_params"`
	// IP address from which the request was made.
	RemoteIPAddress string `json:"remote_ip_address"`
	// Date and time when the request was made.
	RequestedAt string `json:"requested_at"`
	// Status code that is returned in the response.
	StatusCode int64 `json:"status_code"`
	// Permanent API token ID with which the request was made.
	TokenID int64 `json:"token_id"`
	// ID of the user who made the request.
	UserID int64 `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Actions         respjson.Field
		ClientID        respjson.Field
		Data            respjson.Field
		Host            respjson.Field
		Method          respjson.Field
		Path            respjson.Field
		QueryParams     respjson.Field
		RemoteIPAddress respjson.Field
		RequestedAt     respjson.Field
		StatusCode      respjson.Field
		TokenID         respjson.Field
		UserID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnAuditLogEntry) RawJSON() string { return r.JSON.raw }
func (r *CdnAuditLogEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdnAuditLogEntryAction struct {
	// Type of change.
	//
	// Possible values:
	//
	// - **D** - Object is deleted.
	// - **C** - Object is created.
	// - **U** - Object is updated.
	ActionType string `json:"action_type"`
	// JSON representation of object after the request.
	StateAfterRequest any `json:"state_after_request"`
	// JSON representation of object before the request.
	StateBeforeRequest any `json:"state_before_request"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType         respjson.Field
		StateAfterRequest  respjson.Field
		StateBeforeRequest respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnAuditLogEntryAction) RawJSON() string { return r.JSON.raw }
func (r *CdnAuditLogEntryAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditLogListParams struct {
	// Client ID.
	ClientID param.Opt[int64] `query:"client_id,omitzero" json:"-"`
	// Maximum number of items in response.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// End of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	//
	// You can specify a date with a time separated by a space, or just a date.
	//
	// Examples:
	//
	// - &`max_requested_at`=2021-05-05 12:00:00
	// - &`max_requested_at`=2021-05-05
	MaxRequestedAt param.Opt[string] `query:"max_requested_at,omitzero" json:"-"`
	// HTTP method type of requests.
	//
	// Use upper case only.
	//
	// Example:
	//
	// - ?method=DELETE
	Method param.Opt[string] `query:"method,omitzero" json:"-"`
	// Beginning of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	//
	// You can specify a date with a time separated by a space, or just a date.
	//
	// Examples:
	//
	// - &`min_requested_at`=2021-05-05 12:00:00
	// - &`min_requested_at`=2021-05-05
	MinRequestedAt param.Opt[string] `query:"min_requested_at,omitzero" json:"-"`
	// Offset relative to the beginning of activity logs.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Exact URL path.
	Path param.Opt[string] `query:"path,omitzero" json:"-"`
	// Exact IP address from which requests are sent.
	RemoteIPAddress param.Opt[string] `query:"remote_ip_address,omitzero" json:"-"`
	// Status code returned in the response.
	//
	// Specify the first numbers of a status code to get requests for a group of status
	// codes.
	//
	// To filter the activity logs by 4xx codes, use:
	//
	// - &`status_code`=4 -
	StatusCode param.Opt[int64] `query:"status_code,omitzero" json:"-"`
	// Permanent API token ID. Requests made with this token should be displayed.
	TokenID param.Opt[int64] `query:"token_id,omitzero" json:"-"`
	// User ID.
	UserID param.Opt[int64] `query:"user_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuditLogListParams]'s query parameters as `url.Values`.
func (r AuditLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
