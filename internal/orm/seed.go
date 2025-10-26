package orm

import (
	"context"
	"log"
	"time"

	//"project/internal/orm/ent"
)


func SeedData(client *ent.Client) {
	ctx := context.Background()

	log.Println("Seeding initial data...")

	// Create an admin
	admin, err := client.Admin.Create().
		SetEmail("admin@example.com").
		SetPassword("hashed_password").
		SetFirstname("Super").
		SetLastname("Admin").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating admin: %v", err)
	}
	log.Printf("Created admin: %s\n", admin.Email)

	
	user, err := client.User.Create().
		SetEmail("user@example.com").
		SetPassword("hashed_password").
		SetFirstname("Test").
		SetLastname("User").
		SetStatus("active").
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Printf("Created user: %s\n", user.Email)

	log.Println("Seed data inserted successfully.")
}
