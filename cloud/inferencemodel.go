// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// InferenceModelService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceModelService] method instead.
type InferenceModelService struct {
	Options []option.RequestOption
}

// NewInferenceModelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInferenceModelService(opts ...option.RequestOption) (r InferenceModelService) {
	r = InferenceModelService{}
	r.Options = opts
	return
}

// List models from catalog
func (r *InferenceModelService) List(ctx context.Context, query InferenceModelListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[MlcatalogModelCard], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cloud/v3/inference/models"
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

// List models from catalog
func (r *InferenceModelService) ListAutoPaging(ctx context.Context, query InferenceModelListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[MlcatalogModelCard] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Get model from catalog
func (r *InferenceModelService) Get(ctx context.Context, modelID string, opts ...option.RequestOption) (res *MlcatalogModelCard, err error) {
	opts = append(r.Options[:], opts...)
	if modelID == "" {
		err = errors.New("missing required model_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/models/%s", modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// '#/components/schemas/MLCatalogModelCardSerializerV3'
// "$.components.schemas.MLCatalogModelCardSerializerV3"
type MlcatalogModelCard struct {
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/id'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.id"
	ID string `json:"id,required" format:"uuid"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/category/anyOf/0'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.category.anyOf[0]"
	Category string `json:"category,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/default_flavor_name/anyOf/0'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.default_flavor_name.anyOf[0]"
	DefaultFlavorName string `json:"default_flavor_name,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/description'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.description"
	Description string `json:"description,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/developer/anyOf/0'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.developer.anyOf[0]"
	Developer string `json:"developer,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/documentation_page/anyOf/0'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.documentation_page.anyOf[0]"
	DocumentationPage string `json:"documentation_page,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/eula_url/anyOf/0'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.eula_url.anyOf[0]"
	EulaURL string `json:"eula_url,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/example_curl_request/anyOf/0'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.example_curl_request.anyOf[0]"
	ExampleCurlRequest string `json:"example_curl_request,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/has_eula'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.has_eula"
	HasEula bool `json:"has_eula,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/image_registry_id/anyOf/0'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.image_registry_id.anyOf[0]"
	ImageRegistryID string `json:"image_registry_id,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/image_url'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.image_url"
	ImageURL string `json:"image_url,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/inference_backend/anyOf/0'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.inference_backend.anyOf[0]"
	InferenceBackend string `json:"inference_backend,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/inference_frontend/anyOf/0'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.inference_frontend.anyOf[0]"
	InferenceFrontend string `json:"inference_frontend,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/model_id/anyOf/0'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.model_id.anyOf[0]"
	ModelID string `json:"model_id,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/name'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/openai_compatibility/anyOf/0'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.openai_compatibility.anyOf[0]"
	OpenAICompatibility string `json:"openai_compatibility,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/port'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.port"
	Port int64 `json:"port,required"`
	// '#/components/schemas/MLCatalogModelCardSerializerV3/properties/version/anyOf/0'
	// "$.components.schemas.MLCatalogModelCardSerializerV3.properties.version.anyOf[0]"
	Version string `json:"version,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                  resp.Field
		Category            resp.Field
		DefaultFlavorName   resp.Field
		Description         resp.Field
		Developer           resp.Field
		DocumentationPage   resp.Field
		EulaURL             resp.Field
		ExampleCurlRequest  resp.Field
		HasEula             resp.Field
		ImageRegistryID     resp.Field
		ImageURL            resp.Field
		InferenceBackend    resp.Field
		InferenceFrontend   resp.Field
		ModelID             resp.Field
		Name                resp.Field
		OpenAICompatibility resp.Field
		Port                resp.Field
		Version             resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MlcatalogModelCard) RawJSON() string { return r.JSON.raw }
func (r *MlcatalogModelCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/MLCatalogOrderByChoices'
// "$.components.schemas.MLCatalogOrderByChoices"
type MlcatalogOrderByChoices string

const (
	MlcatalogOrderByChoicesNameAsc  MlcatalogOrderByChoices = "name.asc"
	MlcatalogOrderByChoicesNameDesc MlcatalogOrderByChoices = "name.desc"
)

type InferenceModelListParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2Fmodels/get/parameters/0'
	// "$.paths['/cloud/v3/inference/models'].get.parameters[0]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Finference%2Fmodels/get/parameters/1'
	// "$.paths['/cloud/v3/inference/models'].get.parameters[1]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Finference%2Fmodels/get/parameters/2'
	// "$.paths['/cloud/v3/inference/models'].get.parameters[2]"
	//
	// Any of "name.asc", "name.desc".
	OrderBy MlcatalogOrderByChoices `query:"order_by,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceModelListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [InferenceModelListParams]'s query parameters as
// `url.Values`.
func (r InferenceModelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
