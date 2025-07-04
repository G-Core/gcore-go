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

	// TODO set placement group ID before running
	placementGroupID := os.Getenv("GCORE_CLOUD_PLACEMENT_GROUP_ID")

	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectID),
		option.WithCloudRegionID(cloudRegionID),
	)

	uploadedImageID := uploadImage(&client)

	instanceID := createInstance(&client, uploadedImageID)
	instance := getInstance(&client, instanceID)
	getConsole(&client, instanceID)
	listInstances(&client)
	listInstancesWithAutopager(&client)
	updateInstance(&client, instanceID)
	rebootInstance(&client, instanceID)
	resizeInstance(&client, instanceID)

	// Flavors
	listFlavors(&client)

	// Images
	listImages(&client)
	getImage(&client, uploadedImageID)
	updateImage(&client, uploadedImageID)
	deleteImage(&client, uploadedImageID)
	//volumeImageID := createImageFromVolume(&client, instance.Volumes[0].ID)
	//deleteImage(&client, volumeImageID)

	// Interfaces
	interfaces := listInterfaces(&client, instanceID)
	detachInterface(&client, instanceID, interfaces[0].IPAssignments[0].IPAddress, interfaces[0].PortID)
	interfaces = attachInterface(&client, instanceID, interfaces[0].NetworkID)
	//disablePortSecurity(&client, interfaces[0].PortID)
	//enablePortSecurity(&client, interfaces[0].PortID)

	// Metrics
	listMetrics(&client, instanceID)

	// Placement groups
	if placementGroupID != "" {
		addToPlacementGroup(&client, instanceID, placementGroupID)
		removeFromPlacementGroup(&client, instanceID)
	}

	// Security groups
	unassignSecurityGroup(&client, instanceID)
	assignSecurityGroup(&client, instanceID)

	deleteInstance(&client, instanceID, instance.Volumes)
}

func createInstance(client *gcore.Client, imageID string) string {
	fmt.Println("\n=== CREATE INSTANCE ===")

	interfaces := []cloud.InstanceNewParamsInterfaceUnion{
		{OfExternal: &cloud.InstanceNewParamsInterfaceExternal{
			Type: constant.ValueOf[constant.External](),
		}},
	}

	volumes := []cloud.InstanceNewParamsVolumeUnion{
		{OfImage: &cloud.InstanceNewParamsVolumeImage{
			Name:      gcore.String("gcore-go-example-volume"),
			Size:      gcore.Int(10),
			TypeName:  "standard",
			Source:    constant.ValueOf[constant.Image](),
			ImageID:   imageID,
			BootIndex: gcore.Int(0),
		}},
	}

	params := cloud.InstanceNewParams{
		Name:       gcore.String("gcore-go-example-instance"),
		Flavor:     "g1-standard-1-2",
		Interfaces: interfaces,
		Volumes:    volumes,
		Password:   gcore.String("Gcore123!"),
		Tags:       map[string]string{"name": "gcore-go-example"},
	}

	instance, err := client.Cloud.Instances.NewAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating instance: %v", err)
	}

	fmt.Printf("Created Instance: ID=%s, Name=%s, Status=%s\n", instance.ID, instance.Name, instance.Status)
	fmt.Println("========================")

	return instance.ID
}

func getInstance(client *gcore.Client, instanceID string) *cloud.Instance {
	fmt.Println("\n=== GET INSTANCE ===")

	instance, err := client.Cloud.Instances.Get(context.Background(), instanceID, cloud.InstanceGetParams{})
	if err != nil {
		fmt.Printf("Error getting instance: %v\n", err)
		return nil
	}

	fmt.Printf("Instance: ID=%s, Name=%s, Status=%s\n", instance.ID, instance.Name, instance.Status)
	fmt.Println("====================")
	return instance
}

func listInstances(client *gcore.Client) {
	fmt.Println("\n=== LIST INSTANCES ===")

	params := cloud.InstanceListParams{}
	instances, err := client.Cloud.Instances.List(context.Background(), params)
	if err != nil {
		fmt.Printf("Error listing instances: %v\n", err)
		return
	}

	for i, instance := range instances.Results {
		fmt.Printf("  %d. Instance: ID=%s, Name=%s, Status=%s\n", i+1, instance.ID, instance.Name, instance.Status)
	}

	fmt.Println("======================")
}

func listInstancesWithAutopager(client *gcore.Client) {
	fmt.Println("\n=== LIST INSTANCES USING AUTOPAGER ===")

	count := 0

	iter := client.Cloud.Instances.ListAutoPaging(context.Background(), cloud.InstanceListParams{
		Limit: gcore.Int(10), // Process 10 items per page
	})

	for iter.Next() {
		instance := iter.Current()
		count++
		fmt.Printf("  %d. Instance: ID=%s, Name=%s, Status=%s\n",
			count, instance.ID, instance.Name, instance.Status)
	}

	if err := iter.Err(); err != nil {
		fmt.Printf("Error iterating instances: %v\n", err)
		return
	}

	fmt.Println("=======================================")
}

func updateInstance(client *gcore.Client, instanceID string) {
	fmt.Println("\n=== UPDATE INSTANCE ===")

	params := cloud.InstanceUpdateParams{
		Name: "gcore-go-example-updated",
	}
	instance, err := client.Cloud.Instances.Update(context.Background(), instanceID, params)
	if err != nil {
		fmt.Printf("Error updating instance: %v\n", err)
		return
	}

	fmt.Printf("Instance updated: ID=%s, Name changed to %s\n", instance.ID, instance.Name)
	fmt.Println("=======================")
}

func resizeInstance(client *gcore.Client, instanceID string) {
	fmt.Println("\n=== RESIZE INSTANCE ===")

	params := cloud.InstanceResizeParams{
		FlavorID: "g1-standard-2-4",
	}

	_, err := client.Cloud.Instances.ResizeAndPoll(context.Background(), instanceID, params)
	if err != nil {
		fmt.Printf("Error resizing instance: %v\n", err)
		return
	}

	fmt.Printf("Instance resized: ID=%s, New flavor=%s\n", instanceID, params.FlavorID)
	fmt.Println("=======================")
}

func rebootInstance(client *gcore.Client, instanceID string) {
	fmt.Println("\n=== REBOOT INSTANCE ===")

	// Reboot instance (works from active state)
	rebootParams := cloud.InstanceActionParams{
		OfBasicActionInstanceSerializer: &cloud.InstanceActionParamsBodyBasicActionInstanceSerializer{
			Action: "reboot",
		},
	}
	if _, err := client.Cloud.Instances.ActionAndPoll(context.Background(), instanceID, rebootParams); err != nil {
		fmt.Printf("Error rebooting instance: %v\n", err)
		return
	}
	fmt.Printf("Rebooted instance: %s\n", instanceID)

	fmt.Println("=======================")
}

func assignSecurityGroup(client *gcore.Client, instanceID string) {
	fmt.Println("\n=== ASSIGN SECURITY GROUP ===")

	securityGroupParams := cloud.InstanceAssignSecurityGroupParams{Name: gcore.String("default")}
	if err := client.Cloud.Instances.AssignSecurityGroup(context.Background(), instanceID, securityGroupParams); err != nil {
		fmt.Printf("Error assigning security group: %v\n", err)
		return
	}

	fmt.Printf("Assigned security group: %s\n", securityGroupParams.Name)
	fmt.Println("==============================")
}

func unassignSecurityGroup(client *gcore.Client, instanceID string) {
	fmt.Println("\n=== UNASSIGN SECURITY GROUP ===")

	securityGroupParams := cloud.InstanceUnassignSecurityGroupParams{Name: gcore.String("default")}
	if err := client.Cloud.Instances.UnassignSecurityGroup(context.Background(), instanceID, securityGroupParams); err != nil {
		fmt.Printf("Error unassigning security group: %v\n", err)
		return
	}

	fmt.Printf("Unassigned security group: %s\n", securityGroupParams.Name)
	fmt.Println("================================")
}

func deleteInstance(client *gcore.Client, instanceID string, volumes []cloud.InstanceVolume) {
	fmt.Println("\n=== DELETE INSTANCE ===")

	// Create comma-separated string from volume IDs
	volumesStr := ""
	if len(volumes) > 0 {
		volumesStr = volumes[0].ID
		for i := 1; i < len(volumes); i++ {
			volumesStr += "," + volumes[i].ID
		}
	}

	params := cloud.InstanceDeleteParams{
		DeleteFloatings: gcore.Bool(true),
		Volumes:         gcore.String(volumesStr),
	}
	if err := client.Cloud.Instances.DeleteAndPoll(context.Background(), instanceID, params); err != nil {
		log.Fatalf("Error deleting instance: %v", err)
	}

	fmt.Printf("Deleted Instance and related resources: ID=%s, Volumes=%s\n", instanceID, volumesStr)
	fmt.Println("=======================")
}

func getConsole(client *gcore.Client, instanceID string) {
	fmt.Println("\n=== GET CONSOLE ===")

	params := cloud.InstanceGetConsoleParams{}
	console, err := client.Cloud.Instances.GetConsole(context.Background(), instanceID, params)
	if err != nil {
		fmt.Printf("Error getting console: %v\n", err)
		return
	}

	fmt.Printf("Console: protocol=%s, type=%s, url=%s\n", console.RemoteConsole.Protocol, console.RemoteConsole.Type, console.RemoteConsole.URL)
	fmt.Println("===================")
}

func enablePortSecurity(client *gcore.Client, portID string) {
	fmt.Println("\n=== ENABLE PORT SECURITY ===")

	if _, err := client.Cloud.Instances.EnablePortSecurity(context.Background(), portID, cloud.InstanceEnablePortSecurityParams{}); err != nil {
		if strings.Contains(err.Error(), "public network port") {
			fmt.Printf("Skipping port security enable: Cannot change port security on public network ports (port: %s)\n", portID)
		} else {
			fmt.Printf("Error enabling port security: %v\n", err)
		}
	} else {
		fmt.Printf("Enabled port security for port: %s\n", portID)
	}

	fmt.Println("=============================")
}

func disablePortSecurity(client *gcore.Client, portID string) {
	fmt.Println("\n=== DISABLE PORT SECURITY ===")

	if _, err := client.Cloud.Instances.DisablePortSecurity(context.Background(), portID, cloud.InstanceDisablePortSecurityParams{}); err != nil {
		if strings.Contains(err.Error(), "public network port") {
			fmt.Printf("Skipping port security disable: Cannot change port security on public network ports (port: %s)\n", portID)
		} else {
			fmt.Printf("Error disabling port security: %v\n", err)
		}
	} else {
		fmt.Printf("Disabled port security for port: %s\n", portID)
	}

	fmt.Println("==============================")
}
