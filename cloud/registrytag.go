// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/param"
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

// Delete a specific tag from an artifact.
func (r *RegistryTagService) Delete(ctx context.Context, tagName string, body RegistryTagDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
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
	ProjectID      param.Opt[int64] `path:"project_id,omitzero,required" json:"-"`
	RegionID       param.Opt[int64] `path:"region_id,omitzero,required" json:"-"`
	RegistryID     int64            `path:"registry_id,required" json:"-"`
	RepositoryName string           `path:"repository_name,required" json:"-"`
	Digest         string           `path:"digest,required" json:"-"`
	paramObj
}
