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

// CloudV2LoadbalancerMetadataService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2LoadbalancerMetadataService] method instead.
type CloudV2LoadbalancerMetadataService struct {
	Options []option.RequestOption
}

// NewCloudV2LoadbalancerMetadataService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV2LoadbalancerMetadataService(opts ...option.RequestOption) (r *CloudV2LoadbalancerMetadataService) {
	r = &CloudV2LoadbalancerMetadataService{}
	r.Options = opts
	return
}

// Create or update one or more metadata items for a loadbalancer. If an item does
// not exist, it gets created. If an item already exists, it's value is overwritten
func (r *CloudV2LoadbalancerMetadataService) NewOrUpdate(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, body CloudV2LoadbalancerMetadataNewOrUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/loadbalancers/%v/%v/%s/metadata", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// All existing metadata is deleted and replaced with the metadata from the
// request.
func (r *CloudV2LoadbalancerMetadataService) Replace(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, body CloudV2LoadbalancerMetadataReplaceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/loadbalancers/%v/%v/%s/metadata", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CloudV2LoadbalancerMetadataNewOrUpdateParams struct {
	// Serializer used for accepting user's input. Any field is valid for this
	// serializer except for read-only or internal (issued by regular user) keys.
	RawMetadata RawMetadataParam `json:"raw_metadata"`
}

func (r CloudV2LoadbalancerMetadataNewOrUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RawMetadata)
}

type CloudV2LoadbalancerMetadataReplaceParams struct {
	// Serializer used for accepting user's input. Any field is valid for this
	// serializer except for read-only or internal (issued by regular user) keys.
	RawMetadata RawMetadataParam `json:"raw_metadata"`
}

func (r CloudV2LoadbalancerMetadataReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RawMetadata)
}
