// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
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

// InsightService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightService] method instead.
type InsightService struct {
	Options []option.RequestOption
}

// NewInsightService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInsightService(opts ...option.RequestOption) (r InsightService) {
	r = InsightService{}
	r.Options = opts
	return
}

// Insight types are generalized categories that encompass various specific
// occurrences of the same kind.
func (r *InsightService) ListTypes(ctx context.Context, query InsightListTypesParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapInsightType], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "waap/v1/security-insights/types"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Insight types are generalized categories that encompass various specific
// occurrences of the same kind.
func (r *InsightService) ListTypesAutoPaging(ctx context.Context, query InsightListTypesParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapInsightType] {
	return pagination.NewOffsetPageAutoPager(r.ListTypes(ctx, query, opts...))
}

type WaapInsightType struct {
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
	Slug string `json:"slug,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description               respjson.Field
		InsightFrequency          respjson.Field
		InsightGroupingDimensions respjson.Field
		InsightTemplate           respjson.Field
		Labels                    respjson.Field
		Name                      respjson.Field
		RecommendationTemplate    respjson.Field
		Slug                      respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapInsightType) RawJSON() string { return r.JSON.raw }
func (r *WaapInsightType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListTypesParams struct {
	// Filter by the frequency of the insight type
	InsightFrequency param.Opt[int64] `query:"insight_frequency,omitzero" json:"-"`
	// Filter by the name of the insight type
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Filter by the slug of the insight type
	Slug param.Opt[string] `query:"slug,omitzero" json:"-"`
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Sort the response by given field.
	//
	// Any of "name", "-name", "slug", "-slug", "insight_frequency",
	// "-insight_frequency".
	Ordering InsightListTypesParamsOrdering `query:"ordering,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightListTypesParams]'s query parameters as `url.Values`.
func (r InsightListTypesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort the response by given field.
type InsightListTypesParamsOrdering string

const (
	InsightListTypesParamsOrderingName                  InsightListTypesParamsOrdering = "name"
	InsightListTypesParamsOrderingMinusName             InsightListTypesParamsOrdering = "-name"
	InsightListTypesParamsOrderingSlug                  InsightListTypesParamsOrdering = "slug"
	InsightListTypesParamsOrderingMinusSlug             InsightListTypesParamsOrdering = "-slug"
	InsightListTypesParamsOrderingInsightFrequency      InsightListTypesParamsOrdering = "insight_frequency"
	InsightListTypesParamsOrderingMinusInsightFrequency InsightListTypesParamsOrdering = "-insight_frequency"
)
