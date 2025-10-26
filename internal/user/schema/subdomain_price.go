package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

type SubdomainPrice struct {
    ent.Schema
}

func (SubdomainPrice) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.UUID("domain_id", uuid.UUID{}),
        field.Int("min_length"),
        field.Int("max_length"),
        field.Enum("character_pattern").Values("standard", "premium", "numeric", "single_char", "emoji"),
        field.Float("price"),
        field.Bool("is_premium").Default(false),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
        field.Time("deleted_at").Optional(),
    }
}

func (SubdomainPrice) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("domain", Domain.Type).Ref("subdomain_prices").Field("domain_id").Unique(),
        edge.To("subdomains", Subdomain.Type).StorageKey(edge.Column("price_id")),
    }
}
