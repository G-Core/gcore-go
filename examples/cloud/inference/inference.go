package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/option"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	//apiKey := os.Getenv("GCORE_API_KEY")
	// Will use Production API URL if omitted
	//baseURL := os.Getenv("GCORE_BASE_URL")

	// TODO set cloud project ID before running
	cloudProjectID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_PROJECT_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_PROJECT_ID environment variable is required and must be a valid integer")
	}

	// TODO set cloud region ID before running
	cloudRegionID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_REGION_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_REGION_ID environment variable is required and must be a valid integer")
	}

	// TODO set flavor name before running
	cloudInferenceFlavorName := os.Getenv("GCORE_CLOUD_INFERENCE_FLAVOR_NAME")
	if cloudInferenceFlavorName == "" {
		log.Fatalf("GCORE_CLOUD_INFERENCE_FLAVOR_NAME environment variable is required")
	}

	// TODO set model ID before running
	cloudInferenceModelID := os.Getenv("GCORE_CLOUD_INFERENCE_MODEL_ID")
	if cloudInferenceModelID == "" {
		log.Fatalf("GCORE_CLOUD_INFERENCE_MODEL_ID environment variable is required")
	}

	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectID),
		option.WithCloudRegionID(cloudRegionID),
	)

	// Capacity
	getCapacityByRegion(&client)

	// Flavors
	listFlavors(&client)
	getFlavor(&client, cloudInferenceFlavorName)

	// Registry Credentials
	credentialName := createRegistryCredential(&client)
	listRegistryCredentials(&client)
	getRegistryCredential(&client, credentialName)
	replaceRegistryCredential(&client, credentialName)
	deleteRegistryCredential(&client, credentialName)

	// Secrets
	secretName := createSecret(&client)
	listSecrets(&client)
	getSecret(&client, secretName)
	replaceSecret(&client, secretName)
	deleteSecret(&client, secretName)

	// Deployments
	deploymentName := createDeployment(&client, cloudInferenceFlavorName, cloudRegionID)
	listDeployments(&client)
	getDeployment(&client, deploymentName)
	updateDeployment(&client, deploymentName)
	stopDeployment(&client, deploymentName)
	startDeployment(&client, deploymentName)
	listDeploymentLogs(&client, deploymentName)
	deleteDeployment(&client, deploymentName)
}

func getCapacityByRegion(client *gcore.Client) {
	fmt.Println("\n=== GET CAPACITY BY REGION ===")

	capacities, err := client.Cloud.Inference.GetCapacityByRegion(context.Background())
	if err != nil {
		fmt.Printf("Error getting capacity by region: %v\n", err)
		return
	}

	displayCount := 3
	if len(capacities.Results) < displayCount {
		displayCount = len(capacities.Results)
	}

	for i := 0; i < displayCount; i++ {
		regionCapacity := capacities.Results[i]
		fmt.Printf("%d. Region ID: %d, available flavors: %d\n", i+1, regionCapacity.RegionID, len(regionCapacity.Capacity))

		// Display first few flavors per region
		flavorDisplayCount := 3
		if len(regionCapacity.Capacity) < flavorDisplayCount {
			flavorDisplayCount = len(regionCapacity.Capacity)
		}

		for j := 0; j < flavorDisplayCount; j++ {
			flavor := regionCapacity.Capacity[j]
			fmt.Printf("   - %s: %d capacity\n", flavor.FlavorName, flavor.Capacity)
		}

		if len(regionCapacity.Capacity) > flavorDisplayCount {
			fmt.Printf("   ... and %d more flavors\n", len(regionCapacity.Capacity)-flavorDisplayCount)
		}
	}

	if len(capacities.Results) > displayCount {
		fmt.Printf("... and %d more regions\n", len(capacities.Results)-displayCount)
	}

	fmt.Println("========================")
}
