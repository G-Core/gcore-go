// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// APITokenService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPITokenService] method instead.
type APITokenService struct {
	Options []option.RequestOption
}

// NewAPITokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPITokenService(opts ...option.RequestOption) (r APITokenService) {
	r = APITokenService{}
	r.Options = opts
	return
}

// Create an API token in the current account.
func (r *APITokenService) New(ctx context.Context, clientID int64, body APITokenNewParams, opts ...option.RequestOption) (res *APITokenCreate, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("iam/clients/%v/tokens", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get information about your permanent API tokens in the account. A user with the
// Administrators role gets information about all API tokens in the account.
func (r *APITokenService) List(ctx context.Context, clientID int64, query APITokenListParams, opts ...option.RequestOption) (res *APITokenList, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("iam/clients/%v/tokens", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete API token from current account. Ensure that the API token is not being
// used by an active application. After deleting the token, all applications that
// use this token will not be able to get access to your account via API. The
// action cannot be reversed.
func (r *APITokenService) Delete(ctx context.Context, tokenID int64, body APITokenDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("iam/clients/%v/tokens/%v", body.ClientID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get API Token
func (r *APITokenService) Get(ctx context.Context, tokenID int64, query APITokenGetParams, opts ...option.RequestOption) (res *APIToken, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("iam/clients/%v/tokens/%v", query.ClientID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIToken struct {
	ClientUser APITokenClientUser `json:"client_user,required"`
	// Date when the API token becomes expired (ISO 8086/RFC 3339 format), UTC. If
	// null, then the API token will never expire.
	ExpDate string `json:"exp_date,required"`
	// API token name.
	Name string `json:"name,required"`
	// API token ID.
	ID int64 `json:"id"`
	// Date when the API token was issued (ISO 8086/RFC 3339 format), UTC.
	Created string `json:"created"`
	// Deletion flag. If true, then the API token was deleted.
	Deleted bool `json:"deleted"`
	// API token description.
	Description string `json:"description"`
	// Expiration flag. If true, then the API token has expired. When an API token
	// expires it will be automatically deleted.
	Expired bool `json:"expired"`
	// Date when the API token was last used (ISO 8086/RFC 3339 format), UTC.
	LastUsage string `json:"last_usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientUser  respjson.Field
		ExpDate     respjson.Field
		Name        respjson.Field
		ID          respjson.Field
		Created     respjson.Field
		Deleted     respjson.Field
		Description respjson.Field
		Expired     respjson.Field
		LastUsage   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIToken) RawJSON() string { return r.JSON.raw }
func (r *APIToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APITokenClientUser struct {
	// Account's ID.
	ClientID int64 `json:"client_id"`
	// Deletion flag. If true, then the API token was deleted.
	Deleted bool                   `json:"deleted"`
	Role    APITokenClientUserRole `json:"role"`
	// User's email who issued the API token.
	UserEmail string `json:"user_email"`
	// User's ID who issued the API token.
	UserID int64 `json:"user_id"`
	// User's name who issued the API token.
	UserName string `json:"user_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID    respjson.Field
		Deleted     respjson.Field
		Role        respjson.Field
		UserEmail   respjson.Field
		UserID      respjson.Field
		UserName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APITokenClientUser) RawJSON() string { return r.JSON.raw }
func (r *APITokenClientUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APITokenClientUserRole struct {
	// Group's ID: Possible values are:
	//
	//   - 1 - Administrators* 2 - Users* 5 - Engineers* 3009 - Purge and Prefetch only
	//     (API+Web)* 3022 - Purge and Prefetch only (API)
	ID int64 `json:"id"`
	// Group's name.
	//
	// Any of "Users", "Administrators", "Engineers", "Purge and Prefetch only (API)",
	// "Purge and Prefetch only (API+Web)".
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APITokenClientUserRole) RawJSON() string { return r.JSON.raw }
func (r *APITokenClientUserRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APITokenCreate struct {
	// API token. Copy it, because you will not be able to get it again. We do not
	// store tokens. All responsibility for token storage and usage is on the issuer.
	Token string `json:"token"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APITokenCreate) RawJSON() string { return r.JSON.raw }
func (r *APITokenCreate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APITokenList []APITokenListItem

type APITokenListItem struct {
	ClientUser APITokenListItemClientUser `json:"client_user,required"`
	// Date when the API token becomes expired (ISO 8086/RFC 3339 format), UTC. If
	// null, then the API token will never expire.
	ExpDate string `json:"exp_date,required"`
	// API token name.
	Name string `json:"name,required"`
	// API token ID.
	ID int64 `json:"id"`
	// Date when the API token was issued (ISO 8086/RFC 3339 format), UTC.
	Created string `json:"created"`
	// Deletion flag. If true, then the API token was deleted.
	Deleted bool `json:"deleted"`
	// API token description.
	Description string `json:"description"`
	// Expiration flag. If true, then the API token has expired. When an API token
	// expires it will be automatically deleted.
	Expired bool `json:"expired"`
	// Date when the API token was last used (ISO 8086/RFC 3339 format), UTC.
	LastUsage string `json:"last_usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientUser  respjson.Field
		ExpDate     respjson.Field
		Name        respjson.Field
		ID          respjson.Field
		Created     respjson.Field
		Deleted     respjson.Field
		Description respjson.Field
		Expired     respjson.Field
		LastUsage   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APITokenListItem) RawJSON() string { return r.JSON.raw }
func (r *APITokenListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APITokenListItemClientUser struct {
	// Account's ID.
	ClientID int64 `json:"client_id"`
	// Deletion flag. If true, then the API token was deleted.
	Deleted bool                           `json:"deleted"`
	Role    APITokenListItemClientUserRole `json:"role"`
	// User's email who issued the API token.
	UserEmail string `json:"user_email"`
	// User's ID who issued the API token.
	UserID int64 `json:"user_id"`
	// User's name who issued the API token.
	UserName string `json:"user_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID    respjson.Field
		Deleted     respjson.Field
		Role        respjson.Field
		UserEmail   respjson.Field
		UserID      respjson.Field
		UserName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APITokenListItemClientUser) RawJSON() string { return r.JSON.raw }
func (r *APITokenListItemClientUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APITokenListItemClientUserRole struct {
	// Group's ID: Possible values are:
	//
	//   - 1 - Administrators* 2 - Users* 5 - Engineers* 3009 - Purge and Prefetch only
	//     (API+Web)* 3022 - Purge and Prefetch only (API)
	ID int64 `json:"id"`
	// Group's name.
	//
	// Any of "Users", "Administrators", "Engineers", "Purge and Prefetch only (API)",
	// "Purge and Prefetch only (API+Web)".
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APITokenListItemClientUserRole) RawJSON() string { return r.JSON.raw }
func (r *APITokenListItemClientUserRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APITokenNewParams struct {
	// API token role.
	ClientUser APITokenNewParamsClientUser `json:"client_user,omitzero,required"`
	// Date when the API token becomes expired (ISO 8086/RFC 3339 format), UTC. If
	// null, then the API token will never expire.
	ExpDate string `json:"exp_date,required"`
	// API token name.
	Name string `json:"name,required"`
	// API token description.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r APITokenNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APITokenNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APITokenNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API token role.
type APITokenNewParamsClientUser struct {
	Role APITokenNewParamsClientUserRole `json:"role,omitzero"`
	paramObj
}

func (r APITokenNewParamsClientUser) MarshalJSON() (data []byte, err error) {
	type shadow APITokenNewParamsClientUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APITokenNewParamsClientUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APITokenNewParamsClientUserRole struct {
	// Group's ID: Possible values are:
	//
	//   - 1 - Administrators* 2 - Users* 5 - Engineers* 3009 - Purge and Prefetch only
	//     (API+Web)* 3022 - Purge and Prefetch only (API)
	ID param.Opt[int64] `json:"id,omitzero"`
	// Group's name.
	//
	// Any of "Users", "Administrators", "Engineers", "Purge and Prefetch only (API)",
	// "Purge and Prefetch only (API+Web)".
	Name string `json:"name,omitzero"`
	paramObj
}

func (r APITokenNewParamsClientUserRole) MarshalJSON() (data []byte, err error) {
	type shadow APITokenNewParamsClientUserRole
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APITokenNewParamsClientUserRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[APITokenNewParamsClientUserRole](
		"name", "Users", "Administrators", "Engineers", "Purge and Prefetch only (API)", "Purge and Prefetch only (API+Web)",
	)
}

type APITokenListParams struct {
	// The state of API tokens included in the response.
	//
	//	Two possible values:
	//
	// - True - API token was not deleted.\* False - API token was deleted.
	//
	// Example, _&deleted=True_
	Deleted param.Opt[bool] `query:"deleted,omitzero" json:"-"`
	// User's ID. Use to get API tokens issued by a particular user.
	//
	//	Example, _&`issued_by`=1234_
	IssuedBy param.Opt[int64] `query:"issued_by,omitzero" json:"-"`
	// User's ID. Use to get API tokens issued by anyone except a particular user.
	//
	//	Example, _¬_issued_by=1234_
	NotIssuedBy param.Opt[int64] `query:"not_issued_by,omitzero" json:"-"`
	// Group's ID. Possible values are:
	//
	//   - 1 - Administrators* 2 - Users* 5 - Engineers* 3009 - Purge and Prefetch only
	//     (API+Web)* 3022 - Purge and Prefetch only (API)
	//
	// Example, _&role=Engineers_
	Role param.Opt[string] `query:"role,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APITokenListParams]'s query parameters as `url.Values`.
func (r APITokenListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type APITokenDeleteParams struct {
	ClientID int64 `path:"clientId,required" json:"-"`
	paramObj
}

type APITokenGetParams struct {
	ClientID int64 `path:"clientId,required" json:"-"`
	paramObj
}
