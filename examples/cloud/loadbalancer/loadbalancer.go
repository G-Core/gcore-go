package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	//apiKey := os.Getenv("GCORE_API_KEY")
	// Will use Production API URL if omitted
	//baseURL := os.Getenv("GCORE_BASE_URL")

	// TODO set cloud project ID before running
	cloudProjectID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_PROJECT_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_PROJECT_ID environment variable is required and must be a valid integer")
	}

	// TODO set cloud region ID before running
	cloudRegionID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_REGION_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_REGION_ID environment variable is required and must be a valid integer")
	}

	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectID),
		option.WithCloudRegionID(cloudRegionID),
	)

	// Create a Load Balancer and use its details for other operations
	loadBalancer := createLoadBalancer(&client)
	listLoadBalancers(&client)
	listLoadBalancersWithAutoPager(&client)
	updateLoadBalancer(&client, loadBalancer.ID)
	resizeLoadBalancer(&client, loadBalancer.ID)
	failoverLoadBalancer(&client, loadBalancer.ID)

	// Statuses
	listLoadBalancerStatuses(&client)
	getLoadBalancerByID(&client, loadBalancer.ID)
	getLoadBalancerStatus(&client, loadBalancer.ID)

	// Metrics
	getLoadBalancerMetrics(&client, loadBalancer.ID)

	deleteLoadBalancer(&client, loadBalancer.ID)
}

func createLoadBalancer(client *gcore.Client) *cloud.LoadBalancer {
	fmt.Println("\n=== CREATE LOAD BALANCER ===")

	params := cloud.LoadBalancerNewParams{
		Flavor: param.NewOpt("lb1-1-2"),
		Name:   param.NewOpt("gcore-go-example"),
		Tags:   map[string]string{"name": "gcore-go-example"},
	}

	loadBalancer, err := client.Cloud.LoadBalancers.NewAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating load balancer: %v", err)
	}

	fmt.Printf("Created Load Balancer: ID=%s, Name=%s, Status=%s\n",
		loadBalancer.ID, loadBalancer.Name, loadBalancer.ProvisioningStatus)
	fmt.Println("=============================")

	return loadBalancer
}

func listLoadBalancers(client *gcore.Client) {
	fmt.Println("\n=== LIST ALL LOAD BALANCERS ===")

	params := cloud.LoadBalancerListParams{}
	loadBalancers, err := client.Cloud.LoadBalancers.List(context.Background(), params)
	if err != nil {
		log.Fatalf("Error listing load balancers: %v", err)
	}

	for i, lb := range loadBalancers.Results {
		fmt.Printf("  %d. Load Balancer: ID=%s, Name=%s, Status=%s\n",
			i+1, lb.ID, lb.Name, lb.ProvisioningStatus)
	}

	if len(loadBalancers.Results) == 0 {
		fmt.Println("  No load balancers found.")
	}

	fmt.Println("================================")
}

func listLoadBalancersWithAutoPager(client *gcore.Client) {
	fmt.Println("\n=== LIST LOAD BALANCERS USING AUTOPAGER ===")

	count := 0

	iter := client.Cloud.LoadBalancers.ListAutoPaging(context.Background(), cloud.LoadBalancerListParams{
		Limit: gcore.Int(10), // Process 10 items per page
	})

	for iter.Next() {
		lb := iter.Current()
		count++
		fmt.Printf("  %d. Load Balancer: ID=%s, Name=%s, Status=%s\n",
			count, lb.ID, lb.Name, lb.ProvisioningStatus)
	}

	if err := iter.Err(); err != nil {
		log.Fatalf("Error iterating load balancers: %v", err)
	}

	fmt.Println("==============================================")
}

func getLoadBalancerByID(client *gcore.Client, loadBalancerID string) {
	fmt.Println("\n=== GET LOAD BALANCER BY ID ===")

	loadBalancer, err := client.Cloud.LoadBalancers.Get(context.Background(), loadBalancerID, cloud.LoadBalancerGetParams{})
	if err != nil {
		log.Fatalf("Error getting load balancer: %v", err)
	}

	fmt.Printf("Load Balancer: ID=%s, Name=%s, Status=%s, Flavor=%s\n",
		loadBalancer.ID, loadBalancer.Name, loadBalancer.ProvisioningStatus, loadBalancer.Flavor.FlavorName)
	fmt.Println("================================")
}

func updateLoadBalancer(client *gcore.Client, loadBalancerID string) {
	fmt.Println("\n=== UPDATE LOAD BALANCER ===")

	params := cloud.LoadBalancerUpdateParams{
		Name: param.NewOpt("gcore-go-example-updated"),
	}

	loadBalancer, err := client.Cloud.LoadBalancers.Update(context.Background(), loadBalancerID, params)
	if err != nil {
		log.Fatalf("Error updating load balancer: %v", err)
	}

	fmt.Printf("Updated Load Balancer: ID=%s, Name=%s\n",
		loadBalancer.ID, loadBalancer.Name)
	fmt.Println("=============================")
}

func resizeLoadBalancer(client *gcore.Client, loadBalancerID string) {
	fmt.Println("\n=== RESIZE LOAD BALANCER ===")

	params := cloud.LoadBalancerResizeParams{
		Flavor: "lb1-2-4",
	}

	_, err := client.Cloud.LoadBalancers.ResizeAndPoll(context.Background(), loadBalancerID, params)
	if err != nil {
		log.Fatalf("Error resizing load balancer: %v", err)
	}

	fmt.Printf("Resized Load Balancer: ID=%s\n", loadBalancerID)
	fmt.Println("=============================")
}

func failoverLoadBalancer(client *gcore.Client, loadBalancerID string) {
	fmt.Println("\n=== FAILOVER LOAD BALANCER ===")

	params := cloud.LoadBalancerFailoverParams{}

	_, err := client.Cloud.LoadBalancers.FailoverAndPoll(context.Background(), loadBalancerID, params)
	if err != nil {
		log.Fatalf("Error failing over load balancer: %v", err)
	}

	fmt.Printf("Failed over Load Balancer: ID=%s\n", loadBalancerID)
	fmt.Println("===============================")
}

func deleteLoadBalancer(client *gcore.Client, loadBalancerID string) {
	fmt.Println("\n=== DELETE LOAD BALANCER ===")

	params := cloud.LoadBalancerDeleteParams{}

	err := client.Cloud.LoadBalancers.DeleteAndPoll(context.Background(), loadBalancerID, params)
	if err != nil {
		log.Fatalf("Error deleting load balancer: %v", err)
	}

	fmt.Printf("Load Balancer with ID %s successfully deleted\n", loadBalancerID)
	fmt.Println("=============================")
}
