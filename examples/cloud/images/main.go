package main

import (
	"context"
	"fmt"
	"os"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func createImageFromVolume(ctx context.Context, volumeID string) (string, error) {
	client := gcore.NewClient()

	params := cloud.InstanceImageNewFromVolumeParams{
		VolumeID: volumeID,
		Name:     gcore.String("sdk-image-go-example").Value,
		OsType:   cloud.InstanceImageNewFromVolumeParamsOsType(gcore.String("linux").Value),
	}

	taskIDList, err := client.Cloud.Instances.Images.NewFromVolume(ctx, params)
	if err != nil {
		return "", err
	}

	if len(taskIDList.Tasks) != 1 {
		return "", fmt.Errorf("expected exactly one task created, got %d", len(taskIDList.Tasks))
	}

	task, err := client.Cloud.Tasks.Poll(ctx, taskIDList.Tasks[0])
	if err != nil {
		return "", err
	}

	if len(task.CreatedResources.Images) != 1 {
		return "", fmt.Errorf("expected exactly one image created, got %d", len(task.CreatedResources.Images))
	}

	return task.CreatedResources.Images[0], nil
}

func getImage(ctx context.Context, imageID string) (*cloud.Image, error) {
	client := gcore.NewClient()
	return client.Cloud.Instances.Images.Get(ctx, imageID, cloud.InstanceImageGetParams{})
}

func updateImage(ctx context.Context, imageID string) (*cloud.Image, error) {
	client := gcore.NewClient()
	params := cloud.InstanceImageUpdateParams{
		Name: gcore.String("sdk-image-go-example-updated"),
	}
	return client.Cloud.Instances.Images.Update(ctx, imageID, params)
}

func listImages(ctx context.Context) ([]cloud.Image, error) {
	client := gcore.NewClient()
	imageList, err := client.Cloud.Instances.Images.List(ctx, cloud.InstanceImageListParams{})
	if err != nil {
		return nil, err
	}
	return imageList.Results, nil
}

func deleteImage(ctx context.Context, imageID string) error {
	client := gcore.NewClient()
	taskList, err := client.Cloud.Instances.Images.Delete(ctx, imageID, cloud.InstanceImageDeleteParams{})
	if err != nil {
		return err
	}
	_, err = client.Cloud.Tasks.Poll(ctx, taskList.Tasks[0])
	return err
}

func main() {
	ctx := context.Background()

	// TODO: Set bootable volume ID before running
	volumeID := "your-bootable-volume-id"

	imageID, err := createImageFromVolume(ctx, volumeID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating image: %v\n", err)
		return
	}

	image, err := getImage(ctx, imageID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting image: %v\n", err)
		return
	}
	fmt.Printf("Image ID: %s, Name: %s, OS Type: %s, Status: %s\n", image.ID, image.Name, image.OsType, image.Status)

	updatedImage, err := updateImage(ctx, imageID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating image: %v\n", err)
		return
	}
	fmt.Printf("Updated image name to: %s\n", updatedImage.Name)

	images, err := listImages(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing images: %v\n", err)
		return
	}
	for i, img := range images {
		fmt.Printf("%d. Image ID: %s, Name: %s, OS Type: %s, Status: %s\n", i+1, img.ID, img.Name, img.OsType, img.Status)
	}

	err = deleteImage(ctx, imageID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting image: %v\n", err)
		return
	}
	fmt.Printf("Deleted image %s\n", imageID)
}
