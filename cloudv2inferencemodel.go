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

// CloudV2InferenceModelService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2InferenceModelService] method instead.
type CloudV2InferenceModelService struct {
	Options []option.RequestOption
}

// NewCloudV2InferenceModelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV2InferenceModelService(opts ...option.RequestOption) (r *CloudV2InferenceModelService) {
	r = &CloudV2InferenceModelService{}
	r.Options = opts
	return
}

// Get ML Model Catalog Details
func (r *CloudV2InferenceModelService) Get(ctx context.Context, modelID string, opts ...option.RequestOption) (res *ModelCatalog, err error) {
	opts = append(r.Options[:], opts...)
	if modelID == "" {
		err = errors.New("missing required model_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/models/%s", modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List ML Model Catalog
func (r *CloudV2InferenceModelService) List(ctx context.Context, query CloudV2InferenceModelListParams, opts ...option.RequestOption) (res *ModelCatalog, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/inference/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ModelCatalog struct {
	// Model ID.
	ID string `json:"id,required" format:"uuid"`
	// Description of the model.
	Description string `json:"description,required"`
	// Image URL of the model.
	ImageURL string `json:"image_url,required"`
	// Name of the model.
	Name string `json:"name,required"`
	// Port on which the model runs.
	Port int64 `json:"port,required"`
	// Category of the model.
	Category string `json:"category,nullable"`
	// Default flavor for the model.
	DefaultFlavorID string `json:"default_flavor_id,nullable" format:"uuid"`
	// Developer of the model.
	Developer string `json:"developer,nullable"`
	// Path to the documentation page.
	DocumentationPage string `json:"documentation_page,nullable"`
	// URL to the EULA text.
	EulaURL string `json:"eula_url,nullable"`
	// Example curl request to the model.
	ExampleCurlRequest string `json:"example_curl_request,nullable"`
	// Whether the model has an EULA.
	HasEula bool `json:"has_eula"`
	// Image registry of the model.
	ImageRegistryID string `json:"image_registry_id,nullable" format:"uuid"`
	// Describing underlying inference engine.
	InferenceBackend string `json:"inference_backend,nullable"`
	// Describing model frontend type.
	InferenceFrontend string `json:"inference_frontend,nullable"`
	// Model name to perform inference call.
	ModelID string `json:"model_id,nullable"`
	// OpenAI compatibility level.
	OpenAICompatibility string `json:"openai_compatibility,nullable"`
	// Version of the model.
	Version string           `json:"version,nullable"`
	JSON    modelCatalogJSON `json:"-"`
}

// modelCatalogJSON contains the JSON metadata for the struct [ModelCatalog]
type modelCatalogJSON struct {
	ID                  apijson.Field
	Description         apijson.Field
	ImageURL            apijson.Field
	Name                apijson.Field
	Port                apijson.Field
	Category            apijson.Field
	DefaultFlavorID     apijson.Field
	Developer           apijson.Field
	DocumentationPage   apijson.Field
	EulaURL             apijson.Field
	ExampleCurlRequest  apijson.Field
	HasEula             apijson.Field
	ImageRegistryID     apijson.Field
	InferenceBackend    apijson.Field
	InferenceFrontend   apijson.Field
	ModelID             apijson.Field
	OpenAICompatibility apijson.Field
	Version             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ModelCatalog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelCatalogJSON) RawJSON() string {
	return r.raw
}

type CloudV2InferenceModelListParams struct {
	// Limit the number of returned instances. Limited by max limit value of 1000
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
	// Order instances by transmitted fields and directions (name.asc)
	OrderBy param.Field[string] `query:"order_by"`
}

// URLQuery serializes [CloudV2InferenceModelListParams]'s query parameters as
// `url.Values`.
func (r CloudV2InferenceModelListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
