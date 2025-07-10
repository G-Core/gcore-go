package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listDeploymentLogs(client *gcore.Client, deploymentName string) {
	fmt.Println("\n=== LIST DEPLOYMENT LOGS ===")

	iter := client.Cloud.Inference.Deployments.Logs.ListAutoPaging(context.Background(), deploymentName, cloud.InferenceDeploymentLogListParams{})
	count := 0
	for iter.Next() {
		count++
		log := iter.Current()
		fmt.Printf("  %d. Log: %s at %s\n", count, log.Message, log.Time.Format("2006-01-02 15:04:05"))
		if count >= 10 { // Limit to first 10 logs
			break
		}
	}

	if err := iter.Err(); err != nil {
		fmt.Printf("Error iterating deployment logs: %v\n", err)
		return
	}

	if count == 0 {
		fmt.Println("  No logs found")
	}

	fmt.Println("========================")
}
