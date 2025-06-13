package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func createImageFromVolume(client *gcore.Client, volumeID string) string {
	fmt.Println("\n=== CREATE IMAGE FROM VOLUME ===")

	params := cloud.InstanceImageNewFromVolumeParams{
		VolumeID: volumeID,
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

	fmt.Printf("Uploaded Image: ID=%s, Name=%s, OS Type=%s, Arch=%s, Status=%s, Size=%d\n", image.ID, image.Name, image.OsType, image.Architecture, image.Status, image.Size)
	fmt.Println("====================")

	return image.ID
}

func listImages(client *gcore.Client) {
	fmt.Println("\n=== LIST ALL IMAGES ===")

	imageList, err := client.Cloud.Instances.Images.List(context.Background(), cloud.InstanceImageListParams{})
	if err != nil {
		fmt.Printf("Error listing images: %v\n", err)
		return
	}

	displayCount := 3
	if len(imageList.Results) < displayCount {
		displayCount = len(imageList.Results)
	}

	for i := 0; i < displayCount; i++ {
		img := imageList.Results[i]
		fmt.Printf("  %d. Image ID: %s, Name: %s, OS Type: %s, Status: %s\n", i+1, img.ID, img.Name, img.OsType, img.Status)
	}

	if len(imageList.Results) > displayCount {
		fmt.Printf("  ... and %d more images\n", len(imageList.Results)-displayCount)
	}

	fmt.Println("=======================")
}

func getImage(client *gcore.Client, imageID string) {
	fmt.Println("\n=== GET IMAGE BY ID ===")

	image, err := client.Cloud.Instances.Images.Get(context.Background(), imageID, cloud.InstanceImageGetParams{})
	if err != nil {
		fmt.Printf("Error getting image: %v\n", err)
		return
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
		fmt.Printf("Error updating image: %v\n", err)
		return
	}

	fmt.Printf("Updated Image ID: %s, Name: %s\n", updatedImage.ID, updatedImage.Name)
	fmt.Println("====================")
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
