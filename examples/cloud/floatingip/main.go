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
	"github.com/G-Core/gcore-go/packages/param"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	//apiKey := os.Getenv("GCORE_API_KEY")
	// Will use Production API URL if omitted
	//baseURL := os.Getenv("GCORE_API_URL")

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

	// Create a Floating IP and use its details for other operations
	floatingIP := createFloatingIP(&client)
	listAllFloatingIPs(&client)
	getFloatingIPByID(&client, floatingIP.ID)
	assignFloatingIP(&client, floatingIP.ID, "f827a969-faa9-49c4-9d06-268ccfd1f256")
	unassignFloatingIP(&client, floatingIP.ID)
	deleteFloatingIP(&client, floatingIP.ID)
}

func createFloatingIP(client *gcore.Client) *cloud.FloatingIP {
	fmt.Println("\n=== CREATE FLOATING IP ===")

	params := cloud.FloatingIPNewParams{
		PortID:         param.Opt[string]{}, // Unassigned at creation, omitted field
		FixedIPAddress: param.Opt[string]{}, // Omitted field
		Tags:           map[string]string{"name": "gcore-go-example"},
	}

	taskIDs, err := client.Cloud.FloatingIPs.New(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating floating IP: %v", err)
	}

	task, err := client.Cloud.Tasks.Poll(context.Background(), taskIDs.Tasks[0])
	if err != nil {
		log.Fatalf("Error polling task for floating IP creation: %v", err)
	}

	floatingIP, err := client.Cloud.FloatingIPs.Get(context.Background(), task.CreatedResources.Floatingips[0], cloud.FloatingIPGetParams{})
	if err != nil {
		log.Fatalf("Error getting created floating IP details: %v", err)
	}

	fmt.Printf("Created Floating IP: ID=%s, Status=%s, FloatingIPAddress=%s\n",
		floatingIP.ID, floatingIP.Status, floatingIP.FloatingIPAddress)
	fmt.Println("===========================")

	return floatingIP
}

func listAllFloatingIPs(client *gcore.Client) {
	fmt.Println("\n=== LIST ALL FLOATING IPS ===")

	floatingIPs, err := client.Cloud.FloatingIPs.List(context.Background(), cloud.FloatingIPListParams{})
	if err != nil {
		log.Fatalf("Error listing floating IPs: %v", err)
	}

	for i, ip := range floatingIPs.Results {
		fmt.Printf("  %d. Floating IP: ID=%s, Status=%s, FloatingIPAddress=%s, PortID=%v\n",
			i+1, ip.ID, ip.Status, ip.FloatingIPAddress, ip.PortID)
	}

	if len(floatingIPs.Results) == 0 {
		fmt.Println("  No floating IPs found.")
	}

	fmt.Println("==============================")
}

func getFloatingIPByID(client *gcore.Client, floatingIPID string) {
	fmt.Println("\n=== GET FLOATING IP BY ID ===")

	floatingIP, err := client.Cloud.FloatingIPs.Get(context.Background(), floatingIPID, cloud.FloatingIPGetParams{})
	if err != nil {
		log.Fatalf("Error getting floating IP: %v", err)
	}

	fmt.Printf("Floating IP: ID=%s, Status=%s, FloatingIPAddress=%s, PortID=%v, FixedIPAddress=%v\n",
		floatingIP.ID, floatingIP.Status, floatingIP.FloatingIPAddress, floatingIP.PortID, floatingIP.FixedIPAddress)
	fmt.Println("==============================")
}

func assignFloatingIP(client *gcore.Client, floatingIPID string, portID string) {
	fmt.Println("\n=== ASSIGN FLOATING IP ===")

	params := cloud.FloatingIPAssignParams{
		PortID: portID,
	}

	floatingIP, err := client.Cloud.FloatingIPs.Assign(context.Background(), floatingIPID, params)
	if err != nil {
		log.Fatalf("Error assigning floating IP: %v", err)
	}

	fmt.Printf("Assigned Floating IP: ID=%s, PortID=%v, FixedIPAddress=%v\n",
		floatingIP.ID, floatingIP.PortID, floatingIP.FixedIPAddress)
	fmt.Println("===========================")
}

func unassignFloatingIP(client *gcore.Client, floatingIPID string) {
	fmt.Println("\n=== UNASSIGN FLOATING IP ===")

	floatingIP, err := client.Cloud.FloatingIPs.Unassign(context.Background(), floatingIPID, cloud.FloatingIPUnassignParams{})
	if err != nil {
		log.Fatalf("Error unassigning floating IP: %v", err)
	}

	fmt.Printf("Unassigned Floating IP: ID=%s\n", floatingIP.ID)
	fmt.Println("=============================")
}

func deleteFloatingIP(client *gcore.Client, floatingIPID string) {
	fmt.Println("\n=== DELETE FLOATING IP ===")

	taskIDs, err := client.Cloud.FloatingIPs.Delete(context.Background(), floatingIPID, cloud.FloatingIPDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting floating IP: %v", err)
	}

	_, err = client.Cloud.Tasks.Poll(context.Background(), taskIDs.Tasks[0])
	if err != nil {
		log.Fatalf("Error polling task for floating IP deletion: %v", err)
	}

	fmt.Printf("Floating IP with ID %s successfully deleted\n", floatingIPID)
	fmt.Println("===========================")
}
