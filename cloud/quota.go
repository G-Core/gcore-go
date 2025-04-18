// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// QuotaService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuotaService] method instead.
type QuotaService struct {
	Options  []option.RequestOption
	Requests QuotaRequestService
}

// NewQuotaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQuotaService(opts ...option.RequestOption) (r QuotaService) {
	r = QuotaService{}
	r.Options = opts
	r.Requests = NewQuotaRequestService(opts...)
	return
}

// Get combined client quotas, regional and global.
func (r *QuotaService) GetAll(ctx context.Context, opts ...option.RequestOption) (res *QuotaGetAllResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/client_quotas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a quota by region
func (r *QuotaService) GetByRegion(ctx context.Context, clientID int64, query QuotaGetByRegionParams, opts ...option.RequestOption) (res *QuotaGetByRegionResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.RegionID, precfg.RegionID)
	if !query.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/regional_quotas/%v/%v", clientID, query.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get global quota
func (r *QuotaService) GetGlobal(ctx context.Context, clientID int64, opts ...option.RequestOption) (res *QuotaGetGlobalResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/global_quotas/%v", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// '#/paths/%2Fcloud%2Fv2%2Fclient_quotas/get/responses/200/content/application%2Fjson/schema'
// "$.paths['/cloud/v2/client_quotas'].get.responses[200].content['application/json'].schema"
type QuotaGetAllResponse struct {
	// '#/components/schemas/AllClientQuotasSerializer/properties/global_quotas'
	// "$.components.schemas.AllClientQuotasSerializer.properties.global_quotas"
	GlobalQuotas QuotaGetAllResponseGlobalQuotas `json:"global_quotas"`
	// '#/components/schemas/AllClientQuotasSerializer/properties/regional_quotas'
	// "$.components.schemas.AllClientQuotasSerializer.properties.regional_quotas"
	RegionalQuotas []QuotaGetAllResponseRegionalQuota `json:"regional_quotas"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		GlobalQuotas   resp.Field
		RegionalQuotas resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaGetAllResponse) RawJSON() string { return r.JSON.raw }
func (r *QuotaGetAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/AllClientQuotasSerializer/properties/global_quotas'
// "$.components.schemas.AllClientQuotasSerializer.properties.global_quotas"
type QuotaGetAllResponseGlobalQuotas struct {
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_cpu_millicore_count_limit'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_cpu_millicore_count_limit"
	InferenceCPUMillicoreCountLimit int64 `json:"inference_cpu_millicore_count_limit"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_cpu_millicore_count_usage'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_cpu_millicore_count_usage"
	InferenceCPUMillicoreCountUsage int64 `json:"inference_cpu_millicore_count_usage"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_gpu_a100_count_limit'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_gpu_a100_count_limit"
	InferenceGPUA100CountLimit int64 `json:"inference_gpu_a100_count_limit"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_gpu_a100_count_usage'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_gpu_a100_count_usage"
	InferenceGPUA100CountUsage int64 `json:"inference_gpu_a100_count_usage"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_gpu_h100_count_limit'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_gpu_h100_count_limit"
	InferenceGPUH100CountLimit int64 `json:"inference_gpu_h100_count_limit"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_gpu_h100_count_usage'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_gpu_h100_count_usage"
	InferenceGPUH100CountUsage int64 `json:"inference_gpu_h100_count_usage"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_gpu_l40s_count_limit'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_gpu_l40s_count_limit"
	InferenceGPUL40sCountLimit int64 `json:"inference_gpu_l40s_count_limit"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_gpu_l40s_count_usage'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_gpu_l40s_count_usage"
	InferenceGPUL40sCountUsage int64 `json:"inference_gpu_l40s_count_usage"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_instance_count_limit'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_instance_count_limit"
	InferenceInstanceCountLimit int64 `json:"inference_instance_count_limit"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_instance_count_usage'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_instance_count_usage"
	InferenceInstanceCountUsage int64 `json:"inference_instance_count_usage"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/keypair_count_limit'
	// "$.components.schemas.GlobalQuotasSerializer.properties.keypair_count_limit"
	KeypairCountLimit int64 `json:"keypair_count_limit"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/keypair_count_usage'
	// "$.components.schemas.GlobalQuotasSerializer.properties.keypair_count_usage"
	KeypairCountUsage int64 `json:"keypair_count_usage"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/project_count_limit'
	// "$.components.schemas.GlobalQuotasSerializer.properties.project_count_limit"
	ProjectCountLimit int64 `json:"project_count_limit"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/project_count_usage'
	// "$.components.schemas.GlobalQuotasSerializer.properties.project_count_usage"
	ProjectCountUsage int64 `json:"project_count_usage"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		InferenceCPUMillicoreCountLimit resp.Field
		InferenceCPUMillicoreCountUsage resp.Field
		InferenceGPUA100CountLimit      resp.Field
		InferenceGPUA100CountUsage      resp.Field
		InferenceGPUH100CountLimit      resp.Field
		InferenceGPUH100CountUsage      resp.Field
		InferenceGPUL40sCountLimit      resp.Field
		InferenceGPUL40sCountUsage      resp.Field
		InferenceInstanceCountLimit     resp.Field
		InferenceInstanceCountUsage     resp.Field
		KeypairCountLimit               resp.Field
		KeypairCountUsage               resp.Field
		ProjectCountLimit               resp.Field
		ProjectCountUsage               resp.Field
		ExtraFields                     map[string]resp.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaGetAllResponseGlobalQuotas) RawJSON() string { return r.JSON.raw }
func (r *QuotaGetAllResponseGlobalQuotas) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/AllClientQuotasSerializer/properties/regional_quotas/items'
// "$.components.schemas.AllClientQuotasSerializer.properties.regional_quotas.items"
type QuotaGetAllResponseRegionalQuota struct {
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_basic_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_basic_count_limit"
	BaremetalBasicCountLimit int64 `json:"baremetal_basic_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_basic_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_basic_count_usage"
	BaremetalBasicCountUsage int64 `json:"baremetal_basic_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_gpu_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_gpu_count_limit"
	BaremetalGPUCountLimit int64 `json:"baremetal_gpu_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_gpu_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_gpu_count_usage"
	BaremetalGPUCountUsage int64 `json:"baremetal_gpu_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_hf_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_hf_count_limit"
	BaremetalHfCountLimit int64 `json:"baremetal_hf_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_hf_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_hf_count_usage"
	BaremetalHfCountUsage int64 `json:"baremetal_hf_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_infrastructure_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_infrastructure_count_limit"
	BaremetalInfrastructureCountLimit int64 `json:"baremetal_infrastructure_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_infrastructure_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_infrastructure_count_usage"
	BaremetalInfrastructureCountUsage int64 `json:"baremetal_infrastructure_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_network_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_network_count_limit"
	BaremetalNetworkCountLimit int64 `json:"baremetal_network_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_network_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_network_count_usage"
	BaremetalNetworkCountUsage int64 `json:"baremetal_network_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_storage_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_storage_count_limit"
	BaremetalStorageCountLimit int64 `json:"baremetal_storage_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_storage_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_storage_count_usage"
	BaremetalStorageCountUsage int64 `json:"baremetal_storage_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_container_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_container_count_limit"
	CaasContainerCountLimit int64 `json:"caas_container_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_container_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_container_count_usage"
	CaasContainerCountUsage int64 `json:"caas_container_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_cpu_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_cpu_count_limit"
	CaasCPUCountLimit int64 `json:"caas_cpu_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_cpu_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_cpu_count_usage"
	CaasCPUCountUsage int64 `json:"caas_cpu_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_gpu_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_gpu_count_limit"
	CaasGPUCountLimit int64 `json:"caas_gpu_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_gpu_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_gpu_count_usage"
	CaasGPUCountUsage int64 `json:"caas_gpu_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_ram_size_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_ram_size_limit"
	CaasRamSizeLimit int64 `json:"caas_ram_size_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_ram_size_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_ram_size_usage"
	CaasRamSizeUsage int64 `json:"caas_ram_size_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/cluster_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.cluster_count_limit"
	ClusterCountLimit int64 `json:"cluster_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/cluster_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.cluster_count_usage"
	ClusterCountUsage int64 `json:"cluster_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/cpu_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.cpu_count_limit"
	CPUCountLimit int64 `json:"cpu_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/cpu_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.cpu_count_usage"
	CPUCountUsage int64 `json:"cpu_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/dbaas_postgres_cluster_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.dbaas_postgres_cluster_count_limit"
	DbaasPostgresClusterCountLimit int64 `json:"dbaas_postgres_cluster_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/dbaas_postgres_cluster_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.dbaas_postgres_cluster_count_usage"
	DbaasPostgresClusterCountUsage int64 `json:"dbaas_postgres_cluster_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/external_ip_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.external_ip_count_limit"
	ExternalIPCountLimit int64 `json:"external_ip_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/external_ip_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.external_ip_count_usage"
	ExternalIPCountUsage int64 `json:"external_ip_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_cpu_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_cpu_count_limit"
	FaasCPUCountLimit int64 `json:"faas_cpu_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_cpu_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_cpu_count_usage"
	FaasCPUCountUsage int64 `json:"faas_cpu_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_function_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_function_count_limit"
	FaasFunctionCountLimit int64 `json:"faas_function_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_function_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_function_count_usage"
	FaasFunctionCountUsage int64 `json:"faas_function_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_namespace_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_namespace_count_limit"
	FaasNamespaceCountLimit int64 `json:"faas_namespace_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_namespace_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_namespace_count_usage"
	FaasNamespaceCountUsage int64 `json:"faas_namespace_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_ram_size_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_ram_size_limit"
	FaasRamSizeLimit int64 `json:"faas_ram_size_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_ram_size_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_ram_size_usage"
	FaasRamSizeUsage int64 `json:"faas_ram_size_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/firewall_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.firewall_count_limit"
	FirewallCountLimit int64 `json:"firewall_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/firewall_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.firewall_count_usage"
	FirewallCountUsage int64 `json:"firewall_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/floating_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.floating_count_limit"
	FloatingCountLimit int64 `json:"floating_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/floating_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.floating_count_usage"
	FloatingCountUsage int64 `json:"floating_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_count_limit"
	GPUCountLimit int64 `json:"gpu_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_count_usage"
	GPUCountUsage int64 `json:"gpu_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_virtual_a100_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_virtual_a100_count_limit"
	GPUVirtualA100CountLimit int64 `json:"gpu_virtual_a100_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_virtual_a100_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_virtual_a100_count_usage"
	GPUVirtualA100CountUsage int64 `json:"gpu_virtual_a100_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_virtual_h100_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_virtual_h100_count_limit"
	GPUVirtualH100CountLimit int64 `json:"gpu_virtual_h100_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_virtual_h100_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_virtual_h100_count_usage"
	GPUVirtualH100CountUsage int64 `json:"gpu_virtual_h100_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_virtual_l40s_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_virtual_l40s_count_limit"
	GPUVirtualL40sCountLimit int64 `json:"gpu_virtual_l40s_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_virtual_l40s_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_virtual_l40s_count_usage"
	GPUVirtualL40sCountUsage int64 `json:"gpu_virtual_l40s_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/image_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.image_count_limit"
	ImageCountLimit int64 `json:"image_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/image_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.image_count_usage"
	ImageCountUsage int64 `json:"image_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/image_size_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.image_size_limit"
	ImageSizeLimit int64 `json:"image_size_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/image_size_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.image_size_usage"
	ImageSizeUsage int64 `json:"image_size_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/ipu_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.ipu_count_limit"
	IpuCountLimit int64 `json:"ipu_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/ipu_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.ipu_count_usage"
	IpuCountUsage int64 `json:"ipu_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/laas_topic_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.laas_topic_count_limit"
	LaasTopicCountLimit int64 `json:"laas_topic_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/laas_topic_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.laas_topic_count_usage"
	LaasTopicCountUsage int64 `json:"laas_topic_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/loadbalancer_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.loadbalancer_count_limit"
	LoadbalancerCountLimit int64 `json:"loadbalancer_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/loadbalancer_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.loadbalancer_count_usage"
	LoadbalancerCountUsage int64 `json:"loadbalancer_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/network_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.network_count_limit"
	NetworkCountLimit int64 `json:"network_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/network_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.network_count_usage"
	NetworkCountUsage int64 `json:"network_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/ram_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.ram_limit"
	RamLimit int64 `json:"ram_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/ram_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.ram_usage"
	RamUsage int64 `json:"ram_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/region_id'
	// "$.components.schemas.RegionalQuotasSerializer.properties.region_id"
	RegionID int64 `json:"region_id"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/registry_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.registry_count_limit"
	RegistryCountLimit int64 `json:"registry_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/registry_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.registry_count_usage"
	RegistryCountUsage int64 `json:"registry_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/registry_storage_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.registry_storage_limit"
	RegistryStorageLimit int64 `json:"registry_storage_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/registry_storage_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.registry_storage_usage"
	RegistryStorageUsage int64 `json:"registry_storage_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/router_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.router_count_limit"
	RouterCountLimit int64 `json:"router_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/router_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.router_count_usage"
	RouterCountUsage int64 `json:"router_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/secret_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.secret_count_limit"
	SecretCountLimit int64 `json:"secret_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/secret_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.secret_count_usage"
	SecretCountUsage int64 `json:"secret_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/servergroup_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.servergroup_count_limit"
	ServergroupCountLimit int64 `json:"servergroup_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/servergroup_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.servergroup_count_usage"
	ServergroupCountUsage int64 `json:"servergroup_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/sfs_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.sfs_count_limit"
	SfsCountLimit int64 `json:"sfs_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/sfs_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.sfs_count_usage"
	SfsCountUsage int64 `json:"sfs_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/sfs_size_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.sfs_size_limit"
	SfsSizeLimit int64 `json:"sfs_size_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/sfs_size_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.sfs_size_usage"
	SfsSizeUsage int64 `json:"sfs_size_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/shared_vm_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.shared_vm_count_limit"
	SharedVmCountLimit int64 `json:"shared_vm_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/shared_vm_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.shared_vm_count_usage"
	SharedVmCountUsage int64 `json:"shared_vm_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/snapshot_schedule_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.snapshot_schedule_count_limit"
	SnapshotScheduleCountLimit int64 `json:"snapshot_schedule_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/snapshot_schedule_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.snapshot_schedule_count_usage"
	SnapshotScheduleCountUsage int64 `json:"snapshot_schedule_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/subnet_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.subnet_count_limit"
	SubnetCountLimit int64 `json:"subnet_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/subnet_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.subnet_count_usage"
	SubnetCountUsage int64 `json:"subnet_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/vm_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.vm_count_limit"
	VmCountLimit int64 `json:"vm_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/vm_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.vm_count_usage"
	VmCountUsage int64 `json:"vm_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_count_limit"
	VolumeCountLimit int64 `json:"volume_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_count_usage"
	VolumeCountUsage int64 `json:"volume_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_size_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_size_limit"
	VolumeSizeLimit int64 `json:"volume_size_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_size_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_size_usage"
	VolumeSizeUsage int64 `json:"volume_size_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_snapshots_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_snapshots_count_limit"
	VolumeSnapshotsCountLimit int64 `json:"volume_snapshots_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_snapshots_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_snapshots_count_usage"
	VolumeSnapshotsCountUsage int64 `json:"volume_snapshots_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_snapshots_size_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_snapshots_size_limit"
	VolumeSnapshotsSizeLimit int64 `json:"volume_snapshots_size_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_snapshots_size_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_snapshots_size_usage"
	VolumeSnapshotsSizeUsage int64 `json:"volume_snapshots_size_usage"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		BaremetalBasicCountLimit          resp.Field
		BaremetalBasicCountUsage          resp.Field
		BaremetalGPUCountLimit            resp.Field
		BaremetalGPUCountUsage            resp.Field
		BaremetalHfCountLimit             resp.Field
		BaremetalHfCountUsage             resp.Field
		BaremetalInfrastructureCountLimit resp.Field
		BaremetalInfrastructureCountUsage resp.Field
		BaremetalNetworkCountLimit        resp.Field
		BaremetalNetworkCountUsage        resp.Field
		BaremetalStorageCountLimit        resp.Field
		BaremetalStorageCountUsage        resp.Field
		CaasContainerCountLimit           resp.Field
		CaasContainerCountUsage           resp.Field
		CaasCPUCountLimit                 resp.Field
		CaasCPUCountUsage                 resp.Field
		CaasGPUCountLimit                 resp.Field
		CaasGPUCountUsage                 resp.Field
		CaasRamSizeLimit                  resp.Field
		CaasRamSizeUsage                  resp.Field
		ClusterCountLimit                 resp.Field
		ClusterCountUsage                 resp.Field
		CPUCountLimit                     resp.Field
		CPUCountUsage                     resp.Field
		DbaasPostgresClusterCountLimit    resp.Field
		DbaasPostgresClusterCountUsage    resp.Field
		ExternalIPCountLimit              resp.Field
		ExternalIPCountUsage              resp.Field
		FaasCPUCountLimit                 resp.Field
		FaasCPUCountUsage                 resp.Field
		FaasFunctionCountLimit            resp.Field
		FaasFunctionCountUsage            resp.Field
		FaasNamespaceCountLimit           resp.Field
		FaasNamespaceCountUsage           resp.Field
		FaasRamSizeLimit                  resp.Field
		FaasRamSizeUsage                  resp.Field
		FirewallCountLimit                resp.Field
		FirewallCountUsage                resp.Field
		FloatingCountLimit                resp.Field
		FloatingCountUsage                resp.Field
		GPUCountLimit                     resp.Field
		GPUCountUsage                     resp.Field
		GPUVirtualA100CountLimit          resp.Field
		GPUVirtualA100CountUsage          resp.Field
		GPUVirtualH100CountLimit          resp.Field
		GPUVirtualH100CountUsage          resp.Field
		GPUVirtualL40sCountLimit          resp.Field
		GPUVirtualL40sCountUsage          resp.Field
		ImageCountLimit                   resp.Field
		ImageCountUsage                   resp.Field
		ImageSizeLimit                    resp.Field
		ImageSizeUsage                    resp.Field
		IpuCountLimit                     resp.Field
		IpuCountUsage                     resp.Field
		LaasTopicCountLimit               resp.Field
		LaasTopicCountUsage               resp.Field
		LoadbalancerCountLimit            resp.Field
		LoadbalancerCountUsage            resp.Field
		NetworkCountLimit                 resp.Field
		NetworkCountUsage                 resp.Field
		RamLimit                          resp.Field
		RamUsage                          resp.Field
		RegionID                          resp.Field
		RegistryCountLimit                resp.Field
		RegistryCountUsage                resp.Field
		RegistryStorageLimit              resp.Field
		RegistryStorageUsage              resp.Field
		RouterCountLimit                  resp.Field
		RouterCountUsage                  resp.Field
		SecretCountLimit                  resp.Field
		SecretCountUsage                  resp.Field
		ServergroupCountLimit             resp.Field
		ServergroupCountUsage             resp.Field
		SfsCountLimit                     resp.Field
		SfsCountUsage                     resp.Field
		SfsSizeLimit                      resp.Field
		SfsSizeUsage                      resp.Field
		SharedVmCountLimit                resp.Field
		SharedVmCountUsage                resp.Field
		SnapshotScheduleCountLimit        resp.Field
		SnapshotScheduleCountUsage        resp.Field
		SubnetCountLimit                  resp.Field
		SubnetCountUsage                  resp.Field
		VmCountLimit                      resp.Field
		VmCountUsage                      resp.Field
		VolumeCountLimit                  resp.Field
		VolumeCountUsage                  resp.Field
		VolumeSizeLimit                   resp.Field
		VolumeSizeUsage                   resp.Field
		VolumeSnapshotsCountLimit         resp.Field
		VolumeSnapshotsCountUsage         resp.Field
		VolumeSnapshotsSizeLimit          resp.Field
		VolumeSnapshotsSizeUsage          resp.Field
		ExtraFields                       map[string]resp.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaGetAllResponseRegionalQuota) RawJSON() string { return r.JSON.raw }
func (r *QuotaGetAllResponseRegionalQuota) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/paths/%2Fcloud%2Fv2%2Fregional_quotas%2F%7Bclient_id%7D%2F%7Bregion_id%7D/get/responses/200/content/application%2Fjson/schema'
// "$.paths['/cloud/v2/regional_quotas/{client_id}/{region_id}'].get.responses[200].content['application/json'].schema"
type QuotaGetByRegionResponse struct {
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_basic_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_basic_count_limit"
	BaremetalBasicCountLimit int64 `json:"baremetal_basic_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_basic_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_basic_count_usage"
	BaremetalBasicCountUsage int64 `json:"baremetal_basic_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_gpu_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_gpu_count_limit"
	BaremetalGPUCountLimit int64 `json:"baremetal_gpu_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_gpu_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_gpu_count_usage"
	BaremetalGPUCountUsage int64 `json:"baremetal_gpu_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_hf_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_hf_count_limit"
	BaremetalHfCountLimit int64 `json:"baremetal_hf_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_hf_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_hf_count_usage"
	BaremetalHfCountUsage int64 `json:"baremetal_hf_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_infrastructure_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_infrastructure_count_limit"
	BaremetalInfrastructureCountLimit int64 `json:"baremetal_infrastructure_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_infrastructure_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_infrastructure_count_usage"
	BaremetalInfrastructureCountUsage int64 `json:"baremetal_infrastructure_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_network_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_network_count_limit"
	BaremetalNetworkCountLimit int64 `json:"baremetal_network_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_network_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_network_count_usage"
	BaremetalNetworkCountUsage int64 `json:"baremetal_network_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_storage_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_storage_count_limit"
	BaremetalStorageCountLimit int64 `json:"baremetal_storage_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/baremetal_storage_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.baremetal_storage_count_usage"
	BaremetalStorageCountUsage int64 `json:"baremetal_storage_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_container_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_container_count_limit"
	CaasContainerCountLimit int64 `json:"caas_container_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_container_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_container_count_usage"
	CaasContainerCountUsage int64 `json:"caas_container_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_cpu_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_cpu_count_limit"
	CaasCPUCountLimit int64 `json:"caas_cpu_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_cpu_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_cpu_count_usage"
	CaasCPUCountUsage int64 `json:"caas_cpu_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_gpu_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_gpu_count_limit"
	CaasGPUCountLimit int64 `json:"caas_gpu_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_gpu_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_gpu_count_usage"
	CaasGPUCountUsage int64 `json:"caas_gpu_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_ram_size_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_ram_size_limit"
	CaasRamSizeLimit int64 `json:"caas_ram_size_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/caas_ram_size_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.caas_ram_size_usage"
	CaasRamSizeUsage int64 `json:"caas_ram_size_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/cluster_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.cluster_count_limit"
	ClusterCountLimit int64 `json:"cluster_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/cluster_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.cluster_count_usage"
	ClusterCountUsage int64 `json:"cluster_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/cpu_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.cpu_count_limit"
	CPUCountLimit int64 `json:"cpu_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/cpu_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.cpu_count_usage"
	CPUCountUsage int64 `json:"cpu_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/dbaas_postgres_cluster_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.dbaas_postgres_cluster_count_limit"
	DbaasPostgresClusterCountLimit int64 `json:"dbaas_postgres_cluster_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/dbaas_postgres_cluster_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.dbaas_postgres_cluster_count_usage"
	DbaasPostgresClusterCountUsage int64 `json:"dbaas_postgres_cluster_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/external_ip_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.external_ip_count_limit"
	ExternalIPCountLimit int64 `json:"external_ip_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/external_ip_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.external_ip_count_usage"
	ExternalIPCountUsage int64 `json:"external_ip_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_cpu_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_cpu_count_limit"
	FaasCPUCountLimit int64 `json:"faas_cpu_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_cpu_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_cpu_count_usage"
	FaasCPUCountUsage int64 `json:"faas_cpu_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_function_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_function_count_limit"
	FaasFunctionCountLimit int64 `json:"faas_function_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_function_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_function_count_usage"
	FaasFunctionCountUsage int64 `json:"faas_function_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_namespace_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_namespace_count_limit"
	FaasNamespaceCountLimit int64 `json:"faas_namespace_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_namespace_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_namespace_count_usage"
	FaasNamespaceCountUsage int64 `json:"faas_namespace_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_ram_size_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_ram_size_limit"
	FaasRamSizeLimit int64 `json:"faas_ram_size_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/faas_ram_size_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.faas_ram_size_usage"
	FaasRamSizeUsage int64 `json:"faas_ram_size_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/firewall_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.firewall_count_limit"
	FirewallCountLimit int64 `json:"firewall_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/firewall_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.firewall_count_usage"
	FirewallCountUsage int64 `json:"firewall_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/floating_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.floating_count_limit"
	FloatingCountLimit int64 `json:"floating_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/floating_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.floating_count_usage"
	FloatingCountUsage int64 `json:"floating_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_count_limit"
	GPUCountLimit int64 `json:"gpu_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_count_usage"
	GPUCountUsage int64 `json:"gpu_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_virtual_a100_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_virtual_a100_count_limit"
	GPUVirtualA100CountLimit int64 `json:"gpu_virtual_a100_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_virtual_a100_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_virtual_a100_count_usage"
	GPUVirtualA100CountUsage int64 `json:"gpu_virtual_a100_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_virtual_h100_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_virtual_h100_count_limit"
	GPUVirtualH100CountLimit int64 `json:"gpu_virtual_h100_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_virtual_h100_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_virtual_h100_count_usage"
	GPUVirtualH100CountUsage int64 `json:"gpu_virtual_h100_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_virtual_l40s_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_virtual_l40s_count_limit"
	GPUVirtualL40sCountLimit int64 `json:"gpu_virtual_l40s_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/gpu_virtual_l40s_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.gpu_virtual_l40s_count_usage"
	GPUVirtualL40sCountUsage int64 `json:"gpu_virtual_l40s_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/image_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.image_count_limit"
	ImageCountLimit int64 `json:"image_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/image_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.image_count_usage"
	ImageCountUsage int64 `json:"image_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/image_size_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.image_size_limit"
	ImageSizeLimit int64 `json:"image_size_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/image_size_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.image_size_usage"
	ImageSizeUsage int64 `json:"image_size_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/ipu_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.ipu_count_limit"
	IpuCountLimit int64 `json:"ipu_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/ipu_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.ipu_count_usage"
	IpuCountUsage int64 `json:"ipu_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/laas_topic_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.laas_topic_count_limit"
	LaasTopicCountLimit int64 `json:"laas_topic_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/laas_topic_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.laas_topic_count_usage"
	LaasTopicCountUsage int64 `json:"laas_topic_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/loadbalancer_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.loadbalancer_count_limit"
	LoadbalancerCountLimit int64 `json:"loadbalancer_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/loadbalancer_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.loadbalancer_count_usage"
	LoadbalancerCountUsage int64 `json:"loadbalancer_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/network_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.network_count_limit"
	NetworkCountLimit int64 `json:"network_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/network_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.network_count_usage"
	NetworkCountUsage int64 `json:"network_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/ram_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.ram_limit"
	RamLimit int64 `json:"ram_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/ram_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.ram_usage"
	RamUsage int64 `json:"ram_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/region_id'
	// "$.components.schemas.RegionalQuotasSerializer.properties.region_id"
	RegionID int64 `json:"region_id"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/registry_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.registry_count_limit"
	RegistryCountLimit int64 `json:"registry_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/registry_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.registry_count_usage"
	RegistryCountUsage int64 `json:"registry_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/registry_storage_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.registry_storage_limit"
	RegistryStorageLimit int64 `json:"registry_storage_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/registry_storage_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.registry_storage_usage"
	RegistryStorageUsage int64 `json:"registry_storage_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/router_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.router_count_limit"
	RouterCountLimit int64 `json:"router_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/router_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.router_count_usage"
	RouterCountUsage int64 `json:"router_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/secret_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.secret_count_limit"
	SecretCountLimit int64 `json:"secret_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/secret_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.secret_count_usage"
	SecretCountUsage int64 `json:"secret_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/servergroup_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.servergroup_count_limit"
	ServergroupCountLimit int64 `json:"servergroup_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/servergroup_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.servergroup_count_usage"
	ServergroupCountUsage int64 `json:"servergroup_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/sfs_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.sfs_count_limit"
	SfsCountLimit int64 `json:"sfs_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/sfs_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.sfs_count_usage"
	SfsCountUsage int64 `json:"sfs_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/sfs_size_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.sfs_size_limit"
	SfsSizeLimit int64 `json:"sfs_size_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/sfs_size_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.sfs_size_usage"
	SfsSizeUsage int64 `json:"sfs_size_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/shared_vm_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.shared_vm_count_limit"
	SharedVmCountLimit int64 `json:"shared_vm_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/shared_vm_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.shared_vm_count_usage"
	SharedVmCountUsage int64 `json:"shared_vm_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/snapshot_schedule_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.snapshot_schedule_count_limit"
	SnapshotScheduleCountLimit int64 `json:"snapshot_schedule_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/snapshot_schedule_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.snapshot_schedule_count_usage"
	SnapshotScheduleCountUsage int64 `json:"snapshot_schedule_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/subnet_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.subnet_count_limit"
	SubnetCountLimit int64 `json:"subnet_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/subnet_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.subnet_count_usage"
	SubnetCountUsage int64 `json:"subnet_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/vm_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.vm_count_limit"
	VmCountLimit int64 `json:"vm_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/vm_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.vm_count_usage"
	VmCountUsage int64 `json:"vm_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_count_limit"
	VolumeCountLimit int64 `json:"volume_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_count_usage"
	VolumeCountUsage int64 `json:"volume_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_size_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_size_limit"
	VolumeSizeLimit int64 `json:"volume_size_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_size_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_size_usage"
	VolumeSizeUsage int64 `json:"volume_size_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_snapshots_count_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_snapshots_count_limit"
	VolumeSnapshotsCountLimit int64 `json:"volume_snapshots_count_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_snapshots_count_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_snapshots_count_usage"
	VolumeSnapshotsCountUsage int64 `json:"volume_snapshots_count_usage"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_snapshots_size_limit'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_snapshots_size_limit"
	VolumeSnapshotsSizeLimit int64 `json:"volume_snapshots_size_limit"`
	// '#/components/schemas/RegionalQuotasSerializer/properties/volume_snapshots_size_usage'
	// "$.components.schemas.RegionalQuotasSerializer.properties.volume_snapshots_size_usage"
	VolumeSnapshotsSizeUsage int64 `json:"volume_snapshots_size_usage"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		BaremetalBasicCountLimit          resp.Field
		BaremetalBasicCountUsage          resp.Field
		BaremetalGPUCountLimit            resp.Field
		BaremetalGPUCountUsage            resp.Field
		BaremetalHfCountLimit             resp.Field
		BaremetalHfCountUsage             resp.Field
		BaremetalInfrastructureCountLimit resp.Field
		BaremetalInfrastructureCountUsage resp.Field
		BaremetalNetworkCountLimit        resp.Field
		BaremetalNetworkCountUsage        resp.Field
		BaremetalStorageCountLimit        resp.Field
		BaremetalStorageCountUsage        resp.Field
		CaasContainerCountLimit           resp.Field
		CaasContainerCountUsage           resp.Field
		CaasCPUCountLimit                 resp.Field
		CaasCPUCountUsage                 resp.Field
		CaasGPUCountLimit                 resp.Field
		CaasGPUCountUsage                 resp.Field
		CaasRamSizeLimit                  resp.Field
		CaasRamSizeUsage                  resp.Field
		ClusterCountLimit                 resp.Field
		ClusterCountUsage                 resp.Field
		CPUCountLimit                     resp.Field
		CPUCountUsage                     resp.Field
		DbaasPostgresClusterCountLimit    resp.Field
		DbaasPostgresClusterCountUsage    resp.Field
		ExternalIPCountLimit              resp.Field
		ExternalIPCountUsage              resp.Field
		FaasCPUCountLimit                 resp.Field
		FaasCPUCountUsage                 resp.Field
		FaasFunctionCountLimit            resp.Field
		FaasFunctionCountUsage            resp.Field
		FaasNamespaceCountLimit           resp.Field
		FaasNamespaceCountUsage           resp.Field
		FaasRamSizeLimit                  resp.Field
		FaasRamSizeUsage                  resp.Field
		FirewallCountLimit                resp.Field
		FirewallCountUsage                resp.Field
		FloatingCountLimit                resp.Field
		FloatingCountUsage                resp.Field
		GPUCountLimit                     resp.Field
		GPUCountUsage                     resp.Field
		GPUVirtualA100CountLimit          resp.Field
		GPUVirtualA100CountUsage          resp.Field
		GPUVirtualH100CountLimit          resp.Field
		GPUVirtualH100CountUsage          resp.Field
		GPUVirtualL40sCountLimit          resp.Field
		GPUVirtualL40sCountUsage          resp.Field
		ImageCountLimit                   resp.Field
		ImageCountUsage                   resp.Field
		ImageSizeLimit                    resp.Field
		ImageSizeUsage                    resp.Field
		IpuCountLimit                     resp.Field
		IpuCountUsage                     resp.Field
		LaasTopicCountLimit               resp.Field
		LaasTopicCountUsage               resp.Field
		LoadbalancerCountLimit            resp.Field
		LoadbalancerCountUsage            resp.Field
		NetworkCountLimit                 resp.Field
		NetworkCountUsage                 resp.Field
		RamLimit                          resp.Field
		RamUsage                          resp.Field
		RegionID                          resp.Field
		RegistryCountLimit                resp.Field
		RegistryCountUsage                resp.Field
		RegistryStorageLimit              resp.Field
		RegistryStorageUsage              resp.Field
		RouterCountLimit                  resp.Field
		RouterCountUsage                  resp.Field
		SecretCountLimit                  resp.Field
		SecretCountUsage                  resp.Field
		ServergroupCountLimit             resp.Field
		ServergroupCountUsage             resp.Field
		SfsCountLimit                     resp.Field
		SfsCountUsage                     resp.Field
		SfsSizeLimit                      resp.Field
		SfsSizeUsage                      resp.Field
		SharedVmCountLimit                resp.Field
		SharedVmCountUsage                resp.Field
		SnapshotScheduleCountLimit        resp.Field
		SnapshotScheduleCountUsage        resp.Field
		SubnetCountLimit                  resp.Field
		SubnetCountUsage                  resp.Field
		VmCountLimit                      resp.Field
		VmCountUsage                      resp.Field
		VolumeCountLimit                  resp.Field
		VolumeCountUsage                  resp.Field
		VolumeSizeLimit                   resp.Field
		VolumeSizeUsage                   resp.Field
		VolumeSnapshotsCountLimit         resp.Field
		VolumeSnapshotsCountUsage         resp.Field
		VolumeSnapshotsSizeLimit          resp.Field
		VolumeSnapshotsSizeUsage          resp.Field
		ExtraFields                       map[string]resp.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaGetByRegionResponse) RawJSON() string { return r.JSON.raw }
func (r *QuotaGetByRegionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/paths/%2Fcloud%2Fv2%2Fglobal_quotas%2F%7Bclient_id%7D/get/responses/200/content/application%2Fjson/schema'
// "$.paths['/cloud/v2/global_quotas/{client_id}'].get.responses[200].content['application/json'].schema"
type QuotaGetGlobalResponse struct {
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_cpu_millicore_count_limit'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_cpu_millicore_count_limit"
	InferenceCPUMillicoreCountLimit int64 `json:"inference_cpu_millicore_count_limit"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_cpu_millicore_count_usage'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_cpu_millicore_count_usage"
	InferenceCPUMillicoreCountUsage int64 `json:"inference_cpu_millicore_count_usage"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_gpu_a100_count_limit'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_gpu_a100_count_limit"
	InferenceGPUA100CountLimit int64 `json:"inference_gpu_a100_count_limit"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_gpu_a100_count_usage'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_gpu_a100_count_usage"
	InferenceGPUA100CountUsage int64 `json:"inference_gpu_a100_count_usage"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_gpu_h100_count_limit'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_gpu_h100_count_limit"
	InferenceGPUH100CountLimit int64 `json:"inference_gpu_h100_count_limit"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_gpu_h100_count_usage'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_gpu_h100_count_usage"
	InferenceGPUH100CountUsage int64 `json:"inference_gpu_h100_count_usage"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_gpu_l40s_count_limit'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_gpu_l40s_count_limit"
	InferenceGPUL40sCountLimit int64 `json:"inference_gpu_l40s_count_limit"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_gpu_l40s_count_usage'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_gpu_l40s_count_usage"
	InferenceGPUL40sCountUsage int64 `json:"inference_gpu_l40s_count_usage"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_instance_count_limit'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_instance_count_limit"
	InferenceInstanceCountLimit int64 `json:"inference_instance_count_limit"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/inference_instance_count_usage'
	// "$.components.schemas.GlobalQuotasSerializer.properties.inference_instance_count_usage"
	InferenceInstanceCountUsage int64 `json:"inference_instance_count_usage"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/keypair_count_limit'
	// "$.components.schemas.GlobalQuotasSerializer.properties.keypair_count_limit"
	KeypairCountLimit int64 `json:"keypair_count_limit"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/keypair_count_usage'
	// "$.components.schemas.GlobalQuotasSerializer.properties.keypair_count_usage"
	KeypairCountUsage int64 `json:"keypair_count_usage"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/project_count_limit'
	// "$.components.schemas.GlobalQuotasSerializer.properties.project_count_limit"
	ProjectCountLimit int64 `json:"project_count_limit"`
	// '#/components/schemas/GlobalQuotasSerializer/properties/project_count_usage'
	// "$.components.schemas.GlobalQuotasSerializer.properties.project_count_usage"
	ProjectCountUsage int64 `json:"project_count_usage"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		InferenceCPUMillicoreCountLimit resp.Field
		InferenceCPUMillicoreCountUsage resp.Field
		InferenceGPUA100CountLimit      resp.Field
		InferenceGPUA100CountUsage      resp.Field
		InferenceGPUH100CountLimit      resp.Field
		InferenceGPUH100CountUsage      resp.Field
		InferenceGPUL40sCountLimit      resp.Field
		InferenceGPUL40sCountUsage      resp.Field
		InferenceInstanceCountLimit     resp.Field
		InferenceInstanceCountUsage     resp.Field
		KeypairCountLimit               resp.Field
		KeypairCountUsage               resp.Field
		ProjectCountLimit               resp.Field
		ProjectCountUsage               resp.Field
		ExtraFields                     map[string]resp.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaGetGlobalResponse) RawJSON() string { return r.JSON.raw }
func (r *QuotaGetGlobalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuotaGetByRegionParams struct {
	// '#/paths/%2Fcloud%2Fv2%2Fregional_quotas%2F%7Bclient_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v2/regional_quotas/{client_id}/{region_id}'].get.parameters[1].schema"
	//
	// Use [option.WithRegionID] on the client to set a global default for this field.
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f QuotaGetByRegionParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
