// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/apiquery"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/pagination"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
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

// To generate a key, omit the public_key parameter from the request body
func (r *SSHKeyService) New(ctx context.Context, params SSHKeyNewParams, opts ...option.RequestOption) (res *SSHKeyCreated, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/ssh_keys/%v", params.ProjectID.Value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Share or unshare SSH key with users
func (r *SSHKeyService) Update(ctx context.Context, sshKeyID string, params SSHKeyUpdateParams, opts ...option.RequestOption) (res *SSHKey, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if !params.ProjectID.IsPresent() {
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	if !params.ProjectID.IsPresent() {
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.ProjectID)
	if !body.ProjectID.IsPresent() {
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
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.ProjectID)
	if !query.ProjectID.IsPresent() {
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

// '#/components/schemas/SSHKeySerializer' "$.components.schemas.SSHKeySerializer"
type SSHKey struct {
	// '#/components/schemas/SSHKeySerializer/properties/id'
	// "$.components.schemas.SSHKeySerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/SSHKeySerializer/properties/created_at/anyOf/0'
	// "$.components.schemas.SSHKeySerializer.properties.created_at.anyOf[0]"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/SSHKeySerializer/properties/fingerprint'
	// "$.components.schemas.SSHKeySerializer.properties.fingerprint"
	Fingerprint string `json:"fingerprint,required"`
	// '#/components/schemas/SSHKeySerializer/properties/name'
	// "$.components.schemas.SSHKeySerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/SSHKeySerializer/properties/project_id/anyOf/0'
	// "$.components.schemas.SSHKeySerializer.properties.project_id.anyOf[0]"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/SSHKeySerializer/properties/public_key'
	// "$.components.schemas.SSHKeySerializer.properties.public_key"
	PublicKey string `json:"public_key,required"`
	// '#/components/schemas/SSHKeySerializer/properties/shared_in_project'
	// "$.components.schemas.SSHKeySerializer.properties.shared_in_project"
	SharedInProject bool `json:"shared_in_project,required"`
	// '#/components/schemas/SSHKeySerializer/properties/state'
	// "$.components.schemas.SSHKeySerializer.properties.state"
	//
	// Any of "ACTIVE", "DELETING".
	State SSHKeyState `json:"state,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID              resp.Field
		CreatedAt       resp.Field
		Fingerprint     resp.Field
		Name            resp.Field
		ProjectID       resp.Field
		PublicKey       resp.Field
		SharedInProject resp.Field
		State           resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSHKey) RawJSON() string { return r.JSON.raw }
func (r *SSHKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/SSHKeySerializer/properties/state'
// "$.components.schemas.SSHKeySerializer.properties.state"
type SSHKeyState string

const (
	SSHKeyStateActive   SSHKeyState = "ACTIVE"
	SSHKeyStateDeleting SSHKeyState = "DELETING"
)

// '#/components/schemas/CreatedSSHKeySerializer'
// "$.components.schemas.CreatedSSHKeySerializer"
type SSHKeyCreated struct {
	// '#/components/schemas/CreatedSSHKeySerializer/properties/id'
	// "$.components.schemas.CreatedSSHKeySerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/CreatedSSHKeySerializer/properties/created_at'
	// "$.components.schemas.CreatedSSHKeySerializer.properties.created_at"
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// '#/components/schemas/CreatedSSHKeySerializer/properties/fingerprint'
	// "$.components.schemas.CreatedSSHKeySerializer.properties.fingerprint"
	Fingerprint string `json:"fingerprint,required"`
	// '#/components/schemas/CreatedSSHKeySerializer/properties/name'
	// "$.components.schemas.CreatedSSHKeySerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreatedSSHKeySerializer/properties/private_key/anyOf/0'
	// "$.components.schemas.CreatedSSHKeySerializer.properties.private_key.anyOf[0]"
	PrivateKey string `json:"private_key,required"`
	// '#/components/schemas/CreatedSSHKeySerializer/properties/project_id'
	// "$.components.schemas.CreatedSSHKeySerializer.properties.project_id"
	ProjectID int64 `json:"project_id,required"`
	// '#/components/schemas/CreatedSSHKeySerializer/properties/public_key'
	// "$.components.schemas.CreatedSSHKeySerializer.properties.public_key"
	PublicKey string `json:"public_key,required"`
	// '#/components/schemas/CreatedSSHKeySerializer/properties/shared_in_project'
	// "$.components.schemas.CreatedSSHKeySerializer.properties.shared_in_project"
	SharedInProject bool `json:"shared_in_project,required"`
	// '#/components/schemas/CreatedSSHKeySerializer/properties/state'
	// "$.components.schemas.CreatedSSHKeySerializer.properties.state"
	//
	// Any of "ACTIVE", "DELETING".
	State SSHKeyCreatedState `json:"state,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID              resp.Field
		CreatedAt       resp.Field
		Fingerprint     resp.Field
		Name            resp.Field
		PrivateKey      resp.Field
		ProjectID       resp.Field
		PublicKey       resp.Field
		SharedInProject resp.Field
		State           resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSHKeyCreated) RawJSON() string { return r.JSON.raw }
func (r *SSHKeyCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/CreatedSSHKeySerializer/properties/state'
// "$.components.schemas.CreatedSSHKeySerializer.properties.state"
type SSHKeyCreatedState string

const (
	SSHKeyCreatedStateActive   SSHKeyCreatedState = "ACTIVE"
	SSHKeyCreatedStateDeleting SSHKeyCreatedState = "DELETING"
)

type SSHKeyNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fssh_keys%2F%7Bproject_id%7D/post/parameters/0/schema'
	// "$.paths['/cloud/v1/ssh_keys/{project_id}'].post.parameters[0].schema"
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateSSHKeySerializer/properties/name'
	// "$.components.schemas.CreateSSHKeySerializer.properties.name"
	Name string `json:"name,required"`
	// '#/components/schemas/CreateSSHKeySerializer/properties/public_key'
	// "$.components.schemas.CreateSSHKeySerializer.properties.public_key"
	PublicKey param.Opt[string] `json:"public_key,omitzero"`
	// '#/components/schemas/CreateSSHKeySerializer/properties/shared_in_project'
	// "$.components.schemas.CreateSSHKeySerializer.properties.shared_in_project"
	SharedInProject param.Opt[bool] `json:"shared_in_project,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SSHKeyNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SSHKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SSHKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type SSHKeyUpdateParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fssh_keys%2F%7Bproject_id%7D%2F%7Bssh_key_id%7D/patch/parameters/0/schema'
	// "$.paths['/cloud/v1/ssh_keys/{project_id}/{ssh_key_id}'].patch.parameters[0].schema"
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/components/schemas/ShareSSHKeySerializer/properties/shared_in_project'
	// "$.components.schemas.ShareSSHKeySerializer.properties.shared_in_project"
	SharedInProject bool `json:"shared_in_project,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SSHKeyUpdateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SSHKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SSHKeyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type SSHKeyListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fssh_keys%2F%7Bproject_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/ssh_keys/{project_id}'].get.parameters[0].schema"
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fssh_keys%2F%7Bproject_id%7D/get/parameters/1'
	// "$.paths['/cloud/v1/ssh_keys/{project_id}'].get.parameters[1]"
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fssh_keys%2F%7Bproject_id%7D/get/parameters/2'
	// "$.paths['/cloud/v1/ssh_keys/{project_id}'].get.parameters[2]"
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fssh_keys%2F%7Bproject_id%7D/get/parameters/3'
	// "$.paths['/cloud/v1/ssh_keys/{project_id}'].get.parameters[3]"
	//
	// Any of "created_at.asc", "created_at.desc", "name.asc", "name.desc".
	OrderBy SSHKeyListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SSHKeyListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SSHKeyListParams]'s query parameters as `url.Values`.
func (r SSHKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// '#/paths/%2Fcloud%2Fv1%2Fssh_keys%2F%7Bproject_id%7D/get/parameters/3'
// "$.paths['/cloud/v1/ssh_keys/{project_id}'].get.parameters[3]"
type SSHKeyListParamsOrderBy string

const (
	SSHKeyListParamsOrderByCreatedAtAsc  SSHKeyListParamsOrderBy = "created_at.asc"
	SSHKeyListParamsOrderByCreatedAtDesc SSHKeyListParamsOrderBy = "created_at.desc"
	SSHKeyListParamsOrderByNameAsc       SSHKeyListParamsOrderBy = "name.asc"
	SSHKeyListParamsOrderByNameDesc      SSHKeyListParamsOrderBy = "name.desc"
)

type SSHKeyDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fssh_keys%2F%7Bproject_id%7D%2F%7Bssh_key_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/ssh_keys/{project_id}/{ssh_key_id}']['delete'].parameters[0].schema"
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SSHKeyDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type SSHKeyGetParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fssh_keys%2F%7Bproject_id%7D%2F%7Bssh_key_id%7D/get/parameters/0/schema'
	// "$.paths['/cloud/v1/ssh_keys/{project_id}/{ssh_key_id}'].get.parameters[0].schema"
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SSHKeyGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
