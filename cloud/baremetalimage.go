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

// Retrieve the available images list for bare metal servers. Returned entities may
// or may not be owned by the project
func (r *BaremetalImageService) List(ctx context.Context, params BaremetalImageListParams, opts ...option.RequestOption) (res *ImageList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/bmimages/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type BaremetalImageListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Show price
	IncludePrices param.Opt[bool] `query:"include_prices,omitzero" json:"-"`
	// Any value to show private images
	Private param.Opt[string] `query:"private,omitzero" json:"-"`
	// Filter by tag key-value pairs. Must be a valid JSON string. 'curl -G
	// --data-urlencode 'tag_key_value={"key": "value"}' --url
	// 'http://localhost:1111/v1/images/1/1'"
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// Filter by tag keys.
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	// Image visibility. Globally visible images are public
	//
	// Any of "private", "public", "shared".
	Visibility BaremetalImageListParamsVisibility `query:"visibility,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f BaremetalImageListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

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
