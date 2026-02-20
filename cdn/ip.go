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

// IPService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPService] method instead.
type IPService struct {
	Options []option.RequestOption
}

// NewIPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPService(opts ...option.RequestOption) (r IPService) {
	r = IPService{}
	r.Options = opts
	return
}

// Get all IP addresses of CDN servers that can be used to pull content from your
// origin.
//
// This list is updated periodically. If you want to use IP from this list to
// configure IP ACL in your origin, you need to independently monitor its
// relevance. We recommend using a script to automatically update IP ACL.
//
// This request does not require authorization.
func (r *IPService) List(ctx context.Context, params IPListParams, opts ...option.RequestOption) (res *PublicIPList, err error) {
	if !param.IsOmitted(params.Accept) {
		opts = append(opts, option.WithHeader("Accept", fmt.Sprintf("%v", params.Accept)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "cdn/public-ip-list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type PublicIPList struct {
	// List of IPv4 addresses.
	Addresses []string `json:"addresses"`
	// List of IPv6 addresses.
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
func (r PublicIPList) RawJSON() string { return r.JSON.raw }
func (r *PublicIPList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPListParams struct {
	// Optional format override. When set, this takes precedence over the `Accept`
	// header.
	//
	// Any of "json", "plain".
	Format IPListParamsFormat `query:"format,omitzero" json:"-"`
	// Any of "application/json", "text/plain".
	Accept IPListParamsAccept `header:"Accept,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IPListParams]'s query parameters as `url.Values`.
func (r IPListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Optional format override. When set, this takes precedence over the `Accept`
// header.
type IPListParamsFormat string

const (
	IPListParamsFormatJson  IPListParamsFormat = "json"
	IPListParamsFormatPlain IPListParamsFormat = "plain"
)

type IPListParamsAccept string

const (
	IPListParamsAcceptApplicationJson IPListParamsAccept = "application/json"
	IPListParamsAcceptTextPlain       IPListParamsAccept = "text/plain"
)
