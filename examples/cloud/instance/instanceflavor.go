package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listFlavors(client *gcore.Client) {
	fmt.Println("\n=== LIST FLAVORS ===")

	// ListAutoPaging fetches all following pages automatically.
	iter := client.Cloud.Instances.Flavors.ListAutoPaging(context.Background(), cloud.InstanceFlavorListParams{})
	flavors := make([]cloud.InstanceFlavorDetailedUnion, 0)
	for iter.Next() {
		flavors = append(flavors, iter.Current())
	}

	if err := iter.Err(); err != nil {
		fmt.Printf("Error iterating flavors: %v\n", err)
		return
	}

	printFlavorDetails(flavors)
	fmt.Println("====================")
}

func printFlavorDetails(flavors []cloud.InstanceFlavorDetailedUnion) {
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
