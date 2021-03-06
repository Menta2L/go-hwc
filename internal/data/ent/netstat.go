// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/menta2l/go-hwc/internal/data/ent/host"
	"github.com/menta2l/go-hwc/internal/data/ent/netstat"
)

// Netstat is the model entity for the Netstat schema.
type Netstat struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Addr holds the value of the "addr" field.
	Addr string `json:"addr,omitempty"`
	// Port holds the value of the "port" field.
	Port uint64 `json:"port,omitempty"`
	// Proto holds the value of the "proto" field.
	Proto string `json:"proto,omitempty"`
	// Process holds the value of the "process" field.
	Process string `json:"process,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NetstatQuery when eager-loading is set.
	Edges        NetstatEdges `json:"edges"`
	host_netstat *string
}

// NetstatEdges holds the relations/edges for other nodes in the graph.
type NetstatEdges struct {
	// HostID holds the value of the host_id edge.
	HostID *Host `json:"host_id,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// HostIDOrErr returns the HostID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NetstatEdges) HostIDOrErr() (*Host, error) {
	if e.loadedTypes[0] {
		if e.HostID == nil {
			// The edge host_id was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: host.Label}
		}
		return e.HostID, nil
	}
	return nil, &NotLoadedError{edge: "host_id"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Netstat) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case netstat.FieldID, netstat.FieldPort:
			values[i] = new(sql.NullInt64)
		case netstat.FieldAddr, netstat.FieldProto, netstat.FieldProcess:
			values[i] = new(sql.NullString)
		case netstat.FieldCreatedAt, netstat.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case netstat.ForeignKeys[0]: // host_netstat
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Netstat", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Netstat fields.
func (n *Netstat) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case netstat.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case netstat.FieldAddr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field addr", values[i])
			} else if value.Valid {
				n.Addr = value.String
			}
		case netstat.FieldPort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field port", values[i])
			} else if value.Valid {
				n.Port = uint64(value.Int64)
			}
		case netstat.FieldProto:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field proto", values[i])
			} else if value.Valid {
				n.Proto = value.String
			}
		case netstat.FieldProcess:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field process", values[i])
			} else if value.Valid {
				n.Process = value.String
			}
		case netstat.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				n.CreatedAt = value.Time
			}
		case netstat.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = value.Time
			}
		case netstat.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host_netstat", values[i])
			} else if value.Valid {
				n.host_netstat = new(string)
				*n.host_netstat = value.String
			}
		}
	}
	return nil
}

// QueryHostID queries the "host_id" edge of the Netstat entity.
func (n *Netstat) QueryHostID() *HostQuery {
	return (&NetstatClient{config: n.config}).QueryHostID(n)
}

// Update returns a builder for updating this Netstat.
// Note that you need to call Netstat.Unwrap() before calling this method if this Netstat
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Netstat) Update() *NetstatUpdateOne {
	return (&NetstatClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the Netstat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Netstat) Unwrap() *Netstat {
	tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Netstat is not a transactional entity")
	}
	n.config.driver = tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Netstat) String() string {
	var builder strings.Builder
	builder.WriteString("Netstat(")
	builder.WriteString(fmt.Sprintf("id=%v", n.ID))
	builder.WriteString(", addr=")
	builder.WriteString(n.Addr)
	builder.WriteString(", port=")
	builder.WriteString(fmt.Sprintf("%v", n.Port))
	builder.WriteString(", proto=")
	builder.WriteString(n.Proto)
	builder.WriteString(", process=")
	builder.WriteString(n.Process)
	builder.WriteString(", created_at=")
	builder.WriteString(n.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(n.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Netstats is a parsable slice of Netstat.
type Netstats []*Netstat

func (n Netstats) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
