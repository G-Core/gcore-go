// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1DbaaPostgreClusterService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1DbaaPostgreClusterService] method instead.
type CloudV1DbaaPostgreClusterService struct {
	Options []option.RequestOption
	Users   *CloudV1DbaaPostgreClusterUserService
}

// NewCloudV1DbaaPostgreClusterService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1DbaaPostgreClusterService(opts ...option.RequestOption) (r *CloudV1DbaaPostgreClusterService) {
	r = &CloudV1DbaaPostgreClusterService{}
	r.Options = opts
	r.Users = NewCloudV1DbaaPostgreClusterUserService(opts...)
	return
}

// Create DBAAS PostgreSQL cluster.
func (r *CloudV1DbaaPostgreClusterService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1DbaaPostgreClusterNewParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/clusters/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get DBAAS Postgres Cluster.
func (r *CloudV1DbaaPostgreClusterService) Get(ctx context.Context, projectID int64, regionID int64, clusterName string, opts ...option.RequestOption) (res *CloudV1DbaaPostgreClusterGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/clusters/%v/%v/%s", projectID, regionID, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch DBAAS Postgres Cluster.
func (r *CloudV1DbaaPostgreClusterService) Update(ctx context.Context, projectID int64, regionID int64, clusterName string, body CloudV1DbaaPostgreClusterUpdateParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/clusters/%v/%v/%s", projectID, regionID, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List DBAAS PostgreSQL Clusters.
func (r *CloudV1DbaaPostgreClusterService) List(ctx context.Context, projectID int64, regionID int64, query CloudV1DbaaPostgreClusterListParams, opts ...option.RequestOption) (res *CloudV1DbaaPostgreClusterListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/clusters/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete DBAAS PostgreSQL cluster.
func (r *CloudV1DbaaPostgreClusterService) Delete(ctx context.Context, projectID int64, regionID int64, clusterName string, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/clusters/%v/%v/%s", projectID, regionID, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ClusterStatus string

const (
	ClusterStatusDeleting  ClusterStatus = "DELETING"
	ClusterStatusFailed    ClusterStatus = "FAILED"
	ClusterStatusPreparing ClusterStatus = "PREPARING"
	ClusterStatusReady     ClusterStatus = "READY"
	ClusterStatusUnhealthy ClusterStatus = "UNHEALTHY"
	ClusterStatusUnknown   ClusterStatus = "UNKNOWN"
	ClusterStatusUpdating  ClusterStatus = "UPDATING"
)

func (r ClusterStatus) IsKnown() bool {
	switch r {
	case ClusterStatusDeleting, ClusterStatusFailed, ClusterStatusPreparing, ClusterStatusReady, ClusterStatusUnhealthy, ClusterStatusUnknown, ClusterStatusUpdating:
		return true
	}
	return false
}

type CloudV1DbaaPostgreClusterGetResponse struct {
	ClusterName string                                         `json:"cluster_name,required"`
	CreatedAt   time.Time                                      `json:"created_at,required" format:"date-time"`
	Databases   []CloudV1DbaaPostgreClusterGetResponseDatabase `json:"databases,required"`
	// Instance RAM and CPU
	Flavor           FlavorDatabase                              `json:"flavor,required"`
	HighAvailability HighAvailabilityOptions                     `json:"high_availability,required,nullable"`
	Network          CloudV1DbaaPostgreClusterGetResponseNetwork `json:"network,required"`
	// Main PG configuration
	PgServerConfiguration PostgresqlServerConfig `json:"pg_server_configuration,required"`
	// Current cluster status
	Status ClusterStatus `json:"status,required"`
	// PG's storage configuration
	Storage Storage                                    `json:"storage,required"`
	Users   []CloudV1DbaaPostgreClusterGetResponseUser `json:"users,required"`
	JSON    cloudV1DbaaPostgreClusterGetResponseJSON   `json:"-"`
}

// cloudV1DbaaPostgreClusterGetResponseJSON contains the JSON metadata for the
// struct [CloudV1DbaaPostgreClusterGetResponse]
type cloudV1DbaaPostgreClusterGetResponseJSON struct {
	ClusterName           apijson.Field
	CreatedAt             apijson.Field
	Databases             apijson.Field
	Flavor                apijson.Field
	HighAvailability      apijson.Field
	Network               apijson.Field
	PgServerConfiguration apijson.Field
	Status                apijson.Field
	Storage               apijson.Field
	Users                 apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CloudV1DbaaPostgreClusterGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DbaaPostgreClusterGetResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1DbaaPostgreClusterGetResponseDatabase struct {
	// Database name
	Name string `json:"name,required"`
	// Database owner from users list
	Owner string `json:"owner,required"`
	// Size in bytes
	Size int64                                            `json:"size,required"`
	JSON cloudV1DbaaPostgreClusterGetResponseDatabaseJSON `json:"-"`
}

// cloudV1DbaaPostgreClusterGetResponseDatabaseJSON contains the JSON metadata for
// the struct [CloudV1DbaaPostgreClusterGetResponseDatabase]
type cloudV1DbaaPostgreClusterGetResponseDatabaseJSON struct {
	Name        apijson.Field
	Owner       apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1DbaaPostgreClusterGetResponseDatabase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DbaaPostgreClusterGetResponseDatabaseJSON) RawJSON() string {
	return r.raw
}

type CloudV1DbaaPostgreClusterGetResponseNetwork struct {
	// Allowed IPs and subnets for incoming traffic
	ACL []string `json:"acl,required" format:"ipvanyinterface"`
	// Connection string to main database
	ConnectionString string `json:"connection_string,required"`
	// database hostname
	Host string `json:"host,required"`
	// Network Type
	NetworkType CloudV1DbaaPostgreClusterGetResponseNetworkNetworkType `json:"network_type,required"`
	JSON        cloudV1DbaaPostgreClusterGetResponseNetworkJSON        `json:"-"`
}

// cloudV1DbaaPostgreClusterGetResponseNetworkJSON contains the JSON metadata for
// the struct [CloudV1DbaaPostgreClusterGetResponseNetwork]
type cloudV1DbaaPostgreClusterGetResponseNetworkJSON struct {
	ACL              apijson.Field
	ConnectionString apijson.Field
	Host             apijson.Field
	NetworkType      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudV1DbaaPostgreClusterGetResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DbaaPostgreClusterGetResponseNetworkJSON) RawJSON() string {
	return r.raw
}

// Network Type
type CloudV1DbaaPostgreClusterGetResponseNetworkNetworkType string

const (
	CloudV1DbaaPostgreClusterGetResponseNetworkNetworkTypePublic CloudV1DbaaPostgreClusterGetResponseNetworkNetworkType = "public"
)

func (r CloudV1DbaaPostgreClusterGetResponseNetworkNetworkType) IsKnown() bool {
	switch r {
	case CloudV1DbaaPostgreClusterGetResponseNetworkNetworkTypePublic:
		return true
	}
	return false
}

type CloudV1DbaaPostgreClusterGetResponseUser struct {
	// Display was secret revealed or not
	IsSecretRevealed bool `json:"is_secret_revealed,required"`
	// User name
	Name string `json:"name,required"`
	// User's attributes
	RoleAttributes []CloudV1DbaaPostgreClusterGetResponseUsersRoleAttribute `json:"role_attributes,required"`
	JSON           cloudV1DbaaPostgreClusterGetResponseUserJSON             `json:"-"`
}

// cloudV1DbaaPostgreClusterGetResponseUserJSON contains the JSON metadata for the
// struct [CloudV1DbaaPostgreClusterGetResponseUser]
type cloudV1DbaaPostgreClusterGetResponseUserJSON struct {
	IsSecretRevealed apijson.Field
	Name             apijson.Field
	RoleAttributes   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CloudV1DbaaPostgreClusterGetResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DbaaPostgreClusterGetResponseUserJSON) RawJSON() string {
	return r.raw
}

type CloudV1DbaaPostgreClusterGetResponseUsersRoleAttribute string

const (
	CloudV1DbaaPostgreClusterGetResponseUsersRoleAttributeBypassrls  CloudV1DbaaPostgreClusterGetResponseUsersRoleAttribute = "BYPASSRLS"
	CloudV1DbaaPostgreClusterGetResponseUsersRoleAttributeCreatedb   CloudV1DbaaPostgreClusterGetResponseUsersRoleAttribute = "CREATEDB"
	CloudV1DbaaPostgreClusterGetResponseUsersRoleAttributeCreaterole CloudV1DbaaPostgreClusterGetResponseUsersRoleAttribute = "CREATEROLE"
	CloudV1DbaaPostgreClusterGetResponseUsersRoleAttributeInherit    CloudV1DbaaPostgreClusterGetResponseUsersRoleAttribute = "INHERIT"
	CloudV1DbaaPostgreClusterGetResponseUsersRoleAttributeLogin      CloudV1DbaaPostgreClusterGetResponseUsersRoleAttribute = "LOGIN"
	CloudV1DbaaPostgreClusterGetResponseUsersRoleAttributeNologin    CloudV1DbaaPostgreClusterGetResponseUsersRoleAttribute = "NOLOGIN"
)

func (r CloudV1DbaaPostgreClusterGetResponseUsersRoleAttribute) IsKnown() bool {
	switch r {
	case CloudV1DbaaPostgreClusterGetResponseUsersRoleAttributeBypassrls, CloudV1DbaaPostgreClusterGetResponseUsersRoleAttributeCreatedb, CloudV1DbaaPostgreClusterGetResponseUsersRoleAttributeCreaterole, CloudV1DbaaPostgreClusterGetResponseUsersRoleAttributeInherit, CloudV1DbaaPostgreClusterGetResponseUsersRoleAttributeLogin, CloudV1DbaaPostgreClusterGetResponseUsersRoleAttributeNologin:
		return true
	}
	return false
}

type CloudV1DbaaPostgreClusterListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []CloudV1DbaaPostgreClusterListResponseResult `json:"results,required"`
	JSON    cloudV1DbaaPostgreClusterListResponseJSON     `json:"-"`
}

// cloudV1DbaaPostgreClusterListResponseJSON contains the JSON metadata for the
// struct [CloudV1DbaaPostgreClusterListResponse]
type cloudV1DbaaPostgreClusterListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1DbaaPostgreClusterListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DbaaPostgreClusterListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1DbaaPostgreClusterListResponseResult struct {
	// PostgreSQL cluster name
	ClusterName string `json:"cluster_name,required"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Current cluster status
	Status ClusterStatus `json:"status,required"`
	// Cluster version
	Version string                                          `json:"version,required"`
	JSON    cloudV1DbaaPostgreClusterListResponseResultJSON `json:"-"`
}

// cloudV1DbaaPostgreClusterListResponseResultJSON contains the JSON metadata for
// the struct [CloudV1DbaaPostgreClusterListResponseResult]
type cloudV1DbaaPostgreClusterListResponseResultJSON struct {
	ClusterName apijson.Field
	CreatedAt   apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1DbaaPostgreClusterListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DbaaPostgreClusterListResponseResultJSON) RawJSON() string {
	return r.raw
}

type CloudV1DbaaPostgreClusterNewParams struct {
	PgClusterCreate PgClusterCreateParam `json:"pg_cluster_create,required"`
}

func (r CloudV1DbaaPostgreClusterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PgClusterCreate)
}

type CloudV1DbaaPostgreClusterUpdateParams struct {
	Databases param.Field[[]DatabaseParam] `json:"databases"`
	// New instance RAM and CPU
	Flavor param.Field[FlavorDatabaseParam] `json:"flavor"`
	// New High Availability settings
	HighAvailability param.Field[HighAvailabilityOptionsParam] `json:"high_availability"`
	Network          param.Field[PublicNetworkParam]           `json:"network"`
	// New PosgtreSQL cluster configuration
	PgServerConfiguration param.Field[CloudV1DbaaPostgreClusterUpdateParamsPgServerConfiguration] `json:"pg_server_configuration"`
	// New storage configuration
	Storage param.Field[CloudV1DbaaPostgreClusterUpdateParamsStorage] `json:"storage"`
	Users   param.Field[[]UserParam]                                  `json:"users"`
}

func (r CloudV1DbaaPostgreClusterUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// New PosgtreSQL cluster configuration
type CloudV1DbaaPostgreClusterUpdateParamsPgServerConfiguration struct {
	// New pg.conf file settings
	PgConf param.Field[string]                   `json:"pg_conf"`
	Pooler param.Field[PoolerConfigurationParam] `json:"pooler"`
	// New cluster version
	Version param.Field[string] `json:"version"`
}

func (r CloudV1DbaaPostgreClusterUpdateParamsPgServerConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// New storage configuration
type CloudV1DbaaPostgreClusterUpdateParamsStorage struct {
	// Total available storage for database
	SizeGib param.Field[int64] `json:"size_gib,required"`
}

func (r CloudV1DbaaPostgreClusterUpdateParamsStorage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1DbaaPostgreClusterListParams struct {
	// Limit the number of returned tasks. Falls back to default of 10 if not
	// specified.
	Limit param.Field[int64] `query:"limit"`
	// Offset value is used to exclude the first set of records from the result
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CloudV1DbaaPostgreClusterListParams]'s query parameters as
// `url.Values`.
func (r CloudV1DbaaPostgreClusterListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
