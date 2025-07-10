package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listModels(client *gcore.Client) {
	fmt.Println("\n=== LIST MODELS ===")

	iter := client.Cloud.Inference.Models.ListAutoPaging(context.Background(), cloud.InferenceModelListParams{})
	modelList := make([]cloud.InferenceModel, 0)
	for iter.Next() {
		modelList = append(modelList, iter.Current())
	}

	if err := iter.Err(); err != nil {
		fmt.Printf("Error iterating models: %v\n", err)
		return
	}

	displayCount := 3
	if len(modelList) < displayCount {
		displayCount = len(modelList)
	}

	for i := 0; i < displayCount; i++ {
		model := modelList[i]
		fmt.Printf("  %d. Model: %s, ID: %s\n", i+1, model.Name, model.ID)
	}

	if len(modelList) > displayCount {
		fmt.Printf("  ... and %d more models\n", len(modelList)-displayCount)
	}

	fmt.Println("========================")
}

func getModel(client *gcore.Client, modelID string) {
	fmt.Println("\n=== GET MODEL ===")

	model, err := client.Cloud.Inference.Models.Get(context.Background(), modelID)
	if err != nil {
		fmt.Printf("Error getting model: %v\n", err)
		return
	}

	fmt.Printf("Model: %s, ID: %s\n", model.Name, model.ID)
	fmt.Println("========================")
}
