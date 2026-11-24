// Custom code. This file is not generated and is preserved across codegen runs.
// It isolates hand-written *AndPoll convenience methods from generated code to
// eliminate merge conflicts.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// NewAndPoll creates a port, waits for the create task and returns the new port. It does not support dry
// runs, because a dry run creates no task. Use the [PortService.New] method for a dry run.
func (r *PortService) NewAndPoll(ctx context.Context, params PortNewParams, opts ...option.RequestOption) (res *Port, err error) {
	if params.DryRun.Valid() && params.DryRun.Value {
		return nil, errors.New("NewAndPoll does not support DryRun, use New instead")
	}

	// Exclude WithResponseBodyInto for the action (New returns TaskIDList, must deserialize properly)
	allOpts := slices.Concat(r.Options, opts)
	svc := NewPortService()
	resource, err := svc.New(ctx, params, requestconfig.ExcludeResponseBodyInto(allOpts...)...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(allOpts...)
	if err != nil {
		return
	}
	var getParams PortGetParams
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

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Ports) == 0 {
		return nil, errors.New("expected a port to be created in a task")
	}
	resourceID := task.CreatedResources.Ports[0]

	return getFinalResource(ctx, allOpts, func(ctx context.Context, opts ...option.RequestOption) (*Port, error) {
		return svc.Get(ctx, resourceID, getParams, opts...)
	})
}

// UpdateAndPoll updates a port, waits for the first task and returns the port. Use the [TaskService.Poll]
// method if you need to wait for all tasks. An update that changes nothing creates no task, so the method
// returns the port at once.
//
// After the task, the method calls Get once and returns its result or error.
func (r *PortService) UpdateAndPoll(ctx context.Context, portID string, params PortUpdateParams, opts ...option.RequestOption) (res *Port, err error) {
	allOpts := slices.Concat(r.Options, opts)
	svc := NewPortService()
	// Exclude WithResponseBodyInto for the action (Update returns TaskIDList, must deserialize properly)
	resource, err := svc.Update(ctx, portID, params, requestconfig.ExcludeResponseBodyInto(allOpts...)...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(allOpts...)
	if err != nil {
		return
	}
	var getParams PortGetParams
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

	return getFinalResource(ctx, allOpts, func(ctx context.Context, opts ...option.RequestOption) (*Port, error) {
		return svc.Get(ctx, portID, getParams, opts...)
	})
}

// DeleteAndPoll deletes a port and waits for the first task. Use the [TaskService.Poll] method if you need
// to wait for all tasks.
func (r *PortService) DeleteAndPoll(ctx context.Context, portID string, params PortDeleteParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (Delete returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Delete(ctx, portID, params, actionOpts...)
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

// ConvertAndPoll converts a port to a floating IP, waits for the convert task and returns the new floating
// IP. The task deletes the port, so the port ID is not valid after the call. Floating IPs have no global
// view route, so the method gets the floating IP with the [FloatingIPService.Get] method in the region of
// the task.
func (r *PortService) ConvertAndPoll(ctx context.Context, portID string, params PortConvertParams, opts ...option.RequestOption) (res *FloatingIP, err error) {
	// Exclude WithResponseBodyInto for the action (Convert returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Convert(ctx, portID, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams FloatingIPGetParams
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
	task, err := newTaskService(r.Options...).Poll(ctx, taskID, pollOpts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Floatingips) == 0 {
		return nil, errors.New("expected a floating IP to be created in a task")
	}
	if !task.JSON.RegionID.Valid() {
		return nil, fmt.Errorf("task %s has no region ID", taskID)
	}
	// the region comes from the task, not from the client, because the port route has no region
	getParams.RegionID = param.NewOpt(task.RegionID)

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	floatingIPs := NewFloatingIPService(r.Options...)
	return floatingIPs.Get(ctx, task.CreatedResources.Floatingips[0], getParams, getOpts...)
}
