package main

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/internal/apierror"
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

	// Create SSH keys with unique names
	ts := time.Now().Unix()
	key1ID := createSSHKey(&client, fmt.Sprintf("example-key-1-%d", ts))
	key2ID := createSSHKey(&client, fmt.Sprintf("example-key-2-%d", ts))
	listSSHKeys(&client)
	getSSHKey(&client, key1ID)

	// Create an SFTP storage associated with the SSH keys
	sftpStorageID := createSFTPStorageWithKeys(&client, key1ID, key2ID)
	waitForStorageProvisioning(&client, sftpStorageID)

	// Update the storage to only use the first key
	updateStorageSSHKeys(&client, sftpStorageID, key1ID)

	// Cleanup
	cleanup(&client, sftpStorageID, key1ID, key2ID)
}

func createSSHKey(client *gcore.Client, name string) int64 {
	fmt.Printf("\n=== CREATE SSH KEY: %s ===\n", name)

	// Generate a temporary ed25519 key pair for the example.
	// In production, you would use your own pre-existing SSH public key.
	publicKey, err := generateTempSSHKey()
	if err != nil {
		log.Fatalf("Failed to generate SSH key: %v", err)
	}

	params := storage.SSHKeyNewParams{
		Name:      name,
		PublicKey: publicKey,
	}

	key, err := client.Storage.SSHKeys.New(context.Background(), params)
	if err != nil {
		fmt.Printf("Error creating SSH key: %v\n", err)
		log.Fatalf("Failed to create SSH key %s", name)
	}

	fmt.Printf("Created SSH Key: ID=%d, Name=%s\n", key.ID, key.Name)
	fmt.Println("==========================")
	return key.ID
}

func listSSHKeys(client *gcore.Client) {
	fmt.Println("\n=== LIST SSH KEYS ===")

	keys, err := client.Storage.SSHKeys.List(context.Background(), storage.SSHKeyListParams{})
	if err != nil {
		fmt.Printf("Error listing SSH keys: %v\n", err)
		fmt.Println("=====================")
		return
	}

	for i, key := range keys.Results {
		fmt.Printf("  %d. ID=%d, Name=%s, Created=%s\n", i+1, key.ID, key.Name, key.CreatedAt)
	}

	fmt.Println("=====================")
}

func getSSHKey(client *gcore.Client, keyID int64) {
	fmt.Println("\n=== GET SSH KEY ===")

	key, err := client.Storage.SSHKeys.Get(context.Background(), keyID)
	if err != nil {
		fmt.Printf("Error getting SSH key: %v\n", err)
		fmt.Println("==================")
		return
	}

	fmt.Printf("SSH Key: ID=%d, Name=%s, Created=%s\n", key.ID, key.Name, key.CreatedAt)
	fmt.Printf("Public Key: %s...\n", key.PublicKey[:50])

	fmt.Println("==================")
}

func createSFTPStorageWithKeys(client *gcore.Client, keyIDs ...int64) int64 {
	fmt.Println("\n=== CREATE SFTP STORAGE WITH SSH KEYS ===")

	name := fmt.Sprintf("sftp-ssh-%d", time.Now().Unix())

	params := storage.SftpStorageNewParams{
		Name:         name,
		LocationName: "ams",
		PasswordMode: storage.SftpStorageNewParamsPasswordModeNone,
		SSHKeyIDs:    keyIDs,
	}

	newStorage, err := client.Storage.SftpStorages.New(context.Background(), params)
	if err != nil {
		fmt.Printf("Error creating SFTP storage: %v\n", err)
		log.Fatalf("Failed to create SFTP storage")
	}

	fmt.Printf("Created Storage: ID=%d, Name=%s, Location=%s\n",
		newStorage.ID, newStorage.Name, newStorage.LocationName)
	fmt.Printf("SSH Key IDs: %v\n", newStorage.SSHKeyIDs)
	fmt.Printf("Has Password: %t (SSH-key-only access)\n", newStorage.HasPassword)

	fmt.Println("=========================================")
	return newStorage.ID
}

func waitForStorageProvisioning(client *gcore.Client, storageID int64) {
	fmt.Println("\n=== WAIT FOR STORAGE PROVISIONING ===")

	ctx := context.Background()
	maxWait := 30 * time.Second
	start := time.Now()

	for time.Since(start) < maxWait {
		s, err := client.Storage.SftpStorages.Get(ctx, storageID)
		if err != nil {
			fmt.Printf("Error checking storage status: %v\n", err)
			break
		}
		if s.ProvisioningStatus == storage.SftpStorageProvisioningStatusActive {
			fmt.Printf("Storage %d is ready\n", storageID)
			fmt.Println("=====================================")
			return
		}
		fmt.Printf("Storage %d status: %s, waiting...\n", storageID, s.ProvisioningStatus)
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("Storage %d not ready, proceeding anyway...\n", storageID)
	fmt.Println("=====================================")
}

func updateStorageSSHKeys(client *gcore.Client, storageID int64, keyID int64) {
	fmt.Println("\n=== UPDATE STORAGE SSH KEYS ===")

	// Replace the storage's SSH keys with only the specified key
	params := storage.SftpStorageUpdateParams{
		SSHKeyIDs: []int64{keyID},
	}

	updatedStorage, err := client.Storage.SftpStorages.Update(context.Background(), storageID, params)
	if err != nil {
		fmt.Printf("Error updating storage SSH keys: %v\n", err)
		fmt.Println("==============================")
		return
	}

	fmt.Printf("Updated storage %d SSH keys: %v\n", storageID, updatedStorage.SSHKeyIDs)

	fmt.Println("==============================")
}

// generateTempSSHKey generates a temporary ed25519 SSH public key in OpenSSH
// authorized_keys format. In production, you would use your own pre-existing
// SSH public key instead.
func generateTempSSHKey() (string, error) {
	pub, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", err
	}

	// Build the SSH wire format: string "ssh-ed25519" + string <32-byte key>
	keyType := "ssh-ed25519"
	var blob []byte
	// Encode key type length + key type
	typeLen := make([]byte, 4)
	binary.BigEndian.PutUint32(typeLen, uint32(len(keyType)))
	blob = append(blob, typeLen...)
	blob = append(blob, []byte(keyType)...)
	// Encode public key length + public key
	keyLen := make([]byte, 4)
	binary.BigEndian.PutUint32(keyLen, uint32(len(pub)))
	blob = append(blob, keyLen...)
	blob = append(blob, pub...)

	return fmt.Sprintf("%s %s example@demo", keyType, base64.StdEncoding.EncodeToString(blob)), nil
}

func cleanup(client *gcore.Client, storageID int64, keyIDs ...int64) {
	fmt.Println("\n=== CLEANUP ===")

	ctx := context.Background()

	// Delete the SFTP storage and wait for it to be fully removed
	err := client.Storage.SftpStorages.Delete(ctx, storageID)
	if err != nil {
		fmt.Printf("Error deleting storage %d: %v\n", storageID, err)
	} else {
		fmt.Printf("Storage %d delete requested, waiting for completion...\n", storageID)
		// Poll until the storage is fully deleted
		for range 30 {
			time.Sleep(2 * time.Second)
			_, getErr := client.Storage.SftpStorages.Get(ctx, storageID)
			if getErr != nil {
				var apiErr *apierror.Error
				if errors.As(getErr, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
					fmt.Printf("Storage %d deleted successfully\n", storageID)
					break
				}
				// Transient error — keep polling
				fmt.Printf("Error checking storage %d status: %v\n", storageID, getErr)
			}
		}
	}

	// Delete the SSH keys
	for _, keyID := range keyIDs {
		err := client.Storage.SSHKeys.Delete(ctx, keyID)
		if err != nil {
			fmt.Printf("Error deleting SSH key %d: %v\n", keyID, err)
		} else {
			fmt.Printf("SSH key %d deleted successfully\n", keyID)
		}
	}

	fmt.Println("===============")
}
