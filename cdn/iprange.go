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
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// IPRangeService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPRangeService] method instead.
type IPRangeService struct {
	Options []option.RequestOption
}

// NewIPRangeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPRangeService(opts ...option.RequestOption) (r IPRangeService) {
	r = IPRangeService{}
	r.Options = opts
	return
}

// Get all CDN networks that can be used to pull content from your origin.
//
// This list is updated periodically. If you want to use network from this list to
// configure IP ACL on your origin, you need to independently monitor its
// relevance. We recommend using a script for automatically update IP ACL.
//
// This request does not require authorization.
func (r *IPRangeService) List(ctx context.Context, params IPRangeListParams, opts ...option.RequestOption) (res *PublicNetworkList, err error) {
	if !param.IsOmitted(params.Accept) {
		opts = append(opts, option.WithHeader("Accept", fmt.Sprintf("%v", params.Accept)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "cdn/public-net-list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type PublicNetworkList struct {
	// List of IPv4 networks.
	Addresses []string `json:"addresses"`
	// List of IPv6 networks.
	AddressesV6 []string `json:"addresses_v6"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Addresses   respjson.Field
		AddressesV6 respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicNetworkList) RawJSON() string { return r.JSON.raw }
func (r *PublicNetworkList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPRangeListParams struct {
	// Optional format override. When set, this takes precedence over the `Accept`
	// header.
	//
	// Any of "json", "plain".
	Format IPRangeListParamsFormat `query:"format,omitzero" json:"-"`
	// Any of "application/json", "text/plain".
	Accept IPRangeListParamsAccept `header:"Accept,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IPRangeListParams]'s query parameters as `url.Values`.
func (r IPRangeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Optional format override. When set, this takes precedence over the `Accept`
// header.
type IPRangeListParamsFormat string

const (
	IPRangeListParamsFormatJson  IPRangeListParamsFormat = "json"
	IPRangeListParamsFormatPlain IPRangeListParamsFormat = "plain"
)

type IPRangeListParamsAccept string

const (
	IPRangeListParamsAcceptApplicationJson IPRangeListParamsAccept = "application/json"
	IPRangeListParamsAcceptTextPlain       IPRangeListParamsAccept = "text/plain"
)
