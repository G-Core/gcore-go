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
type IPAddress string // Always "ip_address"
type New string       // Always "new"
type Port string      // Always "port"
type Subnet string    // Always "subnet"

func (c AnySubnet) Default() AnySubnet { return "any_subnet" }
func (c External) Default() External   { return "external" }
func (c IPAddress) Default() IPAddress { return "ip_address" }
func (c New) Default() New             { return "new" }
func (c Port) Default() Port           { return "port" }
func (c Subnet) Default() Subnet       { return "subnet" }

func (c AnySubnet) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c External) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c IPAddress) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c New) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Port) MarshalJSON() ([]byte, error)      { return marshalString(c) }
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
