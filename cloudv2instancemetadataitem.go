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

// CloudV2InstanceMetadataItemService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2InstanceMetadataItemService] method instead.
type CloudV2InstanceMetadataItemService struct {
	Options []option.RequestOption
}

// NewCloudV2InstanceMetadataItemService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV2InstanceMetadataItemService(opts ...option.RequestOption) (r *CloudV2InstanceMetadataItemService) {
	r = &CloudV2InstanceMetadataItemService{}
	r.Options = opts
	return
}

// Get instance metadata item by key
func (r *CloudV2InstanceMetadataItemService) Get(ctx context.Context, projectID int64, regionID int64, instanceID string, query CloudV2InstanceMetadataItemGetParams, opts ...option.RequestOption) (res *MetadataItem, err error) {
	opts = append(r.Options[:], opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/instances/%v/%v/%s/metadata_item", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete instance metadata item by key
func (r *CloudV2InstanceMetadataItemService) Delete(ctx context.Context, projectID int64, regionID int64, instanceID string, body CloudV2InstanceMetadataItemDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/instances/%v/%v/%s/metadata_item", projectID, regionID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type CloudV2InstanceMetadataItemGetParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV2InstanceMetadataItemGetParams]'s query parameters as
// `url.Values`.
func (r CloudV2InstanceMetadataItemGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV2InstanceMetadataItemDeleteParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV2InstanceMetadataItemDeleteParams]'s query parameters
// as `url.Values`.
func (r CloudV2InstanceMetadataItemDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
