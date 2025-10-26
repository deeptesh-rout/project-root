package orm

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	//"project/internal/orm/ent"
)

func ConnectDB() (*ent.Client, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/project_db?sslmode=disable"
	}

	driver, err := sql.Open(dialect.Postgres, dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := driver.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed pinging postgres: %w", err)
	}

	client := ent.NewClient(ent.Driver(driver))
	log.Println("Connected to PostgreSQL database successfully")
	return client, nil
}

func CloseDB(client *ent.Client) {
	if err := client.Close(); err != nil {
		log.Printf("failed closing database connection: %v", err)
	} else {
		log.Println("Database connection closed")
	}
}
