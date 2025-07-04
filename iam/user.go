// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam

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

// UserService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// This method updates user's details.
func (r *UserService) Update(ctx context.Context, userID int64, body UserUpdateParams, opts ...option.RequestOption) (res *UserUpdate, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("iam/users/%v", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of users. Pass a value for the `limit` parameter in your request if
// you want retrieve a paginated result. Otherwise API returns a list with all
// users without pagination.
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *pagination.OffsetPageIam[User], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "iam/users"
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

// Get a list of users. Pass a value for the `limit` parameter in your request if
// you want retrieve a paginated result. Otherwise API returns a list with all
// users without pagination.
func (r *UserService) ListAutoPaging(ctx context.Context, query UserListParams, opts ...option.RequestOption) *pagination.OffsetPageIamAutoPager[User] {
	return pagination.NewOffsetPageIamAutoPager(r.List(ctx, query, opts...))
}

// Revokes user's access to the specified account. If the specified user doesn't
// have access to multiple accounts, the user is deleted.
func (r *UserService) Delete(ctx context.Context, userID int64, body UserDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("iam/clients/%v/client-users/%v", body.ClientID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get user's details
func (r *UserService) Get(ctx context.Context, userID int64, opts ...option.RequestOption) (res *UserDetailed, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("iam/users/%v", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Invite a user to the account. User will receive an email. The new user will
// receive an invitation email with a link to create an account password, the
// existing user will be notified about the invitation to the account.
func (r *UserService) Invite(ctx context.Context, body UserInviteParams, opts ...option.RequestOption) (res *UserInvite, err error) {
	opts = append(r.Options[:], opts...)
	path := "iam/clients/invite_user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type User struct {
	// User's ID.
	ID int64 `json:"id"`
	// Email confirmation:
	//
	// - `true` – user confirmed the email;
	// - `false` – user did not confirm the email.
	Activated bool `json:"activated"`
	// System field. List of auth types available for the account.
	//
	// Any of "password", "sso", "github", "google-oauth2".
	AuthTypes []string `json:"auth_types"`
	// User's account ID.
	Client float64 `json:"client"`
	// User's company.
	Company string `json:"company"`
	// Deletion flag. If `true` then user was deleted.
	Deleted bool `json:"deleted"`
	// User's email address.
	Email string `json:"email" format:"email"`
	// User's group in the current account. IAM supports 5 groups:
	//
	// - Users
	// - Administrators
	// - Engineers
	// - Purge and Prefetch only (API)
	// - Purge and Prefetch only (API+Web)
	Groups []UserGroup `json:"groups"`
	// User's language. Defines language of the control panel and email messages.
	//
	// Any of "de", "en", "ru", "zh", "az".
	Lang UserLang `json:"lang"`
	// User's name.
	Name string `json:"name,nullable"`
	// User's phone.
	Phone string `json:"phone,nullable"`
	// Services provider ID.
	Reseller int64 `json:"reseller"`
	// SSO authentication flag. If `true` then user can login via SAML SSO.
	SSOAuth bool `json:"sso_auth"`
	// Two-step verification:
	//
	// - `true` – user enabled two-step verification;
	// - `false` – user disabled two-step verification.
	TwoFa bool `json:"two_fa"`
	// User's type.
	//
	// Any of "common".
	UserType UserUserType `json:"user_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Activated   respjson.Field
		AuthTypes   respjson.Field
		Client      respjson.Field
		Company     respjson.Field
		Deleted     respjson.Field
		Email       respjson.Field
		Groups      respjson.Field
		Lang        respjson.Field
		Name        respjson.Field
		Phone       respjson.Field
		Reseller    respjson.Field
		SSOAuth     respjson.Field
		TwoFa       respjson.Field
		UserType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGroup struct {
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
func (r UserGroup) RawJSON() string { return r.JSON.raw }
func (r *UserGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User's language. Defines language of the control panel and email messages.
type UserLang string

const (
	UserLangDe UserLang = "de"
	UserLangEn UserLang = "en"
	UserLangRu UserLang = "ru"
	UserLangZh UserLang = "zh"
	UserLangAz UserLang = "az"
)

// User's type.
type UserUserType string

const (
	UserUserTypeCommon UserUserType = "common"
)

type UserDetailed struct {
	// User's ID.
	ID int64 `json:"id"`
	// Email confirmation:
	//
	// - `true` – user confirmed the email;
	// - `false` – user did not confirm the email.
	Activated bool `json:"activated"`
	// System field. List of auth types available for the account.
	//
	// Any of "password", "sso", "github", "google-oauth2".
	AuthTypes []string `json:"auth_types"`
	// User's account ID.
	Client float64 `json:"client"`
	// List of user's clients. User can access to one or more clients.
	ClientAndRoles []UserDetailedClientAndRole `json:"client_and_roles"`
	// User's company.
	Company string `json:"company"`
	// Deletion flag. If `true` then user was deleted.
	Deleted bool `json:"deleted"`
	// User's email address.
	Email string `json:"email" format:"email"`
	// User's group in the current account. IAM supports 5 groups:
	//
	// - Users
	// - Administrators
	// - Engineers
	// - Purge and Prefetch only (API)
	// - Purge and Prefetch only (API+Web)
	Groups []UserDetailedGroup `json:"groups"`
	// User activity flag.
	IsActive bool `json:"is_active"`
	// User's language. Defines language of the control panel and email messages.
	//
	// Any of "de", "en", "ru", "zh", "az".
	Lang UserDetailedLang `json:"lang"`
	// User's name.
	Name string `json:"name,nullable"`
	// User's phone.
	Phone string `json:"phone,nullable"`
	// Services provider ID.
	Reseller int64 `json:"reseller"`
	// SSO authentication flag. If `true` then user can login via SAML SSO.
	SSOAuth bool `json:"sso_auth"`
	// Two-step verification:
	//
	// - `true` – user enabled two-step verification;
	// - `false` – user disabled two-step verification.
	TwoFa bool `json:"two_fa"`
	// User's type.
	//
	// Any of "common", "reseller", "seller".
	UserType UserDetailedUserType `json:"user_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Activated      respjson.Field
		AuthTypes      respjson.Field
		Client         respjson.Field
		ClientAndRoles respjson.Field
		Company        respjson.Field
		Deleted        respjson.Field
		Email          respjson.Field
		Groups         respjson.Field
		IsActive       respjson.Field
		Lang           respjson.Field
		Name           respjson.Field
		Phone          respjson.Field
		Reseller       respjson.Field
		SSOAuth        respjson.Field
		TwoFa          respjson.Field
		UserType       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserDetailed) RawJSON() string { return r.JSON.raw }
func (r *UserDetailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserDetailedClientAndRole struct {
	ClientCompanyName string `json:"client_company_name,required"`
	ClientID          int64  `json:"client_id,required"`
	// User's ID.
	UserID int64 `json:"user_id,required"`
	// User role in this client.
	UserRoles []string `json:"user_roles,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientCompanyName respjson.Field
		ClientID          respjson.Field
		UserID            respjson.Field
		UserRoles         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserDetailedClientAndRole) RawJSON() string { return r.JSON.raw }
func (r *UserDetailedClientAndRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserDetailedGroup struct {
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
func (r UserDetailedGroup) RawJSON() string { return r.JSON.raw }
func (r *UserDetailedGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User's language. Defines language of the control panel and email messages.
type UserDetailedLang string

const (
	UserDetailedLangDe UserDetailedLang = "de"
	UserDetailedLangEn UserDetailedLang = "en"
	UserDetailedLangRu UserDetailedLang = "ru"
	UserDetailedLangZh UserDetailedLang = "zh"
	UserDetailedLangAz UserDetailedLang = "az"
)

// User's type.
type UserDetailedUserType string

const (
	UserDetailedUserTypeCommon   UserDetailedUserType = "common"
	UserDetailedUserTypeReseller UserDetailedUserType = "reseller"
	UserDetailedUserTypeSeller   UserDetailedUserType = "seller"
)

type UserInvite struct {
	// Status of the invitation.
	Status string `json:"status"`
	// Invited user ID.
	UserID int64 `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserInvite) RawJSON() string { return r.JSON.raw }
func (r *UserInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdate struct {
	// User's ID.
	ID int64 `json:"id"`
	// Email confirmation:
	//
	// - `true` – user confirmed the email;
	// - `false` – user did not confirm the email.
	Activated bool `json:"activated"`
	// System field. List of auth types available for the account.
	//
	// Any of "password", "sso", "github", "google-oauth2".
	AuthTypes []string `json:"auth_types"`
	// User's account ID.
	Client float64 `json:"client"`
	// List of user's clients. User can access to one or more clients.
	ClientAndRoles []UserUpdateClientAndRole `json:"client_and_roles"`
	// User's company.
	Company string `json:"company"`
	// Deletion flag. If `true` then user was deleted.
	Deleted bool `json:"deleted"`
	// User's email address.
	Email string `json:"email" format:"email"`
	// User's group in the current account. IAM supports 5 groups:
	//
	// - Users
	// - Administrators
	// - Engineers
	// - Purge and Prefetch only (API)
	// - Purge and Prefetch only (API+Web)
	Groups []UserUpdateGroup `json:"groups"`
	// User activity flag.
	IsActive bool `json:"is_active"`
	// User's language. Defines language of the control panel and email messages.
	//
	// Any of "de", "en", "ru", "zh", "az".
	Lang UserUpdateLang `json:"lang"`
	// User's name.
	Name string `json:"name,nullable"`
	// User's phone.
	Phone string `json:"phone,nullable"`
	// Services provider ID.
	Reseller int64 `json:"reseller"`
	// SSO authentication flag. If `true` then user can login via SAML SSO.
	SSOAuth bool `json:"sso_auth"`
	// Two-step verification:
	//
	// - `true` – user enabled two-step verification;
	// - `false` – user disabled two-step verification.
	TwoFa bool `json:"two_fa"`
	// User's type.
	//
	// Any of "common", "reseller", "seller".
	UserType UserUpdateUserType `json:"user_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Activated      respjson.Field
		AuthTypes      respjson.Field
		Client         respjson.Field
		ClientAndRoles respjson.Field
		Company        respjson.Field
		Deleted        respjson.Field
		Email          respjson.Field
		Groups         respjson.Field
		IsActive       respjson.Field
		Lang           respjson.Field
		Name           respjson.Field
		Phone          respjson.Field
		Reseller       respjson.Field
		SSOAuth        respjson.Field
		TwoFa          respjson.Field
		UserType       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserUpdate) RawJSON() string { return r.JSON.raw }
func (r *UserUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdateClientAndRole struct {
	ClientCompanyName string `json:"client_company_name,required"`
	ClientID          int64  `json:"client_id,required"`
	// User's ID.
	UserID int64 `json:"user_id,required"`
	// User role in this client.
	UserRoles []string `json:"user_roles,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientCompanyName respjson.Field
		ClientID          respjson.Field
		UserID            respjson.Field
		UserRoles         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserUpdateClientAndRole) RawJSON() string { return r.JSON.raw }
func (r *UserUpdateClientAndRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdateGroup struct {
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
func (r UserUpdateGroup) RawJSON() string { return r.JSON.raw }
func (r *UserUpdateGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User's language. Defines language of the control panel and email messages.
type UserUpdateLang string

const (
	UserUpdateLangDe UserUpdateLang = "de"
	UserUpdateLangEn UserUpdateLang = "en"
	UserUpdateLangRu UserUpdateLang = "ru"
	UserUpdateLangZh UserUpdateLang = "zh"
	UserUpdateLangAz UserUpdateLang = "az"
)

// User's type.
type UserUpdateUserType string

const (
	UserUpdateUserTypeCommon   UserUpdateUserType = "common"
	UserUpdateUserTypeReseller UserUpdateUserType = "reseller"
	UserUpdateUserTypeSeller   UserUpdateUserType = "seller"
)

type UserUpdateParams struct {
	// User's name.
	Name param.Opt[string] `json:"name,omitzero"`
	// User's phone.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// User's company.
	Company param.Opt[string] `json:"company,omitzero"`
	// User's email address.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// System field. List of auth types available for the account.
	//
	// Any of "password", "sso", "github", "google-oauth2".
	AuthTypes []string `json:"auth_types,omitzero"`
	// User's group in the current account. IAM supports 5 groups:
	//
	// - Users
	// - Administrators
	// - Engineers
	// - Purge and Prefetch only (API)
	// - Purge and Prefetch only (API+Web)
	Groups []UserUpdateParamsGroup `json:"groups,omitzero"`
	// User's language. Defines language of the control panel and email messages.
	//
	// Any of "de", "en", "ru", "zh", "az".
	Lang UserUpdateParamsLang `json:"lang,omitzero"`
	paramObj
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdateParamsGroup struct {
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

func (r UserUpdateParamsGroup) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParamsGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParamsGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UserUpdateParamsGroup](
		"name", "Users", "Administrators", "Engineers", "Purge and Prefetch only (API)", "Purge and Prefetch only (API+Web)",
	)
}

// User's language. Defines language of the control panel and email messages.
type UserUpdateParamsLang string

const (
	UserUpdateParamsLangDe UserUpdateParamsLang = "de"
	UserUpdateParamsLangEn UserUpdateParamsLang = "en"
	UserUpdateParamsLangRu UserUpdateParamsLang = "ru"
	UserUpdateParamsLangZh UserUpdateParamsLang = "zh"
	UserUpdateParamsLangAz UserUpdateParamsLang = "az"
)

type UserListParams struct {
	// The maximum number of items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset relative to the beginning of list.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type UserDeleteParams struct {
	ClientID int64 `path:"clientId,required" json:"-"`
	paramObj
}

type UserInviteParams struct {
	// ID of account.
	ClientID int64 `json:"client_id,required"`
	// User email.
	Email    string                   `json:"email,required" format:"email"`
	UserRole UserInviteParamsUserRole `json:"user_role,omitzero,required"`
	// User name.
	Name param.Opt[string] `json:"name,omitzero"`
	// User's language. Defines language of the control panel and email messages.
	//
	// Any of "de", "en", "ru", "zh", "az".
	Lang UserInviteParamsLang `json:"lang,omitzero"`
	paramObj
}

func (r UserInviteParams) MarshalJSON() (data []byte, err error) {
	type shadow UserInviteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserInviteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteParamsUserRole struct {
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

func (r UserInviteParamsUserRole) MarshalJSON() (data []byte, err error) {
	type shadow UserInviteParamsUserRole
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserInviteParamsUserRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UserInviteParamsUserRole](
		"name", "Users", "Administrators", "Engineers", "Purge and Prefetch only (API)", "Purge and Prefetch only (API+Web)",
	)
}

// User's language. Defines language of the control panel and email messages.
type UserInviteParamsLang string

const (
	UserInviteParamsLangDe UserInviteParamsLang = "de"
	UserInviteParamsLangEn UserInviteParamsLang = "en"
	UserInviteParamsLangRu UserInviteParamsLang = "ru"
	UserInviteParamsLangZh UserInviteParamsLang = "zh"
	UserInviteParamsLangAz UserInviteParamsLang = "az"
)
