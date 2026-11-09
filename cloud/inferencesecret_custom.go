// Custom code. This file is not generated and is preserved across codegen runs.
// It isolates hand-written *AndPoll convenience methods from generated code to
// eliminate merge conflicts.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
)

// DeleteAndPoll deletes an inference secret and polls until the API stops reporting it. The DELETE
// creates no task and answers 204 while the secret is still reported for a few seconds. Polling stops
// on the first 404 and keeps going through other GET errors, since the DELETE already succeeded; the
// last such error is returned only when the polling timeout runs out. Polling honors
// WithPollingIntervalSeconds and WithPollingTimeoutSeconds.
//
// A 404 does not mean the API has finished: a secret created with the same name within a few
// seconds of it can still be removed by the earlier delete.
func (r *InferenceSecretService) DeleteAndPoll(ctx context.Context, secretName string, params InferenceSecretDeleteParams, opts ...option.RequestOption) error {
	err := r.Delete(ctx, secretName, params, opts...)
	if err != nil {
		return err
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return err
	}
	var getParams InferenceSecretGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	getParams.ProjectID = params.ProjectID

	// The GETs must not write into the caller's response target, whether it was
	// passed to this call or set on the client the service was built from.
	getter := *r
	getter.Options = requestconfig.ExcludeResponseBodyInto(r.Options...)
	getOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)

	// Reuse the same polling configuration as TaskService.Poll
	pollingInterval := time.Duration(precfg.PollingIntervalSeconds) * time.Second
	if pollingInterval < time.Second {
		pollingInterval = time.Second
	}
	pollingCtx := ctx
	if precfg.PollingTimeoutSeconds > 0 {
		var cancel context.CancelFunc
		pollingCtx, cancel = context.WithTimeout(ctx, time.Duration(precfg.PollingTimeoutSeconds)*time.Second)
		defer cancel()
	}

	var lastErr error
	for {
		_, err = getter.Get(pollingCtx, secretName, getParams, getOpts...)
		var apiErr *Error
		if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
			return nil
		}
		if err != nil && pollingCtx.Err() == nil {
			lastErr = err
		}

		select {
		case <-pollingCtx.Done():
			if lastErr != nil {
				return fmt.Errorf("inference secret %s is still reported after delete: %w (last error: %w)", secretName, pollingCtx.Err(), lastErr)
			}
			return fmt.Errorf("inference secret %s is still reported after delete: %w", secretName, pollingCtx.Err())
		case <-time.After(pollingInterval):
		}
	}
}
