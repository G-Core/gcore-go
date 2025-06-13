package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func createSubnet(client *gcore.Client, networkID string) string {
	fmt.Println("\n=== CREATE SUBNET ===")
	cidr := "192.168.123.0/24"
	name := "gcore-go-example"
	params := cloud.NetworkSubnetNewParams{
		NetworkID:  networkID,
		Cidr:       cidr,
		Name:       name,
		EnableDhcp: gcore.Bool(true),
		IPVersion:  4,
		Tags:       map[string]string{"name": "gcore-go-example"},
	}

	taskIDList, err := client.Cloud.Networks.Subnets.New(context.Background(), params)
	if err != nil {
		log.Fatalf("Error creating subnet: %v", err)
	}

	task, err := client.Cloud.Tasks.Poll(context.Background(), taskIDList.Tasks[0])
	if err != nil {
		log.Fatalf("Error polling task for subnet creation: %v", err)
	}

	if len(task.CreatedResources.Subnets) == 0 {
		log.Fatalf("No subnet created")
	}

	subnetID := task.CreatedResources.Subnets[0]
	fmt.Printf("Created Subnet ID: %s\n", subnetID)
	fmt.Println("===================")
	return subnetID
}

func getSubnet(client *gcore.Client, subnetID string) {
	fmt.Println("\n=== GET SUBNET ===")
	subnet, err := client.Cloud.Networks.Subnets.Get(context.Background(), subnetID, cloud.NetworkSubnetGetParams{})
	if err != nil {
		log.Fatalf("Error getting subnet: %v", err)
	}

	fmt.Printf("Subnet ID: %s, CIDR: %s, Name: %s\n", subnet.ID, subnet.Cidr, subnet.Name)
	fmt.Println("==================")
}

func updateSubnet(client *gcore.Client, subnetID string) {
	fmt.Println("\n=== UPDATE SUBNET ===")
	newName := "gcore-go-example-updated"
	params := cloud.NetworkSubnetUpdateParams{
		Name: gcore.String(newName),
	}

	subnet, err := client.Cloud.Networks.Subnets.Update(context.Background(), subnetID, params)
	if err != nil {
		log.Fatalf("Error updating subnet: %v", err)
	}

	fmt.Printf("Updated subnet name to: %s\n", subnet.Name)
	fmt.Println("=======================")
}

func listSubnets(client *gcore.Client, networkID string) {
	fmt.Println("\n=== LIST SUBNETS ===")
	subnets, err := client.Cloud.Networks.Subnets.List(context.Background(), cloud.NetworkSubnetListParams{NetworkID: gcore.String(networkID)})
	if err != nil {
		log.Fatalf("Error listing subnets: %v", err)
	}

	for i, subnet := range subnets.Results {
		fmt.Printf("  %d. Subnet ID: %s, CIDR: %s, Name: %s\n", i+1, subnet.ID, subnet.Cidr, subnet.Name)
	}

	if len(subnets.Results) == 0 {
		fmt.Println("  No subnets found.")
	}
	fmt.Println("====================")
}

func deleteSubnet(client *gcore.Client, subnetID string) {
	fmt.Println("\n=== DELETE SUBNET ===")
	params := cloud.NetworkSubnetDeleteParams{}
	err := client.Cloud.Networks.Subnets.Delete(context.Background(), subnetID, params)
	if err != nil {
		log.Fatalf("Error deleting subnet: %v", err)
	}

	fmt.Printf("Subnet with ID %s successfully deleted\n", subnetID)
	fmt.Println("===================")
}
