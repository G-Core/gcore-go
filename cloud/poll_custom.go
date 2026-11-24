// Custom code. This file is not generated and is preserved across codegen runs.
// It provides a shared helper for the *AndPoll convenience methods so they can
// poll tasks without a generated struct field, keeping custom code fully isolated
// from generated code.

package cloud

import (
	"context"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
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
