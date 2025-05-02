// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
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

// '#/components/schemas/AwsIamData' "$.components.schemas.AwsIamData"
type AwsIamData struct {
	// '#/components/schemas/AwsIamData/properties/aws_access_key_id'
	// "$.components.schemas.AwsIamData.properties.aws_access_key_id"
	AwsAccessKeyID string `json:"aws_access_key_id,required"`
	// '#/components/schemas/AwsIamData/properties/aws_secret_access_key'
	// "$.components.schemas.AwsIamData.properties.aws_secret_access_key"
	AwsSecretAccessKey string `json:"aws_secret_access_key,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AwsAccessKeyID     resp.Field
		AwsSecretAccessKey resp.Field
		ExtraFields        map[string]resp.Field
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
// AwsIamDataParam.IsOverridden()
func (r AwsIamData) ToParam() AwsIamDataParam {
	return param.OverrideObj[AwsIamDataParam](r.RawJSON())
}

// '#/components/schemas/AwsIamData' "$.components.schemas.AwsIamData"
//
// The properties AwsAccessKeyID, AwsSecretAccessKey are required.
type AwsIamDataParam struct {
	// '#/components/schemas/AwsIamData/properties/aws_access_key_id'
	// "$.components.schemas.AwsIamData.properties.aws_access_key_id"
	AwsAccessKeyID string `json:"aws_access_key_id,required"`
	// '#/components/schemas/AwsIamData/properties/aws_secret_access_key'
	// "$.components.schemas.AwsIamData.properties.aws_secret_access_key"
	AwsSecretAccessKey string `json:"aws_secret_access_key,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f AwsIamDataParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r AwsIamDataParam) MarshalJSON() (data []byte, err error) {
	type shadow AwsIamDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CapacityDetailsSerializer'
// "$.components.schemas.CapacityDetailsSerializer"
type Capacity struct {
	// '#/components/schemas/CapacityDetailsSerializer/properties/capacity'
	// "$.components.schemas.CapacityDetailsSerializer.properties.capacity"
	Capacity int64 `json:"capacity,required"`
	// '#/components/schemas/CapacityDetailsSerializer/properties/flavor_name'
	// "$.components.schemas.CapacityDetailsSerializer.properties.flavor_name"
	FlavorName string `json:"flavor_name,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Capacity    resp.Field
		FlavorName  resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Capacity) RawJSON() string { return r.JSON.raw }
func (r *Capacity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ContainerProbeOutSerializerV2'
// "$.components.schemas.ContainerProbeOutSerializerV2"
type ContainerProbe struct {
	// '#/components/schemas/ContainerProbeOutSerializerV2/properties/exec/anyOf/0'
	// "$.components.schemas.ContainerProbeOutSerializerV2.properties.exec.anyOf[0]"
	Exec ContainerProbeExec `json:"exec,required"`
	// '#/components/schemas/ContainerProbeOutSerializerV2/properties/failure_threshold'
	// "$.components.schemas.ContainerProbeOutSerializerV2.properties.failure_threshold"
	FailureThreshold int64 `json:"failure_threshold,required"`
	// '#/components/schemas/ContainerProbeOutSerializerV2/properties/http_get/anyOf/0'
	// "$.components.schemas.ContainerProbeOutSerializerV2.properties.http_get.anyOf[0]"
	HTTPGet ContainerProbeHTTPGet `json:"http_get,required"`
	// '#/components/schemas/ContainerProbeOutSerializerV2/properties/initial_delay_seconds'
	// "$.components.schemas.ContainerProbeOutSerializerV2.properties.initial_delay_seconds"
	InitialDelaySeconds int64 `json:"initial_delay_seconds,required"`
	// '#/components/schemas/ContainerProbeOutSerializerV2/properties/period_seconds'
	// "$.components.schemas.ContainerProbeOutSerializerV2.properties.period_seconds"
	PeriodSeconds int64 `json:"period_seconds,required"`
	// '#/components/schemas/ContainerProbeOutSerializerV2/properties/success_threshold'
	// "$.components.schemas.ContainerProbeOutSerializerV2.properties.success_threshold"
	SuccessThreshold int64 `json:"success_threshold,required"`
	// '#/components/schemas/ContainerProbeOutSerializerV2/properties/tcp_socket/anyOf/0'
	// "$.components.schemas.ContainerProbeOutSerializerV2.properties.tcp_socket.anyOf[0]"
	TcpSocket ContainerProbeTcpSocket `json:"tcp_socket,required"`
	// '#/components/schemas/ContainerProbeOutSerializerV2/properties/timeout_seconds'
	// "$.components.schemas.ContainerProbeOutSerializerV2.properties.timeout_seconds"
	TimeoutSeconds int64 `json:"timeout_seconds,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Exec                resp.Field
		FailureThreshold    resp.Field
		HTTPGet             resp.Field
		InitialDelaySeconds resp.Field
		PeriodSeconds       resp.Field
		SuccessThreshold    resp.Field
		TcpSocket           resp.Field
		TimeoutSeconds      resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerProbe) RawJSON() string { return r.JSON.raw }
func (r *ContainerProbe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InferenceInstanceContainerProbeConfigurationOutSerializerV2'
// "$.components.schemas.InferenceInstanceContainerProbeConfigurationOutSerializerV2"
type ContainerProbeConfig struct {
	// '#/components/schemas/InferenceInstanceContainerProbeConfigurationOutSerializerV2/properties/enabled'
	// "$.components.schemas.InferenceInstanceContainerProbeConfigurationOutSerializerV2.properties.enabled"
	Enabled bool `json:"enabled,required"`
	// '#/components/schemas/InferenceInstanceContainerProbeConfigurationOutSerializerV2/properties/probe/anyOf/0'
	// "$.components.schemas.InferenceInstanceContainerProbeConfigurationOutSerializerV2.properties.probe.anyOf[0]"
	Probe ContainerProbe `json:"probe,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Enabled     resp.Field
		Probe       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerProbeConfig) RawJSON() string { return r.JSON.raw }
func (r *ContainerProbeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InferenceInstanceContainerProbeConfigurationSerializerV2'
// "$.components.schemas.InferenceInstanceContainerProbeConfigurationSerializerV2"
//
// The property Enabled is required.
type ContainerProbeConfigCreateParam struct {
	// '#/components/schemas/InferenceInstanceContainerProbeConfigurationSerializerV2/properties/enabled'
	// "$.components.schemas.InferenceInstanceContainerProbeConfigurationSerializerV2.properties.enabled"
	Enabled bool `json:"enabled,required"`
	// '#/components/schemas/InferenceInstanceContainerProbeConfigurationSerializerV2/properties/probe/anyOf/0'
	// "$.components.schemas.InferenceInstanceContainerProbeConfigurationSerializerV2.properties.probe.anyOf[0]"
	Probe ContainerProbeCreateParam `json:"probe,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ContainerProbeConfigCreateParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ContainerProbeConfigCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ContainerProbeConfigCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerProbeSerializerV2'
// "$.components.schemas.ContainerProbeSerializerV2"
type ContainerProbeCreateParam struct {
	// '#/components/schemas/ContainerProbeSerializerV2/properties/failure_threshold'
	// "$.components.schemas.ContainerProbeSerializerV2.properties.failure_threshold"
	FailureThreshold param.Opt[int64] `json:"failure_threshold,omitzero"`
	// '#/components/schemas/ContainerProbeSerializerV2/properties/initial_delay_seconds'
	// "$.components.schemas.ContainerProbeSerializerV2.properties.initial_delay_seconds"
	InitialDelaySeconds param.Opt[int64] `json:"initial_delay_seconds,omitzero"`
	// '#/components/schemas/ContainerProbeSerializerV2/properties/period_seconds'
	// "$.components.schemas.ContainerProbeSerializerV2.properties.period_seconds"
	PeriodSeconds param.Opt[int64] `json:"period_seconds,omitzero"`
	// '#/components/schemas/ContainerProbeSerializerV2/properties/success_threshold'
	// "$.components.schemas.ContainerProbeSerializerV2.properties.success_threshold"
	SuccessThreshold param.Opt[int64] `json:"success_threshold,omitzero"`
	// '#/components/schemas/ContainerProbeSerializerV2/properties/timeout_seconds'
	// "$.components.schemas.ContainerProbeSerializerV2.properties.timeout_seconds"
	TimeoutSeconds param.Opt[int64] `json:"timeout_seconds,omitzero"`
	// '#/components/schemas/ContainerProbeSerializerV2/properties/exec/anyOf/0'
	// "$.components.schemas.ContainerProbeSerializerV2.properties.exec.anyOf[0]"
	Exec ContainerProbeExecCreateParam `json:"exec,omitzero"`
	// '#/components/schemas/ContainerProbeSerializerV2/properties/http_get/anyOf/0'
	// "$.components.schemas.ContainerProbeSerializerV2.properties.http_get.anyOf[0]"
	HTTPGet ContainerProbeHTTPGetCreateParam `json:"http_get,omitzero"`
	// '#/components/schemas/ContainerProbeSerializerV2/properties/tcp_socket/anyOf/0'
	// "$.components.schemas.ContainerProbeSerializerV2.properties.tcp_socket.anyOf[0]"
	TcpSocket ContainerProbeTcpSocketCreateParam `json:"tcp_socket,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ContainerProbeCreateParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ContainerProbeCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ContainerProbeCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerProbeExecConfigOutSerializerV2'
// "$.components.schemas.ContainerProbeExecConfigOutSerializerV2"
type ContainerProbeExec struct {
	// '#/components/schemas/ContainerProbeExecConfigOutSerializerV2/properties/command'
	// "$.components.schemas.ContainerProbeExecConfigOutSerializerV2.properties.command"
	Command []string `json:"command,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Command     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerProbeExec) RawJSON() string { return r.JSON.raw }
func (r *ContainerProbeExec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ContainerProbeExecConfigSerializerV2'
// "$.components.schemas.ContainerProbeExecConfigSerializerV2"
//
// The property Command is required.
type ContainerProbeExecCreateParam struct {
	// '#/components/schemas/ContainerProbeExecConfigSerializerV2/properties/command'
	// "$.components.schemas.ContainerProbeExecConfigSerializerV2.properties.command"
	Command []string `json:"command,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ContainerProbeExecCreateParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ContainerProbeExecCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ContainerProbeExecCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerProbeHttpGetConfigOutSerializerV2'
// "$.components.schemas.ContainerProbeHttpGetConfigOutSerializerV2"
type ContainerProbeHTTPGet struct {
	// '#/components/schemas/ContainerProbeHttpGetConfigOutSerializerV2/properties/headers'
	// "$.components.schemas.ContainerProbeHttpGetConfigOutSerializerV2.properties.headers"
	Headers map[string]string `json:"headers,required"`
	// '#/components/schemas/ContainerProbeHttpGetConfigOutSerializerV2/properties/host/anyOf/0'
	// "$.components.schemas.ContainerProbeHttpGetConfigOutSerializerV2.properties.host.anyOf[0]"
	Host string `json:"host,required"`
	// '#/components/schemas/ContainerProbeHttpGetConfigOutSerializerV2/properties/path'
	// "$.components.schemas.ContainerProbeHttpGetConfigOutSerializerV2.properties.path"
	Path string `json:"path,required"`
	// '#/components/schemas/ContainerProbeHttpGetConfigOutSerializerV2/properties/port'
	// "$.components.schemas.ContainerProbeHttpGetConfigOutSerializerV2.properties.port"
	Port int64 `json:"port,required"`
	// '#/components/schemas/ContainerProbeHttpGetConfigOutSerializerV2/properties/schema'
	// "$.components.schemas.ContainerProbeHttpGetConfigOutSerializerV2.properties.schema"
	Schema string `json:"schema,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Headers     resp.Field
		Host        resp.Field
		Path        resp.Field
		Port        resp.Field
		Schema      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerProbeHTTPGet) RawJSON() string { return r.JSON.raw }
func (r *ContainerProbeHTTPGet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ContainerProbeHttpGetConfigSerializerV2'
// "$.components.schemas.ContainerProbeHttpGetConfigSerializerV2"
//
// The property Port is required.
type ContainerProbeHTTPGetCreateParam struct {
	// '#/components/schemas/ContainerProbeHttpGetConfigSerializerV2/properties/port'
	// "$.components.schemas.ContainerProbeHttpGetConfigSerializerV2.properties.port"
	Port int64 `json:"port,required"`
	// '#/components/schemas/ContainerProbeHttpGetConfigSerializerV2/properties/host/anyOf/0'
	// "$.components.schemas.ContainerProbeHttpGetConfigSerializerV2.properties.host.anyOf[0]"
	Host param.Opt[string] `json:"host,omitzero"`
	// '#/components/schemas/ContainerProbeHttpGetConfigSerializerV2/properties/path'
	// "$.components.schemas.ContainerProbeHttpGetConfigSerializerV2.properties.path"
	Path param.Opt[string] `json:"path,omitzero"`
	// '#/components/schemas/ContainerProbeHttpGetConfigSerializerV2/properties/schema'
	// "$.components.schemas.ContainerProbeHttpGetConfigSerializerV2.properties.schema"
	Schema param.Opt[string] `json:"schema,omitzero"`
	// '#/components/schemas/ContainerProbeHttpGetConfigSerializerV2/properties/headers'
	// "$.components.schemas.ContainerProbeHttpGetConfigSerializerV2.properties.headers"
	Headers map[string]string `json:"headers,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ContainerProbeHTTPGetCreateParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ContainerProbeHTTPGetCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ContainerProbeHTTPGetCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerProbeTcpSocketConfigOutSerializerV2'
// "$.components.schemas.ContainerProbeTcpSocketConfigOutSerializerV2"
type ContainerProbeTcpSocket struct {
	// '#/components/schemas/ContainerProbeTcpSocketConfigOutSerializerV2/properties/port'
	// "$.components.schemas.ContainerProbeTcpSocketConfigOutSerializerV2.properties.port"
	Port int64 `json:"port,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Port        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerProbeTcpSocket) RawJSON() string { return r.JSON.raw }
func (r *ContainerProbeTcpSocket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ContainerProbeTcpSocketConfigSerializerV2'
// "$.components.schemas.ContainerProbeTcpSocketConfigSerializerV2"
//
// The property Port is required.
type ContainerProbeTcpSocketCreateParam struct {
	// '#/components/schemas/ContainerProbeTcpSocketConfigSerializerV2/properties/port'
	// "$.components.schemas.ContainerProbeTcpSocketConfigSerializerV2.properties.port"
	Port int64 `json:"port,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ContainerProbeTcpSocketCreateParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ContainerProbeTcpSocketCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ContainerProbeTcpSocketCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/ContainerScaleOutSerializerV3'
// "$.components.schemas.ContainerScaleOutSerializerV3"
type ContainerScale struct {
	// '#/components/schemas/ContainerScaleOutSerializerV3/properties/cooldown_period/anyOf/0'
	// "$.components.schemas.ContainerScaleOutSerializerV3.properties.cooldown_period.anyOf[0]"
	CooldownPeriod int64 `json:"cooldown_period,required"`
	// '#/components/schemas/ContainerScaleOutSerializerV3/properties/max'
	// "$.components.schemas.ContainerScaleOutSerializerV3.properties.max"
	Max int64 `json:"max,required"`
	// '#/components/schemas/ContainerScaleOutSerializerV3/properties/min'
	// "$.components.schemas.ContainerScaleOutSerializerV3.properties.min"
	Min int64 `json:"min,required"`
	// '#/components/schemas/ContainerScaleOutSerializerV3/properties/polling_interval/anyOf/0'
	// "$.components.schemas.ContainerScaleOutSerializerV3.properties.polling_interval.anyOf[0]"
	PollingInterval int64 `json:"polling_interval,required"`
	// '#/components/schemas/ContainerScaleOutSerializerV3/properties/triggers'
	// "$.components.schemas.ContainerScaleOutSerializerV3.properties.triggers"
	Triggers ContainerScaleTriggers `json:"triggers,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CooldownPeriod  resp.Field
		Max             resp.Field
		Min             resp.Field
		PollingInterval resp.Field
		Triggers        resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerScale) RawJSON() string { return r.JSON.raw }
func (r *ContainerScale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ContainerScaleTriggersRateOutSerializer'
// "$.components.schemas.ContainerScaleTriggersRateOutSerializer"
type ContainerScaleTriggerRate struct {
	// '#/components/schemas/ContainerScaleTriggersRateOutSerializer/properties/rate'
	// "$.components.schemas.ContainerScaleTriggersRateOutSerializer.properties.rate"
	Rate int64 `json:"rate,required"`
	// '#/components/schemas/ContainerScaleTriggersRateOutSerializer/properties/window'
	// "$.components.schemas.ContainerScaleTriggersRateOutSerializer.properties.window"
	Window int64 `json:"window,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Rate        resp.Field
		Window      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerScaleTriggerRate) RawJSON() string { return r.JSON.raw }
func (r *ContainerScaleTriggerRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ContainerScaleTriggersSqsOutSerializer'
// "$.components.schemas.ContainerScaleTriggersSqsOutSerializer"
type ContainerScaleTriggerSqs struct {
	// '#/components/schemas/ContainerScaleTriggersSqsOutSerializer/properties/activation_queue_length'
	// "$.components.schemas.ContainerScaleTriggersSqsOutSerializer.properties.activation_queue_length"
	ActivationQueueLength int64 `json:"activation_queue_length,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsOutSerializer/properties/aws_endpoint/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersSqsOutSerializer.properties.aws_endpoint.anyOf[0]"
	AwsEndpoint string `json:"aws_endpoint,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsOutSerializer/properties/aws_region'
	// "$.components.schemas.ContainerScaleTriggersSqsOutSerializer.properties.aws_region"
	AwsRegion string `json:"aws_region,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsOutSerializer/properties/queue_length'
	// "$.components.schemas.ContainerScaleTriggersSqsOutSerializer.properties.queue_length"
	QueueLength int64 `json:"queue_length,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsOutSerializer/properties/queue_url'
	// "$.components.schemas.ContainerScaleTriggersSqsOutSerializer.properties.queue_url"
	QueueURL string `json:"queue_url,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsOutSerializer/properties/scale_on_delayed'
	// "$.components.schemas.ContainerScaleTriggersSqsOutSerializer.properties.scale_on_delayed"
	ScaleOnDelayed bool `json:"scale_on_delayed,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsOutSerializer/properties/scale_on_flight'
	// "$.components.schemas.ContainerScaleTriggersSqsOutSerializer.properties.scale_on_flight"
	ScaleOnFlight bool `json:"scale_on_flight,required"`
	// '#/components/schemas/ContainerScaleTriggersSqsOutSerializer/properties/secret_name'
	// "$.components.schemas.ContainerScaleTriggersSqsOutSerializer.properties.secret_name"
	SecretName string `json:"secret_name,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ActivationQueueLength resp.Field
		AwsEndpoint           resp.Field
		AwsRegion             resp.Field
		QueueLength           resp.Field
		QueueURL              resp.Field
		ScaleOnDelayed        resp.Field
		ScaleOnFlight         resp.Field
		SecretName            resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerScaleTriggerSqs) RawJSON() string { return r.JSON.raw }
func (r *ContainerScaleTriggerSqs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ContainerScaleTriggersThresholdOutSerializer'
// "$.components.schemas.ContainerScaleTriggersThresholdOutSerializer"
type ContainerScaleTriggerThreshold struct {
	// '#/components/schemas/ContainerScaleTriggersThresholdOutSerializer/properties/threshold'
	// "$.components.schemas.ContainerScaleTriggersThresholdOutSerializer.properties.threshold"
	Threshold int64 `json:"threshold,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Threshold   resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerScaleTriggerThreshold) RawJSON() string { return r.JSON.raw }
func (r *ContainerScaleTriggerThreshold) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/ContainerScaleTriggersOutSerializer'
// "$.components.schemas.ContainerScaleTriggersOutSerializer"
type ContainerScaleTriggers struct {
	// '#/components/schemas/ContainerScaleTriggersOutSerializer/properties/cpu/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersOutSerializer.properties.cpu.anyOf[0]"
	CPU ContainerScaleTriggerThreshold `json:"cpu,required"`
	// '#/components/schemas/ContainerScaleTriggersOutSerializer/properties/gpu_memory/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersOutSerializer.properties.gpu_memory.anyOf[0]"
	GPUMemory ContainerScaleTriggerThreshold `json:"gpu_memory,required"`
	// '#/components/schemas/ContainerScaleTriggersOutSerializer/properties/gpu_utilization/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersOutSerializer.properties.gpu_utilization.anyOf[0]"
	GPUUtilization ContainerScaleTriggerThreshold `json:"gpu_utilization,required"`
	// '#/components/schemas/ContainerScaleTriggersOutSerializer/properties/http/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersOutSerializer.properties.http.anyOf[0]"
	HTTP ContainerScaleTriggerRate `json:"http,required"`
	// '#/components/schemas/ContainerScaleTriggersOutSerializer/properties/memory/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersOutSerializer.properties.memory.anyOf[0]"
	Memory ContainerScaleTriggerThreshold `json:"memory,required"`
	// '#/components/schemas/ContainerScaleTriggersOutSerializer/properties/sqs/anyOf/0'
	// "$.components.schemas.ContainerScaleTriggersOutSerializer.properties.sqs.anyOf[0]"
	Sqs ContainerScaleTriggerSqs `json:"sqs,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CPU            resp.Field
		GPUMemory      resp.Field
		GPUUtilization resp.Field
		HTTP           resp.Field
		Memory         resp.Field
		Sqs            resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContainerScaleTriggers) RawJSON() string { return r.JSON.raw }
func (r *ContainerScaleTriggers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/DeployStatusSerializer'
// "$.components.schemas.DeployStatusSerializer"
type DeployStatus struct {
	// '#/components/schemas/DeployStatusSerializer/properties/ready'
	// "$.components.schemas.DeployStatusSerializer.properties.ready"
	Ready int64 `json:"ready,required"`
	// '#/components/schemas/DeployStatusSerializer/properties/total'
	// "$.components.schemas.DeployStatusSerializer.properties.total"
	Total int64 `json:"total,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Ready       resp.Field
		Total       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeployStatus) RawJSON() string { return r.JSON.raw }
func (r *DeployStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/InferenceInstanceProbesOutSerializerV2'
// "$.components.schemas.InferenceInstanceProbesOutSerializerV2"
type InferenceProbes struct {
	// '#/components/schemas/InferenceInstanceProbesOutSerializerV2/properties/liveness_probe/anyOf/0'
	// "$.components.schemas.InferenceInstanceProbesOutSerializerV2.properties.liveness_probe.anyOf[0]"
	LivenessProbe ContainerProbeConfig `json:"liveness_probe,required"`
	// '#/components/schemas/InferenceInstanceProbesOutSerializerV2/properties/readiness_probe/anyOf/0'
	// "$.components.schemas.InferenceInstanceProbesOutSerializerV2.properties.readiness_probe.anyOf[0]"
	ReadinessProbe ContainerProbeConfig `json:"readiness_probe,required"`
	// '#/components/schemas/InferenceInstanceProbesOutSerializerV2/properties/startup_probe/anyOf/0'
	// "$.components.schemas.InferenceInstanceProbesOutSerializerV2.properties.startup_probe.anyOf[0]"
	StartupProbe ContainerProbeConfig `json:"startup_probe,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		LivenessProbe  resp.Field
		ReadinessProbe resp.Field
		StartupProbe   resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceProbes) RawJSON() string { return r.JSON.raw }
func (r *InferenceProbes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/IngressOptsSerializer'
// "$.components.schemas.IngressOptsSerializer"
type IngressOptsParam struct {
	// '#/components/schemas/IngressOptsSerializer/properties/disable_response_buffering'
	// "$.components.schemas.IngressOptsSerializer.properties.disable_response_buffering"
	DisableResponseBuffering param.Opt[bool] `json:"disable_response_buffering,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f IngressOptsParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r IngressOptsParam) MarshalJSON() (data []byte, err error) {
	type shadow IngressOptsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/IngressOptsOutSerializer'
// "$.components.schemas.IngressOptsOutSerializer"
type IngressOptsOut struct {
	// '#/components/schemas/IngressOptsOutSerializer/properties/disable_response_buffering'
	// "$.components.schemas.IngressOptsOutSerializer.properties.disable_response_buffering"
	DisableResponseBuffering bool `json:"disable_response_buffering,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		DisableResponseBuffering resp.Field
		ExtraFields              map[string]resp.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IngressOptsOut) RawJSON() string { return r.JSON.raw }
func (r *IngressOptsOut) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/RegionCapacityOutSerializerV3'
// "$.components.schemas.RegionCapacityOutSerializerV3"
type RegionCapacity struct {
	// '#/components/schemas/RegionCapacityOutSerializerV3/properties/capacity'
	// "$.components.schemas.RegionCapacityOutSerializerV3.properties.capacity"
	Capacity []Capacity `json:"capacity,required"`
	// '#/components/schemas/RegionCapacityOutSerializerV3/properties/region_id'
	// "$.components.schemas.RegionCapacityOutSerializerV3.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Capacity    resp.Field
		RegionID    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegionCapacity) RawJSON() string { return r.JSON.raw }
func (r *RegionCapacity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/RegionCapacityOutSerializerV3List'
// "$.components.schemas.RegionCapacityOutSerializerV3List"
type RegionCapacityList struct {
	// '#/components/schemas/RegionCapacityOutSerializerV3List/properties/count'
	// "$.components.schemas.RegionCapacityOutSerializerV3List.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/RegionCapacityOutSerializerV3List/properties/results'
	// "$.components.schemas.RegionCapacityOutSerializerV3List.properties.results"
	Results []RegionCapacity `json:"results,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegionCapacityList) RawJSON() string { return r.JSON.raw }
func (r *RegionCapacityList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
