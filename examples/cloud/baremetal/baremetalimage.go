package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listBaremetalImages(client *gcore.Client) []cloud.Image {
	fmt.Println("\n=== LIST BAREMETAL IMAGES ===")

	params := cloud.BaremetalImageListParams{
		IncludePrices: gcore.Bool(true),
	}

	images, err := client.Cloud.Baremetal.Images.List(context.Background(), params)
	if err != nil {
		fmt.Printf("Error listing baremetal images: %v\n", err)
		return nil
	}

	displayCount := 3
	if len(images.Results) < displayCount {
		displayCount = len(images.Results)
	}

	for i := 0; i < displayCount; i++ {
		printImageDetails(images.Results[i], i+1)
	}

	printTruncationMessage(len(images.Results), displayCount, "images")

	fmt.Printf("Total baremetal images: %d\n", len(images.Results))
	fmt.Println("===============================")
	return images.Results
}

func listBaremetalImagesWithFilters(client *gcore.Client) {
	fmt.Println("\n=== LIST BAREMETAL IMAGES WITH FILTERS ===")

	// Example with visibility filter
	params := cloud.BaremetalImageListParams{
		Visibility:    cloud.BaremetalImageListParamsVisibilityPublic,
		IncludePrices: gcore.Bool(true),
	}

	images, err := client.Cloud.Baremetal.Images.List(context.Background(), params)
	if err != nil {
		fmt.Printf("Error listing baremetal images with filters: %v\n", err)
		return
	}

	// Show first few results
	displayCount := 3
	if len(images.Results) < displayCount {
		displayCount = len(images.Results)
	}

	for i := 0; i < displayCount; i++ {
		printImageDetails(images.Results[i], i+1)
	}

	printTruncationMessage(len(images.Results), displayCount, "images")

	fmt.Printf("Total baremetal images: %d\n", len(images.Results))
	fmt.Println("===============================")
	fmt.Println()
}

func getUbuntuImage(images []cloud.Image) string {
	return getOSImage(images, "ubuntu")
}

func getDebianImage(images []cloud.Image) string {
	return getOSImage(images, "debian")
}

func getOSImage(images []cloud.Image, osName string) string {
	var osImages []cloud.Image
	for _, img := range images {
		if strings.Contains(img.Name, osName) {
			osImages = append(osImages, img)
		}
	}

	if len(osImages) == 0 {
		var linuxImages []cloud.Image
		for _, img := range images {
			if img.OsType == "linux" {
				linuxImages = append(linuxImages, img)
			}
		}

		if len(linuxImages) == 0 {
			if len(images) > 0 {
				selectedImage := images[0]
				fmt.Printf("No %s/Linux images found, using first available: %s\n", osName, selectedImage.Name)
				return selectedImage.ID
			} else {
				log.Fatalf("No images found")
			}
		} else {
			selectedImage := linuxImages[0]
			fmt.Printf("No %s images found, using first Linux image: %s\n", osName, selectedImage.Name)
			return selectedImage.ID
		}
	} else {
		selectedImage := osImages[0]
		fmt.Printf("Selected %s image: %s\n", osName, selectedImage.Name)
		return selectedImage.ID
	}

	return ""
}

func printImageDetails(image cloud.Image, index int) {
	fmt.Printf("  %d. Image: ID=%s, name=%s\n", index, image.ID, image.Name)
	fmt.Printf("     OS Type: %s, Status: %s\n", image.OsType, image.Status)
	fmt.Printf("     Architecture: %s, Size: %d MB\n", image.Architecture, image.Size)
	fmt.Printf("     Visibility: %s\n", image.Visibility)

	if image.OsDistro != "" {
		fmt.Printf("     OS Distribution: %s\n", image.OsDistro)
	}

	if image.OsVersion != "" {
		fmt.Printf("     OS Version: %s\n", image.OsVersion)
	}

	if len(image.Tags) > 0 {
		fmt.Printf("     Tags: %v\n", image.Tags)
	}

	fmt.Println()
}
