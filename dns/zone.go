// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	shimjson "github.com/G-Core/gcore-go/internal/encoding/json"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// ZoneService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneService] method instead.
type ZoneService struct {
	Options []option.RequestOption
	Dnssec  ZoneDnssecService
	Rrsets  ZoneRrsetService
}

// NewZoneService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZoneService(opts ...option.RequestOption) (r ZoneService) {
	r = ZoneService{}
	r.Options = opts
	r.Dnssec = NewZoneDnssecService(opts...)
	r.Rrsets = NewZoneRrsetService(opts...)
	return
}

// Add DNS zone.
func (r *ZoneService) New(ctx context.Context, body ZoneNewParams, opts ...option.RequestOption) (res *ZoneNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dns/v2/zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Show created zones with pagination managed by limit and offset params. All query
// params are optional.
func (r *ZoneService) List(ctx context.Context, query ZoneListParams, opts ...option.RequestOption) (res *ZoneListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dns/v2/zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete DNS zone and its records and raws.
func (r *ZoneService) Delete(ctx context.Context, name string, opts ...option.RequestOption) (res *ZoneDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns delegation status for specified domain name. This endpoint has rate
// limit.
func (r *ZoneService) CheckDelegationStatus(ctx context.Context, name string, opts ...option.RequestOption) (res *ZoneCheckDelegationStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/analyze/%s/delegation-status", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Disable DNS zone.
func (r *ZoneService) Disable(ctx context.Context, name string, opts ...option.RequestOption) (res *ZoneDisableResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s/disable", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Enable DNS zone.
func (r *ZoneService) Enable(ctx context.Context, name string, opts ...option.RequestOption) (res *ZoneEnableResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s/enable", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Export zone to bind9 format.
func (r *ZoneService) Export(ctx context.Context, zoneName string, opts ...option.RequestOption) (res *ZoneExportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneName == "" {
		err = errors.New("missing required zoneName parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s/export", zoneName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Zone info by zone name.
func (r *ZoneService) Get(ctx context.Context, name string, opts ...option.RequestOption) (res *ZoneGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Statistics of DNS zone in common and by record types. To get summary statistics
// for all zones use `all` instead of zone name in path. Note: Consumption
// statistics is updated in near real-time as a standard practice. However, the
// frequency of updates can vary, but they are typically available within a 30
// minutes period. Exceptions, such as maintenance periods, may delay data beyond
// 30 minutes until servers resume and backfill missing statistics.
func (r *ZoneService) GetStatistics(ctx context.Context, name string, query ZoneGetStatisticsParams, opts ...option.RequestOption) (res *ZoneGetStatisticsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s/statistics", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Import zone in bind9 format.
func (r *ZoneService) Import(ctx context.Context, zoneName string, body ZoneImportParams, opts ...option.RequestOption) (res *ZoneImportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneName == "" {
		err = errors.New("missing required zoneName parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s/import", zoneName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update DNS zone and SOA record.
func (r *ZoneService) Replace(ctx context.Context, name string, body ZoneReplaceParams, opts ...option.RequestOption) (res *ZoneReplaceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// NameServer
type DNSNameServer struct {
	Ipv4Addresses []string `json:"ipv4Addresses"`
	Ipv6Addresses []string `json:"ipv6Addresses"`
	Name          string   `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ipv4Addresses respjson.Field
		Ipv6Addresses respjson.Field
		Name          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSNameServer) RawJSON() string { return r.JSON.raw }
func (r *DNSNameServer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneNewResponse struct {
	ID       int64    `json:"id"`
	Warnings []string `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneListResponse struct {
	TotalAmount int64                  `json:"total_amount"`
	Zones       []ZoneListResponseZone `json:"zones"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TotalAmount respjson.Field
		Zones       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputZone
type ZoneListResponseZone struct {
	// ID of zone. This field usually is omitted in response and available only in case
	// of getting deleted zones by admin.
	ID       int64 `json:"id"`
	ClientID int64 `json:"client_id"`
	// email address of the administrator responsible for this zone
	Contact string `json:"contact"`
	// describe dnssec status true means dnssec is enabled for the zone false means
	// dnssec is disabled for the zone
	DnssecEnabled bool `json:"dnssec_enabled"`
	// number of seconds after which secondary name servers should stop answering
	// request for this zone
	Expiry int64 `json:"expiry"`
	// arbitrarily data of zone in json format
	Meta any `json:"meta"`
	// name of DNS zone
	Name string `json:"name"`
	// Time To Live of cache
	NxTtl int64 `json:"nx_ttl"`
	// primary master name server for zone
	PrimaryServer string                       `json:"primary_server"`
	Records       []ZoneListResponseZoneRecord `json:"records"`
	// number of seconds after which secondary name servers should query the master for
	// the SOA record, to detect zone changes.
	Refresh int64 `json:"refresh"`
	// number of seconds after which secondary name servers should retry to request the
	// serial number
	Retry        int64                            `json:"retry"`
	RrsetsAmount ZoneListResponseZoneRrsetsAmount `json:"rrsets_amount"`
	// Serial number for this zone or Timestamp of zone modification moment. If a
	// secondary name server slaved to this one observes an increase in this number,
	// the slave will assume that the zone has been updated and initiate a zone
	// transfer.
	Serial int64  `json:"serial"`
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ClientID      respjson.Field
		Contact       respjson.Field
		DnssecEnabled respjson.Field
		Expiry        respjson.Field
		Meta          respjson.Field
		Name          respjson.Field
		NxTtl         respjson.Field
		PrimaryServer respjson.Field
		Records       respjson.Field
		Refresh       respjson.Field
		Retry         respjson.Field
		RrsetsAmount  respjson.Field
		Serial        respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneListResponseZone) RawJSON() string { return r.JSON.raw }
func (r *ZoneListResponseZone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Record - readonly short version of rrset
type ZoneListResponseZoneRecord struct {
	Name         string   `json:"name"`
	ShortAnswers []string `json:"short_answers"`
	Ttl          int64    `json:"ttl"`
	Type         string   `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		ShortAnswers respjson.Field
		Ttl          respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneListResponseZoneRecord) RawJSON() string { return r.JSON.raw }
func (r *ZoneListResponseZoneRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneListResponseZoneRrsetsAmount struct {
	// Amount of dynamic RRsets in zone
	Dynamic ZoneListResponseZoneRrsetsAmountDynamic `json:"dynamic"`
	// Amount of static RRsets in zone
	Static int64 `json:"static"`
	// Total amount of RRsets in zone
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dynamic     respjson.Field
		Static      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneListResponseZoneRrsetsAmount) RawJSON() string { return r.JSON.raw }
func (r *ZoneListResponseZoneRrsetsAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Amount of dynamic RRsets in zone
type ZoneListResponseZoneRrsetsAmountDynamic struct {
	// Amount of RRsets with enabled healthchecks
	Healthcheck int64 `json:"healthcheck"`
	// Total amount of dynamic RRsets in zone
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Healthcheck respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneListResponseZoneRrsetsAmountDynamic) RawJSON() string { return r.JSON.raw }
func (r *ZoneListResponseZoneRrsetsAmountDynamic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDeleteResponse = any

type ZoneCheckDelegationStatusResponse struct {
	AuthoritativeNameServers []DNSNameServer `json:"authoritative_name_servers"`
	GcoreAuthorizedCount     int64           `json:"gcore_authorized_count"`
	IsWhitelabelDelegation   bool            `json:"is_whitelabel_delegation"`
	NonGcoreAuthorizedCount  int64           `json:"non_gcore_authorized_count"`
	ZoneExists               bool            `json:"zone_exists"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthoritativeNameServers respjson.Field
		GcoreAuthorizedCount     respjson.Field
		IsWhitelabelDelegation   respjson.Field
		NonGcoreAuthorizedCount  respjson.Field
		ZoneExists               respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneCheckDelegationStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneCheckDelegationStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDisableResponse = any

type ZoneEnableResponse = any

type ZoneExportResponse struct {
	RawZone string `json:"raw_zone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RawZone     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneExportResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneExportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete zone info with all records included
type ZoneGetResponse struct {
	// OutputZone
	Zone ZoneGetResponseZone `json:"Zone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Zone        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputZone
type ZoneGetResponseZone struct {
	// ID of zone. This field usually is omitted in response and available only in case
	// of getting deleted zones by admin.
	ID       int64 `json:"id"`
	ClientID int64 `json:"client_id"`
	// email address of the administrator responsible for this zone
	Contact string `json:"contact"`
	// describe dnssec status true means dnssec is enabled for the zone false means
	// dnssec is disabled for the zone
	DnssecEnabled bool `json:"dnssec_enabled"`
	// number of seconds after which secondary name servers should stop answering
	// request for this zone
	Expiry int64 `json:"expiry"`
	// arbitrarily data of zone in json format
	Meta any `json:"meta"`
	// name of DNS zone
	Name string `json:"name"`
	// Time To Live of cache
	NxTtl int64 `json:"nx_ttl"`
	// primary master name server for zone
	PrimaryServer string                      `json:"primary_server"`
	Records       []ZoneGetResponseZoneRecord `json:"records"`
	// number of seconds after which secondary name servers should query the master for
	// the SOA record, to detect zone changes.
	Refresh int64 `json:"refresh"`
	// number of seconds after which secondary name servers should retry to request the
	// serial number
	Retry        int64                           `json:"retry"`
	RrsetsAmount ZoneGetResponseZoneRrsetsAmount `json:"rrsets_amount"`
	// Serial number for this zone or Timestamp of zone modification moment. If a
	// secondary name server slaved to this one observes an increase in this number,
	// the slave will assume that the zone has been updated and initiate a zone
	// transfer.
	Serial int64  `json:"serial"`
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ClientID      respjson.Field
		Contact       respjson.Field
		DnssecEnabled respjson.Field
		Expiry        respjson.Field
		Meta          respjson.Field
		Name          respjson.Field
		NxTtl         respjson.Field
		PrimaryServer respjson.Field
		Records       respjson.Field
		Refresh       respjson.Field
		Retry         respjson.Field
		RrsetsAmount  respjson.Field
		Serial        respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneGetResponseZone) RawJSON() string { return r.JSON.raw }
func (r *ZoneGetResponseZone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Record - readonly short version of rrset
type ZoneGetResponseZoneRecord struct {
	Name         string   `json:"name"`
	ShortAnswers []string `json:"short_answers"`
	Ttl          int64    `json:"ttl"`
	Type         string   `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		ShortAnswers respjson.Field
		Ttl          respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneGetResponseZoneRecord) RawJSON() string { return r.JSON.raw }
func (r *ZoneGetResponseZoneRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGetResponseZoneRrsetsAmount struct {
	// Amount of dynamic RRsets in zone
	Dynamic ZoneGetResponseZoneRrsetsAmountDynamic `json:"dynamic"`
	// Amount of static RRsets in zone
	Static int64 `json:"static"`
	// Total amount of RRsets in zone
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dynamic     respjson.Field
		Static      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneGetResponseZoneRrsetsAmount) RawJSON() string { return r.JSON.raw }
func (r *ZoneGetResponseZoneRrsetsAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Amount of dynamic RRsets in zone
type ZoneGetResponseZoneRrsetsAmountDynamic struct {
	// Amount of RRsets with enabled healthchecks
	Healthcheck int64 `json:"healthcheck"`
	// Total amount of dynamic RRsets in zone
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Healthcheck respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneGetResponseZoneRrsetsAmountDynamic) RawJSON() string { return r.JSON.raw }
func (r *ZoneGetResponseZoneRrsetsAmountDynamic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// StatisticsZoneResponse
type ZoneGetStatisticsResponse struct {
	// Requests amount (values) for particular zone fractionated by time intervals
	// (keys). Example of response:
	// `{ "requests": { "1598608080000": 14716, "1598608140000": 51167, "1598608200000": 53432, "1598611020000": 51050, "1598611080000": 52611, "1598611140000": 46884 } }`
	Requests any `json:"requests"`
	// Total - sum of all values
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Requests    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneGetStatisticsResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneGetStatisticsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneImportResponse struct {
	// ImportedRRSets - import statistics
	Imported ZoneImportResponseImported   `json:"imported"`
	Success  bool                         `json:"success"`
	Warnings map[string]map[string]string `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Imported    respjson.Field
		Success     respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneImportResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneImportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ImportedRRSets - import statistics
type ZoneImportResponseImported struct {
	Qtype                  int64 `json:"qtype"`
	ResourceRecords        int64 `json:"resource_records"`
	Rrsets                 int64 `json:"rrsets"`
	SkippedResourceRecords int64 `json:"skipped_resource_records"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Qtype                  respjson.Field
		ResourceRecords        respjson.Field
		Rrsets                 respjson.Field
		SkippedResourceRecords respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneImportResponseImported) RawJSON() string { return r.JSON.raw }
func (r *ZoneImportResponseImported) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneReplaceResponse = any

type ZoneNewParams struct {
	// name of DNS zone
	Name string `json:"name,required"`
	// email address of the administrator responsible for this zone
	Contact param.Opt[string] `json:"contact,omitzero"`
	// If a zone is disabled, then its records will not be resolved on dns servers
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// number of seconds after which secondary name servers should stop answering
	// request for this zone
	Expiry param.Opt[int64] `json:"expiry,omitzero"`
	// Time To Live of cache
	NxTtl param.Opt[int64] `json:"nx_ttl,omitzero"`
	// primary master name server for zone
	PrimaryServer param.Opt[string] `json:"primary_server,omitzero"`
	// number of seconds after which secondary name servers should query the master for
	// the SOA record, to detect zone changes.
	Refresh param.Opt[int64] `json:"refresh,omitzero"`
	// number of seconds after which secondary name servers should retry to request the
	// serial number
	Retry param.Opt[int64] `json:"retry,omitzero"`
	// Serial number for this zone or Timestamp of zone modification moment. If a
	// secondary name server slaved to this one observes an increase in this number,
	// the slave will assume that the zone has been updated and initiate a zone
	// transfer.
	Serial param.Opt[int64] `json:"serial,omitzero"`
	// arbitrarily data of zone in json format you can specify `webhook` url and
	// `webhook_method` here webhook will get a map with three arrays: for created,
	// updated and deleted rrsets `webhook_method` can be omitted, POST will be used by
	// default
	Meta map[string]any `json:"meta,omitzero"`
	paramObj
}

func (r ZoneNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneListParams struct {
	CaseSensitive param.Opt[bool] `query:"case_sensitive,omitzero" json:"-"`
	// Zones with dynamic RRsets
	Dynamic    param.Opt[bool] `query:"dynamic,omitzero" json:"-"`
	Enabled    param.Opt[bool] `query:"enabled,omitzero" json:"-"`
	ExactMatch param.Opt[bool] `query:"exact_match,omitzero" json:"-"`
	// Zones with RRsets that have healthchecks
	Healthcheck param.Opt[bool] `query:"healthcheck,omitzero" json:"-"`
	// Max number of records in response
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Amount of records to skip before beginning to write in response.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Field name to sort by
	OrderBy       param.Opt[string]    `query:"order_by,omitzero" json:"-"`
	Status        param.Opt[string]    `query:"status,omitzero" json:"-"`
	UpdatedAtFrom param.Opt[time.Time] `query:"updated_at_from,omitzero" format:"date-time" json:"-"`
	UpdatedAtTo   param.Opt[time.Time] `query:"updated_at_to,omitzero" format:"date-time" json:"-"`
	// to pass several ids `id=1&id=3&id=5...`
	ID []int64 `query:"id,omitzero" json:"-"`
	// to pass several `client_ids` `client_id=1&`client_id`=3&`client_id`=5...`
	ClientID      []int64 `query:"client_id,omitzero" json:"-"`
	IamResellerID []int64 `query:"iam_reseller_id,omitzero" json:"-"`
	// to pass several names `name=first&name=second...`
	Name []string `query:"name,omitzero" json:"-"`
	// Ascending or descending order
	//
	// Any of "asc", "desc".
	OrderDirection ZoneListParamsOrderDirection `query:"order_direction,omitzero" json:"-"`
	ResellerID     []int64                      `query:"reseller_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneListParams]'s query parameters as `url.Values`.
func (r ZoneListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Ascending or descending order
type ZoneListParamsOrderDirection string

const (
	ZoneListParamsOrderDirectionAsc  ZoneListParamsOrderDirection = "asc"
	ZoneListParamsOrderDirectionDesc ZoneListParamsOrderDirection = "desc"
)

type ZoneGetStatisticsParams struct {
	// Beginning of the requested time period (Unix Timestamp, UTC.) In a query string:
	// &from=1709068637
	From param.Opt[int64] `query:"from,omitzero" json:"-"`
	// Granularity parameter string is a sequence of decimal numbers, each with
	// optional fraction and a unit suffix, such as "300ms", "1.5h" or "2h45m". Valid
	// time units are "s", "m", "h".
	Granularity param.Opt[string] `query:"granularity,omitzero" json:"-"`
	// DNS record type. Possible values:
	//
	// - A
	// - AAAA
	// - NS
	// - CNAME
	// - MX
	// - TXT
	// - SVCB
	// - HTTPS
	RecordType param.Opt[string] `query:"record_type,omitzero" json:"-"`
	// End of the requested time period (Unix Timestamp, UTC.) In a query string:
	// &to=1709673437
	To param.Opt[int64] `query:"to,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneGetStatisticsParams]'s query parameters as
// `url.Values`.
func (r ZoneGetStatisticsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ZoneImportParams struct {
	// Read reads up to len(p) bytes into p. It returns the number of bytes read (0 <=
	// n <= len(p)) and any error encountered. Even if Read returns n < len(p), it may
	// use all of p as scratch space during the call. If some data is available but not
	// len(p) bytes, Read conventionally returns what is available instead of waiting
	// for more. When Read encounters an error or end-of-file condition after
	// successfully reading n > 0 bytes, it returns the number of bytes read. It may
	// return the (non-nil) error from the same call or return the error (and n == 0)
	// from a subsequent call. An instance of this general case is that a Reader
	// returning a non-zero number of bytes at the end of the input stream may return
	// either err == EOF or err == nil. The next Read should return 0, EOF. Callers
	// should always process the n > 0 bytes returned before considering the error err.
	// Doing so correctly handles I/O errors that happen after reading some bytes and
	// also both of the allowed EOF behaviors. If len(p) == 0, Read should always
	// return n == 0. It may return a non-nil error if some error condition is known,
	// such as EOF. Implementations of Read are discouraged from returning a zero byte
	// count with a nil error, except when len(p) == 0. Callers should treat a return
	// of 0 and nil as indicating that nothing happened; in particular it does not
	// indicate EOF. Implementations must not retain p.
	Body any
	paramObj
}

func (r ZoneImportParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ZoneImportParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type ZoneReplaceParams struct {
	// name of DNS zone
	Name string `json:"name,required"`
	// email address of the administrator responsible for this zone
	Contact param.Opt[string] `json:"contact,omitzero"`
	// If a zone is disabled, then its records will not be resolved on dns servers
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// number of seconds after which secondary name servers should stop answering
	// request for this zone
	Expiry param.Opt[int64] `json:"expiry,omitzero"`
	// Time To Live of cache
	NxTtl param.Opt[int64] `json:"nx_ttl,omitzero"`
	// primary master name server for zone
	PrimaryServer param.Opt[string] `json:"primary_server,omitzero"`
	// number of seconds after which secondary name servers should query the master for
	// the SOA record, to detect zone changes.
	Refresh param.Opt[int64] `json:"refresh,omitzero"`
	// number of seconds after which secondary name servers should retry to request the
	// serial number
	Retry param.Opt[int64] `json:"retry,omitzero"`
	// Serial number for this zone or Timestamp of zone modification moment. If a
	// secondary name server slaved to this one observes an increase in this number,
	// the slave will assume that the zone has been updated and initiate a zone
	// transfer.
	Serial param.Opt[int64] `json:"serial,omitzero"`
	// arbitrarily data of zone in json format you can specify `webhook` url and
	// `webhook_method` here webhook will get a map with three arrays: for created,
	// updated and deleted rrsets `webhook_method` can be omitted, POST will be used by
	// default
	Meta map[string]any `json:"meta,omitzero"`
	paramObj
}

func (r ZoneReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
