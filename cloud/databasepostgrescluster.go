// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/G-Core/gcore-go/shared/constant"
)

// DatabasePostgresClusterService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatabasePostgresClusterService] method instead.
type DatabasePostgresClusterService struct {
	Options         []option.RequestOption
	Metrics         DatabasePostgresClusterMetricService
	UserCredentials DatabasePostgresClusterUserCredentialService
	tasks           TaskService
}

// NewDatabasePostgresClusterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDatabasePostgresClusterService(opts ...option.RequestOption) (r DatabasePostgresClusterService) {
	r = DatabasePostgresClusterService{}
	r.Options = opts
	r.Metrics = NewDatabasePostgresClusterMetricService(opts...)
	r.UserCredentials = NewDatabasePostgresClusterUserCredentialService(opts...)
	r.tasks = NewTaskService(opts...)
	return
}

// Create a new PostgreSQL cluster with the specified configuration.
func (r *DatabasePostgresClusterService) New(ctx context.Context, params DatabasePostgresClusterNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/clusters/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// NewAndPoll creates a new PostgreSQL cluster and polls for completion
func (r *DatabasePostgresClusterService) NewAndPoll(ctx context.Context, params DatabasePostgresClusterNewParams, opts ...option.RequestOption) (v *PostgresCluster, err error) {
	resource, err := r.New(ctx, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams DatabasePostgresClusterGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]
	task, err := r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.PostgreSQLClusters) != 1 {
		return nil, errors.New("expected exactly one postgres cluster to be created in a task")
	}
	resourceID := task.CreatedResources.PostgreSQLClusters[0]

	return r.Get(ctx, resourceID, getParams, opts...)
}

// Update the configuration of an existing PostgreSQL cluster.
func (r *DatabasePostgresClusterService) Update(ctx context.Context, clusterName string, params DatabasePostgresClusterUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/clusters/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// UpdateAndPoll updates a PostgreSQL cluster and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *DatabasePostgresClusterService) UpdateAndPoll(ctx context.Context, clusterName string, params DatabasePostgresClusterUpdateParams, opts ...option.RequestOption) (v *PostgresCluster, err error) {
	resource, err := r.Update(ctx, clusterName, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams DatabasePostgresClusterGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) == 0 {
		return nil, errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, clusterName, getParams, opts...)
}

// List all PostgreSQL clusters in the specified project and region. Results can be
// filtered by search query and paginated.
func (r *DatabasePostgresClusterService) List(ctx context.Context, params DatabasePostgresClusterListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[PostgresClusterShort], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/clusters/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all PostgreSQL clusters in the specified project and region. Results can be
// filtered by search query and paginated.
func (r *DatabasePostgresClusterService) ListAutoPaging(ctx context.Context, params DatabasePostgresClusterListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[PostgresClusterShort] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete a PostgreSQL cluster and all its associated resources.
func (r *DatabasePostgresClusterService) Delete(ctx context.Context, clusterName string, body DatabasePostgresClusterDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/clusters/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// DeleteAndPoll deletes a PostgreSQL cluster and polls for completion of the first task. Use the [TaskService.Poll] method if you
// need to poll for all tasks.
func (r *DatabasePostgresClusterService) DeleteAndPoll(ctx context.Context, clusterName string, body DatabasePostgresClusterDeleteParams, opts ...option.RequestOption) error {
	resource, err := r.Delete(ctx, clusterName, body, opts...)
	if err != nil {
		return err
	}

	opts = slices.Concat(r.Options, opts)
	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]
	_, err = r.tasks.Poll(ctx, taskID, opts...)
	return err
}

// Get detailed information about a specific PostgreSQL cluster.
func (r *DatabasePostgresClusterService) Get(ctx context.Context, clusterName string, query DatabasePostgresClusterGetParams, opts ...option.RequestOption) (res *PostgresCluster, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if clusterName == "" {
		err = errors.New("missing required cluster_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/clusters/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, clusterName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PostgresCluster struct {
	ClusterName string                    `json:"cluster_name,required"`
	CreatedAt   time.Time                 `json:"created_at,required" format:"date-time"`
	Databases   []PostgresClusterDatabase `json:"databases,required"`
	// Instance RAM and CPU
	Flavor           PostgresClusterFlavor           `json:"flavor,required"`
	HighAvailability PostgresClusterHighAvailability `json:"high_availability,required"`
	Network          PostgresClusterNetwork          `json:"network,required"`
	// Main PG configuration
	PgServerConfiguration PostgresClusterPgServerConfiguration `json:"pg_server_configuration,required"`
	// Current cluster status
	//
	// Any of "DELETING", "FAILED", "PREPARING", "READY", "UNHEALTHY", "UNKNOWN",
	// "UPDATING".
	Status PostgresClusterStatus `json:"status,required"`
	// PG's storage configuration
	Storage PostgresClusterStorage `json:"storage,required"`
	Users   []PostgresClusterUser  `json:"users,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClusterName           respjson.Field
		CreatedAt             respjson.Field
		Databases             respjson.Field
		Flavor                respjson.Field
		HighAvailability      respjson.Field
		Network               respjson.Field
		PgServerConfiguration respjson.Field
		Status                respjson.Field
		Storage               respjson.Field
		Users                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostgresCluster) RawJSON() string { return r.JSON.raw }
func (r *PostgresCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostgresClusterDatabase struct {
	// Database name
	Name string `json:"name,required"`
	// Database owner from users list
	Owner string `json:"owner,required"`
	// Size in bytes
	Size int64 `json:"size,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Owner       respjson.Field
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostgresClusterDatabase) RawJSON() string { return r.JSON.raw }
func (r *PostgresClusterDatabase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance RAM and CPU
type PostgresClusterFlavor struct {
	// Maximum available cores for instance
	CPU int64 `json:"cpu,required"`
	// Maximum available RAM for instance
	MemoryGib int64 `json:"memory_gib,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPU         respjson.Field
		MemoryGib   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostgresClusterFlavor) RawJSON() string { return r.JSON.raw }
func (r *PostgresClusterFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostgresClusterHighAvailability struct {
	// Type of replication
	//
	// Any of "async", "sync".
	ReplicationMode string `json:"replication_mode,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReplicationMode respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostgresClusterHighAvailability) RawJSON() string { return r.JSON.raw }
func (r *PostgresClusterHighAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostgresClusterNetwork struct {
	// Allowed IPs and subnets for incoming traffic
	ACL []string `json:"acl,required" format:"ipvanyinterface"`
	// Connection string to main database
	ConnectionString string `json:"connection_string,required"`
	// database hostname
	Host string `json:"host,required"`
	// Network Type
	NetworkType constant.Public `json:"network_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ACL              respjson.Field
		ConnectionString respjson.Field
		Host             respjson.Field
		NetworkType      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostgresClusterNetwork) RawJSON() string { return r.JSON.raw }
func (r *PostgresClusterNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Main PG configuration
type PostgresClusterPgServerConfiguration struct {
	// pg.conf settings
	PgConf string `json:"pg_conf,required"`
	// Cluster version
	Version string                                     `json:"version,required"`
	Pooler  PostgresClusterPgServerConfigurationPooler `json:"pooler,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PgConf      respjson.Field
		Version     respjson.Field
		Pooler      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostgresClusterPgServerConfiguration) RawJSON() string { return r.JSON.raw }
func (r *PostgresClusterPgServerConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostgresClusterPgServerConfigurationPooler struct {
	// Any of "session", "statement", "transaction".
	Mode string `json:"mode,required"`
	// Any of "pgbouncer".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostgresClusterPgServerConfigurationPooler) RawJSON() string { return r.JSON.raw }
func (r *PostgresClusterPgServerConfigurationPooler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current cluster status
type PostgresClusterStatus string

const (
	PostgresClusterStatusDeleting  PostgresClusterStatus = "DELETING"
	PostgresClusterStatusFailed    PostgresClusterStatus = "FAILED"
	PostgresClusterStatusPreparing PostgresClusterStatus = "PREPARING"
	PostgresClusterStatusReady     PostgresClusterStatus = "READY"
	PostgresClusterStatusUnhealthy PostgresClusterStatus = "UNHEALTHY"
	PostgresClusterStatusUnknown   PostgresClusterStatus = "UNKNOWN"
	PostgresClusterStatusUpdating  PostgresClusterStatus = "UPDATING"
)

// PG's storage configuration
type PostgresClusterStorage struct {
	// Total available storage for database
	SizeGib int64 `json:"size_gib,required"`
	// Storage type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SizeGib     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostgresClusterStorage) RawJSON() string { return r.JSON.raw }
func (r *PostgresClusterStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostgresClusterUser struct {
	// Display was secret revealed or not
	IsSecretRevealed bool `json:"is_secret_revealed,required"`
	// User name
	Name string `json:"name,required"`
	// User's attributes
	//
	// Any of "BYPASSRLS", "CREATEDB", "CREATEROLE", "INHERIT", "LOGIN", "NOLOGIN".
	RoleAttributes []string `json:"role_attributes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSecretRevealed respjson.Field
		Name             respjson.Field
		RoleAttributes   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostgresClusterUser) RawJSON() string { return r.JSON.raw }
func (r *PostgresClusterUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostgresClusterShort struct {
	// PostgreSQL cluster name
	ClusterName string `json:"cluster_name,required"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Current cluster status
	//
	// Any of "DELETING", "FAILED", "PREPARING", "READY", "UNHEALTHY", "UNKNOWN",
	// "UPDATING".
	Status PostgresClusterShortStatus `json:"status,required"`
	// Cluster version
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClusterName respjson.Field
		CreatedAt   respjson.Field
		Status      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostgresClusterShort) RawJSON() string { return r.JSON.raw }
func (r *PostgresClusterShort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current cluster status
type PostgresClusterShortStatus string

const (
	PostgresClusterShortStatusDeleting  PostgresClusterShortStatus = "DELETING"
	PostgresClusterShortStatusFailed    PostgresClusterShortStatus = "FAILED"
	PostgresClusterShortStatusPreparing PostgresClusterShortStatus = "PREPARING"
	PostgresClusterShortStatusReady     PostgresClusterShortStatus = "READY"
	PostgresClusterShortStatusUnhealthy PostgresClusterShortStatus = "UNHEALTHY"
	PostgresClusterShortStatusUnknown   PostgresClusterShortStatus = "UNKNOWN"
	PostgresClusterShortStatusUpdating  PostgresClusterShortStatus = "UPDATING"
)

type DatabasePostgresClusterNewParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// High Availability settings
	HighAvailability DatabasePostgresClusterNewParamsHighAvailability `json:"high_availability,omitzero,required"`
	// PostgreSQL cluster name
	ClusterName string `json:"cluster_name,required"`
	// Instance RAM and CPU
	Flavor  DatabasePostgresClusterNewParamsFlavor  `json:"flavor,omitzero,required"`
	Network DatabasePostgresClusterNewParamsNetwork `json:"network,omitzero,required"`
	// PosgtreSQL cluster configuration
	PgServerConfiguration DatabasePostgresClusterNewParamsPgServerConfiguration `json:"pg_server_configuration,omitzero,required"`
	// Cluster's storage configuration
	Storage   DatabasePostgresClusterNewParamsStorage    `json:"storage,omitzero,required"`
	Databases []DatabasePostgresClusterNewParamsDatabase `json:"databases,omitzero"`
	Users     []DatabasePostgresClusterNewParamsUser     `json:"users,omitzero"`
	paramObj
}

func (r DatabasePostgresClusterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instance RAM and CPU
//
// The properties CPU, MemoryGib are required.
type DatabasePostgresClusterNewParamsFlavor struct {
	// Maximum available cores for instance
	CPU int64 `json:"cpu,required"`
	// Maximum available RAM for instance
	MemoryGib int64 `json:"memory_gib,required"`
	paramObj
}

func (r DatabasePostgresClusterNewParamsFlavor) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterNewParamsFlavor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterNewParamsFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// High Availability settings
//
// The property ReplicationMode is required.
type DatabasePostgresClusterNewParamsHighAvailability struct {
	// Type of replication
	//
	// Any of "async", "sync".
	ReplicationMode string `json:"replication_mode,omitzero,required"`
	paramObj
}

func (r DatabasePostgresClusterNewParamsHighAvailability) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterNewParamsHighAvailability
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterNewParamsHighAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DatabasePostgresClusterNewParamsHighAvailability](
		"replication_mode", "async", "sync",
	)
}

// The properties ACL, NetworkType are required.
type DatabasePostgresClusterNewParamsNetwork struct {
	// Allowed IPs and subnets for incoming traffic
	ACL []string `json:"acl,omitzero,required" format:"ipvanyinterface"`
	// Network Type
	//
	// This field can be elided, and will marshal its zero value as "public".
	NetworkType constant.Public `json:"network_type,required"`
	paramObj
}

func (r DatabasePostgresClusterNewParamsNetwork) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterNewParamsNetwork
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterNewParamsNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PosgtreSQL cluster configuration
//
// The properties PgConf, Version are required.
type DatabasePostgresClusterNewParamsPgServerConfiguration struct {
	// pg.conf settings
	PgConf string `json:"pg_conf,required"`
	// Cluster version
	Version string                                                      `json:"version,required"`
	Pooler  DatabasePostgresClusterNewParamsPgServerConfigurationPooler `json:"pooler,omitzero"`
	paramObj
}

func (r DatabasePostgresClusterNewParamsPgServerConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterNewParamsPgServerConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterNewParamsPgServerConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Mode is required.
type DatabasePostgresClusterNewParamsPgServerConfigurationPooler struct {
	// Any of "session", "statement", "transaction".
	Mode string `json:"mode,omitzero,required"`
	// Any of "pgbouncer".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r DatabasePostgresClusterNewParamsPgServerConfigurationPooler) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterNewParamsPgServerConfigurationPooler
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterNewParamsPgServerConfigurationPooler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DatabasePostgresClusterNewParamsPgServerConfigurationPooler](
		"mode", "session", "statement", "transaction",
	)
	apijson.RegisterFieldValidator[DatabasePostgresClusterNewParamsPgServerConfigurationPooler](
		"type", "pgbouncer",
	)
}

// Cluster's storage configuration
//
// The properties SizeGib, Type are required.
type DatabasePostgresClusterNewParamsStorage struct {
	// Total available storage for database
	SizeGib int64 `json:"size_gib,required"`
	// Storage type
	Type string `json:"type,required"`
	paramObj
}

func (r DatabasePostgresClusterNewParamsStorage) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterNewParamsStorage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterNewParamsStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Owner are required.
type DatabasePostgresClusterNewParamsDatabase struct {
	// Database name
	Name string `json:"name,required"`
	// Database owner from users list
	Owner string `json:"owner,required"`
	paramObj
}

func (r DatabasePostgresClusterNewParamsDatabase) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterNewParamsDatabase
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterNewParamsDatabase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, RoleAttributes are required.
type DatabasePostgresClusterNewParamsUser struct {
	// User name
	Name string `json:"name,required"`
	// User's attributes
	//
	// Any of "BYPASSRLS", "CREATEDB", "CREATEROLE", "INHERIT", "LOGIN", "NOLOGIN".
	RoleAttributes []string `json:"role_attributes,omitzero,required"`
	paramObj
}

func (r DatabasePostgresClusterNewParamsUser) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterNewParamsUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterNewParamsUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatabasePostgresClusterUpdateParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// New instance RAM and CPU
	Flavor DatabasePostgresClusterUpdateParamsFlavor `json:"flavor,omitzero"`
	// New High Availability settings
	HighAvailability DatabasePostgresClusterUpdateParamsHighAvailability `json:"high_availability,omitzero"`
	Network          DatabasePostgresClusterUpdateParamsNetwork          `json:"network,omitzero"`
	// New PosgtreSQL cluster configuration
	PgServerConfiguration DatabasePostgresClusterUpdateParamsPgServerConfiguration `json:"pg_server_configuration,omitzero"`
	// New storage configuration
	Storage   DatabasePostgresClusterUpdateParamsStorage    `json:"storage,omitzero"`
	Databases []DatabasePostgresClusterUpdateParamsDatabase `json:"databases,omitzero"`
	Users     []DatabasePostgresClusterUpdateParamsUser     `json:"users,omitzero"`
	paramObj
}

func (r DatabasePostgresClusterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Owner are required.
type DatabasePostgresClusterUpdateParamsDatabase struct {
	// Database name
	Name string `json:"name,required"`
	// Database owner from users list
	Owner string `json:"owner,required"`
	paramObj
}

func (r DatabasePostgresClusterUpdateParamsDatabase) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterUpdateParamsDatabase
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterUpdateParamsDatabase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// New instance RAM and CPU
//
// The properties CPU, MemoryGib are required.
type DatabasePostgresClusterUpdateParamsFlavor struct {
	// Maximum available cores for instance
	CPU int64 `json:"cpu,required"`
	// Maximum available RAM for instance
	MemoryGib int64 `json:"memory_gib,required"`
	paramObj
}

func (r DatabasePostgresClusterUpdateParamsFlavor) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterUpdateParamsFlavor
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterUpdateParamsFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// New High Availability settings
//
// The property ReplicationMode is required.
type DatabasePostgresClusterUpdateParamsHighAvailability struct {
	// Type of replication
	//
	// Any of "async", "sync".
	ReplicationMode string `json:"replication_mode,omitzero,required"`
	paramObj
}

func (r DatabasePostgresClusterUpdateParamsHighAvailability) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterUpdateParamsHighAvailability
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterUpdateParamsHighAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DatabasePostgresClusterUpdateParamsHighAvailability](
		"replication_mode", "async", "sync",
	)
}

// The properties ACL, NetworkType are required.
type DatabasePostgresClusterUpdateParamsNetwork struct {
	// Allowed IPs and subnets for incoming traffic
	ACL []string `json:"acl,omitzero,required" format:"ipvanyinterface"`
	// Network Type
	//
	// This field can be elided, and will marshal its zero value as "public".
	NetworkType constant.Public `json:"network_type,required"`
	paramObj
}

func (r DatabasePostgresClusterUpdateParamsNetwork) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterUpdateParamsNetwork
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterUpdateParamsNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// New PosgtreSQL cluster configuration
type DatabasePostgresClusterUpdateParamsPgServerConfiguration struct {
	// New pg.conf file settings
	PgConf param.Opt[string] `json:"pg_conf,omitzero"`
	// New cluster version
	Version param.Opt[string]                                              `json:"version,omitzero"`
	Pooler  DatabasePostgresClusterUpdateParamsPgServerConfigurationPooler `json:"pooler,omitzero"`
	paramObj
}

func (r DatabasePostgresClusterUpdateParamsPgServerConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterUpdateParamsPgServerConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterUpdateParamsPgServerConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Mode is required.
type DatabasePostgresClusterUpdateParamsPgServerConfigurationPooler struct {
	// Any of "session", "statement", "transaction".
	Mode string `json:"mode,omitzero,required"`
	// Any of "pgbouncer".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r DatabasePostgresClusterUpdateParamsPgServerConfigurationPooler) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterUpdateParamsPgServerConfigurationPooler
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterUpdateParamsPgServerConfigurationPooler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DatabasePostgresClusterUpdateParamsPgServerConfigurationPooler](
		"mode", "session", "statement", "transaction",
	)
	apijson.RegisterFieldValidator[DatabasePostgresClusterUpdateParamsPgServerConfigurationPooler](
		"type", "pgbouncer",
	)
}

// New storage configuration
//
// The property SizeGib is required.
type DatabasePostgresClusterUpdateParamsStorage struct {
	// Total available storage for database
	SizeGib int64 `json:"size_gib,required"`
	paramObj
}

func (r DatabasePostgresClusterUpdateParamsStorage) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterUpdateParamsStorage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterUpdateParamsStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, RoleAttributes are required.
type DatabasePostgresClusterUpdateParamsUser struct {
	// User name
	Name string `json:"name,required"`
	// User's attributes
	//
	// Any of "BYPASSRLS", "CREATEDB", "CREATEROLE", "INHERIT", "LOGIN", "NOLOGIN".
	RoleAttributes []string `json:"role_attributes,omitzero,required"`
	paramObj
}

func (r DatabasePostgresClusterUpdateParamsUser) MarshalJSON() (data []byte, err error) {
	type shadow DatabasePostgresClusterUpdateParamsUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DatabasePostgresClusterUpdateParamsUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DatabasePostgresClusterListParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Maximum number of clusters to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of clusters to skip
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DatabasePostgresClusterListParams]'s query parameters as
// `url.Values`.
func (r DatabasePostgresClusterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DatabasePostgresClusterDeleteParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type DatabasePostgresClusterGetParams struct {
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID  param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}
