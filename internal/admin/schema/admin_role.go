package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

type AdminRole struct {
	ent.Schema
}

func (AdminRole) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("admin_id", uuid.UUID{}).Optional(),
		field.UUID("role_id", uuid.UUID{}).Optional(),
		field.Time("assigned_at").Default(time.Now).SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.UUID("assigned_by", uuid.UUID{}).Optional(),
		field.Time("revoked_at").Optional().Nillable().SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.UUID("revoked_by", uuid.UUID{}).Optional(),
	}
}

func (AdminRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("admin", Admin.Type).
			Ref("roles").
			Field("admin_id").
			Unique(),
		edge.From("role", Role.Type).
			Ref("admin_roles").
			Field("role_id").
			Unique(),
		edge.From("assigned_by_admin", Admin.Type).
			Ref("assigned_admin_roles").
			Field("assigned_by").
			Unique(),
		edge.From("revoked_by_admin", Admin.Type).
			Ref("revoked_admin_roles").
			Field("revoked_by").
			Unique(),
	}
}
