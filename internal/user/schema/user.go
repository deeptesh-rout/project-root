package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Unique(),
        field.String("email").Unique().NotEmpty(),
        field.String("password").NotEmpty(),
        field.String("firstname").NotEmpty(),
        field.String("lastname").NotEmpty(),
        field.Time("dob").Optional(),
        field.Enum("gender").Values("male", "female", "others").Optional(),
        field.String("phone_number").Optional(),
        field.Time("email_verified_at").Optional(),
        field.Time("phone_verified_at").Optional(),
        field.UUID("profile_pic_id", uuid.UUID{}).Optional(),
        field.JSON("preferences", map[string]interface{}{}).Optional(),
        field.Enum("status").Values("pending", "blocked", "deleted", "active").Default("pending"),
        field.Time("deleted_at").Optional(),
        field.String("deletion_reason").Optional(),
        field.Time("anonymized_at").Optional(),
        field.Time("last_login_at").Optional(),
        field.String("last_login_ip").Optional(),
        field.String("last_login_device").Optional(),
        field.Time("password_changed_at").Optional(),
        field.Bool("is_two_factor_enabled").Default(false),
        field.Enum("two_factor_method").Values("totp", "sms", "email").Optional(),
        field.String("two_factor_secret").Optional(),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("sessions", UserSession.Type).StorageKey(edge.Column("user_id")),
        edge.To("backup_codes", TwoFactorBackupCode.Type).StorageKey(edge.Column("user_id")),
    }
}
