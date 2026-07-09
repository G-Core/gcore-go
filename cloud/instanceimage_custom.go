// Custom code. This file is not generated and is preserved across codegen runs.
// It isolates hand-written *AndPoll convenience methods from generated code to
// eliminate merge conflicts.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
)

// DeleteAndPoll delete the image and poll for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *InstanceImageService) DeleteAndPoll(ctx context.Context, imageID string, body InstanceImageDeleteParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (Delete returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Delete(ctx, imageID, body, actionOpts...)
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

// NewFromVolumeAndPoll create image from volume and poll for the result
func (r *InstanceImageService) NewFromVolumeAndPoll(ctx context.Context, params InstanceImageNewFromVolumeParams, opts ...option.RequestOption) (v *Image, err error) {
	// Exclude WithResponseBodyInto for the action (NewFromVolume returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.NewFromVolume(ctx, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams InstanceImageGetParams
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

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Images) != 1 {
		return nil, errors.New("expected exactly one image to be created in a task")
	}
	resourceID := task.CreatedResources.Images[0]

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, resourceID, getParams, getOpts...)
}

// UploadAndPoll upload image and poll for the task completion. Use the [TaskService.Poll] method if you need to poll
// for all tasks.
func (r *InstanceImageService) UploadAndPoll(ctx context.Context, params InstanceImageUploadParams, opts ...option.RequestOption) (v *Image, err error) {
	// Exclude WithResponseBodyInto for the action (Upload returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Upload(ctx, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams InstanceImageGetParams
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
	task, err := newTaskService(r.Options...).Poll(ctx, taskID, pollOpts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Images) != 1 {
		return nil, errors.New("expected exactly one image to be created in a task")
	}
	resourceID := task.CreatedResources.Images[0]

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	// Poll using a copy of the options that excludes WithResponseBodyInto so the
	// typed *Image is deserialized and its status/size can be inspected (with
	// WithResponseBodyInto set the typed value is nil and only the raw body is
	// populated).
	pollGetOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)

	// Reuse the same polling configuration as TaskService.Poll
	// (WithPollingIntervalSeconds / WithPollingTimeoutSeconds).
	pollingInterval := time.Duration(precfg.PollingIntervalSeconds) * time.Second
	if pollingInterval < time.Second {
		pollingInterval = time.Second
	}
	pollingCtx := ctx
	if precfg.PollingTimeoutSeconds > 0 {
		var cancel context.CancelFunc
		pollingCtx, cancel = context.WithTimeout(ctx, time.Duration(precfg.PollingTimeoutSeconds)*time.Second)
		defer cancel()
	}

	// The upload task is marked complete as soon as the image bytes are handed off
	// to the image store; the image itself then transitions queued -> saving ->
	// active asynchronously and only reports its final `size` once active. Polling
	// just the task therefore returns a transient image (status:"saving", size:0).
	// Poll the image until it settles so callers observe fully-resolved fields.
	for {
		v, err = r.Get(pollingCtx, resourceID, getParams, pollGetOpts...)
		if err != nil {
			return
		}
		if v.Status == "active" && v.Size > 0 {
			// Re-fetch honoring the caller's options (e.g. WithResponseBodyInto) so
			// the settled image is returned in the requested form.
			return r.Get(pollingCtx, resourceID, getParams, getOpts...)
		}
		switch v.Status {
		case "active", "queued", "saving":
			// still settling (pre-active, or active but size not yet populated)
		default:
			return v, fmt.Errorf("image %s entered unexpected status %q while waiting for it to become active", resourceID, v.Status)
		}
		select {
		case <-pollingCtx.Done():
			return v, pollingCtx.Err()
		case <-time.After(pollingInterval):
		}
	}
}
