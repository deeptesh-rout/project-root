package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

// TwoFactorBackupCode holds the schema definition for backup codes.
type TwoFactorBackupCode struct {
    ent.Schema
}

// Fields of TwoFactorBackupCode.
func (TwoFactorBackupCode) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.UUID("user_id", uuid.UUID{}),
        field.String("code_hash").NotEmpty(),
        field.Time("used_at").Optional(),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
    }
}

// Edges of TwoFactorBackupCode.
func (TwoFactorBackupCode) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).Ref("backup_codes").Field("user_id").Unique(),
    }
}
