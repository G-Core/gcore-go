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
	"github.com/G-Core/gcore-go/packages/param"
)

// NewAndPoll create bare metal server and poll for the result
func (r *BaremetalServerService) NewAndPoll(ctx context.Context, params BaremetalServerNewParams, opts ...option.RequestOption) (v *BaremetalServer, err error) {
	// Exclude WithResponseBodyInto for the action (New returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.New(ctx, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
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

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Instances) != 1 {
		return nil, errors.New("expected exactly one bare metal server to be created in a task")
	}
	resourceID := task.CreatedResources.Instances[0]

	// Use List to find the created server
	var listParams BaremetalServerListParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	listParams.ProjectID = params.ProjectID
	listParams.RegionID = params.RegionID
	listParams.Uuid = param.NewOpt(resourceID)

	// Exclude WithResponseBodyInto and clear request body for List
	listOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	servers, err := r.List(ctx, listParams, listOpts...)
	if err != nil {
		return
	}

	if len(servers.Results) == 0 {
		return nil, errors.New("server not found after creation")
	}

	if err := requestconfig.WriteResponseBodyInto(opts, []byte(servers.Results[0].RawJSON())); err != nil {
		return nil, err
	}

	return &servers.Results[0], nil
}

// RebuildAndPoll rebuild bare metal server and poll for the completion of the first task.  Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *BaremetalServerService) RebuildAndPoll(ctx context.Context, serverID string, params BaremetalServerRebuildParams, opts ...option.RequestOption) (v *BaremetalServer, err error) {
	// Exclude WithResponseBodyInto for the action (Rebuild returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Rebuild(ctx, serverID, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}

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

	// Use List to find the rebuilt server
	var listParams BaremetalServerListParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	listParams.ProjectID = params.ProjectID
	listParams.RegionID = params.RegionID
	listParams.Uuid = param.NewOpt(serverID)

	// Exclude WithResponseBodyInto and clear request body for List
	listOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	servers, err := r.List(ctx, listParams, listOpts...)
	if err != nil {
		return
	}

	if len(servers.Results) == 0 {
		return nil, errors.New("server not found after rebuild")
	}

	if err := requestconfig.WriteResponseBodyInto(opts, []byte(servers.Results[0].RawJSON())); err != nil {
		return nil, err
	}

	return &servers.Results[0], nil
}

// DeleteAndPoll deletes a bare metal server and polls for the completion of the first task.
// Use the [TaskService.Poll] method if you need to poll for all tasks.
func (r *BaremetalServerService) DeleteAndPoll(ctx context.Context, serverID string, params BaremetalServerDeleteParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (Delete returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Delete(ctx, serverID, params, actionOpts...)
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
