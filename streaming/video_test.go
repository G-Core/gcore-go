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

func TestVideoNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Streaming.Videos.New(context.TODO(), streaming.VideoNewParams{
		Video: streaming.CreateVideoParam{
			Name:                           "IBC 2025 - International Broadcasting Convention",
			AutoTranscribeAudioLanguage:    streaming.CreateVideoAutoTranscribeAudioLanguageAuto,
			AutoTranslateSubtitlesLanguage: streaming.CreateVideoAutoTranslateSubtitlesLanguageDisable,
			ClientUserID:                   gcore.Int(10),
			ClipDurationSeconds:            gcore.Int(60),
			ClipStartSeconds:               gcore.Int(137),
			CustomIframeURL:                gcore.String("custom_iframe_url"),
			Description:                    gcore.String("We look forward to welcoming you at IBC2025, which will take place 12-15 September 2025."),
			DirectoryID:                    gcore.Int(800),
			OriginHTTPHeaders:              gcore.String("Authorization: Bearer ..."),
			OriginURL:                      gcore.String("https://www.googleapis.com/drive/v3/files/...?alt=media"),
			Poster:                         gcore.String("data:image/jpeg;base64,/9j/4AA...qf/2Q=="),
			Priority:                       gcore.Int(0),
			Projection:                     gcore.String("regular"),
			QualitySetID:                   gcore.Int(0),
			RemotePosterURL:                gcore.String("remote_poster_url"),
			RemovePoster:                   gcore.Bool(true),
			ScreenshotID:                   gcore.Int(-1),
			ShareURL:                       gcore.String("share_url"),
			SourceBitrateLimit:             gcore.Bool(true),
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

func TestVideoUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Streaming.Videos.Update(
		context.TODO(),
		0,
		streaming.VideoUpdateParams{
			CreateVideo: streaming.CreateVideoParam{
				Name:                           "IBC 2025 - International Broadcasting Convention",
				AutoTranscribeAudioLanguage:    streaming.CreateVideoAutoTranscribeAudioLanguageAuto,
				AutoTranslateSubtitlesLanguage: streaming.CreateVideoAutoTranslateSubtitlesLanguageDisable,
				ClientUserID:                   gcore.Int(10),
				ClipDurationSeconds:            gcore.Int(60),
				ClipStartSeconds:               gcore.Int(137),
				CustomIframeURL:                gcore.String("custom_iframe_url"),
				Description:                    gcore.String("We look forward to welcoming you at IBC2025, which will take place 12-15 September 2025."),
				DirectoryID:                    gcore.Int(800),
				OriginHTTPHeaders:              gcore.String("Authorization: Bearer ..."),
				OriginURL:                      gcore.String("https://www.googleapis.com/drive/v3/files/...?alt=media"),
				Poster:                         gcore.String("data:image/jpeg;base64,/9j/4AA...qf/2Q=="),
				Priority:                       gcore.Int(0),
				Projection:                     gcore.String("regular"),
				QualitySetID:                   gcore.Int(0),
				RemotePosterURL:                gcore.String("remote_poster_url"),
				RemovePoster:                   gcore.Bool(true),
				ScreenshotID:                   gcore.Int(-1),
				ShareURL:                       gcore.String("share_url"),
				SourceBitrateLimit:             gcore.Bool(true),
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

func TestVideoListWithOptionalParams(t *testing.T) {
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
	_, err := client.Streaming.Videos.List(context.TODO(), streaming.VideoListParams{
		ID:           gcore.String("id"),
		ClientUserID: gcore.Int(0),
		Fields:       gcore.String("fields"),
		Page:         gcore.Int(0),
		PerPage:      gcore.Int(0),
		Search:       gcore.String("search"),
		Status:       gcore.String("status"),
		StreamID:     gcore.Int(0),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVideoDelete(t *testing.T) {
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
	err := client.Streaming.Videos.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVideoNewMultipleWithOptionalParams(t *testing.T) {
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
	_, err := client.Streaming.Videos.NewMultiple(context.TODO(), streaming.VideoNewMultipleParams{
		Fields: gcore.String("fields"),
		Videos: []streaming.VideoNewMultipleParamsVideo{{
			CreateVideoParam: streaming.CreateVideoParam{
				Name:                           "IBC 2025 - International Broadcasting Convention",
				AutoTranscribeAudioLanguage:    streaming.CreateVideoAutoTranscribeAudioLanguageAuto,
				AutoTranslateSubtitlesLanguage: streaming.CreateVideoAutoTranslateSubtitlesLanguageDisable,
				ClientUserID:                   gcore.Int(10),
				ClipDurationSeconds:            gcore.Int(60),
				ClipStartSeconds:               gcore.Int(137),
				CustomIframeURL:                gcore.String("custom_iframe_url"),
				Description:                    gcore.String("We look forward to welcoming you at IBC2025, which will take place 12-15 September 2025."),
				DirectoryID:                    gcore.Int(800),
				OriginHTTPHeaders:              gcore.String("Authorization: Bearer ..."),
				OriginURL:                      gcore.String("https://www.googleapis.com/drive/v3/files/...?alt=media"),
				Poster:                         gcore.String("data:image/jpeg;base64,/9j/4AA...qf/2Q=="),
				Priority:                       gcore.Int(0),
				Projection:                     gcore.String("regular"),
				QualitySetID:                   gcore.Int(0),
				RemotePosterURL:                gcore.String("remote_poster_url"),
				RemovePoster:                   gcore.Bool(true),
				ScreenshotID:                   gcore.Int(-1),
				ShareURL:                       gcore.String("share_url"),
				SourceBitrateLimit:             gcore.Bool(true),
			},
			Subtitles: []streaming.SubtitleBaseParam{{
				Language: gcore.String("eng"),
				Vtt:      gcore.String("WEBVTT\n\n1\n00:00:07.154 --> 00:00:12.736\nWe have 100 million registered users or active users who play at least once a week.\n\n2\n00:00:13.236 --> 00:00:20.198\nWe might have 80 or 100,000 playing on a given cluster."),
				Name:     gcore.String("English (AI-generated)"),
			}, {
				Language: gcore.String("ger"),
				Vtt:      gcore.String("WEBVTT\n\n1\n00:00:07.154 --> 00:00:12.736\nWir haben 100 Millionen registrierte Benutzer oder aktive Benutzer, die mindestens einmal pro Woche spielen.\n\n2\n00:00:13.236 --> 00:00:20.198\nWir haben vielleicht 80 oder 100.000, die auf einem bestimmten Cluster spielen."),
				Name:     gcore.String("German (AI-translated)"),
			}},
		}},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVideoGet(t *testing.T) {
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
	_, err := client.Streaming.Videos.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVideoGetParametersForDirectUpload(t *testing.T) {
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
	_, err := client.Streaming.Videos.GetParametersForDirectUpload(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVideoListNamesWithOptionalParams(t *testing.T) {
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
	err := client.Streaming.Videos.ListNames(context.TODO(), streaming.VideoListNamesParams{
		IDs: []int64{0},
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
