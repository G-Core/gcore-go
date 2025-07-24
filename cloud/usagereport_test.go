// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
)

func TestUsageReportGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.UsageReports.Get(context.TODO(), cloud.UsageReportGetParams{
		TimeFrom:      time.Now(),
		TimeTo:        time.Now(),
		EnableLastDay: gcore.Bool(false),
		Limit:         gcore.Int(10),
		Offset:        gcore.Int(0),
		Projects:      []int64{16, 17, 18, 19, 20},
		Regions:       []int64{1, 2, 3},
		SchemaFilter: cloud.UsageReportGetParamsSchemaFilterUnion{
			OfInstance: &cloud.UsageReportGetParamsSchemaFilterInstance{
				Field:  "flavor",
				Values: []string{"g1-standard-1-2"},
			},
		},
		Sorting: []cloud.UsageReportGetParamsSorting{{
			BillingValue: "asc",
			FirstSeen:    "asc",
			LastName:     "asc",
			LastSeen:     "asc",
			Project:      "asc",
			Region:       "asc",
			Type:         "asc",
		}},
		Tags: cloud.UsageReportGetParamsTags{
			Conditions: []cloud.UsageReportGetParamsTagsCondition{{
				Key:    gcore.String("os_version"),
				Strict: gcore.Bool(true),
				Value:  gcore.String("22.04"),
			}, {
				Key:    gcore.String("os_version"),
				Strict: gcore.Bool(true),
				Value:  gcore.String("23.04"),
			}},
			ConditionType: "OR",
		},
		Types: []string{"egress_traffic", "instance"},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
