// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cdn"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

func TestOriginGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.OriginGroups.New(context.TODO(), cdn.OriginGroupNewParams{
		OfNoneAuth: &cdn.OriginGroupNewParamsBodyNoneAuth{
			Name: "YourOriginGroup",
			Sources: []cdn.OriginGroupNewParamsBodyNoneAuthSourceUnion{{
				OfHostSource: &cdn.OriginGroupNewParamsBodyNoneAuthSourceHostSource{
					Source:             "yourwebsite.com",
					Backup:             gcore.Bool(false),
					Enabled:            gcore.Bool(true),
					HostHeaderOverride: param.Null[string](),
					Tag:                gcore.String("default"),
				},
			}},
			AuthType:          gcore.String("none"),
			ProxyNextUpstream: []string{"error", "timeout", "invalid_header", "http_500", "http_502", "http_503", "http_504"},
			UseNext:           gcore.Bool(true),
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

func TestOriginGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.OriginGroups.Update(
		context.TODO(),
		0,
		cdn.OriginGroupUpdateParams{
			OfNoneAuth: &cdn.OriginGroupUpdateParamsBodyNoneAuth{
				Name:              "YourOriginGroup",
				AuthType:          gcore.String("none"),
				Path:              gcore.String(""),
				ProxyNextUpstream: []string{"error", "timeout", "invalid_header", "http_500", "http_502", "http_503", "http_504"},
				Sources: []cdn.OriginGroupUpdateParamsBodyNoneAuthSourceUnion{{
					OfHostSource: &cdn.OriginGroupUpdateParamsBodyNoneAuthSourceHostSource{
						Source:             "yourdomain.com",
						Backup:             gcore.Bool(false),
						Enabled:            gcore.Bool(true),
						HostHeaderOverride: param.Null[string](),
						Tag:                gcore.String("default"),
					},
				}},
				UseNext: gcore.Bool(true),
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

func TestOriginGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.OriginGroups.List(context.TODO(), cdn.OriginGroupListParams{
		HasRelatedResources: gcore.Bool(true),
		Name:                gcore.String("name"),
		Sources:             gcore.String("sources"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOriginGroupDelete(t *testing.T) {
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
	err := client.CDN.OriginGroups.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOriginGroupGet(t *testing.T) {
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
	_, err := client.CDN.OriginGroups.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOriginGroupReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.OriginGroups.Replace(
		context.TODO(),
		0,
		cdn.OriginGroupReplaceParams{
			OfNoneAuth: &cdn.OriginGroupReplaceParamsBodyNoneAuth{
				AuthType: "none",
				Name:     "YourOriginGroup",
				Path:     "",
				Sources: []cdn.OriginGroupReplaceParamsBodyNoneAuthSourceUnion{{
					OfHostSource: &cdn.OriginGroupReplaceParamsBodyNoneAuthSourceHostSource{
						Source:             "yourdomain.com",
						Backup:             gcore.Bool(false),
						Enabled:            gcore.Bool(true),
						HostHeaderOverride: param.Null[string](),
						Tag:                gcore.String("default"),
					},
				}},
				UseNext:           true,
				ProxyNextUpstream: []string{"error", "timeout", "invalid_header", "http_500", "http_502", "http_503", "http_504"},
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
