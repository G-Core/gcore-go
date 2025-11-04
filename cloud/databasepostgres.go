// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"github.com/G-Core/gcore-go/option"
)

// DatabasePostgresService contains methods and other services that help with
// interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatabasePostgresService] method instead.
type DatabasePostgresService struct {
	Options              []option.RequestOption
	Clusters             DatabasePostgresClusterService
	Configurations       DatabasePostgresConfigurationService
	CustomConfigurations DatabasePostgresCustomConfigurationService
}

// NewDatabasePostgresService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDatabasePostgresService(opts ...option.RequestOption) (r DatabasePostgresService) {
	r = DatabasePostgresService{}
	r.Options = opts
	r.Clusters = NewDatabasePostgresClusterService(opts...)
	r.Configurations = NewDatabasePostgresConfigurationService(opts...)
	r.CustomConfigurations = NewDatabasePostgresCustomConfigurationService(opts...)
	return
}
