package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

type AdminAuditLog struct {
	ent.Schema
}

func (AdminAuditLog) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("admin_id", uuid.UUID{}).Optional(),
		field.Enum("actor_type").Values("admin", "system").Default("admin"),
		field.Enum("resource_type").Values(
			"user", "domain", "subdomain", "blocked_subdomains", "invoice",
			"payment", "role", "permission", "admin", "role_permission",
			"admin_permission", "admin_role",
		),
		field.UUID("resource_id", uuid.UUID{}),
		field.Enum("action").Values("create", "read", "update", "delete", "approve", "reject", "export", "login", "logout"),
		field.JSON("changes", map[string]interface{}{}).Optional(),
		field.String("ip_address").Optional(),
		field.Text("user_agent").Optional(),
		field.Time("created_at").Default(time.Now).SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
	}
}

func (AdminAuditLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("admin", Admin.Type).
			Ref("audit_logs").
			Field("admin_id").
			Unique(),
	}
}
