package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
	"project-root/internal/admin/schema" 
)

type VerificationToken struct {
	ent.Schema
}

func (VerificationToken) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Enum("token_type").Values("email_verification", "password_reset", "admin_invite"),
		field.UUID("user_id", uuid.UUID{}).Optional(),
		field.UUID("admin_id", uuid.UUID{}).Optional(),
		field.String("token_hash").NotEmpty().Unique(),
		field.Time("expires_at").SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.Time("used_at").Optional().Nillable().SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.String("ip_address").Optional(),
		field.Time("created_at").Default(time.Now).SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
	}
}

func (VerificationToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", schema.User.Type).
			Ref("verification_tokens").
			Field("user_id").
			Unique(),
		edge.From("admin", schema.Admin.Type).
			Ref("verification_tokens_admin").
			Field("admin_id").
			Unique(),
	}
}
