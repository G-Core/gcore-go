// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1AIService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1AIService] method instead.
type CloudV1AIService struct {
	Options  []option.RequestOption
	Images   *CloudV1AIImageService
	Clusters *CloudV1AIClusterService
}

// NewCloudV1AIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1AIService(opts ...option.RequestOption) (r *CloudV1AIService) {
	r = &CloudV1AIService{}
	r.Options = opts
	r.Images = NewCloudV1AIImageService(opts...)
	r.Clusters = NewCloudV1AIClusterService(opts...)
	return
}

// Get flavors available for a potential AI cluster
func (r *CloudV1AIService) GetFlavors(ctx context.Context, projectID string, regionID string, query CloudV1AIGetFlavorsParams, opts ...option.RequestOption) (res *CloudV1AIGetFlavorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if regionID == "" {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/flavors/%s/%s", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List available regions that provide AI cluster support
func (r *CloudV1AIService) ListRegions(ctx context.Context, opts ...option.RequestOption) (res *CloudV1AIListRegionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/ai/regions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ServiceEndpoint struct {
	// Endpoint ID
	ID int64 `json:"id,required"`
	// List of client IDs to access
	ClientIDs []int64 `json:"client_ids,required"`
	// The endpoint creation datetime timestamp
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Endpoint service
	Service EndpointService `json:"service,required"`
	// The endpoint last update datetime timestamp
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Endpoint URL
	URL string `json:"url,required"`
	// Service version
	Version string              `json:"version,required"`
	JSON    serviceEndpointJSON `json:"-"`
}

// serviceEndpointJSON contains the JSON metadata for the struct [ServiceEndpoint]
type serviceEndpointJSON struct {
	ID          apijson.Field
	ClientIDs   apijson.Field
	CreatedOn   apijson.Field
	RegionID    apijson.Field
	Service     apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEndpointJSON) RawJSON() string {
	return r.raw
}

type CloudV1AIGetFlavorsResponse struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []CloudV1AIGetFlavorsResponseResult `json:"results"`
	JSON    cloudV1AIGetFlavorsResponseJSON     `json:"-"`
}

// cloudV1AIGetFlavorsResponseJSON contains the JSON metadata for the struct
// [CloudV1AIGetFlavorsResponse]
type cloudV1AIGetFlavorsResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1AIGetFlavorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1AIGetFlavorsResponseJSON) RawJSON() string {
	return r.raw
}

// AI Flavor Object
type CloudV1AIGetFlavorsResponseResult struct {
	// Number of available instances of given configuration
	Capacity int64 `json:"capacity"`
	// Currency code. Shown if the include_prices query parameter if set to true
	CurrencyCode string `json:"currency_code"`
	// Disabled flavor flag
	Disabled bool `json:"disabled"`
	// Short ID
	FlavorID string `json:"flavor_id"`
	// Human readable name
	FlavorName string `json:"flavor_name"`
	// Various hardware hints for flavor.
	HardwareDescription DeprecatedFlavorHardwareDescription `json:"hardware_description"`
	// Price per hour. Shown if the include_prices query parameter if set to true
	PricePerHour float64 `json:"price_per_hour"`
	// Price per month. Shown if the include_prices query parameter if set to true
	PricePerMonth float64 `json:"price_per_month"`
	// Price status for the UI
	PriceStatus CloudV1AIGetFlavorsResponseResultsPriceStatus `json:"price_status"`
	// Flavor resource class for mapping to hardware capacity
	ResourceClass string                                `json:"resource_class"`
	JSON          cloudV1AIGetFlavorsResponseResultJSON `json:"-"`
}

// cloudV1AIGetFlavorsResponseResultJSON contains the JSON metadata for the struct
// [CloudV1AIGetFlavorsResponseResult]
type cloudV1AIGetFlavorsResponseResultJSON struct {
	Capacity            apijson.Field
	CurrencyCode        apijson.Field
	Disabled            apijson.Field
	FlavorID            apijson.Field
	FlavorName          apijson.Field
	HardwareDescription apijson.Field
	PricePerHour        apijson.Field
	PricePerMonth       apijson.Field
	PriceStatus         apijson.Field
	ResourceClass       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudV1AIGetFlavorsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1AIGetFlavorsResponseResultJSON) RawJSON() string {
	return r.raw
}

// Price status for the UI
type CloudV1AIGetFlavorsResponseResultsPriceStatus string

const (
	CloudV1AIGetFlavorsResponseResultsPriceStatusError CloudV1AIGetFlavorsResponseResultsPriceStatus = "error"
	CloudV1AIGetFlavorsResponseResultsPriceStatusHide  CloudV1AIGetFlavorsResponseResultsPriceStatus = "hide"
	CloudV1AIGetFlavorsResponseResultsPriceStatusShow  CloudV1AIGetFlavorsResponseResultsPriceStatus = "show"
)

func (r CloudV1AIGetFlavorsResponseResultsPriceStatus) IsKnown() bool {
	switch r {
	case CloudV1AIGetFlavorsResponseResultsPriceStatusError, CloudV1AIGetFlavorsResponseResultsPriceStatusHide, CloudV1AIGetFlavorsResponseResultsPriceStatusShow:
		return true
	}
	return false
}

type CloudV1AIListRegionsResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1AIListRegionsResponseResult `json:"results,required"`
	JSON    cloudV1AIListRegionsResponseJSON     `json:"-"`
}

// cloudV1AIListRegionsResponseJSON contains the JSON metadata for the struct
// [CloudV1AIListRegionsResponse]
type cloudV1AIListRegionsResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1AIListRegionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1AIListRegionsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1AIListRegionsResponseResult struct {
	// Region ID. It automatically is being generated when it is created
	ID int64 `json:"id,required"`
	// The access level of the region in string format.
	AccessLevel RegionAccessType `json:"access_level,required"`
	// A datetime object that is automatically generated when the region is created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Unique constraint. Display name
	DisplayName string `json:"display_name,required"`
	// Endpoint type
	EndpointType RegionEndpointType `json:"endpoint_type,required"`
	// External network ID for Neutron
	ExternalNetworkID string `json:"external_network_id,required"`
	// Flag indicating if a region has AI capabilities available
	HasAI bool `json:"has_ai,required"`
	// Flag indicating if a region has AI GPUs available
	HasAIGPU bool `json:"has_ai_gpu,required"`
	// Region has bare metal capability
	HasBaremetal bool `json:"has_baremetal,required"`
	// Region has basic vm capability
	HasBasicVm bool `json:"has_basic_vm,required"`
	// Flag indicating if a region is managed kubernetes available
	HasK8s bool `json:"has_k8s,required"`
	// Region has KVM virtualization capability
	HasKvm bool `json:"has_kvm,required"`
	// Region has SFS capability
	HasSfs bool `json:"has_sfs,required"`
	// Keystone object in the JSON format
	Keystone CloudV1AIListRegionsResponseResultsKeystone `json:"keystone,required"`
	// Foreign key for Keystone entity
	KeystoneID int64 `json:"keystone_id,required"`
	// Region name exactly as stated in keystone
	KeystoneName string `json:"keystone_name,required"`
	// Metrics database object in JSON format
	MetricsDatabase CloudV1AIListRegionsResponseResultsMetricsDatabase `json:"metrics_database,required,nullable"`
	// NoVNC console proxy URL
	NovncProxyURL string `json:"novnc_proxy_url,required,nullable"`
	// Serial console proxy URL
	SerialProxyURL string `json:"serial_proxy_url,required,nullable"`
	// Region state
	State RegionState `json:"state,required"`
	// Physical network name to create vlan networks
	VlanPhysicalNetwork string `json:"vlan_physical_network,required"`
	// Zone
	Zone Zone `json:"zone,required,nullable"`
	// AI service API endpoint object
	AIServiceEndpoint ServiceEndpoint `json:"ai_service_endpoint,nullable"`
	// AI service API endpoint ID
	AIServiceEndpointID int64 `json:"ai_service_endpoint_id,nullable"`
	// List of available volume types, 'standard', 'ssd_hiiops', 'cold'].
	AvailableVolumeTypes []string `json:"available_volume_types,nullable"`
	// Set of available volume types, 'standard', 'ssd_hiiops', 'cold'].
	AvailableVolumeTypesSet []string `json:"available_volume_types_set,nullable"`
	// Country
	Country string `json:"country,nullable"`
	// DDoS endpoint object
	DdosEndpoint ServiceEndpoint `json:"ddos_endpoint,nullable"`
	// DDoS endpoint ID
	DdosEndpointID int64 `json:"ddos_endpoint_id,nullable"`
	// Foreign key to Metrics database entity
	MetricsDatabaseID int64 `json:"metrics_database_id,nullable"`
	// Spice console proxy URL
	SpiceProxyURL string `json:"spice_proxy_url,nullable"`
	// UUID. Id of the Task entity that is handling the region's state transition
	TaskID string                                 `json:"task_id,nullable"`
	JSON   cloudV1AIListRegionsResponseResultJSON `json:"-"`
}

// cloudV1AIListRegionsResponseResultJSON contains the JSON metadata for the struct
// [CloudV1AIListRegionsResponseResult]
type cloudV1AIListRegionsResponseResultJSON struct {
	ID                      apijson.Field
	AccessLevel             apijson.Field
	CreatedOn               apijson.Field
	DisplayName             apijson.Field
	EndpointType            apijson.Field
	ExternalNetworkID       apijson.Field
	HasAI                   apijson.Field
	HasAIGPU                apijson.Field
	HasBaremetal            apijson.Field
	HasBasicVm              apijson.Field
	HasK8s                  apijson.Field
	HasKvm                  apijson.Field
	HasSfs                  apijson.Field
	Keystone                apijson.Field
	KeystoneID              apijson.Field
	KeystoneName            apijson.Field
	MetricsDatabase         apijson.Field
	NovncProxyURL           apijson.Field
	SerialProxyURL          apijson.Field
	State                   apijson.Field
	VlanPhysicalNetwork     apijson.Field
	Zone                    apijson.Field
	AIServiceEndpoint       apijson.Field
	AIServiceEndpointID     apijson.Field
	AvailableVolumeTypes    apijson.Field
	AvailableVolumeTypesSet apijson.Field
	Country                 apijson.Field
	DdosEndpoint            apijson.Field
	DdosEndpointID          apijson.Field
	MetricsDatabaseID       apijson.Field
	SpiceProxyURL           apijson.Field
	TaskID                  apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CloudV1AIListRegionsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1AIListRegionsResponseResultJSON) RawJSON() string {
	return r.raw
}

// Keystone object in the JSON format
type CloudV1AIListRegionsResponseResultsKeystone struct {
	// ID of Keystone federated domain that was created during initialization
	KeystoneFederatedDomainID string `json:"keystone_federated_domain_id,required"`
	// Keystone state
	State CloudV1AIListRegionsResponseResultsKeystoneState `json:"state,required"`
	// Unique constraint. Keystone endpoint to use when interacting with Openstack API
	URL string `json:"url,required"`
	// Username for admin access
	Username string `json:"username,required"`
	// An automatically generated Region ID when it is created.
	ID int64 `json:"id,nullable"`
	// Datetime object. It automatically is being generated when it is created
	CreatedOn time.Time                                       `json:"created_on,nullable" format:"date-time"`
	JSON      cloudV1AIListRegionsResponseResultsKeystoneJSON `json:"-"`
}

// cloudV1AIListRegionsResponseResultsKeystoneJSON contains the JSON metadata for
// the struct [CloudV1AIListRegionsResponseResultsKeystone]
type cloudV1AIListRegionsResponseResultsKeystoneJSON struct {
	KeystoneFederatedDomainID apijson.Field
	State                     apijson.Field
	URL                       apijson.Field
	Username                  apijson.Field
	ID                        apijson.Field
	CreatedOn                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CloudV1AIListRegionsResponseResultsKeystone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1AIListRegionsResponseResultsKeystoneJSON) RawJSON() string {
	return r.raw
}

// Keystone state
type CloudV1AIListRegionsResponseResultsKeystoneState string

const (
	CloudV1AIListRegionsResponseResultsKeystoneStateDeleted              CloudV1AIListRegionsResponseResultsKeystoneState = "DELETED"
	CloudV1AIListRegionsResponseResultsKeystoneStateInitializationFailed CloudV1AIListRegionsResponseResultsKeystoneState = "INITIALIZATION_FAILED"
	CloudV1AIListRegionsResponseResultsKeystoneStateInitialized          CloudV1AIListRegionsResponseResultsKeystoneState = "INITIALIZED"
	CloudV1AIListRegionsResponseResultsKeystoneStateNew                  CloudV1AIListRegionsResponseResultsKeystoneState = "NEW"
)

func (r CloudV1AIListRegionsResponseResultsKeystoneState) IsKnown() bool {
	switch r {
	case CloudV1AIListRegionsResponseResultsKeystoneStateDeleted, CloudV1AIListRegionsResponseResultsKeystoneStateInitializationFailed, CloudV1AIListRegionsResponseResultsKeystoneStateInitialized, CloudV1AIListRegionsResponseResultsKeystoneStateNew:
		return true
	}
	return false
}

// Metrics database object in JSON format
type CloudV1AIListRegionsResponseResultsMetricsDatabase struct {
	// ID of the metrics database
	ID int64 `json:"id,required"`
	// Metrics database host
	Host string `json:"host,required"`
	// Metrics database port
	Port int64 `json:"port,required"`
	// Metrics database username
	Username string                                                 `json:"username,required"`
	JSON     cloudV1AIListRegionsResponseResultsMetricsDatabaseJSON `json:"-"`
}

// cloudV1AIListRegionsResponseResultsMetricsDatabaseJSON contains the JSON
// metadata for the struct [CloudV1AIListRegionsResponseResultsMetricsDatabase]
type cloudV1AIListRegionsResponseResultsMetricsDatabaseJSON struct {
	ID          apijson.Field
	Host        apijson.Field
	Port        apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1AIListRegionsResponseResultsMetricsDatabase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1AIListRegionsResponseResultsMetricsDatabaseJSON) RawJSON() string {
	return r.raw
}

type CloudV1AIGetFlavorsParams struct {
	// Flag for filtering disabled flavors in the region
	Disabled param.Field[bool] `query:"disabled"`
	// Set to true to see the number of available AI Infrastructure clusters of given
	// configuration
	IncludeCapacity param.Field[bool] `query:"include_capacity"`
	// Set to true if flavor listing should include flavor prices
	IncludePrices param.Field[bool] `query:"include_prices"`
}

// URLQuery serializes [CloudV1AIGetFlavorsParams]'s query parameters as
// `url.Values`.
func (r CloudV1AIGetFlavorsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
