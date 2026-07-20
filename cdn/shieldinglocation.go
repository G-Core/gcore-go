// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

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

// ShieldingLocationService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewShieldingLocationService] method instead.
type ShieldingLocationService struct {
	Options []option.RequestOption
}

// NewShieldingLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewShieldingLocationService(opts ...option.RequestOption) (r ShieldingLocationService) {
	r = ShieldingLocationService{}
	r.Options = opts
	return
}

// Get information about all origin shielding locations available in the account.
func (r *ShieldingLocationService) List(ctx context.Context, query ShieldingLocationListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ShieldingLocation], err error) {
	// CDN API only envelopes when limit>=1; default so List always paginates.
	if !query.Limit.Valid() {
		query.Limit = param.NewOpt[int64](1000)
	}
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cdn/shieldingpop_v2"
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

// Get information about all origin shielding locations available in the account.
func (r *ShieldingLocationService) ListAutoPaging(ctx context.Context, query ShieldingLocationListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ShieldingLocation] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

type ShieldingLocation struct {
	// Origin shielding location ID.
	ID int64 `json:"id"`
	// City of origin shielding location.
	City string `json:"city"`
	// Country of origin shielding location.
	Country string `json:"country"`
	// Name of origin shielding location datacenter.
	Datacenter string `json:"datacenter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		City        respjson.Field
		Country     respjson.Field
		Datacenter  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShieldingLocation) RawJSON() string { return r.JSON.raw }
func (r *ShieldingLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShieldingLocationList struct {
	// Total number of items.
	Count int64 `json:"count" api:"required"`
	// URL to the next page of results. Null if current page is the last one.
	Next string `json:"next" api:"required"`
	// URL to the previous page of results. Null if current page is the first one.
	Previous string              `json:"previous" api:"required"`
	Results  []ShieldingLocation `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Next        respjson.Field
		Previous    respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShieldingLocationList) RawJSON() string { return r.JSON.raw }
func (r *ShieldingLocationList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShieldingLocationListParams struct {
	// Maximum number of items to return in the response. Cannot exceed 1000.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip from the beginning of the list.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ShieldingLocationListParams]'s query parameters as
// `url.Values`.
func (r ShieldingLocationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
