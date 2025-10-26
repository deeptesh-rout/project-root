package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type SystemSetting struct {
	ent.Schema
}

func (SystemSetting) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("key").NotEmpty().Unique(),
		field.JSON("value", map[string]interface{}{}).Default(map[string]interface{}{}),
		field.Enum("data_type").Values("string", "number", "boolean", "json"),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now).SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).SchemaType(map[string]string{
			"postgres": "timestamptz",
		}),
	}
}

func (SystemSetting) Edges() []ent.Edge {
	return nil
}
