// Custom code. This file is not generated and is preserved across codegen runs.
// It provides a shared helper for the *AndPoll convenience methods so they can
// poll tasks without a generated struct field, keeping custom code fully isolated
// from generated code.

package cloud

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// newTaskService builds an addressable *TaskService from the calling service's
// request options. The *AndPoll methods use it to poll task completion without
// depending on any generated struct field.
func newTaskService(opts ...option.RequestOption) *TaskService {
	svc := NewTaskService(opts...)
	return &svc
}

// getFinalResource keeps intermediate response capture private and exposes only the final resource JSON
func getFinalResource[T interface{ RawJSON() string }](ctx context.Context, opts []option.RequestOption, get func(context.Context, ...option.RequestOption) (*T, error)) (*T, error) {
	getOpts := append(requestconfig.ExcludeResponseBodyInto(opts...), requestconfig.WithoutRequestBody())
	resource, err := get(ctx, getOpts...)
	if err != nil {
		return nil, err
	}
	if err := requestconfig.WriteResponseBodyInto(opts, []byte((*resource).RawJSON())); err != nil {
		return nil, err
	}
	return resource, nil
}

// preparePoll checks the effective dry_run query before the helper sends a request
func preparePoll(helper, rawMethod string, dryRun param.Opt[bool], params apiquery.Queryer, opts []option.RequestOption) (requestconfig.RequestConfig, error) {
	reject := func() error {
		return fmt.Errorf("%s does not support DryRun or a dry_run query option, because a dry run creates no task; use %s instead", helper, rawMethod)
	}
	if dryRun.Valid() && dryRun.Value {
		return requestconfig.RequestConfig{}, reject()
	}
	query, err := params.URLQuery()
	if err != nil {
		return requestconfig.RequestConfig{}, err
	}
	probe := requestconfig.NewMinimalConfig()
	probe.Request.URL.RawQuery = query.Encode()
	for _, opt := range opts {
		_ = opt.Apply(probe)
	}
	for _, value := range probe.Request.URL.Query()["dry_run"] {
		switch strings.ToLower(strings.TrimSpace(value)) {
		case "false", "0", "f", "n", "no", "off":
		default:
			return requestconfig.RequestConfig{}, reject()
		}
	}
	return requestconfig.PreRequestOptions(opts...)
}

// waitForTasks gives all task reads one polling budget and keeps their responses private
func waitForTasks(ctx context.Context, taskIDs []string, taskKind string, precfg requestconfig.RequestConfig, opts []option.RequestOption) ([]*Task, error) {
	if len(taskIDs) == 0 {
		return nil, fmt.Errorf("expected at least one %s", taskKind)
	}
	if precfg.PollingTimeoutSeconds > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(precfg.PollingTimeoutSeconds)*time.Second)
		defer cancel()
	}
	pollOpts := append(requestconfig.ExcludeResponseBodyInto(opts...), option.WithResponseInto(nil), requestconfig.WithoutRequestBody())
	taskService := newTaskService(pollOpts...)
	tasks := make([]*Task, 0, len(taskIDs))
	for _, taskID := range taskIDs {
		task, err := taskService.Poll(ctx, taskID)
		if err != nil {
			return nil, fmt.Errorf("wait for %s %s: %w", taskKind, taskID, err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
