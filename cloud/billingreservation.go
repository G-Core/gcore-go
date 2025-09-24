// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
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

// Get reservation
func (r *BillingReservationService) Get(ctx context.Context, reservationID int64, opts ...option.RequestOption) (res *BillingReservation, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("cloud/v1/reservations/%v", reservationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BillingReservation struct {
	// Reservation id
	ID int64 `json:"id,required"`
	// Reservation active from date
	ActiveFrom time.Time `json:"active_from,required" format:"date"`
	// Reservation active to date
	ActiveTo time.Time `json:"active_to,required" format:"date"`
	// Name of the billing period, e.g month
	ActivityPeriod string `json:"activity_period,required"`
	// Length of the full reservation period by `activity_period`
	ActivityPeriodLength int64 `json:"activity_period_length,required"`
	// Reservation amount prices
	AmountPrices BillingReservationAmountPrices `json:"amount_prices,required"`
	// Billing plan id
	BillingPlanID int64 `json:"billing_plan_id,required"`
	// Reservation creation date
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Error message if any occured during reservation
	Error string `json:"error,required"`
	// ETA delivery if bare metal out of stock. Value None means that bare metal in
	// stock.
	Eta time.Time `json:"eta,required" format:"date"`
	// Hide or show expiration message to customer.
	IsExpirationMessageVisible bool `json:"is_expiration_message_visible,required"`
	// Reservation name
	Name string `json:"name,required"`
	// List of possible next reservation statuses
	NextStatuses []string `json:"next_statuses,required"`
	// Region id
	RegionID int64 `json:"region_id,required"`
	// Region name
	RegionName string `json:"region_name,required"`
	// The date when show expiration date to customer
	RemindExpirationMessage time.Time `json:"remind_expiration_message,required" format:"date"`
	// List of reservation resources
	Resources []BillingReservationResource `json:"resources,required"`
	// Reservation status
	Status string `json:"status,required"`
	// User status
	UserStatus string `json:"user_status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		ActiveFrom                 respjson.Field
		ActiveTo                   respjson.Field
		ActivityPeriod             respjson.Field
		ActivityPeriodLength       respjson.Field
		AmountPrices               respjson.Field
		BillingPlanID              respjson.Field
		CreatedAt                  respjson.Field
		Error                      respjson.Field
		Eta                        respjson.Field
		IsExpirationMessageVisible respjson.Field
		Name                       respjson.Field
		NextStatuses               respjson.Field
		RegionID                   respjson.Field
		RegionName                 respjson.Field
		RemindExpirationMessage    respjson.Field
		Resources                  respjson.Field
		Status                     respjson.Field
		UserStatus                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingReservation) RawJSON() string { return r.JSON.raw }
func (r *BillingReservation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reservation amount prices
type BillingReservationAmountPrices struct {
	// Commit price of the item charged per month
	CommitPricePerMonth string `json:"commit_price_per_month,required"`
	// Commit price of the item charged per hour
	CommitPricePerUnit string `json:"commit_price_per_unit,required"`
	// Commit price of the item charged for all period reservation
	CommitPriceTotal string `json:"commit_price_total,required"`
	// Currency code (3 letter code per ISO 4217)
	CurrencyCode string `json:"currency_code,required"`
	// Overcommit price of the item charged per month
	OvercommitPricePerMonth string `json:"overcommit_price_per_month,required"`
	// Overcommit price of the item charged per hour
	OvercommitPricePerUnit string `json:"overcommit_price_per_unit,required"`
	// Overcommit price of the item charged for all period reservation
	OvercommitPriceTotal string `json:"overcommit_price_total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommitPricePerMonth     respjson.Field
		CommitPricePerUnit      respjson.Field
		CommitPriceTotal        respjson.Field
		CurrencyCode            respjson.Field
		OvercommitPricePerMonth respjson.Field
		OvercommitPricePerUnit  respjson.Field
		OvercommitPriceTotal    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingReservationAmountPrices) RawJSON() string { return r.JSON.raw }
func (r *BillingReservationAmountPrices) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingReservationResource struct {
	// Name of the billing period, e.g month
	ActivityPeriod string `json:"activity_period,required"`
	// Length of the full reservation period by `activity_period`
	ActivityPeriodLength int64 `json:"activity_period_length,required"`
	// Billing plan item id
	BillingPlanItemID int64 `json:"billing_plan_item_id,required"`
	// Commit price of the item charged per month
	CommitPricePerMonth string `json:"commit_price_per_month,required"`
	// Commit price of the item charged per hour
	CommitPricePerUnit string `json:"commit_price_per_unit,required"`
	// Commit price of the item charged for all period reservation
	CommitPriceTotal string `json:"commit_price_total,required"`
	// Overcommit billing plan item id
	OvercommitBillingPlanItemID int64 `json:"overcommit_billing_plan_item_id,required"`
	// Overcommit price of the item charged per month
	OvercommitPricePerMonth string `json:"overcommit_price_per_month,required"`
	// Overcommit price of the item charged per hour
	OvercommitPricePerUnit string `json:"overcommit_price_per_unit,required"`
	// Overcommit price of the item charged for all period reservation
	OvercommitPriceTotal string `json:"overcommit_price_total,required"`
	// Number of reserved resource items
	ResourceCount int64 `json:"resource_count,required"`
	// Resource name
	ResourceName string `json:"resource_name,required"`
	// Resource type
	//
	// Any of "flavor".
	ResourceType string `json:"resource_type,required"`
	// Billing unit name
	UnitName string `json:"unit_name,required"`
	// Minimal billing size, for example it is 744 hours per 1 month.
	UnitSizeMonth string `json:"unit_size_month,required"`
	// Unit size month multiplied by count of resources in the reservation
	UnitSizeTotal string `json:"unit_size_total,required"`
	// Baremetal CPU description
	CPU string `json:"cpu,nullable"`
	// Baremetal disk description
	Disk string `json:"disk,nullable"`
	// Baremetal RAM description
	Ram string `json:"ram,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActivityPeriod              respjson.Field
		ActivityPeriodLength        respjson.Field
		BillingPlanItemID           respjson.Field
		CommitPricePerMonth         respjson.Field
		CommitPricePerUnit          respjson.Field
		CommitPriceTotal            respjson.Field
		OvercommitBillingPlanItemID respjson.Field
		OvercommitPricePerMonth     respjson.Field
		OvercommitPricePerUnit      respjson.Field
		OvercommitPriceTotal        respjson.Field
		ResourceCount               respjson.Field
		ResourceName                respjson.Field
		ResourceType                respjson.Field
		UnitName                    respjson.Field
		UnitSizeMonth               respjson.Field
		UnitSizeTotal               respjson.Field
		CPU                         respjson.Field
		Disk                        respjson.Field
		Ram                         respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingReservationResource) RawJSON() string { return r.JSON.raw }
func (r *BillingReservationResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingReservationListParams struct {
	// Lower bound, starting from what date the reservation was/will be activated
	ActivatedFrom param.Opt[time.Time] `query:"activated_from,omitzero" format:"date" json:"-"`
	// High bound, before what date the reservation was/will be activated
	ActivatedTo param.Opt[time.Time] `query:"activated_to,omitzero" format:"date" json:"-"`
	// Lower bound the filter, showing result(s) equal to or greater than date the
	// reservation was created
	CreatedFrom param.Opt[time.Time] `query:"created_from,omitzero" format:"date-time" json:"-"`
	// High bound the filter, showing result(s) equal to or less date the reservation
	// was created
	CreatedTo param.Opt[time.Time] `query:"created_to,omitzero" format:"date-time" json:"-"`
	// Lower bound, starting from what date the reservation was/will be deactivated
	DeactivatedFrom param.Opt[time.Time] `query:"deactivated_from,omitzero" format:"date" json:"-"`
	// High bound, before what date the reservation was/will be deactivated
	DeactivatedTo param.Opt[time.Time] `query:"deactivated_to,omitzero" format:"date" json:"-"`
	// Limit of reservation list page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Name from billing features for specific resource
	MetricName param.Opt[string] `query:"metric_name,omitzero" json:"-"`
	// Offset in reservation list
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Region for reservation
	RegionID param.Opt[int64] `query:"region_id,omitzero" json:"-"`
	// Order by field and direction.
	//
	// Any of "active_from.asc", "active_from.desc", "active_to.asc", "active_to.desc",
	// "created_at.asc", "created_at.desc".
	OrderBy BillingReservationListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Field for fixed a status by reservation workflow
	//
	// Any of "ACTIVATED", "APPROVED", "COPIED", "CREATED", "EXPIRED", "REJECTED",
	// "RESERVED", "WAITING_FOR_PAYMENT".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BillingReservationListParams]'s query parameters as
// `url.Values`.
func (r BillingReservationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Order by field and direction.
type BillingReservationListParamsOrderBy string

const (
	BillingReservationListParamsOrderByActiveFromAsc  BillingReservationListParamsOrderBy = "active_from.asc"
	BillingReservationListParamsOrderByActiveFromDesc BillingReservationListParamsOrderBy = "active_from.desc"
	BillingReservationListParamsOrderByActiveToAsc    BillingReservationListParamsOrderBy = "active_to.asc"
	BillingReservationListParamsOrderByActiveToDesc   BillingReservationListParamsOrderBy = "active_to.desc"
	BillingReservationListParamsOrderByCreatedAtAsc   BillingReservationListParamsOrderBy = "created_at.asc"
	BillingReservationListParamsOrderByCreatedAtDesc  BillingReservationListParamsOrderBy = "created_at.desc"
)
