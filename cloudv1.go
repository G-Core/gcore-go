// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/shared"
	"github.com/tidwall/gjson"
)

// CloudV1Service contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1Service] method instead.
type CloudV1Service struct {
	Options               []option.RequestOption
	Bmflavors             *CloudV1BmflavorService
	Bminstances           *CloudV1BminstanceService
	Caas                  *CloudV1CaaService
	Pricing               *CloudV1PricingService
	Registries            *CloudV1RegistryService
	CostReport            *CloudV1CostReportService
	ReservationCostReport *CloudV1ReservationCostReportService
	Ddos                  *CloudV1DdoService
	FileShares            *CloudV1FileShareService
	Securitygrouprules    *CloudV1SecuritygroupruleService
	Securitygroups        *CloudV1SecuritygroupService
	Floatingips           *CloudV1FloatingipService
	Faas                  *CloudV1FaaService
	AI                    *CloudV1AIService
	Bmimages              *CloudV1BmimageService
	Images                *CloudV1ImageService
	ResellerImage         *CloudV1ResellerImageService
	SharedImage           *CloudV1SharedImageService
	Apptemplates          *CloudV1ApptemplateService
	Instances             *CloudV1InstanceService
	Ports                 *CloudV1PortService
	ResellerNameTemplates *CloudV1ResellerNameTemplateService
	L7policies            *CloudV1L7policyService
	Lblisteners           *CloudV1LblistenerService
	Lbpools               *CloudV1LbpoolService
	Loadbalancers         *CloudV1LoadbalancerService
	Laas                  *CloudV1LaaService
	K8s                   *CloudV1K8Service
	Dbaas                 *CloudV1DbaaService
	Networks              *CloudV1NetworkService
	Subnets               *CloudV1SubnetService
	Servergroups          *CloudV1ServergroupService
	Projects              *CloudV1ProjectService
	Regions               *CloudV1RegionService
	ResellerRegion        *CloudV1ResellerRegionService
	Reservations          *CloudV1ReservationService
	ReservedFixedIPs      *CloudV1ReservedFixedIPService
	Routers               *CloudV1RouterService
	Keypairs              *CloudV1KeypairService
	Secrets               *CloudV1SecretService
	LifecyclePolicy       *CloudV1LifecyclePolicyService
	Schedule              *CloudV1ScheduleService
	Snapshots             *CloudV1SnapshotService
	Tasks                 *CloudV1TaskService
	UserActions           *CloudV1UserActionService
	Users                 *CloudV1UserService
	Volumes               *CloudV1VolumeService
}

// NewCloudV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCloudV1Service(opts ...option.RequestOption) (r *CloudV1Service) {
	r = &CloudV1Service{}
	r.Options = opts
	r.Bmflavors = NewCloudV1BmflavorService(opts...)
	r.Bminstances = NewCloudV1BminstanceService(opts...)
	r.Caas = NewCloudV1CaaService(opts...)
	r.Pricing = NewCloudV1PricingService(opts...)
	r.Registries = NewCloudV1RegistryService(opts...)
	r.CostReport = NewCloudV1CostReportService(opts...)
	r.ReservationCostReport = NewCloudV1ReservationCostReportService(opts...)
	r.Ddos = NewCloudV1DdoService(opts...)
	r.FileShares = NewCloudV1FileShareService(opts...)
	r.Securitygrouprules = NewCloudV1SecuritygroupruleService(opts...)
	r.Securitygroups = NewCloudV1SecuritygroupService(opts...)
	r.Floatingips = NewCloudV1FloatingipService(opts...)
	r.Faas = NewCloudV1FaaService(opts...)
	r.AI = NewCloudV1AIService(opts...)
	r.Bmimages = NewCloudV1BmimageService(opts...)
	r.Images = NewCloudV1ImageService(opts...)
	r.ResellerImage = NewCloudV1ResellerImageService(opts...)
	r.SharedImage = NewCloudV1SharedImageService(opts...)
	r.Apptemplates = NewCloudV1ApptemplateService(opts...)
	r.Instances = NewCloudV1InstanceService(opts...)
	r.Ports = NewCloudV1PortService(opts...)
	r.ResellerNameTemplates = NewCloudV1ResellerNameTemplateService(opts...)
	r.L7policies = NewCloudV1L7policyService(opts...)
	r.Lblisteners = NewCloudV1LblistenerService(opts...)
	r.Lbpools = NewCloudV1LbpoolService(opts...)
	r.Loadbalancers = NewCloudV1LoadbalancerService(opts...)
	r.Laas = NewCloudV1LaaService(opts...)
	r.K8s = NewCloudV1K8Service(opts...)
	r.Dbaas = NewCloudV1DbaaService(opts...)
	r.Networks = NewCloudV1NetworkService(opts...)
	r.Subnets = NewCloudV1SubnetService(opts...)
	r.Servergroups = NewCloudV1ServergroupService(opts...)
	r.Projects = NewCloudV1ProjectService(opts...)
	r.Regions = NewCloudV1RegionService(opts...)
	r.ResellerRegion = NewCloudV1ResellerRegionService(opts...)
	r.Reservations = NewCloudV1ReservationService(opts...)
	r.ReservedFixedIPs = NewCloudV1ReservedFixedIPService(opts...)
	r.Routers = NewCloudV1RouterService(opts...)
	r.Keypairs = NewCloudV1KeypairService(opts...)
	r.Secrets = NewCloudV1SecretService(opts...)
	r.LifecyclePolicy = NewCloudV1LifecyclePolicyService(opts...)
	r.Schedule = NewCloudV1ScheduleService(opts...)
	r.Snapshots = NewCloudV1SnapshotService(opts...)
	r.Tasks = NewCloudV1TaskService(opts...)
	r.UserActions = NewCloudV1UserActionService(opts...)
	r.Users = NewCloudV1UserService(opts...)
	r.Volumes = NewCloudV1VolumeService(opts...)
	return
}

// Get available floating IPs
func (r *CloudV1Service) GetAvailableFloatingIPs(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *CloudV1GetAvailableFloatingIPsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/availablefloatingips/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve the list of directly-attachable networks with subnet details. Returned
// entities may or may not be owned by the project
func (r *CloudV1Service) GetAvailableNetworks(ctx context.Context, projectID int64, regionID int64, query CloudV1GetAvailableNetworksParams, opts ...option.RequestOption) (res *CloudV1GetAvailableNetworksResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/availablenetworks/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get amount of available bare metal nodes It's similar v1/bmcapacity just with
// <project_id> in path
func (r *CloudV1Service) GetBmCapacity(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *BaremetalCapacity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/bmcapacity/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List of bare metal flavors that are available for reservation
func (r *CloudV1Service) GetBmReservationFlavors(ctx context.Context, regionID int64, query CloudV1GetBmReservationFlavorsParams, opts ...option.RequestOption) (res *[]CloudV1GetBmReservationFlavorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/bm_reservation_flavors/%v", regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a list of flavors. When the include_prices query parameter is
// specified, the list shows prices. A client in trial mode gets all price values
// as 0. If you get Pricing Error contact the support
func (r *CloudV1Service) GetFlavors(ctx context.Context, projectID int64, regionID int64, query CloudV1GetFlavorsParams, opts ...option.RequestOption) (res *FlavorList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/flavors/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a list of load balancer flavors. When the include_prices query
// parameter is specified, the list shows prices. A client in trial mode gets all
// price values as 0. If you get Pricing Error contact the support
func (r *CloudV1Service) GetLbFlavors(ctx context.Context, projectID int64, regionID int64, query CloudV1GetLbFlavorsParams, opts ...option.RequestOption) (res *CloudV1GetLbFlavorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/lbflavors/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Client available endpoints
func (r *CloudV1Service) GetServiceEndpoints(ctx context.Context, query CloudV1GetServiceEndpointsParams, opts ...option.RequestOption) (res *CloudV1GetServiceEndpointsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/service_endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Receiving data from the past hour might lead to incomplete statistics. For the
// most accurate data, we recommend accessing the statistics after at least one
// hour. Typically, updates are available within a 24-hour period, although the
// frequency can vary. Maintenance periods or other exceptions may cause delays,
// potentially extending beyond 24 hours until the servers are back online and the
// missing data is filled in.
func (r *CloudV1Service) GetUsageReport(ctx context.Context, body CloudV1GetUsageReportParams, opts ...option.RequestOption) (res *CloudV1GetUsageReportResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/usage_report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List images owned by project
func (r *CloudV1Service) ListProjectImages(ctx context.Context, projectID int64, regionID int64, query CloudV1ListProjectImagesParams, opts ...option.RequestOption) (res *ImageList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/projectimages/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upload image
func (r *CloudV1Service) UploadImage(ctx context.Context, projectID int64, regionID int64, body CloudV1UploadImageParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/downloadimage/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Available bare metal nodes count per flavor
type BaremetalCapacity struct {
	// Bare metal capacity dict. Bare metal flavors are the keys of this dictionary,
	// available nodes count are corresponding values
	Capacity map[string]interface{} `json:"capacity"`
	JSON     baremetalCapacityJSON  `json:"-"`
}

// baremetalCapacityJSON contains the JSON metadata for the struct
// [BaremetalCapacity]
type baremetalCapacityJSON struct {
	Capacity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BaremetalCapacity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r baremetalCapacityJSON) RawJSON() string {
	return r.raw
}

// Various hardware hints for flavor.
type DeprecatedFlavorHardwareDescription struct {
	// Human-readable CPU description
	CPU string `json:"cpu"`
	// Human-readable disk description
	Disk string `json:"disk"`
	// Human-readable ephemeral disk description
	Ephemeral string `json:"ephemeral"`
	// Human-readable GPU description
	GPU string `json:"gpu"`
	// Human-readable IPU description of AI cluster
	Ipu string `json:"ipu"`
	// Human-readable NIC description
	Network string `json:"network"`
	// Human-readable count of poplar servers of AI cluster
	PoplarCount string `json:"poplar_count"`
	// Human-readable RAM description
	Ram  string                                  `json:"ram"`
	JSON deprecatedFlavorHardwareDescriptionJSON `json:"-"`
}

// deprecatedFlavorHardwareDescriptionJSON contains the JSON metadata for the
// struct [DeprecatedFlavorHardwareDescription]
type deprecatedFlavorHardwareDescriptionJSON struct {
	CPU         apijson.Field
	Disk        apijson.Field
	Ephemeral   apijson.Field
	GPU         apijson.Field
	Ipu         apijson.Field
	Network     apijson.Field
	PoplarCount apijson.Field
	Ram         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeprecatedFlavorHardwareDescription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deprecatedFlavorHardwareDescriptionJSON) RawJSON() string {
	return r.raw
}

type EndpointService string

const (
	EndpointServiceAIService                EndpointService = "ai_service"
	EndpointServiceCaas                     EndpointService = "caas"
	EndpointServiceCaasForInference         EndpointService = "caas_for_inference"
	EndpointServiceDbaas                    EndpointService = "dbaas"
	EndpointServiceDdos                     EndpointService = "ddos"
	EndpointServiceFaas                     EndpointService = "faas"
	EndpointServiceInferenceAtTheEdge       EndpointService = "inference_at_the_edge"
	EndpointServiceInferenceBox             EndpointService = "inference_box"
	EndpointServiceLaas                     EndpointService = "laas"
	EndpointServiceMkaas                    EndpointService = "mkaas"
	EndpointServiceRegistry                 EndpointService = "registry"
	EndpointServiceVipuControllerManagerAPI EndpointService = "vipu_controller_manager_api"
)

func (r EndpointService) IsKnown() bool {
	switch r {
	case EndpointServiceAIService, EndpointServiceCaas, EndpointServiceCaasForInference, EndpointServiceDbaas, EndpointServiceDdos, EndpointServiceFaas, EndpointServiceInferenceAtTheEdge, EndpointServiceInferenceBox, EndpointServiceLaas, EndpointServiceMkaas, EndpointServiceRegistry, EndpointServiceVipuControllerManagerAPI:
		return true
	}
	return false
}

type FlavorList struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []FlavorListResult `json:"results"`
	JSON    flavorListJSON     `json:"-"`
}

// flavorListJSON contains the JSON metadata for the struct [FlavorList]
type flavorListJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FlavorList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r flavorListJSON) RawJSON() string {
	return r.raw
}

// Flavor object with price
type FlavorListResult struct {
	// Flavor architecture type
	Architecture string `json:"architecture"`
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
	// Flavor operating system
	OsType string `json:"os_type"`
	// Price per hour. Shown if the include_prices query parameter if set to true
	PricePerHour float64 `json:"price_per_hour"`
	// Price per month. Shown if the include_prices query parameter if set to true
	PricePerMonth float64 `json:"price_per_month"`
	// Price status for the UI
	PriceStatus FlavorListResultsPriceStatus `json:"price_status"`
	// RAM size in MiB
	Ram int64 `json:"ram"`
	// Flavor resource class for mapping to hardware capacity
	ResourceClass string `json:"resource_class"`
	// Virtual CPU count
	Vcpus int64                `json:"vcpus"`
	JSON  flavorListResultJSON `json:"-"`
}

// flavorListResultJSON contains the JSON metadata for the struct
// [FlavorListResult]
type flavorListResultJSON struct {
	Architecture        apijson.Field
	CurrencyCode        apijson.Field
	Disabled            apijson.Field
	FlavorID            apijson.Field
	FlavorName          apijson.Field
	HardwareDescription apijson.Field
	OsType              apijson.Field
	PricePerHour        apijson.Field
	PricePerMonth       apijson.Field
	PriceStatus         apijson.Field
	Ram                 apijson.Field
	ResourceClass       apijson.Field
	Vcpus               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *FlavorListResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r flavorListResultJSON) RawJSON() string {
	return r.raw
}

// Price status for the UI
type FlavorListResultsPriceStatus string

const (
	FlavorListResultsPriceStatusError FlavorListResultsPriceStatus = "error"
	FlavorListResultsPriceStatusHide  FlavorListResultsPriceStatus = "hide"
	FlavorListResultsPriceStatusShow  FlavorListResultsPriceStatus = "show"
)

func (r FlavorListResultsPriceStatus) IsKnown() bool {
	switch r {
	case FlavorListResultsPriceStatusError, FlavorListResultsPriceStatusHide, FlavorListResultsPriceStatusShow:
		return true
	}
	return false
}

type CloudV1GetAvailableFloatingIPsResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []FloatingIP                               `json:"results,required"`
	JSON    cloudV1GetAvailableFloatingIPsResponseJSON `json:"-"`
}

// cloudV1GetAvailableFloatingIPsResponseJSON contains the JSON metadata for the
// struct [CloudV1GetAvailableFloatingIPsResponse]
type cloudV1GetAvailableFloatingIPsResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1GetAvailableFloatingIPsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetAvailableFloatingIPsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1GetAvailableNetworksResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1GetAvailableNetworksResponseResult `json:"results,required"`
	JSON    cloudV1GetAvailableNetworksResponseJSON     `json:"-"`
}

// cloudV1GetAvailableNetworksResponseJSON contains the JSON metadata for the
// struct [CloudV1GetAvailableNetworksResponse]
type cloudV1GetAvailableNetworksResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1GetAvailableNetworksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetAvailableNetworksResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1GetAvailableNetworksResponseResult struct {
	// Network ID
	ID string `json:"id,required" format:"uuid4"`
	// Datetime when the network was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// True if the network `router:external` attribute
	External bool `json:"external,required"`
	// Network name
	Name string `json:"name,required"`
	// Indicates port_security_enabled status of all newly created in the network
	// ports.
	PortSecurityEnabled bool `json:"port_security_enabled,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// True when the network is shared with your project by external owner
	Shared bool `json:"shared,required"`
	// Network type (vlan, vxlan)
	Type string `json:"type,required"`
	// Datetime when the network was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,nullable" format:"uuid4"`
	// True if network has is_default attribute
	Default bool `json:"default,nullable"`
	// Network metadata
	Metadata []DetailedMetadata `json:"metadata"`
	// MTU (maximum transmission unit). Default value is 1450
	Mtu int64 `json:"mtu"`
	// Project ID
	ProjectID int64 `json:"project_id,nullable"`
	// Id of network segment
	SegmentationID int64    `json:"segmentation_id,nullable"`
	Subnets        []Subnet `json:"subnets,nullable"`
	// Active task. If None, action has been performed immediately in the request
	// itself.
	TaskID string                                        `json:"task_id,nullable" format:"uuid4"`
	JSON   cloudV1GetAvailableNetworksResponseResultJSON `json:"-"`
}

// cloudV1GetAvailableNetworksResponseResultJSON contains the JSON metadata for the
// struct [CloudV1GetAvailableNetworksResponseResult]
type cloudV1GetAvailableNetworksResponseResultJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	External            apijson.Field
	Name                apijson.Field
	PortSecurityEnabled apijson.Field
	Region              apijson.Field
	RegionID            apijson.Field
	Shared              apijson.Field
	Type                apijson.Field
	UpdatedAt           apijson.Field
	CreatorTaskID       apijson.Field
	Default             apijson.Field
	Metadata            apijson.Field
	Mtu                 apijson.Field
	ProjectID           apijson.Field
	SegmentationID      apijson.Field
	Subnets             apijson.Field
	TaskID              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudV1GetAvailableNetworksResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetAvailableNetworksResponseResultJSON) RawJSON() string {
	return r.raw
}

// Bare metal flavors that are available for reservations
type CloudV1GetBmReservationFlavorsResponse struct {
	// Description using period for calc slices [`year`, `month`, `day`]
	ActivityPeriod string `json:"activity_period"`
	// List of the full reservation period by `activity_period` length
	ActivityPeriodLength []int64 `json:"activity_period_length"`
	// Flavor architecture type
	Architecture string `json:"architecture"`
	// Count of available nodes
	AvailableCount int64 `json:"available_count"`
	// Disabled flavor flag
	Disabled bool `json:"disabled"`
	// Short ID
	FlavorID string `json:"flavor_id"`
	// Human readable name
	FlavorName string `json:"flavor_name"`
	// Various hardware hints for flavor.
	HardwareDescription DeprecatedFlavorHardwareDescription `json:"hardware_description"`
	// Flavor operating system
	OsType string `json:"os_type"`
	// RAM size in MiB
	Ram int64 `json:"ram"`
	// Flavor resource class for mapping to hardware capacity
	ResourceClass string `json:"resource_class"`
	// Virtual CPU count
	Vcpus int64                                      `json:"vcpus"`
	JSON  cloudV1GetBmReservationFlavorsResponseJSON `json:"-"`
}

// cloudV1GetBmReservationFlavorsResponseJSON contains the JSON metadata for the
// struct [CloudV1GetBmReservationFlavorsResponse]
type cloudV1GetBmReservationFlavorsResponseJSON struct {
	ActivityPeriod       apijson.Field
	ActivityPeriodLength apijson.Field
	Architecture         apijson.Field
	AvailableCount       apijson.Field
	Disabled             apijson.Field
	FlavorID             apijson.Field
	FlavorName           apijson.Field
	HardwareDescription  apijson.Field
	OsType               apijson.Field
	Ram                  apijson.Field
	ResourceClass        apijson.Field
	Vcpus                apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CloudV1GetBmReservationFlavorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetBmReservationFlavorsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1GetLbFlavorsResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1GetLbFlavorsResponseResult `json:"results,required"`
	JSON    cloudV1GetLbFlavorsResponseJSON     `json:"-"`
}

// cloudV1GetLbFlavorsResponseJSON contains the JSON metadata for the struct
// [CloudV1GetLbFlavorsResponse]
type cloudV1GetLbFlavorsResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1GetLbFlavorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetLbFlavorsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1GetLbFlavorsResponseResult struct {
	// Short ID
	FlavorID string `json:"flavor_id,required"`
	// Human readable name
	FlavorName string `json:"flavor_name,required"`
	// Additional hardware description.
	HardwareDescription CloudV1GetLbFlavorsResponseResultsHardwareDescription `json:"hardware_description,required"`
	// RAM size in MiB
	Ram int64 `json:"ram,required"`
	// Virtual CPU count
	Vcpus int64 `json:"vcpus,required"`
	// Currency code. Shown if the include_prices query parameter if set to true
	CurrencyCode string `json:"currency_code,nullable"`
	// Price per hour. Shown if the include_prices query parameter if set to true
	PricePerHour float64 `json:"price_per_hour,nullable"`
	// Price per month. Shown if the include_prices query parameter if set to true
	PricePerMonth float64 `json:"price_per_month,nullable"`
	// Price status for the UI
	PriceStatus PriceDisplayStatus                    `json:"price_status,nullable"`
	JSON        cloudV1GetLbFlavorsResponseResultJSON `json:"-"`
}

// cloudV1GetLbFlavorsResponseResultJSON contains the JSON metadata for the struct
// [CloudV1GetLbFlavorsResponseResult]
type cloudV1GetLbFlavorsResponseResultJSON struct {
	FlavorID            apijson.Field
	FlavorName          apijson.Field
	HardwareDescription apijson.Field
	Ram                 apijson.Field
	Vcpus               apijson.Field
	CurrencyCode        apijson.Field
	PricePerHour        apijson.Field
	PricePerMonth       apijson.Field
	PriceStatus         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CloudV1GetLbFlavorsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetLbFlavorsResponseResultJSON) RawJSON() string {
	return r.raw
}

type CloudV1GetLbFlavorsResponseResultsHardwareDescription struct {
	// Human-readable CPU description
	CPU string `json:"cpu,nullable"`
	// Human-readable disk description
	Disk string `json:"disk,nullable"`
	// Human-readable ephemeral disk description
	Ephemeral string `json:"ephemeral,nullable"`
	// Human-readable GPU description
	GPU string `json:"gpu,nullable"`
	// Human-readable IPU description of AI cluster
	Ipu string `json:"ipu,nullable"`
	// Human-readable NIC description
	Network string `json:"network,nullable"`
	// Human-readable count of poplar servers of AI cluster
	PoplarCount int64 `json:"poplar_count,nullable"`
	// Human-readable RAM description
	Ram  string                                                    `json:"ram,nullable"`
	JSON cloudV1GetLbFlavorsResponseResultsHardwareDescriptionJSON `json:"-"`
}

// cloudV1GetLbFlavorsResponseResultsHardwareDescriptionJSON contains the JSON
// metadata for the struct [CloudV1GetLbFlavorsResponseResultsHardwareDescription]
type cloudV1GetLbFlavorsResponseResultsHardwareDescriptionJSON struct {
	CPU         apijson.Field
	Disk        apijson.Field
	Ephemeral   apijson.Field
	GPU         apijson.Field
	Ipu         apijson.Field
	Network     apijson.Field
	PoplarCount apijson.Field
	Ram         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1GetLbFlavorsResponseResultsHardwareDescription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetLbFlavorsResponseResultsHardwareDescriptionJSON) RawJSON() string {
	return r.raw
}

type CloudV1GetServiceEndpointsResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1GetServiceEndpointsResponseResult `json:"results,required"`
	JSON    cloudV1GetServiceEndpointsResponseJSON     `json:"-"`
}

// cloudV1GetServiceEndpointsResponseJSON contains the JSON metadata for the struct
// [CloudV1GetServiceEndpointsResponse]
type cloudV1GetServiceEndpointsResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1GetServiceEndpointsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetServiceEndpointsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1GetServiceEndpointsResponseResult struct {
	// The endpoint creation datetime timestamp
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Endpoint service
	Service EndpointService `json:"service,required"`
	// The endpoint last update datetime timestamp
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Service version
	Version string                                       `json:"version,required"`
	JSON    cloudV1GetServiceEndpointsResponseResultJSON `json:"-"`
}

// cloudV1GetServiceEndpointsResponseResultJSON contains the JSON metadata for the
// struct [CloudV1GetServiceEndpointsResponseResult]
type cloudV1GetServiceEndpointsResponseResultJSON struct {
	CreatedOn   apijson.Field
	RegionID    apijson.Field
	Service     apijson.Field
	UpdatedAt   apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1GetServiceEndpointsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetServiceEndpointsResponseResultJSON) RawJSON() string {
	return r.raw
}

type CloudV1GetUsageReportResponse struct {
	// Total count of the resources
	Count     int64                                   `json:"count,required"`
	Resources []CloudV1GetUsageReportResponseResource `json:"resources,required"`
	Totals    []CloudV1GetUsageReportResponseTotal    `json:"totals,required"`
	JSON      cloudV1GetUsageReportResponseJSON       `json:"-"`
}

// cloudV1GetUsageReportResponseJSON contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponse]
type cloudV1GetUsageReportResponseJSON struct {
	Count       apijson.Field
	Resources   apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseJSON) RawJSON() string {
	return r.raw
}

// These properties are common for all individual AI clusters in all cost and usage
// reports results (but not in totals)
type CloudV1GetUsageReportResponseResource struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// This field can have the runtime type of
	// [CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceBackupSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceContainerSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceImagesSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceInferenceSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerBillingValueUnion].
	BillingValue interface{} `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit string `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                      `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesType `json:"type,required"`
	// ID of the VM the IP is attached to
	AttachedToVm string `json:"attached_to_vm,nullable" format:"uuid"`
	// Type of the file share
	FileShareType string `json:"file_share_type"`
	// Flavor of the AI cluster
	Flavor string `json:"flavor"`
	// Name of the instance
	InstanceName string `json:"instance_name,nullable"`
	// Type of the instance
	InstanceType EgressTrafficInstanceType `json:"instance_type"`
	// IP address
	IPAddress string `json:"ip_address" format:"ipvanyaddress"`
	// Name of the AI cluster
	LastName string `json:"last_name"`
	// Size of the backup in bytes
	LastSize int64 `json:"last_size"`
	// ID of the network the IP is attached to
	NetworkID string `json:"network_id" format:"uuid"`
	// ID of the port the traffic is associated with
	PortID string `json:"port_id" format:"uuid"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// ID of the backup schedule
	ScheduleID string `json:"schedule_id" format:"uuid"`
	// Unit of size
	SizeUnit string `json:"size_unit"`
	// UUID of the source volume
	SourceVolumeUuid string `json:"source_volume_uuid" format:"uuid"`
	// ID of the subnet the IP is attached to
	SubnetID string `json:"subnet_id" format:"uuid"`
	// This field can have the runtime type of [[]interface{}].
	Tags interface{} `json:"tags"`
	// UUID of the AI cluster
	Uuid string `json:"uuid,nullable" format:"uuid"`
	// ID of the bare metal server the traffic is associated with
	VmID string `json:"vm_id" format:"uuid"`
	// Type of the volume
	VolumeType string                                    `json:"volume_type"`
	JSON       cloudV1GetUsageReportResponseResourceJSON `json:"-"`
	union      CloudV1GetUsageReportResponseResourcesUnion
}

// cloudV1GetUsageReportResponseResourceJSON contains the JSON metadata for the
// struct [CloudV1GetUsageReportResponseResource]
type cloudV1GetUsageReportResponseResourceJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	AttachedToVm      apijson.Field
	FileShareType     apijson.Field
	Flavor            apijson.Field
	InstanceName      apijson.Field
	InstanceType      apijson.Field
	IPAddress         apijson.Field
	LastName          apijson.Field
	LastSize          apijson.Field
	NetworkID         apijson.Field
	PortID            apijson.Field
	RegionName        apijson.Field
	ScheduleID        apijson.Field
	SizeUnit          apijson.Field
	SourceVolumeUuid  apijson.Field
	SubnetID          apijson.Field
	Tags              apijson.Field
	Uuid              apijson.Field
	VmID              apijson.Field
	VolumeType        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r cloudV1GetUsageReportResponseResourceJSON) RawJSON() string {
	return r.raw
}

func (r *CloudV1GetUsageReportResponseResource) UnmarshalJSON(data []byte) (err error) {
	*r = CloudV1GetUsageReportResponseResource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CloudV1GetUsageReportResponseResourcesUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceBackupSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceContainerSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceFileShareSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceImagesSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceInferenceSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceInstanceSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceVolumeSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializer],
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializer].
func (r CloudV1GetUsageReportResponseResource) AsUnion() CloudV1GetUsageReportResponseResourcesUnion {
	return r.union
}

// These properties are common for all individual AI clusters in all cost and usage
// reports results (but not in totals)
//
// Union satisfied by
// [CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceBackupSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceContainerSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceFileShareSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceImagesSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceInferenceSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceInstanceSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceVolumeSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializer],
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializer],
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializer] or
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializer].
type CloudV1GetUsageReportResponseResourcesUnion interface {
	implementsCloudV1GetUsageReportResponseResource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializer{}),
			DiscriminatorValue: "ai_cluster",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializer{}),
			DiscriminatorValue: "baremetal",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializer{}),
			DiscriminatorValue: "basic_vm",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceBackupSerializer{}),
			DiscriminatorValue: "backup",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceContainerSerializer{}),
			DiscriminatorValue: "containers",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializer{}),
			DiscriminatorValue: "egress_traffic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializer{}),
			DiscriminatorValue: "external_ip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceFileShareSerializer{}),
			DiscriminatorValue: "file_share",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializer{}),
			DiscriminatorValue: "floatingip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializer{}),
			DiscriminatorValue: "functions",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializer{}),
			DiscriminatorValue: "functions_calls",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializer{}),
			DiscriminatorValue: "functions_traffic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceImagesSerializer{}),
			DiscriminatorValue: "image",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceInferenceSerializer{}),
			DiscriminatorValue: "inference",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceInstanceSerializer{}),
			DiscriminatorValue: "instance",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializer{}),
			DiscriminatorValue: "load_balancer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializer{}),
			DiscriminatorValue: "log_index",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializer{}),
			DiscriminatorValue: "snapshot",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceVolumeSerializer{}),
			DiscriminatorValue: "volume",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializer{}),
			DiscriminatorValue: "dbaas_postgresql_connection_pooler",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializer{}),
			DiscriminatorValue: "dbaas_postgresql_memory",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializer{}),
			DiscriminatorValue: "dbaas_postgresql_public_network",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializer{}),
			DiscriminatorValue: "dbaas_postgresql_cpu",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializer{}),
			DiscriminatorValue: "dbaas_postgresql_volume",
		},
	)
}

// These properties are common for all individual AI clusters in all cost and usage
// reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Flavor of the AI cluster
	Flavor string `json:"flavor,required"`
	// Name of the AI cluster
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                 `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerType `json:"type,required"`
	// UUID of the AI cluster
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                         `json:"tags,nullable"`
	JSON cloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializer]
type cloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	Flavor            apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	Tags              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerTypeAICluster CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerType = "ai_cluster"
)

func (r CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceAIClusterSerializerTypeAICluster:
		return true
	}
	return false
}

// These properties are common for all individual bare metal servers in all cost
// and usage reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Flavor of the bare metal server
	Flavor string `json:"flavor,required"`
	// Name of the bare metal server
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                 `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerType `json:"type,required"`
	// UUID of the bare metal server
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                         `json:"tags,nullable"`
	JSON cloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializer]
type cloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	Flavor            apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	Tags              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerTypeBaremetal CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerType = "baremetal"
)

func (r CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceBaremetalSerializerTypeBaremetal:
		return true
	}
	return false
}

// These properties are common for all individual basic VMs in all cost and usage
// reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Flavor of the basic VM
	Flavor string `json:"flavor,required"`
	// Name of the basic VM
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                               `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerType `json:"type,required"`
	// UUID of the basic VM
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                       `json:"tags,nullable"`
	JSON cloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerJSON contains the
// JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializer]
type cloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	Flavor            apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	Tags              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerTypeBasicVm CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerType = "basic_vm"
)

func (r CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceBasicVmSerializerTypeBasicVm:
		return true
	}
	return false
}

// These properties are common for all individual backups in all cost and usage
// reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceBackupSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceBackupSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceBackupSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the backup
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// Size of the backup in bytes
	LastSize int64 `json:"last_size,required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// ID of the backup schedule
	ScheduleID string `json:"schedule_id,required" format:"uuid"`
	// UUID of the source volume
	SourceVolumeUuid string                                                             `json:"source_volume_uuid,required" format:"uuid"`
	Type             CloudV1GetUsageReportResponseResourcesResourceBackupSerializerType `json:"type,required"`
	// UUID of the backup
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string                                                             `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseResourcesResourceBackupSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceBackupSerializerJSON contains the
// JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceBackupSerializer]
type cloudV1GetUsageReportResponseResourcesResourceBackupSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	LastSize          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	ScheduleID        apijson.Field
	SourceVolumeUuid  apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceBackupSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceBackupSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceBackupSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceBackupSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceBackupSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceBackupSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceBackupSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceBackupSerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseResourcesResourceBackupSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceBackupSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceBackupSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceBackupSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceBackupSerializerTypeBackup CloudV1GetUsageReportResponseResourcesResourceBackupSerializerType = "backup"
)

func (r CloudV1GetUsageReportResponseResourcesResourceBackupSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceBackupSerializerTypeBackup:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceContainerSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceContainerSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceContainerSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the container
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                 `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceContainerSerializerType `json:"type,required"`
	// UUID of the container
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string                                                                `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseResourcesResourceContainerSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceContainerSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceContainerSerializer]
type cloudV1GetUsageReportResponseResourcesResourceContainerSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceContainerSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceContainerSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceContainerSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceContainerSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceContainerSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceContainerSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceContainerSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceContainerSerializerBillingValueUnitGBs CloudV1GetUsageReportResponseResourcesResourceContainerSerializerBillingValueUnit = "GBS"
)

func (r CloudV1GetUsageReportResponseResourcesResourceContainerSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceContainerSerializerBillingValueUnitGBs:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceContainerSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceContainerSerializerTypeContainers CloudV1GetUsageReportResponseResourcesResourceContainerSerializerType = "containers"
)

func (r CloudV1GetUsageReportResponseResourcesResourceContainerSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceContainerSerializerTypeContainers:
		return true
	}
	return false
}

// These properties are common for all individual egress traffic in all cost and
// usage reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Type of the instance
	InstanceType EgressTrafficInstanceType `json:"instance_type,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the port the traffic is associated with
	PortID string `json:"port_id,required" format:"uuid"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Unit of size
	SizeUnit string                                                                    `json:"size_unit,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerType `json:"type,required"`
	// ID of the bare metal server the traffic is associated with
	VmID string `json:"vm_id,required" format:"uuid"`
	// Name of the instance
	InstanceName string `json:"instance_name,nullable"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                             `json:"tags,nullable"`
	JSON cloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializer]
type cloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	InstanceType      apijson.Field
	LastSeen          apijson.Field
	PortID            apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	SizeUnit          apijson.Field
	Type              apijson.Field
	VmID              apijson.Field
	InstanceName      apijson.Field
	RegionName        apijson.Field
	Tags              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerBillingValueUnitBytes CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerBillingValueUnit = "bytes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerBillingValueUnitBytes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerTypeEgressTraffic CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerType = "egress_traffic"
)

func (r CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceEgressTrafficSerializerTypeEgressTraffic:
		return true
	}
	return false
}

// These properties are common for all individual external IPs in all cost and
// usage reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializer struct {
	// ID of the VM the IP is attached to
	AttachedToVm string `json:"attached_to_vm,required,nullable" format:"uuid"`
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// IP address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the network the IP is attached to
	NetworkID string `json:"network_id,required" format:"uuid"`
	// ID of the port the IP is associated with
	PortID string `json:"port_id,required" format:"uuid"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// ID of the subnet the IP is attached to
	SubnetID string                                                                 `json:"subnet_id,required" format:"uuid"`
	Type     CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                 `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializer]
type cloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerJSON struct {
	AttachedToVm      apijson.Field
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	IPAddress         apijson.Field
	LastSeen          apijson.Field
	NetworkID         apijson.Field
	PortID            apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	SubnetID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerTypeExternalIP CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerType = "external_ip"
)

func (r CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceExternalIPSerializerTypeExternalIP:
		return true
	}
	return false
}

// These properties are common for all individual file shares in all cost and usage
// reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceFileShareSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Type of the file share
	FileShareType string `json:"file_share_type,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the file share
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// Size of the file share in bytes
	LastSize int64 `json:"last_size,required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Unit of size
	SizeUnit CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerSizeUnit `json:"size_unit,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerType     `json:"type,required"`
	// UUID of the file share
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                         `json:"tags,nullable"`
	JSON cloudV1GetUsageReportResponseResourcesResourceFileShareSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceFileShareSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceFileShareSerializer]
type cloudV1GetUsageReportResponseResourcesResourceFileShareSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FileShareType     apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	LastSize          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	SizeUnit          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	Tags              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceFileShareSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceFileShareSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceFileShareSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceFileShareSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

// Unit of size
type CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerSizeUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerSizeUnitGiB CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerSizeUnit = "GiB"
)

func (r CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerSizeUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerSizeUnitGiB:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerTypeFileShare CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerType = "file_share"
)

func (r CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceFileShareSerializerTypeFileShare:
		return true
	}
	return false
}

// These properties are common for all individual floating IPs in all cost and
// usage reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// IP address
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	// Name of the floating IP
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                  `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerType `json:"type,required"`
	// UUID of the floating IP
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                          `json:"tags,nullable"`
	JSON cloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializer]
type cloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	IPAddress         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	Tags              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerTypeFloatingip CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerType = "floatingip"
)

func (r CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceFloatingIPSerializerTypeFloatingip:
		return true
	}
	return false
}

// These properties are common for all individual functions in all cost and usage
// reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the function
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                 `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerType `json:"type,required"`
	// UUID of the function
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string                                                                `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializer]
type cloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerBillingValueUnitGBs CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerBillingValueUnit = "GBS"
)

func (r CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerBillingValueUnitGBs:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerTypeFunctions CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerType = "functions"
)

func (r CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceFunctionsSerializerTypeFunctions:
		return true
	}
	return false
}

// These properties are common for all individual function calls in all cost and
// usage reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the function call
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                     `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerType `json:"type,required"`
	// UUID of the function call
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string                                                                    `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializer]
type cloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerBillingValueUnitMls CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerBillingValueUnit = "MLS"
)

func (r CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerBillingValueUnitMls:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerTypeFunctionsCalls CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerType = "functions_calls"
)

func (r CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceFunctionCallsSerializerTypeFunctionsCalls:
		return true
	}
	return false
}

// These properties are common for all individual function egress traffic in all
// cost and usage reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the function egress traffic
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                             `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerType `json:"type,required"`
	// UUID of the function egress traffic
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string                                                                            `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializer]
type cloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerBillingValueUnitGB CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerBillingValueUnit = "GB"
)

func (r CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerBillingValueUnitGB:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerTypeFunctionsTraffic CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerType = "functions_traffic"
)

func (r CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceFunctionEgressTrafficSerializerTypeFunctionsTraffic:
		return true
	}
	return false
}

// These properties are common for all individual images in all cost and usage
// reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceImagesSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceImagesSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceImagesSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the image
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// Size of the image in bytes
	LastSize int64 `json:"last_size,required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Unit of size
	SizeUnit CloudV1GetUsageReportResponseResourcesResourceImagesSerializerSizeUnit `json:"size_unit,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceImagesSerializerType     `json:"type,required"`
	// UUID of the image
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                      `json:"tags,nullable"`
	JSON cloudV1GetUsageReportResponseResourcesResourceImagesSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceImagesSerializerJSON contains the
// JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceImagesSerializer]
type cloudV1GetUsageReportResponseResourcesResourceImagesSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	LastSize          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	SizeUnit          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	Tags              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceImagesSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceImagesSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceImagesSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceImagesSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceImagesSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceImagesSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceImagesSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceImagesSerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseResourcesResourceImagesSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceImagesSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceImagesSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

// Unit of size
type CloudV1GetUsageReportResponseResourcesResourceImagesSerializerSizeUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceImagesSerializerSizeUnitBytes CloudV1GetUsageReportResponseResourcesResourceImagesSerializerSizeUnit = "bytes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceImagesSerializerSizeUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceImagesSerializerSizeUnitBytes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceImagesSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceImagesSerializerTypeImage CloudV1GetUsageReportResponseResourcesResourceImagesSerializerType = "image"
)

func (r CloudV1GetUsageReportResponseResourcesResourceImagesSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceImagesSerializerTypeImage:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceInferenceSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceInferenceSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit string `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Flavor of the inference deployment
	Flavor string `json:"flavor,required"`
	// Name of the inference deployment
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                 `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceInferenceSerializerType `json:"type,required"`
	// UUID of the inference deployment
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string                                                                `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseResourcesResourceInferenceSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceInferenceSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceInferenceSerializer]
type cloudV1GetUsageReportResponseResourcesResourceInferenceSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	Flavor            apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceInferenceSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceInferenceSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceInferenceSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceInferenceSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceInferenceSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceInferenceSerializerBillingValueUnion)(nil)).Elem(),
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

type CloudV1GetUsageReportResponseResourcesResourceInferenceSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceInferenceSerializerTypeInference CloudV1GetUsageReportResponseResourcesResourceInferenceSerializerType = "inference"
)

func (r CloudV1GetUsageReportResponseResourcesResourceInferenceSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceInferenceSerializerTypeInference:
		return true
	}
	return false
}

// These properties are common for all individual instances in all cost and usage
// reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceInstanceSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Flavor of the instance
	Flavor string `json:"flavor,required"`
	// Name of the instance
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerType `json:"type,required"`
	// UUID of the instance
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                        `json:"tags,nullable"`
	JSON cloudV1GetUsageReportResponseResourcesResourceInstanceSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceInstanceSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceInstanceSerializer]
type cloudV1GetUsageReportResponseResourcesResourceInstanceSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	Flavor            apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	Tags              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceInstanceSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceInstanceSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceInstanceSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceInstanceSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerTypeInstance CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerType = "instance"
)

func (r CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceInstanceSerializerTypeInstance:
		return true
	}
	return false
}

// These properties are common for all individual load balancers in all cost and
// usage reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Flavor of the load balancer
	Flavor string `json:"flavor,required"`
	// Name of the load balancer
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                    `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerType `json:"type,required"`
	// UUID of the load balancer
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                            `json:"tags,nullable"`
	JSON cloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializer]
type cloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	Flavor            apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	Tags              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerTypeLoadBalancer CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerType = "load_balancer"
)

func (r CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceLoadBalancerSerializerTypeLoadBalancer:
		return true
	}
	return false
}

// These properties are common for all individual log indexes in all cost and usage
// reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the log index
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// Size of the log index in bytes
	LastSize int64 `json:"last_size,required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Unit of size
	SizeUnit string                                                               `json:"size_unit,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerType `json:"type,required"`
	// UUID of the log index
	Uuid string `json:"uuid,required,nullable" format:"uuid"`
	// Region name
	RegionName string                                                               `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializer]
type cloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	LastSize          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	SizeUnit          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerTypeLogIndex CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerType = "log_index"
)

func (r CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceLogIndexSerializerTypeLogIndex:
		return true
	}
	return false
}

// These properties are common for all individual snapshots in all cost and usage
// reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the snapshot
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// Size of the snapshot in bytes
	LastSize int64 `json:"last_size,required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Unit of size
	SizeUnit string `json:"size_unit,required"`
	// UUID of the source volume
	SourceVolumeUuid string                                                               `json:"source_volume_uuid,required" format:"uuid"`
	Type             CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerType `json:"type,required"`
	// UUID of the snapshot
	Uuid string `json:"uuid,required" format:"uuid"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                        `json:"tags,nullable"`
	JSON cloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializer]
type cloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	LastSize          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	SizeUnit          apijson.Field
	SourceVolumeUuid  apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	VolumeType        apijson.Field
	RegionName        apijson.Field
	Tags              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerTypeSnapshot CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerType = "snapshot"
)

func (r CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceSnapshotSerializerTypeSnapshot:
		return true
	}
	return false
}

// These properties are common for all individual volumes in all cost and usage
// reports results (but not in totals)
type CloudV1GetUsageReportResponseResourcesResourceVolumeSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the volume
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// Size of the volume in bytes
	LastSize int64 `json:"last_size,required"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Unit of size
	SizeUnit string                                                             `json:"size_unit,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerType `json:"type,required"`
	// UUID of the volume
	Uuid string `json:"uuid,required" format:"uuid"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
	// ID of the VM the volume is attached to
	AttachedToVm string `json:"attached_to_vm,nullable" format:"uuid"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// List of tags
	Tags []interface{}                                                      `json:"tags,nullable"`
	JSON cloudV1GetUsageReportResponseResourcesResourceVolumeSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceVolumeSerializerJSON contains the
// JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceVolumeSerializer]
type cloudV1GetUsageReportResponseResourcesResourceVolumeSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	LastSize          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	SizeUnit          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	VolumeType        apijson.Field
	AttachedToVm      apijson.Field
	RegionName        apijson.Field
	Tags              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceVolumeSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceVolumeSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceVolumeSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceVolumeSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerTypeVolume CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerType = "volume"
)

func (r CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceVolumeSerializerTypeVolume:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                             `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerType `json:"type,required"`
	// UUID of the cluster
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string                                                                            `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializer]
type cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerTypeDbaasPostgresqlConnectionPooler CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerType = "dbaas_postgresql_connection_pooler"
)

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPoolerSerializerTypeDbaasPostgresqlConnectionPooler:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                             `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerType `json:"type,required"`
	// UUID of the cluster
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string                                                                            `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializer]
type cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerTypeDbaasPostgresqlMemory CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerType = "dbaas_postgresql_memory"
)

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlMemorySerializerTypeDbaasPostgresqlMemory:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                                    `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerType `json:"type,required"`
	// UUID of the cluster
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string                                                                                   `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializer]
type cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerTypeDbaasPostgresqlPublicNetwork CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerType = "dbaas_postgresql_public_network"
)

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlPublicNetworkSerializerTypeDbaasPostgresqlPublicNetwork:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                          `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerType `json:"type,required"`
	// UUID of the cluster
	Uuid string `json:"uuid,required" format:"uuid"`
	// Region name
	RegionName string                                                                         `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializer]
type cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerTypeDbaasPostgresqlCPU CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerType = "dbaas_postgresql_cpu"
)

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlcpuSerializerTypeDbaasPostgresqlCPU:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// First time the resource was seen in the given period
	FirstSeen time.Time `json:"first_seen,required" format:"date-time"`
	// Name of the cluster
	LastName string `json:"last_name,required"`
	// Last time the resource was seen in the given period
	LastSeen time.Time `json:"last_seen,required" format:"date-time"`
	// ID of the project the resource belongs to
	ProjectID int64 `json:"project_id,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Unit of size
	SizeUnit string                                                                            `json:"size_unit,required"`
	Type     CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerType `json:"type,required"`
	// UUID of the cluster
	Uuid string `json:"uuid,required" format:"uuid"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
	// Region name
	RegionName string                                                                            `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializer]
type cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FirstSeen         apijson.Field
	LastName          apijson.Field
	LastSeen          apijson.Field
	ProjectID         apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	SizeUnit          apijson.Field
	Type              apijson.Field
	Uuid              apijson.Field
	VolumeType        apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializer) implementsCloudV1GetUsageReportResponseResource() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerType string

const (
	CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerTypeDbaasPostgresqlVolume CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerType = "dbaas_postgresql_volume"
)

func (r CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesResourceDbaasPostgreSqlVolumeSerializerTypeDbaasPostgresqlVolume:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseResourcesType string

const (
	CloudV1GetUsageReportResponseResourcesTypeAICluster                       CloudV1GetUsageReportResponseResourcesType = "ai_cluster"
	CloudV1GetUsageReportResponseResourcesTypeBaremetal                       CloudV1GetUsageReportResponseResourcesType = "baremetal"
	CloudV1GetUsageReportResponseResourcesTypeBasicVm                         CloudV1GetUsageReportResponseResourcesType = "basic_vm"
	CloudV1GetUsageReportResponseResourcesTypeBackup                          CloudV1GetUsageReportResponseResourcesType = "backup"
	CloudV1GetUsageReportResponseResourcesTypeContainers                      CloudV1GetUsageReportResponseResourcesType = "containers"
	CloudV1GetUsageReportResponseResourcesTypeEgressTraffic                   CloudV1GetUsageReportResponseResourcesType = "egress_traffic"
	CloudV1GetUsageReportResponseResourcesTypeExternalIP                      CloudV1GetUsageReportResponseResourcesType = "external_ip"
	CloudV1GetUsageReportResponseResourcesTypeFileShare                       CloudV1GetUsageReportResponseResourcesType = "file_share"
	CloudV1GetUsageReportResponseResourcesTypeFloatingip                      CloudV1GetUsageReportResponseResourcesType = "floatingip"
	CloudV1GetUsageReportResponseResourcesTypeFunctions                       CloudV1GetUsageReportResponseResourcesType = "functions"
	CloudV1GetUsageReportResponseResourcesTypeFunctionsCalls                  CloudV1GetUsageReportResponseResourcesType = "functions_calls"
	CloudV1GetUsageReportResponseResourcesTypeFunctionsTraffic                CloudV1GetUsageReportResponseResourcesType = "functions_traffic"
	CloudV1GetUsageReportResponseResourcesTypeImage                           CloudV1GetUsageReportResponseResourcesType = "image"
	CloudV1GetUsageReportResponseResourcesTypeInference                       CloudV1GetUsageReportResponseResourcesType = "inference"
	CloudV1GetUsageReportResponseResourcesTypeInstance                        CloudV1GetUsageReportResponseResourcesType = "instance"
	CloudV1GetUsageReportResponseResourcesTypeLoadBalancer                    CloudV1GetUsageReportResponseResourcesType = "load_balancer"
	CloudV1GetUsageReportResponseResourcesTypeLogIndex                        CloudV1GetUsageReportResponseResourcesType = "log_index"
	CloudV1GetUsageReportResponseResourcesTypeSnapshot                        CloudV1GetUsageReportResponseResourcesType = "snapshot"
	CloudV1GetUsageReportResponseResourcesTypeVolume                          CloudV1GetUsageReportResponseResourcesType = "volume"
	CloudV1GetUsageReportResponseResourcesTypeDbaasPostgresqlConnectionPooler CloudV1GetUsageReportResponseResourcesType = "dbaas_postgresql_connection_pooler"
	CloudV1GetUsageReportResponseResourcesTypeDbaasPostgresqlMemory           CloudV1GetUsageReportResponseResourcesType = "dbaas_postgresql_memory"
	CloudV1GetUsageReportResponseResourcesTypeDbaasPostgresqlPublicNetwork    CloudV1GetUsageReportResponseResourcesType = "dbaas_postgresql_public_network"
	CloudV1GetUsageReportResponseResourcesTypeDbaasPostgresqlCPU              CloudV1GetUsageReportResponseResourcesType = "dbaas_postgresql_cpu"
	CloudV1GetUsageReportResponseResourcesTypeDbaasPostgresqlVolume           CloudV1GetUsageReportResponseResourcesType = "dbaas_postgresql_volume"
)

func (r CloudV1GetUsageReportResponseResourcesType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseResourcesTypeAICluster, CloudV1GetUsageReportResponseResourcesTypeBaremetal, CloudV1GetUsageReportResponseResourcesTypeBasicVm, CloudV1GetUsageReportResponseResourcesTypeBackup, CloudV1GetUsageReportResponseResourcesTypeContainers, CloudV1GetUsageReportResponseResourcesTypeEgressTraffic, CloudV1GetUsageReportResponseResourcesTypeExternalIP, CloudV1GetUsageReportResponseResourcesTypeFileShare, CloudV1GetUsageReportResponseResourcesTypeFloatingip, CloudV1GetUsageReportResponseResourcesTypeFunctions, CloudV1GetUsageReportResponseResourcesTypeFunctionsCalls, CloudV1GetUsageReportResponseResourcesTypeFunctionsTraffic, CloudV1GetUsageReportResponseResourcesTypeImage, CloudV1GetUsageReportResponseResourcesTypeInference, CloudV1GetUsageReportResponseResourcesTypeInstance, CloudV1GetUsageReportResponseResourcesTypeLoadBalancer, CloudV1GetUsageReportResponseResourcesTypeLogIndex, CloudV1GetUsageReportResponseResourcesTypeSnapshot, CloudV1GetUsageReportResponseResourcesTypeVolume, CloudV1GetUsageReportResponseResourcesTypeDbaasPostgresqlConnectionPooler, CloudV1GetUsageReportResponseResourcesTypeDbaasPostgresqlMemory, CloudV1GetUsageReportResponseResourcesTypeDbaasPostgresqlPublicNetwork, CloudV1GetUsageReportResponseResourcesTypeDbaasPostgresqlCPU, CloudV1GetUsageReportResponseResourcesTypeDbaasPostgresqlVolume:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotal struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// This field can have the runtime type of
	// [CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerBillingValueUnion],
	// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerBillingValueUnion].
	BillingValue interface{} `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit string `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                   `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsType `json:"type,required"`
	// Type of the file share
	FileShareType string `json:"file_share_type"`
	// Flavor of the AI cluster
	Flavor string `json:"flavor"`
	// Type of the instance
	InstanceType EgressTrafficInstanceType `json:"instance_type"`
	// Region name
	RegionName string `json:"region_name,nullable"`
	// Type of the volume
	VolumeType string                                 `json:"volume_type"`
	JSON       cloudV1GetUsageReportResponseTotalJSON `json:"-"`
	union      CloudV1GetUsageReportResponseTotalsUnion
}

// cloudV1GetUsageReportResponseTotalJSON contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotal]
type cloudV1GetUsageReportResponseTotalJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	FileShareType     apijson.Field
	Flavor            apijson.Field
	InstanceType      apijson.Field
	RegionName        apijson.Field
	VolumeType        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r cloudV1GetUsageReportResponseTotalJSON) RawJSON() string {
	return r.raw
}

func (r *CloudV1GetUsageReportResponseTotal) UnmarshalJSON(data []byte) (err error) {
	*r = CloudV1GetUsageReportResponseTotal{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CloudV1GetUsageReportResponseTotalsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializer].
func (r CloudV1GetUsageReportResponseTotal) AsUnion() CloudV1GetUsageReportResponseTotalsUnion {
	return r.union
}

// Union satisfied by
// [CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializer],
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializer]
// or
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializer].
type CloudV1GetUsageReportResponseTotalsUnion interface {
	implementsCloudV1GetUsageReportResponseTotal()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializer{}),
			DiscriminatorValue: "ai_cluster",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializer{}),
			DiscriminatorValue: "baremetal",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializer{}),
			DiscriminatorValue: "basic_vm",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializer{}),
			DiscriminatorValue: "containers",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializer{}),
			DiscriminatorValue: "egress_traffic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializer{}),
			DiscriminatorValue: "external_ip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializer{}),
			DiscriminatorValue: "file_share",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializer{}),
			DiscriminatorValue: "floatingip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializer{}),
			DiscriminatorValue: "functions",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializer{}),
			DiscriminatorValue: "functions_calls",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializer{}),
			DiscriminatorValue: "functions_traffic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializer{}),
			DiscriminatorValue: "image",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializer{}),
			DiscriminatorValue: "inference",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializer{}),
			DiscriminatorValue: "instance",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializer{}),
			DiscriminatorValue: "load_balancer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializer{}),
			DiscriminatorValue: "log_index",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializer{}),
			DiscriminatorValue: "snapshot",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializer{}),
			DiscriminatorValue: "volume",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializer{}),
			DiscriminatorValue: "dbaas_postgresql_connection_pooler",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializer{}),
			DiscriminatorValue: "dbaas_postgresql_memory",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializer{}),
			DiscriminatorValue: "dbaas_postgresql_public_network",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializer{}),
			DiscriminatorValue: "dbaas_postgresql_cpu",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializer{}),
			DiscriminatorValue: "dbaas_postgresql_volume",
		},
	)
}

type CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Flavor of the AI cluster
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                     `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                    `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Flavor            apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerTypeAICluster CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerType = "ai_cluster"
)

func (r CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalAIClusterReportItemSerializerTypeAICluster:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Flavor of the bare metal server
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                     `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                    `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Flavor            apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerTypeBaremetal CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerType = "baremetal"
)

func (r CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalBaremetalReportItemSerializerTypeBaremetal:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Flavor of the basic VM
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                   `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                  `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Flavor            apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerTypeBasicVm CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerType = "basic_vm"
)

func (r CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalBasicVmReportItemSerializerTypeBasicVm:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                     `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                    `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerBillingValueUnitGBs CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerBillingValueUnit = "GBS"
)

func (r CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerBillingValueUnitGBs:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerTypeContainers CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerType = "containers"
)

func (r CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalContainerReportItemSerializerTypeContainers:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Type of the instance
	InstanceType EgressTrafficInstanceType `json:"instance_type,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                         `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                        `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	InstanceType      apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerBillingValueUnitBytes CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerBillingValueUnit = "bytes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerBillingValueUnitBytes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerTypeEgressTraffic CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerType = "egress_traffic"
)

func (r CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalEgressTrafficReportItemSerializerTypeEgressTraffic:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                      `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                     `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerTypeExternalIP CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerType = "external_ip"
)

func (r CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalExternalIPReportItemSerializerTypeExternalIP:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Type of the file share
	FileShareType string `json:"file_share_type,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                     `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                    `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	FileShareType     apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerTypeFileShare CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerType = "file_share"
)

func (r CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalFileShareReportItemSerializerTypeFileShare:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                      `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                     `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerTypeFloatingip CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerType = "floatingip"
)

func (r CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalFloatingIPReportItemSerializerTypeFloatingip:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                     `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                    `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerBillingValueUnitGBs CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerBillingValueUnit = "GBS"
)

func (r CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerBillingValueUnitGBs:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerTypeFunctions CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerType = "functions"
)

func (r CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalFunctionsReportItemSerializerTypeFunctions:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                         `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                        `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerBillingValueUnitMls CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerBillingValueUnit = "MLS"
)

func (r CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerBillingValueUnitMls:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerTypeFunctionsCalls CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerType = "functions_calls"
)

func (r CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalFunctionCallsReportItemSerializerTypeFunctionsCalls:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                                 `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                                `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerBillingValueUnitGB CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerBillingValueUnit = "GB"
)

func (r CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerBillingValueUnitGB:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerTypeFunctionsTraffic CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerType = "functions_traffic"
)

func (r CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalFunctionEgressTrafficReportItemSerializerTypeFunctionsTraffic:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                  `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                 `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerTypeImage CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerType = "image"
)

func (r CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalImagesReportItemSerializerTypeImage:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit string `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                     `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                    `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

type CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerTypeInference CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerType = "inference"
)

func (r CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalInferenceReportItemSerializerTypeInference:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Flavor of the instance
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                    `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                   `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Flavor            apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerTypeInstance CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerType = "instance"
)

func (r CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalInstanceReportItemSerializerTypeInstance:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Flavor of the load balancer
	Flavor string `json:"flavor,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                        `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                       `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Flavor            apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerTypeLoadBalancer CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerType = "load_balancer"
)

func (r CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalLoadBalancerReportItemSerializerTypeLoadBalancer:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                    `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                   `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerTypeLogIndex CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerType = "log_index"
)

func (r CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalLogIndexReportItemSerializerTypeLogIndex:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                    `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerType `json:"type,required"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
	// Region name
	RegionName string                                                                   `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	VolumeType        apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerTypeSnapshot CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerType = "snapshot"
)

func (r CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalSnapshotReportItemSerializerTypeSnapshot:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                  `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerType `json:"type,required"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
	// Region name
	RegionName string                                                                 `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerJSON contains
// the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	VolumeType        apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerTypeVolume CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerType = "volume"
)

func (r CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalVolumeReportItemSerializerTypeVolume:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                                 `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                                `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerTypeDbaasPostgresqlConnectionPooler CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerType = "dbaas_postgresql_connection_pooler"
)

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPoolerReportItemSerializerTypeDbaasPostgresqlConnectionPooler:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                                 `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                                `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerTypeDbaasPostgresqlMemory CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerType = "dbaas_postgresql_memory"
)

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlMemoryReportItemSerializerTypeDbaasPostgresqlMemory:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                                        `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                                       `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerTypeDbaasPostgresqlPublicNetwork CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerType = "dbaas_postgresql_public_network"
)

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlPublicNetworkReportItemSerializerTypeDbaasPostgresqlPublicNetwork:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                              `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerType `json:"type,required"`
	// Region name
	RegionName string                                                                             `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerBillingValueUnitMinutes CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerBillingValueUnit = "minutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerBillingValueUnitMinutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerTypeDbaasPostgresqlCPU CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerType = "dbaas_postgresql_cpu"
)

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlcpuReportItemSerializerTypeDbaasPostgresqlCPU:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializer struct {
	// Name of the billing metric
	BillingMetricName string `json:"billing_metric_name,required"`
	// Value of the billing metric
	BillingValue CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerBillingValueUnion `json:"billing_value,required"`
	// Unit of billing value
	BillingValueUnit CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerBillingValueUnit `json:"billing_value_unit,required"`
	// Region ID
	Region int64 `json:"region,required"`
	// Region ID
	RegionID int64                                                                                 `json:"region_id,required"`
	Type     CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerType `json:"type,required"`
	// Type of the volume
	VolumeType string `json:"volume_type,required"`
	// Region name
	RegionName string                                                                                `json:"region_name,nullable"`
	JSON       cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerJSON `json:"-"`
}

// cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerJSON
// contains the JSON metadata for the struct
// [CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializer]
type cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerJSON struct {
	BillingMetricName apijson.Field
	BillingValue      apijson.Field
	BillingValueUnit  apijson.Field
	Region            apijson.Field
	RegionID          apijson.Field
	Type              apijson.Field
	VolumeType        apijson.Field
	RegionName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerJSON) RawJSON() string {
	return r.raw
}

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializer) implementsCloudV1GetUsageReportResponseTotal() {
}

// Value of the billing metric
//
// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerBillingValueUnion interface {
	ImplementsCloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerBillingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerBillingValueUnion)(nil)).Elem(),
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

// Unit of billing value
type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerBillingValueUnit string

const (
	CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerBillingValueUnitGbminutes CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerBillingValueUnit = "gbminutes"
)

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerBillingValueUnit) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerBillingValueUnitGbminutes:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerType string

const (
	CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerTypeDbaasPostgresqlVolume CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerType = "dbaas_postgresql_volume"
)

func (r CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTotalDbaasPostgreSqlVolumeReportItemSerializerTypeDbaasPostgresqlVolume:
		return true
	}
	return false
}

type CloudV1GetUsageReportResponseTotalsType string

const (
	CloudV1GetUsageReportResponseTotalsTypeAICluster                       CloudV1GetUsageReportResponseTotalsType = "ai_cluster"
	CloudV1GetUsageReportResponseTotalsTypeBaremetal                       CloudV1GetUsageReportResponseTotalsType = "baremetal"
	CloudV1GetUsageReportResponseTotalsTypeBasicVm                         CloudV1GetUsageReportResponseTotalsType = "basic_vm"
	CloudV1GetUsageReportResponseTotalsTypeContainers                      CloudV1GetUsageReportResponseTotalsType = "containers"
	CloudV1GetUsageReportResponseTotalsTypeEgressTraffic                   CloudV1GetUsageReportResponseTotalsType = "egress_traffic"
	CloudV1GetUsageReportResponseTotalsTypeExternalIP                      CloudV1GetUsageReportResponseTotalsType = "external_ip"
	CloudV1GetUsageReportResponseTotalsTypeFileShare                       CloudV1GetUsageReportResponseTotalsType = "file_share"
	CloudV1GetUsageReportResponseTotalsTypeFloatingip                      CloudV1GetUsageReportResponseTotalsType = "floatingip"
	CloudV1GetUsageReportResponseTotalsTypeFunctions                       CloudV1GetUsageReportResponseTotalsType = "functions"
	CloudV1GetUsageReportResponseTotalsTypeFunctionsCalls                  CloudV1GetUsageReportResponseTotalsType = "functions_calls"
	CloudV1GetUsageReportResponseTotalsTypeFunctionsTraffic                CloudV1GetUsageReportResponseTotalsType = "functions_traffic"
	CloudV1GetUsageReportResponseTotalsTypeImage                           CloudV1GetUsageReportResponseTotalsType = "image"
	CloudV1GetUsageReportResponseTotalsTypeInference                       CloudV1GetUsageReportResponseTotalsType = "inference"
	CloudV1GetUsageReportResponseTotalsTypeInstance                        CloudV1GetUsageReportResponseTotalsType = "instance"
	CloudV1GetUsageReportResponseTotalsTypeLoadBalancer                    CloudV1GetUsageReportResponseTotalsType = "load_balancer"
	CloudV1GetUsageReportResponseTotalsTypeLogIndex                        CloudV1GetUsageReportResponseTotalsType = "log_index"
	CloudV1GetUsageReportResponseTotalsTypeSnapshot                        CloudV1GetUsageReportResponseTotalsType = "snapshot"
	CloudV1GetUsageReportResponseTotalsTypeVolume                          CloudV1GetUsageReportResponseTotalsType = "volume"
	CloudV1GetUsageReportResponseTotalsTypeDbaasPostgresqlConnectionPooler CloudV1GetUsageReportResponseTotalsType = "dbaas_postgresql_connection_pooler"
	CloudV1GetUsageReportResponseTotalsTypeDbaasPostgresqlMemory           CloudV1GetUsageReportResponseTotalsType = "dbaas_postgresql_memory"
	CloudV1GetUsageReportResponseTotalsTypeDbaasPostgresqlPublicNetwork    CloudV1GetUsageReportResponseTotalsType = "dbaas_postgresql_public_network"
	CloudV1GetUsageReportResponseTotalsTypeDbaasPostgresqlCPU              CloudV1GetUsageReportResponseTotalsType = "dbaas_postgresql_cpu"
	CloudV1GetUsageReportResponseTotalsTypeDbaasPostgresqlVolume           CloudV1GetUsageReportResponseTotalsType = "dbaas_postgresql_volume"
)

func (r CloudV1GetUsageReportResponseTotalsType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportResponseTotalsTypeAICluster, CloudV1GetUsageReportResponseTotalsTypeBaremetal, CloudV1GetUsageReportResponseTotalsTypeBasicVm, CloudV1GetUsageReportResponseTotalsTypeContainers, CloudV1GetUsageReportResponseTotalsTypeEgressTraffic, CloudV1GetUsageReportResponseTotalsTypeExternalIP, CloudV1GetUsageReportResponseTotalsTypeFileShare, CloudV1GetUsageReportResponseTotalsTypeFloatingip, CloudV1GetUsageReportResponseTotalsTypeFunctions, CloudV1GetUsageReportResponseTotalsTypeFunctionsCalls, CloudV1GetUsageReportResponseTotalsTypeFunctionsTraffic, CloudV1GetUsageReportResponseTotalsTypeImage, CloudV1GetUsageReportResponseTotalsTypeInference, CloudV1GetUsageReportResponseTotalsTypeInstance, CloudV1GetUsageReportResponseTotalsTypeLoadBalancer, CloudV1GetUsageReportResponseTotalsTypeLogIndex, CloudV1GetUsageReportResponseTotalsTypeSnapshot, CloudV1GetUsageReportResponseTotalsTypeVolume, CloudV1GetUsageReportResponseTotalsTypeDbaasPostgresqlConnectionPooler, CloudV1GetUsageReportResponseTotalsTypeDbaasPostgresqlMemory, CloudV1GetUsageReportResponseTotalsTypeDbaasPostgresqlPublicNetwork, CloudV1GetUsageReportResponseTotalsTypeDbaasPostgresqlCPU, CloudV1GetUsageReportResponseTotalsTypeDbaasPostgresqlVolume:
		return true
	}
	return false
}

type CloudV1GetAvailableNetworksParams struct {
	// Limit the number of returned limit request entities.
	Limit param.Field[int64] `query:"limit"`
	// Filter by metadata keys. Must be a valid JSON string. curl -G --data-urlencode
	// "metadata_k=["value", "sense"]" --url
	// "http://localhost:1111/v1/availablenetworks/1/1"
	MetadataK param.Field[string] `query:"metadata_k"`
	// Filter by metadata key-value pairs. Must be a valid JSON string. curl -G
	// --data-urlencode "metadata_kv={"key": "value"}" --url
	// "http://localhost:1111/v1/availablenetworks/1/1"
	MetadataKv param.Field[string] `query:"metadata_kv"`
	// Can be used to only show subnets of the specific network
	NetworkID param.Field[string] `query:"network_id"`
	// Filter network by network type (vlan or vxlan)
	NetworkType param.Field[string] `query:"network_type"`
	// Offset value is used to exclude the first set of records from the result.
	Offset param.Field[int64] `query:"offset"`
	// Order networks by transmitted fields and directions (name.asc).
	OrderBy param.Field[string] `query:"order_by"`
	// Can be used to only show networks with shared state"
	Shared param.Field[bool] `query:"shared"`
}

// URLQuery serializes [CloudV1GetAvailableNetworksParams]'s query parameters as
// `url.Values`.
func (r CloudV1GetAvailableNetworksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1GetBmReservationFlavorsParams struct {
	// Client identifier. Must be used for users w/o client_id in jwt
	ClientID param.Field[int64] `query:"client_id_"`
	// Flag for filtering disabled flavors in the region. Defaults to true
	Disabled param.Field[bool] `query:"disabled"`
	// None
	EnsureCached param.Field[string] `query:"ensure_cached"`
	// Set to true to get only flavors with Windows OS, false to get only with Linux
	// OS, None option to without the OS filter. Defaults to none.
	WindowsOs param.Field[bool] `query:"windows_os"`
}

// URLQuery serializes [CloudV1GetBmReservationFlavorsParams]'s query parameters as
// `url.Values`.
func (r CloudV1GetBmReservationFlavorsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1GetFlavorsParams struct {
	// Flag for filtering disabled flavors in the region. Defaults to true
	Disabled param.Field[bool] `query:"disabled"`
	// Set to true to exclude flavors dedicated to linux images. Default False
	ExcludeLinux param.Field[bool] `query:"exclude_linux"`
	// Set to true to exclude flavors dedicated to windows images. Default False
	ExcludeWindows param.Field[bool] `query:"exclude_windows"`
	// Set to true if the response should include flavor prices
	IncludePrices param.Field[bool] `query:"include_prices"`
}

// URLQuery serializes [CloudV1GetFlavorsParams]'s query parameters as
// `url.Values`.
func (r CloudV1GetFlavorsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1GetLbFlavorsParams struct {
	// Set to true if the response should include flavor prices
	IncludePrices param.Field[bool] `query:"include_prices"`
}

// URLQuery serializes [CloudV1GetLbFlavorsParams]'s query parameters as
// `url.Values`.
func (r CloudV1GetLbFlavorsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1GetServiceEndpointsParams struct {
	// Region ID
	RegionID param.Field[int64] `query:"region_id"`
	// Endpoint service
	Service param.Field[EndpointService] `query:"service"`
}

// URLQuery serializes [CloudV1GetServiceEndpointsParams]'s query parameters as
// `url.Values`.
func (r CloudV1GetServiceEndpointsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1GetUsageReportParams struct {
	// The start date of the report period (ISO 8601). The report starts from the
	// beginning of this day.
	TimeFrom param.Field[time.Time] `json:"time_from,required" format:"date-time"`
	// The end date of the report period (ISO 8601). The report ends just before the
	// beginning of this day.
	TimeTo param.Field[time.Time] `json:"time_to,required" format:"date-time"`
	// Expenses for the last specified day are taken into account. As the default,
	// False.
	EnableLastDay param.Field[bool] `json:"enable_last_day"`
	// The response resources limit. Defaults to 10.
	Limit param.Field[int64] `json:"limit"`
	// The response resources offset.
	Offset param.Field[int64] `json:"offset"`
	// List of project IDs
	Projects param.Field[[]int64] `json:"projects"`
	// List of region IDs.
	Regions param.Field[[]int64] `json:"regions"`
	// Extended filter for field filtering.
	SchemaFilter param.Field[CloudV1GetUsageReportParamsSchemaFilterUnion] `json:"schema_filter"`
	// List of sorting filters (JSON objects) fields: project. directions: asc, desc.
	Sorting param.Field[[]CloudV1GetUsageReportParamsSorting] `json:"sorting"`
	// Filter by tags
	Tags param.Field[TagsFilterParam] `json:"tags"`
	// List of resource types to be filtered in the report.
	Types param.Field[[]PrebillingResourceTypes] `json:"types"`
}

func (r CloudV1GetUsageReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Extended filter for field filtering.
type CloudV1GetUsageReportParamsSchemaFilter struct {
	// Field name to filter by
	Field  param.Field[CloudV1GetUsageReportParamsSchemaFilterField] `json:"field,required"`
	Type   param.Field[CloudV1GetUsageReportParamsSchemaFilterType]  `json:"type,required"`
	Values param.Field[interface{}]                                  `json:"values,required"`
}

func (r CloudV1GetUsageReportParamsSchemaFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CloudV1GetUsageReportParamsSchemaFilter) implementsCloudV1GetUsageReportParamsSchemaFilterUnion() {
}

// Extended filter for field filtering.
//
// Satisfied by [SchemaFilterSnapshotParam], [SchemaFilterInstanceParam],
// [SchemaFilterAIClusterParam], [SchemaFilterBasicVmParam],
// [SchemaFilterBaremetalParam], [SchemaFilterVolumeParam],
// [SchemaFilterFileShareParam], [SchemaFilterImageParam],
// [SchemaFilterFloatingIPParam], [SchemaFilterEgressTrafficParam],
// [SchemaFilterLoadBalancerParam], [SchemaFilterExternalIPParam],
// [SchemaFilterBackupParam], [SchemaFilterLogIndexParam],
// [SchemaFilterFunctionsParam], [SchemaFilterFunctionsCallsParam],
// [SchemaFilterFunctionsTrafficParam], [SchemaFilterContainersParam],
// [SchemaFilterInferenceParam], [SchemaFilterDbaasPostgresqlVolumeParam],
// [SchemaFilterDbaasPostgresqlPublicNetworkParam],
// [SchemaFilterDbaasPostgresqlCPUParam], [SchemaFilterDbaasPostgresqlMemoryParam],
// [SchemaFilterDbaasPostgresqlPoolerParam],
// [CloudV1GetUsageReportParamsSchemaFilter].
type CloudV1GetUsageReportParamsSchemaFilterUnion interface {
	implementsCloudV1GetUsageReportParamsSchemaFilterUnion()
}

// Field name to filter by
type CloudV1GetUsageReportParamsSchemaFilterField string

const (
	CloudV1GetUsageReportParamsSchemaFilterFieldLastName         CloudV1GetUsageReportParamsSchemaFilterField = "last_name"
	CloudV1GetUsageReportParamsSchemaFilterFieldLastSize         CloudV1GetUsageReportParamsSchemaFilterField = "last_size"
	CloudV1GetUsageReportParamsSchemaFilterFieldSourceVolumeUuid CloudV1GetUsageReportParamsSchemaFilterField = "source_volume_uuid"
	CloudV1GetUsageReportParamsSchemaFilterFieldType             CloudV1GetUsageReportParamsSchemaFilterField = "type"
	CloudV1GetUsageReportParamsSchemaFilterFieldUuid             CloudV1GetUsageReportParamsSchemaFilterField = "uuid"
	CloudV1GetUsageReportParamsSchemaFilterFieldVolumeType       CloudV1GetUsageReportParamsSchemaFilterField = "volume_type"
	CloudV1GetUsageReportParamsSchemaFilterFieldFlavor           CloudV1GetUsageReportParamsSchemaFilterField = "flavor"
	CloudV1GetUsageReportParamsSchemaFilterFieldAttachedToVm     CloudV1GetUsageReportParamsSchemaFilterField = "attached_to_vm"
	CloudV1GetUsageReportParamsSchemaFilterFieldFileShareType    CloudV1GetUsageReportParamsSchemaFilterField = "file_share_type"
	CloudV1GetUsageReportParamsSchemaFilterFieldIPAddress        CloudV1GetUsageReportParamsSchemaFilterField = "ip_address"
	CloudV1GetUsageReportParamsSchemaFilterFieldInstanceName     CloudV1GetUsageReportParamsSchemaFilterField = "instance_name"
	CloudV1GetUsageReportParamsSchemaFilterFieldInstanceType     CloudV1GetUsageReportParamsSchemaFilterField = "instance_type"
	CloudV1GetUsageReportParamsSchemaFilterFieldPortID           CloudV1GetUsageReportParamsSchemaFilterField = "port_id"
	CloudV1GetUsageReportParamsSchemaFilterFieldVmID             CloudV1GetUsageReportParamsSchemaFilterField = "vm_id"
	CloudV1GetUsageReportParamsSchemaFilterFieldNetworkID        CloudV1GetUsageReportParamsSchemaFilterField = "network_id"
	CloudV1GetUsageReportParamsSchemaFilterFieldSubnetID         CloudV1GetUsageReportParamsSchemaFilterField = "subnet_id"
	CloudV1GetUsageReportParamsSchemaFilterFieldScheduleID       CloudV1GetUsageReportParamsSchemaFilterField = "schedule_id"
)

func (r CloudV1GetUsageReportParamsSchemaFilterField) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportParamsSchemaFilterFieldLastName, CloudV1GetUsageReportParamsSchemaFilterFieldLastSize, CloudV1GetUsageReportParamsSchemaFilterFieldSourceVolumeUuid, CloudV1GetUsageReportParamsSchemaFilterFieldType, CloudV1GetUsageReportParamsSchemaFilterFieldUuid, CloudV1GetUsageReportParamsSchemaFilterFieldVolumeType, CloudV1GetUsageReportParamsSchemaFilterFieldFlavor, CloudV1GetUsageReportParamsSchemaFilterFieldAttachedToVm, CloudV1GetUsageReportParamsSchemaFilterFieldFileShareType, CloudV1GetUsageReportParamsSchemaFilterFieldIPAddress, CloudV1GetUsageReportParamsSchemaFilterFieldInstanceName, CloudV1GetUsageReportParamsSchemaFilterFieldInstanceType, CloudV1GetUsageReportParamsSchemaFilterFieldPortID, CloudV1GetUsageReportParamsSchemaFilterFieldVmID, CloudV1GetUsageReportParamsSchemaFilterFieldNetworkID, CloudV1GetUsageReportParamsSchemaFilterFieldSubnetID, CloudV1GetUsageReportParamsSchemaFilterFieldScheduleID:
		return true
	}
	return false
}

type CloudV1GetUsageReportParamsSchemaFilterType string

const (
	CloudV1GetUsageReportParamsSchemaFilterTypeSnapshot                        CloudV1GetUsageReportParamsSchemaFilterType = "snapshot"
	CloudV1GetUsageReportParamsSchemaFilterTypeInstance                        CloudV1GetUsageReportParamsSchemaFilterType = "instance"
	CloudV1GetUsageReportParamsSchemaFilterTypeAICluster                       CloudV1GetUsageReportParamsSchemaFilterType = "ai_cluster"
	CloudV1GetUsageReportParamsSchemaFilterTypeBasicVm                         CloudV1GetUsageReportParamsSchemaFilterType = "basic_vm"
	CloudV1GetUsageReportParamsSchemaFilterTypeBaremetal                       CloudV1GetUsageReportParamsSchemaFilterType = "baremetal"
	CloudV1GetUsageReportParamsSchemaFilterTypeVolume                          CloudV1GetUsageReportParamsSchemaFilterType = "volume"
	CloudV1GetUsageReportParamsSchemaFilterTypeFileShare                       CloudV1GetUsageReportParamsSchemaFilterType = "file_share"
	CloudV1GetUsageReportParamsSchemaFilterTypeImage                           CloudV1GetUsageReportParamsSchemaFilterType = "image"
	CloudV1GetUsageReportParamsSchemaFilterTypeFloatingip                      CloudV1GetUsageReportParamsSchemaFilterType = "floatingip"
	CloudV1GetUsageReportParamsSchemaFilterTypeEgressTraffic                   CloudV1GetUsageReportParamsSchemaFilterType = "egress_traffic"
	CloudV1GetUsageReportParamsSchemaFilterTypeLoadBalancer                    CloudV1GetUsageReportParamsSchemaFilterType = "load_balancer"
	CloudV1GetUsageReportParamsSchemaFilterTypeExternalIP                      CloudV1GetUsageReportParamsSchemaFilterType = "external_ip"
	CloudV1GetUsageReportParamsSchemaFilterTypeBackup                          CloudV1GetUsageReportParamsSchemaFilterType = "backup"
	CloudV1GetUsageReportParamsSchemaFilterTypeLogIndex                        CloudV1GetUsageReportParamsSchemaFilterType = "log_index"
	CloudV1GetUsageReportParamsSchemaFilterTypeFunctions                       CloudV1GetUsageReportParamsSchemaFilterType = "functions"
	CloudV1GetUsageReportParamsSchemaFilterTypeFunctionsCalls                  CloudV1GetUsageReportParamsSchemaFilterType = "functions_calls"
	CloudV1GetUsageReportParamsSchemaFilterTypeFunctionsTraffic                CloudV1GetUsageReportParamsSchemaFilterType = "functions_traffic"
	CloudV1GetUsageReportParamsSchemaFilterTypeContainers                      CloudV1GetUsageReportParamsSchemaFilterType = "containers"
	CloudV1GetUsageReportParamsSchemaFilterTypeInference                       CloudV1GetUsageReportParamsSchemaFilterType = "inference"
	CloudV1GetUsageReportParamsSchemaFilterTypeDbaasPostgresqlVolume           CloudV1GetUsageReportParamsSchemaFilterType = "dbaas_postgresql_volume"
	CloudV1GetUsageReportParamsSchemaFilterTypeDbaasPostgresqlPublicNetwork    CloudV1GetUsageReportParamsSchemaFilterType = "dbaas_postgresql_public_network"
	CloudV1GetUsageReportParamsSchemaFilterTypeDbaasPostgresqlCPU              CloudV1GetUsageReportParamsSchemaFilterType = "dbaas_postgresql_cpu"
	CloudV1GetUsageReportParamsSchemaFilterTypeDbaasPostgresqlMemory           CloudV1GetUsageReportParamsSchemaFilterType = "dbaas_postgresql_memory"
	CloudV1GetUsageReportParamsSchemaFilterTypeDbaasPostgresqlConnectionPooler CloudV1GetUsageReportParamsSchemaFilterType = "dbaas_postgresql_connection_pooler"
)

func (r CloudV1GetUsageReportParamsSchemaFilterType) IsKnown() bool {
	switch r {
	case CloudV1GetUsageReportParamsSchemaFilterTypeSnapshot, CloudV1GetUsageReportParamsSchemaFilterTypeInstance, CloudV1GetUsageReportParamsSchemaFilterTypeAICluster, CloudV1GetUsageReportParamsSchemaFilterTypeBasicVm, CloudV1GetUsageReportParamsSchemaFilterTypeBaremetal, CloudV1GetUsageReportParamsSchemaFilterTypeVolume, CloudV1GetUsageReportParamsSchemaFilterTypeFileShare, CloudV1GetUsageReportParamsSchemaFilterTypeImage, CloudV1GetUsageReportParamsSchemaFilterTypeFloatingip, CloudV1GetUsageReportParamsSchemaFilterTypeEgressTraffic, CloudV1GetUsageReportParamsSchemaFilterTypeLoadBalancer, CloudV1GetUsageReportParamsSchemaFilterTypeExternalIP, CloudV1GetUsageReportParamsSchemaFilterTypeBackup, CloudV1GetUsageReportParamsSchemaFilterTypeLogIndex, CloudV1GetUsageReportParamsSchemaFilterTypeFunctions, CloudV1GetUsageReportParamsSchemaFilterTypeFunctionsCalls, CloudV1GetUsageReportParamsSchemaFilterTypeFunctionsTraffic, CloudV1GetUsageReportParamsSchemaFilterTypeContainers, CloudV1GetUsageReportParamsSchemaFilterTypeInference, CloudV1GetUsageReportParamsSchemaFilterTypeDbaasPostgresqlVolume, CloudV1GetUsageReportParamsSchemaFilterTypeDbaasPostgresqlPublicNetwork, CloudV1GetUsageReportParamsSchemaFilterTypeDbaasPostgresqlCPU, CloudV1GetUsageReportParamsSchemaFilterTypeDbaasPostgresqlMemory, CloudV1GetUsageReportParamsSchemaFilterTypeDbaasPostgresqlConnectionPooler:
		return true
	}
	return false
}

type CloudV1GetUsageReportParamsSorting struct {
	BillingValue param.Field[SortingDirections] `json:"billing_value"`
	FirstSeen    param.Field[SortingDirections] `json:"first_seen"`
	LastName     param.Field[SortingDirections] `json:"last_name"`
	LastSeen     param.Field[SortingDirections] `json:"last_seen"`
	Project      param.Field[SortingDirections] `json:"project"`
	Region       param.Field[SortingDirections] `json:"region"`
	Type         param.Field[SortingDirections] `json:"type"`
}

func (r CloudV1GetUsageReportParamsSorting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1ListProjectImagesParams struct {
	// Show price
	IncludePrices param.Field[bool] `query:"include_prices"`
	// Filter by metadata keys. Must be a valid JSON string. ' curl -G --data-urlencode
	// 'metadata_k=["key1", "key2"]' --url 'http://localhost:1111/v1/images/1/1'
	MetadataK param.Field[string] `query:"metadata_k"`
	// Filter by metadata key-value pairs. Must be a valid JSON string. 'curl -G
	// --data-urlencode 'metadata_kv={"key": "value"}' --url
	// 'http://localhost:1111/v1/images/1/1'"
	MetadataKv param.Field[string] `query:"metadata_kv"`
	// Any value to show private images
	Private param.Field[string] `query:"private"`
	// Image visibility. Globally visible images are public
	Visibility param.Field[CloudV1ListProjectImagesParamsVisibility] `query:"visibility"`
}

// URLQuery serializes [CloudV1ListProjectImagesParams]'s query parameters as
// `url.Values`.
func (r CloudV1ListProjectImagesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Image visibility. Globally visible images are public
type CloudV1ListProjectImagesParamsVisibility string

const (
	CloudV1ListProjectImagesParamsVisibilityPrivate CloudV1ListProjectImagesParamsVisibility = "private"
	CloudV1ListProjectImagesParamsVisibilityPublic  CloudV1ListProjectImagesParamsVisibility = "public"
	CloudV1ListProjectImagesParamsVisibilityShared  CloudV1ListProjectImagesParamsVisibility = "shared"
)

func (r CloudV1ListProjectImagesParamsVisibility) IsKnown() bool {
	switch r {
	case CloudV1ListProjectImagesParamsVisibilityPrivate, CloudV1ListProjectImagesParamsVisibilityPublic, CloudV1ListProjectImagesParamsVisibilityShared:
		return true
	}
	return false
}

type CloudV1UploadImageParams struct {
	// Image name
	Name param.Field[string] `json:"name,required"`
	// URL
	URL param.Field[string] `json:"url,required"`
	// an image architecture type: aarch64, x86_64
	Architecture param.Field[CloudV1UploadImageParamsArchitecture] `json:"architecture"`
	// When True, image cannot be deleted unless all volumes, created from it, are
	// deleted. Defaults to False
	CowFormat param.Field[bool] `json:"cow_format"`
	// Specifies the type of firmware with which to boot the guest.
	HwFirmwareType param.Field[CloudV1UploadImageParamsHwFirmwareType] `json:"hw_firmware_type"`
	// A virtual chipset type.
	HwMachineType param.Field[CloudV1UploadImageParamsHwMachineType] `json:"hw_machine_type"`
	// Set to true if the image will be used by bare metal servers. Defaults to false.
	IsBaremetal param.Field[bool] `json:"is_baremetal"`
	// Create one or more metadata items for an image
	Metadata param.Field[interface{}] `json:"metadata"`
	// OS Distribution, i.e. Debian, CentOS, Ubuntu, CoreOS etc.
	OsDistro param.Field[string] `json:"os_distro"`
	// The operating system installed on the image.
	OsType param.Field[CloudV1UploadImageParamsOsType] `json:"os_type"`
	// OS version, i.e. 19.04 (for Ubuntu) or 9.4 for Debian
	OsVersion param.Field[string] `json:"os_version"`
	// Permission to use a ssh key in instances
	SSHKey param.Field[CloudV1UploadImageParamsSSHKey] `json:"ssh_key"`
}

func (r CloudV1UploadImageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// an image architecture type: aarch64, x86_64
type CloudV1UploadImageParamsArchitecture string

const (
	CloudV1UploadImageParamsArchitectureAarch64 CloudV1UploadImageParamsArchitecture = "aarch64"
	CloudV1UploadImageParamsArchitectureX86_64  CloudV1UploadImageParamsArchitecture = "x86_64"
)

func (r CloudV1UploadImageParamsArchitecture) IsKnown() bool {
	switch r {
	case CloudV1UploadImageParamsArchitectureAarch64, CloudV1UploadImageParamsArchitectureX86_64:
		return true
	}
	return false
}

// Specifies the type of firmware with which to boot the guest.
type CloudV1UploadImageParamsHwFirmwareType string

const (
	CloudV1UploadImageParamsHwFirmwareTypeBios CloudV1UploadImageParamsHwFirmwareType = "bios"
	CloudV1UploadImageParamsHwFirmwareTypeUefi CloudV1UploadImageParamsHwFirmwareType = "uefi"
)

func (r CloudV1UploadImageParamsHwFirmwareType) IsKnown() bool {
	switch r {
	case CloudV1UploadImageParamsHwFirmwareTypeBios, CloudV1UploadImageParamsHwFirmwareTypeUefi:
		return true
	}
	return false
}

// A virtual chipset type.
type CloudV1UploadImageParamsHwMachineType string

const (
	CloudV1UploadImageParamsHwMachineTypeI440 CloudV1UploadImageParamsHwMachineType = "i440"
	CloudV1UploadImageParamsHwMachineTypeQ35  CloudV1UploadImageParamsHwMachineType = "q35"
)

func (r CloudV1UploadImageParamsHwMachineType) IsKnown() bool {
	switch r {
	case CloudV1UploadImageParamsHwMachineTypeI440, CloudV1UploadImageParamsHwMachineTypeQ35:
		return true
	}
	return false
}

// The operating system installed on the image.
type CloudV1UploadImageParamsOsType string

const (
	CloudV1UploadImageParamsOsTypeLinux   CloudV1UploadImageParamsOsType = "linux"
	CloudV1UploadImageParamsOsTypeWindows CloudV1UploadImageParamsOsType = "windows"
)

func (r CloudV1UploadImageParamsOsType) IsKnown() bool {
	switch r {
	case CloudV1UploadImageParamsOsTypeLinux, CloudV1UploadImageParamsOsTypeWindows:
		return true
	}
	return false
}

// Permission to use a ssh key in instances
type CloudV1UploadImageParamsSSHKey string

const (
	CloudV1UploadImageParamsSSHKeyAllow    CloudV1UploadImageParamsSSHKey = "allow"
	CloudV1UploadImageParamsSSHKeyDeny     CloudV1UploadImageParamsSSHKey = "deny"
	CloudV1UploadImageParamsSSHKeyRequired CloudV1UploadImageParamsSSHKey = "required"
)

func (r CloudV1UploadImageParamsSSHKey) IsKnown() bool {
	switch r {
	case CloudV1UploadImageParamsSSHKeyAllow, CloudV1UploadImageParamsSSHKeyDeny, CloudV1UploadImageParamsSSHKeyRequired:
		return true
	}
	return false
}
