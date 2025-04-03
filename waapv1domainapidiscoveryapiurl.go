// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// WaapV1DomainAPIDiscoveryAPIURLService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1DomainAPIDiscoveryAPIURLService] method instead.
type WaapV1DomainAPIDiscoveryAPIURLService struct {
	Options []option.RequestOption
}

// NewWaapV1DomainAPIDiscoveryAPIURLService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWaapV1DomainAPIDiscoveryAPIURLService(opts ...option.RequestOption) (r *WaapV1DomainAPIDiscoveryAPIURLService) {
	r = &WaapV1DomainAPIDiscoveryAPIURLService{}
	r.Options = opts
	return
}

// Update the API URLs for a domain. If your domain has a common base URL for all
// API paths, you can set it here.
func (r *WaapV1DomainAPIDiscoveryAPIURLService) Update(ctx context.Context, domainID int64, body WaapV1DomainAPIDiscoveryAPIURLUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/api-urls", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Retrieve the API URLs for a domain
func (r *WaapV1DomainAPIDiscoveryAPIURLService) List(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *WaapV1DomainAPIDiscoveryApiurlListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-discovery/api-urls", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Request model for updating the API URLs
type WaapV1DomainAPIDiscoveryApiurlListResponse struct {
	// The list of URLs
	APIURLs []string                                       `json:"api_urls,required"`
	JSON    waapV1DomainAPIDiscoveryApiurlListResponseJSON `json:"-"`
}

// waapV1DomainAPIDiscoveryApiurlListResponseJSON contains the JSON metadata for
// the struct [WaapV1DomainAPIDiscoveryApiurlListResponse]
type waapV1DomainAPIDiscoveryApiurlListResponseJSON struct {
	APIURLs     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1DomainAPIDiscoveryApiurlListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1DomainAPIDiscoveryApiurlListResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1DomainAPIDiscoveryAPIURLUpdateParams struct {
	// The list of URLs
	APIURLs param.Field[[]string] `json:"api_urls,required"`
}

func (r WaapV1DomainAPIDiscoveryAPIURLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
