package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/option"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	//apiKey := os.Getenv("GCORE_API_KEY")
	// Will use Production API URL if omitted
	//baseURL := os.Getenv("GCORE_BASE_URL")

	// TODO set cloud project ID before running
	cloudProjectID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_PROJECT_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_PROJECT_ID environment variable is required and must be a valid integer")
	}

	// TODO set cloud region ID before running
	cloudRegionID, err := strconv.ParseInt(os.Getenv("GCORE_CLOUD_REGION_ID"), 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_REGION_ID environment variable is required and must be a valid integer")
	}

	client := gcore.NewClient(
		//option.WithAPIKey(apiKey),
		//option.WithBaseURL(baseURL),
		option.WithCloudProjectID(cloudProjectID),
		option.WithCloudRegionID(cloudRegionID),
	)

	// Create a security group and use its ID for other operations
	sgID := createSecurityGroup(&client)
	listSecurityGroups(&client)
	getSecurityGroup(&client, sgID)
	updateSecurityGroup(&client, sgID)

	// Create and manage security group rules
	ruleID := createSecurityGroupRule(&client, sgID)
	ruleID = replaceSecurityGroupRule(&client, ruleID, sgID)
	deleteSecurityGroupRule(&client, ruleID)

	deleteSecurityGroup(&client, sgID)
}

func createSecurityGroup(client *gcore.Client) string {
	fmt.Println("\n=== CREATE SECURITY GROUP ===")

	sg, err := client.Cloud.SecurityGroups.New(context.Background(), cloud.SecurityGroupNewParams{
		SecurityGroup: cloud.SecurityGroupNewParamsSecurityGroup{
			Name:        "gcore-go-example",
			Description: gcore.String("gcore-go-example security group"),
		},
	})
	if err != nil {
		log.Fatalf("Error creating security group: %v", err)
	}

	fmt.Printf("Created Security Group ID: %s, Name: %s\n", sg.ID, sg.Name)
	fmt.Println("==============================")

	return sg.ID
}

func listSecurityGroups(client *gcore.Client) {
	fmt.Println("\n=== LIST SECURITY GROUPS ===")

	sgs, err := client.Cloud.SecurityGroups.List(context.Background(), cloud.SecurityGroupListParams{})
	if err != nil {
		log.Fatalf("Error listing security groups: %v", err)
	}

	for i, sg := range sgs.Results {
		fmt.Printf("  %d. Security Group ID: %s, Name: %s\n", i+1, sg.ID, sg.Name)
	}

	fmt.Println("=============================")
}

func getSecurityGroup(client *gcore.Client, groupID string) {
	fmt.Println("\n=== GET SECURITY GROUP BY ID ===")

	sg, err := client.Cloud.SecurityGroups.Get(context.Background(), groupID, cloud.SecurityGroupGetParams{})
	if err != nil {
		log.Fatalf("Error getting security group: %v", err)
	}

	fmt.Printf("Security Group ID: %s, Name: %s, Description: %s\n", sg.ID, sg.Name, sg.Description)
	fmt.Println("=================================")
}

func updateSecurityGroup(client *gcore.Client, groupID string) {
	fmt.Println("\n=== UPDATE SECURITY GROUP ===")

	sg, err := client.Cloud.SecurityGroups.Update(context.Background(), groupID, cloud.SecurityGroupUpdateParams{
		Name: gcore.String("gcore-go-example-updated"),
	})
	if err != nil {
		log.Fatalf("Error updating security group: %v", err)
	}

	fmt.Printf("Updated Security Group ID: %s, Name: %s\n", sg.ID, sg.Name)
	fmt.Println("==============================")
}

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

func deleteSecurityGroup(client *gcore.Client, groupID string) {
	fmt.Println("\n=== DELETE SECURITY GROUP ===")

	err := client.Cloud.SecurityGroups.Delete(context.Background(), groupID, cloud.SecurityGroupDeleteParams{})
	if err != nil {
		log.Fatalf("Error deleting security group: %v", err)
	}

	fmt.Printf("Security Group with ID %s successfully deleted\n", groupID)
	fmt.Println("==============================")
}
