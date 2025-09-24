// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
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

// Example of success response:
//
// ```
// # HELP healthcheck_state The `healthcheck_state` metric reflects the state of a specific monitor after conducting a health check
// # TYPE healthcheck_state gauge
// healthcheck_state{client_id="1",monitor_id="431",monitor_locations="us-east-1,us-west-1",monitor_name="test-monitor-1",monitor_type="http",rrset_name="rrset-name1",rrset_type="rrset-type1",zone_name="zone-name1"} 0
// healthcheck_state{client_id="1",monitor_id="4871",monitor_locations="fr-1,fr-2",monitor_name="test-monitor-2",monitor_type="tcp",rrset_name="rrset-name2",rrset_type="rrset-type2",zone_name="zone-name2"} 1
// healthcheck_state{client_id="2",monitor_id="7123",monitor_locations="ua-1,ua-2",monitor_name="test-monitor-3",monitor_type="icmp",rrset_name="rrset-name3",rrset_type="rrset-type3",zone_name="zone-name3"} 0
// ```
func (r *MetricService) List(ctx context.Context, query MetricListParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "plain/text")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "dns/v2/monitor/metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type MetricListParams struct {
	// Admin and technical user can specify `client_id` to get metrics for particular
	// client. Ignored for client
	ClientIDs []int64 `query:"client_ids,omitzero" json:"-"`
	// Admin and technical user can specify `monitor_id` to get metrics for particular
	// zone. Ignored for client
	ZoneNames []string `query:"zone_names,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MetricListParams]'s query parameters as `url.Values`.
func (r MetricListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
