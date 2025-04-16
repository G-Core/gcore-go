// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"github.com/stainless-sdks/gcore-go/option"
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
	return
}
