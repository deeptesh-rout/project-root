package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

type Invoice struct {
    ent.Schema
}

func (Invoice) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.String("invoice_number").Unique().NotEmpty(),
        field.UUID("user_id", uuid.UUID{}),
        field.JSON("billing_address_snapshot", map[string]interface{}{}).NotEmpty(),
        field.Enum("status").Values("draft", "issued", "paid", "partially_paid", "cancelled", "refunded").Default("draft"),
        field.Enum("currency").Values("INR").Default("INR"),
        field.Float("subtotal"),
        field.Float("tax_rate"),
        field.Float("tax_amount"),
        field.Float("discount_amount").Default(0),
        field.Float("total_amount"),
        field.Float("amount_paid").Default(0),
        field.Float("amount_due"),
        field.String("notes").Optional(),
        field.String("customer_notes").Optional(),
        field.Time("due_date").Optional(),
        field.Time("issued_at").Optional(),
        field.Time("paid_at").Optional(),
        field.Time("cancelled_at").Optional(),
        field.String("cancellation_reason").Optional(),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
    }
}

func (Invoice) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).Ref("invoices").Field("user_id").Unique(),
        edge.To("items", InvoiceItem.Type).StorageKey(edge.Column("invoice_id")),
        edge.To("coupon_usages", CouponUsage.Type).StorageKey(edge.Column("invoice_id")),
        edge.To("payments", Payment.Type).StorageKey(edge.Column("invoice_id")),
        edge.To("refunds", Refund.Type).StorageKey(edge.Column("invoice_id")),
    }
}
