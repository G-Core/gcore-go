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

// CloudV1LaaService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1LaaService] method instead.
type CloudV1LaaService struct {
	Options []option.RequestOption
	Status  *CloudV1LaaStatusService
	Topics  *CloudV1LaaTopicService
}

// NewCloudV1LaaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV1LaaService(opts ...option.RequestOption) (r *CloudV1LaaService) {
	r = &CloudV1LaaService{}
	r.Options = opts
	r.Status = NewCloudV1LaaStatusService(opts...)
	r.Topics = NewCloudV1LaaTopicService(opts...)
	return
}

// Available destination regions
func (r *CloudV1LaaService) GetDestinationRegions(ctx context.Context, regionID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("cloud/v1/laas/%v/destination_regions", regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// List LaaS hosts
func (r *CloudV1LaaService) ListHosts(ctx context.Context, regionID int64, opts ...option.RequestOption) (res *CloudV1LaaListHostsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/laas/%v/hosts", regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Regenerate LaaS credentials
func (r *CloudV1LaaService) RegenerateCredentials(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *CloudV1LaaRegenerateCredentialsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/laas/%v/%v/users", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Validate custom_namespace creation availability
func (r *CloudV1LaaService) ValidateNamespaceCreation(ctx context.Context, projectID int64, regionID int64, namespaceName string, opts ...option.RequestOption) (res *CloudV1LaaValidateNamespaceCreationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if namespaceName == "" {
		err = errors.New("missing required namespace_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/laas/namespaces/%v/%v/%s", projectID, regionID, namespaceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// LaaS hosts list schema
type CloudV1LaaListHostsResponse struct {
	// Hosts list
	Hosts []string                        `json:"hosts"`
	JSON  cloudV1LaaListHostsResponseJSON `json:"-"`
}

// cloudV1LaaListHostsResponseJSON contains the JSON metadata for the struct
// [CloudV1LaaListHostsResponse]
type cloudV1LaaListHostsResponseJSON struct {
	Hosts       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1LaaListHostsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1LaaListHostsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1LaaRegenerateCredentialsResponse struct {
	// Password to be used in log shipper config to write logs to LaaS
	Password string `json:"password,required"`
	// Username to be used in log shipper config to write logs to LaaS
	Username string                                      `json:"username,required"`
	JSON     cloudV1LaaRegenerateCredentialsResponseJSON `json:"-"`
}

// cloudV1LaaRegenerateCredentialsResponseJSON contains the JSON metadata for the
// struct [CloudV1LaaRegenerateCredentialsResponse]
type cloudV1LaaRegenerateCredentialsResponseJSON struct {
	Password    apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1LaaRegenerateCredentialsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1LaaRegenerateCredentialsResponseJSON) RawJSON() string {
	return r.raw
}

type CloudV1LaaValidateNamespaceCreationResponse struct {
	// Is namespace available
	Available bool `json:"available,required"`
	// Client namespace
	Namespace string                                          `json:"namespace,required"`
	JSON      cloudV1LaaValidateNamespaceCreationResponseJSON `json:"-"`
}

// cloudV1LaaValidateNamespaceCreationResponseJSON contains the JSON metadata for
// the struct [CloudV1LaaValidateNamespaceCreationResponse]
type cloudV1LaaValidateNamespaceCreationResponseJSON struct {
	Available   apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudV1LaaValidateNamespaceCreationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudV1LaaValidateNamespaceCreationResponseJSON) RawJSON() string {
	return r.raw
}
