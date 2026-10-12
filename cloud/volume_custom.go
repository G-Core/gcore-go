// Custom code. This file is not generated and is preserved across codegen runs.
// It isolates hand-written *AndPoll convenience methods from generated code to
// eliminate merge conflicts.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
)

// NewAndPoll creates a new volume and polls the corresponding task until it is completed.
func (r *VolumeService) NewAndPoll(ctx context.Context, params VolumeNewParams, opts ...option.RequestOption) (res *Volume, err error) {
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
	var getParams VolumeGetParams
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

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Volumes) != 1 {
		return nil, errors.New("expected exactly one network to be created in a task")
	}
	resourceID := task.CreatedResources.Volumes[0]

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, resourceID, getParams, getOpts...)
}

// DeleteAndPoll deletes a volume and polls the corresponding task until it is completed.
// Use the [TaskService.Poll] method if you need to poll for all tasks.
func (r *VolumeService) DeleteAndPoll(ctx context.Context, volumeID string, params VolumeDeleteParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (Delete returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Delete(ctx, volumeID, params, actionOpts...)
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

// AttachToInstanceAndPoll attaches the volume to instance and polls the corresponding task until it is completed. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *VolumeService) AttachToInstanceAndPoll(ctx context.Context, volumeID string, params VolumeAttachToInstanceParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (AttachToInstance returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.AttachToInstance(ctx, volumeID, params, actionOpts...)
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

// DetachFromInstanceAndPoll detaches the volume to instance and polls the corresponding task until it is completed.
// Use the [TaskService.Poll] method if you need to poll for all tasks.
func (r *VolumeService) DetachFromInstanceAndPoll(ctx context.Context, volumeID string, params VolumeDetachFromInstanceParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (DetachFromInstance returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.DetachFromInstance(ctx, volumeID, params, actionOpts...)
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

// ResizeAndPoll increases the size of a volume and polls the corresponding task until it is completed. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *VolumeService) ResizeAndPoll(ctx context.Context, volumeID string, params VolumeResizeParams, opts ...option.RequestOption) (res *Volume, err error) {
	// Exclude WithResponseBodyInto for the action (Resize returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Resize(ctx, volumeID, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams VolumeGetParams
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
	return r.Get(ctx, volumeID, getParams, getOpts...)
}

// ChangeTypeAndPoll changes the type of a volume and polls the volume until the
// change is complete. The retype endpoint creates no task: it returns while the
// volume is still "retyping" with its old type, so this polls the volume itself
// until it reports the new type. Polling honors WithPollingIntervalSeconds and
// WithPollingTimeoutSeconds.
func (r *VolumeService) ChangeTypeAndPoll(ctx context.Context, volumeID string, params VolumeChangeTypeParams, opts ...option.RequestOption) (res *Volume, err error) {
	// Exclude WithResponseBodyInto for the action so its typed response can be inspected
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	v, err := r.ChangeType(ctx, volumeID, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams VolumeGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	// Poll with WithResponseBodyInto excluded so the typed *Volume can be
	// inspected (with a custom response target the request layer skips decoding)
	pollGetOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)

	// Reuse the same polling configuration as TaskService.Poll
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

	target := string(params.VolumeType)
	for {
		if strings.HasPrefix(string(v.Status), "error") {
			return v, fmt.Errorf("volume %s entered status %q while changing its type to %q", volumeID, v.Status, target)
		}
		if v.Status != VolumeStatusRetyping {
			if !strings.EqualFold(v.VolumeType, target) {
				// The retype ended without taking effect (rolled back).
				return v, fmt.Errorf("volume %s left status %q with type %q instead of %q", volumeID, VolumeStatusRetyping, v.VolumeType, target)
			}
			// Hand the settled volume's body to the caller's target, if any.
			if err := requestconfig.WriteResponseBodyInto(opts, []byte(v.RawJSON())); err != nil {
				return nil, err
			}
			return v, nil
		}
		select {
		case <-pollingCtx.Done():
			return v, pollingCtx.Err()
		case <-time.After(pollingInterval):
		}
		v, err = r.Get(pollingCtx, volumeID, getParams, pollGetOpts...)
		if err != nil {
			return
		}
	}
}
