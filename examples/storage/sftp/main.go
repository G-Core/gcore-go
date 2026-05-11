// SFTP storage CRUD example demonstrating the polling helpers
// (NewAndPoll, UpdateAndPoll, DeleteAndPoll) on SftpStorageService.
//
// The first half of main runs the simple SDK pattern (typed params, typed
// return). The second half runs the same lifecycle in the Terraform-provider
// pattern: option.WithRequestBody supplies a pre-marshalled JSON body and
// option.WithResponseBodyInto captures the raw *http.Response so the caller
// can read both the typed return and the original API payload.
//
// Reads GCORE_API_KEY from the environment.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

	runSimple(ctx, &client)
	runTerraformStyle(ctx, &client)
}

// runSimple uses the polling helpers with no extra request options — the
// typed return is the only observable.
func runSimple(ctx context.Context, client *gcore.Client) {
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

// runTerraformStyle mirrors how the gcore-terraform provider drives the
// polling helpers: WithRequestBody supplies a pre-marshalled JSON body and
// WithResponseBodyInto(**http.Response) captures the raw API response. The
// provider discards the typed return entirely and decodes the raw body into
// its own state model — so this example discards it too (`_, err :=`) and
// reads the resource ID and other fields out of the captured *http.Response.
func runTerraformStyle(ctx context.Context, client *gcore.Client) {
	name := fmt.Sprintf("sftp-tf-%d", time.Now().Unix())

	// --- Create + poll (Terraform style) ----------------------------------
	fmt.Println("\n=== NewAndPoll (Terraform-style options) ===")
	createBody, err := json.Marshal(map[string]any{
		"location_name": "ams",
		"name":          name,
		"password_mode": "auto",
	})
	if err != nil {
		log.Fatalf("marshal create body: %v", err)
	}

	var rawCreate *http.Response
	if _, err := client.Storage.SftpStorages.NewAndPoll(
		ctx,
		// Empty struct: WithRequestBody overrides the SDK-marshalled body.
		storage.SftpStorageNewParams{},
		option.WithRequestBody("application/json", createBody),
		option.WithResponseBodyInto(&rawCreate),
	); err != nil {
		log.Fatalf("NewAndPoll (tf) failed: %v", err)
	}
	created := decodeRawBody("create", rawCreate)

	// --- Update + poll (Terraform style) ----------------------------------
	fmt.Println("\n=== UpdateAndPoll (Terraform-style options) ===")
	updateBody, err := json.Marshal(map[string]any{
		"password_mode": "auto", // regenerate password
	})
	if err != nil {
		log.Fatalf("marshal update body: %v", err)
	}

	var rawUpdate *http.Response
	if _, err := client.Storage.SftpStorages.UpdateAndPoll(
		ctx,
		created.ID,
		storage.SftpStorageUpdateParams{},
		option.WithRequestBody("application/json", updateBody),
		option.WithResponseBodyInto(&rawUpdate),
	); err != nil {
		log.Fatalf("UpdateAndPoll (tf) failed: %v", err)
	}
	decodeRawBody("update", rawUpdate)

	// --- Delete + poll (Terraform style) ----------------------------------
	fmt.Println("\n=== DeleteAndPoll (Terraform-style options) ===")
	if err := client.Storage.SftpStorages.DeleteAndPoll(ctx, created.ID); err != nil {
		log.Fatalf("DeleteAndPoll (tf) failed: %v", err)
	}
	fmt.Printf("deleted: id=%d\n", created.ID)
}

// decodeRawBody reads raw.Body, decodes it into a storage.SftpStorage so the
// caller can use the typed fields (e.g. the resource id), and prints a summary.
func decodeRawBody(label string, raw *http.Response) storage.SftpStorage {
	if raw == nil || raw.Body == nil {
		log.Fatalf("raw %s body: <nil>", label)
	}
	body, err := io.ReadAll(raw.Body)
	_ = raw.Body.Close()
	if err != nil {
		log.Fatalf("read raw %s body: %v", label, err)
	}
	var decoded storage.SftpStorage
	if err := json.Unmarshal(body, &decoded); err != nil {
		log.Fatalf("decode raw %s body: %v", label, err)
	}
	fmt.Printf("raw %s body: %d bytes, id=%d, provisioning_status=%q, password_set=%t\n",
		label, len(body), decoded.ID, decoded.ProvisioningStatus, decoded.Password != "")
	return decoded
}
