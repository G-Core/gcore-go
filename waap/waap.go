// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// WaapService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapService] method instead.
type WaapService struct {
	Options        []option.RequestOption
	Statistics     StatisticService
	Domains        DomainService
	CustomPageSets CustomPageSetService
	AdvancedRules  AdvancedRuleService
	Tags           TagService
	Organizations  OrganizationService
	Insights       InsightService
	IPInfo         IPInfoService
}

// NewWaapService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWaapService(opts ...option.RequestOption) (r WaapService) {
	r = WaapService{}
	r.Options = opts
	r.Statistics = NewStatisticService(opts...)
	r.Domains = NewDomainService(opts...)
	r.CustomPageSets = NewCustomPageSetService(opts...)
	r.AdvancedRules = NewAdvancedRuleService(opts...)
	r.Tags = NewTagService(opts...)
	r.Organizations = NewOrganizationService(opts...)
	r.Insights = NewInsightService(opts...)
	r.IPInfo = NewIPInfoService(opts...)
	return
}

// Get information about WAAP service for the client
func (r *WaapService) GetAccountOverview(ctx context.Context, opts ...option.RequestOption) (res *WaapGetAccountOverviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/clients/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Represents the WAAP service information for a client
type WaapGetAccountOverviewResponse struct {
	// The client ID
	ID int64 `json:"id" api:"required"`
	// List of enabled features
	Features []string `json:"features" api:"required"`
	// Quotas for the client
	Quotas map[string]WaapGetAccountOverviewResponseQuota `json:"quotas" api:"required"`
	// Information about the WAAP service status
	Service WaapGetAccountOverviewResponseService `json:"service" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Features    respjson.Field
		Quotas      respjson.Field
		Service     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapGetAccountOverviewResponse) RawJSON() string { return r.JSON.raw }
func (r *WaapGetAccountOverviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WaapGetAccountOverviewResponseQuota struct {
	// The maximum allowed number of this resource
	Allowed int64 `json:"allowed" api:"required"`
	// The current number of this resource
	Current int64 `json:"current" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allowed     respjson.Field
		Current     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapGetAccountOverviewResponseQuota) RawJSON() string { return r.JSON.raw }
func (r *WaapGetAccountOverviewResponseQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the WAAP service status
type WaapGetAccountOverviewResponseService struct {
	// Whether the service is enabled
	Enabled bool `json:"enabled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapGetAccountOverviewResponseService) RawJSON() string { return r.JSON.raw }
func (r *WaapGetAccountOverviewResponseService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
