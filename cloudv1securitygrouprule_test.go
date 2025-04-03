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

func TestCloudV1SecuritygroupruleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Securitygrouprules.Update(
		context.TODO(),
		int64(0),
		int64(0),
		"pk",
		gcore.CloudV1SecuritygroupruleUpdateParams{
			Direction:       gcore.F(gcore.CloudV1SecuritygroupruleUpdateParamsDirectionEgress),
			SecurityGroupID: gcore.F("00000000-0000-4000-8000-000000000000"),
			Description:     gcore.F("Some description"),
			Ethertype:       gcore.F(gcore.CloudV1SecuritygroupruleUpdateParamsEthertypeIPv4),
			PortRangeMax:    gcore.F(int64(80)),
			PortRangeMin:    gcore.F(int64(80)),
			Protocol:        gcore.F(gcore.SecurityGroupProtocolAh),
			RemoteGroupID:   gcore.F("00000000-0000-4000-8000-000000000000"),
			RemoteIPPrefix:  gcore.F("10.0.0.0/8"),
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

func TestCloudV1SecuritygroupruleDelete(t *testing.T) {
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
	err := client.Cloud.V1.Securitygrouprules.Delete(
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
