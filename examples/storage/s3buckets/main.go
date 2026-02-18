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

	storageID := createS3Storage(&client)
	waitForStorageProvisioning(&client, storageID)
	bucketNames := createBuckets(&client, storageID)
	listBuckets(&client, storageID)
	setBucketLifecycle(&client, storageID, bucketNames[0])
	setBucketCORS(&client, storageID, bucketNames[0])
	setBucketPolicy(&client, storageID, bucketNames[0])
	listBuckets(&client, storageID)
	deleteBuckets(&client, storageID, bucketNames)
	deleteStorage(&client, storageID)
}

func createS3Storage(client *gcore.Client) int64 {
	fmt.Println("\n=== CREATE S3 STORAGE ===")

	storageName := fmt.Sprintf("s3-bucket-example-%d", time.Now().Unix())

	params := storage.StorageNewParams{
		Name:     storageName,
		Type:     storage.StorageNewParamsTypeS3Compatible,
		Location: "s-stage",
	}

	newStorage, err := client.Storage.New(context.Background(), params)
	if err != nil {
		fmt.Printf("Error creating %s storage: %v\n", storage.StorageNewParamsTypeS3Compatible, err)
		log.Fatalf("Failed to create S3 storage")
	}

	fmt.Printf("Created Storage: ID=%d, Name=%s, Type=%s, Location=%s\n",
		newStorage.ID, newStorage.Name, newStorage.Type, newStorage.Location)
	fmt.Printf("Storage address: %s\n", newStorage.Address)

	// Display S3 credentials
	if newStorage.Credentials.S3.AccessKey != "" {
		fmt.Printf("S3 Access Key: %s\n", newStorage.Credentials.S3.AccessKey)
		fmt.Printf("S3 Secret Key: %s\n", newStorage.Credentials.S3.SecretKey)
	}

	fmt.Println("=========================")
	return newStorage.ID
}

func waitForStorageProvisioning(client *gcore.Client, storageID int64) {
	fmt.Println("\n=== WAIT FOR STORAGE PROVISIONING ===")

	ctx := context.Background()
	maxWait := 30 * time.Second
	start := time.Now()

	for time.Since(start) < maxWait {
		storage, err := client.Storage.Get(ctx, storageID)
		if err != nil {
			fmt.Printf("Error checking storage status: %v\n", err)
			break
		}
		if storage.ProvisioningStatus == "ok" {
			fmt.Printf("Storage %d is ready\n", storageID)
			break
		}
		fmt.Printf("Storage %d status: %s, waiting...\n", storageID, storage.ProvisioningStatus)
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("Storage %d not ready, proceeding anyway...\n", storageID)
	fmt.Println("=====================================")
}

func createBuckets(client *gcore.Client, storageID int64) []string {
	fmt.Println("\n=== CREATE BUCKETS ===")

	bucketNames := []string{"example-bucket-1", "example-bucket-2", "test-bucket"}

	for _, bucketName := range bucketNames {
		params := storage.BucketNewParams{
			StorageID: storageID,
		}

		err := client.Storage.Buckets.New(context.Background(), bucketName, params)
		if err != nil {
			fmt.Printf("Error creating bucket %s: %v\n", bucketName, err)
			continue
		}

		fmt.Printf("Created bucket: %s\n", bucketName)
	}

	fmt.Println("=====================")
	return bucketNames
}

func listBuckets(client *gcore.Client, storageID int64) {
	fmt.Println("\n=== LIST BUCKETS ===")

	buckets, err := client.Storage.Buckets.List(context.Background(), storageID, storage.BucketListParams{})
	if err != nil {
		fmt.Printf("Error listing buckets: %v\n", err)
		fmt.Println("====================")
		return
	}

	for i, bucket := range buckets.Results {
		fmt.Printf("  %d. Bucket: Name=%s", i+1, bucket.Name)
		if bucket.Lifecycle > 0 {
			fmt.Printf(", Lifecycle: %d days", bucket.Lifecycle)
		}
		fmt.Println()
	}

	fmt.Println("====================")
}

func setBucketLifecycle(client *gcore.Client, storageID int64, bucketName string) {
	fmt.Println("\n=== SET BUCKET LIFECYCLE ===")

	// Set lifecycle to expire objects after specified days
	params := storage.BucketLifecycleNewParams{
		StorageID:      storageID,
		ExpirationDays: param.NewOpt(int64(30)),
	}

	err := client.Storage.Buckets.Lifecycle.New(context.Background(), bucketName, params)
	if err != nil {
		fmt.Printf("Error setting bucket lifecycle: %v\n", err)
		fmt.Println("============================")
		return
	}

	fmt.Printf("Set lifecycle policy for bucket %s: objects expire after %d days\n", bucketName, 30)

	fmt.Println("============================")
}

func setBucketCORS(client *gcore.Client, storageID int64, bucketName string) {
	fmt.Println("\n=== SET BUCKET CORS ===")

	// Configure CORS with specified origins
	params := storage.BucketCorNewParams{
		StorageID:      storageID,
		AllowedOrigins: []string{"*"},
	}

	err := client.Storage.Buckets.Cors.New(context.Background(), bucketName, params)
	if err != nil {
		fmt.Printf("Error setting bucket CORS: %v\n", err)
		fmt.Println("=======================")
		return
	}

	fmt.Printf("Set CORS policy for bucket %s with origins: %v\n", bucketName, []string{"*"})

	fmt.Println("=======================")
}

func setBucketPolicy(client *gcore.Client, storageID int64, bucketName string) {
	fmt.Println("\n=== SET BUCKET POLICY ===")

	// Apply a public read policy to the bucket
	params := storage.BucketPolicyNewParams{
		StorageID: storageID,
	}

	err := client.Storage.Buckets.Policy.New(context.Background(), bucketName, params)
	if err != nil {
		fmt.Printf("Error setting bucket policy: %v\n", err)
		fmt.Println("=========================")
		return
	}

	fmt.Printf("Set public read policy for bucket %s\n", bucketName)

	fmt.Println("=========================")
}

func deleteBuckets(client *gcore.Client, storageID int64, bucketNames []string) {
	fmt.Println("\n=== DELETE BUCKETS ===")

	ctx := context.Background()

	for _, bucketName := range bucketNames {
		params := storage.BucketDeleteParams{
			StorageID: storageID,
		}

		err := client.Storage.Buckets.Delete(ctx, bucketName, params)
		if err != nil {
			fmt.Printf("Error deleting bucket %s: %v\n", bucketName, err)
			continue
		}

		fmt.Printf("Deleted bucket: %s\n", bucketName)
	}

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
