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

// NewAndPoll creates an instance, waits for the create task and returns the new instance. It does not
// support dry runs, because a dry run creates no task. Use the [InstanceService.New] method for a dry run.
//
// After the task, the method calls Get once and returns its result or error.
func (r *InstanceService) NewAndPoll(ctx context.Context, params InstanceNewParams, opts ...option.RequestOption) (res *Instance, err error) {
	if params.DryRun.Valid() && params.DryRun.Value {
		return nil, errors.New("NewAndPoll does not support DryRun, use New instead")
	}

	// Exclude WithResponseBodyInto for the action (New returns TaskIDList, must deserialize properly)
	allOpts := slices.Concat(r.Options, opts)
	svc := NewInstanceService()
	resource, err := svc.New(ctx, params, requestconfig.ExcludeResponseBodyInto(allOpts...)...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(allOpts...)
	if err != nil {
		return
	}
	var getParams InstanceGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	getParams.ProjectID = params.ProjectID

	if len(resource.Tasks) == 0 {
		return nil, errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := append(requestconfig.ExcludeResponseBodyInto(allOpts...), option.WithResponseInto(nil), requestconfig.WithoutRequestBody())
	task, err := newTaskService(pollOpts...).Poll(ctx, taskID)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Instances) == 0 {
		return nil, errors.New("expected an instance to be created in a task")
	}
	resourceID := task.CreatedResources.Instances[0]

	return getFinalResource(ctx, allOpts, func(ctx context.Context, opts ...option.RequestOption) (*Instance, error) {
		return svc.Get(ctx, resourceID, getParams, opts...)
	})
}

// UpdateAndPoll updates an instance, waits for the first task and returns the instance. Use the
// [TaskService.Poll] method if you need to wait for all tasks. If the response has no task, the method
// returns the instance at once. It does not support dry runs, because a dry run creates no task. Use the
// [InstanceService.Update] method for a dry run.
//
// After the task, the method calls Get once and returns its result or error.
func (r *InstanceService) UpdateAndPoll(ctx context.Context, instanceID string, params InstanceUpdateParams, opts ...option.RequestOption) (res *Instance, err error) {
	if params.DryRun.Valid() && params.DryRun.Value {
		return nil, errors.New("UpdateAndPoll does not support DryRun, use Update instead")
	}

	// Exclude WithResponseBodyInto for the action (Update returns TaskIDList, must deserialize properly)
	allOpts := slices.Concat(r.Options, opts)
	svc := NewInstanceService()
	resource, err := svc.Update(ctx, instanceID, params, requestconfig.ExcludeResponseBodyInto(allOpts...)...)
	if err != nil {
		return
	}
	precfg, err := requestconfig.PreRequestOptions(allOpts...)
	if err != nil {
		return
	}
	var getParams InstanceGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	getParams.ProjectID = params.ProjectID

	if len(resource.Tasks) > 0 {
		taskID := resource.Tasks[0]
		// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
		pollOpts := append(requestconfig.ExcludeResponseBodyInto(allOpts...), option.WithResponseInto(nil), requestconfig.WithoutRequestBody())
		_, err = newTaskService(pollOpts...).Poll(ctx, taskID)
		if err != nil {
			return nil, err
		}
	}

	return getFinalResource(ctx, allOpts, func(ctx context.Context, opts ...option.RequestOption) (*Instance, error) {
		return svc.Get(ctx, instanceID, getParams, opts...)
	})
}

// DeleteAndPoll deletes an instance and waits for the first task. Use the [TaskService.Poll] method if you
// need to wait for all tasks. It does not support dry runs, because a dry run creates no task. Use the
// [InstanceService.Delete] method for a dry run.
func (r *InstanceService) DeleteAndPoll(ctx context.Context, instanceID string, params InstanceDeleteParams, opts ...option.RequestOption) error {
	if params.DryRun.Valid() && params.DryRun.Value {
		return errors.New("DeleteAndPoll does not support DryRun, use Delete instead")
	}

	// Exclude WithResponseBodyInto for the action (Delete returns InstanceDeleteResponse, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Delete(ctx, instanceID, params, actionOpts...)
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

// ActionAndPoll runs a power action on an instance, waits for the first task and returns the instance.
// Use the [TaskService.Poll] method if you need to wait for all tasks. It does not support dry runs,
// because a dry run creates no task. Use the [InstanceService.Action] method for a dry run.
//
// After the task, the method calls Get once and returns its result or error.
func (r *InstanceService) ActionAndPoll(ctx context.Context, instanceID string, params InstanceActionParams, opts ...option.RequestOption) (res *Instance, err error) {
	if params.DryRun.Valid() && params.DryRun.Value {
		return nil, errors.New("ActionAndPoll does not support DryRun, use Action instead")
	}

	// Exclude WithResponseBodyInto for the action (Action returns TaskIDList, must deserialize properly)
	allOpts := slices.Concat(r.Options, opts)
	svc := NewInstanceService()
	resource, err := svc.Action(ctx, instanceID, params, requestconfig.ExcludeResponseBodyInto(allOpts...)...)
	if err != nil {
		return
	}
	precfg, err := requestconfig.PreRequestOptions(allOpts...)
	if err != nil {
		return
	}
	var getParams InstanceGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	getParams.ProjectID = params.ProjectID

	if len(resource.Tasks) == 0 {
		return nil, errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := append(requestconfig.ExcludeResponseBodyInto(allOpts...), option.WithResponseInto(nil), requestconfig.WithoutRequestBody())
	_, err = newTaskService(pollOpts...).Poll(ctx, taskID)
	if err != nil {
		return
	}

	return getFinalResource(ctx, allOpts, func(ctx context.Context, opts ...option.RequestOption) (*Instance, error) {
		return svc.Get(ctx, instanceID, getParams, opts...)
	})
}

// RebuildAndPoll rebuild the instance and poll for completion of the first task. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *InstanceService) RebuildAndPoll(ctx context.Context, instanceID string, params InstanceRebuildParams, opts ...option.RequestOption) (v *Instance, err error) {
	// Exclude WithResponseBodyInto for the action (Rebuild returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Rebuild(ctx, instanceID, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams InstanceGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	getParams.ProjectID = params.ProjectID

	if len(resource.Tasks) == 0 {
		return nil, errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	// Exclude WithResponseBodyInto and clear request body for Poll (returns Task, must deserialize properly)
	pollOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	_, err = newTaskService(r.Options...).Poll(ctx, taskID, pollOpts...)
	if err != nil {
		return
	}

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, instanceID, getParams, getOpts...)
}
