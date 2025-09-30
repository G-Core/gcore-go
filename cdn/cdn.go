// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// CdnService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCdnService] method instead.
type CdnService struct {
	Options               []option.RequestOption
	Resources             ResourceService
	Shields               ShieldService
	OriginGroups          OriginGroupService
	RuleTemplates         RuleTemplateService
	Certificates          CertificateService
	TrustedCaCertificates TrustedCaCertificateService
	AuditLog              AuditLogService
	Logs                  LogService
	LogsUploader          LogsUploaderService
	Statistics            StatisticService
	NetworkCapacity       NetworkCapacityService
	Metrics               MetricService
	IPRanges              IPRangeService
}

// NewCdnService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCdnService(opts ...option.RequestOption) (r CdnService) {
	r = CdnService{}
	r.Options = opts
	r.Resources = NewResourceService(opts...)
	r.Shields = NewShieldService(opts...)
	r.OriginGroups = NewOriginGroupService(opts...)
	r.RuleTemplates = NewRuleTemplateService(opts...)
	r.Certificates = NewCertificateService(opts...)
	r.TrustedCaCertificates = NewTrustedCaCertificateService(opts...)
	r.AuditLog = NewAuditLogService(opts...)
	r.Logs = NewLogService(opts...)
	r.LogsUploader = NewLogsUploaderService(opts...)
	r.Statistics = NewStatisticService(opts...)
	r.NetworkCapacity = NewNetworkCapacityService(opts...)
	r.Metrics = NewMetricService(opts...)
	r.IPRanges = NewIPRangeService(opts...)
	return
}

// Get information about CDN service limits.
func (r *CdnService) GetAccountLimits(ctx context.Context, opts ...option.RequestOption) (res *CdnAccountLimits, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "cdn/clients/me/limits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get information about CDN service.
func (r *CdnService) GetAccountOverview(ctx context.Context, opts ...option.RequestOption) (res *CdnAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "cdn/clients/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get information about available CDN features.
func (r *CdnService) GetAvailableFeatures(ctx context.Context, opts ...option.RequestOption) (res *CdnAvailableFeatures, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "cdn/clients/me/features"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get purges history.
func (r *CdnService) ListPurgeStatuses(ctx context.Context, query CdnListPurgeStatusesParams, opts ...option.RequestOption) (res *pagination.OffsetPageCdn[PurgeStatus], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "cdn/purge_statuses"
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

// Get purges history.
func (r *CdnService) ListPurgeStatusesAutoPaging(ctx context.Context, query CdnListPurgeStatusesParams, opts ...option.RequestOption) *pagination.OffsetPageCdnAutoPager[PurgeStatus] {
	return pagination.NewOffsetPageCdnAutoPager(r.ListPurgeStatuses(ctx, query, opts...))
}

// Change information about CDN service.
func (r *CdnService) UpdateAccount(ctx context.Context, body CdnUpdateAccountParams, opts ...option.RequestOption) (res *CdnAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "cdn/clients/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type CdnAccount struct {
	// Account ID.
	ID int64 `json:"id"`
	// Defines whether resources will be deactivated automatically by inactivity.
	//
	// Possible values:
	//
	// - **true** - Resources will be deactivated.
	// - **false** - Resources will not be deactivated.
	AutoSuspendEnabled bool `json:"auto_suspend_enabled"`
	// Limit on the number of rules for each CDN resource.
	CdnResourcesRulesMaxCount int64 `json:"cdn_resources_rules_max_count"`
	// Domain zone to which a CNAME record of your CDN resources should be pointed.
	Cname string `json:"cname"`
	// Date of the first synchronization with the Platform (ISO 8601/RFC 3339 format,
	// UTC.)
	Created string `json:"created"`
	// Information about the CDN service status.
	Service CdnAccountService `json:"service"`
	// Date of the last update of information about CDN service (ISO 8601/RFC 3339
	// format, UTC.)
	Updated string `json:"updated"`
	// Defines whether custom balancing is used for content delivery.
	//
	// Possible values:
	//
	// - **true** - Custom balancing is used for content delivery.
	// - **false** - Custom balancing is not used for content delivery.
	UseBalancer bool `json:"use_balancer"`
	// CDN traffic usage limit in gigabytes.
	//
	// When the limit is reached, we will send an email notification.
	UtilizationLevel int64 `json:"utilization_level"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		AutoSuspendEnabled        respjson.Field
		CdnResourcesRulesMaxCount respjson.Field
		Cname                     respjson.Field
		Created                   respjson.Field
		Service                   respjson.Field
		Updated                   respjson.Field
		UseBalancer               respjson.Field
		UtilizationLevel          respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnAccount) RawJSON() string { return r.JSON.raw }
func (r *CdnAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the CDN service status.
type CdnAccountService struct {
	// Defines whether the CDN service is activated.
	//
	// Possible values:
	//
	// - **true** - Service is activated.
	// - **false** - Service is not activated.
	Enabled bool `json:"enabled"`
	// CDN service status.
	//
	// Possible values:
	//
	//   - **new** - CDN service is not activated.
	//   - **trial** - Free trial is in progress.
	//   - **trialend** - Free trial has ended and CDN service is stopped. All CDN
	//     resources are suspended.
	//   - **activating** - CDN service is being activated. It can take up to 15 minutes.
	//   - **active** - CDN service is active.
	//   - **paused** - CDN service is stopped. All CDN resources are suspended.
	//   - **deleted** - CDN service is stopped. All CDN resources are deleted.
	Status string `json:"status"`
	// Date of the last CDN service status update (ISO 8601/RFC 3339 format, UTC.)
	Updated string `json:"updated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Status      respjson.Field
		Updated     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnAccountService) RawJSON() string { return r.JSON.raw }
func (r *CdnAccountService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdnAccountLimits struct {
	// Account ID.
	ID int64 `json:"id"`
	// Maximum number of origins that can be added to the origin group on your tariff
	// plan.
	OriginsInGroupLimit int64 `json:"origins_in_group_limit"`
	// Maximum number of CDN resources that can be created on your tariff plan.
	ResourcesLimit int64 `json:"resources_limit"`
	// Maximum number of rules that can be created per CDN resource on your tariff
	// plan.
	RulesLimit int64 `json:"rules_limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		OriginsInGroupLimit respjson.Field
		ResourcesLimit      respjson.Field
		RulesLimit          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnAccountLimits) RawJSON() string { return r.JSON.raw }
func (r *CdnAccountLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdnAvailableFeatures struct {
	// Account ID.
	ID int64 `json:"id"`
	// Free features available for your account.
	FreeFeatures []CdnAvailableFeaturesFreeFeature `json:"free_features"`
	// Paid features available for your account.
	PaidFeatures []CdnAvailableFeaturesPaidFeature `json:"paid_features"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		FreeFeatures respjson.Field
		PaidFeatures respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnAvailableFeatures) RawJSON() string { return r.JSON.raw }
func (r *CdnAvailableFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdnAvailableFeaturesFreeFeature struct {
	// Date and time when the feature was activated (ISO 8601/RFC 3339 format, UTC.)
	CreateDate string `json:"create_date"`
	// Feature ID.
	FeatureID int64 `json:"feature_id"`
	// Internal feature activation ID.
	FreeFeatureID int64 `json:"free_feature_id"`
	// Feature name.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate    respjson.Field
		FeatureID     respjson.Field
		FreeFeatureID respjson.Field
		Name          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnAvailableFeaturesFreeFeature) RawJSON() string { return r.JSON.raw }
func (r *CdnAvailableFeaturesFreeFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdnAvailableFeaturesPaidFeature struct {
	// Date and time when the feature was activated (ISO 8601/RFC 3339 format, UTC.)
	CreateDate string `json:"create_date"`
	// Feature ID.
	FeatureID int64 `json:"feature_id"`
	// Feature name.
	Name string `json:"name"`
	// Internal feature activation ID.
	PaidFeatureID int64 `json:"paid_feature_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate    respjson.Field
		FeatureID     respjson.Field
		Name          respjson.Field
		PaidFeatureID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnAvailableFeaturesPaidFeature) RawJSON() string { return r.JSON.raw }
func (r *CdnAvailableFeaturesPaidFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PurgeStatus struct {
	// Date and time when the purge was created (ISO 8601/RFC 3339 format, UTC).
	Created string `json:"created"`
	// Purge payload depends on purge type.
	//
	// Possible values:
	//
	// - **urls** - Purge by URL.
	// - **paths** - Purge by Pattern and purge All.
	Payload any `json:"payload"`
	// Purge ID.
	PurgeID int64 `json:"purge_id"`
	// Contains the name of the purge request type.
	//
	// Possible values:
	//
	// - **`purge_by_pattern`** - Purge by Pattern.
	// - **`purge_by_url`** - Purge by URL.
	// - **`purge_all`** - Purge All.
	PurgeType string              `json:"purge_type"`
	Resource  PurgeStatusResource `json:"resource"`
	// Purge status.
	//
	// Possible values:
	//
	// - **In progress** - Purge is in progress.
	// - **Successful** - Purge was successful.
	// - **Failed** - Purge failed.
	//
	// Any of "In progress", "Successful", "Failed".
	Status PurgeStatusStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Created     respjson.Field
		Payload     respjson.Field
		PurgeID     respjson.Field
		PurgeType   respjson.Field
		Resource    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PurgeStatus) RawJSON() string { return r.JSON.raw }
func (r *PurgeStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PurgeStatusResource struct {
	// Resource ID.
	ID int64 `json:"id"`
	// CNAME of the resource.
	Cname string `json:"cname"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cname       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PurgeStatusResource) RawJSON() string { return r.JSON.raw }
func (r *PurgeStatusResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Purge status.
//
// Possible values:
//
// - **In progress** - Purge is in progress.
// - **Successful** - Purge was successful.
// - **Failed** - Purge failed.
type PurgeStatusStatus string

const (
	PurgeStatusStatusInProgress PurgeStatusStatus = "In progress"
	PurgeStatusStatusSuccessful PurgeStatusStatus = "Successful"
	PurgeStatusStatusFailed     PurgeStatusStatus = "Failed"
)

type CdnListPurgeStatusesParams struct {
	// Purges associated with a specific resource CNAME.
	//
	// Example:
	//
	// - &cname=example.com
	Cname param.Opt[string] `query:"cname,omitzero" json:"-"`
	// Start date and time of the requested time period (ISO 8601/RFC 3339 format,
	// UTC.)
	//
	// Examples:
	//
	// - &`from_created`=2021-06-14T00:00:00Z
	// - &`from_created`=2021-06-14T00:00:00.000Z
	FromCreated param.Opt[string] `query:"from_created,omitzero" json:"-"`
	// Maximum number of purges in the response.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of purge requests in the response to skip starting from the beginning of
	// the requested period.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Purge requests with a certain purge type.
	//
	// Possible values:
	//
	// - **`purge_by_pattern`** - Purge by Pattern.
	// - **`purge_by_url`** - Purge by URL.
	// - **`purge_all`** - Purge All.
	PurgeType param.Opt[string] `query:"purge_type,omitzero" json:"-"`
	// Purge with a certain status.
	//
	// Possible values:
	//
	// - **In progress**
	// - **Successful**
	// - **Failed**
	// - **Status report disabled**
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// End date and time of the requested time period (ISO 8601/RFC 3339 format, UTC.)
	//
	// Examples:
	//
	// - &`to_created`=2021-06-15T00:00:00Z
	// - &`to_created`=2021-06-15T00:00:00.000Z
	ToCreated param.Opt[string] `query:"to_created,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CdnListPurgeStatusesParams]'s query parameters as
// `url.Values`.
func (r CdnListPurgeStatusesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CdnUpdateAccountParams struct {
	// CDN traffic usage limit in gigabytes.
	//
	// When the limit is reached, we will send an email notification.
	UtilizationLevel param.Opt[int64] `json:"utilization_level,omitzero"`
	paramObj
}

func (r CdnUpdateAccountParams) MarshalJSON() (data []byte, err error) {
	type shadow CdnUpdateAccountParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CdnUpdateAccountParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
