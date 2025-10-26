package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Role schema
type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("name").
			NotEmpty().
			Unique(),
		field.String("slug").
			NotEmpty().
			Unique(),
		field.String("description").
			Optional().
			Nillable(),
		field.Bool("is_system_role").
			Default(false),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			SchemaType(map[string]string{"postgres": "timestamptz"}),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{"postgres": "timestamptz"}),
		field.Time("deleted_at").
			Optional().
			Nillable().
			SchemaType(map[string]string{"postgres": "timestamptz"}),
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("admin_roles", AdminRole.Type),       // A role can be assigned to multiple admins
		edge.To("role_permissions", RolePermission.Type), // Role can have multiple permissions
	}
}

// ============================================
// RolePermission schema
// ============================================
type RolePermission struct {
	ent.Schema
}

func (RolePermission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.UUID("role_id", uuid.UUID{}),
		field.UUID("permission_id", uuid.UUID{}),
		field.Time("granted_at").
			Default(time.Now).
			Immutable().
			SchemaType(map[string]string{"postgres": "timestamptz"}),
		field.UUID("granted_by", uuid.UUID{}),
		field.Time("revoked_at").
			Optional().
			Nillable().
			SchemaType(map[string]string{"postgres": "timestamptz"}),
		field.UUID("revoked_by", uuid.UUID{}).
			Optional().
			Nillable(),
	}
}

func (RolePermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).
			Field("role_id").
			Ref("role_permissions").
			Unique().
			Required(),
		edge.From("permission", Permission.Type).
			Field("permission_id").
			Ref("role_permissions").
			Unique().
			Required(),
		edge.From("granted_admin", Admin.Type).
			Field("granted_by").
			Ref("role_permissions_granted").
			Unique(),
		edge.From("revoked_admin", Admin.Type).
			Field("revoked_by").
			Ref("role_permissions_revoked").
			Unique(),
	}
}

// ============================================
// Permission schema is defined in internal/admin/schema/permission.go
// Duplicate stub removed to avoid redeclaration.
// ============================================

// AdminRole schema is defined in internal/admin/schema/admin_role.go to avoid duplicate declarations.
