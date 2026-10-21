// Custom code. This file is not generated and is preserved across codegen runs.
// It isolates hand-written *AndPoll convenience methods from generated code to
// eliminate merge conflicts.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
)

// NewAndPoll creates a load balancer pool health monitor and polls until the operation completes,
// then returns the health monitor. A pool has at most one health monitor and the GET is keyed by
// the pool, so the created id from the task is not needed to fetch it.
func (r *LoadBalancerPoolHealthMonitorService) NewAndPoll(ctx context.Context, poolID string, params LoadBalancerPoolHealthMonitorNewParams, opts ...option.RequestOption) (v *HealthMonitor, err error) {
	// Exclude WithResponseBodyInto for the action (New returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.New(ctx, poolID, params, actionOpts...)
	if err != nil {
		return
	}

	// Merge service options only to resolve defaults; the generated methods prepend r.Options
	// themselves, so passing the merged list on would apply them twice.
	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams LoadBalancerPoolHealthMonitorGetParams
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
	_, err = newTaskService(r.Options...).Poll(ctx, taskID, pollOpts...)
	if err != nil {
		return
	}

	return r.getAndWriteBody(ctx, poolID, getParams, opts...)
}

// UpdateAndPoll updates a load balancer pool health monitor and polls until the operation
// completes, then returns the health monitor. A PATCH whose fields already match the current state
// creates no task and returns an empty task list, so polling is skipped in that case.
func (r *LoadBalancerPoolHealthMonitorService) UpdateAndPoll(ctx context.Context, poolID string, params LoadBalancerPoolHealthMonitorUpdateParams, opts ...option.RequestOption) (v *HealthMonitor, err error) {
	// Exclude WithResponseBodyInto for the action (Update returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Update(ctx, poolID, params, actionOpts...)
	if err != nil {
		return
	}

	// Merge service options only to resolve defaults; the generated methods prepend r.Options
	// themselves, so passing the merged list on would apply them twice.
	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return
	}
	var getParams LoadBalancerPoolHealthMonitorGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

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

	return r.getAndWriteBody(ctx, poolID, getParams, opts...)
}

// DeleteAndPoll deletes a load balancer pool health monitor and polls until the API stops
// reporting it. The DELETE creates no task and answers 204 while the monitor is still reported for
// about a second; a create issued in that window fails with "already exists". Polling stops on the
// first 404 and keeps going through other GET errors, since the DELETE already succeeded; the last
// such error is returned only when the polling timeout runs out. Polling honors
// WithPollingIntervalSeconds and WithPollingTimeoutSeconds.
func (r *LoadBalancerPoolHealthMonitorService) DeleteAndPoll(ctx context.Context, poolID string, params LoadBalancerPoolHealthMonitorDeleteParams, opts ...option.RequestOption) error {
	err := r.Delete(ctx, poolID, params, opts...)
	if err != nil {
		return err
	}

	precfg, err := requestconfig.PreRequestOptions(slices.Concat(r.Options, opts)...)
	if err != nil {
		return err
	}
	var getParams LoadBalancerPoolHealthMonitorGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	getOpts := slices.Concat(
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

	var lastErr error
	for {
		_, err = r.Get(pollingCtx, poolID, getParams, getOpts...)
		var apiErr *Error
		if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
			return nil
		}
		if err != nil && pollingCtx.Err() == nil {
			lastErr = err
		}

		select {
		case <-pollingCtx.Done():
			if lastErr != nil {
				return fmt.Errorf("health monitor of pool %s is still reported after delete: %w (last error: %w)", poolID, pollingCtx.Err(), lastErr)
			}
			return fmt.Errorf("health monitor of pool %s is still reported after delete: %w", poolID, pollingCtx.Err())
		case <-time.After(pollingInterval):
		}
	}
}

// getAndWriteBody fetches the pool's health monitor and hands the raw body to the caller's
// WithResponseBodyInto target, if any. With a custom response target the request layer skips
// decoding, so Get would return a nil health monitor; fetch into our own value instead, then
// write the same body to the caller's target.
func (r *LoadBalancerPoolHealthMonitorService) getAndWriteBody(ctx context.Context, poolID string, getParams LoadBalancerPoolHealthMonitorGetParams, opts ...option.RequestOption) (*HealthMonitor, error) {
	getOpts := slices.Concat(
		requestconfig.ExcludeResponseBodyInto(opts...),
		[]option.RequestOption{requestconfig.WithoutRequestBody()},
	)
	hm, err := r.Get(ctx, poolID, getParams, getOpts...)
	if err != nil {
		return nil, err
	}
	if err := requestconfig.WriteResponseBodyInto(opts, []byte(hm.RawJSON())); err != nil {
		return nil, err
	}
	return hm, nil
}
