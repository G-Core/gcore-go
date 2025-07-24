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

type AICluster string                       // Always "ai_cluster"
type AIVirtualCluster string                // Always "ai_virtual_cluster"
type AnySubnet string                       // Always "any_subnet"
type Apptemplate string                     // Always "apptemplate"
type Backup string                          // Always "backup"
type Baremetal string                       // Always "baremetal"
type BasicVm string                         // Always "basic_vm"
type Bytes string                           // Always "bytes"
type Containers string                      // Always "containers"
type DbaasPostgreSQLConnectionPooler string // Always "dbaas_postgresql_connection_pooler"
type DbaasPostgreSQLCPU string              // Always "dbaas_postgresql_cpu"
type DbaasPostgreSQLMemory string           // Always "dbaas_postgresql_memory"
type DbaasPostgreSQLPublicNetwork string    // Always "dbaas_postgresql_public_network"
type DbaasPostgreSQLVolume string           // Always "dbaas_postgresql_volume"
type EgressTraffic string                   // Always "egress_traffic"
type Existing string                        // Always "existing"
type ExistingVolume string                  // Always "existing-volume"
type External string                        // Always "external"
type ExternalIP string                      // Always "external_ip"
type FileShare string                       // Always "file_share"
type Fixed string                           // Always "fixed"
type Floating string                        // Always "floating"
type Floatingip string                      // Always "floatingip"
type Functions string                       // Always "functions"
type FunctionsCalls string                  // Always "functions_calls"
type FunctionsTraffic string                // Always "functions_traffic"
type GB string                              // Always "GB"
type Gbminutes string                       // Always "gbminutes"
type GBs string                             // Always "GBS"
type GiB string                             // Always "GiB"
type Image string                           // Always "image"
type Inference string                       // Always "inference"
type Instance string                        // Always "instance"
type IPAddress string                       // Always "ip_address"
type LoadBalancer string                    // Always "load_balancer"
type LogIndex string                        // Always "log_index"
type Minutes string                         // Always "minutes"
type Mls string                             // Always "MLS"
type New string                             // Always "new"
type NewVolume string                       // Always "new-volume"
type Nfs string                             // Always "NFS"
type Port string                            // Always "port"
type ReservedFixedIP string                 // Always "reserved_fixed_ip"
type Snapshot string                        // Always "snapshot"
type Standard string                        // Always "standard"
type Start string                           // Always "start"
type Subnet string                          // Always "subnet"
type Vast string                            // Always "vast"
type VastShareType string                   // Always "vast_share_type"
type Volume string                          // Always "volume"

func (c AICluster) Default() AICluster               { return "ai_cluster" }
func (c AIVirtualCluster) Default() AIVirtualCluster { return "ai_virtual_cluster" }
func (c AnySubnet) Default() AnySubnet               { return "any_subnet" }
func (c Apptemplate) Default() Apptemplate           { return "apptemplate" }
func (c Backup) Default() Backup                     { return "backup" }
func (c Baremetal) Default() Baremetal               { return "baremetal" }
func (c BasicVm) Default() BasicVm                   { return "basic_vm" }
func (c Bytes) Default() Bytes                       { return "bytes" }
func (c Containers) Default() Containers             { return "containers" }
func (c DbaasPostgreSQLConnectionPooler) Default() DbaasPostgreSQLConnectionPooler {
	return "dbaas_postgresql_connection_pooler"
}
func (c DbaasPostgreSQLCPU) Default() DbaasPostgreSQLCPU       { return "dbaas_postgresql_cpu" }
func (c DbaasPostgreSQLMemory) Default() DbaasPostgreSQLMemory { return "dbaas_postgresql_memory" }
func (c DbaasPostgreSQLPublicNetwork) Default() DbaasPostgreSQLPublicNetwork {
	return "dbaas_postgresql_public_network"
}
func (c DbaasPostgreSQLVolume) Default() DbaasPostgreSQLVolume { return "dbaas_postgresql_volume" }
func (c EgressTraffic) Default() EgressTraffic                 { return "egress_traffic" }
func (c Existing) Default() Existing                           { return "existing" }
func (c ExistingVolume) Default() ExistingVolume               { return "existing-volume" }
func (c External) Default() External                           { return "external" }
func (c ExternalIP) Default() ExternalIP                       { return "external_ip" }
func (c FileShare) Default() FileShare                         { return "file_share" }
func (c Fixed) Default() Fixed                                 { return "fixed" }
func (c Floating) Default() Floating                           { return "floating" }
func (c Floatingip) Default() Floatingip                       { return "floatingip" }
func (c Functions) Default() Functions                         { return "functions" }
func (c FunctionsCalls) Default() FunctionsCalls               { return "functions_calls" }
func (c FunctionsTraffic) Default() FunctionsTraffic           { return "functions_traffic" }
func (c GB) Default() GB                                       { return "GB" }
func (c Gbminutes) Default() Gbminutes                         { return "gbminutes" }
func (c GBs) Default() GBs                                     { return "GBS" }
func (c GiB) Default() GiB                                     { return "GiB" }
func (c Image) Default() Image                                 { return "image" }
func (c Inference) Default() Inference                         { return "inference" }
func (c Instance) Default() Instance                           { return "instance" }
func (c IPAddress) Default() IPAddress                         { return "ip_address" }
func (c LoadBalancer) Default() LoadBalancer                   { return "load_balancer" }
func (c LogIndex) Default() LogIndex                           { return "log_index" }
func (c Minutes) Default() Minutes                             { return "minutes" }
func (c Mls) Default() Mls                                     { return "MLS" }
func (c New) Default() New                                     { return "new" }
func (c NewVolume) Default() NewVolume                         { return "new-volume" }
func (c Nfs) Default() Nfs                                     { return "NFS" }
func (c Port) Default() Port                                   { return "port" }
func (c ReservedFixedIP) Default() ReservedFixedIP             { return "reserved_fixed_ip" }
func (c Snapshot) Default() Snapshot                           { return "snapshot" }
func (c Standard) Default() Standard                           { return "standard" }
func (c Start) Default() Start                                 { return "start" }
func (c Subnet) Default() Subnet                               { return "subnet" }
func (c Vast) Default() Vast                                   { return "vast" }
func (c VastShareType) Default() VastShareType                 { return "vast_share_type" }
func (c Volume) Default() Volume                               { return "volume" }

func (c AICluster) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c AIVirtualCluster) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c AnySubnet) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c Apptemplate) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c Backup) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c Baremetal) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c BasicVm) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c Bytes) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c Containers) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c DbaasPostgreSQLConnectionPooler) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c DbaasPostgreSQLCPU) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c DbaasPostgreSQLMemory) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c DbaasPostgreSQLPublicNetwork) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c DbaasPostgreSQLVolume) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c EgressTraffic) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Existing) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c ExistingVolume) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c External) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c ExternalIP) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c FileShare) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c Fixed) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c Floating) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c Floatingip) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c Functions) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c FunctionsCalls) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c FunctionsTraffic) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c GB) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c Gbminutes) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c GBs) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c GiB) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c Image) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c Inference) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c Instance) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c IPAddress) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c LoadBalancer) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c LogIndex) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c Minutes) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c Mls) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c New) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c NewVolume) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c Nfs) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c Port) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c ReservedFixedIP) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Snapshot) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c Standard) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c Start) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c Subnet) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c Vast) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c VastShareType) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Volume) MarshalJSON() ([]byte, error)                          { return marshalString(c) }

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
