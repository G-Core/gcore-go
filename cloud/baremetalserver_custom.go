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

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// NewAndPoll creates a bare metal server, waits for every task of the create and returns the new server.
// It reads the server ID from created_resources.instances of the tasks. It returns an error when the
// response has no task, when a task fails, or when the tasks do not name exactly one server.
//
// All tasks share one configured polling timeout and respect the caller's deadline. After the tasks finish,
// NewAndPoll calls Get once and returns its result or error. Each HTTP request uses the SDK retry policy,
// including [option.WithMaxRetries].
//
// NewAndPoll does not support dry runs, because a dry run creates no task. With DryRun set to true, or with
// a dry_run query option such as option.WithQuery("dry_run", "true"), it returns an error and sends no
// request. Use the [BaremetalServerService.New] method for a dry run. [option.WithResponseBodyInto] captures
// only the final server JSON. [option.WithResponseInto] captures the action and then the Get responses, but
// task reads do not overwrite it.
func (r *BaremetalServerService) NewAndPoll(ctx context.Context, params BaremetalServerNewParams, opts ...option.RequestOption) (*BaremetalServer, error) {
	allOpts := slices.Concat(r.Options, opts)
	precfg, err := preparePoll("NewAndPoll", "New", params.DryRun, params, allOpts)
	if err != nil {
		return nil, err
	}
	svc := NewBaremetalServerService()

	// exclude WithResponseBodyInto, because New must decode the task list
	resource, err := svc.New(ctx, params, requestconfig.ExcludeResponseBodyInto(allOpts...)...)
	if err != nil {
		return nil, err
	}
	tasks, err := waitForTasks(ctx, resource.Tasks, "task", precfg, allOpts)
	if err != nil {
		return nil, err
	}
	serverID, err := createdBaremetalServerID(tasks)
	if err != nil {
		return nil, err
	}

	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	return getFinalResource(ctx, allOpts, func(ctx context.Context, opts ...option.RequestOption) (*BaremetalServer, error) {
		return svc.Get(ctx, serverID, BaremetalServerGetParams{ProjectID: params.ProjectID}, opts...)
	})
}

// UpdateAndPoll updates a bare metal server, waits for every task of the update and returns the server. It
// returns an error when the response has no task or when a task fails.
//
// All tasks share one configured polling timeout and respect the caller's deadline. After the tasks finish,
// UpdateAndPoll calls Get once and returns its result or error. Each HTTP request uses the SDK retry policy,
// including [option.WithMaxRetries].
//
// UpdateAndPoll does not support dry runs, because a dry run creates no task. With DryRun set to true, or
// with a dry_run query option, it returns an error and sends no request. Use the
// [BaremetalServerService.Update] method for a dry run. [option.WithResponseBodyInto] captures only the
// final server JSON. [option.WithResponseInto] captures the action and then the Get responses, but task
// reads do not overwrite it.
func (r *BaremetalServerService) UpdateAndPoll(ctx context.Context, serverID string, params BaremetalServerUpdateParams, opts ...option.RequestOption) (*BaremetalServer, error) {
	allOpts := slices.Concat(r.Options, opts)
	precfg, err := preparePoll("UpdateAndPoll", "Update", params.DryRun, params, allOpts)
	if err != nil {
		return nil, err
	}
	svc := NewBaremetalServerService()

	// exclude WithResponseBodyInto, because Update must decode the task list
	resource, err := svc.Update(ctx, serverID, params, requestconfig.ExcludeResponseBodyInto(allOpts...)...)
	if err != nil {
		return nil, err
	}
	if _, err := waitForTasks(ctx, resource.Tasks, "task", precfg, allOpts); err != nil {
		return nil, err
	}

	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	return getFinalResource(ctx, allOpts, func(ctx context.Context, opts ...option.RequestOption) (*BaremetalServer, error) {
		return svc.Get(ctx, serverID, BaremetalServerGetParams{ProjectID: params.ProjectID}, opts...)
	})
}

// ActionAndPoll runs a power action on a bare metal server, waits for every task of the action and returns
// the server. It returns an error when the response has no task or when a task fails.
//
// All tasks share one configured polling timeout and respect the caller's deadline. After the tasks finish,
// ActionAndPoll calls Get once and returns its result or error. Each HTTP request uses the SDK retry policy,
// including [option.WithMaxRetries].
//
// ActionAndPoll does not support dry runs, because a dry run creates no task. With DryRun set to true, or
// with a dry_run query option, it returns an error and sends no request. Use the
// [BaremetalServerService.Action] method for a dry run. [option.WithResponseBodyInto] captures only the
// final server JSON. [option.WithResponseInto] captures the action and then the Get responses, but task
// reads do not overwrite it.
func (r *BaremetalServerService) ActionAndPoll(ctx context.Context, serverID string, params BaremetalServerActionParams, opts ...option.RequestOption) (*BaremetalServer, error) {
	allOpts := slices.Concat(r.Options, opts)
	precfg, err := preparePoll("ActionAndPoll", "Action", params.DryRun, params, allOpts)
	if err != nil {
		return nil, err
	}
	svc := NewBaremetalServerService()

	// exclude WithResponseBodyInto, because Action must decode the task list
	resource, err := svc.Action(ctx, serverID, params, requestconfig.ExcludeResponseBodyInto(allOpts...)...)
	if err != nil {
		return nil, err
	}
	if _, err := waitForTasks(ctx, resource.Tasks, "task", precfg, allOpts); err != nil {
		return nil, err
	}

	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	return getFinalResource(ctx, allOpts, func(ctx context.Context, opts ...option.RequestOption) (*BaremetalServer, error) {
		return svc.Get(ctx, serverID, BaremetalServerGetParams{ProjectID: params.ProjectID}, opts...)
	})
}

// DeleteAndPoll deletes a bare metal server and waits for every task of the delete. It returns an error when
// the response has no task or when a task fails. By default the delete keeps the ports and floating IPs of
// the server. Use the fields of [BaremetalServerDeleteParams] to delete them too.
// All tasks share one configured polling timeout and respect the caller's deadline. Each HTTP request uses
// the SDK retry policy, including [option.WithMaxRetries].
//
// DeleteAndPoll does not support dry runs, because a dry run creates no task. With DryRun set to true, or
// with a dry_run query option, it returns an error and sends no request. Use the
// [BaremetalServerService.Delete] method for a dry run. [option.WithResponseInto] captures the delete
// response; [option.WithResponseBodyInto] is unused.
func (r *BaremetalServerService) DeleteAndPoll(ctx context.Context, serverID string, params BaremetalServerDeleteParams, opts ...option.RequestOption) error {
	allOpts := slices.Concat(r.Options, opts)
	precfg, err := preparePoll("DeleteAndPoll", "Delete", params.DryRun, params, allOpts)
	if err != nil {
		return err
	}
	svc := NewBaremetalServerService()

	// exclude WithResponseBodyInto, because Delete must decode the task list
	resource, err := svc.Delete(ctx, serverID, params, requestconfig.ExcludeResponseBodyInto(allOpts...)...)
	if err != nil {
		return err
	}
	_, err = waitForTasks(ctx, resource.Tasks, "task", precfg, allOpts)
	return err
}

// RebuildAndPoll rebuild bare metal server and poll for the completion of the first task.  Use the [TaskService.Poll]
// method if you need to poll for all tasks.
//
// The rebuild route is regional, so RebuildAndPoll reads the server with the regional List of
// [BaremetalRegionalServerService] and returns the regional model [BaremetalRegionalServer]. With
// option.WithResponseBodyInto, it writes the JSON of that regional server.
func (r *BaremetalServerService) RebuildAndPoll(ctx context.Context, serverID string, params BaremetalServerRebuildParams, opts ...option.RequestOption) (v *BaremetalRegionalServer, err error) {
	// Exclude WithResponseBodyInto for the action (Rebuild returns TaskIDList, must deserialize properly)
	actionOpts := requestconfig.ExcludeResponseBodyInto(opts...)
	resource, err := r.withoutResponseCapture().Rebuild(ctx, serverID, params, actionOpts...)
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
	_, err = newTaskService(requestconfig.ExcludeResponseBodyInto(r.Options...)...).Poll(ctx, taskID, pollOpts...)
	if err != nil {
		return
	}

	// Use the regional List to find the rebuilt server
	var listParams BaremetalRegionalServerListParams
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
	regional := NewBaremetalRegionalServerService(requestconfig.ExcludeResponseBodyInto(r.Options...)...)
	servers, err := regional.List(ctx, listParams, listOpts...)
	if err != nil {
		return
	}

	if len(servers.Results) == 0 {
		return nil, errors.New("server not found after rebuild")
	}

	if err := r.captureResponse(opts, servers.Results[0].RawJSON()); err != nil {
		return nil, err
	}

	return &servers.Results[0], nil
}

// createdBaremetalServerID returns the one server ID in created_resources.instances of the tasks
func createdBaremetalServerID(tasks []*Task) (string, error) {
	var ids []string
	for _, task := range tasks {
		for _, id := range task.CreatedResources.Instances {
			if id != "" && !slices.Contains(ids, id) {
				ids = append(ids, id)
			}
		}
	}
	switch len(ids) {
	case 0:
		return "", errors.New("expected a bare metal server ID in created_resources.instances of the tasks, got none")
	case 1:
		return ids[0], nil
	default:
		return "", fmt.Errorf("expected one bare metal server ID in created_resources.instances of the tasks, got %d: %s", len(ids), strings.Join(ids, ", "))
	}
}

// withoutResponseCapture returns a copy of the service without WithResponseBodyInto in its options. The
// generated methods put the service options before the options of the call, so a capture that the client
// sets would otherwise take the task list and leave the result empty
func (r *BaremetalServerService) withoutResponseCapture() *BaremetalServerService {
	svc := *r
	svc.Options = requestconfig.ExcludeResponseBodyInto(r.Options...)
	return &svc
}

// captureResponse writes raw to the WithResponseBodyInto destination of the service or of opts. An option of
// the call wins, as in a normal request
func (r *BaremetalServerService) captureResponse(opts []option.RequestOption, raw string) error {
	return requestconfig.WriteResponseBodyInto(slices.Concat(r.Options, opts), []byte(raw))
}
