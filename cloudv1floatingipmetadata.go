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

// CloudV1FloatingipMetadataService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1FloatingipMetadataService] method instead.
type CloudV1FloatingipMetadataService struct {
	Options []option.RequestOption
}

// NewCloudV1FloatingipMetadataService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1FloatingipMetadataService(opts ...option.RequestOption) (r *CloudV1FloatingipMetadataService) {
	r = &CloudV1FloatingipMetadataService{}
	r.Options = opts
	return
}

// Create or update one or more metadata items for a floating IP. If an item does
// not exist, it gets created. If an item already exists, it's value is overwritten
func (r *CloudV1FloatingipMetadataService) New(ctx context.Context, projectID int64, regionID int64, pk string, body CloudV1FloatingipMetadataNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s/metadata", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Lists all metadata for a floating IP
func (r *CloudV1FloatingipMetadataService) List(ctx context.Context, projectID int64, regionID int64, pk string, opts ...option.RequestOption) (res *DetailedMetadataCollection, err error) {
	opts = append(r.Options[:], opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s/metadata", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// All existing metadata is deleted and replaced with the metadata from the
// request.
func (r *CloudV1FloatingipMetadataService) Replace(ctx context.Context, projectID int64, regionID int64, pk string, body CloudV1FloatingipMetadataReplaceParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if pk == "" {
		err = errors.New("missing required pk parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/floatingips/%v/%v/%s/metadata", projectID, regionID, pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type CloudV1FloatingipMetadataNewParams struct {
	// Serializer used for accepting user's input. Any field is valid for this
	// serializer except for read-only or internal (issued by regular user) keys.
	RawMetadata RawMetadataParam `json:"raw_metadata"`
}

func (r CloudV1FloatingipMetadataNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RawMetadata)
}

type CloudV1FloatingipMetadataReplaceParams struct {
	// Serializer used for accepting user's input. Any field is valid for this
	// serializer except for read-only or internal (issued by regular user) keys.
	RawMetadata RawMetadataParam `json:"raw_metadata"`
}

func (r CloudV1FloatingipMetadataReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RawMetadata)
}
