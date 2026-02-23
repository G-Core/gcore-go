// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
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
	opts = slices.Concat(r.Options, opts)
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

// Acknowledge all tasks
func (r *TaskService) AcknowledgeAll(ctx context.Context, body TaskAcknowledgeAllParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "cloud/v1/tasks/acknowledge_all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Acknowledge one task
func (r *TaskService) AcknowledgeOne(ctx context.Context, taskID string, opts ...option.RequestOption) (res *Task, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/tasks/%s", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Poll for task status until it is finished, an error occurs, or the context is done. It uses a default polling interval
// of 1 second which can be overridden to values greater than 0 (otherwise the default value is used).
func (r *TaskService) Poll(ctx context.Context, taskID string, opts ...requestconfig.RequestOption) (*Task, error) {
	// extract polling interval from options, if not explicitly set, the default value is used
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	pollingInterval := time.Duration(precfg.CloudPollingIntervalSeconds) * time.Second
	// ensure the polling interval is at least 1 second
	if pollingInterval < time.Second {
		pollingInterval = time.Second
	}

	// set up polling timeout if configured, otherwise use the provided context
	pollingCtx := ctx
	var cancel context.CancelFunc
	if precfg.CloudPollingTimeoutSeconds > 0 {
		pollingTimeout := time.Duration(precfg.CloudPollingTimeoutSeconds) * time.Second
		pollingCtx, cancel = context.WithTimeout(ctx, pollingTimeout)
		defer cancel()
	}

	// poll the task status until it is finished or an error occurs
	for {
		task, err := r.Get(pollingCtx, taskID)
		if err != nil {
			return nil, fmt.Errorf("failed to get task status: %w", err)
		}

		if task.State == TaskStateFinished {
			return task, nil
		}

		if task.State == TaskStateError {
			return nil, fmt.Errorf("task %s failed with error: %s", taskID, task.Error)
		}

		// check if the context is done before sleeping
		select {
		// handles both timeout and cancellation
		case <-pollingCtx.Done():
			return nil, pollingCtx.Err()
		case <-time.After(pollingInterval):
		}
	}
}

type Task struct {
	// The task ID
	ID string `json:"id,required"`
	// Created timestamp
	CreatedOn string `json:"created_on,required"`
	// The task state
	//
	// Any of "ERROR", "FINISHED", "NEW", "RUNNING".
	State TaskState `json:"state,required"`
	// The task type
	TaskType string `json:"task_type,required"`
	// The user ID that initiated the task
	UserID int64 `json:"user_id,required"`
	// If task was acknowledged, this field stores acknowledge timestamp
	AcknowledgedAt string `json:"acknowledged_at,nullable"`
	// If task was acknowledged, this field stores `user_id` of the person
	AcknowledgedBy int64 `json:"acknowledged_by,nullable"`
	// The client ID
	ClientID int64 `json:"client_id,nullable"`
	// If the task creates resources, this field will contain their IDs
	CreatedResources TaskCreatedResources `json:"created_resources,nullable"`
	// Task parameters
	Data any `json:"data"`
	// Task detailed state that is more specific to task type
	//
	// Any of "CLUSTER_CLEAN_UP", "CLUSTER_REBUILD", "CLUSTER_RESIZE",
	// "CLUSTER_RESUME", "CLUSTER_SERVER_REBUILD", "CLUSTER_SUSPEND", "ERROR",
	// "FINISHED", "IPU_SERVERS", "NETWORK", "POPLAR_SERVERS", "POST_DEPLOY_SETUP",
	// "VIPU_CONTROLLER".
	DetailedState TaskDetailedState `json:"detailed_state,nullable"`
	// The error value
	Error string `json:"error,nullable"`
	// Finished timestamp
	FinishedOn string `json:"finished_on,nullable"`
	// Job ID
	JobID string `json:"job_id,nullable"`
	// Lifecycle policy ID
	LifecyclePolicyID int64 `json:"lifecycle_policy_id,nullable"`
	// The project ID
	ProjectID int64 `json:"project_id,nullable"`
	// The region ID
	RegionID int64 `json:"region_id,nullable"`
	// The request ID
	RequestID string `json:"request_id,nullable"`
	// Schedule ID
	ScheduleID string `json:"schedule_id,nullable"`
	// Last updated timestamp
	UpdatedOn string `json:"updated_on,nullable"`
	// Client, specified in user's JWT
	UserClientID int64 `json:"user_client_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedOn         respjson.Field
		State             respjson.Field
		TaskType          respjson.Field
		UserID            respjson.Field
		AcknowledgedAt    respjson.Field
		AcknowledgedBy    respjson.Field
		ClientID          respjson.Field
		CreatedResources  respjson.Field
		Data              respjson.Field
		DetailedState     respjson.Field
		Error             respjson.Field
		FinishedOn        respjson.Field
		JobID             respjson.Field
		LifecyclePolicyID respjson.Field
		ProjectID         respjson.Field
		RegionID          respjson.Field
		RequestID         respjson.Field
		ScheduleID        respjson.Field
		UpdatedOn         respjson.Field
		UserClientID      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Task) RawJSON() string { return r.JSON.raw }
func (r *Task) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The task state
type TaskState string

const (
	TaskStateError    TaskState = "ERROR"
	TaskStateFinished TaskState = "FINISHED"
	TaskStateNew      TaskState = "NEW"
	TaskStateRunning  TaskState = "RUNNING"
)

// If the task creates resources, this field will contain their IDs
type TaskCreatedResources struct {
	// Deprecated, use 'clusters' instead. IDs of created GPU clusters
	//
	// Deprecated: deprecated
	AIClusters []string `json:"ai_clusters"`
	// IDs of created API keys
	APIKeys []string `json:"api_keys"`
	// IDs of created CaaS containers
	CaasContainers []string `json:"caas_containers"`
	// IDs of created GPU clusters
	Clusters []string `json:"clusters"`
	// IDs of created ddos protection profiles
	DDOSProfiles []int64 `json:"ddos_profiles"`
	// IDs of created FaaS functions
	FaasFunctions []string `json:"faas_functions"`
	// IDs of created FaaS namespaces
	FaasNamespaces []string `json:"faas_namespaces"`
	// IDs of created file shares
	FileShares []string `json:"file_shares"`
	// IDs of created floating IPs
	Floatingips []string `json:"floatingips"`
	// IDs of created health monitors
	Healthmonitors []string `json:"healthmonitors"`
	// IDs of created images
	Images []string `json:"images"`
	// IDs of created inference instances
	InferenceInstances []string `json:"inference_instances"`
	// IDs of created instances
	Instances []string `json:"instances"`
	// IDs of created Kubernetes clusters
	K8SClusters []string `json:"k8s_clusters"`
	// IDs of created Kubernetes pools
	K8SPools []string `json:"k8s_pools"`
	// IDs of created L7 policies
	L7polices []string `json:"l7polices"`
	// IDs of created L7 rules
	L7rules []string `json:"l7rules"`
	// IDs of created LaaS topics
	LaasTopic []string `json:"laas_topic"`
	// IDs of created load balancer listeners
	Listeners []string `json:"listeners"`
	// IDs of created load balancers
	Loadbalancers []string `json:"loadbalancers"`
	// IDs of created pool members
	Members []string `json:"members"`
	// IDs of created networks
	Networks []string `json:"networks"`
	// IDs of created load balancer pools
	Pools []string `json:"pools"`
	// IDs of created ports
	Ports []string `json:"ports"`
	// IDs of created postgres clusters
	PostgreSQLClusters []string `json:"postgresql_clusters"`
	// IDs of created projects
	Projects []int64 `json:"projects"`
	// IDs of created registry registries
	RegistryRegistries []string `json:"registry_registries"`
	// IDs of created registry users
	RegistryUsers []string `json:"registry_users"`
	// IDs of created routers
	Routers []string `json:"routers"`
	// IDs of created secrets
	Secrets []string `json:"secrets"`
	// IDs of created security group rules
	SecurityGroupRules []string `json:"security_group_rules"`
	// IDs of created security groups
	SecurityGroups []string `json:"security_groups"`
	// IDs of created server groups
	Servergroups []string `json:"servergroups"`
	// IDs of created volume snapshots
	Snapshots []string `json:"snapshots"`
	// IDs of created subnets
	Subnets []string `json:"subnets"`
	// IDs of created volumes
	Volumes []string `json:"volumes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIClusters         respjson.Field
		APIKeys            respjson.Field
		CaasContainers     respjson.Field
		Clusters           respjson.Field
		DDOSProfiles       respjson.Field
		FaasFunctions      respjson.Field
		FaasNamespaces     respjson.Field
		FileShares         respjson.Field
		Floatingips        respjson.Field
		Healthmonitors     respjson.Field
		Images             respjson.Field
		InferenceInstances respjson.Field
		Instances          respjson.Field
		K8SClusters        respjson.Field
		K8SPools           respjson.Field
		L7polices          respjson.Field
		L7rules            respjson.Field
		LaasTopic          respjson.Field
		Listeners          respjson.Field
		Loadbalancers      respjson.Field
		Members            respjson.Field
		Networks           respjson.Field
		Pools              respjson.Field
		Ports              respjson.Field
		PostgreSQLClusters respjson.Field
		Projects           respjson.Field
		RegistryRegistries respjson.Field
		RegistryUsers      respjson.Field
		Routers            respjson.Field
		Secrets            respjson.Field
		SecurityGroupRules respjson.Field
		SecurityGroups     respjson.Field
		Servergroups       respjson.Field
		Snapshots          respjson.Field
		Subnets            respjson.Field
		Volumes            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskCreatedResources) RawJSON() string { return r.JSON.raw }
func (r *TaskCreatedResources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Task detailed state that is more specific to task type
type TaskDetailedState string

const (
	TaskDetailedStateClusterCleanUp       TaskDetailedState = "CLUSTER_CLEAN_UP"
	TaskDetailedStateClusterRebuild       TaskDetailedState = "CLUSTER_REBUILD"
	TaskDetailedStateClusterResize        TaskDetailedState = "CLUSTER_RESIZE"
	TaskDetailedStateClusterResume        TaskDetailedState = "CLUSTER_RESUME"
	TaskDetailedStateClusterServerRebuild TaskDetailedState = "CLUSTER_SERVER_REBUILD"
	TaskDetailedStateClusterSuspend       TaskDetailedState = "CLUSTER_SUSPEND"
	TaskDetailedStateError                TaskDetailedState = "ERROR"
	TaskDetailedStateFinished             TaskDetailedState = "FINISHED"
	TaskDetailedStateIpuServers           TaskDetailedState = "IPU_SERVERS"
	TaskDetailedStateNetwork              TaskDetailedState = "NETWORK"
	TaskDetailedStatePoplarServers        TaskDetailedState = "POPLAR_SERVERS"
	TaskDetailedStatePostDeploySetup      TaskDetailedState = "POST_DEPLOY_SETUP"
	TaskDetailedStateVipuController       TaskDetailedState = "VIPU_CONTROLLER"
)

type TaskListParams struct {
	// ISO formatted datetime string. Filter the tasks by creation date greater than or
	// equal to `from_timestamp`
	FromTimestamp param.Opt[time.Time] `query:"from_timestamp,omitzero" format:"date-time" json:"-"`
	// Filter the tasks by their acknowledgement status
	IsAcknowledged param.Opt[bool] `query:"is_acknowledged,omitzero" json:"-"`
	// Limit the number of returned tasks. Falls back to default of 10 if not
	// specified. Limited by max limit value of 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter the tasks by their type one of ['activate_ddos_profile',
	// 'attach_bm_to_reserved_fixed_ip', 'attach_vm_interface',
	// 'attach_vm_to_reserved_fixed_ip', 'attach_volume', 'create_ai_cluster_gpu',
	// 'create_bm', 'create_caas_container', 'create_dbaas_postgres_cluster',
	// 'create_ddos_profile', 'create_faas_function', 'create_faas_namespace',
	// 'create_fip', 'create_gpu_virtual_cluster', 'create_image',
	// 'create_inference_application', 'create_inference_instance',
	// 'create_k8s_cluster_pool_v2', 'create_k8s_cluster_v2', 'create_l7policy',
	// 'create_l7rule', 'create_lblistener', 'create_lbmember', 'create_lbpool',
	// 'create_lbpool_health_monitor', 'create_loadbalancer', 'create_network',
	// 'create_reserved_fixed_ip', 'create_router', 'create_secret',
	// 'create_security_group', 'create_security_group_rule', 'create_servergroup',
	// 'create_sfs', 'create_snapshot', 'create_subnet', 'create_vm', 'create_volume',
	// 'deactivate_ddos_profile', 'delete_ai_cluster_gpu', 'delete_caas_container',
	// 'delete_dbaas_postgres_cluster', 'delete_ddos_profile', 'delete_faas_function',
	// 'delete_faas_namespace', 'delete_fip', 'delete_gpu_virtual_cluster',
	// 'delete_gpu_virtual_server', 'delete_image', 'delete_inference_application',
	// 'delete_inference_instance', 'delete_k8s_cluster_pool_v2',
	// 'delete_k8s_cluster_v2', 'delete_l7policy', 'delete_l7rule',
	// 'delete_lblistener', 'delete_lbmember', 'delete_lbmetadata', 'delete_lbpool',
	// 'delete_loadbalancer', 'delete_network', 'delete_project',
	// 'delete_reserved_fixed_ip', 'delete_router', 'delete_secret',
	// 'delete_security_group_rule', 'delete_servergroup', 'delete_sfs',
	// 'delete_snapshot', 'delete_subnet', 'delete_vm', 'delete_volume',
	// 'detach_vm_interface', 'detach_volume', 'download_image',
	// 'downscale_ai_cluster_gpu', 'downscale_gpu_virtual_cluster', 'extend_sfs',
	// 'extend_volume', 'failover_loadbalancer', 'hard_reboot_gpu_baremetal_server',
	// 'hard_reboot_gpu_virtual_cluster', 'hard_reboot_gpu_virtual_server',
	// 'hard_reboot_vm', 'patch_caas_container', 'patch_dbaas_postgres_cluster',
	// 'patch_faas_function', 'patch_faas_namespace', 'patch_lblistener',
	// 'patch_lbpool', 'put_into_server_group', 'put_l7rule', 'rebuild_bm',
	// 'rebuild_gpu_baremetal_cluster', 'rebuild_gpu_baremetal_node',
	// 'rebuild_gpu_baremetal_server', 'remove_from_server_group',
	// 'replace_lbmetadata', 'resize_k8s_cluster_v2', 'resize_loadbalancer',
	// 'resize_vm', 'resume_vm', 'revert_volume', 'soft_reboot_gpu_baremetal_server',
	// 'soft_reboot_gpu_virtual_cluster', 'soft_reboot_gpu_virtual_server',
	// 'soft_reboot_vm', 'start_gpu_baremetal_server', 'start_gpu_virtual_cluster',
	// 'start_gpu_virtual_server', 'start_vm', 'stop_gpu_baremetal_server',
	// 'stop_gpu_virtual_cluster', 'stop_gpu_virtual_server', 'stop_vm', 'suspend_vm',
	// 'sync_private_flavors', 'update_ddos_profile', 'update_floating_ip',
	// 'update_inference_application', 'update_inference_instance',
	// 'update_k8s_cluster_v2', 'update_l7policy', 'update_lbmetadata',
	// 'update_loadbalancer', 'update_port_allowed_address_pairs', 'update_router',
	// 'update_security_group', 'update_sfs', 'update_tags_gpu_virtual_cluster',
	// 'upgrade_k8s_cluster_v2', 'upscale_ai_cluster_gpu',
	// 'upscale_gpu_virtual_cluster']
	TaskType param.Opt[string] `query:"task_type,omitzero" json:"-"`
	// ISO formatted datetime string. Filter the tasks by creation date less than or
	// equal to `to_timestamp`
	ToTimestamp param.Opt[time.Time] `query:"to_timestamp,omitzero" format:"date-time" json:"-"`
	// Sorting by creation date. Oldest first, or most recent first
	//
	// Any of "asc", "desc".
	OrderBy TaskListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// The project ID to filter the tasks by project. Supports multiple values of kind
	// key=value1&key=value2
	ProjectID []int64 `query:"project_id,omitzero" json:"-"`
	// The region ID to filter the tasks by region. Supports multiple values of kind
	// key=value1&key=value2
	RegionID []int64 `query:"region_id,omitzero" json:"-"`
	// (DEPRECATED Use 'order_by' instead) Sorting by creation date. Oldest first, or
	// most recent first
	//
	// Any of "asc", "desc".
	Sorting TaskListParamsSorting `query:"sorting,omitzero" json:"-"`
	// Filter the tasks by state. Supports multiple values of kind
	// key=value1&key=value2
	//
	// Any of "ERROR", "FINISHED", "NEW", "RUNNING".
	State []string `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaskListParams]'s query parameters as `url.Values`.
func (r TaskListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sorting by creation date. Oldest first, or most recent first
type TaskListParamsOrderBy string

const (
	TaskListParamsOrderByAsc  TaskListParamsOrderBy = "asc"
	TaskListParamsOrderByDesc TaskListParamsOrderBy = "desc"
)

// (DEPRECATED Use 'order_by' instead) Sorting by creation date. Oldest first, or
// most recent first
type TaskListParamsSorting string

const (
	TaskListParamsSortingAsc  TaskListParamsSorting = "asc"
	TaskListParamsSortingDesc TaskListParamsSorting = "desc"
)

type TaskAcknowledgeAllParams struct {
	// Project ID
	ProjectID param.Opt[int64] `query:"project_id,omitzero" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `query:"region_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaskAcknowledgeAllParams]'s query parameters as
// `url.Values`.
func (r TaskAcknowledgeAllParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
