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

// CloudV3InferenceDeploymentService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV3InferenceDeploymentService] method instead.
type CloudV3InferenceDeploymentService struct {
	Options []option.RequestOption
}

// NewCloudV3InferenceDeploymentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV3InferenceDeploymentService(opts ...option.RequestOption) (r *CloudV3InferenceDeploymentService) {
	r = &CloudV3InferenceDeploymentService{}
	r.Options = opts
	return
}

// Create inference deployment
func (r *CloudV3InferenceDeploymentService) New(ctx context.Context, projectID int64, body CloudV3InferenceDeploymentNewParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get inference deployment
func (r *CloudV3InferenceDeploymentService) Get(ctx context.Context, projectID int64, instanceName string, opts ...option.RequestOption) (res *InferenceInstanceOut, err error) {
	opts = append(r.Options[:], opts...)
	if instanceName == "" {
		err = errors.New("missing required instance_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments/%s", projectID, instanceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update inference instance
func (r *CloudV3InferenceDeploymentService) Update(ctx context.Context, projectID int64, instanceName string, body CloudV3InferenceDeploymentUpdateParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if instanceName == "" {
		err = errors.New("missing required instance_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments/%s", projectID, instanceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List inference deployments
func (r *CloudV3InferenceDeploymentService) List(ctx context.Context, projectID int64, query CloudV3InferenceDeploymentListParams, opts ...option.RequestOption) (res *CloudV3InferenceDeploymentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete inference instance
func (r *CloudV3InferenceDeploymentService) Delete(ctx context.Context, projectID int64, instanceName string, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if instanceName == "" {
		err = errors.New("missing required instance_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments/%s", projectID, instanceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get inference deployment API key
func (r *CloudV3InferenceDeploymentService) GetApikey(ctx context.Context, projectID int64, instanceName string, opts ...option.RequestOption) (res *CloudV3InferenceDeploymentGetApikeyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if instanceName == "" {
		err = errors.New("missing required instance_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments/%s/apikey", projectID, instanceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get inference deployment logs
func (r *CloudV3InferenceDeploymentService) GetLogs(ctx context.Context, projectID int64, instanceName string, query CloudV3InferenceDeploymentGetLogsParams, opts ...option.RequestOption) (res *CloudV3InferenceDeploymentGetLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if instanceName == "" {
		err = errors.New("missing required instance_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments/%s/logs", projectID, instanceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This operation initializes an inference instance after it was stopped, making it
// available to handle inference requests again. The instance will launch with the
// **minimum** number of replicas defined in the scaling settings.
//
//   - If the minimum replicas are set to **0**, the instance will initially start
//     with **0** replicas.
//   - It will automatically scale up when it receives requests or SQS messages,
//     according to the configured scaling rules.
func (r *CloudV3InferenceDeploymentService) Start(ctx context.Context, projectID int64, instanceName string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if instanceName == "" {
		err = errors.New("missing required instance_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments/%s/start", projectID, instanceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// This operation shuts down an inference instance, making it unavailable for
// handling requests. The instance will scale down to **0** replicas, overriding
// any minimum replica settings.
//
//   - Once stopped, the instance will **not** process any inference requests or SQS
//     messages.
//   - It will **not** restart automatically and must be started manually.
//   - While stopped, the instance will **not** incur any charges.
func (r *CloudV3InferenceDeploymentService) Stop(ctx context.Context, projectID int64, instanceName string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if instanceName == "" {
		err = errors.New("missing required instance_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/%v/deployments/%s/stop", projectID, instanceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type ContainerScaleDeploymentV3 struct {
	// Maximum scale for the container
	Max int64 `json:"max,required"`
	// Minimum scale for the container
	Min int64 `json:"min,required"`
	// Cooldown period between scaling actions in seconds
	CooldownPeriod int64 `json:"cooldown_period,nullable"`
	// Polling interval for scaling triggers in seconds
	PollingInterval int64 `json:"polling_interval,nullable"`
	// Triggers for scaling actions
	Triggers ContainerScaleDeploymentV3Triggers `json:"triggers"`
	JSON     containerScaleDeploymentV3JSON     `json:"-"`
}

// containerScaleDeploymentV3JSON contains the JSON metadata for the struct
// [ContainerScaleDeploymentV3]
type containerScaleDeploymentV3JSON struct {
	Max             apijson.Field
	Min             apijson.Field
	CooldownPeriod  apijson.Field
	PollingInterval apijson.Field
	Triggers        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ContainerScaleDeploymentV3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerScaleDeploymentV3JSON) RawJSON() string {
	return r.raw
}

// Triggers for scaling actions
type ContainerScaleDeploymentV3Triggers struct {
	// CPU trigger configuration
	CPU ContainerScaleTriggersThreshold `json:"cpu,nullable"`
	// GPU memory trigger configuration. Calculated by DCGM_FI_DEV_MEM_COPY_UTIL metric
	GPUMemory ContainerScaleTriggersThreshold `json:"gpu_memory,nullable"`
	// GPU utilization trigger configuration. Calculated by DCGM_FI_DEV_GPU_UTIL metric
	GPUUtilization ContainerScaleTriggersThreshold `json:"gpu_utilization,nullable"`
	// HTTP trigger configuration
	HTTP ContainerScaleDeploymentV3TriggersHTTP `json:"http,nullable"`
	// Memory trigger configuration
	Memory ContainerScaleTriggersThreshold `json:"memory,nullable"`
	// SQS trigger configuration
	Sqs  ContainerScaleDeploymentV3TriggersSqs  `json:"sqs,nullable"`
	JSON containerScaleDeploymentV3TriggersJSON `json:"-"`
}

// containerScaleDeploymentV3TriggersJSON contains the JSON metadata for the struct
// [ContainerScaleDeploymentV3Triggers]
type containerScaleDeploymentV3TriggersJSON struct {
	CPU            apijson.Field
	GPUMemory      apijson.Field
	GPUUtilization apijson.Field
	HTTP           apijson.Field
	Memory         apijson.Field
	Sqs            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContainerScaleDeploymentV3Triggers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerScaleDeploymentV3TriggersJSON) RawJSON() string {
	return r.raw
}

// HTTP trigger configuration
type ContainerScaleDeploymentV3TriggersHTTP struct {
	// Request count per 'window' seconds for the http trigger
	Rate int64 `json:"rate,required"`
	// Time window for rate calculation in seconds
	Window int64                                      `json:"window,required"`
	JSON   containerScaleDeploymentV3TriggersHTTPJSON `json:"-"`
}

// containerScaleDeploymentV3TriggersHTTPJSON contains the JSON metadata for the
// struct [ContainerScaleDeploymentV3TriggersHTTP]
type containerScaleDeploymentV3TriggersHTTPJSON struct {
	Rate        apijson.Field
	Window      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContainerScaleDeploymentV3TriggersHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerScaleDeploymentV3TriggersHTTPJSON) RawJSON() string {
	return r.raw
}

// SQS trigger configuration
type ContainerScaleDeploymentV3TriggersSqs struct {
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
	AwsEndpoint string `json:"aws_endpoint,nullable"`
	// Scale on delayed messages
	ScaleOnDelayed bool `json:"scale_on_delayed"`
	// Scale on in-flight messages
	ScaleOnFlight bool                                      `json:"scale_on_flight"`
	JSON          containerScaleDeploymentV3TriggersSqsJSON `json:"-"`
}

// containerScaleDeploymentV3TriggersSqsJSON contains the JSON metadata for the
// struct [ContainerScaleDeploymentV3TriggersSqs]
type containerScaleDeploymentV3TriggersSqsJSON struct {
	ActivationQueueLength apijson.Field
	AwsRegion             apijson.Field
	QueueLength           apijson.Field
	QueueURL              apijson.Field
	SecretName            apijson.Field
	AwsEndpoint           apijson.Field
	ScaleOnDelayed        apijson.Field
	ScaleOnFlight         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContainerScaleDeploymentV3TriggersSqs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerScaleDeploymentV3TriggersSqsJSON) RawJSON() string {
	return r.raw
}

type ContainerScaleDeploymentV3Param struct {
	// Maximum scale for the container
	Max param.Field[int64] `json:"max,required"`
	// Minimum scale for the container
	Min param.Field[int64] `json:"min,required"`
	// Cooldown period between scaling actions in seconds
	CooldownPeriod param.Field[int64] `json:"cooldown_period"`
	// Polling interval for scaling triggers in seconds
	PollingInterval param.Field[int64] `json:"polling_interval"`
	// Triggers for scaling actions
	Triggers param.Field[ContainerScaleDeploymentV3TriggersParam] `json:"triggers"`
}

func (r ContainerScaleDeploymentV3Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Triggers for scaling actions
type ContainerScaleDeploymentV3TriggersParam struct {
	// CPU trigger configuration
	CPU param.Field[ContainerScaleTriggersThresholdParam] `json:"cpu"`
	// GPU memory trigger configuration. Calculated by DCGM_FI_DEV_MEM_COPY_UTIL metric
	GPUMemory param.Field[ContainerScaleTriggersThresholdParam] `json:"gpu_memory"`
	// GPU utilization trigger configuration. Calculated by DCGM_FI_DEV_GPU_UTIL metric
	GPUUtilization param.Field[ContainerScaleTriggersThresholdParam] `json:"gpu_utilization"`
	// HTTP trigger configuration
	HTTP param.Field[ContainerScaleDeploymentV3TriggersHTTPParam] `json:"http"`
	// Memory trigger configuration
	Memory param.Field[ContainerScaleTriggersThresholdParam] `json:"memory"`
	// SQS trigger configuration
	Sqs param.Field[ContainerScaleDeploymentV3TriggersSqsParam] `json:"sqs"`
}

func (r ContainerScaleDeploymentV3TriggersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// HTTP trigger configuration
type ContainerScaleDeploymentV3TriggersHTTPParam struct {
	// Request count per 'window' seconds for the http trigger
	Rate param.Field[int64] `json:"rate,required"`
	// Time window for rate calculation in seconds
	Window param.Field[int64] `json:"window,required"`
}

func (r ContainerScaleDeploymentV3TriggersHTTPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SQS trigger configuration
type ContainerScaleDeploymentV3TriggersSqsParam struct {
	// Number of messages for activation
	ActivationQueueLength param.Field[int64] `json:"activation_queue_length,required"`
	// AWS region
	AwsRegion param.Field[string] `json:"aws_region,required"`
	// Number of messages for one replica
	QueueLength param.Field[int64] `json:"queue_length,required"`
	// SQS queue URL
	QueueURL param.Field[string] `json:"queue_url,required"`
	// Auth secret name
	SecretName param.Field[string] `json:"secret_name,required"`
	// Custom AWS endpoint
	AwsEndpoint param.Field[string] `json:"aws_endpoint"`
	// Scale on delayed messages
	ScaleOnDelayed param.Field[bool] `json:"scale_on_delayed"`
	// Scale on in-flight messages
	ScaleOnFlight param.Field[bool] `json:"scale_on_flight"`
}

func (r ContainerScaleDeploymentV3TriggersSqsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContainerScaleTriggersThreshold struct {
	// Threshold value for the trigger in percentage
	Threshold int64                               `json:"threshold,required"`
	JSON      containerScaleTriggersThresholdJSON `json:"-"`
}

// containerScaleTriggersThresholdJSON contains the JSON metadata for the struct
// [ContainerScaleTriggersThreshold]
type containerScaleTriggersThresholdJSON struct {
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContainerScaleTriggersThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerScaleTriggersThresholdJSON) RawJSON() string {
	return r.raw
}

type ContainerScaleTriggersThresholdParam struct {
	// Threshold value for the trigger in percentage
	Threshold param.Field[int64] `json:"threshold,required"`
}

func (r ContainerScaleTriggersThresholdParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InferenceInstanceOut struct {
	// List of containers for the inference instance
	Containers []InferenceInstanceOutContainer `json:"containers,required"`
	// Registry credentials name
	CredentialsName string `json:"credentials_name,required"`
	// Flavor name for the inference instance
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
	// Project ID. If not provided, your default project ID will be used.
	ProjectID int64 `json:"project_id,required"`
	// Inference instance status
	Status InstanceStatus `json:"status,required"`
	// Address of the inference instance
	Address string `json:"address,nullable" format:"uri"`
	// `true` if instance uses API key authentication. `"X-Api-Key": "*****"` header is
	// required for the requests to the instance if enabled.
	AuthEnabled bool `json:"auth_enabled"`
	// Command to be executed when running a container from an image.
	Command string `json:"command,nullable"`
	// Inference instance creation date in ISO 8601 format.
	CreatedAt string `json:"created_at,nullable"`
	// Inference instance description.
	Description string `json:"description"`
	// Environment variables for the inference instance
	Envs map[string]string `json:"envs,nullable"`
	// Ingress options for the inference instance
	IngressOpts IngressOpts `json:"ingress_opts,nullable"`
	// Logging configuration for the inference instance
	Logging InferenceInstanceOutLogging `json:"logging,nullable"`
	// Probes configured for all containers of the inference instance.
	Probes InstanceProbesDeploymentV2 `json:"probes,nullable"`
	// Specifies the duration in seconds without any requests after which the
	// containers will be downscaled to their minimum scale value as defined by
	// `scale.min`. If set, this helps in optimizing resource usage by reducing the
	// number of container instances during periods of inactivity.
	Timeout int64                    `json:"timeout,nullable"`
	JSON    inferenceInstanceOutJSON `json:"-"`
}

// inferenceInstanceOutJSON contains the JSON metadata for the struct
// [InferenceInstanceOut]
type inferenceInstanceOutJSON struct {
	Containers      apijson.Field
	CredentialsName apijson.Field
	FlavorName      apijson.Field
	Image           apijson.Field
	ListeningPort   apijson.Field
	Name            apijson.Field
	ProjectID       apijson.Field
	Status          apijson.Field
	Address         apijson.Field
	AuthEnabled     apijson.Field
	Command         apijson.Field
	CreatedAt       apijson.Field
	Description     apijson.Field
	Envs            apijson.Field
	IngressOpts     apijson.Field
	Logging         apijson.Field
	Probes          apijson.Field
	Timeout         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InferenceInstanceOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferenceInstanceOutJSON) RawJSON() string {
	return r.raw
}

type InferenceInstanceOutContainer struct {
	// Status of the containers deployment
	DeployStatus DeployStatus `json:"deploy_status,required"`
	// Region name for the container
	RegionID int64 `json:"region_id,required"`
	// Scale for the container
	Scale ContainerScaleDeploymentV3 `json:"scale,required"`
	// Address of the inference instance
	Address string `json:"address,nullable" format:"uri"`
	// Error message if the container deployment failed
	ErrorMessage string                            `json:"error_message,nullable"`
	JSON         inferenceInstanceOutContainerJSON `json:"-"`
}

// inferenceInstanceOutContainerJSON contains the JSON metadata for the struct
// [InferenceInstanceOutContainer]
type inferenceInstanceOutContainerJSON struct {
	DeployStatus apijson.Field
	RegionID     apijson.Field
	Scale        apijson.Field
	Address      apijson.Field
	ErrorMessage apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InferenceInstanceOutContainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferenceInstanceOutContainerJSON) RawJSON() string {
	return r.raw
}

// Logging configuration for the inference instance
type InferenceInstanceOutLogging struct {
	Enabled             bool                             `json:"enabled,required"`
	DestinationRegionID int64                            `json:"destination_region_id,nullable"`
	RetentionPolicy     LaasIndexRetentionPolicyPydantic `json:"retention_policy,nullable"`
	TopicName           string                           `json:"topic_name,nullable"`
	JSON                inferenceInstanceOutLoggingJSON  `json:"-"`
}

// inferenceInstanceOutLoggingJSON contains the JSON metadata for the struct
// [InferenceInstanceOutLogging]
type inferenceInstanceOutLoggingJSON struct {
	Enabled             apijson.Field
	DestinationRegionID apijson.Field
	RetentionPolicy     apijson.Field
	TopicName           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InferenceInstanceOutLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inferenceInstanceOutLoggingJSON) RawJSON() string {
	return r.raw
}

type IngressOpts struct {
	// Disable response buffering if true. A client usually has a much slower
	// connection and can not consume the response data as fast as it is produced by an
	// upstream application. Ingress tries to buffer the whole response in order to
	// release the upstream application as soon as possible.By default, the response
	// buffering is enabled.
	DisableResponseBuffering bool            `json:"disable_response_buffering"`
	JSON                     ingressOptsJSON `json:"-"`
}

// ingressOptsJSON contains the JSON metadata for the struct [IngressOpts]
type ingressOptsJSON struct {
	DisableResponseBuffering apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *IngressOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ingressOptsJSON) RawJSON() string {
	return r.raw
}

type IngressOptsParam struct {
	// Disable response buffering if true. A client usually has a much slower
	// connection and can not consume the response data as fast as it is produced by an
	// upstream application. Ingress tries to buffer the whole response in order to
	// release the upstream application as soon as possible.By default, the response
	// buffering is enabled.
	DisableResponseBuffering param.Field[bool] `json:"disable_response_buffering"`
}

func (r IngressOptsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceContainerProbeConfiguration struct {
	// Whether the probe is enabled or not.
	Enabled bool `json:"enabled,required"`
	// Probe configuration (exec, http_get or tcp_socket)
	Probe InstanceContainerProbeConfigurationProbe `json:"probe,nullable"`
	JSON  instanceContainerProbeConfigurationJSON  `json:"-"`
}

// instanceContainerProbeConfigurationJSON contains the JSON metadata for the
// struct [InstanceContainerProbeConfiguration]
type instanceContainerProbeConfigurationJSON struct {
	Enabled     apijson.Field
	Probe       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceContainerProbeConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceContainerProbeConfigurationJSON) RawJSON() string {
	return r.raw
}

// Probe configuration (exec, http_get or tcp_socket)
type InstanceContainerProbeConfigurationProbe struct {
	// Exec probe configuration
	Exec InstanceContainerProbeConfigurationProbeExec `json:"exec,nullable"`
	// The number of consecutive probe failures that mark the container as unhealthy.
	FailureThreshold int64 `json:"failure_threshold"`
	// HTTP GET probe configuration
	HTTPGet InstanceContainerProbeConfigurationProbeHTTPGet `json:"http_get,nullable"`
	// The initial delay before starting the first probe.
	InitialDelaySeconds int64 `json:"initial_delay_seconds"`
	// How often (in seconds) to perform the probe.
	PeriodSeconds int64 `json:"period_seconds"`
	// The number of consecutive successful probes that mark the container as healthy.
	SuccessThreshold int64 `json:"success_threshold"`
	// TCP socket probe configuration
	TcpSocket InstanceContainerProbeConfigurationProbeTcpSocket `json:"tcp_socket,nullable"`
	// The timeout for each probe.
	TimeoutSeconds int64                                        `json:"timeout_seconds"`
	JSON           instanceContainerProbeConfigurationProbeJSON `json:"-"`
}

// instanceContainerProbeConfigurationProbeJSON contains the JSON metadata for the
// struct [InstanceContainerProbeConfigurationProbe]
type instanceContainerProbeConfigurationProbeJSON struct {
	Exec                apijson.Field
	FailureThreshold    apijson.Field
	HTTPGet             apijson.Field
	InitialDelaySeconds apijson.Field
	PeriodSeconds       apijson.Field
	SuccessThreshold    apijson.Field
	TcpSocket           apijson.Field
	TimeoutSeconds      apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InstanceContainerProbeConfigurationProbe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceContainerProbeConfigurationProbeJSON) RawJSON() string {
	return r.raw
}

// Exec probe configuration
type InstanceContainerProbeConfigurationProbeExec struct {
	// Command to be executed inside the running container.
	Command []string                                         `json:"command,required"`
	JSON    instanceContainerProbeConfigurationProbeExecJSON `json:"-"`
}

// instanceContainerProbeConfigurationProbeExecJSON contains the JSON metadata for
// the struct [InstanceContainerProbeConfigurationProbeExec]
type instanceContainerProbeConfigurationProbeExecJSON struct {
	Command     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceContainerProbeConfigurationProbeExec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceContainerProbeConfigurationProbeExecJSON) RawJSON() string {
	return r.raw
}

// HTTP GET probe configuration
type InstanceContainerProbeConfigurationProbeHTTPGet struct {
	// Port number the probe should connect to.
	Port int64 `json:"port,required"`
	// HTTP headers to be sent with the request.
	Headers map[string]string `json:"headers"`
	// Host name to send HTTP request to.
	Host string `json:"host,nullable"`
	// The endpoint to send the HTTP request to.
	Path string `json:"path"`
	// Schema to use for the HTTP request.
	Schema string                                              `json:"schema"`
	JSON   instanceContainerProbeConfigurationProbeHTTPGetJSON `json:"-"`
}

// instanceContainerProbeConfigurationProbeHTTPGetJSON contains the JSON metadata
// for the struct [InstanceContainerProbeConfigurationProbeHTTPGet]
type instanceContainerProbeConfigurationProbeHTTPGetJSON struct {
	Port        apijson.Field
	Headers     apijson.Field
	Host        apijson.Field
	Path        apijson.Field
	Schema      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceContainerProbeConfigurationProbeHTTPGet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceContainerProbeConfigurationProbeHTTPGetJSON) RawJSON() string {
	return r.raw
}

// TCP socket probe configuration
type InstanceContainerProbeConfigurationProbeTcpSocket struct {
	// Port number to check if it's open.
	Port int64                                                 `json:"port,required"`
	JSON instanceContainerProbeConfigurationProbeTcpSocketJSON `json:"-"`
}

// instanceContainerProbeConfigurationProbeTcpSocketJSON contains the JSON metadata
// for the struct [InstanceContainerProbeConfigurationProbeTcpSocket]
type instanceContainerProbeConfigurationProbeTcpSocketJSON struct {
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceContainerProbeConfigurationProbeTcpSocket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceContainerProbeConfigurationProbeTcpSocketJSON) RawJSON() string {
	return r.raw
}

type InstanceContainerProbeConfigurationParam struct {
	// Whether the probe is enabled or not.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Probe configuration (exec, http_get or tcp_socket)
	Probe param.Field[InstanceContainerProbeConfigurationProbeParam] `json:"probe"`
}

func (r InstanceContainerProbeConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Probe configuration (exec, http_get or tcp_socket)
type InstanceContainerProbeConfigurationProbeParam struct {
	// Exec probe configuration
	Exec param.Field[InstanceContainerProbeConfigurationProbeExecParam] `json:"exec"`
	// The number of consecutive probe failures that mark the container as unhealthy.
	FailureThreshold param.Field[int64] `json:"failure_threshold"`
	// HTTP GET probe configuration
	HTTPGet param.Field[InstanceContainerProbeConfigurationProbeHTTPGetParam] `json:"http_get"`
	// The initial delay before starting the first probe.
	InitialDelaySeconds param.Field[int64] `json:"initial_delay_seconds"`
	// How often (in seconds) to perform the probe.
	PeriodSeconds param.Field[int64] `json:"period_seconds"`
	// The number of consecutive successful probes that mark the container as healthy.
	SuccessThreshold param.Field[int64] `json:"success_threshold"`
	// TCP socket probe configuration
	TcpSocket param.Field[InstanceContainerProbeConfigurationProbeTcpSocketParam] `json:"tcp_socket"`
	// The timeout for each probe.
	TimeoutSeconds param.Field[int64] `json:"timeout_seconds"`
}

func (r InstanceContainerProbeConfigurationProbeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Exec probe configuration
type InstanceContainerProbeConfigurationProbeExecParam struct {
	// Command to be executed inside the running container.
	Command param.Field[[]string] `json:"command,required"`
}

func (r InstanceContainerProbeConfigurationProbeExecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// HTTP GET probe configuration
type InstanceContainerProbeConfigurationProbeHTTPGetParam struct {
	// Port number the probe should connect to.
	Port param.Field[int64] `json:"port,required"`
	// HTTP headers to be sent with the request.
	Headers param.Field[map[string]string] `json:"headers"`
	// Host name to send HTTP request to.
	Host param.Field[string] `json:"host"`
	// The endpoint to send the HTTP request to.
	Path param.Field[string] `json:"path"`
	// Schema to use for the HTTP request.
	Schema param.Field[string] `json:"schema"`
}

func (r InstanceContainerProbeConfigurationProbeHTTPGetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TCP socket probe configuration
type InstanceContainerProbeConfigurationProbeTcpSocketParam struct {
	// Port number to check if it's open.
	Port param.Field[int64] `json:"port,required"`
}

func (r InstanceContainerProbeConfigurationProbeTcpSocketParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceProbesDeploymentV2 struct {
	// Liveness probe configuration
	LivenessProbe InstanceContainerProbeConfiguration `json:"liveness_probe,nullable"`
	// Readiness probe configuration
	ReadinessProbe InstanceContainerProbeConfiguration `json:"readiness_probe,nullable"`
	// Startup probe configuration
	StartupProbe InstanceContainerProbeConfiguration `json:"startup_probe,nullable"`
	JSON         instanceProbesDeploymentV2JSON      `json:"-"`
}

// instanceProbesDeploymentV2JSON contains the JSON metadata for the struct
// [InstanceProbesDeploymentV2]
type instanceProbesDeploymentV2JSON struct {
	LivenessProbe  apijson.Field
	ReadinessProbe apijson.Field
	StartupProbe   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceProbesDeploymentV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceProbesDeploymentV2JSON) RawJSON() string {
	return r.raw
}

type InstanceProbesDeploymentV2Param struct {
	// Liveness probe configuration
	LivenessProbe param.Field[InstanceContainerProbeConfigurationParam] `json:"liveness_probe"`
	// Readiness probe configuration
	ReadinessProbe param.Field[InstanceContainerProbeConfigurationParam] `json:"readiness_probe"`
	// Startup probe configuration
	StartupProbe param.Field[InstanceContainerProbeConfigurationParam] `json:"startup_probe"`
}

func (r InstanceProbesDeploymentV2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoggingInParam struct {
	// ID of the region in which the logs will be stored
	DestinationRegionID param.Field[int64] `json:"destination_region_id"`
	// Enable or disable log streaming
	Enabled param.Field[bool] `json:"enabled"`
	// Logs retention policy
	RetentionPolicy param.Field[LaasIndexRetentionPolicyPydanticParam] `json:"retention_policy"`
	// The topic name to stream logs to
	TopicName param.Field[string] `json:"topic_name"`
}

func (r LoggingInParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV3InferenceDeploymentListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []InferenceInstanceOut                     `json:"results,required"`
	JSON    cloudV3InferenceDeploymentListResponseJSON `json:"-"`
}

// cloudV3InferenceDeploymentListResponseJSON contains the JSON metadata for the
// struct [CloudV3InferenceDeploymentListResponse]
type cloudV3InferenceDeploymentListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV3InferenceDeploymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3InferenceDeploymentListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV3InferenceDeploymentGetApikeyResponse struct {
	// API key secret
	Secret string `json:"secret,required"`
	// API key status
	Status CloudV3InferenceDeploymentGetApikeyResponseStatus `json:"status,required"`
	JSON   cloudV3InferenceDeploymentGetApikeyResponseJSON   `json:"-"`
}

// cloudV3InferenceDeploymentGetApikeyResponseJSON contains the JSON metadata for
// the struct [CloudV3InferenceDeploymentGetApikeyResponse]
type cloudV3InferenceDeploymentGetApikeyResponseJSON struct {
	Secret      apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV3InferenceDeploymentGetApikeyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3InferenceDeploymentGetApikeyResponseJSON) RawJSON() string {
	return r.raw
}

// API key status
type CloudV3InferenceDeploymentGetApikeyResponseStatus string

const (
	CloudV3InferenceDeploymentGetApikeyResponseStatusPending CloudV3InferenceDeploymentGetApikeyResponseStatus = "PENDING"
	CloudV3InferenceDeploymentGetApikeyResponseStatusReady   CloudV3InferenceDeploymentGetApikeyResponseStatus = "READY"
)

func (r CloudV3InferenceDeploymentGetApikeyResponseStatus) IsKnown() bool {
	switch r {
	case CloudV3InferenceDeploymentGetApikeyResponseStatusPending, CloudV3InferenceDeploymentGetApikeyResponseStatusReady:
		return true
	}
	return false
}

type CloudV3InferenceDeploymentGetLogsResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV3InferenceDeploymentGetLogsResponseResult `json:"results,required"`
	JSON    cloudV3InferenceDeploymentGetLogsResponseJSON     `json:"-"`
}

// cloudV3InferenceDeploymentGetLogsResponseJSON contains the JSON metadata for the
// struct [CloudV3InferenceDeploymentGetLogsResponse]
type cloudV3InferenceDeploymentGetLogsResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV3InferenceDeploymentGetLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3InferenceDeploymentGetLogsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV3InferenceDeploymentGetLogsResponseResult struct {
	// Log message.
	Message string `json:"message,required"`
	// Pod name.
	Pod string `json:"pod,required"`
	// Region ID where the container is deployed.
	RegionID int64 `json:"region_id,required"`
	// Log message timestamp in ISO 8601 format.
	Time time.Time                                           `json:"time,required" format:"date-time"`
	JSON cloudV3InferenceDeploymentGetLogsResponseResultJSON `json:"-"`
}

// cloudV3InferenceDeploymentGetLogsResponseResultJSON contains the JSON metadata
// for the struct [CloudV3InferenceDeploymentGetLogsResponseResult]
type cloudV3InferenceDeploymentGetLogsResponseResultJSON struct {
	Message     apijson.Field
	Pod         apijson.Field
	RegionID    apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV3InferenceDeploymentGetLogsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3InferenceDeploymentGetLogsResponseResultJSON) RawJSON() string {
	return r.raw
}

type CloudV3InferenceDeploymentNewParams struct {
	// List of containers for the inference instance.
	Containers param.Field[[]CloudV3InferenceDeploymentNewParamsContainer] `json:"containers,required"`
	// Flavor name for the inference instance.
	FlavorName param.Field[string] `json:"flavor_name,required"`
	// Docker image for the inference instance. This field should contain the image
	// name and tag in the format 'name:tag', e.g., 'nginx:latest'. It defaults to
	// Docker Hub as the image registry, but any accessible Docker image URL can be
	// specified.
	Image param.Field[string] `json:"image,required"`
	// Listening port for the inference instance.
	ListeningPort param.Field[int64] `json:"listening_port,required"`
	// Inference instance name.
	Name param.Field[string] `json:"name,required"`
	// Set to `true` to enable API key authentication for the inference instance.
	// `"X-Api-Key": "*****"` header is required for the requests to the instance if
	// enabled.
	AuthEnabled param.Field[bool] `json:"auth_enabled"`
	// Command to be executed when running a container from an image.
	Command param.Field[[]string] `json:"command"`
	// Registry credentials name
	CredentialsName param.Field[string] `json:"credentials_name"`
	// Inference instance description.
	Description param.Field[string] `json:"description"`
	// Environment variables for the inference instance.
	Envs param.Field[map[string]string] `json:"envs"`
	// Ingress options for the inference instance
	IngressOpts param.Field[IngressOptsParam] `json:"ingress_opts"`
	// Logging configuration for the inference instance
	Logging param.Field[LoggingInParam] `json:"logging"`
	// Probes configured for all containers of the inference instance. If probes are
	// not provided, and the image_name is from a the Model Catalog registry, the
	// default probes will be used.
	Probes param.Field[InstanceProbesDeploymentV2Param] `json:"probes"`
	// Specifies the duration in seconds without any requests after which the
	// containers will be downscaled to their minimum scale value as defined by
	// `scale.min`. If set, this helps in optimizing resource usage by reducing the
	// number of container instances during periods of inactivity. The default value
	// when the parameter is not set is 120.
	Timeout param.Field[int64] `json:"timeout"`
}

func (r CloudV3InferenceDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV3InferenceDeploymentNewParamsContainer struct {
	// Region id for the container
	RegionID param.Field[int64] `json:"region_id,required"`
	// Scale for the container
	Scale param.Field[ContainerScaleDeploymentV3Param] `json:"scale,required"`
}

func (r CloudV3InferenceDeploymentNewParamsContainer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV3InferenceDeploymentUpdateParams struct {
	// Set to `true` to enable API key authentication for the inference instance.
	// `"X-Api-Key": "*****"` header is required for the requests to the instance if
	// enabled.
	AuthEnabled param.Field[bool] `json:"auth_enabled"`
	// Command to be executed when running a container from an image.
	Command param.Field[[]string] `json:"command"`
	// List of containers for the inference instance.
	Containers param.Field[[]CloudV3InferenceDeploymentUpdateParamsContainer] `json:"containers"`
	// Registry credentials name
	CredentialsName param.Field[string] `json:"credentials_name"`
	// Inference instance description.
	Description param.Field[string] `json:"description"`
	// Environment variables for the inference instance.
	Envs param.Field[map[string]string] `json:"envs"`
	// Flavor name for the inference instance.
	FlavorName param.Field[string] `json:"flavor_name"`
	// Docker image for the inference instance. This field should contain the image
	// name and tag in the format 'name:tag', e.g., 'nginx:latest'. It defaults to
	// Docker Hub as the image registry, but any accessible Docker image URL can be
	// specified.
	Image param.Field[string] `json:"image"`
	// Ingress options for the inference instance
	IngressOpts param.Field[IngressOptsParam] `json:"ingress_opts"`
	// Listening port for the inference instance.
	ListeningPort param.Field[int64] `json:"listening_port"`
	// Logging configuration for the inference instance
	Logging param.Field[LoggingInParam] `json:"logging"`
	// Probes configured for all containers of the inference instance.
	Probes param.Field[InstanceProbesDeploymentV2Param] `json:"probes"`
	// Specifies the duration in seconds without any requests after which the
	// containers will be downscaled to their minimum scale value as defined by
	// `scale.min`. If set, this helps in optimizing resource usage by reducing the
	// number of container instances during periods of inactivity. The default value
	// when the parameter is not set is 120.
	Timeout param.Field[int64] `json:"timeout"`
}

func (r CloudV3InferenceDeploymentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV3InferenceDeploymentUpdateParamsContainer struct {
	// Region id for the container
	RegionID param.Field[int64] `json:"region_id,required"`
	// Scale for the container
	Scale param.Field[ContainerScaleDeploymentV3Param] `json:"scale,required"`
}

func (r CloudV3InferenceDeploymentUpdateParamsContainer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV3InferenceDeploymentListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV3InferenceDeploymentListParams]'s query parameters as
// `url.Values`.
func (r CloudV3InferenceDeploymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV3InferenceDeploymentGetLogsParams struct {
	// Optional. Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Field[int64] `query:"offset"`
	// Order by field
	OrderBy param.Field[CloudV3InferenceDeploymentGetLogsParamsOrderBy] `query:"order_by"`
	// Region ID
	RegionID param.Field[int64] `query:"region_id"`
}

// URLQuery serializes [CloudV3InferenceDeploymentGetLogsParams]'s query parameters
// as `url.Values`.
func (r CloudV3InferenceDeploymentGetLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Order by field
type CloudV3InferenceDeploymentGetLogsParamsOrderBy string

const (
	CloudV3InferenceDeploymentGetLogsParamsOrderByTimeAsc  CloudV3InferenceDeploymentGetLogsParamsOrderBy = "time.asc"
	CloudV3InferenceDeploymentGetLogsParamsOrderByTimeDesc CloudV3InferenceDeploymentGetLogsParamsOrderBy = "time.desc"
)

func (r CloudV3InferenceDeploymentGetLogsParamsOrderBy) IsKnown() bool {
	switch r {
	case CloudV3InferenceDeploymentGetLogsParamsOrderByTimeAsc, CloudV3InferenceDeploymentGetLogsParamsOrderByTimeDesc:
		return true
	}
	return false
}
