// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1FileShareMetadataService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1FileShareMetadataService] method instead.
type CloudV1FileShareMetadataService struct {
	Options []option.RequestOption
}

// NewCloudV1FileShareMetadataService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1FileShareMetadataService(opts ...option.RequestOption) (r *CloudV1FileShareMetadataService) {
	r = &CloudV1FileShareMetadataService{}
	r.Options = opts
	return
}

// Create or update one or more tag for a file share. If an item does not exist, it
// gets created. If an item already exists, it's value is overwritten
func (r *CloudV1FileShareMetadataService) New(ctx context.Context, projectID int64, regionID int64, fileShareID string, body CloudV1FileShareMetadataNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s/metadata", projectID, regionID, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// List file share tags
func (r *CloudV1FileShareMetadataService) List(ctx context.Context, projectID int64, regionID int64, fileShareID string, opts ...option.RequestOption) (res *DetailedMetadataCollection, err error) {
	opts = append(r.Options[:], opts...)
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s/metadata", projectID, regionID, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// All existing metadata is deleted and replaced with the metadata from the request
func (r *CloudV1FileShareMetadataService) Replace(ctx context.Context, projectID int64, regionID int64, fileShareID string, body CloudV1FileShareMetadataReplaceParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s/metadata", projectID, regionID, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type DetailedMetadataCollection struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []DetailedMetadata             `json:"results,required"`
	JSON    detailedMetadataCollectionJSON `json:"-"`
}

// detailedMetadataCollectionJSON contains the JSON metadata for the struct
// [DetailedMetadataCollection]
type detailedMetadataCollectionJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DetailedMetadataCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detailedMetadataCollectionJSON) RawJSON() string {
	return r.raw
}

type CloudV1FileShareMetadataNewParams struct {
	// Used to process user input, accepting all fields except read-only or internal
	// keys (issued by a regular user). Users can specify key-value tags here.
	Body map[string]string `json:"body"`
}

func (r CloudV1FileShareMetadataNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CloudV1FileShareMetadataReplaceParams struct {
	// Used to process user input, accepting all fields except read-only or internal
	// keys (issued by a regular user). Users can specify key-value tags here.
	Body map[string]string `json:"body"`
}

func (r CloudV1FileShareMetadataReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
