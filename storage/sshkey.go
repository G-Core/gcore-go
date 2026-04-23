// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

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

// SSH keys enable secure access to SFTP storage by associating public keys with
// user accounts for authentication.
//
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

// Creates an SSH public key for SFTP storage access. These keys are used for
// passwordless authentication to SFTP storages.
func (r *SSHKeyService) New(ctx context.Context, body SSHKeyNewParams, opts ...option.RequestOption) (res *SSHKey, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/v4/ssh_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of all SSH public keys for SFTP storage access.
func (r *SSHKeyService) List(ctx context.Context, query SSHKeyListParams, opts ...option.RequestOption) (res *pagination.OffsetPage[SSHKey], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "storage/v4/ssh_keys"
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

// Returns a paginated list of all SSH public keys for SFTP storage access.
func (r *SSHKeyService) ListAutoPaging(ctx context.Context, query SSHKeyListParams, opts ...option.RequestOption) *pagination.OffsetPageAutoPager[SSHKey] {
	return pagination.NewOffsetPageAutoPager(r.List(ctx, query, opts...))
}

// Removes an SSH public key. This will revoke SFTP access for any storages using
// this key.
func (r *SSHKeyService) Delete(ctx context.Context, sshKeyID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("storage/v4/ssh_keys/%v", sshKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieves a single SSH key by ID.
func (r *SSHKeyService) Get(ctx context.Context, sshKeyID int64, opts ...option.RequestOption) (res *SSHKey, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("storage/v4/ssh_keys/%v", sshKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type SSHKey struct {
	// Unique identifier for the SSH key
	ID int64 `json:"id" api:"required"`
	// ISO 8601 timestamp when the SSH key was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// User-defined name for the SSH key
	Name string `json:"name" api:"required"`
	// The SSH public key content
	PublicKey string `json:"public_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		PublicKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSHKey) RawJSON() string { return r.JSON.raw }
func (r *SSHKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SSHKeyNewParams struct {
	// User-defined name for the SSH key
	Name string `json:"name" api:"required"`
	// The SSH public key content (ssh-rsa or ssh-ed25519 format)
	PublicKey string `json:"public_key" api:"required"`
	paramObj
}

func (r SSHKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SSHKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SSHKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SSHKeyListParams struct {
	// Maximum number of items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by name (partial match)
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Number of items to skip
	Offset  param.Opt[int64]  `query:"offset,omitzero" json:"-"`
	OrderBy param.Opt[string] `query:"order_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SSHKeyListParams]'s query parameters as `url.Values`.
func (r SSHKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
