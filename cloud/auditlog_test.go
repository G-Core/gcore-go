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

func TestAuditLogListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.AuditLogs.List(context.TODO(), cloud.AuditLogListParams{
		ActionType:    []string{"activate", "delete"},
		APIGroup:      []string{"ai_cluster", "image"},
		FromTimestamp: gcore.Time(time.Now()),
		Limit:         gcore.Int(1000),
		Offset:        gcore.Int(0),
		OrderBy:       cloud.AuditLogListParamsOrderByAsc,
		ProjectID:     []int64{1, 2, 3},
		RegionID:      []int64{1, 2, 3},
		ResourceID:    []string{"string"},
		SearchField:   gcore.String("default"),
		Sorting:       cloud.AuditLogListParamsSortingAsc,
		ToTimestamp:   gcore.Time(time.Now()),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
