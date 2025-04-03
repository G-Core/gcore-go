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

func TestCloudV1LbpoolHealthmonitorNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.Lbpools.Healthmonitor.New(
		context.TODO(),
		int64(0),
		int64(0),
		"pool_id",
		gcore.CloudV1LbpoolHealthmonitorNewParams{
			CreateLbHealthMonitor: gcore.CreateLbHealthMonitorParam{
				Delay:          gcore.F(int64(5)),
				MaxRetries:     gcore.F(int64(1)),
				Timeout:        gcore.F(int64(30)),
				Type:           gcore.F(gcore.HealthMonitorTypeHTTP),
				ExpectedCodes:  gcore.F("200,301,302"),
				HTTPMethod:     gcore.F(gcore.HTTPMethodHealthmonitorConnect),
				MaxRetriesDown: gcore.F(int64(3)),
				URLPath:        gcore.F("/"),
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

func TestCloudV1LbpoolHealthmonitorDelete(t *testing.T) {
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
	err := client.Cloud.V1.Lbpools.Healthmonitor.Delete(
		context.TODO(),
		int64(0),
		int64(0),
		"pool_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
