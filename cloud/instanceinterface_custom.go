// Custom code. This file is not generated and is preserved across codegen runs.
// It isolates hand-written *AndPoll convenience methods from generated code to
// eliminate merge conflicts.

package cloud

import (
	"context"
	"errors"
	"slices"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
)

// AttachAndPoll attach interface to instance and poll for the completion of the first task. Use the [TaskService.Poll]
// method if you need to poll for all tasks. The interfaces list is paginated; call [InstanceInterfaceService.List]
// (or [InstanceInterfaceService.ListAutoPaging]) afterwards to read the resulting interfaces.
func (r *InstanceInterfaceService) AttachAndPoll(ctx context.Context, instanceID string, params InstanceInterfaceAttachParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (Attach returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Attach(ctx, instanceID, params, actionOpts...)
	if err != nil {
		return err
	}

	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	_, err = newTaskService(r.Options...).Poll(ctx, taskID, pollOpts...)
	return err
}

// DetachAndPoll interface from instance and poll for the completion of the first task. Use the [TaskService.Poll]
// method if you need to poll for all tasks. The interfaces list is paginated; call [InstanceInterfaceService.List]
// (or [InstanceInterfaceService.ListAutoPaging]) afterwards to read the resulting interfaces.
func (r *InstanceInterfaceService) DetachAndPoll(ctx context.Context, instanceID string, params InstanceInterfaceDetachParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (Detach returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Detach(ctx, instanceID, params, actionOpts...)
	if err != nil {
		return err
	}

	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	_, err = newTaskService(r.Options...).Poll(ctx, taskID, pollOpts...)
	return err
}
