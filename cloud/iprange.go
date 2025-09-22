// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
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

// Returns the complete list of IPv4 and IPv6 address ranges that Cloud uses for
// outbound (egress) traffic.
//
// Typical reasons to call this endpoint:
//
//   - Host-file delivery workflows – You upload images or other assets to the Cloud
//     and share a download link that points to your own infrastructure. Add these
//     egress prefixes to your firewall or object-storage allow-list so our clients
//     can fetch the files without being blocked.
//   - Push integrations / webhooks – You subscribe to the user-actions event log and
//     Cloud pushes events to your listener endpoint. Whitelisting the egress IP
//     ranges lets you accept only traffic that originates from us.
//   - General security controls, audit tooling, or SIEM rules that need to verify
//     that traffic truly comes from the Cloud.
//
// The list is global (covers all regions) and refreshed automatically whenever
// Gcore allocates new egress IP space. The response is an array of CIDR blocks;
// duplicate prefixes are not returned.
func (r *IPRangeService) List(ctx context.Context, opts ...option.RequestOption) (res *IPRanges, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cloud/public/v1/ipranges/egress"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type IPRanges struct {
	// IP ranges list
	Ranges []string `json:"ranges,required" format:"ipv4interface"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ranges      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPRanges) RawJSON() string { return r.JSON.raw }
func (r *IPRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
