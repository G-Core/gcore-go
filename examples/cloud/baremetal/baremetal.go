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

	// TODO set image ID before running, or leave it empty to use an Ubuntu image
	imageID := os.Getenv("GCORE_CLOUD_IMAGE_ID")

	// TODO set flavor name before running, or leave it empty to use the smallest flavor that can boot the image
	flavorName := os.Getenv("GCORE_CLOUD_FLAVOR_NAME")

	// TODO set the ID of a private VLAN network before running. The example attaches a second port in this
	// network. A server can have only one port in each network, so the first port uses the public network
	networkID := os.Getenv("GCORE_CLOUD_NETWORK_ID")
	if networkID == "" {
		log.Fatalf("GCORE_CLOUD_NETWORK_ID environment variable is required")
	}
	// TODO set a subnet of that network before running, or leave it empty to use any subnet of the network
	subnetID := os.Getenv("GCORE_CLOUD_SUBNET_ID")

	// TODO set security group ID before running, or leave it empty. The ports get it only when the flavor
	// supports security groups
	securityGroupID := os.Getenv("GCORE_CLOUD_SECURITY_GROUP_ID")

	// TODO set SSH key ID before running, or leave it empty to create the server without a key
	sshKeyID := os.Getenv("GCORE_CLOUD_SSH_KEY_ID")

	// bare metal servers, flavors and ports take the region from their parameters. Images and rebuild still
	// take it from the client
	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectID),
		option.WithCloudRegionID(cloudRegionID),
	)

	// Images
	if imageID == "" {
		images := listBaremetalImages(&client)
		listBaremetalImagesWithFilters(&client)
		imageID = getUbuntuImage(images)
	}

	// Flavors that can boot the image
	flavor := pickFlavor(&client, cloudRegionID, imageID, flavorName)

	ports := portSettings{
		regionID:        cloudRegionID,
		networkID:       networkID,
		subnetID:        subnetID,
		securityGroupID: securityGroupID,
		capabilities:    flavor.NetworkingCapabilities,
	}
	if err := runServerLifecycle(&client, ports, flavor, imageID, sshKeyID); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// runServerLifecycle creates a server from a new port, changes it and deletes it. It deletes the ports and the
// server that it creates, also when a step fails
func runServerLifecycle(client *gcore.Client, ports portSettings, flavor cloud.BaremetalFlavor, imageID, sshKeyID string) error {
	// a bare metal server connects to networks through ports, so create the ports first: one in the public
	// network for the create, and one in the private network to attach later. A server can have only one
	// port in each network
	public := ports
	public.networkID, public.subnetID = "", ""
	port, err := createPort(client, public, "gcore-go-example-port-1")
	if err != nil {
		return err
	}
	// the server delete keeps its ports. The deferred calls run in reverse order, so the ports go after the
	// server
	defer deletePort(client, port.ID)
	secondPort, err := createPort(client, ports, "gcore-go-example-port-2")
	if err != nil {
		return err
	}
	defer deletePort(client, secondPort.ID)

	server, err := createServer(client, ports.regionID, flavor.Name, imageID, port.ID, sshKeyID)
	if err != nil {
		return err
	}
	serverDeleted := false
	defer func() {
		if !serverDeleted {
			if err := deleteServer(client, server.ID); err != nil {
				fmt.Printf("Error deleting bare metal server %s: %v\n", server.ID, err)
			}
		}
	}()

	if err := listServers(client, ports.regionID); err != nil {
		return err
	}
	if err := getServer(client, server.ID); err != nil {
		return err
	}

	// the full list of interfaces: the current port first, then the new port at the end
	if err := updateServer(client, server.ID, []string{port.ID, secondPort.ID}); err != nil {
		return err
	}

	if err := getConsole(client, server.ID); err != nil {
		return err
	}
	if err := runAction(client, server.ID, cloud.BaremetalServerActionParamsActionReboot); err != nil {
		return err
	}
	if err := rebuildServer(client, server.ID, imageID); err != nil {
		return err
	}

	if err := deleteServer(client, server.ID); err != nil {
		return err
	}
	serverDeleted = true
	return nil
}

// portSettings holds the port settings that the selected flavor accepts
type portSettings struct {
	regionID        int64
	networkID       string
	subnetID        string
	securityGroupID string
	capabilities    cloud.BaremetalFlavorNetworkingCapabilities
}

func createPort(client *gcore.Client, settings portSettings, name string) (*cloud.Port, error) {
	fmt.Println("\n=== CREATE PORT ===")

	// with no network, the port gets one IPv4 address in the public network of the region
	params := cloud.PortNewParams{
		RegionID: settings.regionID,
		Tags:     map[string]string{"name": name},
	}
	if settings.networkID != "" {
		params.NetworkID = gcore.String(settings.networkID)
		if settings.subnetID != "" {
			params.IPAssignments = []cloud.PortNewParamsIPAssignment{{SubnetID: settings.subnetID}}
		}
	}
	// a flavor without security group support refuses a port with groups. An empty list keeps the default
	// security group off the port. This example sets no port security and no allowed address pairs
	switch {
	case !settings.capabilities.SecurityGroups:
		params.SecurityGroupIDs = []string{}
	case settings.securityGroupID != "":
		params.SecurityGroupIDs = []string{settings.securityGroupID}
	}

	port, err := client.Cloud.Ports.NewAndPoll(context.Background(), params)
	if err != nil {
		return nil, fmt.Errorf("create port: %w", err)
	}

	fmt.Printf("Created Port: ID=%s, NetworkID=%s\n", port.ID, port.NetworkID)
	for _, ipAssignment := range port.IPAssignments {
		fmt.Printf("  IP assignment: IPAddress=%s, SubnetID=%s\n", ipAssignment.IPAddress, ipAssignment.SubnetID)
	}
	fmt.Println("===================")

	return port, nil
}

func deletePort(client *gcore.Client, portID string) {
	fmt.Println("\n=== DELETE PORT ===")

	if err := client.Cloud.Ports.DeleteAndPoll(context.Background(), portID, cloud.PortDeleteParams{}); err != nil {
		fmt.Printf("Error deleting port %s: %v\n", portID, err)
		return
	}

	fmt.Printf("Deleted Port: ID=%s\n", portID)
	fmt.Println("===================")
}

func createServer(client *gcore.Client, regionID int64, flavorName, imageID, portID, sshKeyID string) (*cloud.BaremetalServer, error) {
	fmt.Println("\n=== CREATE BAREMETAL SERVER ===")

	params := cloud.BaremetalServerNewParams{
		RegionID:   regionID,
		Name:       gcore.String("gcore-go-example"),
		FlavorName: flavorName,
		ImageID:    imageID,
		// the order of the ports is the NIC order and the boot order. Without PortGroup, all ports go to the
		// first bond. A DPU flavor refuses PortGroup
		Interfaces: []cloud.BaremetalServerNewParamsInterface{
			{PortID: portID},
		},
		Tags: map[string]string{"name": "gcore-go-example"},
	}
	if sshKeyID != "" {
		params.SSHKeyID = gcore.String(sshKeyID)
	}

	server, err := client.Cloud.Baremetal.Servers.NewAndPoll(context.Background(), params)
	if err != nil {
		return nil, fmt.Errorf("create bare metal server: %w", err)
	}

	fmt.Println("Created baremetal server:")
	printServerDetails(*server)
	fmt.Println("===============================")
	return server, nil
}

func listServers(client *gcore.Client, regionID int64) error {
	fmt.Println("\n=== LIST BAREMETAL SERVERS ===")

	count := 0

	// without RegionIDs the list holds the servers of all regions
	iter := client.Cloud.Baremetal.Servers.ListAutoPaging(context.Background(), cloud.BaremetalServerListParams{
		RegionIDs: []int64{regionID},
		Name:      cloud.BaremetalServerListParamsName{Prefix: []string{"gcore-go-example"}},
		OrderBy:   cloud.BaremetalServerListParamsOrderByCreatedAtDesc,
		Limit:     gcore.Int(10), // Process 10 items per page
	})

	for iter.Next() {
		server := iter.Current()
		count++
		fmt.Printf("  %d. Server: ID=%s, Name=%s, Status=%s, RegionID=%d\n",
			count, server.ID, server.Name, server.Status, server.RegionID)
	}

	if err := iter.Err(); err != nil {
		return fmt.Errorf("list bare metal servers: %w", err)
	}

	fmt.Printf("Total baremetal servers: %d\n", count)
	fmt.Println("===============================")
	return nil
}

func getServer(client *gcore.Client, serverID string) error {
	fmt.Println("\n=== GET BAREMETAL SERVER ===")

	// the API finds the region of the server from its ID
	server, err := client.Cloud.Baremetal.Servers.Get(context.Background(), serverID, cloud.BaremetalServerGetParams{})
	if err != nil {
		return fmt.Errorf("get bare metal server: %w", err)
	}

	printServerDetails(*server)
	fmt.Println("============================")
	return nil
}

func updateServer(client *gcore.Client, serverID string, portIDs []string) error {
	fmt.Println("\n=== UPDATE BAREMETAL SERVER ===")

	interfaces := make([]cloud.BaremetalServerUpdateParamsInterface, 0, len(portIDs))
	for _, portID := range portIDs {
		interfaces = append(interfaces, cloud.BaremetalServerUpdateParamsInterface{PortID: portID})
	}
	// tags use JSON Merge Patch: a nil value removes the tag
	params := cloud.BaremetalServerUpdateParams{
		Interfaces: interfaces,
		Tags: cloud.TagUpdateMap{
			"env":  gcore.Ptr("example"),
			"name": nil,
		},
	}

	server, err := client.Cloud.Baremetal.Servers.UpdateAndPoll(context.Background(), serverID, params)
	if err != nil {
		return fmt.Errorf("update bare metal server: %w", err)
	}

	// the server can show old data for a few seconds after the tasks
	fmt.Println("Updated baremetal server:")
	printServerDetails(*server)
	fmt.Println("===============================")
	return nil
}

func getConsole(client *gcore.Client, serverID string) error {
	fmt.Println("\n=== GET CONSOLE ===")

	// without ConsoleType the hardware picks the console type
	console, err := client.Cloud.Baremetal.Servers.Console.Get(context.Background(), serverID, cloud.BaremetalServerConsoleGetParams{})
	if err != nil {
		return fmt.Errorf("get console: %w", err)
	}

	fmt.Printf("Console: Protocol=%s, Type=%s, URL=%s\n", console.RemoteConsole.Protocol, console.RemoteConsole.Type, console.RemoteConsole.URL)
	fmt.Println("===================")
	return nil
}

func runAction(client *gcore.Client, serverID string, action cloud.BaremetalServerActionParamsAction) error {
	fmt.Printf("\n=== %s BAREMETAL SERVER ===\n", action)

	server, err := client.Cloud.Baremetal.Servers.ActionAndPoll(context.Background(), serverID, cloud.BaremetalServerActionParams{
		Action: action,
	})
	if err != nil {
		return fmt.Errorf("run action %s: %w", action, err)
	}

	fmt.Printf("Ran action %s: ID=%s, Status=%s\n", action, server.ID, server.Status)
	fmt.Println("===============================")
	return nil
}

func rebuildServer(client *gcore.Client, serverID string, imageID string) error {
	fmt.Println("\n=== REBUILD BAREMETAL SERVER ===")

	// rebuild keeps the regional route, so it takes the region from the client. It returns the regional
	// server model
	params := cloud.BaremetalServerRebuildParams{
		ImageID: gcore.String(imageID),
	}

	server, err := client.Cloud.Baremetal.Servers.RebuildAndPoll(context.Background(), serverID, params)
	if err != nil {
		return fmt.Errorf("rebuild bare metal server: %w", err)
	}

	fmt.Printf("Rebuilt baremetal server: ID=%s, Name=%s, Status=%s\n", server.ID, server.Name, server.Status)
	fmt.Println("===============================")
	return nil
}

func deleteServer(client *gcore.Client, serverID string) error {
	fmt.Println("\n=== DELETE BAREMETAL SERVER ===")

	// by default the delete keeps the ports and floating IPs of the server. Name the ones to delete in
	// PortIDs and FloatingIPIDs, or set AllPorts and AllFloatingIPs
	err := client.Cloud.Baremetal.Servers.DeleteAndPoll(context.Background(), serverID, cloud.BaremetalServerDeleteParams{})
	if err != nil {
		return fmt.Errorf("delete bare metal server: %w", err)
	}

	fmt.Printf("Deleted baremetal server: ID=%s\n", serverID)
	fmt.Println("===============================")
	return nil
}

func printServerDetails(server cloud.BaremetalServer) {
	fmt.Printf("  ID=%s, Name=%s, Status=%s, RegionID=%d\n", server.ID, server.Name, server.Status, server.RegionID)
	fmt.Printf("  FlavorName=%s, ImageID=%s, Tags=%v\n", server.FlavorName, server.ImageID, server.Tags)
	for _, iface := range server.Interfaces {
		// a DPU server has no bond, so PortGroup is null
		portGroup := "none"
		if iface.JSON.PortGroup.Valid() {
			portGroup = strconv.FormatInt(iface.PortGroup, 10)
		}
		fmt.Printf("  Interface: PortID=%s, NetworkID=%s, PortGroup=%s\n", iface.PortID, iface.NetworkID, portGroup)
		for _, ipAssignment := range iface.IPAssignments {
			fmt.Printf("    IP assignment: IPAddress=%s, SubnetID=%s\n", ipAssignment.IPAddress, ipAssignment.SubnetID)
		}
	}
}
