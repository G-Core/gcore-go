// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
)

// GPUBaremetalClusterFlavorService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGPUBaremetalClusterFlavorService] method instead.
type GPUBaremetalClusterFlavorService struct {
	Options []option.RequestOption
}

// NewGPUBaremetalClusterFlavorService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGPUBaremetalClusterFlavorService(opts ...option.RequestOption) (r GPUBaremetalClusterFlavorService) {
	r = GPUBaremetalClusterFlavorService{}
	r.Options = opts
	return
}

// List bare metal GPU flavors
func (r *GPUBaremetalClusterFlavorService) List(ctx context.Context, params GPUBaremetalClusterFlavorListParams, opts ...option.RequestOption) (res *GPUBaremetalFlavorList, err error) {
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
	path := fmt.Sprintf("cloud/v3/gpu/baremetal/%v/%v/flavors", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type GPUBaremetalClusterFlavorListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Set to `true` to remove the disabled flavors from the response.
	HideDisabled param.Opt[bool] `query:"hide_disabled,omitzero" json:"-"`
	// Set to `true` if the response should include flavor prices.
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GPUBaremetalClusterFlavorListParams]'s query parameters as
// `url.Values`.
func (r GPUBaremetalClusterFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
