package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/packages/param"
)

func listGPUVirtualClusterServers(client *gcore.Client, clusterID string) {
	fmt.Println("\n=== LIST GPU VIRTUAL CLUSTER SERVERS ===")

	params := cloud.GPUVirtualClusterServerListParams{}
	servers, err := client.Cloud.GPUVirtual.Clusters.Servers.List(context.Background(), clusterID, params)
	if err != nil {
		fmt.Printf("Error listing GPU virtual cluster servers: %v\n", err)
		return
	}

	displayCount := min(3, len(servers.Results))

	for i := 0; i < displayCount; i++ {
		printGPUServerDetails(servers.Results[i], i+1)
	}

	printTruncationMessage(len(servers.Results), displayCount, "servers")

	fmt.Printf("Total GPU virtual cluster servers: %d\n", len(servers.Results))
	fmt.Println("===============================")
}

func deleteGPUVirtualClusterServer(client *gcore.Client, clusterID string, serverID string) {
	fmt.Println("\n=== DELETE GPU VIRTUAL CLUSTER SERVER ===")

	params := cloud.GPUVirtualClusterServerDeleteParams{
		ClusterID:      clusterID,
		AllFloatingIPs: param.NewOpt(true),
		AllVolumes:     param.NewOpt(true),
	}

	_, err := client.Cloud.GPUVirtual.Clusters.Servers.Delete(context.Background(), serverID, params)
	if err != nil {
		fmt.Printf("Error deleting GPU virtual server: %v\n", err)
		return
	}

	fmt.Printf("Deleted GPU virtual server: %s from cluster %s\n", serverID, clusterID)
	fmt.Println("===============================")
}

func printGPUServerDetails(server cloud.GPUVirtualClusterServer, index int) {
	fmt.Printf("  %d. Server: ID=%s, name=%s, status=%s\n", index, server.ID, server.Name, server.Status)
	fmt.Printf("     Flavor: %s\n", server.Flavor)
	fmt.Printf("     Image ID: %s\n", server.ImageID)
	fmt.Printf("     Created: %s\n", server.CreatedAt.Format("2006-01-02 15:04:05"))

	if len(server.IPAddresses) > 0 {
		fmt.Printf("     IP Addresses:\n")
		for _, addr := range server.IPAddresses {
			fmt.Printf("       - %s\n", addr)
		}
	}

	if len(server.SecurityGroups) > 0 {
		fmt.Printf("     Security Groups: ")
		for i, sg := range server.SecurityGroups {
			if i > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%s", sg.Name)
		}
		fmt.Println()
	}

	fmt.Println()
}
