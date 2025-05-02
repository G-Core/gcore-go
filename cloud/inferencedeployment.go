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
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.IsPresent() {
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
	if !params.ProjectID.IsPresent() {
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
	if !params.ProjectID.IsPresent() {
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
	if !body.ProjectID.IsPresent() {
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
	if !query.ProjectID.IsPresent() {
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
	if !query.ProjectID.IsPresent() {
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
	if !body.ProjectID.IsPresent() {
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
	if !body.ProjectID.IsPresent() {
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

// '#/components/schemas/ContainerOutSerializerV3'
// "$.components.schemas.ContainerOutSerializerV3"
type Container struct {
	// '#/components/schemas/ContainerOutSerializerV3/properties/address/anyOf/0'
	// "$.components.schemas.ContainerOutSerializerV3.properties.address.anyOf[0]"
	Address string `json:"address,required" format:"uri"`
	// '#/components/schemas/ContainerOutSerializerV3/properties/deploy_status'
	// "$.components.schemas.ContainerOutSerializerV3.properties.deploy_status"
	DeployStatus DeployStatus `json:"deploy_status,required"`
	// '#/components/schemas/ContainerOutSerializerV3/properties/error_message/anyOf/0'
	// "$.components.schemas.ContainerOutSerializerV3.properties.error_message.anyOf[0]"
	ErrorMessage string `json:"error_message,required"`
	// '#/components/schemas/ContainerOutSerializerV3/properties/region_id'
	// "$.components.schemas.ContainerOutSerializerV3.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/ContainerOutSerializerV3/properties/scale'
	// "$.components.schemas.ContainerOutSerializerV3.properties.scale"
	Scale ContainerScale `json:"scale,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Address      resp.Field
		DeployStatus resp.Field
		ErrorMessage resp.Field
		RegionID     resp.Field
		Scale        resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Container) RawJSON() string { return r.JSON.raw }
func (r *Container) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InferenceInstanceOutSerializerV3'
// "$.components.schemas.InferenceInstanceOutSerializerV3"
type Inference struct {
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/address/anyOf/0'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.address.anyOf[0]"
	Address string `json:"address,required" format:"uri"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/auth_enabled'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.auth_enabled"
	AuthEnabled bool `json:"auth_enabled,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/command/anyOf/0'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.command.anyOf[0]"
	Command string `json:"command,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/containers'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.containers"
	Containers []Container `json:"containers,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/created_at/anyOf/0'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.created_at.anyOf[0]"
	CreatedAt string `json:"created_at,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/credentials_name'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.credentials_name"
	CredentialsName string `json:"credentials_name,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/description'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.description"
	Description string `json:"description,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/envs/anyOf/0'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.envs.anyOf[0]"
	Envs map[string]string `json:"envs,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/flavor_name'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.flavor_name"
	FlavorName string `json:"flavor_name,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/image'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.image"
	Image string `json:"image,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/ingress_opts/anyOf/0'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.ingress_opts.anyOf[0]"
	IngressOpts IngressOptsOut `json:"ingress_opts,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/listening_port'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.listening_port"
	ListeningPort int64 `json:"listening_port,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/logging/anyOf/0'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.logging.anyOf[0]"
	Logging Logging `json:"logging,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/name'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/probes/anyOf/0'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.probes.anyOf[0]"
	Probes InferenceProbes `json:"probes,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/project_id'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/status'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.status"
	//
	// Any of "ACTIVE", "DELETING", "DEPLOYING", "DISABLED", "PARTIALLYDEPLOYED".
	Status InferenceStatus `json:"status,required"`
	// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/timeout/anyOf/0'
	// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.timeout.anyOf[0]"
	Timeout int64 `json:"timeout,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Address         resp.Field
		AuthEnabled     resp.Field
		Command         resp.Field
		Containers      resp.Field
		CreatedAt       resp.Field
		CredentialsName resp.Field
		Description     resp.Field
		Envs            resp.Field
		FlavorName      resp.Field
		Image           resp.Field
		IngressOpts     resp.Field
		ListeningPort   resp.Field
		Logging         resp.Field
		Name            resp.Field
		Probes          resp.Field
		ProjectID       resp.Field
		Status          resp.Field
		Timeout         resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Inference) RawJSON() string { return r.JSON.raw }
func (r *Inference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InferenceInstanceOutSerializerV3/properties/status'
// "$.components.schemas.InferenceInstanceOutSerializerV3.properties.status"
type InferenceStatus string

const (
	InferenceStatusActive            InferenceStatus = "ACTIVE"
	InferenceStatusDeleting          InferenceStatus = "DELETING"
	InferenceStatusDeploying         InferenceStatus = "DEPLOYING"
	InferenceStatusDisabled          InferenceStatus = "DISABLED"
	InferenceStatusPartiallydeployed InferenceStatus = "PARTIALLYDEPLOYED"
)

// '#/components/schemas/InferenceInstanceApikeySecretSerializerV3'
// "$.components.schemas.InferenceInstanceApikeySecretSerializerV3"
type InferenceApikeySecret struct {
	// '#/components/schemas/InferenceInstanceApikeySecretSerializerV3/properties/secret'
	// "$.components.schemas.InferenceInstanceApikeySecretSerializerV3.properties.secret"
	Secret string `json:"secret,required"`
	// '#/components/schemas/InferenceInstanceApikeySecretSerializerV3/properties/status'
	// "$.components.schemas.InferenceInstanceApikeySecretSerializerV3.properties.status"
	//
	// Any of "PENDING", "READY".
	Status InferenceApikeySecretStatus `json:"status,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Secret      resp.Field
		Status      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApikeySecret) RawJSON() string { return r.JSON.raw }
func (r *InferenceApikeySecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InferenceInstanceApikeySecretSerializerV3/properties/status'
// "$.components.schemas.InferenceInstanceApikeySecretSerializerV3.properties.status"
type InferenceApikeySecretStatus string

const (
	InferenceApikeySecretStatusPending InferenceApikeySecretStatus = "PENDING"
	InferenceApikeySecretStatusReady   InferenceApikeySecretStatus = "READY"
)

// '#/components/schemas/InferenceInstanceLogSerializerV3'
// "$.components.schemas.InferenceInstanceLogSerializerV3"
type InferenceLog struct {
	// '#/components/schemas/InferenceInstanceLogSerializerV3/properties/message'
	// "$.components.schemas.InferenceInstanceLogSerializerV3.properties.message"
	Message string `json:"message,required"`
	// '#/components/schemas/InferenceInstanceLogSerializerV3/properties/pod'
	// "$.components.schemas.InferenceInstanceLogSerializerV3.properties.pod"
	Pod string `json:"pod,required"`
	// '#/components/schemas/InferenceInstanceLogSerializerV3/properties/region_id'
	// "$.components.schemas.InferenceInstanceLogSerializerV3.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/InferenceInstanceLogSerializerV3/properties/time'
	// "$.components.schemas.InferenceInstanceLogSerializerV3.properties.time"
	Time time.Time `json:"time,required" format:"date-time"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Message     resp.Field
		Pod         resp.Field
		RegionID    resp.Field
		Time        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceLog) RawJSON() string { return r.JSON.raw }
func (r *InferenceLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceDeploymentNewParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments/post/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/components/schemas/InferenceInstanceInSerializerV3/properties/containers'
	// "$.components.schemas.InferenceInstanceInSerializerV3.properties.containers"
	Containers []InferenceDeploymentNewParamsContainer `json:"containers,omitzero,required"`
	// '#/components/schemas/InferenceInstanceInSerializerV3/properties/flavor_name'
	// "$.components.schemas.InferenceInstanceInSerializerV3.properties.flavor_name"
	FlavorName string `json:"flavor_name,required"`
	// '#/components/schemas/InferenceInstanceInSerializerV3/properties/image'
	// "$.components.schemas.InferenceInstanceInSerializerV3.properties.image"
	Image string `json:"image,required"`
	// '#/components/schemas/InferenceInstanceInSerializerV3/properties/listening_port'
	// "$.components.schemas.InferenceInstanceInSerializerV3.properties.listening_port"
	ListeningPort int64 `json:"listening_port,required"`
	// '#/components/schemas/InferenceInstanceInSerializerV3/properties/name'
	// "$.components.schemas.InferenceInstanceInSerializerV3.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/InferenceInstanceInSerializerV3/properties/credentials_name/anyOf/0'
	// "$.components.schemas.InferenceInstanceInSerializerV3.properties.credentials_name.anyOf[0]"
	CredentialsName param.Opt[string] `json:"credentials_name,omitzero"`
	// '#/components/schemas/InferenceInstanceInSerializerV3/properties/description/anyOf/0'
	// "$.components.schemas.InferenceInstanceInSerializerV3.properties.description.anyOf[0]"
	Description param.Opt[string] `json:"description,omitzero"`
	// '#/components/schemas/InferenceInstanceInSerializerV3/properties/timeout/anyOf/0'
	// "$.components.schemas.InferenceInstanceInSerializerV3.properties.timeout.anyOf[0]"
	Timeout param.Opt[int64] `json:"timeout,omitzero"`
	// '#/components/schemas/InferenceInstanceInSerializerV3/properties/auth_enabled'
	// "$.components.schemas.InferenceInstanceInSerializerV3.properties.auth_enabled"
	AuthEnabled param.Opt[bool] `json:"auth_enabled,omitzero"`
	// '#/components/schemas/InferenceInstanceInSerializerV3/properties/command/anyOf/0'
	// "$.components.schemas.InferenceInstanceInSerializerV3.properties.command.anyOf[0]"
	Command []string `json:"command,omitzero"`
	// '#/components/schemas/InferenceInstanceInSerializerV3/properties/logging/anyOf/0'
	// "$.components.schemas.InferenceInstanceInSerializerV3.properties.logging.anyOf[0]"
	Logging InferenceDeploymentNewParamsLogging `json:"logging,omitzero"`
	// '#/components/schemas/InferenceInstanceInSerializerV3/properties/probes/anyOf/0'
	// "$.components.schemas.InferenceInstanceInSerializerV3.properties.probes.anyOf[0]"
	Probes InferenceDeploymentNewParamsProbes `json:"probes,omitzero"`
	// '#/components/schemas/InferenceInstanceInSerializerV3/properties/envs'
	// "$.components.schemas.InferenceInstanceInSerializerV3.properties.envs"
	Envs map[string]string `json:"envs,omitzero"`
	// '#/components/schemas/InferenceInstanceInSerializerV3/properties/ingress_opts/anyOf/0'
	// "$.components.schemas.InferenceInstanceInSerializerV3.properties.ingress_opts.anyOf[0]"
	IngressOpts IngressOptsParam `json:"ingress_opts,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InferenceDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/InferenceInstanceInSerializerV3/properties/containers/items'
// "$.components.schemas.InferenceInstanceInSerializerV3.properties.containers.items"
//
// The properties RegionID, Scale are required.
type InferenceDeploymentNewParamsContainer struct {
	// '#/components/schemas/ContainerInSerializerV3/properties/region_id'
	// "$.components.schemas.ContainerInSerializerV3.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/ContainerInSerializerV3/properties/scale'
	// "$.components.schemas.ContainerInSerializerV3.properties.scale"
	Scale InferenceDeploymentNewParamsContainerScale `json:"scale,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentNewParamsContainer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentNewParamsContainer) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainer
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerInSerializerV3/properties/scale'
// "$.components.schemas.ContainerInSerializerV3.properties.scale"
//
// The properties Max, Min are required.
type InferenceDeploymentNewParamsContainerScale struct {
	// '#/components/schemas/ContainerScaleSerializerV3/properties/max'
	// "$.components.schemas.ContainerScaleSerializerV3.properties.max"
	Max int64 `json:"max,required"`
	// '#/components/schemas/ContainerScaleSerializerV3/properties/min'
	// "$.components.schemas.ContainerScaleSerializerV3.properties.min"
	Min int64 `json:"min,required"`
	// '#/components/schemas/ContainerScaleSerializerV3/properties/cooldown_period/anyOf/0'
	// "$.components.schemas.ContainerScaleSerializerV3.properties.cooldown_period.anyOf[0]"
	CooldownPeriod param.Opt[int64] `json:"cooldown_period,omitzero"`
	// '#/components/schemas/ContainerScaleSerializerV3/properties/polling_interval/anyOf/0'
	// "$.components.schemas.ContainerScaleSerializerV3.properties.polling_interval.anyOf[0]"
	PollingInterval param.Opt[int64] `json:"polling_interval,omitzero"`
	// '#/components/schemas/ContainerScaleSerializerV3/properties/triggers'
	// "$.components.schemas.ContainerScaleSerializerV3.properties.triggers"
	Triggers InferenceDeploymentNewParamsContainerScaleTriggers `json:"triggers,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentNewParamsContainerScale) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentNewParamsContainerScale) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScale
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleSerializerV3/properties/triggers'
// "$.components.schemas.ContainerScaleSerializerV3.properties.triggers"
type InferenceDeploymentNewParamsContainerScaleTriggers struct {
	// '#/components/schemas/ContainerScaleTriggersSerializer/properties/cpu/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSerializer.properties.cpu.anyOf[0]"
	CPU InferenceDeploymentNewParamsContainerScaleTriggersCPU `json:"cpu,omitzero"`
	// '#/components/schemas/ContainerScaleTriggersSerializer/properties/gpu_memory/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSerializer.properties.gpu_memory.anyOf[0]"
	GPUMemory InferenceDeploymentNewParamsContainerScaleTriggersGPUMemory `json:"gpu_memory,omitzero"`
	// '#/components/schemas/ContainerScaleTriggersSerializer/properties/gpu_utilization/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSerializer.properties.gpu_utilization.anyOf[0]"
	GPUUtilization InferenceDeploymentNewParamsContainerScaleTriggersGPUUtilization `json:"gpu_utilization,omitzero"`
	// '#/components/schemas/ContainerScaleTriggersSerializer/properties/http/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSerializer.properties.http.anyOf[0]"
	HTTP InferenceDeploymentNewParamsContainerScaleTriggersHTTP `json:"http,omitzero"`
	// '#/components/schemas/ContainerScaleTriggersSerializer/properties/memory/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSerializer.properties.memory.anyOf[0]"
	Memory InferenceDeploymentNewParamsContainerScaleTriggersMemory `json:"memory,omitzero"`
	// '#/components/schemas/ContainerScaleTriggersSerializer/properties/sqs/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSerializer.properties.sqs.anyOf[0]"
	Sqs InferenceDeploymentNewParamsContainerScaleTriggersSqs `json:"sqs,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentNewParamsContainerScaleTriggers) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentNewParamsContainerScaleTriggers) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScaleTriggers
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleTriggersSerializer/properties/cpu/anyOf/0'
// "$.components.schemas.ContainerScaleTriggersSerializer.properties.cpu.anyOf[0]"
//
// The property Threshold is required.
type InferenceDeploymentNewParamsContainerScaleTriggersCPU struct {
	// '#/components/schemas/ContainerScaleTriggersThresholdSerializer/properties/threshold'
	// "$.components.schemas.ContainerScaleTriggersThresholdSerializer.properties.threshold"
	Threshold int64 `json:"threshold,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentNewParamsContainerScaleTriggersCPU) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentNewParamsContainerScaleTriggersCPU) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScaleTriggersCPU
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleTriggersSerializer/properties/gpu_memory/anyOf/0'
// "$.components.schemas.ContainerScaleTriggersSerializer.properties.gpu_memory.anyOf[0]"
//
// The property Threshold is required.
type InferenceDeploymentNewParamsContainerScaleTriggersGPUMemory struct {
	// '#/components/schemas/ContainerScaleTriggersThresholdSerializer/properties/threshold'
	// "$.components.schemas.ContainerScaleTriggersThresholdSerializer.properties.threshold"
	Threshold int64 `json:"threshold,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentNewParamsContainerScaleTriggersGPUMemory) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentNewParamsContainerScaleTriggersGPUMemory) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScaleTriggersGPUMemory
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleTriggersSerializer/properties/gpu_utilization/anyOf/0'
// "$.components.schemas.ContainerScaleTriggersSerializer.properties.gpu_utilization.anyOf[0]"
//
// The property Threshold is required.
type InferenceDeploymentNewParamsContainerScaleTriggersGPUUtilization struct {
	// '#/components/schemas/ContainerScaleTriggersThresholdSerializer/properties/threshold'
	// "$.components.schemas.ContainerScaleTriggersThresholdSerializer.properties.threshold"
	Threshold int64 `json:"threshold,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentNewParamsContainerScaleTriggersGPUUtilization) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentNewParamsContainerScaleTriggersGPUUtilization) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScaleTriggersGPUUtilization
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleTriggersSerializer/properties/http/anyOf/0'
// "$.components.schemas.ContainerScaleTriggersSerializer.properties.http.anyOf[0]"
//
// The properties Rate, Window are required.
type InferenceDeploymentNewParamsContainerScaleTriggersHTTP struct {
	// '#/components/schemas/ContainerScaleTriggersRateSerializer/properties/rate'
	// "$.components.schemas.ContainerScaleTriggersRateSerializer.properties.rate"
	Rate int64 `json:"rate,required"`
	// '#/components/schemas/ContainerScaleTriggersRateSerializer/properties/window'
	// "$.components.schemas.ContainerScaleTriggersRateSerializer.properties.window"
	Window int64 `json:"window,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentNewParamsContainerScaleTriggersHTTP) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentNewParamsContainerScaleTriggersHTTP) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScaleTriggersHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleTriggersSerializer/properties/memory/anyOf/0'
// "$.components.schemas.ContainerScaleTriggersSerializer.properties.memory.anyOf[0]"
//
// The property Threshold is required.
type InferenceDeploymentNewParamsContainerScaleTriggersMemory struct {
	// '#/components/schemas/ContainerScaleTriggersThresholdSerializer/properties/threshold'
	// "$.components.schemas.ContainerScaleTriggersThresholdSerializer.properties.threshold"
	Threshold int64 `json:"threshold,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentNewParamsContainerScaleTriggersMemory) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentNewParamsContainerScaleTriggersMemory) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScaleTriggersMemory
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleTriggersSerializer/properties/sqs/anyOf/0'
// "$.components.schemas.ContainerScaleTriggersSerializer.properties.sqs.anyOf[0]"
//
// The properties ActivationQueueLength, AwsRegion, QueueLength, QueueURL,
// SecretName are required.
type InferenceDeploymentNewParamsContainerScaleTriggersSqs struct {
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/activation_queue_length'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.activation_queue_length"
	ActivationQueueLength int64 `json:"activation_queue_length,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/aws_region'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.aws_region"
	AwsRegion string `json:"aws_region,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/queue_length'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.queue_length"
	QueueLength int64 `json:"queue_length,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/queue_url'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.queue_url"
	QueueURL string `json:"queue_url,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/secret_name'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.secret_name"
	SecretName string `json:"secret_name,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/aws_endpoint/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.aws_endpoint.anyOf[0]"
	AwsEndpoint param.Opt[string] `json:"aws_endpoint,omitzero"`
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/scale_on_delayed'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.scale_on_delayed"
	ScaleOnDelayed param.Opt[bool] `json:"scale_on_delayed,omitzero"`
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/scale_on_flight'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.scale_on_flight"
	ScaleOnFlight param.Opt[bool] `json:"scale_on_flight,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentNewParamsContainerScaleTriggersSqs) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentNewParamsContainerScaleTriggersSqs) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsContainerScaleTriggersSqs
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/InferenceInstanceInSerializerV3/properties/logging/anyOf/0'
// "$.components.schemas.InferenceInstanceInSerializerV3.properties.logging.anyOf[0]"
type InferenceDeploymentNewParamsLogging struct {
	// '#/components/schemas/LoggingInSerializer/properties/destination_region_id/anyOf/0'
	// "$.components.schemas.LoggingInSerializer.properties.destination_region_id.anyOf[0]"
	DestinationRegionID param.Opt[int64] `json:"destination_region_id,omitzero"`
	// '#/components/schemas/LoggingInSerializer/properties/topic_name/anyOf/0'
	// "$.components.schemas.LoggingInSerializer.properties.topic_name.anyOf[0]"
	TopicName param.Opt[string] `json:"topic_name,omitzero"`
	// '#/components/schemas/LoggingInSerializer/properties/enabled'
	// "$.components.schemas.LoggingInSerializer.properties.enabled"
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// '#/components/schemas/LoggingInSerializer/properties/retention_policy/anyOf/0'
	// "$.components.schemas.LoggingInSerializer.properties.retention_policy.anyOf[0]"
	RetentionPolicy LaasIndexRetentionPolicyParam `json:"retention_policy,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentNewParamsLogging) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentNewParamsLogging) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsLogging
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/InferenceInstanceInSerializerV3/properties/probes/anyOf/0'
// "$.components.schemas.InferenceInstanceInSerializerV3.properties.probes.anyOf[0]"
type InferenceDeploymentNewParamsProbes struct {
	// '#/components/schemas/InferenceInstanceProbesSerializerV2/properties/liveness_probe/anyOf/0'
	// "$.components.schemas.InferenceInstanceProbesSerializerV2.properties.liveness_probe.anyOf[0]"
	LivenessProbe ContainerProbeConfigCreateParam `json:"liveness_probe,omitzero"`
	// '#/components/schemas/InferenceInstanceProbesSerializerV2/properties/readiness_probe/anyOf/0'
	// "$.components.schemas.InferenceInstanceProbesSerializerV2.properties.readiness_probe.anyOf[0]"
	ReadinessProbe ContainerProbeConfigCreateParam `json:"readiness_probe,omitzero"`
	// '#/components/schemas/InferenceInstanceProbesSerializerV2/properties/startup_probe/anyOf/0'
	// "$.components.schemas.InferenceInstanceProbesSerializerV2.properties.startup_probe.anyOf[0]"
	StartupProbe ContainerProbeConfigCreateParam `json:"startup_probe,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentNewParamsProbes) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentNewParamsProbes) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentNewParamsProbes
	return param.MarshalObject(r, (*shadow)(&r))
}

type InferenceDeploymentUpdateParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments%2F%7Bdeployment_name%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments/{deployment_name}'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/auth_enabled/anyOf/0'
	// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.auth_enabled.anyOf[0]"
	AuthEnabled param.Opt[bool] `json:"auth_enabled,omitzero"`
	// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/credentials_name/anyOf/0'
	// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.credentials_name.anyOf[0]"
	CredentialsName param.Opt[string] `json:"credentials_name,omitzero"`
	// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/description/anyOf/0'
	// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.description.anyOf[0]"
	Description param.Opt[string] `json:"description,omitzero"`
	// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/flavor_name/anyOf/0'
	// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.flavor_name.anyOf[0]"
	FlavorName param.Opt[string] `json:"flavor_name,omitzero"`
	// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/image/anyOf/0'
	// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.image.anyOf[0]"
	Image param.Opt[string] `json:"image,omitzero"`
	// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/listening_port/anyOf/0'
	// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.listening_port.anyOf[0]"
	ListeningPort param.Opt[int64] `json:"listening_port,omitzero"`
	// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/timeout/anyOf/0'
	// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.timeout.anyOf[0]"
	Timeout param.Opt[int64] `json:"timeout,omitzero"`
	// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/command/anyOf/0'
	// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.command.anyOf[0]"
	Command []string `json:"command,omitzero"`
	// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/containers/anyOf/0'
	// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.containers.anyOf[0]"
	Containers []InferenceDeploymentUpdateParamsContainer `json:"containers,omitzero"`
	// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/envs/anyOf/0'
	// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.envs.anyOf[0]"
	Envs map[string]string `json:"envs,omitzero"`
	// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/logging/anyOf/0'
	// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.logging.anyOf[0]"
	Logging InferenceDeploymentUpdateParamsLogging `json:"logging,omitzero"`
	// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/probes/anyOf/0'
	// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.probes.anyOf[0]"
	Probes InferenceDeploymentUpdateParamsProbes `json:"probes,omitzero"`
	// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/ingress_opts/anyOf/0'
	// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.ingress_opts.anyOf[0]"
	IngressOpts IngressOptsParam `json:"ingress_opts,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r InferenceDeploymentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/containers/anyOf/0/items'
// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.containers.anyOf[0].items"
//
// The properties RegionID, Scale are required.
type InferenceDeploymentUpdateParamsContainer struct {
	// '#/components/schemas/ContainerInUpdateSerializerV3/properties/region_id/anyOf/0'
	// "$.components.schemas.ContainerInUpdateSerializerV3.properties.region_id.anyOf[0]"
	RegionID param.Opt[int64] `json:"region_id,omitzero,required"`
	// '#/components/schemas/ContainerInUpdateSerializerV3/properties/scale/anyOf/0'
	// "$.components.schemas.ContainerInUpdateSerializerV3.properties.scale.anyOf[0]"
	Scale InferenceDeploymentUpdateParamsContainerScale `json:"scale,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentUpdateParamsContainer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentUpdateParamsContainer) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainer
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerInUpdateSerializerV3/properties/scale/anyOf/0'
// "$.components.schemas.ContainerInUpdateSerializerV3.properties.scale.anyOf[0]"
//
// The properties Max, Min are required.
type InferenceDeploymentUpdateParamsContainerScale struct {
	// '#/components/schemas/ContainerScaleUpdateSerializerV3/properties/max'
	// "$.components.schemas.ContainerScaleUpdateSerializerV3.properties.max"
	Max int64 `json:"max,required"`
	// '#/components/schemas/ContainerScaleUpdateSerializerV3/properties/min'
	// "$.components.schemas.ContainerScaleUpdateSerializerV3.properties.min"
	Min int64 `json:"min,required"`
	// '#/components/schemas/ContainerScaleUpdateSerializerV3/properties/cooldown_period'
	// "$.components.schemas.ContainerScaleUpdateSerializerV3.properties.cooldown_period"
	CooldownPeriod param.Opt[int64] `json:"cooldown_period,omitzero"`
	// '#/components/schemas/ContainerScaleUpdateSerializerV3/properties/polling_interval'
	// "$.components.schemas.ContainerScaleUpdateSerializerV3.properties.polling_interval"
	PollingInterval param.Opt[int64] `json:"polling_interval,omitzero"`
	// '#/components/schemas/ContainerScaleUpdateSerializerV3/properties/triggers'
	// "$.components.schemas.ContainerScaleUpdateSerializerV3.properties.triggers"
	Triggers InferenceDeploymentUpdateParamsContainerScaleTriggers `json:"triggers,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentUpdateParamsContainerScale) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentUpdateParamsContainerScale) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScale
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleUpdateSerializerV3/properties/triggers'
// "$.components.schemas.ContainerScaleUpdateSerializerV3.properties.triggers"
type InferenceDeploymentUpdateParamsContainerScaleTriggers struct {
	// '#/components/schemas/ContainerScaleTriggersSerializer/properties/cpu/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSerializer.properties.cpu.anyOf[0]"
	CPU InferenceDeploymentUpdateParamsContainerScaleTriggersCPU `json:"cpu,omitzero"`
	// '#/components/schemas/ContainerScaleTriggersSerializer/properties/gpu_memory/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSerializer.properties.gpu_memory.anyOf[0]"
	GPUMemory InferenceDeploymentUpdateParamsContainerScaleTriggersGPUMemory `json:"gpu_memory,omitzero"`
	// '#/components/schemas/ContainerScaleTriggersSerializer/properties/gpu_utilization/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSerializer.properties.gpu_utilization.anyOf[0]"
	GPUUtilization InferenceDeploymentUpdateParamsContainerScaleTriggersGPUUtilization `json:"gpu_utilization,omitzero"`
	// '#/components/schemas/ContainerScaleTriggersSerializer/properties/http/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSerializer.properties.http.anyOf[0]"
	HTTP InferenceDeploymentUpdateParamsContainerScaleTriggersHTTP `json:"http,omitzero"`
	// '#/components/schemas/ContainerScaleTriggersSerializer/properties/memory/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSerializer.properties.memory.anyOf[0]"
	Memory InferenceDeploymentUpdateParamsContainerScaleTriggersMemory `json:"memory,omitzero"`
	// '#/components/schemas/ContainerScaleTriggersSerializer/properties/sqs/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSerializer.properties.sqs.anyOf[0]"
	Sqs InferenceDeploymentUpdateParamsContainerScaleTriggersSqs `json:"sqs,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentUpdateParamsContainerScaleTriggers) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentUpdateParamsContainerScaleTriggers) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScaleTriggers
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleTriggersSerializer/properties/cpu/anyOf/0'
// "$.components.schemas.ContainerScaleTriggersSerializer.properties.cpu.anyOf[0]"
//
// The property Threshold is required.
type InferenceDeploymentUpdateParamsContainerScaleTriggersCPU struct {
	// '#/components/schemas/ContainerScaleTriggersThresholdSerializer/properties/threshold'
	// "$.components.schemas.ContainerScaleTriggersThresholdSerializer.properties.threshold"
	Threshold int64 `json:"threshold,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentUpdateParamsContainerScaleTriggersCPU) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentUpdateParamsContainerScaleTriggersCPU) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScaleTriggersCPU
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleTriggersSerializer/properties/gpu_memory/anyOf/0'
// "$.components.schemas.ContainerScaleTriggersSerializer.properties.gpu_memory.anyOf[0]"
//
// The property Threshold is required.
type InferenceDeploymentUpdateParamsContainerScaleTriggersGPUMemory struct {
	// '#/components/schemas/ContainerScaleTriggersThresholdSerializer/properties/threshold'
	// "$.components.schemas.ContainerScaleTriggersThresholdSerializer.properties.threshold"
	Threshold int64 `json:"threshold,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentUpdateParamsContainerScaleTriggersGPUMemory) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentUpdateParamsContainerScaleTriggersGPUMemory) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScaleTriggersGPUMemory
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleTriggersSerializer/properties/gpu_utilization/anyOf/0'
// "$.components.schemas.ContainerScaleTriggersSerializer.properties.gpu_utilization.anyOf[0]"
//
// The property Threshold is required.
type InferenceDeploymentUpdateParamsContainerScaleTriggersGPUUtilization struct {
	// '#/components/schemas/ContainerScaleTriggersThresholdSerializer/properties/threshold'
	// "$.components.schemas.ContainerScaleTriggersThresholdSerializer.properties.threshold"
	Threshold int64 `json:"threshold,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentUpdateParamsContainerScaleTriggersGPUUtilization) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentUpdateParamsContainerScaleTriggersGPUUtilization) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScaleTriggersGPUUtilization
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleTriggersSerializer/properties/http/anyOf/0'
// "$.components.schemas.ContainerScaleTriggersSerializer.properties.http.anyOf[0]"
//
// The properties Rate, Window are required.
type InferenceDeploymentUpdateParamsContainerScaleTriggersHTTP struct {
	// '#/components/schemas/ContainerScaleTriggersRateSerializer/properties/rate'
	// "$.components.schemas.ContainerScaleTriggersRateSerializer.properties.rate"
	Rate int64 `json:"rate,required"`
	// '#/components/schemas/ContainerScaleTriggersRateSerializer/properties/window'
	// "$.components.schemas.ContainerScaleTriggersRateSerializer.properties.window"
	Window int64 `json:"window,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentUpdateParamsContainerScaleTriggersHTTP) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentUpdateParamsContainerScaleTriggersHTTP) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScaleTriggersHTTP
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleTriggersSerializer/properties/memory/anyOf/0'
// "$.components.schemas.ContainerScaleTriggersSerializer.properties.memory.anyOf[0]"
//
// The property Threshold is required.
type InferenceDeploymentUpdateParamsContainerScaleTriggersMemory struct {
	// '#/components/schemas/ContainerScaleTriggersThresholdSerializer/properties/threshold'
	// "$.components.schemas.ContainerScaleTriggersThresholdSerializer.properties.threshold"
	Threshold int64 `json:"threshold,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentUpdateParamsContainerScaleTriggersMemory) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentUpdateParamsContainerScaleTriggersMemory) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScaleTriggersMemory
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleTriggersSerializer/properties/sqs/anyOf/0'
// "$.components.schemas.ContainerScaleTriggersSerializer.properties.sqs.anyOf[0]"
//
// The properties ActivationQueueLength, AwsRegion, QueueLength, QueueURL,
// SecretName are required.
type InferenceDeploymentUpdateParamsContainerScaleTriggersSqs struct {
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/activation_queue_length'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.activation_queue_length"
	ActivationQueueLength int64 `json:"activation_queue_length,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/aws_region'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.aws_region"
	AwsRegion string `json:"aws_region,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/queue_length'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.queue_length"
	QueueLength int64 `json:"queue_length,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/queue_url'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.queue_url"
	QueueURL string `json:"queue_url,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/secret_name'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.secret_name"
	SecretName string `json:"secret_name,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/aws_endpoint/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.aws_endpoint.anyOf[0]"
	AwsEndpoint param.Opt[string] `json:"aws_endpoint,omitzero"`
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/scale_on_delayed'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.scale_on_delayed"
	ScaleOnDelayed param.Opt[bool] `json:"scale_on_delayed,omitzero"`
	// '#/components/schemas/ContainerScaleTriggersSqsSerializer/properties/scale_on_flight'
	// "$.components.schemas.ContainerScaleTriggersSqsSerializer.properties.scale_on_flight"
	ScaleOnFlight param.Opt[bool] `json:"scale_on_flight,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentUpdateParamsContainerScaleTriggersSqs) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentUpdateParamsContainerScaleTriggersSqs) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsContainerScaleTriggersSqs
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/logging/anyOf/0'
// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.logging.anyOf[0]"
type InferenceDeploymentUpdateParamsLogging struct {
	// '#/components/schemas/LoggingInSerializer/properties/destination_region_id/anyOf/0'
	// "$.components.schemas.LoggingInSerializer.properties.destination_region_id.anyOf[0]"
	DestinationRegionID param.Opt[int64] `json:"destination_region_id,omitzero"`
	// '#/components/schemas/LoggingInSerializer/properties/topic_name/anyOf/0'
	// "$.components.schemas.LoggingInSerializer.properties.topic_name.anyOf[0]"
	TopicName param.Opt[string] `json:"topic_name,omitzero"`
	// '#/components/schemas/LoggingInSerializer/properties/enabled'
	// "$.components.schemas.LoggingInSerializer.properties.enabled"
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// '#/components/schemas/LoggingInSerializer/properties/retention_policy/anyOf/0'
	// "$.components.schemas.LoggingInSerializer.properties.retention_policy.anyOf[0]"
	RetentionPolicy LaasIndexRetentionPolicyParam `json:"retention_policy,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentUpdateParamsLogging) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentUpdateParamsLogging) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsLogging
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/InferenceInstanceInUpdateSerializerV3/properties/probes/anyOf/0'
// "$.components.schemas.InferenceInstanceInUpdateSerializerV3.properties.probes.anyOf[0]"
type InferenceDeploymentUpdateParamsProbes struct {
	// '#/components/schemas/InferenceInstanceProbesSerializerV2/properties/liveness_probe/anyOf/0'
	// "$.components.schemas.InferenceInstanceProbesSerializerV2.properties.liveness_probe.anyOf[0]"
	LivenessProbe ContainerProbeConfigCreateParam `json:"liveness_probe,omitzero"`
	// '#/components/schemas/InferenceInstanceProbesSerializerV2/properties/readiness_probe/anyOf/0'
	// "$.components.schemas.InferenceInstanceProbesSerializerV2.properties.readiness_probe.anyOf[0]"
	ReadinessProbe ContainerProbeConfigCreateParam `json:"readiness_probe,omitzero"`
	// '#/components/schemas/InferenceInstanceProbesSerializerV2/properties/startup_probe/anyOf/0'
	// "$.components.schemas.InferenceInstanceProbesSerializerV2.properties.startup_probe.anyOf[0]"
	StartupProbe ContainerProbeConfigCreateParam `json:"startup_probe,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentUpdateParamsProbes) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r InferenceDeploymentUpdateParamsProbes) MarshalJSON() (data []byte, err error) {
	type shadow InferenceDeploymentUpdateParamsProbes
	return param.MarshalObject(r, (*shadow)(&r))
}

type InferenceDeploymentListParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments/get/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments/get/parameters/1'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments'].get.parameters[1]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments/get/parameters/2'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments'].get.parameters[2]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [InferenceDeploymentListParams]'s query parameters as
// `url.Values`.
func (r InferenceDeploymentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InferenceDeploymentDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments%2F%7Bdeployment_name%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments/{deployment_name}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type InferenceDeploymentGetParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments%2F%7Bdeployment_name%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments/{deployment_name}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type InferenceDeploymentGetAPIKeyParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments%2F%7Bdeployment_name%7D%2Fapikey/get/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments/{deployment_name}/apikey'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentGetAPIKeyParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

type InferenceDeploymentStartParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments%2F%7Bdeployment_name%7D%2Fstart/post/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments/{deployment_name}/start'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentStartParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type InferenceDeploymentStopParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2F%7Bproject_id%7D%2Fdeployments%2F%7Bdeployment_name%7D%2Fstop/post/parameters/0/schema'
	// "$.paths['/cloud/v3/inference/{project_id}/deployments/{deployment_name}/stop'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceDeploymentStopParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
