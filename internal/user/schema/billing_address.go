package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "github.com/google/uuid"
    "time"
)

type BillingAddress struct {
    ent.Schema
}

func (BillingAddress) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.UUID("user_id", uuid.UUID{}),
        field.String("full_name").NotEmpty(),
        field.String("address_line1").NotEmpty(),
        field.String("address_line2").Optional(),
        field.String("city").NotEmpty(),
        field.String("state").NotEmpty(),
        field.String("postal_code").NotEmpty(),
        field.String("country").NotEmpty(),
        field.String("phone").NotEmpty(),
        field.Bool("is_default").Default(false),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional(),
        field.Time("deleted_at").Optional(),
    }
}

func (BillingAddress) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).Ref("billing_addresses").Field("user_id").Unique(),
    }
}
