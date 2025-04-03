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

// CloudV1SecuritygroupMetadataItemService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1SecuritygroupMetadataItemService] method instead.
type CloudV1SecuritygroupMetadataItemService struct {
	Options []option.RequestOption
}

// NewCloudV1SecuritygroupMetadataItemService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1SecuritygroupMetadataItemService(opts ...option.RequestOption) (r *CloudV1SecuritygroupMetadataItemService) {
	r = &CloudV1SecuritygroupMetadataItemService{}
	r.Options = opts
	return
}

// Get network metadata item by key
func (r *CloudV1SecuritygroupMetadataItemService) Get(ctx context.Context, projectID int64, regionID int64, pk string, query CloudV1SecuritygroupMetadataItemGetParams, opts ...option.RequestOption) (res *DetailedMetadata, err error) {
	opts = append(r.Options[:], opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s/metadata_item", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete security group metadata item by key
func (r *CloudV1SecuritygroupMetadataItemService) Delete(ctx context.Context, projectID int64, regionID int64, pk string, body CloudV1SecuritygroupMetadataItemDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s/metadata_item", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type CloudV1SecuritygroupMetadataItemGetParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV1SecuritygroupMetadataItemGetParams]'s query
// parameters as `url.Values`.
func (r CloudV1SecuritygroupMetadataItemGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1SecuritygroupMetadataItemDeleteParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV1SecuritygroupMetadataItemDeleteParams]'s query
// parameters as `url.Values`.
func (r CloudV1SecuritygroupMetadataItemDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
