// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1ApptemplateService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1ApptemplateService] method instead.
type CloudV1ApptemplateService struct {
	Options []option.RequestOption
}

// NewCloudV1ApptemplateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCloudV1ApptemplateService(opts ...option.RequestOption) (r *CloudV1ApptemplateService) {
	r = &CloudV1ApptemplateService{}
	r.Options = opts
	return
}

// Get apptemplate
func (r *CloudV1ApptemplateService) Get(ctx context.Context, projectID int64, regionID int64, apptemplateID string, opts ...option.RequestOption) (res *AppTemplate, err error) {
	opts = append(r.Options[:], opts...)
	if apptemplateID == "" {
		err = errors.New("missing required apptemplate_id parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/apptemplates/%v/%v/%s", projectID, regionID, apptemplateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve marketplace application templates list. Application templates are used
// in instance creation API v2
func (r *CloudV1ApptemplateService) List(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *CloudV1ApptemplateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/apptemplates/%v/%v", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// AppTemplate schema
type AppTemplate struct {
	// Application id
	ID string `json:"id,required"`
	// Application category
	Category string `json:"category,required"`
	// Localization key
	Description string `json:"description,required"`
	// Application display name
	DisplayName string `json:"display_name,required"`
	// Image used to deploy the application
	ImageName string `json:"image_name,required"`
	// OS name
	OsName string `json:"os_name,required"`
	// OS version
	OsVersion string `json:"os_version,required"`
	// Region id
	RegionID int64 `json:"region_id,required"`
	// Localization key
	ShortDescription string `json:"short_description,required"`
	// Application version
	Version string `json:"version,required"`
	// Application website
	Website string `json:"website,required"`
	// Config
	AppConfig []interface{} `json:"app_config"`
	// Developer name
	Developer string `json:"developer"`
	// Application disk requirements
	MinDisk int64 `json:"min_disk"`
	// Application RAM requirements
	MinRam int64 `json:"min_ram"`
	// Application VCPUs requirements
	MinVcpus int64 `json:"min_vcpus"`
	// Localization key
	Usage string          `json:"usage"`
	JSON  appTemplateJSON `json:"-"`
}

// appTemplateJSON contains the JSON metadata for the struct [AppTemplate]
type appTemplateJSON struct {
	ID               apijson.Field
	Category         apijson.Field
	Description      apijson.Field
	DisplayName      apijson.Field
	ImageName        apijson.Field
	OsName           apijson.Field
	OsVersion        apijson.Field
	RegionID         apijson.Field
	ShortDescription apijson.Field
	Version          apijson.Field
	Website          apijson.Field
	AppConfig        apijson.Field
	Developer        apijson.Field
	MinDisk          apijson.Field
	MinRam           apijson.Field
	MinVcpus         apijson.Field
	Usage            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppTemplate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appTemplateJSON) RawJSON() string {
	return r.raw
}

type CloudV1ApptemplateListResponse struct {
	// Number of objects
	Count int64 `json:"count"`
	// Objects
	Results []AppTemplate                      `json:"results"`
	JSON    cloudV1ApptemplateListResponseJSON `json:"-"`
}

// cloudV1ApptemplateListResponseJSON contains the JSON metadata for the struct
// [CloudV1ApptemplateListResponse]
type cloudV1ApptemplateListResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1ApptemplateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1ApptemplateListResponseJSON) RawJSON() string {
	return r.raw
}
