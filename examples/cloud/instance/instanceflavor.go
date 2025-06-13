package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listFlavors(client *gcore.Client) {
	fmt.Println("\n=== LIST FLAVORS ===")

	flavors, err := client.Cloud.Instances.Flavors.List(context.Background(), cloud.InstanceFlavorListParams{})
	if err != nil {
		fmt.Printf("Error listing flavors: %v\n", err)
		return
	}

	printFlavorDetails(flavors.Results)
	fmt.Println("====================")
}

func listSuitableFlavors(client *gcore.Client, imageID string) {
	fmt.Println("\n=== LIST SUITABLE FLAVORS ===")

	volumes := []cloud.InstanceFlavorListSuitableParamsVolume{
		{
			Source:  "image",
			Size:    gcore.Int(10),
			ImageID: gcore.String(imageID),
		},
	}

	flavors, err := client.Cloud.Instances.Flavors.ListSuitable(context.Background(), cloud.InstanceFlavorListSuitableParams{
		Volumes: volumes,
	})
	if err != nil {
		fmt.Printf("Error listing suitable flavors: %v\n", err)
		return
	}

	fmt.Printf("Suitable flavors (for 10GB image volume): %d\n", len(flavors.Results))
	if len(flavors.Results) == 0 {
		fmt.Println("  No suitable flavors found for the specified volume requirements.")
	} else {
		printFlavorDetails(flavors.Results)
	}
	fmt.Println("=============================")
}

func printFlavorDetails(flavors []cloud.InstanceFlavor) {
	displayCount := 3
	if len(flavors) < displayCount {
		displayCount = len(flavors)
	}

	for i := 0; i < displayCount; i++ {
		flavor := flavors[i]
		fmt.Printf("  %d. Flavor: ID=%s, Name=%s\n", i+1, flavor.FlavorID, flavor.FlavorName)
		fmt.Printf("     RAM: %d MB, VCPUs: %d\n", flavor.Ram, flavor.Vcpus)
		status := "AVAILABLE"
		if flavor.Disabled {
			status = "DISABLED"
		}
		fmt.Printf("     Status: %s\n", status)
		fmt.Println()
	}

	if len(flavors) > displayCount {
		fmt.Printf("  ... and %d more flavors\n", len(flavors)-displayCount)
	}
}
