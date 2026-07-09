// Custom code. This file is not generated and is preserved across codegen runs.
// It isolates hand-written *AndPoll convenience methods from generated code to
// eliminate merge conflicts.

package cloud

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/requestconfig"
)

// Poll for task status until it is finished, an error occurs, or the context is done. It uses a default polling interval
// of 1 second which can be overridden to values greater than 0 (otherwise the default value is used).
func (r *TaskService) Poll(ctx context.Context, taskID string, opts ...requestconfig.RequestOption) (*Task, error) {
	// extract polling interval from options, if not explicitly set, the default value is used
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	pollingInterval := time.Duration(precfg.PollingIntervalSeconds) * time.Second
	// ensure the polling interval is at least 1 second
	if pollingInterval < time.Second {
		pollingInterval = time.Second
	}

	// set up polling timeout if configured, otherwise use the provided context
	pollingCtx := ctx
	var cancel context.CancelFunc
	if precfg.PollingTimeoutSeconds > 0 {
		pollingTimeout := time.Duration(precfg.PollingTimeoutSeconds) * time.Second
		pollingCtx, cancel = context.WithTimeout(ctx, pollingTimeout)
		defer cancel()
	}

	// poll the task status until it is finished or an error occurs
	for {
		task, err := r.Get(pollingCtx, taskID)
		if err != nil {
			return nil, fmt.Errorf("failed to get task status: %w", err)
		}

		if task.State == TaskStateFinished {
			return task, nil
		}

		if task.State == TaskStateError {
			return nil, fmt.Errorf("task %s failed with error: %s", taskID, task.Error)
		}

		// check if the context is done before sleeping
		select {
		// handles both timeout and cancellation
		case <-pollingCtx.Done():
			return nil, pollingCtx.Err()
		case <-time.After(pollingInterval):
		}
	}
}
