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

// CloudV1VolumeMetadataService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1VolumeMetadataService] method instead.
type CloudV1VolumeMetadataService struct {
	Options []option.RequestOption
}

// NewCloudV1VolumeMetadataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1VolumeMetadataService(opts ...option.RequestOption) (r *CloudV1VolumeMetadataService) {
	r = &CloudV1VolumeMetadataService{}
	r.Options = opts
	return
}

// Create or update one or more metadata items for a volume. If an item does not
// exist, it gets created. If an item already exists, it's value is overwritten
func (r *CloudV1VolumeMetadataService) New(ctx context.Context, projectID int64, regionID int64, volumeID string, body CloudV1VolumeMetadataNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/metadata", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// List all metadata for a volume
func (r *CloudV1VolumeMetadataService) List(ctx context.Context, projectID int64, regionID int64, volumeID string, opts ...option.RequestOption) (res *ListMetadata, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/metadata", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// All existing metadata is deleted and replaced with the metadata from the
// request.
func (r *CloudV1VolumeMetadataService) Replace(ctx context.Context, projectID int64, regionID int64, volumeID string, body CloudV1VolumeMetadataReplaceParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/metadata", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type CloudV1VolumeMetadataNewParams struct {
	// Metadata schema
	Metadata MetadataParam `json:"metadata"`
}

func (r CloudV1VolumeMetadataNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Metadata)
}

type CloudV1VolumeMetadataReplaceParams struct {
	// Metadata schema
	Metadata MetadataParam `json:"metadata"`
}

func (r CloudV1VolumeMetadataReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Metadata)
}
