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

	// TODO set cloud port ID before running
	portID := os.Getenv("GCORE_CLOUD_PORT_ID")
	if portID == "" {
		log.Fatalf("GCORE_CLOUD_PORT_ID environment variable is required")
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
	updateFloatingIP(&client, floatingIP.ID)
	assignFloatingIP(&client, floatingIP.ID, portID)
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

	floatingIP, err := client.Cloud.FloatingIPs.NewAndPoll(context.Background(), params)
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

	err := client.Cloud.FloatingIPs.DeleteAndPoll(context.Background(), floatingIPID, cloud.FloatingIPDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting floating IP: %v", err)
	}

	fmt.Printf("Floating IP with ID %s successfully deleted\n", floatingIPID)
	fmt.Println("===========================")
}

func updateFloatingIP(client *gcore.Client, floatingIPID string) {
	fmt.Println("\n=== UPDATE FLOATING IP ===")

	params := cloud.FloatingIPUpdateParams{
		Tags: map[string]string{"env": "prod"},
	}
	floatingIP, err := client.Cloud.FloatingIPs.Update(context.Background(), floatingIPID, params)
	if err != nil {
		log.Fatalf("Error updating floating IP: %v", err)
	}

	fmt.Printf("Updated Floating IP: ID=%s, Status=%s, FloatingIPAddress=%s\n",
		floatingIP.ID, floatingIP.Status, floatingIP.FloatingIPAddress)
	fmt.Println("===========================")
}
