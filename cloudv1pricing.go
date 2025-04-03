// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/shared"
	"github.com/tidwall/gjson"
)

// CloudV1PricingService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1PricingService] method instead.
type CloudV1PricingService struct {
	Options   []option.RequestOption
	Faas      *CloudV1PricingFaaService
	AI        *CloudV1PricingAIService
	Inference *CloudV1PricingInferenceService
}

// NewCloudV1PricingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1PricingService(opts ...option.RequestOption) (r *CloudV1PricingService) {
	r = &CloudV1PricingService{}
	r.Options = opts
	r.Faas = NewCloudV1PricingFaaService(opts...)
	r.AI = NewCloudV1PricingAIService(opts...)
	r.Inference = NewCloudV1PricingInferenceService(opts...)
	return
}

// For calculate discount of reservation compare price and over_commit_price
func (r *CloudV1PricingService) GetBillingReservationPrices(ctx context.Context, regionID int64, body CloudV1PricingGetBillingReservationPricesParams, opts ...option.RequestOption) (res *CloudV1PricingGetBillingReservationPricesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/pricing/%v/billing_reservation", regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the DDoS protection price based on mitigation capacity.
func (r *CloudV1PricingService) GetDdosPrice(ctx context.Context, regionID int64, opts ...option.RequestOption) (res *CloudV1PricingGetDdosPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/pricing/%v/ddos", regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get billing rate of the existing instance based on its current configuration.
func (r *CloudV1PricingService) GetInstancePrice(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/pricing/%v/%v/instances/%s", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the price for Logging-as-a-Service (LaaS) in the specified region.
func (r *CloudV1PricingService) GetLaasPrice(ctx context.Context, regionID int64, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/pricing/%v/laas", regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get maximum usage cost of resources if all snapshots are created by the policy
func (r *CloudV1PricingService) GetLifecyclePolicyCost(ctx context.Context, projectID int64, regionID int64, body CloudV1PricingGetLifecyclePolicyCostParams, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/pricing/%v/%v/lifecycle_policy", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Preview billing price of the container based on the provided specifications.
func (r *CloudV1PricingService) PreviewContainerPrice(ctx context.Context, projectID int64, regionID int64, body CloudV1PricingPreviewContainerPriceParams, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/pricing/%v/%v/caas", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Preview billing price of the file share of given type and size
func (r *CloudV1PricingService) PreviewFileSharePrice(ctx context.Context, projectID int64, regionID int64, body CloudV1PricingPreviewFileSharePriceParams, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/pricing/%v/%v/file_shares", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Preview billing price of the floating IP addresses.
func (r *CloudV1PricingService) PreviewFloatingIPPrice(ctx context.Context, projectID int64, regionID int64, body CloudV1PricingPreviewFloatingIPPriceParams, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/pricing/%v/%v/floating_ips", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Preview billing price of the image of given size.
func (r *CloudV1PricingService) PreviewImagePrice(ctx context.Context, projectID int64, regionID int64, body CloudV1PricingPreviewImagePriceParams, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/pricing/%v/%v/images", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Preview billing price of the load balancer.
func (r *CloudV1PricingService) PreviewLoadBalancerPrice(ctx context.Context, projectID int64, regionID int64, body CloudV1PricingPreviewLoadBalancerPriceParams, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/pricing/%v/%v/loadbalancers", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Preview billing price of the reserved fixed IP
func (r *CloudV1PricingService) PreviewReservedFixedIPPrice(ctx context.Context, projectID int64, regionID int64, body CloudV1PricingPreviewReservedFixedIPPriceParams, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/pricing/%v/%v/reserved_fixed_ips", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Preview billing price of the snapshots.
func (r *CloudV1PricingService) PreviewSnapshotPrice(ctx context.Context, projectID int64, regionID int64, body CloudV1PricingPreviewSnapshotPriceParams, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/pricing/%v/%v/snapshots", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Preview billing price of the volume of given type and size
func (r *CloudV1PricingService) PreviewVolumePrice(ctx context.Context, projectID int64, regionID int64, body CloudV1PricingPreviewVolumePriceParams, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/pricing/%v/%v/volumes", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BillingReservationAmountPrices struct {
	// Commit price of the item charged per month
	CommitPricePerMonth BillingReservationAmountPricesCommitPricePerMonthUnion `json:"commit_price_per_month,required"`
	// Commit price of the item charged per hour
	CommitPricePerUnit BillingReservationAmountPricesCommitPricePerUnitUnion `json:"commit_price_per_unit,required"`
	// Commit price of the item charged for all period reservation
	CommitPriceTotal BillingReservationAmountPricesCommitPriceTotalUnion `json:"commit_price_total,required"`
	// Currency code (3 letter code per ISO 4217)
	CurrencyCode string `json:"currency_code,required"`
	// Overcommit price of the item charged per month
	OvercommitPricePerMonth BillingReservationAmountPricesOvercommitPricePerMonthUnion `json:"overcommit_price_per_month,required"`
	// Overcommit price of the item charged per hour
	OvercommitPricePerUnit BillingReservationAmountPricesOvercommitPricePerUnitUnion `json:"overcommit_price_per_unit,required"`
	// Overcommit price of the item charged for all period reservation
	OvercommitPriceTotal BillingReservationAmountPricesOvercommitPriceTotalUnion `json:"overcommit_price_total,required"`
	JSON                 billingReservationAmountPricesJSON                      `json:"-"`
}

// billingReservationAmountPricesJSON contains the JSON metadata for the struct
// [BillingReservationAmountPrices]
type billingReservationAmountPricesJSON struct {
	CommitPricePerMonth     apijson.Field
	CommitPricePerUnit      apijson.Field
	CommitPriceTotal        apijson.Field
	CurrencyCode            apijson.Field
	OvercommitPricePerMonth apijson.Field
	OvercommitPricePerUnit  apijson.Field
	OvercommitPriceTotal    apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *BillingReservationAmountPrices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingReservationAmountPricesJSON) RawJSON() string {
	return r.raw
}

// Commit price of the item charged per month
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type BillingReservationAmountPricesCommitPricePerMonthUnion interface {
	ImplementsBillingReservationAmountPricesCommitPricePerMonthUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingReservationAmountPricesCommitPricePerMonthUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Commit price of the item charged per hour
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type BillingReservationAmountPricesCommitPricePerUnitUnion interface {
	ImplementsBillingReservationAmountPricesCommitPricePerUnitUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingReservationAmountPricesCommitPricePerUnitUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Commit price of the item charged for all period reservation
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type BillingReservationAmountPricesCommitPriceTotalUnion interface {
	ImplementsBillingReservationAmountPricesCommitPriceTotalUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingReservationAmountPricesCommitPriceTotalUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Overcommit price of the item charged per month
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type BillingReservationAmountPricesOvercommitPricePerMonthUnion interface {
	ImplementsBillingReservationAmountPricesOvercommitPricePerMonthUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingReservationAmountPricesOvercommitPricePerMonthUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Overcommit price of the item charged per hour
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type BillingReservationAmountPricesOvercommitPricePerUnitUnion interface {
	ImplementsBillingReservationAmountPricesOvercommitPricePerUnitUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingReservationAmountPricesOvercommitPricePerUnitUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Overcommit price of the item charged for all period reservation
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type BillingReservationAmountPricesOvercommitPriceTotalUnion interface {
	ImplementsBillingReservationAmountPricesOvercommitPriceTotalUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BillingReservationAmountPricesOvercommitPriceTotalUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type BillingReservationPricingResourceRequestParam struct {
	// Name from billing features for specific resource
	ResourceName param.Field[string] `json:"resource_name,required"`
	// Resource type
	ResourceType param.Field[BillingReservationType] `json:"resource_type,required"`
	// Resource value
	ResourceValue param.Field[int64] `json:"resource_value,required"`
}

func (r BillingReservationPricingResourceRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillingReservationType string

const (
	BillingReservationTypeFlavor BillingReservationType = "flavor"
)

func (r BillingReservationType) IsKnown() bool {
	switch r {
	case BillingReservationTypeFlavor:
		return true
	}
	return false
}

type CreateReservedFixedIPParam struct {
	// Must be 'external'
	Type param.Field[CreateReservedFixedIPType] `json:"type,required"`
	// Reserved fixed IP will be allocated the given IP address
	IPAddress param.Field[string] `json:"ip_address" format:"ipvanyaddress"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[InterfaceIPFamily] `json:"ip_family"`
	// If reserved fixed IP is a VIP
	IsVip param.Field[bool] `json:"is_vip"`
	// Reserved fixed IP will be allocated in a subnet of this network
	NetworkID param.Field[string] `json:"network_id" format:"uuid4"`
	// Port ID to make a reserved fixed IP (for example, `vip_port_id` of the Load
	// Balancer entity).
	PortID param.Field[string] `json:"port_id" format:"uuid4"`
	// Reserved fixed IP will be allocated in this subnet
	SubnetID param.Field[string] `json:"subnet_id" format:"uuid4"`
}

func (r CreateReservedFixedIPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateReservedFixedIPParam) implementsCreateReservedFixedIPUnionParam() {}

// Satisfied by [CreateReservedFixedIPNewReservedFixedIPExternalSerializerParam],
// [CreateReservedFixedIPNewReservedFixedIPSpecificSubnetSerializerParam],
// [CreateReservedFixedIPNewReservedFixedIPAnySubnetSerializerParam],
// [CreateReservedFixedIPNewReservedFixedIPSpecificIPAddressSerializerParam],
// [CreateReservedFixedIPNewReservedFixedIPSpecificPortSerializerParam],
// [CreateReservedFixedIPParam].
type CreateReservedFixedIPUnionParam interface {
	implementsCreateReservedFixedIPUnionParam()
}

type CreateReservedFixedIPNewReservedFixedIPExternalSerializerParam struct {
	// Must be 'external'
	Type param.Field[CreateReservedFixedIPNewReservedFixedIPExternalSerializerType] `json:"type,required"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[InterfaceIPFamily] `json:"ip_family"`
	// If reserved fixed IP is a VIP
	IsVip param.Field[bool] `json:"is_vip"`
}

func (r CreateReservedFixedIPNewReservedFixedIPExternalSerializerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateReservedFixedIPNewReservedFixedIPExternalSerializerParam) implementsCreateReservedFixedIPUnionParam() {
}

// Must be 'external'
type CreateReservedFixedIPNewReservedFixedIPExternalSerializerType string

const (
	CreateReservedFixedIPNewReservedFixedIPExternalSerializerTypeExternal CreateReservedFixedIPNewReservedFixedIPExternalSerializerType = "external"
)

func (r CreateReservedFixedIPNewReservedFixedIPExternalSerializerType) IsKnown() bool {
	switch r {
	case CreateReservedFixedIPNewReservedFixedIPExternalSerializerTypeExternal:
		return true
	}
	return false
}

type CreateReservedFixedIPNewReservedFixedIPSpecificSubnetSerializerParam struct {
	// Reserved fixed IP will be allocated in this subnet
	SubnetID param.Field[string] `json:"subnet_id,required" format:"uuid4"`
	// Must be 'subnet'.
	Type param.Field[CreateReservedFixedIPNewReservedFixedIPSpecificSubnetSerializerType] `json:"type,required"`
	// If reserved fixed IP is a VIP
	IsVip param.Field[bool] `json:"is_vip"`
}

func (r CreateReservedFixedIPNewReservedFixedIPSpecificSubnetSerializerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateReservedFixedIPNewReservedFixedIPSpecificSubnetSerializerParam) implementsCreateReservedFixedIPUnionParam() {
}

// Must be 'subnet'.
type CreateReservedFixedIPNewReservedFixedIPSpecificSubnetSerializerType string

const (
	CreateReservedFixedIPNewReservedFixedIPSpecificSubnetSerializerTypeSubnet CreateReservedFixedIPNewReservedFixedIPSpecificSubnetSerializerType = "subnet"
)

func (r CreateReservedFixedIPNewReservedFixedIPSpecificSubnetSerializerType) IsKnown() bool {
	switch r {
	case CreateReservedFixedIPNewReservedFixedIPSpecificSubnetSerializerTypeSubnet:
		return true
	}
	return false
}

type CreateReservedFixedIPNewReservedFixedIPAnySubnetSerializerParam struct {
	// Reserved fixed IP will be allocated in a subnet of this network
	NetworkID param.Field[string] `json:"network_id,required" format:"uuid4"`
	// Must be 'any_subnet'.
	Type param.Field[CreateReservedFixedIPNewReservedFixedIPAnySubnetSerializerType] `json:"type,required"`
	// Which subnets should be selected: IPv4, IPv6 or use dual stack.
	IPFamily param.Field[InterfaceIPFamily] `json:"ip_family"`
	// If reserved fixed IP is a VIP
	IsVip param.Field[bool] `json:"is_vip"`
}

func (r CreateReservedFixedIPNewReservedFixedIPAnySubnetSerializerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateReservedFixedIPNewReservedFixedIPAnySubnetSerializerParam) implementsCreateReservedFixedIPUnionParam() {
}

// Must be 'any_subnet'.
type CreateReservedFixedIPNewReservedFixedIPAnySubnetSerializerType string

const (
	CreateReservedFixedIPNewReservedFixedIPAnySubnetSerializerTypeAnySubnet CreateReservedFixedIPNewReservedFixedIPAnySubnetSerializerType = "any_subnet"
)

func (r CreateReservedFixedIPNewReservedFixedIPAnySubnetSerializerType) IsKnown() bool {
	switch r {
	case CreateReservedFixedIPNewReservedFixedIPAnySubnetSerializerTypeAnySubnet:
		return true
	}
	return false
}

type CreateReservedFixedIPNewReservedFixedIPSpecificIPAddressSerializerParam struct {
	// Reserved fixed IP will be allocated the given IP address
	IPAddress param.Field[string] `json:"ip_address,required" format:"ipvanyaddress"`
	// Reserved fixed IP will be allocated in a subnet of this network
	NetworkID param.Field[string] `json:"network_id,required" format:"uuid4"`
	// Must be 'ip_address'.
	Type param.Field[CreateReservedFixedIPNewReservedFixedIPSpecificIPAddressSerializerType] `json:"type,required"`
	// If reserved fixed IP is a VIP
	IsVip param.Field[bool] `json:"is_vip"`
}

func (r CreateReservedFixedIPNewReservedFixedIPSpecificIPAddressSerializerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateReservedFixedIPNewReservedFixedIPSpecificIPAddressSerializerParam) implementsCreateReservedFixedIPUnionParam() {
}

// Must be 'ip_address'.
type CreateReservedFixedIPNewReservedFixedIPSpecificIPAddressSerializerType string

const (
	CreateReservedFixedIPNewReservedFixedIPSpecificIPAddressSerializerTypeIPAddress CreateReservedFixedIPNewReservedFixedIPSpecificIPAddressSerializerType = "ip_address"
)

func (r CreateReservedFixedIPNewReservedFixedIPSpecificIPAddressSerializerType) IsKnown() bool {
	switch r {
	case CreateReservedFixedIPNewReservedFixedIPSpecificIPAddressSerializerTypeIPAddress:
		return true
	}
	return false
}

type CreateReservedFixedIPNewReservedFixedIPSpecificPortSerializerParam struct {
	// Port ID to make a reserved fixed IP (for example, `vip_port_id` of the Load
	// Balancer entity).
	PortID param.Field[string] `json:"port_id,required" format:"uuid4"`
	// Must be 'port'.
	Type param.Field[CreateReservedFixedIPNewReservedFixedIPSpecificPortSerializerType] `json:"type,required"`
}

func (r CreateReservedFixedIPNewReservedFixedIPSpecificPortSerializerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateReservedFixedIPNewReservedFixedIPSpecificPortSerializerParam) implementsCreateReservedFixedIPUnionParam() {
}

// Must be 'port'.
type CreateReservedFixedIPNewReservedFixedIPSpecificPortSerializerType string

const (
	CreateReservedFixedIPNewReservedFixedIPSpecificPortSerializerTypePort CreateReservedFixedIPNewReservedFixedIPSpecificPortSerializerType = "port"
)

func (r CreateReservedFixedIPNewReservedFixedIPSpecificPortSerializerType) IsKnown() bool {
	switch r {
	case CreateReservedFixedIPNewReservedFixedIPSpecificPortSerializerTypePort:
		return true
	}
	return false
}

// Must be 'external'
type CreateReservedFixedIPType string

const (
	CreateReservedFixedIPTypeExternal  CreateReservedFixedIPType = "external"
	CreateReservedFixedIPTypeSubnet    CreateReservedFixedIPType = "subnet"
	CreateReservedFixedIPTypeAnySubnet CreateReservedFixedIPType = "any_subnet"
	CreateReservedFixedIPTypeIPAddress CreateReservedFixedIPType = "ip_address"
	CreateReservedFixedIPTypePort      CreateReservedFixedIPType = "port"
)

func (r CreateReservedFixedIPType) IsKnown() bool {
	switch r {
	case CreateReservedFixedIPTypeExternal, CreateReservedFixedIPTypeSubnet, CreateReservedFixedIPTypeAnySubnet, CreateReservedFixedIPTypeIPAddress, CreateReservedFixedIPTypePort:
		return true
	}
	return false
}

// Response with prices per hour and per month
type ItemPriceResponse struct {
	// Price status for the UI
	PriceStatus PriceDisplayStatus `json:"price_status,required"`
	// Currency code (3 letter code per ISO 4217)
	CurrencyCode ItemPriceResponseCurrencyCode `json:"currency_code,nullable"`
	// Actual discount relative value
	DiscountPercent ItemPriceResponseDiscountPercentUnion `json:"discount_percent,nullable"`
	// Price of the item charged per hour
	PricePerHour ItemPriceResponsePricePerHourUnion `json:"price_per_hour,nullable"`
	// Price of the item charged per month
	PricePerMonth ItemPriceResponsePricePerMonthUnion `json:"price_per_month,nullable"`
	// Total price VAT inclusive per month without discount
	PriceWithoutDiscountPerMonth ItemPriceResponsePriceWithoutDiscountPerMonthUnion `json:"price_without_discount_per_month,nullable"`
	JSON                         itemPriceResponseJSON                              `json:"-"`
}

// itemPriceResponseJSON contains the JSON metadata for the struct
// [ItemPriceResponse]
type itemPriceResponseJSON struct {
	PriceStatus                  apijson.Field
	CurrencyCode                 apijson.Field
	DiscountPercent              apijson.Field
	PricePerHour                 apijson.Field
	PricePerMonth                apijson.Field
	PriceWithoutDiscountPerMonth apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ItemPriceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemPriceResponseJSON) RawJSON() string {
	return r.raw
}

// Currency code (3 letter code per ISO 4217)
type ItemPriceResponseCurrencyCode string

const (
	ItemPriceResponseCurrencyCodeAed ItemPriceResponseCurrencyCode = "AED"
	ItemPriceResponseCurrencyCodeAfn ItemPriceResponseCurrencyCode = "AFN"
	ItemPriceResponseCurrencyCodeAll ItemPriceResponseCurrencyCode = "ALL"
	ItemPriceResponseCurrencyCodeAmd ItemPriceResponseCurrencyCode = "AMD"
	ItemPriceResponseCurrencyCodeAng ItemPriceResponseCurrencyCode = "ANG"
	ItemPriceResponseCurrencyCodeAoa ItemPriceResponseCurrencyCode = "AOA"
	ItemPriceResponseCurrencyCodeArs ItemPriceResponseCurrencyCode = "ARS"
	ItemPriceResponseCurrencyCodeAud ItemPriceResponseCurrencyCode = "AUD"
	ItemPriceResponseCurrencyCodeAwg ItemPriceResponseCurrencyCode = "AWG"
	ItemPriceResponseCurrencyCodeAzn ItemPriceResponseCurrencyCode = "AZN"
	ItemPriceResponseCurrencyCodeBam ItemPriceResponseCurrencyCode = "BAM"
	ItemPriceResponseCurrencyCodeBbd ItemPriceResponseCurrencyCode = "BBD"
	ItemPriceResponseCurrencyCodeBdt ItemPriceResponseCurrencyCode = "BDT"
	ItemPriceResponseCurrencyCodeBgn ItemPriceResponseCurrencyCode = "BGN"
	ItemPriceResponseCurrencyCodeBhd ItemPriceResponseCurrencyCode = "BHD"
	ItemPriceResponseCurrencyCodeBif ItemPriceResponseCurrencyCode = "BIF"
	ItemPriceResponseCurrencyCodeBmd ItemPriceResponseCurrencyCode = "BMD"
	ItemPriceResponseCurrencyCodeBnd ItemPriceResponseCurrencyCode = "BND"
	ItemPriceResponseCurrencyCodeBob ItemPriceResponseCurrencyCode = "BOB"
	ItemPriceResponseCurrencyCodeBrl ItemPriceResponseCurrencyCode = "BRL"
	ItemPriceResponseCurrencyCodeBsd ItemPriceResponseCurrencyCode = "BSD"
	ItemPriceResponseCurrencyCodeBtn ItemPriceResponseCurrencyCode = "BTN"
	ItemPriceResponseCurrencyCodeBwp ItemPriceResponseCurrencyCode = "BWP"
	ItemPriceResponseCurrencyCodeByn ItemPriceResponseCurrencyCode = "BYN"
	ItemPriceResponseCurrencyCodeBzd ItemPriceResponseCurrencyCode = "BZD"
	ItemPriceResponseCurrencyCodeCad ItemPriceResponseCurrencyCode = "CAD"
	ItemPriceResponseCurrencyCodeCdf ItemPriceResponseCurrencyCode = "CDF"
	ItemPriceResponseCurrencyCodeChf ItemPriceResponseCurrencyCode = "CHF"
	ItemPriceResponseCurrencyCodeClp ItemPriceResponseCurrencyCode = "CLP"
	ItemPriceResponseCurrencyCodeCny ItemPriceResponseCurrencyCode = "CNY"
	ItemPriceResponseCurrencyCodeCop ItemPriceResponseCurrencyCode = "COP"
	ItemPriceResponseCurrencyCodeCrc ItemPriceResponseCurrencyCode = "CRC"
	ItemPriceResponseCurrencyCodeCuc ItemPriceResponseCurrencyCode = "CUC"
	ItemPriceResponseCurrencyCodeCup ItemPriceResponseCurrencyCode = "CUP"
	ItemPriceResponseCurrencyCodeCve ItemPriceResponseCurrencyCode = "CVE"
	ItemPriceResponseCurrencyCodeCzk ItemPriceResponseCurrencyCode = "CZK"
	ItemPriceResponseCurrencyCodeDjf ItemPriceResponseCurrencyCode = "DJF"
	ItemPriceResponseCurrencyCodeDkk ItemPriceResponseCurrencyCode = "DKK"
	ItemPriceResponseCurrencyCodeDop ItemPriceResponseCurrencyCode = "DOP"
	ItemPriceResponseCurrencyCodeDzd ItemPriceResponseCurrencyCode = "DZD"
	ItemPriceResponseCurrencyCodeEgp ItemPriceResponseCurrencyCode = "EGP"
	ItemPriceResponseCurrencyCodeErn ItemPriceResponseCurrencyCode = "ERN"
	ItemPriceResponseCurrencyCodeEtb ItemPriceResponseCurrencyCode = "ETB"
	ItemPriceResponseCurrencyCodeEur ItemPriceResponseCurrencyCode = "EUR"
	ItemPriceResponseCurrencyCodeFjd ItemPriceResponseCurrencyCode = "FJD"
	ItemPriceResponseCurrencyCodeFkp ItemPriceResponseCurrencyCode = "FKP"
	ItemPriceResponseCurrencyCodeGbp ItemPriceResponseCurrencyCode = "GBP"
	ItemPriceResponseCurrencyCodeGel ItemPriceResponseCurrencyCode = "GEL"
	ItemPriceResponseCurrencyCodeGhs ItemPriceResponseCurrencyCode = "GHS"
	ItemPriceResponseCurrencyCodeGip ItemPriceResponseCurrencyCode = "GIP"
	ItemPriceResponseCurrencyCodeGmd ItemPriceResponseCurrencyCode = "GMD"
	ItemPriceResponseCurrencyCodeGnf ItemPriceResponseCurrencyCode = "GNF"
	ItemPriceResponseCurrencyCodeGtq ItemPriceResponseCurrencyCode = "GTQ"
	ItemPriceResponseCurrencyCodeGyd ItemPriceResponseCurrencyCode = "GYD"
	ItemPriceResponseCurrencyCodeHkd ItemPriceResponseCurrencyCode = "HKD"
	ItemPriceResponseCurrencyCodeHnl ItemPriceResponseCurrencyCode = "HNL"
	ItemPriceResponseCurrencyCodeHrk ItemPriceResponseCurrencyCode = "HRK"
	ItemPriceResponseCurrencyCodeHtg ItemPriceResponseCurrencyCode = "HTG"
	ItemPriceResponseCurrencyCodeHuf ItemPriceResponseCurrencyCode = "HUF"
	ItemPriceResponseCurrencyCodeIdr ItemPriceResponseCurrencyCode = "IDR"
	ItemPriceResponseCurrencyCodeIls ItemPriceResponseCurrencyCode = "ILS"
	ItemPriceResponseCurrencyCodeInr ItemPriceResponseCurrencyCode = "INR"
	ItemPriceResponseCurrencyCodeIqd ItemPriceResponseCurrencyCode = "IQD"
	ItemPriceResponseCurrencyCodeIrr ItemPriceResponseCurrencyCode = "IRR"
	ItemPriceResponseCurrencyCodeIsk ItemPriceResponseCurrencyCode = "ISK"
	ItemPriceResponseCurrencyCodeJmd ItemPriceResponseCurrencyCode = "JMD"
	ItemPriceResponseCurrencyCodeJod ItemPriceResponseCurrencyCode = "JOD"
	ItemPriceResponseCurrencyCodeJpy ItemPriceResponseCurrencyCode = "JPY"
	ItemPriceResponseCurrencyCodeKes ItemPriceResponseCurrencyCode = "KES"
	ItemPriceResponseCurrencyCodeKgs ItemPriceResponseCurrencyCode = "KGS"
	ItemPriceResponseCurrencyCodeKhr ItemPriceResponseCurrencyCode = "KHR"
	ItemPriceResponseCurrencyCodeKmf ItemPriceResponseCurrencyCode = "KMF"
	ItemPriceResponseCurrencyCodeKpw ItemPriceResponseCurrencyCode = "KPW"
	ItemPriceResponseCurrencyCodeKrw ItemPriceResponseCurrencyCode = "KRW"
	ItemPriceResponseCurrencyCodeKwd ItemPriceResponseCurrencyCode = "KWD"
	ItemPriceResponseCurrencyCodeKyd ItemPriceResponseCurrencyCode = "KYD"
	ItemPriceResponseCurrencyCodeKzt ItemPriceResponseCurrencyCode = "KZT"
	ItemPriceResponseCurrencyCodeLak ItemPriceResponseCurrencyCode = "LAK"
	ItemPriceResponseCurrencyCodeLbp ItemPriceResponseCurrencyCode = "LBP"
	ItemPriceResponseCurrencyCodeLkr ItemPriceResponseCurrencyCode = "LKR"
	ItemPriceResponseCurrencyCodeLrd ItemPriceResponseCurrencyCode = "LRD"
	ItemPriceResponseCurrencyCodeLsl ItemPriceResponseCurrencyCode = "LSL"
	ItemPriceResponseCurrencyCodeLyd ItemPriceResponseCurrencyCode = "LYD"
	ItemPriceResponseCurrencyCodeMad ItemPriceResponseCurrencyCode = "MAD"
	ItemPriceResponseCurrencyCodeMdl ItemPriceResponseCurrencyCode = "MDL"
	ItemPriceResponseCurrencyCodeMga ItemPriceResponseCurrencyCode = "MGA"
	ItemPriceResponseCurrencyCodeMkd ItemPriceResponseCurrencyCode = "MKD"
	ItemPriceResponseCurrencyCodeMmk ItemPriceResponseCurrencyCode = "MMK"
	ItemPriceResponseCurrencyCodeMnt ItemPriceResponseCurrencyCode = "MNT"
	ItemPriceResponseCurrencyCodeMop ItemPriceResponseCurrencyCode = "MOP"
	ItemPriceResponseCurrencyCodeMro ItemPriceResponseCurrencyCode = "MRO"
	ItemPriceResponseCurrencyCodeMur ItemPriceResponseCurrencyCode = "MUR"
	ItemPriceResponseCurrencyCodeMvr ItemPriceResponseCurrencyCode = "MVR"
	ItemPriceResponseCurrencyCodeMwk ItemPriceResponseCurrencyCode = "MWK"
	ItemPriceResponseCurrencyCodeMxn ItemPriceResponseCurrencyCode = "MXN"
	ItemPriceResponseCurrencyCodeMyr ItemPriceResponseCurrencyCode = "MYR"
	ItemPriceResponseCurrencyCodeMzn ItemPriceResponseCurrencyCode = "MZN"
	ItemPriceResponseCurrencyCodeNad ItemPriceResponseCurrencyCode = "NAD"
	ItemPriceResponseCurrencyCodeNgn ItemPriceResponseCurrencyCode = "NGN"
	ItemPriceResponseCurrencyCodeNio ItemPriceResponseCurrencyCode = "NIO"
	ItemPriceResponseCurrencyCodeNok ItemPriceResponseCurrencyCode = "NOK"
	ItemPriceResponseCurrencyCodeNpr ItemPriceResponseCurrencyCode = "NPR"
	ItemPriceResponseCurrencyCodeNzd ItemPriceResponseCurrencyCode = "NZD"
	ItemPriceResponseCurrencyCodeOmr ItemPriceResponseCurrencyCode = "OMR"
	ItemPriceResponseCurrencyCodePab ItemPriceResponseCurrencyCode = "PAB"
	ItemPriceResponseCurrencyCodePen ItemPriceResponseCurrencyCode = "PEN"
	ItemPriceResponseCurrencyCodePgk ItemPriceResponseCurrencyCode = "PGK"
	ItemPriceResponseCurrencyCodePhp ItemPriceResponseCurrencyCode = "PHP"
	ItemPriceResponseCurrencyCodePkr ItemPriceResponseCurrencyCode = "PKR"
	ItemPriceResponseCurrencyCodePln ItemPriceResponseCurrencyCode = "PLN"
	ItemPriceResponseCurrencyCodePyg ItemPriceResponseCurrencyCode = "PYG"
	ItemPriceResponseCurrencyCodeQar ItemPriceResponseCurrencyCode = "QAR"
	ItemPriceResponseCurrencyCodeRon ItemPriceResponseCurrencyCode = "RON"
	ItemPriceResponseCurrencyCodeRsd ItemPriceResponseCurrencyCode = "RSD"
	ItemPriceResponseCurrencyCodeRub ItemPriceResponseCurrencyCode = "RUB"
	ItemPriceResponseCurrencyCodeRwf ItemPriceResponseCurrencyCode = "RWF"
	ItemPriceResponseCurrencyCodeSar ItemPriceResponseCurrencyCode = "SAR"
	ItemPriceResponseCurrencyCodeSbd ItemPriceResponseCurrencyCode = "SBD"
	ItemPriceResponseCurrencyCodeScr ItemPriceResponseCurrencyCode = "SCR"
	ItemPriceResponseCurrencyCodeSdg ItemPriceResponseCurrencyCode = "SDG"
	ItemPriceResponseCurrencyCodeSek ItemPriceResponseCurrencyCode = "SEK"
	ItemPriceResponseCurrencyCodeSgd ItemPriceResponseCurrencyCode = "SGD"
	ItemPriceResponseCurrencyCodeShp ItemPriceResponseCurrencyCode = "SHP"
	ItemPriceResponseCurrencyCodeSll ItemPriceResponseCurrencyCode = "SLL"
	ItemPriceResponseCurrencyCodeSos ItemPriceResponseCurrencyCode = "SOS"
	ItemPriceResponseCurrencyCodeSrd ItemPriceResponseCurrencyCode = "SRD"
	ItemPriceResponseCurrencyCodeSsp ItemPriceResponseCurrencyCode = "SSP"
	ItemPriceResponseCurrencyCodeStd ItemPriceResponseCurrencyCode = "STD"
	ItemPriceResponseCurrencyCodeSvc ItemPriceResponseCurrencyCode = "SVC"
	ItemPriceResponseCurrencyCodeSyp ItemPriceResponseCurrencyCode = "SYP"
	ItemPriceResponseCurrencyCodeSzl ItemPriceResponseCurrencyCode = "SZL"
	ItemPriceResponseCurrencyCodeThb ItemPriceResponseCurrencyCode = "THB"
	ItemPriceResponseCurrencyCodeTjs ItemPriceResponseCurrencyCode = "TJS"
	ItemPriceResponseCurrencyCodeTmt ItemPriceResponseCurrencyCode = "TMT"
	ItemPriceResponseCurrencyCodeTnd ItemPriceResponseCurrencyCode = "TND"
	ItemPriceResponseCurrencyCodeTop ItemPriceResponseCurrencyCode = "TOP"
	ItemPriceResponseCurrencyCodeTry ItemPriceResponseCurrencyCode = "TRY"
	ItemPriceResponseCurrencyCodeTtd ItemPriceResponseCurrencyCode = "TTD"
	ItemPriceResponseCurrencyCodeTwd ItemPriceResponseCurrencyCode = "TWD"
	ItemPriceResponseCurrencyCodeTzs ItemPriceResponseCurrencyCode = "TZS"
	ItemPriceResponseCurrencyCodeUah ItemPriceResponseCurrencyCode = "UAH"
	ItemPriceResponseCurrencyCodeUgx ItemPriceResponseCurrencyCode = "UGX"
	ItemPriceResponseCurrencyCodeUsd ItemPriceResponseCurrencyCode = "USD"
	ItemPriceResponseCurrencyCodeUyu ItemPriceResponseCurrencyCode = "UYU"
	ItemPriceResponseCurrencyCodeUzs ItemPriceResponseCurrencyCode = "UZS"
	ItemPriceResponseCurrencyCodeVef ItemPriceResponseCurrencyCode = "VEF"
	ItemPriceResponseCurrencyCodeVnd ItemPriceResponseCurrencyCode = "VND"
	ItemPriceResponseCurrencyCodeVuv ItemPriceResponseCurrencyCode = "VUV"
	ItemPriceResponseCurrencyCodeWst ItemPriceResponseCurrencyCode = "WST"
	ItemPriceResponseCurrencyCodeXaf ItemPriceResponseCurrencyCode = "XAF"
	ItemPriceResponseCurrencyCodeXag ItemPriceResponseCurrencyCode = "XAG"
	ItemPriceResponseCurrencyCodeXau ItemPriceResponseCurrencyCode = "XAU"
	ItemPriceResponseCurrencyCodeXba ItemPriceResponseCurrencyCode = "XBA"
	ItemPriceResponseCurrencyCodeXbb ItemPriceResponseCurrencyCode = "XBB"
	ItemPriceResponseCurrencyCodeXbc ItemPriceResponseCurrencyCode = "XBC"
	ItemPriceResponseCurrencyCodeXbd ItemPriceResponseCurrencyCode = "XBD"
	ItemPriceResponseCurrencyCodeXcd ItemPriceResponseCurrencyCode = "XCD"
	ItemPriceResponseCurrencyCodeXdr ItemPriceResponseCurrencyCode = "XDR"
	ItemPriceResponseCurrencyCodeXof ItemPriceResponseCurrencyCode = "XOF"
	ItemPriceResponseCurrencyCodeXpd ItemPriceResponseCurrencyCode = "XPD"
	ItemPriceResponseCurrencyCodeXpf ItemPriceResponseCurrencyCode = "XPF"
	ItemPriceResponseCurrencyCodeXpt ItemPriceResponseCurrencyCode = "XPT"
	ItemPriceResponseCurrencyCodeXsu ItemPriceResponseCurrencyCode = "XSU"
	ItemPriceResponseCurrencyCodeXts ItemPriceResponseCurrencyCode = "XTS"
	ItemPriceResponseCurrencyCodeXua ItemPriceResponseCurrencyCode = "XUA"
	ItemPriceResponseCurrencyCodeXxx ItemPriceResponseCurrencyCode = "XXX"
	ItemPriceResponseCurrencyCodeYer ItemPriceResponseCurrencyCode = "YER"
	ItemPriceResponseCurrencyCodeZar ItemPriceResponseCurrencyCode = "ZAR"
	ItemPriceResponseCurrencyCodeZmw ItemPriceResponseCurrencyCode = "ZMW"
	ItemPriceResponseCurrencyCodeZwl ItemPriceResponseCurrencyCode = "ZWL"
)

func (r ItemPriceResponseCurrencyCode) IsKnown() bool {
	switch r {
	case ItemPriceResponseCurrencyCodeAed, ItemPriceResponseCurrencyCodeAfn, ItemPriceResponseCurrencyCodeAll, ItemPriceResponseCurrencyCodeAmd, ItemPriceResponseCurrencyCodeAng, ItemPriceResponseCurrencyCodeAoa, ItemPriceResponseCurrencyCodeArs, ItemPriceResponseCurrencyCodeAud, ItemPriceResponseCurrencyCodeAwg, ItemPriceResponseCurrencyCodeAzn, ItemPriceResponseCurrencyCodeBam, ItemPriceResponseCurrencyCodeBbd, ItemPriceResponseCurrencyCodeBdt, ItemPriceResponseCurrencyCodeBgn, ItemPriceResponseCurrencyCodeBhd, ItemPriceResponseCurrencyCodeBif, ItemPriceResponseCurrencyCodeBmd, ItemPriceResponseCurrencyCodeBnd, ItemPriceResponseCurrencyCodeBob, ItemPriceResponseCurrencyCodeBrl, ItemPriceResponseCurrencyCodeBsd, ItemPriceResponseCurrencyCodeBtn, ItemPriceResponseCurrencyCodeBwp, ItemPriceResponseCurrencyCodeByn, ItemPriceResponseCurrencyCodeBzd, ItemPriceResponseCurrencyCodeCad, ItemPriceResponseCurrencyCodeCdf, ItemPriceResponseCurrencyCodeChf, ItemPriceResponseCurrencyCodeClp, ItemPriceResponseCurrencyCodeCny, ItemPriceResponseCurrencyCodeCop, ItemPriceResponseCurrencyCodeCrc, ItemPriceResponseCurrencyCodeCuc, ItemPriceResponseCurrencyCodeCup, ItemPriceResponseCurrencyCodeCve, ItemPriceResponseCurrencyCodeCzk, ItemPriceResponseCurrencyCodeDjf, ItemPriceResponseCurrencyCodeDkk, ItemPriceResponseCurrencyCodeDop, ItemPriceResponseCurrencyCodeDzd, ItemPriceResponseCurrencyCodeEgp, ItemPriceResponseCurrencyCodeErn, ItemPriceResponseCurrencyCodeEtb, ItemPriceResponseCurrencyCodeEur, ItemPriceResponseCurrencyCodeFjd, ItemPriceResponseCurrencyCodeFkp, ItemPriceResponseCurrencyCodeGbp, ItemPriceResponseCurrencyCodeGel, ItemPriceResponseCurrencyCodeGhs, ItemPriceResponseCurrencyCodeGip, ItemPriceResponseCurrencyCodeGmd, ItemPriceResponseCurrencyCodeGnf, ItemPriceResponseCurrencyCodeGtq, ItemPriceResponseCurrencyCodeGyd, ItemPriceResponseCurrencyCodeHkd, ItemPriceResponseCurrencyCodeHnl, ItemPriceResponseCurrencyCodeHrk, ItemPriceResponseCurrencyCodeHtg, ItemPriceResponseCurrencyCodeHuf, ItemPriceResponseCurrencyCodeIdr, ItemPriceResponseCurrencyCodeIls, ItemPriceResponseCurrencyCodeInr, ItemPriceResponseCurrencyCodeIqd, ItemPriceResponseCurrencyCodeIrr, ItemPriceResponseCurrencyCodeIsk, ItemPriceResponseCurrencyCodeJmd, ItemPriceResponseCurrencyCodeJod, ItemPriceResponseCurrencyCodeJpy, ItemPriceResponseCurrencyCodeKes, ItemPriceResponseCurrencyCodeKgs, ItemPriceResponseCurrencyCodeKhr, ItemPriceResponseCurrencyCodeKmf, ItemPriceResponseCurrencyCodeKpw, ItemPriceResponseCurrencyCodeKrw, ItemPriceResponseCurrencyCodeKwd, ItemPriceResponseCurrencyCodeKyd, ItemPriceResponseCurrencyCodeKzt, ItemPriceResponseCurrencyCodeLak, ItemPriceResponseCurrencyCodeLbp, ItemPriceResponseCurrencyCodeLkr, ItemPriceResponseCurrencyCodeLrd, ItemPriceResponseCurrencyCodeLsl, ItemPriceResponseCurrencyCodeLyd, ItemPriceResponseCurrencyCodeMad, ItemPriceResponseCurrencyCodeMdl, ItemPriceResponseCurrencyCodeMga, ItemPriceResponseCurrencyCodeMkd, ItemPriceResponseCurrencyCodeMmk, ItemPriceResponseCurrencyCodeMnt, ItemPriceResponseCurrencyCodeMop, ItemPriceResponseCurrencyCodeMro, ItemPriceResponseCurrencyCodeMur, ItemPriceResponseCurrencyCodeMvr, ItemPriceResponseCurrencyCodeMwk, ItemPriceResponseCurrencyCodeMxn, ItemPriceResponseCurrencyCodeMyr, ItemPriceResponseCurrencyCodeMzn, ItemPriceResponseCurrencyCodeNad, ItemPriceResponseCurrencyCodeNgn, ItemPriceResponseCurrencyCodeNio, ItemPriceResponseCurrencyCodeNok, ItemPriceResponseCurrencyCodeNpr, ItemPriceResponseCurrencyCodeNzd, ItemPriceResponseCurrencyCodeOmr, ItemPriceResponseCurrencyCodePab, ItemPriceResponseCurrencyCodePen, ItemPriceResponseCurrencyCodePgk, ItemPriceResponseCurrencyCodePhp, ItemPriceResponseCurrencyCodePkr, ItemPriceResponseCurrencyCodePln, ItemPriceResponseCurrencyCodePyg, ItemPriceResponseCurrencyCodeQar, ItemPriceResponseCurrencyCodeRon, ItemPriceResponseCurrencyCodeRsd, ItemPriceResponseCurrencyCodeRub, ItemPriceResponseCurrencyCodeRwf, ItemPriceResponseCurrencyCodeSar, ItemPriceResponseCurrencyCodeSbd, ItemPriceResponseCurrencyCodeScr, ItemPriceResponseCurrencyCodeSdg, ItemPriceResponseCurrencyCodeSek, ItemPriceResponseCurrencyCodeSgd, ItemPriceResponseCurrencyCodeShp, ItemPriceResponseCurrencyCodeSll, ItemPriceResponseCurrencyCodeSos, ItemPriceResponseCurrencyCodeSrd, ItemPriceResponseCurrencyCodeSsp, ItemPriceResponseCurrencyCodeStd, ItemPriceResponseCurrencyCodeSvc, ItemPriceResponseCurrencyCodeSyp, ItemPriceResponseCurrencyCodeSzl, ItemPriceResponseCurrencyCodeThb, ItemPriceResponseCurrencyCodeTjs, ItemPriceResponseCurrencyCodeTmt, ItemPriceResponseCurrencyCodeTnd, ItemPriceResponseCurrencyCodeTop, ItemPriceResponseCurrencyCodeTry, ItemPriceResponseCurrencyCodeTtd, ItemPriceResponseCurrencyCodeTwd, ItemPriceResponseCurrencyCodeTzs, ItemPriceResponseCurrencyCodeUah, ItemPriceResponseCurrencyCodeUgx, ItemPriceResponseCurrencyCodeUsd, ItemPriceResponseCurrencyCodeUyu, ItemPriceResponseCurrencyCodeUzs, ItemPriceResponseCurrencyCodeVef, ItemPriceResponseCurrencyCodeVnd, ItemPriceResponseCurrencyCodeVuv, ItemPriceResponseCurrencyCodeWst, ItemPriceResponseCurrencyCodeXaf, ItemPriceResponseCurrencyCodeXag, ItemPriceResponseCurrencyCodeXau, ItemPriceResponseCurrencyCodeXba, ItemPriceResponseCurrencyCodeXbb, ItemPriceResponseCurrencyCodeXbc, ItemPriceResponseCurrencyCodeXbd, ItemPriceResponseCurrencyCodeXcd, ItemPriceResponseCurrencyCodeXdr, ItemPriceResponseCurrencyCodeXof, ItemPriceResponseCurrencyCodeXpd, ItemPriceResponseCurrencyCodeXpf, ItemPriceResponseCurrencyCodeXpt, ItemPriceResponseCurrencyCodeXsu, ItemPriceResponseCurrencyCodeXts, ItemPriceResponseCurrencyCodeXua, ItemPriceResponseCurrencyCodeXxx, ItemPriceResponseCurrencyCodeYer, ItemPriceResponseCurrencyCodeZar, ItemPriceResponseCurrencyCodeZmw, ItemPriceResponseCurrencyCodeZwl:
		return true
	}
	return false
}

// Actual discount relative value
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type ItemPriceResponseDiscountPercentUnion interface {
	ImplementsItemPriceResponseDiscountPercentUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ItemPriceResponseDiscountPercentUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Price of the item charged per hour
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type ItemPriceResponsePricePerHourUnion interface {
	ImplementsItemPriceResponsePricePerHourUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ItemPriceResponsePricePerHourUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Price of the item charged per month
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type ItemPriceResponsePricePerMonthUnion interface {
	ImplementsItemPriceResponsePricePerMonthUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ItemPriceResponsePricePerMonthUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Total price VAT inclusive per month without discount
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type ItemPriceResponsePriceWithoutDiscountPerMonthUnion interface {
	ImplementsItemPriceResponsePriceWithoutDiscountPerMonthUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ItemPriceResponsePriceWithoutDiscountPerMonthUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PriceDisplayStatus string

const (
	PriceDisplayStatusError PriceDisplayStatus = "error"
	PriceDisplayStatusHide  PriceDisplayStatus = "hide"
	PriceDisplayStatusShow  PriceDisplayStatus = "show"
)

func (r PriceDisplayStatus) IsKnown() bool {
	switch r {
	case PriceDisplayStatusError, PriceDisplayStatusHide, PriceDisplayStatusShow:
		return true
	}
	return false
}

type VolumeTypeName string

const (
	VolumeTypeNameCold          VolumeTypeName = "cold"
	VolumeTypeNameSsdHiiops     VolumeTypeName = "ssd_hiiops"
	VolumeTypeNameSsdLocal      VolumeTypeName = "ssd_local"
	VolumeTypeNameSsdLowlatency VolumeTypeName = "ssd_lowlatency"
	VolumeTypeNameStandard      VolumeTypeName = "standard"
	VolumeTypeNameUltra         VolumeTypeName = "ultra"
)

func (r VolumeTypeName) IsKnown() bool {
	switch r {
	case VolumeTypeNameCold, VolumeTypeNameSsdHiiops, VolumeTypeNameSsdLocal, VolumeTypeNameSsdLowlatency, VolumeTypeNameStandard, VolumeTypeNameUltra:
		return true
	}
	return false
}

type CloudV1PricingGetBillingReservationPricesResponse struct {
	// Description using period for calc slices [`year`, `month`, `day`]
	ActivityPeriod string `json:"activity_period,required"`
	// Length of the full reservation period by `activity_period`
	ActivityPeriodLength int64                                                       `json:"activity_period_length,required"`
	AmountPrices         BillingReservationAmountPrices                              `json:"amount_prices,required"`
	PriceStatus          PriceDisplayStatus                                          `json:"price_status,required"`
	RegionID             int64                                                       `json:"region_id,required"`
	RegionName           string                                                      `json:"region_name,required"`
	Resources            []CloudV1PricingGetBillingReservationPricesResponseResource `json:"resources,required"`
	JSON                 cloudV1PricingGetBillingReservationPricesResponseJSON       `json:"-"`
}

// cloudV1PricingGetBillingReservationPricesResponseJSON contains the JSON metadata
// for the struct [CloudV1PricingGetBillingReservationPricesResponse]
type cloudV1PricingGetBillingReservationPricesResponseJSON struct {
	ActivityPeriod       apijson.Field
	ActivityPeriodLength apijson.Field
	AmountPrices         apijson.Field
	PriceStatus          apijson.Field
	RegionID             apijson.Field
	RegionName           apijson.Field
	Resources            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CloudV1PricingGetBillingReservationPricesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1PricingGetBillingReservationPricesResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1PricingGetBillingReservationPricesResponseResource struct {
	// Name of the billing period, e.g month
	ActivityPeriod string `json:"activity_period,required"`
	// Length of the full reservation period by `activity_period`
	ActivityPeriodLength int64 `json:"activity_period_length,required"`
	// Commit price of the item charged per month
	CommitPricePerMonth CloudV1PricingGetBillingReservationPricesResponseResourcesCommitPricePerMonthUnion `json:"commit_price_per_month,required"`
	// Commit price of the item charged per hour
	CommitPricePerUnit CloudV1PricingGetBillingReservationPricesResponseResourcesCommitPricePerUnitUnion `json:"commit_price_per_unit,required"`
	// Commit price of the item charged for all period reservation
	CommitPriceTotal CloudV1PricingGetBillingReservationPricesResponseResourcesCommitPriceTotalUnion `json:"commit_price_total,required"`
	// Overcommit price of the item charged per month
	OvercommitPricePerMonth CloudV1PricingGetBillingReservationPricesResponseResourcesOvercommitPricePerMonthUnion `json:"overcommit_price_per_month,required"`
	// Overcommit price of the item charged per hour
	OvercommitPricePerUnit CloudV1PricingGetBillingReservationPricesResponseResourcesOvercommitPricePerUnitUnion `json:"overcommit_price_per_unit,required"`
	// Overcommit price of the item charged for all period reservation
	OvercommitPriceTotal CloudV1PricingGetBillingReservationPricesResponseResourcesOvercommitPriceTotalUnion `json:"overcommit_price_total,required"`
	// Resource name
	ResourceName string `json:"resource_name,required"`
	// Billing unit name
	UnitName string `json:"unit_name,required"`
	// Minimal billing size, for example it is 744 hours per 1 month
	UnitSizeMonth CloudV1PricingGetBillingReservationPricesResponseResourcesUnitSizeMonthUnion `json:"unit_size_month,required"`
	// Unit size month multiplied by count of resources in the reservation
	UnitSizeTotal CloudV1PricingGetBillingReservationPricesResponseResourcesUnitSizeTotalUnion `json:"unit_size_total,required"`
	JSON          cloudV1PricingGetBillingReservationPricesResponseResourceJSON                `json:"-"`
}

// cloudV1PricingGetBillingReservationPricesResponseResourceJSON contains the JSON
// metadata for the struct
// [CloudV1PricingGetBillingReservationPricesResponseResource]
type cloudV1PricingGetBillingReservationPricesResponseResourceJSON struct {
	ActivityPeriod          apijson.Field
	ActivityPeriodLength    apijson.Field
	CommitPricePerMonth     apijson.Field
	CommitPricePerUnit      apijson.Field
	CommitPriceTotal        apijson.Field
	OvercommitPricePerMonth apijson.Field
	OvercommitPricePerUnit  apijson.Field
	OvercommitPriceTotal    apijson.Field
	ResourceName            apijson.Field
	UnitName                apijson.Field
	UnitSizeMonth           apijson.Field
	UnitSizeTotal           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CloudV1PricingGetBillingReservationPricesResponseResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1PricingGetBillingReservationPricesResponseResourceJSON) RawJSON() string {
	return r.raw
}

// Commit price of the item charged per month
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1PricingGetBillingReservationPricesResponseResourcesCommitPricePerMonthUnion interface {
	ImplementsCloudV1PricingGetBillingReservationPricesResponseResourcesCommitPricePerMonthUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1PricingGetBillingReservationPricesResponseResourcesCommitPricePerMonthUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Commit price of the item charged per hour
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1PricingGetBillingReservationPricesResponseResourcesCommitPricePerUnitUnion interface {
	ImplementsCloudV1PricingGetBillingReservationPricesResponseResourcesCommitPricePerUnitUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1PricingGetBillingReservationPricesResponseResourcesCommitPricePerUnitUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Commit price of the item charged for all period reservation
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1PricingGetBillingReservationPricesResponseResourcesCommitPriceTotalUnion interface {
	ImplementsCloudV1PricingGetBillingReservationPricesResponseResourcesCommitPriceTotalUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1PricingGetBillingReservationPricesResponseResourcesCommitPriceTotalUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Overcommit price of the item charged per month
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1PricingGetBillingReservationPricesResponseResourcesOvercommitPricePerMonthUnion interface {
	ImplementsCloudV1PricingGetBillingReservationPricesResponseResourcesOvercommitPricePerMonthUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1PricingGetBillingReservationPricesResponseResourcesOvercommitPricePerMonthUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Overcommit price of the item charged per hour
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1PricingGetBillingReservationPricesResponseResourcesOvercommitPricePerUnitUnion interface {
	ImplementsCloudV1PricingGetBillingReservationPricesResponseResourcesOvercommitPricePerUnitUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1PricingGetBillingReservationPricesResponseResourcesOvercommitPricePerUnitUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Overcommit price of the item charged for all period reservation
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1PricingGetBillingReservationPricesResponseResourcesOvercommitPriceTotalUnion interface {
	ImplementsCloudV1PricingGetBillingReservationPricesResponseResourcesOvercommitPriceTotalUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1PricingGetBillingReservationPricesResponseResourcesOvercommitPriceTotalUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Minimal billing size, for example it is 744 hours per 1 month
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1PricingGetBillingReservationPricesResponseResourcesUnitSizeMonthUnion interface {
	ImplementsCloudV1PricingGetBillingReservationPricesResponseResourcesUnitSizeMonthUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1PricingGetBillingReservationPricesResponseResourcesUnitSizeMonthUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Unit size month multiplied by count of resources in the reservation
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1PricingGetBillingReservationPricesResponseResourcesUnitSizeTotalUnion interface {
	ImplementsCloudV1PricingGetBillingReservationPricesResponseResourcesUnitSizeTotalUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1PricingGetBillingReservationPricesResponseResourcesUnitSizeTotalUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1PricingGetDdosPriceResponse struct {
	// Number of days in the current month
	CurrentMonthDays int64 `json:"current_month_days,required"`
	// Number of days left in the current month
	CurrentMonthDaysLast int64 `json:"current_month_days_last,required"`
	// Number of days in the next month
	NextMonthDays int64 `json:"next_month_days,required"`
	// DDoS protection price details
	Price CloudV1PricingGetDdosPriceResponsePrice `json:"price,required"`
	JSON  cloudV1PricingGetDdosPriceResponseJSON  `json:"-"`
}

// cloudV1PricingGetDdosPriceResponseJSON contains the JSON metadata for the struct
// [CloudV1PricingGetDdosPriceResponse]
type cloudV1PricingGetDdosPriceResponseJSON struct {
	CurrentMonthDays     apijson.Field
	CurrentMonthDaysLast apijson.Field
	NextMonthDays        apijson.Field
	Price                apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CloudV1PricingGetDdosPriceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1PricingGetDdosPriceResponseJSON) RawJSON() string {
	return r.raw
}

// DDoS protection price details
type CloudV1PricingGetDdosPriceResponsePrice struct {
	// Name of the billing feature
	FeatureName string `json:"feature_name,required"`
	// Period of the price
	Period string `json:"period,required"`
	// Currency code (3 letter code per ISO 4217)
	PriceCurrency CloudV1PricingGetDdosPriceResponsePricePriceCurrency `json:"price_currency,required"`
	// Name of the unit
	UnitName string `json:"unit_name,required"`
	// Number of units
	UnitSize int64 `json:"unit_size,required"`
	// Price of the feature charged per 30 days
	PricePer30d CloudV1PricingGetDdosPriceResponsePricePricePer30dUnion `json:"price_per_30d,nullable"`
	// Price of the feature charged per day
	PricePerDay CloudV1PricingGetDdosPriceResponsePricePricePerDayUnion `json:"price_per_day,nullable"`
	// Price of the feature charged per hour
	PricePerHour CloudV1PricingGetDdosPriceResponsePricePricePerHourUnion `json:"price_per_hour,nullable"`
	JSON         cloudV1PricingGetDdosPriceResponsePriceJSON              `json:"-"`
}

// cloudV1PricingGetDdosPriceResponsePriceJSON contains the JSON metadata for the
// struct [CloudV1PricingGetDdosPriceResponsePrice]
type cloudV1PricingGetDdosPriceResponsePriceJSON struct {
	FeatureName   apijson.Field
	Period        apijson.Field
	PriceCurrency apijson.Field
	UnitName      apijson.Field
	UnitSize      apijson.Field
	PricePer30d   apijson.Field
	PricePerDay   apijson.Field
	PricePerHour  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CloudV1PricingGetDdosPriceResponsePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1PricingGetDdosPriceResponsePriceJSON) RawJSON() string {
	return r.raw
}

// Currency code (3 letter code per ISO 4217)
type CloudV1PricingGetDdosPriceResponsePricePriceCurrency string

const (
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAed CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "AED"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAfn CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "AFN"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAll CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "ALL"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAmd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "AMD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAng CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "ANG"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAoa CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "AOA"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyArs CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "ARS"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAud CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "AUD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAwg CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "AWG"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAzn CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "AZN"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBam CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BAM"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBbd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BBD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBdt CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BDT"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBgn CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BGN"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBhd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BHD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBif CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BIF"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBmd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BMD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBnd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BND"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBob CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BOB"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBrl CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BRL"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBsd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BSD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBtn CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BTN"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBwp CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BWP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyByn CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BYN"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBzd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "BZD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCad CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "CAD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCdf CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "CDF"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyChf CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "CHF"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyClp CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "CLP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCny CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "CNY"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCop CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "COP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCrc CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "CRC"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCuc CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "CUC"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCup CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "CUP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCve CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "CVE"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCzk CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "CZK"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyDjf CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "DJF"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyDkk CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "DKK"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyDop CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "DOP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyDzd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "DZD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyEgp CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "EGP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyErn CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "ERN"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyEtb CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "ETB"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyEur CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "EUR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyFjd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "FJD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyFkp CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "FKP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGbp CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "GBP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGel CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "GEL"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGhs CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "GHS"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGip CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "GIP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGmd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "GMD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGnf CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "GNF"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGtq CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "GTQ"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGyd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "GYD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyHkd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "HKD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyHnl CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "HNL"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyHrk CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "HRK"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyHtg CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "HTG"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyHuf CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "HUF"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyIdr CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "IDR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyIls CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "ILS"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyInr CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "INR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyIqd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "IQD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyIrr CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "IRR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyIsk CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "ISK"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyJmd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "JMD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyJod CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "JOD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyJpy CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "JPY"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKes CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "KES"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKgs CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "KGS"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKhr CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "KHR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKmf CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "KMF"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKpw CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "KPW"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKrw CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "KRW"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKwd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "KWD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKyd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "KYD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKzt CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "KZT"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyLak CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "LAK"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyLbp CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "LBP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyLkr CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "LKR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyLrd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "LRD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyLsl CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "LSL"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyLyd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "LYD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMad CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "MAD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMdl CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "MDL"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMga CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "MGA"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMkd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "MKD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMmk CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "MMK"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMnt CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "MNT"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMop CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "MOP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMro CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "MRO"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMur CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "MUR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMvr CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "MVR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMwk CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "MWK"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMxn CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "MXN"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMyr CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "MYR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMzn CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "MZN"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyNad CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "NAD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyNgn CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "NGN"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyNio CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "NIO"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyNok CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "NOK"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyNpr CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "NPR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyNzd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "NZD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyOmr CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "OMR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyPab CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "PAB"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyPen CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "PEN"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyPgk CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "PGK"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyPhp CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "PHP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyPkr CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "PKR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyPln CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "PLN"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyPyg CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "PYG"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyQar CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "QAR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyRon CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "RON"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyRsd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "RSD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyRub CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "RUB"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyRwf CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "RWF"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencySar CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "SAR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencySbd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "SBD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyScr CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "SCR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencySdg CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "SDG"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencySek CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "SEK"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencySgd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "SGD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyShp CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "SHP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencySll CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "SLL"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencySos CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "SOS"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencySrd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "SRD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencySsp CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "SSP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyStd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "STD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencySvc CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "SVC"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencySyp CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "SYP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencySzl CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "SZL"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyThb CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "THB"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTjs CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "TJS"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTmt CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "TMT"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTnd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "TND"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTop CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "TOP"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTry CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "TRY"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTtd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "TTD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTwd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "TWD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTzs CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "TZS"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyUah CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "UAH"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyUgx CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "UGX"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyUsd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "USD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyUyu CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "UYU"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyUzs CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "UZS"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyVef CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "VEF"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyVnd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "VND"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyVuv CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "VUV"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyWst CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "WST"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXaf CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XAF"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXag CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XAG"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXau CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XAU"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXba CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XBA"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXbb CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XBB"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXbc CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XBC"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXbd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XBD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXcd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XCD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXdr CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XDR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXof CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XOF"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXpd CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XPD"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXpf CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XPF"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXpt CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XPT"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXsu CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XSU"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXts CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XTS"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXua CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XUA"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXxx CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "XXX"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyYer CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "YER"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyZar CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "ZAR"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyZmw CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "ZMW"
	CloudV1PricingGetDdosPriceResponsePricePriceCurrencyZwl CloudV1PricingGetDdosPriceResponsePricePriceCurrency = "ZWL"
)

func (r CloudV1PricingGetDdosPriceResponsePricePriceCurrency) IsKnown() bool {
	switch r {
	case CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAed, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAfn, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAll, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAmd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAng, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAoa, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyArs, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAud, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAwg, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyAzn, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBam, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBbd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBdt, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBgn, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBhd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBif, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBmd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBnd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBob, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBrl, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBsd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBtn, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBwp, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyByn, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyBzd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCad, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCdf, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyChf, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyClp, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCny, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCop, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCrc, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCuc, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCup, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCve, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyCzk, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyDjf, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyDkk, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyDop, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyDzd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyEgp, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyErn, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyEtb, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyEur, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyFjd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyFkp, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGbp, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGel, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGhs, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGip, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGmd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGnf, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGtq, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyGyd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyHkd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyHnl, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyHrk, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyHtg, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyHuf, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyIdr, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyIls, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyInr, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyIqd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyIrr, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyIsk, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyJmd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyJod, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyJpy, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKes, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKgs, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKhr, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKmf, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKpw, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKrw, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKwd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKyd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyKzt, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyLak, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyLbp, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyLkr, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyLrd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyLsl, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyLyd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMad, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMdl, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMga, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMkd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMmk, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMnt, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMop, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMro, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMur, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMvr, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMwk, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMxn, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMyr, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyMzn, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyNad, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyNgn, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyNio, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyNok, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyNpr, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyNzd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyOmr, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyPab, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyPen, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyPgk, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyPhp, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyPkr, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyPln, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyPyg, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyQar, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyRon, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyRsd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyRub, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyRwf, CloudV1PricingGetDdosPriceResponsePricePriceCurrencySar, CloudV1PricingGetDdosPriceResponsePricePriceCurrencySbd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyScr, CloudV1PricingGetDdosPriceResponsePricePriceCurrencySdg, CloudV1PricingGetDdosPriceResponsePricePriceCurrencySek, CloudV1PricingGetDdosPriceResponsePricePriceCurrencySgd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyShp, CloudV1PricingGetDdosPriceResponsePricePriceCurrencySll, CloudV1PricingGetDdosPriceResponsePricePriceCurrencySos, CloudV1PricingGetDdosPriceResponsePricePriceCurrencySrd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencySsp, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyStd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencySvc, CloudV1PricingGetDdosPriceResponsePricePriceCurrencySyp, CloudV1PricingGetDdosPriceResponsePricePriceCurrencySzl, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyThb, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTjs, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTmt, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTnd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTop, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTry, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTtd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTwd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyTzs, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyUah, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyUgx, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyUsd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyUyu, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyUzs, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyVef, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyVnd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyVuv, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyWst, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXaf, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXag, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXau, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXba, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXbb, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXbc, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXbd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXcd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXdr, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXof, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXpd, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXpf, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXpt, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXsu, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXts, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXua, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyXxx, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyYer, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyZar, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyZmw, CloudV1PricingGetDdosPriceResponsePricePriceCurrencyZwl:
		return true
	}
	return false
}

// Price of the feature charged per 30 days
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1PricingGetDdosPriceResponsePricePricePer30dUnion interface {
	ImplementsCloudV1PricingGetDdosPriceResponsePricePricePer30dUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1PricingGetDdosPriceResponsePricePricePer30dUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Price of the feature charged per day
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1PricingGetDdosPriceResponsePricePricePerDayUnion interface {
	ImplementsCloudV1PricingGetDdosPriceResponsePricePricePerDayUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1PricingGetDdosPriceResponsePricePricePerDayUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Price of the feature charged per hour
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1PricingGetDdosPriceResponsePricePricePerHourUnion interface {
	ImplementsCloudV1PricingGetDdosPriceResponsePricePricePerHourUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1PricingGetDdosPriceResponsePricePricePerHourUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CloudV1PricingGetBillingReservationPricesParams struct {
	Period     param.Field[CloudV1PricingGetBillingReservationPricesParamsPeriod] `json:"period,required"`
	Resources  param.Field[[]BillingReservationPricingResourceRequestParam]       `json:"resources,required"`
	ClientID   param.Field[int64]                                                 `json:"client_id"`
	HasWindows param.Field[bool]                                                  `json:"has_windows"`
}

func (r CloudV1PricingGetBillingReservationPricesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1PricingGetBillingReservationPricesParamsPeriod string

const (
	CloudV1PricingGetBillingReservationPricesParamsPeriodMonths12 CloudV1PricingGetBillingReservationPricesParamsPeriod = "MONTHS12"
	CloudV1PricingGetBillingReservationPricesParamsPeriodMonths36 CloudV1PricingGetBillingReservationPricesParamsPeriod = "MONTHS36"
)

func (r CloudV1PricingGetBillingReservationPricesParamsPeriod) IsKnown() bool {
	switch r {
	case CloudV1PricingGetBillingReservationPricesParamsPeriodMonths12, CloudV1PricingGetBillingReservationPricesParamsPeriodMonths36:
		return true
	}
	return false
}

type CloudV1PricingGetLifecyclePolicyCostParams struct {
	CreateLifecyclePolicy CreateLifecyclePolicyParam `json:"create_lifecycle_policy,required"`
}

func (r CloudV1PricingGetLifecyclePolicyCostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateLifecyclePolicy)
}

type CloudV1PricingPreviewContainerPriceParams struct {
	// Container flavor
	Flavor param.Field[string] `json:"flavor,required"`
	// Container autoscaling
	Scale param.Field[ScaleParam] `json:"scale,required"`
}

func (r CloudV1PricingPreviewContainerPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1PricingPreviewFileSharePriceParams struct {
	// File share size in GiB
	Size param.Field[int64] `json:"size,required"`
	// Only 'new-file-share'
	Source param.Field[CloudV1PricingPreviewFileSharePriceParamsSource] `json:"source,required"`
	// Only 'standard'
	TypeName param.Field[VolumeTypeName] `json:"type_name"`
}

func (r CloudV1PricingPreviewFileSharePriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Only 'new-file-share'
type CloudV1PricingPreviewFileSharePriceParamsSource string

const (
	CloudV1PricingPreviewFileSharePriceParamsSourceNewFileShare CloudV1PricingPreviewFileSharePriceParamsSource = "new-file-share"
)

func (r CloudV1PricingPreviewFileSharePriceParamsSource) IsKnown() bool {
	switch r {
	case CloudV1PricingPreviewFileSharePriceParamsSourceNewFileShare:
		return true
	}
	return false
}

type CloudV1PricingPreviewFloatingIPPriceParams struct {
	// Number of floating IP addresses
	Count param.Field[int64] `json:"count,required"`
}

func (r CloudV1PricingPreviewFloatingIPPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1PricingPreviewImagePriceParams struct {
	// Image size in Gb
	Size param.Field[int64] `json:"size,required"`
}

func (r CloudV1PricingPreviewImagePriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1PricingPreviewLoadBalancerPriceParams struct {
	// Load balancer flavor name
	Flavor param.Field[string] `json:"flavor,required"`
}

func (r CloudV1PricingPreviewLoadBalancerPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1PricingPreviewReservedFixedIPPriceParams struct {
	CreateReservedFixedIP CreateReservedFixedIPUnionParam `json:"create_reserved_fixed_ip"`
}

func (r CloudV1PricingPreviewReservedFixedIPPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateReservedFixedIP)
}

type CloudV1PricingPreviewSnapshotPriceParams struct {
	// Volume ID
	VolumeID param.Field[string] `json:"volume_id,required"`
}

func (r CloudV1PricingPreviewSnapshotPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1PricingPreviewVolumePriceParams struct {
	// One of new-volume, image, snapshot
	Source param.Field[CloudV1PricingPreviewVolumePriceParamsSource] `json:"source,required"`
	// Volume size in GiB
	Size param.Field[int64] `json:"size"`
	// Snapshot ID
	SnapshotID param.Field[string] `json:"snapshot_id"`
	// One of high iops ssd, standard, cold, ultra, ssd low-latency
	TypeName param.Field[VolumeTypeName] `json:"type_name"`
}

func (r CloudV1PricingPreviewVolumePriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// One of new-volume, image, snapshot
type CloudV1PricingPreviewVolumePriceParamsSource string

const (
	CloudV1PricingPreviewVolumePriceParamsSourceImage     CloudV1PricingPreviewVolumePriceParamsSource = "image"
	CloudV1PricingPreviewVolumePriceParamsSourceNewVolume CloudV1PricingPreviewVolumePriceParamsSource = "new-volume"
	CloudV1PricingPreviewVolumePriceParamsSourceSnapshot  CloudV1PricingPreviewVolumePriceParamsSource = "snapshot"
)

func (r CloudV1PricingPreviewVolumePriceParamsSource) IsKnown() bool {
	switch r {
	case CloudV1PricingPreviewVolumePriceParamsSourceImage, CloudV1PricingPreviewVolumePriceParamsSourceNewVolume, CloudV1PricingPreviewVolumePriceParamsSourceSnapshot:
		return true
	}
	return false
}
