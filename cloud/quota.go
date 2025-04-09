// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"github.com/stainless-sdks/gcore-go/option"
)

// QuotaService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuotaService] method instead.
type QuotaService struct {
	Options       []option.RequestOption
	ClientQuotas  QuotaClientQuotaService
	Global        QuotaGlobalService
	Regional      QuotaRegionalService
	LimitsRequest QuotaLimitsRequestService
}

// NewQuotaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQuotaService(opts ...option.RequestOption) (r QuotaService) {
	r = QuotaService{}
	r.Options = opts
	r.ClientQuotas = NewQuotaClientQuotaService(opts...)
	r.Global = NewQuotaGlobalService(opts...)
	r.Regional = NewQuotaRegionalService(opts...)
	r.LimitsRequest = NewQuotaLimitsRequestService(opts...)
	return
}
