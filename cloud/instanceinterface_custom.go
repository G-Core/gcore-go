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
// method if you need to poll for all tasks.
func (r *InstanceInterfaceService) AttachAndPoll(ctx context.Context, instanceID string, params InstanceInterfaceAttachParams, opts ...option.RequestOption) (v *NetworkInterfaceList, err error) {
	// Exclude WithResponseBodyInto for the action (Attach returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Attach(ctx, instanceID, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var listParams InstanceInterfaceListParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	listParams.ProjectID = params.ProjectID
	listParams.RegionID = params.RegionID

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

	// Exclude WithResponseBodyInto and clear request body for List
	listOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	result, err := r.List(ctx, instanceID, listParams, listOpts...)
	if err != nil {
		return nil, err
	}

	if err := requestconfig.WriteResponseBodyInto(opts, []byte(result.RawJSON())); err != nil {
		return nil, err
	}

	return result, nil
}

// DetachAndPoll interface from instance and poll for the completion of the first task. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *InstanceInterfaceService) DetachAndPoll(ctx context.Context, instanceID string, params InstanceInterfaceDetachParams, opts ...option.RequestOption) (v *NetworkInterfaceList, err error) {
	// Exclude WithResponseBodyInto for the action (Detach returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Detach(ctx, instanceID, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var listParams InstanceInterfaceListParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	listParams.ProjectID = params.ProjectID
	listParams.RegionID = params.RegionID

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

	// Exclude WithResponseBodyInto and clear request body for List
	listOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	result, err := r.List(ctx, instanceID, listParams, listOpts...)
	if err != nil {
		return nil, err
	}

	if err := requestconfig.WriteResponseBodyInto(opts, []byte(result.RawJSON())); err != nil {
		return nil, err
	}

	return result, nil
}
