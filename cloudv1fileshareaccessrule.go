// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/param"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1FileShareAccessRuleService contains methods and other services that help
// with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1FileShareAccessRuleService] method instead.
type CloudV1FileShareAccessRuleService struct {
	Options []option.RequestOption
}

// NewCloudV1FileShareAccessRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCloudV1FileShareAccessRuleService(opts ...option.RequestOption) (r *CloudV1FileShareAccessRuleService) {
	r = &CloudV1FileShareAccessRuleService{}
	r.Options = opts
	return
}

// Create file share access rule
func (r *CloudV1FileShareAccessRuleService) New(ctx context.Context, projectID int64, regionID int64, fileShareID string, body CloudV1FileShareAccessRuleNewParams, opts ...option.RequestOption) (res *AccessRule, err error) {
	opts = append(r.Options[:], opts...)
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s/access_rule", projectID, regionID, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get file share access rules
func (r *CloudV1FileShareAccessRuleService) List(ctx context.Context, projectID int64, regionID int64, fileShareID string, opts ...option.RequestOption) (res *CloudV1FileShareAccessRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s/access_rule", projectID, regionID, fileShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete file share access rule
func (r *CloudV1FileShareAccessRuleService) Delete(ctx context.Context, projectID int64, regionID int64, fileShareID string, accessRuleID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if fileShareID == "" {
		err = errors.New("missing required file_share_id parameter")
		return
	}
	if accessRuleID == "" {
		err = errors.New("missing required access_rule_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/file_shares/%v/%v/%s/access_rule/%s", projectID, regionID, fileShareID, accessRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AccessLevelChoices string

const (
	AccessLevelChoicesRo AccessLevelChoices = "ro"
	AccessLevelChoicesRw AccessLevelChoices = "rw"
)

func (r AccessLevelChoices) IsKnown() bool {
	switch r {
	case AccessLevelChoicesRo, AccessLevelChoicesRw:
		return true
	}
	return false
}

type AccessRule struct {
	// Access Rule ID
	ID string `json:"id,required" format:"uuid4"`
	// Access mode
	AccessLevel AccessLevelChoices `json:"access_level,required"`
	// Source IP or network
	AccessTo string `json:"access_to,required" format:"ipvanyaddress"`
	// Access Rule state
	State AccessRuleState `json:"state,required"`
	JSON  accessRuleJSON  `json:"-"`
}

// accessRuleJSON contains the JSON metadata for the struct [AccessRule]
type accessRuleJSON struct {
	ID          apijson.Field
	AccessLevel apijson.Field
	AccessTo    apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleJSON) RawJSON() string {
	return r.raw
}

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

func (r AccessRuleState) IsKnown() bool {
	switch r {
	case AccessRuleStateActive, AccessRuleStateApplying, AccessRuleStateDenying, AccessRuleStateError, AccessRuleStateNew, AccessRuleStateQueuedToApply, AccessRuleStateQueuedToDeny:
		return true
	}
	return false
}

type CreateAccessRuleParam struct {
	// Access mode
	AccessMode param.Field[AccessLevelChoices] `json:"access_mode,required"`
	// Source IP or network
	IPAddress param.Field[string] `json:"ip_address,required" format:"ipvanyaddress"`
}

func (r CreateAccessRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CloudV1FileShareAccessRuleListResponse struct {
	// Number of objects
	Count int64 `json:"count,required"`
	// Objects
	Results []AccessRule                               `json:"results,required"`
	JSON    cloudV1FileShareAccessRuleListResponseJSON `json:"-"`
}

// cloudV1FileShareAccessRuleListResponseJSON contains the JSON metadata for the
// struct [CloudV1FileShareAccessRuleListResponse]
type cloudV1FileShareAccessRuleListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1FileShareAccessRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1FileShareAccessRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1FileShareAccessRuleNewParams struct {
	CreateAccessRule CreateAccessRuleParam `json:"create_access_rule,required"`
}

func (r CloudV1FileShareAccessRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateAccessRule)
}
