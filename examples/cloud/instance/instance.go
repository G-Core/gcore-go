package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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

	// TODO set image ID before running, or leave it empty to upload an Ubuntu image
	imageID := os.Getenv("GCORE_CLOUD_IMAGE_ID")

	// TODO set flavor name before running, or leave it empty to use the smallest flavor that can boot the volume
	flavorName := os.Getenv("GCORE_CLOUD_FLAVOR_NAME")

	// TODO set SSH key ID before running, or leave it empty to create the instance without a key
	sshKeyID := os.Getenv("GCORE_CLOUD_SSH_KEY_ID")

	// instances and ports take the region from their parameters. Volumes, images and metrics still take
	// it from the client
	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectID),
		option.WithCloudRegionID(cloudRegionID),
	)

	uploadedImage := imageID == ""
	if uploadedImage {
		imageID = uploadImage(&client)
	}

	// an instance boots from a volume and connects to networks through ports, so create them first
	bootVolume := createBootVolume(&client, imageID)
	port := createPort(&client, cloudRegionID, networkID, subnetID)
	if flavorName == "" {
		flavorName = pickFlavor(&client, cloudRegionID, bootVolume.ID)
	}

	instance := createInstance(&client, cloudRegionID, flavorName, bootVolume.ID, port.ID, sshKeyID)
	getInstance(&client, instance.ID)
	listInstances(&client, cloudRegionID)

	if resizeFlavorName := pickResizeFlavor(&client, instance.ID); resizeFlavorName != "" {
		resizeInstance(&client, instance.ID, resizeFlavorName)
	}

	getConsole(&client, instance.ID)
	listMetrics(&client, instance.ID)
	runAction(&client, instance.ID, cloud.InstanceActionParamsActionStop)
	runAction(&client, instance.ID, cloud.InstanceActionParamsActionStart)

	deleteInstance(&client, instance.ID, bootVolume.ID, port.ID)
	if uploadedImage {
		deleteImage(&client, imageID)
	}
}

func createBootVolume(client *gcore.Client, imageID string) *cloud.Volume {
	fmt.Println("\n=== CREATE BOOT VOLUME ===")

	params := cloud.VolumeNewParams{
		OfImage: &cloud.VolumeNewParamsBodyImage{
			ImageID:  imageID,
			Name:     "gcore-go-example-boot-volume",
			Size:     10,
			TypeName: "standard",
			Tags:     map[string]string{"name": "gcore-go-example"},
		},
	}

	volume, err := client.Cloud.Volumes.NewAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating boot volume: %v", err)
	}

	fmt.Printf("Created Volume: ID=%s, Name=%s, Size=%d GiB, Bootable=%t\n", volume.ID, volume.Name, volume.Size, volume.Bootable)
	fmt.Println("==========================")

	return volume
}

func createPort(client *gcore.Client, regionID int64, networkID, subnetID string) *cloud.Port {
	fmt.Println("\n=== CREATE PORT ===")

	// security groups belong to the port. Set SecurityGroupIDs here, or later with Ports.UpdateAndPoll
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

	fmt.Printf("Created Port: ID=%s, IPAddress=%s\n", port.ID, port.IPAssignments[0].IPAddress)
	fmt.Println("===================")

	return port
}

func createInstance(client *gcore.Client, regionID int64, flavorName, bootVolumeID, portID, sshKeyID string) *cloud.Instance {
	fmt.Println("\n=== CREATE INSTANCE ===")

	params := cloud.InstanceNewParams{
		RegionID:     regionID,
		Name:         gcore.String("gcore-go-example-instance"),
		FlavorName:   flavorName,
		BootVolumeID: bootVolumeID,
		// the order of the ports is the order of the network interfaces
		Interfaces: []cloud.InstanceNewParamsInterface{
			{PortID: portID},
		},
		Tags: map[string]string{"name": "gcore-go-example"},
	}
	if sshKeyID != "" {
		params.SSHKeyID = gcore.String(sshKeyID)
	}

	instance, err := client.Cloud.Instances.NewAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating instance: %v", err)
	}

	fmt.Println("Created Instance:")
	printInstanceDetails(*instance)
	fmt.Println("=======================")

	return instance
}

func getInstance(client *gcore.Client, instanceID string) {
	fmt.Println("\n=== GET INSTANCE ===")

	// the API finds the region of the instance from its ID
	instance, err := client.Cloud.Instances.Get(context.Background(), instanceID, cloud.InstanceGetParams{})
	if err != nil {
		log.Fatalf("Error getting instance: %v", err)
	}

	printInstanceDetails(*instance)
	fmt.Println("====================")
}

func listInstances(client *gcore.Client, regionID int64) {
	fmt.Println("\n=== LIST INSTANCES ===")

	count := 0

	// without RegionIDs the list holds the instances of all regions
	iter := client.Cloud.Instances.ListAutoPaging(context.Background(), cloud.InstanceListParams{
		RegionIDs: []int64{regionID},
		OrderBy:   cloud.InstanceListParamsOrderByCreatedAtDesc,
		Limit:     gcore.Int(10), // Process 10 items per page
	})

	for iter.Next() {
		instance := iter.Current()
		count++
		fmt.Printf("  %d. Instance: ID=%s, Name=%s, Status=%s, RegionID=%d\n",
			count, instance.ID, instance.Name, instance.Status, instance.RegionID)
	}

	if err := iter.Err(); err != nil {
		log.Fatalf("Error iterating instances: %v", err)
	}

	fmt.Println("======================")
}

func resizeInstance(client *gcore.Client, instanceID, flavorName string) {
	fmt.Println("\n=== RESIZE INSTANCE ===")

	// a resize restarts the instance. The same call can also rename the instance, change its ports and
	// data volumes, or move it to a placement group with PlacementGroupID
	params := cloud.InstanceUpdateParams{
		FlavorName: gcore.String(flavorName),
	}

	instance, err := client.Cloud.Instances.UpdateAndPoll(context.Background(), instanceID, params)
	if err != nil {
		log.Fatalf("Error resizing instance: %v", err)
	}

	fmt.Println("Resized Instance:")
	printInstanceDetails(*instance)
	fmt.Println("=======================")
}

func getConsole(client *gcore.Client, instanceID string) {
	fmt.Println("\n=== GET CONSOLE ===")

	console, err := client.Cloud.Instances.Console.Get(context.Background(), instanceID, cloud.InstanceConsoleGetParams{})
	if err != nil {
		log.Fatalf("Error getting console: %v", err)
	}

	fmt.Printf("Console: Protocol=%s, Type=%s, URL=%s\n", console.RemoteConsole.Protocol, console.RemoteConsole.Type, console.RemoteConsole.URL)
	fmt.Println("===================")
}

func runAction(client *gcore.Client, instanceID string, action cloud.InstanceActionParamsAction) {
	fmt.Printf("\n=== %s INSTANCE ===\n", strings.ToUpper(string(action)))

	params := cloud.InstanceActionParams{
		Action: action,
	}

	instance, err := client.Cloud.Instances.ActionAndPoll(context.Background(), instanceID, params)
	if err != nil {
		log.Fatalf("Error running action %s: %v", action, err)
	}

	fmt.Printf("Ran action %s: ID=%s, Status=%s\n", action, instance.ID, instance.Status)
	fmt.Println("========================")
}

func deleteInstance(client *gcore.Client, instanceID, bootVolumeID, portID string) {
	fmt.Println("\n=== DELETE INSTANCE ===")

	// by default the delete keeps the volumes, ports and floating IPs of the instance. Name the ones to
	// delete, or set AllVolumes, AllPorts and AllFloatingIPs
	params := cloud.InstanceDeleteParams{
		VolumeIDs: []string{bootVolumeID},
		PortIDs:   []string{portID},
	}

	if err := client.Cloud.Instances.DeleteAndPoll(context.Background(), instanceID, params); err != nil {
		log.Fatalf("Error deleting instance: %v", err)
	}

	fmt.Printf("Deleted Instance with its boot volume and port: ID=%s, VolumeID=%s, PortID=%s\n", instanceID, bootVolumeID, portID)
	fmt.Println("=======================")
}

func printInstanceDetails(instance cloud.Instance) {
	fmt.Printf("  ID=%s, Name=%s, Status=%s, RegionID=%d\n", instance.ID, instance.Name, instance.Status, instance.RegionID)
	fmt.Printf("  FlavorName=%s, BootVolumeID=%s\n", instance.FlavorName, instance.BootVolumeID)
	for _, iface := range instance.Interfaces {
		fmt.Printf("  Interface: PortID=%s, NetworkID=%s\n", iface.PortID, iface.NetworkID)
		for _, ipAssignment := range iface.IPAssignments {
			fmt.Printf("    IP assignment: IPAddress=%s, SubnetID=%s\n", ipAssignment.IPAddress, ipAssignment.SubnetID)
		}
	}
}
