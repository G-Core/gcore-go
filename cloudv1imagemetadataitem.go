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

// CloudV1ImageMetadataItemService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1ImageMetadataItemService] method instead.
type CloudV1ImageMetadataItemService struct {
	Options []option.RequestOption
}

// NewCloudV1ImageMetadataItemService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1ImageMetadataItemService(opts ...option.RequestOption) (r *CloudV1ImageMetadataItemService) {
	r = &CloudV1ImageMetadataItemService{}
	r.Options = opts
	return
}

// Get image metadata item by key
func (r *CloudV1ImageMetadataItemService) Get(ctx context.Context, projectID int64, regionID int64, imageID string, query CloudV1ImageMetadataItemGetParams, opts ...option.RequestOption) (res *MetadataItem, err error) {
	opts = append(r.Options[:], opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v/%s/metadata_item", projectID, regionID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete image metadata item by key
func (r *CloudV1ImageMetadataItemService) Delete(ctx context.Context, projectID int64, regionID int64, imageID string, body CloudV1ImageMetadataItemDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/images/%v/%v/%s/metadata_item", projectID, regionID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type CloudV1ImageMetadataItemGetParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV1ImageMetadataItemGetParams]'s query parameters as
// `url.Values`.
func (r CloudV1ImageMetadataItemGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1ImageMetadataItemDeleteParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV1ImageMetadataItemDeleteParams]'s query parameters as
// `url.Values`.
func (r CloudV1ImageMetadataItemDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
