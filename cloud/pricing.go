// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// Pricing gives the prices of instances, bare metal servers and ports in all
// regions of the project: flavor prices, the price of a configuration before you
// create it, and the price of an existing server.
//
// PricingService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPricingService] method instead.
type PricingService struct {
	Options   []option.RequestOption
	Instances PricingInstanceService
	Baremetal PricingBaremetalService
	Ports     PricingPortService
}

// NewPricingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPricingService(opts ...option.RequestOption) (r PricingService) {
	r = PricingService{}
	r.Options = opts
	r.Instances = NewPricingInstanceService(opts...)
	r.Baremetal = NewPricingBaremetalService(opts...)
	r.Ports = NewPricingPortService(opts...)
	return
}

type FlavorPrice struct {
	// Currency code.
	Currency string `json:"currency" api:"required"`
	// Discount in percent. 0 when there is no discount.
	DiscountPercent float64 `json:"discount_percent" api:"required"`
	// Flavor name.
	FlavorName string `json:"flavor_name" api:"required"`
	// Price per hour after the discount.
	Hourly float64 `json:"hourly" api:"required"`
	// Price per month after the discount.
	Monthly float64 `json:"monthly" api:"required"`
	// Price per month before the discount.
	MonthlyBeforeDiscount float64 `json:"monthly_before_discount" api:"required"`
	// Region ID.
	RegionID int64 `json:"region_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency              respjson.Field
		DiscountPercent       respjson.Field
		FlavorName            respjson.Field
		Hourly                respjson.Field
		Monthly               respjson.Field
		MonthlyBeforeDiscount respjson.Field
		RegionID              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlavorPrice) RawJSON() string { return r.JSON.raw }
func (r *FlavorPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SubnetID is required.
type IPAssignmentRequestParam struct {
	// ID of the subnet
	SubnetID string `json:"subnet_id" api:"required" format:"uuid4"`
	// IP address to assign. Omit it to get any free address of the subnet, which is
	// the only option when the subnet's free addresses are not known to the caller.
	IPAddress param.Opt[string] `json:"ip_address,omitzero" format:"ipv4"`
	paramObj
}

func (r IPAssignmentRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow IPAssignmentRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IPAssignmentRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortPrice struct {
	// Currency code. Null when the port has no IPv4 address on the region's public
	// network. Then all prices are 0.
	Currency string `json:"currency" api:"required"`
	// Discount in percent. 0 when there is no discount.
	DiscountPercent float64 `json:"discount_percent" api:"required"`
	// Price per hour after the discount.
	Hourly float64 `json:"hourly" api:"required"`
	// Price per month after the discount.
	Monthly float64 `json:"monthly" api:"required"`
	// Price per month before the discount.
	MonthlyBeforeDiscount float64 `json:"monthly_before_discount" api:"required"`
	// Tax in percent. All prices include it. Null when the port has no IPv4 address on
	// the region's public network. Then all prices are 0.
	TaxPercent float64 `json:"tax_percent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency              respjson.Field
		DiscountPercent       respjson.Field
		Hourly                respjson.Field
		Monthly               respjson.Field
		MonthlyBeforeDiscount respjson.Field
		TaxPercent            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortPrice) RawJSON() string { return r.JSON.raw }
func (r *PortPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PriceFloatingIPParam struct {
	// ID of an existing floating IP. Omit it for a new floating IP.
	ID param.Opt[string] `json:"id,omitzero" format:"uuid4"`
	paramObj
}

func (r PriceFloatingIPParam) MarshalJSON() (data []byte, err error) {
	type shadow PriceFloatingIPParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PriceFloatingIPParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PortID is required.
type PricePortReferenceParam struct {
	// Existing port.
	PortID string `json:"port_id" api:"required" format:"uuid4"`
	// Floating IP of the port.
	FloatingIP PriceFloatingIPParam `json:"floating_ip,omitzero"`
	paramObj
}

func (r PricePortReferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow PricePortReferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PricePortReferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PricePortSpecParam struct {
	// Network of the new port. When omitted, it is the network of the subnets in
	// `ip_assignments`, or the public network of the region if `ip_assignments` is
	// omitted.
	NetworkID param.Opt[string] `json:"network_id,omitzero" format:"uuid4"`
	// Floating IP of the new port.
	FloatingIP PriceFloatingIPParam `json:"floating_ip,omitzero"`
	// IP addresses of the new port. Omit the field to get one address from an IPv4
	// subnet.
	IPAssignments []PricePortSpecIPAssignmentParam `json:"ip_assignments,omitzero"`
	paramObj
}

func (r PricePortSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow PricePortSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PricePortSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SubnetID is required.
type PricePortSpecIPAssignmentParam struct {
	// ID of the subnet
	SubnetID string `json:"subnet_id" api:"required" format:"uuid4"`
	// IP address to assign. Omit it to get any free address of the subnet, which is
	// the only option when the subnet's free addresses are not known to the caller.
	IPAddress param.Opt[string] `json:"ip_address,omitzero" format:"ipv4"`
	// Floating IP of this address of the new port.
	FloatingIP PriceFloatingIPParam `json:"floating_ip,omitzero"`
	paramObj
}

func (r PricePortSpecIPAssignmentParam) MarshalJSON() (data []byte, err error) {
	type shadow PricePortSpecIPAssignmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PricePortSpecIPAssignmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property VolumeID is required.
type PriceVolumeReferenceParam struct {
	// Existing volume.
	VolumeID string `json:"volume_id" api:"required" format:"uuid4"`
	paramObj
}

func (r PriceVolumeReferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow PriceVolumeReferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PriceVolumeReferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PriceVolumeSpecParam struct {
	// Size in GiB. Required for a blank or image volume. A volume from a snapshot or a
	// volume has the size of its source, and a larger size grows it.
	Size param.Opt[int64] `json:"size,omitzero"`
	// Image, snapshot or volume to create the volume from. Omit it for a blank volume.
	Source PriceVolumeSpecSourceParam `json:"source,omitzero"`
	// Volume type. A volume from a snapshot or a volume has the type of its source.
	// Any other volume is `standard` by default.
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	Type PriceVolumeSpecType `json:"type,omitzero"`
	paramObj
}

func (r PriceVolumeSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow PriceVolumeSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PriceVolumeSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image, snapshot or volume to create the volume from. Omit it for a blank volume.
//
// The properties ID, Type are required.
type PriceVolumeSpecSourceParam struct {
	// ID of the image, the snapshot or the volume to clone. Only a source volume of
	// type `standard` or `ssd_hiiops` can be cloned.
	ID string `json:"id" api:"required" format:"uuid4"`
	// What the volume is created from.
	//
	// Any of "image", "snapshot", "volume".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r PriceVolumeSpecSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow PriceVolumeSpecSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PriceVolumeSpecSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PriceVolumeSpecSourceParam](
		"type", "image", "snapshot", "volume",
	)
}

// Volume type. A volume from a snapshot or a volume has the type of its source.
// Any other volume is `standard` by default.
type PriceVolumeSpecType string

const (
	PriceVolumeSpecTypeCold          PriceVolumeSpecType = "cold"
	PriceVolumeSpecTypeSsdHiiops     PriceVolumeSpecType = "ssd_hiiops"
	PriceVolumeSpecTypeSsdLocal      PriceVolumeSpecType = "ssd_local"
	PriceVolumeSpecTypeSsdLowlatency PriceVolumeSpecType = "ssd_lowlatency"
	PriceVolumeSpecTypeStandard      PriceVolumeSpecType = "standard"
	PriceVolumeSpecTypeUltra         PriceVolumeSpecType = "ultra"
)

type ServerPrice struct {
	// Price of the boot volume. Null when the body has no boot volume.
	BootVolume ServerPriceBootVolume `json:"boot_volume" api:"required"`
	// Currency code.
	Currency string `json:"currency" api:"required"`
	// Prices of the data volumes, in the order of the request.
	DataVolumes []ServerPriceDataVolume `json:"data_volumes" api:"required"`
	// Rate of the outgoing traffic. It is not part of the totals. Null when the
	// traffic of the server has no tariff.
	EgressTraffic ServerPriceEgressTraffic `json:"egress_traffic" api:"required"`
	// Price of the IPv4 addresses that ports have on the public network.
	ExternalIPs ServerPriceExternalIPs `json:"external_ips" api:"required"`
	// Price of the flavor.
	Flavor ServerPriceFlavor `json:"flavor" api:"required"`
	// Price of the floating IPs. A floating IP is a separate address that maps to a
	// port.
	FloatingIPs ServerPriceFloatingIPs `json:"floating_ips" api:"required"`
	// Sum of the hourly prices of all sections.
	Hourly float64 `json:"hourly" api:"required"`
	// Sum of the monthly prices of all sections.
	Monthly float64 `json:"monthly" api:"required"`
	// Sum of the monthly prices before the discount.
	MonthlyBeforeDiscount float64 `json:"monthly_before_discount" api:"required"`
	// Tax in percent. The prices include it.
	TaxPercent float64 `json:"tax_percent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BootVolume            respjson.Field
		Currency              respjson.Field
		DataVolumes           respjson.Field
		EgressTraffic         respjson.Field
		ExternalIPs           respjson.Field
		Flavor                respjson.Field
		FloatingIPs           respjson.Field
		Hourly                respjson.Field
		Monthly               respjson.Field
		MonthlyBeforeDiscount respjson.Field
		TaxPercent            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerPrice) RawJSON() string { return r.JSON.raw }
func (r *ServerPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price of the boot volume. Null when the body has no boot volume.
type ServerPriceBootVolume struct {
	// Discount in percent. 0 when there is no discount.
	DiscountPercent float64 `json:"discount_percent" api:"required"`
	// Price per hour after the discount.
	Hourly float64 `json:"hourly" api:"required"`
	// Price per month after the discount.
	Monthly float64 `json:"monthly" api:"required"`
	// Price per month before the discount.
	MonthlyBeforeDiscount float64 `json:"monthly_before_discount" api:"required"`
	// Volume size in GiB. The price uses the size in GB
	Size int64 `json:"size" api:"required"`
	// Volume type.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DiscountPercent       respjson.Field
		Hourly                respjson.Field
		Monthly               respjson.Field
		MonthlyBeforeDiscount respjson.Field
		Size                  respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerPriceBootVolume) RawJSON() string { return r.JSON.raw }
func (r *ServerPriceBootVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServerPriceDataVolume struct {
	// Discount in percent. 0 when there is no discount.
	DiscountPercent float64 `json:"discount_percent" api:"required"`
	// Price per hour after the discount.
	Hourly float64 `json:"hourly" api:"required"`
	// Price per month after the discount.
	Monthly float64 `json:"monthly" api:"required"`
	// Price per month before the discount.
	MonthlyBeforeDiscount float64 `json:"monthly_before_discount" api:"required"`
	// Volume size in GiB. The price uses the size in GB
	Size int64 `json:"size" api:"required"`
	// Volume type.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DiscountPercent       respjson.Field
		Hourly                respjson.Field
		Monthly               respjson.Field
		MonthlyBeforeDiscount respjson.Field
		Size                  respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerPriceDataVolume) RawJSON() string { return r.JSON.raw }
func (r *ServerPriceDataVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rate of the outgoing traffic. It is not part of the totals. Null when the
// traffic of the server has no tariff.
type ServerPriceEgressTraffic struct {
	// Discount in percent. 0 when there is no discount.
	DiscountPercent float64 `json:"discount_percent" api:"required"`
	// Price of one GB of used traffic after the discount.
	PerGB float64 `json:"per_gb" api:"required"`
	// Price of one GB before the discount.
	PerGBBeforeDiscount float64 `json:"per_gb_before_discount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DiscountPercent     respjson.Field
		PerGB               respjson.Field
		PerGBBeforeDiscount respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerPriceEgressTraffic) RawJSON() string { return r.JSON.raw }
func (r *ServerPriceEgressTraffic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price of the IPv4 addresses that ports have on the public network.
type ServerPriceExternalIPs struct {
	// Number of IP addresses. 0 when the server has none.
	Count int64 `json:"count" api:"required"`
	// Discount in percent. 0 when there is no discount.
	DiscountPercent float64 `json:"discount_percent" api:"required"`
	// Price per hour after the discount.
	Hourly float64 `json:"hourly" api:"required"`
	// Price per month after the discount.
	Monthly float64 `json:"monthly" api:"required"`
	// Price per month before the discount.
	MonthlyBeforeDiscount float64 `json:"monthly_before_discount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count                 respjson.Field
		DiscountPercent       respjson.Field
		Hourly                respjson.Field
		Monthly               respjson.Field
		MonthlyBeforeDiscount respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerPriceExternalIPs) RawJSON() string { return r.JSON.raw }
func (r *ServerPriceExternalIPs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price of the flavor.
type ServerPriceFlavor struct {
	// Number of servers. The create makes one server.
	Count int64 `json:"count" api:"required"`
	// Discount in percent. 0 when there is no discount.
	DiscountPercent float64 `json:"discount_percent" api:"required"`
	// Price per hour after the discount.
	Hourly float64 `json:"hourly" api:"required"`
	// Price per month after the discount.
	Monthly float64 `json:"monthly" api:"required"`
	// Price per month before the discount.
	MonthlyBeforeDiscount float64 `json:"monthly_before_discount" api:"required"`
	// Flavor name.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count                 respjson.Field
		DiscountPercent       respjson.Field
		Hourly                respjson.Field
		Monthly               respjson.Field
		MonthlyBeforeDiscount respjson.Field
		Name                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerPriceFlavor) RawJSON() string { return r.JSON.raw }
func (r *ServerPriceFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Price of the floating IPs. A floating IP is a separate address that maps to a
// port.
type ServerPriceFloatingIPs struct {
	// Number of IP addresses. 0 when the server has none.
	Count int64 `json:"count" api:"required"`
	// Discount in percent. 0 when there is no discount.
	DiscountPercent float64 `json:"discount_percent" api:"required"`
	// Price per hour after the discount.
	Hourly float64 `json:"hourly" api:"required"`
	// Price per month after the discount.
	Monthly float64 `json:"monthly" api:"required"`
	// Price per month before the discount.
	MonthlyBeforeDiscount float64 `json:"monthly_before_discount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count                 respjson.Field
		DiscountPercent       respjson.Field
		Hourly                respjson.Field
		Monthly               respjson.Field
		MonthlyBeforeDiscount respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServerPriceFloatingIPs) RawJSON() string { return r.JSON.raw }
func (r *ServerPriceFloatingIPs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
