// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package streaming

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// BroadcastService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBroadcastService] method instead.
type BroadcastService struct {
	Options []option.RequestOption
}

// NewBroadcastService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBroadcastService(opts ...option.RequestOption) (r BroadcastService) {
	r = BroadcastService{}
	r.Options = opts
	return
}

// Broadcast entity is for setting up HTML video player, which serves to combine:
//
//   - many live streams,
//   - advertising,
//   - and design in one config. If you use other players or you get streams by
//     direct .m3u8/.mpd links, then you will not need this entity.
//
// Scheme of "broadcast" entity using:
// ![Scheme of "broadcast" using](https://demo-files.gvideo.io/apidocs/broadcasts.png)
func (r *BroadcastService) New(ctx context.Context, body BroadcastNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/broadcasts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Updates broadcast settings
func (r *BroadcastService) Update(ctx context.Context, broadcastID int64, body BroadcastUpdateParams, opts ...option.RequestOption) (res *Broadcast, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("streaming/broadcasts/%v", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Note: Feature "Broadcast" is outdated, soon it will be replaced by
// "Multicamera".
//
// Returns a list of broadcasts. Please see description in POST method.
func (r *BroadcastService) List(ctx context.Context, query BroadcastListParams, opts ...option.RequestOption) (res *pagination.PageStreaming[Broadcast], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "streaming/broadcasts"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Note: Feature "Broadcast" is outdated, soon it will be replaced by
// "Multicamera".
//
// Returns a list of broadcasts. Please see description in POST method.
func (r *BroadcastService) ListAutoPaging(ctx context.Context, query BroadcastListParams, opts ...option.RequestOption) *pagination.PageStreamingAutoPager[Broadcast] {
	return pagination.NewPageStreamingAutoPager(r.List(ctx, query, opts...))
}

// Delete broadcast
func (r *BroadcastService) Delete(ctx context.Context, broadcastID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("streaming/broadcasts/%v", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns broadcast details
func (r *BroadcastService) Get(ctx context.Context, broadcastID int64, opts ...option.RequestOption) (res *Broadcast, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("streaming/broadcasts/%v", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns number of simultaneous broadcast viewers at the current moment
func (r *BroadcastService) GetSpectatorsCount(ctx context.Context, broadcastID int64, opts ...option.RequestOption) (res *BroadcastSpectatorsCount, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("streaming/broadcasts/%v/spectators", broadcastID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Broadcast struct {
	// Broadcast name
	Name string `json:"name,required"`
	// ID of ad to be displayed in a live stream. If empty the default ad is show. If
	// there is no default ad, no ad is shown
	AdID int64 `json:"ad_id"`
	// Custom URL of iframe for video player to be shared via sharing button in player.
	// Auto generated iframe URL is provided by default
	CustomIframeURL string `json:"custom_iframe_url"`
	// A custom message that is shown if broadcast status is set to pending. If empty,
	// a default message is shown
	PendingMessage string `json:"pending_message"`
	// ID of player to be used with a broadcast. If empty the default player is used
	PlayerID int64 `json:"player_id"`
	// Uploaded poster file
	Poster string `json:"poster"`
	// Custom URL or iframe displayed in the link field when a user clicks on a sharing
	// button in player. If empty, the link field and social network sharing is
	// disabled
	ShareURL string `json:"share_url"`
	// Regulates if a DVR record is shown once a broadcast is finished. Has two
	// possible values:
	//
	// - **true** — record is shown
	// - **false** — record isn't shown
	//
	// Default is false
	ShowDvrAfterFinish bool `json:"show_dvr_after_finish"`
	// Broadcast statuses:
	//
	//	**Pending** — default “Broadcast isn’t started yet” or custom message (see `pending_message`
	//
	// parameter) is shown, users don't see the live stream
	//
	//	**Live** — broadcast is live, and viewers can see it
	//	**Paused** — “Broadcast is paused” message is shown, users don't see the live stream
	//
	// **Finished** — “Broadcast is finished” message is shown, users don't see the
	// live stream
	//
	//	The users' browsers start displaying the message/stream immediately after you change
	//
	// the broadcast status
	Status string `json:"status"`
	// IDs of streams used in a broadcast
	StreamIDs []int64 `json:"stream_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name               respjson.Field
		AdID               respjson.Field
		CustomIframeURL    respjson.Field
		PendingMessage     respjson.Field
		PlayerID           respjson.Field
		Poster             respjson.Field
		ShareURL           respjson.Field
		ShowDvrAfterFinish respjson.Field
		Status             respjson.Field
		StreamIDs          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Broadcast) RawJSON() string { return r.JSON.raw }
func (r *Broadcast) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastSpectatorsCount struct {
	// Number of spectators at the moment
	SpectatorsCount int64 `json:"spectators_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SpectatorsCount respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastSpectatorsCount) RawJSON() string { return r.JSON.raw }
func (r *BroadcastSpectatorsCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastNewParams struct {
	Broadcast BroadcastNewParamsBroadcast `json:"broadcast,omitzero"`
	paramObj
}

func (r BroadcastNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type BroadcastNewParamsBroadcast struct {
	// Broadcast name
	Name string `json:"name,required"`
	// ID of ad to be displayed in a live stream. If empty the default ad is show. If
	// there is no default ad, no ad is shown
	AdID param.Opt[int64] `json:"ad_id,omitzero"`
	// Custom URL of iframe for video player to be shared via sharing button in player.
	// Auto generated iframe URL is provided by default
	CustomIframeURL param.Opt[string] `json:"custom_iframe_url,omitzero"`
	// A custom message that is shown if broadcast status is set to pending. If empty,
	// a default message is shown
	PendingMessage param.Opt[string] `json:"pending_message,omitzero"`
	// ID of player to be used with a broadcast. If empty the default player is used
	PlayerID param.Opt[int64] `json:"player_id,omitzero"`
	// Uploaded poster file
	Poster param.Opt[string] `json:"poster,omitzero"`
	// Custom URL or iframe displayed in the link field when a user clicks on a sharing
	// button in player. If empty, the link field and social network sharing is
	// disabled
	ShareURL param.Opt[string] `json:"share_url,omitzero"`
	// Regulates if a DVR record is shown once a broadcast is finished. Has two
	// possible values:
	//
	// - **true** — record is shown
	// - **false** — record isn't shown
	//
	// Default is false
	ShowDvrAfterFinish param.Opt[bool] `json:"show_dvr_after_finish,omitzero"`
	// Broadcast statuses:
	//
	//	**Pending** — default “Broadcast isn’t started yet” or custom message (see `pending_message`
	//
	// parameter) is shown, users don't see the live stream
	//
	//	**Live** — broadcast is live, and viewers can see it
	//	**Paused** — “Broadcast is paused” message is shown, users don't see the live stream
	//
	// **Finished** — “Broadcast is finished” message is shown, users don't see the
	// live stream
	//
	//	The users' browsers start displaying the message/stream immediately after you change
	//
	// the broadcast status
	Status param.Opt[string] `json:"status,omitzero"`
	// IDs of streams used in a broadcast
	StreamIDs []int64 `json:"stream_ids,omitzero"`
	paramObj
}

func (r BroadcastNewParamsBroadcast) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastNewParamsBroadcast
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastNewParamsBroadcast) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastUpdateParams struct {
	Broadcast BroadcastUpdateParamsBroadcast `json:"broadcast,omitzero"`
	paramObj
}

func (r BroadcastUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type BroadcastUpdateParamsBroadcast struct {
	// Broadcast name
	Name string `json:"name,required"`
	// ID of ad to be displayed in a live stream. If empty the default ad is show. If
	// there is no default ad, no ad is shown
	AdID param.Opt[int64] `json:"ad_id,omitzero"`
	// Custom URL of iframe for video player to be shared via sharing button in player.
	// Auto generated iframe URL is provided by default
	CustomIframeURL param.Opt[string] `json:"custom_iframe_url,omitzero"`
	// A custom message that is shown if broadcast status is set to pending. If empty,
	// a default message is shown
	PendingMessage param.Opt[string] `json:"pending_message,omitzero"`
	// ID of player to be used with a broadcast. If empty the default player is used
	PlayerID param.Opt[int64] `json:"player_id,omitzero"`
	// Uploaded poster file
	Poster param.Opt[string] `json:"poster,omitzero"`
	// Custom URL or iframe displayed in the link field when a user clicks on a sharing
	// button in player. If empty, the link field and social network sharing is
	// disabled
	ShareURL param.Opt[string] `json:"share_url,omitzero"`
	// Regulates if a DVR record is shown once a broadcast is finished. Has two
	// possible values:
	//
	// - **true** — record is shown
	// - **false** — record isn't shown
	//
	// Default is false
	ShowDvrAfterFinish param.Opt[bool] `json:"show_dvr_after_finish,omitzero"`
	// Broadcast statuses:
	//
	//	**Pending** — default “Broadcast isn’t started yet” or custom message (see `pending_message`
	//
	// parameter) is shown, users don't see the live stream
	//
	//	**Live** — broadcast is live, and viewers can see it
	//	**Paused** — “Broadcast is paused” message is shown, users don't see the live stream
	//
	// **Finished** — “Broadcast is finished” message is shown, users don't see the
	// live stream
	//
	//	The users' browsers start displaying the message/stream immediately after you change
	//
	// the broadcast status
	Status param.Opt[string] `json:"status,omitzero"`
	// IDs of streams used in a broadcast
	StreamIDs []int64 `json:"stream_ids,omitzero"`
	paramObj
}

func (r BroadcastUpdateParamsBroadcast) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastUpdateParamsBroadcast
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastUpdateParamsBroadcast) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastListParams struct {
	// Query parameter. Use it to list the paginated content
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BroadcastListParams]'s query parameters as `url.Values`.
func (r BroadcastListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
