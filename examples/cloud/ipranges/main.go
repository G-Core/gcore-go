package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/stainless-sdks/gcore-go"
	"github.com/stainless-sdks/gcore-go/option"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	apiKey := os.Getenv("GCORE_API_KEY")
	// Will use Production API URL if omitted
	baseURL := os.Getenv("GCORE_API_URL")

	client := gcore.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL(baseURL),
	)

	listIPRanges(&client)
}

func listIPRanges(client *gcore.Client) {
	fmt.Println("\n=== LIST IP RANGES ===")

	ctx := context.Background()

	ipRanges, err := client.Cloud.IPRanges.List(ctx)
	if err != nil {
		log.Fatalf("Error listing IP ranges: %v", err)
	}

	for i, ipRange := range ipRanges.Ranges {
		fmt.Printf("  %d. IP Range: %s\n", i+1, ipRange)
	}

	fmt.Println("=======================")
}
