package requestconfig

import (
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
	cfg := newMinimalConfig()
	originalResponseBodyInto := cfg.ResponseBodyInto

	// Apply the option
	_ = optFunc.Apply(cfg)

	// Check if ResponseBodyInto changed
	return !reflect.DeepEqual(originalResponseBodyInto, cfg.ResponseBodyInto)
}

// newMinimalConfig creates a minimally initialized RequestConfig for option detection.
func newMinimalConfig() *RequestConfig {
	req, _ := http.NewRequest("GET", "http://localhost", nil)
	baseURL, _ := url.Parse("http://localhost")
	return &RequestConfig{
		Request: req,
		BaseURL: baseURL,
	}
}
