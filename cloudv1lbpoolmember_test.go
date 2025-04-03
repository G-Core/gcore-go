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

func TestCloudV1LbpoolMemberNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Lbpools.Member.New(
		context.TODO(),
		int64(0),
		int64(0),
		"pool_id",
		gcore.CloudV1LbpoolMemberNewParams{
			CreateLbPoolMember: gcore.CreateLbPoolMemberParam{
				Address:        gcore.F("192.168.40.33"),
				ProtocolPort:   gcore.F(int64(80)),
				AdminStateUp:   gcore.F(false),
				InstanceID:     gcore.F("a7e7e8d6-0bf7-4ac9-8170-831b47ee2ba9"),
				MonitorAddress: gcore.F("monitor_address"),
				MonitorPort:    gcore.F(int64(0)),
				SubnetID:       gcore.F("32283b0b-b560-4690-810c-f672cbb2e28d"),
				Weight:         gcore.F(int64(1)),
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

func TestCloudV1LbpoolMemberDelete(t *testing.T) {
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
	_, err := client.Cloud.V1.Lbpools.Member.Delete(
		context.TODO(),
		int64(0),
		int64(0),
		"pool_id",
		"member_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
