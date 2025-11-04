// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud

import (
	"github.com/G-Core/gcore-go/option"
)

// DatabasePostgresClusterMetricService contains methods and other services that
// help with interacting with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatabasePostgresClusterMetricService] method instead.
type DatabasePostgresClusterMetricService struct {
	Options []option.RequestOption
}

// NewDatabasePostgresClusterMetricService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDatabasePostgresClusterMetricService(opts ...option.RequestOption) (r DatabasePostgresClusterMetricService) {
	r = DatabasePostgresClusterMetricService{}
	r.Options = opts
	return
}
