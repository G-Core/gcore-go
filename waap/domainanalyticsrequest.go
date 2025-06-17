// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"errors"
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

// DomainAnalyticsRequestService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainAnalyticsRequestService] method instead.
type DomainAnalyticsRequestService struct {
	Options []option.RequestOption
}

// NewDomainAnalyticsRequestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDomainAnalyticsRequestService(opts ...option.RequestOption) (r DomainAnalyticsRequestService) {
	r = DomainAnalyticsRequestService{}
	r.Options = opts
	return
}

// Retrieve a domain's requests data.
func (r *DomainAnalyticsRequestService) List(ctx context.Context, domainID int64, query DomainAnalyticsRequestListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapRequestSummary], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/requests", domainID)
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

// Retrieve a domain's requests data.
func (r *DomainAnalyticsRequestService) ListAutoPaging(ctx context.Context, domainID int64, query DomainAnalyticsRequestListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapRequestSummary] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, domainID, query, opts...))
}

// Retrieves all the available information for a request that matches a given
// request id
func (r *DomainAnalyticsRequestService) Get(ctx context.Context, requestID string, query DomainAnalyticsRequestGetParams, opts ...option.RequestOption) (res *WaapRequestDetails, err error) {
	opts = append(r.Options[:], opts...)
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/requests/%s/details", query.DomainID, requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DomainAnalyticsRequestListParams struct {
	// Filter traffic starting from a specified date in ISO 8601 format
	Start time.Time `query:"start,required" format:"date-time" json:"-"`
	// Filter traffic up to a specified end date in ISO 8601 format. If not provided,
	// defaults to the current date and time.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date-time" json:"-"`
	// Filter the response by IP.
	IP param.Opt[string] `query:"ip,omitzero" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Sort the response by given field.
	Ordering param.Opt[string] `query:"ordering,omitzero" json:"-"`
	// Filter the response by reference ID.
	ReferenceID param.Opt[string] `query:"reference_id,omitzero" json:"-"`
	// Filter the response by security rule name.
	SecurityRuleName param.Opt[string] `query:"security_rule_name,omitzero" json:"-"`
	// Filter the response by response code.
	StatusCode param.Opt[int64] `query:"status_code,omitzero" json:"-"`
	// Filter the response by actions.
	//
	// Any of "allow", "block", "captcha", "handshake".
	Actions []string `query:"actions,omitzero" json:"-"`
	// Filter the response by country codes in ISO 3166-1 alpha-2 format.
	Countries []string `query:"countries,omitzero" json:"-"`
	// Filter the response by traffic types.
	TrafficTypes []WaapTrafficType `query:"traffic_types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainAnalyticsRequestListParams]'s query parameters as
// `url.Values`.
func (r DomainAnalyticsRequestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DomainAnalyticsRequestGetParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}
