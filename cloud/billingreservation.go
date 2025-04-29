// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// BillingReservationService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillingReservationService] method instead.
type BillingReservationService struct {
	Options []option.RequestOption
}

// NewBillingReservationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBillingReservationService(opts ...option.RequestOption) (r BillingReservationService) {
	r = BillingReservationService{}
	r.Options = opts
	return
}

// List reservations
func (r *BillingReservationService) List(ctx context.Context, query BillingReservationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[BillingReservation], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cloud/v1/reservations"
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

// List reservations
func (r *BillingReservationService) ListAutoPaging(ctx context.Context, query BillingReservationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[BillingReservation] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Get specific reservation
func (r *BillingReservationService) Get(ctx context.Context, reservationID int64, opts ...option.RequestOption) (res *BillingReservation, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/reservations/%v", reservationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// '#/components/schemas/BillingReservationItemResponseSerializer'
// "$.components.schemas.BillingReservationItemResponseSerializer"
type BillingReservation struct {
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/id'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.id"
	ID int64 `json:"id,required"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/active_from'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.active_from"
	ActiveFrom time.Time `json:"active_from,required" format:"date"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/active_to'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.active_to"
	ActiveTo time.Time `json:"active_to,required" format:"date"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/activity_period'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.activity_period"
	ActivityPeriod string `json:"activity_period,required"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/activity_period_length'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.activity_period_length"
	ActivityPeriodLength int64 `json:"activity_period_length,required"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/amount_prices'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.amount_prices"
	AmountPrices BillingReservationAmountPrices `json:"amount_prices,required"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/billing_plan_id'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.billing_plan_id"
	BillingPlanID int64 `json:"billing_plan_id,required"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/created_at'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/error/anyOf/0'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.error.anyOf[0]"
	Error string `json:"error,required"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/eta/anyOf/0'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.eta.anyOf[0]"
	Eta time.Time `json:"eta,required" format:"date"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/is_expiration_message_visible'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.is_expiration_message_visible"
	IsExpirationMessageVisible bool `json:"is_expiration_message_visible,required"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/name'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/next_statuses'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.next_statuses"
	NextStatuses []string `json:"next_statuses,required"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/region_id'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/region_name'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.region_name"
	RegionName string `json:"region_name,required"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/remind_expiration_message/anyOf/0'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.remind_expiration_message.anyOf[0]"
	RemindExpirationMessage time.Time `json:"remind_expiration_message,required" format:"date"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/resources'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.resources"
	Resources []BillingReservationResource `json:"resources,required"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/status'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.status"
	Status string `json:"status,required"`
	// '#/components/schemas/BillingReservationItemResponseSerializer/properties/user_status'
	// "$.components.schemas.BillingReservationItemResponseSerializer.properties.user_status"
	UserStatus string `json:"user_status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                         resp.Field
		ActiveFrom                 resp.Field
		ActiveTo                   resp.Field
		ActivityPeriod             resp.Field
		ActivityPeriodLength       resp.Field
		AmountPrices               resp.Field
		BillingPlanID              resp.Field
		CreatedAt                  resp.Field
		Error                      resp.Field
		Eta                        resp.Field
		IsExpirationMessageVisible resp.Field
		Name                       resp.Field
		NextStatuses               resp.Field
		RegionID                   resp.Field
		RegionName                 resp.Field
		RemindExpirationMessage    resp.Field
		Resources                  resp.Field
		Status                     resp.Field
		UserStatus                 resp.Field
		ExtraFields                map[string]resp.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingReservation) RawJSON() string { return r.JSON.raw }
func (r *BillingReservation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/BillingReservationItemResponseSerializer/properties/amount_prices'
// "$.components.schemas.BillingReservationItemResponseSerializer.properties.amount_prices"
type BillingReservationAmountPrices struct {
	// '#/components/schemas/BillingReservationAmountPricesResponseSerializer/properties/commit_price_per_month'
	// "$.components.schemas.BillingReservationAmountPricesResponseSerializer.properties.commit_price_per_month"
	CommitPricePerMonth string `json:"commit_price_per_month,required"`
	// '#/components/schemas/BillingReservationAmountPricesResponseSerializer/properties/commit_price_per_unit'
	// "$.components.schemas.BillingReservationAmountPricesResponseSerializer.properties.commit_price_per_unit"
	CommitPricePerUnit string `json:"commit_price_per_unit,required"`
	// '#/components/schemas/BillingReservationAmountPricesResponseSerializer/properties/commit_price_total'
	// "$.components.schemas.BillingReservationAmountPricesResponseSerializer.properties.commit_price_total"
	CommitPriceTotal string `json:"commit_price_total,required"`
	// '#/components/schemas/BillingReservationAmountPricesResponseSerializer/properties/currency_code'
	// "$.components.schemas.BillingReservationAmountPricesResponseSerializer.properties.currency_code"
	CurrencyCode string `json:"currency_code,required"`
	// '#/components/schemas/BillingReservationAmountPricesResponseSerializer/properties/overcommit_price_per_month'
	// "$.components.schemas.BillingReservationAmountPricesResponseSerializer.properties.overcommit_price_per_month"
	OvercommitPricePerMonth string `json:"overcommit_price_per_month,required"`
	// '#/components/schemas/BillingReservationAmountPricesResponseSerializer/properties/overcommit_price_per_unit'
	// "$.components.schemas.BillingReservationAmountPricesResponseSerializer.properties.overcommit_price_per_unit"
	OvercommitPricePerUnit string `json:"overcommit_price_per_unit,required"`
	// '#/components/schemas/BillingReservationAmountPricesResponseSerializer/properties/overcommit_price_total'
	// "$.components.schemas.BillingReservationAmountPricesResponseSerializer.properties.overcommit_price_total"
	OvercommitPriceTotal string `json:"overcommit_price_total,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CommitPricePerMonth     resp.Field
		CommitPricePerUnit      resp.Field
		CommitPriceTotal        resp.Field
		CurrencyCode            resp.Field
		OvercommitPricePerMonth resp.Field
		OvercommitPricePerUnit  resp.Field
		OvercommitPriceTotal    resp.Field
		ExtraFields             map[string]resp.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingReservationAmountPrices) RawJSON() string { return r.JSON.raw }
func (r *BillingReservationAmountPrices) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/BillingReservationItemResponseSerializer/properties/resources/items'
// "$.components.schemas.BillingReservationItemResponseSerializer.properties.resources.items"
type BillingReservationResource struct {
	// '#/components/schemas/BillingReservationResourceSerializer/properties/activity_period'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.activity_period"
	ActivityPeriod string `json:"activity_period,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/activity_period_length'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.activity_period_length"
	ActivityPeriodLength int64 `json:"activity_period_length,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/billing_plan_item_id'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.billing_plan_item_id"
	BillingPlanItemID int64 `json:"billing_plan_item_id,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/commit_price_per_month'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.commit_price_per_month"
	CommitPricePerMonth string `json:"commit_price_per_month,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/commit_price_per_unit'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.commit_price_per_unit"
	CommitPricePerUnit string `json:"commit_price_per_unit,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/commit_price_total'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.commit_price_total"
	CommitPriceTotal string `json:"commit_price_total,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/overcommit_billing_plan_item_id'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.overcommit_billing_plan_item_id"
	OvercommitBillingPlanItemID int64 `json:"overcommit_billing_plan_item_id,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/overcommit_price_per_month'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.overcommit_price_per_month"
	OvercommitPricePerMonth string `json:"overcommit_price_per_month,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/overcommit_price_per_unit'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.overcommit_price_per_unit"
	OvercommitPricePerUnit string `json:"overcommit_price_per_unit,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/overcommit_price_total'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.overcommit_price_total"
	OvercommitPriceTotal string `json:"overcommit_price_total,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/resource_count'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.resource_count"
	ResourceCount int64 `json:"resource_count,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/resource_name'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.resource_name"
	ResourceName string `json:"resource_name,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/resource_type'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.resource_type"
	//
	// Any of "flavor".
	ResourceType string `json:"resource_type,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/unit_name'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.unit_name"
	UnitName string `json:"unit_name,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/unit_size_month'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.unit_size_month"
	UnitSizeMonth string `json:"unit_size_month,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/unit_size_total'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.unit_size_total"
	UnitSizeTotal string `json:"unit_size_total,required"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/cpu/anyOf/0'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.cpu.anyOf[0]"
	CPU string `json:"cpu,nullable"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/disk/anyOf/0'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.disk.anyOf[0]"
	Disk string `json:"disk,nullable"`
	// '#/components/schemas/BillingReservationResourceSerializer/properties/ram/anyOf/0'
	// "$.components.schemas.BillingReservationResourceSerializer.properties.ram.anyOf[0]"
	Ram string `json:"ram,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ActivityPeriod              resp.Field
		ActivityPeriodLength        resp.Field
		BillingPlanItemID           resp.Field
		CommitPricePerMonth         resp.Field
		CommitPricePerUnit          resp.Field
		CommitPriceTotal            resp.Field
		OvercommitBillingPlanItemID resp.Field
		OvercommitPricePerMonth     resp.Field
		OvercommitPricePerUnit      resp.Field
		OvercommitPriceTotal        resp.Field
		ResourceCount               resp.Field
		ResourceName                resp.Field
		ResourceType                resp.Field
		UnitName                    resp.Field
		UnitSizeMonth               resp.Field
		UnitSizeTotal               resp.Field
		CPU                         resp.Field
		Disk                        resp.Field
		Ram                         resp.Field
		ExtraFields                 map[string]resp.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingReservationResource) RawJSON() string { return r.JSON.raw }
func (r *BillingReservationResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingReservationListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Freservations/get/parameters/0'
	// "$.paths['/cloud/v1/reservations'].get.parameters[0]"
	ActivatedFrom param.Opt[time.Time] `query:"activated_from,omitzero" format:"date" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freservations/get/parameters/1'
	// "$.paths['/cloud/v1/reservations'].get.parameters[1]"
	ActivatedTo param.Opt[time.Time] `query:"activated_to,omitzero" format:"date" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freservations/get/parameters/2'
	// "$.paths['/cloud/v1/reservations'].get.parameters[2]"
	CreatedFrom param.Opt[time.Time] `query:"created_from,omitzero" format:"date-time" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freservations/get/parameters/3'
	// "$.paths['/cloud/v1/reservations'].get.parameters[3]"
	CreatedTo param.Opt[time.Time] `query:"created_to,omitzero" format:"date-time" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freservations/get/parameters/4'
	// "$.paths['/cloud/v1/reservations'].get.parameters[4]"
	DeactivatedFrom param.Opt[time.Time] `query:"deactivated_from,omitzero" format:"date" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freservations/get/parameters/5'
	// "$.paths['/cloud/v1/reservations'].get.parameters[5]"
	DeactivatedTo param.Opt[time.Time] `query:"deactivated_to,omitzero" format:"date" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freservations/get/parameters/6'
	// "$.paths['/cloud/v1/reservations'].get.parameters[6]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freservations/get/parameters/7'
	// "$.paths['/cloud/v1/reservations'].get.parameters[7]"
	MetricName param.Opt[string] `query:"metric_name,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freservations/get/parameters/8'
	// "$.paths['/cloud/v1/reservations'].get.parameters[8]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freservations/get/parameters/9'
	// "$.paths['/cloud/v1/reservations'].get.parameters[9]"
	RegionID param.Opt[int64] `query:"region_id,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Freservations/get/parameters/10'
	// "$.paths['/cloud/v1/reservations'].get.parameters[10]"
	//
	// Any of "ACTIVATED", "APPROVED", "COPIED", "CREATED", "EXPIRED", "REJECTED",
	// "RESERVED", "WAITING_FOR_PAYMENT".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BillingReservationListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [BillingReservationListParams]'s query parameters as
// `url.Values`.
func (r BillingReservationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
