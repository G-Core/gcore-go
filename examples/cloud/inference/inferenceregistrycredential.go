package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func createRegistryCredential(client *gcore.Client) string {
	fmt.Println("\n=== CREATE REGISTRY CREDENTIAL ===")

	params := cloud.InferenceRegistryCredentialNewParams{
		Name:        "gcore-go-example",
		Username:    "example-user",
		Password:    "example-password",
		RegistryURL: "https://registry.example.com",
	}

	credential, err := client.Cloud.Inference.RegistryCredentials.New(context.Background(), params)
	if err != nil {
		fmt.Printf("Error creating registry credential: %v\n", err)
		return ""
	}

	fmt.Printf("Created registry credential: %s\n", credential.Name)
	fmt.Println("========================")
	return credential.Name
}

func listRegistryCredentials(client *gcore.Client) {
	fmt.Println("\n=== LIST REGISTRY CREDENTIALS ===")

	iter := client.Cloud.Inference.RegistryCredentials.ListAutoPaging(context.Background(), cloud.InferenceRegistryCredentialListParams{})
	count := 0
	for iter.Next() {
		count++
		credential := iter.Current()
		fmt.Printf("  %d. Registry credential: %s, URL: %s\n", count, credential.Name, credential.RegistryURL)
	}

	if err := iter.Err(); err != nil {
		fmt.Printf("Error iterating registry credentials: %v\n", err)
		return
	}

	fmt.Println("========================")
}

func getRegistryCredential(client *gcore.Client, credentialName string) {
	fmt.Println("\n=== GET REGISTRY CREDENTIAL ===")

	credential, err := client.Cloud.Inference.RegistryCredentials.Get(context.Background(), credentialName, cloud.InferenceRegistryCredentialGetParams{})
	if err != nil {
		fmt.Printf("Error getting registry credential: %v\n", err)
		return
	}

	fmt.Printf("Registry credential: %s, URL: %s\n", credential.Name, credential.RegistryURL)
	fmt.Println("========================")
}

func replaceRegistryCredential(client *gcore.Client, credentialName string) {
	fmt.Println("\n=== REPLACE REGISTRY CREDENTIAL ===")

	params := cloud.InferenceRegistryCredentialReplaceParams{
		Username:    "updated-user",
		Password:    "updated-password",
		RegistryURL: "https://updated-registry.example.com",
	}

	credential, err := client.Cloud.Inference.RegistryCredentials.Replace(context.Background(), credentialName, params)
	if err != nil {
		fmt.Printf("Error replacing registry credential: %v\n", err)
		return
	}

	fmt.Printf("Replaced registry credential: %s\n", credential.Name)
	fmt.Println("========================")
}

func deleteRegistryCredential(client *gcore.Client, credentialName string) {
	fmt.Println("\n=== DELETE REGISTRY CREDENTIAL ===")

	err := client.Cloud.Inference.RegistryCredentials.Delete(context.Background(), credentialName, cloud.InferenceRegistryCredentialDeleteParams{})
	if err != nil {
		fmt.Printf("Error deleting registry credential: %v\n", err)
		return
	}

	fmt.Printf("Deleted registry credential: %s\n", credentialName)
	fmt.Println("========================")
}
