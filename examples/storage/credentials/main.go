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

	s3StorageID := createS3Storage(&client)
	sftpStorageID := createSFTPStorage(&client)
	waitForStorageProvisioning(&client, s3StorageID, sftpStorageID)
	getStorageDetails(&client, s3StorageID, sftpStorageID)
	regenerateS3Credentials(&client, s3StorageID)
	regenerateSFTPPassword(&client, sftpStorageID)
	setCustomSFTPPassword(&client, sftpStorageID)
	removeSFTPPassword(&client, sftpStorageID)
	reenableSFTPPassword(&client, sftpStorageID)
	getStorageDetails(&client, s3StorageID, sftpStorageID)
	cleanup(&client, s3StorageID, sftpStorageID)
}

func createS3Storage(client *gcore.Client) int64 {
	fmt.Println("\n=== CREATE S3 STORAGE ===")

	s3Name := fmt.Sprintf("s3-creds-example-%d", time.Now().Unix())

	params := storage.StorageNewParams{
		Name:     s3Name,
		Type:     storage.StorageNewParamsTypeS3,
		Location: storage.StorageNewParamsLocationFra,
	}

	newStorage, err := client.Storage.New(context.Background(), params)
	if err != nil {
		fmt.Printf("Error creating %s storage: %v\n", storage.StorageNewParamsTypeS3, err)
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

func createSFTPStorage(client *gcore.Client) int64 {
	fmt.Println("\n=== CREATE SFTP STORAGE ===")

	sftpName := fmt.Sprintf("sftp-creds-example-%d", time.Now().Unix())

	params := storage.StorageNewParams{
		Name:                 sftpName,
		Type:                 storage.StorageNewParamsTypeSftp,
		Location:             storage.StorageNewParamsLocationFra,
		GenerateSftpPassword: param.NewOpt(true),
	}

	newStorage, err := client.Storage.New(context.Background(), params)
	if err != nil {
		fmt.Printf("Error creating %s storage: %v\n", storage.StorageNewParamsTypeSftp, err)
		log.Fatalf("Failed to create SFTP storage")
	}

	fmt.Printf("Created Storage: ID=%d, Name=%s, Type=%s, Location=%s\n",
		newStorage.ID, newStorage.Name, newStorage.Type, newStorage.Location)
	fmt.Printf("Storage address: %s\n", newStorage.Address)

	// Display SFTP credentials
	if newStorage.Credentials.SftpPassword != "" {
		fmt.Printf("SFTP Password: %s\n", newStorage.Credentials.SftpPassword)
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
		storage, err := client.Storage.Get(ctx, s3StorageID)
		if err != nil {
			fmt.Printf("Error checking storage status: %v\n", err)
			break
		}
		if storage.ProvisioningStatus == "ok" {
			fmt.Printf("Storage %d is ready\n", s3StorageID)
			break
		}
		fmt.Printf("Storage %d status: %s, waiting...\n", s3StorageID, storage.ProvisioningStatus)
		time.Sleep(2 * time.Second)
	}

	// Wait for SFTP storage
	start = time.Now()
	for time.Since(start) < maxWait {
		storage, err := client.Storage.Get(ctx, sftpStorageID)
		if err != nil {
			fmt.Printf("Error checking storage status: %v\n", err)
			break
		}
		if storage.ProvisioningStatus == "ok" {
			fmt.Printf("Storage %d is ready\n", sftpStorageID)
			break
		}
		fmt.Printf("Storage %d status: %s, waiting...\n", sftpStorageID, storage.ProvisioningStatus)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("=====================================")
}

func getStorageDetails(client *gcore.Client, s3StorageID, sftpStorageID int64) {
	fmt.Println("\n=== GET STORAGE DETAILS ===")

	ctx := context.Background()

	// Get S3 storage details
	storageDetails, err := client.Storage.Get(ctx, s3StorageID)
	if err != nil {
		fmt.Printf("Error getting storage details: %v\n", err)
	} else {
		fmt.Printf("Storage: ID=%d, Name=%s, Type=%s, Location=%s, Status=%s\n",
			storageDetails.ID, storageDetails.Name, storageDetails.Type, storageDetails.Location, storageDetails.ProvisioningStatus)
		fmt.Printf("Address: %s, Created: %s, Can Restore: %t\n",
			storageDetails.Address, storageDetails.CreatedAt, storageDetails.CanRestore)
	}

	// Get SFTP storage details
	storageDetails, err = client.Storage.Get(ctx, sftpStorageID)
	if err != nil {
		fmt.Printf("Error getting storage details: %v\n", err)
	} else {
		fmt.Printf("Storage: ID=%d, Name=%s, Type=%s, Location=%s, Status=%s\n",
			storageDetails.ID, storageDetails.Name, storageDetails.Type, storageDetails.Location, storageDetails.ProvisioningStatus)
		fmt.Printf("Address: %s, Created: %s, Can Restore: %t\n",
			storageDetails.Address, storageDetails.CreatedAt, storageDetails.CanRestore)
	}

	fmt.Println("==========================")
}

func regenerateS3Credentials(client *gcore.Client, s3StorageID int64) {
	fmt.Println("\n=== REGENERATE S3 CREDENTIALS ===")

	params := storage.CredentialRecreateParams{
		GenerateS3Keys: param.NewOpt(true),
	}

	updatedStorage, err := client.Storage.Credentials.Recreate(context.Background(), s3StorageID, params)
	if err != nil {
		fmt.Printf("Error performing credential operation: %v\n", err)
		fmt.Println("================================")
		return
	}

	fmt.Printf("Generated new S3 credentials for storage %d:\n", s3StorageID)
	fmt.Printf("New Access Key: %s\n", updatedStorage.Credentials.S3.AccessKey)
	fmt.Printf("New Secret Key: %s\n", updatedStorage.Credentials.S3.SecretKey)

	fmt.Println("================================")
}

func regenerateSFTPPassword(client *gcore.Client, sftpStorageID int64) {
	fmt.Println("\n=== REGENERATE SFTP PASSWORD ===")

	params := storage.CredentialRecreateParams{
		GenerateSftpPassword: param.NewOpt(true),
	}

	updatedStorage, err := client.Storage.Credentials.Recreate(context.Background(), sftpStorageID, params)
	if err != nil {
		fmt.Printf("Error performing credential operation: %v\n", err)
		fmt.Println("===============================")
		return
	}

	fmt.Printf("Generated new SFTP password for storage %d: %s\n", sftpStorageID, updatedStorage.Credentials.SftpPassword)

	fmt.Println("===============================")
}

func setCustomSFTPPassword(client *gcore.Client, sftpStorageID int64) {
	fmt.Println("\n=== SET CUSTOM SFTP PASSWORD ===")

	params := storage.CredentialRecreateParams{
		SftpPassword: param.NewOpt("MyNewSecurePassword456!"),
	}

	updatedStorage, err := client.Storage.Credentials.Recreate(context.Background(), sftpStorageID, params)
	if err != nil {
		fmt.Printf("Error performing credential operation: %v\n", err)
		fmt.Println("=================================")
		return
	}

	fmt.Printf("Set custom SFTP password for storage %d: %s\n", sftpStorageID, updatedStorage.Credentials.SftpPassword)

	fmt.Println("=================================")
}

func removeSFTPPassword(client *gcore.Client, sftpStorageID int64) {
	fmt.Println("\n=== REMOVE SFTP PASSWORD ===")

	params := storage.CredentialRecreateParams{
		DeleteSftpPassword: param.NewOpt(true),
	}

	updatedStorage, err := client.Storage.Credentials.Recreate(context.Background(), sftpStorageID, params)
	if err != nil {
		fmt.Printf("Error performing credential operation: %v\n", err)
		fmt.Println("============================")
		return
	}

	fmt.Printf("Removed SFTP password for storage %d\n", sftpStorageID)
	fmt.Printf("Password authentication is now disabled\n")
	if updatedStorage.Credentials.SftpPassword == "" {
		fmt.Printf("Confirmed: SFTP password is empty\n")
	}

	fmt.Println("============================")
}

func reenableSFTPPassword(client *gcore.Client, sftpStorageID int64) {
	fmt.Println("\n=== RE-ENABLE SFTP PASSWORD ===")

	params := storage.CredentialRecreateParams{
		GenerateSftpPassword: param.NewOpt(true),
	}

	updatedStorage, err := client.Storage.Credentials.Recreate(context.Background(), sftpStorageID, params)
	if err != nil {
		fmt.Printf("Error performing credential operation: %v\n", err)
		fmt.Println("==============================")
		return
	}

	fmt.Printf("Generated new SFTP password for storage %d: %s\n", sftpStorageID, updatedStorage.Credentials.SftpPassword)

	fmt.Println("==============================")
}

func cleanup(client *gcore.Client, s3StorageID, sftpStorageID int64) {
	fmt.Println("\n=== CLEANUP ===")

	ctx := context.Background()

	// Delete S3 storage
	err := client.Storage.Delete(ctx, s3StorageID)
	if err != nil {
		fmt.Printf("Error deleting storage %d: %v\n", s3StorageID, err)
	} else {
		fmt.Printf("Storage %d deleted successfully\n", s3StorageID)
	}

	// Delete SFTP storage
	err = client.Storage.Delete(ctx, sftpStorageID)
	if err != nil {
		fmt.Printf("Error deleting storage %d: %v\n", sftpStorageID, err)
	} else {
		fmt.Printf("Storage %d deleted successfully\n", sftpStorageID)
	}

	fmt.Println("===============")
}
