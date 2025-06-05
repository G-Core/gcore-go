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

	// TODO set bootable volume ID before running
	volumeId := os.Getenv("GCORE_CLOUD_VOLUME_ID")
	if os.Getenv("GCORE_CLOUD_VOLUME_ID") == "" {
		log.Fatal("GCORE_CLOUD_VOLUME_ID environment variable is not set")
	}

	// Create an image and use its ID for other operations
	imageID := createImageFromVolume(&client, volumeId)
	listImages(&client)
	getImage(&client, imageID)
	updateImage(&client, imageID)
	deleteImage(&client, imageID)
}

func createImageFromVolume(client *gcore.Client, volumeId string) string {
	fmt.Println("\n=== CREATE IMAGE FROM VOLUME ===")

	params := cloud.InstanceImageNewFromVolumeParams{
		VolumeID: volumeId,
		Name:     gcore.String("sdk-image-go-example").Value,
		OsType:   cloud.InstanceImageNewFromVolumeParamsOsType(gcore.String("linux").Value),
	}

	taskIDList, err := client.Cloud.Instances.Images.NewFromVolume(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating image: %v", err)
	}

	task, err := client.Cloud.Tasks.Poll(context.Background(), taskIDList.Tasks[0])
	if err != nil {
		log.Fatalf("Error polling task: %v", err)
	}

	imageID := task.CreatedResources.Images[0]
	fmt.Printf("Created Image ID: %s\n", imageID)
	fmt.Println("================================")

	return imageID
}

func listImages(client *gcore.Client) {
	fmt.Println("\n=== LIST ALL IMAGES ===")

	imageList, err := client.Cloud.Instances.Images.List(context.Background(), cloud.InstanceImageListParams{})
	if err != nil {
		log.Fatalf("Error listing images: %v", err)
	}

	for i, img := range imageList.Results {
		fmt.Printf("  %d. Image ID: %s, Name: %s, OS Type: %s, Status: %s\n", i+1, img.ID, img.Name, img.OsType, img.Status)
	}

	fmt.Println("=======================")
}

func getImage(client *gcore.Client, imageID string) {
	fmt.Println("\n=== GET IMAGE BY ID ===")

	image, err := client.Cloud.Instances.Images.Get(context.Background(), imageID, cloud.InstanceImageGetParams{})
	if err != nil {
		log.Fatalf("Error getting image: %v", err)
	}

	fmt.Printf("Image ID: %s, Name: %s, OS Type: %s, Status: %s\n", image.ID, image.Name, image.OsType, image.Status)
	fmt.Println("=======================")
}

func updateImage(client *gcore.Client, imageID string) {
	fmt.Println("\n=== UPDATE IMAGE ===")

	params := cloud.InstanceImageUpdateParams{
		Name: gcore.String("sdk-image-go-example-updated"),
	}

	updatedImage, err := client.Cloud.Instances.Images.Update(context.Background(), imageID, params)
	if err != nil {
		log.Fatalf("Error updating image: %v", err)
	}

	fmt.Printf("Updated Image ID: %s, Name: %s\n", updatedImage.ID, updatedImage.Name)
	fmt.Println("====================")
}

func deleteImage(client *gcore.Client, imageID string) {
	fmt.Println("\n=== DELETE IMAGE ===")

	taskList, err := client.Cloud.Instances.Images.Delete(context.Background(), imageID, cloud.InstanceImageDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting image: %v", err)
	}

	_, err = client.Cloud.Tasks.Poll(context.Background(), taskList.Tasks[0])
	if err != nil {
		log.Fatalf("Error polling delete task: %v", err)
	}

	fmt.Printf("Image with ID %s successfully deleted\n", imageID)
	fmt.Println("====================")
}
