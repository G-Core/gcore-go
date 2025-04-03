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

// CloudV1LoadbalancerMetadataService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1LoadbalancerMetadataService] method instead.
type CloudV1LoadbalancerMetadataService struct {
	Options []option.RequestOption
}

// NewCloudV1LoadbalancerMetadataService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1LoadbalancerMetadataService(opts ...option.RequestOption) (r *CloudV1LoadbalancerMetadataService) {
	r = &CloudV1LoadbalancerMetadataService{}
	r.Options = opts
	return
}

// Create or update one or more metadata items for a loadbalancer. If an item does
// not exist, it gets created. If an item already exists, it's value is overwritten
func (r *CloudV1LoadbalancerMetadataService) New(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, body CloudV1LoadbalancerMetadataNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/metadata", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Lists all metadata for a loadbalancer
func (r *CloudV1LoadbalancerMetadataService) List(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, opts ...option.RequestOption) (res *DetailedMetadataCollection, err error) {
	opts = append(r.Options[:], opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/metadata", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// All existing metadata is deleted and replaced with the metadata from the
// request.
func (r *CloudV1LoadbalancerMetadataService) Replace(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, body CloudV1LoadbalancerMetadataReplaceParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/loadbalancers/%v/%v/%s/metadata", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type CloudV1LoadbalancerMetadataNewParams struct {
	// Serializer used for accepting user's input. Any field is valid for this
	// serializer except for read-only or internal (issued by regular user) keys.
	RawMetadata RawMetadataParam `json:"raw_metadata"`
}

func (r CloudV1LoadbalancerMetadataNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RawMetadata)
}

type CloudV1LoadbalancerMetadataReplaceParams struct {
	// Serializer used for accepting user's input. Any field is valid for this
	// serializer except for read-only or internal (issued by regular user) keys.
	RawMetadata RawMetadataParam `json:"raw_metadata"`
}

func (r CloudV1LoadbalancerMetadataReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RawMetadata)
}
