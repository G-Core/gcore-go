package main

import (
	"context"
	"fmt"
	"log"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
)

func createFileShareAccessRule(client *gcore.Client, fileShareID string) *cloud.AccessRule {
	fmt.Println("\n=== CREATE FILE SHARE ACCESS RULE ===")

	params := cloud.FileShareAccessRuleNewParams{
		AccessMode: cloud.FileShareAccessRuleNewParamsAccessModeRo,
		IPAddress:  "192.168.1.0/24",
	}

	accessRule, err := client.Cloud.FileShares.AccessRules.New(context.Background(), fileShareID, params)
	if err != nil {
		log.Fatalf("Error creating file share access rule: %v", err)
	}

	fmt.Printf("Created Access Rule: ID=%s, AccessLevel=%s, AccessTo=%s, State=%s\n",
		accessRule.ID, accessRule.AccessLevel, accessRule.AccessTo, accessRule.State)
	fmt.Println("=====================================")

	return accessRule
}

func listFileShareAccessRules(client *gcore.Client, fileShareID string) {
	fmt.Println("\n=== LIST FILE SHARE ACCESS RULES ===")

	accessRules, err := client.Cloud.FileShares.AccessRules.List(context.Background(), fileShareID, cloud.FileShareAccessRuleListParams{})
	if err != nil {
		log.Fatalf("Error listing file share access rules: %v", err)
	}

	for i, rule := range accessRules.Results {
		fmt.Printf("  %d. Access Rule: ID=%s, AccessLevel=%s, AccessTo=%s, State=%s\n",
			i+1, rule.ID, rule.AccessLevel, rule.AccessTo, rule.State)
	}

	fmt.Println("====================================")
}

func deleteFileShareAccessRule(client *gcore.Client, fileShareID, accessRuleID string) {
	fmt.Println("\n=== DELETE FILE SHARE ACCESS RULE ===")

	params := cloud.FileShareAccessRuleDeleteParams{
		FileShareID: fileShareID,
	}

	err := client.Cloud.FileShares.AccessRules.Delete(context.Background(), accessRuleID, params)
	if err != nil {
		log.Fatalf("Error deleting file share access rule: %v", err)
	}

	fmt.Printf("File Share Access Rule %s successfully deleted\n", accessRuleID)
	fmt.Println("=====================================")
}
