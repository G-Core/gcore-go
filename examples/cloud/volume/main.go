package main

import (
	"context"
	"fmt"
	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/option"
	"log"
	"os"
	"strconv"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	//apiKey := os.Getenv("GCORE_API_KEY")
	// Will use Production API URL if omitted
	//baseURL := os.Getenv("GCORE_API_URL")

	// TODO set cloud project and region IDs before running
	cloudProjectId, _ := strconv.ParseInt(os.Getenv("GCORE_CLOUD_PROJECT_ID"), 10, 64)
	cloudRegionId, _ := strconv.ParseInt(os.Getenv("GCORE_CLOUD_REGION_ID"), 10, 64)

	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectId),
		option.WithCloudRegionID(cloudRegionId),
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
			Name:   "gcore-go-sdk-example",
			Size:   1,
			Source: "new-volume",
		},
	}

	taskList, err := client.Cloud.Volumes.New(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating volume: %v", err)
	}

	task, err := client.Cloud.Tasks.Poll(context.Background(), taskList.Tasks[0])
	if err != nil {
		log.Fatalf("Error polling task: %v", err)
	}

	volumeID := task.CreatedResources.Volumes[0]
	fmt.Printf("Created Volume ID: %s\n", volumeID)
	fmt.Println("=====================")

	return volumeID
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
		Name: "gcore-go-sdk-example-updated",
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

	taskList, err := client.Cloud.Volumes.Resize(context.Background(), volumeID, params)
	if err != nil {
		log.Fatalf("Error resizing volume: %v", err)
	}

	_, err = client.Cloud.Tasks.Poll(context.Background(), taskList.Tasks[0])
	if err != nil {
		log.Fatalf("Error polling resize task: %v", err)
	}

	fmt.Printf("Resized volume %s to %d GiB\n", volumeID, params.Size)
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

	taskList, err := client.Cloud.Volumes.AttachToInstance(context.Background(), volumeID, params)
	if err != nil {
		log.Fatalf("Error attaching volume: %v", err)
	}

	_, err = client.Cloud.Tasks.Poll(context.Background(), taskList.Tasks[0])
	if err != nil {
		log.Fatalf("Error polling attach task: %v", err)
	}

	fmt.Printf("Attached volume %s to instance %s\n", volumeID, instanceID)
	fmt.Println("=====================")
}

func detachVolumeToInstance(client *gcore.Client, volumeID, instanceID string) {
	fmt.Println("\n=== DETACH VOLUME ===")

	params := cloud.VolumeDetachFromInstanceParams{
		InstanceID: instanceID,
	}

	taskList, err := client.Cloud.Volumes.DetachFromInstance(context.Background(), volumeID, params)
	if err != nil {
		log.Fatalf("Error detaching volume: %v", err)
	}

	_, err = client.Cloud.Tasks.Poll(context.Background(), taskList.Tasks[0])
	if err != nil {
		log.Fatalf("Error polling detach task: %v", err)
	}

	fmt.Printf("Detached volume %s from instance %s\n", volumeID, instanceID)
	fmt.Println("=====================")
}

func deleteVolume(client *gcore.Client, volumeID string) {
	fmt.Println("\n=== DELETE VOLUME ===")

	taskList, err := client.Cloud.Volumes.Delete(context.Background(), volumeID, cloud.VolumeDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting volume: %v", err)
	}

	_, err = client.Cloud.Tasks.Poll(context.Background(), taskList.Tasks[0])
	if err != nil {
		log.Fatalf("Error polling delete task: %v", err)
	}

	fmt.Printf("Volume with ID %s successfully deleted\n", volumeID)
	fmt.Println("=====================")
}
