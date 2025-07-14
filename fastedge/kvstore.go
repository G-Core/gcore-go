// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fastedge

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// KvStoreService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKvStoreService] method instead.
type KvStoreService struct {
	Options []option.RequestOption
}

// NewKvStoreService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewKvStoreService(opts ...option.RequestOption) (r KvStoreService) {
	r = KvStoreService{}
	r.Options = opts
	return
}

// Add a new KV store
func (r *KvStoreService) New(ctx context.Context, body KvStoreNewParams, opts ...option.RequestOption) (res *KvStore, err error) {
	opts = append(r.Options[:], opts...)
	path := "fastedge/v1/kv"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List available stores
func (r *KvStoreService) List(ctx context.Context, query KvStoreListParams, opts ...option.RequestOption) (res *KvStoreListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "fastedge/v1/kv"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a store
func (r *KvStoreService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("fastedge/v1/kv/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get store by id
func (r *KvStoreService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *KvStoreGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("fastedge/v1/kv/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a store
func (r *KvStoreService) Replace(ctx context.Context, id int64, body KvStoreReplaceParams, opts ...option.RequestOption) (res *KvStore, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("fastedge/v1/kv/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type KvStore struct {
	// The unique identifier of the store
	ID int64 `json:"id"`
	// The number of applications that use this store
	AppCount int64 `json:"app_count"`
	// BYOD (Bring Your Own Data) settings
	Byod KvStoreByod `json:"byod"`
	// A description of the store
	Comment string `json:"comment"`
	// Last update time
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AppCount    respjson.Field
		Byod        respjson.Field
		Comment     respjson.Field
		Updated     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KvStore) RawJSON() string { return r.JSON.raw }
func (r *KvStore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this KvStore to a KvStoreParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// KvStoreParam.Overrides()
func (r KvStore) ToParam() KvStoreParam {
	return param.Override[KvStoreParam](json.RawMessage(r.RawJSON()))
}

// BYOD (Bring Your Own Data) settings
type KvStoreByod struct {
	// Key prefix
	Prefix string `json:"prefix,required"`
	// URL to connect to
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Prefix      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KvStoreByod) RawJSON() string { return r.JSON.raw }
func (r *KvStoreByod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KvStoreParam struct {
	// The unique identifier of the store
	ID param.Opt[int64] `json:"id,omitzero"`
	// The number of applications that use this store
	AppCount param.Opt[int64] `json:"app_count,omitzero"`
	// A description of the store
	Comment param.Opt[string] `json:"comment,omitzero"`
	// Last update time
	Updated param.Opt[time.Time] `json:"updated,omitzero" format:"date-time"`
	// BYOD (Bring Your Own Data) settings
	Byod KvStoreByodParam `json:"byod,omitzero"`
	paramObj
}

func (r KvStoreParam) MarshalJSON() (data []byte, err error) {
	type shadow KvStoreParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KvStoreParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BYOD (Bring Your Own Data) settings
//
// The properties Prefix, URL are required.
type KvStoreByodParam struct {
	// Key prefix
	Prefix string `json:"prefix,required"`
	// URL to connect to
	URL string `json:"url,required"`
	paramObj
}

func (r KvStoreByodParam) MarshalJSON() (data []byte, err error) {
	type shadow KvStoreByodParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KvStoreByodParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KvStoreShort struct {
	// The unique identifier of the store
	ID int64 `json:"id"`
	// A description of the store
	Comment string `json:"comment"`
	// Last update time
	Updated time.Time `json:"updated" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Comment     respjson.Field
		Updated     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KvStoreShort) RawJSON() string { return r.JSON.raw }
func (r *KvStoreShort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KvStoreStats struct {
	// Store statistics
	Stats KvStoreStatsStats `json:"stats"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Stats       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KvStoreStats) RawJSON() string { return r.JSON.raw }
func (r *KvStoreStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Store statistics
type KvStoreStatsStats struct {
	// Total number of Cuckoo filter entries
	CfCount int64 `json:"cf_count,required"`
	// Total number of KV entries
	KvCount int64 `json:"kv_count,required"`
	// Total store size in bytes
	Size int64 `json:"size,required"`
	// Total number of sorted set entries
	ZsetCount int64 `json:"zset_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CfCount     respjson.Field
		KvCount     respjson.Field
		Size        respjson.Field
		ZsetCount   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KvStoreStatsStats) RawJSON() string { return r.JSON.raw }
func (r *KvStoreStatsStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KvStoreListResponse struct {
	// Total number of stores
	Count  int64          `json:"count,required"`
	Stores []KvStoreShort `json:"stores"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Stores      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KvStoreListResponse) RawJSON() string { return r.JSON.raw }
func (r *KvStoreListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KvStoreGetResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	KvStore
	KvStoreStats
}

// Returns the unmodified JSON received from the API
func (r KvStoreGetResponse) RawJSON() string { return r.JSON.raw }
func (r *KvStoreGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KvStoreNewParams struct {
	KvStore KvStoreParam
	paramObj
}

func (r KvStoreNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.KvStore)
}
func (r *KvStoreNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.KvStore)
}

type KvStoreListParams struct {
	// App ID
	AppID param.Opt[int64] `query:"app_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [KvStoreListParams]'s query parameters as `url.Values`.
func (r KvStoreListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type KvStoreReplaceParams struct {
	KvStore KvStoreParam
	paramObj
}

func (r KvStoreReplaceParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.KvStore)
}
func (r *KvStoreReplaceParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.KvStore)
}
