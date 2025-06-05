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
	//baseURL := os.Getenv("GCORE_API_URL")

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

	// TODO set bootable volume ID before running
	volumeID := os.Getenv("GCORE_CLOUD_VOLUME_ID")
	if volumeID == "" {
		log.Fatalf("GCORE_CLOUD_VOLUME_ID environment variable is not set")
	}

	// Create an image from volume and use its ID for other operations
	volumeImageID := createImageFromVolume(&client, volumeID)
	uploadedImageID := uploadImage(&client)
	listImages(&client)
	getImage(&client, volumeImageID)
	updateImage(&client, volumeImageID)
	deleteImage(&client, uploadedImageID)
	deleteImage(&client, volumeImageID)
}

func createImageFromVolume(client *gcore.Client, volumeId string) string {
	fmt.Println("\n=== CREATE IMAGE FROM VOLUME ===")

	params := cloud.InstanceImageNewFromVolumeParams{
		VolumeID: volumeId,
		Name:     "gcore-go-example",
	}

	image, err := client.Cloud.Instances.Images.NewFromVolumeAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating image: %v", err)
	}

	fmt.Printf("Created Image ID: %s\n", image.ID)
	fmt.Println("================================")

	return image.ID
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
		Name: gcore.String("gcore-go-example-updated"),
	}

	updatedImage, err := client.Cloud.Instances.Images.Update(context.Background(), imageID, params)
	if err != nil {
		log.Fatalf("Error updating image: %v", err)
	}

	fmt.Printf("Updated Image ID: %s, Name: %s\n", updatedImage.ID, updatedImage.Name)
	fmt.Println("====================")
}

func uploadImage(client *gcore.Client) string {
	fmt.Println("\n=== UPLOAD IMAGE ===")

	params := cloud.InstanceImageUploadParams{
		Name:         "gcore-go-example-uploaded",
		URL:          "https://cloud-images.ubuntu.com/releases/24.04/release/ubuntu-24.04-server-cloudimg-amd64.img",
		OsType:       cloud.InstanceImageUploadParamsOsTypeLinux,
		Architecture: cloud.InstanceImageUploadParamsArchitectureX86_64,
		SSHKey:       cloud.InstanceImageUploadParamsSSHKeyAllow,
		OsDistro:     gcore.String("Ubuntu"),
		OsVersion:    gcore.String("24.04"),
	}

	image, err := client.Cloud.Instances.Images.UploadAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error uploading image: %v", err)
	}

	fmt.Printf("Uploaded Image ID: %s\n", image.ID)
	fmt.Println("====================")

	return image.ID
}

func deleteImage(client *gcore.Client, imageID string) {
	fmt.Println("\n=== DELETE IMAGE ===")

	err := client.Cloud.Instances.Images.DeleteAndPoll(context.Background(), imageID, cloud.InstanceImageDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting image: %v", err)
	}

	fmt.Printf("Image with ID %s successfully deleted\n", imageID)
	fmt.Println("====================")
}
