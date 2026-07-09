// Custom code. This file is not generated and is preserved across codegen runs.
// It provides a shared helper for the *AndPoll convenience methods so they can
// poll tasks without a generated struct field, keeping custom code fully isolated
// from generated code.

package cloud

import "github.com/G-Core/gcore-go/option"

// newTaskService builds an addressable *TaskService from the calling service's
// request options. The *AndPoll methods use it to poll task completion without
// depending on any generated struct field.
func newTaskService(opts ...option.RequestOption) *TaskService {
	svc := NewTaskService(opts...)
	return &svc
}
