// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1LaaStatusService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1LaaStatusService] method instead.
type CloudV1LaaStatusService struct {
	Options []option.RequestOption
}

// NewCloudV1LaaStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1LaaStatusService(opts ...option.RequestOption) (r *CloudV1LaaStatusService) {
	r = &CloudV1LaaStatusService{}
	r.Options = opts
	return
}

// Update LaaS status
func (r *CloudV1LaaStatusService) Update(ctx context.Context, projectID int64, regionID int64, body CloudV1LaaStatusUpdateParams, opts ...option.RequestOption) (res *LaasStatus, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/laas/%v/%v/status", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get LaaS status
func (r *CloudV1LaaStatusService) Get(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *LaasStatus, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/laas/%v/%v/status", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LaasStatus struct {
	// Boolean showing whether client is initialized with LaaS
	IsInitialized bool `json:"is_initialized,required"`
	// Client namespace
	Namespace string         `json:"namespace,required,nullable"`
	JSON      laasStatusJSON `json:"-"`
}

// laasStatusJSON contains the JSON metadata for the struct [LaasStatus]
type laasStatusJSON struct {
	IsInitialized apijson.Field
	Namespace     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LaasStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r laasStatusJSON) RawJSON() string {
	return r.raw
}

type CloudV1LaaStatusUpdateParams struct {
	// Boolean showing whether client is initialized with LaaS
	IsInitialized param.Field[bool] `json:"is_initialized,required"`
	// Optional. Custom namespace
	Namespace param.Field[string] `json:"namespace"`
}

func (r CloudV1LaaStatusUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
