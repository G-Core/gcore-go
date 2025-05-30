package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/shared/constant"
)

func createFileShare(ctx context.Context, name string, size int64, volumeType string) (string, error) {
	fmt.Println("\n=== CREATE FILE SHARE ===")
	client := gcore.NewClient()

	params := cloud.FileShareNewParams{
		OfCreateStandardFileShareSerializer: &cloud.FileShareNewParamsBodyCreateStandardFileShareSerializer{
			Name: name,
			Network: cloud.FileShareNewParamsBodyCreateStandardFileShareSerializerNetwork{
				NetworkID: "024a29e9-b4b7-4c91-9a46-505be123d9f8",
				SubnetID:  gcore.String("91200a6c-07e0-42aa-98da-32d1f6545ae7"),
			},
			Size:       size,
			VolumeType: volumeType,
			Protocol:   constant.Nfs("NFS"),
		},
	}

	taskList, err := client.Cloud.FileShares.New(ctx, params)
	if err != nil {
		return "", err
	}

	taskID := taskList.Tasks[0]
	task, err := client.Cloud.Tasks.Poll(ctx, taskID)
	if err != nil {
		return "", err
	}

	if len(task.CreatedResources.FileShares) == 0 {
		return "", fmt.Errorf("no file share created")
	}

	fileShareID := task.CreatedResources.FileShares[0]
	fmt.Println("Created File Share ID:", fileShareID)
	fmt.Println("====================")
	return fileShareID, nil
}

func getFileShare(ctx context.Context, fileShareID string) error {
	fmt.Println("\n=== GET FILE SHARE ===")
	client := gcore.NewClient()

	fileShare, err := client.Cloud.FileShares.Get(ctx, fileShareID, cloud.FileShareGetParams{})
	if err != nil {
		return err
	}

	fmt.Printf("File Share ID: %s, Name: %s, Size: %d GiB\n", fileShare.ID, fileShare.Name, fileShare.Size)
	fmt.Println("====================")
	return nil
}

func updateFileShare(ctx context.Context, fileShareID, newName string) error {
	fmt.Println("\n=== UPDATE FILE SHARE ===")
	client := gcore.NewClient()

	params := cloud.FileShareUpdateParams{
		Name: newName,
	}

	fileShare, err := client.Cloud.FileShares.Update(ctx, fileShareID, params)
	if err != nil {
		return err
	}

	fmt.Println("Updated file share name to:", fileShare.Name)
	fmt.Println("======================")
	return nil
}

func listFileShares(ctx context.Context) error {
	fmt.Println("\n=== LIST FILE SHARES ===")
	client := gcore.NewClient()

	fileShares, err := client.Cloud.FileShares.List(ctx, cloud.FileShareListParams{})
	if err != nil {
		return err
	}

	count := 0
	for _, fs := range fileShares.Results {
		count++
		fmt.Printf("%d. File Share ID: %s, Name: %s, Size: %d GiB\n", count, fs.ID, fs.Name, fs.Size)
	}

	if count == 0 {
		fmt.Println("No file shares found.")
	}
	fmt.Println("======================")
	return nil
}

func resizeFileShare(ctx context.Context, fileShareID string, newSize int64) error {
	fmt.Println("\n=== RESIZE FILE SHARE ===")
	client := gcore.NewClient()

	params := cloud.FileShareResizeParams{
		Size: newSize,
	}

	taskList, err := client.Cloud.FileShares.Resize(ctx, fileShareID, params)
	if err != nil {
		return err
	}

	taskID := taskList.Tasks[0]
	_, err = client.Cloud.Tasks.Poll(ctx, taskID)
	if err != nil {
		return err
	}

	fmt.Printf("Resized File Share %s to %d GiB\n", fileShareID, newSize)
	fmt.Println("======================")
	return nil
}

func deleteFileShare(ctx context.Context, fileShareID string) error {
	fmt.Println("\n=== DELETE FILE SHARE ===")
	client := gcore.NewClient()

	taskList, err := client.Cloud.FileShares.Delete(ctx, fileShareID, cloud.FileShareDeleteParams{})
	if err != nil {
		return err
	}

	if len(taskList.Tasks) > 0 {
		_, err := client.Cloud.Tasks.Poll(ctx, taskList.Tasks[0])
		if err != nil {
			return err
		}
	}

	fmt.Println("Deleted File Share", fileShareID)
	fmt.Println("======================")
	return nil
}

func main() {
	ctx := context.Background()

	fileShareName := "example-file-share"
	sizeGiB := int64(10)
	volumeType := "default_share_type" // or "vast_share_type"

	fileShareID, err := createFileShare(ctx, fileShareName, sizeGiB, volumeType)
	if err != nil {
		log.Fatalf("failed to create file share: %v", err)
	}

	if err = getFileShare(ctx, fileShareID); err != nil {
		log.Fatalf("failed to get file share: %v", err)
	}

	if err = updateFileShare(ctx, fileShareID, fileShareName+"-updated"); err != nil {
		log.Fatalf("failed to update file share: %v", err)
	}

	if err = listFileShares(ctx); err != nil {
		log.Fatalf("failed to list file shares: %v", err)
	}

	if err = resizeFileShare(ctx, fileShareID, 20); err != nil {
		log.Fatalf("failed to resize file share: %v", err)
	}

	if err = deleteFileShare(ctx, fileShareID); err != nil {
		log.Fatalf("failed to delete file share: %v", err)
	}

	fmt.Println("Example file share lifecycle complete.")
}
