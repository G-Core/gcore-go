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

// ShieldService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewShieldService] method instead.
type ShieldService struct {
	Options []option.RequestOption
}

// NewShieldService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewShieldService(opts ...option.RequestOption) (r ShieldService) {
	r = ShieldService{}
	r.Options = opts
	return
}

// Get information about all origin shielding locations available in the account.
func (r *ShieldService) List(ctx context.Context, query ShieldListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[ShieldListResponse], err error) {
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
func (r *ShieldService) ListAutoPaging(ctx context.Context, query ShieldListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[ShieldListResponse] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

type ShieldListResponse struct {
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
func (r ShieldListResponse) RawJSON() string { return r.JSON.raw }
func (r *ShieldListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShieldListParams struct {
	// Maximum number of items to return in the response. Cannot exceed 1000.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip from the beginning of the list.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ShieldListParams]'s query parameters as `url.Values`.
func (r ShieldListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
