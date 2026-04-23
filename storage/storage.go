// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"github.com/G-Core/gcore-go/option"
)

// StorageService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageService] method instead.
type StorageService struct {
	Options []option.RequestOption
	// Locations represent cloud regions where new storages can be created.
	Locations      LocationService
	ObjectStorages ObjectStorageService
	SftpStorages   SftpStorageService
	// SSH keys enable secure access to SFTP storage by associating public keys with
	// user accounts for authentication.
	SSHKeys    SSHKeyService
	Statistics StatisticService
}

// NewStorageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStorageService(opts ...option.RequestOption) (r StorageService) {
	r = StorageService{}
	r.Options = opts
	r.Locations = NewLocationService(opts...)
	r.ObjectStorages = NewObjectStorageService(opts...)
	r.SftpStorages = NewSftpStorageService(opts...)
	r.SSHKeys = NewSSHKeyService(opts...)
	r.Statistics = NewStatisticService(opts...)
	return
}
