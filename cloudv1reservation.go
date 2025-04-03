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

// CloudV1ReservationService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1ReservationService] method instead.
type CloudV1ReservationService struct {
	Options []option.RequestOption
}

// NewCloudV1ReservationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1ReservationService(opts ...option.RequestOption) (r *CloudV1ReservationService) {
	r = &CloudV1ReservationService{}
	r.Options = opts
	return
}

// Create reservation
func (r *CloudV1ReservationService) New(ctx context.Context, body CloudV1ReservationNewParams, opts ...option.RequestOption) (res *BillingReservationItem, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/reservations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get specific reservation
func (r *CloudV1ReservationService) Get(ctx context.Context, reservationID int64, opts ...option.RequestOption) (res *BillingReservationItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/reservations/%v", reservationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List reservations
func (r *CloudV1ReservationService) List(ctx context.Context, query CloudV1ReservationListParams, opts ...option.RequestOption) (res *CloudV1ReservationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/reservations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete reservation
func (r *CloudV1ReservationService) Delete(ctx context.Context, reservationID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("cloud/v1/reservations/%v", reservationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type BillingReservationItem struct {
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
	Error string `json:"error,required,nullable"`
	// ETA delivery if bare metal out of stock. Value None means that bare metal in
	// stock.
	Eta time.Time `json:"eta,required,nullable" format:"date"`
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
	RemindExpirationMessage time.Time `json:"remind_expiration_message,required,nullable" format:"date"`
	// List of reservation resources
	Resources []BillingReservationItemResource `json:"resources,required"`
	// Reservation status
	Status string `json:"status,required"`
	// User status
	UserStatus string                     `json:"user_status,required"`
	JSON       billingReservationItemJSON `json:"-"`
}

// billingReservationItemJSON contains the JSON metadata for the struct
// [BillingReservationItem]
type billingReservationItemJSON struct {
	ID                         apijson.Field
	ActiveFrom                 apijson.Field
	ActiveTo                   apijson.Field
	ActivityPeriod             apijson.Field
	ActivityPeriodLength       apijson.Field
	AmountPrices               apijson.Field
	BillingPlanID              apijson.Field
	CreatedAt                  apijson.Field
	Error                      apijson.Field
	Eta                        apijson.Field
	IsExpirationMessageVisible apijson.Field
	Name                       apijson.Field
	NextStatuses               apijson.Field
	RegionID                   apijson.Field
	RegionName                 apijson.Field
	RemindExpirationMessage    apijson.Field
	Resources                  apijson.Field
	Status                     apijson.Field
	UserStatus                 apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *BillingReservationItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingReservationItemJSON) RawJSON() string {
	return r.raw
}

type BillingReservationItemResource struct {
	// Name of the billing period, e.g month
	ActivityPeriod string `json:"activity_period,required"`
	// Length of the full reservation period by `activity_period`
	ActivityPeriodLength int64 `json:"activity_period_length,required"`
	// Billing plan item id
	BillingPlanItemID int64 `json:"billing_plan_item_id,required"`
	// Commit price of the item charged per month
	CommitPricePerMonth BillingReservationItemResourcesCommitPricePerMonthUnion `json:"commit_price_per_month,required"`
	// Commit price of the item charged per hour
	CommitPricePerUnit BillingReservationItemResourcesCommitPricePerUnitUnion `json:"commit_price_per_unit,required"`
	// Commit price of the item charged for all period reservation
	CommitPriceTotal BillingReservationItemResourcesCommitPriceTotalUnion `json:"commit_price_total,required"`
	// Overcommit billing plan item id
	OvercommitBillingPlanItemID int64 `json:"overcommit_billing_plan_item_id,required"`
	// Overcommit price of the item charged per month
	OvercommitPricePerMonth BillingReservationItemResourcesOvercommitPricePerMonthUnion `json:"overcommit_price_per_month,required"`
	// Overcommit price of the item charged per hour
	OvercommitPricePerUnit BillingReservationItemResourcesOvercommitPricePerUnitUnion `json:"overcommit_price_per_unit,required"`
	// Overcommit price of the item charged for all period reservation
	OvercommitPriceTotal BillingReservationItemResourcesOvercommitPriceTotalUnion `json:"overcommit_price_total,required"`
	// Number of reserved resource items
	ResourceCount int64 `json:"resource_count,required"`
	// Resource name
	ResourceName string `json:"resource_name,required"`
	// Resource type
	ResourceType BillingReservationType `json:"resource_type,required"`
	// Billing unit name
	UnitName string `json:"unit_name,required"`
	// Minimal billing size, for example it is 744 hours per 1 month.
	UnitSizeMonth BillingReservationItemResourcesUnitSizeMonthUnion `json:"unit_size_month,required"`
	// Unit size month multiplied by count of resources in the reservation
	UnitSizeTotal BillingReservationItemResourcesUnitSizeTotalUnion `json:"unit_size_total,required"`
	// Baremetal CPU description
	CPU string `json:"cpu,nullable"`
	// Baremetal disk description
	Disk string `json:"disk,nullable"`
	// Baremetal RAM description
	Ram  string                             `json:"ram,nullable"`
	JSON billingReservationItemResourceJSON `json:"-"`
}

// billingReservationItemResourceJSON contains the JSON metadata for the struct
// [BillingReservationItemResource]
type billingReservationItemResourceJSON struct {
	ActivityPeriod              apijson.Field
	ActivityPeriodLength        apijson.Field
	BillingPlanItemID           apijson.Field
	CommitPricePerMonth         apijson.Field
	CommitPricePerUnit          apijson.Field
	CommitPriceTotal            apijson.Field
	OvercommitBillingPlanItemID apijson.Field
	OvercommitPricePerMonth     apijson.Field
	OvercommitPricePerUnit      apijson.Field
	OvercommitPriceTotal        apijson.Field
	ResourceCount               apijson.Field
	ResourceName                apijson.Field
	ResourceType                apijson.Field
	UnitName                    apijson.Field
	UnitSizeMonth               apijson.Field
	UnitSizeTotal               apijson.Field
	CPU                         apijson.Field
	Disk                        apijson.Field
	Ram                         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *BillingReservationItemResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingReservationItemResourceJSON) RawJSON() string {
	return r.raw
}

// Commit price of the item charged per month
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type BillingReservationItemResourcesCommitPricePerMonthUnion interface {
	ImplementsBillingReservationItemResourcesCommitPricePerMonthUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingReservationItemResourcesCommitPricePerMonthUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Commit price of the item charged per hour
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type BillingReservationItemResourcesCommitPricePerUnitUnion interface {
	ImplementsBillingReservationItemResourcesCommitPricePerUnitUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingReservationItemResourcesCommitPricePerUnitUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Commit price of the item charged for all period reservation
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type BillingReservationItemResourcesCommitPriceTotalUnion interface {
	ImplementsBillingReservationItemResourcesCommitPriceTotalUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingReservationItemResourcesCommitPriceTotalUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Overcommit price of the item charged per month
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type BillingReservationItemResourcesOvercommitPricePerMonthUnion interface {
	ImplementsBillingReservationItemResourcesOvercommitPricePerMonthUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingReservationItemResourcesOvercommitPricePerMonthUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Overcommit price of the item charged per hour
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type BillingReservationItemResourcesOvercommitPricePerUnitUnion interface {
	ImplementsBillingReservationItemResourcesOvercommitPricePerUnitUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingReservationItemResourcesOvercommitPricePerUnitUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Overcommit price of the item charged for all period reservation
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type BillingReservationItemResourcesOvercommitPriceTotalUnion interface {
	ImplementsBillingReservationItemResourcesOvercommitPriceTotalUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingReservationItemResourcesOvercommitPriceTotalUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Minimal billing size, for example it is 744 hours per 1 month.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type BillingReservationItemResourcesUnitSizeMonthUnion interface {
	ImplementsBillingReservationItemResourcesUnitSizeMonthUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingReservationItemResourcesUnitSizeMonthUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit size month multiplied by count of resources in the reservation
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type BillingReservationItemResourcesUnitSizeTotalUnion interface {
	ImplementsBillingReservationItemResourcesUnitSizeTotalUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingReservationItemResourcesUnitSizeTotalUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1ReservationListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []BillingReservationItem           `json:"results,required"`
	JSON    cloudV1ReservationListResponseJSON `json:"-"`
}

// cloudV1ReservationListResponseJSON contains the JSON metadata for the struct
// [CloudV1ReservationListResponse]
type cloudV1ReservationListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1ReservationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1ReservationListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1ReservationNewParams struct {
	Name       param.Field[string]                                          `json:"name,required"`
	Period     param.Field[CloudV1ReservationNewParamsPeriod]               `json:"period,required"`
	RegionID   param.Field[int64]                                           `json:"region_id,required"`
	Resources  param.Field[[]BillingReservationPricingResourceRequestParam] `json:"resources,required"`
	HasWindows param.Field[bool]                                            `json:"has_windows"`
}

func (r CloudV1ReservationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1ReservationNewParamsPeriod string

const (
	CloudV1ReservationNewParamsPeriodMonths12 CloudV1ReservationNewParamsPeriod = "MONTHS12"
	CloudV1ReservationNewParamsPeriodMonths36 CloudV1ReservationNewParamsPeriod = "MONTHS36"
)

func (r CloudV1ReservationNewParamsPeriod) IsKnown() bool {
	switch r {
	case CloudV1ReservationNewParamsPeriodMonths12, CloudV1ReservationNewParamsPeriodMonths36:
		return true
	}
	return false
}

type CloudV1ReservationListParams struct {
	// Lower bound, starting from what date the reservation was/will be activated
	ActivatedFrom param.Field[time.Time] `query:"activated_from" format:"date"`
	// High bound, before what date the reservation was/will be activated
	ActivatedTo param.Field[time.Time] `query:"activated_to" format:"date"`
	// Lower bound the filter, showing result(s) equal to or greater than date the
	// reservation was created
	CreatedFrom param.Field[time.Time] `query:"created_from" format:"date-time"`
	// High bound the filter, showing result(s) equal to or less date the reservation
	// was created
	CreatedTo param.Field[time.Time] `query:"created_to" format:"date-time"`
	// Lower bound, starting from what date the reservation was/will be deactivated
	DeactivatedFrom param.Field[time.Time] `query:"deactivated_from" format:"date"`
	// High bound, before what date the reservation was/will be deactivated
	DeactivatedTo param.Field[time.Time] `query:"deactivated_to" format:"date"`
	// Limit of reservation list page
	Limit param.Field[int64] `query:"limit"`
	// Name from billing features for specific resource
	MetricName param.Field[string] `query:"metric_name"`
	// Offset in reservation list
	Offset param.Field[int64] `query:"offset"`
	// Region for reservation
	RegionID param.Field[int64] `query:"region_id"`
	// Field for fixed a status by reservation workflow
	Status param.Field[[]CloudV1ReservationListParamsStatus] `query:"status"`
}

// URLQuery serializes [CloudV1ReservationListParams]'s query parameters as
// `url.Values`.
func (r CloudV1ReservationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1ReservationListParamsStatus string

const (
	CloudV1ReservationListParamsStatusActivated         CloudV1ReservationListParamsStatus = "ACTIVATED"
	CloudV1ReservationListParamsStatusApproved          CloudV1ReservationListParamsStatus = "APPROVED"
	CloudV1ReservationListParamsStatusCopied            CloudV1ReservationListParamsStatus = "COPIED"
	CloudV1ReservationListParamsStatusCreated           CloudV1ReservationListParamsStatus = "CREATED"
	CloudV1ReservationListParamsStatusExpired           CloudV1ReservationListParamsStatus = "EXPIRED"
	CloudV1ReservationListParamsStatusRejected          CloudV1ReservationListParamsStatus = "REJECTED"
	CloudV1ReservationListParamsStatusReserved          CloudV1ReservationListParamsStatus = "RESERVED"
	CloudV1ReservationListParamsStatusWaitingForPayment CloudV1ReservationListParamsStatus = "WAITING_FOR_PAYMENT"
)

func (r CloudV1ReservationListParamsStatus) IsKnown() bool {
	switch r {
	case CloudV1ReservationListParamsStatusActivated, CloudV1ReservationListParamsStatusApproved, CloudV1ReservationListParamsStatusCopied, CloudV1ReservationListParamsStatusCreated, CloudV1ReservationListParamsStatusExpired, CloudV1ReservationListParamsStatusRejected, CloudV1ReservationListParamsStatusReserved, CloudV1ReservationListParamsStatusWaitingForPayment:
		return true
	}
	return false
}
