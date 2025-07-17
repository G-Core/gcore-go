package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

func main() {
	// No need to pass the API key explicitly — it will automatically be read from the GCORE_API_KEY environment variable if omitted
	// Will use Production API URL if omitted

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
		option.WithCloudProjectID(cloudProjectID),
		option.WithCloudRegionID(cloudRegionID),
	)

	listAuditLogs(&client)
	listAuditLogsAutoPaging(&client)
}

func listAuditLogs(client *gcore.Client) {
	fmt.Println("\n=== LIST AUDIT LOGS (first page) ===")

	params := cloud.AuditLogListParams{
		Limit:   param.NewOpt(int64(5)),              // Limit to 5 results for the example
		OrderBy: cloud.AuditLogListParamsOrderByDesc, // Most recent first
		// Example: filter by action type (uncomment to use)
		//ActionType: []string{"create"},
		// Example: filter by time range (uncomment to use)
		//FromTimestamp: param.NewOpt(time.Now().Add(-24 * time.Hour)),
	}

	page, err := client.Cloud.AuditLogs.List(context.Background(), params)
	if err != nil {
		log.Fatalf("Error listing audit logs: %v", err)
	}

	for i, logEntry := range page.Results {
		fmt.Printf("  %d. ID=%s, ActionType=%s, APIGroup=%s, Email=%s, Timestamp=%s\n",
			i+1, logEntry.ID, logEntry.ActionType, logEntry.APIGroup, logEntry.Email, logEntry.Timestamp.Format(time.RFC3339))
	}

	fmt.Println("=====================================")
}

func listAuditLogsAutoPaging(client *gcore.Client) {
	fmt.Println("\n=== LIST AUDIT LOGS USING AUTOPAGER ===")

	params := cloud.AuditLogListParams{
		Limit:   param.NewOpt(int64(3)),             // Process 3 items per page for demonstration
		OrderBy: cloud.AuditLogListParamsOrderByAsc, // Oldest first
	}

	iter := client.Cloud.AuditLogs.ListAutoPaging(context.Background(), params)
	count := 0
	for iter.Next() {
		logEntry := iter.Current()
		count++
		fmt.Printf("  %d. ID=%s, ActionType=%s, APIGroup=%s, Email=%s, Timestamp=%s\n",
			count, logEntry.ID, logEntry.ActionType, logEntry.APIGroup, logEntry.Email, logEntry.Timestamp.Format(time.RFC3339))
		if count >= 10 {
			fmt.Println("  ... (truncated for example)")
			break
		}
	}
	if err := iter.Err(); err != nil {
		log.Fatalf("Error iterating audit logs: %v", err)
	}

	fmt.Println("========================================")
}
