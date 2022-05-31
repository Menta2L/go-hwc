package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Cpu holds the schema definition for the Cpu entity.
type Cpu struct {
	ent.Schema
}

// Fields of the Cpu.
func (Cpu) Fields() []ent.Field {
	return []ent.Field{
		field.String("vendor_id"),
		field.String("family"),
		field.String("model"),
		field.String("model_name"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Cpu.
func (Cpu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("host_id", Host.Type).
			Ref("cpu_id").
			Unique(),
	}
}
