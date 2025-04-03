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

// CloudV1SecuritygroupMetadataService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1SecuritygroupMetadataService] method instead.
type CloudV1SecuritygroupMetadataService struct {
	Options []option.RequestOption
}

// NewCloudV1SecuritygroupMetadataService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1SecuritygroupMetadataService(opts ...option.RequestOption) (r *CloudV1SecuritygroupMetadataService) {
	r = &CloudV1SecuritygroupMetadataService{}
	r.Options = opts
	return
}

// Create or update one or more metadata items for a security group. If an item
// does not exist, it gets created. If an item already exists, it's value is
// overwritten
func (r *CloudV1SecuritygroupMetadataService) New(ctx context.Context, projectID int64, regionID int64, securityGroupID string, body CloudV1SecuritygroupMetadataNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if securityGroupID == "" {
		err = errors.New("missing required security_group_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s/metadata", projectID, regionID, securityGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Lists all metadata for a security group
func (r *CloudV1SecuritygroupMetadataService) List(ctx context.Context, projectID int64, regionID int64, securityGroupID string, opts ...option.RequestOption) (res *DetailedMetadataCollection, err error) {
	opts = append(r.Options[:], opts...)
	if securityGroupID == "" {
		err = errors.New("missing required security_group_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s/metadata", projectID, regionID, securityGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// All existing metadata is deleted and replaced with the metadata from the
// request.
func (r *CloudV1SecuritygroupMetadataService) Replace(ctx context.Context, projectID int64, regionID int64, securityGroupID string, body CloudV1SecuritygroupMetadataReplaceParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if securityGroupID == "" {
		err = errors.New("missing required security_group_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/securitygroups/%v/%v/%s/metadata", projectID, regionID, securityGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type RawMetadataParam map[string]interface{}

type CloudV1SecuritygroupMetadataNewParams struct {
	// Serializer used for accepting user's input. Any field is valid for this
	// serializer except for read-only or internal (issued by regular user) keys.
	RawMetadata RawMetadataParam `json:"raw_metadata"`
}

func (r CloudV1SecuritygroupMetadataNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RawMetadata)
}

type CloudV1SecuritygroupMetadataReplaceParams struct {
	// Serializer used for accepting user's input. Any field is valid for this
	// serializer except for read-only or internal (issued by regular user) keys.
	RawMetadata RawMetadataParam `json:"raw_metadata"`
}

func (r CloudV1SecuritygroupMetadataReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RawMetadata)
}
