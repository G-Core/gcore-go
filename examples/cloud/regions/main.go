package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/stainless-sdks/gcore-go"
	"github.com/stainless-sdks/gcore-go/cloud"
	"github.com/stainless-sdks/gcore-go/option"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	apiKey := os.Getenv("GCORE_API_KEY")
	client := gcore.NewClient(
		option.WithAPIKey(apiKey),
	)

	getRegionByID(&client)
	listAllRegions(&client)
	listRegionsWithFilters(&client)
	listRegionsWithAutopager(&client)
}

func getRegionByID(client *gcore.Client) {
	// Region ID can also be omitted — it defaults to the GCORE_REGION environment variable
	regionID := int64(76)
	if id, err := strconv.ParseInt(os.Getenv("GCORE_REGION"), 10, 64); err == nil {
		regionID = id
	}

	fmt.Println("\n=== GET REGION BY ID ===")

	region, err := client.Cloud.Regions.Get(context.Background(), cloud.RegionGetParams{
		RegionID: gcore.Int(regionID),
	})
	if err != nil {
		log.Fatalf("Error getting region: %v", err)
	}

	fmt.Printf("Region ID: %d, Display Name: %s\n", region.ID, region.DisplayName)
	fmt.Println("========================")
}

func listAllRegions(client *gcore.Client) {
	fmt.Println("\n=== LIST ALL REGIONS ===")

	regions, err := client.Cloud.Regions.List(context.Background(), cloud.RegionListParams{})
	if err != nil {
		log.Fatalf("Error listing regions: %v", err)
	}

	for i, region := range regions.Results {
		fmt.Printf("  %d. Region ID: %d, Display Name: %s\n", i+1, region.ID, region.DisplayName)
	}

	fmt.Println("========================")
}

func listRegionsWithFilters(client *gcore.Client) {
	fmt.Println("\n=== LIST REGIONS WITH FILTERS ===")

	regions, err := client.Cloud.Regions.List(context.Background(), cloud.RegionListParams{
		OrderBy:         cloud.RegionListParamsOrderByCreatedAtAsc,
		Product:         cloud.RegionListParamsProductContainers,
		ShowVolumeTypes: gcore.Bool(true),
	})
	if err != nil {
		log.Fatalf("Error listing regions with filters: %v", err)
	}

	for i, region := range regions.Results {
		fmt.Printf("  %d. Region ID: %d, Display Name: %s\n", i+1, region.ID, region.DisplayName)
	}

	fmt.Println("=================================")
}

func listRegionsWithAutopager(client *gcore.Client) {
	fmt.Println("\n=== LIST REGIONS USING AUTOPAGER ===")

	count := 0

	iter := client.Cloud.Regions.ListAutoPaging(context.Background(), cloud.RegionListParams{
		Limit: gcore.Int(10), // Process 10 items per page
	})

	for iter.Next() {
		region := iter.Current()
		count++
		fmt.Printf("  %d. Region ID: %d, Display Name: %s\n",
			count, region.ID, region.DisplayName)
	}

	if err := iter.Err(); err != nil {
		log.Fatalf("Error iterating regions: %v", err)
	}

	fmt.Println("====================================")
}
