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

	s3StorageID := createS3Storage(&client)
	sftpStorageID := createSFTPStorage(&client)
	waitForStorageProvisioning(&client, s3StorageID, sftpStorageID)
	getStorageDetails(&client, s3StorageID, sftpStorageID)
	newAccessKey := regenerateS3Credentials(&client, s3StorageID)
	listS3AccessKeys(&client, s3StorageID)
	deleteS3AccessKey(&client, s3StorageID, newAccessKey)
	listS3AccessKeys(&client, s3StorageID)
	regenerateSFTPPassword(&client, sftpStorageID)
	removeSFTPPassword(&client, sftpStorageID)
	reenableSFTPPassword(&client, sftpStorageID)
	getStorageDetails(&client, s3StorageID, sftpStorageID)
	cleanup(&client, s3StorageID, sftpStorageID)
}

func createS3Storage(client *gcore.Client) int64 {
	fmt.Println("\n=== CREATE S3 STORAGE ===")

	s3Name := fmt.Sprintf("s3-creds-%d", time.Now().Unix())

	params := storage.ObjectStorageNewParams{
		Name:         s3Name,
		LocationName: "s-ed1",
	}

	newStorage, err := client.Storage.ObjectStorages.New(context.Background(), params)
	if err != nil {
		fmt.Printf("Error creating S3 storage: %v\n", err)
		log.Fatalf("Failed to create S3 storage")
	}

	fmt.Printf("Created Storage: ID=%d, Name=%s, Location=%s\n",
		newStorage.ID, newStorage.Name, newStorage.LocationName)
	fmt.Printf("Storage address: %s\n", newStorage.Address)

	// Display S3 credentials
	if len(newStorage.AccessKeys) > 0 {
		fmt.Printf("S3 Access Key: %s\n", newStorage.AccessKeys[0].AccessKey)
		fmt.Printf("S3 Secret Key: %s\n", newStorage.AccessKeys[0].SecretKey)
	}

	fmt.Println("=========================")
	return newStorage.ID
}

func createSFTPStorage(client *gcore.Client) int64 {
	fmt.Println("\n=== CREATE SFTP STORAGE ===")

	sftpName := fmt.Sprintf("sftp-creds-%d", time.Now().Unix())

	params := storage.SftpStorageNewParams{
		Name:         sftpName,
		LocationName: "ams",
		PasswordMode: storage.SftpStorageNewParamsPasswordModeAuto,
	}

	newStorage, err := client.Storage.SftpStorages.New(context.Background(), params)
	if err != nil {
		fmt.Printf("Error creating SFTP storage: %v\n", err)
		log.Fatalf("Failed to create SFTP storage")
	}

	fmt.Printf("Created Storage: ID=%d, Name=%s, Location=%s\n",
		newStorage.ID, newStorage.Name, newStorage.LocationName)
	fmt.Printf("Storage address: %s\n", newStorage.Address)

	// Display SFTP credentials
	if newStorage.Password != "" {
		fmt.Printf("SFTP Password: %s\n", newStorage.Password)
	}

	fmt.Println("==========================")
	return newStorage.ID
}

func waitForStorageProvisioning(client *gcore.Client, s3StorageID, sftpStorageID int64) {
	fmt.Println("\n=== WAIT FOR STORAGE PROVISIONING ===")

	ctx := context.Background()
	maxWait := 30 * time.Second

	// Wait for S3 storage
	start := time.Now()
	for time.Since(start) < maxWait {
		s, err := client.Storage.ObjectStorages.Get(ctx, s3StorageID)
		if err != nil {
			fmt.Printf("Error checking storage status: %v\n", err)
			break
		}
		if s.ProvisioningStatus == storage.S3StorageProvisioningStatusActive {
			fmt.Printf("Storage %d is ready\n", s3StorageID)
			break
		}
		fmt.Printf("Storage %d status: %s, waiting...\n", s3StorageID, s.ProvisioningStatus)
		time.Sleep(2 * time.Second)
	}

	// Wait for SFTP storage
	start = time.Now()
	for time.Since(start) < maxWait {
		s, err := client.Storage.SftpStorages.Get(ctx, sftpStorageID)
		if err != nil {
			fmt.Printf("Error checking storage status: %v\n", err)
			break
		}
		if s.ProvisioningStatus == storage.SftpStorageProvisioningStatusActive {
			fmt.Printf("Storage %d is ready\n", sftpStorageID)
			break
		}
		fmt.Printf("Storage %d status: %s, waiting...\n", sftpStorageID, s.ProvisioningStatus)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("=====================================")
}

func getStorageDetails(client *gcore.Client, s3StorageID, sftpStorageID int64) {
	fmt.Println("\n=== GET STORAGE DETAILS ===")

	ctx := context.Background()

	// Get S3 storage details
	s3Details, err := client.Storage.ObjectStorages.Get(ctx, s3StorageID)
	if err != nil {
		fmt.Printf("Error getting S3 storage details: %v\n", err)
	} else {
		fmt.Printf("S3 Storage: ID=%d, Name=%s, Location=%s, Status=%s\n",
			s3Details.ID, s3Details.Name, s3Details.LocationName, s3Details.ProvisioningStatus)
		fmt.Printf("Address: %s, Created: %s\n",
			s3Details.Address, s3Details.CreatedAt)
	}

	// Get SFTP storage details
	sftpDetails, err := client.Storage.SftpStorages.Get(ctx, sftpStorageID)
	if err != nil {
		fmt.Printf("Error getting SFTP storage details: %v\n", err)
	} else {
		fmt.Printf("SFTP Storage: ID=%d, Name=%s, Location=%s, Status=%s\n",
			sftpDetails.ID, sftpDetails.Name, sftpDetails.LocationName, sftpDetails.ProvisioningStatus)
		fmt.Printf("Address: %s, Created: %s, HasPassword: %t\n",
			sftpDetails.Address, sftpDetails.CreatedAt, sftpDetails.HasPassword)
	}

	fmt.Println("==========================")
}

func regenerateS3Credentials(client *gcore.Client, s3StorageID int64) string {
	fmt.Println("\n=== REGENERATE S3 CREDENTIALS ===")

	// Create a new access key for the S3 storage (max 2 per storage)
	newKey, err := client.Storage.ObjectStorages.AccessKeys.New(context.Background(), s3StorageID)
	if err != nil {
		fmt.Printf("Error creating new access key: %v\n", err)
		fmt.Println("================================")
		return ""
	}

	fmt.Printf("Created new S3 access key for storage %d:\n", s3StorageID)
	fmt.Printf("New Access Key: %s\n", newKey.AccessKey)
	fmt.Printf("New Secret Key: %s\n", newKey.SecretKey)

	fmt.Println("================================")
	return newKey.AccessKey
}

func listS3AccessKeys(client *gcore.Client, s3StorageID int64) {
	fmt.Println("\n=== LIST S3 ACCESS KEYS ===")

	keys, err := client.Storage.ObjectStorages.AccessKeys.List(context.Background(), s3StorageID, storage.ObjectStorageAccessKeyListParams{})
	if err != nil {
		fmt.Printf("Error listing access keys: %v\n", err)
		fmt.Println("===========================")
		return
	}

	for i, key := range keys.Results {
		fmt.Printf("  %d. Access Key: %s, Created: %s\n", i+1, key.AccessKey, key.CreatedAt)
	}

	fmt.Println("===========================")
}

func deleteS3AccessKey(client *gcore.Client, s3StorageID int64, accessKey string) {
	fmt.Println("\n=== DELETE S3 ACCESS KEY ===")

	if accessKey == "" {
		fmt.Println("No access key to delete")
		fmt.Println("===========================")
		return
	}

	params := storage.ObjectStorageAccessKeyDeleteParams{
		StorageID: s3StorageID,
	}

	err := client.Storage.ObjectStorages.AccessKeys.Delete(context.Background(), accessKey, params)
	if err != nil {
		fmt.Printf("Error deleting access key %s: %v\n", accessKey, err)
		fmt.Println("===========================")
		return
	}

	fmt.Printf("Deleted access key: %s\n", accessKey)
	fmt.Println("===========================")
}

func regenerateSFTPPassword(client *gcore.Client, sftpStorageID int64) {
	fmt.Println("\n=== REGENERATE SFTP PASSWORD ===")

	params := storage.SftpStorageUpdateParams{
		PasswordMode: storage.SftpStorageUpdateParamsPasswordModeAuto,
	}

	updatedStorage, err := client.Storage.SftpStorages.Update(context.Background(), sftpStorageID, params)
	if err != nil {
		fmt.Printf("Error regenerating SFTP password: %v\n", err)
		fmt.Println("===============================")
		return
	}

	fmt.Printf("Generated new SFTP password for storage %d: %s\n", sftpStorageID, updatedStorage.Password)

	fmt.Println("===============================")
}

func removeSFTPPassword(client *gcore.Client, sftpStorageID int64) {
	fmt.Println("\n=== REMOVE SFTP PASSWORD ===")

	params := storage.SftpStorageUpdateParams{
		PasswordMode: storage.SftpStorageUpdateParamsPasswordModeNone,
	}

	updatedStorage, err := client.Storage.SftpStorages.Update(context.Background(), sftpStorageID, params)
	if err != nil {
		fmt.Printf("Error removing SFTP password: %v\n", err)
		fmt.Println("============================")
		return
	}

	fmt.Printf("Removed SFTP password for storage %d\n", sftpStorageID)
	fmt.Printf("Password authentication is now disabled (HasPassword: %t)\n", updatedStorage.HasPassword)

	fmt.Println("============================")
}

func reenableSFTPPassword(client *gcore.Client, sftpStorageID int64) {
	fmt.Println("\n=== RE-ENABLE SFTP PASSWORD ===")

	params := storage.SftpStorageUpdateParams{
		PasswordMode: storage.SftpStorageUpdateParamsPasswordModeAuto,
	}

	updatedStorage, err := client.Storage.SftpStorages.Update(context.Background(), sftpStorageID, params)
	if err != nil {
		fmt.Printf("Error re-enabling SFTP password: %v\n", err)
		fmt.Println("==============================")
		return
	}

	fmt.Printf("Generated new SFTP password for storage %d: %s\n", sftpStorageID, updatedStorage.Password)

	fmt.Println("==============================")
}

func cleanup(client *gcore.Client, s3StorageID, sftpStorageID int64) {
	fmt.Println("\n=== CLEANUP ===")

	ctx := context.Background()

	// Delete S3 storage
	err := client.Storage.ObjectStorages.Delete(ctx, s3StorageID)
	if err != nil {
		fmt.Printf("Error deleting storage %d: %v\n", s3StorageID, err)
	} else {
		fmt.Printf("Storage %d deleted successfully\n", s3StorageID)
	}

	// Delete SFTP storage
	err = client.Storage.SftpStorages.Delete(ctx, sftpStorageID)
	if err != nil {
		fmt.Printf("Error deleting storage %d: %v\n", sftpStorageID, err)
	} else {
		fmt.Printf("Storage %d deleted successfully\n", sftpStorageID)
	}

	fmt.Println("===============")
}
