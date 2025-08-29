// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// ZoneRrsetService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneRrsetService] method instead.
type ZoneRrsetService struct {
	Options []option.RequestOption
}

// NewZoneRrsetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneRrsetService(opts ...option.RequestOption) (r ZoneRrsetService) {
	r = ZoneRrsetService{}
	r.Options = opts
	return
}

// Add the RRSet to the zone specified by zoneName, RRSets can be configured to be
// either dynamic or static.
//
// ### Static RRsets
//
// Staticly configured RRSets provide DNS responses as is.
//
// ### Dynamic RRsets
//
// Dynamic RRSets have picker configuration defined thus it's possible to finely
// customize DNS response. Picking rules are defined on the RRSet level as a list
// of selectors, filters and mutators. Picker considers different resource records
// metadata, requestor IP, and other event-feeds like monitoring. Picker
// configuration is an ordered list defined by "pickers" attribute. Requestor IP is
// determined by EDNS Client Subnet (ECS) if defined, otherwise - by
// client/recursor IP. Selector pickers are used in the specified order until the
// first match, in case of match - all next selectors are bypassed. Filters or
// mutators are applied to the match according to the order they are specified. For
// example, sort records by proximity to user, shuffle based on weights and return
// not more than 3:
// `"pickers": [ { "type": "geodistance" }, { "type": "`weighted_shuffle`" }, { "type": "`first_n`", "limit": 3 } ]`
//
// #### geodns filter
//
// A resource record is included in the answer if resource record's metadata
// matches requestor info. For each resource record in RRSet, the following
// metadata is considered (in the order specified):
//
//   - `ip` - list of network addresses in CIDR format, e.g.
//     `["192.168.15.150/25", "2003:de:2016::/48"]`;
//   - `asn` - list of autonomous system numbers, e.g. `[1234, 5678]`;
//   - `regions` - list of region codes, e.g. `["de-bw", "de-by"]`;
//   - `countries` - list of country codes, e.g. `["de", "lu", "lt"]`;
//   - `continents` - list of continent codes, e.g.
//     `["af", "an", "eu", "as", "na", "sa", "oc"]`. If there is a record (or
//     multiple) with metadata matched IP, it's used as a response. If not - asn,
//     then country and then continent are checked for a match. If there is no match,
//     then the behaviour is defined by _strict_ parameter of the filter. Example:
//     `"pickers": [ { "type": "geodns", "strict": true } ]`
//
// ##### Strict parameter
//
// `strict: true` means that if no records percolate through the geodns filter it
// returns no answers. `strict: false` means that if no records percolate through
// the geodns filter, all records are passed over.
//
// #### asn selector
//
// Resource records which ASN metadata matches ASN of the requestor are picked by
// this selector, and passed to the next non-selector picker, if there is no
// match - next configured picker starts with all records. Example:
// `"pickers": [ {"type": "asn"} ]`
//
// #### country selector
//
// Resource records which country metadata matches country of the requestor are
// picked by this selector, and passed to the next non-selector picker, if there is
// no match - next configured picker starts with all records. Example:
// `"pickers": [ { "type": "country" } ]`
//
// #### continent selector
//
// Resource records which continent metadata matches continent of the requestor are
// picked by this selector, and passed to the next non-selector picker, if there is
// no match - next configured picker starts with all records. Example:
// `"pickers": [ { "type": "continent" } ]`
//
// #### region selector
//
// Resource records which region metadata matches region of the requestor are
// picked by this selector, and passed to the next non-selector picker, if there is
// no match - next configured picker starts with all records. e.g. `fr-nor` for
// France/Normandy. Example: `"pickers": [ { "type": "region" } ]`
//
// #### ip selector
//
// Resource records which IP metadata matches IP of the requestor are picked by
// this selector, and passed to the next non-selector picker, if there is no
// match - next configured picker starts with all records. Maximum 100 subnets are
// allowed to specify in meta of RR. Example: `"pickers": [ { "type": "ip" } ]`
//
// #### default selector
//
// When enabled, records marked as default are selected:
// `"meta": {"default": true}`. Example:
// `"pickers": [ { "type": "geodns", "strict": false }, { "type": "default" }, { "type": "`first_n`", "limit": 2 } ]`
//
// #### geodistance mutator
//
// The resource records are rearranged in ascending order based on the distance (in
// meters) from requestor to the coordinates specified in latlong metadata.
// Distance is calculated using Haversine formula. The "nearest" to the user's IP
// RR goes first. The records without latlong metadata come last. e.g. for Berlin
// `[52.520008, 13.404954]`.; In this configuration the only "nearest" to the
// requestor record to be returned:
// `"pickers": [ { "type": "geodistance" }, { "type": "`first_n`", "limit": 1 } ]`
//
// #### `weighted_shuffle` mutator
//
// The resource records are rearranged in random order based on the `weight`
// metadata. Default weight (if not specified) is 50. Example:
// `"pickers": [ { "type": "`weighted_shuffle`" } ]`
//
// #### `first_n` filter
//
// Slices first N (N specified as a limit parameter value) resource records.
// Example: `"pickers": [ { "type": "`first_n`", "limit": 1 } ]` returns only the
// first resource record.
//
// ##### limit parameter
//
// Can be a positive value for a specific limit. Use zero or leave it blank to
// indicate no limits.
func (r *ZoneRrsetService) New(ctx context.Context, rrsetType string, params ZoneRrsetNewParams, opts ...option.RequestOption) (res *DNSOutputRrset, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if params.ZoneName == "" {
		err = errors.New("missing required zoneName parameter")
		return
	}
	if params.RrsetName == "" {
		err = errors.New("missing required rrsetName parameter")
		return
	}
	if rrsetType == "" {
		err = errors.New("missing required rrsetType parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s/%s/%s", params.ZoneName, params.RrsetName, rrsetType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create/update RRset.
func (r *ZoneRrsetService) Update(ctx context.Context, rrsetType string, params ZoneRrsetUpdateParams, opts ...option.RequestOption) (res *DNSOutputRrset, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if params.ZoneName == "" {
		err = errors.New("missing required zoneName parameter")
		return
	}
	if params.RrsetName == "" {
		err = errors.New("missing required rrsetName parameter")
		return
	}
	if rrsetType == "" {
		err = errors.New("missing required rrsetType parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s/%s/%s", params.ZoneName, params.RrsetName, rrsetType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// List of RRset.
func (r *ZoneRrsetService) List(ctx context.Context, zoneName string, query ZoneRrsetListParams, opts ...option.RequestOption) (res *ZoneRrsetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if zoneName == "" {
		err = errors.New("missing required zoneName parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s/rrsets", zoneName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete RRset.
func (r *ZoneRrsetService) Delete(ctx context.Context, rrsetType string, body ZoneRrsetDeleteParams, opts ...option.RequestOption) (res *ZoneRrsetDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if body.ZoneName == "" {
		err = errors.New("missing required zoneName parameter")
		return
	}
	if body.RrsetName == "" {
		err = errors.New("missing required rrsetName parameter")
		return
	}
	if rrsetType == "" {
		err = errors.New("missing required rrsetType parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s/%s/%s", body.ZoneName, body.RrsetName, rrsetType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Particular RRset item info
func (r *ZoneRrsetService) Get(ctx context.Context, rrsetType string, query ZoneRrsetGetParams, opts ...option.RequestOption) (res *DNSOutputRrset, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if query.ZoneName == "" {
		err = errors.New("missing required zoneName parameter")
		return
	}
	if query.RrsetName == "" {
		err = errors.New("missing required rrsetName parameter")
		return
	}
	if rrsetType == "" {
		err = errors.New("missing required rrsetType parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s/%s/%s", query.ZoneName, query.RrsetName, rrsetType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get failover history for the RRset
func (r *ZoneRrsetService) GetFailoverLogs(ctx context.Context, rrsetType string, params ZoneRrsetGetFailoverLogsParams, opts ...option.RequestOption) (res *ZoneRrsetGetFailoverLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	if params.ZoneName == "" {
		err = errors.New("missing required zoneName parameter")
		return
	}
	if params.RrsetName == "" {
		err = errors.New("missing required rrsetName parameter")
		return
	}
	if rrsetType == "" {
		err = errors.New("missing required rrsetType parameter")
		return
	}
	path := fmt.Sprintf("dns/v2/zones/%s/%s/%s/failover/log", params.ZoneName, params.RrsetName, rrsetType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type DNSFailoverLog []DNSFailoverLogItem

// FailoverLogEntry
type DNSFailoverLogItem struct {
	Action  string `json:"action"`
	Address string `json:"address"`
	Time    int64  `json:"time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Address     respjson.Field
		Time        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSFailoverLogItem) RawJSON() string { return r.JSON.raw }
func (r *DNSFailoverLogItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DNSOutputRrset struct {
	Name string `json:"name,required"`
	// List of resource record from rrset
	ResourceRecords []DNSOutputRrsetResourceRecord `json:"resource_records,required"`
	// RRSet type
	//
	// Any of "A", "AAAA", "NS", "CNAME", "MX", "TXT", "SRV", "SOA".
	Type        DNSOutputRrsetType `json:"type,required"`
	FilterSetID int64              `json:"filter_set_id"`
	// Meta information for rrset. Map with string key and any valid json as value,
	// with valid keys
	//
	//  1. `failover` (object, beta feature, might be changed in the future) can have
	//     fields 1.1. `protocol` (string, required, HTTP, TCP, UDP, ICMP) 1.2. `port`
	//     (int, required, 1-65535) 1.3. `frequency` (int, required, in seconds 10-3600)
	//     1.4. `timeout` (int, required, in seconds 1-10), 1.5. `method` (string, only
	//     for protocol=HTTP) 1.6. `command` (string, bytes to be sent only for
	//     protocol=TCP/UDP) 1.7. `url` (string, only for protocol=HTTP) 1.8. `tls`
	//     (bool, only for protocol=HTTP) 1.9. `regexp` (string regex to match, only for
	//     non-ICMP) 1.10. `http_status_code` (int, only for protocol=HTTP) 1.11. `host`
	//     (string, only for protocol=HTTP)
	//  2. `geodns_link` (string) - name of the geodns link to use, if previously set,
	//     must re-send when updating or CDN integration will be removed for this RRSet
	Meta map[string]any `json:"meta"`
	// Set of pickers
	Pickers []DNSOutputRrsetPicker `json:"pickers"`
	Ttl     int64                  `json:"ttl"`
	// Timestamp marshals/unmarshals date and time as timestamp in json
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Warning about some possible side effects without strictly disallowing operations
	// on rrset readonly Deprecated: use Warnings instead
	Warning string `json:"warning"`
	// Warning about some possible side effects without strictly disallowing operations
	// on rrset readonly
	Warnings []DNSOutputRrsetWarning `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name            respjson.Field
		ResourceRecords respjson.Field
		Type            respjson.Field
		FilterSetID     respjson.Field
		Meta            respjson.Field
		Pickers         respjson.Field
		Ttl             respjson.Field
		UpdatedAt       respjson.Field
		Warning         respjson.Field
		Warnings        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSOutputRrset) RawJSON() string { return r.JSON.raw }
func (r *DNSOutputRrset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DNSOutputRrsetResourceRecord struct {
	// Content of resource record The exact length of the array depends on the type of
	// rrset, each individual record parameter must be a separate element of the array.
	// For example
	//
	//   - SRV-record: `[100, 1, 5061, "example.com"]`
	//   - CNAME-record: `[ "the.target.domain" ]`
	//   - A-record: `[ "1.2.3.4", "5.6.7.8" ]`
	//   - AAAA-record: `[ "2001:db8::1", "2001:db8::2" ]`
	//   - MX-record: `[ "mail1.example.com", "mail2.example.com" ]`
	//   - SVCB/HTTPS-record:
	//     `[ 1, ".", ["alpn", "h3", "h2"], [ "port", 1443 ], [ "ipv4hint", "10.0.0.1" ], [ "ech", "AEn+DQBFKwAgACABWIHUGj4u+PIggYXcR5JF0gYk3dCRioBW8uJq9H4mKAAIAAEAAQABAANAEnB1YmxpYy50bHMtZWNoLmRldgAA" ] ]`
	Content []any `json:"content,required"`
	ID      int64 `json:"id"`
	Enabled bool  `json:"enabled"`
	// Meta information for record Map with string key and any valid json as value,
	// with valid keys
	//
	//  1. `asn` (array of int)
	//  2. `continents` (array of string)
	//  3. `countries` (array of string)
	//  4. `latlong` (array of float64, latitude and longitude)
	//  5. `fallback` (bool)
	//  6. `backup` (bool)
	//  7. `notes` (string)
	//  8. `weight` (float)
	//  9. `ip` (string) Some keys are reserved for balancing, @see
	//     https://api.gcore.com/dns/v2/info/meta This meta will be used to decide which
	//     resource record should pass through filters from the filter set
	Meta map[string]any `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ID          respjson.Field
		Enabled     respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSOutputRrsetResourceRecord) RawJSON() string { return r.JSON.raw }
func (r *DNSOutputRrsetResourceRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RRSet type
type DNSOutputRrsetType string

const (
	DNSOutputRrsetTypeA     DNSOutputRrsetType = "A"
	DNSOutputRrsetTypeAaaa  DNSOutputRrsetType = "AAAA"
	DNSOutputRrsetTypeNs    DNSOutputRrsetType = "NS"
	DNSOutputRrsetTypeCname DNSOutputRrsetType = "CNAME"
	DNSOutputRrsetTypeMx    DNSOutputRrsetType = "MX"
	DNSOutputRrsetTypeTxt   DNSOutputRrsetType = "TXT"
	DNSOutputRrsetTypeSrv   DNSOutputRrsetType = "SRV"
	DNSOutputRrsetTypeSoa   DNSOutputRrsetType = "SOA"
)

type DNSOutputRrsetPicker struct {
	// Filter type
	//
	// Any of "geodns", "asn", "country", "continent", "region", "ip", "geodistance",
	// "weighted_shuffle", "default", "first_n".
	Type string `json:"type,required"`
	// Limits the number of records returned by the filter Can be a positive value for
	// a specific limit. Use zero or leave it blank to indicate no limits.
	Limit int64 `json:"limit"`
	// if strict=false, then the filter will return all records if no records match the
	// filter
	Strict bool `json:"strict"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Limit       respjson.Field
		Strict      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSOutputRrsetPicker) RawJSON() string { return r.JSON.raw }
func (r *DNSOutputRrsetPicker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DNSOutputRrsetWarning struct {
	Key     string `json:"key"`
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSOutputRrsetWarning) RawJSON() string { return r.JSON.raw }
func (r *DNSOutputRrsetWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRrsetListResponse struct {
	Rrsets      []DNSOutputRrset `json:"rrsets"`
	TotalAmount int64            `json:"total_amount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rrsets      respjson.Field
		TotalAmount respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneRrsetListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneRrsetListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRrsetDeleteResponse = any

type ZoneRrsetGetFailoverLogsResponse struct {
	// FailoverLog
	Log         DNSFailoverLog `json:"log"`
	TotalAmount int64          `json:"total_amount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Log         respjson.Field
		TotalAmount respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneRrsetGetFailoverLogsResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneRrsetGetFailoverLogsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRrsetNewParams struct {
	ZoneName  string `path:"zoneName,required" json:"-"`
	RrsetName string `path:"rrsetName,required" json:"-"`
	// List of resource record from rrset
	ResourceRecords []ZoneRrsetNewParamsResourceRecord `json:"resource_records,omitzero,required"`
	Ttl             param.Opt[int64]                   `json:"ttl,omitzero"`
	// Meta information for rrset
	Meta map[string]any `json:"meta,omitzero"`
	// Set of pickers
	Pickers []ZoneRrsetNewParamsPicker `json:"pickers,omitzero"`
	paramObj
}

func (r ZoneRrsetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneRrsetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneRrsetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// nolint: lll
//
// The property Content is required.
type ZoneRrsetNewParamsResourceRecord struct {
	// Content of resource record The exact length of the array depends on the type of
	// rrset, each individual record parameter must be a separate element of the array.
	// For example
	//
	//   - SRV-record: `[100, 1, 5061, "example.com"]`
	//   - CNAME-record: `[ "the.target.domain" ]`
	//   - A-record: `[ "1.2.3.4", "5.6.7.8" ]`
	//   - AAAA-record: `[ "2001:db8::1", "2001:db8::2" ]`
	//   - MX-record: `[ "mail1.example.com", "mail2.example.com" ]`
	//   - SVCB/HTTPS-record:
	//     `[ 1, ".", ["alpn", "h3", "h2"], [ "port", 1443 ], [ "ipv4hint", "10.0.0.1" ], [ "ech", "AEn+DQBFKwAgACABWIHUGj4u+PIggYXcR5JF0gYk3dCRioBW8uJq9H4mKAAIAAEAAQABAANAEnB1YmxpYy50bHMtZWNoLmRldgAA" ] ]`
	Content []any           `json:"content,omitzero,required"`
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// This meta will be used to decide which resource record should pass through
	// filters from the filter set
	Meta map[string]any `json:"meta,omitzero"`
	paramObj
}

func (r ZoneRrsetNewParamsResourceRecord) MarshalJSON() (data []byte, err error) {
	type shadow ZoneRrsetNewParamsResourceRecord
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneRrsetNewParamsResourceRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ZoneRrsetNewParamsPicker struct {
	// Filter type
	//
	// Any of "geodns", "asn", "country", "continent", "region", "ip", "geodistance",
	// "weighted_shuffle", "default", "first_n".
	Type string `json:"type,omitzero,required"`
	// Limits the number of records returned by the filter Can be a positive value for
	// a specific limit. Use zero or leave it blank to indicate no limits.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// if strict=false, then the filter will return all records if no records match the
	// filter
	Strict param.Opt[bool] `json:"strict,omitzero"`
	paramObj
}

func (r ZoneRrsetNewParamsPicker) MarshalJSON() (data []byte, err error) {
	type shadow ZoneRrsetNewParamsPicker
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneRrsetNewParamsPicker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ZoneRrsetNewParamsPicker](
		"type", "geodns", "asn", "country", "continent", "region", "ip", "geodistance", "weighted_shuffle", "default", "first_n",
	)
}

type ZoneRrsetUpdateParams struct {
	ZoneName  string `path:"zoneName,required" json:"-"`
	RrsetName string `path:"rrsetName,required" json:"-"`
	// List of resource record from rrset
	ResourceRecords []ZoneRrsetUpdateParamsResourceRecord `json:"resource_records,omitzero,required"`
	Ttl             param.Opt[int64]                      `json:"ttl,omitzero"`
	// Meta information for rrset
	Meta map[string]any `json:"meta,omitzero"`
	// Set of pickers
	Pickers []ZoneRrsetUpdateParamsPicker `json:"pickers,omitzero"`
	paramObj
}

func (r ZoneRrsetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneRrsetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneRrsetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// nolint: lll
//
// The property Content is required.
type ZoneRrsetUpdateParamsResourceRecord struct {
	// Content of resource record The exact length of the array depends on the type of
	// rrset, each individual record parameter must be a separate element of the array.
	// For example
	//
	//   - SRV-record: `[100, 1, 5061, "example.com"]`
	//   - CNAME-record: `[ "the.target.domain" ]`
	//   - A-record: `[ "1.2.3.4", "5.6.7.8" ]`
	//   - AAAA-record: `[ "2001:db8::1", "2001:db8::2" ]`
	//   - MX-record: `[ "mail1.example.com", "mail2.example.com" ]`
	//   - SVCB/HTTPS-record:
	//     `[ 1, ".", ["alpn", "h3", "h2"], [ "port", 1443 ], [ "ipv4hint", "10.0.0.1" ], [ "ech", "AEn+DQBFKwAgACABWIHUGj4u+PIggYXcR5JF0gYk3dCRioBW8uJq9H4mKAAIAAEAAQABAANAEnB1YmxpYy50bHMtZWNoLmRldgAA" ] ]`
	Content []any           `json:"content,omitzero,required"`
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// This meta will be used to decide which resource record should pass through
	// filters from the filter set
	Meta map[string]any `json:"meta,omitzero"`
	paramObj
}

func (r ZoneRrsetUpdateParamsResourceRecord) MarshalJSON() (data []byte, err error) {
	type shadow ZoneRrsetUpdateParamsResourceRecord
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneRrsetUpdateParamsResourceRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ZoneRrsetUpdateParamsPicker struct {
	// Filter type
	//
	// Any of "geodns", "asn", "country", "continent", "region", "ip", "geodistance",
	// "weighted_shuffle", "default", "first_n".
	Type string `json:"type,omitzero,required"`
	// Limits the number of records returned by the filter Can be a positive value for
	// a specific limit. Use zero or leave it blank to indicate no limits.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// if strict=false, then the filter will return all records if no records match the
	// filter
	Strict param.Opt[bool] `json:"strict,omitzero"`
	paramObj
}

func (r ZoneRrsetUpdateParamsPicker) MarshalJSON() (data []byte, err error) {
	type shadow ZoneRrsetUpdateParamsPicker
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneRrsetUpdateParamsPicker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ZoneRrsetUpdateParamsPicker](
		"type", "geodns", "asn", "country", "continent", "region", "ip", "geodistance", "weighted_shuffle", "default", "first_n",
	)
}

type ZoneRrsetListParams struct {
	// Max number of records in response
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Amount of records to skip before beginning to write in response.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Field name to sort by
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// Ascending or descending order
	//
	// Any of "asc", "desc".
	OrderDirection ZoneRrsetListParamsOrderDirection `query:"order_direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneRrsetListParams]'s query parameters as `url.Values`.
func (r ZoneRrsetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Ascending or descending order
type ZoneRrsetListParamsOrderDirection string

const (
	ZoneRrsetListParamsOrderDirectionAsc  ZoneRrsetListParamsOrderDirection = "asc"
	ZoneRrsetListParamsOrderDirectionDesc ZoneRrsetListParamsOrderDirection = "desc"
)

type ZoneRrsetDeleteParams struct {
	ZoneName  string `path:"zoneName,required" json:"-"`
	RrsetName string `path:"rrsetName,required" json:"-"`
	paramObj
}

type ZoneRrsetGetParams struct {
	ZoneName  string `path:"zoneName,required" json:"-"`
	RrsetName string `path:"rrsetName,required" json:"-"`
	paramObj
}

type ZoneRrsetGetFailoverLogsParams struct {
	ZoneName  string `path:"zoneName,required" json:"-"`
	RrsetName string `path:"rrsetName,required" json:"-"`
	// Max number of records in response
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Amount of records to skip before beginning to write in response.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneRrsetGetFailoverLogsParams]'s query parameters as
// `url.Values`.
func (r ZoneRrsetGetFailoverLogsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
