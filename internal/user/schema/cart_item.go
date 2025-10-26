package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "github.com/google/uuid"
    "time"
)

type CartItem struct {
    ent.Schema
}

func (CartItem) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.UUID("cart_id", uuid.UUID{}),
        field.Enum("item_type").Values("subdomain_purchase"),
        field.UUID("item_id", uuid.UUID{}),
        field.Int("quantity").Default(1),
        field.Float("price_at_addition").Positive(),
        field.Time("added_at").Default(time.Now).Immutable(),
    }
}

func (CartItem) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("cart", Cart.Type).Ref("items").Field("cart_id").Unique(),
    }
}
