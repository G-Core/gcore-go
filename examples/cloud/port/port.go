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

	// TODO set cloud network ID before running
	networkID := os.Getenv("GCORE_CLOUD_NETWORK_ID")
	if networkID == "" {
		log.Fatalf("GCORE_CLOUD_NETWORK_ID environment variable is required")
	}

	// TODO set cloud subnet ID of the network before running
	subnetID := os.Getenv("GCORE_CLOUD_SUBNET_ID")
	if subnetID == "" {
		log.Fatalf("GCORE_CLOUD_SUBNET_ID environment variable is required")
	}

	// TODO set cloud security group ID before running
	securityGroupID := os.Getenv("GCORE_CLOUD_SECURITY_GROUP_ID")
	if securityGroupID == "" {
		log.Fatalf("GCORE_CLOUD_SECURITY_GROUP_ID environment variable is required")
	}

	// ports use the global view: the requests take the region from their parameters, not from the client
	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectID),
	)

	// Create a Port and use its details for other operations
	port := createPort(&client, cloudRegionID, networkID, subnetID)
	// the second port owns the shared address that the first port holds as an allowed address pair
	vipPort := createPort(&client, cloudRegionID, networkID, subnetID)

	getPort(&client, port.ID)
	updatePortTag(&client, port.ID, "env", "example")
	updatePortSecurityGroups(&client, port.ID, []string{securityGroupID})
	disablePortSecurity(&client, port.ID)
	enablePortSecurity(&client, port.ID, []string{securityGroupID})
	updatePortAllowedAddressPairs(&client, port.ID, vipPort.IPAssignments[0].IPAddress)
	listPortsWithFilters(&client, cloudRegionID, networkID, "env", "example")

	// delete the first port before the VIP port, because its allowed address pairs refer to the VIP address
	deletePort(&client, port.ID)
	deletePort(&client, vipPort.ID)

	convertPortToFloatingIP(&client, cloudRegionID)
}

func createPort(client *gcore.Client, regionID int64, networkID, subnetID string) *cloud.Port {
	fmt.Println("\n=== CREATE PORT ===")

	params := cloud.PortNewParams{
		RegionID:  regionID,
		NetworkID: gcore.String(networkID),
		IPAssignments: []cloud.PortNewParamsIPAssignment{
			{SubnetID: subnetID},
		},
		Tags: map[string]string{"name": "gcore-go-example"},
	}

	port, err := client.Cloud.Ports.NewAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating port: %v", err)
	}

	fmt.Println("Created Port:")
	printPortDetails(*port)
	fmt.Println("===================")

	return port
}

func getPort(client *gcore.Client, portID string) {
	fmt.Println("\n=== GET PORT ===")

	port, err := client.Cloud.Ports.Get(context.Background(), portID, cloud.PortGetParams{})
	if err != nil {
		log.Fatalf("Error getting port: %v", err)
	}

	fmt.Println("Port:")
	printPortDetails(*port)
	fmt.Println("================")
}

func updatePortTag(client *gcore.Client, portID, key, value string) {
	fmt.Println("\n=== UPDATE PORT TAGS ===")

	// the sent tags merge into the existing tags. A tag with a nil value removes the tag
	params := cloud.PortUpdateParams{
		Tags: cloud.TagUpdateMap{key: gcore.StringPtr(value)},
	}

	port, err := client.Cloud.Ports.UpdateAndPoll(context.Background(), portID, params)
	if err != nil {
		log.Fatalf("Error updating port tags: %v", err)
	}
	// UpdateAndPoll waits until the port shows the new tag, so the returned port has it
	if got := port.Tags[key]; got != value {
		log.Fatalf("Port tag %s is %q, want %q", key, got, value)
	}

	fmt.Println("Updated Port:")
	printPortDetails(*port)
	fmt.Println("========================")
}

func updatePortSecurityGroups(client *gcore.Client, portID string, securityGroupIDs []string) {
	fmt.Println("\n=== UPDATE PORT SECURITY GROUPS ===")

	// the list replaces all security groups of the port
	params := cloud.PortUpdateParams{
		SecurityGroupIDs: securityGroupIDs,
	}

	port, err := client.Cloud.Ports.UpdateAndPoll(context.Background(), portID, params)
	if err != nil {
		log.Fatalf("Error updating port security groups: %v", err)
	}

	fmt.Println("Updated Port:")
	printPortDetails(*port)
	fmt.Println("===================================")
}

func disablePortSecurity(client *gcore.Client, portID string) {
	fmt.Println("\n=== DISABLE PORT SECURITY ===")

	// a port without port security must have no security groups. An empty slice sends `[]` and detaches
	// all groups, while a nil slice leaves them unchanged
	params := cloud.PortUpdateParams{
		PortSecurityEnabled: gcore.Bool(false),
		SecurityGroupIDs:    []string{},
	}

	port, err := client.Cloud.Ports.UpdateAndPoll(context.Background(), portID, params)
	if err != nil {
		log.Fatalf("Error disabling port security: %v", err)
	}

	fmt.Println("Updated Port:")
	printPortDetails(*port)
	fmt.Println("=============================")
}

func enablePortSecurity(client *gcore.Client, portID string, securityGroupIDs []string) {
	fmt.Println("\n=== ENABLE PORT SECURITY ===")

	// with port security and no security groups the port permits no traffic, so attach the groups again
	params := cloud.PortUpdateParams{
		PortSecurityEnabled: gcore.Bool(true),
		SecurityGroupIDs:    securityGroupIDs,
	}

	port, err := client.Cloud.Ports.UpdateAndPoll(context.Background(), portID, params)
	if err != nil {
		log.Fatalf("Error enabling port security: %v", err)
	}

	fmt.Println("Updated Port:")
	printPortDetails(*port)
	fmt.Println("============================")
}

func updatePortAllowedAddressPairs(client *gcore.Client, portID string, ipAddress string) {
	fmt.Println("\n=== UPDATE PORT ALLOWED ADDRESS PAIRS ===")

	// allowed address pairs need port security on the port
	params := cloud.PortUpdateParams{
		AllowedAddressPairs: []cloud.PortUpdateParamsAllowedAddressPair{
			{IPAddress: ipAddress},
		},
	}

	port, err := client.Cloud.Ports.UpdateAndPoll(context.Background(), portID, params)
	if err != nil {
		log.Fatalf("Error updating port allowed address pairs: %v", err)
	}

	fmt.Println("Updated Port:")
	printPortDetails(*port)
	fmt.Println("=========================================")
}

func listPortsWithFilters(client *gcore.Client, regionID int64, networkID, tagKey, tagValue string) {
	fmt.Println("\n=== LIST PORTS WITH FILTERS ===")

	count := 0

	// without RegionIDs the list holds the ports of all regions. A tag key can list several values, and
	// the port must have one of them
	iter := client.Cloud.Ports.ListAutoPaging(context.Background(), cloud.PortListParams{
		NetworkIDs: []string{networkID},
		RegionIDs:  []int64{regionID},
		Tags:       map[string][]string{tagKey: {tagValue}},
		Limit:      gcore.Int(10), // Process 10 items per page
	})

	for iter.Next() {
		port := iter.Current()
		count++
		fmt.Printf("  %d. Port: ID=%s, RegionID=%d, NetworkID=%s, Tags=%v\n",
			count, port.ID, port.RegionID, port.NetworkID, port.Tags)
	}

	if err := iter.Err(); err != nil {
		log.Fatalf("Error iterating ports: %v", err)
	}

	fmt.Println("===============================")
}

func deletePort(client *gcore.Client, portID string) {
	fmt.Println("\n=== DELETE PORT ===")

	err := client.Cloud.Ports.DeleteAndPoll(context.Background(), portID, cloud.PortDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting port: %v", err)
	}

	fmt.Printf("Port with ID %s successfully deleted\n", portID)
	fmt.Println("===================")
}

func convertPortToFloatingIP(client *gcore.Client, regionID int64) {
	fmt.Println("\n=== CONVERT PORT TO FLOATING IP ===")

	// a port without a network gets one public IPv4 address. Only an unattached port with one public
	// IPv4 address can become a floating IP
	port, err := client.Cloud.Ports.NewAndPoll(context.Background(), cloud.PortNewParams{
		RegionID: regionID,
		Tags:     map[string]string{"name": "gcore-go-example"},
	})
	if err != nil {
		log.Fatalf("Error creating public port: %v", err)
	}
	fmt.Println("Created public Port:")
	printPortDetails(*port)

	// the task deletes the port and creates a floating IP with the same address
	floatingIP, err := client.Cloud.Ports.ConvertAndPoll(context.Background(), port.ID, cloud.PortConvertParams{
		TargetType: cloud.PortConvertParamsTargetTypeFloatingIP,
	})
	if err != nil {
		log.Fatalf("Error converting port to floating IP: %v", err)
	}
	fmt.Printf("Converted Port %s to Floating IP: ID=%s, Address=%s, RegionID=%d\n",
		port.ID, floatingIP.ID, floatingIP.FloatingIPAddress, floatingIP.RegionID)

	err = client.Cloud.FloatingIPs.DeleteAndPoll(context.Background(), floatingIP.ID, cloud.FloatingIPDeleteParams{
		RegionID: gcore.Int(floatingIP.RegionID),
	})
	if err != nil {
		log.Fatalf("Error deleting floating IP: %v", err)
	}
	fmt.Printf("Floating IP with ID %s successfully deleted\n", floatingIP.ID)
	fmt.Println("===================================")
}

func printPortDetails(port cloud.Port) {
	fmt.Printf("  ID=%s, RegionID=%d, NetworkID=%s\n", port.ID, port.RegionID, port.NetworkID)
	for _, ipAssignment := range port.IPAssignments {
		fmt.Printf("  IP assignment: IPAddress=%s, SubnetID=%s\n", ipAssignment.IPAddress, ipAssignment.SubnetID)
	}
	fmt.Printf("  SecurityGroupIDs=%v\n", port.SecurityGroupIDs)
	fmt.Printf("  PortSecurityEnabled=%t\n", port.PortSecurityEnabled)
	for _, pair := range port.AllowedAddressPairs {
		fmt.Printf("  Allowed address pair: IPAddress=%s\n", pair.IPAddress)
	}
	fmt.Printf("  Tags=%v\n", port.Tags)
}
