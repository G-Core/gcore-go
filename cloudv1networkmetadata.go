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

// CloudV1NetworkMetadataService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1NetworkMetadataService] method instead.
type CloudV1NetworkMetadataService struct {
	Options []option.RequestOption
}

// NewCloudV1NetworkMetadataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1NetworkMetadataService(opts ...option.RequestOption) (r *CloudV1NetworkMetadataService) {
	r = &CloudV1NetworkMetadataService{}
	r.Options = opts
	return
}

// Create or update one or more metadata items for a network. If an item does not
// exist, it gets created. If an item already exists, it's value is overwritten
func (r *CloudV1NetworkMetadataService) New(ctx context.Context, projectID int64, regionID int64, networkID string, body CloudV1NetworkMetadataNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s/metadata", projectID, regionID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Lists all metadata for a network
func (r *CloudV1NetworkMetadataService) List(ctx context.Context, projectID int64, regionID int64, networkID string, opts ...option.RequestOption) (res *ListMetadata, err error) {
	opts = append(r.Options[:], opts...)
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s/metadata", projectID, regionID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// All existing metadata is deleted and replaced with the metadata from the
// request.
func (r *CloudV1NetworkMetadataService) Replace(ctx context.Context, projectID int64, regionID int64, networkID string, body CloudV1NetworkMetadataReplaceParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s/metadata", projectID, regionID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type CloudV1NetworkMetadataNewParams struct {
	// Metadata schema
	Metadata MetadataParam `json:"metadata"`
}

func (r CloudV1NetworkMetadataNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Metadata)
}

type CloudV1NetworkMetadataReplaceParams struct {
	// Metadata schema
	Metadata MetadataParam `json:"metadata"`
}

func (r CloudV1NetworkMetadataReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Metadata)
}
