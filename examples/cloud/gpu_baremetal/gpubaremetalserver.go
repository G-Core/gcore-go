package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/packages/param"
)

func listGPUBaremetalClusterServers(client *gcore.Client, clusterID string) {
	fmt.Println("\n=== LIST GPU BAREMETAL CLUSTER SERVERS ===")

	params := cloud.GPUBaremetalClusterServerListParams{}
	servers, err := client.Cloud.GPUBaremetal.Clusters.Servers.List(context.Background(), clusterID, params)
	if err != nil {
		fmt.Printf("Error listing GPU baremetal cluster servers: %v\n", err)
		return
	}

	displayCount := 3
	if len(servers.Results) < displayCount {
		displayCount = len(servers.Results)
	}

	for i := 0; i < displayCount; i++ {
		printGPUServerDetails(servers.Results[i], i+1)
	}

	printTruncationMessage(len(servers.Results), displayCount, "servers")

	fmt.Printf("Total GPU baremetal cluster servers: %d\n", len(servers.Results))
	fmt.Println("===============================")
}

func getGPUBaremetalClusterServerConsole(client *gcore.Client, serverID string) {
	fmt.Println("\n=== GET GPU BAREMETAL CLUSTER SERVER CONSOLE ===")

	params := cloud.GPUBaremetalClusterServerGetConsoleParams{}
	console, err := client.Cloud.GPUBaremetal.Clusters.Servers.GetConsole(context.Background(), serverID, params)
	if err != nil {
		fmt.Printf("Error getting GPU baremetal server console: %v\n", err)
		return
	}

	fmt.Printf("Console URL for server %s: %s\n", serverID, console.RemoteConsole.URL)
	fmt.Printf("Console Type: %s\n", console.RemoteConsole.Type)
	fmt.Printf("Console Protocol: %s\n", console.RemoteConsole.Protocol)
	fmt.Println("===============================")
}

func deleteGPUBaremetalClusterServer(client *gcore.Client, clusterID string, serverID string) {
	fmt.Println("\n=== DELETE GPU BAREMETAL CLUSTER SERVER ===")

	params := cloud.GPUBaremetalClusterServerDeleteParams{
		ClusterID:       clusterID,
		DeleteFloatings: param.NewOpt(true),
	}

	_, err := client.Cloud.GPUBaremetal.Clusters.Servers.Delete(context.Background(), serverID, params)
	if err != nil {
		fmt.Printf("Error deleting GPU baremetal server: %v\n", err)
		return
	}

	fmt.Printf("Deleted GPU baremetal server: %s from cluster %s\n", serverID, clusterID)
	fmt.Println("===============================")
}

func printGPUServerDetails(server cloud.GPUBaremetalClusterServer, index int) {
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
