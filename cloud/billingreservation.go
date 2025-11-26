// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
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

// Get a list of billing reservations along with detailed information on resource
// configurations and associated pricing.
func (r *BillingReservationService) List(ctx context.Context, query BillingReservationListParams, opts ...option.RequestOption) (res *BillingReservations, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cloud/v2/reservations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BillingReservation struct {
	// Active billing plan ID
	ActiveBillingPlanID int64 `json:"active_billing_plan_id,required"`
	// Overcommit pricing details
	ActiveOvercommit BillingReservationActiveOvercommit `json:"active_overcommit,required"`
	// Commit pricing details
	Commit BillingReservationCommit `json:"commit,required"`
	// Hardware specifications
	HardwareInfo BillingReservationHardwareInfo `json:"hardware_info,required"`
	// Region name
	RegionName string `json:"region_name,required"`
	// Number of reserved resource items
	ResourceCount int64 `json:"resource_count,required"`
	// Resource name
	ResourceName string `json:"resource_name,required"`
	// Unit name (e.g., 'H' for hours)
	UnitName string `json:"unit_name,required"`
	// Unit size per month (e.g., 744 hours)
	UnitSizeMonth string `json:"unit_size_month,required"`
	// Unit size month multiplied by count of resources in the reservation
	UnitSizeTotal string `json:"unit_size_total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveBillingPlanID respjson.Field
		ActiveOvercommit    respjson.Field
		Commit              respjson.Field
		HardwareInfo        respjson.Field
		RegionName          respjson.Field
		ResourceCount       respjson.Field
		ResourceName        respjson.Field
		UnitName            respjson.Field
		UnitSizeMonth       respjson.Field
		UnitSizeTotal       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingReservation) RawJSON() string { return r.JSON.raw }
func (r *BillingReservation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Overcommit pricing details
type BillingReservationActiveOvercommit struct {
	// Billing subscription active from date
	ActiveFrom time.Time `json:"active_from,required" format:"date-time"`
	// Billing plan item ID
	PlanItemID int64 `json:"plan_item_id,required"`
	// Price per month
	PricePerMonth string `json:"price_per_month,required"`
	// Price per unit (hourly)
	PricePerUnit string `json:"price_per_unit,required"`
	// Total price for the reservation period
	PriceTotal string `json:"price_total,required"`
	// Billing subscription ID for overcommit
	SubscriptionID int64 `json:"subscription_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveFrom     respjson.Field
		PlanItemID     respjson.Field
		PricePerMonth  respjson.Field
		PricePerUnit   respjson.Field
		PriceTotal     respjson.Field
		SubscriptionID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingReservationActiveOvercommit) RawJSON() string { return r.JSON.raw }
func (r *BillingReservationActiveOvercommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Commit pricing details
type BillingReservationCommit struct {
	// Billing subscription active from date
	ActiveFrom time.Time `json:"active_from,required" format:"date-time"`
	// Billing subscription active to date
	ActiveTo time.Time `json:"active_to,required" format:"date-time"`
	// Price per month, per one resource
	PricePerMonth string `json:"price_per_month,required"`
	// Price per unit, per one resource (hourly)
	PricePerUnit string `json:"price_per_unit,required"`
	// Total price for the reservation period for the full reserved amount
	PriceTotal string `json:"price_total,required"`
	// Billing subscription ID for commit
	SubscriptionID int64 `json:"subscription_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveFrom     respjson.Field
		ActiveTo       respjson.Field
		PricePerMonth  respjson.Field
		PricePerUnit   respjson.Field
		PriceTotal     respjson.Field
		SubscriptionID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingReservationCommit) RawJSON() string { return r.JSON.raw }
func (r *BillingReservationCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Hardware specifications
type BillingReservationHardwareInfo struct {
	// CPU specification
	CPU string `json:"cpu,required"`
	// Disk specification
	Disk string `json:"disk,required"`
	// RAM specification
	Ram string `json:"ram,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU         respjson.Field
		Disk        respjson.Field
		Ram         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingReservationHardwareInfo) RawJSON() string { return r.JSON.raw }
func (r *BillingReservationHardwareInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingReservations struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []BillingReservation `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingReservations) RawJSON() string { return r.JSON.raw }
func (r *BillingReservations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingReservationListParams struct {
	// Metric name for the resource (e.g., 'bm1-hf-`medium_min`')
	MetricName param.Opt[string] `query:"metric_name,omitzero" json:"-"`
	// Region for reservation
	RegionID param.Opt[int64] `query:"region_id,omitzero" json:"-"`
	// Include inactive commits in the response
	ShowInactive param.Opt[bool] `query:"show_inactive,omitzero" json:"-"`
	// Order by field and direction.
	//
	// Any of "active_from.asc", "active_from.desc", "active_to.asc", "active_to.desc".
	OrderBy BillingReservationListParamsOrderBy `query:"order_by,omitzero" json:"-"`
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
)
