// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
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
func (r *K8SFlavorService) List(ctx context.Context, params K8SFlavorListParams, opts ...option.RequestOption) (res *BaremetalFlavorList, err error) {
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
	path := fmt.Sprintf("cloud/v1/k8s/%v/%v/flavors", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type K8SFlavorListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Set to true to exclude GPU flavors. Default is false.
	ExcludeGPU param.Opt[bool] `query:"exclude_gpu,omitzero" json:"-"`
	// Set to true to include flavor capacity. Default is False.
	IncludeCapacity param.Opt[bool] `query:"include_capacity,omitzero" json:"-"`
	// Set to true to include flavor prices. Default is False.
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [K8SFlavorListParams]'s query parameters as `url.Values`.
func (r K8SFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
