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

// CloudV1SubnetMetadataService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1SubnetMetadataService] method instead.
type CloudV1SubnetMetadataService struct {
	Options []option.RequestOption
}

// NewCloudV1SubnetMetadataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1SubnetMetadataService(opts ...option.RequestOption) (r *CloudV1SubnetMetadataService) {
	r = &CloudV1SubnetMetadataService{}
	r.Options = opts
	return
}

// Create or update one or more metadata items for a subnet. If an item does not
// exist, it gets created. If an item already exists, it's value is overwritten
func (r *CloudV1SubnetMetadataService) New(ctx context.Context, projectID int64, regionID int64, subnetID string, body CloudV1SubnetMetadataNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v/%s/metadata", projectID, regionID, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Lists all metadata for a subnet
func (r *CloudV1SubnetMetadataService) List(ctx context.Context, projectID int64, regionID int64, subnetID string, opts ...option.RequestOption) (res *DetailedMetadataCollection, err error) {
	opts = append(r.Options[:], opts...)
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v/%s/metadata", projectID, regionID, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// All existing metadata is deleted and replaced with the metadata from the
// request.
func (r *CloudV1SubnetMetadataService) Replace(ctx context.Context, projectID int64, regionID int64, subnetID string, body CloudV1SubnetMetadataReplaceParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if subnetID == "" {
		err = errors.New("missing required subnet_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/subnets/%v/%v/%s/metadata", projectID, regionID, subnetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type CloudV1SubnetMetadataNewParams struct {
	// Used to process user input, accepting all fields except read-only or internal
	// keys (issued by a regular user). Users can specify key-value tags here.
	Body map[string]string `json:"body"`
}

func (r CloudV1SubnetMetadataNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CloudV1SubnetMetadataReplaceParams struct {
	// Used to process user input, accepting all fields except read-only or internal
	// keys (issued by a regular user). Users can specify key-value tags here.
	Body map[string]string `json:"body"`
}

func (r CloudV1SubnetMetadataReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
