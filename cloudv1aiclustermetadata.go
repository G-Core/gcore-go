// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1AIClusterMetadataService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1AIClusterMetadataService] method instead.
type CloudV1AIClusterMetadataService struct {
	Options []option.RequestOption
}

// NewCloudV1AIClusterMetadataService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1AIClusterMetadataService(opts ...option.RequestOption) (r *CloudV1AIClusterMetadataService) {
	r = &CloudV1AIClusterMetadataService{}
	r.Options = opts
	return
}

// Create or update one or more metadata items of the cluster. If an item does not
// exist, it gets created in the cluster metadata. If an item already exists, it's
// value is overwritten
func (r *CloudV1AIClusterMetadataService) Update(ctx context.Context, projectID int64, regionID int64, clusterID string, body CloudV1AIClusterMetadataUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/metadata", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// List all metadata of AI cluster
func (r *CloudV1AIClusterMetadataService) List(ctx context.Context, projectID int64, regionID int64, clusterID string, opts ...option.RequestOption) (res *ListMetadata, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/clusters/%v/%v/%s/metadata", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Metadata schema
type ListMetadata struct {
	// Number of metadata items returned
	Count int64 `json:"count,required"`
	// Metadata list
	Results []MetadataItem   `json:"results"`
	JSON    listMetadataJSON `json:"-"`
}

// listMetadataJSON contains the JSON metadata for the struct [ListMetadata]
type listMetadataJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listMetadataJSON) RawJSON() string {
	return r.raw
}

// Metadata schema
type MetadataParam struct {
	// Metadata key and value pairs. The maximum size for key is 255 bytes, for value
	// is 1024 bytes.
	Key param.Field[string] `json:"key"`
}

func (r MetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1AIClusterMetadataUpdateParams struct {
	// Metadata schema
	Metadata MetadataParam `json:"metadata"`
}

func (r CloudV1AIClusterMetadataUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Metadata)
}
