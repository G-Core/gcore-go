package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/packages/param"
)

func listInterfaces(client *gcore.Client, instanceID string) []cloud.NetworkInterface {
	fmt.Println("\n=== LIST INTERFACES ===")

	interfaces, err := client.Cloud.Instances.Interfaces.List(context.Background(), instanceID, cloud.InstanceInterfaceListParams{})
	if err != nil {
		fmt.Printf("Error listing interfaces: %v\n", err)
		return nil
	}

	for i, iface := range interfaces.Results {
		fmt.Printf("  %d. Interface: PortID=%s, NetworkID=%s\n", i+1, iface.PortID, iface.NetworkID)
	}

	fmt.Println("=======================")
	return interfaces.Results
}

func attachInterface(client *gcore.Client, instanceID string, networkID string) []cloud.NetworkInterface {
	fmt.Println("\n=== ATTACH INTERFACE ===")

	// Attach interface to any available subnet in the specified network
	params := cloud.InstanceInterfaceAttachParams{
		OfNewInterfaceAnySubnetSchema: &cloud.InstanceInterfaceAttachParamsBodyNewInterfaceAnySubnetSchema{
			Type:      param.Opt[string]{Value: "any_subnet"},
			NetworkID: networkID,
		},
	}

	interfaces, err := client.Cloud.Instances.Interfaces.AttachAndPoll(context.Background(), instanceID, params)
	if err != nil {
		fmt.Printf("Error attaching interface to any subnet in network %s: %v\n", networkID, err)
		return nil
	}

	fmt.Printf("Attached interface to any available subnet in network %s (instance: %s)\n", networkID, instanceID)
	fmt.Println("========================")
	return interfaces.Results
}

func detachInterface(client *gcore.Client, instanceID, ipAddress, portID string) {
	fmt.Println("\n=== DETACH INTERFACE ===")

	params := cloud.InstanceInterfaceDetachParams{
		IPAddress: ipAddress,
		PortID:    portID,
	}

	_, err := client.Cloud.Instances.Interfaces.DetachAndPoll(context.Background(), instanceID, params)
	if err != nil {
		fmt.Printf("Error detaching interface (IP: %s, Port: %s): %v\n", ipAddress, portID, err)
		return
	}

	fmt.Printf("Detached interface (IP: %s, Port: %s) from instance: %s\n", ipAddress, portID, instanceID)
	fmt.Println("========================")
}
