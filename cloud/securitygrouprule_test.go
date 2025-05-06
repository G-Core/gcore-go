// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
)

func TestSecurityGroupRuleNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cloud.SecurityGroups.Rules.New(
		context.TODO(),
		"group_id",
		cloud.SecurityGroupRuleNewParams{
			ProjectID:      gcore.Int(0),
			RegionID:       gcore.Int(0),
			Description:    gcore.String("Some description"),
			Direction:      cloud.SecurityGroupRuleNewParamsDirectionIngress,
			Ethertype:      cloud.SecurityGroupRuleNewParamsEthertypeIPv4,
			PortRangeMax:   gcore.Int(80),
			PortRangeMin:   gcore.Int(80),
			Protocol:       cloud.SecurityGroupRuleNewParamsProtocolTcp,
			RemoteGroupID:  gcore.String("00000000-0000-4000-8000-000000000000"),
			RemoteIPPrefix: gcore.String("10.0.0.0/8"),
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSecurityGroupRuleDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Cloud.SecurityGroups.Rules.Delete(
		context.TODO(),
		"rule_id",
		cloud.SecurityGroupRuleDeleteParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSecurityGroupRuleReplaceWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cloud.SecurityGroups.Rules.Replace(
		context.TODO(),
		"rule_id",
		cloud.SecurityGroupRuleReplaceParams{
			ProjectID:       gcore.Int(0),
			RegionID:        gcore.Int(0),
			Direction:       cloud.SecurityGroupRuleReplaceParamsDirectionIngress,
			SecurityGroupID: "00000000-0000-4000-8000-000000000000",
			Description:     gcore.String("Some description"),
			Ethertype:       cloud.SecurityGroupRuleReplaceParamsEthertypeIPv4,
			PortRangeMax:    gcore.Int(80),
			PortRangeMin:    gcore.Int(80),
			Protocol:        cloud.SecurityGroupRuleReplaceParamsProtocolTcp,
			RemoteGroupID:   gcore.String("00000000-0000-4000-8000-000000000000"),
			RemoteIPPrefix:  gcore.String("10.0.0.0/8"),
		},
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
