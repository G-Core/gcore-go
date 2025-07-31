// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security

import (
	"github.com/G-Core/gcore-go/option"
)

// SecurityService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecurityService] method instead.
type SecurityService struct {
	Options          []option.RequestOption
	Events           EventService
	BgpAnnounces     BgpAnnounceService
	ProfileTemplates ProfileTemplateService
	Profiles         ProfileService
}

// NewSecurityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecurityService(opts ...option.RequestOption) (r SecurityService) {
	r = SecurityService{}
	r.Options = opts
	r.Events = NewEventService(opts...)
	r.BgpAnnounces = NewBgpAnnounceService(opts...)
	r.ProfileTemplates = NewProfileTemplateService(opts...)
	r.Profiles = NewProfileService(opts...)
	return
}
