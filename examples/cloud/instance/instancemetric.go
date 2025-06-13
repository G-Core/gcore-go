package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listMetrics(client *gcore.Client, instanceID string) {
	fmt.Println("\n=== LIST METRICS ===")

	params := cloud.InstanceMetricListParams{
		TimeInterval: 1,
		TimeUnit:     "hour",
	}

	metrics, err := client.Cloud.Instances.Metrics.List(context.Background(), instanceID, params)
	if err != nil {
		fmt.Printf("Error listing metrics: %v\n", err)
		return
	}

	fmt.Printf("Metrics for instance: %d entries\n", len(metrics.Results))
	displayCount := 3
	if len(metrics.Results) < displayCount {
		displayCount = len(metrics.Results)
	}
	for i := 0; i < displayCount; i++ {
		metric := metrics.Results[i]
		fmt.Printf("  %d. Metric: %+v\n", i+1, metric)
	}
	if len(metrics.Results) > displayCount {
		fmt.Printf("  ... and %d more metrics\n", len(metrics.Results)-displayCount)
	}
	fmt.Println("====================")
}
