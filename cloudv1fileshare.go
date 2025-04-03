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

// CloudV1FileShareService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1FileShareService] method instead.
type CloudV1FileShareService struct {
	Options      []option.RequestOption
	AccessRule   *CloudV1FileShareAccessRuleService
	Metadata     *CloudV1FileShareMetadataService
	MetadataItem *CloudV1FileShareMetadataItemService
}

// NewCloudV1FileShareService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1FileShareService(opts ...option.RequestOption) (r *CloudV1FileShareService) {
	r = &CloudV1FileShareService{}
	r.Options = opts
	r.AccessRule = NewCloudV1FileShareAccessRuleService(opts...)
	r.Metadata = NewCloudV1FileShareMetadataService(opts...)
	r.MetadataItem = NewCloudV1FileShareMetadataItemService(opts...)
	return
}

// Create file share
func (r *CloudV1FileShareService) New(ctx context.Context, projectID int64, regionID int64, body CloudV1FileShareNewParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get file share
func (r *CloudV1FileShareService) Get(ctx context.Context, projectID int64, regionID int64, fileShareID string, opts ...option.RequestOption) (res *FileShare, err error) {
	opts = append(r.Options[:], opts...)
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s", projectID, regionID, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Rename file share
func (r *CloudV1FileShareService) Update(ctx context.Context, projectID int64, regionID int64, fileShareID string, body CloudV1FileShareUpdateParams, opts ...option.RequestOption) (res *FileShare, err error) {
	opts = append(r.Options[:], opts...)
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s", projectID, regionID, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List file shares
func (r *CloudV1FileShareService) List(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *CloudV1FileShareListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete file share
func (r *CloudV1FileShareService) Delete(ctx context.Context, projectID int64, regionID int64, fileShareID string, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s", projectID, regionID, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Check if regional quota is exceeded, if yes the number of additional quotas
// needed to create the specified File Share will be calculated
func (r *CloudV1FileShareService) CheckQuota(ctx context.Context, projectID int64, regionID int64, body CloudV1FileShareCheckQuotaParams, opts ...option.RequestOption) (res *CloudV1FileShareCheckQuotaResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/check_limits", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Resize file share
func (r *CloudV1FileShareService) Resize(ctx context.Context, projectID int64, regionID int64, fileShareID string, body CloudV1FileShareResizeParams, opts ...option.RequestOption) (res *TaskIDs, err error) {
	opts = append(r.Options[:], opts...)
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s/extend", projectID, regionID, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FileShare struct {
	// File share ID
	ID string `json:"id,required" format:"uuid4"`
	// Connection point. Can be null during File share creation
	ConnectionPoint string `json:"connection_point,required,nullable"`
	// Datetime when the file share was created
	CreatedAt string `json:"created_at,required"`
	// Task that created this entity
	CreatorTaskID string `json:"creator_task_id,required" format:"uuid4"`
	// Key-value tags
	Metadata map[string]string `json:"metadata,required"`
	// Key-value tags with read-only flag
	MetadataDetailed []DetailedMetadata `json:"metadata_detailed,required"`
	// File share name
	Name string `json:"name,required"`
	// Network ID
	NetworkID string `json:"network_id,required" format:"uuid4"`
	// Network name
	NetworkName string `json:"network_name,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// File share protocol
	Protocol string `json:"protocol,required"`
	// Region name
	Region string `json:"region,required"`
	// Region ID
	RegionID int64 `json:"region_id,required"`
	// Share network name
	ShareNetworkName string `json:"share_network_name,required"`
	// File share size, GiB
	Size int64 `json:"size,required"`
	// File share status
	Status FileShareStatus `json:"status,required"`
	// Subnet ID
	SubnetID string `json:"subnet_id,required" format:"uuid4"`
	// Subnet name
	SubnetName string `json:"subnet_name,required"`
	// Active task. Can be null.
	TaskID string `json:"task_id,required,nullable" format:"uuid4"`
	// File share disk type
	VolumeType string        `json:"volume_type,required"`
	JSON       fileShareJSON `json:"-"`
}

// fileShareJSON contains the JSON metadata for the struct [FileShare]
type fileShareJSON struct {
	ID               apijson.Field
	ConnectionPoint  apijson.Field
	CreatedAt        apijson.Field
	CreatorTaskID    apijson.Field
	Metadata         apijson.Field
	MetadataDetailed apijson.Field
	Name             apijson.Field
	NetworkID        apijson.Field
	NetworkName      apijson.Field
	ProjectID        apijson.Field
	Protocol         apijson.Field
	Region           apijson.Field
	RegionID         apijson.Field
	ShareNetworkName apijson.Field
	Size             apijson.Field
	Status           apijson.Field
	SubnetID         apijson.Field
	SubnetName       apijson.Field
	TaskID           apijson.Field
	VolumeType       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FileShare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileShareJSON) RawJSON() string {
	return r.raw
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

func (r FileShareStatus) IsKnown() bool {
	switch r {
	case FileShareStatusAvailable, FileShareStatusAwaitingTransfer, FileShareStatusBackupCreating, FileShareStatusBackupRestoring, FileShareStatusBackupRestoringError, FileShareStatusCreating, FileShareStatusCreatingFromSnapshot, FileShareStatusDeleted, FileShareStatusDeleting, FileShareStatusError, FileShareStatusErrorDeleting, FileShareStatusExtending, FileShareStatusExtendingError, FileShareStatusInactive, FileShareStatusManageError, FileShareStatusManageStarting, FileShareStatusMigrating, FileShareStatusMigratingTo, FileShareStatusReplicationChange, FileShareStatusReverting, FileShareStatusRevertingError, FileShareStatusShrinking, FileShareStatusShrinkingError, FileShareStatusShrinkingPossibleDataLossError, FileShareStatusUnmanageError, FileShareStatusUnmanageStarting, FileShareStatusUnmanaged:
		return true
	}
	return false
}

type NameParam struct {
	// Name.
	Name param.Field[string] `json:"name,required"`
}

func (r NameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1FileShareListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []FileShare                      `json:"results,required"`
	JSON    cloudV1FileShareListResponseJSON `json:"-"`
}

// cloudV1FileShareListResponseJSON contains the JSON metadata for the struct
// [CloudV1FileShareListResponse]
type cloudV1FileShareListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1FileShareListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FileShareListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1FileShareCheckQuotaResponse struct {
	// Shared file system Count limit
	SfsCountLimit int64 `json:"sfs_count_limit"`
	// Shared file system Count requested
	SfsCountRequested int64 `json:"sfs_count_requested"`
	// Shared file system Count usage
	SfsCountUsage int64 `json:"sfs_count_usage"`
	// Shared file system Size, GiB limit
	SfsSizeLimit int64 `json:"sfs_size_limit"`
	// Shared file system Size, GiB requested
	SfsSizeRequested int64 `json:"sfs_size_requested"`
	// Shared file system Size, GiB usage
	SfsSizeUsage int64                                  `json:"sfs_size_usage"`
	JSON         cloudV1FileShareCheckQuotaResponseJSON `json:"-"`
}

// cloudV1FileShareCheckQuotaResponseJSON contains the JSON metadata for the struct
// [CloudV1FileShareCheckQuotaResponse]
type cloudV1FileShareCheckQuotaResponseJSON struct {
	SfsCountLimit     apijson.Field
	SfsCountRequested apijson.Field
	SfsCountUsage     apijson.Field
	SfsSizeLimit      apijson.Field
	SfsSizeRequested  apijson.Field
	SfsSizeUsage      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CloudV1FileShareCheckQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FileShareCheckQuotaResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1FileShareNewParams struct {
	// File share name
	Name param.Field[string] `json:"name,required"`
	// File share network configuration
	Network param.Field[CloudV1FileShareNewParamsNetwork] `json:"network,required"`
	// File share protocol
	Protocol param.Field[CloudV1FileShareNewParamsProtocol] `json:"protocol,required"`
	// File share size
	Size param.Field[int64] `json:"size,required"`
	// Access Rules
	Access param.Field[[]CreateAccessRuleParam] `json:"access"`
	// Create one or more key-value tags
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r CloudV1FileShareNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// File share network configuration
type CloudV1FileShareNewParamsNetwork struct {
	// Network ID.
	NetworkID param.Field[string] `json:"network_id,required" format:"uuid4"`
	// Subnetwork ID. If the subnet is not selected, it will be selected automatically.
	SubnetID param.Field[string] `json:"subnet_id" format:"uuid4"`
}

func (r CloudV1FileShareNewParamsNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// File share protocol
type CloudV1FileShareNewParamsProtocol string

const (
	CloudV1FileShareNewParamsProtocolNfs CloudV1FileShareNewParamsProtocol = "NFS"
)

func (r CloudV1FileShareNewParamsProtocol) IsKnown() bool {
	switch r {
	case CloudV1FileShareNewParamsProtocolNfs:
		return true
	}
	return false
}

type CloudV1FileShareUpdateParams struct {
	Name NameParam `json:"name,required"`
}

func (r CloudV1FileShareUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Name)
}

type CloudV1FileShareCheckQuotaParams struct {
	// File Share size in GiB
	Size param.Field[int64] `json:"size,required"`
}

func (r CloudV1FileShareCheckQuotaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1FileShareResizeParams struct {
	// File Share new size in GiB.
	Size param.Field[int64] `json:"size,required"`
}

func (r CloudV1FileShareResizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
