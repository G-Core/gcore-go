package requestconfig

import (
	"io"
	"net/http"
	"testing"
)

func TestExcludeResponseBodyInto(t *testing.T) {
	t.Run("WithResponseBodyInto is filtered", func(t *testing.T) {
		var customDest interface{}
		opts := []RequestOption{
			RequestOptionFunc(func(r *RequestConfig) error {
				r.ResponseBodyInto = &customDest
				return nil
			}),
		}

		filtered := ExcludeResponseBodyInto(opts...)
		if len(filtered) != 0 {
			t.Errorf("Expected WithResponseBodyInto to be filtered out, got %d options", len(filtered))
		}
	})

	t.Run("WithResponseInto is NOT filtered", func(t *testing.T) {
		var res *http.Response
		opts := []RequestOption{
			RequestOptionFunc(func(r *RequestConfig) error {
				r.ResponseInto = &res
				return nil
			}),
		}

		filtered := ExcludeResponseBodyInto(opts...)
		if len(filtered) != 1 {
			t.Errorf("Expected WithResponseInto to NOT be filtered, got %d options", len(filtered))
		}
	})

	t.Run("Mixed options", func(t *testing.T) {
		var customDest interface{}
		var res *http.Response
		opts := []RequestOption{
			RequestOptionFunc(func(r *RequestConfig) error {
				r.ResponseBodyInto = &customDest
				return nil
			}),
			RequestOptionFunc(func(r *RequestConfig) error {
				r.ResponseInto = &res
				return nil
			}),
			RequestOptionFunc(func(r *RequestConfig) error {
				r.Request.Header.Set("X-Test", "value")
				return nil
			}),
		}

		filtered := ExcludeResponseBodyInto(opts...)
		if len(filtered) != 2 {
			t.Errorf("Expected 2 options after filtering (ResponseInto + Header), got %d", len(filtered))
		}
	})

	t.Run("Other options are kept", func(t *testing.T) {
		var customDest interface{}
		opts := []RequestOption{
			RequestOptionFunc(func(r *RequestConfig) error {
				r.ResponseBodyInto = &customDest
				return nil
			}),
			RequestOptionFunc(func(r *RequestConfig) error {
				r.Request.Header.Set("X-Test", "value")
				return nil
			}),
			RequestOptionFunc(func(r *RequestConfig) error {
				r.Body = nil
				return nil
			}),
		}

		filtered := ExcludeResponseBodyInto(opts...)
		if len(filtered) != 2 {
			t.Errorf("Expected 2 options after filtering (Header + Body), got %d", len(filtered))
		}
	})
}

func TestWriteResponseBodyInto(t *testing.T) {
	raw := []byte(`{"id":"abc","name":"test"}`)

	t.Run("writes into *[]byte", func(t *testing.T) {
		var dst []byte
		opts := []RequestOption{
			RequestOptionFunc(func(r *RequestConfig) error {
				r.ResponseBodyInto = &dst
				return nil
			}),
		}
		err := WriteResponseBodyInto(opts, raw)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if string(dst) != string(raw) {
			t.Errorf("expected %q, got %q", raw, dst)
		}
	})

	t.Run("writes into **http.Response", func(t *testing.T) {
		var res *http.Response
		opts := []RequestOption{
			RequestOptionFunc(func(r *RequestConfig) error {
				r.ResponseBodyInto = &res
				return nil
			}),
		}
		err := WriteResponseBodyInto(opts, raw)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res == nil {
			t.Fatal("expected non-nil response")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status 200, got %d", res.StatusCode)
		}
		body, _ := io.ReadAll(res.Body)
		if string(body) != string(raw) {
			t.Errorf("expected body %q, got %q", raw, body)
		}
	})

	t.Run("decodes into arbitrary struct", func(t *testing.T) {
		var dst struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}
		opts := []RequestOption{
			RequestOptionFunc(func(r *RequestConfig) error {
				r.ResponseBodyInto = &dst
				return nil
			}),
		}
		err := WriteResponseBodyInto(opts, raw)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if dst.ID != "abc" || dst.Name != "test" {
			t.Errorf("expected {abc test}, got {%s %s}", dst.ID, dst.Name)
		}
	})

	t.Run("no-op when no WithResponseBodyInto", func(t *testing.T) {
		opts := []RequestOption{
			RequestOptionFunc(func(r *RequestConfig) error {
				r.Request.Header.Set("X-Test", "value")
				return nil
			}),
		}
		err := WriteResponseBodyInto(opts, raw)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("no-op with empty opts", func(t *testing.T) {
		err := WriteResponseBodyInto(nil, raw)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
