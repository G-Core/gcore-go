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

// PricingInstanceService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPricingInstanceService] method instead.
type PricingInstanceService struct {
	Options []option.RequestOption
	Flavors PricingInstanceFlavorService
}

// NewPricingInstanceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPricingInstanceService(opts ...option.RequestOption) (r PricingInstanceService) {
	r = PricingInstanceService{}
	r.Options = opts
	r.Flavors = NewPricingInstanceFlavorService(opts...)
	return
}

// Get the price of an existing instance. The price covers the flavor, the volumes,
// the external IP addresses and the floating IPs.
func (r *PricingInstanceService) Get(ctx context.Context, instanceID string, query PricingInstanceGetParams, opts ...option.RequestOption) (res *ServerPrice, err error) {
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
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v1/projects/%v/pricing/instances/%s", query.ProjectID.Value, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the price of an instance before you create it. Send the same body that you
// send to create an instance. The API reads only the fields that change the price.
// The price includes the flavor, the volumes, the ports and the floating IPs. It
// also includes the existing volumes, ports and floating IPs that you attach.
func (r *PricingInstanceService) Preview(ctx context.Context, params PricingInstancePreviewParams, opts ...option.RequestOption) (res *ServerPrice, err error) {
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
	path := fmt.Sprintf("cloud/v1/projects/%v/pricing/instances", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type PricingInstanceGetParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	paramObj
}

type PricingInstancePreviewParams struct {
	// Project ID.
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. The
	// body of the regular create: the volumes exist.
	OfInstancePricePreviewReferenceBody *PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBody `json:",inline"`
	// This field is a request body variant, only one variant field can be set. The
	// body of the workflow create: each volume is an existing one or a new one.
	OfInstancePricePreviewWorkflowBody *PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBody `json:",inline"`

	paramObj
}

func (u PricingInstancePreviewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInstancePricePreviewReferenceBody, u.OfInstancePricePreviewWorkflowBody)
}
func (r *PricingInstancePreviewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The body of the regular create: the volumes exist.
//
// The properties FlavorName, RegionID are required.
type PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBody struct {
	// Flavor of the server.
	FlavorName string `json:"flavor_name" api:"required"`
	// Region of the server.
	RegionID int64 `json:"region_id" api:"required"`
	// Existing boot volume. Omit it to price the server without a boot volume.
	BootVolumeID param.Opt[string] `json:"boot_volume_id,omitzero" format:"uuid4"`
	// Existing data volumes.
	DataVolumeIDs []string `json:"data_volume_ids,omitzero" format:"uuid4"`
	// Ports in NIC order. An entry with `port_id` is an existing port, any other entry
	// is a new one. Omit it to price the server without ports.
	Interfaces []PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBodyInterfaceUnion `json:"interfaces,omitzero"`
	paramObj
}

func (r PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBody) MarshalJSON() (data []byte, err error) {
	type shadow PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBodyInterfaceUnion struct {
	OfPricePortReference *PricePortReferenceParam `json:",omitzero,inline"`
	OfPricePortSpec      *PricePortSpecParam      `json:",omitzero,inline"`
	paramUnion
}

func (u PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBodyInterfaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPricePortReference, u.OfPricePortSpec)
}
func (u *PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBodyInterfaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBodyInterfaceUnion) asAny() any {
	if !param.IsOmitted(u.OfPricePortReference) {
		return u.OfPricePortReference
	} else if !param.IsOmitted(u.OfPricePortSpec) {
		return u.OfPricePortSpec
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBodyInterfaceUnion) GetPortID() *string {
	if vt := u.OfPricePortReference; vt != nil {
		return &vt.PortID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBodyInterfaceUnion) GetIPAssignments() []PricePortSpecIPAssignmentParam {
	if vt := u.OfPricePortSpec; vt != nil {
		return vt.IPAssignments
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBodyInterfaceUnion) GetNetworkID() *string {
	if vt := u.OfPricePortSpec; vt != nil && vt.NetworkID.Valid() {
		return &vt.NetworkID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's FloatingIP property, if present.
func (u PricingInstancePreviewParamsBodyInstancePricePreviewReferenceBodyInterfaceUnion) GetFloatingIP() *PriceFloatingIPParam {
	if vt := u.OfPricePortReference; vt != nil {
		return &vt.FloatingIP
	} else if vt := u.OfPricePortSpec; vt != nil {
		return &vt.FloatingIP
	}
	return nil
}

// The body of the workflow create: each volume is an existing one or a new one.
//
// The properties FlavorName, RegionID are required.
type PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBody struct {
	// Flavor of the server.
	FlavorName string `json:"flavor_name" api:"required"`
	// Region of the server.
	RegionID int64 `json:"region_id" api:"required"`
	// An entry with `volume_id` is an existing volume, any other entry is a new one.
	// Omit it to price the server without a boot volume.
	BootVolume PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyBootVolumeUnion `json:"boot_volume,omitzero"`
	// An entry with `volume_id` is an existing volume, any other entry is a new one.
	DataVolumes []PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyDataVolumeUnion `json:"data_volumes,omitzero"`
	// Ports in NIC order. An entry with `port_id` is an existing port, any other entry
	// is a new one. Omit it to price the server without ports.
	Interfaces []PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyInterfaceUnion `json:"interfaces,omitzero"`
	paramObj
}

func (r PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBody) MarshalJSON() (data []byte, err error) {
	type shadow PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyBootVolumeUnion struct {
	OfPriceVolumeReference *PriceVolumeReferenceParam `json:",omitzero,inline"`
	OfPriceVolumeSpec      *PriceVolumeSpecParam      `json:",omitzero,inline"`
	paramUnion
}

func (u PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyBootVolumeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPriceVolumeReference, u.OfPriceVolumeSpec)
}
func (u *PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyBootVolumeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyBootVolumeUnion) asAny() any {
	if !param.IsOmitted(u.OfPriceVolumeReference) {
		return u.OfPriceVolumeReference
	} else if !param.IsOmitted(u.OfPriceVolumeSpec) {
		return u.OfPriceVolumeSpec
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyDataVolumeUnion struct {
	OfPriceVolumeReference *PriceVolumeReferenceParam `json:",omitzero,inline"`
	OfPriceVolumeSpec      *PriceVolumeSpecParam      `json:",omitzero,inline"`
	paramUnion
}

func (u PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyDataVolumeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPriceVolumeReference, u.OfPriceVolumeSpec)
}
func (u *PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyDataVolumeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyDataVolumeUnion) asAny() any {
	if !param.IsOmitted(u.OfPriceVolumeReference) {
		return u.OfPriceVolumeReference
	} else if !param.IsOmitted(u.OfPriceVolumeSpec) {
		return u.OfPriceVolumeSpec
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyInterfaceUnion struct {
	OfPricePortReference *PricePortReferenceParam `json:",omitzero,inline"`
	OfPricePortSpec      *PricePortSpecParam      `json:",omitzero,inline"`
	paramUnion
}

func (u PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyInterfaceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPricePortReference, u.OfPricePortSpec)
}
func (u *PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyInterfaceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyInterfaceUnion) asAny() any {
	if !param.IsOmitted(u.OfPricePortReference) {
		return u.OfPricePortReference
	} else if !param.IsOmitted(u.OfPricePortSpec) {
		return u.OfPricePortSpec
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyInterfaceUnion) GetPortID() *string {
	if vt := u.OfPricePortReference; vt != nil {
		return &vt.PortID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyInterfaceUnion) GetIPAssignments() []PricePortSpecIPAssignmentParam {
	if vt := u.OfPricePortSpec; vt != nil {
		return vt.IPAssignments
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyInterfaceUnion) GetNetworkID() *string {
	if vt := u.OfPricePortSpec; vt != nil && vt.NetworkID.Valid() {
		return &vt.NetworkID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's FloatingIP property, if present.
func (u PricingInstancePreviewParamsBodyInstancePricePreviewWorkflowBodyInterfaceUnion) GetFloatingIP() *PriceFloatingIPParam {
	if vt := u.OfPricePortReference; vt != nil {
		return &vt.FloatingIP
	} else if vt := u.OfPricePortSpec; vt != nil {
		return &vt.FloatingIP
	}
	return nil
}
