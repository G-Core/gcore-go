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

func TestCloudV2PortAllowAddressPairsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V2.Ports.AllowAddressPairs(
		context.TODO(),
		int64(0),
		int64(0),
		"port_id",
		gcore.CloudV2PortAllowAddressPairsParams{
			AllowedAddressPairs: gcore.F([]gcore.AllowedAddressPairsParam{{
				IPAddress:  gcore.F("192.168.123.20"),
				MacAddress: gcore.F("00:16:3e:f2:87:16"),
			}, {
				IPAddress:  gcore.F("192.168.0.0/17"),
				MacAddress: gcore.F("00:16:3e:f2:87:89"),
			}}),
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
