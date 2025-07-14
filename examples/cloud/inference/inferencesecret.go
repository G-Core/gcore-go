package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func createSecret(client *gcore.Client) string {
	fmt.Println("\n=== CREATE SECRET ===")

	params := cloud.InferenceSecretNewParams{
		Name: "gcore-go-example",
		Type: "aws-iam",
		Data: cloud.InferenceSecretNewParamsData{
			AwsAccessKeyID:     "example-key",
			AwsSecretAccessKey: "example-secret",
		},
	}

	secret, err := client.Cloud.Inference.Secrets.New(context.Background(), params)
	if err != nil {
		fmt.Printf("Error creating secret: %v\n", err)
		return ""
	}

	fmt.Printf("Created secret: %s\n", secret.Name)
	fmt.Println("========================")
	return secret.Name
}

func listSecrets(client *gcore.Client) {
	fmt.Println("\n=== LIST SECRETS ===")

	iter := client.Cloud.Inference.Secrets.ListAutoPaging(context.Background(), cloud.InferenceSecretListParams{})
	count := 0
	for iter.Next() {
		count++
		secret := iter.Current()
		fmt.Printf("  %d. Secret: %s, type: %s\n", count, secret.Name, secret.Type)
	}

	if err := iter.Err(); err != nil {
		fmt.Printf("Error iterating secrets: %v\n", err)
		return
	}

	fmt.Println("========================")
}

func getSecret(client *gcore.Client, secretName string) {
	fmt.Println("\n=== GET SECRET ===")

	secret, err := client.Cloud.Inference.Secrets.Get(context.Background(), secretName, cloud.InferenceSecretGetParams{})
	if err != nil {
		fmt.Printf("Error getting secret: %v\n", err)
		return
	}

	fmt.Printf("Secret: %s, type: %s\n", secret.Name, secret.Type)
	fmt.Println("========================")
}

func replaceSecret(client *gcore.Client, secretName string) {
	fmt.Println("\n=== REPLACE SECRET ===")

	params := cloud.InferenceSecretReplaceParams{
		Type: "aws-iam",
		Data: cloud.InferenceSecretReplaceParamsData{
			AwsAccessKeyID:     "updated-key",
			AwsSecretAccessKey: "updated-secret",
		},
	}

	secret, err := client.Cloud.Inference.Secrets.Replace(context.Background(), secretName, params)
	if err != nil {
		fmt.Printf("Error replacing secret: %v\n", err)
		return
	}

	fmt.Printf("Replaced secret: %s\n", secret.Name)
	fmt.Println("========================")
}

func deleteSecret(client *gcore.Client, secretName string) {
	fmt.Println("\n=== DELETE SECRET ===")

	err := client.Cloud.Inference.Secrets.Delete(context.Background(), secretName, cloud.InferenceSecretDeleteParams{})
	if err != nil {
		fmt.Printf("Error deleting secret: %v\n", err)
		return
	}

	fmt.Printf("Deleted secret: %s\n", secretName)
	fmt.Println("========================")
}
