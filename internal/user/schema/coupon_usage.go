package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "github.com/google/uuid"
    "time"
)

type CouponUsage struct {
    ent.Schema
}

func (CouponUsage) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.UUID("coupon_id", uuid.UUID{}),
        field.UUID("user_id", uuid.UUID{}),
        field.UUID("invoice_id", uuid.UUID{}),
        field.Float("discount_applied").NotEmpty(),
        field.Time("used_at").Default(time.Now).Immutable(),
    }
}

func (CouponUsage) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("coupon", Coupon.Type).Ref("usages").Field("coupon_id").Unique(),
        edge.From("user", User.Type).Ref("coupon_usages").Field("user_id").Unique(),
        edge.From("invoice", Invoice.Type).Ref("coupon_usages").Field("invoice_id").Unique(),
    }
}
