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

// NewAndPoll creates a new GPU bare metal cluster and polls for completion. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *GPUBaremetalClusterService) NewAndPoll(ctx context.Context, params GPUBaremetalClusterNewParams, opts ...option.RequestOption) (v *GPUBaremetalCluster, err error) {
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
	var getParams GPUBaremetalClusterGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

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

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Clusters) != 1 {
		return nil, errors.New("expected exactly one cluster to be created in a task")
	}
	clusterID := task.CreatedResources.Clusters[0]

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, clusterID, getParams, getOpts...)
}

// ActionAndPoll performs an action on a bare metal GPU cluster and polls for completion of the first task. Use the
// [TaskService.Poll] method if you need to poll for all tasks.
func (r *GPUBaremetalClusterService) ActionAndPoll(ctx context.Context, clusterID string, params GPUBaremetalClusterActionParams, opts ...option.RequestOption) (v *GPUBaremetalCluster, err error) {
	// Exclude WithResponseBodyInto for the action (Action returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Action(ctx, clusterID, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams GPUBaremetalClusterGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

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
		return nil, err
	}

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, clusterID, getParams, getOpts...)
}

// ApplySettingsAndPoll applies the cluster's server settings to all of its servers and polls for completion of the
// first task. Use the [TaskService.Poll] method if you need to poll for all tasks.
//
// Patch the settings first with [GPUBaremetalClusterService.Update], then call this to roll them out. Applying
// settings re-images the servers, so params.MaxDisruption must be set to
// [GPUBaremetalClusterApplySettingsParamsMaxDisruptionRebuild] for the request to proceed.
func (r *GPUBaremetalClusterService) ApplySettingsAndPoll(ctx context.Context, clusterID string, params GPUBaremetalClusterApplySettingsParams, opts ...option.RequestOption) (v *GPUBaremetalCluster, err error) {
	// Exclude WithResponseBodyInto for the action (ApplySettings returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.ApplySettings(ctx, clusterID, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams GPUBaremetalClusterGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

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
		return nil, err
	}

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, clusterID, getParams, getOpts...)
}

// RebuildAndPoll rebuilds a GPU bare metal cluster and polls for completion of the first task. Use the
// [TaskService.Poll] method if you need to poll for all tasks.
//
// Deprecated: use [GPUBaremetalClusterService.ApplySettingsAndPoll] instead.
func (r *GPUBaremetalClusterService) RebuildAndPoll(ctx context.Context, clusterID string, params GPUBaremetalClusterRebuildParams, opts ...option.RequestOption) (v *GPUBaremetalCluster, err error) {
	// Exclude WithResponseBodyInto for the action (Rebuild returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Rebuild(ctx, clusterID, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams GPUBaremetalClusterGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

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
	return r.Get(ctx, clusterID, getParams, getOpts...)
}

// ResizeAndPoll resizes a GPU bare metal cluster and polls for completion of the first task. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *GPUBaremetalClusterService) ResizeAndPoll(ctx context.Context, clusterID string, params GPUBaremetalClusterResizeParams, opts ...option.RequestOption) (v *GPUBaremetalCluster, err error) {
	// Exclude WithResponseBodyInto for the action (Resize returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Resize(ctx, clusterID, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams GPUBaremetalClusterGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

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
	return r.Get(ctx, clusterID, getParams, getOpts...)
}

// DeleteAndPoll deletes a bare metal GPU cluster and polls for completion of the first task. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *GPUBaremetalClusterService) DeleteAndPoll(ctx context.Context, clusterID string, params GPUBaremetalClusterDeleteParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (Delete returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Delete(ctx, clusterID, params, actionOpts...)
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
