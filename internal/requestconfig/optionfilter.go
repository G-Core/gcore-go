package requestconfig

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

// WithoutRequestBody clears any request body set by previous options.
// Used internally in polling methods to ensure Poll() and Get() calls have no request body.
func WithoutRequestBody() RequestOption {
	return RequestOptionFunc(func(r *RequestConfig) error {
		r.Body = nil
		return nil
	})
}

// ExcludeResponseBodyInto returns a new slice of options with WithResponseBodyInto removed.
// Used for action methods and Poll() to ensure they use default deserialization (TaskIDList, Task).
func ExcludeResponseBodyInto(opts ...RequestOption) []RequestOption {
	if len(opts) == 0 {
		return opts
	}

	filtered := make([]RequestOption, 0, len(opts))
	for _, opt := range opts {
		if !modifiesResponseBodyInto(opt) {
			filtered = append(filtered, opt)
		}
	}
	return filtered
}

// modifiesResponseBodyInto checks if an option modifies the ResponseBodyInto field.
func modifiesResponseBodyInto(opt RequestOption) bool {
	optFunc, ok := opt.(RequestOptionFunc)
	if !ok {
		return false
	}

	// Create a minimally initialized test config
	cfg := NewMinimalConfig()
	originalResponseBodyInto := cfg.ResponseBodyInto

	// Apply the option
	_ = optFunc.Apply(cfg)

	// Check if ResponseBodyInto changed
	return !reflect.DeepEqual(originalResponseBodyInto, cfg.ResponseBodyInto)
}

// NewMinimalConfig creates a minimally initialized RequestConfig safe for
// probing option effects without triggering nil-pointer panics in options that
// touch headers, base URL, or request body.
func NewMinimalConfig() *RequestConfig {
	req, _ := http.NewRequest("GET", "http://localhost", nil)
	baseURL, _ := url.Parse("http://localhost")
	return &RequestConfig{
		Request: req,
		BaseURL: baseURL,
	}
}

// WriteResponseBodyInto writes raw JSON bytes into the WithResponseBodyInto destination
// extracted from opts. This is used by AndPoll methods that fetch the final resource via
// List (which requires WithResponseBodyInto to be excluded) and need to write the single
// result back to the caller's destination afterward.
//
// If no WithResponseBodyInto option is present in opts, this is a no-op.
func WriteResponseBodyInto(opts []RequestOption, raw []byte) error {
	cfg := NewMinimalConfig()
	for _, opt := range opts {
		if modifiesResponseBodyInto(opt) {
			_ = opt.Apply(cfg)
		}
	}
	if cfg.ResponseBodyInto == nil {
		return nil
	}
	switch dst := cfg.ResponseBodyInto.(type) {
	case *[]byte:
		*dst = raw
	case **http.Response:
		*dst = &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(bytes.NewReader(raw)),
		}
	default:
		return json.NewDecoder(bytes.NewReader(raw)).Decode(dst)
	}
	return nil
}
