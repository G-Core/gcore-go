// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
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
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	if !params.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !params.RegionID.Valid() {
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
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&query.RegionID, precfg.CloudRegionID)
	if !query.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !query.RegionID.Valid() {
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
func (r *FileShareAccessRuleService) Delete(ctx context.Context, accessRuleID string, body FileShareAccessRuleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.Valid() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.Valid() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if body.FileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	if accessRuleID == "" {
		err = errors.New("missing required access_rule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s/access_rule/%s", body.ProjectID.Value, body.RegionID.Value, body.FileShareID, accessRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AccessRule struct {
	// Access Rule ID
	ID string `json:"id,required" format:"uuid4"`
	// Access mode
	//
	// Any of "ro", "rw".
	AccessLevel AccessRuleAccessLevel `json:"access_level,required"`
	// Source IP or network
	AccessTo string `json:"access_to,required" format:"ipvanyaddress"`
	// Access Rule state
	//
	// Any of "active", "applying", "denying", "error", "new", "queued_to_apply",
	// "queued_to_deny".
	State AccessRuleState `json:"state,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccessLevel respjson.Field
		AccessTo    respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessRule) RawJSON() string { return r.JSON.raw }
func (r *AccessRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Access mode
type AccessRuleAccessLevel string

const (
	AccessRuleAccessLevelRo AccessRuleAccessLevel = "ro"
	AccessRuleAccessLevelRw AccessRuleAccessLevel = "rw"
)

// Access Rule state
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

type AccessRuleList struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []AccessRule `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessRuleList) RawJSON() string { return r.JSON.raw }
func (r *AccessRuleList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileShareAccessRuleNewParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// Access mode
	//
	// Any of "ro", "rw".
	AccessMode FileShareAccessRuleNewParamsAccessMode `json:"access_mode,omitzero,required"`
	// Source IP or network
	IPAddress string `json:"ip_address,required" format:"ipvanyaddress"`
	paramObj
}

func (r FileShareAccessRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FileShareAccessRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileShareAccessRuleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Access mode
type FileShareAccessRuleNewParamsAccessMode string

const (
	FileShareAccessRuleNewParamsAccessModeRo FileShareAccessRuleNewParamsAccessMode = "ro"
	FileShareAccessRuleNewParamsAccessModeRw FileShareAccessRuleNewParamsAccessMode = "rw"
)

type FileShareAccessRuleListParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	paramObj
}

type FileShareAccessRuleDeleteParams struct {
	// Project ID
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// Region ID
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// File Share ID
	FileShareID string `path:"file_share_id,required" format:"uuid4" json:"-"`
	paramObj
}
