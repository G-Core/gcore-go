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

// PricingBaremetalServerService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPricingBaremetalServerService] method instead.
type PricingBaremetalServerService struct {
	Options []option.RequestOption
}

// NewPricingBaremetalServerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPricingBaremetalServerService(opts ...option.RequestOption) (r PricingBaremetalServerService) {
	r = PricingBaremetalServerService{}
	r.Options = opts
	return
}

// Get the price of an existing bare metal server. The price covers the flavor, the
// external IP addresses and the floating IPs.
func (r *PricingBaremetalServerService) Get(ctx context.Context, serverID string, query PricingBaremetalServerGetParams, opts ...option.RequestOption) (res *ServerPrice, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return nil, err
	}
	if serverID == "" {
		err = errors.New("missing required server_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/pricing/baremetal/servers/%s", query.ProjectID.Value, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the price of a bare metal server before you create it. Send the same body
// that you send to create a bare metal server. The API reads only the fields that
// change the price. The price includes the flavor, the ports and the floating IPs.
// It also includes the existing ports and floating IPs that you attach.
func (r *PricingBaremetalServerService) Preview(ctx context.Context, params PricingBaremetalServerPreviewParams, opts ...option.RequestOption) (res *ServerPrice, err error) {
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
	path := fmt.Sprintf("cloud/v1/projects/%v/pricing/baremetal/servers", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type PricingBaremetalServerGetParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	paramObj
}

type PricingBaremetalServerPreviewParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Flavor of the server.
	FlavorName string `json:"flavor_name" api:"required"`
	// Region of the server.
	RegionID int64 `json:"region_id" api:"required"`
	// Ports in NIC order. An entry with `port_id` is an existing port, any other entry
	// is a new one. Omit it to price the server without ports.
	Interfaces []PricingBaremetalServerPreviewParamsInterfaceUnion `json:"interfaces,omitzero"`
	paramObj
}

func (r PricingBaremetalServerPreviewParams) MarshalJSON() (data []byte, err error) {
	type shadow PricingBaremetalServerPreviewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PricingBaremetalServerPreviewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PricingBaremetalServerPreviewParamsInterfaceUnion struct {
	OfPricePortReference *PricePortReferenceParam `json:",omitzero,inline"`
	OfPricePortSpec      *PricePortSpecParam      `json:",omitzero,inline"`
	paramUnion
}

func (u PricingBaremetalServerPreviewParamsInterfaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPricePortReference, u.OfPricePortSpec)
}
func (u *PricingBaremetalServerPreviewParamsInterfaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PricingBaremetalServerPreviewParamsInterfaceUnion) asAny() any {
	if !param.IsOmitted(u.OfPricePortReference) {
		return u.OfPricePortReference
	} else if !param.IsOmitted(u.OfPricePortSpec) {
		return u.OfPricePortSpec
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PricingBaremetalServerPreviewParamsInterfaceUnion) GetPortID() *string {
	if vt := u.OfPricePortReference; vt != nil {
		return &vt.PortID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PricingBaremetalServerPreviewParamsInterfaceUnion) GetIPAssignments() []IPAssignmentRequestParam {
	if vt := u.OfPricePortSpec; vt != nil {
		return vt.IPAssignments
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PricingBaremetalServerPreviewParamsInterfaceUnion) GetNetworkID() *string {
	if vt := u.OfPricePortSpec; vt != nil && vt.NetworkID.Valid() {
		return &vt.NetworkID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's FloatingIP property, if present.
func (u PricingBaremetalServerPreviewParamsInterfaceUnion) GetFloatingIP() *PriceFloatingIPParam {
	if vt := u.OfPricePortReference; vt != nil {
		return &vt.FloatingIP
	} else if vt := u.OfPricePortSpec; vt != nil {
		return &vt.FloatingIP
	}
	return nil
}
