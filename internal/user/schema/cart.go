package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "github.com/google/uuid"
    "time"
)

type Cart struct {
    ent.Schema
}

func (Cart) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.UUID("user_id", uuid.UUID{}).Unique(),
        field.Time("updated_at").Optional(),
        field.Time("created_at").Default(time.Now).Immutable(),
    }
}

func (Cart) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).Ref("cart").Field("user_id").Unique(),
        edge.To("items", CartItem.Type).StorageKey(edge.Column("cart_id")),
    }
}
