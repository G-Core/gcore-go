// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"net/http"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
)

// AdvancedRuleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedRuleService] method instead.
type AdvancedRuleService struct {
	Options []option.RequestOption
}

// NewAdvancedRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdvancedRuleService(opts ...option.RequestOption) (r AdvancedRuleService) {
	r = AdvancedRuleService{}
	r.Options = opts
	return
}

// Retrieve an advanced rules descriptor
func (r *AdvancedRuleService) List(ctx context.Context, opts ...option.RequestOption) (res *WaapAdvancedRuleDescriptorList, err error) {
	opts = append(r.Options[:], opts...)
	path := "waap/v1/advanced-rules/descriptor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
