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

// BaremetalImageService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBaremetalImageService] method instead.
type BaremetalImageService struct {
	Options []option.RequestOption
}

// NewBaremetalImageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBaremetalImageService(opts ...option.RequestOption) (r BaremetalImageService) {
	r = BaremetalImageService{}
	r.Options = opts
	return
}

// Retrieve a list of available images for bare metal servers. The list can be
// filtered by visibility, tags, and other parameters. Returned entities may or may
// not be owned by the project.
func (r *BaremetalImageService) List(ctx context.Context, params BaremetalImageListParams, opts ...option.RequestOption) (res *ImageList, err error) {
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
	path := fmt.Sprintf("cloud/v1/bmimages/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type BaremetalImageListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero" api:"required" json:"-"`
	// Show price
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	// Any value to show private images
	Private param.Opt[string] `query:"private,omitzero" json:"-"`
	// Filter by tag key-value pairs. Must be a valid JSON string.
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// Filter by tag keys.
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	// Image visibility. Globally visible images are public
	//
	// Any of "private", "public", "shared".
	Visibility BaremetalImageListParamsVisibility `query:"visibility,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BaremetalImageListParams]'s query parameters as
// `url.Values`.
func (r BaremetalImageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Image visibility. Globally visible images are public
type BaremetalImageListParamsVisibility string

const (
	BaremetalImageListParamsVisibilityPrivate BaremetalImageListParamsVisibility = "private"
	BaremetalImageListParamsVisibilityPublic  BaremetalImageListParamsVisibility = "public"
	BaremetalImageListParamsVisibilityShared  BaremetalImageListParamsVisibility = "shared"
)
