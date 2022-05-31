package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/jackc/pgtype"
)

// Network holds the schema definition for the Network entity.
type Network struct {
	ent.Schema
}

// Fields of the Network.
func (Network) Fields() []ent.Field {
	return []ent.Field{
		field.Int("idx"),
		field.Int("mtu"),
		field.String("name"),
		field.String("mac"),
		field.Other("flags", &pgtype.TextArray{}).
			SchemaType(map[string]string{
				dialect.Postgres: "text[]",
			}).Optional().Nillable(),
		field.Other("addrs", &pgtype.TextArray{}).
			SchemaType(map[string]string{
				dialect.Postgres: "text[]",
			}).Optional().Nillable(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Network.
func (Network) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("host_id", Host.Type).
			Ref("network_id").
			Unique().
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Card cannot be created without its owner.
			Required(),
	}
}
