// Package polling holds helpers shared by generated *AndPoll service methods.
package polling

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
)

// RecoverActionBody reads the raw create/update response left intact by
// option.WithResponseBodyInto(**http.Response) and decodes it into recovered.
// raw.Body is reset to a fresh reader over the same bytes so the caller can
// still read it after this returns. The returned bytes are the original raw
// payload, ready for sjson enrichment with a post-poll status field.
//
// Returns an error if raw.Body has already been drained (e.g. a different
// WithResponseBodyInto destination shape consumed it during decoding), since
// *AndPoll cannot reconstruct the typed value in that case.
func RecoverActionBody[T any](raw *http.Response, recovered *T) ([]byte, error) {
	if raw == nil || raw.Body == nil {
		return nil, errors.New("response body unavailable; pass WithResponseBodyInto with **http.Response to capture the create/update body")
	}
	rawBytes, err := io.ReadAll(raw.Body)
	_ = raw.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("read action response: %w", err)
	}
	if len(rawBytes) == 0 {
		return nil, errors.New("response body already drained; *AndPoll only supports WithResponseBodyInto with **http.Response")
	}
	if err := apijson.UnmarshalRoot(rawBytes, recovered); err != nil {
		return nil, fmt.Errorf("decode action response: %w", err)
	}
	raw.Body = io.NopCloser(bytes.NewReader(rawBytes))
	return rawBytes, nil
}
