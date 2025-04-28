// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// TaskService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskService] method instead.
type TaskService struct {
	Options []option.RequestOption
}

// NewTaskService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTaskService(opts ...option.RequestOption) (r TaskService) {
	r = TaskService{}
	r.Options = opts
	return
}

// List tasks
func (r *TaskService) List(ctx context.Context, query TaskListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Task], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cloud/v1/tasks"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List tasks
func (r *TaskService) ListAutoPaging(ctx context.Context, query TaskListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Task] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Acknowledge all client tasks in project or region
func (r *TaskService) AcknowledgeAll(ctx context.Context, body TaskAcknowledgeAllParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cloud/v1/tasks/acknowledge_all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Acknowledge one task on project scope
func (r *TaskService) AcknowledgeOne(ctx context.Context, taskID string, opts ...option.RequestOption) (res *Task, err error) {
	opts = append(r.Options[:], opts...)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/tasks/%s/acknowledge", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get task
func (r *TaskService) Get(ctx context.Context, taskID string, opts ...option.RequestOption) (res *Task, err error) {
	opts = append(r.Options[:], opts...)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/tasks/%s", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// '#/components/schemas/TaskSerializer' "$.components.schemas.TaskSerializer"
// Poll for task status until it is finished, an error occurs or the context is done. It uses a default polling interval
// of 1 second which can be overridden.
func (r *TaskService) Poll(ctx context.Context, taskID string, opts ...requestconfig.RequestOption) (*Task, error) {
	// extract polling interval from options, if not explicitly set, the default value is used
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	pollingInterval := time.Duration(precfg.PollingIntervalMs) * time.Millisecond

	// poll the task status until it is finished or an error occurs
	for {
		task, err := r.Get(ctx, taskID)
		if err != nil {
			return nil, fmt.Errorf("failed to get task status: %w", err)
		}

		if task.State == "FINISHED" {
			return task, nil
		}

		if task.State == "ERROR" {
			return nil, fmt.Errorf("task %s failed with error: %s", taskID, task.Error)
		}

		time.Sleep(pollingInterval)

		// check if the context is done, if so return the error
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}
}

type Task struct {
	// '#/components/schemas/TaskSerializer/properties/id'
	// "$.components.schemas.TaskSerializer.properties.id"
	ID string `json:"id,required"`
	// '#/components/schemas/TaskSerializer/properties/created_on/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.created_on.anyOf[0]"
	CreatedOn string `json:"created_on,required"`
	// '#/components/schemas/TaskSerializer/properties/state'
	// "$.components.schemas.TaskSerializer.properties.state"
	//
	// Any of "ERROR", "FINISHED", "NEW", "RUNNING".
	State TaskState `json:"state,required"`
	// '#/components/schemas/TaskSerializer/properties/task_type'
	// "$.components.schemas.TaskSerializer.properties.task_type"
	TaskType string `json:"task_type,required"`
	// '#/components/schemas/TaskSerializer/properties/user_id'
	// "$.components.schemas.TaskSerializer.properties.user_id"
	UserID int64 `json:"user_id,required"`
	// '#/components/schemas/TaskSerializer/properties/acknowledged_at/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.acknowledged_at.anyOf[0]"
	AcknowledgedAt string `json:"acknowledged_at,nullable"`
	// '#/components/schemas/TaskSerializer/properties/acknowledged_by/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.acknowledged_by.anyOf[0]"
	AcknowledgedBy int64 `json:"acknowledged_by,nullable"`
	// '#/components/schemas/TaskSerializer/properties/client_id/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.client_id.anyOf[0]"
	ClientID int64 `json:"client_id,nullable"`
	// '#/components/schemas/TaskSerializer/properties/created_resources/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.created_resources.anyOf[0]"
	CreatedResources TaskCreatedResources `json:"created_resources,nullable"`
	// '#/components/schemas/TaskSerializer/properties/data/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.data.anyOf[0]"
	Data any `json:"data"`
	// '#/components/schemas/TaskSerializer/properties/detailed_state/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.detailed_state.anyOf[0]"
	//
	// Any of "CLUSTER_CLEAN_UP", "CLUSTER_RESIZE", "CLUSTER_RESUME",
	// "CLUSTER_SUSPEND", "ERROR", "FINISHED", "IPU_SERVERS", "NETWORK",
	// "POPLAR_SERVERS", "POST_DEPLOY_SETUP", "VIPU_CONTROLLER".
	DetailedState TaskDetailedState `json:"detailed_state,nullable"`
	// '#/components/schemas/TaskSerializer/properties/error/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.error.anyOf[0]"
	Error string `json:"error,nullable"`
	// '#/components/schemas/TaskSerializer/properties/finished_on/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.finished_on.anyOf[0]"
	FinishedOn string `json:"finished_on,nullable"`
	// '#/components/schemas/TaskSerializer/properties/job_id/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.job_id.anyOf[0]"
	JobID string `json:"job_id,nullable"`
	// '#/components/schemas/TaskSerializer/properties/lifecycle_policy_id/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.lifecycle_policy_id.anyOf[0]"
	LifecyclePolicyID int64 `json:"lifecycle_policy_id,nullable"`
	// '#/components/schemas/TaskSerializer/properties/project_id/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.project_id.anyOf[0]"
	ProjectID int64 `json:"project_id,nullable"`
	// '#/components/schemas/TaskSerializer/properties/region_id/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.region_id.anyOf[0]"
	RegionID int64 `json:"region_id,nullable"`
	// '#/components/schemas/TaskSerializer/properties/request_id/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.request_id.anyOf[0]"
	RequestID string `json:"request_id,nullable"`
	// '#/components/schemas/TaskSerializer/properties/schedule_id/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.schedule_id.anyOf[0]"
	ScheduleID string `json:"schedule_id,nullable"`
	// '#/components/schemas/TaskSerializer/properties/updated_on/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.updated_on.anyOf[0]"
	UpdatedOn string `json:"updated_on,nullable"`
	// '#/components/schemas/TaskSerializer/properties/user_client_id/anyOf/0'
	// "$.components.schemas.TaskSerializer.properties.user_client_id.anyOf[0]"
	UserClientID int64 `json:"user_client_id,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                resp.Field
		CreatedOn         resp.Field
		State             resp.Field
		TaskType          resp.Field
		UserID            resp.Field
		AcknowledgedAt    resp.Field
		AcknowledgedBy    resp.Field
		ClientID          resp.Field
		CreatedResources  resp.Field
		Data              resp.Field
		DetailedState     resp.Field
		Error             resp.Field
		FinishedOn        resp.Field
		JobID             resp.Field
		LifecyclePolicyID resp.Field
		ProjectID         resp.Field
		RegionID          resp.Field
		RequestID         resp.Field
		ScheduleID        resp.Field
		UpdatedOn         resp.Field
		UserClientID      resp.Field
		ExtraFields       map[string]resp.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Task) RawJSON() string { return r.JSON.raw }
func (r *Task) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/TaskSerializer/properties/state'
// "$.components.schemas.TaskSerializer.properties.state"
type TaskState string

const (
	TaskStateError    TaskState = "ERROR"
	TaskStateFinished TaskState = "FINISHED"
	TaskStateNew      TaskState = "NEW"
	TaskStateRunning  TaskState = "RUNNING"
)

// '#/components/schemas/TaskSerializer/properties/created_resources/anyOf/0'
// "$.components.schemas.TaskSerializer.properties.created_resources.anyOf[0]"
type TaskCreatedResources struct {
	// '#/components/schemas/CreatedResources/properties/ai_clusters'
	// "$.components.schemas.CreatedResources.properties.ai_clusters"
	AIClusters []string `json:"ai_clusters"`
	// '#/components/schemas/CreatedResources/properties/api_keys'
	// "$.components.schemas.CreatedResources.properties.api_keys"
	APIKeys []string `json:"api_keys"`
	// '#/components/schemas/CreatedResources/properties/caas_containers'
	// "$.components.schemas.CreatedResources.properties.caas_containers"
	CaasContainers []string `json:"caas_containers"`
	// '#/components/schemas/CreatedResources/properties/ddos_profiles'
	// "$.components.schemas.CreatedResources.properties.ddos_profiles"
	DDOSProfiles []int64 `json:"ddos_profiles"`
	// '#/components/schemas/CreatedResources/properties/faas_functions'
	// "$.components.schemas.CreatedResources.properties.faas_functions"
	FaasFunctions []string `json:"faas_functions"`
	// '#/components/schemas/CreatedResources/properties/faas_namespaces'
	// "$.components.schemas.CreatedResources.properties.faas_namespaces"
	FaasNamespaces []string `json:"faas_namespaces"`
	// '#/components/schemas/CreatedResources/properties/file_shares'
	// "$.components.schemas.CreatedResources.properties.file_shares"
	FileShares []string `json:"file_shares"`
	// '#/components/schemas/CreatedResources/properties/floatingips'
	// "$.components.schemas.CreatedResources.properties.floatingips"
	Floatingips []string `json:"floatingips"`
	// '#/components/schemas/CreatedResources/properties/healthmonitors'
	// "$.components.schemas.CreatedResources.properties.healthmonitors"
	Healthmonitors []string `json:"healthmonitors"`
	// '#/components/schemas/CreatedResources/properties/heat'
	// "$.components.schemas.CreatedResources.properties.heat"
	Heat []string `json:"heat"`
	// '#/components/schemas/CreatedResources/properties/images'
	// "$.components.schemas.CreatedResources.properties.images"
	Images []string `json:"images"`
	// '#/components/schemas/CreatedResources/properties/inference_instances'
	// "$.components.schemas.CreatedResources.properties.inference_instances"
	InferenceInstances []string `json:"inference_instances"`
	// '#/components/schemas/CreatedResources/properties/instances'
	// "$.components.schemas.CreatedResources.properties.instances"
	Instances []string `json:"instances"`
	// '#/components/schemas/CreatedResources/properties/k8s_clusters'
	// "$.components.schemas.CreatedResources.properties.k8s_clusters"
	K8sClusters []string `json:"k8s_clusters"`
	// '#/components/schemas/CreatedResources/properties/k8s_pools'
	// "$.components.schemas.CreatedResources.properties.k8s_pools"
	K8sPools []string `json:"k8s_pools"`
	// '#/components/schemas/CreatedResources/properties/l7polices'
	// "$.components.schemas.CreatedResources.properties.l7polices"
	L7polices []string `json:"l7polices"`
	// '#/components/schemas/CreatedResources/properties/l7rules'
	// "$.components.schemas.CreatedResources.properties.l7rules"
	L7rules []string `json:"l7rules"`
	// '#/components/schemas/CreatedResources/properties/laas_topic'
	// "$.components.schemas.CreatedResources.properties.laas_topic"
	LaasTopic []string `json:"laas_topic"`
	// '#/components/schemas/CreatedResources/properties/listeners'
	// "$.components.schemas.CreatedResources.properties.listeners"
	Listeners []string `json:"listeners"`
	// '#/components/schemas/CreatedResources/properties/loadbalancers'
	// "$.components.schemas.CreatedResources.properties.loadbalancers"
	Loadbalancers []string `json:"loadbalancers"`
	// '#/components/schemas/CreatedResources/properties/members'
	// "$.components.schemas.CreatedResources.properties.members"
	Members []string `json:"members"`
	// '#/components/schemas/CreatedResources/properties/networks'
	// "$.components.schemas.CreatedResources.properties.networks"
	Networks []string `json:"networks"`
	// '#/components/schemas/CreatedResources/properties/pools'
	// "$.components.schemas.CreatedResources.properties.pools"
	Pools []string `json:"pools"`
	// '#/components/schemas/CreatedResources/properties/ports'
	// "$.components.schemas.CreatedResources.properties.ports"
	Ports []string `json:"ports"`
	// '#/components/schemas/CreatedResources/properties/postgresql_clusters'
	// "$.components.schemas.CreatedResources.properties.postgresql_clusters"
	PostgreSQLClusters []string `json:"postgresql_clusters"`
	// '#/components/schemas/CreatedResources/properties/projects'
	// "$.components.schemas.CreatedResources.properties.projects"
	Projects []int64 `json:"projects"`
	// '#/components/schemas/CreatedResources/properties/registry_registries'
	// "$.components.schemas.CreatedResources.properties.registry_registries"
	RegistryRegistries []string `json:"registry_registries"`
	// '#/components/schemas/CreatedResources/properties/registry_users'
	// "$.components.schemas.CreatedResources.properties.registry_users"
	RegistryUsers []string `json:"registry_users"`
	// '#/components/schemas/CreatedResources/properties/routers'
	// "$.components.schemas.CreatedResources.properties.routers"
	Routers []string `json:"routers"`
	// '#/components/schemas/CreatedResources/properties/secrets'
	// "$.components.schemas.CreatedResources.properties.secrets"
	Secrets []string `json:"secrets"`
	// '#/components/schemas/CreatedResources/properties/servergroups'
	// "$.components.schemas.CreatedResources.properties.servergroups"
	Servergroups []string `json:"servergroups"`
	// '#/components/schemas/CreatedResources/properties/snapshots'
	// "$.components.schemas.CreatedResources.properties.snapshots"
	Snapshots []string `json:"snapshots"`
	// '#/components/schemas/CreatedResources/properties/subnets'
	// "$.components.schemas.CreatedResources.properties.subnets"
	Subnets []string `json:"subnets"`
	// '#/components/schemas/CreatedResources/properties/volumes'
	// "$.components.schemas.CreatedResources.properties.volumes"
	Volumes []string `json:"volumes"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AIClusters         resp.Field
		APIKeys            resp.Field
		CaasContainers     resp.Field
		DDOSProfiles       resp.Field
		FaasFunctions      resp.Field
		FaasNamespaces     resp.Field
		FileShares         resp.Field
		Floatingips        resp.Field
		Healthmonitors     resp.Field
		Heat               resp.Field
		Images             resp.Field
		InferenceInstances resp.Field
		Instances          resp.Field
		K8sClusters        resp.Field
		K8sPools           resp.Field
		L7polices          resp.Field
		L7rules            resp.Field
		LaasTopic          resp.Field
		Listeners          resp.Field
		Loadbalancers      resp.Field
		Members            resp.Field
		Networks           resp.Field
		Pools              resp.Field
		Ports              resp.Field
		PostgreSQLClusters resp.Field
		Projects           resp.Field
		RegistryRegistries resp.Field
		RegistryUsers      resp.Field
		Routers            resp.Field
		Secrets            resp.Field
		Servergroups       resp.Field
		Snapshots          resp.Field
		Subnets            resp.Field
		Volumes            resp.Field
		ExtraFields        map[string]resp.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskCreatedResources) RawJSON() string { return r.JSON.raw }
func (r *TaskCreatedResources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/TaskSerializer/properties/detailed_state/anyOf/0'
// "$.components.schemas.TaskSerializer.properties.detailed_state.anyOf[0]"
type TaskDetailedState string

const (
	TaskDetailedStateClusterCleanUp  TaskDetailedState = "CLUSTER_CLEAN_UP"
	TaskDetailedStateClusterResize   TaskDetailedState = "CLUSTER_RESIZE"
	TaskDetailedStateClusterResume   TaskDetailedState = "CLUSTER_RESUME"
	TaskDetailedStateClusterSuspend  TaskDetailedState = "CLUSTER_SUSPEND"
	TaskDetailedStateError           TaskDetailedState = "ERROR"
	TaskDetailedStateFinished        TaskDetailedState = "FINISHED"
	TaskDetailedStateIpuServers      TaskDetailedState = "IPU_SERVERS"
	TaskDetailedStateNetwork         TaskDetailedState = "NETWORK"
	TaskDetailedStatePoplarServers   TaskDetailedState = "POPLAR_SERVERS"
	TaskDetailedStatePostDeploySetup TaskDetailedState = "POST_DEPLOY_SETUP"
	TaskDetailedStateVipuController  TaskDetailedState = "VIPU_CONTROLLER"
)

type TaskListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ftasks/get/parameters/0/schema/anyOf/0'
	// "$.paths['/cloud/v1/tasks'].get.parameters[0].schema.anyOf[0]"
	FromTimestamp param.Opt[time.Time] `query:"from_timestamp,omitzero" format:"date-time" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ftasks/get/parameters/1/schema/anyOf/0'
	// "$.paths['/cloud/v1/tasks'].get.parameters[1].schema.anyOf[0]"
	IsAcknowledged param.Opt[bool] `query:"is_acknowledged,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ftasks/get/parameters/8/schema/anyOf/0'
	// "$.paths['/cloud/v1/tasks'].get.parameters[8].schema.anyOf[0]"
	TaskType param.Opt[string] `query:"task_type,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ftasks/get/parameters/9/schema/anyOf/0'
	// "$.paths['/cloud/v1/tasks'].get.parameters[9].schema.anyOf[0]"
	ToTimestamp param.Opt[time.Time] `query:"to_timestamp,omitzero" format:"date-time" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ftasks/get/parameters/2'
	// "$.paths['/cloud/v1/tasks'].get.parameters[2]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ftasks/get/parameters/3'
	// "$.paths['/cloud/v1/tasks'].get.parameters[3]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ftasks/get/parameters/4/schema/anyOf/0'
	// "$.paths['/cloud/v1/tasks'].get.parameters[4].schema.anyOf[0]"
	ProjectID []int64 `query:"project_id,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ftasks/get/parameters/5/schema/anyOf/0'
	// "$.paths['/cloud/v1/tasks'].get.parameters[5].schema.anyOf[0]"
	RegionID []int64 `query:"region_id,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ftasks/get/parameters/6/schema/anyOf/0'
	// "$.paths['/cloud/v1/tasks'].get.parameters[6].schema.anyOf[0]"
	//
	// Any of "asc", "desc".
	Sorting TaskListParamsSorting `query:"sorting,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ftasks/get/parameters/7/schema/anyOf/0'
	// "$.paths['/cloud/v1/tasks'].get.parameters[7].schema.anyOf[0]"
	//
	// Any of "ERROR", "FINISHED", "NEW", "RUNNING".
	State []string `query:"state,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TaskListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [TaskListParams]'s query parameters as `url.Values`.
func (r TaskListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// '#/paths/%2Fcloud%2Fv1%2Ftasks/get/parameters/6/schema/anyOf/0'
// "$.paths['/cloud/v1/tasks'].get.parameters[6].schema.anyOf[0]"
type TaskListParamsSorting string

const (
	TaskListParamsSortingAsc  TaskListParamsSorting = "asc"
	TaskListParamsSortingDesc TaskListParamsSorting = "desc"
)

type TaskAcknowledgeAllParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ftasks%2Facknowledge_all/post/parameters/0/schema/anyOf/0'
	// "$.paths['/cloud/v1/tasks/acknowledge_all'].post.parameters[0].schema.anyOf[0]"
	ProjectID param.Opt[int64] `query:"project_id,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ftasks%2Facknowledge_all/post/parameters/1/schema/anyOf/0'
	// "$.paths['/cloud/v1/tasks/acknowledge_all'].post.parameters[1].schema.anyOf[0]"
	RegionID param.Opt[int64] `query:"region_id,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TaskAcknowledgeAllParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [TaskAcknowledgeAllParams]'s query parameters as
// `url.Values`.
func (r TaskAcknowledgeAllParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
