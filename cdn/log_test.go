// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cdn_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/G-Core/gcore-go"
	"github.com/G-Core/gcore-go/cdn"
	"github.com/G-Core/gcore-go/internal/testutil"
	"github.com/G-Core/gcore-go/option"
)

func TestLogListWithOptionalParams(t *testing.T) {
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
	_, err := client.CDN.Logs.List(context.TODO(), cdn.LogListParams{
		From:             "from",
		To:               "to",
		CacheStatusEq:    gcore.String("cache_status__eq"),
		CacheStatusIn:    gcore.String("cache_status__in"),
		CacheStatusNe:    gcore.String("cache_status__ne"),
		CacheStatusNotIn: gcore.String("cache_status__not_in"),
		ClientIPEq:       gcore.String("client_ip__eq"),
		ClientIPIn:       gcore.String("client_ip__in"),
		ClientIPNe:       gcore.String("client_ip__ne"),
		ClientIPNotIn:    gcore.String("client_ip__not_in"),
		CnameContains:    gcore.String("cname__contains"),
		CnameEq:          gcore.String("cname__eq"),
		CnameIn:          gcore.String("cname__in"),
		CnameNe:          gcore.String("cname__ne"),
		CnameNotIn:       gcore.String("cname__not_in"),
		DatacenterEq:     gcore.String("datacenter__eq"),
		DatacenterIn:     gcore.String("datacenter__in"),
		DatacenterNe:     gcore.String("datacenter__ne"),
		DatacenterNotIn:  gcore.String("datacenter__not_in"),
		Fields:           gcore.String("fields"),
		Limit:            gcore.Int(1),
		MethodEq:         gcore.String("method__eq"),
		MethodIn:         gcore.String("method__in"),
		MethodNe:         gcore.String("method__ne"),
		MethodNotIn:      gcore.String("method__not_in"),
		Offset:           gcore.Int(0),
		Ordering:         gcore.String("ordering"),
		ResourceIDEq:     gcore.Int(0),
		ResourceIDGt:     gcore.Int(0),
		ResourceIDGte:    gcore.Int(0),
		ResourceIDIn:     gcore.String("resource_id__in"),
		ResourceIDLt:     gcore.Int(0),
		ResourceIDLte:    gcore.Int(0),
		ResourceIDNe:     gcore.Int(0),
		ResourceIDNotIn:  gcore.String("resource_id__not_in"),
		SizeEq:           gcore.Int(0),
		SizeGt:           gcore.Int(0),
		SizeGte:          gcore.Int(0),
		SizeIn:           gcore.String("size__in"),
		SizeLt:           gcore.Int(0),
		SizeLte:          gcore.Int(0),
		SizeNe:           gcore.Int(0),
		SizeNotIn:        gcore.String("size__not_in"),
		StatusEq:         gcore.Int(0),
		StatusGt:         gcore.Int(0),
		StatusGte:        gcore.Int(0),
		StatusIn:         gcore.String("status__in"),
		StatusLt:         gcore.Int(0),
		StatusLte:        gcore.Int(0),
		StatusNe:         gcore.Int(0),
		StatusNotIn:      gcore.String("status__not_in"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogDownloadWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := gcore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.CDN.Logs.Download(context.TODO(), cdn.LogDownloadParams{
		Format:           "format",
		From:             "from",
		To:               "to",
		CacheStatusEq:    gcore.String("cache_status__eq"),
		CacheStatusIn:    gcore.String("cache_status__in"),
		CacheStatusNe:    gcore.String("cache_status__ne"),
		CacheStatusNotIn: gcore.String("cache_status__not_in"),
		ClientIPEq:       gcore.String("client_ip__eq"),
		ClientIPIn:       gcore.String("client_ip__in"),
		ClientIPNe:       gcore.String("client_ip__ne"),
		ClientIPNotIn:    gcore.String("client_ip__not_in"),
		CnameContains:    gcore.String("cname__contains"),
		CnameEq:          gcore.String("cname__eq"),
		CnameIn:          gcore.String("cname__in"),
		CnameNe:          gcore.String("cname__ne"),
		CnameNotIn:       gcore.String("cname__not_in"),
		DatacenterEq:     gcore.String("datacenter__eq"),
		DatacenterIn:     gcore.String("datacenter__in"),
		DatacenterNe:     gcore.String("datacenter__ne"),
		DatacenterNotIn:  gcore.String("datacenter__not_in"),
		Fields:           gcore.String("fields"),
		Limit:            gcore.Int(10000),
		MethodEq:         gcore.String("method__eq"),
		MethodIn:         gcore.String("method__in"),
		MethodNe:         gcore.String("method__ne"),
		MethodNotIn:      gcore.String("method__not_in"),
		Offset:           gcore.Int(0),
		ResourceIDEq:     gcore.Int(0),
		ResourceIDGt:     gcore.Int(0),
		ResourceIDGte:    gcore.Int(0),
		ResourceIDIn:     gcore.String("resource_id__in"),
		ResourceIDLt:     gcore.Int(0),
		ResourceIDLte:    gcore.Int(0),
		ResourceIDNe:     gcore.Int(0),
		ResourceIDNotIn:  gcore.String("resource_id__not_in"),
		SizeEq:           gcore.Int(0),
		SizeGt:           gcore.Int(0),
		SizeGte:          gcore.Int(0),
		SizeIn:           gcore.String("size__in"),
		SizeLt:           gcore.Int(0),
		SizeLte:          gcore.Int(0),
		SizeNe:           gcore.Int(0),
		SizeNotIn:        gcore.String("size__not_in"),
		Sort:             gcore.String("sort"),
		StatusEq:         gcore.Int(0),
		StatusGt:         gcore.Int(0),
		StatusGte:        gcore.Int(0),
		StatusIn:         gcore.String("status__in"),
		StatusLt:         gcore.Int(0),
		StatusLte:        gcore.Int(0),
		StatusNe:         gcore.Int(0),
		StatusNotIn:      gcore.String("status__not_in"),
	})
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *gcore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
