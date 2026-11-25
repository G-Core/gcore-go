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
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
)

// PricingBaremetalFlavorService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPricingBaremetalFlavorService] method instead.
type PricingBaremetalFlavorService struct {
	Options []option.RequestOption
}

// NewPricingBaremetalFlavorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPricingBaremetalFlavorService(opts ...option.RequestOption) (r PricingBaremetalFlavorService) {
	r = PricingBaremetalFlavorService{}
	r.Options = opts
	return
}

// Get the prices of the bare metal flavors that the client can use, in all
// regions.
func (r *PricingBaremetalFlavorService) List(ctx context.Context, params PricingBaremetalFlavorListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[FlavorPrice], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/pricing/baremetal/flavors", params.ProjectID.Value)
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

// Get the prices of the bare metal flavors that the client can use, in all
// regions.
func (r *PricingBaremetalFlavorService) ListAutoPaging(ctx context.Context, params PricingBaremetalFlavorListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[FlavorPrice] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

type PricingBaremetalFlavorListParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Max items to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Items to skip.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter by exact flavor names.
	FlavorNames []string `query:"flavor_names,omitzero" json:"-"`
	// Sort field and direction.
	//
	// Any of "flavor_name.asc", "flavor_name.desc", "region_id.asc", "region_id.desc".
	OrderBy PricingBaremetalFlavorListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// Filter by region IDs. Omit to get all regions that the client can see.
	RegionIDs []int64 `query:"region_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PricingBaremetalFlavorListParams]'s query parameters as
// `url.Values`.
func (r PricingBaremetalFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort field and direction.
type PricingBaremetalFlavorListParamsOrderBy string

const (
	PricingBaremetalFlavorListParamsOrderByFlavorNameAsc  PricingBaremetalFlavorListParamsOrderBy = "flavor_name.asc"
	PricingBaremetalFlavorListParamsOrderByFlavorNameDesc PricingBaremetalFlavorListParamsOrderBy = "flavor_name.desc"
	PricingBaremetalFlavorListParamsOrderByRegionIDAsc    PricingBaremetalFlavorListParamsOrderBy = "region_id.asc"
	PricingBaremetalFlavorListParamsOrderByRegionIDDesc   PricingBaremetalFlavorListParamsOrderBy = "region_id.desc"
)
