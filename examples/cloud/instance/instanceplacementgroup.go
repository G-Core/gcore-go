package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func addToPlacementGroup(client *gcore.Client, instanceID, groupID string) {
	fmt.Println("\n=== ADD TO PLACEMENT GROUP ===")

	addParams := cloud.InstanceAddToPlacementGroupParams{ServergroupID: groupID}
	if _, err := client.Cloud.Instances.AddToPlacementGroupAndPoll(context.Background(), instanceID, addParams); err != nil {
		fmt.Printf("Error adding to placement group: %v\n", err)
		return
	}

	fmt.Printf("Added instance %s to placement group: %s\n", instanceID, groupID)
	fmt.Println("===============================")
}

func removeFromPlacementGroup(client *gcore.Client, instanceID string) {
	fmt.Println("\n=== REMOVE FROM PLACEMENT GROUP ===")

	removeParams := cloud.InstanceRemoveFromPlacementGroupParams{}
	if _, err := client.Cloud.Instances.RemoveFromPlacementGroupAndPoll(context.Background(), instanceID, removeParams); err != nil {
		fmt.Printf("Error removing from placement group: %v\n", err)
		return
	}

	fmt.Printf("Removed instance %s from placement group\n", instanceID)
	fmt.Println("====================================")
}
