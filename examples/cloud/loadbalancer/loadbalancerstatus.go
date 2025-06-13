package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listLoadBalancerStatuses(client *gcore.Client) {
	fmt.Println("\n=== LIST LOAD BALANCER STATUSES ===")

	params := cloud.LoadBalancerStatusListParams{}
	statuses, err := client.Cloud.LoadBalancers.Statuses.List(context.Background(), params)
	if err != nil {
		log.Fatalf("Error getting load balancer statuses: %v", err)
	}

	for i, status := range statuses.Results {
		fmt.Printf("  %d. Load Balancer Status: ID=%s, OperatingStatus=%s, ProvisioningStatus=%s\n",
			i+1, status.ID, status.OperatingStatus, status.ProvisioningStatus)
	}

	if len(statuses.Results) == 0 {
		fmt.Println("  No load balancer statuses found.")
	}

	fmt.Println("====================================")
}

func getLoadBalancerStatus(client *gcore.Client, loadBalancerID string) {
	fmt.Println("\n=== GET LOAD BALANCER STATUS ===")

	params := cloud.LoadBalancerStatusGetParams{}
	status, err := client.Cloud.LoadBalancers.Statuses.Get(context.Background(), loadBalancerID, params)
	if err != nil {
		log.Fatalf("Error getting load balancer status: %v", err)
	}

	fmt.Printf("Load Balancer Status: ID=%s, OperatingStatus=%s, ProvisioningStatus=%s\n",
		status.ID, status.OperatingStatus, status.ProvisioningStatus)
	fmt.Println("=================================")
}
