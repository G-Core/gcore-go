package main

import (
	"context"
	"fmt"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/shared/constant"
	"log"
	"os"

	"github.com/stainless-sdks/gcore-go"
	"github.com/stainless-sdks/gcore-go/cloud"
	"github.com/stainless-sdks/gcore-go/option"
)

func main() {
	// Retrieve API key and base URL from environment variables
	apiKey := os.Getenv("GCORE_API_KEY")  // Automatically read by NewClient if omitted
	baseURL := os.Getenv("GCORE_API_URL") // Defaults to Production API URL if omitted

	// Initialize the client
	client := gcore.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL(baseURL),
	)

	// Execute example functions
	portID := createReservedFixedIP(&client)
	listReservedFixedIPs(&client)
	getReservedFixedIPByID(&client, portID)
	toggleReservedFixedIPVip(&client, portID, true)

	// operations on VIPs
	listCandidatePorts(&client, portID)
	listConnectedPorts(&client, portID)

	// IsVIP needs to be false to delete the reserved fixed IP
	toggleReservedFixedIPVip(&client, portID, false)
	deleteReservedFixedIP(&client, portID)

}

func createReservedFixedIP(client *gcore.Client) string {
	fmt.Println("\n=== CREATE RESERVED FIXED IP ===")

	ctx := context.Background()

	// Example parameters for creating a Reserved Fixed IP
	params := cloud.ReservedFixedIPNewParams{
		Body: cloud.ReservedFixedIPNewParamsBodyNewReservedFixedIPExternalSerializer{
			IsVip:    param.NewOpt(false),
			IPFamily: cloud.InterfaceIPFamilyIpv4,
			Type:     constant.ValueOf[constant.External](),
		},
	}

	resp, err := client.Cloud.ReservedFixedIPs.NewAndPoll(ctx, params)
	if err != nil {
		log.Fatalf("Failed to create reserved fixed IP: %v", err)
	}

	fmt.Printf("Created Reserved Fixed IP: Name=%s, Status=%s, PortID:%s\n", resp.Name, resp.Status, resp.PortID)
	fmt.Println("========================")
	return resp.PortID
}

func listReservedFixedIPs(client *gcore.Client) {
	fmt.Println("\n=== LIST RESERVED FIXED IPS ===")

	ctx := context.Background()

	// Simple listing
	params := cloud.ReservedFixedIPListParams{}
	reservedIPs, err := client.Cloud.ReservedFixedIPs.List(ctx, params)
	if err != nil {
		log.Fatalf("Failed to list reserved fixed IPs: %v", err)
	}

	for i, ip := range reservedIPs.Results {
		fmt.Printf("Reserved Fixed IP (%d): Name=%s, Status=%s, PortID=%s\n", i, ip.Name, ip.Status, ip.PortID)
	}

	// Pagination example
	fmt.Println("\n=== LIST RESERVED FIXED IPS WITH AUTOPAGER ===")
	iter := client.Cloud.ReservedFixedIPs.ListAutoPaging(ctx, params)
	counter := 0
	for iter.Next() {
		ip := iter.Current()
		fmt.Printf("Reserved Fixed IP (%d): Name=%s, Status=%s, PortID=%s\n", counter, ip.Name, ip.Status, ip.PortID)
		counter++
	}
	if err := iter.Err(); err != nil {
		log.Fatalf("Failed to paginate reserved fixed IPs: %v", err)
	}
	fmt.Printf("Total Reserved Fixed IPs: %d\n", counter)
	fmt.Println("========================")
}

func getReservedFixedIPByID(client *gcore.Client, portID string) {
	fmt.Println("\n=== GET RESERVED FIXED IP BY ID ===")

	ctx := context.Background()

	params := cloud.ReservedFixedIPGetParams{}
	reservedIP, err := client.Cloud.ReservedFixedIPs.Get(ctx, portID, params)
	if err != nil {
		log.Fatalf("Failed to get reserved fixed IP by ID: %v", err)
	}

	fmt.Printf("Reserved Fixed IP: Name=%s, Status=%s, PortID=%s, IsVip:%t\n", reservedIP.Name, reservedIP.Status,
		reservedIP.PortID, reservedIP.IsVip)
	fmt.Println("========================")
}

func deleteReservedFixedIP(client *gcore.Client, portID string) {
	fmt.Println("\n=== DELETE RESERVED FIXED IP ===")

	ctx := context.Background()
	params := cloud.ReservedFixedIPDeleteParams{}
	tasks, err := client.Cloud.ReservedFixedIPs.Delete(ctx, portID, params)
	if err != nil {
		log.Fatalf("Failed to delete reserved fixed IP: %v", err)
	}

	fmt.Printf("Deleted Reserved Fixed IP: Tasks=%v\n", tasks.Tasks)
	fmt.Println("========================")
}

func toggleReservedFixedIPVip(client *gcore.Client, portID string, isVip bool) {
	fmt.Println("\n=== TOGGLE RESERVED FIXED IP VIP ===")

	ctx := context.Background()

	params := cloud.ReservedFixedIPVipToggleParams{
		IsVip: isVip,
	}
	reservedIP, err := client.Cloud.ReservedFixedIPs.Vip.Toggle(ctx, portID, params)
	if err != nil {
		log.Fatalf("Failed to toggle reserved fixed IP: %v", err)
	}

	fmt.Printf("Reserved Fixed IP: Name=%s, Status=%s, PortID=%s, IsVip:%t\n", reservedIP.Name, reservedIP.Status,
		reservedIP.PortID, reservedIP.IsVip)
	fmt.Println("========================")
}

func listCandidatePorts(client *gcore.Client, portID string) {
	fmt.Println("\n=== LIST CANDIDATE PORTS ===")

	ctx := context.Background()

	// Example parameters for listing candidate ports
	params := cloud.ReservedFixedIPVipListCandidatePortsParams{}

	candidatePorts, err := client.Cloud.ReservedFixedIPs.Vip.ListCandidatePorts(ctx, portID, params)
	if err != nil {
		log.Fatalf("Failed to list candidate ports: %v", err)
	}

	for i, port := range candidatePorts.Results {
		fmt.Printf("Candidate Port (%d): PortID=%s, InstanceName=%s\n", i, port.PortID, port.InstanceName)
	}
	fmt.Println("========================")
}

func listConnectedPorts(client *gcore.Client, portID string) {
	fmt.Println("\n=== LIST CONNECTED PORTS ===")

	ctx := context.Background()

	// Example parameters for listing connected ports
	params := cloud.ReservedFixedIPVipListConnectedPortsParams{}

	connectedPorts, err := client.Cloud.ReservedFixedIPs.Vip.ListConnectedPorts(ctx, portID, params)
	if err != nil {
		log.Fatalf("Failed to list connected ports: %v", err)
	}

	for i, port := range connectedPorts.Results {
		fmt.Printf("Connected Port (%d): PortID=%s, InstanceName=%s\n", i, port.PortID, port.InstanceName)
	}
	fmt.Println("========================")
}
