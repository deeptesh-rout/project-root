package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

type Payment struct {
    ent.Schema
}

func (Payment) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.UUID("invoice_id", uuid.UUID{}),
        field.UUID("user_id", uuid.UUID{}),
        field.Float("amount"),
        field.Enum("currency").Values("INR").Default("INR"),
        field.Enum("status").Values("pending", "processing", "success", "failed", "refunded", "partially_refunded").Default("pending"),
        field.Enum("payment_method").Values("gateway"),
        field.Enum("payment_gateway").Values("phonepe", "razorpay", "stripe").Optional(),
        field.String("order_id").NotEmpty(),
        field.String("gateway_order_id").Optional(),
        field.JSON("gateway_response", map[string]interface{}{}).Optional(),
        field.String("failure_reason").Optional(),
        field.Time("confirmed_at").Optional(),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
    }
}

func (Payment) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("invoice", Invoice.Type).Ref("payments").Field("invoice_id").Unique(),
        edge.From("user", User.Type).Ref("payments").Field("user_id").Unique(),
        edge.To("refunds", Refund.Type).StorageKey(edge.Column("payment_id")),
    }
}
