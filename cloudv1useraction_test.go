// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/gcore-go"
	"github.com/stainless-sdks/gcore-go/internal/testutil"
	"github.com/stainless-sdks/gcore-go/option"
)

func TestCloudV1UserActionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.UserActions.List(context.TODO(), gcore.CloudV1UserActionListParams{
		ActionType:    gcore.F(gcore.CloudV1UserActionListParamsActionTypeActivate),
		APIGroup:      gcore.F(gcore.CloudV1UserActionListParamsAPIGroupAICluster),
		FromTimestamp: gcore.F("from_timestamp"),
		Limit:         gcore.F(int64(0)),
		Offset:        gcore.F(int64(0)),
		ProjectID:     gcore.F(int64(0)),
		ResourceID:    gcore.F("resource_id"),
		SearchField:   gcore.F("search_field"),
		Sorting:       gcore.F(gcore.CloudV1UserActionListParamsSortingAsc),
		ToTimestamp:   gcore.F("to_timestamp"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1UserActionGetSubscriptionsListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cloud.V1.UserActions.GetSubscriptionsList(context.TODO(), gcore.CloudV1UserActionGetSubscriptionsListParams{
		Limit:  gcore.F(int64(0)),
		Offset: gcore.F(int64(0)),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1UserActionSubscribe(t *testing.T) {
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
	err := client.Cloud.V1.UserActions.Subscribe(context.TODO(), gcore.CloudV1UserActionSubscribeParams{
		AuthHeaderName:  gcore.F("Authorization"),
		AuthHeaderValue: gcore.F("Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNTc0OTU0MDY0LCJqdGkiOiJiYXIiLCJ1c2VyX2lkIjoxLCJ1c2VyX3R5cGUiOiJjb21tb24iLCJ1c2VyX2dyb3VwcyI6WyJVc2VycyJdLCJjbGllbnRfaWQiOjEsImVtYWlsIjoidGVzdEB0ZXN0LnRlc3QiLCJ1c2VybmFtZSI6InRlc3RAdGVzdC50ZXN0IiwiaXNfYWRtaW4iOmZhbHNlLCJjbGllbnRfbmFtZSI6Ik5hbWUiLCJleGNoYW5nZWFibGUiOnRydWUsImZha2VfdXNlcl9pZCI6MSwicmp0aSI6ImZvbyJ9.YTH_s67j7xyWlFLy093RxReT5PmitnawLr25Jh7Ix14"),
		URL:             gcore.F("https://your-url.com/receive-user-action-messages"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1UserActionSubscribeAmqpWithOptionalParams(t *testing.T) {
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
	err := client.Cloud.V1.UserActions.SubscribeAmqp(context.TODO(), gcore.CloudV1UserActionSubscribeAmqpParams{
		ConnectionString:         gcore.F("amqps://guest:guest@192.168.123.20:5671/user_action_events"),
		Exchange:                 gcore.F("exchange"),
		ReceiveChildClientEvents: gcore.F(false),
		RoutingKey:               gcore.F("foo"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1UserActionUnsubscribe(t *testing.T) {
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
	err := client.Cloud.V1.UserActions.Unsubscribe(context.TODO())
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCloudV1UserActionUnsubscribeAmqp(t *testing.T) {
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
	err := client.Cloud.V1.UserActions.UnsubscribeAmqp(context.TODO())
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
