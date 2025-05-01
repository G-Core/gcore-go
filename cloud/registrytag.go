// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gcore-go/internal/requestconfig"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/param"
)

// RegistryTagService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryTagService] method instead.
type RegistryTagService struct {
	Options []option.RequestOption
}

// NewRegistryTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRegistryTagService(opts ...option.RequestOption) (r RegistryTagService) {
	r = RegistryTagService{}
	r.Options = opts
	return
}

// Delete a tag
func (r *RegistryTagService) Delete(ctx context.Context, tagName string, body RegistryTagDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&body.RegionID, precfg.CloudRegionID)
	if !body.ProjectID.IsPresent() {
		err = errors.New("missing required project_id parameter")
		return
	}
	if !body.RegionID.IsPresent() {
		err = errors.New("missing required region_id parameter")
		return
	}
	if body.RepositoryName == "" {
		err = errors.New("missing required repository_name parameter")
		return
	}
	if body.Digest == "" {
		err = errors.New("missing required digest parameter")
		return
	}
	if tagName == "" {
		err = errors.New("missing required tag_name parameter")
		return
	}
	path := fmt.Sprintf("cloud/v1/registries/%v/%v/%v/repositories/%s/artifacts/%s/tags/%s", body.ProjectID.Value, body.RegionID.Value, body.RegistryID, body.RepositoryName, body.Digest, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type RegistryTagDeleteParams struct {
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D%2Fartifacts%2F%7Bdigest%7D%2Ftags%2F%7Btag_name%7D/delete/parameters/0/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}/artifacts/{digest}/tags/{tag_name}']['delete'].parameters[0].schema"
	ProjectID param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D%2Fartifacts%2F%7Bdigest%7D%2Ftags%2F%7Btag_name%7D/delete/parameters/1/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}/artifacts/{digest}/tags/{tag_name}']['delete'].parameters[1].schema"
	RegionID param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D%2Fartifacts%2F%7Bdigest%7D%2Ftags%2F%7Btag_name%7D/delete/parameters/2/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}/artifacts/{digest}/tags/{tag_name}']['delete'].parameters[2].schema"
	RegistryID int64 `path:"registry_id,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D%2Fartifacts%2F%7Bdigest%7D%2Ftags%2F%7Btag_name%7D/delete/parameters/3/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}/artifacts/{digest}/tags/{tag_name}']['delete'].parameters[3].schema"
	RepositoryName string `path:"repository_name,required" json:"-"`
	// '#/paths/%2Fcloud%2Fv1%2Fregistries%2F%7Bproject_id%7D%2F%7Bregion_id%7D%2F%7Bregistry_id%7D%2Frepositories%2F%7Brepository_name%7D%2Fartifacts%2F%7Bdigest%7D%2Ftags%2F%7Btag_name%7D/delete/parameters/4/schema'
	// "$.paths['/cloud/v1/registries/{project_id}/{region_id}/{registry_id}/repositories/{repository_name}/artifacts/{digest}/tags/{tag_name}']['delete'].parameters[4].schema"
	Digest string `path:"digest,required" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f RegistryTagDeleteParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
