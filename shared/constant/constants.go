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

type AnySubnet string // Always "any_subnet"
type External string  // Always "external"
type Image string     // Always "image"
type IPAddress string // Always "ip_address"
type NewVolume string // Always "new-volume"
type Port string      // Always "port"
type Snapshot string  // Always "snapshot"
type Subnet string    // Always "subnet"

func (c AnySubnet) Default() AnySubnet { return "any_subnet" }
func (c External) Default() External   { return "external" }
func (c Image) Default() Image         { return "image" }
func (c IPAddress) Default() IPAddress { return "ip_address" }
func (c NewVolume) Default() NewVolume { return "new-volume" }
func (c Port) Default() Port           { return "port" }
func (c Snapshot) Default() Snapshot   { return "snapshot" }
func (c Subnet) Default() Subnet       { return "subnet" }

func (c AnySubnet) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c External) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c Image) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c IPAddress) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c NewVolume) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Port) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Snapshot) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c Subnet) MarshalJSON() ([]byte, error)    { return marshalString(c) }

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
