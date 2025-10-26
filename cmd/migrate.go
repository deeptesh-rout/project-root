package main

import (
	"context"
	"log"
	"time"

	//userent "project-root/internal/user/ent"
	_ "github.com/lib/pq"
)

func migrate() error {

	client, err := userent.Open("postgres", "host=localhost port=5432 user=postgres dbname=mydb password=secret sslmode=disable")
	if err != nil {
		return err
	}
	defer func() {
		if cerr := client.Close(); cerr != nil {
			log.Printf("Error closing client: %v", cerr)
		}
	}()

	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	
	if err := client.Schema.Create(ctx); err != nil {
		return err
	}

	log.Println("Database schema migrated successfully")
	return nil
}
