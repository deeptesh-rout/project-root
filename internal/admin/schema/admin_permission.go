package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

type AdminPermission struct {
	ent.Schema
}

func (AdminPermission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("admin_id", uuid.UUID{}).Optional(),
		field.UUID("permission_id", uuid.UUID{}).Optional(),
		field.Time("granted_at").Default(time.Now).SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.UUID("granted_by", uuid.UUID{}).Optional(),
		field.Time("revoked_at").Optional().Nillable().SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.UUID("revoked_by", uuid.UUID{}).Optional(),
	}
}

func (AdminPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("admin", Admin.Type).
			Ref("permissions").
			Field("admin_id").
			Unique(),
		edge.From("permission", Permission.Type).
			Ref("admin_permissions").
			Field("permission_id").
			Unique(),
		edge.From("granted_by_admin", Admin.Type).
			Ref("granted_admin_permissions").
			Field("granted_by").
			Unique(),
		edge.From("revoked_by_admin", Admin.Type).
			Ref("revoked_admin_permissions").
			Field("revoked_by").
			Unique(),
	}
}
