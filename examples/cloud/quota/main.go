package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/packages/param"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	//apiKey := os.Getenv("GCORE_API_KEY")
	// Will use Production API URL if omitted
	//baseURL := os.Getenv("GCORE_BASE_URL")

	// TODO set cloud project ID before running
	clientID, err := strconv.ParseInt(os.Getenv("GCORE_CLIENT_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLIENT_ID environment variable is required and must be a valid integer")
	}

	// TODO set cloud region ID before running
	regionID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_REGION_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_REGION_ID environment variable is required and must be a valid integer")
	}

	client := gcore.NewClient(
	//option.WithAPIKey(apiKey),
	//option.WithBaseURL(baseURL),
	)

	// Get combined quotas (global and regional)
	getCombinedQuotas(&client)
	getGlobalQuota(&client, clientID)
	getRegionQuota(&client, clientID, regionID)
}

func getCombinedQuotas(client *gcore.Client) {
	fmt.Println("\n=== GET COMBINED QUOTAS ===")

	quotas, err := client.Cloud.Quotas.GetAll(context.Background())
	if err != nil {
		log.Fatalf("Error fetching combined quotas: %v", err)
	}

	printGlobalQuotas(quotas.GlobalQuotas)
	printRegionalQuotas(quotas.RegionalQuotas)

	fmt.Println("===========================")
}

func getGlobalQuota(client *gcore.Client, clientID int64) {
	fmt.Println("\n=== GET GLOBAL QUOTA ===")

	quota, err := client.Cloud.Quotas.GetGlobal(context.Background(), clientID)
	if err != nil {
		log.Fatalf("Error fetching global quota: %v", err)
	}

	fmt.Printf("Global quota for client_id=%d:\n", clientID)
	globalQuota := cloud.QuotaGetAllResponseGlobalQuotas(*quota)
	printGlobalQuotas(globalQuota)

	fmt.Println("========================")
}

func getRegionQuota(client *gcore.Client, clientID int64, regionID int64) {
	fmt.Println("\n=== GET REGION QUOTA ===")

	params := cloud.QuotaGetByRegionParams{
		ClientID: clientID,
		RegionID: param.Opt[int64]{Value: regionID},
	}
	quota, err := client.Cloud.Quotas.GetByRegion(context.Background(), params)
	if err != nil {
		log.Fatalf("Error fetching regional quota: %v", err)
	}

	fmt.Printf("Regional quota for client_id=%d, region_id=%d:\n", clientID, regionID)
	printRegionalQuotasFromResponse([]cloud.QuotaGetByRegionResponse{*quota})

	fmt.Println("========================")
}

func printGlobalQuotas(globalQuotas cloud.QuotaGetAllResponseGlobalQuotas) {
	fmt.Println("\n--- Global Quotas ---")
	fmt.Printf("  inference_cpu_millicore_count_limit: %d (usage: %d)\n",
		globalQuotas.InferenceCPUMillicoreCountLimit, globalQuotas.InferenceCPUMillicoreCountUsage)
	fmt.Printf("  inference_gpu_a100_count_limit: %d (usage: %d)\n",
		globalQuotas.InferenceGPUA100CountLimit, globalQuotas.InferenceGPUA100CountUsage)
	fmt.Printf("  inference_gpu_h100_count_limit: %d (usage: %d)\n",
		globalQuotas.InferenceGPUH100CountLimit, globalQuotas.InferenceGPUH100CountUsage)
	fmt.Printf("  inference_gpu_l40s_count_limit: %d (usage: %d)\n",
		globalQuotas.InferenceGPUL40sCountLimit, globalQuotas.InferenceGPUL40sCountUsage)
	fmt.Printf("  inference_instance_count_limit: %d (usage: %d)\n",
		globalQuotas.InferenceInstanceCountLimit, globalQuotas.InferenceInstanceCountUsage)
	fmt.Printf("  keypair_count_limit: %d (usage: %d)\n",
		globalQuotas.KeypairCountLimit, globalQuotas.KeypairCountUsage)
	fmt.Printf("  project_count_limit: %d (usage: %d)\n",
		globalQuotas.ProjectCountLimit, globalQuotas.ProjectCountUsage)
}

func printRegionalQuotas(regionalQuotas []cloud.QuotaGetAllResponseRegionalQuota) {
	fmt.Println("\n--- Regional Quotas ---")
	for idx, region := range regionalQuotas {
		fmt.Printf("  Region #%d: region_id=%d\n", idx+1, region.RegionID)
		fmt.Printf("    cpu_count_limit: %d (usage: %d)\n", region.CPUCountLimit, region.CPUCountUsage)
		fmt.Printf("    external_ip_count_limit: %d (usage: %d)\n", region.ExternalIPCountLimit, region.ExternalIPCountUsage)
		fmt.Printf("    ram_limit: %d GiB (usage: %d GiB)\n", region.RamLimit, region.RamUsage)
		fmt.Printf("    vm_count_limit: %d (usage: %d)\n", region.VmCountLimit, region.VmCountUsage)
		fmt.Printf("    volume_count_limit: %d (usage: %d)\n", region.VolumeCountLimit, region.VolumeCountUsage)
		fmt.Printf("    volume_size_limit: %d GiB (usage: %d GiB)\n", region.VolumeSizeLimit, region.VolumeSizeUsage)
	}
}

func printRegionalQuotasFromResponse(regionalQuotas []cloud.QuotaGetByRegionResponse) {
	fmt.Println("\n--- Regional Quotas ---")
	for idx, region := range regionalQuotas {
		fmt.Printf("  Region #%d: region_id=%d\n", idx+1, region.RegionID)
		fmt.Printf("    cpu_count_limit: %d (usage: %d)\n", region.CPUCountLimit, region.CPUCountUsage)
		fmt.Printf("    external_ip_count_limit: %d (usage: %d)\n", region.ExternalIPCountLimit, region.ExternalIPCountUsage)
		fmt.Printf("    ram_limit: %d GiB (usage: %d GiB)\n", region.RamLimit, region.RamUsage)
		fmt.Printf("    vm_count_limit: %d (usage: %d)\n", region.VmCountLimit, region.VmCountUsage)
		fmt.Printf("    volume_count_limit: %d (usage: %d)\n", region.VolumeCountLimit, region.VolumeCountUsage)
		fmt.Printf("    volume_size_limit: %d GiB (usage: %d GiB)\n", region.VolumeSizeLimit, region.VolumeSizeUsage)
	}
}
