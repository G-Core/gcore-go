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

// CloudV1PricingAIClusterService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1PricingAIClusterService] method instead.
type CloudV1PricingAIClusterService struct {
	Options []option.RequestOption
}

// NewCloudV1PricingAIClusterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1PricingAIClusterService(opts ...option.RequestOption) (r *CloudV1PricingAIClusterService) {
	r = &CloudV1PricingAIClusterService{}
	r.Options = opts
	return
}

// Get billing rate of the existing AI cluster based on the current configuration.
func (r *CloudV1PricingAIClusterService) GetAIClusterPrice(ctx context.Context, projectID int64, regionID int64, clusterID string, opts ...option.RequestOption) (res *ItemPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/pricing/%v/%v/ai/clusters/%s", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Preview billing price of the AI cluster
func (r *CloudV1PricingAIClusterService) PreviewAIClusterPrice(ctx context.Context, projectID int64, regionID int64, body CloudV1PricingAIClusterPreviewAIClusterPriceParams, opts ...option.RequestOption) (res *InstancePricingPreviewV2, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/pricing/%v/%v/ai/clusters", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Pricing schema with detailed prices per category.
type DetailedInstancePricing struct {
	// Price of external/public IP over the specified period of time.
	ExternalIP DetailedInstancePricingExternalIPUnion `json:"external_ip,nullable"`
	// Price of flavor over the specified period of time.
	Flavor DetailedInstancePricingFlavorUnion `json:"flavor,nullable"`
	// Price of floating IP over the specified period of time.
	FloatingIP DetailedInstancePricingFloatingIPUnion `json:"floating_ip,nullable"`
	// Pricing schema with detailed prices per volume type.
	Volumes DetailedInstancePricingVolumes `json:"volumes,nullable"`
	JSON    detailedInstancePricingJSON    `json:"-"`
}

// detailedInstancePricingJSON contains the JSON metadata for the struct
// [DetailedInstancePricing]
type detailedInstancePricingJSON struct {
	ExternalIP  apijson.Field
	Flavor      apijson.Field
	FloatingIP  apijson.Field
	Volumes     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DetailedInstancePricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailedInstancePricingJSON) RawJSON() string {
	return r.raw
}

// Price of external/public IP over the specified period of time.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type DetailedInstancePricingExternalIPUnion interface {
	ImplementsDetailedInstancePricingExternalIPUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DetailedInstancePricingExternalIPUnion)(nil)).Elem(),
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

// Price of flavor over the specified period of time.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type DetailedInstancePricingFlavorUnion interface {
	ImplementsDetailedInstancePricingFlavorUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DetailedInstancePricingFlavorUnion)(nil)).Elem(),
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

// Price of floating IP over the specified period of time.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type DetailedInstancePricingFloatingIPUnion interface {
	ImplementsDetailedInstancePricingFloatingIPUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DetailedInstancePricingFloatingIPUnion)(nil)).Elem(),
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

// Pricing schema with detailed prices per volume type.
type DetailedInstancePricingVolumes struct {
	// Price of cold volume over the specified period of time.
	Cold DetailedInstancePricingVolumesColdUnion `json:"cold,nullable"`
	// Price of ssd_hiiops volume over the specified period of time.
	SsdHiiops DetailedInstancePricingVolumesSsdHiiopsUnion `json:"ssd_hiiops,nullable"`
	// Price of ssd_lowlatency volume over the specified period of time.
	SsdLowlatency DetailedInstancePricingVolumesSsdLowlatencyUnion `json:"ssd_lowlatency,nullable"`
	// Price of standard volume over the specified period of time.
	Standard DetailedInstancePricingVolumesStandardUnion `json:"standard,nullable"`
	// Price of ultra volume over the specified period of time.
	Ultra DetailedInstancePricingVolumesUltraUnion `json:"ultra,nullable"`
	JSON  detailedInstancePricingVolumesJSON       `json:"-"`
}

// detailedInstancePricingVolumesJSON contains the JSON metadata for the struct
// [DetailedInstancePricingVolumes]
type detailedInstancePricingVolumesJSON struct {
	Cold          apijson.Field
	SsdHiiops     apijson.Field
	SsdLowlatency apijson.Field
	Standard      apijson.Field
	Ultra         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DetailedInstancePricingVolumes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailedInstancePricingVolumesJSON) RawJSON() string {
	return r.raw
}

// Price of cold volume over the specified period of time.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type DetailedInstancePricingVolumesColdUnion interface {
	ImplementsDetailedInstancePricingVolumesColdUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DetailedInstancePricingVolumesColdUnion)(nil)).Elem(),
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

// Price of ssd_hiiops volume over the specified period of time.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type DetailedInstancePricingVolumesSsdHiiopsUnion interface {
	ImplementsDetailedInstancePricingVolumesSsdHiiopsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DetailedInstancePricingVolumesSsdHiiopsUnion)(nil)).Elem(),
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

// Price of ssd_lowlatency volume over the specified period of time.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type DetailedInstancePricingVolumesSsdLowlatencyUnion interface {
	ImplementsDetailedInstancePricingVolumesSsdLowlatencyUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DetailedInstancePricingVolumesSsdLowlatencyUnion)(nil)).Elem(),
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

// Price of standard volume over the specified period of time.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type DetailedInstancePricingVolumesStandardUnion interface {
	ImplementsDetailedInstancePricingVolumesStandardUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DetailedInstancePricingVolumesStandardUnion)(nil)).Elem(),
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

// Price of ultra volume over the specified period of time.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type DetailedInstancePricingVolumesUltraUnion interface {
	ImplementsDetailedInstancePricingVolumesUltraUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DetailedInstancePricingVolumesUltraUnion)(nil)).Elem(),
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

// Billing response v2 preview schema for server instance.
type InstancePricingPreviewV2 struct {
	// Prices per category charged per hour.
	PerHour DetailedInstancePricing `json:"per_hour,required"`
	// Prices per category charged per month.
	PerMonth DetailedInstancePricing `json:"per_month,required"`
	// Price status for the UI.
	PriceStatus PriceDisplayStatus `json:"price_status,required"`
	// Currency code (3-letter code per ISO 4217).
	CurrencyCode InstancePricingPreviewV2CurrencyCode `json:"currency_code,nullable"`
	// Actual discount as a relative value.
	DiscountPercent InstancePricingPreviewV2DiscountPercentUnion `json:"discount_percent,nullable"`
	// Pricing details per category charged per GB.
	PerGB InstancePricingPreviewV2PerGB `json:"per_gb,nullable"`
	// Total price VAT inclusive per month without discount.
	PriceWithoutDiscountPerMonth InstancePricingPreviewV2PriceWithoutDiscountPerMonthUnion `json:"price_without_discount_per_month,nullable"`
	// Total price VAT inclusive per hour.
	TotalPricePerHour InstancePricingPreviewV2TotalPricePerHourUnion `json:"total_price_per_hour,nullable"`
	// Total price VAT inclusive per month.
	TotalPricePerMonth InstancePricingPreviewV2TotalPricePerMonthUnion `json:"total_price_per_month,nullable"`
	// Number of virtual machine instances being created.
	VmCount int64                        `json:"vm_count"`
	JSON    instancePricingPreviewV2JSON `json:"-"`
}

// instancePricingPreviewV2JSON contains the JSON metadata for the struct
// [InstancePricingPreviewV2]
type instancePricingPreviewV2JSON struct {
	PerHour                      apijson.Field
	PerMonth                     apijson.Field
	PriceStatus                  apijson.Field
	CurrencyCode                 apijson.Field
	DiscountPercent              apijson.Field
	PerGB                        apijson.Field
	PriceWithoutDiscountPerMonth apijson.Field
	TotalPricePerHour            apijson.Field
	TotalPricePerMonth           apijson.Field
	VmCount                      apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *InstancePricingPreviewV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instancePricingPreviewV2JSON) RawJSON() string {
	return r.raw
}

// Currency code (3-letter code per ISO 4217).
type InstancePricingPreviewV2CurrencyCode string

const (
	InstancePricingPreviewV2CurrencyCodeAed InstancePricingPreviewV2CurrencyCode = "AED"
	InstancePricingPreviewV2CurrencyCodeAfn InstancePricingPreviewV2CurrencyCode = "AFN"
	InstancePricingPreviewV2CurrencyCodeAll InstancePricingPreviewV2CurrencyCode = "ALL"
	InstancePricingPreviewV2CurrencyCodeAmd InstancePricingPreviewV2CurrencyCode = "AMD"
	InstancePricingPreviewV2CurrencyCodeAng InstancePricingPreviewV2CurrencyCode = "ANG"
	InstancePricingPreviewV2CurrencyCodeAoa InstancePricingPreviewV2CurrencyCode = "AOA"
	InstancePricingPreviewV2CurrencyCodeArs InstancePricingPreviewV2CurrencyCode = "ARS"
	InstancePricingPreviewV2CurrencyCodeAud InstancePricingPreviewV2CurrencyCode = "AUD"
	InstancePricingPreviewV2CurrencyCodeAwg InstancePricingPreviewV2CurrencyCode = "AWG"
	InstancePricingPreviewV2CurrencyCodeAzn InstancePricingPreviewV2CurrencyCode = "AZN"
	InstancePricingPreviewV2CurrencyCodeBam InstancePricingPreviewV2CurrencyCode = "BAM"
	InstancePricingPreviewV2CurrencyCodeBbd InstancePricingPreviewV2CurrencyCode = "BBD"
	InstancePricingPreviewV2CurrencyCodeBdt InstancePricingPreviewV2CurrencyCode = "BDT"
	InstancePricingPreviewV2CurrencyCodeBgn InstancePricingPreviewV2CurrencyCode = "BGN"
	InstancePricingPreviewV2CurrencyCodeBhd InstancePricingPreviewV2CurrencyCode = "BHD"
	InstancePricingPreviewV2CurrencyCodeBif InstancePricingPreviewV2CurrencyCode = "BIF"
	InstancePricingPreviewV2CurrencyCodeBmd InstancePricingPreviewV2CurrencyCode = "BMD"
	InstancePricingPreviewV2CurrencyCodeBnd InstancePricingPreviewV2CurrencyCode = "BND"
	InstancePricingPreviewV2CurrencyCodeBob InstancePricingPreviewV2CurrencyCode = "BOB"
	InstancePricingPreviewV2CurrencyCodeBrl InstancePricingPreviewV2CurrencyCode = "BRL"
	InstancePricingPreviewV2CurrencyCodeBsd InstancePricingPreviewV2CurrencyCode = "BSD"
	InstancePricingPreviewV2CurrencyCodeBtn InstancePricingPreviewV2CurrencyCode = "BTN"
	InstancePricingPreviewV2CurrencyCodeBwp InstancePricingPreviewV2CurrencyCode = "BWP"
	InstancePricingPreviewV2CurrencyCodeByn InstancePricingPreviewV2CurrencyCode = "BYN"
	InstancePricingPreviewV2CurrencyCodeBzd InstancePricingPreviewV2CurrencyCode = "BZD"
	InstancePricingPreviewV2CurrencyCodeCad InstancePricingPreviewV2CurrencyCode = "CAD"
	InstancePricingPreviewV2CurrencyCodeCdf InstancePricingPreviewV2CurrencyCode = "CDF"
	InstancePricingPreviewV2CurrencyCodeChf InstancePricingPreviewV2CurrencyCode = "CHF"
	InstancePricingPreviewV2CurrencyCodeClp InstancePricingPreviewV2CurrencyCode = "CLP"
	InstancePricingPreviewV2CurrencyCodeCny InstancePricingPreviewV2CurrencyCode = "CNY"
	InstancePricingPreviewV2CurrencyCodeCop InstancePricingPreviewV2CurrencyCode = "COP"
	InstancePricingPreviewV2CurrencyCodeCrc InstancePricingPreviewV2CurrencyCode = "CRC"
	InstancePricingPreviewV2CurrencyCodeCuc InstancePricingPreviewV2CurrencyCode = "CUC"
	InstancePricingPreviewV2CurrencyCodeCup InstancePricingPreviewV2CurrencyCode = "CUP"
	InstancePricingPreviewV2CurrencyCodeCve InstancePricingPreviewV2CurrencyCode = "CVE"
	InstancePricingPreviewV2CurrencyCodeCzk InstancePricingPreviewV2CurrencyCode = "CZK"
	InstancePricingPreviewV2CurrencyCodeDjf InstancePricingPreviewV2CurrencyCode = "DJF"
	InstancePricingPreviewV2CurrencyCodeDkk InstancePricingPreviewV2CurrencyCode = "DKK"
	InstancePricingPreviewV2CurrencyCodeDop InstancePricingPreviewV2CurrencyCode = "DOP"
	InstancePricingPreviewV2CurrencyCodeDzd InstancePricingPreviewV2CurrencyCode = "DZD"
	InstancePricingPreviewV2CurrencyCodeEgp InstancePricingPreviewV2CurrencyCode = "EGP"
	InstancePricingPreviewV2CurrencyCodeErn InstancePricingPreviewV2CurrencyCode = "ERN"
	InstancePricingPreviewV2CurrencyCodeEtb InstancePricingPreviewV2CurrencyCode = "ETB"
	InstancePricingPreviewV2CurrencyCodeEur InstancePricingPreviewV2CurrencyCode = "EUR"
	InstancePricingPreviewV2CurrencyCodeFjd InstancePricingPreviewV2CurrencyCode = "FJD"
	InstancePricingPreviewV2CurrencyCodeFkp InstancePricingPreviewV2CurrencyCode = "FKP"
	InstancePricingPreviewV2CurrencyCodeGbp InstancePricingPreviewV2CurrencyCode = "GBP"
	InstancePricingPreviewV2CurrencyCodeGel InstancePricingPreviewV2CurrencyCode = "GEL"
	InstancePricingPreviewV2CurrencyCodeGhs InstancePricingPreviewV2CurrencyCode = "GHS"
	InstancePricingPreviewV2CurrencyCodeGip InstancePricingPreviewV2CurrencyCode = "GIP"
	InstancePricingPreviewV2CurrencyCodeGmd InstancePricingPreviewV2CurrencyCode = "GMD"
	InstancePricingPreviewV2CurrencyCodeGnf InstancePricingPreviewV2CurrencyCode = "GNF"
	InstancePricingPreviewV2CurrencyCodeGtq InstancePricingPreviewV2CurrencyCode = "GTQ"
	InstancePricingPreviewV2CurrencyCodeGyd InstancePricingPreviewV2CurrencyCode = "GYD"
	InstancePricingPreviewV2CurrencyCodeHkd InstancePricingPreviewV2CurrencyCode = "HKD"
	InstancePricingPreviewV2CurrencyCodeHnl InstancePricingPreviewV2CurrencyCode = "HNL"
	InstancePricingPreviewV2CurrencyCodeHrk InstancePricingPreviewV2CurrencyCode = "HRK"
	InstancePricingPreviewV2CurrencyCodeHtg InstancePricingPreviewV2CurrencyCode = "HTG"
	InstancePricingPreviewV2CurrencyCodeHuf InstancePricingPreviewV2CurrencyCode = "HUF"
	InstancePricingPreviewV2CurrencyCodeIdr InstancePricingPreviewV2CurrencyCode = "IDR"
	InstancePricingPreviewV2CurrencyCodeIls InstancePricingPreviewV2CurrencyCode = "ILS"
	InstancePricingPreviewV2CurrencyCodeInr InstancePricingPreviewV2CurrencyCode = "INR"
	InstancePricingPreviewV2CurrencyCodeIqd InstancePricingPreviewV2CurrencyCode = "IQD"
	InstancePricingPreviewV2CurrencyCodeIrr InstancePricingPreviewV2CurrencyCode = "IRR"
	InstancePricingPreviewV2CurrencyCodeIsk InstancePricingPreviewV2CurrencyCode = "ISK"
	InstancePricingPreviewV2CurrencyCodeJmd InstancePricingPreviewV2CurrencyCode = "JMD"
	InstancePricingPreviewV2CurrencyCodeJod InstancePricingPreviewV2CurrencyCode = "JOD"
	InstancePricingPreviewV2CurrencyCodeJpy InstancePricingPreviewV2CurrencyCode = "JPY"
	InstancePricingPreviewV2CurrencyCodeKes InstancePricingPreviewV2CurrencyCode = "KES"
	InstancePricingPreviewV2CurrencyCodeKgs InstancePricingPreviewV2CurrencyCode = "KGS"
	InstancePricingPreviewV2CurrencyCodeKhr InstancePricingPreviewV2CurrencyCode = "KHR"
	InstancePricingPreviewV2CurrencyCodeKmf InstancePricingPreviewV2CurrencyCode = "KMF"
	InstancePricingPreviewV2CurrencyCodeKpw InstancePricingPreviewV2CurrencyCode = "KPW"
	InstancePricingPreviewV2CurrencyCodeKrw InstancePricingPreviewV2CurrencyCode = "KRW"
	InstancePricingPreviewV2CurrencyCodeKwd InstancePricingPreviewV2CurrencyCode = "KWD"
	InstancePricingPreviewV2CurrencyCodeKyd InstancePricingPreviewV2CurrencyCode = "KYD"
	InstancePricingPreviewV2CurrencyCodeKzt InstancePricingPreviewV2CurrencyCode = "KZT"
	InstancePricingPreviewV2CurrencyCodeLak InstancePricingPreviewV2CurrencyCode = "LAK"
	InstancePricingPreviewV2CurrencyCodeLbp InstancePricingPreviewV2CurrencyCode = "LBP"
	InstancePricingPreviewV2CurrencyCodeLkr InstancePricingPreviewV2CurrencyCode = "LKR"
	InstancePricingPreviewV2CurrencyCodeLrd InstancePricingPreviewV2CurrencyCode = "LRD"
	InstancePricingPreviewV2CurrencyCodeLsl InstancePricingPreviewV2CurrencyCode = "LSL"
	InstancePricingPreviewV2CurrencyCodeLyd InstancePricingPreviewV2CurrencyCode = "LYD"
	InstancePricingPreviewV2CurrencyCodeMad InstancePricingPreviewV2CurrencyCode = "MAD"
	InstancePricingPreviewV2CurrencyCodeMdl InstancePricingPreviewV2CurrencyCode = "MDL"
	InstancePricingPreviewV2CurrencyCodeMga InstancePricingPreviewV2CurrencyCode = "MGA"
	InstancePricingPreviewV2CurrencyCodeMkd InstancePricingPreviewV2CurrencyCode = "MKD"
	InstancePricingPreviewV2CurrencyCodeMmk InstancePricingPreviewV2CurrencyCode = "MMK"
	InstancePricingPreviewV2CurrencyCodeMnt InstancePricingPreviewV2CurrencyCode = "MNT"
	InstancePricingPreviewV2CurrencyCodeMop InstancePricingPreviewV2CurrencyCode = "MOP"
	InstancePricingPreviewV2CurrencyCodeMro InstancePricingPreviewV2CurrencyCode = "MRO"
	InstancePricingPreviewV2CurrencyCodeMur InstancePricingPreviewV2CurrencyCode = "MUR"
	InstancePricingPreviewV2CurrencyCodeMvr InstancePricingPreviewV2CurrencyCode = "MVR"
	InstancePricingPreviewV2CurrencyCodeMwk InstancePricingPreviewV2CurrencyCode = "MWK"
	InstancePricingPreviewV2CurrencyCodeMxn InstancePricingPreviewV2CurrencyCode = "MXN"
	InstancePricingPreviewV2CurrencyCodeMyr InstancePricingPreviewV2CurrencyCode = "MYR"
	InstancePricingPreviewV2CurrencyCodeMzn InstancePricingPreviewV2CurrencyCode = "MZN"
	InstancePricingPreviewV2CurrencyCodeNad InstancePricingPreviewV2CurrencyCode = "NAD"
	InstancePricingPreviewV2CurrencyCodeNgn InstancePricingPreviewV2CurrencyCode = "NGN"
	InstancePricingPreviewV2CurrencyCodeNio InstancePricingPreviewV2CurrencyCode = "NIO"
	InstancePricingPreviewV2CurrencyCodeNok InstancePricingPreviewV2CurrencyCode = "NOK"
	InstancePricingPreviewV2CurrencyCodeNpr InstancePricingPreviewV2CurrencyCode = "NPR"
	InstancePricingPreviewV2CurrencyCodeNzd InstancePricingPreviewV2CurrencyCode = "NZD"
	InstancePricingPreviewV2CurrencyCodeOmr InstancePricingPreviewV2CurrencyCode = "OMR"
	InstancePricingPreviewV2CurrencyCodePab InstancePricingPreviewV2CurrencyCode = "PAB"
	InstancePricingPreviewV2CurrencyCodePen InstancePricingPreviewV2CurrencyCode = "PEN"
	InstancePricingPreviewV2CurrencyCodePgk InstancePricingPreviewV2CurrencyCode = "PGK"
	InstancePricingPreviewV2CurrencyCodePhp InstancePricingPreviewV2CurrencyCode = "PHP"
	InstancePricingPreviewV2CurrencyCodePkr InstancePricingPreviewV2CurrencyCode = "PKR"
	InstancePricingPreviewV2CurrencyCodePln InstancePricingPreviewV2CurrencyCode = "PLN"
	InstancePricingPreviewV2CurrencyCodePyg InstancePricingPreviewV2CurrencyCode = "PYG"
	InstancePricingPreviewV2CurrencyCodeQar InstancePricingPreviewV2CurrencyCode = "QAR"
	InstancePricingPreviewV2CurrencyCodeRon InstancePricingPreviewV2CurrencyCode = "RON"
	InstancePricingPreviewV2CurrencyCodeRsd InstancePricingPreviewV2CurrencyCode = "RSD"
	InstancePricingPreviewV2CurrencyCodeRub InstancePricingPreviewV2CurrencyCode = "RUB"
	InstancePricingPreviewV2CurrencyCodeRwf InstancePricingPreviewV2CurrencyCode = "RWF"
	InstancePricingPreviewV2CurrencyCodeSar InstancePricingPreviewV2CurrencyCode = "SAR"
	InstancePricingPreviewV2CurrencyCodeSbd InstancePricingPreviewV2CurrencyCode = "SBD"
	InstancePricingPreviewV2CurrencyCodeScr InstancePricingPreviewV2CurrencyCode = "SCR"
	InstancePricingPreviewV2CurrencyCodeSdg InstancePricingPreviewV2CurrencyCode = "SDG"
	InstancePricingPreviewV2CurrencyCodeSek InstancePricingPreviewV2CurrencyCode = "SEK"
	InstancePricingPreviewV2CurrencyCodeSgd InstancePricingPreviewV2CurrencyCode = "SGD"
	InstancePricingPreviewV2CurrencyCodeShp InstancePricingPreviewV2CurrencyCode = "SHP"
	InstancePricingPreviewV2CurrencyCodeSll InstancePricingPreviewV2CurrencyCode = "SLL"
	InstancePricingPreviewV2CurrencyCodeSos InstancePricingPreviewV2CurrencyCode = "SOS"
	InstancePricingPreviewV2CurrencyCodeSrd InstancePricingPreviewV2CurrencyCode = "SRD"
	InstancePricingPreviewV2CurrencyCodeSsp InstancePricingPreviewV2CurrencyCode = "SSP"
	InstancePricingPreviewV2CurrencyCodeStd InstancePricingPreviewV2CurrencyCode = "STD"
	InstancePricingPreviewV2CurrencyCodeSvc InstancePricingPreviewV2CurrencyCode = "SVC"
	InstancePricingPreviewV2CurrencyCodeSyp InstancePricingPreviewV2CurrencyCode = "SYP"
	InstancePricingPreviewV2CurrencyCodeSzl InstancePricingPreviewV2CurrencyCode = "SZL"
	InstancePricingPreviewV2CurrencyCodeThb InstancePricingPreviewV2CurrencyCode = "THB"
	InstancePricingPreviewV2CurrencyCodeTjs InstancePricingPreviewV2CurrencyCode = "TJS"
	InstancePricingPreviewV2CurrencyCodeTmt InstancePricingPreviewV2CurrencyCode = "TMT"
	InstancePricingPreviewV2CurrencyCodeTnd InstancePricingPreviewV2CurrencyCode = "TND"
	InstancePricingPreviewV2CurrencyCodeTop InstancePricingPreviewV2CurrencyCode = "TOP"
	InstancePricingPreviewV2CurrencyCodeTry InstancePricingPreviewV2CurrencyCode = "TRY"
	InstancePricingPreviewV2CurrencyCodeTtd InstancePricingPreviewV2CurrencyCode = "TTD"
	InstancePricingPreviewV2CurrencyCodeTwd InstancePricingPreviewV2CurrencyCode = "TWD"
	InstancePricingPreviewV2CurrencyCodeTzs InstancePricingPreviewV2CurrencyCode = "TZS"
	InstancePricingPreviewV2CurrencyCodeUah InstancePricingPreviewV2CurrencyCode = "UAH"
	InstancePricingPreviewV2CurrencyCodeUgx InstancePricingPreviewV2CurrencyCode = "UGX"
	InstancePricingPreviewV2CurrencyCodeUsd InstancePricingPreviewV2CurrencyCode = "USD"
	InstancePricingPreviewV2CurrencyCodeUyu InstancePricingPreviewV2CurrencyCode = "UYU"
	InstancePricingPreviewV2CurrencyCodeUzs InstancePricingPreviewV2CurrencyCode = "UZS"
	InstancePricingPreviewV2CurrencyCodeVef InstancePricingPreviewV2CurrencyCode = "VEF"
	InstancePricingPreviewV2CurrencyCodeVnd InstancePricingPreviewV2CurrencyCode = "VND"
	InstancePricingPreviewV2CurrencyCodeVuv InstancePricingPreviewV2CurrencyCode = "VUV"
	InstancePricingPreviewV2CurrencyCodeWst InstancePricingPreviewV2CurrencyCode = "WST"
	InstancePricingPreviewV2CurrencyCodeXaf InstancePricingPreviewV2CurrencyCode = "XAF"
	InstancePricingPreviewV2CurrencyCodeXag InstancePricingPreviewV2CurrencyCode = "XAG"
	InstancePricingPreviewV2CurrencyCodeXau InstancePricingPreviewV2CurrencyCode = "XAU"
	InstancePricingPreviewV2CurrencyCodeXba InstancePricingPreviewV2CurrencyCode = "XBA"
	InstancePricingPreviewV2CurrencyCodeXbb InstancePricingPreviewV2CurrencyCode = "XBB"
	InstancePricingPreviewV2CurrencyCodeXbc InstancePricingPreviewV2CurrencyCode = "XBC"
	InstancePricingPreviewV2CurrencyCodeXbd InstancePricingPreviewV2CurrencyCode = "XBD"
	InstancePricingPreviewV2CurrencyCodeXcd InstancePricingPreviewV2CurrencyCode = "XCD"
	InstancePricingPreviewV2CurrencyCodeXdr InstancePricingPreviewV2CurrencyCode = "XDR"
	InstancePricingPreviewV2CurrencyCodeXof InstancePricingPreviewV2CurrencyCode = "XOF"
	InstancePricingPreviewV2CurrencyCodeXpd InstancePricingPreviewV2CurrencyCode = "XPD"
	InstancePricingPreviewV2CurrencyCodeXpf InstancePricingPreviewV2CurrencyCode = "XPF"
	InstancePricingPreviewV2CurrencyCodeXpt InstancePricingPreviewV2CurrencyCode = "XPT"
	InstancePricingPreviewV2CurrencyCodeXsu InstancePricingPreviewV2CurrencyCode = "XSU"
	InstancePricingPreviewV2CurrencyCodeXts InstancePricingPreviewV2CurrencyCode = "XTS"
	InstancePricingPreviewV2CurrencyCodeXua InstancePricingPreviewV2CurrencyCode = "XUA"
	InstancePricingPreviewV2CurrencyCodeXxx InstancePricingPreviewV2CurrencyCode = "XXX"
	InstancePricingPreviewV2CurrencyCodeYer InstancePricingPreviewV2CurrencyCode = "YER"
	InstancePricingPreviewV2CurrencyCodeZar InstancePricingPreviewV2CurrencyCode = "ZAR"
	InstancePricingPreviewV2CurrencyCodeZmw InstancePricingPreviewV2CurrencyCode = "ZMW"
	InstancePricingPreviewV2CurrencyCodeZwl InstancePricingPreviewV2CurrencyCode = "ZWL"
)

func (r InstancePricingPreviewV2CurrencyCode) IsKnown() bool {
	switch r {
	case InstancePricingPreviewV2CurrencyCodeAed, InstancePricingPreviewV2CurrencyCodeAfn, InstancePricingPreviewV2CurrencyCodeAll, InstancePricingPreviewV2CurrencyCodeAmd, InstancePricingPreviewV2CurrencyCodeAng, InstancePricingPreviewV2CurrencyCodeAoa, InstancePricingPreviewV2CurrencyCodeArs, InstancePricingPreviewV2CurrencyCodeAud, InstancePricingPreviewV2CurrencyCodeAwg, InstancePricingPreviewV2CurrencyCodeAzn, InstancePricingPreviewV2CurrencyCodeBam, InstancePricingPreviewV2CurrencyCodeBbd, InstancePricingPreviewV2CurrencyCodeBdt, InstancePricingPreviewV2CurrencyCodeBgn, InstancePricingPreviewV2CurrencyCodeBhd, InstancePricingPreviewV2CurrencyCodeBif, InstancePricingPreviewV2CurrencyCodeBmd, InstancePricingPreviewV2CurrencyCodeBnd, InstancePricingPreviewV2CurrencyCodeBob, InstancePricingPreviewV2CurrencyCodeBrl, InstancePricingPreviewV2CurrencyCodeBsd, InstancePricingPreviewV2CurrencyCodeBtn, InstancePricingPreviewV2CurrencyCodeBwp, InstancePricingPreviewV2CurrencyCodeByn, InstancePricingPreviewV2CurrencyCodeBzd, InstancePricingPreviewV2CurrencyCodeCad, InstancePricingPreviewV2CurrencyCodeCdf, InstancePricingPreviewV2CurrencyCodeChf, InstancePricingPreviewV2CurrencyCodeClp, InstancePricingPreviewV2CurrencyCodeCny, InstancePricingPreviewV2CurrencyCodeCop, InstancePricingPreviewV2CurrencyCodeCrc, InstancePricingPreviewV2CurrencyCodeCuc, InstancePricingPreviewV2CurrencyCodeCup, InstancePricingPreviewV2CurrencyCodeCve, InstancePricingPreviewV2CurrencyCodeCzk, InstancePricingPreviewV2CurrencyCodeDjf, InstancePricingPreviewV2CurrencyCodeDkk, InstancePricingPreviewV2CurrencyCodeDop, InstancePricingPreviewV2CurrencyCodeDzd, InstancePricingPreviewV2CurrencyCodeEgp, InstancePricingPreviewV2CurrencyCodeErn, InstancePricingPreviewV2CurrencyCodeEtb, InstancePricingPreviewV2CurrencyCodeEur, InstancePricingPreviewV2CurrencyCodeFjd, InstancePricingPreviewV2CurrencyCodeFkp, InstancePricingPreviewV2CurrencyCodeGbp, InstancePricingPreviewV2CurrencyCodeGel, InstancePricingPreviewV2CurrencyCodeGhs, InstancePricingPreviewV2CurrencyCodeGip, InstancePricingPreviewV2CurrencyCodeGmd, InstancePricingPreviewV2CurrencyCodeGnf, InstancePricingPreviewV2CurrencyCodeGtq, InstancePricingPreviewV2CurrencyCodeGyd, InstancePricingPreviewV2CurrencyCodeHkd, InstancePricingPreviewV2CurrencyCodeHnl, InstancePricingPreviewV2CurrencyCodeHrk, InstancePricingPreviewV2CurrencyCodeHtg, InstancePricingPreviewV2CurrencyCodeHuf, InstancePricingPreviewV2CurrencyCodeIdr, InstancePricingPreviewV2CurrencyCodeIls, InstancePricingPreviewV2CurrencyCodeInr, InstancePricingPreviewV2CurrencyCodeIqd, InstancePricingPreviewV2CurrencyCodeIrr, InstancePricingPreviewV2CurrencyCodeIsk, InstancePricingPreviewV2CurrencyCodeJmd, InstancePricingPreviewV2CurrencyCodeJod, InstancePricingPreviewV2CurrencyCodeJpy, InstancePricingPreviewV2CurrencyCodeKes, InstancePricingPreviewV2CurrencyCodeKgs, InstancePricingPreviewV2CurrencyCodeKhr, InstancePricingPreviewV2CurrencyCodeKmf, InstancePricingPreviewV2CurrencyCodeKpw, InstancePricingPreviewV2CurrencyCodeKrw, InstancePricingPreviewV2CurrencyCodeKwd, InstancePricingPreviewV2CurrencyCodeKyd, InstancePricingPreviewV2CurrencyCodeKzt, InstancePricingPreviewV2CurrencyCodeLak, InstancePricingPreviewV2CurrencyCodeLbp, InstancePricingPreviewV2CurrencyCodeLkr, InstancePricingPreviewV2CurrencyCodeLrd, InstancePricingPreviewV2CurrencyCodeLsl, InstancePricingPreviewV2CurrencyCodeLyd, InstancePricingPreviewV2CurrencyCodeMad, InstancePricingPreviewV2CurrencyCodeMdl, InstancePricingPreviewV2CurrencyCodeMga, InstancePricingPreviewV2CurrencyCodeMkd, InstancePricingPreviewV2CurrencyCodeMmk, InstancePricingPreviewV2CurrencyCodeMnt, InstancePricingPreviewV2CurrencyCodeMop, InstancePricingPreviewV2CurrencyCodeMro, InstancePricingPreviewV2CurrencyCodeMur, InstancePricingPreviewV2CurrencyCodeMvr, InstancePricingPreviewV2CurrencyCodeMwk, InstancePricingPreviewV2CurrencyCodeMxn, InstancePricingPreviewV2CurrencyCodeMyr, InstancePricingPreviewV2CurrencyCodeMzn, InstancePricingPreviewV2CurrencyCodeNad, InstancePricingPreviewV2CurrencyCodeNgn, InstancePricingPreviewV2CurrencyCodeNio, InstancePricingPreviewV2CurrencyCodeNok, InstancePricingPreviewV2CurrencyCodeNpr, InstancePricingPreviewV2CurrencyCodeNzd, InstancePricingPreviewV2CurrencyCodeOmr, InstancePricingPreviewV2CurrencyCodePab, InstancePricingPreviewV2CurrencyCodePen, InstancePricingPreviewV2CurrencyCodePgk, InstancePricingPreviewV2CurrencyCodePhp, InstancePricingPreviewV2CurrencyCodePkr, InstancePricingPreviewV2CurrencyCodePln, InstancePricingPreviewV2CurrencyCodePyg, InstancePricingPreviewV2CurrencyCodeQar, InstancePricingPreviewV2CurrencyCodeRon, InstancePricingPreviewV2CurrencyCodeRsd, InstancePricingPreviewV2CurrencyCodeRub, InstancePricingPreviewV2CurrencyCodeRwf, InstancePricingPreviewV2CurrencyCodeSar, InstancePricingPreviewV2CurrencyCodeSbd, InstancePricingPreviewV2CurrencyCodeScr, InstancePricingPreviewV2CurrencyCodeSdg, InstancePricingPreviewV2CurrencyCodeSek, InstancePricingPreviewV2CurrencyCodeSgd, InstancePricingPreviewV2CurrencyCodeShp, InstancePricingPreviewV2CurrencyCodeSll, InstancePricingPreviewV2CurrencyCodeSos, InstancePricingPreviewV2CurrencyCodeSrd, InstancePricingPreviewV2CurrencyCodeSsp, InstancePricingPreviewV2CurrencyCodeStd, InstancePricingPreviewV2CurrencyCodeSvc, InstancePricingPreviewV2CurrencyCodeSyp, InstancePricingPreviewV2CurrencyCodeSzl, InstancePricingPreviewV2CurrencyCodeThb, InstancePricingPreviewV2CurrencyCodeTjs, InstancePricingPreviewV2CurrencyCodeTmt, InstancePricingPreviewV2CurrencyCodeTnd, InstancePricingPreviewV2CurrencyCodeTop, InstancePricingPreviewV2CurrencyCodeTry, InstancePricingPreviewV2CurrencyCodeTtd, InstancePricingPreviewV2CurrencyCodeTwd, InstancePricingPreviewV2CurrencyCodeTzs, InstancePricingPreviewV2CurrencyCodeUah, InstancePricingPreviewV2CurrencyCodeUgx, InstancePricingPreviewV2CurrencyCodeUsd, InstancePricingPreviewV2CurrencyCodeUyu, InstancePricingPreviewV2CurrencyCodeUzs, InstancePricingPreviewV2CurrencyCodeVef, InstancePricingPreviewV2CurrencyCodeVnd, InstancePricingPreviewV2CurrencyCodeVuv, InstancePricingPreviewV2CurrencyCodeWst, InstancePricingPreviewV2CurrencyCodeXaf, InstancePricingPreviewV2CurrencyCodeXag, InstancePricingPreviewV2CurrencyCodeXau, InstancePricingPreviewV2CurrencyCodeXba, InstancePricingPreviewV2CurrencyCodeXbb, InstancePricingPreviewV2CurrencyCodeXbc, InstancePricingPreviewV2CurrencyCodeXbd, InstancePricingPreviewV2CurrencyCodeXcd, InstancePricingPreviewV2CurrencyCodeXdr, InstancePricingPreviewV2CurrencyCodeXof, InstancePricingPreviewV2CurrencyCodeXpd, InstancePricingPreviewV2CurrencyCodeXpf, InstancePricingPreviewV2CurrencyCodeXpt, InstancePricingPreviewV2CurrencyCodeXsu, InstancePricingPreviewV2CurrencyCodeXts, InstancePricingPreviewV2CurrencyCodeXua, InstancePricingPreviewV2CurrencyCodeXxx, InstancePricingPreviewV2CurrencyCodeYer, InstancePricingPreviewV2CurrencyCodeZar, InstancePricingPreviewV2CurrencyCodeZmw, InstancePricingPreviewV2CurrencyCodeZwl:
		return true
	}
	return false
}

// Actual discount as a relative value.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type InstancePricingPreviewV2DiscountPercentUnion interface {
	ImplementsInstancePricingPreviewV2DiscountPercentUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InstancePricingPreviewV2DiscountPercentUnion)(nil)).Elem(),
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

// Pricing details per category charged per GB.
type InstancePricingPreviewV2PerGB struct {
	// The price of egress traffic per GB.
	EgressTrafficBaremetal InstancePricingPreviewV2PerGBEgressTrafficBaremetalUnion `json:"egress_traffic_baremetal,nullable"`
	JSON                   instancePricingPreviewV2PerGBJSON                        `json:"-"`
}

// instancePricingPreviewV2PerGBJSON contains the JSON metadata for the struct
// [InstancePricingPreviewV2PerGB]
type instancePricingPreviewV2PerGBJSON struct {
	EgressTrafficBaremetal apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *InstancePricingPreviewV2PerGB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instancePricingPreviewV2PerGBJSON) RawJSON() string {
	return r.raw
}

// The price of egress traffic per GB.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type InstancePricingPreviewV2PerGBEgressTrafficBaremetalUnion interface {
	ImplementsInstancePricingPreviewV2PerGBEgressTrafficBaremetalUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InstancePricingPreviewV2PerGBEgressTrafficBaremetalUnion)(nil)).Elem(),
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

// Total price VAT inclusive per month without discount.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type InstancePricingPreviewV2PriceWithoutDiscountPerMonthUnion interface {
	ImplementsInstancePricingPreviewV2PriceWithoutDiscountPerMonthUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InstancePricingPreviewV2PriceWithoutDiscountPerMonthUnion)(nil)).Elem(),
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

// Total price VAT inclusive per hour.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type InstancePricingPreviewV2TotalPricePerHourUnion interface {
	ImplementsInstancePricingPreviewV2TotalPricePerHourUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InstancePricingPreviewV2TotalPricePerHourUnion)(nil)).Elem(),
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

// Total price VAT inclusive per month.
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type InstancePricingPreviewV2TotalPricePerMonthUnion interface {
	ImplementsInstancePricingPreviewV2TotalPricePerMonthUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InstancePricingPreviewV2TotalPricePerMonthUnion)(nil)).Elem(),
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

// Schema lists only the fields required to calculate price
type InstanceVolumePricingRequestParam struct {
	// The source of the volume. It must be one of: 'new-volume', 'image', 'snapshot',
	// 'apptemplate'.
	Source param.Field[InstanceVolumePricingRequestSource] `json:"source,required"`
	// Volume size in GiB.
	Size param.Field[int64] `json:"size"`
	// The snapshot ID. It is mandatory if the volume is created from a snapshot.
	SnapshotID param.Field[string] `json:"snapshot_id"`
	// The type of the volume. It must be one of: 'standard', 'ssd_hiiops', 'cold',
	// 'ultra'.
	TypeName param.Field[VolumeTypeName] `json:"type_name"`
}

func (r InstanceVolumePricingRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The source of the volume. It must be one of: 'new-volume', 'image', 'snapshot',
// 'apptemplate'.
type InstanceVolumePricingRequestSource string

const (
	InstanceVolumePricingRequestSourceApptemplate    InstanceVolumePricingRequestSource = "apptemplate"
	InstanceVolumePricingRequestSourceExistingVolume InstanceVolumePricingRequestSource = "existing-volume"
	InstanceVolumePricingRequestSourceImage          InstanceVolumePricingRequestSource = "image"
	InstanceVolumePricingRequestSourceNewVolume      InstanceVolumePricingRequestSource = "new-volume"
	InstanceVolumePricingRequestSourceSnapshot       InstanceVolumePricingRequestSource = "snapshot"
)

func (r InstanceVolumePricingRequestSource) IsKnown() bool {
	switch r {
	case InstanceVolumePricingRequestSourceApptemplate, InstanceVolumePricingRequestSourceExistingVolume, InstanceVolumePricingRequestSourceImage, InstanceVolumePricingRequestSourceNewVolume, InstanceVolumePricingRequestSourceSnapshot:
		return true
	}
	return false
}

// Instance will be attached to default external network
type NewVmInterfaceSerializersPydanticParam struct {
	// Floating IP config for this subnet attachment.
	FloatingIP param.Field[NewInterfaceFloatingIPPydanticParam] `json:"floating_ip"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Fixed IP address.
	IPAddress param.Field[string] `json:"ip_address"`
	// Which subnets should be selected: IPv4, IPv6, or use dual stack.
	IPFamily param.Field[InterfaceIPFamily] `json:"ip_family"`
	// Network ID the subnet belongs to. Port will be plugged in this network.
	NetworkID param.Field[string] `json:"network_id"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64] `json:"port_group"`
	// Network ID the subnet belongs to. Port will be plugged in this network.
	PortID param.Field[string] `json:"port_id"`
	// Port is assigned an IP address from this subnet.
	SubnetID param.Field[string]                                `json:"subnet_id"`
	Type     param.Field[NewVmInterfaceSerializersPydanticType] `json:"type"`
}

func (r NewVmInterfaceSerializersPydanticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewVmInterfaceSerializersPydanticParam) implementsNewVmInterfaceSerializersPydanticUnionParam() {
}

// Instance will be attached to default external network
//
// Satisfied by
// [NewVmInterfaceSerializersPydanticNewInterfaceExternalSerializerPydanticParam],
// [NewVmInterfaceSerializersPydanticNewInterfaceSpecificSubnetFipSerializerPydanticParam],
// [NewVmInterfaceSerializersPydanticNewInterfaceAnySubnetFipSerializerPydanticParam],
// [NewVmInterfaceSerializersPydanticNewInterfaceReservedFixedIPFipSerializerPydanticParam],
// [NewVmInterfaceSerializersPydanticParam].
type NewVmInterfaceSerializersPydanticUnionParam interface {
	implementsNewVmInterfaceSerializersPydanticUnionParam()
}

// Instance will be attached to default external network
type NewVmInterfaceSerializersPydanticNewInterfaceExternalSerializerPydanticParam struct {
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Which subnets should be selected: IPv4, IPv6, or use dual stack.
	IPFamily param.Field[InterfaceIPFamily] `json:"ip_family"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64]                                                                       `json:"port_group"`
	Type      param.Field[NewVmInterfaceSerializersPydanticNewInterfaceExternalSerializerPydanticType] `json:"type"`
}

func (r NewVmInterfaceSerializersPydanticNewInterfaceExternalSerializerPydanticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewVmInterfaceSerializersPydanticNewInterfaceExternalSerializerPydanticParam) implementsNewVmInterfaceSerializersPydanticUnionParam() {
}

type NewVmInterfaceSerializersPydanticNewInterfaceExternalSerializerPydanticType string

const (
	NewVmInterfaceSerializersPydanticNewInterfaceExternalSerializerPydanticTypeExternal NewVmInterfaceSerializersPydanticNewInterfaceExternalSerializerPydanticType = "external"
)

func (r NewVmInterfaceSerializersPydanticNewInterfaceExternalSerializerPydanticType) IsKnown() bool {
	switch r {
	case NewVmInterfaceSerializersPydanticNewInterfaceExternalSerializerPydanticTypeExternal:
		return true
	}
	return false
}

// Instance will be attached to specified subnet, and floating IP can be created
// and attached to this IP
type NewVmInterfaceSerializersPydanticNewInterfaceSpecificSubnetFipSerializerPydanticParam struct {
	// Network ID the subnet belongs to. Port will be plugged in this network.
	NetworkID param.Field[string] `json:"network_id,required"`
	// Port is assigned an IP address from this subnet.
	SubnetID param.Field[string] `json:"subnet_id,required"`
	// Floating IP config for this subnet attachment.
	FloatingIP param.Field[NewInterfaceFloatingIPPydanticParam] `json:"floating_ip"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64]                                                                                `json:"port_group"`
	Type      param.Field[NewVmInterfaceSerializersPydanticNewInterfaceSpecificSubnetFipSerializerPydanticType] `json:"type"`
}

func (r NewVmInterfaceSerializersPydanticNewInterfaceSpecificSubnetFipSerializerPydanticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewVmInterfaceSerializersPydanticNewInterfaceSpecificSubnetFipSerializerPydanticParam) implementsNewVmInterfaceSerializersPydanticUnionParam() {
}

type NewVmInterfaceSerializersPydanticNewInterfaceSpecificSubnetFipSerializerPydanticType string

const (
	NewVmInterfaceSerializersPydanticNewInterfaceSpecificSubnetFipSerializerPydanticTypeSubnet NewVmInterfaceSerializersPydanticNewInterfaceSpecificSubnetFipSerializerPydanticType = "subnet"
)

func (r NewVmInterfaceSerializersPydanticNewInterfaceSpecificSubnetFipSerializerPydanticType) IsKnown() bool {
	switch r {
	case NewVmInterfaceSerializersPydanticNewInterfaceSpecificSubnetFipSerializerPydanticTypeSubnet:
		return true
	}
	return false
}

// Instance will be attached to a subnet with the largest count of free IPs.
// Floating IP will be created and attached to that IP.
type NewVmInterfaceSerializersPydanticNewInterfaceAnySubnetFipSerializerPydanticParam struct {
	// Network ID the subnet belongs to. Port will be plugged in this network.
	NetworkID param.Field[string] `json:"network_id,required"`
	// Floating IP config for this subnet attachment.
	FloatingIP param.Field[NewInterfaceFloatingIPPydanticParam] `json:"floating_ip"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Fixed IP address.
	IPAddress param.Field[string] `json:"ip_address"`
	// Which subnets should be selected: IPv4, IPv6, or use dual stack.
	IPFamily param.Field[InterfaceIPFamily] `json:"ip_family"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64]                                                                           `json:"port_group"`
	Type      param.Field[NewVmInterfaceSerializersPydanticNewInterfaceAnySubnetFipSerializerPydanticType] `json:"type"`
}

func (r NewVmInterfaceSerializersPydanticNewInterfaceAnySubnetFipSerializerPydanticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewVmInterfaceSerializersPydanticNewInterfaceAnySubnetFipSerializerPydanticParam) implementsNewVmInterfaceSerializersPydanticUnionParam() {
}

type NewVmInterfaceSerializersPydanticNewInterfaceAnySubnetFipSerializerPydanticType string

const (
	NewVmInterfaceSerializersPydanticNewInterfaceAnySubnetFipSerializerPydanticTypeAnySubnet NewVmInterfaceSerializersPydanticNewInterfaceAnySubnetFipSerializerPydanticType = "any_subnet"
)

func (r NewVmInterfaceSerializersPydanticNewInterfaceAnySubnetFipSerializerPydanticType) IsKnown() bool {
	switch r {
	case NewVmInterfaceSerializersPydanticNewInterfaceAnySubnetFipSerializerPydanticTypeAnySubnet:
		return true
	}
	return false
}

// Instance will be attached to the given port. Floating IP will be created and
// attached to that IP.
type NewVmInterfaceSerializersPydanticNewInterfaceReservedFixedIPFipSerializerPydanticParam struct {
	// Network ID the subnet belongs to. Port will be plugged in this network.
	PortID param.Field[string] `json:"port_id,required"`
	// Floating IP config for this subnet attachment.
	FloatingIP param.Field[NewInterfaceFloatingIPPydanticParam] `json:"floating_ip"`
	// Interface name
	InterfaceName param.Field[string] `json:"interface_name"`
	// Each group will be added to the separate trunk.
	PortGroup param.Field[int64]                                                                                 `json:"port_group"`
	Type      param.Field[NewVmInterfaceSerializersPydanticNewInterfaceReservedFixedIPFipSerializerPydanticType] `json:"type"`
}

func (r NewVmInterfaceSerializersPydanticNewInterfaceReservedFixedIPFipSerializerPydanticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewVmInterfaceSerializersPydanticNewInterfaceReservedFixedIPFipSerializerPydanticParam) implementsNewVmInterfaceSerializersPydanticUnionParam() {
}

type NewVmInterfaceSerializersPydanticNewInterfaceReservedFixedIPFipSerializerPydanticType string

const (
	NewVmInterfaceSerializersPydanticNewInterfaceReservedFixedIPFipSerializerPydanticTypeReservedFixedIP NewVmInterfaceSerializersPydanticNewInterfaceReservedFixedIPFipSerializerPydanticType = "reserved_fixed_ip"
)

func (r NewVmInterfaceSerializersPydanticNewInterfaceReservedFixedIPFipSerializerPydanticType) IsKnown() bool {
	switch r {
	case NewVmInterfaceSerializersPydanticNewInterfaceReservedFixedIPFipSerializerPydanticTypeReservedFixedIP:
		return true
	}
	return false
}

type NewVmInterfaceSerializersPydanticType string

const (
	NewVmInterfaceSerializersPydanticTypeExternal        NewVmInterfaceSerializersPydanticType = "external"
	NewVmInterfaceSerializersPydanticTypeSubnet          NewVmInterfaceSerializersPydanticType = "subnet"
	NewVmInterfaceSerializersPydanticTypeAnySubnet       NewVmInterfaceSerializersPydanticType = "any_subnet"
	NewVmInterfaceSerializersPydanticTypeReservedFixedIP NewVmInterfaceSerializersPydanticType = "reserved_fixed_ip"
)

func (r NewVmInterfaceSerializersPydanticType) IsKnown() bool {
	switch r {
	case NewVmInterfaceSerializersPydanticTypeExternal, NewVmInterfaceSerializersPydanticTypeSubnet, NewVmInterfaceSerializersPydanticTypeAnySubnet, NewVmInterfaceSerializersPydanticTypeReservedFixedIP:
		return true
	}
	return false
}

type CloudV1PricingAIClusterPreviewAIClusterPriceParams struct {
	// AI cluster flavor
	Flavor param.Field[string] `json:"flavor,required"`
	// Subnet IPs and floating IPs
	Interfaces param.Field[[]NewVmInterfaceSerializersPydanticUnionParam] `json:"interfaces,required"`
	// AI cluster name
	Name param.Field[string] `json:"name,required"`
	// A multiplier that would be used to calculate the price, multiplying the flavor,
	// IP, and volume prices by this number
	InstancesCount param.Field[int64] `json:"instances_count"`
	// Volumes attached to the AI cluster
	Volumes param.Field[[]InstanceVolumePricingRequestParam] `json:"volumes"`
}

func (r CloudV1PricingAIClusterPreviewAIClusterPriceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
