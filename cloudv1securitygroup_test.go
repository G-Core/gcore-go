// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/gcore-go"
	"github.com/stainless-sdks/gcore-go/internal/testutil"
	"github.com/stainless-sdks/gcore-go/option"
)

func TestCloudV1SecuritygroupNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.Securitygroups.New(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1SecuritygroupNewParams{
			SecurityGroup: gcore.F(gcore.CloudV1SecuritygroupNewParamsSecurityGroup{
				Name:        gcore.F("my_security_group"),
				Description: gcore.F("Some description"),
				Metadata: gcore.F(gcore.RawMetadataParam{
					"cost-center": "bar",
					"data-center": "bar",
				}),
				SecurityGroupRules: gcore.F([]gcore.CreateSecurityGroupRuleParam{{
					Description:    gcore.F("Some description"),
					Direction:      gcore.F(gcore.CreateSecurityGroupRuleDirectionEgress),
					Ethertype:      gcore.F(gcore.CreateSecurityGroupRuleEthertypeIPv4),
					PortRangeMax:   gcore.F(int64(80)),
					PortRangeMin:   gcore.F(int64(80)),
					Protocol:       gcore.F(gcore.SecurityGroupProtocolAh),
					RemoteGroupID:  gcore.F("00000000-0000-4000-8000-000000000000"),
					RemoteIPPrefix: gcore.F("10.0.0.0/8"),
				}}),
				Tags: gcore.F([]string{"string"}),
			}),
			Instances: gcore.F([]string{"00000000-0000-4000-8000-000000000000"}),
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

func TestCloudV1SecuritygroupGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.Securitygroups.Get(
		context.TODO(),
		int64(0),
		int64(0),
		"pk",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1SecuritygroupUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.Securitygroups.Update(
		context.TODO(),
		int64(0),
		int64(0),
		"pk",
		gcore.CloudV1SecuritygroupUpdateParams{
			ChangedRules: gcore.F([]gcore.CloudV1SecuritygroupUpdateParamsChangedRule{{
				Action:              gcore.F(gcore.CloudV1SecuritygroupUpdateParamsChangedRulesActionCreate),
				Description:         gcore.F("Some description"),
				Direction:           gcore.F(gcore.CloudV1SecuritygroupUpdateParamsChangedRulesDirectionEgress),
				Ethertype:           gcore.F(gcore.CloudV1SecuritygroupUpdateParamsChangedRulesEthertypeIPv4),
				PortRangeMax:        gcore.F(int64(80)),
				PortRangeMin:        gcore.F(int64(80)),
				Protocol:            gcore.F(gcore.SecurityGroupProtocolAh),
				RemoteGroupID:       gcore.F("00000000-0000-4000-8000-000000000000"),
				RemoteIPPrefix:      gcore.F("10.0.0.0/8"),
				SecurityGroupRuleID: gcore.F("00000000-0000-4000-8000-000000000000"),
			}}),
			Name: gcore.F("some_name"),
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

func TestCloudV1SecuritygroupListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.Securitygroups.List(
		context.TODO(),
		int64(0),
		int64(0),
		gcore.CloudV1SecuritygroupListParams{
			Limit:      gcore.F(int64(0)),
			MetadataK:  gcore.F("metadata_k"),
			MetadataKv: gcore.F("metadata_kv"),
			Offset:     gcore.F(int64(0)),
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

func TestCloudV1SecuritygroupDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	err := client.Cloud.V1.Securitygroups.Delete(
		context.TODO(),
		int64(0),
		int64(0),
		"pk",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1SecuritygroupAddRuleWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.Securitygroups.AddRule(
		context.TODO(),
		int64(0),
		int64(0),
		"pk",
		gcore.CloudV1SecuritygroupAddRuleParams{
			CreateSecurityGroupRule: gcore.CreateSecurityGroupRuleParam{
				Description:    gcore.F("Some description"),
				Direction:      gcore.F(gcore.CreateSecurityGroupRuleDirectionEgress),
				Ethertype:      gcore.F(gcore.CreateSecurityGroupRuleEthertypeIPv4),
				PortRangeMax:   gcore.F(int64(80)),
				PortRangeMin:   gcore.F(int64(80)),
				Protocol:       gcore.F(gcore.SecurityGroupProtocolAh),
				RemoteGroupID:  gcore.F("00000000-0000-4000-8000-000000000000"),
				RemoteIPPrefix: gcore.F("10.0.0.0/8"),
			},
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

func TestCloudV1SecuritygroupCopy(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	err := client.Cloud.V1.Securitygroups.Copy(
		context.TODO(),
		int64(0),
		int64(0),
		"pk",
		gcore.CloudV1SecuritygroupCopyParams{
			NamePydantic: gcore.NamePydanticParam{
				Name: gcore.F("some_name"),
			},
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

func TestCloudV1SecuritygroupFilterInstances(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.Securitygroups.FilterInstances(
		context.TODO(),
		int64(0),
		int64(0),
		"secgroup_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1SecuritygroupRevert(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Cloud.V1.Securitygroups.Revert(
		context.TODO(),
		int64(0),
		int64(0),
		"pk",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
