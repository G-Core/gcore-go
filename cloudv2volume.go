// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2VolumeService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2VolumeService] method instead.
type CloudV2VolumeService struct {
	Options []option.RequestOption
}

// NewCloudV2VolumeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV2VolumeService(opts ...option.RequestOption) (r *CloudV2VolumeService) {
	r = &CloudV2VolumeService{}
	r.Options = opts
	return
}

// Attach the volume to instance. Note: ultra volume can only be attached to an
// instance with shared flavor
func (r *CloudV2VolumeService) Attach(ctx context.Context, projectID int64, regionID int64, volumeID string, body CloudV2VolumeAttachParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/volumes/%v/%v/%s/attach", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detach the volume from instance
func (r *CloudV2VolumeService) Detach(ctx context.Context, projectID int64, regionID int64, volumeID string, body CloudV2VolumeDetachParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/volumes/%v/%v/%s/detach", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CloudV2VolumeAttachParams struct {
	// Attach volume to instance schema
	AttachVolumeToInstance AttachVolumeToInstanceParam `json:"attach_volume_to_instance,required"`
}

func (r CloudV2VolumeAttachParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AttachVolumeToInstance)
}

type CloudV2VolumeDetachParams struct {
	InstanceID InstanceIDParam `json:"instance_id,required"`
}

func (r CloudV2VolumeDetachParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.InstanceID)
}
