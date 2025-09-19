package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/storage"
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

	listStorages(&client)
	storageID := createStorage(&client)
	getStorage(&client, storageID)
	updateStorage(&client, storageID)
	listStorages(&client)
	getStorage(&client, storageID)
	deleteStorage(&client, storageID)
}

func createStorage(client *gcore.Client) int64 {
	fmt.Println("\n=== CREATE STORAGE ===")

	name := fmt.Sprintf("example-s3-storage-%d", time.Now().Unix())

	params := storage.StorageNewParams{
		Name:     name,
		Type:     storage.StorageNewParamsTypeS3,
		Location: storage.StorageNewParamsLocationAms,
	}

	newStorage, err := client.Storage.New(context.Background(), params)
	if err != nil {
		fmt.Printf("Error creating %s storage: %v\n", storage.StorageNewParamsTypeS3, err)
		log.Fatalf("Failed to create storage")
	}

	fmt.Printf("Created Storage: ID=%d, Name=%s, Type=%s, Location=%s\n",
		newStorage.ID, newStorage.Name, newStorage.Type, newStorage.Location)
	fmt.Printf("Storage address: %s\n", newStorage.Address)

	// Display S3 credentials
	if newStorage.Credentials.S3.AccessKey != "" {
		fmt.Printf("S3 Access Key: %s\n", newStorage.Credentials.S3.AccessKey)
		fmt.Printf("S3 Secret Key: %s\n", newStorage.Credentials.S3.SecretKey)
	}

	fmt.Println("======================")
	return newStorage.ID
}

func getStorage(client *gcore.Client, storageID int64) {
	fmt.Println("\n=== GET STORAGE ===")

	storageDetails, err := client.Storage.Get(context.Background(), storageID)
	if err != nil {
		fmt.Printf("Error getting storage details: %v\n", err)
		fmt.Println("==================")
		return
	}

	fmt.Printf("Storage: ID=%d, Name=%s, Type=%s, Location=%s, Status=%s\n",
		storageDetails.ID, storageDetails.Name, storageDetails.Type, storageDetails.Location, storageDetails.ProvisioningStatus)
	fmt.Printf("Address: %s, Created: %s, Can Restore: %t\n",
		storageDetails.Address, storageDetails.CreatedAt, storageDetails.CanRestore)

	if storageDetails.Expires != "" {
		fmt.Printf("Expires: %s\n", storageDetails.Expires)
	}

	if storageDetails.ServerAlias != "" {
		fmt.Printf("Server Alias: %s\n", storageDetails.ServerAlias)
	}

	// Note: Credentials are only visible at creation time and during regeneration operations
	if storageDetails.Type == "s3" && storageDetails.Credentials.S3.AccessKey != "" {
		fmt.Printf("S3 Access Key: %s\n", storageDetails.Credentials.S3.AccessKey)
		fmt.Printf("S3 Secret Key: %s\n", storageDetails.Credentials.S3.SecretKey)
	}

	if storageDetails.Type == "sftp" {
		if storageDetails.Credentials.SftpPassword != "" {
			fmt.Printf("SFTP Password: %s\n", storageDetails.Credentials.SftpPassword)
		}
		if len(storageDetails.Credentials.Keys) > 0 {
			fmt.Printf("SSH Keys (%d):\n", len(storageDetails.Credentials.Keys))
			for _, key := range storageDetails.Credentials.Keys {
				fmt.Printf("  - ID: %d, Name: %s, Created: %s\n", key.ID, key.Name, key.CreatedAt)
			}
		}
	}

	fmt.Println("==================")
}

func listStorages(client *gcore.Client) {
	fmt.Println("\n=== LIST STORAGES ===")

	storages, err := client.Storage.List(context.Background(), storage.StorageListParams{})
	if err != nil {
		fmt.Printf("Error listing storages: %v\n", err)
		fmt.Println("====================")
		return
	}

	for i, s := range storages.Results {
		fmt.Printf("  %d. Storage: ID=%d, Name=%s, Type=%s, Location=%s, Status=%s\n",
			i+1, s.ID, s.Name, s.Type, s.Location, s.ProvisioningStatus)
	}

	fmt.Println("====================")
}

func updateStorage(client *gcore.Client, storageID int64) {
	fmt.Println("\n=== UPDATE STORAGE ===")

	params := storage.StorageUpdateParams{
		Expires: param.NewOpt("30 days"),
	}

	updatedStorage, err := client.Storage.Update(context.Background(), storageID, params)
	if err != nil {
		fmt.Printf("Error updating storage: %v\n", err)
		fmt.Println("======================")
		return
	}

	fmt.Printf("Updated Storage: ID=%d, Expires: %s\n", updatedStorage.ID, updatedStorage.Expires)

	fmt.Println("======================")
}

func deleteStorage(client *gcore.Client, storageID int64) {
	fmt.Println("\n=== DELETE STORAGE ===")

	err := client.Storage.Delete(context.Background(), storageID)
	if err != nil {
		fmt.Printf("Error deleting storage %d: %v\n", storageID, err)
		fmt.Println("======================")
		return
	}

	fmt.Printf("Storage %d deleted successfully\n", storageID)
	fmt.Println("======================")
}
