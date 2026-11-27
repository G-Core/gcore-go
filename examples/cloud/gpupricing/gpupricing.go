package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
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

	catalog := listGPUPrices(&client)
	filterGPUPrices(catalog, cloud.GPUCurrencyEur, cloud.GPUDeploymentTypeVirtual)
}

func listGPUPrices(client *gcore.Client) *cloud.GPUPriceCatalog {
	fmt.Println("\n=== LIST GPU PRICES ===")

	ctx := context.Background()

	catalog, err := client.Cloud.Pricing.GPU.List(ctx)
	if err != nil {
		log.Fatalf("Error listing GPU prices: %v", err)
	}

	fmt.Printf("Catalog generated at %s, %d prices\n", catalog.GeneratedAt, catalog.Count)
	for i, price := range catalog.Results {
		fmt.Printf("  %d. %dx %s (%s) in %s: %s %s/hour, %s %s/GPU/hour, SKU %s\n",
			i+1, price.GPUCount, price.GPUModel, price.DeploymentType, price.RegionName,
			price.PricePerHour, price.Currency, price.PricePerGPUHour, price.Currency,
			price.PricingDetail.SKU)
	}

	fmt.Println("=======================")
	return catalog
}

func filterGPUPrices(catalog *cloud.GPUPriceCatalog, currency cloud.GPUCurrency, deploymentType cloud.GPUDeploymentType) {
	fmt.Printf("\n=== %s GPU PRICES IN %s ===\n", deploymentType, currency)

	// the API has no filters, so filter the catalog on the client
	count := 0
	for _, price := range catalog.Results {
		if price.Currency != currency || price.DeploymentType != deploymentType {
			continue
		}
		count++
		// prices are decimal strings, so print them as is to keep their precision
		fmt.Printf("  %d. %dx %s in %s: %s/hour, flavors %v\n",
			count, price.GPUCount, price.GPUModel, price.RegionName, price.PricePerHour,
			price.PricingDetail.FlavorNames)
	}

	fmt.Println("=======================")
}
