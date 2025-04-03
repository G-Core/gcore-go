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

func TestCloudV1CostReportGetResourcesWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.CostReport.GetResources(context.TODO(), gcore.CloudV1CostReportGetResourcesParams{
		TimeFrom:       gcore.F(time.Now()),
		TimeTo:         gcore.F(time.Now()),
		EnableLastDay:  gcore.F(false),
		Limit:          gcore.F(int64(10)),
		Offset:         gcore.F(int64(0)),
		Projects:       gcore.F([]int64{int64(16), int64(17), int64(18), int64(19), int64(20)}),
		Regions:        gcore.F([]int64{int64(1), int64(2), int64(3)}),
		ResponseFormat: gcore.F(gcore.CloudV1CostReportGetResourcesParamsResponseFormatCsvRecords),
		SchemaFilter: gcore.F[gcore.CloudV1CostReportGetResourcesParamsSchemaFilterUnion](gcore.SchemaFilterInstanceParam{
			Field:  gcore.F(gcore.SchemaFilterInstanceEnumFlavor),
			Type:   gcore.F(gcore.SchemaFilterInstanceTypeInstance),
			Values: gcore.F([]string{"g1-standard-1-2"}),
		}),
		Sorting: gcore.F([]gcore.CloudV1CostReportGetResourcesParamsSorting{{
			BillingValue: gcore.F(gcore.SortingDirectionsAsc),
			FirstSeen:    gcore.F(gcore.SortingDirectionsAsc),
			LastName:     gcore.F(gcore.SortingDirectionsAsc),
			LastSeen:     gcore.F(gcore.SortingDirectionsAsc),
			Project:      gcore.F(gcore.SortingDirectionsAsc),
			Region:       gcore.F(gcore.SortingDirectionsAsc),
			Type:         gcore.F(gcore.SortingDirectionsAsc),
		}}),
		Tags: gcore.F(gcore.TagsFilterParam{
			Conditions: gcore.F([]gcore.TagsFilterConditionParam{{
				Key:    gcore.F("os_version"),
				Strict: gcore.F(true),
				Value:  gcore.F("22.04"),
			}}),
			ConditionType: gcore.F(gcore.TagsFilterConditionTypeAnd),
		}),
		Types: gcore.F([]gcore.PrebillingResourceTypes{gcore.PrebillingResourceTypesAICluster, gcore.PrebillingResourceTypesBackup}),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1CostReportGetTotalsWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.CostReport.GetTotals(context.TODO(), gcore.CloudV1CostReportGetTotalsParams{
		TimeFrom:       gcore.F(time.Now()),
		TimeTo:         gcore.F(time.Now()),
		EnableLastDay:  gcore.F(false),
		Projects:       gcore.F([]int64{int64(16), int64(17), int64(18), int64(19), int64(20)}),
		Regions:        gcore.F([]int64{int64(1), int64(2), int64(3)}),
		ResponseFormat: gcore.F(gcore.CloudV1CostReportGetTotalsParamsResponseFormatCsvTotals),
		SchemaFilter: gcore.F[gcore.CloudV1CostReportGetTotalsParamsSchemaFilterUnion](gcore.SchemaFilterInstanceParam{
			Field:  gcore.F(gcore.SchemaFilterInstanceEnumFlavor),
			Type:   gcore.F(gcore.SchemaFilterInstanceTypeInstance),
			Values: gcore.F([]string{"g1-standard-1-2"}),
		}),
		Tags: gcore.F(gcore.TagsFilterParam{
			Conditions: gcore.F([]gcore.TagsFilterConditionParam{{
				Key:    gcore.F("os_version"),
				Strict: gcore.F(true),
				Value:  gcore.F("22.04"),
			}}),
			ConditionType: gcore.F(gcore.TagsFilterConditionTypeAnd),
		}),
		Types: gcore.F([]gcore.PrebillingResourceTypes{gcore.PrebillingResourceTypesAICluster, gcore.PrebillingResourceTypesBackup}),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
