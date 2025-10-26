package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

type Refund struct {
    ent.Schema
}

func (Refund) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.UUID("payment_id", uuid.UUID{}),
        field.UUID("invoice_id", uuid.UUID{}),
        field.UUID("user_id", uuid.UUID{}),
        field.Enum("refund_type").Values("full", "partial"),
        field.Float("amount"),
        field.String("reason").NotEmpty(),
        field.Enum("refund_method").Values("original"),
        field.String("gateway_refund_id").Optional(),
        field.JSON("gateway_response", map[string]interface{}{}).Optional(),
        field.Enum("status").Values("pending", "processing", "completed", "failed").Default("pending"),
        field.UUID("processed_by", uuid.UUID{}).Optional(),
        field.Time("processed_at").Optional(),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
    }
}

func (Refund) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("payment", Payment.Type).Ref("refunds").Field("payment_id").Unique(),
        edge.From("invoice", Invoice.Type).Ref("refunds").Field("invoice_id").Unique(),
        edge.From("user", User.Type).Ref("refunds").Field("user_id").Unique(),
        edge.From("processed_by_admin", Admin.Type).Ref("refunds_processed").Field("processed_by").Unique(),
    }
}
