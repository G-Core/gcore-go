// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
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

type MlcatalogModelCard struct {
	// Model ID.
	ID string `json:"id,required" format:"uuid"`
	// Category of the model.
	Category string `json:"category,required"`
	// Default flavor for the model.
	DefaultFlavorName string `json:"default_flavor_name,required"`
	// Description of the model.
	Description string `json:"description,required"`
	// Developer of the model.
	Developer string `json:"developer,required"`
	// Path to the documentation page.
	DocumentationPage string `json:"documentation_page,required"`
	// URL to the EULA text.
	EulaURL string `json:"eula_url,required"`
	// Example curl request to the model.
	ExampleCurlRequest string `json:"example_curl_request,required"`
	// Whether the model has an EULA.
	HasEula bool `json:"has_eula,required"`
	// Image registry of the model.
	ImageRegistryID string `json:"image_registry_id,required"`
	// Image URL of the model.
	ImageURL string `json:"image_url,required"`
	// Describing underlying inference engine.
	InferenceBackend string `json:"inference_backend,required"`
	// Describing model frontend type.
	InferenceFrontend string `json:"inference_frontend,required"`
	// Model name to perform inference call.
	ModelID string `json:"model_id,required"`
	// Name of the model.
	Name string `json:"name,required"`
	// OpenAI compatibility level.
	OpenAICompatibility string `json:"openai_compatibility,required"`
	// Port on which the model runs.
	Port int64 `json:"port,required"`
	// Version of the model.
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Category            respjson.Field
		DefaultFlavorName   respjson.Field
		Description         respjson.Field
		Developer           respjson.Field
		DocumentationPage   respjson.Field
		EulaURL             respjson.Field
		ExampleCurlRequest  respjson.Field
		HasEula             respjson.Field
		ImageRegistryID     respjson.Field
		ImageURL            respjson.Field
		InferenceBackend    respjson.Field
		InferenceFrontend   respjson.Field
		ModelID             respjson.Field
		Name                respjson.Field
		OpenAICompatibility respjson.Field
		Port                respjson.Field
		Version             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MlcatalogModelCard) RawJSON() string { return r.JSON.raw }
func (r *MlcatalogModelCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MlcatalogOrderByChoices string

const (
	MlcatalogOrderByChoicesNameAsc  MlcatalogOrderByChoices = "name.asc"
	MlcatalogOrderByChoicesNameDesc MlcatalogOrderByChoices = "name.desc"
)

type InferenceModelListParams struct {
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Order instances by transmitted fields and directions
	//
	// Any of "name.asc", "name.desc".
	OrderBy MlcatalogOrderByChoices `query:"order_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InferenceModelListParams]'s query parameters as
// `url.Values`.
func (r InferenceModelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
