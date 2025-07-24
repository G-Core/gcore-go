// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package streaming_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/streaming"
)

func TestPlayerNewWithOptionalParams(t *testing.T) {
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
	err := client.Streaming.Players.New(context.TODO(), streaming.PlayerNewParams{
		Player: streaming.PlayerParam{
			Name:                 "name",
			ID:                   gcore.Int(0),
			Autoplay:             gcore.Bool(true),
			BgColor:              gcore.String("bg_color"),
			ClientID:             gcore.Int(0),
			CustomCss:            gcore.String("custom_css"),
			Design:               gcore.String("design"),
			DisableSkin:          gcore.Bool(true),
			FgColor:              gcore.String("fg_color"),
			Framework:            gcore.String("framework"),
			HoverColor:           gcore.String("hover_color"),
			JsURL:                gcore.String("js_url"),
			Logo:                 gcore.String("logo"),
			LogoPosition:         gcore.String("logo_position"),
			Mute:                 gcore.Bool(true),
			SaveOptionsToCookies: gcore.Bool(true),
			ShowSharing:          gcore.Bool(true),
			SkinIsURL:            gcore.String("skin_is_url"),
			SpeedControl:         gcore.Bool(true),
			TextColor:            gcore.String("text_color"),
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

func TestPlayerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Streaming.Players.Update(
		context.TODO(),
		0,
		streaming.PlayerUpdateParams{
			Player: streaming.PlayerParam{
				Name:                 "name",
				ID:                   gcore.Int(0),
				Autoplay:             gcore.Bool(true),
				BgColor:              gcore.String("bg_color"),
				ClientID:             gcore.Int(0),
				CustomCss:            gcore.String("custom_css"),
				Design:               gcore.String("design"),
				DisableSkin:          gcore.Bool(true),
				FgColor:              gcore.String("fg_color"),
				Framework:            gcore.String("framework"),
				HoverColor:           gcore.String("hover_color"),
				JsURL:                gcore.String("js_url"),
				Logo:                 gcore.String("logo"),
				LogoPosition:         gcore.String("logo_position"),
				Mute:                 gcore.Bool(true),
				SaveOptionsToCookies: gcore.Bool(true),
				ShowSharing:          gcore.Bool(true),
				SkinIsURL:            gcore.String("skin_is_url"),
				SpeedControl:         gcore.Bool(true),
				TextColor:            gcore.String("text_color"),
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

func TestPlayerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Streaming.Players.List(context.TODO(), streaming.PlayerListParams{
		Page: gcore.Int(0),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlayerDelete(t *testing.T) {
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
	err := client.Streaming.Players.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlayerGet(t *testing.T) {
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
	_, err := client.Streaming.Players.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlayerPreview(t *testing.T) {
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
	err := client.Streaming.Players.Preview(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
