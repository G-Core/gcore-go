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
	// Field for fixed a status by reservation workflow
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
