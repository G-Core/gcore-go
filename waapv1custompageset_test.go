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

func TestWaapV1CustomPageSetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.V1.CustomPageSets.New(context.TODO(), gcore.WaapV1CustomPageSetNewParams{
		Name: gcore.F("x"),
		Block: gcore.F(gcore.BlockPageParam{
			Enabled: gcore.F(true),
			Header:  gcore.F("xxx"),
			Logo:    gcore.F("logo"),
			Text:    gcore.F("xxxxxxxxxxxxxxxxxxxx"),
			Title:   gcore.F("xxx"),
		}),
		BlockCsrf: gcore.F(gcore.BlockCsrfPageParam{
			Enabled: gcore.F(true),
			Header:  gcore.F("xxx"),
			Logo:    gcore.F("logo"),
			Text:    gcore.F("xxxxxxxxxxxxxxxxxxxx"),
			Title:   gcore.F("xxx"),
		}),
		Captcha: gcore.F(gcore.CaptchaPageParam{
			Enabled: gcore.F(true),
			Error:   gcore.F("xxxxxxxxxx"),
			Header:  gcore.F("xxx"),
			Logo:    gcore.F("logo"),
			Text:    gcore.F("xxxxxxxxxxxxxxxxxxxx"),
			Title:   gcore.F("xxx"),
		}),
		CookieDisabled: gcore.F(gcore.CookieDisabledPageParam{
			Enabled: gcore.F(true),
			Header:  gcore.F("xxx"),
			Text:    gcore.F("xxxxxxxxxxxxxxxxxxxx"),
		}),
		Domains: gcore.F([]int64{int64(0)}),
		Handshake: gcore.F(gcore.HandshakePageParam{
			Enabled: gcore.F(true),
			Header:  gcore.F("xxx"),
			Logo:    gcore.F("logo"),
			Title:   gcore.F("xxx"),
		}),
		JavascriptDisabled: gcore.F(gcore.JavascriptDisabledPageParam{
			Enabled: gcore.F(true),
			Header:  gcore.F("xxx"),
			Text:    gcore.F("xxxxxxxxxxxxxxxxxxxx"),
		}),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaapV1CustomPageSetGet(t *testing.T) {
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
	_, err := client.Waap.V1.CustomPageSets.Get(context.TODO(), int64(0))
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaapV1CustomPageSetUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Waap.V1.CustomPageSets.Update(
		context.TODO(),
		int64(0),
		gcore.WaapV1CustomPageSetUpdateParams{
			Block: gcore.F(gcore.BlockPageParam{
				Enabled: gcore.F(true),
				Header:  gcore.F("xxx"),
				Logo:    gcore.F("logo"),
				Text:    gcore.F("xxxxxxxxxxxxxxxxxxxx"),
				Title:   gcore.F("xxx"),
			}),
			BlockCsrf: gcore.F(gcore.BlockCsrfPageParam{
				Enabled: gcore.F(true),
				Header:  gcore.F("xxx"),
				Logo:    gcore.F("logo"),
				Text:    gcore.F("xxxxxxxxxxxxxxxxxxxx"),
				Title:   gcore.F("xxx"),
			}),
			Captcha: gcore.F(gcore.CaptchaPageParam{
				Enabled: gcore.F(true),
				Error:   gcore.F("xxxxxxxxxx"),
				Header:  gcore.F("xxx"),
				Logo:    gcore.F("logo"),
				Text:    gcore.F("xxxxxxxxxxxxxxxxxxxx"),
				Title:   gcore.F("xxx"),
			}),
			CookieDisabled: gcore.F(gcore.CookieDisabledPageParam{
				Enabled: gcore.F(true),
				Header:  gcore.F("xxx"),
				Text:    gcore.F("xxxxxxxxxxxxxxxxxxxx"),
			}),
			Domains: gcore.F([]int64{int64(0)}),
			Handshake: gcore.F(gcore.HandshakePageParam{
				Enabled: gcore.F(true),
				Header:  gcore.F("xxx"),
				Logo:    gcore.F("logo"),
				Title:   gcore.F("xxx"),
			}),
			JavascriptDisabled: gcore.F(gcore.JavascriptDisabledPageParam{
				Enabled: gcore.F(true),
				Header:  gcore.F("xxx"),
				Text:    gcore.F("xxxxxxxxxxxxxxxxxxxx"),
			}),
			Name: gcore.F("x"),
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

func TestWaapV1CustomPageSetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.V1.CustomPageSets.List(context.TODO(), gcore.WaapV1CustomPageSetListParams{
		IDs:      gcore.F([]int64{int64(0)}),
		Limit:    gcore.F(int64(0)),
		Name:     gcore.F("name"),
		Offset:   gcore.F(int64(0)),
		Ordering: gcore.F(gcore.WaapV1CustomPageSetListParamsOrderingName),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaapV1CustomPageSetDelete(t *testing.T) {
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
	err := client.Waap.V1.CustomPageSets.Delete(context.TODO(), int64(0))
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
