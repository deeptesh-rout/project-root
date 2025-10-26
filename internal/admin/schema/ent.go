package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// EntMixin provides shared fields for admin module schemas.
type EntMixin struct{}


func (EntMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Optional().
			Nillable(),
	}
}
