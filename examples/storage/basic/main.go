package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/G-Core/gcore-go"
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
	listStorages(&client)
	getStorage(&client, storageID)
	deleteStorage(&client, storageID)
}

func createStorage(client *gcore.Client) int64 {
	fmt.Println("\n=== CREATE STORAGE ===")

	name := fmt.Sprintf("s3-basic-%d", time.Now().Unix())

	params := storage.ObjectStorageNewParams{
		Name:         name,
		LocationName: "s-ed1",
	}

	newStorage, err := client.Storage.ObjectStorages.New(context.Background(), params)
	if err != nil {
		fmt.Printf("Error creating S3 storage: %v\n", err)
		log.Fatalf("Failed to create storage")
	}

	fmt.Printf("Created Storage: ID=%d, Name=%s, Location=%s\n",
		newStorage.ID, newStorage.Name, newStorage.LocationName)
	fmt.Printf("Storage address: %s\n", newStorage.Address)

	// Display S3 credentials
	if len(newStorage.AccessKeys) > 0 {
		fmt.Printf("S3 Access Key: %s\n", newStorage.AccessKeys[0].AccessKey)
		fmt.Printf("S3 Secret Key: %s\n", newStorage.AccessKeys[0].SecretKey)
	}

	fmt.Println("======================")
	return newStorage.ID
}

func getStorage(client *gcore.Client, storageID int64) {
	fmt.Println("\n=== GET STORAGE ===")

	storageDetails, err := client.Storage.ObjectStorages.Get(context.Background(), storageID)
	if err != nil {
		fmt.Printf("Error getting storage details: %v\n", err)
		fmt.Println("==================")
		return
	}

	fmt.Printf("Storage: ID=%d, Name=%s, Location=%s, Status=%s\n",
		storageDetails.ID, storageDetails.Name, storageDetails.LocationName, storageDetails.ProvisioningStatus)
	fmt.Printf("Address: %s, Created: %s\n",
		storageDetails.Address, storageDetails.CreatedAt)

	fmt.Println("==================")
}

func listStorages(client *gcore.Client) {
	fmt.Println("\n=== LIST STORAGES ===")

	storages, err := client.Storage.ObjectStorages.List(context.Background(), storage.ObjectStorageListParams{})
	if err != nil {
		fmt.Printf("Error listing storages: %v\n", err)
		fmt.Println("====================")
		return
	}

	for i, s := range storages.Results {
		fmt.Printf("  %d. Storage: ID=%d, Name=%s, Location=%s, Status=%s\n",
			i+1, s.ID, s.Name, s.LocationName, s.ProvisioningStatus)
	}

	fmt.Println("====================")
}

func deleteStorage(client *gcore.Client, storageID int64) {
	fmt.Println("\n=== DELETE STORAGE ===")

	err := client.Storage.ObjectStorages.Delete(context.Background(), storageID)
	if err != nil {
		fmt.Printf("Error deleting storage %d: %v\n", storageID, err)
		fmt.Println("======================")
		return
	}

	fmt.Printf("Storage %d deleted successfully\n", storageID)
	fmt.Println("======================")
}
