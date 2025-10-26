package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

type SubdomainHistory struct {
    ent.Schema
}

func (SubdomainHistory) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.UUID("subdomain_id", uuid.UUID{}),
        field.UUID("user_id", uuid.UUID{}),
        field.Int64("interval_seconds").Comment("Interval duration in seconds"),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
    }
}

func (SubdomainHistory) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("subdomain", Subdomain.Type).Ref("history").Field("subdomain_id").Unique(),
        edge.From("user", User.Type).Ref("subdomain_histories").Field("user_id").Unique(),
    }
}
