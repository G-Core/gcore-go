// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cdn"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
)

func TestMetricListWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.Metrics.List(context.TODO(), cdn.MetricListParams{
		From:    "2021-06-14T00:00:00Z",
		Metrics: []string{"edge_status_2xx", "edge_status_3xx", "edge_status_4xx", "edge_status_5xx"},
		To:      "2021-06-15T00:00:00Z",
		FilterBy: []cdn.MetricListParamsFilterBy{{
			Field: "resource",
			Op:    "eq",
			Values: []cdn.MetricListParamsFilterByValueUnion{{
				OfFloat: gcore.Float(1234),
			}},
		}},
		Granularity: gcore.String("P1D"),
		GroupBy:     []string{"cname"},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
