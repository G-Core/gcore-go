package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listGPUVirtualClusterImages(client *gcore.Client) []cloud.GPUImage {
	fmt.Println("\n=== LIST GPU VIRTUAL CLUSTER IMAGES ===")

	params := cloud.GPUVirtualClusterImageListParams{}
	images, err := client.Cloud.GPUVirtual.Clusters.Images.List(context.Background(), params)
	if err != nil {
		fmt.Printf("Error listing GPU virtual cluster images: %v\n", err)
		return nil
	}

	displayCount := min(3, len(images.Results))

	for i := 0; i < displayCount; i++ {
		printGPUImageDetails(images.Results[i], i+1)
	}

	printTruncationMessage(len(images.Results), displayCount, "images")

	fmt.Printf("Total GPU virtual cluster images: %d\n", len(images.Results))
	fmt.Println("===============================")
	return images.Results
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

	log.Fatalf("No active GPU virtual cluster images found")
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
