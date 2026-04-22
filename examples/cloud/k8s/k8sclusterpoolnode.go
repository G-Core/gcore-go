package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listPoolNodes(client *gcore.Client, clusterName, poolName string) {
	fmt.Println("\n=== LIST K8S POOL NODES ===")

	nodes, err := client.Cloud.K8S.Clusters.Pools.Nodes.List(
		context.Background(), poolName, cloud.K8SClusterPoolNodeListParams{
			ClusterName: clusterName,
		})
	if err != nil {
		fmt.Printf("Error listing k8s pool nodes: %v\n", err)
		return
	}

	for i, n := range nodes.Results {
		fmt.Printf("  %d. Node: ID=%s, Name=%s, Status=%s\n",
			i+1, n.ID, n.Name, n.Status)
	}

	if len(nodes.Results) == 0 {
		fmt.Println("  No pool nodes found.")
	}

	fmt.Println("===========================")
}
