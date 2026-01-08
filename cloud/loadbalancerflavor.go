// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// LoadBalancerFlavorService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoadBalancerFlavorService] method instead.
type LoadBalancerFlavorService struct {
	Options []option.RequestOption
}

// NewLoadBalancerFlavorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerFlavorService(opts ...option.RequestOption) (r LoadBalancerFlavorService) {
	r = LoadBalancerFlavorService{}
	r.Options = opts
	return
}

// Retrieve a list of load balancer flavors. When the `include_prices` query
// parameter is specified, the list shows prices. A client in trial mode gets all
// price values as 0. If you get Pricing Error contact the support
func (r *LoadBalancerFlavorService) List(ctx context.Context, params LoadBalancerFlavorListParams, opts ...option.RequestOption) (res *LoadBalancerFlavorListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/lbflavors/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type LoadBalancerFlavorListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []LoadBalancerFlavorListResponseResultUnion `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerFlavorListResponse) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerFlavorListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LoadBalancerFlavorListResponseResultUnion contains all possible properties and
// values from [LoadBalancerFlavorListResponseResultLbFlavorSerializer],
// [LoadBalancerFlavorDetail].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LoadBalancerFlavorListResponseResultUnion struct {
	FlavorID   string `json:"flavor_id"`
	FlavorName string `json:"flavor_name"`
	Ram        int64  `json:"ram"`
	Vcpus      int64  `json:"vcpus"`
	// This field is from variant [LoadBalancerFlavorDetail].
	HardwareDescription LoadBalancerFlavorDetailHardwareDescriptionUnion `json:"hardware_description"`
	// This field is from variant [LoadBalancerFlavorDetail].
	CurrencyCode string `json:"currency_code"`
	// This field is from variant [LoadBalancerFlavorDetail].
	PricePerHour float64 `json:"price_per_hour"`
	// This field is from variant [LoadBalancerFlavorDetail].
	PricePerMonth float64 `json:"price_per_month"`
	// This field is from variant [LoadBalancerFlavorDetail].
	PriceStatus LoadBalancerFlavorDetailPriceStatus `json:"price_status"`
	JSON        struct {
		FlavorID            respjson.Field
		FlavorName          respjson.Field
		Ram                 respjson.Field
		Vcpus               respjson.Field
		HardwareDescription respjson.Field
		CurrencyCode        respjson.Field
		PricePerHour        respjson.Field
		PricePerMonth       respjson.Field
		PriceStatus         respjson.Field
		raw                 string
	} `json:"-"`
}

func (u LoadBalancerFlavorListResponseResultUnion) AsLbFlavorSerializer() (v LoadBalancerFlavorListResponseResultLbFlavorSerializer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LoadBalancerFlavorListResponseResultUnion) AsOptionalPriceSchema() (v LoadBalancerFlavorDetail) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LoadBalancerFlavorListResponseResultUnion) RawJSON() string { return u.JSON.raw }

func (r *LoadBalancerFlavorListResponseResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerFlavorListResponseResultLbFlavorSerializer struct {
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id,required"`
	// Flavor name
	FlavorName string `json:"flavor_name,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Virtual CPU count. For bare metal flavors, it's a physical CPU count
	Vcpus int64 `json:"vcpus,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FlavorID    respjson.Field
		FlavorName  respjson.Field
		Ram         respjson.Field
		Vcpus       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LoadBalancerFlavorListResponseResultLbFlavorSerializer) RawJSON() string { return r.JSON.raw }
func (r *LoadBalancerFlavorListResponseResultLbFlavorSerializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerFlavorListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Set to true if the response should include flavor prices
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LoadBalancerFlavorListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
