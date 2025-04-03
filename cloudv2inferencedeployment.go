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

// CloudV2InferenceDeploymentService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2InferenceDeploymentService] method instead.
type CloudV2InferenceDeploymentService struct {
	Options []option.RequestOption
	Regions *CloudV2InferenceDeploymentRegionService
}

// NewCloudV2InferenceDeploymentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV2InferenceDeploymentService(opts ...option.RequestOption) (r *CloudV2InferenceDeploymentService) {
	r = &CloudV2InferenceDeploymentService{}
	r.Options = opts
	r.Regions = NewCloudV2InferenceDeploymentRegionService(opts...)
	return
}

// Deprecated. Create Inference Instance
func (r *CloudV2InferenceDeploymentService) New(ctx context.Context, body CloudV2InferenceDeploymentNewParams, opts ...option.RequestOption) (res *CloudV2InferenceDeploymentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/inference/deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deprecated. Get Inference Instance
func (r *CloudV2InferenceDeploymentService) Get(ctx context.Context, instanceID string, opts ...option.RequestOption) (res *InstanceOut, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/deployments/%s", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deprecated. Update Inference Instance
func (r *CloudV2InferenceDeploymentService) Update(ctx context.Context, instanceID string, body CloudV2InferenceDeploymentUpdateParams, opts ...option.RequestOption) (res *InstanceOut, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/deployments/%s", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deprecated. List Inference Instances
func (r *CloudV2InferenceDeploymentService) List(ctx context.Context, query CloudV2InferenceDeploymentListParams, opts ...option.RequestOption) (res *CloudV2InferenceDeploymentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/inference/deployments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deprecated. Delete Inference Instance
func (r *CloudV2InferenceDeploymentService) Delete(ctx context.Context, instanceID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/deployments/%s", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Deprecated. Get Inference Instance Logs
func (r *CloudV2InferenceDeploymentService) GetLogs(ctx context.Context, instanceID string, opts ...option.RequestOption) (res *CloudV2InferenceDeploymentGetLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/deployments/%s/logs", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deprecated. Start Inference Instance
func (r *CloudV2InferenceDeploymentService) Start(ctx context.Context, instanceID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/deployments/%s/start", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Deprecated. Stop Inference Instance
func (r *CloudV2InferenceDeploymentService) Stop(ctx context.Context, instanceID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/deployments/%s/stop", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type ContainerOut struct {
	// Status of the containers deployment
	DeployStatus DeployStatus `json:"deploy_status,required"`
	// Region ID for the container
	RegionID int64 `json:"region_id,required"`
	// Scale for the container
	Scale ContainerScaleDeployment `json:"scale,required"`
	// Error message if the container deployment failed
	ErrorMessage string           `json:"error_message,nullable"`
	JSON         containerOutJSON `json:"-"`
}

// containerOutJSON contains the JSON metadata for the struct [ContainerOut]
type containerOutJSON struct {
	DeployStatus apijson.Field
	RegionID     apijson.Field
	Scale        apijson.Field
	ErrorMessage apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContainerOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerOutJSON) RawJSON() string {
	return r.raw
}

type ContainerProbeConfiguration struct {
	// Whether the probe is enabled or not.
	Enabled bool `json:"enabled,required"`
	// Probe configuration (exec, http_get or tcp_socket)
	Probe ContainerProbeConfigurationProbe `json:"probe,nullable"`
	JSON  containerProbeConfigurationJSON  `json:"-"`
}

// containerProbeConfigurationJSON contains the JSON metadata for the struct
// [ContainerProbeConfiguration]
type containerProbeConfigurationJSON struct {
	Enabled     apijson.Field
	Probe       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContainerProbeConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerProbeConfigurationJSON) RawJSON() string {
	return r.raw
}

// Probe configuration (exec, http_get or tcp_socket)
type ContainerProbeConfigurationProbe struct {
	// Exec probe configuration
	Exec ContainerProbeConfigurationProbeExec `json:"exec,nullable"`
	// The number of consecutive probe failures that mark the container as unhealthy.
	FailureThreshold int64 `json:"failure_threshold"`
	// HTTP GET probe configuration
	HTTPGet ContainerProbeConfigurationProbeHTTPGet `json:"http_get,nullable"`
	// The initial delay before starting the first probe.
	InitialDelaySeconds int64 `json:"initial_delay_seconds"`
	// How often (in seconds) to perform the probe.
	PeriodSeconds int64 `json:"period_seconds"`
	// The number of consecutive successful probes that mark the container as healthy.
	SuccessThreshold int64 `json:"success_threshold"`
	// TCP socket probe configuration
	TcpSocket ContainerProbeConfigurationProbeTcpSocket `json:"tcp_socket,nullable"`
	// The timeout for each probe.
	TimeoutSeconds int64                                `json:"timeout_seconds"`
	JSON           containerProbeConfigurationProbeJSON `json:"-"`
}

// containerProbeConfigurationProbeJSON contains the JSON metadata for the struct
// [ContainerProbeConfigurationProbe]
type containerProbeConfigurationProbeJSON struct {
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

func (r *ContainerProbeConfigurationProbe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerProbeConfigurationProbeJSON) RawJSON() string {
	return r.raw
}

// Exec probe configuration
type ContainerProbeConfigurationProbeExec struct {
	// Command to be executed inside the running container.
	Command []string                                 `json:"command,required"`
	JSON    containerProbeConfigurationProbeExecJSON `json:"-"`
}

// containerProbeConfigurationProbeExecJSON contains the JSON metadata for the
// struct [ContainerProbeConfigurationProbeExec]
type containerProbeConfigurationProbeExecJSON struct {
	Command     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContainerProbeConfigurationProbeExec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerProbeConfigurationProbeExecJSON) RawJSON() string {
	return r.raw
}

// HTTP GET probe configuration
type ContainerProbeConfigurationProbeHTTPGet struct {
	// Port number the probe should connect to.
	Port int64 `json:"port,required"`
	// HTTP headers to be sent with the request.
	Headers map[string]string `json:"headers"`
	// Host name to send HTTP request to.
	Host string `json:"host,nullable"`
	// The endpoint to send the HTTP request to.
	Path string `json:"path"`
	// Schema to use for the HTTP request.
	Schema string                                      `json:"schema"`
	JSON   containerProbeConfigurationProbeHTTPGetJSON `json:"-"`
}

// containerProbeConfigurationProbeHTTPGetJSON contains the JSON metadata for the
// struct [ContainerProbeConfigurationProbeHTTPGet]
type containerProbeConfigurationProbeHTTPGetJSON struct {
	Port        apijson.Field
	Headers     apijson.Field
	Host        apijson.Field
	Path        apijson.Field
	Schema      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContainerProbeConfigurationProbeHTTPGet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerProbeConfigurationProbeHTTPGetJSON) RawJSON() string {
	return r.raw
}

// TCP socket probe configuration
type ContainerProbeConfigurationProbeTcpSocket struct {
	// Port number to check if it's open.
	Port int64                                         `json:"port,required"`
	JSON containerProbeConfigurationProbeTcpSocketJSON `json:"-"`
}

// containerProbeConfigurationProbeTcpSocketJSON contains the JSON metadata for the
// struct [ContainerProbeConfigurationProbeTcpSocket]
type containerProbeConfigurationProbeTcpSocketJSON struct {
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContainerProbeConfigurationProbeTcpSocket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerProbeConfigurationProbeTcpSocketJSON) RawJSON() string {
	return r.raw
}

type ContainerProbeConfigurationParam struct {
	// Whether the probe is enabled or not.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Probe configuration (exec, http_get or tcp_socket)
	Probe param.Field[ContainerProbeConfigurationProbeParam] `json:"probe"`
}

func (r ContainerProbeConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Probe configuration (exec, http_get or tcp_socket)
type ContainerProbeConfigurationProbeParam struct {
	// Exec probe configuration
	Exec param.Field[ContainerProbeConfigurationProbeExecParam] `json:"exec"`
	// The number of consecutive probe failures that mark the container as unhealthy.
	FailureThreshold param.Field[int64] `json:"failure_threshold"`
	// HTTP GET probe configuration
	HTTPGet param.Field[ContainerProbeConfigurationProbeHTTPGetParam] `json:"http_get"`
	// The initial delay before starting the first probe.
	InitialDelaySeconds param.Field[int64] `json:"initial_delay_seconds"`
	// How often (in seconds) to perform the probe.
	PeriodSeconds param.Field[int64] `json:"period_seconds"`
	// The number of consecutive successful probes that mark the container as healthy.
	SuccessThreshold param.Field[int64] `json:"success_threshold"`
	// TCP socket probe configuration
	TcpSocket param.Field[ContainerProbeConfigurationProbeTcpSocketParam] `json:"tcp_socket"`
	// The timeout for each probe.
	TimeoutSeconds param.Field[int64] `json:"timeout_seconds"`
}

func (r ContainerProbeConfigurationProbeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Exec probe configuration
type ContainerProbeConfigurationProbeExecParam struct {
	// Command to be executed inside the running container.
	Command param.Field[[]string] `json:"command,required"`
}

func (r ContainerProbeConfigurationProbeExecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// HTTP GET probe configuration
type ContainerProbeConfigurationProbeHTTPGetParam struct {
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

func (r ContainerProbeConfigurationProbeHTTPGetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TCP socket probe configuration
type ContainerProbeConfigurationProbeTcpSocketParam struct {
	// Port number to check if it's open.
	Port param.Field[int64] `json:"port,required"`
}

func (r ContainerProbeConfigurationProbeTcpSocketParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContainerScaleDeployment struct {
	// Maximum scale for the container
	Max int64 `json:"max,required"`
	// Minimum scale for the container
	Min int64 `json:"min,required"`
	// Cooldown period between scaling actions in seconds
	CooldownPeriod int64 `json:"cooldown_period,nullable"`
	// Triggers for scaling actions
	Triggers ContainerScaleDeploymentTriggers `json:"triggers,nullable"`
	JSON     containerScaleDeploymentJSON     `json:"-"`
}

// containerScaleDeploymentJSON contains the JSON metadata for the struct
// [ContainerScaleDeployment]
type containerScaleDeploymentJSON struct {
	Max            apijson.Field
	Min            apijson.Field
	CooldownPeriod apijson.Field
	Triggers       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContainerScaleDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerScaleDeploymentJSON) RawJSON() string {
	return r.raw
}

// Triggers for scaling actions
type ContainerScaleDeploymentTriggers struct {
	CPU            Threshold                            `json:"cpu,nullable"`
	GPUMemory      Threshold                            `json:"gpu_memory,nullable"`
	GPUUtilization Threshold                            `json:"gpu_utilization,nullable"`
	HTTP           ContainerScaleDeploymentTriggersHTTP `json:"http,nullable"`
	Memory         Threshold                            `json:"memory,nullable"`
	JSON           containerScaleDeploymentTriggersJSON `json:"-"`
}

// containerScaleDeploymentTriggersJSON contains the JSON metadata for the struct
// [ContainerScaleDeploymentTriggers]
type containerScaleDeploymentTriggersJSON struct {
	CPU            apijson.Field
	GPUMemory      apijson.Field
	GPUUtilization apijson.Field
	HTTP           apijson.Field
	Memory         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContainerScaleDeploymentTriggers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerScaleDeploymentTriggersJSON) RawJSON() string {
	return r.raw
}

type ContainerScaleDeploymentTriggersHTTP struct {
	Rate   int64                                    `json:"rate,required"`
	Window int64                                    `json:"window,required"`
	JSON   containerScaleDeploymentTriggersHTTPJSON `json:"-"`
}

// containerScaleDeploymentTriggersHTTPJSON contains the JSON metadata for the
// struct [ContainerScaleDeploymentTriggersHTTP]
type containerScaleDeploymentTriggersHTTPJSON struct {
	Rate        apijson.Field
	Window      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContainerScaleDeploymentTriggersHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r containerScaleDeploymentTriggersHTTPJSON) RawJSON() string {
	return r.raw
}

type ContainerScaleDeploymentParam struct {
	// Maximum scale for the container
	Max param.Field[int64] `json:"max,required"`
	// Minimum scale for the container
	Min param.Field[int64] `json:"min,required"`
	// Cooldown period between scaling actions in seconds
	CooldownPeriod param.Field[int64] `json:"cooldown_period"`
	// Triggers for scaling actions
	Triggers param.Field[ContainerScaleDeploymentTriggersParam] `json:"triggers"`
}

func (r ContainerScaleDeploymentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Triggers for scaling actions
type ContainerScaleDeploymentTriggersParam struct {
	CPU            param.Field[ThresholdParam]                            `json:"cpu"`
	GPUMemory      param.Field[ThresholdParam]                            `json:"gpu_memory"`
	GPUUtilization param.Field[ThresholdParam]                            `json:"gpu_utilization"`
	HTTP           param.Field[ContainerScaleDeploymentTriggersHTTPParam] `json:"http"`
	Memory         param.Field[ThresholdParam]                            `json:"memory"`
}

func (r ContainerScaleDeploymentTriggersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContainerScaleDeploymentTriggersHTTPParam struct {
	Rate   param.Field[int64] `json:"rate,required"`
	Window param.Field[int64] `json:"window,required"`
}

func (r ContainerScaleDeploymentTriggersHTTPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceInParam struct {
	// List of containers for the inference instance.
	Containers param.Field[[]InstanceInContainerParam] `json:"containers,required"`
	// Flavor ID for the inference instance.
	FlavorID param.Field[string] `json:"flavor_id,required" format:"uuid"`
	// Docker image for the inference instance. This field should contain the image
	// name and tag in the format 'name:tag', e.g., 'nginx:latest'. It defaults to
	// Docker Hub as the image registry, but any accessible Docker image URL can be
	// specified.
	Image param.Field[string] `json:"image,required"`
	// Listening port for the inference instance.
	ListeningPort param.Field[int64] `json:"listening_port,required"`
	// Inference instance name.
	Name param.Field[string] `json:"name,required"`
	// List of API keys IDs to attach to the inference instance
	APIKeyIDs param.Field[[]string] `json:"api_key_ids" format:"uuid"`
	// Set to true to enable API key authentication for the inference instance. Manage
	// API keys through the '/v1/inference_instances/keys' endpoint.
	AuthEnabled param.Field[bool] `json:"auth_enabled"`
	// Command to be executed when running a container from an image.
	Command param.Field[[]string] `json:"command"`
	// Inference instance description.
	Description param.Field[string] `json:"description"`
	// Environment variables for the inference instance.
	Envs param.Field[map[string]string] `json:"envs"`
	// Image registry ID for authentication in private registries. Leave this parameter
	// empty if no authentication is required for the repository.
	ImageRegistryID param.Field[string] `json:"image_registry_id" format:"uuid"`
	// Probes configured for all containers of the inference instance. If probes are
	// not provided, and the image_name is from a the Model Catalog registry, the
	// default probes will be used.
	Probes param.Field[InstanceProbesDeploymentParam] `json:"probes"`
	// Project ID. If not provided, your default project ID will be used.
	ProjectID param.Field[int64] `json:"project_id"`
	// Specifies the duration in seconds without any requests after which the
	// containers will be downscaled to their minimum scale value as defined by
	// `scale.min`. If set, this helps in optimizing resource usage by reducing the
	// number of container instances during periods of inactivity. The default value
	// when the parameter is not set is 120.
	Timeout param.Field[int64] `json:"timeout"`
}

func (r InstanceInParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceInContainerParam struct {
	// Region ID for the container
	RegionID param.Field[int64] `json:"region_id,required"`
	// Scale for the container
	Scale param.Field[ContainerScaleDeploymentParam] `json:"scale,required"`
}

func (r InstanceInContainerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceOut struct {
	// Inference instance ID.
	ID string `json:"id,required" format:"uuid"`
	// List of containers for the inference instance
	Containers []ContainerOut `json:"containers,required"`
	// Flavor ID for the inference instance
	FlavorID string `json:"flavor_id,required" format:"uuid"`
	// Docker image for the inference instance. This field should contain the image
	// name and tag in the format 'name:tag', e.g., 'nginx:latest'. It defaults to
	// Docker Hub as the image registry, but any accessible Docker image URL can be
	// specified.
	Image string `json:"image,required"`
	// Listening port for the inference instance.
	ListeningPort int64 `json:"listening_port,required"`
	// Inference instance name.
	Name string `json:"name,required"`
	// Inference instance status
	Status InstanceStatus `json:"status,required"`
	// Address of the inference instance
	Address string `json:"address,nullable" format:"uri"`
	// List of API keys IDs attached to the inference instance
	APIKeyIDs []string `json:"api_key_ids" format:"uuid"`
	// Set to true if instance uses API key authentication. Manage API keys through the
	// '/v1/inference_instances/keys' endpoint.
	AuthEnabled bool `json:"auth_enabled"`
	// Command to be executed when running a container from an image.
	Command []string `json:"command,nullable"`
	// Inference instance creation date in ISO 8601 format.
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// Inference instance description.
	Description string `json:"description,nullable"`
	// Environment variables for the inference instance
	Envs map[string]string `json:"envs,nullable"`
	// Image registry ID for authentication in private registries. This parameter is
	// empty if no authentication is required for the repository.
	ImageRegistryID string `json:"image_registry_id,nullable" format:"uuid"`
	// Probes configured for all containers of the inference instance.
	Probes InstanceProbesDeployment `json:"probes,nullable"`
	// Specifies the duration in seconds without any requests after which the
	// containers will be downscaled to their minimum scale value as defined by
	// `scale.min`. If set, this helps in optimizing resource usage by reducing the
	// number of container instances during periods of inactivity.
	Timeout int64           `json:"timeout,nullable"`
	JSON    instanceOutJSON `json:"-"`
}

// instanceOutJSON contains the JSON metadata for the struct [InstanceOut]
type instanceOutJSON struct {
	ID              apijson.Field
	Containers      apijson.Field
	FlavorID        apijson.Field
	Image           apijson.Field
	ListeningPort   apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	Address         apijson.Field
	APIKeyIDs       apijson.Field
	AuthEnabled     apijson.Field
	Command         apijson.Field
	CreatedAt       apijson.Field
	Description     apijson.Field
	Envs            apijson.Field
	ImageRegistryID apijson.Field
	Probes          apijson.Field
	Timeout         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InstanceOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceOutJSON) RawJSON() string {
	return r.raw
}

type InstanceProbesDeployment struct {
	// Liveness probe configuration
	LivenessProbe ContainerProbeConfiguration `json:"liveness_probe,nullable"`
	// Readiness probe configuration
	ReadinessProbe ContainerProbeConfiguration `json:"readiness_probe,nullable"`
	// Startup probe configuration
	StartupProbe ContainerProbeConfiguration  `json:"startup_probe,nullable"`
	JSON         instanceProbesDeploymentJSON `json:"-"`
}

// instanceProbesDeploymentJSON contains the JSON metadata for the struct
// [InstanceProbesDeployment]
type instanceProbesDeploymentJSON struct {
	LivenessProbe  apijson.Field
	ReadinessProbe apijson.Field
	StartupProbe   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceProbesDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceProbesDeploymentJSON) RawJSON() string {
	return r.raw
}

type InstanceProbesDeploymentParam struct {
	// Liveness probe configuration
	LivenessProbe param.Field[ContainerProbeConfigurationParam] `json:"liveness_probe"`
	// Readiness probe configuration
	ReadinessProbe param.Field[ContainerProbeConfigurationParam] `json:"readiness_probe"`
	// Startup probe configuration
	StartupProbe param.Field[ContainerProbeConfigurationParam] `json:"startup_probe"`
}

func (r InstanceProbesDeploymentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceStatus string

const (
	InstanceStatusActive            InstanceStatus = "ACTIVE"
	InstanceStatusDeleted           InstanceStatus = "DELETED"
	InstanceStatusDeleting          InstanceStatus = "DELETING"
	InstanceStatusDeploying         InstanceStatus = "DEPLOYING"
	InstanceStatusDisabled          InstanceStatus = "DISABLED"
	InstanceStatusError             InstanceStatus = "ERROR"
	InstanceStatusFailed            InstanceStatus = "FAILED"
	InstanceStatusNew               InstanceStatus = "NEW"
	InstanceStatusPartiallydeployed InstanceStatus = "PARTIALLYDEPLOYED"
	InstanceStatusPending           InstanceStatus = "PENDING"
)

func (r InstanceStatus) IsKnown() bool {
	switch r {
	case InstanceStatusActive, InstanceStatusDeleted, InstanceStatusDeleting, InstanceStatusDeploying, InstanceStatusDisabled, InstanceStatusError, InstanceStatusFailed, InstanceStatusNew, InstanceStatusPartiallydeployed, InstanceStatusPending:
		return true
	}
	return false
}

type Threshold struct {
	Threshold int64         `json:"threshold,required"`
	JSON      thresholdJSON `json:"-"`
}

// thresholdJSON contains the JSON metadata for the struct [Threshold]
type thresholdJSON struct {
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Threshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r thresholdJSON) RawJSON() string {
	return r.raw
}

type ThresholdParam struct {
	Threshold param.Field[int64] `json:"threshold,required"`
}

func (r ThresholdParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV2InferenceDeploymentNewResponse struct {
	// Inference instance ID.
	ID string `json:"id,required" format:"uuid"`
	// Flavor ID for the inference instance
	FlavorID string `json:"flavor_id,required" format:"uuid"`
	// Docker image for the inference instance. This field should contain the image
	// name and tag in the format 'name:tag', e.g., 'nginx:latest'. It defaults to
	// Docker Hub as the image registry, but any accessible Docker image URL can be
	// specified.
	Image string `json:"image,required"`
	// Listening port for the inference instance.
	ListeningPort int64 `json:"listening_port,required"`
	// Inference instance name.
	Name string `json:"name,required"`
	// Inference instance status
	Status InstanceStatus `json:"status,required"`
	// Address of the inference instance
	Address string `json:"address,nullable" format:"uri"`
	// List of API keys IDs attached to the inference instance
	APIKeyIDs []string `json:"api_key_ids" format:"uuid"`
	// Set to true if instance uses API key authentication. Manage API keys through the
	// '/v1/inference_instances/keys' endpoint.
	AuthEnabled bool `json:"auth_enabled"`
	// Command to be executed when running a container from an image.
	Command []string `json:"command,nullable"`
	// List of containers for the inference instance
	Containers []ContainerOut `json:"containers"`
	// Inference instance creation date in ISO 8601 format.
	CreatedAt time.Time `json:"created_at,nullable" format:"date-time"`
	// Inference instance description.
	Description string `json:"description,nullable"`
	// Environment variables for the inference instance
	Envs map[string]string `json:"envs,nullable"`
	// Image registry ID for authentication in private registries. This parameter is
	// empty if no authentication is required for the repository.
	ImageRegistryID string `json:"image_registry_id,nullable" format:"uuid"`
	// Probes configured for all containers of the inference instance.
	Probes InstanceProbesDeployment `json:"probes,nullable"`
	// Specifies the duration in seconds without any requests after which the
	// containers will be downscaled to their minimum scale value as defined by
	// `scale.min`. If set, this helps in optimizing resource usage by reducing the
	// number of container instances during periods of inactivity.
	Timeout int64                                     `json:"timeout,nullable"`
	JSON    cloudV2InferenceDeploymentNewResponseJSON `json:"-"`
}

// cloudV2InferenceDeploymentNewResponseJSON contains the JSON metadata for the
// struct [CloudV2InferenceDeploymentNewResponse]
type cloudV2InferenceDeploymentNewResponseJSON struct {
	ID              apijson.Field
	FlavorID        apijson.Field
	Image           apijson.Field
	ListeningPort   apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	Address         apijson.Field
	APIKeyIDs       apijson.Field
	AuthEnabled     apijson.Field
	Command         apijson.Field
	Containers      apijson.Field
	CreatedAt       apijson.Field
	Description     apijson.Field
	Envs            apijson.Field
	ImageRegistryID apijson.Field
	Probes          apijson.Field
	Timeout         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CloudV2InferenceDeploymentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV2InferenceDeploymentNewResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV2InferenceDeploymentListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []InstanceOut                              `json:"results,required"`
	JSON    cloudV2InferenceDeploymentListResponseJSON `json:"-"`
}

// cloudV2InferenceDeploymentListResponseJSON contains the JSON metadata for the
// struct [CloudV2InferenceDeploymentListResponse]
type cloudV2InferenceDeploymentListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV2InferenceDeploymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV2InferenceDeploymentListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV2InferenceDeploymentGetLogsResponse struct {
	// Log message.
	Message string `json:"message,required"`
	// Pod name.
	Pod string `json:"pod,required"`
	// Region ID where the container is deployed.
	RegionID int64 `json:"region_id,required"`
	// Log message timestamp in ISO 8601 format.
	Time time.Time                                     `json:"time,required" format:"date-time"`
	JSON cloudV2InferenceDeploymentGetLogsResponseJSON `json:"-"`
}

// cloudV2InferenceDeploymentGetLogsResponseJSON contains the JSON metadata for the
// struct [CloudV2InferenceDeploymentGetLogsResponse]
type cloudV2InferenceDeploymentGetLogsResponseJSON struct {
	Message     apijson.Field
	Pod         apijson.Field
	RegionID    apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV2InferenceDeploymentGetLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV2InferenceDeploymentGetLogsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV2InferenceDeploymentNewParams struct {
	InstanceIn InstanceInParam `json:"instance_in,required"`
}

func (r CloudV2InferenceDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.InstanceIn)
}

type CloudV2InferenceDeploymentUpdateParams struct {
	InstanceIn InstanceInParam `json:"instance_in,required"`
}

func (r CloudV2InferenceDeploymentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.InstanceIn)
}

type CloudV2InferenceDeploymentListParams struct {
	// Limit the number of returned instances. Limited by max limit value of 1000
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
	// Order instances by transmitted fields and directions (name.asc)
	OrderBy param.Field[string] `query:"order_by"`
	// Project ID
	ProjectID param.Field[int64] `query:"project_id"`
}

// URLQuery serializes [CloudV2InferenceDeploymentListParams]'s query parameters as
// `url.Values`.
func (r CloudV2InferenceDeploymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
