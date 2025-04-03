// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV3InferenceModelService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV3InferenceModelService] method instead.
type CloudV3InferenceModelService struct {
	Options []option.RequestOption
}

// NewCloudV3InferenceModelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV3InferenceModelService(opts ...option.RequestOption) (r *CloudV3InferenceModelService) {
	r = &CloudV3InferenceModelService{}
	r.Options = opts
	return
}

// Get model from catalog
func (r *CloudV3InferenceModelService) Get(ctx context.Context, modelID string, opts ...option.RequestOption) (res *MlCatalogModelCard, err error) {
	opts = append(r.Options[:], opts...)
	if modelID == "" {
		err = errors.New("missing required model_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/models/%s", modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List models from catalog
func (r *CloudV3InferenceModelService) List(ctx context.Context, query CloudV3InferenceModelListParams, opts ...option.RequestOption) (res *CloudV3InferenceModelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v3/inference/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type MlCatalogModelCard struct {
	// Model ID.
	ID string `json:"id,required" format:"uuid"`
	// Category of the model.
	Category string `json:"category,required,nullable"`
	// Default flavor for the model.
	DefaultFlavorName string `json:"default_flavor_name,required,nullable"`
	// Description of the model.
	Description string `json:"description,required"`
	// Developer of the model.
	Developer string `json:"developer,required,nullable"`
	// Path to the documentation page.
	DocumentationPage string `json:"documentation_page,required,nullable"`
	// URL to the EULA text.
	EulaURL string `json:"eula_url,required,nullable"`
	// Example curl request to the model.
	ExampleCurlRequest string `json:"example_curl_request,required,nullable"`
	// Whether the model has an EULA.
	HasEula bool `json:"has_eula,required"`
	// Image registry of the model.
	ImageRegistryID string `json:"image_registry_id,required,nullable"`
	// Image URL of the model.
	ImageURL string `json:"image_url,required"`
	// Describing underlying inference engine.
	InferenceBackend string `json:"inference_backend,required,nullable"`
	// Describing model frontend type.
	InferenceFrontend string `json:"inference_frontend,required,nullable"`
	// Model name to perform inference call.
	ModelID string `json:"model_id,required,nullable"`
	// Name of the model.
	Name string `json:"name,required"`
	// OpenAI compatibility level.
	OpenAICompatibility string `json:"openai_compatibility,required,nullable"`
	// Port on which the model runs.
	Port int64 `json:"port,required"`
	// Version of the model.
	Version string                 `json:"version,required,nullable"`
	JSON    mlCatalogModelCardJSON `json:"-"`
}

// mlCatalogModelCardJSON contains the JSON metadata for the struct
// [MlCatalogModelCard]
type mlCatalogModelCardJSON struct {
	ID                  apijson.Field
	Category            apijson.Field
	DefaultFlavorName   apijson.Field
	Description         apijson.Field
	Developer           apijson.Field
	DocumentationPage   apijson.Field
	EulaURL             apijson.Field
	ExampleCurlRequest  apijson.Field
	HasEula             apijson.Field
	ImageRegistryID     apijson.Field
	ImageURL            apijson.Field
	InferenceBackend    apijson.Field
	InferenceFrontend   apijson.Field
	ModelID             apijson.Field
	Name                apijson.Field
	OpenAICompatibility apijson.Field
	Port                apijson.Field
	Version             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MlCatalogModelCard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mlCatalogModelCardJSON) RawJSON() string {
	return r.raw
}

type CloudV3InferenceModelListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []MlCatalogModelCard                  `json:"results,required"`
	JSON    cloudV3InferenceModelListResponseJSON `json:"-"`
}

// cloudV3InferenceModelListResponseJSON contains the JSON metadata for the struct
// [CloudV3InferenceModelListResponse]
type cloudV3InferenceModelListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV3InferenceModelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3InferenceModelListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV3InferenceModelListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Field[int64] `query:"offset"`
	// Order instances by transmitted fields and directions (name.asc)
	OrderBy param.Field[string] `query:"order_by"`
}

// URLQuery serializes [CloudV3InferenceModelListParams]'s query parameters as
// `url.Values`.
func (r CloudV3InferenceModelListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
