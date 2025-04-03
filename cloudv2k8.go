// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV2K8Service contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV2K8Service] method instead.
type CloudV2K8Service struct {
	Options  []option.RequestOption
	Clusters *CloudV2K8ClusterService
}

// NewCloudV2K8Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudV2K8Service(opts ...option.RequestOption) (r *CloudV2K8Service) {
	r = &CloudV2K8Service{}
	r.Options = opts
	r.Clusters = NewCloudV2K8ClusterService(opts...)
	return
}

// List available k8s cluster versions for cluster creation
func (r *CloudV2K8Service) GetNewVersions(ctx context.Context, projectID int64, regionID int64, opts ...option.RequestOption) (res *VersionList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v2/k8s/%v/%v/create_versions", projectID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
