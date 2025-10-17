// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
	"github.com/G-Core/gcore-go/shared/constant"
)

// FileShareService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileShareService] method instead.
type FileShareService struct {
	Options     []option.RequestOption
	AccessRules FileShareAccessRuleService
}

// NewFileShareService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFileShareService(opts ...option.RequestOption) (r FileShareService) {
	r = FileShareService{}
	r.Options = opts
	r.AccessRules = NewFileShareAccessRuleService(opts...)
	return
}

// Create file share
func (r *FileShareService) New(ctx context.Context, params FileShareNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Rename file share, update tags or set share specific properties
func (r *FileShareService) Update(ctx context.Context, fileShareID string, params FileShareUpdateParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v3/file_shares/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List file shares
func (r *FileShareService) List(ctx context.Context, params FileShareListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[FileShare], err error) {
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
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v", params.ProjectID.Value, params.RegionID.Value)
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

// List file shares
func (r *FileShareService) ListAutoPaging(ctx context.Context, params FileShareListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[FileShare] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete file share
func (r *FileShareService) Delete(ctx context.Context, fileShareID string, body FileShareDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s", body.ProjectID.Value, body.RegionID.Value, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get file share
func (r *FileShareService) Get(ctx context.Context, fileShareID string, query FileShareGetParams, opts ...option.RequestOption) (res *FileShare, err error) {
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
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Resize file share
func (r *FileShareService) Resize(ctx context.Context, fileShareID string, params FileShareResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s/extend", params.ProjectID.Value, params.RegionID.Value, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type FileShare struct {
	// File share ID
	ID string `json:"id,required" format:"uuid4"`
	// Connection point. Can be null during File share creation
	ConnectionPoint string `json:"connection_point,required"`
	// Datetime when the file share was created
	CreatedAt string `json:"created_at,required"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required" format:"uuid4"`
	// File share name
	Name string `json:"name,required"`
	// Network ID.
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// Network name.
	NetworkName string `json:"network_name,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// File share protocol
	Protocol string `json:"protocol,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Share network name. May be null if the file share was created with volume type
	// VAST
	ShareNetworkName string `json:"share_network_name,required"`
	// Share settings specific to the file share type
	ShareSettings FileShareShareSettingsUnion `json:"share_settings,required"`
	// File share size in GiB
	Size int64 `json:"size,required"`
	// File share status
	//
	// Any of "available", "awaiting_transfer", "backup_creating", "backup_restoring",
	// "backup_restoring_error", "creating", "creating_from_snapshot", "deleted",
	// "deleting", "ensuring", "error", "error_deleting", "extending",
	// "extending_error", "inactive", "manage_error", "manage_starting", "migrating",
	// "migrating_to", "replication_change", "reverting", "reverting_error",
	// "shrinking", "shrinking_error", "shrinking_possible_data_loss_error",
	// "unmanage_error", "unmanage_starting", "unmanaged".
	Status FileShareStatus `json:"status,required"`
	// Subnet ID.
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Subnet name.
	SubnetName string `json:"subnet_name,required"`
	// List of key-value tags associated with the resource. A tag is a key-value pair
	// that can be associated with a resource, enabling efficient filtering and
	// grouping for better organization and management. Some tags are read-only and
	// cannot be modified by the user. Tags are also integrated with cost reports,
	// allowing cost data to be filtered based on tag keys or values.
	Tags []Tag `json:"tags,required"`
	// The UUID of the active task that currently holds a lock on the resource. This
	// lock prevents concurrent modifications to ensure consistency. If `null`, the
	// resource is not locked.
	TaskID string `json:"task_id,required" format:"uuid4"`
	// File share type name
	//
	// Any of "standard", "vast".
	TypeName FileShareTypeName `json:"type_name,required"`
	// Deprecated. Use `type_name` instead. File share disk type
	//
	// Any of "default_share_type", "vast_share_type".
	//
	// Deprecated: deprecated
	VolumeType FileShareVolumeType `json:"volume_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ConnectionPoint  respjson.Field
		CreatedAt        respjson.Field
		CreatorTaskID    respjson.Field
		Name             respjson.Field
		NetworkID        respjson.Field
		NetworkName      respjson.Field
		ProjectID        respjson.Field
		Protocol         respjson.Field
		Region           respjson.Field
		RegionID         respjson.Field
		ShareNetworkName respjson.Field
		ShareSettings    respjson.Field
		Size             respjson.Field
		Status           respjson.Field
		SubnetID         respjson.Field
		SubnetName       respjson.Field
		Tags             respjson.Field
		TaskID           respjson.Field
		TypeName         respjson.Field
		VolumeType       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileShare) RawJSON() string { return r.JSON.raw }
func (r *FileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FileShareShareSettingsUnion contains all possible properties and values from
// [FileShareShareSettingsStandard], [FileShareShareSettingsVast].
//
// Use the [FileShareShareSettingsUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FileShareShareSettingsUnion struct {
	// Any of "standard", "vast".
	TypeName string `json:"type_name"`
	// This field is from variant [FileShareShareSettingsVast].
	AllowedCharacters string `json:"allowed_characters"`
	// This field is from variant [FileShareShareSettingsVast].
	PathLength string `json:"path_length"`
	// This field is from variant [FileShareShareSettingsVast].
	RootSquash bool `json:"root_squash"`
	JSON       struct {
		TypeName          respjson.Field
		AllowedCharacters respjson.Field
		PathLength        respjson.Field
		RootSquash        respjson.Field
		raw               string
	} `json:"-"`
}

// anyFileShareShareSettings is implemented by each variant of
// [FileShareShareSettingsUnion] to add type safety for the return type of
// [FileShareShareSettingsUnion.AsAny]
type anyFileShareShareSettings interface {
	implFileShareShareSettingsUnion()
}

func (FileShareShareSettingsStandard) implFileShareShareSettingsUnion() {}
func (FileShareShareSettingsVast) implFileShareShareSettingsUnion()     {}

// Use the following switch statement to find the correct variant
//
//	switch variant := FileShareShareSettingsUnion.AsAny().(type) {
//	case cloud.FileShareShareSettingsStandard:
//	case cloud.FileShareShareSettingsVast:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FileShareShareSettingsUnion) AsAny() anyFileShareShareSettings {
	switch u.TypeName {
	case "standard":
		return u.AsStandard()
	case "vast":
		return u.AsVast()
	}
	return nil
}

func (u FileShareShareSettingsUnion) AsStandard() (v FileShareShareSettingsStandard) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileShareShareSettingsUnion) AsVast() (v FileShareShareSettingsVast) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FileShareShareSettingsUnion) RawJSON() string { return u.JSON.raw }

func (r *FileShareShareSettingsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileShareShareSettingsStandard struct {
	// Standard file share type
	TypeName constant.Standard `json:"type_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TypeName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileShareShareSettingsStandard) RawJSON() string { return r.JSON.raw }
func (r *FileShareShareSettingsStandard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileShareShareSettingsVast struct {
	// Any of "LCD", "NPL".
	AllowedCharacters string `json:"allowed_characters,required"`
	// Any of "LCD", "NPL".
	PathLength string `json:"path_length,required"`
	// Enables or disables root squash for NFS clients.
	//
	//   - If `true`, root squash is enabled: the root user is mapped to nobody for all
	//     file and folder management operations on the export.
	//   - If `false`, root squash is disabled: the NFS client `root` user retains root
	//     privileges.
	RootSquash bool `json:"root_squash,required"`
	// Vast file share type
	TypeName constant.Vast `json:"type_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedCharacters respjson.Field
		PathLength        respjson.Field
		RootSquash        respjson.Field
		TypeName          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileShareShareSettingsVast) RawJSON() string { return r.JSON.raw }
func (r *FileShareShareSettingsVast) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File share status
type FileShareStatus string

const (
	FileShareStatusAvailable                      FileShareStatus = "available"
	FileShareStatusAwaitingTransfer               FileShareStatus = "awaiting_transfer"
	FileShareStatusBackupCreating                 FileShareStatus = "backup_creating"
	FileShareStatusBackupRestoring                FileShareStatus = "backup_restoring"
	FileShareStatusBackupRestoringError           FileShareStatus = "backup_restoring_error"
	FileShareStatusCreating                       FileShareStatus = "creating"
	FileShareStatusCreatingFromSnapshot           FileShareStatus = "creating_from_snapshot"
	FileShareStatusDeleted                        FileShareStatus = "deleted"
	FileShareStatusDeleting                       FileShareStatus = "deleting"
	FileShareStatusEnsuring                       FileShareStatus = "ensuring"
	FileShareStatusError                          FileShareStatus = "error"
	FileShareStatusErrorDeleting                  FileShareStatus = "error_deleting"
	FileShareStatusExtending                      FileShareStatus = "extending"
	FileShareStatusExtendingError                 FileShareStatus = "extending_error"
	FileShareStatusInactive                       FileShareStatus = "inactive"
	FileShareStatusManageError                    FileShareStatus = "manage_error"
	FileShareStatusManageStarting                 FileShareStatus = "manage_starting"
	FileShareStatusMigrating                      FileShareStatus = "migrating"
	FileShareStatusMigratingTo                    FileShareStatus = "migrating_to"
	FileShareStatusReplicationChange              FileShareStatus = "replication_change"
	FileShareStatusReverting                      FileShareStatus = "reverting"
	FileShareStatusRevertingError                 FileShareStatus = "reverting_error"
	FileShareStatusShrinking                      FileShareStatus = "shrinking"
	FileShareStatusShrinkingError                 FileShareStatus = "shrinking_error"
	FileShareStatusShrinkingPossibleDataLossError FileShareStatus = "shrinking_possible_data_loss_error"
	FileShareStatusUnmanageError                  FileShareStatus = "unmanage_error"
	FileShareStatusUnmanageStarting               FileShareStatus = "unmanage_starting"
	FileShareStatusUnmanaged                      FileShareStatus = "unmanaged"
)

// File share type name
type FileShareTypeName string

const (
	FileShareTypeNameStandard FileShareTypeName = "standard"
	FileShareTypeNameVast     FileShareTypeName = "vast"
)

// Deprecated. Use `type_name` instead. File share disk type
type FileShareVolumeType string

const (
	FileShareVolumeTypeDefaultShareType FileShareVolumeType = "default_share_type"
	FileShareVolumeTypeVastShareType    FileShareVolumeType = "vast_share_type"
)

type FileShareNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfCreateStandardFileShareSerializer *FileShareNewParamsBodyCreateStandardFileShareSerializer `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfCreateVastFileShareSerializer *FileShareNewParamsBodyCreateVastFileShareSerializer `json:",inline"`

	paramObj
}

func (u FileShareNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCreateStandardFileShareSerializer, u.OfCreateVastFileShareSerializer)
}
func (r *FileShareNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Network, Protocol, Size are required.
type FileShareNewParamsBodyCreateStandardFileShareSerializer struct {
	// File share name
	Name string `json:"name,required"`
	// File share network configuration
	Network FileShareNewParamsBodyCreateStandardFileShareSerializerNetwork `json:"network,omitzero,required"`
	// File share size in GiB
	Size int64 `json:"size,required"`
	// Access Rules
	Access []FileShareNewParamsBodyCreateStandardFileShareSerializerAccess `json:"access,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	// Standard file share type
	//
	// Any of "standard".
	TypeName string `json:"type_name,omitzero"`
	// Deprecated. Use `type_name` instead.
	//
	// Any of "default_share_type".
	//
	// Deprecated: deprecated
	VolumeType string `json:"volume_type,omitzero"`
	// File share protocol
	//
	// This field can be elided, and will marshal its zero value as "NFS".
	Protocol constant.Nfs `json:"protocol,required"`
	paramObj
}

func (r FileShareNewParamsBodyCreateStandardFileShareSerializer) MarshalJSON() (data []byte, err error) {
	type shadow FileShareNewParamsBodyCreateStandardFileShareSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileShareNewParamsBodyCreateStandardFileShareSerializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileShareNewParamsBodyCreateStandardFileShareSerializer](
		"type_name", "standard",
	)
	apijson.RegisterFieldValidator[FileShareNewParamsBodyCreateStandardFileShareSerializer](
		"volume_type", "default_share_type",
	)
}

// File share network configuration
//
// The property NetworkID is required.
type FileShareNewParamsBodyCreateStandardFileShareSerializerNetwork struct {
	// Network ID.
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// Subnetwork ID. If the subnet is not selected, it will be selected automatically.
	SubnetID param.Opt[string] `json:"subnet_id,omitzero" format:"uuid4"`
	paramObj
}

func (r FileShareNewParamsBodyCreateStandardFileShareSerializerNetwork) MarshalJSON() (data []byte, err error) {
	type shadow FileShareNewParamsBodyCreateStandardFileShareSerializerNetwork
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileShareNewParamsBodyCreateStandardFileShareSerializerNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessMode, IPAddress are required.
type FileShareNewParamsBodyCreateStandardFileShareSerializerAccess struct {
	// Access mode
	//
	// Any of "ro", "rw".
	AccessMode string `json:"access_mode,omitzero,required"`
	// Source IP or network
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	paramObj
}

func (r FileShareNewParamsBodyCreateStandardFileShareSerializerAccess) MarshalJSON() (data []byte, err error) {
	type shadow FileShareNewParamsBodyCreateStandardFileShareSerializerAccess
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileShareNewParamsBodyCreateStandardFileShareSerializerAccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileShareNewParamsBodyCreateStandardFileShareSerializerAccess](
		"access_mode", "ro", "rw",
	)
}

// The properties Name, Protocol, Size are required.
type FileShareNewParamsBodyCreateVastFileShareSerializer struct {
	// File share name
	Name string `json:"name,required"`
	// File share size in GiB
	Size int64 `json:"size,required"`
	// Configuration settings for the share
	ShareSettings FileShareNewParamsBodyCreateVastFileShareSerializerShareSettings `json:"share_settings,omitzero"`
	// Key-value tags to associate with the resource. A tag is a key-value pair that
	// can be associated with a resource, enabling efficient filtering and grouping for
	// better organization and management. Some tags are read-only and cannot be
	// modified by the user. Tags are also integrated with cost reports, allowing cost
	// data to be filtered based on tag keys or values.
	Tags map[string]string `json:"tags,omitzero"`
	// Vast file share type
	//
	// Any of "vast".
	TypeName string `json:"type_name,omitzero"`
	// Deprecated. Use `type_name` instead.
	//
	// Any of "vast_share_type".
	//
	// Deprecated: deprecated
	VolumeType string `json:"volume_type,omitzero"`
	// File share protocol
	//
	// This field can be elided, and will marshal its zero value as "NFS".
	Protocol constant.Nfs `json:"protocol,required"`
	paramObj
}

func (r FileShareNewParamsBodyCreateVastFileShareSerializer) MarshalJSON() (data []byte, err error) {
	type shadow FileShareNewParamsBodyCreateVastFileShareSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileShareNewParamsBodyCreateVastFileShareSerializer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileShareNewParamsBodyCreateVastFileShareSerializer](
		"type_name", "vast",
	)
	apijson.RegisterFieldValidator[FileShareNewParamsBodyCreateVastFileShareSerializer](
		"volume_type", "vast_share_type",
	)
}

// Configuration settings for the share
type FileShareNewParamsBodyCreateVastFileShareSerializerShareSettings struct {
	// Enables or disables root squash for NFS clients.
	//
	//   - If `true` (default), root squash is enabled: the root user is mapped to nobody
	//     for all file and folder management operations on the export.
	//   - If `false`, root squash is disabled: the NFS client `root` user retains root
	//     privileges. Use this option if you trust the root user not to perform
	//     operations that will corrupt data.
	RootSquash param.Opt[bool] `json:"root_squash,omitzero"`
	// Determines which characters are allowed in file names. Choose between:
	//
	//   - Lowest Common Denominator (LCD), allows only characters allowed by all VAST
	//     Cluster-supported protocols
	//   - Native Protocol Limit (NPL), imposes no limitation beyond that of the client
	//     protocol.
	//
	// Any of "LCD", "NPL".
	AllowedCharacters string `json:"allowed_characters,omitzero"`
	// Affects the maximum limit of file path component name length. Choose between:
	//
	//   - Lowest Common Denominator (LCD), imposes the lowest common denominator file
	//     length limit of all VAST Cluster-supported protocols. With this (default)
	//     option, the limitation on the length of a single component of the path is 255
	//     characters
	//   - Native Protocol Limit (NPL), imposes no limitation beyond that of the client
	//     protocol.
	//
	// Any of "LCD", "NPL".
	PathLength string `json:"path_length,omitzero"`
	paramObj
}

func (r FileShareNewParamsBodyCreateVastFileShareSerializerShareSettings) MarshalJSON() (data []byte, err error) {
	type shadow FileShareNewParamsBodyCreateVastFileShareSerializerShareSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileShareNewParamsBodyCreateVastFileShareSerializerShareSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileShareNewParamsBodyCreateVastFileShareSerializerShareSettings](
		"allowed_characters", "LCD", "NPL",
	)
	apijson.RegisterFieldValidator[FileShareNewParamsBodyCreateVastFileShareSerializerShareSettings](
		"path_length", "LCD", "NPL",
	)
}

type FileShareUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Name
	Name param.Opt[string] `json:"name,omitzero"`
	// Configuration settings for the share
	ShareSettings FileShareUpdateParamsShareSettings `json:"share_settings,omitzero"`
	// Update key-value tags using JSON Merge Patch semantics (RFC 7386). Provide
	// key-value pairs to add or update tags. Set tag values to `null` to remove tags.
	// Unspecified tags remain unchanged. Read-only tags are always preserved and
	// cannot be modified.
	//
	// **Examples:**
	//
	//   - **Add/update tags:**
	//     `{'tags': {'environment': 'production', 'team': 'backend'}}` adds new tags or
	//     updates existing ones.
	//   - **Delete tags:** `{'tags': {'`old_tag`': null}}` removes specific tags.
	//   - **Remove all tags:** `{'tags': null}` removes all user-managed tags (read-only
	//     tags are preserved).
	//   - **Partial update:** `{'tags': {'environment': 'staging'}}` only updates
	//     specified tags.
	//   - **Mixed operations:**
	//     `{'tags': {'environment': 'production', '`cost_center`': 'engineering', '`deprecated_tag`': null}}`
	//     adds/updates 'environment' and '`cost_center`' while removing
	//     '`deprecated_tag`', preserving other existing tags.
	//   - **Replace all:** first delete existing tags with null values, then add new
	//     ones in the same request.
	Tags TagUpdateMap `json:"tags,omitzero"`
	paramObj
}

func (r FileShareUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FileShareUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileShareUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration settings for the share
type FileShareUpdateParamsShareSettings struct {
	// Enables or disables root squash for NFS clients.
	//
	//   - If `true` (default), root squash is enabled: the root user is mapped to nobody
	//     for all file and folder management operations on the export.
	//   - If `false`, root squash is disabled: the NFS client `root` user retains root
	//     privileges. Use this option if you trust the root user not to perform
	//     operations that will corrupt data.
	RootSquash param.Opt[bool] `json:"root_squash,omitzero"`
	// Determines which characters are allowed in file names. Choose between:
	//
	//   - Lowest Common Denominator (LCD), allows only characters allowed by all VAST
	//     Cluster-supported protocols
	//   - Native Protocol Limit (NPL), imposes no limitation beyond that of the client
	//     protocol.
	//
	// Any of "LCD", "NPL".
	AllowedCharacters string `json:"allowed_characters,omitzero"`
	// Affects the maximum limit of file path component name length. Choose between:
	//
	//   - Lowest Common Denominator (LCD), imposes the lowest common denominator file
	//     length limit of all VAST Cluster-supported protocols. With this (default)
	//     option, the limitation on the length of a single component of the path is 255
	//     characters
	//   - Native Protocol Limit (NPL), imposes no limitation beyond that of the client
	//     protocol.
	//
	// Any of "LCD", "NPL".
	PathLength string `json:"path_length,omitzero"`
	paramObj
}

func (r FileShareUpdateParamsShareSettings) MarshalJSON() (data []byte, err error) {
	type shadow FileShareUpdateParamsShareSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileShareUpdateParamsShareSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileShareUpdateParamsShareSettings](
		"allowed_characters", "LCD", "NPL",
	)
	apijson.RegisterFieldValidator[FileShareUpdateParamsShareSettings](
		"path_length", "LCD", "NPL",
	)
}

type FileShareListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Optional. Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// File share name. Uses partial match.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Optional. Offset value is used to exclude the first set of records from the
	// result
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// File share type name
	//
	// Any of "standard", "vast".
	TypeName FileShareListParamsTypeName `query:"type_name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileShareListParams]'s query parameters as `url.Values`.
func (r FileShareListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// File share type name
type FileShareListParamsTypeName string

const (
	FileShareListParamsTypeNameStandard FileShareListParamsTypeName = "standard"
	FileShareListParamsTypeNameVast     FileShareListParamsTypeName = "vast"
)

type FileShareDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type FileShareGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type FileShareResizeParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// File Share new size in GiB.
	Size int64 `json:"size,required"`
	paramObj
}

func (r FileShareResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow FileShareResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileShareResizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
