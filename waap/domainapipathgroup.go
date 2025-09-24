// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waap

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// DomainAPIPathGroupService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainAPIPathGroupService] method instead.
type DomainAPIPathGroupService struct {
	Options []option.RequestOption
}

// NewDomainAPIPathGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDomainAPIPathGroupService(opts ...option.RequestOption) (r DomainAPIPathGroupService) {
	r = DomainAPIPathGroupService{}
	r.Options = opts
	return
}

// Retrieve a list of API path groups for a specific domain
func (r *DomainAPIPathGroupService) List(ctx context.Context, domainID int64, opts ...option.RequestOption) (res *APIPathGroupList, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.gcore.com/")}, opts...)
	path := fmt.Sprintf("waap/v1/domains/%v/api-path-groups", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response model for the API path groups
type APIPathGroupList struct {
	// An array of api groups associated with the API path
	APIPathGroups []string `json:"api_path_groups,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIPathGroups respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIPathGroupList) RawJSON() string { return r.JSON.raw }
func (r *APIPathGroupList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
