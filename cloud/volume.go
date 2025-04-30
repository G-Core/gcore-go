// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
	"github.com/stainless-sdks/gcore-go/shared/constant"
)

// VolumeService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVolumeService] method instead.
type VolumeService struct {
	Options []option.RequestOption
}

// NewVolumeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVolumeService(opts ...option.RequestOption) (r VolumeService) {
	r = VolumeService{}
	r.Options = opts
	return
}

// Create volume
func (r *VolumeService) New(ctx context.Context, params VolumeNewParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v", params.ProjectID.Value, params.RegionID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Rename volume
func (r *VolumeService) Update(ctx context.Context, volumeID string, params VolumeUpdateParams, opts ...option.RequestOption) (res *Volume, err error) {
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List volumes
func (r *VolumeService) List(ctx context.Context, params VolumeListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[Volume], err error) {
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
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v", params.ProjectID.Value, params.RegionID.Value)
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

// List volumes
func (r *VolumeService) ListAutoPaging(ctx context.Context, params VolumeListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[Volume] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete volume
func (r *VolumeService) Delete(ctx context.Context, volumeID string, params VolumeDeleteParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s", params.ProjectID.Value, params.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Attach the volume to instance. Note: ultra volume can only be attached to an
// instance with shared flavor
func (r *VolumeService) AttachToInstance(ctx context.Context, volumeID string, params VolumeAttachToInstanceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/volumes/%v/%v/%s/attach", params.ProjectID.Value, params.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Change volume type
func (r *VolumeService) ChangeType(ctx context.Context, volumeID string, params VolumeChangeTypeParams, opts ...option.RequestOption) (res *Volume, err error) {
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/retype", params.ProjectID.Value, params.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Detach the volume from instance
func (r *VolumeService) DetachFromInstance(ctx context.Context, volumeID string, params VolumeDetachFromInstanceParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v2/volumes/%v/%v/%s/detach", params.ProjectID.Value, params.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get volume
func (r *VolumeService) Get(ctx context.Context, volumeID string, query VolumeGetParams, opts ...option.RequestOption) (res *Volume, err error) {
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s", query.ProjectID.Value, query.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Extend volume
func (r *VolumeService) Resize(ctx context.Context, volumeID string, params VolumeResizeParams, opts ...option.RequestOption) (res *TaskIDList, err error) {
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/extend", params.ProjectID.Value, params.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Revert volume to it's last snapshot
func (r *VolumeService) RevertToLastSnapshot(ctx context.Context, volumeID string, body VolumeRevertToLastSnapshotParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	if volumeID == "" {
		err = errors.New("missing required volume_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/volumes/%v/%v/%s/revert", body.ProjectID.Value, body.RegionID.Value, volumeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// '#/components/schemas/VolumeSerializer' "$.components.schemas.VolumeSerializer"
type Volume struct {
	// '#/components/schemas/VolumeSerializer/properties/id'
	// "$.components.schemas.VolumeSerializer.properties.id"
	ID string `json:"id,required"`
	// '#/components/schemas/VolumeSerializer/properties/bootable'
	// "$.components.schemas.VolumeSerializer.properties.bootable"
	Bootable bool `json:"bootable,required"`
	// '#/components/schemas/VolumeSerializer/properties/created_at'
	// "$.components.schemas.VolumeSerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/VolumeSerializer/properties/is_root_volume'
	// "$.components.schemas.VolumeSerializer.properties.is_root_volume"
	IsRootVolume bool `json:"is_root_volume,required"`
	// '#/components/schemas/VolumeSerializer/properties/name'
	// "$.components.schemas.VolumeSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/VolumeSerializer/properties/project_id'
	// "$.components.schemas.VolumeSerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/VolumeSerializer/properties/region'
	// "$.components.schemas.VolumeSerializer.properties.region"
	Region string `json:"region,required"`
	// '#/components/schemas/VolumeSerializer/properties/region_id'
	// "$.components.schemas.VolumeSerializer.properties.region_id"
	RegionID int64 `json:"region_id,required"`
	// '#/components/schemas/VolumeSerializer/properties/size'
	// "$.components.schemas.VolumeSerializer.properties.size"
	Size int64 `json:"size,required"`
	// '#/components/schemas/VolumeSerializer/properties/status'
	// "$.components.schemas.VolumeSerializer.properties.status"
	//
	// Any of "attaching", "available", "awaiting-transfer", "backing-up", "creating",
	// "deleting", "detaching", "downloading", "error", "error_backing-up",
	// "error_deleting", "error_extending", "error_restoring", "extending", "in-use",
	// "maintenance", "reserved", "restoring-backup", "retyping", "reverting",
	// "uploading".
	Status VolumeStatus `json:"status,required"`
	// '#/components/schemas/VolumeSerializer/properties/tags'
	// "$.components.schemas.VolumeSerializer.properties.tags"
	Tags []Tag `json:"tags,required"`
	// '#/components/schemas/VolumeSerializer/properties/volume_type'
	// "$.components.schemas.VolumeSerializer.properties.volume_type"
	VolumeType string `json:"volume_type,required"`
	// '#/components/schemas/VolumeSerializer/properties/attachments/anyOf/0'
	// "$.components.schemas.VolumeSerializer.properties.attachments.anyOf[0]"
	Attachments []VolumeAttachment `json:"attachments,nullable"`
	// '#/components/schemas/VolumeSerializer/properties/creator_task_id/anyOf/0'
	// "$.components.schemas.VolumeSerializer.properties.creator_task_id.anyOf[0]"
	CreatorTaskID string `json:"creator_task_id,nullable"`
	// '#/components/schemas/VolumeSerializer/properties/limiter_stats/anyOf/0'
	// "$.components.schemas.VolumeSerializer.properties.limiter_stats.anyOf[0]"
	LimiterStats VolumeLimiterStats `json:"limiter_stats,nullable"`
	// '#/components/schemas/VolumeSerializer/properties/snapshot_ids/anyOf/0'
	// "$.components.schemas.VolumeSerializer.properties.snapshot_ids.anyOf[0]"
	SnapshotIDs []string `json:"snapshot_ids,nullable"`
	// '#/components/schemas/VolumeSerializer/properties/task_id/anyOf/0'
	// "$.components.schemas.VolumeSerializer.properties.task_id.anyOf[0]"
	TaskID string `json:"task_id,nullable"`
	// '#/components/schemas/VolumeSerializer/properties/updated_at/anyOf/0'
	// "$.components.schemas.VolumeSerializer.properties.updated_at.anyOf[0]"
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// '#/components/schemas/VolumeSerializer/properties/volume_image_metadata/anyOf/0'
	// "$.components.schemas.VolumeSerializer.properties.volume_image_metadata.anyOf[0]"
	VolumeImageMetadata map[string]string `json:"volume_image_metadata,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                  resp.Field
		Bootable            resp.Field
		CreatedAt           resp.Field
		IsRootVolume        resp.Field
		Name                resp.Field
		ProjectID           resp.Field
		Region              resp.Field
		RegionID            resp.Field
		Size                resp.Field
		Status              resp.Field
		Tags                resp.Field
		VolumeType          resp.Field
		Attachments         resp.Field
		CreatorTaskID       resp.Field
		LimiterStats        resp.Field
		SnapshotIDs         resp.Field
		TaskID              resp.Field
		UpdatedAt           resp.Field
		VolumeImageMetadata resp.Field
		ExtraFields         map[string]resp.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Volume) RawJSON() string { return r.JSON.raw }
func (r *Volume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/VolumeSerializer/properties/status'
// "$.components.schemas.VolumeSerializer.properties.status"
type VolumeStatus string

const (
	VolumeStatusAttaching        VolumeStatus = "attaching"
	VolumeStatusAvailable        VolumeStatus = "available"
	VolumeStatusAwaitingTransfer VolumeStatus = "awaiting-transfer"
	VolumeStatusBackingUp        VolumeStatus = "backing-up"
	VolumeStatusCreating         VolumeStatus = "creating"
	VolumeStatusDeleting         VolumeStatus = "deleting"
	VolumeStatusDetaching        VolumeStatus = "detaching"
	VolumeStatusDownloading      VolumeStatus = "downloading"
	VolumeStatusError            VolumeStatus = "error"
	VolumeStatusErrorBackingUp   VolumeStatus = "error_backing-up"
	VolumeStatusErrorDeleting    VolumeStatus = "error_deleting"
	VolumeStatusErrorExtending   VolumeStatus = "error_extending"
	VolumeStatusErrorRestoring   VolumeStatus = "error_restoring"
	VolumeStatusExtending        VolumeStatus = "extending"
	VolumeStatusInUse            VolumeStatus = "in-use"
	VolumeStatusMaintenance      VolumeStatus = "maintenance"
	VolumeStatusReserved         VolumeStatus = "reserved"
	VolumeStatusRestoringBackup  VolumeStatus = "restoring-backup"
	VolumeStatusRetyping         VolumeStatus = "retyping"
	VolumeStatusReverting        VolumeStatus = "reverting"
	VolumeStatusUploading        VolumeStatus = "uploading"
)

// '#/components/schemas/VolumeSerializer/properties/attachments/anyOf/0/items'
// "$.components.schemas.VolumeSerializer.properties.attachments.anyOf[0].items"
type VolumeAttachment struct {
	// '#/components/schemas/VolumeAttachmentSerializer/properties/attachment_id'
	// "$.components.schemas.VolumeAttachmentSerializer.properties.attachment_id"
	AttachmentID string `json:"attachment_id,required"`
	// '#/components/schemas/VolumeAttachmentSerializer/properties/volume_id'
	// "$.components.schemas.VolumeAttachmentSerializer.properties.volume_id"
	VolumeID string `json:"volume_id,required"`
	// '#/components/schemas/VolumeAttachmentSerializer/properties/attached_at/anyOf/0'
	// "$.components.schemas.VolumeAttachmentSerializer.properties.attached_at.anyOf[0]"
	AttachedAt time.Time `json:"attached_at,nullable" format:"date-time"`
	// '#/components/schemas/VolumeAttachmentSerializer/properties/device/anyOf/0'
	// "$.components.schemas.VolumeAttachmentSerializer.properties.device.anyOf[0]"
	Device string `json:"device,nullable"`
	// '#/components/schemas/VolumeAttachmentSerializer/properties/flavor_id/anyOf/0'
	// "$.components.schemas.VolumeAttachmentSerializer.properties.flavor_id.anyOf[0]"
	FlavorID string `json:"flavor_id,nullable"`
	// '#/components/schemas/VolumeAttachmentSerializer/properties/instance_name/anyOf/0'
	// "$.components.schemas.VolumeAttachmentSerializer.properties.instance_name.anyOf[0]"
	InstanceName string `json:"instance_name,nullable"`
	// '#/components/schemas/VolumeAttachmentSerializer/properties/server_id/anyOf/0'
	// "$.components.schemas.VolumeAttachmentSerializer.properties.server_id.anyOf[0]"
	ServerID string `json:"server_id,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AttachmentID resp.Field
		VolumeID     resp.Field
		AttachedAt   resp.Field
		Device       resp.Field
		FlavorID     resp.Field
		InstanceName resp.Field
		ServerID     resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeAttachment) RawJSON() string { return r.JSON.raw }
func (r *VolumeAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/VolumeSerializer/properties/limiter_stats/anyOf/0'
// "$.components.schemas.VolumeSerializer.properties.limiter_stats.anyOf[0]"
type VolumeLimiterStats struct {
	// '#/components/schemas/VolumeLimiterStatsSerializer/properties/iops_base_limit'
	// "$.components.schemas.VolumeLimiterStatsSerializer.properties.iops_base_limit"
	IopsBaseLimit int64 `json:"iops_base_limit,required"`
	// '#/components/schemas/VolumeLimiterStatsSerializer/properties/iops_burst_limit'
	// "$.components.schemas.VolumeLimiterStatsSerializer.properties.iops_burst_limit"
	IopsBurstLimit int64 `json:"iops_burst_limit,required"`
	// '#/components/schemas/VolumeLimiterStatsSerializer/properties/MBps_base_limit'
	// "$.components.schemas.VolumeLimiterStatsSerializer.properties.MBps_base_limit"
	MBpsBaseLimit int64 `json:"MBps_base_limit,required"`
	// '#/components/schemas/VolumeLimiterStatsSerializer/properties/MBps_burst_limit'
	// "$.components.schemas.VolumeLimiterStatsSerializer.properties.MBps_burst_limit"
	MBpsBurstLimit int64 `json:"MBps_burst_limit,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IopsBaseLimit  resp.Field
		IopsBurstLimit resp.Field
		MBpsBaseLimit  resp.Field
		MBpsBurstLimit resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeLimiterStats) RawJSON() string { return r.JSON.raw }
func (r *VolumeLimiterStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/parameters/1/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/requestBody/content/application%2Fjson/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].post.requestBody.content['application/json'].schema"
	Body VolumeNewParamsBodyUnion
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VolumeNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r VolumeNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}

// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/post/requestBody/content/application%2Fjson/schema'
// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].post.requestBody.content['application/json'].schema"
//
// Satisfied by [VolumeNewParamsBodyCreateVolumeFromImageSerializer],
// [VolumeNewParamsBodyCreateVolumeFromSnapshotSerializer] and
// [VolumeNewParamsBodyCreateNewVolumeSerializer]
type VolumeNewParamsBodyUnion interface {
	implVolumeNewParamsBodyUnion()
}

func (VolumeNewParamsBodyCreateVolumeFromImageSerializer) implVolumeNewParamsBodyUnion()    {}
func (VolumeNewParamsBodyCreateVolumeFromSnapshotSerializer) implVolumeNewParamsBodyUnion() {}
func (VolumeNewParamsBodyCreateNewVolumeSerializer) implVolumeNewParamsBodyUnion()          {}

// '#/components/schemas/CreateVolumeSerializer/anyOf/0'
// "$.components.schemas.CreateVolumeSerializer.anyOf[0]"
//
// The properties ImageID, Name, Size, Source are required.
type VolumeNewParamsBodyCreateVolumeFromImageSerializer struct {
	// '#/components/schemas/CreateVolumeFromImageSerializer/properties/image_id'
	// "$.components.schemas.CreateVolumeFromImageSerializer.properties.image_id"
	ImageID string `json:"image_id,required" format:"uuid4"`
	// '#/components/schemas/CreateVolumeFromImageSerializer/properties/name'
	// "$.components.schemas.CreateVolumeFromImageSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateVolumeFromImageSerializer/properties/size'
	// "$.components.schemas.CreateVolumeFromImageSerializer.properties.size"
	Size int64 `json:"size,required"`
	// '#/components/schemas/CreateVolumeFromImageSerializer/properties/attachment_tag'
	// "$.components.schemas.CreateVolumeFromImageSerializer.properties.attachment_tag"
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	// '#/components/schemas/CreateVolumeFromImageSerializer/properties/instance_id_to_attach_to'
	// "$.components.schemas.CreateVolumeFromImageSerializer.properties.instance_id_to_attach_to"
	InstanceIDToAttachTo param.Opt[string] `json:"instance_id_to_attach_to,omitzero"`
	// '#/components/schemas/CreateVolumeFromImageSerializer/properties/lifecycle_policy_ids'
	// "$.components.schemas.CreateVolumeFromImageSerializer.properties.lifecycle_policy_ids"
	LifecyclePolicyIDs []int64 `json:"lifecycle_policy_ids,omitzero"`
	// '#/components/schemas/CreateVolumeFromImageSerializer/properties/tags'
	// "$.components.schemas.CreateVolumeFromImageSerializer.properties.tags"
	Tags map[string]string `json:"tags,omitzero"`
	// '#/components/schemas/CreateVolumeFromImageSerializer/properties/type_name'
	// "$.components.schemas.CreateVolumeFromImageSerializer.properties.type_name"
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	TypeName string `json:"type_name,omitzero"`
	// '#/components/schemas/CreateVolumeFromImageSerializer/properties/source'
	// "$.components.schemas.CreateVolumeFromImageSerializer.properties.source"
	//
	// This field can be elided, and will marshal its zero value as "image".
	Source constant.Image `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VolumeNewParamsBodyCreateVolumeFromImageSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r VolumeNewParamsBodyCreateVolumeFromImageSerializer) MarshalJSON() (data []byte, err error) {
	type shadow VolumeNewParamsBodyCreateVolumeFromImageSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[VolumeNewParamsBodyCreateVolumeFromImageSerializer](
		"TypeName", false, "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

// '#/components/schemas/CreateVolumeSerializer/anyOf/1'
// "$.components.schemas.CreateVolumeSerializer.anyOf[1]"
//
// The properties Name, SnapshotID, Source are required.
type VolumeNewParamsBodyCreateVolumeFromSnapshotSerializer struct {
	// '#/components/schemas/CreateVolumeFromSnapshotSerializer/properties/name'
	// "$.components.schemas.CreateVolumeFromSnapshotSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateVolumeFromSnapshotSerializer/properties/snapshot_id'
	// "$.components.schemas.CreateVolumeFromSnapshotSerializer.properties.snapshot_id"
	SnapshotID string `json:"snapshot_id,required" format:"uuid4"`
	// '#/components/schemas/CreateVolumeFromSnapshotSerializer/properties/attachment_tag'
	// "$.components.schemas.CreateVolumeFromSnapshotSerializer.properties.attachment_tag"
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	// '#/components/schemas/CreateVolumeFromSnapshotSerializer/properties/instance_id_to_attach_to'
	// "$.components.schemas.CreateVolumeFromSnapshotSerializer.properties.instance_id_to_attach_to"
	InstanceIDToAttachTo param.Opt[string] `json:"instance_id_to_attach_to,omitzero"`
	// '#/components/schemas/CreateVolumeFromSnapshotSerializer/properties/size'
	// "$.components.schemas.CreateVolumeFromSnapshotSerializer.properties.size"
	Size param.Opt[int64] `json:"size,omitzero"`
	// '#/components/schemas/CreateVolumeFromSnapshotSerializer/properties/lifecycle_policy_ids'
	// "$.components.schemas.CreateVolumeFromSnapshotSerializer.properties.lifecycle_policy_ids"
	LifecyclePolicyIDs []int64 `json:"lifecycle_policy_ids,omitzero"`
	// '#/components/schemas/CreateVolumeFromSnapshotSerializer/properties/tags'
	// "$.components.schemas.CreateVolumeFromSnapshotSerializer.properties.tags"
	Tags map[string]string `json:"tags,omitzero"`
	// '#/components/schemas/CreateVolumeFromSnapshotSerializer/properties/type_name'
	// "$.components.schemas.CreateVolumeFromSnapshotSerializer.properties.type_name"
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	TypeName string `json:"type_name,omitzero"`
	// '#/components/schemas/CreateVolumeFromSnapshotSerializer/properties/source'
	// "$.components.schemas.CreateVolumeFromSnapshotSerializer.properties.source"
	//
	// This field can be elided, and will marshal its zero value as "snapshot".
	Source constant.Snapshot `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VolumeNewParamsBodyCreateVolumeFromSnapshotSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r VolumeNewParamsBodyCreateVolumeFromSnapshotSerializer) MarshalJSON() (data []byte, err error) {
	type shadow VolumeNewParamsBodyCreateVolumeFromSnapshotSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[VolumeNewParamsBodyCreateVolumeFromSnapshotSerializer](
		"TypeName", false, "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

// '#/components/schemas/CreateVolumeSerializer/anyOf/2'
// "$.components.schemas.CreateVolumeSerializer.anyOf[2]"
//
// The properties Name, Size, Source are required.
type VolumeNewParamsBodyCreateNewVolumeSerializer struct {
	// '#/components/schemas/CreateNewVolumeSerializer/properties/name'
	// "$.components.schemas.CreateNewVolumeSerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateNewVolumeSerializer/properties/size'
	// "$.components.schemas.CreateNewVolumeSerializer.properties.size"
	Size int64 `json:"size,required"`
	// '#/components/schemas/CreateNewVolumeSerializer/properties/attachment_tag'
	// "$.components.schemas.CreateNewVolumeSerializer.properties.attachment_tag"
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	// '#/components/schemas/CreateNewVolumeSerializer/properties/instance_id_to_attach_to'
	// "$.components.schemas.CreateNewVolumeSerializer.properties.instance_id_to_attach_to"
	InstanceIDToAttachTo param.Opt[string] `json:"instance_id_to_attach_to,omitzero"`
	// '#/components/schemas/CreateNewVolumeSerializer/properties/lifecycle_policy_ids'
	// "$.components.schemas.CreateNewVolumeSerializer.properties.lifecycle_policy_ids"
	LifecyclePolicyIDs []int64 `json:"lifecycle_policy_ids,omitzero"`
	// '#/components/schemas/CreateNewVolumeSerializer/properties/tags'
	// "$.components.schemas.CreateNewVolumeSerializer.properties.tags"
	Tags map[string]string `json:"tags,omitzero"`
	// '#/components/schemas/CreateNewVolumeSerializer/properties/type_name'
	// "$.components.schemas.CreateNewVolumeSerializer.properties.type_name"
	//
	// Any of "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra".
	TypeName string `json:"type_name,omitzero"`
	// '#/components/schemas/CreateNewVolumeSerializer/properties/source'
	// "$.components.schemas.CreateNewVolumeSerializer.properties.source"
	//
	// This field can be elided, and will marshal its zero value as "new-volume".
	Source constant.NewVolume `json:"source,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VolumeNewParamsBodyCreateNewVolumeSerializer) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r VolumeNewParamsBodyCreateNewVolumeSerializer) MarshalJSON() (data []byte, err error) {
	type shadow VolumeNewParamsBodyCreateNewVolumeSerializer
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[VolumeNewParamsBodyCreateNewVolumeSerializer](
		"TypeName", false, "cold", "ssd_hiiops", "ssd_local", "ssd_lowlatency", "standard", "ultra",
	)
}

type VolumeUpdateParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}/{volume_id}'].patch.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D/patch/parameters/1/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}/{volume_id}'].patch.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/NameSerializer/properties/name'
	// "$.components.schemas.NameSerializer.properties.name"
	Name string `json:"name,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VolumeUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r VolumeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type VolumeListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].get.parameters[2]"
	Bootable param.Opt[bool] `query:"bootable,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].get.parameters[3]"
	ClusterID param.Opt[string] `query:"cluster_id,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/4'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].get.parameters[4]"
	HasAttachments param.Opt[bool] `query:"has_attachments,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/5'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].get.parameters[5]"
	IDPart param.Opt[string] `query:"id_part,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/6'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].get.parameters[6]"
	InstanceID param.Opt[string] `query:"instance_id,omitzero" format:"uuid4" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/7'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].get.parameters[7]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/8'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].get.parameters[8]"
	NamePart param.Opt[string] `query:"name_part,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/9'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].get.parameters[9]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/11'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].get.parameters[11]"
	TagKeyValue param.Opt[string] `query:"tag_key_value,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D/get/parameters/10'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}'].get.parameters[10]"
	TagKey []string `query:"tag_key,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VolumeListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [VolumeListParams]'s query parameters as `url.Values`.
func (r VolumeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type VolumeDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}/{volume_id}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}/{volume_id}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D/delete/parameters/3'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}/{volume_id}']['delete'].parameters[3]"
	Snapshots param.Opt[string] `query:"snapshots,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VolumeDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [VolumeDeleteParams]'s query parameters as `url.Values`.
func (r VolumeDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type VolumeAttachToInstanceParams struct {
	// '#/paths/%2Fcloud%2Fv2%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D%2Fattach/post/parameters/0/schema'
	// "$.paths['/cloud/v2/volumes/{project_id}/{region_id}/{volume_id}/attach'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D%2Fattach/post/parameters/1/schema'
	// "$.paths['/cloud/v2/volumes/{project_id}/{region_id}/{volume_id}/attach'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/AttachVolumeToInstanceSerializer/properties/instance_id'
	// "$.components.schemas.AttachVolumeToInstanceSerializer.properties.instance_id"
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	// '#/components/schemas/AttachVolumeToInstanceSerializer/properties/attachment_tag'
	// "$.components.schemas.AttachVolumeToInstanceSerializer.properties.attachment_tag"
	AttachmentTag param.Opt[string] `json:"attachment_tag,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VolumeAttachToInstanceParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r VolumeAttachToInstanceParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeAttachToInstanceParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type VolumeChangeTypeParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D%2Fretype/post/parameters/0/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}/{volume_id}/retype'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D%2Fretype/post/parameters/1/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}/{volume_id}/retype'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/VolumeRetypeSerializer/properties/volume_type'
	// "$.components.schemas.VolumeRetypeSerializer.properties.volume_type"
	//
	// Any of "ssd_hiiops", "standard".
	VolumeType VolumeChangeTypeParamsVolumeType `json:"volume_type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VolumeChangeTypeParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r VolumeChangeTypeParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeChangeTypeParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/VolumeRetypeSerializer/properties/volume_type'
// "$.components.schemas.VolumeRetypeSerializer.properties.volume_type"
type VolumeChangeTypeParamsVolumeType string

const (
	VolumeChangeTypeParamsVolumeTypeSsdHiiops VolumeChangeTypeParamsVolumeType = "ssd_hiiops"
	VolumeChangeTypeParamsVolumeTypeStandard  VolumeChangeTypeParamsVolumeType = "standard"
)

type VolumeDetachFromInstanceParams struct {
	// '#/paths/%2Fcloud%2Fv2%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D%2Fdetach/post/parameters/0/schema'
	// "$.paths['/cloud/v2/volumes/{project_id}/{region_id}/{volume_id}/detach'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv2%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D%2Fdetach/post/parameters/1/schema'
	// "$.paths['/cloud/v2/volumes/{project_id}/{region_id}/{volume_id}/detach'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/InstanceIdSerializer/properties/instance_id'
	// "$.components.schemas.InstanceIdSerializer.properties.instance_id"
	InstanceID string `json:"instance_id,required" format:"uuid4"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VolumeDetachFromInstanceParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r VolumeDetachFromInstanceParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeDetachFromInstanceParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type VolumeGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}/{volume_id}'].get.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D/get/parameters/1/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}/{volume_id}'].get.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VolumeGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type VolumeResizeParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D%2Fextend/post/parameters/0/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}/{volume_id}/extend'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D%2Fextend/post/parameters/1/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}/{volume_id}/extend'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/SizeSerializer/properties/size'
	// "$.components.schemas.SizeSerializer.properties.size"
	Size int64 `json:"size,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VolumeResizeParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r VolumeResizeParams) MarshalJSON() (data []byte, err error) {
	type shadow VolumeResizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type VolumeRevertToLastSnapshotParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D%2Frevert/post/parameters/0/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}/{volume_id}/revert'].post.parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fvolumes%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bvolume_id%7D%2Frevert/post/parameters/1/schema'
	// "$.paths['/cloud/v1/volumes/{project_id}/{region_id}/{volume_id}/revert'].post.parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VolumeRevertToLastSnapshotParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
