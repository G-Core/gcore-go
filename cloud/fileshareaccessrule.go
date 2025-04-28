// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// FileShareAccessRuleService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileShareAccessRuleService] method instead.
type FileShareAccessRuleService struct {
	Options []option.RequestOption
}

// NewFileShareAccessRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFileShareAccessRuleService(opts ...option.RequestOption) (r FileShareAccessRuleService) {
	r = FileShareAccessRuleService{}
	r.Options = opts
	return
}

// Create file share access rule
func (r *FileShareAccessRuleService) New(ctx context.Context, fileShareID string, params FileShareAccessRuleNewParams, opts ...option.RequestOption) (res *AccessRule, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.RegionID)
	if !params.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s/access_rule", params.ProjectID.Value, params.RegionID.Value, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get file share access rules
func (r *FileShareAccessRuleService) List(ctx context.Context, fileShareID string, query FileShareAccessRuleListParams, opts ...option.RequestOption) (res *AccessRuleList, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.RegionID)
	if !query.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s/access_rule", query.ProjectID.Value, query.RegionID.Value, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete file share access rule
func (r *FileShareAccessRuleService) Delete(ctx context.Context, fileShareID string, accessRuleID string, body FileShareAccessRuleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.ProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.RegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	if accessRuleID == "" {
		err = errors.New("missing required access_rule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s/access_rule/%s", body.ProjectID.Value, body.RegionID.Value, fileShareID, accessRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// '#/components/schemas/AccessRuleSerializer'
// "$.components.schemas.AccessRuleSerializer"
type AccessRule struct {
	// '#/components/schemas/AccessRuleSerializer/properties/id'
	// "$.components.schemas.AccessRuleSerializer.properties.id"
	ID string `json:"id,required" format:"uuid4"`
	// '#/components/schemas/AccessRuleSerializer/properties/access_level'
	// "$.components.schemas.AccessRuleSerializer.properties.access_level"
	//
	// Any of "ro", "rw".
	AccessLevel AccessRuleAccessLevel `json:"access_level,required"`
	// '#/components/schemas/AccessRuleSerializer/properties/access_to/anyOf/0'
	// "$.components.schemas.AccessRuleSerializer.properties.access_to.anyOf[0]"
	AccessTo string `json:"access_to,required" format:"ipvanyaddress"`
	// '#/components/schemas/AccessRuleSerializer/properties/state'
	// "$.components.schemas.AccessRuleSerializer.properties.state"
	//
	// Any of "active", "applying", "denying", "error", "new", "queued_to_apply",
	// "queued_to_deny".
	State AccessRuleState `json:"state,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		AccessLevel resp.Field
		AccessTo    resp.Field
		State       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessRule) RawJSON() string { return r.JSON.raw }
func (r *AccessRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// '#/components/schemas/AccessRuleSerializer/properties/access_level'
// "$.components.schemas.AccessRuleSerializer.properties.access_level"
type AccessRuleAccessLevel string

const (
	AccessRuleAccessLevelRo AccessRuleAccessLevel = "ro"
	AccessRuleAccessLevelRw AccessRuleAccessLevel = "rw"
)

// '#/components/schemas/AccessRuleSerializer/properties/state'
// "$.components.schemas.AccessRuleSerializer.properties.state"
type AccessRuleState string

const (
	AccessRuleStateActive        AccessRuleState = "active"
	AccessRuleStateApplying      AccessRuleState = "applying"
	AccessRuleStateDenying       AccessRuleState = "denying"
	AccessRuleStateError         AccessRuleState = "error"
	AccessRuleStateNew           AccessRuleState = "new"
	AccessRuleStateQueuedToApply AccessRuleState = "queued_to_apply"
	AccessRuleStateQueuedToDeny  AccessRuleState = "queued_to_deny"
)

// '#/components/schemas/AccessRuleCollectionSerializer'
// "$.components.schemas.AccessRuleCollectionSerializer"
type AccessRuleList struct {
	// '#/components/schemas/AccessRuleCollectionSerializer/properties/count'
	// "$.components.schemas.AccessRuleCollectionSerializer.properties.count"
	Count int64 `json:"count,required"`
	// '#/components/schemas/AccessRuleCollectionSerializer/properties/results'
	// "$.components.schemas.AccessRuleCollectionSerializer.properties.results"
	Results []AccessRule `json:"results,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Results     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessRuleList) RawJSON() string { return r.JSON.raw }
func (r *AccessRuleList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileShareAccessRuleNewParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfile_share_id%7D%2Faccess_rule/post/parameters/0/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}/{file_share_id}/access_rule'].post.parameters[0].schema"
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfile_share_id%7D%2Faccess_rule/post/parameters/1/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}/{file_share_id}/access_rule'].post.parameters[1].schema"
	//
	// Use [option.WithRegionID] on the client to set a global default for this field.
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/components/schemas/CreateAccessRuleSerializer/properties/access_mode'
	// "$.components.schemas.CreateAccessRuleSerializer.properties.access_mode"
	//
	// Any of "ro", "rw".
	AccessMode FileShareAccessRuleNewParamsAccessMode `json:"access_mode,omitzero,required"`
	// '#/components/schemas/CreateAccessRuleSerializer/properties/ip_address/anyOf/0'
	// "$.components.schemas.CreateAccessRuleSerializer.properties.ip_address.anyOf[0]"
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FileShareAccessRuleNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r FileShareAccessRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FileShareAccessRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// '#/components/schemas/CreateAccessRuleSerializer/properties/access_mode'
// "$.components.schemas.CreateAccessRuleSerializer.properties.access_mode"
type FileShareAccessRuleNewParamsAccessMode string

const (
	FileShareAccessRuleNewParamsAccessModeRo FileShareAccessRuleNewParamsAccessMode = "ro"
	FileShareAccessRuleNewParamsAccessModeRw FileShareAccessRuleNewParamsAccessMode = "rw"
)

type FileShareAccessRuleListParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfile_share_id%7D%2Faccess_rule/get/parameters/0/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}/{file_share_id}/access_rule'].get.parameters[0].schema"
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfile_share_id%7D%2Faccess_rule/get/parameters/1/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}/{file_share_id}/access_rule'].get.parameters[1].schema"
	//
	// Use [option.WithRegionID] on the client to set a global default for this field.
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FileShareAccessRuleListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

type FileShareAccessRuleDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfile_share_id%7D%2Faccess_rule%2F%7Baccess_rule_id%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}/{file_share_id}/access_rule/{access_rule_id}']['delete'].parameters[0].schema"
	//
	// Use [option.WithProjectID] on the client to set a global default for this field.
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Ffile_shares%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bfile_share_id%7D%2Faccess_rule%2F%7Baccess_rule_id%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/file_shares/{project_id}/{region_id}/{file_share_id}/access_rule/{access_rule_id}']['delete'].parameters[1].schema"
	//
	// Use [option.WithRegionID] on the client to set a global default for this field.
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f FileShareAccessRuleDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
