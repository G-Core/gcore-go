// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// LocationService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocationService] method instead.
type LocationService struct {
	Options []option.RequestOption
}

// NewLocationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLocationService(opts ...option.RequestOption) (r LocationService) {
	r = LocationService{}
	r.Options = opts
	return
}

// List of All locations continents/countries/regions.
func (r *LocationService) List(ctx context.Context, opts ...option.RequestOption) (res *LocationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dns/v2/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List of All locations continents.
func (r *LocationService) ListContinents(ctx context.Context, opts ...option.RequestOption) (res *LocationListContinentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dns/v2/locations/continents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List of All locations countries.
func (r *LocationService) ListCountries(ctx context.Context, opts ...option.RequestOption) (res *LocationListCountriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dns/v2/locations/countries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List of All locations regions.
func (r *LocationService) ListRegions(ctx context.Context, opts ...option.RequestOption) (res *LocationListRegionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dns/v2/locations/regions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DNSLocationTranslations struct {
	Names map[string]string `json:"names"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Names       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSLocationTranslations) RawJSON() string { return r.JSON.raw }
func (r *DNSLocationTranslations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationListResponse struct {
	Continents map[string]DNSLocationTranslations `json:"continents"`
	Countries  map[string]DNSLocationTranslations `json:"countries"`
	Regions    map[string]DNSLocationTranslations `json:"regions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Continents  respjson.Field
		Countries   respjson.Field
		Regions     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationListResponse) RawJSON() string { return r.JSON.raw }
func (r *LocationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationListContinentsResponse map[string]DNSLocationTranslations

type LocationListCountriesResponse map[string]DNSLocationTranslations

type LocationListRegionsResponse map[string]DNSLocationTranslations
