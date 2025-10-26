package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// EntMixin provides a simple base mixin for all user-related schemas.
type EntMixin struct{}

func (EntMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			DefaultNow().
			Immutable(),
		field.Time("updated_at").
			Optional().
			Nillable(),
	}
}
