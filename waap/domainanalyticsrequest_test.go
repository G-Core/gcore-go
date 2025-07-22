// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/waap"
)

func TestDomainAnalyticsRequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.Domains.Analytics.Requests.List(
		context.TODO(),
		1,
		waap.DomainAnalyticsRequestListParams{
			Start:            time.Now(),
			Actions:          []string{"allow"},
			Countries:        []string{"Mv"},
			End:              gcore.Time(time.Now()),
			IP:               gcore.String(".:"),
			Limit:            gcore.Int(0),
			Offset:           gcore.Int(0),
			Ordering:         gcore.String("ordering"),
			ReferenceID:      gcore.String("2c02efDd09B3BA1AEaDd3dCAa7aC7A37"),
			SecurityRuleName: gcore.String("security_rule_name"),
			StatusCode:       gcore.Int(100),
			TrafficTypes:     []waap.WaapTrafficType{waap.WaapTrafficTypePolicyAllowed},
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

func TestDomainAnalyticsRequestGet(t *testing.T) {
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
	_, err := client.Waap.Domains.Analytics.Requests.Get(
		context.TODO(),
		"request_id",
		waap.DomainAnalyticsRequestGetParams{
			DomainID: 1,
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
