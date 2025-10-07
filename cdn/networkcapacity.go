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

// NetworkCapacityService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkCapacityService] method instead.
type NetworkCapacityService struct {
	Options []option.RequestOption
}

// NewNetworkCapacityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkCapacityService(opts ...option.RequestOption) (r NetworkCapacityService) {
	r = NetworkCapacityService{}
	r.Options = opts
	return
}

// Get network capacity per country.
func (r *NetworkCapacityService) List(ctx context.Context, opts ...option.RequestOption) (res *NetworkCapacity, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/advanced/v1/capacity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type NetworkCapacity []NetworkCapacityItem

type NetworkCapacityItem struct {
	// Network capacity in Gbit/s.
	Capacity float64 `json:"capacity"`
	// Country name.
	Country string `json:"country"`
	// ISO country code.
	CountryCode string `json:"country_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capacity    respjson.Field
		Country     respjson.Field
		CountryCode respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkCapacityItem) RawJSON() string { return r.JSON.raw }
func (r *NetworkCapacityItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
