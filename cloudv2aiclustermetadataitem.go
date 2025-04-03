// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2AIClusterMetadataItemService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2AIClusterMetadataItemService] method instead.
type CloudV2AIClusterMetadataItemService struct {
	Options []option.RequestOption
}

// NewCloudV2AIClusterMetadataItemService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV2AIClusterMetadataItemService(opts ...option.RequestOption) (r *CloudV2AIClusterMetadataItemService) {
	r = &CloudV2AIClusterMetadataItemService{}
	r.Options = opts
	return
}

// Get AI cluster metadata item by key
func (r *CloudV2AIClusterMetadataItemService) Get(ctx context.Context, projectID int64, regionID int64, clusterID string, query CloudV2AIClusterMetadataItemGetParams, opts ...option.RequestOption) (res *MetadataItem, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s/metadata_item", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete AI cluster metadata item by key
func (r *CloudV2AIClusterMetadataItemService) Delete(ctx context.Context, projectID int64, regionID int64, clusterID string, body CloudV2AIClusterMetadataItemDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/ai/clusters/%v/%v/%s/metadata_item", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type CloudV2AIClusterMetadataItemGetParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV2AIClusterMetadataItemGetParams]'s query parameters
// as `url.Values`.
func (r CloudV2AIClusterMetadataItemGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV2AIClusterMetadataItemDeleteParams struct {
	// Meta key
	Key param.Field[string] `query:"key,required"`
}

// URLQuery serializes [CloudV2AIClusterMetadataItemDeleteParams]'s query
// parameters as `url.Values`.
func (r CloudV2AIClusterMetadataItemDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
