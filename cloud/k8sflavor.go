// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
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

// K8SFlavorService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewK8SFlavorService] method instead.
type K8SFlavorService struct {
	Options []option.RequestOption
}

// NewK8SFlavorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewK8SFlavorService(opts ...option.RequestOption) (r K8SFlavorService) {
	r = K8SFlavorService{}
	r.Options = opts
	return
}

// Retrieve a list of flavors for k8s pool. When the `include_prices` query
// parameter is specified, the list shows prices. A client in trial mode gets all
// price values as 0. If you get Pricing Error contact the support
func (r *K8SFlavorService) List(ctx context.Context, params K8SFlavorListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[K8SFlavor], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/k8s/%v/%v/flavors", params.ProjectID.Value, params.RegionID.Value)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Retrieve a list of flavors for k8s pool. When the `include_prices` query
// parameter is specified, the list shows prices. A client in trial mode gets all
// price values as 0. If you get Pricing Error contact the support
func (r *K8SFlavorService) ListAutoPaging(ctx context.Context, params K8SFlavorListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[K8SFlavor] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Kubernetes flavor schema. The flavors on this route are general flavors, so they
// may have no resource class.
type K8SFlavor struct {
	// Flavor architecture type
	Architecture string `json:"architecture" api:"required"`
	// Number of available instances of given configuration. The value is null if it is
	// not loaded.
	Capacity int64 `json:"capacity" api:"required"`
	// Currency code. The value is null if the price is not available.
	CurrencyCode string `json:"currency_code" api:"required"`
	// Disabled flavor flag
	Disabled bool `json:"disabled" api:"required"`
	// Flavor ID is the same as name
	FlavorID string `json:"flavor_id" api:"required"`
	// Flavor name
	FlavorName string `json:"flavor_name" api:"required"`
	// Additional hardware description
	HardwareDescription map[string]string `json:"hardware_description" api:"required"`
	// Flavor operating system
	OsType string `json:"os_type" api:"required"`
	// Price per hour. The value is null if the price is not available.
	PricePerHour float64 `json:"price_per_hour" api:"required"`
	// Price per month. The value is null if the price is not available.
	PricePerMonth float64 `json:"price_per_month" api:"required"`
	// Price status for the UI. The value is null if the price is not available.
	//
	// Any of "error", "hide", "show".
	PriceStatus K8SFlavorPriceStatus `json:"price_status" api:"required"`
	// RAM size in MiB
	Ram int64 `json:"ram" api:"required"`
	// Number of available instances of given flavor from reservations. The value is
	// null if it is not loaded.
	ReservedCapacity int64 `json:"reserved_capacity" api:"required"`
	// Flavor resource class for mapping to hardware capacity. The value is null if the
	// flavor is not a bare metal flavor.
	ResourceClass string `json:"resource_class" api:"required"`
	// Virtual CPU count. For bare metal flavors, it's a physical CPU count
	Vcpus int64 `json:"vcpus" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Architecture        respjson.Field
		Capacity            respjson.Field
		CurrencyCode        respjson.Field
		Disabled            respjson.Field
		FlavorID            respjson.Field
		FlavorName          respjson.Field
		HardwareDescription respjson.Field
		OsType              respjson.Field
		PricePerHour        respjson.Field
		PricePerMonth       respjson.Field
		PriceStatus         respjson.Field
		Ram                 respjson.Field
		ReservedCapacity    respjson.Field
		ResourceClass       respjson.Field
		Vcpus               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r K8SFlavor) RawJSON() string { return r.JSON.raw }
func (r *K8SFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price status for the UI. The value is null if the price is not available.
type K8SFlavorPriceStatus string

const (
	K8SFlavorPriceStatusError K8SFlavorPriceStatus = "error"
	K8SFlavorPriceStatusHide  K8SFlavorPriceStatus = "hide"
	K8SFlavorPriceStatusShow  K8SFlavorPriceStatus = "show"
)

type K8SFlavorListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Set to true to exclude GPU flavors. Default is false.
	ExcludeGPU param.Opt[bool] `query:"exclude_gpu,omitzero" json:"-"`
	// Set to true to include flavor capacity. Default is False.
	IncludeCapacity param.Opt[bool] `query:"include_capacity,omitzero" json:"-"`
	// Set to true to include flavor prices. Default is False.
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [K8SFlavorListParams]'s query parameters as `url.Values`.
func (r K8SFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
