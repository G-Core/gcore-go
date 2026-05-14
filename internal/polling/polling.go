// Package polling holds helpers shared by generated *AndPoll service methods.
package polling

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
)

// RecoverActionBody recovers the action response payload after the synchronous
// create/update returned, so the *AndPoll polling loop can enrich it.
//
// It tries, in order:
//   - reading raw.Body (works when the caller passed WithResponseBodyInto with
//     a **http.Response target — the request executor leaves the body intact).
//   - the bytes already populated into the caller's *[]byte target by the
//     request executor (the body itself has been drained, but the bytes survive
//     in the user's destination).
//
// On success, raw.Body is reset to a fresh reader over the recovered bytes so
// downstream code can re-read it.
func RecoverActionBody[T any](raw *http.Response, recovered *T, opts ...option.RequestOption) ([]byte, error) {
	if raw == nil {
		return nil, errors.New("response unavailable; pass WithResponseBodyInto with **http.Response or *[]byte")
	}

	var rawBytes []byte

	// Prefer the response body — only intact when the caller used
	// WithResponseBodyInto(**http.Response), which short-circuits body reading
	// inside the request executor.
	if raw.Body != nil {
		b, err := io.ReadAll(raw.Body)
		_ = raw.Body.Close()
		if err == nil && len(b) > 0 {
			rawBytes = b
		}
	}

	// Fall back to the bytes the request executor already wrote into the
	// caller's *[]byte target. This is the *[]byte path: the body itself has
	// been drained, but the decoded bytes are sitting in the user's slice.
	if rawBytes == nil {
		if dst := findResponseBytesTarget(opts); dst != nil && len(*dst) > 0 {
			rawBytes = *dst
		}
	}

	if rawBytes == nil {
		return nil, errors.New("response body unavailable; pass WithResponseBodyInto with **http.Response or *[]byte to *AndPoll")
	}

	if err := apijson.UnmarshalRoot(rawBytes, recovered); err != nil {
		return nil, fmt.Errorf("decode action response: %w", err)
	}

	raw.Body = io.NopCloser(bytes.NewReader(rawBytes))
	return rawBytes, nil
}

// WriteResponseBodyInto propagates the post-poll (status-enriched) body back to
// the *[]byte target the caller configured via option.WithResponseBodyInto, so
// they observe the polled state rather than the original action response.
// No-op when the caller did not request a *[]byte body.
func WriteResponseBodyInto(opts []option.RequestOption, enriched []byte) {
	if dst := findResponseBytesTarget(opts); dst != nil {
		*dst = enriched
	}
}

// findResponseBytesTarget returns the *[]byte the caller configured via
// option.WithResponseBodyInto, or nil if none was set.
//
// WithResponseBodyInto is a regular RequestOption (not pre-request), so it is
// not surfaced by PreRequestOptions. We probe each option in isolation on a
// minimally-initialized config to avoid panics in options like WithHeader /
// WithRequestBody that expect a populated config — same pattern as
// requestconfig.ExcludeResponseBodyInto.
func findResponseBytesTarget(opts []option.RequestOption) *[]byte {
	var dst *[]byte
	for _, opt := range opts {
		fn, ok := opt.(requestconfig.RequestOptionFunc)
		if !ok {
			continue
		}
		cfg := requestconfig.NewMinimalConfig()
		if err := fn.Apply(cfg); err != nil {
			continue
		}
		if cur, ok := cfg.ResponseBodyInto.(*[]byte); ok {
			dst = cur
		}
	}
	return dst
}
