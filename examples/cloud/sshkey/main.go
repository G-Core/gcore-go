package main

import (
	"context"
	"fmt"
	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/option"
	"log"
	"os"
	"strconv"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	//apiKey := os.Getenv("GCORE_API_KEY")
	// Will use Production API URL if omitted
	//baseURL := os.Getenv("GCORE_API_URL")

	// TODO set cloud project ID before running
	cloudProjectId, _ := strconv.ParseInt(os.Getenv("GCORE_CLOUD_PROJECT_ID"), 10, 64)

	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectId),
	)

	// Create an SSH key and use its ID for other operations
	sshKeyID := createSSHKey(&client)
	listAllSSHKeys(&client)
	listSSHKeysWithFilters(&client)
	listSSHKeysWithAutopager(&client)
	getSSHKeyByID(&client, sshKeyID)
	updateSSHKey(&client, sshKeyID)
	deleteSSHKey(&client, sshKeyID)
}

func createSSHKey(client *gcore.Client) string {
	fmt.Println("\n=== CREATE SSH KEY ===")

	sshKey, err := client.Cloud.SSHKeys.New(context.Background(), cloud.SSHKeyNewParams{
		Name: "New Test SSH Key",
	})
	if err != nil {
		log.Fatalf("Error creating SSH key: %v", err)
	}

	fmt.Printf("Created SSH Key ID: %s, Name: %s\n", sshKey.ID, sshKey.Name)
	fmt.Println("=======================")

	return sshKey.ID
}

func listAllSSHKeys(client *gcore.Client) {
	fmt.Println("\n=== LIST ALL SSH KEYS ===")

	sshKeys, err := client.Cloud.SSHKeys.List(context.Background(), cloud.SSHKeyListParams{})
	if err != nil {
		log.Fatalf("Error listing SSH keys: %v", err)
	}

	for i, sshKey := range sshKeys.Results {
		fmt.Printf("  %d. SSH Key ID: %s, Name: %s\n", i+1, sshKey.ID, sshKey.Name)
	}

	fmt.Println("=========================")
}

func listSSHKeysWithFilters(client *gcore.Client) {
	fmt.Println("\n=== LIST SSH KEYS WITH FILTERS ===")

	sshKeys, err := client.Cloud.SSHKeys.List(context.Background(), cloud.SSHKeyListParams{
		OrderBy: cloud.SSHKeyListParamsOrderByCreatedAtAsc,
	})
	if err != nil {
		log.Fatalf("Error listing SSH keys with filters: %v", err)
	}

	for i, sshKey := range sshKeys.Results {
		fmt.Printf("  %d. SSH Key ID: %s, Name: %16s, CreatedAt: %s\n", i+1, sshKey.ID, sshKey.Name,
			sshKey.CreatedAt)
	}

	fmt.Println("==================================")
}

func listSSHKeysWithAutopager(client *gcore.Client) {
	fmt.Println("\n=== LIST SSH KEYS USING AUTOPAGER ===")

	count := 0

	iter := client.Cloud.SSHKeys.ListAutoPaging(context.Background(), cloud.SSHKeyListParams{
		Limit: gcore.Int(10), // Process 10 items per page
	})

	for iter.Next() {
		sshKey := iter.Current()
		count++
		fmt.Printf("  %d. SSH Key ID: %s, Name: %s\n", count, sshKey.ID, sshKey.Name)
	}

	if err := iter.Err(); err != nil {
		log.Fatalf("Error iterating SSH keys: %v", err)
	}

	fmt.Println("=====================================")
}

func getSSHKeyByID(client *gcore.Client, sshKeyID string) {
	fmt.Println("\n=== GET SSH KEY BY ID ===")

	sshKey, err := client.Cloud.SSHKeys.Get(context.Background(), sshKeyID, cloud.SSHKeyGetParams{})
	if err != nil {
		log.Fatalf("Error getting SSH key: %v", err)
	}

	fmt.Printf("SSH Key ID: %s, Name: %s\n", sshKey.ID, sshKey.Name)
	fmt.Println("=========================")
}

func updateSSHKey(client *gcore.Client, sshKeyID string) {
	fmt.Println("\n=== UPDATE SSH KEY ===")

	sshKey, err := client.Cloud.SSHKeys.Update(context.Background(), sshKeyID, cloud.SSHKeyUpdateParams{
		SharedInProject: true,
	})
	if err != nil {
		log.Fatalf("Error updating SSH key: %v", err)
	}

	fmt.Printf("Updated SSH Key ID: %s, Name: %s, SharedInProject: %t\n", sshKey.ID, sshKey.Name,
		sshKey.SharedInProject)
	fmt.Println("=======================")
}

func deleteSSHKey(client *gcore.Client, sshKeyID string) {
	fmt.Println("\n=== DELETE SSH KEY ===")

	err := client.Cloud.SSHKeys.Delete(context.Background(), sshKeyID, cloud.SSHKeyDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting SSH key: %v", err)
	}

	fmt.Printf("SSH Key with ID %s successfully deleted\n", sshKeyID)
	fmt.Println("=======================")
}
