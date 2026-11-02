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

// InferenceApplicationTemplateService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceApplicationTemplateService] method instead.
type InferenceApplicationTemplateService struct {
	Options []option.RequestOption
}

// NewInferenceApplicationTemplateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewInferenceApplicationTemplateService(opts ...option.RequestOption) (r InferenceApplicationTemplateService) {
	r = InferenceApplicationTemplateService{}
	r.Options = opts
	return
}

// Returns a list of available machine learning application templates from the
// catalog. Each template includes metadata such as name, description, cover image,
// documentation, tags, and a set of configurable components (e.g., `model`, `ui`).
// Components define parameters, supported deployment flavors, and other attributes
// required to create a fully functional application deployment.
func (r *InferenceApplicationTemplateService) List(ctx context.Context, query InferenceApplicationTemplateListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[InferenceApplicationTemplate], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cloud/v3/inference/applications/catalog"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Returns a list of available machine learning application templates from the
// catalog. Each template includes metadata such as name, description, cover image,
// documentation, tags, and a set of configurable components (e.g., `model`, `ui`).
// Components define parameters, supported deployment flavors, and other attributes
// required to create a fully functional application deployment.
func (r *InferenceApplicationTemplateService) ListAutoPaging(ctx context.Context, query InferenceApplicationTemplateListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[InferenceApplicationTemplate] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Retrieves detailed information about a specific machine learning application
// template from the catalog. The response includes the application's metadata,
// documentation, tags, and a complete set of components with configuration
// options, compatible flavors, and deployment capabilities — all necessary for
// building and customizing an AI application.
func (r *InferenceApplicationTemplateService) Get(ctx context.Context, applicationName string, opts ...option.RequestOption) (res *InferenceApplicationTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	if applicationName == "" {
		err = errors.New("missing required application_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("cloud/v3/inference/applications/catalog/%s", applicationName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type InferenceApplicationTemplate struct {
	// Configurable components of the application
	Components map[string]InferenceApplicationTemplateComponent `json:"components" api:"required"`
	// URL to the application's cover image
	CoverURL string `json:"cover_url" api:"required"`
	// Brief overview of the application
	Description string `json:"description" api:"required"`
	// Human-readable name of the application
	DisplayName string `json:"display_name" api:"required"`
	// Unique application identifier in the catalog
	Name string `json:"name" api:"required"`
	// Detailed documentation or instructions
	Readme string `json:"readme" api:"required"`
	// Categorization key-value pairs
	Tags map[string]string `json:"tags" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Components  respjson.Field
		CoverURL    respjson.Field
		Description respjson.Field
		DisplayName respjson.Field
		Name        respjson.Field
		Readme      respjson.Field
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationTemplate) RawJSON() string { return r.JSON.raw }
func (r *InferenceApplicationTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationTemplateComponent struct {
	// Summary of the component's functionality
	Description string `json:"description" api:"required"`
	// Human-readable name of the component
	DisplayName string `json:"display_name" api:"required"`
	// Indicates whether this component can expose a public-facing endpoint (e.g., for
	// inference or API access).
	Exposable bool `json:"exposable" api:"required"`
	// URL to the component's license information
	LicenseURL string `json:"license_url" api:"required"`
	// Configurable parameters for the component
	Parameters map[string]InferenceApplicationTemplateComponentParameter `json:"parameters" api:"required"`
	// Detailed documentation or usage instructions
	Readme string `json:"readme" api:"required"`
	// Specifies if the component is required for the application
	Required bool `json:"required" api:"required"`
	// List of compatible flavors or configurations
	SuitableFlavors []InferenceApplicationTemplateComponentSuitableFlavor `json:"suitable_flavors" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description     respjson.Field
		DisplayName     respjson.Field
		Exposable       respjson.Field
		LicenseURL      respjson.Field
		Parameters      respjson.Field
		Readme          respjson.Field
		Required        respjson.Field
		SuitableFlavors respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationTemplateComponent) RawJSON() string { return r.JSON.raw }
func (r *InferenceApplicationTemplateComponent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationTemplateComponentParameter struct {
	// Default value assigned if not provided
	DefaultValue string `json:"default_value" api:"required"`
	// Description of the parameter's purpose
	Description string `json:"description" api:"required"`
	// User-friendly name of the parameter
	DisplayName string `json:"display_name" api:"required"`
	// Allowed values when type is "enum"
	EnumValues []string `json:"enum_values" api:"required"`
	// Maximum value (applies to integer and float types)
	MaxValue string `json:"max_value" api:"required"`
	// Minimum value (applies to integer and float types)
	MinValue string `json:"min_value" api:"required"`
	// Regexp pattern when type is "string"
	Pattern string `json:"pattern" api:"required"`
	// Indicates is parameter mandatory
	Required bool `json:"required" api:"required"`
	// Determines the type of the parameter
	//
	// Any of "enum", "float", "integer", "string".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultValue respjson.Field
		Description  respjson.Field
		DisplayName  respjson.Field
		EnumValues   respjson.Field
		MaxValue     respjson.Field
		MinValue     respjson.Field
		Pattern      respjson.Field
		Required     respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationTemplateComponentParameter) RawJSON() string { return r.JSON.raw }
func (r *InferenceApplicationTemplateComponentParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationTemplateComponentSuitableFlavor struct {
	// Name of the flavor
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationTemplateComponentSuitableFlavor) RawJSON() string { return r.JSON.raw }
func (r *InferenceApplicationTemplateComponentSuitableFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationTemplateList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []InferenceApplicationTemplate `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceApplicationTemplateList) RawJSON() string { return r.JSON.raw }
func (r *InferenceApplicationTemplateList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceApplicationTemplateListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InferenceApplicationTemplateListParams]'s query parameters
// as `url.Values`.
func (r InferenceApplicationTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
