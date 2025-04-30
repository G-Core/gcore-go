// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
	"github.com/stainless-sdks/gcore-go/shared/constant"
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
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Rename file share
func (r *FileShareService) Update(ctx context.Context, fileShareID string, params FileShareUpdateParams, opts ...option.RequestOption) (res *FileShare, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List file shares
func (r *FileShareService) List(ctx context.Context, params FileShareListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[FileShare], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
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
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
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
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
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

// '#/components/schemas/FileShareSerializer'
// "$.components.schemas.FileShareSerializer"
type FileShare struct {
	// '#/components/schemas/FileShareSerializer/properties/id'
	// "$.components.schemas.FileShareSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/FileShareSerializer/properties/connection_point/anyOf/0'
	// "$.components.schemas.FileShareSerializer.properties.connection_point.anyOf[0]"
	ConnectionPoint string `json:"connection_point,required"`
	// '#/components/schemas/FileShareSerializer/properties/created_at'
	// "$.components.schemas.FileShareSerializer.properties.created_at"
	CreatedAt string `json:"created_at,required"`
	// '#/components/schemas/FileShareSerializer/properties/creator_task_id'
	// "$.components.schemas.FileShareSerializer.properties.creator_task_id"
	CreatorTaskID string `json:"creator_task_id,required" format:"uuid4"`
	// '#/components/schemas/FileShareSerializer/properties/name'
	// "$.components.schemas.FileShareSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/FileShareSerializer/properties/network_id'
	// "$.components.schemas.FileShareSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/FileShareSerializer/properties/network_name'
	// "$.components.schemas.FileShareSerializer.properties.network_name"
	NetworkName string `json:"network_name,required"`
	// '#/components/schemas/FileShareSerializer/properties/project_id'
	// "$.components.schemas.FileShareSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/FileShareSerializer/properties/protocol'
	// "$.components.schemas.FileShareSerializer.properties.protocol"
	Protocol string `json:"protocol,required"`
	// '#/components/schemas/FileShareSerializer/properties/region'
	// "$.components.schemas.FileShareSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/FileShareSerializer/properties/region_id'
	// "$.components.schemas.FileShareSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/FileShareSerializer/properties/share_network_name/anyOf/0'
	// "$.components.schemas.FileShareSerializer.properties.share_network_name.anyOf[0]"
	ShareNetworkName string `json:"share_network_name,required"`
	// '#/components/schemas/FileShareSerializer/properties/size'
	// "$.components.schemas.FileShareSerializer.properties.size"
	Size int64 `json:"size,required"`
	// '#/components/schemas/FileShareSerializer/properties/status'
	// "$.components.schemas.FileShareSerializer.properties.status"
	//
	// Any of "available", "awaiting_transfer", "backup_creating", "backup_restoring",
	// "backup_restoring_error", "creating", "creating_from_snapshot", "deleted",
	// "deleting", "ensuring", "error", "error_deleting", "extending",
	// "extending_error", "inactive", "manage_error", "manage_starting", "migrating",
	// "migrating_to", "replication_change", "reverting", "reverting_error",
	// "shrinking", "shrinking_error", "shrinking_possible_data_loss_error",
	// "unmanage_error", "unmanage_starting", "unmanaged".
	Status FileShareStatus `json:"status,required"`
	// '#/components/schemas/FileShareSerializer/properties/subnet_id'
	// "$.components.schemas.FileShareSerializer.properties.subnet_id"
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// '#/components/schemas/FileShareSerializer/properties/subnet_name'
	// "$.components.schemas.FileShareSerializer.properties.subnet_name"
	SubnetName string `json:"subnet_name,required"`
	// '#/components/schemas/FileShareSerializer/properties/tags'
	// "$.components.schemas.FileShareSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/FileShareSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.FileShareSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,required" format:"uuid4"`
	// '#/components/schemas/FileShareSerializer/properties/volume_type'
	// "$.components.schemas.FileShareSerializer.properties.volume_type"
	//
	// Any of "default_share_type", "vast_share_type".
	VolumeType FileShareVolumeType `json:"volume_type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID               resp.Field
		ConnectionPoint  resp.Field
		CreatedAt        resp.Field
		CreatorTaskID    resp.Field
		Name             resp.Field
		NetworkID        resp.Field
		NetworkName      resp.Field
		ProjectID        resp.Field
		Protocol         resp.Field
		Region           resp.Field
		RegionID         resp.Field
		ShareNetworkName resp.Field
		Size             resp.Field
		Status           resp.Field
		SubnetID         resp.Field
		SubnetName       resp.Field
		Tags             resp.Field
		TaskID           resp.Field
		VolumeType       resp.Field
		ExtraFields      map[string]resp.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileShare) RawJSON() string { return r.JSON.raw }
func (r *FileShare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/FileShareSerializer/properties/status'
// "$.components.schemas.FileShareSerializer.properties.status"
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

// '#/components/schemas/FileShareSerializer/properties/volume_type'
// "$.components.schemas.FileShareSerializer.properties.volume_type"
type FileShareVolumeType string

const (
	FileShareVolumeTypeDefaultShareType FileShareVolumeType = "default_share_type"
	FileShareVolumeTypeVastShareType    FileShareVolumeType = "vast_share_type"
)

type FileShareNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/requestBody/content/application%2Fjson/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}'].post.requestBody.content['application/json'].schema"
	Body FileShareNewParamsBodyUnion
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FileShareNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r FileShareNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/requestBody/content/application%2Fjson/schema'
// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}'].post.requestBody.content['application/json'].schema"
//
// Satisfied by [FileShareNewParamsBodyCreateStandardFileShareSerializer] and
// [FileShareNewParamsBodyCreateVastFileShareSerializer]
type FileShareNewParamsBodyUnion interface {
	implFileShareNewParamsBodyUnion()
}

func (FileShareNewParamsBodyCreateStandardFileShareSerializer) implFileShareNewParamsBodyUnion() {}
func (FileShareNewParamsBodyCreateVastFileShareSerializer) implFileShareNewParamsBodyUnion()     {}

// '#/components/schemas/CreateFileShareSerializer/anyOf/0'
// "$.components.schemas.CreateFileShareSerializer.anyOf[0]"
//
// The properties Name, Network, Protocol, Size are required.
type FileShareNewParamsBodyCreateStandardFileShareSerializer struct {
	// '#/components/schemas/CreateStandardFileShareSerializer/properties/name'
	// "$.components.schemas.CreateStandardFileShareSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateStandardFileShareSerializer/properties/network'
	// "$.components.schemas.CreateStandardFileShareSerializer.properties.network"
	Network FileShareNewParamsBodyCreateStandardFileShareSerializerNetwork `json:"network,omitzero,required"`
	// '#/components/schemas/CreateStandardFileShareSerializer/properties/size'
	// "$.components.schemas.CreateStandardFileShareSerializer.properties.size"
	Size int64 `json:"size,required"`
	// '#/components/schemas/CreateStandardFileShareSerializer/properties/access'
	// "$.components.schemas.CreateStandardFileShareSerializer.properties.access"
	Access []FileShareNewParamsBodyCreateStandardFileShareSerializerAccess `json:"access,omitzero"`
	// '#/components/schemas/CreateStandardFileShareSerializer/properties/tags'
	// "$.components.schemas.CreateStandardFileShareSerializer.properties.tags"
	Tags TagUpdateList `json:"tags,omitzero"`
	// '#/components/schemas/CreateStandardFileShareSerializer/properties/volume_type'
	// "$.components.schemas.CreateStandardFileShareSerializer.properties.volume_type"
	//
	// Any of "default_share_type".
	VolumeType string `json:"volume_type,omitzero"`
	// '#/components/schemas/CreateStandardFileShareSerializer/properties/protocol'
	// "$.components.schemas.CreateStandardFileShareSerializer.properties.protocol"
	//
	// This field can be elided, and will marshal its zero value as "NFS".
	Protocol constant.Nfs `json:"protocol,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FileShareNewParamsBodyCreateStandardFileShareSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FileShareNewParamsBodyCreateStandardFileShareSerializer) MarshalJSON() (data []byte, err error) {
	type shadow FileShareNewParamsBodyCreateStandardFileShareSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FileShareNewParamsBodyCreateStandardFileShareSerializer](
		"VolumeType", false, "default_share_type",
	)
}

// '#/components/schemas/CreateStandardFileShareSerializer/properties/network'
// "$.components.schemas.CreateStandardFileShareSerializer.properties.network"
//
// The property NetworkID is required.
type FileShareNewParamsBodyCreateStandardFileShareSerializerNetwork struct {
	// '#/components/schemas/FileShareNetworkSerializer/properties/network_id'
	// "$.components.schemas.FileShareNetworkSerializer.properties.network_id"
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// '#/components/schemas/FileShareNetworkSerializer/properties/subnet_id'
	// "$.components.schemas.FileShareNetworkSerializer.properties.subnet_id"
	SubnetID param.Opt[string] `json:"subnet_id,omitzero" format:"uuid4"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FileShareNewParamsBodyCreateStandardFileShareSerializerNetwork) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FileShareNewParamsBodyCreateStandardFileShareSerializerNetwork) MarshalJSON() (data []byte, err error) {
	type shadow FileShareNewParamsBodyCreateStandardFileShareSerializerNetwork
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateStandardFileShareSerializer/properties/access/items'
// "$.components.schemas.CreateStandardFileShareSerializer.properties.access.items"
//
// The properties AccessMode, IPAddress are required.
type FileShareNewParamsBodyCreateStandardFileShareSerializerAccess struct {
	// '#/components/schemas/CreateAccessRuleSerializer/properties/access_mode'
	// "$.components.schemas.CreateAccessRuleSerializer.properties.access_mode"
	//
	// Any of "ro", "rw".
	AccessMode string `json:"access_mode,omitzero,required"`
	// '#/components/schemas/CreateAccessRuleSerializer/properties/ip_address/anyOf/0'
	// "$.components.schemas.CreateAccessRuleSerializer.properties.ip_address.anyOf[0]"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FileShareNewParamsBodyCreateStandardFileShareSerializerAccess) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FileShareNewParamsBodyCreateStandardFileShareSerializerAccess) MarshalJSON() (data []byte, err error) {
	type shadow FileShareNewParamsBodyCreateStandardFileShareSerializerAccess
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[FileShareNewParamsBodyCreateStandardFileShareSerializerAccess](
		"AccessMode", false, "ro", "rw",
	)
}

// '#/components/schemas/CreateFileShareSerializer/anyOf/1'
// "$.components.schemas.CreateFileShareSerializer.anyOf[1]"
//
// The properties Name, Protocol, Size, VolumeType are required.
type FileShareNewParamsBodyCreateVastFileShareSerializer struct {
	// '#/components/schemas/CreateVastFileShareSerializer/properties/name'
	// "$.components.schemas.CreateVastFileShareSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateVastFileShareSerializer/properties/size'
	// "$.components.schemas.CreateVastFileShareSerializer.properties.size"
	Size int64 `json:"size,required"`
	// '#/components/schemas/CreateVastFileShareSerializer/properties/tags'
	// "$.components.schemas.CreateVastFileShareSerializer.properties.tags"
	Tags TagUpdateList `json:"tags,omitzero"`
	// '#/components/schemas/CreateVastFileShareSerializer/properties/protocol'
	// "$.components.schemas.CreateVastFileShareSerializer.properties.protocol"
	//
	// This field can be elided, and will marshal its zero value as "NFS".
	Protocol constant.Nfs `json:"protocol,required"`
	// '#/components/schemas/CreateVastFileShareSerializer/properties/volume_type'
	// "$.components.schemas.CreateVastFileShareSerializer.properties.volume_type"
	//
	// This field can be elided, and will marshal its zero value as "vast_share_type".
	VolumeType constant.VastShareType `json:"volume_type,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FileShareNewParamsBodyCreateVastFileShareSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r FileShareNewParamsBodyCreateVastFileShareSerializer) MarshalJSON() (data []byte, err error) {
	type shadow FileShareNewParamsBodyCreateVastFileShareSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

type FileShareUpdateParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfile_share_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}/{file_share_id}'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfile_share_id%7D/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}/{file_share_id}'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/NameSerializer/properties/name'
	// "$.components.schemas.NameSerializer.properties.name"
	Name string `json:"name,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FileShareUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r FileShareUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FileShareUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type FileShareListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}'].get.parameters[2]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}'].get.parameters[3]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FileShareListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [FileShareListParams]'s query parameters as `url.Values`.
func (r FileShareListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FileShareDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfile_share_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}/{file_share_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfile_share_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}/{file_share_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FileShareDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type FileShareGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfile_share_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}/{file_share_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfile_share_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}/{file_share_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FileShareGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type FileShareResizeParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfile_share_id%7D%2Fextend/post/parameters/0/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}/{file_share_id}/extend'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfile_share_id%7D%2Fextend/post/parameters/1/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}/{file_share_id}/extend'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/ResizeSfsSerializer/properties/size'
	// "$.components.schemas.ResizeSfsSerializer.properties.size"
	Size int64 `json:"size,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FileShareResizeParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r FileShareResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow FileShareResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
