package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type AdminSession struct {
	ent.Schema
}

func (AdminSession) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("admin_id", uuid.UUID{}).Optional(),
		field.String("token_hash").NotEmpty().Unique(),
		field.String("ip_address").Optional(),
		field.String("user_agent").Optional(),
		field.Time("expires_at").Optional().SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.Time("revoked_at").Optional().Nillable().SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.Time("last_activity_at").Optional().Nillable().SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.Time("created_at").Default(time.Now).SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
	}
}

func (AdminSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("admin", Admin.Type).
			Ref("sessions").
			Field("admin_id").
			Unique(),
	}
}
