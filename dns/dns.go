// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

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

// DNSService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSService] method instead.
type DNSService struct {
	Options         []option.RequestOption
	Locations       LocationService
	Metrics         MetricService
	Pickers         PickerService
	Zones           ZoneService
	NetworkMappings NetworkMappingService
}

// NewDNSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSService(opts ...option.RequestOption) (r DNSService) {
	r = DNSService{}
	r.Options = opts
	r.Locations = NewLocationService(opts...)
	r.Metrics = NewMetricService(opts...)
	r.Pickers = NewPickerService(opts...)
	r.Zones = NewZoneService(opts...)
	r.NetworkMappings = NewNetworkMappingService(opts...)
	return
}

// Get info about client
func (r *DNSService) GetAccountOverview(ctx context.Context, opts ...option.RequestOption) (res *DNSGetAccountOverviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dns/v2/platform/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the dns records from a specific domain or ip.
func (r *DNSService) Lookup(ctx context.Context, query DNSLookupParams, opts ...option.RequestOption) (res *[]DNSLookupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dns/v2/lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DNSGetAccountOverviewResponse struct {
	Info DNSGetAccountOverviewResponseInfo `json:"Info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSGetAccountOverviewResponse) RawJSON() string { return r.JSON.raw }
func (r *DNSGetAccountOverviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DNSGetAccountOverviewResponseInfo struct {
	Contact     string `json:"contact"`
	NameServer1 string `json:"name_server_1"`
	NameServer2 string `json:"name_server_2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contact     respjson.Field
		NameServer1 respjson.Field
		NameServer2 respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSGetAccountOverviewResponseInfo) RawJSON() string { return r.JSON.raw }
func (r *DNSGetAccountOverviewResponseInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DNSLookupResponse struct {
	Content []string `json:"content"`
	Name    string   `json:"name"`
	Ttl     int64    `json:"ttl"`
	Type    string   `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Name        respjson.Field
		Ttl         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSLookupResponse) RawJSON() string { return r.JSON.raw }
func (r *DNSLookupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DNSLookupParams struct {
	// Domain name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Server that will be used as resolver
	//
	// Any of "authoritative_dns", "google", "cloudflare", "open_dns", "quad9",
	// "gcore".
	RequestServer DNSLookupParamsRequestServer `query:"request_server,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DNSLookupParams]'s query parameters as `url.Values`.
func (r DNSLookupParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Server that will be used as resolver
type DNSLookupParamsRequestServer string

const (
	DNSLookupParamsRequestServerAuthoritativeDNS DNSLookupParamsRequestServer = "authoritative_dns"
	DNSLookupParamsRequestServerGoogle           DNSLookupParamsRequestServer = "google"
	DNSLookupParamsRequestServerCloudflare       DNSLookupParamsRequestServer = "cloudflare"
	DNSLookupParamsRequestServerOpenDNS          DNSLookupParamsRequestServer = "open_dns"
	DNSLookupParamsRequestServerQuad9            DNSLookupParamsRequestServer = "quad9"
	DNSLookupParamsRequestServerGcore            DNSLookupParamsRequestServer = "gcore"
)
