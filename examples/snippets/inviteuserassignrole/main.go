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
	"github.com/G-Core/gcore-go/iam"
	"github.com/G-Core/gcore-go/packages/param"
)

func main() {
	// TODO set API key before running
	// apiKey := os.Getenv("GCORE_API_KEY")
	// TODO set client ID before running
	clientIDStr := os.Getenv("GCORE_CLIENT_ID")
	if clientIDStr == "" {
		log.Fatalf("GCORE_CLIENT_ID environment variable is required")
	}
	clientID, err := strconv.ParseInt(clientIDStr, 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLIENT_ID must be a valid integer: %v", err)
	}

	// TODO set cloud project ID before running
	cloudProjectIDStr := os.Getenv("GCORE_CLOUD_PROJECT_ID")
	if cloudProjectIDStr == "" {
		log.Fatalf("GCORE_CLOUD_PROJECT_ID environment variable is required")
	}
	cloudProjectID, err := strconv.ParseInt(cloudProjectIDStr, 10, 64)
	if err != nil {
		log.Fatalf("GCORE_CLOUD_PROJECT_ID must be a valid integer: %v", err)
	}

	client := gcore.NewClient(
	// No need to explicitly pass API key if using environment variables
	// option.WithAPIKey(apiKey),
	)

	userEmail := "john.doe@example.com"
	userName := "John Doe"

	userRole := iam.UserInviteParamsUserRole{
		ID:   param.Opt[int64]{Value: 5},
		Name: "Engineers",
	}

	// Step 1: Invite user via IAM
	invitedUser, err := client.Iam.Users.Invite(context.TODO(), iam.UserInviteParams{
		ClientID: clientID,
		Email:    userEmail,
		Name:     param.Opt[string]{Value: userName},
		UserRole: userRole,
		Lang:     iam.UserInviteParamsLangEn,
	})
	if err != nil {
		log.Fatalf("Failed to invite user: %v", err)
	}

	// Step 2: Create role in cloud for the invited user
	if invitedUser.UserID != 0 {
		_, err := client.Cloud.Users.RoleAssignments.New(context.TODO(), cloud.UserRoleAssignmentNewParams{
			UserID:    invitedUser.UserID,
			Role:      "ProjectAdministrator",
			ProjectID: param.Opt[int64]{Value: cloudProjectID},
			ClientID:  param.Opt[int64]{Value: clientID},
		})
		if err != nil {
			log.Fatalf("Failed to create role assignment: %v", err)
		}

		// Step 3: Create API token for the invited user (expires in 1 month)
		expDate := time.Now().Add(30 * 24 * time.Hour).Format(time.RFC3339)
		_, err = client.Iam.APITokens.New(context.TODO(), clientID, iam.APITokenNewParams{
			ClientUser: iam.APITokenNewParamsClientUser{
				Role: iam.APITokenNewParamsClientUserRole{
					ID:   param.Opt[int64]{Value: 5},
					Name: "Engineers",
				},
			},
			ExpDate:     gcore.String(expDate),
			Name:        fmt.Sprintf("Token for %s", userEmail),
			Description: param.Opt[string]{Value: fmt.Sprintf("API token for invited user %s", userName)},
		})
		if err != nil {
			log.Fatalf("Failed to create API token: %v", err)
		}
	}
}
