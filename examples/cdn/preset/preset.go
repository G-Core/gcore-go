package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cdn"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read
	// from the GCORE_API_KEY environment variable if omitted.
	client := gcore.NewClient()

	// List the presets available to your account.
	presets := listPresets(&client)

	// Get details about the first preset.
	if len(presets) > 0 {
		getPreset(&client, presets[0].ID)
	}
}

func listPresets(client *gcore.Client) []cdn.PresetDetail {
	fmt.Println("\n=== LIST PRESETS ===")

	result, err := client.CDN.Presets.List(context.Background(), cdn.PresetListParams{})
	if err != nil {
		log.Fatalf("Error listing presets: %v", err)
	}

	for i, preset := range result.Results {
		fmt.Printf("  %d. Preset: ID=%d, Name=%s, ObjectType=%s\n",
			i+1, preset.ID, preset.Name, preset.ObjectType)
	}
	fmt.Println("====================")

	return result.Results
}

func getPreset(client *gcore.Client, presetID int64) {
	fmt.Println("\n=== GET PRESET ===")

	preset, err := client.CDN.Presets.Get(context.Background(), presetID)
	if err != nil {
		log.Fatalf("Error getting preset: %v", err)
	}

	fmt.Printf("Preset: ID=%d, Name=%s, ObjectType=%s, Service=%s, Settings=%d field(s)\n",
		preset.ID, preset.Name, preset.ObjectType, preset.Service, len(preset.PresetSettings))
	fmt.Println("==================")
}
