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

// CloudV2LoadbalancerService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2LoadbalancerService] method instead.
type CloudV2LoadbalancerService struct {
	Options  []option.RequestOption
	Metadata *CloudV2LoadbalancerMetadataService
}

// NewCloudV2LoadbalancerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV2LoadbalancerService(opts ...option.RequestOption) (r *CloudV2LoadbalancerService) {
	r = &CloudV2LoadbalancerService{}
	r.Options = opts
	r.Metadata = NewCloudV2LoadbalancerMetadataService(opts...)
	return
}

// Delete loadbalancer metadata item by key
func (r *CloudV2LoadbalancerService) DeleteMetadataItem(ctx context.Context, projectID int64, regionID int64, loadbalancerID string, body CloudV2LoadbalancerDeleteMetadataItemParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if loadbalancerID == "" {
		err = errors.New("missing required loadbalancer_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/loadbalancers/%v/%v/%s/metadata_item", projectID, regionID, loadbalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type CloudV2LoadbalancerDeleteMetadataItemParams struct {
	// Meta key
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [CloudV2LoadbalancerDeleteMetadataItemParams]'s query
// parameters as `url.Values`.
func (r CloudV2LoadbalancerDeleteMetadataItemParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
