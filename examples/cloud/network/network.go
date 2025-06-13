package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
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

	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectID),
		option.WithCloudRegionID(cloudRegionID),
	)

	networkID := createNetwork(&client)
	getNetwork(&client, networkID)
	updateNetwork(&client, networkID)
	listNetworks(&client)

	// Routers
	routerID := createRouter(&client)
	listAllRouters(&client)
	listRoutersWithFilters(&client)
	listRoutersWithAutopager(&client)
	getRouterByID(&client, routerID)
	updateRouter(&client, routerID)
	deleteRouter(&client, routerID)

	// Subnets
	subnetID := createSubnet(&client, networkID)
	getSubnet(&client, subnetID)
	updateSubnet(&client, subnetID)
	listSubnets(&client, networkID)
	deleteSubnet(&client, subnetID)

	deleteNetwork(&client, networkID)

}

func createNetwork(client *gcore.Client) string {
	fmt.Println("\n=== CREATE NETWORK ===")
	name := "gcore-go-example"
	params := cloud.NetworkNewParams{
		Name:         name,
		CreateRouter: gcore.Bool(true),
		Type:         cloud.NetworkNewParamsTypeVxlan,
		Tags:         map[string]string{"name": "gcore-go-example"},
	}

	taskIDList, err := client.Cloud.Networks.New(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating network: %v", err)
	}

	task, err := client.Cloud.Tasks.Poll(context.Background(), taskIDList.Tasks[0])
	if err != nil {
		log.Fatalf("Error polling task for network creation: %v", err)
	}

	if len(task.CreatedResources.Networks) == 0 {
		log.Fatalf("No network created")
	}

	networkID := task.CreatedResources.Networks[0]
	fmt.Printf("Created Network ID: %s\n", networkID)
	fmt.Println("====================")
	return networkID
}

func getNetwork(client *gcore.Client, networkID string) {
	fmt.Println("\n=== GET NETWORK ===")
	network, err := client.Cloud.Networks.Get(context.Background(), networkID, cloud.NetworkGetParams{})
	if err != nil {
		log.Fatalf("Error getting network: %v", err)
	}

	fmt.Printf("Network ID: %s, Name: %s, Type: %s\n", network.ID, network.Name, network.Type)
	fmt.Println("===================")
}

func updateNetwork(client *gcore.Client, networkID string) {
	fmt.Println("\n=== UPDATE NETWORK ===")
	newName := "gcore-go-example-updated"
	params := cloud.NetworkUpdateParams{
		Name: newName,
	}

	network, err := client.Cloud.Networks.Update(context.Background(), networkID, params)
	if err != nil {
		log.Fatalf("Error updating network: %v", err)
	}

	fmt.Printf("Updated network name to: %s\n", network.Name)
	fmt.Println("========================")
}

func listNetworks(client *gcore.Client) {
	fmt.Println("\n=== LIST NETWORKS ===")
	networks, err := client.Cloud.Networks.List(context.Background(), cloud.NetworkListParams{})
	if err != nil {
		log.Fatalf("Error listing networks: %v", err)
	}

	for i, network := range networks.Results {
		fmt.Printf("  %d. Network ID: %s, Name: %s, Type: %s\n", i+1, network.ID, network.Name, network.Type)
	}

	if len(networks.Results) == 0 {
		fmt.Println("  No networks found.")
	}
	fmt.Println("=====================")
}

func deleteNetwork(client *gcore.Client, networkID string) {
	fmt.Println("\n=== DELETE NETWORK ===")
	params := cloud.NetworkDeleteParams{}
	taskIDList, err := client.Cloud.Networks.Delete(context.Background(), networkID, params)
	if err != nil {
		log.Fatalf("Error deleting network: %v", err)
	}

	if len(taskIDList.Tasks) > 0 {
		_, err = client.Cloud.Tasks.Poll(context.Background(), taskIDList.Tasks[0])
		if err != nil {
			log.Fatalf("Error polling task for network deletion: %v", err)
		}
	}

	fmt.Printf("Network with ID %s successfully deleted\n", networkID)
	fmt.Println("====================")
}
