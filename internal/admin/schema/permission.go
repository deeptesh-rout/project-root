package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Permission struct {
	ent.Schema
}

func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("resource").NotEmpty(),
		field.Enum("action").Values("CREATE", "READ", "UPDATE", "DELETE", "APPROVE", "EXPORT"),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now).SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
	}
}

func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role_permissions", RolePermission.Type),
		edge.To("admin_permissions", AdminPermission.Type),
	}
}
