// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1TaskService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1TaskService] method instead.
type CloudV1TaskService struct {
	Options []option.RequestOption
}

// NewCloudV1TaskService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1TaskService(opts ...option.RequestOption) (r *CloudV1TaskService) {
	r = &CloudV1TaskService{}
	r.Options = opts
	return
}

// Get task
func (r *CloudV1TaskService) Get(ctx context.Context, taskID string, opts ...option.RequestOption) (res *Task, err error) {
	opts = append(r.Options[:], opts...)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/tasks/%s", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List tasks
func (r *CloudV1TaskService) List(ctx context.Context, query CloudV1TaskListParams, opts ...option.RequestOption) (res *TaskCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v1/tasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deprecated
//
// Use /v1/tasks/acknowledge_all instead
func (r *CloudV1TaskService) Acknowledge(ctx context.Context, params CloudV1TaskAcknowledgeParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cloud/v1/tasks/acknowledge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Acknowledge all client tasks in project or region
func (r *CloudV1TaskService) AcknowledgeAll(ctx context.Context, body CloudV1TaskAcknowledgeAllParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "cloud/v1/tasks/acknowledge_all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Acknowledge one task on project scope
func (r *CloudV1TaskService) AcknowledgeTask(ctx context.Context, taskID string, opts ...option.RequestOption) (res *Task, err error) {
	opts = append(r.Options[:], opts...)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/tasks/%s/acknowledge", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// List active tasks
func (r *CloudV1TaskService) ListActive(ctx context.Context, projectID int64, regionID int64, query CloudV1TaskListActiveParams, opts ...option.RequestOption) (res *TaskCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/tasks/%v/%v/active", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Task struct {
	// The task ID
	ID string `json:"id,required"`
	// Created timestamp
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The task state
	State string `json:"state,required"`
	// The task type
	TaskType string `json:"task_type,required"`
	// The user ID that initiated the task
	UserID int64 `json:"user_id,required"`
	// If task was acknowledged, this field stores acknowledge timestamp
	AcknowledgedAt time.Time `json:"acknowledged_at,nullable" format:"date-time"`
	// If task was acknowledged, this field stores user_id of the person
	AcknowledgedBy int64 `json:"acknowledged_by,nullable"`
	// The client ID
	ClientID int64 `json:"client_id,nullable"`
	// If the task creates resources, this field will contain their IDs
	CreatedResources TaskCreatedResources `json:"created_resources,nullable"`
	// Task parameters
	Data string `json:"data,nullable"`
	// Task detailed state that is more specific to task type
	DetailedState TaskDetailedState `json:"detailed_state,nullable"`
	// The error value
	Error string `json:"error,nullable"`
	// Finished timestamp
	FinishedOn time.Time `json:"finished_on,nullable" format:"date-time"`
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
	UpdatedOn time.Time `json:"updated_on,nullable" format:"date-time"`
	// Client, specified in user's JWT
	UserClientID int64    `json:"user_client_id,nullable"`
	JSON         taskJSON `json:"-"`
}

// taskJSON contains the JSON metadata for the struct [Task]
type taskJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	State             apijson.Field
	TaskType          apijson.Field
	UserID            apijson.Field
	AcknowledgedAt    apijson.Field
	AcknowledgedBy    apijson.Field
	ClientID          apijson.Field
	CreatedResources  apijson.Field
	Data              apijson.Field
	DetailedState     apijson.Field
	Error             apijson.Field
	FinishedOn        apijson.Field
	JobID             apijson.Field
	LifecyclePolicyID apijson.Field
	ProjectID         apijson.Field
	RegionID          apijson.Field
	RequestID         apijson.Field
	ScheduleID        apijson.Field
	UpdatedOn         apijson.Field
	UserClientID      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Task) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskJSON) RawJSON() string {
	return r.raw
}

// If the task creates resources, this field will contain their IDs
type TaskCreatedResources struct {
	// IDs of created AI clusters
	AIClusters []string `json:"ai_clusters"`
	// IDs of created API keys
	APIKeys []string `json:"api_keys"`
	// IDs of created CaaS containers
	CaasContainers []string `json:"caas_containers"`
	// IDs of created ddos protection profiles
	DdosProfiles []int64 `json:"ddos_profiles"`
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
	// IDs of created heat resources
	Heat []string `json:"heat"`
	// IDs of created images
	Images []string `json:"images"`
	// IDs of created inference instances
	InferenceInstances []string `json:"inference_instances"`
	// IDs of created instances
	Instances []string `json:"instances"`
	// IDs of created Kubernetes clusters
	K8sClusters []string `json:"k8s_clusters"`
	// IDs of created Kubernetes pools
	K8sPools []string `json:"k8s_pools"`
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
	PostgresqlClusters []string `json:"postgresql_clusters"`
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
	// IDs of created server groups
	Servergroups []string `json:"servergroups"`
	// IDs of created volume snapshots
	Snapshots []string `json:"snapshots"`
	// IDs of created subnets
	Subnets []string `json:"subnets"`
	// IDs of created volumes
	Volumes []string                 `json:"volumes"`
	JSON    taskCreatedResourcesJSON `json:"-"`
}

// taskCreatedResourcesJSON contains the JSON metadata for the struct
// [TaskCreatedResources]
type taskCreatedResourcesJSON struct {
	AIClusters         apijson.Field
	APIKeys            apijson.Field
	CaasContainers     apijson.Field
	DdosProfiles       apijson.Field
	FaasFunctions      apijson.Field
	FaasNamespaces     apijson.Field
	FileShares         apijson.Field
	Floatingips        apijson.Field
	Healthmonitors     apijson.Field
	Heat               apijson.Field
	Images             apijson.Field
	InferenceInstances apijson.Field
	Instances          apijson.Field
	K8sClusters        apijson.Field
	K8sPools           apijson.Field
	L7polices          apijson.Field
	L7rules            apijson.Field
	LaasTopic          apijson.Field
	Listeners          apijson.Field
	Loadbalancers      apijson.Field
	Members            apijson.Field
	Networks           apijson.Field
	Pools              apijson.Field
	Ports              apijson.Field
	PostgresqlClusters apijson.Field
	Projects           apijson.Field
	RegistryRegistries apijson.Field
	RegistryUsers      apijson.Field
	Routers            apijson.Field
	Secrets            apijson.Field
	Servergroups       apijson.Field
	Snapshots          apijson.Field
	Subnets            apijson.Field
	Volumes            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TaskCreatedResources) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskCreatedResourcesJSON) RawJSON() string {
	return r.raw
}

// Task detailed state that is more specific to task type
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

func (r TaskDetailedState) IsKnown() bool {
	switch r {
	case TaskDetailedStateClusterCleanUp, TaskDetailedStateClusterResize, TaskDetailedStateClusterResume, TaskDetailedStateClusterSuspend, TaskDetailedStateError, TaskDetailedStateFinished, TaskDetailedStateIpuServers, TaskDetailedStateNetwork, TaskDetailedStatePoplarServers, TaskDetailedStatePostDeploySetup, TaskDetailedStateVipuController:
		return true
	}
	return false
}

type TaskCollection struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []Task             `json:"results,required"`
	JSON    taskCollectionJSON `json:"-"`
}

// taskCollectionJSON contains the JSON metadata for the struct [TaskCollection]
type taskCollectionJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskCollectionJSON) RawJSON() string {
	return r.raw
}

type CloudV1TaskListParams struct {
	// ISO formatted datetime string. Filter the tasks by creation date greater than or
	// equal to from_timestamp
	FromTimestamp param.Field[time.Time] `query:"from_timestamp" format:"date-time"`
	// Filter the tasks by their acknowledgement status
	IsAcknowledged param.Field[bool] `query:"is_acknowledged"`
	// Limit the number of returned tasks. Falls back to default of 10 if not
	// specified. Limited by max limit value of 1000
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
	// The project ID to filter the tasks by project. Supports multiple values of kind
	// key=value1&key=value2
	ProjectID param.Field[[]int64] `query:"project_id"`
	// The region ID to filter the tasks by region. Supports multiple values of kind
	// key=value1&key=value2
	RegionID param.Field[[]int64] `query:"region_id"`
	// Sorting by creation date. Oldest first, or most recent first
	Sorting param.Field[CloudV1TaskListParamsSorting] `query:"sorting"`
	// Filter the tasks by state. Supports multiple values of kind
	// key=value1&key=value2
	State param.Field[[]CloudV1TaskListParamsState] `query:"state"`
	// Filter the tasks by their type one of ['activate_ddos_profile',
	// 'attach_bm_to_reserved_fixed_ip', 'attach_vm_interface',
	// 'attach_vm_to_reserved_fixed_ip', 'attach_volume', 'create_ai_cluster_gpu',
	// 'create_bm', 'create_caas_container', 'create_dbaas_postgres_cluster',
	// 'create_ddos_profile', 'create_faas_function', 'create_faas_namespace',
	// 'create_fip', 'create_gpu_virtual_cluster', 'create_image',
	// 'create_inference_instance', 'create_inference_instance_key',
	// 'create_k8s_cluster_pool_v2', 'create_k8s_cluster_v2', 'create_l7policy',
	// 'create_l7rule', 'create_lblistener', 'create_lbmember', 'create_lbpool',
	// 'create_lbpool_health_monitor', 'create_loadbalancer', 'create_network',
	// 'create_reserved_fixed_ip', 'create_router', 'create_secret',
	// 'create_servergroup', 'create_sfs', 'create_snapshot', 'create_subnet',
	// 'create_vm', 'create_volume', 'deactivate_ddos_profile',
	// 'delete_ai_cluster_gpu', 'delete_caas_container',
	// 'delete_dbaas_postgres_cluster', 'delete_ddos_profile', 'delete_faas_function',
	// 'delete_faas_namespace', 'delete_fip', 'delete_gpu_virtual_cluster',
	// 'delete_image', 'delete_inference_instance', 'delete_k8s_cluster_pool_v2',
	// 'delete_k8s_cluster_v2', 'delete_l7policy', 'delete_l7rule',
	// 'delete_lblistener', 'delete_lbmember', 'delete_lbmetadata', 'delete_lbpool',
	// 'delete_loadbalancer', 'delete_network', 'delete_reserved_fixed_ip',
	// 'delete_router', 'delete_secret', 'delete_servergroup', 'delete_sfs',
	// 'delete_snapshot', 'delete_subnet', 'delete_vm', 'delete_volume',
	// 'detach_vm_from_subnet', 'detach_vm_interface', 'detach_volume',
	// 'download_image', 'downscale_ai_cluster_gpu', 'extend_sfs', 'extend_volume',
	// 'failover_loadbalancer', 'hard_reboot_vm', 'patch_caas_container',
	// 'patch_dbaas_postgres_cluster', 'patch_faas_function', 'patch_faas_namespace',
	// 'patch_lblistener', 'patch_lbpool', 'put_into_server_group', 'put_l7policy',
	// 'put_l7rule', 'rebuild_bm', 'rebuild_gpu_baremetal_node',
	// 'remove_from_server_group', 'replace_lbmetadata', 'resize_k8s_cluster_v2',
	// 'resize_loadbalancer', 'resize_vm', 'resume_vm', 'revert_volume',
	// 'soft_reboot_vm', 'start_vm', 'stop_vm', 'suspend_vm', 'sync_private_flavors',
	// 'update_ddos_profile', 'update_inference_instance',
	// 'update_inference_instance_key', 'update_k8s_cluster_v2', 'update_lbmetadata',
	// 'update_port_allowed_address_pairs', 'upgrade_k8s_cluster_v2',
	// 'upscale_ai_cluster_gpu']
	TaskType param.Field[string] `query:"task_type"`
	// ISO formatted datetime string. Filter the tasks by creation date less than or
	// equal to to_timestamp
	ToTimestamp param.Field[time.Time] `query:"to_timestamp" format:"date-time"`
}

// URLQuery serializes [CloudV1TaskListParams]'s query parameters as `url.Values`.
func (r CloudV1TaskListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sorting by creation date. Oldest first, or most recent first
type CloudV1TaskListParamsSorting string

const (
	CloudV1TaskListParamsSortingAsc  CloudV1TaskListParamsSorting = "asc"
	CloudV1TaskListParamsSortingDesc CloudV1TaskListParamsSorting = "desc"
)

func (r CloudV1TaskListParamsSorting) IsKnown() bool {
	switch r {
	case CloudV1TaskListParamsSortingAsc, CloudV1TaskListParamsSortingDesc:
		return true
	}
	return false
}

type CloudV1TaskListParamsState string

const (
	CloudV1TaskListParamsStateError    CloudV1TaskListParamsState = "ERROR"
	CloudV1TaskListParamsStateFinished CloudV1TaskListParamsState = "FINISHED"
	CloudV1TaskListParamsStateNew      CloudV1TaskListParamsState = "NEW"
	CloudV1TaskListParamsStateRunning  CloudV1TaskListParamsState = "RUNNING"
)

func (r CloudV1TaskListParamsState) IsKnown() bool {
	switch r {
	case CloudV1TaskListParamsStateError, CloudV1TaskListParamsStateFinished, CloudV1TaskListParamsStateNew, CloudV1TaskListParamsStateRunning:
		return true
	}
	return false
}

type CloudV1TaskAcknowledgeParams struct {
	// Project ID of the project where tasks should be acknowledged
	ProjectID param.Field[int64] `query:"project_id"`
	// Region ID of the region where tasks should be acknowledged
	RegionID param.Field[int64] `query:"region_id"`
	// List of tasks ids
	TaskIDs param.Field[[]string] `json:"task_ids"`
}

func (r CloudV1TaskAcknowledgeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [CloudV1TaskAcknowledgeParams]'s query parameters as
// `url.Values`.
func (r CloudV1TaskAcknowledgeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1TaskAcknowledgeAllParams struct {
	// Project ID
	ProjectID param.Field[int64] `query:"project_id"`
	// Region ID
	RegionID param.Field[int64] `query:"region_id"`
}

// URLQuery serializes [CloudV1TaskAcknowledgeAllParams]'s query parameters as
// `url.Values`.
func (r CloudV1TaskAcknowledgeAllParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1TaskListActiveParams struct {
	// Limit the number of returned tasks. Falls back to default of 10 if not
	// specified. Limited by max limit value of 1000
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
	// Filter the tasks by their type one of ['activate_ddos_profile',
	// 'attach_bm_to_reserved_fixed_ip', 'attach_vm_interface',
	// 'attach_vm_to_reserved_fixed_ip', 'attach_volume', 'create_ai_cluster_gpu',
	// 'create_bm', 'create_caas_container', 'create_dbaas_postgres_cluster',
	// 'create_ddos_profile', 'create_faas_function', 'create_faas_namespace',
	// 'create_fip', 'create_gpu_virtual_cluster', 'create_image',
	// 'create_inference_instance', 'create_inference_instance_key',
	// 'create_k8s_cluster_pool_v2', 'create_k8s_cluster_v2', 'create_l7policy',
	// 'create_l7rule', 'create_lblistener', 'create_lbmember', 'create_lbpool',
	// 'create_lbpool_health_monitor', 'create_loadbalancer', 'create_network',
	// 'create_reserved_fixed_ip', 'create_router', 'create_secret',
	// 'create_servergroup', 'create_sfs', 'create_snapshot', 'create_subnet',
	// 'create_vm', 'create_volume', 'deactivate_ddos_profile',
	// 'delete_ai_cluster_gpu', 'delete_caas_container',
	// 'delete_dbaas_postgres_cluster', 'delete_ddos_profile', 'delete_faas_function',
	// 'delete_faas_namespace', 'delete_fip', 'delete_gpu_virtual_cluster',
	// 'delete_image', 'delete_inference_instance', 'delete_k8s_cluster_pool_v2',
	// 'delete_k8s_cluster_v2', 'delete_l7policy', 'delete_l7rule',
	// 'delete_lblistener', 'delete_lbmember', 'delete_lbmetadata', 'delete_lbpool',
	// 'delete_loadbalancer', 'delete_network', 'delete_reserved_fixed_ip',
	// 'delete_router', 'delete_secret', 'delete_servergroup', 'delete_sfs',
	// 'delete_snapshot', 'delete_subnet', 'delete_vm', 'delete_volume',
	// 'detach_vm_from_subnet', 'detach_vm_interface', 'detach_volume',
	// 'download_image', 'downscale_ai_cluster_gpu', 'extend_sfs', 'extend_volume',
	// 'failover_loadbalancer', 'hard_reboot_vm', 'patch_caas_container',
	// 'patch_dbaas_postgres_cluster', 'patch_faas_function', 'patch_faas_namespace',
	// 'patch_lblistener', 'patch_lbpool', 'put_into_server_group', 'put_l7policy',
	// 'put_l7rule', 'rebuild_bm', 'rebuild_gpu_baremetal_node',
	// 'remove_from_server_group', 'replace_lbmetadata', 'resize_k8s_cluster_v2',
	// 'resize_loadbalancer', 'resize_vm', 'resume_vm', 'revert_volume',
	// 'soft_reboot_vm', 'start_vm', 'stop_vm', 'suspend_vm', 'sync_private_flavors',
	// 'update_ddos_profile', 'update_inference_instance',
	// 'update_inference_instance_key', 'update_k8s_cluster_v2', 'update_lbmetadata',
	// 'update_port_allowed_address_pairs', 'upgrade_k8s_cluster_v2',
	// 'upscale_ai_cluster_gpu']
	TaskType param.Field[[]string] `query:"task_type"`
}

// URLQuery serializes [CloudV1TaskListActiveParams]'s query parameters as
// `url.Values`.
func (r CloudV1TaskListActiveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
