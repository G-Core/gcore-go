package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listBaremetalFlavors(client *gcore.Client) []cloud.BaremetalFlavor {
	fmt.Println("\n=== LIST BAREMETAL FLAVORS ===")

	params := cloud.BaremetalFlavorListParams{
		IncludeCapacity: gcore.Bool(true),
	}

	flavors, err := client.Cloud.Baremetal.Flavors.List(context.Background(), params)
	if err != nil {
		fmt.Printf("Error listing baremetal flavors: %v\n", err)
		return nil
	}

	displayCount := 3
	if len(flavors.Results) < displayCount {
		displayCount = len(flavors.Results)
	}

	for i := 0; i < displayCount; i++ {
		printFlavorDetails(flavors.Results[i], i+1)
	}

	printTruncationMessage(len(flavors.Results), displayCount, "flavors")

	fmt.Printf("Total baremetal flavors: %d\n", len(flavors.Results))
	fmt.Println("===============================")
	return flavors.Results
}

func getSmallestFlavor(flavors []cloud.BaremetalFlavor) string {
	var availableFlavors []cloud.BaremetalFlavor
	for _, f := range flavors {
		if !f.Disabled && f.Capacity > 0 {
			availableFlavors = append(availableFlavors, f)
		}
	}

	if len(availableFlavors) == 0 {
		log.Fatalf("No available flavors with capacity found")
	}

	smallestFlavor := availableFlavors[0]
	for _, f := range availableFlavors {
		if f.Ram < smallestFlavor.Ram {
			smallestFlavor = f
		}
	}

	fmt.Printf("Selected smallest flavor: %s (RAM: %d MB, Capacity: %d)\n", smallestFlavor.FlavorName, smallestFlavor.Ram, smallestFlavor.Capacity)
	return smallestFlavor.FlavorID
}

func printFlavorDetails(flavor cloud.BaremetalFlavor, index int) {
	fmt.Printf("  %d. Flavor: ID=%s, name=%s\n", index, flavor.FlavorID, flavor.FlavorName)
	fmt.Printf("     RAM: %d MB, VCPUs: %d\n", flavor.Ram, flavor.Vcpus)
	fmt.Printf("     Architecture: %s, OS: %s\n", flavor.Architecture, flavor.OsType)
	fmt.Printf("     Resource Class: %s\n", flavor.ResourceClass)

	status := "AVAILABLE"
	if flavor.Disabled {
		status = "DISABLED"
	}
	fmt.Printf("     Status: %s\n", status)

	if flavor.Capacity != 0 {
		fmt.Printf("     Capacity: %d\n", flavor.Capacity)
	}

	if flavor.PricePerHour != 0 {
		fmt.Printf("     Price per Hour: %.4f %s\n", flavor.PricePerHour, flavor.CurrencyCode)
	}

	if flavor.PricePerMonth != 0 {
		fmt.Printf("     Price per Month: %.2f %s\n", flavor.PricePerMonth, flavor.CurrencyCode)
	}

	fmt.Println()
}

func printTruncationMessage(totalCount, displayCount int, itemType string) {
	if totalCount > displayCount {
		fmt.Printf("  ... and %d more %s\n", totalCount-displayCount, itemType)
	}
}
