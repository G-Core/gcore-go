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

func TestInstanceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.New(context.TODO(), cloud.InstanceNewParams{
		ProjectID:    gcore.Int(1),
		BootVolumeID: "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
		FlavorName:   "g2-standard-32-64",
		Interfaces: []cloud.InstanceNewParamsInterface{{
			PortID: "bbbbbbbb-cccc-dddd-eeee-ffffffffffff",
		}},
		RegionID:         7,
		DryRun:           gcore.Bool(true),
		DataVolumeIDs:    []string{"string"},
		Name:             gcore.String("web-1"),
		Password:         gcore.String("x"),
		PlacementGroupID: gcore.String("placement_group_id"),
		SSHKeyID:         gcore.String("ssh_key_id"),
		Tags: map[string]string{
			"foo": "my-tag-value",
		},
		UserData: gcore.String("x"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInstanceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Update(
		context.TODO(),
		"6c6aa80f-4234-4678-9abc-def012345678",
		cloud.InstanceUpdateParams{
			ProjectID:     gcore.Int(1),
			DryRun:        gcore.Bool(true),
			DataVolumeIDs: []string{"string"},
			FlavorName:    gcore.String("flavor_name"),
			Interfaces: []cloud.InstanceUpdateParamsInterface{{
				PortID: "port_id",
			}},
			Name:             gcore.String("name"),
			PlacementGroupID: gcore.String("placement_group_id"),
			Tags: cloud.TagUpdateMap{
				"foo": gcore.Ptr("my-tag-value"),
			},
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

func TestInstanceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.List(context.TODO(), cloud.InstanceListParams{
		ProjectID: gcore.Int(1),
		CreatedAt: cloud.InstanceListParamsCreatedAt{
			Gt:  gcore.Time(time.Now()),
			Gte: gcore.Time(time.Now()),
			Lt:  gcore.Time(time.Now()),
			Lte: gcore.Time(time.Now()),
		},
		FlavorName: cloud.InstanceListParamsFlavorName{
			Contains: []string{"string"},
			Exact:    []string{"string"},
			Prefix:   []string{"string"},
			Suffix:   []string{"string"},
		},
		IDs: []string{"string"},
		IPAddress: cloud.InstanceListParamsIPAddress{
			Contains: []string{"string"},
			Exact:    []string{"string"},
		},
		Limit:                  gcore.Int(1),
		Managed:                gcore.Bool(true),
		ManagedBy:              cloud.InstanceListParamsManagedByK8SCluster,
		ManagedByResourceIDs:   []string{"string"},
		ManagedByResourceNames: []string{"string"},
		Name: cloud.InstanceListParamsName{
			Contains: []string{"string"},
			Exact:    []string{"string"},
			Prefix:   []string{"string"},
			Suffix:   []string{"string"},
		},
		NetworkIDs:       []string{"string"},
		Offset:           gcore.Int(0),
		OrderBy:          cloud.InstanceListParamsOrderByCreatedAtAsc,
		RegionIDs:        []int64{0},
		SecurityGroupIDs: []string{"string"},
		Statuses:         []string{"active"},
		SubnetIDs:        []string{"string"},
		TagKey: cloud.InstanceListParamsTagKey{
			Contains: []string{"string"},
			Exact:    []string{"string"},
			Prefix:   []string{"string"},
			Suffix:   []string{"string"},
		},
		TagValue: cloud.InstanceListParamsTagValue{
			Contains: []string{"string"},
			Exact:    []string{"string"},
			Prefix:   []string{"string"},
			Suffix:   []string{"string"},
		},
		Tags: map[string][]string{
			"env": {"prod"},
		},
		UpdatedAt: cloud.InstanceListParamsUpdatedAt{
			Gt:  gcore.Time(time.Now()),
			Gte: gcore.Time(time.Now()),
			Lt:  gcore.Time(time.Now()),
			Lte: gcore.Time(time.Now()),
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

func TestInstanceDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Delete(
		context.TODO(),
		"6c6aa80f-4234-4678-9abc-def012345678",
		cloud.InstanceDeleteParams{
			ProjectID:      gcore.Int(1),
			AllFloatingIPs: gcore.Bool(true),
			AllPorts:       gcore.Bool(true),
			AllVolumes:     gcore.Bool(true),
			DryRun:         gcore.Bool(true),
			FloatingIPIDs:  []string{"string"},
			PortIDs:        []string{"string"},
			VolumeIDs:      []string{"string"},
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

func TestInstanceActionWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Action(
		context.TODO(),
		"6c6aa80f-4234-4678-9abc-def012345678",
		cloud.InstanceActionParams{
			ProjectID: gcore.Int(1),
			Action:    cloud.InstanceActionParamsActionStop,
			DryRun:    gcore.Bool(true),
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

func TestInstanceGet(t *testing.T) {
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
	_, err := client.Cloud.Instances.Get(
		context.TODO(),
		"6c6aa80f-4234-4678-9abc-def012345678",
		cloud.InstanceGetParams{
			ProjectID: gcore.Int(1),
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

func TestInstanceRebuildWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Instances.Rebuild(
		context.TODO(),
		"a314d566-2772-4124-a4c4-5688ab76b928",
		cloud.InstanceRebuildParams{
			ProjectID:  gcore.Int(1),
			RegionID:   gcore.Int(1),
			ImageID:    gcore.String("f01fd9a0-9548-48ba-82dc-a8c8b2d6f2f1"),
			SSHKeyName: gcore.String("my-ssh-key"),
			UserData:   gcore.String("user_data"),
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
