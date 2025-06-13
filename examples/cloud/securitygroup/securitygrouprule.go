package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func createSecurityGroupRule(client *gcore.Client, groupID string) string {
	fmt.Println("\n=== CREATE SECURITY GROUP RULE ===")

	httpPort := int64(80)
	remoteIP := "0.0.0.0/0"
	httpDesc := "Allow HTTP traffic"

	rule, err := client.Cloud.SecurityGroups.Rules.New(context.Background(), groupID, cloud.SecurityGroupRuleNewParams{
		Direction:      cloud.SecurityGroupRuleNewParamsDirectionIngress,
		Protocol:       cloud.SecurityGroupRuleNewParamsProtocolTcp,
		Ethertype:      cloud.SecurityGroupRuleNewParamsEthertypeIPv4,
		PortRangeMin:   gcore.Int(httpPort),
		PortRangeMax:   gcore.Int(httpPort),
		RemoteIPPrefix: gcore.String(remoteIP),
		Description:    gcore.String(httpDesc),
	})
	if err != nil {
		log.Fatalf("Error creating security group rule: %v", err)
	}

	fmt.Printf("Created Security Group Rule ID: %s, Protocol: %s, Port: %d\n", rule.ID, rule.Protocol, rule.PortRangeMin)
	fmt.Println("===================================")

	return rule.ID
}

func replaceSecurityGroupRule(client *gcore.Client, ruleID, groupID string) string {
	fmt.Println("\n=== REPLACE SECURITY GROUP RULE ===")

	httpsPort := int64(443)
	remoteIP := "0.0.0.0/0"
	httpsDesc := "Allow HTTPS traffic"

	rule, err := client.Cloud.SecurityGroups.Rules.Replace(context.Background(), ruleID, cloud.SecurityGroupRuleReplaceParams{
		Direction:       cloud.SecurityGroupRuleReplaceParamsDirectionIngress,
		SecurityGroupID: groupID,
		Protocol:        cloud.SecurityGroupRuleReplaceParamsProtocolTcp,
		Ethertype:       cloud.SecurityGroupRuleReplaceParamsEthertypeIPv4,
		PortRangeMin:    gcore.Int(httpsPort),
		PortRangeMax:    gcore.Int(httpsPort),
		RemoteIPPrefix:  gcore.String(remoteIP),
		Description:     gcore.String(httpsDesc),
	})
	if err != nil {
		log.Fatalf("Error replacing security group rule: %v", err)
	}

	fmt.Printf("Replaced Security Group Rule ID: %s, Protocol: %s, Port: %d\n", rule.ID, rule.Protocol, rule.PortRangeMin)
	fmt.Println("====================================")

	return rule.ID
}

func deleteSecurityGroupRule(client *gcore.Client, ruleID string) {
	fmt.Println("\n=== DELETE SECURITY GROUP RULE ===")

	err := client.Cloud.SecurityGroups.Rules.Delete(context.Background(), ruleID, cloud.SecurityGroupRuleDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting security group rule: %v", err)
	}

	fmt.Printf("Security Group Rule with ID %s successfully deleted\n", ruleID)
	fmt.Println("===================================")
}
