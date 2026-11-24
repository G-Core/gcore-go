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

// NewAndPoll creates a placement group, waits for all tasks, and returns the group.
// The tasks must report exactly one distinct ID in created_resources.servergroups.
// All tasks share one configured polling timeout and respect the caller's deadline. After the tasks finish,
// NewAndPoll calls Get once and returns its result or error. Each HTTP request uses the SDK retry policy,
// including [option.WithMaxRetries].
//
// Dry runs are rejected before any request. Use [PlacementGroupService.New] instead.
// [option.WithResponseBodyInto] captures only the final group JSON. [option.WithResponseInto] captures the
// action and then the Get responses, but task reads do not overwrite it.
func (r *PlacementGroupService) NewAndPoll(ctx context.Context, params PlacementGroupNewParams, opts ...option.RequestOption) (*PlacementGroup, error) {
	allOpts := slices.Concat(r.Options, opts)
	precfg, err := preparePoll("NewAndPoll", "New", params.DryRun, params, allOpts)
	if err != nil {
		return nil, err
	}
	svc := NewPlacementGroupService()
	resource, err := svc.New(ctx, params, requestconfig.ExcludeResponseBodyInto(allOpts...)...)
	if err != nil {
		return nil, err
	}
	tasks, err := waitForTasks(ctx, resource.Tasks, "placement group task", precfg, allOpts)
	if err != nil {
		return nil, err
	}
	var id string
	for _, task := range tasks {
		for _, candidate := range task.CreatedResources.Servergroups {
			if candidate == "" {
				continue
			}
			if id != "" && id != candidate {
				return nil, errors.New("expected one placement group ID in created_resources.servergroups, got multiple IDs")
			}
			id = candidate
		}
	}
	if id == "" {
		return nil, errors.New("expected a placement group ID in created_resources.servergroups, got none")
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	return getFinalResource(ctx, allOpts, func(ctx context.Context, opts ...option.RequestOption) (*PlacementGroup, error) {
		return svc.Get(ctx, id, PlacementGroupGetParams{ProjectID: params.ProjectID}, opts...)
	})
}

// DeleteAndPoll deletes a placement group and waits for all returned tasks.
// All tasks share one configured polling timeout and respect the caller's deadline. Each HTTP request uses
// the SDK retry policy, including [option.WithMaxRetries].
// Dry runs are rejected before any request. Use [PlacementGroupService.Delete] instead.
// [option.WithResponseInto] captures the delete response; [option.WithResponseBodyInto] is unused.
func (r *PlacementGroupService) DeleteAndPoll(ctx context.Context, placementGroupID string, params PlacementGroupDeleteParams, opts ...option.RequestOption) error {
	allOpts := slices.Concat(r.Options, opts)
	precfg, err := preparePoll("DeleteAndPoll", "Delete", params.DryRun, params, allOpts)
	if err != nil {
		return err
	}
	svc := NewPlacementGroupService()
	resource, err := svc.Delete(ctx, placementGroupID, params, requestconfig.ExcludeResponseBodyInto(allOpts...)...)
	if err != nil {
		return err
	}
	_, err = waitForTasks(ctx, resource.Tasks, "placement group task", precfg, allOpts)
	return err
}
