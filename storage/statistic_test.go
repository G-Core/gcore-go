// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/storage"
)

func TestStatisticGetUsageAggregatedWithOptionalParams(t *testing.T) {
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
	_, err := client.Storage.Statistics.GetUsageAggregated(context.TODO(), storage.StatisticGetUsageAggregatedParams{
		From:      gcore.String("2006-01-02"),
		Locations: []string{"s-ed1", "s-drc2", "s-sgc1"},
		Storages:  []string{"123-myStorage"},
		To:        gcore.String("2006-01-02"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStatisticGetUsageSeriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Storage.Statistics.GetUsageSeries(context.TODO(), storage.StatisticGetUsageSeriesParams{
		From:        gcore.String("2006-01-02"),
		Granularity: gcore.String("12h"),
		Locations:   []string{"s-ed1", "s-drc2", "s-sgc1"},
		Source:      gcore.Int(0),
		Storages:    []string{"123-myStorage"},
		To:          gcore.String("2006-01-02"),
		TsString:    gcore.Bool(true),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
