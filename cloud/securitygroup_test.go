// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/gcore-go"
	"github.com/stainless-sdks/gcore-go/cloud"
	"github.com/stainless-sdks/gcore-go/internal/testutil"
	"github.com/stainless-sdks/gcore-go/option"
)

func TestSecurityGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.SecurityGroups.New(context.TODO(), cloud.SecurityGroupNewParams{
		ProjectID: gcore.Int(0),
		RegionID:  gcore.Int(0),
		SecurityGroup: cloud.SecurityGroupNewParamsSecurityGroup{
			Name:        "my_security_group",
			Description: gcore.String("Some description"),
			SecurityGroupRules: []cloud.SecurityGroupNewParamsSecurityGroupSecurityGroupRule{{
				Description:    gcore.String("Some description"),
				Direction:      "ingress",
				Ethertype:      "IPv4",
				PortRangeMax:   gcore.Int(80),
				PortRangeMin:   gcore.Int(80),
				Protocol:       "tcp",
				RemoteGroupID:  gcore.String("00000000-0000-4000-8000-000000000000"),
				RemoteIPPrefix: gcore.String("10.0.0.0/8"),
			}},
			Tags: map[string]any{
				"my-tag": "bar",
			},
		},
		Instances: []string{"00000000-0000-4000-8000-000000000000"},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSecurityGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.SecurityGroups.Update(
		context.TODO(),
		"group_id",
		cloud.SecurityGroupUpdateParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			ChangedRules: []cloud.SecurityGroupUpdateParamsChangedRule{{
				Action:              "delete",
				Description:         gcore.String("Some description"),
				Direction:           "egress",
				Ethertype:           "IPv4",
				PortRangeMax:        gcore.Int(80),
				PortRangeMin:        gcore.Int(80),
				Protocol:            "tcp",
				RemoteGroupID:       gcore.String("00000000-0000-4000-8000-000000000000"),
				RemoteIPPrefix:      gcore.String("10.0.0.0/8"),
				SecurityGroupRuleID: gcore.String("00000000-0000-4000-8000-000000000000"),
			}},
			Name: gcore.String("some_name"),
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

func TestSecurityGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.SecurityGroups.List(context.TODO(), cloud.SecurityGroupListParams{
		ProjectID:   gcore.Int(0),
		RegionID:    gcore.Int(0),
		Limit:       gcore.Int(0),
		Offset:      gcore.Int(0),
		TagKey:      []string{"string"},
		TagKeyValue: gcore.String("tag_key_value"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSecurityGroupDelete(t *testing.T) {
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
	err := client.Cloud.SecurityGroups.Delete(
		context.TODO(),
		"group_id",
		cloud.SecurityGroupDeleteParams{
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

func TestSecurityGroupCopy(t *testing.T) {
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
	err := client.Cloud.SecurityGroups.Copy(
		context.TODO(),
		"group_id",
		cloud.SecurityGroupCopyParams{
			ProjectID: gcore.Int(0),
			RegionID:  gcore.Int(0),
			Name:      "some_name",
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

func TestSecurityGroupGet(t *testing.T) {
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
	_, err := client.Cloud.SecurityGroups.Get(
		context.TODO(),
		"group_id",
		cloud.SecurityGroupGetParams{
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

func TestSecurityGroupRevertToDefault(t *testing.T) {
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
	_, err := client.Cloud.SecurityGroups.RevertToDefault(
		context.TODO(),
		"group_id",
		cloud.SecurityGroupRevertToDefaultParams{
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
