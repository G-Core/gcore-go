// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// InferenceService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceService] method instead.
type InferenceService struct {
	Options             []option.RequestOption
	Flavors             InferenceFlavorService
	Models              InferenceModelService
	Deployments         InferenceDeploymentService
	RegistryCredentials InferenceRegistryCredentialService
	Secrets             InferenceSecretService
}

// NewInferenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInferenceService(opts ...option.RequestOption) (r InferenceService) {
	r = InferenceService{}
	r.Options = opts
	r.Flavors = NewInferenceFlavorService(opts...)
	r.Models = NewInferenceModelService(opts...)
	r.Deployments = NewInferenceDeploymentService(opts...)
	r.RegistryCredentials = NewInferenceRegistryCredentialService(opts...)
	r.Secrets = NewInferenceSecretService(opts...)
	return
}

// Get inference capacity by region
func (r *InferenceService) GetCapacityByRegion(ctx context.Context, opts ...option.RequestOption) (res *RegionCapacityList, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v3/inference/capacity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AwsIamData struct {
	// AWS IAM key ID.
	AwsAccessKeyID string `json:"aws_access_key_id,required"`
	// AWS IAM secret key.
	AwsSecretAccessKey string `json:"aws_secret_access_key,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwsAccessKeyID     respjson.Field
		AwsSecretAccessKey respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AwsIamData) RawJSON() string { return r.JSON.raw }
func (r *AwsIamData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AwsIamData to a AwsIamDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AwsIamDataParam.Overrides()
func (r AwsIamData) ToParam() AwsIamDataParam {
	return param.Override[AwsIamDataParam](json.RawMessage(r.RawJSON()))
}

// The properties AwsAccessKeyID, AwsSecretAccessKey are required.
type AwsIamDataParam struct {
	// AWS IAM key ID.
	AwsAccessKeyID string `json:"aws_access_key_id,required"`
	// AWS IAM secret key.
	AwsSecretAccessKey string `json:"aws_secret_access_key,required"`
	paramObj
}

func (r AwsIamDataParam) MarshalJSON() (data []byte, err error) {
	type shadow AwsIamDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AwsIamDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Capacity struct {
	// Available capacity.
	Capacity int64 `json:"capacity,required"`
	// Flavor name.
	FlavorName string `json:"flavor_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capacity    respjson.Field
		FlavorName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Capacity) RawJSON() string { return r.JSON.raw }
func (r *Capacity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContainerProbe struct {
	// Exec probe configuration
	Exec ContainerProbeExec `json:"exec,required"`
	// The number of consecutive probe failures that mark the container as unhealthy.
	FailureThreshold int64 `json:"failure_threshold,required"`
	// HTTP GET probe configuration
	HTTPGet ContainerProbeHTTPGet `json:"http_get,required"`
	// The initial delay before starting the first probe.
	InitialDelaySeconds int64 `json:"initial_delay_seconds,required"`
	// How often (in seconds) to perform the probe.
	PeriodSeconds int64 `json:"period_seconds,required"`
	// The number of consecutive successful probes that mark the container as healthy.
	SuccessThreshold int64 `json:"success_threshold,required"`
	// TCP socket probe configuration
	TcpSocket ContainerProbeTcpSocket `json:"tcp_socket,required"`
	// The timeout for each probe.
	TimeoutSeconds int64 `json:"timeout_seconds,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exec                respjson.Field
		FailureThreshold    respjson.Field
		HTTPGet             respjson.Field
		InitialDelaySeconds respjson.Field
		PeriodSeconds       respjson.Field
		SuccessThreshold    respjson.Field
		TcpSocket           respjson.Field
		TimeoutSeconds      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerProbe) RawJSON() string { return r.JSON.raw }
func (r *ContainerProbe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContainerProbeConfig struct {
	// Whether the probe is enabled or not.
	Enabled bool `json:"enabled,required"`
	// Probe configuration (exec, http_get or tcp_socket)
	Probe ContainerProbe `json:"probe,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Probe       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerProbeConfig) RawJSON() string { return r.JSON.raw }
func (r *ContainerProbeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type ContainerProbeConfigCreateParam struct {
	// Whether the probe is enabled or not.
	Enabled bool `json:"enabled,required"`
	// Probe configuration (exec, http_get or tcp_socket)
	Probe ContainerProbeCreateParam `json:"probe,omitzero"`
	paramObj
}

func (r ContainerProbeConfigCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ContainerProbeConfigCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContainerProbeConfigCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContainerProbeCreateParam struct {
	// The number of consecutive probe failures that mark the container as unhealthy.
	FailureThreshold param.Opt[int64] `json:"failure_threshold,omitzero"`
	// The initial delay before starting the first probe.
	InitialDelaySeconds param.Opt[int64] `json:"initial_delay_seconds,omitzero"`
	// How often (in seconds) to perform the probe.
	PeriodSeconds param.Opt[int64] `json:"period_seconds,omitzero"`
	// The number of consecutive successful probes that mark the container as healthy.
	SuccessThreshold param.Opt[int64] `json:"success_threshold,omitzero"`
	// The timeout for each probe.
	TimeoutSeconds param.Opt[int64] `json:"timeout_seconds,omitzero"`
	// Exec probe configuration
	Exec ContainerProbeExecCreateParam `json:"exec,omitzero"`
	// HTTP GET probe configuration
	HTTPGet ContainerProbeHTTPGetCreateParam `json:"http_get,omitzero"`
	// TCP socket probe configuration
	TcpSocket ContainerProbeTcpSocketCreateParam `json:"tcp_socket,omitzero"`
	paramObj
}

func (r ContainerProbeCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ContainerProbeCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContainerProbeCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContainerProbeExec struct {
	// Command to be executed inside the running container.
	Command []string `json:"command,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Command     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerProbeExec) RawJSON() string { return r.JSON.raw }
func (r *ContainerProbeExec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Command is required.
type ContainerProbeExecCreateParam struct {
	// Command to be executed inside the running container.
	Command []string `json:"command,omitzero,required"`
	paramObj
}

func (r ContainerProbeExecCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ContainerProbeExecCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContainerProbeExecCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContainerProbeHTTPGet struct {
	// HTTP headers to be sent with the request.
	Headers map[string]string `json:"headers,required"`
	// Host name to send HTTP request to.
	Host string `json:"host,required"`
	// The endpoint to send the HTTP request to.
	Path string `json:"path,required"`
	// Port number the probe should connect to.
	Port int64 `json:"port,required"`
	// Schema to use for the HTTP request.
	Schema string `json:"schema,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers     respjson.Field
		Host        respjson.Field
		Path        respjson.Field
		Port        respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerProbeHTTPGet) RawJSON() string { return r.JSON.raw }
func (r *ContainerProbeHTTPGet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Port is required.
type ContainerProbeHTTPGetCreateParam struct {
	// Port number the probe should connect to.
	Port int64 `json:"port,required"`
	// Host name to send HTTP request to.
	Host param.Opt[string] `json:"host,omitzero"`
	// The endpoint to send the HTTP request to.
	Path param.Opt[string] `json:"path,omitzero"`
	// Schema to use for the HTTP request.
	Schema param.Opt[string] `json:"schema,omitzero"`
	// HTTP headers to be sent with the request.
	Headers map[string]string `json:"headers,omitzero"`
	paramObj
}

func (r ContainerProbeHTTPGetCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ContainerProbeHTTPGetCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContainerProbeHTTPGetCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContainerProbeTcpSocket struct {
	// Port number to check if it's open.
	Port int64 `json:"port,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Port        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerProbeTcpSocket) RawJSON() string { return r.JSON.raw }
func (r *ContainerProbeTcpSocket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Port is required.
type ContainerProbeTcpSocketCreateParam struct {
	// Port number to check if it's open.
	Port int64 `json:"port,required"`
	paramObj
}

func (r ContainerProbeTcpSocketCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ContainerProbeTcpSocketCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContainerProbeTcpSocketCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContainerScale struct {
	// Cooldown period between scaling actions in seconds
	CooldownPeriod int64 `json:"cooldown_period,required"`
	// Maximum scale for the container
	Max int64 `json:"max,required"`
	// Minimum scale for the container
	Min int64 `json:"min,required"`
	// Polling interval for scaling triggers in seconds
	PollingInterval int64 `json:"polling_interval,required"`
	// Triggers for scaling actions
	Triggers ContainerScaleTriggers `json:"triggers,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CooldownPeriod  respjson.Field
		Max             respjson.Field
		Min             respjson.Field
		PollingInterval respjson.Field
		Triggers        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerScale) RawJSON() string { return r.JSON.raw }
func (r *ContainerScale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContainerScaleTriggerRate struct {
	// Request count per 'window' seconds for the http trigger
	Rate int64 `json:"rate,required"`
	// Time window for rate calculation in seconds
	Window int64 `json:"window,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rate        respjson.Field
		Window      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerScaleTriggerRate) RawJSON() string { return r.JSON.raw }
func (r *ContainerScaleTriggerRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContainerScaleTriggerSqs struct {
	// Number of messages for activation
	ActivationQueueLength int64 `json:"activation_queue_length,required"`
	// Custom AWS endpoint
	AwsEndpoint string `json:"aws_endpoint,required"`
	// AWS region
	AwsRegion string `json:"aws_region,required"`
	// Number of messages for one replica
	QueueLength int64 `json:"queue_length,required"`
	// SQS queue URL
	QueueURL string `json:"queue_url,required"`
	// Scale on delayed messages
	ScaleOnDelayed bool `json:"scale_on_delayed,required"`
	// Scale on in-flight messages
	ScaleOnFlight bool `json:"scale_on_flight,required"`
	// Auth secret name
	SecretName string `json:"secret_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActivationQueueLength respjson.Field
		AwsEndpoint           respjson.Field
		AwsRegion             respjson.Field
		QueueLength           respjson.Field
		QueueURL              respjson.Field
		ScaleOnDelayed        respjson.Field
		ScaleOnFlight         respjson.Field
		SecretName            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerScaleTriggerSqs) RawJSON() string { return r.JSON.raw }
func (r *ContainerScaleTriggerSqs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContainerScaleTriggerThreshold struct {
	// Threshold value for the trigger in percentage
	Threshold int64 `json:"threshold,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Threshold   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerScaleTriggerThreshold) RawJSON() string { return r.JSON.raw }
func (r *ContainerScaleTriggerThreshold) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContainerScaleTriggers struct {
	// CPU trigger configuration
	CPU ContainerScaleTriggerThreshold `json:"cpu,required"`
	// GPU memory trigger configuration. Calculated by DCGM_FI_DEV_MEM_COPY_UTIL metric
	GPUMemory ContainerScaleTriggerThreshold `json:"gpu_memory,required"`
	// GPU utilization trigger configuration. Calculated by DCGM_FI_DEV_GPU_UTIL metric
	GPUUtilization ContainerScaleTriggerThreshold `json:"gpu_utilization,required"`
	// HTTP trigger configuration
	HTTP ContainerScaleTriggerRate `json:"http,required"`
	// Memory trigger configuration
	Memory ContainerScaleTriggerThreshold `json:"memory,required"`
	// SQS trigger configuration
	Sqs ContainerScaleTriggerSqs `json:"sqs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU            respjson.Field
		GPUMemory      respjson.Field
		GPUUtilization respjson.Field
		HTTP           respjson.Field
		Memory         respjson.Field
		Sqs            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerScaleTriggers) RawJSON() string { return r.JSON.raw }
func (r *ContainerScaleTriggers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeployStatus struct {
	// Number of ready instances
	Ready int64 `json:"ready,required"`
	// Total number of instances
	Total int64 `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ready       respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeployStatus) RawJSON() string { return r.JSON.raw }
func (r *DeployStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceProbes struct {
	// Liveness probe configuration
	LivenessProbe ContainerProbeConfig `json:"liveness_probe,required"`
	// Readiness probe configuration
	ReadinessProbe ContainerProbeConfig `json:"readiness_probe,required"`
	// Startup probe configuration
	StartupProbe ContainerProbeConfig `json:"startup_probe,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LivenessProbe  respjson.Field
		ReadinessProbe respjson.Field
		StartupProbe   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceProbes) RawJSON() string { return r.JSON.raw }
func (r *InferenceProbes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IngressOptsParam struct {
	// Disable response buffering if true. A client usually has a much slower
	// connection and can not consume the response data as fast as it is produced by an
	// upstream application. Ingress tries to buffer the whole response in order to
	// release the upstream application as soon as possible.By default, the response
	// buffering is enabled.
	DisableResponseBuffering param.Opt[bool] `json:"disable_response_buffering,omitzero"`
	paramObj
}

func (r IngressOptsParam) MarshalJSON() (data []byte, err error) {
	type shadow IngressOptsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IngressOptsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IngressOptsOut struct {
	// Disable response buffering if true. A client usually has a much slower
	// connection and can not consume the response data as fast as it is produced by an
	// upstream application. Ingress tries to buffer the whole response in order to
	// release the upstream application as soon as possible.By default, the response
	// buffering is enabled.
	DisableResponseBuffering bool `json:"disable_response_buffering,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisableResponseBuffering respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IngressOptsOut) RawJSON() string { return r.JSON.raw }
func (r *IngressOptsOut) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegionCapacity struct {
	// List of capacities by flavor.
	Capacity []Capacity `json:"capacity,required"`
	// Region ID.
	RegionID int64 `json:"region_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capacity    respjson.Field
		RegionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegionCapacity) RawJSON() string { return r.JSON.raw }
func (r *RegionCapacity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegionCapacityList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []RegionCapacity `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegionCapacityList) RawJSON() string { return r.JSON.raw }
func (r *RegionCapacityList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
