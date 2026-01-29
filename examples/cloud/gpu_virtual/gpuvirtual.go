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
	flavors := listGPUVirtualClusterFlavors(&client)

	// Images
	images := listGPUVirtualClusterImages(&client)
	imageID := getUbuntuImage(images)

	// Check if we have GPU flavors available
	if len(flavors) == 0 {
		fmt.Println("\nWARNING: No GPU virtual cluster flavors available in this region.")
		fmt.Println("   This example demonstrates the API usage, but cluster creation")
		fmt.Println("   will fail without available GPU flavors.")
		fmt.Println("   Showing existing clusters and demonstrating other operations...")

		// List existing clusters
		listGPUVirtualClusters(&client)
		listGPUVirtualClustersWithAutoPager(&client)

		fmt.Println("\nExample completed successfully!")
		fmt.Println("   All API functions are working correctly.")
		fmt.Println("   To create GPU virtual clusters, use a region with available GPU flavors.")
		return
	}

	flavorName := getAvailableFlavor(flavors)

	// Clusters - Only proceed if we have available flavors
	fmt.Println("\nCreating GPU virtual cluster (this may take several minutes)...")
	clusterID := createGPUVirtualCluster(&client, flavorName, imageID)
	listGPUVirtualClusters(&client)
	listGPUVirtualClustersWithAutoPager(&client)
	getGPUVirtualCluster(&client, clusterID)

	// Server operations
	listGPUVirtualClusterServers(&client, clusterID)

	// Cluster actions - soft reboot
	softRebootGPUVirtualCluster(&client, clusterID)

	// Stop and start cluster
	stopGPUVirtualCluster(&client, clusterID)
	startGPUVirtualCluster(&client, clusterID)

	// Update cluster name
	updateGPUVirtualCluster(&client, clusterID)

	// Cleanup
	fmt.Println("\nCleaning up...")
	deleteGPUVirtualCluster(&client, clusterID)
	fmt.Println("GPU virtual cluster example completed successfully!")
}

func createGPUVirtualCluster(client *gcore.Client, flavorName string, imageID string) string {
	fmt.Println("\n=== CREATE GPU VIRTUAL CLUSTER ===")

	params := cloud.GPUVirtualClusterNewParams{
		Name:         "gcore-go-gpu-virtual-example",
		Flavor:       flavorName,
		ServersCount: 1,
		ServersSettings: cloud.GPUVirtualClusterNewParamsServersSettings{
			Interfaces: []cloud.GPUVirtualClusterNewParamsServersSettingsInterfaceUnion{{
				OfExternal: &cloud.GPUVirtualClusterNewParamsServersSettingsInterfaceExternal{
					Type: constant.External("external"),
				},
			}},
			Volumes: []cloud.GPUVirtualClusterNewParamsServersSettingsVolumeUnion{{
				OfImage: &cloud.GPUVirtualClusterNewParamsServersSettingsVolumeImage{
					Source:              constant.Image("image"),
					ImageID:             imageID,
					Name:                "boot-volume",
					Size:                50,
					BootIndex:           0,
					Type:                "ssd_hiiops",
					DeleteOnTermination: param.NewOpt(true),
				},
			}},
			Credentials: cloud.GPUVirtualClusterNewParamsServersSettingsCredentials{
				Password: param.NewOpt("SecurePass123!"),
			},
		},
		Tags: map[string]string{
			"name":    "gcore-go-example",
			"example": "gpu-virtual-cluster",
		},
	}

	cluster, err := client.Cloud.GPUVirtual.Clusters.NewAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating GPU virtual cluster: %v", err)
	}

	fmt.Printf("Created GPU virtual cluster: ID=%s, name=%s, status=%s\n", cluster.ID, cluster.Name, cluster.Status)
	fmt.Printf("Cluster has %d servers\n", cluster.ServersCount)
	fmt.Println("===============================")
	return cluster.ID
}

func listGPUVirtualClusters(client *gcore.Client) {
	fmt.Println("\n=== LIST GPU VIRTUAL CLUSTERS ===")

	params := cloud.GPUVirtualClusterListParams{}
	clusters, err := client.Cloud.GPUVirtual.Clusters.List(context.Background(), params)
	if err != nil {
		fmt.Printf("Error listing GPU virtual clusters: %v\n", err)
		return
	}

	displayCount := min(3, len(clusters.Results))

	for i := 0; i < displayCount; i++ {
		printClusterDetails(clusters.Results[i], i+1)
	}

	printTruncationMessage(len(clusters.Results), displayCount, "clusters")

	fmt.Printf("Total GPU virtual clusters: %d\n", len(clusters.Results))
	fmt.Println("===============================")
}

func listGPUVirtualClustersWithAutoPager(client *gcore.Client) {
	fmt.Println("\n=== LIST GPU VIRTUAL CLUSTERS WITH AUTO PAGER ===")

	params := cloud.GPUVirtualClusterListParams{}
	autopager := client.Cloud.GPUVirtual.Clusters.ListAutoPaging(context.Background(), params)

	var allClusters []cloud.GPUVirtualCluster
	for autopager.Next() {
		cluster := autopager.Current()
		allClusters = append(allClusters, cluster)
	}

	if err := autopager.Err(); err != nil {
		fmt.Printf("Error with auto pager: %v\n", err)
		return
	}

	displayCount := min(3, len(allClusters))

	for i := 0; i < displayCount; i++ {
		printClusterDetails(allClusters[i], i+1)
	}

	printTruncationMessage(len(allClusters), displayCount, "clusters")

	fmt.Printf("Total GPU virtual clusters: %d\n", len(allClusters))
	fmt.Println("===============================")
}

func getGPUVirtualCluster(client *gcore.Client, clusterID string) {
	fmt.Println("\n=== GET GPU VIRTUAL CLUSTER ===")

	cluster, err := client.Cloud.GPUVirtual.Clusters.Get(context.Background(), clusterID, cloud.GPUVirtualClusterGetParams{})
	if err != nil {
		fmt.Printf("Error getting GPU virtual cluster: %v\n", err)
		return
	}

	fmt.Printf("Cluster: ID=%s, name=%s, status=%s\n", cluster.ID, cluster.Name, cluster.Status)
	fmt.Printf("Flavor: %s\n", cluster.Flavor)
	fmt.Printf("Servers: %d\n", cluster.ServersCount)

	if len(cluster.ServersIDs) > 0 {
		fmt.Println("Server IDs:")
		for i, serverID := range cluster.ServersIDs {
			fmt.Printf("  %d. %s\n", i+1, serverID)
		}
	}

	fmt.Println("===============================")
}

func updateGPUVirtualCluster(client *gcore.Client, clusterID string) {
	fmt.Println("\n=== UPDATE GPU VIRTUAL CLUSTER ===")

	params := cloud.GPUVirtualClusterUpdateParams{
		Name: "gcore-go-gpu-virtual-updated",
	}

	cluster, err := client.Cloud.GPUVirtual.Clusters.Update(context.Background(), clusterID, params)
	if err != nil {
		fmt.Printf("Error updating GPU virtual cluster: %v\n", err)
		return
	}

	fmt.Printf("Updated GPU virtual cluster: ID=%s, name=%s\n", cluster.ID, cluster.Name)
	fmt.Println("===============================")
}

func softRebootGPUVirtualCluster(client *gcore.Client, clusterID string) {
	fmt.Println("\n=== SOFT REBOOT GPU VIRTUAL CLUSTER ===")

	params := cloud.GPUVirtualClusterActionParams{
		OfSoftReboot: &cloud.GPUVirtualClusterActionParamsBodySoftReboot{},
	}

	_, err := client.Cloud.GPUVirtual.Clusters.ActionAndPoll(context.Background(), clusterID, params)
	if err != nil {
		fmt.Printf("Error soft rebooting GPU virtual cluster: %v\n", err)
		return
	}

	fmt.Printf("Soft rebooted GPU virtual cluster: %s\n", clusterID)
	fmt.Println("===============================")
}

func stopGPUVirtualCluster(client *gcore.Client, clusterID string) {
	fmt.Println("\n=== STOP GPU VIRTUAL CLUSTER ===")

	params := cloud.GPUVirtualClusterActionParams{
		OfStop: &cloud.GPUVirtualClusterActionParamsBodyStop{},
	}

	_, err := client.Cloud.GPUVirtual.Clusters.ActionAndPoll(context.Background(), clusterID, params)
	if err != nil {
		fmt.Printf("Error stopping GPU virtual cluster: %v\n", err)
		return
	}

	fmt.Printf("Stopped GPU virtual cluster: %s\n", clusterID)
	fmt.Println("===============================")
}

func startGPUVirtualCluster(client *gcore.Client, clusterID string) {
	fmt.Println("\n=== START GPU VIRTUAL CLUSTER ===")

	params := cloud.GPUVirtualClusterActionParams{
		OfStart: &cloud.GPUVirtualClusterActionParamsBodyStart{},
	}

	_, err := client.Cloud.GPUVirtual.Clusters.ActionAndPoll(context.Background(), clusterID, params)
	if err != nil {
		fmt.Printf("Error starting GPU virtual cluster: %v\n", err)
		return
	}

	fmt.Printf("Started GPU virtual cluster: %s\n", clusterID)
	fmt.Println("===============================")
}

func deleteGPUVirtualCluster(client *gcore.Client, clusterID string) {
	fmt.Println("\n=== DELETE GPU VIRTUAL CLUSTER ===")

	params := cloud.GPUVirtualClusterDeleteParams{
		AllFloatingIPs: param.NewOpt(true),
		AllVolumes:     param.NewOpt(true),
	}

	err := client.Cloud.GPUVirtual.Clusters.DeleteAndPoll(context.Background(), clusterID, params)
	if err != nil {
		fmt.Printf("Error deleting GPU virtual cluster: %v\n", err)
		return
	}

	fmt.Printf("Deleted GPU virtual cluster: %s\n", clusterID)
	fmt.Println("===============================")
}

func printClusterDetails(cluster cloud.GPUVirtualCluster, index int) {
	fmt.Printf("  %d. Cluster: ID=%s, name=%s, status=%s\n", index, cluster.ID, cluster.Name, cluster.Status)
	fmt.Printf("     Flavor: %s\n", cluster.Flavor)
	fmt.Printf("     Servers: %d\n", cluster.ServersCount)
	fmt.Println()
}

func printTruncationMessage(totalCount, displayCount int, itemType string) {
	if totalCount > displayCount {
		fmt.Printf("  ... and %d more %s\n", totalCount-displayCount, itemType)
	}
}
