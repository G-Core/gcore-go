package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listVersions(client *gcore.Client) []cloud.K8SClusterVersion {
	fmt.Println("\n=== LIST K8S VERSIONS ===")

	versions, err := client.Cloud.K8S.ListVersions(context.Background(), cloud.K8SListVersionsParams{})
	if err != nil {
		log.Fatalf("Error listing k8s versions: %v", err)
	}

	for i, v := range versions.Results {
		fmt.Printf("  %d. Version: %s\n", i+1, v.Version)
	}

	fmt.Printf("Total k8s versions: %d\n", len(versions.Results))
	fmt.Println("=========================")
	return versions.Results
}

func pickCreateVersion(versions []cloud.K8SClusterVersion) string {
	if len(versions) == 0 {
		log.Fatalf("No k8s versions available for cluster creation")
	}
	// The API returns versions in ascending order. Pick the penultimate
	// version when possible so the example always has a newer version to
	// upgrade to later. Fall back to the only available version otherwise.
	idx := max(len(versions)-2, 0)
	picked := versions[idx].Version
	fmt.Printf("Selected k8s version for creation: %s\n", picked)
	return picked
}
