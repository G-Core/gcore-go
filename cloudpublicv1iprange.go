// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudPublicV1IprangeService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudPublicV1IprangeService] method instead.
type CloudPublicV1IprangeService struct {
	Options []option.RequestOption
}

// NewCloudPublicV1IprangeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudPublicV1IprangeService(opts ...option.RequestOption) (r *CloudPublicV1IprangeService) {
	r = &CloudPublicV1IprangeService{}
	r.Options = opts
	return
}

// List of all Edge Cloud Egress Public IPs.
func (r *CloudPublicV1IprangeService) ListEgress(ctx context.Context, opts ...option.RequestOption) (res *CloudPublicV1IprangeListEgressResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/public/v1/ipranges/egress"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CloudPublicV1IprangeListEgressResponse struct {
	// IP ranges list
	Ranges []string                                   `json:"ranges,required" format:"ipv4interface"`
	JSON   cloudPublicV1IprangeListEgressResponseJSON `json:"-"`
}

// cloudPublicV1IprangeListEgressResponseJSON contains the JSON metadata for the
// struct [CloudPublicV1IprangeListEgressResponse]
type cloudPublicV1IprangeListEgressResponseJSON struct {
	Ranges      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudPublicV1IprangeListEgressResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudPublicV1IprangeListEgressResponseJSON) RawJSON() string {
	return r.raw
}
