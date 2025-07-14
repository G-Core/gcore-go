// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fastedge_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/fastedge"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
)

func TestAppNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Fastedge.Apps.New(context.TODO(), fastedge.AppNewParams{
		App: fastedge.AppParam{
			Binary:  gcore.Int(0),
			Comment: gcore.String("comment"),
			Debug:   gcore.Bool(true),
			Env: map[string]string{
				"var1": "value1",
				"var2": "value2",
			},
			Log:  fastedge.AppLogKafka,
			Name: gcore.String("name"),
			RspHeaders: map[string]string{
				"header1": "value1",
				"header2": "value2",
			},
			Secrets: map[string]fastedge.AppSecretParam{
				"foo": {
					ID: 0,
				},
			},
			Status: gcore.Int(0),
			Stores: map[string]int64{
				"country_allow": 1,
				"ip_block":      2,
			},
			Template: gcore.Int(0),
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

func TestAppUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Fastedge.Apps.Update(
		context.TODO(),
		0,
		fastedge.AppUpdateParams{
			App: fastedge.AppParam{
				Binary:  gcore.Int(0),
				Comment: gcore.String("comment"),
				Debug:   gcore.Bool(true),
				Env: map[string]string{
					"var1": "value1",
					"var2": "value2",
				},
				Log:  fastedge.AppLogKafka,
				Name: gcore.String("name"),
				RspHeaders: map[string]string{
					"header1": "value1",
					"header2": "value2",
				},
				Secrets: map[string]fastedge.AppSecretParam{
					"foo": {
						ID: 0,
					},
				},
				Status: gcore.Int(0),
				Stores: map[string]int64{
					"country_allow": 1,
					"ip_block":      2,
				},
				Template: gcore.Int(0),
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

func TestAppListWithOptionalParams(t *testing.T) {
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
	_, err := client.Fastedge.Apps.List(context.TODO(), fastedge.AppListParams{
		APIType:  fastedge.AppListParamsAPITypeWasiHTTP,
		Binary:   gcore.Int(0),
		Limit:    gcore.Int(0),
		Name:     gcore.String("name"),
		Offset:   gcore.Int(0),
		Ordering: fastedge.AppListParamsOrderingName,
		Plan:     gcore.Int(0),
		Status:   gcore.Int(0),
		Template: gcore.Int(0),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppDelete(t *testing.T) {
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
	err := client.Fastedge.Apps.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppGet(t *testing.T) {
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
	_, err := client.Fastedge.Apps.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.Fastedge.Apps.Replace(
		context.TODO(),
		0,
		fastedge.AppReplaceParams{
			Body: fastedge.AppReplaceParamsBody{
				AppParam: fastedge.AppParam{
					Binary:  gcore.Int(0),
					Comment: gcore.String("comment"),
					Debug:   gcore.Bool(true),
					Env: map[string]string{
						"var1": "value1",
						"var2": "value2",
					},
					Log:  fastedge.AppLogKafka,
					Name: gcore.String("name"),
					RspHeaders: map[string]string{
						"header1": "value1",
						"header2": "value2",
					},
					Secrets: map[string]fastedge.AppSecretParam{
						"foo": {
							ID: 0,
						},
					},
					Status: gcore.Int(0),
					Stores: map[string]int64{
						"country_allow": 1,
						"ip_block":      2,
					},
					Template: gcore.Int(0),
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
