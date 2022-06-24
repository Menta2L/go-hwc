package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Disk holds the schema definition for the Disk entity.
type Disk struct {
	ent.Schema
}

// Fields of the Disk.
func (Disk) Fields() []ent.Field {
	return []ent.Field{
		field.String("device"),
		field.String("Mountpoint").StorageKey("mount"),
		field.String("Fstype").StorageKey("fs_type"),
		field.JSON("opts", []string{}).
			Optional().Optional(),
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
			Ref("disk").
			Unique().
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Card cannot be created without its owner.
			Required(),
	}
}
