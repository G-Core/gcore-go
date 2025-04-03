// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gcore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
)

// CloudV1DbaaPostgreClusterUserCredentialService contains methods and other
// services that help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudV1DbaaPostgreClusterUserCredentialService] method instead.
type CloudV1DbaaPostgreClusterUserCredentialService struct {
	Options []option.RequestOption
}

// NewCloudV1DbaaPostgreClusterUserCredentialService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewCloudV1DbaaPostgreClusterUserCredentialService(opts ...option.RequestOption) (r *CloudV1DbaaPostgreClusterUserCredentialService) {
	r = &CloudV1DbaaPostgreClusterUserCredentialService{}
	r.Options = opts
	return
}

// Get credentials for given user in PG Cluster once.
func (r *CloudV1DbaaPostgreClusterUserCredentialService) Get(ctx context.Context, projectID int64, regionID int64, clusterName int64, username int64, opts ...option.RequestOption) (res *UserCredentials, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/clusters/%v/%v/%v/users/%v/credentials", projectID, regionID, clusterName, username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Recreate credentials for given user in PG Cluster.
func (r *CloudV1DbaaPostgreClusterUserCredentialService) Recreate(ctx context.Context, projectID int64, regionID int64, clusterName int64, username int64, opts ...option.RequestOption) (res *UserCredentials, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("cloud/v1/dbaas/postgres/clusters/%v/%v/%v/users/%v/credentials", projectID, regionID, clusterName, username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type UserCredentials struct {
	// Password
	Password string `json:"password,required"`
	// Username
	Username string              `json:"username,required"`
	JSON     userCredentialsJSON `json:"-"`
}

// userCredentialsJSON contains the JSON metadata for the struct [UserCredentials]
type userCredentialsJSON struct {
	Password    apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserCredentials) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userCredentialsJSON) RawJSON() string {
	return r.raw
}
