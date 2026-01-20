package requestconfig

import (
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
