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

// NewAndPoll creates a load balancer pool member and polls until the operation completes
func (r *LoadBalancerPoolMemberService) NewAndPoll(ctx context.Context, poolID string, params LoadBalancerPoolMemberNewParams, opts ...option.RequestOption) (v *Member, err error) {
	// Exclude WithResponseBodyInto for the action (New returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.New(ctx, poolID, params, actionOpts...)
	if err != nil {
		return
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams LoadBalancerPoolMemberGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID
	getParams.PoolID = poolID

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

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Members) != 1 {
		return nil, errors.New("expected exactly one member to be created in a task")
	}
	memberID := task.CreatedResources.Members[0]

	// Exclude WithResponseBodyInto and clear request body for Get: with a custom
	// response target the request layer skips decoding, so Get would return a nil
	// member. Fetch into our own value, then hand the same body to the caller's target.
	getOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	member, err := r.Get(ctx, memberID, getParams, getOpts...)
	if err != nil {
		return nil, err
	}
	if err := requestconfig.WriteResponseBodyInto(opts, []byte(member.RawJSON())); err != nil {
		return nil, err
	}
	return member, nil
}

// UpdateAndPoll updates a load balancer pool member and polls until the operation completes
func (r *LoadBalancerPoolMemberService) UpdateAndPoll(ctx context.Context, memberID string, params LoadBalancerPoolMemberUpdateParams, opts ...option.RequestOption) (v *Member, err error) {
	// Exclude WithResponseBodyInto for the action (Update returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Update(ctx, memberID, params, actionOpts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams LoadBalancerPoolMemberGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID
	getParams.PoolID = params.PoolID

	// A PATCH whose fields already match the member's current state creates no task and returns an
	// empty task list, so only poll when there is a task to poll for.
	if len(resource.Tasks) > 0 {
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
	}

	// Exclude WithResponseBodyInto and clear request body for Get: with a custom
	// response target the request layer skips decoding, so Get would return a nil
	// member. Fetch into our own value, then hand the same body to the caller's target.
	getOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	member, err := r.Get(ctx, memberID, getParams, getOpts...)
	if err != nil {
		return nil, err
	}
	if err := requestconfig.WriteResponseBodyInto(opts, []byte(member.RawJSON())); err != nil {
		return nil, err
	}
	return member, nil
}

// ReplaceAndPoll replaces the pool's entire member list and polls until the operation completes,
// then returns the resulting members. The API handles a replace as a single pool patch: it returns
// one task when the member set changes and an empty task list when the requested set already
// matches (members are matched by address and port). Both cases are covered by polling whatever
// the action returns, so no task is a completed no-op rather than an error.
func (r *LoadBalancerPoolMemberService) ReplaceAndPoll(ctx context.Context, poolID string, params LoadBalancerPoolMemberReplaceParams, opts ...option.RequestOption) (v []Member, err error) {
	// Exclude WithResponseBodyInto for the action (Replace returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Replace(ctx, poolID, params, actionOpts...)
	if err != nil {
		return nil, err
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	var listParams LoadBalancerPoolMemberListParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	listParams.ProjectID = params.ProjectID
	listParams.RegionID = params.RegionID

	// A replace that changes nothing creates no task and returns an empty task list.
	// Exclude WithResponseBodyInto and clear request body for Poll and List.
	quietOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	taskService := newTaskService(r.Options...)
	for _, taskID := range resource.Tasks {
		if _, err = taskService.Poll(ctx, taskID, quietOpts...); err != nil {
			return nil, err
		}
	}

	page, err := r.List(ctx, poolID, listParams, quietOpts...)
	if err != nil {
		return nil, err
	}
	if err := requestconfig.WriteResponseBodyInto(opts, []byte(page.RawJSON())); err != nil {
		return nil, err
	}
	return page.Results, nil
}

// DeleteAndPoll deletes a load balancer pool member and polls for completion
func (r *LoadBalancerPoolMemberService) DeleteAndPoll(ctx context.Context, memberID string, params LoadBalancerPoolMemberDeleteParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (Delete returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Delete(ctx, memberID, params, actionOpts...)
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
