// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// InferenceDeploymentService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceDeploymentService] method instead.
type InferenceDeploymentService struct {
	Options []option.RequestOption
	Logs    InferenceDeploymentLogService
	tasks   TaskService
}

// NewInferenceDeploymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInferenceDeploymentService(opts ...option.RequestOption) (r InferenceDeploymentService) {
	r = InferenceDeploymentService{}
	r.Options = opts
	r.Logs = NewInferenceDeploymentLogService(opts...)
	r.tasks = NewTaskService(opts...)
	return
}

// Create inference deployment
func (r *InferenceDeploymentService) New(ctx context.Context, params InferenceDeploymentNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update inference deployment
func (r *InferenceDeploymentService) Update(ctx context.Context, deploymentName string, params InferenceDeploymentUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if deploymentName == "" {
		err = errors.New("missing required deployment_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments/%s", params.ProjectID.Value, deploymentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List inference deployments
func (r *InferenceDeploymentService) List(ctx context.Context, params InferenceDeploymentListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Inference], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments", params.ProjectID.Value)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List inference deployments
func (r *InferenceDeploymentService) ListAutoPaging(ctx context.Context, params InferenceDeploymentListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Inference] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete inference deployment
func (r *InferenceDeploymentService) Delete(ctx context.Context, deploymentName string, body InferenceDeploymentDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if deploymentName == "" {
		err = errors.New("missing required deployment_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments/%s", body.ProjectID.Value, deploymentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get inference deployment
func (r *InferenceDeploymentService) Get(ctx context.Context, deploymentName string, query InferenceDeploymentGetParams, opts ...option.RequestOption) (res *Inference, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if deploymentName == "" {
		err = errors.New("missing required deployment_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments/%s", query.ProjectID.Value, deploymentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get inference deployment API key
func (r *InferenceDeploymentService) GetAPIKey(ctx context.Context, deploymentName string, query InferenceDeploymentGetAPIKeyParams, opts ...option.RequestOption) (res *InferenceApikeySecret, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if deploymentName == "" {
		err = errors.New("missing required deployment_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments/%s/apikey", query.ProjectID.Value, deploymentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This operation initializes an inference deployment after it was stopped, making
// it available to handle inference requests again. The instance will launch with
// the **minimum** number of replicas defined in the scaling settings.
//
//   - If the minimum replicas are set to **0**, the instance will initially start
//     with **0** replicas.
//   - It will automatically scale up when it receives requests or SQS messages,
//     according to the configured scaling rules.
func (r *InferenceDeploymentService) Start(ctx context.Context, deploymentName string, body InferenceDeploymentStartParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if deploymentName == "" {
		err = errors.New("missing required deployment_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments/%s/start", body.ProjectID.Value, deploymentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// This operation shuts down an inference deployment, making it unavailable for
// handling requests. The deployment will scale down to **0** replicas, overriding
// any minimum replica settings.
//
//   - Once stopped, the deployment will **not** process any inference requests or
//     SQS messages.
//   - It will **not** restart automatically and must be started manually.
//   - While stopped, the deployment will **not** incur any charges.
func (r *InferenceDeploymentService) Stop(ctx context.Context, deploymentName string, body InferenceDeploymentStopParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if deploymentName == "" {
		err = errors.New("missing required deployment_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments/%s/stop", body.ProjectID.Value, deploymentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type Container struct {
	// Address of the inference instance
	Address string `json:"address,required" format:"uri"`
	// Status of the containers deployment
	DeployStatus DeployStatus `json:"deploy_status,required"`
	// Error message if the container deployment failed
	ErrorMessage string `json:"error_message,required"`
	// Region name for the container
	RegionID int64 `json:"region_id,required"`
	// Scale for the container
	Scale ContainerScale `json:"scale,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address      respjson.Field
		DeployStatus respjson.Field
		ErrorMessage respjson.Field
		RegionID     respjson.Field
		Scale        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Container) RawJSON() string { return r.JSON.raw }
func (r *Container) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Inference struct {
	// Address of the inference instance
	Address string `json:"address,required" format:"uri"`
	// `true` if instance uses API key authentication.
	// `"Authorization": "Bearer *****"` or `"X-Api-Key": "*****"` header is required
	// for the requests to the instance if enabled.
	AuthEnabled bool `json:"auth_enabled,required"`
	// Command to be executed when running a container from an image.
	Command string `json:"command,required"`
	// List of containers for the inference instance
	Containers []Container `json:"containers,required"`
	// Inference instance creation date in ISO 8601 format.
	CreatedAt string `json:"created_at,required"`
	// Registry credentials name
	CredentialsName string `json:"credentials_name,required"`
	// Inference instance description.
	Description string `json:"description,required"`
	// Environment variables for the inference instance
	Envs map[string]string `json:"envs,required"`
	// Flavor name for the inference instance
	FlavorName string `json:"flavor_name,required"`
	// Docker image for the inference instance. This field should contain the image
	// name and tag in the format 'name:tag', e.g., 'nginx:latest'. It defaults to
	// Docker Hub as the image registry, but any accessible Docker image URL can be
	// specified.
	Image string `json:"image,required"`
	// Ingress options for the inference instance
	IngressOpts IngressOptsOut `json:"ingress_opts,required"`
	// Listening port for the inference instance.
	ListeningPort int64 `json:"listening_port,required"`
	// Logging configuration for the inference instance
	Logging Logging `json:"logging,required"`
	// Inference instance name.
	Name string `json:"name,required"`
	// Probes configured for all containers of the inference instance.
	Probes InferenceProbes `json:"probes,required"`
	// Project ID. If not provided, your default project ID will be used.
	ProjectID int64 `json:"project_id,required"`
	// Inference instance status.
	//
	// Value can be one of the following:
	//
	//   - `DEPLOYING` - The instance is being deployed. Containers are not yet created.
	//   - `PARTIALLYDEPLOYED` - All containers have been created, but some may not be
	//     ready yet. Instances stuck in this state typically indicate either image being
	//     pulled, or a failure of some kind. In the latter case, the `error_message`
	//     field of the respective container object in the `containers` collection
	//     explains the failure reason.
	//   - `ACTIVE` - The instance is running and ready to accept requests.
	//   - `DISABLED` - The instance is disabled and not accepting any requests.
	//   - `DELETING` - The instance is being deleted.
	//
	// Any of "ACTIVE", "DELETING", "DEPLOYING", "DISABLED", "PARTIALLYDEPLOYED".
	Status InferenceStatus `json:"status,required"`
	// Specifies the duration in seconds without any requests after which the
	// containers will be downscaled to their minimum scale value as defined by
	// `scale.min`. If set, this helps in optimizing resource usage by reducing the
	// number of container instances during periods of inactivity.
	Timeout int64 `json:"timeout,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address         respjson.Field
		AuthEnabled     respjson.Field
		Command         respjson.Field
		Containers      respjson.Field
		CreatedAt       respjson.Field
		CredentialsName respjson.Field
		Description     respjson.Field
		Envs            respjson.Field
		FlavorName      respjson.Field
		Image           respjson.Field
		IngressOpts     respjson.Field
		ListeningPort   respjson.Field
		Logging         respjson.Field
		Name            respjson.Field
		Probes          respjson.Field
		ProjectID       respjson.Field
		Status          respjson.Field
		Timeout         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Inference) RawJSON() string { return r.JSON.raw }
func (r *Inference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inference instance status.
//
// Value can be one of the following:
//
//   - `DEPLOYING` - The instance is being deployed. Containers are not yet created.
//   - `PARTIALLYDEPLOYED` - All containers have been created, but some may not be
//     ready yet. Instances stuck in this state typically indicate either image being
//     pulled, or a failure of some kind. In the latter case, the `error_message`
//     field of the respective container object in the `containers` collection
//     explains the failure reason.
//   - `ACTIVE` - The instance is running and ready to accept requests.
//   - `DISABLED` - The instance is disabled and not accepting any requests.
//   - `DELETING` - The instance is being deleted.
type InferenceStatus string

const (
	InferenceStatusActive            InferenceStatus = "ACTIVE"
	InferenceStatusDeleting          InferenceStatus = "DELETING"
	InferenceStatusDeploying         InferenceStatus = "DEPLOYING"
	InferenceStatusDisabled          InferenceStatus = "DISABLED"
	InferenceStatusPartiallydeployed InferenceStatus = "PARTIALLYDEPLOYED"
)

type InferenceApikeySecret struct {
	// API key secret
	Secret string `json:"secret,required"`
	// API key status
	//
	// Any of "PENDING", "READY".
	Status InferenceApikeySecretStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Secret      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApikeySecret) RawJSON() string { return r.JSON.raw }
func (r *InferenceApikeySecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API key status
type InferenceApikeySecretStatus string

const (
	InferenceApikeySecretStatusPending InferenceApikeySecretStatus = "PENDING"
	InferenceApikeySecretStatusReady   InferenceApikeySecretStatus = "READY"
)

type InferenceLog struct {
	// Log message.
	Message string `json:"message,required"`
	// Pod name.
	Pod string `json:"pod,required"`
	// Region ID where the container is deployed.
	RegionID int64 `json:"region_id,required"`
	// Log message timestamp in ISO 8601 format.
	Time time.Time `json:"time,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Pod         respjson.Field
		RegionID    respjson.Field
		Time        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceLog) RawJSON() string { return r.JSON.raw }
func (r *InferenceLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceDeploymentNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// List of containers for the inference instance.
	Containers []InferenceDeploymentNewParamsContainer `json:"containers,omitzero,required"`
	// Flavor name for the inference instance.
	FlavorName string `json:"flavor_name,required"`
	// Docker image for the inference instance. This field should contain the image
	// name and tag in the format 'name:tag', e.g., 'nginx:latest'. It defaults to
	// Docker Hub as the image registry, but any accessible Docker image URL can be
	// specified.
	Image string `json:"image,required"`
	// Listening port for the inference instance.
	ListeningPort int64 `json:"listening_port,required"`
	// Inference instance name.
	Name string `json:"name,required"`
	// Registry credentials name
	CredentialsName param.Opt[string] `json:"credentials_name,omitzero"`
	// Inference instance description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Specifies the duration in seconds without any requests after which the
	// containers will be downscaled to their minimum scale value as defined by
	// `scale.min`. If set, this helps in optimizing resource usage by reducing the
	// number of container instances during periods of inactivity. The default value
	// when the parameter is not set is 120.
	Timeout param.Opt[int64] `json:"timeout,omitzero"`
	// Set to `true` to enable API key authentication for the inference instance.
	// `"Authorization": "Bearer *****"` or `"X-Api-Key": "*****"` header is required
	// for the requests to the instance if enabled
	AuthEnabled param.Opt[bool] `json:"auth_enabled,omitzero"`
	// Command to be executed when running a container from an image.
	Command []string `json:"command,omitzero"`
	// Logging configuration for the inference instance
	Logging InferenceDeploymentNewParamsLogging `json:"logging,omitzero"`
	// Probes configured for all containers of the inference instance. If probes are
	// not provided, and the image_name is from a the Model Catalog registry, the
	// default probes will be used.
	Probes InferenceDeploymentNewParamsProbes `json:"probes,omitzero"`
	// Environment variables for the inference instance.
	Envs map[string]string `json:"envs,omitzero"`
	// Ingress options for the inference instance
	IngressOpts IngressOptsParam `json:"ingress_opts,omitzero"`
	paramObj
}

func (r InferenceDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties RegionID, Scale are required.
type InferenceDeploymentNewParamsContainer struct {
	// Region id for the container
	RegionID int64 `json:"region_id,required"`
	// Scale for the container
	Scale InferenceDeploymentNewParamsContainerScale `json:"scale,omitzero,required"`
	paramObj
}

func (r InferenceDeploymentNewParamsContainer) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsContainer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scale for the container
//
// The properties Max, Min are required.
type InferenceDeploymentNewParamsContainerScale struct {
	// Maximum scale for the container
	Max int64 `json:"max,required"`
	// Minimum scale for the container
	Min int64 `json:"min,required"`
	// Cooldown period between scaling actions in seconds
	CooldownPeriod param.Opt[int64] `json:"cooldown_period,omitzero"`
	// Polling interval for scaling triggers in seconds
	PollingInterval param.Opt[int64] `json:"polling_interval,omitzero"`
	// Triggers for scaling actions
	Triggers InferenceDeploymentNewParamsContainerScaleTriggers `json:"triggers,omitzero"`
	paramObj
}

func (r InferenceDeploymentNewParamsContainerScale) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsContainerScale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Triggers for scaling actions
type InferenceDeploymentNewParamsContainerScaleTriggers struct {
	// CPU trigger configuration
	CPU InferenceDeploymentNewParamsContainerScaleTriggersCPU `json:"cpu,omitzero"`
	// GPU memory trigger configuration. Calculated by DCGM_FI_DEV_MEM_COPY_UTIL metric
	GPUMemory InferenceDeploymentNewParamsContainerScaleTriggersGPUMemory `json:"gpu_memory,omitzero"`
	// GPU utilization trigger configuration. Calculated by DCGM_FI_DEV_GPU_UTIL metric
	GPUUtilization InferenceDeploymentNewParamsContainerScaleTriggersGPUUtilization `json:"gpu_utilization,omitzero"`
	// HTTP trigger configuration
	HTTP InferenceDeploymentNewParamsContainerScaleTriggersHTTP `json:"http,omitzero"`
	// Memory trigger configuration
	Memory InferenceDeploymentNewParamsContainerScaleTriggersMemory `json:"memory,omitzero"`
	// SQS trigger configuration
	Sqs InferenceDeploymentNewParamsContainerScaleTriggersSqs `json:"sqs,omitzero"`
	paramObj
}

func (r InferenceDeploymentNewParamsContainerScaleTriggers) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScaleTriggers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsContainerScaleTriggers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CPU trigger configuration
//
// The property Threshold is required.
type InferenceDeploymentNewParamsContainerScaleTriggersCPU struct {
	// Threshold value for the trigger in percentage
	Threshold int64 `json:"threshold,required"`
	paramObj
}

func (r InferenceDeploymentNewParamsContainerScaleTriggersCPU) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScaleTriggersCPU
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsContainerScaleTriggersCPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPU memory trigger configuration. Calculated by DCGM_FI_DEV_MEM_COPY_UTIL metric
//
// The property Threshold is required.
type InferenceDeploymentNewParamsContainerScaleTriggersGPUMemory struct {
	// Threshold value for the trigger in percentage
	Threshold int64 `json:"threshold,required"`
	paramObj
}

func (r InferenceDeploymentNewParamsContainerScaleTriggersGPUMemory) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScaleTriggersGPUMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsContainerScaleTriggersGPUMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPU utilization trigger configuration. Calculated by DCGM_FI_DEV_GPU_UTIL metric
//
// The property Threshold is required.
type InferenceDeploymentNewParamsContainerScaleTriggersGPUUtilization struct {
	// Threshold value for the trigger in percentage
	Threshold int64 `json:"threshold,required"`
	paramObj
}

func (r InferenceDeploymentNewParamsContainerScaleTriggersGPUUtilization) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScaleTriggersGPUUtilization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsContainerScaleTriggersGPUUtilization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP trigger configuration
//
// The properties Rate, Window are required.
type InferenceDeploymentNewParamsContainerScaleTriggersHTTP struct {
	// Request count per 'window' seconds for the http trigger
	Rate int64 `json:"rate,required"`
	// Time window for rate calculation in seconds
	Window int64 `json:"window,required"`
	paramObj
}

func (r InferenceDeploymentNewParamsContainerScaleTriggersHTTP) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScaleTriggersHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsContainerScaleTriggersHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Memory trigger configuration
//
// The property Threshold is required.
type InferenceDeploymentNewParamsContainerScaleTriggersMemory struct {
	// Threshold value for the trigger in percentage
	Threshold int64 `json:"threshold,required"`
	paramObj
}

func (r InferenceDeploymentNewParamsContainerScaleTriggersMemory) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScaleTriggersMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsContainerScaleTriggersMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SQS trigger configuration
//
// The properties ActivationQueueLength, AwsRegion, QueueLength, QueueURL,
// SecretName are required.
type InferenceDeploymentNewParamsContainerScaleTriggersSqs struct {
	// Number of messages for activation
	ActivationQueueLength int64 `json:"activation_queue_length,required"`
	// AWS region
	AwsRegion string `json:"aws_region,required"`
	// Number of messages for one replica
	QueueLength int64 `json:"queue_length,required"`
	// SQS queue URL
	QueueURL string `json:"queue_url,required"`
	// Auth secret name
	SecretName string `json:"secret_name,required"`
	// Custom AWS endpoint
	AwsEndpoint param.Opt[string] `json:"aws_endpoint,omitzero"`
	// Scale on delayed messages
	ScaleOnDelayed param.Opt[bool] `json:"scale_on_delayed,omitzero"`
	// Scale on in-flight messages
	ScaleOnFlight param.Opt[bool] `json:"scale_on_flight,omitzero"`
	paramObj
}

func (r InferenceDeploymentNewParamsContainerScaleTriggersSqs) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScaleTriggersSqs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsContainerScaleTriggersSqs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Logging configuration for the inference instance
type InferenceDeploymentNewParamsLogging struct {
	// ID of the region in which the logs will be stored
	DestinationRegionID param.Opt[int64] `json:"destination_region_id,omitzero"`
	// The topic name to stream logs to
	TopicName param.Opt[string] `json:"topic_name,omitzero"`
	// Enable or disable log streaming
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Logs retention policy
	RetentionPolicy LaasIndexRetentionPolicyParam `json:"retention_policy,omitzero"`
	paramObj
}

func (r InferenceDeploymentNewParamsLogging) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsLogging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsLogging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Probes configured for all containers of the inference instance. If probes are
// not provided, and the image_name is from a the Model Catalog registry, the
// default probes will be used.
type InferenceDeploymentNewParamsProbes struct {
	// Liveness probe configuration
	LivenessProbe ContainerProbeConfigCreateParam `json:"liveness_probe,omitzero"`
	// Readiness probe configuration
	ReadinessProbe ContainerProbeConfigCreateParam `json:"readiness_probe,omitzero"`
	// Startup probe configuration
	StartupProbe ContainerProbeConfigCreateParam `json:"startup_probe,omitzero"`
	paramObj
}

func (r InferenceDeploymentNewParamsProbes) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbes
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceDeploymentUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Set to `true` to enable API key authentication for the inference instance.
	// `"Authorization": "Bearer *****"` or `"X-Api-Key": "*****"` header is required
	// for the requests to the instance if enabled
	AuthEnabled param.Opt[bool] `json:"auth_enabled,omitzero"`
	// Registry credentials name
	CredentialsName param.Opt[string] `json:"credentials_name,omitzero"`
	// Inference instance description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Flavor name for the inference instance.
	FlavorName param.Opt[string] `json:"flavor_name,omitzero"`
	// Docker image for the inference instance. This field should contain the image
	// name and tag in the format 'name:tag', e.g., 'nginx:latest'. It defaults to
	// Docker Hub as the image registry, but any accessible Docker image URL can be
	// specified.
	Image param.Opt[string] `json:"image,omitzero"`
	// Listening port for the inference instance.
	ListeningPort param.Opt[int64] `json:"listening_port,omitzero"`
	// Specifies the duration in seconds without any requests after which the
	// containers will be downscaled to their minimum scale value as defined by
	// `scale.min`. If set, this helps in optimizing resource usage by reducing the
	// number of container instances during periods of inactivity. The default value
	// when the parameter is not set is 120.
	Timeout param.Opt[int64] `json:"timeout,omitzero"`
	// Command to be executed when running a container from an image.
	Command []string `json:"command,omitzero"`
	// List of containers for the inference instance.
	Containers []InferenceDeploymentUpdateParamsContainer `json:"containers,omitzero"`
	// Environment variables for the inference instance.
	Envs map[string]string `json:"envs,omitzero"`
	// Logging configuration for the inference instance
	Logging InferenceDeploymentUpdateParamsLogging `json:"logging,omitzero"`
	// Probes configured for all containers of the inference instance.
	Probes InferenceDeploymentUpdateParamsProbes `json:"probes,omitzero"`
	// Ingress options for the inference instance
	IngressOpts IngressOptsParam `json:"ingress_opts,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties RegionID, Scale are required.
type InferenceDeploymentUpdateParamsContainer struct {
	// Region id for the container
	RegionID param.Opt[int64] `json:"region_id,omitzero,required"`
	// Scale for the container
	Scale InferenceDeploymentUpdateParamsContainerScale `json:"scale,omitzero,required"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsContainer) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsContainer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scale for the container
//
// The properties Max, Min are required.
type InferenceDeploymentUpdateParamsContainerScale struct {
	// Maximum scale for the container
	Max int64 `json:"max,required"`
	// Minimum scale for the container
	Min int64 `json:"min,required"`
	// Cooldown period between scaling actions in seconds
	CooldownPeriod param.Opt[int64] `json:"cooldown_period,omitzero"`
	// Polling interval for scaling triggers in seconds
	PollingInterval param.Opt[int64] `json:"polling_interval,omitzero"`
	// Triggers for scaling actions
	Triggers InferenceDeploymentUpdateParamsContainerScaleTriggers `json:"triggers,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsContainerScale) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsContainerScale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Triggers for scaling actions
type InferenceDeploymentUpdateParamsContainerScaleTriggers struct {
	// CPU trigger configuration
	CPU InferenceDeploymentUpdateParamsContainerScaleTriggersCPU `json:"cpu,omitzero"`
	// GPU memory trigger configuration. Calculated by DCGM_FI_DEV_MEM_COPY_UTIL metric
	GPUMemory InferenceDeploymentUpdateParamsContainerScaleTriggersGPUMemory `json:"gpu_memory,omitzero"`
	// GPU utilization trigger configuration. Calculated by DCGM_FI_DEV_GPU_UTIL metric
	GPUUtilization InferenceDeploymentUpdateParamsContainerScaleTriggersGPUUtilization `json:"gpu_utilization,omitzero"`
	// HTTP trigger configuration
	HTTP InferenceDeploymentUpdateParamsContainerScaleTriggersHTTP `json:"http,omitzero"`
	// Memory trigger configuration
	Memory InferenceDeploymentUpdateParamsContainerScaleTriggersMemory `json:"memory,omitzero"`
	// SQS trigger configuration
	Sqs InferenceDeploymentUpdateParamsContainerScaleTriggersSqs `json:"sqs,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsContainerScaleTriggers) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScaleTriggers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsContainerScaleTriggers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CPU trigger configuration
//
// The property Threshold is required.
type InferenceDeploymentUpdateParamsContainerScaleTriggersCPU struct {
	// Threshold value for the trigger in percentage
	Threshold int64 `json:"threshold,required"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsContainerScaleTriggersCPU) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScaleTriggersCPU
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsContainerScaleTriggersCPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPU memory trigger configuration. Calculated by DCGM_FI_DEV_MEM_COPY_UTIL metric
//
// The property Threshold is required.
type InferenceDeploymentUpdateParamsContainerScaleTriggersGPUMemory struct {
	// Threshold value for the trigger in percentage
	Threshold int64 `json:"threshold,required"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsContainerScaleTriggersGPUMemory) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScaleTriggersGPUMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsContainerScaleTriggersGPUMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPU utilization trigger configuration. Calculated by DCGM_FI_DEV_GPU_UTIL metric
//
// The property Threshold is required.
type InferenceDeploymentUpdateParamsContainerScaleTriggersGPUUtilization struct {
	// Threshold value for the trigger in percentage
	Threshold int64 `json:"threshold,required"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsContainerScaleTriggersGPUUtilization) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScaleTriggersGPUUtilization
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsContainerScaleTriggersGPUUtilization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP trigger configuration
//
// The properties Rate, Window are required.
type InferenceDeploymentUpdateParamsContainerScaleTriggersHTTP struct {
	// Request count per 'window' seconds for the http trigger
	Rate int64 `json:"rate,required"`
	// Time window for rate calculation in seconds
	Window int64 `json:"window,required"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsContainerScaleTriggersHTTP) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScaleTriggersHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsContainerScaleTriggersHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Memory trigger configuration
//
// The property Threshold is required.
type InferenceDeploymentUpdateParamsContainerScaleTriggersMemory struct {
	// Threshold value for the trigger in percentage
	Threshold int64 `json:"threshold,required"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsContainerScaleTriggersMemory) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScaleTriggersMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsContainerScaleTriggersMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SQS trigger configuration
//
// The properties ActivationQueueLength, AwsRegion, QueueLength, QueueURL,
// SecretName are required.
type InferenceDeploymentUpdateParamsContainerScaleTriggersSqs struct {
	// Number of messages for activation
	ActivationQueueLength int64 `json:"activation_queue_length,required"`
	// AWS region
	AwsRegion string `json:"aws_region,required"`
	// Number of messages for one replica
	QueueLength int64 `json:"queue_length,required"`
	// SQS queue URL
	QueueURL string `json:"queue_url,required"`
	// Auth secret name
	SecretName string `json:"secret_name,required"`
	// Custom AWS endpoint
	AwsEndpoint param.Opt[string] `json:"aws_endpoint,omitzero"`
	// Scale on delayed messages
	ScaleOnDelayed param.Opt[bool] `json:"scale_on_delayed,omitzero"`
	// Scale on in-flight messages
	ScaleOnFlight param.Opt[bool] `json:"scale_on_flight,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsContainerScaleTriggersSqs) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScaleTriggersSqs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsContainerScaleTriggersSqs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Logging configuration for the inference instance
type InferenceDeploymentUpdateParamsLogging struct {
	// ID of the region in which the logs will be stored
	DestinationRegionID param.Opt[int64] `json:"destination_region_id,omitzero"`
	// The topic name to stream logs to
	TopicName param.Opt[string] `json:"topic_name,omitzero"`
	// Enable or disable log streaming
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Logs retention policy
	RetentionPolicy LaasIndexRetentionPolicyParam `json:"retention_policy,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsLogging) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsLogging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsLogging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Probes configured for all containers of the inference instance.
type InferenceDeploymentUpdateParamsProbes struct {
	// Liveness probe configuration
	LivenessProbe ContainerProbeConfigCreateParam `json:"liveness_probe,omitzero"`
	// Readiness probe configuration
	ReadinessProbe ContainerProbeConfigCreateParam `json:"readiness_probe,omitzero"`
	// Startup probe configuration
	StartupProbe ContainerProbeConfigCreateParam `json:"startup_probe,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbes) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbes
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceDeploymentListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InferenceDeploymentListParams]'s query parameters as
// `url.Values`.
func (r InferenceDeploymentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InferenceDeploymentDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

type InferenceDeploymentGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

type InferenceDeploymentGetAPIKeyParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

type InferenceDeploymentStartParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

type InferenceDeploymentStopParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

// NewAndPoll creates a new inference deployment and polls for completion
func (r *InferenceDeploymentService) NewAndPoll(ctx context.Context, params InferenceDeploymentNewParams, opts ...option.RequestOption) (v *Inference, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams InferenceDeploymentGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	getParams.ProjectID = params.ProjectID

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]
	task, err := r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.InferenceInstances) != 1 {
		return nil, errors.New("expected exactly one inference deployment to be created in a task")
	}
	resourceID := task.CreatedResources.InferenceInstances[0]

	return r.Get(ctx, resourceID, getParams, opts...)
}

// DeleteAndPoll deletes an inference deployment and polls for completion of the first task. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *InferenceDeploymentService) DeleteAndPoll(ctx context.Context, deploymentName string, params InferenceDeploymentDeleteParams, opts ...option.RequestOption) error {
	resource, err := r.Delete(ctx, deploymentName, params, opts...)
	if err != nil {
		return err
	}

	opts = append(r.Options[:], opts...)
	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	return err
}

// UpdateAndPoll updates an inference deployment and polls for completion of the first task. Use the [TaskService.Poll]
// method if you need to poll for all tasks.
func (r *InferenceDeploymentService) UpdateAndPoll(ctx context.Context, deploymentName string, params InferenceDeploymentUpdateParams, opts ...option.RequestOption) (v *Inference, err error) {
	resource, err := r.Update(ctx, deploymentName, params, opts...)
	if err != nil {
		return
	}

	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams InferenceDeploymentGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	getParams.ProjectID = params.ProjectID

	if len(resource.Tasks) == 0 {
		return nil, errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, deploymentName, getParams, opts...)
}
