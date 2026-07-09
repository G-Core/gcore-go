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

// DeleteAndPoll deletes a bare metal GPU server from cluster and polls for completion of the first task. Use the
// [TaskService.Poll] method if you need to poll for all tasks.
func (r *GPUBaremetalClusterServerService) DeleteAndPoll(ctx context.Context, instanceID string, params GPUBaremetalClusterServerDeleteParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (Delete returns TaskIDList, must deserialize properly)
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

// RebuildAndPoll rebuilds a bare metal GPU cluster server and polls for completion of the first task. Use the
// [TaskService.Poll] method if you need to poll for all tasks.
func (r *GPUBaremetalClusterServerService) RebuildAndPoll(ctx context.Context, serverID string, body GPUBaremetalClusterServerRebuildParams, opts ...option.RequestOption) (v *GPUBaremetalClusterServer, err error) {
	// Exclude WithResponseBodyInto for the action (Rebuild returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Rebuild(ctx, serverID, body, actionOpts...)
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
	var listParams GPUBaremetalClusterServerListParams
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	listParams.ProjectID = body.ProjectID
	listParams.RegionID = body.RegionID
	listParams.Uuids = []string{serverID}

	// Exclude WithResponseBodyInto and clear request body for List
	listOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	page, err := r.List(ctx, body.ClusterID, listParams, listOpts...)
	if err != nil {
		return
	}

	if len(page.Results) == 0 {
		return nil, errors.New("server not found after rebuild")
	}

	if err := requestconfig.WriteResponseBodyInto(opts, []byte(page.Results[0].RawJSON())); err != nil {
		return nil, err
	}

	return &page.Results[0], nil
}

// ReplaceAndPoll replaces a bare metal GPU cluster server and polls for completion of the first task. Use the
// [TaskService.Poll] method if you need to poll for all tasks.
func (r *GPUBaremetalClusterServerService) ReplaceAndPoll(ctx context.Context, serverID string, body GPUBaremetalClusterServerReplaceParams, opts ...option.RequestOption) (v *GPUBaremetalClusterServer, err error) {
	// Exclude WithResponseBodyInto for the action (Replace returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Replace(ctx, serverID, body, actionOpts...)
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
	task, err := newTaskService(r.Options...).Poll(ctx, taskID, pollOpts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Instances) != 1 {
		return nil, errors.New("expected exactly one instance to be created in a task")
	}
	newServerID := task.CreatedResources.Instances[0]

	// Use List to find the new server
	var listParams GPUBaremetalClusterServerListParams
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	listParams.ProjectID = body.ProjectID
	listParams.RegionID = body.RegionID
	listParams.Uuids = []string{newServerID}

	// Exclude WithResponseBodyInto and clear request body for List
	listOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	page, err := r.List(ctx, body.ClusterID, listParams, listOpts...)
	if err != nil {
		return
	}

	if len(page.Results) == 0 {
		return nil, errors.New("server not found after replace")
	}

	if err := requestconfig.WriteResponseBodyInto(opts, []byte(page.Results[0].RawJSON())); err != nil {
		return nil, err
	}

	return &page.Results[0], nil
}
