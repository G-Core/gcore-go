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

// CloudV1LoadbalancerMetadataItemService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1LoadbalancerMetadataItemService] method instead.
type CloudV1LoadbalancerMetadataItemService struct {
	Options []option.RequestOption
}

// NewCloudV1LoadbalancerMetadataItemService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1LoadbalancerMetadataItemService(opts ...option.RequestOption) (r *CloudV1LoadbalancerMetadataItemService) {
	r = &CloudV1LoadbalancerMetadataItemService{}
	r.Options = opts
	return
}

// Get loadbalancer metadata item by key
func (r *CloudV1LoadbalancerMetadataItemService) Get(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, query CloudV1LoadbalancerMetadataItemGetParams, opts ...option.RequestOption) (res *DetailedMetadata, err error) {
	opts = append(r.Options[:], opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/metadata_item", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete loadbalancer metadata item by key
func (r *CloudV1LoadbalancerMetadataItemService) Delete(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, body CloudV1LoadbalancerMetadataItemDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/metadata_item", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type CloudV1LoadbalancerMetadataItemGetParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV1LoadbalancerMetadataItemGetParams]'s query
// parameters as `url.Values`.
func (r CloudV1LoadbalancerMetadataItemGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1LoadbalancerMetadataItemDeleteParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV1LoadbalancerMetadataItemDeleteParams]'s query
// parameters as `url.Values`.
func (r CloudV1LoadbalancerMetadataItemDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
