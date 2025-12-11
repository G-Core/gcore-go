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
	"github.com/G-Core/gcore-go/shared/constant"
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

	// Create a Reserved Fixed IP and use its details for other operations
	fixedIP := createReservedFixedIP(&client)
	listAllReservedFixedIPs(&client)
	listReservedFixedIPsWithAutopager(&client)
	getReservedFixedIPByID(&client, fixedIP.PortID)
	toggleReservedFixedIPVip(&client, fixedIP.PortID, true)
	listCandidatePorts(&client, fixedIP.PortID)
	listConnectedPorts(&client, fixedIP.PortID)
	// IsVIP needs to be false to delete the reserved fixed IP
	toggleReservedFixedIPVip(&client, fixedIP.PortID, false)
	deleteReservedFixedIP(&client, fixedIP.PortID)
}

func createReservedFixedIP(client *gcore.Client) *cloud.ReservedFixedIP {
	fmt.Println("\n=== CREATE RESERVED FIXED IP ===")

	params := cloud.ReservedFixedIPNewParams{
		OfExternal: &cloud.ReservedFixedIPNewParamsBodyExternal{
			IsVip:    param.NewOpt(false),
			IPFamily: cloud.InterfaceIPFamilyIpv4,
			Type:     constant.ValueOf[constant.External](),
		},
	}

	resp, err := client.Cloud.ReservedFixedIPs.NewAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating reserved fixed IP: %v", err)
	}

	fmt.Printf("Created Reserved Fixed IP: Name=%s, Status=%s, PortID=%s, IP=%s\n",
		resp.Name, resp.Status, resp.PortID, resp.FixedIPAddress)
	fmt.Println("=================================")

	return resp
}

func listAllReservedFixedIPs(client *gcore.Client) {
	fmt.Println("\n=== LIST ALL RESERVED FIXED IPS ===")

	reservedIPs, err := client.Cloud.ReservedFixedIPs.List(context.Background(), cloud.ReservedFixedIPListParams{})
	if err != nil {
		log.Fatalf("Error listing reserved fixed IPs: %v", err)
	}

	for i, ip := range reservedIPs.Results {
		fmt.Printf("  %d. Reserved Fixed IP: Name=%s, Status=%s, PortID=%s\n", i+1, ip.Name, ip.Status, ip.PortID)
	}

	fmt.Println("===================================")
}

func listReservedFixedIPsWithAutopager(client *gcore.Client) {
	fmt.Println("\n=== LIST RESERVED FIXED IPS USING AUTOPAGER ===")

	count := 0

	iter := client.Cloud.ReservedFixedIPs.ListAutoPaging(context.Background(), cloud.ReservedFixedIPListParams{
		Limit: gcore.Int(10), // Process 10 items per page
	})

	for iter.Next() {
		ip := iter.Current()
		count++
		fmt.Printf("  %d. Reserved Fixed IP: Name=%s, Status=%s, PortID=%s\n", count, ip.Name, ip.Status, ip.PortID)
	}

	if err := iter.Err(); err != nil {
		log.Fatalf("Error iterating reserved fixed IPs: %v", err)
	}

	fmt.Println("===============================================")
}

func getReservedFixedIPByID(client *gcore.Client, portID string) {
	fmt.Println("\n=== GET RESERVED FIXED IP BY ID ===")

	reservedIP, err := client.Cloud.ReservedFixedIPs.Get(context.Background(), portID, cloud.ReservedFixedIPGetParams{})
	if err != nil {
		log.Fatalf("Error getting reserved fixed IP: %v", err)
	}

	fmt.Printf("Reserved Fixed IP: Name=%s, Status=%s, PortID=%s, IsVip=%t\n", reservedIP.Name, reservedIP.Status,
		reservedIP.PortID, reservedIP.IsVip)
	fmt.Println("===================================")
}

func toggleReservedFixedIPVip(client *gcore.Client, portID string, isVip bool) {
	fmt.Println("\n=== TOGGLE RESERVED FIXED IP VIP ===")

	params := cloud.ReservedFixedIPVipToggleParams{
		IsVip: isVip,
	}
	reservedIP, err := client.Cloud.ReservedFixedIPs.Vip.Toggle(context.Background(), portID, params)
	if err != nil {
		log.Fatalf("Error toggling reserved fixed IP VIP: %v", err)
	}

	fmt.Printf("Toggled Reserved Fixed IP: Name=%s, Status=%s, PortID=%s, IsVip=%t\n", reservedIP.Name, reservedIP.Status,
		reservedIP.PortID, reservedIP.IsVip)
	fmt.Println("====================================")
}

func listCandidatePorts(client *gcore.Client, portID string) {
	fmt.Println("\n=== LIST CANDIDATE PORTS ===")

	candidatePorts, err := client.Cloud.ReservedFixedIPs.Vip.CandidatePorts.List(context.Background(), portID, cloud.ReservedFixedIPVipCandidatePortListParams{})
	if err != nil {
		log.Fatalf("Error listing candidate ports: %v", err)
	}

	for i, port := range candidatePorts.Results {
		fmt.Printf("  %d. Candidate Port: PortID=%s, InstanceName=%s\n", i+1, port.PortID, port.InstanceName)
	}

	fmt.Println("============================")
}

func listConnectedPorts(client *gcore.Client, portID string) {
	fmt.Println("\n=== LIST CONNECTED PORTS ===")

	connectedPorts, err := client.Cloud.ReservedFixedIPs.Vip.ConnectedPorts.List(context.Background(), portID, cloud.ReservedFixedIPVipConnectedPortListParams{})
	if err != nil {
		log.Fatalf("Error listing connected ports: %v", err)
	}

	for i, port := range connectedPorts.Results {
		fmt.Printf("  %d. Connected Port: PortID=%s, InstanceName=%s\n", i+1, port.PortID, port.InstanceName)
	}

	fmt.Println("============================")
}

func deleteReservedFixedIP(client *gcore.Client, portID string) {
	fmt.Println("\n=== DELETE RESERVED FIXED IP ===")

	err := client.Cloud.ReservedFixedIPs.DeleteAndPoll(context.Background(), portID, cloud.ReservedFixedIPDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting reserved fixed IP: %v", err)
	}

	fmt.Printf("Reserved Fixed IP with PortID %s successfully deleted\n", portID)
	fmt.Println("=================================")
}
