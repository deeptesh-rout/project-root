package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

type Admin struct {
	ent.Schema
}

func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("email").NotEmpty().Unique(),
		field.String("password_hash").NotEmpty(),
		field.String("firstname").NotEmpty(),
		field.String("lastname").NotEmpty(),
		field.Time("email_verified_at").Optional().Nillable().SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.JSON("preferences", map[string]interface{}{}).Optional(),
		field.Bool("is_system_admin").Default(false),
		field.Enum("status").Values("pending", "active", "blocked", "deleted").Default("pending"),
		field.Time("last_login_at").Optional().Nillable().SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.String("last_login_ip").Optional(),
		field.String("last_login_device").Optional(),
		field.Time("password_changed_at").Optional().Nillable().SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.Bool("is_two_factor_enabled").Default(false),
		field.Enum("two_factor_method").Values("totp", "email").Optional(),
		field.String("two_factor_secret").Optional(),
		field.Time("deleted_at").Optional().Nillable().SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.String("deletion_reason").Optional(),
		field.UUID("created_by", uuid.UUID{}).Optional(),
		field.UUID("deleted_by", uuid.UUID{}).Optional(),
		field.Time("created_at").Default(time.Now).SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
	}
}

func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sessions", AdminSession.Type),
		edge.To("roles", AdminRole.Type),
		edge.To("permissions", AdminPermission.Type),
		edge.To("audit_logs", AdminAuditLog.Type),
		// explicit referential edges
		edge.From("created_by_admin", Admin.Type).
			Ref("created_admins").
			Field("created_by").
			Unique(),
		edge.From("deleted_by_admin", Admin.Type).
			Ref("deleted_admins").
			Field("deleted_by").
			Unique(),
	}
}
