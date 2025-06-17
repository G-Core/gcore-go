// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// ClientService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientService] method instead.
type ClientService struct {
	Options []option.RequestOption
}

// NewClientService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewClientService(opts ...option.RequestOption) (r ClientService) {
	r = ClientService{}
	r.Options = opts
	return
}

// Get information about WAAP service for the client
func (r *ClientService) Me(ctx context.Context, opts ...option.RequestOption) (res *ClientMeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/clients/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Represents the WAAP service information for a client
type ClientMeResponse struct {
	// The client ID
	ID int64 `json:"id,required"`
	// List of enabled features
	Features []string `json:"features,required"`
	// Quotas for the client
	Quotas map[string]ClientMeResponseQuota `json:"quotas,required"`
	// Information about the WAAP service status
	Service ClientMeResponseService `json:"service,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Features    respjson.Field
		Quotas      respjson.Field
		Service     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientMeResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientMeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientMeResponseQuota struct {
	// The maximum allowed number of this resource
	Allowed int64 `json:"allowed,required"`
	// The current number of this resource
	Current int64 `json:"current,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allowed     respjson.Field
		Current     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientMeResponseQuota) RawJSON() string { return r.JSON.raw }
func (r *ClientMeResponseQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the WAAP service status
type ClientMeResponseService struct {
	// Whether the service is enabled
	Enabled bool `json:"enabled,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientMeResponseService) RawJSON() string { return r.JSON.raw }
func (r *ClientMeResponseService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
