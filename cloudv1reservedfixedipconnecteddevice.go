// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1ReservedFixedIPConnectedDeviceService contains methods and other services
// that help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1ReservedFixedIPConnectedDeviceService] method instead.
type CloudV1ReservedFixedIPConnectedDeviceService struct {
	Options []option.RequestOption
}

// NewCloudV1ReservedFixedIPConnectedDeviceService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewCloudV1ReservedFixedIPConnectedDeviceService(opts ...option.RequestOption) (r *CloudV1ReservedFixedIPConnectedDeviceService) {
	r = &CloudV1ReservedFixedIPConnectedDeviceService{}
	r.Options = opts
	return
}

// List instance ports that share VIP
func (r *CloudV1ReservedFixedIPConnectedDeviceService) List(ctx context.Context, projectID int64, regionID int64, portID string, opts ...option.RequestOption) (res *ConnectedDevicesVip, err error) {
	opts = append(r.Options[:], opts...)
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s/connected_devices", projectID, regionID, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Add ports that share VIP
func (r *CloudV1ReservedFixedIPConnectedDeviceService) Add(ctx context.Context, projectID int64, regionID int64, portID string, body CloudV1ReservedFixedIPConnectedDeviceAddParams, opts ...option.RequestOption) (res *ConnectedDevicesVip, err error) {
	opts = append(r.Options[:], opts...)
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s/connected_devices", projectID, regionID, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Replace ports that share VIP
func (r *CloudV1ReservedFixedIPConnectedDeviceService) Replace(ctx context.Context, projectID int64, regionID int64, portID string, body CloudV1ReservedFixedIPConnectedDeviceReplaceParams, opts ...option.RequestOption) (res *ConnectedDevicesVip, err error) {
	opts = append(r.Options[:], opts...)
	if portID == "" {
		err = errors.New("missing required port_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/reserved_fixed_ips/%v/%v/%s/connected_devices", projectID, regionID, portID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ConnectedDevicesVip struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []ConnectedDevicesVipResult `json:"results,required"`
	JSON    connectedDevicesVipJSON     `json:"-"`
}

// connectedDevicesVipJSON contains the JSON metadata for the struct
// [ConnectedDevicesVip]
type connectedDevicesVipJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectedDevicesVip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectedDevicesVipJSON) RawJSON() string {
	return r.raw
}

type ConnectedDevicesVipResult struct {
	// ID of the instance that owns the port
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	// Name of the instance that owns the port
	InstanceName string `json:"instance_name,required"`
	// IP addresses assigned to this port
	IPAssignments []PortIPWithSubnet `json:"ip_assignments,required"`
	// Network details
	Network Network `json:"network,required"`
	// Port ID that shares VIP
	PortID string                        `json:"port_id,required" format:"uuid4"`
	JSON   connectedDevicesVipResultJSON `json:"-"`
}

// connectedDevicesVipResultJSON contains the JSON metadata for the struct
// [ConnectedDevicesVipResult]
type connectedDevicesVipResultJSON struct {
	InstanceID    apijson.Field
	InstanceName  apijson.Field
	IPAssignments apijson.Field
	Network       apijson.Field
	PortID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ConnectedDevicesVipResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectedDevicesVipResultJSON) RawJSON() string {
	return r.raw
}

type PortIDsForVipParam struct {
	// List of port IDs that will share one VIP
	PortIDs param.Field[[]string] `json:"port_ids" format:"uuid4"`
}

func (r PortIDsForVipParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1ReservedFixedIPConnectedDeviceAddParams struct {
	PortIDsForVip PortIDsForVipParam `json:"port_ids_for_vip"`
}

func (r CloudV1ReservedFixedIPConnectedDeviceAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PortIDsForVip)
}

type CloudV1ReservedFixedIPConnectedDeviceReplaceParams struct {
	PortIDsForVip PortIDsForVipParam `json:"port_ids_for_vip"`
}

func (r CloudV1ReservedFixedIPConnectedDeviceReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PortIDsForVip)
}
