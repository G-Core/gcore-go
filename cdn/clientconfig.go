// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn

import (
	"context"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
)

// Information about the current state of the CDN service in your account.
//
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

// Get information about CDN service.
func (r *ClientConfigService) Get(ctx context.Context, opts ...option.RequestOption) (res *CDNAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cdn/clients/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
