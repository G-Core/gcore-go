package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/packages/param"
)

func listGPUVirtualClusterFlavors(client *gcore.Client) []cloud.GPUVirtualFlavorUnion {
	fmt.Println("\n=== LIST GPU VIRTUAL CLUSTER FLAVORS ===")

	params := cloud.GPUVirtualClusterFlavorListParams{
		HideDisabled:  param.NewOpt(true),
		IncludePrices: param.NewOpt(true),
	}

	flavors, err := client.Cloud.GPUVirtual.Clusters.Flavors.List(context.Background(), params)
	if err != nil {
		fmt.Printf("Error listing GPU virtual cluster flavors: %v\n", err)
		return nil
	}

	displayCount := min(3, len(flavors.Results))

	for i := 0; i < displayCount; i++ {
		printGPUFlavorDetails(flavors.Results[i], i+1)
	}

	printTruncationMessage(len(flavors.Results), displayCount, "flavors")

	fmt.Printf("Total GPU virtual cluster flavors: %d\n", len(flavors.Results))
	fmt.Println("===============================")
	return flavors.Results
}

func getAvailableFlavor(flavors []cloud.GPUVirtualFlavorUnion) string {
	var availableFlavors []cloud.GPUVirtualFlavorUnion
	for _, f := range flavors {
		if !f.Disabled && f.Capacity > 0 {
			availableFlavors = append(availableFlavors, f)
		}
	}

	if len(availableFlavors) == 0 {
		fmt.Printf("WARNING: No available GPU virtual flavors with capacity found in this region.\n")
		fmt.Printf("   GPU virtual clusters require specialized hardware that may not be available in all regions.\n")
		fmt.Printf("   Please try a different region or contact support for GPU virtual cluster availability.\n")

		if len(flavors) > 0 {
			fmt.Printf("   Using first available flavor for demonstration: %s\n", flavors[0].Name)
			return flavors[0].Name
		}
		log.Fatalf("No GPU virtual flavors found at all")
	}

	selectedFlavor := availableFlavors[0]
	fmt.Printf("Selected GPU virtual flavor: %s (Capacity: %d)\n", selectedFlavor.Name, selectedFlavor.Capacity)
	return selectedFlavor.Name
}

func printGPUFlavorDetails(flavor cloud.GPUVirtualFlavorUnion, index int) {
	fmt.Printf("  %d. Flavor: name=%s\n", index, flavor.Name)
	fmt.Printf("     Architecture: %s\n", flavor.Architecture)

	if flavor.HardwareDescription.GPU != "" {
		fmt.Printf("     GPU: %s\n", flavor.HardwareDescription.GPU)
	}
	if flavor.HardwareDescription.Vcpus > 0 {
		fmt.Printf("     vCPUs: %d\n", flavor.HardwareDescription.Vcpus)
	}
	if flavor.HardwareDescription.Ram > 0 {
		fmt.Printf("     RAM: %d GB\n", flavor.HardwareDescription.Ram)
	}
	if flavor.HardwareDescription.LocalStorage > 0 {
		fmt.Printf("     Local Storage: %d GB\n", flavor.HardwareDescription.LocalStorage)
	}

	if flavor.HardwareProperties.GPUCount > 0 {
		fmt.Printf("     GPU Count: %d\n", flavor.HardwareProperties.GPUCount)
	}
	if flavor.HardwareProperties.GPUManufacturer != "" {
		fmt.Printf("     GPU Manufacturer: %s\n", flavor.HardwareProperties.GPUManufacturer)
	}
	if flavor.HardwareProperties.GPUModel != "" {
		fmt.Printf("     GPU Model: %s\n", flavor.HardwareProperties.GPUModel)
	}

	status := "AVAILABLE"
	if flavor.Disabled {
		status = "DISABLED"
	}
	fmt.Printf("     Status: %s\n", status)

	if flavor.Capacity != 0 {
		fmt.Printf("     Capacity: %d\n", flavor.Capacity)
	}

	if withPrice := flavor.AsVirtualGPUFlavorsChemaWithPrice(); withPrice.Price.PricePerHour != 0 {
		fmt.Printf("     Price per Hour: %.4f %s\n", withPrice.Price.PricePerHour, withPrice.Price.CurrencyCode)
		if withPrice.Price.PricePerMonth != 0 {
			fmt.Printf("     Price per Month: %.2f %s\n", withPrice.Price.PricePerMonth, withPrice.Price.CurrencyCode)
		}
	}

	fmt.Println()
}
