package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listFlavors(client *gcore.Client) []cloud.BaremetalFlavor {
	fmt.Println("\n=== LIST K8S FLAVORS ===")

	params := cloud.K8SFlavorListParams{
		ExcludeGPU:      gcore.Bool(true),
		IncludeCapacity: gcore.Bool(true),
	}

	flavors, err := client.Cloud.K8S.Flavors.List(context.Background(), params)
	if err != nil {
		log.Fatalf("Error listing k8s flavors: %v", err)
	}

	displayCount := 3
	if len(flavors.Results) < displayCount {
		displayCount = len(flavors.Results)
	}

	for i := 0; i < displayCount; i++ {
		f := flavors.Results[i]
		fmt.Printf("  %d. Flavor: ID=%s, Name=%s, RAM=%d MB, VCPUs=%d\n",
			i+1, f.FlavorID, f.FlavorName, f.Ram, f.Vcpus)
	}

	if len(flavors.Results) > displayCount {
		fmt.Printf("  ... and %d more flavors\n", len(flavors.Results)-displayCount)
	}

	fmt.Printf("Total k8s flavors: %d\n", len(flavors.Results))
	fmt.Println("========================")
	return flavors.Results
}

func pickPoolFlavor(flavors []cloud.BaremetalFlavor) string {
	// Prefer the smallest g1-standard VM flavor with capacity; fall back to
	// any available non-baremetal flavor. Skip disabled, zero-capacity, and
	// test/fake flavors.
	var candidates []cloud.BaremetalFlavor
	for _, f := range flavors {
		if f.Disabled || f.Capacity <= 0 {
			continue
		}
		if strings.Contains(f.FlavorName, "test") || strings.Contains(f.FlavorName, "fake") {
			continue
		}
		if !strings.HasPrefix(f.FlavorName, "g1-standard-") {
			continue
		}
		candidates = append(candidates, f)
	}

	if len(candidates) == 0 {
		log.Fatalf("No suitable g1-standard k8s flavor with capacity found")
	}

	pick := candidates[0]
	for _, f := range candidates {
		if f.Ram < pick.Ram {
			pick = f
		}
	}

	fmt.Printf("Selected k8s pool flavor: %s (RAM: %d MB, VCPUs: %d)\n",
		pick.FlavorName, pick.Ram, pick.Vcpus)
	return pick.FlavorID
}
