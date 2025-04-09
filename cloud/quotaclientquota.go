// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// QuotaClientQuotaService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuotaClientQuotaService] method instead.
type QuotaClientQuotaService struct {
	Options []option.RequestOption
}

// NewQuotaClientQuotaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewQuotaClientQuotaService(opts ...option.RequestOption) (r QuotaClientQuotaService) {
	r = QuotaClientQuotaService{}
	r.Options = opts
	return
}

// Get combined client quotas, regional and global.
func (r *QuotaClientQuotaService) Get(ctx context.Context, opts ...option.RequestOption) (res *AllClientQuotas, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/client_quotas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AllClientQuotas struct {
	// Global entity quotas
	GlobalQuotas GlobalQuotas `json:"global_quotas"`
	// Regional entity quotas. Only contains initialized quotas.
	RegionalQuotas []RegionalQuotas `json:"regional_quotas"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		GlobalQuotas   resp.Field
		RegionalQuotas resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllClientQuotas) RawJSON() string { return r.JSON.raw }
func (r *AllClientQuotas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
