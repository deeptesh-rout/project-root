package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

type Domain struct {
    ent.Schema
}

func (Domain) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.String("name").Unique().NotEmpty(),
        field.Enum("status").Values("disabled", "enabled").Default("disabled"),
        field.String("status_reason").Optional(),
        field.Bool("registration_enabled").Default(true),
        field.Time("deleted_at").Optional(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
        field.Time("created_at").Default(time.Now).Immutable(),
    }
}

func (Domain) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("subdomain_prices", SubdomainPrice.Type).StorageKey(edge.Column("domain_id")),
        edge.To("subdomains", Subdomain.Type).StorageKey(edge.Column("domain_id")),
        edge.To("blocked_subdomains", BlockedSubdomain.Type).StorageKey(edge.Column("domain_id")),
    }
}
