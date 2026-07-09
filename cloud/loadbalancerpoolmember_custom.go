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
	var getParams LoadBalancerPoolGetParams
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

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.Members) != 1 {
		return nil, errors.New("expected exactly one member to be created in a task")
	}
	memberID := task.CreatedResources.Members[0]

	// Exclude WithResponseBodyInto and clear request body for Get
	getOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	poolService := NewLoadBalancerPoolService(slices.Concat(r.Options, getOpts)...)
	pool, err := poolService.Get(ctx, poolID, getParams, getOpts...)
	if err != nil {
		return
	}

	for _, member := range pool.Members {
		if member.ID == memberID {
			if err := requestconfig.WriteResponseBodyInto(opts, []byte(member.RawJSON())); err != nil {
				return nil, err
			}
			return &member, nil
		}
	}

	return nil, fmt.Errorf("member %s not found in pool %s after creation", memberID, poolID)
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
