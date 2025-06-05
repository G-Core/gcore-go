package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	//apiKey := os.Getenv("GCORE_API_KEY")
	// Will use Production API URL if omitted
	//baseURL := os.Getenv("GCORE_API_URL")

	client := gcore.NewClient(
	//option.WithAPIKey(apiKey),
	//option.WithBaseURL(baseURL),
	)

	// Create a project and use its ID for other operations
	projectID := createProject(&client)

	listAllProjects(&client)
	listProjectsWithFilters(&client)
	listProjectsWithAutopager(&client)
	getProjectByID(&client, projectID)
	updateProject(&client, projectID)
	deleteProject(&client, projectID)
}

func createProject(client *gcore.Client) int64 {
	fmt.Println("\n=== CREATE PROJECT ===")

	project, err := client.Cloud.Projects.New(context.Background(), cloud.ProjectNewParams{
		Name:        "gcore-go-example",
		Description: gcore.String("Project created via SDK example"),
	})
	if err != nil {
		log.Fatalf("Error creating project: %v", err)
	}

	fmt.Printf("Created Project ID: %d, Name: %s\n", project.ID, project.Name)
	fmt.Println("======================")

	return project.ID
}

func getProjectByID(client *gcore.Client, projectID int64) {
	fmt.Println("\n=== GET PROJECT BY ID ===")

	project, err := client.Cloud.Projects.Get(context.Background(), cloud.ProjectGetParams{
		ProjectID: gcore.Int(projectID),
	})
	if err != nil {
		log.Fatalf("Error getting project: %v", err)
	}

	fmt.Printf("Project ID: %d, Name: %s\n", project.ID, project.Name)
	fmt.Println("=========================")
}

func listAllProjects(client *gcore.Client) {
	fmt.Println("\n=== LIST ALL PROJECTS ===")

	projects, err := client.Cloud.Projects.List(context.Background(), cloud.ProjectListParams{})
	if err != nil {
		log.Fatalf("Error listing projects: %v", err)
	}

	for i, project := range projects.Results {
		fmt.Printf("  %d. Project ID: %d, Name: %s\n", i+1, project.ID, project.Name)
	}

	fmt.Println("=========================")
}

func listProjectsWithFilters(client *gcore.Client) {
	fmt.Println("\n=== LIST PROJECTS WITH FILTERS ===")

	projects, err := client.Cloud.Projects.List(context.Background(), cloud.ProjectListParams{
		OrderBy: cloud.ProjectListParamsOrderByNameAsc,
	})
	if err != nil {
		log.Fatalf("Error listing projects with filters: %v", err)
	}

	for i, project := range projects.Results {
		fmt.Printf("  %d. Project ID: %d, Name: %s\n", i+1, project.ID, project.Name)
	}

	fmt.Println("==================================")
}

func listProjectsWithAutopager(client *gcore.Client) {
	fmt.Println("\n=== LIST PROJECTS USING AUTOPAGER ===")

	count := 0

	iter := client.Cloud.Projects.ListAutoPaging(context.Background(), cloud.ProjectListParams{
		Limit: gcore.Int(10), // Process 10 items per page
	})

	for iter.Next() {
		project := iter.Current()
		count++
		fmt.Printf("  %d. Project ID: %d, Name: %s\n",
			count, project.ID, project.Name)
	}

	if err := iter.Err(); err != nil {
		log.Fatalf("Error iterating projects: %v", err)
	}

	fmt.Println("=====================================")
}

func updateProject(client *gcore.Client, projectID int64) {
	fmt.Println("\n=== UPDATE PROJECT ===")

	project, err := client.Cloud.Projects.Replace(context.Background(), cloud.ProjectReplaceParams{
		ProjectID:   gcore.Int(projectID),
		Name:        "gcore-go-example-updated",
		Description: gcore.String("Updated project description"),
	})
	if err != nil {
		log.Fatalf("Error updating project: %v", err)
	}

	fmt.Printf("Updated Project ID: %d, Name: %s\n", project.ID, project.Name)
	fmt.Println("======================")
}

func deleteProject(client *gcore.Client, projectID int64) {
	fmt.Println("\n=== DELETE PROJECT ===")

	_, err := client.Cloud.Projects.Delete(context.Background(), cloud.ProjectDeleteParams{
		ProjectID: gcore.Int(projectID),
	})
	if err != nil {
		log.Fatalf("Error deleting project: %v", err)
	}

	fmt.Printf("Project with ID %d successfully deleted\n", projectID)
	fmt.Println("======================")
}
