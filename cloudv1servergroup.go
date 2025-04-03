// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1ServergroupService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1ServergroupService] method instead.
type CloudV1ServergroupService struct {
	Options []option.RequestOption
}

// NewCloudV1ServergroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1ServergroupService(opts ...option.RequestOption) (r *CloudV1ServergroupService) {
	r = &CloudV1ServergroupService{}
	r.Options = opts
	return
}

// Create an affinity or anti-affinity or soft-anti-affinity server group
func (r *CloudV1ServergroupService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1ServergroupNewParams, opts ...option.RequestOption) (res *ServerGroup, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/servergroups/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get server group
func (r *CloudV1ServergroupService) Get(ctx context.Context, projectID int64, regionID int64, servergroupID string, opts ...option.RequestOption) (res *ServerGroup, err error) {
	opts = append(r.Options[:], opts...)
	if servergroupID == "" {
		err = errors.New("missing required servergroup_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/servergroups/%v/%v/%s", projectID, regionID, servergroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List server groups
func (r *CloudV1ServergroupService) List(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *CloudV1ServergroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/servergroups/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete server group
func (r *CloudV1ServergroupService) Delete(ctx context.Context, projectID int64, regionID int64, servergroupID string, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = append(r.Options[:], opts...)
	if servergroupID == "" {
		err = errors.New("missing required servergroup_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/servergroups/%v/%v/%s", projectID, regionID, servergroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ServerGroup struct {
	// The list of instances in this server group.
	Instances []ServerGroupInstance `json:"instances,required"`
	// The name of the server group.
	Name string `json:"name,required"`
	// The server group policy. Options are: anti-affinity, affinity, or
	// soft-anti-affinity.
	Policy string `json:"policy,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// The ID of the server group.
	ServergroupID string          `json:"servergroup_id,required"`
	JSON          serverGroupJSON `json:"-"`
}

// serverGroupJSON contains the JSON metadata for the struct [ServerGroup]
type serverGroupJSON struct {
	Instances     apijson.Field
	Name          apijson.Field
	Policy        apijson.Field
	ProjectID     apijson.Field
	Region        apijson.Field
	RegionID      apijson.Field
	ServergroupID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ServerGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serverGroupJSON) RawJSON() string {
	return r.raw
}

type ServerGroupInstance struct {
	// The ID of the instance, corresponding to the attribute 'id'.
	InstanceID string `json:"instance_id,required"`
	// The name of the instance, corresponding to the attribute 'name'.
	InstanceName string                  `json:"instance_name,required"`
	JSON         serverGroupInstanceJSON `json:"-"`
}

// serverGroupInstanceJSON contains the JSON metadata for the struct
// [ServerGroupInstance]
type serverGroupInstanceJSON struct {
	InstanceID   apijson.Field
	InstanceName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServerGroupInstance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serverGroupInstanceJSON) RawJSON() string {
	return r.raw
}

type ServerGroupPolicy string

const (
	ServerGroupPolicyAffinity         ServerGroupPolicy = "affinity"
	ServerGroupPolicyAntiAffinity     ServerGroupPolicy = "anti-affinity"
	ServerGroupPolicySoftAntiAffinity ServerGroupPolicy = "soft-anti-affinity"
)

func (r ServerGroupPolicy) IsKnown() bool {
	switch r {
	case ServerGroupPolicyAffinity, ServerGroupPolicyAntiAffinity, ServerGroupPolicySoftAntiAffinity:
		return true
	}
	return false
}

type CloudV1ServergroupListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []ServerGroup                      `json:"results,required"`
	JSON    cloudV1ServergroupListResponseJSON `json:"-"`
}

// cloudV1ServergroupListResponseJSON contains the JSON metadata for the struct
// [CloudV1ServergroupListResponse]
type cloudV1ServergroupListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1ServergroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1ServergroupListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1ServergroupNewParams struct {
	// The name of the server group.
	Name param.Field[string] `json:"name,required"`
	// The server group policy.
	Policy param.Field[ServerGroupPolicy] `json:"policy,required"`
}

func (r CloudV1ServergroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
