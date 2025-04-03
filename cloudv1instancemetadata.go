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

// CloudV1InstanceMetadataService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1InstanceMetadataService] method instead.
type CloudV1InstanceMetadataService struct {
	Options []option.RequestOption
}

// NewCloudV1InstanceMetadataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1InstanceMetadataService(opts ...option.RequestOption) (r *CloudV1InstanceMetadataService) {
	r = &CloudV1InstanceMetadataService{}
	r.Options = opts
	return
}

// Create or update one or more metadata items for an instance. If an item does not
// exist, it gets created in the server metadata. If an item already exists, it's
// value is overwritten
func (r *CloudV1InstanceMetadataService) New(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1InstanceMetadataNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/metadata", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Lists all metadata for an instance
func (r *CloudV1InstanceMetadataService) List(ctx context.Context, projectID int64, regionID int64, instanceID string, opts ...option.RequestOption) (res *ListMetadata, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/metadata", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete instance metadata item by key. This endpoint does not support special
// characters in keys. Use /v2/instances/.../metadata_item instead.
func (r *CloudV1InstanceMetadataService) DeleteItem(ctx context.Context, projectID int64, regionID int64, instanceID string, key string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/metadata/%s", projectID, regionID, instanceID, key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get instance metadata item by key. This endpoint does not support special
// characters in keys. Use /v2/instances/.../metadata_item instead.
func (r *CloudV1InstanceMetadataService) GetItem(ctx context.Context, projectID int64, regionID int64, instanceID string, key string, opts ...option.RequestOption) (res *MetadataItem, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/metadata/%s", projectID, regionID, instanceID, key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// All existing metadata is deleted and replaced with the metadata from the request
func (r *CloudV1InstanceMetadataService) Replace(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV1InstanceMetadataReplaceParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/instances/%v/%v/%s/metadata", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type CloudV1InstanceMetadataNewParams struct {
	// Metadata schema
	Metadata MetadataParam `json:"metadata"`
}

func (r CloudV1InstanceMetadataNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Metadata)
}

type CloudV1InstanceMetadataReplaceParams struct {
	// Metadata schema
	Metadata MetadataParam `json:"metadata"`
}

func (r CloudV1InstanceMetadataReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Metadata)
}
