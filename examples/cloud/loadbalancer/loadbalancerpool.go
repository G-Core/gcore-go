package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/packages/param"
)

func createPool(client *gcore.Client, loadBalancerID string) *cloud.LoadBalancerPool {
	fmt.Println("\n=== CREATE LOAD BALANCER POOL ===")

	params := cloud.LoadBalancerPoolNewParams{
		Name:           "gcore-go-example-pool",
		LbAlgorithm:    cloud.LbAlgorithmRoundRobin,
		Protocol:       cloud.LbPoolProtocolHTTP,
		LoadBalancerID: param.NewOpt(loadBalancerID),
	}

	pool, err := client.Cloud.LoadBalancers.Pools.NewAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating load balancer pool: %v", err)
	}

	fmt.Printf("Created Pool: ID=%s, Name=%s, Protocol=%s, Algorithm=%s\n",
		pool.ID, pool.Name, pool.Protocol, pool.LbAlgorithm)
	fmt.Println("=================================")

	return pool
}

func listPools(client *gcore.Client) {
	fmt.Println("\n=== LIST LOAD BALANCER POOLS ===")

	params := cloud.LoadBalancerPoolListParams{}
	pools, err := client.Cloud.LoadBalancers.Pools.List(context.Background(), params)
	if err != nil {
		log.Fatalf("Error listing pools: %v", err)
	}

	for i, pool := range pools.Results {
		fmt.Printf("  %d. Pool: ID=%s, Name=%s, Protocol=%s\n",
			i+1, pool.ID, pool.Name, pool.Protocol)
	}

	if len(pools.Results) == 0 {
		fmt.Println("  No pools found.")
	}

	fmt.Println("================================")
}

func deletePool(client *gcore.Client, poolID string) {
	fmt.Println("\n=== DELETE LOAD BALANCER POOL ===")

	params := cloud.LoadBalancerPoolDeleteParams{}

	err := client.Cloud.LoadBalancers.Pools.DeleteAndPoll(context.Background(), poolID, params)
	if err != nil {
		log.Fatalf("Error deleting pool: %v", err)
	}

	fmt.Printf("Pool with ID %s successfully deleted\n", poolID)
	fmt.Println("=================================")
}
