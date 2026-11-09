// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

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
func (r *UserService) Update(ctx context.Context, userID int64, body UserUpdateParams, opts ...option.RequestOption) (res *UserUpdated, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("iam/users/%v", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List the users of the caller's own client. Replaces the deprecated `GET /users`.
//
// Filtering follows the standard query filtering rules: `ids` and `langs` are
// repeated flat parameters (OR), e.g. `ids=1&ids=2`; `email`, `name` and
// `created_at` require a bracket operator, e.g. `email[contains]=example.com`.
// Sorting is `order_by=field.direction` (e.g. `order_by=created_at.desc`, the
// default); results are paginated with `limit`/`offset`.
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[User], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "iam/v2/users"
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

// List the users of the caller's own client. Replaces the deprecated `GET /users`.
//
// Filtering follows the standard query filtering rules: `ids` and `langs` are
// repeated flat parameters (OR), e.g. `ids=1&ids=2`; `email`, `name` and
// `created_at` require a bracket operator, e.g. `email[contains]=example.com`.
// Sorting is `order_by=field.direction` (e.g. `order_by=created_at.desc`, the
// default); results are paginated with `limit`/`offset`.
func (r *UserService) ListAutoPaging(ctx context.Context, query UserListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[User] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Revokes user's access to the specified account. If the specified user doesn't
// have access to multiple accounts, the user is deleted.
func (r *UserService) Delete(ctx context.Context, userID int64, body UserDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("iam/clients/%v/client-users/%v", body.ClientID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get a single user of the caller's own client. Replaces the deprecated
// `GET /users/{userId}`.
func (r *UserService) Get(ctx context.Context, userID int64, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("iam/v2/users/%v", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The authenticated user's own profile, including the clients they belong to.
func (r *UserService) GetMe(ctx context.Context, opts ...option.RequestOption) (res *UserMe, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "iam/v2/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Invite a user to the account.
//
// User will receive an email. The new user will receive an invitation email with a
// link to create an account password, the existing user will be notified about the
// invitation to the account.
func (r *UserService) Invite(ctx context.Context, body UserInviteParams, opts ...option.RequestOption) (res *UserInvited, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "iam/clients/invite_user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fixed field set for /v2/users.
type User struct {
	ID int64 `json:"id" api:"required"`
	// Whether the user's membership in this client is activated.
	Activated bool `json:"activated" api:"required"`
	// List of authentication methods available to the user (e.g., 'password', 'sso').
	AuthTypes []string  `json:"auth_types" api:"required"`
	Client    int64     `json:"client" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// User's email address used for login and notifications.
	Email  string        `json:"email" api:"required" format:"email"`
	Groups []UserGroupV2 `json:"groups" api:"required"`
	// Full name of the user.
	Name string `json:"name" api:"required"`
	// Indicates whether two-factor authentication (TOTP) is enabled for the user.
	TwoFa    bool   `json:"two_fa" api:"required"`
	UserType string `json:"user_type" api:"required"`
	// Any of "az", "de", "en", "ru", "zh".
	Lang UserLang `json:"lang"`
	// User phone
	Phone string `json:"phone" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Activated   respjson.Field
		AuthTypes   respjson.Field
		Client      respjson.Field
		CreatedAt   respjson.Field
		Email       respjson.Field
		Groups      respjson.Field
		Name        respjson.Field
		TwoFa       respjson.Field
		UserType    respjson.Field
		Lang        respjson.Field
		Phone       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGroupV2 struct {
	ID   int64  `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGroupV2) RawJSON() string { return r.JSON.raw }
func (r *UserGroupV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserInvited struct {
	// Status of the invitation.
	Status string `json:"status" api:"required"`
	// Invited user ID.
	UserID int64 `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserInvited) RawJSON() string { return r.JSON.raw }
func (r *UserInvited) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserLang string

const (
	UserLangAz UserLang = "az"
	UserLangDe UserLang = "de"
	UserLangEn UserLang = "en"
	UserLangRu UserLang = "ru"
	UserLangZh UserLang = "zh"
)

// The client-tier caller's own profile, including the clients they belong to.
type UserMe struct {
	ID int64 `json:"id" api:"required"`
	// Whether the user's membership in this client is activated.
	Activated bool `json:"activated" api:"required"`
	// List of authentication methods available to the user (e.g., 'password', 'sso').
	AuthTypes []string `json:"auth_types" api:"required"`
	// False when the user cannot disable their own 2FA because they are an active
	// member of an account that enforces 2FA. Use it to disable the Disable-2FA
	// control.
	CanDisable2fa   bool                   `json:"can_disable_2fa" api:"required"`
	Client          int64                  `json:"client" api:"required"`
	ClientsAndRoles []UserMeClientsAndRole `json:"clients_and_roles" api:"required"`
	CreatedAt       time.Time              `json:"created_at" api:"required" format:"date-time"`
	// User's email address used for login and notifications.
	Email  string        `json:"email" api:"required" format:"email"`
	Groups []UserGroupV2 `json:"groups" api:"required"`
	// Full name of the user.
	Name string `json:"name" api:"required"`
	// Whether the support chat widget should be shown.
	ShowChat bool `json:"show_chat" api:"required"`
	// Indicates whether two-factor authentication (TOTP) is enabled for the user.
	TwoFa    bool   `json:"two_fa" api:"required"`
	UserType string `json:"user_type" api:"required"`
	// Any of "az", "de", "en", "ru", "zh".
	Lang UserLang `json:"lang"`
	// User phone
	Phone string `json:"phone" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Activated       respjson.Field
		AuthTypes       respjson.Field
		CanDisable2fa   respjson.Field
		Client          respjson.Field
		ClientsAndRoles respjson.Field
		CreatedAt       respjson.Field
		Email           respjson.Field
		Groups          respjson.Field
		Name            respjson.Field
		ShowChat        respjson.Field
		TwoFa           respjson.Field
		UserType        respjson.Field
		Lang            respjson.Field
		Phone           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserMe) RawJSON() string { return r.JSON.raw }
func (r *UserMe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserMeClientsAndRole struct {
	AuthTypes         []string `json:"auth_types" api:"required"`
	ClientCompanyName string   `json:"client_company_name" api:"required"`
	ClientID          int64    `json:"client_id" api:"required"`
	UserID            int64    `json:"user_id" api:"required"`
	UserRoles         []string `json:"user_roles" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthTypes         respjson.Field
		ClientCompanyName respjson.Field
		ClientID          respjson.Field
		UserID            respjson.Field
		UserRoles         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserMeClientsAndRole) RawJSON() string { return r.JSON.raw }
func (r *UserMeClientsAndRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User's type.
type UserType string

const (
	UserTypeCommon   UserType = "common"
	UserTypeReseller UserType = "reseller"
	UserTypeSeller   UserType = "seller"
)

type UserUpdated struct {
	// User's ID.
	ID int64 `json:"id" api:"required"`
	// Email confirmation:
	//
	// - `true` – user confirmed the email;
	// - `false` – user did not confirm the email.
	Activated bool `json:"activated" api:"required"`
	// System field. List of auth types available for the account.
	AuthTypes []AuthType `json:"auth_types" api:"required"`
	// User's account ID.
	Client float64 `json:"client" api:"required"`
	// List of user's clients. User can access to one or more clients.
	ClientAndRoles []UserUpdatedClientAndRole `json:"client_and_roles" api:"required"`
	// User's company.
	Company string `json:"company" api:"required"`
	// Deletion flag. If `true` then user was deleted.
	Deleted bool `json:"deleted" api:"required"`
	// User's email address.
	Email string `json:"email" api:"required" format:"email"`
	// User's group in the current account.
	//
	// IAM supports 5 groups:
	//
	// - Users
	// - Administrators
	// - Engineers
	// - Purge and Prefetch only (API)
	// - Purge and Prefetch only (API+Web)
	Groups []UserGroup `json:"groups" api:"required"`
	// User activity flag.
	IsActive bool `json:"is_active" api:"required"`
	// User's language.
	//
	// Defines language of the control panel and email messages.
	//
	// Any of "de", "en", "ru", "zh", "az".
	Lang UserLanguage `json:"lang" api:"required"`
	// User's name.
	Name string `json:"name" api:"required"`
	// User's phone.
	Phone string `json:"phone" api:"required"`
	// Services provider ID.
	Reseller int64 `json:"reseller" api:"required"`
	// SSO authentication flag. If `true` then user can login via SAML SSO.
	SSOAuth bool `json:"sso_auth" api:"required"`
	// Two-step verification:
	//
	// - `true` – user enabled two-step verification;
	// - `false` – user disabled two-step verification.
	TwoFa bool `json:"two_fa" api:"required"`
	// User's type.
	//
	// Any of "common", "reseller", "seller".
	UserType UserType `json:"user_type" api:"required"`
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
func (r UserUpdated) RawJSON() string { return r.JSON.raw }
func (r *UserUpdated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdatedClientAndRole struct {
	ClientCompanyName string `json:"client_company_name" api:"required"`
	ClientID          int64  `json:"client_id" api:"required"`
	// User's ID.
	UserID int64 `json:"user_id" api:"required"`
	// User role in this client.
	UserRoles []string `json:"user_roles" api:"required"`
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
func (r UserUpdatedClientAndRole) RawJSON() string { return r.JSON.raw }
func (r *UserUpdatedClientAndRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdateParams struct {
	// User's name.
	Name param.Opt[string] `json:"name,omitzero"`
	// User's phone.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// User's email address.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// System field. List of auth types available for the account.
	AuthTypes []AuthType `json:"auth_types,omitzero"`
	// User's language.
	//
	// Defines language of the control panel and email messages.
	//
	// Any of "de", "en", "ru", "zh", "az".
	Lang UserLanguage `json:"lang,omitzero"`
	paramObj
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListParams struct {
	// The maximum number of items.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset relative to the beginning of list.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Sort order as `field.direction`, e.g. `order_by=created_at.desc`. Allowed
	// fields: `created_at`, `email`, `name`. Defaults to `created_at.desc`.
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	// Filter by login method - a specific social provider, 'sso', or 'password'.
	//
	// Any of "password", "sso", "google-oauth2", "github".
	AuthType UserListParamsAuthType `query:"auth_type,omitzero" json:"-"`
	// Filter by creation time. Bracket operator required: `created_at[gte]`,
	// `created_at[gt]`, `created_at[lt]` or `created_at[lte]` (ISO 8601), e.g.
	// `created_at[gte]=2026-01-01T00:00:00Z&created_at[lt]=2026-02-01T00:00:00Z`.
	CreatedAt map[string]any `query:"created_at,omitzero" json:"-"`
	// Filter by email, case-insensitive. Bracket operator required: `email[exact]`,
	// `email[prefix]`, `email[contains]` or `email[suffix]`, e.g.
	// `email[contains]=example.com`.
	Email map[string]any `query:"email,omitzero" json:"-"`
	// Filter by user id. Repeat the parameter for several values (OR), e.g.
	// `ids=1&ids=2`.
	IDs []int64 `query:"ids,omitzero" json:"-"`
	// Filter by interface language. Repeat the parameter for several values (OR), e.g.
	// `langs=en&langs=ru`.
	//
	// Any of "az", "de", "en", "ru", "zh".
	Langs []string `query:"langs,omitzero" json:"-"`
	// Filter by name, case-insensitive. Bracket operator required: `name[exact]`,
	// `name[prefix]`, `name[contains]` or `name[suffix]`, e.g. `name[prefix]=Jo`.
	Name map[string]any `query:"name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by login method - a specific social provider, 'sso', or 'password'.
type UserListParamsAuthType string

const (
	UserListParamsAuthTypePassword     UserListParamsAuthType = "password"
	UserListParamsAuthTypeSSO          UserListParamsAuthType = "sso"
	UserListParamsAuthTypeGoogleOauth2 UserListParamsAuthType = "google-oauth2"
	UserListParamsAuthTypeGitHub       UserListParamsAuthType = "github"
)

type UserDeleteParams struct {
	ClientID int64 `path:"clientId" api:"required" json:"-"`
	paramObj
}

type UserInviteParams struct {
	// ID of account.
	ClientID int64 `json:"client_id" api:"required"`
	// User email.
	Email    string         `json:"email" api:"required" format:"email"`
	UserRole UserGroupParam `json:"user_role,omitzero" api:"required"`
	// User name.
	Name param.Opt[string] `json:"name,omitzero"`
	// User's language.
	//
	// Defines language of the control panel and email messages.
	//
	// Any of "de", "en", "ru", "zh", "az".
	Lang UserLanguage `json:"lang,omitzero"`
	paramObj
}

func (r UserInviteParams) MarshalJSON() (data []byte, err error) {
	type shadow UserInviteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserInviteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
