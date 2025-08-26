// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/waap"
)

func TestCustomPageSetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.CustomPageSets.New(context.TODO(), waap.CustomPageSetNewParams{
		Name: "x",
		Block: waap.CustomPageSetNewParamsBlock{
			Enabled: true,
			Header:  gcore.String("xxx"),
			Logo:    gcore.String("logo"),
			Text:    gcore.String("xxxxxxxxxxxxxxxxxxxx"),
			Title:   gcore.String("xxx"),
		},
		BlockCsrf: waap.CustomPageSetNewParamsBlockCsrf{
			Enabled: true,
			Header:  gcore.String("xxx"),
			Logo:    gcore.String("logo"),
			Text:    gcore.String("xxxxxxxxxxxxxxxxxxxx"),
			Title:   gcore.String("xxx"),
		},
		Captcha: waap.CustomPageSetNewParamsCaptcha{
			Enabled: true,
			Error:   gcore.String("xxxxxxxxxx"),
			Header:  gcore.String("xxx"),
			Logo:    gcore.String("logo"),
			Text:    gcore.String("xxxxxxxxxxxxxxxxxxxx"),
			Title:   gcore.String("xxx"),
		},
		CookieDisabled: waap.CustomPageSetNewParamsCookieDisabled{
			Enabled: true,
			Header:  gcore.String("xxx"),
			Text:    gcore.String("xxxxxxxxxxxxxxxxxxxx"),
		},
		Domains: []int64{1},
		Handshake: waap.CustomPageSetNewParamsHandshake{
			Enabled: true,
			Header:  gcore.String("xxx"),
			Logo:    gcore.String("logo"),
			Title:   gcore.String("xxx"),
		},
		JavascriptDisabled: waap.CustomPageSetNewParamsJavascriptDisabled{
			Enabled: true,
			Header:  gcore.String("xxx"),
			Text:    gcore.String("xxxxxxxxxxxxxxxxxxxx"),
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

func TestCustomPageSetUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Waap.CustomPageSets.Update(
		context.TODO(),
		0,
		waap.CustomPageSetUpdateParams{
			Block: waap.CustomPageSetUpdateParamsBlock{
				Enabled: true,
				Header:  gcore.String("xxx"),
				Logo:    gcore.String("logo"),
				Text:    gcore.String("xxxxxxxxxxxxxxxxxxxx"),
				Title:   gcore.String("xxx"),
			},
			BlockCsrf: waap.CustomPageSetUpdateParamsBlockCsrf{
				Enabled: true,
				Header:  gcore.String("xxx"),
				Logo:    gcore.String("logo"),
				Text:    gcore.String("xxxxxxxxxxxxxxxxxxxx"),
				Title:   gcore.String("xxx"),
			},
			Captcha: waap.CustomPageSetUpdateParamsCaptcha{
				Enabled: true,
				Error:   gcore.String("xxxxxxxxxx"),
				Header:  gcore.String("xxx"),
				Logo:    gcore.String("logo"),
				Text:    gcore.String("xxxxxxxxxxxxxxxxxxxx"),
				Title:   gcore.String("xxx"),
			},
			CookieDisabled: waap.CustomPageSetUpdateParamsCookieDisabled{
				Enabled: true,
				Header:  gcore.String("xxx"),
				Text:    gcore.String("xxxxxxxxxxxxxxxxxxxx"),
			},
			Domains: []int64{1},
			Handshake: waap.CustomPageSetUpdateParamsHandshake{
				Enabled: true,
				Header:  gcore.String("xxx"),
				Logo:    gcore.String("logo"),
				Title:   gcore.String("xxx"),
			},
			JavascriptDisabled: waap.CustomPageSetUpdateParamsJavascriptDisabled{
				Enabled: true,
				Header:  gcore.String("xxx"),
				Text:    gcore.String("xxxxxxxxxxxxxxxxxxxx"),
			},
			Name: gcore.String("x"),
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

func TestCustomPageSetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.CustomPageSets.List(context.TODO(), waap.CustomPageSetListParams{
		IDs:      []int64{0},
		Limit:    gcore.Int(0),
		Name:     gcore.String("*example"),
		Offset:   gcore.Int(0),
		Ordering: waap.CustomPageSetListParamsOrderingName,
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomPageSetDelete(t *testing.T) {
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
	err := client.Waap.CustomPageSets.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomPageSetGet(t *testing.T) {
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
	_, err := client.Waap.CustomPageSets.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomPageSetPreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Waap.CustomPageSets.Preview(context.TODO(), waap.CustomPageSetPreviewParams{
		PageType: waap.CustomPageSetPreviewParamsPageTypeBlockHTML,
		Error:    gcore.String("xxxxxxxxxx"),
		Header:   gcore.String("xxx"),
		Logo:     gcore.String("logo"),
		Text:     gcore.String("xxxxxxxxxxxxxxxxxxxx"),
		Title:    gcore.String("xxx"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
