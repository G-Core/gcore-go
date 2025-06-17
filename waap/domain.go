// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
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
	Policies        DomainPolicyService
	Analytics       DomainAnalyticsService
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
	r.Policies = NewDomainPolicyService(opts...)
	r.Analytics = NewDomainAnalyticsService(opts...)
	r.CustomRules = NewDomainCustomRuleService(opts...)
	r.FirewallRules = NewDomainFirewallRuleService(opts...)
	r.AdvancedRules = NewDomainAdvancedRuleService(opts...)
	return
}

// Update Domain
func (r *DomainService) Update(ctx context.Context, domainID int64, body DomainUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Retrieve a list of domains associated with the client
func (r *DomainService) List(ctx context.Context, query DomainListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapSummaryDomain], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve detailed information about a specific domain
func (r *DomainService) Get(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *WaapDetailedDomain, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve all rule sets linked to a particular domain
func (r *DomainService) ListRuleSets(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *[]WaapRuleSet, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/rule-sets", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DomainUpdateParams struct {
	// Domain statuses that can be used when updating a domain
	//
	// Any of "active", "monitor".
	Status DomainUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r DomainUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Domain statuses that can be used when updating a domain
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
	// The different statuses a domain can have
	//
	// Any of "active", "bypass", "monitor", "locked".
	Status WaapDomainStatus `query:"status,omitzero" json:"-"`
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
