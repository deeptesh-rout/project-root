package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "github.com/google/uuid"
    "time"
)

type BlockedSubdomain struct {
    ent.Schema
}

type Admin struct {
    ent.Schema
}

func (BlockedSubdomain) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.String("subdomain").NotEmpty(),
        field.UUID("domain_id", uuid.UUID{}),
        field.String("reason").NotEmpty(),
        field.UUID("created_by", uuid.UUID{}),
        field.Bool("is_available").Default(false),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional(),
        field.Time("deleted_at").Optional(),
    }
}

func (BlockedSubdomain) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("domain", Domain.Type).Ref("blocked_subdomains").Field("domain_id").Unique(),
        edge.From("creator", Admin.Type).Ref("blocked_subdomains_created").Field("created_by").Unique(),
    }
}
