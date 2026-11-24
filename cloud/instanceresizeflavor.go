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

// InstanceResizeFlavorService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceResizeFlavorService] method instead.
type InstanceResizeFlavorService struct {
	Options []option.RequestOption
}

// NewInstanceResizeFlavorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInstanceResizeFlavorService(opts ...option.RequestOption) (r InstanceResizeFlavorService) {
	r = InstanceResizeFlavorService{}
	r.Options = opts
	return
}

// Returns the flavors the instance can be resized to: the flavors in its region
// that keep its flavor family and network profile, without its current flavor. An
// empty list means there is nothing to resize to. A 400 means the instance cannot
// be resized at all.
func (r *InstanceResizeFlavorService) List(ctx context.Context, instanceID string, params InstanceResizeFlavorListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[InstanceFlavor], err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/instances/%s/flavors", params.ProjectID.Value, instanceID)
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

// Returns the flavors the instance can be resized to: the flavors in its region
// that keep its flavor family and network profile, without its current flavor. An
// empty list means there is nothing to resize to. A 400 means the instance cannot
// be resized at all.
func (r *InstanceResizeFlavorService) ListAutoPaging(ctx context.Context, instanceID string, params InstanceResizeFlavorListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[InstanceFlavor] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, instanceID, params, opts...))
}

type InstanceResizeFlavorListParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Max items to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Items to skip.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Sort field and direction.
	//
	// Any of "name.asc", "name.desc", "ram_mib.asc", "ram_mib.desc", "vcpus.asc",
	// "vcpus.desc".
	OrderBy InstanceResizeFlavorListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InstanceResizeFlavorListParams]'s query parameters as
// `url.Values`.
func (r InstanceResizeFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort field and direction.
type InstanceResizeFlavorListParamsOrderBy string

const (
	InstanceResizeFlavorListParamsOrderByNameAsc    InstanceResizeFlavorListParamsOrderBy = "name.asc"
	InstanceResizeFlavorListParamsOrderByNameDesc   InstanceResizeFlavorListParamsOrderBy = "name.desc"
	InstanceResizeFlavorListParamsOrderByRamMibAsc  InstanceResizeFlavorListParamsOrderBy = "ram_mib.asc"
	InstanceResizeFlavorListParamsOrderByRamMibDesc InstanceResizeFlavorListParamsOrderBy = "ram_mib.desc"
	InstanceResizeFlavorListParamsOrderByVcpusAsc   InstanceResizeFlavorListParamsOrderBy = "vcpus.asc"
	InstanceResizeFlavorListParamsOrderByVcpusDesc  InstanceResizeFlavorListParamsOrderBy = "vcpus.desc"
)
