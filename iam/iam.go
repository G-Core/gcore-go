// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// IamService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIamService] method instead.
type IamService struct {
	Options   []option.RequestOption
	APITokens APITokenService
	Users     UserService
}

// NewIamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIamService(opts ...option.RequestOption) (r IamService) {
	r = IamService{}
	r.Options = opts
	r.APITokens = NewAPITokenService(opts...)
	r.Users = NewUserService(opts...)
	return
}

// Get information about your profile, users and other account details.
func (r *IamService) GetAccountOverview(ctx context.Context, opts ...option.RequestOption) (res *AccountOverview, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "iam/clients/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountOverview struct {
	// The account ID.
	ID int64 `json:"id,required"`
	// System field. Billing type of the account.
	BillType string `json:"bill_type,required"`
	// System field. List of services available for the account.
	//
	// Any of "CDN", "STORAGE", "STREAMING", "DNS", "DDOS", "CLOUD".
	Capabilities []string `json:"capabilities,required"`
	// The company name.
	CompanyName string `json:"companyName,required"`
	// ID of the current user.
	CurrentUser int64 `json:"currentUser,required"`
	// The field shows the status of the account:
	//
	// - `true` – the account has been deleted
	// - `false` – the account is not deleted
	Deleted bool `json:"deleted,required"`
	// The account email.
	Email string `json:"email,required" format:"email"`
	// System field. Control panel domain.
	EntryBaseDomain string `json:"entryBaseDomain,required"`
	// An object of arrays which contains information about free features available for
	// the requested account.
	FreeFeatures AccountOverviewFreeFeatures `json:"freeFeatures,required"`
	// System field.
	HasActiveAdmin bool `json:"has_active_admin,required"`
	// System field:
	//
	// - `true` — a test account;
	// - `false` — a production account.
	IsTest bool `json:"is_test,required"`
	// Name of a user who registered the requested account.
	Name string `json:"name,required"`
	// An object of arrays which contains information about paid features available for
	// the requested account.
	PaidFeatures AccountOverviewPaidFeatures `json:"paidFeatures,required"`
	// An object of arrays which contains information about all services available for
	// the requested account.
	ServiceStatuses AccountOverviewServiceStatuses `json:"serviceStatuses,required"`
	// Status of the account.
	//
	// Any of "new", "trial", "trialend", "active", "integration", "paused",
	// "preparation", "ready".
	Status AccountOverviewStatus `json:"status,required"`
	// System field. The company country (ISO 3166-1 alpha-2 format).
	CountryCode string `json:"country_code"`
	// The account custom ID.
	CustomID string `json:"custom_id,nullable"`
	// Phone of a user who registered the requested account.
	Phone string `json:"phone,nullable"`
	// System field. Type of the account registration process.
	//
	// Any of "sign_up_full", "sign_up_simple".
	SignupProcess AccountOverviewSignupProcess `json:"signup_process,nullable"`
	// List of account users.
	Users []AccountOverviewUser `json:"users"`
	// The company website.
	Website string `json:"website"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		BillType        respjson.Field
		Capabilities    respjson.Field
		CompanyName     respjson.Field
		CurrentUser     respjson.Field
		Deleted         respjson.Field
		Email           respjson.Field
		EntryBaseDomain respjson.Field
		FreeFeatures    respjson.Field
		HasActiveAdmin  respjson.Field
		IsTest          respjson.Field
		Name            respjson.Field
		PaidFeatures    respjson.Field
		ServiceStatuses respjson.Field
		Status          respjson.Field
		CountryCode     respjson.Field
		CustomID        respjson.Field
		Phone           respjson.Field
		SignupProcess   respjson.Field
		Users           respjson.Field
		Website         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverview) RawJSON() string { return r.JSON.raw }
func (r *AccountOverview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object of arrays which contains information about free features available for
// the requested account.
type AccountOverviewFreeFeatures struct {
	CDN       []AccountOverviewFreeFeaturesCDN       `json:"CDN"`
	Cloud     []AccountOverviewFreeFeaturesCloud     `json:"CLOUD"`
	DDOS      []AccountOverviewFreeFeaturesDDOS      `json:"DDOS"`
	DNS       []AccountOverviewFreeFeaturesDNS       `json:"DNS"`
	Storage   []AccountOverviewFreeFeaturesStorage   `json:"STORAGE"`
	Streaming []AccountOverviewFreeFeaturesStreaming `json:"STREAMING"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CDN         respjson.Field
		Cloud       respjson.Field
		DDOS        respjson.Field
		DNS         respjson.Field
		Storage     respjson.Field
		Streaming   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewFreeFeatures) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewFreeFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature object.
type AccountOverviewFreeFeaturesCDN struct {
	// Date and time when the feature was activated (ISO 8086/RFC 3339 format).
	CreateDate string `json:"create_date"`
	// Feature ID.
	FeatureID int64 `json:"feature_id"`
	// Internal feature activation ID.
	FreeFeatureID int64 `json:"free_feature_id"`
	// Name of the feature.
	Name string `json:"name"`
	// Service's name.
	//
	// Any of "CDN", "STORAGE", "STREAMING", "DNS", "DDOS", "CLOUD".
	Service string `json:"service"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate    respjson.Field
		FeatureID     respjson.Field
		FreeFeatureID respjson.Field
		Name          respjson.Field
		Service       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewFreeFeaturesCDN) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewFreeFeaturesCDN) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature object.
type AccountOverviewFreeFeaturesCloud struct {
	// Date and time when the feature was activated (ISO 8086/RFC 3339 format).
	CreateDate string `json:"create_date"`
	// Feature ID.
	FeatureID int64 `json:"feature_id"`
	// Internal feature activation ID.
	FreeFeatureID int64 `json:"free_feature_id"`
	// Name of the feature.
	Name string `json:"name"`
	// Service's name.
	//
	// Any of "CDN", "STORAGE", "STREAMING", "DNS", "DDOS", "CLOUD".
	Service string `json:"service"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate    respjson.Field
		FeatureID     respjson.Field
		FreeFeatureID respjson.Field
		Name          respjson.Field
		Service       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewFreeFeaturesCloud) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewFreeFeaturesCloud) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature object.
type AccountOverviewFreeFeaturesDDOS struct {
	// Date and time when the feature was activated (ISO 8086/RFC 3339 format).
	CreateDate string `json:"create_date"`
	// Feature ID.
	FeatureID int64 `json:"feature_id"`
	// Internal feature activation ID.
	FreeFeatureID int64 `json:"free_feature_id"`
	// Name of the feature.
	Name string `json:"name"`
	// Service's name.
	//
	// Any of "CDN", "STORAGE", "STREAMING", "DNS", "DDOS", "CLOUD".
	Service string `json:"service"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate    respjson.Field
		FeatureID     respjson.Field
		FreeFeatureID respjson.Field
		Name          respjson.Field
		Service       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewFreeFeaturesDDOS) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewFreeFeaturesDDOS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature object.
type AccountOverviewFreeFeaturesDNS struct {
	// Date and time when the feature was activated (ISO 8086/RFC 3339 format).
	CreateDate string `json:"create_date"`
	// Feature ID.
	FeatureID int64 `json:"feature_id"`
	// Internal feature activation ID.
	FreeFeatureID int64 `json:"free_feature_id"`
	// Name of the feature.
	Name string `json:"name"`
	// Service's name.
	//
	// Any of "CDN", "STORAGE", "STREAMING", "DNS", "DDOS", "CLOUD".
	Service string `json:"service"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate    respjson.Field
		FeatureID     respjson.Field
		FreeFeatureID respjson.Field
		Name          respjson.Field
		Service       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewFreeFeaturesDNS) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewFreeFeaturesDNS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature object.
type AccountOverviewFreeFeaturesStorage struct {
	// Date and time when the feature was activated (ISO 8086/RFC 3339 format).
	CreateDate string `json:"create_date"`
	// Feature ID.
	FeatureID int64 `json:"feature_id"`
	// Internal feature activation ID.
	FreeFeatureID int64 `json:"free_feature_id"`
	// Name of the feature.
	Name string `json:"name"`
	// Service's name.
	//
	// Any of "CDN", "STORAGE", "STREAMING", "DNS", "DDOS", "CLOUD".
	Service string `json:"service"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate    respjson.Field
		FeatureID     respjson.Field
		FreeFeatureID respjson.Field
		Name          respjson.Field
		Service       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewFreeFeaturesStorage) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewFreeFeaturesStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature object.
type AccountOverviewFreeFeaturesStreaming struct {
	// Date and time when the feature was activated (ISO 8086/RFC 3339 format).
	CreateDate string `json:"create_date"`
	// Feature ID.
	FeatureID int64 `json:"feature_id"`
	// Internal feature activation ID.
	FreeFeatureID int64 `json:"free_feature_id"`
	// Name of the feature.
	Name string `json:"name"`
	// Service's name.
	//
	// Any of "CDN", "STORAGE", "STREAMING", "DNS", "DDOS", "CLOUD".
	Service string `json:"service"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate    respjson.Field
		FeatureID     respjson.Field
		FreeFeatureID respjson.Field
		Name          respjson.Field
		Service       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewFreeFeaturesStreaming) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewFreeFeaturesStreaming) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object of arrays which contains information about paid features available for
// the requested account.
type AccountOverviewPaidFeatures struct {
	CDN       []AccountOverviewPaidFeaturesCDN       `json:"CDN"`
	Cloud     []AccountOverviewPaidFeaturesCloud     `json:"CLOUD"`
	DDOS      []AccountOverviewPaidFeaturesDDOS      `json:"DDOS"`
	DNS       []AccountOverviewPaidFeaturesDNS       `json:"DNS"`
	Storage   []AccountOverviewPaidFeaturesStorage   `json:"STORAGE"`
	Streaming []AccountOverviewPaidFeaturesStreaming `json:"STREAMING"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CDN         respjson.Field
		Cloud       respjson.Field
		DDOS        respjson.Field
		DNS         respjson.Field
		Storage     respjson.Field
		Streaming   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewPaidFeatures) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewPaidFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature object.
type AccountOverviewPaidFeaturesCDN struct {
	// Date and time when the feature was activated (ISO 8086/RFC 3339 format).
	CreateDate string `json:"create_date"`
	// Feature ID.
	FeatureID int64 `json:"feature_id"`
	// Name of the feature.
	Name string `json:"name"`
	// Internal feature activation ID.
	PaidFeatureID int64 `json:"paid_feature_id"`
	// Service's name.
	//
	// Any of "CDN", "STORAGE", "STREAMING", "DNS", "DDOS", "CLOUD".
	Service string `json:"service"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate    respjson.Field
		FeatureID     respjson.Field
		Name          respjson.Field
		PaidFeatureID respjson.Field
		Service       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewPaidFeaturesCDN) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewPaidFeaturesCDN) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature object.
type AccountOverviewPaidFeaturesCloud struct {
	// Date and time when the feature was activated (ISO 8086/RFC 3339 format).
	CreateDate string `json:"create_date"`
	// Feature ID.
	FeatureID int64 `json:"feature_id"`
	// Name of the feature.
	Name string `json:"name"`
	// Internal feature activation ID.
	PaidFeatureID int64 `json:"paid_feature_id"`
	// Service's name.
	//
	// Any of "CDN", "STORAGE", "STREAMING", "DNS", "DDOS", "CLOUD".
	Service string `json:"service"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate    respjson.Field
		FeatureID     respjson.Field
		Name          respjson.Field
		PaidFeatureID respjson.Field
		Service       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewPaidFeaturesCloud) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewPaidFeaturesCloud) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature object.
type AccountOverviewPaidFeaturesDDOS struct {
	// Date and time when the feature was activated (ISO 8086/RFC 3339 format).
	CreateDate string `json:"create_date"`
	// Feature ID.
	FeatureID int64 `json:"feature_id"`
	// Name of the feature.
	Name string `json:"name"`
	// Internal feature activation ID.
	PaidFeatureID int64 `json:"paid_feature_id"`
	// Service's name.
	//
	// Any of "CDN", "STORAGE", "STREAMING", "DNS", "DDOS", "CLOUD".
	Service string `json:"service"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate    respjson.Field
		FeatureID     respjson.Field
		Name          respjson.Field
		PaidFeatureID respjson.Field
		Service       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewPaidFeaturesDDOS) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewPaidFeaturesDDOS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature object.
type AccountOverviewPaidFeaturesDNS struct {
	// Date and time when the feature was activated (ISO 8086/RFC 3339 format).
	CreateDate string `json:"create_date"`
	// Feature ID.
	FeatureID int64 `json:"feature_id"`
	// Name of the feature.
	Name string `json:"name"`
	// Internal feature activation ID.
	PaidFeatureID int64 `json:"paid_feature_id"`
	// Service's name.
	//
	// Any of "CDN", "STORAGE", "STREAMING", "DNS", "DDOS", "CLOUD".
	Service string `json:"service"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate    respjson.Field
		FeatureID     respjson.Field
		Name          respjson.Field
		PaidFeatureID respjson.Field
		Service       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewPaidFeaturesDNS) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewPaidFeaturesDNS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature object.
type AccountOverviewPaidFeaturesStorage struct {
	// Date and time when the feature was activated (ISO 8086/RFC 3339 format).
	CreateDate string `json:"create_date"`
	// Feature ID.
	FeatureID int64 `json:"feature_id"`
	// Name of the feature.
	Name string `json:"name"`
	// Internal feature activation ID.
	PaidFeatureID int64 `json:"paid_feature_id"`
	// Service's name.
	//
	// Any of "CDN", "STORAGE", "STREAMING", "DNS", "DDOS", "CLOUD".
	Service string `json:"service"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate    respjson.Field
		FeatureID     respjson.Field
		Name          respjson.Field
		PaidFeatureID respjson.Field
		Service       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewPaidFeaturesStorage) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewPaidFeaturesStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feature object.
type AccountOverviewPaidFeaturesStreaming struct {
	// Date and time when the feature was activated (ISO 8086/RFC 3339 format).
	CreateDate string `json:"create_date"`
	// Feature ID.
	FeatureID int64 `json:"feature_id"`
	// Name of the feature.
	Name string `json:"name"`
	// Internal feature activation ID.
	PaidFeatureID int64 `json:"paid_feature_id"`
	// Service's name.
	//
	// Any of "CDN", "STORAGE", "STREAMING", "DNS", "DDOS", "CLOUD".
	Service string `json:"service"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate    respjson.Field
		FeatureID     respjson.Field
		Name          respjson.Field
		PaidFeatureID respjson.Field
		Service       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewPaidFeaturesStreaming) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewPaidFeaturesStreaming) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object of arrays which contains information about all services available for
// the requested account.
type AccountOverviewServiceStatuses struct {
	CDN       AccountOverviewServiceStatusesCDN       `json:"CDN"`
	Cloud     AccountOverviewServiceStatusesCloud     `json:"CLOUD"`
	DDOS      AccountOverviewServiceStatusesDDOS      `json:"DDOS"`
	DNS       AccountOverviewServiceStatusesDNS       `json:"DNS"`
	Storage   AccountOverviewServiceStatusesStorage   `json:"STORAGE"`
	Streaming AccountOverviewServiceStatusesStreaming `json:"STREAMING"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CDN         respjson.Field
		Cloud       respjson.Field
		DDOS        respjson.Field
		DNS         respjson.Field
		Storage     respjson.Field
		Streaming   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewServiceStatuses) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewServiceStatuses) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOverviewServiceStatusesCDN struct {
	// `true` - service is available in the Control Panel.
	Enabled bool `json:"enabled"`
	// Status of the service.
	//
	// Any of "new", "trial", "trialend", "active", "paused", "activating", "deleted".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewServiceStatusesCDN) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewServiceStatusesCDN) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOverviewServiceStatusesCloud struct {
	// `true` - service is available in the Control Panel.
	Enabled bool `json:"enabled"`
	// Status of the service.
	//
	// Any of "new", "trial", "trialend", "active", "paused", "activating", "deleted".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewServiceStatusesCloud) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewServiceStatusesCloud) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOverviewServiceStatusesDDOS struct {
	// `true` - service is available in the Control Panel.
	Enabled bool `json:"enabled"`
	// Status of the service.
	//
	// Any of "new", "trial", "trialend", "active", "paused", "activating", "deleted".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewServiceStatusesDDOS) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewServiceStatusesDDOS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOverviewServiceStatusesDNS struct {
	// `true` - service is available in the Control Panel.
	Enabled bool `json:"enabled"`
	// Status of the service.
	//
	// Any of "new", "trial", "trialend", "active", "paused", "activating", "deleted".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewServiceStatusesDNS) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewServiceStatusesDNS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOverviewServiceStatusesStorage struct {
	// `true` - service is available in the Control Panel.
	Enabled bool `json:"enabled"`
	// Status of the service.
	//
	// Any of "new", "trial", "trialend", "active", "paused", "activating", "deleted".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewServiceStatusesStorage) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewServiceStatusesStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOverviewServiceStatusesStreaming struct {
	// `true` - service is available in the Control Panel.
	Enabled bool `json:"enabled"`
	// Status of the service.
	//
	// Any of "new", "trial", "trialend", "active", "paused", "activating", "deleted".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewServiceStatusesStreaming) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewServiceStatusesStreaming) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the account.
type AccountOverviewStatus string

const (
	AccountOverviewStatusNew         AccountOverviewStatus = "new"
	AccountOverviewStatusTrial       AccountOverviewStatus = "trial"
	AccountOverviewStatusTrialend    AccountOverviewStatus = "trialend"
	AccountOverviewStatusActive      AccountOverviewStatus = "active"
	AccountOverviewStatusIntegration AccountOverviewStatus = "integration"
	AccountOverviewStatusPaused      AccountOverviewStatus = "paused"
	AccountOverviewStatusPreparation AccountOverviewStatus = "preparation"
	AccountOverviewStatusReady       AccountOverviewStatus = "ready"
)

// System field. Type of the account registration process.
type AccountOverviewSignupProcess string

const (
	AccountOverviewSignupProcessSignUpFull   AccountOverviewSignupProcess = "sign_up_full"
	AccountOverviewSignupProcessSignUpSimple AccountOverviewSignupProcess = "sign_up_simple"
)

type AccountOverviewUser struct {
	// User's ID.
	ID int64 `json:"id"`
	// Email confirmation:
	//
	// - `true` – user confirmed the email;
	// - `false` – user did not confirm the email.
	Activated bool `json:"activated"`
	// System field. List of auth types available for the account.
	AuthTypes []AuthType `json:"auth_types"`
	// User's account ID.
	Client float64 `json:"client"`
	// User's company.
	Company string `json:"company"`
	// Deletion flag. If `true` then user was deleted.
	Deleted bool `json:"deleted"`
	// User's email address.
	Email string `json:"email" format:"email"`
	// User's group in the current account.
	//
	// IAM supports 5 groups:
	//
	// - Users
	// - Administrators
	// - Engineers
	// - Purge and Prefetch only (API)
	// - Purge and Prefetch only (API+Web)
	Groups []UserGroup `json:"groups"`
	// User's language.
	//
	// Defines language of the control panel and email messages.
	//
	// Any of "de", "en", "ru", "zh", "az".
	Lang UserLanguage `json:"lang"`
	// User's name.
	Name string `json:"name,nullable"`
	// User's phone.
	Phone string `json:"phone,nullable"`
	// Services provider ID.
	Reseller int64 `json:"reseller"`
	// SSO authentication flag. If `true` then user can login via SAML SSO.
	SSOAuth bool `json:"sso_auth"`
	// Two-step verification:
	//
	// - `true` – user enabled two-step verification;
	// - `false` – user disabled two-step verification.
	TwoFa bool `json:"two_fa"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Activated   respjson.Field
		AuthTypes   respjson.Field
		Client      respjson.Field
		Company     respjson.Field
		Deleted     respjson.Field
		Email       respjson.Field
		Groups      respjson.Field
		Lang        respjson.Field
		Name        respjson.Field
		Phone       respjson.Field
		Reseller    respjson.Field
		SSOAuth     respjson.Field
		TwoFa       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountOverviewUser) RawJSON() string { return r.JSON.raw }
func (r *AccountOverviewUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Auth types.
type AuthType string

const (
	AuthTypePassword     AuthType = "password"
	AuthTypeSSO          AuthType = "sso"
	AuthTypeGitHub       AuthType = "github"
	AuthTypeGoogleOauth2 AuthType = "google-oauth2"
)

type UserGroup struct {
	// Group's ID: Possible values are:
	//
	//   - 1 - Administrators* 2 - Users* 5 - Engineers* 3009 - Purge and Prefetch only
	//     (API+Web)* 3022 - Purge and Prefetch only (API)
	ID int64 `json:"id"`
	// Group's name.
	//
	// Any of "Users", "Administrators", "Engineers", "Purge and Prefetch only (API)",
	// "Purge and Prefetch only (API+Web)".
	Name UserGroupName `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGroup) RawJSON() string { return r.JSON.raw }
func (r *UserGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UserGroup to a UserGroupParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UserGroupParam.Overrides()
func (r UserGroup) ToParam() UserGroupParam {
	return param.Override[UserGroupParam](json.RawMessage(r.RawJSON()))
}

// Group's name.
type UserGroupName string

const (
	UserGroupNameUsers                      UserGroupName = "Users"
	UserGroupNameAdministrators             UserGroupName = "Administrators"
	UserGroupNameEngineers                  UserGroupName = "Engineers"
	UserGroupNamePurgeAndPrefetchOnlyAPI    UserGroupName = "Purge and Prefetch only (API)"
	UserGroupNamePurgeAndPrefetchOnlyAPIWeb UserGroupName = "Purge and Prefetch only (API+Web)"
)

type UserGroupParam struct {
	// Group's ID: Possible values are:
	//
	//   - 1 - Administrators* 2 - Users* 5 - Engineers* 3009 - Purge and Prefetch only
	//     (API+Web)* 3022 - Purge and Prefetch only (API)
	ID param.Opt[int64] `json:"id,omitzero"`
	// Group's name.
	//
	// Any of "Users", "Administrators", "Engineers", "Purge and Prefetch only (API)",
	// "Purge and Prefetch only (API+Web)".
	Name UserGroupName `json:"name,omitzero"`
	paramObj
}

func (r UserGroupParam) MarshalJSON() (data []byte, err error) {
	type shadow UserGroupParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserGroupParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User's language.
//
// Defines language of the control panel and email messages.
type UserLanguage string

const (
	UserLanguageDe UserLanguage = "de"
	UserLanguageEn UserLanguage = "en"
	UserLanguageRu UserLanguage = "ru"
	UserLanguageZh UserLanguage = "zh"
	UserLanguageAz UserLanguage = "az"
)
