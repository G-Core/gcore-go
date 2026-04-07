// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
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
func (r *ShieldService) List(ctx context.Context, query ShieldListParams, opts ...option.RequestOption) (res *ShieldListResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/shieldingpop_v2"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// ShieldListResponseUnion contains all possible properties and values from
// [[]ShieldListResponsePlainListItem], [ShieldListResponsePaginatedList].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPlainList]
type ShieldListResponseUnion struct {
	// This field will be present if the value is a [[]ShieldListResponsePlainListItem]
	// instead of an object.
	OfPlainList []ShieldListResponsePlainListItem `json:",inline"`
	// This field is from variant [ShieldListResponsePaginatedList].
	Count int64 `json:"count"`
	// This field is from variant [ShieldListResponsePaginatedList].
	Next string `json:"next"`
	// This field is from variant [ShieldListResponsePaginatedList].
	Previous string `json:"previous"`
	// This field is from variant [ShieldListResponsePaginatedList].
	Results []ShieldListResponsePaginatedListResult `json:"results"`
	JSON    struct {
		OfPlainList respjson.Field
		Count       respjson.Field
		Next        respjson.Field
		Previous    respjson.Field
		Results     respjson.Field
		raw         string
	} `json:"-"`
}

func (u ShieldListResponseUnion) AsPlainList() (v []ShieldListResponsePlainListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShieldListResponseUnion) AsPaginatedList() (v ShieldListResponsePaginatedList) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShieldListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ShieldListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShieldListResponsePlainListItem struct {
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
func (r ShieldListResponsePlainListItem) RawJSON() string { return r.JSON.raw }
func (r *ShieldListResponsePlainListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShieldListResponsePaginatedList struct {
	// Total number of items.
	Count int64 `json:"count" api:"required"`
	// URL to the next page of results. Null if current page is the last one.
	Next string `json:"next" api:"required"`
	// URL to the previous page of results. Null if current page is the first one.
	Previous string                                  `json:"previous" api:"required"`
	Results  []ShieldListResponsePaginatedListResult `json:"results" api:"required"`
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
func (r ShieldListResponsePaginatedList) RawJSON() string { return r.JSON.raw }
func (r *ShieldListResponsePaginatedList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShieldListResponsePaginatedListResult struct {
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
func (r ShieldListResponsePaginatedListResult) RawJSON() string { return r.JSON.raw }
func (r *ShieldListResponsePaginatedListResult) UnmarshalJSON(data []byte) error {
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
