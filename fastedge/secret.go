// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fastedge

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	shimjson "github.com/G-Core/gcore-go/internal/encoding/json"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// FastEdge secrets store sensitive values such as API keys and tokens that can be
// referenced by FastEdge applications.
//
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

// Create a new encrypted secret for use in edge applications. Secrets are
// encrypted with AES-256-GCM and can have multiple time-based slots for rotation.
func (r *SecretService) New(ctx context.Context, body SecretNewParams, opts ...option.RequestOption) (res *SecretNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fastedge/v1/secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Partially updates secret metadata and/or modifies specific slots
func (r *SecretService) Update(ctx context.Context, secretID int64, body SecretUpdateParams, opts ...option.RequestOption) (res *Secret, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("fastedge/v1/secrets/%v", secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieve encrypted secrets available to the authenticated client. Secrets can be
// filtered by application ID or name. Values are encrypted and require decryption.
func (r *SecretService) List(ctx context.Context, query SecretListParams, opts ...option.RequestOption) (res *SecretListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "fastedge/v1/secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently delete a secret and all its slot values. Secrets in use by
// applications require force=true to delete.
func (r *SecretService) Delete(ctx context.Context, secretID int64, body SecretDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("fastedge/v1/secrets/%v", secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Retrieve complete metadata for a specific secret including all time-based slots.
// Secret values remain encrypted; use the encryption service to decrypt when
// needed.
func (r *SecretService) Get(ctx context.Context, secretID int64, opts ...option.RequestOption) (res *Secret, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("fastedge/v1/secrets/%v", secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates secret metadata and/or adds new time-based slots with encrypted values
func (r *SecretService) Replace(ctx context.Context, secretID int64, body SecretReplaceParams, opts ...option.RequestOption) (res *Secret, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("fastedge/v1/secrets/%v", secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
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
	// Unix timestamp (seconds since epoch) indicating when this secret version becomes
	// active. Use for time-based secret rotation.
	Slot int64 `json:"slot" api:"required"`
	// SHA-256 hash of the decrypted value for integrity verification (auto-generated)
	Checksum string `json:"checksum"`
	// The plaintext secret value. Will be encrypted with AES-256-GCM before storage.
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
	// Unix timestamp (seconds since epoch) indicating when this secret version becomes
	// active. Use for time-based secret rotation.
	Slot int64 `json:"slot" api:"required"`
	// The plaintext secret value. Will be encrypted with AES-256-GCM before storage.
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
	Name string `json:"name" api:"required"`
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
	// Total number of secrets matching the filters
	Count   int64         `json:"count" api:"required"`
	Secrets []SecretShort `json:"secrets" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
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
	// When true, deletes secret even if used by applications. Defaults to false.
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
