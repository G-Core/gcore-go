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

// CloudV1FileShareMetadataItemService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1FileShareMetadataItemService] method instead.
type CloudV1FileShareMetadataItemService struct {
	Options []option.RequestOption
}

// NewCloudV1FileShareMetadataItemService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1FileShareMetadataItemService(opts ...option.RequestOption) (r *CloudV1FileShareMetadataItemService) {
	r = &CloudV1FileShareMetadataItemService{}
	r.Options = opts
	return
}

// Get file share tag by key
func (r *CloudV1FileShareMetadataItemService) Get(ctx context.Context, projectID int64, regionID int64, fileShareID string, query CloudV1FileShareMetadataItemGetParams, opts ...option.RequestOption) (res *DetailedMetadata, err error) {
	opts = append(r.Options[:], opts...)
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s/metadata_item", projectID, regionID, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete file share tag by key
func (r *CloudV1FileShareMetadataItemService) Delete(ctx context.Context, projectID int64, regionID int64, fileShareID string, body CloudV1FileShareMetadataItemDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s/metadata_item", projectID, regionID, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Serializer used for displaying a specific resource tag information in formatted
// version.
type DetailedMetadata struct {
	// Tag key. The maximum size for a key is 255 bytes.
	Key string `json:"key,required"`
	// Whether the tag key is read-only, being managed by the system
	ReadOnly bool `json:"read_only,required"`
	// Tag value. The maximum size for a value is 1024 bytes.
	Value string               `json:"value,required"`
	JSON  detailedMetadataJSON `json:"-"`
}

// detailedMetadataJSON contains the JSON metadata for the struct
// [DetailedMetadata]
type detailedMetadataJSON struct {
	Key         apijson.Field
	ReadOnly    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DetailedMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailedMetadataJSON) RawJSON() string {
	return r.raw
}

type CloudV1FileShareMetadataItemGetParams struct {
	// Tag key
	Key param.Field[string] `query:"key,required"`
}

// URLQuery serializes [CloudV1FileShareMetadataItemGetParams]'s query parameters
// as `url.Values`.
func (r CloudV1FileShareMetadataItemGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CloudV1FileShareMetadataItemDeleteParams struct {
	// Tag key
	Key param.Field[string] `query:"key,required"`
}

// URLQuery serializes [CloudV1FileShareMetadataItemDeleteParams]'s query
// parameters as `url.Values`.
func (r CloudV1FileShareMetadataItemDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
