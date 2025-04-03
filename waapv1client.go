// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// WaapV1ClientService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWaapV1ClientService] method instead.
type WaapV1ClientService struct {
	Options []option.RequestOption
}

// NewWaapV1ClientService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWaapV1ClientService(opts ...option.RequestOption) (r *WaapV1ClientService) {
	r = &WaapV1ClientService{}
	r.Options = opts
	return
}

// Get information about WAAP service for the client
func (r *WaapV1ClientService) GetMe(ctx context.Context, opts ...option.RequestOption) (res *WaapV1ClientGetMeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/clients/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Represents the WAAP service information for a client
type WaapV1ClientGetMeResponse struct {
	// The client ID
	ID int64 `json:"id,required,nullable"`
	// List of enabled features
	Features []string `json:"features,required"`
	// Quotas for the client
	Quotas map[string]WaapV1ClientGetMeResponseQuota `json:"quotas,required"`
	// Information about the WAAP service status
	Service WaapV1ClientGetMeResponseService `json:"service,required"`
	JSON    waapV1ClientGetMeResponseJSON    `json:"-"`
}

// waapV1ClientGetMeResponseJSON contains the JSON metadata for the struct
// [WaapV1ClientGetMeResponse]
type waapV1ClientGetMeResponseJSON struct {
	ID          apijson.Field
	Features    apijson.Field
	Quotas      apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1ClientGetMeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1ClientGetMeResponseJSON) RawJSON() string {
	return r.raw
}

type WaapV1ClientGetMeResponseQuota struct {
	// The maximum allowed number of this resource
	Allowed int64 `json:"allowed,required"`
	// The current number of this resource
	Current int64                              `json:"current,required"`
	JSON    waapV1ClientGetMeResponseQuotaJSON `json:"-"`
}

// waapV1ClientGetMeResponseQuotaJSON contains the JSON metadata for the struct
// [WaapV1ClientGetMeResponseQuota]
type waapV1ClientGetMeResponseQuotaJSON struct {
	Allowed     apijson.Field
	Current     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1ClientGetMeResponseQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1ClientGetMeResponseQuotaJSON) RawJSON() string {
	return r.raw
}

// Information about the WAAP service status
type WaapV1ClientGetMeResponseService struct {
	// Whether the service is enabled
	Enabled bool                                 `json:"enabled,required"`
	JSON    waapV1ClientGetMeResponseServiceJSON `json:"-"`
}

// waapV1ClientGetMeResponseServiceJSON contains the JSON metadata for the struct
// [WaapV1ClientGetMeResponseService]
type waapV1ClientGetMeResponseServiceJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaapV1ClientGetMeResponseService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waapV1ClientGetMeResponseServiceJSON) RawJSON() string {
	return r.raw
}
