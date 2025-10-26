package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "github.com/google/uuid"
    "time"
)

type AuditLog struct {
    ent.Schema
}

func (AuditLog) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.UUID("user_id", uuid.UUID{}).Optional(),
        field.Enum("resource_type").Values("user", "subdomain", "invoice", "payment", "dns_record", "cart"),
        field.UUID("resource_id", uuid.UUID{}),
        field.Enum("action").Values("create", "read", "update", "delete", "login", "logout", "verify"),
        field.JSON("changes", map[string]interface{}{}).Optional(),
        field.String("ip_address").Optional(),
        field.String("user_agent").Optional(),
        field.Time("created_at").Default(time.Now).Immutable(),
    }
}

func (AuditLog) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).Ref("audit_logs").Field("user_id").Unique(),
    }
}
