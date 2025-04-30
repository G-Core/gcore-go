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

// InferenceFlavorService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceFlavorService] method instead.
type InferenceFlavorService struct {
	Options []option.RequestOption
}

// NewInferenceFlavorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInferenceFlavorService(opts ...option.RequestOption) (r InferenceFlavorService) {
	r = InferenceFlavorService{}
	r.Options = opts
	return
}

// List inference flavors
func (r *InferenceFlavorService) List(ctx context.Context, query InferenceFlavorListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[InferenceFlavor], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cloud/v3/inference/flavors"
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

// List inference flavors
func (r *InferenceFlavorService) ListAutoPaging(ctx context.Context, query InferenceFlavorListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[InferenceFlavor] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Get inference flavor
func (r *InferenceFlavorService) Get(ctx context.Context, flavorName string, opts ...option.RequestOption) (res *InferenceFlavor, err error) {
	opts = append(r.Options[:], opts...)
	if flavorName == "" {
		err = errors.New("missing required flavor_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/inference/flavors/%s", flavorName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// '#/components/schemas/InferenceFlavorOutSerializerV3'
// "$.components.schemas.InferenceFlavorOutSerializerV3"
type InferenceFlavor struct {
	// '#/components/schemas/InferenceFlavorOutSerializerV3/properties/cpu'
	// "$.components.schemas.InferenceFlavorOutSerializerV3.properties.cpu"
	CPU float64 `json:"cpu,required"`
	// '#/components/schemas/InferenceFlavorOutSerializerV3/properties/description'
	// "$.components.schemas.InferenceFlavorOutSerializerV3.properties.description"
	Description string `json:"description,required"`
	// '#/components/schemas/InferenceFlavorOutSerializerV3/properties/gpu'
	// "$.components.schemas.InferenceFlavorOutSerializerV3.properties.gpu"
	GPU int64 `json:"gpu,required"`
	// '#/components/schemas/InferenceFlavorOutSerializerV3/properties/gpu_compute_capability'
	// "$.components.schemas.InferenceFlavorOutSerializerV3.properties.gpu_compute_capability"
	GPUComputeCapability string `json:"gpu_compute_capability,required"`
	// '#/components/schemas/InferenceFlavorOutSerializerV3/properties/gpu_memory'
	// "$.components.schemas.InferenceFlavorOutSerializerV3.properties.gpu_memory"
	GPUMemory float64 `json:"gpu_memory,required"`
	// '#/components/schemas/InferenceFlavorOutSerializerV3/properties/gpu_model'
	// "$.components.schemas.InferenceFlavorOutSerializerV3.properties.gpu_model"
	GPUModel string `json:"gpu_model,required"`
	// '#/components/schemas/InferenceFlavorOutSerializerV3/properties/is_gpu_shared'
	// "$.components.schemas.InferenceFlavorOutSerializerV3.properties.is_gpu_shared"
	IsGPUShared bool `json:"is_gpu_shared,required"`
	// '#/components/schemas/InferenceFlavorOutSerializerV3/properties/memory'
	// "$.components.schemas.InferenceFlavorOutSerializerV3.properties.memory"
	Memory float64 `json:"memory,required"`
	// '#/components/schemas/InferenceFlavorOutSerializerV3/properties/name'
	// "$.components.schemas.InferenceFlavorOutSerializerV3.properties.name"
	Name string `json:"name,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CPU                  resp.Field
		Description          resp.Field
		GPU                  resp.Field
		GPUComputeCapability resp.Field
		GPUMemory            resp.Field
		GPUModel             resp.Field
		IsGPUShared          resp.Field
		Memory               resp.Field
		Name                 resp.Field
		ExtraFields          map[string]resp.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceFlavor) RawJSON() string { return r.JSON.raw }
func (r *InferenceFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceFlavorListParams struct {
	// '#/paths/%2Fcloud%2Fv3%2Finference%2Fflavors/get/parameters/0'
	// "$.paths['/cloud/v3/inference/flavors'].get.parameters[0]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv3%2Finference%2Fflavors/get/parameters/1'
	// "$.paths['/cloud/v3/inference/flavors'].get.parameters[1]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InferenceFlavorListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [InferenceFlavorListParams]'s query parameters as
// `url.Values`.
func (r InferenceFlavorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
