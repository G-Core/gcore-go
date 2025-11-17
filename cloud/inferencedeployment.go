// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

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
}

// NewInferenceDeploymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInferenceDeploymentService(opts ...option.RequestOption) (r InferenceDeploymentService) {
	r = InferenceDeploymentService{}
	r.Options = opts
	r.Logs = NewInferenceDeploymentLogService(opts...)
	return
}

// Create inference deployment
func (r *InferenceDeploymentService) New(ctx context.Context, params InferenceDeploymentNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
func (r *InferenceDeploymentService) List(ctx context.Context, params InferenceDeploymentListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[InferenceDeployment], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
func (r *InferenceDeploymentService) ListAutoPaging(ctx context.Context, params InferenceDeploymentListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[InferenceDeployment] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete inference deployment
func (r *InferenceDeploymentService) Delete(ctx context.Context, deploymentName string, body InferenceDeploymentDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
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
func (r *InferenceDeploymentService) Get(ctx context.Context, deploymentName string, query InferenceDeploymentGetParams, opts ...option.RequestOption) (res *InferenceDeployment, err error) {
	opts = slices.Concat(r.Options, opts)
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
//
// Deprecated: deprecated
func (r *InferenceDeploymentService) GetAPIKey(ctx context.Context, deploymentName string, query InferenceDeploymentGetAPIKeyParams, opts ...option.RequestOption) (res *InferenceDeploymentAPIKey, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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

type InferenceDeployment struct {
	// Address of the inference instance
	Address string `json:"address,required" format:"uri"`
	// `true` if instance uses API key authentication.
	// `"Authorization": "Bearer *****"` or `"X-Api-Key": "*****"` header is required
	// for the requests to the instance if enabled.
	//
	// Deprecated: deprecated
	AuthEnabled bool `json:"auth_enabled,required"`
	// Command to be executed when running a container from an image.
	Command string `json:"command,required"`
	// List of containers for the inference instance
	Containers []InferenceDeploymentContainer `json:"containers,required"`
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
	IngressOpts InferenceDeploymentIngressOpts `json:"ingress_opts,required"`
	// Listening port for the inference instance.
	ListeningPort int64 `json:"listening_port,required"`
	// Logging configuration for the inference instance
	Logging Logging `json:"logging,required"`
	// Inference instance name.
	Name string `json:"name,required"`
	// Indicates to which parent object this inference belongs to.
	ObjectReferences []InferenceDeploymentObjectReference `json:"object_references,required"`
	// Probes configured for all containers of the inference instance.
	Probes InferenceDeploymentProbes `json:"probes,required"`
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
	//   - `PENDING` - The instance is running but scaled to zero. It will be
	//     automatically scaled up when a request is made.
	//   - `DELETING` - The instance is being deleted.
	//
	// Any of "ACTIVE", "DELETING", "DEPLOYING", "DISABLED", "PARTIALLYDEPLOYED",
	// "PENDING".
	Status InferenceDeploymentStatus `json:"status,required"`
	// Specifies the duration in seconds without any requests after which the
	// containers will be downscaled to their minimum scale value as defined by
	// `scale.min`. If set, this helps in optimizing resource usage by reducing the
	// number of container instances during periods of inactivity.
	Timeout int64 `json:"timeout,required"`
	// List of API keys for the inference instance
	APIKeys []string `json:"api_keys,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address          respjson.Field
		AuthEnabled      respjson.Field
		Command          respjson.Field
		Containers       respjson.Field
		CreatedAt        respjson.Field
		CredentialsName  respjson.Field
		Description      respjson.Field
		Envs             respjson.Field
		FlavorName       respjson.Field
		Image            respjson.Field
		IngressOpts      respjson.Field
		ListeningPort    respjson.Field
		Logging          respjson.Field
		Name             respjson.Field
		ObjectReferences respjson.Field
		Probes           respjson.Field
		ProjectID        respjson.Field
		Status           respjson.Field
		Timeout          respjson.Field
		APIKeys          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceDeployment) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeployment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceDeploymentContainer struct {
	// Address of the inference instance
	Address string `json:"address,required" format:"uri"`
	// Status of the containers deployment
	DeployStatus InferenceDeploymentContainerDeployStatus `json:"deploy_status,required"`
	// Error message if the container deployment failed
	ErrorMessage string `json:"error_message,required"`
	// Region name for the container
	RegionID int64 `json:"region_id,required"`
	// Scale for the container
	Scale InferenceDeploymentContainerScale `json:"scale,required"`
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
func (r InferenceDeploymentContainer) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentContainer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the containers deployment
type InferenceDeploymentContainerDeployStatus struct {
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
func (r InferenceDeploymentContainerDeployStatus) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentContainerDeployStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scale for the container
type InferenceDeploymentContainerScale struct {
	// Cooldown period between scaling actions in seconds
	CooldownPeriod int64 `json:"cooldown_period,required"`
	// Maximum scale for the container
	Max int64 `json:"max,required"`
	// Minimum scale for the container
	Min int64 `json:"min,required"`
	// Polling interval for scaling triggers in seconds
	PollingInterval int64 `json:"polling_interval,required"`
	// Triggers for scaling actions
	Triggers InferenceDeploymentContainerScaleTriggers `json:"triggers,required"`
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
func (r InferenceDeploymentContainerScale) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentContainerScale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Triggers for scaling actions
type InferenceDeploymentContainerScaleTriggers struct {
	// CPU trigger configuration
	CPU InferenceDeploymentContainerScaleTriggersCPU `json:"cpu,required"`
	// GPU memory trigger configuration. Calculated by `DCGM_FI_DEV_MEM_COPY_UTIL`
	// metric
	GPUMemory InferenceDeploymentContainerScaleTriggersGPUMemory `json:"gpu_memory,required"`
	// GPU utilization trigger configuration. Calculated by `DCGM_FI_DEV_GPU_UTIL`
	// metric
	GPUUtilization InferenceDeploymentContainerScaleTriggersGPUUtilization `json:"gpu_utilization,required"`
	// HTTP trigger configuration
	HTTP InferenceDeploymentContainerScaleTriggersHTTP `json:"http,required"`
	// Memory trigger configuration
	Memory InferenceDeploymentContainerScaleTriggersMemory `json:"memory,required"`
	// SQS trigger configuration
	Sqs InferenceDeploymentContainerScaleTriggersSqs `json:"sqs,required"`
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
func (r InferenceDeploymentContainerScaleTriggers) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentContainerScaleTriggers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CPU trigger configuration
type InferenceDeploymentContainerScaleTriggersCPU struct {
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
func (r InferenceDeploymentContainerScaleTriggersCPU) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentContainerScaleTriggersCPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPU memory trigger configuration. Calculated by `DCGM_FI_DEV_MEM_COPY_UTIL`
// metric
type InferenceDeploymentContainerScaleTriggersGPUMemory struct {
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
func (r InferenceDeploymentContainerScaleTriggersGPUMemory) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentContainerScaleTriggersGPUMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GPU utilization trigger configuration. Calculated by `DCGM_FI_DEV_GPU_UTIL`
// metric
type InferenceDeploymentContainerScaleTriggersGPUUtilization struct {
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
func (r InferenceDeploymentContainerScaleTriggersGPUUtilization) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentContainerScaleTriggersGPUUtilization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP trigger configuration
type InferenceDeploymentContainerScaleTriggersHTTP struct {
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
func (r InferenceDeploymentContainerScaleTriggersHTTP) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentContainerScaleTriggersHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Memory trigger configuration
type InferenceDeploymentContainerScaleTriggersMemory struct {
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
func (r InferenceDeploymentContainerScaleTriggersMemory) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentContainerScaleTriggersMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SQS trigger configuration
type InferenceDeploymentContainerScaleTriggersSqs struct {
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
func (r InferenceDeploymentContainerScaleTriggersSqs) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentContainerScaleTriggersSqs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ingress options for the inference instance
type InferenceDeploymentIngressOpts struct {
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
func (r InferenceDeploymentIngressOpts) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentIngressOpts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceDeploymentObjectReference struct {
	// Kind of the inference object to be referenced
	//
	// Any of "AppDeployment".
	Kind string `json:"kind,required"`
	// Name of the inference object to be referenced
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceDeploymentObjectReference) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentObjectReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Probes configured for all containers of the inference instance.
type InferenceDeploymentProbes struct {
	// Liveness probe configuration
	LivenessProbe ProbeConfig `json:"liveness_probe,required"`
	// Readiness probe configuration
	ReadinessProbe ProbeConfig `json:"readiness_probe,required"`
	// Startup probe configuration
	StartupProbe ProbeConfig `json:"startup_probe,required"`
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
func (r InferenceDeploymentProbes) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentProbes) UnmarshalJSON(data []byte) error {
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
//   - `PENDING` - The instance is running but scaled to zero. It will be
//     automatically scaled up when a request is made.
//   - `DELETING` - The instance is being deleted.
type InferenceDeploymentStatus string

const (
	InferenceDeploymentStatusActive            InferenceDeploymentStatus = "ACTIVE"
	InferenceDeploymentStatusDeleting          InferenceDeploymentStatus = "DELETING"
	InferenceDeploymentStatusDeploying         InferenceDeploymentStatus = "DEPLOYING"
	InferenceDeploymentStatusDisabled          InferenceDeploymentStatus = "DISABLED"
	InferenceDeploymentStatusPartiallydeployed InferenceDeploymentStatus = "PARTIALLYDEPLOYED"
	InferenceDeploymentStatusPending           InferenceDeploymentStatus = "PENDING"
)

type InferenceDeploymentAPIKey struct {
	// API key secret
	Secret string `json:"secret,required"`
	// API key status
	//
	// Any of "PENDING", "READY".
	Status InferenceDeploymentAPIKeyStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Secret      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceDeploymentAPIKey) RawJSON() string { return r.JSON.raw }
func (r *InferenceDeploymentAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API key status
type InferenceDeploymentAPIKeyStatus string

const (
	InferenceDeploymentAPIKeyStatusPending InferenceDeploymentAPIKeyStatus = "PENDING"
	InferenceDeploymentAPIKeyStatusReady   InferenceDeploymentAPIKeyStatus = "READY"
)

type Probe struct {
	// Exec probe configuration
	Exec ProbeExec `json:"exec,required"`
	// The number of consecutive probe failures that mark the container as unhealthy.
	FailureThreshold int64 `json:"failure_threshold,required"`
	// HTTP GET probe configuration
	HTTPGet ProbeHTTPGet `json:"http_get,required"`
	// The initial delay before starting the first probe.
	InitialDelaySeconds int64 `json:"initial_delay_seconds,required"`
	// How often (in seconds) to perform the probe.
	PeriodSeconds int64 `json:"period_seconds,required"`
	// The number of consecutive successful probes that mark the container as healthy.
	SuccessThreshold int64 `json:"success_threshold,required"`
	// TCP socket probe configuration
	TcpSocket ProbeTcpSocket `json:"tcp_socket,required"`
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
func (r Probe) RawJSON() string { return r.JSON.raw }
func (r *Probe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProbeConfig struct {
	// Whether the probe is enabled or not.
	Enabled bool `json:"enabled,required"`
	// Probe configuration (exec, `http_get` or `tcp_socket`)
	Probe Probe `json:"probe,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Probe       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProbeConfig) RawJSON() string { return r.JSON.raw }
func (r *ProbeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProbeExec struct {
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
func (r ProbeExec) RawJSON() string { return r.JSON.raw }
func (r *ProbeExec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProbeHTTPGet struct {
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
func (r ProbeHTTPGet) RawJSON() string { return r.JSON.raw }
func (r *ProbeHTTPGet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProbeTcpSocket struct {
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
func (r ProbeTcpSocket) RawJSON() string { return r.JSON.raw }
func (r *ProbeTcpSocket) UnmarshalJSON(data []byte) error {
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
	// for the requests to the instance if enabled. This field is deprecated and will
	// be removed in the future. Use `api_keys` field instead.If `auth_enabled` and
	// `api_keys` are both specified, a ValidationError will be raised.
	AuthEnabled param.Opt[bool] `json:"auth_enabled,omitzero"`
	// Command to be executed when running a container from an image.
	Command []string `json:"command,omitzero"`
	// Ingress options for the inference instance
	IngressOpts InferenceDeploymentNewParamsIngressOpts `json:"ingress_opts,omitzero"`
	// Logging configuration for the inference instance
	Logging InferenceDeploymentNewParamsLogging `json:"logging,omitzero"`
	// Probes configured for all containers of the inference instance. If probes are
	// not provided, and the `image_name` is from a the Model Catalog registry, the
	// default probes will be used.
	Probes InferenceDeploymentNewParamsProbes `json:"probes,omitzero"`
	// List of API keys for the inference instance. Multiple keys can be attached to
	// one deployment.If `auth_enabled` and `api_keys` are both specified, a
	// ValidationError will be raised.
	APIKeys []string `json:"api_keys,omitzero"`
	// Environment variables for the inference instance.
	Envs map[string]string `json:"envs,omitzero"`
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
	// GPU memory trigger configuration. Calculated by `DCGM_FI_DEV_MEM_COPY_UTIL`
	// metric
	GPUMemory InferenceDeploymentNewParamsContainerScaleTriggersGPUMemory `json:"gpu_memory,omitzero"`
	// GPU utilization trigger configuration. Calculated by `DCGM_FI_DEV_GPU_UTIL`
	// metric
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

// GPU memory trigger configuration. Calculated by `DCGM_FI_DEV_MEM_COPY_UTIL`
// metric
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

// GPU utilization trigger configuration. Calculated by `DCGM_FI_DEV_GPU_UTIL`
// metric
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

// Ingress options for the inference instance
type InferenceDeploymentNewParamsIngressOpts struct {
	// Disable response buffering if true. A client usually has a much slower
	// connection and can not consume the response data as fast as it is produced by an
	// upstream application. Ingress tries to buffer the whole response in order to
	// release the upstream application as soon as possible.By default, the response
	// buffering is enabled.
	DisableResponseBuffering param.Opt[bool] `json:"disable_response_buffering,omitzero"`
	paramObj
}

func (r InferenceDeploymentNewParamsIngressOpts) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsIngressOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsIngressOpts) UnmarshalJSON(data []byte) error {
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
// not provided, and the `image_name` is from a the Model Catalog registry, the
// default probes will be used.
type InferenceDeploymentNewParamsProbes struct {
	// Liveness probe configuration
	LivenessProbe InferenceDeploymentNewParamsProbesLivenessProbe `json:"liveness_probe,omitzero"`
	// Readiness probe configuration
	ReadinessProbe InferenceDeploymentNewParamsProbesReadinessProbe `json:"readiness_probe,omitzero"`
	// Startup probe configuration
	StartupProbe InferenceDeploymentNewParamsProbesStartupProbe `json:"startup_probe,omitzero"`
	paramObj
}

func (r InferenceDeploymentNewParamsProbes) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbes
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Liveness probe configuration
//
// The property Enabled is required.
type InferenceDeploymentNewParamsProbesLivenessProbe struct {
	// Whether the probe is enabled or not.
	Enabled bool `json:"enabled,required"`
	// Probe configuration (exec, `http_get` or `tcp_socket`)
	Probe InferenceDeploymentNewParamsProbesLivenessProbeProbe `json:"probe,omitzero"`
	paramObj
}

func (r InferenceDeploymentNewParamsProbesLivenessProbe) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesLivenessProbe
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesLivenessProbe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Probe configuration (exec, `http_get` or `tcp_socket`)
type InferenceDeploymentNewParamsProbesLivenessProbeProbe struct {
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
	Exec InferenceDeploymentNewParamsProbesLivenessProbeProbeExec `json:"exec,omitzero"`
	// HTTP GET probe configuration
	HTTPGet InferenceDeploymentNewParamsProbesLivenessProbeProbeHTTPGet `json:"http_get,omitzero"`
	// TCP socket probe configuration
	TcpSocket InferenceDeploymentNewParamsProbesLivenessProbeProbeTcpSocket `json:"tcp_socket,omitzero"`
	paramObj
}

func (r InferenceDeploymentNewParamsProbesLivenessProbeProbe) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesLivenessProbeProbe
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesLivenessProbeProbe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exec probe configuration
//
// The property Command is required.
type InferenceDeploymentNewParamsProbesLivenessProbeProbeExec struct {
	// Command to be executed inside the running container.
	Command []string `json:"command,omitzero,required"`
	paramObj
}

func (r InferenceDeploymentNewParamsProbesLivenessProbeProbeExec) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesLivenessProbeProbeExec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesLivenessProbeProbeExec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP GET probe configuration
//
// The property Port is required.
type InferenceDeploymentNewParamsProbesLivenessProbeProbeHTTPGet struct {
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

func (r InferenceDeploymentNewParamsProbesLivenessProbeProbeHTTPGet) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesLivenessProbeProbeHTTPGet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesLivenessProbeProbeHTTPGet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TCP socket probe configuration
//
// The property Port is required.
type InferenceDeploymentNewParamsProbesLivenessProbeProbeTcpSocket struct {
	// Port number to check if it's open.
	Port int64 `json:"port,required"`
	paramObj
}

func (r InferenceDeploymentNewParamsProbesLivenessProbeProbeTcpSocket) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesLivenessProbeProbeTcpSocket
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesLivenessProbeProbeTcpSocket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Readiness probe configuration
//
// The property Enabled is required.
type InferenceDeploymentNewParamsProbesReadinessProbe struct {
	// Whether the probe is enabled or not.
	Enabled bool `json:"enabled,required"`
	// Probe configuration (exec, `http_get` or `tcp_socket`)
	Probe InferenceDeploymentNewParamsProbesReadinessProbeProbe `json:"probe,omitzero"`
	paramObj
}

func (r InferenceDeploymentNewParamsProbesReadinessProbe) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesReadinessProbe
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesReadinessProbe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Probe configuration (exec, `http_get` or `tcp_socket`)
type InferenceDeploymentNewParamsProbesReadinessProbeProbe struct {
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
	Exec InferenceDeploymentNewParamsProbesReadinessProbeProbeExec `json:"exec,omitzero"`
	// HTTP GET probe configuration
	HTTPGet InferenceDeploymentNewParamsProbesReadinessProbeProbeHTTPGet `json:"http_get,omitzero"`
	// TCP socket probe configuration
	TcpSocket InferenceDeploymentNewParamsProbesReadinessProbeProbeTcpSocket `json:"tcp_socket,omitzero"`
	paramObj
}

func (r InferenceDeploymentNewParamsProbesReadinessProbeProbe) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesReadinessProbeProbe
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesReadinessProbeProbe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exec probe configuration
//
// The property Command is required.
type InferenceDeploymentNewParamsProbesReadinessProbeProbeExec struct {
	// Command to be executed inside the running container.
	Command []string `json:"command,omitzero,required"`
	paramObj
}

func (r InferenceDeploymentNewParamsProbesReadinessProbeProbeExec) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesReadinessProbeProbeExec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesReadinessProbeProbeExec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP GET probe configuration
//
// The property Port is required.
type InferenceDeploymentNewParamsProbesReadinessProbeProbeHTTPGet struct {
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

func (r InferenceDeploymentNewParamsProbesReadinessProbeProbeHTTPGet) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesReadinessProbeProbeHTTPGet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesReadinessProbeProbeHTTPGet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TCP socket probe configuration
//
// The property Port is required.
type InferenceDeploymentNewParamsProbesReadinessProbeProbeTcpSocket struct {
	// Port number to check if it's open.
	Port int64 `json:"port,required"`
	paramObj
}

func (r InferenceDeploymentNewParamsProbesReadinessProbeProbeTcpSocket) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesReadinessProbeProbeTcpSocket
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesReadinessProbeProbeTcpSocket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Startup probe configuration
//
// The property Enabled is required.
type InferenceDeploymentNewParamsProbesStartupProbe struct {
	// Whether the probe is enabled or not.
	Enabled bool `json:"enabled,required"`
	// Probe configuration (exec, `http_get` or `tcp_socket`)
	Probe InferenceDeploymentNewParamsProbesStartupProbeProbe `json:"probe,omitzero"`
	paramObj
}

func (r InferenceDeploymentNewParamsProbesStartupProbe) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesStartupProbe
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesStartupProbe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Probe configuration (exec, `http_get` or `tcp_socket`)
type InferenceDeploymentNewParamsProbesStartupProbeProbe struct {
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
	Exec InferenceDeploymentNewParamsProbesStartupProbeProbeExec `json:"exec,omitzero"`
	// HTTP GET probe configuration
	HTTPGet InferenceDeploymentNewParamsProbesStartupProbeProbeHTTPGet `json:"http_get,omitzero"`
	// TCP socket probe configuration
	TcpSocket InferenceDeploymentNewParamsProbesStartupProbeProbeTcpSocket `json:"tcp_socket,omitzero"`
	paramObj
}

func (r InferenceDeploymentNewParamsProbesStartupProbeProbe) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesStartupProbeProbe
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesStartupProbeProbe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exec probe configuration
//
// The property Command is required.
type InferenceDeploymentNewParamsProbesStartupProbeProbeExec struct {
	// Command to be executed inside the running container.
	Command []string `json:"command,omitzero,required"`
	paramObj
}

func (r InferenceDeploymentNewParamsProbesStartupProbeProbeExec) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesStartupProbeProbeExec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesStartupProbeProbeExec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP GET probe configuration
//
// The property Port is required.
type InferenceDeploymentNewParamsProbesStartupProbeProbeHTTPGet struct {
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

func (r InferenceDeploymentNewParamsProbesStartupProbeProbeHTTPGet) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesStartupProbeProbeHTTPGet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesStartupProbeProbeHTTPGet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TCP socket probe configuration
//
// The property Port is required.
type InferenceDeploymentNewParamsProbesStartupProbeProbeTcpSocket struct {
	// Port number to check if it's open.
	Port int64 `json:"port,required"`
	paramObj
}

func (r InferenceDeploymentNewParamsProbesStartupProbeProbeTcpSocket) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbesStartupProbeProbeTcpSocket
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentNewParamsProbesStartupProbeProbeTcpSocket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceDeploymentUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Registry credentials name
	CredentialsName param.Opt[string] `json:"credentials_name,omitzero"`
	// Inference instance description.
	Description param.Opt[string] `json:"description,omitzero"`
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
	// Set to `true` to enable API key authentication for the inference instance.
	// `"Authorization": "Bearer *****"` or `"X-Api-Key": "*****"` header is required
	// for the requests to the instance if enabled. This field is deprecated and will
	// be removed in the future. Use `api_keys` field instead.If `auth_enabled` and
	// `api_keys` are both specified, a ValidationError will be raised.
	AuthEnabled param.Opt[bool] `json:"auth_enabled,omitzero"`
	// Flavor name for the inference instance.
	FlavorName param.Opt[string] `json:"flavor_name,omitzero"`
	// List of API keys for the inference instance. Multiple keys can be attached to
	// one deployment.If `auth_enabled` and `api_keys` are both specified, a
	// ValidationError will be raised.If `[]` is provided, the API keys will be removed
	// and auth will be disabled on the deployment.
	APIKeys []string `json:"api_keys,omitzero"`
	// Command to be executed when running a container from an image.
	Command []string `json:"command,omitzero"`
	// List of containers for the inference instance.
	Containers []InferenceDeploymentUpdateParamsContainer `json:"containers,omitzero"`
	// Environment variables for the inference instance.
	Envs map[string]string `json:"envs,omitzero"`
	// Ingress options for the inference instance
	IngressOpts InferenceDeploymentUpdateParamsIngressOpts `json:"ingress_opts,omitzero"`
	// Logging configuration for the inference instance
	Logging InferenceDeploymentUpdateParamsLogging `json:"logging,omitzero"`
	// Probes configured for all containers of the inference instance.
	Probes InferenceDeploymentUpdateParamsProbes `json:"probes,omitzero"`
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
	// GPU memory trigger configuration. Calculated by `DCGM_FI_DEV_MEM_COPY_UTIL`
	// metric
	GPUMemory InferenceDeploymentUpdateParamsContainerScaleTriggersGPUMemory `json:"gpu_memory,omitzero"`
	// GPU utilization trigger configuration. Calculated by `DCGM_FI_DEV_GPU_UTIL`
	// metric
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

// GPU memory trigger configuration. Calculated by `DCGM_FI_DEV_MEM_COPY_UTIL`
// metric
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

// GPU utilization trigger configuration. Calculated by `DCGM_FI_DEV_GPU_UTIL`
// metric
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

// Ingress options for the inference instance
type InferenceDeploymentUpdateParamsIngressOpts struct {
	// Disable response buffering if true. A client usually has a much slower
	// connection and can not consume the response data as fast as it is produced by an
	// upstream application. Ingress tries to buffer the whole response in order to
	// release the upstream application as soon as possible.By default, the response
	// buffering is enabled.
	DisableResponseBuffering param.Opt[bool] `json:"disable_response_buffering,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsIngressOpts) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsIngressOpts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsIngressOpts) UnmarshalJSON(data []byte) error {
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
	LivenessProbe InferenceDeploymentUpdateParamsProbesLivenessProbe `json:"liveness_probe,omitzero"`
	// Readiness probe configuration
	ReadinessProbe InferenceDeploymentUpdateParamsProbesReadinessProbe `json:"readiness_probe,omitzero"`
	// Startup probe configuration
	StartupProbe InferenceDeploymentUpdateParamsProbesStartupProbe `json:"startup_probe,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbes) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbes
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Liveness probe configuration
type InferenceDeploymentUpdateParamsProbesLivenessProbe struct {
	// Whether the probe is enabled or not.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Probe configuration (exec, `http_get` or `tcp_socket`)
	Probe InferenceDeploymentUpdateParamsProbesLivenessProbeProbe `json:"probe,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesLivenessProbe) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesLivenessProbe
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesLivenessProbe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Probe configuration (exec, `http_get` or `tcp_socket`)
type InferenceDeploymentUpdateParamsProbesLivenessProbeProbe struct {
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
	Exec InferenceDeploymentUpdateParamsProbesLivenessProbeProbeExec `json:"exec,omitzero"`
	// HTTP GET probe configuration
	HTTPGet InferenceDeploymentUpdateParamsProbesLivenessProbeProbeHTTPGet `json:"http_get,omitzero"`
	// TCP socket probe configuration
	TcpSocket InferenceDeploymentUpdateParamsProbesLivenessProbeProbeTcpSocket `json:"tcp_socket,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesLivenessProbeProbe) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesLivenessProbeProbe
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesLivenessProbeProbe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exec probe configuration
type InferenceDeploymentUpdateParamsProbesLivenessProbeProbeExec struct {
	// Command to be executed inside the running container.
	Command []string `json:"command,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesLivenessProbeProbeExec) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesLivenessProbeProbeExec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesLivenessProbeProbeExec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP GET probe configuration
type InferenceDeploymentUpdateParamsProbesLivenessProbeProbeHTTPGet struct {
	// Host name to send HTTP request to.
	Host param.Opt[string] `json:"host,omitzero"`
	// The endpoint to send the HTTP request to.
	Path param.Opt[string] `json:"path,omitzero"`
	// Port number the probe should connect to.
	Port param.Opt[int64] `json:"port,omitzero"`
	// Schema to use for the HTTP request.
	Schema param.Opt[string] `json:"schema,omitzero"`
	// HTTP headers to be sent with the request.
	Headers map[string]string `json:"headers,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesLivenessProbeProbeHTTPGet) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesLivenessProbeProbeHTTPGet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesLivenessProbeProbeHTTPGet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TCP socket probe configuration
type InferenceDeploymentUpdateParamsProbesLivenessProbeProbeTcpSocket struct {
	// Port number to check if it's open.
	Port param.Opt[int64] `json:"port,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesLivenessProbeProbeTcpSocket) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesLivenessProbeProbeTcpSocket
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesLivenessProbeProbeTcpSocket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Readiness probe configuration
type InferenceDeploymentUpdateParamsProbesReadinessProbe struct {
	// Whether the probe is enabled or not.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Probe configuration (exec, `http_get` or `tcp_socket`)
	Probe InferenceDeploymentUpdateParamsProbesReadinessProbeProbe `json:"probe,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesReadinessProbe) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesReadinessProbe
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesReadinessProbe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Probe configuration (exec, `http_get` or `tcp_socket`)
type InferenceDeploymentUpdateParamsProbesReadinessProbeProbe struct {
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
	Exec InferenceDeploymentUpdateParamsProbesReadinessProbeProbeExec `json:"exec,omitzero"`
	// HTTP GET probe configuration
	HTTPGet InferenceDeploymentUpdateParamsProbesReadinessProbeProbeHTTPGet `json:"http_get,omitzero"`
	// TCP socket probe configuration
	TcpSocket InferenceDeploymentUpdateParamsProbesReadinessProbeProbeTcpSocket `json:"tcp_socket,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesReadinessProbeProbe) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesReadinessProbeProbe
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesReadinessProbeProbe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exec probe configuration
type InferenceDeploymentUpdateParamsProbesReadinessProbeProbeExec struct {
	// Command to be executed inside the running container.
	Command []string `json:"command,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesReadinessProbeProbeExec) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesReadinessProbeProbeExec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesReadinessProbeProbeExec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP GET probe configuration
type InferenceDeploymentUpdateParamsProbesReadinessProbeProbeHTTPGet struct {
	// Host name to send HTTP request to.
	Host param.Opt[string] `json:"host,omitzero"`
	// The endpoint to send the HTTP request to.
	Path param.Opt[string] `json:"path,omitzero"`
	// Port number the probe should connect to.
	Port param.Opt[int64] `json:"port,omitzero"`
	// Schema to use for the HTTP request.
	Schema param.Opt[string] `json:"schema,omitzero"`
	// HTTP headers to be sent with the request.
	Headers map[string]string `json:"headers,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesReadinessProbeProbeHTTPGet) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesReadinessProbeProbeHTTPGet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesReadinessProbeProbeHTTPGet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TCP socket probe configuration
type InferenceDeploymentUpdateParamsProbesReadinessProbeProbeTcpSocket struct {
	// Port number to check if it's open.
	Port param.Opt[int64] `json:"port,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesReadinessProbeProbeTcpSocket) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesReadinessProbeProbeTcpSocket
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesReadinessProbeProbeTcpSocket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Startup probe configuration
type InferenceDeploymentUpdateParamsProbesStartupProbe struct {
	// Whether the probe is enabled or not.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Probe configuration (exec, `http_get` or `tcp_socket`)
	Probe InferenceDeploymentUpdateParamsProbesStartupProbeProbe `json:"probe,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesStartupProbe) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesStartupProbe
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesStartupProbe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Probe configuration (exec, `http_get` or `tcp_socket`)
type InferenceDeploymentUpdateParamsProbesStartupProbeProbe struct {
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
	Exec InferenceDeploymentUpdateParamsProbesStartupProbeProbeExec `json:"exec,omitzero"`
	// HTTP GET probe configuration
	HTTPGet InferenceDeploymentUpdateParamsProbesStartupProbeProbeHTTPGet `json:"http_get,omitzero"`
	// TCP socket probe configuration
	TcpSocket InferenceDeploymentUpdateParamsProbesStartupProbeProbeTcpSocket `json:"tcp_socket,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesStartupProbeProbe) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesStartupProbeProbe
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesStartupProbeProbe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Exec probe configuration
type InferenceDeploymentUpdateParamsProbesStartupProbeProbeExec struct {
	// Command to be executed inside the running container.
	Command []string `json:"command,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesStartupProbeProbeExec) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesStartupProbeProbeExec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesStartupProbeProbeExec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP GET probe configuration
type InferenceDeploymentUpdateParamsProbesStartupProbeProbeHTTPGet struct {
	// Host name to send HTTP request to.
	Host param.Opt[string] `json:"host,omitzero"`
	// The endpoint to send the HTTP request to.
	Path param.Opt[string] `json:"path,omitzero"`
	// Port number the probe should connect to.
	Port param.Opt[int64] `json:"port,omitzero"`
	// Schema to use for the HTTP request.
	Schema param.Opt[string] `json:"schema,omitzero"`
	// HTTP headers to be sent with the request.
	Headers map[string]string `json:"headers,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesStartupProbeProbeHTTPGet) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesStartupProbeProbeHTTPGet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesStartupProbeProbeHTTPGet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TCP socket probe configuration
type InferenceDeploymentUpdateParamsProbesStartupProbeProbeTcpSocket struct {
	// Port number to check if it's open.
	Port param.Opt[int64] `json:"port,omitzero"`
	paramObj
}

func (r InferenceDeploymentUpdateParamsProbesStartupProbeProbeTcpSocket) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbesStartupProbeProbeTcpSocket
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceDeploymentUpdateParamsProbesStartupProbeProbeTcpSocket) UnmarshalJSON(data []byte) error {
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
