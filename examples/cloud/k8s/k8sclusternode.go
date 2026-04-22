package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listClusterNodes(client *gcore.Client, clusterName string) {
	fmt.Println("\n=== LIST K8S CLUSTER NODES ===")

	nodes, err := client.Cloud.K8S.Clusters.Nodes.List(
		context.Background(), clusterName, cloud.K8SClusterNodeListParams{})
	if err != nil {
		fmt.Printf("Error listing k8s cluster nodes: %v\n", err)
		return
	}

	for i, n := range nodes.Results {
		fmt.Printf("  %d. Node: ID=%s, Name=%s, Status=%s\n",
			i+1, n.ID, n.Name, n.Status)
	}

	if len(nodes.Results) == 0 {
		fmt.Println("  No nodes found.")
	}

	fmt.Println("==============================")
}
