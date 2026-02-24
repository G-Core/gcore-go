// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// IPInfoMetricService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPInfoMetricService] method instead.
type IPInfoMetricService struct {
	Options []option.RequestOption
}

// NewIPInfoMetricService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIPInfoMetricService(opts ...option.RequestOption) (r IPInfoMetricService) {
	r = IPInfoMetricService{}
	r.Options = opts
	return
}

// Retrieve metrics encompassing the counts of total requests, blocked requests and
// unique sessions associated with a specified IP address. Metrics provide a
// statistical overview, aiding in analyzing the interaction and access patterns of
// the IP address in context.
func (r *IPInfoMetricService) List(ctx context.Context, query IPInfoMetricListParams, opts ...option.RequestOption) (res *WaapIPInfoCounts, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/ip-info/counts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WaapIPInfoCounts struct {
	// The number of requests from the IP address that were blocked
	BlockedRequests int64 `json:"blocked_requests" api:"required"`
	// The total number of requests made by the IP address
	TotalRequests int64 `json:"total_requests" api:"required"`
	// The number of unique sessions from the IP address
	UniqueSessions int64 `json:"unique_sessions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockedRequests respjson.Field
		TotalRequests   respjson.Field
		UniqueSessions  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapIPInfoCounts) RawJSON() string { return r.JSON.raw }
func (r *WaapIPInfoCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPInfoMetricListParams struct {
	// The IP address to check
	IP string `query:"ip" api:"required" format:"ipv4" json:"-"`
	// The identifier for a domain. When specified, the response will exclusively
	// contain data pertinent to the indicated domain, filtering out information from
	// other domains.
	DomainID param.Opt[int64] `query:"domain_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IPInfoMetricListParams]'s query parameters as `url.Values`.
func (r IPInfoMetricListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
