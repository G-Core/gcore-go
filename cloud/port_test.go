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

func TestPortNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Ports.New(context.TODO(), cloud.PortNewParams{
		ProjectID: gcore.Int(1),
		RegionID:  1,
		DryRun:    gcore.Bool(true),
		AllowedAddressPairs: []cloud.PortNewParamsAllowedAddressPair{{
			IPAddress: "192.168.123.20",
		}},
		IPAssignments: []cloud.IPAssignmentRequestParam{{
			SubnetID:  "b39792c3-3160-4356-912e-ba396c95cdcf",
			IPAddress: gcore.String("192.168.123.20"),
		}},
		NetworkID:           gcore.String("bc688791-f1b0-44eb-97d4-07697294b1e1"),
		PortSecurityEnabled: gcore.Bool(true),
		SecurityGroupIDs:    []string{"351b0dd7-ca09-431c-be53-935db3785067"},
		Tags: map[string]string{
			"my-tag": "my-tag-value",
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

func TestPortUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Ports.Update(
		context.TODO(),
		"1f0ca628-a73b-42c0-bdac-7b10d023e097",
		cloud.PortUpdateParams{
			ProjectID: gcore.Int(1),
			AllowedAddressPairs: []cloud.PortUpdateParamsAllowedAddressPair{{
				IPAddress: "192.168.123.20",
			}, {
				IPAddress: "192.168.0.0/17",
			}},
			IPAssignments: []cloud.IPAssignmentRequestParam{{
				SubnetID:  "b39792c3-3160-4356-912e-ba396c95cdcf",
				IPAddress: gcore.String("192.168.123.20"),
			}},
			PortSecurityEnabled: gcore.Bool(true),
			SecurityGroupIDs:    []string{"351b0dd7-ca09-431c-be53-935db3785067"},
			Tags: cloud.TagUpdateMap{
				"my-tag":           gcore.Ptr("my-tag-value"),
				"my-tag-to-remove": nil,
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

func TestPortListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.Ports.List(context.TODO(), cloud.PortListParams{
		ProjectID:               gcore.Int(1),
		Attached:                gcore.Bool(true),
		AttachedToResourceIDs:   []string{"string"},
		AttachedToResourceTypes: []string{"baremetal_server"},
		CreatedAt: cloud.PortListParamsCreatedAt{
			Gt:  gcore.Time(time.Now()),
			Gte: gcore.Time(time.Now()),
			Lt:  gcore.Time(time.Now()),
			Lte: gcore.Time(time.Now()),
		},
		IDs: []string{"string"},
		IPAddress: cloud.PortListParamsIPAddress{
			Contains: []string{"string"},
			Exact:    []string{"string"},
		},
		Limit:                  gcore.Int(1),
		Managed:                gcore.Bool(true),
		ManagedBy:              cloud.PortListParamsManagedByFileShare,
		ManagedByResourceIDs:   []string{"string"},
		ManagedByResourceNames: []string{"string"},
		NetworkIDs:             []string{"string"},
		Offset:                 gcore.Int(0),
		OrderBy:                cloud.PortListParamsOrderByCreatedAtAsc,
		Public:                 gcore.Bool(true),
		RegionIDs:              []int64{0},
		SecurityGroupIDs:       []string{"string"},
		SubnetIDs:              []string{"string"},
		TagKey: cloud.PortListParamsTagKey{
			Contains: []string{"string"},
			Exact:    []string{"string"},
			Prefix:   []string{"string"},
			Suffix:   []string{"string"},
		},
		TagValue: cloud.PortListParamsTagValue{
			Contains: []string{"string"},
			Exact:    []string{"string"},
			Prefix:   []string{"string"},
			Suffix:   []string{"string"},
		},
		Tags: map[string][]string{
			"env": {"prod"},
		},
		UpdatedAt: cloud.PortListParamsUpdatedAt{
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

func TestPortDelete(t *testing.T) {
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
	_, err := client.Cloud.Ports.Delete(
		context.TODO(),
		"1f0ca628-a73b-42c0-bdac-7b10d023e097",
		cloud.PortDeleteParams{
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

func TestPortConvert(t *testing.T) {
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
	_, err := client.Cloud.Ports.Convert(
		context.TODO(),
		"1f0ca628-a73b-42c0-bdac-7b10d023e097",
		cloud.PortConvertParams{
			ProjectID:  gcore.Int(1),
			TargetType: cloud.PortConvertParamsTargetTypeFloatingIP,
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

func TestPortGet(t *testing.T) {
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
	_, err := client.Cloud.Ports.Get(
		context.TODO(),
		"1f0ca628-a73b-42c0-bdac-7b10d023e097",
		cloud.PortGetParams{
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
