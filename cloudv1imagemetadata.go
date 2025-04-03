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

// CloudV1ImageMetadataService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1ImageMetadataService] method instead.
type CloudV1ImageMetadataService struct {
	Options []option.RequestOption
}

// NewCloudV1ImageMetadataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1ImageMetadataService(opts ...option.RequestOption) (r *CloudV1ImageMetadataService) {
	r = &CloudV1ImageMetadataService{}
	r.Options = opts
	return
}

// Create or update one or more metadata items for an image. If an item does not
// exist, it gets created. If an item already exists, it's value is overwritten
func (r *CloudV1ImageMetadataService) New(ctx context.Context, projectID int64, regionID int64, imageID string, body CloudV1ImageMetadataNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v/%s/metadata", projectID, regionID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// List all metadata for an image
func (r *CloudV1ImageMetadataService) List(ctx context.Context, projectID int64, regionID int64, imageID string, opts ...option.RequestOption) (res *ListMetadata, err error) {
	opts = append(r.Options[:], opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v/%s/metadata", projectID, regionID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// All existing metadata is deleted and replaced with the metadata from the
// request.
func (r *CloudV1ImageMetadataService) Replace(ctx context.Context, projectID int64, regionID int64, imageID string, body CloudV1ImageMetadataReplaceParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v/%s/metadata", projectID, regionID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type CloudV1ImageMetadataNewParams struct {
	// Metadata schema
	Metadata MetadataParam `json:"metadata"`
}

func (r CloudV1ImageMetadataNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Metadata)
}

type CloudV1ImageMetadataReplaceParams struct {
	// Metadata schema
	Metadata MetadataParam `json:"metadata"`
}

func (r CloudV1ImageMetadataReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Metadata)
}
