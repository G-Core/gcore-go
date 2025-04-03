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

// CloudV1DbaaStatusService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1DbaaStatusService] method instead.
type CloudV1DbaaStatusService struct {
	Options []option.RequestOption
}

// NewCloudV1DbaaStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1DbaaStatusService(opts ...option.RequestOption) (r *CloudV1DbaaStatusService) {
	r = &CloudV1DbaaStatusService{}
	r.Options = opts
	return
}

// Activate or deactivate DBAAS Service for a specified project.
func (r *CloudV1DbaaStatusService) Update(ctx context.Context, projectID int64, regionID int64, body CloudV1DbaaStatusUpdateParams, opts ...option.RequestOption) (res *DbaasStatus, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/dbaas/status/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get DBAAS service status for a specified project. It may be initialized or not
func (r *CloudV1DbaaStatusService) Get(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *DbaasStatus, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/dbaas/status/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DbaasStatus struct {
	// DBAAS service status
	IsInitialized bool            `json:"is_initialized,required"`
	JSON          dbaasStatusJSON `json:"-"`
}

// dbaasStatusJSON contains the JSON metadata for the struct [DbaasStatus]
type dbaasStatusJSON struct {
	IsInitialized apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DbaasStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dbaasStatusJSON) RawJSON() string {
	return r.raw
}

type DbaasStatusParam struct {
	// DBAAS service status
	IsInitialized param.Field[bool] `json:"is_initialized,required"`
}

func (r DbaasStatusParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1DbaaStatusUpdateParams struct {
	DbaasStatus DbaasStatusParam `json:"dbaas_status,required"`
}

func (r CloudV1DbaaStatusUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DbaasStatus)
}
