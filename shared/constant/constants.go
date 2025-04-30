// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	"encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type AnySubnet string       // Always "any_subnet"
type Existing string        // Always "existing"
type External string        // Always "external"
type Fixed string           // Always "fixed"
type Floating string        // Always "floating"
type Image string           // Always "image"
type IPAddress string       // Always "ip_address"
type New string             // Always "new"
type NewVolume string       // Always "new-volume"
type Nfs string             // Always "NFS"
type Port string            // Always "port"
type ReservedFixedIP string // Always "reserved_fixed_ip"
type Snapshot string        // Always "snapshot"
type Subnet string          // Always "subnet"
type VastShareType string   // Always "vast_share_type"

func (c AnySubnet) Default() AnySubnet             { return "any_subnet" }
func (c Existing) Default() Existing               { return "existing" }
func (c External) Default() External               { return "external" }
func (c Fixed) Default() Fixed                     { return "fixed" }
func (c Floating) Default() Floating               { return "floating" }
func (c Image) Default() Image                     { return "image" }
func (c IPAddress) Default() IPAddress             { return "ip_address" }
func (c New) Default() New                         { return "new" }
func (c NewVolume) Default() NewVolume             { return "new-volume" }
func (c Nfs) Default() Nfs                         { return "NFS" }
func (c Port) Default() Port                       { return "port" }
func (c ReservedFixedIP) Default() ReservedFixedIP { return "reserved_fixed_ip" }
func (c Snapshot) Default() Snapshot               { return "snapshot" }
func (c Subnet) Default() Subnet                   { return "subnet" }
func (c VastShareType) Default() VastShareType     { return "vast_share_type" }

func (c AnySubnet) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Existing) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c External) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Fixed) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Floating) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Image) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c IPAddress) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c New) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c NewVolume) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Nfs) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Port) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c ReservedFixedIP) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Snapshot) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Subnet) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c VastShareType) MarshalJSON() ([]byte, error)   { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return json.Marshal(string(v))
}
