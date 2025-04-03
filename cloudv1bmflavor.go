// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1BmflavorService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1BmflavorService] method instead.
type CloudV1BmflavorService struct {
	Options []option.RequestOption
}

// NewCloudV1BmflavorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1BmflavorService(opts ...option.RequestOption) (r *CloudV1BmflavorService) {
	r = &CloudV1BmflavorService{}
	r.Options = opts
	return
}

// Retrieve a list of flavors. When the include_prices query parameter is
// specified, the list shows prices. A client in trial mode gets all price values
// as 0. If you get Pricing Error contact the support
func (r *CloudV1BmflavorService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1BmflavorListParams, opts ...option.RequestOption) (res *BaremetalFlavorList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/bmflavors/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a list of flavors. When the include_prices query parameter is
// specified, the list shows prices. A client in trial mode gets all price values
// as 0. If you get Pricing Error contact the support
func (r *CloudV1BmflavorService) ListDefault(ctx context.Context, regionID int64, query CloudV1BmflavorListDefaultParams, opts ...option.RequestOption) (res *BaremetalFlavorList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/bmflavors/%v", regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BaremetalFlavorList struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []BaremetalFlavorListResult `json:"results"`
	JSON    baremetalFlavorListJSON     `json:"-"`
}

// baremetalFlavorListJSON contains the JSON metadata for the struct
// [BaremetalFlavorList]
type baremetalFlavorListJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BaremetalFlavorList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r baremetalFlavorListJSON) RawJSON() string {
	return r.raw
}

// Bare metal flavors
type BaremetalFlavorListResult struct {
	// Currency code. Shown if the include_prices query parameter if set to true
	CurrencyCode string `json:"currency_code"`
	// Disabled flavor flag
	Disabled bool `json:"disabled"`
	// Short ID
	FlavorID string `json:"flavor_id"`
	// Human readable name
	FlavorName string `json:"flavor_name"`
	// Various hardware hints for flavor.
	HardwareDescription DeprecatedFlavorHardwareDescription `json:"hardware_description"`
	// Price per hour. Shown if the include_prices query parameter if set to true
	PricePerHour float64 `json:"price_per_hour"`
	// Price per month. Shown if the include_prices query parameter if set to true
	PricePerMonth float64 `json:"price_per_month"`
	// Price status for the UI
	PriceStatus BaremetalFlavorListResultsPriceStatus `json:"price_status"`
	// Flavor resource class for mapping to hardware capacity
	ResourceClass string                        `json:"resource_class"`
	JSON          baremetalFlavorListResultJSON `json:"-"`
}

// baremetalFlavorListResultJSON contains the JSON metadata for the struct
// [BaremetalFlavorListResult]
type baremetalFlavorListResultJSON struct {
	CurrencyCode        apijson.Field
	Disabled            apijson.Field
	FlavorID            apijson.Field
	FlavorName          apijson.Field
	HardwareDescription apijson.Field
	PricePerHour        apijson.Field
	PricePerMonth       apijson.Field
	PriceStatus         apijson.Field
	ResourceClass       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *BaremetalFlavorListResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r baremetalFlavorListResultJSON) RawJSON() string {
	return r.raw
}

// Price status for the UI
type BaremetalFlavorListResultsPriceStatus string

const (
	BaremetalFlavorListResultsPriceStatusError BaremetalFlavorListResultsPriceStatus = "error"
	BaremetalFlavorListResultsPriceStatusHide  BaremetalFlavorListResultsPriceStatus = "hide"
	BaremetalFlavorListResultsPriceStatusShow  BaremetalFlavorListResultsPriceStatus = "show"
)

func (r BaremetalFlavorListResultsPriceStatus) IsKnown() bool {
	switch r {
	case BaremetalFlavorListResultsPriceStatusError, BaremetalFlavorListResultsPriceStatusHide, BaremetalFlavorListResultsPriceStatusShow:
		return true
	}
	return false
}

type CloudV1BmflavorListParams struct {
	// Flag for filtering disabled flavors in the region. Defaults to true
	Disabled param.Field[bool] `query:"disabled"`
	// Set to true to exclude flavors dedicated to linux images. Default False
	ExcludeLinux param.Field[bool] `query:"exclude_linux"`
	// Set to true to exclude flavors dedicated to windows images. Default False
	ExcludeWindows param.Field[bool] `query:"exclude_windows"`
	// Set to true if the response should include flavor capacity
	IncludeCapacity param.Field[bool] `query:"include_capacity"`
	// Set to true if the response should include flavor prices
	IncludePrices param.Field[bool] `query:"include_prices"`
	// Optional. Set to true if flavor listing should include count of reserved
	// resources in stock.
	IncludeReservationStock param.Field[bool] `query:"include_reservation_stock"`
}

// URLQuery serializes [CloudV1BmflavorListParams]'s query parameters as
// `url.Values`.
func (r CloudV1BmflavorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1BmflavorListDefaultParams struct {
	// Client identifier. Must be used for users w/o client_id in jwt
	ClientID param.Field[int64] `query:"client_id_"`
	// Flag for filtering disabled flavors in the region. Defaults to true
	Disabled param.Field[bool] `query:"disabled"`
	// Set to true to exclude flavors dedicated to linux images. Default False
	ExcludeLinux param.Field[bool] `query:"exclude_linux"`
	// Set to true to exclude flavors dedicated to windows images. Default False
	ExcludeWindows param.Field[bool] `query:"exclude_windows"`
	// Set to true if the response should include flavor capacity
	IncludeCapacity param.Field[bool] `query:"include_capacity"`
	// Set to true if the response should include flavor prices
	IncludePrices param.Field[bool] `query:"include_prices"`
	// Optional. Set to true if flavor listing should include count of reserved
	// resources in stock.
	IncludeReservationStock param.Field[bool] `query:"include_reservation_stock"`
}

// URLQuery serializes [CloudV1BmflavorListDefaultParams]'s query parameters as
// `url.Values`.
func (r CloudV1BmflavorListDefaultParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
