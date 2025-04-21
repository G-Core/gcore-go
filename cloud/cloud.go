// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"github.com/stainless-sdks/gcore-go/internal/apijson"
	"github.com/stainless-sdks/gcore-go/option"
	"github.com/stainless-sdks/gcore-go/packages/resp"
)

// CloudService contains methods and other services that help with interacting with
// the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudService] method instead.
type CloudService struct {
	Options  []option.RequestOption
	Projects ProjectService
	Tasks    TaskService
	Regions  RegionService
	Quotas   QuotaService
	Secrets  SecretService
	SSHKeys  SSHKeyService
	IPRanges IPRangeService
}

// NewCloudService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCloudService(opts ...option.RequestOption) (r CloudService) {
	r = CloudService{}
	r.Options = opts
	r.Projects = NewProjectService(opts...)
	r.Tasks = NewTaskService(opts...)
	r.Regions = NewRegionService(opts...)
	r.Quotas = NewQuotaService(opts...)
	r.Secrets = NewSecretService(opts...)
	r.SSHKeys = NewSSHKeyService(opts...)
	r.IPRanges = NewIPRangeService(opts...)
	return
}

// '#/components/schemas/TaskIDsSerializer'
// "$.components.schemas.TaskIDsSerializer"
type TaskIDList struct {
	// '#/components/schemas/TaskIDsSerializer/properties/tasks'
	// "$.components.schemas.TaskIDsSerializer.properties.tasks"
	Tasks []string `json:"tasks,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Tasks       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskIDList) RawJSON() string { return r.JSON.raw }
func (r *TaskIDList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
