package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	//apiKey := os.Getenv("GCORE_API_KEY")
	// Will use Production API URL if omitted
	//baseURL := os.Getenv("GCORE_BASE_URL")

	client := gcore.NewClient(
	//option.WithAPIKey(apiKey),
	//option.WithBaseURL(baseURL),
	)

	getTaskByID(&client)
	listTasks(&client)
	listTasksWithAutopager(&client)
}

func getTaskByID(client *gcore.Client) {
	// This example assumes you have a task ID. Replace with a valid ID from your account.
	// You might need to run listTasks first to find one.
	taskID := "your-task-id-here" // TODO: Replace with an actual task ID

	fmt.Printf("\n=== GET TASK BY ID (%s) ===\n", taskID)

	ctx := context.Background()
	task, err := client.Cloud.Tasks.Get(ctx, taskID)
	if err != nil {
		// It's common to not have a specific task ID readily available,
		// so log as info instead of fatal if not found or invalid.
		log.Printf("Info: Could not get task '%s': %v", taskID, err)
		fmt.Println("============================")
		return // Don't proceed if task couldn't be fetched
	}

	fmt.Printf("Task ID: %s, Type: %s, State: %s, Created: %s\n",
		task.ID, task.TaskType, task.State, task.CreatedOn)
	fmt.Println("============================")
}

func listTasks(client *gcore.Client) {
	// Get project ID from environment variable GCORE_CLOUD_PROJECT_ID, defaulting to 1 if not set or invalid
	projectID := int64(1)
	if idStr := os.Getenv("GCORE_CLOUD_PROJECT_ID"); idStr != "" {
		if id, err := strconv.ParseInt(idStr, 10, 64); err == nil {
			projectID = id
		} else {
			log.Printf("Warning: Invalid GCORE_CLOUD_PROJECT_ID value '%s', using default 1. Error: %v", idStr, err)
		}
	}
	// Get region ID from environment variable GCORE_CLOUD_REGION_ID, defaulting to 1 if not set or invalid
	regionID := int64(1)
	if idStr := os.Getenv("GCORE_CLOUD_REGION_ID"); idStr != "" {
		if id, err := strconv.ParseInt(idStr, 10, 64); err == nil {
			regionID = id
		} else {
			log.Printf("Warning: Invalid GCORE_CLOUD_REGION_ID value '%s', using default 1. Error: %v", idStr, err)
		}
	}

	fmt.Printf("\n=== LIST ALL TASKS (Project: %d, Region: %d) ===\n", projectID, regionID)

	ctx := context.Background()
	tasksPage, err := client.Cloud.Tasks.List(ctx, cloud.TaskListParams{
		// Pass project_id and region_id as slices, as expected by the API
		ProjectID: []int64{projectID},
		RegionID:  []int64{regionID},
	})
	if err != nil {
		log.Fatalf("Error listing tasks: %v", err)
	}

	count := 0
	fmt.Println("Results:")
	if tasksPage != nil {
		for _, task := range tasksPage.Results {
			fmt.Printf(
				"- Task ID: %s, Type: %s, State: %s, Created: %s\n",
				task.ID, task.TaskType, task.State, task.CreatedOn,
			)
			count++
		}

		if count == 0 {
			fmt.Println("No tasks found.")
		}
	} else {
		fmt.Println("Could not retrieve tasks.") // Handle case where tasksPage might be nil (though unlikely on success)
	}

	fmt.Println("=================================================")
}

func listTasksWithAutopager(client *gcore.Client) {
	// Get project ID from environment variable GCORE_CLOUD_PROJECT_ID, defaulting to 1 if not set or invalid
	projectID := int64(1)
	if idStr := os.Getenv("GCORE_CLOUD_PROJECT_ID"); idStr != "" {
		if id, err := strconv.ParseInt(idStr, 10, 64); err == nil {
			projectID = id
		} else {
			log.Printf("Warning: Invalid GCORE_CLOUD_PROJECT_ID value '%s', using default 1. Error: %v", idStr, err)
		}
	}
	// Get region ID from environment variable GCORE_CLOUD_REGION_ID, defaulting to 1 if not set or invalid
	regionID := int64(1)
	if idStr := os.Getenv("GCORE_CLOUD_REGION_ID"); idStr != "" {
		if id, err := strconv.ParseInt(idStr, 10, 64); err == nil {
			regionID = id
		} else {
			log.Printf("Warning: Invalid GCORE_CLOUD_REGION_ID value '%s', using default 1. Error: %v", idStr, err)
		}
	}

	fmt.Printf("\n=== LIST TASKS USING AUTOPAGER (Project: %d, Region: %d) ===\n", projectID, regionID)

	ctx := context.Background()
	count := 0
	iter := client.Cloud.Tasks.ListAutoPaging(ctx, cloud.TaskListParams{
		// Pass project_id and region_id as slices
		ProjectID: []int64{projectID},
		RegionID:  []int64{regionID},
		Limit:     gcore.Int(10), // Process 10 items per page
	})

	fmt.Println("Results:")
	for iter.Next() {
		task := iter.Current()
		count++
		fmt.Printf(
			"  %d. Task ID: %s, Type: %s, State: %s, Created: %s\n",
			count, task.ID, task.TaskType, task.State, task.CreatedOn,
		)
	}

	if err := iter.Err(); err != nil {
		log.Fatalf("Error iterating tasks: %v", err)
	}

	if count == 0 {
		fmt.Println("No tasks found.")
	}

	fmt.Println("======================================================")
}
