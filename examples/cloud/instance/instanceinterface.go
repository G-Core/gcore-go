package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listInterfaces(client *gcore.Client, instanceID string) []cloud.NetworkInterfaceUnion {
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

func attachInterface(client *gcore.Client, instanceID string, networkID string) []cloud.NetworkInterfaceUnion {
	fmt.Println("\n=== ATTACH INTERFACE ===")

	// Attach interface to any available subnet in the specified network
	params := cloud.InstanceInterfaceAttachParams{
		OfAnySubnet: &cloud.InstanceInterfaceAttachParamsBodyAnySubnet{
			Type:      "any_subnet",
			NetworkID: networkID,
		},
	}

	if err := client.Cloud.Instances.Interfaces.AttachAndPoll(context.Background(), instanceID, params); err != nil {
		fmt.Printf("Error attaching interface to any subnet in network %s: %v\n", networkID, err)
		return nil
	}

	fmt.Printf("Attached interface to any available subnet in network %s (instance: %s)\n", networkID, instanceID)
	fmt.Println("========================")

	// AttachAndPoll returns only tasks; list the interfaces to read the current state.
	return listInterfaces(client, instanceID)
}

func detachInterface(client *gcore.Client, instanceID, ipAddress, portID string) {
	fmt.Println("\n=== DETACH INTERFACE ===")

	params := cloud.InstanceInterfaceDetachParams{
		IPAddress: ipAddress,
		PortID:    portID,
	}

	if err := client.Cloud.Instances.Interfaces.DetachAndPoll(context.Background(), instanceID, params); err != nil {
		fmt.Printf("Error detaching interface (IP: %s, Port: %s): %v\n", ipAddress, portID, err)
		return
	}

	fmt.Printf("Detached interface (IP: %s, Port: %s) from instance: %s\n", ipAddress, portID, instanceID)
	fmt.Println("========================")
}
