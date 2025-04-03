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

// CloudV2AIClusterMetadataService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2AIClusterMetadataService] method instead.
type CloudV2AIClusterMetadataService struct {
	Options []option.RequestOption
}

// NewCloudV2AIClusterMetadataService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV2AIClusterMetadataService(opts ...option.RequestOption) (r *CloudV2AIClusterMetadataService) {
	r = &CloudV2AIClusterMetadataService{}
	r.Options = opts
	return
}

// Create or update one or more metadata items of the cluster. If an item does not
// exist, it gets created in the cluster metadata. If an item already exists, it's
// value is overwritten
func (r *CloudV2AIClusterMetadataService) NewOrUpdate(ctx context.Context, projectID int64, regionID int64, clusterID string, body CloudV2AIClusterMetadataNewOrUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s/metadata", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// All existing metadata is deleted and replaced with the metadata from the request
func (r *CloudV2AIClusterMetadataService) Replace(ctx context.Context, projectID int64, regionID int64, clusterID string, body CloudV2AIClusterMetadataReplaceParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s/metadata", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type CloudV2AIClusterMetadataNewOrUpdateParams struct {
	// Metadata schema
	Metadata MetadataParam `json:"metadata"`
}

func (r CloudV2AIClusterMetadataNewOrUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Metadata)
}

type CloudV2AIClusterMetadataReplaceParams struct {
	// Metadata schema
	Metadata MetadataParam `json:"metadata"`
}

func (r CloudV2AIClusterMetadataReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Metadata)
}
