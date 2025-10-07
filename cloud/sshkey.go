// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
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

// SSHKeyService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSSHKeyService] method instead.
type SSHKeyService struct {
	Options []option.RequestOption
}

// NewSSHKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSSHKeyService(opts ...option.RequestOption) (r SSHKeyService) {
	r = SSHKeyService{}
	r.Options = opts
	return
}

// To generate a key, omit the `public_key` parameter from the request body
func (r *SSHKeyService) New(ctx context.Context, params SSHKeyNewParams, opts ...option.RequestOption) (res *SSHKeyCreated, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ssh_keys/%v", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Share or unshare SSH key with users
func (r *SSHKeyService) Update(ctx context.Context, sshKeyID string, params SSHKeyUpdateParams, opts ...option.RequestOption) (res *SSHKey, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if sshKeyID == "" {
		err = errors.New("missing required ssh_key_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ssh_keys/%v/%s", params.ProjectID.Value, sshKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List SSH keys
func (r *SSHKeyService) List(ctx context.Context, params SSHKeyListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SSHKey], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ssh_keys/%v", params.ProjectID.Value)
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

// List SSH keys
func (r *SSHKeyService) ListAutoPaging(ctx context.Context, params SSHKeyListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SSHKey] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, params, opts...))
}

// Delete SSH key
func (r *SSHKeyService) Delete(ctx context.Context, sshKeyID string, body SSHKeyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if sshKeyID == "" {
		err = errors.New("missing required ssh_key_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ssh_keys/%v/%s", body.ProjectID.Value, sshKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get SSH key
func (r *SSHKeyService) Get(ctx context.Context, sshKeyID string, query SSHKeyGetParams, opts ...option.RequestOption) (res *SSHKey, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if sshKeyID == "" {
		err = errors.New("missing required ssh_key_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ssh_keys/%v/%s", query.ProjectID.Value, sshKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SSHKey struct {
	// SSH key ID
	ID string `json:"id,required" format:"uuid4"`
	// SSH key creation time
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Fingerprint
	Fingerprint string `json:"fingerprint,required"`
	// SSH key name
	Name string `json:"name,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// The public part of an SSH key is the shareable portion of an SSH key pair. It
	// can be safely sent to servers or services to grant access. It does not contain
	// sensitive information.
	PublicKey string `json:"public_key,required"`
	// SSH key will be visible to all users in the project
	SharedInProject bool `json:"shared_in_project,required"`
	// SSH key state
	//
	// Any of "ACTIVE", "DELETING".
	State SSHKeyState `json:"state,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Fingerprint     respjson.Field
		Name            respjson.Field
		ProjectID       respjson.Field
		PublicKey       respjson.Field
		SharedInProject respjson.Field
		State           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSHKey) RawJSON() string { return r.JSON.raw }
func (r *SSHKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SSH key state
type SSHKeyState string

const (
	SSHKeyStateActive   SSHKeyState = "ACTIVE"
	SSHKeyStateDeleting SSHKeyState = "DELETING"
)

type SSHKeyCreated struct {
	// SSH key ID
	ID string `json:"id,required" format:"uuid4"`
	// SSH key creation time
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Fingerprint
	Fingerprint string `json:"fingerprint,required"`
	// SSH key name
	Name string `json:"name,required"`
	// The private part of an SSH key is the confidential portion of the key pair. It
	// should never be shared or exposed. This key is used to prove your identity when
	// connecting to a server.
	//
	// If you omit the `public_key`, the platform will generate a key for you. The
	// `private_key` will be returned **once** in the API response. Be sure to save it
	// securely, as it cannot be retrieved again later.
	//
	// Best practice: Save the private key to a secure location on your machine (e.g.,
	// `~/.ssh/id_ed25519`) and set the file permissions to be readable only by you.
	PrivateKey string `json:"private_key,required"`
	// Project ID
	ProjectID int64 `json:"project_id,required"`
	// The public part of an SSH key is the shareable portion of an SSH key pair. It
	// can be safely sent to servers or services to grant access. It does not contain
	// sensitive information.
	PublicKey string `json:"public_key,required"`
	// SSH key will be visible to all users in the project
	SharedInProject bool `json:"shared_in_project,required"`
	// SSH key state
	//
	// Any of "ACTIVE", "DELETING".
	State SSHKeyCreatedState `json:"state,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Fingerprint     respjson.Field
		Name            respjson.Field
		PrivateKey      respjson.Field
		ProjectID       respjson.Field
		PublicKey       respjson.Field
		SharedInProject respjson.Field
		State           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSHKeyCreated) RawJSON() string { return r.JSON.raw }
func (r *SSHKeyCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SSH key state
type SSHKeyCreatedState string

const (
	SSHKeyCreatedStateActive   SSHKeyCreatedState = "ACTIVE"
	SSHKeyCreatedStateDeleting SSHKeyCreatedState = "DELETING"
)

type SSHKeyNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// SSH key name
	Name string `json:"name,required"`
	// The public part of an SSH key is the shareable portion of an SSH key pair. It
	// can be safely sent to servers or services to grant access. It does not contain
	// sensitive information.
	//
	//   - If you’re uploading your own key, provide the public part here (usually found
	//     in a file like `id_ed25519.pub`).
	//   - If you want the platform to generate an Ed25519 key pair for you, leave this
	//     field empty — the system will return the private key in the response **once
	//     only**.
	PublicKey param.Opt[string] `json:"public_key,omitzero"`
	// SSH key is shared with all users in the project
	SharedInProject param.Opt[bool] `json:"shared_in_project,omitzero"`
	paramObj
}

func (r SSHKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SSHKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SSHKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SSHKeyUpdateParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Share your ssh key with all users in the project
	SharedInProject bool `json:"shared_in_project,required"`
	paramObj
}

func (r SSHKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SSHKeyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SSHKeyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SSHKeyListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Maximum number of SSH keys to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Offset for pagination
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Sort order for the SSH keys
	//
	// Any of "created_at.asc", "created_at.desc", "name.asc", "name.desc".
	OrderBy SSHKeyListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SSHKeyListParams]'s query parameters as `url.Values`.
func (r SSHKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort order for the SSH keys
type SSHKeyListParamsOrderBy string

const (
	SSHKeyListParamsOrderByCreatedAtAsc  SSHKeyListParamsOrderBy = "created_at.asc"
	SSHKeyListParamsOrderByCreatedAtDesc SSHKeyListParamsOrderBy = "created_at.desc"
	SSHKeyListParamsOrderByNameAsc       SSHKeyListParamsOrderBy = "name.asc"
	SSHKeyListParamsOrderByNameDesc      SSHKeyListParamsOrderBy = "name.desc"
)

type SSHKeyDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

type SSHKeyGetParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}
