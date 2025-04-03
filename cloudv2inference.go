// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2InferenceService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2InferenceService] method instead.
type CloudV2InferenceService struct {
	Options     []option.RequestOption
	APIKeys     *CloudV2InferenceAPIKeyService
	Deployments *CloudV2InferenceDeploymentService
	Flavors     *CloudV2InferenceFlavorService
	Models      *CloudV2InferenceModelService
	Registries  *CloudV2InferenceRegistryService
}

// NewCloudV2InferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV2InferenceService(opts ...option.RequestOption) (r *CloudV2InferenceService) {
	r = &CloudV2InferenceService{}
	r.Options = opts
	r.APIKeys = NewCloudV2InferenceAPIKeyService(opts...)
	r.Deployments = NewCloudV2InferenceDeploymentService(opts...)
	r.Flavors = NewCloudV2InferenceFlavorService(opts...)
	r.Models = NewCloudV2InferenceModelService(opts...)
	r.Registries = NewCloudV2InferenceRegistryService(opts...)
	return
}

// Deprecated. Get Capacity for regions
func (r *CloudV2InferenceService) GetCapacity(ctx context.Context, opts ...option.RequestOption) (res *CloudV2InferenceGetCapacityResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/inference/capacity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CloudV2InferenceGetCapacityResponse struct {
	// List of capacities per flavor.
	Capacity []CloudV2InferenceGetCapacityResponseCapacity `json:"capacity,required"`
	// Region ID.
	RegionID int64                                   `json:"region_id,required"`
	JSON     cloudV2InferenceGetCapacityResponseJSON `json:"-"`
}

// cloudV2InferenceGetCapacityResponseJSON contains the JSON metadata for the
// struct [CloudV2InferenceGetCapacityResponse]
type cloudV2InferenceGetCapacityResponseJSON struct {
	Capacity    apijson.Field
	RegionID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV2InferenceGetCapacityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV2InferenceGetCapacityResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV2InferenceGetCapacityResponseCapacity struct {
	// Available capacity.
	Capacity int64 `json:"capacity,required"`
	// Flavor ID.
	FlavorID string                                          `json:"flavor_id,required" format:"uuid"`
	JSON     cloudV2InferenceGetCapacityResponseCapacityJSON `json:"-"`
}

// cloudV2InferenceGetCapacityResponseCapacityJSON contains the JSON metadata for
// the struct [CloudV2InferenceGetCapacityResponseCapacity]
type cloudV2InferenceGetCapacityResponseCapacityJSON struct {
	Capacity    apijson.Field
	FlavorID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV2InferenceGetCapacityResponseCapacity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV2InferenceGetCapacityResponseCapacityJSON) RawJSON() string {
	return r.raw
}
