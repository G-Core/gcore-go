// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
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
func (r *IPRangeService) List(ctx context.Context, opts ...option.RequestOption) (res *PublicNetworkList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/public-net-list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *IPRangeService) ListIPs(ctx context.Context, opts ...option.RequestOption) (res *PublicIPList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/public-ip-list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
