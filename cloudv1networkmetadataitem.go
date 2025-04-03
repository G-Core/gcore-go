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

// CloudV1NetworkMetadataItemService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1NetworkMetadataItemService] method instead.
type CloudV1NetworkMetadataItemService struct {
	Options []option.RequestOption
}

// NewCloudV1NetworkMetadataItemService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1NetworkMetadataItemService(opts ...option.RequestOption) (r *CloudV1NetworkMetadataItemService) {
	r = &CloudV1NetworkMetadataItemService{}
	r.Options = opts
	return
}

// Get network metadata item by key
func (r *CloudV1NetworkMetadataItemService) Get(ctx context.Context, projectID int64, regionID int64, networkID string, query CloudV1NetworkMetadataItemGetParams, opts ...option.RequestOption) (res *MetadataItem, err error) {
	opts = append(r.Options[:], opts...)
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s/metadata_item", projectID, regionID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete network metadata item by key
func (r *CloudV1NetworkMetadataItemService) Delete(ctx context.Context, projectID int64, regionID int64, networkID string, body CloudV1NetworkMetadataItemDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/networks/%v/%v/%s/metadata_item", projectID, regionID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type CloudV1NetworkMetadataItemGetParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV1NetworkMetadataItemGetParams]'s query parameters as
// `url.Values`.
func (r CloudV1NetworkMetadataItemGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1NetworkMetadataItemDeleteParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV1NetworkMetadataItemDeleteParams]'s query parameters
// as `url.Values`.
func (r CloudV1NetworkMetadataItemDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
