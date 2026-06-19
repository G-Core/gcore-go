package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func createPoolMember(client *gcore.Client, poolID string) *cloud.Member {
	fmt.Println("\n=== CREATE POOL MEMBER ===")

	params := cloud.LoadBalancerPoolMemberNewParams{
		Address:      "192.168.1.10",
		ProtocolPort: 80,
	}

	member, err := client.Cloud.LoadBalancers.Pools.Members.NewAndPoll(context.Background(), poolID, params)
	if err != nil {
		log.Fatalf("Error creating pool member: %v", err)
	}

	fmt.Printf("Created Member: ID=%s, Address=%s, Port=%d, OperatingStatus=%s\n",
		member.ID, member.Address, member.ProtocolPort, member.OperatingStatus)
	fmt.Println("==========================")

	return member
}

func deletePoolMember(client *gcore.Client, memberID, poolID string) {
	fmt.Println("\n=== DELETE POOL MEMBER ===")

	params := cloud.LoadBalancerPoolMemberDeleteParams{
		PoolID: poolID,
	}

	err := client.Cloud.LoadBalancers.Pools.Members.DeleteAndPoll(context.Background(), memberID, params)
	if err != nil {
		log.Fatalf("Error deleting pool member: %v", err)
	}

	fmt.Printf("Pool member with ID %s successfully deleted\n", memberID)
	fmt.Println("==========================")
}
