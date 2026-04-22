package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func getKubeconfig(client *gcore.Client, clusterName string) {
	fmt.Println("\n=== GET K8S KUBECONFIG ===")

	kubeconfig, err := client.Cloud.K8S.Clusters.Kubeconfig.Get(
		context.Background(), clusterName, cloud.K8SClusterKubeconfigGetParams{})
	if err != nil {
		fmt.Printf("Error getting k8s kubeconfig: %v\n", err)
		return
	}

	fmt.Printf("Kubeconfig: host=%s, created_at=%s, expires_at=%s, bytes=%d\n",
		kubeconfig.Host, kubeconfig.CreatedAt, kubeconfig.ExpiresAt, len(kubeconfig.Config))
	fmt.Println("==========================")
}
