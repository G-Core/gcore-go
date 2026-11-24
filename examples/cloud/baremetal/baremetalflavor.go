package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

// listBaremetalFlavors lists the flavors of the region that can boot the image. With flavorName, it lists
// only that flavor
func listBaremetalFlavors(client *gcore.Client, regionID int64, imageID, flavorName string) []cloud.BaremetalFlavor {
	fmt.Println("\n=== LIST BAREMETAL FLAVORS ===")

	// without RegionIDs the list holds the flavors of all regions
	params := cloud.BaremetalFlavorListParams{
		RegionIDs:             []int64{regionID},
		CompatibleWithImageID: gcore.String(imageID),
		OrderBy:               cloud.BaremetalFlavorListParamsOrderByRamMibAsc,
	}
	if flavorName != "" {
		params.Name = cloud.BaremetalFlavorListParamsName{Exact: []string{flavorName}}
	}

	var flavors []cloud.BaremetalFlavor
	iter := client.Cloud.Baremetal.Flavors.ListAutoPaging(context.Background(), params)
	for iter.Next() {
		flavors = append(flavors, iter.Current())
	}
	if err := iter.Err(); err != nil {
		log.Fatalf("Error listing baremetal flavors: %v", err)
	}

	displayCount := 3
	if len(flavors) < displayCount {
		displayCount = len(flavors)
	}

	for i := 0; i < displayCount; i++ {
		printFlavorDetails(flavors[i], i+1)
	}

	printTruncationMessage(len(flavors), displayCount, "flavors")

	fmt.Printf("Total baremetal flavors: %d\n", len(flavors))
	fmt.Println("===============================")
	return flavors
}

// pickFlavor returns the flavor with the smallest RAM that can boot the image and has capacity in the region
func pickFlavor(client *gcore.Client, regionID int64, imageID, flavorName string) cloud.BaremetalFlavor {
	flavors := listBaremetalFlavors(client, regionID, imageID, flavorName)

	// the list is in RAM order, so the first flavor with capacity is the smallest
	for _, flavor := range flavors {
		if flavor.Capacity > 0 {
			fmt.Printf("Selected flavor: %s (RAM: %d MiB, CPU threads: %d, Capacity: %d)\n",
				flavor.Name, flavor.RamMib, flavor.CPUThreads, flavor.Capacity)
			printNetworkingCapabilities(flavor)
			return flavor
		}
	}

	log.Fatalf("No flavor with capacity can boot image %s in region %d", imageID, regionID)
	return cloud.BaremetalFlavor{}
}

// printNetworkingCapabilities prints the port features that the flavor supports. The ports of the server
// must use only these features
func printNetworkingCapabilities(flavor cloud.BaremetalFlavor) {
	capabilities := flavor.NetworkingCapabilities
	fmt.Printf("Networking capabilities: SecurityGroups=%t, PortSecurity=%t, AllowedAddressPairs=%t\n",
		capabilities.SecurityGroups, capabilities.PortSecurity, capabilities.AllowedAddressPairs)
}

func printFlavorDetails(flavor cloud.BaremetalFlavor, index int) {
	fmt.Printf("  %d. Flavor: Name=%s, RegionID=%d\n", index, flavor.Name, flavor.RegionID)
	fmt.Printf("     RAM: %d MiB, CPU threads: %d\n", flavor.RamMib, flavor.CPUThreads)
	fmt.Printf("     Architecture: %s, OS: %s\n", flavor.Architecture, flavor.OsType)
	fmt.Printf("     Resource Class: %s\n", flavor.ResourceClass)
	fmt.Printf("     Capacity: %d (reserved: %d)\n", flavor.Capacity, flavor.ReservedCapacity)

	if flavor.JSON.HardwareDescription.Valid() {
		hardware := flavor.HardwareDescription
		fmt.Printf("     Hardware: CPU=%s, RAM=%s, Disk=%s, Network=%s\n", hardware.CPU, hardware.Ram, hardware.Disk, hardware.Network)
	}

	fmt.Println()
}

func printTruncationMessage(totalCount, displayCount int, itemType string) {
	if totalCount > displayCount {
		fmt.Printf("  ... and %d more %s\n", totalCount-displayCount, itemType)
	}
}
