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

func TestCostReportGetAggregatedWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.CostReports.GetAggregated(context.TODO(), cloud.CostReportGetAggregatedParams{
		TimeFrom:       time.Now(),
		TimeTo:         time.Now(),
		EnableLastDay:  gcore.Bool(false),
		Projects:       []int64{16, 17, 18, 19, 20},
		Regions:        []int64{1, 2, 3},
		ResponseFormat: cloud.CostReportGetAggregatedParamsResponseFormatCsvTotals,
		SchemaFilter: cloud.CostReportGetAggregatedParamsSchemaFilterUnion{
			OfInstance: &cloud.CostReportGetAggregatedParamsSchemaFilterInstance{
				Field:  "flavor",
				Values: []string{"g1-standard-1-2"},
			},
		},
		Tags: cloud.CostReportGetAggregatedParamsTags{
			Conditions: []cloud.CostReportGetAggregatedParamsTagsCondition{{
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

func TestCostReportGetAggregatedMonthlyWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.CostReports.GetAggregatedMonthly(context.TODO(), cloud.CostReportGetAggregatedMonthlyParams{
		TimeFrom:       time.Now(),
		TimeTo:         time.Now(),
		Regions:        []int64{1, 2, 3},
		ResponseFormat: cloud.CostReportGetAggregatedMonthlyParamsResponseFormatCsvTotals,
		SchemaFilter: cloud.CostReportGetAggregatedMonthlyParamsSchemaFilterUnion{
			OfInstance: &cloud.CostReportGetAggregatedMonthlyParamsSchemaFilterInstance{
				Field:  "flavor",
				Values: []string{"g1-standard-1-2"},
			},
		},
		Tags: cloud.CostReportGetAggregatedMonthlyParamsTags{
			Conditions: []cloud.CostReportGetAggregatedMonthlyParamsTagsCondition{{
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

func TestCostReportGetDetailedWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.CostReports.GetDetailed(context.TODO(), cloud.CostReportGetDetailedParams{
		TimeFrom:       time.Now(),
		TimeTo:         time.Now(),
		EnableLastDay:  gcore.Bool(false),
		Limit:          gcore.Int(10),
		Offset:         gcore.Int(0),
		Projects:       []int64{16, 17, 18, 19, 20},
		Regions:        []int64{1, 2, 3},
		ResponseFormat: cloud.CostReportGetDetailedParamsResponseFormatCsvRecords,
		SchemaFilter: cloud.CostReportGetDetailedParamsSchemaFilterUnion{
			OfInstance: &cloud.CostReportGetDetailedParamsSchemaFilterInstance{
				Field:  "flavor",
				Values: []string{"g1-standard-1-2"},
			},
		},
		Sorting: []cloud.CostReportGetDetailedParamsSorting{{
			BillingValue: "asc",
			FirstSeen:    "asc",
			LastName:     "asc",
			LastSeen:     "asc",
			Project:      "asc",
			Region:       "asc",
			Type:         "asc",
		}},
		Tags: cloud.CostReportGetDetailedParamsTags{
			Conditions: []cloud.CostReportGetDetailedParamsTagsCondition{{
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
