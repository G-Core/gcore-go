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

// CloudV1DbaaPostgreService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1DbaaPostgreService] method instead.
type CloudV1DbaaPostgreService struct {
	Options  []option.RequestOption
	Clusters *CloudV1DbaaPostgreClusterService
}

// NewCloudV1DbaaPostgreService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1DbaaPostgreService(opts ...option.RequestOption) (r *CloudV1DbaaPostgreService) {
	r = &CloudV1DbaaPostgreService{}
	r.Options = opts
	r.Clusters = NewCloudV1DbaaPostgreClusterService(opts...)
	return
}

// Check quota PostgreSQL cluster creation
func (r *CloudV1DbaaPostgreService) CheckQuota(ctx context.Context, projectID int64, regionID int64, body CloudV1DbaaPostgreCheckQuotaParams, opts ...option.RequestOption) (res *RegionalDiffQuotas, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/check_limit/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get available PostgtreSQL configurations for a region.
func (r *CloudV1DbaaPostgreService) GetConfigurations(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *CloudV1DbaaPostgreGetConfigurationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/configuration/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Validate user's custom PostgreSQL configuration file.
func (r *CloudV1DbaaPostgreService) ValidateConfiguration(ctx context.Context, projectID int64, regionID int64, body CloudV1DbaaPostgreValidateConfigurationParams, opts ...option.RequestOption) (res *CloudV1DbaaPostgreValidateConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/validate_pg_conf/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DatabaseParam struct {
	// Database name
	Name param.Field[string] `json:"name,required"`
	// Database owner from users list
	Owner param.Field[string] `json:"owner,required"`
}

func (r DatabaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FlavorDatabase struct {
	// Maximum available cores for instance
	CPU int64 `json:"cpu,required"`
	// Maximum available RAM for instance
	MemoryGib int64              `json:"memory_gib,required"`
	JSON      flavorDatabaseJSON `json:"-"`
}

// flavorDatabaseJSON contains the JSON metadata for the struct [FlavorDatabase]
type flavorDatabaseJSON struct {
	CPU         apijson.Field
	MemoryGib   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FlavorDatabase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r flavorDatabaseJSON) RawJSON() string {
	return r.raw
}

type FlavorDatabaseParam struct {
	// Maximum available cores for instance
	CPU param.Field[int64] `json:"cpu,required"`
	// Maximum available RAM for instance
	MemoryGib param.Field[int64] `json:"memory_gib,required"`
}

func (r FlavorDatabaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HighAvailabilityOptions struct {
	// Type of replication
	ReplicationMode HighAvailabilityOptionsReplicationMode `json:"replication_mode,required"`
	JSON            highAvailabilityOptionsJSON            `json:"-"`
}

// highAvailabilityOptionsJSON contains the JSON metadata for the struct
// [HighAvailabilityOptions]
type highAvailabilityOptionsJSON struct {
	ReplicationMode apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *HighAvailabilityOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r highAvailabilityOptionsJSON) RawJSON() string {
	return r.raw
}

// Type of replication
type HighAvailabilityOptionsReplicationMode string

const (
	HighAvailabilityOptionsReplicationModeAsync HighAvailabilityOptionsReplicationMode = "async"
	HighAvailabilityOptionsReplicationModeSync  HighAvailabilityOptionsReplicationMode = "sync"
)

func (r HighAvailabilityOptionsReplicationMode) IsKnown() bool {
	switch r {
	case HighAvailabilityOptionsReplicationModeAsync, HighAvailabilityOptionsReplicationModeSync:
		return true
	}
	return false
}

type HighAvailabilityOptionsParam struct {
	// Type of replication
	ReplicationMode param.Field[HighAvailabilityOptionsReplicationMode] `json:"replication_mode,required"`
}

func (r HighAvailabilityOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PgClusterCreateParam struct {
	// PostgreSQL cluster name
	ClusterName param.Field[string] `json:"cluster_name,required"`
	// Instance RAM and CPU
	Flavor param.Field[FlavorDatabaseParam] `json:"flavor,required"`
	// High Availability settings
	HighAvailability param.Field[HighAvailabilityOptionsParam] `json:"high_availability,required"`
	Network          param.Field[PublicNetworkParam]           `json:"network,required"`
	// PosgtreSQL cluster configuration
	PgServerConfiguration param.Field[PostgresqlServerConfigParam] `json:"pg_server_configuration,required"`
	// Cluster's storage configuration
	Storage   param.Field[StorageParam]    `json:"storage,required"`
	Databases param.Field[[]DatabaseParam] `json:"databases"`
	Users     param.Field[[]UserParam]     `json:"users"`
}

func (r PgClusterCreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PoolerConfiguration struct {
	Mode PoolerConfigurationMode `json:"mode,required"`
	Type PoolerConfigurationType `json:"type"`
	JSON poolerConfigurationJSON `json:"-"`
}

// poolerConfigurationJSON contains the JSON metadata for the struct
// [PoolerConfiguration]
type poolerConfigurationJSON struct {
	Mode        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolerConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolerConfigurationJSON) RawJSON() string {
	return r.raw
}

type PoolerConfigurationMode string

const (
	PoolerConfigurationModeSession     PoolerConfigurationMode = "session"
	PoolerConfigurationModeStatement   PoolerConfigurationMode = "statement"
	PoolerConfigurationModeTransaction PoolerConfigurationMode = "transaction"
)

func (r PoolerConfigurationMode) IsKnown() bool {
	switch r {
	case PoolerConfigurationModeSession, PoolerConfigurationModeStatement, PoolerConfigurationModeTransaction:
		return true
	}
	return false
}

type PoolerConfigurationType string

const (
	PoolerConfigurationTypePgbouncer PoolerConfigurationType = "pgbouncer"
)

func (r PoolerConfigurationType) IsKnown() bool {
	switch r {
	case PoolerConfigurationTypePgbouncer:
		return true
	}
	return false
}

type PoolerConfigurationParam struct {
	Mode param.Field[PoolerConfigurationMode] `json:"mode,required"`
	Type param.Field[PoolerConfigurationType] `json:"type"`
}

func (r PoolerConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PostgresqlServerConfig struct {
	// pg.conf settings
	PgConf string `json:"pg_conf,required"`
	// Cluster version
	Version string                     `json:"version,required"`
	Pooler  PoolerConfiguration        `json:"pooler,nullable"`
	JSON    postgresqlServerConfigJSON `json:"-"`
}

// postgresqlServerConfigJSON contains the JSON metadata for the struct
// [PostgresqlServerConfig]
type postgresqlServerConfigJSON struct {
	PgConf      apijson.Field
	Version     apijson.Field
	Pooler      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostgresqlServerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postgresqlServerConfigJSON) RawJSON() string {
	return r.raw
}

type PostgresqlServerConfigParam struct {
	// pg.conf settings
	PgConf param.Field[string] `json:"pg_conf,required"`
	// Cluster version
	Version param.Field[string]                   `json:"version,required"`
	Pooler  param.Field[PoolerConfigurationParam] `json:"pooler"`
}

func (r PostgresqlServerConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PublicNetworkParam struct {
	// Allowed IPs and subnets for incoming traffic
	ACL param.Field[[]string] `json:"acl,required" format:"ipvanyinterface"`
	// Network Type
	NetworkType param.Field[PublicNetworkNetworkType] `json:"network_type,required"`
}

func (r PublicNetworkParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Network Type
type PublicNetworkNetworkType string

const (
	PublicNetworkNetworkTypePublic PublicNetworkNetworkType = "public"
)

func (r PublicNetworkNetworkType) IsKnown() bool {
	switch r {
	case PublicNetworkNetworkTypePublic:
		return true
	}
	return false
}

type Storage struct {
	// Total available storage for database
	SizeGib int64 `json:"size_gib,required"`
	// Storage type
	Type string      `json:"type,required"`
	JSON storageJSON `json:"-"`
}

// storageJSON contains the JSON metadata for the struct [Storage]
type storageJSON struct {
	SizeGib     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Storage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storageJSON) RawJSON() string {
	return r.raw
}

type StorageParam struct {
	// Total available storage for database
	SizeGib param.Field[int64] `json:"size_gib,required"`
	// Storage type
	Type param.Field[string] `json:"type,required"`
}

func (r StorageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserParam struct {
	// User name
	Name param.Field[string] `json:"name,required"`
	// User's attributes
	RoleAttributes param.Field[[]UserRoleAttribute] `json:"role_attributes,required"`
}

func (r UserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserRoleAttribute string

const (
	UserRoleAttributeBypassrls  UserRoleAttribute = "BYPASSRLS"
	UserRoleAttributeCreatedb   UserRoleAttribute = "CREATEDB"
	UserRoleAttributeCreaterole UserRoleAttribute = "CREATEROLE"
	UserRoleAttributeInherit    UserRoleAttribute = "INHERIT"
	UserRoleAttributeLogin      UserRoleAttribute = "LOGIN"
	UserRoleAttributeNologin    UserRoleAttribute = "NOLOGIN"
)

func (r UserRoleAttribute) IsKnown() bool {
	switch r {
	case UserRoleAttributeBypassrls, UserRoleAttributeCreatedb, UserRoleAttributeCreaterole, UserRoleAttributeInherit, UserRoleAttributeLogin, UserRoleAttributeNologin:
		return true
	}
	return false
}

type CloudV1DbaaPostgreGetConfigurationsResponse struct {
	Flavors        []CloudV1DbaaPostgreGetConfigurationsResponseFlavor       `json:"flavors,required"`
	StorageClasses []CloudV1DbaaPostgreGetConfigurationsResponseStorageClass `json:"storage_classes,required"`
	// Available versions
	Versions []string                                        `json:"versions,required"`
	JSON     cloudV1DbaaPostgreGetConfigurationsResponseJSON `json:"-"`
}

// cloudV1DbaaPostgreGetConfigurationsResponseJSON contains the JSON metadata for
// the struct [CloudV1DbaaPostgreGetConfigurationsResponse]
type cloudV1DbaaPostgreGetConfigurationsResponseJSON struct {
	Flavors        apijson.Field
	StorageClasses apijson.Field
	Versions       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CloudV1DbaaPostgreGetConfigurationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DbaaPostgreGetConfigurationsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1DbaaPostgreGetConfigurationsResponseFlavor struct {
	// Maximum available cores for instance
	CPU int64 `json:"cpu,required"`
	// Maximum available RAM for instance
	MemoryGib int64                                                 `json:"memory_gib,required"`
	PgConf    string                                                `json:"pg_conf,required"`
	JSON      cloudV1DbaaPostgreGetConfigurationsResponseFlavorJSON `json:"-"`
}

// cloudV1DbaaPostgreGetConfigurationsResponseFlavorJSON contains the JSON metadata
// for the struct [CloudV1DbaaPostgreGetConfigurationsResponseFlavor]
type cloudV1DbaaPostgreGetConfigurationsResponseFlavorJSON struct {
	CPU         apijson.Field
	MemoryGib   apijson.Field
	PgConf      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1DbaaPostgreGetConfigurationsResponseFlavor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DbaaPostgreGetConfigurationsResponseFlavorJSON) RawJSON() string {
	return r.raw
}

type CloudV1DbaaPostgreGetConfigurationsResponseStorageClass struct {
	// Storage type
	Name string                                                      `json:"name,required"`
	JSON cloudV1DbaaPostgreGetConfigurationsResponseStorageClassJSON `json:"-"`
}

// cloudV1DbaaPostgreGetConfigurationsResponseStorageClassJSON contains the JSON
// metadata for the struct
// [CloudV1DbaaPostgreGetConfigurationsResponseStorageClass]
type cloudV1DbaaPostgreGetConfigurationsResponseStorageClassJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1DbaaPostgreGetConfigurationsResponseStorageClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DbaaPostgreGetConfigurationsResponseStorageClassJSON) RawJSON() string {
	return r.raw
}

type CloudV1DbaaPostgreValidateConfigurationResponse struct {
	// Errors list
	Errors []string `json:"errors,required"`
	// Validity of pg.conf file
	IsValid bool                                                `json:"is_valid,required"`
	JSON    cloudV1DbaaPostgreValidateConfigurationResponseJSON `json:"-"`
}

// cloudV1DbaaPostgreValidateConfigurationResponseJSON contains the JSON metadata
// for the struct [CloudV1DbaaPostgreValidateConfigurationResponse]
type cloudV1DbaaPostgreValidateConfigurationResponseJSON struct {
	Errors      apijson.Field
	IsValid     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1DbaaPostgreValidateConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1DbaaPostgreValidateConfigurationResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1DbaaPostgreCheckQuotaParams struct {
	PgClusterCreate PgClusterCreateParam `json:"pg_cluster_create,required"`
}

func (r CloudV1DbaaPostgreCheckQuotaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PgClusterCreate)
}

type CloudV1DbaaPostgreValidateConfigurationParams struct {
	// PostgreSQL configuration
	PgConf param.Field[string] `json:"pg_conf,required"`
	// PostgreSQL version
	Version param.Field[string] `json:"version,required"`
}

func (r CloudV1DbaaPostgreValidateConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
