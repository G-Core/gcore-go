package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/option"
)

// This example demonstrates how polling methods handle WithRequestBody and
// WithResponseBodyInto options correctly. When users pass these options to
// *AndPoll methods, the SDK filters them internally so that:
//
//   - WithResponseBodyInto is excluded from the action call (New/Update/etc.)
//     and the Poll call, so TaskIDList and Task deserialize correctly.
//     It is kept for the final Get call, allowing custom deserialization.
//   - WithRequestBody is applied to the action call (New/Update/etc.) for
//     custom serialization. It is cleared for Poll and Get calls via
//     WithoutRequestBody(), preventing leftover request bodies.
//
// This is the same pattern used by the Terraform provider, which passes
// WithRequestBody for custom JSON serialization and WithResponseBodyInto
// for custom response deserialization.
func main() {
	cloudProjectID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_PROJECT_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_PROJECT_ID environment variable is required and must be a valid integer")
	}

	cloudRegionID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_REGION_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_REGION_ID environment variable is required and must be a valid integer")
	}

	client := gcore.NewClient(
		option.WithCloudProjectID(cloudProjectID),
		option.WithCloudRegionID(cloudRegionID),
	)

	// 1. Standard polling: NewAndPoll without extra options
	networkID := standardPolling(&client)

	// 2. Polling with WithRequestBody: custom request serialization
	withRequestBodyPolling(&client)

	// 3. Polling with WithResponseBodyInto: custom response deserialization
	withResponseBodyIntoPolling(&client)

	// 4. Polling with both WithRequestBody and WithResponseBodyInto
	withBothOptionsPolling(&client)

	// 5. Cleanup
	deleteNetwork(&client, networkID)
}

// standardPolling demonstrates the basic NewAndPoll flow.
// The SDK internally calls: New -> Poll -> Get, and returns the typed result.
func standardPolling(client *gcore.Client) string {
	fmt.Println("\n=== STANDARD POLLING ===")

	network, err := client.Cloud.Networks.NewAndPoll(context.Background(), cloud.NetworkNewParams{
		Name:         "gcore-go-polling-example",
		CreateRouter: gcore.Bool(false),
		Type:         cloud.NetworkNewParamsTypeVxlan,
	})
	if err != nil {
		log.Fatalf("Error creating network: %v", err)
	}

	fmt.Printf("Created network: ID=%s, Name=%s, Type=%s\n", network.ID, network.Name, network.Type)
	fmt.Println("========================")
	return network.ID
}

// withRequestBodyPolling demonstrates using WithRequestBody with a polling
// method. This is commonly used by the Terraform provider to control JSON
// serialization (e.g., sending "rules":[] to create a security group with
// no default egress rules).
//
// The SDK ensures that:
//   - The custom request body is sent to the action call (New)
//   - The request body is cleared for Poll and Get calls
func withRequestBodyPolling(client *gcore.Client) {
	fmt.Println("\n=== POLLING WITH WithRequestBody ===")

	// Build a custom JSON body. In practice, this is used to control
	// serialization details that the params struct doesn't support
	// (e.g., explicit empty arrays, null fields, extra fields).
	body := map[string]any{
		"name":          "gcore-go-polling-custom-body",
		"create_router": false,
		"type":          "vxlan",
	}
	bodyBytes, _ := json.Marshal(body)

	// NewAndPoll with WithRequestBody: the custom body is sent to the New call,
	// but cleared for the Poll and Get calls via WithoutRequestBody().
	network, err := client.Cloud.Networks.NewAndPoll(
		context.Background(),
		cloud.NetworkNewParams{},
		option.WithRequestBody("application/json", bodyBytes),
	)
	if err != nil {
		log.Fatalf("Error creating network with WithRequestBody: %v", err)
	}

	fmt.Printf("Created network: ID=%s, Name=%s\n", network.ID, network.Name)

	if network.Name != "gcore-go-polling-custom-body" {
		log.Fatalf("Expected name 'gcore-go-polling-custom-body', got '%s'", network.Name)
	}
	fmt.Println("Custom request body was applied to the action call correctly")
	fmt.Println("====================================")

	deleteNetwork(client, network.ID)
}

// CustomNetwork is a user-defined struct for custom deserialization.
// WithResponseBodyInto uses standard json.Unmarshal, so field types must match
// the raw API JSON format (e.g., tags are returned as an array by the API,
// while the SDK's Network struct uses a custom deserializer to convert them to a map).
type CustomNetwork struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Subnets   []string `json:"subnets"`
	ProjectID int64    `json:"project_id"`
	RegionID  int64    `json:"region_id"`
}

// withResponseBodyIntoPolling demonstrates using WithResponseBodyInto with a
// polling method. The SDK correctly filters it:
//   - Excluded from the New call (so TaskIDList deserializes correctly)
//   - Excluded from the Poll call (so Task deserializes correctly)
//   - Applied to the final Get call (user gets custom deserialization)
//
// Note: WithResponseBodyInto replaces the default deserialization target,
// so the standard typed return value (e.g. *Network) will be nil.
// The data is deserialized into the custom struct instead.
func withResponseBodyIntoPolling(client *gcore.Client) {
	fmt.Println("\n=== POLLING WITH WithResponseBodyInto ===")

	var customNet CustomNetwork

	// NewAndPoll with WithResponseBodyInto: the response from the final Get
	// call is deserialized into customNet instead of the default *Network.
	// The returned *Network will be nil — use customNet for the data.
	_, err := client.Cloud.Networks.NewAndPoll(
		context.Background(),
		cloud.NetworkNewParams{
			Name:         "gcore-go-polling-custom-deser",
			CreateRouter: gcore.Bool(false),
			Type:         cloud.NetworkNewParamsTypeVxlan,
		},
		option.WithResponseBodyInto(&customNet),
	)
	if err != nil {
		log.Fatalf("Error creating network with WithResponseBodyInto: %v", err)
	}

	fmt.Printf("Custom struct: ID=%s, Name=%s, Type=%s\n", customNet.ID, customNet.Name, customNet.Type)
	fmt.Printf("Custom struct: ProjectID=%d, RegionID=%d\n", customNet.ProjectID, customNet.RegionID)

	if customNet.ID == "" {
		log.Fatal("Expected customNet.ID to be populated")
	}

	raw, _ := json.MarshalIndent(customNet, "", "  ")
	fmt.Printf("Custom struct as JSON:\n%s\n", string(raw))
	fmt.Println("=========================================")

	deleteNetwork(client, customNet.ID)
}

// withBothOptionsPolling demonstrates the Terraform provider pattern:
// using both WithRequestBody (custom serialization) and WithResponseBodyInto
// (custom deserialization) with a polling method.
//
// The SDK ensures:
//   - WithRequestBody is applied to the action call only (cleared for Poll/Get)
//   - WithResponseBodyInto is applied to the final Get call only (excluded from action/Poll)
func withBothOptionsPolling(client *gcore.Client) {
	fmt.Println("\n=== POLLING WITH WithRequestBody + WithResponseBodyInto ===")

	body := map[string]any{
		"name":          "gcore-go-polling-both-options",
		"create_router": false,
		"type":          "vxlan",
	}
	bodyBytes, _ := json.Marshal(body)

	var customNet CustomNetwork

	// This is the pattern used by the Terraform provider:
	// - WithRequestBody: custom JSON body for the New call
	// - WithResponseBodyInto: custom struct for the Get call
	_, err := client.Cloud.Networks.NewAndPoll(
		context.Background(),
		cloud.NetworkNewParams{},
		option.WithRequestBody("application/json", bodyBytes),
		option.WithResponseBodyInto(&customNet),
	)
	if err != nil {
		log.Fatalf("Error creating network with both options: %v", err)
	}

	fmt.Printf("Custom struct: ID=%s, Name=%s, Type=%s\n", customNet.ID, customNet.Name, customNet.Type)

	if customNet.Name != "gcore-go-polling-both-options" {
		log.Fatalf("Expected name 'gcore-go-polling-both-options', got '%s'", customNet.Name)
	}
	fmt.Println("Both options worked correctly together")
	fmt.Println("=========================================================")

	deleteNetwork(client, customNet.ID)
}

func deleteNetwork(client *gcore.Client, networkID string) {
	fmt.Printf("\n=== DELETE NETWORK %s ===\n", networkID)
	err := client.Cloud.Networks.DeleteAndPoll(context.Background(), networkID, cloud.NetworkDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting network: %v", err)
	}
	fmt.Printf("Network %s successfully deleted\n", networkID)
	fmt.Println("========================")
}
