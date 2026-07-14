// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

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

// CDN presets are predefined sets of CDN resource or rule settings that can be
// applied to an object in a single request, letting you configure caching,
// delivery, and security options consistently.
//
// PresetService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPresetService] method instead.
type PresetService struct {
	Options []option.RequestOption
	// Applied presets represent the association between a preset and a CDN resource or
	// rule. Use them to apply a preset to an object, list the objects a preset is
	// applied to, unapply it, and inspect which object fields a preset manages.
	Applied PresetAppliedService
}

// NewPresetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPresetService(opts ...option.RequestOption) (r PresetService) {
	r = PresetService{}
	r.Options = opts
	r.Applied = NewPresetAppliedService(opts...)
	return
}

// Get the list of presets available to your account.
//
// A preset is a predefined set of CDN resource or rule settings that can be
// applied to an object in one request.
func (r *PresetService) List(ctx context.Context, query PresetListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[PresetDetail], err error) {
	// Limit set with a default so List always paginates.
	if !query.Limit.Valid() {
		query.Limit = param.NewOpt[int64](1000)
	}
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cdn/presets"
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

// Get the list of presets available to your account.
//
// A preset is a predefined set of CDN resource or rule settings that can be
// applied to an object in one request.
func (r *PresetService) ListAutoPaging(ctx context.Context, query PresetListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[PresetDetail] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Get information about a preset.
func (r *PresetService) Get(ctx context.Context, presetID int64, opts ...option.RequestOption) (res *PresetDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("cdn/presets/%v", presetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PresetDetail struct {
	// Preset ID.
	ID int64 `json:"id"`
	// Preset name.
	Name string `json:"name"`
	// Type of object the preset can be applied to.
	//
	// Possible values:
	//
	// - **CDNResource** - Preset is applied to a CDN resource.
	// - **Rule** - Preset is applied to a rule.
	//
	// Any of "CDNResource", "Rule".
	ObjectType PresetDetailObjectType `json:"object_type"`
	// CDN resource or rule settings that the preset applies to the target object,
	// including the **options** object.
	//
	// The available keys match the writable fields of the object type the preset
	// targets (`object_type`). Options included in the preset cannot be edited on the
	// object while the preset is applied.
	PresetSettings map[string]any `json:"preset_settings"`
	// Service the preset belongs to.
	//
	// Any of "cdn".
	Service PresetDetailService `json:"service"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Name           respjson.Field
		ObjectType     respjson.Field
		PresetSettings respjson.Field
		Service        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PresetDetail) RawJSON() string { return r.JSON.raw }
func (r *PresetDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of object the preset can be applied to.
//
// Possible values:
//
// - **CDNResource** - Preset is applied to a CDN resource.
// - **Rule** - Preset is applied to a rule.
type PresetDetailObjectType string

const (
	PresetDetailObjectTypeCDNResource PresetDetailObjectType = "CDNResource"
	PresetDetailObjectTypeRule        PresetDetailObjectType = "Rule"
)

// Service the preset belongs to.
type PresetDetailService string

const (
	PresetDetailServiceCDN PresetDetailService = "cdn"
)

type PresetListParams struct {
	// Maximum number of items to return in the response. Cannot exceed 1000.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip from the beginning of the list.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PresetListParams]'s query parameters as `url.Values`.
func (r PresetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
