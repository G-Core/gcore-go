package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func createNetwork(ctx context.Context, name string) (string, error) {
	fmt.Println("\n=== CREATE NETWORK ===")
	client := gcore.NewClient()
	params := cloud.NetworkNewParams{
		Name:         name,
		CreateRouter: gcore.Bool(true),
		Type:         cloud.NetworkNewParamsTypeVxlan,
	}

	taskIDList, err := client.Cloud.Networks.New(ctx, params)
	if err != nil {
		return "", err
	}

	taskID := taskIDList.Tasks[0]
	task, err := client.Cloud.Tasks.Poll(ctx, taskID)
	if err != nil {
		return "", err
	}

	if len(task.CreatedResources.Networks) == 0 {
		return "", fmt.Errorf("no network created")
	}

	networkID := task.CreatedResources.Networks[0]
	fmt.Println("Created Network ID:", networkID)
	fmt.Println("====================")
	return networkID, nil
}

func getNetwork(ctx context.Context, networkID string) error {
	fmt.Println("\n=== GET NETWORK ===")
	client := gcore.NewClient()
	network, err := client.Cloud.Networks.Get(ctx, networkID, cloud.NetworkGetParams{})
	if err != nil {
		return err
	}

	fmt.Printf("Network ID: %s, Name: %s, Type: %s\n", network.ID, network.Name, network.Type)
	fmt.Println("===================")
	return nil
}

func updateNetwork(ctx context.Context, networkID, newName string) error {
	fmt.Println("\n=== UPDATE NETWORK ===")
	client := gcore.NewClient()
	params := cloud.NetworkUpdateParams{
		Name: newName,
	}

	network, err := client.Cloud.Networks.Update(ctx, networkID, params)
	if err != nil {
		return err
	}

	fmt.Println("Updated network name to:", network.Name)
	fmt.Println("========================")
	return nil
}

func deleteNetwork(ctx context.Context, networkID string) error {
	fmt.Println("\n=== DELETE NETWORK ===")
	client := gcore.NewClient()
	params := cloud.NetworkDeleteParams{}
	taskIDList, err := client.Cloud.Networks.Delete(ctx, networkID, params)
	if err != nil {
		return err
	}

	if len(taskIDList.Tasks) > 0 {
		_, err = client.Cloud.Tasks.Poll(ctx, taskIDList.Tasks[0])
		if err != nil {
			return err
		}
	}

	fmt.Println("Deleted Network", networkID)
	fmt.Println("====================")
	return nil
}

func listNetworks(ctx context.Context) error {
	fmt.Println("\n=== LIST NETWORKS ===")
	client := gcore.NewClient()
	networks, err := client.Cloud.Networks.List(ctx, cloud.NetworkListParams{})
	if err != nil {
		return err
	}

	count := 0
	for _, network := range networks.Results {
		count++
		fmt.Printf("%d. Network ID: %s, Name: %s, Type: %s\n", count, network.ID, network.Name, network.Type)
	}

	if count == 0 {
		fmt.Println("No networks found.")
	}
	fmt.Println("=====================")
	return nil
}

func createSubnet(ctx context.Context, networkID, cidr, name string) (string, error) {
	fmt.Println("\n=== CREATE SUBNET ===")
	client := gcore.NewClient()
	params := cloud.NetworkSubnetNewParams{
		NetworkID:  networkID,
		Cidr:       cidr,
		Name:       name,
		EnableDhcp: gcore.Bool(true),
		IPVersion:  4,
	}

	taskIDList, err := client.Cloud.Networks.Subnets.New(ctx, params)
	if err != nil {
		return "", err
	}

	taskID := taskIDList.Tasks[0]
	task, err := client.Cloud.Tasks.Poll(ctx, taskID)
	if err != nil {
		return "", err
	}

	if len(task.CreatedResources.Subnets) == 0 {
		return "", fmt.Errorf("no subnet created")
	}

	subnetID := task.CreatedResources.Subnets[0]
	fmt.Println("Created Subnet ID:", subnetID)
	fmt.Println("===================")
	return subnetID, nil
}

func getSubnet(ctx context.Context, subnetID string) error {
	fmt.Println("\n=== GET SUBNET ===")
	client := gcore.NewClient()
	subnet, err := client.Cloud.Networks.Subnets.Get(ctx, subnetID, cloud.NetworkSubnetGetParams{})
	if err != nil {
		return err
	}

	fmt.Printf("Subnet ID: %s, CIDR: %s, Name: %s\n", subnet.ID, subnet.Cidr, subnet.Name)
	fmt.Println("==================")
	return nil
}

func updateSubnet(ctx context.Context, subnetID, newName string) error {
	fmt.Println("\n=== UPDATE SUBNET ===")
	client := gcore.NewClient()
	params := cloud.NetworkSubnetUpdateParams{
		Name: gcore.String(newName),
	}

	subnet, err := client.Cloud.Networks.Subnets.Update(ctx, subnetID, params)
	if err != nil {
		return err
	}

	fmt.Println("Updated subnet name to:", subnet.Name)
	fmt.Println("=======================")
	return nil
}

func deleteSubnet(ctx context.Context, subnetID string) error {
	fmt.Println("\n=== DELETE SUBNET ===")
	client := gcore.NewClient()
	params := cloud.NetworkSubnetDeleteParams{}
	err := client.Cloud.Networks.Subnets.Delete(ctx, subnetID, params)
	if err != nil {
		return err
	}

	fmt.Println("Deleted Subnet", subnetID)
	fmt.Println("===================")
	return nil
}

func listSubnets(ctx context.Context, networkID string) error {
	fmt.Println("\n=== LIST SUBNETS ===")
	client := gcore.NewClient()
	subnets, err := client.Cloud.Networks.Subnets.List(ctx, cloud.NetworkSubnetListParams{NetworkID: gcore.String(networkID)})
	if err != nil {
		return err
	}

	count := 0
	for _, subnet := range subnets.Results {
		count++
		fmt.Printf("%d. Subnet ID: %s, CIDR: %s, Name: %s\n", count, subnet.ID, subnet.Cidr, subnet.Name)
	}

	if count == 0 {
		fmt.Println("No subnets found.")
	}
	fmt.Println("====================")
	return nil
}

func createRouter(ctx context.Context) (string, error) {
	fmt.Println("\n=== CREATE ROUTER ===")
	client := gcore.NewClient()
	params := cloud.NetworkRouterNewParams{
		Name: "example-router",
	}

	taskIDList, err := client.Cloud.Networks.Routers.New(ctx, params)
	if err != nil {
		return "", err
	}

	taskID := taskIDList.Tasks[0]
	task, err := client.Cloud.Tasks.Poll(ctx, taskID)
	if err != nil {
		return "", err
	}

	if len(task.CreatedResources.Routers) == 0 {
		return "", fmt.Errorf("no router created")
	}

	routerID := task.CreatedResources.Routers[0]
	fmt.Println("Created Router ID:", routerID)
	fmt.Println("===================")
	return routerID, nil
}

func getRouter(ctx context.Context, routerID string) error {
	fmt.Println("\n=== GET ROUTER ===")
	client := gcore.NewClient()
	router, err := client.Cloud.Networks.Routers.Get(ctx, routerID, cloud.NetworkRouterGetParams{})
	if err != nil {
		return err
	}

	fmt.Printf("Router ID: %s, Name: %s\n", router.ID, router.Name)
	fmt.Println("==================")
	return nil
}

func updateRouter(ctx context.Context, routerID, newName string) error {
	fmt.Println("\n=== UPDATE ROUTER ===")
	client := gcore.NewClient()
	params := cloud.NetworkRouterUpdateParams{
		Name: gcore.String(newName),
	}

	router, err := client.Cloud.Networks.Routers.Update(ctx, routerID, params)
	if err != nil {
		return err
	}

	fmt.Println("Updated router name to:", router.Name)
	fmt.Println("=======================")
	return nil
}

func deleteRouter(ctx context.Context, routerID string) error {
	fmt.Println("\n=== DELETE ROUTER ===")
	client := gcore.NewClient()
	params := cloud.NetworkRouterDeleteParams{}
	taskIDList, err := client.Cloud.Networks.Routers.Delete(ctx, routerID, params)
	if err != nil {
		return err
	}

	if len(taskIDList.Tasks) > 0 {
		_, err = client.Cloud.Tasks.Poll(ctx, taskIDList.Tasks[0])
		if err != nil {
			return err
		}
	}

	fmt.Println("Deleted Router", routerID)
	fmt.Println("===================")
	return nil
}

func listRouters(ctx context.Context) error {
	fmt.Println("\n=== LIST ROUTERS ===")
	client := gcore.NewClient()
	params := cloud.NetworkRouterListParams{
		Limit: gcore.Int(100),
	}
	pager := client.Cloud.Networks.Routers.ListAutoPaging(ctx, params)
	count := 0
	for pager.Next() {
		router := pager.Current()
		count++
		fmt.Printf("%d. Router ID: %s, Name: %s\n", count, router.ID, router.Name)
	}
	if count == 0 {
		fmt.Println("No routers found.")
	}
	fmt.Println("=====================")
	return pager.Err()
}

func main() {
	ctx := context.Background()

	networkName := "example-network"
	networkID, err := createNetwork(ctx, networkName)
	if err != nil {
		log.Fatalf("failed to create network: %v", err)
	}

	err = getNetwork(ctx, networkID)
	if err != nil {
		log.Fatalf("failed to get network: %v", err)
	}

	err = updateNetwork(ctx, networkID, networkName+"-updated")
	if err != nil {
		log.Fatalf("failed to update network: %v", err)
	}

	err = listNetworks(ctx)
	if err != nil {
		log.Fatalf("failed to list networks: %v", err)
	}

	err = listRouters(ctx)
	if err != nil {
		log.Fatalf("failed to list routers: %v", err)
	}

	subnetCIDR := "192.168.123.0/24"
	subnetName := "example-subnet"

	subnetID, err := createSubnet(ctx, networkID, subnetCIDR, subnetName)
	if err != nil {
		log.Fatalf("failed to create subnet: %v", err)
	}

	err = getSubnet(ctx, subnetID)
	if err != nil {
		log.Fatalf("failed to get subnet: %v", err)
	}

	err = updateSubnet(ctx, subnetID, subnetName+"-updated")
	if err != nil {
		log.Fatalf("failed to update subnet: %v", err)
	}

	err = listSubnets(ctx, networkID)
	if err != nil {
		log.Fatalf("failed to list subnets: %v", err)
	}

	routerID, err := createRouter(ctx)
	if err != nil {
		log.Fatalf("failed to create router: %v", err)
	}

	err = getRouter(ctx, routerID)
	if err != nil {
		log.Fatalf("failed to get router: %v", err)
	}

	err = updateRouter(ctx, routerID, "example-router-updated")
	if err != nil {
		log.Fatalf("failed to update router: %v", err)
	}

	err = listSubnets(ctx, networkID)
	if err != nil {
		log.Fatalf("failed to list subnets: %v", err)
	}

	err = deleteSubnet(ctx, subnetID)
	if err != nil {
		log.Fatalf("failed to delete subnet: %v", err)
	}

	err = deleteRouter(ctx, routerID)
	if err != nil {
		log.Fatalf("failed to delete router: %v", err)
	}

	err = deleteNetwork(ctx, networkID)
	if err != nil {
		log.Fatalf("failed to delete network: %v", err)
	}

	fmt.Println("Example network lifecycle complete")
}
