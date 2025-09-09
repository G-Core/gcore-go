package main

import (
	"context"
	"fmt"
	"github.com/G-Core/gcore-go/packages/param"
	"log"
	"os"
	"strconv"

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

	// TODO set cloud project and region IDs before running
	cloudProjectID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_PROJECT_ID"), 10, 64)
	if err != nil {
		log.Fatalf("Error parsing GCORE_CLOUD_PROJECT_ID: %v", err)
	}
	cloudRegionID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_REGION_ID"), 10, 64)
	if err != nil {
		log.Fatalf("Error parsing GCORE_CLOUD_REGION_ID: %v", err)
	}

	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectID),
		option.WithCloudRegionID(cloudRegionID),
	)

	// TODO set instance ID before running
	instanceID := os.Getenv("GCORE_CLOUD_INSTANCE_ID")
	if os.Getenv("GCORE_CLOUD_INSTANCE_ID") == "" {
		log.Fatal("GCORE_CLOUD_INSTANCE_ID environment variable is not set")
	}

	// Create a volume and use its ID for other operations
	volumeID := createVolume(&client)
	listAllVolumes(&client)
	getVolumeByID(&client, volumeID)
	updateVolume(&client, volumeID)
	resizeVolume(&client, volumeID)
	changeVolumeType(&client, volumeID)
	attachVolumeToInstance(&client, volumeID, instanceID)
	detachVolumeToInstance(&client, volumeID, instanceID)
	deleteVolume(&client, volumeID)
}

func createVolume(client *gcore.Client) string {
	fmt.Println("\n=== CREATE VOLUME ===")

	params := cloud.VolumeNewParams{
		OfNewVolume: &cloud.VolumeNewParamsBodyNewVolume{
			Name:   "gcore-go-example",
			Size:   1,
			Source: constant.NewVolume("").Default(),
		},
	}

	volume, err := client.Cloud.Volumes.NewAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating volume: %v", err)
	}

	fmt.Printf("Created Volume ID: %s\n", volume.ID)
	fmt.Println("=====================")

	return volume.ID
}

func getVolumeByID(client *gcore.Client, volumeID string) {
	fmt.Println("\n=== GET VOLUME BY ID ===")

	volume, err := client.Cloud.Volumes.Get(context.Background(), volumeID, cloud.VolumeGetParams{})
	if err != nil {
		log.Fatalf("Error getting volume: %v", err)
	}

	fmt.Printf("Volume ID: %s, Name: %s, Size: %d GiB\n", volume.ID, volume.Name, volume.Size)
	fmt.Println("========================")
}

func listAllVolumes(client *gcore.Client) {
	fmt.Println("\n=== LIST ALL VOLUMES ===")

	volumesPage, err := client.Cloud.Volumes.List(context.Background(), cloud.VolumeListParams{})
	if err != nil {
		log.Fatalf("Error listing volumes: %v", err)
	}

	for i, volume := range volumesPage.Results {
		fmt.Printf("  %d. Volume ID: %s, Name: %s, Size: %d GiB\n", i+1, volume.ID, volume.Name, volume.Size)
	}

	fmt.Println("========================")
}

func updateVolume(client *gcore.Client, volumeID string) {
	fmt.Println("\n=== UPDATE VOLUME ===")

	params := cloud.VolumeUpdateParams{
		Name: param.NewOpt("gcore-go-example-updated"),
	}

	volume, err := client.Cloud.Volumes.Update(context.Background(), volumeID, params)
	if err != nil {
		log.Fatalf("Error updating volume: %v", err)
	}

	fmt.Printf("Updated Volume ID: %s, Name: %s\n", volume.ID, volume.Name)
	fmt.Println("=====================")
}

func resizeVolume(client *gcore.Client, volumeID string) {
	fmt.Println("\n=== RESIZE VOLUME ===")

	params := cloud.VolumeResizeParams{
		Size: 2,
	}

	volume, err := client.Cloud.Volumes.ResizeAndPoll(context.Background(), volumeID, params)
	if err != nil {
		log.Fatalf("Error resizing volume: %v", err)
	}

	fmt.Printf("Resized volume %s to %d GiB\n", volumeID, volume.Size)
	fmt.Println("=====================")
}

func changeVolumeType(client *gcore.Client, volumeID string) {
	fmt.Println("\n=== CHANGE VOLUME TYPE ===")

	params := cloud.VolumeChangeTypeParams{
		VolumeType: cloud.VolumeChangeTypeParamsVolumeTypeSsdHiiops,
	}

	_, err := client.Cloud.Volumes.ChangeType(context.Background(), volumeID, params)
	if err != nil {
		log.Fatalf("Error changing volume type: %v", err)
	}

	fmt.Printf("Changed volume %s type to: %s\n", volumeID, params.VolumeType)
	fmt.Println("==========================")
}

func attachVolumeToInstance(client *gcore.Client, volumeID, instanceID string) {
	fmt.Println("\n=== ATTACH VOLUME ===")

	params := cloud.VolumeAttachToInstanceParams{
		InstanceID: instanceID,
	}

	err := client.Cloud.Volumes.AttachToInstanceAndPoll(context.Background(), volumeID, params)
	if err != nil {
		log.Fatalf("Error attaching volume: %v", err)
	}
	fmt.Printf("Attached volume %s to instance %s\n", volumeID, instanceID)
	fmt.Println("=====================")
}

func detachVolumeToInstance(client *gcore.Client, volumeID, instanceID string) {
	fmt.Println("\n=== DETACH VOLUME ===")

	params := cloud.VolumeDetachFromInstanceParams{
		InstanceID: instanceID,
	}

	err := client.Cloud.Volumes.DetachFromInstanceAndPoll(context.Background(), volumeID, params)
	if err != nil {
		log.Fatalf("Error detaching volume: %v", err)
	}
	fmt.Printf("Detached volume %s from instance %s\n", volumeID, instanceID)
	fmt.Println("=====================")
}

func deleteVolume(client *gcore.Client, volumeID string) {
	fmt.Println("\n=== DELETE VOLUME ===")

	err := client.Cloud.Volumes.DeleteAndPoll(context.Background(), volumeID, cloud.VolumeDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting volume: %v", err)
	}

	fmt.Printf("Volume with ID %s successfully deleted\n", volumeID)
	fmt.Println("=====================")
}
