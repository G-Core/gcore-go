// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1AIClusterMetadataItemService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1AIClusterMetadataItemService] method instead.
type CloudV1AIClusterMetadataItemService struct {
	Options []option.RequestOption
}

// NewCloudV1AIClusterMetadataItemService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1AIClusterMetadataItemService(opts ...option.RequestOption) (r *CloudV1AIClusterMetadataItemService) {
	r = &CloudV1AIClusterMetadataItemService{}
	r.Options = opts
	return
}

// Get AI cluster metadata item by key
func (r *CloudV1AIClusterMetadataItemService) Get(ctx context.Context, projectID int64, regionID int64, clusterID string, query CloudV1AIClusterMetadataItemGetParams, opts ...option.RequestOption) (res *MetadataItem, err error) {
	opts = append(r.Options[:], opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/cluster/%v/%v/%s/metadata_item", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete AI cluster metadata item by key
func (r *CloudV1AIClusterMetadataItemService) Delete(ctx context.Context, projectID int64, regionID int64, clusterID string, body CloudV1AIClusterMetadataItemDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if clusterID == "" {
		err = errors.New("missing required cluster_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ai/cluster/%v/%v/%s/metadata_item", projectID, regionID, clusterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Metadata item schema
type MetadataItem struct {
	// Metadata item key. Maximum size is 255 bytes.
	Key string `json:"key,required"`
	// Flag that shows if metadata item can be modified or not
	ReadOnly bool `json:"read_only,required"`
	// Metadata item value. Maximum size is 1024 bytes.
	Value string           `json:"value,required"`
	JSON  metadataItemJSON `json:"-"`
}

// metadataItemJSON contains the JSON metadata for the struct [MetadataItem]
type metadataItemJSON struct {
	Key         apijson.Field
	ReadOnly    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetadataItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metadataItemJSON) RawJSON() string {
	return r.raw
}

type CloudV1AIClusterMetadataItemGetParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV1AIClusterMetadataItemGetParams]'s query parameters
// as `url.Values`.
func (r CloudV1AIClusterMetadataItemGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1AIClusterMetadataItemDeleteParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV1AIClusterMetadataItemDeleteParams]'s query
// parameters as `url.Values`.
func (r CloudV1AIClusterMetadataItemDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
