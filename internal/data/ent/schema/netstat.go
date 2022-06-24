package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Sockets holds the schema definition for the Sockets entity.
type Netstat struct {
	ent.Schema
}

// Fields of the Sockets.
func (Netstat) Fields() []ent.Field {
	return []ent.Field{
		field.String("addr"),
		field.Uint64("port"),
		field.String("proto"),
		field.String("process"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Sockets.
func (Netstat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("host_id", Host.Type).
			Ref("netstat").
			Unique().
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Card cannot be created without its owner.
			Required(),
	}
}
