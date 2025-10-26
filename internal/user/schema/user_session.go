package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

// UserSession holds the schema definition for user sessions.
type UserSession struct {
    ent.Schema
}

// Fields of UserSession.
func (UserSession) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.UUID("user_id", uuid.UUID{}),
        field.String("token_hash").Unique().NotEmpty(),
        field.String("ip_address").NotEmpty(),
        field.String("user_agent").NotEmpty(),
        field.Time("expires_at"),
        field.Time("revoked_at").Optional(),
        field.Time("last_activity_at").Optional(),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
    }
}

// Edges of UserSession.
func (UserSession) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).Ref("sessions").Field("user_id").Unique(),
    }
}
