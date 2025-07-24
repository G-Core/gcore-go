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

func TestVideoSubtitleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Streaming.Videos.Subtitles.New(
		context.TODO(),
		0,
		streaming.VideoSubtitleNewParams{
			Body: streaming.VideoSubtitleNewParamsBody{
				SubtitleBaseParam: streaming.SubtitleBaseParam{
					Language: gcore.String("language"),
					Name:     gcore.String("German (AI-generated)"),
					Vtt:      gcore.String("WEBVTT\n\n1\n00:00:07.154 --> 00:00:12.736\nWir haben 100 Millionen registrierte Benutzer oder aktive Benutzer, die mindestens einmal pro Woche spielen.\n\n2\n00:00:13.236 --> 00:00:20.198\nWir haben vielleicht 80 oder 100.000, die auf einem bestimmten Cluster spielen."),
				},
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

func TestVideoSubtitleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Streaming.Videos.Subtitles.Update(
		context.TODO(),
		0,
		streaming.VideoSubtitleUpdateParams{
			VideoID: 0,
			SubtitleBase: streaming.SubtitleBaseParam{
				Language: gcore.String("ltz"),
				Name:     gcore.String("name"),
				Vtt:      gcore.String("vtt"),
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

func TestVideoSubtitleList(t *testing.T) {
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
	_, err := client.Streaming.Videos.Subtitles.List(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVideoSubtitleDelete(t *testing.T) {
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
	err := client.Streaming.Videos.Subtitles.Delete(
		context.TODO(),
		0,
		streaming.VideoSubtitleDeleteParams{
			VideoID: 0,
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

func TestVideoSubtitleGet(t *testing.T) {
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
	_, err := client.Streaming.Videos.Subtitles.Get(
		context.TODO(),
		0,
		streaming.VideoSubtitleGetParams{
			VideoID: 0,
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
