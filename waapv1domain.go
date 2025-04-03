// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/shared"
	"github.com/tidwall/gjson"
)

// WaapV1DomainService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1DomainService] method instead.
type WaapV1DomainService struct {
	Options         []option.RequestOption
	Policies        *WaapV1DomainPolicyService
	Requests        *WaapV1DomainRequestService
	CustomRules     *WaapV1DomainCustomRuleService
	FirewallRules   *WaapV1DomainFirewallRuleService
	AdvancedRules   *WaapV1DomainAdvancedRuleService
	APIPaths        *WaapV1DomainAPIPathService
	APIDiscovery    *WaapV1DomainAPIDiscoveryService
	Insights        *WaapV1DomainInsightService
	InsightSilences *WaapV1DomainInsightSilenceService
}

// NewWaapV1DomainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWaapV1DomainService(opts ...option.RequestOption) (r *WaapV1DomainService) {
	r = &WaapV1DomainService{}
	r.Options = opts
	r.Policies = NewWaapV1DomainPolicyService(opts...)
	r.Requests = NewWaapV1DomainRequestService(opts...)
	r.CustomRules = NewWaapV1DomainCustomRuleService(opts...)
	r.FirewallRules = NewWaapV1DomainFirewallRuleService(opts...)
	r.AdvancedRules = NewWaapV1DomainAdvancedRuleService(opts...)
	r.APIPaths = NewWaapV1DomainAPIPathService(opts...)
	r.APIDiscovery = NewWaapV1DomainAPIDiscoveryService(opts...)
	r.Insights = NewWaapV1DomainInsightService(opts...)
	r.InsightSilences = NewWaapV1DomainInsightSilenceService(opts...)
	return
}

// Retrieve detailed information about a specific domain
func (r *WaapV1DomainService) Get(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *WaapV1DomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Domain
func (r *WaapV1DomainService) Update(ctx context.Context, domainID int64, body WaapV1DomainUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Retrieve a list of domains associated with the client
func (r *WaapV1DomainService) List(ctx context.Context, query WaapV1DomainListParams, opts ...option.RequestOption) (res *WaapV1DomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an inactive domain by ID. Only domains with status 'bypass' can be
// deleted.
func (r *WaapV1DomainService) Delete(ctx context.Context, domainID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve a list of API path groups for a specific domain
func (r *WaapV1DomainService) GetAPIPathGroups(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *WaapV1DomainGetAPIPathGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-path-groups", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a domain's DDoS attacks
func (r *WaapV1DomainService) GetDdosAttacks(ctx context.Context, domainID int64, query WaapV1DomainGetDdosAttacksParams, opts ...option.RequestOption) (res *WaapV1DomainGetDdosAttacksResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/ddos-attacks", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns the top DDoS counts grouped by URL, User-Agent or IP
func (r *WaapV1DomainService) GetDdosInfo(ctx context.Context, domainID int64, query WaapV1DomainGetDdosInfoParams, opts ...option.RequestOption) (res *WaapV1DomainGetDdosInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/ddos-info", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve an domain's event statistics
func (r *WaapV1DomainService) GetEventStatistics(ctx context.Context, domainID int64, query WaapV1DomainGetEventStatisticsParams, opts ...option.RequestOption) (res *WaapV1DomainGetEventStatisticsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/stats", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve all rule sets linked to a particular domain
func (r *WaapV1DomainService) GetRuleSets(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *[]WaapV1DomainGetRuleSetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/rule-sets", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves a comprehensive report on a domain's traffic statistics based on
// Clickhouse. The report includes details such as API requests, blocked events,
// error counts, and many more traffic-related metrics.
func (r *WaapV1DomainService) GetTraffic(ctx context.Context, domainID int64, query WaapV1DomainGetTrafficParams, opts ...option.RequestOption) (res *[]WaapV1DomainGetTrafficResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/traffic", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The different statuses a domain can have
type DomainStatus string

const (
	DomainStatusActive  DomainStatus = "active"
	DomainStatusBypass  DomainStatus = "bypass"
	DomainStatusMonitor DomainStatus = "monitor"
	DomainStatusLocked  DomainStatus = "locked"
)

func (r DomainStatus) IsKnown() bool {
	switch r {
	case DomainStatusActive, DomainStatusBypass, DomainStatusMonitor, DomainStatusLocked:
		return true
	}
	return false
}

// Represents a WAAP domain, serving as a singular unit within the WAAP service.
//
// Each domain functions autonomously, possessing its own set of rules and
// configurations to manage web application firewall settings and behaviors.
type WaapV1DomainGetResponse struct {
	// The domain ID
	ID int64 `json:"id,required"`
	// The date and time the domain was created in ISO 8601 format
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The ID of the custom page set
	CustomPageSet int64 `json:"custom_page_set,required,nullable"`
	// The domain name
	Name string `json:"name,required"`
	// The different statuses a domain can have
	Status DomainStatus `json:"status,required"`
	// Domain level quotas
	Quotas map[string]WaapV1DomainGetResponseQuota `json:"quotas,nullable"`
	JSON   waapV1DomainGetResponseJSON             `json:"-"`
}

// waapV1DomainGetResponseJSON contains the JSON metadata for the struct
// [WaapV1DomainGetResponse]
type waapV1DomainGetResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	CustomPageSet apijson.Field
	Name          apijson.Field
	Status        apijson.Field
	Quotas        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WaapV1DomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainGetResponseQuota struct {
	// The maximum allowed number of this resource
	Allowed int64 `json:"allowed,required"`
	// The current number of this resource
	Current int64                            `json:"current,required"`
	JSON    waapV1DomainGetResponseQuotaJSON `json:"-"`
}

// waapV1DomainGetResponseQuotaJSON contains the JSON metadata for the struct
// [WaapV1DomainGetResponseQuota]
type waapV1DomainGetResponseQuotaJSON struct {
	Allowed     apijson.Field
	Current     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainGetResponseQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainGetResponseQuotaJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainListResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []WaapV1DomainListResponseResult `json:"results,required"`
	JSON    waapV1DomainListResponseJSON     `json:"-"`
}

// waapV1DomainListResponseJSON contains the JSON metadata for the struct
// [WaapV1DomainListResponse]
type waapV1DomainListResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainListResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a WAAP domain when getting a list of domains.
type WaapV1DomainListResponseResult struct {
	// The domain ID
	ID int64 `json:"id,required"`
	// The date and time the domain was created in ISO 8601 format
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The ID of the custom page set
	CustomPageSet int64 `json:"custom_page_set,required,nullable"`
	// The domain name
	Name string `json:"name,required"`
	// The different statuses a domain can have
	Status DomainStatus                       `json:"status,required"`
	JSON   waapV1DomainListResponseResultJSON `json:"-"`
}

// waapV1DomainListResponseResultJSON contains the JSON metadata for the struct
// [WaapV1DomainListResponseResult]
type waapV1DomainListResponseResultJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	CustomPageSet apijson.Field
	Name          apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WaapV1DomainListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Response model for the API path groups
type WaapV1DomainGetAPIPathGroupsResponse struct {
	// An array of api groups associated with the API path
	APIPathGroups []string                                 `json:"api_path_groups,required"`
	JSON          waapV1DomainGetAPIPathGroupsResponseJSON `json:"-"`
}

// waapV1DomainGetAPIPathGroupsResponseJSON contains the JSON metadata for the
// struct [WaapV1DomainGetAPIPathGroupsResponse]
type waapV1DomainGetAPIPathGroupsResponseJSON struct {
	APIPathGroups apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WaapV1DomainGetAPIPathGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainGetAPIPathGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainGetDdosAttacksResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []WaapV1DomainGetDdosAttacksResponseResult `json:"results,required"`
	JSON    waapV1DomainGetDdosAttacksResponseJSON     `json:"-"`
}

// waapV1DomainGetDdosAttacksResponseJSON contains the JSON metadata for the struct
// [WaapV1DomainGetDdosAttacksResponse]
type waapV1DomainGetDdosAttacksResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainGetDdosAttacksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainGetDdosAttacksResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainGetDdosAttacksResponseResult struct {
	// End time of DDoS attack
	EndTime time.Time `json:"end_time,nullable" format:"date-time"`
	// Start time of DDoS attack
	StartTime time.Time                                    `json:"start_time" format:"date-time"`
	JSON      waapV1DomainGetDdosAttacksResponseResultJSON `json:"-"`
}

// waapV1DomainGetDdosAttacksResponseResultJSON contains the JSON metadata for the
// struct [WaapV1DomainGetDdosAttacksResponseResult]
type waapV1DomainGetDdosAttacksResponseResultJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainGetDdosAttacksResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainGetDdosAttacksResponseResultJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainGetDdosInfoResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []WaapV1DomainGetDdosInfoResponseResult `json:"results,required"`
	JSON    waapV1DomainGetDdosInfoResponseJSON     `json:"-"`
}

// waapV1DomainGetDdosInfoResponseJSON contains the JSON metadata for the struct
// [WaapV1DomainGetDdosInfoResponse]
type waapV1DomainGetDdosInfoResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainGetDdosInfoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainGetDdosInfoResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainGetDdosInfoResponseResult struct {
	// The number of requests made
	Count int64 `json:"count,required"`
	// The value for the grouped by type
	Identity string                                     `json:"identity,required"`
	Type     WaapV1DomainGetDdosInfoResponseResultsType `json:"type,required"`
	JSON     waapV1DomainGetDdosInfoResponseResultJSON  `json:"-"`
}

// waapV1DomainGetDdosInfoResponseResultJSON contains the JSON metadata for the
// struct [WaapV1DomainGetDdosInfoResponseResult]
type waapV1DomainGetDdosInfoResponseResultJSON struct {
	Count       apijson.Field
	Identity    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainGetDdosInfoResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainGetDdosInfoResponseResultJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainGetDdosInfoResponseResultsType string

const (
	WaapV1DomainGetDdosInfoResponseResultsTypeURL       WaapV1DomainGetDdosInfoResponseResultsType = "URL"
	WaapV1DomainGetDdosInfoResponseResultsTypeIP        WaapV1DomainGetDdosInfoResponseResultsType = "IP"
	WaapV1DomainGetDdosInfoResponseResultsTypeUserAgent WaapV1DomainGetDdosInfoResponseResultsType = "User-Agent"
)

func (r WaapV1DomainGetDdosInfoResponseResultsType) IsKnown() bool {
	switch r {
	case WaapV1DomainGetDdosInfoResponseResultsTypeURL, WaapV1DomainGetDdosInfoResponseResultsTypeIP, WaapV1DomainGetDdosInfoResponseResultsTypeUserAgent:
		return true
	}
	return false
}

// A collection of event metrics over a time span
type WaapV1DomainGetEventStatisticsResponse struct {
	// A collection of total numbers of events with blocked results per criteria
	Blocked WaapV1DomainGetEventStatisticsResponseBlocked `json:"blocked,required"`
	// A collection of total numbers of events per criteria
	Count WaapV1DomainGetEventStatisticsResponseCount `json:"count,required"`
	JSON  waapV1DomainGetEventStatisticsResponseJSON  `json:"-"`
}

// waapV1DomainGetEventStatisticsResponseJSON contains the JSON metadata for the
// struct [WaapV1DomainGetEventStatisticsResponse]
type waapV1DomainGetEventStatisticsResponseJSON struct {
	Blocked     apijson.Field
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainGetEventStatisticsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainGetEventStatisticsResponseJSON) RawJSON() string {
	return r.raw
}

// A collection of total numbers of events with blocked results per criteria
type WaapV1DomainGetEventStatisticsResponseBlocked struct {
	// A collection of event counts per action. The first item is the action's
	// abbreviation/full action name, and the second item is the number of events
	Action [][]WaapV1DomainGetEventStatisticsResponseBlockedActionUnion `json:"action,required"`
	// A collection of event counts per country of origin. The first item is the
	// country's ISO 3166-1 alpha-2, and the second item is the number of events
	Country [][]WaapV1DomainGetEventStatisticsResponseBlockedCountryUnion `json:"country,required"`
	// A collection of event counts per organization that owns the event's client IP.
	// The first item is the organization's name, and the second item is the number of
	// events
	Org [][]WaapV1DomainGetEventStatisticsResponseBlockedOrgUnion `json:"org,required"`
	// A collection of event counts per rule that triggered the event. The first item
	// is the rule's name, and the second item is the number of events
	RuleName [][]WaapV1DomainGetEventStatisticsResponseBlockedRuleNameUnion `json:"rule_name,required"`
	JSON     waapV1DomainGetEventStatisticsResponseBlockedJSON              `json:"-"`
}

// waapV1DomainGetEventStatisticsResponseBlockedJSON contains the JSON metadata for
// the struct [WaapV1DomainGetEventStatisticsResponseBlocked]
type waapV1DomainGetEventStatisticsResponseBlockedJSON struct {
	Action      apijson.Field
	Country     apijson.Field
	Org         apijson.Field
	RuleName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainGetEventStatisticsResponseBlocked) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainGetEventStatisticsResponseBlockedJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type WaapV1DomainGetEventStatisticsResponseBlockedActionUnion interface {
	ImplementsWaapV1DomainGetEventStatisticsResponseBlockedActionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WaapV1DomainGetEventStatisticsResponseBlockedActionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type WaapV1DomainGetEventStatisticsResponseBlockedCountryUnion interface {
	ImplementsWaapV1DomainGetEventStatisticsResponseBlockedCountryUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WaapV1DomainGetEventStatisticsResponseBlockedCountryUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type WaapV1DomainGetEventStatisticsResponseBlockedOrgUnion interface {
	ImplementsWaapV1DomainGetEventStatisticsResponseBlockedOrgUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WaapV1DomainGetEventStatisticsResponseBlockedOrgUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type WaapV1DomainGetEventStatisticsResponseBlockedRuleNameUnion interface {
	ImplementsWaapV1DomainGetEventStatisticsResponseBlockedRuleNameUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WaapV1DomainGetEventStatisticsResponseBlockedRuleNameUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// A collection of total numbers of events per criteria
type WaapV1DomainGetEventStatisticsResponseCount struct {
	// A collection of event counts per action. The first item is the action's
	// abbreviation/full action name, and the second item is the number of events
	Action [][]WaapV1DomainGetEventStatisticsResponseCountActionUnion `json:"action,required"`
	// A collection of event counts per country of origin. The first item is the
	// country's ISO 3166-1 alpha-2, and the second item is the number of events
	Country [][]WaapV1DomainGetEventStatisticsResponseCountCountryUnion `json:"country,required"`
	// A collection of event counts per organization that owns the event's client IP.
	// The first item is the organization's name, and the second item is the number of
	// events
	Org [][]WaapV1DomainGetEventStatisticsResponseCountOrgUnion `json:"org,required"`
	// A collection of event counts per rule that triggered the event. The first item
	// is the rule's name, and the second item is the number of events
	RuleName [][]WaapV1DomainGetEventStatisticsResponseCountRuleNameUnion `json:"rule_name,required"`
	JSON     waapV1DomainGetEventStatisticsResponseCountJSON              `json:"-"`
}

// waapV1DomainGetEventStatisticsResponseCountJSON contains the JSON metadata for
// the struct [WaapV1DomainGetEventStatisticsResponseCount]
type waapV1DomainGetEventStatisticsResponseCountJSON struct {
	Action      apijson.Field
	Country     apijson.Field
	Org         apijson.Field
	RuleName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainGetEventStatisticsResponseCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainGetEventStatisticsResponseCountJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type WaapV1DomainGetEventStatisticsResponseCountActionUnion interface {
	ImplementsWaapV1DomainGetEventStatisticsResponseCountActionUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WaapV1DomainGetEventStatisticsResponseCountActionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type WaapV1DomainGetEventStatisticsResponseCountCountryUnion interface {
	ImplementsWaapV1DomainGetEventStatisticsResponseCountCountryUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WaapV1DomainGetEventStatisticsResponseCountCountryUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type WaapV1DomainGetEventStatisticsResponseCountOrgUnion interface {
	ImplementsWaapV1DomainGetEventStatisticsResponseCountOrgUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WaapV1DomainGetEventStatisticsResponseCountOrgUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type WaapV1DomainGetEventStatisticsResponseCountRuleNameUnion interface {
	ImplementsWaapV1DomainGetEventStatisticsResponseCountRuleNameUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WaapV1DomainGetEventStatisticsResponseCountRuleNameUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Represents a custom rule set.
type WaapV1DomainGetRuleSetsResponse struct {
	// Identifier of the rule set.
	ID int64 `json:"id,required"`
	// Detailed description of the rule set.
	Description string `json:"description,required"`
	// Indicates if the rule set is currently active.
	IsActive bool `json:"is_active,required"`
	// Name of the rule set.
	Name string `json:"name,required"`
	// Collection of tags associated with the rule set.
	Tags []WaapV1DomainGetRuleSetsResponseTag `json:"tags,required"`
	// The resource slug associated with the rule set.
	ResourceSlug string                                `json:"resource_slug,nullable"`
	Rules        []WaapV1DomainGetRuleSetsResponseRule `json:"rules"`
	JSON         waapV1DomainGetRuleSetsResponseJSON   `json:"-"`
}

// waapV1DomainGetRuleSetsResponseJSON contains the JSON metadata for the struct
// [WaapV1DomainGetRuleSetsResponse]
type waapV1DomainGetRuleSetsResponseJSON struct {
	ID           apijson.Field
	Description  apijson.Field
	IsActive     apijson.Field
	Name         apijson.Field
	Tags         apijson.Field
	ResourceSlug apijson.Field
	Rules        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WaapV1DomainGetRuleSetsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainGetRuleSetsResponseJSON) RawJSON() string {
	return r.raw
}

// A single tag associated with a rule set.
type WaapV1DomainGetRuleSetsResponseTag struct {
	// Identifier of the tag.
	ID int64 `json:"id,required"`
	// Detailed description of the tag.
	Description string `json:"description,required"`
	// Name of the tag.
	Name string                                 `json:"name,required"`
	JSON waapV1DomainGetRuleSetsResponseTagJSON `json:"-"`
}

// waapV1DomainGetRuleSetsResponseTagJSON contains the JSON metadata for the struct
// [WaapV1DomainGetRuleSetsResponseTag]
type waapV1DomainGetRuleSetsResponseTagJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainGetRuleSetsResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainGetRuleSetsResponseTagJSON) RawJSON() string {
	return r.raw
}

// Represents a configurable WAAP security rule, also known as a policy.
type WaapV1DomainGetRuleSetsResponseRule struct {
	// Unique identifier for the security rule
	ID string `json:"id,required"`
	// The action taken by the WAAP upon rule activation.
	Action WaapV1DomainGetRuleSetsResponseRulesAction `json:"action,required"`
	// Detailed description of the security rule
	Description string `json:"description,required"`
	// The rule set group name to which the rule belongs
	Group string `json:"group,required"`
	// Indicates if the security rule is active
	Mode bool `json:"mode,required"`
	// Name of the security rule
	Name string `json:"name,required"`
	// Identifier of the rule set to which the rule belongs
	RuleSetID int64                                   `json:"rule_set_id,required"`
	JSON      waapV1DomainGetRuleSetsResponseRuleJSON `json:"-"`
}

// waapV1DomainGetRuleSetsResponseRuleJSON contains the JSON metadata for the
// struct [WaapV1DomainGetRuleSetsResponseRule]
type waapV1DomainGetRuleSetsResponseRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Group       apijson.Field
	Mode        apijson.Field
	Name        apijson.Field
	RuleSetID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainGetRuleSetsResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainGetRuleSetsResponseRuleJSON) RawJSON() string {
	return r.raw
}

// The action taken by the WAAP upon rule activation.
type WaapV1DomainGetRuleSetsResponseRulesAction string

const (
	WaapV1DomainGetRuleSetsResponseRulesActionAllow     WaapV1DomainGetRuleSetsResponseRulesAction = "Allow"
	WaapV1DomainGetRuleSetsResponseRulesActionBlock     WaapV1DomainGetRuleSetsResponseRulesAction = "Block"
	WaapV1DomainGetRuleSetsResponseRulesActionCaptcha   WaapV1DomainGetRuleSetsResponseRulesAction = "Captcha"
	WaapV1DomainGetRuleSetsResponseRulesActionGateway   WaapV1DomainGetRuleSetsResponseRulesAction = "Gateway"
	WaapV1DomainGetRuleSetsResponseRulesActionHandshake WaapV1DomainGetRuleSetsResponseRulesAction = "Handshake"
	WaapV1DomainGetRuleSetsResponseRulesActionMonitor   WaapV1DomainGetRuleSetsResponseRulesAction = "Monitor"
	WaapV1DomainGetRuleSetsResponseRulesActionComposite WaapV1DomainGetRuleSetsResponseRulesAction = "Composite"
)

func (r WaapV1DomainGetRuleSetsResponseRulesAction) IsKnown() bool {
	switch r {
	case WaapV1DomainGetRuleSetsResponseRulesActionAllow, WaapV1DomainGetRuleSetsResponseRulesActionBlock, WaapV1DomainGetRuleSetsResponseRulesActionCaptcha, WaapV1DomainGetRuleSetsResponseRulesActionGateway, WaapV1DomainGetRuleSetsResponseRulesActionHandshake, WaapV1DomainGetRuleSetsResponseRulesActionMonitor, WaapV1DomainGetRuleSetsResponseRulesActionComposite:
		return true
	}
	return false
}

// Represents the traffic metrics for a domain at a given time window
type WaapV1DomainGetTrafficResponse struct {
	// UNIX timestamp indicating when the traffic data was recorded
	Timestamp int64 `json:"timestamp,required"`
	// Number of AJAX requests made
	Ajax int64 `json:"ajax"`
	// Number of API requests made
	API int64 `json:"api"`
	// Number of requests allowed through custom rules
	CustomAllowed int64 `json:"customAllowed"`
	// Number of requests blocked due to custom rules
	CustomBlocked int64 `json:"customBlocked"`
	// Number of DDoS attack attempts successfully blocked
	DdosBlocked int64 `json:"ddosBlocked"`
	// Number of requests triggering monitoring actions
	Monitored int64 `json:"monitored"`
	// Number of successful HTTP 2xx responses from the origin server
	Origin2xx int64 `json:"origin2xx"`
	// Number of HTTP 3xx redirects issued by the origin server
	Origin3xx int64 `json:"origin3xx"`
	// Number of HTTP 4xx errors from the origin server
	OriginError4xx int64 `json:"originError4xx"`
	// Number of HTTP 5xx errors from the origin server
	OriginError5xx int64 `json:"originError5xx"`
	// Number of timeouts experienced at the origin server
	OriginTimeout int64 `json:"originTimeout"`
	// Number of requests served directly by the origin server
	PassedToOrigin int64 `json:"passedToOrigin"`
	// Number of requests allowed by security policies
	PolicyAllowed int64 `json:"policyAllowed"`
	// Number of requests blocked by security policies
	PolicyBlocked int64 `json:"policyBlocked"`
	// Average origin server response time in milliseconds
	ResponseTime int64 `json:"responseTime"`
	// Number of static asset requests
	Static int64 `json:"static"`
	// Total number of requests
	Total int64 `json:"total"`
	// Requests resulting in neither blocks nor sanctions
	Uncategorized int64                              `json:"uncategorized"`
	JSON          waapV1DomainGetTrafficResponseJSON `json:"-"`
}

// waapV1DomainGetTrafficResponseJSON contains the JSON metadata for the struct
// [WaapV1DomainGetTrafficResponse]
type waapV1DomainGetTrafficResponseJSON struct {
	Timestamp      apijson.Field
	Ajax           apijson.Field
	API            apijson.Field
	CustomAllowed  apijson.Field
	CustomBlocked  apijson.Field
	DdosBlocked    apijson.Field
	Monitored      apijson.Field
	Origin2xx      apijson.Field
	Origin3xx      apijson.Field
	OriginError4xx apijson.Field
	OriginError5xx apijson.Field
	OriginTimeout  apijson.Field
	PassedToOrigin apijson.Field
	PolicyAllowed  apijson.Field
	PolicyBlocked  apijson.Field
	ResponseTime   apijson.Field
	Static         apijson.Field
	Total          apijson.Field
	Uncategorized  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WaapV1DomainGetTrafficResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainGetTrafficResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainUpdateParams struct {
	// Domain statuses that can be used when updating a domain
	Status param.Field[WaapV1DomainUpdateParamsStatus] `json:"status"`
}

func (r WaapV1DomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Domain statuses that can be used when updating a domain
type WaapV1DomainUpdateParamsStatus string

const (
	WaapV1DomainUpdateParamsStatusActive  WaapV1DomainUpdateParamsStatus = "active"
	WaapV1DomainUpdateParamsStatusMonitor WaapV1DomainUpdateParamsStatus = "monitor"
)

func (r WaapV1DomainUpdateParamsStatus) IsKnown() bool {
	switch r {
	case WaapV1DomainUpdateParamsStatusActive, WaapV1DomainUpdateParamsStatusMonitor:
		return true
	}
	return false
}

type WaapV1DomainListParams struct {
	// Filter domains based on their IDs
	IDs param.Field[[]int64] `query:"ids"`
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Filter domains based on the domain name. Supports '\*' as a wildcard character
	Name param.Field[string] `query:"name"`
	// Number of items to skip
	Offset param.Field[int64] `query:"offset"`
	// Sort the response by given field.
	Ordering param.Field[WaapV1DomainListParamsOrdering] `query:"ordering"`
	// The different statuses a domain can have
	Status param.Field[DomainStatus] `query:"status"`
}

// URLQuery serializes [WaapV1DomainListParams]'s query parameters as `url.Values`.
func (r WaapV1DomainListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort the response by given field.
type WaapV1DomainListParamsOrdering string

const (
	WaapV1DomainListParamsOrderingID             WaapV1DomainListParamsOrdering = "id"
	WaapV1DomainListParamsOrderingName           WaapV1DomainListParamsOrdering = "name"
	WaapV1DomainListParamsOrderingStatus         WaapV1DomainListParamsOrdering = "status"
	WaapV1DomainListParamsOrderingCreatedAt      WaapV1DomainListParamsOrdering = "created_at"
	WaapV1DomainListParamsOrderingMinusID        WaapV1DomainListParamsOrdering = "-id"
	WaapV1DomainListParamsOrderingMinusName      WaapV1DomainListParamsOrdering = "-name"
	WaapV1DomainListParamsOrderingMinusStatus    WaapV1DomainListParamsOrdering = "-status"
	WaapV1DomainListParamsOrderingMinusCreatedAt WaapV1DomainListParamsOrdering = "-created_at"
)

func (r WaapV1DomainListParamsOrdering) IsKnown() bool {
	switch r {
	case WaapV1DomainListParamsOrderingID, WaapV1DomainListParamsOrderingName, WaapV1DomainListParamsOrderingStatus, WaapV1DomainListParamsOrderingCreatedAt, WaapV1DomainListParamsOrderingMinusID, WaapV1DomainListParamsOrderingMinusName, WaapV1DomainListParamsOrderingMinusStatus, WaapV1DomainListParamsOrderingMinusCreatedAt:
		return true
	}
	return false
}

type WaapV1DomainGetDdosAttacksParams struct {
	// Filter attacks up to a specified end date in ISO 8601 format
	EndTime param.Field[time.Time] `query:"end_time" format:"date-time"`
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip
	Offset param.Field[int64] `query:"offset"`
	// Sort the response by given field.
	Ordering param.Field[WaapV1DomainGetDdosAttacksParamsOrdering] `query:"ordering"`
	// Filter attacks starting from a specified date in ISO 8601 format
	StartTime param.Field[time.Time] `query:"start_time" format:"date-time"`
}

// URLQuery serializes [WaapV1DomainGetDdosAttacksParams]'s query parameters as
// `url.Values`.
func (r WaapV1DomainGetDdosAttacksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort the response by given field.
type WaapV1DomainGetDdosAttacksParamsOrdering string

const (
	WaapV1DomainGetDdosAttacksParamsOrderingStartTime      WaapV1DomainGetDdosAttacksParamsOrdering = "start_time"
	WaapV1DomainGetDdosAttacksParamsOrderingMinusStartTime WaapV1DomainGetDdosAttacksParamsOrdering = "-start_time"
	WaapV1DomainGetDdosAttacksParamsOrderingEndTime        WaapV1DomainGetDdosAttacksParamsOrdering = "end_time"
	WaapV1DomainGetDdosAttacksParamsOrderingMinusEndTime   WaapV1DomainGetDdosAttacksParamsOrdering = "-end_time"
)

func (r WaapV1DomainGetDdosAttacksParamsOrdering) IsKnown() bool {
	switch r {
	case WaapV1DomainGetDdosAttacksParamsOrderingStartTime, WaapV1DomainGetDdosAttacksParamsOrderingMinusStartTime, WaapV1DomainGetDdosAttacksParamsOrderingEndTime, WaapV1DomainGetDdosAttacksParamsOrderingMinusEndTime:
		return true
	}
	return false
}

type WaapV1DomainGetDdosInfoParams struct {
	// The identity of the requests to group by
	GroupBy param.Field[WaapV1DomainGetDdosInfoParamsGroupBy] `query:"group_by,required"`
	// Filter traffic starting from a specified date in ISO 8601 format
	Start param.Field[time.Time] `query:"start,required" format:"date-time"`
	// Filter traffic up to a specified end date in ISO 8601 format. If not provided,
	// defaults to the current date and time.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Number of items to skip
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [WaapV1DomainGetDdosInfoParams]'s query parameters as
// `url.Values`.
func (r WaapV1DomainGetDdosInfoParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The identity of the requests to group by
type WaapV1DomainGetDdosInfoParamsGroupBy string

const (
	WaapV1DomainGetDdosInfoParamsGroupByURL       WaapV1DomainGetDdosInfoParamsGroupBy = "URL"
	WaapV1DomainGetDdosInfoParamsGroupByUserAgent WaapV1DomainGetDdosInfoParamsGroupBy = "User-Agent"
	WaapV1DomainGetDdosInfoParamsGroupByIP        WaapV1DomainGetDdosInfoParamsGroupBy = "IP"
)

func (r WaapV1DomainGetDdosInfoParamsGroupBy) IsKnown() bool {
	switch r {
	case WaapV1DomainGetDdosInfoParamsGroupByURL, WaapV1DomainGetDdosInfoParamsGroupByUserAgent, WaapV1DomainGetDdosInfoParamsGroupByIP:
		return true
	}
	return false
}

type WaapV1DomainGetEventStatisticsParams struct {
	// Filter traffic starting from a specified date in ISO 8601 format
	Start param.Field[time.Time] `query:"start,required" format:"date-time"`
	// A list of action names to filter on.
	Action param.Field[[]WaapV1DomainGetEventStatisticsParamsAction] `query:"action"`
	// Filter traffic up to a specified end date in ISO 8601 format. If not provided,
	// defaults to the current date and time.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// A list of IPs to filter event statistics.
	IP param.Field[[]string] `query:"ip" format:"ipvanyaddress"`
	// A list of reference IDs to filter event statistics.
	ReferenceID param.Field[[]string] `query:"reference_id"`
	// A list of results to filter event statistics.
	Result param.Field[[]WaapV1DomainGetEventStatisticsParamsResult] `query:"result"`
}

// URLQuery serializes [WaapV1DomainGetEventStatisticsParams]'s query parameters as
// `url.Values`.
func (r WaapV1DomainGetEventStatisticsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WaapV1DomainGetEventStatisticsParamsAction string

const (
	WaapV1DomainGetEventStatisticsParamsActionBlock     WaapV1DomainGetEventStatisticsParamsAction = "block"
	WaapV1DomainGetEventStatisticsParamsActionCaptcha   WaapV1DomainGetEventStatisticsParamsAction = "captcha"
	WaapV1DomainGetEventStatisticsParamsActionHandshake WaapV1DomainGetEventStatisticsParamsAction = "handshake"
	WaapV1DomainGetEventStatisticsParamsActionMonitor   WaapV1DomainGetEventStatisticsParamsAction = "monitor"
)

func (r WaapV1DomainGetEventStatisticsParamsAction) IsKnown() bool {
	switch r {
	case WaapV1DomainGetEventStatisticsParamsActionBlock, WaapV1DomainGetEventStatisticsParamsActionCaptcha, WaapV1DomainGetEventStatisticsParamsActionHandshake, WaapV1DomainGetEventStatisticsParamsActionMonitor:
		return true
	}
	return false
}

type WaapV1DomainGetEventStatisticsParamsResult string

const (
	WaapV1DomainGetEventStatisticsParamsResultPassed    WaapV1DomainGetEventStatisticsParamsResult = "passed"
	WaapV1DomainGetEventStatisticsParamsResultBlocked   WaapV1DomainGetEventStatisticsParamsResult = "blocked"
	WaapV1DomainGetEventStatisticsParamsResultMonitored WaapV1DomainGetEventStatisticsParamsResult = "monitored"
	WaapV1DomainGetEventStatisticsParamsResultAllowed   WaapV1DomainGetEventStatisticsParamsResult = "allowed"
)

func (r WaapV1DomainGetEventStatisticsParamsResult) IsKnown() bool {
	switch r {
	case WaapV1DomainGetEventStatisticsParamsResultPassed, WaapV1DomainGetEventStatisticsParamsResultBlocked, WaapV1DomainGetEventStatisticsParamsResultMonitored, WaapV1DomainGetEventStatisticsParamsResultAllowed:
		return true
	}
	return false
}

type WaapV1DomainGetTrafficParams struct {
	// Specifies the granularity of the result data.
	Resolution param.Field[WaapV1DomainGetTrafficParamsResolution] `query:"resolution,required"`
	// Filter traffic starting from a specified date in ISO 8601 format
	Start param.Field[time.Time] `query:"start,required" format:"date-time"`
	// Filter traffic up to a specified end date in ISO 8601 format. If not provided,
	// defaults to the current date and time.
	End param.Field[time.Time] `query:"end" format:"date-time"`
}

// URLQuery serializes [WaapV1DomainGetTrafficParams]'s query parameters as
// `url.Values`.
func (r WaapV1DomainGetTrafficParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the granularity of the result data.
type WaapV1DomainGetTrafficParamsResolution string

const (
	WaapV1DomainGetTrafficParamsResolutionDaily    WaapV1DomainGetTrafficParamsResolution = "daily"
	WaapV1DomainGetTrafficParamsResolutionHourly   WaapV1DomainGetTrafficParamsResolution = "hourly"
	WaapV1DomainGetTrafficParamsResolutionMinutely WaapV1DomainGetTrafficParamsResolution = "minutely"
)

func (r WaapV1DomainGetTrafficParamsResolution) IsKnown() bool {
	switch r {
	case WaapV1DomainGetTrafficParamsResolutionDaily, WaapV1DomainGetTrafficParamsResolutionHourly, WaapV1DomainGetTrafficParamsResolutionMinutely:
		return true
	}
	return false
}
