package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

// Coupon holds the schema definition for the Coupon entity.
type Coupon struct {
    ent.Schema
}

// Fields of the Coupon.
func (Coupon) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.String("code").Unique().NotEmpty(),
        field.String("description").Optional(),
        field.Enum("discount_type").Values("percentage", "fixed_amount"),
        field.Float("discount_value"),
        field.Float("min_purchase_amount").Default(0),
        field.Float("max_discount_amount").Optional(),
        field.Int("max_uses").Optional(),
        field.Int("current_uses").Default(0),
        field.Int("max_uses_per_user").Default(1),
        field.Enum("applicable_to").Values("all", "subdomains", "renewals").Default("all"),
        field.Bool("stackable").Default(false),
        field.Time("valid_from"),
        field.Time("valid_until"),
        field.Bool("is_active").Default(true),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
        field.Time("deleted_at").Optional(),
    }
}

// Edges of the Coupon.
func (Coupon) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("usages", CouponUsage.Type).StorageKey(edge.Column("coupon_id")),
    }
}
