// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
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

// FilterTemplateService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFilterTemplateService] method instead.
type FilterTemplateService struct {
	Options []option.RequestOption
}

// NewFilterTemplateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFilterTemplateService(opts ...option.RequestOption) (r FilterTemplateService) {
	r = FilterTemplateService{}
	r.Options = opts
	return
}

// Create a filter template for the authenticated user.
func (r *FilterTemplateService) New(ctx context.Context, body FilterTemplateNewParams, opts ...option.RequestOption) (res *WaapFilterTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "waap/v1/filter-templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update or rename a filter template owned by the authenticated user.
func (r *FilterTemplateService) Update(ctx context.Context, templateID int64, body FilterTemplateUpdateParams, opts ...option.RequestOption) (res *WaapFilterTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/filter-templates/%v", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List filter templates owned by the authenticated user.
func (r *FilterTemplateService) List(ctx context.Context, query FilterTemplateListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[WaapFilterTemplate], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "waap/v1/filter-templates"
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

// List filter templates owned by the authenticated user.
func (r *FilterTemplateService) ListAutoPaging(ctx context.Context, query FilterTemplateListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[WaapFilterTemplate] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a filter template owned by the authenticated user.
func (r *FilterTemplateService) Delete(ctx context.Context, templateID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("waap/v1/filter-templates/%v", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve a single filter template, scoped to the authenticated user.
func (r *FilterTemplateService) Get(ctx context.Context, templateID int64, opts ...option.RequestOption) (res *WaapFilterTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("waap/v1/filter-templates/%v", templateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Response body for a single filter template.
type WaapFilterTemplate struct {
	// The ID of the filter template
	ID int64 `json:"id" api:"required"`
	// Arbitrary JSON object storing the frontend filter state
	FilterConfig map[string]any `json:"filter_config" api:"required"`
	// Human-readable template name
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		FilterConfig respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WaapFilterTemplate) RawJSON() string { return r.JSON.raw }
func (r *WaapFilterTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilterTemplateNewParams struct {
	// Arbitrary JSON object storing the frontend filter state
	FilterConfig map[string]any `json:"filter_config,omitzero" api:"required"`
	// Human-readable template name
	Name string `json:"name" api:"required"`
	paramObj
}

func (r FilterTemplateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FilterTemplateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilterTemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilterTemplateUpdateParams struct {
	// Human-readable template name
	Name param.Opt[string] `json:"name,omitzero"`
	// Arbitrary JSON object storing the frontend filter state
	FilterConfig map[string]any `json:"filter_config,omitzero"`
	paramObj
}

func (r FilterTemplateUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FilterTemplateUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilterTemplateUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilterTemplateListParams struct {
	// Number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter template names by case-insensitive substring match
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FilterTemplateListParams]'s query parameters as
// `url.Values`.
func (r FilterTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
