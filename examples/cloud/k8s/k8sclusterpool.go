package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listPools(client *gcore.Client, clusterName string) {
	fmt.Println("\n=== LIST K8S CLUSTER POOLS ===")

	pools, err := client.Cloud.K8S.Clusters.Pools.List(
		context.Background(), clusterName, cloud.K8SClusterPoolListParams{})
	if err != nil {
		fmt.Printf("Error listing k8s cluster pools: %v\n", err)
		return
	}

	for i, p := range pools.Results {
		fmt.Printf("  %d. Pool: Name=%s, Flavor=%s, NodeCount=%d, Status=%s\n",
			i+1, p.Name, p.FlavorID, p.NodeCount, p.Status)
	}

	if len(pools.Results) == 0 {
		fmt.Println("  No pools found.")
	}

	fmt.Println("==============================")
}

func getFirstPool(client *gcore.Client, clusterName string) string {
	pools, err := client.Cloud.K8S.Clusters.Pools.List(
		context.Background(), clusterName, cloud.K8SClusterPoolListParams{})
	if err != nil {
		log.Fatalf("Error listing k8s cluster pools: %v", err)
	}
	if len(pools.Results) == 0 {
		log.Fatalf("Cluster %s has no pools", clusterName)
	}
	return pools.Results[0].Name
}

func getPool(client *gcore.Client, clusterName, poolName string) {
	fmt.Println("\n=== GET K8S CLUSTER POOL ===")

	pool, err := client.Cloud.K8S.Clusters.Pools.Get(
		context.Background(), poolName, cloud.K8SClusterPoolGetParams{
			ClusterName: clusterName,
		})
	if err != nil {
		fmt.Printf("Error getting k8s cluster pool: %v\n", err)
		return
	}

	fmt.Printf("Pool: Name=%s, Flavor=%s, NodeCount=%d, Min=%d, Max=%d, Status=%s\n",
		pool.Name, pool.FlavorID, pool.NodeCount, pool.MinNodeCount, pool.MaxNodeCount, pool.Status)
	fmt.Println("============================")
}

func checkPoolQuota(client *gcore.Client, flavorID string) {
	fmt.Println("\n=== CHECK K8S POOL QUOTA ===")

	params := cloud.K8SClusterPoolCheckQuotaParams{
		FlavorID:          flavorID,
		Name:              gcore.String("gcore-go-pool-2"),
		MinNodeCount:      gcore.Int(1),
		MaxNodeCount:      gcore.Int(2),
		BootVolumeSize:    gcore.Int(50),
		ServergroupPolicy: cloud.K8SClusterPoolCheckQuotaParamsServergroupPolicySoftAntiAffinity,
	}

	quota, err := client.Cloud.K8S.Clusters.Pools.CheckQuota(context.Background(), params)
	if err != nil {
		fmt.Printf("Error checking k8s pool quota: %v\n", err)
		return
	}

	fmt.Printf("CPU: limit=%d, requested=%d, usage=%d\n",
		quota.CPUCountLimit, quota.CPUCountRequested, quota.CPUCountUsage)
	fmt.Printf("RAM (MiB): limit=%d, requested=%d, usage=%d\n",
		quota.RamLimit, quota.RamRequested, quota.RamUsage)
	fmt.Printf("Volumes: count_limit=%d, count_usage=%d, size_limit=%d GiB\n",
		quota.VolumeCountLimit, quota.VolumeCountUsage, quota.VolumeSizeLimit)
	fmt.Println("============================")
}

func createPool(client *gcore.Client, clusterName, flavorID string) string {
	fmt.Println("\n=== CREATE K8S CLUSTER POOL ===")

	params := cloud.K8SClusterPoolNewParams{
		Name:               "gcore-go-pool-2",
		FlavorID:           flavorID,
		MinNodeCount:       1,
		MaxNodeCount:       gcore.Int(2),
		BootVolumeSize:     gcore.Int(50),
		BootVolumeType:     cloud.K8SClusterPoolNewParamsBootVolumeTypeStandard,
		AutoHealingEnabled: gcore.Bool(true),
		IsPublicIpv4:       gcore.Bool(true),
		ServergroupPolicy:  cloud.K8SClusterPoolNewParamsServergroupPolicySoftAntiAffinity,
		Labels: map[string]string{
			"pool": "gcore-go-example-2",
		},
	}

	pool, err := client.Cloud.K8S.Clusters.Pools.NewAndPoll(context.Background(), clusterName, params)
	if err != nil {
		log.Fatalf("Error creating k8s cluster pool: %v", err)
	}

	fmt.Printf("Created pool: Name=%s, Flavor=%s, NodeCount=%d\n",
		pool.Name, pool.FlavorID, pool.NodeCount)
	fmt.Println("===============================")

	return pool.Name
}

func updatePool(client *gcore.Client, clusterName, poolName string) {
	fmt.Println("\n=== UPDATE K8S CLUSTER POOL ===")

	params := cloud.K8SClusterPoolUpdateParams{
		ClusterName: clusterName,
		Labels: map[string]string{
			"pool":  "gcore-go-example-2",
			"stage": "updated",
		},
	}

	pool, err := client.Cloud.K8S.Clusters.Pools.Update(context.Background(), poolName, params)
	if err != nil {
		fmt.Printf("Error updating k8s cluster pool: %v\n", err)
		return
	}

	fmt.Printf("Updated pool: Name=%s, Labels=%v\n", pool.Name, pool.Labels)
	fmt.Println("===============================")
}

func resizePool(client *gcore.Client, clusterName, poolName string) {
	fmt.Println("\n=== RESIZE K8S CLUSTER POOL ===")

	params := cloud.K8SClusterPoolResizeParams{
		ClusterName: clusterName,
		NodeCount:   2,
	}

	pool, err := client.Cloud.K8S.Clusters.Pools.ResizeAndPoll(context.Background(), poolName, params)
	if err != nil {
		log.Fatalf("Error resizing k8s cluster pool: %v", err)
	}

	fmt.Printf("Resized pool: Name=%s, NodeCount=%d\n", pool.Name, pool.NodeCount)
	fmt.Println("===============================")
}

func deletePool(client *gcore.Client, clusterName, poolName string) {
	fmt.Println("\n=== DELETE K8S CLUSTER POOL ===")

	params := cloud.K8SClusterPoolDeleteParams{
		ClusterName: clusterName,
	}

	err := client.Cloud.K8S.Clusters.Pools.DeleteAndPoll(context.Background(), poolName, params)
	if err != nil {
		log.Fatalf("Error deleting k8s cluster pool: %v", err)
	}

	fmt.Printf("Deleted pool: Name=%s\n", poolName)
	fmt.Println("===============================")
}
