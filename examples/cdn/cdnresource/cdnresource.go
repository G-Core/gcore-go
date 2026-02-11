package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cdn"
	"github.com/G-Core/gcore-go/packages/param"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read
	// from the GCORE_API_KEY environment variable if omitted

	// Origin group ID is required to create a CDN resource
	originGroupID, err := strconv.ParseInt(os.Getenv("GCORE_CDN_ORIGIN_GROUP_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CDN_ORIGIN_GROUP_ID environment variable is required and must be a valid integer")
	}

	client := gcore.NewClient()

	// Create a CDN resource
	resource := createCDNResource(&client, originGroupID)

	// Demonstrate DeactivateAndDelete - this first deactivates the resource
	// by setting active=false, then deletes it
	deactivateAndDeleteCDNResource(&client, resource.ID)
}

func createCDNResource(client *gcore.Client, originGroupID int64) *cdn.CDNResource {
	fmt.Println("\n=== CREATE CDN RESOURCE ===")

	// Use a unique CNAME to avoid conflicts
	cname := fmt.Sprintf("cdn-example-%d.example.com", time.Now().Unix())

	params := cdn.CDNResourceNewParams{
		Cname:       cname,
		OriginGroup: param.NewOpt(originGroupID),
		Active:      gcore.Bool(true),
	}

	resource, err := client.CDN.CDNResources.New(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating CDN resource: %v", err)
	}

	fmt.Printf("Created CDN Resource: ID=%d, Cname=%s, Active=%v\n",
		resource.ID, resource.Cname, resource.Active)
	fmt.Println("============================")

	return resource
}

func deactivateAndDeleteCDNResource(client *gcore.Client, resourceID int64) {
	fmt.Println("\n=== DEACTIVATE AND DELETE CDN RESOURCE ===")

	// DeactivateAndDelete is a utility method that:
	// 1. First deactivates the CDN resource by setting active=false
	// 2. Then deletes the resource permanently
	//
	// This is useful because the Delete operation requires the CDN resource
	// to be deactivated first.
	err := client.CDN.CDNResources.DeactivateAndDelete(context.Background(), resourceID)
	if err != nil {
		log.Fatalf("Error deactivating and deleting CDN resource: %v", err)
	}

	fmt.Printf("CDN Resource with ID %d successfully deactivated and deleted\n", resourceID)
	fmt.Println("==========================================")
}
