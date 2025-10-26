package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

type InvoiceItem struct {
    ent.Schema
}

func (InvoiceItem) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.UUID("invoice_id", uuid.UUID{}),
        field.Enum("item_type").Values("subdomain_purchase", "subdomain_renewal", "custom"),
        field.UUID("item_id", uuid.UUID{}).Optional(),
        field.String("title").NotEmpty(),
        field.String("description").NotEmpty(),
        field.Int("quantity").Default(1),
        field.Float("unit_price"),
        field.Float("discount_amount").Default(0),
        field.Float("total_amount"),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
    }
}

func (InvoiceItem) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("invoice", Invoice.Type).
            Ref("items").
            Field("invoice_id").
            Unique(),
    }
}
