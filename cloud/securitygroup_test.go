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

func TestSecurityGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.SecurityGroups.New(context.TODO(), cloud.SecurityGroupNewParams{
		ProjectID:   gcore.Int(1),
		RegionID:    gcore.Int(1),
		Name:        "my_security_group",
		Description: gcore.String("My security group description"),
		Rules: []cloud.SecurityGroupNewParamsRule{{
			Direction:      "ingress",
			Description:    gcore.String("Some description"),
			Ethertype:      "IPv4",
			PortRangeMax:   gcore.Int(80),
			PortRangeMin:   gcore.Int(80),
			Protocol:       "tcp",
			RemoteGroupID:  gcore.String("00000000-0000-4000-8000-000000000000"),
			RemoteIPPrefix: gcore.String("10.0.0.0/8"),
		}},
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

func TestSecurityGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.SecurityGroups.Update(
		context.TODO(),
		"00000000-0000-4000-8000-000000000000",
		cloud.SecurityGroupUpdateParams{
			ProjectID:   gcore.Int(1),
			RegionID:    gcore.Int(1),
			Description: gcore.String("Some description"),
			Name:        gcore.String("some_name"),
			Rules: []cloud.SecurityGroupUpdateParamsRule{{
				Description:    gcore.String("Some description"),
				Direction:      "egress",
				Ethertype:      "IPv4",
				PortRangeMax:   gcore.Int(80),
				PortRangeMin:   gcore.Int(80),
				Protocol:       "tcp",
				RemoteGroupID:  gcore.String("00000000-0000-4000-8000-000000000000"),
				RemoteIPPrefix: gcore.String("10.0.0.0/8"),
			}},
			Tags: cloud.TagUpdateMap{
				"foo": "my-tag-value",
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

func TestSecurityGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.SecurityGroups.List(context.TODO(), cloud.SecurityGroupListParams{
		ProjectID:   gcore.Int(1),
		RegionID:    gcore.Int(1),
		Limit:       gcore.Int(10),
		Name:        gcore.String("my_security_group"),
		Offset:      gcore.Int(0),
		TagKey:      []string{"key1", "key2"},
		TagKeyValue: gcore.String("tag_key_value"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSecurityGroupDelete(t *testing.T) {
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
	err := client.Cloud.SecurityGroups.Delete(
		context.TODO(),
		"024a29e9-b4b7-4c91-9a46-505be123d9f8",
		cloud.SecurityGroupDeleteParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
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

func TestSecurityGroupCopy(t *testing.T) {
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
	_, err := client.Cloud.SecurityGroups.Copy(
		context.TODO(),
		"024a29e9-b4b7-4c91-9a46-505be123d9f8",
		cloud.SecurityGroupCopyParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
			Name:      "some_name",
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

func TestSecurityGroupGet(t *testing.T) {
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
	_, err := client.Cloud.SecurityGroups.Get(
		context.TODO(),
		"024a29e9-b4b7-4c91-9a46-505be123d9f8",
		cloud.SecurityGroupGetParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
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

func TestSecurityGroupRevertToDefault(t *testing.T) {
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
	_, err := client.Cloud.SecurityGroups.RevertToDefault(
		context.TODO(),
		"024a29e9-b4b7-4c91-9a46-505be123d9f8",
		cloud.SecurityGroupRevertToDefaultParams{
			ProjectID: gcore.Int(1),
			RegionID:  gcore.Int(1),
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
