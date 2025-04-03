// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// WaapV1SecurityInsightService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1SecurityInsightService] method instead.
type WaapV1SecurityInsightService struct {
	Options []option.RequestOption
}

// NewWaapV1SecurityInsightService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWaapV1SecurityInsightService(opts ...option.RequestOption) (r *WaapV1SecurityInsightService) {
	r = &WaapV1SecurityInsightService{}
	r.Options = opts
	return
}

// Insight types are generalized categories that encompass various specific
// occurrences of the same kind.
func (r *WaapV1SecurityInsightService) ListTypes(ctx context.Context, query WaapV1SecurityInsightListTypesParams, opts ...option.RequestOption) (res *WaapV1SecurityInsightListTypesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/security-insights/types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WaapV1SecurityInsightListTypesResponse struct {
	// Number of items contain in the response
	Count int64 `json:"count,required"`
	// Number of items requested in the response
	Limit int64 `json:"limit,required"`
	// Items response offset used
	Offset int64 `json:"offset,required"`
	// List of items returned in the response following given criteria
	Results []WaapV1SecurityInsightListTypesResponseResult `json:"results,required"`
	JSON    waapV1SecurityInsightListTypesResponseJSON     `json:"-"`
}

// waapV1SecurityInsightListTypesResponseJSON contains the JSON metadata for the
// struct [WaapV1SecurityInsightListTypesResponse]
type waapV1SecurityInsightListTypesResponseJSON struct {
	Count       apijson.Field
	Limit       apijson.Field
	Offset      apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1SecurityInsightListTypesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1SecurityInsightListTypesResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1SecurityInsightListTypesResponseResult struct {
	// The description of the insight type
	Description string `json:"description,required"`
	// The frequency of the insight type
	InsightFrequency int64 `json:"insight_frequency,required"`
	// The grouping dimensions of the insight type
	InsightGroupingDimensions []string `json:"insight_grouping_dimensions,required"`
	// The insight template
	InsightTemplate string `json:"insight_template,required"`
	// The labels of the insight type
	Labels []string `json:"labels,required"`
	// The name of the insight type
	Name string `json:"name,required"`
	// The recommendation template
	RecommendationTemplate string `json:"recommendation_template,required"`
	// The slug of the insight type
	Slug string                                           `json:"slug,required"`
	JSON waapV1SecurityInsightListTypesResponseResultJSON `json:"-"`
}

// waapV1SecurityInsightListTypesResponseResultJSON contains the JSON metadata for
// the struct [WaapV1SecurityInsightListTypesResponseResult]
type waapV1SecurityInsightListTypesResponseResultJSON struct {
	Description               apijson.Field
	InsightFrequency          apijson.Field
	InsightGroupingDimensions apijson.Field
	InsightTemplate           apijson.Field
	Labels                    apijson.Field
	Name                      apijson.Field
	RecommendationTemplate    apijson.Field
	Slug                      apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *WaapV1SecurityInsightListTypesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1SecurityInsightListTypesResponseResultJSON) RawJSON() string {
	return r.raw
}

type WaapV1SecurityInsightListTypesParams struct {
	// Filter by the frequency of the insight type
	InsightFrequency param.Field[int64] `query:"insight_frequency"`
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Filter by the name of the insight type
	Name param.Field[string] `query:"name"`
	// Number of items to skip
	Offset param.Field[int64] `query:"offset"`
	// Sort the response by given field.
	Ordering param.Field[WaapV1SecurityInsightListTypesParamsOrdering] `query:"ordering"`
	// Filter by the slug of the insight type
	Slug param.Field[string] `query:"slug"`
}

// URLQuery serializes [WaapV1SecurityInsightListTypesParams]'s query parameters as
// `url.Values`.
func (r WaapV1SecurityInsightListTypesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort the response by given field.
type WaapV1SecurityInsightListTypesParamsOrdering string

const (
	WaapV1SecurityInsightListTypesParamsOrderingName                  WaapV1SecurityInsightListTypesParamsOrdering = "name"
	WaapV1SecurityInsightListTypesParamsOrderingMinusName             WaapV1SecurityInsightListTypesParamsOrdering = "-name"
	WaapV1SecurityInsightListTypesParamsOrderingSlug                  WaapV1SecurityInsightListTypesParamsOrdering = "slug"
	WaapV1SecurityInsightListTypesParamsOrderingMinusSlug             WaapV1SecurityInsightListTypesParamsOrdering = "-slug"
	WaapV1SecurityInsightListTypesParamsOrderingInsightFrequency      WaapV1SecurityInsightListTypesParamsOrdering = "insight_frequency"
	WaapV1SecurityInsightListTypesParamsOrderingMinusInsightFrequency WaapV1SecurityInsightListTypesParamsOrdering = "-insight_frequency"
)

func (r WaapV1SecurityInsightListTypesParamsOrdering) IsKnown() bool {
	switch r {
	case WaapV1SecurityInsightListTypesParamsOrderingName, WaapV1SecurityInsightListTypesParamsOrderingMinusName, WaapV1SecurityInsightListTypesParamsOrderingSlug, WaapV1SecurityInsightListTypesParamsOrderingMinusSlug, WaapV1SecurityInsightListTypesParamsOrderingInsightFrequency, WaapV1SecurityInsightListTypesParamsOrderingMinusInsightFrequency:
		return true
	}
	return false
}
