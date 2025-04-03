// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2ClientQuotaService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2ClientQuotaService] method instead.
type CloudV2ClientQuotaService struct {
	Options               []option.RequestOption
	NotificationThreshold *CloudV2ClientQuotaNotificationThresholdService
}

// NewCloudV2ClientQuotaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV2ClientQuotaService(opts ...option.RequestOption) (r *CloudV2ClientQuotaService) {
	r = &CloudV2ClientQuotaService{}
	r.Options = opts
	r.NotificationThreshold = NewCloudV2ClientQuotaNotificationThresholdService(opts...)
	return
}

// Get combined client quotas, regional and global.
func (r *CloudV2ClientQuotaService) Get(ctx context.Context, query CloudV2ClientQuotaGetParams, opts ...option.RequestOption) (res *CloudV2ClientQuotaGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cloud/v2/client_quotas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CloudV2ClientQuotaGetResponse struct {
	// Global entity quotas
	GlobalQuotas GlobalQuotas `json:"global_quotas"`
	// Regional entity quotas. Only contains initialized quotas.
	RegionalQuotas []RegionalQuotas                  `json:"regional_quotas"`
	JSON           cloudV2ClientQuotaGetResponseJSON `json:"-"`
}

// cloudV2ClientQuotaGetResponseJSON contains the JSON metadata for the struct
// [CloudV2ClientQuotaGetResponse]
type cloudV2ClientQuotaGetResponseJSON struct {
	GlobalQuotas   apijson.Field
	RegionalQuotas apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CloudV2ClientQuotaGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV2ClientQuotaGetResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV2ClientQuotaGetParams struct {
	// Id of the client, if not present will be inferred from jwt.
	ClientID param.Field[int64] `query:"client_id"`
}

// URLQuery serializes [CloudV2ClientQuotaGetParams]'s query parameters as
// `url.Values`.
func (r CloudV2ClientQuotaGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
