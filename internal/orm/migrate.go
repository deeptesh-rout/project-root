package orm

import (
    "context"
    "log"
    "time"

    //"project-root/internal/orm/ent"
    _ "github.com/lib/pq"
)

func Migrate() error {
    client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=mydb password=secret sslmode=disable")
    if err != nil {
        return err
    }
    defer client.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
    defer cancel()

    if err := client.Schema.Create(ctx); err != nil {
        return err
    }

    log.Println("Database schema migrated successfully")
    return nil
}
