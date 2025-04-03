// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV3InferenceService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV3InferenceService] method instead.
type CloudV3InferenceService struct {
	Options             []option.RequestOption
	Flavors             *CloudV3InferenceFlavorService
	Models              *CloudV3InferenceModelService
	Deployments         *CloudV3InferenceDeploymentService
	RegistryCredentials *CloudV3InferenceRegistryCredentialService
	Secrets             *CloudV3InferenceSecretService
}

// NewCloudV3InferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV3InferenceService(opts ...option.RequestOption) (r *CloudV3InferenceService) {
	r = &CloudV3InferenceService{}
	r.Options = opts
	r.Flavors = NewCloudV3InferenceFlavorService(opts...)
	r.Models = NewCloudV3InferenceModelService(opts...)
	r.Deployments = NewCloudV3InferenceDeploymentService(opts...)
	r.RegistryCredentials = NewCloudV3InferenceRegistryCredentialService(opts...)
	r.Secrets = NewCloudV3InferenceSecretService(opts...)
	return
}

// Get inference capacity by region
func (r *CloudV3InferenceService) GetCapacity(ctx context.Context, opts ...option.RequestOption) (res *CloudV3InferenceGetCapacityResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v3/inference/capacity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CloudV3InferenceGetCapacityResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV3InferenceGetCapacityResponseResult `json:"results,required"`
	JSON    cloudV3InferenceGetCapacityResponseJSON     `json:"-"`
}

// cloudV3InferenceGetCapacityResponseJSON contains the JSON metadata for the
// struct [CloudV3InferenceGetCapacityResponse]
type cloudV3InferenceGetCapacityResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV3InferenceGetCapacityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3InferenceGetCapacityResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV3InferenceGetCapacityResponseResult struct {
	// List of capacities by flavor.
	Capacity []CloudV3InferenceGetCapacityResponseResultsCapacity `json:"capacity,required"`
	// Region ID.
	RegionID int64                                         `json:"region_id,required"`
	JSON     cloudV3InferenceGetCapacityResponseResultJSON `json:"-"`
}

// cloudV3InferenceGetCapacityResponseResultJSON contains the JSON metadata for the
// struct [CloudV3InferenceGetCapacityResponseResult]
type cloudV3InferenceGetCapacityResponseResultJSON struct {
	Capacity    apijson.Field
	RegionID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV3InferenceGetCapacityResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3InferenceGetCapacityResponseResultJSON) RawJSON() string {
	return r.raw
}

type CloudV3InferenceGetCapacityResponseResultsCapacity struct {
	// Available capacity.
	Capacity int64 `json:"capacity,required"`
	// Flavor name.
	FlavorName string                                                 `json:"flavor_name,required"`
	JSON       cloudV3InferenceGetCapacityResponseResultsCapacityJSON `json:"-"`
}

// cloudV3InferenceGetCapacityResponseResultsCapacityJSON contains the JSON
// metadata for the struct [CloudV3InferenceGetCapacityResponseResultsCapacity]
type cloudV3InferenceGetCapacityResponseResultsCapacityJSON struct {
	Capacity    apijson.Field
	FlavorName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV3InferenceGetCapacityResponseResultsCapacity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV3InferenceGetCapacityResponseResultsCapacityJSON) RawJSON() string {
	return r.raw
}
