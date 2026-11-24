// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cloud"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
)

func TestInstanceFlavorListWithOptionalParams(t *testing.T) {
	t.Skip("GCLOUD2-27612: the mock rejects arrays inside object query filters")
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
	_, err := client.Cloud.Instances.Flavors.List(context.TODO(), cloud.InstanceFlavorListParams{
		ProjectID:                  gcore.Int(1),
		Architectures:              []string{"aarch64"},
		Available:                  gcore.Bool(true),
		CompatibleWithBootVolumeID: gcore.String("compatible_with_boot_volume_id"),
		CompatibleWithImageID:      gcore.String("compatible_with_image_id"),
		CompatibleWithSnapshotID:   gcore.String("compatible_with_snapshot_id"),
		Limit:                      gcore.Int(1),
		Name: cloud.InstanceFlavorListParamsName{
			Contains: []string{"string"},
			Exact:    []string{"string"},
			Prefix:   []string{"string"},
			Suffix:   []string{"string"},
		},
		Offset:  gcore.Int(0),
		OrderBy: cloud.InstanceFlavorListParamsOrderByNameAsc,
		OsTypes: []string{"linux"},
		RamMib: cloud.InstanceFlavorListParamsRamMib{
			Gt:  gcore.Int(0),
			Gte: gcore.Int(0),
			Lt:  gcore.Int(0),
			Lte: gcore.Int(0),
		},
		RegionIDs: []int64{0},
		Vcpus: cloud.InstanceFlavorListParamsVcpus{
			Gt:  gcore.Int(0),
			Gte: gcore.Int(0),
			Lt:  gcore.Int(0),
			Lte: gcore.Int(0),
		},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
