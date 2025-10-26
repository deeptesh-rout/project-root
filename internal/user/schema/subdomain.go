package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

type Subdomain struct {
    ent.Schema
}

func (Subdomain) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.String("name").NotEmpty(),
        field.UUID("user_id", uuid.UUID{}),
        field.UUID("domain_id", uuid.UUID{}),
        field.UUID("price_id", uuid.UUID{}),
        field.Enum("status").Values("active", "disabled", "suspended", "expired").Default("active"),
        field.String("suspension_reason").Optional(),
        field.Time("suspended_at").Optional(),
        field.String("disabled_reason").Optional(),
        field.Time("disabled_at").Optional(),
        field.Bool("auto_renew").Default(true),
        field.Time("expires_at"),
        field.Time("renewal_reminder_sent_at").Optional(),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
    }
}

func (Subdomain) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).Ref("subdomains").Field("user_id").Unique(),
        edge.From("domain", Domain.Type).Ref("subdomains").Field("domain_id").Unique(),
        edge.From("price", SubdomainPrice.Type).Ref("subdomains").Field("price_id").Unique(),
        edge.To("history", SubdomainHistory.Type).StorageKey(edge.Column("subdomain_id")),
        edge.To("dns_records", DNSRecord.Type).StorageKey(edge.Column("subdomain_id")),
    }
}
