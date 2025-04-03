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

// CloudV1VolumeMetadataItemService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1VolumeMetadataItemService] method instead.
type CloudV1VolumeMetadataItemService struct {
	Options []option.RequestOption
}

// NewCloudV1VolumeMetadataItemService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1VolumeMetadataItemService(opts ...option.RequestOption) (r *CloudV1VolumeMetadataItemService) {
	r = &CloudV1VolumeMetadataItemService{}
	r.Options = opts
	return
}

// Get volume metadata item by key
func (r *CloudV1VolumeMetadataItemService) Get(ctx context.Context, projectID int64, regionID int64, volumeID string, query CloudV1VolumeMetadataItemGetParams, opts ...option.RequestOption) (res *MetadataItem, err error) {
	opts = append(r.Options[:], opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/metadata_item", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete volume metadata item by key
func (r *CloudV1VolumeMetadataItemService) Delete(ctx context.Context, projectID int64, regionID int64, volumeID string, body CloudV1VolumeMetadataItemDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/metadata_item", projectID, regionID, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type CloudV1VolumeMetadataItemGetParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV1VolumeMetadataItemGetParams]'s query parameters as
// `url.Values`.
func (r CloudV1VolumeMetadataItemGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1VolumeMetadataItemDeleteParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV1VolumeMetadataItemDeleteParams]'s query parameters
// as `url.Values`.
func (r CloudV1VolumeMetadataItemDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
