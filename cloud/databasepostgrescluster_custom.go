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

// NewAndPoll creates a new PostgreSQL cluster and polls for completion
func (r *DatabasePostgresClusterService) NewAndPoll(ctx context.Context, params DatabasePostgresClusterNewParams, opts ...option.RequestOption) (v *PostgresCluster, err error) {
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
	var getParams DatabasePostgresClusterGetParams
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

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.PostgreSQLClusters) != 1 {
		return nil, errors.New("expected exactly one postgres cluster to be created in a task")
	}
	resourceID := task.CreatedResources.PostgreSQLClusters[0]

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, resourceID, getParams, getOpts...)
}

// UpdateAndPoll updates a PostgreSQL cluster and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *DatabasePostgresClusterService) UpdateAndPoll(ctx context.Context, clusterName string, params DatabasePostgresClusterUpdateParams, opts ...option.RequestOption) (v *PostgresCluster, err error) {
	// Exclude WithResponseBodyInto for the action (Update returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Update(ctx, clusterName, params, actionOpts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams DatabasePostgresClusterGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	// Depending on which fields were being updated the Update method might not create Tasks. For instance, if the user
	// only updates tags no task will be created. Therefore, we only poll when there are tasks to poll for.
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

	// Clear request body for Get
	getOpts := slices.Concat(opts, []option.RequestOption{requestconfig.WithoutRequestBody()})
	return r.Get(ctx, clusterName, getParams, getOpts...)
}

// DeleteAndPoll deletes a PostgreSQL cluster and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *DatabasePostgresClusterService) DeleteAndPoll(ctx context.Context, clusterName string, body DatabasePostgresClusterDeleteParams, opts ...option.RequestOption) error {
	// Exclude WithResponseBodyInto for the action (Delete returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.Delete(ctx, clusterName, body, actionOpts...)
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
