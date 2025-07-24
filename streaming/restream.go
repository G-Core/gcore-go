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

// RestreamService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRestreamService] method instead.
type RestreamService struct {
	Options []option.RequestOption
}

// NewRestreamService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRestreamService(opts ...option.RequestOption) (r RestreamService) {
	r = RestreamService{}
	r.Options = opts
	return
}

// Create restream
func (r *RestreamService) New(ctx context.Context, body RestreamNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "streaming/restreams"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Updates restream settings
func (r *RestreamService) Update(ctx context.Context, restreamID int64, body RestreamUpdateParams, opts ...option.RequestOption) (res *Restream, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("streaming/restreams/%v", restreamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of created restreams
func (r *RestreamService) List(ctx context.Context, query RestreamListParams, opts ...option.RequestOption) (res *pagination.PageStreaming[Restream], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "streaming/restreams"
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

// Returns a list of created restreams
func (r *RestreamService) ListAutoPaging(ctx context.Context, query RestreamListParams, opts ...option.RequestOption) *pagination.PageStreamingAutoPager[Restream] {
	return pagination.NewPageStreamingAutoPager(r.List(ctx, query, opts...))
}

// Delete restream
func (r *RestreamService) Delete(ctx context.Context, restreamID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("streaming/restreams/%v", restreamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns restream details
func (r *RestreamService) Get(ctx context.Context, restreamID int64, opts ...option.RequestOption) (res *Restream, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("streaming/restreams/%v", restreamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Restream struct {
	// Enables/Disables restream. Has two possible values:
	//
	// - **true** — restream is enabled and can be started
	// - **false** — restream is disabled.
	//
	// Default is true
	Active bool `json:"active"`
	// Custom field where you can specify user ID in your system
	ClientUserID int64 `json:"client_user_id"`
	// Indicates that the stream is being published. Has two possible values:
	//
	// - **true** — stream is being published
	// - **false** — stream isn't published
	Live bool `json:"live"`
	// Restream name
	Name string `json:"name"`
	// ID of the stream to restream
	StreamID int64 `json:"stream_id"`
	// A URL to push the stream to
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active       respjson.Field
		ClientUserID respjson.Field
		Live         respjson.Field
		Name         respjson.Field
		StreamID     respjson.Field
		Uri          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Restream) RawJSON() string { return r.JSON.raw }
func (r *Restream) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestreamNewParams struct {
	Restream RestreamNewParamsRestream `json:"restream,omitzero"`
	paramObj
}

func (r RestreamNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RestreamNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestreamNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestreamNewParamsRestream struct {
	// Enables/Disables restream. Has two possible values:
	//
	// - **true** — restream is enabled and can be started
	// - **false** — restream is disabled.
	//
	// Default is true
	Active param.Opt[bool] `json:"active,omitzero"`
	// Custom field where you can specify user ID in your system
	ClientUserID param.Opt[int64] `json:"client_user_id,omitzero"`
	// Indicates that the stream is being published. Has two possible values:
	//
	// - **true** — stream is being published
	// - **false** — stream isn't published
	Live param.Opt[bool] `json:"live,omitzero"`
	// Restream name
	Name param.Opt[string] `json:"name,omitzero"`
	// ID of the stream to restream
	StreamID param.Opt[int64] `json:"stream_id,omitzero"`
	// A URL to push the stream to
	Uri param.Opt[string] `json:"uri,omitzero"`
	paramObj
}

func (r RestreamNewParamsRestream) MarshalJSON() (data []byte, err error) {
	type shadow RestreamNewParamsRestream
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestreamNewParamsRestream) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestreamUpdateParams struct {
	Restream RestreamUpdateParamsRestream `json:"restream,omitzero"`
	paramObj
}

func (r RestreamUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RestreamUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestreamUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestreamUpdateParamsRestream struct {
	// Enables/Disables restream. Has two possible values:
	//
	// - **true** — restream is enabled and can be started
	// - **false** — restream is disabled.
	//
	// Default is true
	Active param.Opt[bool] `json:"active,omitzero"`
	// Custom field where you can specify user ID in your system
	ClientUserID param.Opt[int64] `json:"client_user_id,omitzero"`
	// Indicates that the stream is being published. Has two possible values:
	//
	// - **true** — stream is being published
	// - **false** — stream isn't published
	Live param.Opt[bool] `json:"live,omitzero"`
	// Restream name
	Name param.Opt[string] `json:"name,omitzero"`
	// ID of the stream to restream
	StreamID param.Opt[int64] `json:"stream_id,omitzero"`
	// A URL to push the stream to
	Uri param.Opt[string] `json:"uri,omitzero"`
	paramObj
}

func (r RestreamUpdateParamsRestream) MarshalJSON() (data []byte, err error) {
	type shadow RestreamUpdateParamsRestream
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestreamUpdateParamsRestream) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestreamListParams struct {
	// Query parameter. Use it to list the paginated content
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RestreamListParams]'s query parameters as `url.Values`.
func (r RestreamListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
