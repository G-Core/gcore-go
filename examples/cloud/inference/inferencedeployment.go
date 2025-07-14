package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func createDeployment(client *gcore.Client, flavorName string, regionID int64) string {
	fmt.Println("\n=== CREATE DEPLOYMENT ===")

	container := cloud.InferenceDeploymentNewParamsContainer{
		RegionID: regionID,
		Scale: cloud.InferenceDeploymentNewParamsContainerScale{
			Min:             1,
			Max:             3,
			CooldownPeriod:  gcore.Int(60),
			PollingInterval: gcore.Int(30),
			Triggers: cloud.InferenceDeploymentNewParamsContainerScaleTriggers{
				CPU: cloud.InferenceDeploymentNewParamsContainerScaleTriggersCPU{
					Threshold: 80,
				},
			},
		},
	}

	params := cloud.InferenceDeploymentNewParams{
		Name:          "gcore-go-example",
		Image:         "nginx:latest",
		FlavorName:    flavorName,
		ListeningPort: 80,
		Containers:    []cloud.InferenceDeploymentNewParamsContainer{container},
	}

	deployment, err := client.Cloud.Inference.Deployments.NewAndPoll(context.Background(), params)
	if err != nil {
		fmt.Printf("Error creating deployment: %v\n", err)
		return ""
	}

	fmt.Printf("Created deployment: %s, status: %s\n", deployment.Name, deployment.Status)
	fmt.Println("========================")
	return deployment.Name
}

func listDeployments(client *gcore.Client) {
	fmt.Println("\n=== LIST DEPLOYMENTS ===")

	iter := client.Cloud.Inference.Deployments.ListAutoPaging(context.Background(), cloud.InferenceDeploymentListParams{})
	count := 0
	for iter.Next() {
		count++
		deployment := iter.Current()
		fmt.Printf("  %d. Deployment: %s, status: %s\n", count, deployment.Name, deployment.Status)
	}

	if err := iter.Err(); err != nil {
		fmt.Printf("Error iterating deployments: %v\n", err)
		return
	}

	fmt.Println("========================")
}

func getDeployment(client *gcore.Client, deploymentName string) {
	fmt.Println("\n=== GET DEPLOYMENT ===")

	deployment, err := client.Cloud.Inference.Deployments.Get(context.Background(), deploymentName, cloud.InferenceDeploymentGetParams{})
	if err != nil {
		fmt.Printf("Error getting deployment: %v\n", err)
		return
	}

	fmt.Printf("Deployment: %s, status: %s\n", deployment.Name, deployment.Status)
	fmt.Println("========================")
}

func updateDeployment(client *gcore.Client, deploymentName string) {
	fmt.Println("\n=== UPDATE DEPLOYMENT ===")

	params := cloud.InferenceDeploymentUpdateParams{
		Description: gcore.String("Updated description"),
	}

	deployment, err := client.Cloud.Inference.Deployments.UpdateAndPoll(context.Background(), deploymentName, params)
	if err != nil {
		fmt.Printf("Error updating deployment: %v\n", err)
		return
	}

	fmt.Printf("Updated deployment: %s\n", deployment.Name)
	fmt.Println("========================")
}

func startDeployment(client *gcore.Client, deploymentName string) {
	fmt.Println("\n=== START DEPLOYMENT ===")

	params := cloud.InferenceDeploymentStartParams{}
	err := client.Cloud.Inference.Deployments.Start(context.Background(), deploymentName, params)
	if err != nil {
		fmt.Printf("Error starting deployment: %v\n", err)
		return
	}

	fmt.Printf("Started deployment: %s\n", deploymentName)
	fmt.Println("========================")
}

func stopDeployment(client *gcore.Client, deploymentName string) {
	fmt.Println("\n=== STOP DEPLOYMENT ===")

	params := cloud.InferenceDeploymentStopParams{}
	err := client.Cloud.Inference.Deployments.Stop(context.Background(), deploymentName, params)
	if err != nil {
		fmt.Printf("Error stopping deployment: %v\n", err)
		return
	}

	fmt.Printf("Stopped deployment: %s\n", deploymentName)
	fmt.Println("========================")
}

func deleteDeployment(client *gcore.Client, deploymentName string) {
	fmt.Println("\n=== DELETE DEPLOYMENT ===")

	params := cloud.InferenceDeploymentDeleteParams{}
	err := client.Cloud.Inference.Deployments.DeleteAndPoll(context.Background(), deploymentName, params)
	if err != nil {
		fmt.Printf("Error deleting deployment: %v\n", err)
		return
	}

	fmt.Printf("Deleted deployment: %s\n", deploymentName)
	fmt.Println("========================")
}
