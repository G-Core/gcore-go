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
	"github.com/G-Core/gcore-go/shared/constant"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	//apiKey := os.Getenv("GCORE_API_KEY")
	// Will use Production API URL if omitted
	//baseURL := os.Getenv("GCORE_API_URL")

	// TODO set cloud project ID before running
	cloudProjectId, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_PROJECT_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_PROJECT_ID environment variable is required and must be a valid integer")
	}

	// TODO set cloud region ID before running
	cloudRegionId, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_REGION_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_REGION_ID environment variable is required and must be a valid integer")
	}

	// TODO set cloud network ID before running
	networkID := os.Getenv("GCORE_CLOUD_NETWORK_ID")
	if networkID == "" {
		log.Fatalf("GCORE_CLOUD_NETWORK_ID environment variable is required")
	}

	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectId),
		option.WithCloudRegionID(cloudRegionId),
	)

	// Create a File Share and use its details for other operations
	fileShare := createFileShare(&client, networkID)
	listAllFileShares(&client)
	listFileSharesWithAutopager(&client)
	getFileShareByID(&client, fileShare.ID)
	updateFileShare(&client, fileShare.ID, "gcore-go-example-updated")
	resizeFileShare(&client, fileShare.ID, 2)
	deleteFileShare(&client, fileShare.ID)
}

func createFileShare(client *gcore.Client, networkID string) *cloud.FileShare {
	fmt.Println("\n=== CREATE FILE SHARE ===")

	params := cloud.FileShareNewParams{
		OfCreateStandardFileShareSerializer: &cloud.FileShareNewParamsBodyCreateStandardFileShareSerializer{
			Name: "gcore-go-example",
			Network: cloud.FileShareNewParamsBodyCreateStandardFileShareSerializerNetwork{
				NetworkID: networkID,
			},
			Size:       1,
			VolumeType: "default_share_type",
			Protocol:   constant.ValueOf[constant.Nfs](),
		},
	}

	taskList, err := client.Cloud.FileShares.New(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating file share: %v", err)
	}

	task, err := client.Cloud.Tasks.Poll(context.Background(), taskList.Tasks[0])
	if err != nil {
		log.Fatalf("Error polling task: %v", err)
	}

	fileShare, err := client.Cloud.FileShares.Get(context.Background(), task.CreatedResources.FileShares[0], cloud.FileShareGetParams{})
	if err != nil {
		log.Fatalf("Error getting created file share: %v", err)
	}

	fmt.Printf("Created File Share: ID=%s, Name=%s, Size=%d GiB, Status=%s\n",
		fileShare.ID, fileShare.Name, fileShare.Size, fileShare.Status)
	fmt.Println("=========================")

	return fileShare
}

func listAllFileShares(client *gcore.Client) {
	fmt.Println("\n=== LIST ALL FILE SHARES ===")

	fileShares, err := client.Cloud.FileShares.List(context.Background(), cloud.FileShareListParams{})
	if err != nil {
		log.Fatalf("Error listing file shares: %v", err)
	}

	for i, fs := range fileShares.Results {
		fmt.Printf("  %d. File Share: ID=%s, Name=%s, Size=%d GiB, Status=%s\n",
			i+1, fs.ID, fs.Name, fs.Size, fs.Status)
	}

	fmt.Println("============================")
}

func listFileSharesWithAutopager(client *gcore.Client) {
	fmt.Println("\n=== LIST FILE SHARES USING AUTOPAGER ===")

	count := 0

	iter := client.Cloud.FileShares.ListAutoPaging(context.Background(), cloud.FileShareListParams{
		Limit: gcore.Int(10), // Process 10 items per page
	})

	for iter.Next() {
		fs := iter.Current()
		count++
		fmt.Printf("  %d. File Share: ID=%s, Name=%s, Size=%d GiB, Status=%s\n",
			count, fs.ID, fs.Name, fs.Size, fs.Status)
	}

	if err := iter.Err(); err != nil {
		log.Fatalf("Error iterating file shares: %v", err)
	}

	fmt.Println("=========================================")
}

func getFileShareByID(client *gcore.Client, fileShareID string) {
	fmt.Println("\n=== GET FILE SHARE BY ID ===")

	fileShare, err := client.Cloud.FileShares.Get(context.Background(), fileShareID, cloud.FileShareGetParams{})
	if err != nil {
		log.Fatalf("Error getting file share: %v", err)
	}

	fmt.Printf("File Share: ID=%s, Name=%s, Size=%d GiB, Status=%s, Protocol=%s\n",
		fileShare.ID, fileShare.Name, fileShare.Size, fileShare.Status, fileShare.Protocol)
	fmt.Println("============================")
}

func updateFileShare(client *gcore.Client, fileShareID, newName string) {
	fmt.Println("\n=== UPDATE FILE SHARE ===")

	params := cloud.FileShareUpdateParams{
		Name: newName,
	}

	fileShare, err := client.Cloud.FileShares.Update(context.Background(), fileShareID, params)
	if err != nil {
		log.Fatalf("Error updating file share: %v", err)
	}

	fmt.Printf("Updated File Share: ID=%s, Name=%s, Size=%d GiB, Status=%s\n",
		fileShare.ID, fileShare.Name, fileShare.Size, fileShare.Status)
	fmt.Println("=========================")
}

func resizeFileShare(client *gcore.Client, fileShareID string, newSize int64) {
	fmt.Println("\n=== RESIZE FILE SHARE ===")

	params := cloud.FileShareResizeParams{
		Size: newSize,
	}

	taskList, err := client.Cloud.FileShares.Resize(context.Background(), fileShareID, params)
	if err != nil {
		log.Fatalf("Error resizing file share: %v", err)
	}

	taskID := taskList.Tasks[0]
	_, err = client.Cloud.Tasks.Poll(context.Background(), taskID)
	if err != nil {
		log.Fatalf("Error polling resize task: %v", err)
	}

	fmt.Printf("File Share %s successfully resized to %d GiB\n", fileShareID, newSize)
	fmt.Println("=========================")
}

func deleteFileShare(client *gcore.Client, fileShareID string) {
	fmt.Println("\n=== DELETE FILE SHARE ===")

	taskList, err := client.Cloud.FileShares.Delete(context.Background(), fileShareID, cloud.FileShareDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting file share: %v", err)
	}

	if len(taskList.Tasks) > 0 {
		_, err := client.Cloud.Tasks.Poll(context.Background(), taskList.Tasks[0])
		if err != nil {
			log.Fatalf("Error polling delete task: %v", err)
		}
	}

	fmt.Printf("File Share %s successfully deleted, Tasks: %v\n", fileShareID, taskList.Tasks)
	fmt.Println("=========================")
}
