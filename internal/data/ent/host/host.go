// Code generated by entc, DO NOT EDIT.

package host

import (
	"time"
)

const (
	// Label holds the string label denoting the host type in the database.
	Label = "host"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHostname holds the string denoting the hostname field in the database.
	FieldHostname = "hostname"
	// FieldOs holds the string denoting the os field in the database.
	FieldOs = "os"
	// FieldPlatform holds the string denoting the platform field in the database.
	FieldPlatform = "platform"
	// FieldPlatformFamily holds the string denoting the platform_family field in the database.
	FieldPlatformFamily = "platform_family"
	// FieldPlatformVersion holds the string denoting the platform_version field in the database.
	FieldPlatformVersion = "platform_version"
	// FieldKernelVersion holds the string denoting the kernel_version field in the database.
	FieldKernelVersion = "kernel_version"
	// FieldKernelArch holds the string denoting the kernel_arch field in the database.
	FieldKernelArch = "kernel_arch"
	// FieldVirtualizationSystem holds the string denoting the virtualization_system field in the database.
	FieldVirtualizationSystem = "virtualization_system"
	// FieldVirtualizationRole holds the string denoting the virtualization_role field in the database.
	FieldVirtualizationRole = "virtualization_role"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeCPUID holds the string denoting the cpu_id edge name in mutations.
	EdgeCPUID = "cpu_id"
	// EdgeNetworkID holds the string denoting the network_id edge name in mutations.
	EdgeNetworkID = "network_id"
	// EdgeNetstatID holds the string denoting the netstat_id edge name in mutations.
	EdgeNetstatID = "netstat_id"
	// EdgeDiskID holds the string denoting the disk_id edge name in mutations.
	EdgeDiskID = "disk_id"
	// Table holds the table name of the host in the database.
	Table = "hosts"
	// CPUIDTable is the table that holds the cpu_id relation/edge.
	CPUIDTable = "cpus"
	// CPUIDInverseTable is the table name for the Cpu entity.
	// It exists in this package in order to avoid circular dependency with the "cpu" package.
	CPUIDInverseTable = "cpus"
	// CPUIDColumn is the table column denoting the cpu_id relation/edge.
	CPUIDColumn = "host_cpu_id"
	// NetworkIDTable is the table that holds the network_id relation/edge.
	NetworkIDTable = "networks"
	// NetworkIDInverseTable is the table name for the Network entity.
	// It exists in this package in order to avoid circular dependency with the "network" package.
	NetworkIDInverseTable = "networks"
	// NetworkIDColumn is the table column denoting the network_id relation/edge.
	NetworkIDColumn = "host_network_id"
	// NetstatIDTable is the table that holds the netstat_id relation/edge.
	NetstatIDTable = "netstats"
	// NetstatIDInverseTable is the table name for the Netstat entity.
	// It exists in this package in order to avoid circular dependency with the "netstat" package.
	NetstatIDInverseTable = "netstats"
	// NetstatIDColumn is the table column denoting the netstat_id relation/edge.
	NetstatIDColumn = "host_netstat_id"
	// DiskIDTable is the table that holds the disk_id relation/edge.
	DiskIDTable = "disks"
	// DiskIDInverseTable is the table name for the Disk entity.
	// It exists in this package in order to avoid circular dependency with the "disk" package.
	DiskIDInverseTable = "disks"
	// DiskIDColumn is the table column denoting the disk_id relation/edge.
	DiskIDColumn = "host_disk_id"
)

// Columns holds all SQL columns for host fields.
var Columns = []string{
	FieldID,
	FieldHostname,
	FieldOs,
	FieldPlatform,
	FieldPlatformFamily,
	FieldPlatformVersion,
	FieldKernelVersion,
	FieldKernelArch,
	FieldVirtualizationSystem,
	FieldVirtualizationRole,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)