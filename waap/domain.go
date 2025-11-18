// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"errors"
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

// DomainService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainService] method instead.
type DomainService struct {
	Options         []option.RequestOption
	Settings        DomainSettingService
	APIPaths        DomainAPIPathService
	APIPathGroups   DomainAPIPathGroupService
	APIDiscovery    DomainAPIDiscoveryService
	Insights        DomainInsightService
	InsightSilences DomainInsightSilenceService
	Statistics      DomainStatisticService
	CustomRules     DomainCustomRuleService
	FirewallRules   DomainFirewallRuleService
	AdvancedRules   DomainAdvancedRuleService
}

// NewDomainService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomainService(opts ...option.RequestOption) (r DomainService) {
	r = DomainService{}
	r.Options = opts
	r.Settings = NewDomainSettingService(opts...)
	r.APIPaths = NewDomainAPIPathService(opts...)
	r.APIPathGroups = NewDomainAPIPathGroupService(opts...)
	r.APIDiscovery = NewDomainAPIDiscoveryService(opts...)
	r.Insights = NewDomainInsightService(opts...)
	r.InsightSilences = NewDomainInsightSilenceService(opts...)
	r.Statistics = NewDomainStatisticService(opts...)
	r.CustomRules = NewDomainCustomRuleService(opts...)
	r.FirewallRules = NewDomainFirewallRuleService(opts...)
	r.AdvancedRules = NewDomainAdvancedRuleService(opts...)
	return
}

// Update Domain
func (r *DomainService) Update(ctx context.Context, domainID int64, body DomainUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Retrieve a list of domains associated with the client
func (r *DomainService) List(ctx context.Context, query DomainListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapSummaryDomain], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "waap/v1/domains"
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

// Retrieve a list of domains associated with the client
func (r *DomainService) ListAutoPaging(ctx context.Context, query DomainListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapSummaryDomain] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Delete an inactive domain by ID. Only domains with status 'bypass' can be
// deleted.
func (r *DomainService) Delete(ctx context.Context, domainID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve detailed information about a specific domain
func (r *DomainService) Get(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *WaapDetailedDomain, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve all rule sets linked to a particular domain
func (r *DomainService) ListRuleSets(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *[]WaapRuleSet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/domains/%v/rule-sets", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify the activation state of a policy associated with a domain
func (r *DomainService) TogglePolicy(ctx context.Context, policyID string, body DomainTogglePolicyParams, opts ...option.RequestOption) (res *WaapPolicyMode, err error) {
	opts = slices.Concat(r.Options, opts)
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("waap/v1/domains/%v/policies/%s/toggle", body.DomainID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Represents a WAAP domain, serving as a singular unit within the WAAP service.
//
// Each domain functions autonomously, possessing its own set of rules and
// configurations to manage web application firewall settings and behaviors.
type WaapDetailedDomain struct {
	// The domain ID
	ID int64 `json:"id,required"`
	// The date and time the domain was created in ISO 8601 format
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The ID of the custom page set
	CustomPageSet int64 `json:"custom_page_set,required"`
	// The domain name
	Name string `json:"name,required"`
	// The different statuses a domain can have
	//
	// Any of "active", "bypass", "monitor", "locked".
	Status WaapDetailedDomainStatus `json:"status,required"`
	// Domain level quotas
	Quotas map[string]WaapDetailedDomainQuota `json:"quotas,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		CustomPageSet respjson.Field
		Name          respjson.Field
		Status        respjson.Field
		Quotas        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDetailedDomain) RawJSON() string { return r.JSON.raw }
func (r *WaapDetailedDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The different statuses a domain can have
type WaapDetailedDomainStatus string

const (
	WaapDetailedDomainStatusActive  WaapDetailedDomainStatus = "active"
	WaapDetailedDomainStatusBypass  WaapDetailedDomainStatus = "bypass"
	WaapDetailedDomainStatusMonitor WaapDetailedDomainStatus = "monitor"
	WaapDetailedDomainStatusLocked  WaapDetailedDomainStatus = "locked"
)

type WaapDetailedDomainQuota struct {
	// The maximum allowed number of this resource
	Allowed int64 `json:"allowed,required"`
	// The current number of this resource
	Current int64 `json:"current,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allowed     respjson.Field
		Current     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDetailedDomainQuota) RawJSON() string { return r.JSON.raw }
func (r *WaapDetailedDomainQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API settings of a domain
type WaapDomainAPISettings struct {
	// The API URLs for a domain. If your domain has a common base URL for all API
	// paths, it can be set here
	APIURLs []string `json:"api_urls"`
	// Indicates if the domain is an API domain. All requests to an API domain are
	// treated as API requests. If this is set to true then the `api_urls` field is
	// ignored.
	IsAPI bool `json:"is_api"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIURLs     respjson.Field
		IsAPI       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDomainAPISettings) RawJSON() string { return r.JSON.raw }
func (r *WaapDomainAPISettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DDoS settings for a domain.
type WaapDomainDDOSSettings struct {
	// The burst threshold detects sudden rises in traffic. If it is met and the number
	// of requests is at least five times the last 2-second interval, DDoS protection
	// will activate. Default is 1000.
	BurstThreshold int64 `json:"burst_threshold"`
	// The global threshold is responsible for identifying DDoS attacks with a slow
	// rise in traffic. If the threshold is met and the current number of requests is
	// at least double that of the previous 10-second window, DDoS protection will
	// activate. Default is 5000.
	GlobalThreshold int64 `json:"global_threshold"`
	// The sub-second threshold protects WAAP servers against attacks from traffic
	// bursts. When this threshold is reached, the DDoS mode will activate on the
	// affected WAAP server, not the whole WAAP cluster. Default is 50.
	SubSecondThreshold int64 `json:"sub_second_threshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BurstThreshold     respjson.Field
		GlobalThreshold    respjson.Field
		SubSecondThreshold respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDomainDDOSSettings) RawJSON() string { return r.JSON.raw }
func (r *WaapDomainDDOSSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for a domain.
type WaapDomainSettingsModel struct {
	// API settings of a domain
	API WaapDomainAPISettings `json:"api,required"`
	// DDoS settings for a domain.
	DDOS WaapDomainDDOSSettings `json:"ddos,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		API         respjson.Field
		DDOS        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapDomainSettingsModel) RawJSON() string { return r.JSON.raw }
func (r *WaapDomainSettingsModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents the mode of a security rule.
type WaapPolicyMode struct {
	// Indicates if the security rule is active
	Mode bool `json:"mode,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapPolicyMode) RawJSON() string { return r.JSON.raw }
func (r *WaapPolicyMode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a custom rule set.
type WaapRuleSet struct {
	// Identifier of the rule set.
	ID int64 `json:"id,required"`
	// Detailed description of the rule set.
	Description string `json:"description,required"`
	// Indicates if the rule set is currently active.
	IsActive bool `json:"is_active,required"`
	// Name of the rule set.
	Name string `json:"name,required"`
	// Collection of tags associated with the rule set.
	Tags []WaapRuleSetTag `json:"tags,required"`
	// The resource slug associated with the rule set.
	ResourceSlug string            `json:"resource_slug,nullable"`
	Rules        []WaapRuleSetRule `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Description  respjson.Field
		IsActive     respjson.Field
		Name         respjson.Field
		Tags         respjson.Field
		ResourceSlug respjson.Field
		Rules        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRuleSet) RawJSON() string { return r.JSON.raw }
func (r *WaapRuleSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single tag associated with a rule set.
type WaapRuleSetTag struct {
	// Identifier of the tag.
	ID int64 `json:"id,required"`
	// Detailed description of the tag.
	Description string `json:"description,required"`
	// Name of the tag.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRuleSetTag) RawJSON() string { return r.JSON.raw }
func (r *WaapRuleSetTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a configurable WAAP security rule, also known as a policy.
type WaapRuleSetRule struct {
	// Unique identifier for the security rule
	ID string `json:"id,required"`
	// Specifies the action taken by the WAAP upon rule activation
	//
	// Any of "Allow", "Block", "Captcha", "Gateway", "Handshake", "Monitor",
	// "Composite".
	Action string `json:"action,required"`
	// Detailed description of the security rule
	Description string `json:"description,required"`
	// The rule set group name to which the rule belongs
	Group string `json:"group,required"`
	// Indicates if the security rule is active
	Mode bool `json:"mode,required"`
	// Name of the security rule
	Name string `json:"name,required"`
	// Identifier of the rule set to which the rule belongs
	RuleSetID int64 `json:"rule_set_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Action      respjson.Field
		Description respjson.Field
		Group       respjson.Field
		Mode        respjson.Field
		Name        respjson.Field
		RuleSetID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapRuleSetRule) RawJSON() string { return r.JSON.raw }
func (r *WaapRuleSetRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a WAAP domain when getting a list of domains.
type WaapSummaryDomain struct {
	// The domain ID
	ID int64 `json:"id,required"`
	// The date and time the domain was created in ISO 8601 format
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The ID of the custom page set
	CustomPageSet int64 `json:"custom_page_set,required"`
	// The domain name
	Name string `json:"name,required"`
	// The different statuses a domain can have
	//
	// Any of "active", "bypass", "monitor", "locked".
	Status WaapSummaryDomainStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		CustomPageSet respjson.Field
		Name          respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapSummaryDomain) RawJSON() string { return r.JSON.raw }
func (r *WaapSummaryDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The different statuses a domain can have
type WaapSummaryDomainStatus string

const (
	WaapSummaryDomainStatusActive  WaapSummaryDomainStatus = "active"
	WaapSummaryDomainStatusBypass  WaapSummaryDomainStatus = "bypass"
	WaapSummaryDomainStatusMonitor WaapSummaryDomainStatus = "monitor"
	WaapSummaryDomainStatusLocked  WaapSummaryDomainStatus = "locked"
)

type DomainUpdateParams struct {
	// The current status of the domain
	//
	// Any of "active", "monitor".
	Status DomainUpdateParamsStatus `json:"status,omitzero,required"`
	paramObj
}

func (r DomainUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the domain
type DomainUpdateParamsStatus string

const (
	DomainUpdateParamsStatusActive  DomainUpdateParamsStatus = "active"
	DomainUpdateParamsStatusMonitor DomainUpdateParamsStatus = "monitor"
)

type DomainListParams struct {
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter domains based on the domain name. Supports '\*' as a wildcard character
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter domains based on their IDs
	IDs []int64 `query:"ids,omitzero" json:"-"`
	// Sort the response by given field.
	//
	// Any of "id", "name", "status", "created_at", "-id", "-name", "-status",
	// "-created_at".
	Ordering DomainListParamsOrdering `query:"ordering,omitzero" json:"-"`
	// Filter domains based on the domain status
	//
	// Any of "active", "bypass", "monitor", "locked".
	Status DomainListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainListParams]'s query parameters as `url.Values`.
func (r DomainListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort the response by given field.
type DomainListParamsOrdering string

const (
	DomainListParamsOrderingID             DomainListParamsOrdering = "id"
	DomainListParamsOrderingName           DomainListParamsOrdering = "name"
	DomainListParamsOrderingStatus         DomainListParamsOrdering = "status"
	DomainListParamsOrderingCreatedAt      DomainListParamsOrdering = "created_at"
	DomainListParamsOrderingMinusID        DomainListParamsOrdering = "-id"
	DomainListParamsOrderingMinusName      DomainListParamsOrdering = "-name"
	DomainListParamsOrderingMinusStatus    DomainListParamsOrdering = "-status"
	DomainListParamsOrderingMinusCreatedAt DomainListParamsOrdering = "-created_at"
)

// Filter domains based on the domain status
type DomainListParamsStatus string

const (
	DomainListParamsStatusActive  DomainListParamsStatus = "active"
	DomainListParamsStatusBypass  DomainListParamsStatus = "bypass"
	DomainListParamsStatusMonitor DomainListParamsStatus = "monitor"
	DomainListParamsStatusLocked  DomainListParamsStatus = "locked"
)

type DomainTogglePolicyParams struct {
	// The domain ID
	DomainID int64 `path:"domain_id,required" json:"-"`
	paramObj
}
