package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/jackc/pgtype"
)

// Disk holds the schema definition for the Disk entity.
type Disk struct {
	ent.Schema
}

// Fields of the Disk.
func (Disk) Fields() []ent.Field {
	return []ent.Field{
		field.String("device"),
		field.String("mount"),
		field.String("fs_type"),
		field.Other("opts", &pgtype.TextArray{}).
			SchemaType(map[string]string{
				dialect.Postgres: "text[]",
			}),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Disk.
func (Disk) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("host_id", Host.Type).
			Ref("disk_id").
			Unique().
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Card cannot be created without its owner.
			Required(),
	}
}
