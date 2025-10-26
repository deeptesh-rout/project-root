package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

type Notification struct {
    ent.Schema
}

func (Notification) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.UUID("user_id", uuid.UUID{}),
        field.Enum("type").Values("info", "warning", "success", "error").Default("info"),
        field.Enum("category").Values("account", "billing", "domain", "support", "system"),
        field.String("title").NotEmpty(),
        field.String("description").NotEmpty(),
        field.String("action_url").Optional(),
        field.Time("read_at").Optional(),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
        field.Time("deleted_at").Optional(),
    }
}

func (Notification) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).Ref("notifications").Field("user_id").Unique(),
    }
}
