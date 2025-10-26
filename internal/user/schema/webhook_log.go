package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
    "time"
)

// WebhookLog holds the schema definition for webhook logs.
type WebhookLog struct {
    ent.Schema
}

// Fields of WebhookLog.
func (WebhookLog) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
        field.Enum("source").Values("phonepe", "razorpay", "dns_provider", "email_provider"),
        field.String("event").NotEmpty(),
        field.JSON("payload", map[string]interface{}{}),
        field.String("signature").Optional(),
        field.Enum("status").Values("pending", "processing", "processed", "failed", "ignored").Default("pending"),
        field.Enum("related_entity_type").Values("payment", "invoice", "subdomain", "email").Optional(),
        field.UUID("related_entity_id", uuid.UUID{}).Optional(),
        field.Int("processing_attempts").Default(0),
        field.String("error_message").Optional(),
        field.Time("processed_at").Optional(),
        field.Time("received_at").Default(time.Now).Immutable(),
        field.Time("updated_at").Optional().UpdateDefault(time.Now),
    }
}

// Edges of WebhookLog.
func (WebhookLog) Edges() []ent.Edge {
    return nil
}
