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

// CloudV1CaaContainerService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1CaaContainerService] method instead.
type CloudV1CaaContainerService struct {
	Options []option.RequestOption
}

// NewCloudV1CaaContainerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1CaaContainerService(opts ...option.RequestOption) (r *CloudV1CaaContainerService) {
	r = &CloudV1CaaContainerService{}
	r.Options = opts
	return
}

// Create container
func (r *CloudV1CaaContainerService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1CaaContainerNewParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/caas/%v/%v/containers", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get container
func (r *CloudV1CaaContainerService) Get(ctx context.Context, projectID int64, regionID int64, containerName string, opts ...option.RequestOption) (res *ContainerSerializer, err error) {
	opts = append(r.Options[:], opts...)
	if containerName == "" {
		err = errors.New("missing required container_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/caas/%v/%v/containers/%s", projectID, regionID, containerName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change container
func (r *CloudV1CaaContainerService) Update(ctx context.Context, projectID int64, regionID int64, containerName string, body CloudV1CaaContainerUpdateParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if containerName == "" {
		err = errors.New("missing required container_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/caas/%v/%v/containers/%s", projectID, regionID, containerName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List containers
func (r *CloudV1CaaContainerService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1CaaContainerListParams, opts ...option.RequestOption) (res *CloudV1CaaContainerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/caas/%v/%v/containers", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete container
func (r *CloudV1CaaContainerService) Delete(ctx context.Context, projectID int64, regionID int64, containerName string, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if containerName == "" {
		err = errors.New("missing required container_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/caas/%v/%v/containers/%s", projectID, regionID, containerName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get container logs
func (r *CloudV1CaaContainerService) GetLogs(ctx context.Context, projectID int64, regionID int64, containerName string, query CloudV1CaaContainerGetLogsParams, opts ...option.RequestOption) (res *CloudV1CaaContainerGetLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if containerName == "" {
		err = errors.New("missing required container_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/caas/%v/%v/containers/%s/logs", projectID, regionID, containerName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ContainerSerializer struct {
	// Container address
	Address string `json:"address,required"`
	// Container creation date
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Container deploy status
	DeployStatus DeployStatus `json:"deploy_status,required"`
	// Container description
	Description string `json:"description,required"`
	// Container environment variables
	Envs interface{} `json:"envs,required"`
	// Container flavor
	Flavor string `json:"flavor,required"`
	// Container image
	Image string `json:"image,required"`
	// Enable/Disable api key authentication
	IsAPIKeyAuth bool `json:"is_api_key_auth,required"`
	// Set to true if container is disabled
	IsDisabled bool `json:"is_disabled,required"`
	// Container listening port
	ListeningPort int64 `json:"listening_port,required"`
	// Container name
	Name string `json:"name,required"`
	// Container namespace
	Namespace string `json:"namespace,required"`
	// Image pull secret
	PullSecret string `json:"pull_secret,required,nullable"`
	// Container autoscaling
	Scale Scale `json:"scale,required"`
	// Source of the container, can be 'cloud' or 'iate'
	Source string `json:"source,required"`
	// Container status
	Status string `json:"status,required"`
	// Container status message
	StatusMessage string `json:"status_message,required,nullable"`
	// Container timeout in seconds
	Timeout int64 `json:"timeout,required"`
	// Container's commands
	Commands string `json:"commands,nullable"`
	// Logging configuration
	Logging Logging                 `json:"logging,nullable"`
	JSON    containerSerializerJSON `json:"-"`
}

// containerSerializerJSON contains the JSON metadata for the struct
// [ContainerSerializer]
type containerSerializerJSON struct {
	Address       apijson.Field
	CreatedAt     apijson.Field
	DeployStatus  apijson.Field
	Description   apijson.Field
	Envs          apijson.Field
	Flavor        apijson.Field
	Image         apijson.Field
	IsAPIKeyAuth  apijson.Field
	IsDisabled    apijson.Field
	ListeningPort apijson.Field
	Name          apijson.Field
	Namespace     apijson.Field
	PullSecret    apijson.Field
	Scale         apijson.Field
	Source        apijson.Field
	Status        apijson.Field
	StatusMessage apijson.Field
	Timeout       apijson.Field
	Commands      apijson.Field
	Logging       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ContainerSerializer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerSerializerJSON) RawJSON() string {
	return r.raw
}

type DeployStatus struct {
	// Number of ready instances
	Ready int64 `json:"ready,required"`
	// Total number of instances
	Total int64            `json:"total,required"`
	JSON  deployStatusJSON `json:"-"`
}

// deployStatusJSON contains the JSON metadata for the struct [DeployStatus]
type deployStatusJSON struct {
	Ready       apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeployStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deployStatusJSON) RawJSON() string {
	return r.raw
}

type Logging struct {
	// Destination region id to which the logs will be written
	DestinationRegionID int64 `json:"destination_region_id,nullable"`
	// Enable/disable forwarding logs to LaaS
	Enabled bool `json:"enabled"`
	// The logs retention policy
	RetentionPolicy LaasIndexRetentionPolicyPydantic `json:"retention_policy,nullable"`
	// The topic name to which the logs will be written
	TopicName string      `json:"topic_name,nullable"`
	JSON      loggingJSON `json:"-"`
}

// loggingJSON contains the JSON metadata for the struct [Logging]
type loggingJSON struct {
	DestinationRegionID apijson.Field
	Enabled             apijson.Field
	RetentionPolicy     apijson.Field
	TopicName           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Logging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loggingJSON) RawJSON() string {
	return r.raw
}

type LoggingParam struct {
	// Destination region id to which the logs will be written
	DestinationRegionID param.Field[int64] `json:"destination_region_id"`
	// Enable/disable forwarding logs to LaaS
	Enabled param.Field[bool] `json:"enabled"`
	// The logs retention policy
	RetentionPolicy param.Field[LaasIndexRetentionPolicyPydanticParam] `json:"retention_policy"`
	// The topic name to which the logs will be written
	TopicName param.Field[string] `json:"topic_name"`
}

func (r LoggingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Scale struct {
	// Maximum number of instances
	Max int64 `json:"max,required"`
	// Minimum number of instances
	Min int64 `json:"min,required"`
	// Cooldown period in seconds
	CooldownPeriod int64 `json:"cooldown_period,nullable"`
	// Scale triggers configuration
	Triggers ScaleTriggers `json:"triggers,nullable"`
	JSON     scaleJSON     `json:"-"`
}

// scaleJSON contains the JSON metadata for the struct [Scale]
type scaleJSON struct {
	Max            apijson.Field
	Min            apijson.Field
	CooldownPeriod apijson.Field
	Triggers       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Scale) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scaleJSON) RawJSON() string {
	return r.raw
}

// Scale triggers configuration
type ScaleTriggers struct {
	// CPU trigger configuration
	CPU ScaleTriggersCPU `json:"cpu,nullable"`
	// HTTP trigger configuration
	HTTP ScaleTriggersHTTP `json:"http,nullable"`
	// Memory trigger configuration
	Memory ScaleTriggersMemory `json:"memory,nullable"`
	JSON   scaleTriggersJSON   `json:"-"`
}

// scaleTriggersJSON contains the JSON metadata for the struct [ScaleTriggers]
type scaleTriggersJSON struct {
	CPU         apijson.Field
	HTTP        apijson.Field
	Memory      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScaleTriggers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scaleTriggersJSON) RawJSON() string {
	return r.raw
}

// CPU trigger configuration
type ScaleTriggersCPU struct {
	// CPU usage percentage
	Threshold int64                `json:"threshold,required"`
	JSON      scaleTriggersCPUJSON `json:"-"`
}

// scaleTriggersCPUJSON contains the JSON metadata for the struct
// [ScaleTriggersCPU]
type scaleTriggersCPUJSON struct {
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScaleTriggersCPU) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scaleTriggersCPUJSON) RawJSON() string {
	return r.raw
}

// HTTP trigger configuration
type ScaleTriggersHTTP struct {
	// HTTP requests rate
	Rate int64 `json:"rate,required"`
	// Time window in seconds
	Window int64                 `json:"window,required"`
	JSON   scaleTriggersHTTPJSON `json:"-"`
}

// scaleTriggersHTTPJSON contains the JSON metadata for the struct
// [ScaleTriggersHTTP]
type scaleTriggersHTTPJSON struct {
	Rate        apijson.Field
	Window      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScaleTriggersHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scaleTriggersHTTPJSON) RawJSON() string {
	return r.raw
}

// Memory trigger configuration
type ScaleTriggersMemory struct {
	// Memory usage percentage
	Threshold int64                   `json:"threshold,required"`
	JSON      scaleTriggersMemoryJSON `json:"-"`
}

// scaleTriggersMemoryJSON contains the JSON metadata for the struct
// [ScaleTriggersMemory]
type scaleTriggersMemoryJSON struct {
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScaleTriggersMemory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scaleTriggersMemoryJSON) RawJSON() string {
	return r.raw
}

type ScaleParam struct {
	// Maximum number of instances
	Max param.Field[int64] `json:"max,required"`
	// Minimum number of instances
	Min param.Field[int64] `json:"min,required"`
	// Cooldown period in seconds
	CooldownPeriod param.Field[int64] `json:"cooldown_period"`
	// Scale triggers configuration
	Triggers param.Field[ScaleTriggersParam] `json:"triggers"`
}

func (r ScaleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Scale triggers configuration
type ScaleTriggersParam struct {
	// CPU trigger configuration
	CPU param.Field[ScaleTriggersCPUParam] `json:"cpu"`
	// HTTP trigger configuration
	HTTP param.Field[ScaleTriggersHTTPParam] `json:"http"`
	// Memory trigger configuration
	Memory param.Field[ScaleTriggersMemoryParam] `json:"memory"`
}

func (r ScaleTriggersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// CPU trigger configuration
type ScaleTriggersCPUParam struct {
	// CPU usage percentage
	Threshold param.Field[int64] `json:"threshold,required"`
}

func (r ScaleTriggersCPUParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// HTTP trigger configuration
type ScaleTriggersHTTPParam struct {
	// HTTP requests rate
	Rate param.Field[int64] `json:"rate,required"`
	// Time window in seconds
	Window param.Field[int64] `json:"window,required"`
}

func (r ScaleTriggersHTTPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Memory trigger configuration
type ScaleTriggersMemoryParam struct {
	// Memory usage percentage
	Threshold param.Field[int64] `json:"threshold,required"`
}

func (r ScaleTriggersMemoryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TaskIDs struct {
	// List of task IDs
	Tasks []string    `json:"tasks,required"`
	JSON  taskIDsJSON `json:"-"`
}

// taskIDsJSON contains the JSON metadata for the struct [TaskIDs]
type taskIDsJSON struct {
	Tasks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskIDs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskIDsJSON) RawJSON() string {
	return r.raw
}

type CloudV1CaaContainerListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []ContainerSerializer               `json:"results,required"`
	JSON    cloudV1CaaContainerListResponseJSON `json:"-"`
}

// cloudV1CaaContainerListResponseJSON contains the JSON metadata for the struct
// [CloudV1CaaContainerListResponse]
type cloudV1CaaContainerListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1CaaContainerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CaaContainerListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1CaaContainerGetLogsResponse struct {
	// Container logs
	Logs []interface{}                          `json:"logs,required"`
	JSON cloudV1CaaContainerGetLogsResponseJSON `json:"-"`
}

// cloudV1CaaContainerGetLogsResponseJSON contains the JSON metadata for the struct
// [CloudV1CaaContainerGetLogsResponse]
type cloudV1CaaContainerGetLogsResponseJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1CaaContainerGetLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1CaaContainerGetLogsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1CaaContainerNewParams struct {
	// Container flavor
	Flavor param.Field[string] `json:"flavor,required"`
	// Container image
	Image param.Field[string] `json:"image,required"`
	// Container listening port
	ListeningPort param.Field[int64] `json:"listening_port,required"`
	// Container name
	Name param.Field[string] `json:"name,required"`
	// Container autoscaling
	Scale param.Field[ScaleParam] `json:"scale,required"`
	// Container's commands
	Commands param.Field[string] `json:"commands"`
	// Container description
	Description param.Field[string] `json:"description"`
	// Container environment variables
	Envs param.Field[interface{}] `json:"envs"`
	// Enable/Disable api key authentication
	IsAPIKeyAuth param.Field[bool] `json:"is_api_key_auth"`
	// Set to true if container is disabled
	IsDisabled param.Field[bool] `json:"is_disabled"`
	// Logging configuration
	Logging param.Field[LoggingParam] `json:"logging"`
	// Image pull secret
	PullSecret param.Field[string] `json:"pull_secret"`
	// Container timeout in seconds
	Timeout param.Field[int64] `json:"timeout"`
}

func (r CloudV1CaaContainerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1CaaContainerUpdateParams struct {
	// Container's commands
	Commands param.Field[string] `json:"commands"`
	// Container description
	Description param.Field[string] `json:"description"`
	// Container environment variables
	Envs param.Field[interface{}] `json:"envs"`
	// Container flavor
	Flavor param.Field[string] `json:"flavor"`
	// Container image
	Image param.Field[string] `json:"image"`
	// Enable/Disable api key authentication
	IsAPIKeyAuth param.Field[bool] `json:"is_api_key_auth"`
	// Set to true if container is disabled
	IsDisabled param.Field[bool] `json:"is_disabled"`
	// Container listening port
	ListeningPort param.Field[int64] `json:"listening_port"`
	// Logging configuration
	Logging param.Field[LoggingParam] `json:"logging"`
	// Image pull secret
	PullSecret param.Field[string] `json:"pull_secret"`
	// Container autoscaling
	Scale param.Field[ScaleParam] `json:"scale"`
	// Container timeout in seconds
	Timeout param.Field[int64] `json:"timeout"`
}

func (r CloudV1CaaContainerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1CaaContainerListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV1CaaContainerListParams]'s query parameters as
// `url.Values`.
func (r CloudV1CaaContainerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1CaaContainerGetLogsParams struct {
	// Limit the number of logs lines
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CloudV1CaaContainerGetLogsParams]'s query parameters as
// `url.Values`.
func (r CloudV1CaaContainerGetLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
