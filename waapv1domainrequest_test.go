// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/gcore-go"
	"github.com/stainless-sdks/gcore-go/internal/testutil"
	"github.com/stainless-sdks/gcore-go/option"
)

func TestWaapV1DomainRequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.V1.Domains.Requests.List(
		context.TODO(),
		int64(0),
		gcore.WaapV1DomainRequestListParams{
			Start:            gcore.F(time.Now()),
			Actions:          gcore.F([]gcore.WaapV1DomainRequestListParamsAction{gcore.WaapV1DomainRequestListParamsActionAllow}),
			Countries:        gcore.F([]string{"Mv"}),
			End:              gcore.F(time.Now()),
			IP:               gcore.F(".:"),
			Limit:            gcore.F(int64(0)),
			Offset:           gcore.F(int64(0)),
			Ordering:         gcore.F("ordering"),
			ReferenceID:      gcore.F("2c02efDd09B3BA1AEaDd3dCAa7aC7A37"),
			SecurityRuleName: gcore.F("security_rule_name"),
			StatusCode:       gcore.F(int64(100)),
			TrafficTypes:     gcore.F([]gcore.WaapV1DomainRequestListParamsTrafficType{gcore.WaapV1DomainRequestListParamsTrafficTypePolicyAllowed}),
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

func TestWaapV1DomainRequestGetDetails(t *testing.T) {
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
	_, err := client.Waap.V1.Domains.Requests.GetDetails(
		context.TODO(),
		int64(0),
		"request_id",
	)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
