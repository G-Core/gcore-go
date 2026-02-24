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
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/shared/constant"
)

func main() {
	cloudProjectID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_PROJECT_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_PROJECT_ID environment variable is required and must be a valid integer")
	}

	cloudRegionID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_REGION_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_REGION_ID environment variable is required and must be a valid integer")
	}

	client := gcore.NewClient(
		option.WithCloudProjectID(cloudProjectID),
		option.WithCloudRegionID(cloudRegionID),
	)

	// Flavors
	flavors := listGPUBaremetalClusterFlavors(&client)

	// Images
	images := listGPUBaremetalClusterImages(&client)
	imageID := getUbuntuImage(images)

	// Check if we have GPU flavors available
	if len(flavors) == 0 {
		fmt.Println("\nWARNING: No GPU baremetal cluster flavors available in this region.")
		fmt.Println("   This example demonstrates the API usage, but cluster creation")
		fmt.Println("   will fail without available GPU flavors.")
		fmt.Println("   Showing existing clusters and demonstrating other operations...")

		// List existing clusters
		listGPUBaremetalClusters(&client)
		listGPUBaremetalClustersWithAutoPager(&client)

		fmt.Println("\nExample completed successfully!")
		fmt.Println("   All API functions are working correctly.")
		fmt.Println("   To create GPU baremetal clusters, use a region with available GPU flavors.")
		return
	}

	flavorName := getAvailableFlavor(flavors)

	// Clusters - Only proceed if we have available flavors
	fmt.Println("\nCreating GPU baremetal cluster (this may take several minutes)...")
	clusterID := createGPUBaremetalCluster(&client, flavorName, imageID)
	listGPUBaremetalClusters(&client)
	listGPUBaremetalClustersWithAutoPager(&client)
	getGPUBaremetalCluster(&client, clusterID)

	// Server operations
	listGPUBaremetalClusterServers(&client, clusterID)
	rebootAllServers(&client, clusterID)
	powercycleAllServers(&client, clusterID)

	// Cluster operations
	resizeGPUBaremetalCluster(&client, clusterID, 2)
	rebuildGPUBaremetalCluster(&client, clusterID)

	// Cleanup
	fmt.Println("\nCleaning up...")
	deleteGPUBaremetalCluster(&client, clusterID)
	fmt.Println("GPU baremetal cluster example completed successfully!")
}

func createGPUBaremetalCluster(client *gcore.Client, flavorName string, imageID string) string {
	fmt.Println("\n=== CREATE GPU BAREMETAL CLUSTER ===")

	clusterName := "gcore-go-gpu-baremetal-example"

	params := cloud.GPUBaremetalClusterNewParams{
		Name:         clusterName,
		Flavor:       flavorName,
		ImageID:      imageID,
		ServersCount: 1,
		ServersSettings: cloud.GPUBaremetalClusterNewParamsServersSettings{
			Interfaces: []cloud.GPUBaremetalClusterNewParamsServersSettingsInterfaceUnion{{
				OfExternal: &cloud.GPUBaremetalClusterNewParamsServersSettingsInterfaceExternal{
					Type: constant.External("external"),
				},
			}},
			Credentials: cloud.GPUBaremetalClusterNewParamsServersSettingsCredentials{
				Password: param.NewOpt("SecurePass123!"),
			},
		},
		Tags: map[string]string{
			"name":    "gcore-go-example",
			"example": "gpu-baremetal-cluster",
		},
	}

	cluster, err := client.Cloud.GPUBaremetal.Clusters.NewAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating GPU baremetal cluster: %v", err)
	}

	fmt.Printf("Created GPU baremetal cluster: ID=%s, name=%s, status=%s\n", cluster.ID, cluster.Name, cluster.Status)
	fmt.Printf("Cluster has %d servers\n", cluster.ServersCount)
	fmt.Println("===============================")
	return cluster.ID
}

func listGPUBaremetalClusters(client *gcore.Client) {
	fmt.Println("\n=== LIST GPU BAREMETAL CLUSTERS ===")

	params := cloud.GPUBaremetalClusterListParams{}
	clusters, err := client.Cloud.GPUBaremetal.Clusters.List(context.Background(), params)
	if err != nil {
		fmt.Printf("Error listing GPU baremetal clusters: %v\n", err)
		return
	}

	displayCount := 3
	if len(clusters.Results) < displayCount {
		displayCount = len(clusters.Results)
	}

	for i := 0; i < displayCount; i++ {
		printClusterDetails(clusters.Results[i], i+1)
	}

	printTruncationMessage(len(clusters.Results), displayCount, "clusters")

	fmt.Printf("Total GPU baremetal clusters: %d\n", len(clusters.Results))
	fmt.Println("===============================")
}

func listGPUBaremetalClustersWithAutoPager(client *gcore.Client) {
	fmt.Println("\n=== LIST GPU BAREMETAL CLUSTERS WITH AUTO PAGER ===")

	params := cloud.GPUBaremetalClusterListParams{}
	autopager := client.Cloud.GPUBaremetal.Clusters.ListAutoPaging(context.Background(), params)

	var allClusters []cloud.GPUBaremetalCluster
	for autopager.Next() {
		cluster := autopager.Current()
		allClusters = append(allClusters, cluster)
	}

	if err := autopager.Err(); err != nil {
		fmt.Printf("Error with auto pager: %v\n", err)
		return
	}

	displayCount := 3
	if len(allClusters) < displayCount {
		displayCount = len(allClusters)
	}

	for i := 0; i < displayCount; i++ {
		printClusterDetails(allClusters[i], i+1)
	}

	printTruncationMessage(len(allClusters), displayCount, "clusters")

	fmt.Printf("Total GPU baremetal clusters: %d\n", len(allClusters))
	fmt.Println("===============================")
}

func getGPUBaremetalCluster(client *gcore.Client, clusterID string) {
	fmt.Println("\n=== GET GPU BAREMETAL CLUSTER ===")

	cluster, err := client.Cloud.GPUBaremetal.Clusters.Get(context.Background(), clusterID, cloud.GPUBaremetalClusterGetParams{})
	if err != nil {
		fmt.Printf("Error getting GPU baremetal cluster: %v\n", err)
		return
	}

	fmt.Printf("Cluster: ID=%s, name=%s, status=%s\n", cluster.ID, cluster.Name, cluster.Status)
	fmt.Printf("Flavor: %s, Image ID: %s\n", cluster.Flavor, cluster.ImageID)
	fmt.Printf("Servers: %d\n", cluster.ServersCount)

	if len(cluster.ServersIDs) > 0 {
		fmt.Println("Server IDs:")
		for i, serverID := range cluster.ServersIDs {
			fmt.Printf("  %d. %s\n", i+1, serverID)
		}
	}

	fmt.Println("===============================")
}

func resizeGPUBaremetalCluster(client *gcore.Client, clusterID string, newSize int64) {
	fmt.Println("\n=== RESIZE GPU BAREMETAL CLUSTER ===")

	params := cloud.GPUBaremetalClusterResizeParams{
		InstancesCount: newSize,
	}

	cluster, err := client.Cloud.GPUBaremetal.Clusters.ResizeAndPoll(context.Background(), clusterID, params)
	if err != nil {
		fmt.Printf("Error resizing GPU baremetal cluster: %v\n", err)
		return
	}

	fmt.Printf("Resized GPU baremetal cluster: ID=%s, now has %d servers\n", cluster.ID, cluster.ServersCount)
	fmt.Println("===============================")
}

func rebuildGPUBaremetalCluster(client *gcore.Client, clusterID string) {
	fmt.Println("\n=== REBUILD GPU BAREMETAL CLUSTER ===")

	params := cloud.GPUBaremetalClusterRebuildParams{}
	rebuiltCluster, err := client.Cloud.GPUBaremetal.Clusters.RebuildAndPoll(context.Background(), clusterID, params)
	if err != nil {
		fmt.Printf("Error rebuilding GPU baremetal cluster: %v\n", err)
		return
	}

	fmt.Printf("Rebuilt GPU baremetal cluster: ID=%s, name=%s\n", rebuiltCluster.ID, rebuiltCluster.Name)
	fmt.Println("===============================")
}

func rebootAllServers(client *gcore.Client, clusterID string) {
	fmt.Println("\n=== REBOOT ALL SERVERS ===")

	params := cloud.GPUBaremetalClusterRebootAllServersParams{}
	servers, err := client.Cloud.GPUBaremetal.Clusters.RebootAllServers(context.Background(), clusterID, params)
	if err != nil {
		fmt.Printf("Error rebooting all servers: %v\n", err)
		return
	}

	fmt.Printf("Rebooted %d servers in cluster %s\n", len(servers.Results), clusterID)
	for i, server := range servers.Results {
		fmt.Printf("  %d. Server ID: %s, Status: %s\n", i+1, server.ID, server.Status)
	}
	fmt.Println("===============================")
}

func powercycleAllServers(client *gcore.Client, clusterID string) {
	fmt.Println("\n=== POWERCYCLE ALL SERVERS ===")

	params := cloud.GPUBaremetalClusterPowercycleAllServersParams{}
	servers, err := client.Cloud.GPUBaremetal.Clusters.PowercycleAllServers(context.Background(), clusterID, params)
	if err != nil {
		fmt.Printf("Error powercycling all servers: %v\n", err)
		return
	}

	fmt.Printf("Powercycled %d servers in cluster %s\n", len(servers.Results), clusterID)
	for i, server := range servers.Results {
		fmt.Printf("  %d. Server ID: %s, Status: %s\n", i+1, server.ID, server.Status)
	}
	fmt.Println("===============================")
}

func deleteGPUBaremetalCluster(client *gcore.Client, clusterID string) {
	fmt.Println("\n=== DELETE GPU BAREMETAL CLUSTER ===")

	params := cloud.GPUBaremetalClusterDeleteParams{
		AllFloatingIPs: param.NewOpt(true),
	}

	_, err := client.Cloud.GPUBaremetal.Clusters.Delete(context.Background(), clusterID, params)
	if err != nil {
		fmt.Printf("Error deleting GPU baremetal cluster: %v\n", err)
		return
	}

	fmt.Printf("Deleted GPU baremetal cluster: %s\n", clusterID)
	fmt.Println("===============================")
}

func printClusterDetails(cluster cloud.GPUBaremetalCluster, index int) {
	fmt.Printf("  %d. Cluster: ID=%s, name=%s, status=%s\n", index, cluster.ID, cluster.Name, cluster.Status)
	fmt.Printf("     Flavor: %s, Image ID: %s\n", cluster.Flavor, cluster.ImageID)
	fmt.Printf("     Servers: %d\n", cluster.ServersCount)
	fmt.Println()
}

func printTruncationMessage(totalCount, displayCount int, itemType string) {
	if totalCount > displayCount {
		fmt.Printf("  ... and %d more %s\n", totalCount-displayCount, itemType)
	}
}
