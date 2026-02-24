// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package streaming

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// DirectoryService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDirectoryService] method instead.
type DirectoryService struct {
	Options []option.RequestOption
}

// NewDirectoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDirectoryService(opts ...option.RequestOption) (r DirectoryService) {
	r = DirectoryService{}
	r.Options = opts
	return
}

// Use this method to create a new directory entity.
func (r *DirectoryService) New(ctx context.Context, body DirectoryNewParams, opts ...option.RequestOption) (res *DirectoryBase, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "streaming/directories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change a directory name or move to another "parent_id".
func (r *DirectoryService) Update(ctx context.Context, directoryID int64, body DirectoryUpdateParams, opts ...option.RequestOption) (res *DirectoryBase, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/directories/%v", directoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete a directory **and all entities inside**.
//
// After its execution, all contents of the directory will be deleted recursively:
//
// - Subdirectories
// - Videos
//
// The directory and contents are deleted permanently and irreversibly. Therefore,
// it is impossible to restore files after this.
//
// For details, see the Product Documentation.
func (r *DirectoryService) Delete(ctx context.Context, directoryID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("streaming/directories/%v", directoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Complete directory structure with contents. The structure contains both
// subfolders and videos in a continuous list.
func (r *DirectoryService) Get(ctx context.Context, directoryID int64, opts ...option.RequestOption) (res *DirectoryGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("streaming/directories/%v", directoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Tree structure of directories.
//
// This endpoint returns hierarchical data about directories in video hosting.
func (r *DirectoryService) GetTree(ctx context.Context, opts ...option.RequestOption) (res *DirectoriesTree, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "streaming/directories/tree"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DirectoriesTree struct {
	Tree []DirectoriesTreeTree `json:"tree"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tree        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectoriesTree) RawJSON() string { return r.JSON.raw }
func (r *DirectoriesTree) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoriesTreeTree struct {
	// Array of subdirectories, if any.
	Descendants []DirectoriesTree `json:"descendants"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Descendants respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	DirectoryBase
}

// Returns the unmodified JSON received from the API
func (r DirectoriesTreeTree) RawJSON() string { return r.JSON.raw }
func (r *DirectoriesTreeTree) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryBase struct {
	// ID of the directory
	ID int64 `json:"id"`
	// Time of creation. Datetime in ISO 8601 format.
	CreatedAt string `json:"created_at"`
	// Number of objects in this directory. Counting files and folders. The quantity is
	// calculated only at one level (not recursively in all subfolders).
	ItemsCount int64 `json:"items_count"`
	// Title of the directory
	Name string `json:"name"`
	// ID of a parent directory. "null" if it's in the root.
	ParentID int64 `json:"parent_id"`
	// Time of last update of the directory entity. Datetime in ISO 8601 format.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ItemsCount  respjson.Field
		Name        respjson.Field
		ParentID    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectoryBase) RawJSON() string { return r.JSON.raw }
func (r *DirectoryBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryItem struct {
	// Type of the entity: directory, or video
	//
	// Any of "Directory".
	ItemType string `json:"item_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	DirectoryBase
}

// Returns the unmodified JSON received from the API
func (r DirectoryItem) RawJSON() string { return r.JSON.raw }
func (r *DirectoryItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryVideo struct {
	// Type of the entity: directory, or video
	//
	// Any of "Video".
	ItemType string `json:"item_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Video
}

// Returns the unmodified JSON received from the API
func (r DirectoryVideo) RawJSON() string { return r.JSON.raw }
func (r *DirectoryVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryGetResponse struct {
	Directory DirectoryGetResponseDirectory `json:"directory"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Directory   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectoryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DirectoryGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryGetResponseDirectory struct {
	// Array of subdirectories, if any.
	Items []DirectoryGetResponseDirectoryItemUnion `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	DirectoryBase
}

// Returns the unmodified JSON received from the API
func (r DirectoryGetResponseDirectory) RawJSON() string { return r.JSON.raw }
func (r *DirectoryGetResponseDirectory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DirectoryGetResponseDirectoryItemUnion contains all possible properties and
// values from [DirectoryVideo], [DirectoryItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DirectoryGetResponseDirectoryItemUnion struct {
	ID int64 `json:"id"`
	// This field is from variant [DirectoryVideo].
	AdID int64 `json:"ad_id"`
	// This field is from variant [DirectoryVideo].
	CDNViews int64 `json:"cdn_views"`
	// This field is from variant [DirectoryVideo].
	ClientID int64 `json:"client_id"`
	// This field is from variant [DirectoryVideo].
	ClientUserID int64 `json:"client_user_id"`
	// This field is from variant [DirectoryVideo].
	ConvertedVideos []VideoConvertedVideo `json:"converted_videos"`
	// This field is from variant [DirectoryVideo].
	CustomIframeURL string `json:"custom_iframe_url"`
	// This field is from variant [DirectoryVideo].
	DashURL string `json:"dash_url"`
	// This field is from variant [DirectoryVideo].
	Description string `json:"description"`
	// This field is from variant [DirectoryVideo].
	Duration int64 `json:"duration"`
	// This field is from variant [DirectoryVideo].
	Error string `json:"error"`
	// This field is from variant [DirectoryVideo].
	HlsCmafURL string `json:"hls_cmaf_url"`
	// This field is from variant [DirectoryVideo].
	HlsURL string `json:"hls_url"`
	// This field is from variant [DirectoryVideo].
	IframeURL string `json:"iframe_url"`
	Name      string `json:"name"`
	// This field is from variant [DirectoryVideo].
	OriginSize int64 `json:"origin_size"`
	// This field is from variant [DirectoryVideo].
	OriginURL string `json:"origin_url"`
	// This field is from variant [DirectoryVideo].
	OriginVideoDuration int64 `json:"origin_video_duration"`
	// This field is from variant [DirectoryVideo].
	Poster string `json:"poster"`
	// This field is from variant [DirectoryVideo].
	PosterThumb string `json:"poster_thumb"`
	// This field is from variant [DirectoryVideo].
	Projection string `json:"projection"`
	// This field is from variant [DirectoryVideo].
	RecordingStartedAt string `json:"recording_started_at"`
	// This field is from variant [DirectoryVideo].
	Screenshot string `json:"screenshot"`
	// This field is from variant [DirectoryVideo].
	ScreenshotID int64 `json:"screenshot_id"`
	// This field is from variant [DirectoryVideo].
	Screenshots []string `json:"screenshots"`
	// This field is from variant [DirectoryVideo].
	ShareURL string `json:"share_url"`
	// This field is from variant [DirectoryVideo].
	Slug string `json:"slug"`
	// This field is from variant [DirectoryVideo].
	Sprite string `json:"sprite"`
	// This field is from variant [DirectoryVideo].
	SpriteVtt string `json:"sprite_vtt"`
	// This field is from variant [DirectoryVideo].
	Status VideoStatus `json:"status"`
	// This field is from variant [DirectoryVideo].
	StreamID int64 `json:"stream_id"`
	// This field is from variant [DirectoryVideo].
	Views    int64  `json:"views"`
	ItemType string `json:"item_type"`
	// This field is from variant [DirectoryItem].
	CreatedAt string `json:"created_at"`
	// This field is from variant [DirectoryItem].
	ItemsCount int64 `json:"items_count"`
	// This field is from variant [DirectoryItem].
	ParentID int64 `json:"parent_id"`
	// This field is from variant [DirectoryItem].
	UpdatedAt string `json:"updated_at"`
	JSON      struct {
		ID                  respjson.Field
		AdID                respjson.Field
		CDNViews            respjson.Field
		ClientID            respjson.Field
		ClientUserID        respjson.Field
		ConvertedVideos     respjson.Field
		CustomIframeURL     respjson.Field
		DashURL             respjson.Field
		Description         respjson.Field
		Duration            respjson.Field
		Error               respjson.Field
		HlsCmafURL          respjson.Field
		HlsURL              respjson.Field
		IframeURL           respjson.Field
		Name                respjson.Field
		OriginSize          respjson.Field
		OriginURL           respjson.Field
		OriginVideoDuration respjson.Field
		Poster              respjson.Field
		PosterThumb         respjson.Field
		Projection          respjson.Field
		RecordingStartedAt  respjson.Field
		Screenshot          respjson.Field
		ScreenshotID        respjson.Field
		Screenshots         respjson.Field
		ShareURL            respjson.Field
		Slug                respjson.Field
		Sprite              respjson.Field
		SpriteVtt           respjson.Field
		Status              respjson.Field
		StreamID            respjson.Field
		Views               respjson.Field
		ItemType            respjson.Field
		CreatedAt           respjson.Field
		ItemsCount          respjson.Field
		ParentID            respjson.Field
		UpdatedAt           respjson.Field
		raw                 string
	} `json:"-"`
}

func (u DirectoryGetResponseDirectoryItemUnion) AsDirectoryVideo() (v DirectoryVideo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DirectoryGetResponseDirectoryItemUnion) AsDirectoryItem() (v DirectoryItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DirectoryGetResponseDirectoryItemUnion) RawJSON() string { return u.JSON.raw }

func (r *DirectoryGetResponseDirectoryItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryNewParams struct {
	// Title of the directory.
	Name string `json:"name" api:"required"`
	// ID of a parent directory. "null" if it's in the root.
	ParentID param.Opt[int64] `json:"parent_id,omitzero"`
	paramObj
}

func (r DirectoryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DirectoryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirectoryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryUpdateParams struct {
	// Title of the directory. Omit this if you don't want to change.
	Name param.Opt[string] `json:"name,omitzero"`
	// ID of a parent directory. "null" if it's in the root. Omit this if you don't
	// want to change.
	ParentID param.Opt[int64] `json:"parent_id,omitzero"`
	paramObj
}

func (r DirectoryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DirectoryUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirectoryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
