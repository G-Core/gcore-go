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

func TestGPUVirtualClusterServerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.GPUVirtual.Clusters.Servers.List(
		context.TODO(),
		"1aaaab48-10d0-46d9-80cc-85209284ceb4",
		cloud.GPUVirtualClusterServerListParams{
			ProjectID:     gcore.Int(1),
			RegionID:      gcore.Int(7),
			ChangedBefore: gcore.Time(time.Now()),
			ChangedSince:  gcore.Time(time.Now()),
			IPAddress:     gcore.String("237.84.2.178"),
			Limit:         gcore.Int(10),
			Name:          gcore.String("name"),
			Offset:        gcore.Int(0),
			OrderBy:       cloud.GPUVirtualClusterServerListParamsOrderByCreatedAtAsc,
			Status:        cloud.GPUVirtualClusterServerListParamsStatusActive,
			Uuids:         []string{"string"},
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

func TestGPUVirtualClusterServerDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.GPUVirtual.Clusters.Servers.Delete(
		context.TODO(),
		"f1c1eeb6-1834-48c9-a7b0-daafce64872b",
		cloud.GPUVirtualClusterServerDeleteParams{
			ProjectID:           gcore.Int(1),
			RegionID:            gcore.Int(7),
			ClusterID:           "1aaaab48-10d0-46d9-80cc-85209284ceb4",
			AllFloatingIPs:      gcore.Bool(true),
			AllReservedFixedIPs: gcore.Bool(true),
			AllVolumes:          gcore.Bool(true),
			FloatingIPIDs:       []string{"e4a01208-d6ac-4304-bf86-3028154b070a"},
			ReservedFixedIPIDs:  []string{"a29b8e1e-08d3-4cec-91fb-06e81e5f46d5"},
			VolumeIDs:           []string{"1333c684-c3da-4b91-ac9e-a92706aa2824"},
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
