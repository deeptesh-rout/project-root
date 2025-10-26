package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

type DNSRecord struct {
    ent.Schema
}

func (DNSRecord) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.UUID("subdomain_id", uuid.UUID{}),
        field.Enum("type").Values("A", "AAAA", "CNAME", "TXT", "MX", "NS", "SOA").Default("A"),
        field.String("name").NotEmpty(),
        field.String("value").NotEmpty(),
        field.Int("ttl").Default(3600),
        field.Int("priority").Optional(),
        field.Bool("is_proxied").Default(false),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
        field.Time("deleted_at").Optional(),
    }
}

func (DNSRecord) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("subdomain", Subdomain.Type).
            Ref("dns_records").
            Field("subdomain_id").
            Unique(),
    }
}
