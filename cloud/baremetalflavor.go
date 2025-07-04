// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// BaremetalFlavorService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBaremetalFlavorService] method instead.
type BaremetalFlavorService struct {
	Options []option.RequestOption
}

// NewBaremetalFlavorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBaremetalFlavorService(opts ...option.RequestOption) (r BaremetalFlavorService) {
	r = BaremetalFlavorService{}
	r.Options = opts
	return
}

// List all available bare metal flavors in the specified project and region. When
// `include_prices` is specified, the list includes pricing information. A client
// in trial mode gets all price values as 0. If you get Pricing Error contact the
// support.
func (r *BaremetalFlavorService) List(ctx context.Context, params BaremetalFlavorListParams, opts ...option.RequestOption) (res *BaremetalFlavorList, err error) {
	opts = append(r.Options[:], opts...)
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
	path := fmt.Sprintf("cloud/v1/bmflavors/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type BaremetalFlavorListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Flag for filtering disabled flavors in the region. Defaults to true
	Disabled param.Opt[bool] `query:"disabled,omitzero" json:"-"`
	// Set to true to exclude flavors dedicated to linux images. Default False
	ExcludeLinux param.Opt[bool] `query:"exclude_linux,omitzero" json:"-"`
	// Set to true to exclude flavors dedicated to windows images. Default False
	ExcludeWindows param.Opt[bool] `query:"exclude_windows,omitzero" json:"-"`
	// Set to true if the response should include flavor capacity
	IncludeCapacity param.Opt[bool] `query:"include_capacity,omitzero" json:"-"`
	// Set to true if the response should include flavor prices
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	// Optional. Set to true if flavor listing should include count of reserved
	// resources in stock.
	IncludeReservationStock param.Opt[bool] `query:"include_reservation_stock,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalFlavorListParams]'s query parameters as
// `url.Values`.
func (r BaremetalFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
