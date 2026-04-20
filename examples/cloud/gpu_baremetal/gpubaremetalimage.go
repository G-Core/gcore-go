package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listGPUBaremetalClusterImages(client *gcore.Client) []cloud.GPUImage {
	fmt.Println("\n=== LIST GPU BAREMETAL CLUSTER IMAGES ===")

	params := cloud.GPUBaremetalClusterImageListParams{}
	images, err := client.Cloud.GPUBaremetal.Clusters.Images.List(context.Background(), params)
	if err != nil {
		fmt.Printf("Error listing GPU baremetal cluster images: %v\n", err)
		return nil
	}

	displayCount := 3
	if len(images.Results) < displayCount {
		displayCount = len(images.Results)
	}

	for i := 0; i < displayCount; i++ {
		printGPUImageDetails(images.Results[i], i+1)
	}

	printTruncationMessage(len(images.Results), displayCount, "images")

	fmt.Printf("Total GPU baremetal cluster images: %d\n", len(images.Results))
	fmt.Println("===============================")
	return images.Results
}

func uploadGPUBaremetalClusterImage(client *gcore.Client) string {
	fmt.Println("\n=== UPLOAD GPU BAREMETAL CLUSTER IMAGE ===")

	params := cloud.GPUBaremetalClusterImageUploadParams{
		Name:         "gcore-go-example-gpu-baremetal-image",
		URL:          "http://mirror.noris.net/cirros/0.4.0/cirros-0.4.0-x86_64-disk.img",
		Architecture: cloud.GPUBaremetalClusterImageUploadParamsArchitectureX86_64,
		OsType:       cloud.GPUBaremetalClusterImageUploadParamsOsTypeLinux,
		SSHKey:       cloud.GPUBaremetalClusterImageUploadParamsSSHKeyAllow,
		Tags: map[string]string{
			"name": "gcore-go-example",
		},
	}

	image, err := client.Cloud.GPUBaremetal.Clusters.Images.UploadAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error uploading GPU baremetal cluster image: %v", err)
	}

	fmt.Printf("Uploaded image: ID=%s, name=%s, status=%s\n", image.ID, image.Name, image.Status)
	fmt.Println("===============================")
	return image.ID
}

func getGPUBaremetalClusterImage(client *gcore.Client, imageID string) {
	fmt.Println("\n=== GET GPU BAREMETAL CLUSTER IMAGE ===")

	image, err := client.Cloud.GPUBaremetal.Clusters.Images.Get(context.Background(), imageID, cloud.GPUBaremetalClusterImageGetParams{})
	if err != nil {
		fmt.Printf("Error getting GPU baremetal cluster image: %v\n", err)
		return
	}

	printGPUImageDetails(*image, 1)
	fmt.Println("===============================")
}

func deleteGPUBaremetalClusterImage(client *gcore.Client, imageID string) {
	fmt.Println("\n=== DELETE GPU BAREMETAL CLUSTER IMAGE ===")

	err := client.Cloud.GPUBaremetal.Clusters.Images.DeleteAndPoll(context.Background(), imageID, cloud.GPUBaremetalClusterImageDeleteParams{})
	if err != nil {
		fmt.Printf("Error deleting GPU baremetal cluster image: %v\n", err)
		return
	}

	fmt.Printf("Deleted image: ID=%s\n", imageID)
	fmt.Println("===============================")
}

func getUbuntuImage(images []cloud.GPUImage) string {
	for _, img := range images {
		if strings.Contains(strings.ToLower(img.Name), "ubuntu") &&
			img.Status == "active" {
			fmt.Printf("Selected Ubuntu image: %s (ID: %s)\n", img.Name, img.ID)
			return img.ID
		}
	}

	for _, img := range images {
		if strings.Contains(strings.ToLower(img.OsType), "linux") &&
			img.Status == "active" {
			fmt.Printf("Selected Linux image: %s (ID: %s)\n", img.Name, img.ID)
			return img.ID
		}
	}

	for _, img := range images {
		if img.Status == "active" {
			fmt.Printf("Selected image: %s (ID: %s)\n", img.Name, img.ID)
			return img.ID
		}
	}

	log.Fatalf("No active GPU baremetal cluster images found")
	return ""
}

func printGPUImageDetails(image cloud.GPUImage, index int) {
	fmt.Printf("  %d. Image: ID=%s, name=%s\n", index, image.ID, image.Name)
	fmt.Printf("     Status: %s\n", image.Status)

	if image.OsType != "" {
		fmt.Printf("     OS Type: %s\n", image.OsType)
	}
	if image.OsDistro != "" {
		fmt.Printf("     OS Distribution: %s\n", image.OsDistro)
	}
	if image.OsVersion != "" {
		fmt.Printf("     OS Version: %s\n", image.OsVersion)
	}
	if image.Architecture != "" {
		fmt.Printf("     Architecture: %s\n", image.Architecture)
	}

	fmt.Printf("     Min RAM: %d MB, Min Disk: %d GB\n", image.MinRam, image.MinDisk)

	if image.Size > 0 {
		fmt.Printf("     Size: %d bytes\n", image.Size)
	}

	fmt.Printf("     Visibility: %s\n", image.Visibility)

	if image.SSHKey != "" {
		fmt.Printf("     SSH Key Support: %s\n", image.SSHKey)
	}

	fmt.Println()
}
