// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1K8Service contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1K8Service] method instead.
type CloudV1K8Service struct {
	Options []option.RequestOption
}

// NewCloudV1K8Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1K8Service(opts ...option.RequestOption) (r *CloudV1K8Service) {
	r = &CloudV1K8Service{}
	r.Options = opts
	return
}

// Retrieve a list of flavors for k8s pool. When the include_prices query parameter
// is specified, the list shows prices. A client in trial mode gets all price
// values as 0. If you get Pricing Error contact the support
func (r *CloudV1K8Service) ListFlavors(ctx context.Context, projectID int64, regionID int64, query CloudV1K8ListFlavorsParams, opts ...option.RequestOption) (res *FlavorList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/k8s/%v/%v/flavors", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CloudV1K8ListFlavorsParams struct {
	// Set to false to include gpu flavors. Defaults True
	ExcludeGPU param.Field[bool] `query:"exclude_gpu"`
	// Set to true if the response should include flavor prices
	IncludePrices param.Field[bool] `query:"include_prices"`
}

// URLQuery serializes [CloudV1K8ListFlavorsParams]'s query parameters as
// `url.Values`.
func (r CloudV1K8ListFlavorsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
