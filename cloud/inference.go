// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
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
	Deployments         InferenceDeploymentService
	RegistryCredentials InferenceRegistryCredentialService
	Secrets             InferenceSecretService
	APIKeys             InferenceAPIKeyService
	Applications        InferenceApplicationService
}

// NewInferenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInferenceService(opts ...option.RequestOption) (r InferenceService) {
	r = InferenceService{}
	r.Options = opts
	r.Flavors = NewInferenceFlavorService(opts...)
	r.Deployments = NewInferenceDeploymentService(opts...)
	r.RegistryCredentials = NewInferenceRegistryCredentialService(opts...)
	r.Secrets = NewInferenceSecretService(opts...)
	r.APIKeys = NewInferenceAPIKeyService(opts...)
	r.Applications = NewInferenceApplicationService(opts...)
	return
}

// Get inference capacity by region
func (r *InferenceService) GetCapacityByRegion(ctx context.Context, opts ...option.RequestOption) (res *InferenceRegionCapacityList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "cloud/v3/inference/capacity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type InferenceRegionCapacity struct {
	// List of capacities by flavor.
	Capacity []InferenceRegionCapacityCapacity `json:"capacity" api:"required"`
	// Region ID.
	RegionID int64 `json:"region_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capacity    respjson.Field
		RegionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceRegionCapacity) RawJSON() string { return r.JSON.raw }
func (r *InferenceRegionCapacity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceRegionCapacityCapacity struct {
	// Available capacity.
	Capacity int64 `json:"capacity" api:"required"`
	// Flavor name.
	FlavorName string `json:"flavor_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capacity    respjson.Field
		FlavorName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceRegionCapacityCapacity) RawJSON() string { return r.JSON.raw }
func (r *InferenceRegionCapacityCapacity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceRegionCapacityList struct {
	// Number of objects
	Count int64 `json:"count" api:"required"`
	// Objects
	Results []InferenceRegionCapacity `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceRegionCapacityList) RawJSON() string { return r.JSON.raw }
func (r *InferenceRegionCapacityList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
