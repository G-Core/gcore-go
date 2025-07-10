package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listFlavors(client *gcore.Client) {
	fmt.Println("\n=== LIST FLAVORS ===")

	iter := client.Cloud.Inference.Flavors.ListAutoPaging(context.Background(), cloud.InferenceFlavorListParams{})
	flavorList := make([]cloud.InferenceFlavor, 0)
	for iter.Next() {
		flavorList = append(flavorList, iter.Current())
	}

	if err := iter.Err(); err != nil {
		fmt.Printf("Error iterating flavors: %v\n", err)
		return
	}

	displayCount := 3
	if len(flavorList) < displayCount {
		displayCount = len(flavorList)
	}

	for i := 0; i < displayCount; i++ {
		flavor := flavorList[i]
		fmt.Printf("  %d. Flavor: %s, CPU: %.1f, Memory: %.1f Gi\n", i+1, flavor.Name, flavor.CPU, flavor.Memory)
	}

	if len(flavorList) > displayCount {
		fmt.Printf("  ... and %d more flavors\n", len(flavorList)-displayCount)
	}

	fmt.Println("========================")
}

func getFlavor(client *gcore.Client, flavorName string) {
	fmt.Println("\n=== GET FLAVOR ===")

	flavor, err := client.Cloud.Inference.Flavors.Get(context.Background(), flavorName)
	if err != nil {
		fmt.Printf("Error getting flavor: %v\n", err)
		return
	}

	fmt.Printf("Flavor: %s, CPU: %.1f, Memory: %.1f Gi\n", flavor.Name, flavor.CPU, flavor.Memory)
	fmt.Println("========================")
}
