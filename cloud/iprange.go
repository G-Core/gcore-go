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

// IPRangeService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPRangeService] method instead.
type IPRangeService struct {
	Options []option.RequestOption
}

// NewIPRangeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPRangeService(opts ...option.RequestOption) (r IPRangeService) {
	r = IPRangeService{}
	r.Options = opts
	return
}

// List of all Edge Cloud Egress Public IPs.
func (r *IPRangeService) List(ctx context.Context, opts ...option.RequestOption) (res *IPRanges, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/public/v1/ipranges/egress"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type IPRanges struct {
	// IP ranges list
	Ranges []string `json:"ranges,required" format:"ipv4interface"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Ranges      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPRanges) RawJSON() string { return r.JSON.raw }
func (r *IPRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
