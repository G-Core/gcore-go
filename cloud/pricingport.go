// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// PricingPortService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPricingPortService] method instead.
type PricingPortService struct {
	Options []option.RequestOption
}

// NewPricingPortService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPricingPortService(opts ...option.RequestOption) (r PricingPortService) {
	r = PricingPortService{}
	r.Options = opts
	return
}

// Get the price of a port before you create it. Send the same body that you send
// to create a port. The API reads only the fields that change the price.
func (r *PricingPortService) Preview(ctx context.Context, params PricingPortPreviewParams, opts ...option.RequestOption) (res *PortPrice, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/pricing/ports", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type PricingPortPreviewParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Region ID where the port is created
	RegionID int64 `json:"region_id" api:"required"`
	// Network to create the port in. When omitted, it is the network of the subnets in
	// `ip_assignments`, or the region's public network if `ip_assignments` is omitted.
	// Every subnet in `ip_assignments` must belong to it.
	NetworkID param.Opt[string] `json:"network_id,omitzero" format:"uuid4"`
	// IP addresses to assign. All subnets must belong to the same network. Repeat a
	// subnet to get several addresses from it. Omit the field to get one address from
	// an IPv4 subnet.
	IPAssignments []IPAssignmentRequestParam `json:"ip_assignments,omitzero"`
	paramObj
}

func (r PricingPortPreviewParams) MarshalJSON() (data []byte, err error) {
	type shadow PricingPortPreviewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PricingPortPreviewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
