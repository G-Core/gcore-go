// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// ClientConfigService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientConfigService] method instead.
type ClientConfigService struct {
	Options []option.RequestOption
}

// NewClientConfigService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClientConfigService(opts ...option.RequestOption) (r ClientConfigService) {
	r = ClientConfigService{}
	r.Options = opts
	return
}

// Change information about CDN service.
func (r *ClientConfigService) Update(ctx context.Context, body ClientConfigUpdateParams, opts ...option.RequestOption) (res *CDNAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/clients/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get information about CDN service.
func (r *ClientConfigService) Get(ctx context.Context, opts ...option.RequestOption) (res *CDNAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/clients/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change information about CDN service.
func (r *ClientConfigService) Replace(ctx context.Context, body ClientConfigReplaceParams, opts ...option.RequestOption) (res *CDNAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/clients/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ClientConfigUpdateParams struct {
	// CDN traffic usage limit in gigabytes.
	//
	// When the limit is reached, we will send an email notification.
	UtilizationLevel param.Opt[int64] `json:"utilization_level,omitzero"`
	paramObj
}

func (r ClientConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientConfigUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientConfigUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientConfigReplaceParams struct {
	// CDN traffic usage limit in gigabytes.
	//
	// When the limit is reached, we will send an email notification.
	UtilizationLevel param.Opt[int64] `json:"utilization_level,omitzero"`
	paramObj
}

func (r ClientConfigReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientConfigReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientConfigReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
