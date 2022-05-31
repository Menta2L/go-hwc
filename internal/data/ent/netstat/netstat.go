// Code generated by entc, DO NOT EDIT.

package netstat

import (
	"time"
)

const (
	// Label holds the string label denoting the netstat type in the database.
	Label = "netstat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAddr holds the string denoting the addr field in the database.
	FieldAddr = "addr"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// FieldProto holds the string denoting the proto field in the database.
	FieldProto = "proto"
	// FieldProcess holds the string denoting the process field in the database.
	FieldProcess = "process"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeHostID holds the string denoting the host_id edge name in mutations.
	EdgeHostID = "host_id"
	// Table holds the table name of the netstat in the database.
	Table = "netstats"
	// HostIDTable is the table that holds the host_id relation/edge.
	HostIDTable = "netstats"
	// HostIDInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostIDInverseTable = "hosts"
	// HostIDColumn is the table column denoting the host_id relation/edge.
	HostIDColumn = "host_netstat_id"
)

// Columns holds all SQL columns for netstat fields.
var Columns = []string{
	FieldID,
	FieldAddr,
	FieldPort,
	FieldProto,
	FieldProcess,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "netstats"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"host_netstat_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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