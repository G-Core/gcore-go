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
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	//apiKey := os.Getenv("GCORE_API_KEY")
	// Will use Production API URL if omitted
	//baseURL := os.Getenv("GCORE_API_URL")

	// TODO set cloud project ID before running
	cloudProjectID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_PROJECT_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_PROJECT_ID environment variable is required and must be a valid integer")
	}

	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectID),
	)

	// Create a router and use its ID for other operations
	routerID := createRouter(&client)
	listAllRouters(&client)
	listRoutersWithFilters(&client)
	listRoutersWithAutopager(&client)
	getRouterByID(&client, routerID)
	updateRouter(&client, routerID)
	deleteRouter(&client, routerID)
}

func createRouter(client *gcore.Client) string {
	fmt.Println("\n=== CREATE ROUTER ===")

	taskIDList, err := client.Cloud.Networks.Routers.New(context.Background(), cloud.NetworkRouterNewParams{
		Name: "gcore-go-example",
	})
	if err != nil {
		log.Fatalf("Error creating router: %v", err)
	}

	task, err := client.Cloud.Tasks.Poll(context.Background(), taskIDList.Tasks[0])
	if err != nil {
		log.Fatalf("Error polling task: %v", err)
	}

	routerID := task.CreatedResources.Routers[0]
	fmt.Printf("Created Router ID: %s\n", routerID)
	fmt.Println("======================")

	return routerID
}

func listAllRouters(client *gcore.Client) {
	fmt.Println("\n=== LIST ALL ROUTERS ===")

	routers, err := client.Cloud.Networks.Routers.List(context.Background(), cloud.NetworkRouterListParams{})
	if err != nil {
		log.Fatalf("Error listing routers: %v", err)
	}

	for i, router := range routers.Results {
		fmt.Printf("  %d. Router ID: %s, Name: %s\n", i+1, router.ID, router.Name)
	}

	fmt.Println("========================")
}

func listRoutersWithFilters(client *gcore.Client) {
	fmt.Println("\n=== LIST ROUTERS WITH FILTERS ===")

	routers, err := client.Cloud.Networks.Routers.List(context.Background(), cloud.NetworkRouterListParams{
		Limit: gcore.Int(10),
	})
	if err != nil {
		log.Fatalf("Error listing routers with filters: %v", err)
	}

	for i, router := range routers.Results {
		fmt.Printf("  %d. Router ID: %s, Name: %16s, Status: %s\n", i+1, router.ID, router.Name,
			router.Status)
	}

	fmt.Println("=================================")
}

func listRoutersWithAutopager(client *gcore.Client) {
	fmt.Println("\n=== LIST ROUTERS USING AUTOPAGER ===")

	count := 0

	iter := client.Cloud.Networks.Routers.ListAutoPaging(context.Background(), cloud.NetworkRouterListParams{
		Limit: gcore.Int(10), // Process 10 items per page
	})

	for iter.Next() {
		router := iter.Current()
		count++
		fmt.Printf("  %d. Router ID: %s, Name: %s\n", count, router.ID, router.Name)
	}

	if err := iter.Err(); err != nil {
		log.Fatalf("Error iterating routers: %v", err)
	}

	fmt.Println("====================================")
}

func getRouterByID(client *gcore.Client, routerID string) {
	fmt.Println("\n=== GET ROUTER BY ID ===")

	router, err := client.Cloud.Networks.Routers.Get(context.Background(), routerID, cloud.NetworkRouterGetParams{})
	if err != nil {
		log.Fatalf("Error getting router: %v", err)
	}

	fmt.Printf("Router ID: %s, Name: %s\n", router.ID, router.Name)
	fmt.Println("========================")
}

func updateRouter(client *gcore.Client, routerID string) {
	fmt.Println("\n=== UPDATE ROUTER ===")

	params := cloud.NetworkRouterUpdateParams{
		Name: gcore.String("gcore-go-example-updated"),
	}

	router, err := client.Cloud.Networks.Routers.Update(context.Background(), routerID, params)
	if err != nil {
		log.Fatalf("Error updating router: %v", err)
	}

	fmt.Printf("Updated Router ID: %s, New Name: %s\n", router.ID, router.Name)
	fmt.Println("=====================")
}

func deleteRouter(client *gcore.Client, routerID string) {
	fmt.Println("\n=== DELETE ROUTER ===")

	taskIDList, err := client.Cloud.Networks.Routers.Delete(context.Background(), routerID, cloud.NetworkRouterDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting router: %v", err)
	}

	_, err = client.Cloud.Tasks.Poll(context.Background(), taskIDList.Tasks[0])
	if err != nil {
		log.Fatalf("Error polling delete task: %v", err)
	}

	fmt.Printf("Router with ID %s successfully deleted\n", routerID)
	fmt.Println("=====================")
}
