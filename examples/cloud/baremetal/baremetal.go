package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/shared/constant"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	//apiKey:= os.Getenv("GCORE_API_KEY")
	// Will use Production API URL if omitted
	//baseURL:= os.Getenv("GCORE_BASE_URL")

	// TODO set cloud project ID before running
	cloudProjectID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_PROJECT_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_PROJECT_ID environment variable is required and must be a valid integer")
	}

	// TODO set cloud region ID before running
	cloudRegionID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_REGION_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_REGION_ID environment variable is required and must be a valid integer")
	}

	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectID),
		option.WithCloudRegionID(cloudRegionID),
	)

	// Flavors
	flavors := listBaremetalFlavors(&client)

	// Images
	images := listBaremetalImages(&client)
	listBaremetalImagesWithFilters(&client)

	flavorID := getSmallestFlavor(flavors)
	ubuntuImageID := getUbuntuImage(images)
	debianImageID := getDebianImage(images)

	// Servers
	serverID := createServer(&client, flavorID, ubuntuImageID)
	listServers(&client)
	listServersWithAutoPager(&client)
	rebuildServer(&client, serverID, debianImageID)
}

func createServer(client *gcore.Client, flavorID string, imageID string) string {
	fmt.Println("\n=== CREATE BAREMETAL SERVER ===")

	params := cloud.BaremetalServerNewParams{
		Name:    gcore.String("gcore-go-example"),
		Flavor:  flavorID,
		ImageID: gcore.String(imageID),
		Interfaces: []cloud.BaremetalServerNewParamsInterfaceUnion{{
			OfExternal: &cloud.BaremetalServerNewParamsInterfaceExternal{
				Type: constant.External("external"),
			},
		}},
		Password: gcore.String("Gcore123!"),
		Tags: map[string]string{
			"name": "gcore-go-example",
		},
	}

	server, err := client.Cloud.Baremetal.Servers.NewAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating baremetal server: %v", err)
	}

	fmt.Printf("Created baremetal server: ID=%s, name=%s, status=%s\n", server.ID, server.Name, server.Status)
	fmt.Println("===============================")
	return server.ID
}

func listServers(client *gcore.Client) {
	fmt.Println("\n=== LIST BAREMETAL SERVERS ===")

	params := cloud.BaremetalServerListParams{}
	servers, err := client.Cloud.Baremetal.Servers.List(context.Background(), params)
	if err != nil {
		fmt.Printf("Error listing baremetal servers: %v\n", err)
		return
	}

	displayCount := 3
	if len(servers.Results) < displayCount {
		displayCount = len(servers.Results)
	}

	for i := 0; i < displayCount; i++ {
		printServerDetails(servers.Results[i], i+1)
	}

	printTruncationMessage(len(servers.Results), displayCount, "servers")

	fmt.Printf("Total baremetal servers: %d\n", len(servers.Results))
	fmt.Println("===============================")
}

func listServersWithAutoPager(client *gcore.Client) {
	fmt.Println("\n=== LIST BAREMETAL SERVERS WITH AUTO PAGER ===")

	params := cloud.BaremetalServerListParams{}
	autopager := client.Cloud.Baremetal.Servers.ListAutoPaging(context.Background(), params)

	// Collect all servers first to enable truncation
	var allServers []cloud.BaremetalServer
	for autopager.Next() {
		server := autopager.Current()
		allServers = append(allServers, server)
	}

	if err := autopager.Err(); err != nil {
		fmt.Printf("Error with auto pager: %v\n", err)
		return
	}

	displayCount := 3
	if len(allServers) < displayCount {
		displayCount = len(allServers)
	}

	for i := 0; i < displayCount; i++ {
		printServerDetails(allServers[i], i+1)
	}

	printTruncationMessage(len(allServers), displayCount, "servers")

	fmt.Printf("Total baremetal servers: %d\n", len(allServers))
	fmt.Println("===============================")
}

func rebuildServer(client *gcore.Client, serverID string, imageID string) {
	fmt.Println("\n=== REBUILD BAREMETAL SERVER ===")

	params := cloud.BaremetalServerRebuildParams{
		ImageID: gcore.String(imageID),
	}

	server, err := client.Cloud.Baremetal.Servers.RebuildAndPoll(context.Background(), serverID, params)
	if err != nil {
		fmt.Printf("Error rebuilding baremetal server: %v\n", err)
		return
	}

	fmt.Printf("Rebuilt baremetal server: ID=%s, name=%s\n", server.ID, server.Name)
	fmt.Println("===============================")
}

func printServerDetails(server cloud.BaremetalServer, index int) {
	fmt.Printf("  %d. Server: ID=%s, name=%s, status=%s\n", index, server.ID, server.Name, server.Status)
}
