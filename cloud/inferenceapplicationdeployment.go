// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// InferenceApplicationDeploymentService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceApplicationDeploymentService] method instead.
type InferenceApplicationDeploymentService struct {
	Options []option.RequestOption
}

// NewInferenceApplicationDeploymentService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewInferenceApplicationDeploymentService(opts ...option.RequestOption) (r InferenceApplicationDeploymentService) {
	r = InferenceApplicationDeploymentService{}
	r.Options = opts
	return
}

// Creates a new application deployment based on a selected catalog application.
// Specify the desired deployment name, target regions, and configuration for each
// component. The platform will provision the necessary resources and initialize
// the application accordingly.
func (r *InferenceApplicationDeploymentService) New(ctx context.Context, params InferenceApplicationDeploymentNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v3/inference/applications/%v/deployments", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates an existing application deployment. You can modify the target regions
// and update configurations for individual components. To disable a component, set
// its value to null. Only the provided fields will be updated; all others remain
// unchanged.
func (r *InferenceApplicationDeploymentService) Update(ctx context.Context, deploymentName string, params InferenceApplicationDeploymentUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v3/inference/applications/%v/deployments/%s", params.ProjectID.Value, deploymentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Returns a list of your application deployments, including deployment names,
// associated catalog applications, regions, component configurations, and current
// status. Useful for monitoring and managing all active AI application instances.
func (r *InferenceApplicationDeploymentService) List(ctx context.Context, query InferenceApplicationDeploymentListParams, opts ...option.RequestOption) (res *InferenceApplicationDeploymentList, err error) {
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
	path := fmt.Sprintf("cloud/v3/inference/applications/%v/deployments", query.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an existing application deployment along with all associated resources.
// This action will permanently remove the deployment and **terminate all related
// inference instances** that are part of the application.
func (r *InferenceApplicationDeploymentService) Delete(ctx context.Context, deploymentName string, body InferenceApplicationDeploymentDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v3/inference/applications/%v/deployments/%s", body.ProjectID.Value, deploymentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieves detailed information about a specific application deployment. The
// response includes the catalog application it was created from, deployment name,
// active regions, configuration of each component, and the current status of the
// deployment.
func (r *InferenceApplicationDeploymentService) Get(ctx context.Context, deploymentName string, query InferenceApplicationDeploymentGetParams, opts ...option.RequestOption) (res *InferenceApplicationDeployment, err error) {
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
	path := fmt.Sprintf("cloud/v3/inference/applications/%v/deployments/%s", query.ProjectID.Value, deploymentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type InferenceApplicationDeployment struct {
	// List of API keys for the application
	APIKeys []string `json:"api_keys" api:"required"`
	// Identifier of the application template from the catalog
	ApplicationName string `json:"application_name" api:"required"`
	// Mapping of component names to their configuration (e.g., `"model": {...}`)
	ComponentsConfiguration map[string]InferenceApplicationDeploymentComponentsConfiguration `json:"components_configuration" api:"required"`
	// Unique identifier of the deployment
	Name string `json:"name" api:"required"`
	// Geographical regions where the deployment is active
	Regions []int64 `json:"regions" api:"required"`
	// Current state of the deployment across regions
	Status InferenceApplicationDeploymentStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeys                 respjson.Field
		ApplicationName         respjson.Field
		ComponentsConfiguration respjson.Field
		Name                    respjson.Field
		Regions                 respjson.Field
		Status                  respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationDeployment) RawJSON() string { return r.JSON.raw }
func (r *InferenceApplicationDeployment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationDeploymentComponentsConfiguration struct {
	// Indicates if the component will obtain a public address
	Exposed bool `json:"exposed" api:"required"`
	// Chosen flavor or variant of the component
	Flavor string `json:"flavor" api:"required"`
	// Map of parameter overrides for customization
	ParameterOverrides map[string]InferenceApplicationDeploymentComponentsConfigurationParameterOverride `json:"parameter_overrides" api:"required"`
	// Scaling parameters of the component
	Scale InferenceApplicationDeploymentComponentsConfigurationScale `json:"scale" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exposed            respjson.Field
		Flavor             respjson.Field
		ParameterOverrides respjson.Field
		Scale              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationDeploymentComponentsConfiguration) RawJSON() string { return r.JSON.raw }
func (r *InferenceApplicationDeploymentComponentsConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationDeploymentComponentsConfigurationParameterOverride struct {
	// New value assigned to the overridden parameter
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationDeploymentComponentsConfigurationParameterOverride) RawJSON() string {
	return r.JSON.raw
}
func (r *InferenceApplicationDeploymentComponentsConfigurationParameterOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scaling parameters of the component
type InferenceApplicationDeploymentComponentsConfigurationScale struct {
	// Maximum number of replicas the container can be scaled up to
	Max int64 `json:"max" api:"required"`
	// Minimum number of replicas the component can be scaled down to
	Min int64 `json:"min" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationDeploymentComponentsConfigurationScale) RawJSON() string {
	return r.JSON.raw
}
func (r *InferenceApplicationDeploymentComponentsConfigurationScale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current state of the deployment across regions
type InferenceApplicationDeploymentStatus struct {
	// Map of components and their inferences
	ComponentInferences map[string]InferenceApplicationDeploymentStatusComponentInference `json:"component_inferences" api:"required"`
	// High-level summary of the deployment status across all regions
	//
	// Any of "Active", "Failed", "PartiallyDeployed", "Unknown".
	ConsolidatedStatus string `json:"consolidated_status" api:"required"`
	// Map of component keys to their global access endpoints
	ExposeAddresses map[string]InferenceApplicationDeploymentStatusExposeAddress `json:"expose_addresses" api:"required"`
	// Status details for each deployment region
	Regions []InferenceApplicationDeploymentStatusRegion `json:"regions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComponentInferences respjson.Field
		ConsolidatedStatus  respjson.Field
		ExposeAddresses     respjson.Field
		Regions             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationDeploymentStatus) RawJSON() string { return r.JSON.raw }
func (r *InferenceApplicationDeploymentStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationDeploymentStatusComponentInference struct {
	// Flavor of the inference
	Flavor string `json:"flavor" api:"required"`
	// Name of the inference
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Flavor      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationDeploymentStatusComponentInference) RawJSON() string { return r.JSON.raw }
func (r *InferenceApplicationDeploymentStatusComponentInference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationDeploymentStatusExposeAddress struct {
	// Global access endpoint for the component
	Address string `json:"address" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationDeploymentStatusExposeAddress) RawJSON() string { return r.JSON.raw }
func (r *InferenceApplicationDeploymentStatusExposeAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationDeploymentStatusRegion struct {
	// Mapping of component names to their status in the region
	Components map[string]InferenceApplicationDeploymentStatusRegionComponent `json:"components" api:"required"`
	// Region ID
	RegionID int64 `json:"region_id" api:"required"`
	// Current state of the deployment in a specific region
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Components  respjson.Field
		RegionID    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationDeploymentStatusRegion) RawJSON() string { return r.JSON.raw }
func (r *InferenceApplicationDeploymentStatusRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationDeploymentStatusRegionComponent struct {
	// Error message if the component is in an error state
	Error string `json:"error" api:"required"`
	// Current state of the component in a specific region
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationDeploymentStatusRegionComponent) RawJSON() string { return r.JSON.raw }
func (r *InferenceApplicationDeploymentStatusRegionComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationDeploymentList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []InferenceApplicationDeployment `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationDeploymentList) RawJSON() string { return r.JSON.raw }
func (r *InferenceApplicationDeploymentList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationDeploymentNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// Identifier of the application from the catalog
	ApplicationName string `json:"application_name" api:"required"`
	// Mapping of component names to their configuration (e.g., `"model": {...}`)
	ComponentsConfiguration map[string]InferenceApplicationDeploymentNewParamsComponentsConfiguration `json:"components_configuration,omitzero" api:"required"`
	// Desired name for the new deployment
	Name string `json:"name" api:"required"`
	// Geographical regions where the deployment should be created
	Regions []int64 `json:"regions,omitzero" api:"required"`
	// List of API keys for the application
	APIKeys []string `json:"api_keys,omitzero"`
	paramObj
}

func (r InferenceApplicationDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceApplicationDeploymentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceApplicationDeploymentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Exposed, Flavor, Scale are required.
type InferenceApplicationDeploymentNewParamsComponentsConfiguration struct {
	// Whether the component should be exposed via a public endpoint (e.g., for
	// external inference/API access).
	Exposed bool `json:"exposed" api:"required"`
	// Specifies the compute configuration (e.g., CPU/GPU size) to be used for the
	// component.
	Flavor string `json:"flavor" api:"required"`
	// Scaling parameters of the component
	Scale InferenceApplicationDeploymentNewParamsComponentsConfigurationScale `json:"scale,omitzero" api:"required"`
	// Map of parameter overrides for customization
	ParameterOverrides map[string]InferenceApplicationDeploymentNewParamsComponentsConfigurationParameterOverride `json:"parameter_overrides,omitzero"`
	paramObj
}

func (r InferenceApplicationDeploymentNewParamsComponentsConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow InferenceApplicationDeploymentNewParamsComponentsConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceApplicationDeploymentNewParamsComponentsConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scaling parameters of the component
//
// The properties Max, Min are required.
type InferenceApplicationDeploymentNewParamsComponentsConfigurationScale struct {
	// Maximum number of replicas the container can be scaled up to
	Max int64 `json:"max" api:"required"`
	// Minimum number of replicas the component can be scaled down to
	Min int64 `json:"min" api:"required"`
	paramObj
}

func (r InferenceApplicationDeploymentNewParamsComponentsConfigurationScale) MarshalJSON() (data []byte, err error) {
	type shadow InferenceApplicationDeploymentNewParamsComponentsConfigurationScale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceApplicationDeploymentNewParamsComponentsConfigurationScale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Value is required.
type InferenceApplicationDeploymentNewParamsComponentsConfigurationParameterOverride struct {
	// New value assigned to the overridden parameter
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InferenceApplicationDeploymentNewParamsComponentsConfigurationParameterOverride) MarshalJSON() (data []byte, err error) {
	type shadow InferenceApplicationDeploymentNewParamsComponentsConfigurationParameterOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceApplicationDeploymentNewParamsComponentsConfigurationParameterOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationDeploymentUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	// List of API keys for the application
	APIKeys []string `json:"api_keys,omitzero"`
	// Mapping of component names to their configuration (e.g., `"model": {...}`)
	ComponentsConfiguration map[string]InferenceApplicationDeploymentUpdateParamsComponentsConfiguration `json:"components_configuration,omitzero"`
	// Geographical regions to be updated for the deployment
	Regions []int64 `json:"regions,omitzero"`
	paramObj
}

func (r InferenceApplicationDeploymentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceApplicationDeploymentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceApplicationDeploymentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationDeploymentUpdateParamsComponentsConfiguration struct {
	// Whether the component should be exposed via a public endpoint (e.g., for
	// external inference/API access).
	Exposed param.Opt[bool] `json:"exposed,omitzero"`
	// Specifies the compute configuration (e.g., CPU/GPU size) to be used for the
	// component.
	Flavor param.Opt[string] `json:"flavor,omitzero"`
	// Map of parameter overrides for customization
	ParameterOverrides map[string]InferenceApplicationDeploymentUpdateParamsComponentsConfigurationParameterOverride `json:"parameter_overrides,omitzero"`
	// Scaling parameters of the component
	Scale InferenceApplicationDeploymentUpdateParamsComponentsConfigurationScale `json:"scale,omitzero"`
	paramObj
}

func (r InferenceApplicationDeploymentUpdateParamsComponentsConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow InferenceApplicationDeploymentUpdateParamsComponentsConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceApplicationDeploymentUpdateParamsComponentsConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Value is required.
type InferenceApplicationDeploymentUpdateParamsComponentsConfigurationParameterOverride struct {
	// New value assigned to the overridden parameter
	Value string `json:"value" api:"required"`
	paramObj
}

func (r InferenceApplicationDeploymentUpdateParamsComponentsConfigurationParameterOverride) MarshalJSON() (data []byte, err error) {
	type shadow InferenceApplicationDeploymentUpdateParamsComponentsConfigurationParameterOverride
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceApplicationDeploymentUpdateParamsComponentsConfigurationParameterOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scaling parameters of the component
type InferenceApplicationDeploymentUpdateParamsComponentsConfigurationScale struct {
	// Maximum number of replicas the container can be scaled up to
	Max param.Opt[int64] `json:"max,omitzero"`
	// Minimum number of replicas the component can be scaled down to
	Min param.Opt[int64] `json:"min,omitzero"`
	paramObj
}

func (r InferenceApplicationDeploymentUpdateParamsComponentsConfigurationScale) MarshalJSON() (data []byte, err error) {
	type shadow InferenceApplicationDeploymentUpdateParamsComponentsConfigurationScale
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceApplicationDeploymentUpdateParamsComponentsConfigurationScale) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationDeploymentListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	paramObj
}

type InferenceApplicationDeploymentDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	paramObj
}

type InferenceApplicationDeploymentGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero" api:"required" json:"-"`
	paramObj
}
