// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2InferenceDeploymentRegionService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2InferenceDeploymentRegionService] method instead.
type CloudV2InferenceDeploymentRegionService struct {
	Options []option.RequestOption
}

// NewCloudV2InferenceDeploymentRegionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV2InferenceDeploymentRegionService(opts ...option.RequestOption) (r *CloudV2InferenceDeploymentRegionService) {
	r = &CloudV2InferenceDeploymentRegionService{}
	r.Options = opts
	return
}

// Deprecated. Get Inference Instance Logs by Region
func (r *CloudV2InferenceDeploymentRegionService) GetLogs(ctx context.Context, instanceID string, regionID int64, opts ...option.RequestOption) (res *CloudV2InferenceDeploymentRegionGetLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/inference/deployments/%s/regions/%v/logs", instanceID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CloudV2InferenceDeploymentRegionGetLogsResponse struct {
	// Log message.
	Message string `json:"message,required"`
	// Pod name.
	Pod string `json:"pod,required"`
	// Log message timestamp in ISO 8601 format.
	Time time.Time                                           `json:"time,required" format:"date-time"`
	JSON cloudV2InferenceDeploymentRegionGetLogsResponseJSON `json:"-"`
}

// cloudV2InferenceDeploymentRegionGetLogsResponseJSON contains the JSON metadata
// for the struct [CloudV2InferenceDeploymentRegionGetLogsResponse]
type cloudV2InferenceDeploymentRegionGetLogsResponseJSON struct {
	Message     apijson.Field
	Pod         apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV2InferenceDeploymentRegionGetLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV2InferenceDeploymentRegionGetLogsResponseJSON) RawJSON() string {
	return r.raw
}
