// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
)

// DomainSettingService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainSettingService] method instead.
type DomainSettingService struct {
	Options []option.RequestOption
}

// NewDomainSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDomainSettingService(opts ...option.RequestOption) (r DomainSettingService) {
	r = DomainSettingService{}
	r.Options = opts
	return
}

// Update settings for a specific domain
func (r *DomainSettingService) Update(ctx context.Context, domainID int64, body DomainSettingUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/settings", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Retrieve settings for a specific domain
func (r *DomainSettingService) Get(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *WaapDomainSettingsModel, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/settings", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DomainSettingUpdateParams struct {
	// Editable API settings of a domain
	API DomainSettingUpdateParamsAPI `json:"api,omitzero"`
	// Editable DDoS settings for a domain.
	DDOS DomainSettingUpdateParamsDDOS `json:"ddos,omitzero"`
	paramObj
}

func (r DomainSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainSettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Editable API settings of a domain
type DomainSettingUpdateParamsAPI struct {
	// Indicates if the domain is an API domain. All requests to an API domain are
	// treated as API requests. If this is set to true then the `api_urls` field is
	// ignored.
	IsAPI param.Opt[bool] `json:"is_api,omitzero"`
	// The API URLs for a domain. If your domain has a common base URL for all API
	// paths, it can be set here
	APIURLs []string `json:"api_urls,omitzero"`
	paramObj
}

func (r DomainSettingUpdateParamsAPI) MarshalJSON() (data []byte, err error) {
	type shadow DomainSettingUpdateParamsAPI
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainSettingUpdateParamsAPI) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Editable DDoS settings for a domain.
type DomainSettingUpdateParamsDDOS struct {
	// The burst threshold detects sudden rises in traffic. If it is met and the number
	// of requests is at least five times the last 2-second interval, DDoS protection
	// will activate. Default is 1000.
	BurstThreshold param.Opt[int64] `json:"burst_threshold,omitzero"`
	// The global threshold is responsible for identifying DDoS attacks with a slow
	// rise in traffic. If the threshold is met and the current number of requests is
	// at least double that of the previous 10-second window, DDoS protection will
	// activate. Default is 5000.
	GlobalThreshold param.Opt[int64] `json:"global_threshold,omitzero"`
	paramObj
}

func (r DomainSettingUpdateParamsDDOS) MarshalJSON() (data []byte, err error) {
	type shadow DomainSettingUpdateParamsDDOS
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainSettingUpdateParamsDDOS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
