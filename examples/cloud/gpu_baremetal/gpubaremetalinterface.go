package main

import (
	"context"
	"fmt"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func listGPUBaremetalClusterInterfaces(client *gcore.Client, clusterID string) []cloud.NetworkInterface {
	fmt.Println("\n=== LIST GPU BAREMETAL CLUSTER INTERFACES ===")

	params := cloud.GPUBaremetalClusterInterfaceListParams{}
	interfaces, err := client.Cloud.GPUBaremetal.Clusters.Interfaces.List(context.Background(), clusterID, params)
	if err != nil {
		fmt.Printf("Error listing GPU baremetal cluster interfaces: %v\n", err)
		return nil
	}

	if len(interfaces.Results) == 0 {
		fmt.Println("No network interfaces found")
	} else {
		for i, iface := range interfaces.Results {
			fmt.Printf("  %d. Interface: PortID=%s\n", i+1, iface.PortID)
			if len(iface.IPAssignments) > 0 {
				for j, assignment := range iface.IPAssignments {
					fmt.Printf("     IP %d: %s\n", j+1, assignment.IPAddress)
				}
			}
			if iface.NetworkID != "" {
				fmt.Printf("     Network ID: %s\n", iface.NetworkID)
			}
			if iface.InterfaceName != "" {
				fmt.Printf("     Interface Name: %s\n", iface.InterfaceName)
			}
		}
	}

	fmt.Printf("Total interfaces: %d\n", len(interfaces.Results))
	fmt.Println("===============================")
	return interfaces.Results
}
