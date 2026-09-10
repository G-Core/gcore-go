package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/packages/param"
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

func listPoolMembers(client *gcore.Client, poolID string) {
	fmt.Println("\n=== LIST POOL MEMBERS ===")

	params := cloud.LoadBalancerPoolMemberListParams{}
	members, err := client.Cloud.LoadBalancers.Pools.Members.List(context.Background(), poolID, params)
	if err != nil {
		log.Fatalf("Error listing pool members: %v", err)
	}

	for i, member := range members.Results {
		fmt.Printf("  %d. Member: ID=%s, Address=%s, Port=%d, OperatingStatus=%s\n",
			i+1, member.ID, member.Address, member.ProtocolPort, member.OperatingStatus)
	}

	if len(members.Results) == 0 {
		fmt.Println("  No members found.")
	}

	fmt.Println("=========================")
}

func getPoolMember(client *gcore.Client, memberID, poolID string) {
	fmt.Println("\n=== GET POOL MEMBER ===")

	params := cloud.LoadBalancerPoolMemberGetParams{
		PoolID: poolID,
	}

	member, err := client.Cloud.LoadBalancers.Pools.Members.Get(context.Background(), memberID, params)
	if err != nil {
		log.Fatalf("Error getting pool member: %v", err)
	}

	fmt.Printf("Member: ID=%s, Address=%s, Port=%d, Weight=%d, ProvisioningStatus=%s\n",
		member.ID, member.Address, member.ProtocolPort, member.Weight, member.ProvisioningStatus)
	fmt.Println("=======================")
}

func updatePoolMember(client *gcore.Client, memberID, poolID string) *cloud.Member {
	fmt.Println("\n=== UPDATE POOL MEMBER ===")

	params := cloud.LoadBalancerPoolMemberUpdateParams{
		PoolID: poolID,
		Weight: param.NewOpt(int64(2)),
	}

	member, err := client.Cloud.LoadBalancers.Pools.Members.UpdateAndPoll(context.Background(), memberID, params)
	if err != nil {
		log.Fatalf("Error updating pool member: %v", err)
	}

	fmt.Printf("Updated Member: ID=%s, Weight=%d, OperatingStatus=%s\n",
		member.ID, member.Weight, member.OperatingStatus)
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
