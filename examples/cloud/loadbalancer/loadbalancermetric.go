package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func getLoadBalancerMetrics(client *gcore.Client, loadBalancerID string) {
	fmt.Println("\n=== GET LOAD BALANCER METRICS ===")

	params := cloud.LoadBalancerMetricListParams{
		TimeInterval: 1,
		TimeUnit:     cloud.InstanceMetricsTimeUnitHour,
	}

	metrics, err := client.Cloud.LoadBalancers.Metrics.List(context.Background(), loadBalancerID, params)
	if err != nil {
		log.Fatalf("Error getting load balancer metrics: %v", err)
	}

	fmt.Printf("Load Balancer Metrics for ID %s:\n", loadBalancerID)
	for i, metric := range metrics.Results {
		fmt.Printf("  %d. Metric: CPUUtil=%.2f%%, MemoryUtil=%.2f%%, Time=%s\n",
			i+1, metric.CPUUtil, metric.MemoryUtil, metric.Time)
	}

	if len(metrics.Results) == 0 {
		fmt.Println("  No metrics found.")
	}

	fmt.Println("==================================")
}
