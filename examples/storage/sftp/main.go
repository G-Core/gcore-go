// SFTP storage CRUD example demonstrating the polling helpers
// (NewAndPoll, UpdateAndPoll, DeleteAndPoll) on SftpStorageService.
//
// Reads GCORE_API_KEY from the environment.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/storage"
)

func main() {
	client := gcore.NewClient(
		// Bound the polling loops so a wedged provisioning won't hang the example.
		option.WithPollingIntervalSeconds(2),
		option.WithPollingTimeoutSeconds(180),
	)

	ctx := context.Background()
	name := fmt.Sprintf("sftp-crud-%d", time.Now().Unix())

	// --- Create + poll ----------------------------------------------------
	fmt.Println("=== NewAndPoll ===")
	created, err := client.Storage.SftpStorages.NewAndPoll(ctx, storage.SftpStorageNewParams{
		LocationName: "ams",
		Name:         name,
		PasswordMode: storage.SftpStorageNewParamsPasswordModeAuto,
	})
	if err != nil {
		log.Fatalf("NewAndPoll failed: %v", err)
	}
	fmt.Printf("created: id=%d name=%s status=%s address=%s\n", created.ID, created.Name, created.ProvisioningStatus, created.Address)
	fmt.Printf("one-time password set: %t (length=%d)\n", created.Password != "", len(created.Password))

	// --- Update + poll ----------------------------------------------------
	fmt.Println("\n=== UpdateAndPoll ===")
	updated, err := client.Storage.SftpStorages.UpdateAndPoll(ctx, created.ID, storage.SftpStorageUpdateParams{
		PasswordMode: storage.SftpStorageUpdateParamsPasswordModeAuto, // regenerate password
	})
	if err != nil {
		log.Fatalf("UpdateAndPoll failed: %v", err)
	}
	fmt.Printf("updated: id=%d status=%s\n", updated.ID, updated.ProvisioningStatus)
	fmt.Printf("regenerated password set: %t (length=%d)\n", updated.Password != "", len(updated.Password))

	// --- Delete + poll ----------------------------------------------------
	fmt.Println("\n=== DeleteAndPoll ===")
	if err := client.Storage.SftpStorages.DeleteAndPoll(ctx, created.ID); err != nil {
		log.Fatalf("DeleteAndPoll failed: %v", err)
	}
	fmt.Printf("deleted: id=%d\n", created.ID)
}
