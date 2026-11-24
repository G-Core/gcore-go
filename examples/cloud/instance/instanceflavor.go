package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

// pickFlavor returns the smallest available flavor that can boot the volume
func pickFlavor(client *gcore.Client, regionID int64, bootVolumeID string) string {
	fmt.Println("\n=== LIST FLAVORS ===")

	// without RegionIDs the list holds the flavors of all regions
	iter := client.Cloud.Instances.Flavors.ListAutoPaging(context.Background(), cloud.InstanceFlavorListParams{
		RegionIDs:                  []int64{regionID},
		CompatibleWithBootVolumeID: gcore.String(bootVolumeID),
		Available:                  gcore.Bool(true),
		OrderBy:                    cloud.InstanceFlavorListParamsOrderByVcpusAsc,
	})
	flavors := make([]cloud.InstanceFlavor, 0)
	for iter.Next() {
		flavors = append(flavors, iter.Current())
	}

	if err := iter.Err(); err != nil {
		log.Fatalf("Error iterating flavors: %v", err)
	}
	if len(flavors) == 0 {
		log.Fatalf("No available flavor can boot volume %s in region %d", bootVolumeID, regionID)
	}

	printFlavorDetails(flavors)
	fmt.Println("====================")

	return flavors[0].Name
}

// pickResizeFlavor returns the smallest available flavor that the instance can resize to, or an empty string
func pickResizeFlavor(client *gcore.Client, instanceID string) string {
	fmt.Println("\n=== LIST RESIZE FLAVORS ===")

	iter := client.Cloud.Instances.ResizeFlavors.ListAutoPaging(context.Background(), instanceID, cloud.InstanceResizeFlavorListParams{
		OrderBy: cloud.InstanceResizeFlavorListParamsOrderByVcpusAsc,
	})
	flavors := make([]cloud.InstanceFlavor, 0)
	for iter.Next() {
		flavors = append(flavors, iter.Current())
	}

	if err := iter.Err(); err != nil {
		log.Fatalf("Error iterating resize flavors: %v", err)
	}

	printFlavorDetails(flavors)
	fmt.Println("===========================")

	for _, flavor := range flavors {
		if flavor.Available {
			return flavor.Name
		}
	}
	fmt.Println("No available flavor to resize to")
	return ""
}

func printFlavorDetails(flavors []cloud.InstanceFlavor) {
	displayCount := min(3, len(flavors))

	for i := 0; i < displayCount; i++ {
		flavor := flavors[i]
		fmt.Printf("  %d. Flavor: Name=%s, RegionID=%d\n", i+1, flavor.Name, flavor.RegionID)
		fmt.Printf("     RAM: %d MiB, VCPUs: %d, Available: %t\n", flavor.RamMib, flavor.Vcpus, flavor.Available)
	}

	if len(flavors) > displayCount {
		fmt.Printf("  ... and %d more flavors\n", len(flavors)-displayCount)
	}
}
