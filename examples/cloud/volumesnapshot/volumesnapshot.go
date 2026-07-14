package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/shared/constant"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	//apiKey := os.Getenv("GCORE_API_KEY")
	// Will use Production API URL if omitted
	//baseURL := os.Getenv("GCORE_BASE_URL")

	// TODO set cloud project and region IDs before running
	cloudProjectID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_PROJECT_ID"), 10, 64)
	if err != nil {
		log.Fatalf("Error parsing GCORE_CLOUD_PROJECT_ID: %v", err)
	}
	cloudRegionID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_REGION_ID"), 10, 64)
	if err != nil {
		log.Fatalf("Error parsing GCORE_CLOUD_REGION_ID: %v", err)
	}

	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectID),
		option.WithCloudRegionID(cloudRegionID),
	)

	// A snapshot is taken of a volume, so create a volume to use as the source.
	volumeID := createSourceVolume(&client)
	snapshotID := createSnapshot(&client, volumeID)
	listSnapshots(&client)
	getSnapshot(&client, snapshotID)
	updateSnapshot(&client, snapshotID)
	deleteSnapshot(&client, snapshotID)
	deleteSourceVolume(&client, volumeID)
}

func createSourceVolume(client *gcore.Client) string {
	fmt.Println("\n=== CREATE SOURCE VOLUME ===")

	params := cloud.VolumeNewParams{
		OfNewVolume: &cloud.VolumeNewParamsBodyNewVolume{
			Name:   "gcore-go-example-snapshot-source",
			Size:   1,
			Source: constant.NewVolume("").Default(),
		},
	}

	volume, err := client.Cloud.Volumes.NewAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating source volume: %v", err)
	}

	fmt.Printf("Created source Volume ID: %s\n", volume.ID)
	fmt.Println("============================")

	return volume.ID
}

func createSnapshot(client *gcore.Client, volumeID string) string {
	fmt.Println("\n=== CREATE SNAPSHOT ===")

	params := cloud.VolumeSnapshotNewParams{
		Name:        "gcore-go-example-snapshot",
		VolumeID:    volumeID,
		Description: param.NewOpt("Created by the gcore-go volume snapshot example"),
	}

	// NewAndPoll creates the snapshot and waits for the underlying task to finish.
	snapshot, err := client.Cloud.VolumeSnapshots.NewAndPoll(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating snapshot: %v", err)
	}

	fmt.Printf("Created Snapshot ID: %s, Status: %s\n", snapshot.ID, snapshot.Status)
	fmt.Println("=======================")

	return snapshot.ID
}

func listSnapshots(client *gcore.Client) {
	fmt.Println("\n=== LIST SNAPSHOTS ===")

	snapshotsPage, err := client.Cloud.VolumeSnapshots.List(context.Background(), cloud.VolumeSnapshotListParams{})
	if err != nil {
		log.Fatalf("Error listing snapshots: %v", err)
	}

	for i, snapshot := range snapshotsPage.Results {
		fmt.Printf("  %d. Snapshot ID: %s, Name: %s, Size: %d GiB, Status: %s\n", i+1, snapshot.ID, snapshot.Name, snapshot.Size, snapshot.Status)
	}

	fmt.Println("======================")
}

func getSnapshot(client *gcore.Client, snapshotID string) {
	fmt.Println("\n=== GET SNAPSHOT BY ID ===")

	snapshot, err := client.Cloud.VolumeSnapshots.Get(context.Background(), snapshotID, cloud.VolumeSnapshotGetParams{})
	if err != nil {
		log.Fatalf("Error getting snapshot: %v", err)
	}

	fmt.Printf("Snapshot ID: %s, Name: %s, Volume ID: %s, Size: %d GiB\n", snapshot.ID, snapshot.Name, snapshot.VolumeID, snapshot.Size)
	fmt.Println("==========================")
}

func updateSnapshot(client *gcore.Client, snapshotID string) {
	fmt.Println("\n=== UPDATE SNAPSHOT ===")

	params := cloud.VolumeSnapshotUpdateParams{
		Name: param.NewOpt("gcore-go-example-snapshot-updated"),
	}

	snapshot, err := client.Cloud.VolumeSnapshots.Update(context.Background(), snapshotID, params)
	if err != nil {
		log.Fatalf("Error updating snapshot: %v", err)
	}

	fmt.Printf("Updated Snapshot ID: %s, Name: %s\n", snapshot.ID, snapshot.Name)
	fmt.Println("=======================")
}

func deleteSnapshot(client *gcore.Client, snapshotID string) {
	fmt.Println("\n=== DELETE SNAPSHOT ===")

	// DeleteAndPoll waits for the delete task to complete.
	err := client.Cloud.VolumeSnapshots.DeleteAndPoll(context.Background(), snapshotID, cloud.VolumeSnapshotDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting snapshot: %v", err)
	}

	fmt.Printf("Snapshot with ID %s successfully deleted\n", snapshotID)
	fmt.Println("=======================")
}

func deleteSourceVolume(client *gcore.Client, volumeID string) {
	fmt.Println("\n=== DELETE SOURCE VOLUME ===")

	err := client.Cloud.Volumes.DeleteAndPoll(context.Background(), volumeID, cloud.VolumeDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting source volume: %v", err)
	}

	fmt.Printf("Source volume with ID %s successfully deleted\n", volumeID)
	fmt.Println("============================")
}
