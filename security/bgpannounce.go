// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security

import (
	"context"
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

// BgpAnnounceService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBgpAnnounceService] method instead.
type BgpAnnounceService struct {
	Options []option.RequestOption
}

// NewBgpAnnounceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBgpAnnounceService(opts ...option.RequestOption) (r BgpAnnounceService) {
	r = BgpAnnounceService{}
	r.Options = opts
	return
}

// Get BGP announces filtered by parameters. Shows announces in active profiles,
// meaning that to get a non-empty response, the client must have at least one
// active profile.
func (r *BgpAnnounceService) List(ctx context.Context, query BgpAnnounceListParams, opts ...option.RequestOption) (res *[]ClientAnnounce, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "security/sifter/v2/protected_addresses/announces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Change BGP announces (it can be enabled or disabled, but not created or
// updated). Can be applied to already existing announces in active profiles,
// meaning that the client must have at least one active profile.
func (r *BgpAnnounceService) Toggle(ctx context.Context, params BgpAnnounceToggleParams, opts ...option.RequestOption) (res *BgpAnnounceToggleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "security/sifter/v2/protected_addresses/announces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ClientAnnounce struct {
	Announced    []string `json:"announced" api:"required" format:"ipv4network"`
	ClientID     int64    `json:"client_id" api:"required"`
	NotAnnounced []string `json:"not_announced" api:"required" format:"ipv4network"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Announced    respjson.Field
		ClientID     respjson.Field
		NotAnnounced respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientAnnounce) RawJSON() string { return r.JSON.raw }
func (r *ClientAnnounce) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BgpAnnounceToggleResponse = any

type BgpAnnounceListParams struct {
	Announced param.Opt[bool]   `query:"announced,omitzero" json:"-"`
	Site      param.Opt[string] `query:"site,omitzero" json:"-"`
	// Any of "STATIC", "DYNAMIC".
	Origin BgpAnnounceListParamsOrigin `query:"origin,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BgpAnnounceListParams]'s query parameters as `url.Values`.
func (r BgpAnnounceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type BgpAnnounceListParamsOrigin string

const (
	BgpAnnounceListParamsOriginStatic  BgpAnnounceListParamsOrigin = "STATIC"
	BgpAnnounceListParamsOriginDynamic BgpAnnounceListParamsOrigin = "DYNAMIC"
)

type BgpAnnounceToggleParams struct {
	Announce string           `json:"announce" api:"required" format:"ipv4network"`
	Enabled  bool             `json:"enabled" api:"required"`
	ClientID param.Opt[int64] `query:"client_id,omitzero" json:"-"`
	paramObj
}

func (r BgpAnnounceToggleParams) MarshalJSON() (data []byte, err error) {
	type shadow BgpAnnounceToggleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BgpAnnounceToggleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BgpAnnounceToggleParams]'s query parameters as
// `url.Values`.
func (r BgpAnnounceToggleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
