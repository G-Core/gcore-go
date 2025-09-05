// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fastedge

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	shimjson "github.com/G-Core/gcore-go/internal/encoding/json"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// SecretService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecretService] method instead.
type SecretService struct {
	Options []option.RequestOption
}

// NewSecretService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSecretService(opts ...option.RequestOption) (r SecretService) {
	r = SecretService{}
	r.Options = opts
	return
}

// Add a new secret
func (r *SecretService) New(ctx context.Context, body SecretNewParams, opts ...option.RequestOption) (res *SecretNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "fastedge/v1/secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a secret
func (r *SecretService) Update(ctx context.Context, id int64, body SecretUpdateParams, opts ...option.RequestOption) (res *Secret, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("fastedge/v1/secrets/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List available secrets
func (r *SecretService) List(ctx context.Context, query SecretListParams, opts ...option.RequestOption) (res *SecretListResponse, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := "fastedge/v1/secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a secret
func (r *SecretService) Delete(ctx context.Context, id int64, body SecretDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("fastedge/v1/secrets/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Get secret by id
func (r *SecretService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *Secret, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("fastedge/v1/secrets/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a secret
func (r *SecretService) Replace(ctx context.Context, id int64, body SecretReplaceParams, opts ...option.RequestOption) (res *Secret, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("fastedge/v1/secrets/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type Secret struct {
	// The number of applications that use this secret.
	AppCount int64 `json:"app_count"`
	// A description or comment about the secret.
	Comment string `json:"comment"`
	// The unique name of the secret.
	Name string `json:"name"`
	// A list of secret slots associated with this secret.
	SecretSlots []SecretSecretSlot `json:"secret_slots"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppCount    respjson.Field
		Comment     respjson.Field
		Name        respjson.Field
		SecretSlots respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Secret) RawJSON() string { return r.JSON.raw }
func (r *Secret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Secret to a SecretParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SecretParam.Overrides()
func (r Secret) ToParam() SecretParam {
	return param.Override[SecretParam](json.RawMessage(r.RawJSON()))
}

type SecretSecretSlot struct {
	// Secret slot ID.
	Slot int64 `json:"slot,required"`
	// A checksum of the secret value for integrity verification.
	Checksum string `json:"checksum"`
	// The value of the secret.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Slot        respjson.Field
		Checksum    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecretSecretSlot) RawJSON() string { return r.JSON.raw }
func (r *SecretSecretSlot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecretParam struct {
	// The number of applications that use this secret.
	AppCount param.Opt[int64] `json:"app_count,omitzero"`
	// A description or comment about the secret.
	Comment param.Opt[string] `json:"comment,omitzero"`
	// The unique name of the secret.
	Name param.Opt[string] `json:"name,omitzero"`
	// A list of secret slots associated with this secret.
	SecretSlots []SecretSecretSlotParam `json:"secret_slots,omitzero"`
	paramObj
}

func (r SecretParam) MarshalJSON() (data []byte, err error) {
	type shadow SecretParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecretParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Slot is required.
type SecretSecretSlotParam struct {
	// Secret slot ID.
	Slot int64 `json:"slot,required"`
	// A checksum of the secret value for integrity verification.
	Checksum param.Opt[string] `json:"checksum,omitzero"`
	// The value of the secret.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r SecretSecretSlotParam) MarshalJSON() (data []byte, err error) {
	type shadow SecretSecretSlotParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecretSecretSlotParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecretShort struct {
	// The unique name of the secret.
	Name string `json:"name,required"`
	// The unique identifier of the secret.
	ID int64 `json:"id"`
	// The number of applications that use this secret.
	AppCount int64 `json:"app_count"`
	// A description or comment about the secret.
	Comment string `json:"comment"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ID          respjson.Field
		AppCount    respjson.Field
		Comment     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecretShort) RawJSON() string { return r.JSON.raw }
func (r *SecretShort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecretNewResponse struct {
	// The unique identifier of the secret.
	ID int64 `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Secret
}

// Returns the unmodified JSON received from the API
func (r SecretNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SecretNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecretListResponse struct {
	Secrets []SecretShort `json:"secrets,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Secrets     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SecretListResponse) RawJSON() string { return r.JSON.raw }
func (r *SecretListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecretNewParams struct {
	Secret SecretParam
	paramObj
}

func (r SecretNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Secret)
}
func (r *SecretNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Secret)
}

type SecretUpdateParams struct {
	Secret SecretParam
	paramObj
}

func (r SecretUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Secret)
}
func (r *SecretUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Secret)
}

type SecretListParams struct {
	// App ID
	AppID param.Opt[int64] `query:"app_id,omitzero" json:"-"`
	// Secret name
	SecretName param.Opt[string] `query:"secret_name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SecretListParams]'s query parameters as `url.Values`.
func (r SecretListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SecretDeleteParams struct {
	// Force delete secret even if it is used in slots
	Force param.Opt[bool] `query:"force,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SecretDeleteParams]'s query parameters as `url.Values`.
func (r SecretDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SecretReplaceParams struct {
	Secret SecretParam
	paramObj
}

func (r SecretReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Secret)
}
func (r *SecretReplaceParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Secret)
}
